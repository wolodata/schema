// Code generated by ent, DO NOT EDIT.

package article

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/wolodata/schema/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldID, id))
}

// OriginShortID applies equality check predicate on the "origin_short_id" field. It's identical to OriginShortIDEQ.
func OriginShortID(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldOriginShortID, v))
}

// IsChinese applies equality check predicate on the "is_chinese" field. It's identical to IsChineseEQ.
func IsChinese(v bool) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldIsChinese, v))
}

// OriginType applies equality check predicate on the "origin_type" field. It's identical to OriginTypeEQ.
func OriginType(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldOriginType, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldURL, v))
}

// TitleChinese applies equality check predicate on the "title_chinese" field. It's identical to TitleChineseEQ.
func TitleChinese(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTitleChinese, v))
}

// TitleEnglish applies equality check predicate on the "title_english" field. It's identical to TitleEnglishEQ.
func TitleEnglish(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTitleEnglish, v))
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldAuthor, v))
}

// PublishedAt applies equality check predicate on the "published_at" field. It's identical to PublishedAtEQ.
func PublishedAt(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldPublishedAt, v))
}

// HTMLChinese applies equality check predicate on the "html_chinese" field. It's identical to HTMLChineseEQ.
func HTMLChinese(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldHTMLChinese, v))
}

// HTMLEnglish applies equality check predicate on the "html_english" field. It's identical to HTMLEnglishEQ.
func HTMLEnglish(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldHTMLEnglish, v))
}

// TextChinese applies equality check predicate on the "text_chinese" field. It's identical to TextChineseEQ.
func TextChinese(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTextChinese, v))
}

// TextEnglish applies equality check predicate on the "text_english" field. It's identical to TextEnglishEQ.
func TextEnglish(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTextEnglish, v))
}

// CrawledAt applies equality check predicate on the "crawled_at" field. It's identical to CrawledAtEQ.
func CrawledAt(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldCrawledAt, v))
}

// SummaryChinese applies equality check predicate on the "summary_chinese" field. It's identical to SummaryChineseEQ.
func SummaryChinese(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldSummaryChinese, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldCategory, v))
}

// OriginShortIDEQ applies the EQ predicate on the "origin_short_id" field.
func OriginShortIDEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldOriginShortID, v))
}

// OriginShortIDNEQ applies the NEQ predicate on the "origin_short_id" field.
func OriginShortIDNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldOriginShortID, v))
}

// OriginShortIDIn applies the In predicate on the "origin_short_id" field.
func OriginShortIDIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldOriginShortID, vs...))
}

// OriginShortIDNotIn applies the NotIn predicate on the "origin_short_id" field.
func OriginShortIDNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldOriginShortID, vs...))
}

// OriginShortIDGT applies the GT predicate on the "origin_short_id" field.
func OriginShortIDGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldOriginShortID, v))
}

// OriginShortIDGTE applies the GTE predicate on the "origin_short_id" field.
func OriginShortIDGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldOriginShortID, v))
}

// OriginShortIDLT applies the LT predicate on the "origin_short_id" field.
func OriginShortIDLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldOriginShortID, v))
}

// OriginShortIDLTE applies the LTE predicate on the "origin_short_id" field.
func OriginShortIDLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldOriginShortID, v))
}

// OriginShortIDContains applies the Contains predicate on the "origin_short_id" field.
func OriginShortIDContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldOriginShortID, v))
}

// OriginShortIDHasPrefix applies the HasPrefix predicate on the "origin_short_id" field.
func OriginShortIDHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldOriginShortID, v))
}

// OriginShortIDHasSuffix applies the HasSuffix predicate on the "origin_short_id" field.
func OriginShortIDHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldOriginShortID, v))
}

// OriginShortIDEqualFold applies the EqualFold predicate on the "origin_short_id" field.
func OriginShortIDEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldOriginShortID, v))
}

// OriginShortIDContainsFold applies the ContainsFold predicate on the "origin_short_id" field.
func OriginShortIDContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldOriginShortID, v))
}

// IsChineseEQ applies the EQ predicate on the "is_chinese" field.
func IsChineseEQ(v bool) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldIsChinese, v))
}

// IsChineseNEQ applies the NEQ predicate on the "is_chinese" field.
func IsChineseNEQ(v bool) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldIsChinese, v))
}

// OriginTypeEQ applies the EQ predicate on the "origin_type" field.
func OriginTypeEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldOriginType, v))
}

