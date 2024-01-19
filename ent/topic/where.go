// Code generated by ent, DO NOT EDIT.

package topic

import (
	"entgo.io/ent/dialect/sql"
	"github.com/wolodata/schema/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Topic {
	return predicate.Topic(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Topic {
	return predicate.Topic(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Topic {
	return predicate.Topic(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Topic {
	return predicate.Topic(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Topic {
	return predicate.Topic(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Topic {
	return predicate.Topic(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Topic {
	return predicate.Topic(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldUserID, v))
}

// Keyword applies equality check predicate on the "keyword" field. It's identical to KeywordEQ.
func Keyword(v string) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldKeyword, v))
}

// FollowTitle applies equality check predicate on the "follow_title" field. It's identical to FollowTitleEQ.
func FollowTitle(v bool) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldFollowTitle, v))
}

// FollowContent applies equality check predicate on the "follow_content" field. It's identical to FollowContentEQ.
func FollowContent(v bool) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldFollowContent, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Topic {
	return predicate.Topic(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Topic {
	return predicate.Topic(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Topic {
	return predicate.Topic(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.Topic {
	return predicate.Topic(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.Topic {
	return predicate.Topic(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.Topic {
	return predicate.Topic(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.Topic {
	return predicate.Topic(sql.FieldLTE(FieldUserID, v))
}

// KeywordEQ applies the EQ predicate on the "keyword" field.
func KeywordEQ(v string) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldKeyword, v))
}

// KeywordNEQ applies the NEQ predicate on the "keyword" field.
func KeywordNEQ(v string) predicate.Topic {
	return predicate.Topic(sql.FieldNEQ(FieldKeyword, v))
}

// KeywordIn applies the In predicate on the "keyword" field.
func KeywordIn(vs ...string) predicate.Topic {
	return predicate.Topic(sql.FieldIn(FieldKeyword, vs...))
}

// KeywordNotIn applies the NotIn predicate on the "keyword" field.
func KeywordNotIn(vs ...string) predicate.Topic {
	return predicate.Topic(sql.FieldNotIn(FieldKeyword, vs...))
}

// KeywordGT applies the GT predicate on the "keyword" field.
func KeywordGT(v string) predicate.Topic {
	return predicate.Topic(sql.FieldGT(FieldKeyword, v))
}

// KeywordGTE applies the GTE predicate on the "keyword" field.
func KeywordGTE(v string) predicate.Topic {
	return predicate.Topic(sql.FieldGTE(FieldKeyword, v))
}

// KeywordLT applies the LT predicate on the "keyword" field.
func KeywordLT(v string) predicate.Topic {
	return predicate.Topic(sql.FieldLT(FieldKeyword, v))
}

// KeywordLTE applies the LTE predicate on the "keyword" field.
func KeywordLTE(v string) predicate.Topic {
	return predicate.Topic(sql.FieldLTE(FieldKeyword, v))
}

// KeywordContains applies the Contains predicate on the "keyword" field.
func KeywordContains(v string) predicate.Topic {
	return predicate.Topic(sql.FieldContains(FieldKeyword, v))
}

// KeywordHasPrefix applies the HasPrefix predicate on the "keyword" field.
func KeywordHasPrefix(v string) predicate.Topic {
	return predicate.Topic(sql.FieldHasPrefix(FieldKeyword, v))
}

// KeywordHasSuffix applies the HasSuffix predicate on the "keyword" field.
func KeywordHasSuffix(v string) predicate.Topic {
	return predicate.Topic(sql.FieldHasSuffix(FieldKeyword, v))
}

// KeywordEqualFold applies the EqualFold predicate on the "keyword" field.
func KeywordEqualFold(v string) predicate.Topic {
	return predicate.Topic(sql.FieldEqualFold(FieldKeyword, v))
}

// KeywordContainsFold applies the ContainsFold predicate on the "keyword" field.
func KeywordContainsFold(v string) predicate.Topic {
	return predicate.Topic(sql.FieldContainsFold(FieldKeyword, v))
}

// FollowTitleEQ applies the EQ predicate on the "follow_title" field.
func FollowTitleEQ(v bool) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldFollowTitle, v))
}

// FollowTitleNEQ applies the NEQ predicate on the "follow_title" field.
func FollowTitleNEQ(v bool) predicate.Topic {
	return predicate.Topic(sql.FieldNEQ(FieldFollowTitle, v))
}

// FollowContentEQ applies the EQ predicate on the "follow_content" field.
func FollowContentEQ(v bool) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldFollowContent, v))
}

// FollowContentNEQ applies the NEQ predicate on the "follow_content" field.
func FollowContentNEQ(v bool) predicate.Topic {
	return predicate.Topic(sql.FieldNEQ(FieldFollowContent, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Topic) predicate.Topic {
	return predicate.Topic(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Topic) predicate.Topic {
	return predicate.Topic(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Topic) predicate.Topic {
	return predicate.Topic(sql.NotPredicates(p))
}
