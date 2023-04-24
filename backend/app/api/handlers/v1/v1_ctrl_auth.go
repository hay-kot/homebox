package v1

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/helper"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/safeserve/errchain"
	"github.com/hay-kot/safeserve/server"
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

func (ctrl *V1Controller) HandleSsoHeaderLogin() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		log.Info().Msg("Header SSO Login Attempt")
		if !ctrl.headerSSOEnabled {
			return validate.NewRequestError(errors.New("authentication failed. Header SSO is disaled"), http.StatusInternalServerError)			
		}

		{
			t := strings.Split(r.RemoteAddr, ":")
			if t[0] != ctrl.headerSSOAllowedIP {
				return validate.NewRequestError(errors.New("authentication failed. Header SSO not allowed for this remote IP"), http.StatusInternalServerError)			
			}
			log.Info().Msgf("Header SSO Login Attempt allowed from IP '%s'", t[0])
		}

		email := r.Header.Get("Remote-Email")

		if email == "" {
			return validate.NewRequestError(errors.New("authentication failed. not SSO header found or empty"), http.StatusInternalServerError)
		}

		// check if a user matching provided email is existing already
		_, err := ctrl.repo.Users.GetOneEmail(r.Context(), email)

		if err != nil {
			// user not found -> create it
			var username = r.Header.Get("Remote-Name")
			
			/* TODO: decide how to handle group information provided by HTTP header
			// if groups are provided, they will be comma-separated. take only the first group
			var groups = r.Header.Get("Remote-Groups")
			var groupArr = strings.Split(groups, ",")
			groupTok := ""
			if len(groupArr) > 0 {
				groupTok = groupArr[0]
			}
			*/
			
			// Use a randomly generatd password. Not meant to be used as login. Only a dummy.
			regData := services.UserRegistration {
				GroupToken: "",  // don't set group for now
				Name : username,
				Email : email,
				Password : helper.GenerateRandomPassword(64, 12, 5, 5),
			}

			_, err := ctrl.svc.User.RegisterUser(r.Context(), regData)
			if err != nil {
				log.Err(err).Msg("failed to register user from SSO HTTP headers")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}
		}

		// login as user with provided password
		newToken, err := ctrl.svc.User.LoginWithoutPassword(r.Context(), strings.ToLower(email), false)

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
