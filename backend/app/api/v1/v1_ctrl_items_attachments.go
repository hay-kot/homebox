package v1

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/services"
	"github.com/hay-kot/homebox/backend/internal/types"
	"github.com/hay-kot/homebox/backend/pkgs/server"
	"github.com/rs/zerolog/log"
)

// HandleItemAttachmentGet godocs
// @Summary   retrieves an attachment for an item
// @Tags      Items
// @Produce   application/octet-stream
// @Param     id     path   string  true  "Item ID"
// @Param     token  query  string  true  "Attachment token"
// @Success   200
// @Router    /v1/items/{id}/attachments/download [GET]
// @Security  Bearer
func (ctrl *V1Controller) HandleItemAttachmentDownload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := server.GetParam(r, "token", "")

		path, err := ctrl.svc.Items.AttachmentPath(r.Context(), token)

		if err != nil {
			log.Err(err).Msg("failed to get attachment")
			server.RespondServerError(w)
			return
		}

		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filepath.Base(path)))
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeFile(w, r, path)
	}
}

// HandleItemAttachmentToken godocs
// @Summary   retrieves an attachment for an item
// @Tags      Items
// @Produce   application/octet-stream
// @Param     id             path      string  true  "Item ID"
// @Param     attachment_id  path      string  true  "Attachment ID"
// @Success   200            {object}  types.ItemAttachmentToken
// @Router    /v1/items/{id}/attachments/{attachment_id} [GET]
// @Security  Bearer
func (ctrl *V1Controller) HandleItemAttachmentToken() http.HandlerFunc {
	return ctrl.handleItemAttachmentsHandler
}

// HandleItemAttachmentDelete godocs
// @Summary   retrieves an attachment for an item
// @Tags      Items
// @Param     id             path  string  true  "Item ID"
// @Param     attachment_id  path  string  true  "Attachment ID"
// @Success   204
// @Router    /v1/items/{id}/attachments/{attachment_id} [DELETE]
// @Security  Bearer
func (ctrl *V1Controller) HandleItemAttachmentDelete() http.HandlerFunc {
	return ctrl.handleItemAttachmentsHandler
}

// HandleItemAttachmentUpdate godocs
// @Summary   retrieves an attachment for an item
// @Tags      Items
// @Param     id             path      string                      true  "Item ID"
// @Param     attachment_id  path      string                      true  "Attachment ID"
// @Param     payload        body      types.ItemAttachmentUpdate  true  "Attachment Update"
// @Success   200            {object}  types.ItemOut
// @Router    /v1/items/{id}/attachments/{attachment_id} [PUT]
// @Security  Bearer
func (ctrl *V1Controller) HandleItemAttachmentUpdate() http.HandlerFunc {
	return ctrl.handleItemAttachmentsHandler
}

func (ctrl *V1Controller) handleItemAttachmentsHandler(w http.ResponseWriter, r *http.Request) {
	uid, user, err := ctrl.partialParseIdAndUser(w, r)
	if err != nil {
		return
	}

	attachmentId, err := uuid.Parse(chi.URLParam(r, "attachment_id"))
	if err != nil {
		log.Err(err).Msg("failed to parse attachment_id param")
		server.RespondError(w, http.StatusBadRequest, err)
		return
	}

	ctx := services.NewContext(r.Context())

	switch r.Method {

	// Token Handler
	case http.MethodGet:
		token, err := ctrl.svc.Items.AttachmentToken(ctx, uid, attachmentId)
		if err != nil {
			switch err {
			case services.ErrNotFound:
				log.Err(err).
					Str("id", attachmentId.String()).
					Msg("failed to find attachment with id")

				server.RespondError(w, http.StatusNotFound, err)

			case services.ErrFileNotFound:
				log.Err(err).
					Str("id", attachmentId.String()).
					Msg("failed to find file path for attachment with id")
				log.Warn().Msg("attachment with no file path removed from database")

				server.RespondError(w, http.StatusNotFound, err)

			default:
				log.Err(err).Msg("failed to get attachment")
				server.RespondServerError(w)
				return
			}
		}

		server.Respond(w, http.StatusOK, types.ItemAttachmentToken{Token: token})

	// Delete Attachment Handler
	case http.MethodDelete:
		err = ctrl.svc.Items.AttachmentDelete(r.Context(), user.GroupID, uid, attachmentId)
		if err != nil {
			log.Err(err).Msg("failed to delete attachment")
			server.RespondServerError(w)
			return
		}

		server.Respond(w, http.StatusNoContent, nil)

	// Update Attachment Handler
	case http.MethodPut:
		var attachment types.ItemAttachmentUpdate
		err = server.Decode(r, &attachment)
		if err != nil {
			log.Err(err).Msg("failed to decode attachment")
			server.RespondError(w, http.StatusBadRequest, err)
			return
		}

		attachment.ID = attachmentId
		val, err := ctrl.svc.Items.AttachmentUpdate(ctx, uid, &attachment)
		if err != nil {
			log.Err(err).Msg("failed to delete attachment")
			server.RespondServerError(w)
			return
		}

		server.Respond(w, http.StatusOK, val)
	}
}
