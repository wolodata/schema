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
		field.Time("published_at").Immutable().SchemaType(map[string]string{dialect.MySQL: "datetime"}),
		field.Text("html_chinese").Default(""),
		field.Text("html_english").Default(""),
		field.Text("text_chinese").Default(""),
		field.Text("text_english").Default(""),

		// 图片
		field.Strings("images").Default([]string{}),
		field.Bool("image_uploaded").Default(false),

		// 弱关键词 可能命中多个
		field.Bool("weak_keyword_processed").Default(false),
		field.Bool("weak_keyword_related").Default(false),
		field.JSON("weak_keywords", []WeakKeyword{}).Default([]WeakKeyword{}),

		// 强关键词 最多命中一个
		field.Bool("strong_keyword_processed").Default(false),
		field.Bool("strong_keyword_related").Default(false),
		field.JSON("strong_keyword", StrongKeyword{}).Optional(),

		// 是否处理过强相关
		field.Bool("strong_related_processed").Default(false),
		// 是否强相关
		field.Bool("strong_related").Default(false),

		// 是否处理过所属范畴
		field.Bool("strong_related_category_processed").Default(false),
		// 所属范畴
		field.String("strong_related_category").Default(""),

		// 中文总结
		field.Text("summary_chinese").Default(""),

		// 是否处理过强相关内容的中文总结
		field.Bool("strong_related_summary_chinese_processed").Default(false),
		// 强相关内容的中文总结
		field.Text("strong_related_summary_chinese").Default(""),
	}
}

type WeakKeyword struct {
	Word     string
	Category uint64
}

type StrongKeyword struct {
	Main      string
	MainCount uint64
	Sub       string
	SubCount  uint64
}

func (Article) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_chinese"),
		index.Fields("title_chinese"),
		index.Fields("title_english"),
		index.Fields("origin_short_id"),
		index.Fields("published_at"),
		index.Fields("weak_keyword_processed"),
		index.Fields("weak_keyword_related"),
		index.Fields("strong_keyword_processed"),
		index.Fields("strong_keyword_related"),
		// 是否处理过强相关
		index.Fields("strong_related_processed"),
		// 是否强相关
		index.Fields("strong_related"),
		// 是否处理过所属范畴
		index.Fields("strong_related_category_processed"),
		// 所属范畴
		index.Fields("strong_related_category"),
		// 是否处理过强相关内容的中文总结
		index.Fields("strong_related_summary_chinese_processed"),
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
