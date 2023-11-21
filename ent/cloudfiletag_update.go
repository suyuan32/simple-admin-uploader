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
	"github.com/suyuan32/simple-admin-file/ent/cloudfile"
	"github.com/suyuan32/simple-admin-file/ent/cloudfiletag"
	"github.com/suyuan32/simple-admin-file/ent/predicate"
)

// CloudFileTagUpdate is the builder for updating CloudFileTag entities.
type CloudFileTagUpdate struct {
	config
	hooks    []Hook
	mutation *CloudFileTagMutation
}

// Where appends a list predicates to the CloudFileTagUpdate builder.
func (cftu *CloudFileTagUpdate) Where(ps ...predicate.CloudFileTag) *CloudFileTagUpdate {
	cftu.mutation.Where(ps...)
	return cftu
}

// SetUpdatedAt sets the "updated_at" field.
func (cftu *CloudFileTagUpdate) SetUpdatedAt(t time.Time) *CloudFileTagUpdate {
	cftu.mutation.SetUpdatedAt(t)
	return cftu
}

// SetStatus sets the "status" field.
func (cftu *CloudFileTagUpdate) SetStatus(u uint8) *CloudFileTagUpdate {
	cftu.mutation.ResetStatus()
	cftu.mutation.SetStatus(u)
	return cftu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cftu *CloudFileTagUpdate) SetNillableStatus(u *uint8) *CloudFileTagUpdate {
	if u != nil {
		cftu.SetStatus(*u)
	}
	return cftu
}

// AddStatus adds u to the "status" field.
func (cftu *CloudFileTagUpdate) AddStatus(u int8) *CloudFileTagUpdate {
	cftu.mutation.AddStatus(u)
	return cftu
}

// ClearStatus clears the value of the "status" field.
func (cftu *CloudFileTagUpdate) ClearStatus() *CloudFileTagUpdate {
	cftu.mutation.ClearStatus()
	return cftu
}

// SetName sets the "name" field.
func (cftu *CloudFileTagUpdate) SetName(s string) *CloudFileTagUpdate {
	cftu.mutation.SetName(s)
	return cftu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cftu *CloudFileTagUpdate) SetNillableName(s *string) *CloudFileTagUpdate {
	if s != nil {
		cftu.SetName(*s)
	}
	return cftu
}

// SetRemark sets the "remark" field.
func (cftu *CloudFileTagUpdate) SetRemark(s string) *CloudFileTagUpdate {
	cftu.mutation.SetRemark(s)
	return cftu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (cftu *CloudFileTagUpdate) SetNillableRemark(s *string) *CloudFileTagUpdate {
	if s != nil {
		cftu.SetRemark(*s)
	}
	return cftu
}

// ClearRemark clears the value of the "remark" field.
func (cftu *CloudFileTagUpdate) ClearRemark() *CloudFileTagUpdate {
	cftu.mutation.ClearRemark()
	return cftu
}

// AddCloudFileIDs adds the "cloud_files" edge to the CloudFile entity by IDs.
func (cftu *CloudFileTagUpdate) AddCloudFileIDs(ids ...uuid.UUID) *CloudFileTagUpdate {
	cftu.mutation.AddCloudFileIDs(ids...)
	return cftu
}

// AddCloudFiles adds the "cloud_files" edges to the CloudFile entity.
func (cftu *CloudFileTagUpdate) AddCloudFiles(c ...*CloudFile) *CloudFileTagUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cftu.AddCloudFileIDs(ids...)
}

// Mutation returns the CloudFileTagMutation object of the builder.
func (cftu *CloudFileTagUpdate) Mutation() *CloudFileTagMutation {
	return cftu.mutation
}

// ClearCloudFiles clears all "cloud_files" edges to the CloudFile entity.
func (cftu *CloudFileTagUpdate) ClearCloudFiles() *CloudFileTagUpdate {
	cftu.mutation.ClearCloudFiles()
	return cftu
}

// RemoveCloudFileIDs removes the "cloud_files" edge to CloudFile entities by IDs.
func (cftu *CloudFileTagUpdate) RemoveCloudFileIDs(ids ...uuid.UUID) *CloudFileTagUpdate {
	cftu.mutation.RemoveCloudFileIDs(ids...)
	return cftu
}

