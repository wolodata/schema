// Code generated by ent, DO NOT EDIT.

package promotconfig

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the promotconfig type in the database.
	Label = "promot_config"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldAPIModel holds the string denoting the api_model field in the database.
	FieldAPIModel = "api_model"
	// FieldAPIURL holds the string denoting the api_url field in the database.
	FieldAPIURL = "api_url"
	// FieldAPIKey holds the string denoting the api_key field in the database.
	FieldAPIKey = "api_key"
	// FieldPromptSystem holds the string denoting the prompt_system field in the database.
	FieldPromptSystem = "prompt_system"
	// FieldPromptUser holds the string denoting the prompt_user field in the database.
	FieldPromptUser = "prompt_user"
	// Table holds the table name of the promotconfig in the database.
	Table = "t_prompt_config"
)

// Columns holds all SQL columns for promotconfig fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldAPIModel,
	FieldAPIURL,
	FieldAPIKey,
	FieldPromptSystem,
	FieldPromptUser,
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
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultAPIModel holds the default value on creation for the "api_model" field.
	DefaultAPIModel string
	// DefaultAPIURL holds the default value on creation for the "api_url" field.
	DefaultAPIURL string
	// DefaultAPIKey holds the default value on creation for the "api_key" field.
	DefaultAPIKey string
	// DefaultPromptSystem holds the default value on creation for the "prompt_system" field.
	DefaultPromptSystem string
	// DefaultPromptUser holds the default value on creation for the "prompt_user" field.
	DefaultPromptUser string
)

// OrderOption defines the ordering options for the PromotConfig queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByAPIModel orders the results by the api_model field.
func ByAPIModel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAPIModel, opts...).ToFunc()
}

// ByAPIURL orders the results by the api_url field.
func ByAPIURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAPIURL, opts...).ToFunc()
}

// ByAPIKey orders the results by the api_key field.
func ByAPIKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAPIKey, opts...).ToFunc()
}

// ByPromptSystem orders the results by the prompt_system field.
func ByPromptSystem(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPromptSystem, opts...).ToFunc()
}

// ByPromptUser orders the results by the prompt_user field.
func ByPromptUser(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPromptUser, opts...).ToFunc()
}
