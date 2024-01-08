package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Artwork holds the schema definition for the Artwork entity.
type Artwork struct {
	ent.Schema
}

// Fields of the Artwork.
func (Artwork) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").StructTag(`json:"name"`).Comment("作品名称"),
		field.String("url").Default("").StructTag(`json:"url"`).Comment("作品链接"),
		field.Int64("author_id").Default(0).StructTag(`json:"author_id"`).Comment("作者的用户 id"),
	}
}

// Edges of the Artwork.
func (Artwork) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("author", User.Type).Ref("artworks").Field("author_id").Unique().Required(),
		edge.To("artwork_likes", ArtworkLike.Type),
	}
}

// Mixin of Artwork
func (Artwork) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Artwork) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("author_id"),
	}
}

func (Artwork) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("艺术作品，参与投票等逻辑；Artwork"),
	}
}
