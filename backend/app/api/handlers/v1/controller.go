package v1

import (
	"net/http"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/safeserve/errchain"
	"github.com/hay-kot/safeserve/server"
)

type Results[T any] struct {
	Items []T `json:"items"`
}

func WrapResults[T any](items []T) Results[T] {
	return Results[T]{Items: items}
}

type Wrapped struct {
	Item interface{} `json:"item"`
}

func Wrap(v any) Wrapped {
	return Wrapped{Item: v}
}

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

func WithHeaderSSO(headerSSOEnabled bool) func(*V1Controller) {
	return func(ctrl *V1Controller) {
		ctrl.headerSSOEnabled = headerSSOEnabled
	}
}

func WithHeaderSSOAllowedIP(headerSSOAllowedIP string) func(*V1Controller) {
	return func(ctrl *V1Controller) {
		ctrl.headerSSOAllowedIP = headerSSOAllowedIP
	}
}

func WithHeaderSSOAutoRegister(headerSSOAutoRegister bool) func(*V1Controller) {
	return func(ctrl *V1Controller) {
		ctrl.headerSSOAutoRegister = headerSSOAutoRegister
	}
}

func WithHeaderSSOHeaderName(headerSSOHeaderName string) func(*V1Controller) {
	return func(ctrl *V1Controller) {
		ctrl.headerSSOHeaderName = headerSSOHeaderName
	}
}

func WithHeaderSSOHeaderEmail(headerSSOHeaderEmail string) func(*V1Controller) {
	return func(ctrl *V1Controller) {
		ctrl.headerSSOHeaderEmail = headerSSOHeaderEmail
	}
}

type V1Controller struct {
	repo                        *repo.AllRepos
	svc                         *services.AllServices
	maxUploadSize               int64
	isDemo                      bool
	allowRegistration           bool
	headerSSOEnabled            bool
	headerSSOAllowedIP          string
	headerSSOAutoRegister       bool
	headerSSOHeaderName			string
	headerSSOHeaderEmail		string
}

type (
	ReadyFunc func() bool

	Build struct {
		Version   string `json:"version"`
		Commit    string `json:"commit"`
		BuildTime string `json:"buildTime"`
	}

	ApiSummary struct {
		Healthy           bool     `json:"health"`
		Versions          []string `json:"versions"`
		Title             string   `json:"title"`
		Message           string   `json:"message"`
		Build             Build    `json:"build"`
		Demo              bool     `json:"demo"`
		AllowRegistration bool     `json:"allowRegistration"`
	}
)

func BaseUrlFunc(prefix string) func(s string) string {
	return func(s string) string {
		return prefix + "/v1" + s
	}
}

func NewControllerV1(svc *services.AllServices, repos *repo.AllRepos, options ...func(*V1Controller)) *V1Controller {
	ctrl := &V1Controller{
		repo:              repos,
		svc:               svc,
		allowRegistration: true,
	}

	for _, opt := range options {
		opt(ctrl)
	}

	return ctrl
}

// HandleBase godoc
//
//	@Summary Application Info
//	@Tags    Base
//	@Produce json
//	@Success 200 {object} ApiSummary
//	@Router  /v1/status [GET]
func (ctrl *V1Controller) HandleBase(ready ReadyFunc, build Build) errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return server.JSON(w, http.StatusOK, ApiSummary{
			Healthy:           ready(),
			Title:             "Homebox",
			Message:           "Track, Manage, and Organize your Things",
			Build:             build,
			Demo:              ctrl.isDemo,
			AllowRegistration: ctrl.allowRegistration,
		})
	}
}
