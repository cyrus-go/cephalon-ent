// Code generated by ent, DO NOT EDIT.

package platformaccount

import (
	"cephalon-ent/pkg/cep_ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldDeletedAt, v))
}

// SumTotalCep applies equality check predicate on the "sum_total_cep" field. It's identical to SumTotalCepEQ.
func SumTotalCep(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldSumTotalCep, v))
}

// TotalCep applies equality check predicate on the "total_cep" field. It's identical to TotalCepEQ.
func TotalCep(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldTotalCep, v))
}

// SumPureCep applies equality check predicate on the "sum_pure_cep" field. It's identical to SumPureCepEQ.
func SumPureCep(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldSumPureCep, v))
}

// PureCep applies equality check predicate on the "pure_cep" field. It's identical to PureCepEQ.
func PureCep(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldPureCep, v))
}

// SumGiftCep applies equality check predicate on the "sum_gift_cep" field. It's identical to SumGiftCepEQ.
func SumGiftCep(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldSumGiftCep, v))
}

// GiftCep applies equality check predicate on the "gift_cep" field. It's identical to GiftCepEQ.
func GiftCep(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldGiftCep, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldDeletedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldType, vs...))
}

// SumTotalCepEQ applies the EQ predicate on the "sum_total_cep" field.
func SumTotalCepEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldSumTotalCep, v))
}

// SumTotalCepNEQ applies the NEQ predicate on the "sum_total_cep" field.
func SumTotalCepNEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldSumTotalCep, v))
}

// SumTotalCepIn applies the In predicate on the "sum_total_cep" field.
func SumTotalCepIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldSumTotalCep, vs...))
}

// SumTotalCepNotIn applies the NotIn predicate on the "sum_total_cep" field.
func SumTotalCepNotIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldSumTotalCep, vs...))
}

// SumTotalCepGT applies the GT predicate on the "sum_total_cep" field.
func SumTotalCepGT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldSumTotalCep, v))
}

// SumTotalCepGTE applies the GTE predicate on the "sum_total_cep" field.
func SumTotalCepGTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldSumTotalCep, v))
}

// SumTotalCepLT applies the LT predicate on the "sum_total_cep" field.
func SumTotalCepLT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldSumTotalCep, v))
}

// SumTotalCepLTE applies the LTE predicate on the "sum_total_cep" field.
func SumTotalCepLTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldSumTotalCep, v))
}

// TotalCepEQ applies the EQ predicate on the "total_cep" field.
func TotalCepEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldTotalCep, v))
}

// TotalCepNEQ applies the NEQ predicate on the "total_cep" field.
func TotalCepNEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldTotalCep, v))
}

// TotalCepIn applies the In predicate on the "total_cep" field.
func TotalCepIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldTotalCep, vs...))
}

// TotalCepNotIn applies the NotIn predicate on the "total_cep" field.
func TotalCepNotIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldTotalCep, vs...))
}

// TotalCepGT applies the GT predicate on the "total_cep" field.
func TotalCepGT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldTotalCep, v))
}

// TotalCepGTE applies the GTE predicate on the "total_cep" field.
func TotalCepGTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldTotalCep, v))
}

// TotalCepLT applies the LT predicate on the "total_cep" field.
func TotalCepLT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldTotalCep, v))
}

// TotalCepLTE applies the LTE predicate on the "total_cep" field.
func TotalCepLTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldTotalCep, v))
}

// SumPureCepEQ applies the EQ predicate on the "sum_pure_cep" field.
func SumPureCepEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldSumPureCep, v))
}

// SumPureCepNEQ applies the NEQ predicate on the "sum_pure_cep" field.
func SumPureCepNEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldSumPureCep, v))
}

// SumPureCepIn applies the In predicate on the "sum_pure_cep" field.
func SumPureCepIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldSumPureCep, vs...))
}

// SumPureCepNotIn applies the NotIn predicate on the "sum_pure_cep" field.
func SumPureCepNotIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldSumPureCep, vs...))
}

// SumPureCepGT applies the GT predicate on the "sum_pure_cep" field.
func SumPureCepGT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldSumPureCep, v))
}

// SumPureCepGTE applies the GTE predicate on the "sum_pure_cep" field.
func SumPureCepGTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldSumPureCep, v))
}

