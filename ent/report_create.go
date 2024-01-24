// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wolodata/schema/ent/report"
)

// ReportCreate is the builder for creating a Report entity.
type ReportCreate struct {
	config
	mutation *ReportMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetReportType sets the "report_type" field.
func (rc *ReportCreate) SetReportType(s string) *ReportCreate {
	rc.mutation.SetReportType(s)
	return rc
}

// SetTriggerUserID sets the "trigger_user_id" field.
func (rc *ReportCreate) SetTriggerUserID(i int32) *ReportCreate {
	rc.mutation.SetTriggerUserID(i)
	return rc
}

// SetNillableTriggerUserID sets the "trigger_user_id" field if the given value is not nil.
func (rc *ReportCreate) SetNillableTriggerUserID(i *int32) *ReportCreate {
	if i != nil {
		rc.SetTriggerUserID(*i)
	}
	return rc
}

// SetTriggerAt sets the "trigger_at" field.
func (rc *ReportCreate) SetTriggerAt(t time.Time) *ReportCreate {
	rc.mutation.SetTriggerAt(t)
	return rc
}

// SetNillableTriggerAt sets the "trigger_at" field if the given value is not nil.
func (rc *ReportCreate) SetNillableTriggerAt(t *time.Time) *ReportCreate {
	if t != nil {
		rc.SetTriggerAt(*t)
	}
	return rc
}

// SetRelatedArticleIds sets the "related_article_ids" field.
func (rc *ReportCreate) SetRelatedArticleIds(s []string) *ReportCreate {
	rc.mutation.SetRelatedArticleIds(s)
	return rc
}

// SetContent sets the "content" field.
func (rc *ReportCreate) SetContent(s string) *ReportCreate {
	rc.mutation.SetContent(s)
	return rc
}

// SetReason sets the "reason" field.
func (rc *ReportCreate) SetReason(s string) *ReportCreate {
	rc.mutation.SetReason(s)
	return rc
}

// SetGeneratedAt sets the "generated_at" field.
func (rc *ReportCreate) SetGeneratedAt(t time.Time) *ReportCreate {
	rc.mutation.SetGeneratedAt(t)
	return rc
}

// SetID sets the "id" field.
func (rc *ReportCreate) SetID(i int32) *ReportCreate {
	rc.mutation.SetID(i)
	return rc
}

// Mutation returns the ReportMutation object of the builder.
func (rc *ReportCreate) Mutation() *ReportMutation {
	return rc.mutation
}

// Save creates the Report in the database.
func (rc *ReportCreate) Save(ctx context.Context) (*Report, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReportCreate) SaveX(ctx context.Context) *Report {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReportCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReportCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ReportCreate) defaults() {
	if _, ok := rc.mutation.TriggerAt(); !ok {
		v := report.DefaultTriggerAt()
		rc.mutation.SetTriggerAt(v)
	}
	if _, ok := rc.mutation.RelatedArticleIds(); !ok {
		v := report.DefaultRelatedArticleIds
		rc.mutation.SetRelatedArticleIds(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReportCreate) check() error {
	if _, ok := rc.mutation.ReportType(); !ok {
		return &ValidationError{Name: "report_type", err: errors.New(`ent: missing required field "Report.report_type"`)}
	}
	if v, ok := rc.mutation.ReportType(); ok {
		if err := report.ReportTypeValidator(v); err != nil {
			return &ValidationError{Name: "report_type", err: fmt.Errorf(`ent: validator failed for field "Report.report_type": %w`, err)}
		}
	}
	if _, ok := rc.mutation.TriggerAt(); !ok {
		return &ValidationError{Name: "trigger_at", err: errors.New(`ent: missing required field "Report.trigger_at"`)}
	}
	if _, ok := rc.mutation.RelatedArticleIds(); !ok {
		return &ValidationError{Name: "related_article_ids", err: errors.New(`ent: missing required field "Report.related_article_ids"`)}
	}
	if _, ok := rc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Report.content"`)}
	}
	if _, ok := rc.mutation.Reason(); !ok {
		return &ValidationError{Name: "reason", err: errors.New(`ent: missing required field "Report.reason"`)}
	}
	if _, ok := rc.mutation.GeneratedAt(); !ok {
		return &ValidationError{Name: "generated_at", err: errors.New(`ent: missing required field "Report.generated_at"`)}
	}
	return nil
}

func (rc *ReportCreate) sqlSave(ctx context.Context) (*Report, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ReportCreate) createSpec() (*Report, *sqlgraph.CreateSpec) {
	var (
		_node = &Report{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(report.Table, sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt32))
	)
	_spec.OnConflict = rc.conflict
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.ReportType(); ok {
		_spec.SetField(report.FieldReportType, field.TypeString, value)
		_node.ReportType = value
	}
	if value, ok := rc.mutation.TriggerUserID(); ok {
		_spec.SetField(report.FieldTriggerUserID, field.TypeInt32, value)
		_node.TriggerUserID = value
	}
	if value, ok := rc.mutation.TriggerAt(); ok {
		_spec.SetField(report.FieldTriggerAt, field.TypeTime, value)
		_node.TriggerAt = value
	}
	if value, ok := rc.mutation.RelatedArticleIds(); ok {
		_spec.SetField(report.FieldRelatedArticleIds, field.TypeJSON, value)
		_node.RelatedArticleIds = value
	}
	if value, ok := rc.mutation.Content(); ok {
		_spec.SetField(report.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := rc.mutation.Reason(); ok {
		_spec.SetField(report.FieldReason, field.TypeString, value)
		_node.Reason = value
	}
	if value, ok := rc.mutation.GeneratedAt(); ok {
		_spec.SetField(report.FieldGeneratedAt, field.TypeTime, value)
		_node.GeneratedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Report.Create().
//		SetReportType(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReportUpsert) {
//			SetReportType(v+v).
//		}).
//		Exec(ctx)
func (rc *ReportCreate) OnConflict(opts ...sql.ConflictOption) *ReportUpsertOne {
	rc.conflict = opts
	return &ReportUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Report.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *ReportCreate) OnConflictColumns(columns ...string) *ReportUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &ReportUpsertOne{
		create: rc,
	}
}

type (
	// ReportUpsertOne is the builder for "upsert"-ing
	//  one Report node.
	ReportUpsertOne struct {
		create *ReportCreate
	}

	// ReportUpsert is the "OnConflict" setter.
	ReportUpsert struct {
		*sql.UpdateSet
	}
)

// SetContent sets the "content" field.
func (u *ReportUpsert) SetContent(v string) *ReportUpsert {
	u.Set(report.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *ReportUpsert) UpdateContent() *ReportUpsert {
	u.SetExcluded(report.FieldContent)
	return u
}

// SetReason sets the "reason" field.
func (u *ReportUpsert) SetReason(v string) *ReportUpsert {
	u.Set(report.FieldReason, v)
	return u
}

// UpdateReason sets the "reason" field to the value that was provided on create.
func (u *ReportUpsert) UpdateReason() *ReportUpsert {
	u.SetExcluded(report.FieldReason)
	return u
}

// SetGeneratedAt sets the "generated_at" field.
func (u *ReportUpsert) SetGeneratedAt(v time.Time) *ReportUpsert {
	u.Set(report.FieldGeneratedAt, v)
	return u
}

// UpdateGeneratedAt sets the "generated_at" field to the value that was provided on create.
func (u *ReportUpsert) UpdateGeneratedAt() *ReportUpsert {
	u.SetExcluded(report.FieldGeneratedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Report.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(report.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ReportUpsertOne) UpdateNewValues() *ReportUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(report.FieldID)
		}
		if _, exists := u.create.mutation.ReportType(); exists {
			s.SetIgnore(report.FieldReportType)
		}
		if _, exists := u.create.mutation.TriggerUserID(); exists {
			s.SetIgnore(report.FieldTriggerUserID)
		}
		if _, exists := u.create.mutation.TriggerAt(); exists {
			s.SetIgnore(report.FieldTriggerAt)
		}
		if _, exists := u.create.mutation.RelatedArticleIds(); exists {
			s.SetIgnore(report.FieldRelatedArticleIds)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Report.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ReportUpsertOne) Ignore() *ReportUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReportUpsertOne) DoNothing() *ReportUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReportCreate.OnConflict
// documentation for more info.
func (u *ReportUpsertOne) Update(set func(*ReportUpsert)) *ReportUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReportUpsert{UpdateSet: update})
	}))
	return u
}

// SetContent sets the "content" field.
func (u *ReportUpsertOne) SetContent(v string) *ReportUpsertOne {
	return u.Update(func(s *ReportUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *ReportUpsertOne) UpdateContent() *ReportUpsertOne {
	return u.Update(func(s *ReportUpsert) {
		s.UpdateContent()
	})
}

// SetReason sets the "reason" field.
func (u *ReportUpsertOne) SetReason(v string) *ReportUpsertOne {
	return u.Update(func(s *ReportUpsert) {
		s.SetReason(v)
	})
}

// UpdateReason sets the "reason" field to the value that was provided on create.
func (u *ReportUpsertOne) UpdateReason() *ReportUpsertOne {
	return u.Update(func(s *ReportUpsert) {
		s.UpdateReason()
	})
}

// SetGeneratedAt sets the "generated_at" field.
func (u *ReportUpsertOne) SetGeneratedAt(v time.Time) *ReportUpsertOne {
	return u.Update(func(s *ReportUpsert) {
		s.SetGeneratedAt(v)
	})
}

// UpdateGeneratedAt sets the "generated_at" field to the value that was provided on create.
func (u *ReportUpsertOne) UpdateGeneratedAt() *ReportUpsertOne {
	return u.Update(func(s *ReportUpsert) {
		s.UpdateGeneratedAt()
	})
}

// Exec executes the query.
func (u *ReportUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ReportCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReportUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ReportUpsertOne) ID(ctx context.Context) (id int32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ReportUpsertOne) IDX(ctx context.Context) int32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ReportCreateBulk is the builder for creating many Report entities in bulk.
type ReportCreateBulk struct {
	config
	err      error
	builders []*ReportCreate
	conflict []sql.ConflictOption
}

// Save creates the Report entities in the database.
func (rcb *ReportCreateBulk) Save(ctx context.Context) ([]*Report, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Report, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReportMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReportCreateBulk) SaveX(ctx context.Context) []*Report {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReportCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReportCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Report.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReportUpsert) {
//			SetReportType(v+v).
//		}).
//		Exec(ctx)
func (rcb *ReportCreateBulk) OnConflict(opts ...sql.ConflictOption) *ReportUpsertBulk {
	rcb.conflict = opts
	return &ReportUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Report.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *ReportCreateBulk) OnConflictColumns(columns ...string) *ReportUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &ReportUpsertBulk{
		create: rcb,
	}
}

// ReportUpsertBulk is the builder for "upsert"-ing
// a bulk of Report nodes.
type ReportUpsertBulk struct {
	create *ReportCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Report.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(report.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ReportUpsertBulk) UpdateNewValues() *ReportUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(report.FieldID)
			}
			if _, exists := b.mutation.ReportType(); exists {
				s.SetIgnore(report.FieldReportType)
			}
			if _, exists := b.mutation.TriggerUserID(); exists {
				s.SetIgnore(report.FieldTriggerUserID)
			}
			if _, exists := b.mutation.TriggerAt(); exists {
				s.SetIgnore(report.FieldTriggerAt)
			}
			if _, exists := b.mutation.RelatedArticleIds(); exists {
				s.SetIgnore(report.FieldRelatedArticleIds)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Report.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ReportUpsertBulk) Ignore() *ReportUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReportUpsertBulk) DoNothing() *ReportUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReportCreateBulk.OnConflict
// documentation for more info.
func (u *ReportUpsertBulk) Update(set func(*ReportUpsert)) *ReportUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReportUpsert{UpdateSet: update})
	}))
	return u
}

