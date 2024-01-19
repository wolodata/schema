// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/wolodata/schema/ent/topic"
)

// Topic is the model entity for the Topic schema.
type Topic struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Keyword holds the value of the "keyword" field.
	Keyword string `json:"keyword,omitempty"`
	// FollowTitle holds the value of the "follow_title" field.
	FollowTitle bool `json:"follow_title,omitempty"`
	// FollowContent holds the value of the "follow_content" field.
	FollowContent bool `json:"follow_content,omitempty"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Topic) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case topic.FieldFollowTitle, topic.FieldFollowContent:
			values[i] = new(sql.NullBool)
		case topic.FieldID, topic.FieldUserID:
			values[i] = new(sql.NullInt64)
		case topic.FieldKeyword:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Topic fields.
func (t *Topic) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case topic.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case topic.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				t.UserID = int(value.Int64)
			}
		case topic.FieldKeyword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field keyword", values[i])
			} else if value.Valid {
				t.Keyword = value.String
			}
		case topic.FieldFollowTitle:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field follow_title", values[i])
			} else if value.Valid {
				t.FollowTitle = value.Bool
			}
		case topic.FieldFollowContent:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field follow_content", values[i])
			} else if value.Valid {
				t.FollowContent = value.Bool
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Topic.
// This includes values selected through modifiers, order, etc.
func (t *Topic) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Topic.
// Note that you need to call Topic.Unwrap() before calling this method if this Topic
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Topic) Update() *TopicUpdateOne {
	return NewTopicClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Topic entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Topic) Unwrap() *Topic {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Topic is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Topic) String() string {
	var builder strings.Builder
	builder.WriteString("Topic(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", t.UserID))
	builder.WriteString(", ")
	builder.WriteString("keyword=")
	builder.WriteString(t.Keyword)
	builder.WriteString(", ")
	builder.WriteString("follow_title=")
	builder.WriteString(fmt.Sprintf("%v", t.FollowTitle))
	builder.WriteString(", ")
	builder.WriteString("follow_content=")
	builder.WriteString(fmt.Sprintf("%v", t.FollowContent))
	builder.WriteByte(')')
	return builder.String()
}

// Topics is a parsable slice of Topic.
type Topics []*Topic