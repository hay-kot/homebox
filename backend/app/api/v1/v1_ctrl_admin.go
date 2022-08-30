package v1

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/hay-kot/git-web-template/backend/internal/services"
	"github.com/hay-kot/git-web-template/backend/internal/types"
	"github.com/hay-kot/git-web-template/backend/pkgs/hasher"
	"github.com/hay-kot/git-web-template/backend/pkgs/logger"
	"github.com/hay-kot/git-web-template/backend/pkgs/server"
)

// HandleAdminUserGetAll godoc
// @Summary   Gets all users from the database
// @Tags      Admin: Users
// @Produce   json
// @Success   200  {object}  server.Result{item=[]types.UserOut}
// @Router    /v1/admin/users [get]
// @Security  Bearer
func (ctrl *V1Controller) HandleAdminUserGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := ctrl.svc.Admin.GetAll(r.Context())

		if err != nil {
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		server.Respond(w, http.StatusOK, server.Wrap(users))
	}
}

// HandleAdminUserGet godoc
// @Summary   Get a user from the database
// @Tags      Admin: Users
// @Produce   json
// @Param     id  path  string  true  "User ID"
// @Success   200      {object}  server.Result{item=types.UserOut}
// @Router    /v1/admin/users/{id} [get]
// @Security  Bearer
func (ctrl *V1Controller) HandleAdminUserGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, err := uuid.Parse(chi.URLParam(r, "id"))

		if err != nil {
			ctrl.log.Debug(err.Error(), logger.Props{
				"scope":   "admin",
				"details": "failed to convert id to valid UUID",
			})
			server.RespondError(w, http.StatusBadRequest, err)
			return
		}

		user, err := ctrl.svc.Admin.GetByID(r.Context(), uid)

		if err != nil {
			ctrl.log.Error(err, nil)
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}
		server.Respond(w, http.StatusOK, server.Wrap(user))

	}
}

// HandleAdminUserCreate godoc
// @Summary   Create a new user
// @Tags      Admin: Users
// @Produce   json
// @Param     payload  body      types.UserCreate  true  "User Data"
// @Success   200  {object}  server.Result{item=types.UserOut}
// @Router    /v1/admin/users [POST]
// @Security  Bearer
func (ctrl *V1Controller) HandleAdminUserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createData := types.UserCreate{}

		if err := server.Decode(r, &createData); err != nil {
			ctrl.log.Error(err, logger.Props{
				"scope":   "admin",
				"details": "failed to decode user create data",
			})
			server.RespondError(w, http.StatusBadRequest, err)
			return
		}

		err := createData.Validate()

		if err != nil {
			server.RespondError(w, http.StatusUnprocessableEntity, err)
			return
		}

		hashedPw, err := hasher.HashPassword(createData.Password)

		if err != nil {
			ctrl.log.Error(err, logger.Props{
				"scope":   "admin",
				"details": "failed to hash password",
			})

			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		createData.Password = hashedPw
		userOut, err := ctrl.svc.Admin.Create(r.Context(), createData)

		if err != nil {
			ctrl.log.Error(err, logger.Props{
				"scope":   "admin",
				"details": "failed to create user",
			})

			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		server.Respond(w, http.StatusCreated, server.Wrap(userOut))
	}
}

// HandleAdminUserUpdate godoc
// @Summary   Update a User
// @Tags      Admin: Users
// @Param     id       path  string            true  "User ID"
// @Param     payload  body  types.UserUpdate  true  "User Data"
// @Produce   json
// @Success   200  {object}  server.Result{item=types.UserOut}
// @Router    /v1/admin/users/{id} [PUT]
// @Security  Bearer
func (ctrl *V1Controller) HandleAdminUserUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			ctrl.log.Debug(err.Error(), logger.Props{
				"scope":   "admin",
				"details": "failed to convert id to valid UUID",
			})
		}

		updateData := types.UserUpdate{}

		if err := server.Decode(r, &updateData); err != nil {
			ctrl.log.Error(err, logger.Props{
				"scope":   "admin",
				"details": "failed to decode user update data",
			})
			server.RespondError(w, http.StatusBadRequest, err)
			return
		}

		newData, err := ctrl.svc.Admin.UpdateProperties(r.Context(), uid, updateData)

		if err != nil {
			ctrl.log.Error(err, logger.Props{
				"scope":   "admin",
				"details": "failed to update user",
			})
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		server.Respond(w, http.StatusOK, server.Wrap(newData))
	}
}

// HandleAdminUserDelete godoc
// @Summary   Delete a User
// @Tags      Admin: Users
// @Param     id   path      string  true  "User ID"
// @Produce   json
// @Success   204
// @Router    /v1/admin/users/{id} [DELETE]
// @Security  Bearer
func (ctrl *V1Controller) HandleAdminUserDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			ctrl.log.Debug(err.Error(), logger.Props{
				"scope":   "admin",
				"details": "failed to convert id to valid UUID",
			})
		}

		actor := services.UseUserCtx(r.Context())

		if actor.ID == uid {
			server.RespondError(w, http.StatusBadRequest, errors.New("cannot delete yourself"))
			return
		}

		err = ctrl.svc.Admin.Delete(r.Context(), uid)

		if err != nil {
			ctrl.log.Error(err, logger.Props{
				"scope":   "admin",
				"details": "failed to delete user",
			})
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}
	}
}
