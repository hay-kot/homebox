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

func (svc *LocationService) GetAll(ctx context.Context, groupId uuid.UUID) ([]*types.LocationOut, error) {
	panic("not implemented")
}
