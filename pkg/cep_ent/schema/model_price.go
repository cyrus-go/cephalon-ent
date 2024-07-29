package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ModelPrice holds the schema definition for the ModelPrice entity.
type ModelPrice struct {
	ent.Schema
}

// Fields of the ModelPrice.
func (ModelPrice) Fields() []ent.Field {
	return []ent.Field{
		field.Int("input_price").Default(0).StructTag(`json:"input_price"`).Comment("输入价格"),
	}
}

// Edges of the ModelPrice.
func (ModelPrice) Edges() []ent.Edge {
	return nil
}
