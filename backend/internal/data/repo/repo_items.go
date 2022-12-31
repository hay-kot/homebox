package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent"
	"github.com/hay-kot/homebox/backend/internal/data/ent/group"
	"github.com/hay-kot/homebox/backend/internal/data/ent/item"
	"github.com/hay-kot/homebox/backend/internal/data/ent/itemfield"
	"github.com/hay-kot/homebox/backend/internal/data/ent/label"
	"github.com/hay-kot/homebox/backend/internal/data/ent/location"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
)

type ItemsRepository struct {
	db *ent.Client
}

type (
	ItemQuery struct {
		Page            int
		PageSize        int
		Search          string      `json:"search"`
		LocationIDs     []uuid.UUID `json:"locationIds"`
		LabelIDs        []uuid.UUID `json:"labelIds"`
		SortBy          string      `json:"sortBy"`
		IncludeArchived bool        `json:"includeArchived"`
	}

	ItemField struct {
		ID           uuid.UUID `json:"id,omitempty"`
		Type         string    `json:"type"`
		Name         string    `json:"name"`
		TextValue    string    `json:"textValue"`
		NumberValue  int       `json:"numberValue"`
		BooleanValue bool      `json:"booleanValue"`
		TimeValue    time.Time `json:"timeValue,omitempty"`
	}

	ItemCreate struct {
		ImportRef   string    `json:"-"`
		ParentID    uuid.UUID `json:"parentId" extensions:"x-nullable"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		AssetID     AssetID   `json:"-"`

		// Edges
		LocationID uuid.UUID   `json:"locationId"`
		LabelIDs   []uuid.UUID `json:"labelIds"`
	}
	ItemUpdate struct {
		ParentID    uuid.UUID `json:"parentId" extensions:"x-nullable,x-omitempty"`
		ID          uuid.UUID `json:"id"`
		AssetID     AssetID   `json:"assetId"`
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
		LifetimeWarranty bool      `json:"lifetimeWarranty"`
		WarrantyExpires  time.Time `json:"warrantyExpires"`
		WarrantyDetails  string    `json:"warrantyDetails"`

		// Purchase
		PurchaseTime  time.Time `json:"purchaseTime"`
		PurchaseFrom  string    `json:"purchaseFrom"`
		PurchasePrice float64   `json:"purchasePrice,string"`

		// Sold
		SoldTime  time.Time `json:"soldTime"`
		SoldTo    string    `json:"soldTo"`
		SoldPrice float64   `json:"soldPrice,string"`
		SoldNotes string    `json:"soldNotes"`

		// Extras
		Notes  string      `json:"notes"`
		Fields []ItemField `json:"fields"`
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
	}

	ItemOut struct {
		Parent *ItemSummary `json:"parent,omitempty" extensions:"x-nullable,x-omitempty"`
		ItemSummary
		AssetID AssetID `json:"assetId,string"`

		SerialNumber string `json:"serialNumber"`
		ModelNumber  string `json:"modelNumber"`
		Manufacturer string `json:"manufacturer"`

		// Warranty
		LifetimeWarranty bool      `json:"lifetimeWarranty"`
		WarrantyExpires  time.Time `json:"warrantyExpires"`
		WarrantyDetails  string    `json:"warrantyDetails"`

		// Purchase
		PurchaseTime time.Time `json:"purchaseTime"`
		PurchaseFrom string    `json:"purchaseFrom"`

		// Sold
		SoldTime  time.Time `json:"soldTime"`
		SoldTo    string    `json:"soldTo"`
		SoldPrice float64   `json:"soldPrice,string"`
		SoldNotes string    `json:"soldNotes"`

		// Extras
		Notes string `json:"notes"`

		Attachments []ItemAttachment `json:"attachments"`
		Fields      []ItemField      `json:"fields"`
		Children    []ItemSummary    `json:"children"`
	}
)

var (
	mapItemsSummaryErr = mapTEachErrFunc(mapItemSummary)
)

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

	return ItemSummary{
		ID:            item.ID,
		Name:          item.Name,
		Description:   item.Description,
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
	}
}

var (
	mapItemOutErr = mapTErrFunc(mapItemOut)
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
			TimeValue:    f.TimeValue,
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

	var children []ItemSummary
	if item.Edges.Children != nil {
		children = mapEach(item.Edges.Children, mapItemSummary)
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
		WarrantyExpires:  item.WarrantyExpires,
		WarrantyDetails:  item.WarrantyDetails,

		// Identification
		SerialNumber: item.SerialNumber,
		ModelNumber:  item.ModelNumber,
		Manufacturer: item.Manufacturer,

		// Purchase
		PurchaseTime: item.PurchaseTime,
		PurchaseFrom: item.PurchaseFrom,

		// Sold
		SoldTime:  item.SoldTime,
		SoldTo:    item.SoldTo,
		SoldPrice: item.SoldPrice,
		SoldNotes: item.SoldNotes,

		// Extras
		Notes:       item.Notes,
		Attachments: attachments,
		Fields:      fields,
		Children:    children,
	}
}

func (e *ItemsRepository) getOne(ctx context.Context, where ...predicate.Item) (ItemOut, error) {
	q := e.db.Item.Query().Where(where...)

	return mapItemOutErr(q.
		WithFields().
		WithLabel().
		WithLocation().
		WithGroup().
		WithChildren().
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

// GetOneByGroup returns a single item by ID. If the item does not exist, an error is returned.
// GetOneByGroup ensures that the item belongs to a specific group.
func (e *ItemsRepository) GetOneByGroup(ctx context.Context, gid, id uuid.UUID) (ItemOut, error) {
	return e.getOne(ctx, item.ID(id), item.HasGroupWith(group.ID(gid)))
}

func (e *ItemsRepository) GetIDsByAssetID(ctx context.Context, assetID AssetID) ([]uuid.UUID, error) {
	return e.db.Item.Query().Where(item.AssetID(int(assetID))).Order(ent.Desc(item.FieldCreatedAt)).IDs(ctx)
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

	if len(q.LabelIDs) > 0 {
		labels := make([]predicate.Item, 0, len(q.LabelIDs))
		for _, l := range q.LabelIDs {
			labels = append(labels, item.HasLabelWith(label.ID(l)))
		}
		qb = qb.Where(item.Or(labels...))
	}

	if len(q.LocationIDs) > 0 {
		locations := make([]predicate.Item, 0, len(q.LocationIDs))
		for _, l := range q.LocationIDs {
			locations = append(locations, item.HasLocationWith(location.ID(l)))
		}
		qb = qb.Where(item.Or(locations...))
	}

	if q.Search != "" {
		qb.Where(
			item.Or(
				item.NameContainsFold(q.Search),
				item.DescriptionContainsFold(q.Search),
				item.NotesContainsFold(q.Search),
			),
		)
	}

	if q.Page != -1 || q.PageSize != -1 {
		qb = qb.
			Offset(calculateOffset(q.Page, q.PageSize)).
			Limit(q.PageSize)
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

	count, err := qb.Count(ctx)
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

// GetAll returns all the items in the database with the Labels and Locations eager loaded.
func (e *ItemsRepository) GetAll(ctx context.Context, gid uuid.UUID) ([]ItemSummary, error) {
	return mapItemsSummaryErr(e.db.Item.Query().
		Where(item.HasGroupWith(group.ID(gid))).
		WithLabel().
		WithLocation().
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

	return e.GetOne(ctx, result.ID)
}

func (e *ItemsRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return e.db.Item.DeleteOneID(id).Exec(ctx)
}

func (e *ItemsRepository) DeleteByGroup(ctx context.Context, gid, id uuid.UUID) error {
	_, err := e.db.Item.
		Delete().
		Where(
			item.ID(id),
			item.HasGroupWith(group.ID(gid)),
		).Exec(ctx)
	return err
}

func (e *ItemsRepository) UpdateByGroup(ctx context.Context, gid uuid.UUID, data ItemUpdate) (ItemOut, error) {
	q := e.db.Item.Update().Where(item.ID(data.ID), item.HasGroupWith(group.ID(gid))).
		SetName(data.Name).
		SetDescription(data.Description).
		SetLocationID(data.LocationID).
		SetSerialNumber(data.SerialNumber).
		SetModelNumber(data.ModelNumber).
		SetManufacturer(data.Manufacturer).
		SetArchived(data.Archived).
		SetPurchaseTime(data.PurchaseTime).
		SetPurchaseFrom(data.PurchaseFrom).
		SetPurchasePrice(data.PurchasePrice).
		SetSoldTime(data.SoldTime).
		SetSoldTo(data.SoldTo).
		SetSoldPrice(data.SoldPrice).
		SetSoldNotes(data.SoldNotes).
		SetNotes(data.Notes).
		SetLifetimeWarranty(data.LifetimeWarranty).
		SetInsured(data.Insured).
		SetWarrantyExpires(data.WarrantyExpires).
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
				SetTimeValue(f.TimeValue).
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
			SetBooleanValue(f.BooleanValue).
			SetTimeValue(f.TimeValue)

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

	return e.GetOne(ctx, data.ID)
}
