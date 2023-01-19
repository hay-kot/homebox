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
	"github.com/hay-kot/homebox/backend/internal/data/ent/item"
	"github.com/hay-kot/homebox/backend/internal/data/ent/label"
)

// LabelCreate is the builder for creating a Label entity.
type LabelCreate struct {
	config
	mutation *LabelMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (lc *LabelCreate) SetCreatedAt(t time.Time) *LabelCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LabelCreate) SetNillableCreatedAt(t *time.Time) *LabelCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LabelCreate) SetUpdatedAt(t time.Time) *LabelCreate {
	lc.mutation.SetUpdatedAt(t)
	return lc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lc *LabelCreate) SetNillableUpdatedAt(t *time.Time) *LabelCreate {
	if t != nil {
		lc.SetUpdatedAt(*t)
	}
	return lc
}

// SetName sets the "name" field.
func (lc *LabelCreate) SetName(s string) *LabelCreate {
	lc.mutation.SetName(s)
	return lc
}

// SetDescription sets the "description" field.
func (lc *LabelCreate) SetDescription(s string) *LabelCreate {
	lc.mutation.SetDescription(s)
	return lc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (lc *LabelCreate) SetNillableDescription(s *string) *LabelCreate {
	if s != nil {
		lc.SetDescription(*s)
	}
	return lc
}

// SetColor sets the "color" field.
func (lc *LabelCreate) SetColor(s string) *LabelCreate {
	lc.mutation.SetColor(s)
	return lc
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (lc *LabelCreate) SetNillableColor(s *string) *LabelCreate {
	if s != nil {
		lc.SetColor(*s)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LabelCreate) SetID(u uuid.UUID) *LabelCreate {
	lc.mutation.SetID(u)
	return lc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lc *LabelCreate) SetNillableID(u *uuid.UUID) *LabelCreate {
	if u != nil {
		lc.SetID(*u)
	}
	return lc
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (lc *LabelCreate) SetGroupID(id uuid.UUID) *LabelCreate {
	lc.mutation.SetGroupID(id)
	return lc
}

// SetGroup sets the "group" edge to the Group entity.
func (lc *LabelCreate) SetGroup(g *Group) *LabelCreate {
	return lc.SetGroupID(g.ID)
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (lc *LabelCreate) AddItemIDs(ids ...uuid.UUID) *LabelCreate {
	lc.mutation.AddItemIDs(ids...)
	return lc
}

// AddItems adds the "items" edges to the Item entity.
func (lc *LabelCreate) AddItems(i ...*Item) *LabelCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return lc.AddItemIDs(ids...)
}

// Mutation returns the LabelMutation object of the builder.
func (lc *LabelCreate) Mutation() *LabelMutation {
	return lc.mutation
}

// Save creates the Label in the database.
func (lc *LabelCreate) Save(ctx context.Context) (*Label, error) {
	lc.defaults()
	return withHooks[*Label, LabelMutation](ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LabelCreate) SaveX(ctx context.Context) *Label {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LabelCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LabelCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LabelCreate) defaults() {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		v := label.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		v := label.DefaultUpdatedAt()
		lc.mutation.SetUpdatedAt(v)
	}
	if _, ok := lc.mutation.ID(); !ok {
		v := label.DefaultID()
		lc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LabelCreate) check() error {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Label.created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Label.updated_at"`)}
	}
	if _, ok := lc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Label.name"`)}
	}
	if v, ok := lc.mutation.Name(); ok {
		if err := label.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Label.name": %w`, err)}
		}
	}
	if v, ok := lc.mutation.Description(); ok {
		if err := label.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Label.description": %w`, err)}
		}
	}
	if v, ok := lc.mutation.Color(); ok {
		if err := label.ColorValidator(v); err != nil {
			return &ValidationError{Name: "color", err: fmt.Errorf(`ent: validator failed for field "Label.color": %w`, err)}
		}
	}
	if _, ok := lc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "Label.group"`)}
	}
	return nil
}

func (lc *LabelCreate) sqlSave(ctx context.Context) (*Label, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
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
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LabelCreate) createSpec() (*Label, *sqlgraph.CreateSpec) {
	var (
		_node = &Label{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: label.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: label.FieldID,
			},
		}
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(label.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.SetField(label.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := lc.mutation.Name(); ok {
		_spec.SetField(label.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := lc.mutation.Description(); ok {
		_spec.SetField(label.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := lc.mutation.Color(); ok {
		_spec.SetField(label.FieldColor, field.TypeString, value)
		_node.Color = value
	}
	if nodes := lc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   label.GroupTable,
			Columns: []string{label.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_labels = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.ItemsTable,
			Columns: label.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LabelCreateBulk is the builder for creating many Label entities in bulk.
type LabelCreateBulk struct {
	config
	builders []*LabelCreate
}

// Save creates the Label entities in the database.
func (lcb *LabelCreateBulk) Save(ctx context.Context) ([]*Label, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Label, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LabelMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LabelCreateBulk) SaveX(ctx context.Context) []*Label {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LabelCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LabelCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
