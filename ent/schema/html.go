package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Html struct {
	ent.Schema
}

func (Html) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Immutable(),
		field.String("origin_short_id").NotEmpty().Immutable(),
		field.Bool("is_chinese").Default(false).Immutable(),
		field.String("url").MaxLen(768).NotEmpty().Unique().Immutable(),
		field.Text("html").NotEmpty().Immutable(),
		field.Time("crawled_at").Immutable().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")),
	}
}

func (Html) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("origin_short_id"),
		index.Fields("is_chinese"),
		index.Fields("crawled_at"),
	}
}

func (Html) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "t_html",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		},
	}
}
