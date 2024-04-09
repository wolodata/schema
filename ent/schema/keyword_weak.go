package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type KeywordWeak struct {
	ent.Schema
}

func (KeywordWeak) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable(),
		field.Uint64("category"),
		field.String("word").Unique(),
	}
}

func (KeywordWeak) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("category"),
	}
}

func (KeywordWeak) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "t_keyword_weak",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		},
	}
}
