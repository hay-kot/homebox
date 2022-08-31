package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/internal/repo"
	"github.com/hay-kot/content/backend/internal/types"
)

type LocationService struct {
	repos *repo.AllRepos
}

func ToLocationOut(location *ent.Location, err error) (*types.LocationOut, error) {
	return &types.LocationOut{
		ID:          location.ID,
		GroupID:     location.Edges.Group.ID,
		Name:        location.Name,
		Description: location.Description,
		CreatedAt:   location.CreatedAt,
		UpdatedAt:   location.UpdatedAt,
	}, err
}

func (svc *LocationService) Create(ctx context.Context, groupId uuid.UUID, data types.LocationCreate) (*types.LocationOut, error) {
	location, err := svc.repos.Locations.Create(ctx, groupId, data)
	return ToLocationOut(location, err)
}

func (svc *LocationService) GetAll(ctx context.Context, groupId uuid.UUID) ([]*types.LocationOut, error) {
	locations, err := svc.repos.Locations.GetAll(ctx, groupId)
	if err != nil {
		return nil, err
	}

	locationsOut := make([]*types.LocationOut, len(locations))
	for i, location := range locations {
		locationOut, _ := ToLocationOut(location, nil)
		locationsOut[i] = locationOut
	}

	return locationsOut, nil
}
