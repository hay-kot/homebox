package repo

import (
	"context"

	"github.com/hay-kot/git-web-template/backend/internal/types"
)

type TokenRepository interface {
	// GetUserFromToken get's a user from a token
	GetUserFromToken(ctx context.Context, token []byte) (types.UserOut, error)
	// Creates a token for a user
	CreateToken(ctx context.Context, createToken types.UserAuthTokenCreate) (types.UserAuthToken, error)
	// DeleteToken remove a single token from the database - equivalent to revoke or logout
	DeleteToken(ctx context.Context, token []byte) error
	// PurgeExpiredTokens removes all expired tokens from the database
	PurgeExpiredTokens(ctx context.Context) (int, error)
	// DeleteAll removes all tokens from the database
	DeleteAll(ctx context.Context) (int, error)
}
