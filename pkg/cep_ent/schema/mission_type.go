package schema

import (
	"cephalon-ent/pkg/enums"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// MissionType holds the schema definition for the MissionType entity.
type MissionType struct {
	ent.Schema
}

// Fields of the MissionType.
func (MissionType) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").GoType(enums.MissionTypeSdTxt2Img).Default(string(enums.MissionTypeSdTxt2Img)).StructTag(`json:"type"`).Comment("任务类型"),
		field.Enum("gpu").GoType(enums.GPU3070).Default(string(enums.GPU3070)).StructTag(`json:"gpu"`).Comment("显卡型号"),
		field.Int64("cep").Default(0).StructTag(`json:"cep"`).Comment("单价消耗 cep"),
		field.Bool("is_time").Default(false).StructTag(`json:"is_time"`).Comment("是否计时任务"),
		field.Enum("category").GoType(enums.MissionCategorySD).Default(string(enums.MissionCategorySD)).StructTag(`json:"category"`).Comment("任务种类，SD，Jupyter 等"),
	}
}

// Edges of the MissionType.
func (MissionType) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (MissionType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of the BaseMixin.
func (MissionType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("任务类型等信息，给任务定价归类等"),
	}
}
