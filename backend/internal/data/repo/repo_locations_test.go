package repo

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent"
	"github.com/stretchr/testify/assert"
)

func locationFactory() LocationCreate {
	return LocationCreate{
		Name:        fk.Str(10),
		Description: fk.Str(100),
	}
}

func useLocations(t *testing.T, len int) []LocationOut {
	t.Helper()

	out := make([]LocationOut, len)

	for i := 0; i < len; i++ {
		loc, err := tRepos.Locations.Create(context.Background(), tGroup.ID, locationFactory())
		assert.NoError(t, err)
		out[i] = loc
	}

	t.Cleanup(func() {
		for _, loc := range out {
			err := tRepos.Locations.Delete(context.Background(), loc.ID)

			if err != nil {
				assert.True(t, ent.IsNotFound(err))
			}
		}
	})

	return out
}

func TestLocationRepository_Get(t *testing.T) {
	loc, err := tRepos.Locations.Create(context.Background(), tGroup.ID, locationFactory())
	assert.NoError(t, err)

	// Get by ID
	foundLoc, err := tRepos.Locations.Get(context.Background(), loc.ID)
	assert.NoError(t, err)
	assert.Equal(t, loc.ID, foundLoc.ID)

	err = tRepos.Locations.Delete(context.Background(), loc.ID)
	assert.NoError(t, err)
}

func TestLocationRepositoryGetAllWithCount(t *testing.T) {
	ctx := context.Background()
	result := useLocations(t, 1)[0]

	_, err := tRepos.Items.Create(ctx, tGroup.ID, ItemCreate{
		Name:        fk.Str(10),
		Description: fk.Str(100),
		LocationID:  result.ID,
	})

	assert.NoError(t, err)

	results, err := tRepos.Locations.GetAll(context.Background(), tGroup.ID, LocationQuery{})
	assert.NoError(t, err)

	for _, loc := range results {
		if loc.ID == result.ID {
			assert.Equal(t, 1, loc.ItemCount)
		}
	}

}

func TestLocationRepository_Create(t *testing.T) {
	loc := useLocations(t, 1)[0]

	// Get by ID
	foundLoc, err := tRepos.Locations.Get(context.Background(), loc.ID)
	assert.NoError(t, err)
	assert.Equal(t, loc.ID, foundLoc.ID)

	err = tRepos.Locations.Delete(context.Background(), loc.ID)
	assert.NoError(t, err)
}

func TestLocationRepository_Update(t *testing.T) {
	loc := useLocations(t, 1)[0]

	updateData := LocationUpdate{
		ID:          loc.ID,
		Name:        fk.Str(10),
		Description: fk.Str(100),
	}

	update, err := tRepos.Locations.Update(context.Background(), updateData)
	assert.NoError(t, err)

	foundLoc, err := tRepos.Locations.Get(context.Background(), loc.ID)
	assert.NoError(t, err)

	assert.Equal(t, update.ID, foundLoc.ID)
	assert.Equal(t, update.Name, foundLoc.Name)
	assert.Equal(t, update.Description, foundLoc.Description)

	err = tRepos.Locations.Delete(context.Background(), loc.ID)
	assert.NoError(t, err)
}

func TestLocationRepository_Delete(t *testing.T) {
	loc := useLocations(t, 1)[0]

	err := tRepos.Locations.Delete(context.Background(), loc.ID)
	assert.NoError(t, err)

	_, err = tRepos.Locations.Get(context.Background(), loc.ID)
	assert.Error(t, err)
}

func TestItemRepository_TreeQuery(t *testing.T) {
	locs := useLocations(t, 3)

	// Set relations
	_, err := tRepos.Locations.UpdateOneByGroup(context.Background(), tGroup.ID, locs[0].ID, LocationUpdate{
		ID:          locs[0].ID,
		ParentID:    locs[1].ID,
		Name:        locs[0].Name,
		Description: locs[0].Description,
	})
	assert.NoError(t, err)

	locations, err := tRepos.Locations.Tree(context.Background(), tGroup.ID, TreeQuery{WithItems: true})

	assert.NoError(t, err)

	assert.Equal(t, 2, len(locations))

	// Check roots
	for _, loc := range locations {
		if loc.ID == locs[1].ID {
			assert.Equal(t, 1, len(loc.Children))
		}
	}
}

func TestConvertLocationsToTree(t *testing.T) {
	uuid1, uuid2, uuid3, uuid4 := uuid.New(), uuid.New(), uuid.New(), uuid.New()

	testCases := []struct {
		name      string
		locations []FlatTreeItem
		expected  []TreeItem
	}{
		{
			name: "Convert locations to tree",
			locations: []FlatTreeItem{
				{
					ID:       uuid1,
					Name:     "Root1",
					ParentID: uuid.Nil,
					Level:    0,
				},
				{
					ID:       uuid2,
					Name:     "Child1",
					ParentID: uuid1,
					Level:    1,
				},
				{
					ID:       uuid3,
					Name:     "Child2",
					ParentID: uuid1,
					Level:    1,
				},
			},
			expected: []TreeItem{
				{
					ID:   uuid1,
					Name: "Root1",
					Children: []*TreeItem{
						{
							ID:       uuid2,
							Name:     "Child1",
							Children: []*TreeItem{},
						},
						{
							ID:       uuid3,
							Name:     "Child2",
							Children: []*TreeItem{},
						},
					},
				},
			},
		},
		{
			name: "Convert locations to tree with deeply nested children",
			locations: []FlatTreeItem{
				{
					ID:       uuid1,
					Name:     "Root1",
					ParentID: uuid.Nil,
					Level:    0,
				},
				{
					ID:       uuid2,
					Name:     "Child1",
					ParentID: uuid1,
					Level:    1,
				},
				{
					ID:       uuid3,
					Name:     "Child2",
					ParentID: uuid2,
					Level:    2,
				},
				{
					ID:       uuid4,
					Name:     "Child3",
					ParentID: uuid3,
					Level:    3,
				},
			},
			expected: []TreeItem{
				{
					ID:   uuid1,
					Name: "Root1",
					Children: []*TreeItem{
						{
							ID:   uuid2,
							Name: "Child1",
							Children: []*TreeItem{
								{
									ID:   uuid3,
									Name: "Child2",
									Children: []*TreeItem{
										{
											ID:       uuid4,
											Name:     "Child3",
											Children: []*TreeItem{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ConvertLocationsToTree(tc.locations)

			// Compare JSON strings
			expected, _ := json.Marshal(tc.expected)
			got, _ := json.Marshal(result)
			assert.Equal(t, string(expected), string(got))
		})
	}
}
