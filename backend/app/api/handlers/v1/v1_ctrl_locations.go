package v1

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/homebox/backend/internal/web/adapters"
	"github.com/hay-kot/httpkit/errchain"
)

// HandleLocationTreeQuery godoc
//
//	@Summary  Get Locations Tree
//	@Tags     Locations
//	@Produce  json
//	@Param    withItems         query    bool   false "include items in response tree"
//	@Success  200 {object} []repo.TreeItem
//	@Router   /v1/locations/tree [GET]
//	@Security Bearer
func (ctrl *V1Controller) HandleLocationTreeQuery() errchain.HandlerFunc {
	fn := func(r *http.Request, query repo.TreeQuery) ([]repo.TreeItem, error) {
		auth := services.NewContext(r.Context())
		return ctrl.repo.Locations.Tree(auth, auth.GID, query)
	}

	return adapters.Query(fn, http.StatusOK)
}

// HandleLocationGetAll godoc
//
//	@Summary  Get All Locations
//	@Tags     Locations
//	@Produce  json
//	@Param    filterChildren query bool false "Filter locations with parents"
//	@Success  200 {object} []repo.LocationOutCount
//	@Router   /v1/locations [GET]
//	@Security Bearer
func (ctrl *V1Controller) HandleLocationGetAll() errchain.HandlerFunc {
	fn := func(r *http.Request, q repo.LocationQuery) ([]repo.LocationOutCount, error) {
		auth := services.NewContext(r.Context())
		return ctrl.repo.Locations.GetAll(auth, auth.GID, q)
	}

	return adapters.Query(fn, http.StatusOK)
}

// HandleLocationCreate godoc
//
//	@Summary  Create Location
//	@Tags     Locations
//	@Produce  json
//	@Param    payload body     repo.LocationCreate true "Location Data"
//	@Success  200     {object} repo.LocationSummary
//	@Router   /v1/locations [POST]
//	@Security Bearer
func (ctrl *V1Controller) HandleLocationCreate() errchain.HandlerFunc {
	fn := func(r *http.Request, createData repo.LocationCreate) (repo.LocationOut, error) {
		auth := services.NewContext(r.Context())
		return ctrl.repo.Locations.Create(auth, auth.GID, createData)
	}

	return adapters.Action(fn, http.StatusCreated)
}

// HandleLocationDelete godoc
//
//	@Summary  Delete Location
//	@Tags     Locations
//	@Produce  json
//	@Param    id path string true "Location ID"
//	@Success  204
//	@Router   /v1/locations/{id} [DELETE]
//	@Security Bearer
func (ctrl *V1Controller) HandleLocationDelete() errchain.HandlerFunc {
	fn := func(r *http.Request, ID uuid.UUID) (any, error) {
		auth := services.NewContext(r.Context())
		err := ctrl.repo.Locations.DeleteByGroup(auth, auth.GID, ID)
		return nil, err
	}

	return adapters.CommandID("id", fn, http.StatusNoContent)
}

// HandleLocationGet godoc
//
//	@Summary  Get Location
//	@Tags     Locations
//	@Produce  json
//	@Param    id  path     string true "Location ID"
//	@Success  200 {object} repo.LocationOut
//	@Router   /v1/locations/{id} [GET]
//	@Security Bearer
func (ctrl *V1Controller) HandleLocationGet() errchain.HandlerFunc {
	fn := func(r *http.Request, ID uuid.UUID) (repo.LocationOut, error) {
		auth := services.NewContext(r.Context())
		return ctrl.repo.Locations.GetOneByGroup(auth, auth.GID, ID)
	}

	return adapters.CommandID("id", fn, http.StatusOK)
}

// HandleLocationUpdate godoc
//
//	@Summary  Update Location
//	@Tags     Locations
//	@Produce  json
//	@Param    id      path     string              true "Location ID"
//	@Param    payload body     repo.LocationUpdate true "Location Data"
//	@Success  200     {object} repo.LocationOut
//	@Router   /v1/locations/{id} [PUT]
//	@Security Bearer
func (ctrl *V1Controller) HandleLocationUpdate() errchain.HandlerFunc {
	fn := func(r *http.Request, ID uuid.UUID, body repo.LocationUpdate) (repo.LocationOut, error) {
		auth := services.NewContext(r.Context())
		body.ID = ID
		return ctrl.repo.Locations.UpdateByGroup(auth, auth.GID, ID, body)
	}

	return adapters.ActionID("id", fn, http.StatusOK)
}
