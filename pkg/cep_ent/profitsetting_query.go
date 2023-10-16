// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/profitsetting"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// ProfitSettingQuery is the builder for querying ProfitSetting entities.
type ProfitSettingQuery struct {
	config
	ctx        *QueryContext
	order      []profitsetting.OrderOption
	inters     []Interceptor
	predicates []predicate.ProfitSetting
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProfitSettingQuery builder.
func (psq *ProfitSettingQuery) Where(ps ...predicate.ProfitSetting) *ProfitSettingQuery {
	psq.predicates = append(psq.predicates, ps...)
	return psq
}

// Limit the number of records to be returned by this query.
func (psq *ProfitSettingQuery) Limit(limit int) *ProfitSettingQuery {
	psq.ctx.Limit = &limit
	return psq
}

// Offset to start from.
func (psq *ProfitSettingQuery) Offset(offset int) *ProfitSettingQuery {
	psq.ctx.Offset = &offset
	return psq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (psq *ProfitSettingQuery) Unique(unique bool) *ProfitSettingQuery {
	psq.ctx.Unique = &unique
	return psq
}

// Order specifies how the records should be ordered.
func (psq *ProfitSettingQuery) Order(o ...profitsetting.OrderOption) *ProfitSettingQuery {
	psq.order = append(psq.order, o...)
	return psq
}

// QueryUser chains the current query on the "user" edge.
func (psq *ProfitSettingQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: psq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profitsetting.Table, profitsetting.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, profitsetting.UserTable, profitsetting.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(psq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProfitSetting entity from the query.
// Returns a *NotFoundError when no ProfitSetting was found.
func (psq *ProfitSettingQuery) First(ctx context.Context) (*ProfitSetting, error) {
	nodes, err := psq.Limit(1).All(setContextOp(ctx, psq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{profitsetting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (psq *ProfitSettingQuery) FirstX(ctx context.Context) *ProfitSetting {
	node, err := psq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProfitSetting ID from the query.
// Returns a *NotFoundError when no ProfitSetting ID was found.
func (psq *ProfitSettingQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = psq.Limit(1).IDs(setContextOp(ctx, psq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{profitsetting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (psq *ProfitSettingQuery) FirstIDX(ctx context.Context) int64 {
	id, err := psq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProfitSetting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProfitSetting entity is found.
// Returns a *NotFoundError when no ProfitSetting entities are found.
func (psq *ProfitSettingQuery) Only(ctx context.Context) (*ProfitSetting, error) {
	nodes, err := psq.Limit(2).All(setContextOp(ctx, psq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{profitsetting.Label}
	default:
		return nil, &NotSingularError{profitsetting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (psq *ProfitSettingQuery) OnlyX(ctx context.Context) *ProfitSetting {
	node, err := psq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProfitSetting ID in the query.
// Returns a *NotSingularError when more than one ProfitSetting ID is found.
// Returns a *NotFoundError when no entities are found.
func (psq *ProfitSettingQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = psq.Limit(2).IDs(setContextOp(ctx, psq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{profitsetting.Label}
	default:
		err = &NotSingularError{profitsetting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (psq *ProfitSettingQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := psq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProfitSettings.
func (psq *ProfitSettingQuery) All(ctx context.Context) ([]*ProfitSetting, error) {
	ctx = setContextOp(ctx, psq.ctx, "All")
	if err := psq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProfitSetting, *ProfitSettingQuery]()
	return withInterceptors[[]*ProfitSetting](ctx, psq, qr, psq.inters)
}

// AllX is like All, but panics if an error occurs.
func (psq *ProfitSettingQuery) AllX(ctx context.Context) []*ProfitSetting {
	nodes, err := psq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProfitSetting IDs.
func (psq *ProfitSettingQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if psq.ctx.Unique == nil && psq.path != nil {
		psq.Unique(true)
	}
	ctx = setContextOp(ctx, psq.ctx, "IDs")
	if err = psq.Select(profitsetting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (psq *ProfitSettingQuery) IDsX(ctx context.Context) []int64 {
	ids, err := psq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (psq *ProfitSettingQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, psq.ctx, "Count")
	if err := psq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, psq, querierCount[*ProfitSettingQuery](), psq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (psq *ProfitSettingQuery) CountX(ctx context.Context) int {
	count, err := psq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (psq *ProfitSettingQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, psq.ctx, "Exist")
	switch _, err := psq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (psq *ProfitSettingQuery) ExistX(ctx context.Context) bool {
	exist, err := psq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProfitSettingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (psq *ProfitSettingQuery) Clone() *ProfitSettingQuery {
	if psq == nil {
		return nil
	}
	return &ProfitSettingQuery{
		config:     psq.config,
		ctx:        psq.ctx.Clone(),
		order:      append([]profitsetting.OrderOption{}, psq.order...),
		inters:     append([]Interceptor{}, psq.inters...),
		predicates: append([]predicate.ProfitSetting{}, psq.predicates...),
		withUser:   psq.withUser.Clone(),
		// clone intermediate query.
		sql:  psq.sql.Clone(),
		path: psq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (psq *ProfitSettingQuery) WithUser(opts ...func(*UserQuery)) *ProfitSettingQuery {
	query := (&UserClient{config: psq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	psq.withUser = query
	return psq
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
//	client.ProfitSetting.Query().
//		GroupBy(profitsetting.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (psq *ProfitSettingQuery) GroupBy(field string, fields ...string) *ProfitSettingGroupBy {
	psq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProfitSettingGroupBy{build: psq}
	grbuild.flds = &psq.ctx.Fields
	grbuild.label = profitsetting.Label
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
//	client.ProfitSetting.Query().
//		Select(profitsetting.FieldCreatedBy).
//		Scan(ctx, &v)
func (psq *ProfitSettingQuery) Select(fields ...string) *ProfitSettingSelect {
	psq.ctx.Fields = append(psq.ctx.Fields, fields...)
	sbuild := &ProfitSettingSelect{ProfitSettingQuery: psq}
	sbuild.label = profitsetting.Label
	sbuild.flds, sbuild.scan = &psq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProfitSettingSelect configured with the given aggregations.
func (psq *ProfitSettingQuery) Aggregate(fns ...AggregateFunc) *ProfitSettingSelect {
	return psq.Select().Aggregate(fns...)
}

func (psq *ProfitSettingQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range psq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, psq); err != nil {
				return err
			}
		}
	}
	for _, f := range psq.ctx.Fields {
		if !profitsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if psq.path != nil {
		prev, err := psq.path(ctx)
		if err != nil {
			return err
		}
		psq.sql = prev
	}
	return nil
}

func (psq *ProfitSettingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProfitSetting, error) {
	var (
		nodes       = []*ProfitSetting{}
		_spec       = psq.querySpec()
		loadedTypes = [1]bool{
			psq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProfitSetting).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProfitSetting{config: psq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, psq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := psq.withUser; query != nil {
		if err := psq.loadUser(ctx, query, nodes, nil,
			func(n *ProfitSetting, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (psq *ProfitSettingQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*ProfitSetting, init func(*ProfitSetting), assign func(*ProfitSetting, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*ProfitSetting)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (psq *ProfitSettingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := psq.querySpec()
	_spec.Node.Columns = psq.ctx.Fields
	if len(psq.ctx.Fields) > 0 {
		_spec.Unique = psq.ctx.Unique != nil && *psq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, psq.driver, _spec)
}

func (psq *ProfitSettingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(profitsetting.Table, profitsetting.Columns, sqlgraph.NewFieldSpec(profitsetting.FieldID, field.TypeInt64))
	_spec.From = psq.sql
	if unique := psq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if psq.path != nil {
		_spec.Unique = true
	}
	if fields := psq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profitsetting.FieldID)
		for i := range fields {
			if fields[i] != profitsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if psq.withUser != nil {
			_spec.Node.AddColumnOnce(profitsetting.FieldUserID)
		}
	}
	if ps := psq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := psq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := psq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := psq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (psq *ProfitSettingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(psq.driver.Dialect())
	t1 := builder.Table(profitsetting.Table)
	columns := psq.ctx.Fields
	if len(columns) == 0 {
		columns = profitsetting.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if psq.sql != nil {
		selector = psq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if psq.ctx.Unique != nil && *psq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range psq.predicates {
		p(selector)
	}
	for _, p := range psq.order {
		p(selector)
	}
	if offset := psq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := psq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProfitSettingGroupBy is the group-by builder for ProfitSetting entities.
type ProfitSettingGroupBy struct {
	selector
	build *ProfitSettingQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (psgb *ProfitSettingGroupBy) Aggregate(fns ...AggregateFunc) *ProfitSettingGroupBy {
	psgb.fns = append(psgb.fns, fns...)
	return psgb
}

// Scan applies the selector query and scans the result into the given value.
func (psgb *ProfitSettingGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, psgb.build.ctx, "GroupBy")
	if err := psgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProfitSettingQuery, *ProfitSettingGroupBy](ctx, psgb.build, psgb, psgb.build.inters, v)
}

func (psgb *ProfitSettingGroupBy) sqlScan(ctx context.Context, root *ProfitSettingQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(psgb.fns))
	for _, fn := range psgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*psgb.flds)+len(psgb.fns))
		for _, f := range *psgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*psgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := psgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProfitSettingSelect is the builder for selecting fields of ProfitSetting entities.
type ProfitSettingSelect struct {
	*ProfitSettingQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pss *ProfitSettingSelect) Aggregate(fns ...AggregateFunc) *ProfitSettingSelect {
	pss.fns = append(pss.fns, fns...)
	return pss
}

// Scan applies the selector query and scans the result into the given value.
func (pss *ProfitSettingSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pss.ctx, "Select")
	if err := pss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProfitSettingQuery, *ProfitSettingSelect](ctx, pss.ProfitSettingQuery, pss, pss.inters, v)
}

func (pss *ProfitSettingSelect) sqlScan(ctx context.Context, root *ProfitSettingQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pss.fns))
	for _, fn := range pss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
