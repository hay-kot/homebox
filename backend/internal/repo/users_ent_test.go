package repo

import (
	"context"
	"fmt"
	"testing"

	"github.com/hay-kot/git-web-template/backend/internal/types"
	"github.com/hay-kot/git-web-template/backend/pkgs/faker"
	"github.com/stretchr/testify/assert"
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

func Test_EntUserRepo_GetOneEmail(t *testing.T) {
	assert := assert.New(t)
	user := UserFactory()
	ctx := context.Background()

	testRepos.Users.Create(ctx, user)

	foundUser, err := testRepos.Users.GetOneEmail(ctx, user.Email)

	assert.NotNil(foundUser)
	assert.Nil(err)
	assert.Equal(user.Email, foundUser.Email)
	assert.Equal(user.Name, foundUser.Name)

	// Cleanup
	testRepos.Users.DeleteAll(ctx)
}

func Test_EntUserRepo_GetOneId(t *testing.T) {
	assert := assert.New(t)
	user := UserFactory()
	ctx := context.Background()

	userOut, _ := testRepos.Users.Create(ctx, user)
	foundUser, err := testRepos.Users.GetOneId(ctx, userOut.ID)

	assert.NotNil(foundUser)
	assert.Nil(err)
	assert.Equal(user.Email, foundUser.Email)
	assert.Equal(user.Name, foundUser.Name)

	// Cleanup
	testRepos.Users.DeleteAll(ctx)
}

func Test_EntUserRepo_GetAll(t *testing.T) {
	// Setup
	toCreate := []types.UserCreate{
		UserFactory(),
		UserFactory(),
		UserFactory(),
		UserFactory(),
	}

	ctx := context.Background()

	created := []types.UserOut{}

	for _, usr := range toCreate {
		usrOut, _ := testRepos.Users.Create(ctx, usr)
		created = append(created, usrOut)
	}

	// Validate
	allUsers, err := testRepos.Users.GetAll(ctx)

	assert.Nil(t, err)
	assert.Equal(t, len(created), len(allUsers))

	for _, usr := range created {
		fmt.Printf("%+v\n", usr)
		assert.Contains(t, allUsers, usr)
	}

	for _, usr := range created {
		testRepos.Users.Delete(ctx, usr.ID)
	}

	// Cleanup
	testRepos.Users.DeleteAll(ctx)
}

func Test_EntUserRepo_Update(t *testing.T) {
	t.Skip()
}

func Test_EntUserRepo_Delete(t *testing.T) {
	// Create 10 Users
	for i := 0; i < 10; i++ {
		user := UserFactory()
		ctx := context.Background()
		_, _ = testRepos.Users.Create(ctx, user)
	}

	// Delete all
	ctx := context.Background()
	allUsers, _ := testRepos.Users.GetAll(ctx)

	assert.Greater(t, len(allUsers), 0)
	testRepos.Users.DeleteAll(ctx)

	allUsers, _ = testRepos.Users.GetAll(ctx)
	assert.Equal(t, len(allUsers), 0)

}

func Test_EntUserRepo_GetSuperusers(t *testing.T) {
	// Create 10 Users
	superuser := 0
	users := 0

	for i := 0; i < 10; i++ {
		user := UserFactory()
		ctx := context.Background()
		_, _ = testRepos.Users.Create(ctx, user)

		if user.IsSuperuser {
			superuser++
		} else {
			users++
		}
	}

	// Delete all
	ctx := context.Background()

	superUsers, err := testRepos.Users.GetSuperusers(ctx)
	assert.NoError(t, err)

	for _, usr := range superUsers {
		assert.True(t, usr.IsSuperuser)
	}

	// Cleanup
	testRepos.Users.DeleteAll(ctx)
}
