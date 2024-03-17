package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/hay-kot/homebox/backend/internal/data/ent/schema/mixins"
)

type ActionToken struct {
	ent.Schema
}

func (ActionToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UserMixin{
			ref:   "action_tokens",
			field: "user_id",
		},
		mixins.BaseMixin{},
	}
}

// Fields of the ActionToken.
func (ActionToken) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("action").
			Values("reset_password").
			Default("reset_password"),
		field.Bytes("token").
			Unique(),
	}
}

func (ActionToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("token"),
		index.Fields("action"),
		index.Fields("user_id"),
	}
}
