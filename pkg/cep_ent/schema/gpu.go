package schema

import (
	"entgo.io/ent"
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
	}
}

// Edges of the ServerOrder.
func (Gpu) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.To("device_gpu_missions", DeviceGpuMission.Type),
	}
}

// Mixin of ServerOrder
func (Gpu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
