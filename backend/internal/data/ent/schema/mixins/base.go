package mixins

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

type DetailsMixin struct {
	mixin.Schema
}

func (DetailsMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(255).
			NotEmpty(),
		field.String("description").
			MaxLen(1000).
			Optional(),
	}
}
