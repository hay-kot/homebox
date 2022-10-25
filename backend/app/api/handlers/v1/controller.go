package v1

import (
	"net/http"

	"github.com/hay-kot/homebox/backend/internal/services"
	"github.com/hay-kot/homebox/backend/pkgs/server"
)

func WithMaxUploadSize(maxUploadSize int64) func(*V1Controller) {
	return func(ctrl *V1Controller) {
		ctrl.maxUploadSize = maxUploadSize
	}
}

func WithDemoStatus(demoStatus bool) func(*V1Controller) {
	return func(ctrl *V1Controller) {
		ctrl.isDemo = demoStatus
	}
}

func WithRegistration(allowRegistration bool) func(*V1Controller) {
	return func(ctrl *V1Controller) {
		ctrl.allowRegistration = allowRegistration
	}
}

type V1Controller struct {
	svc               *services.AllServices
	maxUploadSize     int64
	isDemo            bool
	allowRegistration bool
}

type (
	ReadyFunc func() bool

	Build struct {
		Version   string `json:"version"`
		Commit    string `json:"commit"`
		BuildTime string `json:"buildTime"`
	}

	ApiSummary struct {
		Healthy  bool     `json:"health"`
		Versions []string `json:"versions"`
		Title    string   `json:"title"`
		Message  string   `json:"message"`
		Build    Build    `json:"build"`
		Demo     bool     `json:"demo"`
	}
)

func BaseUrlFunc(prefix string) func(s string) string {
	return func(s string) string {
		return prefix + "/v1" + s
	}
}

func NewControllerV1(svc *services.AllServices, options ...func(*V1Controller)) *V1Controller {
	ctrl := &V1Controller{
		svc:               svc,
		allowRegistration: true,
	}

	for _, opt := range options {
		opt(ctrl)
	}

	return ctrl
}

// HandleBase godoc
// @Summary Retrieves the basic information about the API
// @Tags    Base
// @Produce json
// @Success 200 {object} ApiSummary
// @Router  /v1/status [GET]
func (ctrl *V1Controller) HandleBase(ready ReadyFunc, build Build) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.Respond(w, http.StatusOK, ApiSummary{
			Healthy: ready(),
			Title:   "Go API Template",
			Message: "Welcome to the Go API Template Application!",
			Build:   build,
			Demo:    ctrl.isDemo,
		})
	}
}
