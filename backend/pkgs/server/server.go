package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

var (
	ErrServerNotStarted     = errors.New("server not started")
	ErrServerAlreadyStarted = errors.New("server already started")
)

type Server struct {
	Host   string
	Port   string
	Worker Worker

	wg  sync.WaitGroup
	mux *chi.Mux

	// mw is the global middleware chain for the server.
	mw []Middleware

	started      bool
	activeServer *http.Server

	idleTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
}

func NewServer(opts ...Option) *Server {
	s := &Server{
		Host:         "localhost",
		Port:         "8080",
		mux:          chi.NewRouter(),
		Worker:       NewSimpleWorker(),
		idleTimeout:  30 * time.Second,
		readTimeout:  10 * time.Second,
		writeTimeout: 10 * time.Second,
	}

	for _, opt := range opts {
		err := opt(s)
		if err != nil {
			panic(err)
		}
	}

	return s
}

func (s *Server) Shutdown(sig string) error {
	if !s.started {
		return ErrServerNotStarted
	}
	fmt.Printf("Received %s signal, shutting down\n", sig)

	// Create a context with a 5-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.activeServer.Shutdown(ctx)
	s.started = false
	if err != nil {
		return err
	}

	fmt.Println("Http server shutdown, waiting for all tasks to finish")
	s.wg.Wait()

	return nil
}

func (s *Server) Start() error {
	if s.started {
		return ErrServerAlreadyStarted
	}

	s.activeServer = &http.Server{
		Addr:         s.Host + ":" + s.Port,
		Handler:      s.mux,
		IdleTimeout:  s.idleTimeout,
		ReadTimeout:  s.readTimeout,
		WriteTimeout: s.writeTimeout,
	}

	shutdownError := make(chan error)

	go func() {
		// Create a quit channel which carries os.Signal values.
		quit := make(chan os.Signal, 1)

		// Use signal.Notify() to listen for incoming SIGINT and SIGTERM signals and
		// relay them to the quit channel.
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// Read the signal from the quit channel. block until received
		sig := <-quit

		err := s.Shutdown(sig.String())
		if err != nil {
			shutdownError <- err
		}

		// Exit the application with a 0 (success) status code.
		os.Exit(0)
	}()

	s.started = true
	err := s.activeServer.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	fmt.Println("Server shutdown successfully")

	return nil
}

// Background starts a go routine that runs on the servers pool. In the event of a shutdown
// request, the server will wait until all open goroutines have finished before shutting down.
func (svr *Server) Background(task func()) {
	svr.wg.Add(1)
	svr.Worker.Add(func() {
		defer svr.wg.Done()
		task()
	})
}
