package repo

import (
	"context"
	"testing"

	"github.com/hay-kot/content/backend/internal/types"
	"github.com/stretchr/testify/assert"
)

func locationFactory() types.LocationCreate {
	return types.LocationCreate{
		Name:        fk.Str(10),
		Description: fk.Str(100),
	}
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
	result, err := tRepos.Locations.Create(ctx, tGroup.ID, types.LocationCreate{
		Name:        fk.Str(10),
		Description: fk.Str(100),
	})
	assert.NoError(t, err)

	_, err = tRepos.Items.Create(ctx, tGroup.ID, types.ItemCreate{
		Name:        fk.Str(10),
		Description: fk.Str(100),
		LocationID:  result.ID,
	})

	assert.NoError(t, err)

	results, err := tRepos.Locations.GetAll(context.Background(), tGroup.ID)
	assert.NoError(t, err)

	for _, loc := range results {
		if loc.ID == result.ID {
			assert.Equal(t, 1, loc.ItemCount)
		}
	}

}

func TestLocationRepository_Create(t *testing.T) {
	loc, err := tRepos.Locations.Create(context.Background(), tGroup.ID, locationFactory())
	assert.NoError(t, err)

	// Get by ID
	foundLoc, err := tRepos.Locations.Get(context.Background(), loc.ID)
	assert.NoError(t, err)
	assert.Equal(t, loc.ID, foundLoc.ID)

	err = tRepos.Locations.Delete(context.Background(), loc.ID)
	assert.NoError(t, err)
}

func TestLocationRepository_Update(t *testing.T) {
	loc, err := tRepos.Locations.Create(context.Background(), tGroup.ID, locationFactory())
	assert.NoError(t, err)

	updateData := types.LocationUpdate{
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
	loc, err := tRepos.Locations.Create(context.Background(), tGroup.ID, locationFactory())
	assert.NoError(t, err)

	err = tRepos.Locations.Delete(context.Background(), loc.ID)
	assert.NoError(t, err)

	_, err = tRepos.Locations.Get(context.Background(), loc.ID)
	assert.Error(t, err)
}
