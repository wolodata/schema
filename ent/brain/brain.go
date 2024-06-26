// Code generated by ent, DO NOT EDIT.

package brain

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the brain type in the database.
	Label = "brain"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldQuestion holds the string denoting the question field in the database.
	FieldQuestion = "question"
	// FieldAnswer holds the string denoting the answer field in the database.
	FieldAnswer = "answer"
	// Table holds the table name of the brain in the database.
	Table = "t_brain"
)

// Columns holds all SQL columns for brain fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldQuestion,
	FieldAnswer,
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

// OrderOption defines the ordering options for the Brain queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByQuestion orders the results by the question field.
func ByQuestion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuestion, opts...).ToFunc()
}

// ByAnswer orders the results by the answer field.
func ByAnswer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAnswer, opts...).ToFunc()
}
