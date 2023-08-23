// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/predicate"
	"cephalon-ent/pkg/cep_ent/user"
	"cephalon-ent/pkg/cep_ent/vxaccount"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VXAccountUpdate is the builder for updating VXAccount entities.
type VXAccountUpdate struct {
	config
	hooks    []Hook
	mutation *VXAccountMutation
}

// Where appends a list predicates to the VXAccountUpdate builder.
func (vau *VXAccountUpdate) Where(ps ...predicate.VXAccount) *VXAccountUpdate {
	vau.mutation.Where(ps...)
	return vau
}

// SetCreatedBy sets the "created_by" field.
func (vau *VXAccountUpdate) SetCreatedBy(i int64) *VXAccountUpdate {
	vau.mutation.ResetCreatedBy()
	vau.mutation.SetCreatedBy(i)
	return vau
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (vau *VXAccountUpdate) SetNillableCreatedBy(i *int64) *VXAccountUpdate {
	if i != nil {
		vau.SetCreatedBy(*i)
	}
	return vau
}

// AddCreatedBy adds i to the "created_by" field.
func (vau *VXAccountUpdate) AddCreatedBy(i int64) *VXAccountUpdate {
	vau.mutation.AddCreatedBy(i)
	return vau
}

// SetUpdatedBy sets the "updated_by" field.
func (vau *VXAccountUpdate) SetUpdatedBy(i int64) *VXAccountUpdate {
	vau.mutation.ResetUpdatedBy()
	vau.mutation.SetUpdatedBy(i)
	return vau
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (vau *VXAccountUpdate) SetNillableUpdatedBy(i *int64) *VXAccountUpdate {
	if i != nil {
		vau.SetUpdatedBy(*i)
	}
	return vau
}

// AddUpdatedBy adds i to the "updated_by" field.
func (vau *VXAccountUpdate) AddUpdatedBy(i int64) *VXAccountUpdate {
	vau.mutation.AddUpdatedBy(i)
	return vau
}

// SetUpdatedAt sets the "updated_at" field.
func (vau *VXAccountUpdate) SetUpdatedAt(t time.Time) *VXAccountUpdate {
	vau.mutation.SetUpdatedAt(t)
	return vau
}

// SetDeletedAt sets the "deleted_at" field.
func (vau *VXAccountUpdate) SetDeletedAt(t time.Time) *VXAccountUpdate {
	vau.mutation.SetDeletedAt(t)
	return vau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (vau *VXAccountUpdate) SetNillableDeletedAt(t *time.Time) *VXAccountUpdate {
	if t != nil {
		vau.SetDeletedAt(*t)
	}
	return vau
}

// SetOpenID sets the "open_id" field.
func (vau *VXAccountUpdate) SetOpenID(s string) *VXAccountUpdate {
	vau.mutation.SetOpenID(s)
	return vau
}

// SetNillableOpenID sets the "open_id" field if the given value is not nil.
func (vau *VXAccountUpdate) SetNillableOpenID(s *string) *VXAccountUpdate {
	if s != nil {
		vau.SetOpenID(*s)
	}
	return vau
}

// SetUnionID sets the "union_id" field.
func (vau *VXAccountUpdate) SetUnionID(s string) *VXAccountUpdate {
	vau.mutation.SetUnionID(s)
	return vau
}

// SetNillableUnionID sets the "union_id" field if the given value is not nil.
func (vau *VXAccountUpdate) SetNillableUnionID(s *string) *VXAccountUpdate {
	if s != nil {
		vau.SetUnionID(*s)
	}
	return vau
}

// SetScope sets the "scope" field.
func (vau *VXAccountUpdate) SetScope(s string) *VXAccountUpdate {
	vau.mutation.SetScope(s)
	return vau
}

// SetNillableScope sets the "scope" field if the given value is not nil.
func (vau *VXAccountUpdate) SetNillableScope(s *string) *VXAccountUpdate {
	if s != nil {
		vau.SetScope(*s)
	}
	return vau
}

// SetSessionKey sets the "session_key" field.
func (vau *VXAccountUpdate) SetSessionKey(s string) *VXAccountUpdate {
	vau.mutation.SetSessionKey(s)
	return vau
}

// SetNillableSessionKey sets the "session_key" field if the given value is not nil.
func (vau *VXAccountUpdate) SetNillableSessionKey(s *string) *VXAccountUpdate {
	if s != nil {
		vau.SetSessionKey(*s)
	}
	return vau
}

// SetUserID sets the "user_id" field.
func (vau *VXAccountUpdate) SetUserID(i int64) *VXAccountUpdate {
	vau.mutation.SetUserID(i)
	return vau
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (vau *VXAccountUpdate) SetNillableUserID(i *int64) *VXAccountUpdate {
	if i != nil {
		vau.SetUserID(*i)
	}
	return vau
}

// SetUser sets the "user" edge to the User entity.
func (vau *VXAccountUpdate) SetUser(u *User) *VXAccountUpdate {
	return vau.SetUserID(u.ID)
}

// Mutation returns the VXAccountMutation object of the builder.
func (vau *VXAccountUpdate) Mutation() *VXAccountMutation {
	return vau.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (vau *VXAccountUpdate) ClearUser() *VXAccountUpdate {
	vau.mutation.ClearUser()
	return vau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vau *VXAccountUpdate) Save(ctx context.Context) (int, error) {
	vau.defaults()
	return withHooks(ctx, vau.sqlSave, vau.mutation, vau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vau *VXAccountUpdate) SaveX(ctx context.Context) int {
	affected, err := vau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vau *VXAccountUpdate) Exec(ctx context.Context) error {
	_, err := vau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vau *VXAccountUpdate) ExecX(ctx context.Context) {
	if err := vau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vau *VXAccountUpdate) defaults() {
	if _, ok := vau.mutation.UpdatedAt(); !ok {
		v := vxaccount.UpdateDefaultUpdatedAt()
		vau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vau *VXAccountUpdate) check() error {
	if _, ok := vau.mutation.UserID(); vau.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "VXAccount.user"`)
	}
	return nil
}

func (vau *VXAccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := vau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(vxaccount.Table, vxaccount.Columns, sqlgraph.NewFieldSpec(vxaccount.FieldID, field.TypeInt64))
	if ps := vau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vau.mutation.CreatedBy(); ok {
		_spec.SetField(vxaccount.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := vau.mutation.AddedCreatedBy(); ok {
		_spec.AddField(vxaccount.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := vau.mutation.UpdatedBy(); ok {
		_spec.SetField(vxaccount.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := vau.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(vxaccount.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := vau.mutation.UpdatedAt(); ok {
		_spec.SetField(vxaccount.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := vau.mutation.DeletedAt(); ok {
		_spec.SetField(vxaccount.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := vau.mutation.OpenID(); ok {
		_spec.SetField(vxaccount.FieldOpenID, field.TypeString, value)
	}
	if value, ok := vau.mutation.UnionID(); ok {
		_spec.SetField(vxaccount.FieldUnionID, field.TypeString, value)
	}
	if value, ok := vau.mutation.Scope(); ok {
		_spec.SetField(vxaccount.FieldScope, field.TypeString, value)
	}
	if value, ok := vau.mutation.SessionKey(); ok {
		_spec.SetField(vxaccount.FieldSessionKey, field.TypeString, value)
	}
	if vau.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vxaccount.UserTable,
			Columns: []string{vxaccount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vau.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vxaccount.UserTable,
			Columns: []string{vxaccount.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, vau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vxaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vau.mutation.done = true
	return n, nil
}

// VXAccountUpdateOne is the builder for updating a single VXAccount entity.
type VXAccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VXAccountMutation
}

// SetCreatedBy sets the "created_by" field.
func (vauo *VXAccountUpdateOne) SetCreatedBy(i int64) *VXAccountUpdateOne {
	vauo.mutation.ResetCreatedBy()
	vauo.mutation.SetCreatedBy(i)
	return vauo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (vauo *VXAccountUpdateOne) SetNillableCreatedBy(i *int64) *VXAccountUpdateOne {
	if i != nil {
		vauo.SetCreatedBy(*i)
	}
	return vauo
}

// AddCreatedBy adds i to the "created_by" field.
func (vauo *VXAccountUpdateOne) AddCreatedBy(i int64) *VXAccountUpdateOne {
	vauo.mutation.AddCreatedBy(i)
	return vauo
}

// SetUpdatedBy sets the "updated_by" field.
func (vauo *VXAccountUpdateOne) SetUpdatedBy(i int64) *VXAccountUpdateOne {
	vauo.mutation.ResetUpdatedBy()
	vauo.mutation.SetUpdatedBy(i)
	return vauo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (vauo *VXAccountUpdateOne) SetNillableUpdatedBy(i *int64) *VXAccountUpdateOne {
	if i != nil {
		vauo.SetUpdatedBy(*i)
	}
	return vauo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (vauo *VXAccountUpdateOne) AddUpdatedBy(i int64) *VXAccountUpdateOne {
	vauo.mutation.AddUpdatedBy(i)
	return vauo
}

// SetUpdatedAt sets the "updated_at" field.
func (vauo *VXAccountUpdateOne) SetUpdatedAt(t time.Time) *VXAccountUpdateOne {
	vauo.mutation.SetUpdatedAt(t)
	return vauo
}

// SetDeletedAt sets the "deleted_at" field.
func (vauo *VXAccountUpdateOne) SetDeletedAt(t time.Time) *VXAccountUpdateOne {
	vauo.mutation.SetDeletedAt(t)
	return vauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (vauo *VXAccountUpdateOne) SetNillableDeletedAt(t *time.Time) *VXAccountUpdateOne {
	if t != nil {
		vauo.SetDeletedAt(*t)
	}
	return vauo
}

// SetOpenID sets the "open_id" field.
func (vauo *VXAccountUpdateOne) SetOpenID(s string) *VXAccountUpdateOne {
	vauo.mutation.SetOpenID(s)
	return vauo
}

// SetNillableOpenID sets the "open_id" field if the given value is not nil.
func (vauo *VXAccountUpdateOne) SetNillableOpenID(s *string) *VXAccountUpdateOne {
	if s != nil {
		vauo.SetOpenID(*s)
	}
	return vauo
}

// SetUnionID sets the "union_id" field.
func (vauo *VXAccountUpdateOne) SetUnionID(s string) *VXAccountUpdateOne {
	vauo.mutation.SetUnionID(s)
	return vauo
}

// SetNillableUnionID sets the "union_id" field if the given value is not nil.
func (vauo *VXAccountUpdateOne) SetNillableUnionID(s *string) *VXAccountUpdateOne {
	if s != nil {
		vauo.SetUnionID(*s)
	}
	return vauo
}

// SetScope sets the "scope" field.
func (vauo *VXAccountUpdateOne) SetScope(s string) *VXAccountUpdateOne {
	vauo.mutation.SetScope(s)
	return vauo
}

// SetNillableScope sets the "scope" field if the given value is not nil.
func (vauo *VXAccountUpdateOne) SetNillableScope(s *string) *VXAccountUpdateOne {
	if s != nil {
		vauo.SetScope(*s)
	}
	return vauo
}

// SetSessionKey sets the "session_key" field.
func (vauo *VXAccountUpdateOne) SetSessionKey(s string) *VXAccountUpdateOne {
	vauo.mutation.SetSessionKey(s)
	return vauo
}

// SetNillableSessionKey sets the "session_key" field if the given value is not nil.
func (vauo *VXAccountUpdateOne) SetNillableSessionKey(s *string) *VXAccountUpdateOne {
	if s != nil {
		vauo.SetSessionKey(*s)
	}
	return vauo
}

// SetUserID sets the "user_id" field.
func (vauo *VXAccountUpdateOne) SetUserID(i int64) *VXAccountUpdateOne {
	vauo.mutation.SetUserID(i)
	return vauo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (vauo *VXAccountUpdateOne) SetNillableUserID(i *int64) *VXAccountUpdateOne {
	if i != nil {
		vauo.SetUserID(*i)
	}
	return vauo
}

// SetUser sets the "user" edge to the User entity.
func (vauo *VXAccountUpdateOne) SetUser(u *User) *VXAccountUpdateOne {
	return vauo.SetUserID(u.ID)
}

// Mutation returns the VXAccountMutation object of the builder.
func (vauo *VXAccountUpdateOne) Mutation() *VXAccountMutation {
	return vauo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (vauo *VXAccountUpdateOne) ClearUser() *VXAccountUpdateOne {
	vauo.mutation.ClearUser()
	return vauo
}

// Where appends a list predicates to the VXAccountUpdate builder.
func (vauo *VXAccountUpdateOne) Where(ps ...predicate.VXAccount) *VXAccountUpdateOne {
	vauo.mutation.Where(ps...)
	return vauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vauo *VXAccountUpdateOne) Select(field string, fields ...string) *VXAccountUpdateOne {
	vauo.fields = append([]string{field}, fields...)
	return vauo
}

// Save executes the query and returns the updated VXAccount entity.
func (vauo *VXAccountUpdateOne) Save(ctx context.Context) (*VXAccount, error) {
	vauo.defaults()
	return withHooks(ctx, vauo.sqlSave, vauo.mutation, vauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vauo *VXAccountUpdateOne) SaveX(ctx context.Context) *VXAccount {
	node, err := vauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vauo *VXAccountUpdateOne) Exec(ctx context.Context) error {
	_, err := vauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vauo *VXAccountUpdateOne) ExecX(ctx context.Context) {
	if err := vauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vauo *VXAccountUpdateOne) defaults() {
	if _, ok := vauo.mutation.UpdatedAt(); !ok {
		v := vxaccount.UpdateDefaultUpdatedAt()
		vauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vauo *VXAccountUpdateOne) check() error {
	if _, ok := vauo.mutation.UserID(); vauo.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "VXAccount.user"`)
	}
	return nil
}

func (vauo *VXAccountUpdateOne) sqlSave(ctx context.Context) (_node *VXAccount, err error) {
	if err := vauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(vxaccount.Table, vxaccount.Columns, sqlgraph.NewFieldSpec(vxaccount.FieldID, field.TypeInt64))
	id, ok := vauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "VXAccount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vxaccount.FieldID)
		for _, f := range fields {
			if !vxaccount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != vxaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vauo.mutation.CreatedBy(); ok {
		_spec.SetField(vxaccount.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := vauo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(vxaccount.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := vauo.mutation.UpdatedBy(); ok {
		_spec.SetField(vxaccount.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := vauo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(vxaccount.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := vauo.mutation.UpdatedAt(); ok {
		_spec.SetField(vxaccount.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := vauo.mutation.DeletedAt(); ok {
		_spec.SetField(vxaccount.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := vauo.mutation.OpenID(); ok {
		_spec.SetField(vxaccount.FieldOpenID, field.TypeString, value)
	}
	if value, ok := vauo.mutation.UnionID(); ok {
		_spec.SetField(vxaccount.FieldUnionID, field.TypeString, value)
	}
	if value, ok := vauo.mutation.Scope(); ok {
		_spec.SetField(vxaccount.FieldScope, field.TypeString, value)
	}
	if value, ok := vauo.mutation.SessionKey(); ok {
		_spec.SetField(vxaccount.FieldSessionKey, field.TypeString, value)
	}
	if vauo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vxaccount.UserTable,
			Columns: []string{vxaccount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vauo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vxaccount.UserTable,
			Columns: []string{vxaccount.UserColumn},
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
	_node = &VXAccount{config: vauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vxaccount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vauo.mutation.done = true
	return _node, nil
}
