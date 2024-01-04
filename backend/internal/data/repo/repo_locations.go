package repo

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/core/services/reporting/eventbus"
	"github.com/hay-kot/homebox/backend/internal/data/ent"
	"github.com/hay-kot/homebox/backend/internal/data/ent/group"
	"github.com/hay-kot/homebox/backend/internal/data/ent/location"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
)

type LocationRepository struct {
	db  *ent.Client
	bus *eventbus.EventBus
}

type (
	LocationCreate struct {
		Name        string    `json:"name"`
		ParentID    uuid.UUID `json:"parentId"    extensions:"x-nullable"`
		Description string    `json:"description"`
	}

	LocationUpdate struct {
		ParentID    uuid.UUID `json:"parentId"    extensions:"x-nullable"`
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

var mapLocationOutErr = mapTErrFunc(mapLocationOut)

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
	}
}

func (r *LocationRepository) publishMutationEvent(GID uuid.UUID) {
	if r.bus != nil {
		r.bus.Publish(eventbus.EventLocationMutation, eventbus.GroupMutationEvent{GID: GID})
	}
}

type LocationQuery struct {
	FilterChildren bool `json:"filterChildren" schema:"filterChildren"`
}

// GetAll returns all locations with item count field populated
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
					SUM(items.quantity)
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
	defer func() { _ = rows.Close() }()

	list := []LocationOutCount{}
	for rows.Next() {
		var ct LocationOutCount

		var maybeCount *int

		err := rows.Scan(&ct.ID, &ct.Name, &ct.Description, &ct.CreatedAt, &ct.UpdatedAt, &maybeCount)
		if err != nil {
			return nil, err
		}

		if maybeCount != nil {
			ct.ItemCount = *maybeCount
		}

		list = append(list, ct)
	}

	return list, err
}

