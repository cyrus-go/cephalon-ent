package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type SurveyAnswer struct {
	ent.Schema
}

func (SurveyAnswer) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("survey_response_id").StructTag(`json:"survey_response_id,string"`).Default(0).Comment("问卷用户调查结果 ID，一个用户同一个问卷只能回答一次，这个问卷会有多个问题和答案"),
		field.Int64("survey_question_id").StructTag(`json:"survey_question_id,string"`).Default(0).Comment("问题 id"),
		field.Text("survey_answer").StructTag(`json:"survey_answer"`).Default("").Comment("答案的内容"),
	}
}

// Edges of the SurveyAnswer.
func (SurveyAnswer) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("survey_response", SurveyResponse.Type).Ref("survey_answers").Field("survey_response_id").Unique().Required(),
		edge.From("survey_question", SurveyQuestion.Type).Ref("survey_answers").Field("survey_question_id").Unique().Required(),
	}
}

// Indexes of SurveyAnswer
func (SurveyAnswer) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of SurveyAnswer
func (SurveyAnswer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (SurveyAnswer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("问卷问题答案表"),
	}
}
