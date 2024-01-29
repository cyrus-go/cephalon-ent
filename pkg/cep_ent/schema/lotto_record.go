package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type LottoRecord struct {
	ent.Schema
}

func (LottoRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键：用户 ID"),
		field.Int64("lotto_id").StructTag(`json:"lotto_id"`).Default(0).Comment("外键：抽奖活动 ID"),
		field.Enum("result").Values("unknow", "winning", "losing", "failed").Default("unknow").StructTag(`json:"result"`).Comment("抽奖结果"),
		field.Int64("lotto_prize_id").StructTag(`json:"lotto_prize_id"`).Default(0).Comment("外键：奖品 ID"),
		field.Enum("status").Values("waiting", "granted", "not_grant", "invalid").Default("waiting").StructTag(`json:"status"`).Comment("抽奖状态"),
		field.Int64("remain_lotto_count").StructTag(`json:"remain_lotto_count"`).Default(0).Comment("剩余抽奖次数"),
	}
}

// Edges of the LottoRecord.
func (LottoRecord) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("lotto_records").Field("user_id").Unique().Required(),
		edge.From("lotto", Lotto.Type).Ref("lotto_records").Field("lotto_id").Unique().Required(),
		edge.From("lotto_prize", LottoPrize.Type).Ref("lotto_records").Field("lotto_prize_id").Unique().Required(),
	}
}

// Indexes of LottoRecord
func (LottoRecord) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of LottoRecord
func (LottoRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (LottoRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("用户抽奖记录表"),
	}
}
