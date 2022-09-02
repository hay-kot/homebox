package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/ent/group"
	"github.com/hay-kot/content/backend/ent/label"
	"github.com/hay-kot/content/backend/internal/types"
)

type EntLabelRepository struct {
	db *ent.Client
}

func (r *EntLabelRepository) Get(ctx context.Context, ID uuid.UUID) (*ent.Label, error) {
	return r.db.Label.Query().
		Where(label.ID(ID)).
		WithGroup().
		WithItems().
		Only(ctx)
}

func (r *EntLabelRepository) GetAll(ctx context.Context, groupId uuid.UUID) ([]*ent.Label, error) {
	return r.db.Label.Query().
		Where(label.HasGroupWith(group.ID(groupId))).
		WithGroup().
		All(ctx)
}

func (r *EntLabelRepository) Create(ctx context.Context, groupdId uuid.UUID, data types.LabelCreate) (*ent.Label, error) {
	label, err := r.db.Label.Create().
		SetName(data.Name).
		SetDescription(data.Description).
		SetColor(data.Color).
		SetGroupID(groupdId).
		Save(ctx)

	label.Edges.Group = &ent.Group{ID: groupdId} // bootstrap group ID
	return label, err
}

func (r *EntLabelRepository) Update(ctx context.Context, data types.LabelUpdate) (*ent.Label, error) {
	_, err := r.db.Label.UpdateOneID(data.ID).
		SetName(data.Name).
		SetDescription(data.Description).
		SetColor(data.Color).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return r.Get(ctx, data.ID)
}

func (r *EntLabelRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Label.DeleteOneID(id).Exec(ctx)
}
