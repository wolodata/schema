// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wolodata/schema/ent/article"
	"github.com/wolodata/schema/ent/predicate"
)

// ArticleQuery is the builder for querying Article entities.
type ArticleQuery struct {
	config
	ctx        *QueryContext
	order      []article.OrderOption
	inters     []Interceptor
	predicates []predicate.Article
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ArticleQuery builder.
func (aq *ArticleQuery) Where(ps ...predicate.Article) *ArticleQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *ArticleQuery) Limit(limit int) *ArticleQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *ArticleQuery) Offset(offset int) *ArticleQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *ArticleQuery) Unique(unique bool) *ArticleQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *ArticleQuery) Order(o ...article.OrderOption) *ArticleQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// First returns the first Article entity from the query.
// Returns a *NotFoundError when no Article was found.
func (aq *ArticleQuery) First(ctx context.Context) (*Article, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{article.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *ArticleQuery) FirstX(ctx context.Context) *Article {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Article ID from the query.
// Returns a *NotFoundError when no Article ID was found.
func (aq *ArticleQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{article.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *ArticleQuery) FirstIDX(ctx context.Context) string {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Article entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Article entity is found.
// Returns a *NotFoundError when no Article entities are found.
func (aq *ArticleQuery) Only(ctx context.Context) (*Article, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{article.Label}
	default:
		return nil, &NotSingularError{article.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *ArticleQuery) OnlyX(ctx context.Context) *Article {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Article ID in the query.
// Returns a *NotSingularError when more than one Article ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *ArticleQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{article.Label}
	default:
		err = &NotSingularError{article.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *ArticleQuery) OnlyIDX(ctx context.Context) string {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Articles.
func (aq *ArticleQuery) All(ctx context.Context) ([]*Article, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Article, *ArticleQuery]()
	return withInterceptors[[]*Article](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *ArticleQuery) AllX(ctx context.Context) []*Article {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Article IDs.
func (aq *ArticleQuery) IDs(ctx context.Context) (ids []string, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(article.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *ArticleQuery) IDsX(ctx context.Context) []string {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *ArticleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*ArticleQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *ArticleQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *ArticleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *ArticleQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ArticleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *ArticleQuery) Clone() *ArticleQuery {
	if aq == nil {
		return nil
	}
	return &ArticleQuery{
		config:     aq.config,
		ctx:        aq.ctx.Clone(),
		order:      append([]article.OrderOption{}, aq.order...),
		inters:     append([]Interceptor{}, aq.inters...),
		predicates: append([]predicate.Article{}, aq.predicates...),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OriginShortID string `json:"origin_short_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Article.Query().
//		GroupBy(article.FieldOriginShortID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *ArticleQuery) GroupBy(field string, fields ...string) *ArticleGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ArticleGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = article.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OriginShortID string `json:"origin_short_id,omitempty"`
//	}
//
//	client.Article.Query().
//		Select(article.FieldOriginShortID).
//		Scan(ctx, &v)
func (aq *ArticleQuery) Select(fields ...string) *ArticleSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &ArticleSelect{ArticleQuery: aq}
	sbuild.label = article.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ArticleSelect configured with the given aggregations.
func (aq *ArticleQuery) Aggregate(fns ...AggregateFunc) *ArticleSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *ArticleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !article.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *ArticleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Article, error) {
	var (
		nodes = []*Article{}
		_spec = aq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Article).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Article{config: aq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aq *ArticleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *ArticleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(article.Table, article.Columns, sqlgraph.NewFieldSpec(article.FieldID, field.TypeString))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, article.FieldID)
		for i := range fields {
			if fields[i] != article.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *ArticleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(article.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = article.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ArticleGroupBy is the group-by builder for Article entities.
type ArticleGroupBy struct {
	selector
	build *ArticleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *ArticleGroupBy) Aggregate(fns ...AggregateFunc) *ArticleGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *ArticleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArticleQuery, *ArticleGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *ArticleGroupBy) sqlScan(ctx context.Context, root *ArticleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ArticleSelect is the builder for selecting fields of Article entities.
type ArticleSelect struct {
	*ArticleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *ArticleSelect) Aggregate(fns ...AggregateFunc) *ArticleSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *ArticleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArticleQuery, *ArticleSelect](ctx, as.ArticleQuery, as, as.inters, v)
}

func (as *ArticleSelect) sqlScan(ctx context.Context, root *ArticleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
