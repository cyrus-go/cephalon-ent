package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type Gpu struct {
	ent.Schema
}

func (Gpu) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("version").GoType(enums.GpuVersion2060).Default(string(enums.GpuVersion2060)).StructTag(`json:"version"`).Comment("显卡型号"),
		field.Int("power").Default(0).StructTag(`json:"power"`).Comment("显卡能力值"),
		field.Int("video_memory").Default(40).StructTag(`json:"video_memory"`).Comment("显存"),
		field.Int("cpu").Default(12).StructTag(`json:"cpu"`).Comment("CPU"),
		field.Int("memory").Default(128).StructTag(`json:"memory"`).Comment("内存"),
		field.Int64("lowest_earn_month").Default(0).StructTag(`json:"lowest_earn_month"`).Comment("保底最低月收益"),
		field.Int64("highest_earn_month").Default(0).StructTag(`json:"highest_earn_month"`).Comment("保底最高月收益"),
		field.Int64("trouble_deduct_amount").Default(0).StructTag(`json:"trouble_deduct_amount"`).Comment("故障扣费金额，单位：厘/小时"),
	}
}

// Edges of the Gpu.
func (Gpu) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.To("device_gpu_missions", DeviceGpuMission.Type),
		edge.To("prices", Price.Type),
	}
}

// Mixin of Gpu
func (Gpu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Gpu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("显卡信息，在表里有显卡型号的，才能在系统中选择使用等"),
	}
}
