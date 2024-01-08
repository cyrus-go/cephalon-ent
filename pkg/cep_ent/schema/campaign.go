package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/common"
)

// Campaign holds the schema definition for the Campaign entity.
type Campaign struct {
	ent.Schema
}

// Fields of the Campaign.
func (Campaign) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").StructTag(`json:"name"`).Comment("活动名称"),
		field.String("type").Default("").StructTag(`json:"type"`).Comment("活动类型"),
		field.Time("started_at").Default(common.ZeroTime).StructTag(`json:"started_at"`).Comment("活动开始时间"),
		field.Time("ended_at").Default(common.ZeroTime).StructTag(`json:"ended_at"`).Comment("活动结束时间"),
		field.Int("status").Default(0).StructTag(`json:"status"`).Comment("活动状态"),

		field.String("invite_id").Default("").StructTag(`json:"invite_id"`).Comment("外键邀请码 id"),
	}
}

// Edges of the Campaign.
func (Campaign) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.To("invites", Invite.Type),
		edge.To("campaign_orders", CampaignOrder.Type),
	}
}

// Mixin of Campaign
func (Campaign) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Campaign) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("活动，计划废弃"),
	}
}
