package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type UserModel struct {
	ent.Schema
}

func (UserModel) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Comment("用户ID"),
		field.Int64("model_id").StructTag(`json:"model_id"`).Comment("模型ID"),
		field.Enum("relation").StructTag(`json:"relation"`).Default(string(enums.UnknownRelation)).GoType(enums.StarRelation).Comment("关系"),
		field.Enum("status").StructTag(`json:"status"`).Default(string(enums.UnknownModelRelation)).GoType(enums.StarModel).Comment("状态"),
	}
}

func (UserModel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required().Field("user_id"),
		edge.To("model", Model.Type).Unique().Required().Field("model_id"),
	}
}

func (UserModel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
