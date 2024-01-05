// Package v1 provides the API handlers for version 1 of the API.
package v1

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/core/services/reporting/eventbus"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/hay-kot/httpkit/errchain"
	"github.com/hay-kot/httpkit/server"
	"github.com/rs/zerolog/log"

	"github.com/olahol/melody"
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

func WithSecureCookies(secure bool) func(*V1Controller) {
	return func(ctrl *V1Controller) {
		ctrl.cookieSecure = secure
	}
}

type V1Controller struct {
	cookieSecure      bool
	repo              *repo.AllRepos
	svc               *services.AllServices
	maxUploadSize     int64
	isDemo            bool
	allowRegistration bool
	bus               *eventbus.EventBus
}

type (
	ReadyFunc func() bool

	Build struct {
		Version   string `json:"version"`
		Commit    string `json:"commit"`
		BuildTime string `json:"buildTime"`
	}

	APISummary struct {
		Healthy           bool     `json:"health"`
		Versions          []string `json:"versions"`
		Title             string   `json:"title"`
		Message           string   `json:"message"`
		Build             Build    `json:"build"`
		Demo              bool     `json:"demo"`
		AllowRegistration bool     `json:"allowRegistration"`
	}
)

func BaseURLFunc(prefix string) func(s string) string {
	return func(s string) string {
		return prefix + "/v1" + s
	}
}

func NewControllerV1(svc *services.AllServices, repos *repo.AllRepos, bus *eventbus.EventBus, options ...func(*V1Controller)) *V1Controller {
	ctrl := &V1Controller{
		repo:              repos,
		svc:               svc,
		allowRegistration: true,
		bus:               bus,
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
//	@Success 200 {object} APISummary
//	@Router  /v1/status [GET]
func (ctrl *V1Controller) HandleBase(ready ReadyFunc, build Build) errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return server.JSON(w, http.StatusOK, APISummary{
			Healthy:           ready(),
			Title:             "Homebox",
			Message:           "Track, Manage, and Organize your Things",
			Build:             build,
			Demo:              ctrl.isDemo,
			AllowRegistration: ctrl.allowRegistration,
		})
	}
}

// HandleCurrency godoc
//
// @Summary Currency
// @Tags    Base
// @Produce json
// @Success 200 {object} currencies.Currency
// @Router  /v1/currency [GET]
func (ctrl *V1Controller) HandleCurrency() errchain.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		return server.JSON(w, http.StatusOK, ctrl.svc.Currencies.Slice())
	}
}

func (ctrl *V1Controller) HandleCacheWS() errchain.HandlerFunc {
	m := melody.New()

	m.HandleConnect(func(s *melody.Session) {
		auth := services.NewContext(s.Request.Context())
		s.Set("gid", auth.GID)
	})

	factory := func(e string) func(data any) {
		return func(data any) {
			eventData, ok := data.(eventbus.GroupMutationEvent)
			if !ok {
				log.Log().Msgf("invalid event data: %v", data)
				return
			}

			jsonStr := fmt.Sprintf(`{"event": "%s"}`, e)

			_ = m.BroadcastFilter([]byte(jsonStr), func(s *melody.Session) bool {
				groupIDStr, ok := s.Get("gid")
				if !ok {
					return false
				}

				GID := groupIDStr.(uuid.UUID)
				return GID == eventData.GID
			})
		}
	}

	ctrl.bus.Subscribe(eventbus.EventLabelMutation, factory("label.mutation"))
	ctrl.bus.Subscribe(eventbus.EventLocationMutation, factory("location.mutation"))
	ctrl.bus.Subscribe(eventbus.EventItemMutation, factory("item.mutation"))

	return func(w http.ResponseWriter, r *http.Request) error {
		return m.HandleRequest(w, r)
	}
}
