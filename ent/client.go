// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/wolodata/schema/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/wolodata/schema/ent/article"
	"github.com/wolodata/schema/ent/html"
	"github.com/wolodata/schema/ent/report"
	"github.com/wolodata/schema/ent/topic"
	"github.com/wolodata/schema/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Article is the client for interacting with the Article builders.
	Article *ArticleClient
	// Html is the client for interacting with the Html builders.
	Html *HTMLClient
	// Report is the client for interacting with the Report builders.
	Report *ReportClient
	// Topic is the client for interacting with the Topic builders.
	Topic *TopicClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Article = NewArticleClient(c.config)
	c.Html = NewHTMLClient(c.config)
	c.Report = NewReportClient(c.config)
	c.Topic = NewTopicClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Article: NewArticleClient(cfg),
		Html:    NewHTMLClient(cfg),
		Report:  NewReportClient(cfg),
		Topic:   NewTopicClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Article: NewArticleClient(cfg),
		Html:    NewHTMLClient(cfg),
		Report:  NewReportClient(cfg),
		Topic:   NewTopicClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Article.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Article.Use(hooks...)
	c.Html.Use(hooks...)
	c.Report.Use(hooks...)
	c.Topic.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Article.Intercept(interceptors...)
	c.Html.Intercept(interceptors...)
	c.Report.Intercept(interceptors...)
	c.Topic.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ArticleMutation:
		return c.Article.mutate(ctx, m)
	case *HTMLMutation:
		return c.Html.mutate(ctx, m)
	case *ReportMutation:
		return c.Report.mutate(ctx, m)
	case *TopicMutation:
		return c.Topic.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ArticleClient is a client for the Article schema.
type ArticleClient struct {
	config
}

