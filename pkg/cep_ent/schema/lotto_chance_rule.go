package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type LottoChanceRule struct {
	ent.Schema
}

func (LottoChanceRule) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("lotto_id").StructTag(`json:"lotto_id"`).Default(0).Comment("外键：抽奖活动 ID"),
		field.Enum("condition").GoType(enums.LottoConditionUnknown).Default(string(enums.LottoConditionUnknown)).StructTag(`json:"condition"`).Comment("条件"),
		field.Int64("award_count").StructTag(`json:"award_count"`).Default(0).Comment("奖励抽奖次数"),
		field.Int64("recharge_amount").StructTag(`json:"recharge_amount"`).Default(0).Comment("充值金额，类型为充值时才有数据"),
	}
}

// Edges of the LottoChanceRule.
func (LottoChanceRule) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("lotto", Lotto.Type).Ref("lotto_Change_rules").Field("lotto_id").Unique().Required(),
	}
}

// Indexes of LottoRecord
func (LottoChanceRule) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of LottoRecord
func (LottoChanceRule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (LottoChanceRule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("获取抽奖次数规则表"),
	}
}
