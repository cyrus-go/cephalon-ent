// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/device"
	"cephalon-ent/pkg/cep_ent/user"
	"cephalon-ent/pkg/cep_ent/userdevice"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserDeviceCreate is the builder for creating a UserDevice entity.
type UserDeviceCreate struct {
	config
	mutation *UserDeviceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (udc *UserDeviceCreate) SetCreatedBy(i int64) *UserDeviceCreate {
	udc.mutation.SetCreatedBy(i)
	return udc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableCreatedBy(i *int64) *UserDeviceCreate {
	if i != nil {
		udc.SetCreatedBy(*i)
	}
	return udc
}

// SetUpdatedBy sets the "updated_by" field.
func (udc *UserDeviceCreate) SetUpdatedBy(i int64) *UserDeviceCreate {
	udc.mutation.SetUpdatedBy(i)
	return udc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableUpdatedBy(i *int64) *UserDeviceCreate {
	if i != nil {
		udc.SetUpdatedBy(*i)
	}
	return udc
}

// SetCreatedAt sets the "created_at" field.
func (udc *UserDeviceCreate) SetCreatedAt(t time.Time) *UserDeviceCreate {
	udc.mutation.SetCreatedAt(t)
	return udc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableCreatedAt(t *time.Time) *UserDeviceCreate {
	if t != nil {
		udc.SetCreatedAt(*t)
	}
	return udc
}

// SetUpdatedAt sets the "updated_at" field.
func (udc *UserDeviceCreate) SetUpdatedAt(t time.Time) *UserDeviceCreate {
	udc.mutation.SetUpdatedAt(t)
	return udc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableUpdatedAt(t *time.Time) *UserDeviceCreate {
	if t != nil {
		udc.SetUpdatedAt(*t)
	}
	return udc
}

// SetDeletedAt sets the "deleted_at" field.
func (udc *UserDeviceCreate) SetDeletedAt(t time.Time) *UserDeviceCreate {
	udc.mutation.SetDeletedAt(t)
	return udc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableDeletedAt(t *time.Time) *UserDeviceCreate {
	if t != nil {
		udc.SetDeletedAt(*t)
	}
	return udc
}

// SetUserID sets the "user_id" field.
func (udc *UserDeviceCreate) SetUserID(i int64) *UserDeviceCreate {
	udc.mutation.SetUserID(i)
	return udc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableUserID(i *int64) *UserDeviceCreate {
	if i != nil {
		udc.SetUserID(*i)
	}
	return udc
}

// SetDeviceID sets the "device_id" field.
func (udc *UserDeviceCreate) SetDeviceID(i int64) *UserDeviceCreate {
	udc.mutation.SetDeviceID(i)
	return udc
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableDeviceID(i *int64) *UserDeviceCreate {
	if i != nil {
		udc.SetDeviceID(*i)
	}
	return udc
}

// SetID sets the "id" field.
func (udc *UserDeviceCreate) SetID(i int64) *UserDeviceCreate {
	udc.mutation.SetID(i)
	return udc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableID(i *int64) *UserDeviceCreate {
	if i != nil {
		udc.SetID(*i)
	}
	return udc
}

// SetUser sets the "user" edge to the User entity.
func (udc *UserDeviceCreate) SetUser(u *User) *UserDeviceCreate {
	return udc.SetUserID(u.ID)
}

// SetDevice sets the "device" edge to the Device entity.
func (udc *UserDeviceCreate) SetDevice(d *Device) *UserDeviceCreate {
	return udc.SetDeviceID(d.ID)
}

// Mutation returns the UserDeviceMutation object of the builder.
func (udc *UserDeviceCreate) Mutation() *UserDeviceMutation {
	return udc.mutation
}

// Save creates the UserDevice in the database.
func (udc *UserDeviceCreate) Save(ctx context.Context) (*UserDevice, error) {
	udc.defaults()
	return withHooks(ctx, udc.sqlSave, udc.mutation, udc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (udc *UserDeviceCreate) SaveX(ctx context.Context) *UserDevice {
	v, err := udc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (udc *UserDeviceCreate) Exec(ctx context.Context) error {
	_, err := udc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udc *UserDeviceCreate) ExecX(ctx context.Context) {
	if err := udc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (udc *UserDeviceCreate) defaults() {
	if _, ok := udc.mutation.CreatedBy(); !ok {
		v := userdevice.DefaultCreatedBy
		udc.mutation.SetCreatedBy(v)
	}
	if _, ok := udc.mutation.UpdatedBy(); !ok {
		v := userdevice.DefaultUpdatedBy
		udc.mutation.SetUpdatedBy(v)
	}
	if _, ok := udc.mutation.CreatedAt(); !ok {
		v := userdevice.DefaultCreatedAt()
		udc.mutation.SetCreatedAt(v)
	}
	if _, ok := udc.mutation.UpdatedAt(); !ok {
		v := userdevice.DefaultUpdatedAt()
		udc.mutation.SetUpdatedAt(v)
	}
	if _, ok := udc.mutation.DeletedAt(); !ok {
		v := userdevice.DefaultDeletedAt
		udc.mutation.SetDeletedAt(v)
	}
	if _, ok := udc.mutation.UserID(); !ok {
		v := userdevice.DefaultUserID
		udc.mutation.SetUserID(v)
	}
	if _, ok := udc.mutation.DeviceID(); !ok {
		v := userdevice.DefaultDeviceID
		udc.mutation.SetDeviceID(v)
	}
	if _, ok := udc.mutation.ID(); !ok {
		v := userdevice.DefaultID()
		udc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (udc *UserDeviceCreate) check() error {
	if _, ok := udc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "UserDevice.created_by"`)}
	}
	if _, ok := udc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "UserDevice.updated_by"`)}
	}
	if _, ok := udc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "UserDevice.created_at"`)}
	}
	if _, ok := udc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "UserDevice.updated_at"`)}
	}
	if _, ok := udc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "UserDevice.deleted_at"`)}
	}
	if _, ok := udc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`cep_ent: missing required field "UserDevice.user_id"`)}
	}
	if _, ok := udc.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device_id", err: errors.New(`cep_ent: missing required field "UserDevice.device_id"`)}
	}
	if _, ok := udc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`cep_ent: missing required edge "UserDevice.user"`)}
	}
	if _, ok := udc.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device", err: errors.New(`cep_ent: missing required edge "UserDevice.device"`)}
	}
	return nil
}

func (udc *UserDeviceCreate) sqlSave(ctx context.Context) (*UserDevice, error) {
	if err := udc.check(); err != nil {
		return nil, err
	}
	_node, _spec := udc.createSpec()
	if err := sqlgraph.CreateNode(ctx, udc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	udc.mutation.id = &_node.ID
	udc.mutation.done = true
	return _node, nil
}

func (udc *UserDeviceCreate) createSpec() (*UserDevice, *sqlgraph.CreateSpec) {
	var (
		_node = &UserDevice{config: udc.config}
		_spec = sqlgraph.NewCreateSpec(userdevice.Table, sqlgraph.NewFieldSpec(userdevice.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = udc.conflict
	if id, ok := udc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := udc.mutation.CreatedBy(); ok {
		_spec.SetField(userdevice.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := udc.mutation.UpdatedBy(); ok {
		_spec.SetField(userdevice.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := udc.mutation.CreatedAt(); ok {
		_spec.SetField(userdevice.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := udc.mutation.UpdatedAt(); ok {
		_spec.SetField(userdevice.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := udc.mutation.DeletedAt(); ok {
		_spec.SetField(userdevice.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if nodes := udc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userdevice.UserTable,
			Columns: []string{userdevice.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := udc.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userdevice.DeviceTable,
			Columns: []string{userdevice.DeviceColumn},
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
//	client.UserDevice.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserDeviceUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (udc *UserDeviceCreate) OnConflict(opts ...sql.ConflictOption) *UserDeviceUpsertOne {
	udc.conflict = opts
	return &UserDeviceUpsertOne{
		create: udc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserDevice.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (udc *UserDeviceCreate) OnConflictColumns(columns ...string) *UserDeviceUpsertOne {
	udc.conflict = append(udc.conflict, sql.ConflictColumns(columns...))
	return &UserDeviceUpsertOne{
		create: udc,
	}
}

type (
	// UserDeviceUpsertOne is the builder for "upsert"-ing
	//  one UserDevice node.
	UserDeviceUpsertOne struct {
		create *UserDeviceCreate
	}

	// UserDeviceUpsert is the "OnConflict" setter.
	UserDeviceUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *UserDeviceUpsert) SetCreatedBy(v int64) *UserDeviceUpsert {
	u.Set(userdevice.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *UserDeviceUpsert) UpdateCreatedBy() *UserDeviceUpsert {
	u.SetExcluded(userdevice.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *UserDeviceUpsert) AddCreatedBy(v int64) *UserDeviceUpsert {
	u.Add(userdevice.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *UserDeviceUpsert) SetUpdatedBy(v int64) *UserDeviceUpsert {
	u.Set(userdevice.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *UserDeviceUpsert) UpdateUpdatedBy() *UserDeviceUpsert {
	u.SetExcluded(userdevice.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *UserDeviceUpsert) AddUpdatedBy(v int64) *UserDeviceUpsert {
	u.Add(userdevice.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserDeviceUpsert) SetUpdatedAt(v time.Time) *UserDeviceUpsert {
	u.Set(userdevice.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserDeviceUpsert) UpdateUpdatedAt() *UserDeviceUpsert {
	u.SetExcluded(userdevice.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserDeviceUpsert) SetDeletedAt(v time.Time) *UserDeviceUpsert {
	u.Set(userdevice.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserDeviceUpsert) UpdateDeletedAt() *UserDeviceUpsert {
	u.SetExcluded(userdevice.FieldDeletedAt)
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserDeviceUpsert) SetUserID(v int64) *UserDeviceUpsert {
	u.Set(userdevice.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserDeviceUpsert) UpdateUserID() *UserDeviceUpsert {
	u.SetExcluded(userdevice.FieldUserID)
	return u
}

// SetDeviceID sets the "device_id" field.
func (u *UserDeviceUpsert) SetDeviceID(v int64) *UserDeviceUpsert {
	u.Set(userdevice.FieldDeviceID, v)
	return u
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *UserDeviceUpsert) UpdateDeviceID() *UserDeviceUpsert {
	u.SetExcluded(userdevice.FieldDeviceID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.UserDevice.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userdevice.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserDeviceUpsertOne) UpdateNewValues() *UserDeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(userdevice.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(userdevice.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserDevice.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserDeviceUpsertOne) Ignore() *UserDeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserDeviceUpsertOne) DoNothing() *UserDeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserDeviceCreate.OnConflict
// documentation for more info.
func (u *UserDeviceUpsertOne) Update(set func(*UserDeviceUpsert)) *UserDeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserDeviceUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *UserDeviceUpsertOne) SetCreatedBy(v int64) *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *UserDeviceUpsertOne) AddCreatedBy(v int64) *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *UserDeviceUpsertOne) UpdateCreatedBy() *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *UserDeviceUpsertOne) SetUpdatedBy(v int64) *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *UserDeviceUpsertOne) AddUpdatedBy(v int64) *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *UserDeviceUpsertOne) UpdateUpdatedBy() *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserDeviceUpsertOne) SetUpdatedAt(v time.Time) *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserDeviceUpsertOne) UpdateUpdatedAt() *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserDeviceUpsertOne) SetDeletedAt(v time.Time) *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserDeviceUpsertOne) UpdateDeletedAt() *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserDeviceUpsertOne) SetUserID(v int64) *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserDeviceUpsertOne) UpdateUserID() *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateUserID()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *UserDeviceUpsertOne) SetDeviceID(v int64) *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *UserDeviceUpsertOne) UpdateDeviceID() *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateDeviceID()
	})
}

// Exec executes the query.
func (u *UserDeviceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for UserDeviceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserDeviceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserDeviceUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserDeviceUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserDeviceCreateBulk is the builder for creating many UserDevice entities in bulk.
type UserDeviceCreateBulk struct {
	config
	builders []*UserDeviceCreate
	conflict []sql.ConflictOption
}

// Save creates the UserDevice entities in the database.
func (udcb *UserDeviceCreateBulk) Save(ctx context.Context) ([]*UserDevice, error) {
	specs := make([]*sqlgraph.CreateSpec, len(udcb.builders))
	nodes := make([]*UserDevice, len(udcb.builders))
	mutators := make([]Mutator, len(udcb.builders))
	for i := range udcb.builders {
		func(i int, root context.Context) {
			builder := udcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserDeviceMutation)
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
					_, err = mutators[i+1].Mutate(root, udcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = udcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, udcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, udcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (udcb *UserDeviceCreateBulk) SaveX(ctx context.Context) []*UserDevice {
	v, err := udcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (udcb *UserDeviceCreateBulk) Exec(ctx context.Context) error {
	_, err := udcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udcb *UserDeviceCreateBulk) ExecX(ctx context.Context) {
	if err := udcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserDevice.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserDeviceUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (udcb *UserDeviceCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserDeviceUpsertBulk {
	udcb.conflict = opts
	return &UserDeviceUpsertBulk{
		create: udcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserDevice.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (udcb *UserDeviceCreateBulk) OnConflictColumns(columns ...string) *UserDeviceUpsertBulk {
	udcb.conflict = append(udcb.conflict, sql.ConflictColumns(columns...))
	return &UserDeviceUpsertBulk{
		create: udcb,
	}
}

// UserDeviceUpsertBulk is the builder for "upsert"-ing
// a bulk of UserDevice nodes.
type UserDeviceUpsertBulk struct {
	create *UserDeviceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserDevice.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userdevice.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserDeviceUpsertBulk) UpdateNewValues() *UserDeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(userdevice.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(userdevice.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserDevice.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserDeviceUpsertBulk) Ignore() *UserDeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserDeviceUpsertBulk) DoNothing() *UserDeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserDeviceCreateBulk.OnConflict
// documentation for more info.
func (u *UserDeviceUpsertBulk) Update(set func(*UserDeviceUpsert)) *UserDeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserDeviceUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *UserDeviceUpsertBulk) SetCreatedBy(v int64) *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *UserDeviceUpsertBulk) AddCreatedBy(v int64) *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *UserDeviceUpsertBulk) UpdateCreatedBy() *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *UserDeviceUpsertBulk) SetUpdatedBy(v int64) *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *UserDeviceUpsertBulk) AddUpdatedBy(v int64) *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *UserDeviceUpsertBulk) UpdateUpdatedBy() *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserDeviceUpsertBulk) SetUpdatedAt(v time.Time) *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserDeviceUpsertBulk) UpdateUpdatedAt() *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserDeviceUpsertBulk) SetDeletedAt(v time.Time) *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserDeviceUpsertBulk) UpdateDeletedAt() *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserDeviceUpsertBulk) SetUserID(v int64) *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserDeviceUpsertBulk) UpdateUserID() *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateUserID()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *UserDeviceUpsertBulk) SetDeviceID(v int64) *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *UserDeviceUpsertBulk) UpdateDeviceID() *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateDeviceID()
	})
}

// Exec executes the query.
func (u *UserDeviceUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("cep_ent: OnConflict was set for builder %d. Set it on the UserDeviceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for UserDeviceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserDeviceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
