package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/repo"
)

var (
	ErrNotOwner = errors.New("not owner")
)

type LocationService struct {
	repos *repo.AllRepos
}

func (svc *LocationService) GetOne(ctx context.Context, groupId uuid.UUID, id uuid.UUID) (repo.LocationOut, error) {
	return svc.repos.Locations.GetOneByGroup(ctx, groupId, id)
}

func (svc *LocationService) GetAll(ctx context.Context, groupId uuid.UUID) ([]repo.LocationOutCount, error) {
	return svc.repos.Locations.GetAll(ctx, groupId)
}

func (svc *LocationService) Create(ctx context.Context, groupId uuid.UUID, data repo.LocationCreate) (repo.LocationOut, error) {
	return svc.repos.Locations.Create(ctx, groupId, data)
}

func (svc *LocationService) Delete(ctx context.Context, groupId uuid.UUID, id uuid.UUID) error {
	_, err := svc.repos.Locations.GetOneByGroup(ctx, groupId, id)
	if err != nil {
		return err
	}
	return svc.repos.Locations.Delete(ctx, id)
}

func (svc *LocationService) Update(ctx context.Context, groupId uuid.UUID, data repo.LocationUpdate) (repo.LocationOut, error) {
	location, err := svc.repos.Locations.GetOneByGroup(ctx, groupId, data.ID)
	if err != nil {
		return repo.LocationOut{}, err
	}

	data.ID = location.ID
	return svc.repos.Locations.Update(ctx, data)
}
