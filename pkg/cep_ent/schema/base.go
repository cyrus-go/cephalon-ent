package schema

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"github.com/stark-sim/cephalon-ent/common"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	falsePtr := false
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable().
			Annotations(entsql.Annotation{Incremental: &falsePtr}).
			DefaultFunc(func() int64 {
				return common.GenSnowflakeInt64()
			}),
		//field.Int64("id").Immutable().DefaultFunc(tools.GenSnowflakeID()),
		field.Int64("created_by").Default(0).StructTag(`json:"created_by"`),
		field.Int64("updated_by").Default(0).StructTag(`json:"updated_by"`),
		field.Time("created_at").Immutable().Default(time.Now).StructTag(`json:"created_at"`),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).StructTag(`json:"updated_at"`),
		field.Time("deleted_at").Default(common.ZeroTime).StructTag(`json:"deleted_at"`),
	}
}

// Annotations of the BaseMixin.
func (BaseMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
	}
}
