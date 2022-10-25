package v1

import (
	"errors"
	"net/http"
	"time"

	"github.com/hay-kot/homebox/backend/internal/services"
	"github.com/hay-kot/homebox/backend/pkgs/server"
	"github.com/rs/zerolog/log"
)

type (
	TokenResponse struct {
		Token     string    `json:"token"`
		ExpiresAt time.Time `json:"expiresAt"`
	}

	LoginForm struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

// HandleAuthLogin godoc
// @Summary User Login
// @Tags    Authentication
// @Accept  x-www-form-urlencoded
// @Accept  application/json
// @Param   username formData string false "string" example(admin@admin.com)
// @Param   password formData string false "string" example(admin)
// @Produce json
// @Success 200 {object} TokenResponse
// @Router  /v1/users/login [POST]
func (ctrl *V1Controller) HandleAuthLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginForm := &LoginForm{}

		switch r.Header.Get("Content-Type") {
		case server.ContentFormUrlEncoded:
			err := r.ParseForm()
			if err != nil {
				server.Respond(w, http.StatusBadRequest, server.Wrap(err))
				log.Error().Err(err).Msg("failed to parse form")
				return
			}

			loginForm.Username = r.PostFormValue("username")
			loginForm.Password = r.PostFormValue("password")
		case server.ContentJSON:
			err := server.Decode(r, loginForm)

			if err != nil {
				log.Err(err).Msg("failed to decode login form")
				server.Respond(w, http.StatusBadRequest, server.Wrap(err))
				return
			}
		default:
			server.Respond(w, http.StatusBadRequest, errors.New("invalid content type"))
			return
		}

		if loginForm.Username == "" || loginForm.Password == "" {
			server.RespondError(w, http.StatusBadRequest, errors.New("username and password are required"))
			return
		}

		newToken, err := ctrl.svc.User.Login(r.Context(), loginForm.Username, loginForm.Password)

		if err != nil {
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		server.Respond(w, http.StatusOK, TokenResponse{
			Token:     "Bearer " + newToken.Raw,
			ExpiresAt: newToken.ExpiresAt,
		})
	}
}

// HandleAuthLogout godoc
// @Summary  User Logout
// @Tags     Authentication
// @Success  204
// @Router   /v1/users/logout [POST]
// @Security Bearer
func (ctrl *V1Controller) HandleAuthLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := services.UseTokenCtx(r.Context())

		if token == "" {
			server.RespondError(w, http.StatusUnauthorized, errors.New("no token within request context"))
			return
		}

		err := ctrl.svc.User.Logout(r.Context(), token)

		if err != nil {
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		server.Respond(w, http.StatusNoContent, nil)
	}
}

// HandleAuthLogout godoc
// @Summary     User Token Refresh
// @Description handleAuthRefresh returns a handler that will issue a new token from an existing token.
// @Description This does not validate that the user still exists within the database.
// @Tags        Authentication
// @Success     200
// @Router      /v1/users/refresh [GET]
// @Security    Bearer
func (ctrl *V1Controller) HandleAuthRefresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestToken := services.UseTokenCtx(r.Context())

		if requestToken == "" {
			server.RespondError(w, http.StatusUnauthorized, errors.New("no user token found"))
			return
		}

		newToken, err := ctrl.svc.User.RenewToken(r.Context(), requestToken)

		if err != nil {
			server.RespondUnauthorized(w)
			return
		}

		server.Respond(w, http.StatusOK, newToken)
	}
}
