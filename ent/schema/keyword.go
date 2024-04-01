package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Keyword struct {
	ent.Schema
}

func (Keyword) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable(),
		field.String("name").Unique(),
		field.Strings("words").Default([]string{}),
		field.String("color").Default(""),
		field.Uint64("order").Default(0),
	}
}

func (Keyword) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order"),
	}
}

func (Keyword) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "t_keyword",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		},
	}
}
