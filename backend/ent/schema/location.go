package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/hay-kot/content/backend/ent/schema/mixins"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

func (Location) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.DetailsMixin{},
	}
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return nil
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).
			Ref("locations").
			Unique().
			Required(),
		edge.To("items", Item.Type),
	}
}
