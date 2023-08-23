package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProfitAccount holds the schema definition for the ProfitAccount entity.
type ProfitAccount struct {
	ent.Schema
}

// Fields of the ProfitAccount.
func (ProfitAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键用户 id"),
		field.Int64("sum_cep").Default(0).StructTag(`json:"sum_cep"`).Comment("累计分润余额"),
		field.Int64("remain_cep").Default(0).StructTag(`json:"remain_cep"`).Comment("剩余分润余额，未提现的"),
	}
}

// Edges of the ProfitAccount.
func (ProfitAccount) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("profit_account").Field("user_id").Unique().Required(),
		edge.To("earn_bills", EarnBill.Type),
	}
}

// Mixin of ProfitAccount
func (ProfitAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
