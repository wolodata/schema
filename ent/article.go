// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/wolodata/schema/ent/article"
)

// Article is the model entity for the Article schema.
type Article struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// OriginShortID holds the value of the "origin_short_id" field.
	OriginShortID string `json:"origin_short_id,omitempty"`
	// IsChinese holds the value of the "is_chinese" field.
	IsChinese bool `json:"is_chinese,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// TitleChinese holds the value of the "title_chinese" field.
	TitleChinese string `json:"title_chinese,omitempty"`
	// TitleEnglish holds the value of the "title_english" field.
	TitleEnglish string `json:"title_english,omitempty"`
	// Author holds the value of the "author" field.
	Author []string `json:"author,omitempty"`
	// PublishedAt holds the value of the "published_at" field.
	PublishedAt time.Time `json:"published_at,omitempty"`
	// HTMLChinese holds the value of the "html_chinese" field.
	HTMLChinese string `json:"html_chinese,omitempty"`
	// HTMLEnglish holds the value of the "html_english" field.
	HTMLEnglish string `json:"html_english,omitempty"`
	// TextChinese holds the value of the "text_chinese" field.
	TextChinese string `json:"text_chinese,omitempty"`
	// TextEnglish holds the value of the "text_english" field.
	TextEnglish string `json:"text_english,omitempty"`
	// IsChinaRelated holds the value of the "is_china_related" field.
	IsChinaRelated bool `json:"is_china_related,omitempty"`
	// ChinaRelatedKeywords holds the value of the "china_related_keywords" field.
	ChinaRelatedKeywords []string `json:"china_related_keywords,omitempty"`
	// IsChinaStrongRelated holds the value of the "is_china_strong_related" field.
	IsChinaStrongRelated bool `json:"is_china_strong_related,omitempty"`
	// ChinaRelatedCategory holds the value of the "china_related_category" field.
	ChinaRelatedCategory string `json:"china_related_category,omitempty"`
	// SummaryChinese holds the value of the "summary_chinese" field.
	SummaryChinese string `json:"summary_chinese,omitempty"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Article) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case article.FieldAuthor, article.FieldChinaRelatedKeywords:
			values[i] = new([]byte)
		case article.FieldIsChinese, article.FieldIsChinaRelated, article.FieldIsChinaStrongRelated:
			values[i] = new(sql.NullBool)
		case article.FieldID, article.FieldOriginShortID, article.FieldURL, article.FieldTitleChinese, article.FieldTitleEnglish, article.FieldHTMLChinese, article.FieldHTMLEnglish, article.FieldTextChinese, article.FieldTextEnglish, article.FieldChinaRelatedCategory, article.FieldSummaryChinese:
			values[i] = new(sql.NullString)
		case article.FieldPublishedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Article fields.
func (a *Article) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case article.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				a.ID = value.String
			}
		case article.FieldOriginShortID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field origin_short_id", values[i])
			} else if value.Valid {
				a.OriginShortID = value.String
			}
		case article.FieldIsChinese:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_chinese", values[i])
			} else if value.Valid {
				a.IsChinese = value.Bool
			}
		case article.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				a.URL = value.String
			}
		case article.FieldTitleChinese:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title_chinese", values[i])
			} else if value.Valid {
				a.TitleChinese = value.String
			}
		case article.FieldTitleEnglish:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title_english", values[i])
			} else if value.Valid {
				a.TitleEnglish = value.String
			}
		case article.FieldAuthor:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.Author); err != nil {
					return fmt.Errorf("unmarshal field author: %w", err)
				}
			}
		case article.FieldPublishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field published_at", values[i])
			} else if value.Valid {
				a.PublishedAt = value.Time
			}
		case article.FieldHTMLChinese:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field html_chinese", values[i])
			} else if value.Valid {
				a.HTMLChinese = value.String
			}
		case article.FieldHTMLEnglish:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field html_english", values[i])
			} else if value.Valid {
				a.HTMLEnglish = value.String
			}
		case article.FieldTextChinese:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text_chinese", values[i])
			} else if value.Valid {
				a.TextChinese = value.String
			}
		case article.FieldTextEnglish:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text_english", values[i])
			} else if value.Valid {
				a.TextEnglish = value.String
			}
		case article.FieldIsChinaRelated:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_china_related", values[i])
			} else if value.Valid {
				a.IsChinaRelated = value.Bool
			}
		case article.FieldChinaRelatedKeywords:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field china_related_keywords", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.ChinaRelatedKeywords); err != nil {
					return fmt.Errorf("unmarshal field china_related_keywords: %w", err)
				}
			}
		case article.FieldIsChinaStrongRelated:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_china_strong_related", values[i])
			} else if value.Valid {
				a.IsChinaStrongRelated = value.Bool
			}
		case article.FieldChinaRelatedCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field china_related_category", values[i])
			} else if value.Valid {
				a.ChinaRelatedCategory = value.String
			}
		case article.FieldSummaryChinese:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field summary_chinese", values[i])
			} else if value.Valid {
				a.SummaryChinese = value.String
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Article.
// This includes values selected through modifiers, order, etc.
func (a *Article) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this Article.
// Note that you need to call Article.Unwrap() before calling this method if this Article
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Article) Update() *ArticleUpdateOne {
	return NewArticleClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Article entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Article) Unwrap() *Article {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Article is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Article) String() string {
	var builder strings.Builder
	builder.WriteString("Article(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("origin_short_id=")
	builder.WriteString(a.OriginShortID)
	builder.WriteString(", ")
	builder.WriteString("is_chinese=")
	builder.WriteString(fmt.Sprintf("%v", a.IsChinese))
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(a.URL)
	builder.WriteString(", ")
	builder.WriteString("title_chinese=")
	builder.WriteString(a.TitleChinese)
	builder.WriteString(", ")
	builder.WriteString("title_english=")
	builder.WriteString(a.TitleEnglish)
	builder.WriteString(", ")
	builder.WriteString("author=")
	builder.WriteString(fmt.Sprintf("%v", a.Author))
	builder.WriteString(", ")
	builder.WriteString("published_at=")
	builder.WriteString(a.PublishedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("html_chinese=")
	builder.WriteString(a.HTMLChinese)
	builder.WriteString(", ")
	builder.WriteString("html_english=")
	builder.WriteString(a.HTMLEnglish)
	builder.WriteString(", ")
	builder.WriteString("text_chinese=")
	builder.WriteString(a.TextChinese)
	builder.WriteString(", ")
	builder.WriteString("text_english=")
	builder.WriteString(a.TextEnglish)
	builder.WriteString(", ")
	builder.WriteString("is_china_related=")
	builder.WriteString(fmt.Sprintf("%v", a.IsChinaRelated))
	builder.WriteString(", ")
	builder.WriteString("china_related_keywords=")
	builder.WriteString(fmt.Sprintf("%v", a.ChinaRelatedKeywords))
	builder.WriteString(", ")
	builder.WriteString("is_china_strong_related=")
	builder.WriteString(fmt.Sprintf("%v", a.IsChinaStrongRelated))
	builder.WriteString(", ")
	builder.WriteString("china_related_category=")
	builder.WriteString(a.ChinaRelatedCategory)
	builder.WriteString(", ")
	builder.WriteString("summary_chinese=")
	builder.WriteString(a.SummaryChinese)
	builder.WriteByte(')')
	return builder.String()
}

// Articles is a parsable slice of Article.
type Articles []*Article
