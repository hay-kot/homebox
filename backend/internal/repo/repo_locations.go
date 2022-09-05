package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/ent/location"
	"github.com/hay-kot/content/backend/internal/types"
)

type LocationRepository struct {
	db *ent.Client
}

type LocationWithCount struct {
	*ent.Location
	ItemCount int `json:"itemCount"`
}

// GetALlWithCount returns all locations with item count field populated
func (r *LocationRepository) GetAll(ctx context.Context, groupId uuid.UUID) ([]LocationWithCount, error) {
	query := `--sql
		SELECT
			id,
			name,
			description,
			created_at,
			updated_at,
			(
				SELECT
					COUNT(*)
				FROM
					items
				WHERE
					items.location_items = locations.id
			) as item_count
		FROM
			locations
		WHERE
			locations.group_locations = ?
	`

	rows, err := r.db.Sql().QueryContext(ctx, query, groupId)
	if err != nil {
		return nil, err
	}

	list := []LocationWithCount{}
	for rows.Next() {
		var loc ent.Location
		var ct LocationWithCount
		err := rows.Scan(&loc.ID, &loc.Name, &loc.Description, &loc.CreatedAt, &loc.UpdatedAt, &ct.ItemCount)
		if err != nil {
			return nil, err
		}
		ct.Location = &loc
		list = append(list, ct)
	}

	return list, err
}

func (r *LocationRepository) Get(ctx context.Context, ID uuid.UUID) (*ent.Location, error) {
	return r.db.Location.Query().
		Where(location.ID(ID)).
		WithGroup().
		WithItems(func(iq *ent.ItemQuery) {
			iq.WithLabel()
		}).
		Only(ctx)
}

func (r *LocationRepository) Create(ctx context.Context, groupdId uuid.UUID, data types.LocationCreate) (*ent.Location, error) {
	location, err := r.db.Location.Create().
		SetName(data.Name).
		SetDescription(data.Description).
		SetGroupID(groupdId).
		Save(ctx)

	location.Edges.Group = &ent.Group{ID: groupdId} // bootstrap group ID
	return location, err
}

func (r *LocationRepository) Update(ctx context.Context, data types.LocationUpdate) (*ent.Location, error) {
	_, err := r.db.Location.UpdateOneID(data.ID).
		SetName(data.Name).
		SetDescription(data.Description).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return r.Get(ctx, data.ID)
}

func (r *LocationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Location.DeleteOneID(id).Exec(ctx)
}
