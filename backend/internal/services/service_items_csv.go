package services

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

var ErrInvalidCsv = errors.New("invalid csv")

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

type csvRow struct {
	Location       string
	Labels         string
	Name           string
	Description    string
	SerialNumber   string
	ModelNumber    string
	Manufacturer   string
	Notes          string
	PurchaseFrom   string
	PurchasedPrice string
	PurchasedAt    string
	SoldTo         string
	SoldPrice      string
	SoldAt         string
}

func newCsvRow(row []string) csvRow {
	return csvRow{
		Location:       row[0],
		Labels:         row[1],
		Name:           row[2],
		Description:    row[3],
		SerialNumber:   row[4],
		ModelNumber:    row[5],
		Manufacturer:   row[6],
		Notes:          row[7],
		PurchaseFrom:   row[8],
		PurchasedPrice: row[9],
		PurchasedAt:    row[10],
		SoldTo:         row[11],
		SoldPrice:      row[12],
		SoldAt:         row[13],
	}
}

func (c csvRow) parsedSoldPrice() float64 {
	return parseFloat(c.SoldPrice)
}

func (c csvRow) parsedPurchasedPrice() float64 {
	return parseFloat(c.PurchasedPrice)
}

func (c csvRow) parsedPurchasedAt() time.Time {
	return parseDate(c.PurchasedAt)
}

func (c csvRow) parsedSoldAt() time.Time {
	return parseDate(c.SoldAt)
}

func (c csvRow) getLabels() []string {
	split := strings.Split(c.Labels, ";")

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
