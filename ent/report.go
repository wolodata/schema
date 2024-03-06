// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/wolodata/schema/ent/report"
)

// Report is the model entity for the Report schema.
type Report struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// ReportType holds the value of the "report_type" field.
	ReportType string `json:"report_type,omitempty"`
	// TriggerUserID holds the value of the "trigger_user_id" field.
	TriggerUserID int32 `json:"trigger_user_id,omitempty"`
	// TriggerAt holds the value of the "trigger_at" field.
	TriggerAt time.Time `json:"trigger_at,omitempty"`
	// RelatedArticleIds holds the value of the "related_article_ids" field.
	RelatedArticleIds []int32 `json:"related_article_ids,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Reason holds the value of the "reason" field.
	Reason string `json:"reason,omitempty"`
	// GeneratedAt holds the value of the "generated_at" field.
	GeneratedAt  time.Time `json:"generated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Report) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case report.FieldRelatedArticleIds:
			values[i] = new([]byte)
		case report.FieldID, report.FieldTriggerUserID:
			values[i] = new(sql.NullInt64)
		case report.FieldReportType, report.FieldContent, report.FieldReason:
			values[i] = new(sql.NullString)
		case report.FieldTriggerAt, report.FieldGeneratedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Report fields.
func (r *Report) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case report.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int32(value.Int64)
		case report.FieldReportType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field report_type", values[i])
			} else if value.Valid {
				r.ReportType = value.String
			}
		case report.FieldTriggerUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field trigger_user_id", values[i])
			} else if value.Valid {
				r.TriggerUserID = int32(value.Int64)
			}
		case report.FieldTriggerAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field trigger_at", values[i])
			} else if value.Valid {
				r.TriggerAt = value.Time
			}
		case report.FieldRelatedArticleIds:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field related_article_ids", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.RelatedArticleIds); err != nil {
					return fmt.Errorf("unmarshal field related_article_ids: %w", err)
				}
			}
		case report.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				r.Content = value.String
			}
		case report.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				r.Reason = value.String
			}
		case report.FieldGeneratedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field generated_at", values[i])
			} else if value.Valid {
				r.GeneratedAt = value.Time
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Report.
// This includes values selected through modifiers, order, etc.
func (r *Report) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// Update returns a builder for updating this Report.
// Note that you need to call Report.Unwrap() before calling this method if this Report
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Report) Update() *ReportUpdateOne {
	return NewReportClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Report entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Report) Unwrap() *Report {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Report is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Report) String() string {
	var builder strings.Builder
	builder.WriteString("Report(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("report_type=")
	builder.WriteString(r.ReportType)
	builder.WriteString(", ")
	builder.WriteString("trigger_user_id=")
	builder.WriteString(fmt.Sprintf("%v", r.TriggerUserID))
	builder.WriteString(", ")
	builder.WriteString("trigger_at=")
	builder.WriteString(r.TriggerAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("related_article_ids=")
	builder.WriteString(fmt.Sprintf("%v", r.RelatedArticleIds))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(r.Content)
	builder.WriteString(", ")
	builder.WriteString("reason=")
	builder.WriteString(r.Reason)
	builder.WriteString(", ")
	builder.WriteString("generated_at=")
	builder.WriteString(r.GeneratedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Reports is a parsable slice of Report.
type Reports []*Report
