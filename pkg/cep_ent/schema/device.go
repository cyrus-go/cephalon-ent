package schema

import (
	"cephalon-ent/pkg/enums"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Device struct {
	ent.Schema
}

func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键用户 id"),
		field.Enum("status").GoType(enums.DeviceStatusOnline).Default(string(enums.DeviceStatusOnline)).Comment("设备状态"),
		field.Enum("binding_status").GoType(enums.DeviceBindingStatusInit).Default(string(enums.DeviceBindingStatusInit)).StructTag(`json:"binding_status"`).Comment("设备的绑定状态"),
	}
}

// Edges of the ServerOrder.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("devices").Field("user_id").Unique().Required(),
		edge.To("mission_produce_orders", MissionProduceOrder.Type),
		edge.To("mission_productions", MissionProduction.Type),
		edge.To("user_devices", UserDevice.Type),
	}
}

// Mixin of ServerOrder
func (Device) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of the BaseMixin.
func (Device) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("设备表，与用户一对多"),
	}
}
