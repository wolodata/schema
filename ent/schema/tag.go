package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Tag struct {
	ent.Schema
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("english").NotEmpty().Immutable().Unique(),
		field.String("chinese").Default(""),
	}
}

func (Tag) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Tag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "t_tag",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		},
	}
}
