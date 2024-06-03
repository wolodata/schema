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
	"github.com/wolodata/schema/ent/systemconfig"
)

// SystemConfigCreate is the builder for creating a SystemConfig entity.
type SystemConfigCreate struct {
	config
	mutation *SystemConfigMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (scc *SystemConfigCreate) SetName(s string) *SystemConfigCreate {
	scc.mutation.SetName(s)
	return scc
}

// SetDescription sets the "description" field.
func (scc *SystemConfigCreate) SetDescription(s string) *SystemConfigCreate {
	scc.mutation.SetDescription(s)
	return scc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (scc *SystemConfigCreate) SetNillableDescription(s *string) *SystemConfigCreate {
	if s != nil {
		scc.SetDescription(*s)
	}
	return scc
}

// SetValue sets the "value" field.
func (scc *SystemConfigCreate) SetValue(s string) *SystemConfigCreate {
	scc.mutation.SetValue(s)
	return scc
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (scc *SystemConfigCreate) SetNillableValue(s *string) *SystemConfigCreate {
	if s != nil {
		scc.SetValue(*s)
	}
	return scc
}

// SetID sets the "id" field.
func (scc *SystemConfigCreate) SetID(s string) *SystemConfigCreate {
	scc.mutation.SetID(s)
	return scc
}

// Mutation returns the SystemConfigMutation object of the builder.
func (scc *SystemConfigCreate) Mutation() *SystemConfigMutation {
	return scc.mutation
}

// Save creates the SystemConfig in the database.
func (scc *SystemConfigCreate) Save(ctx context.Context) (*SystemConfig, error) {
	scc.defaults()
	return withHooks(ctx, scc.sqlSave, scc.mutation, scc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (scc *SystemConfigCreate) SaveX(ctx context.Context) *SystemConfig {
	v, err := scc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scc *SystemConfigCreate) Exec(ctx context.Context) error {
	_, err := scc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scc *SystemConfigCreate) ExecX(ctx context.Context) {
	if err := scc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scc *SystemConfigCreate) defaults() {
	if _, ok := scc.mutation.Description(); !ok {
		v := systemconfig.DefaultDescription
		scc.mutation.SetDescription(v)
	}
	if _, ok := scc.mutation.Value(); !ok {
		v := systemconfig.DefaultValue
		scc.mutation.SetValue(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scc *SystemConfigCreate) check() error {
	if _, ok := scc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "SystemConfig.name"`)}
	}
	if _, ok := scc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "SystemConfig.description"`)}
	}
	if _, ok := scc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "SystemConfig.value"`)}
	}
	return nil
}

func (scc *SystemConfigCreate) sqlSave(ctx context.Context) (*SystemConfig, error) {
	if err := scc.check(); err != nil {
		return nil, err
	}
	_node, _spec := scc.createSpec()
	if err := sqlgraph.CreateNode(ctx, scc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected SystemConfig.ID type: %T", _spec.ID.Value)
		}
	}
	scc.mutation.id = &_node.ID
	scc.mutation.done = true
	return _node, nil
}

func (scc *SystemConfigCreate) createSpec() (*SystemConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &SystemConfig{config: scc.config}
		_spec = sqlgraph.NewCreateSpec(systemconfig.Table, sqlgraph.NewFieldSpec(systemconfig.FieldID, field.TypeString))
	)
	_spec.OnConflict = scc.conflict
	if id, ok := scc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := scc.mutation.Name(); ok {
		_spec.SetField(systemconfig.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := scc.mutation.Description(); ok {
		_spec.SetField(systemconfig.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := scc.mutation.Value(); ok {
		_spec.SetField(systemconfig.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SystemConfig.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SystemConfigUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (scc *SystemConfigCreate) OnConflict(opts ...sql.ConflictOption) *SystemConfigUpsertOne {
	scc.conflict = opts
	return &SystemConfigUpsertOne{
		create: scc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SystemConfig.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scc *SystemConfigCreate) OnConflictColumns(columns ...string) *SystemConfigUpsertOne {
	scc.conflict = append(scc.conflict, sql.ConflictColumns(columns...))
	return &SystemConfigUpsertOne{
		create: scc,
	}
}

type (
	// SystemConfigUpsertOne is the builder for "upsert"-ing
	//  one SystemConfig node.
	SystemConfigUpsertOne struct {
		create *SystemConfigCreate
	}

	// SystemConfigUpsert is the "OnConflict" setter.
	SystemConfigUpsert struct {
		*sql.UpdateSet
	}
)

// SetDescription sets the "description" field.
func (u *SystemConfigUpsert) SetDescription(v string) *SystemConfigUpsert {
	u.Set(systemconfig.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *SystemConfigUpsert) UpdateDescription() *SystemConfigUpsert {
	u.SetExcluded(systemconfig.FieldDescription)
	return u
}

// SetValue sets the "value" field.
func (u *SystemConfigUpsert) SetValue(v string) *SystemConfigUpsert {
	u.Set(systemconfig.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SystemConfigUpsert) UpdateValue() *SystemConfigUpsert {
	u.SetExcluded(systemconfig.FieldValue)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SystemConfig.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(systemconfig.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SystemConfigUpsertOne) UpdateNewValues() *SystemConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(systemconfig.FieldID)
		}
		if _, exists := u.create.mutation.Name(); exists {
			s.SetIgnore(systemconfig.FieldName)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SystemConfig.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SystemConfigUpsertOne) Ignore() *SystemConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SystemConfigUpsertOne) DoNothing() *SystemConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SystemConfigCreate.OnConflict
// documentation for more info.
func (u *SystemConfigUpsertOne) Update(set func(*SystemConfigUpsert)) *SystemConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SystemConfigUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *SystemConfigUpsertOne) SetDescription(v string) *SystemConfigUpsertOne {
	return u.Update(func(s *SystemConfigUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *SystemConfigUpsertOne) UpdateDescription() *SystemConfigUpsertOne {
	return u.Update(func(s *SystemConfigUpsert) {
		s.UpdateDescription()
	})
}

// SetValue sets the "value" field.
func (u *SystemConfigUpsertOne) SetValue(v string) *SystemConfigUpsertOne {
	return u.Update(func(s *SystemConfigUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SystemConfigUpsertOne) UpdateValue() *SystemConfigUpsertOne {
	return u.Update(func(s *SystemConfigUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *SystemConfigUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SystemConfigCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SystemConfigUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SystemConfigUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: SystemConfigUpsertOne.ID is not supported by MySQL driver. Use SystemConfigUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SystemConfigUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SystemConfigCreateBulk is the builder for creating many SystemConfig entities in bulk.
type SystemConfigCreateBulk struct {
	config
	err      error
	builders []*SystemConfigCreate
	conflict []sql.ConflictOption
}

// Save creates the SystemConfig entities in the database.
func (sccb *SystemConfigCreateBulk) Save(ctx context.Context) ([]*SystemConfig, error) {
	if sccb.err != nil {
		return nil, sccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sccb.builders))
	nodes := make([]*SystemConfig, len(sccb.builders))
	mutators := make([]Mutator, len(sccb.builders))
	for i := range sccb.builders {
		func(i int, root context.Context) {
			builder := sccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SystemConfigMutation)
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
					_, err = mutators[i+1].Mutate(root, sccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sccb *SystemConfigCreateBulk) SaveX(ctx context.Context) []*SystemConfig {
	v, err := sccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sccb *SystemConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := sccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sccb *SystemConfigCreateBulk) ExecX(ctx context.Context) {
	if err := sccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SystemConfig.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SystemConfigUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (sccb *SystemConfigCreateBulk) OnConflict(opts ...sql.ConflictOption) *SystemConfigUpsertBulk {
	sccb.conflict = opts
	return &SystemConfigUpsertBulk{
		create: sccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SystemConfig.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sccb *SystemConfigCreateBulk) OnConflictColumns(columns ...string) *SystemConfigUpsertBulk {
	sccb.conflict = append(sccb.conflict, sql.ConflictColumns(columns...))
	return &SystemConfigUpsertBulk{
		create: sccb,
	}
}

// SystemConfigUpsertBulk is the builder for "upsert"-ing
// a bulk of SystemConfig nodes.
type SystemConfigUpsertBulk struct {
	create *SystemConfigCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SystemConfig.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(systemconfig.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SystemConfigUpsertBulk) UpdateNewValues() *SystemConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(systemconfig.FieldID)
			}
			if _, exists := b.mutation.Name(); exists {
				s.SetIgnore(systemconfig.FieldName)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SystemConfig.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SystemConfigUpsertBulk) Ignore() *SystemConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SystemConfigUpsertBulk) DoNothing() *SystemConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SystemConfigCreateBulk.OnConflict
// documentation for more info.
func (u *SystemConfigUpsertBulk) Update(set func(*SystemConfigUpsert)) *SystemConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SystemConfigUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *SystemConfigUpsertBulk) SetDescription(v string) *SystemConfigUpsertBulk {
	return u.Update(func(s *SystemConfigUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *SystemConfigUpsertBulk) UpdateDescription() *SystemConfigUpsertBulk {
	return u.Update(func(s *SystemConfigUpsert) {
		s.UpdateDescription()
	})
}

// SetValue sets the "value" field.
func (u *SystemConfigUpsertBulk) SetValue(v string) *SystemConfigUpsertBulk {
	return u.Update(func(s *SystemConfigUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SystemConfigUpsertBulk) UpdateValue() *SystemConfigUpsertBulk {
	return u.Update(func(s *SystemConfigUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *SystemConfigUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SystemConfigCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SystemConfigCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SystemConfigUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
