package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// WithdrawAccount holds the schema definition for the WithdrawAccount entity.
type WithdrawAccount struct {
	ent.Schema
}

// Fields of the Wallet.
func (WithdrawAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Default(0).StructTag(`json:"user_id,string"`).Comment("外键用户 id"),
		field.String("business_name").Default("0").StructTag(`json:"symbol_id,string"`).Comment("商户名称"),
		field.String("business_type").Default("0").StructTag(`json:"symbol_id,string"`).Comment("商户名称"),
		field.Int64("business_id").Default(0).StructTag(`json:"amount"`).Comment("货币余额"),
	}
}

// Edges of the WithdrawAccount.
func (WithdrawAccount) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("withdraw_accounts").Field("user_id").Unique().Required(),
	}
}

// Mixin of WithdrawAccount
func (WithdrawAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (WithdrawAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("提现账户，用来提供提现渠道"),
	}
}
