package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/git-web-template/backend/internal/repo"
	"github.com/hay-kot/git-web-template/backend/internal/types"
)

type AdminService struct {
	repos *repo.AllRepos
}

func (svc *AdminService) Create(ctx context.Context, usr types.UserCreate) (types.UserOut, error) {
	return svc.repos.Users.Create(ctx, usr)
}

func (svc *AdminService) GetAll(ctx context.Context) ([]types.UserOut, error) {
	return svc.repos.Users.GetAll(ctx)
}

func (svc *AdminService) GetByID(ctx context.Context, id uuid.UUID) (types.UserOut, error) {
	return svc.repos.Users.GetOneId(ctx, id)
}

func (svc *AdminService) GetByEmail(ctx context.Context, email string) (types.UserOut, error) {
	return svc.repos.Users.GetOneEmail(ctx, email)
}

func (svc *AdminService) UpdateProperties(ctx context.Context, ID uuid.UUID, data types.UserUpdate) (types.UserOut, error) {
	err := svc.repos.Users.Update(ctx, ID, data)

	if err != nil {
		return types.UserOut{}, err
	}

	return svc.repos.Users.GetOneId(ctx, ID)
}

func (svc *AdminService) Delete(ctx context.Context, id uuid.UUID) error {
	return svc.repos.Users.Delete(ctx, id)
}

func (svc *AdminService) DeleteAll(ctx context.Context) error {
	return svc.repos.Users.DeleteAll(ctx)
}
