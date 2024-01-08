package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type Price struct {
	ent.Schema
}

// Fields of Price.
func (Price) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("gpu_id").Default(0).StructTag(`json:"gpu_id,string"`).Comment("外键 gpu id"),
		field.Enum("gpu_version").GoType(enums.GpuVersion2060).Default(string(enums.GpuVersion2060)).StructTag(`json:"gpu_version"`).Comment("显卡型号"),
		field.Enum("mission_type").GoType(enums.MissionTypeSdTxt2Img).Default(string(enums.MissionTypeSdTxt2Img)).StructTag(`json:"mission_type"`).Comment("任务类型"),
		field.Enum("mission_category").GoType(enums.MissionCategorySD).Default(string(enums.MissionCategorySD)).StructTag(`json:"mission_category"`).Comment("任务大类"),
		field.Enum("mission_billing_type").GoType(enums.MissionBillingTypeCount).Default(string(enums.MissionBillingTypeCount)).StructTag(`json:"mission_billing_type"`).Comment("任务计费类型"),
		//field.Enum("renewal_type").GoType(enums.RenewalTypeHour).Default(string(enums.RenewalTypeUnknow)).StructTag("renewal_type").Comment("包时类型，只有包时任务才有"),
		field.Int64("cep").Default(0).StructTag(`json:"cep"`).Comment("任务单价"),
		field.Time("started_at").Optional().Nillable().StructTag(`json:"started_at"`).Comment("价格有效时间开始，为空表示永久有效"),
		field.Time("finished_at").Optional().Nillable().StructTag(`json:"finished_at"`).Comment("价格有效时间结束，为空表示永久有效"),
		field.Bool("is_deprecated").Default(false).StructTag(`json:"is_deprecated"`).Comment("价格是否屏蔽，前端置灰，硬选也可以"),
		field.Bool("is_sensitive").Default(false).StructTag(`json:"is_sensitive"`).Comment("价格是否敏感，用于特殊类型任务，不能让外部看到选项"),
		field.Bool("is_hot_gpu").Default(false).StructTag(`json:"is_hot_gpu"`).Comment("是否为热点 Gpu"),
	}
}

// Edges of the Price.
func (Price) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("gpu", Gpu.Type).Ref("prices").Field("gpu_id").Unique().Required(),
	}
}

// Indexes of the Price.
func (Price) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("gpu_id"),
	}
}

// Mixin of Price.
func (Price) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of Price.
func (Price) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务定价表，表里有数据，任务才有单价，才可以被创建"),
	}
}
