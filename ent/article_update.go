// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/wolodata/schema/ent/article"
	"github.com/wolodata/schema/ent/predicate"
)

// ArticleUpdate is the builder for updating Article entities.
type ArticleUpdate struct {
	config
	hooks    []Hook
	mutation *ArticleMutation
}

// Where appends a list predicates to the ArticleUpdate builder.
func (au *ArticleUpdate) Where(ps ...predicate.Article) *ArticleUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetTitleChinese sets the "title_chinese" field.
func (au *ArticleUpdate) SetTitleChinese(s string) *ArticleUpdate {
	au.mutation.SetTitleChinese(s)
	return au
}

// SetNillableTitleChinese sets the "title_chinese" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableTitleChinese(s *string) *ArticleUpdate {
	if s != nil {
		au.SetTitleChinese(*s)
	}
	return au
}

// SetTitleEnglish sets the "title_english" field.
func (au *ArticleUpdate) SetTitleEnglish(s string) *ArticleUpdate {
	au.mutation.SetTitleEnglish(s)
	return au
}

// SetNillableTitleEnglish sets the "title_english" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableTitleEnglish(s *string) *ArticleUpdate {
	if s != nil {
		au.SetTitleEnglish(*s)
	}
	return au
}

// SetAuthor sets the "author" field.
func (au *ArticleUpdate) SetAuthor(s string) *ArticleUpdate {
	au.mutation.SetAuthor(s)
	return au
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableAuthor(s *string) *ArticleUpdate {
	if s != nil {
		au.SetAuthor(*s)
	}
	return au
}

// SetHTMLChinese sets the "html_chinese" field.
func (au *ArticleUpdate) SetHTMLChinese(s string) *ArticleUpdate {
	au.mutation.SetHTMLChinese(s)
	return au
}

// SetNillableHTMLChinese sets the "html_chinese" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableHTMLChinese(s *string) *ArticleUpdate {
	if s != nil {
		au.SetHTMLChinese(*s)
	}
	return au
}

// SetHTMLEnglish sets the "html_english" field.
func (au *ArticleUpdate) SetHTMLEnglish(s string) *ArticleUpdate {
	au.mutation.SetHTMLEnglish(s)
	return au
}

// SetNillableHTMLEnglish sets the "html_english" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableHTMLEnglish(s *string) *ArticleUpdate {
	if s != nil {
		au.SetHTMLEnglish(*s)
	}
	return au
}

// SetTextChinese sets the "text_chinese" field.
func (au *ArticleUpdate) SetTextChinese(s string) *ArticleUpdate {
	au.mutation.SetTextChinese(s)
	return au
}

// SetNillableTextChinese sets the "text_chinese" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableTextChinese(s *string) *ArticleUpdate {
	if s != nil {
		au.SetTextChinese(*s)
	}
	return au
}

// SetTextEnglish sets the "text_english" field.
func (au *ArticleUpdate) SetTextEnglish(s string) *ArticleUpdate {
	au.mutation.SetTextEnglish(s)
	return au
}

// SetNillableTextEnglish sets the "text_english" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableTextEnglish(s *string) *ArticleUpdate {
	if s != nil {
		au.SetTextEnglish(*s)
	}
	return au
}

// SetIsChinaRelated sets the "is_china_related" field.
func (au *ArticleUpdate) SetIsChinaRelated(b bool) *ArticleUpdate {
	au.mutation.SetIsChinaRelated(b)
	return au
}

// SetNillableIsChinaRelated sets the "is_china_related" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableIsChinaRelated(b *bool) *ArticleUpdate {
	if b != nil {
		au.SetIsChinaRelated(*b)
	}
	return au
}

// SetChinaRelatedKeywords sets the "china_related_keywords" field.
func (au *ArticleUpdate) SetChinaRelatedKeywords(s []string) *ArticleUpdate {
	au.mutation.SetChinaRelatedKeywords(s)
	return au
}

