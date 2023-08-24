// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/devicegpumission"
	"cephalon-ent/pkg/cep_ent/gpu"
	"cephalon-ent/pkg/enums"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GpuCreate is the builder for creating a Gpu entity.
type GpuCreate struct {
	config
	mutation *GpuMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (gc *GpuCreate) SetCreatedBy(i int64) *GpuCreate {
	gc.mutation.SetCreatedBy(i)
	return gc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (gc *GpuCreate) SetNillableCreatedBy(i *int64) *GpuCreate {
	if i != nil {
		gc.SetCreatedBy(*i)
	}
	return gc
}

// SetUpdatedBy sets the "updated_by" field.
func (gc *GpuCreate) SetUpdatedBy(i int64) *GpuCreate {
	gc.mutation.SetUpdatedBy(i)
	return gc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gc *GpuCreate) SetNillableUpdatedBy(i *int64) *GpuCreate {
	if i != nil {
		gc.SetUpdatedBy(*i)
	}
	return gc
}

// SetCreatedAt sets the "created_at" field.
func (gc *GpuCreate) SetCreatedAt(t time.Time) *GpuCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GpuCreate) SetNillableCreatedAt(t *time.Time) *GpuCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetUpdatedAt sets the "updated_at" field.
func (gc *GpuCreate) SetUpdatedAt(t time.Time) *GpuCreate {
	gc.mutation.SetUpdatedAt(t)
	return gc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gc *GpuCreate) SetNillableUpdatedAt(t *time.Time) *GpuCreate {
	if t != nil {
		gc.SetUpdatedAt(*t)
	}
	return gc
}

// SetDeletedAt sets the "deleted_at" field.
func (gc *GpuCreate) SetDeletedAt(t time.Time) *GpuCreate {
	gc.mutation.SetDeletedAt(t)
	return gc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gc *GpuCreate) SetNillableDeletedAt(t *time.Time) *GpuCreate {
	if t != nil {
		gc.SetDeletedAt(*t)
	}
	return gc
}

// SetVersion sets the "version" field.
func (gc *GpuCreate) SetVersion(ev enums.GpuVersion) *GpuCreate {
	gc.mutation.SetVersion(ev)
	return gc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (gc *GpuCreate) SetNillableVersion(ev *enums.GpuVersion) *GpuCreate {
	if ev != nil {
		gc.SetVersion(*ev)
	}
	return gc
}

// SetID sets the "id" field.
func (gc *GpuCreate) SetID(i int64) *GpuCreate {
	gc.mutation.SetID(i)
	return gc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gc *GpuCreate) SetNillableID(i *int64) *GpuCreate {
	if i != nil {
		gc.SetID(*i)
	}
	return gc
}

// AddDeviceGpuMissionIDs adds the "device_gpu_missions" edge to the DeviceGpuMission entity by IDs.
func (gc *GpuCreate) AddDeviceGpuMissionIDs(ids ...int64) *GpuCreate {
	gc.mutation.AddDeviceGpuMissionIDs(ids...)
	return gc
}

// AddDeviceGpuMissions adds the "device_gpu_missions" edges to the DeviceGpuMission entity.
func (gc *GpuCreate) AddDeviceGpuMissions(d ...*DeviceGpuMission) *GpuCreate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return gc.AddDeviceGpuMissionIDs(ids...)
}

// Mutation returns the GpuMutation object of the builder.
func (gc *GpuCreate) Mutation() *GpuMutation {
	return gc.mutation
}

// Save creates the Gpu in the database.
func (gc *GpuCreate) Save(ctx context.Context) (*Gpu, error) {
	gc.defaults()
	return withHooks(ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GpuCreate) SaveX(ctx context.Context) *Gpu {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GpuCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GpuCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GpuCreate) defaults() {
	if _, ok := gc.mutation.CreatedBy(); !ok {
		v := gpu.DefaultCreatedBy
		gc.mutation.SetCreatedBy(v)
	}
	if _, ok := gc.mutation.UpdatedBy(); !ok {
		v := gpu.DefaultUpdatedBy
		gc.mutation.SetUpdatedBy(v)
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		v := gpu.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		v := gpu.DefaultUpdatedAt()
		gc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gc.mutation.DeletedAt(); !ok {
		v := gpu.DefaultDeletedAt
		gc.mutation.SetDeletedAt(v)
	}
	if _, ok := gc.mutation.Version(); !ok {
		v := gpu.DefaultVersion
		gc.mutation.SetVersion(v)
	}
	if _, ok := gc.mutation.ID(); !ok {
		v := gpu.DefaultID()
		gc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GpuCreate) check() error {
	if _, ok := gc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "Gpu.created_by"`)}
	}
	if _, ok := gc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "Gpu.updated_by"`)}
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "Gpu.created_at"`)}
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "Gpu.updated_at"`)}
	}
	if _, ok := gc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "Gpu.deleted_at"`)}
	}
	if _, ok := gc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`cep_ent: missing required field "Gpu.version"`)}
	}
	if v, ok := gc.mutation.Version(); ok {
		if err := gpu.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`cep_ent: validator failed for field "Gpu.version": %w`, err)}
		}
	}
	return nil
}

func (gc *GpuCreate) sqlSave(ctx context.Context) (*Gpu, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GpuCreate) createSpec() (*Gpu, *sqlgraph.CreateSpec) {
	var (
		_node = &Gpu{config: gc.config}
		_spec = sqlgraph.NewCreateSpec(gpu.Table, sqlgraph.NewFieldSpec(gpu.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = gc.conflict
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.CreatedBy(); ok {
		_spec.SetField(gpu.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := gc.mutation.UpdatedBy(); ok {
		_spec.SetField(gpu.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.SetField(gpu.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.UpdatedAt(); ok {
		_spec.SetField(gpu.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := gc.mutation.DeletedAt(); ok {
		_spec.SetField(gpu.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := gc.mutation.Version(); ok {
		_spec.SetField(gpu.FieldVersion, field.TypeEnum, value)
		_node.Version = value
	}
	if nodes := gc.mutation.DeviceGpuMissionsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Gpu.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GpuUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (gc *GpuCreate) OnConflict(opts ...sql.ConflictOption) *GpuUpsertOne {
	gc.conflict = opts
	return &GpuUpsertOne{
		create: gc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Gpu.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gc *GpuCreate) OnConflictColumns(columns ...string) *GpuUpsertOne {
	gc.conflict = append(gc.conflict, sql.ConflictColumns(columns...))
	return &GpuUpsertOne{
		create: gc,
	}
}

type (
	// GpuUpsertOne is the builder for "upsert"-ing
	//  one Gpu node.
	GpuUpsertOne struct {
		create *GpuCreate
	}

	// GpuUpsert is the "OnConflict" setter.
	GpuUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *GpuUpsert) SetCreatedBy(v int64) *GpuUpsert {
	u.Set(gpu.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *GpuUpsert) UpdateCreatedBy() *GpuUpsert {
	u.SetExcluded(gpu.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *GpuUpsert) AddCreatedBy(v int64) *GpuUpsert {
	u.Add(gpu.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *GpuUpsert) SetUpdatedBy(v int64) *GpuUpsert {
	u.Set(gpu.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *GpuUpsert) UpdateUpdatedBy() *GpuUpsert {
	u.SetExcluded(gpu.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *GpuUpsert) AddUpdatedBy(v int64) *GpuUpsert {
	u.Add(gpu.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GpuUpsert) SetUpdatedAt(v time.Time) *GpuUpsert {
	u.Set(gpu.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GpuUpsert) UpdateUpdatedAt() *GpuUpsert {
	u.SetExcluded(gpu.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *GpuUpsert) SetDeletedAt(v time.Time) *GpuUpsert {
	u.Set(gpu.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *GpuUpsert) UpdateDeletedAt() *GpuUpsert {
	u.SetExcluded(gpu.FieldDeletedAt)
	return u
}

// SetVersion sets the "version" field.
func (u *GpuUpsert) SetVersion(v enums.GpuVersion) *GpuUpsert {
	u.Set(gpu.FieldVersion, v)
	return u
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *GpuUpsert) UpdateVersion() *GpuUpsert {
	u.SetExcluded(gpu.FieldVersion)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Gpu.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(gpu.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GpuUpsertOne) UpdateNewValues() *GpuUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(gpu.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(gpu.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Gpu.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GpuUpsertOne) Ignore() *GpuUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GpuUpsertOne) DoNothing() *GpuUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GpuCreate.OnConflict
// documentation for more info.
func (u *GpuUpsertOne) Update(set func(*GpuUpsert)) *GpuUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GpuUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *GpuUpsertOne) SetCreatedBy(v int64) *GpuUpsertOne {
	return u.Update(func(s *GpuUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *GpuUpsertOne) AddCreatedBy(v int64) *GpuUpsertOne {
	return u.Update(func(s *GpuUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *GpuUpsertOne) UpdateCreatedBy() *GpuUpsertOne {
	return u.Update(func(s *GpuUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *GpuUpsertOne) SetUpdatedBy(v int64) *GpuUpsertOne {
	return u.Update(func(s *GpuUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *GpuUpsertOne) AddUpdatedBy(v int64) *GpuUpsertOne {
	return u.Update(func(s *GpuUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *GpuUpsertOne) UpdateUpdatedBy() *GpuUpsertOne {
	return u.Update(func(s *GpuUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GpuUpsertOne) SetUpdatedAt(v time.Time) *GpuUpsertOne {
	return u.Update(func(s *GpuUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GpuUpsertOne) UpdateUpdatedAt() *GpuUpsertOne {
	return u.Update(func(s *GpuUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *GpuUpsertOne) SetDeletedAt(v time.Time) *GpuUpsertOne {
	return u.Update(func(s *GpuUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *GpuUpsertOne) UpdateDeletedAt() *GpuUpsertOne {
	return u.Update(func(s *GpuUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetVersion sets the "version" field.
func (u *GpuUpsertOne) SetVersion(v enums.GpuVersion) *GpuUpsertOne {
	return u.Update(func(s *GpuUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *GpuUpsertOne) UpdateVersion() *GpuUpsertOne {
	return u.Update(func(s *GpuUpsert) {
		s.UpdateVersion()
	})
}

// Exec executes the query.
func (u *GpuUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for GpuCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GpuUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GpuUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GpuUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GpuCreateBulk is the builder for creating many Gpu entities in bulk.
type GpuCreateBulk struct {
	config
	builders []*GpuCreate
	conflict []sql.ConflictOption
}

// Save creates the Gpu entities in the database.
func (gcb *GpuCreateBulk) Save(ctx context.Context) ([]*Gpu, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Gpu, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GpuMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GpuCreateBulk) SaveX(ctx context.Context) []*Gpu {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GpuCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GpuCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Gpu.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GpuUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (gcb *GpuCreateBulk) OnConflict(opts ...sql.ConflictOption) *GpuUpsertBulk {
	gcb.conflict = opts
	return &GpuUpsertBulk{
		create: gcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Gpu.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gcb *GpuCreateBulk) OnConflictColumns(columns ...string) *GpuUpsertBulk {
	gcb.conflict = append(gcb.conflict, sql.ConflictColumns(columns...))
	return &GpuUpsertBulk{
		create: gcb,
	}
}

// GpuUpsertBulk is the builder for "upsert"-ing
// a bulk of Gpu nodes.
type GpuUpsertBulk struct {
	create *GpuCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Gpu.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(gpu.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GpuUpsertBulk) UpdateNewValues() *GpuUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(gpu.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(gpu.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Gpu.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GpuUpsertBulk) Ignore() *GpuUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GpuUpsertBulk) DoNothing() *GpuUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GpuCreateBulk.OnConflict
// documentation for more info.
func (u *GpuUpsertBulk) Update(set func(*GpuUpsert)) *GpuUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GpuUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *GpuUpsertBulk) SetCreatedBy(v int64) *GpuUpsertBulk {
	return u.Update(func(s *GpuUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *GpuUpsertBulk) AddCreatedBy(v int64) *GpuUpsertBulk {
	return u.Update(func(s *GpuUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *GpuUpsertBulk) UpdateCreatedBy() *GpuUpsertBulk {
	return u.Update(func(s *GpuUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *GpuUpsertBulk) SetUpdatedBy(v int64) *GpuUpsertBulk {
	return u.Update(func(s *GpuUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *GpuUpsertBulk) AddUpdatedBy(v int64) *GpuUpsertBulk {
	return u.Update(func(s *GpuUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *GpuUpsertBulk) UpdateUpdatedBy() *GpuUpsertBulk {
	return u.Update(func(s *GpuUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GpuUpsertBulk) SetUpdatedAt(v time.Time) *GpuUpsertBulk {
	return u.Update(func(s *GpuUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GpuUpsertBulk) UpdateUpdatedAt() *GpuUpsertBulk {
	return u.Update(func(s *GpuUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *GpuUpsertBulk) SetDeletedAt(v time.Time) *GpuUpsertBulk {
	return u.Update(func(s *GpuUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *GpuUpsertBulk) UpdateDeletedAt() *GpuUpsertBulk {
	return u.Update(func(s *GpuUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetVersion sets the "version" field.
func (u *GpuUpsertBulk) SetVersion(v enums.GpuVersion) *GpuUpsertBulk {
	return u.Update(func(s *GpuUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *GpuUpsertBulk) UpdateVersion() *GpuUpsertBulk {
	return u.Update(func(s *GpuUpsert) {
		s.UpdateVersion()
	})
}

// Exec executes the query.
func (u *GpuUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("cep_ent: OnConflict was set for builder %d. Set it on the GpuCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for GpuCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GpuUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
