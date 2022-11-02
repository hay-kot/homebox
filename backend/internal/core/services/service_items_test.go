package services

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/stretchr/testify/assert"
)

func TestItemService_CsvImport(t *testing.T) {
	data := loadcsv()
	svc := &ItemService{
		repo: tRepos,
	}
	count, err := svc.CsvImport(context.Background(), tGroup.ID, data)
	assert.Equal(t, 6, count)
	assert.NoError(t, err)

	// Check import refs are deduplicated
	count, err = svc.CsvImport(context.Background(), tGroup.ID, data)
	assert.Equal(t, 0, count)
	assert.NoError(t, err)

	items, err := svc.repo.Items.GetAll(context.Background(), tGroup.ID)
	assert.NoError(t, err)
	t.Cleanup(func() {
		for _, item := range items {
			err := svc.repo.Items.Delete(context.Background(), item.ID)
			assert.NoError(t, err)
		}
	})

	assert.Equal(t, len(items), 6)

	dataCsv := []csvRow{}
	for _, item := range data {
		dataCsv = append(dataCsv, newCsvRow(item))
	}

	allLocation, err := tRepos.Locations.GetAll(context.Background(), tGroup.ID, repo.LocationQuery{})
	assert.NoError(t, err)
	locNames := []string{}
	for _, loc := range allLocation {
		locNames = append(locNames, loc.Name)
	}

	allLabels, err := tRepos.Labels.GetAll(context.Background(), tGroup.ID)
	assert.NoError(t, err)
	labelNames := []string{}
	for _, label := range allLabels {
		labelNames = append(labelNames, label.Name)
	}

	ids := []uuid.UUID{}
	t.Cleanup((func() {
		for _, id := range ids {
			err := svc.repo.Items.Delete(context.Background(), id)
			assert.NoError(t, err)
		}
	}))

	for _, item := range items {
		assert.Contains(t, locNames, item.Location.Name)
		for _, label := range item.Labels {
			assert.Contains(t, labelNames, label.Name)
		}

		for _, csvRow := range dataCsv {
			if csvRow.Item.Name == item.Name {
				assert.Equal(t, csvRow.Item.Description, item.Description)
				assert.Equal(t, csvRow.Item.Quantity, item.Quantity)
				assert.Equal(t, csvRow.Item.Insured, item.Insured)
			}
		}
	}
}
