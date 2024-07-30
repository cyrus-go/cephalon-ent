package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ApiToken holds the schema definition for the ApiToken entity.
type ApiToken struct {
	ent.Schema
}

// Fields of the Token.
func (ApiToken) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").StructTag(`json:"user_id"`).Comment("用户ID"),
		field.String("name").Default("").StructTag(`json:"name"`).Comment("token 名称"),
		field.String("token").Default("").Sensitive().Comment("token 内容"),
	}
}

// Edges of the Token.
func (ApiToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("tokens").Field("user_id").Unique().Required(),
	}
}

func (ApiToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
