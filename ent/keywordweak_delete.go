// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wolodata/schema/ent/keywordweak"
	"github.com/wolodata/schema/ent/predicate"
)

// KeywordWeakDelete is the builder for deleting a KeywordWeak entity.
type KeywordWeakDelete struct {
	config
	hooks    []Hook
	mutation *KeywordWeakMutation
}

// Where appends a list predicates to the KeywordWeakDelete builder.
func (kwd *KeywordWeakDelete) Where(ps ...predicate.KeywordWeak) *KeywordWeakDelete {
	kwd.mutation.Where(ps...)
	return kwd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (kwd *KeywordWeakDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, kwd.sqlExec, kwd.mutation, kwd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (kwd *KeywordWeakDelete) ExecX(ctx context.Context) int {
	n, err := kwd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (kwd *KeywordWeakDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(keywordweak.Table, sqlgraph.NewFieldSpec(keywordweak.FieldID, field.TypeString))
	if ps := kwd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, kwd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	kwd.mutation.done = true
	return affected, err
}

// KeywordWeakDeleteOne is the builder for deleting a single KeywordWeak entity.
type KeywordWeakDeleteOne struct {
	kwd *KeywordWeakDelete
}

// Where appends a list predicates to the KeywordWeakDelete builder.
func (kwdo *KeywordWeakDeleteOne) Where(ps ...predicate.KeywordWeak) *KeywordWeakDeleteOne {
	kwdo.kwd.mutation.Where(ps...)
	return kwdo
}

// Exec executes the deletion query.
func (kwdo *KeywordWeakDeleteOne) Exec(ctx context.Context) error {
	n, err := kwdo.kwd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{keywordweak.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (kwdo *KeywordWeakDeleteOne) ExecX(ctx context.Context) {
	if err := kwdo.Exec(ctx); err != nil {
		panic(err)
	}
}
