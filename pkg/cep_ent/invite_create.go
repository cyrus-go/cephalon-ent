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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/bill"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/campaign"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/invite"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// InviteCreate is the builder for creating a Invite entity.
type InviteCreate struct {
	config
	mutation *InviteMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (ic *InviteCreate) SetCreatedBy(i int64) *InviteCreate {
	ic.mutation.SetCreatedBy(i)
	return ic
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ic *InviteCreate) SetNillableCreatedBy(i *int64) *InviteCreate {
	if i != nil {
		ic.SetCreatedBy(*i)
	}
	return ic
}

// SetUpdatedBy sets the "updated_by" field.
func (ic *InviteCreate) SetUpdatedBy(i int64) *InviteCreate {
	ic.mutation.SetUpdatedBy(i)
	return ic
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ic *InviteCreate) SetNillableUpdatedBy(i *int64) *InviteCreate {
	if i != nil {
		ic.SetUpdatedBy(*i)
	}
	return ic
}

// SetCreatedAt sets the "created_at" field.
func (ic *InviteCreate) SetCreatedAt(t time.Time) *InviteCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *InviteCreate) SetNillableCreatedAt(t *time.Time) *InviteCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *InviteCreate) SetUpdatedAt(t time.Time) *InviteCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *InviteCreate) SetNillableUpdatedAt(t *time.Time) *InviteCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetDeletedAt sets the "deleted_at" field.
func (ic *InviteCreate) SetDeletedAt(t time.Time) *InviteCreate {
	ic.mutation.SetDeletedAt(t)
	return ic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ic *InviteCreate) SetNillableDeletedAt(t *time.Time) *InviteCreate {
	if t != nil {
		ic.SetDeletedAt(*t)
	}
	return ic
}

// SetInviteCode sets the "invite_code" field.
func (ic *InviteCreate) SetInviteCode(s string) *InviteCreate {
	ic.mutation.SetInviteCode(s)
	return ic
}

// SetNillableInviteCode sets the "invite_code" field if the given value is not nil.
func (ic *InviteCreate) SetNillableInviteCode(s *string) *InviteCreate {
	if s != nil {
		ic.SetInviteCode(*s)
	}
	return ic
}

// SetShareCep sets the "share_cep" field.
func (ic *InviteCreate) SetShareCep(i int64) *InviteCreate {
	ic.mutation.SetShareCep(i)
	return ic
}

// SetNillableShareCep sets the "share_cep" field if the given value is not nil.
func (ic *InviteCreate) SetNillableShareCep(i *int64) *InviteCreate {
	if i != nil {
		ic.SetShareCep(*i)
	}
	return ic
}

// SetRegCep sets the "reg_cep" field.
func (ic *InviteCreate) SetRegCep(i int64) *InviteCreate {
	ic.mutation.SetRegCep(i)
	return ic
}

// SetNillableRegCep sets the "reg_cep" field if the given value is not nil.
func (ic *InviteCreate) SetNillableRegCep(i *int64) *InviteCreate {
	if i != nil {
		ic.SetRegCep(*i)
	}
	return ic
}

// SetFirstRechargeCep sets the "first_recharge_cep" field.
func (ic *InviteCreate) SetFirstRechargeCep(i int64) *InviteCreate {
	ic.mutation.SetFirstRechargeCep(i)
	return ic
}

// SetNillableFirstRechargeCep sets the "first_recharge_cep" field if the given value is not nil.
func (ic *InviteCreate) SetNillableFirstRechargeCep(i *int64) *InviteCreate {
	if i != nil {
		ic.SetFirstRechargeCep(*i)
	}
	return ic
}

// SetType sets the "type" field.
func (ic *InviteCreate) SetType(et enums.InviteType) *InviteCreate {
	ic.mutation.SetType(et)
	return ic
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ic *InviteCreate) SetNillableType(et *enums.InviteType) *InviteCreate {
	if et != nil {
		ic.SetType(*et)
	}
	return ic
}

// SetUserID sets the "user_id" field.
func (ic *InviteCreate) SetUserID(i int64) *InviteCreate {
	ic.mutation.SetUserID(i)
	return ic
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ic *InviteCreate) SetNillableUserID(i *int64) *InviteCreate {
	if i != nil {
		ic.SetUserID(*i)
	}
	return ic
}

