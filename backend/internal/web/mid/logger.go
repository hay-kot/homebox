package mid

import (
	"fmt"
	"net/http"

	"github.com/hay-kot/homebox/backend/pkgs/server"
	"github.com/rs/zerolog"
)

type statusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func Logger(log zerolog.Logger) server.Middleware {
	return func(next server.Handler) server.Handler {
		return server.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
			traceId := server.GetTraceID(r.Context())

			log.Info().
				Str("trace_id", traceId).
				Str("method", r.Method).
				Str("path", r.URL.Path).
				Str("remove_address", r.RemoteAddr).
				Msg("request started")

			record := &statusRecorder{ResponseWriter: w, Status: http.StatusOK}

			err := next.ServeHTTP(record, r)

			log.Info().
				Str("trace_id", traceId).
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

	atLeast6 := func(s string) string {
		for len(s) <= 6 {
			s += " "
		}
		return s
	}

	return func(next server.Handler) server.Handler {
		return server.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
			record := &statusRecorder{ResponseWriter: w, Status: http.StatusOK}

			err := next.ServeHTTP(record, r) // Blocks until the next handler returns.

			url := fmt.Sprintf("%s %s", r.RequestURI, r.Proto)

			log.Info().
				Str("trace_id", server.GetTraceID(r.Context())).
				Msgf("%s %s %s",
					bold(fmtCode(record.Status)),
					bold(orange(atLeast6(r.Method))),
					aqua(url),
				)

			return err
		})
	}
}
