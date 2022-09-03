package types

import (
	"time"

	"github.com/google/uuid"
)

type LocationCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type LocationUpdate struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type LocationSummary struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type LocationCount struct {
	LocationSummary
	ItemCount int `json:"itemCount"`
}

type LocationOut struct {
	LocationSummary
	Items []*ItemSummary `json:"items"`
}
