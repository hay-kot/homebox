package types

import (
	"time"

	"github.com/google/uuid"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	BearerToken string    `json:"token"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

type UserAuthTokenDetail struct {
	Raw       string    `json:"raw"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type UserAuthToken struct {
	TokenHash []byte    `json:"token"`
	UserID    uuid.UUID `json:"userId"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u UserAuthToken) IsExpired() bool {
	return u.ExpiresAt.Before(time.Now())
}

type UserAuthTokenCreate struct {
	TokenHash []byte    `json:"token"`
	UserID    uuid.UUID `json:"userId"`
	ExpiresAt time.Time `json:"expiresAt"`
}
