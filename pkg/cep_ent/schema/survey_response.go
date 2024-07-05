package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type SurveyResponse struct {
	ent.Schema
}

func (SurveyResponse) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id,string"`).Default(0).Comment("用户 ID"),
		field.Int64("survey_id").StructTag(`json:"survey_id,string"`).Default(0).Comment("问卷 ID"),
	}
}

// Edges of the SurveyResponse.
func (SurveyResponse) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("survey_responses").Field("user_id").Unique().Required(),
		edge.From("survey", Survey.Type).Ref("survey_responses").Field("survey_id").Unique().Required(),

		edge.To("survey_answers", SurveyAnswer.Type),
	}
}

// Indexes of SurveyResponse
func (SurveyResponse) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of SurveyResponse
func (SurveyResponse) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (SurveyResponse) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("问卷调查结果储存表"),
	}
}