// SetCampaignID sets the "campaign_id" field.
func (ic *InviteCreate) SetCampaignID(i int64) *InviteCreate {
	ic.mutation.SetCampaignID(i)
	return ic
}

// SetNillableCampaignID sets the "campaign_id" field if the given value is not nil.
func (ic *InviteCreate) SetNillableCampaignID(i *int64) *InviteCreate {
	if i != nil {
		ic.SetCampaignID(*i)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *InviteCreate) SetID(i int64) *InviteCreate {
	ic.mutation.SetID(i)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *InviteCreate) SetNillableID(i *int64) *InviteCreate {
	if i != nil {
		ic.SetID(*i)
	}
	return ic
}

// SetUser sets the "user" edge to the User entity.
func (ic *InviteCreate) SetUser(u *User) *InviteCreate {
	return ic.SetUserID(u.ID)
}

// SetCampaign sets the "campaign" edge to the Campaign entity.
func (ic *InviteCreate) SetCampaign(c *Campaign) *InviteCreate {
	return ic.SetCampaignID(c.ID)
}

// AddBillIDs adds the "bills" edge to the Bill entity by IDs.
func (ic *InviteCreate) AddBillIDs(ids ...int64) *InviteCreate {
	ic.mutation.AddBillIDs(ids...)
	return ic
}

// AddBills adds the "bills" edges to the Bill entity.
func (ic *InviteCreate) AddBills(b ...*Bill) *InviteCreate {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ic.AddBillIDs(ids...)
}

// Mutation returns the InviteMutation object of the builder.
func (ic *InviteCreate) Mutation() *InviteMutation {
	return ic.mutation
}

// Save creates the Invite in the database.
func (ic *InviteCreate) Save(ctx context.Context) (*Invite, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InviteCreate) SaveX(ctx context.Context) *Invite {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InviteCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InviteCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *InviteCreate) defaults() {
	if _, ok := ic.mutation.CreatedBy(); !ok {
		v := invite.DefaultCreatedBy
		ic.mutation.SetCreatedBy(v)
	}
	if _, ok := ic.mutation.UpdatedBy(); !ok {
		v := invite.DefaultUpdatedBy
		ic.mutation.SetUpdatedBy(v)
	}
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := invite.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := invite.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
	if _, ok := ic.mutation.DeletedAt(); !ok {
		v := invite.DefaultDeletedAt
		ic.mutation.SetDeletedAt(v)
	}
	if _, ok := ic.mutation.InviteCode(); !ok {
		v := invite.DefaultInviteCode
		ic.mutation.SetInviteCode(v)
	}
	if _, ok := ic.mutation.ShareCep(); !ok {
		v := invite.DefaultShareCep
		ic.mutation.SetShareCep(v)
	}
	if _, ok := ic.mutation.RegCep(); !ok {
		v := invite.DefaultRegCep
		ic.mutation.SetRegCep(v)
	}
	if _, ok := ic.mutation.FirstRechargeCep(); !ok {
		v := invite.DefaultFirstRechargeCep
		ic.mutation.SetFirstRechargeCep(v)
	}
	if _, ok := ic.mutation.GetType(); !ok {
		v := invite.DefaultType
		ic.mutation.SetType(v)
	}
	if _, ok := ic.mutation.UserID(); !ok {
		v := invite.DefaultUserID
		ic.mutation.SetUserID(v)
	}
	if _, ok := ic.mutation.CampaignID(); !ok {
		v := invite.DefaultCampaignID
		ic.mutation.SetCampaignID(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := invite.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InviteCreate) check() error {
	if _, ok := ic.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cep_ent: missing required field "Invite.created_by"`)}
	}
	if _, ok := ic.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cep_ent: missing required field "Invite.updated_by"`)}
	}
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cep_ent: missing required field "Invite.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cep_ent: missing required field "Invite.updated_at"`)}
	}
	if _, ok := ic.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`cep_ent: missing required field "Invite.deleted_at"`)}
	}
	if _, ok := ic.mutation.InviteCode(); !ok {
		return &ValidationError{Name: "invite_code", err: errors.New(`cep_ent: missing required field "Invite.invite_code"`)}
	}
	if _, ok := ic.mutation.ShareCep(); !ok {
		return &ValidationError{Name: "share_cep", err: errors.New(`cep_ent: missing required field "Invite.share_cep"`)}
	}
	if _, ok := ic.mutation.RegCep(); !ok {
		return &ValidationError{Name: "reg_cep", err: errors.New(`cep_ent: missing required field "Invite.reg_cep"`)}
	}
	if _, ok := ic.mutation.FirstRechargeCep(); !ok {
		return &ValidationError{Name: "first_recharge_cep", err: errors.New(`cep_ent: missing required field "Invite.first_recharge_cep"`)}
	}
	if _, ok := ic.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`cep_ent: missing required field "Invite.type"`)}
	}
	if v, ok := ic.mutation.GetType(); ok {
		if err := invite.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`cep_ent: validator failed for field "Invite.type": %w`, err)}
		}
	}
	if _, ok := ic.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`cep_ent: missing required field "Invite.user_id"`)}
	}
	if _, ok := ic.mutation.CampaignID(); !ok {
		return &ValidationError{Name: "campaign_id", err: errors.New(`cep_ent: missing required field "Invite.campaign_id"`)}
	}
	if _, ok := ic.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`cep_ent: missing required edge "Invite.user"`)}
	}
	if _, ok := ic.mutation.CampaignID(); !ok {
		return &ValidationError{Name: "campaign", err: errors.New(`cep_ent: missing required edge "Invite.campaign"`)}
	}
	return nil
}

