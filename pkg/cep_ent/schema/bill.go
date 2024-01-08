package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type Bill struct {
	ent.Schema
}

// Fields of the Bill.
func (Bill) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").GoType(enums.BillTypeRecharge).Default(string(enums.BillTypeUnknown)).StructTag(`json:"type"`).Comment("流水的类型，对应的 order_id 关联哪张表依赖于该字段"),
		field.Int64("order_id").Default(0).Optional().StructTag(`json:"order_id,string"`).Comment("比如 type 为 mission 时关联任务订单。当为 0 时，流水没有详细订单信息"),
		field.Enum("way").GoType(enums.BillWayRechargeWechat).Default(string(enums.BillWayUnknown)).StructTag(`json:"way"`).Comment("额度账户流水的产生方式，微信、支付宝、计时消耗等，偏向于业务展示"),
		field.Int64("symbol_id").Default(0).StructTag(`json:"symbol_id,string"`).Comment("外键币种 id"),
		field.Int64("profit_symbol_id").Default(3).StructTag(`json:"profit_symbol_id,string"`).Comment("外键分润币种 id"),
		field.Int64("amount").Default(0).StructTag(`json:"amount"`).Comment("消耗多少货币金额"),
		field.Int64("target_user_id").Default(0).StructTag(`json:"target_user_id,string"`).Comment("流水目标钱包 id"),
		field.Int64("target_before_amount").Default(0).StructTag(`json:"target_before_amount"`).Comment("目标钱包期初金额"),
		field.Int64("target_after_amount").Default(0).StructTag(`json:"target_after_amount"`).Comment("目标钱包期末金额"),
		field.Int64("source_user_id").Default(0).StructTag(`json:"source_user_id,string"`).Comment("流水来源钱包 id"),
		field.Int64("source_before_amount").Default(0).StructTag(`json:"source_before_amount"`).Comment("来源钱包期初金额"),
		field.Int64("source_after_amount").Default(0).StructTag(`json:"source_after_amount"`).Comment("来源钱包期初金额"),
		field.String("serial_number").Default("").StructTag(`json:"serial_number"`).Comment("流水号，唯一"),
		field.Int64("invite_id").Default(0).StructTag(`json:"invite_id,string"`).Comment("外键关联某个邀请码"),
	}
}

// Edges of the Bill.
func (Bill) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("source_user", User.Type).Ref("outcome_bills").Field("source_user_id").Unique().Required(),
		edge.From("target_user", User.Type).Ref("income_bills").Field("target_user_id").Unique().Required(),
		edge.From("transfer_order", TransferOrder.Type).Ref("bills").Field("order_id").Unique(),
		edge.From("mission_order", MissionOrder.Type).Ref("bills").Field("order_id").Unique(),
		edge.From("invite", Invite.Type).Ref("bills").Field("invite_id").Unique().Required(),
		edge.From("symbol", Symbol.Type).Ref("bills").Field("symbol_id").Unique().Required(),
	}
}

// Indexes of Bill
func (Bill) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("source_user_id"),
		index.Fields("target_user_id"),
		index.Fields("order_id"),
		index.Fields("invite_id"),
		index.Fields("symbol_id"),
	}
}

// Mixin of Bill
func (Bill) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Bill) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("记录用户账户的变动，流水记录"),
	}
}
