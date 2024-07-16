package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type DeviceOfflineRecord struct {
	ent.Schema
}

func (DeviceOfflineRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("device_id").Default(0).StructTag(`json:"device_id,string"`).Comment("外键设备 id"),
	}
}

// Edges of the DeviceState.
func (DeviceOfflineRecord) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("device", Device.Type).Ref("device_offline_records").Field("device_id").Unique().Required(),
	}
}

func (DeviceOfflineRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("device_id"),
	}
}

// Mixin of DeviceState
func (DeviceOfflineRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (DeviceOfflineRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("设备离线时间记录"),
	}
}
