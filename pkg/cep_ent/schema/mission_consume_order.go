package schema

import (
	"cephalon-ent/common"
	"cephalon-ent/pkg/enums"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.Int64("mission_id").Default(0).StructTag(`json:"mission_id"`).Comment("外键任务 id，关联任务"),
		field.Enum("status").Values("waiting", "canceled", "doing", "succeed", "failed").Default("waiting").StructTag(`json:"status"`).Comment("任务订单的状态，注意不强关联任务的状态"),
		field.Int64("cep").Default(0).StructTag(`json:"cep"`).Comment("发布任务需消耗的 cep 量"),
		field.Enum("type").GoType(enums.MissionTypeSdTxt2Img).Default(string(enums.MissionTypeSdTxt2Img)).StructTag(`json:"type"`).Comment("任务类型，等于任务表的类型字段"),
		field.Bool("is_time").Default(false).StructTag(`json:"is_time"`).Comment("是否为计时类型任务"),
		field.Enum("call_way").Values("api", "yuan_hui", "dev_platform").Default("api").StructTag(`json:"call_way"`).Comment("调用方式，API 调用或者微信小程序调用"),
		field.String("serial_number").Default("").StructTag(`json:"serial_number"`).Comment("订单序列号"),
		field.Time("started_at").Default(common.ZeroTime).Optional().Nillable().StructTag(`json:"started_at"`).Comment("任务开始执行时刻"),
		field.Time("finished_at").Default(common.ZeroTime).Optional().Nillable().StructTag(`json:"finished_at"`).Comment("任务结束执行时刻"),
		field.Int64("mission_batch_id").Default(0).StructTag(`json:"mission_batch_id"`).Comment("外键任务批次 id"),
		field.String("mission_batch_number").Default("").StructTag(`json:"mission_batch_number"`).Comment("任务批次号，用于方便检索"),
	}
}

// Edges of the MissionConsumeOrder.
func (MissionConsumeOrder) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("mission_consume_orders").Field("user_id").Unique().Required(),
		edge.To("bills", Bill.Type),
		// 和任务一对一
		edge.From("mission", Mission.Type).Ref("mission_consume_order").Unique().Required(),
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

// Annotations of the BaseMixin.
func (MissionConsumeOrder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务消费订单，和任务一对一"),
	}
}
