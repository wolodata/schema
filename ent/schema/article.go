package schema

import (
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
		field.String("id").Immutable(),
		field.String("origin_short_id").NotEmpty().Immutable(),
		field.Bool("is_chinese").Default(false).Immutable(),
		field.String("url").MaxLen(768).NotEmpty().Unique().Immutable(),
		field.String("title_chinese").Default(""),
		field.String("title_english").Default(""),
		field.Strings("author").Default([]string{}),
		field.Time("published_at").Immutable().SchemaType(map[string]string{dialect.MySQL: "datetime"}),
		field.Text("html_chinese").Default(""),
		field.Text("html_english").Default(""),
		field.Text("text_chinese").Default(""),
		field.Text("text_english").Default(""),
		field.Strings("images").Default([]string{}),
		field.Strings("weak_keyword_ids").Default([]string{}),
		field.String("strong_keyword_id").Default(""),
		field.Text("summary_chinese").Default(""),
	}
}

func (Article) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_chinese"),
		index.Fields("title_chinese"),
		index.Fields("title_english"),
		index.Fields("origin_short_id"),
		index.Fields("weak_keyword_ids"),
		index.Fields("strong_keyword_id"),
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
