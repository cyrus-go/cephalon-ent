package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/common"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// MissionProduction holds the schema definition for the MissionProduction entity.
type MissionProduction struct {
	ent.Schema
}

// Fields of the MissionProduction.
func (MissionProduction) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("mission_id").StructTag(`json:"mission_id,string"`).Comment("任务 ID"),
		field.Int64("user_id").StructTag(`json:"user_id,string"`).Comment("用户 ID"),
		field.Time("started_at").Default(common.ZeroTime).StructTag(`json:"started_at"`).Comment("任务开始时刻"),
		field.Time("finished_at").Default(common.ZeroTime).StructTag(`json:"finished_at"`).Comment("任务完成时刻"),
		field.Enum("state").Default(string(enums.MissionStateUnknown)).GoType(enums.MissionStateWaiting).StructTag(`json:"state"`).Comment("任务执行状态情况"),
		field.Int64("device_id").Default(0).StructTag(`json:"device_id,string"`).Comment("领到任务的设备 ID"),
		field.Enum("gpu_version").Default(string(enums.GpuVersionUnknown)).GoType(enums.GpuVersion3080).StructTag(`json:"gpu_version"`).Comment("任务使用什么显卡在执行"),
		field.Int8("device_slot").Default(0).StructTag(`json:"device_slot"`).Comment("显卡占用设备的插槽"),
		field.String("urls").Default("").StructTag(`json:"urls"`).Comment("任务结果链接列表，json 序列化后存储"),
		field.Int32("resp_status_code").Default(0).StructTag(`json:"resp_status_code"`).Comment("内部功能返回码"),
		field.String("resp_body").Default("").StructTag(`json:"resp_body"`).Comment("返回内容体 json 转 string"),
	}
}

// Edges of the MissionProduction.
func (MissionProduction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("mission", Mission.Type).Ref("mission_productions").Field("mission_id").Unique().Required(),
		edge.From("user", User.Type).Ref("mission_productions").Field("user_id").Unique().Required(),
		edge.To("mission_produce_order", MissionProduceOrder.Type).Unique(),
	}
}

func (MissionProduction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (MissionProduction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务执行情况，记录谁接了什么任务，用什么接，做得怎么样"),
	}
}
