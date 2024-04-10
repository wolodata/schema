// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/wolodata/schema/ent/keywordweak"
)

// KeywordWeak is the model entity for the KeywordWeak schema.
type KeywordWeak struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Category holds the value of the "category" field.
	Category uint64 `json:"category,omitempty"`
	// Word holds the value of the "word" field.
	Word         string `json:"word,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*KeywordWeak) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case keywordweak.FieldCategory:
			values[i] = new(sql.NullInt64)
		case keywordweak.FieldID, keywordweak.FieldWord:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the KeywordWeak fields.
func (kw *KeywordWeak) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case keywordweak.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				kw.ID = value.String
			}
		case keywordweak.FieldCategory:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				kw.Category = uint64(value.Int64)
			}
		case keywordweak.FieldWord:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field word", values[i])
			} else if value.Valid {
				kw.Word = value.String
			}
		default:
			kw.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the KeywordWeak.
// This includes values selected through modifiers, order, etc.
func (kw *KeywordWeak) Value(name string) (ent.Value, error) {
	return kw.selectValues.Get(name)
}

// Update returns a builder for updating this KeywordWeak.
// Note that you need to call KeywordWeak.Unwrap() before calling this method if this KeywordWeak
// was returned from a transaction, and the transaction was committed or rolled back.
func (kw *KeywordWeak) Update() *KeywordWeakUpdateOne {
	return NewKeywordWeakClient(kw.config).UpdateOne(kw)
}

// Unwrap unwraps the KeywordWeak entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (kw *KeywordWeak) Unwrap() *KeywordWeak {
	_tx, ok := kw.config.driver.(*txDriver)
	if !ok {
		panic("ent: KeywordWeak is not a transactional entity")
	}
	kw.config.driver = _tx.drv
	return kw
}

// String implements the fmt.Stringer.
func (kw *KeywordWeak) String() string {
	var builder strings.Builder
	builder.WriteString("KeywordWeak(")
	builder.WriteString(fmt.Sprintf("id=%v, ", kw.ID))
	builder.WriteString("category=")
	builder.WriteString(fmt.Sprintf("%v", kw.Category))
	builder.WriteString(", ")
	builder.WriteString("word=")
	builder.WriteString(kw.Word)
	builder.WriteByte(')')
	return builder.String()
}

// KeywordWeaks is a parsable slice of KeywordWeak.
type KeywordWeaks []*KeywordWeak