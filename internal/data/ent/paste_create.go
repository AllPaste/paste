// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/AllPaste/paste/internal/data/ent/paste"
)

// PasteCreate is the builder for creating a Paste entity.
type PasteCreate struct {
	config
	mutation *PasteMutation
	hooks    []Hook
}

// SetContent sets the "content" field.
func (pc *PasteCreate) SetContent(s string) *PasteCreate {
	pc.mutation.SetContent(s)
	return pc
}

// SetCreator sets the "creator" field.
func (pc *PasteCreate) SetCreator(s string) *PasteCreate {
	pc.mutation.SetCreator(s)
	return pc
}

// Mutation returns the PasteMutation object of the builder.
func (pc *PasteCreate) Mutation() *PasteMutation {
	return pc.mutation
}

// Save creates the Paste in the database.
func (pc *PasteCreate) Save(ctx context.Context) (*Paste, error) {
	var (
		err  error
		node *Paste
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PasteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (pc *PasteCreate) SaveX(ctx context.Context) *Paste {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PasteCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PasteCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PasteCreate) check() error {
	if _, ok := pc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Paste.content"`)}
	}
	if _, ok := pc.mutation.Creator(); !ok {
		return &ValidationError{Name: "creator", err: errors.New(`ent: missing required field "Paste.creator"`)}
	}
	return nil
}

func (pc *PasteCreate) sqlSave(ctx context.Context) (*Paste, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PasteCreate) createSpec() (*Paste, *sqlgraph.CreateSpec) {
	var (
		_node = &Paste{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: paste.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: paste.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Content(); ok {
		_spec.SetField(paste.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := pc.mutation.Creator(); ok {
		_spec.SetField(paste.FieldCreator, field.TypeString, value)
		_node.Creator = value
	}
	return _node, _spec
}

// PasteCreateBulk is the builder for creating many Paste entities in bulk.
type PasteCreateBulk struct {
	config
	builders []*PasteCreate
}

// Save creates the Paste entities in the database.
func (pcb *PasteCreateBulk) Save(ctx context.Context) ([]*Paste, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Paste, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PasteMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PasteCreateBulk) SaveX(ctx context.Context) []*Paste {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PasteCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PasteCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
