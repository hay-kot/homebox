package server

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type vkey int

const (
	// Key is the key for the server in the request context.
	key vkey = 1
)

type Values struct {
	TraceID string
}

func GetTraceID(ctx context.Context) string {
	v, ok := ctx.Value(key).(Values)
	if !ok {
		return ""
	}
	return v.TraceID
}

func (s *Server) toHttpHandler(handler Handler, mw ...Middleware) http.HandlerFunc {
	handler = wrapMiddleware(mw, handler)

	handler = wrapMiddleware(s.mw, handler)

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Add the trace ID to the context
		ctx = context.WithValue(ctx, key, Values{
			TraceID: uuid.NewString(),
		})

		err := handler.ServeHTTP(w, r.WithContext(ctx))

		if err != nil {
			if IsShutdownError(err) {
				s.Shutdown("SIGTERM")
			}
		}
	}
}

func (s *Server) handle(method, pattern string, handler Handler, mw ...Middleware) {
	h := s.toHttpHandler(handler, mw...)

	switch method {
	case http.MethodGet:
		s.mux.Get(pattern, h)
	case http.MethodPost:
		s.mux.Post(pattern, h)
	case http.MethodPut:
		s.mux.Put(pattern, h)
	case http.MethodDelete:
		s.mux.Delete(pattern, h)
	case http.MethodPatch:
		s.mux.Patch(pattern, h)
	case http.MethodHead:
		s.mux.Head(pattern, h)
	case http.MethodOptions:
		s.mux.Options(pattern, h)
	}
}

func (s *Server) Get(pattern string, handler Handler, mw ...Middleware) {
	s.handle(http.MethodGet, pattern, handler, mw...)
}

func (s *Server) Post(pattern string, handler Handler, mw ...Middleware) {
	s.handle(http.MethodPost, pattern, handler, mw...)
}

func (s *Server) Put(pattern string, handler Handler, mw ...Middleware) {
	s.handle(http.MethodPut, pattern, handler, mw...)
}

func (s *Server) Delete(pattern string, handler Handler, mw ...Middleware) {
	s.handle(http.MethodDelete, pattern, handler, mw...)
}

func (s *Server) Patch(pattern string, handler Handler, mw ...Middleware) {
	s.handle(http.MethodPatch, pattern, handler, mw...)
}

func (s *Server) Head(pattern string, handler Handler, mw ...Middleware) {
	s.handle(http.MethodHead, pattern, handler, mw...)
}

func (s *Server) Options(pattern string, handler Handler, mw ...Middleware) {
	s.handle(http.MethodOptions, pattern, handler, mw...)
}

func (s *Server) NotFound(handler Handler) {
	s.mux.NotFound(s.toHttpHandler(handler))
}
