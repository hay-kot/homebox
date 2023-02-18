package v1

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/homebox/backend/pkgs/server"
	"github.com/rs/zerolog/log"
)

// HandleUserSelf godoc
// @Summary Get the current user
// @Tags    User
// @Produce json
// @Param   payload body services.UserRegistration true "User Data"
// @Success 204
// @Router  /v1/users/register [Post]
func (ctrl *V1Controller) HandleUserRegistration() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		regData := services.UserRegistration{}

		if err := server.Decode(r, &regData); err != nil {
			log.Err(err).Msg("failed to decode user registration data")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		if !ctrl.allowRegistration && regData.GroupToken == "" {
			return validate.NewRequestError(fmt.Errorf("user registration disabled"), http.StatusForbidden)
		}

		_, err := ctrl.svc.User.RegisterUser(r.Context(), regData)
		if err != nil {
			log.Err(err).Msg("failed to register user")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusNoContent, nil)
	}
}

// HandleUserSelf godoc
// @Summary  Get the current user
// @Tags     User
// @Produce  json
// @Success  200 {object} server.Result{item=repo.UserOut}
// @Router   /v1/users/self [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleUserSelf() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		token := services.UseTokenCtx(r.Context())
		usr, err := ctrl.svc.User.GetSelf(r.Context(), token)
		if usr.ID == uuid.Nil || err != nil {
			log.Err(err).Msg("failed to get user")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusOK, server.Wrap(usr))
	}
}

// HandleUserSelfUpdate godoc
// @Summary  Update the current user
// @Tags     User
// @Produce  json
// @Param    payload body     repo.UserUpdate true "User Data"
// @Success  200     {object} server.Result{item=repo.UserUpdate}
// @Router   /v1/users/self [PUT]
// @Security Bearer
func (ctrl *V1Controller) HandleUserSelfUpdate() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		updateData := repo.UserUpdate{}
		if err := server.Decode(r, &updateData); err != nil {
			log.Err(err).Msg("failed to decode user update data")
			return validate.NewRequestError(err, http.StatusBadRequest)
		}

		actor := services.UseUserCtx(r.Context())
		newData, err := ctrl.svc.User.UpdateSelf(r.Context(), actor.ID, updateData)
		if err != nil {
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusOK, server.Wrap(newData))
	}
}

// HandleUserSelfDelete godoc
// @Summary  Deletes the user account
// @Tags     User
// @Produce  json
// @Success  204
// @Router   /v1/users/self [DELETE]
// @Security Bearer
func (ctrl *V1Controller) HandleUserSelfDelete() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if ctrl.isDemo {
			return validate.NewRequestError(nil, http.StatusForbidden)
		}

		actor := services.UseUserCtx(r.Context())
		if err := ctrl.svc.User.DeleteSelf(r.Context(), actor.ID); err != nil {
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusNoContent, nil)
	}
}

type (
	ChangePassword struct {
		Current string `json:"current,omitempty"`
		New     string `json:"new,omitempty"`
	}
)

// HandleUserSelfChangePassword godoc
// @Summary  Updates the users password
// @Tags     User
// @Success  204
// @Param    payload body ChangePassword true "Password Payload"
// @Router   /v1/users/change-password [PUT]
// @Security Bearer
func (ctrl *V1Controller) HandleUserSelfChangePassword() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if ctrl.isDemo {
			return validate.NewRequestError(nil, http.StatusForbidden)
		}

		var cp ChangePassword
		err := server.Decode(r, &cp)
		if err != nil {
			log.Err(err).Msg("user failed to change password")
		}

		ctx := services.NewContext(r.Context())

		ok := ctrl.svc.User.ChangePassword(ctx, cp.Current, cp.New)
		if !ok {
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusNoContent, nil)
	}
}
