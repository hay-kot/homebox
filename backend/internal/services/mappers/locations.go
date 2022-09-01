package mappers

import (
	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/internal/types"
)

func ToLocationSummary(location *ent.Location) *types.LocationSummary {
	return &types.LocationSummary{
		ID:          location.ID,
		GroupID:     location.Edges.Group.ID,
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
		LocationSummary: *ToLocationSummary(location),
		Items:           MapEach(location.Edges.Items, ToItemSummary),
	}
}

func ToLocationOutErr(location *ent.Location, err error) (*types.LocationOut, error) {
	return ToLocationOut(location), err
}
