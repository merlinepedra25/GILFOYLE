// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dreamvo/gilfoyle/ent/media"
	"github.com/dreamvo/gilfoyle/ent/mediafile"
	"github.com/dreamvo/gilfoyle/ent/probe"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// MediaCreate is the builder for creating a Media entity.
type MediaCreate struct {
	config
	mutation *MediaMutation
	hooks    []Hook
}

// SetTitle sets the title field.
func (mc *MediaCreate) SetTitle(s string) *MediaCreate {
	mc.mutation.SetTitle(s)
	return mc
}

// SetOriginalFilename sets the original_filename field.
func (mc *MediaCreate) SetOriginalFilename(s string) *MediaCreate {
	mc.mutation.SetOriginalFilename(s)
	return mc
}

// SetNillableOriginalFilename sets the original_filename field if the given value is not nil.
func (mc *MediaCreate) SetNillableOriginalFilename(s *string) *MediaCreate {
	if s != nil {
		mc.SetOriginalFilename(*s)
	}
	return mc
}

// SetStatus sets the status field.
func (mc *MediaCreate) SetStatus(m media.Status) *MediaCreate {
	mc.mutation.SetStatus(m)
	return mc
}

// SetMessage sets the message field.
func (mc *MediaCreate) SetMessage(s string) *MediaCreate {
	mc.mutation.SetMessage(s)
	return mc
}

// SetNillableMessage sets the message field if the given value is not nil.
func (mc *MediaCreate) SetNillableMessage(s *string) *MediaCreate {
	if s != nil {
		mc.SetMessage(*s)
	}
	return mc
}

// SetPlayable sets the playable field.
func (mc *MediaCreate) SetPlayable(b bool) *MediaCreate {
	mc.mutation.SetPlayable(b)
	return mc
}

// SetNillablePlayable sets the playable field if the given value is not nil.
func (mc *MediaCreate) SetNillablePlayable(b *bool) *MediaCreate {
	if b != nil {
		mc.SetPlayable(*b)
	}
	return mc
}

// SetCreatedAt sets the created_at field.
func (mc *MediaCreate) SetCreatedAt(t time.Time) *MediaCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (mc *MediaCreate) SetNillableCreatedAt(t *time.Time) *MediaCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the updated_at field.
func (mc *MediaCreate) SetUpdatedAt(t time.Time) *MediaCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (mc *MediaCreate) SetNillableUpdatedAt(t *time.Time) *MediaCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetID sets the id field.
func (mc *MediaCreate) SetID(u uuid.UUID) *MediaCreate {
	mc.mutation.SetID(u)
	return mc
}

// AddMediaFileIDs adds the media_files edge to MediaFile by ids.
func (mc *MediaCreate) AddMediaFileIDs(ids ...uuid.UUID) *MediaCreate {
	mc.mutation.AddMediaFileIDs(ids...)
	return mc
}

// AddMediaFiles adds the media_files edges to MediaFile.
func (mc *MediaCreate) AddMediaFiles(m ...*MediaFile) *MediaCreate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddMediaFileIDs(ids...)
}

// SetProbeID sets the probe edge to Probe by id.
func (mc *MediaCreate) SetProbeID(id uuid.UUID) *MediaCreate {
	mc.mutation.SetProbeID(id)
	return mc
}

// SetNillableProbeID sets the probe edge to Probe by id if the given value is not nil.
func (mc *MediaCreate) SetNillableProbeID(id *uuid.UUID) *MediaCreate {
	if id != nil {
		mc = mc.SetProbeID(*id)
	}
	return mc
}

// SetProbe sets the probe edge to Probe.
func (mc *MediaCreate) SetProbe(p *Probe) *MediaCreate {
	return mc.SetProbeID(p.ID)
}

// Mutation returns the MediaMutation object of the builder.
func (mc *MediaCreate) Mutation() *MediaMutation {
	return mc.mutation
}

// Save creates the Media in the database.
func (mc *MediaCreate) Save(ctx context.Context) (*Media, error) {
	var (
		err  error
		node *Media
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MediaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MediaCreate) SaveX(ctx context.Context) *Media {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (mc *MediaCreate) defaults() {
	if _, ok := mc.mutation.OriginalFilename(); !ok {
		v := media.DefaultOriginalFilename
		mc.mutation.SetOriginalFilename(v)
	}
	if _, ok := mc.mutation.Message(); !ok {
		v := media.DefaultMessage
		mc.mutation.SetMessage(v)
	}
	if _, ok := mc.mutation.Playable(); !ok {
		v := media.DefaultPlayable
		mc.mutation.SetPlayable(v)
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := media.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := media.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := media.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MediaCreate) check() error {
	if _, ok := mc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New("ent: missing required field \"title\"")}
	}
	if v, ok := mc.mutation.Title(); ok {
		if err := media.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := mc.mutation.OriginalFilename(); ok {
		if err := media.OriginalFilenameValidator(v); err != nil {
			return &ValidationError{Name: "original_filename", err: fmt.Errorf("ent: validator failed for field \"original_filename\": %w", err)}
		}
	}
	if _, ok := mc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if v, ok := mc.mutation.Status(); ok {
		if err := media.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if v, ok := mc.mutation.Message(); ok {
		if err := media.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf("ent: validator failed for field \"message\": %w", err)}
		}
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (mc *MediaCreate) sqlSave(ctx context.Context) (*Media, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (mc *MediaCreate) createSpec() (*Media, *sqlgraph.CreateSpec) {
	var (
		_node = &Media{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: media.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: media.FieldID,
			},
		}
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := mc.mutation.OriginalFilename(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldOriginalFilename,
		})
		_node.OriginalFilename = value
	}
	if value, ok := mc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: media.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := mc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := mc.mutation.Playable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: media.FieldPlayable,
		})
		_node.Playable = value
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: media.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: media.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := mc.mutation.MediaFilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   media.MediaFilesTable,
			Columns: []string{media.MediaFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mediafile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ProbeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   media.ProbeTable,
			Columns: []string{media.ProbeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: probe.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MediaCreateBulk is the builder for creating a bulk of Media entities.
type MediaCreateBulk struct {
	config
	builders []*MediaCreate
}

// Save creates the Media entities in the database.
func (mcb *MediaCreateBulk) Save(ctx context.Context) ([]*Media, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Media, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MediaMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (mcb *MediaCreateBulk) SaveX(ctx context.Context) []*Media {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
