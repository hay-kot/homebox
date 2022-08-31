package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/ent/group"
	"github.com/hay-kot/content/backend/ent/location"
	"github.com/hay-kot/content/backend/internal/types"
)

type EntLocationRepository struct {
	db *ent.Client
}

func (r *EntLocationRepository) Get(ctx context.Context, ID uuid.UUID) (*ent.Location, error) {
	return r.db.Location.Get(ctx, ID)
}

func (r *EntLocationRepository) GetAll(ctx context.Context, groupId uuid.UUID) ([]*ent.Location, error) {
	return r.db.Location.Query().
		Where(location.HasGroupWith(group.ID(groupId))).
		All(ctx)
}

func (r *EntLocationRepository) Create(ctx context.Context, groupdId uuid.UUID, data types.LocationCreate) (*ent.Location, error) {
	location, err := r.db.Location.Create().
		SetName(data.Name).
		SetDescription(data.Description).
		SetGroupID(groupdId).
		Save(ctx)

	location.Edges.Group = &ent.Group{ID: groupdId} // bootstrap group ID
	return location, err
}

func (r *EntLocationRepository) Update(ctx context.Context, data types.LocationUpdate) (*ent.Location, error) {
	return r.db.Location.UpdateOneID(data.ID).
		SetName(data.Name).
		SetDescription(data.Description).
		Save(ctx)
}

func (r *EntLocationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Location.DeleteOneID(id).Exec(ctx)
}
