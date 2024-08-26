package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type GiftMissionConfig struct {
	ent.Schema
}

func (GiftMissionConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("stability_level").GoType(enums.DeviceStabilityTypeGood).Default(string(enums.DeviceStabilityTypeGood)).StructTag(`json:"stability_level"`).Comment("稳定性等级"),
		field.Enum("gpu_version").GoType(enums.GpuVersion3080).Default(string(enums.GpuVersion3080)).StructTag(`json:"gpu_version"`).Comment("GPU 版本"),
		field.Int64("gap_base").StructTag(`json:"gap_base"`).Default(0).Comment("间隔基数"),
		field.Int64("gap_random_max").StructTag(`json:"gap_random_max"`).Default(0).Comment("间隔随机范围最大值"),
		field.Int64("gap_random_min").StructTag(`json:"gap_random_min"`).Default(0).Comment("间隔随机范围最小值"),
	}
}

// Edges of the GiftMissionConfig.
func (GiftMissionConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("devices", Device.Type),
	}
}

// Indexes of GiftMissionConfig
func (GiftMissionConfig) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("stability_level", "gpu_version"),
	}
}

// Mixin of GiftMissionConfig
func (GiftMissionConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (GiftMissionConfig) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("补贴任务参数配置（需要初始化数据）"),
	}
}
