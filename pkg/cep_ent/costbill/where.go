// Code generated by ent, DO NOT EDIT.

package costbill

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldDeletedAt, v))
}

// IsAdd applies equality check predicate on the "is_add" field. It's identical to IsAddEQ.
func IsAdd(v bool) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldIsAdd, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldUserID, v))
}

// SerialNumber applies equality check predicate on the "serial_number" field. It's identical to SerialNumberEQ.
func SerialNumber(v string) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldSerialNumber, v))
}

// CostAccountID applies equality check predicate on the "cost_account_id" field. It's identical to CostAccountIDEQ.
func CostAccountID(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldCostAccountID, v))
}

// PureCep applies equality check predicate on the "pure_cep" field. It's identical to PureCepEQ.
func PureCep(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldPureCep, v))
}

// GiftCep applies equality check predicate on the "gift_cep" field. It's identical to GiftCepEQ.
func GiftCep(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldGiftCep, v))
}

// ReasonID applies equality check predicate on the "reason_id" field. It's identical to ReasonIDEQ.
func ReasonID(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldReasonID, v))
}

// MarketAccountID applies equality check predicate on the "market_account_id" field. It's identical to MarketAccountIDEQ.
func MarketAccountID(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldMarketAccountID, v))
}

// CampaignOrderID applies equality check predicate on the "campaign_order_id" field. It's identical to CampaignOrderIDEQ.
func CampaignOrderID(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldCampaignOrderID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.CostBill {
	return predicate.CostBill(sql.FieldLTE(FieldDeletedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldType, vs...))
}

// WayEQ applies the EQ predicate on the "way" field.
func WayEQ(v Way) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldWay, v))
}

// WayNEQ applies the NEQ predicate on the "way" field.
func WayNEQ(v Way) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldWay, v))
}

// WayIn applies the In predicate on the "way" field.
func WayIn(vs ...Way) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldWay, vs...))
}

// WayNotIn applies the NotIn predicate on the "way" field.
func WayNotIn(vs ...Way) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldWay, vs...))
}

// IsAddEQ applies the EQ predicate on the "is_add" field.
func IsAddEQ(v bool) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldIsAdd, v))
}

// IsAddNEQ applies the NEQ predicate on the "is_add" field.
func IsAddNEQ(v bool) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldIsAdd, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldUserID, vs...))
}

// SerialNumberEQ applies the EQ predicate on the "serial_number" field.
func SerialNumberEQ(v string) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldSerialNumber, v))
}

// SerialNumberNEQ applies the NEQ predicate on the "serial_number" field.
func SerialNumberNEQ(v string) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldSerialNumber, v))
}

// SerialNumberIn applies the In predicate on the "serial_number" field.
func SerialNumberIn(vs ...string) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldSerialNumber, vs...))
}

// SerialNumberNotIn applies the NotIn predicate on the "serial_number" field.
func SerialNumberNotIn(vs ...string) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldSerialNumber, vs...))
}

// SerialNumberGT applies the GT predicate on the "serial_number" field.
func SerialNumberGT(v string) predicate.CostBill {
	return predicate.CostBill(sql.FieldGT(FieldSerialNumber, v))
}

// SerialNumberGTE applies the GTE predicate on the "serial_number" field.
func SerialNumberGTE(v string) predicate.CostBill {
	return predicate.CostBill(sql.FieldGTE(FieldSerialNumber, v))
}

// SerialNumberLT applies the LT predicate on the "serial_number" field.
func SerialNumberLT(v string) predicate.CostBill {
	return predicate.CostBill(sql.FieldLT(FieldSerialNumber, v))
}

// SerialNumberLTE applies the LTE predicate on the "serial_number" field.
func SerialNumberLTE(v string) predicate.CostBill {
	return predicate.CostBill(sql.FieldLTE(FieldSerialNumber, v))
}

// SerialNumberContains applies the Contains predicate on the "serial_number" field.
func SerialNumberContains(v string) predicate.CostBill {
	return predicate.CostBill(sql.FieldContains(FieldSerialNumber, v))
}

// SerialNumberHasPrefix applies the HasPrefix predicate on the "serial_number" field.
func SerialNumberHasPrefix(v string) predicate.CostBill {
	return predicate.CostBill(sql.FieldHasPrefix(FieldSerialNumber, v))
}

// SerialNumberHasSuffix applies the HasSuffix predicate on the "serial_number" field.
func SerialNumberHasSuffix(v string) predicate.CostBill {
	return predicate.CostBill(sql.FieldHasSuffix(FieldSerialNumber, v))
}

