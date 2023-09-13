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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/frpcinfo"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/frpsinfo"
)

// FrpsInfoCreate is the builder for creating a FrpsInfo entity.
type FrpsInfoCreate struct {
	config
	mutation *FrpsInfoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (fic *FrpsInfoCreate) SetCreatedBy(i int64) *FrpsInfoCreate {
	fic.mutation.SetCreatedBy(i)
	return fic
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (fic *FrpsInfoCreate) SetNillableCreatedBy(i *int64) *FrpsInfoCreate {
	if i != nil {
		fic.SetCreatedBy(*i)
	}
	return fic
}

// SetUpdatedBy sets the "updated_by" field.
func (fic *FrpsInfoCreate) SetUpdatedBy(i int64) *FrpsInfoCreate {
	fic.mutation.SetUpdatedBy(i)
	return fic
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fic *FrpsInfoCreate) SetNillableUpdatedBy(i *int64) *FrpsInfoCreate {
	if i != nil {
		fic.SetUpdatedBy(*i)
	}
	return fic
}

// SetCreatedAt sets the "created_at" field.
func (fic *FrpsInfoCreate) SetCreatedAt(t time.Time) *FrpsInfoCreate {
	fic.mutation.SetCreatedAt(t)
	return fic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fic *FrpsInfoCreate) SetNillableCreatedAt(t *time.Time) *FrpsInfoCreate {
	if t != nil {
		fic.SetCreatedAt(*t)
	}
	return fic
}

// SetUpdatedAt sets the "updated_at" field.
func (fic *FrpsInfoCreate) SetUpdatedAt(t time.Time) *FrpsInfoCreate {
	fic.mutation.SetUpdatedAt(t)
	return fic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fic *FrpsInfoCreate) SetNillableUpdatedAt(t *time.Time) *FrpsInfoCreate {
	if t != nil {
		fic.SetUpdatedAt(*t)
	}
	return fic
}

// SetDeletedAt sets the "deleted_at" field.
func (fic *FrpsInfoCreate) SetDeletedAt(t time.Time) *FrpsInfoCreate {
	fic.mutation.SetDeletedAt(t)
	return fic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fic *FrpsInfoCreate) SetNillableDeletedAt(t *time.Time) *FrpsInfoCreate {
	if t != nil {
		fic.SetDeletedAt(*t)
	}
	return fic
}

// SetTag sets the "tag" field.
func (fic *FrpsInfoCreate) SetTag(s string) *FrpsInfoCreate {
	fic.mutation.SetTag(s)
	return fic
}

// SetNillableTag sets the "tag" field if the given value is not nil.
func (fic *FrpsInfoCreate) SetNillableTag(s *string) *FrpsInfoCreate {
	if s != nil {
		fic.SetTag(*s)
	}
	return fic
}

// SetServerAddr sets the "server_addr" field.
func (fic *FrpsInfoCreate) SetServerAddr(s string) *FrpsInfoCreate {
	fic.mutation.SetServerAddr(s)
	return fic
}

// SetNillableServerAddr sets the "server_addr" field if the given value is not nil.
func (fic *FrpsInfoCreate) SetNillableServerAddr(s *string) *FrpsInfoCreate {
	if s != nil {
		fic.SetServerAddr(*s)
	}
	return fic
}

// SetServerPort sets the "server_port" field.
func (fic *FrpsInfoCreate) SetServerPort(i int) *FrpsInfoCreate {
	fic.mutation.SetServerPort(i)
	return fic
}

// SetNillableServerPort sets the "server_port" field if the given value is not nil.
func (fic *FrpsInfoCreate) SetNillableServerPort(i *int) *FrpsInfoCreate {
	if i != nil {
		fic.SetServerPort(*i)
	}
	return fic
}

// SetAuthenticationMethod sets the "authentication_method" field.
func (fic *FrpsInfoCreate) SetAuthenticationMethod(s string) *FrpsInfoCreate {
	fic.mutation.SetAuthenticationMethod(s)
	return fic
}

// SetNillableAuthenticationMethod sets the "authentication_method" field if the given value is not nil.
func (fic *FrpsInfoCreate) SetNillableAuthenticationMethod(s *string) *FrpsInfoCreate {
	if s != nil {
		fic.SetAuthenticationMethod(*s)
	}
	return fic
}

// SetToken sets the "token" field.
func (fic *FrpsInfoCreate) SetToken(s string) *FrpsInfoCreate {
	fic.mutation.SetToken(s)
	return fic
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (fic *FrpsInfoCreate) SetNillableToken(s *string) *FrpsInfoCreate {
	if s != nil {
		fic.SetToken(*s)
	}
	return fic
}

// SetType sets the "type" field.
func (fic *FrpsInfoCreate) SetType(s string) *FrpsInfoCreate {
	fic.mutation.SetType(s)
	return fic
}

// SetNillableType sets the "type" field if the given value is not nil.
func (fic *FrpsInfoCreate) SetNillableType(s *string) *FrpsInfoCreate {
	if s != nil {
		fic.SetType(*s)
	}
	return fic
}

// SetID sets the "id" field.
func (fic *FrpsInfoCreate) SetID(i int64) *FrpsInfoCreate {
	fic.mutation.SetID(i)
	return fic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fic *FrpsInfoCreate) SetNillableID(i *int64) *FrpsInfoCreate {
	if i != nil {
		fic.SetID(*i)
	}
	return fic
}

// AddFrpcInfoIDs adds the "frpc_infos" edge to the FrpcInfo entity by IDs.
func (fic *FrpsInfoCreate) AddFrpcInfoIDs(ids ...int64) *FrpsInfoCreate {
	fic.mutation.AddFrpcInfoIDs(ids...)
	return fic
}

// AddFrpcInfos adds the "frpc_infos" edges to the FrpcInfo entity.
func (fic *FrpsInfoCreate) AddFrpcInfos(f ...*FrpcInfo) *FrpsInfoCreate {
	ids := make([]int64, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fic.AddFrpcInfoIDs(ids...)
}

// Mutation returns the FrpsInfoMutation object of the builder.
func (fic *FrpsInfoCreate) Mutation() *FrpsInfoMutation {
	return fic.mutation
}

// Save creates the FrpsInfo in the database.
func (fic *FrpsInfoCreate) Save(ctx context.Context) (*FrpsInfo, error) {
	fic.defaults()
	return withHooks(ctx, fic.sqlSave, fic.mutation, fic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fic *FrpsInfoCreate) SaveX(ctx context.Context) *FrpsInfo {
	v, err := fic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fic *FrpsInfoCreate) Exec(ctx context.Context) error {
	_, err := fic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fic *FrpsInfoCreate) ExecX(ctx context.Context) {
	if err := fic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fic *FrpsInfoCreate) defaults() {
	if _, ok := fic.mutation.CreatedBy(); !ok {
		v := frpsinfo.DefaultCreatedBy
		fic.mutation.SetCreatedBy(v)
	}
	if _, ok := fic.mutation.UpdatedBy(); !ok {
		v := frpsinfo.DefaultUpdatedBy
		fic.mutation.SetUpdatedBy(v)
	}
	if _, ok := fic.mutation.CreatedAt(); !ok {
		v := frpsinfo.DefaultCreatedAt()
		fic.mutation.SetCreatedAt(v)
	}
	if _, ok := fic.mutation.UpdatedAt(); !ok {
		v := frpsinfo.DefaultUpdatedAt()
		fic.mutation.SetUpdatedAt(v)
	}
	if _, ok := fic.mutation.DeletedAt(); !ok {
		v := frpsinfo.DefaultDeletedAt
		fic.mutation.SetDeletedAt(v)
	}
	if _, ok := fic.mutation.Tag(); !ok {
		v := frpsinfo.DefaultTag
		fic.mutation.SetTag(v)
	}
	if _, ok := fic.mutation.ServerAddr(); !ok {
		v := frpsinfo.DefaultServerAddr
		fic.mutation.SetServerAddr(v)
	}
	if _, ok := fic.mutation.ServerPort(); !ok {
		v := frpsinfo.DefaultServerPort
		fic.mutation.SetServerPort(v)
	}
	if _, ok := fic.mutation.AuthenticationMethod(); !ok {
		v := frpsinfo.DefaultAuthenticationMethod
		fic.mutation.SetAuthenticationMethod(v)
	}
	if _, ok := fic.mutation.Token(); !ok {
		v := frpsinfo.DefaultToken
		fic.mutation.SetToken(v)
	}
	if _, ok := fic.mutation.GetType(); !ok {
		v := frpsinfo.DefaultType
		fic.mutation.SetType(v)
	}
	if _, ok := fic.mutation.ID(); !ok {
		v := frpsinfo.DefaultID()
		fic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fic *FrpsInfoCreate) check() error {
	if _, ok := fic.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "FrpsInfo.created_by"`)}
	}
	if _, ok := fic.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "FrpsInfo.updated_by"`)}
	}
	if _, ok := fic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "FrpsInfo.created_at"`)}
	}
	if _, ok := fic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "FrpsInfo.updated_at"`)}
	}
	if _, ok := fic.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "FrpsInfo.deleted_at"`)}
	}
	if _, ok := fic.mutation.Tag(); !ok {
		return &ValidationError{Name: "tag", err: errors.New(`cep_ent: missing required field "FrpsInfo.tag"`)}
	}
	if _, ok := fic.mutation.ServerAddr(); !ok {
		return &ValidationError{Name: "server_addr", err: errors.New(`cep_ent: missing required field "FrpsInfo.server_addr"`)}
	}
	if _, ok := fic.mutation.ServerPort(); !ok {
		return &ValidationError{Name: "server_port", err: errors.New(`cep_ent: missing required field "FrpsInfo.server_port"`)}
	}
	if _, ok := fic.mutation.AuthenticationMethod(); !ok {
		return &ValidationError{Name: "authentication_method", err: errors.New(`cep_ent: missing required field "FrpsInfo.authentication_method"`)}
	}
	if _, ok := fic.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`cep_ent: missing required field "FrpsInfo.token"`)}
	}
	if _, ok := fic.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`cep_ent: missing required field "FrpsInfo.type"`)}
	}
	return nil
}

func (fic *FrpsInfoCreate) sqlSave(ctx context.Context) (*FrpsInfo, error) {
	if err := fic.check(); err != nil {
		return nil, err
	}
	_node, _spec := fic.createSpec()
	if err := sqlgraph.CreateNode(ctx, fic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	fic.mutation.id = &_node.ID
	fic.mutation.done = true
	return _node, nil
}

func (fic *FrpsInfoCreate) createSpec() (*FrpsInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &FrpsInfo{config: fic.config}
		_spec = sqlgraph.NewCreateSpec(frpsinfo.Table, sqlgraph.NewFieldSpec(frpsinfo.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = fic.conflict
	if id, ok := fic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fic.mutation.CreatedBy(); ok {
		_spec.SetField(frpsinfo.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := fic.mutation.UpdatedBy(); ok {
		_spec.SetField(frpsinfo.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := fic.mutation.CreatedAt(); ok {
		_spec.SetField(frpsinfo.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fic.mutation.UpdatedAt(); ok {
		_spec.SetField(frpsinfo.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fic.mutation.DeletedAt(); ok {
		_spec.SetField(frpsinfo.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := fic.mutation.Tag(); ok {
		_spec.SetField(frpsinfo.FieldTag, field.TypeString, value)
		_node.Tag = value
	}
	if value, ok := fic.mutation.ServerAddr(); ok {
		_spec.SetField(frpsinfo.FieldServerAddr, field.TypeString, value)
		_node.ServerAddr = value
	}
	if value, ok := fic.mutation.ServerPort(); ok {
		_spec.SetField(frpsinfo.FieldServerPort, field.TypeInt, value)
		_node.ServerPort = value
	}
	if value, ok := fic.mutation.AuthenticationMethod(); ok {
		_spec.SetField(frpsinfo.FieldAuthenticationMethod, field.TypeString, value)
		_node.AuthenticationMethod = value
	}
	if value, ok := fic.mutation.Token(); ok {
		_spec.SetField(frpsinfo.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := fic.mutation.GetType(); ok {
		_spec.SetField(frpsinfo.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if nodes := fic.mutation.FrpcInfosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   frpsinfo.FrpcInfosTable,
			Columns: []string{frpsinfo.FrpcInfosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(frpcinfo.FieldID, field.TypeInt64),
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
//	client.FrpsInfo.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FrpsInfoUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (fic *FrpsInfoCreate) OnConflict(opts ...sql.ConflictOption) *FrpsInfoUpsertOne {
	fic.conflict = opts
	return &FrpsInfoUpsertOne{
		create: fic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FrpsInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fic *FrpsInfoCreate) OnConflictColumns(columns ...string) *FrpsInfoUpsertOne {
	fic.conflict = append(fic.conflict, sql.ConflictColumns(columns...))
	return &FrpsInfoUpsertOne{
		create: fic,
	}
}

type (
	// FrpsInfoUpsertOne is the builder for "upsert"-ing
	//  one FrpsInfo node.
	FrpsInfoUpsertOne struct {
		create *FrpsInfoCreate
	}

	// FrpsInfoUpsert is the "OnConflict" setter.
	FrpsInfoUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *FrpsInfoUpsert) SetCreatedBy(v int64) *FrpsInfoUpsert {
	u.Set(frpsinfo.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *FrpsInfoUpsert) UpdateCreatedBy() *FrpsInfoUpsert {
	u.SetExcluded(frpsinfo.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *FrpsInfoUpsert) AddCreatedBy(v int64) *FrpsInfoUpsert {
	u.Add(frpsinfo.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *FrpsInfoUpsert) SetUpdatedBy(v int64) *FrpsInfoUpsert {
	u.Set(frpsinfo.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *FrpsInfoUpsert) UpdateUpdatedBy() *FrpsInfoUpsert {
	u.SetExcluded(frpsinfo.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *FrpsInfoUpsert) AddUpdatedBy(v int64) *FrpsInfoUpsert {
	u.Add(frpsinfo.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FrpsInfoUpsert) SetUpdatedAt(v time.Time) *FrpsInfoUpsert {
	u.Set(frpsinfo.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FrpsInfoUpsert) UpdateUpdatedAt() *FrpsInfoUpsert {
	u.SetExcluded(frpsinfo.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FrpsInfoUpsert) SetDeletedAt(v time.Time) *FrpsInfoUpsert {
	u.Set(frpsinfo.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FrpsInfoUpsert) UpdateDeletedAt() *FrpsInfoUpsert {
	u.SetExcluded(frpsinfo.FieldDeletedAt)
	return u
}

// SetTag sets the "tag" field.
func (u *FrpsInfoUpsert) SetTag(v string) *FrpsInfoUpsert {
	u.Set(frpsinfo.FieldTag, v)
	return u
}

// UpdateTag sets the "tag" field to the value that was provided on create.
func (u *FrpsInfoUpsert) UpdateTag() *FrpsInfoUpsert {
	u.SetExcluded(frpsinfo.FieldTag)
	return u
}

// SetServerAddr sets the "server_addr" field.
func (u *FrpsInfoUpsert) SetServerAddr(v string) *FrpsInfoUpsert {
	u.Set(frpsinfo.FieldServerAddr, v)
	return u
}

// UpdateServerAddr sets the "server_addr" field to the value that was provided on create.
func (u *FrpsInfoUpsert) UpdateServerAddr() *FrpsInfoUpsert {
	u.SetExcluded(frpsinfo.FieldServerAddr)
	return u
}

// SetServerPort sets the "server_port" field.
func (u *FrpsInfoUpsert) SetServerPort(v int) *FrpsInfoUpsert {
	u.Set(frpsinfo.FieldServerPort, v)
	return u
}

// UpdateServerPort sets the "server_port" field to the value that was provided on create.
func (u *FrpsInfoUpsert) UpdateServerPort() *FrpsInfoUpsert {
	u.SetExcluded(frpsinfo.FieldServerPort)
	return u
}

// AddServerPort adds v to the "server_port" field.
func (u *FrpsInfoUpsert) AddServerPort(v int) *FrpsInfoUpsert {
	u.Add(frpsinfo.FieldServerPort, v)
	return u
}

// SetAuthenticationMethod sets the "authentication_method" field.
func (u *FrpsInfoUpsert) SetAuthenticationMethod(v string) *FrpsInfoUpsert {
	u.Set(frpsinfo.FieldAuthenticationMethod, v)
	return u
}

// UpdateAuthenticationMethod sets the "authentication_method" field to the value that was provided on create.
func (u *FrpsInfoUpsert) UpdateAuthenticationMethod() *FrpsInfoUpsert {
	u.SetExcluded(frpsinfo.FieldAuthenticationMethod)
	return u
}

// SetToken sets the "token" field.
func (u *FrpsInfoUpsert) SetToken(v string) *FrpsInfoUpsert {
	u.Set(frpsinfo.FieldToken, v)
	return u
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *FrpsInfoUpsert) UpdateToken() *FrpsInfoUpsert {
	u.SetExcluded(frpsinfo.FieldToken)
	return u
}

// SetType sets the "type" field.
func (u *FrpsInfoUpsert) SetType(v string) *FrpsInfoUpsert {
	u.Set(frpsinfo.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *FrpsInfoUpsert) UpdateType() *FrpsInfoUpsert {
	u.SetExcluded(frpsinfo.FieldType)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.FrpsInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(frpsinfo.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FrpsInfoUpsertOne) UpdateNewValues() *FrpsInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(frpsinfo.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(frpsinfo.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FrpsInfo.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FrpsInfoUpsertOne) Ignore() *FrpsInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FrpsInfoUpsertOne) DoNothing() *FrpsInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FrpsInfoCreate.OnConflict
// documentation for more info.
func (u *FrpsInfoUpsertOne) Update(set func(*FrpsInfoUpsert)) *FrpsInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FrpsInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *FrpsInfoUpsertOne) SetCreatedBy(v int64) *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *FrpsInfoUpsertOne) AddCreatedBy(v int64) *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *FrpsInfoUpsertOne) UpdateCreatedBy() *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *FrpsInfoUpsertOne) SetUpdatedBy(v int64) *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *FrpsInfoUpsertOne) AddUpdatedBy(v int64) *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *FrpsInfoUpsertOne) UpdateUpdatedBy() *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FrpsInfoUpsertOne) SetUpdatedAt(v time.Time) *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FrpsInfoUpsertOne) UpdateUpdatedAt() *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FrpsInfoUpsertOne) SetDeletedAt(v time.Time) *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FrpsInfoUpsertOne) UpdateDeletedAt() *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetTag sets the "tag" field.
func (u *FrpsInfoUpsertOne) SetTag(v string) *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetTag(v)
	})
}

// UpdateTag sets the "tag" field to the value that was provided on create.
func (u *FrpsInfoUpsertOne) UpdateTag() *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateTag()
	})
}

// SetServerAddr sets the "server_addr" field.
func (u *FrpsInfoUpsertOne) SetServerAddr(v string) *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetServerAddr(v)
	})
}

// UpdateServerAddr sets the "server_addr" field to the value that was provided on create.
func (u *FrpsInfoUpsertOne) UpdateServerAddr() *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateServerAddr()
	})
}

// SetServerPort sets the "server_port" field.
func (u *FrpsInfoUpsertOne) SetServerPort(v int) *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetServerPort(v)
	})
}

// AddServerPort adds v to the "server_port" field.
func (u *FrpsInfoUpsertOne) AddServerPort(v int) *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.AddServerPort(v)
	})
}

// UpdateServerPort sets the "server_port" field to the value that was provided on create.
func (u *FrpsInfoUpsertOne) UpdateServerPort() *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateServerPort()
	})
}

// SetAuthenticationMethod sets the "authentication_method" field.
func (u *FrpsInfoUpsertOne) SetAuthenticationMethod(v string) *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetAuthenticationMethod(v)
	})
}

// UpdateAuthenticationMethod sets the "authentication_method" field to the value that was provided on create.
func (u *FrpsInfoUpsertOne) UpdateAuthenticationMethod() *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateAuthenticationMethod()
	})
}

// SetToken sets the "token" field.
func (u *FrpsInfoUpsertOne) SetToken(v string) *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *FrpsInfoUpsertOne) UpdateToken() *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateToken()
	})
}

// SetType sets the "type" field.
func (u *FrpsInfoUpsertOne) SetType(v string) *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *FrpsInfoUpsertOne) UpdateType() *FrpsInfoUpsertOne {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateType()
	})
}

// Exec executes the query.
func (u *FrpsInfoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for FrpsInfoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FrpsInfoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FrpsInfoUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FrpsInfoUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FrpsInfoCreateBulk is the builder for creating many FrpsInfo entities in bulk.
type FrpsInfoCreateBulk struct {
	config
	err      error
	builders []*FrpsInfoCreate
	conflict []sql.ConflictOption
}

// Save creates the FrpsInfo entities in the database.
func (ficb *FrpsInfoCreateBulk) Save(ctx context.Context) ([]*FrpsInfo, error) {
	if ficb.err != nil {
		return nil, ficb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ficb.builders))
	nodes := make([]*FrpsInfo, len(ficb.builders))
	mutators := make([]Mutator, len(ficb.builders))
	for i := range ficb.builders {
		func(i int, root context.Context) {
			builder := ficb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FrpsInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, ficb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ficb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ficb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ficb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ficb *FrpsInfoCreateBulk) SaveX(ctx context.Context) []*FrpsInfo {
	v, err := ficb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ficb *FrpsInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := ficb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ficb *FrpsInfoCreateBulk) ExecX(ctx context.Context) {
	if err := ficb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FrpsInfo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FrpsInfoUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (ficb *FrpsInfoCreateBulk) OnConflict(opts ...sql.ConflictOption) *FrpsInfoUpsertBulk {
	ficb.conflict = opts
	return &FrpsInfoUpsertBulk{
		create: ficb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FrpsInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ficb *FrpsInfoCreateBulk) OnConflictColumns(columns ...string) *FrpsInfoUpsertBulk {
	ficb.conflict = append(ficb.conflict, sql.ConflictColumns(columns...))
	return &FrpsInfoUpsertBulk{
		create: ficb,
	}
}

// FrpsInfoUpsertBulk is the builder for "upsert"-ing
// a bulk of FrpsInfo nodes.
type FrpsInfoUpsertBulk struct {
	create *FrpsInfoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.FrpsInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(frpsinfo.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FrpsInfoUpsertBulk) UpdateNewValues() *FrpsInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(frpsinfo.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(frpsinfo.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FrpsInfo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FrpsInfoUpsertBulk) Ignore() *FrpsInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FrpsInfoUpsertBulk) DoNothing() *FrpsInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FrpsInfoCreateBulk.OnConflict
// documentation for more info.
func (u *FrpsInfoUpsertBulk) Update(set func(*FrpsInfoUpsert)) *FrpsInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FrpsInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *FrpsInfoUpsertBulk) SetCreatedBy(v int64) *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *FrpsInfoUpsertBulk) AddCreatedBy(v int64) *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *FrpsInfoUpsertBulk) UpdateCreatedBy() *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *FrpsInfoUpsertBulk) SetUpdatedBy(v int64) *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *FrpsInfoUpsertBulk) AddUpdatedBy(v int64) *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *FrpsInfoUpsertBulk) UpdateUpdatedBy() *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FrpsInfoUpsertBulk) SetUpdatedAt(v time.Time) *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FrpsInfoUpsertBulk) UpdateUpdatedAt() *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FrpsInfoUpsertBulk) SetDeletedAt(v time.Time) *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FrpsInfoUpsertBulk) UpdateDeletedAt() *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetTag sets the "tag" field.
func (u *FrpsInfoUpsertBulk) SetTag(v string) *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetTag(v)
	})
}

// UpdateTag sets the "tag" field to the value that was provided on create.
func (u *FrpsInfoUpsertBulk) UpdateTag() *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateTag()
	})
}

// SetServerAddr sets the "server_addr" field.
func (u *FrpsInfoUpsertBulk) SetServerAddr(v string) *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetServerAddr(v)
	})
}

// UpdateServerAddr sets the "server_addr" field to the value that was provided on create.
func (u *FrpsInfoUpsertBulk) UpdateServerAddr() *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateServerAddr()
	})
}

// SetServerPort sets the "server_port" field.
func (u *FrpsInfoUpsertBulk) SetServerPort(v int) *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetServerPort(v)
	})
}

// AddServerPort adds v to the "server_port" field.
func (u *FrpsInfoUpsertBulk) AddServerPort(v int) *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.AddServerPort(v)
	})
}

// UpdateServerPort sets the "server_port" field to the value that was provided on create.
func (u *FrpsInfoUpsertBulk) UpdateServerPort() *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateServerPort()
	})
}

// SetAuthenticationMethod sets the "authentication_method" field.
func (u *FrpsInfoUpsertBulk) SetAuthenticationMethod(v string) *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetAuthenticationMethod(v)
	})
}

// UpdateAuthenticationMethod sets the "authentication_method" field to the value that was provided on create.
func (u *FrpsInfoUpsertBulk) UpdateAuthenticationMethod() *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateAuthenticationMethod()
	})
}

// SetToken sets the "token" field.
func (u *FrpsInfoUpsertBulk) SetToken(v string) *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *FrpsInfoUpsertBulk) UpdateToken() *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateToken()
	})
}

// SetType sets the "type" field.
func (u *FrpsInfoUpsertBulk) SetType(v string) *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *FrpsInfoUpsertBulk) UpdateType() *FrpsInfoUpsertBulk {
	return u.Update(func(s *FrpsInfoUpsert) {
		s.UpdateType()
	})
}

// Exec executes the query.
func (u *FrpsInfoUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("cep_ent: OnConflict was set for builder %d. Set it on the FrpsInfoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for FrpsInfoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FrpsInfoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}