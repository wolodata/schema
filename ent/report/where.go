// Code generated by ent, DO NOT EDIT.

package report

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/wolodata/schema/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldID, id))
}

// ReportType applies equality check predicate on the "report_type" field. It's identical to ReportTypeEQ.
func ReportType(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldReportType, v))
}

// TriggerUserID applies equality check predicate on the "trigger_user_id" field. It's identical to TriggerUserIDEQ.
func TriggerUserID(v int32) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldTriggerUserID, v))
}

// TriggerAt applies equality check predicate on the "trigger_at" field. It's identical to TriggerAtEQ.
func TriggerAt(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldTriggerAt, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldContent, v))
}

// Reason applies equality check predicate on the "reason" field. It's identical to ReasonEQ.
func Reason(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldReason, v))
}

// GeneratedAt applies equality check predicate on the "generated_at" field. It's identical to GeneratedAtEQ.
func GeneratedAt(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldGeneratedAt, v))
}

// ReportTypeEQ applies the EQ predicate on the "report_type" field.
func ReportTypeEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldReportType, v))
}

// ReportTypeNEQ applies the NEQ predicate on the "report_type" field.
func ReportTypeNEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldReportType, v))
}

// ReportTypeIn applies the In predicate on the "report_type" field.
func ReportTypeIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldReportType, vs...))
}

// ReportTypeNotIn applies the NotIn predicate on the "report_type" field.
func ReportTypeNotIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldReportType, vs...))
}

// ReportTypeGT applies the GT predicate on the "report_type" field.
func ReportTypeGT(v string) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldReportType, v))
}

// ReportTypeGTE applies the GTE predicate on the "report_type" field.
func ReportTypeGTE(v string) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldReportType, v))
}

// ReportTypeLT applies the LT predicate on the "report_type" field.
func ReportTypeLT(v string) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldReportType, v))
}

// ReportTypeLTE applies the LTE predicate on the "report_type" field.
func ReportTypeLTE(v string) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldReportType, v))
}

// ReportTypeContains applies the Contains predicate on the "report_type" field.
func ReportTypeContains(v string) predicate.Report {
	return predicate.Report(sql.FieldContains(FieldReportType, v))
}

// ReportTypeHasPrefix applies the HasPrefix predicate on the "report_type" field.
func ReportTypeHasPrefix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasPrefix(FieldReportType, v))
}

// ReportTypeHasSuffix applies the HasSuffix predicate on the "report_type" field.
func ReportTypeHasSuffix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasSuffix(FieldReportType, v))
}

// ReportTypeEqualFold applies the EqualFold predicate on the "report_type" field.
func ReportTypeEqualFold(v string) predicate.Report {
	return predicate.Report(sql.FieldEqualFold(FieldReportType, v))
}

// ReportTypeContainsFold applies the ContainsFold predicate on the "report_type" field.
func ReportTypeContainsFold(v string) predicate.Report {
	return predicate.Report(sql.FieldContainsFold(FieldReportType, v))
}

// TriggerUserIDEQ applies the EQ predicate on the "trigger_user_id" field.
func TriggerUserIDEQ(v int32) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldTriggerUserID, v))
}

// TriggerUserIDNEQ applies the NEQ predicate on the "trigger_user_id" field.
func TriggerUserIDNEQ(v int32) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldTriggerUserID, v))
}

// TriggerUserIDIn applies the In predicate on the "trigger_user_id" field.
func TriggerUserIDIn(vs ...int32) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldTriggerUserID, vs...))
}

// TriggerUserIDNotIn applies the NotIn predicate on the "trigger_user_id" field.
func TriggerUserIDNotIn(vs ...int32) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldTriggerUserID, vs...))
}

// TriggerUserIDGT applies the GT predicate on the "trigger_user_id" field.
func TriggerUserIDGT(v int32) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldTriggerUserID, v))
}

// TriggerUserIDGTE applies the GTE predicate on the "trigger_user_id" field.
func TriggerUserIDGTE(v int32) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldTriggerUserID, v))
}

// TriggerUserIDLT applies the LT predicate on the "trigger_user_id" field.
func TriggerUserIDLT(v int32) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldTriggerUserID, v))
}

// TriggerUserIDLTE applies the LTE predicate on the "trigger_user_id" field.
func TriggerUserIDLTE(v int32) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldTriggerUserID, v))
}

// TriggerUserIDIsNil applies the IsNil predicate on the "trigger_user_id" field.
func TriggerUserIDIsNil() predicate.Report {
	return predicate.Report(sql.FieldIsNull(FieldTriggerUserID))
}

