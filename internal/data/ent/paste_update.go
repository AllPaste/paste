// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/AllPaste/paste/internal/data/ent/paste"
	"github.com/AllPaste/paste/internal/data/ent/predicate"
)

// PasteUpdate is the builder for updating Paste entities.
type PasteUpdate struct {
	config
	hooks    []Hook
	mutation *PasteMutation
}

// Where appends a list predicates to the PasteUpdate builder.
func (pu *PasteUpdate) Where(ps ...predicate.Paste) *PasteUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetContent sets the "content" field.
func (pu *PasteUpdate) SetContent(s string) *PasteUpdate {
	pu.mutation.SetContent(s)
	return pu
}

// SetCreator sets the "creator" field.
func (pu *PasteUpdate) SetCreator(s string) *PasteUpdate {
	pu.mutation.SetCreator(s)
	return pu
}

// Mutation returns the PasteMutation object of the builder.
func (pu *PasteUpdate) Mutation() *PasteMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PasteUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PasteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PasteUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PasteUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PasteUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PasteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   paste.Table,
			Columns: paste.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: paste.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Content(); ok {
		_spec.SetField(paste.FieldContent, field.TypeString, value)
	}
	if value, ok := pu.mutation.Creator(); ok {
		_spec.SetField(paste.FieldCreator, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paste.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PasteUpdateOne is the builder for updating a single Paste entity.
type PasteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PasteMutation
}

// SetContent sets the "content" field.
func (puo *PasteUpdateOne) SetContent(s string) *PasteUpdateOne {
	puo.mutation.SetContent(s)
	return puo
}

// SetCreator sets the "creator" field.
func (puo *PasteUpdateOne) SetCreator(s string) *PasteUpdateOne {
	puo.mutation.SetCreator(s)
	return puo
}

// Mutation returns the PasteMutation object of the builder.
func (puo *PasteUpdateOne) Mutation() *PasteMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PasteUpdateOne) Select(field string, fields ...string) *PasteUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Paste entity.
func (puo *PasteUpdateOne) Save(ctx context.Context) (*Paste, error) {
	var (
		err  error
		node *Paste
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PasteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Paste)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PasteMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PasteUpdateOne) SaveX(ctx context.Context) *Paste {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PasteUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PasteUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PasteUpdateOne) sqlSave(ctx context.Context) (_node *Paste, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   paste.Table,
			Columns: paste.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: paste.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Paste.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, paste.FieldID)
		for _, f := range fields {
			if !paste.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != paste.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Content(); ok {
		_spec.SetField(paste.FieldContent, field.TypeString, value)
	}
	if value, ok := puo.mutation.Creator(); ok {
		_spec.SetField(paste.FieldCreator, field.TypeString, value)
	}
	_node = &Paste{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paste.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}