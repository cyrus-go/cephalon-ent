package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ExtraServicePrice holds the schema definition for the ExtraServicePrice entity.
type ExtraServicePrice struct {
	ent.Schema
}

// Fields of the ExtraServicePrice
func (ExtraServicePrice) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("extra_service_type").GoType(enums.ExtraServiceTypeVPN).Default(string(enums.ExtraServiceTypeUnknown)).StructTag(`json:"extra_service_type"`).Comment("附加服务类型"),
		field.Enum("extra_service_billing_type").GoType(enums.ExtraServiceBillingTypeTimePlanDay).Default(string(enums.ExtraServiceBillingTypeUnknown)).StructTag(`json:"extra_service_billing_type"`).Comment("附加服务订单类型"),
		field.Int64("extra_service_id").Default(0).StructTag(`json:extra_service_id,string`).Comment("附加服务 id，外键关联附加服务"),
		field.Int64("cep").Default(0).StructTag(`json:"cep"`).Comment("附加服务单价"),
		field.Time("started_at").Optional().Nillable().StructTag(`json:"started_at"`).Comment("价格有效时间开始，为空表示永久有效"),
		field.Time("finished_at").Optional().Nillable().StructTag(`json:"finished_at"`).Comment("价格有效时间结束，为空表示永久有效"),
		field.Bool("is_deprecated").Default(false).StructTag(`json:"is_deprecated"`).Comment("价格是否屏蔽，前端置灰，硬选也可以"),
		field.Bool("is_sensitive").Default(false).StructTag(`json:"is_sensitive"`).Comment("价格是否敏感，用于特殊类型任务，不能让外部看到选项"),
	}
}

// Edges of the ExtraServicePrice
func (ExtraServicePrice) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("extra_service", ExtraService.Type).Ref("extra_service_prices").Field("extra_service_id").Unique().Required(),
	}
}

// Indexes of the ExtraServicePrice
func (ExtraServicePrice) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("extra_service_id"),
	}
}

// Mixin of ExtraServicePrice
func (ExtraServicePrice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (ExtraServicePrice) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("附加服务价格表"),
	}
}
