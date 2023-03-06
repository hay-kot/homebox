package schema

import (
	"entgo.io/ent"

	"github.com/hay-kot/homebox/backend/internal/data/ent/schema/mixins"
)

type Test struct {
	ent.Schema
}

func (Test) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		GroupMixin{ref: "tests"},
	}
}

// Fields of the Test.
func (Test) Fields() []ent.Field {
	return []ent.Field{
		// field.String("name").
	}
}

// Edges of the Test.
func (Test) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.From("group", Group.Type).
	}
}

func (Test) Indexes() []ent.Index {
	return []ent.Index{
		// index.Fields("token"),
	}
}