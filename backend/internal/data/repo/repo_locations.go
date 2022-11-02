package repo

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent"
	"github.com/hay-kot/homebox/backend/internal/data/ent/group"
	"github.com/hay-kot/homebox/backend/internal/data/ent/item"
	"github.com/hay-kot/homebox/backend/internal/data/ent/location"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
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
		ParentID    uuid.UUID `json:"parentId" extensions:"x-nullable"`
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
		Parent *LocationSummary `json:"parent,omitempty"`
		LocationSummary
		Items    []ItemSummary     `json:"items"`
		Children []LocationSummary `json:"children"`
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
	var parent *LocationSummary
	if location.Edges.Parent != nil {
		p := mapLocationSummary(location.Edges.Parent)
		parent = &p
	}

	children := make([]LocationSummary, 0, len(location.Edges.Children))
	for _, c := range location.Edges.Children {
		children = append(children, mapLocationSummary(c))
	}

	return LocationOut{
		Parent:   parent,
		Children: children,
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

type LocationQuery struct {
	FilterChildren bool `json:"filterChildren"`
}

// GetALlWithCount returns all locations with item count field populated
func (r *LocationRepository) GetAll(ctx context.Context, GID uuid.UUID, filter LocationQuery) ([]LocationOutCount, error) {
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
					AND items.archived = false
			) as item_count
		FROM
			locations
		WHERE
			locations.group_locations = ? {{ FILTER_CHILDREN }}
		ORDER BY
			locations.name ASC
`

	if filter.FilterChildren {
		query = strings.Replace(query, "{{ FILTER_CHILDREN }}", "AND locations.location_children IS NULL", 1)
	} else {
		query = strings.Replace(query, "{{ FILTER_CHILDREN }}", "", 1)
	}

	rows, err := r.db.Sql().QueryContext(ctx, query, GID)
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
			iq.Where(item.Archived(false)).WithLabel()
		}).
		WithParent().
		WithChildren().
		Only(ctx))
}

func (r *LocationRepository) Get(ctx context.Context, ID uuid.UUID) (LocationOut, error) {
	return r.getOne(ctx, location.ID(ID))
}

func (r *LocationRepository) GetOneByGroup(ctx context.Context, GID, ID uuid.UUID) (LocationOut, error) {
	return r.getOne(ctx, location.ID(ID), location.HasGroupWith(group.ID(GID)))
}

func (r *LocationRepository) Create(ctx context.Context, GID uuid.UUID, data LocationCreate) (LocationOut, error) {
	location, err := r.db.Location.Create().
		SetName(data.Name).
		SetDescription(data.Description).
		SetGroupID(GID).
		Save(ctx)

	if err != nil {
		return LocationOut{}, err
	}

	location.Edges.Group = &ent.Group{ID: GID} // bootstrap group ID
	return mapLocationOut(location), nil
}

func (r *LocationRepository) update(ctx context.Context, data LocationUpdate, where ...predicate.Location) (LocationOut, error) {
	q := r.db.Location.Update().
		Where(where...).
		SetName(data.Name).
		SetDescription(data.Description)

	if data.ParentID != uuid.Nil {
		q.SetParentID(data.ParentID)
	} else {
		q.ClearParent()
	}

	_, err := q.Save(ctx)
	if err != nil {
		return LocationOut{}, err
	}

	return r.Get(ctx, data.ID)
}

func (r *LocationRepository) Update(ctx context.Context, data LocationUpdate) (LocationOut, error) {
	return r.update(ctx, data, location.ID(data.ID))
}

func (r *LocationRepository) UpdateOneByGroup(ctx context.Context, GID, ID uuid.UUID, data LocationUpdate) (LocationOut, error) {
	return r.update(ctx, data, location.ID(ID), location.HasGroupWith(group.ID(GID)))
}

func (r *LocationRepository) Delete(ctx context.Context, ID uuid.UUID) error {
	return r.db.Location.DeleteOneID(ID).Exec(ctx)
}

func (r *LocationRepository) DeleteByGroup(ctx context.Context, GID, ID uuid.UUID) error {
	_, err := r.db.Location.Delete().Where(location.ID(ID), location.HasGroupWith(group.ID(GID))).Exec(ctx)
	return err
}
