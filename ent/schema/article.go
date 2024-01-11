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
		field.Bool("is_chinese").Immutable(),
		field.String("origin_type").Immutable(),
		field.String("url").MaxLen(768).NotEmpty().Unique().Immutable(),
		field.String("title_chinese"),
		field.String("title_english"),
		field.String("author"),
		field.Strings("tags").Default([]string{}),
		field.Time("published_at").Immutable().SchemaType(map[string]string{dialect.MySQL: "datetime"}),
		field.Text("html_chinese"),
		field.Text("html_english"),
		field.Text("text_chinese"),
		field.Text("text_english"),
		field.Time("crawled_at").Immutable().SchemaType(map[string]string{dialect.MySQL: "datetime"}).Default(time.Now).Annotations(entsql.Default("CURRENT_TIMESTAMP")),
		field.Text("summary_chinese"),
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
