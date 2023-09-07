package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FrpsInfo holds the schema definition for the FrpsInfo entity.
type FrpsInfo struct {
	ent.Schema
}

// Fields of the VXSocial.
func (FrpsInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("tag").Default("").StructTag(`json:"tag"`).Comment("ini 文件服务端 tag"),
		field.String("server_addr").Default("").StructTag(`json:"server_addr"`).Comment("frps 服务地址"),
		field.Int("server_port").Default(0).StructTag(`json:"server_port"`).Comment("frps 服务端口"),
	}
}

// Edges of the VXSocial.
func (FrpsInfo) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.To("frpc_infos", FrpcInfo.Type),
	}
}

// Mixin of VXSocial
func (FrpsInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
