package services

import (
	"context"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/types"
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
				assert.Equal(t, csvRow.Item.SerialNumber, item.SerialNumber)
				assert.Equal(t, csvRow.Item.Manufacturer, item.Manufacturer)
				assert.Equal(t, csvRow.Item.Notes, item.Notes)

				// Purchase Fields
				assert.Equal(t, csvRow.Item.PurchaseTime, item.PurchaseTime)
				assert.Equal(t, csvRow.Item.PurchaseFrom, item.PurchaseFrom)
				assert.Equal(t, csvRow.Item.PurchasePrice, item.PurchasePrice)

				// Sold Fields
				assert.Equal(t, csvRow.Item.SoldTime, item.SoldTime)
				assert.Equal(t, csvRow.Item.SoldTo, item.SoldTo)
				assert.Equal(t, csvRow.Item.SoldPrice, item.SoldPrice)
			}
		}
	}
}

func TestItemService_AddAttachment(t *testing.T) {
	temp := os.TempDir()

	svc := &ItemService{
		repo:     tRepos,
		filepath: temp,
	}

	loc, err := tSvc.Location.Create(context.Background(), tGroup.ID, types.LocationCreate{
		Description: "test",
		Name:        "test",
	})
	assert.NoError(t, err)
	assert.NotNil(t, loc)

	itmC := types.ItemCreate{
		Name:        fk.Str(10),
		Description: fk.Str(10),
		LocationID:  loc.ID,
	}

	itm, err := svc.Create(context.Background(), tGroup.ID, itmC)
	assert.NoError(t, err)
	assert.NotNil(t, itm)
	t.Cleanup(func() {
		err := svc.repo.Items.Delete(context.Background(), itm.ID)
		assert.NoError(t, err)
	})

	contents := fk.Str(1000)
	reader := strings.NewReader(contents)

	// Setup
	afterAttachment, err := svc.AddAttachment(context.Background(), tGroup.ID, itm.ID, "testfile.txt", "attachment", reader)
	assert.NoError(t, err)
	assert.NotNil(t, afterAttachment)

	// Check that the file exists
	storedPath := afterAttachment.Attachments[0].Document.Path

	// {root}/{group}/{item}/{attachment}
	assert.Equal(t, path.Join(temp, tGroup.ID.String(), itm.ID.String(), "testfile.txt"), storedPath)

	// Check that the file contents are correct
	bts, err := os.ReadFile(storedPath)
	assert.NoError(t, err)
	assert.Equal(t, contents, string(bts))

}
