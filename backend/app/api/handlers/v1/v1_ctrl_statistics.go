package v1

import (
	"net/http"
	"time"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/homebox/backend/pkgs/server"
)

// HandleGroupGet godoc
// @Summary  Get the current user's group statistics
// @Tags     Statistics
// @Produce  json
// @Success  200 {object} []repo.TotalsByOrganizer
// @Router   /v1/groups/statistics/locations [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleGroupStatisticsLocations() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())

		stats, err := ctrl.repo.Groups.StatsLocationsByPurchasePrice(ctx, ctx.GID)
		if err != nil {
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusOK, stats)
	}
}

// HandleGroupGet godoc
// @Summary  Get the current user's group statistics
// @Tags     Statistics
// @Produce  json
// @Success  200 {object} []repo.TotalsByOrganizer
// @Router   /v1/groups/statistics/labels [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleGroupStatisticsLabels() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())

		stats, err := ctrl.repo.Groups.StatsLabelsByPurchasePrice(ctx, ctx.GID)
		if err != nil {
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusOK, stats)
	}
}

// HandleGroupGet godoc
// @Summary  Get the current user's group statistics
// @Tags     Statistics
// @Produce  json
// @Success  200 {object} repo.GroupStatistics
// @Router   /v1/groups/statistics [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleGroupStatistics() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())

		stats, err := ctrl.repo.Groups.StatsGroup(ctx, ctx.GID)
		if err != nil {
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusOK, stats)
	}
}

// HandleGroupGet godoc
// @Summary  Queries the changes overtime of the purchase price over time
// @Tags     Statistics
// @Produce  json
// @Success  200 {object} repo.ValueOverTime
// @Param 	 start query string false "start date"
// @Param 	 end query string false "end date"
// @Router   /v1/groups/statistics/purchase-price [GET]
// @Security Bearer
func (ctrl *V1Controller) HandleGroupStatisticsPriceOverTime() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())

		startDate, err := time.Parse("2006-01-02", r.URL.Query().Get("start"))
		if err != nil {
			return validate.NewRequestError(err, http.StatusBadRequest)
		}

		endDate, err := time.Parse("2006-01-02", r.URL.Query().Get("end"))
		if err != nil {
			endDate = time.Now()
		}

		stats, err := ctrl.repo.Groups.StatsPurchasePrice(ctx, ctx.GID, startDate, endDate)
		if err != nil {
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.Respond(w, http.StatusOK, stats)
	}
}
