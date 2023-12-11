package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/schema/mixins"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the Home.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(255).
			NotEmpty(),
		field.Enum("currency").
			Default("usd").
			Values(
				"aed",
				"aud",
				"bgn",
				"brl",
				"cad",
				"chf",
				"czk",
				"dkk",
				"eur",
				"gbp",
				"hkd",
				"idr",
				"inr",
				"jpy",
				"krw",
				"mxn",
				"nok",
				"nzd",
				"pln",
				"rmb",
				"ron",
				"rub",
				"sar",
				"sek",
				"sgd",
				"thb",
				"try",
				"usd",
				"xag",
				"xau",
				"zar",
			),
	}
}

// Edges of the Home.
func (Group) Edges() []ent.Edge {
	owned := func(name string, t any) ent.Edge {
		return edge.To(name, t).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			})
	}

	return []ent.Edge{
		owned("users", User.Type),
		owned("locations", Location.Type),
		owned("items", Item.Type),
		owned("labels", Label.Type),
		owned("documents", Document.Type),
		owned("invitation_tokens", GroupInvitationToken.Type),
		owned("notifiers", Notifier.Type),
		// $scaffold_edge
	}
}

// GroupMixin when embedded in an ent.Schema, adds a reference to
// the Group entity.
type GroupMixin struct {
	ref   string
	field string
	mixin.Schema
}

func (g GroupMixin) Fields() []ent.Field {
	if g.field != "" {
		return []ent.Field{
			field.UUID(g.field, uuid.UUID{}),
		}
	}

	return nil
}

func (g GroupMixin) Edges() []ent.Edge {
	edge := edge.From("group", Group.Type).
		Ref(g.ref).
		Unique().
		Required()

	if g.field != "" {
		edge = edge.Field(g.field)
	}

	return []ent.Edge{edge}
}
