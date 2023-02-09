package v1

import (
	"net/http"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/homebox/backend/pkgs/server"
	"github.com/rs/zerolog/log"
)

type ActionAmountResult struct {
	Completed int `json:"completed"`
}

// HandleGroupInvitationsCreate godoc
// @Summary  Ensures all items in the database have an asset id
// @Tags     Group
// @Produce  json
// @Success  200     {object} ActionAmountResult
// @Router   /v1/actions/ensure-asset-ids [Post]
// @Security Bearer
func (ctrl *V1Controller) HandleEnsureAssetID() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())

		totalCompleted, err := ctrl.svc.Items.EnsureAssetID(ctx, ctx.GID)
		if err != nil {
			log.Err(err).Msg("failed to ensure asset id")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusOK, ActionAmountResult{Completed: totalCompleted})
	}
}

// HandleItemDateZeroOut godoc
// @Summary  Resets all item date fields to the beginning of the day
// @Tags     Group
// @Produce  json
// @Success  200     {object} ActionAmountResult
// @Router   /v1/actions/zero-item-time-fields [Post]
// @Security Bearer
func (ctrl *V1Controller) HandleItemDateZeroOut() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())

		totalCompleted, err := ctrl.repo.Items.ZeroOutTimeFields(ctx, ctx.GID)
		if err != nil {
			log.Err(err).Msg("failed to ensure asset id")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusOK, ActionAmountResult{Completed: totalCompleted})
	}
}
