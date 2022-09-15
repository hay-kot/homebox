package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/internal/types"
)

type AdminService struct {
	repos *repo.AllRepos
}

func (svc *AdminService) Create(ctx context.Context, usr types.UserCreate) (*ent.User, error) {
	return svc.repos.Users.Create(ctx, usr)
}

func (svc *AdminService) GetAll(ctx context.Context) ([]*ent.User, error) {
	return svc.repos.Users.GetAll(ctx)
}

func (svc *AdminService) GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return svc.repos.Users.GetOneId(ctx, id)
}

func (svc *AdminService) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	return svc.repos.Users.GetOneEmail(ctx, email)
}

func (svc *AdminService) UpdateProperties(ctx context.Context, ID uuid.UUID, data types.UserUpdate) (*ent.User, error) {
	err := svc.repos.Users.Update(ctx, ID, data)

	if err != nil {
		return &ent.User{}, err
	}

	return svc.repos.Users.GetOneId(ctx, ID)
}

func (svc *AdminService) Delete(ctx context.Context, id uuid.UUID) error {
	return svc.repos.Users.Delete(ctx, id)
}

func (svc *AdminService) DeleteAll(ctx context.Context) error {
	return svc.repos.Users.DeleteAll(ctx)
}
