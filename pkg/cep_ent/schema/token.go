package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

// Fields of the Token.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").StructTag(`json:"user_id"`).Comment("用户ID"),
		field.String("name").Default("").StructTag(`json:"name"`).Comment("token 名称"),
		field.String("token").Default("").Sensitive().Comment("token 内容"),
	}
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("tokens").Field("user_id").Unique().Required(),
	}
}

func (Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
