// Code generated by ent, DO NOT EDIT.

package missionfailedfeedback

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldDeletedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldUserID, v))
}

// DeviceID applies equality check predicate on the "device_id" field. It's identical to DeviceIDEQ.
func DeviceID(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldDeviceID, v))
}

// MissionID applies equality check predicate on the "mission_id" field. It's identical to MissionIDEQ.
func MissionID(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldMissionID, v))
}

// MissionName applies equality check predicate on the "mission_name" field. It's identical to MissionNameEQ.
func MissionName(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldMissionName, v))
}

// Reason applies equality check predicate on the "reason" field. It's identical to ReasonEQ.
func Reason(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldReason, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLTE(FieldDeletedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNotIn(FieldUserID, vs...))
}

// DeviceIDEQ applies the EQ predicate on the "device_id" field.
func DeviceIDEQ(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldDeviceID, v))
}

// DeviceIDNEQ applies the NEQ predicate on the "device_id" field.
func DeviceIDNEQ(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNEQ(FieldDeviceID, v))
}

// DeviceIDIn applies the In predicate on the "device_id" field.
func DeviceIDIn(vs ...int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldIn(FieldDeviceID, vs...))
}

// DeviceIDNotIn applies the NotIn predicate on the "device_id" field.
func DeviceIDNotIn(vs ...int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNotIn(FieldDeviceID, vs...))
}

// MissionIDEQ applies the EQ predicate on the "mission_id" field.
func MissionIDEQ(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldMissionID, v))
}

// MissionIDNEQ applies the NEQ predicate on the "mission_id" field.
func MissionIDNEQ(v int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNEQ(FieldMissionID, v))
}

// MissionIDIn applies the In predicate on the "mission_id" field.
func MissionIDIn(vs ...int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldIn(FieldMissionID, vs...))
}

// MissionIDNotIn applies the NotIn predicate on the "mission_id" field.
func MissionIDNotIn(vs ...int64) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNotIn(FieldMissionID, vs...))
}

// MissionNameEQ applies the EQ predicate on the "mission_name" field.
func MissionNameEQ(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldMissionName, v))
}

// MissionNameNEQ applies the NEQ predicate on the "mission_name" field.
func MissionNameNEQ(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNEQ(FieldMissionName, v))
}

// MissionNameIn applies the In predicate on the "mission_name" field.
func MissionNameIn(vs ...string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldIn(FieldMissionName, vs...))
}

// MissionNameNotIn applies the NotIn predicate on the "mission_name" field.
func MissionNameNotIn(vs ...string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNotIn(FieldMissionName, vs...))
}

// MissionNameGT applies the GT predicate on the "mission_name" field.
func MissionNameGT(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGT(FieldMissionName, v))
}

// MissionNameGTE applies the GTE predicate on the "mission_name" field.
func MissionNameGTE(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGTE(FieldMissionName, v))
}

// MissionNameLT applies the LT predicate on the "mission_name" field.
func MissionNameLT(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLT(FieldMissionName, v))
}

// MissionNameLTE applies the LTE predicate on the "mission_name" field.
func MissionNameLTE(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLTE(FieldMissionName, v))
}

// MissionNameContains applies the Contains predicate on the "mission_name" field.
func MissionNameContains(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldContains(FieldMissionName, v))
}

// MissionNameHasPrefix applies the HasPrefix predicate on the "mission_name" field.
func MissionNameHasPrefix(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldHasPrefix(FieldMissionName, v))
}

// MissionNameHasSuffix applies the HasSuffix predicate on the "mission_name" field.
func MissionNameHasSuffix(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldHasSuffix(FieldMissionName, v))
}

// MissionNameEqualFold applies the EqualFold predicate on the "mission_name" field.
func MissionNameEqualFold(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEqualFold(FieldMissionName, v))
}

// MissionNameContainsFold applies the ContainsFold predicate on the "mission_name" field.
func MissionNameContainsFold(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldContainsFold(FieldMissionName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v enums.MissionFailedFeedbackStatus) predicate.MissionFailedFeedback {
	vc := v
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v enums.MissionFailedFeedbackStatus) predicate.MissionFailedFeedback {
	vc := v
	return predicate.MissionFailedFeedback(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...enums.MissionFailedFeedbackStatus) predicate.MissionFailedFeedback {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionFailedFeedback(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...enums.MissionFailedFeedbackStatus) predicate.MissionFailedFeedback {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MissionFailedFeedback(sql.FieldNotIn(FieldStatus, v...))
}

// ReasonEQ applies the EQ predicate on the "reason" field.
func ReasonEQ(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEQ(FieldReason, v))
}

// ReasonNEQ applies the NEQ predicate on the "reason" field.
func ReasonNEQ(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNEQ(FieldReason, v))
}

// ReasonIn applies the In predicate on the "reason" field.
func ReasonIn(vs ...string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldIn(FieldReason, vs...))
}

// ReasonNotIn applies the NotIn predicate on the "reason" field.
func ReasonNotIn(vs ...string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldNotIn(FieldReason, vs...))
}

// ReasonGT applies the GT predicate on the "reason" field.
func ReasonGT(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGT(FieldReason, v))
}

// ReasonGTE applies the GTE predicate on the "reason" field.
func ReasonGTE(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldGTE(FieldReason, v))
}

// ReasonLT applies the LT predicate on the "reason" field.
func ReasonLT(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLT(FieldReason, v))
}

// ReasonLTE applies the LTE predicate on the "reason" field.
func ReasonLTE(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldLTE(FieldReason, v))
}

// ReasonContains applies the Contains predicate on the "reason" field.
func ReasonContains(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldContains(FieldReason, v))
}

// ReasonHasPrefix applies the HasPrefix predicate on the "reason" field.
func ReasonHasPrefix(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldHasPrefix(FieldReason, v))
}

// ReasonHasSuffix applies the HasSuffix predicate on the "reason" field.
func ReasonHasSuffix(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldHasSuffix(FieldReason, v))
}

// ReasonEqualFold applies the EqualFold predicate on the "reason" field.
func ReasonEqualFold(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldEqualFold(FieldReason, v))
}

// ReasonContainsFold applies the ContainsFold predicate on the "reason" field.
func ReasonContainsFold(v string) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.FieldContainsFold(FieldReason, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDevice applies the HasEdge predicate on the "device" edge.
func HasDevice() predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeviceTable, DeviceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeviceWith applies the HasEdge predicate on the "device" edge with a given conditions (other predicates).
func HasDeviceWith(preds ...predicate.Device) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(func(s *sql.Selector) {
		step := newDeviceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMission applies the HasEdge predicate on the "mission" edge.
func HasMission() predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, MissionTable, MissionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMissionWith applies the HasEdge predicate on the "mission" edge with a given conditions (other predicates).
func HasMissionWith(preds ...predicate.Mission) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(func(s *sql.Selector) {
		step := newMissionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MissionFailedFeedback) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MissionFailedFeedback) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MissionFailedFeedback) predicate.MissionFailedFeedback {
	return predicate.MissionFailedFeedback(sql.NotPredicates(p))
}
