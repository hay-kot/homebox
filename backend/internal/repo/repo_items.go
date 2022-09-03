package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/ent/group"
	"github.com/hay-kot/content/backend/ent/item"
	"github.com/hay-kot/content/backend/internal/types"
)

type ItemsRepository struct {
	db *ent.Client
}

func (e *ItemsRepository) GetOne(ctx context.Context, id uuid.UUID) (*ent.Item, error) {
	return e.db.Item.Query().
		Where(item.ID(id)).
		WithFields().
		WithLabel().
		WithLocation().
		WithGroup().
		Only(ctx)
}

func (e *ItemsRepository) GetAll(ctx context.Context, gid uuid.UUID) ([]*ent.Item, error) {
	return e.db.Item.Query().
		Where(item.HasGroupWith(group.ID(gid))).
		WithLabel().
		WithLocation().
		All(ctx)
}

func (e *ItemsRepository) Create(ctx context.Context, gid uuid.UUID, data types.ItemCreate) (*ent.Item, error) {
	return e.db.Item.Create().
		SetName(data.Name).
		SetDescription(data.Description).
		SetGroupID(gid).
		AddLabelIDs(data.LabelIDs...).
		SetLocationID(data.LocationID).
		Save(ctx)
}

func (e *ItemsRepository) Delete(ctx context.Context, gid uuid.UUID, id uuid.UUID) error {
	panic("implement me")
}

func (e *ItemsRepository) Update(ctx context.Context, gid uuid.UUID, data types.ItemUpdate) (*ent.Item, error) {
	panic("implement me")
}
