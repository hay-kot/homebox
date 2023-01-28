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
	"github.com/hay-kot/homebox/backend/internal/data/ent/authroles"
	"github.com/hay-kot/homebox/backend/internal/data/ent/authtokens"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
)

// AuthRolesQuery is the builder for querying AuthRoles entities.
type AuthRolesQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.AuthRoles
	withToken  *AuthTokensQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthRolesQuery builder.
func (arq *AuthRolesQuery) Where(ps ...predicate.AuthRoles) *AuthRolesQuery {
	arq.predicates = append(arq.predicates, ps...)
	return arq
}

// Limit the number of records to be returned by this query.
func (arq *AuthRolesQuery) Limit(limit int) *AuthRolesQuery {
	arq.ctx.Limit = &limit
	return arq
}

// Offset to start from.
func (arq *AuthRolesQuery) Offset(offset int) *AuthRolesQuery {
	arq.ctx.Offset = &offset
	return arq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (arq *AuthRolesQuery) Unique(unique bool) *AuthRolesQuery {
	arq.ctx.Unique = &unique
	return arq
}

// Order specifies how the records should be ordered.
func (arq *AuthRolesQuery) Order(o ...OrderFunc) *AuthRolesQuery {
	arq.order = append(arq.order, o...)
	return arq
}

// QueryToken chains the current query on the "token" edge.
func (arq *AuthRolesQuery) QueryToken() *AuthTokensQuery {
	query := (&AuthTokensClient{config: arq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := arq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(authroles.Table, authroles.FieldID, selector),
			sqlgraph.To(authtokens.Table, authtokens.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, authroles.TokenTable, authroles.TokenColumn),
		)
		fromU = sqlgraph.SetNeighbors(arq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AuthRoles entity from the query.
// Returns a *NotFoundError when no AuthRoles was found.
func (arq *AuthRolesQuery) First(ctx context.Context) (*AuthRoles, error) {
	nodes, err := arq.Limit(1).All(setContextOp(ctx, arq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authroles.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (arq *AuthRolesQuery) FirstX(ctx context.Context) *AuthRoles {
	node, err := arq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthRoles ID from the query.
// Returns a *NotFoundError when no AuthRoles ID was found.
func (arq *AuthRolesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = arq.Limit(1).IDs(setContextOp(ctx, arq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authroles.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (arq *AuthRolesQuery) FirstIDX(ctx context.Context) int {
	id, err := arq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AuthRoles entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AuthRoles entity is found.
// Returns a *NotFoundError when no AuthRoles entities are found.
func (arq *AuthRolesQuery) Only(ctx context.Context) (*AuthRoles, error) {
	nodes, err := arq.Limit(2).All(setContextOp(ctx, arq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authroles.Label}
	default:
		return nil, &NotSingularError{authroles.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (arq *AuthRolesQuery) OnlyX(ctx context.Context) *AuthRoles {
	node, err := arq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AuthRoles ID in the query.
// Returns a *NotSingularError when more than one AuthRoles ID is found.
// Returns a *NotFoundError when no entities are found.
func (arq *AuthRolesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = arq.Limit(2).IDs(setContextOp(ctx, arq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authroles.Label}
	default:
		err = &NotSingularError{authroles.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (arq *AuthRolesQuery) OnlyIDX(ctx context.Context) int {
	id, err := arq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthRolesSlice.
func (arq *AuthRolesQuery) All(ctx context.Context) ([]*AuthRoles, error) {
	ctx = setContextOp(ctx, arq.ctx, "All")
	if err := arq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AuthRoles, *AuthRolesQuery]()
	return withInterceptors[[]*AuthRoles](ctx, arq, qr, arq.inters)
}

// AllX is like All, but panics if an error occurs.
func (arq *AuthRolesQuery) AllX(ctx context.Context) []*AuthRoles {
	nodes, err := arq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthRoles IDs.
func (arq *AuthRolesQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = setContextOp(ctx, arq.ctx, "IDs")
	if err := arq.Select(authroles.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (arq *AuthRolesQuery) IDsX(ctx context.Context) []int {
	ids, err := arq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (arq *AuthRolesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, arq.ctx, "Count")
	if err := arq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, arq, querierCount[*AuthRolesQuery](), arq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (arq *AuthRolesQuery) CountX(ctx context.Context) int {
	count, err := arq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (arq *AuthRolesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, arq.ctx, "Exist")
	switch _, err := arq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (arq *AuthRolesQuery) ExistX(ctx context.Context) bool {
	exist, err := arq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthRolesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (arq *AuthRolesQuery) Clone() *AuthRolesQuery {
	if arq == nil {
		return nil
	}
	return &AuthRolesQuery{
		config:     arq.config,
		ctx:        arq.ctx.Clone(),
		order:      append([]OrderFunc{}, arq.order...),
		inters:     append([]Interceptor{}, arq.inters...),
		predicates: append([]predicate.AuthRoles{}, arq.predicates...),
		withToken:  arq.withToken.Clone(),
		// clone intermediate query.
		sql:  arq.sql.Clone(),
		path: arq.path,
	}
}

// WithToken tells the query-builder to eager-load the nodes that are connected to
// the "token" edge. The optional arguments are used to configure the query builder of the edge.
func (arq *AuthRolesQuery) WithToken(opts ...func(*AuthTokensQuery)) *AuthRolesQuery {
	query := (&AuthTokensClient{config: arq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	arq.withToken = query
	return arq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Role authroles.Role `json:"role,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AuthRoles.Query().
//		GroupBy(authroles.FieldRole).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (arq *AuthRolesQuery) GroupBy(field string, fields ...string) *AuthRolesGroupBy {
	arq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AuthRolesGroupBy{build: arq}
	grbuild.flds = &arq.ctx.Fields
	grbuild.label = authroles.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Role authroles.Role `json:"role,omitempty"`
//	}
//
//	client.AuthRoles.Query().
//		Select(authroles.FieldRole).
//		Scan(ctx, &v)
func (arq *AuthRolesQuery) Select(fields ...string) *AuthRolesSelect {
	arq.ctx.Fields = append(arq.ctx.Fields, fields...)
	sbuild := &AuthRolesSelect{AuthRolesQuery: arq}
	sbuild.label = authroles.Label
	sbuild.flds, sbuild.scan = &arq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AuthRolesSelect configured with the given aggregations.
func (arq *AuthRolesQuery) Aggregate(fns ...AggregateFunc) *AuthRolesSelect {
	return arq.Select().Aggregate(fns...)
}

func (arq *AuthRolesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range arq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, arq); err != nil {
				return err
			}
		}
	}
	for _, f := range arq.ctx.Fields {
		if !authroles.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if arq.path != nil {
		prev, err := arq.path(ctx)
		if err != nil {
			return err
		}
		arq.sql = prev
	}
	return nil
}

func (arq *AuthRolesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AuthRoles, error) {
	var (
		nodes       = []*AuthRoles{}
		withFKs     = arq.withFKs
		_spec       = arq.querySpec()
		loadedTypes = [1]bool{
			arq.withToken != nil,
		}
	)
	if arq.withToken != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, authroles.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AuthRoles).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AuthRoles{config: arq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, arq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := arq.withToken; query != nil {
		if err := arq.loadToken(ctx, query, nodes, nil,
			func(n *AuthRoles, e *AuthTokens) { n.Edges.Token = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (arq *AuthRolesQuery) loadToken(ctx context.Context, query *AuthTokensQuery, nodes []*AuthRoles, init func(*AuthRoles), assign func(*AuthRoles, *AuthTokens)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*AuthRoles)
	for i := range nodes {
		if nodes[i].auth_tokens_roles == nil {
			continue
		}
		fk := *nodes[i].auth_tokens_roles
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(authtokens.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "auth_tokens_roles" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (arq *AuthRolesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := arq.querySpec()
	_spec.Node.Columns = arq.ctx.Fields
	if len(arq.ctx.Fields) > 0 {
		_spec.Unique = arq.ctx.Unique != nil && *arq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, arq.driver, _spec)
}

func (arq *AuthRolesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authroles.Table,
			Columns: authroles.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authroles.FieldID,
			},
		},
		From:   arq.sql,
		Unique: true,
	}
	if unique := arq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := arq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authroles.FieldID)
		for i := range fields {
			if fields[i] != authroles.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := arq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := arq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := arq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := arq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (arq *AuthRolesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(arq.driver.Dialect())
	t1 := builder.Table(authroles.Table)
	columns := arq.ctx.Fields
	if len(columns) == 0 {
		columns = authroles.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if arq.sql != nil {
		selector = arq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if arq.ctx.Unique != nil && *arq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range arq.predicates {
		p(selector)
	}
	for _, p := range arq.order {
		p(selector)
	}
	if offset := arq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := arq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthRolesGroupBy is the group-by builder for AuthRoles entities.
type AuthRolesGroupBy struct {
	selector
	build *AuthRolesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (argb *AuthRolesGroupBy) Aggregate(fns ...AggregateFunc) *AuthRolesGroupBy {
	argb.fns = append(argb.fns, fns...)
	return argb
}

// Scan applies the selector query and scans the result into the given value.
func (argb *AuthRolesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, argb.build.ctx, "GroupBy")
	if err := argb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AuthRolesQuery, *AuthRolesGroupBy](ctx, argb.build, argb, argb.build.inters, v)
}

func (argb *AuthRolesGroupBy) sqlScan(ctx context.Context, root *AuthRolesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(argb.fns))
	for _, fn := range argb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*argb.flds)+len(argb.fns))
		for _, f := range *argb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*argb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := argb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AuthRolesSelect is the builder for selecting fields of AuthRoles entities.
type AuthRolesSelect struct {
	*AuthRolesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ars *AuthRolesSelect) Aggregate(fns ...AggregateFunc) *AuthRolesSelect {
	ars.fns = append(ars.fns, fns...)
	return ars
}

// Scan applies the selector query and scans the result into the given value.
func (ars *AuthRolesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ars.ctx, "Select")
	if err := ars.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AuthRolesQuery, *AuthRolesSelect](ctx, ars.AuthRolesQuery, ars, ars.inters, v)
}

func (ars *AuthRolesSelect) sqlScan(ctx context.Context, root *AuthRolesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ars.fns))
	for _, fn := range ars.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ars.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ars.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
