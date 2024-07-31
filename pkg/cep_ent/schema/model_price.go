package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ModelPrice holds the schema definition for the ModelPrice entity.
type ModelPrice struct {
	ent.Schema
}

// Fields of the ModelPrice.
func (ModelPrice) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("model_id").StructTag(`json:"model_id,string"`).Comment("模型ID"),
		field.Enum("invoke_type").Default(string(enums.UnKnownInvokeType)).GoType(enums.ApiInvokeType).StructTag(`json:"invoke_type"`).Comment("调用类型"),
		field.Enum("gpu_version").Default(string(enums.GpuVersionUnknown)).GoType(enums.GpuVersion3060).StructTag(`json:"gpu_version"`).Comment("GPU 版本"),
		field.Int("input_gpu_price").Default(0).StructTag(`json:"input_gpu_price"`).Comment("输入算力价格"),
		field.Int("output_gpu_price").Default(0).StructTag(`json:"output_gpu_price"`).Comment("输出算力价格"),
		field.Int("input_model_price").Default(0).StructTag(`json:"input_model_price"`).Comment("输入模型使用价格"),
		field.Int("output_model_price").Default(0).StructTag(`json:"output_model_price"`).Comment("输出模型使用价格"),
	}
}

// Edges of the ModelPrice.
func (ModelPrice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("model", Model.Type).Ref("model_prices").Field("model_id").Unique().Required(),
	}
}

func (ModelPrice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
