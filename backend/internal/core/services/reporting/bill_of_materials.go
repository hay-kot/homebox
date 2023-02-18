package reporting

import (
	"context"

	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/types"
)

// =================================================================================================

type BillOfMaterialsEntry struct {
	PurchaseDate types.Date `csv:"Purchase Date"`
	Name         string     `csv:"Name"`
	Description  string     `csv:"Description"`
	Manufacturer string     `csv:"Manufacturer"`
	SerialNumber string     `csv:"Serial Number"`
	ModelNumber  string     `csv:"Model Number"`
	Quantity     int        `csv:"Quantity"`
	Price        float64    `csv:"Price"`
	TotalPrice   float64    `csv:"Total Price"`
}

// BillOfMaterialsTSV returns a byte slice of the Bill of Materials for a given GID in TSV format
// See BillOfMaterialsEntry for the format of the output
func (rs *ReportingService) BillOfMaterialsTSV(ctx context.Context, GID uuid.UUID) ([]byte, error) {
	entities, err := rs.repos.Items.GetAll(ctx, GID)
	if err != nil {
		rs.l.Debug().Err(err).Msg("failed to get all items for BOM Csv Reporting")
		return nil, err
	}

	bomEntries := make([]BillOfMaterialsEntry, len(entities))
	for i, entity := range entities {
		bomEntries[i] = BillOfMaterialsEntry{
			PurchaseDate: entity.PurchaseTime,
			Name:         entity.Name,
			Description:  entity.Description,
			Manufacturer: entity.Manufacturer,
			SerialNumber: entity.SerialNumber,
			ModelNumber:  entity.ModelNumber,
			Quantity:     entity.Quantity,
			Price:        entity.PurchasePrice,
			TotalPrice:   entity.PurchasePrice * float64(entity.Quantity),
		}
	}

	return gocsv.MarshalBytes(&bomEntries)
}
