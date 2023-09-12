package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/homebox/backend/pkgs/lookup"
	"github.com/hay-kot/httpkit/errchain"
)

type tokenHasKey struct {
	key string
}

var hashedToken = tokenHasKey{key: "hashedToken"}

type RoleMode int

const (
	RoleModeOr  RoleMode = 0
	RoleModeAnd RoleMode = 1
)

// mwRoles is a middleware that will validate the required roles are met. All roles
// are required to be met for the request to be allowed. If the user does not have
// the required roles, a 403 Forbidden will be returned.
//
// WARNING: This middleware _MUST_ be called after mwAuthToken or else it will panic
func (a *app) mwRoles(rm RoleMode, required ...string) errchain.Middleware {
	return func(next errchain.Handler) errchain.Handler {
		return errchain.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
			ctx := r.Context()

			maybeToken := ctx.Value(hashedToken)
			if maybeToken == nil {
				panic("mwRoles: token not found in context, you must call mwAuthToken before mwRoles")
			}

			token := maybeToken.(string)

			roles, err := a.repos.AuthTokens.GetRoles(r.Context(), token)
			if err != nil {
				return err
			}

		outer:
			switch rm {
			case RoleModeOr:
				for _, role := range required {
					if roles.Contains(role) {
						break outer
					}
				}
				return validate.NewRequestError(errors.New("Forbidden"), http.StatusForbidden)
			case RoleModeAnd:
				for _, req := range required {
					if !roles.Contains(req) {
						return validate.NewRequestError(errors.New("Unauthorized"), http.StatusForbidden)
					}
				}
			}

			return next.ServeHTTP(w, r)
		})
	}
}

// mwAuthToken is a middleware that will check the database for a stateful token
// and attach it's user to the request context, or return an appropriate error.
// Authorization support is by token via Headers or Query Parameter
//
// Example:
//   - header = "Bearer 1234567890"
//   - query = "?access_token=1234567890"
//   - cookie = hb.auth.token = 1234567890
func (a *app) mwAuthToken(next errchain.Handler) errchain.Handler {
	return errchain.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		extractor := lookup.MultiExtractor{
			lookup.HeaderExtractor{Key: "Authorization", Prefix: "Bearer"},
			lookup.ArgumentExtractor("access_token"),
			lookup.CookieExtractor("hb.auth.token"),
		}
		requestToken, err := extractor.ExtractValue(r)
		if err != nil || requestToken == "" {
			return validate.NewRequestError(errors.New("Authorization header or query is required"), http.StatusUnauthorized)
		}

		r = r.WithContext(context.WithValue(r.Context(), hashedToken, requestToken))

		usr, err := a.services.User.GetSelf(r.Context(), requestToken)
		// Check the database for the token
		if err != nil {
			return validate.NewRequestError(errors.New("valid authorization header is required"), http.StatusUnauthorized)
		}

		r = r.WithContext(services.SetUserCtx(r.Context(), &usr, requestToken))
		return next.ServeHTTP(w, r)
	})
}
