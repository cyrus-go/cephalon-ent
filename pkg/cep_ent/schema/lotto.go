package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/common"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type Lotto struct {
	ent.Schema
}

func (Lotto) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").StructTag(`json:"name"`).Default("").Comment("抽奖活动名称"),
		field.Int64("total_weight").StructTag(`json:"total_weight"`).Default(0).Comment("活动总权重"),
		field.Time("started_at").Default(common.ZeroTime).StructTag(`json:"started_at"`).Comment("活动开始时间"),
		field.Time("ended_at").Default(common.ZeroTime).StructTag(`json:"ended_at"`).Comment("活动结束时间"),
		field.Enum("status").GoType(enums.LottoStatusNormal).Default(string(enums.LottoStatusUnknown)).StructTag(`json:"status"`).Comment("状态"),
		field.String("remark").StructTag(`json:"remark"`).Default("").Comment("备注信息"),
	}
}

// Edges of the Lotto.
func (Lotto) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.To("lotto_prizes", LottoPrize.Type),
		edge.To("lotto_records", LottoRecord.Type),
		edge.To("lotto_user_counts", LottoUserCount.Type),
		edge.To("lotto_get_count_records", LottoGetCountRecord.Type),
		edge.To("lotto_Change_rules", LottoChanceRule.Type),
	}
}

// Indexes of Lotto
func (Lotto) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
	}
}

// Mixin of Lotto
func (Lotto) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Lotto) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("抽奖活动表"),
	}
}
