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
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/model"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/modelprice"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ModelPriceUpdate is the builder for updating ModelPrice entities.
type ModelPriceUpdate struct {
	config
	hooks     []Hook
	mutation  *ModelPriceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ModelPriceUpdate builder.
func (mpu *ModelPriceUpdate) Where(ps ...predicate.ModelPrice) *ModelPriceUpdate {
	mpu.mutation.Where(ps...)
	return mpu
}

// SetCreatedBy sets the "created_by" field.
func (mpu *ModelPriceUpdate) SetCreatedBy(i int64) *ModelPriceUpdate {
	mpu.mutation.ResetCreatedBy()
	mpu.mutation.SetCreatedBy(i)
	return mpu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mpu *ModelPriceUpdate) SetNillableCreatedBy(i *int64) *ModelPriceUpdate {
	if i != nil {
		mpu.SetCreatedBy(*i)
	}
	return mpu
}

// AddCreatedBy adds i to the "created_by" field.
func (mpu *ModelPriceUpdate) AddCreatedBy(i int64) *ModelPriceUpdate {
	mpu.mutation.AddCreatedBy(i)
	return mpu
}

// SetUpdatedBy sets the "updated_by" field.
func (mpu *ModelPriceUpdate) SetUpdatedBy(i int64) *ModelPriceUpdate {
	mpu.mutation.ResetUpdatedBy()
	mpu.mutation.SetUpdatedBy(i)
	return mpu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mpu *ModelPriceUpdate) SetNillableUpdatedBy(i *int64) *ModelPriceUpdate {
	if i != nil {
		mpu.SetUpdatedBy(*i)
	}
	return mpu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (mpu *ModelPriceUpdate) AddUpdatedBy(i int64) *ModelPriceUpdate {
	mpu.mutation.AddUpdatedBy(i)
	return mpu
}

// SetUpdatedAt sets the "updated_at" field.
func (mpu *ModelPriceUpdate) SetUpdatedAt(t time.Time) *ModelPriceUpdate {
	mpu.mutation.SetUpdatedAt(t)
	return mpu
}

// SetDeletedAt sets the "deleted_at" field.
func (mpu *ModelPriceUpdate) SetDeletedAt(t time.Time) *ModelPriceUpdate {
	mpu.mutation.SetDeletedAt(t)
	return mpu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mpu *ModelPriceUpdate) SetNillableDeletedAt(t *time.Time) *ModelPriceUpdate {
	if t != nil {
		mpu.SetDeletedAt(*t)
	}
	return mpu
}

// SetModelID sets the "model_id" field.
func (mpu *ModelPriceUpdate) SetModelID(i int64) *ModelPriceUpdate {
	mpu.mutation.SetModelID(i)
	return mpu
}

// SetInvokeType sets the "invoke_type" field.
func (mpu *ModelPriceUpdate) SetInvokeType(e enums.Model) *ModelPriceUpdate {
	mpu.mutation.SetInvokeType(e)
	return mpu
}

// SetNillableInvokeType sets the "invoke_type" field if the given value is not nil.
func (mpu *ModelPriceUpdate) SetNillableInvokeType(e *enums.Model) *ModelPriceUpdate {
	if e != nil {
		mpu.SetInvokeType(*e)
	}
	return mpu
}

// SetGpuVersion sets the "gpu_version" field.
func (mpu *ModelPriceUpdate) SetGpuVersion(ev enums.GpuVersion) *ModelPriceUpdate {
	mpu.mutation.SetGpuVersion(ev)
	return mpu
}

// SetNillableGpuVersion sets the "gpu_version" field if the given value is not nil.
func (mpu *ModelPriceUpdate) SetNillableGpuVersion(ev *enums.GpuVersion) *ModelPriceUpdate {
	if ev != nil {
		mpu.SetGpuVersion(*ev)
	}
	return mpu
}

// SetInputGpuPrice sets the "input_gpu_price" field.
func (mpu *ModelPriceUpdate) SetInputGpuPrice(i int) *ModelPriceUpdate {
	mpu.mutation.ResetInputGpuPrice()
	mpu.mutation.SetInputGpuPrice(i)
	return mpu
}

// SetNillableInputGpuPrice sets the "input_gpu_price" field if the given value is not nil.
func (mpu *ModelPriceUpdate) SetNillableInputGpuPrice(i *int) *ModelPriceUpdate {
	if i != nil {
		mpu.SetInputGpuPrice(*i)
	}
	return mpu
}

// AddInputGpuPrice adds i to the "input_gpu_price" field.
func (mpu *ModelPriceUpdate) AddInputGpuPrice(i int) *ModelPriceUpdate {
	mpu.mutation.AddInputGpuPrice(i)
	return mpu
}

// SetOutputGpuPrice sets the "output_gpu_price" field.
func (mpu *ModelPriceUpdate) SetOutputGpuPrice(i int) *ModelPriceUpdate {
	mpu.mutation.ResetOutputGpuPrice()
	mpu.mutation.SetOutputGpuPrice(i)
	return mpu
}

// SetNillableOutputGpuPrice sets the "output_gpu_price" field if the given value is not nil.
func (mpu *ModelPriceUpdate) SetNillableOutputGpuPrice(i *int) *ModelPriceUpdate {
	if i != nil {
		mpu.SetOutputGpuPrice(*i)
	}
	return mpu
}

// AddOutputGpuPrice adds i to the "output_gpu_price" field.
func (mpu *ModelPriceUpdate) AddOutputGpuPrice(i int) *ModelPriceUpdate {
	mpu.mutation.AddOutputGpuPrice(i)
	return mpu
}

// SetInputModelPrice sets the "input_model_price" field.
func (mpu *ModelPriceUpdate) SetInputModelPrice(i int) *ModelPriceUpdate {
	mpu.mutation.ResetInputModelPrice()
	mpu.mutation.SetInputModelPrice(i)
	return mpu
}

// SetNillableInputModelPrice sets the "input_model_price" field if the given value is not nil.
func (mpu *ModelPriceUpdate) SetNillableInputModelPrice(i *int) *ModelPriceUpdate {
	if i != nil {
		mpu.SetInputModelPrice(*i)
	}
	return mpu
}

// AddInputModelPrice adds i to the "input_model_price" field.
func (mpu *ModelPriceUpdate) AddInputModelPrice(i int) *ModelPriceUpdate {
	mpu.mutation.AddInputModelPrice(i)
	return mpu
}

// SetOutputModelPrice sets the "output_model_price" field.
func (mpu *ModelPriceUpdate) SetOutputModelPrice(i int) *ModelPriceUpdate {
	mpu.mutation.ResetOutputModelPrice()
	mpu.mutation.SetOutputModelPrice(i)
	return mpu
}

// SetNillableOutputModelPrice sets the "output_model_price" field if the given value is not nil.
func (mpu *ModelPriceUpdate) SetNillableOutputModelPrice(i *int) *ModelPriceUpdate {
	if i != nil {
		mpu.SetOutputModelPrice(*i)
	}
	return mpu
}

// AddOutputModelPrice adds i to the "output_model_price" field.
func (mpu *ModelPriceUpdate) AddOutputModelPrice(i int) *ModelPriceUpdate {
	mpu.mutation.AddOutputModelPrice(i)
	return mpu
}

// SetModel sets the "model" edge to the Model entity.
func (mpu *ModelPriceUpdate) SetModel(m *Model) *ModelPriceUpdate {
	return mpu.SetModelID(m.ID)
}

// Mutation returns the ModelPriceMutation object of the builder.
func (mpu *ModelPriceUpdate) Mutation() *ModelPriceMutation {
	return mpu.mutation
}

// ClearModel clears the "model" edge to the Model entity.
func (mpu *ModelPriceUpdate) ClearModel() *ModelPriceUpdate {
	mpu.mutation.ClearModel()
	return mpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mpu *ModelPriceUpdate) Save(ctx context.Context) (int, error) {
	mpu.defaults()
	return withHooks(ctx, mpu.sqlSave, mpu.mutation, mpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mpu *ModelPriceUpdate) SaveX(ctx context.Context) int {
	affected, err := mpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mpu *ModelPriceUpdate) Exec(ctx context.Context) error {
	_, err := mpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpu *ModelPriceUpdate) ExecX(ctx context.Context) {
	if err := mpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mpu *ModelPriceUpdate) defaults() {
	if _, ok := mpu.mutation.UpdatedAt(); !ok {
		v := modelprice.UpdateDefaultUpdatedAt()
		mpu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mpu *ModelPriceUpdate) check() error {
	if v, ok := mpu.mutation.InvokeType(); ok {
		if err := modelprice.InvokeTypeValidator(v); err != nil {
			return &ValidationError{Name: "invoke_type", err: fmt.Errorf(`cep_ent: validator failed for field "ModelPrice.invoke_type": %w`, err)}
		}
	}
	if v, ok := mpu.mutation.GpuVersion(); ok {
		if err := modelprice.GpuVersionValidator(v); err != nil {
			return &ValidationError{Name: "gpu_version", err: fmt.Errorf(`cep_ent: validator failed for field "ModelPrice.gpu_version": %w`, err)}
		}
	}
	if _, ok := mpu.mutation.ModelID(); mpu.mutation.ModelCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ModelPrice.model"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mpu *ModelPriceUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ModelPriceUpdate {
	mpu.modifiers = append(mpu.modifiers, modifiers...)
	return mpu
}

func (mpu *ModelPriceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(modelprice.Table, modelprice.Columns, sqlgraph.NewFieldSpec(modelprice.FieldID, field.TypeInt64))
	if ps := mpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mpu.mutation.CreatedBy(); ok {
		_spec.SetField(modelprice.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mpu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(modelprice.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mpu.mutation.UpdatedBy(); ok {
		_spec.SetField(modelprice.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mpu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(modelprice.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mpu.mutation.UpdatedAt(); ok {
		_spec.SetField(modelprice.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mpu.mutation.DeletedAt(); ok {
		_spec.SetField(modelprice.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := mpu.mutation.InvokeType(); ok {
		_spec.SetField(modelprice.FieldInvokeType, field.TypeEnum, value)
	}
	if value, ok := mpu.mutation.GpuVersion(); ok {
		_spec.SetField(modelprice.FieldGpuVersion, field.TypeEnum, value)
	}
	if value, ok := mpu.mutation.InputGpuPrice(); ok {
		_spec.SetField(modelprice.FieldInputGpuPrice, field.TypeInt, value)
	}
	if value, ok := mpu.mutation.AddedInputGpuPrice(); ok {
		_spec.AddField(modelprice.FieldInputGpuPrice, field.TypeInt, value)
	}
	if value, ok := mpu.mutation.OutputGpuPrice(); ok {
		_spec.SetField(modelprice.FieldOutputGpuPrice, field.TypeInt, value)
	}
	if value, ok := mpu.mutation.AddedOutputGpuPrice(); ok {
		_spec.AddField(modelprice.FieldOutputGpuPrice, field.TypeInt, value)
	}
	if value, ok := mpu.mutation.InputModelPrice(); ok {
		_spec.SetField(modelprice.FieldInputModelPrice, field.TypeInt, value)
	}
	if value, ok := mpu.mutation.AddedInputModelPrice(); ok {
		_spec.AddField(modelprice.FieldInputModelPrice, field.TypeInt, value)
	}
	if value, ok := mpu.mutation.OutputModelPrice(); ok {
		_spec.SetField(modelprice.FieldOutputModelPrice, field.TypeInt, value)
	}
	if value, ok := mpu.mutation.AddedOutputModelPrice(); ok {
		_spec.AddField(modelprice.FieldOutputModelPrice, field.TypeInt, value)
	}
	if mpu.mutation.ModelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   modelprice.ModelTable,
			Columns: []string{modelprice.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpu.mutation.ModelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   modelprice.ModelTable,
			Columns: []string{modelprice.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(mpu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{modelprice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mpu.mutation.done = true
	return n, nil
}

// ModelPriceUpdateOne is the builder for updating a single ModelPrice entity.
type ModelPriceUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ModelPriceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (mpuo *ModelPriceUpdateOne) SetCreatedBy(i int64) *ModelPriceUpdateOne {
	mpuo.mutation.ResetCreatedBy()
	mpuo.mutation.SetCreatedBy(i)
	return mpuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mpuo *ModelPriceUpdateOne) SetNillableCreatedBy(i *int64) *ModelPriceUpdateOne {
	if i != nil {
		mpuo.SetCreatedBy(*i)
	}
	return mpuo
}

// AddCreatedBy adds i to the "created_by" field.
func (mpuo *ModelPriceUpdateOne) AddCreatedBy(i int64) *ModelPriceUpdateOne {
	mpuo.mutation.AddCreatedBy(i)
	return mpuo
}

// SetUpdatedBy sets the "updated_by" field.
func (mpuo *ModelPriceUpdateOne) SetUpdatedBy(i int64) *ModelPriceUpdateOne {
	mpuo.mutation.ResetUpdatedBy()
	mpuo.mutation.SetUpdatedBy(i)
	return mpuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mpuo *ModelPriceUpdateOne) SetNillableUpdatedBy(i *int64) *ModelPriceUpdateOne {
	if i != nil {
		mpuo.SetUpdatedBy(*i)
	}
	return mpuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (mpuo *ModelPriceUpdateOne) AddUpdatedBy(i int64) *ModelPriceUpdateOne {
	mpuo.mutation.AddUpdatedBy(i)
	return mpuo
}

// SetUpdatedAt sets the "updated_at" field.
func (mpuo *ModelPriceUpdateOne) SetUpdatedAt(t time.Time) *ModelPriceUpdateOne {
	mpuo.mutation.SetUpdatedAt(t)
	return mpuo
}

// SetDeletedAt sets the "deleted_at" field.
func (mpuo *ModelPriceUpdateOne) SetDeletedAt(t time.Time) *ModelPriceUpdateOne {
	mpuo.mutation.SetDeletedAt(t)
	return mpuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mpuo *ModelPriceUpdateOne) SetNillableDeletedAt(t *time.Time) *ModelPriceUpdateOne {
	if t != nil {
		mpuo.SetDeletedAt(*t)
	}
	return mpuo
}

// SetModelID sets the "model_id" field.
func (mpuo *ModelPriceUpdateOne) SetModelID(i int64) *ModelPriceUpdateOne {
	mpuo.mutation.SetModelID(i)
	return mpuo
}

// SetInvokeType sets the "invoke_type" field.
func (mpuo *ModelPriceUpdateOne) SetInvokeType(e enums.Model) *ModelPriceUpdateOne {
	mpuo.mutation.SetInvokeType(e)
	return mpuo
}

// SetNillableInvokeType sets the "invoke_type" field if the given value is not nil.
func (mpuo *ModelPriceUpdateOne) SetNillableInvokeType(e *enums.Model) *ModelPriceUpdateOne {
	if e != nil {
		mpuo.SetInvokeType(*e)
	}
	return mpuo
}

// SetGpuVersion sets the "gpu_version" field.
func (mpuo *ModelPriceUpdateOne) SetGpuVersion(ev enums.GpuVersion) *ModelPriceUpdateOne {
	mpuo.mutation.SetGpuVersion(ev)
	return mpuo
}

// SetNillableGpuVersion sets the "gpu_version" field if the given value is not nil.
func (mpuo *ModelPriceUpdateOne) SetNillableGpuVersion(ev *enums.GpuVersion) *ModelPriceUpdateOne {
	if ev != nil {
		mpuo.SetGpuVersion(*ev)
	}
	return mpuo
}

// SetInputGpuPrice sets the "input_gpu_price" field.
func (mpuo *ModelPriceUpdateOne) SetInputGpuPrice(i int) *ModelPriceUpdateOne {
	mpuo.mutation.ResetInputGpuPrice()
	mpuo.mutation.SetInputGpuPrice(i)
	return mpuo
}

// SetNillableInputGpuPrice sets the "input_gpu_price" field if the given value is not nil.
func (mpuo *ModelPriceUpdateOne) SetNillableInputGpuPrice(i *int) *ModelPriceUpdateOne {
	if i != nil {
		mpuo.SetInputGpuPrice(*i)
	}
	return mpuo
}

// AddInputGpuPrice adds i to the "input_gpu_price" field.
func (mpuo *ModelPriceUpdateOne) AddInputGpuPrice(i int) *ModelPriceUpdateOne {
	mpuo.mutation.AddInputGpuPrice(i)
	return mpuo
}

// SetOutputGpuPrice sets the "output_gpu_price" field.
func (mpuo *ModelPriceUpdateOne) SetOutputGpuPrice(i int) *ModelPriceUpdateOne {
	mpuo.mutation.ResetOutputGpuPrice()
	mpuo.mutation.SetOutputGpuPrice(i)
	return mpuo
}

// SetNillableOutputGpuPrice sets the "output_gpu_price" field if the given value is not nil.
func (mpuo *ModelPriceUpdateOne) SetNillableOutputGpuPrice(i *int) *ModelPriceUpdateOne {
	if i != nil {
		mpuo.SetOutputGpuPrice(*i)
	}
	return mpuo
}

// AddOutputGpuPrice adds i to the "output_gpu_price" field.
func (mpuo *ModelPriceUpdateOne) AddOutputGpuPrice(i int) *ModelPriceUpdateOne {
	mpuo.mutation.AddOutputGpuPrice(i)
	return mpuo
}

// SetInputModelPrice sets the "input_model_price" field.
func (mpuo *ModelPriceUpdateOne) SetInputModelPrice(i int) *ModelPriceUpdateOne {
	mpuo.mutation.ResetInputModelPrice()
	mpuo.mutation.SetInputModelPrice(i)
	return mpuo
}

// SetNillableInputModelPrice sets the "input_model_price" field if the given value is not nil.
func (mpuo *ModelPriceUpdateOne) SetNillableInputModelPrice(i *int) *ModelPriceUpdateOne {
	if i != nil {
		mpuo.SetInputModelPrice(*i)
	}
	return mpuo
}

// AddInputModelPrice adds i to the "input_model_price" field.
func (mpuo *ModelPriceUpdateOne) AddInputModelPrice(i int) *ModelPriceUpdateOne {
	mpuo.mutation.AddInputModelPrice(i)
	return mpuo
}

// SetOutputModelPrice sets the "output_model_price" field.
func (mpuo *ModelPriceUpdateOne) SetOutputModelPrice(i int) *ModelPriceUpdateOne {
	mpuo.mutation.ResetOutputModelPrice()
	mpuo.mutation.SetOutputModelPrice(i)
	return mpuo
}

// SetNillableOutputModelPrice sets the "output_model_price" field if the given value is not nil.
func (mpuo *ModelPriceUpdateOne) SetNillableOutputModelPrice(i *int) *ModelPriceUpdateOne {
	if i != nil {
		mpuo.SetOutputModelPrice(*i)
	}
	return mpuo
}

// AddOutputModelPrice adds i to the "output_model_price" field.
func (mpuo *ModelPriceUpdateOne) AddOutputModelPrice(i int) *ModelPriceUpdateOne {
	mpuo.mutation.AddOutputModelPrice(i)
	return mpuo
}

// SetModel sets the "model" edge to the Model entity.
func (mpuo *ModelPriceUpdateOne) SetModel(m *Model) *ModelPriceUpdateOne {
	return mpuo.SetModelID(m.ID)
}

// Mutation returns the ModelPriceMutation object of the builder.
func (mpuo *ModelPriceUpdateOne) Mutation() *ModelPriceMutation {
	return mpuo.mutation
}

// ClearModel clears the "model" edge to the Model entity.
func (mpuo *ModelPriceUpdateOne) ClearModel() *ModelPriceUpdateOne {
	mpuo.mutation.ClearModel()
	return mpuo
}

// Where appends a list predicates to the ModelPriceUpdate builder.
func (mpuo *ModelPriceUpdateOne) Where(ps ...predicate.ModelPrice) *ModelPriceUpdateOne {
	mpuo.mutation.Where(ps...)
	return mpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mpuo *ModelPriceUpdateOne) Select(field string, fields ...string) *ModelPriceUpdateOne {
	mpuo.fields = append([]string{field}, fields...)
	return mpuo
}

// Save executes the query and returns the updated ModelPrice entity.
func (mpuo *ModelPriceUpdateOne) Save(ctx context.Context) (*ModelPrice, error) {
	mpuo.defaults()
	return withHooks(ctx, mpuo.sqlSave, mpuo.mutation, mpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mpuo *ModelPriceUpdateOne) SaveX(ctx context.Context) *ModelPrice {
	node, err := mpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mpuo *ModelPriceUpdateOne) Exec(ctx context.Context) error {
	_, err := mpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpuo *ModelPriceUpdateOne) ExecX(ctx context.Context) {
	if err := mpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mpuo *ModelPriceUpdateOne) defaults() {
	if _, ok := mpuo.mutation.UpdatedAt(); !ok {
		v := modelprice.UpdateDefaultUpdatedAt()
		mpuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mpuo *ModelPriceUpdateOne) check() error {
	if v, ok := mpuo.mutation.InvokeType(); ok {
		if err := modelprice.InvokeTypeValidator(v); err != nil {
			return &ValidationError{Name: "invoke_type", err: fmt.Errorf(`cep_ent: validator failed for field "ModelPrice.invoke_type": %w`, err)}
		}
	}
	if v, ok := mpuo.mutation.GpuVersion(); ok {
		if err := modelprice.GpuVersionValidator(v); err != nil {
			return &ValidationError{Name: "gpu_version", err: fmt.Errorf(`cep_ent: validator failed for field "ModelPrice.gpu_version": %w`, err)}
		}
	}
	if _, ok := mpuo.mutation.ModelID(); mpuo.mutation.ModelCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "ModelPrice.model"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mpuo *ModelPriceUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ModelPriceUpdateOne {
	mpuo.modifiers = append(mpuo.modifiers, modifiers...)
	return mpuo
}

func (mpuo *ModelPriceUpdateOne) sqlSave(ctx context.Context) (_node *ModelPrice, err error) {
	if err := mpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(modelprice.Table, modelprice.Columns, sqlgraph.NewFieldSpec(modelprice.FieldID, field.TypeInt64))
	id, ok := mpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "ModelPrice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, modelprice.FieldID)
		for _, f := range fields {
			if !modelprice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != modelprice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mpuo.mutation.CreatedBy(); ok {
		_spec.SetField(modelprice.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mpuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(modelprice.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := mpuo.mutation.UpdatedBy(); ok {
		_spec.SetField(modelprice.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mpuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(modelprice.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := mpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(modelprice.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mpuo.mutation.DeletedAt(); ok {
		_spec.SetField(modelprice.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := mpuo.mutation.InvokeType(); ok {
		_spec.SetField(modelprice.FieldInvokeType, field.TypeEnum, value)
	}
	if value, ok := mpuo.mutation.GpuVersion(); ok {
		_spec.SetField(modelprice.FieldGpuVersion, field.TypeEnum, value)
	}
	if value, ok := mpuo.mutation.InputGpuPrice(); ok {
		_spec.SetField(modelprice.FieldInputGpuPrice, field.TypeInt, value)
	}
	if value, ok := mpuo.mutation.AddedInputGpuPrice(); ok {
		_spec.AddField(modelprice.FieldInputGpuPrice, field.TypeInt, value)
	}
	if value, ok := mpuo.mutation.OutputGpuPrice(); ok {
		_spec.SetField(modelprice.FieldOutputGpuPrice, field.TypeInt, value)
	}
	if value, ok := mpuo.mutation.AddedOutputGpuPrice(); ok {
		_spec.AddField(modelprice.FieldOutputGpuPrice, field.TypeInt, value)
	}
	if value, ok := mpuo.mutation.InputModelPrice(); ok {
		_spec.SetField(modelprice.FieldInputModelPrice, field.TypeInt, value)
	}
	if value, ok := mpuo.mutation.AddedInputModelPrice(); ok {
		_spec.AddField(modelprice.FieldInputModelPrice, field.TypeInt, value)
	}
	if value, ok := mpuo.mutation.OutputModelPrice(); ok {
		_spec.SetField(modelprice.FieldOutputModelPrice, field.TypeInt, value)
	}
	if value, ok := mpuo.mutation.AddedOutputModelPrice(); ok {
		_spec.AddField(modelprice.FieldOutputModelPrice, field.TypeInt, value)
	}
	if mpuo.mutation.ModelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   modelprice.ModelTable,
			Columns: []string{modelprice.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpuo.mutation.ModelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   modelprice.ModelTable,
			Columns: []string{modelprice.ModelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(model.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(mpuo.modifiers...)
	_node = &ModelPrice{config: mpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{modelprice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mpuo.mutation.done = true
	return _node, nil
}
