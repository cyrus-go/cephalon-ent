package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type GpuPeak struct {
	ent.Schema
}

func (GpuPeak) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("version").GoType(enums.GpuVersion2060).Default(string(enums.GpuVersion2060)).StructTag(`json:"version"`).Comment("显卡型号"),
		field.Int("peak").Default(0).StructTag(`json:"power"`).Comment("每日同时占用峰值"),
	}
}

func (GpuPeak) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
	}
}

func (GpuPeak) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
