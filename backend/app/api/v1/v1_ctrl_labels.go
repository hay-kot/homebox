package v1

import (
	"net/http"
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
	}
}
