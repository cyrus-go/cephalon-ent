package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/common"
)

// RenewalAgreement holds the schema definition for the RenewalAgreement entity.
type RenewalAgreement struct {
	ent.Schema
}

// Fields of the RenewalAgreement.
func (RenewalAgreement) Fields() []ent.Field {
	return []ent.Field{
		field.Time("next_pay_time").Default(common.ZeroTime).StructTag(`json:"next_pay_time"`).Comment("下次扣款时间"),
		field.Enum("type").Values("unknow", "hour", "day", "month").Default("unknow").StructTag(`json:"type"`).Comment("自动续费类型（按小时、按天等）"),
		field.Enum("sub_status").Values("unknow", "subscribing", "finished").Default("unknow").StructTag(`json:"sub_status"`).Comment("订阅自动续费状态（订阅中、已结束等）"),
		field.Enum("pay_status").Values("unknow", "waiting", "succeed", "failed").Default("unknow").StructTag(`json:"pay_status"`).Comment("支付状态（待支付、已支付、支付失败等）"),
		field.Int64("first_pay").Default(0).StructTag(`json:"first_pay"`).Comment("首次扣款价格"),
		field.Int64("after_pay").Default(0).StructTag(`json:"after_pay"`).Comment("后续扣款价格"),
		field.Time("sub_finished_time").Default(common.ZeroTime).StructTag(`json:"sub_finished_time"`).Comment("订阅自动续费结束时间"),

		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键用户 id"),
		field.Int64("mission_id").StructTag(`json:"mission_id"`).Default(0).Comment("外键任务 id"),
	}
}

// Edges of the RenewalAgreement.
func (RenewalAgreement) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("renewal_agreements").Field("user_id").Unique().Required(),
		edge.From("mission", Mission.Type).Ref("renewal_agreements").Field("mission_id").Unique().Required(),
	}
}

// Mixin of RenewalAgreement
func (RenewalAgreement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (RenewalAgreement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("自动续费协议"),
	}
}
