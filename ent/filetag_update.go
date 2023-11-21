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
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-file/ent/file"
	"github.com/suyuan32/simple-admin-file/ent/filetag"
	"github.com/suyuan32/simple-admin-file/ent/predicate"
)

// FileTagUpdate is the builder for updating FileTag entities.
type FileTagUpdate struct {
	config
	hooks    []Hook
	mutation *FileTagMutation
}

// Where appends a list predicates to the FileTagUpdate builder.
func (ftu *FileTagUpdate) Where(ps ...predicate.FileTag) *FileTagUpdate {
	ftu.mutation.Where(ps...)
	return ftu
}

// SetUpdatedAt sets the "updated_at" field.
func (ftu *FileTagUpdate) SetUpdatedAt(t time.Time) *FileTagUpdate {
	ftu.mutation.SetUpdatedAt(t)
	return ftu
}

// SetStatus sets the "status" field.
func (ftu *FileTagUpdate) SetStatus(u uint8) *FileTagUpdate {
	ftu.mutation.ResetStatus()
	ftu.mutation.SetStatus(u)
	return ftu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ftu *FileTagUpdate) SetNillableStatus(u *uint8) *FileTagUpdate {
	if u != nil {
		ftu.SetStatus(*u)
	}
	return ftu
}

// AddStatus adds u to the "status" field.
func (ftu *FileTagUpdate) AddStatus(u int8) *FileTagUpdate {
	ftu.mutation.AddStatus(u)
	return ftu
}

// ClearStatus clears the value of the "status" field.
func (ftu *FileTagUpdate) ClearStatus() *FileTagUpdate {
	ftu.mutation.ClearStatus()
	return ftu
}

// SetName sets the "name" field.
func (ftu *FileTagUpdate) SetName(s string) *FileTagUpdate {
	ftu.mutation.SetName(s)
	return ftu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ftu *FileTagUpdate) SetNillableName(s *string) *FileTagUpdate {
	if s != nil {
		ftu.SetName(*s)
	}
	return ftu
}

// SetRemark sets the "remark" field.
func (ftu *FileTagUpdate) SetRemark(s string) *FileTagUpdate {
	ftu.mutation.SetRemark(s)
	return ftu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ftu *FileTagUpdate) SetNillableRemark(s *string) *FileTagUpdate {
	if s != nil {
		ftu.SetRemark(*s)
	}
	return ftu
}

// ClearRemark clears the value of the "remark" field.
func (ftu *FileTagUpdate) ClearRemark() *FileTagUpdate {
	ftu.mutation.ClearRemark()
	return ftu
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (ftu *FileTagUpdate) AddFileIDs(ids ...uuid.UUID) *FileTagUpdate {
	ftu.mutation.AddFileIDs(ids...)
	return ftu
}

// AddFiles adds the "files" edges to the File entity.
func (ftu *FileTagUpdate) AddFiles(f ...*File) *FileTagUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.AddFileIDs(ids...)
}

// Mutation returns the FileTagMutation object of the builder.
func (ftu *FileTagUpdate) Mutation() *FileTagMutation {
	return ftu.mutation
}

