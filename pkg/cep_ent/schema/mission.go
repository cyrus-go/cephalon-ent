package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// Mission holds the schema definition for the Mission entity.
type Mission struct {
	ent.Schema
}

// Fields of the Mission.
func (Mission) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").GoType(enums.MissionTypeSdTxt2Img).Default(string(enums.MissionTypeSdTxt2Img)).StructTag(`json:"type"`).Comment("任务类型"),
		field.String("body").StructTag(`json:"body"`).Default("").Comment("任务的请求参数体"),
		field.String("call_back_url").Default("").StructTag(`json:"call_back_url"`).Comment("回调地址，空字符串表示不回调"),
		field.Enum("status").GoType(enums.MissionStatusWaiting).Default(string(enums.MissionStatusWaiting)).StructTag(`json:"status"`).Comment("任务状态"),
		field.Enum("result").GoType(enums.MissionResultPending).Default(string(enums.MissionResultPending)).StructTag(`json:"result"`).Comment("任务结果，pending 表示还没有结果"),
		field.Strings("result_urls").Optional().Sensitive().Comment("任务结果资源位置列表序列化"),
		field.Int64("key_pair_id").Default(0).StructTag(`json:"key_pair_id"`).Comment("任务创建者的密钥对 ID"),
		field.String("mission_batch_number").Default("").StructTag(`json:"mission_batch_number"`).Comment("任务批次号"),
		field.Enum("gpu_version").GoType(enums.GpuVersion2060).Default(string(enums.GpuVersion2060)).StructTag(`json:"gpu_version"`).Comment("最低可接显卡"),
		field.Int64("unit_cep").Default(0).StructTag(`json:"unit_cep"`).Comment("任务单价，按次就是 unit_cep/次，按时就是 unit_cep/分钟"),
	}
}

// Edges of the Mission.
func (Mission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mission_key_pairs", MissionKeyPair.Type),
		edge.From("key_pair", HmacKeyPair.Type).Ref("created_missions").Field("key_pair_id").Unique().Required(),
	}
}

func (Mission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