func (r *LocationRepository) getOne(ctx context.Context, where ...predicate.Location) (LocationOut, error) {
	return mapLocationOutErr(r.db.Location.Query().
		Where(where...).
		WithGroup().
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
	q := r.db.Location.Create().
		SetName(data.Name).
		SetDescription(data.Description).
		SetGroupID(GID)

	if data.ParentID != uuid.Nil {
		q.SetParentID(data.ParentID)
	}

	location, err := q.Save(ctx)
	if err != nil {
		return LocationOut{}, err
	}

	location.Edges.Group = &ent.Group{ID: GID} // bootstrap group ID
	r.publishMutationEvent(GID)
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

func (r *LocationRepository) UpdateByGroup(ctx context.Context, GID, ID uuid.UUID, data LocationUpdate) (LocationOut, error) {
	v, err := r.update(ctx, data, location.ID(ID), location.HasGroupWith(group.ID(GID)))
	if err != nil {
		return LocationOut{}, err
	}

	r.publishMutationEvent(GID)
	return v, err
}

// delete should only be used after checking that the location is owned by the
// group. Otherwise, use DeleteByGroup
func (r *LocationRepository) delete(ctx context.Context, ID uuid.UUID) error {
	return r.db.Location.DeleteOneID(ID).Exec(ctx)
}

func (r *LocationRepository) DeleteByGroup(ctx context.Context, GID, ID uuid.UUID) error {
	_, err := r.db.Location.Delete().Where(location.ID(ID), location.HasGroupWith(group.ID(GID))).Exec(ctx)
	if err != nil {
		return err
	}
	r.publishMutationEvent(GID)

	return err
}

type TreeItem struct {
	ID       uuid.UUID   `json:"id"`
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Children []*TreeItem `json:"children"`
}

type FlatTreeItem struct {
	ID       uuid.UUID
	Name     string
	Type     string
	ParentID uuid.UUID
	Level    int
}

type TreeQuery struct {
	WithItems bool `json:"withItems" schema:"withItems"`
}

type LocationPath struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (r *LocationRepository) PathForLoc(ctx context.Context, GID, locID uuid.UUID) ([]LocationPath, error) {
	query := `WITH RECURSIVE location_path AS (
		SELECT id, name, location_children
		FROM locations
		WHERE id = ? -- Replace ? with the ID of the item's location
		AND group_locations = ? -- Replace ? with the ID of the group

		UNION ALL

		SELECT loc.id, loc.name, loc.location_children
		FROM locations loc
		JOIN location_path lp ON loc.id = lp.location_children
	  )

	  SELECT id, name
	  FROM location_path`

	rows, err := r.db.Sql().QueryContext(ctx, query, locID, GID)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var locations []LocationPath

	for rows.Next() {
		var location LocationPath
		if err := rows.Scan(&location.ID, &location.Name); err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Reverse the order of the locations so that the root is last
	for i := len(locations)/2 - 1; i >= 0; i-- {
		opp := len(locations) - 1 - i
		locations[i], locations[opp] = locations[opp], locations[i]
	}

	return locations, nil
}

func (r *LocationRepository) Tree(ctx context.Context, GID uuid.UUID, tq TreeQuery) ([]TreeItem, error) {
	query := `
		WITH recursive location_tree(id, NAME, parent_id, level, node_type) AS
		(
			SELECT  id,
					NAME,
					location_children AS parent_id,
					0 AS level,
					'location' AS node_type
			FROM    locations
			WHERE   location_children IS NULL
			AND     group_locations = ?

			UNION ALL
			SELECT  c.id,
					c.NAME,
					c.location_children AS parent_id,
					level + 1,
					'location' AS node_type
			FROM   locations c
			JOIN   location_tree p
			ON     c.location_children = p.id
			WHERE  level < 10 -- prevent infinite loop & excessive recursion
		){{ WITH_ITEMS }}

		SELECT   id,
				 NAME,
				 level,
				 parent_id,
				 node_type
		FROM    (
					SELECT  *
					FROM    location_tree


					{{ WITH_ITEMS_FROM }}


				) tree
		ORDER BY node_type DESC, -- sort locations before items
				 level,
				 lower(NAME)`

	if tq.WithItems {
		itemQuery := `, item_tree(id, NAME, parent_id, level, node_type) AS
		(
			SELECT  id,
					NAME,
					location_items as parent_id,
					0 AS level,
					'item' AS node_type
			FROM    items
			WHERE   item_children IS NULL
			AND     location_items IN (SELECT id FROM location_tree)

			UNION ALL

			SELECT  c.id,
					c.NAME,
					c.item_children AS parent_id,
					level + 1,
					'item' AS node_type
			FROM    items c
			JOIN    item_tree p
			ON      c.item_children = p.id
			WHERE   c.item_children IS NOT NULL
			AND     level < 10 -- prevent infinite loop & excessive recursion
		)`

		// Conditional table joined to main query
		itemsFrom := `
		UNION ALL
		SELECT  *
		FROM    item_tree`

		query = strings.ReplaceAll(query, "{{ WITH_ITEMS }}", itemQuery)
		query = strings.ReplaceAll(query, "{{ WITH_ITEMS_FROM }}", itemsFrom)
	} else {
		query = strings.ReplaceAll(query, "{{ WITH_ITEMS }}", "")
		query = strings.ReplaceAll(query, "{{ WITH_ITEMS_FROM }}", "")
	}

	rows, err := r.db.Sql().QueryContext(ctx, query, GID)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var locations []FlatTreeItem
	for rows.Next() {
		var location FlatTreeItem
		if err := rows.Scan(&location.ID, &location.Name, &location.Level, &location.ParentID, &location.Type); err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ConvertLocationsToTree(locations), nil
}

func ConvertLocationsToTree(locations []FlatTreeItem) []TreeItem {
	locationMap := make(map[uuid.UUID]*TreeItem, len(locations))

	var rootIds []uuid.UUID

	for _, location := range locations {
		loc := &TreeItem{
			ID:       location.ID,
			Name:     location.Name,
			Type:     location.Type,
			Children: []*TreeItem{},
		}

		locationMap[location.ID] = loc
		if location.ParentID != uuid.Nil {
			parent, ok := locationMap[location.ParentID]
			if ok {
				parent.Children = append(parent.Children, loc)
			}
		} else {
			rootIds = append(rootIds, location.ID)
		}
	}

	roots := make([]TreeItem, 0, len(rootIds))
	for _, id := range rootIds {
		roots = append(roots, *locationMap[id])
	}

	return roots
}
