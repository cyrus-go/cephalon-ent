// Code generated by ent, DO NOT EDIT.

package missionkind

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldEQ(FieldDeletedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.MissionKind {
	return predicate.MissionKind(sql.FieldLTE(FieldDeletedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v enums.MissionType) predicate.MissionKind {
	vc := v
	return predicate.MissionKind(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v enums.MissionType) predicate.MissionKind {
	vc := v
	return predicate.MissionKind(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...enums.MissionType) predicate.MissionKind {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionKind(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...enums.MissionType) predicate.MissionKind {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionKind(sql.FieldNotIn(FieldType, v...))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v enums.MissionCategory) predicate.MissionKind {
	vc := v
	return predicate.MissionKind(sql.FieldEQ(FieldCategory, vc))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v enums.MissionCategory) predicate.MissionKind {
	vc := v
	return predicate.MissionKind(sql.FieldNEQ(FieldCategory, vc))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...enums.MissionCategory) predicate.MissionKind {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionKind(sql.FieldIn(FieldCategory, v...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...enums.MissionCategory) predicate.MissionKind {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionKind(sql.FieldNotIn(FieldCategory, v...))
}

// BillingTypeEQ applies the EQ predicate on the "billing_type" field.
func BillingTypeEQ(v enums.MissionBillingType) predicate.MissionKind {
	vc := v
	return predicate.MissionKind(sql.FieldEQ(FieldBillingType, vc))
}

// BillingTypeNEQ applies the NEQ predicate on the "billing_type" field.
func BillingTypeNEQ(v enums.MissionBillingType) predicate.MissionKind {
	vc := v
	return predicate.MissionKind(sql.FieldNEQ(FieldBillingType, vc))
}

// BillingTypeIn applies the In predicate on the "billing_type" field.
func BillingTypeIn(vs ...enums.MissionBillingType) predicate.MissionKind {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionKind(sql.FieldIn(FieldBillingType, v...))
}

// BillingTypeNotIn applies the NotIn predicate on the "billing_type" field.
func BillingTypeNotIn(vs ...enums.MissionBillingType) predicate.MissionKind {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionKind(sql.FieldNotIn(FieldBillingType, v...))
}

// HasMissions applies the HasEdge predicate on the "missions" edge.
func HasMissions() predicate.MissionKind {
	return predicate.MissionKind(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MissionsTable, MissionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMissionsWith applies the HasEdge predicate on the "missions" edge with a given conditions (other predicates).
func HasMissionsWith(preds ...predicate.Mission) predicate.MissionKind {
	return predicate.MissionKind(func(s *sql.Selector) {
		step := newMissionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MissionKind) predicate.MissionKind {
	return predicate.MissionKind(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MissionKind) predicate.MissionKind {
	return predicate.MissionKind(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MissionKind) predicate.MissionKind {
	return predicate.MissionKind(sql.NotPredicates(p))
}
