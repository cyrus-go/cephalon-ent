package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// FrpcInfo holds the schema definition for the FrpcInfo entity.
type FrpcInfo struct {
	ent.Schema
}

// Fields of the FrpcInfo.
func (FrpcInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("tag").Default("").StructTag(`json:"tag"`).Comment("ini 文件客户端 tag"),
		field.String("type").Default("").StructTag(`json:"type"`).Comment("frpc 通讯类型"),
		field.String("local_ip").Default("").StructTag(`json:"local_ip"`).Comment("frpc 本地地址"),
		field.Int("local_port").Default(0).StructTag(`json:"local_port"`).Comment("frpc 本地要使用端口"),
		field.Int("remote_port").Default(0).StructTag(`json:"remote_port"`).Comment("frpc 本地要使用端口对应的远程端口"),
		field.Bool("is_using").Default(false).StructTag(`json:"is_using"`).Comment("端口是否已经在使用"),

		field.Int64("frps_id").StructTag(`json:"frps_id,string"`).Default(0).Comment("外键 frps id"),
		field.Int64("device_id").StructTag(`json:"device_id,string"`).Default(0).Comment("外键设备 id"),
	}
}

// Edges of the FrpcInfo.
func (FrpcInfo) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("frps_info", FrpsInfo.Type).Ref("frpc_infos").Field("frps_id").Unique().Required(),
		edge.From("device", Device.Type).Ref("frpc_infos").Field("device_id").Unique().Required(),
	}
}

// Indexes of the FrpcInfo
func (FrpcInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("frps_id"),
		index.Fields("device_id"),
	}
}

// Mixin of FrpcInfo
func (FrpcInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (FrpcInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("fRPC 客户端端口分配情况"),
	}
}
