package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/common"
)

type DeviceRebootTime struct {
	ent.Schema
}

func (DeviceRebootTime) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("device_id").Default(0).StructTag(`json:"device_id,string"`).Comment("外键设备 id"),
		field.Time("start_time").Default(common.ZeroTime).StructTag(`json:"start_time"`).Comment("设备开机时间"),
		field.Time("end_time").Default(common.ZeroTime).StructTag(`json:"end_time"`).Comment("设备关机时间"),
		field.Time("now_time").Default(common.ZeroTime).StructTag(`json:"now_time"`).Comment("设备上线时间"),
		field.String("online_time").Default("None").StructTag(`json:"online_time"`).Comment("设备运行时间"),
		field.String("offline_time").Default("None").StructTag(`json:"offline_time"`).Comment("设备宕机时间"),
	}
}

// Edges of the DeviceGpuMission.
func (DeviceRebootTime) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("device", Device.Type).Ref("device_reboot_times").Field("device_id").Unique().Required(),
	}
}

func (DeviceRebootTime) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("device_id"),
	}
}

// Mixin of DeviceGpuMission
func (DeviceRebootTime) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (DeviceRebootTime) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("设备重启时间记录，已弃用"),
	}
}
