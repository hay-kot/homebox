package services

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/hay-kot/homebox/backend/internal/data/repo"
)

var ErrInvalidCsv = errors.New("invalid csv")

const NumOfCols = 21

func parseFloat(s string) float64 {
	if s == "" {
		return 0
	}
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func parseDate(s string) time.Time {
	if s == "" {
		return time.Time{}
	}

	p, _ := time.Parse("01/02/2006", s)
	return p
}

func parseBool(s string) bool {
	switch strings.ToLower(s) {
	case "true", "yes", "1":
		return true
	default:
		return false
	}
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type csvRow struct {
	Item     repo.ItemOut
	Location string
	LabelStr string
}

func newCsvRow(row []string) csvRow {
	return csvRow{
		Location: row[1],
		LabelStr: row[2],
		Item: repo.ItemOut{
			ItemSummary: repo.ItemSummary{
				ImportRef:   row[0],
				Quantity:    parseInt(row[3]),
				Name:        row[4],
				Description: row[5],
				Insured:     parseBool(row[6]),
			},
			SerialNumber:     row[7],
			ModelNumber:      row[8],
			Manufacturer:     row[9],
			Notes:            row[10],
			PurchaseFrom:     row[11],
			PurchasePrice:    parseFloat(row[12]),
			PurchaseTime:     parseDate(row[13]),
			LifetimeWarranty: parseBool(row[14]),
			WarrantyExpires:  parseDate(row[15]),
			WarrantyDetails:  row[16],
			SoldTo:           row[17],
			SoldPrice:        parseFloat(row[18]),
			SoldTime:         parseDate(row[19]),
			SoldNotes:        row[20],
		},
	}
}

func (c csvRow) getLabels() []string {
	split := strings.Split(c.LabelStr, ";")

	// Trim each
	for i, s := range split {
		split[i] = strings.TrimSpace(s)
	}

	// Remove empty
	for i, s := range split {
		if s == "" {
			split = append(split[:i], split[i+1:]...)
		}
	}

	return split
}

func (c csvRow) validate() []error {
	var errs []error

	add := func(err error) {
		errs = append(errs, err)
	}

	required := func(s string, name string) {
		if s == "" {
			add(errors.New(name + " is required"))
		}
	}

	required(c.Location, "Location")
	required(c.Item.Name, "Name")

	return errs
}
