package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// CloudFile holds the schema definition for the CloudFile entity.
type CloudFile struct {
	ent.Schema
}

// Fields of the CloudFile.
func (CloudFile) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键用户id"),
		field.String("name").Default("").StructTag(`json:"name"`).Comment("文件名"),
		field.String("icon").Default("").StructTag(`json:"icon"`).Comment("文件图标"),
		field.Int64("size").Default(0).StructTag(`json:"size"`).Comment("文件大小"),
		field.String("md5").Default("").StructTag(`json:"md5"`).Comment("md5"),
	}
}

// Edges of the CloudFile.
func (CloudFile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("cloud_files").Field("user_id").Unique().Required(),
	}
}

// Indexes of CloudFile
func (CloudFile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

// Mixin of CloudFile
func (CloudFile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