// AppendChinaRelatedKeywords appends s to the "china_related_keywords" field.
func (au *ArticleUpdate) AppendChinaRelatedKeywords(s []string) *ArticleUpdate {
	au.mutation.AppendChinaRelatedKeywords(s)
	return au
}

// SetIsChinaStrongRelated sets the "is_china_strong_related" field.
func (au *ArticleUpdate) SetIsChinaStrongRelated(b bool) *ArticleUpdate {
	au.mutation.SetIsChinaStrongRelated(b)
	return au
}

// SetNillableIsChinaStrongRelated sets the "is_china_strong_related" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableIsChinaStrongRelated(b *bool) *ArticleUpdate {
	if b != nil {
		au.SetIsChinaStrongRelated(*b)
	}
	return au
}

// SetChinaRelatedCategory sets the "china_related_category" field.
func (au *ArticleUpdate) SetChinaRelatedCategory(s string) *ArticleUpdate {
	au.mutation.SetChinaRelatedCategory(s)
	return au
}

// SetNillableChinaRelatedCategory sets the "china_related_category" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableChinaRelatedCategory(s *string) *ArticleUpdate {
	if s != nil {
		au.SetChinaRelatedCategory(*s)
	}
	return au
}

// SetSummaryChinese sets the "summary_chinese" field.
func (au *ArticleUpdate) SetSummaryChinese(s string) *ArticleUpdate {
	au.mutation.SetSummaryChinese(s)
	return au
}

// SetNillableSummaryChinese sets the "summary_chinese" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableSummaryChinese(s *string) *ArticleUpdate {
	if s != nil {
		au.SetSummaryChinese(*s)
	}
	return au
}

// Mutation returns the ArticleMutation object of the builder.
func (au *ArticleUpdate) Mutation() *ArticleMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArticleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArticleUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArticleUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArticleUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ArticleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(article.Table, article.Columns, sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.TitleChinese(); ok {
		_spec.SetField(article.FieldTitleChinese, field.TypeString, value)
	}
	if value, ok := au.mutation.TitleEnglish(); ok {
		_spec.SetField(article.FieldTitleEnglish, field.TypeString, value)
	}
	if value, ok := au.mutation.Author(); ok {
		_spec.SetField(article.FieldAuthor, field.TypeString, value)
	}
	if value, ok := au.mutation.HTMLChinese(); ok {
		_spec.SetField(article.FieldHTMLChinese, field.TypeString, value)
	}
	if value, ok := au.mutation.HTMLEnglish(); ok {
		_spec.SetField(article.FieldHTMLEnglish, field.TypeString, value)
	}
	if value, ok := au.mutation.TextChinese(); ok {
		_spec.SetField(article.FieldTextChinese, field.TypeString, value)
	}
	if value, ok := au.mutation.TextEnglish(); ok {
		_spec.SetField(article.FieldTextEnglish, field.TypeString, value)
	}
	if value, ok := au.mutation.IsChinaRelated(); ok {
		_spec.SetField(article.FieldIsChinaRelated, field.TypeBool, value)
	}
	if value, ok := au.mutation.ChinaRelatedKeywords(); ok {
		_spec.SetField(article.FieldChinaRelatedKeywords, field.TypeJSON, value)
	}
	if value, ok := au.mutation.AppendedChinaRelatedKeywords(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, article.FieldChinaRelatedKeywords, value)
		})
	}
	if value, ok := au.mutation.IsChinaStrongRelated(); ok {
		_spec.SetField(article.FieldIsChinaStrongRelated, field.TypeBool, value)
	}
	if value, ok := au.mutation.ChinaRelatedCategory(); ok {
		_spec.SetField(article.FieldChinaRelatedCategory, field.TypeString, value)
	}
	if value, ok := au.mutation.SummaryChinese(); ok {
		_spec.SetField(article.FieldSummaryChinese, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ArticleUpdateOne is the builder for updating a single Article entity.
type ArticleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArticleMutation
}

// SetTitleChinese sets the "title_chinese" field.
func (auo *ArticleUpdateOne) SetTitleChinese(s string) *ArticleUpdateOne {
	auo.mutation.SetTitleChinese(s)
	return auo
}

// SetNillableTitleChinese sets the "title_chinese" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableTitleChinese(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetTitleChinese(*s)
	}
	return auo
}

