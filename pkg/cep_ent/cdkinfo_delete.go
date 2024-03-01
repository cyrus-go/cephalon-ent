// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/cdkinfo"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// CDKInfoDelete is the builder for deleting a CDKInfo entity.
type CDKInfoDelete struct {
	config
	hooks    []Hook
	mutation *CDKInfoMutation
}

// Where appends a list predicates to the CDKInfoDelete builder.
func (cid *CDKInfoDelete) Where(ps ...predicate.CDKInfo) *CDKInfoDelete {
	cid.mutation.Where(ps...)
	return cid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cid *CDKInfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cid.sqlExec, cid.mutation, cid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cid *CDKInfoDelete) ExecX(ctx context.Context) int {
	n, err := cid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cid *CDKInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(cdkinfo.Table, sqlgraph.NewFieldSpec(cdkinfo.FieldID, field.TypeInt64))
	if ps := cid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cid.mutation.done = true
	return affected, err
}

// CDKInfoDeleteOne is the builder for deleting a single CDKInfo entity.
type CDKInfoDeleteOne struct {
	cid *CDKInfoDelete
}

// Where appends a list predicates to the CDKInfoDelete builder.
func (cido *CDKInfoDeleteOne) Where(ps ...predicate.CDKInfo) *CDKInfoDeleteOne {
	cido.cid.mutation.Where(ps...)
	return cido
}

// Exec executes the deletion query.
func (cido *CDKInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := cido.cid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cdkinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cido *CDKInfoDeleteOne) ExecX(ctx context.Context) {
	if err := cido.Exec(ctx); err != nil {
		panic(err)
	}
}