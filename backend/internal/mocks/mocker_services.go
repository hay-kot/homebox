package mocks

import (
	"github.com/hay-kot/content/backend/internal/repo"
	"github.com/hay-kot/content/backend/internal/services"
)

func GetMockServices(repos *repo.AllRepos) *services.AllServices {
	return services.NewServices(repos)
}
