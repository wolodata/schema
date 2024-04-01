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
	"github.com/wolodata/schema/ent/keyword"
	"github.com/wolodata/schema/ent/predicate"
)

// KeywordUpdate is the builder for updating Keyword entities.
type KeywordUpdate struct {
	config
	hooks    []Hook
	mutation *KeywordMutation
}

// Where appends a list predicates to the KeywordUpdate builder.
func (ku *KeywordUpdate) Where(ps ...predicate.Keyword) *KeywordUpdate {
	ku.mutation.Where(ps...)
	return ku
}

// SetName sets the "name" field.
func (ku *KeywordUpdate) SetName(s string) *KeywordUpdate {
	ku.mutation.SetName(s)
	return ku
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ku *KeywordUpdate) SetNillableName(s *string) *KeywordUpdate {
	if s != nil {
		ku.SetName(*s)
	}
	return ku
}

// SetWords sets the "words" field.
func (ku *KeywordUpdate) SetWords(s []string) *KeywordUpdate {
	ku.mutation.SetWords(s)
	return ku
}

// AppendWords appends s to the "words" field.
func (ku *KeywordUpdate) AppendWords(s []string) *KeywordUpdate {
	ku.mutation.AppendWords(s)
	return ku
}

// SetColor sets the "color" field.
func (ku *KeywordUpdate) SetColor(s string) *KeywordUpdate {
	ku.mutation.SetColor(s)
	return ku
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (ku *KeywordUpdate) SetNillableColor(s *string) *KeywordUpdate {
	if s != nil {
		ku.SetColor(*s)
	}
	return ku
}

// SetOrder sets the "order" field.
func (ku *KeywordUpdate) SetOrder(u uint64) *KeywordUpdate {
	ku.mutation.ResetOrder()
	ku.mutation.SetOrder(u)
	return ku
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (ku *KeywordUpdate) SetNillableOrder(u *uint64) *KeywordUpdate {
	if u != nil {
		ku.SetOrder(*u)
	}
	return ku
}

// AddOrder adds u to the "order" field.
func (ku *KeywordUpdate) AddOrder(u int64) *KeywordUpdate {
	ku.mutation.AddOrder(u)
	return ku
}

// Mutation returns the KeywordMutation object of the builder.
func (ku *KeywordUpdate) Mutation() *KeywordMutation {
	return ku.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ku *KeywordUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ku.sqlSave, ku.mutation, ku.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ku *KeywordUpdate) SaveX(ctx context.Context) int {
	affected, err := ku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ku *KeywordUpdate) Exec(ctx context.Context) error {
	_, err := ku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ku *KeywordUpdate) ExecX(ctx context.Context) {
	if err := ku.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ku *KeywordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(keyword.Table, keyword.Columns, sqlgraph.NewFieldSpec(keyword.FieldID, field.TypeString))
	if ps := ku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ku.mutation.Name(); ok {
		_spec.SetField(keyword.FieldName, field.TypeString, value)
	}
	if value, ok := ku.mutation.Words(); ok {
		_spec.SetField(keyword.FieldWords, field.TypeJSON, value)
	}
	if value, ok := ku.mutation.AppendedWords(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, keyword.FieldWords, value)
		})
	}
	if value, ok := ku.mutation.Color(); ok {
		_spec.SetField(keyword.FieldColor, field.TypeString, value)
	}
	if value, ok := ku.mutation.Order(); ok {
		_spec.SetField(keyword.FieldOrder, field.TypeUint64, value)
	}
	if value, ok := ku.mutation.AddedOrder(); ok {
		_spec.AddField(keyword.FieldOrder, field.TypeUint64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{keyword.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ku.mutation.done = true
	return n, nil
}

// KeywordUpdateOne is the builder for updating a single Keyword entity.
type KeywordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *KeywordMutation
}

// SetName sets the "name" field.
func (kuo *KeywordUpdateOne) SetName(s string) *KeywordUpdateOne {
	kuo.mutation.SetName(s)
	return kuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (kuo *KeywordUpdateOne) SetNillableName(s *string) *KeywordUpdateOne {
	if s != nil {
		kuo.SetName(*s)
	}
	return kuo
}

// SetWords sets the "words" field.
func (kuo *KeywordUpdateOne) SetWords(s []string) *KeywordUpdateOne {
	kuo.mutation.SetWords(s)
	return kuo
}

// AppendWords appends s to the "words" field.
func (kuo *KeywordUpdateOne) AppendWords(s []string) *KeywordUpdateOne {
	kuo.mutation.AppendWords(s)
	return kuo
}

// SetColor sets the "color" field.
func (kuo *KeywordUpdateOne) SetColor(s string) *KeywordUpdateOne {
	kuo.mutation.SetColor(s)
	return kuo
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (kuo *KeywordUpdateOne) SetNillableColor(s *string) *KeywordUpdateOne {
	if s != nil {
		kuo.SetColor(*s)
	}
	return kuo
}

// SetOrder sets the "order" field.
func (kuo *KeywordUpdateOne) SetOrder(u uint64) *KeywordUpdateOne {
	kuo.mutation.ResetOrder()
	kuo.mutation.SetOrder(u)
	return kuo
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (kuo *KeywordUpdateOne) SetNillableOrder(u *uint64) *KeywordUpdateOne {
	if u != nil {
		kuo.SetOrder(*u)
	}
	return kuo
}

// AddOrder adds u to the "order" field.
func (kuo *KeywordUpdateOne) AddOrder(u int64) *KeywordUpdateOne {
	kuo.mutation.AddOrder(u)
	return kuo
}

// Mutation returns the KeywordMutation object of the builder.
func (kuo *KeywordUpdateOne) Mutation() *KeywordMutation {
	return kuo.mutation
}

// Where appends a list predicates to the KeywordUpdate builder.
func (kuo *KeywordUpdateOne) Where(ps ...predicate.Keyword) *KeywordUpdateOne {
	kuo.mutation.Where(ps...)
	return kuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kuo *KeywordUpdateOne) Select(field string, fields ...string) *KeywordUpdateOne {
	kuo.fields = append([]string{field}, fields...)
	return kuo
}

// Save executes the query and returns the updated Keyword entity.
func (kuo *KeywordUpdateOne) Save(ctx context.Context) (*Keyword, error) {
	return withHooks(ctx, kuo.sqlSave, kuo.mutation, kuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (kuo *KeywordUpdateOne) SaveX(ctx context.Context) *Keyword {
	node, err := kuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kuo *KeywordUpdateOne) Exec(ctx context.Context) error {
	_, err := kuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kuo *KeywordUpdateOne) ExecX(ctx context.Context) {
	if err := kuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (kuo *KeywordUpdateOne) sqlSave(ctx context.Context) (_node *Keyword, err error) {
	_spec := sqlgraph.NewUpdateSpec(keyword.Table, keyword.Columns, sqlgraph.NewFieldSpec(keyword.FieldID, field.TypeString))
	id, ok := kuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Keyword.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := kuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, keyword.FieldID)
		for _, f := range fields {
			if !keyword.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != keyword.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kuo.mutation.Name(); ok {
		_spec.SetField(keyword.FieldName, field.TypeString, value)
	}
	if value, ok := kuo.mutation.Words(); ok {
		_spec.SetField(keyword.FieldWords, field.TypeJSON, value)
	}
	if value, ok := kuo.mutation.AppendedWords(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, keyword.FieldWords, value)
		})
	}
	if value, ok := kuo.mutation.Color(); ok {
		_spec.SetField(keyword.FieldColor, field.TypeString, value)
	}
	if value, ok := kuo.mutation.Order(); ok {
		_spec.SetField(keyword.FieldOrder, field.TypeUint64, value)
	}
	if value, ok := kuo.mutation.AddedOrder(); ok {
		_spec.AddField(keyword.FieldOrder, field.TypeUint64, value)
	}
	_node = &Keyword{config: kuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{keyword.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	kuo.mutation.done = true
	return _node, nil
}
