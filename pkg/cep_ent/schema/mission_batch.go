package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type MissionBatch struct {
	ent.Schema
}

func (MissionBatch) Fields() []ent.Field {
	return []ent.Field{
		field.String("number").Default("").StructTag(`json:"number"`).Comment("批次号码"),
		field.Int64("user_id").Default(0).StructTag(`json:"user_id,string"`).Comment("创建者"),
	}
}

// Edges of the ServerOrder.
func (MissionBatch) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("mission_batches").Field("user_id").Unique().Required(),
		edge.To("mission_consume_orders", MissionConsumeOrder.Type),
		edge.To("missions", Mission.Type),
		edge.To("mission_orders", MissionOrder.Type),
	}
}

// Mixin of ServerOrder
func (MissionBatch) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (MissionBatch) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务批次，为同一次性创建的多个任务赋予关系记录"),
	}
}
