package types

import (
	"time"

	"github.com/google/uuid"
)

type DocumentCreate struct {
	Title string `json:"name"`
	Path  string `json:"path"`
}

type DocumentUpdate struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"name"`
	Path  string    `json:"path"`
}

type DocumentToken struct {
	Raw       string    `json:"raw"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type DocumentTokenCreate struct {
	TokenHash  []byte    `json:"tokenHash"`
	DocumentID uuid.UUID `json:"documentId"`
	ExpiresAt  time.Time `json:"expiresAt"`
}
