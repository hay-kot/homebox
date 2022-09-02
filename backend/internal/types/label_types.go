package types

import (
	"time"

	"github.com/google/uuid"
)

type LabelCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type LabelUpdate struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
}

type LabelSummary struct {
	ID          uuid.UUID `json:"id"`
	GroupID     uuid.UUID `json:"groupId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type LabelOut struct {
	LabelSummary
	Items []*ItemSummary `json:"items"`
}