// SetTitleEnglish sets the "title_english" field.
func (auo *ArticleUpdateOne) SetTitleEnglish(s string) *ArticleUpdateOne {
	auo.mutation.SetTitleEnglish(s)
	return auo
}

// SetNillableTitleEnglish sets the "title_english" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableTitleEnglish(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetTitleEnglish(*s)
	}
	return auo
}

// SetAuthor sets the "author" field.
func (auo *ArticleUpdateOne) SetAuthor(s string) *ArticleUpdateOne {
	auo.mutation.SetAuthor(s)
	return auo
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableAuthor(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetAuthor(*s)
	}
	return auo
}

// SetHTMLChinese sets the "html_chinese" field.
func (auo *ArticleUpdateOne) SetHTMLChinese(s string) *ArticleUpdateOne {
	auo.mutation.SetHTMLChinese(s)
	return auo
}

// SetNillableHTMLChinese sets the "html_chinese" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableHTMLChinese(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetHTMLChinese(*s)
	}
	return auo
}

// SetHTMLEnglish sets the "html_english" field.
func (auo *ArticleUpdateOne) SetHTMLEnglish(s string) *ArticleUpdateOne {
	auo.mutation.SetHTMLEnglish(s)
	return auo
}

// SetNillableHTMLEnglish sets the "html_english" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableHTMLEnglish(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetHTMLEnglish(*s)
	}
	return auo
}

// SetTextChinese sets the "text_chinese" field.
func (auo *ArticleUpdateOne) SetTextChinese(s string) *ArticleUpdateOne {
	auo.mutation.SetTextChinese(s)
	return auo
}

// SetNillableTextChinese sets the "text_chinese" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableTextChinese(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetTextChinese(*s)
	}
	return auo
}

// SetTextEnglish sets the "text_english" field.
func (auo *ArticleUpdateOne) SetTextEnglish(s string) *ArticleUpdateOne {
	auo.mutation.SetTextEnglish(s)
	return auo
}

// SetNillableTextEnglish sets the "text_english" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableTextEnglish(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetTextEnglish(*s)
	}
	return auo
}

// SetIsChinaRelated sets the "is_china_related" field.
func (auo *ArticleUpdateOne) SetIsChinaRelated(b bool) *ArticleUpdateOne {
	auo.mutation.SetIsChinaRelated(b)
	return auo
}

// SetNillableIsChinaRelated sets the "is_china_related" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableIsChinaRelated(b *bool) *ArticleUpdateOne {
	if b != nil {
		auo.SetIsChinaRelated(*b)
	}
	return auo
}

// SetChinaRelatedKeywords sets the "china_related_keywords" field.
func (auo *ArticleUpdateOne) SetChinaRelatedKeywords(s []string) *ArticleUpdateOne {
	auo.mutation.SetChinaRelatedKeywords(s)
	return auo
}

// AppendChinaRelatedKeywords appends s to the "china_related_keywords" field.
func (auo *ArticleUpdateOne) AppendChinaRelatedKeywords(s []string) *ArticleUpdateOne {
	auo.mutation.AppendChinaRelatedKeywords(s)
	return auo
}

// SetIsChinaStrongRelated sets the "is_china_strong_related" field.
func (auo *ArticleUpdateOne) SetIsChinaStrongRelated(b bool) *ArticleUpdateOne {
	auo.mutation.SetIsChinaStrongRelated(b)
	return auo
}

