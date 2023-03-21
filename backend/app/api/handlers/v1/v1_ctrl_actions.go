package v1

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/safeserve/errchain"
	"github.com/hay-kot/safeserve/server"
	"github.com/rs/zerolog/log"
)

type ActionAmountResult struct {
	Completed int `json:"completed"`
}

func actionHandlerFactory(ref string, fn func(context.Context, uuid.UUID) (int, error)) errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())

		totalCompleted, err := fn(ctx, ctx.GID)
		if err != nil {
			log.Err(err).Str("action_ref", ref).Msg("failed to run action")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.JSON(w, http.StatusOK, ActionAmountResult{Completed: totalCompleted})
	}
}

// HandleEnsureAssetID godoc
//
//	@Summary     Ensure Asset IDs
//	@Description Ensures all items in the database have an asset ID
//	@Tags        Actions
//	@Produce     json
//	@Success     200     {object} ActionAmountResult
//	@Router      /v1/actions/ensure-asset-ids [Post]
//	@Security    Bearer
func (ctrl *V1Controller) HandleEnsureAssetID() errchain.HandlerFunc {
	return actionHandlerFactory("ensure asset IDs", ctrl.svc.Items.EnsureAssetID)
}

// HandleEnsureImportRefs godoc
//
//	@Summary  Ensures Import Refs
//	@Description  Ensures all items in the database have an import ref
//	@Tags     Actions
//	@Produce  json
//	@Success  200     {object} ActionAmountResult
//	@Router   /v1/actions/ensure-import-refs [Post]
//	@Security Bearer
func (ctrl *V1Controller) HandleEnsureImportRefs() errchain.HandlerFunc {
	return actionHandlerFactory("ensure import refs", ctrl.svc.Items.EnsureImportRef)
}

// HandleItemDateZeroOut godoc
//
//	@Summary      Zero Out Time Fields
//	@Description  Resets all item date fields to the beginning of the day
//	@Tags         Actions
//	@Produce      json
//	@Success      200     {object} ActionAmountResult
//	@Router       /v1/actions/zero-item-time-fields [Post]
//	@Security Bearer
func (ctrl *V1Controller) HandleItemDateZeroOut() errchain.HandlerFunc {
	return actionHandlerFactory("zero out date time", ctrl.repo.Items.ZeroOutTimeFields)
}
