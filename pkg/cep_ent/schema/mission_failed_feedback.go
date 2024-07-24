package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type MissionFailedFeedback struct {
	ent.Schema
}

func (MissionFailedFeedback) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Default(0).StructTag(`json:"user_id,string"`).Comment("外键，反馈的用户 ID"),
		field.Int64("device_id").Default(0).StructTag(`json:"device_id,string"`).Comment("外键，反馈关联的设备 ID"),
		field.Int64("mission_id").Default(0).StructTag(`json:"mission_id,string"`).Comment("外键，反馈关联的任务 ID"),
		field.String("mission_name").StructTag(`json:"mission_name"`).Default("").Comment("应用名称"),
		field.Enum("status").GoType(enums.MissionFailedFeedbackStatusInit).Default(string(enums.MissionFailedFeedbackStatusInit)).StructTag(`json:"status"`).Comment("状态"),
		field.String("reason").StructTag(`json:"reason"`).Default("").Comment("任务失败的原因"),
	}
}

// Edges of the MissionFailedFeedback.
func (MissionFailedFeedback) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("mission_failed_feedbacks").Field("user_id").Unique().Required(),
		edge.From("device", Device.Type).Ref("mission_failed_feedbacks").Field("device_id").Unique().Required(),
		edge.From("mission", Mission.Type).Ref("mission_failed_feedback").Field("mission_id").Unique().Required(),
	}
}

// Indexes of MissionFailedFeedback
func (MissionFailedFeedback) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("device_id"),
		index.Fields("mission_id"),
		index.Fields("user_id", "mission_id"),
	}
}

// Mixin of MissionFailedFeedback
func (MissionFailedFeedback) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (MissionFailedFeedback) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("应用启动失败反馈信息表"),
	}
}
