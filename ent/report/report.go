// Code generated by ent, DO NOT EDIT.

package report

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the report type in the database.
	Label = "report"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldReportType holds the string denoting the report_type field in the database.
	FieldReportType = "report_type"
	// FieldTriggerUserID holds the string denoting the trigger_user_id field in the database.
	FieldTriggerUserID = "trigger_user_id"
	// FieldTriggerAt holds the string denoting the trigger_at field in the database.
	FieldTriggerAt = "trigger_at"
	// FieldRelatedArticleIds holds the string denoting the related_article_ids field in the database.
	FieldRelatedArticleIds = "related_article_ids"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldReason holds the string denoting the reason field in the database.
	FieldReason = "reason"
	// FieldGeneratedAt holds the string denoting the generated_at field in the database.
	FieldGeneratedAt = "generated_at"
	// Table holds the table name of the report in the database.
	Table = "t_report"
)

// Columns holds all SQL columns for report fields.
var Columns = []string{
	FieldID,
	FieldReportType,
	FieldTriggerUserID,
	FieldTriggerAt,
	FieldRelatedArticleIds,
	FieldContent,
	FieldReason,
	FieldGeneratedAt,
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
	// ReportTypeValidator is a validator for the "report_type" field. It is called by the builders before save.
	ReportTypeValidator func(string) error
	// DefaultTriggerAt holds the default value on creation for the "trigger_at" field.
	DefaultTriggerAt func() time.Time
	// DefaultRelatedArticleIds holds the default value on creation for the "related_article_ids" field.
	DefaultRelatedArticleIds []int32
)

// OrderOption defines the ordering options for the Report queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByReportType orders the results by the report_type field.
func ByReportType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReportType, opts...).ToFunc()
}

// ByTriggerUserID orders the results by the trigger_user_id field.
func ByTriggerUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTriggerUserID, opts...).ToFunc()
}

// ByTriggerAt orders the results by the trigger_at field.
func ByTriggerAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTriggerAt, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByReason orders the results by the reason field.
func ByReason(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReason, opts...).ToFunc()
}

// ByGeneratedAt orders the results by the generated_at field.
func ByGeneratedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGeneratedAt, opts...).ToFunc()
}
