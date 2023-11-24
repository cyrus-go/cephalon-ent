package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ExtraService holds the schema definition for the ExtraService entity.
type ExtraService struct {
	ent.Schema
}

func (ExtraService) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").StructTag(`json:"name"`).Comment("附加服务名称"),
		field.Enum("extra_service_type").GoType(enums.ExtraServiceTypeVPN).Default(string(enums.ExtraServiceTypeUnknown)).StructTag(`json:"extra_service_type"`).Comment("附加服务类型"),
		field.Time("started_at").Optional().Nillable().StructTag(`json:"started_at"`).Comment("附加服务购买时间开始，为空表示永久有效"),
		field.Time("finished_at").Optional().Nillable().StructTag(`json:"finished_at"`).Comment("附加服务购买时间结束，为空表示永久有效"),
	}
}

// Edges of the ExtraService
func (ExtraService) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.To("missions", Mission.Type),
		edge.To("mission_extra_services", MissionExtraService.Type),
		edge.To("extra_service_prices", ExtraServicePrice.Type),
	}
}

// Mixin of ExtraService
func (ExtraService) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (ExtraService) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("附加服务表"),
	}
}
