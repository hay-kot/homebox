package repo

import (
	"context"
	"testing"
	"time"

	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/internal/types"
	"github.com/stretchr/testify/assert"
)

func itemFactory() types.ItemCreate {
	return types.ItemCreate{
		Name:        fk.RandomString(10),
		Description: fk.RandomString(100),
	}
}

func useItems(t *testing.T, len int) ([]*ent.Item, func()) {
	t.Helper()

	location, err := tRepos.Locations.Create(context.Background(), tGroup.ID, locationFactory())
	assert.NoError(t, err)

	items := make([]*ent.Item, len)
	for i := 0; i < len; i++ {
		itm := itemFactory()
		itm.LocationID = location.ID

		item, err := tRepos.Items.Create(context.Background(), tGroup.ID, itm)
		assert.NoError(t, err)
		items[i] = item
	}

	return items, func() {
		for _, item := range items {
			err := tRepos.Items.Delete(context.Background(), item.ID)
			assert.NoError(t, err)
		}
	}
}

func TestItemsRepository_GetOne(t *testing.T) {
	entity, cleanup := useItems(t, 3)
	defer cleanup()

	for _, item := range entity {
		result, err := tRepos.Items.GetOne(context.Background(), item.ID)
		assert.NoError(t, err)
		assert.Equal(t, item.ID, result.ID)
	}
}

func TestItemsRepository_GetAll(t *testing.T) {
	length := 10
	expected, cleanup := useItems(t, length)
	defer cleanup()

	results, err := tRepos.Items.GetAll(context.Background(), tGroup.ID)
	assert.NoError(t, err)

	assert.Equal(t, length, len(results))

	for _, item := range results {
		for _, expectedItem := range expected {
			if item.ID == expectedItem.ID {
				assert.Equal(t, expectedItem.ID, item.ID)
				assert.Equal(t, expectedItem.Name, item.Name)
				assert.Equal(t, expectedItem.Description, item.Description)
			}
		}
	}
}

func TestItemsRepository_Create(t *testing.T) {
	location, err := tRepos.Locations.Create(context.Background(), tGroup.ID, locationFactory())
	assert.NoError(t, err)

	itm := itemFactory()
	itm.LocationID = location.ID

	result, err := tRepos.Items.Create(context.Background(), tGroup.ID, itm)
	assert.NoError(t, err)
	assert.NotEmpty(t, result.ID)

	// Cleanup
	err = tRepos.Locations.Delete(context.Background(), location.ID)
	assert.NoError(t, err)

	err = tRepos.Items.Delete(context.Background(), result.ID)
	assert.NoError(t, err)
}

func TestItemsRepository_Create_Location(t *testing.T) {
	location, err := tRepos.Locations.Create(context.Background(), tGroup.ID, locationFactory())
	assert.NoError(t, err)
	assert.NotEmpty(t, location.ID)

	item := itemFactory()
	item.LocationID = location.ID

	// Create Resource
	result, err := tRepos.Items.Create(context.Background(), tGroup.ID, item)
	assert.NoError(t, err)
	assert.NotEmpty(t, result.ID)

	// Get Resource
	foundItem, err := tRepos.Items.GetOne(context.Background(), result.ID)
	assert.NoError(t, err)
	assert.Equal(t, result.ID, foundItem.ID)
	assert.Equal(t, location.ID, foundItem.Edges.Location.ID)

	// Cleanup
	err = tRepos.Locations.Delete(context.Background(), location.ID)
	assert.NoError(t, err)
	err = tRepos.Items.Delete(context.Background(), result.ID)
	assert.NoError(t, err)
}

func TestItemsRepository_Delete(t *testing.T) {
	entities, _ := useItems(t, 3)

	for _, item := range entities {
		err := tRepos.Items.Delete(context.Background(), item.ID)
		assert.NoError(t, err)
	}

	results, err := tRepos.Items.GetAll(context.Background(), tGroup.ID)
	assert.NoError(t, err)
	assert.Empty(t, results)
}

func TestItemsRepository_Update(t *testing.T) {
	entities, cleanup := useItems(t, 3)
	defer cleanup()

	entity := entities[0]

	updateData := types.ItemUpdate{
		ID:               entity.ID,
		Name:             entity.Name,
		LocationID:       entity.Edges.Location.ID,
		SerialNumber:     fk.RandomString(10),
		LabelIDs:         nil,
		ModelNumber:      fk.RandomString(10),
		Manufacturer:     fk.RandomString(10),
		PurchaseTime:     time.Now(),
		PurchaseFrom:     fk.RandomString(10),
		PurchasePrice:    300.99,
		SoldTime:         time.Now(),
		SoldTo:           fk.RandomString(10),
		SoldPrice:        300.99,
		SoldNotes:        fk.RandomString(10),
		Notes:            fk.RandomString(10),
		WarrantyExpires:  time.Now(),
		WarrantyDetails:  fk.RandomString(10),
		LifetimeWarranty: true,
	}

	updatedEntity, err := tRepos.Items.Update(context.Background(), updateData)
	assert.NoError(t, err)

	got, err := tRepos.Items.GetOne(context.Background(), updatedEntity.ID)
	assert.NoError(t, err)

	assert.Equal(t, updateData.ID, got.ID)
	assert.Equal(t, updateData.Name, got.Name)
	assert.Equal(t, updateData.LocationID, got.Edges.Location.ID)
	assert.Equal(t, updateData.SerialNumber, got.SerialNumber)
	assert.Equal(t, updateData.ModelNumber, got.ModelNumber)
	assert.Equal(t, updateData.Manufacturer, got.Manufacturer)
	// assert.Equal(t, updateData.PurchaseTime, got.PurchaseTime)
	assert.Equal(t, updateData.PurchaseFrom, got.PurchaseFrom)
	assert.Equal(t, updateData.PurchasePrice, got.PurchasePrice)
	// assert.Equal(t, updateData.SoldTime, got.SoldTime)
	assert.Equal(t, updateData.SoldTo, got.SoldTo)
	assert.Equal(t, updateData.SoldPrice, got.SoldPrice)
	assert.Equal(t, updateData.SoldNotes, got.SoldNotes)
	assert.Equal(t, updateData.Notes, got.Notes)
	// assert.Equal(t, updateData.WarrantyExpires, got.WarrantyExpires)
	assert.Equal(t, updateData.WarrantyDetails, got.WarrantyDetails)
	assert.Equal(t, updateData.LifetimeWarranty, got.LifetimeWarranty)
}
