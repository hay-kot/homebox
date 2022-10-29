package server

import (
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	return f(w, r)
}

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request) error
}

// ToHandler converts a function to a customer implementation of the Handler interface.
// that returns an error. This wrapper around the handler function and simply
// returns the nil in all cases
func ToHandler(handler http.Handler) Handler {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		handler.ServeHTTP(w, r)
		return nil
	})
}
