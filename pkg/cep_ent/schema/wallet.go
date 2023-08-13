package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Wallet holds the schema definition for the Wallet entity.
type Wallet struct {
	ent.Schema
}

// Fields of the Wallet.
func (Wallet) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键用户 id"),
		field.Int64("cep").Default(0).StructTag(`json:"cep"`).Comment("当前余额"),
		field.Int64("sum_cep").Default(0).StructTag(`json:"sum_cep"`).Comment("累计余额"),
	}
}

// Edges of the Wallet.
func (Wallet) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		// 和用户一对一
		edge.From("user", User.Type).Ref("wallet").Field("user_id").Unique().Required(),
		edge.To("bills", Bill.Type),
	}
}

// Mixin of Wallet
func (Wallet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of the BaseMixin.
func (Wallet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("钱包，与用户一对一"),
	}
}
