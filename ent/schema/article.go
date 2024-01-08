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

type Article struct {
	ent.Schema
}

func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Immutable(),
		field.String("origin_name").Immutable(),
		field.String("origin_type").Immutable(),
		field.String("url").MaxLen(768).Unique().Immutable(),
		field.String("title_cn").Optional(),
		field.String("title_en").Optional(),
		field.String("author").Immutable().Optional(),
		field.Strings("tags").Immutable().Optional(),
		field.Time("published_at").Immutable().SchemaType(map[string]string{dialect.MySQL: "datetime"}),
		field.Text("html_cn").Optional(),
		field.Text("html_en").Optional(),
		field.Text("text_cn").Optional(),
		field.Text("text_en").Optional(),
		field.Time("crawled_at").Immutable().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")),
		field.Text("summary_cn").Optional(),
	}
}

func (Article) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title_cn"),
		index.Fields("title_en"),
		index.Fields("origin_name"),
		index.Fields("origin_type"),
	}
}

func (Article) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "t_article",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_general_ci",
		},
	}
}
