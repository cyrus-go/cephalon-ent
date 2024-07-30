package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Token holds the schema definition for the Token entity.
type ModelOrder struct {
	ent.Schema
}

// Fields of the Token.
func (ModelOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").StructTag(`json:"user_id"`).Comment("用户ID"),
		field.String("token").Default("").Sensitive().Comment("有token为api调用，没token为web调用"),
		field.String("model_id").Default("").StructTag(`json:"model_id"`).Comment("模型ID"),
		field.Int("input_token_cost").Default(0).StructTag(`json:"input_token_cost"`).Comment("输入token消耗"),
		field.Int("output_token_cost").Default(0).StructTag(`json:"output_token_cost"`).Comment("输出token消耗"),
		field.Int("total_cost_cep").Default(0).StructTag(`json:"total_cost_cep"`).Comment("总消耗cep"),
	}
}

// Edges of the Token.
func (ModelOrder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("model_orders").Field("user_id").Unique().Required(),
		edge.From("model", Model.Type).Ref("model_orders").Field("model_id").Unique().Required(),
	}
}

func (ModelOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
