// Code generated by ent, DO NOT EDIT.

package cdkinfo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldDeletedAt, v))
}

// IssueUserID applies equality check predicate on the "issue_user_id" field. It's identical to IssueUserIDEQ.
func IssueUserID(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldIssueUserID, v))
}

// CdkNumber applies equality check predicate on the "cdk_number" field. It's identical to CdkNumberEQ.
func CdkNumber(v string) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldCdkNumber, v))
}

// GetCep applies equality check predicate on the "get_cep" field. It's identical to GetCepEQ.
func GetCep(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldGetCep, v))
}

// GetTime applies equality check predicate on the "get_time" field. It's identical to GetTimeEQ.
func GetTime(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldGetTime, v))
}

// ExpiredAt applies equality check predicate on the "expired_at" field. It's identical to ExpiredAtEQ.
func ExpiredAt(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldExpiredAt, v))
}

// UseTimes applies equality check predicate on the "use_times" field. It's identical to UseTimesEQ.
func UseTimes(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldUseTimes, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLTE(FieldDeletedAt, v))
}

// IssueUserIDEQ applies the EQ predicate on the "issue_user_id" field.
func IssueUserIDEQ(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldIssueUserID, v))
}

// IssueUserIDNEQ applies the NEQ predicate on the "issue_user_id" field.
func IssueUserIDNEQ(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNEQ(FieldIssueUserID, v))
}

// IssueUserIDIn applies the In predicate on the "issue_user_id" field.
func IssueUserIDIn(vs ...int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldIn(FieldIssueUserID, vs...))
}

// IssueUserIDNotIn applies the NotIn predicate on the "issue_user_id" field.
func IssueUserIDNotIn(vs ...int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNotIn(FieldIssueUserID, vs...))
}

// CdkNumberEQ applies the EQ predicate on the "cdk_number" field.
func CdkNumberEQ(v string) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldCdkNumber, v))
}

// CdkNumberNEQ applies the NEQ predicate on the "cdk_number" field.
func CdkNumberNEQ(v string) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNEQ(FieldCdkNumber, v))
}

// CdkNumberIn applies the In predicate on the "cdk_number" field.
func CdkNumberIn(vs ...string) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldIn(FieldCdkNumber, vs...))
}

// CdkNumberNotIn applies the NotIn predicate on the "cdk_number" field.
func CdkNumberNotIn(vs ...string) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNotIn(FieldCdkNumber, vs...))
}

// CdkNumberGT applies the GT predicate on the "cdk_number" field.
func CdkNumberGT(v string) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGT(FieldCdkNumber, v))
}

// CdkNumberGTE applies the GTE predicate on the "cdk_number" field.
func CdkNumberGTE(v string) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGTE(FieldCdkNumber, v))
}

// CdkNumberLT applies the LT predicate on the "cdk_number" field.
func CdkNumberLT(v string) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLT(FieldCdkNumber, v))
}

// CdkNumberLTE applies the LTE predicate on the "cdk_number" field.
func CdkNumberLTE(v string) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLTE(FieldCdkNumber, v))
}

// CdkNumberContains applies the Contains predicate on the "cdk_number" field.
func CdkNumberContains(v string) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldContains(FieldCdkNumber, v))
}

// CdkNumberHasPrefix applies the HasPrefix predicate on the "cdk_number" field.
func CdkNumberHasPrefix(v string) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldHasPrefix(FieldCdkNumber, v))
}

// CdkNumberHasSuffix applies the HasSuffix predicate on the "cdk_number" field.
func CdkNumberHasSuffix(v string) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldHasSuffix(FieldCdkNumber, v))
}

// CdkNumberEqualFold applies the EqualFold predicate on the "cdk_number" field.
func CdkNumberEqualFold(v string) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEqualFold(FieldCdkNumber, v))
}

