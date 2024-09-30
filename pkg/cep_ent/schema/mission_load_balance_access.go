package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/common"
)

// MissionLoadBalanceAccess holds the schema definition for the MissionLoadBalanceAccess entity.
type MissionLoadBalanceAccess struct {
	ent.Schema
}

// Fields of the MissionLoadBalanceAccess.
func (MissionLoadBalanceAccess) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("mission_id").StructTag(`json:"mission_id,string"`).Comment("任务 ID"),
		field.Int64("mission_load_balance_id").StructTag(`json:"mission_load_balance_id,string"`).Comment("父 loadbanlace ID"),
		field.Time("last_access").Default(common.ZeroTime).StructTag(`json:"last_access"`).Comment("上次获取 api 时间"),
		field.Int32("access_count").Default(0).StructTag(`json:"access_count"`).Comment("已访问次数"),
	}
}

// Edges of the MissionLoadBalanceAccess.
func (MissionLoadBalanceAccess) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.From("mission", Mission.Type).Ref("mission_load_balance_accesss").Field("mission_id").Unique().Required(),
		// edge.From("mission_load_balance", Mission.Type).Ref("mission_load_balance_accesss").Field("mission_load_balance_id").Unique().Required(),
	}
}

// Indexes of the MissionLoadBalanceAccess
func (MissionLoadBalanceAccess) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("mission_id"),
		index.Fields("mission_load_balance_id"),
	}
}

func (MissionLoadBalanceAccess) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (MissionLoadBalanceAccess) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("mission负载均衡访问记录"),
	}
}
