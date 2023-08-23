// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/costbill"
	"cephalon-ent/pkg/cep_ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CostBillDelete is the builder for deleting a CostBill entity.
type CostBillDelete struct {
	config
	hooks    []Hook
	mutation *CostBillMutation
}

// Where appends a list predicates to the CostBillDelete builder.
func (cbd *CostBillDelete) Where(ps ...predicate.CostBill) *CostBillDelete {
	cbd.mutation.Where(ps...)
	return cbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cbd *CostBillDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cbd.sqlExec, cbd.mutation, cbd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cbd *CostBillDelete) ExecX(ctx context.Context) int {
	n, err := cbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cbd *CostBillDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(costbill.Table, sqlgraph.NewFieldSpec(costbill.FieldID, field.TypeInt64))
	if ps := cbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cbd.mutation.done = true
	return affected, err
}

// CostBillDeleteOne is the builder for deleting a single CostBill entity.
type CostBillDeleteOne struct {
	cbd *CostBillDelete
}

// Where appends a list predicates to the CostBillDelete builder.
func (cbdo *CostBillDeleteOne) Where(ps ...predicate.CostBill) *CostBillDeleteOne {
	cbdo.cbd.mutation.Where(ps...)
	return cbdo
}

// Exec executes the deletion query.
func (cbdo *CostBillDeleteOne) Exec(ctx context.Context) error {
	n, err := cbdo.cbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{costbill.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cbdo *CostBillDeleteOne) ExecX(ctx context.Context) {
	if err := cbdo.Exec(ctx); err != nil {
		panic(err)
	}
}
