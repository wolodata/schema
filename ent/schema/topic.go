package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Topic struct {
	ent.Schema
}

func (Topic) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable(),
		field.Int32("user_id").Immutable(),
		field.String("keyword").NotEmpty(),
		field.Bool("follow_title").Default(false),
		field.Bool("follow_content").Default(false),
	}
}

func (Topic) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "keyword").Unique(),
		index.Fields("user_id"),
		index.Fields("keyword"),
	}
}

func (Topic) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "t_topic",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		},
	}
}