func (ic *InviteCreate) sqlSave(ctx context.Context) (*Invite, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *InviteCreate) createSpec() (*Invite, *sqlgraph.CreateSpec) {
	var (
		_node = &Invite{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(invite.Table, sqlgraph.NewFieldSpec(invite.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = ic.conflict
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ic.mutation.CreatedBy(); ok {
		_spec.SetField(invite.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := ic.mutation.UpdatedBy(); ok {
		_spec.SetField(invite.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(invite.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(invite.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.DeletedAt(); ok {
		_spec.SetField(invite.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ic.mutation.InviteCode(); ok {
		_spec.SetField(invite.FieldInviteCode, field.TypeString, value)
		_node.InviteCode = value
	}
	if value, ok := ic.mutation.ShareCep(); ok {
		_spec.SetField(invite.FieldShareCep, field.TypeInt64, value)
		_node.ShareCep = value
	}
	if value, ok := ic.mutation.RegCep(); ok {
		_spec.SetField(invite.FieldRegCep, field.TypeInt64, value)
		_node.RegCep = value
	}
	if value, ok := ic.mutation.FirstRechargeCep(); ok {
		_spec.SetField(invite.FieldFirstRechargeCep, field.TypeInt64, value)
		_node.FirstRechargeCep = value
	}
	if value, ok := ic.mutation.GetType(); ok {
		_spec.SetField(invite.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if nodes := ic.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invite.UserTable,
			Columns: []string{invite.UserColumn},
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
	if nodes := ic.mutation.CampaignIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invite.CampaignTable,
			Columns: []string{invite.CampaignColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(campaign.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CampaignID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   invite.BillsTable,
			Columns: []string{invite.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt64),
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
//	client.Invite.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.InviteUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (ic *InviteCreate) OnConflict(opts ...sql.ConflictOption) *InviteUpsertOne {
	ic.conflict = opts
	return &InviteUpsertOne{
		create: ic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Invite.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ic *InviteCreate) OnConflictColumns(columns ...string) *InviteUpsertOne {
	ic.conflict = append(ic.conflict, sql.ConflictColumns(columns...))
	return &InviteUpsertOne{
		create: ic,
	}
}

type (
	// InviteUpsertOne is the builder for "upsert"-ing
	//  one Invite node.
	InviteUpsertOne struct {
		create *InviteCreate
	}

	// InviteUpsert is the "OnConflict" setter.
	InviteUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *InviteUpsert) SetCreatedBy(v int64) *InviteUpsert {
	u.Set(invite.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *InviteUpsert) UpdateCreatedBy() *InviteUpsert {
	u.SetExcluded(invite.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *InviteUpsert) AddCreatedBy(v int64) *InviteUpsert {
	u.Add(invite.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *InviteUpsert) SetUpdatedBy(v int64) *InviteUpsert {
	u.Set(invite.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *InviteUpsert) UpdateUpdatedBy() *InviteUpsert {
	u.SetExcluded(invite.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *InviteUpsert) AddUpdatedBy(v int64) *InviteUpsert {
	u.Add(invite.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *InviteUpsert) SetUpdatedAt(v time.Time) *InviteUpsert {
	u.Set(invite.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *InviteUpsert) UpdateUpdatedAt() *InviteUpsert {
	u.SetExcluded(invite.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *InviteUpsert) SetDeletedAt(v time.Time) *InviteUpsert {
	u.Set(invite.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *InviteUpsert) UpdateDeletedAt() *InviteUpsert {
	u.SetExcluded(invite.FieldDeletedAt)
	return u
}

// SetInviteCode sets the "invite_code" field.
func (u *InviteUpsert) SetInviteCode(v string) *InviteUpsert {
	u.Set(invite.FieldInviteCode, v)
	return u
}

// UpdateInviteCode sets the "invite_code" field to the value that was provided on create.
func (u *InviteUpsert) UpdateInviteCode() *InviteUpsert {
	u.SetExcluded(invite.FieldInviteCode)
	return u
}

// SetShareCep sets the "share_cep" field.
func (u *InviteUpsert) SetShareCep(v int64) *InviteUpsert {
	u.Set(invite.FieldShareCep, v)
	return u
}

// UpdateShareCep sets the "share_cep" field to the value that was provided on create.
func (u *InviteUpsert) UpdateShareCep() *InviteUpsert {
	u.SetExcluded(invite.FieldShareCep)
	return u
}

// AddShareCep adds v to the "share_cep" field.
func (u *InviteUpsert) AddShareCep(v int64) *InviteUpsert {
	u.Add(invite.FieldShareCep, v)
	return u
}

// SetRegCep sets the "reg_cep" field.
func (u *InviteUpsert) SetRegCep(v int64) *InviteUpsert {
	u.Set(invite.FieldRegCep, v)
	return u
}

// UpdateRegCep sets the "reg_cep" field to the value that was provided on create.
func (u *InviteUpsert) UpdateRegCep() *InviteUpsert {
	u.SetExcluded(invite.FieldRegCep)
	return u
}

// AddRegCep adds v to the "reg_cep" field.
func (u *InviteUpsert) AddRegCep(v int64) *InviteUpsert {
	u.Add(invite.FieldRegCep, v)
	return u
}

// SetFirstRechargeCep sets the "first_recharge_cep" field.
func (u *InviteUpsert) SetFirstRechargeCep(v int64) *InviteUpsert {
	u.Set(invite.FieldFirstRechargeCep, v)
	return u
}

// UpdateFirstRechargeCep sets the "first_recharge_cep" field to the value that was provided on create.
func (u *InviteUpsert) UpdateFirstRechargeCep() *InviteUpsert {
	u.SetExcluded(invite.FieldFirstRechargeCep)
	return u
}

// AddFirstRechargeCep adds v to the "first_recharge_cep" field.
func (u *InviteUpsert) AddFirstRechargeCep(v int64) *InviteUpsert {
	u.Add(invite.FieldFirstRechargeCep, v)
	return u
}

// SetType sets the "type" field.
func (u *InviteUpsert) SetType(v enums.InviteType) *InviteUpsert {
	u.Set(invite.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *InviteUpsert) UpdateType() *InviteUpsert {
	u.SetExcluded(invite.FieldType)
	return u
}

// SetUserID sets the "user_id" field.
func (u *InviteUpsert) SetUserID(v int64) *InviteUpsert {
	u.Set(invite.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *InviteUpsert) UpdateUserID() *InviteUpsert {
	u.SetExcluded(invite.FieldUserID)
	return u
}

// SetCampaignID sets the "campaign_id" field.
func (u *InviteUpsert) SetCampaignID(v int64) *InviteUpsert {
	u.Set(invite.FieldCampaignID, v)
	return u
}

// UpdateCampaignID sets the "campaign_id" field to the value that was provided on create.
func (u *InviteUpsert) UpdateCampaignID() *InviteUpsert {
	u.SetExcluded(invite.FieldCampaignID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Invite.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(invite.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *InviteUpsertOne) UpdateNewValues() *InviteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(invite.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(invite.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Invite.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *InviteUpsertOne) Ignore() *InviteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *InviteUpsertOne) DoNothing() *InviteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the InviteCreate.OnConflict
// documentation for more info.
func (u *InviteUpsertOne) Update(set func(*InviteUpsert)) *InviteUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&InviteUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *InviteUpsertOne) SetCreatedBy(v int64) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *InviteUpsertOne) AddCreatedBy(v int64) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *InviteUpsertOne) UpdateCreatedBy() *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *InviteUpsertOne) SetUpdatedBy(v int64) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *InviteUpsertOne) AddUpdatedBy(v int64) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *InviteUpsertOne) UpdateUpdatedBy() *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *InviteUpsertOne) SetUpdatedAt(v time.Time) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *InviteUpsertOne) UpdateUpdatedAt() *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *InviteUpsertOne) SetDeletedAt(v time.Time) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *InviteUpsertOne) UpdateDeletedAt() *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetInviteCode sets the "invite_code" field.
func (u *InviteUpsertOne) SetInviteCode(v string) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.SetInviteCode(v)
	})
}

// UpdateInviteCode sets the "invite_code" field to the value that was provided on create.
func (u *InviteUpsertOne) UpdateInviteCode() *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateInviteCode()
	})
}

// SetShareCep sets the "share_cep" field.
func (u *InviteUpsertOne) SetShareCep(v int64) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.SetShareCep(v)
	})
}

// AddShareCep adds v to the "share_cep" field.
func (u *InviteUpsertOne) AddShareCep(v int64) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.AddShareCep(v)
	})
}

// UpdateShareCep sets the "share_cep" field to the value that was provided on create.
func (u *InviteUpsertOne) UpdateShareCep() *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateShareCep()
	})
}

// SetRegCep sets the "reg_cep" field.
func (u *InviteUpsertOne) SetRegCep(v int64) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.SetRegCep(v)
	})
}

// AddRegCep adds v to the "reg_cep" field.
func (u *InviteUpsertOne) AddRegCep(v int64) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.AddRegCep(v)
	})
}

