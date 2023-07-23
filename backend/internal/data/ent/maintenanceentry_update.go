// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/item"
	"github.com/hay-kot/homebox/backend/internal/data/ent/maintenanceentry"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
)

// MaintenanceEntryUpdate is the builder for updating MaintenanceEntry entities.
type MaintenanceEntryUpdate struct {
	config
	hooks    []Hook
	mutation *MaintenanceEntryMutation
}

// Where appends a list predicates to the MaintenanceEntryUpdate builder.
func (meu *MaintenanceEntryUpdate) Where(ps ...predicate.MaintenanceEntry) *MaintenanceEntryUpdate {
	meu.mutation.Where(ps...)
	return meu
}

// SetUpdatedAt sets the "updated_at" field.
func (meu *MaintenanceEntryUpdate) SetUpdatedAt(t time.Time) *MaintenanceEntryUpdate {
	meu.mutation.SetUpdatedAt(t)
	return meu
}

// SetItemID sets the "item_id" field.
func (meu *MaintenanceEntryUpdate) SetItemID(u uuid.UUID) *MaintenanceEntryUpdate {
	meu.mutation.SetItemID(u)
	return meu
}

// SetDate sets the "date" field.
func (meu *MaintenanceEntryUpdate) SetDate(t time.Time) *MaintenanceEntryUpdate {
	meu.mutation.SetDate(t)
	return meu
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (meu *MaintenanceEntryUpdate) SetNillableDate(t *time.Time) *MaintenanceEntryUpdate {
	if t != nil {
		meu.SetDate(*t)
	}
	return meu
}

// ClearDate clears the value of the "date" field.
func (meu *MaintenanceEntryUpdate) ClearDate() *MaintenanceEntryUpdate {
	meu.mutation.ClearDate()
	return meu
}

// SetScheduledDate sets the "scheduled_date" field.
func (meu *MaintenanceEntryUpdate) SetScheduledDate(t time.Time) *MaintenanceEntryUpdate {
	meu.mutation.SetScheduledDate(t)
	return meu
}

// SetNillableScheduledDate sets the "scheduled_date" field if the given value is not nil.
func (meu *MaintenanceEntryUpdate) SetNillableScheduledDate(t *time.Time) *MaintenanceEntryUpdate {
	if t != nil {
		meu.SetScheduledDate(*t)
	}
	return meu
}

// ClearScheduledDate clears the value of the "scheduled_date" field.
func (meu *MaintenanceEntryUpdate) ClearScheduledDate() *MaintenanceEntryUpdate {
	meu.mutation.ClearScheduledDate()
	return meu
}

// SetName sets the "name" field.
func (meu *MaintenanceEntryUpdate) SetName(s string) *MaintenanceEntryUpdate {
	meu.mutation.SetName(s)
	return meu
}

// SetDescription sets the "description" field.
func (meu *MaintenanceEntryUpdate) SetDescription(s string) *MaintenanceEntryUpdate {
	meu.mutation.SetDescription(s)
	return meu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (meu *MaintenanceEntryUpdate) SetNillableDescription(s *string) *MaintenanceEntryUpdate {
	if s != nil {
		meu.SetDescription(*s)
	}
	return meu
}

// ClearDescription clears the value of the "description" field.
func (meu *MaintenanceEntryUpdate) ClearDescription() *MaintenanceEntryUpdate {
	meu.mutation.ClearDescription()
	return meu
}

// SetCost sets the "cost" field.
func (meu *MaintenanceEntryUpdate) SetCost(f float64) *MaintenanceEntryUpdate {
	meu.mutation.ResetCost()
	meu.mutation.SetCost(f)
	return meu
}

// SetNillableCost sets the "cost" field if the given value is not nil.
func (meu *MaintenanceEntryUpdate) SetNillableCost(f *float64) *MaintenanceEntryUpdate {
	if f != nil {
		meu.SetCost(*f)
	}
	return meu
}

// AddCost adds f to the "cost" field.
func (meu *MaintenanceEntryUpdate) AddCost(f float64) *MaintenanceEntryUpdate {
	meu.mutation.AddCost(f)
	return meu
}

// SetItem sets the "item" edge to the Item entity.
func (meu *MaintenanceEntryUpdate) SetItem(i *Item) *MaintenanceEntryUpdate {
	return meu.SetItemID(i.ID)
}

// Mutation returns the MaintenanceEntryMutation object of the builder.
func (meu *MaintenanceEntryUpdate) Mutation() *MaintenanceEntryMutation {
	return meu.mutation
}

// ClearItem clears the "item" edge to the Item entity.
func (meu *MaintenanceEntryUpdate) ClearItem() *MaintenanceEntryUpdate {
	meu.mutation.ClearItem()
	return meu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (meu *MaintenanceEntryUpdate) Save(ctx context.Context) (int, error) {
	meu.defaults()
	return withHooks(ctx, meu.sqlSave, meu.mutation, meu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (meu *MaintenanceEntryUpdate) SaveX(ctx context.Context) int {
	affected, err := meu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (meu *MaintenanceEntryUpdate) Exec(ctx context.Context) error {
	_, err := meu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (meu *MaintenanceEntryUpdate) ExecX(ctx context.Context) {
	if err := meu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (meu *MaintenanceEntryUpdate) defaults() {
	if _, ok := meu.mutation.UpdatedAt(); !ok {
		v := maintenanceentry.UpdateDefaultUpdatedAt()
		meu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (meu *MaintenanceEntryUpdate) check() error {
	if v, ok := meu.mutation.Name(); ok {
		if err := maintenanceentry.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MaintenanceEntry.name": %w`, err)}
		}
	}
	if v, ok := meu.mutation.Description(); ok {
		if err := maintenanceentry.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "MaintenanceEntry.description": %w`, err)}
		}
	}
	if _, ok := meu.mutation.ItemID(); meu.mutation.ItemCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MaintenanceEntry.item"`)
	}
	return nil
}

func (meu *MaintenanceEntryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := meu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(maintenanceentry.Table, maintenanceentry.Columns, sqlgraph.NewFieldSpec(maintenanceentry.FieldID, field.TypeUUID))
	if ps := meu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := meu.mutation.UpdatedAt(); ok {
		_spec.SetField(maintenanceentry.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := meu.mutation.Date(); ok {
		_spec.SetField(maintenanceentry.FieldDate, field.TypeTime, value)
	}
	if meu.mutation.DateCleared() {
		_spec.ClearField(maintenanceentry.FieldDate, field.TypeTime)
	}
	if value, ok := meu.mutation.ScheduledDate(); ok {
		_spec.SetField(maintenanceentry.FieldScheduledDate, field.TypeTime, value)
	}
	if meu.mutation.ScheduledDateCleared() {
		_spec.ClearField(maintenanceentry.FieldScheduledDate, field.TypeTime)
	}
	if value, ok := meu.mutation.Name(); ok {
		_spec.SetField(maintenanceentry.FieldName, field.TypeString, value)
	}
	if value, ok := meu.mutation.Description(); ok {
		_spec.SetField(maintenanceentry.FieldDescription, field.TypeString, value)
	}
	if meu.mutation.DescriptionCleared() {
		_spec.ClearField(maintenanceentry.FieldDescription, field.TypeString)
	}
	if value, ok := meu.mutation.Cost(); ok {
		_spec.SetField(maintenanceentry.FieldCost, field.TypeFloat64, value)
	}
	if value, ok := meu.mutation.AddedCost(); ok {
		_spec.AddField(maintenanceentry.FieldCost, field.TypeFloat64, value)
	}
	if meu.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   maintenanceentry.ItemTable,
			Columns: []string{maintenanceentry.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := meu.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   maintenanceentry.ItemTable,
			Columns: []string{maintenanceentry.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, meu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{maintenanceentry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	meu.mutation.done = true
	return n, nil
}

// MaintenanceEntryUpdateOne is the builder for updating a single MaintenanceEntry entity.
type MaintenanceEntryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MaintenanceEntryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (meuo *MaintenanceEntryUpdateOne) SetUpdatedAt(t time.Time) *MaintenanceEntryUpdateOne {
	meuo.mutation.SetUpdatedAt(t)
	return meuo
}

// SetItemID sets the "item_id" field.
func (meuo *MaintenanceEntryUpdateOne) SetItemID(u uuid.UUID) *MaintenanceEntryUpdateOne {
	meuo.mutation.SetItemID(u)
	return meuo
}

// SetDate sets the "date" field.
func (meuo *MaintenanceEntryUpdateOne) SetDate(t time.Time) *MaintenanceEntryUpdateOne {
	meuo.mutation.SetDate(t)
	return meuo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (meuo *MaintenanceEntryUpdateOne) SetNillableDate(t *time.Time) *MaintenanceEntryUpdateOne {
	if t != nil {
		meuo.SetDate(*t)
	}
	return meuo
}

// ClearDate clears the value of the "date" field.
func (meuo *MaintenanceEntryUpdateOne) ClearDate() *MaintenanceEntryUpdateOne {
	meuo.mutation.ClearDate()
	return meuo
}

// SetScheduledDate sets the "scheduled_date" field.
func (meuo *MaintenanceEntryUpdateOne) SetScheduledDate(t time.Time) *MaintenanceEntryUpdateOne {
	meuo.mutation.SetScheduledDate(t)
	return meuo
}

// SetNillableScheduledDate sets the "scheduled_date" field if the given value is not nil.
func (meuo *MaintenanceEntryUpdateOne) SetNillableScheduledDate(t *time.Time) *MaintenanceEntryUpdateOne {
	if t != nil {
		meuo.SetScheduledDate(*t)
	}
	return meuo
}

// ClearScheduledDate clears the value of the "scheduled_date" field.
func (meuo *MaintenanceEntryUpdateOne) ClearScheduledDate() *MaintenanceEntryUpdateOne {
	meuo.mutation.ClearScheduledDate()
	return meuo
}

// SetName sets the "name" field.
func (meuo *MaintenanceEntryUpdateOne) SetName(s string) *MaintenanceEntryUpdateOne {
	meuo.mutation.SetName(s)
	return meuo
}

// SetDescription sets the "description" field.
func (meuo *MaintenanceEntryUpdateOne) SetDescription(s string) *MaintenanceEntryUpdateOne {
	meuo.mutation.SetDescription(s)
	return meuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (meuo *MaintenanceEntryUpdateOne) SetNillableDescription(s *string) *MaintenanceEntryUpdateOne {
	if s != nil {
		meuo.SetDescription(*s)
	}
	return meuo
}

// ClearDescription clears the value of the "description" field.
func (meuo *MaintenanceEntryUpdateOne) ClearDescription() *MaintenanceEntryUpdateOne {
	meuo.mutation.ClearDescription()
	return meuo
}

// SetCost sets the "cost" field.
func (meuo *MaintenanceEntryUpdateOne) SetCost(f float64) *MaintenanceEntryUpdateOne {
	meuo.mutation.ResetCost()
	meuo.mutation.SetCost(f)
	return meuo
}

// SetNillableCost sets the "cost" field if the given value is not nil.
func (meuo *MaintenanceEntryUpdateOne) SetNillableCost(f *float64) *MaintenanceEntryUpdateOne {
	if f != nil {
		meuo.SetCost(*f)
	}
	return meuo
}

// AddCost adds f to the "cost" field.
func (meuo *MaintenanceEntryUpdateOne) AddCost(f float64) *MaintenanceEntryUpdateOne {
	meuo.mutation.AddCost(f)
	return meuo
}

// SetItem sets the "item" edge to the Item entity.
func (meuo *MaintenanceEntryUpdateOne) SetItem(i *Item) *MaintenanceEntryUpdateOne {
	return meuo.SetItemID(i.ID)
}

// Mutation returns the MaintenanceEntryMutation object of the builder.
func (meuo *MaintenanceEntryUpdateOne) Mutation() *MaintenanceEntryMutation {
	return meuo.mutation
}

// ClearItem clears the "item" edge to the Item entity.
func (meuo *MaintenanceEntryUpdateOne) ClearItem() *MaintenanceEntryUpdateOne {
	meuo.mutation.ClearItem()
	return meuo
}

// Where appends a list predicates to the MaintenanceEntryUpdate builder.
func (meuo *MaintenanceEntryUpdateOne) Where(ps ...predicate.MaintenanceEntry) *MaintenanceEntryUpdateOne {
	meuo.mutation.Where(ps...)
	return meuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (meuo *MaintenanceEntryUpdateOne) Select(field string, fields ...string) *MaintenanceEntryUpdateOne {
	meuo.fields = append([]string{field}, fields...)
	return meuo
}

// Save executes the query and returns the updated MaintenanceEntry entity.
func (meuo *MaintenanceEntryUpdateOne) Save(ctx context.Context) (*MaintenanceEntry, error) {
	meuo.defaults()
	return withHooks(ctx, meuo.sqlSave, meuo.mutation, meuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (meuo *MaintenanceEntryUpdateOne) SaveX(ctx context.Context) *MaintenanceEntry {
	node, err := meuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (meuo *MaintenanceEntryUpdateOne) Exec(ctx context.Context) error {
	_, err := meuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (meuo *MaintenanceEntryUpdateOne) ExecX(ctx context.Context) {
	if err := meuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (meuo *MaintenanceEntryUpdateOne) defaults() {
	if _, ok := meuo.mutation.UpdatedAt(); !ok {
		v := maintenanceentry.UpdateDefaultUpdatedAt()
		meuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (meuo *MaintenanceEntryUpdateOne) check() error {
	if v, ok := meuo.mutation.Name(); ok {
		if err := maintenanceentry.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MaintenanceEntry.name": %w`, err)}
		}
	}
	if v, ok := meuo.mutation.Description(); ok {
		if err := maintenanceentry.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "MaintenanceEntry.description": %w`, err)}
		}
	}
	if _, ok := meuo.mutation.ItemID(); meuo.mutation.ItemCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MaintenanceEntry.item"`)
	}
	return nil
}

func (meuo *MaintenanceEntryUpdateOne) sqlSave(ctx context.Context) (_node *MaintenanceEntry, err error) {
	if err := meuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(maintenanceentry.Table, maintenanceentry.Columns, sqlgraph.NewFieldSpec(maintenanceentry.FieldID, field.TypeUUID))
	id, ok := meuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MaintenanceEntry.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := meuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, maintenanceentry.FieldID)
		for _, f := range fields {
			if !maintenanceentry.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != maintenanceentry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := meuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := meuo.mutation.UpdatedAt(); ok {
		_spec.SetField(maintenanceentry.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := meuo.mutation.Date(); ok {
		_spec.SetField(maintenanceentry.FieldDate, field.TypeTime, value)
	}
	if meuo.mutation.DateCleared() {
		_spec.ClearField(maintenanceentry.FieldDate, field.TypeTime)
	}
	if value, ok := meuo.mutation.ScheduledDate(); ok {
		_spec.SetField(maintenanceentry.FieldScheduledDate, field.TypeTime, value)
	}
	if meuo.mutation.ScheduledDateCleared() {
		_spec.ClearField(maintenanceentry.FieldScheduledDate, field.TypeTime)
	}
	if value, ok := meuo.mutation.Name(); ok {
		_spec.SetField(maintenanceentry.FieldName, field.TypeString, value)
	}
	if value, ok := meuo.mutation.Description(); ok {
		_spec.SetField(maintenanceentry.FieldDescription, field.TypeString, value)
	}
	if meuo.mutation.DescriptionCleared() {
		_spec.ClearField(maintenanceentry.FieldDescription, field.TypeString)
	}
	if value, ok := meuo.mutation.Cost(); ok {
		_spec.SetField(maintenanceentry.FieldCost, field.TypeFloat64, value)
	}
	if value, ok := meuo.mutation.AddedCost(); ok {
		_spec.AddField(maintenanceentry.FieldCost, field.TypeFloat64, value)
	}
	if meuo.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   maintenanceentry.ItemTable,
			Columns: []string{maintenanceentry.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := meuo.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   maintenanceentry.ItemTable,
			Columns: []string{maintenanceentry.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MaintenanceEntry{config: meuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, meuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{maintenanceentry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	meuo.mutation.done = true
	return _node, nil
}
