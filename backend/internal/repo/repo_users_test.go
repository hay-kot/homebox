package repo

import (
	"context"
	"fmt"
	"testing"

	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/internal/types"
	"github.com/stretchr/testify/assert"
)

func userFactory() types.UserCreate {

	return types.UserCreate{
		Name:        fk.Str(10),
		Email:       fk.Email(),
		Password:    fk.Str(10),
		IsSuperuser: fk.Bool(),
		GroupID:     tGroup.ID,
	}
}

func TestUserRepo_GetOneEmail(t *testing.T) {
	assert := assert.New(t)
	user := userFactory()
	ctx := context.Background()

	_, err := tRepos.Users.Create(ctx, user)
	assert.NoError(err)

	foundUser, err := tRepos.Users.GetOneEmail(ctx, user.Email)

	assert.NotNil(foundUser)
	assert.Nil(err)
	assert.Equal(user.Email, foundUser.Email)
	assert.Equal(user.Name, foundUser.Name)

	// Cleanup
	err = tRepos.Users.DeleteAll(ctx)
	assert.NoError(err)
}

func TestUserRepo_GetOneId(t *testing.T) {
	assert := assert.New(t)
	user := userFactory()
	ctx := context.Background()

	userOut, _ := tRepos.Users.Create(ctx, user)
	foundUser, err := tRepos.Users.GetOneId(ctx, userOut.ID)

	assert.NotNil(foundUser)
	assert.Nil(err)
	assert.Equal(user.Email, foundUser.Email)
	assert.Equal(user.Name, foundUser.Name)

	// Cleanup
	err = tRepos.Users.DeleteAll(ctx)
	assert.NoError(err)
}

func TestUserRepo_GetAll(t *testing.T) {
	// Setup
	toCreate := []types.UserCreate{
		userFactory(),
		userFactory(),
		userFactory(),
		userFactory(),
	}

	ctx := context.Background()

	created := []*ent.User{}

	for _, usr := range toCreate {
		usrOut, _ := tRepos.Users.Create(ctx, usr)
		created = append(created, usrOut)
	}

	// Validate
	allUsers, err := tRepos.Users.GetAll(ctx)

	assert.NoError(t, err)
	assert.Equal(t, len(created), len(allUsers))

	for _, usr := range created {
		fmt.Printf("%+v\n", usr)
		for _, usr2 := range allUsers {
			if usr.ID == usr2.ID {
				assert.Equal(t, usr.Email, usr2.Email)

				// Check groups are loaded
				assert.NotNil(t, usr2.Edges.Group)
			}
		}
	}

	for _, usr := range created {
		_ = tRepos.Users.Delete(ctx, usr.ID)
	}

	// Cleanup
	err = tRepos.Users.DeleteAll(ctx)
	assert.NoError(t, err)
}

func TestUserRepo_Update(t *testing.T) {
	user, err := tRepos.Users.Create(context.Background(), userFactory())
	assert.NoError(t, err)

	updateData := types.UserUpdate{
		Name:  fk.Str(10),
		Email: fk.Email(),
	}

	// Update
	err = tRepos.Users.Update(context.Background(), user.ID, updateData)
	assert.NoError(t, err)

	// Validate
	updated, err := tRepos.Users.GetOneId(context.Background(), user.ID)
	assert.NoError(t, err)
	assert.NotEqual(t, user.Name, updated.Name)
	assert.NotEqual(t, user.Email, updated.Email)
}

func TestUserRepo_Delete(t *testing.T) {
	// Create 10 Users
	for i := 0; i < 10; i++ {
		user := userFactory()
		ctx := context.Background()
		_, _ = tRepos.Users.Create(ctx, user)
	}

	// Delete all
	ctx := context.Background()
	allUsers, _ := tRepos.Users.GetAll(ctx)

	assert.Greater(t, len(allUsers), 0)
	err := tRepos.Users.DeleteAll(ctx)
	assert.NoError(t, err)

	allUsers, _ = tRepos.Users.GetAll(ctx)
	assert.Equal(t, len(allUsers), 0)

}

func TestUserRepo_GetSuperusers(t *testing.T) {
	// Create 10 Users
	superuser := 0
	users := 0

	for i := 0; i < 10; i++ {
		user := userFactory()
		ctx := context.Background()
		_, _ = tRepos.Users.Create(ctx, user)

		if user.IsSuperuser {
			superuser++
		} else {
			users++
		}
	}

	// Delete all
	ctx := context.Background()

	superUsers, err := tRepos.Users.GetSuperusers(ctx)
	assert.NoError(t, err)

	for _, usr := range superUsers {
		assert.True(t, usr.IsSuperuser)
	}

	// Cleanup
	err = tRepos.Users.DeleteAll(ctx)
	assert.NoError(t, err)
}
