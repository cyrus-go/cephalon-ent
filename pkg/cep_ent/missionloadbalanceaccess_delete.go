// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionloadbalanceaccess"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// MissionLoadBalanceAccessDelete is the builder for deleting a MissionLoadBalanceAccess entity.
type MissionLoadBalanceAccessDelete struct {
	config
	hooks    []Hook
	mutation *MissionLoadBalanceAccessMutation
}

// Where appends a list predicates to the MissionLoadBalanceAccessDelete builder.
func (mlbad *MissionLoadBalanceAccessDelete) Where(ps ...predicate.MissionLoadBalanceAccess) *MissionLoadBalanceAccessDelete {
	mlbad.mutation.Where(ps...)
	return mlbad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mlbad *MissionLoadBalanceAccessDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mlbad.sqlExec, mlbad.mutation, mlbad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mlbad *MissionLoadBalanceAccessDelete) ExecX(ctx context.Context) int {
	n, err := mlbad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mlbad *MissionLoadBalanceAccessDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(missionloadbalanceaccess.Table, sqlgraph.NewFieldSpec(missionloadbalanceaccess.FieldID, field.TypeInt64))
	if ps := mlbad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mlbad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mlbad.mutation.done = true
	return affected, err
}

// MissionLoadBalanceAccessDeleteOne is the builder for deleting a single MissionLoadBalanceAccess entity.
type MissionLoadBalanceAccessDeleteOne struct {
	mlbad *MissionLoadBalanceAccessDelete
}

// Where appends a list predicates to the MissionLoadBalanceAccessDelete builder.
func (mlbado *MissionLoadBalanceAccessDeleteOne) Where(ps ...predicate.MissionLoadBalanceAccess) *MissionLoadBalanceAccessDeleteOne {
	mlbado.mlbad.mutation.Where(ps...)
	return mlbado
}

// Exec executes the deletion query.
func (mlbado *MissionLoadBalanceAccessDeleteOne) Exec(ctx context.Context) error {
	n, err := mlbado.mlbad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{missionloadbalanceaccess.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mlbado *MissionLoadBalanceAccessDeleteOne) ExecX(ctx context.Context) {
	if err := mlbado.Exec(ctx); err != nil {
		panic(err)
	}
}
