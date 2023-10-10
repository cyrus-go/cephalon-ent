// Code generated by ent, DO NOT EDIT.

package renewalagreement

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldDeletedAt, v))
}

// NextPayTime applies equality check predicate on the "next_pay_time" field. It's identical to NextPayTimeEQ.
func NextPayTime(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldNextPayTime, v))
}

// SymbolID applies equality check predicate on the "symbol_id" field. It's identical to SymbolIDEQ.
func SymbolID(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldSymbolID, v))
}

// FirstPay applies equality check predicate on the "first_pay" field. It's identical to FirstPayEQ.
func FirstPay(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldFirstPay, v))
}

// AfterPay applies equality check predicate on the "after_pay" field. It's identical to AfterPayEQ.
func AfterPay(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldAfterPay, v))
}

// LastWarningTime applies equality check predicate on the "last_warning_time" field. It's identical to LastWarningTimeEQ.
func LastWarningTime(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldLastWarningTime, v))
}

// SubFinishedTime applies equality check predicate on the "sub_finished_time" field. It's identical to SubFinishedTimeEQ.
func SubFinishedTime(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldSubFinishedTime, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldUserID, v))
}

// MissionID applies equality check predicate on the "mission_id" field. It's identical to MissionIDEQ.
func MissionID(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldMissionID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLTE(FieldDeletedAt, v))
}

// NextPayTimeEQ applies the EQ predicate on the "next_pay_time" field.
func NextPayTimeEQ(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldNextPayTime, v))
}

// NextPayTimeNEQ applies the NEQ predicate on the "next_pay_time" field.
func NextPayTimeNEQ(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldNextPayTime, v))
}

// NextPayTimeIn applies the In predicate on the "next_pay_time" field.
func NextPayTimeIn(vs ...time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldIn(FieldNextPayTime, vs...))
}

// NextPayTimeNotIn applies the NotIn predicate on the "next_pay_time" field.
func NextPayTimeNotIn(vs ...time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldNextPayTime, vs...))
}

// NextPayTimeGT applies the GT predicate on the "next_pay_time" field.
func NextPayTimeGT(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGT(FieldNextPayTime, v))
}

// NextPayTimeGTE applies the GTE predicate on the "next_pay_time" field.
func NextPayTimeGTE(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGTE(FieldNextPayTime, v))
}

// NextPayTimeLT applies the LT predicate on the "next_pay_time" field.
func NextPayTimeLT(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLT(FieldNextPayTime, v))
}

// NextPayTimeLTE applies the LTE predicate on the "next_pay_time" field.
func NextPayTimeLTE(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLTE(FieldNextPayTime, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v enums.RenewalType) predicate.RenewalAgreement {
	vc := v
	return predicate.RenewalAgreement(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v enums.RenewalType) predicate.RenewalAgreement {
	vc := v
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...enums.RenewalType) predicate.RenewalAgreement {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RenewalAgreement(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...enums.RenewalType) predicate.RenewalAgreement {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldType, v...))
}

// SubStatusEQ applies the EQ predicate on the "sub_status" field.
func SubStatusEQ(v enums.RenewalSubStatus) predicate.RenewalAgreement {
	vc := v
	return predicate.RenewalAgreement(sql.FieldEQ(FieldSubStatus, vc))
}

// SubStatusNEQ applies the NEQ predicate on the "sub_status" field.
func SubStatusNEQ(v enums.RenewalSubStatus) predicate.RenewalAgreement {
	vc := v
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldSubStatus, vc))
}

// SubStatusIn applies the In predicate on the "sub_status" field.
func SubStatusIn(vs ...enums.RenewalSubStatus) predicate.RenewalAgreement {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RenewalAgreement(sql.FieldIn(FieldSubStatus, v...))
}

// SubStatusNotIn applies the NotIn predicate on the "sub_status" field.
func SubStatusNotIn(vs ...enums.RenewalSubStatus) predicate.RenewalAgreement {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldSubStatus, v...))
}

// PayStatusEQ applies the EQ predicate on the "pay_status" field.
func PayStatusEQ(v enums.RenewalPayStatus) predicate.RenewalAgreement {
	vc := v
	return predicate.RenewalAgreement(sql.FieldEQ(FieldPayStatus, vc))
}

