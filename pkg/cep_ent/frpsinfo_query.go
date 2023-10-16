// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/frpcinfo"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/frpsinfo"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// FrpsInfoQuery is the builder for querying FrpsInfo entities.
type FrpsInfoQuery struct {
	config
	ctx           *QueryContext
	order         []frpsinfo.OrderOption
	inters        []Interceptor
	predicates    []predicate.FrpsInfo
	withFrpcInfos *FrpcInfoQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FrpsInfoQuery builder.
func (fiq *FrpsInfoQuery) Where(ps ...predicate.FrpsInfo) *FrpsInfoQuery {
	fiq.predicates = append(fiq.predicates, ps...)
	return fiq
}

// Limit the number of records to be returned by this query.
func (fiq *FrpsInfoQuery) Limit(limit int) *FrpsInfoQuery {
	fiq.ctx.Limit = &limit
	return fiq
}

// Offset to start from.
func (fiq *FrpsInfoQuery) Offset(offset int) *FrpsInfoQuery {
	fiq.ctx.Offset = &offset
	return fiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fiq *FrpsInfoQuery) Unique(unique bool) *FrpsInfoQuery {
	fiq.ctx.Unique = &unique
	return fiq
}

// Order specifies how the records should be ordered.
func (fiq *FrpsInfoQuery) Order(o ...frpsinfo.OrderOption) *FrpsInfoQuery {
	fiq.order = append(fiq.order, o...)
	return fiq
}

// QueryFrpcInfos chains the current query on the "frpc_infos" edge.
func (fiq *FrpsInfoQuery) QueryFrpcInfos() *FrpcInfoQuery {
	query := (&FrpcInfoClient{config: fiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(frpsinfo.Table, frpsinfo.FieldID, selector),
			sqlgraph.To(frpcinfo.Table, frpcinfo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, frpsinfo.FrpcInfosTable, frpsinfo.FrpcInfosColumn),
		)
		fromU = sqlgraph.SetNeighbors(fiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FrpsInfo entity from the query.
// Returns a *NotFoundError when no FrpsInfo was found.
func (fiq *FrpsInfoQuery) First(ctx context.Context) (*FrpsInfo, error) {
	nodes, err := fiq.Limit(1).All(setContextOp(ctx, fiq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{frpsinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fiq *FrpsInfoQuery) FirstX(ctx context.Context) *FrpsInfo {
	node, err := fiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FrpsInfo ID from the query.
// Returns a *NotFoundError when no FrpsInfo ID was found.
func (fiq *FrpsInfoQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = fiq.Limit(1).IDs(setContextOp(ctx, fiq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{frpsinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fiq *FrpsInfoQuery) FirstIDX(ctx context.Context) int64 {
	id, err := fiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FrpsInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FrpsInfo entity is found.
// Returns a *NotFoundError when no FrpsInfo entities are found.
func (fiq *FrpsInfoQuery) Only(ctx context.Context) (*FrpsInfo, error) {
	nodes, err := fiq.Limit(2).All(setContextOp(ctx, fiq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{frpsinfo.Label}
	default:
		return nil, &NotSingularError{frpsinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fiq *FrpsInfoQuery) OnlyX(ctx context.Context) *FrpsInfo {
	node, err := fiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FrpsInfo ID in the query.
// Returns a *NotSingularError when more than one FrpsInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (fiq *FrpsInfoQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = fiq.Limit(2).IDs(setContextOp(ctx, fiq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{frpsinfo.Label}
	default:
		err = &NotSingularError{frpsinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fiq *FrpsInfoQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := fiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FrpsInfos.
func (fiq *FrpsInfoQuery) All(ctx context.Context) ([]*FrpsInfo, error) {
	ctx = setContextOp(ctx, fiq.ctx, "All")
	if err := fiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FrpsInfo, *FrpsInfoQuery]()
	return withInterceptors[[]*FrpsInfo](ctx, fiq, qr, fiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fiq *FrpsInfoQuery) AllX(ctx context.Context) []*FrpsInfo {
	nodes, err := fiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FrpsInfo IDs.
func (fiq *FrpsInfoQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if fiq.ctx.Unique == nil && fiq.path != nil {
		fiq.Unique(true)
	}
	ctx = setContextOp(ctx, fiq.ctx, "IDs")
	if err = fiq.Select(frpsinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fiq *FrpsInfoQuery) IDsX(ctx context.Context) []int64 {
	ids, err := fiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fiq *FrpsInfoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fiq.ctx, "Count")
	if err := fiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fiq, querierCount[*FrpsInfoQuery](), fiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fiq *FrpsInfoQuery) CountX(ctx context.Context) int {
	count, err := fiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fiq *FrpsInfoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fiq.ctx, "Exist")
	switch _, err := fiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fiq *FrpsInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := fiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FrpsInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fiq *FrpsInfoQuery) Clone() *FrpsInfoQuery {
	if fiq == nil {
		return nil
	}
	return &FrpsInfoQuery{
		config:        fiq.config,
		ctx:           fiq.ctx.Clone(),
		order:         append([]frpsinfo.OrderOption{}, fiq.order...),
		inters:        append([]Interceptor{}, fiq.inters...),
		predicates:    append([]predicate.FrpsInfo{}, fiq.predicates...),
		withFrpcInfos: fiq.withFrpcInfos.Clone(),
		// clone intermediate query.
		sql:  fiq.sql.Clone(),
		path: fiq.path,
	}
}

// WithFrpcInfos tells the query-builder to eager-load the nodes that are connected to
// the "frpc_infos" edge. The optional arguments are used to configure the query builder of the edge.
func (fiq *FrpsInfoQuery) WithFrpcInfos(opts ...func(*FrpcInfoQuery)) *FrpsInfoQuery {
	query := (&FrpcInfoClient{config: fiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fiq.withFrpcInfos = query
	return fiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FrpsInfo.Query().
//		GroupBy(frpsinfo.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (fiq *FrpsInfoQuery) GroupBy(field string, fields ...string) *FrpsInfoGroupBy {
	fiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FrpsInfoGroupBy{build: fiq}
	grbuild.flds = &fiq.ctx.Fields
	grbuild.label = frpsinfo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by"`
//	}
//
//	client.FrpsInfo.Query().
//		Select(frpsinfo.FieldCreatedBy).
//		Scan(ctx, &v)
func (fiq *FrpsInfoQuery) Select(fields ...string) *FrpsInfoSelect {
	fiq.ctx.Fields = append(fiq.ctx.Fields, fields...)
	sbuild := &FrpsInfoSelect{FrpsInfoQuery: fiq}
	sbuild.label = frpsinfo.Label
	sbuild.flds, sbuild.scan = &fiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FrpsInfoSelect configured with the given aggregations.
func (fiq *FrpsInfoQuery) Aggregate(fns ...AggregateFunc) *FrpsInfoSelect {
	return fiq.Select().Aggregate(fns...)
}

func (fiq *FrpsInfoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fiq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fiq); err != nil {
				return err
			}
		}
	}
	for _, f := range fiq.ctx.Fields {
		if !frpsinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if fiq.path != nil {
		prev, err := fiq.path(ctx)
		if err != nil {
			return err
		}
		fiq.sql = prev
	}
	return nil
}

func (fiq *FrpsInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FrpsInfo, error) {
	var (
		nodes       = []*FrpsInfo{}
		_spec       = fiq.querySpec()
		loadedTypes = [1]bool{
			fiq.withFrpcInfos != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FrpsInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FrpsInfo{config: fiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fiq.withFrpcInfos; query != nil {
		if err := fiq.loadFrpcInfos(ctx, query, nodes,
			func(n *FrpsInfo) { n.Edges.FrpcInfos = []*FrpcInfo{} },
			func(n *FrpsInfo, e *FrpcInfo) { n.Edges.FrpcInfos = append(n.Edges.FrpcInfos, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fiq *FrpsInfoQuery) loadFrpcInfos(ctx context.Context, query *FrpcInfoQuery, nodes []*FrpsInfo, init func(*FrpsInfo), assign func(*FrpsInfo, *FrpcInfo)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*FrpsInfo)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(frpcinfo.FieldFrpsID)
	}
	query.Where(predicate.FrpcInfo(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(frpsinfo.FrpcInfosColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.FrpsID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "frps_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (fiq *FrpsInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fiq.querySpec()
	_spec.Node.Columns = fiq.ctx.Fields
	if len(fiq.ctx.Fields) > 0 {
		_spec.Unique = fiq.ctx.Unique != nil && *fiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fiq.driver, _spec)
}

func (fiq *FrpsInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(frpsinfo.Table, frpsinfo.Columns, sqlgraph.NewFieldSpec(frpsinfo.FieldID, field.TypeInt64))
	_spec.From = fiq.sql
	if unique := fiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fiq.path != nil {
		_spec.Unique = true
	}
	if fields := fiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, frpsinfo.FieldID)
		for i := range fields {
			if fields[i] != frpsinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fiq *FrpsInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fiq.driver.Dialect())
	t1 := builder.Table(frpsinfo.Table)
	columns := fiq.ctx.Fields
	if len(columns) == 0 {
		columns = frpsinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fiq.sql != nil {
		selector = fiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fiq.ctx.Unique != nil && *fiq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fiq.predicates {
		p(selector)
	}
	for _, p := range fiq.order {
		p(selector)
	}
	if offset := fiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FrpsInfoGroupBy is the group-by builder for FrpsInfo entities.
type FrpsInfoGroupBy struct {
	selector
	build *FrpsInfoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (figb *FrpsInfoGroupBy) Aggregate(fns ...AggregateFunc) *FrpsInfoGroupBy {
	figb.fns = append(figb.fns, fns...)
	return figb
}

// Scan applies the selector query and scans the result into the given value.
func (figb *FrpsInfoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, figb.build.ctx, "GroupBy")
	if err := figb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FrpsInfoQuery, *FrpsInfoGroupBy](ctx, figb.build, figb, figb.build.inters, v)
}

func (figb *FrpsInfoGroupBy) sqlScan(ctx context.Context, root *FrpsInfoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(figb.fns))
	for _, fn := range figb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*figb.flds)+len(figb.fns))
		for _, f := range *figb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*figb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := figb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FrpsInfoSelect is the builder for selecting fields of FrpsInfo entities.
type FrpsInfoSelect struct {
	*FrpsInfoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fis *FrpsInfoSelect) Aggregate(fns ...AggregateFunc) *FrpsInfoSelect {
	fis.fns = append(fis.fns, fns...)
	return fis
}

// Scan applies the selector query and scans the result into the given value.
func (fis *FrpsInfoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fis.ctx, "Select")
	if err := fis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FrpsInfoQuery, *FrpsInfoSelect](ctx, fis.FrpsInfoQuery, fis, fis.inters, v)
}

func (fis *FrpsInfoSelect) sqlScan(ctx context.Context, root *FrpsInfoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fis.fns))
	for _, fn := range fis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
