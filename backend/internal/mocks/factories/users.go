package factories

import (
	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/pkgs/faker"
)

func UserFactory() repo.UserCreate {
	f := faker.NewFaker()
	return repo.UserCreate{
		Name:        f.Str(10),
		Email:       f.Email(),
		Password:    f.Str(10),
		IsSuperuser: f.Bool(),
	}
}
