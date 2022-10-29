package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/ent/attachment"
	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/internal/services"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/homebox/backend/pkgs/server"
	"github.com/rs/zerolog/log"
)

type (
	ItemAttachmentToken struct {
		Token string `json:"token"`
	}
)

// HandleItemsImport godocs
// @Summary  imports items into the database
// @Tags     Items Attachments
// @Produce  json
// @Param    id   path     string true "Item ID"
// @Param    file formData file   true "File attachment"
// @Param    type formData string true "Type of file"
// @Param    name formData string true "name of the file including extension"
// @Success  200  {object} repo.ItemOut
// @Failure  422  {object} []server.ValidationError
// @Router   /v1/items/{id}/attachments [POST]
// @Security Bearer
func (ctrl *V1Controller) HandleItemAttachmentCreate() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := r.ParseMultipartForm(ctrl.maxUploadSize << 20)
		if err != nil {
			log.Err(err).Msg("failed to parse multipart form")
			return validate.NewRequestError(errors.New("failed to parse multipart form"), http.StatusBadRequest)

		}

		errs := validate.NewFieldErrors()

		file, _, err := r.FormFile("file")
		if err != nil {
			switch {
			case errors.Is(err, http.ErrMissingFile):
				log.Debug().Msg("file for attachment is missing")
				errs = errs.Append("file", "file is required")
			default:
				log.Err(err).Msg("failed to get file from form")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}
		}

		attachmentName := r.FormValue("name")
		if attachmentName == "" {
			log.Debug().Msg("failed to get name from form")
			errs = errs.Append("name", "name is required")
		}

		if !errs.Nil() {
			return server.Respond(w, http.StatusUnprocessableEntity, errs)
		}

		attachmentType := r.FormValue("type")
		if attachmentType == "" {
			attachmentType = attachment.TypeAttachment.String()
		}

		id, err := ctrl.routeID(w, r)
		if err != nil {
			return err
		}

		ctx := services.NewContext(r.Context())

		item, err := ctrl.svc.Items.AttachmentAdd(
			ctx,
			id,
			attachmentName,
			attachment.Type(attachmentType),
			file,
		)

		if err != nil {
			log.Err(err).Msg("failed to add attachment")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusCreated, item)
	}
}

// HandleItemAttachmentGet godocs
// @Summary  retrieves an attachment for an item
// @Tags     Items Attachments
// @Produce  application/octet-stream
// @Param    id    path  string true "Item ID"
// @Param    token query string true "Attachment token"
// @Success  200
// @Router   /v1/items/{id}/attachments/download [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleItemAttachmentDownload() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		token := server.GetParam(r, "token", "")

		doc, err := ctrl.svc.Items.AttachmentPath(r.Context(), token)

		if err != nil {
			log.Err(err).Msg("failed to get attachment")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", doc.Title))
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeFile(w, r, doc.Path)
		return nil
	}
}

// HandleItemAttachmentToken godocs
// @Summary  retrieves an attachment for an item
// @Tags     Items Attachments
// @Produce  application/octet-stream
// @Param    id            path     string true "Item ID"
// @Param    attachment_id path     string true "Attachment ID"
// @Success  200           {object} ItemAttachmentToken
// @Router   /v1/items/{id}/attachments/{attachment_id} [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleItemAttachmentToken() server.HandlerFunc {
	return ctrl.handleItemAttachmentsHandler
}

// HandleItemAttachmentDelete godocs
// @Summary  retrieves an attachment for an item
// @Tags     Items Attachments
// @Param    id            path string true "Item ID"
// @Param    attachment_id path string true "Attachment ID"
// @Success  204
// @Router   /v1/items/{id}/attachments/{attachment_id} [DELETE]
// @Security Bearer
func (ctrl *V1Controller) HandleItemAttachmentDelete() server.HandlerFunc {
	return ctrl.handleItemAttachmentsHandler
}

// HandleItemAttachmentUpdate godocs
// @Summary  retrieves an attachment for an item
// @Tags     Items Attachments
// @Param    id            path     string                    true "Item ID"
// @Param    attachment_id path     string                    true "Attachment ID"
// @Param    payload       body     repo.ItemAttachmentUpdate true "Attachment Update"
// @Success  200           {object} repo.ItemOut
// @Router   /v1/items/{id}/attachments/{attachment_id} [PUT]
// @Security Bearer
func (ctrl *V1Controller) HandleItemAttachmentUpdate() server.HandlerFunc {
	return ctrl.handleItemAttachmentsHandler
}

func (ctrl *V1Controller) handleItemAttachmentsHandler(w http.ResponseWriter, r *http.Request) error {
	ID, err := ctrl.routeID(w, r)
	if err != nil {
		return err
	}

	attachmentId, err := uuid.Parse(chi.URLParam(r, "attachment_id"))
	if err != nil {
		log.Err(err).Msg("failed to parse attachment_id param")
		return validate.NewRequestError(err, http.StatusBadRequest)
	}

	ctx := services.NewContext(r.Context())

	switch r.Method {

	// Token Handler
	case http.MethodGet:
		token, err := ctrl.svc.Items.AttachmentToken(ctx, ID, attachmentId)
		if err != nil {
			switch err {
			case services.ErrNotFound:
				log.Err(err).
					Str("id", attachmentId.String()).
					Msg("failed to find attachment with id")

				return validate.NewRequestError(err, http.StatusNotFound)

			case services.ErrFileNotFound:
				log.Err(err).
					Str("id", attachmentId.String()).
					Msg("failed to find file path for attachment with id")
				log.Warn().Msg("attachment with no file path removed from database")

				return validate.NewRequestError(err, http.StatusNotFound)

			default:
				log.Err(err).Msg("failed to get attachment")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}
		}

		return server.Respond(w, http.StatusOK, ItemAttachmentToken{Token: token})

	// Delete Attachment Handler
	case http.MethodDelete:
		err = ctrl.svc.Items.AttachmentDelete(r.Context(), ctx.GID, ID, attachmentId)
		if err != nil {
			log.Err(err).Msg("failed to delete attachment")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusNoContent, nil)

	// Update Attachment Handler
	case http.MethodPut:
		var attachment repo.ItemAttachmentUpdate
		err = server.Decode(r, &attachment)
		if err != nil {
			log.Err(err).Msg("failed to decode attachment")
			return validate.NewRequestError(err, http.StatusBadRequest)
		}

		attachment.ID = attachmentId
		val, err := ctrl.svc.Items.AttachmentUpdate(ctx, ID, &attachment)
		if err != nil {
			log.Err(err).Msg("failed to delete attachment")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusOK, val)
	}

	return nil
}