// OriginTypeNEQ applies the NEQ predicate on the "origin_type" field.
func OriginTypeNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldOriginType, v))
}

// OriginTypeIn applies the In predicate on the "origin_type" field.
func OriginTypeIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldOriginType, vs...))
}

// OriginTypeNotIn applies the NotIn predicate on the "origin_type" field.
func OriginTypeNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldOriginType, vs...))
}

// OriginTypeGT applies the GT predicate on the "origin_type" field.
func OriginTypeGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldOriginType, v))
}

// OriginTypeGTE applies the GTE predicate on the "origin_type" field.
func OriginTypeGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldOriginType, v))
}

// OriginTypeLT applies the LT predicate on the "origin_type" field.
func OriginTypeLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldOriginType, v))
}

// OriginTypeLTE applies the LTE predicate on the "origin_type" field.
func OriginTypeLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldOriginType, v))
}

// OriginTypeContains applies the Contains predicate on the "origin_type" field.
func OriginTypeContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldOriginType, v))
}

// OriginTypeHasPrefix applies the HasPrefix predicate on the "origin_type" field.
func OriginTypeHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldOriginType, v))
}

// OriginTypeHasSuffix applies the HasSuffix predicate on the "origin_type" field.
func OriginTypeHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldOriginType, v))
}

// OriginTypeEqualFold applies the EqualFold predicate on the "origin_type" field.
func OriginTypeEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldOriginType, v))
}

// OriginTypeContainsFold applies the ContainsFold predicate on the "origin_type" field.
func OriginTypeContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldOriginType, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldURL, v))
}

// TitleChineseEQ applies the EQ predicate on the "title_chinese" field.
func TitleChineseEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTitleChinese, v))
}

// TitleChineseNEQ applies the NEQ predicate on the "title_chinese" field.
func TitleChineseNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldTitleChinese, v))
}

// TitleChineseIn applies the In predicate on the "title_chinese" field.
func TitleChineseIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldTitleChinese, vs...))
}

// TitleChineseNotIn applies the NotIn predicate on the "title_chinese" field.
func TitleChineseNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldTitleChinese, vs...))
}

// TitleChineseGT applies the GT predicate on the "title_chinese" field.
func TitleChineseGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldTitleChinese, v))
}

// TitleChineseGTE applies the GTE predicate on the "title_chinese" field.
func TitleChineseGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldTitleChinese, v))
}

// TitleChineseLT applies the LT predicate on the "title_chinese" field.
func TitleChineseLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldTitleChinese, v))
}

// TitleChineseLTE applies the LTE predicate on the "title_chinese" field.
func TitleChineseLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldTitleChinese, v))
}

// TitleChineseContains applies the Contains predicate on the "title_chinese" field.
func TitleChineseContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldTitleChinese, v))
}

// TitleChineseHasPrefix applies the HasPrefix predicate on the "title_chinese" field.
func TitleChineseHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldTitleChinese, v))
}

// TitleChineseHasSuffix applies the HasSuffix predicate on the "title_chinese" field.
func TitleChineseHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldTitleChinese, v))
}

// TitleChineseEqualFold applies the EqualFold predicate on the "title_chinese" field.
func TitleChineseEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldTitleChinese, v))
}

// TitleChineseContainsFold applies the ContainsFold predicate on the "title_chinese" field.
func TitleChineseContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldTitleChinese, v))
}

// TitleEnglishEQ applies the EQ predicate on the "title_english" field.
func TitleEnglishEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTitleEnglish, v))
}

// TitleEnglishNEQ applies the NEQ predicate on the "title_english" field.
func TitleEnglishNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldTitleEnglish, v))
}

// TitleEnglishIn applies the In predicate on the "title_english" field.
func TitleEnglishIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldTitleEnglish, vs...))
}

// TitleEnglishNotIn applies the NotIn predicate on the "title_english" field.
func TitleEnglishNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldTitleEnglish, vs...))
}

// TitleEnglishGT applies the GT predicate on the "title_english" field.
func TitleEnglishGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldTitleEnglish, v))
}

// TitleEnglishGTE applies the GTE predicate on the "title_english" field.
func TitleEnglishGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldTitleEnglish, v))
}

// TitleEnglishLT applies the LT predicate on the "title_english" field.
func TitleEnglishLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldTitleEnglish, v))
}

// TitleEnglishLTE applies the LTE predicate on the "title_english" field.
func TitleEnglishLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldTitleEnglish, v))
}

