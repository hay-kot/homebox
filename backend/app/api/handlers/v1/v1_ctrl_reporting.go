package v1

import (
	"net/http"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/pkgs/server"
)

// HandleBillOfMaterialsExport godoc
//
//	@Summary  Generates a Bill of Materials CSV
//	@Tags     Reporting
//	@Produce  json
//	@Success 200 {string} string "text/csv"
//	@Router   /v1/reporting/bill-of-materials [GET]
//	@Security Bearer
func (ctrl *V1Controller) HandleBillOfMaterialsExport() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		actor := services.UseUserCtx(r.Context())

		csv, err := ctrl.svc.Reporting.BillOfMaterialsTSV(r.Context(), actor.GroupID)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment; filename=bom.csv")
		_, err = w.Write(csv)
		return err
	}
}
