package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent"
	"github.com/hay-kot/homebox/backend/internal/data/ent/maintenanceentry"
)

// MaintenanceEntryRepository is a repository for maintenance entries that are
// associated with an item in the database. An entry represents a maintenance event
// that has been performed on an item.
type MaintenanceEntryRepository struct {
	db *ent.Client
}
type (
	MaintenanceEntryCreate struct {
		Date        time.Time `json:"date"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Cost        float64   `json:"cost,string"`
	}

	MaintenanceEntry struct {
		ID          uuid.UUID `json:"id"`
		Date        time.Time `json:"date"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Cost        float64   `json:"cost,string"`
	}

	MaintenanceEntryUpdate struct {
		Date        time.Time `json:"date"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Cost        float64   `json:"cost,string"`
	}

	MaintenanceLog struct {
		ItemID      uuid.UUID          `json:"itemId"`
		CostAverage float64            `json:"costAverage"`
		CostTotal   float64            `json:"costTotal"`
		Entries     []MaintenanceEntry `json:"entries"`
	}
)

var (
	mapMaintenanceEntryErr  = mapTErrFunc(mapMaintenanceEntry)
	mapEachMaintenanceEntry = mapTEachFunc(mapMaintenanceEntry)
)

func mapMaintenanceEntry(entry *ent.MaintenanceEntry) MaintenanceEntry {
	return MaintenanceEntry{
		ID:          entry.ID,
		Date:        entry.Date,
		Name:        entry.Name,
		Description: entry.Description,
		Cost:        entry.Cost,
	}
}

func (r *MaintenanceEntryRepository) Create(ctx context.Context, itemID uuid.UUID, input MaintenanceEntryCreate) (MaintenanceEntry, error) {
	item, err := r.db.MaintenanceEntry.Create().
		SetItemID(itemID).
		SetDate(input.Date).
		SetName(input.Name).
		SetDescription(input.Description).
		SetCost(input.Cost).
		Save(ctx)

	return mapMaintenanceEntryErr(item, err)
}

func (r *MaintenanceEntryRepository) Update(ctx context.Context, ID uuid.UUID, input MaintenanceEntryUpdate) (MaintenanceEntry, error) {
	item, err := r.db.MaintenanceEntry.UpdateOneID(ID).
		SetDate(input.Date).
		SetName(input.Name).
		SetDescription(input.Description).
		SetCost(input.Cost).
		Save(ctx)

	return mapMaintenanceEntryErr(item, err)
}

func (r *MaintenanceEntryRepository) GetLog(ctx context.Context, itemID uuid.UUID) (MaintenanceLog, error) {
	log := MaintenanceLog{
		ItemID: itemID,
	}

	entries, err := r.db.MaintenanceEntry.Query().
		Where(maintenanceentry.ItemID(itemID)).
		Order(ent.Desc(maintenanceentry.FieldDate)).
		All(ctx)
	if err != nil {
		return MaintenanceLog{}, err
	}

	log.Entries = mapEachMaintenanceEntry(entries)

	var maybeTotal *float64
	var maybeAverage *float64

	q := `
SELECT
  SUM(cost_total) AS total_of_totals,
  AVG(cost_total) AS avg_of_averages
FROM
  (
    SELECT
      strftime('%m-%Y', date) AS my,
      SUM(cost) AS cost_total
    FROM
      maintenance_entries
    WHERE
      item_id = ?
    GROUP BY
      my
  )`

	row := r.db.Sql().QueryRowContext(ctx, q, itemID)
	err = row.Scan(&maybeTotal, &maybeAverage)
	if err != nil {
		return MaintenanceLog{}, err
	}

	log.CostAverage = orDefault(maybeAverage, 0)
	log.CostTotal = orDefault(maybeTotal, 0)
	return log, nil
}

func (r *MaintenanceEntryRepository) Delete(ctx context.Context, ID uuid.UUID) error {
	return r.db.MaintenanceEntry.DeleteOneID(ID).Exec(ctx)
}
