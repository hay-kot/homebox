package mappers

import (
	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/internal/types"
)

func ToItemSummary(item *ent.Item) *types.ItemSummary {
	var location *types.LocationSummary
	if item.Edges.Location != nil {
		location = ToLocationSummary(item.Edges.Location)
	}

	var labels []*types.LabelSummary
	if item.Edges.Label != nil {
		labels = MapEach(item.Edges.Label, ToLabelSummary)
	}

	return &types.ItemSummary{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,

		// Edges
		Location: location,
		Labels:   labels,

		// Identification
		SerialNumber: item.SerialNumber,
		ModelNumber:  item.ModelNumber,
		Manufacturer: item.Manufacturer,

		// Purchase
		PurchaseTime:  item.PurchaseTime,
		PurchaseFrom:  item.PurchaseFrom,
		PurchasePrice: item.PurchasePrice,

		// Sold
		SoldTime:  item.SoldTime,
		SoldTo:    item.SoldTo,
		SoldPrice: item.SoldPrice,
		SoldNotes: item.SoldNotes,

		// Extras
		Notes: item.Notes,
	}
}

func ToItemSummaryErr(item *ent.Item, err error) (*types.ItemSummary, error) {
	return ToItemSummary(item), err
}

func ToItemOut(item *ent.Item) *types.ItemOut {
	return &types.ItemOut{
		ItemSummary: *ToItemSummary(item),
	}
}

func ToItemOutErr(item *ent.Item, err error) (*types.ItemOut, error) {
	return ToItemOut(item), err
}
