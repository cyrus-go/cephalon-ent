// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/device"
	"cephalon-ent/pkg/cep_ent/devicegpumission"
	"cephalon-ent/pkg/cep_ent/gpu"
	"cephalon-ent/pkg/cep_ent/missionkind"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeviceGpuMissionCreate is the builder for creating a DeviceGpuMission entity.
type DeviceGpuMissionCreate struct {
	config
	mutation *DeviceGpuMissionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (dgmc *DeviceGpuMissionCreate) SetCreatedBy(i int64) *DeviceGpuMissionCreate {
	dgmc.mutation.SetCreatedBy(i)
	return dgmc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (dgmc *DeviceGpuMissionCreate) SetNillableCreatedBy(i *int64) *DeviceGpuMissionCreate {
	if i != nil {
		dgmc.SetCreatedBy(*i)
	}
	return dgmc
}

// SetUpdatedBy sets the "updated_by" field.
func (dgmc *DeviceGpuMissionCreate) SetUpdatedBy(i int64) *DeviceGpuMissionCreate {
	dgmc.mutation.SetUpdatedBy(i)
	return dgmc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (dgmc *DeviceGpuMissionCreate) SetNillableUpdatedBy(i *int64) *DeviceGpuMissionCreate {
	if i != nil {
		dgmc.SetUpdatedBy(*i)
	}
	return dgmc
}

// SetCreatedAt sets the "created_at" field.
func (dgmc *DeviceGpuMissionCreate) SetCreatedAt(t time.Time) *DeviceGpuMissionCreate {
	dgmc.mutation.SetCreatedAt(t)
	return dgmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dgmc *DeviceGpuMissionCreate) SetNillableCreatedAt(t *time.Time) *DeviceGpuMissionCreate {
	if t != nil {
		dgmc.SetCreatedAt(*t)
	}
	return dgmc
}

// SetUpdatedAt sets the "updated_at" field.
func (dgmc *DeviceGpuMissionCreate) SetUpdatedAt(t time.Time) *DeviceGpuMissionCreate {
	dgmc.mutation.SetUpdatedAt(t)
	return dgmc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dgmc *DeviceGpuMissionCreate) SetNillableUpdatedAt(t *time.Time) *DeviceGpuMissionCreate {
	if t != nil {
		dgmc.SetUpdatedAt(*t)
	}
	return dgmc
}

// SetDeletedAt sets the "deleted_at" field.
func (dgmc *DeviceGpuMissionCreate) SetDeletedAt(t time.Time) *DeviceGpuMissionCreate {
	dgmc.mutation.SetDeletedAt(t)
	return dgmc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dgmc *DeviceGpuMissionCreate) SetNillableDeletedAt(t *time.Time) *DeviceGpuMissionCreate {
	if t != nil {
		dgmc.SetDeletedAt(*t)
	}
	return dgmc
}

// SetDeviceID sets the "device_id" field.
func (dgmc *DeviceGpuMissionCreate) SetDeviceID(i int64) *DeviceGpuMissionCreate {
	dgmc.mutation.SetDeviceID(i)
	return dgmc
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (dgmc *DeviceGpuMissionCreate) SetNillableDeviceID(i *int64) *DeviceGpuMissionCreate {
	if i != nil {
		dgmc.SetDeviceID(*i)
	}
	return dgmc
}

// SetGpuID sets the "gpu_id" field.
func (dgmc *DeviceGpuMissionCreate) SetGpuID(i int64) *DeviceGpuMissionCreate {
	dgmc.mutation.SetGpuID(i)
	return dgmc
}

// SetNillableGpuID sets the "gpu_id" field if the given value is not nil.
func (dgmc *DeviceGpuMissionCreate) SetNillableGpuID(i *int64) *DeviceGpuMissionCreate {
	if i != nil {
		dgmc.SetGpuID(*i)
	}
	return dgmc
}

// SetMissionKindID sets the "mission_kind_id" field.
func (dgmc *DeviceGpuMissionCreate) SetMissionKindID(i int64) *DeviceGpuMissionCreate {
	dgmc.mutation.SetMissionKindID(i)
	return dgmc
}

// SetNillableMissionKindID sets the "mission_kind_id" field if the given value is not nil.
func (dgmc *DeviceGpuMissionCreate) SetNillableMissionKindID(i *int64) *DeviceGpuMissionCreate {
	if i != nil {
		dgmc.SetMissionKindID(*i)
	}
	return dgmc
}

// SetID sets the "id" field.
func (dgmc *DeviceGpuMissionCreate) SetID(i int64) *DeviceGpuMissionCreate {
	dgmc.mutation.SetID(i)
	return dgmc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dgmc *DeviceGpuMissionCreate) SetNillableID(i *int64) *DeviceGpuMissionCreate {
	if i != nil {
		dgmc.SetID(*i)
	}
	return dgmc
}

// SetDevice sets the "device" edge to the Device entity.
func (dgmc *DeviceGpuMissionCreate) SetDevice(d *Device) *DeviceGpuMissionCreate {
	return dgmc.SetDeviceID(d.ID)
}

// SetMissionKind sets the "mission_kind" edge to the MissionKind entity.
func (dgmc *DeviceGpuMissionCreate) SetMissionKind(m *MissionKind) *DeviceGpuMissionCreate {
	return dgmc.SetMissionKindID(m.ID)
}

// SetGpu sets the "gpu" edge to the Gpu entity.
func (dgmc *DeviceGpuMissionCreate) SetGpu(g *Gpu) *DeviceGpuMissionCreate {
	return dgmc.SetGpuID(g.ID)
}

// Mutation returns the DeviceGpuMissionMutation object of the builder.
func (dgmc *DeviceGpuMissionCreate) Mutation() *DeviceGpuMissionMutation {
	return dgmc.mutation
}

// Save creates the DeviceGpuMission in the database.
func (dgmc *DeviceGpuMissionCreate) Save(ctx context.Context) (*DeviceGpuMission, error) {
	dgmc.defaults()
	return withHooks(ctx, dgmc.sqlSave, dgmc.mutation, dgmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dgmc *DeviceGpuMissionCreate) SaveX(ctx context.Context) *DeviceGpuMission {
	v, err := dgmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dgmc *DeviceGpuMissionCreate) Exec(ctx context.Context) error {
	_, err := dgmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dgmc *DeviceGpuMissionCreate) ExecX(ctx context.Context) {
	if err := dgmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dgmc *DeviceGpuMissionCreate) defaults() {
	if _, ok := dgmc.mutation.CreatedBy(); !ok {
		v := devicegpumission.DefaultCreatedBy
		dgmc.mutation.SetCreatedBy(v)
	}
	if _, ok := dgmc.mutation.UpdatedBy(); !ok {
		v := devicegpumission.DefaultUpdatedBy
		dgmc.mutation.SetUpdatedBy(v)
	}
	if _, ok := dgmc.mutation.CreatedAt(); !ok {
		v := devicegpumission.DefaultCreatedAt()
		dgmc.mutation.SetCreatedAt(v)
	}
	if _, ok := dgmc.mutation.UpdatedAt(); !ok {
		v := devicegpumission.DefaultUpdatedAt()
		dgmc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dgmc.mutation.DeletedAt(); !ok {
		v := devicegpumission.DefaultDeletedAt
		dgmc.mutation.SetDeletedAt(v)
	}
	if _, ok := dgmc.mutation.DeviceID(); !ok {
		v := devicegpumission.DefaultDeviceID
		dgmc.mutation.SetDeviceID(v)
	}
	if _, ok := dgmc.mutation.GpuID(); !ok {
		v := devicegpumission.DefaultGpuID
		dgmc.mutation.SetGpuID(v)
	}
	if _, ok := dgmc.mutation.MissionKindID(); !ok {
		v := devicegpumission.DefaultMissionKindID
		dgmc.mutation.SetMissionKindID(v)
	}
	if _, ok := dgmc.mutation.ID(); !ok {
		v := devicegpumission.DefaultID()
		dgmc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dgmc *DeviceGpuMissionCreate) check() error {
	if _, ok := dgmc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "DeviceGpuMission.created_by"`)}
	}
	if _, ok := dgmc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "DeviceGpuMission.updated_by"`)}
	}
	if _, ok := dgmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "DeviceGpuMission.created_at"`)}
	}
	if _, ok := dgmc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "DeviceGpuMission.updated_at"`)}
	}
	if _, ok := dgmc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "DeviceGpuMission.deleted_at"`)}
	}
	if _, ok := dgmc.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device_id", err: errors.New(`cep_ent: missing required field "DeviceGpuMission.device_id"`)}
	}
	if _, ok := dgmc.mutation.GpuID(); !ok {
		return &ValidationError{Name: "gpu_id", err: errors.New(`cep_ent: missing required field "DeviceGpuMission.gpu_id"`)}
	}
	if _, ok := dgmc.mutation.MissionKindID(); !ok {
		return &ValidationError{Name: "mission_kind_id", err: errors.New(`cep_ent: missing required field "DeviceGpuMission.mission_kind_id"`)}
	}
	if _, ok := dgmc.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device", err: errors.New(`cep_ent: missing required edge "DeviceGpuMission.device"`)}
	}
	if _, ok := dgmc.mutation.MissionKindID(); !ok {
		return &ValidationError{Name: "mission_kind", err: errors.New(`cep_ent: missing required edge "DeviceGpuMission.mission_kind"`)}
	}
	if _, ok := dgmc.mutation.GpuID(); !ok {
		return &ValidationError{Name: "gpu", err: errors.New(`cep_ent: missing required edge "DeviceGpuMission.gpu"`)}
	}
	return nil
}

func (dgmc *DeviceGpuMissionCreate) sqlSave(ctx context.Context) (*DeviceGpuMission, error) {
	if err := dgmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dgmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dgmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	dgmc.mutation.id = &_node.ID
	dgmc.mutation.done = true
	return _node, nil
}

func (dgmc *DeviceGpuMissionCreate) createSpec() (*DeviceGpuMission, *sqlgraph.CreateSpec) {
	var (
		_node = &DeviceGpuMission{config: dgmc.config}
		_spec = sqlgraph.NewCreateSpec(devicegpumission.Table, sqlgraph.NewFieldSpec(devicegpumission.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = dgmc.conflict
	if id, ok := dgmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dgmc.mutation.CreatedBy(); ok {
		_spec.SetField(devicegpumission.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := dgmc.mutation.UpdatedBy(); ok {
		_spec.SetField(devicegpumission.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := dgmc.mutation.CreatedAt(); ok {
		_spec.SetField(devicegpumission.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dgmc.mutation.UpdatedAt(); ok {
		_spec.SetField(devicegpumission.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dgmc.mutation.DeletedAt(); ok {
		_spec.SetField(devicegpumission.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if nodes := dgmc.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicegpumission.DeviceTable,
			Columns: []string{devicegpumission.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DeviceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dgmc.mutation.MissionKindIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicegpumission.MissionKindTable,
			Columns: []string{devicegpumission.MissionKindColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionkind.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MissionKindID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dgmc.mutation.GpuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   devicegpumission.GpuTable,
			Columns: []string{devicegpumission.GpuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gpu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.GpuID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeviceGpuMission.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceGpuMissionUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (dgmc *DeviceGpuMissionCreate) OnConflict(opts ...sql.ConflictOption) *DeviceGpuMissionUpsertOne {
	dgmc.conflict = opts
	return &DeviceGpuMissionUpsertOne{
		create: dgmc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceGpuMission.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dgmc *DeviceGpuMissionCreate) OnConflictColumns(columns ...string) *DeviceGpuMissionUpsertOne {
	dgmc.conflict = append(dgmc.conflict, sql.ConflictColumns(columns...))
	return &DeviceGpuMissionUpsertOne{
		create: dgmc,
	}
}

type (
	// DeviceGpuMissionUpsertOne is the builder for "upsert"-ing
	//  one DeviceGpuMission node.
	DeviceGpuMissionUpsertOne struct {
		create *DeviceGpuMissionCreate
	}

	// DeviceGpuMissionUpsert is the "OnConflict" setter.
	DeviceGpuMissionUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *DeviceGpuMissionUpsert) SetCreatedBy(v int64) *DeviceGpuMissionUpsert {
	u.Set(devicegpumission.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsert) UpdateCreatedBy() *DeviceGpuMissionUpsert {
	u.SetExcluded(devicegpumission.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *DeviceGpuMissionUpsert) AddCreatedBy(v int64) *DeviceGpuMissionUpsert {
	u.Add(devicegpumission.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *DeviceGpuMissionUpsert) SetUpdatedBy(v int64) *DeviceGpuMissionUpsert {
	u.Set(devicegpumission.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsert) UpdateUpdatedBy() *DeviceGpuMissionUpsert {
	u.SetExcluded(devicegpumission.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *DeviceGpuMissionUpsert) AddUpdatedBy(v int64) *DeviceGpuMissionUpsert {
	u.Add(devicegpumission.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceGpuMissionUpsert) SetUpdatedAt(v time.Time) *DeviceGpuMissionUpsert {
	u.Set(devicegpumission.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsert) UpdateUpdatedAt() *DeviceGpuMissionUpsert {
	u.SetExcluded(devicegpumission.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceGpuMissionUpsert) SetDeletedAt(v time.Time) *DeviceGpuMissionUpsert {
	u.Set(devicegpumission.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsert) UpdateDeletedAt() *DeviceGpuMissionUpsert {
	u.SetExcluded(devicegpumission.FieldDeletedAt)
	return u
}

// SetDeviceID sets the "device_id" field.
func (u *DeviceGpuMissionUpsert) SetDeviceID(v int64) *DeviceGpuMissionUpsert {
	u.Set(devicegpumission.FieldDeviceID, v)
	return u
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsert) UpdateDeviceID() *DeviceGpuMissionUpsert {
	u.SetExcluded(devicegpumission.FieldDeviceID)
	return u
}

// SetGpuID sets the "gpu_id" field.
func (u *DeviceGpuMissionUpsert) SetGpuID(v int64) *DeviceGpuMissionUpsert {
	u.Set(devicegpumission.FieldGpuID, v)
	return u
}

// UpdateGpuID sets the "gpu_id" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsert) UpdateGpuID() *DeviceGpuMissionUpsert {
	u.SetExcluded(devicegpumission.FieldGpuID)
	return u
}

// SetMissionKindID sets the "mission_kind_id" field.
func (u *DeviceGpuMissionUpsert) SetMissionKindID(v int64) *DeviceGpuMissionUpsert {
	u.Set(devicegpumission.FieldMissionKindID, v)
	return u
}

// UpdateMissionKindID sets the "mission_kind_id" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsert) UpdateMissionKindID() *DeviceGpuMissionUpsert {
	u.SetExcluded(devicegpumission.FieldMissionKindID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DeviceGpuMission.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(devicegpumission.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceGpuMissionUpsertOne) UpdateNewValues() *DeviceGpuMissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(devicegpumission.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(devicegpumission.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeviceGpuMission.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DeviceGpuMissionUpsertOne) Ignore() *DeviceGpuMissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceGpuMissionUpsertOne) DoNothing() *DeviceGpuMissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceGpuMissionCreate.OnConflict
// documentation for more info.
func (u *DeviceGpuMissionUpsertOne) Update(set func(*DeviceGpuMissionUpsert)) *DeviceGpuMissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceGpuMissionUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *DeviceGpuMissionUpsertOne) SetCreatedBy(v int64) *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *DeviceGpuMissionUpsertOne) AddCreatedBy(v int64) *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsertOne) UpdateCreatedBy() *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *DeviceGpuMissionUpsertOne) SetUpdatedBy(v int64) *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *DeviceGpuMissionUpsertOne) AddUpdatedBy(v int64) *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsertOne) UpdateUpdatedBy() *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceGpuMissionUpsertOne) SetUpdatedAt(v time.Time) *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsertOne) UpdateUpdatedAt() *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceGpuMissionUpsertOne) SetDeletedAt(v time.Time) *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsertOne) UpdateDeletedAt() *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *DeviceGpuMissionUpsertOne) SetDeviceID(v int64) *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.SetDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsertOne) UpdateDeviceID() *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.UpdateDeviceID()
	})
}

// SetGpuID sets the "gpu_id" field.
func (u *DeviceGpuMissionUpsertOne) SetGpuID(v int64) *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.SetGpuID(v)
	})
}

// UpdateGpuID sets the "gpu_id" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsertOne) UpdateGpuID() *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.UpdateGpuID()
	})
}

// SetMissionKindID sets the "mission_kind_id" field.
func (u *DeviceGpuMissionUpsertOne) SetMissionKindID(v int64) *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.SetMissionKindID(v)
	})
}

// UpdateMissionKindID sets the "mission_kind_id" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsertOne) UpdateMissionKindID() *DeviceGpuMissionUpsertOne {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.UpdateMissionKindID()
	})
}

// Exec executes the query.
func (u *DeviceGpuMissionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for DeviceGpuMissionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceGpuMissionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeviceGpuMissionUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeviceGpuMissionUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeviceGpuMissionCreateBulk is the builder for creating many DeviceGpuMission entities in bulk.
type DeviceGpuMissionCreateBulk struct {
	config
	builders []*DeviceGpuMissionCreate
	conflict []sql.ConflictOption
}

// Save creates the DeviceGpuMission entities in the database.
func (dgmcb *DeviceGpuMissionCreateBulk) Save(ctx context.Context) ([]*DeviceGpuMission, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dgmcb.builders))
	nodes := make([]*DeviceGpuMission, len(dgmcb.builders))
	mutators := make([]Mutator, len(dgmcb.builders))
	for i := range dgmcb.builders {
		func(i int, root context.Context) {
			builder := dgmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceGpuMissionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dgmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dgmcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dgmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dgmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dgmcb *DeviceGpuMissionCreateBulk) SaveX(ctx context.Context) []*DeviceGpuMission {
	v, err := dgmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dgmcb *DeviceGpuMissionCreateBulk) Exec(ctx context.Context) error {
	_, err := dgmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dgmcb *DeviceGpuMissionCreateBulk) ExecX(ctx context.Context) {
	if err := dgmcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeviceGpuMission.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceGpuMissionUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (dgmcb *DeviceGpuMissionCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeviceGpuMissionUpsertBulk {
	dgmcb.conflict = opts
	return &DeviceGpuMissionUpsertBulk{
		create: dgmcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceGpuMission.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dgmcb *DeviceGpuMissionCreateBulk) OnConflictColumns(columns ...string) *DeviceGpuMissionUpsertBulk {
	dgmcb.conflict = append(dgmcb.conflict, sql.ConflictColumns(columns...))
	return &DeviceGpuMissionUpsertBulk{
		create: dgmcb,
	}
}

// DeviceGpuMissionUpsertBulk is the builder for "upsert"-ing
// a bulk of DeviceGpuMission nodes.
type DeviceGpuMissionUpsertBulk struct {
	create *DeviceGpuMissionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DeviceGpuMission.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(devicegpumission.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceGpuMissionUpsertBulk) UpdateNewValues() *DeviceGpuMissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(devicegpumission.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(devicegpumission.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeviceGpuMission.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DeviceGpuMissionUpsertBulk) Ignore() *DeviceGpuMissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceGpuMissionUpsertBulk) DoNothing() *DeviceGpuMissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceGpuMissionCreateBulk.OnConflict
// documentation for more info.
func (u *DeviceGpuMissionUpsertBulk) Update(set func(*DeviceGpuMissionUpsert)) *DeviceGpuMissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceGpuMissionUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *DeviceGpuMissionUpsertBulk) SetCreatedBy(v int64) *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *DeviceGpuMissionUpsertBulk) AddCreatedBy(v int64) *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsertBulk) UpdateCreatedBy() *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *DeviceGpuMissionUpsertBulk) SetUpdatedBy(v int64) *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *DeviceGpuMissionUpsertBulk) AddUpdatedBy(v int64) *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsertBulk) UpdateUpdatedBy() *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceGpuMissionUpsertBulk) SetUpdatedAt(v time.Time) *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsertBulk) UpdateUpdatedAt() *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceGpuMissionUpsertBulk) SetDeletedAt(v time.Time) *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsertBulk) UpdateDeletedAt() *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *DeviceGpuMissionUpsertBulk) SetDeviceID(v int64) *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.SetDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsertBulk) UpdateDeviceID() *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.UpdateDeviceID()
	})
}

// SetGpuID sets the "gpu_id" field.
func (u *DeviceGpuMissionUpsertBulk) SetGpuID(v int64) *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.SetGpuID(v)
	})
}

// UpdateGpuID sets the "gpu_id" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsertBulk) UpdateGpuID() *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.UpdateGpuID()
	})
}

// SetMissionKindID sets the "mission_kind_id" field.
func (u *DeviceGpuMissionUpsertBulk) SetMissionKindID(v int64) *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.SetMissionKindID(v)
	})
}

// UpdateMissionKindID sets the "mission_kind_id" field to the value that was provided on create.
func (u *DeviceGpuMissionUpsertBulk) UpdateMissionKindID() *DeviceGpuMissionUpsertBulk {
	return u.Update(func(s *DeviceGpuMissionUpsert) {
		s.UpdateMissionKindID()
	})
}

// Exec executes the query.
func (u *DeviceGpuMissionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("cep_ent: OnConflict was set for builder %d. Set it on the DeviceGpuMissionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for DeviceGpuMissionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceGpuMissionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
