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
	q := e.db.Item.Create().
		SetName(data.Name).
		SetDescription(data.Description).
		SetGroupID(gid).
		SetLocationID(data.LocationID)

	if data.LabelIDs != nil && len(data.LabelIDs) > 0 {
		q.AddLabelIDs(data.LabelIDs...)
	}

	result, err := q.Save(ctx)
	if err != nil {
		return nil, err
	}

	return e.GetOne(ctx, result.ID)
}

func (e *ItemsRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.db.Item.DeleteOneID(id).Exec(ctx)
}

func (e *ItemsRepository) Update(ctx context.Context, data types.ItemUpdate) (*ent.Item, error) {
	q := e.db.Item.UpdateOneID(data.ID).
		SetName(data.Name).
		SetDescription(data.Description).
		SetLocationID(data.LocationID).
		SetSerialNumber(data.SerialNumber).
		SetModelNumber(data.ModelNumber).
		SetManufacturer(data.Manufacturer).
		SetPurchaseTime(data.PurchaseTime).
		SetPurchaseFrom(data.PurchaseFrom).
		SetPurchasePrice(data.PurchasePrice).
		SetSoldTime(data.SoldTime).
		SetSoldTo(data.SoldTo).
		SetSoldPrice(data.SoldPrice).
		SetSoldNotes(data.SoldNotes).
		SetNotes(data.Notes)

	err := q.Exec(ctx)

	if err != nil {
		return nil, err
	}

	return e.GetOne(ctx, data.ID)
}
