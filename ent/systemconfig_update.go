// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wolodata/schema/ent/predicate"
	"github.com/wolodata/schema/ent/systemconfig"
)

// SystemConfigUpdate is the builder for updating SystemConfig entities.
type SystemConfigUpdate struct {
	config
	hooks    []Hook
	mutation *SystemConfigMutation
}

// Where appends a list predicates to the SystemConfigUpdate builder.
func (scu *SystemConfigUpdate) Where(ps ...predicate.SystemConfig) *SystemConfigUpdate {
	scu.mutation.Where(ps...)
	return scu
}

// SetDescription sets the "description" field.
func (scu *SystemConfigUpdate) SetDescription(s string) *SystemConfigUpdate {
	scu.mutation.SetDescription(s)
	return scu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (scu *SystemConfigUpdate) SetNillableDescription(s *string) *SystemConfigUpdate {
	if s != nil {
		scu.SetDescription(*s)
	}
	return scu
}

// SetAPIURL sets the "api_url" field.
func (scu *SystemConfigUpdate) SetAPIURL(s string) *SystemConfigUpdate {
	scu.mutation.SetAPIURL(s)
	return scu
}

// SetNillableAPIURL sets the "api_url" field if the given value is not nil.
func (scu *SystemConfigUpdate) SetNillableAPIURL(s *string) *SystemConfigUpdate {
	if s != nil {
		scu.SetAPIURL(*s)
	}
	return scu
}

// SetAPIKey sets the "api_key" field.
func (scu *SystemConfigUpdate) SetAPIKey(s string) *SystemConfigUpdate {
	scu.mutation.SetAPIKey(s)
	return scu
}

// SetNillableAPIKey sets the "api_key" field if the given value is not nil.
func (scu *SystemConfigUpdate) SetNillableAPIKey(s *string) *SystemConfigUpdate {
	if s != nil {
		scu.SetAPIKey(*s)
	}
	return scu
}

// SetPromptSystem sets the "prompt_system" field.
func (scu *SystemConfigUpdate) SetPromptSystem(s string) *SystemConfigUpdate {
	scu.mutation.SetPromptSystem(s)
	return scu
}

// SetNillablePromptSystem sets the "prompt_system" field if the given value is not nil.
func (scu *SystemConfigUpdate) SetNillablePromptSystem(s *string) *SystemConfigUpdate {
	if s != nil {
		scu.SetPromptSystem(*s)
	}
	return scu
}

// SetPromptUser sets the "prompt_user" field.
func (scu *SystemConfigUpdate) SetPromptUser(s string) *SystemConfigUpdate {
	scu.mutation.SetPromptUser(s)
	return scu
}

// SetNillablePromptUser sets the "prompt_user" field if the given value is not nil.
func (scu *SystemConfigUpdate) SetNillablePromptUser(s *string) *SystemConfigUpdate {
	if s != nil {
		scu.SetPromptUser(*s)
	}
	return scu
}

// Mutation returns the SystemConfigMutation object of the builder.
func (scu *SystemConfigUpdate) Mutation() *SystemConfigMutation {
	return scu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (scu *SystemConfigUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, scu.sqlSave, scu.mutation, scu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scu *SystemConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := scu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (scu *SystemConfigUpdate) Exec(ctx context.Context) error {
	_, err := scu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scu *SystemConfigUpdate) ExecX(ctx context.Context) {
	if err := scu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (scu *SystemConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(systemconfig.Table, systemconfig.Columns, sqlgraph.NewFieldSpec(systemconfig.FieldID, field.TypeString))
	if ps := scu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scu.mutation.Description(); ok {
		_spec.SetField(systemconfig.FieldDescription, field.TypeString, value)
	}
	if value, ok := scu.mutation.APIURL(); ok {
		_spec.SetField(systemconfig.FieldAPIURL, field.TypeString, value)
	}
	if value, ok := scu.mutation.APIKey(); ok {
		_spec.SetField(systemconfig.FieldAPIKey, field.TypeString, value)
	}
	if value, ok := scu.mutation.PromptSystem(); ok {
		_spec.SetField(systemconfig.FieldPromptSystem, field.TypeString, value)
	}
	if value, ok := scu.mutation.PromptUser(); ok {
		_spec.SetField(systemconfig.FieldPromptUser, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, scu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{systemconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	scu.mutation.done = true
	return n, nil
}

// SystemConfigUpdateOne is the builder for updating a single SystemConfig entity.
type SystemConfigUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SystemConfigMutation
}

// SetDescription sets the "description" field.
func (scuo *SystemConfigUpdateOne) SetDescription(s string) *SystemConfigUpdateOne {
	scuo.mutation.SetDescription(s)
	return scuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (scuo *SystemConfigUpdateOne) SetNillableDescription(s *string) *SystemConfigUpdateOne {
	if s != nil {
		scuo.SetDescription(*s)
	}
	return scuo
}

// SetAPIURL sets the "api_url" field.
func (scuo *SystemConfigUpdateOne) SetAPIURL(s string) *SystemConfigUpdateOne {
	scuo.mutation.SetAPIURL(s)
	return scuo
}

// SetNillableAPIURL sets the "api_url" field if the given value is not nil.
func (scuo *SystemConfigUpdateOne) SetNillableAPIURL(s *string) *SystemConfigUpdateOne {
	if s != nil {
		scuo.SetAPIURL(*s)
	}
	return scuo
}

// SetAPIKey sets the "api_key" field.
func (scuo *SystemConfigUpdateOne) SetAPIKey(s string) *SystemConfigUpdateOne {
	scuo.mutation.SetAPIKey(s)
	return scuo
}

// SetNillableAPIKey sets the "api_key" field if the given value is not nil.
func (scuo *SystemConfigUpdateOne) SetNillableAPIKey(s *string) *SystemConfigUpdateOne {
	if s != nil {
		scuo.SetAPIKey(*s)
	}
	return scuo
}

// SetPromptSystem sets the "prompt_system" field.
func (scuo *SystemConfigUpdateOne) SetPromptSystem(s string) *SystemConfigUpdateOne {
	scuo.mutation.SetPromptSystem(s)
	return scuo
}

// SetNillablePromptSystem sets the "prompt_system" field if the given value is not nil.
func (scuo *SystemConfigUpdateOne) SetNillablePromptSystem(s *string) *SystemConfigUpdateOne {
	if s != nil {
		scuo.SetPromptSystem(*s)
	}
	return scuo
}

// SetPromptUser sets the "prompt_user" field.
func (scuo *SystemConfigUpdateOne) SetPromptUser(s string) *SystemConfigUpdateOne {
	scuo.mutation.SetPromptUser(s)
	return scuo
}

// SetNillablePromptUser sets the "prompt_user" field if the given value is not nil.
func (scuo *SystemConfigUpdateOne) SetNillablePromptUser(s *string) *SystemConfigUpdateOne {
	if s != nil {
		scuo.SetPromptUser(*s)
	}
	return scuo
}

// Mutation returns the SystemConfigMutation object of the builder.
func (scuo *SystemConfigUpdateOne) Mutation() *SystemConfigMutation {
	return scuo.mutation
}

// Where appends a list predicates to the SystemConfigUpdate builder.
func (scuo *SystemConfigUpdateOne) Where(ps ...predicate.SystemConfig) *SystemConfigUpdateOne {
	scuo.mutation.Where(ps...)
	return scuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (scuo *SystemConfigUpdateOne) Select(field string, fields ...string) *SystemConfigUpdateOne {
	scuo.fields = append([]string{field}, fields...)
	return scuo
}

// Save executes the query and returns the updated SystemConfig entity.
func (scuo *SystemConfigUpdateOne) Save(ctx context.Context) (*SystemConfig, error) {
	return withHooks(ctx, scuo.sqlSave, scuo.mutation, scuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scuo *SystemConfigUpdateOne) SaveX(ctx context.Context) *SystemConfig {
	node, err := scuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (scuo *SystemConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := scuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scuo *SystemConfigUpdateOne) ExecX(ctx context.Context) {
	if err := scuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (scuo *SystemConfigUpdateOne) sqlSave(ctx context.Context) (_node *SystemConfig, err error) {
	_spec := sqlgraph.NewUpdateSpec(systemconfig.Table, systemconfig.Columns, sqlgraph.NewFieldSpec(systemconfig.FieldID, field.TypeString))
	id, ok := scuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SystemConfig.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := scuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, systemconfig.FieldID)
		for _, f := range fields {
			if !systemconfig.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != systemconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := scuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scuo.mutation.Description(); ok {
		_spec.SetField(systemconfig.FieldDescription, field.TypeString, value)
	}
	if value, ok := scuo.mutation.APIURL(); ok {
		_spec.SetField(systemconfig.FieldAPIURL, field.TypeString, value)
	}
	if value, ok := scuo.mutation.APIKey(); ok {
		_spec.SetField(systemconfig.FieldAPIKey, field.TypeString, value)
	}
	if value, ok := scuo.mutation.PromptSystem(); ok {
		_spec.SetField(systemconfig.FieldPromptSystem, field.TypeString, value)
	}
	if value, ok := scuo.mutation.PromptUser(); ok {
		_spec.SetField(systemconfig.FieldPromptUser, field.TypeString, value)
	}
	_node = &SystemConfig{config: scuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, scuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{systemconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	scuo.mutation.done = true
	return _node, nil
}
