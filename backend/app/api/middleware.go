package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hay-kot/content/backend/internal/config"
	"github.com/hay-kot/content/backend/internal/services"
	"github.com/hay-kot/content/backend/pkgs/server"
	"github.com/rs/zerolog/log"
)

func (a *app) setGlobalMiddleware(r *chi.Mux) {
	// =========================================================================
	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(mwStripTrailingSlash)

	// Use struct logger in production for requests, but use
	// pretty console logger in development.
	if a.conf.Mode == config.ModeDevelopment {
		r.Use(a.mwSummaryLogger)
	} else {
		r.Use(a.mwStructLogger)
	}
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
}

// mwAuthToken is a middleware that will check the database for a stateful token
// and attach it to the request context with the user, or return a 401 if it doesn't exist.
func (a *app) mwAuthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestToken := r.Header.Get("Authorization")

		if requestToken == "" {
			server.RespondUnauthorized(w)
			return
		}

		requestToken = strings.TrimPrefix(requestToken, "Bearer ")
		usr, err := a.services.User.GetSelf(r.Context(), requestToken)

		// Check the database for the token

		if err != nil {
			server.RespondUnauthorized(w)
			return
		}

		r = r.WithContext(services.SetUserCtx(r.Context(), usr, requestToken))

		next.ServeHTTP(w, r)
	})
}

// mwAdminOnly is a middleware that extends the mwAuthToken middleware to only allow
// requests from superusers.
// func (a *app) mwAdminOnly(next http.Handler) http.Handler {
// 	mw := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		usr := services.UseUserCtx(r.Context())

// 		if !usr.IsSuperuser {
// 			server.RespondUnauthorized(w)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// 	return a.mwAuthToken(mw)
// }

// mqStripTrailingSlash is a middleware that will strip trailing slashes from the request path.
func mwStripTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func (a *app) mwStructLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		record := &StatusRecorder{ResponseWriter: w, Status: http.StatusOK}
		next.ServeHTTP(record, r)

		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}

		url := fmt.Sprintf("%s://%s%s %s", scheme, r.Host, r.RequestURI, r.Proto)

		log.Info().
			Str("id", middleware.GetReqID(r.Context())).
			Str("url", url).
			Str("method", r.Method).
			Str("remote_addr", r.RemoteAddr).
			Int("status", record.Status).
			Msg(url)
	})
}

func (a *app) mwSummaryLogger(next http.Handler) http.Handler {
	bold := func(s string) string { return "\033[1m" + s + "\033[0m" }
	orange := func(s string) string { return "\033[33m" + s + "\033[0m" }
	aqua := func(s string) string { return "\033[36m" + s + "\033[0m" }
	red := func(s string) string { return "\033[31m" + s + "\033[0m" }
	green := func(s string) string { return "\033[32m" + s + "\033[0m" }

	fmtCode := func(code int) string {
		switch {
		case code >= 500:
			return red(fmt.Sprintf("%d", code))
		case code >= 400:
			return orange(fmt.Sprintf("%d", code))
		case code >= 300:
			return aqua(fmt.Sprintf("%d", code))
		default:
			return green(fmt.Sprintf("%d", code))
		}
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		record := &StatusRecorder{ResponseWriter: w, Status: http.StatusOK}
		next.ServeHTTP(record, r) // Blocks until the next handler returns.

		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}

		url := fmt.Sprintf("%s://%s%s %s", scheme, r.Host, r.RequestURI, r.Proto)

		log.Info().
			Msgf("%s  %s  %s",
				bold(orange(""+r.Method+"")),
				aqua(url),
				bold(fmtCode(record.Status)),
			)
	})
}
