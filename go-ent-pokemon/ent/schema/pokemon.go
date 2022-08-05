package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Pokemon holds the schema definition for the Pokemon entity.
type Pokemon struct {
	ent.Schema
}

// Fields of the Pokemon.
func (Pokemon) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StructTag(`json:"oid,omitempty"`),
		field.Text("name").NotEmpty(),
		field.Text("description").NotEmpty(),
		field.Float("weight"),
		field.Float("height"),
		field.Time("create_at").Default(time.Now).Immutable(),
		field.Time("update_at").Default(time.Now),
	}
}

// Edges of the Pokemon.
func (Pokemon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("fights", Battle.Type),
		edge.To("opponents", Battle.Type),
	}
}
