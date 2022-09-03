package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/internal/repo"
	"github.com/hay-kot/content/backend/internal/services/mappers"
	"github.com/hay-kot/content/backend/internal/types"
)

type ItemService struct {
	repo *repo.AllRepos
}

func (svc *ItemService) GetOne(ctx context.Context, gid uuid.UUID, id uuid.UUID) (*types.ItemOut, error) {
	panic("implement me")
}
func (svc *ItemService) GetAll(ctx context.Context, gid uuid.UUID) ([]*types.ItemSummary, error) {
	items, err := svc.repo.Items.GetAll(ctx, gid)
	if err != nil {
		return nil, err
	}

	itemsOut := make([]*types.ItemSummary, len(items))
	for i, item := range items {
		itemsOut[i] = mappers.ToItemSummary(item)
	}

	return itemsOut, nil
}
func (svc *ItemService) Create(ctx context.Context, gid uuid.UUID, data types.ItemCreate) (*types.ItemOut, error) {
	item, err := svc.repo.Items.Create(ctx, gid, data)
	if err != nil {
		return nil, err
	}

	return mappers.ToItemOut(item), nil
}
func (svc *ItemService) Delete(ctx context.Context, gid uuid.UUID, id uuid.UUID) error {
	panic("implement me")
}
func (svc *ItemService) Update(ctx context.Context, gid uuid.UUID, data types.ItemUpdate) (*types.ItemOut, error) {
	panic("implement me")
}
