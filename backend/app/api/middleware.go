package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/hay-kot/homebox/backend/internal/services"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/homebox/backend/pkgs/server"
)

// mwAuthToken is a middleware that will check the database for a stateful token
// and attach it to the request context with the user, or return a 401 if it doesn't exist.
func (a *app) mwAuthToken(next server.Handler) server.Handler {
	return server.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		requestToken := r.Header.Get("Authorization")

		if requestToken == "" {
			return validate.NewRequestError(errors.New("Authorization header is required"), http.StatusUnauthorized)
		}

		requestToken = strings.TrimPrefix(requestToken, "Bearer ")
		usr, err := a.services.User.GetSelf(r.Context(), requestToken)

		// Check the database for the token
		if err != nil {
			return validate.NewRequestError(errors.New("Authorization header is required"), http.StatusUnauthorized)
		}

		r = r.WithContext(services.SetUserCtx(r.Context(), &usr, requestToken))
		return next.ServeHTTP(w, r)
	})
}
