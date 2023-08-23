package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Collect struct {
	ent.Schema
}

func (Collect) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").Default("").StructTag(`json:"url"`).Comment("路径"),
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键用户 id"),
		field.Int64("jpg_name").StructTag(`json:"jpg_name"`).Default(0).Comment("照片名字"),
	}
}

// Edges of the ServerOrder.
func (Collect) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("collects").Field("user_id").Unique().Required(),
	}
}

// Mixin of ServerOrder
func (Collect) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
