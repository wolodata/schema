// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/wolodata/schema/ent/brain"
)

// Brain is the model entity for the Brain schema.
type Brain struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Question holds the value of the "question" field.
	Question string `json:"question,omitempty"`
	// Answer holds the value of the "answer" field.
	Answer       string `json:"answer,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Brain) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case brain.FieldID, brain.FieldUserID, brain.FieldQuestion, brain.FieldAnswer:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Brain fields.
func (b *Brain) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case brain.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				b.ID = value.String
			}
		case brain.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				b.UserID = value.String
			}
		case brain.FieldQuestion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field question", values[i])
			} else if value.Valid {
				b.Question = value.String
			}
		case brain.FieldAnswer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field answer", values[i])
			} else if value.Valid {
				b.Answer = value.String
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Brain.
// This includes values selected through modifiers, order, etc.
func (b *Brain) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// Update returns a builder for updating this Brain.
// Note that you need to call Brain.Unwrap() before calling this method if this Brain
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Brain) Update() *BrainUpdateOne {
	return NewBrainClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Brain entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Brain) Unwrap() *Brain {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Brain is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Brain) String() string {
	var builder strings.Builder
	builder.WriteString("Brain(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("user_id=")
	builder.WriteString(b.UserID)
	builder.WriteString(", ")
	builder.WriteString("question=")
	builder.WriteString(b.Question)
	builder.WriteString(", ")
	builder.WriteString("answer=")
	builder.WriteString(b.Answer)
	builder.WriteByte(')')
	return builder.String()
}

// Brains is a parsable slice of Brain.
type Brains []*Brain
