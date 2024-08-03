// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/model"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/modelprice"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// ModelPriceQuery is the builder for querying ModelPrice entities.
type ModelPriceQuery struct {
	config
	ctx        *QueryContext
	order      []modelprice.OrderOption
	inters     []Interceptor
	predicates []predicate.ModelPrice
	withModel  *ModelQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ModelPriceQuery builder.
func (mpq *ModelPriceQuery) Where(ps ...predicate.ModelPrice) *ModelPriceQuery {
	mpq.predicates = append(mpq.predicates, ps...)
	return mpq
}

// Limit the number of records to be returned by this query.
func (mpq *ModelPriceQuery) Limit(limit int) *ModelPriceQuery {
	mpq.ctx.Limit = &limit
	return mpq
}

// Offset to start from.
func (mpq *ModelPriceQuery) Offset(offset int) *ModelPriceQuery {
	mpq.ctx.Offset = &offset
	return mpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mpq *ModelPriceQuery) Unique(unique bool) *ModelPriceQuery {
	mpq.ctx.Unique = &unique
	return mpq
}

// Order specifies how the records should be ordered.
func (mpq *ModelPriceQuery) Order(o ...modelprice.OrderOption) *ModelPriceQuery {
	mpq.order = append(mpq.order, o...)
	return mpq
}

// QueryModel chains the current query on the "model" edge.
func (mpq *ModelPriceQuery) QueryModel() *ModelQuery {
	query := (&ModelClient{config: mpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(modelprice.Table, modelprice.FieldID, selector),
			sqlgraph.To(model.Table, model.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, modelprice.ModelTable, modelprice.ModelColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ModelPrice entity from the query.
// Returns a *NotFoundError when no ModelPrice was found.
func (mpq *ModelPriceQuery) First(ctx context.Context) (*ModelPrice, error) {
	nodes, err := mpq.Limit(1).All(setContextOp(ctx, mpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{modelprice.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mpq *ModelPriceQuery) FirstX(ctx context.Context) *ModelPrice {
	node, err := mpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ModelPrice ID from the query.
// Returns a *NotFoundError when no ModelPrice ID was found.
func (mpq *ModelPriceQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mpq.Limit(1).IDs(setContextOp(ctx, mpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{modelprice.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mpq *ModelPriceQuery) FirstIDX(ctx context.Context) int64 {
	id, err := mpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ModelPrice entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ModelPrice entity is found.
// Returns a *NotFoundError when no ModelPrice entities are found.
func (mpq *ModelPriceQuery) Only(ctx context.Context) (*ModelPrice, error) {
	nodes, err := mpq.Limit(2).All(setContextOp(ctx, mpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{modelprice.Label}
	default:
		return nil, &NotSingularError{modelprice.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mpq *ModelPriceQuery) OnlyX(ctx context.Context) *ModelPrice {
	node, err := mpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ModelPrice ID in the query.
// Returns a *NotSingularError when more than one ModelPrice ID is found.
// Returns a *NotFoundError when no entities are found.
func (mpq *ModelPriceQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mpq.Limit(2).IDs(setContextOp(ctx, mpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{modelprice.Label}
	default:
		err = &NotSingularError{modelprice.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mpq *ModelPriceQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := mpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ModelPrices.
func (mpq *ModelPriceQuery) All(ctx context.Context) ([]*ModelPrice, error) {
	ctx = setContextOp(ctx, mpq.ctx, "All")
	if err := mpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ModelPrice, *ModelPriceQuery]()
	return withInterceptors[[]*ModelPrice](ctx, mpq, qr, mpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mpq *ModelPriceQuery) AllX(ctx context.Context) []*ModelPrice {
	nodes, err := mpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ModelPrice IDs.
func (mpq *ModelPriceQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if mpq.ctx.Unique == nil && mpq.path != nil {
		mpq.Unique(true)
	}
	ctx = setContextOp(ctx, mpq.ctx, "IDs")
	if err = mpq.Select(modelprice.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mpq *ModelPriceQuery) IDsX(ctx context.Context) []int64 {
	ids, err := mpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mpq *ModelPriceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mpq.ctx, "Count")
	if err := mpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mpq, querierCount[*ModelPriceQuery](), mpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mpq *ModelPriceQuery) CountX(ctx context.Context) int {
	count, err := mpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mpq *ModelPriceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mpq.ctx, "Exist")
	switch _, err := mpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mpq *ModelPriceQuery) ExistX(ctx context.Context) bool {
	exist, err := mpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ModelPriceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mpq *ModelPriceQuery) Clone() *ModelPriceQuery {
	if mpq == nil {
		return nil
	}
	return &ModelPriceQuery{
		config:     mpq.config,
		ctx:        mpq.ctx.Clone(),
		order:      append([]modelprice.OrderOption{}, mpq.order...),
		inters:     append([]Interceptor{}, mpq.inters...),
		predicates: append([]predicate.ModelPrice{}, mpq.predicates...),
		withModel:  mpq.withModel.Clone(),
		// clone intermediate query.
		sql:  mpq.sql.Clone(),
		path: mpq.path,
	}
}

// WithModel tells the query-builder to eager-load the nodes that are connected to
// the "model" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *ModelPriceQuery) WithModel(opts ...func(*ModelQuery)) *ModelPriceQuery {
	query := (&ModelClient{config: mpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mpq.withModel = query
	return mpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by,string"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ModelPrice.Query().
//		GroupBy(modelprice.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (mpq *ModelPriceQuery) GroupBy(field string, fields ...string) *ModelPriceGroupBy {
	mpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ModelPriceGroupBy{build: mpq}
	grbuild.flds = &mpq.ctx.Fields
	grbuild.label = modelprice.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by,string"`
//	}
//
//	client.ModelPrice.Query().
//		Select(modelprice.FieldCreatedBy).
//		Scan(ctx, &v)
func (mpq *ModelPriceQuery) Select(fields ...string) *ModelPriceSelect {
	mpq.ctx.Fields = append(mpq.ctx.Fields, fields...)
	sbuild := &ModelPriceSelect{ModelPriceQuery: mpq}
	sbuild.label = modelprice.Label
	sbuild.flds, sbuild.scan = &mpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ModelPriceSelect configured with the given aggregations.
func (mpq *ModelPriceQuery) Aggregate(fns ...AggregateFunc) *ModelPriceSelect {
	return mpq.Select().Aggregate(fns...)
}

func (mpq *ModelPriceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mpq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mpq); err != nil {
				return err
			}
		}
	}
	for _, f := range mpq.ctx.Fields {
		if !modelprice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if mpq.path != nil {
		prev, err := mpq.path(ctx)
		if err != nil {
			return err
		}
		mpq.sql = prev
	}
	return nil
}

func (mpq *ModelPriceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ModelPrice, error) {
	var (
		nodes       = []*ModelPrice{}
		_spec       = mpq.querySpec()
		loadedTypes = [1]bool{
			mpq.withModel != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ModelPrice).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ModelPrice{config: mpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(mpq.modifiers) > 0 {
		_spec.Modifiers = mpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mpq.withModel; query != nil {
		if err := mpq.loadModel(ctx, query, nodes, nil,
			func(n *ModelPrice, e *Model) { n.Edges.Model = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mpq *ModelPriceQuery) loadModel(ctx context.Context, query *ModelQuery, nodes []*ModelPrice, init func(*ModelPrice), assign func(*ModelPrice, *Model)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*ModelPrice)
	for i := range nodes {
		fk := nodes[i].ModelID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(model.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "model_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mpq *ModelPriceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mpq.querySpec()
	if len(mpq.modifiers) > 0 {
		_spec.Modifiers = mpq.modifiers
	}
	_spec.Node.Columns = mpq.ctx.Fields
	if len(mpq.ctx.Fields) > 0 {
		_spec.Unique = mpq.ctx.Unique != nil && *mpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mpq.driver, _spec)
}

func (mpq *ModelPriceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(modelprice.Table, modelprice.Columns, sqlgraph.NewFieldSpec(modelprice.FieldID, field.TypeInt64))
	_spec.From = mpq.sql
	if unique := mpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mpq.path != nil {
		_spec.Unique = true
	}
	if fields := mpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, modelprice.FieldID)
		for i := range fields {
			if fields[i] != modelprice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mpq.withModel != nil {
			_spec.Node.AddColumnOnce(modelprice.FieldModelID)
		}
	}
	if ps := mpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mpq *ModelPriceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mpq.driver.Dialect())
	t1 := builder.Table(modelprice.Table)
	columns := mpq.ctx.Fields
	if len(columns) == 0 {
		columns = modelprice.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mpq.sql != nil {
		selector = mpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mpq.ctx.Unique != nil && *mpq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range mpq.modifiers {
		m(selector)
	}
	for _, p := range mpq.predicates {
		p(selector)
	}
	for _, p := range mpq.order {
		p(selector)
	}
	if offset := mpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mpq *ModelPriceQuery) Modify(modifiers ...func(s *sql.Selector)) *ModelPriceSelect {
	mpq.modifiers = append(mpq.modifiers, modifiers...)
	return mpq.Select()
}

// ModelPriceGroupBy is the group-by builder for ModelPrice entities.
type ModelPriceGroupBy struct {
	selector
	build *ModelPriceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mpgb *ModelPriceGroupBy) Aggregate(fns ...AggregateFunc) *ModelPriceGroupBy {
	mpgb.fns = append(mpgb.fns, fns...)
	return mpgb
}

// Scan applies the selector query and scans the result into the given value.
func (mpgb *ModelPriceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mpgb.build.ctx, "GroupBy")
	if err := mpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ModelPriceQuery, *ModelPriceGroupBy](ctx, mpgb.build, mpgb, mpgb.build.inters, v)
}

func (mpgb *ModelPriceGroupBy) sqlScan(ctx context.Context, root *ModelPriceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mpgb.fns))
	for _, fn := range mpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mpgb.flds)+len(mpgb.fns))
		for _, f := range *mpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ModelPriceSelect is the builder for selecting fields of ModelPrice entities.
type ModelPriceSelect struct {
	*ModelPriceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mps *ModelPriceSelect) Aggregate(fns ...AggregateFunc) *ModelPriceSelect {
	mps.fns = append(mps.fns, fns...)
	return mps
}

// Scan applies the selector query and scans the result into the given value.
func (mps *ModelPriceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mps.ctx, "Select")
	if err := mps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ModelPriceQuery, *ModelPriceSelect](ctx, mps.ModelPriceQuery, mps, mps.inters, v)
}

func (mps *ModelPriceSelect) sqlScan(ctx context.Context, root *ModelPriceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mps.fns))
	for _, fn := range mps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mps *ModelPriceSelect) Modify(modifiers ...func(s *sql.Selector)) *ModelPriceSelect {
	mps.modifiers = append(mps.modifiers, modifiers...)
	return mps
}