// TitleEnglishContains applies the Contains predicate on the "title_english" field.
func TitleEnglishContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldTitleEnglish, v))
}

// TitleEnglishHasPrefix applies the HasPrefix predicate on the "title_english" field.
func TitleEnglishHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldTitleEnglish, v))
}

// TitleEnglishHasSuffix applies the HasSuffix predicate on the "title_english" field.
func TitleEnglishHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldTitleEnglish, v))
}

// TitleEnglishEqualFold applies the EqualFold predicate on the "title_english" field.
func TitleEnglishEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldTitleEnglish, v))
}

// TitleEnglishContainsFold applies the ContainsFold predicate on the "title_english" field.
func TitleEnglishContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldTitleEnglish, v))
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldAuthor, v))
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldAuthor, v))
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldAuthor, vs...))
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldAuthor, vs...))
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldAuthor, v))
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldAuthor, v))
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldAuthor, v))
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldAuthor, v))
}

// AuthorContains applies the Contains predicate on the "author" field.
func AuthorContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldAuthor, v))
}

// AuthorHasPrefix applies the HasPrefix predicate on the "author" field.
func AuthorHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldAuthor, v))
}

// AuthorHasSuffix applies the HasSuffix predicate on the "author" field.
func AuthorHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldAuthor, v))
}

// AuthorEqualFold applies the EqualFold predicate on the "author" field.
func AuthorEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldAuthor, v))
}

// AuthorContainsFold applies the ContainsFold predicate on the "author" field.
func AuthorContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldAuthor, v))
}

// PublishedAtEQ applies the EQ predicate on the "published_at" field.
func PublishedAtEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldPublishedAt, v))
}

// PublishedAtNEQ applies the NEQ predicate on the "published_at" field.
func PublishedAtNEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldPublishedAt, v))
}

// PublishedAtIn applies the In predicate on the "published_at" field.
func PublishedAtIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldPublishedAt, vs...))
}

// PublishedAtNotIn applies the NotIn predicate on the "published_at" field.
func PublishedAtNotIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldPublishedAt, vs...))
}

// PublishedAtGT applies the GT predicate on the "published_at" field.
func PublishedAtGT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldPublishedAt, v))
}

// PublishedAtGTE applies the GTE predicate on the "published_at" field.
func PublishedAtGTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldPublishedAt, v))
}

// PublishedAtLT applies the LT predicate on the "published_at" field.
func PublishedAtLT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldPublishedAt, v))
}

// PublishedAtLTE applies the LTE predicate on the "published_at" field.
func PublishedAtLTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldPublishedAt, v))
}

// HTMLChineseEQ applies the EQ predicate on the "html_chinese" field.
func HTMLChineseEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldHTMLChinese, v))
}

// HTMLChineseNEQ applies the NEQ predicate on the "html_chinese" field.
func HTMLChineseNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldHTMLChinese, v))
}

// HTMLChineseIn applies the In predicate on the "html_chinese" field.
func HTMLChineseIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldHTMLChinese, vs...))
}

// HTMLChineseNotIn applies the NotIn predicate on the "html_chinese" field.
func HTMLChineseNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldHTMLChinese, vs...))
}

// HTMLChineseGT applies the GT predicate on the "html_chinese" field.
func HTMLChineseGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldHTMLChinese, v))
}

// HTMLChineseGTE applies the GTE predicate on the "html_chinese" field.
func HTMLChineseGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldHTMLChinese, v))
}

// HTMLChineseLT applies the LT predicate on the "html_chinese" field.
func HTMLChineseLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldHTMLChinese, v))
}

// HTMLChineseLTE applies the LTE predicate on the "html_chinese" field.
func HTMLChineseLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldHTMLChinese, v))
}

// HTMLChineseContains applies the Contains predicate on the "html_chinese" field.
func HTMLChineseContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldHTMLChinese, v))
}

// HTMLChineseHasPrefix applies the HasPrefix predicate on the "html_chinese" field.
func HTMLChineseHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldHTMLChinese, v))
}

// HTMLChineseHasSuffix applies the HasSuffix predicate on the "html_chinese" field.
func HTMLChineseHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldHTMLChinese, v))
}

// HTMLChineseEqualFold applies the EqualFold predicate on the "html_chinese" field.
func HTMLChineseEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldHTMLChinese, v))
}

