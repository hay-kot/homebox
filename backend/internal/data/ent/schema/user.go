package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/hay-kot/homebox/backend/internal/data/ent/schema/mixins"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(255).
			NotEmpty(),
		field.String("email").
			MaxLen(255).
			NotEmpty().
			Unique(),
		field.String("password").
			MaxLen(255).
			NotEmpty().
			Sensitive(),
		field.Bool("is_superuser").
			Default(false),
		field.Bool("superuser").
			Default(false),
		field.Enum("role").
			Default("user").
			Values("user", "owner"),
		field.Time("activated_on").
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).
			Ref("users").
			Required().
			Unique(),
		edge.To("auth_tokens", AuthTokens.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("notifiers", Notifier.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
