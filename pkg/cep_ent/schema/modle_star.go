package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Token holds the schema definition for the Token entity.
type ModleStar struct {
	ent.Schema
}

// Fields of the Token.
func (ModleStar) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").StructTag(`json:"user_id"`).Comment("用户ID"),
		field.Int("model_id").StructTag(`json:"model_id"`).Comment("模型ID"),
	}
}

// Edges of the Token.
func (ModleStar) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("tokens").Field("user_id").Unique().Required(),
		edge.From("model", Model.Type).Ref("tokens").Field("model_id").Unique().Required(),
	}
}

func (ModleStar) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