// UpdateRegCep sets the "reg_cep" field to the value that was provided on create.
func (u *InviteUpsertOne) UpdateRegCep() *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateRegCep()
	})
}

// SetFirstRechargeCep sets the "first_recharge_cep" field.
func (u *InviteUpsertOne) SetFirstRechargeCep(v int64) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.SetFirstRechargeCep(v)
	})
}

// AddFirstRechargeCep adds v to the "first_recharge_cep" field.
func (u *InviteUpsertOne) AddFirstRechargeCep(v int64) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.AddFirstRechargeCep(v)
	})
}

// UpdateFirstRechargeCep sets the "first_recharge_cep" field to the value that was provided on create.
func (u *InviteUpsertOne) UpdateFirstRechargeCep() *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateFirstRechargeCep()
	})
}

// SetType sets the "type" field.
func (u *InviteUpsertOne) SetType(v enums.InviteType) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *InviteUpsertOne) UpdateType() *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateType()
	})
}

// SetUserID sets the "user_id" field.
func (u *InviteUpsertOne) SetUserID(v int64) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *InviteUpsertOne) UpdateUserID() *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateUserID()
	})
}

// SetCampaignID sets the "campaign_id" field.
func (u *InviteUpsertOne) SetCampaignID(v int64) *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.SetCampaignID(v)
	})
}

