package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/hay-kot/homebox/backend/internal/data/ent/schema/mixins"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

func (Item) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.DetailsMixin{},
	}
}

func (Item) Indexes() []ent.Index {
	return []ent.Index{
		// Unique index on the "title" field.
		index.Fields("name"),
		index.Fields("manufacturer"),
		index.Fields("model_number"),
		index.Fields("serial_number"),
		index.Fields("archived"),
	}
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("import_ref").
			Optional().
			MaxLen(100).
			Immutable(),
		field.String("notes").
			MaxLen(1000).
			Optional(),
		field.Int("quantity").
			Default(1),
		field.Bool("insured").
			Default(false),
		field.Bool("archived").
			Default(false),

		// ------------------------------------
		// item identification
		field.String("serial_number").
			MaxLen(255).
			Optional(),
		field.String("model_number").
			MaxLen(255).
			Optional(),
		field.String("manufacturer").
			MaxLen(255).
			Optional(),

		// ------------------------------------
		// Item Warranty
		field.Bool("lifetime_warranty").
			Default(false),
		field.Time("warranty_expires").
			Optional(),
		field.Text("warranty_details").
			MaxLen(1000).
			Optional(),

		// ------------------------------------
		// item purchase
		field.Time("purchase_time").
			Optional(),
		field.String("purchase_from").
			Optional(),
		field.Float("purchase_price").
			Default(0),

		// ------------------------------------
		// Sold Details
		field.Time("sold_time").
			Optional(),
		field.String("sold_to").
			Optional(),
		field.Float("sold_price").
			Default(0),
		field.String("sold_notes").
			MaxLen(1000).
			Optional(),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Item.Type).
			From("parent").
			Unique(),
		edge.From("group", Group.Type).
			Ref("items").
			Required().
			Unique(),
		edge.From("label", Label.Type).
			Ref("items"),
		edge.From("location", Location.Type).
			Ref("items").
			Unique(),
		edge.To("fields", ItemField.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("attachments", Attachment.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
