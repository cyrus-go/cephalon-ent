package schema

import (
	"cephalon-ent/common"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MissionProduction holds the schema definition for the MissionProduction entity.
type MissionProduction struct {
	ent.Schema
}

// Fields of the MissionProduction.
func (MissionProduction) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("mission_id").StructTag(`json:"mission_id"`).Comment("任务 ID"),
		field.Int64("hmac_key_pair_id").StructTag(`json:"hmac_key_pair_id"`).Comment("密钥对 ID"),
		field.Time("started_at").Default(common.ZeroTime).StructTag(`json:"started_at"`).Comment("任务开始时刻"),
		field.Time("finished_at").Default(common.ZeroTime).StructTag(`json:"finished_at"`).Comment("任务完成时刻"),
		field.Enum("status").Values("pending", "failed", "succeed").Default("pending").StructTag(`json:"result"`).Comment("任务结果"),
		field.Int64("device_id").Default(0).StructTag(`json:"device_id"`).Comment("领到任务的设备 ID"),
		field.String("result_urls").Default("").Sensitive().Comment("任务结果资源位置列表序列化"),
		field.String("additional_result").Default("").Sensitive().Comment("额外需要返回的结果数据，格式不定"),
	}
}

// Edges of the MissionProduction.
func (MissionProduction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mission_produce_order", MissionProduceOrder.Type).Unique(),
		edge.From("mission", Mission.Type).Ref("mission_productions").Field("mission_id").Unique().Required(),
		edge.From("hmac_key_pair", HmacKeyPair.Type).Ref("mission_productions").Field("hmac_key_pair_id").Unique().Required(),
		edge.From("device", Device.Type).Ref("mission_productions").Field("device_id").Unique().Required(),
	}
}

func (MissionProduction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of the BaseMixin.
func (MissionProduction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务生产记录，任务被人接了就会产生一条记录，这一次任务完成情况就在这，同一个任务可以被多次接"),
	}
}
