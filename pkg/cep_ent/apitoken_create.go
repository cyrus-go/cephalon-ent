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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/apitoken"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ApiTokenCreate is the builder for creating a ApiToken entity.
type ApiTokenCreate struct {
	config
	mutation *ApiTokenMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (atc *ApiTokenCreate) SetCreatedBy(i int64) *ApiTokenCreate {
	atc.mutation.SetCreatedBy(i)
	return atc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableCreatedBy(i *int64) *ApiTokenCreate {
	if i != nil {
		atc.SetCreatedBy(*i)
	}
	return atc
}

// SetUpdatedBy sets the "updated_by" field.
func (atc *ApiTokenCreate) SetUpdatedBy(i int64) *ApiTokenCreate {
	atc.mutation.SetUpdatedBy(i)
	return atc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableUpdatedBy(i *int64) *ApiTokenCreate {
	if i != nil {
		atc.SetUpdatedBy(*i)
	}
	return atc
}

// SetCreatedAt sets the "created_at" field.
func (atc *ApiTokenCreate) SetCreatedAt(t time.Time) *ApiTokenCreate {
	atc.mutation.SetCreatedAt(t)
	return atc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableCreatedAt(t *time.Time) *ApiTokenCreate {
	if t != nil {
		atc.SetCreatedAt(*t)
	}
	return atc
}

// SetUpdatedAt sets the "updated_at" field.
func (atc *ApiTokenCreate) SetUpdatedAt(t time.Time) *ApiTokenCreate {
	atc.mutation.SetUpdatedAt(t)
	return atc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableUpdatedAt(t *time.Time) *ApiTokenCreate {
	if t != nil {
		atc.SetUpdatedAt(*t)
	}
	return atc
}

// SetDeletedAt sets the "deleted_at" field.
func (atc *ApiTokenCreate) SetDeletedAt(t time.Time) *ApiTokenCreate {
	atc.mutation.SetDeletedAt(t)
	return atc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableDeletedAt(t *time.Time) *ApiTokenCreate {
	if t != nil {
		atc.SetDeletedAt(*t)
	}
	return atc
}

// SetUserID sets the "user_id" field.
func (atc *ApiTokenCreate) SetUserID(i int64) *ApiTokenCreate {
	atc.mutation.SetUserID(i)
	return atc
}

// SetName sets the "name" field.
func (atc *ApiTokenCreate) SetName(s string) *ApiTokenCreate {
	atc.mutation.SetName(s)
	return atc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableName(s *string) *ApiTokenCreate {
	if s != nil {
		atc.SetName(*s)
	}
	return atc
}

// SetToken sets the "token" field.
func (atc *ApiTokenCreate) SetToken(s string) *ApiTokenCreate {
	atc.mutation.SetToken(s)
	return atc
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableToken(s *string) *ApiTokenCreate {
	if s != nil {
		atc.SetToken(*s)
	}
	return atc
}

// SetStatus sets the "status" field.
func (atc *ApiTokenCreate) SetStatus(ets enums.ApiTokenStatus) *ApiTokenCreate {
	atc.mutation.SetStatus(ets)
	return atc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableStatus(ets *enums.ApiTokenStatus) *ApiTokenCreate {
	if ets != nil {
		atc.SetStatus(*ets)
	}
	return atc
}

// SetID sets the "id" field.
func (atc *ApiTokenCreate) SetID(i int64) *ApiTokenCreate {
	atc.mutation.SetID(i)
	return atc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableID(i *int64) *ApiTokenCreate {
	if i != nil {
		atc.SetID(*i)
	}
	return atc
}

// SetUser sets the "user" edge to the User entity.
func (atc *ApiTokenCreate) SetUser(u *User) *ApiTokenCreate {
	return atc.SetUserID(u.ID)
}

// Mutation returns the ApiTokenMutation object of the builder.
func (atc *ApiTokenCreate) Mutation() *ApiTokenMutation {
	return atc.mutation
}

// Save creates the ApiToken in the database.
func (atc *ApiTokenCreate) Save(ctx context.Context) (*ApiToken, error) {
	atc.defaults()
	return withHooks(ctx, atc.sqlSave, atc.mutation, atc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (atc *ApiTokenCreate) SaveX(ctx context.Context) *ApiToken {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *ApiTokenCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *ApiTokenCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atc *ApiTokenCreate) defaults() {
	if _, ok := atc.mutation.CreatedBy(); !ok {
		v := apitoken.DefaultCreatedBy
		atc.mutation.SetCreatedBy(v)
	}
	if _, ok := atc.mutation.UpdatedBy(); !ok {
		v := apitoken.DefaultUpdatedBy
		atc.mutation.SetUpdatedBy(v)
	}
	if _, ok := atc.mutation.CreatedAt(); !ok {
		v := apitoken.DefaultCreatedAt()
		atc.mutation.SetCreatedAt(v)
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		v := apitoken.DefaultUpdatedAt()
		atc.mutation.SetUpdatedAt(v)
	}
	if _, ok := atc.mutation.DeletedAt(); !ok {
		v := apitoken.DefaultDeletedAt
		atc.mutation.SetDeletedAt(v)
	}
	if _, ok := atc.mutation.Name(); !ok {
		v := apitoken.DefaultName
		atc.mutation.SetName(v)
	}
	if _, ok := atc.mutation.Token(); !ok {
		v := apitoken.DefaultToken
		atc.mutation.SetToken(v)
	}
	if _, ok := atc.mutation.Status(); !ok {
		v := apitoken.DefaultStatus
		atc.mutation.SetStatus(v)
	}
	if _, ok := atc.mutation.ID(); !ok {
		v := apitoken.DefaultID()
		atc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *ApiTokenCreate) check() error {
	if _, ok := atc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "ApiToken.created_by"`)}
	}
	if _, ok := atc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "ApiToken.updated_by"`)}
	}
	if _, ok := atc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "ApiToken.created_at"`)}
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "ApiToken.updated_at"`)}
	}
	if _, ok := atc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "ApiToken.deleted_at"`)}
	}
	if _, ok := atc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`cep_ent: missing required field "ApiToken.user_id"`)}
	}
	if _, ok := atc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`cep_ent: missing required field "ApiToken.name"`)}
	}
	if _, ok := atc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`cep_ent: missing required field "ApiToken.token"`)}
	}
	if _, ok := atc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`cep_ent: missing required field "ApiToken.status"`)}
	}
	if v, ok := atc.mutation.Status(); ok {
		if err := apitoken.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`cep_ent: validator failed for field "ApiToken.status": %w`, err)}
		}
	}
	if _, ok := atc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`cep_ent: missing required edge "ApiToken.user"`)}
	}
	return nil
}

func (atc *ApiTokenCreate) sqlSave(ctx context.Context) (*ApiToken, error) {
	if err := atc.check(); err != nil {
		return nil, err
	}
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	atc.mutation.id = &_node.ID
	atc.mutation.done = true
	return _node, nil
}

func (atc *ApiTokenCreate) createSpec() (*ApiToken, *sqlgraph.CreateSpec) {
	var (
		_node = &ApiToken{config: atc.config}
		_spec = sqlgraph.NewCreateSpec(apitoken.Table, sqlgraph.NewFieldSpec(apitoken.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = atc.conflict
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := atc.mutation.CreatedBy(); ok {
		_spec.SetField(apitoken.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := atc.mutation.UpdatedBy(); ok {
		_spec.SetField(apitoken.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := atc.mutation.CreatedAt(); ok {
		_spec.SetField(apitoken.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := atc.mutation.UpdatedAt(); ok {
		_spec.SetField(apitoken.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := atc.mutation.DeletedAt(); ok {
		_spec.SetField(apitoken.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := atc.mutation.Name(); ok {
		_spec.SetField(apitoken.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := atc.mutation.Token(); ok {
		_spec.SetField(apitoken.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := atc.mutation.Status(); ok {
		_spec.SetField(apitoken.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := atc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.UserTable,
			Columns: []string{apitoken.UserColumn},
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
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ApiToken.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ApiTokenUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (atc *ApiTokenCreate) OnConflict(opts ...sql.ConflictOption) *ApiTokenUpsertOne {
	atc.conflict = opts
	return &ApiTokenUpsertOne{
		create: atc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (atc *ApiTokenCreate) OnConflictColumns(columns ...string) *ApiTokenUpsertOne {
	atc.conflict = append(atc.conflict, sql.ConflictColumns(columns...))
	return &ApiTokenUpsertOne{
		create: atc,
	}
}

type (
	// ApiTokenUpsertOne is the builder for "upsert"-ing
	//  one ApiToken node.
	ApiTokenUpsertOne struct {
		create *ApiTokenCreate
	}

	// ApiTokenUpsert is the "OnConflict" setter.
	ApiTokenUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *ApiTokenUpsert) SetCreatedBy(v int64) *ApiTokenUpsert {
	u.Set(apitoken.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *ApiTokenUpsert) UpdateCreatedBy() *ApiTokenUpsert {
	u.SetExcluded(apitoken.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *ApiTokenUpsert) AddCreatedBy(v int64) *ApiTokenUpsert {
	u.Add(apitoken.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *ApiTokenUpsert) SetUpdatedBy(v int64) *ApiTokenUpsert {
	u.Set(apitoken.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *ApiTokenUpsert) UpdateUpdatedBy() *ApiTokenUpsert {
	u.SetExcluded(apitoken.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *ApiTokenUpsert) AddUpdatedBy(v int64) *ApiTokenUpsert {
	u.Add(apitoken.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ApiTokenUpsert) SetUpdatedAt(v time.Time) *ApiTokenUpsert {
	u.Set(apitoken.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ApiTokenUpsert) UpdateUpdatedAt() *ApiTokenUpsert {
	u.SetExcluded(apitoken.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ApiTokenUpsert) SetDeletedAt(v time.Time) *ApiTokenUpsert {
	u.Set(apitoken.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ApiTokenUpsert) UpdateDeletedAt() *ApiTokenUpsert {
	u.SetExcluded(apitoken.FieldDeletedAt)
	return u
}

// SetUserID sets the "user_id" field.
func (u *ApiTokenUpsert) SetUserID(v int64) *ApiTokenUpsert {
	u.Set(apitoken.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ApiTokenUpsert) UpdateUserID() *ApiTokenUpsert {
	u.SetExcluded(apitoken.FieldUserID)
	return u
}

// SetName sets the "name" field.
func (u *ApiTokenUpsert) SetName(v string) *ApiTokenUpsert {
	u.Set(apitoken.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ApiTokenUpsert) UpdateName() *ApiTokenUpsert {
	u.SetExcluded(apitoken.FieldName)
	return u
}

// SetToken sets the "token" field.
func (u *ApiTokenUpsert) SetToken(v string) *ApiTokenUpsert {
	u.Set(apitoken.FieldToken, v)
	return u
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *ApiTokenUpsert) UpdateToken() *ApiTokenUpsert {
	u.SetExcluded(apitoken.FieldToken)
	return u
}

// SetStatus sets the "status" field.
func (u *ApiTokenUpsert) SetStatus(v enums.ApiTokenStatus) *ApiTokenUpsert {
	u.Set(apitoken.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ApiTokenUpsert) UpdateStatus() *ApiTokenUpsert {
	u.SetExcluded(apitoken.FieldStatus)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(apitoken.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ApiTokenUpsertOne) UpdateNewValues() *ApiTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(apitoken.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(apitoken.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ApiTokenUpsertOne) Ignore() *ApiTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ApiTokenUpsertOne) DoNothing() *ApiTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ApiTokenCreate.OnConflict
// documentation for more info.
func (u *ApiTokenUpsertOne) Update(set func(*ApiTokenUpsert)) *ApiTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ApiTokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *ApiTokenUpsertOne) SetCreatedBy(v int64) *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *ApiTokenUpsertOne) AddCreatedBy(v int64) *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *ApiTokenUpsertOne) UpdateCreatedBy() *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *ApiTokenUpsertOne) SetUpdatedBy(v int64) *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *ApiTokenUpsertOne) AddUpdatedBy(v int64) *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *ApiTokenUpsertOne) UpdateUpdatedBy() *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ApiTokenUpsertOne) SetUpdatedAt(v time.Time) *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ApiTokenUpsertOne) UpdateUpdatedAt() *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ApiTokenUpsertOne) SetDeletedAt(v time.Time) *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ApiTokenUpsertOne) UpdateDeletedAt() *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *ApiTokenUpsertOne) SetUserID(v int64) *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ApiTokenUpsertOne) UpdateUserID() *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateUserID()
	})
}

// SetName sets the "name" field.
func (u *ApiTokenUpsertOne) SetName(v string) *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ApiTokenUpsertOne) UpdateName() *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateName()
	})
}

// SetToken sets the "token" field.
func (u *ApiTokenUpsertOne) SetToken(v string) *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *ApiTokenUpsertOne) UpdateToken() *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateToken()
	})
}

// SetStatus sets the "status" field.
func (u *ApiTokenUpsertOne) SetStatus(v enums.ApiTokenStatus) *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ApiTokenUpsertOne) UpdateStatus() *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *ApiTokenUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for ApiTokenCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ApiTokenUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ApiTokenUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ApiTokenUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ApiTokenCreateBulk is the builder for creating many ApiToken entities in bulk.
type ApiTokenCreateBulk struct {
	config
	err      error
	builders []*ApiTokenCreate
	conflict []sql.ConflictOption
}

// Save creates the ApiToken entities in the database.
func (atcb *ApiTokenCreateBulk) Save(ctx context.Context) ([]*ApiToken, error) {
	if atcb.err != nil {
		return nil, atcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*ApiToken, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApiTokenMutation)
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
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = atcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *ApiTokenCreateBulk) SaveX(ctx context.Context) []*ApiToken {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *ApiTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *ApiTokenCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ApiToken.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ApiTokenUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (atcb *ApiTokenCreateBulk) OnConflict(opts ...sql.ConflictOption) *ApiTokenUpsertBulk {
	atcb.conflict = opts
	return &ApiTokenUpsertBulk{
		create: atcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (atcb *ApiTokenCreateBulk) OnConflictColumns(columns ...string) *ApiTokenUpsertBulk {
	atcb.conflict = append(atcb.conflict, sql.ConflictColumns(columns...))
	return &ApiTokenUpsertBulk{
		create: atcb,
	}
}

// ApiTokenUpsertBulk is the builder for "upsert"-ing
// a bulk of ApiToken nodes.
type ApiTokenUpsertBulk struct {
	create *ApiTokenCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(apitoken.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ApiTokenUpsertBulk) UpdateNewValues() *ApiTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(apitoken.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(apitoken.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ApiTokenUpsertBulk) Ignore() *ApiTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ApiTokenUpsertBulk) DoNothing() *ApiTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ApiTokenCreateBulk.OnConflict
// documentation for more info.
func (u *ApiTokenUpsertBulk) Update(set func(*ApiTokenUpsert)) *ApiTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ApiTokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *ApiTokenUpsertBulk) SetCreatedBy(v int64) *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *ApiTokenUpsertBulk) AddCreatedBy(v int64) *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *ApiTokenUpsertBulk) UpdateCreatedBy() *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *ApiTokenUpsertBulk) SetUpdatedBy(v int64) *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *ApiTokenUpsertBulk) AddUpdatedBy(v int64) *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *ApiTokenUpsertBulk) UpdateUpdatedBy() *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ApiTokenUpsertBulk) SetUpdatedAt(v time.Time) *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ApiTokenUpsertBulk) UpdateUpdatedAt() *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ApiTokenUpsertBulk) SetDeletedAt(v time.Time) *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ApiTokenUpsertBulk) UpdateDeletedAt() *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *ApiTokenUpsertBulk) SetUserID(v int64) *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ApiTokenUpsertBulk) UpdateUserID() *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateUserID()
	})
}

// SetName sets the "name" field.
func (u *ApiTokenUpsertBulk) SetName(v string) *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ApiTokenUpsertBulk) UpdateName() *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateName()
	})
}

// SetToken sets the "token" field.
func (u *ApiTokenUpsertBulk) SetToken(v string) *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *ApiTokenUpsertBulk) UpdateToken() *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateToken()
	})
}

// SetStatus sets the "status" field.
func (u *ApiTokenUpsertBulk) SetStatus(v enums.ApiTokenStatus) *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ApiTokenUpsertBulk) UpdateStatus() *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *ApiTokenUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("cep_ent: OnConflict was set for builder %d. Set it on the ApiTokenCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for ApiTokenCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ApiTokenUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
