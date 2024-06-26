// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wolodata/schema/ent/brain"
)

// BrainCreate is the builder for creating a Brain entity.
type BrainCreate struct {
	config
	mutation *BrainMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUserID sets the "user_id" field.
func (bc *BrainCreate) SetUserID(s string) *BrainCreate {
	bc.mutation.SetUserID(s)
	return bc
}

// SetQuestion sets the "question" field.
func (bc *BrainCreate) SetQuestion(s string) *BrainCreate {
	bc.mutation.SetQuestion(s)
	return bc
}

// SetAnswer sets the "answer" field.
func (bc *BrainCreate) SetAnswer(s string) *BrainCreate {
	bc.mutation.SetAnswer(s)
	return bc
}

// SetID sets the "id" field.
func (bc *BrainCreate) SetID(s string) *BrainCreate {
	bc.mutation.SetID(s)
	return bc
}

// Mutation returns the BrainMutation object of the builder.
func (bc *BrainCreate) Mutation() *BrainMutation {
	return bc.mutation
}

// Save creates the Brain in the database.
func (bc *BrainCreate) Save(ctx context.Context) (*Brain, error) {
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BrainCreate) SaveX(ctx context.Context) *Brain {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BrainCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BrainCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BrainCreate) check() error {
	if _, ok := bc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Brain.user_id"`)}
	}
	if _, ok := bc.mutation.Question(); !ok {
		return &ValidationError{Name: "question", err: errors.New(`ent: missing required field "Brain.question"`)}
	}
	if _, ok := bc.mutation.Answer(); !ok {
		return &ValidationError{Name: "answer", err: errors.New(`ent: missing required field "Brain.answer"`)}
	}
	return nil
}