// HTMLChineseContainsFold applies the ContainsFold predicate on the "html_chinese" field.
func HTMLChineseContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldHTMLChinese, v))
}

// HTMLEnglishEQ applies the EQ predicate on the "html_english" field.
func HTMLEnglishEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldHTMLEnglish, v))
}

// HTMLEnglishNEQ applies the NEQ predicate on the "html_english" field.
func HTMLEnglishNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldHTMLEnglish, v))
}

// HTMLEnglishIn applies the In predicate on the "html_english" field.
func HTMLEnglishIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldHTMLEnglish, vs...))
}

// HTMLEnglishNotIn applies the NotIn predicate on the "html_english" field.
func HTMLEnglishNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldHTMLEnglish, vs...))
}

// HTMLEnglishGT applies the GT predicate on the "html_english" field.
func HTMLEnglishGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldHTMLEnglish, v))
}

// HTMLEnglishGTE applies the GTE predicate on the "html_english" field.
func HTMLEnglishGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldHTMLEnglish, v))
}

// HTMLEnglishLT applies the LT predicate on the "html_english" field.
func HTMLEnglishLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldHTMLEnglish, v))
}

// HTMLEnglishLTE applies the LTE predicate on the "html_english" field.
func HTMLEnglishLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldHTMLEnglish, v))
}

// HTMLEnglishContains applies the Contains predicate on the "html_english" field.
func HTMLEnglishContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldHTMLEnglish, v))
}

// HTMLEnglishHasPrefix applies the HasPrefix predicate on the "html_english" field.
func HTMLEnglishHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldHTMLEnglish, v))
}

// HTMLEnglishHasSuffix applies the HasSuffix predicate on the "html_english" field.
func HTMLEnglishHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldHTMLEnglish, v))
}

// HTMLEnglishEqualFold applies the EqualFold predicate on the "html_english" field.
func HTMLEnglishEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldHTMLEnglish, v))
}

// HTMLEnglishContainsFold applies the ContainsFold predicate on the "html_english" field.
func HTMLEnglishContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldHTMLEnglish, v))
}

// TextChineseEQ applies the EQ predicate on the "text_chinese" field.
func TextChineseEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTextChinese, v))
}

// TextChineseNEQ applies the NEQ predicate on the "text_chinese" field.
func TextChineseNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldTextChinese, v))
}

// TextChineseIn applies the In predicate on the "text_chinese" field.
func TextChineseIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldTextChinese, vs...))
}

// TextChineseNotIn applies the NotIn predicate on the "text_chinese" field.
func TextChineseNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldTextChinese, vs...))
}

// TextChineseGT applies the GT predicate on the "text_chinese" field.
func TextChineseGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldTextChinese, v))
}

// TextChineseGTE applies the GTE predicate on the "text_chinese" field.
func TextChineseGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldTextChinese, v))
}

// TextChineseLT applies the LT predicate on the "text_chinese" field.
func TextChineseLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldTextChinese, v))
}

// TextChineseLTE applies the LTE predicate on the "text_chinese" field.
func TextChineseLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldTextChinese, v))
}

// TextChineseContains applies the Contains predicate on the "text_chinese" field.
func TextChineseContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldTextChinese, v))
}

// TextChineseHasPrefix applies the HasPrefix predicate on the "text_chinese" field.
func TextChineseHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldTextChinese, v))
}

// TextChineseHasSuffix applies the HasSuffix predicate on the "text_chinese" field.
func TextChineseHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldTextChinese, v))
}

// TextChineseEqualFold applies the EqualFold predicate on the "text_chinese" field.
func TextChineseEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldTextChinese, v))
}

// TextChineseContainsFold applies the ContainsFold predicate on the "text_chinese" field.
func TextChineseContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldTextChinese, v))
}

// TextEnglishEQ applies the EQ predicate on the "text_english" field.
func TextEnglishEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTextEnglish, v))
}

// TextEnglishNEQ applies the NEQ predicate on the "text_english" field.
func TextEnglishNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldTextEnglish, v))
}

// TextEnglishIn applies the In predicate on the "text_english" field.
func TextEnglishIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldTextEnglish, vs...))
}

// TextEnglishNotIn applies the NotIn predicate on the "text_english" field.
func TextEnglishNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldTextEnglish, vs...))
}

// TextEnglishGT applies the GT predicate on the "text_english" field.
func TextEnglishGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldTextEnglish, v))
}

