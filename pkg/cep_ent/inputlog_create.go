// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/inputlog"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InputLogCreate is the builder for creating a InputLog entity.
type InputLogCreate struct {
	config
	mutation *InputLogMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (ilc *InputLogCreate) SetCreatedBy(i int64) *InputLogCreate {
	ilc.mutation.SetCreatedBy(i)
	return ilc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ilc *InputLogCreate) SetNillableCreatedBy(i *int64) *InputLogCreate {
	if i != nil {
		ilc.SetCreatedBy(*i)
	}
	return ilc
}

// SetUpdatedBy sets the "updated_by" field.
func (ilc *InputLogCreate) SetUpdatedBy(i int64) *InputLogCreate {
	ilc.mutation.SetUpdatedBy(i)
	return ilc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ilc *InputLogCreate) SetNillableUpdatedBy(i *int64) *InputLogCreate {
	if i != nil {
		ilc.SetUpdatedBy(*i)
	}
	return ilc
}

// SetCreatedAt sets the "created_at" field.
func (ilc *InputLogCreate) SetCreatedAt(t time.Time) *InputLogCreate {
	ilc.mutation.SetCreatedAt(t)
	return ilc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ilc *InputLogCreate) SetNillableCreatedAt(t *time.Time) *InputLogCreate {
	if t != nil {
		ilc.SetCreatedAt(*t)
	}
	return ilc
}

// SetUpdatedAt sets the "updated_at" field.
func (ilc *InputLogCreate) SetUpdatedAt(t time.Time) *InputLogCreate {
	ilc.mutation.SetUpdatedAt(t)
	return ilc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ilc *InputLogCreate) SetNillableUpdatedAt(t *time.Time) *InputLogCreate {
	if t != nil {
		ilc.SetUpdatedAt(*t)
	}
	return ilc
}

// SetDeletedAt sets the "deleted_at" field.
func (ilc *InputLogCreate) SetDeletedAt(t time.Time) *InputLogCreate {
	ilc.mutation.SetDeletedAt(t)
	return ilc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ilc *InputLogCreate) SetNillableDeletedAt(t *time.Time) *InputLogCreate {
	if t != nil {
		ilc.SetDeletedAt(*t)
	}
	return ilc
}

// SetTraceID sets the "trace_id" field.
func (ilc *InputLogCreate) SetTraceID(i int64) *InputLogCreate {
	ilc.mutation.SetTraceID(i)
	return ilc
}

// SetHeaders sets the "headers" field.
func (ilc *InputLogCreate) SetHeaders(s string) *InputLogCreate {
	ilc.mutation.SetHeaders(s)
	return ilc
}

// SetBody sets the "body" field.
func (ilc *InputLogCreate) SetBody(s string) *InputLogCreate {
	ilc.mutation.SetBody(s)
	return ilc
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (ilc *InputLogCreate) SetNillableBody(s *string) *InputLogCreate {
	if s != nil {
		ilc.SetBody(*s)
	}
	return ilc
}

// SetQuery sets the "query" field.
func (ilc *InputLogCreate) SetQuery(s string) *InputLogCreate {
	ilc.mutation.SetQuery(s)
	return ilc
}

// SetNillableQuery sets the "query" field if the given value is not nil.
func (ilc *InputLogCreate) SetNillableQuery(s *string) *InputLogCreate {
	if s != nil {
		ilc.SetQuery(*s)
	}
	return ilc
}

// SetURL sets the "url" field.
func (ilc *InputLogCreate) SetURL(s string) *InputLogCreate {
	ilc.mutation.SetURL(s)
	return ilc
}

// SetIP sets the "ip" field.
func (ilc *InputLogCreate) SetIP(s string) *InputLogCreate {
	ilc.mutation.SetIP(s)
	return ilc
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (ilc *InputLogCreate) SetNillableIP(s *string) *InputLogCreate {
	if s != nil {
		ilc.SetIP(*s)
	}
	return ilc
}

// SetCaller sets the "caller" field.
func (ilc *InputLogCreate) SetCaller(s string) *InputLogCreate {
	ilc.mutation.SetCaller(s)
	return ilc
}

// SetNillableCaller sets the "caller" field if the given value is not nil.
func (ilc *InputLogCreate) SetNillableCaller(s *string) *InputLogCreate {
	if s != nil {
		ilc.SetCaller(*s)
	}
	return ilc
}

// SetMethod sets the "method" field.
func (ilc *InputLogCreate) SetMethod(i inputlog.Method) *InputLogCreate {
	ilc.mutation.SetMethod(i)
	return ilc
}

// SetHmacKey sets the "hmac_key" field.
func (ilc *InputLogCreate) SetHmacKey(s string) *InputLogCreate {
	ilc.mutation.SetHmacKey(s)
	return ilc
}

// SetNillableHmacKey sets the "hmac_key" field if the given value is not nil.
func (ilc *InputLogCreate) SetNillableHmacKey(s *string) *InputLogCreate {
	if s != nil {
		ilc.SetHmacKey(*s)
	}
	return ilc
}

// SetID sets the "id" field.
func (ilc *InputLogCreate) SetID(i int64) *InputLogCreate {
	ilc.mutation.SetID(i)
	return ilc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ilc *InputLogCreate) SetNillableID(i *int64) *InputLogCreate {
	if i != nil {
		ilc.SetID(*i)
	}
	return ilc
}

// Mutation returns the InputLogMutation object of the builder.
func (ilc *InputLogCreate) Mutation() *InputLogMutation {
	return ilc.mutation
}

// Save creates the InputLog in the database.
func (ilc *InputLogCreate) Save(ctx context.Context) (*InputLog, error) {
	ilc.defaults()
	return withHooks(ctx, ilc.sqlSave, ilc.mutation, ilc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ilc *InputLogCreate) SaveX(ctx context.Context) *InputLog {
	v, err := ilc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ilc *InputLogCreate) Exec(ctx context.Context) error {
	_, err := ilc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ilc *InputLogCreate) ExecX(ctx context.Context) {
	if err := ilc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ilc *InputLogCreate) defaults() {
	if _, ok := ilc.mutation.CreatedBy(); !ok {
		v := inputlog.DefaultCreatedBy
		ilc.mutation.SetCreatedBy(v)
	}
	if _, ok := ilc.mutation.UpdatedBy(); !ok {
		v := inputlog.DefaultUpdatedBy
		ilc.mutation.SetUpdatedBy(v)
	}
	if _, ok := ilc.mutation.CreatedAt(); !ok {
		v := inputlog.DefaultCreatedAt()
		ilc.mutation.SetCreatedAt(v)
	}
	if _, ok := ilc.mutation.UpdatedAt(); !ok {
		v := inputlog.DefaultUpdatedAt()
		ilc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ilc.mutation.DeletedAt(); !ok {
		v := inputlog.DefaultDeletedAt
		ilc.mutation.SetDeletedAt(v)
	}
	if _, ok := ilc.mutation.Body(); !ok {
		v := inputlog.DefaultBody
		ilc.mutation.SetBody(v)
	}
	if _, ok := ilc.mutation.Query(); !ok {
		v := inputlog.DefaultQuery
		ilc.mutation.SetQuery(v)
	}
	if _, ok := ilc.mutation.IP(); !ok {
		v := inputlog.DefaultIP
		ilc.mutation.SetIP(v)
	}
	if _, ok := ilc.mutation.Caller(); !ok {
		v := inputlog.DefaultCaller
		ilc.mutation.SetCaller(v)
	}
	if _, ok := ilc.mutation.HmacKey(); !ok {
		v := inputlog.DefaultHmacKey
		ilc.mutation.SetHmacKey(v)
	}
	if _, ok := ilc.mutation.ID(); !ok {
		v := inputlog.DefaultID()
		ilc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ilc *InputLogCreate) check() error {
	if _, ok := ilc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "InputLog.created_by"`)}
	}
	if _, ok := ilc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "InputLog.updated_by"`)}
	}
	if _, ok := ilc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "InputLog.created_at"`)}
	}
	if _, ok := ilc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "InputLog.updated_at"`)}
	}
	if _, ok := ilc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "InputLog.deleted_at"`)}
	}
	if _, ok := ilc.mutation.TraceID(); !ok {
		return &ValidationError{Name: "trace_id", err: errors.New(`cep_ent: missing required field "InputLog.trace_id"`)}
	}
	if _, ok := ilc.mutation.Headers(); !ok {
		return &ValidationError{Name: "headers", err: errors.New(`cep_ent: missing required field "InputLog.headers"`)}
	}
	if _, ok := ilc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`cep_ent: missing required field "InputLog.url"`)}
	}
	if _, ok := ilc.mutation.Caller(); !ok {
		return &ValidationError{Name: "caller", err: errors.New(`cep_ent: missing required field "InputLog.caller"`)}
	}
	if _, ok := ilc.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`cep_ent: missing required field "InputLog.method"`)}
	}
	if v, ok := ilc.mutation.Method(); ok {
		if err := inputlog.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf(`cep_ent: validator failed for field "InputLog.method": %w`, err)}
		}
	}
	if _, ok := ilc.mutation.HmacKey(); !ok {
		return &ValidationError{Name: "hmac_key", err: errors.New(`cep_ent: missing required field "InputLog.hmac_key"`)}
	}
	return nil
}

func (ilc *InputLogCreate) sqlSave(ctx context.Context) (*InputLog, error) {
	if err := ilc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ilc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ilc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ilc.mutation.id = &_node.ID
	ilc.mutation.done = true
	return _node, nil
}

func (ilc *InputLogCreate) createSpec() (*InputLog, *sqlgraph.CreateSpec) {
	var (
		_node = &InputLog{config: ilc.config}
		_spec = sqlgraph.NewCreateSpec(inputlog.Table, sqlgraph.NewFieldSpec(inputlog.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = ilc.conflict
	if id, ok := ilc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ilc.mutation.CreatedBy(); ok {
		_spec.SetField(inputlog.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := ilc.mutation.UpdatedBy(); ok {
		_spec.SetField(inputlog.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := ilc.mutation.CreatedAt(); ok {
		_spec.SetField(inputlog.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ilc.mutation.UpdatedAt(); ok {
		_spec.SetField(inputlog.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ilc.mutation.DeletedAt(); ok {
		_spec.SetField(inputlog.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ilc.mutation.TraceID(); ok {
		_spec.SetField(inputlog.FieldTraceID, field.TypeInt64, value)
		_node.TraceID = value
	}
	if value, ok := ilc.mutation.Headers(); ok {
		_spec.SetField(inputlog.FieldHeaders, field.TypeString, value)
		_node.Headers = value
	}
	if value, ok := ilc.mutation.Body(); ok {
		_spec.SetField(inputlog.FieldBody, field.TypeString, value)
		_node.Body = value
	}
	if value, ok := ilc.mutation.Query(); ok {
		_spec.SetField(inputlog.FieldQuery, field.TypeString, value)
		_node.Query = value
	}
	if value, ok := ilc.mutation.URL(); ok {
		_spec.SetField(inputlog.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := ilc.mutation.IP(); ok {
		_spec.SetField(inputlog.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if value, ok := ilc.mutation.Caller(); ok {
		_spec.SetField(inputlog.FieldCaller, field.TypeString, value)
		_node.Caller = value
	}
	if value, ok := ilc.mutation.Method(); ok {
		_spec.SetField(inputlog.FieldMethod, field.TypeEnum, value)
		_node.Method = value
	}
	if value, ok := ilc.mutation.HmacKey(); ok {
		_spec.SetField(inputlog.FieldHmacKey, field.TypeString, value)
		_node.HmacKey = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.InputLog.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.InputLogUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (ilc *InputLogCreate) OnConflict(opts ...sql.ConflictOption) *InputLogUpsertOne {
	ilc.conflict = opts
	return &InputLogUpsertOne{
		create: ilc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.InputLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ilc *InputLogCreate) OnConflictColumns(columns ...string) *InputLogUpsertOne {
	ilc.conflict = append(ilc.conflict, sql.ConflictColumns(columns...))
	return &InputLogUpsertOne{
		create: ilc,
	}
}

type (
	// InputLogUpsertOne is the builder for "upsert"-ing
	//  one InputLog node.
	InputLogUpsertOne struct {
		create *InputLogCreate
	}

	// InputLogUpsert is the "OnConflict" setter.
	InputLogUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *InputLogUpsert) SetCreatedBy(v int64) *InputLogUpsert {
	u.Set(inputlog.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *InputLogUpsert) UpdateCreatedBy() *InputLogUpsert {
	u.SetExcluded(inputlog.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *InputLogUpsert) AddCreatedBy(v int64) *InputLogUpsert {
	u.Add(inputlog.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *InputLogUpsert) SetUpdatedBy(v int64) *InputLogUpsert {
	u.Set(inputlog.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *InputLogUpsert) UpdateUpdatedBy() *InputLogUpsert {
	u.SetExcluded(inputlog.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *InputLogUpsert) AddUpdatedBy(v int64) *InputLogUpsert {
	u.Add(inputlog.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *InputLogUpsert) SetUpdatedAt(v time.Time) *InputLogUpsert {
	u.Set(inputlog.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *InputLogUpsert) UpdateUpdatedAt() *InputLogUpsert {
	u.SetExcluded(inputlog.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *InputLogUpsert) SetDeletedAt(v time.Time) *InputLogUpsert {
	u.Set(inputlog.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *InputLogUpsert) UpdateDeletedAt() *InputLogUpsert {
	u.SetExcluded(inputlog.FieldDeletedAt)
	return u
}

// SetHeaders sets the "headers" field.
func (u *InputLogUpsert) SetHeaders(v string) *InputLogUpsert {
	u.Set(inputlog.FieldHeaders, v)
	return u
}

// UpdateHeaders sets the "headers" field to the value that was provided on create.
func (u *InputLogUpsert) UpdateHeaders() *InputLogUpsert {
	u.SetExcluded(inputlog.FieldHeaders)
	return u
}

// SetBody sets the "body" field.
func (u *InputLogUpsert) SetBody(v string) *InputLogUpsert {
	u.Set(inputlog.FieldBody, v)
	return u
}

// UpdateBody sets the "body" field to the value that was provided on create.
func (u *InputLogUpsert) UpdateBody() *InputLogUpsert {
	u.SetExcluded(inputlog.FieldBody)
	return u
}

// ClearBody clears the value of the "body" field.
func (u *InputLogUpsert) ClearBody() *InputLogUpsert {
	u.SetNull(inputlog.FieldBody)
	return u
}

// SetQuery sets the "query" field.
func (u *InputLogUpsert) SetQuery(v string) *InputLogUpsert {
	u.Set(inputlog.FieldQuery, v)
	return u
}

// UpdateQuery sets the "query" field to the value that was provided on create.
func (u *InputLogUpsert) UpdateQuery() *InputLogUpsert {
	u.SetExcluded(inputlog.FieldQuery)
	return u
}

// ClearQuery clears the value of the "query" field.
func (u *InputLogUpsert) ClearQuery() *InputLogUpsert {
	u.SetNull(inputlog.FieldQuery)
	return u
}

// SetURL sets the "url" field.
func (u *InputLogUpsert) SetURL(v string) *InputLogUpsert {
	u.Set(inputlog.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *InputLogUpsert) UpdateURL() *InputLogUpsert {
	u.SetExcluded(inputlog.FieldURL)
	return u
}

// SetIP sets the "ip" field.
func (u *InputLogUpsert) SetIP(v string) *InputLogUpsert {
	u.Set(inputlog.FieldIP, v)
	return u
}

// UpdateIP sets the "ip" field to the value that was provided on create.
func (u *InputLogUpsert) UpdateIP() *InputLogUpsert {
	u.SetExcluded(inputlog.FieldIP)
	return u
}

// ClearIP clears the value of the "ip" field.
func (u *InputLogUpsert) ClearIP() *InputLogUpsert {
	u.SetNull(inputlog.FieldIP)
	return u
}

// SetCaller sets the "caller" field.
func (u *InputLogUpsert) SetCaller(v string) *InputLogUpsert {
	u.Set(inputlog.FieldCaller, v)
	return u
}

// UpdateCaller sets the "caller" field to the value that was provided on create.
func (u *InputLogUpsert) UpdateCaller() *InputLogUpsert {
	u.SetExcluded(inputlog.FieldCaller)
	return u
}

// SetMethod sets the "method" field.
func (u *InputLogUpsert) SetMethod(v inputlog.Method) *InputLogUpsert {
	u.Set(inputlog.FieldMethod, v)
	return u
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *InputLogUpsert) UpdateMethod() *InputLogUpsert {
	u.SetExcluded(inputlog.FieldMethod)
	return u
}

// SetHmacKey sets the "hmac_key" field.
func (u *InputLogUpsert) SetHmacKey(v string) *InputLogUpsert {
	u.Set(inputlog.FieldHmacKey, v)
	return u
}

// UpdateHmacKey sets the "hmac_key" field to the value that was provided on create.
func (u *InputLogUpsert) UpdateHmacKey() *InputLogUpsert {
	u.SetExcluded(inputlog.FieldHmacKey)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.InputLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(inputlog.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *InputLogUpsertOne) UpdateNewValues() *InputLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(inputlog.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(inputlog.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.TraceID(); exists {
			s.SetIgnore(inputlog.FieldTraceID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.InputLog.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *InputLogUpsertOne) Ignore() *InputLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *InputLogUpsertOne) DoNothing() *InputLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the InputLogCreate.OnConflict
// documentation for more info.
func (u *InputLogUpsertOne) Update(set func(*InputLogUpsert)) *InputLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&InputLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *InputLogUpsertOne) SetCreatedBy(v int64) *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *InputLogUpsertOne) AddCreatedBy(v int64) *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *InputLogUpsertOne) UpdateCreatedBy() *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *InputLogUpsertOne) SetUpdatedBy(v int64) *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *InputLogUpsertOne) AddUpdatedBy(v int64) *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *InputLogUpsertOne) UpdateUpdatedBy() *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *InputLogUpsertOne) SetUpdatedAt(v time.Time) *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *InputLogUpsertOne) UpdateUpdatedAt() *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *InputLogUpsertOne) SetDeletedAt(v time.Time) *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *InputLogUpsertOne) UpdateDeletedAt() *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetHeaders sets the "headers" field.
func (u *InputLogUpsertOne) SetHeaders(v string) *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.SetHeaders(v)
	})
}

// UpdateHeaders sets the "headers" field to the value that was provided on create.
func (u *InputLogUpsertOne) UpdateHeaders() *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateHeaders()
	})
}

// SetBody sets the "body" field.
func (u *InputLogUpsertOne) SetBody(v string) *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.SetBody(v)
	})
}

// UpdateBody sets the "body" field to the value that was provided on create.
func (u *InputLogUpsertOne) UpdateBody() *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateBody()
	})
}

// ClearBody clears the value of the "body" field.
func (u *InputLogUpsertOne) ClearBody() *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.ClearBody()
	})
}

// SetQuery sets the "query" field.
func (u *InputLogUpsertOne) SetQuery(v string) *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.SetQuery(v)
	})
}

// UpdateQuery sets the "query" field to the value that was provided on create.
func (u *InputLogUpsertOne) UpdateQuery() *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateQuery()
	})
}

// ClearQuery clears the value of the "query" field.
func (u *InputLogUpsertOne) ClearQuery() *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.ClearQuery()
	})
}

// SetURL sets the "url" field.
func (u *InputLogUpsertOne) SetURL(v string) *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *InputLogUpsertOne) UpdateURL() *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateURL()
	})
}

// SetIP sets the "ip" field.
func (u *InputLogUpsertOne) SetIP(v string) *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.SetIP(v)
	})
}

// UpdateIP sets the "ip" field to the value that was provided on create.
func (u *InputLogUpsertOne) UpdateIP() *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateIP()
	})
}

// ClearIP clears the value of the "ip" field.
func (u *InputLogUpsertOne) ClearIP() *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.ClearIP()
	})
}

// SetCaller sets the "caller" field.
func (u *InputLogUpsertOne) SetCaller(v string) *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.SetCaller(v)
	})
}

// UpdateCaller sets the "caller" field to the value that was provided on create.
func (u *InputLogUpsertOne) UpdateCaller() *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateCaller()
	})
}

// SetMethod sets the "method" field.
func (u *InputLogUpsertOne) SetMethod(v inputlog.Method) *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *InputLogUpsertOne) UpdateMethod() *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateMethod()
	})
}

// SetHmacKey sets the "hmac_key" field.
func (u *InputLogUpsertOne) SetHmacKey(v string) *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.SetHmacKey(v)
	})
}

// UpdateHmacKey sets the "hmac_key" field to the value that was provided on create.
func (u *InputLogUpsertOne) UpdateHmacKey() *InputLogUpsertOne {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateHmacKey()
	})
}

// Exec executes the query.
func (u *InputLogUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for InputLogCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *InputLogUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *InputLogUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *InputLogUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// InputLogCreateBulk is the builder for creating many InputLog entities in bulk.
type InputLogCreateBulk struct {
	config
	builders []*InputLogCreate
	conflict []sql.ConflictOption
}

// Save creates the InputLog entities in the database.
func (ilcb *InputLogCreateBulk) Save(ctx context.Context) ([]*InputLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ilcb.builders))
	nodes := make([]*InputLog, len(ilcb.builders))
	mutators := make([]Mutator, len(ilcb.builders))
	for i := range ilcb.builders {
		func(i int, root context.Context) {
			builder := ilcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InputLogMutation)
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
					_, err = mutators[i+1].Mutate(root, ilcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ilcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ilcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ilcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ilcb *InputLogCreateBulk) SaveX(ctx context.Context) []*InputLog {
	v, err := ilcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ilcb *InputLogCreateBulk) Exec(ctx context.Context) error {
	_, err := ilcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ilcb *InputLogCreateBulk) ExecX(ctx context.Context) {
	if err := ilcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.InputLog.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.InputLogUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (ilcb *InputLogCreateBulk) OnConflict(opts ...sql.ConflictOption) *InputLogUpsertBulk {
	ilcb.conflict = opts
	return &InputLogUpsertBulk{
		create: ilcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.InputLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ilcb *InputLogCreateBulk) OnConflictColumns(columns ...string) *InputLogUpsertBulk {
	ilcb.conflict = append(ilcb.conflict, sql.ConflictColumns(columns...))
	return &InputLogUpsertBulk{
		create: ilcb,
	}
}

// InputLogUpsertBulk is the builder for "upsert"-ing
// a bulk of InputLog nodes.
type InputLogUpsertBulk struct {
	create *InputLogCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.InputLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(inputlog.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *InputLogUpsertBulk) UpdateNewValues() *InputLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(inputlog.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(inputlog.FieldCreatedAt)
			}
			if _, exists := b.mutation.TraceID(); exists {
				s.SetIgnore(inputlog.FieldTraceID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.InputLog.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *InputLogUpsertBulk) Ignore() *InputLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *InputLogUpsertBulk) DoNothing() *InputLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the InputLogCreateBulk.OnConflict
// documentation for more info.
func (u *InputLogUpsertBulk) Update(set func(*InputLogUpsert)) *InputLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&InputLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *InputLogUpsertBulk) SetCreatedBy(v int64) *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *InputLogUpsertBulk) AddCreatedBy(v int64) *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *InputLogUpsertBulk) UpdateCreatedBy() *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *InputLogUpsertBulk) SetUpdatedBy(v int64) *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *InputLogUpsertBulk) AddUpdatedBy(v int64) *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *InputLogUpsertBulk) UpdateUpdatedBy() *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *InputLogUpsertBulk) SetUpdatedAt(v time.Time) *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *InputLogUpsertBulk) UpdateUpdatedAt() *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *InputLogUpsertBulk) SetDeletedAt(v time.Time) *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *InputLogUpsertBulk) UpdateDeletedAt() *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetHeaders sets the "headers" field.
func (u *InputLogUpsertBulk) SetHeaders(v string) *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.SetHeaders(v)
	})
}

// UpdateHeaders sets the "headers" field to the value that was provided on create.
func (u *InputLogUpsertBulk) UpdateHeaders() *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateHeaders()
	})
}

// SetBody sets the "body" field.
func (u *InputLogUpsertBulk) SetBody(v string) *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.SetBody(v)
	})
}

// UpdateBody sets the "body" field to the value that was provided on create.
func (u *InputLogUpsertBulk) UpdateBody() *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateBody()
	})
}

// ClearBody clears the value of the "body" field.
func (u *InputLogUpsertBulk) ClearBody() *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.ClearBody()
	})
}

// SetQuery sets the "query" field.
func (u *InputLogUpsertBulk) SetQuery(v string) *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.SetQuery(v)
	})
}

// UpdateQuery sets the "query" field to the value that was provided on create.
func (u *InputLogUpsertBulk) UpdateQuery() *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateQuery()
	})
}

// ClearQuery clears the value of the "query" field.
func (u *InputLogUpsertBulk) ClearQuery() *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.ClearQuery()
	})
}

// SetURL sets the "url" field.
func (u *InputLogUpsertBulk) SetURL(v string) *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *InputLogUpsertBulk) UpdateURL() *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateURL()
	})
}

// SetIP sets the "ip" field.
func (u *InputLogUpsertBulk) SetIP(v string) *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.SetIP(v)
	})
}

// UpdateIP sets the "ip" field to the value that was provided on create.
func (u *InputLogUpsertBulk) UpdateIP() *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateIP()
	})
}

// ClearIP clears the value of the "ip" field.
func (u *InputLogUpsertBulk) ClearIP() *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.ClearIP()
	})
}

// SetCaller sets the "caller" field.
func (u *InputLogUpsertBulk) SetCaller(v string) *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.SetCaller(v)
	})
}

// UpdateCaller sets the "caller" field to the value that was provided on create.
func (u *InputLogUpsertBulk) UpdateCaller() *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateCaller()
	})
}

// SetMethod sets the "method" field.
func (u *InputLogUpsertBulk) SetMethod(v inputlog.Method) *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *InputLogUpsertBulk) UpdateMethod() *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateMethod()
	})
}

// SetHmacKey sets the "hmac_key" field.
func (u *InputLogUpsertBulk) SetHmacKey(v string) *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.SetHmacKey(v)
	})
}

// UpdateHmacKey sets the "hmac_key" field to the value that was provided on create.
func (u *InputLogUpsertBulk) UpdateHmacKey() *InputLogUpsertBulk {
	return u.Update(func(s *InputLogUpsert) {
		s.UpdateHmacKey()
	})
}

// Exec executes the query.
func (u *InputLogUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("cep_ent: OnConflict was set for builder %d. Set it on the InputLogCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for InputLogCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *InputLogUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
