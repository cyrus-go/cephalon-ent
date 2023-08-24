package schema

import (
	"cephalon-ent/common"
	"cephalon-ent/pkg/enums"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MissionConsumeOrder holds the schema definition for the MissionConsumeOrder entity.
type MissionConsumeOrder struct {
	ent.Schema
}

// Fields of the MissionConsumeOrder.
func (MissionConsumeOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键关联用户 id"),
		field.Int64("mission_id").Default(0).StructTag(`json:"mission_id"`).Comment("任务 id，关联任务中枢的任务"),
		field.Enum("status").GoType(enums.MissionOrderStatusWaiting).Default(string(enums.MissionOrderStatusWaiting)).StructTag(`json:"status"`).Comment("任务订单的状态，注意不强关联任务的状态"),
		field.Int64("pure_cep").Default(0).StructTag(`json:"pure_cep"`).Comment("任务消耗的本金 cep 量"),
		field.Int64("gift_cep").Default(0).StructTag(`json:"gift_cep"`).Comment("任务消耗的赠送 cep 量"),
		field.Enum("type").GoType(enums.MissionTypeSdTxt2Img).Default(string(enums.MissionTypeSdTxt2Img)).StructTag(`json:"type"`).Comment("任务类型，等于任务表的类型字段"),
		field.Bool("is_time").Default(false).StructTag(`json:"is_time"`).Comment("是否为计时类型任务"),
		field.Enum("call_way").GoType(enums.MissionCallWayApi).Default(string(enums.MissionCallWayApi)).StructTag(`json:"call_way"`).Comment("调用方式，API 调用或者微信小程序调用"),
		field.String("serial_number").Default("").StructTag(`json:"serial_number"`).Comment("订单序列号"),
		field.Time("started_at").Default(common.ZeroTime).StructTag(`json:"started_at"`).Comment("任务开始执行时刻"),
		field.Time("finished_at").Default(common.ZeroTime).StructTag(`json:"finished_at"`).Comment("任务结束执行时刻"),
		field.Int64("mission_batch_id").Default(0).StructTag(`json:"mission_batch_id"`).Comment("任务批次外键"),
		field.String("mission_batch_number").Default("").StructTag(`json:"mission_batch_number"`).Comment("任务批次号，用于方便检索"),
	}
}

// Edges of the MissionConsumeOrder.
func (MissionConsumeOrder) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("mission_consume_orders").Field("user_id").Unique().Required(),
		edge.To("cost_bills", CostBill.Type),
		edge.To("mission_produce_orders", MissionProduceOrder.Type),
		edge.From("mission_batch", MissionBatch.Type).Ref("mission_consume_orders").Field("mission_batch_id").Unique().Required(),
	}
}

// Mixin of MissionConsumeOrder
func (MissionConsumeOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
