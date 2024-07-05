package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Survey struct {
	ent.Schema
}

func (Survey) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").StructTag(`json:"title"`).Default("").Comment("标题"),
	}
}

// Edges of the Survey.
func (Survey) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.To("survey_questions", SurveyQuestion.Type),
		edge.To("survey_responses", SurveyResponse.Type),
	}
}

// Indexes of Survey
func (Survey) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of Survey
func (Survey) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Survey) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("问卷表"),
	}
}
