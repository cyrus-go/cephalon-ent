package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CostAccount holds the schema definition for the CostAccount entity.
type CostAccount struct {
	ent.Schema
}

// Fields of the CostAccount.
func (CostAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键用户 id"),
		field.Int64("total_cep").Default(0).StructTag(`json:"total_cep"`).Comment("总余额"),
		field.Int64("sum_total_cep").Default(0).StructTag(`json:"sum_total_cep"`).Comment("累计总余额"),
		field.Int64("frozen_total_cep").Default(0).StructTag(`json:"frozen_total_cep"`).Comment("暂时冻结的总余额"),
		field.Int64("pure_cep").Default(0).StructTag(`json:"pure_cep"`).Comment("本金余额"),
		field.Int64("sum_pure_cep").Default(0).StructTag(`json:"sum_pure_cep"`).Comment("累计本金余额"),
		field.Int64("frozen_pure_cep").Default(0).StructTag(`json:"frozen_pure_cep"`).Comment("暂时冻结的本金余额"),
		field.Int64("gift_cep").Default(0).StructTag(`json:"gift_cep"`).Comment("赠送余额"),
		field.Int64("sum_gift_cep").Default(0).StructTag(`json:"sum_gift_cep"`).Comment("累计赠送余额"),
		field.Int64("frozen_gift_cep").Default(0).StructTag(`json:"frozen_gift_cep"`).Comment("暂时冻结的赠金余额"),
	}
}

// Edges of the CostAccount.
func (CostAccount) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("cost_account").Field("user_id").Unique().Required(),
		edge.To("cost_bills", CostBill.Type),
	}
}

// Mixin of CostAccount
func (CostAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
