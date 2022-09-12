package repo

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/internal/types"
	"github.com/stretchr/testify/assert"
)

func itemFactory() types.ItemCreate {
	return types.ItemCreate{
		Name:        fk.Str(10),
		Description: fk.Str(100),
	}
}

func useItems(t *testing.T, len int) []*ent.Item {
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

	t.Cleanup(func() {
		for _, item := range items {
			_ = tRepos.Items.Delete(context.Background(), item.ID)
		}
	})

	return items
}

func TestItemsRepository_GetOne(t *testing.T) {
	entity := useItems(t, 3)

	for _, item := range entity {
		result, err := tRepos.Items.GetOne(context.Background(), item.ID)
		assert.NoError(t, err)
		assert.Equal(t, item.ID, result.ID)
	}
}

func TestItemsRepository_GetAll(t *testing.T) {
	length := 10
	expected := useItems(t, length)

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
	entities := useItems(t, 3)

	for _, item := range entities {
		err := tRepos.Items.Delete(context.Background(), item.ID)
		assert.NoError(t, err)
	}

	results, err := tRepos.Items.GetAll(context.Background(), tGroup.ID)
	assert.NoError(t, err)
	assert.Empty(t, results)
}

func TestItemsRepository_Update_Labels(t *testing.T) {
	entity := useItems(t, 1)[0]
	labels := useLabels(t, 3)

	labelsIDs := []uuid.UUID{labels[0].ID, labels[1].ID, labels[2].ID}

	type args struct {
		labelIds []uuid.UUID
	}

	tests := []struct {
		name string
		args args
		want []uuid.UUID
	}{
		{
			name: "add all labels",
			args: args{
				labelIds: labelsIDs,
			},
			want: labelsIDs,
		},
		{
			name: "update with one label",
			args: args{
				labelIds: labelsIDs[:1],
			},
			want: labelsIDs[:1],
		},
		{
			name: "add one new label to existing single label",
			args: args{
				labelIds: labelsIDs[1:],
			},
			want: labelsIDs[1:],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Apply all labels to entity
			updateData := types.ItemUpdate{
				ID:         entity.ID,
				Name:       entity.Name,
				LocationID: entity.Edges.Location.ID,
				LabelIDs:   tt.args.labelIds,
			}

			updated, err := tRepos.Items.Update(context.Background(), updateData)
			assert.NoError(t, err)
			assert.Len(t, tt.want, len(updated.Edges.Label))

			for _, label := range updated.Edges.Label {
				assert.Contains(t, tt.want, label.ID)
			}
		})
	}

}

func TestItemsRepository_Update(t *testing.T) {
	entities := useItems(t, 3)

	entity := entities[0]

	updateData := types.ItemUpdate{
		ID:               entity.ID,
		Name:             entity.Name,
		LocationID:       entity.Edges.Location.ID,
		SerialNumber:     fk.Str(10),
		LabelIDs:         nil,
		ModelNumber:      fk.Str(10),
		Manufacturer:     fk.Str(10),
		PurchaseTime:     time.Now(),
		PurchaseFrom:     fk.Str(10),
		PurchasePrice:    300.99,
		SoldTime:         time.Now(),
		SoldTo:           fk.Str(10),
		SoldPrice:        300.99,
		SoldNotes:        fk.Str(10),
		Notes:            fk.Str(10),
		WarrantyExpires:  time.Now(),
		WarrantyDetails:  fk.Str(10),
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
