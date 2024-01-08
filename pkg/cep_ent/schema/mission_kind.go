package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type MissionKind struct {
	ent.Schema
}

func (MissionKind) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").GoType(enums.MissionTypeSdTxt2Img).Default(string(enums.MissionTypeUnknown)).StructTag(`json:"type"`).Comment("任务单位类型"),
		field.Enum("category").GoType(enums.MissionCategorySD).Default(string(enums.MissionCategoryUnknown)).StructTag(`json:"category"`).Comment("任务大类"),
		field.Enum("billing_type").GoType(enums.MissionBillingTypeCount).Default(string(enums.MissionBillingTypeUnknown)).StructTag(`json:"billing_type"`).Comment("计费类型"),
	}
}

// Edges of the MissionKind.
func (MissionKind) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		//	edge.To("device_gpu_missions", DeviceGpuMission.Type),
		edge.To("missions", Mission.Type),
	}
}

// Mixin of MissionKind
func (MissionKind) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (MissionKind) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务种类，任务类型的抽象层，记录了任务计费类型等信息"),
	}
}
