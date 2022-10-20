package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"github.com/hay-kot/homebox/backend/ent/schema/mixins"
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
		edge.To("children", Location.Type).
			From("parent").
			Unique(),
		edge.From("group", Group.Type).
			Ref("locations").
			Unique().
			Required(),
		edge.To("items", Item.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
