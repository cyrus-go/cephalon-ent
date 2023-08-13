package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Device struct {
	ent.Schema
}

func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键用户 id"),
		field.Enum("status").Values("online", "busy", "free", "offline").Default("online").Comment("设备状态"),
		field.Enum("binding_status").Values("init", "bound", "unbound", "rebinding").Default("init").StructTag(`json:"binding_status"`).Comment("设备的绑定状态"),
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
