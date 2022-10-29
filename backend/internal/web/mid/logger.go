package mid

import (
	"fmt"
	"net/http"

	"github.com/hay-kot/homebox/backend/pkgs/server"
	"github.com/rs/zerolog"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func Logger(log zerolog.Logger) server.Middleware {
	return func(next server.Handler) server.Handler {
		return server.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {

			log.Info().
				Str("trace_id", "TODO").
				Str("method", r.Method).
				Str("path", r.URL.Path).
				Str("remove_address", r.RemoteAddr).
				Msg("request started")

			record := &StatusRecorder{ResponseWriter: w, Status: http.StatusOK}

			err := next.ServeHTTP(record, r)

			log.Info().
				Str("trave_id", "TODO").
				Str("method", r.Method).
				Str("url", r.URL.Path).
				Str("remote_address", r.RemoteAddr).
				Int("status_code", record.Status).
				Msg("request completed")

			return err

		})
	}
}

func SugarLogger(log zerolog.Logger) server.Middleware {
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
	bold := func(s string) string { return "\033[1m" + s + "\033[0m" }

	return func(next server.Handler) server.Handler {
		return server.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {

			record := &StatusRecorder{ResponseWriter: w, Status: http.StatusOK}

			err := next.ServeHTTP(record, r) // Blocks until the next handler returns.

			scheme := "http"
			if r.TLS != nil {
				scheme = "https"
			}

			url := fmt.Sprintf("%s://%s%s %s", scheme, r.Host, r.RequestURI, r.Proto)

			log.Info().
				Msgf("%s  %s  %s",
					bold(fmtCode(record.Status)),
					bold(orange(""+r.Method+"")),
					aqua(url),
				)

			return err
		})
	}
}
