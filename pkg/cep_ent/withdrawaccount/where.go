// Code generated by ent, DO NOT EDIT.

package withdrawaccount

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldDeletedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldUserID, v))
}

// BusinessName applies equality check predicate on the "business_name" field. It's identical to BusinessNameEQ.
func BusinessName(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldBusinessName, v))
}

// BusinessType applies equality check predicate on the "business_type" field. It's identical to BusinessTypeEQ.
func BusinessType(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldBusinessType, v))
}

// BusinessID applies equality check predicate on the "business_id" field. It's identical to BusinessIDEQ.
func BusinessID(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldBusinessID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLTE(FieldDeletedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNotIn(FieldUserID, vs...))
}

// BusinessNameEQ applies the EQ predicate on the "business_name" field.
func BusinessNameEQ(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldBusinessName, v))
}

// BusinessNameNEQ applies the NEQ predicate on the "business_name" field.
func BusinessNameNEQ(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNEQ(FieldBusinessName, v))
}

// BusinessNameIn applies the In predicate on the "business_name" field.
func BusinessNameIn(vs ...string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldIn(FieldBusinessName, vs...))
}

// BusinessNameNotIn applies the NotIn predicate on the "business_name" field.
func BusinessNameNotIn(vs ...string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNotIn(FieldBusinessName, vs...))
}

// BusinessNameGT applies the GT predicate on the "business_name" field.
func BusinessNameGT(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGT(FieldBusinessName, v))
}

// BusinessNameGTE applies the GTE predicate on the "business_name" field.
func BusinessNameGTE(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGTE(FieldBusinessName, v))
}

// BusinessNameLT applies the LT predicate on the "business_name" field.
func BusinessNameLT(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLT(FieldBusinessName, v))
}

// BusinessNameLTE applies the LTE predicate on the "business_name" field.
func BusinessNameLTE(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLTE(FieldBusinessName, v))
}

// BusinessNameContains applies the Contains predicate on the "business_name" field.
func BusinessNameContains(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldContains(FieldBusinessName, v))
}

// BusinessNameHasPrefix applies the HasPrefix predicate on the "business_name" field.
func BusinessNameHasPrefix(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldHasPrefix(FieldBusinessName, v))
}

// BusinessNameHasSuffix applies the HasSuffix predicate on the "business_name" field.
func BusinessNameHasSuffix(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldHasSuffix(FieldBusinessName, v))
}

// BusinessNameEqualFold applies the EqualFold predicate on the "business_name" field.
func BusinessNameEqualFold(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEqualFold(FieldBusinessName, v))
}

// BusinessNameContainsFold applies the ContainsFold predicate on the "business_name" field.
func BusinessNameContainsFold(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldContainsFold(FieldBusinessName, v))
}

// BusinessTypeEQ applies the EQ predicate on the "business_type" field.
func BusinessTypeEQ(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldBusinessType, v))
}

// BusinessTypeNEQ applies the NEQ predicate on the "business_type" field.
func BusinessTypeNEQ(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNEQ(FieldBusinessType, v))
}

// BusinessTypeIn applies the In predicate on the "business_type" field.
func BusinessTypeIn(vs ...string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldIn(FieldBusinessType, vs...))
}

// BusinessTypeNotIn applies the NotIn predicate on the "business_type" field.
func BusinessTypeNotIn(vs ...string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNotIn(FieldBusinessType, vs...))
}

// BusinessTypeGT applies the GT predicate on the "business_type" field.
func BusinessTypeGT(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGT(FieldBusinessType, v))
}

// BusinessTypeGTE applies the GTE predicate on the "business_type" field.
func BusinessTypeGTE(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGTE(FieldBusinessType, v))
}

// BusinessTypeLT applies the LT predicate on the "business_type" field.
func BusinessTypeLT(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLT(FieldBusinessType, v))
}

// BusinessTypeLTE applies the LTE predicate on the "business_type" field.
func BusinessTypeLTE(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLTE(FieldBusinessType, v))
}

// BusinessTypeContains applies the Contains predicate on the "business_type" field.
func BusinessTypeContains(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldContains(FieldBusinessType, v))
}

// BusinessTypeHasPrefix applies the HasPrefix predicate on the "business_type" field.
func BusinessTypeHasPrefix(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldHasPrefix(FieldBusinessType, v))
}

// BusinessTypeHasSuffix applies the HasSuffix predicate on the "business_type" field.
func BusinessTypeHasSuffix(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldHasSuffix(FieldBusinessType, v))
}

// BusinessTypeEqualFold applies the EqualFold predicate on the "business_type" field.
func BusinessTypeEqualFold(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEqualFold(FieldBusinessType, v))
}

// BusinessTypeContainsFold applies the ContainsFold predicate on the "business_type" field.
func BusinessTypeContainsFold(v string) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldContainsFold(FieldBusinessType, v))
}

// BusinessIDEQ applies the EQ predicate on the "business_id" field.
func BusinessIDEQ(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldEQ(FieldBusinessID, v))
}

// BusinessIDNEQ applies the NEQ predicate on the "business_id" field.
func BusinessIDNEQ(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNEQ(FieldBusinessID, v))
}

// BusinessIDIn applies the In predicate on the "business_id" field.
func BusinessIDIn(vs ...int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldIn(FieldBusinessID, vs...))
}

// BusinessIDNotIn applies the NotIn predicate on the "business_id" field.
func BusinessIDNotIn(vs ...int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldNotIn(FieldBusinessID, vs...))
}

// BusinessIDGT applies the GT predicate on the "business_id" field.
func BusinessIDGT(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGT(FieldBusinessID, v))
}

// BusinessIDGTE applies the GTE predicate on the "business_id" field.
func BusinessIDGTE(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldGTE(FieldBusinessID, v))
}

// BusinessIDLT applies the LT predicate on the "business_id" field.
func BusinessIDLT(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLT(FieldBusinessID, v))
}

// BusinessIDLTE applies the LTE predicate on the "business_id" field.
func BusinessIDLTE(v int64) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.FieldLTE(FieldBusinessID, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.WithdrawAccount {
	return predicate.WithdrawAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WithdrawAccount) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WithdrawAccount) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WithdrawAccount) predicate.WithdrawAccount {
	return predicate.WithdrawAccount(sql.NotPredicates(p))
}
