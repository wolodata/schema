// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wolodata/schema/ent/article"
)

// ArticleCreate is the builder for creating a Article entity.
type ArticleCreate struct {
	config
	mutation *ArticleMutation
	hooks    []Hook
}

// SetOriginShortID sets the "origin_short_id" field.
func (ac *ArticleCreate) SetOriginShortID(s string) *ArticleCreate {
	ac.mutation.SetOriginShortID(s)
	return ac
}

// SetOriginType sets the "origin_type" field.
func (ac *ArticleCreate) SetOriginType(s string) *ArticleCreate {
	ac.mutation.SetOriginType(s)
	return ac
}

// SetURL sets the "url" field.
func (ac *ArticleCreate) SetURL(s string) *ArticleCreate {
	ac.mutation.SetURL(s)
	return ac
}

// SetTitleChinese sets the "title_chinese" field.
func (ac *ArticleCreate) SetTitleChinese(s string) *ArticleCreate {
	ac.mutation.SetTitleChinese(s)
	return ac
}

// SetNillableTitleChinese sets the "title_chinese" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableTitleChinese(s *string) *ArticleCreate {
	if s != nil {
		ac.SetTitleChinese(*s)
	}
	return ac
}

// SetTitleEnglish sets the "title_english" field.
func (ac *ArticleCreate) SetTitleEnglish(s string) *ArticleCreate {
	ac.mutation.SetTitleEnglish(s)
	return ac
}

// SetNillableTitleEnglish sets the "title_english" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableTitleEnglish(s *string) *ArticleCreate {
	if s != nil {
		ac.SetTitleEnglish(*s)
	}
	return ac
}

// SetAuthor sets the "author" field.
func (ac *ArticleCreate) SetAuthor(s string) *ArticleCreate {
	ac.mutation.SetAuthor(s)
	return ac
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableAuthor(s *string) *ArticleCreate {
	if s != nil {
		ac.SetAuthor(*s)
	}
	return ac
}

// SetTags sets the "tags" field.
func (ac *ArticleCreate) SetTags(s []string) *ArticleCreate {
	ac.mutation.SetTags(s)
	return ac
}

// SetPublishedAt sets the "published_at" field.
func (ac *ArticleCreate) SetPublishedAt(t time.Time) *ArticleCreate {
	ac.mutation.SetPublishedAt(t)
	return ac
}

// SetHTMLChinese sets the "html_chinese" field.
func (ac *ArticleCreate) SetHTMLChinese(s string) *ArticleCreate {
	ac.mutation.SetHTMLChinese(s)
	return ac
}

// SetNillableHTMLChinese sets the "html_chinese" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableHTMLChinese(s *string) *ArticleCreate {
	if s != nil {
		ac.SetHTMLChinese(*s)
	}
	return ac
}

// SetHTMLEnglish sets the "html_english" field.
func (ac *ArticleCreate) SetHTMLEnglish(s string) *ArticleCreate {
	ac.mutation.SetHTMLEnglish(s)
	return ac
}

// SetNillableHTMLEnglish sets the "html_english" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableHTMLEnglish(s *string) *ArticleCreate {
	if s != nil {
		ac.SetHTMLEnglish(*s)
	}
	return ac
}

// SetTextChinese sets the "text_chinese" field.
func (ac *ArticleCreate) SetTextChinese(s string) *ArticleCreate {
	ac.mutation.SetTextChinese(s)
	return ac
}

// SetNillableTextChinese sets the "text_chinese" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableTextChinese(s *string) *ArticleCreate {
	if s != nil {
		ac.SetTextChinese(*s)
	}
	return ac
}

// SetTextEnglish sets the "text_english" field.
func (ac *ArticleCreate) SetTextEnglish(s string) *ArticleCreate {
	ac.mutation.SetTextEnglish(s)
	return ac
}

// SetNillableTextEnglish sets the "text_english" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableTextEnglish(s *string) *ArticleCreate {
	if s != nil {
		ac.SetTextEnglish(*s)
	}
	return ac
}

// SetCrawledAt sets the "crawled_at" field.
func (ac *ArticleCreate) SetCrawledAt(t time.Time) *ArticleCreate {
	ac.mutation.SetCrawledAt(t)
	return ac
}

// SetNillableCrawledAt sets the "crawled_at" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableCrawledAt(t *time.Time) *ArticleCreate {
	if t != nil {
		ac.SetCrawledAt(*t)
	}
	return ac
}

// SetSummaryChinese sets the "summary_chinese" field.
func (ac *ArticleCreate) SetSummaryChinese(s string) *ArticleCreate {
	ac.mutation.SetSummaryChinese(s)
	return ac
}

