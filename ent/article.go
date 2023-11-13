// Code generated by ent, DO NOT EDIT.

package ent

import (
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
	ID int `json:"id,omitempty"`
	// OriginName holds the value of the "origin_name" field.
	OriginName string `json:"origin_name,omitempty"`
	// OriginType holds the value of the "origin_type" field.
	OriginType string `json:"origin_type,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// TitleChinese holds the value of the "title_chinese" field.
	TitleChinese string `json:"title_chinese,omitempty"`
	// Author holds the value of the "author" field.
	Author string `json:"author,omitempty"`
	// PublishedAt holds the value of the "published_at" field.
	PublishedAt time.Time `json:"published_at,omitempty"`
	// Raw holds the value of the "raw" field.
	Raw string `json:"raw,omitempty"`
	// CrawledAt holds the value of the "crawled_at" field.
	CrawledAt time.Time `json:"crawled_at,omitempty"`
	// SummaryChinese holds the value of the "summary_chinese" field.
	SummaryChinese string `json:"summary_chinese,omitempty"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Article) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case article.FieldID:
			values[i] = new(sql.NullInt64)
		case article.FieldOriginName, article.FieldOriginType, article.FieldURL, article.FieldTitle, article.FieldTitleChinese, article.FieldAuthor, article.FieldRaw, article.FieldSummaryChinese:
			values[i] = new(sql.NullString)
		case article.FieldPublishedAt, article.FieldCrawledAt:
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
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case article.FieldOriginName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field origin_name", values[i])
			} else if value.Valid {
				a.OriginName = value.String
			}
		case article.FieldOriginType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field origin_type", values[i])
			} else if value.Valid {
				a.OriginType = value.String
			}
		case article.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				a.URL = value.String
			}
		case article.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case article.FieldTitleChinese:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title_chinese", values[i])
			} else if value.Valid {
				a.TitleChinese = value.String
			}
		case article.FieldAuthor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				a.Author = value.String
			}
		case article.FieldPublishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field published_at", values[i])
			} else if value.Valid {
				a.PublishedAt = value.Time
			}
		case article.FieldRaw:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field raw", values[i])
			} else if value.Valid {
				a.Raw = value.String
			}
		case article.FieldCrawledAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field crawled_at", values[i])
			} else if value.Valid {
				a.CrawledAt = value.Time
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
	builder.WriteString("origin_name=")
	builder.WriteString(a.OriginName)
	builder.WriteString(", ")
	builder.WriteString("origin_type=")
	builder.WriteString(a.OriginType)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(a.URL)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(a.Title)
	builder.WriteString(", ")
	builder.WriteString("title_chinese=")
	builder.WriteString(a.TitleChinese)
	builder.WriteString(", ")
	builder.WriteString("author=")
	builder.WriteString(a.Author)
	builder.WriteString(", ")
	builder.WriteString("published_at=")
	builder.WriteString(a.PublishedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("raw=")
	builder.WriteString(a.Raw)
	builder.WriteString(", ")
	builder.WriteString("crawled_at=")
	builder.WriteString(a.CrawledAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("summary_chinese=")
	builder.WriteString(a.SummaryChinese)
	builder.WriteByte(')')
	return builder.String()
}

// Articles is a parsable slice of Article.
type Articles []*Article
