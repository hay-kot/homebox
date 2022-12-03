package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/repo"
	"github.com/rs/zerolog/log"
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
func (svc *ItemService) CsvImport(ctx context.Context, GID uuid.UUID, data [][]string) (int, error) {
	loaded := []csvRow{}

	// Skip first row
	for _, row := range data[1:] {
		// Skip empty rows
		if len(row) == 0 {
			continue
		}

		if len(row) != NumOfCols {
			return 0, ErrInvalidCsv
		}

		r := newCsvRow(row)
		loaded = append(loaded, r)
	}

	// validate rows
	var errMap = map[int][]error{}
	var hasErr bool
	for i, r := range loaded {

		errs := r.validate()

		if len(errs) > 0 {
			hasErr = true
			lineNum := i + 2

			errMap[lineNum] = errs
		}
	}

	if hasErr {
		for lineNum, errs := range errMap {
			for _, err := range errs {
				log.Error().Err(err).Int("line", lineNum).Msg("csv import error")
			}
		}
	}

	// Bootstrap the locations and labels so we can reuse the created IDs for the items
	locations := map[string]uuid.UUID{}
	existingLocation, err := svc.repo.Locations.GetAll(ctx, GID, repo.LocationQuery{})
	if err != nil {
		return 0, err
	}
	for _, loc := range existingLocation {
		locations[loc.Name] = loc.ID
	}

	labels := map[string]uuid.UUID{}
	existingLabels, err := svc.repo.Labels.GetAll(ctx, GID)
	if err != nil {
		return 0, err
	}
	for _, label := range existingLabels {
		labels[label.Name] = label.ID
	}

	for _, row := range loaded {

		// Locations
		if _, exists := locations[row.Location]; !exists {
			result, err := svc.repo.Locations.Create(ctx, GID, repo.LocationCreate{
				Name:        row.Location,
				Description: "",
			})
			if err != nil {
				return 0, err
			}
			locations[row.Location] = result.ID
		}

		// Labels

		for _, label := range row.getLabels() {
			if _, exists := labels[label]; exists {
				continue
			}
			result, err := svc.repo.Labels.Create(ctx, GID, repo.LabelCreate{
				Name:        label,
				Description: "",
			})
			if err != nil {
				return 0, err
			}
			labels[label] = result.ID
		}
	}

	highest := repo.AssetID(-1)
	if svc.autoIncrementAssetID {
		highest, err = svc.repo.Items.GetHighestAssetID(ctx, GID)
		if err != nil {
			return 0, err
		}
	}

	// Create the items
	var count int
	for _, row := range loaded {
		// Check Import Ref
		if row.Item.ImportRef != "" {
			exists, err := svc.repo.Items.CheckRef(ctx, GID, row.Item.ImportRef)
			if exists {
				continue
			}
			if err != nil {
				log.Err(err).Msg("error checking import ref")
			}
		}

		locationID := locations[row.Location]
		labelIDs := []uuid.UUID{}
		for _, label := range row.getLabels() {
			labelIDs = append(labelIDs, labels[label])
		}

		log.Info().
			Str("name", row.Item.Name).
			Str("location", row.Location).
			Msgf("Creating Item: %s", row.Item.Name)

		data := repo.ItemCreate{
			ImportRef:   row.Item.ImportRef,
			Name:        row.Item.Name,
			Description: row.Item.Description,
			LabelIDs:    labelIDs,
			LocationID:  locationID,
		}

		if svc.autoIncrementAssetID {
			highest++
			data.AssetID = highest
		}

		result, err := svc.repo.Items.Create(ctx, GID, data)

		if err != nil {
			return count, err
		}

		// Update the item with the rest of the data
		_, err = svc.repo.Items.UpdateByGroup(ctx, GID, repo.ItemUpdate{
			// Edges
			LocationID: locationID,
			LabelIDs:   labelIDs,
			AssetID:    data.AssetID,

			// General Fields
			ID:          result.ID,
			Name:        result.Name,
			Description: result.Description,
			Insured:     row.Item.Insured,
			Notes:       row.Item.Notes,
			Quantity:    row.Item.Quantity,

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
			return count, err
		}

		count++
	}
	return count, nil
}