func (bc *BrainCreate) sqlSave(ctx context.Context) (*Brain, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Brain.ID type: %T", _spec.ID.Value)
		}
	}
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BrainCreate) createSpec() (*Brain, *sqlgraph.CreateSpec) {
	var (
		_node = &Brain{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(brain.Table, sqlgraph.NewFieldSpec(brain.FieldID, field.TypeString))
	)
	_spec.OnConflict = bc.conflict
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.UserID(); ok {
		_spec.SetField(brain.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := bc.mutation.Question(); ok {
		_spec.SetField(brain.FieldQuestion, field.TypeString, value)
		_node.Question = value
	}
	if value, ok := bc.mutation.Answer(); ok {
		_spec.SetField(brain.FieldAnswer, field.TypeString, value)
		_node.Answer = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Brain.Create().
//		SetUserID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BrainUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
func (bc *BrainCreate) OnConflict(opts ...sql.ConflictOption) *BrainUpsertOne {
	bc.conflict = opts
	return &BrainUpsertOne{
		create: bc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Brain.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bc *BrainCreate) OnConflictColumns(columns ...string) *BrainUpsertOne {
	bc.conflict = append(bc.conflict, sql.ConflictColumns(columns...))
	return &BrainUpsertOne{
		create: bc,
	}
}

type (
	// BrainUpsertOne is the builder for "upsert"-ing
	//  one Brain node.
	BrainUpsertOne struct {
		create *BrainCreate
	}

	// BrainUpsert is the "OnConflict" setter.
	BrainUpsert struct {
		*sql.UpdateSet
	}
)

// SetQuestion sets the "question" field.
func (u *BrainUpsert) SetQuestion(v string) *BrainUpsert {
	u.Set(brain.FieldQuestion, v)
	return u
}

// UpdateQuestion sets the "question" field to the value that was provided on create.
func (u *BrainUpsert) UpdateQuestion() *BrainUpsert {
	u.SetExcluded(brain.FieldQuestion)
	return u
}

// SetAnswer sets the "answer" field.
func (u *BrainUpsert) SetAnswer(v string) *BrainUpsert {
	u.Set(brain.FieldAnswer, v)
	return u
}

// UpdateAnswer sets the "answer" field to the value that was provided on create.
func (u *BrainUpsert) UpdateAnswer() *BrainUpsert {
	u.SetExcluded(brain.FieldAnswer)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Brain.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(brain.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BrainUpsertOne) UpdateNewValues() *BrainUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(brain.FieldID)
		}
		if _, exists := u.create.mutation.UserID(); exists {
			s.SetIgnore(brain.FieldUserID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Brain.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *BrainUpsertOne) Ignore() *BrainUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BrainUpsertOne) DoNothing() *BrainUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BrainCreate.OnConflict
// documentation for more info.
func (u *BrainUpsertOne) Update(set func(*BrainUpsert)) *BrainUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BrainUpsert{UpdateSet: update})
	}))
	return u
}

// SetQuestion sets the "question" field.
func (u *BrainUpsertOne) SetQuestion(v string) *BrainUpsertOne {
	return u.Update(func(s *BrainUpsert) {
		s.SetQuestion(v)
	})
}

// UpdateQuestion sets the "question" field to the value that was provided on create.
func (u *BrainUpsertOne) UpdateQuestion() *BrainUpsertOne {
	return u.Update(func(s *BrainUpsert) {
		s.UpdateQuestion()
	})
}

// SetAnswer sets the "answer" field.
func (u *BrainUpsertOne) SetAnswer(v string) *BrainUpsertOne {
	return u.Update(func(s *BrainUpsert) {
		s.SetAnswer(v)
	})
}

// UpdateAnswer sets the "answer" field to the value that was provided on create.
func (u *BrainUpsertOne) UpdateAnswer() *BrainUpsertOne {
	return u.Update(func(s *BrainUpsert) {
		s.UpdateAnswer()
	})
}

// Exec executes the query.
func (u *BrainUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BrainCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BrainUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BrainUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: BrainUpsertOne.ID is not supported by MySQL driver. Use BrainUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BrainUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BrainCreateBulk is the builder for creating many Brain entities in bulk.
type BrainCreateBulk struct {
	config
	err      error
	builders []*BrainCreate
	conflict []sql.ConflictOption
}

// Save creates the Brain entities in the database.
func (bcb *BrainCreateBulk) Save(ctx context.Context) ([]*Brain, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Brain, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BrainMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = bcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BrainCreateBulk) SaveX(ctx context.Context) []*Brain {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BrainCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BrainCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Brain.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BrainUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
func (bcb *BrainCreateBulk) OnConflict(opts ...sql.ConflictOption) *BrainUpsertBulk {
	bcb.conflict = opts
	return &BrainUpsertBulk{
		create: bcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Brain.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bcb *BrainCreateBulk) OnConflictColumns(columns ...string) *BrainUpsertBulk {
	bcb.conflict = append(bcb.conflict, sql.ConflictColumns(columns...))
	return &BrainUpsertBulk{
		create: bcb,
	}
}

// BrainUpsertBulk is the builder for "upsert"-ing
// a bulk of Brain nodes.
type BrainUpsertBulk struct {
	create *BrainCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Brain.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(brain.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BrainUpsertBulk) UpdateNewValues() *BrainUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(brain.FieldID)
			}
			if _, exists := b.mutation.UserID(); exists {
				s.SetIgnore(brain.FieldUserID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Brain.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *BrainUpsertBulk) Ignore() *BrainUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BrainUpsertBulk) DoNothing() *BrainUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BrainCreateBulk.OnConflict
// documentation for more info.
func (u *BrainUpsertBulk) Update(set func(*BrainUpsert)) *BrainUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BrainUpsert{UpdateSet: update})
	}))
	return u
}

// SetQuestion sets the "question" field.
func (u *BrainUpsertBulk) SetQuestion(v string) *BrainUpsertBulk {
	return u.Update(func(s *BrainUpsert) {
		s.SetQuestion(v)
	})
}

// UpdateQuestion sets the "question" field to the value that was provided on create.
func (u *BrainUpsertBulk) UpdateQuestion() *BrainUpsertBulk {
	return u.Update(func(s *BrainUpsert) {
		s.UpdateQuestion()
	})
}

// SetAnswer sets the "answer" field.
func (u *BrainUpsertBulk) SetAnswer(v string) *BrainUpsertBulk {
	return u.Update(func(s *BrainUpsert) {
		s.SetAnswer(v)
	})
}

// UpdateAnswer sets the "answer" field to the value that was provided on create.
func (u *BrainUpsertBulk) UpdateAnswer() *BrainUpsertBulk {
	return u.Update(func(s *BrainUpsert) {
		s.UpdateAnswer()
	})
}

// Exec executes the query.
func (u *BrainUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the BrainCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BrainCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BrainUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
