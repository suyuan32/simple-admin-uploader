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
	"github.com/suyuan32/simple-admin-file/ent/predicate"
	"github.com/suyuan32/simple-admin-file/ent/storageprovider"
)

// StorageProviderUpdate is the builder for updating StorageProvider entities.
type StorageProviderUpdate struct {
	config
	hooks    []Hook
	mutation *StorageProviderMutation
}

// Where appends a list predicates to the StorageProviderUpdate builder.
func (spu *StorageProviderUpdate) Where(ps ...predicate.StorageProvider) *StorageProviderUpdate {
	spu.mutation.Where(ps...)
	return spu
}

// SetUpdatedAt sets the "updated_at" field.
func (spu *StorageProviderUpdate) SetUpdatedAt(t time.Time) *StorageProviderUpdate {
	spu.mutation.SetUpdatedAt(t)
	return spu
}

// SetState sets the "state" field.
func (spu *StorageProviderUpdate) SetState(b bool) *StorageProviderUpdate {
	spu.mutation.SetState(b)
	return spu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (spu *StorageProviderUpdate) SetNillableState(b *bool) *StorageProviderUpdate {
	if b != nil {
		spu.SetState(*b)
	}
	return spu
}

// ClearState clears the value of the "state" field.
func (spu *StorageProviderUpdate) ClearState() *StorageProviderUpdate {
	spu.mutation.ClearState()
	return spu
}

// SetName sets the "name" field.
func (spu *StorageProviderUpdate) SetName(s string) *StorageProviderUpdate {
	spu.mutation.SetName(s)
	return spu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (spu *StorageProviderUpdate) SetNillableName(s *string) *StorageProviderUpdate {
	if s != nil {
		spu.SetName(*s)
	}
	return spu
}

// SetBucket sets the "bucket" field.
func (spu *StorageProviderUpdate) SetBucket(s string) *StorageProviderUpdate {
	spu.mutation.SetBucket(s)
	return spu
}

// SetNillableBucket sets the "bucket" field if the given value is not nil.
func (spu *StorageProviderUpdate) SetNillableBucket(s *string) *StorageProviderUpdate {
	if s != nil {
		spu.SetBucket(*s)
	}
	return spu
}

// SetSecretID sets the "secret_id" field.
func (spu *StorageProviderUpdate) SetSecretID(s string) *StorageProviderUpdate {
	spu.mutation.SetSecretID(s)
	return spu
}

// SetNillableSecretID sets the "secret_id" field if the given value is not nil.
func (spu *StorageProviderUpdate) SetNillableSecretID(s *string) *StorageProviderUpdate {
	if s != nil {
		spu.SetSecretID(*s)
	}
	return spu
}

// SetSecretKey sets the "secret_key" field.
func (spu *StorageProviderUpdate) SetSecretKey(s string) *StorageProviderUpdate {
	spu.mutation.SetSecretKey(s)
	return spu
}

// SetNillableSecretKey sets the "secret_key" field if the given value is not nil.
func (spu *StorageProviderUpdate) SetNillableSecretKey(s *string) *StorageProviderUpdate {
	if s != nil {
		spu.SetSecretKey(*s)
	}
	return spu
}

// SetEndpoint sets the "endpoint" field.
func (spu *StorageProviderUpdate) SetEndpoint(s string) *StorageProviderUpdate {
	spu.mutation.SetEndpoint(s)
	return spu
}

// SetNillableEndpoint sets the "endpoint" field if the given value is not nil.
func (spu *StorageProviderUpdate) SetNillableEndpoint(s *string) *StorageProviderUpdate {
	if s != nil {
		spu.SetEndpoint(*s)
	}
	return spu
}

// SetFolder sets the "folder" field.
func (spu *StorageProviderUpdate) SetFolder(s string) *StorageProviderUpdate {
	spu.mutation.SetFolder(s)
	return spu
}

// SetNillableFolder sets the "folder" field if the given value is not nil.
func (spu *StorageProviderUpdate) SetNillableFolder(s *string) *StorageProviderUpdate {
	if s != nil {
		spu.SetFolder(*s)
	}
	return spu
}

// ClearFolder clears the value of the "folder" field.
func (spu *StorageProviderUpdate) ClearFolder() *StorageProviderUpdate {
	spu.mutation.ClearFolder()
	return spu
}

// SetRegion sets the "region" field.
func (spu *StorageProviderUpdate) SetRegion(s string) *StorageProviderUpdate {
	spu.mutation.SetRegion(s)
	return spu
}

// SetNillableRegion sets the "region" field if the given value is not nil.
func (spu *StorageProviderUpdate) SetNillableRegion(s *string) *StorageProviderUpdate {
	if s != nil {
		spu.SetRegion(*s)
	}
	return spu
}

// SetIsDefault sets the "is_default" field.
func (spu *StorageProviderUpdate) SetIsDefault(b bool) *StorageProviderUpdate {
	spu.mutation.SetIsDefault(b)
	return spu
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (spu *StorageProviderUpdate) SetNillableIsDefault(b *bool) *StorageProviderUpdate {
	if b != nil {
		spu.SetIsDefault(*b)
	}
	return spu
}

// AddCloudfileIDs adds the "cloudfiles" edge to the CloudFile entity by IDs.
func (spu *StorageProviderUpdate) AddCloudfileIDs(ids ...uuid.UUID) *StorageProviderUpdate {
	spu.mutation.AddCloudfileIDs(ids...)
	return spu
}

// AddCloudfiles adds the "cloudfiles" edges to the CloudFile entity.
func (spu *StorageProviderUpdate) AddCloudfiles(c ...*CloudFile) *StorageProviderUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return spu.AddCloudfileIDs(ids...)
}

// Mutation returns the StorageProviderMutation object of the builder.
func (spu *StorageProviderUpdate) Mutation() *StorageProviderMutation {
	return spu.mutation
}

// ClearCloudfiles clears all "cloudfiles" edges to the CloudFile entity.
func (spu *StorageProviderUpdate) ClearCloudfiles() *StorageProviderUpdate {
	spu.mutation.ClearCloudfiles()
	return spu
}

// RemoveCloudfileIDs removes the "cloudfiles" edge to CloudFile entities by IDs.
func (spu *StorageProviderUpdate) RemoveCloudfileIDs(ids ...uuid.UUID) *StorageProviderUpdate {
	spu.mutation.RemoveCloudfileIDs(ids...)
	return spu
}

// RemoveCloudfiles removes "cloudfiles" edges to CloudFile entities.
func (spu *StorageProviderUpdate) RemoveCloudfiles(c ...*CloudFile) *StorageProviderUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return spu.RemoveCloudfileIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (spu *StorageProviderUpdate) Save(ctx context.Context) (int, error) {
	spu.defaults()
	return withHooks(ctx, spu.sqlSave, spu.mutation, spu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (spu *StorageProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := spu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (spu *StorageProviderUpdate) Exec(ctx context.Context) error {
	_, err := spu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spu *StorageProviderUpdate) ExecX(ctx context.Context) {
	if err := spu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spu *StorageProviderUpdate) defaults() {
	if _, ok := spu.mutation.UpdatedAt(); !ok {
		v := storageprovider.UpdateDefaultUpdatedAt()
		spu.mutation.SetUpdatedAt(v)
	}
}

func (spu *StorageProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(storageprovider.Table, storageprovider.Columns, sqlgraph.NewFieldSpec(storageprovider.FieldID, field.TypeUint64))
	if ps := spu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spu.mutation.UpdatedAt(); ok {
		_spec.SetField(storageprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := spu.mutation.State(); ok {
		_spec.SetField(storageprovider.FieldState, field.TypeBool, value)
	}
	if spu.mutation.StateCleared() {
		_spec.ClearField(storageprovider.FieldState, field.TypeBool)
	}
	if value, ok := spu.mutation.Name(); ok {
		_spec.SetField(storageprovider.FieldName, field.TypeString, value)
	}
	if value, ok := spu.mutation.Bucket(); ok {
		_spec.SetField(storageprovider.FieldBucket, field.TypeString, value)
	}
	if value, ok := spu.mutation.SecretID(); ok {
		_spec.SetField(storageprovider.FieldSecretID, field.TypeString, value)
	}
	if value, ok := spu.mutation.SecretKey(); ok {
		_spec.SetField(storageprovider.FieldSecretKey, field.TypeString, value)
	}
	if value, ok := spu.mutation.Endpoint(); ok {
		_spec.SetField(storageprovider.FieldEndpoint, field.TypeString, value)
	}
	if value, ok := spu.mutation.Folder(); ok {
		_spec.SetField(storageprovider.FieldFolder, field.TypeString, value)
	}
	if spu.mutation.FolderCleared() {
		_spec.ClearField(storageprovider.FieldFolder, field.TypeString)
	}
	if value, ok := spu.mutation.Region(); ok {
		_spec.SetField(storageprovider.FieldRegion, field.TypeString, value)
	}
	if value, ok := spu.mutation.IsDefault(); ok {
		_spec.SetField(storageprovider.FieldIsDefault, field.TypeBool, value)
	}
	if spu.mutation.CloudfilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   storageprovider.CloudfilesTable,
			Columns: []string{storageprovider.CloudfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cloudfile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spu.mutation.RemovedCloudfilesIDs(); len(nodes) > 0 && !spu.mutation.CloudfilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   storageprovider.CloudfilesTable,
			Columns: []string{storageprovider.CloudfilesColumn},
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
	if nodes := spu.mutation.CloudfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   storageprovider.CloudfilesTable,
			Columns: []string{storageprovider.CloudfilesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, spu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storageprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	spu.mutation.done = true
	return n, nil
}

// StorageProviderUpdateOne is the builder for updating a single StorageProvider entity.
type StorageProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StorageProviderMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (spuo *StorageProviderUpdateOne) SetUpdatedAt(t time.Time) *StorageProviderUpdateOne {
	spuo.mutation.SetUpdatedAt(t)
	return spuo
}

// SetState sets the "state" field.
func (spuo *StorageProviderUpdateOne) SetState(b bool) *StorageProviderUpdateOne {
	spuo.mutation.SetState(b)
	return spuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (spuo *StorageProviderUpdateOne) SetNillableState(b *bool) *StorageProviderUpdateOne {
	if b != nil {
		spuo.SetState(*b)
	}
	return spuo
}

// ClearState clears the value of the "state" field.
func (spuo *StorageProviderUpdateOne) ClearState() *StorageProviderUpdateOne {
	spuo.mutation.ClearState()
	return spuo
}

// SetName sets the "name" field.
func (spuo *StorageProviderUpdateOne) SetName(s string) *StorageProviderUpdateOne {
	spuo.mutation.SetName(s)
	return spuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (spuo *StorageProviderUpdateOne) SetNillableName(s *string) *StorageProviderUpdateOne {
	if s != nil {
		spuo.SetName(*s)
	}
	return spuo
}

// SetBucket sets the "bucket" field.
func (spuo *StorageProviderUpdateOne) SetBucket(s string) *StorageProviderUpdateOne {
	spuo.mutation.SetBucket(s)
	return spuo
}

// SetNillableBucket sets the "bucket" field if the given value is not nil.
func (spuo *StorageProviderUpdateOne) SetNillableBucket(s *string) *StorageProviderUpdateOne {
	if s != nil {
		spuo.SetBucket(*s)
	}
	return spuo
}

// SetSecretID sets the "secret_id" field.
func (spuo *StorageProviderUpdateOne) SetSecretID(s string) *StorageProviderUpdateOne {
	spuo.mutation.SetSecretID(s)
	return spuo
}

// SetNillableSecretID sets the "secret_id" field if the given value is not nil.
func (spuo *StorageProviderUpdateOne) SetNillableSecretID(s *string) *StorageProviderUpdateOne {
	if s != nil {
		spuo.SetSecretID(*s)
	}
	return spuo
}

// SetSecretKey sets the "secret_key" field.
func (spuo *StorageProviderUpdateOne) SetSecretKey(s string) *StorageProviderUpdateOne {
	spuo.mutation.SetSecretKey(s)
	return spuo
}

// SetNillableSecretKey sets the "secret_key" field if the given value is not nil.
func (spuo *StorageProviderUpdateOne) SetNillableSecretKey(s *string) *StorageProviderUpdateOne {
	if s != nil {
		spuo.SetSecretKey(*s)
	}
	return spuo
}

// SetEndpoint sets the "endpoint" field.
func (spuo *StorageProviderUpdateOne) SetEndpoint(s string) *StorageProviderUpdateOne {
	spuo.mutation.SetEndpoint(s)
	return spuo
}

// SetNillableEndpoint sets the "endpoint" field if the given value is not nil.
func (spuo *StorageProviderUpdateOne) SetNillableEndpoint(s *string) *StorageProviderUpdateOne {
	if s != nil {
		spuo.SetEndpoint(*s)
	}
	return spuo
}

// SetFolder sets the "folder" field.
func (spuo *StorageProviderUpdateOne) SetFolder(s string) *StorageProviderUpdateOne {
	spuo.mutation.SetFolder(s)
	return spuo
}

// SetNillableFolder sets the "folder" field if the given value is not nil.
func (spuo *StorageProviderUpdateOne) SetNillableFolder(s *string) *StorageProviderUpdateOne {
	if s != nil {
		spuo.SetFolder(*s)
	}
	return spuo
}

// ClearFolder clears the value of the "folder" field.
func (spuo *StorageProviderUpdateOne) ClearFolder() *StorageProviderUpdateOne {
	spuo.mutation.ClearFolder()
	return spuo
}

// SetRegion sets the "region" field.
func (spuo *StorageProviderUpdateOne) SetRegion(s string) *StorageProviderUpdateOne {
	spuo.mutation.SetRegion(s)
	return spuo
}

// SetNillableRegion sets the "region" field if the given value is not nil.
func (spuo *StorageProviderUpdateOne) SetNillableRegion(s *string) *StorageProviderUpdateOne {
	if s != nil {
		spuo.SetRegion(*s)
	}
	return spuo
}

// SetIsDefault sets the "is_default" field.
func (spuo *StorageProviderUpdateOne) SetIsDefault(b bool) *StorageProviderUpdateOne {
	spuo.mutation.SetIsDefault(b)
	return spuo
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (spuo *StorageProviderUpdateOne) SetNillableIsDefault(b *bool) *StorageProviderUpdateOne {
	if b != nil {
		spuo.SetIsDefault(*b)
	}
	return spuo
}

// AddCloudfileIDs adds the "cloudfiles" edge to the CloudFile entity by IDs.
func (spuo *StorageProviderUpdateOne) AddCloudfileIDs(ids ...uuid.UUID) *StorageProviderUpdateOne {
	spuo.mutation.AddCloudfileIDs(ids...)
	return spuo
}

// AddCloudfiles adds the "cloudfiles" edges to the CloudFile entity.
func (spuo *StorageProviderUpdateOne) AddCloudfiles(c ...*CloudFile) *StorageProviderUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return spuo.AddCloudfileIDs(ids...)
}

// Mutation returns the StorageProviderMutation object of the builder.
func (spuo *StorageProviderUpdateOne) Mutation() *StorageProviderMutation {
	return spuo.mutation
}

// ClearCloudfiles clears all "cloudfiles" edges to the CloudFile entity.
func (spuo *StorageProviderUpdateOne) ClearCloudfiles() *StorageProviderUpdateOne {
	spuo.mutation.ClearCloudfiles()
	return spuo
}

// RemoveCloudfileIDs removes the "cloudfiles" edge to CloudFile entities by IDs.
func (spuo *StorageProviderUpdateOne) RemoveCloudfileIDs(ids ...uuid.UUID) *StorageProviderUpdateOne {
	spuo.mutation.RemoveCloudfileIDs(ids...)
	return spuo
}

// RemoveCloudfiles removes "cloudfiles" edges to CloudFile entities.
func (spuo *StorageProviderUpdateOne) RemoveCloudfiles(c ...*CloudFile) *StorageProviderUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return spuo.RemoveCloudfileIDs(ids...)
}

// Where appends a list predicates to the StorageProviderUpdate builder.
func (spuo *StorageProviderUpdateOne) Where(ps ...predicate.StorageProvider) *StorageProviderUpdateOne {
	spuo.mutation.Where(ps...)
	return spuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (spuo *StorageProviderUpdateOne) Select(field string, fields ...string) *StorageProviderUpdateOne {
	spuo.fields = append([]string{field}, fields...)
	return spuo
}

// Save executes the query and returns the updated StorageProvider entity.
func (spuo *StorageProviderUpdateOne) Save(ctx context.Context) (*StorageProvider, error) {
	spuo.defaults()
	return withHooks(ctx, spuo.sqlSave, spuo.mutation, spuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (spuo *StorageProviderUpdateOne) SaveX(ctx context.Context) *StorageProvider {
	node, err := spuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (spuo *StorageProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := spuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spuo *StorageProviderUpdateOne) ExecX(ctx context.Context) {
	if err := spuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spuo *StorageProviderUpdateOne) defaults() {
	if _, ok := spuo.mutation.UpdatedAt(); !ok {
		v := storageprovider.UpdateDefaultUpdatedAt()
		spuo.mutation.SetUpdatedAt(v)
	}
}

func (spuo *StorageProviderUpdateOne) sqlSave(ctx context.Context) (_node *StorageProvider, err error) {
	_spec := sqlgraph.NewUpdateSpec(storageprovider.Table, storageprovider.Columns, sqlgraph.NewFieldSpec(storageprovider.FieldID, field.TypeUint64))
	id, ok := spuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "StorageProvider.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := spuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, storageprovider.FieldID)
		for _, f := range fields {
			if !storageprovider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != storageprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := spuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spuo.mutation.UpdatedAt(); ok {
		_spec.SetField(storageprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := spuo.mutation.State(); ok {
		_spec.SetField(storageprovider.FieldState, field.TypeBool, value)
	}
	if spuo.mutation.StateCleared() {
		_spec.ClearField(storageprovider.FieldState, field.TypeBool)
	}
	if value, ok := spuo.mutation.Name(); ok {
		_spec.SetField(storageprovider.FieldName, field.TypeString, value)
	}
	if value, ok := spuo.mutation.Bucket(); ok {
		_spec.SetField(storageprovider.FieldBucket, field.TypeString, value)
	}
	if value, ok := spuo.mutation.SecretID(); ok {
		_spec.SetField(storageprovider.FieldSecretID, field.TypeString, value)
	}
	if value, ok := spuo.mutation.SecretKey(); ok {
		_spec.SetField(storageprovider.FieldSecretKey, field.TypeString, value)
	}
	if value, ok := spuo.mutation.Endpoint(); ok {
		_spec.SetField(storageprovider.FieldEndpoint, field.TypeString, value)
	}
	if value, ok := spuo.mutation.Folder(); ok {
		_spec.SetField(storageprovider.FieldFolder, field.TypeString, value)
	}
	if spuo.mutation.FolderCleared() {
		_spec.ClearField(storageprovider.FieldFolder, field.TypeString)
	}
	if value, ok := spuo.mutation.Region(); ok {
		_spec.SetField(storageprovider.FieldRegion, field.TypeString, value)
	}
	if value, ok := spuo.mutation.IsDefault(); ok {
		_spec.SetField(storageprovider.FieldIsDefault, field.TypeBool, value)
	}
	if spuo.mutation.CloudfilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   storageprovider.CloudfilesTable,
			Columns: []string{storageprovider.CloudfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cloudfile.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spuo.mutation.RemovedCloudfilesIDs(); len(nodes) > 0 && !spuo.mutation.CloudfilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   storageprovider.CloudfilesTable,
			Columns: []string{storageprovider.CloudfilesColumn},
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
	if nodes := spuo.mutation.CloudfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   storageprovider.CloudfilesTable,
			Columns: []string{storageprovider.CloudfilesColumn},
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
	_node = &StorageProvider{config: spuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, spuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storageprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	spuo.mutation.done = true
	return _node, nil
}
