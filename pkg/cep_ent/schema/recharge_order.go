package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RechargeOrder holds the schema definition for the RechargeOrder entity.
type RechargeOrder struct {
	ent.Schema
}

// Fields of the RechargeOrder.
func (RechargeOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("充值的用户 id"),
		field.Enum("status").Values("pending", "canceled", "succeed", "failed").Default("pending").StructTag(`json:"status"`).Comment("充值订单的状态，比如微信发起支付后可能没完成支付"),
		field.Int64("pure_cep").Default(0).StructTag(`json:"pure_cep"`).Comment("充值多少本金"),
		field.Int64("gift_cep").Default(0).StructTag(`json:"gift_cep"`).Comment("赠金"),
		field.Int64("social_id").Default(0).Optional().StructTag(`json:"social_id"`).Comment("关联充值来源的身份源 id"),
		field.Enum("type").Values("vx", "alipay", "manual").Default("vx").StructTag(`json:"type"`).Comment("充值订单的类型"),
		field.String("serial_number").Default("").StructTag(`json:"serial_number"`).Comment("充值订单的序列号"),
		field.String("third_api_resp").Default("").StructTag(`json:"third_api_resp"`).Comment("第三方平台的返回，给到前端才能发起支付"),
		field.Int64("from_user_id").Default(0).StructTag(`json:"from_user_id"`).Comment("由谁发起的充值"),
		field.String("out_transaction_id").Default("").StructTag(`json:"out_transaction_id"`).Comment("平台方订单号"),
	}
}

// Edges of the RechargeOrder.
func (RechargeOrder) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("recharge_orders").Field("user_id").Unique().Required(),
		edge.To("cost_bills", CostBill.Type),
		edge.From("vx_social", VXSocial.Type).Ref("recharge_orders").Field("social_id").Unique(),
	}
}

// Mixin of RechargeOrder
func (RechargeOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
