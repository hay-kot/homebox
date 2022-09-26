package v1

import (
	"context"
	"testing"

	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/internal/mocks/factories"
	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/internal/services"
	"github.com/hay-kot/homebox/backend/internal/types"
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
	c, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}

	if err := c.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	repos := repo.EntAllRepos(c, "/tmp/homebox")
	mockHandler.svc = services.NewServices(repos)

	purge := userPool()
	defer purge()

	m.Run()
}
