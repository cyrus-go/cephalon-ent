package schema

import (
	"cephalon-ent/pkg/enums"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Mission holds the schema definition for the Mission entity.
type Mission struct {
	ent.Schema
}

// Fields of the Mission.
func (Mission) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").GoType(enums.MissionTypeSdTxt2Img).Default(string(enums.MissionTypeSdTxt2Img)).StructTag(`json:"type"`).Comment("任务类型"),
		field.Bool("is_time").Default(false).StructTag(`json:"is_time"`).Comment("是否为计时类型任务"),
		field.String("body").Default("").StructTag(`json:"body"`).Comment("任务的请求参数体"),
		field.String("call_back_url").Default("").StructTag(`json:"call_back_url"`).Comment("回调地址，空字符串表示不回调"),
		field.Enum("status").Values("waiting", "canceled", "doing", "supplying", "closing", "succeed", "failed").Default("waiting").StructTag(`json:"status"`).Comment("任务状态"),
		field.String("result_urls").Default("").Sensitive().Comment("任务结果资源位置列表序列化"),
		field.String("additional_result").Default("").Sensitive().Comment("有的任务除了链接外还有其他有用的结果，都塞在这个字段，比如 sd 的实际入参"),
		field.Int64("hmac_key_pair_id").Default(0).StructTag(`json:"hmac_key_pair_id"`).Comment("外键任务创建者的密钥对 ID"),
		field.Int64("user_id").Default(0).StructTag(`json:"user_id"`).Comment("外键任务创建者的 ID"),
		field.String("mission_batch_number").Default("").StructTag(`json:"batch_number"`).Comment("任务批次号"),
		field.Int64("mission_batch_id").Default(0).StructTag(`json:"mission_batch_id"`).Comment("外键任务批次"),
	}
}

// Edges of the Mission.
func (Mission) Edges() []ent.Edge {
	return []ent.Edge{
		// 同一个任务可以被多次执行
		edge.To("mission_productions", MissionProduction.Type),
		edge.To("mission_consume_order", MissionConsumeOrder.Type).Unique(),
		edge.To("mission_produce_orders", MissionProduceOrder.Type),
		edge.From("hmac_key_pair", HmacKeyPair.Type).Ref("created_missions").Field("hmac_key_pair_id").Unique().Required(),
		edge.From("user", User.Type).Ref("created_missions").Field("user_id").Unique().Required(),
		edge.From("mission_batch", MissionBatch.Type).Ref("missions").Field("mission_batch_id").Unique().Required(),
	}
}

func (Mission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
