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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/devicegpumission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/gpu"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// GpuUpdate is the builder for updating Gpu entities.
type GpuUpdate struct {
	config
	hooks    []Hook
	mutation *GpuMutation
}

// Where appends a list predicates to the GpuUpdate builder.
func (gu *GpuUpdate) Where(ps ...predicate.Gpu) *GpuUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetCreatedBy sets the "created_by" field.
func (gu *GpuUpdate) SetCreatedBy(i int64) *GpuUpdate {
	gu.mutation.ResetCreatedBy()
	gu.mutation.SetCreatedBy(i)
	return gu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (gu *GpuUpdate) SetNillableCreatedBy(i *int64) *GpuUpdate {
	if i != nil {
		gu.SetCreatedBy(*i)
	}
	return gu
}

// AddCreatedBy adds i to the "created_by" field.
func (gu *GpuUpdate) AddCreatedBy(i int64) *GpuUpdate {
	gu.mutation.AddCreatedBy(i)
	return gu
}

// SetUpdatedBy sets the "updated_by" field.
func (gu *GpuUpdate) SetUpdatedBy(i int64) *GpuUpdate {
	gu.mutation.ResetUpdatedBy()
	gu.mutation.SetUpdatedBy(i)
	return gu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gu *GpuUpdate) SetNillableUpdatedBy(i *int64) *GpuUpdate {
	if i != nil {
		gu.SetUpdatedBy(*i)
	}
	return gu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (gu *GpuUpdate) AddUpdatedBy(i int64) *GpuUpdate {
	gu.mutation.AddUpdatedBy(i)
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GpuUpdate) SetUpdatedAt(t time.Time) *GpuUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// SetDeletedAt sets the "deleted_at" field.
func (gu *GpuUpdate) SetDeletedAt(t time.Time) *GpuUpdate {
	gu.mutation.SetDeletedAt(t)
	return gu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gu *GpuUpdate) SetNillableDeletedAt(t *time.Time) *GpuUpdate {
	if t != nil {
		gu.SetDeletedAt(*t)
	}
	return gu
}

// SetVersion sets the "version" field.
func (gu *GpuUpdate) SetVersion(ev enums.GpuVersion) *GpuUpdate {
	gu.mutation.SetVersion(ev)
	return gu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (gu *GpuUpdate) SetNillableVersion(ev *enums.GpuVersion) *GpuUpdate {
	if ev != nil {
		gu.SetVersion(*ev)
	}
	return gu
}

// AddDeviceGpuMissionIDs adds the "device_gpu_missions" edge to the DeviceGpuMission entity by IDs.
func (gu *GpuUpdate) AddDeviceGpuMissionIDs(ids ...int64) *GpuUpdate {
	gu.mutation.AddDeviceGpuMissionIDs(ids...)
	return gu
}

// AddDeviceGpuMissions adds the "device_gpu_missions" edges to the DeviceGpuMission entity.
func (gu *GpuUpdate) AddDeviceGpuMissions(d ...*DeviceGpuMission) *GpuUpdate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gu.AddDeviceGpuMissionIDs(ids...)
}

// Mutation returns the GpuMutation object of the builder.
func (gu *GpuUpdate) Mutation() *GpuMutation {
	return gu.mutation
}

// ClearDeviceGpuMissions clears all "device_gpu_missions" edges to the DeviceGpuMission entity.
func (gu *GpuUpdate) ClearDeviceGpuMissions() *GpuUpdate {
	gu.mutation.ClearDeviceGpuMissions()
	return gu
}

// RemoveDeviceGpuMissionIDs removes the "device_gpu_missions" edge to DeviceGpuMission entities by IDs.
func (gu *GpuUpdate) RemoveDeviceGpuMissionIDs(ids ...int64) *GpuUpdate {
	gu.mutation.RemoveDeviceGpuMissionIDs(ids...)
	return gu
}

// RemoveDeviceGpuMissions removes "device_gpu_missions" edges to DeviceGpuMission entities.
func (gu *GpuUpdate) RemoveDeviceGpuMissions(d ...*DeviceGpuMission) *GpuUpdate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gu.RemoveDeviceGpuMissionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GpuUpdate) Save(ctx context.Context) (int, error) {
	gu.defaults()
	return withHooks(ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GpuUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GpuUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GpuUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GpuUpdate) defaults() {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		v := gpu.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GpuUpdate) check() error {
	if v, ok := gu.mutation.Version(); ok {
		if err := gpu.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`cep_ent: validator failed for field "Gpu.version": %w`, err)}
		}
	}
	return nil
}

