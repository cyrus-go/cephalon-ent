// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/collect"
	"cephalon-ent/pkg/cep_ent/costaccount"
	"cephalon-ent/pkg/cep_ent/costbill"
	"cephalon-ent/pkg/cep_ent/device"
	"cephalon-ent/pkg/cep_ent/earnbill"
	"cephalon-ent/pkg/cep_ent/missionbatch"
	"cephalon-ent/pkg/cep_ent/missionconsumeorder"
	"cephalon-ent/pkg/cep_ent/missionproduceorder"
	"cephalon-ent/pkg/cep_ent/profitaccount"
	"cephalon-ent/pkg/cep_ent/profitsetting"
	"cephalon-ent/pkg/cep_ent/rechargeorder"
	"cephalon-ent/pkg/cep_ent/user"
	"cephalon-ent/pkg/cep_ent/userdevice"
	"cephalon-ent/pkg/cep_ent/vxaccount"
	"cephalon-ent/pkg/cep_ent/vxsocial"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (uc *UserCreate) SetCreatedBy(i int64) *UserCreate {
	uc.mutation.SetCreatedBy(i)
	return uc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedBy(i *int64) *UserCreate {
	if i != nil {
		uc.SetCreatedBy(*i)
	}
	return uc
}

// SetUpdatedBy sets the "updated_by" field.
func (uc *UserCreate) SetUpdatedBy(i int64) *UserCreate {
	uc.mutation.SetUpdatedBy(i)
	return uc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedBy(i *int64) *UserCreate {
	if i != nil {
		uc.SetUpdatedBy(*i)
	}
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetDeletedAt sets the "deleted_at" field.
func (uc *UserCreate) SetDeletedAt(t time.Time) *UserCreate {
	uc.mutation.SetDeletedAt(t)
	return uc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeletedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetDeletedAt(*t)
	}
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uc *UserCreate) SetNillableName(s *string) *UserCreate {
	if s != nil {
		uc.SetName(*s)
	}
	return uc
}

// SetJpgURL sets the "jpg_url" field.
func (uc *UserCreate) SetJpgURL(s string) *UserCreate {
	uc.mutation.SetJpgURL(s)
	return uc
}

// SetNillableJpgURL sets the "jpg_url" field if the given value is not nil.
func (uc *UserCreate) SetNillableJpgURL(s *string) *UserCreate {
	if s != nil {
		uc.SetJpgURL(*s)
	}
	return uc
}

// SetKey sets the "key" field.
func (uc *UserCreate) SetKey(s string) *UserCreate {
	uc.mutation.SetKey(s)
	return uc
}

// SetNillableKey sets the "key" field if the given value is not nil.
func (uc *UserCreate) SetNillableKey(s *string) *UserCreate {
	if s != nil {
		uc.SetKey(*s)
	}
	return uc
}

// SetSecret sets the "secret" field.
func (uc *UserCreate) SetSecret(s string) *UserCreate {
	uc.mutation.SetSecret(s)
	return uc
}

// SetNillableSecret sets the "secret" field if the given value is not nil.
func (uc *UserCreate) SetNillableSecret(s *string) *UserCreate {
	if s != nil {
		uc.SetSecret(*s)
	}
	return uc
}

// SetPhone sets the "phone" field.
func (uc *UserCreate) SetPhone(s string) *UserCreate {
	uc.mutation.SetPhone(s)
	return uc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uc *UserCreate) SetNillablePhone(s *string) *UserCreate {
	if s != nil {
		uc.SetPhone(*s)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uc *UserCreate) SetNillablePassword(s *string) *UserCreate {
	if s != nil {
		uc.SetPassword(*s)
	}
	return uc
}

// SetIsFrozen sets the "is_frozen" field.
func (uc *UserCreate) SetIsFrozen(b bool) *UserCreate {
	uc.mutation.SetIsFrozen(b)
	return uc
}

// SetNillableIsFrozen sets the "is_frozen" field if the given value is not nil.
func (uc *UserCreate) SetNillableIsFrozen(b *bool) *UserCreate {
	if b != nil {
		uc.SetIsFrozen(*b)
	}
	return uc
}

// SetUserType sets the "user_type" field.
func (uc *UserCreate) SetUserType(ut user.UserType) *UserCreate {
	uc.mutation.SetUserType(ut)
	return uc
}

// SetNillableUserType sets the "user_type" field if the given value is not nil.
func (uc *UserCreate) SetNillableUserType(ut *user.UserType) *UserCreate {
	if ut != nil {
		uc.SetUserType(*ut)
	}
	return uc
}

// SetParentID sets the "parent_id" field.
func (uc *UserCreate) SetParentID(i int64) *UserCreate {
	uc.mutation.SetParentID(i)
	return uc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableParentID(i *int64) *UserCreate {
	if i != nil {
		uc.SetParentID(*i)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(i int64) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(i *int64) *UserCreate {
	if i != nil {
		uc.SetID(*i)
	}
	return uc
}

// AddVxAccountIDs adds the "vx_accounts" edge to the VXAccount entity by IDs.
func (uc *UserCreate) AddVxAccountIDs(ids ...int64) *UserCreate {
	uc.mutation.AddVxAccountIDs(ids...)
	return uc
}

// AddVxAccounts adds the "vx_accounts" edges to the VXAccount entity.
func (uc *UserCreate) AddVxAccounts(v ...*VXAccount) *UserCreate {
	ids := make([]int64, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uc.AddVxAccountIDs(ids...)
}

// AddCollectIDs adds the "collects" edge to the Collect entity by IDs.
func (uc *UserCreate) AddCollectIDs(ids ...int64) *UserCreate {
	uc.mutation.AddCollectIDs(ids...)
	return uc
}

// AddCollects adds the "collects" edges to the Collect entity.
func (uc *UserCreate) AddCollects(c ...*Collect) *UserCreate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCollectIDs(ids...)
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (uc *UserCreate) AddDeviceIDs(ids ...int64) *UserCreate {
	uc.mutation.AddDeviceIDs(ids...)
	return uc
}

// AddDevices adds the "devices" edges to the Device entity.
func (uc *UserCreate) AddDevices(d ...*Device) *UserCreate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uc.AddDeviceIDs(ids...)
}

// AddProfitSettingIDs adds the "profit_settings" edge to the ProfitSetting entity by IDs.
func (uc *UserCreate) AddProfitSettingIDs(ids ...int64) *UserCreate {
	uc.mutation.AddProfitSettingIDs(ids...)
	return uc
}

// AddProfitSettings adds the "profit_settings" edges to the ProfitSetting entity.
func (uc *UserCreate) AddProfitSettings(p ...*ProfitSetting) *UserCreate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddProfitSettingIDs(ids...)
}

// SetCostAccountID sets the "cost_account" edge to the CostAccount entity by ID.
func (uc *UserCreate) SetCostAccountID(id int64) *UserCreate {
	uc.mutation.SetCostAccountID(id)
	return uc
}

// SetNillableCostAccountID sets the "cost_account" edge to the CostAccount entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableCostAccountID(id *int64) *UserCreate {
	if id != nil {
		uc = uc.SetCostAccountID(*id)
	}
	return uc
}

// SetCostAccount sets the "cost_account" edge to the CostAccount entity.
func (uc *UserCreate) SetCostAccount(c *CostAccount) *UserCreate {
	return uc.SetCostAccountID(c.ID)
}

// SetProfitAccountID sets the "profit_account" edge to the ProfitAccount entity by ID.
func (uc *UserCreate) SetProfitAccountID(id int64) *UserCreate {
	uc.mutation.SetProfitAccountID(id)
	return uc
}

// SetNillableProfitAccountID sets the "profit_account" edge to the ProfitAccount entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableProfitAccountID(id *int64) *UserCreate {
	if id != nil {
		uc = uc.SetProfitAccountID(*id)
	}
	return uc
}

// SetProfitAccount sets the "profit_account" edge to the ProfitAccount entity.
func (uc *UserCreate) SetProfitAccount(p *ProfitAccount) *UserCreate {
	return uc.SetProfitAccountID(p.ID)
}

// AddCostBillIDs adds the "cost_bills" edge to the CostBill entity by IDs.
func (uc *UserCreate) AddCostBillIDs(ids ...int64) *UserCreate {
	uc.mutation.AddCostBillIDs(ids...)
	return uc
}

// AddCostBills adds the "cost_bills" edges to the CostBill entity.
func (uc *UserCreate) AddCostBills(c ...*CostBill) *UserCreate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCostBillIDs(ids...)
}

// AddEarnBillIDs adds the "earn_bills" edge to the EarnBill entity by IDs.
func (uc *UserCreate) AddEarnBillIDs(ids ...int64) *UserCreate {
	uc.mutation.AddEarnBillIDs(ids...)
	return uc
}

// AddEarnBills adds the "earn_bills" edges to the EarnBill entity.
func (uc *UserCreate) AddEarnBills(e ...*EarnBill) *UserCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uc.AddEarnBillIDs(ids...)
}

// AddMissionConsumeOrderIDs adds the "mission_consume_orders" edge to the MissionConsumeOrder entity by IDs.
func (uc *UserCreate) AddMissionConsumeOrderIDs(ids ...int64) *UserCreate {
	uc.mutation.AddMissionConsumeOrderIDs(ids...)
	return uc
}

// AddMissionConsumeOrders adds the "mission_consume_orders" edges to the MissionConsumeOrder entity.
func (uc *UserCreate) AddMissionConsumeOrders(m ...*MissionConsumeOrder) *UserCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uc.AddMissionConsumeOrderIDs(ids...)
}

