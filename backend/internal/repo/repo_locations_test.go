package repo

import (
	"context"
	"testing"

	"github.com/hay-kot/content/backend/internal/types"
	"github.com/hay-kot/content/backend/pkgs/faker"
	"github.com/stretchr/testify/assert"
)

var fk = faker.NewFaker()

func locationFactory() types.LocationCreate {
	return types.LocationCreate{
		Name:        fk.RandomString(10),
		Description: fk.RandomString(100),
	}
}

func Test_Locations_Get(t *testing.T) {
	loc, err := testRepos.Locations.Create(context.Background(), testGroup.ID, locationFactory())
	assert.NoError(t, err)

	// Get by ID
	foundLoc, err := testRepos.Locations.Get(context.Background(), loc.ID)
	assert.NoError(t, err)
	assert.Equal(t, loc.ID, foundLoc.ID)

	err = testRepos.Locations.Delete(context.Background(), loc.ID)
	assert.NoError(t, err)
}

func Test_LocationsGetAllWithCount(t *testing.T) {
	ctx := context.Background()
	result, err := testRepos.Locations.Create(ctx, testGroup.ID, types.LocationCreate{
		Name:        fk.RandomString(10),
		Description: fk.RandomString(100),
	})
	assert.NoError(t, err)

	_, err = testRepos.Items.Create(ctx, testGroup.ID, types.ItemCreate{
		Name:        fk.RandomString(10),
		Description: fk.RandomString(100),
		LocationID:  result.ID,
	})

	assert.NoError(t, err)

	results, err := testRepos.Locations.GetAll(context.Background(), testGroup.ID)
	assert.NoError(t, err)

	for _, loc := range results {
		if loc.ID == result.ID {
			assert.Equal(t, 1, loc.ItemCount)
		}
	}

}

func Test_Locations_Create(t *testing.T) {
	loc, err := testRepos.Locations.Create(context.Background(), testGroup.ID, locationFactory())
	assert.NoError(t, err)

	// Get by ID
	foundLoc, err := testRepos.Locations.Get(context.Background(), loc.ID)
	assert.NoError(t, err)
	assert.Equal(t, loc.ID, foundLoc.ID)

	err = testRepos.Locations.Delete(context.Background(), loc.ID)
	assert.NoError(t, err)
}

func Test_Locations_Update(t *testing.T) {
	loc, err := testRepos.Locations.Create(context.Background(), testGroup.ID, locationFactory())
	assert.NoError(t, err)

	updateData := types.LocationUpdate{
		ID:          loc.ID,
		Name:        fk.RandomString(10),
		Description: fk.RandomString(100),
	}

	update, err := testRepos.Locations.Update(context.Background(), updateData)
	assert.NoError(t, err)

	foundLoc, err := testRepos.Locations.Get(context.Background(), loc.ID)
	assert.NoError(t, err)

	assert.Equal(t, update.ID, foundLoc.ID)
	assert.Equal(t, update.Name, foundLoc.Name)
	assert.Equal(t, update.Description, foundLoc.Description)

	err = testRepos.Locations.Delete(context.Background(), loc.ID)
	assert.NoError(t, err)
}

func Test_Locations_Delete(t *testing.T) {
	loc, err := testRepos.Locations.Create(context.Background(), testGroup.ID, locationFactory())
	assert.NoError(t, err)

	err = testRepos.Locations.Delete(context.Background(), loc.ID)
	assert.NoError(t, err)

	_, err = testRepos.Locations.Get(context.Background(), loc.ID)
	assert.Error(t, err)
}
