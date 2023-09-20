package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type EarnBill struct {
	ent.Schema
}

// Fields of the StoreHouse.
func (EarnBill) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values("mission", "withdraw").Default("mission").StructTag(`json:"type"`).Comment("分润账户变动类型"),
		field.Bool("is_minus").Default(false).StructTag(`json:"is_minus"`).Comment("是否减少分润钱包余额，默认为否"),
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键用户 id"),
		field.String("serial_number").Default("").StructTag(`json:"serial_number"`).Comment("账单序列号"),
		field.Int64("profit_account_id").Default(0).StructTag(`json:"profit_account_id"`).Comment("外键分润账户 id"),
		field.Int64("pure_cep").Default(0).StructTag(`json:"pure_cep"`).Comment("分润多少本金余额"),
		field.Int64("gift_cep").Default(0).StructTag(`json:"gift_cep"`).Comment("分润多少赠送余额"),
		field.Int64("platform_account_id").Default(0).StructTag(`json:"-"`).Comment("平台分润账户 id"),
		field.Int64("platform_pure_cep").Default(0).StructTag(`json:"-"`).Comment("平台收取多少本金余额"),
		field.Int64("platform_gift_cep").Default(0).StructTag(`json:"-"`).Comment("平台收取多少赠送余额"),
		field.Int64("reason_id").Default(0).Optional().StructTag(`json:"reason_id"`).Comment("关联分润产生的来源外键 id，比如 type 为 mission 时关联任务订单"),
	}
}

// Edges of the StoreHouse.
func (EarnBill) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("earn_bills").Field("user_id").Unique().Required(),
		edge.From("profit_account", ProfitAccount.Type).Ref("earn_bills").Field("profit_account_id").Unique().Required(),
		edge.From("platform_account", PlatformAccount.Type).Ref("earn_bills").Field("platform_account_id").Unique().Required(),
		edge.From("mission_produce_orders", MissionProduceOrder.Type).Ref("earn_bills").Field("reason_id").Unique(),
	}
}

// Mixin of StoreHouse
func (EarnBill) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (EarnBill) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("分润流水，被 bill 取代"),
	}
}
