package v1

import (
	"net/http"

	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/internal/services"
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
func (ctrl *V1Controller) HandleLocationGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := services.UseUserCtx(r.Context())
		locations, err := ctrl.svc.Location.GetAll(r.Context(), user.GroupID)
		if err != nil {
			log.Err(err).Msg("failed to get locations")
			server.RespondServerError(w)
			return
		}

		server.Respond(w, http.StatusOK, server.Results{Items: locations})
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
func (ctrl *V1Controller) HandleLocationCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createData := repo.LocationCreate{}
		if err := server.Decode(r, &createData); err != nil {
			log.Err(err).Msg("failed to decode location create data")
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		user := services.UseUserCtx(r.Context())
		location, err := ctrl.svc.Location.Create(r.Context(), user.GroupID, createData)
		if err != nil {
			log.Err(err).Msg("failed to create location")
			server.RespondServerError(w)
			return
		}

		server.Respond(w, http.StatusCreated, location)
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
func (ctrl *V1Controller) HandleLocationDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, user, err := ctrl.partialParseIdAndUser(w, r)
		if err != nil {
			return
		}

		err = ctrl.svc.Location.Delete(r.Context(), user.GroupID, uid)
		if err != nil {
			log.Err(err).Msg("failed to delete location")
			server.RespondServerError(w)
			return
		}
		server.Respond(w, http.StatusNoContent, nil)
	}
}

// HandleLocationGet godocs
// @Summary  Gets a location and fields
// @Tags     Locations
// @Produce  json
// @Param    id  path     string true "Location ID"
// @Success  200 {object} repo.LocationOut
// @Router   /v1/locations/{id} [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleLocationGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, user, err := ctrl.partialParseIdAndUser(w, r)
		if err != nil {
			return
		}

		location, err := ctrl.svc.Location.GetOne(r.Context(), user.GroupID, uid)
		if err != nil {
			if ent.IsNotFound(err) {
				log.Err(err).
					Str("id", uid.String()).
					Str("gid", user.GroupID.String()).
					Msg("location not found")
				server.RespondError(w, http.StatusNotFound, err)
				return
			}

			log.Err(err).
				Str("id", uid.String()).
				Str("gid", user.GroupID.String()).
				Msg("failed to get location")
			server.RespondServerError(w)
			return
		}
		server.Respond(w, http.StatusOK, location)
	}
}

// HandleLocationUpdate godocs
// @Summary  updates a location
// @Tags     Locations
// @Produce  json
// @Param    id  path     string true "Location ID"
// @Success  200 {object} repo.LocationOut
// @Router   /v1/locations/{id} [PUT]
// @Security Bearer
func (ctrl *V1Controller) HandleLocationUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := repo.LocationUpdate{}
		if err := server.Decode(r, &body); err != nil {
			log.Err(err).Msg("failed to decode location update data")
			server.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		uid, user, err := ctrl.partialParseIdAndUser(w, r)
		if err != nil {
			return
		}

		body.ID = uid

		result, err := ctrl.svc.Location.Update(r.Context(), user.GroupID, body)
		if err != nil {
			log.Err(err).Msg("failed to update location")
			server.RespondServerError(w)
			return
		}
		server.Respond(w, http.StatusOK, result)
	}
}
