// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/attachment"
	"github.com/hay-kot/homebox/backend/internal/data/ent/document"
	"github.com/hay-kot/homebox/backend/internal/data/ent/item"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
)

// AttachmentQuery is the builder for querying Attachment entities.
type AttachmentQuery struct {
	config
	ctx          *QueryContext
	order        []OrderFunc
	inters       []Interceptor
	predicates   []predicate.Attachment
	withItem     *ItemQuery
	withDocument *DocumentQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttachmentQuery builder.
func (aq *AttachmentQuery) Where(ps ...predicate.Attachment) *AttachmentQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AttachmentQuery) Limit(limit int) *AttachmentQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AttachmentQuery) Offset(offset int) *AttachmentQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AttachmentQuery) Unique(unique bool) *AttachmentQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AttachmentQuery) Order(o ...OrderFunc) *AttachmentQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryItem chains the current query on the "item" edge.
func (aq *AttachmentQuery) QueryItem() *ItemQuery {
	query := (&ItemClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attachment.Table, attachment.FieldID, selector),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attachment.ItemTable, attachment.ItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDocument chains the current query on the "document" edge.
func (aq *AttachmentQuery) QueryDocument() *DocumentQuery {
	query := (&DocumentClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attachment.Table, attachment.FieldID, selector),
			sqlgraph.To(document.Table, document.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attachment.DocumentTable, attachment.DocumentColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Attachment entity from the query.
// Returns a *NotFoundError when no Attachment was found.
func (aq *AttachmentQuery) First(ctx context.Context) (*Attachment, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attachment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AttachmentQuery) FirstX(ctx context.Context) *Attachment {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Attachment ID from the query.
// Returns a *NotFoundError when no Attachment ID was found.
func (aq *AttachmentQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attachment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AttachmentQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Attachment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Attachment entity is found.
// Returns a *NotFoundError when no Attachment entities are found.
func (aq *AttachmentQuery) Only(ctx context.Context) (*Attachment, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attachment.Label}
	default:
		return nil, &NotSingularError{attachment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AttachmentQuery) OnlyX(ctx context.Context) *Attachment {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Attachment ID in the query.
// Returns a *NotSingularError when more than one Attachment ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AttachmentQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attachment.Label}
	default:
		err = &NotSingularError{attachment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AttachmentQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Attachments.
func (aq *AttachmentQuery) All(ctx context.Context) ([]*Attachment, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Attachment, *AttachmentQuery]()
	return withInterceptors[[]*Attachment](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AttachmentQuery) AllX(ctx context.Context) []*Attachment {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Attachment IDs.
func (aq *AttachmentQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err := aq.Select(attachment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AttachmentQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AttachmentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AttachmentQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AttachmentQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AttachmentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AttachmentQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttachmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AttachmentQuery) Clone() *AttachmentQuery {
	if aq == nil {
		return nil
	}
	return &AttachmentQuery{
		config:       aq.config,
		ctx:          aq.ctx.Clone(),
		order:        append([]OrderFunc{}, aq.order...),
		inters:       append([]Interceptor{}, aq.inters...),
		predicates:   append([]predicate.Attachment{}, aq.predicates...),
		withItem:     aq.withItem.Clone(),
		withDocument: aq.withDocument.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithItem tells the query-builder to eager-load the nodes that are connected to
// the "item" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttachmentQuery) WithItem(opts ...func(*ItemQuery)) *AttachmentQuery {
	query := (&ItemClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withItem = query
	return aq
}

// WithDocument tells the query-builder to eager-load the nodes that are connected to
// the "document" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttachmentQuery) WithDocument(opts ...func(*DocumentQuery)) *AttachmentQuery {
	query := (&DocumentClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withDocument = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Attachment.Query().
//		GroupBy(attachment.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AttachmentQuery) GroupBy(field string, fields ...string) *AttachmentGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AttachmentGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = attachment.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Attachment.Query().
//		Select(attachment.FieldCreatedAt).
//		Scan(ctx, &v)
func (aq *AttachmentQuery) Select(fields ...string) *AttachmentSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AttachmentSelect{AttachmentQuery: aq}
	sbuild.label = attachment.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AttachmentSelect configured with the given aggregations.
func (aq *AttachmentQuery) Aggregate(fns ...AggregateFunc) *AttachmentSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AttachmentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !attachment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AttachmentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Attachment, error) {
	var (
		nodes       = []*Attachment{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withItem != nil,
			aq.withDocument != nil,
		}
	)
	if aq.withItem != nil || aq.withDocument != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, attachment.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Attachment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Attachment{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withItem; query != nil {
		if err := aq.loadItem(ctx, query, nodes, nil,
			func(n *Attachment, e *Item) { n.Edges.Item = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withDocument; query != nil {
		if err := aq.loadDocument(ctx, query, nodes, nil,
			func(n *Attachment, e *Document) { n.Edges.Document = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AttachmentQuery) loadItem(ctx context.Context, query *ItemQuery, nodes []*Attachment, init func(*Attachment), assign func(*Attachment, *Item)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Attachment)
	for i := range nodes {
		if nodes[i].item_attachments == nil {
			continue
		}
		fk := *nodes[i].item_attachments
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(item.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "item_attachments" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AttachmentQuery) loadDocument(ctx context.Context, query *DocumentQuery, nodes []*Attachment, init func(*Attachment), assign func(*Attachment, *Document)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Attachment)
	for i := range nodes {
		if nodes[i].document_attachments == nil {
			continue
		}
		fk := *nodes[i].document_attachments
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(document.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "document_attachments" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aq *AttachmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AttachmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   attachment.Table,
			Columns: attachment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: attachment.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attachment.FieldID)
		for i := range fields {
			if fields[i] != attachment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AttachmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(attachment.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = attachment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AttachmentGroupBy is the group-by builder for Attachment entities.
type AttachmentGroupBy struct {
	selector
	build *AttachmentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AttachmentGroupBy) Aggregate(fns ...AggregateFunc) *AttachmentGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AttachmentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttachmentQuery, *AttachmentGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AttachmentGroupBy) sqlScan(ctx context.Context, root *AttachmentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AttachmentSelect is the builder for selecting fields of Attachment entities.
type AttachmentSelect struct {
	*AttachmentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AttachmentSelect) Aggregate(fns ...AggregateFunc) *AttachmentSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AttachmentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttachmentQuery, *AttachmentSelect](ctx, as.AttachmentQuery, as, as.inters, v)
}

func (as *AttachmentSelect) sqlScan(ctx context.Context, root *AttachmentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
