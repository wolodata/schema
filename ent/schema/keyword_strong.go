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
		field.String("main"),
		field.Uint64("main_count"),
		field.String("sub").Default(""),
		field.Uint64("sub_count").Default(0),
	}
}

func (KeywordStrong) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("category"),
		index.Fields("main", "sub").Unique(),
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
