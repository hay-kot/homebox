// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AttachmentsColumns holds the columns for the "attachments" table.
	AttachmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"photo", "manual", "warranty", "attachment"}, Default: "attachment"},
		{Name: "document_attachments", Type: field.TypeUUID},
		{Name: "item_attachments", Type: field.TypeUUID},
	}
	// AttachmentsTable holds the schema information for the "attachments" table.
	AttachmentsTable = &schema.Table{
		Name:       "attachments",
		Columns:    AttachmentsColumns,
		PrimaryKey: []*schema.Column{AttachmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attachments_documents_attachments",
				Columns:    []*schema.Column{AttachmentsColumns[4]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "attachments_items_attachments",
				Columns:    []*schema.Column{AttachmentsColumns[5]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// AuthTokensColumns holds the columns for the "auth_tokens" table.
	AuthTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "token", Type: field.TypeBytes, Unique: true},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "user_auth_tokens", Type: field.TypeUUID, Nullable: true},
	}
	// AuthTokensTable holds the schema information for the "auth_tokens" table.
	AuthTokensTable = &schema.Table{
		Name:       "auth_tokens",
		Columns:    AuthTokensColumns,
		PrimaryKey: []*schema.Column{AuthTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "auth_tokens_users_auth_tokens",
				Columns:    []*schema.Column{AuthTokensColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "authtokens_token",
				Unique:  false,
				Columns: []*schema.Column{AuthTokensColumns[3]},
			},
		},
	}
	// DocumentsColumns holds the columns for the "documents" table.
	DocumentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "path", Type: field.TypeString, Size: 500},
		{Name: "group_documents", Type: field.TypeUUID},
	}
	// DocumentsTable holds the schema information for the "documents" table.
	DocumentsTable = &schema.Table{
		Name:       "documents",
		Columns:    DocumentsColumns,
		PrimaryKey: []*schema.Column{DocumentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "documents_groups_documents",
				Columns:    []*schema.Column{DocumentsColumns[5]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// DocumentTokensColumns holds the columns for the "document_tokens" table.
	DocumentTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "token", Type: field.TypeBytes, Unique: true},
		{Name: "uses", Type: field.TypeInt, Default: 1},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "document_document_tokens", Type: field.TypeUUID, Nullable: true},
	}
	// DocumentTokensTable holds the schema information for the "document_tokens" table.
	DocumentTokensTable = &schema.Table{
		Name:       "document_tokens",
		Columns:    DocumentTokensColumns,
		PrimaryKey: []*schema.Column{DocumentTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "document_tokens_documents_document_tokens",
				Columns:    []*schema.Column{DocumentTokensColumns[6]},
				RefColumns: []*schema.Column{DocumentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "documenttoken_token",
				Unique:  false,
				Columns: []*schema.Column{DocumentTokensColumns[3]},
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "currency", Type: field.TypeEnum, Enums: []string{"usd"}, Default: "usd"},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 1000},
		{Name: "notes", Type: field.TypeString, Nullable: true, Size: 1000},
		{Name: "serial_number", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "model_number", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "manufacturer", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "lifetime_warranty", Type: field.TypeBool, Default: false},
		{Name: "warranty_expires", Type: field.TypeTime, Nullable: true},
		{Name: "warranty_details", Type: field.TypeString, Nullable: true, Size: 1000},
		{Name: "purchase_time", Type: field.TypeTime, Nullable: true},
		{Name: "purchase_from", Type: field.TypeString, Nullable: true},
		{Name: "purchase_price", Type: field.TypeFloat64, Default: 0},
		{Name: "sold_time", Type: field.TypeTime, Nullable: true},
		{Name: "sold_to", Type: field.TypeString, Nullable: true},
		{Name: "sold_price", Type: field.TypeFloat64, Default: 0},
		{Name: "sold_notes", Type: field.TypeString, Nullable: true, Size: 1000},
		{Name: "group_items", Type: field.TypeUUID},
		{Name: "location_items", Type: field.TypeUUID, Nullable: true},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "items_groups_items",
				Columns:    []*schema.Column{ItemsColumns[19]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "items_locations_items",
				Columns:    []*schema.Column{ItemsColumns[20]},
				RefColumns: []*schema.Column{LocationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "item_name",
				Unique:  false,
				Columns: []*schema.Column{ItemsColumns[3]},
			},
			{
				Name:    "item_manufacturer",
				Unique:  false,
				Columns: []*schema.Column{ItemsColumns[8]},
			},
			{
				Name:    "item_model_number",
				Unique:  false,
				Columns: []*schema.Column{ItemsColumns[7]},
			},
			{
				Name:    "item_serial_number",
				Unique:  false,
				Columns: []*schema.Column{ItemsColumns[6]},
			},
		},
	}
	// ItemFieldsColumns holds the columns for the "item_fields" table.
	ItemFieldsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 1000},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"text", "number", "boolean", "time"}},
		{Name: "text_value", Type: field.TypeString, Nullable: true, Size: 500},
		{Name: "number_value", Type: field.TypeInt, Nullable: true},
		{Name: "boolean_value", Type: field.TypeBool, Default: false},
		{Name: "time_value", Type: field.TypeTime},
		{Name: "item_fields", Type: field.TypeUUID, Nullable: true},
	}
	// ItemFieldsTable holds the schema information for the "item_fields" table.
	ItemFieldsTable = &schema.Table{
		Name:       "item_fields",
		Columns:    ItemFieldsColumns,
		PrimaryKey: []*schema.Column{ItemFieldsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "item_fields_items_fields",
				Columns:    []*schema.Column{ItemFieldsColumns[10]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// LabelsColumns holds the columns for the "labels" table.
	LabelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 1000},
		{Name: "color", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "group_labels", Type: field.TypeUUID},
	}
	// LabelsTable holds the schema information for the "labels" table.
	LabelsTable = &schema.Table{
		Name:       "labels",
		Columns:    LabelsColumns,
		PrimaryKey: []*schema.Column{LabelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "labels_groups_labels",
				Columns:    []*schema.Column{LabelsColumns[6]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// LocationsColumns holds the columns for the "locations" table.
	LocationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 1000},
		{Name: "group_locations", Type: field.TypeUUID},
	}
	// LocationsTable holds the schema information for the "locations" table.
	LocationsTable = &schema.Table{
		Name:       "locations",
		Columns:    LocationsColumns,
		PrimaryKey: []*schema.Column{LocationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "locations_groups_locations",
				Columns:    []*schema.Column{LocationsColumns[5]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "password", Type: field.TypeString, Size: 255},
		{Name: "is_superuser", Type: field.TypeBool, Default: false},
		{Name: "group_users", Type: field.TypeUUID},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_groups_users",
				Columns:    []*schema.Column{UsersColumns[7]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// LabelItemsColumns holds the columns for the "label_items" table.
	LabelItemsColumns = []*schema.Column{
		{Name: "label_id", Type: field.TypeUUID},
		{Name: "item_id", Type: field.TypeUUID},
	}
	// LabelItemsTable holds the schema information for the "label_items" table.
	LabelItemsTable = &schema.Table{
		Name:       "label_items",
		Columns:    LabelItemsColumns,
		PrimaryKey: []*schema.Column{LabelItemsColumns[0], LabelItemsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "label_items_label_id",
				Columns:    []*schema.Column{LabelItemsColumns[0]},
				RefColumns: []*schema.Column{LabelsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "label_items_item_id",
				Columns:    []*schema.Column{LabelItemsColumns[1]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AttachmentsTable,
		AuthTokensTable,
		DocumentsTable,
		DocumentTokensTable,
		GroupsTable,
		ItemsTable,
		ItemFieldsTable,
		LabelsTable,
		LocationsTable,
		UsersTable,
		LabelItemsTable,
	}
)

func init() {
	AttachmentsTable.ForeignKeys[0].RefTable = DocumentsTable
	AttachmentsTable.ForeignKeys[1].RefTable = ItemsTable
	AuthTokensTable.ForeignKeys[0].RefTable = UsersTable
	DocumentsTable.ForeignKeys[0].RefTable = GroupsTable
	DocumentTokensTable.ForeignKeys[0].RefTable = DocumentsTable
	ItemsTable.ForeignKeys[0].RefTable = GroupsTable
	ItemsTable.ForeignKeys[1].RefTable = LocationsTable
	ItemFieldsTable.ForeignKeys[0].RefTable = ItemsTable
	LabelsTable.ForeignKeys[0].RefTable = GroupsTable
	LocationsTable.ForeignKeys[0].RefTable = GroupsTable
	UsersTable.ForeignKeys[0].RefTable = GroupsTable
	LabelItemsTable.ForeignKeys[0].RefTable = LabelsTable
	LabelItemsTable.ForeignKeys[1].RefTable = ItemsTable
}
