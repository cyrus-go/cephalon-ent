package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/common"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// Mission holds the schema definition for the Mission entity.
type Mission struct {
	ent.Schema
}

// Fields of the Mission.
func (Mission) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").GoType(enums.MissionTypeSdTxt2Img).Default(string(enums.MissionTypeUnknown)).StructTag(`json:"type"`).Comment("任务类型"),
		field.Int64("mission_kind_id").Default(0).StructTag(`json:"mission_kind_id,string"`).Comment("外键，任务种类 id"),
		field.String("body").StructTag(`json:"body"`).Default("").Comment("任务的请求参数体"),
		field.String("call_back_url").Default("").StructTag(`json:"call_back_url"`).Comment("回调地址，空字符串表示不回调"),
		field.String("call_back_info").Optional().Nillable().StructTag(`json:"call_back_info"`).Comment("回调时带上的参数，接收任何类型数据后 json 压缩"),
		field.Bytes("call_back_data").Optional().Sensitive().Comment("回调时带上的参数，需 json 反序列化后才可使用，所以没有直接 json 序列化字段 (sensitive)"),
		field.Enum("status").GoType(enums.MissionStatusWaiting).Default(string(enums.MissionStatusWaiting)).StructTag(`json:"status"`).Comment("任务状态"),
		field.Enum("result").GoType(enums.MissionResultPending).Default(string(enums.MissionResultPending)).StructTag(`json:"result"`).Comment("任务结果，pending 表示还没有结果"),
		field.Enum("state").GoType(enums.MissionStateWaiting).Default(string(enums.MissionStateUnknown)).StructTag(`json:"state"`).Comment("新任务状态，整合原状态和结果"),
		field.Strings("result_urls").Optional().Sensitive().Comment("任务结果资源位置列表序列化"),
		field.String("urls").Default("").StructTag(`json:"urls"`).Comment("任务结果链接列表，json 序列化后存储"),
		field.Int64("key_pair_id").Default(0).StructTag(`json:"key_pair_id"`).Comment("任务创建者的密钥对 ID"),
		field.Int64("user_id").Default(0).StructTag(`json:"user_id,string"`).Comment("外键任务的创建者 id"),
		field.Int64("mission_batch_id").Default(0).StructTag(`json:"mission_batch_id,string"`).Comment("外键关联任务批次"),
		field.String("mission_batch_number").Default("").StructTag(`json:"mission_batch_number"`).Comment("任务批次号"),
		field.Enum("gpu_version").GoType(enums.GpuVersion2060).Default(string(enums.GpuVersionUnknown)).StructTag(`json:"gpu_version"`).Comment("最低可接显卡"),
		field.Int64("unit_cep").Default(0).StructTag(`json:"unit_cep"`).Comment("任务单价，按次(count)就是 unit_cep/次，按时(time)就是 unit_cep/分钟"),
		field.Int32("resp_status_code").Default(0).StructTag(`json:"resp_status_code"`).Comment("内部功能返回码"),
		field.String("resp_body").Default("").StructTag(`json:"resp_body"`).Comment("返回内容体 json 转 string"),
		field.String("inner_uri").Default("").StructTag(`json:"inner_uri"`).Comment("当 type 为 sd_api 时使用，为转发的 sd 内部接口相对路径"),
		field.Enum("inner_method").GoType(enums.InnerMethodPost).Default(string(enums.InnerMethodPost)).StructTag(`json:"inner_method"`).Comment("内部转发接口的请求方式，POST 或者 GET 等"),
		field.String("temp_hmac_key").Default("").StructTag(`json:"temp_hmac_key"`).Comment("当 type 为 key_pair 时，使用的临时密钥对的键"),
		field.String("temp_hmac_secret").Default("").StructTag(`json:"temp_hmac_secret"`).Comment("当 type 为 key_pair 时，使用的临时密钥对的值"),
		field.String("second_hmac_key").Default("").StructTag(`json:"second_hmac_key"`).Comment("创建任务时使用了的 二级 hmac_key"),
		field.String("username").Default("admin").StructTag(`json:"username"`).Comment("某些任务会使用到的验证用户名"),
		field.String("password").Default("cephalon").StructTag(`json:"password"`).Comment("某些任务会使用到的验证密码"),
		field.String("white_device_ids").GoType([]string{}).Optional().ValueScanner(common.Bytes2StringSliceValueScanner{}).SchemaType(map[string]string{dialect.Postgres: "bytea"}).StructTag(`json:"white_device_ids"`).Comment("任务的设备白名单"),
		field.String("black_device_ids").GoType([]string{}).Optional().ValueScanner(common.Bytes2StringSliceValueScanner{}).SchemaType(map[string]string{dialect.Postgres: "bytea"}).StructTag(`json:"black_device_ids"`).Comment("任务的设备黑名单"),
		field.Time("started_at").Default(common.ZeroTime).Nillable().Optional().StructTag(`json:"started_at"`).Comment("任务开始时间"),
		field.Time("finished_at").Default(common.ZeroTime).Nillable().Optional().StructTag(`json:"finished_at"`).Comment("任务结束时间"),
		field.Time("expired_at").Default(common.ZeroTime).Nillable().Optional().StructTag(`json:"expired_at"`).Comment("任务到期时间（包时任务才有）"),
		field.Time("free_at").Default(common.ZeroTime).SchemaType(map[string]string{dialect.Postgres: "timestamptz default '0001-01-01 00:00:00.00000 +00:00'"}).StructTag(`json:"free_at"`).Comment("任务释放时刻"),
	}
}

// Edges of the Mission.
func (Mission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("mission_kind", MissionKind.Type).Ref("missions").Field("mission_kind_id").Unique().Required(),
		edge.From("user", User.Type).Ref("missions").Field("user_id").Unique().Required(),
		edge.To("mission_key_pairs", MissionKeyPair.Type),
		edge.From("key_pair", HmacKeyPair.Type).Ref("created_missions").Field("key_pair_id").Unique().Required(),
		edge.To("mission_consume_order", MissionConsumeOrder.Type).Unique(),
		edge.To("mission_produce_orders", MissionProduceOrder.Type),
		edge.From("mission_batch", MissionBatch.Type).Ref("missions").Field("mission_batch_id").Unique().Required(),
		edge.To("mission_productions", MissionProduction.Type),
		edge.To("mission_orders", MissionOrder.Type),
		edge.To("renewal_agreements", RenewalAgreement.Type),
		edge.To("mission_extra_services", MissionExtraService.Type),
		edge.To("extra_services", ExtraService.Type),
		edge.To("extra_service_orders", ExtraServiceOrder.Type),
	}
}

func (Mission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Mission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务，具备任务要求，记录完成情况和结果，金额相关信息在 mission_consume_orders 等订单侧"),
	}
}
