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

// IncomeManage holds the schema definition for the IncomeManage entity.
type IncomeManage struct {
	ent.Schema
}

// Fields of the IncomeManage.
func (IncomeManage) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id,string"`).Default(0).Comment("用戶 id"),
		field.String("phone").StructTag(`json:"phone"`).Default("").Comment("用户手机号"),
		field.Enum("type").GoType(enums.IncomeManageTypeIncomeReplacement).Default(string(enums.IncomeManageTypeUnknown)).StructTag(`json:"type"`).Comment("类型"),
		field.Int64("amount").Default(0).StructTag(`json:"amount"`).Comment("金额，单位：厘"),
		field.String("reason").StructTag(`json:"reason"`).Default("").Comment("操作该用户收益钱包的原因"),
		field.Int64("current_balance").Default(0).StructTag(`json:"current_balance"`).Comment("当前余额（在生成这条记录时刻的余额），单位：厘"),
		field.Time("last_edited_at").Default(common.ZeroTime).StructTag(`json:"last_updated_at"`).Comment("审批前最后一次编辑的时间"),
		field.String("reject_reason").StructTag(`json:"reject_reason"`).Default("").Comment("拒绝此条记录原因"),
		field.Enum("status").GoType(enums.IncomeManageStatusPending).Default(string(enums.IncomeManageStatusPending)).StructTag(`json:"status"`).Comment("状态"),
		field.Int64("approve_user_id").StructTag(`json:"approve_user_id,string"`).Default(0).Comment("审批人 id"),
	}
}

// Edges of the IncomeManage.
func (IncomeManage) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("income_manages").Field("user_id").Unique().Required(),
		edge.From("approve_user", User.Type).Ref("approve_income_manages").Field("approve_user_id").Unique().Required(),
	}
}

// Indexes of the IncomeManage
func (IncomeManage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("approve_user_id"),
	}
}

// Mixin of IncomeManage
func (IncomeManage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of IncomeManage
func (IncomeManage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("收益补发记录，需要系统补发的收益，记录到这个表里"),
	}
}
