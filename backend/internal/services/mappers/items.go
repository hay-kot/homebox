package mappers

import (
	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/internal/types"
)

func ToItemSummary(item *ent.Item) *types.ItemSummary {
	return &types.ItemSummary{
		ID:          item.ID,
		LocationID:  item.Edges.Location.ID,
		Name:        item.Name,
		Description: item.Description,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}