// PayStatusNEQ applies the NEQ predicate on the "pay_status" field.
func PayStatusNEQ(v enums.RenewalPayStatus) predicate.RenewalAgreement {
	vc := v
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldPayStatus, vc))
}

// PayStatusIn applies the In predicate on the "pay_status" field.
func PayStatusIn(vs ...enums.RenewalPayStatus) predicate.RenewalAgreement {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RenewalAgreement(sql.FieldIn(FieldPayStatus, v...))
}

// PayStatusNotIn applies the NotIn predicate on the "pay_status" field.
func PayStatusNotIn(vs ...enums.RenewalPayStatus) predicate.RenewalAgreement {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldPayStatus, v...))
}

// SymbolIDEQ applies the EQ predicate on the "symbol_id" field.
func SymbolIDEQ(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldSymbolID, v))
}

// SymbolIDNEQ applies the NEQ predicate on the "symbol_id" field.
func SymbolIDNEQ(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldSymbolID, v))
}

// SymbolIDIn applies the In predicate on the "symbol_id" field.
func SymbolIDIn(vs ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldIn(FieldSymbolID, vs...))
}

// SymbolIDNotIn applies the NotIn predicate on the "symbol_id" field.
func SymbolIDNotIn(vs ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldSymbolID, vs...))
}

// SymbolIDGT applies the GT predicate on the "symbol_id" field.
func SymbolIDGT(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGT(FieldSymbolID, v))
}

// SymbolIDGTE applies the GTE predicate on the "symbol_id" field.
func SymbolIDGTE(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGTE(FieldSymbolID, v))
}

// SymbolIDLT applies the LT predicate on the "symbol_id" field.
func SymbolIDLT(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLT(FieldSymbolID, v))
}

// SymbolIDLTE applies the LTE predicate on the "symbol_id" field.
func SymbolIDLTE(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLTE(FieldSymbolID, v))
}

// FirstPayEQ applies the EQ predicate on the "first_pay" field.
func FirstPayEQ(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldFirstPay, v))
}

// FirstPayNEQ applies the NEQ predicate on the "first_pay" field.
func FirstPayNEQ(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldFirstPay, v))
}

// FirstPayIn applies the In predicate on the "first_pay" field.
func FirstPayIn(vs ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldIn(FieldFirstPay, vs...))
}

// FirstPayNotIn applies the NotIn predicate on the "first_pay" field.
func FirstPayNotIn(vs ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldFirstPay, vs...))
}

// FirstPayGT applies the GT predicate on the "first_pay" field.
func FirstPayGT(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGT(FieldFirstPay, v))
}

// FirstPayGTE applies the GTE predicate on the "first_pay" field.
func FirstPayGTE(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGTE(FieldFirstPay, v))
}

// FirstPayLT applies the LT predicate on the "first_pay" field.
func FirstPayLT(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLT(FieldFirstPay, v))
}

// FirstPayLTE applies the LTE predicate on the "first_pay" field.
func FirstPayLTE(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLTE(FieldFirstPay, v))
}

// AfterPayEQ applies the EQ predicate on the "after_pay" field.
func AfterPayEQ(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldAfterPay, v))
}

// AfterPayNEQ applies the NEQ predicate on the "after_pay" field.
func AfterPayNEQ(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldAfterPay, v))
}

// AfterPayIn applies the In predicate on the "after_pay" field.
func AfterPayIn(vs ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldIn(FieldAfterPay, vs...))
}

// AfterPayNotIn applies the NotIn predicate on the "after_pay" field.
func AfterPayNotIn(vs ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldAfterPay, vs...))
}

// AfterPayGT applies the GT predicate on the "after_pay" field.
func AfterPayGT(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGT(FieldAfterPay, v))
}

// AfterPayGTE applies the GTE predicate on the "after_pay" field.
func AfterPayGTE(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGTE(FieldAfterPay, v))
}

// AfterPayLT applies the LT predicate on the "after_pay" field.
func AfterPayLT(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLT(FieldAfterPay, v))
}

