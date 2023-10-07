package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LoginRecord holds the schema definition for the LoginRecord entity.
type LoginRecord struct {
	ent.Schema
}

// Fields of the LoginRecord.
func (LoginRecord) Fields() []ent.Field {
	return []ent.Field{
		field.String("ua").Default("").StructTag(`json:"ua"`).Comment("用户登录时的 user-agent"),
		field.String("ip").Default("").StructTag(`json:"ip"`).Comment("用户登录时的 ip 地址"),
		field.String("way").Default("").StructTag(`json:"way"`).Comment("用户登录时的登录方式，比如手机号验证码等"),

		field.Int64("user_id").Default(0).StructTag(`json:"user_id"`).Comment("用户 id，外键关联用户"),
	}
}

// Edges of the LoginRecord.
func (LoginRecord) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("login_records").Field("user_id").Unique().Required(),
	}
}

// Mixin of LoginRecord
func (LoginRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (LoginRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("登录记录，记录用户登录的一些信息"),
	}
}
