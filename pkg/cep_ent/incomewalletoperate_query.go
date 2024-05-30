// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/incomewalletoperate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// IncomeWalletOperateQuery is the builder for querying IncomeWalletOperate entities.
type IncomeWalletOperateQuery struct {
	config
	ctx             *QueryContext
	order           []incomewalletoperate.OrderOption
	inters          []Interceptor
	predicates      []predicate.IncomeWalletOperate
	withUser        *UserQuery
	withApproveUser *UserQuery
	modifiers       []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IncomeWalletOperateQuery builder.
func (iwoq *IncomeWalletOperateQuery) Where(ps ...predicate.IncomeWalletOperate) *IncomeWalletOperateQuery {
	iwoq.predicates = append(iwoq.predicates, ps...)
	return iwoq
}

// Limit the number of records to be returned by this query.
func (iwoq *IncomeWalletOperateQuery) Limit(limit int) *IncomeWalletOperateQuery {
	iwoq.ctx.Limit = &limit
	return iwoq
}

// Offset to start from.
func (iwoq *IncomeWalletOperateQuery) Offset(offset int) *IncomeWalletOperateQuery {
	iwoq.ctx.Offset = &offset
	return iwoq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iwoq *IncomeWalletOperateQuery) Unique(unique bool) *IncomeWalletOperateQuery {
	iwoq.ctx.Unique = &unique
	return iwoq
}

// Order specifies how the records should be ordered.
func (iwoq *IncomeWalletOperateQuery) Order(o ...incomewalletoperate.OrderOption) *IncomeWalletOperateQuery {
	iwoq.order = append(iwoq.order, o...)
	return iwoq
}

// QueryUser chains the current query on the "user" edge.
func (iwoq *IncomeWalletOperateQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: iwoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iwoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iwoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(incomewalletoperate.Table, incomewalletoperate.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, incomewalletoperate.UserTable, incomewalletoperate.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(iwoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryApproveUser chains the current query on the "approve_user" edge.
func (iwoq *IncomeWalletOperateQuery) QueryApproveUser() *UserQuery {
	query := (&UserClient{config: iwoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iwoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iwoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(incomewalletoperate.Table, incomewalletoperate.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, incomewalletoperate.ApproveUserTable, incomewalletoperate.ApproveUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(iwoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first IncomeWalletOperate entity from the query.
// Returns a *NotFoundError when no IncomeWalletOperate was found.
func (iwoq *IncomeWalletOperateQuery) First(ctx context.Context) (*IncomeWalletOperate, error) {
	nodes, err := iwoq.Limit(1).All(setContextOp(ctx, iwoq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{incomewalletoperate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iwoq *IncomeWalletOperateQuery) FirstX(ctx context.Context) *IncomeWalletOperate {
	node, err := iwoq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first IncomeWalletOperate ID from the query.
// Returns a *NotFoundError when no IncomeWalletOperate ID was found.
func (iwoq *IncomeWalletOperateQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = iwoq.Limit(1).IDs(setContextOp(ctx, iwoq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{incomewalletoperate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iwoq *IncomeWalletOperateQuery) FirstIDX(ctx context.Context) int64 {
	id, err := iwoq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single IncomeWalletOperate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one IncomeWalletOperate entity is found.
// Returns a *NotFoundError when no IncomeWalletOperate entities are found.
func (iwoq *IncomeWalletOperateQuery) Only(ctx context.Context) (*IncomeWalletOperate, error) {
	nodes, err := iwoq.Limit(2).All(setContextOp(ctx, iwoq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{incomewalletoperate.Label}
	default:
		return nil, &NotSingularError{incomewalletoperate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iwoq *IncomeWalletOperateQuery) OnlyX(ctx context.Context) *IncomeWalletOperate {
	node, err := iwoq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only IncomeWalletOperate ID in the query.
// Returns a *NotSingularError when more than one IncomeWalletOperate ID is found.
// Returns a *NotFoundError when no entities are found.
func (iwoq *IncomeWalletOperateQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = iwoq.Limit(2).IDs(setContextOp(ctx, iwoq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{incomewalletoperate.Label}
	default:
		err = &NotSingularError{incomewalletoperate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iwoq *IncomeWalletOperateQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := iwoq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of IncomeWalletOperates.
func (iwoq *IncomeWalletOperateQuery) All(ctx context.Context) ([]*IncomeWalletOperate, error) {
	ctx = setContextOp(ctx, iwoq.ctx, "All")
	if err := iwoq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*IncomeWalletOperate, *IncomeWalletOperateQuery]()
	return withInterceptors[[]*IncomeWalletOperate](ctx, iwoq, qr, iwoq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iwoq *IncomeWalletOperateQuery) AllX(ctx context.Context) []*IncomeWalletOperate {
	nodes, err := iwoq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of IncomeWalletOperate IDs.
func (iwoq *IncomeWalletOperateQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if iwoq.ctx.Unique == nil && iwoq.path != nil {
		iwoq.Unique(true)
	}
	ctx = setContextOp(ctx, iwoq.ctx, "IDs")
	if err = iwoq.Select(incomewalletoperate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iwoq *IncomeWalletOperateQuery) IDsX(ctx context.Context) []int64 {
	ids, err := iwoq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iwoq *IncomeWalletOperateQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iwoq.ctx, "Count")
	if err := iwoq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iwoq, querierCount[*IncomeWalletOperateQuery](), iwoq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iwoq *IncomeWalletOperateQuery) CountX(ctx context.Context) int {
	count, err := iwoq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iwoq *IncomeWalletOperateQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iwoq.ctx, "Exist")
	switch _, err := iwoq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iwoq *IncomeWalletOperateQuery) ExistX(ctx context.Context) bool {
	exist, err := iwoq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IncomeWalletOperateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iwoq *IncomeWalletOperateQuery) Clone() *IncomeWalletOperateQuery {
	if iwoq == nil {
		return nil
	}
	return &IncomeWalletOperateQuery{
		config:          iwoq.config,
		ctx:             iwoq.ctx.Clone(),
		order:           append([]incomewalletoperate.OrderOption{}, iwoq.order...),
		inters:          append([]Interceptor{}, iwoq.inters...),
		predicates:      append([]predicate.IncomeWalletOperate{}, iwoq.predicates...),
		withUser:        iwoq.withUser.Clone(),
		withApproveUser: iwoq.withApproveUser.Clone(),
		// clone intermediate query.
		sql:  iwoq.sql.Clone(),
		path: iwoq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (iwoq *IncomeWalletOperateQuery) WithUser(opts ...func(*UserQuery)) *IncomeWalletOperateQuery {
	query := (&UserClient{config: iwoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iwoq.withUser = query
	return iwoq
}

// WithApproveUser tells the query-builder to eager-load the nodes that are connected to
// the "approve_user" edge. The optional arguments are used to configure the query builder of the edge.
func (iwoq *IncomeWalletOperateQuery) WithApproveUser(opts ...func(*UserQuery)) *IncomeWalletOperateQuery {
	query := (&UserClient{config: iwoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iwoq.withApproveUser = query
	return iwoq
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
//	client.IncomeWalletOperate.Query().
//		GroupBy(incomewalletoperate.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (iwoq *IncomeWalletOperateQuery) GroupBy(field string, fields ...string) *IncomeWalletOperateGroupBy {
	iwoq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &IncomeWalletOperateGroupBy{build: iwoq}
	grbuild.flds = &iwoq.ctx.Fields
	grbuild.label = incomewalletoperate.Label
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
//	client.IncomeWalletOperate.Query().
//		Select(incomewalletoperate.FieldCreatedBy).
//		Scan(ctx, &v)
func (iwoq *IncomeWalletOperateQuery) Select(fields ...string) *IncomeWalletOperateSelect {
	iwoq.ctx.Fields = append(iwoq.ctx.Fields, fields...)
	sbuild := &IncomeWalletOperateSelect{IncomeWalletOperateQuery: iwoq}
	sbuild.label = incomewalletoperate.Label
	sbuild.flds, sbuild.scan = &iwoq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a IncomeWalletOperateSelect configured with the given aggregations.
func (iwoq *IncomeWalletOperateQuery) Aggregate(fns ...AggregateFunc) *IncomeWalletOperateSelect {
	return iwoq.Select().Aggregate(fns...)
}

func (iwoq *IncomeWalletOperateQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iwoq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iwoq); err != nil {
				return err
			}
		}
	}
	for _, f := range iwoq.ctx.Fields {
		if !incomewalletoperate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if iwoq.path != nil {
		prev, err := iwoq.path(ctx)
		if err != nil {
			return err
		}
		iwoq.sql = prev
	}
	return nil
}

func (iwoq *IncomeWalletOperateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*IncomeWalletOperate, error) {
	var (
		nodes       = []*IncomeWalletOperate{}
		_spec       = iwoq.querySpec()
		loadedTypes = [2]bool{
			iwoq.withUser != nil,
			iwoq.withApproveUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*IncomeWalletOperate).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &IncomeWalletOperate{config: iwoq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(iwoq.modifiers) > 0 {
		_spec.Modifiers = iwoq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iwoq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iwoq.withUser; query != nil {
		if err := iwoq.loadUser(ctx, query, nodes, nil,
			func(n *IncomeWalletOperate, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := iwoq.withApproveUser; query != nil {
		if err := iwoq.loadApproveUser(ctx, query, nodes, nil,
			func(n *IncomeWalletOperate, e *User) { n.Edges.ApproveUser = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iwoq *IncomeWalletOperateQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*IncomeWalletOperate, init func(*IncomeWalletOperate), assign func(*IncomeWalletOperate, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*IncomeWalletOperate)
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
func (iwoq *IncomeWalletOperateQuery) loadApproveUser(ctx context.Context, query *UserQuery, nodes []*IncomeWalletOperate, init func(*IncomeWalletOperate), assign func(*IncomeWalletOperate, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*IncomeWalletOperate)
	for i := range nodes {
		fk := nodes[i].ApproveUserID
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
			return fmt.Errorf(`unexpected foreign-key "approve_user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (iwoq *IncomeWalletOperateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iwoq.querySpec()
	if len(iwoq.modifiers) > 0 {
		_spec.Modifiers = iwoq.modifiers
	}
	_spec.Node.Columns = iwoq.ctx.Fields
	if len(iwoq.ctx.Fields) > 0 {
		_spec.Unique = iwoq.ctx.Unique != nil && *iwoq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iwoq.driver, _spec)
}

func (iwoq *IncomeWalletOperateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(incomewalletoperate.Table, incomewalletoperate.Columns, sqlgraph.NewFieldSpec(incomewalletoperate.FieldID, field.TypeInt64))
	_spec.From = iwoq.sql
	if unique := iwoq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iwoq.path != nil {
		_spec.Unique = true
	}
	if fields := iwoq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, incomewalletoperate.FieldID)
		for i := range fields {
			if fields[i] != incomewalletoperate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if iwoq.withUser != nil {
			_spec.Node.AddColumnOnce(incomewalletoperate.FieldUserID)
		}
		if iwoq.withApproveUser != nil {
			_spec.Node.AddColumnOnce(incomewalletoperate.FieldApproveUserID)
		}
	}
	if ps := iwoq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iwoq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iwoq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iwoq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iwoq *IncomeWalletOperateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iwoq.driver.Dialect())
	t1 := builder.Table(incomewalletoperate.Table)
	columns := iwoq.ctx.Fields
	if len(columns) == 0 {
		columns = incomewalletoperate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iwoq.sql != nil {
		selector = iwoq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iwoq.ctx.Unique != nil && *iwoq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range iwoq.modifiers {
		m(selector)
	}
	for _, p := range iwoq.predicates {
		p(selector)
	}
	for _, p := range iwoq.order {
		p(selector)
	}
	if offset := iwoq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iwoq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (iwoq *IncomeWalletOperateQuery) Modify(modifiers ...func(s *sql.Selector)) *IncomeWalletOperateSelect {
	iwoq.modifiers = append(iwoq.modifiers, modifiers...)
	return iwoq.Select()
}

// IncomeWalletOperateGroupBy is the group-by builder for IncomeWalletOperate entities.
type IncomeWalletOperateGroupBy struct {
	selector
	build *IncomeWalletOperateQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (iwogb *IncomeWalletOperateGroupBy) Aggregate(fns ...AggregateFunc) *IncomeWalletOperateGroupBy {
	iwogb.fns = append(iwogb.fns, fns...)
	return iwogb
}

// Scan applies the selector query and scans the result into the given value.
func (iwogb *IncomeWalletOperateGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, iwogb.build.ctx, "GroupBy")
	if err := iwogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IncomeWalletOperateQuery, *IncomeWalletOperateGroupBy](ctx, iwogb.build, iwogb, iwogb.build.inters, v)
}

func (iwogb *IncomeWalletOperateGroupBy) sqlScan(ctx context.Context, root *IncomeWalletOperateQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(iwogb.fns))
	for _, fn := range iwogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*iwogb.flds)+len(iwogb.fns))
		for _, f := range *iwogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*iwogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := iwogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// IncomeWalletOperateSelect is the builder for selecting fields of IncomeWalletOperate entities.
type IncomeWalletOperateSelect struct {
	*IncomeWalletOperateQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (iwos *IncomeWalletOperateSelect) Aggregate(fns ...AggregateFunc) *IncomeWalletOperateSelect {
	iwos.fns = append(iwos.fns, fns...)
	return iwos
}

// Scan applies the selector query and scans the result into the given value.
func (iwos *IncomeWalletOperateSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, iwos.ctx, "Select")
	if err := iwos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IncomeWalletOperateQuery, *IncomeWalletOperateSelect](ctx, iwos.IncomeWalletOperateQuery, iwos, iwos.inters, v)
}

func (iwos *IncomeWalletOperateSelect) sqlScan(ctx context.Context, root *IncomeWalletOperateQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(iwos.fns))
	for _, fn := range iwos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*iwos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := iwos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (iwos *IncomeWalletOperateSelect) Modify(modifiers ...func(s *sql.Selector)) *IncomeWalletOperateSelect {
	iwos.modifiers = append(iwos.modifiers, modifiers...)
	return iwos
}
