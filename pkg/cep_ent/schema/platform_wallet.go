package schema

import (
	"cephalon-ent/pkg/enums"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PlatformWallet holds the schema definition for the PlatformWallet entity.
type PlatformWallet struct {
	ent.Schema
}

// Fields of the PlatformWallet.
func (PlatformWallet) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").GoType(enums.PlatformWalletTypeProfit).Default(string(enums.PlatformWalletTypeProfit)).StructTag(`json:"type"`),
		field.Int64("sum_cep").Default(0).StructTag(`json:"sum_total_cep"`).Comment("累计总余额"),
		field.Int64("cep").Default(0).StructTag(`json:"total_cep"`).Comment("剩余总余额"),
	}
}

// Edges of the PlatformWallet.
func (PlatformWallet) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.To("bills", Bill.Type),
	}
}

// Indexes of the PlatformWallet
func (PlatformWallet) Indexes() []ent.Index {
	return []ent.Index{
		// 平台方每种类型的账户只有一个
		index.Fields("type", "deleted_at").Unique(),
	}
}

// Mixin of PlatformWallet
func (PlatformWallet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of the BaseMixin.
func (PlatformWallet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("平台钱包账户，独立于用户，每种类型的 cep 余额，比如分润获取的都集中到一条数据，即一个钱包"),
	}
}