func (gu *GpuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(gpu.Table, gpu.Columns, sqlgraph.NewFieldSpec(gpu.FieldID, field.TypeInt64))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.CreatedBy(); ok {
		_spec.SetField(gpu.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := gu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(gpu.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := gu.mutation.UpdatedBy(); ok {
		_spec.SetField(gpu.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := gu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(gpu.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.SetField(gpu.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gu.mutation.DeletedAt(); ok {
		_spec.SetField(gpu.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := gu.mutation.Version(); ok {
		_spec.SetField(gpu.FieldVersion, field.TypeEnum, value)
	}
	if gu.mutation.DeviceGpuMissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gpu.DeviceGpuMissionsTable,
			Columns: []string{gpu.DeviceGpuMissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(devicegpumission.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedDeviceGpuMissionsIDs(); len(nodes) > 0 && !gu.mutation.DeviceGpuMissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gpu.DeviceGpuMissionsTable,
			Columns: []string{gpu.DeviceGpuMissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(devicegpumission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.DeviceGpuMissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gpu.DeviceGpuMissionsTable,
			Columns: []string{gpu.DeviceGpuMissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(devicegpumission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gpu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GpuUpdateOne is the builder for updating a single Gpu entity.
type GpuUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GpuMutation
}

// SetCreatedBy sets the "created_by" field.
func (guo *GpuUpdateOne) SetCreatedBy(i int64) *GpuUpdateOne {
	guo.mutation.ResetCreatedBy()
	guo.mutation.SetCreatedBy(i)
	return guo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (guo *GpuUpdateOne) SetNillableCreatedBy(i *int64) *GpuUpdateOne {
	if i != nil {
		guo.SetCreatedBy(*i)
	}
	return guo
}

// AddCreatedBy adds i to the "created_by" field.
func (guo *GpuUpdateOne) AddCreatedBy(i int64) *GpuUpdateOne {
	guo.mutation.AddCreatedBy(i)
	return guo
}

// SetUpdatedBy sets the "updated_by" field.
func (guo *GpuUpdateOne) SetUpdatedBy(i int64) *GpuUpdateOne {
	guo.mutation.ResetUpdatedBy()
	guo.mutation.SetUpdatedBy(i)
	return guo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (guo *GpuUpdateOne) SetNillableUpdatedBy(i *int64) *GpuUpdateOne {
	if i != nil {
		guo.SetUpdatedBy(*i)
	}
	return guo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (guo *GpuUpdateOne) AddUpdatedBy(i int64) *GpuUpdateOne {
	guo.mutation.AddUpdatedBy(i)
	return guo
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GpuUpdateOne) SetUpdatedAt(t time.Time) *GpuUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// SetDeletedAt sets the "deleted_at" field.
func (guo *GpuUpdateOne) SetDeletedAt(t time.Time) *GpuUpdateOne {
	guo.mutation.SetDeletedAt(t)
	return guo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guo *GpuUpdateOne) SetNillableDeletedAt(t *time.Time) *GpuUpdateOne {
	if t != nil {
		guo.SetDeletedAt(*t)
	}
	return guo
}

// SetVersion sets the "version" field.
func (guo *GpuUpdateOne) SetVersion(ev enums.GpuVersion) *GpuUpdateOne {
	guo.mutation.SetVersion(ev)
	return guo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (guo *GpuUpdateOne) SetNillableVersion(ev *enums.GpuVersion) *GpuUpdateOne {
	if ev != nil {
		guo.SetVersion(*ev)
	}
	return guo
}

// AddDeviceGpuMissionIDs adds the "device_gpu_missions" edge to the DeviceGpuMission entity by IDs.
func (guo *GpuUpdateOne) AddDeviceGpuMissionIDs(ids ...int64) *GpuUpdateOne {
	guo.mutation.AddDeviceGpuMissionIDs(ids...)
	return guo
}

// AddDeviceGpuMissions adds the "device_gpu_missions" edges to the DeviceGpuMission entity.
func (guo *GpuUpdateOne) AddDeviceGpuMissions(d ...*DeviceGpuMission) *GpuUpdateOne {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return guo.AddDeviceGpuMissionIDs(ids...)
}

// Mutation returns the GpuMutation object of the builder.
func (guo *GpuUpdateOne) Mutation() *GpuMutation {
	return guo.mutation
}

// ClearDeviceGpuMissions clears all "device_gpu_missions" edges to the DeviceGpuMission entity.
func (guo *GpuUpdateOne) ClearDeviceGpuMissions() *GpuUpdateOne {
	guo.mutation.ClearDeviceGpuMissions()
	return guo
}

// RemoveDeviceGpuMissionIDs removes the "device_gpu_missions" edge to DeviceGpuMission entities by IDs.
func (guo *GpuUpdateOne) RemoveDeviceGpuMissionIDs(ids ...int64) *GpuUpdateOne {
	guo.mutation.RemoveDeviceGpuMissionIDs(ids...)
	return guo
}

// RemoveDeviceGpuMissions removes "device_gpu_missions" edges to DeviceGpuMission entities.
func (guo *GpuUpdateOne) RemoveDeviceGpuMissions(d ...*DeviceGpuMission) *GpuUpdateOne {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return guo.RemoveDeviceGpuMissionIDs(ids...)
}

// Where appends a list predicates to the GpuUpdate builder.
func (guo *GpuUpdateOne) Where(ps ...predicate.Gpu) *GpuUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GpuUpdateOne) Select(field string, fields ...string) *GpuUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Gpu entity.
func (guo *GpuUpdateOne) Save(ctx context.Context) (*Gpu, error) {
	guo.defaults()
	return withHooks(ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GpuUpdateOne) SaveX(ctx context.Context) *Gpu {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GpuUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GpuUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GpuUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		v := gpu.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GpuUpdateOne) check() error {
	if v, ok := guo.mutation.Version(); ok {
		if err := gpu.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`cep_ent: validator failed for field "Gpu.version": %w`, err)}
		}
	}
	return nil
}

func (guo *GpuUpdateOne) sqlSave(ctx context.Context) (_node *Gpu, err error) {
	if err := guo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(gpu.Table, gpu.Columns, sqlgraph.NewFieldSpec(gpu.FieldID, field.TypeInt64))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "Gpu.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gpu.FieldID)
		for _, f := range fields {
			if !gpu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != gpu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.CreatedBy(); ok {
		_spec.SetField(gpu.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := guo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(gpu.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := guo.mutation.UpdatedBy(); ok {
		_spec.SetField(gpu.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := guo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(gpu.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.SetField(gpu.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := guo.mutation.DeletedAt(); ok {
		_spec.SetField(gpu.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := guo.mutation.Version(); ok {
		_spec.SetField(gpu.FieldVersion, field.TypeEnum, value)
	}
	if guo.mutation.DeviceGpuMissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gpu.DeviceGpuMissionsTable,
			Columns: []string{gpu.DeviceGpuMissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(devicegpumission.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedDeviceGpuMissionsIDs(); len(nodes) > 0 && !guo.mutation.DeviceGpuMissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gpu.DeviceGpuMissionsTable,
			Columns: []string{gpu.DeviceGpuMissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(devicegpumission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.DeviceGpuMissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gpu.DeviceGpuMissionsTable,
			Columns: []string{gpu.DeviceGpuMissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(devicegpumission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Gpu{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gpu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}