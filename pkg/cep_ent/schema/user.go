package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/common"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("请设置昵称").StructTag(`json:"name"`).Comment("用户名"),
		field.String("nick_name").Default("请设置昵称").StructTag(`json:"nick_name"`).Comment("用户昵称"),
		field.String("jpg_url").Default("").StructTag(`json:"jpg_url"`).Comment("头像"),
		field.String("key").Default("").StructTag(`json:"key"`).Comment("用户在任务中枢密钥对的键"),
		field.String("secret").Default("").Sensitive().Comment("用户在任务中枢的密钥对的值"),
		field.String("phone").Default("").StructTag(`json:"phone"`).Comment("用户的手机号"),
		field.String("password").Default("").Sensitive().Comment("密码"),
		field.Bool("is_frozen").Default(false).StructTag(`json:"is_frozen"`).Comment("是否冻结"),
		field.Bool("is_recharge").Default(false).StructTag(`json:"is_recharge"`).Comment("是否充值过"),
		field.Enum("user_type").GoType(enums.UserTypePersonal).Default(string(enums.UserTypePersonal)).StructTag(`json:"user_type"`).Comment("用户类型"),
		field.Int64("parent_id").Default(0).StructTag(`json:"parent_id,string"`).Comment("邀请人用户 id"),
		field.Int64("applet_parent_id").Default(0).StructTag(`json:"applet_parent_id"`).Comment("小程序邀请人用户 id"),
		field.String("pop_version").Default("").StructTag(`json:"pop_version"`).Comment("用户最新弹窗版本"),
		field.String("area_code").Default("+86").StructTag(`json:"area_code"`).Comment("国家区号"),
		field.String("email").Default("").StructTag(`json:"email"'`).Comment("邮箱"),
		field.String("github_id").Default("").StructTag(`json:"github_id"`).Comment("第三方登录github id"),
		field.String("google_id").Default("").StructTag(`json:"google_id"`).Comment("第三方登录google id"),
		field.Int64("cloud_space").Default(0).StructTag(`json:"cloud_space"`).Comment("云盘空间"),
		field.String("baidu_access_token").Default("").Sensitive().Comment("百度网盘 token"),
		field.String("baidu_refresh_token").Default("").Sensitive().Comment("百度网盘刷新 token"),
		field.Time("bound_at").Default(common.ZeroTime).Nillable().Optional().StructTag(`json:"bound_at"`).Comment("用户绑定邀请码的时间"),
		field.Enum("user_status").GoType(enums.UserStatusNormal).Default(string(enums.UserStatusNormal)).StructTag(`json:"user_status"`).Comment("用户状态"),
		field.Enum("channel").GoType(enums.UserChannelTypeNoChannel).Default(string(enums.UserChannelTypeNoChannel)).StructTag(`json:"channel"`).Comment("渠道身份，默认为非渠道用户"),
		//field.Int64("channel_ratio").StructTag(`json:"channel_ratio"`).Default(0).Comment("渠道分成比例"),
		field.Enum("mission_tag").GoType(enums.UserMissionTagNo).Default(string(enums.UserMissionTagNo)).StructTag(`json:"mission_tag"`).Comment("可跳过验证启动特殊任务类型标签"),
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
		edge.To("applet_children", User.Type).From("applet_parent").Unique().Required().Field("applet_parent_id"),
		edge.To("invites", Invite.Type),
		edge.To("campaign_orders", CampaignOrder.Type),
		edge.To("wallets", Wallet.Type),
		edge.To("withdraw_account", WithdrawAccount.Type).Unique(),
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
		edge.To("artworks", Artwork.Type),
		edge.To("artwork_likes", ArtworkLike.Type),
		edge.To("cdk_infos", CDKInfo.Type),
		edge.To("use_cdk_infos", CDKInfo.Type),
		edge.To("lotto_records", LottoRecord.Type),
		edge.To("lotto_user_counts", LottoUserCount.Type),
		edge.To("lotto_get_count_records", LottoGetCountRecord.Type),
		edge.To("cloud_files", CloudFile.Type),
		// edge.To("operate_transfer_orders", TransferOrder.Type),
		edge.To("withdraw_records", WithdrawRecord.Type),
		edge.To("operate_withdraw_records", WithdrawRecord.Type),
		edge.To("trouble_deducts", TroubleDeduct.Type),
		edge.To("income_manages", IncomeManage.Type),
		edge.To("approve_income_manages", IncomeManage.Type),
		edge.To("survey_responses", SurveyResponse.Type),
		edge.To("approve_survey_responses", SurveyResponse.Type),
		edge.To("mission_failed_feedbacks", MissionFailedFeedback.Type),
		edge.To("api_tokens", ApiToken.Type),
		edge.To("user_models", UserModel.Type),
		edge.To("invoke_model_orders", InvokeModelOrder.Type),
	}
}

// Indexes of User
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("phone"),
		index.Fields("email"),
		index.Fields("parent_id"),
		index.Fields("applet_parent_id"),
		index.Fields("channel"),
		//index.Fields("phone", "deleted_at", "email", "github_id", "google_id").Unique(),
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
