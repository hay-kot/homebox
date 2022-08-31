package base

import (
	"net/http"

	"github.com/hay-kot/content/backend/internal/types"
	"github.com/hay-kot/content/backend/pkgs/logger"
	"github.com/hay-kot/content/backend/pkgs/server"
)

type ReadyFunc func() bool

type BaseController struct {
	log *logger.Logger
	svr *server.Server
}

func NewBaseController(log *logger.Logger, svr *server.Server) *BaseController {
	h := &BaseController{
		log: log,
		svr: svr,
	}
	return h
}

// HandleBase godoc
// @Summary  Retrieves the basic information about the API
// @Tags     Base
// @Produce  json
// @Success  200  {object}  server.Result{item=types.ApiSummary}
// @Router   /status [GET]
func (ctrl *BaseController) HandleBase(ready ReadyFunc, versions ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := types.ApiSummary{
			Healthy:  ready(),
			Versions: versions,
			Title:    "Go API Template",
			Message:  "Welcome to the Go API Template Application!",
		}

		server.Respond(w, http.StatusOK, server.Wrap(data))
	}
}
