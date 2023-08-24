package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/common"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// MissionKeyPair holds the schema definition for the MissionKeyPair entity.
type MissionKeyPair struct {
	ent.Schema
}

// Fields of the MissionKeyPair.
func (MissionKeyPair) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("mission_id").StructTag(`json:"mission_id"`).Comment("任务 ID"),
		field.Int64("key_pair_id").StructTag(`json:"key_pair_id"`).Comment("密钥对 ID"),
		field.Time("started_at").Default(common.ZeroTime).StructTag(`json:"started_at"`).Comment("任务开始时刻"),
		field.Time("finished_at").Default(common.ZeroTime).StructTag(`json:"finished_at"`).Comment("任务完成时刻"),
		field.Enum("result").GoType(enums.MissionResultPending).Default(string(enums.MissionResultPending)).StructTag(`json:"result"`).Comment("任务结果"),
		field.Int64("device_id").Default(0).StructTag(`json:"device_id"`).Comment("领到任务的设备 ID"),
		field.Strings("result_urls").Optional().Sensitive().Comment("任务结果资源位置列表序列化"),
	}
}

// Edges of the MissionKeyPair.
func (MissionKeyPair) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("mission", Mission.Type).Ref("mission_key_pairs").Field("mission_id").Unique().Required(),
		edge.From("key_pair", HmacKeyPair.Type).Ref("mission_key_pairs").Field("key_pair_id").Unique().Required(),
	}
}

func (MissionKeyPair) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
