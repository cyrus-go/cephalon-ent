// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/bill"
	"cephalon-ent/pkg/cep_ent/platformwallet"
	"cephalon-ent/pkg/cep_ent/predicate"
	"cephalon-ent/pkg/enums"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlatformWalletUpdate is the builder for updating PlatformWallet entities.
type PlatformWalletUpdate struct {
	config
	hooks    []Hook
	mutation *PlatformWalletMutation
}

// Where appends a list predicates to the PlatformWalletUpdate builder.
func (pwu *PlatformWalletUpdate) Where(ps ...predicate.PlatformWallet) *PlatformWalletUpdate {
	pwu.mutation.Where(ps...)
	return pwu
}

// SetCreatedBy sets the "created_by" field.
func (pwu *PlatformWalletUpdate) SetCreatedBy(i int64) *PlatformWalletUpdate {
	pwu.mutation.ResetCreatedBy()
	pwu.mutation.SetCreatedBy(i)
	return pwu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pwu *PlatformWalletUpdate) SetNillableCreatedBy(i *int64) *PlatformWalletUpdate {
	if i != nil {
		pwu.SetCreatedBy(*i)
	}
	return pwu
}

// AddCreatedBy adds i to the "created_by" field.
func (pwu *PlatformWalletUpdate) AddCreatedBy(i int64) *PlatformWalletUpdate {
	pwu.mutation.AddCreatedBy(i)
	return pwu
}

// SetUpdatedBy sets the "updated_by" field.
func (pwu *PlatformWalletUpdate) SetUpdatedBy(i int64) *PlatformWalletUpdate {
	pwu.mutation.ResetUpdatedBy()
	pwu.mutation.SetUpdatedBy(i)
	return pwu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pwu *PlatformWalletUpdate) SetNillableUpdatedBy(i *int64) *PlatformWalletUpdate {
	if i != nil {
		pwu.SetUpdatedBy(*i)
	}
	return pwu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (pwu *PlatformWalletUpdate) AddUpdatedBy(i int64) *PlatformWalletUpdate {
	pwu.mutation.AddUpdatedBy(i)
	return pwu
}

// SetUpdatedAt sets the "updated_at" field.
func (pwu *PlatformWalletUpdate) SetUpdatedAt(t time.Time) *PlatformWalletUpdate {
	pwu.mutation.SetUpdatedAt(t)
	return pwu
}

// SetDeletedAt sets the "deleted_at" field.
func (pwu *PlatformWalletUpdate) SetDeletedAt(t time.Time) *PlatformWalletUpdate {
	pwu.mutation.SetDeletedAt(t)
	return pwu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pwu *PlatformWalletUpdate) SetNillableDeletedAt(t *time.Time) *PlatformWalletUpdate {
	if t != nil {
		pwu.SetDeletedAt(*t)
	}
	return pwu
}

// SetType sets the "type" field.
func (pwu *PlatformWalletUpdate) SetType(ewt enums.PlatformWalletType) *PlatformWalletUpdate {
	pwu.mutation.SetType(ewt)
	return pwu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pwu *PlatformWalletUpdate) SetNillableType(ewt *enums.PlatformWalletType) *PlatformWalletUpdate {
	if ewt != nil {
		pwu.SetType(*ewt)
	}
	return pwu
}

// SetSumCep sets the "sum_cep" field.
func (pwu *PlatformWalletUpdate) SetSumCep(i int64) *PlatformWalletUpdate {
	pwu.mutation.ResetSumCep()
	pwu.mutation.SetSumCep(i)
	return pwu
}

// SetNillableSumCep sets the "sum_cep" field if the given value is not nil.
func (pwu *PlatformWalletUpdate) SetNillableSumCep(i *int64) *PlatformWalletUpdate {
	if i != nil {
		pwu.SetSumCep(*i)
	}
	return pwu
}

// AddSumCep adds i to the "sum_cep" field.
func (pwu *PlatformWalletUpdate) AddSumCep(i int64) *PlatformWalletUpdate {
	pwu.mutation.AddSumCep(i)
	return pwu
}

// SetCep sets the "cep" field.
func (pwu *PlatformWalletUpdate) SetCep(i int64) *PlatformWalletUpdate {
	pwu.mutation.ResetCep()
	pwu.mutation.SetCep(i)
	return pwu
}

// SetNillableCep sets the "cep" field if the given value is not nil.
func (pwu *PlatformWalletUpdate) SetNillableCep(i *int64) *PlatformWalletUpdate {
	if i != nil {
		pwu.SetCep(*i)
	}
	return pwu
}

// AddCep adds i to the "cep" field.
func (pwu *PlatformWalletUpdate) AddCep(i int64) *PlatformWalletUpdate {
	pwu.mutation.AddCep(i)
	return pwu
}

// AddBillIDs adds the "bills" edge to the Bill entity by IDs.
func (pwu *PlatformWalletUpdate) AddBillIDs(ids ...int64) *PlatformWalletUpdate {
	pwu.mutation.AddBillIDs(ids...)
	return pwu
}

// AddBills adds the "bills" edges to the Bill entity.
func (pwu *PlatformWalletUpdate) AddBills(b ...*Bill) *PlatformWalletUpdate {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pwu.AddBillIDs(ids...)
}

// Mutation returns the PlatformWalletMutation object of the builder.
func (pwu *PlatformWalletUpdate) Mutation() *PlatformWalletMutation {
	return pwu.mutation
}

// ClearBills clears all "bills" edges to the Bill entity.
func (pwu *PlatformWalletUpdate) ClearBills() *PlatformWalletUpdate {
	pwu.mutation.ClearBills()
	return pwu
}

// RemoveBillIDs removes the "bills" edge to Bill entities by IDs.
func (pwu *PlatformWalletUpdate) RemoveBillIDs(ids ...int64) *PlatformWalletUpdate {
	pwu.mutation.RemoveBillIDs(ids...)
	return pwu
}

// RemoveBills removes "bills" edges to Bill entities.
func (pwu *PlatformWalletUpdate) RemoveBills(b ...*Bill) *PlatformWalletUpdate {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pwu.RemoveBillIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pwu *PlatformWalletUpdate) Save(ctx context.Context) (int, error) {
	pwu.defaults()
	return withHooks(ctx, pwu.sqlSave, pwu.mutation, pwu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pwu *PlatformWalletUpdate) SaveX(ctx context.Context) int {
	affected, err := pwu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pwu *PlatformWalletUpdate) Exec(ctx context.Context) error {
	_, err := pwu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pwu *PlatformWalletUpdate) ExecX(ctx context.Context) {
	if err := pwu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pwu *PlatformWalletUpdate) defaults() {
	if _, ok := pwu.mutation.UpdatedAt(); !ok {
		v := platformwallet.UpdateDefaultUpdatedAt()
		pwu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pwu *PlatformWalletUpdate) check() error {
	if v, ok := pwu.mutation.GetType(); ok {
		if err := platformwallet.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`cep_ent: validator failed for field "PlatformWallet.type": %w`, err)}
		}
	}
	return nil
}

