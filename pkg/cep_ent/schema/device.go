package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/common"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type Device struct {
	ent.Schema
}

func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id,omitempty,string"`).Default(0).Comment("外键用户 id"),
		field.String("serial_number").StructTag(`json:"serial_number"`).Default("").Comment("设备唯一序列号"),
		field.Enum("state").Values("On", "Down", "Init").Default("Init").Comment("设备状态"),
		field.Int64("sum_cep").StructTag(`json:"sum_cep"`).Default(0).Comment("该设备总获得利润"),
		field.Bool("linking").Default(false).StructTag(`json:"linking"`).Comment("设备是否正在对接中"),
		field.Enum("binding_status").GoType(enums.DeviceBindingStatusInit).Default(string(enums.DeviceBindingStatusInit)).StructTag(`json:"binding_status"`).Comment("设备的绑定状态"),
		field.Enum("status").Values("online", "offline", "busy").Default("online").StructTag(`json:"status"`).Comment("设备状态"),
		field.String("name").StructTag(`json:"name"`).Default("").Comment("设备名称"),
		field.Enum("type").GoType(enums.DeviceTypeOrdinary).Default(string(enums.DeviceTypeOrdinary)).StructTag(`json:"type"`).Comment("设备类型"),
		field.Int64("cores_number").StructTag(`json:"cores_number"`).Default(0).Comment("核心数"),
		field.String("cpu").StructTag(`json:cpu`).Default("").Comment("CPU型号"),
		field.String("cpus").GoType([]string{}).Optional().ValueScanner(common.Bytes2StringSliceValueScanner{}).SchemaType(map[string]string{dialect.Postgres: "bytea"}).StructTag(`json:"cpus"`).Comment("CPU型号"),
		field.Int64("memory").StructTag(`json:"memory"`).Default(0).Comment("内存(单位:G)"),
		field.Int64("disk").StructTag(`json:disk`).Default(0).Comment("硬盘(单位:T)"),
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
		edge.To("frpc_infos", FrpcInfo.Type),
		edge.To("mission_orders", MissionOrder.Type),
	}
}

// Mixin of ServerOrder
func (Device) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Device) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("设备，对应客户端 core，记录心跳等信息"),
	}
}
