package server

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func testServer(t *testing.T, r http.Handler) *Server {
	svr := NewServer("127.0.0.1", "19245")

	go func() {
		svr.Start(r)
	}()

	ping := func() error {
		_, err := http.Get("http://127.0.0.1:19245")
		return err
	}

	for {
		if err := ping(); err == nil {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	return svr
}

func Test_ServerShutdown_Error(t *testing.T) {
	svr := NewServer("127.0.0.1", "19245")

	err := svr.Shutdown("test")
	assert.ErrorIs(t, err, ErrServerNotStarted)
}

func Test_ServerStarts_Error(t *testing.T) {
	svr := testServer(t, nil)

	err := svr.Start(nil)
	assert.ErrorIs(t, err, ErrServerAlreadyStarted)

	err = svr.Shutdown("test")
	assert.NoError(t, err)
}

func Test_ServerStarts(t *testing.T) {
	svr := testServer(t, nil)
	err := svr.Shutdown("test")
	assert.NoError(t, err)
}

func Test_GracefulServerShutdownWithWorkers(t *testing.T) {
	isFinished := false

	svr := testServer(t, nil)

	svr.Background(func() {
		time.Sleep(time.Second * 4)
		isFinished = true
	})

	err := svr.Shutdown("test")

	assert.NoError(t, err)
	assert.True(t, isFinished)

}

func Test_GracefulServerShutdownWithRequests(t *testing.T) {
	isFinished := false

	router := http.NewServeMux()

	// add long running handler func
	router.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 3)
		isFinished = true
	})

	svr := testServer(t, router)

	// Make request to "/test"
	go func() {
		http.Get("http://127.0.0.1:19245/test") // This is probably bad?
	}()

	time.Sleep(time.Second) // Hack to wait for the request to be made

	err := svr.Shutdown("test")
	assert.NoError(t, err)

	assert.True(t, isFinished)
}
