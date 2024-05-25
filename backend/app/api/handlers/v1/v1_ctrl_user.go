package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/homebox/backend/internal/web/adapters"
	"github.com/hay-kot/httpkit/errchain"
	"github.com/hay-kot/httpkit/server"
	"github.com/rs/zerolog/log"
)

// HandleUserRegistration godoc
//
//	@Summary Register New User
//	@Tags    User
//	@Produce json
//	@Param   payload body services.UserRegistration true "User Data"
//	@Success 204
//	@Router  /v1/users/register [Post]
func (ctrl *V1Controller) HandleUserRegistration() errchain.HandlerFunc {
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

		return server.JSON(w, http.StatusNoContent, nil)
	}
}

// HandleUserSelf godoc
//
//	@Summary  Get User Self
//	@Tags     User
//	@Produce  json
//	@Success  200 {object} Wrapped{item=repo.UserOut}
//	@Router   /v1/users/self [GET]
//	@Security Bearer
func (ctrl *V1Controller) HandleUserSelf() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		token := services.UseTokenCtx(r.Context())
		usr, err := ctrl.svc.User.GetSelf(r.Context(), token)
		if usr.ID == uuid.Nil || err != nil {
			log.Err(err).Msg("failed to get user")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.JSON(w, http.StatusOK, Wrap(usr))
	}
}

// HandleUserSelfUpdate godoc
//
//	@Summary  Update Account
//	@Tags     User
//	@Produce  json
//	@Param    payload body     repo.UserUpdate true "User Data"
//	@Success  200     {object} Wrapped{item=repo.UserUpdate}
//	@Router   /v1/users/self [PUT]
//	@Security Bearer
func (ctrl *V1Controller) HandleUserSelfUpdate() errchain.HandlerFunc {
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

		return server.JSON(w, http.StatusOK, Wrap(newData))
	}
}

// HandleUserSelfDelete godoc
//
//	@Summary  Delete Account
//	@Tags     User
//	@Produce  json
//	@Success  204
//	@Router   /v1/users/self [DELETE]
//	@Security Bearer
func (ctrl *V1Controller) HandleUserSelfDelete() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if ctrl.isDemo {
			return validate.NewRequestError(nil, http.StatusForbidden)
		}

		actor := services.UseUserCtx(r.Context())
		if err := ctrl.svc.User.DeleteSelf(r.Context(), actor.ID); err != nil {
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.JSON(w, http.StatusNoContent, nil)
	}
}

type ChangePassword struct {
	Current string `json:"current,omitempty"`
	New     string `json:"new,omitempty"`
}

// HandleUserSelfChangePassword godoc
//
//	@Summary  Change Password
//	@Tags     User
//	@Success  204
//	@Param    payload body ChangePassword true "Password Payload"
//	@Router   /v1/users/change-password [PUT]
//	@Security Bearer
func (ctrl *V1Controller) HandleUserSelfChangePassword() errchain.HandlerFunc {
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

		ok := ctrl.svc.User.PasswordChange(ctx, cp.Current, cp.New)
		if !ok {
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.JSON(w, http.StatusNoContent, nil)
	}
}

// HandleUserSelfChangePasswordWithToken godoc
//
//	@Summary  Change Password
//	@Tags     User
//	@Success  204
//	@Param    payload body ChangePassword true "Password Payload"
//	@Router   /v1/users/change-password-token [PUT]
func (ctrl *V1Controller) HandleUserSelfChangePasswordWithToken() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		tokenQueryParam := r.URL.Query().Get("token")
		if tokenQueryParam == "" {
			return validate.NewRequestError(fmt.Errorf("missing token query param"), http.StatusBadRequest)
		}

		if ctrl.isDemo {
			return validate.NewRequestError(nil, http.StatusForbidden)
		}

		var cp ChangePassword
		err := server.Decode(r, &cp)
		if err != nil {
			log.Err(err).Msg("user failed to change password")
		}

		ctx := services.NewContext(r.Context())

		ok := ctrl.svc.User.PasswordChange(ctx, cp.Current, cp.New)
		if !ok {
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.JSON(w, http.StatusNoContent, nil)
	}
}

// HandleUserRequestPasswordReset godoc
//
// @Summary Request Password Reset
// @Tags    User
// @Produce json
// @Param   payload body services.PasswordResetRequest true "User Data"
// @Success 204
// @Router  /v1/users/request-password-reset [Post]
func (ctrl *V1Controller) HandleUserRequestPasswordReset() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if ctrl.isDemo {
			return validate.NewRequestError(nil, http.StatusForbidden)
		}

		v, err := adapters.DecodeBody[services.PasswordResetRequest](r)
		if err != nil {
			return err
		}

		go func() {
			ctx := context.Background()

			err = ctrl.svc.User.PasswordResetRequest(ctx, v)
			if err != nil {
				log.Warn().
					Err(err).
					Str("email", v.Email).
					Msg("failed to request password reset")
			}
		}()

		return server.JSON(w, http.StatusNoContent, nil)
	}
}
