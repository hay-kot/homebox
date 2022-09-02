package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/internal/repo"
	"github.com/hay-kot/content/backend/internal/services/mappers"
	"github.com/hay-kot/content/backend/internal/types"
)

type LabelService struct {
	repos *repo.AllRepos
}

func (svc *LabelService) Create(ctx context.Context, groupId uuid.UUID, data types.LabelCreate) (*types.LabelSummary, error) {
	label, err := svc.repos.Labels.Create(ctx, groupId, data)
	return mappers.ToLabelSummaryErr(label, err)
}

func (svc *LabelService) Update(ctx context.Context, groupId uuid.UUID, data types.LabelUpdate) (*types.LabelSummary, error) {
	label, err := svc.repos.Labels.Update(ctx, data)
	return mappers.ToLabelSummaryErr(label, err)
}

func (svc *LabelService) Delete(ctx context.Context, groupId uuid.UUID, id uuid.UUID) error {
	label, err := svc.repos.Labels.Get(ctx, id)
	if err != nil {
		return err
	}
	if label.Edges.Group.ID != groupId {
		return ErrNotOwner
	}
	return svc.repos.Labels.Delete(ctx, id)
}

func (svc *LabelService) Get(ctx context.Context, groupId uuid.UUID, id uuid.UUID) (*types.LabelOut, error) {
	label, err := svc.repos.Labels.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	if label.Edges.Group.ID != groupId {
		return nil, ErrNotOwner
	}

	return mappers.ToLabelOut(label), nil
}

func (svc *LabelService) GetAll(ctx context.Context, groupId uuid.UUID) ([]*types.LabelSummary, error) {
	labels, err := svc.repos.Labels.GetAll(ctx, groupId)
	if err != nil {
		return nil, err
	}

	labelsOut := make([]*types.LabelSummary, len(labels))
	for i, label := range labels {
		labelsOut[i] = mappers.ToLabelSummary(label)
	}

	return labelsOut, nil
}
