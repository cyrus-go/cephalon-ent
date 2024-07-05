package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/common"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type SurveyQuestion struct {
	ent.Schema
}

func (SurveyQuestion) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("survey_id").StructTag(`json:"survey_id,string"`).Default(0).Comment("问卷 ID"),
		field.Text("text").StructTag(`json:"text"`).Default("").Comment("问题的内容"),
		field.Enum("type").GoType(enums.SurveyQuestionTypeSingle).Default(string(enums.SurveyQuestionTypeUnknown)).StructTag(`json:"type"`).Comment("问题类型，单选/多选等"),
		field.String("options").GoType([]string{}).Optional().ValueScanner(common.Bytes2StringSliceValueScanner{}).SchemaType(map[string]string{dialect.Postgres: "bytea"}).StructTag(`json:"options"`).Comment("选项内容，单选/多选的内容"),
	}
}

// Edges of the SurveyQuestion.
func (SurveyQuestion) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("survey", Survey.Type).Ref("survey_questions").Field("survey_id").Unique().Required(),

		edge.To("survey_answers", SurveyAnswer.Type),
	}
}

// Indexes of SurveyQuestion
func (SurveyQuestion) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of SurveyQuestion
func (SurveyQuestion) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (SurveyQuestion) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("问卷问题表"),
	}
}
