package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type DeviceGpuMission struct {
	ent.Schema
}

func (DeviceGpuMission) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("device_id").Default(0).StructTag(`json:"device_id,string"`).Comment("外键设备 id"),
		field.Int64("gpu_id").Default(0).StructTag(`json:"gpu_id,string"`).Comment("外键 gpu id"),
		field.Int64("mission_kind_id").Default(0).StructTag(`json:"mission_kind_id,string"`).Comment("外键任务种类 id"),
		field.Int8("device_slot").Default(0).StructTag(`json:"device_slot"`).Comment("显卡占用设备的插槽"),
		field.Enum("gpu_status").Default(string(enums.GpuStatusOffline)).GoType(enums.GpuStatusFree).StructTag(`json:"gpu_status"`).Comment("gpu 当前状态"),
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

func (DeviceGpuMission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("登记设备的显卡信息，以及设备的任务执行能力配置状态"),
	}
}
