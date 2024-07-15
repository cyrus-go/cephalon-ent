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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// InviteUpdate is the builder for updating Invite entities.
type InviteUpdate struct {
	config
	hooks     []Hook
	mutation  *InviteMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the InviteUpdate builder.
func (iu *InviteUpdate) Where(ps ...predicate.Invite) *InviteUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetCreatedBy sets the "created_by" field.
func (iu *InviteUpdate) SetCreatedBy(i int64) *InviteUpdate {
	iu.mutation.ResetCreatedBy()
	iu.mutation.SetCreatedBy(i)
	return iu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (iu *InviteUpdate) SetNillableCreatedBy(i *int64) *InviteUpdate {
	if i != nil {
		iu.SetCreatedBy(*i)
	}
	return iu
}

// AddCreatedBy adds i to the "created_by" field.
func (iu *InviteUpdate) AddCreatedBy(i int64) *InviteUpdate {
	iu.mutation.AddCreatedBy(i)
	return iu
}

// SetUpdatedBy sets the "updated_by" field.
func (iu *InviteUpdate) SetUpdatedBy(i int64) *InviteUpdate {
	iu.mutation.ResetUpdatedBy()
	iu.mutation.SetUpdatedBy(i)
	return iu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (iu *InviteUpdate) SetNillableUpdatedBy(i *int64) *InviteUpdate {
	if i != nil {
		iu.SetUpdatedBy(*i)
	}
	return iu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (iu *InviteUpdate) AddUpdatedBy(i int64) *InviteUpdate {
	iu.mutation.AddUpdatedBy(i)
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *InviteUpdate) SetUpdatedAt(t time.Time) *InviteUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetDeletedAt sets the "deleted_at" field.
func (iu *InviteUpdate) SetDeletedAt(t time.Time) *InviteUpdate {
	iu.mutation.SetDeletedAt(t)
	return iu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iu *InviteUpdate) SetNillableDeletedAt(t *time.Time) *InviteUpdate {
	if t != nil {
		iu.SetDeletedAt(*t)
	}
	return iu
}

// SetInviteCode sets the "invite_code" field.
func (iu *InviteUpdate) SetInviteCode(s string) *InviteUpdate {
	iu.mutation.SetInviteCode(s)
	return iu
}

// SetNillableInviteCode sets the "invite_code" field if the given value is not nil.
func (iu *InviteUpdate) SetNillableInviteCode(s *string) *InviteUpdate {
	if s != nil {
		iu.SetInviteCode(*s)
	}
	return iu
}

// SetShareCep sets the "share_cep" field.
func (iu *InviteUpdate) SetShareCep(i int64) *InviteUpdate {
	iu.mutation.ResetShareCep()
	iu.mutation.SetShareCep(i)
	return iu
}

// SetNillableShareCep sets the "share_cep" field if the given value is not nil.
func (iu *InviteUpdate) SetNillableShareCep(i *int64) *InviteUpdate {
	if i != nil {
		iu.SetShareCep(*i)
	}
	return iu
}

// AddShareCep adds i to the "share_cep" field.
func (iu *InviteUpdate) AddShareCep(i int64) *InviteUpdate {
	iu.mutation.AddShareCep(i)
	return iu
}

// SetRegCep sets the "reg_cep" field.
func (iu *InviteUpdate) SetRegCep(i int64) *InviteUpdate {
	iu.mutation.ResetRegCep()
	iu.mutation.SetRegCep(i)
	return iu
}

// SetNillableRegCep sets the "reg_cep" field if the given value is not nil.
func (iu *InviteUpdate) SetNillableRegCep(i *int64) *InviteUpdate {
	if i != nil {
		iu.SetRegCep(*i)
	}
	return iu
}

// AddRegCep adds i to the "reg_cep" field.
func (iu *InviteUpdate) AddRegCep(i int64) *InviteUpdate {
	iu.mutation.AddRegCep(i)
	return iu
}

// SetFirstRechargeCep sets the "first_recharge_cep" field.
func (iu *InviteUpdate) SetFirstRechargeCep(i int64) *InviteUpdate {
	iu.mutation.ResetFirstRechargeCep()
	iu.mutation.SetFirstRechargeCep(i)
	return iu
}

// SetNillableFirstRechargeCep sets the "first_recharge_cep" field if the given value is not nil.
func (iu *InviteUpdate) SetNillableFirstRechargeCep(i *int64) *InviteUpdate {
	if i != nil {
		iu.SetFirstRechargeCep(*i)
	}
	return iu
}

// AddFirstRechargeCep adds i to the "first_recharge_cep" field.
func (iu *InviteUpdate) AddFirstRechargeCep(i int64) *InviteUpdate {
	iu.mutation.AddFirstRechargeCep(i)
	return iu
}

// SetType sets the "type" field.
func (iu *InviteUpdate) SetType(et enums.InviteType) *InviteUpdate {
	iu.mutation.SetType(et)
	return iu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (iu *InviteUpdate) SetNillableType(et *enums.InviteType) *InviteUpdate {
	if et != nil {
		iu.SetType(*et)
	}
	return iu
}

// SetUserID sets the "user_id" field.
func (iu *InviteUpdate) SetUserID(i int64) *InviteUpdate {
	iu.mutation.SetUserID(i)
	return iu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (iu *InviteUpdate) SetNillableUserID(i *int64) *InviteUpdate {
	if i != nil {
		iu.SetUserID(*i)
	}
	return iu
}

// SetCampaignID sets the "campaign_id" field.
func (iu *InviteUpdate) SetCampaignID(i int64) *InviteUpdate {
	iu.mutation.SetCampaignID(i)
	return iu
}

// SetNillableCampaignID sets the "campaign_id" field if the given value is not nil.
func (iu *InviteUpdate) SetNillableCampaignID(i *int64) *InviteUpdate {
	if i != nil {
		iu.SetCampaignID(*i)
	}
	return iu
}

// SetUser sets the "user" edge to the User entity.
func (iu *InviteUpdate) SetUser(u *User) *InviteUpdate {
	return iu.SetUserID(u.ID)
}

// SetCampaign sets the "campaign" edge to the Campaign entity.
func (iu *InviteUpdate) SetCampaign(c *Campaign) *InviteUpdate {
	return iu.SetCampaignID(c.ID)
}

// AddBillIDs adds the "bills" edge to the Bill entity by IDs.
func (iu *InviteUpdate) AddBillIDs(ids ...int64) *InviteUpdate {
	iu.mutation.AddBillIDs(ids...)
	return iu
}

// AddBills adds the "bills" edges to the Bill entity.
func (iu *InviteUpdate) AddBills(b ...*Bill) *InviteUpdate {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return iu.AddBillIDs(ids...)
}

// Mutation returns the InviteMutation object of the builder.
func (iu *InviteUpdate) Mutation() *InviteMutation {
	return iu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (iu *InviteUpdate) ClearUser() *InviteUpdate {
	iu.mutation.ClearUser()
	return iu
}

// ClearCampaign clears the "campaign" edge to the Campaign entity.
func (iu *InviteUpdate) ClearCampaign() *InviteUpdate {
	iu.mutation.ClearCampaign()
	return iu
}

// ClearBills clears all "bills" edges to the Bill entity.
func (iu *InviteUpdate) ClearBills() *InviteUpdate {
	iu.mutation.ClearBills()
	return iu
}

// RemoveBillIDs removes the "bills" edge to Bill entities by IDs.
func (iu *InviteUpdate) RemoveBillIDs(ids ...int64) *InviteUpdate {
	iu.mutation.RemoveBillIDs(ids...)
	return iu
}

// RemoveBills removes "bills" edges to Bill entities.
func (iu *InviteUpdate) RemoveBills(b ...*Bill) *InviteUpdate {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return iu.RemoveBillIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *InviteUpdate) Save(ctx context.Context) (int, error) {
	iu.defaults()
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InviteUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InviteUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InviteUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *InviteUpdate) defaults() {
	if _, ok := iu.mutation.UpdatedAt(); !ok {
		v := invite.UpdateDefaultUpdatedAt()
		iu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *InviteUpdate) check() error {
	if v, ok := iu.mutation.GetType(); ok {
		if err := invite.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`cep_ent: validator failed for field "Invite.type": %w`, err)}
		}
	}
	if _, ok := iu.mutation.UserID(); iu.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "Invite.user"`)
	}
	if _, ok := iu.mutation.CampaignID(); iu.mutation.CampaignCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "Invite.campaign"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (iu *InviteUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *InviteUpdate {
	iu.modifiers = append(iu.modifiers, modifiers...)
	return iu
}

func (iu *InviteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(invite.Table, invite.Columns, sqlgraph.NewFieldSpec(invite.FieldID, field.TypeInt64))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.CreatedBy(); ok {
		_spec.SetField(invite.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := iu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(invite.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := iu.mutation.UpdatedBy(); ok {
		_spec.SetField(invite.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := iu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(invite.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.SetField(invite.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iu.mutation.DeletedAt(); ok {
		_spec.SetField(invite.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := iu.mutation.InviteCode(); ok {
		_spec.SetField(invite.FieldInviteCode, field.TypeString, value)
	}
	if value, ok := iu.mutation.ShareCep(); ok {
		_spec.SetField(invite.FieldShareCep, field.TypeInt64, value)
	}
	if value, ok := iu.mutation.AddedShareCep(); ok {
		_spec.AddField(invite.FieldShareCep, field.TypeInt64, value)
	}
	if value, ok := iu.mutation.RegCep(); ok {
		_spec.SetField(invite.FieldRegCep, field.TypeInt64, value)
	}
	if value, ok := iu.mutation.AddedRegCep(); ok {
		_spec.AddField(invite.FieldRegCep, field.TypeInt64, value)
	}
	if value, ok := iu.mutation.FirstRechargeCep(); ok {
		_spec.SetField(invite.FieldFirstRechargeCep, field.TypeInt64, value)
	}
	if value, ok := iu.mutation.AddedFirstRechargeCep(); ok {
		_spec.AddField(invite.FieldFirstRechargeCep, field.TypeInt64, value)
	}
	if value, ok := iu.mutation.GetType(); ok {
		_spec.SetField(invite.FieldType, field.TypeEnum, value)
	}
	if iu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.CampaignCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.CampaignIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.BillsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedBillsIDs(); len(nodes) > 0 && !iu.mutation.BillsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.BillsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(iu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invite.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// InviteUpdateOne is the builder for updating a single Invite entity.
type InviteUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *InviteMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (iuo *InviteUpdateOne) SetCreatedBy(i int64) *InviteUpdateOne {
	iuo.mutation.ResetCreatedBy()
	iuo.mutation.SetCreatedBy(i)
	return iuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (iuo *InviteUpdateOne) SetNillableCreatedBy(i *int64) *InviteUpdateOne {
	if i != nil {
		iuo.SetCreatedBy(*i)
	}
	return iuo
}

// AddCreatedBy adds i to the "created_by" field.
func (iuo *InviteUpdateOne) AddCreatedBy(i int64) *InviteUpdateOne {
	iuo.mutation.AddCreatedBy(i)
	return iuo
}

// SetUpdatedBy sets the "updated_by" field.
func (iuo *InviteUpdateOne) SetUpdatedBy(i int64) *InviteUpdateOne {
	iuo.mutation.ResetUpdatedBy()
	iuo.mutation.SetUpdatedBy(i)
	return iuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (iuo *InviteUpdateOne) SetNillableUpdatedBy(i *int64) *InviteUpdateOne {
	if i != nil {
		iuo.SetUpdatedBy(*i)
	}
	return iuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (iuo *InviteUpdateOne) AddUpdatedBy(i int64) *InviteUpdateOne {
	iuo.mutation.AddUpdatedBy(i)
	return iuo
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *InviteUpdateOne) SetUpdatedAt(t time.Time) *InviteUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetDeletedAt sets the "deleted_at" field.
func (iuo *InviteUpdateOne) SetDeletedAt(t time.Time) *InviteUpdateOne {
	iuo.mutation.SetDeletedAt(t)
	return iuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iuo *InviteUpdateOne) SetNillableDeletedAt(t *time.Time) *InviteUpdateOne {
	if t != nil {
		iuo.SetDeletedAt(*t)
	}
	return iuo
}

// SetInviteCode sets the "invite_code" field.
func (iuo *InviteUpdateOne) SetInviteCode(s string) *InviteUpdateOne {
	iuo.mutation.SetInviteCode(s)
	return iuo
}

// SetNillableInviteCode sets the "invite_code" field if the given value is not nil.
func (iuo *InviteUpdateOne) SetNillableInviteCode(s *string) *InviteUpdateOne {
	if s != nil {
		iuo.SetInviteCode(*s)
	}
	return iuo
}

// SetShareCep sets the "share_cep" field.
func (iuo *InviteUpdateOne) SetShareCep(i int64) *InviteUpdateOne {
	iuo.mutation.ResetShareCep()
	iuo.mutation.SetShareCep(i)
	return iuo
}

// SetNillableShareCep sets the "share_cep" field if the given value is not nil.
func (iuo *InviteUpdateOne) SetNillableShareCep(i *int64) *InviteUpdateOne {
	if i != nil {
		iuo.SetShareCep(*i)
	}
	return iuo
}

// AddShareCep adds i to the "share_cep" field.
func (iuo *InviteUpdateOne) AddShareCep(i int64) *InviteUpdateOne {
	iuo.mutation.AddShareCep(i)
	return iuo
}

// SetRegCep sets the "reg_cep" field.
func (iuo *InviteUpdateOne) SetRegCep(i int64) *InviteUpdateOne {
	iuo.mutation.ResetRegCep()
	iuo.mutation.SetRegCep(i)
	return iuo
}

// SetNillableRegCep sets the "reg_cep" field if the given value is not nil.
func (iuo *InviteUpdateOne) SetNillableRegCep(i *int64) *InviteUpdateOne {
	if i != nil {
		iuo.SetRegCep(*i)
	}
	return iuo
}

// AddRegCep adds i to the "reg_cep" field.
func (iuo *InviteUpdateOne) AddRegCep(i int64) *InviteUpdateOne {
	iuo.mutation.AddRegCep(i)
	return iuo
}

// SetFirstRechargeCep sets the "first_recharge_cep" field.
func (iuo *InviteUpdateOne) SetFirstRechargeCep(i int64) *InviteUpdateOne {
	iuo.mutation.ResetFirstRechargeCep()
	iuo.mutation.SetFirstRechargeCep(i)
	return iuo
}

// SetNillableFirstRechargeCep sets the "first_recharge_cep" field if the given value is not nil.
func (iuo *InviteUpdateOne) SetNillableFirstRechargeCep(i *int64) *InviteUpdateOne {
	if i != nil {
		iuo.SetFirstRechargeCep(*i)
	}
	return iuo
}

// AddFirstRechargeCep adds i to the "first_recharge_cep" field.
func (iuo *InviteUpdateOne) AddFirstRechargeCep(i int64) *InviteUpdateOne {
	iuo.mutation.AddFirstRechargeCep(i)
	return iuo
}

// SetType sets the "type" field.
func (iuo *InviteUpdateOne) SetType(et enums.InviteType) *InviteUpdateOne {
	iuo.mutation.SetType(et)
	return iuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (iuo *InviteUpdateOne) SetNillableType(et *enums.InviteType) *InviteUpdateOne {
	if et != nil {
		iuo.SetType(*et)
	}
	return iuo
}

// SetUserID sets the "user_id" field.
func (iuo *InviteUpdateOne) SetUserID(i int64) *InviteUpdateOne {
	iuo.mutation.SetUserID(i)
	return iuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (iuo *InviteUpdateOne) SetNillableUserID(i *int64) *InviteUpdateOne {
	if i != nil {
		iuo.SetUserID(*i)
	}
	return iuo
}

// SetCampaignID sets the "campaign_id" field.
func (iuo *InviteUpdateOne) SetCampaignID(i int64) *InviteUpdateOne {
	iuo.mutation.SetCampaignID(i)
	return iuo
}

// SetNillableCampaignID sets the "campaign_id" field if the given value is not nil.
func (iuo *InviteUpdateOne) SetNillableCampaignID(i *int64) *InviteUpdateOne {
	if i != nil {
		iuo.SetCampaignID(*i)
	}
	return iuo
}

// SetUser sets the "user" edge to the User entity.
func (iuo *InviteUpdateOne) SetUser(u *User) *InviteUpdateOne {
	return iuo.SetUserID(u.ID)
}

// SetCampaign sets the "campaign" edge to the Campaign entity.
func (iuo *InviteUpdateOne) SetCampaign(c *Campaign) *InviteUpdateOne {
	return iuo.SetCampaignID(c.ID)
}

// AddBillIDs adds the "bills" edge to the Bill entity by IDs.
func (iuo *InviteUpdateOne) AddBillIDs(ids ...int64) *InviteUpdateOne {
	iuo.mutation.AddBillIDs(ids...)
	return iuo
}

// AddBills adds the "bills" edges to the Bill entity.
func (iuo *InviteUpdateOne) AddBills(b ...*Bill) *InviteUpdateOne {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return iuo.AddBillIDs(ids...)
}

// Mutation returns the InviteMutation object of the builder.
func (iuo *InviteUpdateOne) Mutation() *InviteMutation {
	return iuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (iuo *InviteUpdateOne) ClearUser() *InviteUpdateOne {
	iuo.mutation.ClearUser()
	return iuo
}

// ClearCampaign clears the "campaign" edge to the Campaign entity.
func (iuo *InviteUpdateOne) ClearCampaign() *InviteUpdateOne {
	iuo.mutation.ClearCampaign()
	return iuo
}

// ClearBills clears all "bills" edges to the Bill entity.
func (iuo *InviteUpdateOne) ClearBills() *InviteUpdateOne {
	iuo.mutation.ClearBills()
	return iuo
}

// RemoveBillIDs removes the "bills" edge to Bill entities by IDs.
func (iuo *InviteUpdateOne) RemoveBillIDs(ids ...int64) *InviteUpdateOne {
	iuo.mutation.RemoveBillIDs(ids...)
	return iuo
}

// RemoveBills removes "bills" edges to Bill entities.
func (iuo *InviteUpdateOne) RemoveBills(b ...*Bill) *InviteUpdateOne {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return iuo.RemoveBillIDs(ids...)
}

// Where appends a list predicates to the InviteUpdate builder.
func (iuo *InviteUpdateOne) Where(ps ...predicate.Invite) *InviteUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *InviteUpdateOne) Select(field string, fields ...string) *InviteUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Invite entity.
func (iuo *InviteUpdateOne) Save(ctx context.Context) (*Invite, error) {
	iuo.defaults()
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InviteUpdateOne) SaveX(ctx context.Context) *Invite {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *InviteUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InviteUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *InviteUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdatedAt(); !ok {
		v := invite.UpdateDefaultUpdatedAt()
		iuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *InviteUpdateOne) check() error {
	if v, ok := iuo.mutation.GetType(); ok {
		if err := invite.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`cep_ent: validator failed for field "Invite.type": %w`, err)}
		}
	}
	if _, ok := iuo.mutation.UserID(); iuo.mutation.UserCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "Invite.user"`)
	}
	if _, ok := iuo.mutation.CampaignID(); iuo.mutation.CampaignCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "Invite.campaign"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (iuo *InviteUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *InviteUpdateOne {
	iuo.modifiers = append(iuo.modifiers, modifiers...)
	return iuo
}

func (iuo *InviteUpdateOne) sqlSave(ctx context.Context) (_node *Invite, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(invite.Table, invite.Columns, sqlgraph.NewFieldSpec(invite.FieldID, field.TypeInt64))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "Invite.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invite.FieldID)
		for _, f := range fields {
			if !invite.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != invite.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.CreatedBy(); ok {
		_spec.SetField(invite.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := iuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(invite.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := iuo.mutation.UpdatedBy(); ok {
		_spec.SetField(invite.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := iuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(invite.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.SetField(invite.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.DeletedAt(); ok {
		_spec.SetField(invite.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.InviteCode(); ok {
		_spec.SetField(invite.FieldInviteCode, field.TypeString, value)
	}
	if value, ok := iuo.mutation.ShareCep(); ok {
		_spec.SetField(invite.FieldShareCep, field.TypeInt64, value)
	}
	if value, ok := iuo.mutation.AddedShareCep(); ok {
		_spec.AddField(invite.FieldShareCep, field.TypeInt64, value)
	}
	if value, ok := iuo.mutation.RegCep(); ok {
		_spec.SetField(invite.FieldRegCep, field.TypeInt64, value)
	}
	if value, ok := iuo.mutation.AddedRegCep(); ok {
		_spec.AddField(invite.FieldRegCep, field.TypeInt64, value)
	}
	if value, ok := iuo.mutation.FirstRechargeCep(); ok {
		_spec.SetField(invite.FieldFirstRechargeCep, field.TypeInt64, value)
	}
	if value, ok := iuo.mutation.AddedFirstRechargeCep(); ok {
		_spec.AddField(invite.FieldFirstRechargeCep, field.TypeInt64, value)
	}
	if value, ok := iuo.mutation.GetType(); ok {
		_spec.SetField(invite.FieldType, field.TypeEnum, value)
	}
	if iuo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.CampaignCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.CampaignIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.BillsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedBillsIDs(); len(nodes) > 0 && !iuo.mutation.BillsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.BillsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(iuo.modifiers...)
	_node = &Invite{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invite.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
