package types

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNameEmpty  = errors.New("name is empty")
	ErrEmailEmpty = errors.New("email is empty")
)

// UserIn is a basic user input struct containing only the fields that are
// required for user creation.
type UserIn struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserCreate is the Data object contain the requirements of creating a user
// in the database. It should to create users from an API unless the user has
// rights to create SuperUsers. For regular user in data use the UserIn struct.
type UserCreate struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	IsSuperuser bool      `json:"isSuperuser"`
	GroupID     uuid.UUID `json:"groupID"`
}

func (u *UserCreate) Validate() error {
	if u.Name == "" {
		return ErrNameEmpty
	}
	if u.Email == "" {
		return ErrEmailEmpty
	}
	return nil
}

type UserUpdate struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

type UserRegistration struct {
	User      UserIn `json:"user"`
	GroupName string `json:"groupName"`
}
