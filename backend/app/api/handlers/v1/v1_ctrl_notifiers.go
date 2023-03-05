package v1

import (
	"context"
	"net/http"

	"github.com/containrrr/shoutrrr"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/homebox/backend/internal/web/adapters"
	"github.com/hay-kot/homebox/backend/pkgs/server"
)

// HandleGetUserNotifiers godoc
// @Summary  Get All notifier
// @Tags     Notifiers
// @Produce  json
// @Success  200 {object} server.Results{items=[]repo.NotifierOut}
// @Router   /v1/notifiers [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleGetUserNotifiers() server.HandlerFunc {
	fn := func(ctx context.Context, _ struct{}) ([]repo.NotifierOut, error) {
		user := services.UseUserCtx(ctx)
		return ctrl.repo.Notifiers.GetByUser(ctx, user.ID)
	}

	return adapters.Query(fn, http.StatusOK)
}

// HandleCreateNotifier godoc
// @Summary  Create a new notifier
// @Tags     Notifiers
// @Produce  json
// @Param    payload body     repo.NotifierCreate true "Notifier Data"
// @Success  200     {object} repo.NotifierOut
// @Router   /v1/notifiers [POST]
// @Security Bearer
func (ctrl *V1Controller) HandleCreateNotifier() server.HandlerFunc {
	fn := func(ctx context.Context, in repo.NotifierCreate) (repo.NotifierOut, error) {
		auth := services.NewContext(ctx)
		return ctrl.repo.Notifiers.Create(ctx, auth.GID, auth.UID, in)
	}

	return adapters.Action(fn, http.StatusCreated)
}

// HandleDeleteNotifier godocs
// @Summary Delete a notifier
// @Tags    Notifiers
// @Param   id path string true "Notifier ID"
// @Success 204
// @Router  /v1/notifiers/{id} [DELETE]
// @Security Bearer
func (ctrl *V1Controller) HandleDeleteNotifier() server.HandlerFunc {
	fn := func(ctx context.Context, ID uuid.UUID) (any, error) {
		auth := services.NewContext(ctx)
		return nil, ctrl.repo.Notifiers.Delete(ctx, auth.UID, ID)
	}

	return adapters.CommandID("id", fn, http.StatusNoContent)
}

// HandleUpdateNotifier godocs
// @Summary Update a notifier
// @Tags    Notifiers
// @Param   id path string true "Notifier ID"
// @Param   payload body repo.NotifierUpdate true "Notifier Data"
// @Success 200 {object} repo.NotifierOut
// @Router  /v1/notifiers/{id} [PUT]
// @Security Bearer
func (ctrl *V1Controller) HandleUpdateNotifier() server.HandlerFunc {
	fn := func(ctx context.Context, ID uuid.UUID, in repo.NotifierUpdate) (repo.NotifierOut, error) {
		auth := services.NewContext(ctx)
		return ctrl.repo.Notifiers.Update(ctx, auth.UID, ID, in)
	}

	return adapters.ActionID("id", fn, http.StatusOK)
}

// HandlerNotifierTest godoc
// @Summary  Test notifier
// @Tags     Notifiers
// @Produce  json
// @Param    id path string true "Notifier ID"
// @Param url query string true "URL"
// @Success  204
// @Router   /v1/notifiers/test [POST]
// @Security Bearer
func (ctrl *V1Controller) HandlerNotifierTest() server.HandlerFunc {
	type body struct {
		URL string `json:"url" validate:"required"`
	}

	fn := func(ctx context.Context, q body) (any, error) {
		err := shoutrrr.Send(q.URL, "Test message from Homebox")
		return nil, err
	}

	return adapters.Action(fn, http.StatusOK)
}
