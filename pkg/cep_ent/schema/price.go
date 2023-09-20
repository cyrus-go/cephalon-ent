package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.Time("started_at").Optional().Nillable().StructTag(`json:"started_at"`).Comment("价格有效时间开始，为空表示永久有效"),
		field.Time("finished_at").Optional().Nillable().StructTag(`json:"finished_at"`).Comment("价格有效时间结束，为空表示永久有效"),
		field.Bool("is_deprecated").Default(false).StructTag(`json:"is_deprecated"`).Comment("价格是否屏蔽，前端置灰，硬选也可以"),
		field.Bool("is_sensitive").Default(false).StructTag(`json:"is_sensitive"`).Comment("价格是否敏感，用于特殊类型任务，不能让外部看到选项"),
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

func (Price) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务定价表，表里有数据，任务才有单价，才可以被创建"),
	}
}
