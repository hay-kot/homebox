package reporting

import (
	"context"
	"encoding/csv"
	"io"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/rs/zerolog"
)

type ReportingService struct {
	repos *repo.AllRepos
	l     *zerolog.Logger
}

func NewReportingService(repos *repo.AllRepos, l *zerolog.Logger) *ReportingService {
	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = '\t'
		return gocsv.NewSafeCSVWriter(writer)
	})

	return &ReportingService{
		repos: repos,
		l:     l,
	}
}

// =================================================================================================

// NullableTime is a custom type that implements the MarshalCSV interface
// to allow for nullable time.Time fields in the CSV output to be empty
// and not "0001-01-01". It also overrides the default CSV output format
type NullableTime time.Time

func (t NullableTime) MarshalCSV() (string, error) {
	if time.Time(t).IsZero() {
		return "", nil
	}
	// YYYY-MM-DD
	return time.Time(t).Format("2006-01-02"), nil
}

type BillOfMaterialsEntry struct {
	PurchaseDate NullableTime `csv:"Purchase Date"`
	Name         string       `csv:"Name"`
	Description  string       `csv:"Description"`
	Manufacturer string       `csv:"Manufacturer"`
	SerialNumber string       `csv:"Serial Number"`
	ModelNumber  string       `csv:"Model Number"`
	Quantity     int          `csv:"Quantity"`
	Price        float64      `csv:"Price"`
	TotalPrice   float64      `csv:"Total Price"`
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
			PurchaseDate: NullableTime(entity.PurchaseTime),
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
