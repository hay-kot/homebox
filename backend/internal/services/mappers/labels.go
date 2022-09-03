package mappers

import (
	"github.com/hay-kot/content/backend/ent"
	"github.com/hay-kot/content/backend/internal/types"
)

func ToLabelSummary(label *ent.Label) *types.LabelSummary {
	return &types.LabelSummary{
		ID:          label.ID,
		Name:        label.Name,
		Description: label.Description,
		CreatedAt:   label.CreatedAt,
		UpdatedAt:   label.UpdatedAt,
	}
}

func ToLabelSummaryErr(label *ent.Label, err error) (*types.LabelSummary, error) {
	return ToLabelSummary(label), err
}

func ToLabelOut(label *ent.Label) *types.LabelOut {
	return &types.LabelOut{
		LabelSummary: *ToLabelSummary(label),
		Items:        MapEach(label.Edges.Items, ToItemSummary),
	}
}

func ToLabelOutErr(label *ent.Label, err error) (*types.LabelOut, error) {
	return ToLabelOut(label), err
}
