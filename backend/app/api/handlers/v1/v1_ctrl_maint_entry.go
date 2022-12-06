package v1

import (
	"net/http"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/homebox/backend/pkgs/server"
	"github.com/rs/zerolog/log"
)

// HandleMaintenanceGetLog godoc
// @Summary  Get Maintenance Log
// @Tags     Maintenance
// @Produce  json
// @Success  200       {object} repo.MaintenanceLog
// @Router   /v1/items/{id}/maintenance [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleMaintenanceLogGet() server.HandlerFunc {
	return ctrl.handleMaintenanceLog()
}

// HandleMaintenanceEntryCreate godoc
// @Summary  Create Maintenance Entry
// @Tags     Maintenance
// @Produce  json
// @Param    payload body     repo.MaintenanceEntryCreate true "Entry Data"
// @Success  200     {object} repo.MaintenanceEntry
// @Router   /v1/items/{id}/maintenance [POST]
// @Security Bearer
func (ctrl *V1Controller) HandleMaintenanceEntryCreate() server.HandlerFunc {
	return ctrl.handleMaintenanceLog()
}

// HandleMaintenanceEntryDelete godoc
// @Summary  Delete Maintenance Entry
// @Tags     Maintenance
// @Produce  json
// @Success  204
// @Router   /v1/items/{id}/maintenance/{entry_id} [DELETE]
// @Security Bearer
func (ctrl *V1Controller) HandleMaintenanceEntryDelete() server.HandlerFunc {
	return ctrl.handleMaintenanceLog()
}

// HandleMaintenanceEntryUpdate godoc
// @Summary  Update Maintenance Entry
// @Tags     Maintenance
// @Produce  json
// @Param    payload body     repo.MaintenanceEntryUpdate true "Entry Data"
// @Success  200     {object} repo.MaintenanceEntry
// @Router   /v1/items/{id}/maintenance/{entry_id} [PUT]
// @Security Bearer
func (ctrl *V1Controller) HandleMaintenanceEntryUpdate() server.HandlerFunc {
	return ctrl.handleMaintenanceLog()
}

func (ctrl *V1Controller) handleMaintenanceLog() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())
		itemID, err := ctrl.routeID(r)
		if err != nil {
			return err
		}

		switch r.Method {
		case http.MethodGet:
			mlog, err := ctrl.repo.MaintEntry.GetLog(ctx, itemID)
			if err != nil {
				log.Err(err).Msg("failed to get items")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}
			return server.Respond(w, http.StatusOK, mlog)
		case http.MethodPost:
			var create repo.MaintenanceEntryCreate
			err := server.Decode(r, &create)
			if err != nil {
				return validate.NewRequestError(err, http.StatusBadRequest)
			}

			entry, err := ctrl.repo.MaintEntry.Create(ctx, itemID, create)
			if err != nil {
				log.Err(err).Msg("failed to create item")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}

			return server.Respond(w, http.StatusCreated, entry)
		case http.MethodPut:
			entryID, err := ctrl.routeUUID(r, "entry_id")
			if err != nil {
				return err
			}

			var update repo.MaintenanceEntryUpdate
			err = server.Decode(r, &update)
			if err != nil {
				return validate.NewRequestError(err, http.StatusBadRequest)
			}

			entry, err := ctrl.repo.MaintEntry.Update(ctx, entryID, update)
			if err != nil {
				log.Err(err).Msg("failed to update item")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}

			return server.Respond(w, http.StatusOK, entry)
		case http.MethodDelete:
			entryID, err := ctrl.routeUUID(r, "entry_id")
			if err != nil {
				return err
			}

			err = ctrl.repo.MaintEntry.Delete(ctx, entryID)
			if err != nil {
				log.Err(err).Msg("failed to delete item")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}

			return server.Respond(w, http.StatusNoContent, nil)
		}

		return nil
	}
}
