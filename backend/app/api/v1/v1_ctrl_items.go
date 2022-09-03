package v1

import (
	"net/http"

	"github.com/hay-kot/content/backend/internal/services"
	"github.com/hay-kot/content/backend/internal/types"
	"github.com/hay-kot/content/backend/pkgs/server"
	"github.com/rs/zerolog/log"
)

// HandleItemsGetAll godoc
// @Summary   Get All Items
// @Tags      Items
// @Produce   json
// @Success   200  {object}  server.Results{items=[]types.ItemOut}
// @Router    /v1/items [GET]
// @Security  Bearer
func (ctrl *V1Controller) HandleItemsGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := services.UseUserCtx(r.Context())
		items, err := ctrl.svc.Items.GetAll(r.Context(), user.GroupID)
		if err != nil {
			log.Err(err).Msg("failed to get items")
			server.RespondServerError(w)
			return
		}
		server.Respond(w, http.StatusOK, server.Results{Items: items})
	}
}

// HandleItemsCreate godoc
// @Summary   Create a new item
// @Tags      Items
// @Produce   json
// @Param     payload  body      types.ItemCreate  true  "Item Data"
// @Success   200      {object}  types.ItemSummary
// @Router    /v1/items [POST]
// @Security  Bearer
func (ctrl *V1Controller) HandleItemsCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createData := types.ItemCreate{}
		if err := server.Decode(r, &createData); err != nil {
			log.Err(err).Msg("failed to decode request body")
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		user := services.UseUserCtx(r.Context())
		item, err := ctrl.svc.Items.Create(r.Context(), user.GroupID, createData)
		if err != nil {
			log.Err(err).Msg("failed to create item")
			server.RespondServerError(w)
			return
		}

		server.Respond(w, http.StatusCreated, item)

	}
}

// HandleItemDelete godocs
// @Summary   deletes a item
// @Tags      Items
// @Produce   json
// @Param     id   path      string  true  "Item ID"
// @Success   204
// @Router    /v1/items/{id} [DELETE]
// @Security  Bearer
func (ctrl *V1Controller) HandleItemDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, user, err := ctrl.partialParseIdAndUser(w, r)
		if err != nil {
			return
		}

		err = ctrl.svc.Items.Delete(r.Context(), user.GroupID, uid)
		if err != nil {
			log.Err(err).Msg("failed to delete item")
			server.RespondServerError(w)
			return
		}
		server.Respond(w, http.StatusNoContent, nil)
	}
}

// HandleItemGet godocs
// @Summary   Gets a item and fields
// @Tags      Items
// @Produce   json
// @Param     id   path      string  true  "Item ID"
// @Success   200  {object}  types.ItemOut
// @Router    /v1/items/{id} [GET]
// @Security  Bearer
func (ctrl *V1Controller) HandleItemGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, user, err := ctrl.partialParseIdAndUser(w, r)
		if err != nil {
			return
		}

		items, err := ctrl.svc.Items.GetOne(r.Context(), user.GroupID, uid)
		if err != nil {
			log.Err(err).Msg("failed to get item")
			server.RespondServerError(w)
			return
		}
		server.Respond(w, http.StatusOK, items)
	}
}

// HandleItemUpdate godocs
// @Summary   updates a item
// @Tags      Items
// @Produce   json
// @Param     id  path  string  true  "Item ID"
// @Success   200  {object}  types.ItemOut
// @Router    /v1/items/{id} [PUT]
// @Security  Bearer
func (ctrl *V1Controller) HandleItemUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := types.ItemUpdate{}
		if err := server.Decode(r, &body); err != nil {
			log.Err(err).Msg("failed to decode request body")
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}
		uid, user, err := ctrl.partialParseIdAndUser(w, r)
		if err != nil {
			return
		}

		body.ID = uid
		result, err := ctrl.svc.Items.Update(r.Context(), user.GroupID, body)
		if err != nil {
			log.Err(err).Msg("failed to update item")
			server.RespondServerError(w)
			return
		}
		server.Respond(w, http.StatusOK, result)
	}
}
