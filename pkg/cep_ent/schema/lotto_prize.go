package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type LottoPrize struct {
	ent.Schema
}

func (LottoPrize) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("lotto_id").StructTag(`json:"lotto_id"`).Default(0).Comment("外键：抽奖活动 ID"),
		field.String("level_name").StructTag(`json:"level_name"`).Default("").Comment("奖品等级名称"),
		field.Int64("weight").StructTag(`json:"weight"`).Default(0).Comment("奖品等级权重"),
		field.String("name").StructTag(`json:"name"`).Default("").Comment("奖品名称"),
		field.Enum("status").Values("unknow", "normal", "canceled").Default("unknow").StructTag(`json:"status"`).Comment("状态"),
		field.Enum("type").Values("unknow", "get_cep").Default("unknow").StructTag(`json:"type"`).Comment("类型"),
		field.Int64("cep_amount").StructTag(`json:"cep_amount"`).Default(0).Comment("类型为 get_cep 时，cep 的数量"),
	}
}

// Edges of the LottoPrize.
func (LottoPrize) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("lotto", Lotto.Type).Ref("lotto_prizes").Field("lotto_id").Unique().Required(),
		edge.To("lotto_records", LottoRecord.Type),
	}
}

// Indexes of LottoPrize
func (LottoPrize) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}

// Mixin of LottoPrize
func (LottoPrize) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (LottoPrize) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "prizes"},
		schema.Comment("抽奖活动奖品表"),
	}
}
