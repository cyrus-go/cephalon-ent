package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type CostBill struct {
	ent.Schema
}

// Fields of the StoreHouse.
func (CostBill) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values("unknow", "mission", "recharge", "active").Default("unknow").StructTag(`json:"type"`).Comment("额度账户流水的类型，充值或者任务消耗"),
		field.Enum("way").Values("unknow", "recharge_wechat", "recharge_alipay", "recharge_manual", "mission_timing", "mission_counting", "active_register", "active_share", "active_recharge", "first_invite_recharge").Default("unknow").StructTag(`json:"way"`).Comment("额度账户流水的产生方式，微信、支付宝、计时消耗等"),
		field.Bool("is_add").Default(false).StructTag(`json:"is_add"`).Comment("是否增加余额，布尔值默认为否"),
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键用户 id"),
		field.String("serial_number").Default("").StructTag(`json:"serial_number"`).Comment("账单序列号"),
		field.Int64("cost_account_id").Default(0).StructTag(`json:"cost_account_id"`).Comment("外键消耗账户 id"),
		field.Int64("pure_cep").Default(0).StructTag(`json:"pure_cep"`).Comment("消耗多少本金余额"),
		field.Int64("gift_cep").Default(0).StructTag(`json:"gift_cep"`).Comment("消耗多少赠送余额"),
		field.Int64("reason_id").Default(0).Optional().StructTag(`json:"reason_id"`).Comment("关联消耗产生的来源外键 id，比如 type 为 mission 时关联任务订单"),
		field.Enum("status").GoType(enums.BillStatusPending).Default(string(enums.BillStatusPending)).StructTag(`json:"status"`).Comment("消耗流水一开始不能直接生效，确定关联的消耗时间成功后才可以扣费"),

		field.Int64("market_account_id").Default(0).StructTag(`json:"market_account_id"`).Comment("营销账户 id，表示这是一条营销流水（此方案为临时方案）"),
		field.Int64("campaign_order_id").Default(0).StructTag(`json:"campaign_order_id"`).Comment("活动订单 id"),
	}
}

// Edges of the StoreHouse.
func (CostBill) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("cost_bills").Field("user_id").Unique().Required(),
		edge.From("cost_account", CostAccount.Type).Ref("cost_bills").Field("cost_account_id").Unique().Required(),
		edge.From("recharge_order", RechargeOrder.Type).Ref("cost_bills").Field("reason_id").Unique(),
		edge.From("mission_consume_order", MissionConsumeOrder.Type).Ref("cost_bills").Field("reason_id").Unique(),
		edge.From("platform_account", PlatformAccount.Type).Ref("cost_bills").Field("market_account_id").Unique().Required(),
		edge.From("campaign_order", CampaignOrder.Type).Ref("cost_bills").Field("campaign_order_id").Unique().Required(),
	}
}

// Mixin of StoreHouse
func (CostBill) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (CostBill) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("记录额度账户的变动的主流水账单记录"),
	}
}
