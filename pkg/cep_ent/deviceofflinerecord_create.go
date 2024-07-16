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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/deviceofflinerecord"
)

// DeviceOfflineRecordCreate is the builder for creating a DeviceOfflineRecord entity.
type DeviceOfflineRecordCreate struct {
	config
	mutation *DeviceOfflineRecordMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (dorc *DeviceOfflineRecordCreate) SetCreatedBy(i int64) *DeviceOfflineRecordCreate {
	dorc.mutation.SetCreatedBy(i)
	return dorc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (dorc *DeviceOfflineRecordCreate) SetNillableCreatedBy(i *int64) *DeviceOfflineRecordCreate {
	if i != nil {
		dorc.SetCreatedBy(*i)
	}
	return dorc
}

// SetUpdatedBy sets the "updated_by" field.
func (dorc *DeviceOfflineRecordCreate) SetUpdatedBy(i int64) *DeviceOfflineRecordCreate {
	dorc.mutation.SetUpdatedBy(i)
	return dorc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (dorc *DeviceOfflineRecordCreate) SetNillableUpdatedBy(i *int64) *DeviceOfflineRecordCreate {
	if i != nil {
		dorc.SetUpdatedBy(*i)
	}
	return dorc
}

// SetCreatedAt sets the "created_at" field.
func (dorc *DeviceOfflineRecordCreate) SetCreatedAt(t time.Time) *DeviceOfflineRecordCreate {
	dorc.mutation.SetCreatedAt(t)
	return dorc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dorc *DeviceOfflineRecordCreate) SetNillableCreatedAt(t *time.Time) *DeviceOfflineRecordCreate {
	if t != nil {
		dorc.SetCreatedAt(*t)
	}
	return dorc
}

// SetUpdatedAt sets the "updated_at" field.
func (dorc *DeviceOfflineRecordCreate) SetUpdatedAt(t time.Time) *DeviceOfflineRecordCreate {
	dorc.mutation.SetUpdatedAt(t)
	return dorc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dorc *DeviceOfflineRecordCreate) SetNillableUpdatedAt(t *time.Time) *DeviceOfflineRecordCreate {
	if t != nil {
		dorc.SetUpdatedAt(*t)
	}
	return dorc
}

// SetDeletedAt sets the "deleted_at" field.
func (dorc *DeviceOfflineRecordCreate) SetDeletedAt(t time.Time) *DeviceOfflineRecordCreate {
	dorc.mutation.SetDeletedAt(t)
	return dorc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dorc *DeviceOfflineRecordCreate) SetNillableDeletedAt(t *time.Time) *DeviceOfflineRecordCreate {
	if t != nil {
		dorc.SetDeletedAt(*t)
	}
	return dorc
}

// SetDeviceID sets the "device_id" field.
func (dorc *DeviceOfflineRecordCreate) SetDeviceID(i int64) *DeviceOfflineRecordCreate {
	dorc.mutation.SetDeviceID(i)
	return dorc
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (dorc *DeviceOfflineRecordCreate) SetNillableDeviceID(i *int64) *DeviceOfflineRecordCreate {
	if i != nil {
		dorc.SetDeviceID(*i)
	}
	return dorc
}

// SetID sets the "id" field.
func (dorc *DeviceOfflineRecordCreate) SetID(i int64) *DeviceOfflineRecordCreate {
	dorc.mutation.SetID(i)
	return dorc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dorc *DeviceOfflineRecordCreate) SetNillableID(i *int64) *DeviceOfflineRecordCreate {
	if i != nil {
		dorc.SetID(*i)
	}
	return dorc
}

// SetDevice sets the "device" edge to the Device entity.
func (dorc *DeviceOfflineRecordCreate) SetDevice(d *Device) *DeviceOfflineRecordCreate {
	return dorc.SetDeviceID(d.ID)
}

// Mutation returns the DeviceOfflineRecordMutation object of the builder.
func (dorc *DeviceOfflineRecordCreate) Mutation() *DeviceOfflineRecordMutation {
	return dorc.mutation
}

// Save creates the DeviceOfflineRecord in the database.
func (dorc *DeviceOfflineRecordCreate) Save(ctx context.Context) (*DeviceOfflineRecord, error) {
	dorc.defaults()
	return withHooks(ctx, dorc.sqlSave, dorc.mutation, dorc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dorc *DeviceOfflineRecordCreate) SaveX(ctx context.Context) *DeviceOfflineRecord {
	v, err := dorc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dorc *DeviceOfflineRecordCreate) Exec(ctx context.Context) error {
	_, err := dorc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dorc *DeviceOfflineRecordCreate) ExecX(ctx context.Context) {
	if err := dorc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dorc *DeviceOfflineRecordCreate) defaults() {
	if _, ok := dorc.mutation.CreatedBy(); !ok {
		v := deviceofflinerecord.DefaultCreatedBy
		dorc.mutation.SetCreatedBy(v)
	}
	if _, ok := dorc.mutation.UpdatedBy(); !ok {
		v := deviceofflinerecord.DefaultUpdatedBy
		dorc.mutation.SetUpdatedBy(v)
	}
	if _, ok := dorc.mutation.CreatedAt(); !ok {
		v := deviceofflinerecord.DefaultCreatedAt()
		dorc.mutation.SetCreatedAt(v)
	}
	if _, ok := dorc.mutation.UpdatedAt(); !ok {
		v := deviceofflinerecord.DefaultUpdatedAt()
		dorc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dorc.mutation.DeletedAt(); !ok {
		v := deviceofflinerecord.DefaultDeletedAt
		dorc.mutation.SetDeletedAt(v)
	}
	if _, ok := dorc.mutation.DeviceID(); !ok {
		v := deviceofflinerecord.DefaultDeviceID
		dorc.mutation.SetDeviceID(v)
	}
	if _, ok := dorc.mutation.ID(); !ok {
		v := deviceofflinerecord.DefaultID()
		dorc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dorc *DeviceOfflineRecordCreate) check() error {
	if _, ok := dorc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "DeviceOfflineRecord.created_by"`)}
	}
	if _, ok := dorc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "DeviceOfflineRecord.updated_by"`)}
	}
	if _, ok := dorc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "DeviceOfflineRecord.created_at"`)}
	}
	if _, ok := dorc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "DeviceOfflineRecord.updated_at"`)}
	}
	if _, ok := dorc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "DeviceOfflineRecord.deleted_at"`)}
	}
	if _, ok := dorc.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device_id", err: errors.New(`cep_ent: missing required field "DeviceOfflineRecord.device_id"`)}
	}
	if _, ok := dorc.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device", err: errors.New(`cep_ent: missing required edge "DeviceOfflineRecord.device"`)}
	}
	return nil
}

func (dorc *DeviceOfflineRecordCreate) sqlSave(ctx context.Context) (*DeviceOfflineRecord, error) {
	if err := dorc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dorc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dorc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	dorc.mutation.id = &_node.ID
	dorc.mutation.done = true
	return _node, nil
}

func (dorc *DeviceOfflineRecordCreate) createSpec() (*DeviceOfflineRecord, *sqlgraph.CreateSpec) {
	var (
		_node = &DeviceOfflineRecord{config: dorc.config}
		_spec = sqlgraph.NewCreateSpec(deviceofflinerecord.Table, sqlgraph.NewFieldSpec(deviceofflinerecord.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = dorc.conflict
	if id, ok := dorc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dorc.mutation.CreatedBy(); ok {
		_spec.SetField(deviceofflinerecord.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := dorc.mutation.UpdatedBy(); ok {
		_spec.SetField(deviceofflinerecord.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := dorc.mutation.CreatedAt(); ok {
		_spec.SetField(deviceofflinerecord.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dorc.mutation.UpdatedAt(); ok {
		_spec.SetField(deviceofflinerecord.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dorc.mutation.DeletedAt(); ok {
		_spec.SetField(deviceofflinerecord.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if nodes := dorc.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deviceofflinerecord.DeviceTable,
			Columns: []string{deviceofflinerecord.DeviceColumn},
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
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeviceOfflineRecord.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceOfflineRecordUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (dorc *DeviceOfflineRecordCreate) OnConflict(opts ...sql.ConflictOption) *DeviceOfflineRecordUpsertOne {
	dorc.conflict = opts
	return &DeviceOfflineRecordUpsertOne{
		create: dorc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceOfflineRecord.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dorc *DeviceOfflineRecordCreate) OnConflictColumns(columns ...string) *DeviceOfflineRecordUpsertOne {
	dorc.conflict = append(dorc.conflict, sql.ConflictColumns(columns...))
	return &DeviceOfflineRecordUpsertOne{
		create: dorc,
	}
}

type (
	// DeviceOfflineRecordUpsertOne is the builder for "upsert"-ing
	//  one DeviceOfflineRecord node.
	DeviceOfflineRecordUpsertOne struct {
		create *DeviceOfflineRecordCreate
	}

	// DeviceOfflineRecordUpsert is the "OnConflict" setter.
	DeviceOfflineRecordUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *DeviceOfflineRecordUpsert) SetCreatedBy(v int64) *DeviceOfflineRecordUpsert {
	u.Set(deviceofflinerecord.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *DeviceOfflineRecordUpsert) UpdateCreatedBy() *DeviceOfflineRecordUpsert {
	u.SetExcluded(deviceofflinerecord.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *DeviceOfflineRecordUpsert) AddCreatedBy(v int64) *DeviceOfflineRecordUpsert {
	u.Add(deviceofflinerecord.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *DeviceOfflineRecordUpsert) SetUpdatedBy(v int64) *DeviceOfflineRecordUpsert {
	u.Set(deviceofflinerecord.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *DeviceOfflineRecordUpsert) UpdateUpdatedBy() *DeviceOfflineRecordUpsert {
	u.SetExcluded(deviceofflinerecord.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *DeviceOfflineRecordUpsert) AddUpdatedBy(v int64) *DeviceOfflineRecordUpsert {
	u.Add(deviceofflinerecord.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceOfflineRecordUpsert) SetUpdatedAt(v time.Time) *DeviceOfflineRecordUpsert {
	u.Set(deviceofflinerecord.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceOfflineRecordUpsert) UpdateUpdatedAt() *DeviceOfflineRecordUpsert {
	u.SetExcluded(deviceofflinerecord.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceOfflineRecordUpsert) SetDeletedAt(v time.Time) *DeviceOfflineRecordUpsert {
	u.Set(deviceofflinerecord.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceOfflineRecordUpsert) UpdateDeletedAt() *DeviceOfflineRecordUpsert {
	u.SetExcluded(deviceofflinerecord.FieldDeletedAt)
	return u
}

// SetDeviceID sets the "device_id" field.
func (u *DeviceOfflineRecordUpsert) SetDeviceID(v int64) *DeviceOfflineRecordUpsert {
	u.Set(deviceofflinerecord.FieldDeviceID, v)
	return u
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *DeviceOfflineRecordUpsert) UpdateDeviceID() *DeviceOfflineRecordUpsert {
	u.SetExcluded(deviceofflinerecord.FieldDeviceID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DeviceOfflineRecord.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deviceofflinerecord.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceOfflineRecordUpsertOne) UpdateNewValues() *DeviceOfflineRecordUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(deviceofflinerecord.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(deviceofflinerecord.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeviceOfflineRecord.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DeviceOfflineRecordUpsertOne) Ignore() *DeviceOfflineRecordUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceOfflineRecordUpsertOne) DoNothing() *DeviceOfflineRecordUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceOfflineRecordCreate.OnConflict
// documentation for more info.
func (u *DeviceOfflineRecordUpsertOne) Update(set func(*DeviceOfflineRecordUpsert)) *DeviceOfflineRecordUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceOfflineRecordUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *DeviceOfflineRecordUpsertOne) SetCreatedBy(v int64) *DeviceOfflineRecordUpsertOne {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *DeviceOfflineRecordUpsertOne) AddCreatedBy(v int64) *DeviceOfflineRecordUpsertOne {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *DeviceOfflineRecordUpsertOne) UpdateCreatedBy() *DeviceOfflineRecordUpsertOne {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *DeviceOfflineRecordUpsertOne) SetUpdatedBy(v int64) *DeviceOfflineRecordUpsertOne {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *DeviceOfflineRecordUpsertOne) AddUpdatedBy(v int64) *DeviceOfflineRecordUpsertOne {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *DeviceOfflineRecordUpsertOne) UpdateUpdatedBy() *DeviceOfflineRecordUpsertOne {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceOfflineRecordUpsertOne) SetUpdatedAt(v time.Time) *DeviceOfflineRecordUpsertOne {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceOfflineRecordUpsertOne) UpdateUpdatedAt() *DeviceOfflineRecordUpsertOne {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceOfflineRecordUpsertOne) SetDeletedAt(v time.Time) *DeviceOfflineRecordUpsertOne {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceOfflineRecordUpsertOne) UpdateDeletedAt() *DeviceOfflineRecordUpsertOne {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *DeviceOfflineRecordUpsertOne) SetDeviceID(v int64) *DeviceOfflineRecordUpsertOne {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.SetDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *DeviceOfflineRecordUpsertOne) UpdateDeviceID() *DeviceOfflineRecordUpsertOne {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.UpdateDeviceID()
	})
}

// Exec executes the query.
func (u *DeviceOfflineRecordUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for DeviceOfflineRecordCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceOfflineRecordUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeviceOfflineRecordUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeviceOfflineRecordUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeviceOfflineRecordCreateBulk is the builder for creating many DeviceOfflineRecord entities in bulk.
type DeviceOfflineRecordCreateBulk struct {
	config
	err      error
	builders []*DeviceOfflineRecordCreate
	conflict []sql.ConflictOption
}

// Save creates the DeviceOfflineRecord entities in the database.
func (dorcb *DeviceOfflineRecordCreateBulk) Save(ctx context.Context) ([]*DeviceOfflineRecord, error) {
	if dorcb.err != nil {
		return nil, dorcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dorcb.builders))
	nodes := make([]*DeviceOfflineRecord, len(dorcb.builders))
	mutators := make([]Mutator, len(dorcb.builders))
	for i := range dorcb.builders {
		func(i int, root context.Context) {
			builder := dorcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceOfflineRecordMutation)
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
					_, err = mutators[i+1].Mutate(root, dorcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dorcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dorcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dorcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dorcb *DeviceOfflineRecordCreateBulk) SaveX(ctx context.Context) []*DeviceOfflineRecord {
	v, err := dorcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dorcb *DeviceOfflineRecordCreateBulk) Exec(ctx context.Context) error {
	_, err := dorcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dorcb *DeviceOfflineRecordCreateBulk) ExecX(ctx context.Context) {
	if err := dorcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeviceOfflineRecord.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceOfflineRecordUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (dorcb *DeviceOfflineRecordCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeviceOfflineRecordUpsertBulk {
	dorcb.conflict = opts
	return &DeviceOfflineRecordUpsertBulk{
		create: dorcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceOfflineRecord.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dorcb *DeviceOfflineRecordCreateBulk) OnConflictColumns(columns ...string) *DeviceOfflineRecordUpsertBulk {
	dorcb.conflict = append(dorcb.conflict, sql.ConflictColumns(columns...))
	return &DeviceOfflineRecordUpsertBulk{
		create: dorcb,
	}
}

// DeviceOfflineRecordUpsertBulk is the builder for "upsert"-ing
// a bulk of DeviceOfflineRecord nodes.
type DeviceOfflineRecordUpsertBulk struct {
	create *DeviceOfflineRecordCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DeviceOfflineRecord.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deviceofflinerecord.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceOfflineRecordUpsertBulk) UpdateNewValues() *DeviceOfflineRecordUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(deviceofflinerecord.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(deviceofflinerecord.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeviceOfflineRecord.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DeviceOfflineRecordUpsertBulk) Ignore() *DeviceOfflineRecordUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceOfflineRecordUpsertBulk) DoNothing() *DeviceOfflineRecordUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceOfflineRecordCreateBulk.OnConflict
// documentation for more info.
func (u *DeviceOfflineRecordUpsertBulk) Update(set func(*DeviceOfflineRecordUpsert)) *DeviceOfflineRecordUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceOfflineRecordUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *DeviceOfflineRecordUpsertBulk) SetCreatedBy(v int64) *DeviceOfflineRecordUpsertBulk {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *DeviceOfflineRecordUpsertBulk) AddCreatedBy(v int64) *DeviceOfflineRecordUpsertBulk {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *DeviceOfflineRecordUpsertBulk) UpdateCreatedBy() *DeviceOfflineRecordUpsertBulk {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *DeviceOfflineRecordUpsertBulk) SetUpdatedBy(v int64) *DeviceOfflineRecordUpsertBulk {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *DeviceOfflineRecordUpsertBulk) AddUpdatedBy(v int64) *DeviceOfflineRecordUpsertBulk {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *DeviceOfflineRecordUpsertBulk) UpdateUpdatedBy() *DeviceOfflineRecordUpsertBulk {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceOfflineRecordUpsertBulk) SetUpdatedAt(v time.Time) *DeviceOfflineRecordUpsertBulk {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceOfflineRecordUpsertBulk) UpdateUpdatedAt() *DeviceOfflineRecordUpsertBulk {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceOfflineRecordUpsertBulk) SetDeletedAt(v time.Time) *DeviceOfflineRecordUpsertBulk {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceOfflineRecordUpsertBulk) UpdateDeletedAt() *DeviceOfflineRecordUpsertBulk {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *DeviceOfflineRecordUpsertBulk) SetDeviceID(v int64) *DeviceOfflineRecordUpsertBulk {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.SetDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *DeviceOfflineRecordUpsertBulk) UpdateDeviceID() *DeviceOfflineRecordUpsertBulk {
	return u.Update(func(s *DeviceOfflineRecordUpsert) {
		s.UpdateDeviceID()
	})
}

// Exec executes the query.
func (u *DeviceOfflineRecordUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("cep_ent: OnConflict was set for builder %d. Set it on the DeviceOfflineRecordCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for DeviceOfflineRecordCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceOfflineRecordUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
