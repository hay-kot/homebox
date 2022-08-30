package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/git-web-template/backend/internal/types"
)

type UserRepository interface {
	// GetOneId returns a user by id
	GetOneId(ctx context.Context, ID uuid.UUID) (types.UserOut, error)
	// GetOneEmail returns a user by email
	GetOneEmail(ctx context.Context, email string) (types.UserOut, error)
	// GetAll returns all users
	GetAll(ctx context.Context) ([]types.UserOut, error)
	// Get Super Users
	GetSuperusers(ctx context.Context) ([]types.UserOut, error)
	// Create creates a new user
	Create(ctx context.Context, user types.UserCreate) (types.UserOut, error)
	// Update updates a user
	Update(ctx context.Context, ID uuid.UUID, user types.UserUpdate) error
	// Delete deletes a user
	Delete(ctx context.Context, ID uuid.UUID) error

	DeleteAll(ctx context.Context) error
}
