package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestItemService_CsvImport(t *testing.T) {
	data := loadcsv()
	svc := &ItemService{
		repo: tRepos,
	}
	err := svc.CsvImport(context.Background(), tGroup.ID, data)
	assert.NoError(t, err)

	items, err := svc.GetAll(context.Background(), tGroup.ID)
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

	locationService := &LocationService{
		repos: tRepos,
	}

	LabelService := &LabelService{
		repos: tRepos,
	}

	allLocation, err := locationService.GetAll(context.Background(), tGroup.ID)
	assert.NoError(t, err)
	locNames := []string{}
	for _, loc := range allLocation {
		locNames = append(locNames, loc.Name)
	}

	allLabels, err := LabelService.GetAll(context.Background(), tGroup.ID)
	assert.NoError(t, err)
	labelNames := []string{}
	for _, label := range allLabels {
		labelNames = append(labelNames, label.Name)
	}

	for _, item := range items {
		assert.Contains(t, locNames, item.Location.Name)
		for _, label := range item.Labels {
			assert.Contains(t, labelNames, label.Name)
		}

		for _, csvRow := range dataCsv {
			if csvRow.Name == item.Name {
				assert.Equal(t, csvRow.Description, item.Description)
				assert.Equal(t, csvRow.SerialNumber, item.SerialNumber)
				assert.Equal(t, csvRow.Manufacturer, item.Manufacturer)
				assert.Equal(t, csvRow.Notes, item.Notes)

				// Purchase Fields
				assert.Equal(t, csvRow.parsedPurchasedAt(), item.PurchaseTime)
				assert.Equal(t, csvRow.PurchaseFrom, item.PurchaseFrom)
				assert.Equal(t, csvRow.parsedPurchasedPrice(), item.PurchasePrice)

				// Sold Fields
				assert.Equal(t, csvRow.parsedSoldAt(), item.SoldTime)
				assert.Equal(t, csvRow.SoldTo, item.SoldTo)
				assert.Equal(t, csvRow.parsedSoldPrice(), item.SoldPrice)
			}
		}

	}
}
