// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wolodata/schema/ent/html"
)

// HTMLCreate is the builder for creating a Html entity.
type HTMLCreate struct {
	config
	mutation *HTMLMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetOriginShortID sets the "origin_short_id" field.
func (hc *HTMLCreate) SetOriginShortID(s string) *HTMLCreate {
	hc.mutation.SetOriginShortID(s)
	return hc
}

// SetIsChinese sets the "is_chinese" field.
func (hc *HTMLCreate) SetIsChinese(b bool) *HTMLCreate {
	hc.mutation.SetIsChinese(b)
	return hc
}

// SetNillableIsChinese sets the "is_chinese" field if the given value is not nil.
func (hc *HTMLCreate) SetNillableIsChinese(b *bool) *HTMLCreate {
	if b != nil {
		hc.SetIsChinese(*b)
	}
	return hc
}

// SetURL sets the "url" field.
func (hc *HTMLCreate) SetURL(s string) *HTMLCreate {
	hc.mutation.SetURL(s)
	return hc
}

// SetHTML sets the "html" field.
func (hc *HTMLCreate) SetHTML(s string) *HTMLCreate {
	hc.mutation.SetHTML(s)
	return hc
}

// SetCrawledAt sets the "crawled_at" field.
func (hc *HTMLCreate) SetCrawledAt(t time.Time) *HTMLCreate {
	hc.mutation.SetCrawledAt(t)
	return hc
}

// SetNillableCrawledAt sets the "crawled_at" field if the given value is not nil.
func (hc *HTMLCreate) SetNillableCrawledAt(t *time.Time) *HTMLCreate {
	if t != nil {
		hc.SetCrawledAt(*t)
	}
	return hc
}

// SetAnalyzedAt sets the "analyzed_at" field.
func (hc *HTMLCreate) SetAnalyzedAt(t time.Time) *HTMLCreate {
	hc.mutation.SetAnalyzedAt(t)
	return hc
}

// SetNillableAnalyzedAt sets the "analyzed_at" field if the given value is not nil.
func (hc *HTMLCreate) SetNillableAnalyzedAt(t *time.Time) *HTMLCreate {
	if t != nil {
		hc.SetAnalyzedAt(*t)
	}
	return hc
}

// SetID sets the "id" field.
func (hc *HTMLCreate) SetID(s string) *HTMLCreate {
	hc.mutation.SetID(s)
	return hc
}

// Mutation returns the HTMLMutation object of the builder.
func (hc *HTMLCreate) Mutation() *HTMLMutation {
	return hc.mutation
}

// Save creates the Html in the database.
func (hc *HTMLCreate) Save(ctx context.Context) (*Html, error) {
	hc.defaults()
	return withHooks(ctx, hc.sqlSave, hc.mutation, hc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HTMLCreate) SaveX(ctx context.Context) *Html {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HTMLCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HTMLCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hc *HTMLCreate) defaults() {
	if _, ok := hc.mutation.IsChinese(); !ok {
		v := html.DefaultIsChinese
		hc.mutation.SetIsChinese(v)
	}
	if _, ok := hc.mutation.CrawledAt(); !ok {
		v := html.DefaultCrawledAt()
		hc.mutation.SetCrawledAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HTMLCreate) check() error {
	if _, ok := hc.mutation.OriginShortID(); !ok {
		return &ValidationError{Name: "origin_short_id", err: errors.New(`ent: missing required field "Html.origin_short_id"`)}
	}
	if v, ok := hc.mutation.OriginShortID(); ok {
		if err := html.OriginShortIDValidator(v); err != nil {
			return &ValidationError{Name: "origin_short_id", err: fmt.Errorf(`ent: validator failed for field "Html.origin_short_id": %w`, err)}
		}
	}
	if _, ok := hc.mutation.IsChinese(); !ok {
		return &ValidationError{Name: "is_chinese", err: errors.New(`ent: missing required field "Html.is_chinese"`)}
	}
	if _, ok := hc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Html.url"`)}
	}
	if v, ok := hc.mutation.URL(); ok {
		if err := html.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Html.url": %w`, err)}
		}
	}
	if _, ok := hc.mutation.HTML(); !ok {
		return &ValidationError{Name: "html", err: errors.New(`ent: missing required field "Html.html"`)}
	}
	if v, ok := hc.mutation.HTML(); ok {
		if err := html.HTMLValidator(v); err != nil {
			return &ValidationError{Name: "html", err: fmt.Errorf(`ent: validator failed for field "Html.html": %w`, err)}
		}
	}
	if _, ok := hc.mutation.CrawledAt(); !ok {
		return &ValidationError{Name: "crawled_at", err: errors.New(`ent: missing required field "Html.crawled_at"`)}
	}
	return nil
}

func (hc *HTMLCreate) sqlSave(ctx context.Context) (*Html, error) {
	if err := hc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Html.ID type: %T", _spec.ID.Value)
		}
	}
	hc.mutation.id = &_node.ID
	hc.mutation.done = true
	return _node, nil
}

