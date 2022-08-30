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
)

// TODO: #2 Implement Go routine pool/job queue

var ErrServerNotStarted = errors.New("server not started")
var ErrServerAlreadyStarted = errors.New("server already started")

type Server struct {
	Host string
	Port string

	Worker Worker
	wg     sync.WaitGroup

	started      bool
	activeServer *http.Server
}

func NewServer(host, port string) *Server {
	return &Server{
		Host:   host,
		Port:   port,
		wg:     sync.WaitGroup{},
		Worker: NewSimpleWorker(),
	}
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

func (s *Server) Start(router http.Handler) error {
	if s.started {
		return ErrServerAlreadyStarted
	}

	s.activeServer = &http.Server{
		Addr:         s.Host + ":" + s.Port,
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
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
