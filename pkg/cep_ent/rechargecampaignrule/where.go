// Code generated by ent, DO NOT EDIT.

package rechargecampaignrule

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldDeletedAt, v))
}

// LittleValue applies equality check predicate on the "little_value" field. It's identical to LittleValueEQ.
func LittleValue(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldLittleValue, v))
}

// LargeValue applies equality check predicate on the "large_value" field. It's identical to LargeValueEQ.
func LargeValue(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldLargeValue, v))
}

// GiftPercent applies equality check predicate on the "gift_percent" field. It's identical to GiftPercentEQ.
func GiftPercent(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldGiftPercent, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLTE(FieldDeletedAt, v))
}

// LittleValueEQ applies the EQ predicate on the "little_value" field.
func LittleValueEQ(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldLittleValue, v))
}

// LittleValueNEQ applies the NEQ predicate on the "little_value" field.
func LittleValueNEQ(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNEQ(FieldLittleValue, v))
}

// LittleValueIn applies the In predicate on the "little_value" field.
func LittleValueIn(vs ...int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldIn(FieldLittleValue, vs...))
}

// LittleValueNotIn applies the NotIn predicate on the "little_value" field.
func LittleValueNotIn(vs ...int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNotIn(FieldLittleValue, vs...))
}

// LittleValueGT applies the GT predicate on the "little_value" field.
func LittleValueGT(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGT(FieldLittleValue, v))
}

// LittleValueGTE applies the GTE predicate on the "little_value" field.
func LittleValueGTE(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGTE(FieldLittleValue, v))
}

// LittleValueLT applies the LT predicate on the "little_value" field.
func LittleValueLT(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLT(FieldLittleValue, v))
}

// LittleValueLTE applies the LTE predicate on the "little_value" field.
func LittleValueLTE(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLTE(FieldLittleValue, v))
}

// LargeValueEQ applies the EQ predicate on the "large_value" field.
func LargeValueEQ(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldLargeValue, v))
}

// LargeValueNEQ applies the NEQ predicate on the "large_value" field.
func LargeValueNEQ(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNEQ(FieldLargeValue, v))
}

// LargeValueIn applies the In predicate on the "large_value" field.
func LargeValueIn(vs ...int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldIn(FieldLargeValue, vs...))
}

// LargeValueNotIn applies the NotIn predicate on the "large_value" field.
func LargeValueNotIn(vs ...int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNotIn(FieldLargeValue, vs...))
}

// LargeValueGT applies the GT predicate on the "large_value" field.
func LargeValueGT(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGT(FieldLargeValue, v))
}

// LargeValueGTE applies the GTE predicate on the "large_value" field.
func LargeValueGTE(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGTE(FieldLargeValue, v))
}

// LargeValueLT applies the LT predicate on the "large_value" field.
func LargeValueLT(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLT(FieldLargeValue, v))
}

// LargeValueLTE applies the LTE predicate on the "large_value" field.
func LargeValueLTE(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLTE(FieldLargeValue, v))
}

// GiftPercentEQ applies the EQ predicate on the "gift_percent" field.
func GiftPercentEQ(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldEQ(FieldGiftPercent, v))
}

// GiftPercentNEQ applies the NEQ predicate on the "gift_percent" field.
func GiftPercentNEQ(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNEQ(FieldGiftPercent, v))
}

// GiftPercentIn applies the In predicate on the "gift_percent" field.
func GiftPercentIn(vs ...int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldIn(FieldGiftPercent, vs...))
}

// GiftPercentNotIn applies the NotIn predicate on the "gift_percent" field.
func GiftPercentNotIn(vs ...int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldNotIn(FieldGiftPercent, vs...))
}

// GiftPercentGT applies the GT predicate on the "gift_percent" field.
func GiftPercentGT(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGT(FieldGiftPercent, v))
}

// GiftPercentGTE applies the GTE predicate on the "gift_percent" field.
func GiftPercentGTE(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldGTE(FieldGiftPercent, v))
}

// GiftPercentLT applies the LT predicate on the "gift_percent" field.
func GiftPercentLT(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLT(FieldGiftPercent, v))
}

// GiftPercentLTE applies the LTE predicate on the "gift_percent" field.
func GiftPercentLTE(v int64) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.FieldLTE(FieldGiftPercent, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RechargeCampaignRule) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RechargeCampaignRule) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RechargeCampaignRule) predicate.RechargeCampaignRule {
	return predicate.RechargeCampaignRule(sql.NotPredicates(p))
}
