package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/common"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ExtraServiceOrder holds the schema definition for the ExtraServiceOrder entity.
type ExtraServiceOrder struct {
	ent.Schema
}

// Fields of the ExtraServiceOrder
func (ExtraServiceOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("mission_id").Default(0).StructTag(`json:"mission_id,string"`).Comment("任务 id，外键关联任务"),
		field.Int64("mission_order_id").Default(0).StructTag(`json:"mission_order_id,string"`).Comment("任务订单 id，外键关联任务订单"),
		field.Int64("amount").Default(0).StructTag(`json:"amount"`).Comment("订单的货币消耗量"),
		field.Int64("symbol_id").Default(0).StructTag(`json:"symbol_id,string"`).Comment("币种 id"),
		field.Int64("unit_cep").Default(0).StructTag(`json:"unit_cep"`).Comment("任务单价，按次(count)就是 unit_cep/次，按时(time)就是 unit_cep/分钟"),
		field.Enum("extra_service_type").GoType(enums.ExtraServiceTypeVPN).Default(string(enums.ExtraServiceTypeUnknown)).StructTag(`json:"extra_service_type"`).Comment("附加服务类型"),
		field.Int64("buy_duration").Default(0).StructTag(`json:"buy_duration"`).Comment("包时任务订单购买的时长"),
		field.Time("plan_started_at").Default(common.ZeroTime).Nillable().Optional().StructTag(`json:"plan_started_at"`).Comment("任务计划开始时间（包时）"),
		field.Time("plan_finished_at").Default(common.ZeroTime).Nillable().Optional().StructTag(`json:"plan_finished_at"`).Comment("任务计划结束时间（包时）"),
		field.Int64("mission_batch_id").Default(0).StructTag(`json:"mission_batch_id,string"`).Comment("任务批次外键"),
	}
}

// Edges of the ExtraServiceOrder
func (ExtraServiceOrder) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("mission", Mission.Type).Ref("extra_service_orders").Field("mission_id").Unique().Required(),
		edge.From("mission_order", MissionOrder.Type).Ref("extra_service_orders").Field("mission_order_id").Unique().Required(),
		edge.From("symbol", Symbol.Type).Ref("extra_service_order").Field("symbol_id").Unique().Required(),
		edge.From("mission_batch", MissionBatch.Type).Ref("extra_service_order").Field("mission_batch_id").Unique().Required(),
	}
}

// Mixin of ExtraServiceOrder
func (ExtraServiceOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (ExtraServiceOrder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("附加服务订单表"),
	}
}
