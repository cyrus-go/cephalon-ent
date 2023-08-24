package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type MissionKind struct {
	ent.Schema
}

func (MissionKind) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").GoType(enums.MissionTypeSdTxt2Img).Default(string(enums.MissionTypeSdTxt2Img)).StructTag(`json:"type"`).Comment("任务单位类型"),
		field.Enum("category").GoType(enums.MissionCategorySD).Default(string(enums.MissionCategorySD)).StructTag(`json:"category"`).Comment("任务大类"),
		field.Enum("billing_type").GoType(enums.MissionBillingTypeCount).Default(string(enums.MissionBillingTypeCount)).StructTag(`json:"billing_type"`).Comment("计费类型"),
	}
}

// Edges of the ServerOrder.
func (MissionKind) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.To("device_gpu_missions", DeviceGpuMission.Type),
	}
}

// Mixin of ServerOrder
func (MissionKind) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
