// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/chapter"
	"github.com/zibbp/ganymede/ent/predicate"
	"github.com/zibbp/ganymede/ent/vod"
)

// ChapterUpdate is the builder for updating Chapter entities.
type ChapterUpdate struct {
	config
	hooks    []Hook
	mutation *ChapterMutation
}

// Where appends a list predicates to the ChapterUpdate builder.
func (cu *ChapterUpdate) Where(ps ...predicate.Chapter) *ChapterUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetType sets the "type" field.
func (cu *ChapterUpdate) SetType(s string) *ChapterUpdate {
	cu.mutation.SetType(s)
	return cu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cu *ChapterUpdate) SetNillableType(s *string) *ChapterUpdate {
	if s != nil {
		cu.SetType(*s)
	}
	return cu
}

// ClearType clears the value of the "type" field.
func (cu *ChapterUpdate) ClearType() *ChapterUpdate {
	cu.mutation.ClearType()
	return cu
}

// SetTitle sets the "title" field.
func (cu *ChapterUpdate) SetTitle(s string) *ChapterUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cu *ChapterUpdate) SetNillableTitle(s *string) *ChapterUpdate {
	if s != nil {
		cu.SetTitle(*s)
	}
	return cu
}

// ClearTitle clears the value of the "title" field.
func (cu *ChapterUpdate) ClearTitle() *ChapterUpdate {
	cu.mutation.ClearTitle()
	return cu
}

// SetStart sets the "start" field.
func (cu *ChapterUpdate) SetStart(i int) *ChapterUpdate {
	cu.mutation.ResetStart()
	cu.mutation.SetStart(i)
	return cu
}

// SetNillableStart sets the "start" field if the given value is not nil.
func (cu *ChapterUpdate) SetNillableStart(i *int) *ChapterUpdate {
	if i != nil {
		cu.SetStart(*i)
	}
	return cu
}

// AddStart adds i to the "start" field.
func (cu *ChapterUpdate) AddStart(i int) *ChapterUpdate {
	cu.mutation.AddStart(i)
	return cu
}

// ClearStart clears the value of the "start" field.
func (cu *ChapterUpdate) ClearStart() *ChapterUpdate {
	cu.mutation.ClearStart()
	return cu
}

// SetEnd sets the "end" field.
func (cu *ChapterUpdate) SetEnd(i int) *ChapterUpdate {
	cu.mutation.ResetEnd()
	cu.mutation.SetEnd(i)
	return cu
}

// SetNillableEnd sets the "end" field if the given value is not nil.
func (cu *ChapterUpdate) SetNillableEnd(i *int) *ChapterUpdate {
	if i != nil {
		cu.SetEnd(*i)
	}
	return cu
}

// AddEnd adds i to the "end" field.
func (cu *ChapterUpdate) AddEnd(i int) *ChapterUpdate {
	cu.mutation.AddEnd(i)
	return cu
}

// ClearEnd clears the value of the "end" field.
func (cu *ChapterUpdate) ClearEnd() *ChapterUpdate {
	cu.mutation.ClearEnd()
	return cu
}

// SetVodID sets the "vod" edge to the Vod entity by ID.
func (cu *ChapterUpdate) SetVodID(id uuid.UUID) *ChapterUpdate {
	cu.mutation.SetVodID(id)
	return cu
}

// SetVod sets the "vod" edge to the Vod entity.
func (cu *ChapterUpdate) SetVod(v *Vod) *ChapterUpdate {
	return cu.SetVodID(v.ID)
}

// Mutation returns the ChapterMutation object of the builder.
func (cu *ChapterUpdate) Mutation() *ChapterMutation {
	return cu.mutation
}

// ClearVod clears the "vod" edge to the Vod entity.
func (cu *ChapterUpdate) ClearVod() *ChapterUpdate {
	cu.mutation.ClearVod()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChapterUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChapterUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChapterUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChapterUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ChapterUpdate) check() error {
	if _, ok := cu.mutation.VodID(); cu.mutation.VodCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Chapter.vod"`)
	}
	return nil
}

