package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type LottoGetCountRecord struct {
	ent.Schema
}

func (LottoGetCountRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键：用户 ID"),
		field.Int64("lotto_id").StructTag(`json:"lotto_id"`).Default(0).Comment("外键：抽奖活动 ID"),
		field.Int64("count").StructTag(`json:"count"`).Default(0).Comment("此次奖励的抽奖次数"),
		field.Enum("type").Values("unknow", "register", "invite_register", "recharge").Default("unknow").StructTag(`json:"type"`).Comment("抽奖结果"),
		field.Int64("recharge_amount").StructTag(`json:"recharge_amount"`).Default(0).Comment("充值金额，类型为充值时才有数据"),
	}
}

// Edges of the LottoGetCountRecord.
func (LottoGetCountRecord) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("lotto_get_count_records").Field("user_id").Unique().Required(),
		edge.From("lotto", Lotto.Type).Ref("lotto_get_count_records").Field("lotto_id").Unique().Required(),
	}
}

// Indexes of LottoGetCountRecord
func (LottoGetCountRecord) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of LottoGetCountRecord
func (LottoGetCountRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (LottoGetCountRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("用户获得抽奖次数记录表"),
	}
}
