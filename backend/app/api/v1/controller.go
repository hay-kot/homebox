package v1

import (
	"github.com/hay-kot/git-web-template/backend/internal/services"
	"github.com/hay-kot/git-web-template/backend/pkgs/logger"
)

type V1Controller struct {
	log *logger.Logger
	svc *services.AllServices
}

func BaseUrlFunc(prefix string) func(s string) string {
	v1Base := prefix + "/v1"
	prefixFunc := func(s string) string {
		return v1Base + s
	}

	return prefixFunc
}

func NewControllerV1(log *logger.Logger, svc *services.AllServices) *V1Controller {
	ctrl := &V1Controller{
		log: log,
		svc: svc,
	}

	return ctrl
}
