package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CampaignOrder holds the schema definition for the CampaignOrder entity.
type CampaignOrder struct {
	ent.Schema
}

// Fields of the CampaignOrder.
func (CampaignOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("用户 id"),
		field.Int64("campaign_id").Default(0).StructTag(`json:"campaign_id"`).Comment("活动 id"),
	}
}

// Edges of the CampaignOrder.
func (CampaignOrder) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("campaign_orders").Field("user_id").Unique().Required(),
		edge.From("campaign", Campaign.Type).Ref("campaign_orders").Field("campaign_id").Unique().Required(),
		edge.To("cost_bills", CostBill.Type),
		edge.To("recharge_order", RechargeOrder.Type).Unique(),
	}
}

// Mixin of CampaignOrder
func (CampaignOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (CampaignOrder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("活动订单，计划废弃"),
	}
}