// ClearFiles clears all "files" edges to the File entity.
func (ftu *FileTagUpdate) ClearFiles() *FileTagUpdate {
	ftu.mutation.ClearFiles()
	return ftu
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (ftu *FileTagUpdate) RemoveFileIDs(ids ...uuid.UUID) *FileTagUpdate {
	ftu.mutation.RemoveFileIDs(ids...)
	return ftu
}

// RemoveFiles removes "files" edges to File entities.
func (ftu *FileTagUpdate) RemoveFiles(f ...*File) *FileTagUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.RemoveFileIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ftu *FileTagUpdate) Save(ctx context.Context) (int, error) {
	ftu.defaults()
	return withHooks(ctx, ftu.sqlSave, ftu.mutation, ftu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ftu *FileTagUpdate) SaveX(ctx context.Context) int {
	affected, err := ftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ftu *FileTagUpdate) Exec(ctx context.Context) error {
	_, err := ftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftu *FileTagUpdate) ExecX(ctx context.Context) {
	if err := ftu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftu *FileTagUpdate) defaults() {
	if _, ok := ftu.mutation.UpdatedAt(); !ok {
		v := filetag.UpdateDefaultUpdatedAt()
		ftu.mutation.SetUpdatedAt(v)
	}
}

func (ftu *FileTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(filetag.Table, filetag.Columns, sqlgraph.NewFieldSpec(filetag.FieldID, field.TypeUint64))
	if ps := ftu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftu.mutation.UpdatedAt(); ok {
		_spec.SetField(filetag.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ftu.mutation.Status(); ok {
		_spec.SetField(filetag.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := ftu.mutation.AddedStatus(); ok {
		_spec.AddField(filetag.FieldStatus, field.TypeUint8, value)
	}
	if ftu.mutation.StatusCleared() {
		_spec.ClearField(filetag.FieldStatus, field.TypeUint8)
	}
	if value, ok := ftu.mutation.Name(); ok {
		_spec.SetField(filetag.FieldName, field.TypeString, value)
	}
	if value, ok := ftu.mutation.Remark(); ok {
		_spec.SetField(filetag.FieldRemark, field.TypeString, value)
	}
	if ftu.mutation.RemarkCleared() {
		_spec.ClearField(filetag.FieldRemark, field.TypeString)
	}
	if ftu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   filetag.FilesTable,
			Columns: filetag.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftu.mutation.RemovedFilesIDs(); len(nodes) > 0 && !ftu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   filetag.FilesTable,
			Columns: filetag.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftu.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   filetag.FilesTable,
			Columns: filetag.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ftu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filetag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ftu.mutation.done = true
	return n, nil
}

// FileTagUpdateOne is the builder for updating a single FileTag entity.
type FileTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileTagMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ftuo *FileTagUpdateOne) SetUpdatedAt(t time.Time) *FileTagUpdateOne {
	ftuo.mutation.SetUpdatedAt(t)
	return ftuo
}

// SetStatus sets the "status" field.
func (ftuo *FileTagUpdateOne) SetStatus(u uint8) *FileTagUpdateOne {
	ftuo.mutation.ResetStatus()
	ftuo.mutation.SetStatus(u)
	return ftuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ftuo *FileTagUpdateOne) SetNillableStatus(u *uint8) *FileTagUpdateOne {
	if u != nil {
		ftuo.SetStatus(*u)
	}
	return ftuo
}

// AddStatus adds u to the "status" field.
func (ftuo *FileTagUpdateOne) AddStatus(u int8) *FileTagUpdateOne {
	ftuo.mutation.AddStatus(u)
	return ftuo
}

// ClearStatus clears the value of the "status" field.
func (ftuo *FileTagUpdateOne) ClearStatus() *FileTagUpdateOne {
	ftuo.mutation.ClearStatus()
	return ftuo
}

// SetName sets the "name" field.
func (ftuo *FileTagUpdateOne) SetName(s string) *FileTagUpdateOne {
	ftuo.mutation.SetName(s)
	return ftuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ftuo *FileTagUpdateOne) SetNillableName(s *string) *FileTagUpdateOne {
	if s != nil {
		ftuo.SetName(*s)
	}
	return ftuo
}

// SetRemark sets the "remark" field.
func (ftuo *FileTagUpdateOne) SetRemark(s string) *FileTagUpdateOne {
	ftuo.mutation.SetRemark(s)
	return ftuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ftuo *FileTagUpdateOne) SetNillableRemark(s *string) *FileTagUpdateOne {
	if s != nil {
		ftuo.SetRemark(*s)
	}
	return ftuo
}

// ClearRemark clears the value of the "remark" field.
func (ftuo *FileTagUpdateOne) ClearRemark() *FileTagUpdateOne {
	ftuo.mutation.ClearRemark()
	return ftuo
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (ftuo *FileTagUpdateOne) AddFileIDs(ids ...uuid.UUID) *FileTagUpdateOne {
	ftuo.mutation.AddFileIDs(ids...)
	return ftuo
}

// AddFiles adds the "files" edges to the File entity.
func (ftuo *FileTagUpdateOne) AddFiles(f ...*File) *FileTagUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.AddFileIDs(ids...)
}

// Mutation returns the FileTagMutation object of the builder.
func (ftuo *FileTagUpdateOne) Mutation() *FileTagMutation {
	return ftuo.mutation
}

// ClearFiles clears all "files" edges to the File entity.
func (ftuo *FileTagUpdateOne) ClearFiles() *FileTagUpdateOne {
	ftuo.mutation.ClearFiles()
	return ftuo
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (ftuo *FileTagUpdateOne) RemoveFileIDs(ids ...uuid.UUID) *FileTagUpdateOne {
	ftuo.mutation.RemoveFileIDs(ids...)
	return ftuo
}

// RemoveFiles removes "files" edges to File entities.
func (ftuo *FileTagUpdateOne) RemoveFiles(f ...*File) *FileTagUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.RemoveFileIDs(ids...)
}

// Where appends a list predicates to the FileTagUpdate builder.
func (ftuo *FileTagUpdateOne) Where(ps ...predicate.FileTag) *FileTagUpdateOne {
	ftuo.mutation.Where(ps...)
	return ftuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ftuo *FileTagUpdateOne) Select(field string, fields ...string) *FileTagUpdateOne {
	ftuo.fields = append([]string{field}, fields...)
	return ftuo
}

// Save executes the query and returns the updated FileTag entity.
func (ftuo *FileTagUpdateOne) Save(ctx context.Context) (*FileTag, error) {
	ftuo.defaults()
	return withHooks(ctx, ftuo.sqlSave, ftuo.mutation, ftuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ftuo *FileTagUpdateOne) SaveX(ctx context.Context) *FileTag {
	node, err := ftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ftuo *FileTagUpdateOne) Exec(ctx context.Context) error {
	_, err := ftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftuo *FileTagUpdateOne) ExecX(ctx context.Context) {
	if err := ftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftuo *FileTagUpdateOne) defaults() {
	if _, ok := ftuo.mutation.UpdatedAt(); !ok {
		v := filetag.UpdateDefaultUpdatedAt()
		ftuo.mutation.SetUpdatedAt(v)
	}
}

func (ftuo *FileTagUpdateOne) sqlSave(ctx context.Context) (_node *FileTag, err error) {
	_spec := sqlgraph.NewUpdateSpec(filetag.Table, filetag.Columns, sqlgraph.NewFieldSpec(filetag.FieldID, field.TypeUint64))
	id, ok := ftuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FileTag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ftuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, filetag.FieldID)
		for _, f := range fields {
			if !filetag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != filetag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ftuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftuo.mutation.UpdatedAt(); ok {
		_spec.SetField(filetag.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ftuo.mutation.Status(); ok {
		_spec.SetField(filetag.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := ftuo.mutation.AddedStatus(); ok {
		_spec.AddField(filetag.FieldStatus, field.TypeUint8, value)
	}
	if ftuo.mutation.StatusCleared() {
		_spec.ClearField(filetag.FieldStatus, field.TypeUint8)
	}
	if value, ok := ftuo.mutation.Name(); ok {
		_spec.SetField(filetag.FieldName, field.TypeString, value)
	}
	if value, ok := ftuo.mutation.Remark(); ok {
		_spec.SetField(filetag.FieldRemark, field.TypeString, value)
	}
	if ftuo.mutation.RemarkCleared() {
		_spec.ClearField(filetag.FieldRemark, field.TypeString)
	}
	if ftuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   filetag.FilesTable,
			Columns: filetag.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftuo.mutation.RemovedFilesIDs(); len(nodes) > 0 && !ftuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   filetag.FilesTable,
			Columns: filetag.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftuo.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   filetag.FilesTable,
			Columns: filetag.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FileTag{config: ftuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ftuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filetag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ftuo.mutation.done = true
	return _node, nil
}
