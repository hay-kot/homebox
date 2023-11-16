// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/group"
	"github.com/hay-kot/homebox/backend/internal/data/ent/groupinvitationtoken"
)

// GroupInvitationTokenCreate is the builder for creating a GroupInvitationToken entity.
type GroupInvitationTokenCreate struct {
	config
	mutation *GroupInvitationTokenMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (gitc *GroupInvitationTokenCreate) SetCreatedAt(t time.Time) *GroupInvitationTokenCreate {
	gitc.mutation.SetCreatedAt(t)
	return gitc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gitc *GroupInvitationTokenCreate) SetNillableCreatedAt(t *time.Time) *GroupInvitationTokenCreate {
	if t != nil {
		gitc.SetCreatedAt(*t)
	}
	return gitc
}

// SetUpdatedAt sets the "updated_at" field.
func (gitc *GroupInvitationTokenCreate) SetUpdatedAt(t time.Time) *GroupInvitationTokenCreate {
	gitc.mutation.SetUpdatedAt(t)
	return gitc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gitc *GroupInvitationTokenCreate) SetNillableUpdatedAt(t *time.Time) *GroupInvitationTokenCreate {
	if t != nil {
		gitc.SetUpdatedAt(*t)
	}
	return gitc
}

// SetToken sets the "token" field.
func (gitc *GroupInvitationTokenCreate) SetToken(b []byte) *GroupInvitationTokenCreate {
	gitc.mutation.SetToken(b)
	return gitc
}

// SetExpiresAt sets the "expires_at" field.
func (gitc *GroupInvitationTokenCreate) SetExpiresAt(t time.Time) *GroupInvitationTokenCreate {
	gitc.mutation.SetExpiresAt(t)
	return gitc
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (gitc *GroupInvitationTokenCreate) SetNillableExpiresAt(t *time.Time) *GroupInvitationTokenCreate {
	if t != nil {
		gitc.SetExpiresAt(*t)
	}
	return gitc
}

// SetUses sets the "uses" field.
func (gitc *GroupInvitationTokenCreate) SetUses(i int) *GroupInvitationTokenCreate {
	gitc.mutation.SetUses(i)
	return gitc
}

// SetNillableUses sets the "uses" field if the given value is not nil.
func (gitc *GroupInvitationTokenCreate) SetNillableUses(i *int) *GroupInvitationTokenCreate {
	if i != nil {
		gitc.SetUses(*i)
	}
	return gitc
}

// SetID sets the "id" field.
func (gitc *GroupInvitationTokenCreate) SetID(u uuid.UUID) *GroupInvitationTokenCreate {
	gitc.mutation.SetID(u)
	return gitc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gitc *GroupInvitationTokenCreate) SetNillableID(u *uuid.UUID) *GroupInvitationTokenCreate {
	if u != nil {
		gitc.SetID(*u)
	}
	return gitc
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gitc *GroupInvitationTokenCreate) SetGroupID(id uuid.UUID) *GroupInvitationTokenCreate {
	gitc.mutation.SetGroupID(id)
	return gitc
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (gitc *GroupInvitationTokenCreate) SetNillableGroupID(id *uuid.UUID) *GroupInvitationTokenCreate {
	if id != nil {
		gitc = gitc.SetGroupID(*id)
	}
	return gitc
}

// SetGroup sets the "group" edge to the Group entity.
func (gitc *GroupInvitationTokenCreate) SetGroup(g *Group) *GroupInvitationTokenCreate {
	return gitc.SetGroupID(g.ID)
}

// Mutation returns the GroupInvitationTokenMutation object of the builder.
func (gitc *GroupInvitationTokenCreate) Mutation() *GroupInvitationTokenMutation {
	return gitc.mutation
}

// Save creates the GroupInvitationToken in the database.
func (gitc *GroupInvitationTokenCreate) Save(ctx context.Context) (*GroupInvitationToken, error) {
	gitc.defaults()
	return withHooks(ctx, gitc.sqlSave, gitc.mutation, gitc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gitc *GroupInvitationTokenCreate) SaveX(ctx context.Context) *GroupInvitationToken {
	v, err := gitc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gitc *GroupInvitationTokenCreate) Exec(ctx context.Context) error {
	_, err := gitc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gitc *GroupInvitationTokenCreate) ExecX(ctx context.Context) {
	if err := gitc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gitc *GroupInvitationTokenCreate) defaults() {
	if _, ok := gitc.mutation.CreatedAt(); !ok {
		v := groupinvitationtoken.DefaultCreatedAt()
		gitc.mutation.SetCreatedAt(v)
	}
	if _, ok := gitc.mutation.UpdatedAt(); !ok {
		v := groupinvitationtoken.DefaultUpdatedAt()
		gitc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gitc.mutation.ExpiresAt(); !ok {
		v := groupinvitationtoken.DefaultExpiresAt()
		gitc.mutation.SetExpiresAt(v)
	}
	if _, ok := gitc.mutation.Uses(); !ok {
		v := groupinvitationtoken.DefaultUses
		gitc.mutation.SetUses(v)
	}
	if _, ok := gitc.mutation.ID(); !ok {
		v := groupinvitationtoken.DefaultID()
		gitc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gitc *GroupInvitationTokenCreate) check() error {
	if _, ok := gitc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "GroupInvitationToken.created_at"`)}
	}
	if _, ok := gitc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "GroupInvitationToken.updated_at"`)}
	}
	if _, ok := gitc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "GroupInvitationToken.token"`)}
	}
	if _, ok := gitc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`ent: missing required field "GroupInvitationToken.expires_at"`)}
	}
	if _, ok := gitc.mutation.Uses(); !ok {
		return &ValidationError{Name: "uses", err: errors.New(`ent: missing required field "GroupInvitationToken.uses"`)}
	}
	return nil
}

func (gitc *GroupInvitationTokenCreate) sqlSave(ctx context.Context) (*GroupInvitationToken, error) {
	if err := gitc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gitc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gitc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	gitc.mutation.id = &_node.ID
	gitc.mutation.done = true
	return _node, nil
}

func (gitc *GroupInvitationTokenCreate) createSpec() (*GroupInvitationToken, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupInvitationToken{config: gitc.config}
		_spec = sqlgraph.NewCreateSpec(groupinvitationtoken.Table, sqlgraph.NewFieldSpec(groupinvitationtoken.FieldID, field.TypeUUID))
	)
	if id, ok := gitc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := gitc.mutation.CreatedAt(); ok {
		_spec.SetField(groupinvitationtoken.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gitc.mutation.UpdatedAt(); ok {
		_spec.SetField(groupinvitationtoken.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := gitc.mutation.Token(); ok {
		_spec.SetField(groupinvitationtoken.FieldToken, field.TypeBytes, value)
		_node.Token = value
	}
	if value, ok := gitc.mutation.ExpiresAt(); ok {
		_spec.SetField(groupinvitationtoken.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if value, ok := gitc.mutation.Uses(); ok {
		_spec.SetField(groupinvitationtoken.FieldUses, field.TypeInt, value)
		_node.Uses = value
	}
	if nodes := gitc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupinvitationtoken.GroupTable,
			Columns: []string{groupinvitationtoken.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_invitation_tokens = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GroupInvitationTokenCreateBulk is the builder for creating many GroupInvitationToken entities in bulk.
type GroupInvitationTokenCreateBulk struct {
	config
	err      error
	builders []*GroupInvitationTokenCreate
}

// Save creates the GroupInvitationToken entities in the database.
func (gitcb *GroupInvitationTokenCreateBulk) Save(ctx context.Context) ([]*GroupInvitationToken, error) {
	if gitcb.err != nil {
		return nil, gitcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gitcb.builders))
	nodes := make([]*GroupInvitationToken, len(gitcb.builders))
	mutators := make([]Mutator, len(gitcb.builders))
	for i := range gitcb.builders {
		func(i int, root context.Context) {
			builder := gitcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupInvitationTokenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gitcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gitcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gitcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gitcb *GroupInvitationTokenCreateBulk) SaveX(ctx context.Context) []*GroupInvitationToken {
	v, err := gitcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gitcb *GroupInvitationTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := gitcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gitcb *GroupInvitationTokenCreateBulk) ExecX(ctx context.Context) {
	if err := gitcb.Exec(ctx); err != nil {
		panic(err)
	}
}
