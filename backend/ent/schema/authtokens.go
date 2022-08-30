package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// AuthTokens holds the schema definition for the AuthTokens entity.
type AuthTokens struct {
	ent.Schema
}

// Fields of the AuthTokens.
func (AuthTokens) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("token").
			Unique(),
		field.Time("expires_at").
			Default(func() time.Time { return time.Now().Add(time.Hour * 24 * 7) }),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the AuthTokens.
func (AuthTokens) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("auth_tokens").
			Unique(),
	}
}

func (AuthTokens) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("token"),
	}
}
