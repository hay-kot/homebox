package mappers

import (
	"github.com/hay-kot/homebox/backend/ent"
	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/internal/types"
)

func ToLocationCount(location *repo.LocationWithCount) *types.LocationCount {
	return &types.LocationCount{
		LocationSummary: types.LocationSummary{
			ID:          location.ID,
			Name:        location.Name,
			Description: location.Description,
			CreatedAt:   location.CreatedAt,
			UpdatedAt:   location.UpdatedAt,
		},
		ItemCount: location.ItemCount,
	}
}

func ToLocationCountErr(location *repo.LocationWithCount, err error) (*types.LocationCount, error) {
	return ToLocationCount(location), err
}

func ToLocationSummary(location *ent.Location) *types.LocationSummary {
	return &types.LocationSummary{
		ID:          location.ID,
		Name:        location.Name,
		Description: location.Description,
		CreatedAt:   location.CreatedAt,
		UpdatedAt:   location.UpdatedAt,
	}
}

func ToLocationSummaryErr(location *ent.Location, err error) (*types.LocationSummary, error) {
	return ToLocationSummary(location), err
}

func ToLocationOut(location *ent.Location) *types.LocationOut {
	return &types.LocationOut{
		LocationSummary: types.LocationSummary{
			ID:          location.ID,
			Name:        location.Name,
			Description: location.Description,
			CreatedAt:   location.CreatedAt,
			UpdatedAt:   location.UpdatedAt,
		},
		Items: MapEach(location.Edges.Items, ToItemSummary),
	}
}

func ToLocationOutErr(location *ent.Location, err error) (*types.LocationOut, error) {
	return ToLocationOut(location), err
}
