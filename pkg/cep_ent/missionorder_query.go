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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/bill"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/device"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/mission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionbatch"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/symbol"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// MissionOrderQuery is the builder for querying MissionOrder entities.
type MissionOrderQuery struct {
	config
	ctx              *QueryContext
	order            []missionorder.OrderOption
	inters           []Interceptor
	predicates       []predicate.MissionOrder
	withConsumeUser  *UserQuery
	withProduceUser  *UserQuery
	withSymbol       *SymbolQuery
	withBills        *BillQuery
	withMissionBatch *MissionBatchQuery
	withMission      *MissionQuery
	withDevice       *DeviceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MissionOrderQuery builder.
func (moq *MissionOrderQuery) Where(ps ...predicate.MissionOrder) *MissionOrderQuery {
	moq.predicates = append(moq.predicates, ps...)
	return moq
}

// Limit the number of records to be returned by this query.
func (moq *MissionOrderQuery) Limit(limit int) *MissionOrderQuery {
	moq.ctx.Limit = &limit
	return moq
}

// Offset to start from.
func (moq *MissionOrderQuery) Offset(offset int) *MissionOrderQuery {
	moq.ctx.Offset = &offset
	return moq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (moq *MissionOrderQuery) Unique(unique bool) *MissionOrderQuery {
	moq.ctx.Unique = &unique
	return moq
}

// Order specifies how the records should be ordered.
func (moq *MissionOrderQuery) Order(o ...missionorder.OrderOption) *MissionOrderQuery {
	moq.order = append(moq.order, o...)
	return moq
}

// QueryConsumeUser chains the current query on the "consume_user" edge.
func (moq *MissionOrderQuery) QueryConsumeUser() *UserQuery {
	query := (&UserClient{config: moq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := moq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := moq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionorder.Table, missionorder.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, missionorder.ConsumeUserTable, missionorder.ConsumeUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(moq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProduceUser chains the current query on the "produce_user" edge.
func (moq *MissionOrderQuery) QueryProduceUser() *UserQuery {
	query := (&UserClient{config: moq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := moq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := moq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionorder.Table, missionorder.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, missionorder.ProduceUserTable, missionorder.ProduceUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(moq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySymbol chains the current query on the "symbol" edge.
func (moq *MissionOrderQuery) QuerySymbol() *SymbolQuery {
	query := (&SymbolClient{config: moq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := moq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := moq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionorder.Table, missionorder.FieldID, selector),
			sqlgraph.To(symbol.Table, symbol.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, missionorder.SymbolTable, missionorder.SymbolColumn),
		)
		fromU = sqlgraph.SetNeighbors(moq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBills chains the current query on the "bills" edge.
func (moq *MissionOrderQuery) QueryBills() *BillQuery {
	query := (&BillClient{config: moq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := moq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := moq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionorder.Table, missionorder.FieldID, selector),
			sqlgraph.To(bill.Table, bill.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, missionorder.BillsTable, missionorder.BillsColumn),
		)
		fromU = sqlgraph.SetNeighbors(moq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMissionBatch chains the current query on the "mission_batch" edge.
func (moq *MissionOrderQuery) QueryMissionBatch() *MissionBatchQuery {
	query := (&MissionBatchClient{config: moq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := moq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := moq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionorder.Table, missionorder.FieldID, selector),
			sqlgraph.To(missionbatch.Table, missionbatch.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, missionorder.MissionBatchTable, missionorder.MissionBatchColumn),
		)
		fromU = sqlgraph.SetNeighbors(moq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMission chains the current query on the "mission" edge.
func (moq *MissionOrderQuery) QueryMission() *MissionQuery {
	query := (&MissionClient{config: moq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := moq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := moq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionorder.Table, missionorder.FieldID, selector),
			sqlgraph.To(mission.Table, mission.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, missionorder.MissionTable, missionorder.MissionColumn),
		)
		fromU = sqlgraph.SetNeighbors(moq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDevice chains the current query on the "device" edge.
func (moq *MissionOrderQuery) QueryDevice() *DeviceQuery {
	query := (&DeviceClient{config: moq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := moq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := moq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(missionorder.Table, missionorder.FieldID, selector),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, missionorder.DeviceTable, missionorder.DeviceColumn),
		)
		fromU = sqlgraph.SetNeighbors(moq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MissionOrder entity from the query.
// Returns a *NotFoundError when no MissionOrder was found.
func (moq *MissionOrderQuery) First(ctx context.Context) (*MissionOrder, error) {
	nodes, err := moq.Limit(1).All(setContextOp(ctx, moq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{missionorder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (moq *MissionOrderQuery) FirstX(ctx context.Context) *MissionOrder {
	node, err := moq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MissionOrder ID from the query.
// Returns a *NotFoundError when no MissionOrder ID was found.
func (moq *MissionOrderQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = moq.Limit(1).IDs(setContextOp(ctx, moq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{missionorder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (moq *MissionOrderQuery) FirstIDX(ctx context.Context) int64 {
	id, err := moq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MissionOrder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MissionOrder entity is found.
// Returns a *NotFoundError when no MissionOrder entities are found.
func (moq *MissionOrderQuery) Only(ctx context.Context) (*MissionOrder, error) {
	nodes, err := moq.Limit(2).All(setContextOp(ctx, moq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{missionorder.Label}
	default:
		return nil, &NotSingularError{missionorder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (moq *MissionOrderQuery) OnlyX(ctx context.Context) *MissionOrder {
	node, err := moq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MissionOrder ID in the query.
// Returns a *NotSingularError when more than one MissionOrder ID is found.
// Returns a *NotFoundError when no entities are found.
func (moq *MissionOrderQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = moq.Limit(2).IDs(setContextOp(ctx, moq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{missionorder.Label}
	default:
		err = &NotSingularError{missionorder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (moq *MissionOrderQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := moq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MissionOrders.
func (moq *MissionOrderQuery) All(ctx context.Context) ([]*MissionOrder, error) {
	ctx = setContextOp(ctx, moq.ctx, "All")
	if err := moq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MissionOrder, *MissionOrderQuery]()
	return withInterceptors[[]*MissionOrder](ctx, moq, qr, moq.inters)
}

// AllX is like All, but panics if an error occurs.
func (moq *MissionOrderQuery) AllX(ctx context.Context) []*MissionOrder {
	nodes, err := moq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MissionOrder IDs.
func (moq *MissionOrderQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if moq.ctx.Unique == nil && moq.path != nil {
		moq.Unique(true)
	}
	ctx = setContextOp(ctx, moq.ctx, "IDs")
	if err = moq.Select(missionorder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (moq *MissionOrderQuery) IDsX(ctx context.Context) []int64 {
	ids, err := moq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (moq *MissionOrderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, moq.ctx, "Count")
	if err := moq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, moq, querierCount[*MissionOrderQuery](), moq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (moq *MissionOrderQuery) CountX(ctx context.Context) int {
	count, err := moq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (moq *MissionOrderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, moq.ctx, "Exist")
	switch _, err := moq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cep_ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (moq *MissionOrderQuery) ExistX(ctx context.Context) bool {
	exist, err := moq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MissionOrderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (moq *MissionOrderQuery) Clone() *MissionOrderQuery {
	if moq == nil {
		return nil
	}
	return &MissionOrderQuery{
		config:           moq.config,
		ctx:              moq.ctx.Clone(),
		order:            append([]missionorder.OrderOption{}, moq.order...),
		inters:           append([]Interceptor{}, moq.inters...),
		predicates:       append([]predicate.MissionOrder{}, moq.predicates...),
		withConsumeUser:  moq.withConsumeUser.Clone(),
		withProduceUser:  moq.withProduceUser.Clone(),
		withSymbol:       moq.withSymbol.Clone(),
		withBills:        moq.withBills.Clone(),
		withMissionBatch: moq.withMissionBatch.Clone(),
		withMission:      moq.withMission.Clone(),
		withDevice:       moq.withDevice.Clone(),
		// clone intermediate query.
		sql:  moq.sql.Clone(),
		path: moq.path,
	}
}

// WithConsumeUser tells the query-builder to eager-load the nodes that are connected to
// the "consume_user" edge. The optional arguments are used to configure the query builder of the edge.
func (moq *MissionOrderQuery) WithConsumeUser(opts ...func(*UserQuery)) *MissionOrderQuery {
	query := (&UserClient{config: moq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	moq.withConsumeUser = query
	return moq
}

// WithProduceUser tells the query-builder to eager-load the nodes that are connected to
// the "produce_user" edge. The optional arguments are used to configure the query builder of the edge.
func (moq *MissionOrderQuery) WithProduceUser(opts ...func(*UserQuery)) *MissionOrderQuery {
	query := (&UserClient{config: moq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	moq.withProduceUser = query
	return moq
}

// WithSymbol tells the query-builder to eager-load the nodes that are connected to
// the "symbol" edge. The optional arguments are used to configure the query builder of the edge.
func (moq *MissionOrderQuery) WithSymbol(opts ...func(*SymbolQuery)) *MissionOrderQuery {
	query := (&SymbolClient{config: moq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	moq.withSymbol = query
	return moq
}

// WithBills tells the query-builder to eager-load the nodes that are connected to
// the "bills" edge. The optional arguments are used to configure the query builder of the edge.
func (moq *MissionOrderQuery) WithBills(opts ...func(*BillQuery)) *MissionOrderQuery {
	query := (&BillClient{config: moq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	moq.withBills = query
	return moq
}

// WithMissionBatch tells the query-builder to eager-load the nodes that are connected to
// the "mission_batch" edge. The optional arguments are used to configure the query builder of the edge.
func (moq *MissionOrderQuery) WithMissionBatch(opts ...func(*MissionBatchQuery)) *MissionOrderQuery {
	query := (&MissionBatchClient{config: moq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	moq.withMissionBatch = query
	return moq
}

// WithMission tells the query-builder to eager-load the nodes that are connected to
// the "mission" edge. The optional arguments are used to configure the query builder of the edge.
func (moq *MissionOrderQuery) WithMission(opts ...func(*MissionQuery)) *MissionOrderQuery {
	query := (&MissionClient{config: moq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	moq.withMission = query
	return moq
}

// WithDevice tells the query-builder to eager-load the nodes that are connected to
// the "device" edge. The optional arguments are used to configure the query builder of the edge.
func (moq *MissionOrderQuery) WithDevice(opts ...func(*DeviceQuery)) *MissionOrderQuery {
	query := (&DeviceClient{config: moq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	moq.withDevice = query
	return moq
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
//	client.MissionOrder.Query().
//		GroupBy(missionorder.FieldCreatedBy).
//		Aggregate(cep_ent.Count()).
//		Scan(ctx, &v)
func (moq *MissionOrderQuery) GroupBy(field string, fields ...string) *MissionOrderGroupBy {
	moq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MissionOrderGroupBy{build: moq}
	grbuild.flds = &moq.ctx.Fields
	grbuild.label = missionorder.Label
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
//	client.MissionOrder.Query().
//		Select(missionorder.FieldCreatedBy).
//		Scan(ctx, &v)
func (moq *MissionOrderQuery) Select(fields ...string) *MissionOrderSelect {
	moq.ctx.Fields = append(moq.ctx.Fields, fields...)
	sbuild := &MissionOrderSelect{MissionOrderQuery: moq}
	sbuild.label = missionorder.Label
	sbuild.flds, sbuild.scan = &moq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MissionOrderSelect configured with the given aggregations.
func (moq *MissionOrderQuery) Aggregate(fns ...AggregateFunc) *MissionOrderSelect {
	return moq.Select().Aggregate(fns...)
}

func (moq *MissionOrderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range moq.inters {
		if inter == nil {
			return fmt.Errorf("cep_ent: uninitialized interceptor (forgotten import cep_ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, moq); err != nil {
				return err
			}
		}
	}
	for _, f := range moq.ctx.Fields {
		if !missionorder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
		}
	}
	if moq.path != nil {
		prev, err := moq.path(ctx)
		if err != nil {
			return err
		}
		moq.sql = prev
	}
	return nil
}

func (moq *MissionOrderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MissionOrder, error) {
	var (
		nodes       = []*MissionOrder{}
		_spec       = moq.querySpec()
		loadedTypes = [7]bool{
			moq.withConsumeUser != nil,
			moq.withProduceUser != nil,
			moq.withSymbol != nil,
			moq.withBills != nil,
			moq.withMissionBatch != nil,
			moq.withMission != nil,
			moq.withDevice != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MissionOrder).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MissionOrder{config: moq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, moq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := moq.withConsumeUser; query != nil {
		if err := moq.loadConsumeUser(ctx, query, nodes, nil,
			func(n *MissionOrder, e *User) { n.Edges.ConsumeUser = e }); err != nil {
			return nil, err
		}
	}
	if query := moq.withProduceUser; query != nil {
		if err := moq.loadProduceUser(ctx, query, nodes, nil,
			func(n *MissionOrder, e *User) { n.Edges.ProduceUser = e }); err != nil {
			return nil, err
		}
	}
	if query := moq.withSymbol; query != nil {
		if err := moq.loadSymbol(ctx, query, nodes, nil,
			func(n *MissionOrder, e *Symbol) { n.Edges.Symbol = e }); err != nil {
			return nil, err
		}
	}
	if query := moq.withBills; query != nil {
		if err := moq.loadBills(ctx, query, nodes,
			func(n *MissionOrder) { n.Edges.Bills = []*Bill{} },
			func(n *MissionOrder, e *Bill) { n.Edges.Bills = append(n.Edges.Bills, e) }); err != nil {
			return nil, err
		}
	}
	if query := moq.withMissionBatch; query != nil {
		if err := moq.loadMissionBatch(ctx, query, nodes, nil,
			func(n *MissionOrder, e *MissionBatch) { n.Edges.MissionBatch = e }); err != nil {
			return nil, err
		}
	}
	if query := moq.withMission; query != nil {
		if err := moq.loadMission(ctx, query, nodes, nil,
			func(n *MissionOrder, e *Mission) { n.Edges.Mission = e }); err != nil {
			return nil, err
		}
	}
	if query := moq.withDevice; query != nil {
		if err := moq.loadDevice(ctx, query, nodes, nil,
			func(n *MissionOrder, e *Device) { n.Edges.Device = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (moq *MissionOrderQuery) loadConsumeUser(ctx context.Context, query *UserQuery, nodes []*MissionOrder, init func(*MissionOrder), assign func(*MissionOrder, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MissionOrder)
	for i := range nodes {
		fk := nodes[i].ConsumeUserID
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
			return fmt.Errorf(`unexpected foreign-key "consume_user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (moq *MissionOrderQuery) loadProduceUser(ctx context.Context, query *UserQuery, nodes []*MissionOrder, init func(*MissionOrder), assign func(*MissionOrder, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MissionOrder)
	for i := range nodes {
		fk := nodes[i].ProduceUserID
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
			return fmt.Errorf(`unexpected foreign-key "produce_user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (moq *MissionOrderQuery) loadSymbol(ctx context.Context, query *SymbolQuery, nodes []*MissionOrder, init func(*MissionOrder), assign func(*MissionOrder, *Symbol)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MissionOrder)
	for i := range nodes {
		fk := nodes[i].SymbolID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(symbol.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "symbol_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (moq *MissionOrderQuery) loadBills(ctx context.Context, query *BillQuery, nodes []*MissionOrder, init func(*MissionOrder), assign func(*MissionOrder, *Bill)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*MissionOrder)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(bill.FieldOrderID)
	}
	query.Where(predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(missionorder.BillsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OrderID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "order_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (moq *MissionOrderQuery) loadMissionBatch(ctx context.Context, query *MissionBatchQuery, nodes []*MissionOrder, init func(*MissionOrder), assign func(*MissionOrder, *MissionBatch)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MissionOrder)
	for i := range nodes {
		fk := nodes[i].MissionBatchID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(missionbatch.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "mission_batch_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (moq *MissionOrderQuery) loadMission(ctx context.Context, query *MissionQuery, nodes []*MissionOrder, init func(*MissionOrder), assign func(*MissionOrder, *Mission)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MissionOrder)
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
func (moq *MissionOrderQuery) loadDevice(ctx context.Context, query *DeviceQuery, nodes []*MissionOrder, init func(*MissionOrder), assign func(*MissionOrder, *Device)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MissionOrder)
	for i := range nodes {
		fk := nodes[i].DeviceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(device.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "device_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (moq *MissionOrderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := moq.querySpec()
	_spec.Node.Columns = moq.ctx.Fields
	if len(moq.ctx.Fields) > 0 {
		_spec.Unique = moq.ctx.Unique != nil && *moq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, moq.driver, _spec)
}

func (moq *MissionOrderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(missionorder.Table, missionorder.Columns, sqlgraph.NewFieldSpec(missionorder.FieldID, field.TypeInt64))
	_spec.From = moq.sql
	if unique := moq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if moq.path != nil {
		_spec.Unique = true
	}
	if fields := moq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, missionorder.FieldID)
		for i := range fields {
			if fields[i] != missionorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if moq.withConsumeUser != nil {
			_spec.Node.AddColumnOnce(missionorder.FieldConsumeUserID)
		}
		if moq.withProduceUser != nil {
			_spec.Node.AddColumnOnce(missionorder.FieldProduceUserID)
		}
		if moq.withSymbol != nil {
			_spec.Node.AddColumnOnce(missionorder.FieldSymbolID)
		}
		if moq.withMissionBatch != nil {
			_spec.Node.AddColumnOnce(missionorder.FieldMissionBatchID)
		}
		if moq.withMission != nil {
			_spec.Node.AddColumnOnce(missionorder.FieldMissionID)
		}
		if moq.withDevice != nil {
			_spec.Node.AddColumnOnce(missionorder.FieldDeviceID)
		}
	}
	if ps := moq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := moq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := moq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := moq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (moq *MissionOrderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(moq.driver.Dialect())
	t1 := builder.Table(missionorder.Table)
	columns := moq.ctx.Fields
	if len(columns) == 0 {
		columns = missionorder.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if moq.sql != nil {
		selector = moq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if moq.ctx.Unique != nil && *moq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range moq.predicates {
		p(selector)
	}
	for _, p := range moq.order {
		p(selector)
	}
	if offset := moq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := moq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MissionOrderGroupBy is the group-by builder for MissionOrder entities.
type MissionOrderGroupBy struct {
	selector
	build *MissionOrderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mogb *MissionOrderGroupBy) Aggregate(fns ...AggregateFunc) *MissionOrderGroupBy {
	mogb.fns = append(mogb.fns, fns...)
	return mogb
}

// Scan applies the selector query and scans the result into the given value.
func (mogb *MissionOrderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mogb.build.ctx, "GroupBy")
	if err := mogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MissionOrderQuery, *MissionOrderGroupBy](ctx, mogb.build, mogb, mogb.build.inters, v)
}

func (mogb *MissionOrderGroupBy) sqlScan(ctx context.Context, root *MissionOrderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mogb.fns))
	for _, fn := range mogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mogb.flds)+len(mogb.fns))
		for _, f := range *mogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MissionOrderSelect is the builder for selecting fields of MissionOrder entities.
type MissionOrderSelect struct {
	*MissionOrderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mos *MissionOrderSelect) Aggregate(fns ...AggregateFunc) *MissionOrderSelect {
	mos.fns = append(mos.fns, fns...)
	return mos
}

// Scan applies the selector query and scans the result into the given value.
func (mos *MissionOrderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mos.ctx, "Select")
	if err := mos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MissionOrderQuery, *MissionOrderSelect](ctx, mos.MissionOrderQuery, mos, mos.inters, v)
}

func (mos *MissionOrderSelect) sqlScan(ctx context.Context, root *MissionOrderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mos.fns))
	for _, fn := range mos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