// RemoveCloudFiles removes "cloud_files" edges to CloudFile entities.
func (cftu *CloudFileTagUpdate) RemoveCloudFiles(c ...*CloudFile) *CloudFileTagUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cftu.RemoveCloudFileIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cftu *CloudFileTagUpdate) Save(ctx context.Context) (int, error) {
	cftu.defaults()
	return withHooks(ctx, cftu.sqlSave, cftu.mutation, cftu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cftu *CloudFileTagUpdate) SaveX(ctx context.Context) int {
	affected, err := cftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cftu *CloudFileTagUpdate) Exec(ctx context.Context) error {
	_, err := cftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cftu *CloudFileTagUpdate) ExecX(ctx context.Context) {
	if err := cftu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cftu *CloudFileTagUpdate) defaults() {
	if _, ok := cftu.mutation.UpdatedAt(); !ok {
		v := cloudfiletag.UpdateDefaultUpdatedAt()
		cftu.mutation.SetUpdatedAt(v)
	}
}

func (cftu *CloudFileTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(cloudfiletag.Table, cloudfiletag.Columns, sqlgraph.NewFieldSpec(cloudfiletag.FieldID, field.TypeUint64))
	if ps := cftu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cftu.mutation.UpdatedAt(); ok {
		_spec.SetField(cloudfiletag.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cftu.mutation.Status(); ok {
		_spec.SetField(cloudfiletag.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := cftu.mutation.AddedStatus(); ok {
		_spec.AddField(cloudfiletag.FieldStatus, field.TypeUint8, value)
	}
	if cftu.mutation.StatusCleared() {
		_spec.ClearField(cloudfiletag.FieldStatus, field.TypeUint8)
	}
	if value, ok := cftu.mutation.Name(); ok {
		_spec.SetField(cloudfiletag.FieldName, field.TypeString, value)
	}
	if value, ok := cftu.mutation.Remark(); ok {
		_spec.SetField(cloudfiletag.FieldRemark, field.TypeString, value)
	}
	if cftu.mutation.RemarkCleared() {
		_spec.ClearField(cloudfiletag.FieldRemark, field.TypeString)
	}
	if cftu.mutation.CloudFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cloudfiletag.CloudFilesTable,
			Columns: cloudfiletag.CloudFilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cloudfile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cftu.mutation.RemovedCloudFilesIDs(); len(nodes) > 0 && !cftu.mutation.CloudFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cloudfiletag.CloudFilesTable,
			Columns: cloudfiletag.CloudFilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cloudfile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cftu.mutation.CloudFilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cloudfiletag.CloudFilesTable,
			Columns: cloudfiletag.CloudFilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cloudfile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cftu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cloudfiletag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cftu.mutation.done = true
	return n, nil
}

// CloudFileTagUpdateOne is the builder for updating a single CloudFileTag entity.
type CloudFileTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CloudFileTagMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cftuo *CloudFileTagUpdateOne) SetUpdatedAt(t time.Time) *CloudFileTagUpdateOne {
	cftuo.mutation.SetUpdatedAt(t)
	return cftuo
}

// SetStatus sets the "status" field.
func (cftuo *CloudFileTagUpdateOne) SetStatus(u uint8) *CloudFileTagUpdateOne {
	cftuo.mutation.ResetStatus()
	cftuo.mutation.SetStatus(u)
	return cftuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cftuo *CloudFileTagUpdateOne) SetNillableStatus(u *uint8) *CloudFileTagUpdateOne {
	if u != nil {
		cftuo.SetStatus(*u)
	}
	return cftuo
}

// AddStatus adds u to the "status" field.
func (cftuo *CloudFileTagUpdateOne) AddStatus(u int8) *CloudFileTagUpdateOne {
	cftuo.mutation.AddStatus(u)
	return cftuo
}

// ClearStatus clears the value of the "status" field.
func (cftuo *CloudFileTagUpdateOne) ClearStatus() *CloudFileTagUpdateOne {
	cftuo.mutation.ClearStatus()
	return cftuo
}

// SetName sets the "name" field.
func (cftuo *CloudFileTagUpdateOne) SetName(s string) *CloudFileTagUpdateOne {
	cftuo.mutation.SetName(s)
	return cftuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cftuo *CloudFileTagUpdateOne) SetNillableName(s *string) *CloudFileTagUpdateOne {
	if s != nil {
		cftuo.SetName(*s)
	}
	return cftuo
}

// SetRemark sets the "remark" field.
func (cftuo *CloudFileTagUpdateOne) SetRemark(s string) *CloudFileTagUpdateOne {
	cftuo.mutation.SetRemark(s)
	return cftuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (cftuo *CloudFileTagUpdateOne) SetNillableRemark(s *string) *CloudFileTagUpdateOne {
	if s != nil {
		cftuo.SetRemark(*s)
	}
	return cftuo
}

// ClearRemark clears the value of the "remark" field.
func (cftuo *CloudFileTagUpdateOne) ClearRemark() *CloudFileTagUpdateOne {
	cftuo.mutation.ClearRemark()
	return cftuo
}

// AddCloudFileIDs adds the "cloud_files" edge to the CloudFile entity by IDs.
func (cftuo *CloudFileTagUpdateOne) AddCloudFileIDs(ids ...uuid.UUID) *CloudFileTagUpdateOne {
	cftuo.mutation.AddCloudFileIDs(ids...)
	return cftuo
}

// AddCloudFiles adds the "cloud_files" edges to the CloudFile entity.
func (cftuo *CloudFileTagUpdateOne) AddCloudFiles(c ...*CloudFile) *CloudFileTagUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cftuo.AddCloudFileIDs(ids...)
}

// Mutation returns the CloudFileTagMutation object of the builder.
func (cftuo *CloudFileTagUpdateOne) Mutation() *CloudFileTagMutation {
	return cftuo.mutation
}

// ClearCloudFiles clears all "cloud_files" edges to the CloudFile entity.
func (cftuo *CloudFileTagUpdateOne) ClearCloudFiles() *CloudFileTagUpdateOne {
	cftuo.mutation.ClearCloudFiles()
	return cftuo
}

// RemoveCloudFileIDs removes the "cloud_files" edge to CloudFile entities by IDs.
func (cftuo *CloudFileTagUpdateOne) RemoveCloudFileIDs(ids ...uuid.UUID) *CloudFileTagUpdateOne {
	cftuo.mutation.RemoveCloudFileIDs(ids...)
	return cftuo
}

// RemoveCloudFiles removes "cloud_files" edges to CloudFile entities.
func (cftuo *CloudFileTagUpdateOne) RemoveCloudFiles(c ...*CloudFile) *CloudFileTagUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cftuo.RemoveCloudFileIDs(ids...)
}

// Where appends a list predicates to the CloudFileTagUpdate builder.
func (cftuo *CloudFileTagUpdateOne) Where(ps ...predicate.CloudFileTag) *CloudFileTagUpdateOne {
	cftuo.mutation.Where(ps...)
	return cftuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cftuo *CloudFileTagUpdateOne) Select(field string, fields ...string) *CloudFileTagUpdateOne {
	cftuo.fields = append([]string{field}, fields...)
	return cftuo
}

// Save executes the query and returns the updated CloudFileTag entity.
func (cftuo *CloudFileTagUpdateOne) Save(ctx context.Context) (*CloudFileTag, error) {
	cftuo.defaults()
	return withHooks(ctx, cftuo.sqlSave, cftuo.mutation, cftuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cftuo *CloudFileTagUpdateOne) SaveX(ctx context.Context) *CloudFileTag {
	node, err := cftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cftuo *CloudFileTagUpdateOne) Exec(ctx context.Context) error {
	_, err := cftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cftuo *CloudFileTagUpdateOne) ExecX(ctx context.Context) {
	if err := cftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cftuo *CloudFileTagUpdateOne) defaults() {
	if _, ok := cftuo.mutation.UpdatedAt(); !ok {
		v := cloudfiletag.UpdateDefaultUpdatedAt()
		cftuo.mutation.SetUpdatedAt(v)
	}
}

func (cftuo *CloudFileTagUpdateOne) sqlSave(ctx context.Context) (_node *CloudFileTag, err error) {
	_spec := sqlgraph.NewUpdateSpec(cloudfiletag.Table, cloudfiletag.Columns, sqlgraph.NewFieldSpec(cloudfiletag.FieldID, field.TypeUint64))
	id, ok := cftuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CloudFileTag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cftuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cloudfiletag.FieldID)
		for _, f := range fields {
			if !cloudfiletag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cloudfiletag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cftuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cftuo.mutation.UpdatedAt(); ok {
		_spec.SetField(cloudfiletag.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cftuo.mutation.Status(); ok {
		_spec.SetField(cloudfiletag.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := cftuo.mutation.AddedStatus(); ok {
		_spec.AddField(cloudfiletag.FieldStatus, field.TypeUint8, value)
	}
	if cftuo.mutation.StatusCleared() {
		_spec.ClearField(cloudfiletag.FieldStatus, field.TypeUint8)
	}
	if value, ok := cftuo.mutation.Name(); ok {
		_spec.SetField(cloudfiletag.FieldName, field.TypeString, value)
	}
	if value, ok := cftuo.mutation.Remark(); ok {
		_spec.SetField(cloudfiletag.FieldRemark, field.TypeString, value)
	}
	if cftuo.mutation.RemarkCleared() {
		_spec.ClearField(cloudfiletag.FieldRemark, field.TypeString)
	}
	if cftuo.mutation.CloudFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cloudfiletag.CloudFilesTable,
			Columns: cloudfiletag.CloudFilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cloudfile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cftuo.mutation.RemovedCloudFilesIDs(); len(nodes) > 0 && !cftuo.mutation.CloudFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cloudfiletag.CloudFilesTable,
			Columns: cloudfiletag.CloudFilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cloudfile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cftuo.mutation.CloudFilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cloudfiletag.CloudFilesTable,
			Columns: cloudfiletag.CloudFilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cloudfile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CloudFileTag{config: cftuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cftuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cloudfiletag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cftuo.mutation.done = true
	return _node, nil
}
