package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/common"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
	"time"
)

// UserCloseRecord holds the schema definition for the UserCloseRecord entity.
type UserCloseRecord struct {
	ent.Schema
}

// Fields of the UserCloseRecord.
func (UserCloseRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Default(0).StructTag(`json:"user_id,string"`).Comment("用户 id"),
		field.Time("registered_at").Default(common.ZeroTime).StructTag(`json:"registered_at"`).Comment("本次注销时的注册时间"),
		field.Time("closed_at").Default(time.Now()).StructTag(`json:"closed_at"`).Comment("本次注销的时间"),
		field.Enum("type").GoType(enums.UserCloseTypeSelf).Default(string(enums.UserCloseTypeUnknown)).StructTag(`json:"type"`).Comment("注销类型，用户自己注销或管理人员注销等"),
		field.Int64("operate_user_id").Default(0).StructTag(`json:"operate_user_id,string"`).Comment("操作人用户 id（只有管理人员注销时才有值）"),
	}
}

// Edges of the UserCloseRecord.
func (UserCloseRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_close_records").Field("user_id").Unique().Required(),
		edge.From("operate_user", User.Type).Ref("operate_user_close_records").Field("operate_user_id").Unique().Required(),
	}
}

// Indexes of UserCloseRecord
func (UserCloseRecord) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixin of UserCloseRecord
func (UserCloseRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (UserCloseRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("用户注销记录表"),
	}
}
