package mid

import (
	"net/http"

	"github.com/hay-kot/homebox/backend/internal/data/ent"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/homebox/backend/pkgs/server"
	"github.com/rs/zerolog"
)

func Errors(log zerolog.Logger) server.Middleware {
	return func(h server.Handler) server.Handler {
		return server.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
			err := h.ServeHTTP(w, r)
			if err != nil {
				var resp server.ErrorResponse
				var code int

				log.Err(err).
					Str("trace_id", server.GetTraceID(r.Context())).
					Msg("ERROR occurred")

				switch {
				case validate.IsUnauthorizedError(err):
					code = http.StatusUnauthorized
					resp = server.ErrorResponse{
						Error: "unauthorized",
					}
				case validate.IsInvalidRouteKeyError(err):
					code = http.StatusBadRequest
					resp = server.ErrorResponse{
						Error: err.Error(),
					}
				case validate.IsFieldError(err):
					code = http.StatusUnprocessableEntity

					fieldErrors := err.(validate.FieldErrors)
					resp.Error = "Validation Error"
					resp.Fields = map[string]string{}

					for _, fieldError := range fieldErrors {
						resp.Fields[fieldError.Field] = fieldError.Error
					}
				case validate.IsRequestError(err):
					requestError := err.(*validate.RequestError)
					resp.Error = requestError.Error()

					if requestError.Status == 0 {
						code = http.StatusBadRequest
					} else {
						code = requestError.Status
					}
				case ent.IsNotFound(err):
					resp.Error = "Not Found"
					code = http.StatusNotFound
				default:
					resp.Error = "Unknown Error"
					code = http.StatusInternalServerError
				}

				if err := server.Respond(w, code, resp); err != nil {
					return err
				}

				// If Showdown error, return error
				if server.IsShutdownError(err) {
					return err
				}
			}

			return nil
		})
	}
}