// TriggerUserIDNotNil applies the NotNil predicate on the "trigger_user_id" field.
func TriggerUserIDNotNil() predicate.Report {
	return predicate.Report(sql.FieldNotNull(FieldTriggerUserID))
}

// TriggerAtEQ applies the EQ predicate on the "trigger_at" field.
func TriggerAtEQ(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldTriggerAt, v))
}

// TriggerAtNEQ applies the NEQ predicate on the "trigger_at" field.
func TriggerAtNEQ(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldTriggerAt, v))
}

// TriggerAtIn applies the In predicate on the "trigger_at" field.
func TriggerAtIn(vs ...time.Time) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldTriggerAt, vs...))
}

// TriggerAtNotIn applies the NotIn predicate on the "trigger_at" field.
func TriggerAtNotIn(vs ...time.Time) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldTriggerAt, vs...))
}

// TriggerAtGT applies the GT predicate on the "trigger_at" field.
func TriggerAtGT(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldTriggerAt, v))
}

// TriggerAtGTE applies the GTE predicate on the "trigger_at" field.
func TriggerAtGTE(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldTriggerAt, v))
}

// TriggerAtLT applies the LT predicate on the "trigger_at" field.
func TriggerAtLT(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldTriggerAt, v))
}

// TriggerAtLTE applies the LTE predicate on the "trigger_at" field.
func TriggerAtLTE(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldTriggerAt, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Report {
	return predicate.Report(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Report {
	return predicate.Report(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Report {
	return predicate.Report(sql.FieldContainsFold(FieldContent, v))
}

// ReasonEQ applies the EQ predicate on the "reason" field.
func ReasonEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldReason, v))
}

// ReasonNEQ applies the NEQ predicate on the "reason" field.
func ReasonNEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldReason, v))
}

// ReasonIn applies the In predicate on the "reason" field.
func ReasonIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldReason, vs...))
}

// ReasonNotIn applies the NotIn predicate on the "reason" field.
func ReasonNotIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldReason, vs...))
}

// ReasonGT applies the GT predicate on the "reason" field.
func ReasonGT(v string) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldReason, v))
}

// ReasonGTE applies the GTE predicate on the "reason" field.
func ReasonGTE(v string) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldReason, v))
}

// ReasonLT applies the LT predicate on the "reason" field.
func ReasonLT(v string) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldReason, v))
}

// ReasonLTE applies the LTE predicate on the "reason" field.
func ReasonLTE(v string) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldReason, v))
}

// ReasonContains applies the Contains predicate on the "reason" field.
func ReasonContains(v string) predicate.Report {
	return predicate.Report(sql.FieldContains(FieldReason, v))
}

// ReasonHasPrefix applies the HasPrefix predicate on the "reason" field.
func ReasonHasPrefix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasPrefix(FieldReason, v))
}

// ReasonHasSuffix applies the HasSuffix predicate on the "reason" field.
func ReasonHasSuffix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasSuffix(FieldReason, v))
}

// ReasonEqualFold applies the EqualFold predicate on the "reason" field.
func ReasonEqualFold(v string) predicate.Report {
	return predicate.Report(sql.FieldEqualFold(FieldReason, v))
}

// ReasonContainsFold applies the ContainsFold predicate on the "reason" field.
func ReasonContainsFold(v string) predicate.Report {
	return predicate.Report(sql.FieldContainsFold(FieldReason, v))
}

// GeneratedAtEQ applies the EQ predicate on the "generated_at" field.
func GeneratedAtEQ(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldGeneratedAt, v))
}

// GeneratedAtNEQ applies the NEQ predicate on the "generated_at" field.
func GeneratedAtNEQ(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldGeneratedAt, v))
}

// GeneratedAtIn applies the In predicate on the "generated_at" field.
func GeneratedAtIn(vs ...time.Time) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldGeneratedAt, vs...))
}

// GeneratedAtNotIn applies the NotIn predicate on the "generated_at" field.
func GeneratedAtNotIn(vs ...time.Time) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldGeneratedAt, vs...))
}

// GeneratedAtGT applies the GT predicate on the "generated_at" field.
func GeneratedAtGT(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldGeneratedAt, v))
}

// GeneratedAtGTE applies the GTE predicate on the "generated_at" field.
func GeneratedAtGTE(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldGeneratedAt, v))
}

// GeneratedAtLT applies the LT predicate on the "generated_at" field.
func GeneratedAtLT(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldGeneratedAt, v))
}

// GeneratedAtLTE applies the LTE predicate on the "generated_at" field.
func GeneratedAtLTE(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldGeneratedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Report) predicate.Report {
	return predicate.Report(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Report) predicate.Report {
	return predicate.Report(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Report) predicate.Report {
	return predicate.Report(sql.NotPredicates(p))
}