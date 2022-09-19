package types

import (
	"time"

	"github.com/google/uuid"
)

type ItemCreate struct {
	ImportRef   string `json:"-"`
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
	Quantity    int       `json:"quantity"`
	Insured     bool      `json:"insured"`

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
	PurchasePrice float64   `json:"purchasePrice,string"`

	// Sold
	SoldTime  time.Time `json:"soldTime"`
	SoldTo    string    `json:"soldTo"`
	SoldPrice float64   `json:"soldPrice,string"`
	SoldNotes string    `json:"soldNotes"`

	// Extras
	Notes string `json:"notes"`
	// Fields []*FieldSummary `json:"fields"`
}

type ItemSummary struct {
	ImportRef   string    `json:"-"`
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Quantity    int       `json:"quantity"`
	Insured     bool      `json:"insured"`

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
	PurchasePrice float64   `json:"purchasePrice,string"`

	// Sold
	SoldTime  time.Time `json:"soldTime"`
	SoldTo    string    `json:"soldTo"`
	SoldPrice float64   `json:"soldPrice,string"`
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
	Type      string      `json:"type"`
	Document  DocumentOut `json:"document"`
}

type ItemAttachmentToken struct {
	Token string `json:"token"`
}
