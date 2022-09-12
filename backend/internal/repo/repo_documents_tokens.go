package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/ent/documenttoken"
	"github.com/hay-kot/content/backend/internal/types"
)

// DocumentTokensRepository is a repository for Document entity
type DocumentTokensRepository struct {
	db *ent.Client
}

func (r *DocumentTokensRepository) Create(ctx context.Context, data types.DocumentTokenCreate) (*ent.DocumentToken, error) {
	result, err := r.db.DocumentToken.Create().
		SetDocumentID(data.DocumentID).
		SetToken(data.TokenHash).
		SetExpiresAt(data.ExpiresAt).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return r.db.DocumentToken.Query().
		Where(documenttoken.ID(result.ID)).
		WithDocument().
		Only(ctx)
}

func (r *DocumentTokensRepository) PurgeExpiredTokens(ctx context.Context) (int, error) {
	return r.db.DocumentToken.Delete().Where(documenttoken.ExpiresAtLT(time.Now())).Exec(ctx)
}

func (r *DocumentTokensRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.DocumentToken.DeleteOneID(id).Exec(ctx)
}
