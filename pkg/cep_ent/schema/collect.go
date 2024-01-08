package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Collect struct {
	ent.Schema
}

func (Collect) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").Default("").StructTag(`json:"url"`).Comment("路径"),
		field.Int64("user_id").StructTag(`json:"user_id,string"`).Default(0).Comment("外键用户 id"),
		field.Int64("jpg_name").StructTag(`json:"jpg_name"`).Default(0).Comment("照片名字"),
	}
}

// Edges of the Collect.
func (Collect) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("collects").Field("user_id").Unique().Required(),
	}
}

// Indexes of the Collect.
func (Collect) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

// Mixin of Collect.
func (Collect) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of Collect.
func (Collect) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("用户的图片收藏"),
	}
}
