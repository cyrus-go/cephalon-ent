package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// WithdrawRecord holds the schema definition for the WithdrawRecord entity.
type WithdrawRecord struct {
	ent.Schema
}

// Fields of the WithdrawRecord.
func (WithdrawRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id,string"`).Default(0).Comment("提现的用户 id"),
		field.String("withdraw_account").Default("").StructTag(`json:"withdraw_account"`).Comment("提现账户"),
		field.String("business_name").Default("").StructTag(`json:"business_name"`).Comment("威付通商户名称，对公时为户名"),
		field.String("bank").Default("未知银行").StructTag(`json:"bank"`).Comment("开户支行"),
		field.Enum("type").GoType(enums.WithdrawTypeWithdrawWX).Default(string(enums.WithdrawTypeUnknown)).StructTag(`json:"type"`).Comment("提现类型"),
		field.Int64("amount").Default(0).StructTag(`json:"amount"`).Comment("提现金额，单位：厘"),
		field.Int64("remain_amount").Default(0).StructTag(`json:"remain_amount"`).Comment("本次提现后余额，单位：厘"),
		field.Int64("rate").StructTag(`json:"withdraw_rate"`).Default(7).Comment("提现手续费率，100 为基准，比如手续费 7%，值就应该为 7，最大值不能超过 100, 默认 7%"),
		field.Int64("real_amount").StructTag(`json:"withdraw_real_amount"`).Default(0).Comment("提现实际到账（扣除手续费），单位：厘"),
		field.Enum("status").GoType(enums.WithdrawStatusPending).Default(string(enums.WithdrawStatusPending)).StructTag(`json:"status"`).Comment("转账订单的状态，比如微信发起支付后可能没完成支付"),
		field.String("reject_reason").Default("").StructTag(`json:"reject_reason"`).Comment("提现审批拒绝的理由"),
		field.Int64("operate_user_id").StructTag(`json:"operate_user_id,string"`).Default(0).Comment("操作的用户 id，手动充值或者提现审批才有数据，默认为 0"),
		field.Int64("transfer_order_id").StructTag(`json:"transfer_order_id,string"`).Default(0).Comment("对应的交易订单 id（一对一）"),
		field.Int64("symbol_id").Default(0).StructTag(`json:"symbol_id,string"`).Comment("消费的外键币种 id"),
	}
}

// Edges of the WithdrawRecord.
func (WithdrawRecord) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("withdraw_records").Field("user_id").Unique().Required(),
		edge.From("operate_user", User.Type).Ref("operate_withdraw_records").Field("operate_user_id").Unique().Required(),
		edge.From("transfer_order", TransferOrder.Type).Ref("withdraw_record").Field("transfer_order_id").Unique().Required(),
		edge.From("symbol", Symbol.Type).Ref("withdraw_records").Field("symbol_id").Unique().Required(),
	}
}

// Indexes of the WithdrawRecord
func (WithdrawRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("operate_user_id"),
		index.Fields("transfer_order_id"),
	}
}

// Mixin of WithdrawRecord
func (WithdrawRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of WithdrawRecord
func (WithdrawRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("提现记录，记录所有的提现信息"),
	}
}
