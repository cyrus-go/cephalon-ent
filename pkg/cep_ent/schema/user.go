package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("请设置昵称").StructTag(`json:"name"`).Comment("用户名"),
		field.String("jpg_url").Default("").StructTag(`json:"jpg_url"`).Comment("头像"),
		field.String("key").Default("").StructTag(`json:"key"`).Comment("用户在任务中枢密钥对的键"),
		field.String("secret").Default("").Sensitive().Comment("用户在任务中枢的密钥对的值"),
		field.String("phone").Default("").StructTag(`json:"phone"`).Comment("用户的手机号"),
		field.String("password").Default("").Sensitive().Comment("密码"),
		field.Bool("is_frozen").Default(false).StructTag(`json:"is_frozen"`).Comment("是否冻结"),
		field.Bool("is_recharge").Default(false).StructTag(`json:"is_recharge"`).Comment("是否充值过"),
		field.Enum("user_type").Values("personal", "enterprise").Default("personal").StructTag(`json:"user_type"`).Comment("用户类型"),
		field.Int64("parent_id").Default(0).StructTag(`json:"parent_id"`).Comment("邀请人用户 id"),
		field.String("pop_version").Default("").StructTag(`json:"pop_version"`).Comment("用户最新弹窗版本"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("vx_accounts", VXAccount.Type),
		edge.To("collects", Collect.Type),
		edge.To("devices", Device.Type),
		edge.To("profit_settings", ProfitSetting.Type),
		edge.To("cost_account", CostAccount.Type).Unique(),
		edge.To("profit_account", ProfitAccount.Type).Unique(),
		edge.To("cost_bills", CostBill.Type),
		edge.To("earn_bills", EarnBill.Type),
		edge.To("mission_consume_orders", MissionConsumeOrder.Type),
		edge.To("mission_produce_orders", MissionProduceOrder.Type),
		edge.To("recharge_orders", RechargeOrder.Type),
		edge.To("vx_socials", VXSocial.Type),
		edge.To("mission_batches", MissionBatch.Type),
		edge.To("user_devices", UserDevice.Type),
		edge.To("children", User.Type).From("parent").Unique().Required().Field("parent_id"),
		edge.To("invites", Invite.Type),
		edge.To("campaign_orders", CampaignOrder.Type),
		edge.To("wallets", Wallet.Type),
		edge.To("income_bills", Bill.Type),
		edge.To("outcome_bills", Bill.Type),
		edge.To("mission_productions", MissionProduction.Type),
		edge.To("missions", Mission.Type),
		edge.To("income_transfer_orders", TransferOrder.Type),
		edge.To("outcome_transfer_orders", TransferOrder.Type),
		edge.To("consume_mission_orders", MissionOrder.Type),
		edge.To("produce_mission_orders", MissionOrder.Type),
		edge.To("login_records", LoginRecord.Type),
		edge.To("renewal_agreements", RenewalAgreement.Type),
	}
}

// Mixin of User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("用户表"),
	}
}