// SetNillableIsChinaStrongRelated sets the "is_china_strong_related" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableIsChinaStrongRelated(b *bool) *ArticleUpdateOne {
	if b != nil {
		auo.SetIsChinaStrongRelated(*b)
	}
	return auo
}

// SetChinaRelatedCategory sets the "china_related_category" field.
func (auo *ArticleUpdateOne) SetChinaRelatedCategory(s string) *ArticleUpdateOne {
	auo.mutation.SetChinaRelatedCategory(s)
	return auo
}

// SetNillableChinaRelatedCategory sets the "china_related_category" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableChinaRelatedCategory(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetChinaRelatedCategory(*s)
	}
	return auo
}

// SetSummaryChinese sets the "summary_chinese" field.
func (auo *ArticleUpdateOne) SetSummaryChinese(s string) *ArticleUpdateOne {
	auo.mutation.SetSummaryChinese(s)
	return auo
}

// SetNillableSummaryChinese sets the "summary_chinese" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableSummaryChinese(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetSummaryChinese(*s)
	}
	return auo
}

// Mutation returns the ArticleMutation object of the builder.
func (auo *ArticleUpdateOne) Mutation() *ArticleMutation {
	return auo.mutation
}

// Where appends a list predicates to the ArticleUpdate builder.
func (auo *ArticleUpdateOne) Where(ps ...predicate.Article) *ArticleUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArticleUpdateOne) Select(field string, fields ...string) *ArticleUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Article entity.
func (auo *ArticleUpdateOne) Save(ctx context.Context) (*Article, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArticleUpdateOne) SaveX(ctx context.Context) *Article {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArticleUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArticleUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ArticleUpdateOne) sqlSave(ctx context.Context) (_node *Article, err error) {
	_spec := sqlgraph.NewUpdateSpec(article.Table, article.Columns, sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Article.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, article.FieldID)
		for _, f := range fields {
			if !article.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != article.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.TitleChinese(); ok {
		_spec.SetField(article.FieldTitleChinese, field.TypeString, value)
	}
	if value, ok := auo.mutation.TitleEnglish(); ok {
		_spec.SetField(article.FieldTitleEnglish, field.TypeString, value)
	}
	if value, ok := auo.mutation.Author(); ok {
		_spec.SetField(article.FieldAuthor, field.TypeString, value)
	}
	if value, ok := auo.mutation.HTMLChinese(); ok {
		_spec.SetField(article.FieldHTMLChinese, field.TypeString, value)
	}
	if value, ok := auo.mutation.HTMLEnglish(); ok {
		_spec.SetField(article.FieldHTMLEnglish, field.TypeString, value)
	}
	if value, ok := auo.mutation.TextChinese(); ok {
		_spec.SetField(article.FieldTextChinese, field.TypeString, value)
	}
	if value, ok := auo.mutation.TextEnglish(); ok {
		_spec.SetField(article.FieldTextEnglish, field.TypeString, value)
	}
	if value, ok := auo.mutation.IsChinaRelated(); ok {
		_spec.SetField(article.FieldIsChinaRelated, field.TypeBool, value)
	}
	if value, ok := auo.mutation.ChinaRelatedKeywords(); ok {
		_spec.SetField(article.FieldChinaRelatedKeywords, field.TypeJSON, value)
	}
	if value, ok := auo.mutation.AppendedChinaRelatedKeywords(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, article.FieldChinaRelatedKeywords, value)
		})
	}
	if value, ok := auo.mutation.IsChinaStrongRelated(); ok {
		_spec.SetField(article.FieldIsChinaStrongRelated, field.TypeBool, value)
	}
	if value, ok := auo.mutation.ChinaRelatedCategory(); ok {
		_spec.SetField(article.FieldChinaRelatedCategory, field.TypeString, value)
	}
	if value, ok := auo.mutation.SummaryChinese(); ok {
		_spec.SetField(article.FieldSummaryChinese, field.TypeString, value)
	}
	_node = &Article{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
