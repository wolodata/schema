// Code generated by ent, DO NOT EDIT.

package article

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the article type in the database.
	Label = "article"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOriginShortID holds the string denoting the origin_short_id field in the database.
	FieldOriginShortID = "origin_short_id"
	// FieldIsChinese holds the string denoting the is_chinese field in the database.
	FieldIsChinese = "is_chinese"
	// FieldOriginType holds the string denoting the origin_type field in the database.
	FieldOriginType = "origin_type"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldTitleChinese holds the string denoting the title_chinese field in the database.
	FieldTitleChinese = "title_chinese"
	// FieldTitleEnglish holds the string denoting the title_english field in the database.
	FieldTitleEnglish = "title_english"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldPublishedAt holds the string denoting the published_at field in the database.
	FieldPublishedAt = "published_at"
	// FieldHTMLChinese holds the string denoting the html_chinese field in the database.
	FieldHTMLChinese = "html_chinese"
	// FieldHTMLEnglish holds the string denoting the html_english field in the database.
	FieldHTMLEnglish = "html_english"
	// FieldTextChinese holds the string denoting the text_chinese field in the database.
	FieldTextChinese = "text_chinese"
	// FieldTextEnglish holds the string denoting the text_english field in the database.
	FieldTextEnglish = "text_english"
	// FieldSummaryChinese holds the string denoting the summary_chinese field in the database.
	FieldSummaryChinese = "summary_chinese"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// Table holds the table name of the article in the database.
	Table = "t_article"
)

// Columns holds all SQL columns for article fields.
var Columns = []string{
	FieldID,
	FieldOriginShortID,
	FieldIsChinese,
	FieldOriginType,
	FieldURL,
	FieldTitleChinese,
	FieldTitleEnglish,
	FieldAuthor,
	FieldTags,
	FieldPublishedAt,
	FieldHTMLChinese,
	FieldHTMLEnglish,
	FieldTextChinese,
	FieldTextEnglish,
	FieldSummaryChinese,
	FieldCategory,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// OriginShortIDValidator is a validator for the "origin_short_id" field. It is called by the builders before save.
	OriginShortIDValidator func(string) error
	// DefaultIsChinese holds the default value on creation for the "is_chinese" field.
	DefaultIsChinese bool
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
	// DefaultTitleChinese holds the default value on creation for the "title_chinese" field.
	DefaultTitleChinese string
	// DefaultTitleEnglish holds the default value on creation for the "title_english" field.
	DefaultTitleEnglish string
	// DefaultAuthor holds the default value on creation for the "author" field.
	DefaultAuthor string
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// DefaultHTMLChinese holds the default value on creation for the "html_chinese" field.
	DefaultHTMLChinese string
	// DefaultHTMLEnglish holds the default value on creation for the "html_english" field.
	DefaultHTMLEnglish string
	// DefaultTextChinese holds the default value on creation for the "text_chinese" field.
	DefaultTextChinese string
	// DefaultTextEnglish holds the default value on creation for the "text_english" field.
	DefaultTextEnglish string
	// DefaultSummaryChinese holds the default value on creation for the "summary_chinese" field.
	DefaultSummaryChinese string
	// DefaultCategory holds the default value on creation for the "category" field.
	DefaultCategory string
)

// OrderOption defines the ordering options for the Article queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOriginShortID orders the results by the origin_short_id field.
func ByOriginShortID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOriginShortID, opts...).ToFunc()
}

// ByIsChinese orders the results by the is_chinese field.
func ByIsChinese(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsChinese, opts...).ToFunc()
}

// ByOriginType orders the results by the origin_type field.
func ByOriginType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOriginType, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByTitleChinese orders the results by the title_chinese field.
func ByTitleChinese(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitleChinese, opts...).ToFunc()
}

// ByTitleEnglish orders the results by the title_english field.
func ByTitleEnglish(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitleEnglish, opts...).ToFunc()
}

// ByAuthor orders the results by the author field.
func ByAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthor, opts...).ToFunc()
}

// ByPublishedAt orders the results by the published_at field.
func ByPublishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublishedAt, opts...).ToFunc()
}

// ByHTMLChinese orders the results by the html_chinese field.
func ByHTMLChinese(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHTMLChinese, opts...).ToFunc()
}

// ByHTMLEnglish orders the results by the html_english field.
func ByHTMLEnglish(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHTMLEnglish, opts...).ToFunc()
}

// ByTextChinese orders the results by the text_chinese field.
func ByTextChinese(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTextChinese, opts...).ToFunc()
}

// ByTextEnglish orders the results by the text_english field.
func ByTextEnglish(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTextEnglish, opts...).ToFunc()
}

// BySummaryChinese orders the results by the summary_chinese field.
func BySummaryChinese(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSummaryChinese, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}