// AddMissionProduceOrderIDs adds the "mission_produce_orders" edge to the MissionProduceOrder entity by IDs.
func (uc *UserCreate) AddMissionProduceOrderIDs(ids ...int64) *UserCreate {
	uc.mutation.AddMissionProduceOrderIDs(ids...)
	return uc
}

// AddMissionProduceOrders adds the "mission_produce_orders" edges to the MissionProduceOrder entity.
func (uc *UserCreate) AddMissionProduceOrders(m ...*MissionProduceOrder) *UserCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uc.AddMissionProduceOrderIDs(ids...)
}

// AddRechargeOrderIDs adds the "recharge_orders" edge to the RechargeOrder entity by IDs.
func (uc *UserCreate) AddRechargeOrderIDs(ids ...int64) *UserCreate {
	uc.mutation.AddRechargeOrderIDs(ids...)
	return uc
}

// AddRechargeOrders adds the "recharge_orders" edges to the RechargeOrder entity.
func (uc *UserCreate) AddRechargeOrders(r ...*RechargeOrder) *UserCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddRechargeOrderIDs(ids...)
}

// AddVxSocialIDs adds the "vx_socials" edge to the VXSocial entity by IDs.
func (uc *UserCreate) AddVxSocialIDs(ids ...int64) *UserCreate {
	uc.mutation.AddVxSocialIDs(ids...)
	return uc
}

