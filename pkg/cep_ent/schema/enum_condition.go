package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type EnumCondition struct {
	ent.Schema
}

func (EnumCondition) Fields() []ent.Field {
	return []ent.Field{
		field.String("front_type").Default("").StructTag(`json:"front_type"`).Comment("给到前端的类型"),
		field.String("mission_type").Default("").StructTag(`json:"mission_type"`).Comment("任务类型"),
		field.String("mission_call_way").Default("").StructTag(`json:"mission_call_way"`).Comment("任务调用方式"),
	}
}

// Edges of the ServerOrder.
func (EnumCondition) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of ServerOrder
func (EnumCondition) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (EnumCondition) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务类型枚举值对应表，计划用代码实现取代"),
	}
}
