// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hay-kot/homebox/backend/internal/data/ent/authroles"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
)

// AuthRolesDelete is the builder for deleting a AuthRoles entity.
type AuthRolesDelete struct {
	config
	hooks    []Hook
	mutation *AuthRolesMutation
}

// Where appends a list predicates to the AuthRolesDelete builder.
func (ard *AuthRolesDelete) Where(ps ...predicate.AuthRoles) *AuthRolesDelete {
	ard.mutation.Where(ps...)
	return ard
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ard *AuthRolesDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, AuthRolesMutation](ctx, ard.sqlExec, ard.mutation, ard.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ard *AuthRolesDelete) ExecX(ctx context.Context) int {
	n, err := ard.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ard *AuthRolesDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(authroles.Table, sqlgraph.NewFieldSpec(authroles.FieldID, field.TypeInt))
	if ps := ard.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ard.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ard.mutation.done = true
	return affected, err
}

// AuthRolesDeleteOne is the builder for deleting a single AuthRoles entity.
type AuthRolesDeleteOne struct {
	ard *AuthRolesDelete
}

// Where appends a list predicates to the AuthRolesDelete builder.
func (ardo *AuthRolesDeleteOne) Where(ps ...predicate.AuthRoles) *AuthRolesDeleteOne {
	ardo.ard.mutation.Where(ps...)
	return ardo
}

// Exec executes the deletion query.
func (ardo *AuthRolesDeleteOne) Exec(ctx context.Context) error {
	n, err := ardo.ard.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{authroles.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ardo *AuthRolesDeleteOne) ExecX(ctx context.Context) {
	if err := ardo.Exec(ctx); err != nil {
		panic(err)
	}
}
