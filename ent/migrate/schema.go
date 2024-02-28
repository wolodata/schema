// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TArticleColumns holds the columns for the "t_article" table.
	TArticleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "origin_short_id", Type: field.TypeString},
		{Name: "is_chinese", Type: field.TypeBool, Default: false},
		{Name: "origin_type", Type: field.TypeString},
		{Name: "url", Type: field.TypeString, Unique: true, Size: 768},
		{Name: "title_chinese", Type: field.TypeString, Default: ""},
		{Name: "title_english", Type: field.TypeString, Default: ""},
		{Name: "author", Type: field.TypeString, Default: ""},
		{Name: "tags", Type: field.TypeJSON},
		{Name: "published_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "html_chinese", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "html_english", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "text_chinese", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "text_english", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "crawled_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "summary_chinese", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "category", Type: field.TypeString, Default: ""},
	}
	// TArticleTable holds the schema information for the "t_article" table.
	TArticleTable = &schema.Table{
		Name:       "t_article",
		Columns:    TArticleColumns,
		PrimaryKey: []*schema.Column{TArticleColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "article_is_chinese",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[2]},
			},
			{
				Name:    "article_title_chinese",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[5]},
			},
			{
				Name:    "article_title_english",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[6]},
			},
			{
				Name:    "article_origin_short_id",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[1]},
			},
			{
				Name:    "article_origin_type",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[3]},
			},
			{
				Name:    "article_category",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[16]},
			},
		},
	}
	// TReportColumns holds the columns for the "t_report" table.
	TReportColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "report_type", Type: field.TypeString},
		{Name: "trigger_user_id", Type: field.TypeInt32, Nullable: true},
		{Name: "trigger_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "related_article_ids", Type: field.TypeJSON},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "reason", Type: field.TypeString, Size: 2147483647},
		{Name: "generated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// TReportTable holds the schema information for the "t_report" table.
	TReportTable = &schema.Table{
		Name:       "t_report",
		Columns:    TReportColumns,
		PrimaryKey: []*schema.Column{TReportColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "report_report_type",
				Unique:  false,
				Columns: []*schema.Column{TReportColumns[1]},
			},
			{
				Name:    "report_trigger_user_id",
				Unique:  false,
				Columns: []*schema.Column{TReportColumns[2]},
			},
			{
				Name:    "report_trigger_at",
				Unique:  false,
				Columns: []*schema.Column{TReportColumns[3]},
			},
			{
				Name:    "report_generated_at",
				Unique:  false,
				Columns: []*schema.Column{TReportColumns[7]},
			},
		},
	}
	// TTagColumns holds the columns for the "t_tag" table.
	TTagColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "english", Type: field.TypeString, Unique: true},
		{Name: "chinese", Type: field.TypeString, Default: ""},
	}
	// TTagTable holds the schema information for the "t_tag" table.
	TTagTable = &schema.Table{
		Name:       "t_tag",
		Columns:    TTagColumns,
		PrimaryKey: []*schema.Column{TTagColumns[0]},
	}
	// TTopicColumns holds the columns for the "t_topic" table.
	TTopicColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "user_id", Type: field.TypeInt32},
		{Name: "keyword", Type: field.TypeString},
		{Name: "follow_title", Type: field.TypeBool, Default: false},
		{Name: "follow_content", Type: field.TypeBool, Default: false},
	}
	// TTopicTable holds the schema information for the "t_topic" table.
	TTopicTable = &schema.Table{
		Name:       "t_topic",
		Columns:    TTopicColumns,
		PrimaryKey: []*schema.Column{TTopicColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "topic_user_id_keyword",
				Unique:  true,
				Columns: []*schema.Column{TTopicColumns[1], TTopicColumns[2]},
			},
			{
				Name:    "topic_user_id",
				Unique:  false,
				Columns: []*schema.Column{TTopicColumns[1]},
			},
			{
				Name:    "topic_keyword",
				Unique:  false,
				Columns: []*schema.Column{TTopicColumns[2]},
			},
		},
	}
	// TUserColumns holds the columns for the "t_user" table.
	TUserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// TUserTable holds the schema information for the "t_user" table.
	TUserTable = &schema.Table{
		Name:       "t_user",
		Columns:    TUserColumns,
		PrimaryKey: []*schema.Column{TUserColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_username",
				Unique:  false,
				Columns: []*schema.Column{TUserColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TArticleTable,
		TReportTable,
		TTagTable,
		TTopicTable,
		TUserTable,
	}
)

func init() {
	TArticleTable.Annotation = &entsql.Annotation{
		Table:     "t_article",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_general_ci",
	}
	TReportTable.Annotation = &entsql.Annotation{
		Table:     "t_report",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_general_ci",
	}
	TTagTable.Annotation = &entsql.Annotation{
		Table:     "t_tag",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_general_ci",
	}
	TTopicTable.Annotation = &entsql.Annotation{
		Table:     "t_topic",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_general_ci",
	}
	TUserTable.Annotation = &entsql.Annotation{
		Table:     "t_user",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_general_ci",
	}
}
