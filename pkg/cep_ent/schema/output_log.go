package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// OutputLog holds the schema definition for the OutputLog entity.
type OutputLog struct {
	ent.Schema
}

// Fields of the OutputLog.
func (OutputLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("trace_id").Immutable().StructTag(`json:"trace_id"`).Comment("请求追踪 id"),
		field.String("headers").StructTag(`json:"headers"`).Comment("请求头"),
		field.String("body").Default("").Optional().StructTag(`json:"body"`).Comment("请求体"),
		field.String("url").StructTag(`json:"url"`).Comment("请求地址"),
		field.String("ip").Default("").Optional().StructTag(`json:"ip"`).Comment("客户端 IP"),
		field.String("caller").Default("general").StructTag(`json:"caller"`).Comment("调用方"),
		field.Int16("status").Default(200).Optional().StructTag(`json:"status"`).Comment("状态码"),
		field.String("hmac_key").Default("").StructTag(`json:"hmac_key"`).Comment("记录调用者的密钥对"),
	}
}

// Edges of the OutputLog.
func (OutputLog) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (OutputLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
