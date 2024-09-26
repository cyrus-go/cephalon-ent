package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// RechargeCampaignRuleOversea holds the schema definition for the RechargeCampaignRuleOversea entity.
type RechargeCampaignRuleOversea struct {
	ent.Schema
}

// Fields of the RechargeCampaignRuleOversea.
func (RechargeCampaignRuleOversea) Fields() []ent.Field {
	return []ent.Field{
		field.Float("dollar_price").Default(0).StructTag(`json:"dollar_price"`).Comment("美元价格"),
		field.Float("rmb_price").Default(0).StructTag(`json:"rmb_price"`).Comment("RMB 价格"),
		field.Float("original_rmb_price").Default(0).StructTag(`json:"original_rmb_price"`).Comment("RMB 原价"),
		field.Int64("total_cep").Default(0).StructTag(`json:"total_cep"`).Comment("总共能获得的脑力值"),
		field.Int64("before_discount_cep").Default(0).StructTag(`json:"before_discount_cep"`).Comment("优惠前能获得的脑力值"),
		field.Int64("discount_ratio").Default(0).StructTag(`json:"discount_ratio"`).Comment("优惠力度（现价基于原价的比例）"),
	}
}

// Edges of the RechargeCampaignRuleOversea.
func (RechargeCampaignRuleOversea) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
	}
}

// Mixin of RechargeCampaignRuleOversea.
func (RechargeCampaignRuleOversea) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of RechargeCampaignRuleOversea.
func (RechargeCampaignRuleOversea) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("海外充值活动的规则，死表，与国内区分，国内是赠送充值比例，海外是固定"),
	}
}
