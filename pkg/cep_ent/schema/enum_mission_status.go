package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type EnumMissionStatus struct {
	ent.Schema
}

func (EnumMissionStatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("front_status").Default("").StructTag(`json:"front_status"`).Comment("给到前端的任务状态"),
		field.String("mission_type").Default("").StructTag(`json:"mission_type"`).Comment("任务类型"),
		field.String("mission_status").Default("").StructTag(`json:"mission_status"`).Comment("任务状态"),
	}
}

// Edges of the ServerOrder.
func (EnumMissionStatus) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of ServerOrder
func (EnumMissionStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
