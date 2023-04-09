package v1

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/httpkit/errchain"
	"github.com/hay-kot/httpkit/server"
	"github.com/rs/zerolog/log"
)

type (
	TokenResponse struct {
		Token           string    `json:"token"`
		ExpiresAt       time.Time `json:"expiresAt"`
		AttachmentToken string    `json:"attachmentToken"`
	}

	LoginForm struct {
		Username     string `json:"username"`
		Password     string `json:"password"`
		StayLoggedIn bool   `json:"stayLoggedIn"`
	}
)

// HandleAuthLogin godoc
//
//	@Summary User Login
//	@Tags    Authentication
//	@Accept  x-www-form-urlencoded
//	@Accept  application/json
//	@Param   username formData string false "string" example(admin@admin.com)
//	@Param   password formData string false "string" example(admin)
//	@Param    payload body     LoginForm true "Login Data"
//	@Produce json
//	@Success 200 {object} TokenResponse
//	@Router  /v1/users/login [POST]
func (ctrl *V1Controller) HandleAuthLogin() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		loginForm := &LoginForm{}

		switch r.Header.Get("Content-Type") {
		case "application/x-www-form-urlencoded":
			err := r.ParseForm()
			if err != nil {
				return errors.New("failed to parse form")
			}

			loginForm.Username = r.PostFormValue("username")
			loginForm.Password = r.PostFormValue("password")
			loginForm.StayLoggedIn = r.PostFormValue("stayLoggedIn") == "true"
		case "application/json":
			err := server.Decode(r, loginForm)
			if err != nil {
				log.Err(err).Msg("failed to decode login form")
				return errors.New("failed to decode login form")
			}
		default:
			return server.JSON(w, http.StatusBadRequest, errors.New("invalid content type"))
		}

		if loginForm.Username == "" || loginForm.Password == "" {
			return validate.NewFieldErrors(
				validate.FieldError{
					Field: "username",
					Error: "username or password is empty",
				},
				validate.FieldError{
					Field: "password",
					Error: "username or password is empty",
				},
			)
		}

		newToken, err := ctrl.svc.User.Login(r.Context(), strings.ToLower(loginForm.Username), loginForm.Password, loginForm.StayLoggedIn)
		if err != nil {
			return validate.NewRequestError(errors.New("authentication failed"), http.StatusInternalServerError)
		}

		return server.JSON(w, http.StatusOK, TokenResponse{
			Token:           "Bearer " + newToken.Raw,
			ExpiresAt:       newToken.ExpiresAt,
			AttachmentToken: newToken.AttachmentToken,
		})
	}
}

// HandleAuthLogout godoc
//
//	@Summary  User Logout
//	@Tags     Authentication
//	@Success  204
//	@Router   /v1/users/logout [POST]
//	@Security Bearer
func (ctrl *V1Controller) HandleAuthLogout() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		token := services.UseTokenCtx(r.Context())
		if token == "" {
			return validate.NewRequestError(errors.New("no token within request context"), http.StatusUnauthorized)
		}

		err := ctrl.svc.User.Logout(r.Context(), token)
		if err != nil {
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.JSON(w, http.StatusNoContent, nil)
	}
}

// HandleAuthLogout godoc
//
//	@Summary     User Token Refresh
//	@Description handleAuthRefresh returns a handler that will issue a new token from an existing token.
//	@Description This does not validate that the user still exists within the database.
//	@Tags        Authentication
//	@Success     200
//	@Router      /v1/users/refresh [GET]
//	@Security    Bearer
func (ctrl *V1Controller) HandleAuthRefresh() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		requestToken := services.UseTokenCtx(r.Context())
		if requestToken == "" {
			return validate.NewRequestError(errors.New("no token within request context"), http.StatusUnauthorized)
		}

		newToken, err := ctrl.svc.User.RenewToken(r.Context(), requestToken)
		if err != nil {
			return validate.NewUnauthorizedError()
		}

		return server.JSON(w, http.StatusOK, newToken)
	}
}
