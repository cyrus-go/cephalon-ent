package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// Model holds the schema definition for the Model entity.
type Model struct {
	ent.Schema
}

// Fields of the LanguageModel.
func (Model) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").StructTag(`json:"name"`).Comment("模型名称"),
		field.String("author").Default("").StructTag(`json:"author"`).Comment("模型作者"),
		field.String("description").Default("").StructTag(`json:"description"`).Comment("模型描述"),
		field.Enum("model_type").Default(string(enums.UnknownModel)).GoType(enums.LanguageModel).StructTag(`json:"model_type"`).Comment("模型类型"),
		field.Enum("model_status").Default(string(enums.InitModelStatus)).GoType(enums.InitModelStatus).StructTag(`json:"model_status"`).Comment("模型状态"),
		field.Bool("is_official").Default(false).StructTag(`json:"is_official"`).Comment("是否为官方模型"),
		field.Int("total_usage").Default(0).StructTag(`json:"total_usage"`).Comment("模型的总使用次数"),
	}
}

// Edges of the LanguageModel.
func (Model) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("model_prices", ModelPrice.Type),
		edge.From("star_user", User.Type).Ref("star_model").Through("star_model", ModleStar.Type),
		// edge.To("model_orders", ModelOrder.Type).StructTag(`json:"model_orders"`),
	}
}

func (Model) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