// SetContent sets the "content" field.
func (u *ReportUpsertBulk) SetContent(v string) *ReportUpsertBulk {
	return u.Update(func(s *ReportUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *ReportUpsertBulk) UpdateContent() *ReportUpsertBulk {
	return u.Update(func(s *ReportUpsert) {
		s.UpdateContent()
	})
}

// SetReason sets the "reason" field.
func (u *ReportUpsertBulk) SetReason(v string) *ReportUpsertBulk {
	return u.Update(func(s *ReportUpsert) {
		s.SetReason(v)
	})
}

// UpdateReason sets the "reason" field to the value that was provided on create.
func (u *ReportUpsertBulk) UpdateReason() *ReportUpsertBulk {
	return u.Update(func(s *ReportUpsert) {
		s.UpdateReason()
	})
}

// SetGeneratedAt sets the "generated_at" field.
func (u *ReportUpsertBulk) SetGeneratedAt(v time.Time) *ReportUpsertBulk {
	return u.Update(func(s *ReportUpsert) {
		s.SetGeneratedAt(v)
	})
}

// UpdateGeneratedAt sets the "generated_at" field to the value that was provided on create.
func (u *ReportUpsertBulk) UpdateGeneratedAt() *ReportUpsertBulk {
	return u.Update(func(s *ReportUpsert) {
		s.UpdateGeneratedAt()
	})
}

// Exec executes the query.
func (u *ReportUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ReportCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ReportCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReportUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
