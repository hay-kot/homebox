package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/ent"
)

type GroupRepository struct {
	db *ent.Client
}

func (r *GroupRepository) Create(ctx context.Context, name string) (*ent.Group, error) {
	return r.db.Group.Create().
		SetName(name).
		Save(ctx)
}

func (r *GroupRepository) GetOneId(ctx context.Context, id uuid.UUID) (*ent.Group, error) {
	return r.db.Group.Get(ctx, id)
}