// UpdateCampaignID sets the "campaign_id" field to the value that was provided on create.
func (u *InviteUpsertOne) UpdateCampaignID() *InviteUpsertOne {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateCampaignID()
	})
}

// Exec executes the query.
func (u *InviteUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for InviteCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *InviteUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *InviteUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *InviteUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// InviteCreateBulk is the builder for creating many Invite entities in bulk.
type InviteCreateBulk struct {
	config
	err      error
	builders []*InviteCreate
	conflict []sql.ConflictOption
}

// Save creates the Invite entities in the database.
func (icb *InviteCreateBulk) Save(ctx context.Context) ([]*Invite, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Invite, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InviteMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = icb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InviteCreateBulk) SaveX(ctx context.Context) []*Invite {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *InviteCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *InviteCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Invite.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.InviteUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (icb *InviteCreateBulk) OnConflict(opts ...sql.ConflictOption) *InviteUpsertBulk {
	icb.conflict = opts
	return &InviteUpsertBulk{
		create: icb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Invite.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (icb *InviteCreateBulk) OnConflictColumns(columns ...string) *InviteUpsertBulk {
	icb.conflict = append(icb.conflict, sql.ConflictColumns(columns...))
	return &InviteUpsertBulk{
		create: icb,
	}
}

// InviteUpsertBulk is the builder for "upsert"-ing
// a bulk of Invite nodes.
type InviteUpsertBulk struct {
	create *InviteCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Invite.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(invite.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *InviteUpsertBulk) UpdateNewValues() *InviteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(invite.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(invite.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Invite.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *InviteUpsertBulk) Ignore() *InviteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *InviteUpsertBulk) DoNothing() *InviteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the InviteCreateBulk.OnConflict
// documentation for more info.
func (u *InviteUpsertBulk) Update(set func(*InviteUpsert)) *InviteUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&InviteUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *InviteUpsertBulk) SetCreatedBy(v int64) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *InviteUpsertBulk) AddCreatedBy(v int64) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *InviteUpsertBulk) UpdateCreatedBy() *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *InviteUpsertBulk) SetUpdatedBy(v int64) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *InviteUpsertBulk) AddUpdatedBy(v int64) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *InviteUpsertBulk) UpdateUpdatedBy() *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *InviteUpsertBulk) SetUpdatedAt(v time.Time) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *InviteUpsertBulk) UpdateUpdatedAt() *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *InviteUpsertBulk) SetDeletedAt(v time.Time) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *InviteUpsertBulk) UpdateDeletedAt() *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetInviteCode sets the "invite_code" field.
func (u *InviteUpsertBulk) SetInviteCode(v string) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.SetInviteCode(v)
	})
}

