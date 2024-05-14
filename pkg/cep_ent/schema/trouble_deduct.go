package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		field.Int64("device_id").StructTag(`json:"device_id,string"`).Default(0).Comment("设备 id"),
		field.Time("started_at").Default(common.ZeroTime).StructTag(`json:"started_at"`).Comment("故障开始时刻"),
		field.Time("finished_at").Default(common.ZeroTime).StructTag(`json:"finished_at"`).Comment("故障结束时刻"),
		field.Float("time_of_duration").Default(0).StructTag(`json:"time_of_duration,string"`).Comment("持续时长，单位：小时"),
		field.Int64("amount").Default(0).StructTag(`json:"amount"`).Comment("扣费金额，单位：分"),
		field.Enum("status").GoType(enums.TroubleDeductStatusPending).Default(string(enums.TroubleDeductStatusPending)).StructTag(`json:"status"`).Comment("状态"),
		field.String("reason").StructTag(`json:"reason"`).Default("").Comment("扣费原因"),
		field.String("cancel_reason").StructTag(`json:"cancel_reason"`).Default("").Comment("取消扣费原因"),
	}
}

// Edges of the TroubleDeduct.
func (TroubleDeduct) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("device", Device.Type).Ref("trouble_deducts").Field("device_id").Unique().Required(),
	}
}

// Indexes of the TroubleDeduct
func (TroubleDeduct) Indexes() []ent.Index {
	return []ent.Index{}
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
