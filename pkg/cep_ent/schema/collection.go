package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Collection struct {
	ent.Schema
}

func (Collection) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").Default("").StructTag(`json:"url"`).Comment("路径"),
		field.Int64("user_id").Default(0).StructTag(`json:"user_id"`).Comment("外键用户 id"),
		field.String("picture_name").Default("").StructTag(`json:"picture_name"`).Comment("照片名称"),
	}
}

// Edges of the ServerOrder.
func (Collection) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("collections").Field("user_id").Unique().Required(),
	}
}

// Mixin of ServerOrder
func (Collection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of the BaseMixin.
func (Collection) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("资源收藏，与用户一对多"),
	}
}
