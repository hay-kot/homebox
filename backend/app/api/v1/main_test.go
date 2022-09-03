package v1

import (
	"context"
	"testing"

	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/internal/mocks"
	"github.com/hay-kot/content/backend/internal/mocks/factories"
	"github.com/hay-kot/content/backend/internal/types"
)

var mockHandler = &V1Controller{}
var users = []*ent.User{}

func userPool() func() {
	create := []types.UserCreate{
		factories.UserFactory(),
		factories.UserFactory(),
		factories.UserFactory(),
		factories.UserFactory(),
	}

	userOut := []*ent.User{}

	for _, user := range create {
		usrOut, _ := mockHandler.svc.Admin.Create(context.Background(), user)
		userOut = append(userOut, usrOut)
	}

	users = userOut

	purge := func() {
		_ = mockHandler.svc.Admin.DeleteAll(context.Background())
	}

	return purge
}

func TestMain(m *testing.M) {
	// Set Handler Vars
	repos, closeDb := mocks.GetEntRepos()
	mockHandler.svc = mocks.GetMockServices(repos)

	defer func() {
		_ = closeDb()
	}()

	purge := userPool()
	defer purge()

	m.Run()
}
