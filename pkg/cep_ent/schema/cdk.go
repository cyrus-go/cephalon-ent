package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/common"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type CDKInfo struct {
	ent.Schema
}

func (CDKInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("issue_user_id").StructTag(`json:"issue_user_id,omitempty,string"`).Default(0).Comment("外键：发行用户 id"),
		field.String("cdk_number").StructTag(`json:"cdk_number"`).Default("").Comment("cdk 序列号"),
		field.Enum("type").GoType(enums.CDKTypeGetCep).Default(string(enums.CDKTypeUnknown)).StructTag(`json:"type"`).Comment("cdk 类型"),
		field.Int64("get_cep").StructTag(`json:"get_cep"`).Default(0).Comment("cdk 能兑换的 cep 数量"),
		field.Int64("get_time").StructTag(`json:"get_time"`).Default(0).Comment("cdk 能兑换的 gpu 使用时长"),
		field.Enum("billing_type").GoType(enums.MissionBillingTypeTimePlanHour).Default(string(enums.MissionBillingTypeUnknown)).StructTag(`json:"billing_type"`).Comment("兑换 gpu 使用时长的类型"),
		field.Time("expired_at").Default(common.ZeroTime).Nillable().Optional().StructTag(`json:"expired_at"`).Comment("过期时间"),
		field.Int64("use_times").StructTag(`json:"use_times"`).Default(0).Comment("cdk 能使用的次数"),
		field.Enum("status").GoType(enums.CDKStatusNormal).Default(string(enums.CDKStatusUnknown)).StructTag(`json:"status"`).Comment("cdk 状态"),
	}
}

// Edges of the CDKInfo.
func (CDKInfo) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("issue_user", User.Type).Ref("cdk_infos").Field("issue_user_id").Unique().Required(),
	}
}

// Indexes of CDKInfo
func (CDKInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("cdk_number"),
	}
}

// Mixin of CDKInfo
func (CDKInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (CDKInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "cdk_infos"},
		schema.Comment("兑换码，可以兑换脑力值、gpu 使用权等"),
	}
}
