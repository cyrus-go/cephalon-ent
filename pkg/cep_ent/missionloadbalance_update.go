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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionloadbalance"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// MissionLoadBalanceUpdate is the builder for updating MissionLoadBalance entities.
type MissionLoadBalanceUpdate struct {
	config
	hooks     []Hook
	mutation  *MissionLoadBalanceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MissionLoadBalanceUpdate builder.
func (mlbu *MissionLoadBalanceUpdate) Where(ps ...predicate.MissionLoadBalance) *MissionLoadBalanceUpdate {
	mlbu.mutation.Where(ps...)
	return mlbu
}

// SetCreatedBy sets the "created_by" field.
func (mlbu *MissionLoadBalanceUpdate) SetCreatedBy(i int64) *MissionLoadBalanceUpdate {
	mlbu.mutation.ResetCreatedBy()
	mlbu.mutation.SetCreatedBy(i)
	return mlbu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mlbu *MissionLoadBalanceUpdate) SetNillableCreatedBy(i *int64) *MissionLoadBalanceUpdate {
	if i != nil {
		mlbu.SetCreatedBy(*i)
	}
	return mlbu
}

// AddCreatedBy adds i to the "created_by" field.
func (mlbu *MissionLoadBalanceUpdate) AddCreatedBy(i int64) *MissionLoadBalanceUpdate {
	mlbu.mutation.AddCreatedBy(i)
	return mlbu
}

// SetUpdatedBy sets the "updated_by" field.
func (mlbu *MissionLoadBalanceUpdate) SetUpdatedBy(i int64) *MissionLoadBalanceUpdate {
	mlbu.mutation.ResetUpdatedBy()
	mlbu.mutation.SetUpdatedBy(i)
	return mlbu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mlbu *MissionLoadBalanceUpdate) SetNillableUpdatedBy(i *int64) *MissionLoadBalanceUpdate {
	if i != nil {
		mlbu.SetUpdatedBy(*i)
	}
	return mlbu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (mlbu *MissionLoadBalanceUpdate) AddUpdatedBy(i int64) *MissionLoadBalanceUpdate {
	mlbu.mutation.AddUpdatedBy(i)
	return mlbu
}

// SetUpdatedAt sets the "updated_at" field.
func (mlbu *MissionLoadBalanceUpdate) SetUpdatedAt(t time.Time) *MissionLoadBalanceUpdate {
	mlbu.mutation.SetUpdatedAt(t)
	return mlbu
}

// SetDeletedAt sets the "deleted_at" field.
func (mlbu *MissionLoadBalanceUpdate) SetDeletedAt(t time.Time) *MissionLoadBalanceUpdate {
	mlbu.mutation.SetDeletedAt(t)
	return mlbu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mlbu *MissionLoadBalanceUpdate) SetNillableDeletedAt(t *time.Time) *MissionLoadBalanceUpdate {
	if t != nil {
		mlbu.SetDeletedAt(*t)
	}
	return mlbu
}

// SetMissionType sets the "mission_type" field.
func (mlbu *MissionLoadBalanceUpdate) SetMissionType(et enums.MissionType) *MissionLoadBalanceUpdate {
	mlbu.mutation.SetMissionType(et)
	return mlbu
}

// SetNillableMissionType sets the "mission_type" field if the given value is not nil.
func (mlbu *MissionLoadBalanceUpdate) SetNillableMissionType(et *enums.MissionType) *MissionLoadBalanceUpdate {
	if et != nil {
		mlbu.SetMissionType(*et)
	}
	return mlbu
}

// SetUserID sets the "user_id" field.
func (mlbu *MissionLoadBalanceUpdate) SetUserID(i int64) *MissionLoadBalanceUpdate {
	mlbu.mutation.ResetUserID()
	mlbu.mutation.SetUserID(i)
	return mlbu
}

// AddUserID adds i to the "user_id" field.
func (mlbu *MissionLoadBalanceUpdate) AddUserID(i int64) *MissionLoadBalanceUpdate {
	mlbu.mutation.AddUserID(i)
	return mlbu
}

// SetStartedAt sets the "started_at" field.
func (mlbu *MissionLoadBalanceUpdate) SetStartedAt(t time.Time) *MissionLoadBalanceUpdate {
	mlbu.mutation.SetStartedAt(t)
	return mlbu
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (mlbu *MissionLoadBalanceUpdate) SetNillableStartedAt(t *time.Time) *MissionLoadBalanceUpdate {
	if t != nil {
		mlbu.SetStartedAt(*t)
	}
	return mlbu
}

// SetFinishedAt sets the "finished_at" field.
func (mlbu *MissionLoadBalanceUpdate) SetFinishedAt(t time.Time) *MissionLoadBalanceUpdate {
	mlbu.mutation.SetFinishedAt(t)
	return mlbu
}

// SetNillableFinishedAt sets the "finished_at" field if the given value is not nil.
func (mlbu *MissionLoadBalanceUpdate) SetNillableFinishedAt(t *time.Time) *MissionLoadBalanceUpdate {
	if t != nil {
		mlbu.SetFinishedAt(*t)
	}
	return mlbu
}

// SetGpuVersion sets the "gpu_version" field.
func (mlbu *MissionLoadBalanceUpdate) SetGpuVersion(ev enums.GpuVersion) *MissionLoadBalanceUpdate {
	mlbu.mutation.SetGpuVersion(ev)
	return mlbu
}

// SetNillableGpuVersion sets the "gpu_version" field if the given value is not nil.
func (mlbu *MissionLoadBalanceUpdate) SetNillableGpuVersion(ev *enums.GpuVersion) *MissionLoadBalanceUpdate {
	if ev != nil {
		mlbu.SetGpuVersion(*ev)
	}
	return mlbu
}

// SetGpuNum sets the "gpu_num" field.
func (mlbu *MissionLoadBalanceUpdate) SetGpuNum(i int8) *MissionLoadBalanceUpdate {
	mlbu.mutation.ResetGpuNum()
	mlbu.mutation.SetGpuNum(i)
	return mlbu
}

// SetNillableGpuNum sets the "gpu_num" field if the given value is not nil.
func (mlbu *MissionLoadBalanceUpdate) SetNillableGpuNum(i *int8) *MissionLoadBalanceUpdate {
	if i != nil {
		mlbu.SetGpuNum(*i)
	}
	return mlbu
}

// AddGpuNum adds i to the "gpu_num" field.
func (mlbu *MissionLoadBalanceUpdate) AddGpuNum(i int8) *MissionLoadBalanceUpdate {
	mlbu.mutation.AddGpuNum(i)
	return mlbu
}

// SetMaxMissionCount sets the "max_mission_count" field.
func (mlbu *MissionLoadBalanceUpdate) SetMaxMissionCount(i int8) *MissionLoadBalanceUpdate {
	mlbu.mutation.ResetMaxMissionCount()
	mlbu.mutation.SetMaxMissionCount(i)
	return mlbu
}

// SetNillableMaxMissionCount sets the "max_mission_count" field if the given value is not nil.
func (mlbu *MissionLoadBalanceUpdate) SetNillableMaxMissionCount(i *int8) *MissionLoadBalanceUpdate {
	if i != nil {
		mlbu.SetMaxMissionCount(*i)
	}
	return mlbu
}

// AddMaxMissionCount adds i to the "max_mission_count" field.
func (mlbu *MissionLoadBalanceUpdate) AddMaxMissionCount(i int8) *MissionLoadBalanceUpdate {
	mlbu.mutation.AddMaxMissionCount(i)
	return mlbu
}

// SetMinMissionCount sets the "min_mission_count" field.
func (mlbu *MissionLoadBalanceUpdate) SetMinMissionCount(i int8) *MissionLoadBalanceUpdate {
	mlbu.mutation.ResetMinMissionCount()
	mlbu.mutation.SetMinMissionCount(i)
	return mlbu
}

// SetNillableMinMissionCount sets the "min_mission_count" field if the given value is not nil.
func (mlbu *MissionLoadBalanceUpdate) SetNillableMinMissionCount(i *int8) *MissionLoadBalanceUpdate {
	if i != nil {
		mlbu.SetMinMissionCount(*i)
	}
	return mlbu
}

// AddMinMissionCount adds i to the "min_mission_count" field.
func (mlbu *MissionLoadBalanceUpdate) AddMinMissionCount(i int8) *MissionLoadBalanceUpdate {
	mlbu.mutation.AddMinMissionCount(i)
	return mlbu
}

// SetMissionBatchID sets the "mission_batch_id" field.
func (mlbu *MissionLoadBalanceUpdate) SetMissionBatchID(i int64) *MissionLoadBalanceUpdate {
	mlbu.mutation.ResetMissionBatchID()
	mlbu.mutation.SetMissionBatchID(i)
	return mlbu
}

// SetNillableMissionBatchID sets the "mission_batch_id" field if the given value is not nil.
func (mlbu *MissionLoadBalanceUpdate) SetNillableMissionBatchID(i *int64) *MissionLoadBalanceUpdate {
	if i != nil {
		mlbu.SetMissionBatchID(*i)
	}
	return mlbu
}

// AddMissionBatchID adds i to the "mission_batch_id" field.
func (mlbu *MissionLoadBalanceUpdate) AddMissionBatchID(i int64) *MissionLoadBalanceUpdate {
	mlbu.mutation.AddMissionBatchID(i)
	return mlbu
}

// SetMissionBatchNumber sets the "mission_batch_number" field.
func (mlbu *MissionLoadBalanceUpdate) SetMissionBatchNumber(s string) *MissionLoadBalanceUpdate {
	mlbu.mutation.SetMissionBatchNumber(s)
	return mlbu
}

// SetNillableMissionBatchNumber sets the "mission_batch_number" field if the given value is not nil.
func (mlbu *MissionLoadBalanceUpdate) SetNillableMissionBatchNumber(s *string) *MissionLoadBalanceUpdate {
	if s != nil {
		mlbu.SetMissionBatchNumber(*s)
	}
	return mlbu
}

// Mutation returns the MissionLoadBalanceMutation object of the builder.
func (mlbu *MissionLoadBalanceUpdate) Mutation() *MissionLoadBalanceMutation {
	return mlbu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mlbu *MissionLoadBalanceUpdate) Save(ctx context.Context) (int, error) {
	mlbu.defaults()
	return withHooks(ctx, mlbu.sqlSave, mlbu.mutation, mlbu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mlbu *MissionLoadBalanceUpdate) SaveX(ctx context.Context) int {
	affected, err := mlbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mlbu *MissionLoadBalanceUpdate) Exec(ctx context.Context) error {
	_, err := mlbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mlbu *MissionLoadBalanceUpdate) ExecX(ctx context.Context) {
	if err := mlbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mlbu *MissionLoadBalanceUpdate) defaults() {
	if _, ok := mlbu.mutation.UpdatedAt(); !ok {
		v := missionloadbalance.UpdateDefaultUpdatedAt()
		mlbu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mlbu *MissionLoadBalanceUpdate) check() error {
	if v, ok := mlbu.mutation.MissionType(); ok {
		if err := missionloadbalance.MissionTypeValidator(v); err != nil {
			return &ValidationError{Name: "mission_type", err: fmt.Errorf(`cep_ent: validator failed for field "MissionLoadBalance.mission_type": %w`, err)}
		}
	}
	if v, ok := mlbu.mutation.GpuVersion(); ok {
		if err := missionloadbalance.GpuVersionValidator(v); err != nil {
			return &ValidationError{Name: "gpu_version", err: fmt.Errorf(`cep_ent: validator failed for field "MissionLoadBalance.gpu_version": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mlbu *MissionLoadBalanceUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MissionLoadBalanceUpdate {
	mlbu.modifiers = append(mlbu.modifiers, modifiers...)
	return mlbu
}

func (mlbu *MissionLoadBalanceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mlbu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(missionloadbalance.Table, missionloadbalance.Columns, sqlgraph.NewFieldSpec(missionloadbalance.FieldID, field.TypeInt64))
	if ps := mlbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mlbu.mutation.CreatedBy(); ok {
		_spec.SetField(missionloadbalance.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mlbu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(missionloadbalance.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mlbu.mutation.UpdatedBy(); ok {
		_spec.SetField(missionloadbalance.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mlbu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(missionloadbalance.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mlbu.mutation.UpdatedAt(); ok {
		_spec.SetField(missionloadbalance.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mlbu.mutation.DeletedAt(); ok {
		_spec.SetField(missionloadbalance.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := mlbu.mutation.MissionType(); ok {
		_spec.SetField(missionloadbalance.FieldMissionType, field.TypeEnum, value)
	}
	if value, ok := mlbu.mutation.UserID(); ok {
		_spec.SetField(missionloadbalance.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := mlbu.mutation.AddedUserID(); ok {
		_spec.AddField(missionloadbalance.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := mlbu.mutation.StartedAt(); ok {
		_spec.SetField(missionloadbalance.FieldStartedAt, field.TypeTime, value)
	}
	if value, ok := mlbu.mutation.FinishedAt(); ok {
		_spec.SetField(missionloadbalance.FieldFinishedAt, field.TypeTime, value)
	}
	if value, ok := mlbu.mutation.GpuVersion(); ok {
		_spec.SetField(missionloadbalance.FieldGpuVersion, field.TypeEnum, value)
	}
	if value, ok := mlbu.mutation.GpuNum(); ok {
		_spec.SetField(missionloadbalance.FieldGpuNum, field.TypeInt8, value)
	}
	if value, ok := mlbu.mutation.AddedGpuNum(); ok {
		_spec.AddField(missionloadbalance.FieldGpuNum, field.TypeInt8, value)
	}
	if value, ok := mlbu.mutation.MaxMissionCount(); ok {
		_spec.SetField(missionloadbalance.FieldMaxMissionCount, field.TypeInt8, value)
	}
	if value, ok := mlbu.mutation.AddedMaxMissionCount(); ok {
		_spec.AddField(missionloadbalance.FieldMaxMissionCount, field.TypeInt8, value)
	}
	if value, ok := mlbu.mutation.MinMissionCount(); ok {
		_spec.SetField(missionloadbalance.FieldMinMissionCount, field.TypeInt8, value)
	}
	if value, ok := mlbu.mutation.AddedMinMissionCount(); ok {
		_spec.AddField(missionloadbalance.FieldMinMissionCount, field.TypeInt8, value)
	}
	if value, ok := mlbu.mutation.MissionBatchID(); ok {
		_spec.SetField(missionloadbalance.FieldMissionBatchID, field.TypeInt64, value)
	}
	if value, ok := mlbu.mutation.AddedMissionBatchID(); ok {
		_spec.AddField(missionloadbalance.FieldMissionBatchID, field.TypeInt64, value)
	}
	if value, ok := mlbu.mutation.MissionBatchNumber(); ok {
		_spec.SetField(missionloadbalance.FieldMissionBatchNumber, field.TypeString, value)
	}
	_spec.AddModifiers(mlbu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mlbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{missionloadbalance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mlbu.mutation.done = true
	return n, nil
}

// MissionLoadBalanceUpdateOne is the builder for updating a single MissionLoadBalance entity.
type MissionLoadBalanceUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MissionLoadBalanceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (mlbuo *MissionLoadBalanceUpdateOne) SetCreatedBy(i int64) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.ResetCreatedBy()
	mlbuo.mutation.SetCreatedBy(i)
	return mlbuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mlbuo *MissionLoadBalanceUpdateOne) SetNillableCreatedBy(i *int64) *MissionLoadBalanceUpdateOne {
	if i != nil {
		mlbuo.SetCreatedBy(*i)
	}
	return mlbuo
}

// AddCreatedBy adds i to the "created_by" field.
func (mlbuo *MissionLoadBalanceUpdateOne) AddCreatedBy(i int64) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.AddCreatedBy(i)
	return mlbuo
}

// SetUpdatedBy sets the "updated_by" field.
func (mlbuo *MissionLoadBalanceUpdateOne) SetUpdatedBy(i int64) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.ResetUpdatedBy()
	mlbuo.mutation.SetUpdatedBy(i)
	return mlbuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mlbuo *MissionLoadBalanceUpdateOne) SetNillableUpdatedBy(i *int64) *MissionLoadBalanceUpdateOne {
	if i != nil {
		mlbuo.SetUpdatedBy(*i)
	}
	return mlbuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (mlbuo *MissionLoadBalanceUpdateOne) AddUpdatedBy(i int64) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.AddUpdatedBy(i)
	return mlbuo
}

// SetUpdatedAt sets the "updated_at" field.
func (mlbuo *MissionLoadBalanceUpdateOne) SetUpdatedAt(t time.Time) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.SetUpdatedAt(t)
	return mlbuo
}

// SetDeletedAt sets the "deleted_at" field.
func (mlbuo *MissionLoadBalanceUpdateOne) SetDeletedAt(t time.Time) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.SetDeletedAt(t)
	return mlbuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mlbuo *MissionLoadBalanceUpdateOne) SetNillableDeletedAt(t *time.Time) *MissionLoadBalanceUpdateOne {
	if t != nil {
		mlbuo.SetDeletedAt(*t)
	}
	return mlbuo
}

// SetMissionType sets the "mission_type" field.
func (mlbuo *MissionLoadBalanceUpdateOne) SetMissionType(et enums.MissionType) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.SetMissionType(et)
	return mlbuo
}

// SetNillableMissionType sets the "mission_type" field if the given value is not nil.
func (mlbuo *MissionLoadBalanceUpdateOne) SetNillableMissionType(et *enums.MissionType) *MissionLoadBalanceUpdateOne {
	if et != nil {
		mlbuo.SetMissionType(*et)
	}
	return mlbuo
}

// SetUserID sets the "user_id" field.
func (mlbuo *MissionLoadBalanceUpdateOne) SetUserID(i int64) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.ResetUserID()
	mlbuo.mutation.SetUserID(i)
	return mlbuo
}

// AddUserID adds i to the "user_id" field.
func (mlbuo *MissionLoadBalanceUpdateOne) AddUserID(i int64) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.AddUserID(i)
	return mlbuo
}

// SetStartedAt sets the "started_at" field.
func (mlbuo *MissionLoadBalanceUpdateOne) SetStartedAt(t time.Time) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.SetStartedAt(t)
	return mlbuo
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (mlbuo *MissionLoadBalanceUpdateOne) SetNillableStartedAt(t *time.Time) *MissionLoadBalanceUpdateOne {
	if t != nil {
		mlbuo.SetStartedAt(*t)
	}
	return mlbuo
}

// SetFinishedAt sets the "finished_at" field.
func (mlbuo *MissionLoadBalanceUpdateOne) SetFinishedAt(t time.Time) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.SetFinishedAt(t)
	return mlbuo
}

// SetNillableFinishedAt sets the "finished_at" field if the given value is not nil.
func (mlbuo *MissionLoadBalanceUpdateOne) SetNillableFinishedAt(t *time.Time) *MissionLoadBalanceUpdateOne {
	if t != nil {
		mlbuo.SetFinishedAt(*t)
	}
	return mlbuo
}

// SetGpuVersion sets the "gpu_version" field.
func (mlbuo *MissionLoadBalanceUpdateOne) SetGpuVersion(ev enums.GpuVersion) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.SetGpuVersion(ev)
	return mlbuo
}

// SetNillableGpuVersion sets the "gpu_version" field if the given value is not nil.
func (mlbuo *MissionLoadBalanceUpdateOne) SetNillableGpuVersion(ev *enums.GpuVersion) *MissionLoadBalanceUpdateOne {
	if ev != nil {
		mlbuo.SetGpuVersion(*ev)
	}
	return mlbuo
}

// SetGpuNum sets the "gpu_num" field.
func (mlbuo *MissionLoadBalanceUpdateOne) SetGpuNum(i int8) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.ResetGpuNum()
	mlbuo.mutation.SetGpuNum(i)
	return mlbuo
}

// SetNillableGpuNum sets the "gpu_num" field if the given value is not nil.
func (mlbuo *MissionLoadBalanceUpdateOne) SetNillableGpuNum(i *int8) *MissionLoadBalanceUpdateOne {
	if i != nil {
		mlbuo.SetGpuNum(*i)
	}
	return mlbuo
}

// AddGpuNum adds i to the "gpu_num" field.
func (mlbuo *MissionLoadBalanceUpdateOne) AddGpuNum(i int8) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.AddGpuNum(i)
	return mlbuo
}

// SetMaxMissionCount sets the "max_mission_count" field.
func (mlbuo *MissionLoadBalanceUpdateOne) SetMaxMissionCount(i int8) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.ResetMaxMissionCount()
	mlbuo.mutation.SetMaxMissionCount(i)
	return mlbuo
}

// SetNillableMaxMissionCount sets the "max_mission_count" field if the given value is not nil.
func (mlbuo *MissionLoadBalanceUpdateOne) SetNillableMaxMissionCount(i *int8) *MissionLoadBalanceUpdateOne {
	if i != nil {
		mlbuo.SetMaxMissionCount(*i)
	}
	return mlbuo
}

// AddMaxMissionCount adds i to the "max_mission_count" field.
func (mlbuo *MissionLoadBalanceUpdateOne) AddMaxMissionCount(i int8) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.AddMaxMissionCount(i)
	return mlbuo
}

// SetMinMissionCount sets the "min_mission_count" field.
func (mlbuo *MissionLoadBalanceUpdateOne) SetMinMissionCount(i int8) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.ResetMinMissionCount()
	mlbuo.mutation.SetMinMissionCount(i)
	return mlbuo
}

// SetNillableMinMissionCount sets the "min_mission_count" field if the given value is not nil.
func (mlbuo *MissionLoadBalanceUpdateOne) SetNillableMinMissionCount(i *int8) *MissionLoadBalanceUpdateOne {
	if i != nil {
		mlbuo.SetMinMissionCount(*i)
	}
	return mlbuo
}

// AddMinMissionCount adds i to the "min_mission_count" field.
func (mlbuo *MissionLoadBalanceUpdateOne) AddMinMissionCount(i int8) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.AddMinMissionCount(i)
	return mlbuo
}

// SetMissionBatchID sets the "mission_batch_id" field.
func (mlbuo *MissionLoadBalanceUpdateOne) SetMissionBatchID(i int64) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.ResetMissionBatchID()
	mlbuo.mutation.SetMissionBatchID(i)
	return mlbuo
}

// SetNillableMissionBatchID sets the "mission_batch_id" field if the given value is not nil.
func (mlbuo *MissionLoadBalanceUpdateOne) SetNillableMissionBatchID(i *int64) *MissionLoadBalanceUpdateOne {
	if i != nil {
		mlbuo.SetMissionBatchID(*i)
	}
	return mlbuo
}

// AddMissionBatchID adds i to the "mission_batch_id" field.
func (mlbuo *MissionLoadBalanceUpdateOne) AddMissionBatchID(i int64) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.AddMissionBatchID(i)
	return mlbuo
}

// SetMissionBatchNumber sets the "mission_batch_number" field.
func (mlbuo *MissionLoadBalanceUpdateOne) SetMissionBatchNumber(s string) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.SetMissionBatchNumber(s)
	return mlbuo
}

// SetNillableMissionBatchNumber sets the "mission_batch_number" field if the given value is not nil.
func (mlbuo *MissionLoadBalanceUpdateOne) SetNillableMissionBatchNumber(s *string) *MissionLoadBalanceUpdateOne {
	if s != nil {
		mlbuo.SetMissionBatchNumber(*s)
	}
	return mlbuo
}

// Mutation returns the MissionLoadBalanceMutation object of the builder.
func (mlbuo *MissionLoadBalanceUpdateOne) Mutation() *MissionLoadBalanceMutation {
	return mlbuo.mutation
}

// Where appends a list predicates to the MissionLoadBalanceUpdate builder.
func (mlbuo *MissionLoadBalanceUpdateOne) Where(ps ...predicate.MissionLoadBalance) *MissionLoadBalanceUpdateOne {
	mlbuo.mutation.Where(ps...)
	return mlbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mlbuo *MissionLoadBalanceUpdateOne) Select(field string, fields ...string) *MissionLoadBalanceUpdateOne {
	mlbuo.fields = append([]string{field}, fields...)
	return mlbuo
}

// Save executes the query and returns the updated MissionLoadBalance entity.
func (mlbuo *MissionLoadBalanceUpdateOne) Save(ctx context.Context) (*MissionLoadBalance, error) {
	mlbuo.defaults()
	return withHooks(ctx, mlbuo.sqlSave, mlbuo.mutation, mlbuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mlbuo *MissionLoadBalanceUpdateOne) SaveX(ctx context.Context) *MissionLoadBalance {
	node, err := mlbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mlbuo *MissionLoadBalanceUpdateOne) Exec(ctx context.Context) error {
	_, err := mlbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mlbuo *MissionLoadBalanceUpdateOne) ExecX(ctx context.Context) {
	if err := mlbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mlbuo *MissionLoadBalanceUpdateOne) defaults() {
	if _, ok := mlbuo.mutation.UpdatedAt(); !ok {
		v := missionloadbalance.UpdateDefaultUpdatedAt()
		mlbuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mlbuo *MissionLoadBalanceUpdateOne) check() error {
	if v, ok := mlbuo.mutation.MissionType(); ok {
		if err := missionloadbalance.MissionTypeValidator(v); err != nil {
			return &ValidationError{Name: "mission_type", err: fmt.Errorf(`cep_ent: validator failed for field "MissionLoadBalance.mission_type": %w`, err)}
		}
	}
	if v, ok := mlbuo.mutation.GpuVersion(); ok {
		if err := missionloadbalance.GpuVersionValidator(v); err != nil {
			return &ValidationError{Name: "gpu_version", err: fmt.Errorf(`cep_ent: validator failed for field "MissionLoadBalance.gpu_version": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mlbuo *MissionLoadBalanceUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MissionLoadBalanceUpdateOne {
	mlbuo.modifiers = append(mlbuo.modifiers, modifiers...)
	return mlbuo
}

func (mlbuo *MissionLoadBalanceUpdateOne) sqlSave(ctx context.Context) (_node *MissionLoadBalance, err error) {
	if err := mlbuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(missionloadbalance.Table, missionloadbalance.Columns, sqlgraph.NewFieldSpec(missionloadbalance.FieldID, field.TypeInt64))
	id, ok := mlbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "MissionLoadBalance.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mlbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, missionloadbalance.FieldID)
		for _, f := range fields {
			if !missionloadbalance.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != missionloadbalance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mlbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mlbuo.mutation.CreatedBy(); ok {
		_spec.SetField(missionloadbalance.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mlbuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(missionloadbalance.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mlbuo.mutation.UpdatedBy(); ok {
		_spec.SetField(missionloadbalance.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mlbuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(missionloadbalance.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mlbuo.mutation.UpdatedAt(); ok {
		_spec.SetField(missionloadbalance.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mlbuo.mutation.DeletedAt(); ok {
		_spec.SetField(missionloadbalance.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := mlbuo.mutation.MissionType(); ok {
		_spec.SetField(missionloadbalance.FieldMissionType, field.TypeEnum, value)
	}
	if value, ok := mlbuo.mutation.UserID(); ok {
		_spec.SetField(missionloadbalance.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := mlbuo.mutation.AddedUserID(); ok {
		_spec.AddField(missionloadbalance.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := mlbuo.mutation.StartedAt(); ok {
		_spec.SetField(missionloadbalance.FieldStartedAt, field.TypeTime, value)
	}
	if value, ok := mlbuo.mutation.FinishedAt(); ok {
		_spec.SetField(missionloadbalance.FieldFinishedAt, field.TypeTime, value)
	}
	if value, ok := mlbuo.mutation.GpuVersion(); ok {
		_spec.SetField(missionloadbalance.FieldGpuVersion, field.TypeEnum, value)
	}
	if value, ok := mlbuo.mutation.GpuNum(); ok {
		_spec.SetField(missionloadbalance.FieldGpuNum, field.TypeInt8, value)
	}
	if value, ok := mlbuo.mutation.AddedGpuNum(); ok {
		_spec.AddField(missionloadbalance.FieldGpuNum, field.TypeInt8, value)
	}
	if value, ok := mlbuo.mutation.MaxMissionCount(); ok {
		_spec.SetField(missionloadbalance.FieldMaxMissionCount, field.TypeInt8, value)
	}
	if value, ok := mlbuo.mutation.AddedMaxMissionCount(); ok {
		_spec.AddField(missionloadbalance.FieldMaxMissionCount, field.TypeInt8, value)
	}
	if value, ok := mlbuo.mutation.MinMissionCount(); ok {
		_spec.SetField(missionloadbalance.FieldMinMissionCount, field.TypeInt8, value)
	}
	if value, ok := mlbuo.mutation.AddedMinMissionCount(); ok {
		_spec.AddField(missionloadbalance.FieldMinMissionCount, field.TypeInt8, value)
	}
	if value, ok := mlbuo.mutation.MissionBatchID(); ok {
		_spec.SetField(missionloadbalance.FieldMissionBatchID, field.TypeInt64, value)
	}
	if value, ok := mlbuo.mutation.AddedMissionBatchID(); ok {
		_spec.AddField(missionloadbalance.FieldMissionBatchID, field.TypeInt64, value)
	}
	if value, ok := mlbuo.mutation.MissionBatchNumber(); ok {
		_spec.SetField(missionloadbalance.FieldMissionBatchNumber, field.TypeString, value)
	}
	_spec.AddModifiers(mlbuo.modifiers...)
	_node = &MissionLoadBalance{config: mlbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mlbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{missionloadbalance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mlbuo.mutation.done = true
	return _node, nil
}
