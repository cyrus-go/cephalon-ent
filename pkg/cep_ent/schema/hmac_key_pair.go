package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// HmacKeyPair holds the schema definition for the HmacKeyPair entity.
type HmacKeyPair struct {
	ent.Schema
}

// Fields of the HmacKeyPair.
func (HmacKeyPair) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Default("").StructTag(`json:"key"`).Comment("密钥对的 key 值，用于检索密钥"),
		field.String("secret").Default("").StructTag(`json:"secret"`).Comment("加密密钥"),
		field.String("caller").Default("").StructTag(`json:"caller"`).Comment("请求方"),
		field.Int64("user_id").Default(0).StructTag(`json:"user_id"`).Comment("外键用户 ID"),
	}
}

// Edges of the HmacKeyPair.
func (HmacKeyPair) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mission_productions", MissionProduction.Type),
		edge.To("created_missions", Mission.Type),
		edge.From("user", User.Type).Ref("hmac_key_pair").Unique().Required(),
	}
}

func (HmacKeyPair) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (HmacKeyPair) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key"),
	}
}

// Annotations of the BaseMixin.
func (HmacKeyPair) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("Hmac 密钥对，用于没有登录态时安全调用任务相关接口的场景"),
	}
}
