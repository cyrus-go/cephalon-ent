package schema

import (
	"cephalon-ent/common"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").DefaultFunc(common.GenerateRandomString).StructTag(`json:"name"`).Comment("用户名"),
		field.String("nick_name").Default("请设置昵称").StructTag(`json:"nick_name"`).Comment("用户昵称"),
		field.String("phone").Default("").StructTag(`json:"phone"`).Comment("手机号"),
		field.String("password").Default("").Sensitive().Comment("密码"),
		field.String("avatar_url").Default("").StructTag(`json:"avatar_url"`).Comment("头像路径"),
		field.Enum("status").Values("normal", "frozen").Default("normal").StructTag(`json:"status"`).Comment("用户状态"),
		field.Enum("type").Values("personal", "enterprise").Default("personal").StructTag(`json:"type"`).Comment("用户类型"),
		field.Int("platform").Default(0).StructTag(`json:"platform"`).Comment("用户可以在什么平台登录，二进制开关数据"),
		field.String("hmac_key").Default("").StructTag(`json:"hmac_key"`).Comment("用户使用任务相关功能的密钥对的键，唯一"),
		field.String("hmac_secret").Default("").StructTag(`json:"hmac_secret"`).Comment("用户使用任务相关功能的密钥对的值"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bills", Bill.Type),
		edge.To("hmac_key_pair", HmacKeyPair.Type).Unique(),
		edge.To("created_missions", Mission.Type),
		edge.To("wallet", Wallet.Type).Unique(),
		edge.To("collections", Collection.Type),
		edge.To("devices", Device.Type),
		edge.To("profit_settings", ProfitSetting.Type),
		edge.To("mission_consume_orders", MissionConsumeOrder.Type),
		edge.To("mission_produce_orders", MissionProduceOrder.Type),
		edge.To("recharge_orders", RechargeOrder.Type),
		edge.To("vx_socials", VXSocial.Type),
		edge.To("mission_batches", MissionBatch.Type),
		edge.To("user_devices", UserDevice.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "deleted_at").Unique(),
		index.Fields("phone", "deleted_at").Unique(),
		index.Fields("hmac_key", "deleted_at").Unique(),
	}
}

// Mixin of User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of the BaseMixin.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("用户表，手机号唯一，用户名唯一，hmac_key 唯一"),
	}
}
