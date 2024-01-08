package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type MissionExtraService struct {
	ent.Schema
}

func (MissionExtraService) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("mission_id").StructTag(`json:"mission_id,string"`).Default(0).Comment("外键任务 id"),
		field.Int64("extra_service_id").StructTag(`json:"extra_service_id,string"`).Default(0).Comment("外键附加服务 id"),
	}
}

// Edges of the MissionExtraService.
func (MissionExtraService) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("mission", Mission.Type).Ref("mission_extra_services").Field("mission_id").Unique().Required(),
		edge.From("extra_service", ExtraService.Type).Ref("mission_extra_services").Field("extra_service_id").Unique().Required(),
	}
}

// Indexes of the MissionExtraService
func (MissionExtraService) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("mission_id"),
		index.Fields("extra_service_id"),
	}
}

// Mixin of MissionExtraService
func (MissionExtraService) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (MissionExtraService) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务与附加服务的中间关系"),
	}
}
