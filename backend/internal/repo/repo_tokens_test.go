package repo

import (
	"context"
	"testing"
	"time"

	"github.com/hay-kot/homebox/backend/pkgs/hasher"
	"github.com/stretchr/testify/assert"
)

func TestAuthTokenRepo_CreateToken(t *testing.T) {
	asrt := assert.New(t)
	ctx := context.Background()
	user := userFactory()

	userOut, err := tRepos.Users.Create(ctx, user)
	asrt.NoError(err)

	expiresAt := time.Now().Add(time.Hour)

	generatedToken := hasher.GenerateToken()

	token, err := tRepos.AuthTokens.CreateToken(ctx, UserAuthTokenCreate{
		TokenHash: generatedToken.Hash,
		ExpiresAt: expiresAt,
		UserID:    userOut.ID,
	})

	asrt.NoError(err)
	asrt.Equal(userOut.ID, token.UserID)
	asrt.Equal(expiresAt, token.ExpiresAt)

	// Cleanup
	asrt.NoError(tRepos.Users.Delete(ctx, userOut.ID))
	_, err = tRepos.AuthTokens.DeleteAll(ctx)
	asrt.NoError(err)
}

func TestAuthTokenRepo_DeleteToken(t *testing.T) {
	asrt := assert.New(t)
	ctx := context.Background()
	user := userFactory()

	userOut, err := tRepos.Users.Create(ctx, user)
	asrt.NoError(err)

	expiresAt := time.Now().Add(time.Hour)

	generatedToken := hasher.GenerateToken()

	_, err = tRepos.AuthTokens.CreateToken(ctx, UserAuthTokenCreate{
		TokenHash: generatedToken.Hash,
		ExpiresAt: expiresAt,
		UserID:    userOut.ID,
	})
	asrt.NoError(err)

	// Delete token
	err = tRepos.AuthTokens.DeleteToken(ctx, []byte(generatedToken.Raw))
	asrt.NoError(err)
}

func TestAuthTokenRepo_GetUserByToken(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	user := userFactory()
	userOut, _ := tRepos.Users.Create(ctx, user)

	expiresAt := time.Now().Add(time.Hour)
	generatedToken := hasher.GenerateToken()

	token, err := tRepos.AuthTokens.CreateToken(ctx, UserAuthTokenCreate{
		TokenHash: generatedToken.Hash,
		ExpiresAt: expiresAt,
		UserID:    userOut.ID,
	})

	assert.NoError(err)

	// Get User from token
	foundUser, err := tRepos.AuthTokens.GetUserFromToken(ctx, token.TokenHash)

	assert.NoError(err)
	assert.Equal(userOut.ID, foundUser.ID)
	assert.Equal(userOut.Name, foundUser.Name)
	assert.Equal(userOut.Email, foundUser.Email)

	// Cleanup
	assert.NoError(tRepos.Users.Delete(ctx, userOut.ID))
	_, err = tRepos.AuthTokens.DeleteAll(ctx)
	assert.NoError(err)
}

func TestAuthTokenRepo_PurgeExpiredTokens(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	user := userFactory()
	userOut, _ := tRepos.Users.Create(ctx, user)

	createdTokens := []UserAuthToken{}

	for i := 0; i < 5; i++ {
		expiresAt := time.Now()
		generatedToken := hasher.GenerateToken()

		createdToken, err := tRepos.AuthTokens.CreateToken(ctx, UserAuthTokenCreate{
			TokenHash: generatedToken.Hash,
			ExpiresAt: expiresAt,
			UserID:    userOut.ID,
		})

		assert.NoError(err)
		assert.NotNil(createdToken)

		createdTokens = append(createdTokens, createdToken)

	}

	// Purge expired tokens
	tokensDeleted, err := tRepos.AuthTokens.PurgeExpiredTokens(ctx)

	assert.NoError(err)
	assert.Equal(5, tokensDeleted)

	// Check if tokens are deleted
	for _, token := range createdTokens {
		_, err := tRepos.AuthTokens.GetUserFromToken(ctx, token.TokenHash)
		assert.Error(err)
	}

	// Cleanup
	assert.NoError(tRepos.Users.Delete(ctx, userOut.ID))
	_, err = tRepos.AuthTokens.DeleteAll(ctx)
	assert.NoError(err)
}
