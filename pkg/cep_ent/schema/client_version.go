package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// ClientVersion holds the schema definition for the ClientVersion entity.
type ClientVersion struct {
	ent.Schema
}

// Fields of the ClientVersion.
func (ClientVersion) Fields() []ent.Field {
	return []ent.Field{
		field.String("client_url").Default("").StructTag(`json:"client_url"`).Comment("客户端文件地址"),
		field.String("config_url").Default("").StructTag(`json:"config_url"`).Comment("主配置文件地址"),
		field.String("version").Unique().StructTag(`json:"version"`).Comment("版本号"),
		field.Enum("status").GoType(enums.ClientStatusDisable).Default(string(enums.ClientStatusDisable)).StructTag(`json:"status"`).Comment("状态：只允许有一条数据的状态为启用（可被自动更新的版本）"),
	}
}

// Edges of the ClientVersion.
func (ClientVersion) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
	}
}

// Indexes of ClientVersion.
func (ClientVersion) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status"),
	}
}

// Mixin of ClientVersion.
func (ClientVersion) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of ClientVersion.
func (ClientVersion) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("客户端版本，OTA 服务端功能支持"),
	}
}
