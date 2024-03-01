// Code generated by ent, DO NOT EDIT.

package storageprovider

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/suyuan32/simple-admin-file/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldUpdatedAt, v))
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v bool) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldState, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldName, v))
}

// Bucket applies equality check predicate on the "bucket" field. It's identical to BucketEQ.
func Bucket(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldBucket, v))
}

// SecretID applies equality check predicate on the "secret_id" field. It's identical to SecretIDEQ.
func SecretID(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldSecretID, v))
}

// SecretKey applies equality check predicate on the "secret_key" field. It's identical to SecretKeyEQ.
func SecretKey(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldSecretKey, v))
}

// Endpoint applies equality check predicate on the "endpoint" field. It's identical to EndpointEQ.
func Endpoint(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldEndpoint, v))
}

// Folder applies equality check predicate on the "folder" field. It's identical to FolderEQ.
func Folder(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldFolder, v))
}

// Region applies equality check predicate on the "region" field. It's identical to RegionEQ.
func Region(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldRegion, v))
}

// IsDefault applies equality check predicate on the "is_default" field. It's identical to IsDefaultEQ.
func IsDefault(v bool) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldIsDefault, v))
}

// UseCdn applies equality check predicate on the "use_cdn" field. It's identical to UseCdnEQ.
func UseCdn(v bool) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldUseCdn, v))
}

// CdnURL applies equality check predicate on the "cdn_url" field. It's identical to CdnURLEQ.
func CdnURL(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldCdnURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLTE(FieldUpdatedAt, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v bool) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v bool) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNEQ(FieldState, v))
}

// StateIsNil applies the IsNil predicate on the "state" field.
func StateIsNil() predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldIsNull(FieldState))
}

// StateNotNil applies the NotNil predicate on the "state" field.
func StateNotNil() predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNotNull(FieldState))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContainsFold(FieldName, v))
}

// BucketEQ applies the EQ predicate on the "bucket" field.
func BucketEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldBucket, v))
}

// BucketNEQ applies the NEQ predicate on the "bucket" field.
func BucketNEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNEQ(FieldBucket, v))
}

// BucketIn applies the In predicate on the "bucket" field.
func BucketIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldIn(FieldBucket, vs...))
}

// BucketNotIn applies the NotIn predicate on the "bucket" field.
func BucketNotIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNotIn(FieldBucket, vs...))
}

// BucketGT applies the GT predicate on the "bucket" field.
func BucketGT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGT(FieldBucket, v))
}

// BucketGTE applies the GTE predicate on the "bucket" field.
func BucketGTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGTE(FieldBucket, v))
}

// BucketLT applies the LT predicate on the "bucket" field.
func BucketLT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLT(FieldBucket, v))
}

// BucketLTE applies the LTE predicate on the "bucket" field.
func BucketLTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLTE(FieldBucket, v))
}

// BucketContains applies the Contains predicate on the "bucket" field.
func BucketContains(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContains(FieldBucket, v))
}

// BucketHasPrefix applies the HasPrefix predicate on the "bucket" field.
func BucketHasPrefix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasPrefix(FieldBucket, v))
}

// BucketHasSuffix applies the HasSuffix predicate on the "bucket" field.
func BucketHasSuffix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasSuffix(FieldBucket, v))
}

// BucketEqualFold applies the EqualFold predicate on the "bucket" field.
func BucketEqualFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEqualFold(FieldBucket, v))
}

// BucketContainsFold applies the ContainsFold predicate on the "bucket" field.
func BucketContainsFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContainsFold(FieldBucket, v))
}

// SecretIDEQ applies the EQ predicate on the "secret_id" field.
func SecretIDEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldSecretID, v))
}

// SecretIDNEQ applies the NEQ predicate on the "secret_id" field.
func SecretIDNEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNEQ(FieldSecretID, v))
}

// SecretIDIn applies the In predicate on the "secret_id" field.
func SecretIDIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldIn(FieldSecretID, vs...))
}

// SecretIDNotIn applies the NotIn predicate on the "secret_id" field.
func SecretIDNotIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNotIn(FieldSecretID, vs...))
}

