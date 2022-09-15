package repo

import (
	"context"
	"time"

	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/ent/authtokens"
	"github.com/hay-kot/homebox/backend/internal/types"
)

type TokenRepository struct {
	db *ent.Client
}

// GetUserFromToken get's a user from a token
func (r *TokenRepository) GetUserFromToken(ctx context.Context, token []byte) (*ent.User, error) {
	user, err := r.db.AuthTokens.Query().
		Where(authtokens.Token(token)).
		Where(authtokens.ExpiresAtGTE(time.Now())).
		WithUser().
		QueryUser().
		WithGroup().
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Creates a token for a user
func (r *TokenRepository) CreateToken(ctx context.Context, createToken types.UserAuthTokenCreate) (types.UserAuthToken, error) {
	tokenOut := types.UserAuthToken{}

	dbToken, err := r.db.AuthTokens.Create().
		SetToken(createToken.TokenHash).
		SetUserID(createToken.UserID).
		SetExpiresAt(createToken.ExpiresAt).
		Save(ctx)

	if err != nil {
		return tokenOut, err
	}

	tokenOut.TokenHash = dbToken.Token
	tokenOut.UserID = createToken.UserID
	tokenOut.CreatedAt = dbToken.CreatedAt
	tokenOut.ExpiresAt = dbToken.ExpiresAt

	return tokenOut, nil
}

// DeleteToken remove a single token from the database - equivalent to revoke or logout
func (r *TokenRepository) DeleteToken(ctx context.Context, token []byte) error {
	_, err := r.db.AuthTokens.Delete().Where(authtokens.Token(token)).Exec(ctx)
	return err
}

// PurgeExpiredTokens removes all expired tokens from the database
func (r *TokenRepository) PurgeExpiredTokens(ctx context.Context) (int, error) {
	tokensDeleted, err := r.db.AuthTokens.Delete().Where(authtokens.ExpiresAtLTE(time.Now())).Exec(ctx)

	if err != nil {
		return 0, err
	}

	return tokensDeleted, nil
}

func (r *TokenRepository) DeleteAll(ctx context.Context) (int, error) {
	amount, err := r.db.AuthTokens.Delete().Exec(ctx)
	return amount, err
}
