package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/common"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type DeviceGpuMission struct {
	ent.Schema
}

func (DeviceGpuMission) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("device_id").Default(0).StructTag(`json:"device_id,string"`).Comment("外键设备 id"),
		field.Int64("gpu_id").Default(0).StructTag(`json:"gpu_id,string"`).Comment("外键 gpu id"),
		field.String("able_mission_kind").GoType([]string{}).Optional().ValueScanner(common.Bytes2StringSliceValueScanner{}).SchemaType(map[string]string{dialect.Postgres: "bytea"}).StructTag(`json:"able_mission_kind"`).Comment("可以接的任务类型"),
		field.Int8("device_slot").Default(0).StructTag(`json:"device_slot"`).Comment("显卡占用设备的插槽"),
		field.Int8("max_online_mission").Default(0).StructTag(`json:"max_online_mission"`).Comment("最大同时在线任务"),
		field.Enum("gpu_status").Default(string(enums.GpuStatusOffline)).GoType(enums.GpuStatusFree).StructTag(`json:"gpu_status"`).Comment("gpu 当前状态"),
		field.String("mission_id").GoType([]int64{}).Optional().ValueScanner(common.Bytes2Int64SliceValueScanner{}).SchemaType(map[string]string{dialect.Postgres: "bytea"}).StructTag(`json:"mission_id"`).Comment("正在做的任务 id"),
	}
}

// Edges of the DeviceGpuMission.
func (DeviceGpuMission) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("device", Device.Type).Ref("device_gpu_missions").Field("device_id").Unique().Required(),
		//	edge.From("mission_kind", MissionKind.Type).Ref("device_gpu_missions").Field("mission_kind_id").Unique().Required(),
		edge.From("gpu", Gpu.Type).Ref("device_gpu_missions").Field("gpu_id").Unique().Required(),
	}
}

func (DeviceGpuMission) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("device_id"),
		index.Fields("gpu_id"),
	}
}

// Mixin of DeviceGpuMission
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
