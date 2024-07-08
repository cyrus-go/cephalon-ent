package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/common"
)

type Survey struct {
	ent.Schema
}

func (Survey) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").StructTag(`json:"title"`).Default("").Comment("标题"),
		field.Time("started_at").Default(common.ZeroTime).Nillable().Optional().StructTag(`json:"started_at"`).Comment("填写问卷开始的时间"),
		field.Time("ended_at").Default(common.ZeroTime).Nillable().Optional().StructTag(`json:"ended_at"`).Comment("填写问卷结束的时间"),
		field.Int64("sort_num").StructTag(`json:"sort_num"`).Default(1).Comment("分组排序序列号"),
		field.String("group").StructTag(`json:"group"`).Default("").Comment("问卷分组（自定义，可以为空），同组问卷可以根据序号强关联"),
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