func (hc *HTMLCreate) createSpec() (*Html, *sqlgraph.CreateSpec) {
	var (
		_node = &Html{config: hc.config}
		_spec = sqlgraph.NewCreateSpec(html.Table, sqlgraph.NewFieldSpec(html.FieldID, field.TypeString))
	)
	_spec.OnConflict = hc.conflict
	if id, ok := hc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := hc.mutation.OriginShortID(); ok {
		_spec.SetField(html.FieldOriginShortID, field.TypeString, value)
		_node.OriginShortID = value
	}
	if value, ok := hc.mutation.IsChinese(); ok {
		_spec.SetField(html.FieldIsChinese, field.TypeBool, value)
		_node.IsChinese = value
	}
	if value, ok := hc.mutation.URL(); ok {
		_spec.SetField(html.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := hc.mutation.HTML(); ok {
		_spec.SetField(html.FieldHTML, field.TypeString, value)
		_node.HTML = value
	}
	if value, ok := hc.mutation.CrawledAt(); ok {
		_spec.SetField(html.FieldCrawledAt, field.TypeTime, value)
		_node.CrawledAt = value
	}
	if value, ok := hc.mutation.AnalyzedAt(); ok {
		_spec.SetField(html.FieldAnalyzedAt, field.TypeTime, value)
		_node.AnalyzedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Html.Create().
//		SetOriginShortID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HtmlUpsert) {
//			SetOriginShortID(v+v).
//		}).
//		Exec(ctx)
func (hc *HTMLCreate) OnConflict(opts ...sql.ConflictOption) *HtmlUpsertOne {
	hc.conflict = opts
	return &HtmlUpsertOne{
		create: hc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Html.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (hc *HTMLCreate) OnConflictColumns(columns ...string) *HtmlUpsertOne {
	hc.conflict = append(hc.conflict, sql.ConflictColumns(columns...))
	return &HtmlUpsertOne{
		create: hc,
	}
}

type (
	// HtmlUpsertOne is the builder for "upsert"-ing
	//  one Html node.
	HtmlUpsertOne struct {
		create *HTMLCreate
	}

	// HtmlUpsert is the "OnConflict" setter.
	HtmlUpsert struct {
		*sql.UpdateSet
	}
)

// SetAnalyzedAt sets the "analyzed_at" field.
func (u *HtmlUpsert) SetAnalyzedAt(v time.Time) *HtmlUpsert {
	u.Set(html.FieldAnalyzedAt, v)
	return u
}

// UpdateAnalyzedAt sets the "analyzed_at" field to the value that was provided on create.
func (u *HtmlUpsert) UpdateAnalyzedAt() *HtmlUpsert {
	u.SetExcluded(html.FieldAnalyzedAt)
	return u
}

// ClearAnalyzedAt clears the value of the "analyzed_at" field.
func (u *HtmlUpsert) ClearAnalyzedAt() *HtmlUpsert {
	u.SetNull(html.FieldAnalyzedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Html.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(html.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *HtmlUpsertOne) UpdateNewValues() *HtmlUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(html.FieldID)
		}
		if _, exists := u.create.mutation.OriginShortID(); exists {
			s.SetIgnore(html.FieldOriginShortID)
		}
		if _, exists := u.create.mutation.IsChinese(); exists {
			s.SetIgnore(html.FieldIsChinese)
		}
		if _, exists := u.create.mutation.URL(); exists {
			s.SetIgnore(html.FieldURL)
		}
		if _, exists := u.create.mutation.HTML(); exists {
			s.SetIgnore(html.FieldHTML)
		}
		if _, exists := u.create.mutation.CrawledAt(); exists {
			s.SetIgnore(html.FieldCrawledAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Html.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *HtmlUpsertOne) Ignore() *HtmlUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HtmlUpsertOne) DoNothing() *HtmlUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HTMLCreate.OnConflict
// documentation for more info.
func (u *HtmlUpsertOne) Update(set func(*HtmlUpsert)) *HtmlUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HtmlUpsert{UpdateSet: update})
	}))
	return u
}

// SetAnalyzedAt sets the "analyzed_at" field.
func (u *HtmlUpsertOne) SetAnalyzedAt(v time.Time) *HtmlUpsertOne {
	return u.Update(func(s *HtmlUpsert) {
		s.SetAnalyzedAt(v)
	})
}

// UpdateAnalyzedAt sets the "analyzed_at" field to the value that was provided on create.
func (u *HtmlUpsertOne) UpdateAnalyzedAt() *HtmlUpsertOne {
	return u.Update(func(s *HtmlUpsert) {
		s.UpdateAnalyzedAt()
	})
}

// ClearAnalyzedAt clears the value of the "analyzed_at" field.
func (u *HtmlUpsertOne) ClearAnalyzedAt() *HtmlUpsertOne {
	return u.Update(func(s *HtmlUpsert) {
		s.ClearAnalyzedAt()
	})
}

// Exec executes the query.
func (u *HtmlUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for HTMLCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HtmlUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *HtmlUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: HtmlUpsertOne.ID is not supported by MySQL driver. Use HtmlUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *HtmlUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// HTMLCreateBulk is the builder for creating many Html entities in bulk.
type HTMLCreateBulk struct {
	config
	err      error
	builders []*HTMLCreate
	conflict []sql.ConflictOption
}

// Save creates the Html entities in the database.
func (hcb *HTMLCreateBulk) Save(ctx context.Context) ([]*Html, error) {
	if hcb.err != nil {
		return nil, hcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Html, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HTMLMutation)
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
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = hcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HTMLCreateBulk) SaveX(ctx context.Context) []*Html {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HTMLCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HTMLCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Html.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HtmlUpsert) {
//			SetOriginShortID(v+v).
//		}).
//		Exec(ctx)
func (hcb *HTMLCreateBulk) OnConflict(opts ...sql.ConflictOption) *HtmlUpsertBulk {
	hcb.conflict = opts
	return &HtmlUpsertBulk{
		create: hcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Html.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (hcb *HTMLCreateBulk) OnConflictColumns(columns ...string) *HtmlUpsertBulk {
	hcb.conflict = append(hcb.conflict, sql.ConflictColumns(columns...))
	return &HtmlUpsertBulk{
		create: hcb,
	}
}

// HtmlUpsertBulk is the builder for "upsert"-ing
// a bulk of Html nodes.
type HtmlUpsertBulk struct {
	create *HTMLCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Html.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(html.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *HtmlUpsertBulk) UpdateNewValues() *HtmlUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(html.FieldID)
			}
			if _, exists := b.mutation.OriginShortID(); exists {
				s.SetIgnore(html.FieldOriginShortID)
			}
			if _, exists := b.mutation.IsChinese(); exists {
				s.SetIgnore(html.FieldIsChinese)
			}
			if _, exists := b.mutation.URL(); exists {
				s.SetIgnore(html.FieldURL)
			}
			if _, exists := b.mutation.HTML(); exists {
				s.SetIgnore(html.FieldHTML)
			}
			if _, exists := b.mutation.CrawledAt(); exists {
				s.SetIgnore(html.FieldCrawledAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Html.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *HtmlUpsertBulk) Ignore() *HtmlUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HtmlUpsertBulk) DoNothing() *HtmlUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HTMLCreateBulk.OnConflict
// documentation for more info.
func (u *HtmlUpsertBulk) Update(set func(*HtmlUpsert)) *HtmlUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HtmlUpsert{UpdateSet: update})
	}))
	return u
}

// SetAnalyzedAt sets the "analyzed_at" field.
func (u *HtmlUpsertBulk) SetAnalyzedAt(v time.Time) *HtmlUpsertBulk {
	return u.Update(func(s *HtmlUpsert) {
		s.SetAnalyzedAt(v)
	})
}

// UpdateAnalyzedAt sets the "analyzed_at" field to the value that was provided on create.
func (u *HtmlUpsertBulk) UpdateAnalyzedAt() *HtmlUpsertBulk {
	return u.Update(func(s *HtmlUpsert) {
		s.UpdateAnalyzedAt()
	})
}

// ClearAnalyzedAt clears the value of the "analyzed_at" field.
func (u *HtmlUpsertBulk) ClearAnalyzedAt() *HtmlUpsertBulk {
	return u.Update(func(s *HtmlUpsert) {
		s.ClearAnalyzedAt()
	})
}

// Exec executes the query.
func (u *HtmlUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the HTMLCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for HTMLCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HtmlUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
