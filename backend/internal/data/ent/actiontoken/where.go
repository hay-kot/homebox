// Code generated by ent, DO NOT EDIT.

package actiontoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldEQ(FieldUserID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v []byte) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldEQ(FieldToken, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldNotIn(FieldUserID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldLTE(FieldUpdatedAt, v))
}

// ActionEQ applies the EQ predicate on the "action" field.
func ActionEQ(v Action) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldEQ(FieldAction, v))
}

// ActionNEQ applies the NEQ predicate on the "action" field.
func ActionNEQ(v Action) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldNEQ(FieldAction, v))
}

// ActionIn applies the In predicate on the "action" field.
func ActionIn(vs ...Action) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldIn(FieldAction, vs...))
}

// ActionNotIn applies the NotIn predicate on the "action" field.
func ActionNotIn(vs ...Action) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldNotIn(FieldAction, vs...))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v []byte) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v []byte) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...[]byte) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...[]byte) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v []byte) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v []byte) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v []byte) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v []byte) predicate.ActionToken {
	return predicate.ActionToken(sql.FieldLTE(FieldToken, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.ActionToken {
	return predicate.ActionToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.ActionToken {
	return predicate.ActionToken(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ActionToken) predicate.ActionToken {
	return predicate.ActionToken(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ActionToken) predicate.ActionToken {
	return predicate.ActionToken(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ActionToken) predicate.ActionToken {
	return predicate.ActionToken(sql.NotPredicates(p))
}
