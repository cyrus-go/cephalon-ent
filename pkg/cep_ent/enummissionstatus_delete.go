// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"cephalon-ent/pkg/cep_ent/enummissionstatus"
	"cephalon-ent/pkg/cep_ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EnumMissionStatusDelete is the builder for deleting a EnumMissionStatus entity.
type EnumMissionStatusDelete struct {
	config
	hooks    []Hook
	mutation *EnumMissionStatusMutation
}

// Where appends a list predicates to the EnumMissionStatusDelete builder.
func (emsd *EnumMissionStatusDelete) Where(ps ...predicate.EnumMissionStatus) *EnumMissionStatusDelete {
	emsd.mutation.Where(ps...)
	return emsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (emsd *EnumMissionStatusDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, emsd.sqlExec, emsd.mutation, emsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (emsd *EnumMissionStatusDelete) ExecX(ctx context.Context) int {
	n, err := emsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (emsd *EnumMissionStatusDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(enummissionstatus.Table, sqlgraph.NewFieldSpec(enummissionstatus.FieldID, field.TypeInt64))
	if ps := emsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, emsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	emsd.mutation.done = true
	return affected, err
}

// EnumMissionStatusDeleteOne is the builder for deleting a single EnumMissionStatus entity.
type EnumMissionStatusDeleteOne struct {
	emsd *EnumMissionStatusDelete
}

// Where appends a list predicates to the EnumMissionStatusDelete builder.
func (emsdo *EnumMissionStatusDeleteOne) Where(ps ...predicate.EnumMissionStatus) *EnumMissionStatusDeleteOne {
	emsdo.emsd.mutation.Where(ps...)
	return emsdo
}

// Exec executes the deletion query.
func (emsdo *EnumMissionStatusDeleteOne) Exec(ctx context.Context) error {
	n, err := emsdo.emsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{enummissionstatus.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (emsdo *EnumMissionStatusDeleteOne) ExecX(ctx context.Context) {
	if err := emsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
