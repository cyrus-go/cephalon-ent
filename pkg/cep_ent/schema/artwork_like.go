package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ArtworkLike holds the schema definition for the ArtworkLike entity.
type ArtworkLike struct {
	ent.Schema
}

// Fields of the ArtworkLike.
func (ArtworkLike) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Default(0).StructTag(`json:"user_id"`).Comment("外键，用户 id"),
		field.Int64("artwork_id").Default(0).StructTag(`json:"artwork_id"`).Comment("外键，作品 id"),
		field.Int32("date").Default(0).StructTag(`json:"date"`).Comment("投票日期"),
	}
}

// Edges of the ArtworkLike.
func (ArtworkLike) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("artwork_likes").Field("user_id").Unique().Required(),
		edge.From("artwork", Artwork.Type).Ref("artwork_likes").Field("artwork_id").Unique().Required(),
	}
}

// Mixin of ArtworkLike
func (ArtworkLike) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (ArtworkLike) Indexes() []ent.Index {
	return []ent.Index{}
}

func (ArtworkLike) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("用户投赞成票给艺术作品；ArtworkLike"),
	}
}
