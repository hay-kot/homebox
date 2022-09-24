package services

import "github.com/hay-kot/homebox/backend/internal/types"

func defaultLocations() []types.LocationCreate {
	return []types.LocationCreate{
		{
			Name: "Living Room",
		},
		{
			Name: "Garage",
		},
		{
			Name: "Kitchen",
		},
		{
			Name: "Bedroom",
		},
		{
			Name: "Bathroom",
		},
		{
			Name: "Office",
		},
		{
			Name: "Attic",
		},
		{
			Name: "Basement",
		},
	}
}

func defaultLabels() []types.LabelCreate {
	return []types.LabelCreate{
		{
			Name: "Appliances",
		},
		{
			Name: "IOT",
		},
		{
			Name: "Electronics",
		},
		{
			Name: "Servers",
		},
		{
			Name: "General",
		},
		{
			Name: "Important",
		},
	}
}