// SumPureCepLT applies the LT predicate on the "sum_pure_cep" field.
func SumPureCepLT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldSumPureCep, v))
}

// SumPureCepLTE applies the LTE predicate on the "sum_pure_cep" field.
func SumPureCepLTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldSumPureCep, v))
}

// PureCepEQ applies the EQ predicate on the "pure_cep" field.
func PureCepEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldPureCep, v))
}

// PureCepNEQ applies the NEQ predicate on the "pure_cep" field.
func PureCepNEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldPureCep, v))
}

// PureCepIn applies the In predicate on the "pure_cep" field.
func PureCepIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldPureCep, vs...))
}

// PureCepNotIn applies the NotIn predicate on the "pure_cep" field.
func PureCepNotIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldPureCep, vs...))
}

// PureCepGT applies the GT predicate on the "pure_cep" field.
func PureCepGT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldPureCep, v))
}

// PureCepGTE applies the GTE predicate on the "pure_cep" field.
func PureCepGTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldPureCep, v))
}

// PureCepLT applies the LT predicate on the "pure_cep" field.
func PureCepLT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldPureCep, v))
}

// PureCepLTE applies the LTE predicate on the "pure_cep" field.
func PureCepLTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldPureCep, v))
}

// SumGiftCepEQ applies the EQ predicate on the "sum_gift_cep" field.
func SumGiftCepEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldSumGiftCep, v))
}

// SumGiftCepNEQ applies the NEQ predicate on the "sum_gift_cep" field.
func SumGiftCepNEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldSumGiftCep, v))
}

// SumGiftCepIn applies the In predicate on the "sum_gift_cep" field.
func SumGiftCepIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldSumGiftCep, vs...))
}

// SumGiftCepNotIn applies the NotIn predicate on the "sum_gift_cep" field.
func SumGiftCepNotIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldSumGiftCep, vs...))
}

// SumGiftCepGT applies the GT predicate on the "sum_gift_cep" field.
func SumGiftCepGT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldSumGiftCep, v))
}

// SumGiftCepGTE applies the GTE predicate on the "sum_gift_cep" field.
func SumGiftCepGTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldSumGiftCep, v))
}

// SumGiftCepLT applies the LT predicate on the "sum_gift_cep" field.
func SumGiftCepLT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldSumGiftCep, v))
}

// SumGiftCepLTE applies the LTE predicate on the "sum_gift_cep" field.
func SumGiftCepLTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldSumGiftCep, v))
}

// GiftCepEQ applies the EQ predicate on the "gift_cep" field.
func GiftCepEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldGiftCep, v))
}

// GiftCepNEQ applies the NEQ predicate on the "gift_cep" field.
func GiftCepNEQ(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldGiftCep, v))
}

// GiftCepIn applies the In predicate on the "gift_cep" field.
func GiftCepIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldGiftCep, vs...))
}

// GiftCepNotIn applies the NotIn predicate on the "gift_cep" field.
func GiftCepNotIn(vs ...int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldGiftCep, vs...))
}

// GiftCepGT applies the GT predicate on the "gift_cep" field.
func GiftCepGT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldGiftCep, v))
}

// GiftCepGTE applies the GTE predicate on the "gift_cep" field.
func GiftCepGTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldGiftCep, v))
}

// GiftCepLT applies the LT predicate on the "gift_cep" field.
func GiftCepLT(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldGiftCep, v))
}

// GiftCepLTE applies the LTE predicate on the "gift_cep" field.
func GiftCepLTE(v int64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldGiftCep, v))
}

// HasEarnBills applies the HasEdge predicate on the "earn_bills" edge.
func HasEarnBills() predicate.PlatformAccount {
	return predicate.PlatformAccount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EarnBillsTable, EarnBillsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEarnBillsWith applies the HasEdge predicate on the "earn_bills" edge with a given conditions (other predicates).
func HasEarnBillsWith(preds ...predicate.EarnBill) predicate.PlatformAccount {
	return predicate.PlatformAccount(func(s *sql.Selector) {
		step := newEarnBillsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PlatformAccount) predicate.PlatformAccount {
	return predicate.PlatformAccount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PlatformAccount) predicate.PlatformAccount {
	return predicate.PlatformAccount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PlatformAccount) predicate.PlatformAccount {
	return predicate.PlatformAccount(func(s *sql.Selector) {
		p(s.Not())
	})
}
