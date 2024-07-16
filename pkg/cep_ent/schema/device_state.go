package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type DeviceState struct {
	ent.Schema
}

func (DeviceState) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("device_id").Default(0).StructTag(`json:"device_id,string"`).Comment("外键设备 id"),
		field.Float("delay").Default(0).StructTag(`json:"delay"`).Comment("延迟(单位:ms)"),
	}
}

// Edges of the DeviceState.
func (DeviceState) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("device", Device.Type).Ref("device_states").Field("device_id").Unique().Required(),
	}
}

func (DeviceState) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("device_id"),
	}
}

// Mixin of DeviceState
func (DeviceState) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (DeviceState) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("设备状态信息 目前只有延迟信息"),
	}
}
