package v1

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/internal/services"
	"github.com/hay-kot/homebox/backend/pkgs/server"
	"github.com/rs/zerolog/log"
)

// HandleUserSelf godoc
// @Summary  Get the current user
// @Tags     User
// @Produce  json
// @Param    payload  body  services.UserRegistration  true  "User Data"
// @Success  204
// @Router   /v1/users/register [Post]
func (ctrl *V1Controller) HandleUserRegistration() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		regData := services.UserRegistration{}

		if err := server.Decode(r, &regData); err != nil {
			log.Err(err).Msg("failed to decode user registration data")
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		_, err := ctrl.svc.User.RegisterUser(r.Context(), regData)
		if err != nil {
			log.Err(err).Msg("failed to register user")
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		server.Respond(w, http.StatusNoContent, nil)
	}
}

// HandleUserSelf godoc
// @Summary   Get the current user
// @Tags      User
// @Produce   json
// @Success   200  {object}  server.Result{item=repo.UserOut}
// @Router    /v1/users/self [GET]
// @Security  Bearer
func (ctrl *V1Controller) HandleUserSelf() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := services.UseTokenCtx(r.Context())
		usr, err := ctrl.svc.User.GetSelf(r.Context(), token)
		if usr.ID == uuid.Nil || err != nil {
			log.Err(err).Msg("failed to get user")
			server.RespondServerError(w)
			return
		}

		server.Respond(w, http.StatusOK, server.Wrap(usr))
	}
}

// HandleUserSelfUpdate godoc
// @Summary   Update the current user
// @Tags      User
// @Produce   json
// @Param     payload  body      repo.UserUpdate  true  "User Data"
// @Success   200      {object}  server.Result{item=repo.UserUpdate}
// @Router    /v1/users/self [PUT]
// @Security  Bearer
func (ctrl *V1Controller) HandleUserSelfUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updateData := repo.UserUpdate{}
		if err := server.Decode(r, &updateData); err != nil {
			log.Err(err).Msg("failed to decode user update data")
			server.RespondError(w, http.StatusBadRequest, err)
			return
		}

		actor := services.UseUserCtx(r.Context())
		newData, err := ctrl.svc.User.UpdateSelf(r.Context(), actor.ID, updateData)

		if err != nil {

			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		server.Respond(w, http.StatusOK, server.Wrap(newData))
	}
}

// HandleUserUpdatePassword godoc
// @Summary   Update the current user's password // TODO:
// @Tags      User
// @Produce   json
// @Success   204
// @Router    /v1/users/self/password [PUT]
// @Security  Bearer
func (ctrl *V1Controller) HandleUserUpdatePassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// HandleUserSelfDelete godoc
// @Summary   Deletes the user account
// @Tags      User
// @Produce   json
// @Success   204
// @Router    /v1/users/self [DELETE]
// @Security  Bearer
func (ctrl *V1Controller) HandleUserSelfDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		actor := services.UseUserCtx(r.Context())
		if err := ctrl.svc.User.DeleteSelf(r.Context(), actor.ID); err != nil {
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		server.Respond(w, http.StatusNoContent, nil)
	}
}

type (
	ChangePassword struct {
		Current string `json:"current,omitempty"`
		New     string `json:"new,omitempty"`
	}
)

// HandleUserSelfChangePassword godoc
// @Summary   Updates the users password
// @Tags      User
// @Success   204
// @Param     payload  body  ChangePassword  true  "Password Payload"
// @Router    /v1/users/change-password [PUT]
// @Security  Bearer
func (ctrl *V1Controller) HandleUserSelfChangePassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if ctrl.disablePasswordChange {
			server.RespondError(w, http.StatusForbidden, nil)
			return
		}

		var cp ChangePassword
		err := server.Decode(r, &cp)
		if err != nil {
			log.Err(err).Msg("user failed to change password")
		}

		ctx := services.NewContext(r.Context())

		ok := ctrl.svc.User.ChangePassword(ctx, cp.Current, cp.New)
		if !ok {
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		server.Respond(w, http.StatusNoContent, nil)
	}
}
