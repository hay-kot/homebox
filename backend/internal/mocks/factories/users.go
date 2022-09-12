package factories

import (
	"github.com/hay-kot/content/backend/internal/types"
	"github.com/hay-kot/content/backend/pkgs/faker"
)

func UserFactory() types.UserCreate {
	f := faker.NewFaker()
	return types.UserCreate{
		Name:        f.Str(10),
		Email:       f.Email(),
		Password:    f.Str(10),
		IsSuperuser: f.Bool(),
	}
}
