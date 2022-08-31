package v1

import (
	"net/http"

	"github.com/hay-kot/content/backend/internal/services"
	"github.com/hay-kot/content/backend/internal/types"
	"github.com/hay-kot/content/backend/pkgs/server"
)

// HandleUserSelf godoc
// @Summary   Get All Locations
// @Tags      Locations
// @Produce   json
// @Success   200  {object}  server.Results{items=[]types.LocationOut}
// @Router    /v1/locations [GET]
// @Security  Bearer
func (ctrl *V1Controller) HandleLocationGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := services.UseUserCtx(r.Context())
		locations, err := ctrl.svc.Location.GetAll(r.Context(), user.GroupID)
		if err != nil {
			ctrl.log.Error(err, nil)
			server.RespondServerError(w)
			return
		}

		server.Respond(w, http.StatusOK, server.Results{Items: locations})
	}
}

// HandleUserSelf godoc
// @Summary   Create a new location
// @Tags      Locations
// @Produce   json
// @Param     payload  body      types.LocationCreate  true  "Location Data"
// @Success   200      {object}  types.LocationOut
// @Router    /v1/locations [POST]
// @Security  Bearer
func (ctrl *V1Controller) HandleLocationCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createData := types.LocationCreate{}
		if err := server.Decode(r, &createData); err != nil {
			ctrl.log.Error(err, nil)
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		user := services.UseUserCtx(r.Context())
		location, err := ctrl.svc.Location.Create(r.Context(), user.GroupID, createData)
		if err != nil {
			ctrl.log.Error(err, nil)
			server.RespondServerError(w)
			return
		}

		server.Respond(w, http.StatusCreated, location)
	}
}