// SetNillableSummaryChinese sets the "summary_chinese" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableSummaryChinese(s *string) *ArticleCreate {
	if s != nil {
		ac.SetSummaryChinese(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *ArticleCreate) SetID(i int) *ArticleCreate {
	ac.mutation.SetID(i)
	return ac
}

// Mutation returns the ArticleMutation object of the builder.
func (ac *ArticleCreate) Mutation() *ArticleMutation {
	return ac.mutation
}

// Save creates the Article in the database.
func (ac *ArticleCreate) Save(ctx context.Context) (*Article, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArticleCreate) SaveX(ctx context.Context) *Article {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArticleCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArticleCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ArticleCreate) defaults() {
	if _, ok := ac.mutation.CrawledAt(); !ok {
		v := article.DefaultCrawledAt()
		ac.mutation.SetCrawledAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArticleCreate) check() error {
	if _, ok := ac.mutation.OriginShortID(); !ok {
		return &ValidationError{Name: "origin_short_id", err: errors.New(`ent: missing required field "Article.origin_short_id"`)}
	}
	if _, ok := ac.mutation.OriginType(); !ok {
		return &ValidationError{Name: "origin_type", err: errors.New(`ent: missing required field "Article.origin_type"`)}
	}
	if _, ok := ac.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Article.url"`)}
	}
	if v, ok := ac.mutation.URL(); ok {
		if err := article.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Article.url": %w`, err)}
		}
	}
	if _, ok := ac.mutation.PublishedAt(); !ok {
		return &ValidationError{Name: "published_at", err: errors.New(`ent: missing required field "Article.published_at"`)}
	}
	if _, ok := ac.mutation.CrawledAt(); !ok {
		return &ValidationError{Name: "crawled_at", err: errors.New(`ent: missing required field "Article.crawled_at"`)}
	}
	return nil
}

func (ac *ArticleCreate) sqlSave(ctx context.Context) (*Article, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ArticleCreate) createSpec() (*Article, *sqlgraph.CreateSpec) {
	var (
		_node = &Article{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(article.Table, sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.OriginShortID(); ok {
		_spec.SetField(article.FieldOriginShortID, field.TypeString, value)
		_node.OriginShortID = value
	}
	if value, ok := ac.mutation.OriginType(); ok {
		_spec.SetField(article.FieldOriginType, field.TypeString, value)
		_node.OriginType = value
	}
	if value, ok := ac.mutation.URL(); ok {
		_spec.SetField(article.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := ac.mutation.TitleChinese(); ok {
		_spec.SetField(article.FieldTitleChinese, field.TypeString, value)
		_node.TitleChinese = value
	}
	if value, ok := ac.mutation.TitleEnglish(); ok {
		_spec.SetField(article.FieldTitleEnglish, field.TypeString, value)
		_node.TitleEnglish = value
	}
	if value, ok := ac.mutation.Author(); ok {
		_spec.SetField(article.FieldAuthor, field.TypeString, value)
		_node.Author = value
	}
	if value, ok := ac.mutation.Tags(); ok {
		_spec.SetField(article.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := ac.mutation.PublishedAt(); ok {
		_spec.SetField(article.FieldPublishedAt, field.TypeTime, value)
		_node.PublishedAt = value
	}
	if value, ok := ac.mutation.HTMLChinese(); ok {
		_spec.SetField(article.FieldHTMLChinese, field.TypeString, value)
		_node.HTMLChinese = value
	}
	if value, ok := ac.mutation.HTMLEnglish(); ok {
		_spec.SetField(article.FieldHTMLEnglish, field.TypeString, value)
		_node.HTMLEnglish = value
	}
	if value, ok := ac.mutation.TextChinese(); ok {
		_spec.SetField(article.FieldTextChinese, field.TypeString, value)
		_node.TextChinese = value
	}
	if value, ok := ac.mutation.TextEnglish(); ok {
		_spec.SetField(article.FieldTextEnglish, field.TypeString, value)
		_node.TextEnglish = value
	}
	if value, ok := ac.mutation.CrawledAt(); ok {
		_spec.SetField(article.FieldCrawledAt, field.TypeTime, value)
		_node.CrawledAt = value
	}
	if value, ok := ac.mutation.SummaryChinese(); ok {
		_spec.SetField(article.FieldSummaryChinese, field.TypeString, value)
		_node.SummaryChinese = value
	}
	return _node, _spec
}

// ArticleCreateBulk is the builder for creating many Article entities in bulk.
type ArticleCreateBulk struct {
	config
	err      error
	builders []*ArticleCreate
}

// Save creates the Article entities in the database.
func (acb *ArticleCreateBulk) Save(ctx context.Context) ([]*Article, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Article, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArticleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArticleCreateBulk) SaveX(ctx context.Context) []*Article {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArticleCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArticleCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
