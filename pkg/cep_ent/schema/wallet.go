package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Wallet holds the schema definition for the Wallet entity.
type Wallet struct {
	ent.Schema
}

// Fields of the Wallet.
func (Wallet) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Default(0).StructTag(`json:"user_id"`).Comment("外键用户 id"),
		field.Int64("symbol_id").Default(0).StructTag(`json:"symbol_id"`).Comment("外键币种 id"),
		field.Int64("amount").Default(0).StructTag(`json:"amount"`).Comment("货币余额"),
	}
}

// Edges of the Wallet.
func (Wallet) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("wallets").Field("user_id").Unique().Required(),
		edge.From("symbol", Symbol.Type).Ref("wallets").Field("symbol_id").Unique().Required(),
	}
}

// Mixin of Wallet
func (Wallet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Wallet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "symbol_id", "deleted_at").Unique(),
	}
}

func (Wallet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("钱包，作为用户和各币种的中间关系，记录各余额"),
	}
}
