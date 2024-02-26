package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/core/services/reporting/eventbus"
	"github.com/hay-kot/homebox/backend/internal/data/ent"
	"github.com/hay-kot/homebox/backend/internal/data/ent/attachment"
	"github.com/hay-kot/homebox/backend/internal/data/ent/group"
	"github.com/hay-kot/homebox/backend/internal/data/ent/item"
	"github.com/hay-kot/homebox/backend/internal/data/ent/itemfield"
	"github.com/hay-kot/homebox/backend/internal/data/ent/label"
	"github.com/hay-kot/homebox/backend/internal/data/ent/location"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
	"github.com/hay-kot/homebox/backend/internal/data/types"
)

type ItemsRepository struct {
	db  *ent.Client
	bus *eventbus.EventBus
}

type (
	FieldQuery struct {
		Name  string
		Value string
	}

	ItemQuery struct {
		Page            int
		PageSize        int
		Search          string       `json:"search"`
		AssetID         AssetID      `json:"assetId"`
		LocationIDs     []uuid.UUID  `json:"locationIds"`
		LabelIDs        []uuid.UUID  `json:"labelIds"`
		ParentItemIDs   []uuid.UUID  `json:"parentIds"`
		SortBy          string       `json:"sortBy"`
		IncludeArchived bool         `json:"includeArchived"`
		Fields          []FieldQuery `json:"fields"`
		OrderBy         string       `json:"orderBy"`
	}

	ItemField struct {
		ID           uuid.UUID `json:"id,omitempty"`
		Type         string    `json:"type"`
		Name         string    `json:"name"`
		TextValue    string    `json:"textValue"`
		NumberValue  int       `json:"numberValue"`
		BooleanValue bool      `json:"booleanValue"`
		// TimeValue    time.Time `json:"timeValue,omitempty"`
	}

	ItemCreate struct {
		ImportRef   string    `json:"-"`
		ParentID    uuid.UUID `json:"parentId"    extensions:"x-nullable"`
		Name        string    `json:"name"        validate:"required,min=1,max=255"`
		Description string    `json:"description" validate:"max=1000"`
		AssetID     AssetID   `json:"-"`

		// Edges
		LocationID uuid.UUID   `json:"locationId"`
		LabelIDs   []uuid.UUID `json:"labelIds"`
	}

	ItemUpdate struct {
		ParentID    uuid.UUID `json:"parentId"    extensions:"x-nullable,x-omitempty"`
		ID          uuid.UUID `json:"id"`
		AssetID     AssetID   `json:"assetId"     swaggertype:"string"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Quantity    int       `json:"quantity"`
		Insured     bool      `json:"insured"`
		Archived    bool      `json:"archived"`

		// Edges
		LocationID uuid.UUID   `json:"locationId"`
		LabelIDs   []uuid.UUID `json:"labelIds"`

		// Identifications
		SerialNumber string `json:"serialNumber"`
		ModelNumber  string `json:"modelNumber"`
		Manufacturer string `json:"manufacturer"`

		// Warranty
		LifetimeWarranty bool       `json:"lifetimeWarranty"`
		WarrantyExpires  types.Date `json:"warrantyExpires"`
		WarrantyDetails  string     `json:"warrantyDetails"`

		// Purchase
		PurchaseTime  types.Date `json:"purchaseTime"`
		PurchaseFrom  string     `json:"purchaseFrom"`
		PurchasePrice float64    `json:"purchasePrice,string"`

		// Sold
		SoldTime  types.Date `json:"soldTime"`
		SoldTo    string     `json:"soldTo"`
		SoldPrice float64    `json:"soldPrice,string"`
		SoldNotes string     `json:"soldNotes"`

		// Extras
		Notes  string      `json:"notes"`
		Fields []ItemField `json:"fields"`
	}

	ItemPatch struct {
		ID        uuid.UUID `json:"id"`
		Quantity  *int      `json:"quantity,omitempty" extensions:"x-nullable,x-omitempty"`
		ImportRef *string   `json:"-,omitempty"        extensions:"x-nullable,x-omitempty"`
	}

	ItemSummary struct {
		ImportRef   string    `json:"-"`
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Quantity    int       `json:"quantity"`
		Insured     bool      `json:"insured"`
		Archived    bool      `json:"archived"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`

		PurchasePrice float64 `json:"purchasePrice,string"`

		// Edges
		Location *LocationSummary `json:"location,omitempty" extensions:"x-nullable,x-omitempty"`
		Labels   []LabelSummary   `json:"labels"`

		ImageID *uuid.UUID `json:"imageId,omitempty"`
	}

	ItemOut struct {
		Parent *ItemSummary `json:"parent,omitempty" extensions:"x-nullable,x-omitempty"`
		ItemSummary
		AssetID AssetID `json:"assetId,string"`

		SerialNumber string `json:"serialNumber"`
		ModelNumber  string `json:"modelNumber"`
		Manufacturer string `json:"manufacturer"`

		// Warranty
		LifetimeWarranty bool       `json:"lifetimeWarranty"`
		WarrantyExpires  types.Date `json:"warrantyExpires"`
		WarrantyDetails  string     `json:"warrantyDetails"`

		// Purchase
		PurchaseTime types.Date `json:"purchaseTime"`
		PurchaseFrom string     `json:"purchaseFrom"`

		// Sold
		SoldTime  types.Date `json:"soldTime"`
		SoldTo    string     `json:"soldTo"`
		SoldPrice float64    `json:"soldPrice,string"`
		SoldNotes string     `json:"soldNotes"`

		// Extras
		Notes string `json:"notes"`

		Attachments []ItemAttachment `json:"attachments"`
		Fields      []ItemField      `json:"fields"`
	}
)

var mapItemsSummaryErr = mapTEachErrFunc(mapItemSummary)

func mapItemSummary(item *ent.Item) ItemSummary {
	var location *LocationSummary
	if item.Edges.Location != nil {
		loc := mapLocationSummary(item.Edges.Location)
		location = &loc
	}

	labels := make([]LabelSummary, len(item.Edges.Label))
	if item.Edges.Label != nil {
		labels = mapEach(item.Edges.Label, mapLabelSummary)
	}

	var imageID *uuid.UUID
	if item.Edges.Attachments != nil {
		for _, a := range item.Edges.Attachments {
			if a.Primary && a.Edges.Document != nil {
				imageID = &a.ID
				break
			}
		}
	}

	return ItemSummary{
		ID:            item.ID,
		Name:          item.Name,
		Description:   item.Description,
		ImportRef:     item.ImportRef,
		Quantity:      item.Quantity,
		CreatedAt:     item.CreatedAt,
		UpdatedAt:     item.UpdatedAt,
		Archived:      item.Archived,
		PurchasePrice: item.PurchasePrice,

		// Edges
		Location: location,
		Labels:   labels,

		// Warranty
		Insured: item.Insured,
		ImageID: imageID,
	}
}

var (
	mapItemOutErr  = mapTErrFunc(mapItemOut)
	mapItemsOutErr = mapTEachErrFunc(mapItemOut)
)

func mapFields(fields []*ent.ItemField) []ItemField {
	result := make([]ItemField, len(fields))
	for i, f := range fields {
		result[i] = ItemField{
			ID:           f.ID,
			Type:         f.Type.String(),
			Name:         f.Name,
			TextValue:    f.TextValue,
			NumberValue:  f.NumberValue,
			BooleanValue: f.BooleanValue,
			// TimeValue:    f.TimeValue,
		}
	}
	return result
}

func mapItemOut(item *ent.Item) ItemOut {
	var attachments []ItemAttachment
	if item.Edges.Attachments != nil {
		attachments = mapEach(item.Edges.Attachments, ToItemAttachment)
	}

	var fields []ItemField
	if item.Edges.Fields != nil {
		fields = mapFields(item.Edges.Fields)
	}

	var parent *ItemSummary
	if item.Edges.Parent != nil {
		v := mapItemSummary(item.Edges.Parent)
		parent = &v
	}

	return ItemOut{
		Parent:           parent,
		AssetID:          AssetID(item.AssetID),
		ItemSummary:      mapItemSummary(item),
		LifetimeWarranty: item.LifetimeWarranty,
		WarrantyExpires:  types.DateFromTime(item.WarrantyExpires),
		WarrantyDetails:  item.WarrantyDetails,

		// Identification
		SerialNumber: item.SerialNumber,
		ModelNumber:  item.ModelNumber,
		Manufacturer: item.Manufacturer,

		// Purchase
		PurchaseTime: types.DateFromTime(item.PurchaseTime),
		PurchaseFrom: item.PurchaseFrom,

		// Sold
		SoldTime:  types.DateFromTime(item.SoldTime),
		SoldTo:    item.SoldTo,
		SoldPrice: item.SoldPrice,
		SoldNotes: item.SoldNotes,

		// Extras
		Notes:       item.Notes,
		Attachments: attachments,
		Fields:      fields,
	}
}

func (e *ItemsRepository) publishMutationEvent(GID uuid.UUID) {
	if e.bus != nil {
		e.bus.Publish(eventbus.EventItemMutation, eventbus.GroupMutationEvent{GID: GID})
	}
}

func (e *ItemsRepository) getOne(ctx context.Context, where ...predicate.Item) (ItemOut, error) {
	q := e.db.Item.Query().Where(where...)

	return mapItemOutErr(q.
		WithFields().
		WithLabel().
		WithLocation().
		WithGroup().
		WithParent().
		WithAttachments(func(aq *ent.AttachmentQuery) {
			aq.WithDocument()
		}).
		Only(ctx),
	)
}

// GetOne returns a single item by ID. If the item does not exist, an error is returned.
// See also: GetOneByGroup to ensure that the item belongs to a specific group.
func (e *ItemsRepository) GetOne(ctx context.Context, id uuid.UUID) (ItemOut, error) {
	return e.getOne(ctx, item.ID(id))
}

func (e *ItemsRepository) CheckRef(ctx context.Context, GID uuid.UUID, ref string) (bool, error) {
	q := e.db.Item.Query().Where(item.HasGroupWith(group.ID(GID)))
	return q.Where(item.ImportRef(ref)).Exist(ctx)
}

func (e *ItemsRepository) GetByRef(ctx context.Context, GID uuid.UUID, ref string) (ItemOut, error) {
	return e.getOne(ctx, item.ImportRef(ref), item.HasGroupWith(group.ID(GID)))
}

// GetOneByGroup returns a single item by ID. If the item does not exist, an error is returned.
// GetOneByGroup ensures that the item belongs to a specific group.
func (e *ItemsRepository) GetOneByGroup(ctx context.Context, gid, id uuid.UUID) (ItemOut, error) {
	return e.getOne(ctx, item.ID(id), item.HasGroupWith(group.ID(gid)))
}

// QueryByGroup returns a list of items that belong to a specific group based on the provided query.
func (e *ItemsRepository) QueryByGroup(ctx context.Context, gid uuid.UUID, q ItemQuery) (PaginationResult[ItemSummary], error) {
	qb := e.db.Item.Query().Where(
		item.HasGroupWith(group.ID(gid)),
	)

	if q.IncludeArchived {
		qb = qb.Where(
			item.Or(
				item.Archived(true),
				item.Archived(false),
			),
		)
	} else {
		qb = qb.Where(item.Archived(false))
	}

	if q.Search != "" {
		qb.Where(
			item.Or(
				item.NameContainsFold(q.Search),
				item.DescriptionContainsFold(q.Search),
				item.SerialNumberContainsFold(q.Search),
				item.ModelNumberContainsFold(q.Search),
				item.ManufacturerContainsFold(q.Search),
				item.NotesContainsFold(q.Search),
			),
		)
	}

	if !q.AssetID.Nil() {
		qb = qb.Where(item.AssetID(q.AssetID.Int()))
	}

	// Filters within this block define a AND relationship where each subset
	// of filters is OR'd together.
	//
	// The goal is to allow matches like where the item has
	//  - one of the selected labels AND
	//  - one of the selected locations AND
	//  - one of the selected fields key/value matches
	var andPredicates []predicate.Item
	{
		if len(q.LabelIDs) > 0 {
			labelPredicates := make([]predicate.Item, 0, len(q.LabelIDs))
			for _, l := range q.LabelIDs {
				labelPredicates = append(labelPredicates, item.HasLabelWith(label.ID(l)))
			}

			andPredicates = append(andPredicates, item.Or(labelPredicates...))
		}

		if len(q.LocationIDs) > 0 {
			locationPredicates := make([]predicate.Item, 0, len(q.LocationIDs))
			for _, l := range q.LocationIDs {
				locationPredicates = append(locationPredicates, item.HasLocationWith(location.ID(l)))
			}

			andPredicates = append(andPredicates, item.Or(locationPredicates...))
		}

		if len(q.Fields) > 0 {
			fieldPredicates := make([]predicate.Item, 0, len(q.Fields))
			for _, f := range q.Fields {
				fieldPredicates = append(fieldPredicates, item.HasFieldsWith(
					itemfield.And(
						itemfield.Name(f.Name),
						itemfield.TextValue(f.Value),
					),
				))
			}

			andPredicates = append(andPredicates, item.Or(fieldPredicates...))
		}

		if len(q.ParentItemIDs) > 0 {
			andPredicates = append(andPredicates, item.HasParentWith(item.IDIn(q.ParentItemIDs...)))
		}
	}

	if len(andPredicates) > 0 {
		qb = qb.Where(item.And(andPredicates...))
	}

	count, err := qb.Count(ctx)
	if err != nil {
		return PaginationResult[ItemSummary]{}, err
	}

	// Order
	switch q.OrderBy {
	case "createdAt":
		qb = qb.Order(ent.Desc(item.FieldCreatedAt))
	case "updatedAt":
		qb = qb.Order(ent.Desc(item.FieldUpdatedAt))
	default: // "name"
		qb = qb.Order(ent.Asc(item.FieldName))
	}

	qb = qb.
		WithLabel().
		WithLocation().
		WithAttachments(func(aq *ent.AttachmentQuery) {
			aq.Where(
				attachment.Primary(true),
			).
				WithDocument()
		})

	if q.Page != -1 || q.PageSize != -1 {
		qb = qb.
			Offset(calculateOffset(q.Page, q.PageSize)).
			Limit(q.PageSize)
	}

	items, err := mapItemsSummaryErr(qb.All(ctx))
	if err != nil {
		return PaginationResult[ItemSummary]{}, err
	}

	return PaginationResult[ItemSummary]{
		Page:     q.Page,
		PageSize: q.PageSize,
		Total:    count,
		Items:    items,
	}, nil
}

// QueryByAssetID returns items by asset ID. If the item does not exist, an error is returned.
func (e *ItemsRepository) QueryByAssetID(ctx context.Context, gid uuid.UUID, assetID AssetID, page int, pageSize int) (PaginationResult[ItemSummary], error) {
	qb := e.db.Item.Query().Where(
		item.HasGroupWith(group.ID(gid)),
		item.AssetID(int(assetID)),
	)

	if page != -1 || pageSize != -1 {
		qb.Offset(calculateOffset(page, pageSize)).
			Limit(pageSize)
	} else {
		page = -1
		pageSize = -1
	}

	items, err := mapItemsSummaryErr(
		qb.Order(ent.Asc(item.FieldName)).
			WithLabel().
			WithLocation().
			All(ctx),
	)
	if err != nil {
		return PaginationResult[ItemSummary]{}, err
	}

	return PaginationResult[ItemSummary]{
		Page:     page,
		PageSize: pageSize,
		Total:    len(items),
		Items:    items,
	}, nil
}

// GetAll returns all the items in the database with the Labels and Locations eager loaded.
func (e *ItemsRepository) GetAll(ctx context.Context, gid uuid.UUID) ([]ItemOut, error) {
	return mapItemsOutErr(e.db.Item.Query().
		Where(item.HasGroupWith(group.ID(gid))).
		WithLabel().
		WithLocation().
		WithFields().
		All(ctx))
}

func (e *ItemsRepository) GetAllZeroAssetID(ctx context.Context, GID uuid.UUID) ([]ItemSummary, error) {
	q := e.db.Item.Query().Where(
		item.HasGroupWith(group.ID(GID)),
		item.AssetID(0),
	).Order(
		ent.Asc(item.FieldCreatedAt),
	)

	return mapItemsSummaryErr(q.All(ctx))
}

func (e *ItemsRepository) GetHighestAssetID(ctx context.Context, GID uuid.UUID) (AssetID, error) {
	q := e.db.Item.Query().Where(
		item.HasGroupWith(group.ID(GID)),
	).Order(
		ent.Desc(item.FieldAssetID),
	).Limit(1)

	result, err := q.First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return 0, nil
		}
		return 0, err
	}

	return AssetID(result.AssetID), nil
}

func (e *ItemsRepository) SetAssetID(ctx context.Context, GID uuid.UUID, ID uuid.UUID, assetID AssetID) error {
	q := e.db.Item.Update().Where(
		item.HasGroupWith(group.ID(GID)),
		item.ID(ID),
	)

	_, err := q.SetAssetID(int(assetID)).Save(ctx)
	return err
}

func (e *ItemsRepository) Create(ctx context.Context, gid uuid.UUID, data ItemCreate) (ItemOut, error) {
	q := e.db.Item.Create().
		SetImportRef(data.ImportRef).
		SetName(data.Name).
		SetDescription(data.Description).
		SetGroupID(gid).
		SetLocationID(data.LocationID).
		SetAssetID(int(data.AssetID))

	if data.LabelIDs != nil && len(data.LabelIDs) > 0 {
		q.AddLabelIDs(data.LabelIDs...)
	}

	result, err := q.Save(ctx)
	if err != nil {
		return ItemOut{}, err
	}

	e.publishMutationEvent(gid)
	return e.GetOne(ctx, result.ID)
}

func (e *ItemsRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := e.db.Item.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}

	e.publishMutationEvent(id)
	return nil
}

func (e *ItemsRepository) DeleteByGroup(ctx context.Context, gid, id uuid.UUID) error {
	_, err := e.db.Item.
		Delete().
		Where(
			item.ID(id),
			item.HasGroupWith(group.ID(gid)),
		).Exec(ctx)
	if err != nil {
		return err
	}

	e.publishMutationEvent(gid)
	return err
}

func (e *ItemsRepository) UpdateByGroup(ctx context.Context, GID uuid.UUID, data ItemUpdate) (ItemOut, error) {
	q := e.db.Item.Update().Where(item.ID(data.ID), item.HasGroupWith(group.ID(GID))).
		SetName(data.Name).
		SetDescription(data.Description).
		SetLocationID(data.LocationID).
		SetSerialNumber(data.SerialNumber).
		SetModelNumber(data.ModelNumber).
		SetManufacturer(data.Manufacturer).
		SetArchived(data.Archived).
		SetPurchaseTime(data.PurchaseTime.Time()).
		SetPurchaseFrom(data.PurchaseFrom).
		SetPurchasePrice(data.PurchasePrice).
		SetSoldTime(data.SoldTime.Time()).
		SetSoldTo(data.SoldTo).
		SetSoldPrice(data.SoldPrice).
		SetSoldNotes(data.SoldNotes).
		SetNotes(data.Notes).
		SetLifetimeWarranty(data.LifetimeWarranty).
		SetInsured(data.Insured).
		SetWarrantyExpires(data.WarrantyExpires.Time()).
		SetWarrantyDetails(data.WarrantyDetails).
		SetQuantity(data.Quantity).
		SetAssetID(int(data.AssetID))

	currentLabels, err := e.db.Item.Query().Where(item.ID(data.ID)).QueryLabel().All(ctx)
	if err != nil {
		return ItemOut{}, err
	}

	set := newIDSet(currentLabels)

	for _, l := range data.LabelIDs {
		if set.Contains(l) {
			set.Remove(l)
			continue
		}
		q.AddLabelIDs(l)
	}

	if set.Len() > 0 {
		q.RemoveLabelIDs(set.Slice()...)
	}

	if data.ParentID != uuid.Nil {
		q.SetParentID(data.ParentID)
	} else {
		q.ClearParent()
	}

	err = q.Exec(ctx)
	if err != nil {
		return ItemOut{}, err
	}

	fields, err := e.db.ItemField.Query().Where(itemfield.HasItemWith(item.ID(data.ID))).All(ctx)
	if err != nil {
		return ItemOut{}, err
	}

	fieldIds := newIDSet(fields)

	// Update Existing Fields
	for _, f := range data.Fields {
		if f.ID == uuid.Nil {
			// Create New Field
			_, err = e.db.ItemField.Create().
				SetItemID(data.ID).
				SetType(itemfield.Type(f.Type)).
				SetName(f.Name).
				SetTextValue(f.TextValue).
				SetNumberValue(f.NumberValue).
				SetBooleanValue(f.BooleanValue).
				// SetTimeValue(f.TimeValue).
				Save(ctx)
			if err != nil {
				return ItemOut{}, err
			}
		}

		opt := e.db.ItemField.Update().
			Where(
				itemfield.ID(f.ID),
				itemfield.HasItemWith(item.ID(data.ID)),
			).
			SetType(itemfield.Type(f.Type)).
			SetName(f.Name).
			SetTextValue(f.TextValue).
			SetNumberValue(f.NumberValue).
			SetBooleanValue(f.BooleanValue)
			// SetTimeValue(f.TimeValue)

		_, err = opt.Save(ctx)
		if err != nil {
			return ItemOut{}, err
		}

		fieldIds.Remove(f.ID)
		continue
	}

	// Delete Fields that are no longer present
	if fieldIds.Len() > 0 {
		_, err = e.db.ItemField.Delete().
			Where(
				itemfield.IDIn(fieldIds.Slice()...),
				itemfield.HasItemWith(item.ID(data.ID)),
			).Exec(ctx)
		if err != nil {
			return ItemOut{}, err
		}
	}

	e.publishMutationEvent(GID)
	return e.GetOne(ctx, data.ID)
}

func (e *ItemsRepository) GetAllZeroImportRef(ctx context.Context, GID uuid.UUID) ([]uuid.UUID, error) {
	var ids []uuid.UUID

	err := e.db.Item.Query().
		Where(
			item.HasGroupWith(group.ID(GID)),
			item.Or(
				item.ImportRefEQ(""),
				item.ImportRefIsNil(),
			),
		).
		Select(item.FieldID).
		Scan(ctx, &ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (e *ItemsRepository) Patch(ctx context.Context, GID, ID uuid.UUID, data ItemPatch) error {
	q := e.db.Item.Update().
		Where(
			item.ID(ID),
			item.HasGroupWith(group.ID(GID)),
		)

	if data.ImportRef != nil {
		q.SetImportRef(*data.ImportRef)
	}

	if data.Quantity != nil {
		q.SetQuantity(*data.Quantity)
	}

	e.publishMutationEvent(GID)
	return q.Exec(ctx)
}

func (e *ItemsRepository) GetAllCustomFieldValues(ctx context.Context, GID uuid.UUID, name string) ([]string, error) {
	type st struct {
		Value string `json:"text_value"`
	}

	var values []st

	err := e.db.Item.Query().
		Where(
			item.HasGroupWith(group.ID(GID)),
		).
		QueryFields().
		Where(
			itemfield.Name(name),
		).
		Unique(true).
		Select(itemfield.FieldTextValue).
		Scan(ctx, &values)
	if err != nil {
		return nil, fmt.Errorf("failed to get field values: %w", err)
	}

	valueStrings := make([]string, len(values))
	for i, f := range values {
		valueStrings[i] = f.Value
	}

	return valueStrings, nil
}

func (e *ItemsRepository) GetAllCustomFieldNames(ctx context.Context, GID uuid.UUID) ([]string, error) {
	type st struct {
		Name string `json:"name"`
	}

	var fields []st

	err := e.db.Item.Query().
		Where(
			item.HasGroupWith(group.ID(GID)),
		).
		QueryFields().
		Unique(true).
		Select(itemfield.FieldName).
		Scan(ctx, &fields)
	if err != nil {
		return nil, fmt.Errorf("failed to get custom fields: %w", err)
	}

	fieldNames := make([]string, len(fields))
	for i, f := range fields {
		fieldNames[i] = f.Name
	}

	return fieldNames, nil
}

// ZeroOutTimeFields is a helper function that can be invoked via the UI by a group member which will
// set all date fields to the beginning of the day.
//
// This is designed to resolve a long-time bug that has since been fixed with the time selector on the
// frontend. This function is intended to be used as a one-time fix for existing databases and may be
// removed in the future.
func (e *ItemsRepository) ZeroOutTimeFields(ctx context.Context, GID uuid.UUID) (int, error) {
	q := e.db.Item.Query().Where(
		item.HasGroupWith(group.ID(GID)),
		item.Or(
			item.PurchaseTimeNotNil(),
			item.PurchaseFromLT("0002-01-01"),
			item.SoldTimeNotNil(),
			item.SoldToLT("0002-01-01"),
			item.WarrantyExpiresNotNil(),
			item.WarrantyDetailsLT("0002-01-01"),
		),
	)

	items, err := q.All(ctx)
	if err != nil {
		return -1, fmt.Errorf("ZeroOutTimeFields() -> failed to get items: %w", err)
	}

	toDateOnly := func(t time.Time) time.Time {
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	}

	updated := 0

	for _, i := range items {
		updateQ := e.db.Item.Update().Where(item.ID(i.ID))

		if !i.PurchaseTime.IsZero() {
			switch {
			case i.PurchaseTime.Year() < 100:
				updateQ.ClearPurchaseTime()
			default:
				updateQ.SetPurchaseTime(toDateOnly(i.PurchaseTime))
			}
		} else {
			updateQ.ClearPurchaseTime()
		}

		if !i.SoldTime.IsZero() {
			switch {
			case i.SoldTime.Year() < 100:
				updateQ.ClearSoldTime()
			default:
				updateQ.SetSoldTime(toDateOnly(i.SoldTime))
			}
		} else {
			updateQ.ClearSoldTime()
		}

		if !i.WarrantyExpires.IsZero() {
			switch {
			case i.WarrantyExpires.Year() < 100:
				updateQ.ClearWarrantyExpires()
			default:
				updateQ.SetWarrantyExpires(toDateOnly(i.WarrantyExpires))
			}
		} else {
			updateQ.ClearWarrantyExpires()
		}

		_, err = updateQ.Save(ctx)
		if err != nil {
			return updated, fmt.Errorf("ZeroOutTimeFields() -> failed to update item: %w", err)
		}

		updated++
	}

	return updated, nil
}

func (e *ItemsRepository) SetPrimaryPhotos(ctx context.Context, GID uuid.UUID) (int, error) {
	// All items where there is no primary photo
	itemIDs, err := e.db.Item.Query().
		Where(
			item.HasGroupWith(group.ID(GID)),
			item.HasAttachmentsWith(
				attachment.TypeEQ(attachment.TypePhoto),
				attachment.Not(
					attachment.And(
						attachment.Primary(true),
						attachment.TypeEQ(attachment.TypePhoto),
					),
				),
			),
		).
		IDs(ctx)
	if err != nil {
		return -1, err
	}

	updated := 0
	for _, id := range itemIDs {
		// Find the first photo attachment
		a, err := e.db.Attachment.Query().
			Where(
				attachment.HasItemWith(item.ID(id)),
				attachment.TypeEQ(attachment.TypePhoto),
				attachment.Primary(false),
			).
			First(ctx)
		if err != nil {
			return updated, err
		}

		// Set it as primary
		_, err = e.db.Attachment.UpdateOne(a).
			SetPrimary(true).
			Save(ctx)
		if err != nil {
			return updated, err
		}

		updated++
	}

	return updated, nil
}
