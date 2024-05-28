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

// TroubleDeduct holds the schema definition for the TroubleDeduct entity.
type TroubleDeduct struct {
	ent.Schema
}

// Fields of the TroubleDeduct.
func (TroubleDeduct) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id,string"`).Default(0).Comment("用戶 id"),
		field.Int64("device_id").StructTag(`json:"device_id,string"`).Default(0).Comment("设备 id"),
		field.Time("started_at").Default(common.ZeroTime).StructTag(`json:"started_at"`).Comment("故障开始时刻"),
		field.Time("finished_at").Default(common.ZeroTime).StructTag(`json:"finished_at"`).Comment("故障结束时刻"),
		field.Float("time_of_duration").Default(0).StructTag(`json:"time_of_duration"`).Comment("持续时长，单位：小时"),
		field.Int64("deduct_standard").Default(0).StructTag(`json:"deduct_standard"`).Comment("扣费标准，单位：厘"),
		field.Int64("amount").Default(0).StructTag(`json:"amount"`).Comment("扣费金额，单位：厘"),
		field.Int64("current_balance").Default(0).StructTag(`json:"current_balance"`).Comment("当前余额（在生成这条记录时刻的余额），单位：厘"),
		field.Enum("status").GoType(enums.TroubleDeductStatusPending).Default(string(enums.TroubleDeductStatusPending)).StructTag(`json:"status"`).Comment("状态"),
		field.String("reason").StructTag(`json:"reason"`).Default("").Comment("扣费原因"),
		field.String("reject_reason").StructTag(`json:"reject_reason"`).Default("").Comment("拒绝扣费原因"),
	}
}

// Edges of the TroubleDeduct.
func (TroubleDeduct) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("trouble_deducts").Field("user_id").Unique().Required(),
		edge.From("device", Device.Type).Ref("trouble_deducts").Field("device_id").Unique().Required(),
	}
}

// Indexes of the TroubleDeduct
func (TroubleDeduct) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("device_id"),
	}
}

// Mixin of TroubleDeduct
func (TroubleDeduct) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of TroubleDeduct
func (TroubleDeduct) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("故障扣费记录，节点故障时，需要扣费，记录到这个表里"),
	}
}
