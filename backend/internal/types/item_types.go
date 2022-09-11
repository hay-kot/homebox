package types

import (
	"time"

	"github.com/google/uuid"
)

type ItemCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	// Edges
	LocationID uuid.UUID   `json:"locationId"`
	LabelIDs   []uuid.UUID `json:"labelIds"`
}

type ItemUpdate struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`

	// Edges
	LocationID uuid.UUID   `json:"locationId"`
	LabelIDs   []uuid.UUID `json:"labelIds"`

	// Identifications
	SerialNumber string `json:"serialNumber"`
	ModelNumber  string `json:"modelNumber"`
	Manufacturer string `json:"manufacturer"`

	// Warranty
	LifetimeWarranty bool      `json:"lifetimeWarranty"`
	WarrantyExpires  time.Time `json:"warrantyExpires"`
	WarrantyDetails  string    `json:"warrantyDetails"`

	// Purchase
	PurchaseTime  time.Time `json:"purchaseTime"`
	PurchaseFrom  string    `json:"purchaseFrom"`
	PurchasePrice float64   `json:"purchasePrice"`

	// Sold
	SoldTime  time.Time `json:"soldTime"`
	SoldTo    string    `json:"soldTo"`
	SoldPrice float64   `json:"soldPrice"`
	SoldNotes string    `json:"soldNotes"`

	// Extras
	Notes string `json:"notes"`
	// Fields []*FieldSummary `json:"fields"`
}

type ItemSummary struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`

	// Edges
	Location *LocationSummary `json:"location"`
	Labels   []*LabelSummary  `json:"labels"`

	// Identifications
	SerialNumber string `json:"serialNumber"`
	ModelNumber  string `json:"modelNumber"`
	Manufacturer string `json:"manufacturer"`

	// Warranty
	LifetimeWarranty bool      `json:"lifetimeWarranty"`
	WarrantyExpires  time.Time `json:"warrantyExpires"`
	WarrantyDetails  string    `json:"warrantyDetails"`

	// Purchase
	PurchaseTime  time.Time `json:"purchaseTime"`
	PurchaseFrom  string    `json:"purchaseFrom"`
	PurchasePrice float64   `json:"purchasePrice"`

	// Sold
	SoldTime  time.Time `json:"soldTime"`
	SoldTo    string    `json:"soldTo"`
	SoldPrice float64   `json:"soldPrice"`
	SoldNotes string    `json:"soldNotes"`

	// Extras
	Notes string `json:"notes"`
}

type ItemOut struct {
	ItemSummary
	Attachments []*ItemAttachment `json:"attachments"`
	// Future
	// Fields []*FieldSummary `json:"fields"`
}

type ItemAttachment struct {
	ID        uuid.UUID   `json:"id"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	Document  DocumentOut `json:"document"`
}
