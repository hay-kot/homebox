package v1

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/ent/attachment"
	"github.com/hay-kot/homebox/backend/internal/services"
	"github.com/hay-kot/homebox/backend/internal/types"
	"github.com/hay-kot/homebox/backend/pkgs/server"
	"github.com/rs/zerolog/log"
)

// HandleItemsGetAll godoc
// @Summary   Get All Items
// @Tags      Items
// @Produce   json
// @Success   200  {object}  server.Results{items=[]types.ItemSummary}
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
// @Param     id       path      string            true  "Item ID"
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
// @Success   200      {object}  types.ItemOut
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
// @Param     payload  body      types.ItemUpdate  true  "Item Data"
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

// HandleItemsImport godocs
// @Summary   imports items into the database
// @Tags      Items
// @Produce   json
// @Success   204
// @Param     csv  formData  file  true  "Image to upload"
// @Router    /v1/items/import [Post]
// @Security  Bearer
func (ctrl *V1Controller) HandleItemsImport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Max upload size of 10 MB - TODO: Set via config
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			log.Err(err).Msg("failed to parse multipart form")
			server.RespondServerError(w)
			return
		}

		file, _, err := r.FormFile("csv")
		if err != nil {
			log.Err(err).Msg("failed to get file from form")
			server.RespondServerError(w)
			return
		}

		reader := csv.NewReader(file)
		data, err := reader.ReadAll()
		if err != nil {
			log.Err(err).Msg("failed to read csv")
			server.RespondServerError(w)
			return
		}

		user := services.UseUserCtx(r.Context())

		err = ctrl.svc.Items.CsvImport(r.Context(), user.GroupID, data)
		if err != nil {
			log.Err(err).Msg("failed to import items")
			server.RespondServerError(w)
			return
		}

		server.Respond(w, http.StatusNoContent, nil)
	}
}

// HandleItemsImport godocs
// @Summary   imports items into the database
// @Tags      Items
// @Produce   json
// @Param     id    path      string  true  "Item ID"
// @Param     file  formData  file    true  "File attachment"
// @Param     type  formData  string  true  "Type of file"
// @Param     name  formData  string  true  "name of the file including extension"
// @Success   200   {object}  types.ItemOut
// @Router    /v1/items/{id}/attachments [POST]
// @Security  Bearer
func (ctrl *V1Controller) HandleItemAttachmentCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Max upload size of 10 MB - TODO: Set via config
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			log.Err(err).Msg("failed to parse multipart form")
			server.RespondServerError(w)
			return
		}
		file, _, err := r.FormFile("file")
		if err != nil {
			log.Err(err).Msg("failed to get file from form")
			server.RespondServerError(w)
			return
		}
		attachmentName := r.FormValue("name")
		if attachmentName == "" {
			log.Err(err).Msg("failed to get name from form")
			server.RespondError(w, http.StatusBadRequest, errors.New("name is required"))
		}

		attachmentType := r.FormValue("type")
		if attachmentType == "" {
			attachmentName = "attachment"
		}

		uid, user, err := ctrl.partialParseIdAndUser(w, r)
		if err != nil {
			return
		}

		item, err := ctrl.svc.Items.AddAttachment(
			r.Context(),
			user.GroupID,
			uid,
			attachmentName,
			attachment.Type(attachmentType),
			file,
		)

		if err != nil {
			log.Err(err).Msg("failed to add attachment")
			server.RespondServerError(w)
			return
		}

		server.Respond(w, http.StatusOK, item)
	}
}

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

		path, err := ctrl.svc.Items.GetAttachment(r.Context(), token)

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
	return func(w http.ResponseWriter, r *http.Request) {
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

		token, err := ctrl.svc.Items.NewAttachmentToken(r.Context(), user.GroupID, uid, attachmentId)

		if err != nil {
			log.Err(err).Msg("failed to get attachment")
			server.RespondServerError(w)
			return
		}

		server.Respond(w, http.StatusOK, types.ItemAttachmentToken{
			Token: token,
		})

	}
}
