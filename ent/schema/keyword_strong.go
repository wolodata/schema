package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type KeywordStrong struct {
	ent.Schema
}

func (KeywordStrong) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable(),
		field.Uint64("category"),
		field.String("main").Unique(),
		field.Uint64("main_count").Default(1),
		field.String("sub").Default(""),
		field.Uint64("sub_count").Default(1),
	}
}

func (KeywordStrong) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("category"),
	}
}

func (KeywordStrong) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "t_keyword_strong",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		},
	}
}
