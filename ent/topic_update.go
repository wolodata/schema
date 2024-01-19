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
	"github.com/wolodata/schema/ent/predicate"
	"github.com/wolodata/schema/ent/topic"
)

// TopicUpdate is the builder for updating Topic entities.
type TopicUpdate struct {
	config
	hooks    []Hook
	mutation *TopicMutation
}

// Where appends a list predicates to the TopicUpdate builder.
func (tu *TopicUpdate) Where(ps ...predicate.Topic) *TopicUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetKeyword sets the "keyword" field.
func (tu *TopicUpdate) SetKeyword(s string) *TopicUpdate {
	tu.mutation.SetKeyword(s)
	return tu
}

// SetNillableKeyword sets the "keyword" field if the given value is not nil.
func (tu *TopicUpdate) SetNillableKeyword(s *string) *TopicUpdate {
	if s != nil {
		tu.SetKeyword(*s)
	}
	return tu
}

// SetFollowTitle sets the "follow_title" field.
func (tu *TopicUpdate) SetFollowTitle(b bool) *TopicUpdate {
	tu.mutation.SetFollowTitle(b)
	return tu
}

// SetNillableFollowTitle sets the "follow_title" field if the given value is not nil.
func (tu *TopicUpdate) SetNillableFollowTitle(b *bool) *TopicUpdate {
	if b != nil {
		tu.SetFollowTitle(*b)
	}
	return tu
}

// SetFollowContent sets the "follow_content" field.
func (tu *TopicUpdate) SetFollowContent(b bool) *TopicUpdate {
	tu.mutation.SetFollowContent(b)
	return tu
}

// SetNillableFollowContent sets the "follow_content" field if the given value is not nil.
func (tu *TopicUpdate) SetNillableFollowContent(b *bool) *TopicUpdate {
	if b != nil {
		tu.SetFollowContent(*b)
	}
	return tu
}

// SetUserIds sets the "user_ids" field.
func (tu *TopicUpdate) SetUserIds(i []int) *TopicUpdate {
	tu.mutation.SetUserIds(i)
	return tu
}

// AppendUserIds appends i to the "user_ids" field.
func (tu *TopicUpdate) AppendUserIds(i []int) *TopicUpdate {
	tu.mutation.AppendUserIds(i)
	return tu
}

// Mutation returns the TopicMutation object of the builder.
func (tu *TopicUpdate) Mutation() *TopicMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TopicUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TopicUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TopicUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TopicUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TopicUpdate) check() error {
	if v, ok := tu.mutation.Keyword(); ok {
		if err := topic.KeywordValidator(v); err != nil {
			return &ValidationError{Name: "keyword", err: fmt.Errorf(`ent: validator failed for field "Topic.keyword": %w`, err)}
		}
	}
	return nil
}

func (tu *TopicUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(topic.Table, topic.Columns, sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Keyword(); ok {
		_spec.SetField(topic.FieldKeyword, field.TypeString, value)
	}
	if value, ok := tu.mutation.FollowTitle(); ok {
		_spec.SetField(topic.FieldFollowTitle, field.TypeBool, value)
	}
	if value, ok := tu.mutation.FollowContent(); ok {
		_spec.SetField(topic.FieldFollowContent, field.TypeBool, value)
	}
	if value, ok := tu.mutation.UserIds(); ok {
		_spec.SetField(topic.FieldUserIds, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.AppendedUserIds(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, topic.FieldUserIds, value)
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TopicUpdateOne is the builder for updating a single Topic entity.
type TopicUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TopicMutation
}

// SetKeyword sets the "keyword" field.
func (tuo *TopicUpdateOne) SetKeyword(s string) *TopicUpdateOne {
	tuo.mutation.SetKeyword(s)
	return tuo
}

// SetNillableKeyword sets the "keyword" field if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableKeyword(s *string) *TopicUpdateOne {
	if s != nil {
		tuo.SetKeyword(*s)
	}
	return tuo
}

// SetFollowTitle sets the "follow_title" field.
func (tuo *TopicUpdateOne) SetFollowTitle(b bool) *TopicUpdateOne {
	tuo.mutation.SetFollowTitle(b)
	return tuo
}

// SetNillableFollowTitle sets the "follow_title" field if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableFollowTitle(b *bool) *TopicUpdateOne {
	if b != nil {
		tuo.SetFollowTitle(*b)
	}
	return tuo
}

// SetFollowContent sets the "follow_content" field.
func (tuo *TopicUpdateOne) SetFollowContent(b bool) *TopicUpdateOne {
	tuo.mutation.SetFollowContent(b)
	return tuo
}

// SetNillableFollowContent sets the "follow_content" field if the given value is not nil.
func (tuo *TopicUpdateOne) SetNillableFollowContent(b *bool) *TopicUpdateOne {
	if b != nil {
		tuo.SetFollowContent(*b)
	}
	return tuo
}

// SetUserIds sets the "user_ids" field.
func (tuo *TopicUpdateOne) SetUserIds(i []int) *TopicUpdateOne {
	tuo.mutation.SetUserIds(i)
	return tuo
}

// AppendUserIds appends i to the "user_ids" field.
func (tuo *TopicUpdateOne) AppendUserIds(i []int) *TopicUpdateOne {
	tuo.mutation.AppendUserIds(i)
	return tuo
}

// Mutation returns the TopicMutation object of the builder.
func (tuo *TopicUpdateOne) Mutation() *TopicMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TopicUpdate builder.
func (tuo *TopicUpdateOne) Where(ps ...predicate.Topic) *TopicUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TopicUpdateOne) Select(field string, fields ...string) *TopicUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Topic entity.
func (tuo *TopicUpdateOne) Save(ctx context.Context) (*Topic, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TopicUpdateOne) SaveX(ctx context.Context) *Topic {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TopicUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TopicUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TopicUpdateOne) check() error {
	if v, ok := tuo.mutation.Keyword(); ok {
		if err := topic.KeywordValidator(v); err != nil {
			return &ValidationError{Name: "keyword", err: fmt.Errorf(`ent: validator failed for field "Topic.keyword": %w`, err)}
		}
	}
	return nil
}

func (tuo *TopicUpdateOne) sqlSave(ctx context.Context) (_node *Topic, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(topic.Table, topic.Columns, sqlgraph.NewFieldSpec(topic.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Topic.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, topic.FieldID)
		for _, f := range fields {
			if !topic.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != topic.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Keyword(); ok {
		_spec.SetField(topic.FieldKeyword, field.TypeString, value)
	}
	if value, ok := tuo.mutation.FollowTitle(); ok {
		_spec.SetField(topic.FieldFollowTitle, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.FollowContent(); ok {
		_spec.SetField(topic.FieldFollowContent, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.UserIds(); ok {
		_spec.SetField(topic.FieldUserIds, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.AppendedUserIds(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, topic.FieldUserIds, value)
		})
	}
	_node = &Topic{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