// AddVxSocials adds the "vx_socials" edges to the VXSocial entity.
func (uc *UserCreate) AddVxSocials(v ...*VXSocial) *UserCreate {
	ids := make([]int64, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uc.AddVxSocialIDs(ids...)
}

// AddMissionBatchIDs adds the "mission_batches" edge to the MissionBatch entity by IDs.
func (uc *UserCreate) AddMissionBatchIDs(ids ...int64) *UserCreate {
	uc.mutation.AddMissionBatchIDs(ids...)
	return uc
}

// AddMissionBatches adds the "mission_batches" edges to the MissionBatch entity.
func (uc *UserCreate) AddMissionBatches(m ...*MissionBatch) *UserCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uc.AddMissionBatchIDs(ids...)
}

// AddUserDeviceIDs adds the "user_devices" edge to the UserDevice entity by IDs.
func (uc *UserCreate) AddUserDeviceIDs(ids ...int64) *UserCreate {
	uc.mutation.AddUserDeviceIDs(ids...)
	return uc
}

// AddUserDevices adds the "user_devices" edges to the UserDevice entity.
func (uc *UserCreate) AddUserDevices(u ...*UserDevice) *UserCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddUserDeviceIDs(ids...)
}

// SetParent sets the "parent" edge to the User entity.
func (uc *UserCreate) SetParent(u *User) *UserCreate {
	return uc.SetParentID(u.ID)
}

// AddChildIDs adds the "children" edge to the User entity by IDs.
func (uc *UserCreate) AddChildIDs(ids ...int64) *UserCreate {
	uc.mutation.AddChildIDs(ids...)
	return uc
}

