package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PlatformAccount holds the schema definition for the PlatformAccount entity.
type PlatformAccount struct {
	ent.Schema
}

// Fields of the PlatformAccount.
func (PlatformAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values("profit", "market").Default("profit").StructTag(`json:"type"`),
		field.Int64("sum_total_cep").Default(0).StructTag(`json:"sum_total_cep"`).Comment("累计总余额"),
		field.Int64("total_cep").Default(0).StructTag(`json:"total_cep"`).Comment("剩余总余额"),
		field.Int64("sum_pure_cep").Default(0).StructTag(`json:"sum_pure_cep"`).Comment("累计本金"),
		field.Int64("pure_cep").Default(0).StructTag(`json:"pure_cep"`).Comment("剩余本金"),
		field.Int64("sum_gift_cep").Default(0).StructTag(`json:"sum_gift_cep"`).Comment("累计赠金"),
		field.Int64("gift_cep").Default(0).StructTag(`json:"gift_cep"`).Comment("剩余赠金"),
	}
}

// Edges of the PlatformAccount.
func (PlatformAccount) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.To("earn_bills", EarnBill.Type),
		edge.To("cost_bills", CostBill.Type),
	}
}

// Indexes of the PlatformAccount
func (PlatformAccount) Indexes() []ent.Index {
	return []ent.Index{
		// 平台方每种类型的账户只有一个
		index.Fields("type", "deleted_at").Unique(),
	}
}

// Mixin of PlatformAccount
func (PlatformAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
