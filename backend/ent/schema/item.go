package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/ent/schema/mixins"
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

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("notes").
			MaxLen(1000).
			Optional(),
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
		// item purchase
		field.Time("purchase_time").
			Optional(),
		field.String("purchase_from").
			Optional(),
		field.Float("purchase_price").
			Default(0),
		field.UUID("purchase_receipt_id", uuid.UUID{}).
			Optional(),
		// ------------------------------------
		// Sold Details
		field.Time("sold_time").
			Optional(),
		field.String("sold_to").
			Optional(),
		field.Float("sold_price").
			Default(0),
		field.UUID("sold_receipt_id", uuid.UUID{}).
			Optional(),
		field.String("sold_notes").
			MaxLen(1000).
			Optional(),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).
			Ref("items").
			Required().
			Unique(),
		edge.From("location", Location.Type).
			Ref("items").
			Unique(),
		edge.To("fields", ItemField.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
		edge.From("label", Label.Type).
			Ref("items"),
	}
}
