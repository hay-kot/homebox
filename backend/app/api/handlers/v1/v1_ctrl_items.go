package v1

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/homebox/backend/pkgs/server"
	"github.com/rs/zerolog/log"
)

// HandleItemsGetAll godoc
// @Summary  Get All Items
// @Tags     Items
// @Produce  json
// @Param    q         query    string   false "search string"
// @Param    page      query    int      false "page number"
// @Param    pageSize  query    int      false "items per page"
// @Param    labels    query    []string false "label Ids"    collectionFormat(multi)
// @Param    locations query    []string false "location Ids" collectionFormat(multi)
// @Success  200       {object} repo.PaginationResult[repo.ItemSummary]{}
// @Router   /v1/items [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleItemsGetAll() server.HandlerFunc {
	extractQuery := func(r *http.Request) repo.ItemQuery {
		params := r.URL.Query()

		filterFieldItems := func(raw []string) []repo.FieldQuery {
			var items []repo.FieldQuery

			for _, v := range raw {
				parts := strings.SplitN(v, "=", 2)
				if len(parts) == 2 {
					items = append(items, repo.FieldQuery{
						Name:  parts[0],
						Value: parts[1],
					})
				}
			}

			return items
		}

		v := repo.ItemQuery{
			Page:            queryIntOrNegativeOne(params.Get("page")),
			PageSize:        queryIntOrNegativeOne(params.Get("pageSize")),
			Search:          params.Get("q"),
			LocationIDs:     queryUUIDList(params, "locations"),
			LabelIDs:        queryUUIDList(params, "labels"),
			IncludeArchived: queryBool(params.Get("includeArchived")),
			Fields:          filterFieldItems(params["fields"]),
		}

		if strings.HasPrefix(v.Search, "#") {
			aidStr := strings.TrimPrefix(v.Search, "#")

			aid, ok := repo.ParseAssetID(aidStr)
			if ok {
				v.Search = ""
				v.AssetID = aid
			}
		}

		return v
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())

		items, err := ctrl.repo.Items.QueryByGroup(ctx, ctx.GID, extractQuery(r))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return server.Respond(w, http.StatusOK, repo.PaginationResult[repo.ItemSummary]{
					Items: []repo.ItemSummary{},
				})
			}
			log.Err(err).Msg("failed to get items")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}
		return server.Respond(w, http.StatusOK, items)
	}
}

// HandleItemsCreate godoc
// @Summary  Create a new item
// @Tags     Items
// @Produce  json
// @Param    payload body     repo.ItemCreate true "Item Data"
// @Success  200     {object} repo.ItemSummary
// @Router   /v1/items [POST]
// @Security Bearer
func (ctrl *V1Controller) HandleItemsCreate() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		createData := repo.ItemCreate{}
		if err := server.Decode(r, &createData); err != nil {
			log.Err(err).Msg("failed to decode request body")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		ctx := services.NewContext(r.Context())
		item, err := ctrl.svc.Items.Create(ctx, createData)
		if err != nil {
			log.Err(err).Msg("failed to create item")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusCreated, item)
	}
}

// HandleItemGet godocs
// @Summary  Gets a item and fields
// @Tags     Items
// @Produce  json
// @Param    id  path     string true "Item ID"
// @Success  200 {object} repo.ItemOut
// @Router   /v1/items/{id} [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleItemGet() server.HandlerFunc {
	return ctrl.handleItemsGeneral()
}

// HandleItemDelete godocs
// @Summary  deletes a item
// @Tags     Items
// @Produce  json
// @Param    id path string true "Item ID"
// @Success  204
// @Router   /v1/items/{id} [DELETE]
// @Security Bearer
func (ctrl *V1Controller) HandleItemDelete() server.HandlerFunc {
	return ctrl.handleItemsGeneral()
}

// HandleItemUpdate godocs
// @Summary  updates a item
// @Tags     Items
// @Produce  json
// @Param    id      path     string          true "Item ID"
// @Param    payload body     repo.ItemUpdate true "Item Data"
// @Success  200     {object} repo.ItemOut
// @Router   /v1/items/{id} [PUT]
// @Security Bearer
func (ctrl *V1Controller) HandleItemUpdate() server.HandlerFunc {
	return ctrl.handleItemsGeneral()
}

func (ctrl *V1Controller) handleItemsGeneral() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())
		ID, err := ctrl.routeID(r)
		if err != nil {
			return err
		}

		switch r.Method {
		case http.MethodGet:
			items, err := ctrl.repo.Items.GetOneByGroup(r.Context(), ctx.GID, ID)
			if err != nil {
				log.Err(err).Msg("failed to get item")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}
			return server.Respond(w, http.StatusOK, items)
		case http.MethodDelete:
			err = ctrl.repo.Items.DeleteByGroup(r.Context(), ctx.GID, ID)
			if err != nil {
				log.Err(err).Msg("failed to delete item")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}
			return server.Respond(w, http.StatusNoContent, nil)
		case http.MethodPut:
			body := repo.ItemUpdate{}
			if err := server.Decode(r, &body); err != nil {
				log.Err(err).Msg("failed to decode request body")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}
			body.ID = ID
			result, err := ctrl.repo.Items.UpdateByGroup(r.Context(), ctx.GID, body)
			if err != nil {
				log.Err(err).Msg("failed to update item")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}
			return server.Respond(w, http.StatusOK, result)
		}

		return nil
	}
}

// HandleGetAllCustomFieldNames godocs
// @Summary  imports items into the database
// @Tags     Items
// @Produce  json
// @Success  200
// @Router   /v1/items/fields [GET]
// @Success  200     {object} []string
// @Security Bearer
func (ctrl *V1Controller) HandleGetAllCustomFieldNames() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())

		v, err := ctrl.repo.Items.GetAllCustomFieldNames(r.Context(), ctx.GID)
		if err != nil {
			return err
		}

		return server.Respond(w, http.StatusOK, v)
	}
}

// HandleGetAllCustomFieldValues godocs
// @Summary  imports items into the database
// @Tags     Items
// @Produce  json
// @Success  200
// @Router   /v1/items/fields/values [GET]
// @Success  200     {object} []string
// @Security Bearer
func (ctrl *V1Controller) HandleGetAllCustomFieldValues() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())

		v, err := ctrl.repo.Items.GetAllCustomFieldValues(r.Context(), ctx.GID, r.URL.Query().Get("field"))
		if err != nil {
			return err
		}

		return server.Respond(w, http.StatusOK, v)
	}
}

// HandleItemsImport godocs
// @Summary  imports items into the database
// @Tags     Items
// @Produce  json
// @Success  204
// @Param    csv formData file true "Image to upload"
// @Router   /v1/items/import [Post]
// @Security Bearer
func (ctrl *V1Controller) HandleItemsImport() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := r.ParseMultipartForm(ctrl.maxUploadSize << 20)
		if err != nil {
			log.Err(err).Msg("failed to parse multipart form")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		file, _, err := r.FormFile("csv")
		if err != nil {
			log.Err(err).Msg("failed to get file from form")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		data, err := services.ReadCsv(file)
		if err != nil {
			log.Err(err).Msg("failed to read csv")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		user := services.UseUserCtx(r.Context())

		_, err = ctrl.svc.Items.CsvImport(r.Context(), user.GroupID, data)
		if err != nil {
			log.Err(err).Msg("failed to import items")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusNoContent, nil)
	}
}
