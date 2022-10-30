package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent"
	"github.com/hay-kot/homebox/backend/internal/data/ent/documenttoken"
)

// DocumentTokensRepository is a repository for Document entity
type DocumentTokensRepository struct {
	db *ent.Client
}

type (
	DocumentToken struct {
		ID         uuid.UUID `json:"-"`
		TokenHash  []byte    `json:"tokenHash"`
		ExpiresAt  time.Time `json:"expiresAt"`
		DocumentID uuid.UUID `json:"documentId"`
	}

	DocumentTokenCreate struct {
		TokenHash  []byte    `json:"tokenHash"`
		DocumentID uuid.UUID `json:"documentId"`
		ExpiresAt  time.Time `json:"expiresAt"`
	}
)

var (
	mapDocumentTokenErr = mapTErrFunc(mapDocumentToken)
)

func mapDocumentToken(e *ent.DocumentToken) DocumentToken {
	return DocumentToken{
		ID:         e.ID,
		TokenHash:  e.Token,
		ExpiresAt:  e.ExpiresAt,
		DocumentID: e.Edges.Document.ID,
	}
}

func (r *DocumentTokensRepository) Create(ctx context.Context, data DocumentTokenCreate) (DocumentToken, error) {
	result, err := r.db.DocumentToken.Create().
		SetDocumentID(data.DocumentID).
		SetToken(data.TokenHash).
		SetExpiresAt(data.ExpiresAt).
		Save(ctx)

	if err != nil {
		return DocumentToken{}, err
	}

	return mapDocumentTokenErr(r.db.DocumentToken.Query().
		Where(documenttoken.ID(result.ID)).
		WithDocument().
		Only(ctx))
}

func (r *DocumentTokensRepository) PurgeExpiredTokens(ctx context.Context) (int, error) {
	return r.db.DocumentToken.Delete().Where(documenttoken.ExpiresAtLT(time.Now())).Exec(ctx)
}

func (r *DocumentTokensRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.DocumentToken.DeleteOneID(id).Exec(ctx)
}
