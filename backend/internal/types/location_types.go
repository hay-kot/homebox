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

type LocationOut struct {
	ID          uuid.UUID `json:"id"`
	GroupID     uuid.UUID `json:"groupId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
