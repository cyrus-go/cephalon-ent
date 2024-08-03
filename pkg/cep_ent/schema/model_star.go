package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type ModelStar struct {
	ent.Schema
}

func (ModelStar) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Comment("用户ID"),
		field.Int64("model_id").StructTag(`json:"model_id"`).Comment("模型ID"),
		field.Enum("status").StructTag(`json:"status"`).Default(string(enums.UnknownStartStatus)).GoType(enums.Star).Comment("收藏状态"),
	}
}

func (ModelStar) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required().Field("user_id"),
		edge.To("model", Model.Type).Unique().Required().Field("model_id"),
	}
}

func (ModelStar) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