// SecretIDGT applies the GT predicate on the "secret_id" field.
func SecretIDGT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGT(FieldSecretID, v))
}

// SecretIDGTE applies the GTE predicate on the "secret_id" field.
func SecretIDGTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGTE(FieldSecretID, v))
}

// SecretIDLT applies the LT predicate on the "secret_id" field.
func SecretIDLT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLT(FieldSecretID, v))
}

// SecretIDLTE applies the LTE predicate on the "secret_id" field.
func SecretIDLTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLTE(FieldSecretID, v))
}

// SecretIDContains applies the Contains predicate on the "secret_id" field.
func SecretIDContains(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContains(FieldSecretID, v))
}

// SecretIDHasPrefix applies the HasPrefix predicate on the "secret_id" field.
func SecretIDHasPrefix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasPrefix(FieldSecretID, v))
}

// SecretIDHasSuffix applies the HasSuffix predicate on the "secret_id" field.
func SecretIDHasSuffix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasSuffix(FieldSecretID, v))
}

// SecretIDEqualFold applies the EqualFold predicate on the "secret_id" field.
func SecretIDEqualFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEqualFold(FieldSecretID, v))
}

// SecretIDContainsFold applies the ContainsFold predicate on the "secret_id" field.
func SecretIDContainsFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContainsFold(FieldSecretID, v))
}

// SecretKeyEQ applies the EQ predicate on the "secret_key" field.
func SecretKeyEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldSecretKey, v))
}

// SecretKeyNEQ applies the NEQ predicate on the "secret_key" field.
func SecretKeyNEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNEQ(FieldSecretKey, v))
}

// SecretKeyIn applies the In predicate on the "secret_key" field.
func SecretKeyIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldIn(FieldSecretKey, vs...))
}

// SecretKeyNotIn applies the NotIn predicate on the "secret_key" field.
func SecretKeyNotIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNotIn(FieldSecretKey, vs...))
}

// SecretKeyGT applies the GT predicate on the "secret_key" field.
func SecretKeyGT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGT(FieldSecretKey, v))
}

// SecretKeyGTE applies the GTE predicate on the "secret_key" field.
func SecretKeyGTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGTE(FieldSecretKey, v))
}

// SecretKeyLT applies the LT predicate on the "secret_key" field.
func SecretKeyLT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLT(FieldSecretKey, v))
}

// SecretKeyLTE applies the LTE predicate on the "secret_key" field.
func SecretKeyLTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLTE(FieldSecretKey, v))
}

// SecretKeyContains applies the Contains predicate on the "secret_key" field.
func SecretKeyContains(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContains(FieldSecretKey, v))
}

// SecretKeyHasPrefix applies the HasPrefix predicate on the "secret_key" field.
func SecretKeyHasPrefix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasPrefix(FieldSecretKey, v))
}

// SecretKeyHasSuffix applies the HasSuffix predicate on the "secret_key" field.
func SecretKeyHasSuffix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasSuffix(FieldSecretKey, v))
}

// SecretKeyEqualFold applies the EqualFold predicate on the "secret_key" field.
func SecretKeyEqualFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEqualFold(FieldSecretKey, v))
}

// SecretKeyContainsFold applies the ContainsFold predicate on the "secret_key" field.
func SecretKeyContainsFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContainsFold(FieldSecretKey, v))
}

// EndpointEQ applies the EQ predicate on the "endpoint" field.
func EndpointEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldEndpoint, v))
}

// EndpointNEQ applies the NEQ predicate on the "endpoint" field.
func EndpointNEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNEQ(FieldEndpoint, v))
}

// EndpointIn applies the In predicate on the "endpoint" field.
func EndpointIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldIn(FieldEndpoint, vs...))
}

// EndpointNotIn applies the NotIn predicate on the "endpoint" field.
func EndpointNotIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNotIn(FieldEndpoint, vs...))
}

// EndpointGT applies the GT predicate on the "endpoint" field.
func EndpointGT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGT(FieldEndpoint, v))
}

// EndpointGTE applies the GTE predicate on the "endpoint" field.
func EndpointGTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGTE(FieldEndpoint, v))
}

// EndpointLT applies the LT predicate on the "endpoint" field.
func EndpointLT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLT(FieldEndpoint, v))
}

