package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AuthRoles holds the schema definition for the AuthRoles entity.
type AuthRoles struct {
	ent.Schema
}

// Fields of the AuthRoles.
func (AuthRoles) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("role").
			Default("user").
			Values(
				"admin",       // can do everything - currently unused
				"user",        // default login role
				"attachments", // Read Attachments
			),
	}
}

// Edges of the AuthRoles.
func (AuthRoles) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("token", AuthTokens.Type).
			Ref("roles").
			Unique(),
	}
}
