package v1

import (
	"github.com/hay-kot/content/backend/internal/services"
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
