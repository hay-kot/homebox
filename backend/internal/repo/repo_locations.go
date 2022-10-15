package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/ent/group"
	"github.com/hay-kot/homebox/backend/ent/location"
	"github.com/hay-kot/homebox/backend/ent/predicate"
)

type LocationRepository struct {
	db *ent.Client
}

type (
	LocationCreate struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	LocationUpdate struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
	}

	LocationSummary struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}

	LocationOutCount struct {
		LocationSummary
		ItemCount int `json:"itemCount"`
	}

	LocationOut struct {
		LocationSummary
		Items []ItemSummary `json:"items"`
	}
)

func mapLocationSummary(location *ent.Location) LocationSummary {
	return LocationSummary{
		ID:          location.ID,
		Name:        location.Name,
		Description: location.Description,
		CreatedAt:   location.CreatedAt,
		UpdatedAt:   location.UpdatedAt,
	}
}

var (
	mapLocationOutErr = mapTErrFunc(mapLocationOut)
)

func mapLocationOut(location *ent.Location) LocationOut {
	return LocationOut{
		LocationSummary: LocationSummary{
			ID:          location.ID,
			Name:        location.Name,
			Description: location.Description,
			CreatedAt:   location.CreatedAt,
			UpdatedAt:   location.UpdatedAt,
		},
		Items: mapEach(location.Edges.Items, mapItemSummary),
	}
}

// GetALlWithCount returns all locations with item count field populated
func (r *LocationRepository) GetAll(ctx context.Context, groupId uuid.UUID) ([]LocationOutCount, error) {
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
		ORDER BY
			locations.name ASC
`

	rows, err := r.db.Sql().QueryContext(ctx, query, groupId)
	if err != nil {
		return nil, err
	}

	list := []LocationOutCount{}
	for rows.Next() {
		var ct LocationOutCount

		err := rows.Scan(&ct.ID, &ct.Name, &ct.Description, &ct.CreatedAt, &ct.UpdatedAt, &ct.ItemCount)
		if err != nil {
			return nil, err
		}

		list = append(list, ct)
	}

	return list, err
}

func (r *LocationRepository) getOne(ctx context.Context, where ...predicate.Location) (LocationOut, error) {
	return mapLocationOutErr(r.db.Location.Query().
		Where(where...).
		WithGroup().
		WithItems(func(iq *ent.ItemQuery) {
			iq.WithLabel()
		}).
		Only(ctx))
}

func (r *LocationRepository) Get(ctx context.Context, ID uuid.UUID) (LocationOut, error) {
	return r.getOne(ctx, location.ID(ID))
}

func (r *LocationRepository) GetOneByGroup(ctx context.Context, GID, ID uuid.UUID) (LocationOut, error) {
	return r.getOne(ctx, location.ID(ID), location.HasGroupWith(group.ID(GID)))
}

func (r *LocationRepository) Create(ctx context.Context, gid uuid.UUID, data LocationCreate) (LocationOut, error) {
	location, err := r.db.Location.Create().
		SetName(data.Name).
		SetDescription(data.Description).
		SetGroupID(gid).
		Save(ctx)

	if err != nil {
		return LocationOut{}, err
	}

	location.Edges.Group = &ent.Group{ID: gid} // bootstrap group ID
	return mapLocationOut(location), nil
}

func (r *LocationRepository) Update(ctx context.Context, data LocationUpdate) (LocationOut, error) {
	_, err := r.db.Location.UpdateOneID(data.ID).
		SetName(data.Name).
		SetDescription(data.Description).
		Save(ctx)

	if err != nil {
		return LocationOut{}, err
	}

	return r.Get(ctx, data.ID)
}

func (r *LocationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Location.DeleteOneID(id).Exec(ctx)
}
