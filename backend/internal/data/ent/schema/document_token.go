package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/hay-kot/homebox/backend/internal/data/ent/schema/mixins"
)

// DocumentToken holds the schema definition for the DocumentToken entity.
type DocumentToken struct {
	ent.Schema
}

func (DocumentToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the DocumentToken.
func (DocumentToken) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("token").
			NotEmpty().
			Unique(),
		field.Int("uses").
			Default(1),
		field.Time("expires_at").
			Default(func() time.Time { return time.Now().Add(time.Minute * 10) }),
	}
}

// Edges of the DocumentToken.
func (DocumentToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("document", Document.Type).
			Ref("document_tokens").
			Unique(),
	}
}

func (DocumentToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("token"),
	}
}