// SerialNumberEqualFold applies the EqualFold predicate on the "serial_number" field.
func SerialNumberEqualFold(v string) predicate.CostBill {
	return predicate.CostBill(sql.FieldEqualFold(FieldSerialNumber, v))
}

// SerialNumberContainsFold applies the ContainsFold predicate on the "serial_number" field.
func SerialNumberContainsFold(v string) predicate.CostBill {
	return predicate.CostBill(sql.FieldContainsFold(FieldSerialNumber, v))
}

// CostAccountIDEQ applies the EQ predicate on the "cost_account_id" field.
func CostAccountIDEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldCostAccountID, v))
}

// CostAccountIDNEQ applies the NEQ predicate on the "cost_account_id" field.
func CostAccountIDNEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldCostAccountID, v))
}

// CostAccountIDIn applies the In predicate on the "cost_account_id" field.
func CostAccountIDIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldCostAccountID, vs...))
}

// CostAccountIDNotIn applies the NotIn predicate on the "cost_account_id" field.
func CostAccountIDNotIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldCostAccountID, vs...))
}

// PureCepEQ applies the EQ predicate on the "pure_cep" field.
func PureCepEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldPureCep, v))
}

// PureCepNEQ applies the NEQ predicate on the "pure_cep" field.
func PureCepNEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldPureCep, v))
}

// PureCepIn applies the In predicate on the "pure_cep" field.
func PureCepIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldPureCep, vs...))
}

// PureCepNotIn applies the NotIn predicate on the "pure_cep" field.
func PureCepNotIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldPureCep, vs...))
}

// PureCepGT applies the GT predicate on the "pure_cep" field.
func PureCepGT(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldGT(FieldPureCep, v))
}

// PureCepGTE applies the GTE predicate on the "pure_cep" field.
func PureCepGTE(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldGTE(FieldPureCep, v))
}

// PureCepLT applies the LT predicate on the "pure_cep" field.
func PureCepLT(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldLT(FieldPureCep, v))
}

// PureCepLTE applies the LTE predicate on the "pure_cep" field.
func PureCepLTE(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldLTE(FieldPureCep, v))
}

// GiftCepEQ applies the EQ predicate on the "gift_cep" field.
func GiftCepEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldGiftCep, v))
}

// GiftCepNEQ applies the NEQ predicate on the "gift_cep" field.
func GiftCepNEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldGiftCep, v))
}

// GiftCepIn applies the In predicate on the "gift_cep" field.
func GiftCepIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldGiftCep, vs...))
}

// GiftCepNotIn applies the NotIn predicate on the "gift_cep" field.
func GiftCepNotIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldGiftCep, vs...))
}

// GiftCepGT applies the GT predicate on the "gift_cep" field.
func GiftCepGT(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldGT(FieldGiftCep, v))
}

// GiftCepGTE applies the GTE predicate on the "gift_cep" field.
func GiftCepGTE(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldGTE(FieldGiftCep, v))
}

// GiftCepLT applies the LT predicate on the "gift_cep" field.
func GiftCepLT(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldLT(FieldGiftCep, v))
}

// GiftCepLTE applies the LTE predicate on the "gift_cep" field.
func GiftCepLTE(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldLTE(FieldGiftCep, v))
}

// ReasonIDEQ applies the EQ predicate on the "reason_id" field.
func ReasonIDEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldReasonID, v))
}

// ReasonIDNEQ applies the NEQ predicate on the "reason_id" field.
func ReasonIDNEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldReasonID, v))
}

// ReasonIDIn applies the In predicate on the "reason_id" field.
func ReasonIDIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldReasonID, vs...))
}

// ReasonIDNotIn applies the NotIn predicate on the "reason_id" field.
func ReasonIDNotIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldReasonID, vs...))
}

// ReasonIDIsNil applies the IsNil predicate on the "reason_id" field.
func ReasonIDIsNil() predicate.CostBill {
	return predicate.CostBill(sql.FieldIsNull(FieldReasonID))
}

