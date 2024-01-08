package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// VXAccount holds the schema definition for the VXAccount entity.
type VXAccount struct {
	ent.Schema
}

// Fields of the VXAccount.
func (VXAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("open_id").Default("").StructTag(`json:"open_id"`).Comment("微信账户的 open_id"),
		field.String("union_id").Default("").StructTag(`json:"union_id"`).Comment("微信账户的 union_id"),
		field.String("scope").Default("base").StructTag(`json:"scope"`).Comment("账户的权限级别，一般为 base"),
		field.String("session_key").Default("").StructTag(`json:"session_key"`).Comment("会话密钥"),
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键用户 id"),
	}
}

// Edges of the VXAccount.
func (VXAccount) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("vx_accounts").Field("user_id").Unique().Required(),
	}
}

// Indexes of the VXAccount
func (VXAccount) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

// Mixin of VXAccount
func (VXAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of VXAccount
func (VXAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("原微信身份源，被 vx_socials 取代"),
	}
}
