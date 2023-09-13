package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// RechargeCampaignRule holds the schema definition for the RechargeCampaignRule entity.
type RechargeCampaignRule struct {
	ent.Schema
}

// Fields of the RechargeCampaignRule.
func (RechargeCampaignRule) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("little_value").Default(0).StructTag(`json:"little_value"`).Comment("充值范围下限"),
		field.Int64("large_value").Default(0).StructTag(`json:"large_value"`).Comment("充值范围上限"),
		field.Int64("gift_percent").Default(0).StructTag(`json:"gift_percent"`).Comment("赠送的比例"),
	}
}

// Edges of the RechargeCampaignRule.
func (RechargeCampaignRule) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
	}
}

// Mixin of RechargeCampaignRule
func (RechargeCampaignRule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
