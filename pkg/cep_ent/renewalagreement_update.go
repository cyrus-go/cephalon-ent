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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/mission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/renewalagreement"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// RenewalAgreementUpdate is the builder for updating RenewalAgreement entities.
type RenewalAgreementUpdate struct {
	config
	hooks    []Hook
	mutation *RenewalAgreementMutation
}

// Where appends a list predicates to the RenewalAgreementUpdate builder.
func (rau *RenewalAgreementUpdate) Where(ps ...predicate.RenewalAgreement) *RenewalAgreementUpdate {
	rau.mutation.Where(ps...)
	return rau
}

// SetCreatedBy sets the "created_by" field.
func (rau *RenewalAgreementUpdate) SetCreatedBy(i int64) *RenewalAgreementUpdate {
	rau.mutation.ResetCreatedBy()
	rau.mutation.SetCreatedBy(i)
	return rau
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (rau *RenewalAgreementUpdate) SetNillableCreatedBy(i *int64) *RenewalAgreementUpdate {
	if i != nil {
		rau.SetCreatedBy(*i)
	}
	return rau
}

// AddCreatedBy adds i to the "created_by" field.
func (rau *RenewalAgreementUpdate) AddCreatedBy(i int64) *RenewalAgreementUpdate {
	rau.mutation.AddCreatedBy(i)
	return rau
}

// SetUpdatedBy sets the "updated_by" field.
func (rau *RenewalAgreementUpdate) SetUpdatedBy(i int64) *RenewalAgreementUpdate {
	rau.mutation.ResetUpdatedBy()
	rau.mutation.SetUpdatedBy(i)
	return rau
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (rau *RenewalAgreementUpdate) SetNillableUpdatedBy(i *int64) *RenewalAgreementUpdate {
	if i != nil {
		rau.SetUpdatedBy(*i)
	}
	return rau
}

// AddUpdatedBy adds i to the "updated_by" field.
func (rau *RenewalAgreementUpdate) AddUpdatedBy(i int64) *RenewalAgreementUpdate {
	rau.mutation.AddUpdatedBy(i)
	return rau
}

// SetUpdatedAt sets the "updated_at" field.
func (rau *RenewalAgreementUpdate) SetUpdatedAt(t time.Time) *RenewalAgreementUpdate {
	rau.mutation.SetUpdatedAt(t)
	return rau
}

// SetDeletedAt sets the "deleted_at" field.
func (rau *RenewalAgreementUpdate) SetDeletedAt(t time.Time) *RenewalAgreementUpdate {
	rau.mutation.SetDeletedAt(t)
	return rau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rau *RenewalAgreementUpdate) SetNillableDeletedAt(t *time.Time) *RenewalAgreementUpdate {
	if t != nil {
		rau.SetDeletedAt(*t)
	}
	return rau
}

// SetNextPayTime sets the "next_pay_time" field.
func (rau *RenewalAgreementUpdate) SetNextPayTime(t time.Time) *RenewalAgreementUpdate {
	rau.mutation.SetNextPayTime(t)
	return rau
}

// SetNillableNextPayTime sets the "next_pay_time" field if the given value is not nil.
func (rau *RenewalAgreementUpdate) SetNillableNextPayTime(t *time.Time) *RenewalAgreementUpdate {
	if t != nil {
		rau.SetNextPayTime(*t)
	}
	return rau
}

// SetType sets the "type" field.
func (rau *RenewalAgreementUpdate) SetType(et enums.RenewalType) *RenewalAgreementUpdate {
	rau.mutation.SetType(et)
	return rau
}

// SetNillableType sets the "type" field if the given value is not nil.
func (rau *RenewalAgreementUpdate) SetNillableType(et *enums.RenewalType) *RenewalAgreementUpdate {
	if et != nil {
		rau.SetType(*et)
	}
	return rau
}

// SetSubStatus sets the "sub_status" field.
func (rau *RenewalAgreementUpdate) SetSubStatus(ess enums.RenewalSubStatus) *RenewalAgreementUpdate {
	rau.mutation.SetSubStatus(ess)
	return rau
}

// SetNillableSubStatus sets the "sub_status" field if the given value is not nil.
func (rau *RenewalAgreementUpdate) SetNillableSubStatus(ess *enums.RenewalSubStatus) *RenewalAgreementUpdate {
	if ess != nil {
		rau.SetSubStatus(*ess)
	}
	return rau
}

// SetPayStatus sets the "pay_status" field.
func (rau *RenewalAgreementUpdate) SetPayStatus(eps enums.RenewalPayStatus) *RenewalAgreementUpdate {
	rau.mutation.SetPayStatus(eps)
	return rau
}

// SetNillablePayStatus sets the "pay_status" field if the given value is not nil.
func (rau *RenewalAgreementUpdate) SetNillablePayStatus(eps *enums.RenewalPayStatus) *RenewalAgreementUpdate {
	if eps != nil {
		rau.SetPayStatus(*eps)
	}
	return rau
}

// SetSymbolID sets the "symbol_id" field.
func (rau *RenewalAgreementUpdate) SetSymbolID(i int64) *RenewalAgreementUpdate {
	rau.mutation.ResetSymbolID()
	rau.mutation.SetSymbolID(i)
	return rau
}

// SetNillableSymbolID sets the "symbol_id" field if the given value is not nil.
func (rau *RenewalAgreementUpdate) SetNillableSymbolID(i *int64) *RenewalAgreementUpdate {
	if i != nil {
		rau.SetSymbolID(*i)
	}
	return rau
}

// AddSymbolID adds i to the "symbol_id" field.
func (rau *RenewalAgreementUpdate) AddSymbolID(i int64) *RenewalAgreementUpdate {
	rau.mutation.AddSymbolID(i)
	return rau
}

// SetFirstPay sets the "first_pay" field.
func (rau *RenewalAgreementUpdate) SetFirstPay(i int64) *RenewalAgreementUpdate {
	rau.mutation.ResetFirstPay()
	rau.mutation.SetFirstPay(i)
	return rau
}

// SetNillableFirstPay sets the "first_pay" field if the given value is not nil.
func (rau *RenewalAgreementUpdate) SetNillableFirstPay(i *int64) *RenewalAgreementUpdate {
	if i != nil {
		rau.SetFirstPay(*i)
	}
	return rau
}

// AddFirstPay adds i to the "first_pay" field.
func (rau *RenewalAgreementUpdate) AddFirstPay(i int64) *RenewalAgreementUpdate {
	rau.mutation.AddFirstPay(i)
	return rau
}

// SetAfterPay sets the "after_pay" field.
func (rau *RenewalAgreementUpdate) SetAfterPay(i int64) *RenewalAgreementUpdate {
	rau.mutation.ResetAfterPay()
	rau.mutation.SetAfterPay(i)
	return rau
}

// SetNillableAfterPay sets the "after_pay" field if the given value is not nil.
func (rau *RenewalAgreementUpdate) SetNillableAfterPay(i *int64) *RenewalAgreementUpdate {
	if i != nil {
		rau.SetAfterPay(*i)
	}
	return rau
}

// AddAfterPay adds i to the "after_pay" field.
func (rau *RenewalAgreementUpdate) AddAfterPay(i int64) *RenewalAgreementUpdate {
	rau.mutation.AddAfterPay(i)
	return rau
}

// SetLastWarningTime sets the "last_warning_time" field.
func (rau *RenewalAgreementUpdate) SetLastWarningTime(t time.Time) *RenewalAgreementUpdate {
	rau.mutation.SetLastWarningTime(t)
	return rau
}

// SetNillableLastWarningTime sets the "last_warning_time" field if the given value is not nil.
func (rau *RenewalAgreementUpdate) SetNillableLastWarningTime(t *time.Time) *RenewalAgreementUpdate {
	if t != nil {
		rau.SetLastWarningTime(*t)
	}
	return rau
}

// SetSubFinishedTime sets the "sub_finished_time" field.
func (rau *RenewalAgreementUpdate) SetSubFinishedTime(t time.Time) *RenewalAgreementUpdate {
	rau.mutation.SetSubFinishedTime(t)
	return rau
}

// SetNillableSubFinishedTime sets the "sub_finished_time" field if the given value is not nil.
func (rau *RenewalAgreementUpdate) SetNillableSubFinishedTime(t *time.Time) *RenewalAgreementUpdate {
	if t != nil {
		rau.SetSubFinishedTime(*t)
	}
	return rau
}

// SetUserID sets the "user_id" field.
func (rau *RenewalAgreementUpdate) SetUserID(i int64) *RenewalAgreementUpdate {
	rau.mutation.SetUserID(i)
	return rau
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (rau *RenewalAgreementUpdate) SetNillableUserID(i *int64) *RenewalAgreementUpdate {
	if i != nil {
		rau.SetUserID(*i)
	}
	return rau
}

// SetMissionID sets the "mission_id" field.
func (rau *RenewalAgreementUpdate) SetMissionID(i int64) *RenewalAgreementUpdate {
	rau.mutation.SetMissionID(i)
	return rau
}

// SetNillableMissionID sets the "mission_id" field if the given value is not nil.
func (rau *RenewalAgreementUpdate) SetNillableMissionID(i *int64) *RenewalAgreementUpdate {
	if i != nil {
		rau.SetMissionID(*i)
	}
	return rau
}

// SetUser sets the "user" edge to the User entity.
func (rau *RenewalAgreementUpdate) SetUser(u *User) *RenewalAgreementUpdate {
	return rau.SetUserID(u.ID)
}

// SetMission sets the "mission" edge to the Mission entity.
func (rau *RenewalAgreementUpdate) SetMission(m *Mission) *RenewalAgreementUpdate {
	return rau.SetMissionID(m.ID)
}

// Mutation returns the RenewalAgreementMutation object of the builder.
func (rau *RenewalAgreementUpdate) Mutation() *RenewalAgreementMutation {
	return rau.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (rau *RenewalAgreementUpdate) ClearUser() *RenewalAgreementUpdate {
	rau.mutation.ClearUser()
	return rau
}

// ClearMission clears the "mission" edge to the Mission entity.
func (rau *RenewalAgreementUpdate) ClearMission() *RenewalAgreementUpdate {
	rau.mutation.ClearMission()
	return rau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rau *RenewalAgreementUpdate) Save(ctx context.Context) (int, error) {
	rau.defaults()
	return withHooks(ctx, rau.sqlSave, rau.mutation, rau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rau *RenewalAgreementUpdate) SaveX(ctx context.Context) int {
	affected, err := rau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rau *RenewalAgreementUpdate) Exec(ctx context.Context) error {
	_, err := rau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rau *RenewalAgreementUpdate) ExecX(ctx context.Context) {
	if err := rau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rau *RenewalAgreementUpdate) defaults() {
	if _, ok := rau.mutation.UpdatedAt(); !ok {
		v := renewalagreement.UpdateDefaultUpdatedAt()
		rau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rau *RenewalAgreementUpdate) check() error {
	if v, ok := rau.mutation.GetType(); ok {
		if err := renewalagreement.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`cep_ent: validator failed for field "RenewalAgreement.type": %w`, err)}
		}
	}
	if v, ok := rau.mutation.SubStatus(); ok {
		if err := renewalagreement.SubStatusValidator(v); err != nil {
			return &ValidationError{Name: "sub_status", err: fmt.Errorf(`cep_ent: validator failed for field "RenewalAgreement.sub_status": %w`, err)}
		}
	}
	if v, ok := rau.mutation.PayStatus(); ok {
		if err := renewalagreement.PayStatusValidator(v); err != nil {
			return &ValidationError{Name: "pay_status", err: fmt.Errorf(`cep_ent: validator failed for field "RenewalAgreement.pay_status": %w`, err)}
		}
	}
	if _, ok := rau.mutation.UserID(); rau.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "RenewalAgreement.user"`)
	}
	if _, ok := rau.mutation.MissionID(); rau.mutation.MissionCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "RenewalAgreement.mission"`)
	}
	return nil
}

func (rau *RenewalAgreementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(renewalagreement.Table, renewalagreement.Columns, sqlgraph.NewFieldSpec(renewalagreement.FieldID, field.TypeInt64))
	if ps := rau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rau.mutation.CreatedBy(); ok {
		_spec.SetField(renewalagreement.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := rau.mutation.AddedCreatedBy(); ok {
		_spec.AddField(renewalagreement.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := rau.mutation.UpdatedBy(); ok {
		_spec.SetField(renewalagreement.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := rau.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(renewalagreement.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := rau.mutation.UpdatedAt(); ok {
		_spec.SetField(renewalagreement.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rau.mutation.DeletedAt(); ok {
		_spec.SetField(renewalagreement.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := rau.mutation.NextPayTime(); ok {
		_spec.SetField(renewalagreement.FieldNextPayTime, field.TypeTime, value)
	}
	if value, ok := rau.mutation.GetType(); ok {
		_spec.SetField(renewalagreement.FieldType, field.TypeEnum, value)
	}
	if value, ok := rau.mutation.SubStatus(); ok {
		_spec.SetField(renewalagreement.FieldSubStatus, field.TypeEnum, value)
	}
	if value, ok := rau.mutation.PayStatus(); ok {
		_spec.SetField(renewalagreement.FieldPayStatus, field.TypeEnum, value)
	}
	if value, ok := rau.mutation.SymbolID(); ok {
		_spec.SetField(renewalagreement.FieldSymbolID, field.TypeInt64, value)
	}
	if value, ok := rau.mutation.AddedSymbolID(); ok {
		_spec.AddField(renewalagreement.FieldSymbolID, field.TypeInt64, value)
	}
	if value, ok := rau.mutation.FirstPay(); ok {
		_spec.SetField(renewalagreement.FieldFirstPay, field.TypeInt64, value)
	}
	if value, ok := rau.mutation.AddedFirstPay(); ok {
		_spec.AddField(renewalagreement.FieldFirstPay, field.TypeInt64, value)
	}
	if value, ok := rau.mutation.AfterPay(); ok {
		_spec.SetField(renewalagreement.FieldAfterPay, field.TypeInt64, value)
	}
	if value, ok := rau.mutation.AddedAfterPay(); ok {
		_spec.AddField(renewalagreement.FieldAfterPay, field.TypeInt64, value)
	}
	if value, ok := rau.mutation.LastWarningTime(); ok {
		_spec.SetField(renewalagreement.FieldLastWarningTime, field.TypeTime, value)
	}
	if value, ok := rau.mutation.SubFinishedTime(); ok {
		_spec.SetField(renewalagreement.FieldSubFinishedTime, field.TypeTime, value)
	}
	if rau.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   renewalagreement.UserTable,
			Columns: []string{renewalagreement.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rau.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   renewalagreement.UserTable,
			Columns: []string{renewalagreement.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rau.mutation.MissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   renewalagreement.MissionTable,
			Columns: []string{renewalagreement.MissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mission.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rau.mutation.MissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   renewalagreement.MissionTable,
			Columns: []string{renewalagreement.MissionColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, rau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{renewalagreement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rau.mutation.done = true
	return n, nil
}

// RenewalAgreementUpdateOne is the builder for updating a single RenewalAgreement entity.
type RenewalAgreementUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RenewalAgreementMutation
}

// SetCreatedBy sets the "created_by" field.
func (rauo *RenewalAgreementUpdateOne) SetCreatedBy(i int64) *RenewalAgreementUpdateOne {
	rauo.mutation.ResetCreatedBy()
	rauo.mutation.SetCreatedBy(i)
	return rauo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (rauo *RenewalAgreementUpdateOne) SetNillableCreatedBy(i *int64) *RenewalAgreementUpdateOne {
	if i != nil {
		rauo.SetCreatedBy(*i)
	}
	return rauo
}

// AddCreatedBy adds i to the "created_by" field.
func (rauo *RenewalAgreementUpdateOne) AddCreatedBy(i int64) *RenewalAgreementUpdateOne {
	rauo.mutation.AddCreatedBy(i)
	return rauo
}

// SetUpdatedBy sets the "updated_by" field.
func (rauo *RenewalAgreementUpdateOne) SetUpdatedBy(i int64) *RenewalAgreementUpdateOne {
	rauo.mutation.ResetUpdatedBy()
	rauo.mutation.SetUpdatedBy(i)
	return rauo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (rauo *RenewalAgreementUpdateOne) SetNillableUpdatedBy(i *int64) *RenewalAgreementUpdateOne {
	if i != nil {
		rauo.SetUpdatedBy(*i)
	}
	return rauo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (rauo *RenewalAgreementUpdateOne) AddUpdatedBy(i int64) *RenewalAgreementUpdateOne {
	rauo.mutation.AddUpdatedBy(i)
	return rauo
}

// SetUpdatedAt sets the "updated_at" field.
func (rauo *RenewalAgreementUpdateOne) SetUpdatedAt(t time.Time) *RenewalAgreementUpdateOne {
	rauo.mutation.SetUpdatedAt(t)
	return rauo
}

// SetDeletedAt sets the "deleted_at" field.
func (rauo *RenewalAgreementUpdateOne) SetDeletedAt(t time.Time) *RenewalAgreementUpdateOne {
	rauo.mutation.SetDeletedAt(t)
	return rauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rauo *RenewalAgreementUpdateOne) SetNillableDeletedAt(t *time.Time) *RenewalAgreementUpdateOne {
	if t != nil {
		rauo.SetDeletedAt(*t)
	}
	return rauo
}

// SetNextPayTime sets the "next_pay_time" field.
func (rauo *RenewalAgreementUpdateOne) SetNextPayTime(t time.Time) *RenewalAgreementUpdateOne {
	rauo.mutation.SetNextPayTime(t)
	return rauo
}

// SetNillableNextPayTime sets the "next_pay_time" field if the given value is not nil.
func (rauo *RenewalAgreementUpdateOne) SetNillableNextPayTime(t *time.Time) *RenewalAgreementUpdateOne {
	if t != nil {
		rauo.SetNextPayTime(*t)
	}
	return rauo
}

// SetType sets the "type" field.
func (rauo *RenewalAgreementUpdateOne) SetType(et enums.RenewalType) *RenewalAgreementUpdateOne {
	rauo.mutation.SetType(et)
	return rauo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (rauo *RenewalAgreementUpdateOne) SetNillableType(et *enums.RenewalType) *RenewalAgreementUpdateOne {
	if et != nil {
		rauo.SetType(*et)
	}
	return rauo
}

// SetSubStatus sets the "sub_status" field.
func (rauo *RenewalAgreementUpdateOne) SetSubStatus(ess enums.RenewalSubStatus) *RenewalAgreementUpdateOne {
	rauo.mutation.SetSubStatus(ess)
	return rauo
}

// SetNillableSubStatus sets the "sub_status" field if the given value is not nil.
func (rauo *RenewalAgreementUpdateOne) SetNillableSubStatus(ess *enums.RenewalSubStatus) *RenewalAgreementUpdateOne {
	if ess != nil {
		rauo.SetSubStatus(*ess)
	}
	return rauo
}

// SetPayStatus sets the "pay_status" field.
func (rauo *RenewalAgreementUpdateOne) SetPayStatus(eps enums.RenewalPayStatus) *RenewalAgreementUpdateOne {
	rauo.mutation.SetPayStatus(eps)
	return rauo
}

// SetNillablePayStatus sets the "pay_status" field if the given value is not nil.
func (rauo *RenewalAgreementUpdateOne) SetNillablePayStatus(eps *enums.RenewalPayStatus) *RenewalAgreementUpdateOne {
	if eps != nil {
		rauo.SetPayStatus(*eps)
	}
	return rauo
}

// SetSymbolID sets the "symbol_id" field.
func (rauo *RenewalAgreementUpdateOne) SetSymbolID(i int64) *RenewalAgreementUpdateOne {
	rauo.mutation.ResetSymbolID()
	rauo.mutation.SetSymbolID(i)
	return rauo
}

// SetNillableSymbolID sets the "symbol_id" field if the given value is not nil.
func (rauo *RenewalAgreementUpdateOne) SetNillableSymbolID(i *int64) *RenewalAgreementUpdateOne {
	if i != nil {
		rauo.SetSymbolID(*i)
	}
	return rauo
}

// AddSymbolID adds i to the "symbol_id" field.
func (rauo *RenewalAgreementUpdateOne) AddSymbolID(i int64) *RenewalAgreementUpdateOne {
	rauo.mutation.AddSymbolID(i)
	return rauo
}

// SetFirstPay sets the "first_pay" field.
func (rauo *RenewalAgreementUpdateOne) SetFirstPay(i int64) *RenewalAgreementUpdateOne {
	rauo.mutation.ResetFirstPay()
	rauo.mutation.SetFirstPay(i)
	return rauo
}

// SetNillableFirstPay sets the "first_pay" field if the given value is not nil.
func (rauo *RenewalAgreementUpdateOne) SetNillableFirstPay(i *int64) *RenewalAgreementUpdateOne {
	if i != nil {
		rauo.SetFirstPay(*i)
	}
	return rauo
}

// AddFirstPay adds i to the "first_pay" field.
func (rauo *RenewalAgreementUpdateOne) AddFirstPay(i int64) *RenewalAgreementUpdateOne {
	rauo.mutation.AddFirstPay(i)
	return rauo
}

// SetAfterPay sets the "after_pay" field.
func (rauo *RenewalAgreementUpdateOne) SetAfterPay(i int64) *RenewalAgreementUpdateOne {
	rauo.mutation.ResetAfterPay()
	rauo.mutation.SetAfterPay(i)
	return rauo
}

// SetNillableAfterPay sets the "after_pay" field if the given value is not nil.
func (rauo *RenewalAgreementUpdateOne) SetNillableAfterPay(i *int64) *RenewalAgreementUpdateOne {
	if i != nil {
		rauo.SetAfterPay(*i)
	}
	return rauo
}

// AddAfterPay adds i to the "after_pay" field.
func (rauo *RenewalAgreementUpdateOne) AddAfterPay(i int64) *RenewalAgreementUpdateOne {
	rauo.mutation.AddAfterPay(i)
	return rauo
}

// SetLastWarningTime sets the "last_warning_time" field.
func (rauo *RenewalAgreementUpdateOne) SetLastWarningTime(t time.Time) *RenewalAgreementUpdateOne {
	rauo.mutation.SetLastWarningTime(t)
	return rauo
}

// SetNillableLastWarningTime sets the "last_warning_time" field if the given value is not nil.
func (rauo *RenewalAgreementUpdateOne) SetNillableLastWarningTime(t *time.Time) *RenewalAgreementUpdateOne {
	if t != nil {
		rauo.SetLastWarningTime(*t)
	}
	return rauo
}

// SetSubFinishedTime sets the "sub_finished_time" field.
func (rauo *RenewalAgreementUpdateOne) SetSubFinishedTime(t time.Time) *RenewalAgreementUpdateOne {
	rauo.mutation.SetSubFinishedTime(t)
	return rauo
}

// SetNillableSubFinishedTime sets the "sub_finished_time" field if the given value is not nil.
func (rauo *RenewalAgreementUpdateOne) SetNillableSubFinishedTime(t *time.Time) *RenewalAgreementUpdateOne {
	if t != nil {
		rauo.SetSubFinishedTime(*t)
	}
	return rauo
}

// SetUserID sets the "user_id" field.
func (rauo *RenewalAgreementUpdateOne) SetUserID(i int64) *RenewalAgreementUpdateOne {
	rauo.mutation.SetUserID(i)
	return rauo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (rauo *RenewalAgreementUpdateOne) SetNillableUserID(i *int64) *RenewalAgreementUpdateOne {
	if i != nil {
		rauo.SetUserID(*i)
	}
	return rauo
}

// SetMissionID sets the "mission_id" field.
func (rauo *RenewalAgreementUpdateOne) SetMissionID(i int64) *RenewalAgreementUpdateOne {
	rauo.mutation.SetMissionID(i)
	return rauo
}

// SetNillableMissionID sets the "mission_id" field if the given value is not nil.
func (rauo *RenewalAgreementUpdateOne) SetNillableMissionID(i *int64) *RenewalAgreementUpdateOne {
	if i != nil {
		rauo.SetMissionID(*i)
	}
	return rauo
}

// SetUser sets the "user" edge to the User entity.
func (rauo *RenewalAgreementUpdateOne) SetUser(u *User) *RenewalAgreementUpdateOne {
	return rauo.SetUserID(u.ID)
}

// SetMission sets the "mission" edge to the Mission entity.
func (rauo *RenewalAgreementUpdateOne) SetMission(m *Mission) *RenewalAgreementUpdateOne {
	return rauo.SetMissionID(m.ID)
}

// Mutation returns the RenewalAgreementMutation object of the builder.
func (rauo *RenewalAgreementUpdateOne) Mutation() *RenewalAgreementMutation {
	return rauo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (rauo *RenewalAgreementUpdateOne) ClearUser() *RenewalAgreementUpdateOne {
	rauo.mutation.ClearUser()
	return rauo
}

// ClearMission clears the "mission" edge to the Mission entity.
func (rauo *RenewalAgreementUpdateOne) ClearMission() *RenewalAgreementUpdateOne {
	rauo.mutation.ClearMission()
	return rauo
}

// Where appends a list predicates to the RenewalAgreementUpdate builder.
func (rauo *RenewalAgreementUpdateOne) Where(ps ...predicate.RenewalAgreement) *RenewalAgreementUpdateOne {
	rauo.mutation.Where(ps...)
	return rauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rauo *RenewalAgreementUpdateOne) Select(field string, fields ...string) *RenewalAgreementUpdateOne {
	rauo.fields = append([]string{field}, fields...)
	return rauo
}

// Save executes the query and returns the updated RenewalAgreement entity.
func (rauo *RenewalAgreementUpdateOne) Save(ctx context.Context) (*RenewalAgreement, error) {
	rauo.defaults()
	return withHooks(ctx, rauo.sqlSave, rauo.mutation, rauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rauo *RenewalAgreementUpdateOne) SaveX(ctx context.Context) *RenewalAgreement {
	node, err := rauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rauo *RenewalAgreementUpdateOne) Exec(ctx context.Context) error {
	_, err := rauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rauo *RenewalAgreementUpdateOne) ExecX(ctx context.Context) {
	if err := rauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rauo *RenewalAgreementUpdateOne) defaults() {
	if _, ok := rauo.mutation.UpdatedAt(); !ok {
		v := renewalagreement.UpdateDefaultUpdatedAt()
		rauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rauo *RenewalAgreementUpdateOne) check() error {
	if v, ok := rauo.mutation.GetType(); ok {
		if err := renewalagreement.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`cep_ent: validator failed for field "RenewalAgreement.type": %w`, err)}
		}
	}
	if v, ok := rauo.mutation.SubStatus(); ok {
		if err := renewalagreement.SubStatusValidator(v); err != nil {
			return &ValidationError{Name: "sub_status", err: fmt.Errorf(`cep_ent: validator failed for field "RenewalAgreement.sub_status": %w`, err)}
		}
	}
	if v, ok := rauo.mutation.PayStatus(); ok {
		if err := renewalagreement.PayStatusValidator(v); err != nil {
			return &ValidationError{Name: "pay_status", err: fmt.Errorf(`cep_ent: validator failed for field "RenewalAgreement.pay_status": %w`, err)}
		}
	}
	if _, ok := rauo.mutation.UserID(); rauo.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "RenewalAgreement.user"`)
	}
	if _, ok := rauo.mutation.MissionID(); rauo.mutation.MissionCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "RenewalAgreement.mission"`)
	}
	return nil
}

func (rauo *RenewalAgreementUpdateOne) sqlSave(ctx context.Context) (_node *RenewalAgreement, err error) {
	if err := rauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(renewalagreement.Table, renewalagreement.Columns, sqlgraph.NewFieldSpec(renewalagreement.FieldID, field.TypeInt64))
	id, ok := rauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "RenewalAgreement.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, renewalagreement.FieldID)
		for _, f := range fields {
			if !renewalagreement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != renewalagreement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rauo.mutation.CreatedBy(); ok {
		_spec.SetField(renewalagreement.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := rauo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(renewalagreement.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := rauo.mutation.UpdatedBy(); ok {
		_spec.SetField(renewalagreement.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := rauo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(renewalagreement.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := rauo.mutation.UpdatedAt(); ok {
		_spec.SetField(renewalagreement.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rauo.mutation.DeletedAt(); ok {
		_spec.SetField(renewalagreement.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := rauo.mutation.NextPayTime(); ok {
		_spec.SetField(renewalagreement.FieldNextPayTime, field.TypeTime, value)
	}
	if value, ok := rauo.mutation.GetType(); ok {
		_spec.SetField(renewalagreement.FieldType, field.TypeEnum, value)
	}
	if value, ok := rauo.mutation.SubStatus(); ok {
		_spec.SetField(renewalagreement.FieldSubStatus, field.TypeEnum, value)
	}
	if value, ok := rauo.mutation.PayStatus(); ok {
		_spec.SetField(renewalagreement.FieldPayStatus, field.TypeEnum, value)
	}
	if value, ok := rauo.mutation.SymbolID(); ok {
		_spec.SetField(renewalagreement.FieldSymbolID, field.TypeInt64, value)
	}
	if value, ok := rauo.mutation.AddedSymbolID(); ok {
		_spec.AddField(renewalagreement.FieldSymbolID, field.TypeInt64, value)
	}
	if value, ok := rauo.mutation.FirstPay(); ok {
		_spec.SetField(renewalagreement.FieldFirstPay, field.TypeInt64, value)
	}
	if value, ok := rauo.mutation.AddedFirstPay(); ok {
		_spec.AddField(renewalagreement.FieldFirstPay, field.TypeInt64, value)
	}
	if value, ok := rauo.mutation.AfterPay(); ok {
		_spec.SetField(renewalagreement.FieldAfterPay, field.TypeInt64, value)
	}
	if value, ok := rauo.mutation.AddedAfterPay(); ok {
		_spec.AddField(renewalagreement.FieldAfterPay, field.TypeInt64, value)
	}
	if value, ok := rauo.mutation.LastWarningTime(); ok {
		_spec.SetField(renewalagreement.FieldLastWarningTime, field.TypeTime, value)
	}
	if value, ok := rauo.mutation.SubFinishedTime(); ok {
		_spec.SetField(renewalagreement.FieldSubFinishedTime, field.TypeTime, value)
	}
	if rauo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   renewalagreement.UserTable,
			Columns: []string{renewalagreement.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rauo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   renewalagreement.UserTable,
			Columns: []string{renewalagreement.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rauo.mutation.MissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   renewalagreement.MissionTable,
			Columns: []string{renewalagreement.MissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mission.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rauo.mutation.MissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   renewalagreement.MissionTable,
			Columns: []string{renewalagreement.MissionColumn},
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
	_node = &RenewalAgreement{config: rauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{renewalagreement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rauo.mutation.done = true
	return _node, nil
}
