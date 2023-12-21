package v1

import (
	"net/http"
	"time"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/homebox/backend/internal/sys/validate"
	"github.com/hay-kot/homebox/backend/internal/web/adapters"
	"github.com/hay-kot/httpkit/errchain"
	"github.com/hay-kot/httpkit/server"
)

// HandleGroupStatisticsLocations godoc
//
//	@Summary  Get Location Statistics
//	@Tags     Statistics
//	@Produce  json
//	@Success  200 {object} []repo.TotalsByOrganizer
//	@Router   /v1/groups/statistics/locations [GET]
//	@Security Bearer
func (ctrl *V1Controller) HandleGroupStatisticsLocations() errchain.HandlerFunc {
	fn := func(r *http.Request) ([]repo.TotalsByOrganizer, error) {
		auth := services.NewContext(r.Context())
		return ctrl.repo.Groups.StatsLocationsByPurchasePrice(auth, auth.GID)
	}

	return adapters.Command(fn, http.StatusOK)
}

// HandleGroupStatisticsLabels godoc
//
//	@Summary  Get Label Statistics
//	@Tags     Statistics
//	@Produce  json
//	@Success  200 {object} []repo.TotalsByOrganizer
//	@Router   /v1/groups/statistics/labels [GET]
//	@Security Bearer
func (ctrl *V1Controller) HandleGroupStatisticsLabels() errchain.HandlerFunc {
	fn := func(r *http.Request) ([]repo.TotalsByOrganizer, error) {
		auth := services.NewContext(r.Context())
		return ctrl.repo.Groups.StatsLabelsByPurchasePrice(auth, auth.GID)
	}

	return adapters.Command(fn, http.StatusOK)
}

// HandleGroupStatistics godoc
//
//	@Summary  Get Group Statistics
//	@Tags     Statistics
//	@Produce  json
//	@Success  200 {object} repo.GroupStatistics
//	@Router   /v1/groups/statistics [GET]
//	@Security Bearer
func (ctrl *V1Controller) HandleGroupStatistics() errchain.HandlerFunc {
	fn := func(r *http.Request) (repo.GroupStatistics, error) {
		auth := services.NewContext(r.Context())
		return ctrl.repo.Groups.StatsGroup(auth, auth.GID)
	}

	return adapters.Command(fn, http.StatusOK)
}

// HandleGroupStatisticsPriceOverTime godoc
//
//	@Summary  Get Purchase Price Statistics
//	@Tags     Statistics
//	@Produce  json
//	@Success  200 {object} repo.ValueOverTime
//	@Param 	 start query string false "start date"
//	@Param 	 end query string false "end date"
//	@Router   /v1/groups/statistics/purchase-price [GET]
//	@Security Bearer
func (ctrl *V1Controller) HandleGroupStatisticsPriceOverTime() errchain.HandlerFunc {
	parseDate := func(datestr string, defaultDate time.Time) (time.Time, error) {
		if datestr == "" {
			return defaultDate, nil
		}
		return time.Parse("2006-01-02", datestr)
	}

	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := services.NewContext(r.Context())

		startDate, err := parseDate(r.URL.Query().Get("start"), time.Now().AddDate(0, -1, 0))
		if err != nil {
			return validate.NewRequestError(err, http.StatusBadRequest)
		}

		endDate, err := parseDate(r.URL.Query().Get("end"), time.Now())
		if err != nil {
			return validate.NewRequestError(err, http.StatusBadRequest)
		}

		stats, err := ctrl.repo.Groups.StatsPurchasePrice(ctx, ctx.GID, startDate, endDate)
		if err != nil {
			return validate.NewRequestError(err, http.StatusInternalServerError)
		}

		return server.JSON(w, http.StatusOK, stats)
	}
}
