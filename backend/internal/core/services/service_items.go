package services

import (
	"context"
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/core/services/reporting"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrFileNotFound = errors.New("file not found")
)

type ItemService struct {
	repo *repo.AllRepos

	filepath string

	autoIncrementAssetID bool
}

func (svc *ItemService) Create(ctx Context, item repo.ItemCreate) (repo.ItemOut, error) {
	if svc.autoIncrementAssetID {
		highest, err := svc.repo.Items.GetHighestAssetID(ctx, ctx.GID)
		if err != nil {
			return repo.ItemOut{}, err
		}

		item.AssetID = repo.AssetID(highest + 1)
	}

	return svc.repo.Items.Create(ctx, ctx.GID, item)
}

func (svc *ItemService) EnsureAssetID(ctx context.Context, GID uuid.UUID) (int, error) {
	items, err := svc.repo.Items.GetAllZeroAssetID(ctx, GID)
	if err != nil {
		return 0, err
	}

	highest, err := svc.repo.Items.GetHighestAssetID(ctx, GID)
	if err != nil {
		return 0, err
	}

	finished := 0
	for _, item := range items {
		highest++

		err = svc.repo.Items.SetAssetID(ctx, GID, item.ID, repo.AssetID(highest))
		if err != nil {
			return 0, err
		}

		finished++
	}

	return finished, nil
}

func (svc *ItemService) CsvImport(ctx context.Context, GID uuid.UUID, data io.Reader) (int, error) {
	// loaded, err := reporting.ReadCSV(data)
	// if err != nil {
	// 	return 0, err
	// }

	return 0, nil
}

func (svc *ItemService) ExportTSV(ctx context.Context, GID uuid.UUID) ([][]string, error) {
	items, err := svc.repo.Items.GetAll(ctx, GID)
	if err != nil {
		return nil, err
	}

	sheet := reporting.IOSheet{}

	sheet.ReadItems(items)

	return sheet.TSV()
}
