// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wolodata/schema/ent/keyword"
	"github.com/wolodata/schema/ent/predicate"
)

// KeywordDelete is the builder for deleting a Keyword entity.
type KeywordDelete struct {
	config
	hooks    []Hook
	mutation *KeywordMutation
}

// Where appends a list predicates to the KeywordDelete builder.
func (kd *KeywordDelete) Where(ps ...predicate.Keyword) *KeywordDelete {
	kd.mutation.Where(ps...)
	return kd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (kd *KeywordDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, kd.sqlExec, kd.mutation, kd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (kd *KeywordDelete) ExecX(ctx context.Context) int {
	n, err := kd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (kd *KeywordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(keyword.Table, sqlgraph.NewFieldSpec(keyword.FieldID, field.TypeString))
	if ps := kd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, kd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	kd.mutation.done = true
	return affected, err
}

// KeywordDeleteOne is the builder for deleting a single Keyword entity.
type KeywordDeleteOne struct {
	kd *KeywordDelete
}

// Where appends a list predicates to the KeywordDelete builder.
func (kdo *KeywordDeleteOne) Where(ps ...predicate.Keyword) *KeywordDeleteOne {
	kdo.kd.mutation.Where(ps...)
	return kdo
}

// Exec executes the deletion query.
func (kdo *KeywordDeleteOne) Exec(ctx context.Context) error {
	n, err := kdo.kd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{keyword.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (kdo *KeywordDeleteOne) ExecX(ctx context.Context) {
	if err := kdo.Exec(ctx); err != nil {
		panic(err)
	}
}