func (cu *ChapterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(chapter.Table, chapter.Columns, sqlgraph.NewFieldSpec(chapter.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.SetField(chapter.FieldType, field.TypeString, value)
	}
	if cu.mutation.TypeCleared() {
		_spec.ClearField(chapter.FieldType, field.TypeString)
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.SetField(chapter.FieldTitle, field.TypeString, value)
	}
	if cu.mutation.TitleCleared() {
		_spec.ClearField(chapter.FieldTitle, field.TypeString)
	}
	if value, ok := cu.mutation.Start(); ok {
		_spec.SetField(chapter.FieldStart, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedStart(); ok {
		_spec.AddField(chapter.FieldStart, field.TypeInt, value)
	}
	if cu.mutation.StartCleared() {
		_spec.ClearField(chapter.FieldStart, field.TypeInt)
	}
	if value, ok := cu.mutation.End(); ok {
		_spec.SetField(chapter.FieldEnd, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedEnd(); ok {
		_spec.AddField(chapter.FieldEnd, field.TypeInt, value)
	}
	if cu.mutation.EndCleared() {
		_spec.ClearField(chapter.FieldEnd, field.TypeInt)
	}
	if cu.mutation.VodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chapter.VodTable,
			Columns: []string{chapter.VodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vod.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.VodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chapter.VodTable,
			Columns: []string{chapter.VodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vod.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chapter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ChapterUpdateOne is the builder for updating a single Chapter entity.
type ChapterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChapterMutation
}

// SetType sets the "type" field.
func (cuo *ChapterUpdateOne) SetType(s string) *ChapterUpdateOne {
	cuo.mutation.SetType(s)
	return cuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cuo *ChapterUpdateOne) SetNillableType(s *string) *ChapterUpdateOne {
	if s != nil {
		cuo.SetType(*s)
	}
	return cuo
}

// ClearType clears the value of the "type" field.
func (cuo *ChapterUpdateOne) ClearType() *ChapterUpdateOne {
	cuo.mutation.ClearType()
	return cuo
}

// SetTitle sets the "title" field.
func (cuo *ChapterUpdateOne) SetTitle(s string) *ChapterUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cuo *ChapterUpdateOne) SetNillableTitle(s *string) *ChapterUpdateOne {
	if s != nil {
		cuo.SetTitle(*s)
	}
	return cuo
}

// ClearTitle clears the value of the "title" field.
func (cuo *ChapterUpdateOne) ClearTitle() *ChapterUpdateOne {
	cuo.mutation.ClearTitle()
	return cuo
}

// SetStart sets the "start" field.
func (cuo *ChapterUpdateOne) SetStart(i int) *ChapterUpdateOne {
	cuo.mutation.ResetStart()
	cuo.mutation.SetStart(i)
	return cuo
}

// SetNillableStart sets the "start" field if the given value is not nil.
func (cuo *ChapterUpdateOne) SetNillableStart(i *int) *ChapterUpdateOne {
	if i != nil {
		cuo.SetStart(*i)
	}
	return cuo
}

// AddStart adds i to the "start" field.
func (cuo *ChapterUpdateOne) AddStart(i int) *ChapterUpdateOne {
	cuo.mutation.AddStart(i)
	return cuo
}

// ClearStart clears the value of the "start" field.
func (cuo *ChapterUpdateOne) ClearStart() *ChapterUpdateOne {
	cuo.mutation.ClearStart()
	return cuo
}

// SetEnd sets the "end" field.
func (cuo *ChapterUpdateOne) SetEnd(i int) *ChapterUpdateOne {
	cuo.mutation.ResetEnd()
	cuo.mutation.SetEnd(i)
	return cuo
}

// SetNillableEnd sets the "end" field if the given value is not nil.
func (cuo *ChapterUpdateOne) SetNillableEnd(i *int) *ChapterUpdateOne {
	if i != nil {
		cuo.SetEnd(*i)
	}
	return cuo
}

// AddEnd adds i to the "end" field.
func (cuo *ChapterUpdateOne) AddEnd(i int) *ChapterUpdateOne {
	cuo.mutation.AddEnd(i)
	return cuo
}

// ClearEnd clears the value of the "end" field.
func (cuo *ChapterUpdateOne) ClearEnd() *ChapterUpdateOne {
	cuo.mutation.ClearEnd()
	return cuo
}

// SetVodID sets the "vod" edge to the Vod entity by ID.
func (cuo *ChapterUpdateOne) SetVodID(id uuid.UUID) *ChapterUpdateOne {
	cuo.mutation.SetVodID(id)
	return cuo
}

// SetVod sets the "vod" edge to the Vod entity.
func (cuo *ChapterUpdateOne) SetVod(v *Vod) *ChapterUpdateOne {
	return cuo.SetVodID(v.ID)
}

// Mutation returns the ChapterMutation object of the builder.
func (cuo *ChapterUpdateOne) Mutation() *ChapterMutation {
	return cuo.mutation
}

// ClearVod clears the "vod" edge to the Vod entity.
func (cuo *ChapterUpdateOne) ClearVod() *ChapterUpdateOne {
	cuo.mutation.ClearVod()
	return cuo
}

// Where appends a list predicates to the ChapterUpdate builder.
func (cuo *ChapterUpdateOne) Where(ps ...predicate.Chapter) *ChapterUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChapterUpdateOne) Select(field string, fields ...string) *ChapterUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Chapter entity.
func (cuo *ChapterUpdateOne) Save(ctx context.Context) (*Chapter, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChapterUpdateOne) SaveX(ctx context.Context) *Chapter {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChapterUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChapterUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ChapterUpdateOne) check() error {
	if _, ok := cuo.mutation.VodID(); cuo.mutation.VodCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Chapter.vod"`)
	}
	return nil
}

func (cuo *ChapterUpdateOne) sqlSave(ctx context.Context) (_node *Chapter, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(chapter.Table, chapter.Columns, sqlgraph.NewFieldSpec(chapter.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Chapter.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chapter.FieldID)
		for _, f := range fields {
			if !chapter.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chapter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.SetField(chapter.FieldType, field.TypeString, value)
	}
	if cuo.mutation.TypeCleared() {
		_spec.ClearField(chapter.FieldType, field.TypeString)
	}
	if value, ok := cuo.mutation.Title(); ok {
		_spec.SetField(chapter.FieldTitle, field.TypeString, value)
	}
	if cuo.mutation.TitleCleared() {
		_spec.ClearField(chapter.FieldTitle, field.TypeString)
	}
	if value, ok := cuo.mutation.Start(); ok {
		_spec.SetField(chapter.FieldStart, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedStart(); ok {
		_spec.AddField(chapter.FieldStart, field.TypeInt, value)
	}
	if cuo.mutation.StartCleared() {
		_spec.ClearField(chapter.FieldStart, field.TypeInt)
	}
	if value, ok := cuo.mutation.End(); ok {
		_spec.SetField(chapter.FieldEnd, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedEnd(); ok {
		_spec.AddField(chapter.FieldEnd, field.TypeInt, value)
	}
	if cuo.mutation.EndCleared() {
		_spec.ClearField(chapter.FieldEnd, field.TypeInt)
	}
	if cuo.mutation.VodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chapter.VodTable,
			Columns: []string{chapter.VodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vod.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.VodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chapter.VodTable,
			Columns: []string{chapter.VodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vod.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Chapter{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chapter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