// CdkNumberContainsFold applies the ContainsFold predicate on the "cdk_number" field.
func CdkNumberContainsFold(v string) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldContainsFold(FieldCdkNumber, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v enums.CDKType) predicate.CDKInfo {
	vc := v
	return predicate.CDKInfo(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v enums.CDKType) predicate.CDKInfo {
	vc := v
	return predicate.CDKInfo(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...enums.CDKType) predicate.CDKInfo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CDKInfo(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...enums.CDKType) predicate.CDKInfo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CDKInfo(sql.FieldNotIn(FieldType, v...))
}

// GetCepEQ applies the EQ predicate on the "get_cep" field.
func GetCepEQ(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldGetCep, v))
}

// GetCepNEQ applies the NEQ predicate on the "get_cep" field.
func GetCepNEQ(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNEQ(FieldGetCep, v))
}

// GetCepIn applies the In predicate on the "get_cep" field.
func GetCepIn(vs ...int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldIn(FieldGetCep, vs...))
}

// GetCepNotIn applies the NotIn predicate on the "get_cep" field.
func GetCepNotIn(vs ...int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNotIn(FieldGetCep, vs...))
}

// GetCepGT applies the GT predicate on the "get_cep" field.
func GetCepGT(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGT(FieldGetCep, v))
}

// GetCepGTE applies the GTE predicate on the "get_cep" field.
func GetCepGTE(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGTE(FieldGetCep, v))
}

// GetCepLT applies the LT predicate on the "get_cep" field.
func GetCepLT(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLT(FieldGetCep, v))
}

// GetCepLTE applies the LTE predicate on the "get_cep" field.
func GetCepLTE(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLTE(FieldGetCep, v))
}

// GetTimeEQ applies the EQ predicate on the "get_time" field.
func GetTimeEQ(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldGetTime, v))
}

// GetTimeNEQ applies the NEQ predicate on the "get_time" field.
func GetTimeNEQ(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNEQ(FieldGetTime, v))
}

// GetTimeIn applies the In predicate on the "get_time" field.
func GetTimeIn(vs ...int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldIn(FieldGetTime, vs...))
}

// GetTimeNotIn applies the NotIn predicate on the "get_time" field.
func GetTimeNotIn(vs ...int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNotIn(FieldGetTime, vs...))
}

// GetTimeGT applies the GT predicate on the "get_time" field.
func GetTimeGT(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGT(FieldGetTime, v))
}

// GetTimeGTE applies the GTE predicate on the "get_time" field.
func GetTimeGTE(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGTE(FieldGetTime, v))
}

// GetTimeLT applies the LT predicate on the "get_time" field.
func GetTimeLT(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLT(FieldGetTime, v))
}

// GetTimeLTE applies the LTE predicate on the "get_time" field.
func GetTimeLTE(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLTE(FieldGetTime, v))
}

// BillingTypeEQ applies the EQ predicate on the "billing_type" field.
func BillingTypeEQ(v enums.MissionBillingType) predicate.CDKInfo {
	vc := v
	return predicate.CDKInfo(sql.FieldEQ(FieldBillingType, vc))
}

// BillingTypeNEQ applies the NEQ predicate on the "billing_type" field.
func BillingTypeNEQ(v enums.MissionBillingType) predicate.CDKInfo {
	vc := v
	return predicate.CDKInfo(sql.FieldNEQ(FieldBillingType, vc))
}

// BillingTypeIn applies the In predicate on the "billing_type" field.
func BillingTypeIn(vs ...enums.MissionBillingType) predicate.CDKInfo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CDKInfo(sql.FieldIn(FieldBillingType, v...))
}

// BillingTypeNotIn applies the NotIn predicate on the "billing_type" field.
func BillingTypeNotIn(vs ...enums.MissionBillingType) predicate.CDKInfo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CDKInfo(sql.FieldNotIn(FieldBillingType, v...))
}

// ExpiredAtEQ applies the EQ predicate on the "expired_at" field.
func ExpiredAtEQ(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldExpiredAt, v))
}

