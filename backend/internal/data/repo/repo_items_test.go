package repo

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func itemFactory() ItemCreate {
	return ItemCreate{
		Name:        fk.Str(10),
		Description: fk.Str(100),
	}
}

func useItems(t *testing.T, len int) []ItemOut {
	t.Helper()

	location, err := tRepos.Locations.Create(context.Background(), tGroup.ID, locationFactory())
	assert.NoError(t, err)

	items := make([]ItemOut, len)
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

		_ = tRepos.Locations.delete(context.Background(), location.ID)
	})

	return items
}

func TestItemsRepository_RecursiveRelationships(t *testing.T) {
	parent := useItems(t, 1)[0]

	children := useItems(t, 3)

	for _, child := range children {
		update := ItemUpdate{
			ID:          child.ID,
			ParentID:    parent.ID,
			Name:        "note-important",
			Description: "This is a note",
			LocationID:  child.Location.ID,
		}

		// Append Parent ID
		_, err := tRepos.Items.UpdateByGroup(context.Background(), tGroup.ID, update)
		assert.NoError(t, err)

		// Check Parent ID
		updated, err := tRepos.Items.GetOne(context.Background(), child.ID)
		assert.NoError(t, err)
		assert.Equal(t, parent.ID, updated.Parent.ID)

		// Remove Parent ID
		update.ParentID = uuid.Nil
		_, err = tRepos.Items.UpdateByGroup(context.Background(), tGroup.ID, update)
		assert.NoError(t, err)

		// Check Parent ID
		updated, err = tRepos.Items.GetOne(context.Background(), child.ID)
		assert.NoError(t, err)
		assert.Nil(t, updated.Parent)

	}
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

	// Cleanup - Also deletes item
	err = tRepos.Locations.delete(context.Background(), location.ID)
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
	assert.Equal(t, location.ID, foundItem.Location.ID)

	// Cleanup - Also deletes item
	err = tRepos.Locations.delete(context.Background(), location.ID)
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
			updateData := ItemUpdate{
				ID:         entity.ID,
				Name:       entity.Name,
				LocationID: entity.Location.ID,
				LabelIDs:   tt.args.labelIds,
			}

			updated, err := tRepos.Items.UpdateByGroup(context.Background(), tGroup.ID, updateData)
			assert.NoError(t, err)
			assert.Len(t, tt.want, len(updated.Labels))

			for _, label := range updated.Labels {
				assert.Contains(t, tt.want, label.ID)
			}
		})
	}
}

func TestItemsRepository_Update(t *testing.T) {
	entities := useItems(t, 3)

	entity := entities[0]

	updateData := ItemUpdate{
		ID:               entity.ID,
		Name:             entity.Name,
		LocationID:       entity.Location.ID,
		SerialNumber:     fk.Str(10),
		LabelIDs:         nil,
		ModelNumber:      fk.Str(10),
		Manufacturer:     fk.Str(10),
		PurchaseTime:     types.DateFromTime(time.Now()),
		PurchaseFrom:     fk.Str(10),
		PurchasePrice:    300.99,
		SoldTime:         types.DateFromTime(time.Now()),
		SoldTo:           fk.Str(10),
		SoldPrice:        300.99,
		SoldNotes:        fk.Str(10),
		Notes:            fk.Str(10),
		WarrantyExpires:  types.DateFromTime(time.Now()),
		WarrantyDetails:  fk.Str(10),
		LifetimeWarranty: true,
	}

	updatedEntity, err := tRepos.Items.UpdateByGroup(context.Background(), tGroup.ID, updateData)
	assert.NoError(t, err)

	got, err := tRepos.Items.GetOne(context.Background(), updatedEntity.ID)
	assert.NoError(t, err)

	assert.Equal(t, updateData.ID, got.ID)
	assert.Equal(t, updateData.Name, got.Name)
	assert.Equal(t, updateData.LocationID, got.Location.ID)
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

func TestItemRepository_GetAllCustomFields(t *testing.T) {
	const FIELDS_COUNT = 5

	entity := useItems(t, 1)[0]

	fields := make([]ItemField, FIELDS_COUNT)
	names := make([]string, FIELDS_COUNT)
	values := make([]string, FIELDS_COUNT)

	for i := 0; i < FIELDS_COUNT; i++ {
		name := fk.Str(10)
		fields[i] = ItemField{
			Name:      name,
			Type:      "text",
			TextValue: fk.Str(10),
		}
		names[i] = name
		values[i] = fields[i].TextValue
	}

	_, err := tRepos.Items.UpdateByGroup(context.Background(), tGroup.ID, ItemUpdate{
		ID:         entity.ID,
		Name:       entity.Name,
		LocationID: entity.Location.ID,
		Fields:     fields,
	})

	require.NoError(t, err)

	// Test getting all fields
	{
		results, err := tRepos.Items.GetAllCustomFieldNames(context.Background(), tGroup.ID)
		assert.NoError(t, err)
		assert.ElementsMatch(t, names, results)
	}

	// Test getting all values from field
	{
		results, err := tRepos.Items.GetAllCustomFieldValues(context.Background(), tUser.GroupID, names[0])

		assert.NoError(t, err)
		assert.ElementsMatch(t, values[:1], results)
	}
}