// UpdateInviteCode sets the "invite_code" field to the value that was provided on create.
func (u *InviteUpsertBulk) UpdateInviteCode() *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateInviteCode()
	})
}

// SetShareCep sets the "share_cep" field.
func (u *InviteUpsertBulk) SetShareCep(v int64) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.SetShareCep(v)
	})
}

// AddShareCep adds v to the "share_cep" field.
func (u *InviteUpsertBulk) AddShareCep(v int64) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.AddShareCep(v)
	})
}

// UpdateShareCep sets the "share_cep" field to the value that was provided on create.
func (u *InviteUpsertBulk) UpdateShareCep() *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateShareCep()
	})
}

// SetRegCep sets the "reg_cep" field.
func (u *InviteUpsertBulk) SetRegCep(v int64) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.SetRegCep(v)
	})
}

// AddRegCep adds v to the "reg_cep" field.
func (u *InviteUpsertBulk) AddRegCep(v int64) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.AddRegCep(v)
	})
}

// UpdateRegCep sets the "reg_cep" field to the value that was provided on create.
func (u *InviteUpsertBulk) UpdateRegCep() *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateRegCep()
	})
}

// SetFirstRechargeCep sets the "first_recharge_cep" field.
func (u *InviteUpsertBulk) SetFirstRechargeCep(v int64) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.SetFirstRechargeCep(v)
	})
}

// AddFirstRechargeCep adds v to the "first_recharge_cep" field.
func (u *InviteUpsertBulk) AddFirstRechargeCep(v int64) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.AddFirstRechargeCep(v)
	})
}

// UpdateFirstRechargeCep sets the "first_recharge_cep" field to the value that was provided on create.
func (u *InviteUpsertBulk) UpdateFirstRechargeCep() *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateFirstRechargeCep()
	})
}

// SetType sets the "type" field.
func (u *InviteUpsertBulk) SetType(v enums.InviteType) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *InviteUpsertBulk) UpdateType() *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateType()
	})
}

// SetUserID sets the "user_id" field.
func (u *InviteUpsertBulk) SetUserID(v int64) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *InviteUpsertBulk) UpdateUserID() *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateUserID()
	})
}

// SetCampaignID sets the "campaign_id" field.
func (u *InviteUpsertBulk) SetCampaignID(v int64) *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.SetCampaignID(v)
	})
}

// UpdateCampaignID sets the "campaign_id" field to the value that was provided on create.
func (u *InviteUpsertBulk) UpdateCampaignID() *InviteUpsertBulk {
	return u.Update(func(s *InviteUpsert) {
		s.UpdateCampaignID()
	})
}

// Exec executes the query.
func (u *InviteUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("cep_ent: OnConflict was set for builder %d. Set it on the InviteCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("cep_ent: missing options for InviteCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *InviteUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
