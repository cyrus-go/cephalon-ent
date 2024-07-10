// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/predicate"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/surveyanswer"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/surveyquestion"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/surveyresponse"
)

// SurveyAnswerUpdate is the builder for updating SurveyAnswer entities.
type SurveyAnswerUpdate struct {
	config
	hooks     []Hook
	mutation  *SurveyAnswerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SurveyAnswerUpdate builder.
func (sau *SurveyAnswerUpdate) Where(ps ...predicate.SurveyAnswer) *SurveyAnswerUpdate {
	sau.mutation.Where(ps...)
	return sau
}

// SetCreatedBy sets the "created_by" field.
func (sau *SurveyAnswerUpdate) SetCreatedBy(i int64) *SurveyAnswerUpdate {
	sau.mutation.ResetCreatedBy()
	sau.mutation.SetCreatedBy(i)
	return sau
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (sau *SurveyAnswerUpdate) SetNillableCreatedBy(i *int64) *SurveyAnswerUpdate {
	if i != nil {
		sau.SetCreatedBy(*i)
	}
	return sau
}

// AddCreatedBy adds i to the "created_by" field.
func (sau *SurveyAnswerUpdate) AddCreatedBy(i int64) *SurveyAnswerUpdate {
	sau.mutation.AddCreatedBy(i)
	return sau
}

// SetUpdatedBy sets the "updated_by" field.
func (sau *SurveyAnswerUpdate) SetUpdatedBy(i int64) *SurveyAnswerUpdate {
	sau.mutation.ResetUpdatedBy()
	sau.mutation.SetUpdatedBy(i)
	return sau
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sau *SurveyAnswerUpdate) SetNillableUpdatedBy(i *int64) *SurveyAnswerUpdate {
	if i != nil {
		sau.SetUpdatedBy(*i)
	}
	return sau
}

// AddUpdatedBy adds i to the "updated_by" field.
func (sau *SurveyAnswerUpdate) AddUpdatedBy(i int64) *SurveyAnswerUpdate {
	sau.mutation.AddUpdatedBy(i)
	return sau
}

// SetUpdatedAt sets the "updated_at" field.
func (sau *SurveyAnswerUpdate) SetUpdatedAt(t time.Time) *SurveyAnswerUpdate {
	sau.mutation.SetUpdatedAt(t)
	return sau
}

// SetDeletedAt sets the "deleted_at" field.
func (sau *SurveyAnswerUpdate) SetDeletedAt(t time.Time) *SurveyAnswerUpdate {
	sau.mutation.SetDeletedAt(t)
	return sau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sau *SurveyAnswerUpdate) SetNillableDeletedAt(t *time.Time) *SurveyAnswerUpdate {
	if t != nil {
		sau.SetDeletedAt(*t)
	}
	return sau
}

// SetSurveyResponseID sets the "survey_response_id" field.
func (sau *SurveyAnswerUpdate) SetSurveyResponseID(i int64) *SurveyAnswerUpdate {
	sau.mutation.SetSurveyResponseID(i)
	return sau
}

// SetNillableSurveyResponseID sets the "survey_response_id" field if the given value is not nil.
func (sau *SurveyAnswerUpdate) SetNillableSurveyResponseID(i *int64) *SurveyAnswerUpdate {
	if i != nil {
		sau.SetSurveyResponseID(*i)
	}
	return sau
}

// SetSurveyQuestionID sets the "survey_question_id" field.
func (sau *SurveyAnswerUpdate) SetSurveyQuestionID(i int64) *SurveyAnswerUpdate {
	sau.mutation.SetSurveyQuestionID(i)
	return sau
}

// SetNillableSurveyQuestionID sets the "survey_question_id" field if the given value is not nil.
func (sau *SurveyAnswerUpdate) SetNillableSurveyQuestionID(i *int64) *SurveyAnswerUpdate {
	if i != nil {
		sau.SetSurveyQuestionID(*i)
	}
	return sau
}

// SetSurveyAnswer sets the "survey_answer" field.
func (sau *SurveyAnswerUpdate) SetSurveyAnswer(s string) *SurveyAnswerUpdate {
	sau.mutation.SetSurveyAnswer(s)
	return sau
}

// SetNillableSurveyAnswer sets the "survey_answer" field if the given value is not nil.
func (sau *SurveyAnswerUpdate) SetNillableSurveyAnswer(s *string) *SurveyAnswerUpdate {
	if s != nil {
		sau.SetSurveyAnswer(*s)
	}
	return sau
}

// SetSurveyResponse sets the "survey_response" edge to the SurveyResponse entity.
func (sau *SurveyAnswerUpdate) SetSurveyResponse(s *SurveyResponse) *SurveyAnswerUpdate {
	return sau.SetSurveyResponseID(s.ID)
}

// SetSurveyQuestion sets the "survey_question" edge to the SurveyQuestion entity.
func (sau *SurveyAnswerUpdate) SetSurveyQuestion(s *SurveyQuestion) *SurveyAnswerUpdate {
	return sau.SetSurveyQuestionID(s.ID)
}

// Mutation returns the SurveyAnswerMutation object of the builder.
func (sau *SurveyAnswerUpdate) Mutation() *SurveyAnswerMutation {
	return sau.mutation
}

// ClearSurveyResponse clears the "survey_response" edge to the SurveyResponse entity.
func (sau *SurveyAnswerUpdate) ClearSurveyResponse() *SurveyAnswerUpdate {
	sau.mutation.ClearSurveyResponse()
	return sau
}

// ClearSurveyQuestion clears the "survey_question" edge to the SurveyQuestion entity.
func (sau *SurveyAnswerUpdate) ClearSurveyQuestion() *SurveyAnswerUpdate {
	sau.mutation.ClearSurveyQuestion()
	return sau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sau *SurveyAnswerUpdate) Save(ctx context.Context) (int, error) {
	sau.defaults()
	return withHooks(ctx, sau.sqlSave, sau.mutation, sau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sau *SurveyAnswerUpdate) SaveX(ctx context.Context) int {
	affected, err := sau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sau *SurveyAnswerUpdate) Exec(ctx context.Context) error {
	_, err := sau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sau *SurveyAnswerUpdate) ExecX(ctx context.Context) {
	if err := sau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sau *SurveyAnswerUpdate) defaults() {
	if _, ok := sau.mutation.UpdatedAt(); !ok {
		v := surveyanswer.UpdateDefaultUpdatedAt()
		sau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sau *SurveyAnswerUpdate) check() error {
	if _, ok := sau.mutation.SurveyResponseID(); sau.mutation.SurveyResponseCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "SurveyAnswer.survey_response"`)
	}
	if _, ok := sau.mutation.SurveyQuestionID(); sau.mutation.SurveyQuestionCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "SurveyAnswer.survey_question"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sau *SurveyAnswerUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SurveyAnswerUpdate {
	sau.modifiers = append(sau.modifiers, modifiers...)
	return sau
}

func (sau *SurveyAnswerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := sau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(surveyanswer.Table, surveyanswer.Columns, sqlgraph.NewFieldSpec(surveyanswer.FieldID, field.TypeInt64))
	if ps := sau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sau.mutation.CreatedBy(); ok {
		_spec.SetField(surveyanswer.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := sau.mutation.AddedCreatedBy(); ok {
		_spec.AddField(surveyanswer.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := sau.mutation.UpdatedBy(); ok {
		_spec.SetField(surveyanswer.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := sau.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(surveyanswer.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := sau.mutation.UpdatedAt(); ok {
		_spec.SetField(surveyanswer.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := sau.mutation.DeletedAt(); ok {
		_spec.SetField(surveyanswer.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := sau.mutation.SurveyAnswer(); ok {
		_spec.SetField(surveyanswer.FieldSurveyAnswer, field.TypeString, value)
	}
	if sau.mutation.SurveyResponseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyanswer.SurveyResponseTable,
			Columns: []string{surveyanswer.SurveyResponseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyresponse.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sau.mutation.SurveyResponseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyanswer.SurveyResponseTable,
			Columns: []string{surveyanswer.SurveyResponseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyresponse.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sau.mutation.SurveyQuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyanswer.SurveyQuestionTable,
			Columns: []string{surveyanswer.SurveyQuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyquestion.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sau.mutation.SurveyQuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyanswer.SurveyQuestionTable,
			Columns: []string{surveyanswer.SurveyQuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyquestion.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(sau.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, sau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{surveyanswer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sau.mutation.done = true
	return n, nil
}

// SurveyAnswerUpdateOne is the builder for updating a single SurveyAnswer entity.
type SurveyAnswerUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SurveyAnswerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (sauo *SurveyAnswerUpdateOne) SetCreatedBy(i int64) *SurveyAnswerUpdateOne {
	sauo.mutation.ResetCreatedBy()
	sauo.mutation.SetCreatedBy(i)
	return sauo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (sauo *SurveyAnswerUpdateOne) SetNillableCreatedBy(i *int64) *SurveyAnswerUpdateOne {
	if i != nil {
		sauo.SetCreatedBy(*i)
	}
	return sauo
}

// AddCreatedBy adds i to the "created_by" field.
func (sauo *SurveyAnswerUpdateOne) AddCreatedBy(i int64) *SurveyAnswerUpdateOne {
	sauo.mutation.AddCreatedBy(i)
	return sauo
}

// SetUpdatedBy sets the "updated_by" field.
func (sauo *SurveyAnswerUpdateOne) SetUpdatedBy(i int64) *SurveyAnswerUpdateOne {
	sauo.mutation.ResetUpdatedBy()
	sauo.mutation.SetUpdatedBy(i)
	return sauo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sauo *SurveyAnswerUpdateOne) SetNillableUpdatedBy(i *int64) *SurveyAnswerUpdateOne {
	if i != nil {
		sauo.SetUpdatedBy(*i)
	}
	return sauo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (sauo *SurveyAnswerUpdateOne) AddUpdatedBy(i int64) *SurveyAnswerUpdateOne {
	sauo.mutation.AddUpdatedBy(i)
	return sauo
}

// SetUpdatedAt sets the "updated_at" field.
func (sauo *SurveyAnswerUpdateOne) SetUpdatedAt(t time.Time) *SurveyAnswerUpdateOne {
	sauo.mutation.SetUpdatedAt(t)
	return sauo
}

// SetDeletedAt sets the "deleted_at" field.
func (sauo *SurveyAnswerUpdateOne) SetDeletedAt(t time.Time) *SurveyAnswerUpdateOne {
	sauo.mutation.SetDeletedAt(t)
	return sauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sauo *SurveyAnswerUpdateOne) SetNillableDeletedAt(t *time.Time) *SurveyAnswerUpdateOne {
	if t != nil {
		sauo.SetDeletedAt(*t)
	}
	return sauo
}

// SetSurveyResponseID sets the "survey_response_id" field.
func (sauo *SurveyAnswerUpdateOne) SetSurveyResponseID(i int64) *SurveyAnswerUpdateOne {
	sauo.mutation.SetSurveyResponseID(i)
	return sauo
}

// SetNillableSurveyResponseID sets the "survey_response_id" field if the given value is not nil.
func (sauo *SurveyAnswerUpdateOne) SetNillableSurveyResponseID(i *int64) *SurveyAnswerUpdateOne {
	if i != nil {
		sauo.SetSurveyResponseID(*i)
	}
	return sauo
}

// SetSurveyQuestionID sets the "survey_question_id" field.
func (sauo *SurveyAnswerUpdateOne) SetSurveyQuestionID(i int64) *SurveyAnswerUpdateOne {
	sauo.mutation.SetSurveyQuestionID(i)
	return sauo
}

// SetNillableSurveyQuestionID sets the "survey_question_id" field if the given value is not nil.
func (sauo *SurveyAnswerUpdateOne) SetNillableSurveyQuestionID(i *int64) *SurveyAnswerUpdateOne {
	if i != nil {
		sauo.SetSurveyQuestionID(*i)
	}
	return sauo
}

// SetSurveyAnswer sets the "survey_answer" field.
func (sauo *SurveyAnswerUpdateOne) SetSurveyAnswer(s string) *SurveyAnswerUpdateOne {
	sauo.mutation.SetSurveyAnswer(s)
	return sauo
}

// SetNillableSurveyAnswer sets the "survey_answer" field if the given value is not nil.
func (sauo *SurveyAnswerUpdateOne) SetNillableSurveyAnswer(s *string) *SurveyAnswerUpdateOne {
	if s != nil {
		sauo.SetSurveyAnswer(*s)
	}
	return sauo
}

// SetSurveyResponse sets the "survey_response" edge to the SurveyResponse entity.
func (sauo *SurveyAnswerUpdateOne) SetSurveyResponse(s *SurveyResponse) *SurveyAnswerUpdateOne {
	return sauo.SetSurveyResponseID(s.ID)
}

// SetSurveyQuestion sets the "survey_question" edge to the SurveyQuestion entity.
func (sauo *SurveyAnswerUpdateOne) SetSurveyQuestion(s *SurveyQuestion) *SurveyAnswerUpdateOne {
	return sauo.SetSurveyQuestionID(s.ID)
}

// Mutation returns the SurveyAnswerMutation object of the builder.
func (sauo *SurveyAnswerUpdateOne) Mutation() *SurveyAnswerMutation {
	return sauo.mutation
}

// ClearSurveyResponse clears the "survey_response" edge to the SurveyResponse entity.
func (sauo *SurveyAnswerUpdateOne) ClearSurveyResponse() *SurveyAnswerUpdateOne {
	sauo.mutation.ClearSurveyResponse()
	return sauo
}

// ClearSurveyQuestion clears the "survey_question" edge to the SurveyQuestion entity.
func (sauo *SurveyAnswerUpdateOne) ClearSurveyQuestion() *SurveyAnswerUpdateOne {
	sauo.mutation.ClearSurveyQuestion()
	return sauo
}

// Where appends a list predicates to the SurveyAnswerUpdate builder.
func (sauo *SurveyAnswerUpdateOne) Where(ps ...predicate.SurveyAnswer) *SurveyAnswerUpdateOne {
	sauo.mutation.Where(ps...)
	return sauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sauo *SurveyAnswerUpdateOne) Select(field string, fields ...string) *SurveyAnswerUpdateOne {
	sauo.fields = append([]string{field}, fields...)
	return sauo
}

// Save executes the query and returns the updated SurveyAnswer entity.
func (sauo *SurveyAnswerUpdateOne) Save(ctx context.Context) (*SurveyAnswer, error) {
	sauo.defaults()
	return withHooks(ctx, sauo.sqlSave, sauo.mutation, sauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sauo *SurveyAnswerUpdateOne) SaveX(ctx context.Context) *SurveyAnswer {
	node, err := sauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sauo *SurveyAnswerUpdateOne) Exec(ctx context.Context) error {
	_, err := sauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sauo *SurveyAnswerUpdateOne) ExecX(ctx context.Context) {
	if err := sauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sauo *SurveyAnswerUpdateOne) defaults() {
	if _, ok := sauo.mutation.UpdatedAt(); !ok {
		v := surveyanswer.UpdateDefaultUpdatedAt()
		sauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sauo *SurveyAnswerUpdateOne) check() error {
	if _, ok := sauo.mutation.SurveyResponseID(); sauo.mutation.SurveyResponseCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "SurveyAnswer.survey_response"`)
	}
	if _, ok := sauo.mutation.SurveyQuestionID(); sauo.mutation.SurveyQuestionCleared() && !ok {
		return errors.New(`cep_ent: clearing a required unique edge "SurveyAnswer.survey_question"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sauo *SurveyAnswerUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SurveyAnswerUpdateOne {
	sauo.modifiers = append(sauo.modifiers, modifiers...)
	return sauo
}

func (sauo *SurveyAnswerUpdateOne) sqlSave(ctx context.Context) (_node *SurveyAnswer, err error) {
	if err := sauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(surveyanswer.Table, surveyanswer.Columns, sqlgraph.NewFieldSpec(surveyanswer.FieldID, field.TypeInt64))
	id, ok := sauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cep_ent: missing "SurveyAnswer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, surveyanswer.FieldID)
		for _, f := range fields {
			if !surveyanswer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cep_ent: invalid field %q for query", f)}
			}
			if f != surveyanswer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sauo.mutation.CreatedBy(); ok {
		_spec.SetField(surveyanswer.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := sauo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(surveyanswer.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := sauo.mutation.UpdatedBy(); ok {
		_spec.SetField(surveyanswer.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := sauo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(surveyanswer.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := sauo.mutation.UpdatedAt(); ok {
		_spec.SetField(surveyanswer.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := sauo.mutation.DeletedAt(); ok {
		_spec.SetField(surveyanswer.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := sauo.mutation.SurveyAnswer(); ok {
		_spec.SetField(surveyanswer.FieldSurveyAnswer, field.TypeString, value)
	}
	if sauo.mutation.SurveyResponseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyanswer.SurveyResponseTable,
			Columns: []string{surveyanswer.SurveyResponseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyresponse.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sauo.mutation.SurveyResponseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyanswer.SurveyResponseTable,
			Columns: []string{surveyanswer.SurveyResponseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyresponse.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sauo.mutation.SurveyQuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyanswer.SurveyQuestionTable,
			Columns: []string{surveyanswer.SurveyQuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyquestion.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sauo.mutation.SurveyQuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surveyanswer.SurveyQuestionTable,
			Columns: []string{surveyanswer.SurveyQuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(surveyquestion.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(sauo.modifiers...)
	_node = &SurveyAnswer{config: sauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{surveyanswer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sauo.mutation.done = true
	return _node, nil
}
