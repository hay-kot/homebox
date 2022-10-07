// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hay-kot/homebox/backend/ent/groupinvitationtoken"
	"github.com/hay-kot/homebox/backend/ent/predicate"
)

// GroupInvitationTokenDelete is the builder for deleting a GroupInvitationToken entity.
type GroupInvitationTokenDelete struct {
	config
	hooks    []Hook
	mutation *GroupInvitationTokenMutation
}

// Where appends a list predicates to the GroupInvitationTokenDelete builder.
func (gitd *GroupInvitationTokenDelete) Where(ps ...predicate.GroupInvitationToken) *GroupInvitationTokenDelete {
	gitd.mutation.Where(ps...)
	return gitd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gitd *GroupInvitationTokenDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gitd.hooks) == 0 {
		affected, err = gitd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupInvitationTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gitd.mutation = mutation
			affected, err = gitd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gitd.hooks) - 1; i >= 0; i-- {
			if gitd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gitd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gitd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gitd *GroupInvitationTokenDelete) ExecX(ctx context.Context) int {
	n, err := gitd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gitd *GroupInvitationTokenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: groupinvitationtoken.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: groupinvitationtoken.FieldID,
			},
		},
	}
	if ps := gitd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gitd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// GroupInvitationTokenDeleteOne is the builder for deleting a single GroupInvitationToken entity.
type GroupInvitationTokenDeleteOne struct {
	gitd *GroupInvitationTokenDelete
}

// Exec executes the deletion query.
func (gitdo *GroupInvitationTokenDeleteOne) Exec(ctx context.Context) error {
	n, err := gitdo.gitd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{groupinvitationtoken.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gitdo *GroupInvitationTokenDeleteOne) ExecX(ctx context.Context) {
	gitdo.gitd.ExecX(ctx)
}
