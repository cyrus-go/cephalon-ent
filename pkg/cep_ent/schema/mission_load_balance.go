package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/common"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// MissionProduction holds the schema definition for the MissionProduction entity.
type MissionLoadBalance struct {
	ent.Schema
}

// Fields of the MissionProduction.
func (MissionLoadBalance) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("mission_type").GoType(enums.MissionTypeSdTxt2Img).Default(string(enums.MissionTypeUnknown)).StructTag(`json:"type"`).Comment("任务类型"),
		field.Int64("user_id").StructTag(`json:"user_id,string"`).Comment("用户 ID"),
		field.Enum("state").GoType(enums.MissionLoadBalanceStateWaiting).Default(string(enums.MissionLoadBalanceStateUnknown)).StructTag(`json:"state"`).Comment("mission 负载均衡总状态"),
		field.Time("started_at").Default(common.ZeroTime).StructTag(`json:"started_at"`).Comment("任务开始时刻"),
		field.Time("finished_at").Default(common.ZeroTime).StructTag(`json:"finished_at"`).Comment("任务完成时刻"),
		field.Enum("gpu_version").Default(string(enums.GpuVersionUnknown)).GoType(enums.GpuVersion3080).StructTag(`json:"gpu_version"`).Comment("任务使用什么显卡在执行"),
		field.Int8("gpu_num").Default(0).StructTag(`json:"gpu_num"`).Comment("使用显卡数量"),
		field.Int8("max_mission_count").Default(1).StructTag(`json:"max_mission_count"`).Comment("应用数浮动上限"),
		field.Int8("min_mission_count").Default(1).StructTag(`json:"min_mission_count"`).Comment("应用数浮动下限"),
		field.Int8("current_mission_count").Default(1).StructTag(`json:"current_mission_count"`).Comment("当前应用数（调整中的以此为依据）"),
		field.Int64("mission_batch_id").Default(0).StructTag(`json:"mission_batch_id,string"`).Comment("外键关联任务批次"),
		field.String("mission_batch_number").Default("").StructTag(`json:"mission_batch_number"`).Comment("任务批次号"),
	}
}

// Edges of the MissionProduction.
func (MissionLoadBalance) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.From("mission", Mission.Type).Ref("mission_load_balances").Field("mission_id").Unique().Required(),
		// edge.From("user", User.Type).Ref("mission_load_balances").Field("user_id").Unique().Required(),
	}
}

// Indexes of the MissionProduction
func (MissionLoadBalance) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("mission_batch_id"),
		index.Fields("mission_batch_number"),
	}
}

func (MissionLoadBalance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (MissionLoadBalance) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务负载均衡，多个应用自动启动，访问多的时候多启点，访问少的时候减少点"),
	}
}
