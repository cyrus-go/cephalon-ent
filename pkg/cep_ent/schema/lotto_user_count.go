package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type LottoUserCount struct {
	ent.Schema
}

func (LottoUserCount) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键：用户 ID"),
		field.Int64("lotto_id").StructTag(`json:"lotto_id"`).Default(0).Comment("外键：抽奖活动 ID"),
		field.Int64("remain_lotto_count").StructTag(`json:"remain_lotto_count"`).Default(0).Comment("剩余抽奖次数"),
	}
}

// Edges of the LottoUserCount.
func (LottoUserCount) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("lotto_user_counts").Field("user_id").Unique().Required(),
		edge.From("lotto", Lotto.Type).Ref("lotto_user_counts").Field("lotto_id").Unique().Required(),
	}
}

// Indexes of LottoUserCount
func (LottoUserCount) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of LottoUserCount
func (LottoUserCount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (LottoUserCount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("用户可用抽奖次数表"),
	}
}
