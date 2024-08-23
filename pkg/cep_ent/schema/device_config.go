package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type DeviceConfig struct {
	ent.Schema
}

func (DeviceConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("device_id").StructTag(`json:"device_id,string"`).Default(0).Comment("外键设备 id"),
		field.Int64("gap_base").StructTag(`json:"gap_base"`).Default(0).Comment("间隔基数"),
		field.Int64("gap_random_max").StructTag(`json:"gap_random_max"`).Default(0).Comment("间隔随机范围最大值"),
		field.Int64("gap_random_min").StructTag(`json:"gap_random_min"`).Default(0).Comment("间隔随机范围最小值"),
	}
}

// Edges of the DeviceConfig.
func (DeviceConfig) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("device", Device.Type).Ref("device_config").Field("device_id").Unique().Required(),
	}
}

// Indexes of DeviceConfig
func (DeviceConfig) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("device_id"),
	}
}

// Mixin of DeviceConfig
func (DeviceConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (DeviceConfig) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("设备配置信息"),
	}
}
