package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/internal/services/mappers"
	"github.com/hay-kot/homebox/backend/internal/types"
	"github.com/rs/zerolog/log"
)

var (
	ErrNotFound = errors.New("not found")
)

type ItemService struct {
	repo *repo.AllRepos

	filepath string
	// at is a map of tokens to attachment IDs. This is used to store the attachment ID
	// for issued URLs
	at attachmentTokens
}

func (svc *ItemService) GetOne(ctx context.Context, gid uuid.UUID, id uuid.UUID) (*types.ItemOut, error) {
	result, err := svc.repo.Items.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}

	if result.Edges.Group.ID != gid {
		return nil, ErrNotOwner
	}

	return mappers.ToItemOut(result), nil
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
	item, err := svc.repo.Items.GetOne(ctx, id)
	if err != nil {
		return err
	}

	if item.Edges.Group.ID != gid {
		return ErrNotOwner
	}

	err = svc.repo.Items.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (svc *ItemService) Update(ctx context.Context, gid uuid.UUID, data types.ItemUpdate) (*types.ItemOut, error) {
	item, err := svc.repo.Items.GetOne(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if item.Edges.Group.ID != gid {
		return nil, ErrNotOwner
	}

	item, err = svc.repo.Items.Update(ctx, data)
	if err != nil {
		return nil, err
	}

	return mappers.ToItemOut(item), nil
}

func (svc *ItemService) CsvImport(ctx context.Context, gid uuid.UUID, data [][]string) error {
	loaded := []csvRow{}

	// Skip first row
	for _, row := range data[1:] {
		// Skip empty rows
		if len(row) == 0 {
			continue
		}

		if len(row) != NumOfCols {
			return ErrInvalidCsv
		}

		r := newCsvRow(row)
		loaded = append(loaded, r)
	}

	// Bootstrap the locations and labels so we can reuse the created IDs for the items
	locations := map[string]uuid.UUID{}
	existingLocation, err := svc.repo.Locations.GetAll(ctx, gid)
	if err != nil {
		return err
	}
	for _, loc := range existingLocation {
		locations[loc.Name] = loc.ID
	}

	labels := map[string]uuid.UUID{}
	existingLabels, err := svc.repo.Labels.GetAll(ctx, gid)
	if err != nil {
		return err
	}
	for _, label := range existingLabels {
		labels[label.Name] = label.ID
	}

	for _, row := range loaded {

		// Locations
		if _, ok := locations[row.Location]; ok {
			continue
		}

		fmt.Println("Creating Location: ", row.Location)

		result, err := svc.repo.Locations.Create(ctx, gid, types.LocationCreate{
			Name:        row.Location,
			Description: "",
		})
		if err != nil {
			return err
		}
		locations[row.Location] = result.ID

		// Labels

		for _, label := range row.getLabels() {
			if _, ok := labels[label]; ok {
				continue
			}
			result, err := svc.repo.Labels.Create(ctx, gid, types.LabelCreate{
				Name:        label,
				Description: "",
			})
			if err != nil {
				return err
			}
			labels[label] = result.ID
		}
	}

	// Create the items
	for _, row := range loaded {
		locationID := locations[row.Location]
		labelIDs := []uuid.UUID{}
		for _, label := range row.getLabels() {
			labelIDs = append(labelIDs, labels[label])
		}

		log.Info().
			Str("name", row.Item.Name).
			Str("location", row.Location).
			Strs("labels", row.getLabels()).
			Str("locationId", locationID.String()).
			Msgf("Creating Item: %s", row.Item.Name)

		result, err := svc.repo.Items.Create(ctx, gid, types.ItemCreate{
			ImportRef:   row.Item.ImportRef,
			Name:        row.Item.Name,
			Description: row.Item.Description,
			LabelIDs:    labelIDs,
			LocationID:  locationID,
		})

		if err != nil {
			return err
		}

		// Update the item with the rest of the data
		_, err = svc.repo.Items.Update(ctx, types.ItemUpdate{
			// Edges
			LocationID: locationID,
			LabelIDs:   labelIDs,

			// General Fields
			ID:          result.ID,
			Name:        result.Name,
			Description: result.Description,
			Insured:     row.Item.Insured,
			Notes:       row.Item.Notes,

			// Identifies the item as imported
			SerialNumber: row.Item.SerialNumber,
			ModelNumber:  row.Item.ModelNumber,
			Manufacturer: row.Item.Manufacturer,

			// Purchase
			PurchaseFrom:  row.Item.PurchaseFrom,
			PurchasePrice: row.Item.PurchasePrice,
			PurchaseTime:  row.Item.PurchaseTime,

			// Warranty
			LifetimeWarranty: row.Item.LifetimeWarranty,
			WarrantyExpires:  row.Item.WarrantyExpires,
			WarrantyDetails:  row.Item.WarrantyDetails,

			SoldTo:    row.Item.SoldTo,
			SoldPrice: row.Item.SoldPrice,
			SoldTime:  row.Item.SoldTime,
			SoldNotes: row.Item.SoldNotes,
		})

		if err != nil {
			return err
		}
	}
	return nil
}
