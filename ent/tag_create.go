// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wolodata/schema/ent/tag"
)

// TagCreate is the builder for creating a Tag entity.
type TagCreate struct {
	config
	mutation *TagMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEnglish sets the "english" field.
func (tc *TagCreate) SetEnglish(s string) *TagCreate {
	tc.mutation.SetEnglish(s)
	return tc
}

// SetChinese sets the "chinese" field.
func (tc *TagCreate) SetChinese(s string) *TagCreate {
	tc.mutation.SetChinese(s)
	return tc
}

// Mutation returns the TagMutation object of the builder.
func (tc *TagCreate) Mutation() *TagMutation {
	return tc.mutation
}

// Save creates the Tag in the database.
func (tc *TagCreate) Save(ctx context.Context) (*Tag, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TagCreate) SaveX(ctx context.Context) *Tag {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TagCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TagCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TagCreate) check() error {
	if _, ok := tc.mutation.English(); !ok {
		return &ValidationError{Name: "english", err: errors.New(`ent: missing required field "Tag.english"`)}
	}
	if v, ok := tc.mutation.English(); ok {
		if err := tag.EnglishValidator(v); err != nil {
			return &ValidationError{Name: "english", err: fmt.Errorf(`ent: validator failed for field "Tag.english": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Chinese(); !ok {
		return &ValidationError{Name: "chinese", err: errors.New(`ent: missing required field "Tag.chinese"`)}
	}
	if v, ok := tc.mutation.Chinese(); ok {
		if err := tag.ChineseValidator(v); err != nil {
			return &ValidationError{Name: "chinese", err: fmt.Errorf(`ent: validator failed for field "Tag.chinese": %w`, err)}
		}
	}
	return nil
}

func (tc *TagCreate) sqlSave(ctx context.Context) (*Tag, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TagCreate) createSpec() (*Tag, *sqlgraph.CreateSpec) {
	var (
		_node = &Tag{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tag.Table, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt))
	)
	_spec.OnConflict = tc.conflict
	if value, ok := tc.mutation.English(); ok {
		_spec.SetField(tag.FieldEnglish, field.TypeString, value)
		_node.English = value
	}
	if value, ok := tc.mutation.Chinese(); ok {
		_spec.SetField(tag.FieldChinese, field.TypeString, value)
		_node.Chinese = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Tag.Create().
//		SetEnglish(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TagUpsert) {
//			SetEnglish(v+v).
//		}).
//		Exec(ctx)
func (tc *TagCreate) OnConflict(opts ...sql.ConflictOption) *TagUpsertOne {
	tc.conflict = opts
	return &TagUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tag.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TagCreate) OnConflictColumns(columns ...string) *TagUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TagUpsertOne{
		create: tc,
	}
}

type (
	// TagUpsertOne is the builder for "upsert"-ing
	//  one Tag node.
	TagUpsertOne struct {
		create *TagCreate
	}

	// TagUpsert is the "OnConflict" setter.
	TagUpsert struct {
		*sql.UpdateSet
	}
)

// SetChinese sets the "chinese" field.
func (u *TagUpsert) SetChinese(v string) *TagUpsert {
	u.Set(tag.FieldChinese, v)
	return u
}

// UpdateChinese sets the "chinese" field to the value that was provided on create.
func (u *TagUpsert) UpdateChinese() *TagUpsert {
	u.SetExcluded(tag.FieldChinese)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Tag.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TagUpsertOne) UpdateNewValues() *TagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.English(); exists {
			s.SetIgnore(tag.FieldEnglish)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Tag.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TagUpsertOne) Ignore() *TagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TagUpsertOne) DoNothing() *TagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TagCreate.OnConflict
// documentation for more info.
func (u *TagUpsertOne) Update(set func(*TagUpsert)) *TagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TagUpsert{UpdateSet: update})
	}))
	return u
}

// SetChinese sets the "chinese" field.
func (u *TagUpsertOne) SetChinese(v string) *TagUpsertOne {
	return u.Update(func(s *TagUpsert) {
		s.SetChinese(v)
	})
}

// UpdateChinese sets the "chinese" field to the value that was provided on create.
func (u *TagUpsertOne) UpdateChinese() *TagUpsertOne {
	return u.Update(func(s *TagUpsert) {
		s.UpdateChinese()
	})
}

// Exec executes the query.
func (u *TagUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TagCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TagUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TagUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TagUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TagCreateBulk is the builder for creating many Tag entities in bulk.
type TagCreateBulk struct {
	config
	err      error
	builders []*TagCreate
	conflict []sql.ConflictOption
}

// Save creates the Tag entities in the database.
func (tcb *TagCreateBulk) Save(ctx context.Context) ([]*Tag, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tag, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TagMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TagCreateBulk) SaveX(ctx context.Context) []*Tag {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TagCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TagCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Tag.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TagUpsert) {
//			SetEnglish(v+v).
//		}).
//		Exec(ctx)
func (tcb *TagCreateBulk) OnConflict(opts ...sql.ConflictOption) *TagUpsertBulk {
	tcb.conflict = opts
	return &TagUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tag.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TagCreateBulk) OnConflictColumns(columns ...string) *TagUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TagUpsertBulk{
		create: tcb,
	}
}

// TagUpsertBulk is the builder for "upsert"-ing
// a bulk of Tag nodes.
type TagUpsertBulk struct {
	create *TagCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Tag.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TagUpsertBulk) UpdateNewValues() *TagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.English(); exists {
				s.SetIgnore(tag.FieldEnglish)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Tag.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TagUpsertBulk) Ignore() *TagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TagUpsertBulk) DoNothing() *TagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TagCreateBulk.OnConflict
// documentation for more info.
func (u *TagUpsertBulk) Update(set func(*TagUpsert)) *TagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TagUpsert{UpdateSet: update})
	}))
	return u
}

// SetChinese sets the "chinese" field.
func (u *TagUpsertBulk) SetChinese(v string) *TagUpsertBulk {
	return u.Update(func(s *TagUpsert) {
		s.SetChinese(v)
	})
}

// UpdateChinese sets the "chinese" field to the value that was provided on create.
func (u *TagUpsertBulk) UpdateChinese() *TagUpsertBulk {
	return u.Update(func(s *TagUpsert) {
		s.UpdateChinese()
	})
}

// Exec executes the query.
func (u *TagUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TagCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TagCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TagUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