// EndpointLTE applies the LTE predicate on the "endpoint" field.
func EndpointLTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLTE(FieldEndpoint, v))
}

// EndpointContains applies the Contains predicate on the "endpoint" field.
func EndpointContains(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContains(FieldEndpoint, v))
}

// EndpointHasPrefix applies the HasPrefix predicate on the "endpoint" field.
func EndpointHasPrefix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasPrefix(FieldEndpoint, v))
}

// EndpointHasSuffix applies the HasSuffix predicate on the "endpoint" field.
func EndpointHasSuffix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasSuffix(FieldEndpoint, v))
}

// EndpointEqualFold applies the EqualFold predicate on the "endpoint" field.
func EndpointEqualFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEqualFold(FieldEndpoint, v))
}

// EndpointContainsFold applies the ContainsFold predicate on the "endpoint" field.
func EndpointContainsFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContainsFold(FieldEndpoint, v))
}

// FolderEQ applies the EQ predicate on the "folder" field.
func FolderEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldFolder, v))
}

// FolderNEQ applies the NEQ predicate on the "folder" field.
func FolderNEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNEQ(FieldFolder, v))
}

// FolderIn applies the In predicate on the "folder" field.
func FolderIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldIn(FieldFolder, vs...))
}

// FolderNotIn applies the NotIn predicate on the "folder" field.
func FolderNotIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNotIn(FieldFolder, vs...))
}

// FolderGT applies the GT predicate on the "folder" field.
func FolderGT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGT(FieldFolder, v))
}

// FolderGTE applies the GTE predicate on the "folder" field.
func FolderGTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGTE(FieldFolder, v))
}

// FolderLT applies the LT predicate on the "folder" field.
func FolderLT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLT(FieldFolder, v))
}

// FolderLTE applies the LTE predicate on the "folder" field.
func FolderLTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLTE(FieldFolder, v))
}

// FolderContains applies the Contains predicate on the "folder" field.
func FolderContains(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContains(FieldFolder, v))
}

// FolderHasPrefix applies the HasPrefix predicate on the "folder" field.
func FolderHasPrefix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasPrefix(FieldFolder, v))
}

// FolderHasSuffix applies the HasSuffix predicate on the "folder" field.
func FolderHasSuffix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasSuffix(FieldFolder, v))
}

// FolderIsNil applies the IsNil predicate on the "folder" field.
func FolderIsNil() predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldIsNull(FieldFolder))
}

// FolderNotNil applies the NotNil predicate on the "folder" field.
func FolderNotNil() predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNotNull(FieldFolder))
}

// FolderEqualFold applies the EqualFold predicate on the "folder" field.
func FolderEqualFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEqualFold(FieldFolder, v))
}

// FolderContainsFold applies the ContainsFold predicate on the "folder" field.
func FolderContainsFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContainsFold(FieldFolder, v))
}

// RegionEQ applies the EQ predicate on the "region" field.
func RegionEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldRegion, v))
}

// RegionNEQ applies the NEQ predicate on the "region" field.
func RegionNEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNEQ(FieldRegion, v))
}

// RegionIn applies the In predicate on the "region" field.
func RegionIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldIn(FieldRegion, vs...))
}

// RegionNotIn applies the NotIn predicate on the "region" field.
func RegionNotIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNotIn(FieldRegion, vs...))
}

// RegionGT applies the GT predicate on the "region" field.
func RegionGT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGT(FieldRegion, v))
}

// RegionGTE applies the GTE predicate on the "region" field.
func RegionGTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGTE(FieldRegion, v))
}

// RegionLT applies the LT predicate on the "region" field.
func RegionLT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLT(FieldRegion, v))
}

// RegionLTE applies the LTE predicate on the "region" field.
func RegionLTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLTE(FieldRegion, v))
}

// RegionContains applies the Contains predicate on the "region" field.
func RegionContains(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContains(FieldRegion, v))
}

// RegionHasPrefix applies the HasPrefix predicate on the "region" field.
func RegionHasPrefix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasPrefix(FieldRegion, v))
}

// RegionHasSuffix applies the HasSuffix predicate on the "region" field.
func RegionHasSuffix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasSuffix(FieldRegion, v))
}

