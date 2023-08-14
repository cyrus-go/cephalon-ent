package schema

import (
	"cephalon-ent/pkg/enums"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Bill struct {
	ent.Schema
}

// Fields of the StoreHouse.
func (Bill) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").GoType(enums.BillTypeMissionConsume).Default(string(enums.BillTypeMissionConsume)).StructTag(`json:"type"`).Comment("次级账单流水的类型，充值或者任务消耗或任务收益"),
		field.Bool("is_add").Default(false).StructTag(`json:"is_add"`).Comment("是否增加余额，布尔值默认为否"),
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键用户 id"),
		field.String("serial_number").Default("").StructTag(`json:"serial_number"`).Comment("账单序列号"),
		field.Int64("wallet_id").Default(0).StructTag(`json:"wallet_id"`).Comment("外键钱包 id"),
		field.Int64("cep").Default(0).StructTag(`json:"cep"`).Comment("消费或收益多少 cep"),
		field.Int64("reason_id").Default(0).Optional().StructTag(`json:"reason_id"`).Comment("关联消耗产生的来源外键 id，比如 type 为 mission 时关联任务订单"),
		field.Enum("status").GoType(enums.BillStatusPending).Default(string(enums.BillStatusPending)).StructTag(`json:"status"`).Comment("消耗流水一开始不能直接生效，确定关联的消耗时间成功后才可以扣费"),
		field.Int64("market_bill_id").Default(0).StructTag(`json:"market_bill_id"`).Comment("营销流水 id"),
		field.Int64("platform_wallet_id").Default(0).StructTag(`json:"-"`).Comment("平台分润钱包 id"),
		field.Int64("platform_cep").Default(0).StructTag(`json:"-"`).Comment("平台收取多少本金余额"),
	}
}

// Edges of the StoreHouse.
func (Bill) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("bills").Field("user_id").Unique().Required(),
		edge.From("wallet", Wallet.Type).Ref("bills").Field("wallet_id").Unique().Required(),
		edge.From("platform_wallet", PlatformWallet.Type).Ref("bills").Field("platform_wallet_id").Unique().Required(),
		edge.From("recharge_order", RechargeOrder.Type).Ref("bills").Field("reason_id").Unique(),
		edge.From("mission_consume_order", MissionConsumeOrder.Type).Ref("bills").Field("reason_id").Unique(),
		edge.From("mission_produce_order", MissionProduceOrder.Type).Ref("bills").Field("reason_id").Unique(),
	}
}

// Mixin of StoreHouse
func (Bill) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Bill) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("记录额度账户的变动的主流水账单记录"),
	}
}
