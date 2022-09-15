package v1

import (
	"net/http"

	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/internal/services"
	"github.com/hay-kot/homebox/backend/internal/types"
	"github.com/hay-kot/homebox/backend/pkgs/server"
	"github.com/rs/zerolog/log"
)

// HandleLabelsGetAll godoc
// @Summary   Get All Labels
// @Tags      Labels
// @Produce   json
// @Success   200  {object}  server.Results{items=[]types.LabelOut}
// @Router    /v1/labels [GET]
// @Security  Bearer
func (ctrl *V1Controller) HandleLabelsGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := services.UseUserCtx(r.Context())
		labels, err := ctrl.svc.Labels.GetAll(r.Context(), user.GroupID)
		if err != nil {
			log.Err(err).Msg("error getting labels")
			server.RespondServerError(w)
			return
		}
		server.Respond(w, http.StatusOK, server.Results{Items: labels})
	}
}

// HandleLabelsCreate godoc
// @Summary   Create a new label
// @Tags      Labels
// @Produce   json
// @Param     payload  body      types.LabelCreate  true  "Label Data"
// @Success   200      {object}  types.LabelSummary
// @Router    /v1/labels [POST]
// @Security  Bearer
func (ctrl *V1Controller) HandleLabelsCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createData := types.LabelCreate{}
		if err := server.Decode(r, &createData); err != nil {
			log.Err(err).Msg("error decoding label create data")
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		user := services.UseUserCtx(r.Context())
		label, err := ctrl.svc.Labels.Create(r.Context(), user.GroupID, createData)
		if err != nil {
			log.Err(err).Msg("error creating label")
			server.RespondServerError(w)
			return
		}

		server.Respond(w, http.StatusCreated, label)

	}
}

// HandleLabelDelete godocs
// @Summary   deletes a label
// @Tags      Labels
// @Produce   json
// @Param     id   path      string  true  "Label ID"
// @Success   204
// @Router    /v1/labels/{id} [DELETE]
// @Security  Bearer
func (ctrl *V1Controller) HandleLabelDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, user, err := ctrl.partialParseIdAndUser(w, r)
		if err != nil {
			return
		}

		err = ctrl.svc.Labels.Delete(r.Context(), user.GroupID, uid)
		if err != nil {
			log.Err(err).Msg("error deleting label")
			server.RespondServerError(w)
			return
		}
		server.Respond(w, http.StatusNoContent, nil)
	}
}

// HandleLabelGet godocs
// @Summary   Gets a label and fields
// @Tags      Labels
// @Produce   json
// @Param     id   path      string  true  "Label ID"
// @Success   200  {object}  types.LabelOut
// @Router    /v1/labels/{id} [GET]
// @Security  Bearer
func (ctrl *V1Controller) HandleLabelGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, user, err := ctrl.partialParseIdAndUser(w, r)
		if err != nil {
			return
		}

		labels, err := ctrl.svc.Labels.Get(r.Context(), user.GroupID, uid)
		if err != nil {
			if ent.IsNotFound(err) {
				log.Err(err).
					Str("id", uid.String()).
					Msg("label not found")
				server.RespondError(w, http.StatusNotFound, err)
				return
			}
			log.Err(err).Msg("error getting label")
			server.RespondServerError(w)
			return
		}
		server.Respond(w, http.StatusOK, labels)
	}
}

// HandleLabelUpdate godocs
// @Summary   updates a label
// @Tags      Labels
// @Produce   json
// @Param     id  path  string  true  "Label ID"
// @Success   200  {object}  types.LabelOut
// @Router    /v1/labels/{id} [PUT]
// @Security  Bearer
func (ctrl *V1Controller) HandleLabelUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := types.LabelUpdate{}
		if err := server.Decode(r, &body); err != nil {
			log.Err(err).Msg("error decoding label update data")
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}
		uid, user, err := ctrl.partialParseIdAndUser(w, r)
		if err != nil {
			return
		}

		body.ID = uid
		result, err := ctrl.svc.Labels.Update(r.Context(), user.GroupID, body)
		if err != nil {
			log.Err(err).Msg("error updating label")
			server.RespondServerError(w)
			return
		}
		server.Respond(w, http.StatusOK, result)
	}
}
