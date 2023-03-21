package v1

import (
	"net/http"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/safeserve/errchain"
)

// HandleBillOfMaterialsExport godoc
//
//	@Summary  Export Bill of Materials
//	@Tags     Reporting
//	@Produce  json
//	@Success 200 {string} string "text/csv"
//	@Router   /v1/reporting/bill-of-materials [GET]
//	@Security Bearer
func (ctrl *V1Controller) HandleBillOfMaterialsExport() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		actor := services.UseUserCtx(r.Context())

		csv, err := ctrl.svc.Items.ExportBillOfMaterialsTSV(r.Context(), actor.GroupID)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "text/tsv")
		w.Header().Set("Content-Disposition", "attachment; filename=bill-of-materials.tsv")
		_, err = w.Write(csv)
		return err
	}
}
