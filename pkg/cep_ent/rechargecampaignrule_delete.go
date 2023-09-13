// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/rechargecampaignrule"
)

// RechargeCampaignRuleDelete is the builder for deleting a RechargeCampaignRule entity.
type RechargeCampaignRuleDelete struct {
	config
	hooks    []Hook
	mutation *RechargeCampaignRuleMutation
}

// Where appends a list predicates to the RechargeCampaignRuleDelete builder.
func (rcrd *RechargeCampaignRuleDelete) Where(ps ...predicate.RechargeCampaignRule) *RechargeCampaignRuleDelete {
	rcrd.mutation.Where(ps...)
	return rcrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rcrd *RechargeCampaignRuleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rcrd.sqlExec, rcrd.mutation, rcrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rcrd *RechargeCampaignRuleDelete) ExecX(ctx context.Context) int {
	n, err := rcrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rcrd *RechargeCampaignRuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(rechargecampaignrule.Table, sqlgraph.NewFieldSpec(rechargecampaignrule.FieldID, field.TypeInt64))
	if ps := rcrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rcrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rcrd.mutation.done = true
	return affected, err
}

// RechargeCampaignRuleDeleteOne is the builder for deleting a single RechargeCampaignRule entity.
type RechargeCampaignRuleDeleteOne struct {
	rcrd *RechargeCampaignRuleDelete
}

// Where appends a list predicates to the RechargeCampaignRuleDelete builder.
func (rcrdo *RechargeCampaignRuleDeleteOne) Where(ps ...predicate.RechargeCampaignRule) *RechargeCampaignRuleDeleteOne {
	rcrdo.rcrd.mutation.Where(ps...)
	return rcrdo
}

// Exec executes the deletion query.
func (rcrdo *RechargeCampaignRuleDeleteOne) Exec(ctx context.Context) error {
	n, err := rcrdo.rcrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rechargecampaignrule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rcrdo *RechargeCampaignRuleDeleteOne) ExecX(ctx context.Context) {
	if err := rcrdo.Exec(ctx); err != nil {
		panic(err)
	}
}