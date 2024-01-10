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
		field.String("origin_short_id").NotEmpty().Immutable(),
		field.Bool("is_chinese").Default(false).Immutable(),
		field.String("origin_type").Immutable(),
		field.String("url").MaxLen(768).NotEmpty().Unique().Immutable(),
		field.String("title_chinese").Optional(),
		field.String("title_english").Optional(),
		field.String("author").Immutable().Optional(),
		field.Strings("tags").Immutable().Optional(),
		field.Time("published_at").Immutable().SchemaType(map[string]string{dialect.MySQL: "datetime"}),
		field.Text("html_chinese").Optional(),
		field.Text("html_english").Optional(),
		field.Text("text_chinese").Optional(),
		field.Text("text_english").Optional(),
		field.Time("crawled_at").Immutable().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")),
		field.Text("summary_chinese").Optional(),
	}
}

func (Article) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_chinese"),
		index.Fields("title_chinese"),
		index.Fields("title_english"),
		index.Fields("origin_short_id"),
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
