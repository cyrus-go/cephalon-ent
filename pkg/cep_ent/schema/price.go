package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type Price struct {
	ent.Schema
}

func (Price) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("gpu_version").GoType(enums.GpuVersion2060).Default(string(enums.GpuVersion2060)).StructTag(`json:"gpu_version"`).Comment("显卡型号"),
		field.Enum("mission_type").GoType(enums.MissionTypeSdTxt2Img).Default(string(enums.MissionTypeSdTxt2Img)).StructTag(`json:"mission_type"`).Comment("任务类型"),
		field.Enum("mission_category").GoType(enums.MissionCategorySD).Default(string(enums.MissionCategorySD)).StructTag(`json:"mission_category"`).Comment("任务大类"),
		field.Enum("mission_billing_type").GoType(enums.MissionBillingTypeCount).Default(string(enums.MissionBillingTypeCount)).StructTag(`json:"mission_billing_type"`).Comment("任务计费类型"),
		field.Int64("cep").Default(0).StructTag(`json:"cep"`).Comment("任务单价"),
	}
}

// Edges of the ServerOrder.
func (Price) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
	}
}

// Mixin of ServerOrder
func (Price) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
