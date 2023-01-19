// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hay-kot/homebox/backend/internal/data/ent/location"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
)

// LocationDelete is the builder for deleting a Location entity.
type LocationDelete struct {
	config
	hooks    []Hook
	mutation *LocationMutation
}

// Where appends a list predicates to the LocationDelete builder.
func (ld *LocationDelete) Where(ps ...predicate.Location) *LocationDelete {
	ld.mutation.Where(ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LocationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, LocationMutation](ctx, ld.sqlExec, ld.mutation, ld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LocationDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LocationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: location.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: location.FieldID,
			},
		},
	}
	if ps := ld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ld.mutation.done = true
	return affected, err
}

// LocationDeleteOne is the builder for deleting a single Location entity.
type LocationDeleteOne struct {
	ld *LocationDelete
}

// Exec executes the deletion query.
func (ldo *LocationDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{location.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LocationDeleteOne) ExecX(ctx context.Context) {
	ldo.ld.ExecX(ctx)
}
