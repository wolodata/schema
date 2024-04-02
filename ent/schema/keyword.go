package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type Keyword struct {
	ent.Schema
}

func (Keyword) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable(),
		field.Uint64("category"),
		field.String("word").Unique(),
		field.Uint64("china_weak_related_count"),
		field.Uint64("china_strong_related_count"),
		field.String("sub_word").Default(""),
		field.Uint64("sub_word_count").Default(0),
		field.Time("updated_at").SchemaType(map[string]string{"mysql": "datetime"}).Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Keyword) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("category"),
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
