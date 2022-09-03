package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/internal/repo"
	"github.com/hay-kot/content/backend/internal/services/mappers"
	"github.com/hay-kot/content/backend/internal/types"
)

var (
	ErrNotOwner = errors.New("not owner")
)

type LocationService struct {
	repos *repo.AllRepos
}

func (svc *LocationService) GetOne(ctx context.Context, groupId uuid.UUID, id uuid.UUID) (*types.LocationOut, error) {
	location, err := svc.repos.Locations.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	if location.Edges.Group.ID != groupId {
		return nil, ErrNotOwner
	}

	return mappers.ToLocationOut(location), nil
}

func (svc *LocationService) GetAll(ctx context.Context, groupId uuid.UUID) ([]*types.LocationCount, error) {
	locations, err := svc.repos.Locations.GetAll(ctx, groupId)
	if err != nil {
		return nil, err
	}

	locationsOut := make([]*types.LocationCount, len(locations))
	for i, location := range locations {
		locationsOut[i] = mappers.ToLocationCount(&location)
	}

	return locationsOut, nil
}

func (svc *LocationService) Create(ctx context.Context, groupId uuid.UUID, data types.LocationCreate) (*types.LocationOut, error) {
	location, err := svc.repos.Locations.Create(ctx, groupId, data)
	return mappers.ToLocationOutErr(location, err)
}

func (svc *LocationService) Delete(ctx context.Context, groupId uuid.UUID, id uuid.UUID) error {
	location, err := svc.repos.Locations.Get(ctx, id)
	if err != nil {
		return err
	}
	if location.Edges.Group.ID != groupId {
		return ErrNotOwner
	}

	return svc.repos.Locations.Delete(ctx, id)
}

func (svc *LocationService) Update(ctx context.Context, groupId uuid.UUID, data types.LocationUpdate) (*types.LocationOut, error) {
	location, err := svc.repos.Locations.Get(ctx, data.ID)
	if err != nil {
		return nil, err
	}
	if location.Edges.Group.ID != groupId {
		return nil, ErrNotOwner
	}

	return mappers.ToLocationOutErr(svc.repos.Locations.Update(ctx, data))
}
