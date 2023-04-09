package v1

import (
	"net/http"

	"github.com/containrrr/shoutrrr"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/homebox/backend/internal/web/adapters"
	"github.com/hay-kot/httpkit/errchain"
)

// HandleGetUserNotifiers godoc
//
//	@Summary  Get Notifiers
//	@Tags     Notifiers
//	@Produce  json
//	@Success  200 {object} []repo.NotifierOut
//	@Router   /v1/notifiers [GET]
//	@Security Bearer
func (ctrl *V1Controller) HandleGetUserNotifiers() errchain.HandlerFunc {
	fn := func(r *http.Request, _ struct{}) ([]repo.NotifierOut, error) {
		user := services.UseUserCtx(r.Context())
		return ctrl.repo.Notifiers.GetByUser(r.Context(), user.ID)
	}

	return adapters.Query(fn, http.StatusOK)
}

// HandleCreateNotifier godoc
//
//	@Summary  Create Notifier
//	@Tags     Notifiers
//	@Produce  json
//	@Param    payload body     repo.NotifierCreate true "Notifier Data"
//	@Success  200     {object} repo.NotifierOut
//	@Router   /v1/notifiers [POST]
//	@Security Bearer
func (ctrl *V1Controller) HandleCreateNotifier() errchain.HandlerFunc {
	fn := func(r *http.Request, in repo.NotifierCreate) (repo.NotifierOut, error) {
		auth := services.NewContext(r.Context())
		return ctrl.repo.Notifiers.Create(auth, auth.GID, auth.UID, in)
	}

	return adapters.Action(fn, http.StatusCreated)
}

// HandleDeleteNotifier godocs
//
//	@Summary Delete a Notifier
//	@Tags    Notifiers
//	@Param   id path string true "Notifier ID"
//	@Success 204
//	@Router  /v1/notifiers/{id} [DELETE]
//	@Security Bearer
func (ctrl *V1Controller) HandleDeleteNotifier() errchain.HandlerFunc {
	fn := func(r *http.Request, ID uuid.UUID) (any, error) {
		auth := services.NewContext(r.Context())
		return nil, ctrl.repo.Notifiers.Delete(auth, auth.UID, ID)
	}

	return adapters.CommandID("id", fn, http.StatusNoContent)
}

// HandleUpdateNotifier godocs
//
//	@Summary Update Notifier
//	@Tags    Notifiers
//	@Param   id path string true "Notifier ID"
//	@Param   payload body repo.NotifierUpdate true "Notifier Data"
//	@Success 200 {object} repo.NotifierOut
//	@Router  /v1/notifiers/{id} [PUT]
//	@Security Bearer
func (ctrl *V1Controller) HandleUpdateNotifier() errchain.HandlerFunc {
	fn := func(r *http.Request, ID uuid.UUID, in repo.NotifierUpdate) (repo.NotifierOut, error) {
		auth := services.NewContext(r.Context())
		return ctrl.repo.Notifiers.Update(auth, auth.UID, ID, in)
	}

	return adapters.ActionID("id", fn, http.StatusOK)
}

// HandlerNotifierTest godoc
//
//	@Summary  Test Notifier
//	@Tags     Notifiers
//	@Produce  json
//	@Param    id path string true "Notifier ID"
//	@Param url query string true "URL"
//	@Success  204
//	@Router   /v1/notifiers/test [POST]
//	@Security Bearer
func (ctrl *V1Controller) HandlerNotifierTest() errchain.HandlerFunc {
	type body struct {
		URL string `json:"url" validate:"required"`
	}

	fn := func(r *http.Request, q body) (any, error) {
		err := shoutrrr.Send(q.URL, "Test message from Homebox")
		return nil, err
	}

	return adapters.Action(fn, http.StatusOK)
}
