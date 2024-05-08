package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Brain struct {
	ent.Schema
}

func (Brain) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable(),
		field.String("user_id").Immutable(),
		field.String("question"),
		field.String("answer"),
	}
}

func (Brain) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Brain) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "t_brain",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		},
	}
}
