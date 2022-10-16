package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/hay-kot/homebox/backend/ent/schema/mixins"
)

// Attachment holds the schema definition for the Attachment entity.
type Attachment struct {
	ent.Schema
}

func (Attachment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the Attachment.
func (Attachment) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values("photo", "manual", "warranty", "attachment", "receipt").
			Default("attachment"),
	}
}

// Edges of the Attachment.
func (Attachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("item", Item.Type).
			Ref("attachments").
			Required().
			Unique(),
		edge.From("document", Document.Type).
			Ref("attachments").
			Required().
			Unique(),
	}
}
