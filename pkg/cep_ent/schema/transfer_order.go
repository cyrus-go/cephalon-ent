package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// TransferOrder holds the schema definition for the TransferOrder entity.
type TransferOrder struct {
	ent.Schema
}

// Fields of the TransferOrder.
func (TransferOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("source_user_id").StructTag(`json:"source_user_id,string"`).Default(0).Comment("转账来源的用户 id"),
		field.Int64("target_user_id").StructTag(`json:"target_user_id,string"`).Default(0).Comment("转账目标的用户 id"),
		field.Enum("status").Values("pending", "canceled", "succeed", "failed", "reexchange", "pending_order", "approved", "reject").Default("pending").StructTag(`json:"status"`).Comment("转账订单的状态，比如微信发起支付后可能没完成支付"),
		field.Int64("symbol_id").Default(0).StructTag(`json:"symbol_id,string"`).Comment("币种 id"),
		field.Int64("amount").Default(0).StructTag(`json:"amount"`).Comment("充值多少货币量"),
		field.Enum("type").GoType(enums.TransferOrderTypeRecharge).Default(string(enums.TransferOrderTypeUnknown)).StructTag(`json:"type"`).Comment("充值订单的类型"),
		field.String("serial_number").Default("").StructTag(`json:"serial_number"`).Comment("充值订单的序列号"),
		field.Int64("social_id").Default(0).Optional().StructTag(`json:"social_id,string"`).Comment("关联充值来源的身份源 id"),
		field.String("third_api_resp").Default("").StructTag(`json:"third_api_resp"`).Comment("第三方平台的返回，给到前端才能发起支付"),
		field.String("out_transaction_id").Default("").StructTag(`json:"out_transaction_id"`).Comment("平台方订单号"),
		field.Int64("operate_user_id").StructTag(`json:"operate_user_id,string"`).Default(0).Comment("操作的用户 id，管理后台手动充值才有数据，默认为 0"),
		//field.String("withdraw_account").Default("").StructTag(`json:"withdraw_account"`).Comment("提现账户（类型为提现才有数据）"),
		//field.Int64("withdraw_rate").StructTag(`json:"withdraw_rate"`).Default(7).Comment("提现手续费率，100 为基准，比如手续费 7%，值就应该为 7，最大值不能超过 100, 默认 7%"),
		//field.Int64("withdraw_real_amount").StructTag(`json:"withdraw_real_amount"`).Default(0).Comment("提现实际到账，单位：cep"),
		//field.String("reject_reason").Default("").StructTag(`json:"reject_reason"`).Comment("提现审批拒绝的理由"),
	}
}

// Edges of the TransferOrder.
func (TransferOrder) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("source_user", User.Type).Ref("outcome_transfer_orders").Field("source_user_id").Unique().Required(),
		edge.From("target_user", User.Type).Ref("income_transfer_orders").Field("target_user_id").Unique().Required(),
		edge.To("bills", Bill.Type),
		edge.From("vx_social", VXSocial.Type).Ref("transfer_orders").Field("social_id").Unique(),
		edge.From("symbol", Symbol.Type).Ref("transfer_orders").Field("symbol_id").Unique().Required(),
		//edge.From("operate_user", User.Type).Ref("operate_transfer_orders").Field("operate_user_id").Unique().Required(),
		edge.To("withdraw_record", WithdrawRecord.Type).Unique(),
	}
}

// Indexes of the TransferOrder
func (TransferOrder) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("source_user_id"),
		index.Fields("target_user_id"),
		index.Fields("social_id"),
		index.Fields("symbol_id"),
	}
}

// Mixin of TransferOrder
func (TransferOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of TransferOrder
func (TransferOrder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("转账订单，谁给谁转了多少什么货币"),
	}
}
