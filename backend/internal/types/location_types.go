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
	GroupID     uuid.UUID `json:"groupId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ItemSummary struct {
	ID          uuid.UUID `json:"id"`
	LocationID  uuid.UUID `json:"locationId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type LocationOut struct {
	LocationSummary
	Items []*ItemSummary `json:"items"`
}
