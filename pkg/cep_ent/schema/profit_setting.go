package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type ProfitSetting struct {
	ent.Schema
}

// Fields of the ProfitSetting.
func (ProfitSetting) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id,string"`).Default(0).Comment("外键用户 id"),
		field.Int64("ratio").StructTag(`json:"ratio"`).Default(75).Comment("分润比例"),
	}
}

// Edges of the ProfitSetting.
func (ProfitSetting) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("profit_settings").Field("user_id").Unique().Required(),
	}
}

// Indexes of the ProfitSetting.
func (ProfitSetting) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

// Mixin of ProfitSetting.
func (ProfitSetting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of ProfitSetting.
func (ProfitSetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务分润比例设置，属于用户的信息，所有人默认 75%"),
	}
}
