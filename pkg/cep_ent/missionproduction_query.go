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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/mission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionproduceorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionproduction"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// MissionProductionQuery is the builder for querying MissionProduction entities.
type MissionProductionQuery struct {
	config
	ctx                     *QueryContext
	order                   []missionproduction.OrderOption
	inters                  []Interceptor
	predicates              []predicate.MissionProduction
	withMission             *MissionQuery
	withUser                *UserQuery
	withMissionProduceOrder *MissionProduceOrderQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MissionProductionQuery builder.
func (mpq *MissionProductionQuery) Where(ps ...predicate.MissionProduction) *MissionProductionQuery {
	mpq.predicates = append(mpq.predicates, ps...)
	return mpq
}

// Limit the number of records to be returned by this query.
func (mpq *MissionProductionQuery) Limit(limit int) *MissionProductionQuery {
	mpq.ctx.Limit = &limit
	return mpq
}

// Offset to start from.
func (mpq *MissionProductionQuery) Offset(offset int) *MissionProductionQuery {
	mpq.ctx.Offset = &offset
	return mpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mpq *MissionProductionQuery) Unique(unique bool) *MissionProductionQuery {
	mpq.ctx.Unique = &unique
	return mpq
}

// Order specifies how the records should be ordered.
func (mpq *MissionProductionQuery) Order(o ...missionproduction.OrderOption) *MissionProductionQuery {
	mpq.order = append(mpq.order, o...)
	return mpq
}

// QueryMission chains the current query on the "mission" edge.
func (mpq *MissionProductionQuery) QueryMission() *MissionQuery {
	query := (&MissionClient{config: mpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionproduction.Table, missionproduction.FieldID, selector),
			sqlgraph.To(mission.Table, mission.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, missionproduction.MissionTable, missionproduction.MissionColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (mpq *MissionProductionQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: mpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionproduction.Table, missionproduction.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, missionproduction.UserTable, missionproduction.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMissionProduceOrder chains the current query on the "mission_produce_order" edge.
func (mpq *MissionProductionQuery) QueryMissionProduceOrder() *MissionProduceOrderQuery {
	query := (&MissionProduceOrderClient{config: mpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionproduction.Table, missionproduction.FieldID, selector),
			sqlgraph.To(missionproduceorder.Table, missionproduceorder.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, missionproduction.MissionProduceOrderTable, missionproduction.MissionProduceOrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MissionProduction entity from the query.
// Returns a *NotFoundError when no MissionProduction was found.
func (mpq *MissionProductionQuery) First(ctx context.Context) (*MissionProduction, error) {
	nodes, err := mpq.Limit(1).All(setContextOp(ctx, mpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{missionproduction.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mpq *MissionProductionQuery) FirstX(ctx context.Context) *MissionProduction {
	node, err := mpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MissionProduction ID from the query.
// Returns a *NotFoundError when no MissionProduction ID was found.
func (mpq *MissionProductionQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mpq.Limit(1).IDs(setContextOp(ctx, mpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{missionproduction.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mpq *MissionProductionQuery) FirstIDX(ctx context.Context) int64 {
	id, err := mpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MissionProduction entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MissionProduction entity is found.
// Returns a *NotFoundError when no MissionProduction entities are found.
func (mpq *MissionProductionQuery) Only(ctx context.Context) (*MissionProduction, error) {
	nodes, err := mpq.Limit(2).All(setContextOp(ctx, mpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{missionproduction.Label}
	default:
		return nil, &NotSingularError{missionproduction.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mpq *MissionProductionQuery) OnlyX(ctx context.Context) *MissionProduction {
	node, err := mpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MissionProduction ID in the query.
// Returns a *NotSingularError when more than one MissionProduction ID is found.
// Returns a *NotFoundError when no entities are found.
func (mpq *MissionProductionQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mpq.Limit(2).IDs(setContextOp(ctx, mpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{missionproduction.Label}
	default:
		err = &NotSingularError{missionproduction.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mpq *MissionProductionQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := mpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MissionProductions.
func (mpq *MissionProductionQuery) All(ctx context.Context) ([]*MissionProduction, error) {
	ctx = setContextOp(ctx, mpq.ctx, "All")
	if err := mpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MissionProduction, *MissionProductionQuery]()
	return withInterceptors[[]*MissionProduction](ctx, mpq, qr, mpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mpq *MissionProductionQuery) AllX(ctx context.Context) []*MissionProduction {
	nodes, err := mpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MissionProduction IDs.
func (mpq *MissionProductionQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if mpq.ctx.Unique == nil && mpq.path != nil {
		mpq.Unique(true)
	}
	ctx = setContextOp(ctx, mpq.ctx, "IDs")
	if err = mpq.Select(missionproduction.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mpq *MissionProductionQuery) IDsX(ctx context.Context) []int64 {
	ids, err := mpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mpq *MissionProductionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mpq.ctx, "Count")
	if err := mpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mpq, querierCount[*MissionProductionQuery](), mpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mpq *MissionProductionQuery) CountX(ctx context.Context) int {
	count, err := mpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mpq *MissionProductionQuery) Exist(ctx context.Context) (bool, error) {
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
func (mpq *MissionProductionQuery) ExistX(ctx context.Context) bool {
	exist, err := mpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MissionProductionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mpq *MissionProductionQuery) Clone() *MissionProductionQuery {
	if mpq == nil {
		return nil
	}
	return &MissionProductionQuery{
		config:                  mpq.config,
		ctx:                     mpq.ctx.Clone(),
		order:                   append([]missionproduction.OrderOption{}, mpq.order...),
		inters:                  append([]Interceptor{}, mpq.inters...),
		predicates:              append([]predicate.MissionProduction{}, mpq.predicates...),
		withMission:             mpq.withMission.Clone(),
		withUser:                mpq.withUser.Clone(),
		withMissionProduceOrder: mpq.withMissionProduceOrder.Clone(),
		// clone intermediate query.
		sql:  mpq.sql.Clone(),
		path: mpq.path,
	}
}

// WithMission tells the query-builder to eager-load the nodes that are connected to
// the "mission" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MissionProductionQuery) WithMission(opts ...func(*MissionQuery)) *MissionProductionQuery {
	query := (&MissionClient{config: mpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mpq.withMission = query
	return mpq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MissionProductionQuery) WithUser(opts ...func(*UserQuery)) *MissionProductionQuery {
	query := (&UserClient{config: mpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mpq.withUser = query
	return mpq
}

// WithMissionProduceOrder tells the query-builder to eager-load the nodes that are connected to
// the "mission_produce_order" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MissionProductionQuery) WithMissionProduceOrder(opts ...func(*MissionProduceOrderQuery)) *MissionProductionQuery {
	query := (&MissionProduceOrderClient{config: mpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mpq.withMissionProduceOrder = query
	return mpq
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
//	client.MissionProduction.Query().
//		GroupBy(missionproduction.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (mpq *MissionProductionQuery) GroupBy(field string, fields ...string) *MissionProductionGroupBy {
	mpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MissionProductionGroupBy{build: mpq}
	grbuild.flds = &mpq.ctx.Fields
	grbuild.label = missionproduction.Label
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
//	client.MissionProduction.Query().
//		Select(missionproduction.FieldCreatedBy).
//		Scan(ctx, &v)
func (mpq *MissionProductionQuery) Select(fields ...string) *MissionProductionSelect {
	mpq.ctx.Fields = append(mpq.ctx.Fields, fields...)
	sbuild := &MissionProductionSelect{MissionProductionQuery: mpq}
	sbuild.label = missionproduction.Label
	sbuild.flds, sbuild.scan = &mpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MissionProductionSelect configured with the given aggregations.
func (mpq *MissionProductionQuery) Aggregate(fns ...AggregateFunc) *MissionProductionSelect {
	return mpq.Select().Aggregate(fns...)
}

func (mpq *MissionProductionQuery) prepareQuery(ctx context.Context) error {
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
		if !missionproduction.ValidColumn(f) {
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

func (mpq *MissionProductionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MissionProduction, error) {
	var (
		nodes       = []*MissionProduction{}
		_spec       = mpq.querySpec()
		loadedTypes = [3]bool{
			mpq.withMission != nil,
			mpq.withUser != nil,
			mpq.withMissionProduceOrder != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MissionProduction).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MissionProduction{config: mpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
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
	if query := mpq.withMission; query != nil {
		if err := mpq.loadMission(ctx, query, nodes, nil,
			func(n *MissionProduction, e *Mission) { n.Edges.Mission = e }); err != nil {
			return nil, err
		}
	}
	if query := mpq.withUser; query != nil {
		if err := mpq.loadUser(ctx, query, nodes, nil,
			func(n *MissionProduction, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := mpq.withMissionProduceOrder; query != nil {
		if err := mpq.loadMissionProduceOrder(ctx, query, nodes, nil,
			func(n *MissionProduction, e *MissionProduceOrder) { n.Edges.MissionProduceOrder = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mpq *MissionProductionQuery) loadMission(ctx context.Context, query *MissionQuery, nodes []*MissionProduction, init func(*MissionProduction), assign func(*MissionProduction, *Mission)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MissionProduction)
	for i := range nodes {
		fk := nodes[i].MissionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(mission.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "mission_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mpq *MissionProductionQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*MissionProduction, init func(*MissionProduction), assign func(*MissionProduction, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MissionProduction)
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
func (mpq *MissionProductionQuery) loadMissionProduceOrder(ctx context.Context, query *MissionProduceOrderQuery, nodes []*MissionProduction, init func(*MissionProduction), assign func(*MissionProduction, *MissionProduceOrder)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*MissionProduction)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(missionproduceorder.FieldMissionProductionID)
	}
	query.Where(predicate.MissionProduceOrder(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(missionproduction.MissionProduceOrderColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MissionProductionID
		if fk == nil {
			return fmt.Errorf(`foreign-key "mission_production_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "mission_production_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mpq *MissionProductionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mpq.querySpec()
	_spec.Node.Columns = mpq.ctx.Fields
	if len(mpq.ctx.Fields) > 0 {
		_spec.Unique = mpq.ctx.Unique != nil && *mpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mpq.driver, _spec)
}

func (mpq *MissionProductionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(missionproduction.Table, missionproduction.Columns, sqlgraph.NewFieldSpec(missionproduction.FieldID, field.TypeInt64))
	_spec.From = mpq.sql
	if unique := mpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mpq.path != nil {
		_spec.Unique = true
	}
	if fields := mpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, missionproduction.FieldID)
		for i := range fields {
			if fields[i] != missionproduction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mpq.withMission != nil {
			_spec.Node.AddColumnOnce(missionproduction.FieldMissionID)
		}
		if mpq.withUser != nil {
			_spec.Node.AddColumnOnce(missionproduction.FieldUserID)
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

func (mpq *MissionProductionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mpq.driver.Dialect())
	t1 := builder.Table(missionproduction.Table)
	columns := mpq.ctx.Fields
	if len(columns) == 0 {
		columns = missionproduction.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mpq.sql != nil {
		selector = mpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mpq.ctx.Unique != nil && *mpq.ctx.Unique {
		selector.Distinct()
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

// MissionProductionGroupBy is the group-by builder for MissionProduction entities.
type MissionProductionGroupBy struct {
	selector
	build *MissionProductionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mpgb *MissionProductionGroupBy) Aggregate(fns ...AggregateFunc) *MissionProductionGroupBy {
	mpgb.fns = append(mpgb.fns, fns...)
	return mpgb
}

// Scan applies the selector query and scans the result into the given value.
func (mpgb *MissionProductionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mpgb.build.ctx, "GroupBy")
	if err := mpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MissionProductionQuery, *MissionProductionGroupBy](ctx, mpgb.build, mpgb, mpgb.build.inters, v)
}

func (mpgb *MissionProductionGroupBy) sqlScan(ctx context.Context, root *MissionProductionQuery, v any) error {
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

// MissionProductionSelect is the builder for selecting fields of MissionProduction entities.
type MissionProductionSelect struct {
	*MissionProductionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mps *MissionProductionSelect) Aggregate(fns ...AggregateFunc) *MissionProductionSelect {
	mps.fns = append(mps.fns, fns...)
	return mps
}

// Scan applies the selector query and scans the result into the given value.
func (mps *MissionProductionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mps.ctx, "Select")
	if err := mps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MissionProductionQuery, *MissionProductionSelect](ctx, mps.MissionProductionQuery, mps, mps.inters, v)
}

func (mps *MissionProductionSelect) sqlScan(ctx context.Context, root *MissionProductionQuery, v any) error {
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