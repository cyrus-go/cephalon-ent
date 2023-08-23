package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DeviceGpuMission struct {
	ent.Schema
}

func (DeviceGpuMission) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("device_id").Default(0).StructTag(`json:"device_id"`).Comment("外键设备 id"),
		field.Int64("gpu_id").Default(0).StructTag(`json:"gpu_id"`).Comment("外键 gpu id"),
		field.Int64("mission_kind_id").Default(0).StructTag(`json:"mission_kind_id"`).Comment("外键任务种类 id"),
	}
}

// Edges of the ServerOrder.
func (DeviceGpuMission) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("device", Device.Type).Ref("device_gpu_missions").Field("device_id").Unique().Required(),
		edge.From("mission_kind", MissionKind.Type).Ref("device_gpu_missions").Field("mission_kind_id").Unique().Required(),
		edge.From("gpu", Gpu.Type).Ref("device_gpu_missions").Field("gpu_id").Unique().Required(),
	}
}

// Mixin of ServerOrder
func (DeviceGpuMission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