// RegionEqualFold applies the EqualFold predicate on the "region" field.
func RegionEqualFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEqualFold(FieldRegion, v))
}

// RegionContainsFold applies the ContainsFold predicate on the "region" field.
func RegionContainsFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContainsFold(FieldRegion, v))
}

// IsDefaultEQ applies the EQ predicate on the "is_default" field.
func IsDefaultEQ(v bool) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldIsDefault, v))
}

// IsDefaultNEQ applies the NEQ predicate on the "is_default" field.
func IsDefaultNEQ(v bool) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNEQ(FieldIsDefault, v))
}

// UseCdnEQ applies the EQ predicate on the "use_cdn" field.
func UseCdnEQ(v bool) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldUseCdn, v))
}

// UseCdnNEQ applies the NEQ predicate on the "use_cdn" field.
func UseCdnNEQ(v bool) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNEQ(FieldUseCdn, v))
}

// CdnURLEQ applies the EQ predicate on the "cdn_url" field.
func CdnURLEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEQ(FieldCdnURL, v))
}

// CdnURLNEQ applies the NEQ predicate on the "cdn_url" field.
func CdnURLNEQ(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNEQ(FieldCdnURL, v))
}

// CdnURLIn applies the In predicate on the "cdn_url" field.
func CdnURLIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldIn(FieldCdnURL, vs...))
}

// CdnURLNotIn applies the NotIn predicate on the "cdn_url" field.
func CdnURLNotIn(vs ...string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNotIn(FieldCdnURL, vs...))
}

// CdnURLGT applies the GT predicate on the "cdn_url" field.
func CdnURLGT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGT(FieldCdnURL, v))
}

// CdnURLGTE applies the GTE predicate on the "cdn_url" field.
func CdnURLGTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldGTE(FieldCdnURL, v))
}

// CdnURLLT applies the LT predicate on the "cdn_url" field.
func CdnURLLT(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLT(FieldCdnURL, v))
}

// CdnURLLTE applies the LTE predicate on the "cdn_url" field.
func CdnURLLTE(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldLTE(FieldCdnURL, v))
}

// CdnURLContains applies the Contains predicate on the "cdn_url" field.
func CdnURLContains(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContains(FieldCdnURL, v))
}

// CdnURLHasPrefix applies the HasPrefix predicate on the "cdn_url" field.
func CdnURLHasPrefix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasPrefix(FieldCdnURL, v))
}

// CdnURLHasSuffix applies the HasSuffix predicate on the "cdn_url" field.
func CdnURLHasSuffix(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldHasSuffix(FieldCdnURL, v))
}

// CdnURLIsNil applies the IsNil predicate on the "cdn_url" field.
func CdnURLIsNil() predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldIsNull(FieldCdnURL))
}

// CdnURLNotNil applies the NotNil predicate on the "cdn_url" field.
func CdnURLNotNil() predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldNotNull(FieldCdnURL))
}

// CdnURLEqualFold applies the EqualFold predicate on the "cdn_url" field.
func CdnURLEqualFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldEqualFold(FieldCdnURL, v))
}

// CdnURLContainsFold applies the ContainsFold predicate on the "cdn_url" field.
func CdnURLContainsFold(v string) predicate.StorageProvider {
	return predicate.StorageProvider(sql.FieldContainsFold(FieldCdnURL, v))
}

// HasCloudfiles applies the HasEdge predicate on the "cloudfiles" edge.
func HasCloudfiles() predicate.StorageProvider {
	return predicate.StorageProvider(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, CloudfilesTable, CloudfilesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCloudfilesWith applies the HasEdge predicate on the "cloudfiles" edge with a given conditions (other predicates).
func HasCloudfilesWith(preds ...predicate.CloudFile) predicate.StorageProvider {
	return predicate.StorageProvider(func(s *sql.Selector) {
		step := newCloudfilesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StorageProvider) predicate.StorageProvider {
	return predicate.StorageProvider(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StorageProvider) predicate.StorageProvider {
	return predicate.StorageProvider(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.StorageProvider) predicate.StorageProvider {
	return predicate.StorageProvider(sql.NotPredicates(p))
}
