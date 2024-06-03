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
		{Name: "id", Type: field.TypeString},
		{Name: "origin_short_id", Type: field.TypeString},
		{Name: "is_chinese", Type: field.TypeBool, Default: false},
		{Name: "url", Type: field.TypeString, Unique: true, Size: 768},
		{Name: "title_chinese", Type: field.TypeString, Default: ""},
		{Name: "title_english", Type: field.TypeString, Default: ""},
		{Name: "published_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "html_chinese", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "html_english", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "text_chinese", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "text_english", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "images", Type: field.TypeJSON},
		{Name: "image_uploaded", Type: field.TypeBool, Default: false},
		{Name: "weak_keyword_processed", Type: field.TypeBool, Default: false},
		{Name: "weak_keyword_related", Type: field.TypeBool, Default: false},
		{Name: "weak_keywords", Type: field.TypeJSON},
		{Name: "strong_keyword_processed", Type: field.TypeBool, Default: false},
		{Name: "strong_keyword_related", Type: field.TypeBool, Default: false},
		{Name: "strong_keyword", Type: field.TypeJSON, Nullable: true},
		{Name: "strong_related_processed", Type: field.TypeBool, Default: false},
		{Name: "strong_related", Type: field.TypeBool, Default: false},
		{Name: "strong_related_category_processed", Type: field.TypeBool, Default: false},
		{Name: "strong_related_category", Type: field.TypeString, Default: ""},
		{Name: "summary_chinese", Type: field.TypeString, Size: 2147483647, Default: ""},
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
				Columns: []*schema.Column{TArticleColumns[4]},
			},
			{
				Name:    "article_title_english",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[5]},
			},
			{
				Name:    "article_origin_short_id",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[1]},
			},
			{
				Name:    "article_published_at",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[6]},
			},
			{
				Name:    "article_weak_keyword_processed",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[13]},
			},
			{
				Name:    "article_weak_keyword_related",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[14]},
			},
			{
				Name:    "article_strong_keyword_processed",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[16]},
			},
			{
				Name:    "article_strong_keyword_related",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[17]},
			},
			{
				Name:    "article_strong_related_processed",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[19]},
			},
			{
				Name:    "article_strong_related",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[20]},
			},
			{
				Name:    "article_strong_related_category_processed",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[21]},
			},
			{
				Name:    "article_strong_related_category",
				Unique:  false,
				Columns: []*schema.Column{TArticleColumns[22]},
			},
		},
	}
	// TBrainColumns holds the columns for the "t_brain" table.
	TBrainColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
		{Name: "question", Type: field.TypeString, Size: 2147483647},
		{Name: "answer", Type: field.TypeString, Size: 2147483647},
	}
	// TBrainTable holds the schema information for the "t_brain" table.
	TBrainTable = &schema.Table{
		Name:       "t_brain",
		Columns:    TBrainColumns,
		PrimaryKey: []*schema.Column{TBrainColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "brain_user_id",
				Unique:  false,
				Columns: []*schema.Column{TBrainColumns[1]},
			},
		},
	}
	// THTMLColumns holds the columns for the "t_html" table.
	THTMLColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "origin_short_id", Type: field.TypeString},
		{Name: "is_chinese", Type: field.TypeBool, Default: false},
		{Name: "url", Type: field.TypeString, Unique: true, Size: 768},
		{Name: "html", Type: field.TypeString, Size: 2147483647},
		{Name: "crawled_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "parsed_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "analyzed_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "reason", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
	}
	// THTMLTable holds the schema information for the "t_html" table.
	THTMLTable = &schema.Table{
		Name:       "t_html",
		Columns:    THTMLColumns,
		PrimaryKey: []*schema.Column{THTMLColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "html_origin_short_id",
				Unique:  false,
				Columns: []*schema.Column{THTMLColumns[1]},
			},
			{
				Name:    "html_is_chinese",
				Unique:  false,
				Columns: []*schema.Column{THTMLColumns[2]},
			},
			{
				Name:    "html_crawled_at",
				Unique:  false,
				Columns: []*schema.Column{THTMLColumns[5]},
			},
			{
				Name:    "html_parsed_at",
				Unique:  false,
				Columns: []*schema.Column{THTMLColumns[6]},
			},
			{
				Name:    "html_analyzed_at",
				Unique:  false,
				Columns: []*schema.Column{THTMLColumns[7]},
			},
		},
	}
	// TKeywordStrongColumns holds the columns for the "t_keyword_strong" table.
	TKeywordStrongColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "category", Type: field.TypeUint64},
		{Name: "main", Type: field.TypeString},
		{Name: "main_count", Type: field.TypeUint64},
		{Name: "sub", Type: field.TypeString, Default: ""},
		{Name: "sub_count", Type: field.TypeUint64, Default: 0},
	}
	// TKeywordStrongTable holds the schema information for the "t_keyword_strong" table.
	TKeywordStrongTable = &schema.Table{
		Name:       "t_keyword_strong",
		Columns:    TKeywordStrongColumns,
		PrimaryKey: []*schema.Column{TKeywordStrongColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "keywordstrong_category",
				Unique:  false,
				Columns: []*schema.Column{TKeywordStrongColumns[1]},
			},
			{
				Name:    "keywordstrong_main_sub",
				Unique:  true,
				Columns: []*schema.Column{TKeywordStrongColumns[2], TKeywordStrongColumns[4]},
			},
		},
	}
	// TKeywordWeakColumns holds the columns for the "t_keyword_weak" table.
	TKeywordWeakColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "category", Type: field.TypeUint64},
		{Name: "word", Type: field.TypeString, Unique: true},
	}
	// TKeywordWeakTable holds the schema information for the "t_keyword_weak" table.
	TKeywordWeakTable = &schema.Table{
		Name:       "t_keyword_weak",
		Columns:    TKeywordWeakColumns,
		PrimaryKey: []*schema.Column{TKeywordWeakColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "keywordweak_category",
				Unique:  false,
				Columns: []*schema.Column{TKeywordWeakColumns[1]},
			},
		},
	}
	// TPromptConfigColumns holds the columns for the "t_prompt_config" table.
	TPromptConfigColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Default: ""},
		{Name: "api_model", Type: field.TypeString, Default: ""},
		{Name: "api_url", Type: field.TypeString, Default: ""},
		{Name: "api_key", Type: field.TypeString, Default: ""},
		{Name: "prompt_system", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "prompt_user", Type: field.TypeString, Size: 2147483647, Default: ""},
	}
	// TPromptConfigTable holds the schema information for the "t_prompt_config" table.
	TPromptConfigTable = &schema.Table{
		Name:       "t_prompt_config",
		Columns:    TPromptConfigColumns,
		PrimaryKey: []*schema.Column{TPromptConfigColumns[0]},
	}
	// TReportColumns holds the columns for the "t_report" table.
	TReportColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "report_type", Type: field.TypeString},
		{Name: "start_time", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "end_time", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "source_ids", Type: field.TypeJSON},
		{Name: "category", Type: field.TypeString, Default: ""},
		{Name: "article_ids", Type: field.TypeJSON},
		{Name: "trigger_user_id", Type: field.TypeString, Nullable: true},
		{Name: "trigger_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "content", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "generated_at", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP", SchemaType: map[string]string{"mysql": "datetime"}},
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
				Columns: []*schema.Column{TReportColumns[7]},
			},
			{
				Name:    "report_trigger_at",
				Unique:  false,
				Columns: []*schema.Column{TReportColumns[8]},
			},
			{
				Name:    "report_generated_at",
				Unique:  false,
				Columns: []*schema.Column{TReportColumns[10]},
			},
		},
	}
	// TSystemConfigColumns holds the columns for the "t_system_config" table.
	TSystemConfigColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Default: ""},
		{Name: "value", Type: field.TypeString, Default: ""},
	}
	// TSystemConfigTable holds the schema information for the "t_system_config" table.
	TSystemConfigTable = &schema.Table{
		Name:       "t_system_config",
		Columns:    TSystemConfigColumns,
		PrimaryKey: []*schema.Column{TSystemConfigColumns[0]},
	}
	// TTopicColumns holds the columns for the "t_topic" table.
	TTopicColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
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
		{Name: "id", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "is_admin", Type: field.TypeBool, Default: false},
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
			{
				Name:    "user_is_admin",
				Unique:  false,
				Columns: []*schema.Column{TUserColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TArticleTable,
		TBrainTable,
		THTMLTable,
		TKeywordStrongTable,
		TKeywordWeakTable,
		TPromptConfigTable,
		TReportTable,
		TSystemConfigTable,
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
	TBrainTable.Annotation = &entsql.Annotation{
		Table:     "t_brain",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_general_ci",
	}
	THTMLTable.Annotation = &entsql.Annotation{
		Table:     "t_html",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_general_ci",
	}
	TKeywordStrongTable.Annotation = &entsql.Annotation{
		Table:     "t_keyword_strong",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_general_ci",
	}
	TKeywordWeakTable.Annotation = &entsql.Annotation{
		Table:     "t_keyword_weak",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_general_ci",
	}
	TPromptConfigTable.Annotation = &entsql.Annotation{
		Table:     "t_prompt_config",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_general_ci",
	}
	TReportTable.Annotation = &entsql.Annotation{
		Table:     "t_report",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_general_ci",
	}
	TSystemConfigTable.Annotation = &entsql.Annotation{
		Table:     "t_system_config",
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
