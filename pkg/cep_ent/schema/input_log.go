package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// InputLog holds the schema definition for the InputLog entity.
type InputLog struct {
	ent.Schema
}

// Fields of the InputLog.
func (InputLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("trace_id").Immutable().StructTag(`json:"trace_id"`).Comment("请求追踪 id"),
		field.String("headers").StructTag(`json:"headers"`).Comment("请求头"),
		field.String("body").Default("").Optional().StructTag(`json:"body"`).Comment("请求体"),
		field.String("query").Default("").Optional().StructTag(`json:"query"`).Comment("Query 参数"),
		field.String("url").StructTag(`json:"url"`).Comment("请求地址"),
		field.String("ip").Default("").Optional().StructTag(`json:"ip"`).Comment("客户端 IP"),
		field.String("caller").Default("general").StructTag(`json:"caller"`).Comment("调用方"),
		field.Enum("method").Values("GET", "POST", "PUT", "DELETE", "PATCH").StructTag(`json:"method"`).Comment("请求方式"),
		field.String("hmac_key").Default("").StructTag(`json:"hmac_key"`).Comment("记录调用者的密钥对"),
	}
}

// Edges of the InputLog.
func (InputLog) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (InputLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (InputLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("输入日志，初期风变专用"),
	}
}
