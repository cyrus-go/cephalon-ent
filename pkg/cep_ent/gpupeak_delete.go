// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/gpupeak"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// GpuPeakDelete is the builder for deleting a GpuPeak entity.
type GpuPeakDelete struct {
	config
	hooks    []Hook
	mutation *GpuPeakMutation
}

// Where appends a list predicates to the GpuPeakDelete builder.
func (gpd *GpuPeakDelete) Where(ps ...predicate.GpuPeak) *GpuPeakDelete {
	gpd.mutation.Where(ps...)
	return gpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gpd *GpuPeakDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, gpd.sqlExec, gpd.mutation, gpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gpd *GpuPeakDelete) ExecX(ctx context.Context) int {
	n, err := gpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gpd *GpuPeakDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(gpupeak.Table, sqlgraph.NewFieldSpec(gpupeak.FieldID, field.TypeInt64))
	if ps := gpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gpd.mutation.done = true
	return affected, err
}

// GpuPeakDeleteOne is the builder for deleting a single GpuPeak entity.
type GpuPeakDeleteOne struct {
	gpd *GpuPeakDelete
}

// Where appends a list predicates to the GpuPeakDelete builder.
func (gpdo *GpuPeakDeleteOne) Where(ps ...predicate.GpuPeak) *GpuPeakDeleteOne {
	gpdo.gpd.mutation.Where(ps...)
	return gpdo
}

// Exec executes the deletion query.
func (gpdo *GpuPeakDeleteOne) Exec(ctx context.Context) error {
	n, err := gpdo.gpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{gpupeak.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gpdo *GpuPeakDeleteOne) ExecX(ctx context.Context) {
	if err := gpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