// AfterPayLTE applies the LTE predicate on the "after_pay" field.
func AfterPayLTE(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLTE(FieldAfterPay, v))
}

// LastWarningTimeEQ applies the EQ predicate on the "last_warning_time" field.
func LastWarningTimeEQ(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldLastWarningTime, v))
}

// LastWarningTimeNEQ applies the NEQ predicate on the "last_warning_time" field.
func LastWarningTimeNEQ(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldLastWarningTime, v))
}

// LastWarningTimeIn applies the In predicate on the "last_warning_time" field.
func LastWarningTimeIn(vs ...time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldIn(FieldLastWarningTime, vs...))
}

// LastWarningTimeNotIn applies the NotIn predicate on the "last_warning_time" field.
func LastWarningTimeNotIn(vs ...time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldLastWarningTime, vs...))
}

// LastWarningTimeGT applies the GT predicate on the "last_warning_time" field.
func LastWarningTimeGT(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGT(FieldLastWarningTime, v))
}

// LastWarningTimeGTE applies the GTE predicate on the "last_warning_time" field.
func LastWarningTimeGTE(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGTE(FieldLastWarningTime, v))
}

// LastWarningTimeLT applies the LT predicate on the "last_warning_time" field.
func LastWarningTimeLT(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLT(FieldLastWarningTime, v))
}

// LastWarningTimeLTE applies the LTE predicate on the "last_warning_time" field.
func LastWarningTimeLTE(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLTE(FieldLastWarningTime, v))
}

// SubFinishedTimeEQ applies the EQ predicate on the "sub_finished_time" field.
func SubFinishedTimeEQ(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldSubFinishedTime, v))
}

// SubFinishedTimeNEQ applies the NEQ predicate on the "sub_finished_time" field.
func SubFinishedTimeNEQ(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldSubFinishedTime, v))
}

// SubFinishedTimeIn applies the In predicate on the "sub_finished_time" field.
func SubFinishedTimeIn(vs ...time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldIn(FieldSubFinishedTime, vs...))
}

// SubFinishedTimeNotIn applies the NotIn predicate on the "sub_finished_time" field.
func SubFinishedTimeNotIn(vs ...time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldSubFinishedTime, vs...))
}

// SubFinishedTimeGT applies the GT predicate on the "sub_finished_time" field.
func SubFinishedTimeGT(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGT(FieldSubFinishedTime, v))
}

// SubFinishedTimeGTE applies the GTE predicate on the "sub_finished_time" field.
func SubFinishedTimeGTE(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldGTE(FieldSubFinishedTime, v))
}

// SubFinishedTimeLT applies the LT predicate on the "sub_finished_time" field.
func SubFinishedTimeLT(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLT(FieldSubFinishedTime, v))
}

// SubFinishedTimeLTE applies the LTE predicate on the "sub_finished_time" field.
func SubFinishedTimeLTE(v time.Time) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldLTE(FieldSubFinishedTime, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldUserID, vs...))
}

// MissionIDEQ applies the EQ predicate on the "mission_id" field.
func MissionIDEQ(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldEQ(FieldMissionID, v))
}

// MissionIDNEQ applies the NEQ predicate on the "mission_id" field.
func MissionIDNEQ(v int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNEQ(FieldMissionID, v))
}

// MissionIDIn applies the In predicate on the "mission_id" field.
func MissionIDIn(vs ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldIn(FieldMissionID, vs...))
}

// MissionIDNotIn applies the NotIn predicate on the "mission_id" field.
func MissionIDNotIn(vs ...int64) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.FieldNotIn(FieldMissionID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.RenewalAgreement {
	return predicate.RenewalAgreement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMission applies the HasEdge predicate on the "mission" edge.
func HasMission() predicate.RenewalAgreement {
	return predicate.RenewalAgreement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, MissionTable, MissionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMissionWith applies the HasEdge predicate on the "mission" edge with a given conditions (other predicates).
func HasMissionWith(preds ...predicate.Mission) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(func(s *sql.Selector) {
		step := newMissionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RenewalAgreement) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RenewalAgreement) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RenewalAgreement) predicate.RenewalAgreement {
	return predicate.RenewalAgreement(sql.NotPredicates(p))
}
