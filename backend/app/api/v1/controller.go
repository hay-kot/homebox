package v1

import (
	"net/http"

	"github.com/hay-kot/homebox/backend/internal/services"
	"github.com/hay-kot/homebox/backend/internal/types"
	"github.com/hay-kot/homebox/backend/pkgs/server"
)

type V1Controller struct {
	svc *services.AllServices
}

func BaseUrlFunc(prefix string) func(s string) string {
	v1Base := prefix + "/v1"
	prefixFunc := func(s string) string {
		return v1Base + s
	}

	return prefixFunc
}

func NewControllerV1(svc *services.AllServices) *V1Controller {
	ctrl := &V1Controller{
		svc: svc,
	}

	return ctrl
}

type ReadyFunc func() bool

// HandleBase godoc
// @Summary  Retrieves the basic information about the API
// @Tags     Base
// @Produce  json
// @Success  200  {object}  types.ApiSummary
// @Router   /v1/status [GET]
func (ctrl *V1Controller) HandleBase(ready ReadyFunc, build types.Build) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.Respond(w, http.StatusOK, types.ApiSummary{
			Healthy: ready(),
			Title:   "Go API Template",
			Message: "Welcome to the Go API Template Application!",
			Build:   build,
		})
	}
}
