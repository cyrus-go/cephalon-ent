// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/extraserviceorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/mission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionbatch"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/symbol"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ExtraServiceOrderUpdate is the builder for updating ExtraServiceOrder entities.
type ExtraServiceOrderUpdate struct {
	config
	hooks     []Hook
	mutation  *ExtraServiceOrderMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ExtraServiceOrderUpdate builder.
func (esou *ExtraServiceOrderUpdate) Where(ps ...predicate.ExtraServiceOrder) *ExtraServiceOrderUpdate {
	esou.mutation.Where(ps...)
	return esou
}

// SetCreatedBy sets the "created_by" field.
func (esou *ExtraServiceOrderUpdate) SetCreatedBy(i int64) *ExtraServiceOrderUpdate {
	esou.mutation.ResetCreatedBy()
	esou.mutation.SetCreatedBy(i)
	return esou
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (esou *ExtraServiceOrderUpdate) SetNillableCreatedBy(i *int64) *ExtraServiceOrderUpdate {
	if i != nil {
		esou.SetCreatedBy(*i)
	}
	return esou
}

// AddCreatedBy adds i to the "created_by" field.
func (esou *ExtraServiceOrderUpdate) AddCreatedBy(i int64) *ExtraServiceOrderUpdate {
	esou.mutation.AddCreatedBy(i)
	return esou
}

// SetUpdatedBy sets the "updated_by" field.
func (esou *ExtraServiceOrderUpdate) SetUpdatedBy(i int64) *ExtraServiceOrderUpdate {
	esou.mutation.ResetUpdatedBy()
	esou.mutation.SetUpdatedBy(i)
	return esou
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (esou *ExtraServiceOrderUpdate) SetNillableUpdatedBy(i *int64) *ExtraServiceOrderUpdate {
	if i != nil {
		esou.SetUpdatedBy(*i)
	}
	return esou
}

// AddUpdatedBy adds i to the "updated_by" field.
func (esou *ExtraServiceOrderUpdate) AddUpdatedBy(i int64) *ExtraServiceOrderUpdate {
	esou.mutation.AddUpdatedBy(i)
	return esou
}

// SetUpdatedAt sets the "updated_at" field.
func (esou *ExtraServiceOrderUpdate) SetUpdatedAt(t time.Time) *ExtraServiceOrderUpdate {
	esou.mutation.SetUpdatedAt(t)
	return esou
}

// SetDeletedAt sets the "deleted_at" field.
func (esou *ExtraServiceOrderUpdate) SetDeletedAt(t time.Time) *ExtraServiceOrderUpdate {
	esou.mutation.SetDeletedAt(t)
	return esou
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (esou *ExtraServiceOrderUpdate) SetNillableDeletedAt(t *time.Time) *ExtraServiceOrderUpdate {
	if t != nil {
		esou.SetDeletedAt(*t)
	}
	return esou
}

// SetMissionID sets the "mission_id" field.
func (esou *ExtraServiceOrderUpdate) SetMissionID(i int64) *ExtraServiceOrderUpdate {
	esou.mutation.SetMissionID(i)
	return esou
}

// SetNillableMissionID sets the "mission_id" field if the given value is not nil.
func (esou *ExtraServiceOrderUpdate) SetNillableMissionID(i *int64) *ExtraServiceOrderUpdate {
	if i != nil {
		esou.SetMissionID(*i)
	}
	return esou
}

// SetMissionOrderID sets the "mission_order_id" field.
func (esou *ExtraServiceOrderUpdate) SetMissionOrderID(i int64) *ExtraServiceOrderUpdate {
	esou.mutation.SetMissionOrderID(i)
	return esou
}

// SetNillableMissionOrderID sets the "mission_order_id" field if the given value is not nil.
func (esou *ExtraServiceOrderUpdate) SetNillableMissionOrderID(i *int64) *ExtraServiceOrderUpdate {
	if i != nil {
		esou.SetMissionOrderID(*i)
	}
	return esou
}

// SetAmount sets the "amount" field.
func (esou *ExtraServiceOrderUpdate) SetAmount(i int64) *ExtraServiceOrderUpdate {
	esou.mutation.ResetAmount()
	esou.mutation.SetAmount(i)
	return esou
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (esou *ExtraServiceOrderUpdate) SetNillableAmount(i *int64) *ExtraServiceOrderUpdate {
	if i != nil {
		esou.SetAmount(*i)
	}
	return esou
}

// AddAmount adds i to the "amount" field.
func (esou *ExtraServiceOrderUpdate) AddAmount(i int64) *ExtraServiceOrderUpdate {
	esou.mutation.AddAmount(i)
	return esou
}

// SetSymbolID sets the "symbol_id" field.
func (esou *ExtraServiceOrderUpdate) SetSymbolID(i int64) *ExtraServiceOrderUpdate {
	esou.mutation.SetSymbolID(i)
	return esou
}

// SetNillableSymbolID sets the "symbol_id" field if the given value is not nil.
func (esou *ExtraServiceOrderUpdate) SetNillableSymbolID(i *int64) *ExtraServiceOrderUpdate {
	if i != nil {
		esou.SetSymbolID(*i)
	}
	return esou
}

// SetUnitCep sets the "unit_cep" field.
func (esou *ExtraServiceOrderUpdate) SetUnitCep(i int64) *ExtraServiceOrderUpdate {
	esou.mutation.ResetUnitCep()
	esou.mutation.SetUnitCep(i)
	return esou
}

// SetNillableUnitCep sets the "unit_cep" field if the given value is not nil.
func (esou *ExtraServiceOrderUpdate) SetNillableUnitCep(i *int64) *ExtraServiceOrderUpdate {
	if i != nil {
		esou.SetUnitCep(*i)
	}
	return esou
}

// AddUnitCep adds i to the "unit_cep" field.
func (esou *ExtraServiceOrderUpdate) AddUnitCep(i int64) *ExtraServiceOrderUpdate {
	esou.mutation.AddUnitCep(i)
	return esou
}

// SetExtraServiceType sets the "extra_service_type" field.
func (esou *ExtraServiceOrderUpdate) SetExtraServiceType(est enums.ExtraServiceType) *ExtraServiceOrderUpdate {
	esou.mutation.SetExtraServiceType(est)
	return esou
}

// SetNillableExtraServiceType sets the "extra_service_type" field if the given value is not nil.
func (esou *ExtraServiceOrderUpdate) SetNillableExtraServiceType(est *enums.ExtraServiceType) *ExtraServiceOrderUpdate {
	if est != nil {
		esou.SetExtraServiceType(*est)
	}
	return esou
}

// SetBuyDuration sets the "buy_duration" field.
func (esou *ExtraServiceOrderUpdate) SetBuyDuration(i int64) *ExtraServiceOrderUpdate {
	esou.mutation.ResetBuyDuration()
	esou.mutation.SetBuyDuration(i)
	return esou
}

// SetNillableBuyDuration sets the "buy_duration" field if the given value is not nil.
func (esou *ExtraServiceOrderUpdate) SetNillableBuyDuration(i *int64) *ExtraServiceOrderUpdate {
	if i != nil {
		esou.SetBuyDuration(*i)
	}
	return esou
}

// AddBuyDuration adds i to the "buy_duration" field.
func (esou *ExtraServiceOrderUpdate) AddBuyDuration(i int64) *ExtraServiceOrderUpdate {
	esou.mutation.AddBuyDuration(i)
	return esou
}

// SetPlanStartedAt sets the "plan_started_at" field.
func (esou *ExtraServiceOrderUpdate) SetPlanStartedAt(t time.Time) *ExtraServiceOrderUpdate {
	esou.mutation.SetPlanStartedAt(t)
	return esou
}

// SetNillablePlanStartedAt sets the "plan_started_at" field if the given value is not nil.
func (esou *ExtraServiceOrderUpdate) SetNillablePlanStartedAt(t *time.Time) *ExtraServiceOrderUpdate {
	if t != nil {
		esou.SetPlanStartedAt(*t)
	}
	return esou
}

// ClearPlanStartedAt clears the value of the "plan_started_at" field.
func (esou *ExtraServiceOrderUpdate) ClearPlanStartedAt() *ExtraServiceOrderUpdate {
	esou.mutation.ClearPlanStartedAt()
	return esou
}

// SetPlanFinishedAt sets the "plan_finished_at" field.
func (esou *ExtraServiceOrderUpdate) SetPlanFinishedAt(t time.Time) *ExtraServiceOrderUpdate {
	esou.mutation.SetPlanFinishedAt(t)
	return esou
}

// SetNillablePlanFinishedAt sets the "plan_finished_at" field if the given value is not nil.
func (esou *ExtraServiceOrderUpdate) SetNillablePlanFinishedAt(t *time.Time) *ExtraServiceOrderUpdate {
	if t != nil {
		esou.SetPlanFinishedAt(*t)
	}
	return esou
}

// ClearPlanFinishedAt clears the value of the "plan_finished_at" field.
func (esou *ExtraServiceOrderUpdate) ClearPlanFinishedAt() *ExtraServiceOrderUpdate {
	esou.mutation.ClearPlanFinishedAt()
	return esou
}

// SetMissionBatchID sets the "mission_batch_id" field.
func (esou *ExtraServiceOrderUpdate) SetMissionBatchID(i int64) *ExtraServiceOrderUpdate {
	esou.mutation.SetMissionBatchID(i)
	return esou
}

// SetNillableMissionBatchID sets the "mission_batch_id" field if the given value is not nil.
func (esou *ExtraServiceOrderUpdate) SetNillableMissionBatchID(i *int64) *ExtraServiceOrderUpdate {
	if i != nil {
		esou.SetMissionBatchID(*i)
	}
	return esou
}

// SetMission sets the "mission" edge to the Mission entity.
func (esou *ExtraServiceOrderUpdate) SetMission(m *Mission) *ExtraServiceOrderUpdate {
	return esou.SetMissionID(m.ID)
}

// SetMissionOrder sets the "mission_order" edge to the MissionOrder entity.
func (esou *ExtraServiceOrderUpdate) SetMissionOrder(m *MissionOrder) *ExtraServiceOrderUpdate {
	return esou.SetMissionOrderID(m.ID)
}

// SetSymbol sets the "symbol" edge to the Symbol entity.
func (esou *ExtraServiceOrderUpdate) SetSymbol(s *Symbol) *ExtraServiceOrderUpdate {
	return esou.SetSymbolID(s.ID)
}

// SetMissionBatch sets the "mission_batch" edge to the MissionBatch entity.
func (esou *ExtraServiceOrderUpdate) SetMissionBatch(m *MissionBatch) *ExtraServiceOrderUpdate {
	return esou.SetMissionBatchID(m.ID)
}

// Mutation returns the ExtraServiceOrderMutation object of the builder.
func (esou *ExtraServiceOrderUpdate) Mutation() *ExtraServiceOrderMutation {
	return esou.mutation
}

// ClearMission clears the "mission" edge to the Mission entity.
func (esou *ExtraServiceOrderUpdate) ClearMission() *ExtraServiceOrderUpdate {
	esou.mutation.ClearMission()
	return esou
}

// ClearMissionOrder clears the "mission_order" edge to the MissionOrder entity.
func (esou *ExtraServiceOrderUpdate) ClearMissionOrder() *ExtraServiceOrderUpdate {
	esou.mutation.ClearMissionOrder()
	return esou
}

// ClearSymbol clears the "symbol" edge to the Symbol entity.
func (esou *ExtraServiceOrderUpdate) ClearSymbol() *ExtraServiceOrderUpdate {
	esou.mutation.ClearSymbol()
	return esou
}

// ClearMissionBatch clears the "mission_batch" edge to the MissionBatch entity.
func (esou *ExtraServiceOrderUpdate) ClearMissionBatch() *ExtraServiceOrderUpdate {
	esou.mutation.ClearMissionBatch()
	return esou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (esou *ExtraServiceOrderUpdate) Save(ctx context.Context) (int, error) {
	esou.defaults()
	return withHooks(ctx, esou.sqlSave, esou.mutation, esou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (esou *ExtraServiceOrderUpdate) SaveX(ctx context.Context) int {
	affected, err := esou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (esou *ExtraServiceOrderUpdate) Exec(ctx context.Context) error {
	_, err := esou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (esou *ExtraServiceOrderUpdate) ExecX(ctx context.Context) {
	if err := esou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (esou *ExtraServiceOrderUpdate) defaults() {
	if _, ok := esou.mutation.UpdatedAt(); !ok {
		v := extraserviceorder.UpdateDefaultUpdatedAt()
		esou.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (esou *ExtraServiceOrderUpdate) check() error {
	if v, ok := esou.mutation.ExtraServiceType(); ok {
		if err := extraserviceorder.ExtraServiceTypeValidator(v); err != nil {
			return &ValidationError{Name: "extra_service_type", err: fmt.Errorf(`cep_ent: validator failed for field "ExtraServiceOrder.extra_service_type": %w`, err)}
		}
	}
	if _, ok := esou.mutation.MissionID(); esou.mutation.MissionCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ExtraServiceOrder.mission"`)
	}
	if _, ok := esou.mutation.MissionOrderID(); esou.mutation.MissionOrderCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ExtraServiceOrder.mission_order"`)
	}
	if _, ok := esou.mutation.SymbolID(); esou.mutation.SymbolCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ExtraServiceOrder.symbol"`)
	}
	if _, ok := esou.mutation.MissionBatchID(); esou.mutation.MissionBatchCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ExtraServiceOrder.mission_batch"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (esou *ExtraServiceOrderUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ExtraServiceOrderUpdate {
	esou.modifiers = append(esou.modifiers, modifiers...)
	return esou
}

func (esou *ExtraServiceOrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := esou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(extraserviceorder.Table, extraserviceorder.Columns, sqlgraph.NewFieldSpec(extraserviceorder.FieldID, field.TypeInt64))
	if ps := esou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := esou.mutation.CreatedBy(); ok {
		_spec.SetField(extraserviceorder.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := esou.mutation.AddedCreatedBy(); ok {
		_spec.AddField(extraserviceorder.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := esou.mutation.UpdatedBy(); ok {
		_spec.SetField(extraserviceorder.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := esou.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(extraserviceorder.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := esou.mutation.UpdatedAt(); ok {
		_spec.SetField(extraserviceorder.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := esou.mutation.DeletedAt(); ok {
		_spec.SetField(extraserviceorder.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := esou.mutation.Amount(); ok {
		_spec.SetField(extraserviceorder.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := esou.mutation.AddedAmount(); ok {
		_spec.AddField(extraserviceorder.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := esou.mutation.UnitCep(); ok {
		_spec.SetField(extraserviceorder.FieldUnitCep, field.TypeInt64, value)
	}
	if value, ok := esou.mutation.AddedUnitCep(); ok {
		_spec.AddField(extraserviceorder.FieldUnitCep, field.TypeInt64, value)
	}
	if value, ok := esou.mutation.ExtraServiceType(); ok {
		_spec.SetField(extraserviceorder.FieldExtraServiceType, field.TypeEnum, value)
	}
	if value, ok := esou.mutation.BuyDuration(); ok {
		_spec.SetField(extraserviceorder.FieldBuyDuration, field.TypeInt64, value)
	}
	if value, ok := esou.mutation.AddedBuyDuration(); ok {
		_spec.AddField(extraserviceorder.FieldBuyDuration, field.TypeInt64, value)
	}
	if value, ok := esou.mutation.PlanStartedAt(); ok {
		_spec.SetField(extraserviceorder.FieldPlanStartedAt, field.TypeTime, value)
	}
	if esou.mutation.PlanStartedAtCleared() {
		_spec.ClearField(extraserviceorder.FieldPlanStartedAt, field.TypeTime)
	}
	if value, ok := esou.mutation.PlanFinishedAt(); ok {
		_spec.SetField(extraserviceorder.FieldPlanFinishedAt, field.TypeTime, value)
	}
	if esou.mutation.PlanFinishedAtCleared() {
		_spec.ClearField(extraserviceorder.FieldPlanFinishedAt, field.TypeTime)
	}
	if esou.mutation.MissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.MissionTable,
			Columns: []string{extraserviceorder.MissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mission.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esou.mutation.MissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.MissionTable,
			Columns: []string{extraserviceorder.MissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if esou.mutation.MissionOrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.MissionOrderTable,
			Columns: []string{extraserviceorder.MissionOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionorder.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esou.mutation.MissionOrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.MissionOrderTable,
			Columns: []string{extraserviceorder.MissionOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if esou.mutation.SymbolCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.SymbolTable,
			Columns: []string{extraserviceorder.SymbolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(symbol.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esou.mutation.SymbolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.SymbolTable,
			Columns: []string{extraserviceorder.SymbolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(symbol.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if esou.mutation.MissionBatchCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.MissionBatchTable,
			Columns: []string{extraserviceorder.MissionBatchColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionbatch.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esou.mutation.MissionBatchIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.MissionBatchTable,
			Columns: []string{extraserviceorder.MissionBatchColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionbatch.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(esou.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, esou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{extraserviceorder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	esou.mutation.done = true
	return n, nil
}

// ExtraServiceOrderUpdateOne is the builder for updating a single ExtraServiceOrder entity.
type ExtraServiceOrderUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ExtraServiceOrderMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (esouo *ExtraServiceOrderUpdateOne) SetCreatedBy(i int64) *ExtraServiceOrderUpdateOne {
	esouo.mutation.ResetCreatedBy()
	esouo.mutation.SetCreatedBy(i)
	return esouo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (esouo *ExtraServiceOrderUpdateOne) SetNillableCreatedBy(i *int64) *ExtraServiceOrderUpdateOne {
	if i != nil {
		esouo.SetCreatedBy(*i)
	}
	return esouo
}

// AddCreatedBy adds i to the "created_by" field.
func (esouo *ExtraServiceOrderUpdateOne) AddCreatedBy(i int64) *ExtraServiceOrderUpdateOne {
	esouo.mutation.AddCreatedBy(i)
	return esouo
}

// SetUpdatedBy sets the "updated_by" field.
func (esouo *ExtraServiceOrderUpdateOne) SetUpdatedBy(i int64) *ExtraServiceOrderUpdateOne {
	esouo.mutation.ResetUpdatedBy()
	esouo.mutation.SetUpdatedBy(i)
	return esouo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (esouo *ExtraServiceOrderUpdateOne) SetNillableUpdatedBy(i *int64) *ExtraServiceOrderUpdateOne {
	if i != nil {
		esouo.SetUpdatedBy(*i)
	}
	return esouo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (esouo *ExtraServiceOrderUpdateOne) AddUpdatedBy(i int64) *ExtraServiceOrderUpdateOne {
	esouo.mutation.AddUpdatedBy(i)
	return esouo
}

// SetUpdatedAt sets the "updated_at" field.
func (esouo *ExtraServiceOrderUpdateOne) SetUpdatedAt(t time.Time) *ExtraServiceOrderUpdateOne {
	esouo.mutation.SetUpdatedAt(t)
	return esouo
}

// SetDeletedAt sets the "deleted_at" field.
func (esouo *ExtraServiceOrderUpdateOne) SetDeletedAt(t time.Time) *ExtraServiceOrderUpdateOne {
	esouo.mutation.SetDeletedAt(t)
	return esouo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (esouo *ExtraServiceOrderUpdateOne) SetNillableDeletedAt(t *time.Time) *ExtraServiceOrderUpdateOne {
	if t != nil {
		esouo.SetDeletedAt(*t)
	}
	return esouo
}

// SetMissionID sets the "mission_id" field.
func (esouo *ExtraServiceOrderUpdateOne) SetMissionID(i int64) *ExtraServiceOrderUpdateOne {
	esouo.mutation.SetMissionID(i)
	return esouo
}

// SetNillableMissionID sets the "mission_id" field if the given value is not nil.
func (esouo *ExtraServiceOrderUpdateOne) SetNillableMissionID(i *int64) *ExtraServiceOrderUpdateOne {
	if i != nil {
		esouo.SetMissionID(*i)
	}
	return esouo
}

// SetMissionOrderID sets the "mission_order_id" field.
func (esouo *ExtraServiceOrderUpdateOne) SetMissionOrderID(i int64) *ExtraServiceOrderUpdateOne {
	esouo.mutation.SetMissionOrderID(i)
	return esouo
}

// SetNillableMissionOrderID sets the "mission_order_id" field if the given value is not nil.
func (esouo *ExtraServiceOrderUpdateOne) SetNillableMissionOrderID(i *int64) *ExtraServiceOrderUpdateOne {
	if i != nil {
		esouo.SetMissionOrderID(*i)
	}
	return esouo
}

// SetAmount sets the "amount" field.
func (esouo *ExtraServiceOrderUpdateOne) SetAmount(i int64) *ExtraServiceOrderUpdateOne {
	esouo.mutation.ResetAmount()
	esouo.mutation.SetAmount(i)
	return esouo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (esouo *ExtraServiceOrderUpdateOne) SetNillableAmount(i *int64) *ExtraServiceOrderUpdateOne {
	if i != nil {
		esouo.SetAmount(*i)
	}
	return esouo
}

// AddAmount adds i to the "amount" field.
func (esouo *ExtraServiceOrderUpdateOne) AddAmount(i int64) *ExtraServiceOrderUpdateOne {
	esouo.mutation.AddAmount(i)
	return esouo
}

// SetSymbolID sets the "symbol_id" field.
func (esouo *ExtraServiceOrderUpdateOne) SetSymbolID(i int64) *ExtraServiceOrderUpdateOne {
	esouo.mutation.SetSymbolID(i)
	return esouo
}

// SetNillableSymbolID sets the "symbol_id" field if the given value is not nil.
func (esouo *ExtraServiceOrderUpdateOne) SetNillableSymbolID(i *int64) *ExtraServiceOrderUpdateOne {
	if i != nil {
		esouo.SetSymbolID(*i)
	}
	return esouo
}

// SetUnitCep sets the "unit_cep" field.
func (esouo *ExtraServiceOrderUpdateOne) SetUnitCep(i int64) *ExtraServiceOrderUpdateOne {
	esouo.mutation.ResetUnitCep()
	esouo.mutation.SetUnitCep(i)
	return esouo
}

// SetNillableUnitCep sets the "unit_cep" field if the given value is not nil.
func (esouo *ExtraServiceOrderUpdateOne) SetNillableUnitCep(i *int64) *ExtraServiceOrderUpdateOne {
	if i != nil {
		esouo.SetUnitCep(*i)
	}
	return esouo
}

// AddUnitCep adds i to the "unit_cep" field.
func (esouo *ExtraServiceOrderUpdateOne) AddUnitCep(i int64) *ExtraServiceOrderUpdateOne {
	esouo.mutation.AddUnitCep(i)
	return esouo
}

// SetExtraServiceType sets the "extra_service_type" field.
func (esouo *ExtraServiceOrderUpdateOne) SetExtraServiceType(est enums.ExtraServiceType) *ExtraServiceOrderUpdateOne {
	esouo.mutation.SetExtraServiceType(est)
	return esouo
}

// SetNillableExtraServiceType sets the "extra_service_type" field if the given value is not nil.
func (esouo *ExtraServiceOrderUpdateOne) SetNillableExtraServiceType(est *enums.ExtraServiceType) *ExtraServiceOrderUpdateOne {
	if est != nil {
		esouo.SetExtraServiceType(*est)
	}
	return esouo
}

// SetBuyDuration sets the "buy_duration" field.
func (esouo *ExtraServiceOrderUpdateOne) SetBuyDuration(i int64) *ExtraServiceOrderUpdateOne {
	esouo.mutation.ResetBuyDuration()
	esouo.mutation.SetBuyDuration(i)
	return esouo
}

// SetNillableBuyDuration sets the "buy_duration" field if the given value is not nil.
func (esouo *ExtraServiceOrderUpdateOne) SetNillableBuyDuration(i *int64) *ExtraServiceOrderUpdateOne {
	if i != nil {
		esouo.SetBuyDuration(*i)
	}
	return esouo
}

// AddBuyDuration adds i to the "buy_duration" field.
func (esouo *ExtraServiceOrderUpdateOne) AddBuyDuration(i int64) *ExtraServiceOrderUpdateOne {
	esouo.mutation.AddBuyDuration(i)
	return esouo
}

// SetPlanStartedAt sets the "plan_started_at" field.
func (esouo *ExtraServiceOrderUpdateOne) SetPlanStartedAt(t time.Time) *ExtraServiceOrderUpdateOne {
	esouo.mutation.SetPlanStartedAt(t)
	return esouo
}

// SetNillablePlanStartedAt sets the "plan_started_at" field if the given value is not nil.
func (esouo *ExtraServiceOrderUpdateOne) SetNillablePlanStartedAt(t *time.Time) *ExtraServiceOrderUpdateOne {
	if t != nil {
		esouo.SetPlanStartedAt(*t)
	}
	return esouo
}

// ClearPlanStartedAt clears the value of the "plan_started_at" field.
func (esouo *ExtraServiceOrderUpdateOne) ClearPlanStartedAt() *ExtraServiceOrderUpdateOne {
	esouo.mutation.ClearPlanStartedAt()
	return esouo
}

// SetPlanFinishedAt sets the "plan_finished_at" field.
func (esouo *ExtraServiceOrderUpdateOne) SetPlanFinishedAt(t time.Time) *ExtraServiceOrderUpdateOne {
	esouo.mutation.SetPlanFinishedAt(t)
	return esouo
}

// SetNillablePlanFinishedAt sets the "plan_finished_at" field if the given value is not nil.
func (esouo *ExtraServiceOrderUpdateOne) SetNillablePlanFinishedAt(t *time.Time) *ExtraServiceOrderUpdateOne {
	if t != nil {
		esouo.SetPlanFinishedAt(*t)
	}
	return esouo
}

// ClearPlanFinishedAt clears the value of the "plan_finished_at" field.
func (esouo *ExtraServiceOrderUpdateOne) ClearPlanFinishedAt() *ExtraServiceOrderUpdateOne {
	esouo.mutation.ClearPlanFinishedAt()
	return esouo
}

// SetMissionBatchID sets the "mission_batch_id" field.
func (esouo *ExtraServiceOrderUpdateOne) SetMissionBatchID(i int64) *ExtraServiceOrderUpdateOne {
	esouo.mutation.SetMissionBatchID(i)
	return esouo
}

// SetNillableMissionBatchID sets the "mission_batch_id" field if the given value is not nil.
func (esouo *ExtraServiceOrderUpdateOne) SetNillableMissionBatchID(i *int64) *ExtraServiceOrderUpdateOne {
	if i != nil {
		esouo.SetMissionBatchID(*i)
	}
	return esouo
}

// SetMission sets the "mission" edge to the Mission entity.
func (esouo *ExtraServiceOrderUpdateOne) SetMission(m *Mission) *ExtraServiceOrderUpdateOne {
	return esouo.SetMissionID(m.ID)
}

// SetMissionOrder sets the "mission_order" edge to the MissionOrder entity.
func (esouo *ExtraServiceOrderUpdateOne) SetMissionOrder(m *MissionOrder) *ExtraServiceOrderUpdateOne {
	return esouo.SetMissionOrderID(m.ID)
}

// SetSymbol sets the "symbol" edge to the Symbol entity.
func (esouo *ExtraServiceOrderUpdateOne) SetSymbol(s *Symbol) *ExtraServiceOrderUpdateOne {
	return esouo.SetSymbolID(s.ID)
}

// SetMissionBatch sets the "mission_batch" edge to the MissionBatch entity.
func (esouo *ExtraServiceOrderUpdateOne) SetMissionBatch(m *MissionBatch) *ExtraServiceOrderUpdateOne {
	return esouo.SetMissionBatchID(m.ID)
}

// Mutation returns the ExtraServiceOrderMutation object of the builder.
func (esouo *ExtraServiceOrderUpdateOne) Mutation() *ExtraServiceOrderMutation {
	return esouo.mutation
}

// ClearMission clears the "mission" edge to the Mission entity.
func (esouo *ExtraServiceOrderUpdateOne) ClearMission() *ExtraServiceOrderUpdateOne {
	esouo.mutation.ClearMission()
	return esouo
}

// ClearMissionOrder clears the "mission_order" edge to the MissionOrder entity.
func (esouo *ExtraServiceOrderUpdateOne) ClearMissionOrder() *ExtraServiceOrderUpdateOne {
	esouo.mutation.ClearMissionOrder()
	return esouo
}

// ClearSymbol clears the "symbol" edge to the Symbol entity.
func (esouo *ExtraServiceOrderUpdateOne) ClearSymbol() *ExtraServiceOrderUpdateOne {
	esouo.mutation.ClearSymbol()
	return esouo
}

// ClearMissionBatch clears the "mission_batch" edge to the MissionBatch entity.
func (esouo *ExtraServiceOrderUpdateOne) ClearMissionBatch() *ExtraServiceOrderUpdateOne {
	esouo.mutation.ClearMissionBatch()
	return esouo
}

// Where appends a list predicates to the ExtraServiceOrderUpdate builder.
func (esouo *ExtraServiceOrderUpdateOne) Where(ps ...predicate.ExtraServiceOrder) *ExtraServiceOrderUpdateOne {
	esouo.mutation.Where(ps...)
	return esouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (esouo *ExtraServiceOrderUpdateOne) Select(field string, fields ...string) *ExtraServiceOrderUpdateOne {
	esouo.fields = append([]string{field}, fields...)
	return esouo
}

// Save executes the query and returns the updated ExtraServiceOrder entity.
func (esouo *ExtraServiceOrderUpdateOne) Save(ctx context.Context) (*ExtraServiceOrder, error) {
	esouo.defaults()
	return withHooks(ctx, esouo.sqlSave, esouo.mutation, esouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (esouo *ExtraServiceOrderUpdateOne) SaveX(ctx context.Context) *ExtraServiceOrder {
	node, err := esouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (esouo *ExtraServiceOrderUpdateOne) Exec(ctx context.Context) error {
	_, err := esouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (esouo *ExtraServiceOrderUpdateOne) ExecX(ctx context.Context) {
	if err := esouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (esouo *ExtraServiceOrderUpdateOne) defaults() {
	if _, ok := esouo.mutation.UpdatedAt(); !ok {
		v := extraserviceorder.UpdateDefaultUpdatedAt()
		esouo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (esouo *ExtraServiceOrderUpdateOne) check() error {
	if v, ok := esouo.mutation.ExtraServiceType(); ok {
		if err := extraserviceorder.ExtraServiceTypeValidator(v); err != nil {
			return &ValidationError{Name: "extra_service_type", err: fmt.Errorf(`cep_ent: validator failed for field "ExtraServiceOrder.extra_service_type": %w`, err)}
		}
	}
	if _, ok := esouo.mutation.MissionID(); esouo.mutation.MissionCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ExtraServiceOrder.mission"`)
	}
	if _, ok := esouo.mutation.MissionOrderID(); esouo.mutation.MissionOrderCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ExtraServiceOrder.mission_order"`)
	}
	if _, ok := esouo.mutation.SymbolID(); esouo.mutation.SymbolCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ExtraServiceOrder.symbol"`)
	}
	if _, ok := esouo.mutation.MissionBatchID(); esouo.mutation.MissionBatchCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ExtraServiceOrder.mission_batch"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (esouo *ExtraServiceOrderUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ExtraServiceOrderUpdateOne {
	esouo.modifiers = append(esouo.modifiers, modifiers...)
	return esouo
}

func (esouo *ExtraServiceOrderUpdateOne) sqlSave(ctx context.Context) (_node *ExtraServiceOrder, err error) {
	if err := esouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(extraserviceorder.Table, extraserviceorder.Columns, sqlgraph.NewFieldSpec(extraserviceorder.FieldID, field.TypeInt64))
	id, ok := esouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "ExtraServiceOrder.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := esouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, extraserviceorder.FieldID)
		for _, f := range fields {
			if !extraserviceorder.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != extraserviceorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := esouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := esouo.mutation.CreatedBy(); ok {
		_spec.SetField(extraserviceorder.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := esouo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(extraserviceorder.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := esouo.mutation.UpdatedBy(); ok {
		_spec.SetField(extraserviceorder.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := esouo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(extraserviceorder.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := esouo.mutation.UpdatedAt(); ok {
		_spec.SetField(extraserviceorder.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := esouo.mutation.DeletedAt(); ok {
		_spec.SetField(extraserviceorder.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := esouo.mutation.Amount(); ok {
		_spec.SetField(extraserviceorder.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := esouo.mutation.AddedAmount(); ok {
		_spec.AddField(extraserviceorder.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := esouo.mutation.UnitCep(); ok {
		_spec.SetField(extraserviceorder.FieldUnitCep, field.TypeInt64, value)
	}
	if value, ok := esouo.mutation.AddedUnitCep(); ok {
		_spec.AddField(extraserviceorder.FieldUnitCep, field.TypeInt64, value)
	}
	if value, ok := esouo.mutation.ExtraServiceType(); ok {
		_spec.SetField(extraserviceorder.FieldExtraServiceType, field.TypeEnum, value)
	}
	if value, ok := esouo.mutation.BuyDuration(); ok {
		_spec.SetField(extraserviceorder.FieldBuyDuration, field.TypeInt64, value)
	}
	if value, ok := esouo.mutation.AddedBuyDuration(); ok {
		_spec.AddField(extraserviceorder.FieldBuyDuration, field.TypeInt64, value)
	}
	if value, ok := esouo.mutation.PlanStartedAt(); ok {
		_spec.SetField(extraserviceorder.FieldPlanStartedAt, field.TypeTime, value)
	}
	if esouo.mutation.PlanStartedAtCleared() {
		_spec.ClearField(extraserviceorder.FieldPlanStartedAt, field.TypeTime)
	}
	if value, ok := esouo.mutation.PlanFinishedAt(); ok {
		_spec.SetField(extraserviceorder.FieldPlanFinishedAt, field.TypeTime, value)
	}
	if esouo.mutation.PlanFinishedAtCleared() {
		_spec.ClearField(extraserviceorder.FieldPlanFinishedAt, field.TypeTime)
	}
	if esouo.mutation.MissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.MissionTable,
			Columns: []string{extraserviceorder.MissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mission.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esouo.mutation.MissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.MissionTable,
			Columns: []string{extraserviceorder.MissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if esouo.mutation.MissionOrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.MissionOrderTable,
			Columns: []string{extraserviceorder.MissionOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionorder.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esouo.mutation.MissionOrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.MissionOrderTable,
			Columns: []string{extraserviceorder.MissionOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if esouo.mutation.SymbolCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.SymbolTable,
			Columns: []string{extraserviceorder.SymbolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(symbol.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esouo.mutation.SymbolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.SymbolTable,
			Columns: []string{extraserviceorder.SymbolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(symbol.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if esouo.mutation.MissionBatchCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.MissionBatchTable,
			Columns: []string{extraserviceorder.MissionBatchColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionbatch.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esouo.mutation.MissionBatchIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   extraserviceorder.MissionBatchTable,
			Columns: []string{extraserviceorder.MissionBatchColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionbatch.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(esouo.modifiers...)
	_node = &ExtraServiceOrder{config: esouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, esouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{extraserviceorder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	esouo.mutation.done = true
	return _node, nil
}
