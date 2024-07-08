// Code generated by ent, DO NOT EDIT.

package survey

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Survey {
	return predicate.Survey(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Survey {
	return predicate.Survey(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Survey {
	return predicate.Survey(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Survey {
	return predicate.Survey(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Survey {
	return predicate.Survey(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Survey {
	return predicate.Survey(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Survey {
	return predicate.Survey(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldDeletedAt, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldTitle, v))
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldStartedAt, v))
}

// EndedAt applies equality check predicate on the "ended_at" field. It's identical to EndedAtEQ.
func EndedAt(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldEndedAt, v))
}

// SortNum applies equality check predicate on the "sort_num" field. It's identical to SortNumEQ.
func SortNum(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldSortNum, v))
}

// Group applies equality check predicate on the "group" field. It's identical to GroupEQ.
func Group(v string) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldGroup, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.Survey {
	return predicate.Survey(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.Survey {
	return predicate.Survey(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.Survey {
	return predicate.Survey(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.Survey {
	return predicate.Survey(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldLTE(FieldDeletedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Survey {
	return predicate.Survey(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Survey {
	return predicate.Survey(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Survey {
	return predicate.Survey(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Survey {
	return predicate.Survey(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Survey {
	return predicate.Survey(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Survey {
	return predicate.Survey(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Survey {
	return predicate.Survey(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Survey {
	return predicate.Survey(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Survey {
	return predicate.Survey(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Survey {
	return predicate.Survey(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Survey {
	return predicate.Survey(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Survey {
	return predicate.Survey(sql.FieldContainsFold(FieldTitle, v))
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldStartedAt, v))
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldNEQ(FieldStartedAt, v))
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldIn(FieldStartedAt, vs...))
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldNotIn(FieldStartedAt, vs...))
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldGT(FieldStartedAt, v))
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldGTE(FieldStartedAt, v))
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldLT(FieldStartedAt, v))
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldLTE(FieldStartedAt, v))
}

// StartedAtIsNil applies the IsNil predicate on the "started_at" field.
func StartedAtIsNil() predicate.Survey {
	return predicate.Survey(sql.FieldIsNull(FieldStartedAt))
}

// StartedAtNotNil applies the NotNil predicate on the "started_at" field.
func StartedAtNotNil() predicate.Survey {
	return predicate.Survey(sql.FieldNotNull(FieldStartedAt))
}

// EndedAtEQ applies the EQ predicate on the "ended_at" field.
func EndedAtEQ(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldEndedAt, v))
}

// EndedAtNEQ applies the NEQ predicate on the "ended_at" field.
func EndedAtNEQ(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldNEQ(FieldEndedAt, v))
}

// EndedAtIn applies the In predicate on the "ended_at" field.
func EndedAtIn(vs ...time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldIn(FieldEndedAt, vs...))
}

// EndedAtNotIn applies the NotIn predicate on the "ended_at" field.
func EndedAtNotIn(vs ...time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldNotIn(FieldEndedAt, vs...))
}

// EndedAtGT applies the GT predicate on the "ended_at" field.
func EndedAtGT(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldGT(FieldEndedAt, v))
}

// EndedAtGTE applies the GTE predicate on the "ended_at" field.
func EndedAtGTE(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldGTE(FieldEndedAt, v))
}

// EndedAtLT applies the LT predicate on the "ended_at" field.
func EndedAtLT(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldLT(FieldEndedAt, v))
}

// EndedAtLTE applies the LTE predicate on the "ended_at" field.
func EndedAtLTE(v time.Time) predicate.Survey {
	return predicate.Survey(sql.FieldLTE(FieldEndedAt, v))
}

// EndedAtIsNil applies the IsNil predicate on the "ended_at" field.
func EndedAtIsNil() predicate.Survey {
	return predicate.Survey(sql.FieldIsNull(FieldEndedAt))
}

// EndedAtNotNil applies the NotNil predicate on the "ended_at" field.
func EndedAtNotNil() predicate.Survey {
	return predicate.Survey(sql.FieldNotNull(FieldEndedAt))
}

// SortNumEQ applies the EQ predicate on the "sort_num" field.
func SortNumEQ(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldSortNum, v))
}

// SortNumNEQ applies the NEQ predicate on the "sort_num" field.
func SortNumNEQ(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldNEQ(FieldSortNum, v))
}

// SortNumIn applies the In predicate on the "sort_num" field.
func SortNumIn(vs ...int64) predicate.Survey {
	return predicate.Survey(sql.FieldIn(FieldSortNum, vs...))
}

// SortNumNotIn applies the NotIn predicate on the "sort_num" field.
func SortNumNotIn(vs ...int64) predicate.Survey {
	return predicate.Survey(sql.FieldNotIn(FieldSortNum, vs...))
}

// SortNumGT applies the GT predicate on the "sort_num" field.
func SortNumGT(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldGT(FieldSortNum, v))
}

// SortNumGTE applies the GTE predicate on the "sort_num" field.
func SortNumGTE(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldGTE(FieldSortNum, v))
}

// SortNumLT applies the LT predicate on the "sort_num" field.
func SortNumLT(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldLT(FieldSortNum, v))
}

// SortNumLTE applies the LTE predicate on the "sort_num" field.
func SortNumLTE(v int64) predicate.Survey {
	return predicate.Survey(sql.FieldLTE(FieldSortNum, v))
}

// GroupEQ applies the EQ predicate on the "group" field.
func GroupEQ(v string) predicate.Survey {
	return predicate.Survey(sql.FieldEQ(FieldGroup, v))
}

// GroupNEQ applies the NEQ predicate on the "group" field.
func GroupNEQ(v string) predicate.Survey {
	return predicate.Survey(sql.FieldNEQ(FieldGroup, v))
}

// GroupIn applies the In predicate on the "group" field.
func GroupIn(vs ...string) predicate.Survey {
	return predicate.Survey(sql.FieldIn(FieldGroup, vs...))
}

// GroupNotIn applies the NotIn predicate on the "group" field.
func GroupNotIn(vs ...string) predicate.Survey {
	return predicate.Survey(sql.FieldNotIn(FieldGroup, vs...))
}

// GroupGT applies the GT predicate on the "group" field.
func GroupGT(v string) predicate.Survey {
	return predicate.Survey(sql.FieldGT(FieldGroup, v))
}

// GroupGTE applies the GTE predicate on the "group" field.
func GroupGTE(v string) predicate.Survey {
	return predicate.Survey(sql.FieldGTE(FieldGroup, v))
}

// GroupLT applies the LT predicate on the "group" field.
func GroupLT(v string) predicate.Survey {
	return predicate.Survey(sql.FieldLT(FieldGroup, v))
}

// GroupLTE applies the LTE predicate on the "group" field.
func GroupLTE(v string) predicate.Survey {
	return predicate.Survey(sql.FieldLTE(FieldGroup, v))
}

// GroupContains applies the Contains predicate on the "group" field.
func GroupContains(v string) predicate.Survey {
	return predicate.Survey(sql.FieldContains(FieldGroup, v))
}

// GroupHasPrefix applies the HasPrefix predicate on the "group" field.
func GroupHasPrefix(v string) predicate.Survey {
	return predicate.Survey(sql.FieldHasPrefix(FieldGroup, v))
}

// GroupHasSuffix applies the HasSuffix predicate on the "group" field.
func GroupHasSuffix(v string) predicate.Survey {
	return predicate.Survey(sql.FieldHasSuffix(FieldGroup, v))
}

// GroupEqualFold applies the EqualFold predicate on the "group" field.
func GroupEqualFold(v string) predicate.Survey {
	return predicate.Survey(sql.FieldEqualFold(FieldGroup, v))
}

// GroupContainsFold applies the ContainsFold predicate on the "group" field.
func GroupContainsFold(v string) predicate.Survey {
	return predicate.Survey(sql.FieldContainsFold(FieldGroup, v))
}

// HasSurveyQuestions applies the HasEdge predicate on the "survey_questions" edge.
func HasSurveyQuestions() predicate.Survey {
	return predicate.Survey(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SurveyQuestionsTable, SurveyQuestionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSurveyQuestionsWith applies the HasEdge predicate on the "survey_questions" edge with a given conditions (other predicates).
func HasSurveyQuestionsWith(preds ...predicate.SurveyQuestion) predicate.Survey {
	return predicate.Survey(func(s *sql.Selector) {
		step := newSurveyQuestionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSurveyResponses applies the HasEdge predicate on the "survey_responses" edge.
func HasSurveyResponses() predicate.Survey {
	return predicate.Survey(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SurveyResponsesTable, SurveyResponsesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSurveyResponsesWith applies the HasEdge predicate on the "survey_responses" edge with a given conditions (other predicates).
func HasSurveyResponsesWith(preds ...predicate.SurveyResponse) predicate.Survey {
	return predicate.Survey(func(s *sql.Selector) {
		step := newSurveyResponsesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Survey) predicate.Survey {
	return predicate.Survey(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Survey) predicate.Survey {
	return predicate.Survey(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Survey) predicate.Survey {
	return predicate.Survey(sql.NotPredicates(p))
}
