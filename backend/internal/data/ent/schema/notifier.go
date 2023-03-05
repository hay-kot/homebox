package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/schema/mixins"
)

type Notifier struct {
	ent.Schema
}

func (Notifier) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the Notifier.
func (Notifier) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.String("name").
			MaxLen(255).
			NotEmpty(),
		field.String("url").
			Sensitive().
			MaxLen(2083). // supposed max length of URL
			NotEmpty(),
		field.Bool("is_active").
			Default(true),
	}
}

// Edges of the Notifier.
func (Notifier) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Field("user_id").
			Ref("notifiers").
			Required().
			Unique(),
	}
}

func (Notifier) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("user_id", "is_active"),
	}
}
