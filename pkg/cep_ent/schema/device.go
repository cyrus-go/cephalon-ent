package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type Device struct {
	ent.Schema
}

func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id,omitempty"`).Default(0).Comment("外键用户 id"),
		field.String("serial_number").StructTag(`json:"serial_number"`).Default("").Comment("设备唯一序列号"),
		field.Enum("state").Values("On", "Down", "Init").Default("Init").Comment("设备状态"),
		field.Int64("sum_cep").StructTag(`json:"sum_cep"`).Default(0).Comment("该设备总获得利润"),
		field.Bool("linking").Default(false).StructTag(`json:"linking"`).Comment("设备是否正在对接中"),
		field.Enum("binding_status").GoType(enums.DeviceBindingStatusInit).Default(string(enums.DeviceBindingStatusInit)).StructTag(`json:"binding_status"`).Comment("设备的绑定状态"),
		field.Enum("status").Values("online", "offline", "busy").Default("online").StructTag(`json:"status"`).Comment("设备状态"),
	}
}

// Edges of the ServerOrder.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("devices").Field("user_id").Unique().Required(),
		edge.To("mission_produce_orders", MissionProduceOrder.Type),
		edge.To("user_devices", UserDevice.Type),
		edge.To("device_gpu_missions", DeviceGpuMission.Type),
	}
}

// Mixin of ServerOrder
func (Device) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}