package mocks

import (
	"github.com/hay-kot/git-web-template/backend/internal/repo"
	"github.com/hay-kot/git-web-template/backend/internal/services"
)

func GetMockServices(repos *repo.AllRepos) *services.AllServices {
	return services.NewServices(repos)
}
