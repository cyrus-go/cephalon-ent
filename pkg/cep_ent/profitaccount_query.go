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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/earnbill"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/profitaccount"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// ProfitAccountQuery is the builder for querying ProfitAccount entities.
type ProfitAccountQuery struct {
	config
	ctx           *QueryContext
	order         []profitaccount.OrderOption
	inters        []Interceptor
	predicates    []predicate.ProfitAccount
	withUser      *UserQuery
	withEarnBills *EarnBillQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProfitAccountQuery builder.
func (paq *ProfitAccountQuery) Where(ps ...predicate.ProfitAccount) *ProfitAccountQuery {
	paq.predicates = append(paq.predicates, ps...)
	return paq
}

// Limit the number of records to be returned by this query.
func (paq *ProfitAccountQuery) Limit(limit int) *ProfitAccountQuery {
	paq.ctx.Limit = &limit
	return paq
}

// Offset to start from.
func (paq *ProfitAccountQuery) Offset(offset int) *ProfitAccountQuery {
	paq.ctx.Offset = &offset
	return paq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (paq *ProfitAccountQuery) Unique(unique bool) *ProfitAccountQuery {
	paq.ctx.Unique = &unique
	return paq
}

// Order specifies how the records should be ordered.
func (paq *ProfitAccountQuery) Order(o ...profitaccount.OrderOption) *ProfitAccountQuery {
	paq.order = append(paq.order, o...)
	return paq
}

// QueryUser chains the current query on the "user" edge.
func (paq *ProfitAccountQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: paq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := paq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profitaccount.Table, profitaccount.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, profitaccount.UserTable, profitaccount.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(paq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEarnBills chains the current query on the "earn_bills" edge.
func (paq *ProfitAccountQuery) QueryEarnBills() *EarnBillQuery {
	query := (&EarnBillClient{config: paq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := paq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profitaccount.Table, profitaccount.FieldID, selector),
			sqlgraph.To(earnbill.Table, earnbill.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, profitaccount.EarnBillsTable, profitaccount.EarnBillsColumn),
		)
		fromU = sqlgraph.SetNeighbors(paq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProfitAccount entity from the query.
// Returns a *NotFoundError when no ProfitAccount was found.
func (paq *ProfitAccountQuery) First(ctx context.Context) (*ProfitAccount, error) {
	nodes, err := paq.Limit(1).All(setContextOp(ctx, paq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{profitaccount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (paq *ProfitAccountQuery) FirstX(ctx context.Context) *ProfitAccount {
	node, err := paq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProfitAccount ID from the query.
// Returns a *NotFoundError when no ProfitAccount ID was found.
func (paq *ProfitAccountQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = paq.Limit(1).IDs(setContextOp(ctx, paq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{profitaccount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (paq *ProfitAccountQuery) FirstIDX(ctx context.Context) int64 {
	id, err := paq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProfitAccount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProfitAccount entity is found.
// Returns a *NotFoundError when no ProfitAccount entities are found.
func (paq *ProfitAccountQuery) Only(ctx context.Context) (*ProfitAccount, error) {
	nodes, err := paq.Limit(2).All(setContextOp(ctx, paq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{profitaccount.Label}
	default:
		return nil, &NotSingularError{profitaccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (paq *ProfitAccountQuery) OnlyX(ctx context.Context) *ProfitAccount {
	node, err := paq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProfitAccount ID in the query.
// Returns a *NotSingularError when more than one ProfitAccount ID is found.
// Returns a *NotFoundError when no entities are found.
func (paq *ProfitAccountQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = paq.Limit(2).IDs(setContextOp(ctx, paq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{profitaccount.Label}
	default:
		err = &NotSingularError{profitaccount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (paq *ProfitAccountQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := paq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProfitAccounts.
func (paq *ProfitAccountQuery) All(ctx context.Context) ([]*ProfitAccount, error) {
	ctx = setContextOp(ctx, paq.ctx, "All")
	if err := paq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProfitAccount, *ProfitAccountQuery]()
	return withInterceptors[[]*ProfitAccount](ctx, paq, qr, paq.inters)
}

// AllX is like All, but panics if an error occurs.
func (paq *ProfitAccountQuery) AllX(ctx context.Context) []*ProfitAccount {
	nodes, err := paq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProfitAccount IDs.
func (paq *ProfitAccountQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if paq.ctx.Unique == nil && paq.path != nil {
		paq.Unique(true)
	}
	ctx = setContextOp(ctx, paq.ctx, "IDs")
	if err = paq.Select(profitaccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (paq *ProfitAccountQuery) IDsX(ctx context.Context) []int64 {
	ids, err := paq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (paq *ProfitAccountQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, paq.ctx, "Count")
	if err := paq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, paq, querierCount[*ProfitAccountQuery](), paq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (paq *ProfitAccountQuery) CountX(ctx context.Context) int {
	count, err := paq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (paq *ProfitAccountQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, paq.ctx, "Exist")
	switch _, err := paq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (paq *ProfitAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := paq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProfitAccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (paq *ProfitAccountQuery) Clone() *ProfitAccountQuery {
	if paq == nil {
		return nil
	}
	return &ProfitAccountQuery{
		config:        paq.config,
		ctx:           paq.ctx.Clone(),
		order:         append([]profitaccount.OrderOption{}, paq.order...),
		inters:        append([]Interceptor{}, paq.inters...),
		predicates:    append([]predicate.ProfitAccount{}, paq.predicates...),
		withUser:      paq.withUser.Clone(),
		withEarnBills: paq.withEarnBills.Clone(),
		// clone intermediate query.
		sql:  paq.sql.Clone(),
		path: paq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (paq *ProfitAccountQuery) WithUser(opts ...func(*UserQuery)) *ProfitAccountQuery {
	query := (&UserClient{config: paq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	paq.withUser = query
	return paq
}

// WithEarnBills tells the query-builder to eager-load the nodes that are connected to
// the "earn_bills" edge. The optional arguments are used to configure the query builder of the edge.
func (paq *ProfitAccountQuery) WithEarnBills(opts ...func(*EarnBillQuery)) *ProfitAccountQuery {
	query := (&EarnBillClient{config: paq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	paq.withEarnBills = query
	return paq
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
//	client.ProfitAccount.Query().
//		GroupBy(profitaccount.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (paq *ProfitAccountQuery) GroupBy(field string, fields ...string) *ProfitAccountGroupBy {
	paq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProfitAccountGroupBy{build: paq}
	grbuild.flds = &paq.ctx.Fields
	grbuild.label = profitaccount.Label
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
//	client.ProfitAccount.Query().
//		Select(profitaccount.FieldCreatedBy).
//		Scan(ctx, &v)
func (paq *ProfitAccountQuery) Select(fields ...string) *ProfitAccountSelect {
	paq.ctx.Fields = append(paq.ctx.Fields, fields...)
	sbuild := &ProfitAccountSelect{ProfitAccountQuery: paq}
	sbuild.label = profitaccount.Label
	sbuild.flds, sbuild.scan = &paq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProfitAccountSelect configured with the given aggregations.
func (paq *ProfitAccountQuery) Aggregate(fns ...AggregateFunc) *ProfitAccountSelect {
	return paq.Select().Aggregate(fns...)
}

func (paq *ProfitAccountQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range paq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, paq); err != nil {
				return err
			}
		}
	}
	for _, f := range paq.ctx.Fields {
		if !profitaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if paq.path != nil {
		prev, err := paq.path(ctx)
		if err != nil {
			return err
		}
		paq.sql = prev
	}
	return nil
}

func (paq *ProfitAccountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProfitAccount, error) {
	var (
		nodes       = []*ProfitAccount{}
		_spec       = paq.querySpec()
		loadedTypes = [2]bool{
			paq.withUser != nil,
			paq.withEarnBills != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProfitAccount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProfitAccount{config: paq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, paq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := paq.withUser; query != nil {
		if err := paq.loadUser(ctx, query, nodes, nil,
			func(n *ProfitAccount, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := paq.withEarnBills; query != nil {
		if err := paq.loadEarnBills(ctx, query, nodes,
			func(n *ProfitAccount) { n.Edges.EarnBills = []*EarnBill{} },
			func(n *ProfitAccount, e *EarnBill) { n.Edges.EarnBills = append(n.Edges.EarnBills, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (paq *ProfitAccountQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*ProfitAccount, init func(*ProfitAccount), assign func(*ProfitAccount, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*ProfitAccount)
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
func (paq *ProfitAccountQuery) loadEarnBills(ctx context.Context, query *EarnBillQuery, nodes []*ProfitAccount, init func(*ProfitAccount), assign func(*ProfitAccount, *EarnBill)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*ProfitAccount)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(earnbill.FieldProfitAccountID)
	}
	query.Where(predicate.EarnBill(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(profitaccount.EarnBillsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProfitAccountID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "profit_account_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (paq *ProfitAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := paq.querySpec()
	_spec.Node.Columns = paq.ctx.Fields
	if len(paq.ctx.Fields) > 0 {
		_spec.Unique = paq.ctx.Unique != nil && *paq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, paq.driver, _spec)
}

func (paq *ProfitAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(profitaccount.Table, profitaccount.Columns, sqlgraph.NewFieldSpec(profitaccount.FieldID, field.TypeInt64))
	_spec.From = paq.sql
	if unique := paq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if paq.path != nil {
		_spec.Unique = true
	}
	if fields := paq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profitaccount.FieldID)
		for i := range fields {
			if fields[i] != profitaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if paq.withUser != nil {
			_spec.Node.AddColumnOnce(profitaccount.FieldUserID)
		}
	}
	if ps := paq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := paq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := paq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := paq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (paq *ProfitAccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(paq.driver.Dialect())
	t1 := builder.Table(profitaccount.Table)
	columns := paq.ctx.Fields
	if len(columns) == 0 {
		columns = profitaccount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if paq.sql != nil {
		selector = paq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if paq.ctx.Unique != nil && *paq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range paq.predicates {
		p(selector)
	}
	for _, p := range paq.order {
		p(selector)
	}
	if offset := paq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := paq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProfitAccountGroupBy is the group-by builder for ProfitAccount entities.
type ProfitAccountGroupBy struct {
	selector
	build *ProfitAccountQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pagb *ProfitAccountGroupBy) Aggregate(fns ...AggregateFunc) *ProfitAccountGroupBy {
	pagb.fns = append(pagb.fns, fns...)
	return pagb
}

// Scan applies the selector query and scans the result into the given value.
func (pagb *ProfitAccountGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pagb.build.ctx, "GroupBy")
	if err := pagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProfitAccountQuery, *ProfitAccountGroupBy](ctx, pagb.build, pagb, pagb.build.inters, v)
}

func (pagb *ProfitAccountGroupBy) sqlScan(ctx context.Context, root *ProfitAccountQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pagb.fns))
	for _, fn := range pagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pagb.flds)+len(pagb.fns))
		for _, f := range *pagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProfitAccountSelect is the builder for selecting fields of ProfitAccount entities.
type ProfitAccountSelect struct {
	*ProfitAccountQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pas *ProfitAccountSelect) Aggregate(fns ...AggregateFunc) *ProfitAccountSelect {
	pas.fns = append(pas.fns, fns...)
	return pas
}

// Scan applies the selector query and scans the result into the given value.
func (pas *ProfitAccountSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pas.ctx, "Select")
	if err := pas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProfitAccountQuery, *ProfitAccountSelect](ctx, pas.ProfitAccountQuery, pas, pas.inters, v)
}

func (pas *ProfitAccountSelect) sqlScan(ctx context.Context, root *ProfitAccountQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pas.fns))
	for _, fn := range pas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}