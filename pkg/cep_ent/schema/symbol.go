package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Symbol holds the schema definition for the Symbol entity.
type Symbol struct {
	ent.Schema
}

// Fields of the Symbol.
func (Symbol) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").StructTag(`json:"name"`).Comment("币种名称，唯一"),
	}
}

// Edges of the Symbol.
func (Symbol) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.To("wallets", Wallet.Type),
		edge.To("bills", Bill.Type),
		edge.To("mission_orders", MissionOrder.Type),
		edge.To("transfer_orders", TransferOrder.Type),
	}
}

// Mixin of Symbol
func (Symbol) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Symbol) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "deleted_at").Unique(),
	}
}

func (Symbol) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("币种，与用户多对多，通过钱包 Wallet"),
	}
}
