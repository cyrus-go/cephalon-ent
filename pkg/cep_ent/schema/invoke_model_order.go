package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type InvokeModelOrder struct {
	ent.Schema
}

func (InvokeModelOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Comment("用户ID"),
		field.Int64("model_id").StructTag(`json:"model_id"`).Comment("模型ID"),
		field.Int64("api_token_id").StructTag(`json:"api_token"`).Comment("API Token ID"),
		field.Enum("invoke_type").StructTag(`json:"invoke_type"`).Default(string(enums.UnKnownInvokeType)).GoType(enums.WebInvokeType).Comment("调用类型"),
		field.Int("invoke_times").StructTag(`json:"invoke_times"`).Default(0).Comment("调用次数"),
		field.Int("input_token_cost").StructTag(`json:"input_token_cost"`).Default(0).Comment("输入token消耗"),
		field.Int("output_token_cost").StructTag(`json:"output_token_cost"`).Default(0).Comment("输出token消耗"),
		field.Int("input_cep_cost").StructTag(`json:"input_cep_cost"`).Default(0).Comment("输入cep消耗"),
		field.Int("output_cep_cost").StructTag(`json:"output_cep_cost"`).Default(0).Comment("输出cep消耗"),
	}
}

func (InvokeModelOrder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bills", Bill.Type),
		edge.From("user", User.Type).Ref("invoke_model_orders").Field("user_id").Unique().Required(),
		edge.From("model", Model.Type).Ref("invoke_model_orders").Field("model_id").Unique().Required(),
		edge.From("api_token", ApiToken.Type).Ref("invoke_model_orders").Field("api_token_id").Unique().Required(),
	}
}

func (InvokeModelOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
