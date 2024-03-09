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
		field.Int32("id").Immutable(),
		field.String("origin_short_id").NotEmpty().Immutable(),
		field.Bool("is_chinese").Default(false).Immutable(),
		field.String("origin_type").Immutable(),
		field.String("url").MaxLen(768).NotEmpty().Unique().Immutable(),
		field.String("title_chinese").Default(""),
		field.String("title_english").Default(""),
		field.String("author").Default(""),
		field.Strings("tags").Default([]string{}),
		field.Time("published_at").Immutable().SchemaType(map[string]string{dialect.MySQL: "datetime"}),
		field.Text("html_chinese").Default(""),
		field.Text("html_english").Default(""),
		field.Text("text_chinese").Default(""),
		field.Text("text_english").Default(""),
		field.Bool("is_china_related").Default(false),
		field.Strings("china_related_keywords").Default([]string{}),
		field.Bool("is_china_strong_related").Default(false),
		field.String("china_related_category").Default(""),
		field.Text("summary_chinese").Default(""),
	}
}

func (Article) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_chinese"),
		index.Fields("title_chinese"),
		index.Fields("title_english"),
		index.Fields("origin_short_id"),
		index.Fields("origin_type"),
		index.Fields("is_china_related"),
		index.Fields("china_related_keywords"),
		index.Fields("is_china_strong_related"),
		index.Fields("china_related_category"),
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
