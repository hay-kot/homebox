package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/ent/document"
	"github.com/hay-kot/homebox/backend/ent/group"
	"github.com/hay-kot/homebox/backend/internal/types"
)

// DocumentRepository is a repository for Document entity
type DocumentRepository struct {
	db *ent.Client
}

func (r *DocumentRepository) Create(ctx context.Context, gid uuid.UUID, doc types.DocumentCreate) (*ent.Document, error) {
	return r.db.Document.Create().
		SetGroupID(gid).
		SetTitle(doc.Title).
		SetPath(doc.Path).
		Save(ctx)
}

func (r *DocumentRepository) GetAll(ctx context.Context, gid uuid.UUID) ([]*ent.Document, error) {
	return r.db.Document.Query().
		Where(document.HasGroupWith(group.ID(gid))).
		All(ctx)
}

func (r *DocumentRepository) Get(ctx context.Context, id uuid.UUID) (*ent.Document, error) {
	return r.db.Document.Query().
		Where(document.ID(id)).
		Only(ctx)
}

func (r *DocumentRepository) Update(ctx context.Context, id uuid.UUID, doc types.DocumentUpdate) (*ent.Document, error) {
	return r.db.Document.UpdateOneID(id).
		SetTitle(doc.Title).
		SetPath(doc.Path).
		Save(ctx)
}

func (r *DocumentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Document.DeleteOneID(id).Exec(ctx)
}
