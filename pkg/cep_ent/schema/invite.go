package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Invite holds the schema definition for the Invite entity.
type Invite struct {
	ent.Schema
}

// Fields of the Invite.
func (Invite) Fields() []ent.Field {
	return []ent.Field{
		field.String("invite_code").Default("").StructTag(`json:"invite_code"`).Comment("邀请码"),
		field.Int64("share_cep").Default(0).StructTag(`json:"share_cep"`).Comment("通过此邀请码分享能获得的收益"),
		field.Int64("reg_cep").Default(0).StructTag(`json:"reg_cep"`).Comment("通过此邀请码注册能获得的收益"),
		field.String("type").Default("").StructTag(`json:"type"`).Comment("邀请码类型（可以用来区分不同的活动）"),

		field.Int64("user_id").StructTag(`json:"user_id"`).Default(0).Comment("外键用户 id"),
	}
}

// Edges of the Invite.
func (Invite) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("invites").Field("user_id").Unique().Required(),
	}
}

// Indexes of the Invite
func (Invite) Indexes() []ent.Index {
	return []ent.Index{
		// 邀请码索引
		index.Fields("invite_code"),
	}
}

// Mixin of Invite
func (Invite) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}