package repo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMaintenanceEntryRepository_GetLog(t *testing.T) {
	item := useItems(t, 1)[0]

	// Create 10 maintenance entries for the item
	created := make([]MaintenanceEntryCreate, 10)

	lastMonth := time.Now().AddDate(0, -1, 0)
	thisMonth := time.Now()

	for i := 0; i < 10; i++ {
		dt := lastMonth
		if i%2 == 0 {
			dt = thisMonth
		}

		created[i] = MaintenanceEntryCreate{
			Date:        dt,
			Name:        "Maintenance",
			Description: "Maintenance description",
			Cost:        10,
		}
	}

	for _, entry := range created {
		_, err := tRepos.MaintEntry.Create(context.Background(), item.ID, entry)
		if err != nil {
			t.Fatalf("failed to create maintenance entry: %v", err)
		}
	}

	// Get the log for the item
	log, err := tRepos.MaintEntry.GetLog(context.Background(), item.ID)

	if err != nil {
		t.Fatalf("failed to get maintenance log: %v", err)
	}

	assert.Equal(t, item.ID, log.ItemID)
	assert.Equal(t, 10, len(log.Entries))

	// Calculate the average cost
	var total float64

	for _, entry := range log.Entries {
		total += entry.Cost
	}

	assert.Equal(t, total, log.CostTotal, "total cost should be equal to the sum of all entries")
	assert.Equal(t, total/2, log.CostAverage, "average cost should be the average of the two months")

	for _, entry := range log.Entries {
		err := tRepos.MaintEntry.Delete(context.Background(), entry.ID)
		assert.NoError(t, err)
	}
}
