package repo

import (
	"context"
	"testing"
	"time"

	"github.com/hay-kot/git-web-template/backend/internal/types"
	"github.com/hay-kot/git-web-template/backend/pkgs/hasher"
	"github.com/stretchr/testify/assert"
)

func Test_EntAuthTokenRepo_CreateToken(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	user := UserFactory()

	userOut, _ := testRepos.Users.Create(ctx, user)

	expiresAt := time.Now().Add(time.Hour)

	generatedToken := hasher.GenerateToken()

	token, err := testRepos.AuthTokens.CreateToken(ctx, types.UserAuthTokenCreate{
		TokenHash: generatedToken.Hash,
		ExpiresAt: expiresAt,
		UserID:    userOut.ID,
	})

	assert.NoError(err)
	assert.Equal(userOut.ID, token.UserID)
	assert.Equal(expiresAt, token.ExpiresAt)

	// Cleanup
	err = testRepos.Users.Delete(ctx, userOut.ID)
	_, err = testRepos.AuthTokens.DeleteAll(ctx)
}

func Test_EntAuthTokenRepo_GetUserByToken(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	user := UserFactory()
	userOut, _ := testRepos.Users.Create(ctx, user)

	expiresAt := time.Now().Add(time.Hour)
	generatedToken := hasher.GenerateToken()

	token, err := testRepos.AuthTokens.CreateToken(ctx, types.UserAuthTokenCreate{
		TokenHash: generatedToken.Hash,
		ExpiresAt: expiresAt,
		UserID:    userOut.ID,
	})

	// Get User from token
	foundUser, err := testRepos.AuthTokens.GetUserFromToken(ctx, token.TokenHash)

	assert.NoError(err)
	assert.Equal(userOut.ID, foundUser.ID)
	assert.Equal(userOut.Name, foundUser.Name)
	assert.Equal(userOut.Email, foundUser.Email)

	// Cleanup
	err = testRepos.Users.Delete(ctx, userOut.ID)
	_, err = testRepos.AuthTokens.DeleteAll(ctx)
}

func Test_EntAuthTokenRepo_PurgeExpiredTokens(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	user := UserFactory()
	userOut, _ := testRepos.Users.Create(ctx, user)

	createdTokens := []types.UserAuthToken{}

	for i := 0; i < 5; i++ {
		expiresAt := time.Now()
		generatedToken := hasher.GenerateToken()

		createdToken, err := testRepos.AuthTokens.CreateToken(ctx, types.UserAuthTokenCreate{
			TokenHash: generatedToken.Hash,
			ExpiresAt: expiresAt,
			UserID:    userOut.ID,
		})

		assert.NoError(err)
		assert.NotNil(createdToken)

		createdTokens = append(createdTokens, createdToken)

	}

	// Purge expired tokens
	tokensDeleted, err := testRepos.AuthTokens.PurgeExpiredTokens(ctx)

	assert.NoError(err)
	assert.Equal(5, tokensDeleted)

	// Check if tokens are deleted
	for _, token := range createdTokens {
		_, err := testRepos.AuthTokens.GetUserFromToken(ctx, token.TokenHash)
		assert.Error(err)
	}

	// Cleanup
	err = testRepos.Users.Delete(ctx, userOut.ID)
	_, err = testRepos.AuthTokens.DeleteAll(ctx)
}
