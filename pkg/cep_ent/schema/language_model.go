package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// LanguageModel holds the schema definition for the LanguageModel entity.
type LanguageModel struct {
	ent.Schema
}

// Fields of the LanguageModel.
func (LanguageModel) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").StructTag(`json:"name"`).Comment("语言模型名称"),
		field.String("author").Default("").StructTag(`json:"author"`).Comment("语言模型作者"),
		field.Bool("is_official").Default(false).StructTag(`json:"is_official"`).Comment("是否为官方语言模型"),
		field.Int("star").Default(0).StructTag(`json:"star"`).Comment("模型收藏人数"),
		field.Int("total_usage").Default(0).StructTag(`json:"total_usage"`).Comment("语言模型的总使用次数"),
	}
}

// Edges of the LanguageModel.
func (LanguageModel) Edges() []ent.Edge {
	return nil
}
