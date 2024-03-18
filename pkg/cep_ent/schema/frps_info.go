package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FrpsInfo holds the schema definition for the FrpsInfo entity.
type FrpsInfo struct {
	ent.Schema
}

// Fields of the FrpsInfo.
func (FrpsInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("tag").Default("").StructTag(`json:"tag"`).Comment("ini 文件服务端 tag"),
		field.String("domain").Default("").StructTag(`json:"domain"`).Comment("域名"),
		field.String("server_addr").Default("").StructTag(`json:"server_addr"`).Comment("frps 服务地址"),
		field.Int("server_port").Default(0).StructTag(`json:"server_port"`).Comment("frps 服务端口"),
		field.String("authentication_method").Default("").StructTag(`json:"authentication_method"`).Comment("frps 认证方式"),
		field.String("token").Default("").StructTag(`json:"token"`).Comment("frps 认证 token"),
		field.String("type").Default("").StructTag(`json:"type"`).Comment("类型：public、private"),
	}
}

// Edges of the FrpsInfo.
func (FrpsInfo) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.To("frpc_infos", FrpcInfo.Type),
	}
}

// Mixin of FrpsInfo
func (FrpsInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (FrpsInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("fRPS 服务器列表"),
	}
}
