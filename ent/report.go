// Code generated by ent, DO NOT EDIT.

package ent

import (
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
	ID string `json:"id,omitempty"`
	// ReportType holds the value of the "report_type" field.
	ReportType string `json:"report_type,omitempty"`
	// StartTime holds the value of the "start_time" field.
	StartTime time.Time `json:"start_time,omitempty"`
	// EndTime holds the value of the "end_time" field.
	EndTime time.Time `json:"end_time,omitempty"`
	// TriggerUserID holds the value of the "trigger_user_id" field.
	TriggerUserID string `json:"trigger_user_id,omitempty"`
	// TriggerAt holds the value of the "trigger_at" field.
	TriggerAt time.Time `json:"trigger_at,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// GeneratedAt holds the value of the "generated_at" field.
	GeneratedAt  time.Time `json:"generated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Report) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case report.FieldID, report.FieldReportType, report.FieldTriggerUserID, report.FieldContent:
			values[i] = new(sql.NullString)
		case report.FieldStartTime, report.FieldEndTime, report.FieldTriggerAt, report.FieldGeneratedAt:
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
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				r.ID = value.String
			}
		case report.FieldReportType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field report_type", values[i])
			} else if value.Valid {
				r.ReportType = value.String
			}
		case report.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				r.StartTime = value.Time
			}
		case report.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_time", values[i])
			} else if value.Valid {
				r.EndTime = value.Time
			}
		case report.FieldTriggerUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field trigger_user_id", values[i])
			} else if value.Valid {
				r.TriggerUserID = value.String
			}
		case report.FieldTriggerAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field trigger_at", values[i])
			} else if value.Valid {
				r.TriggerAt = value.Time
			}
		case report.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				r.Content = value.String
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
	builder.WriteString("start_time=")
	builder.WriteString(r.StartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_time=")
	builder.WriteString(r.EndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("trigger_user_id=")
	builder.WriteString(r.TriggerUserID)
	builder.WriteString(", ")
	builder.WriteString("trigger_at=")
	builder.WriteString(r.TriggerAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(r.Content)
	builder.WriteString(", ")
	builder.WriteString("generated_at=")
	builder.WriteString(r.GeneratedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Reports is a parsable slice of Report.
type Reports []*Report
