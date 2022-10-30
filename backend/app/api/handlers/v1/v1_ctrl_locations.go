package v1

import (
	"net/http"

	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/internal/services"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/homebox/backend/pkgs/server"
	"github.com/rs/zerolog/log"
)

// HandleLocationGetAll godoc
// @Summary  Get All Locations
// @Tags     Locations
// @Produce  json
// @Success  200 {object} server.Results{items=[]repo.LocationOutCount}
// @Router   /v1/locations [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleLocationGetAll() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		user := services.UseUserCtx(r.Context())
		locations, err := ctrl.repo.Locations.GetAll(r.Context(), user.GroupID)
		if err != nil {
			log.Err(err).Msg("failed to get locations")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusOK, server.Results{Items: locations})
	}
}

// HandleLocationCreate godoc
// @Summary  Create a new location
// @Tags     Locations
// @Produce  json
// @Param    payload body     repo.LocationCreate true "Location Data"
// @Success  200     {object} repo.LocationSummary
// @Router   /v1/locations [POST]
// @Security Bearer
func (ctrl *V1Controller) HandleLocationCreate() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		createData := repo.LocationCreate{}
		if err := server.Decode(r, &createData); err != nil {
			log.Err(err).Msg("failed to decode location create data")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		user := services.UseUserCtx(r.Context())
		location, err := ctrl.repo.Locations.Create(r.Context(), user.GroupID, createData)
		if err != nil {
			log.Err(err).Msg("failed to create location")
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusCreated, location)
	}
}

// HandleLocationDelete godocs
// @Summary  deletes a location
// @Tags     Locations
// @Produce  json
// @Param    id path string true "Location ID"
// @Success  204
// @Router   /v1/locations/{id} [DELETE]
// @Security Bearer
func (ctrl *V1Controller) HandleLocationDelete() server.HandlerFunc {
	return ctrl.handleLocationGeneral()
}

// HandleLocationGet godocs
// @Summary  Gets a location and fields
// @Tags     Locations
// @Produce  json
// @Param    id  path     string true "Location ID"
// @Success  200 {object} repo.LocationOut
// @Router   /v1/locations/{id} [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleLocationGet() server.HandlerFunc {
	return ctrl.handleLocationGeneral()
}

// HandleLocationUpdate godocs
// @Summary  updates a location
// @Tags     Locations
// @Produce  json
// @Param    id      path     string              true "Location ID"
// @Param    payload body     repo.LocationUpdate true "Location Data"
// @Success  200     {object} repo.LocationOut
// @Router   /v1/locations/{id} [PUT]
// @Security Bearer
func (ctrl *V1Controller) HandleLocationUpdate() server.HandlerFunc {
	return ctrl.handleLocationGeneral()
}

func (ctrl *V1Controller) handleLocationGeneral() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())
		ID, err := ctrl.routeID(r)
		if err != nil {
			return err
		}

		switch r.Method {
		case http.MethodGet:
			location, err := ctrl.repo.Locations.GetOneByGroup(r.Context(), ctx.GID, ID)
			if err != nil {
				l := log.Err(err).
					Str("ID", ID.String()).
					Str("GID", ctx.GID.String())

				if ent.IsNotFound(err) {
					l.Msg("location not found")
					return validate.NewRequestError(err, http.StatusNotFound)
				}

				l.Msg("failed to get location")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}
			return server.Respond(w, http.StatusOK, location)
		case http.MethodPut:
			body := repo.LocationUpdate{}
			if err := server.Decode(r, &body); err != nil {
				log.Err(err).Msg("failed to decode location update data")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}

			body.ID = ID

			result, err := ctrl.repo.Locations.UpdateOneByGroup(r.Context(), ctx.GID, ID, body)
			if err != nil {
				log.Err(err).Msg("failed to update location")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}
			return server.Respond(w, http.StatusOK, result)
		case http.MethodDelete:
			err = ctrl.repo.Locations.DeleteByGroup(r.Context(), ctx.GID, ID)
			if err != nil {
				log.Err(err).Msg("failed to delete location")
				return validate.NewRequestError(err, http.StatusInternalServerError)
			}
			return server.Respond(w, http.StatusNoContent, nil)
		}
		return nil
	}
}
