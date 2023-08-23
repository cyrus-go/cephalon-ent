package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// VXSocial holds the schema definition for the VXSocial entity.
type VXSocial struct {
	ent.Schema
}

// Fields of the VXSocial.
func (VXSocial) Fields() []ent.Field {
	return []ent.Field{
		field.String("app_id").Default("").StructTag(`json:"app_id"`).Comment("微信应用 id"),
		field.String("open_id").Default("").StructTag(`json:"open_id"`).Comment("微信身份源的 open_id"),
		field.String("union_id").Default("").StructTag(`json:"union_id"`).Comment("微信身份源的 union_id"),
		field.Enum("scope").Values("base").Default("base").StructTag(`json:"scope"`).Comment("账户的权限级别，一般为 base"),
		field.String("session_key").Default("").Sensitive().Comment("小程序专用的会话密钥，不可以返回给前端"),
		field.String("access_token").Default("").StructTag(`json:"access_token"`).Comment("微信能力访问凭证"),
		field.String("refresh_token").Default("").StructTag(`json:"refresh_token"`).Comment("刷新微信凭证的刷新凭证"),
		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键用户 id"),
	}
}

// Edges of the VXSocial.
func (VXSocial) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("vx_socials").Field("user_id").Unique().Required(),
		edge.To("recharge_orders", RechargeOrder.Type),
	}
}

// Mixin of VXSocial
func (VXSocial) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