// AddChildren adds the "children" edges to the User entity.
func (uc *UserCreate) AddChildren(u ...*User) *UserCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddChildIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedBy(); !ok {
		v := user.DefaultCreatedBy
		uc.mutation.SetCreatedBy(v)
	}
	if _, ok := uc.mutation.UpdatedBy(); !ok {
		v := user.DefaultUpdatedBy
		uc.mutation.SetUpdatedBy(v)
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.DeletedAt(); !ok {
		v := user.DefaultDeletedAt
		uc.mutation.SetDeletedAt(v)
	}
	if _, ok := uc.mutation.Name(); !ok {
		v := user.DefaultName
		uc.mutation.SetName(v)
	}
	if _, ok := uc.mutation.JpgURL(); !ok {
		v := user.DefaultJpgURL
		uc.mutation.SetJpgURL(v)
	}
	if _, ok := uc.mutation.Key(); !ok {
		v := user.DefaultKey
		uc.mutation.SetKey(v)
	}
	if _, ok := uc.mutation.Secret(); !ok {
		v := user.DefaultSecret
		uc.mutation.SetSecret(v)
	}
	if _, ok := uc.mutation.Phone(); !ok {
		v := user.DefaultPhone
		uc.mutation.SetPhone(v)
	}
	if _, ok := uc.mutation.Password(); !ok {
		v := user.DefaultPassword
		uc.mutation.SetPassword(v)
	}
	if _, ok := uc.mutation.IsFrozen(); !ok {
		v := user.DefaultIsFrozen
		uc.mutation.SetIsFrozen(v)
	}
	if _, ok := uc.mutation.UserType(); !ok {
		v := user.DefaultUserType
		uc.mutation.SetUserType(v)
	}
	if _, ok := uc.mutation.ParentID(); !ok {
		v := user.DefaultParentID
		uc.mutation.SetParentID(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "User.created_by"`)}
	}
	if _, ok := uc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "User.updated_by"`)}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "User.updated_at"`)}
	}
	if _, ok := uc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "User.deleted_at"`)}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`cep_ent: missing required field "User.name"`)}
	}
	if _, ok := uc.mutation.JpgURL(); !ok {
		return &ValidationError{Name: "jpg_url", err: errors.New(`cep_ent: missing required field "User.jpg_url"`)}
	}
	if _, ok := uc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`cep_ent: missing required field "User.key"`)}
	}
	if _, ok := uc.mutation.Secret(); !ok {
		return &ValidationError{Name: "secret", err: errors.New(`cep_ent: missing required field "User.secret"`)}
	}
	if _, ok := uc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`cep_ent: missing required field "User.phone"`)}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`cep_ent: missing required field "User.password"`)}
	}
	if _, ok := uc.mutation.IsFrozen(); !ok {
		return &ValidationError{Name: "is_frozen", err: errors.New(`cep_ent: missing required field "User.is_frozen"`)}
	}
	if _, ok := uc.mutation.UserType(); !ok {
		return &ValidationError{Name: "user_type", err: errors.New(`cep_ent: missing required field "User.user_type"`)}
	}
	if v, ok := uc.mutation.UserType(); ok {
		if err := user.UserTypeValidator(v); err != nil {
			return &ValidationError{Name: "user_type", err: fmt.Errorf(`cep_ent: validator failed for field "User.user_type": %w`, err)}
		}
	}
	if _, ok := uc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent_id", err: errors.New(`cep_ent: missing required field "User.parent_id"`)}
	}
	if _, ok := uc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent", err: errors.New(`cep_ent: missing required edge "User.parent"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.CreatedBy(); ok {
		_spec.SetField(user.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := uc.mutation.UpdatedBy(); ok {
		_spec.SetField(user.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.JpgURL(); ok {
		_spec.SetField(user.FieldJpgURL, field.TypeString, value)
		_node.JpgURL = value
	}
	if value, ok := uc.mutation.Key(); ok {
		_spec.SetField(user.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := uc.mutation.Secret(); ok {
		_spec.SetField(user.FieldSecret, field.TypeString, value)
		_node.Secret = value
	}
	if value, ok := uc.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.IsFrozen(); ok {
		_spec.SetField(user.FieldIsFrozen, field.TypeBool, value)
		_node.IsFrozen = value
	}
	if value, ok := uc.mutation.UserType(); ok {
		_spec.SetField(user.FieldUserType, field.TypeEnum, value)
		_node.UserType = value
	}
	if nodes := uc.mutation.VxAccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VxAccountsTable,
			Columns: []string{user.VxAccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vxaccount.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CollectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CollectsTable,
			Columns: []string{user.CollectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collect.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DevicesTable,
			Columns: []string{user.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ProfitSettingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProfitSettingsTable,
			Columns: []string{user.ProfitSettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profitsetting.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CostAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CostAccountTable,
			Columns: []string{user.CostAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(costaccount.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ProfitAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.ProfitAccountTable,
			Columns: []string{user.ProfitAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profitaccount.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CostBillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CostBillsTable,
			Columns: []string{user.CostBillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(costbill.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.EarnBillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.EarnBillsTable,
			Columns: []string{user.EarnBillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earnbill.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.MissionConsumeOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MissionConsumeOrdersTable,
			Columns: []string{user.MissionConsumeOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionconsumeorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.MissionProduceOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MissionProduceOrdersTable,
			Columns: []string{user.MissionProduceOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionproduceorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.RechargeOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RechargeOrdersTable,
			Columns: []string{user.RechargeOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rechargeorder.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.VxSocialsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VxSocialsTable,
			Columns: []string{user.VxSocialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vxsocial.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.MissionBatchesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MissionBatchesTable,
			Columns: []string{user.MissionBatchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missionbatch.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserDevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserDevicesTable,
			Columns: []string{user.UserDevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userdevice.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