// ExpiredAtNEQ applies the NEQ predicate on the "expired_at" field.
func ExpiredAtNEQ(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNEQ(FieldExpiredAt, v))
}

// ExpiredAtIn applies the In predicate on the "expired_at" field.
func ExpiredAtIn(vs ...time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldIn(FieldExpiredAt, vs...))
}

// ExpiredAtNotIn applies the NotIn predicate on the "expired_at" field.
func ExpiredAtNotIn(vs ...time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNotIn(FieldExpiredAt, vs...))
}

// ExpiredAtGT applies the GT predicate on the "expired_at" field.
func ExpiredAtGT(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGT(FieldExpiredAt, v))
}

// ExpiredAtGTE applies the GTE predicate on the "expired_at" field.
func ExpiredAtGTE(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGTE(FieldExpiredAt, v))
}

// ExpiredAtLT applies the LT predicate on the "expired_at" field.
func ExpiredAtLT(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLT(FieldExpiredAt, v))
}

// ExpiredAtLTE applies the LTE predicate on the "expired_at" field.
func ExpiredAtLTE(v time.Time) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLTE(FieldExpiredAt, v))
}

// ExpiredAtIsNil applies the IsNil predicate on the "expired_at" field.
func ExpiredAtIsNil() predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldIsNull(FieldExpiredAt))
}

// ExpiredAtNotNil applies the NotNil predicate on the "expired_at" field.
func ExpiredAtNotNil() predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNotNull(FieldExpiredAt))
}

// UseTimesEQ applies the EQ predicate on the "use_times" field.
func UseTimesEQ(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldEQ(FieldUseTimes, v))
}

// UseTimesNEQ applies the NEQ predicate on the "use_times" field.
func UseTimesNEQ(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNEQ(FieldUseTimes, v))
}

// UseTimesIn applies the In predicate on the "use_times" field.
func UseTimesIn(vs ...int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldIn(FieldUseTimes, vs...))
}

// UseTimesNotIn applies the NotIn predicate on the "use_times" field.
func UseTimesNotIn(vs ...int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldNotIn(FieldUseTimes, vs...))
}

// UseTimesGT applies the GT predicate on the "use_times" field.
func UseTimesGT(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGT(FieldUseTimes, v))
}

// UseTimesGTE applies the GTE predicate on the "use_times" field.
func UseTimesGTE(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldGTE(FieldUseTimes, v))
}

// UseTimesLT applies the LT predicate on the "use_times" field.
func UseTimesLT(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLT(FieldUseTimes, v))
}

// UseTimesLTE applies the LTE predicate on the "use_times" field.
func UseTimesLTE(v int64) predicate.CDKInfo {
	return predicate.CDKInfo(sql.FieldLTE(FieldUseTimes, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v enums.CDKStatus) predicate.CDKInfo {
	vc := v
	return predicate.CDKInfo(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v enums.CDKStatus) predicate.CDKInfo {
	vc := v
	return predicate.CDKInfo(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...enums.CDKStatus) predicate.CDKInfo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CDKInfo(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...enums.CDKStatus) predicate.CDKInfo {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CDKInfo(sql.FieldNotIn(FieldStatus, v...))
}

// HasIssueUser applies the HasEdge predicate on the "issue_user" edge.
func HasIssueUser() predicate.CDKInfo {
	return predicate.CDKInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, IssueUserTable, IssueUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIssueUserWith applies the HasEdge predicate on the "issue_user" edge with a given conditions (other predicates).
func HasIssueUserWith(preds ...predicate.User) predicate.CDKInfo {
	return predicate.CDKInfo(func(s *sql.Selector) {
		step := newIssueUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CDKInfo) predicate.CDKInfo {
	return predicate.CDKInfo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CDKInfo) predicate.CDKInfo {
	return predicate.CDKInfo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CDKInfo) predicate.CDKInfo {
	return predicate.CDKInfo(sql.NotPredicates(p))
}
