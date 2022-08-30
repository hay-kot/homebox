package v1

import (
	"errors"
	"net/http"

	"github.com/hay-kot/git-web-template/backend/internal/services"
	"github.com/hay-kot/git-web-template/backend/internal/types"
	"github.com/hay-kot/git-web-template/backend/pkgs/logger"
	"github.com/hay-kot/git-web-template/backend/pkgs/server"
)

var (
	HeaderFormData = "application/x-www-form-urlencoded"
	HeaderJSON     = "application/json"
)

// HandleAuthLogin godoc
// @Summary  User Login
// @Tags     Authentication
// @Accept   x-www-form-urlencoded
// @Accept   application/json
// @Param    username  formData  string  false  "string"  example(admin@admin.com)
// @Param    password  formData  string  false  "string"  example(admin)
// @Produce  json
// @Success  200  {object}  types.TokenResponse
// @Router   /v1/users/login [POST]
func (ctrl *V1Controller) HandleAuthLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginForm := &types.LoginForm{}

		if r.Header.Get("Content-Type") == HeaderFormData {
			err := r.ParseForm()

			if err != nil {
				server.Respond(w, http.StatusBadRequest, server.Wrap(err))
				return
			}

			loginForm.Username = r.PostFormValue("username")
			loginForm.Password = r.PostFormValue("password")
		} else if r.Header.Get("Content-Type") == HeaderJSON {
			err := server.Decode(r, loginForm)

			if err != nil {
				server.Respond(w, http.StatusBadRequest, server.Wrap(err))
				return
			}
		} else {
			server.Respond(w, http.StatusBadRequest, errors.New("invalid content type"))
			return
		}

		if loginForm.Username == "" || loginForm.Password == "" {
			server.RespondError(w, http.StatusBadRequest, errors.New("username and password are required"))
			return
		}

		newToken, err := ctrl.svc.User.Login(r.Context(), loginForm.Username, loginForm.Password)

		if err != nil {
			server.RespondError(w, http.StatusUnauthorized, err)
			return
		}

		err = server.Respond(w, http.StatusOK, types.TokenResponse{
			BearerToken: "Bearer " + newToken.Raw,
			ExpiresAt:   newToken.ExpiresAt,
		})

		if err != nil {
			ctrl.log.Error(err, logger.Props{
				"user": loginForm.Username,
			})
			return
		}
	}
}

// HandleAuthLogout godoc
// @Summary   User Logout
// @Tags      Authentication
// @Success   204
// @Router    /v1/users/logout [POST]
// @Security  Bearer
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

		err = server.Respond(w, http.StatusNoContent, nil)
	}
}

// HandleAuthLogout godoc
// @Summary      User Token Refresh
// @Description  handleAuthRefresh returns a handler that will issue a new token from an existing token.
// @Description  This does not validate that the user still exists within the database.
// @Tags         Authentication
// @Success      200
// @Router       /v1/users/refresh [GET]
// @Security     Bearer
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

		err = server.Respond(w, http.StatusOK, newToken)

		if err != nil {
			return
		}
	}
}
