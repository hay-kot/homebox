package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/ent/attachment"
	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/internal/services/mappers"
	"github.com/hay-kot/homebox/backend/internal/types"
	"github.com/hay-kot/homebox/backend/pkgs/hasher"
	"github.com/hay-kot/homebox/backend/pkgs/pathlib"
	"github.com/rs/zerolog/log"
)

var (
	ErrNotFound = errors.New("not found")
)

// TODO: this isn't a scalable solution, tokens should be stored in the database
type attachmentTokens map[string]uuid.UUID

func (at attachmentTokens) Add(token string, id uuid.UUID) {
	at[token] = id

	log.Debug().Str("token", token).Str("uuid", id.String()).Msg("added token")

	go func() {
		ch := time.After(1 * time.Minute)
		<-ch
		at.Delete(token)
		log.Debug().Str("token", token).Msg("deleted token")
	}()
}

func (at attachmentTokens) Get(token string) (uuid.UUID, bool) {
	id, ok := at[token]
	return id, ok
}

func (at attachmentTokens) Delete(token string) {
	delete(at, token)
}

type ItemService struct {
	repo *repo.AllRepos

	// filepath is the root of the storage location that will be used to store all files from.
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

func (svc *ItemService) attachmentPath(gid, itemId uuid.UUID, filename string) string {
	path := filepath.Join(svc.filepath, gid.String(), itemId.String(), filename)
	return pathlib.Safe(path)
}

func (svc *ItemService) NewAttachmentToken(ctx context.Context, gid, itemId, attachmentId uuid.UUID) (string, error) {
	item, err := svc.repo.Items.GetOne(ctx, itemId)
	if err != nil {
		return "", err
	}
	if item.Edges.Group.ID != gid {
		return "", ErrNotOwner
	}

	token := hasher.GenerateToken()

	svc.at.Add(token.Raw, attachmentId)

	return token.Raw, nil
}

func (svc *ItemService) GetAttachment(ctx context.Context, token string) (string, error) {
	attachmentId, ok := svc.at.Get(token)
	if !ok {
		return "", ErrNotFound
	}

	attachment, err := svc.repo.Attachments.Get(ctx, attachmentId)
	if err != nil {
		return "", err
	}

	return attachment.Edges.Document.Path, nil
}

// AddAttachment adds an attachment to an item by creating an entry in the Documents table and linking it to the Attachment
// Table and Items table. The file provided via the reader is stored on the file system based on the provided
// relative path during construction of the service.
func (svc *ItemService) AddAttachment(ctx context.Context, gid, itemId uuid.UUID, filename string, attachmentType attachment.Type, file io.Reader) (*types.ItemOut, error) {
	// Get the Item
	item, err := svc.repo.Items.GetOne(ctx, itemId)
	if err != nil {
		return nil, err
	}

	if item.Edges.Group.ID != gid {
		return nil, ErrNotOwner
	}

	// Create the document
	doc, err := svc.repo.Docs.Create(ctx, gid, types.DocumentCreate{
		Title: filename,
		Path:  svc.attachmentPath(gid, itemId, filename),
	})
	if err != nil {
		return nil, err
	}

	// Create the attachment
	_, err = svc.repo.Attachments.Create(ctx, itemId, doc.ID, attachmentType)
	if err != nil {
		return nil, err
	}

	// Read the contents and write them to a file on the file system
	err = os.MkdirAll(filepath.Dir(doc.Path), os.ModePerm)
	if err != nil {
		return nil, err
	}

	f, err := os.Create(doc.Path)
	if err != nil {
		log.Err(err).Msg("failed to create file")
		return nil, err
	}

	_, err = io.Copy(f, file)
	if err != nil {
		return nil, err
	}

	return svc.GetOne(ctx, gid, itemId)
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
