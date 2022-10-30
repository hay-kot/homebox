package server

import (
	"net/http"
	"strings"
)

type Middleware func(Handler) Handler

// wrapMiddleware creates a new handler by wrapping middleware around a final
// handler. The middlewares' Handlers will be executed by requests in the order
// they are provided.
func wrapMiddleware(mw []Middleware, handler Handler) Handler {

	// Loop backwards through the middleware invoking each one. Replace the
	// handler with the new wrapped handler. Looping backwards ensures that the
	// first middleware of the slice is the first to be executed by requests.
	for i := len(mw) - 1; i >= 0; i-- {
		h := mw[i]
		if h != nil {
			handler = h(handler)
		}
	}

	return handler
}

// StripTrailingSlash is a middleware that will strip trailing slashes from the request path.
//
// Example: /api/v1/ -> /api/v1
func StripTrailingSlash() Middleware {
	return func(h Handler) Handler {
		return HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
			return h.ServeHTTP(w, r)
		})
	}
}