func (pwu *PlatformWalletUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pwu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(platformwallet.Table, platformwallet.Columns, sqlgraph.NewFieldSpec(platformwallet.FieldID, field.TypeInt64))
	if ps := pwu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pwu.mutation.CreatedBy(); ok {
		_spec.SetField(platformwallet.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := pwu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(platformwallet.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := pwu.mutation.UpdatedBy(); ok {
		_spec.SetField(platformwallet.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := pwu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(platformwallet.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := pwu.mutation.UpdatedAt(); ok {
		_spec.SetField(platformwallet.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pwu.mutation.DeletedAt(); ok {
		_spec.SetField(platformwallet.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := pwu.mutation.GetType(); ok {
		_spec.SetField(platformwallet.FieldType, field.TypeEnum, value)
	}
	if value, ok := pwu.mutation.SumCep(); ok {
		_spec.SetField(platformwallet.FieldSumCep, field.TypeInt64, value)
	}
	if value, ok := pwu.mutation.AddedSumCep(); ok {
		_spec.AddField(platformwallet.FieldSumCep, field.TypeInt64, value)
	}
	if value, ok := pwu.mutation.Cep(); ok {
		_spec.SetField(platformwallet.FieldCep, field.TypeInt64, value)
	}
	if value, ok := pwu.mutation.AddedCep(); ok {
		_spec.AddField(platformwallet.FieldCep, field.TypeInt64, value)
	}
	if pwu.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platformwallet.BillsTable,
			Columns: []string{platformwallet.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pwu.mutation.RemovedBillsIDs(); len(nodes) > 0 && !pwu.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platformwallet.BillsTable,
			Columns: []string{platformwallet.BillsColumn},
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
	if nodes := pwu.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platformwallet.BillsTable,
			Columns: []string{platformwallet.BillsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pwu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{platformwallet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pwu.mutation.done = true
	return n, nil
}

// PlatformWalletUpdateOne is the builder for updating a single PlatformWallet entity.
type PlatformWalletUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlatformWalletMutation
}

// SetCreatedBy sets the "created_by" field.
func (pwuo *PlatformWalletUpdateOne) SetCreatedBy(i int64) *PlatformWalletUpdateOne {
	pwuo.mutation.ResetCreatedBy()
	pwuo.mutation.SetCreatedBy(i)
	return pwuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pwuo *PlatformWalletUpdateOne) SetNillableCreatedBy(i *int64) *PlatformWalletUpdateOne {
	if i != nil {
		pwuo.SetCreatedBy(*i)
	}
	return pwuo
}

// AddCreatedBy adds i to the "created_by" field.
func (pwuo *PlatformWalletUpdateOne) AddCreatedBy(i int64) *PlatformWalletUpdateOne {
	pwuo.mutation.AddCreatedBy(i)
	return pwuo
}

// SetUpdatedBy sets the "updated_by" field.
func (pwuo *PlatformWalletUpdateOne) SetUpdatedBy(i int64) *PlatformWalletUpdateOne {
	pwuo.mutation.ResetUpdatedBy()
	pwuo.mutation.SetUpdatedBy(i)
	return pwuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pwuo *PlatformWalletUpdateOne) SetNillableUpdatedBy(i *int64) *PlatformWalletUpdateOne {
	if i != nil {
		pwuo.SetUpdatedBy(*i)
	}
	return pwuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (pwuo *PlatformWalletUpdateOne) AddUpdatedBy(i int64) *PlatformWalletUpdateOne {
	pwuo.mutation.AddUpdatedBy(i)
	return pwuo
}

// SetUpdatedAt sets the "updated_at" field.
func (pwuo *PlatformWalletUpdateOne) SetUpdatedAt(t time.Time) *PlatformWalletUpdateOne {
	pwuo.mutation.SetUpdatedAt(t)
	return pwuo
}

// SetDeletedAt sets the "deleted_at" field.
func (pwuo *PlatformWalletUpdateOne) SetDeletedAt(t time.Time) *PlatformWalletUpdateOne {
	pwuo.mutation.SetDeletedAt(t)
	return pwuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pwuo *PlatformWalletUpdateOne) SetNillableDeletedAt(t *time.Time) *PlatformWalletUpdateOne {
	if t != nil {
		pwuo.SetDeletedAt(*t)
	}
	return pwuo
}

// SetType sets the "type" field.
func (pwuo *PlatformWalletUpdateOne) SetType(ewt enums.PlatformWalletType) *PlatformWalletUpdateOne {
	pwuo.mutation.SetType(ewt)
	return pwuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pwuo *PlatformWalletUpdateOne) SetNillableType(ewt *enums.PlatformWalletType) *PlatformWalletUpdateOne {
	if ewt != nil {
		pwuo.SetType(*ewt)
	}
	return pwuo
}

// SetSumCep sets the "sum_cep" field.
func (pwuo *PlatformWalletUpdateOne) SetSumCep(i int64) *PlatformWalletUpdateOne {
	pwuo.mutation.ResetSumCep()
	pwuo.mutation.SetSumCep(i)
	return pwuo
}

// SetNillableSumCep sets the "sum_cep" field if the given value is not nil.
func (pwuo *PlatformWalletUpdateOne) SetNillableSumCep(i *int64) *PlatformWalletUpdateOne {
	if i != nil {
		pwuo.SetSumCep(*i)
	}
	return pwuo
}

// AddSumCep adds i to the "sum_cep" field.
func (pwuo *PlatformWalletUpdateOne) AddSumCep(i int64) *PlatformWalletUpdateOne {
	pwuo.mutation.AddSumCep(i)
	return pwuo
}

// SetCep sets the "cep" field.
func (pwuo *PlatformWalletUpdateOne) SetCep(i int64) *PlatformWalletUpdateOne {
	pwuo.mutation.ResetCep()
	pwuo.mutation.SetCep(i)
	return pwuo
}

// SetNillableCep sets the "cep" field if the given value is not nil.
func (pwuo *PlatformWalletUpdateOne) SetNillableCep(i *int64) *PlatformWalletUpdateOne {
	if i != nil {
		pwuo.SetCep(*i)
	}
	return pwuo
}

// AddCep adds i to the "cep" field.
func (pwuo *PlatformWalletUpdateOne) AddCep(i int64) *PlatformWalletUpdateOne {
	pwuo.mutation.AddCep(i)
	return pwuo
}

// AddBillIDs adds the "bills" edge to the Bill entity by IDs.
func (pwuo *PlatformWalletUpdateOne) AddBillIDs(ids ...int64) *PlatformWalletUpdateOne {
	pwuo.mutation.AddBillIDs(ids...)
	return pwuo
}

// AddBills adds the "bills" edges to the Bill entity.
func (pwuo *PlatformWalletUpdateOne) AddBills(b ...*Bill) *PlatformWalletUpdateOne {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pwuo.AddBillIDs(ids...)
}

// Mutation returns the PlatformWalletMutation object of the builder.
func (pwuo *PlatformWalletUpdateOne) Mutation() *PlatformWalletMutation {
	return pwuo.mutation
}

// ClearBills clears all "bills" edges to the Bill entity.
func (pwuo *PlatformWalletUpdateOne) ClearBills() *PlatformWalletUpdateOne {
	pwuo.mutation.ClearBills()
	return pwuo
}

// RemoveBillIDs removes the "bills" edge to Bill entities by IDs.
func (pwuo *PlatformWalletUpdateOne) RemoveBillIDs(ids ...int64) *PlatformWalletUpdateOne {
	pwuo.mutation.RemoveBillIDs(ids...)
	return pwuo
}

// RemoveBills removes "bills" edges to Bill entities.
func (pwuo *PlatformWalletUpdateOne) RemoveBills(b ...*Bill) *PlatformWalletUpdateOne {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pwuo.RemoveBillIDs(ids...)
}

// Where appends a list predicates to the PlatformWalletUpdate builder.
func (pwuo *PlatformWalletUpdateOne) Where(ps ...predicate.PlatformWallet) *PlatformWalletUpdateOne {
	pwuo.mutation.Where(ps...)
	return pwuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pwuo *PlatformWalletUpdateOne) Select(field string, fields ...string) *PlatformWalletUpdateOne {
	pwuo.fields = append([]string{field}, fields...)
	return pwuo
}

// Save executes the query and returns the updated PlatformWallet entity.
func (pwuo *PlatformWalletUpdateOne) Save(ctx context.Context) (*PlatformWallet, error) {
	pwuo.defaults()
	return withHooks(ctx, pwuo.sqlSave, pwuo.mutation, pwuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pwuo *PlatformWalletUpdateOne) SaveX(ctx context.Context) *PlatformWallet {
	node, err := pwuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pwuo *PlatformWalletUpdateOne) Exec(ctx context.Context) error {
	_, err := pwuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pwuo *PlatformWalletUpdateOne) ExecX(ctx context.Context) {
	if err := pwuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pwuo *PlatformWalletUpdateOne) defaults() {
	if _, ok := pwuo.mutation.UpdatedAt(); !ok {
		v := platformwallet.UpdateDefaultUpdatedAt()
		pwuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pwuo *PlatformWalletUpdateOne) check() error {
	if v, ok := pwuo.mutation.GetType(); ok {
		if err := platformwallet.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`cep_ent: validator failed for field "PlatformWallet.type": %w`, err)}
		}
	}
	return nil
}

func (pwuo *PlatformWalletUpdateOne) sqlSave(ctx context.Context) (_node *PlatformWallet, err error) {
	if err := pwuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(platformwallet.Table, platformwallet.Columns, sqlgraph.NewFieldSpec(platformwallet.FieldID, field.TypeInt64))
	id, ok := pwuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "PlatformWallet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pwuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, platformwallet.FieldID)
		for _, f := range fields {
			if !platformwallet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != platformwallet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pwuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pwuo.mutation.CreatedBy(); ok {
		_spec.SetField(platformwallet.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := pwuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(platformwallet.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := pwuo.mutation.UpdatedBy(); ok {
		_spec.SetField(platformwallet.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := pwuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(platformwallet.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := pwuo.mutation.UpdatedAt(); ok {
		_spec.SetField(platformwallet.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pwuo.mutation.DeletedAt(); ok {
		_spec.SetField(platformwallet.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := pwuo.mutation.GetType(); ok {
		_spec.SetField(platformwallet.FieldType, field.TypeEnum, value)
	}
	if value, ok := pwuo.mutation.SumCep(); ok {
		_spec.SetField(platformwallet.FieldSumCep, field.TypeInt64, value)
	}
	if value, ok := pwuo.mutation.AddedSumCep(); ok {
		_spec.AddField(platformwallet.FieldSumCep, field.TypeInt64, value)
	}
	if value, ok := pwuo.mutation.Cep(); ok {
		_spec.SetField(platformwallet.FieldCep, field.TypeInt64, value)
	}
	if value, ok := pwuo.mutation.AddedCep(); ok {
		_spec.AddField(platformwallet.FieldCep, field.TypeInt64, value)
	}
	if pwuo.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platformwallet.BillsTable,
			Columns: []string{platformwallet.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pwuo.mutation.RemovedBillsIDs(); len(nodes) > 0 && !pwuo.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platformwallet.BillsTable,
			Columns: []string{platformwallet.BillsColumn},
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
	if nodes := pwuo.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   platformwallet.BillsTable,
			Columns: []string{platformwallet.BillsColumn},
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
	_node = &PlatformWallet{config: pwuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pwuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{platformwallet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pwuo.mutation.done = true
	return _node, nil
}