// NewArticleClient returns a client for the Article from the given config.
func NewArticleClient(c config) *ArticleClient {
	return &ArticleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `article.Hooks(f(g(h())))`.
func (c *ArticleClient) Use(hooks ...Hook) {
	c.hooks.Article = append(c.hooks.Article, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `article.Intercept(f(g(h())))`.
func (c *ArticleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Article = append(c.inters.Article, interceptors...)
}

// Create returns a builder for creating a Article entity.
func (c *ArticleClient) Create() *ArticleCreate {
	mutation := newArticleMutation(c.config, OpCreate)
	return &ArticleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Article entities.
func (c *ArticleClient) CreateBulk(builders ...*ArticleCreate) *ArticleCreateBulk {
	return &ArticleCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ArticleClient) MapCreateBulk(slice any, setFunc func(*ArticleCreate, int)) *ArticleCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ArticleCreateBulk{err: fmt.Errorf("calling to ArticleClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ArticleCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ArticleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Article.
func (c *ArticleClient) Update() *ArticleUpdate {
	mutation := newArticleMutation(c.config, OpUpdate)
	return &ArticleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArticleClient) UpdateOne(a *Article) *ArticleUpdateOne {
	mutation := newArticleMutation(c.config, OpUpdateOne, withArticle(a))
	return &ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArticleClient) UpdateOneID(id uint64) *ArticleUpdateOne {
	mutation := newArticleMutation(c.config, OpUpdateOne, withArticleID(id))
	return &ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Article.
func (c *ArticleClient) Delete() *ArticleDelete {
	mutation := newArticleMutation(c.config, OpDelete)
	return &ArticleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ArticleClient) DeleteOne(a *Article) *ArticleDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ArticleClient) DeleteOneID(id uint64) *ArticleDeleteOne {
	builder := c.Delete().Where(article.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArticleDeleteOne{builder}
}

// Query returns a query builder for Article.
func (c *ArticleClient) Query() *ArticleQuery {
	return &ArticleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeArticle},
		inters: c.Interceptors(),
	}
}

// Get returns a Article entity by its id.
func (c *ArticleClient) Get(ctx context.Context, id uint64) (*Article, error) {
	return c.Query().Where(article.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArticleClient) GetX(ctx context.Context, id uint64) *Article {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ArticleClient) Hooks() []Hook {
	return c.hooks.Article
}

// Interceptors returns the client interceptors.
func (c *ArticleClient) Interceptors() []Interceptor {
	return c.inters.Article
}

func (c *ArticleClient) mutate(ctx context.Context, m *ArticleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ArticleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ArticleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ArticleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Article mutation op: %q", m.Op())
	}
}

// HTMLClient is a client for the Html schema.
type HTMLClient struct {
	config
}

// NewHTMLClient returns a client for the Html from the given config.
func NewHTMLClient(c config) *HTMLClient {
	return &HTMLClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `html.Hooks(f(g(h())))`.
func (c *HTMLClient) Use(hooks ...Hook) {
	c.hooks.Html = append(c.hooks.Html, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `html.Intercept(f(g(h())))`.
func (c *HTMLClient) Intercept(interceptors ...Interceptor) {
	c.inters.Html = append(c.inters.Html, interceptors...)
}

// Create returns a builder for creating a Html entity.
func (c *HTMLClient) Create() *HTMLCreate {
	mutation := newHTMLMutation(c.config, OpCreate)
	return &HTMLCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Html entities.
func (c *HTMLClient) CreateBulk(builders ...*HTMLCreate) *HTMLCreateBulk {
	return &HTMLCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *HTMLClient) MapCreateBulk(slice any, setFunc func(*HTMLCreate, int)) *HTMLCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &HTMLCreateBulk{err: fmt.Errorf("calling to HTMLClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*HTMLCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &HTMLCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Html.
func (c *HTMLClient) Update() *HTMLUpdate {
	mutation := newHTMLMutation(c.config, OpUpdate)
	return &HTMLUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HTMLClient) UpdateOne(h *Html) *HTMLUpdateOne {
	mutation := newHTMLMutation(c.config, OpUpdateOne, withHtml(h))
	return &HTMLUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HTMLClient) UpdateOneID(id uint64) *HTMLUpdateOne {
	mutation := newHTMLMutation(c.config, OpUpdateOne, withHtmlID(id))
	return &HTMLUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Html.
func (c *HTMLClient) Delete() *HTMLDelete {
	mutation := newHTMLMutation(c.config, OpDelete)
	return &HTMLDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *HTMLClient) DeleteOne(h *Html) *HTMLDeleteOne {
	return c.DeleteOneID(h.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *HTMLClient) DeleteOneID(id uint64) *HTMLDeleteOne {
	builder := c.Delete().Where(html.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HTMLDeleteOne{builder}
}

// Query returns a query builder for Html.
func (c *HTMLClient) Query() *HTMLQuery {
	return &HTMLQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeHTML},
		inters: c.Interceptors(),
	}
}

// Get returns a Html entity by its id.
func (c *HTMLClient) Get(ctx context.Context, id uint64) (*Html, error) {
	return c.Query().Where(html.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HTMLClient) GetX(ctx context.Context, id uint64) *Html {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *HTMLClient) Hooks() []Hook {
	return c.hooks.Html
}

// Interceptors returns the client interceptors.
func (c *HTMLClient) Interceptors() []Interceptor {
	return c.inters.Html
}

func (c *HTMLClient) mutate(ctx context.Context, m *HTMLMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&HTMLCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&HTMLUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&HTMLUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&HTMLDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Html mutation op: %q", m.Op())
	}
}

// ReportClient is a client for the Report schema.
type ReportClient struct {
	config
}

// NewReportClient returns a client for the Report from the given config.
func NewReportClient(c config) *ReportClient {
	return &ReportClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `report.Hooks(f(g(h())))`.
func (c *ReportClient) Use(hooks ...Hook) {
	c.hooks.Report = append(c.hooks.Report, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `report.Intercept(f(g(h())))`.
func (c *ReportClient) Intercept(interceptors ...Interceptor) {
	c.inters.Report = append(c.inters.Report, interceptors...)
}

// Create returns a builder for creating a Report entity.
func (c *ReportClient) Create() *ReportCreate {
	mutation := newReportMutation(c.config, OpCreate)
	return &ReportCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Report entities.
func (c *ReportClient) CreateBulk(builders ...*ReportCreate) *ReportCreateBulk {
	return &ReportCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ReportClient) MapCreateBulk(slice any, setFunc func(*ReportCreate, int)) *ReportCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ReportCreateBulk{err: fmt.Errorf("calling to ReportClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ReportCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ReportCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Report.
func (c *ReportClient) Update() *ReportUpdate {
	mutation := newReportMutation(c.config, OpUpdate)
	return &ReportUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReportClient) UpdateOne(r *Report) *ReportUpdateOne {
	mutation := newReportMutation(c.config, OpUpdateOne, withReport(r))
	return &ReportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReportClient) UpdateOneID(id uint64) *ReportUpdateOne {
	mutation := newReportMutation(c.config, OpUpdateOne, withReportID(id))
	return &ReportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Report.
func (c *ReportClient) Delete() *ReportDelete {
	mutation := newReportMutation(c.config, OpDelete)
	return &ReportDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ReportClient) DeleteOne(r *Report) *ReportDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ReportClient) DeleteOneID(id uint64) *ReportDeleteOne {
	builder := c.Delete().Where(report.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReportDeleteOne{builder}
}

// Query returns a query builder for Report.
func (c *ReportClient) Query() *ReportQuery {
	return &ReportQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeReport},
		inters: c.Interceptors(),
	}
}

// Get returns a Report entity by its id.
func (c *ReportClient) Get(ctx context.Context, id uint64) (*Report, error) {
	return c.Query().Where(report.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReportClient) GetX(ctx context.Context, id uint64) *Report {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ReportClient) Hooks() []Hook {
	return c.hooks.Report
}

// Interceptors returns the client interceptors.
func (c *ReportClient) Interceptors() []Interceptor {
	return c.inters.Report
}

func (c *ReportClient) mutate(ctx context.Context, m *ReportMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ReportCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ReportUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ReportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ReportDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Report mutation op: %q", m.Op())
	}
}

// TopicClient is a client for the Topic schema.
type TopicClient struct {
	config
}

// NewTopicClient returns a client for the Topic from the given config.
func NewTopicClient(c config) *TopicClient {
	return &TopicClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `topic.Hooks(f(g(h())))`.
func (c *TopicClient) Use(hooks ...Hook) {
	c.hooks.Topic = append(c.hooks.Topic, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `topic.Intercept(f(g(h())))`.
func (c *TopicClient) Intercept(interceptors ...Interceptor) {
	c.inters.Topic = append(c.inters.Topic, interceptors...)
}

// Create returns a builder for creating a Topic entity.
func (c *TopicClient) Create() *TopicCreate {
	mutation := newTopicMutation(c.config, OpCreate)
	return &TopicCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Topic entities.
func (c *TopicClient) CreateBulk(builders ...*TopicCreate) *TopicCreateBulk {
	return &TopicCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TopicClient) MapCreateBulk(slice any, setFunc func(*TopicCreate, int)) *TopicCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TopicCreateBulk{err: fmt.Errorf("calling to TopicClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TopicCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TopicCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Topic.
func (c *TopicClient) Update() *TopicUpdate {
	mutation := newTopicMutation(c.config, OpUpdate)
	return &TopicUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TopicClient) UpdateOne(t *Topic) *TopicUpdateOne {
	mutation := newTopicMutation(c.config, OpUpdateOne, withTopic(t))
	return &TopicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TopicClient) UpdateOneID(id uint64) *TopicUpdateOne {
	mutation := newTopicMutation(c.config, OpUpdateOne, withTopicID(id))
	return &TopicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Topic.
func (c *TopicClient) Delete() *TopicDelete {
	mutation := newTopicMutation(c.config, OpDelete)
	return &TopicDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TopicClient) DeleteOne(t *Topic) *TopicDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TopicClient) DeleteOneID(id uint64) *TopicDeleteOne {
	builder := c.Delete().Where(topic.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TopicDeleteOne{builder}
}

// Query returns a query builder for Topic.
func (c *TopicClient) Query() *TopicQuery {
	return &TopicQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTopic},
		inters: c.Interceptors(),
	}
}

// Get returns a Topic entity by its id.
func (c *TopicClient) Get(ctx context.Context, id uint64) (*Topic, error) {
	return c.Query().Where(topic.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TopicClient) GetX(ctx context.Context, id uint64) *Topic {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TopicClient) Hooks() []Hook {
	return c.hooks.Topic
}

// Interceptors returns the client interceptors.
func (c *TopicClient) Interceptors() []Interceptor {
	return c.inters.Topic
}

func (c *TopicClient) mutate(ctx context.Context, m *TopicMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TopicCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TopicUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TopicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TopicDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Topic mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uint64) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uint64) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uint64) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uint64) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Article, Html, Report, Topic, User []ent.Hook
	}
	inters struct {
		Article, Html, Report, Topic, User []ent.Interceptor
	}
)
