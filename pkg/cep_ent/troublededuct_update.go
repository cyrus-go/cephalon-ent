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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/device"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/troublededuct"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// TroubleDeductUpdate is the builder for updating TroubleDeduct entities.
type TroubleDeductUpdate struct {
	config
	hooks     []Hook
	mutation  *TroubleDeductMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TroubleDeductUpdate builder.
func (tdu *TroubleDeductUpdate) Where(ps ...predicate.TroubleDeduct) *TroubleDeductUpdate {
	tdu.mutation.Where(ps...)
	return tdu
}

// SetCreatedBy sets the "created_by" field.
func (tdu *TroubleDeductUpdate) SetCreatedBy(i int64) *TroubleDeductUpdate {
	tdu.mutation.ResetCreatedBy()
	tdu.mutation.SetCreatedBy(i)
	return tdu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (tdu *TroubleDeductUpdate) SetNillableCreatedBy(i *int64) *TroubleDeductUpdate {
	if i != nil {
		tdu.SetCreatedBy(*i)
	}
	return tdu
}

// AddCreatedBy adds i to the "created_by" field.
func (tdu *TroubleDeductUpdate) AddCreatedBy(i int64) *TroubleDeductUpdate {
	tdu.mutation.AddCreatedBy(i)
	return tdu
}

// SetUpdatedBy sets the "updated_by" field.
func (tdu *TroubleDeductUpdate) SetUpdatedBy(i int64) *TroubleDeductUpdate {
	tdu.mutation.ResetUpdatedBy()
	tdu.mutation.SetUpdatedBy(i)
	return tdu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tdu *TroubleDeductUpdate) SetNillableUpdatedBy(i *int64) *TroubleDeductUpdate {
	if i != nil {
		tdu.SetUpdatedBy(*i)
	}
	return tdu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (tdu *TroubleDeductUpdate) AddUpdatedBy(i int64) *TroubleDeductUpdate {
	tdu.mutation.AddUpdatedBy(i)
	return tdu
}

// SetUpdatedAt sets the "updated_at" field.
func (tdu *TroubleDeductUpdate) SetUpdatedAt(t time.Time) *TroubleDeductUpdate {
	tdu.mutation.SetUpdatedAt(t)
	return tdu
}

// SetDeletedAt sets the "deleted_at" field.
func (tdu *TroubleDeductUpdate) SetDeletedAt(t time.Time) *TroubleDeductUpdate {
	tdu.mutation.SetDeletedAt(t)
	return tdu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tdu *TroubleDeductUpdate) SetNillableDeletedAt(t *time.Time) *TroubleDeductUpdate {
	if t != nil {
		tdu.SetDeletedAt(*t)
	}
	return tdu
}

// SetDeviceID sets the "device_id" field.
func (tdu *TroubleDeductUpdate) SetDeviceID(i int64) *TroubleDeductUpdate {
	tdu.mutation.SetDeviceID(i)
	return tdu
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (tdu *TroubleDeductUpdate) SetNillableDeviceID(i *int64) *TroubleDeductUpdate {
	if i != nil {
		tdu.SetDeviceID(*i)
	}
	return tdu
}

// SetStartedAt sets the "started_at" field.
func (tdu *TroubleDeductUpdate) SetStartedAt(t time.Time) *TroubleDeductUpdate {
	tdu.mutation.SetStartedAt(t)
	return tdu
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (tdu *TroubleDeductUpdate) SetNillableStartedAt(t *time.Time) *TroubleDeductUpdate {
	if t != nil {
		tdu.SetStartedAt(*t)
	}
	return tdu
}

// SetFinishedAt sets the "finished_at" field.
func (tdu *TroubleDeductUpdate) SetFinishedAt(t time.Time) *TroubleDeductUpdate {
	tdu.mutation.SetFinishedAt(t)
	return tdu
}

// SetNillableFinishedAt sets the "finished_at" field if the given value is not nil.
func (tdu *TroubleDeductUpdate) SetNillableFinishedAt(t *time.Time) *TroubleDeductUpdate {
	if t != nil {
		tdu.SetFinishedAt(*t)
	}
	return tdu
}

// SetTimeOfDuration sets the "time_of_duration" field.
func (tdu *TroubleDeductUpdate) SetTimeOfDuration(f float64) *TroubleDeductUpdate {
	tdu.mutation.ResetTimeOfDuration()
	tdu.mutation.SetTimeOfDuration(f)
	return tdu
}

// SetNillableTimeOfDuration sets the "time_of_duration" field if the given value is not nil.
func (tdu *TroubleDeductUpdate) SetNillableTimeOfDuration(f *float64) *TroubleDeductUpdate {
	if f != nil {
		tdu.SetTimeOfDuration(*f)
	}
	return tdu
}

// AddTimeOfDuration adds f to the "time_of_duration" field.
func (tdu *TroubleDeductUpdate) AddTimeOfDuration(f float64) *TroubleDeductUpdate {
	tdu.mutation.AddTimeOfDuration(f)
	return tdu
}

// SetAmount sets the "amount" field.
func (tdu *TroubleDeductUpdate) SetAmount(i int64) *TroubleDeductUpdate {
	tdu.mutation.ResetAmount()
	tdu.mutation.SetAmount(i)
	return tdu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tdu *TroubleDeductUpdate) SetNillableAmount(i *int64) *TroubleDeductUpdate {
	if i != nil {
		tdu.SetAmount(*i)
	}
	return tdu
}

// AddAmount adds i to the "amount" field.
func (tdu *TroubleDeductUpdate) AddAmount(i int64) *TroubleDeductUpdate {
	tdu.mutation.AddAmount(i)
	return tdu
}

// SetStatus sets the "status" field.
func (tdu *TroubleDeductUpdate) SetStatus(eds enums.TroubleDeductStatus) *TroubleDeductUpdate {
	tdu.mutation.SetStatus(eds)
	return tdu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tdu *TroubleDeductUpdate) SetNillableStatus(eds *enums.TroubleDeductStatus) *TroubleDeductUpdate {
	if eds != nil {
		tdu.SetStatus(*eds)
	}
	return tdu
}

// SetReason sets the "reason" field.
func (tdu *TroubleDeductUpdate) SetReason(s string) *TroubleDeductUpdate {
	tdu.mutation.SetReason(s)
	return tdu
}

// SetNillableReason sets the "reason" field if the given value is not nil.
func (tdu *TroubleDeductUpdate) SetNillableReason(s *string) *TroubleDeductUpdate {
	if s != nil {
		tdu.SetReason(*s)
	}
	return tdu
}

// SetRejectReason sets the "reject_reason" field.
func (tdu *TroubleDeductUpdate) SetRejectReason(s string) *TroubleDeductUpdate {
	tdu.mutation.SetRejectReason(s)
	return tdu
}

// SetNillableRejectReason sets the "reject_reason" field if the given value is not nil.
func (tdu *TroubleDeductUpdate) SetNillableRejectReason(s *string) *TroubleDeductUpdate {
	if s != nil {
		tdu.SetRejectReason(*s)
	}
	return tdu
}

// SetDevice sets the "device" edge to the Device entity.
func (tdu *TroubleDeductUpdate) SetDevice(d *Device) *TroubleDeductUpdate {
	return tdu.SetDeviceID(d.ID)
}

// Mutation returns the TroubleDeductMutation object of the builder.
func (tdu *TroubleDeductUpdate) Mutation() *TroubleDeductMutation {
	return tdu.mutation
}

// ClearDevice clears the "device" edge to the Device entity.
func (tdu *TroubleDeductUpdate) ClearDevice() *TroubleDeductUpdate {
	tdu.mutation.ClearDevice()
	return tdu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tdu *TroubleDeductUpdate) Save(ctx context.Context) (int, error) {
	tdu.defaults()
	return withHooks(ctx, tdu.sqlSave, tdu.mutation, tdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tdu *TroubleDeductUpdate) SaveX(ctx context.Context) int {
	affected, err := tdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tdu *TroubleDeductUpdate) Exec(ctx context.Context) error {
	_, err := tdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdu *TroubleDeductUpdate) ExecX(ctx context.Context) {
	if err := tdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tdu *TroubleDeductUpdate) defaults() {
	if _, ok := tdu.mutation.UpdatedAt(); !ok {
		v := troublededuct.UpdateDefaultUpdatedAt()
		tdu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tdu *TroubleDeductUpdate) check() error {
	if v, ok := tdu.mutation.Status(); ok {
		if err := troublededuct.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`cep_ent: validator failed for field "TroubleDeduct.status": %w`, err)}
		}
	}
	if _, ok := tdu.mutation.DeviceID(); tdu.mutation.DeviceCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "TroubleDeduct.device"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tdu *TroubleDeductUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TroubleDeductUpdate {
	tdu.modifiers = append(tdu.modifiers, modifiers...)
	return tdu
}

func (tdu *TroubleDeductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tdu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(troublededuct.Table, troublededuct.Columns, sqlgraph.NewFieldSpec(troublededuct.FieldID, field.TypeInt64))
	if ps := tdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tdu.mutation.CreatedBy(); ok {
		_spec.SetField(troublededuct.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := tdu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(troublededuct.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := tdu.mutation.UpdatedBy(); ok {
		_spec.SetField(troublededuct.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := tdu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(troublededuct.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := tdu.mutation.UpdatedAt(); ok {
		_spec.SetField(troublededuct.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tdu.mutation.DeletedAt(); ok {
		_spec.SetField(troublededuct.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := tdu.mutation.StartedAt(); ok {
		_spec.SetField(troublededuct.FieldStartedAt, field.TypeTime, value)
	}
	if value, ok := tdu.mutation.FinishedAt(); ok {
		_spec.SetField(troublededuct.FieldFinishedAt, field.TypeTime, value)
	}
	if value, ok := tdu.mutation.TimeOfDuration(); ok {
		_spec.SetField(troublededuct.FieldTimeOfDuration, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.AddedTimeOfDuration(); ok {
		_spec.AddField(troublededuct.FieldTimeOfDuration, field.TypeFloat64, value)
	}
	if value, ok := tdu.mutation.Amount(); ok {
		_spec.SetField(troublededuct.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := tdu.mutation.AddedAmount(); ok {
		_spec.AddField(troublededuct.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := tdu.mutation.Status(); ok {
		_spec.SetField(troublededuct.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tdu.mutation.Reason(); ok {
		_spec.SetField(troublededuct.FieldReason, field.TypeString, value)
	}
	if value, ok := tdu.mutation.RejectReason(); ok {
		_spec.SetField(troublededuct.FieldRejectReason, field.TypeString, value)
	}
	if tdu.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   troublededuct.DeviceTable,
			Columns: []string{troublededuct.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tdu.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   troublededuct.DeviceTable,
			Columns: []string{troublededuct.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(tdu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{troublededuct.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tdu.mutation.done = true
	return n, nil
}

// TroubleDeductUpdateOne is the builder for updating a single TroubleDeduct entity.
type TroubleDeductUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TroubleDeductMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (tduo *TroubleDeductUpdateOne) SetCreatedBy(i int64) *TroubleDeductUpdateOne {
	tduo.mutation.ResetCreatedBy()
	tduo.mutation.SetCreatedBy(i)
	return tduo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (tduo *TroubleDeductUpdateOne) SetNillableCreatedBy(i *int64) *TroubleDeductUpdateOne {
	if i != nil {
		tduo.SetCreatedBy(*i)
	}
	return tduo
}

// AddCreatedBy adds i to the "created_by" field.
func (tduo *TroubleDeductUpdateOne) AddCreatedBy(i int64) *TroubleDeductUpdateOne {
	tduo.mutation.AddCreatedBy(i)
	return tduo
}

// SetUpdatedBy sets the "updated_by" field.
func (tduo *TroubleDeductUpdateOne) SetUpdatedBy(i int64) *TroubleDeductUpdateOne {
	tduo.mutation.ResetUpdatedBy()
	tduo.mutation.SetUpdatedBy(i)
	return tduo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tduo *TroubleDeductUpdateOne) SetNillableUpdatedBy(i *int64) *TroubleDeductUpdateOne {
	if i != nil {
		tduo.SetUpdatedBy(*i)
	}
	return tduo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (tduo *TroubleDeductUpdateOne) AddUpdatedBy(i int64) *TroubleDeductUpdateOne {
	tduo.mutation.AddUpdatedBy(i)
	return tduo
}

// SetUpdatedAt sets the "updated_at" field.
func (tduo *TroubleDeductUpdateOne) SetUpdatedAt(t time.Time) *TroubleDeductUpdateOne {
	tduo.mutation.SetUpdatedAt(t)
	return tduo
}

// SetDeletedAt sets the "deleted_at" field.
func (tduo *TroubleDeductUpdateOne) SetDeletedAt(t time.Time) *TroubleDeductUpdateOne {
	tduo.mutation.SetDeletedAt(t)
	return tduo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tduo *TroubleDeductUpdateOne) SetNillableDeletedAt(t *time.Time) *TroubleDeductUpdateOne {
	if t != nil {
		tduo.SetDeletedAt(*t)
	}
	return tduo
}

// SetDeviceID sets the "device_id" field.
func (tduo *TroubleDeductUpdateOne) SetDeviceID(i int64) *TroubleDeductUpdateOne {
	tduo.mutation.SetDeviceID(i)
	return tduo
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (tduo *TroubleDeductUpdateOne) SetNillableDeviceID(i *int64) *TroubleDeductUpdateOne {
	if i != nil {
		tduo.SetDeviceID(*i)
	}
	return tduo
}

// SetStartedAt sets the "started_at" field.
func (tduo *TroubleDeductUpdateOne) SetStartedAt(t time.Time) *TroubleDeductUpdateOne {
	tduo.mutation.SetStartedAt(t)
	return tduo
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (tduo *TroubleDeductUpdateOne) SetNillableStartedAt(t *time.Time) *TroubleDeductUpdateOne {
	if t != nil {
		tduo.SetStartedAt(*t)
	}
	return tduo
}

// SetFinishedAt sets the "finished_at" field.
func (tduo *TroubleDeductUpdateOne) SetFinishedAt(t time.Time) *TroubleDeductUpdateOne {
	tduo.mutation.SetFinishedAt(t)
	return tduo
}

// SetNillableFinishedAt sets the "finished_at" field if the given value is not nil.
func (tduo *TroubleDeductUpdateOne) SetNillableFinishedAt(t *time.Time) *TroubleDeductUpdateOne {
	if t != nil {
		tduo.SetFinishedAt(*t)
	}
	return tduo
}

// SetTimeOfDuration sets the "time_of_duration" field.
func (tduo *TroubleDeductUpdateOne) SetTimeOfDuration(f float64) *TroubleDeductUpdateOne {
	tduo.mutation.ResetTimeOfDuration()
	tduo.mutation.SetTimeOfDuration(f)
	return tduo
}

// SetNillableTimeOfDuration sets the "time_of_duration" field if the given value is not nil.
func (tduo *TroubleDeductUpdateOne) SetNillableTimeOfDuration(f *float64) *TroubleDeductUpdateOne {
	if f != nil {
		tduo.SetTimeOfDuration(*f)
	}
	return tduo
}

// AddTimeOfDuration adds f to the "time_of_duration" field.
func (tduo *TroubleDeductUpdateOne) AddTimeOfDuration(f float64) *TroubleDeductUpdateOne {
	tduo.mutation.AddTimeOfDuration(f)
	return tduo
}

// SetAmount sets the "amount" field.
func (tduo *TroubleDeductUpdateOne) SetAmount(i int64) *TroubleDeductUpdateOne {
	tduo.mutation.ResetAmount()
	tduo.mutation.SetAmount(i)
	return tduo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tduo *TroubleDeductUpdateOne) SetNillableAmount(i *int64) *TroubleDeductUpdateOne {
	if i != nil {
		tduo.SetAmount(*i)
	}
	return tduo
}

// AddAmount adds i to the "amount" field.
func (tduo *TroubleDeductUpdateOne) AddAmount(i int64) *TroubleDeductUpdateOne {
	tduo.mutation.AddAmount(i)
	return tduo
}

// SetStatus sets the "status" field.
func (tduo *TroubleDeductUpdateOne) SetStatus(eds enums.TroubleDeductStatus) *TroubleDeductUpdateOne {
	tduo.mutation.SetStatus(eds)
	return tduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tduo *TroubleDeductUpdateOne) SetNillableStatus(eds *enums.TroubleDeductStatus) *TroubleDeductUpdateOne {
	if eds != nil {
		tduo.SetStatus(*eds)
	}
	return tduo
}

// SetReason sets the "reason" field.
func (tduo *TroubleDeductUpdateOne) SetReason(s string) *TroubleDeductUpdateOne {
	tduo.mutation.SetReason(s)
	return tduo
}

// SetNillableReason sets the "reason" field if the given value is not nil.
func (tduo *TroubleDeductUpdateOne) SetNillableReason(s *string) *TroubleDeductUpdateOne {
	if s != nil {
		tduo.SetReason(*s)
	}
	return tduo
}

// SetRejectReason sets the "reject_reason" field.
func (tduo *TroubleDeductUpdateOne) SetRejectReason(s string) *TroubleDeductUpdateOne {
	tduo.mutation.SetRejectReason(s)
	return tduo
}

// SetNillableRejectReason sets the "reject_reason" field if the given value is not nil.
func (tduo *TroubleDeductUpdateOne) SetNillableRejectReason(s *string) *TroubleDeductUpdateOne {
	if s != nil {
		tduo.SetRejectReason(*s)
	}
	return tduo
}

// SetDevice sets the "device" edge to the Device entity.
func (tduo *TroubleDeductUpdateOne) SetDevice(d *Device) *TroubleDeductUpdateOne {
	return tduo.SetDeviceID(d.ID)
}

// Mutation returns the TroubleDeductMutation object of the builder.
func (tduo *TroubleDeductUpdateOne) Mutation() *TroubleDeductMutation {
	return tduo.mutation
}

// ClearDevice clears the "device" edge to the Device entity.
func (tduo *TroubleDeductUpdateOne) ClearDevice() *TroubleDeductUpdateOne {
	tduo.mutation.ClearDevice()
	return tduo
}

// Where appends a list predicates to the TroubleDeductUpdate builder.
func (tduo *TroubleDeductUpdateOne) Where(ps ...predicate.TroubleDeduct) *TroubleDeductUpdateOne {
	tduo.mutation.Where(ps...)
	return tduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tduo *TroubleDeductUpdateOne) Select(field string, fields ...string) *TroubleDeductUpdateOne {
	tduo.fields = append([]string{field}, fields...)
	return tduo
}

// Save executes the query and returns the updated TroubleDeduct entity.
func (tduo *TroubleDeductUpdateOne) Save(ctx context.Context) (*TroubleDeduct, error) {
	tduo.defaults()
	return withHooks(ctx, tduo.sqlSave, tduo.mutation, tduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tduo *TroubleDeductUpdateOne) SaveX(ctx context.Context) *TroubleDeduct {
	node, err := tduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tduo *TroubleDeductUpdateOne) Exec(ctx context.Context) error {
	_, err := tduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tduo *TroubleDeductUpdateOne) ExecX(ctx context.Context) {
	if err := tduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tduo *TroubleDeductUpdateOne) defaults() {
	if _, ok := tduo.mutation.UpdatedAt(); !ok {
		v := troublededuct.UpdateDefaultUpdatedAt()
		tduo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tduo *TroubleDeductUpdateOne) check() error {
	if v, ok := tduo.mutation.Status(); ok {
		if err := troublededuct.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`cep_ent: validator failed for field "TroubleDeduct.status": %w`, err)}
		}
	}
	if _, ok := tduo.mutation.DeviceID(); tduo.mutation.DeviceCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "TroubleDeduct.device"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tduo *TroubleDeductUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TroubleDeductUpdateOne {
	tduo.modifiers = append(tduo.modifiers, modifiers...)
	return tduo
}

func (tduo *TroubleDeductUpdateOne) sqlSave(ctx context.Context) (_node *TroubleDeduct, err error) {
	if err := tduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(troublededuct.Table, troublededuct.Columns, sqlgraph.NewFieldSpec(troublededuct.FieldID, field.TypeInt64))
	id, ok := tduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "TroubleDeduct.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, troublededuct.FieldID)
		for _, f := range fields {
			if !troublededuct.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != troublededuct.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tduo.mutation.CreatedBy(); ok {
		_spec.SetField(troublededuct.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := tduo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(troublededuct.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := tduo.mutation.UpdatedBy(); ok {
		_spec.SetField(troublededuct.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := tduo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(troublededuct.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := tduo.mutation.UpdatedAt(); ok {
		_spec.SetField(troublededuct.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tduo.mutation.DeletedAt(); ok {
		_spec.SetField(troublededuct.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := tduo.mutation.StartedAt(); ok {
		_spec.SetField(troublededuct.FieldStartedAt, field.TypeTime, value)
	}
	if value, ok := tduo.mutation.FinishedAt(); ok {
		_spec.SetField(troublededuct.FieldFinishedAt, field.TypeTime, value)
	}
	if value, ok := tduo.mutation.TimeOfDuration(); ok {
		_spec.SetField(troublededuct.FieldTimeOfDuration, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.AddedTimeOfDuration(); ok {
		_spec.AddField(troublededuct.FieldTimeOfDuration, field.TypeFloat64, value)
	}
	if value, ok := tduo.mutation.Amount(); ok {
		_spec.SetField(troublededuct.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := tduo.mutation.AddedAmount(); ok {
		_spec.AddField(troublededuct.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := tduo.mutation.Status(); ok {
		_spec.SetField(troublededuct.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tduo.mutation.Reason(); ok {
		_spec.SetField(troublededuct.FieldReason, field.TypeString, value)
	}
	if value, ok := tduo.mutation.RejectReason(); ok {
		_spec.SetField(troublededuct.FieldRejectReason, field.TypeString, value)
	}
	if tduo.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   troublededuct.DeviceTable,
			Columns: []string{troublededuct.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tduo.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   troublededuct.DeviceTable,
			Columns: []string{troublededuct.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(tduo.modifiers...)
	_node = &TroubleDeduct{config: tduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{troublededuct.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tduo.mutation.done = true
	return _node, nil
}
