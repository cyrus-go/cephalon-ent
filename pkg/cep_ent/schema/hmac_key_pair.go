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
		field.String("key").StructTag(`json:"key"`).Comment("密钥对的 key 值，用于检索密钥"),
		field.String("secret").StructTag(`json:"secret"`).Comment("加密密钥"),
		field.String("caller").Default("").StructTag(`json:"caller"`).Comment("请求方"),
	}
}

// Edges of the HmacKeyPair.
func (HmacKeyPair) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mission_key_pairs", MissionKeyPair.Type),
		edge.To("created_missions", Mission.Type),
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

func (HmacKeyPair) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("密钥对，由于必须有用户才能使用密钥对，所以直接并入 user 的 key 和 secret 两个字段，废弃该表"),
	}
}
