package providers

import (
	"errors"
	"net/http"

	"github.com/hay-kot/homebox/backend/internal/core/services"
	"github.com/hay-kot/homebox/backend/internal/sys/config"
	"github.com/hay-kot/homebox/backend/pkgs/ipcheck"
)

type ForwardAuthProvider struct {
	service    *services.UserService
	authConfig *config.AuthConfig
}

func NewForwardAuthProvider(service *services.UserService, authConfig *config.AuthConfig) *ForwardAuthProvider {
	return &ForwardAuthProvider{
		service:    service,
		authConfig: authConfig,
	}
}

func (p *ForwardAuthProvider) Name() string {
	return "forwardauth"
}

func (p *ForwardAuthProvider) Authenticate(w http.ResponseWriter, r *http.Request) (services.UserAuthTokenDetail, error) {
	if !ipcheck.ValidateAgainstList(r.RemoteAddr, p.authConfig.ForwardAuthAllowedIps) {
		return services.UserAuthTokenDetail{}, errors.New("forward authentication denied, IP address not allowed")
	}

	username := r.Header.Get(p.authConfig.ForwardAuthHeader)

	return p.service.PasswordlessLogin(r.Context(), username, p.authConfig.ForwardAuthAutoRegister)
}