// ReasonIDNotNil applies the NotNil predicate on the "reason_id" field.
func ReasonIDNotNil() predicate.CostBill {
	return predicate.CostBill(sql.FieldNotNull(FieldReasonID))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v enums.BillStatus) predicate.CostBill {
	vc := v
	return predicate.CostBill(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v enums.BillStatus) predicate.CostBill {
	vc := v
	return predicate.CostBill(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...enums.BillStatus) predicate.CostBill {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CostBill(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...enums.BillStatus) predicate.CostBill {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CostBill(sql.FieldNotIn(FieldStatus, v...))
}

// MarketAccountIDEQ applies the EQ predicate on the "market_account_id" field.
func MarketAccountIDEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldMarketAccountID, v))
}

// MarketAccountIDNEQ applies the NEQ predicate on the "market_account_id" field.
func MarketAccountIDNEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldMarketAccountID, v))
}

// MarketAccountIDIn applies the In predicate on the "market_account_id" field.
func MarketAccountIDIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldMarketAccountID, vs...))
}

// MarketAccountIDNotIn applies the NotIn predicate on the "market_account_id" field.
func MarketAccountIDNotIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldMarketAccountID, vs...))
}

// CampaignOrderIDEQ applies the EQ predicate on the "campaign_order_id" field.
func CampaignOrderIDEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldEQ(FieldCampaignOrderID, v))
}

// CampaignOrderIDNEQ applies the NEQ predicate on the "campaign_order_id" field.
func CampaignOrderIDNEQ(v int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNEQ(FieldCampaignOrderID, v))
}

// CampaignOrderIDIn applies the In predicate on the "campaign_order_id" field.
func CampaignOrderIDIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldIn(FieldCampaignOrderID, vs...))
}

// CampaignOrderIDNotIn applies the NotIn predicate on the "campaign_order_id" field.
func CampaignOrderIDNotIn(vs ...int64) predicate.CostBill {
	return predicate.CostBill(sql.FieldNotIn(FieldCampaignOrderID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.CostBill {
	return predicate.CostBill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.CostBill {
	return predicate.CostBill(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCostAccount applies the HasEdge predicate on the "cost_account" edge.
func HasCostAccount() predicate.CostBill {
	return predicate.CostBill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CostAccountTable, CostAccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCostAccountWith applies the HasEdge predicate on the "cost_account" edge with a given conditions (other predicates).
func HasCostAccountWith(preds ...predicate.CostAccount) predicate.CostBill {
	return predicate.CostBill(func(s *sql.Selector) {
		step := newCostAccountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRechargeOrder applies the HasEdge predicate on the "recharge_order" edge.
func HasRechargeOrder() predicate.CostBill {
	return predicate.CostBill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RechargeOrderTable, RechargeOrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRechargeOrderWith applies the HasEdge predicate on the "recharge_order" edge with a given conditions (other predicates).
func HasRechargeOrderWith(preds ...predicate.RechargeOrder) predicate.CostBill {
	return predicate.CostBill(func(s *sql.Selector) {
		step := newRechargeOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMissionConsumeOrder applies the HasEdge predicate on the "mission_consume_order" edge.
func HasMissionConsumeOrder() predicate.CostBill {
	return predicate.CostBill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MissionConsumeOrderTable, MissionConsumeOrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMissionConsumeOrderWith applies the HasEdge predicate on the "mission_consume_order" edge with a given conditions (other predicates).
func HasMissionConsumeOrderWith(preds ...predicate.MissionConsumeOrder) predicate.CostBill {
	return predicate.CostBill(func(s *sql.Selector) {
		step := newMissionConsumeOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlatformAccount applies the HasEdge predicate on the "platform_account" edge.
func HasPlatformAccount() predicate.CostBill {
	return predicate.CostBill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlatformAccountTable, PlatformAccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlatformAccountWith applies the HasEdge predicate on the "platform_account" edge with a given conditions (other predicates).
func HasPlatformAccountWith(preds ...predicate.PlatformAccount) predicate.CostBill {
	return predicate.CostBill(func(s *sql.Selector) {
		step := newPlatformAccountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCampaignOrder applies the HasEdge predicate on the "campaign_order" edge.
func HasCampaignOrder() predicate.CostBill {
	return predicate.CostBill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CampaignOrderTable, CampaignOrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCampaignOrderWith applies the HasEdge predicate on the "campaign_order" edge with a given conditions (other predicates).
func HasCampaignOrderWith(preds ...predicate.CampaignOrder) predicate.CostBill {
	return predicate.CostBill(func(s *sql.Selector) {
		step := newCampaignOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CostBill) predicate.CostBill {
	return predicate.CostBill(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CostBill) predicate.CostBill {
	return predicate.CostBill(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CostBill) predicate.CostBill {
	return predicate.CostBill(sql.NotPredicates(p))
}
