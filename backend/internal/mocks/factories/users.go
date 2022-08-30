package factories

import (
	"github.com/hay-kot/git-web-template/backend/internal/types"
	"github.com/hay-kot/git-web-template/backend/pkgs/faker"
)

func UserFactory() types.UserCreate {
	f := faker.NewFaker()
	return types.UserCreate{
		Name:        f.RandomString(10),
		Email:       f.RandomEmail(),
		Password:    f.RandomString(10),
		IsSuperuser: f.RandomBool(),
	}
}
