package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

type MissionCategory struct {
	ent.Schema
}

func (MissionCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("category").Unique().StructTag(`json:"category"`).Comment("任务大类"),
		field.Enum("type").GoType(enums.CategoryTypeHot).Default(string(enums.CategoryTypeUnknown)).StructTag(`json:"type"`).Comment("任务大类类型，比如热门类型等"),
		field.Int("weight").Default(0).StructTag(`json:"weight"`).Comment("权重（目前排序可以用到）"),
	}
}

// Edges of the MissionKind.
func (MissionCategory) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of MissionCategory
func (MissionCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (MissionCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务大类，任务类型的最高抽象层，记录了所有任务大类"),
	}
}