// TextEnglishGTE applies the GTE predicate on the "text_english" field.
func TextEnglishGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldTextEnglish, v))
}

// TextEnglishLT applies the LT predicate on the "text_english" field.
func TextEnglishLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldTextEnglish, v))
}

// TextEnglishLTE applies the LTE predicate on the "text_english" field.
func TextEnglishLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldTextEnglish, v))
}

// TextEnglishContains applies the Contains predicate on the "text_english" field.
func TextEnglishContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldTextEnglish, v))
}

// TextEnglishHasPrefix applies the HasPrefix predicate on the "text_english" field.
func TextEnglishHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldTextEnglish, v))
}

// TextEnglishHasSuffix applies the HasSuffix predicate on the "text_english" field.
func TextEnglishHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldTextEnglish, v))
}

// TextEnglishEqualFold applies the EqualFold predicate on the "text_english" field.
func TextEnglishEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldTextEnglish, v))
}

// TextEnglishContainsFold applies the ContainsFold predicate on the "text_english" field.
func TextEnglishContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldTextEnglish, v))
}

// CrawledAtEQ applies the EQ predicate on the "crawled_at" field.
func CrawledAtEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldCrawledAt, v))
}

// CrawledAtNEQ applies the NEQ predicate on the "crawled_at" field.
func CrawledAtNEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldCrawledAt, v))
}

// CrawledAtIn applies the In predicate on the "crawled_at" field.
func CrawledAtIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldCrawledAt, vs...))
}

// CrawledAtNotIn applies the NotIn predicate on the "crawled_at" field.
func CrawledAtNotIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldCrawledAt, vs...))
}

// CrawledAtGT applies the GT predicate on the "crawled_at" field.
func CrawledAtGT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldCrawledAt, v))
}

// CrawledAtGTE applies the GTE predicate on the "crawled_at" field.
func CrawledAtGTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldCrawledAt, v))
}

// CrawledAtLT applies the LT predicate on the "crawled_at" field.
func CrawledAtLT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldCrawledAt, v))
}

// CrawledAtLTE applies the LTE predicate on the "crawled_at" field.
func CrawledAtLTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldCrawledAt, v))
}

// SummaryChineseEQ applies the EQ predicate on the "summary_chinese" field.
func SummaryChineseEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldSummaryChinese, v))
}

// SummaryChineseNEQ applies the NEQ predicate on the "summary_chinese" field.
func SummaryChineseNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldSummaryChinese, v))
}

// SummaryChineseIn applies the In predicate on the "summary_chinese" field.
func SummaryChineseIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldSummaryChinese, vs...))
}

// SummaryChineseNotIn applies the NotIn predicate on the "summary_chinese" field.
func SummaryChineseNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldSummaryChinese, vs...))
}

// SummaryChineseGT applies the GT predicate on the "summary_chinese" field.
func SummaryChineseGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldSummaryChinese, v))
}

// SummaryChineseGTE applies the GTE predicate on the "summary_chinese" field.
func SummaryChineseGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldSummaryChinese, v))
}

// SummaryChineseLT applies the LT predicate on the "summary_chinese" field.
func SummaryChineseLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldSummaryChinese, v))
}

// SummaryChineseLTE applies the LTE predicate on the "summary_chinese" field.
func SummaryChineseLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldSummaryChinese, v))
}

// SummaryChineseContains applies the Contains predicate on the "summary_chinese" field.
func SummaryChineseContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldSummaryChinese, v))
}

// SummaryChineseHasPrefix applies the HasPrefix predicate on the "summary_chinese" field.
func SummaryChineseHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldSummaryChinese, v))
}

// SummaryChineseHasSuffix applies the HasSuffix predicate on the "summary_chinese" field.
func SummaryChineseHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldSummaryChinese, v))
}

// SummaryChineseEqualFold applies the EqualFold predicate on the "summary_chinese" field.
func SummaryChineseEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldSummaryChinese, v))
}

// SummaryChineseContainsFold applies the ContainsFold predicate on the "summary_chinese" field.
func SummaryChineseContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldSummaryChinese, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldCategory, v))
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldCategory, v))
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldCategory, v))
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldCategory, v))
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldCategory, v))
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldCategory, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Article) predicate.Article {
	return predicate.Article(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Article) predicate.Article {
	return predicate.Article(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Article) predicate.Article {
	return predicate.Article(sql.NotPredicates(p))
}
