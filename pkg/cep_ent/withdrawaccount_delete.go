// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/withdrawaccount"
)

// WithdrawAccountDelete is the builder for deleting a WithdrawAccount entity.
type WithdrawAccountDelete struct {
	config
	hooks    []Hook
	mutation *WithdrawAccountMutation
}

// Where appends a list predicates to the WithdrawAccountDelete builder.
func (wad *WithdrawAccountDelete) Where(ps ...predicate.WithdrawAccount) *WithdrawAccountDelete {
	wad.mutation.Where(ps...)
	return wad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wad *WithdrawAccountDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wad.sqlExec, wad.mutation, wad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wad *WithdrawAccountDelete) ExecX(ctx context.Context) int {
	n, err := wad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wad *WithdrawAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(withdrawaccount.Table, sqlgraph.NewFieldSpec(withdrawaccount.FieldID, field.TypeInt64))
	if ps := wad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wad.mutation.done = true
	return affected, err
}

// WithdrawAccountDeleteOne is the builder for deleting a single WithdrawAccount entity.
type WithdrawAccountDeleteOne struct {
	wad *WithdrawAccountDelete
}

// Where appends a list predicates to the WithdrawAccountDelete builder.
func (wado *WithdrawAccountDeleteOne) Where(ps ...predicate.WithdrawAccount) *WithdrawAccountDeleteOne {
	wado.wad.mutation.Where(ps...)
	return wado
}

// Exec executes the deletion query.
func (wado *WithdrawAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := wado.wad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{withdrawaccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wado *WithdrawAccountDeleteOne) ExecX(ctx context.Context) {
	if err := wado.Exec(ctx); err != nil {
		panic(err)
	}
}
