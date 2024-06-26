// Code generated by ent, DO NOT EDIT.

package keywordstrong

import (
	"entgo.io/ent/dialect/sql"
	"github.com/wolodata/schema/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldContainsFold(FieldID, id))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldEQ(FieldCategory, v))
}

// Main applies equality check predicate on the "main" field. It's identical to MainEQ.
func Main(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldEQ(FieldMain, v))
}

// MainCount applies equality check predicate on the "main_count" field. It's identical to MainCountEQ.
func MainCount(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldEQ(FieldMainCount, v))
}

// Sub applies equality check predicate on the "sub" field. It's identical to SubEQ.
func Sub(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldEQ(FieldSub, v))
}

// SubCount applies equality check predicate on the "sub_count" field. It's identical to SubCountEQ.
func SubCount(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldEQ(FieldSubCount, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldLTE(FieldCategory, v))
}

// MainEQ applies the EQ predicate on the "main" field.
func MainEQ(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldEQ(FieldMain, v))
}

// MainNEQ applies the NEQ predicate on the "main" field.
func MainNEQ(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldNEQ(FieldMain, v))
}

// MainIn applies the In predicate on the "main" field.
func MainIn(vs ...string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldIn(FieldMain, vs...))
}

// MainNotIn applies the NotIn predicate on the "main" field.
func MainNotIn(vs ...string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldNotIn(FieldMain, vs...))
}

// MainGT applies the GT predicate on the "main" field.
func MainGT(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldGT(FieldMain, v))
}

// MainGTE applies the GTE predicate on the "main" field.
func MainGTE(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldGTE(FieldMain, v))
}

// MainLT applies the LT predicate on the "main" field.
func MainLT(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldLT(FieldMain, v))
}

// MainLTE applies the LTE predicate on the "main" field.
func MainLTE(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldLTE(FieldMain, v))
}

// MainContains applies the Contains predicate on the "main" field.
func MainContains(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldContains(FieldMain, v))
}

// MainHasPrefix applies the HasPrefix predicate on the "main" field.
func MainHasPrefix(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldHasPrefix(FieldMain, v))
}

// MainHasSuffix applies the HasSuffix predicate on the "main" field.
func MainHasSuffix(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldHasSuffix(FieldMain, v))
}

// MainEqualFold applies the EqualFold predicate on the "main" field.
func MainEqualFold(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldEqualFold(FieldMain, v))
}

// MainContainsFold applies the ContainsFold predicate on the "main" field.
func MainContainsFold(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldContainsFold(FieldMain, v))
}

// MainCountEQ applies the EQ predicate on the "main_count" field.
func MainCountEQ(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldEQ(FieldMainCount, v))
}

// MainCountNEQ applies the NEQ predicate on the "main_count" field.
func MainCountNEQ(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldNEQ(FieldMainCount, v))
}

// MainCountIn applies the In predicate on the "main_count" field.
func MainCountIn(vs ...uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldIn(FieldMainCount, vs...))
}

// MainCountNotIn applies the NotIn predicate on the "main_count" field.
func MainCountNotIn(vs ...uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldNotIn(FieldMainCount, vs...))
}

// MainCountGT applies the GT predicate on the "main_count" field.
func MainCountGT(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldGT(FieldMainCount, v))
}

// MainCountGTE applies the GTE predicate on the "main_count" field.
func MainCountGTE(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldGTE(FieldMainCount, v))
}

// MainCountLT applies the LT predicate on the "main_count" field.
func MainCountLT(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldLT(FieldMainCount, v))
}

// MainCountLTE applies the LTE predicate on the "main_count" field.
func MainCountLTE(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldLTE(FieldMainCount, v))
}

// SubEQ applies the EQ predicate on the "sub" field.
func SubEQ(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldEQ(FieldSub, v))
}

// SubNEQ applies the NEQ predicate on the "sub" field.
func SubNEQ(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldNEQ(FieldSub, v))
}

// SubIn applies the In predicate on the "sub" field.
func SubIn(vs ...string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldIn(FieldSub, vs...))
}

// SubNotIn applies the NotIn predicate on the "sub" field.
func SubNotIn(vs ...string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldNotIn(FieldSub, vs...))
}

// SubGT applies the GT predicate on the "sub" field.
func SubGT(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldGT(FieldSub, v))
}

// SubGTE applies the GTE predicate on the "sub" field.
func SubGTE(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldGTE(FieldSub, v))
}

// SubLT applies the LT predicate on the "sub" field.
func SubLT(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldLT(FieldSub, v))
}

// SubLTE applies the LTE predicate on the "sub" field.
func SubLTE(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldLTE(FieldSub, v))
}

// SubContains applies the Contains predicate on the "sub" field.
func SubContains(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldContains(FieldSub, v))
}

// SubHasPrefix applies the HasPrefix predicate on the "sub" field.
func SubHasPrefix(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldHasPrefix(FieldSub, v))
}

// SubHasSuffix applies the HasSuffix predicate on the "sub" field.
func SubHasSuffix(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldHasSuffix(FieldSub, v))
}

// SubEqualFold applies the EqualFold predicate on the "sub" field.
func SubEqualFold(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldEqualFold(FieldSub, v))
}

// SubContainsFold applies the ContainsFold predicate on the "sub" field.
func SubContainsFold(v string) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldContainsFold(FieldSub, v))
}

// SubCountEQ applies the EQ predicate on the "sub_count" field.
func SubCountEQ(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldEQ(FieldSubCount, v))
}

// SubCountNEQ applies the NEQ predicate on the "sub_count" field.
func SubCountNEQ(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldNEQ(FieldSubCount, v))
}

// SubCountIn applies the In predicate on the "sub_count" field.
func SubCountIn(vs ...uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldIn(FieldSubCount, vs...))
}

// SubCountNotIn applies the NotIn predicate on the "sub_count" field.
func SubCountNotIn(vs ...uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldNotIn(FieldSubCount, vs...))
}

// SubCountGT applies the GT predicate on the "sub_count" field.
func SubCountGT(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldGT(FieldSubCount, v))
}

// SubCountGTE applies the GTE predicate on the "sub_count" field.
func SubCountGTE(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldGTE(FieldSubCount, v))
}

// SubCountLT applies the LT predicate on the "sub_count" field.
func SubCountLT(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldLT(FieldSubCount, v))
}

// SubCountLTE applies the LTE predicate on the "sub_count" field.
func SubCountLTE(v uint64) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.FieldLTE(FieldSubCount, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.KeywordStrong) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.KeywordStrong) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.KeywordStrong) predicate.KeywordStrong {
	return predicate.KeywordStrong(sql.NotPredicates(p))
}
