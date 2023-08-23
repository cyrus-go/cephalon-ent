// Code generated by ent, DO NOT EDIT.

package enummissionstatus

import (
	"cephalon-ent/pkg/cep_ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldDeletedAt, v))
}

// FrontStatus applies equality check predicate on the "front_status" field. It's identical to FrontStatusEQ.
func FrontStatus(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldFrontStatus, v))
}

// MissionType applies equality check predicate on the "mission_type" field. It's identical to MissionTypeEQ.
func MissionType(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldMissionType, v))
}

// MissionStatus applies equality check predicate on the "mission_status" field. It's identical to MissionStatusEQ.
func MissionStatus(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldMissionStatus, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLTE(FieldDeletedAt, v))
}

// FrontStatusEQ applies the EQ predicate on the "front_status" field.
func FrontStatusEQ(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldFrontStatus, v))
}

// FrontStatusNEQ applies the NEQ predicate on the "front_status" field.
func FrontStatusNEQ(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNEQ(FieldFrontStatus, v))
}

// FrontStatusIn applies the In predicate on the "front_status" field.
func FrontStatusIn(vs ...string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldIn(FieldFrontStatus, vs...))
}

// FrontStatusNotIn applies the NotIn predicate on the "front_status" field.
func FrontStatusNotIn(vs ...string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNotIn(FieldFrontStatus, vs...))
}

// FrontStatusGT applies the GT predicate on the "front_status" field.
func FrontStatusGT(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGT(FieldFrontStatus, v))
}

// FrontStatusGTE applies the GTE predicate on the "front_status" field.
func FrontStatusGTE(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGTE(FieldFrontStatus, v))
}

// FrontStatusLT applies the LT predicate on the "front_status" field.
func FrontStatusLT(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLT(FieldFrontStatus, v))
}

// FrontStatusLTE applies the LTE predicate on the "front_status" field.
func FrontStatusLTE(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLTE(FieldFrontStatus, v))
}

// FrontStatusContains applies the Contains predicate on the "front_status" field.
func FrontStatusContains(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldContains(FieldFrontStatus, v))
}

// FrontStatusHasPrefix applies the HasPrefix predicate on the "front_status" field.
func FrontStatusHasPrefix(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldHasPrefix(FieldFrontStatus, v))
}

// FrontStatusHasSuffix applies the HasSuffix predicate on the "front_status" field.
func FrontStatusHasSuffix(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldHasSuffix(FieldFrontStatus, v))
}

// FrontStatusEqualFold applies the EqualFold predicate on the "front_status" field.
func FrontStatusEqualFold(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEqualFold(FieldFrontStatus, v))
}

// FrontStatusContainsFold applies the ContainsFold predicate on the "front_status" field.
func FrontStatusContainsFold(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldContainsFold(FieldFrontStatus, v))
}

// MissionTypeEQ applies the EQ predicate on the "mission_type" field.
func MissionTypeEQ(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldMissionType, v))
}

// MissionTypeNEQ applies the NEQ predicate on the "mission_type" field.
func MissionTypeNEQ(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNEQ(FieldMissionType, v))
}

// MissionTypeIn applies the In predicate on the "mission_type" field.
func MissionTypeIn(vs ...string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldIn(FieldMissionType, vs...))
}

// MissionTypeNotIn applies the NotIn predicate on the "mission_type" field.
func MissionTypeNotIn(vs ...string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNotIn(FieldMissionType, vs...))
}

// MissionTypeGT applies the GT predicate on the "mission_type" field.
func MissionTypeGT(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGT(FieldMissionType, v))
}

// MissionTypeGTE applies the GTE predicate on the "mission_type" field.
func MissionTypeGTE(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGTE(FieldMissionType, v))
}

// MissionTypeLT applies the LT predicate on the "mission_type" field.
func MissionTypeLT(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLT(FieldMissionType, v))
}

// MissionTypeLTE applies the LTE predicate on the "mission_type" field.
func MissionTypeLTE(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLTE(FieldMissionType, v))
}

// MissionTypeContains applies the Contains predicate on the "mission_type" field.
func MissionTypeContains(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldContains(FieldMissionType, v))
}

// MissionTypeHasPrefix applies the HasPrefix predicate on the "mission_type" field.
func MissionTypeHasPrefix(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldHasPrefix(FieldMissionType, v))
}

// MissionTypeHasSuffix applies the HasSuffix predicate on the "mission_type" field.
func MissionTypeHasSuffix(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldHasSuffix(FieldMissionType, v))
}

// MissionTypeEqualFold applies the EqualFold predicate on the "mission_type" field.
func MissionTypeEqualFold(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEqualFold(FieldMissionType, v))
}

// MissionTypeContainsFold applies the ContainsFold predicate on the "mission_type" field.
func MissionTypeContainsFold(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldContainsFold(FieldMissionType, v))
}

// MissionStatusEQ applies the EQ predicate on the "mission_status" field.
func MissionStatusEQ(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEQ(FieldMissionStatus, v))
}

// MissionStatusNEQ applies the NEQ predicate on the "mission_status" field.
func MissionStatusNEQ(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNEQ(FieldMissionStatus, v))
}

// MissionStatusIn applies the In predicate on the "mission_status" field.
func MissionStatusIn(vs ...string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldIn(FieldMissionStatus, vs...))
}

// MissionStatusNotIn applies the NotIn predicate on the "mission_status" field.
func MissionStatusNotIn(vs ...string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldNotIn(FieldMissionStatus, vs...))
}

// MissionStatusGT applies the GT predicate on the "mission_status" field.
func MissionStatusGT(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGT(FieldMissionStatus, v))
}

// MissionStatusGTE applies the GTE predicate on the "mission_status" field.
func MissionStatusGTE(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldGTE(FieldMissionStatus, v))
}

// MissionStatusLT applies the LT predicate on the "mission_status" field.
func MissionStatusLT(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLT(FieldMissionStatus, v))
}

// MissionStatusLTE applies the LTE predicate on the "mission_status" field.
func MissionStatusLTE(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldLTE(FieldMissionStatus, v))
}

// MissionStatusContains applies the Contains predicate on the "mission_status" field.
func MissionStatusContains(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldContains(FieldMissionStatus, v))
}

// MissionStatusHasPrefix applies the HasPrefix predicate on the "mission_status" field.
func MissionStatusHasPrefix(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldHasPrefix(FieldMissionStatus, v))
}

// MissionStatusHasSuffix applies the HasSuffix predicate on the "mission_status" field.
func MissionStatusHasSuffix(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldHasSuffix(FieldMissionStatus, v))
}

// MissionStatusEqualFold applies the EqualFold predicate on the "mission_status" field.
func MissionStatusEqualFold(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldEqualFold(FieldMissionStatus, v))
}

// MissionStatusContainsFold applies the ContainsFold predicate on the "mission_status" field.
func MissionStatusContainsFold(v string) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(sql.FieldContainsFold(FieldMissionStatus, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EnumMissionStatus) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EnumMissionStatus) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(func(s *sql.Selector) {
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
func Not(p predicate.EnumMissionStatus) predicate.EnumMissionStatus {
	return predicate.EnumMissionStatus(func(s *sql.Selector) {
		p(s.Not())
	})
}
