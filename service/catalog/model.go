// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AccountsCreateMetastore struct {

	// Wire name: 'metastore_info'
	MetastoreInfo *CreateMetastore
}

func accountsCreateMetastoreToPb(st *AccountsCreateMetastore) (*accountsCreateMetastorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsCreateMetastorePb{}
	metastoreInfoPb, err := createMetastoreToPb(st.MetastoreInfo)
	if err != nil {
		return nil, err
	}
	if metastoreInfoPb != nil {
		pb.MetastoreInfo = metastoreInfoPb
	}

	return pb, nil
}

func (st *AccountsCreateMetastore) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &accountsCreateMetastorePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := accountsCreateMetastoreFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AccountsCreateMetastore) MarshalJSON() ([]byte, error) {
	pb, err := accountsCreateMetastoreToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type accountsCreateMetastorePb struct {
	MetastoreInfo *createMetastorePb `json:"metastore_info,omitempty"`
}

func accountsCreateMetastoreFromPb(pb *accountsCreateMetastorePb) (*AccountsCreateMetastore, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsCreateMetastore{}
	metastoreInfoField, err := createMetastoreFromPb(pb.MetastoreInfo)
	if err != nil {
		return nil, err
	}
	if metastoreInfoField != nil {
		st.MetastoreInfo = metastoreInfoField
	}

	return st, nil
}

type AccountsCreateMetastoreAssignment struct {

	// Wire name: 'metastore_assignment'
	MetastoreAssignment *CreateMetastoreAssignment
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func accountsCreateMetastoreAssignmentToPb(st *AccountsCreateMetastoreAssignment) (*accountsCreateMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsCreateMetastoreAssignmentPb{}
	metastoreAssignmentPb, err := createMetastoreAssignmentToPb(st.MetastoreAssignment)
	if err != nil {
		return nil, err
	}
	if metastoreAssignmentPb != nil {
		pb.MetastoreAssignment = metastoreAssignmentPb
	}

	pb.MetastoreId = st.MetastoreId

	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func (st *AccountsCreateMetastoreAssignment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &accountsCreateMetastoreAssignmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := accountsCreateMetastoreAssignmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AccountsCreateMetastoreAssignment) MarshalJSON() ([]byte, error) {
	pb, err := accountsCreateMetastoreAssignmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type accountsCreateMetastoreAssignmentPb struct {
	MetastoreAssignment *createMetastoreAssignmentPb `json:"metastore_assignment,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func accountsCreateMetastoreAssignmentFromPb(pb *accountsCreateMetastoreAssignmentPb) (*AccountsCreateMetastoreAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsCreateMetastoreAssignment{}
	metastoreAssignmentField, err := createMetastoreAssignmentFromPb(pb.MetastoreAssignment)
	if err != nil {
		return nil, err
	}
	if metastoreAssignmentField != nil {
		st.MetastoreAssignment = metastoreAssignmentField
	}
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

type AccountsCreateStorageCredential struct {

	// Wire name: 'credential_info'
	CredentialInfo *CreateStorageCredential
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
}

func accountsCreateStorageCredentialToPb(st *AccountsCreateStorageCredential) (*accountsCreateStorageCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsCreateStorageCredentialPb{}
	credentialInfoPb, err := createStorageCredentialToPb(st.CredentialInfo)
	if err != nil {
		return nil, err
	}
	if credentialInfoPb != nil {
		pb.CredentialInfo = credentialInfoPb
	}

	pb.MetastoreId = st.MetastoreId

	return pb, nil
}

func (st *AccountsCreateStorageCredential) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &accountsCreateStorageCredentialPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := accountsCreateStorageCredentialFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AccountsCreateStorageCredential) MarshalJSON() ([]byte, error) {
	pb, err := accountsCreateStorageCredentialToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type accountsCreateStorageCredentialPb struct {
	CredentialInfo *createStorageCredentialPb `json:"credential_info,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
}

func accountsCreateStorageCredentialFromPb(pb *accountsCreateStorageCredentialPb) (*AccountsCreateStorageCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsCreateStorageCredential{}
	credentialInfoField, err := createStorageCredentialFromPb(pb.CredentialInfo)
	if err != nil {
		return nil, err
	}
	if credentialInfoField != nil {
		st.CredentialInfo = credentialInfoField
	}
	st.MetastoreId = pb.MetastoreId

	return st, nil
}

type AccountsMetastoreAssignment struct {

	// Wire name: 'metastore_assignment'
	MetastoreAssignment *MetastoreAssignment
}

func accountsMetastoreAssignmentToPb(st *AccountsMetastoreAssignment) (*accountsMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsMetastoreAssignmentPb{}
	metastoreAssignmentPb, err := metastoreAssignmentToPb(st.MetastoreAssignment)
	if err != nil {
		return nil, err
	}
	if metastoreAssignmentPb != nil {
		pb.MetastoreAssignment = metastoreAssignmentPb
	}

	return pb, nil
}

func (st *AccountsMetastoreAssignment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &accountsMetastoreAssignmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := accountsMetastoreAssignmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AccountsMetastoreAssignment) MarshalJSON() ([]byte, error) {
	pb, err := accountsMetastoreAssignmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type accountsMetastoreAssignmentPb struct {
	MetastoreAssignment *metastoreAssignmentPb `json:"metastore_assignment,omitempty"`
}

func accountsMetastoreAssignmentFromPb(pb *accountsMetastoreAssignmentPb) (*AccountsMetastoreAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsMetastoreAssignment{}
	metastoreAssignmentField, err := metastoreAssignmentFromPb(pb.MetastoreAssignment)
	if err != nil {
		return nil, err
	}
	if metastoreAssignmentField != nil {
		st.MetastoreAssignment = metastoreAssignmentField
	}

	return st, nil
}

type AccountsMetastoreInfo struct {

	// Wire name: 'metastore_info'
	MetastoreInfo *MetastoreInfo
}

func accountsMetastoreInfoToPb(st *AccountsMetastoreInfo) (*accountsMetastoreInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsMetastoreInfoPb{}
	metastoreInfoPb, err := metastoreInfoToPb(st.MetastoreInfo)
	if err != nil {
		return nil, err
	}
	if metastoreInfoPb != nil {
		pb.MetastoreInfo = metastoreInfoPb
	}

	return pb, nil
}

func (st *AccountsMetastoreInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &accountsMetastoreInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := accountsMetastoreInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AccountsMetastoreInfo) MarshalJSON() ([]byte, error) {
	pb, err := accountsMetastoreInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type accountsMetastoreInfoPb struct {
	MetastoreInfo *metastoreInfoPb `json:"metastore_info,omitempty"`
}

func accountsMetastoreInfoFromPb(pb *accountsMetastoreInfoPb) (*AccountsMetastoreInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsMetastoreInfo{}
	metastoreInfoField, err := metastoreInfoFromPb(pb.MetastoreInfo)
	if err != nil {
		return nil, err
	}
	if metastoreInfoField != nil {
		st.MetastoreInfo = metastoreInfoField
	}

	return st, nil
}

type AccountsStorageCredentialInfo struct {

	// Wire name: 'credential_info'
	CredentialInfo *StorageCredentialInfo
}

func accountsStorageCredentialInfoToPb(st *AccountsStorageCredentialInfo) (*accountsStorageCredentialInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsStorageCredentialInfoPb{}
	credentialInfoPb, err := storageCredentialInfoToPb(st.CredentialInfo)
	if err != nil {
		return nil, err
	}
	if credentialInfoPb != nil {
		pb.CredentialInfo = credentialInfoPb
	}

	return pb, nil
}

func (st *AccountsStorageCredentialInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &accountsStorageCredentialInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := accountsStorageCredentialInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AccountsStorageCredentialInfo) MarshalJSON() ([]byte, error) {
	pb, err := accountsStorageCredentialInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type accountsStorageCredentialInfoPb struct {
	CredentialInfo *storageCredentialInfoPb `json:"credential_info,omitempty"`
}

func accountsStorageCredentialInfoFromPb(pb *accountsStorageCredentialInfoPb) (*AccountsStorageCredentialInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsStorageCredentialInfo{}
	credentialInfoField, err := storageCredentialInfoFromPb(pb.CredentialInfo)
	if err != nil {
		return nil, err
	}
	if credentialInfoField != nil {
		st.CredentialInfo = credentialInfoField
	}

	return st, nil
}

type AccountsUpdateMetastore struct {
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`

	// Wire name: 'metastore_info'
	MetastoreInfo *UpdateMetastore
}

func accountsUpdateMetastoreToPb(st *AccountsUpdateMetastore) (*accountsUpdateMetastorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsUpdateMetastorePb{}
	pb.MetastoreId = st.MetastoreId

	metastoreInfoPb, err := updateMetastoreToPb(st.MetastoreInfo)
	if err != nil {
		return nil, err
	}
	if metastoreInfoPb != nil {
		pb.MetastoreInfo = metastoreInfoPb
	}

	return pb, nil
}

func (st *AccountsUpdateMetastore) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &accountsUpdateMetastorePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := accountsUpdateMetastoreFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AccountsUpdateMetastore) MarshalJSON() ([]byte, error) {
	pb, err := accountsUpdateMetastoreToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type accountsUpdateMetastorePb struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`

	MetastoreInfo *updateMetastorePb `json:"metastore_info,omitempty"`
}

func accountsUpdateMetastoreFromPb(pb *accountsUpdateMetastorePb) (*AccountsUpdateMetastore, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsUpdateMetastore{}
	st.MetastoreId = pb.MetastoreId
	metastoreInfoField, err := updateMetastoreFromPb(pb.MetastoreInfo)
	if err != nil {
		return nil, err
	}
	if metastoreInfoField != nil {
		st.MetastoreInfo = metastoreInfoField
	}

	return st, nil
}

type AccountsUpdateMetastoreAssignment struct {

	// Wire name: 'metastore_assignment'
	MetastoreAssignment *UpdateMetastoreAssignment
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func accountsUpdateMetastoreAssignmentToPb(st *AccountsUpdateMetastoreAssignment) (*accountsUpdateMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsUpdateMetastoreAssignmentPb{}
	metastoreAssignmentPb, err := updateMetastoreAssignmentToPb(st.MetastoreAssignment)
	if err != nil {
		return nil, err
	}
	if metastoreAssignmentPb != nil {
		pb.MetastoreAssignment = metastoreAssignmentPb
	}

	pb.MetastoreId = st.MetastoreId

	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func (st *AccountsUpdateMetastoreAssignment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &accountsUpdateMetastoreAssignmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := accountsUpdateMetastoreAssignmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AccountsUpdateMetastoreAssignment) MarshalJSON() ([]byte, error) {
	pb, err := accountsUpdateMetastoreAssignmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type accountsUpdateMetastoreAssignmentPb struct {
	MetastoreAssignment *updateMetastoreAssignmentPb `json:"metastore_assignment,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func accountsUpdateMetastoreAssignmentFromPb(pb *accountsUpdateMetastoreAssignmentPb) (*AccountsUpdateMetastoreAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsUpdateMetastoreAssignment{}
	metastoreAssignmentField, err := updateMetastoreAssignmentFromPb(pb.MetastoreAssignment)
	if err != nil {
		return nil, err
	}
	if metastoreAssignmentField != nil {
		st.MetastoreAssignment = metastoreAssignmentField
	}
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

type AccountsUpdateStorageCredential struct {

	// Wire name: 'credential_info'
	CredentialInfo *UpdateStorageCredential
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Name of the storage credential.
	// Wire name: 'storage_credential_name'
	StorageCredentialName string `tf:"-"`
}

func accountsUpdateStorageCredentialToPb(st *AccountsUpdateStorageCredential) (*accountsUpdateStorageCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsUpdateStorageCredentialPb{}
	credentialInfoPb, err := updateStorageCredentialToPb(st.CredentialInfo)
	if err != nil {
		return nil, err
	}
	if credentialInfoPb != nil {
		pb.CredentialInfo = credentialInfoPb
	}

	pb.MetastoreId = st.MetastoreId

	pb.StorageCredentialName = st.StorageCredentialName

	return pb, nil
}

func (st *AccountsUpdateStorageCredential) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &accountsUpdateStorageCredentialPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := accountsUpdateStorageCredentialFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AccountsUpdateStorageCredential) MarshalJSON() ([]byte, error) {
	pb, err := accountsUpdateStorageCredentialToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type accountsUpdateStorageCredentialPb struct {
	CredentialInfo *updateStorageCredentialPb `json:"credential_info,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Name of the storage credential.
	StorageCredentialName string `json:"-" url:"-"`
}

func accountsUpdateStorageCredentialFromPb(pb *accountsUpdateStorageCredentialPb) (*AccountsUpdateStorageCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsUpdateStorageCredential{}
	credentialInfoField, err := updateStorageCredentialFromPb(pb.CredentialInfo)
	if err != nil {
		return nil, err
	}
	if credentialInfoField != nil {
		st.CredentialInfo = credentialInfoField
	}
	st.MetastoreId = pb.MetastoreId
	st.StorageCredentialName = pb.StorageCredentialName

	return st, nil
}

type ArtifactAllowlistInfo struct {
	// A list of allowed artifact match patterns.
	// Wire name: 'artifact_matchers'
	ArtifactMatchers []ArtifactMatcher
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of the user who set the artifact allowlist.
	// Wire name: 'created_by'
	CreatedBy string
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string

	ForceSendFields []string `tf:"-"`
}

func artifactAllowlistInfoToPb(st *ArtifactAllowlistInfo) (*artifactAllowlistInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &artifactAllowlistInfoPb{}

	var artifactMatchersPb []artifactMatcherPb
	for _, item := range st.ArtifactMatchers {
		itemPb, err := artifactMatcherToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			artifactMatchersPb = append(artifactMatchersPb, *itemPb)
		}
	}
	pb.ArtifactMatchers = artifactMatchersPb

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	pb.MetastoreId = st.MetastoreId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ArtifactAllowlistInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &artifactAllowlistInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := artifactAllowlistInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ArtifactAllowlistInfo) MarshalJSON() ([]byte, error) {
	pb, err := artifactAllowlistInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type artifactAllowlistInfoPb struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers []artifactMatcherPb `json:"artifact_matchers,omitempty"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of the user who set the artifact allowlist.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func artifactAllowlistInfoFromPb(pb *artifactAllowlistInfoPb) (*ArtifactAllowlistInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ArtifactAllowlistInfo{}

	var artifactMatchersField []ArtifactMatcher
	for _, itemPb := range pb.ArtifactMatchers {
		item, err := artifactMatcherFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			artifactMatchersField = append(artifactMatchersField, *item)
		}
	}
	st.ArtifactMatchers = artifactMatchersField
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.MetastoreId = pb.MetastoreId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *artifactAllowlistInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st artifactAllowlistInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ArtifactMatcher struct {
	// The artifact path or maven coordinate
	// Wire name: 'artifact'
	Artifact string
	// The pattern matching type of the artifact
	// Wire name: 'match_type'
	MatchType MatchType
}

func artifactMatcherToPb(st *ArtifactMatcher) (*artifactMatcherPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &artifactMatcherPb{}
	pb.Artifact = st.Artifact

	pb.MatchType = st.MatchType

	return pb, nil
}

func (st *ArtifactMatcher) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &artifactMatcherPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := artifactMatcherFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ArtifactMatcher) MarshalJSON() ([]byte, error) {
	pb, err := artifactMatcherToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type artifactMatcherPb struct {
	// The artifact path or maven coordinate
	Artifact string `json:"artifact"`
	// The pattern matching type of the artifact
	MatchType MatchType `json:"match_type"`
}

func artifactMatcherFromPb(pb *artifactMatcherPb) (*ArtifactMatcher, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ArtifactMatcher{}
	st.Artifact = pb.Artifact
	st.MatchType = pb.MatchType

	return st, nil
}

// The artifact type
type ArtifactType string
type artifactTypePb string

const ArtifactTypeInitScript ArtifactType = `INIT_SCRIPT`

const ArtifactTypeLibraryJar ArtifactType = `LIBRARY_JAR`

const ArtifactTypeLibraryMaven ArtifactType = `LIBRARY_MAVEN`

// String representation for [fmt.Print]
func (f *ArtifactType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ArtifactType) Set(v string) error {
	switch v {
	case `INIT_SCRIPT`, `LIBRARY_JAR`, `LIBRARY_MAVEN`:
		*f = ArtifactType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INIT_SCRIPT", "LIBRARY_JAR", "LIBRARY_MAVEN"`, v)
	}
}

// Type always returns ArtifactType to satisfy [pflag.Value] interface
func (f *ArtifactType) Type() string {
	return "ArtifactType"
}

func artifactTypeToPb(st *ArtifactType) (*artifactTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := artifactTypePb(*st)
	return &pb, nil
}

func artifactTypeFromPb(pb *artifactTypePb) (*ArtifactType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ArtifactType(*pb)
	return &st, nil
}

type AssignResponse struct {
}

func assignResponseToPb(st *AssignResponse) (*assignResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &assignResponsePb{}

	return pb, nil
}

func (st *AssignResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &assignResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := assignResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AssignResponse) MarshalJSON() ([]byte, error) {
	pb, err := assignResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type assignResponsePb struct {
}

func assignResponseFromPb(pb *assignResponsePb) (*AssignResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AssignResponse{}

	return st, nil
}

// AWS temporary credentials for API authentication. Read more at
// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
type AwsCredentials struct {
	// The access key ID that identifies the temporary credentials.
	// Wire name: 'access_key_id'
	AccessKeyId string
	// The Amazon Resource Name (ARN) of the S3 access point for temporary
	// credentials related the external location.
	// Wire name: 'access_point'
	AccessPoint string
	// The secret access key that can be used to sign AWS API requests.
	// Wire name: 'secret_access_key'
	SecretAccessKey string
	// The token that users must pass to AWS API to use the temporary
	// credentials.
	// Wire name: 'session_token'
	SessionToken string

	ForceSendFields []string `tf:"-"`
}

func awsCredentialsToPb(st *AwsCredentials) (*awsCredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsCredentialsPb{}
	pb.AccessKeyId = st.AccessKeyId

	pb.AccessPoint = st.AccessPoint

	pb.SecretAccessKey = st.SecretAccessKey

	pb.SessionToken = st.SessionToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *AwsCredentials) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &awsCredentialsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := awsCredentialsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AwsCredentials) MarshalJSON() ([]byte, error) {
	pb, err := awsCredentialsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type awsCredentialsPb struct {
	// The access key ID that identifies the temporary credentials.
	AccessKeyId string `json:"access_key_id,omitempty"`
	// The Amazon Resource Name (ARN) of the S3 access point for temporary
	// credentials related the external location.
	AccessPoint string `json:"access_point,omitempty"`
	// The secret access key that can be used to sign AWS API requests.
	SecretAccessKey string `json:"secret_access_key,omitempty"`
	// The token that users must pass to AWS API to use the temporary
	// credentials.
	SessionToken string `json:"session_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func awsCredentialsFromPb(pb *awsCredentialsPb) (*AwsCredentials, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsCredentials{}
	st.AccessKeyId = pb.AccessKeyId
	st.AccessPoint = pb.AccessPoint
	st.SecretAccessKey = pb.SecretAccessKey
	st.SessionToken = pb.SessionToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *awsCredentialsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st awsCredentialsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The AWS IAM role configuration
type AwsIamRole struct {
	// The external ID used in role assumption to prevent the confused deputy
	// problem.
	// Wire name: 'external_id'
	ExternalId string
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	// Wire name: 'role_arn'
	RoleArn string
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	// Wire name: 'unity_catalog_iam_arn'
	UnityCatalogIamArn string

	ForceSendFields []string `tf:"-"`
}

func awsIamRoleToPb(st *AwsIamRole) (*awsIamRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsIamRolePb{}
	pb.ExternalId = st.ExternalId

	pb.RoleArn = st.RoleArn

	pb.UnityCatalogIamArn = st.UnityCatalogIamArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *AwsIamRole) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &awsIamRolePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := awsIamRoleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AwsIamRole) MarshalJSON() ([]byte, error) {
	pb, err := awsIamRoleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type awsIamRolePb struct {
	// The external ID used in role assumption to prevent the confused deputy
	// problem.
	ExternalId string `json:"external_id,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	RoleArn string `json:"role_arn,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func awsIamRoleFromPb(pb *awsIamRolePb) (*AwsIamRole, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsIamRole{}
	st.ExternalId = pb.ExternalId
	st.RoleArn = pb.RoleArn
	st.UnityCatalogIamArn = pb.UnityCatalogIamArn

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *awsIamRolePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st awsIamRolePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AwsIamRoleRequest struct {
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	// Wire name: 'role_arn'
	RoleArn string
}

func awsIamRoleRequestToPb(st *AwsIamRoleRequest) (*awsIamRoleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsIamRoleRequestPb{}
	pb.RoleArn = st.RoleArn

	return pb, nil
}

func (st *AwsIamRoleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &awsIamRoleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := awsIamRoleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AwsIamRoleRequest) MarshalJSON() ([]byte, error) {
	pb, err := awsIamRoleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type awsIamRoleRequestPb struct {
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn string `json:"role_arn"`
}

func awsIamRoleRequestFromPb(pb *awsIamRoleRequestPb) (*AwsIamRoleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsIamRoleRequest{}
	st.RoleArn = pb.RoleArn

	return st, nil
}

type AwsIamRoleResponse struct {
	// The external ID used in role assumption to prevent confused deputy
	// problem..
	// Wire name: 'external_id'
	ExternalId string
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	// Wire name: 'role_arn'
	RoleArn string
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	// Wire name: 'unity_catalog_iam_arn'
	UnityCatalogIamArn string

	ForceSendFields []string `tf:"-"`
}

func awsIamRoleResponseToPb(st *AwsIamRoleResponse) (*awsIamRoleResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsIamRoleResponsePb{}
	pb.ExternalId = st.ExternalId

	pb.RoleArn = st.RoleArn

	pb.UnityCatalogIamArn = st.UnityCatalogIamArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *AwsIamRoleResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &awsIamRoleResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := awsIamRoleResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AwsIamRoleResponse) MarshalJSON() ([]byte, error) {
	pb, err := awsIamRoleResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type awsIamRoleResponsePb struct {
	// The external ID used in role assumption to prevent confused deputy
	// problem..
	ExternalId string `json:"external_id,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn string `json:"role_arn"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func awsIamRoleResponseFromPb(pb *awsIamRoleResponsePb) (*AwsIamRoleResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsIamRoleResponse{}
	st.ExternalId = pb.ExternalId
	st.RoleArn = pb.RoleArn
	st.UnityCatalogIamArn = pb.UnityCatalogIamArn

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *awsIamRoleResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st awsIamRoleResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Azure Active Directory token, essentially the Oauth token for Azure Service
// Principal or Managed Identity. Read more at
// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
type AzureActiveDirectoryToken struct {
	// Opaque token that contains claims that you can use in Azure Active
	// Directory to access cloud services.
	// Wire name: 'aad_token'
	AadToken string

	ForceSendFields []string `tf:"-"`
}

func azureActiveDirectoryTokenToPb(st *AzureActiveDirectoryToken) (*azureActiveDirectoryTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureActiveDirectoryTokenPb{}
	pb.AadToken = st.AadToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *AzureActiveDirectoryToken) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &azureActiveDirectoryTokenPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := azureActiveDirectoryTokenFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AzureActiveDirectoryToken) MarshalJSON() ([]byte, error) {
	pb, err := azureActiveDirectoryTokenToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type azureActiveDirectoryTokenPb struct {
	// Opaque token that contains claims that you can use in Azure Active
	// Directory to access cloud services.
	AadToken string `json:"aad_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func azureActiveDirectoryTokenFromPb(pb *azureActiveDirectoryTokenPb) (*AzureActiveDirectoryToken, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureActiveDirectoryToken{}
	st.AadToken = pb.AadToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *azureActiveDirectoryTokenPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st azureActiveDirectoryTokenPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The Azure managed identity configuration.
type AzureManagedIdentity struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}`.
	// Wire name: 'access_connector_id'
	AccessConnectorId string
	// The Databricks internal ID that represents this managed identity. This
	// field is only used to persist the credential_id once it is fetched from
	// the credentials manager - as we only use the protobuf serializer to store
	// credentials, this ID gets persisted to the database. .
	// Wire name: 'credential_id'
	CredentialId string
	// The Azure resource ID of the managed identity. Use the format,
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}`
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// using the system-assigned identity.
	// Wire name: 'managed_identity_id'
	ManagedIdentityId string

	ForceSendFields []string `tf:"-"`
}

func azureManagedIdentityToPb(st *AzureManagedIdentity) (*azureManagedIdentityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureManagedIdentityPb{}
	pb.AccessConnectorId = st.AccessConnectorId

	pb.CredentialId = st.CredentialId

	pb.ManagedIdentityId = st.ManagedIdentityId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *AzureManagedIdentity) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &azureManagedIdentityPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := azureManagedIdentityFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AzureManagedIdentity) MarshalJSON() ([]byte, error) {
	pb, err := azureManagedIdentityToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type azureManagedIdentityPb struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}`.
	AccessConnectorId string `json:"access_connector_id"`
	// The Databricks internal ID that represents this managed identity. This
	// field is only used to persist the credential_id once it is fetched from
	// the credentials manager - as we only use the protobuf serializer to store
	// credentials, this ID gets persisted to the database. .
	CredentialId string `json:"credential_id,omitempty"`
	// The Azure resource ID of the managed identity. Use the format,
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}`
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// using the system-assigned identity.
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func azureManagedIdentityFromPb(pb *azureManagedIdentityPb) (*AzureManagedIdentity, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureManagedIdentity{}
	st.AccessConnectorId = pb.AccessConnectorId
	st.CredentialId = pb.CredentialId
	st.ManagedIdentityId = pb.ManagedIdentityId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *azureManagedIdentityPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st azureManagedIdentityPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AzureManagedIdentityRequest struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}.
	// Wire name: 'access_connector_id'
	AccessConnectorId string
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	// Wire name: 'managed_identity_id'
	ManagedIdentityId string

	ForceSendFields []string `tf:"-"`
}

func azureManagedIdentityRequestToPb(st *AzureManagedIdentityRequest) (*azureManagedIdentityRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureManagedIdentityRequestPb{}
	pb.AccessConnectorId = st.AccessConnectorId

	pb.ManagedIdentityId = st.ManagedIdentityId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *AzureManagedIdentityRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &azureManagedIdentityRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := azureManagedIdentityRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AzureManagedIdentityRequest) MarshalJSON() ([]byte, error) {
	pb, err := azureManagedIdentityRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type azureManagedIdentityRequestPb struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}.
	AccessConnectorId string `json:"access_connector_id"`
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func azureManagedIdentityRequestFromPb(pb *azureManagedIdentityRequestPb) (*AzureManagedIdentityRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureManagedIdentityRequest{}
	st.AccessConnectorId = pb.AccessConnectorId
	st.ManagedIdentityId = pb.ManagedIdentityId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *azureManagedIdentityRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st azureManagedIdentityRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AzureManagedIdentityResponse struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}.
	// Wire name: 'access_connector_id'
	AccessConnectorId string
	// The Databricks internal ID that represents this managed identity.
	// Wire name: 'credential_id'
	CredentialId string
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	// Wire name: 'managed_identity_id'
	ManagedIdentityId string

	ForceSendFields []string `tf:"-"`
}

func azureManagedIdentityResponseToPb(st *AzureManagedIdentityResponse) (*azureManagedIdentityResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureManagedIdentityResponsePb{}
	pb.AccessConnectorId = st.AccessConnectorId

	pb.CredentialId = st.CredentialId

	pb.ManagedIdentityId = st.ManagedIdentityId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *AzureManagedIdentityResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &azureManagedIdentityResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := azureManagedIdentityResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AzureManagedIdentityResponse) MarshalJSON() ([]byte, error) {
	pb, err := azureManagedIdentityResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type azureManagedIdentityResponsePb struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}.
	AccessConnectorId string `json:"access_connector_id"`
	// The Databricks internal ID that represents this managed identity.
	CredentialId string `json:"credential_id,omitempty"`
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func azureManagedIdentityResponseFromPb(pb *azureManagedIdentityResponsePb) (*AzureManagedIdentityResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureManagedIdentityResponse{}
	st.AccessConnectorId = pb.AccessConnectorId
	st.CredentialId = pb.CredentialId
	st.ManagedIdentityId = pb.ManagedIdentityId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *azureManagedIdentityResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st azureManagedIdentityResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The Azure service principal configuration. Only applicable when purpose is
// **STORAGE**.
type AzureServicePrincipal struct {
	// The application ID of the application registration within the referenced
	// AAD tenant.
	// Wire name: 'application_id'
	ApplicationId string
	// The client secret generated for the above app ID in AAD.
	// Wire name: 'client_secret'
	ClientSecret string
	// The directory ID corresponding to the Azure Active Directory (AAD) tenant
	// of the application.
	// Wire name: 'directory_id'
	DirectoryId string
}

func azureServicePrincipalToPb(st *AzureServicePrincipal) (*azureServicePrincipalPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureServicePrincipalPb{}
	pb.ApplicationId = st.ApplicationId

	pb.ClientSecret = st.ClientSecret

	pb.DirectoryId = st.DirectoryId

	return pb, nil
}

func (st *AzureServicePrincipal) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &azureServicePrincipalPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := azureServicePrincipalFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AzureServicePrincipal) MarshalJSON() ([]byte, error) {
	pb, err := azureServicePrincipalToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type azureServicePrincipalPb struct {
	// The application ID of the application registration within the referenced
	// AAD tenant.
	ApplicationId string `json:"application_id"`
	// The client secret generated for the above app ID in AAD.
	ClientSecret string `json:"client_secret"`
	// The directory ID corresponding to the Azure Active Directory (AAD) tenant
	// of the application.
	DirectoryId string `json:"directory_id"`
}

func azureServicePrincipalFromPb(pb *azureServicePrincipalPb) (*AzureServicePrincipal, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureServicePrincipal{}
	st.ApplicationId = pb.ApplicationId
	st.ClientSecret = pb.ClientSecret
	st.DirectoryId = pb.DirectoryId

	return st, nil
}

// Azure temporary credentials for API authentication. Read more at
// https://docs.microsoft.com/en-us/rest/api/storageservices/create-user-delegation-sas
type AzureUserDelegationSas struct {
	// The signed URI (SAS Token) used to access blob services for a given path
	// Wire name: 'sas_token'
	SasToken string

	ForceSendFields []string `tf:"-"`
}

func azureUserDelegationSasToPb(st *AzureUserDelegationSas) (*azureUserDelegationSasPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureUserDelegationSasPb{}
	pb.SasToken = st.SasToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *AzureUserDelegationSas) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &azureUserDelegationSasPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := azureUserDelegationSasFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AzureUserDelegationSas) MarshalJSON() ([]byte, error) {
	pb, err := azureUserDelegationSasToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type azureUserDelegationSasPb struct {
	// The signed URI (SAS Token) used to access blob services for a given path
	SasToken string `json:"sas_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func azureUserDelegationSasFromPb(pb *azureUserDelegationSasPb) (*AzureUserDelegationSas, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureUserDelegationSas{}
	st.SasToken = pb.SasToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *azureUserDelegationSasPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st azureUserDelegationSasPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Cancel refresh
type CancelRefreshRequest struct {
	// ID of the refresh.
	// Wire name: 'refresh_id'
	RefreshId string `tf:"-"`
	// Full name of the table.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func cancelRefreshRequestToPb(st *CancelRefreshRequest) (*cancelRefreshRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelRefreshRequestPb{}
	pb.RefreshId = st.RefreshId

	pb.TableName = st.TableName

	return pb, nil
}

func (st *CancelRefreshRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelRefreshRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelRefreshRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelRefreshRequest) MarshalJSON() ([]byte, error) {
	pb, err := cancelRefreshRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cancelRefreshRequestPb struct {
	// ID of the refresh.
	RefreshId string `json:"-" url:"-"`
	// Full name of the table.
	TableName string `json:"-" url:"-"`
}

func cancelRefreshRequestFromPb(pb *cancelRefreshRequestPb) (*CancelRefreshRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelRefreshRequest{}
	st.RefreshId = pb.RefreshId
	st.TableName = pb.TableName

	return st, nil
}

type CancelRefreshResponse struct {
}

func cancelRefreshResponseToPb(st *CancelRefreshResponse) (*cancelRefreshResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelRefreshResponsePb{}

	return pb, nil
}

func (st *CancelRefreshResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelRefreshResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelRefreshResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelRefreshResponse) MarshalJSON() ([]byte, error) {
	pb, err := cancelRefreshResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cancelRefreshResponsePb struct {
}

func cancelRefreshResponseFromPb(pb *cancelRefreshResponsePb) (*CancelRefreshResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelRefreshResponse{}

	return st, nil
}

type CatalogInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool
	// The type of the catalog.
	// Wire name: 'catalog_type'
	CatalogType CatalogType
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// The name of the connection to an external data source.
	// Wire name: 'connection_name'
	ConnectionName string
	// Time at which this catalog was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of catalog creator.
	// Wire name: 'created_by'
	CreatedBy string

	// Wire name: 'effective_predictive_optimization_flag'
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'enable_predictive_optimization'
	EnablePredictiveOptimization EnablePredictiveOptimization
	// The full name of the catalog. Corresponds with the name field.
	// Wire name: 'full_name'
	FullName string
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode CatalogIsolationMode
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string
	// Name of catalog.
	// Wire name: 'name'
	Name string
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string
	// Username of current owner of catalog.
	// Wire name: 'owner'
	Owner string
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	// Wire name: 'provider_name'
	ProviderName string
	// Status of an asynchronously provisioned resource.
	// Wire name: 'provisioning_info'
	ProvisioningInfo *ProvisioningInfo

	// Wire name: 'securable_type'
	SecurableType string
	// The name of the share under the share provider.
	// Wire name: 'share_name'
	ShareName string
	// Storage Location URL (full path) for managed tables within catalog.
	// Wire name: 'storage_location'
	StorageLocation string
	// Storage root URL for managed tables within catalog.
	// Wire name: 'storage_root'
	StorageRoot string
	// Time at which this catalog was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// Username of user who last modified catalog.
	// Wire name: 'updated_by'
	UpdatedBy string

	ForceSendFields []string `tf:"-"`
}

func catalogInfoToPb(st *CatalogInfo) (*catalogInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogInfoPb{}
	pb.BrowseOnly = st.BrowseOnly

	pb.CatalogType = st.CatalogType

	pb.Comment = st.Comment

	pb.ConnectionName = st.ConnectionName

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	effectivePredictiveOptimizationFlagPb, err := effectivePredictiveOptimizationFlagToPb(st.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagPb != nil {
		pb.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagPb
	}

	pb.EnablePredictiveOptimization = st.EnablePredictiveOptimization

	pb.FullName = st.FullName

	pb.IsolationMode = st.IsolationMode

	pb.MetastoreId = st.MetastoreId

	pb.Name = st.Name

	pb.Options = st.Options

	pb.Owner = st.Owner

	pb.Properties = st.Properties

	pb.ProviderName = st.ProviderName

	provisioningInfoPb, err := provisioningInfoToPb(st.ProvisioningInfo)
	if err != nil {
		return nil, err
	}
	if provisioningInfoPb != nil {
		pb.ProvisioningInfo = provisioningInfoPb
	}

	pb.SecurableType = st.SecurableType

	pb.ShareName = st.ShareName

	pb.StorageLocation = st.StorageLocation

	pb.StorageRoot = st.StorageRoot

	pb.UpdatedAt = st.UpdatedAt

	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CatalogInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &catalogInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := catalogInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CatalogInfo) MarshalJSON() ([]byte, error) {
	pb, err := catalogInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type catalogInfoPb struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The type of the catalog.
	CatalogType CatalogType `json:"catalog_type,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// The name of the connection to an external data source.
	ConnectionName string `json:"connection_name,omitempty"`
	// Time at which this catalog was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of catalog creator.
	CreatedBy string `json:"created_by,omitempty"`

	EffectivePredictiveOptimizationFlag *effectivePredictiveOptimizationFlagPb `json:"effective_predictive_optimization_flag,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// The full name of the catalog. Corresponds with the name field.
	FullName string `json:"full_name,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode CatalogIsolationMode `json:"isolation_mode,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of catalog.
	Name string `json:"name,omitempty"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options,omitempty"`
	// Username of current owner of catalog.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName string `json:"provider_name,omitempty"`
	// Status of an asynchronously provisioned resource.
	ProvisioningInfo *provisioningInfoPb `json:"provisioning_info,omitempty"`

	SecurableType string `json:"securable_type,omitempty"`
	// The name of the share under the share provider.
	ShareName string `json:"share_name,omitempty"`
	// Storage Location URL (full path) for managed tables within catalog.
	StorageLocation string `json:"storage_location,omitempty"`
	// Storage root URL for managed tables within catalog.
	StorageRoot string `json:"storage_root,omitempty"`
	// Time at which this catalog was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified catalog.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func catalogInfoFromPb(pb *catalogInfoPb) (*CatalogInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CatalogInfo{}
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogType = pb.CatalogType
	st.Comment = pb.Comment
	st.ConnectionName = pb.ConnectionName
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	effectivePredictiveOptimizationFlagField, err := effectivePredictiveOptimizationFlagFromPb(pb.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagField != nil {
		st.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagField
	}
	st.EnablePredictiveOptimization = pb.EnablePredictiveOptimization
	st.FullName = pb.FullName
	st.IsolationMode = pb.IsolationMode
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Options = pb.Options
	st.Owner = pb.Owner
	st.Properties = pb.Properties
	st.ProviderName = pb.ProviderName
	provisioningInfoField, err := provisioningInfoFromPb(pb.ProvisioningInfo)
	if err != nil {
		return nil, err
	}
	if provisioningInfoField != nil {
		st.ProvisioningInfo = provisioningInfoField
	}
	st.SecurableType = pb.SecurableType
	st.ShareName = pb.ShareName
	st.StorageLocation = pb.StorageLocation
	st.StorageRoot = pb.StorageRoot
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *catalogInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st catalogInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Whether the current securable is accessible from all workspaces or a specific
// set of workspaces.
type CatalogIsolationMode string
type catalogIsolationModePb string

const CatalogIsolationModeIsolated CatalogIsolationMode = `ISOLATED`

const CatalogIsolationModeOpen CatalogIsolationMode = `OPEN`

// String representation for [fmt.Print]
func (f *CatalogIsolationMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CatalogIsolationMode) Set(v string) error {
	switch v {
	case `ISOLATED`, `OPEN`:
		*f = CatalogIsolationMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ISOLATED", "OPEN"`, v)
	}
}

// Type always returns CatalogIsolationMode to satisfy [pflag.Value] interface
func (f *CatalogIsolationMode) Type() string {
	return "CatalogIsolationMode"
}

func catalogIsolationModeToPb(st *CatalogIsolationMode) (*catalogIsolationModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogIsolationModePb(*st)
	return &pb, nil
}

func catalogIsolationModeFromPb(pb *catalogIsolationModePb) (*CatalogIsolationMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := CatalogIsolationMode(*pb)
	return &st, nil
}

// The type of the catalog.
type CatalogType string
type catalogTypePb string

const CatalogTypeDeltasharingCatalog CatalogType = `DELTASHARING_CATALOG`

const CatalogTypeForeignCatalog CatalogType = `FOREIGN_CATALOG`

const CatalogTypeManagedCatalog CatalogType = `MANAGED_CATALOG`

const CatalogTypeSystemCatalog CatalogType = `SYSTEM_CATALOG`

// String representation for [fmt.Print]
func (f *CatalogType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CatalogType) Set(v string) error {
	switch v {
	case `DELTASHARING_CATALOG`, `FOREIGN_CATALOG`, `MANAGED_CATALOG`, `SYSTEM_CATALOG`:
		*f = CatalogType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELTASHARING_CATALOG", "FOREIGN_CATALOG", "MANAGED_CATALOG", "SYSTEM_CATALOG"`, v)
	}
}

// Type always returns CatalogType to satisfy [pflag.Value] interface
func (f *CatalogType) Type() string {
	return "CatalogType"
}

func catalogTypeToPb(st *CatalogType) (*catalogTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogTypePb(*st)
	return &pb, nil
}

func catalogTypeFromPb(pb *catalogTypePb) (*CatalogType, error) {
	if pb == nil {
		return nil, nil
	}
	st := CatalogType(*pb)
	return &st, nil
}

type CloudflareApiToken struct {
	// The Cloudflare access key id of the token.
	// Wire name: 'access_key_id'
	AccessKeyId string
	// The account id associated with the API token.
	// Wire name: 'account_id'
	AccountId string
	// The secret access token generated for the access key id
	// Wire name: 'secret_access_key'
	SecretAccessKey string
}

func cloudflareApiTokenToPb(st *CloudflareApiToken) (*cloudflareApiTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cloudflareApiTokenPb{}
	pb.AccessKeyId = st.AccessKeyId

	pb.AccountId = st.AccountId

	pb.SecretAccessKey = st.SecretAccessKey

	return pb, nil
}

func (st *CloudflareApiToken) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cloudflareApiTokenPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cloudflareApiTokenFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CloudflareApiToken) MarshalJSON() ([]byte, error) {
	pb, err := cloudflareApiTokenToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cloudflareApiTokenPb struct {
	// The Cloudflare access key id of the token.
	AccessKeyId string `json:"access_key_id"`
	// The account id associated with the API token.
	AccountId string `json:"account_id"`
	// The secret access token generated for the access key id
	SecretAccessKey string `json:"secret_access_key"`
}

func cloudflareApiTokenFromPb(pb *cloudflareApiTokenPb) (*CloudflareApiToken, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CloudflareApiToken{}
	st.AccessKeyId = pb.AccessKeyId
	st.AccountId = pb.AccountId
	st.SecretAccessKey = pb.SecretAccessKey

	return st, nil
}

type ColumnInfo struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string

	// Wire name: 'mask'
	Mask *ColumnMask
	// Name of Column.
	// Wire name: 'name'
	Name string
	// Whether field may be Null (default: true).
	// Wire name: 'nullable'
	Nullable bool
	// Partition index for column.
	// Wire name: 'partition_index'
	PartitionIndex int
	// Ordinal position of column (starting at position 0).
	// Wire name: 'position'
	Position int
	// Format of IntervalType.
	// Wire name: 'type_interval_type'
	TypeIntervalType string
	// Full data type specification, JSON-serialized.
	// Wire name: 'type_json'
	TypeJson string

	// Wire name: 'type_name'
	TypeName ColumnTypeName
	// Digits of precision; required for DecimalTypes.
	// Wire name: 'type_precision'
	TypePrecision int
	// Digits to right of decimal; Required for DecimalTypes.
	// Wire name: 'type_scale'
	TypeScale int
	// Full data type specification as SQL/catalogString text.
	// Wire name: 'type_text'
	TypeText string

	ForceSendFields []string `tf:"-"`
}

func ColumnInfoToPb(st *ColumnInfo) (*ColumnInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ColumnInfoPb{}
	pb.Comment = st.Comment

	maskPb, err := columnMaskToPb(st.Mask)
	if err != nil {
		return nil, err
	}
	if maskPb != nil {
		pb.Mask = maskPb
	}

	pb.Name = st.Name

	pb.Nullable = st.Nullable

	pb.PartitionIndex = st.PartitionIndex

	pb.Position = st.Position

	pb.TypeIntervalType = st.TypeIntervalType

	pb.TypeJson = st.TypeJson

	pb.TypeName = st.TypeName

	pb.TypePrecision = st.TypePrecision

	pb.TypeScale = st.TypeScale

	pb.TypeText = st.TypeText

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ColumnInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &ColumnInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ColumnInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ColumnInfo) MarshalJSON() ([]byte, error) {
	pb, err := ColumnInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ColumnInfoPb struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`

	Mask *columnMaskPb `json:"mask,omitempty"`
	// Name of Column.
	Name string `json:"name,omitempty"`
	// Whether field may be Null (default: true).
	Nullable bool `json:"nullable,omitempty"`
	// Partition index for column.
	PartitionIndex int `json:"partition_index,omitempty"`
	// Ordinal position of column (starting at position 0).
	Position int `json:"position,omitempty"`
	// Format of IntervalType.
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// Full data type specification, JSON-serialized.
	TypeJson string `json:"type_json,omitempty"`

	TypeName ColumnTypeName `json:"type_name,omitempty"`
	// Digits of precision; required for DecimalTypes.
	TypePrecision int `json:"type_precision,omitempty"`
	// Digits to right of decimal; Required for DecimalTypes.
	TypeScale int `json:"type_scale,omitempty"`
	// Full data type specification as SQL/catalogString text.
	TypeText string `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func ColumnInfoFromPb(pb *ColumnInfoPb) (*ColumnInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ColumnInfo{}
	st.Comment = pb.Comment
	maskField, err := columnMaskFromPb(pb.Mask)
	if err != nil {
		return nil, err
	}
	if maskField != nil {
		st.Mask = maskField
	}
	st.Name = pb.Name
	st.Nullable = pb.Nullable
	st.PartitionIndex = pb.PartitionIndex
	st.Position = pb.Position
	st.TypeIntervalType = pb.TypeIntervalType
	st.TypeJson = pb.TypeJson
	st.TypeName = pb.TypeName
	st.TypePrecision = pb.TypePrecision
	st.TypeScale = pb.TypeScale
	st.TypeText = pb.TypeText

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *ColumnInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ColumnInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ColumnMask struct {
	// The full name of the column mask SQL UDF.
	// Wire name: 'function_name'
	FunctionName string
	// The list of additional table columns to be passed as input to the column
	// mask function. The first arg of the mask function should be of the type
	// of the column being masked and the types of the rest of the args should
	// match the types of columns in 'using_column_names'.
	// Wire name: 'using_column_names'
	UsingColumnNames []string

	ForceSendFields []string `tf:"-"`
}

func columnMaskToPb(st *ColumnMask) (*columnMaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &columnMaskPb{}
	pb.FunctionName = st.FunctionName

	pb.UsingColumnNames = st.UsingColumnNames

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ColumnMask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &columnMaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := columnMaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ColumnMask) MarshalJSON() ([]byte, error) {
	pb, err := columnMaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type columnMaskPb struct {
	// The full name of the column mask SQL UDF.
	FunctionName string `json:"function_name,omitempty"`
	// The list of additional table columns to be passed as input to the column
	// mask function. The first arg of the mask function should be of the type
	// of the column being masked and the types of the rest of the args should
	// match the types of columns in 'using_column_names'.
	UsingColumnNames []string `json:"using_column_names,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func columnMaskFromPb(pb *columnMaskPb) (*ColumnMask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ColumnMask{}
	st.FunctionName = pb.FunctionName
	st.UsingColumnNames = pb.UsingColumnNames

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *columnMaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st columnMaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ColumnTypeName string
type columnTypeNamePb string

const ColumnTypeNameArray ColumnTypeName = `ARRAY`

const ColumnTypeNameBinary ColumnTypeName = `BINARY`

const ColumnTypeNameBoolean ColumnTypeName = `BOOLEAN`

const ColumnTypeNameByte ColumnTypeName = `BYTE`

const ColumnTypeNameChar ColumnTypeName = `CHAR`

const ColumnTypeNameDate ColumnTypeName = `DATE`

const ColumnTypeNameDecimal ColumnTypeName = `DECIMAL`

const ColumnTypeNameDouble ColumnTypeName = `DOUBLE`

const ColumnTypeNameFloat ColumnTypeName = `FLOAT`

const ColumnTypeNameGeography ColumnTypeName = `GEOGRAPHY`

const ColumnTypeNameGeometry ColumnTypeName = `GEOMETRY`

const ColumnTypeNameInt ColumnTypeName = `INT`

const ColumnTypeNameInterval ColumnTypeName = `INTERVAL`

const ColumnTypeNameLong ColumnTypeName = `LONG`

const ColumnTypeNameMap ColumnTypeName = `MAP`

const ColumnTypeNameNull ColumnTypeName = `NULL`

const ColumnTypeNameShort ColumnTypeName = `SHORT`

const ColumnTypeNameString ColumnTypeName = `STRING`

const ColumnTypeNameStruct ColumnTypeName = `STRUCT`

const ColumnTypeNameTableType ColumnTypeName = `TABLE_TYPE`

const ColumnTypeNameTimestamp ColumnTypeName = `TIMESTAMP`

const ColumnTypeNameTimestampNtz ColumnTypeName = `TIMESTAMP_NTZ`

const ColumnTypeNameUserDefinedType ColumnTypeName = `USER_DEFINED_TYPE`

const ColumnTypeNameVariant ColumnTypeName = `VARIANT`

// String representation for [fmt.Print]
func (f *ColumnTypeName) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ColumnTypeName) Set(v string) error {
	switch v {
	case `ARRAY`, `BINARY`, `BOOLEAN`, `BYTE`, `CHAR`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `GEOGRAPHY`, `GEOMETRY`, `INT`, `INTERVAL`, `LONG`, `MAP`, `NULL`, `SHORT`, `STRING`, `STRUCT`, `TABLE_TYPE`, `TIMESTAMP`, `TIMESTAMP_NTZ`, `USER_DEFINED_TYPE`, `VARIANT`:
		*f = ColumnTypeName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARRAY", "BINARY", "BOOLEAN", "BYTE", "CHAR", "DATE", "DECIMAL", "DOUBLE", "FLOAT", "GEOGRAPHY", "GEOMETRY", "INT", "INTERVAL", "LONG", "MAP", "NULL", "SHORT", "STRING", "STRUCT", "TABLE_TYPE", "TIMESTAMP", "TIMESTAMP_NTZ", "USER_DEFINED_TYPE", "VARIANT"`, v)
	}
}

// Type always returns ColumnTypeName to satisfy [pflag.Value] interface
func (f *ColumnTypeName) Type() string {
	return "ColumnTypeName"
}

func columnTypeNameToPb(st *ColumnTypeName) (*columnTypeNamePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := columnTypeNamePb(*st)
	return &pb, nil
}

func columnTypeNameFromPb(pb *columnTypeNamePb) (*ColumnTypeName, error) {
	if pb == nil {
		return nil, nil
	}
	st := ColumnTypeName(*pb)
	return &st, nil
}

type ConnectionInfo struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// Unique identifier of the Connection.
	// Wire name: 'connection_id'
	ConnectionId string
	// The type of connection.
	// Wire name: 'connection_type'
	ConnectionType ConnectionType
	// Time at which this connection was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of connection creator.
	// Wire name: 'created_by'
	CreatedBy string
	// The type of credential.
	// Wire name: 'credential_type'
	CredentialType CredentialType
	// Full name of connection.
	// Wire name: 'full_name'
	FullName string
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string
	// Name of the connection.
	// Wire name: 'name'
	Name string
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string
	// Username of current owner of the connection.
	// Wire name: 'owner'
	Owner string
	// An object containing map of key-value properties attached to the
	// connection.
	// Wire name: 'properties'
	Properties map[string]string
	// Status of an asynchronously provisioned resource.
	// Wire name: 'provisioning_info'
	ProvisioningInfo *ProvisioningInfo
	// If the connection is read only.
	// Wire name: 'read_only'
	ReadOnly bool

	// Wire name: 'securable_type'
	SecurableType string
	// Time at which this connection was updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// Username of user who last modified connection.
	// Wire name: 'updated_by'
	UpdatedBy string
	// URL of the remote data source, extracted from options.
	// Wire name: 'url'
	Url string

	ForceSendFields []string `tf:"-"`
}

func connectionInfoToPb(st *ConnectionInfo) (*connectionInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &connectionInfoPb{}
	pb.Comment = st.Comment

	pb.ConnectionId = st.ConnectionId

	pb.ConnectionType = st.ConnectionType

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	pb.CredentialType = st.CredentialType

	pb.FullName = st.FullName

	pb.MetastoreId = st.MetastoreId

	pb.Name = st.Name

	pb.Options = st.Options

	pb.Owner = st.Owner

	pb.Properties = st.Properties

	provisioningInfoPb, err := provisioningInfoToPb(st.ProvisioningInfo)
	if err != nil {
		return nil, err
	}
	if provisioningInfoPb != nil {
		pb.ProvisioningInfo = provisioningInfoPb
	}

	pb.ReadOnly = st.ReadOnly

	pb.SecurableType = st.SecurableType

	pb.UpdatedAt = st.UpdatedAt

	pb.UpdatedBy = st.UpdatedBy

	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ConnectionInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &connectionInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := connectionInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ConnectionInfo) MarshalJSON() ([]byte, error) {
	pb, err := connectionInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type connectionInfoPb struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Unique identifier of the Connection.
	ConnectionId string `json:"connection_id,omitempty"`
	// The type of connection.
	ConnectionType ConnectionType `json:"connection_type,omitempty"`
	// Time at which this connection was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of connection creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The type of credential.
	CredentialType CredentialType `json:"credential_type,omitempty"`
	// Full name of connection.
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of the connection.
	Name string `json:"name,omitempty"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options,omitempty"`
	// Username of current owner of the connection.
	Owner string `json:"owner,omitempty"`
	// An object containing map of key-value properties attached to the
	// connection.
	Properties map[string]string `json:"properties,omitempty"`
	// Status of an asynchronously provisioned resource.
	ProvisioningInfo *provisioningInfoPb `json:"provisioning_info,omitempty"`
	// If the connection is read only.
	ReadOnly bool `json:"read_only,omitempty"`

	SecurableType string `json:"securable_type,omitempty"`
	// Time at which this connection was updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified connection.
	UpdatedBy string `json:"updated_by,omitempty"`
	// URL of the remote data source, extracted from options.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func connectionInfoFromPb(pb *connectionInfoPb) (*ConnectionInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ConnectionInfo{}
	st.Comment = pb.Comment
	st.ConnectionId = pb.ConnectionId
	st.ConnectionType = pb.ConnectionType
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.CredentialType = pb.CredentialType
	st.FullName = pb.FullName
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Options = pb.Options
	st.Owner = pb.Owner
	st.Properties = pb.Properties
	provisioningInfoField, err := provisioningInfoFromPb(pb.ProvisioningInfo)
	if err != nil {
		return nil, err
	}
	if provisioningInfoField != nil {
		st.ProvisioningInfo = provisioningInfoField
	}
	st.ReadOnly = pb.ReadOnly
	st.SecurableType = pb.SecurableType
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *connectionInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st connectionInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The type of connection.
type ConnectionType string
type connectionTypePb string

const ConnectionTypeBigquery ConnectionType = `BIGQUERY`

const ConnectionTypeDatabricks ConnectionType = `DATABRICKS`

const ConnectionTypeGlue ConnectionType = `GLUE`

const ConnectionTypeHiveMetastore ConnectionType = `HIVE_METASTORE`

const ConnectionTypeHttp ConnectionType = `HTTP`

const ConnectionTypeMysql ConnectionType = `MYSQL`

const ConnectionTypeOracle ConnectionType = `ORACLE`

const ConnectionTypePostgresql ConnectionType = `POSTGRESQL`

const ConnectionTypeRedshift ConnectionType = `REDSHIFT`

const ConnectionTypeSnowflake ConnectionType = `SNOWFLAKE`

const ConnectionTypeSqldw ConnectionType = `SQLDW`

const ConnectionTypeSqlserver ConnectionType = `SQLSERVER`

const ConnectionTypeTeradata ConnectionType = `TERADATA`

// String representation for [fmt.Print]
func (f *ConnectionType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ConnectionType) Set(v string) error {
	switch v {
	case `BIGQUERY`, `DATABRICKS`, `GLUE`, `HIVE_METASTORE`, `HTTP`, `MYSQL`, `ORACLE`, `POSTGRESQL`, `REDSHIFT`, `SNOWFLAKE`, `SQLDW`, `SQLSERVER`, `TERADATA`:
		*f = ConnectionType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BIGQUERY", "DATABRICKS", "GLUE", "HIVE_METASTORE", "HTTP", "MYSQL", "ORACLE", "POSTGRESQL", "REDSHIFT", "SNOWFLAKE", "SQLDW", "SQLSERVER", "TERADATA"`, v)
	}
}

// Type always returns ConnectionType to satisfy [pflag.Value] interface
func (f *ConnectionType) Type() string {
	return "ConnectionType"
}

func connectionTypeToPb(st *ConnectionType) (*connectionTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := connectionTypePb(*st)
	return &pb, nil
}

func connectionTypeFromPb(pb *connectionTypePb) (*ConnectionType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ConnectionType(*pb)
	return &st, nil
}

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_CONTINUOUS_UPDATE or the ONLINE_UPDATING_PIPELINE_RESOURCES state.
type ContinuousUpdateStatus struct {
	// Progress of the initial data synchronization.
	// Wire name: 'initial_pipeline_sync_progress'
	InitialPipelineSyncProgress *PipelineProgress
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	// Wire name: 'last_processed_commit_version'
	LastProcessedCommitVersion int64
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	// Wire name: 'timestamp'
	Timestamp string

	ForceSendFields []string `tf:"-"`
}

func continuousUpdateStatusToPb(st *ContinuousUpdateStatus) (*continuousUpdateStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &continuousUpdateStatusPb{}
	initialPipelineSyncProgressPb, err := pipelineProgressToPb(st.InitialPipelineSyncProgress)
	if err != nil {
		return nil, err
	}
	if initialPipelineSyncProgressPb != nil {
		pb.InitialPipelineSyncProgress = initialPipelineSyncProgressPb
	}

	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion

	pb.Timestamp = st.Timestamp

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ContinuousUpdateStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &continuousUpdateStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := continuousUpdateStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ContinuousUpdateStatus) MarshalJSON() ([]byte, error) {
	pb, err := continuousUpdateStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type continuousUpdateStatusPb struct {
	// Progress of the initial data synchronization.
	InitialPipelineSyncProgress *pipelineProgressPb `json:"initial_pipeline_sync_progress,omitempty"`
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func continuousUpdateStatusFromPb(pb *continuousUpdateStatusPb) (*ContinuousUpdateStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ContinuousUpdateStatus{}
	initialPipelineSyncProgressField, err := pipelineProgressFromPb(pb.InitialPipelineSyncProgress)
	if err != nil {
		return nil, err
	}
	if initialPipelineSyncProgressField != nil {
		st.InitialPipelineSyncProgress = initialPipelineSyncProgressField
	}
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
	st.Timestamp = pb.Timestamp

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *continuousUpdateStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st continuousUpdateStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCatalog struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// The name of the connection to an external data source.
	// Wire name: 'connection_name'
	ConnectionName string
	// Name of catalog.
	// Wire name: 'name'
	Name string
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	// Wire name: 'provider_name'
	ProviderName string
	// The name of the share under the share provider.
	// Wire name: 'share_name'
	ShareName string
	// Storage root URL for managed tables within catalog.
	// Wire name: 'storage_root'
	StorageRoot string

	ForceSendFields []string `tf:"-"`
}

func createCatalogToPb(st *CreateCatalog) (*createCatalogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCatalogPb{}
	pb.Comment = st.Comment

	pb.ConnectionName = st.ConnectionName

	pb.Name = st.Name

	pb.Options = st.Options

	pb.Properties = st.Properties

	pb.ProviderName = st.ProviderName

	pb.ShareName = st.ShareName

	pb.StorageRoot = st.StorageRoot

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateCatalog) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCatalogPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCatalogFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCatalog) MarshalJSON() ([]byte, error) {
	pb, err := createCatalogToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createCatalogPb struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// The name of the connection to an external data source.
	ConnectionName string `json:"connection_name,omitempty"`
	// Name of catalog.
	Name string `json:"name"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName string `json:"provider_name,omitempty"`
	// The name of the share under the share provider.
	ShareName string `json:"share_name,omitempty"`
	// Storage root URL for managed tables within catalog.
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createCatalogFromPb(pb *createCatalogPb) (*CreateCatalog, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCatalog{}
	st.Comment = pb.Comment
	st.ConnectionName = pb.ConnectionName
	st.Name = pb.Name
	st.Options = pb.Options
	st.Properties = pb.Properties
	st.ProviderName = pb.ProviderName
	st.ShareName = pb.ShareName
	st.StorageRoot = pb.StorageRoot

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createCatalogPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createCatalogPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateConnection struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// The type of connection.
	// Wire name: 'connection_type'
	ConnectionType ConnectionType
	// Name of the connection.
	// Wire name: 'name'
	Name string
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string
	// An object containing map of key-value properties attached to the
	// connection.
	// Wire name: 'properties'
	Properties map[string]string
	// If the connection is read only.
	// Wire name: 'read_only'
	ReadOnly bool

	ForceSendFields []string `tf:"-"`
}

func createConnectionToPb(st *CreateConnection) (*createConnectionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createConnectionPb{}
	pb.Comment = st.Comment

	pb.ConnectionType = st.ConnectionType

	pb.Name = st.Name

	pb.Options = st.Options

	pb.Properties = st.Properties

	pb.ReadOnly = st.ReadOnly

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateConnection) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createConnectionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createConnectionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateConnection) MarshalJSON() ([]byte, error) {
	pb, err := createConnectionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createConnectionPb struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// The type of connection.
	ConnectionType ConnectionType `json:"connection_type"`
	// Name of the connection.
	Name string `json:"name"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options"`
	// An object containing map of key-value properties attached to the
	// connection.
	Properties map[string]string `json:"properties,omitempty"`
	// If the connection is read only.
	ReadOnly bool `json:"read_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createConnectionFromPb(pb *createConnectionPb) (*CreateConnection, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateConnection{}
	st.Comment = pb.Comment
	st.ConnectionType = pb.ConnectionType
	st.Name = pb.Name
	st.Options = pb.Options
	st.Properties = pb.Properties
	st.ReadOnly = pb.ReadOnly

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createConnectionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createConnectionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCredentialRequest struct {
	// The AWS IAM role configuration
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRole
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentity
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	// Wire name: 'name'
	Name string
	// Indicates the purpose of the credential.
	// Wire name: 'purpose'
	Purpose CredentialPurpose
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	// Wire name: 'read_only'
	ReadOnly bool
	// Optional. Supplying true to this argument skips validation of the created
	// set of credentials.
	// Wire name: 'skip_validation'
	SkipValidation bool

	ForceSendFields []string `tf:"-"`
}

func createCredentialRequestToPb(st *CreateCredentialRequest) (*createCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCredentialRequestPb{}
	awsIamRolePb, err := awsIamRoleToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}

	azureManagedIdentityPb, err := azureManagedIdentityToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}

	azureServicePrincipalPb, err := azureServicePrincipalToPb(st.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalPb != nil {
		pb.AzureServicePrincipal = azureServicePrincipalPb
	}

	pb.Comment = st.Comment

	databricksGcpServiceAccountPb, err := databricksGcpServiceAccountToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}

	pb.Name = st.Name

	pb.Purpose = st.Purpose

	pb.ReadOnly = st.ReadOnly

	pb.SkipValidation = st.SkipValidation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := createCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createCredentialRequestPb struct {
	// The AWS IAM role configuration
	AwsIamRole *awsIamRolePb `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *azureManagedIdentityPb `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	AzureServicePrincipal *azureServicePrincipalPb `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount *databricksGcpServiceAccountPb `json:"databricks_gcp_service_account,omitempty"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	Name string `json:"name"`
	// Indicates the purpose of the credential.
	Purpose CredentialPurpose `json:"purpose,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool `json:"read_only,omitempty"`
	// Optional. Supplying true to this argument skips validation of the created
	// set of credentials.
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createCredentialRequestFromPb(pb *createCredentialRequestPb) (*CreateCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialRequest{}
	awsIamRoleField, err := awsIamRoleFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := azureManagedIdentityFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	azureServicePrincipalField, err := azureServicePrincipalFromPb(pb.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalField != nil {
		st.AzureServicePrincipal = azureServicePrincipalField
	}
	st.Comment = pb.Comment
	databricksGcpServiceAccountField, err := databricksGcpServiceAccountFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	st.Name = pb.Name
	st.Purpose = pb.Purpose
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateExternalLocation struct {
	// The AWS access point to use when accesing s3 for this external location.
	// Wire name: 'access_point'
	AccessPoint string
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// Name of the storage credential used with this location.
	// Wire name: 'credential_name'
	CredentialName string
	// Encryption options that apply to clients connecting to cloud storage.
	// Wire name: 'encryption_details'
	EncryptionDetails *EncryptionDetails
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	// Wire name: 'fallback'
	Fallback bool
	// Name of the external location.
	// Wire name: 'name'
	Name string
	// Indicates whether the external location is read-only.
	// Wire name: 'read_only'
	ReadOnly bool
	// Skips validation of the storage credential associated with the external
	// location.
	// Wire name: 'skip_validation'
	SkipValidation bool
	// Path URL of the external location.
	// Wire name: 'url'
	Url string

	ForceSendFields []string `tf:"-"`
}

func createExternalLocationToPb(st *CreateExternalLocation) (*createExternalLocationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createExternalLocationPb{}
	pb.AccessPoint = st.AccessPoint

	pb.Comment = st.Comment

	pb.CredentialName = st.CredentialName

	encryptionDetailsPb, err := encryptionDetailsToPb(st.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsPb != nil {
		pb.EncryptionDetails = encryptionDetailsPb
	}

	pb.Fallback = st.Fallback

	pb.Name = st.Name

	pb.ReadOnly = st.ReadOnly

	pb.SkipValidation = st.SkipValidation

	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateExternalLocation) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createExternalLocationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createExternalLocationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateExternalLocation) MarshalJSON() ([]byte, error) {
	pb, err := createExternalLocationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createExternalLocationPb struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string `json:"access_point,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of the storage credential used with this location.
	CredentialName string `json:"credential_name"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *encryptionDetailsPb `json:"encryption_details,omitempty"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback bool `json:"fallback,omitempty"`
	// Name of the external location.
	Name string `json:"name"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation bool `json:"skip_validation,omitempty"`
	// Path URL of the external location.
	Url string `json:"url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createExternalLocationFromPb(pb *createExternalLocationPb) (*CreateExternalLocation, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExternalLocation{}
	st.AccessPoint = pb.AccessPoint
	st.Comment = pb.Comment
	st.CredentialName = pb.CredentialName
	encryptionDetailsField, err := encryptionDetailsFromPb(pb.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsField != nil {
		st.EncryptionDetails = encryptionDetailsField
	}
	st.Fallback = pb.Fallback
	st.Name = pb.Name
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createExternalLocationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createExternalLocationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateFunction struct {
	// Name of parent catalog.
	// Wire name: 'catalog_name'
	CatalogName string
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// Scalar function return data type.
	// Wire name: 'data_type'
	DataType ColumnTypeName
	// External function language.
	// Wire name: 'external_language'
	ExternalLanguage string
	// External function name.
	// Wire name: 'external_name'
	ExternalName string
	// Pretty printed function data type.
	// Wire name: 'full_data_type'
	FullDataType string

	// Wire name: 'input_params'
	InputParams FunctionParameterInfos
	// Whether the function is deterministic.
	// Wire name: 'is_deterministic'
	IsDeterministic bool
	// Function null call.
	// Wire name: 'is_null_call'
	IsNullCall bool
	// Name of function, relative to parent schema.
	// Wire name: 'name'
	Name string
	// Function parameter style. **S** is the value for SQL.
	// Wire name: 'parameter_style'
	ParameterStyle CreateFunctionParameterStyle
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	// Wire name: 'properties'
	Properties string
	// Table function return parameters.
	// Wire name: 'return_params'
	ReturnParams *FunctionParameterInfos
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	// Wire name: 'routine_body'
	RoutineBody CreateFunctionRoutineBody
	// Function body.
	// Wire name: 'routine_definition'
	RoutineDefinition string
	// Function dependencies.
	// Wire name: 'routine_dependencies'
	RoutineDependencies *DependencyList
	// Name of parent schema relative to its parent catalog.
	// Wire name: 'schema_name'
	SchemaName string
	// Function security type.
	// Wire name: 'security_type'
	SecurityType CreateFunctionSecurityType
	// Specific name of the function; Reserved for future use.
	// Wire name: 'specific_name'
	SpecificName string
	// Function SQL data access.
	// Wire name: 'sql_data_access'
	SqlDataAccess CreateFunctionSqlDataAccess
	// List of schemes whose objects can be referenced without qualification.
	// Wire name: 'sql_path'
	SqlPath string

	ForceSendFields []string `tf:"-"`
}

func createFunctionToPb(st *CreateFunction) (*createFunctionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createFunctionPb{}
	pb.CatalogName = st.CatalogName

	pb.Comment = st.Comment

	pb.DataType = st.DataType

	pb.ExternalLanguage = st.ExternalLanguage

	pb.ExternalName = st.ExternalName

	pb.FullDataType = st.FullDataType

	inputParamsPb, err := functionParameterInfosToPb(&st.InputParams)
	if err != nil {
		return nil, err
	}
	if inputParamsPb != nil {
		pb.InputParams = *inputParamsPb
	}

	pb.IsDeterministic = st.IsDeterministic

	pb.IsNullCall = st.IsNullCall

	pb.Name = st.Name

	pb.ParameterStyle = st.ParameterStyle

	pb.Properties = st.Properties

	returnParamsPb, err := functionParameterInfosToPb(st.ReturnParams)
	if err != nil {
		return nil, err
	}
	if returnParamsPb != nil {
		pb.ReturnParams = returnParamsPb
	}

	pb.RoutineBody = st.RoutineBody

	pb.RoutineDefinition = st.RoutineDefinition

	routineDependenciesPb, err := dependencyListToPb(st.RoutineDependencies)
	if err != nil {
		return nil, err
	}
	if routineDependenciesPb != nil {
		pb.RoutineDependencies = routineDependenciesPb
	}

	pb.SchemaName = st.SchemaName

	pb.SecurityType = st.SecurityType

	pb.SpecificName = st.SpecificName

	pb.SqlDataAccess = st.SqlDataAccess

	pb.SqlPath = st.SqlPath

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateFunction) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createFunctionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createFunctionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateFunction) MarshalJSON() ([]byte, error) {
	pb, err := createFunctionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createFunctionPb struct {
	// Name of parent catalog.
	CatalogName string `json:"catalog_name"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Scalar function return data type.
	DataType ColumnTypeName `json:"data_type"`
	// External function language.
	ExternalLanguage string `json:"external_language,omitempty"`
	// External function name.
	ExternalName string `json:"external_name,omitempty"`
	// Pretty printed function data type.
	FullDataType string `json:"full_data_type"`

	InputParams functionParameterInfosPb `json:"input_params"`
	// Whether the function is deterministic.
	IsDeterministic bool `json:"is_deterministic"`
	// Function null call.
	IsNullCall bool `json:"is_null_call"`
	// Name of function, relative to parent schema.
	Name string `json:"name"`
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle CreateFunctionParameterStyle `json:"parameter_style"`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties string `json:"properties,omitempty"`
	// Table function return parameters.
	ReturnParams *functionParameterInfosPb `json:"return_params,omitempty"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody CreateFunctionRoutineBody `json:"routine_body"`
	// Function body.
	RoutineDefinition string `json:"routine_definition"`
	// Function dependencies.
	RoutineDependencies *dependencyListPb `json:"routine_dependencies,omitempty"`
	// Name of parent schema relative to its parent catalog.
	SchemaName string `json:"schema_name"`
	// Function security type.
	SecurityType CreateFunctionSecurityType `json:"security_type"`
	// Specific name of the function; Reserved for future use.
	SpecificName string `json:"specific_name"`
	// Function SQL data access.
	SqlDataAccess CreateFunctionSqlDataAccess `json:"sql_data_access"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string `json:"sql_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createFunctionFromPb(pb *createFunctionPb) (*CreateFunction, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateFunction{}
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.DataType = pb.DataType
	st.ExternalLanguage = pb.ExternalLanguage
	st.ExternalName = pb.ExternalName
	st.FullDataType = pb.FullDataType
	inputParamsField, err := functionParameterInfosFromPb(&pb.InputParams)
	if err != nil {
		return nil, err
	}
	if inputParamsField != nil {
		st.InputParams = *inputParamsField
	}
	st.IsDeterministic = pb.IsDeterministic
	st.IsNullCall = pb.IsNullCall
	st.Name = pb.Name
	st.ParameterStyle = pb.ParameterStyle
	st.Properties = pb.Properties
	returnParamsField, err := functionParameterInfosFromPb(pb.ReturnParams)
	if err != nil {
		return nil, err
	}
	if returnParamsField != nil {
		st.ReturnParams = returnParamsField
	}
	st.RoutineBody = pb.RoutineBody
	st.RoutineDefinition = pb.RoutineDefinition
	routineDependenciesField, err := dependencyListFromPb(pb.RoutineDependencies)
	if err != nil {
		return nil, err
	}
	if routineDependenciesField != nil {
		st.RoutineDependencies = routineDependenciesField
	}
	st.SchemaName = pb.SchemaName
	st.SecurityType = pb.SecurityType
	st.SpecificName = pb.SpecificName
	st.SqlDataAccess = pb.SqlDataAccess
	st.SqlPath = pb.SqlPath

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createFunctionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createFunctionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Function parameter style. **S** is the value for SQL.
type CreateFunctionParameterStyle string
type createFunctionParameterStylePb string

const CreateFunctionParameterStyleS CreateFunctionParameterStyle = `S`

// String representation for [fmt.Print]
func (f *CreateFunctionParameterStyle) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateFunctionParameterStyle) Set(v string) error {
	switch v {
	case `S`:
		*f = CreateFunctionParameterStyle(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "S"`, v)
	}
}

// Type always returns CreateFunctionParameterStyle to satisfy [pflag.Value] interface
func (f *CreateFunctionParameterStyle) Type() string {
	return "CreateFunctionParameterStyle"
}

func createFunctionParameterStyleToPb(st *CreateFunctionParameterStyle) (*createFunctionParameterStylePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := createFunctionParameterStylePb(*st)
	return &pb, nil
}

func createFunctionParameterStyleFromPb(pb *createFunctionParameterStylePb) (*CreateFunctionParameterStyle, error) {
	if pb == nil {
		return nil, nil
	}
	st := CreateFunctionParameterStyle(*pb)
	return &st, nil
}

type CreateFunctionRequest struct {
	// Partial __FunctionInfo__ specifying the function to be created.
	// Wire name: 'function_info'
	FunctionInfo CreateFunction
}

func createFunctionRequestToPb(st *CreateFunctionRequest) (*createFunctionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createFunctionRequestPb{}
	functionInfoPb, err := createFunctionToPb(&st.FunctionInfo)
	if err != nil {
		return nil, err
	}
	if functionInfoPb != nil {
		pb.FunctionInfo = *functionInfoPb
	}

	return pb, nil
}

func (st *CreateFunctionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createFunctionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createFunctionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateFunctionRequest) MarshalJSON() ([]byte, error) {
	pb, err := createFunctionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createFunctionRequestPb struct {
	// Partial __FunctionInfo__ specifying the function to be created.
	FunctionInfo createFunctionPb `json:"function_info"`
}

func createFunctionRequestFromPb(pb *createFunctionRequestPb) (*CreateFunctionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateFunctionRequest{}
	functionInfoField, err := createFunctionFromPb(&pb.FunctionInfo)
	if err != nil {
		return nil, err
	}
	if functionInfoField != nil {
		st.FunctionInfo = *functionInfoField
	}

	return st, nil
}

// Function language. When **EXTERNAL** is used, the language of the routine
// function should be specified in the __external_language__ field, and the
// __return_params__ of the function cannot be used (as **TABLE** return type is
// not supported), and the __sql_data_access__ field must be **NO_SQL**.
type CreateFunctionRoutineBody string
type createFunctionRoutineBodyPb string

const CreateFunctionRoutineBodyExternal CreateFunctionRoutineBody = `EXTERNAL`

const CreateFunctionRoutineBodySql CreateFunctionRoutineBody = `SQL`

// String representation for [fmt.Print]
func (f *CreateFunctionRoutineBody) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateFunctionRoutineBody) Set(v string) error {
	switch v {
	case `EXTERNAL`, `SQL`:
		*f = CreateFunctionRoutineBody(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "SQL"`, v)
	}
}

// Type always returns CreateFunctionRoutineBody to satisfy [pflag.Value] interface
func (f *CreateFunctionRoutineBody) Type() string {
	return "CreateFunctionRoutineBody"
}

func createFunctionRoutineBodyToPb(st *CreateFunctionRoutineBody) (*createFunctionRoutineBodyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := createFunctionRoutineBodyPb(*st)
	return &pb, nil
}

func createFunctionRoutineBodyFromPb(pb *createFunctionRoutineBodyPb) (*CreateFunctionRoutineBody, error) {
	if pb == nil {
		return nil, nil
	}
	st := CreateFunctionRoutineBody(*pb)
	return &st, nil
}

// The security type of the function.
type CreateFunctionSecurityType string
type createFunctionSecurityTypePb string

const CreateFunctionSecurityTypeDefiner CreateFunctionSecurityType = `DEFINER`

// String representation for [fmt.Print]
func (f *CreateFunctionSecurityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateFunctionSecurityType) Set(v string) error {
	switch v {
	case `DEFINER`:
		*f = CreateFunctionSecurityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEFINER"`, v)
	}
}

// Type always returns CreateFunctionSecurityType to satisfy [pflag.Value] interface
func (f *CreateFunctionSecurityType) Type() string {
	return "CreateFunctionSecurityType"
}

func createFunctionSecurityTypeToPb(st *CreateFunctionSecurityType) (*createFunctionSecurityTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := createFunctionSecurityTypePb(*st)
	return &pb, nil
}

func createFunctionSecurityTypeFromPb(pb *createFunctionSecurityTypePb) (*CreateFunctionSecurityType, error) {
	if pb == nil {
		return nil, nil
	}
	st := CreateFunctionSecurityType(*pb)
	return &st, nil
}

// Function SQL data access.
type CreateFunctionSqlDataAccess string
type createFunctionSqlDataAccessPb string

const CreateFunctionSqlDataAccessContainsSql CreateFunctionSqlDataAccess = `CONTAINS_SQL`

const CreateFunctionSqlDataAccessNoSql CreateFunctionSqlDataAccess = `NO_SQL`

const CreateFunctionSqlDataAccessReadsSqlData CreateFunctionSqlDataAccess = `READS_SQL_DATA`

// String representation for [fmt.Print]
func (f *CreateFunctionSqlDataAccess) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateFunctionSqlDataAccess) Set(v string) error {
	switch v {
	case `CONTAINS_SQL`, `NO_SQL`, `READS_SQL_DATA`:
		*f = CreateFunctionSqlDataAccess(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTAINS_SQL", "NO_SQL", "READS_SQL_DATA"`, v)
	}
}

// Type always returns CreateFunctionSqlDataAccess to satisfy [pflag.Value] interface
func (f *CreateFunctionSqlDataAccess) Type() string {
	return "CreateFunctionSqlDataAccess"
}

func createFunctionSqlDataAccessToPb(st *CreateFunctionSqlDataAccess) (*createFunctionSqlDataAccessPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := createFunctionSqlDataAccessPb(*st)
	return &pb, nil
}

func createFunctionSqlDataAccessFromPb(pb *createFunctionSqlDataAccessPb) (*CreateFunctionSqlDataAccess, error) {
	if pb == nil {
		return nil, nil
	}
	st := CreateFunctionSqlDataAccess(*pb)
	return &st, nil
}

type CreateMetastore struct {
	// The user-specified name of the metastore.
	// Wire name: 'name'
	Name string
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	// The field can be omitted in the __workspace-level__ __API__ but not in
	// the __account-level__ __API__. If this field is omitted, the region of
	// the workspace receiving the request will be used.
	// Wire name: 'region'
	Region string
	// The storage root URL for metastore
	// Wire name: 'storage_root'
	StorageRoot string

	ForceSendFields []string `tf:"-"`
}

func createMetastoreToPb(st *CreateMetastore) (*createMetastorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createMetastorePb{}
	pb.Name = st.Name

	pb.Region = st.Region

	pb.StorageRoot = st.StorageRoot

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateMetastore) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createMetastorePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createMetastoreFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateMetastore) MarshalJSON() ([]byte, error) {
	pb, err := createMetastoreToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createMetastorePb struct {
	// The user-specified name of the metastore.
	Name string `json:"name"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	// The field can be omitted in the __workspace-level__ __API__ but not in
	// the __account-level__ __API__. If this field is omitted, the region of
	// the workspace receiving the request will be used.
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createMetastoreFromPb(pb *createMetastorePb) (*CreateMetastore, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateMetastore{}
	st.Name = pb.Name
	st.Region = pb.Region
	st.StorageRoot = pb.StorageRoot

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createMetastorePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createMetastorePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateMetastoreAssignment struct {
	// The name of the default catalog in the metastore. This field is
	// depracted. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	// Wire name: 'default_catalog_name'
	DefaultCatalogName string
	// The unique ID of the metastore.
	// Wire name: 'metastore_id'
	MetastoreId string
	// A workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func createMetastoreAssignmentToPb(st *CreateMetastoreAssignment) (*createMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createMetastoreAssignmentPb{}
	pb.DefaultCatalogName = st.DefaultCatalogName

	pb.MetastoreId = st.MetastoreId

	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func (st *CreateMetastoreAssignment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createMetastoreAssignmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createMetastoreAssignmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateMetastoreAssignment) MarshalJSON() ([]byte, error) {
	pb, err := createMetastoreAssignmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createMetastoreAssignmentPb struct {
	// The name of the default catalog in the metastore. This field is
	// depracted. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	DefaultCatalogName string `json:"default_catalog_name"`
	// The unique ID of the metastore.
	MetastoreId string `json:"metastore_id"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func createMetastoreAssignmentFromPb(pb *createMetastoreAssignmentPb) (*CreateMetastoreAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateMetastoreAssignment{}
	st.DefaultCatalogName = pb.DefaultCatalogName
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

type CreateMonitor struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	// Wire name: 'assets_dir'
	AssetsDir string
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	// Wire name: 'baseline_table_name'
	BaselineTableName string
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	// Wire name: 'custom_metrics'
	CustomMetrics []MonitorMetric
	// The data classification config for the monitor.
	// Wire name: 'data_classification_config'
	DataClassificationConfig *MonitorDataClassificationConfig
	// Configuration for monitoring inference logs.
	// Wire name: 'inference_log'
	InferenceLog *MonitorInferenceLog
	// The notification settings for the monitor.
	// Wire name: 'notifications'
	Notifications *MonitorNotifications
	// Schema where output metric tables are created.
	// Wire name: 'output_schema_name'
	OutputSchemaName string
	// The schedule for automatically updating and refreshing metric tables.
	// Wire name: 'schedule'
	Schedule *MonitorCronSchedule
	// Whether to skip creating a default dashboard summarizing data quality
	// metrics.
	// Wire name: 'skip_builtin_dashboard'
	SkipBuiltinDashboard bool
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	// Wire name: 'slicing_exprs'
	SlicingExprs []string
	// Configuration for monitoring snapshot tables.
	// Wire name: 'snapshot'
	Snapshot *MonitorSnapshot
	// Full name of the table.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
	// Configuration for monitoring time series tables.
	// Wire name: 'time_series'
	TimeSeries *MonitorTimeSeries
	// Optional argument to specify the warehouse for dashboard creation. If not
	// specified, the first running warehouse will be used.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func createMonitorToPb(st *CreateMonitor) (*createMonitorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createMonitorPb{}
	pb.AssetsDir = st.AssetsDir

	pb.BaselineTableName = st.BaselineTableName

	var customMetricsPb []monitorMetricPb
	for _, item := range st.CustomMetrics {
		itemPb, err := monitorMetricToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customMetricsPb = append(customMetricsPb, *itemPb)
		}
	}
	pb.CustomMetrics = customMetricsPb

	dataClassificationConfigPb, err := monitorDataClassificationConfigToPb(st.DataClassificationConfig)
	if err != nil {
		return nil, err
	}
	if dataClassificationConfigPb != nil {
		pb.DataClassificationConfig = dataClassificationConfigPb
	}

	inferenceLogPb, err := monitorInferenceLogToPb(st.InferenceLog)
	if err != nil {
		return nil, err
	}
	if inferenceLogPb != nil {
		pb.InferenceLog = inferenceLogPb
	}

	notificationsPb, err := monitorNotificationsToPb(st.Notifications)
	if err != nil {
		return nil, err
	}
	if notificationsPb != nil {
		pb.Notifications = notificationsPb
	}

	pb.OutputSchemaName = st.OutputSchemaName

	schedulePb, err := monitorCronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}

	pb.SkipBuiltinDashboard = st.SkipBuiltinDashboard

	pb.SlicingExprs = st.SlicingExprs

	snapshotPb, err := monitorSnapshotToPb(st.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotPb != nil {
		pb.Snapshot = snapshotPb
	}

	pb.TableName = st.TableName

	timeSeriesPb, err := monitorTimeSeriesToPb(st.TimeSeries)
	if err != nil {
		return nil, err
	}
	if timeSeriesPb != nil {
		pb.TimeSeries = timeSeriesPb
	}

	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateMonitor) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createMonitorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createMonitorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateMonitor) MarshalJSON() ([]byte, error) {
	pb, err := createMonitorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createMonitorPb struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	AssetsDir string `json:"assets_dir"`
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName string `json:"baseline_table_name,omitempty"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics []monitorMetricPb `json:"custom_metrics,omitempty"`
	// The data classification config for the monitor.
	DataClassificationConfig *monitorDataClassificationConfigPb `json:"data_classification_config,omitempty"`
	// Configuration for monitoring inference logs.
	InferenceLog *monitorInferenceLogPb `json:"inference_log,omitempty"`
	// The notification settings for the monitor.
	Notifications *monitorNotificationsPb `json:"notifications,omitempty"`
	// Schema where output metric tables are created.
	OutputSchemaName string `json:"output_schema_name"`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule *monitorCronSchedulePb `json:"schedule,omitempty"`
	// Whether to skip creating a default dashboard summarizing data quality
	// metrics.
	SkipBuiltinDashboard bool `json:"skip_builtin_dashboard,omitempty"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs []string `json:"slicing_exprs,omitempty"`
	// Configuration for monitoring snapshot tables.
	Snapshot *monitorSnapshotPb `json:"snapshot,omitempty"`
	// Full name of the table.
	TableName string `json:"-" url:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries *monitorTimeSeriesPb `json:"time_series,omitempty"`
	// Optional argument to specify the warehouse for dashboard creation. If not
	// specified, the first running warehouse will be used.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createMonitorFromPb(pb *createMonitorPb) (*CreateMonitor, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateMonitor{}
	st.AssetsDir = pb.AssetsDir
	st.BaselineTableName = pb.BaselineTableName

	var customMetricsField []MonitorMetric
	for _, itemPb := range pb.CustomMetrics {
		item, err := monitorMetricFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customMetricsField = append(customMetricsField, *item)
		}
	}
	st.CustomMetrics = customMetricsField
	dataClassificationConfigField, err := monitorDataClassificationConfigFromPb(pb.DataClassificationConfig)
	if err != nil {
		return nil, err
	}
	if dataClassificationConfigField != nil {
		st.DataClassificationConfig = dataClassificationConfigField
	}
	inferenceLogField, err := monitorInferenceLogFromPb(pb.InferenceLog)
	if err != nil {
		return nil, err
	}
	if inferenceLogField != nil {
		st.InferenceLog = inferenceLogField
	}
	notificationsField, err := monitorNotificationsFromPb(pb.Notifications)
	if err != nil {
		return nil, err
	}
	if notificationsField != nil {
		st.Notifications = notificationsField
	}
	st.OutputSchemaName = pb.OutputSchemaName
	scheduleField, err := monitorCronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	st.SkipBuiltinDashboard = pb.SkipBuiltinDashboard
	st.SlicingExprs = pb.SlicingExprs
	snapshotField, err := monitorSnapshotFromPb(pb.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotField != nil {
		st.Snapshot = snapshotField
	}
	st.TableName = pb.TableName
	timeSeriesField, err := monitorTimeSeriesFromPb(pb.TimeSeries)
	if err != nil {
		return nil, err
	}
	if timeSeriesField != nil {
		st.TimeSeries = timeSeriesField
	}
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createMonitorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createMonitorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Create an Online Table
type CreateOnlineTableRequest struct {
	// Online Table information.
	// Wire name: 'table'
	Table OnlineTable
}

func createOnlineTableRequestToPb(st *CreateOnlineTableRequest) (*createOnlineTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createOnlineTableRequestPb{}
	tablePb, err := onlineTableToPb(&st.Table)
	if err != nil {
		return nil, err
	}
	if tablePb != nil {
		pb.Table = *tablePb
	}

	return pb, nil
}

func (st *CreateOnlineTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createOnlineTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createOnlineTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateOnlineTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := createOnlineTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createOnlineTableRequestPb struct {
	// Online Table information.
	Table onlineTablePb `json:"table"`
}

func createOnlineTableRequestFromPb(pb *createOnlineTableRequestPb) (*CreateOnlineTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateOnlineTableRequest{}
	tableField, err := onlineTableFromPb(&pb.Table)
	if err != nil {
		return nil, err
	}
	if tableField != nil {
		st.Table = *tableField
	}

	return st, nil
}

type CreateRegisteredModelRequest struct {
	// The name of the catalog where the schema and the registered model reside
	// Wire name: 'catalog_name'
	CatalogName string
	// The comment attached to the registered model
	// Wire name: 'comment'
	Comment string
	// The name of the registered model
	// Wire name: 'name'
	Name string
	// The name of the schema where the registered model resides
	// Wire name: 'schema_name'
	SchemaName string
	// The storage location on the cloud under which model version data files
	// are stored
	// Wire name: 'storage_location'
	StorageLocation string

	ForceSendFields []string `tf:"-"`
}

func createRegisteredModelRequestToPb(st *CreateRegisteredModelRequest) (*createRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createRegisteredModelRequestPb{}
	pb.CatalogName = st.CatalogName

	pb.Comment = st.Comment

	pb.Name = st.Name

	pb.SchemaName = st.SchemaName

	pb.StorageLocation = st.StorageLocation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createRegisteredModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createRegisteredModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := createRegisteredModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createRegisteredModelRequestPb struct {
	// The name of the catalog where the schema and the registered model reside
	CatalogName string `json:"catalog_name"`
	// The comment attached to the registered model
	Comment string `json:"comment,omitempty"`
	// The name of the registered model
	Name string `json:"name"`
	// The name of the schema where the registered model resides
	SchemaName string `json:"schema_name"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string `json:"storage_location,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createRegisteredModelRequestFromPb(pb *createRegisteredModelRequestPb) (*CreateRegisteredModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRegisteredModelRequest{}
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.SchemaName = pb.SchemaName
	st.StorageLocation = pb.StorageLocation

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createRegisteredModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createRegisteredModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateResponse struct {
}

func createResponseToPb(st *CreateResponse) (*createResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createResponsePb{}

	return pb, nil
}

func (st *CreateResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateResponse) MarshalJSON() ([]byte, error) {
	pb, err := createResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createResponsePb struct {
}

func createResponseFromPb(pb *createResponsePb) (*CreateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateResponse{}

	return st, nil
}

type CreateSchema struct {
	// Name of parent catalog.
	// Wire name: 'catalog_name'
	CatalogName string
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// Name of schema, relative to parent catalog.
	// Wire name: 'name'
	Name string
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string
	// Storage root URL for managed tables within schema.
	// Wire name: 'storage_root'
	StorageRoot string

	ForceSendFields []string `tf:"-"`
}

func createSchemaToPb(st *CreateSchema) (*createSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createSchemaPb{}
	pb.CatalogName = st.CatalogName

	pb.Comment = st.Comment

	pb.Name = st.Name

	pb.Properties = st.Properties

	pb.StorageRoot = st.StorageRoot

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateSchema) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createSchemaPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createSchemaFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateSchema) MarshalJSON() ([]byte, error) {
	pb, err := createSchemaToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createSchemaPb struct {
	// Name of parent catalog.
	CatalogName string `json:"catalog_name"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of schema, relative to parent catalog.
	Name string `json:"name"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// Storage root URL for managed tables within schema.
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createSchemaFromPb(pb *createSchemaPb) (*CreateSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateSchema{}
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Properties = pb.Properties
	st.StorageRoot = pb.StorageRoot

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createSchemaPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createSchemaPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateStorageCredential struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRoleRequest
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentityRequest
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal
	// The Cloudflare API token configuration.
	// Wire name: 'cloudflare_api_token'
	CloudflareApiToken *CloudflareApiToken
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string
	// The Databricks managed GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest
	// The credential name. The name must be unique within the metastore.
	// Wire name: 'name'
	Name string
	// Whether the storage credential is only usable for read operations.
	// Wire name: 'read_only'
	ReadOnly bool
	// Supplying true to this argument skips validation of the created
	// credential.
	// Wire name: 'skip_validation'
	SkipValidation bool

	ForceSendFields []string `tf:"-"`
}

func createStorageCredentialToPb(st *CreateStorageCredential) (*createStorageCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createStorageCredentialPb{}
	awsIamRolePb, err := awsIamRoleRequestToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}

	azureManagedIdentityPb, err := azureManagedIdentityRequestToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}

	azureServicePrincipalPb, err := azureServicePrincipalToPb(st.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalPb != nil {
		pb.AzureServicePrincipal = azureServicePrincipalPb
	}

	cloudflareApiTokenPb, err := cloudflareApiTokenToPb(st.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenPb != nil {
		pb.CloudflareApiToken = cloudflareApiTokenPb
	}

	pb.Comment = st.Comment

	databricksGcpServiceAccountPb, err := databricksGcpServiceAccountRequestToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}

	pb.Name = st.Name

	pb.ReadOnly = st.ReadOnly

	pb.SkipValidation = st.SkipValidation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateStorageCredential) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createStorageCredentialPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createStorageCredentialFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateStorageCredential) MarshalJSON() ([]byte, error) {
	pb, err := createStorageCredentialToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createStorageCredentialPb struct {
	// The AWS IAM role configuration.
	AwsIamRole *awsIamRoleRequestPb `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *azureManagedIdentityRequestPb `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *azureServicePrincipalPb `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *cloudflareApiTokenPb `json:"cloudflare_api_token,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *databricksGcpServiceAccountRequestPb `json:"databricks_gcp_service_account,omitempty"`
	// The credential name. The name must be unique within the metastore.
	Name string `json:"name"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `json:"read_only,omitempty"`
	// Supplying true to this argument skips validation of the created
	// credential.
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createStorageCredentialFromPb(pb *createStorageCredentialPb) (*CreateStorageCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateStorageCredential{}
	awsIamRoleField, err := awsIamRoleRequestFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := azureManagedIdentityRequestFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	azureServicePrincipalField, err := azureServicePrincipalFromPb(pb.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalField != nil {
		st.AzureServicePrincipal = azureServicePrincipalField
	}
	cloudflareApiTokenField, err := cloudflareApiTokenFromPb(pb.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenField != nil {
		st.CloudflareApiToken = cloudflareApiTokenField
	}
	st.Comment = pb.Comment
	databricksGcpServiceAccountField, err := databricksGcpServiceAccountRequestFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	st.Name = pb.Name
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createStorageCredentialPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createStorageCredentialPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateTableConstraint struct {
	// A table constraint, as defined by *one* of the following fields being
	// set: __primary_key_constraint__, __foreign_key_constraint__,
	// __named_table_constraint__.
	// Wire name: 'constraint'
	Constraint TableConstraint
	// The full name of the table referenced by the constraint.
	// Wire name: 'full_name_arg'
	FullNameArg string
}

func createTableConstraintToPb(st *CreateTableConstraint) (*createTableConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createTableConstraintPb{}
	constraintPb, err := tableConstraintToPb(&st.Constraint)
	if err != nil {
		return nil, err
	}
	if constraintPb != nil {
		pb.Constraint = *constraintPb
	}

	pb.FullNameArg = st.FullNameArg

	return pb, nil
}

func (st *CreateTableConstraint) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createTableConstraintPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createTableConstraintFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateTableConstraint) MarshalJSON() ([]byte, error) {
	pb, err := createTableConstraintToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createTableConstraintPb struct {
	// A table constraint, as defined by *one* of the following fields being
	// set: __primary_key_constraint__, __foreign_key_constraint__,
	// __named_table_constraint__.
	Constraint tableConstraintPb `json:"constraint"`
	// The full name of the table referenced by the constraint.
	FullNameArg string `json:"full_name_arg"`
}

func createTableConstraintFromPb(pb *createTableConstraintPb) (*CreateTableConstraint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateTableConstraint{}
	constraintField, err := tableConstraintFromPb(&pb.Constraint)
	if err != nil {
		return nil, err
	}
	if constraintField != nil {
		st.Constraint = *constraintField
	}
	st.FullNameArg = pb.FullNameArg

	return st, nil
}

type CreateVolumeRequestContent struct {
	// The name of the catalog where the schema and the volume are
	// Wire name: 'catalog_name'
	CatalogName string
	// The comment attached to the volume
	// Wire name: 'comment'
	Comment string
	// The name of the volume
	// Wire name: 'name'
	Name string
	// The name of the schema where the volume is
	// Wire name: 'schema_name'
	SchemaName string
	// The storage location on the cloud
	// Wire name: 'storage_location'
	StorageLocation string
	// The type of the volume. An external volume is located in the specified
	// external location. A managed volume is located in the default location
	// which is specified by the parent schema, or the parent catalog, or the
	// Metastore. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/aws/en/volumes/managed-vs-external
	// Wire name: 'volume_type'
	VolumeType VolumeType

	ForceSendFields []string `tf:"-"`
}

func createVolumeRequestContentToPb(st *CreateVolumeRequestContent) (*createVolumeRequestContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createVolumeRequestContentPb{}
	pb.CatalogName = st.CatalogName

	pb.Comment = st.Comment

	pb.Name = st.Name

	pb.SchemaName = st.SchemaName

	pb.StorageLocation = st.StorageLocation

	pb.VolumeType = st.VolumeType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateVolumeRequestContent) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createVolumeRequestContentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createVolumeRequestContentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateVolumeRequestContent) MarshalJSON() ([]byte, error) {
	pb, err := createVolumeRequestContentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createVolumeRequestContentPb struct {
	// The name of the catalog where the schema and the volume are
	CatalogName string `json:"catalog_name"`
	// The comment attached to the volume
	Comment string `json:"comment,omitempty"`
	// The name of the volume
	Name string `json:"name"`
	// The name of the schema where the volume is
	SchemaName string `json:"schema_name"`
	// The storage location on the cloud
	StorageLocation string `json:"storage_location,omitempty"`
	// The type of the volume. An external volume is located in the specified
	// external location. A managed volume is located in the default location
	// which is specified by the parent schema, or the parent catalog, or the
	// Metastore. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/aws/en/volumes/managed-vs-external
	VolumeType VolumeType `json:"volume_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createVolumeRequestContentFromPb(pb *createVolumeRequestContentPb) (*CreateVolumeRequestContent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateVolumeRequestContent{}
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.SchemaName = pb.SchemaName
	st.StorageLocation = pb.StorageLocation
	st.VolumeType = pb.VolumeType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createVolumeRequestContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createVolumeRequestContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CredentialInfo struct {
	// The AWS IAM role configuration
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRole
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentity
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string
	// Time at which this credential was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of credential creator.
	// Wire name: 'created_by'
	CreatedBy string
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount
	// The full name of the credential.
	// Wire name: 'full_name'
	FullName string
	// The unique identifier of the credential.
	// Wire name: 'id'
	Id string
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode
	// Unique identifier of the parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	// Wire name: 'name'
	Name string
	// Username of current owner of credential.
	// Wire name: 'owner'
	Owner string
	// Indicates the purpose of the credential.
	// Wire name: 'purpose'
	Purpose CredentialPurpose
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	// Wire name: 'read_only'
	ReadOnly bool
	// Time at which this credential was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// Username of user who last modified the credential.
	// Wire name: 'updated_by'
	UpdatedBy string
	// Whether this credential is the current metastore's root storage
	// credential. Only applicable when purpose is **STORAGE**.
	// Wire name: 'used_for_managed_storage'
	UsedForManagedStorage bool

	ForceSendFields []string `tf:"-"`
}

func credentialInfoToPb(st *CredentialInfo) (*credentialInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &credentialInfoPb{}
	awsIamRolePb, err := awsIamRoleToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}

	azureManagedIdentityPb, err := azureManagedIdentityToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}

	azureServicePrincipalPb, err := azureServicePrincipalToPb(st.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalPb != nil {
		pb.AzureServicePrincipal = azureServicePrincipalPb
	}

	pb.Comment = st.Comment

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	databricksGcpServiceAccountPb, err := databricksGcpServiceAccountToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}

	pb.FullName = st.FullName

	pb.Id = st.Id

	pb.IsolationMode = st.IsolationMode

	pb.MetastoreId = st.MetastoreId

	pb.Name = st.Name

	pb.Owner = st.Owner

	pb.Purpose = st.Purpose

	pb.ReadOnly = st.ReadOnly

	pb.UpdatedAt = st.UpdatedAt

	pb.UpdatedBy = st.UpdatedBy

	pb.UsedForManagedStorage = st.UsedForManagedStorage

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CredentialInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &credentialInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := credentialInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CredentialInfo) MarshalJSON() ([]byte, error) {
	pb, err := credentialInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type credentialInfoPb struct {
	// The AWS IAM role configuration
	AwsIamRole *awsIamRolePb `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *azureManagedIdentityPb `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	AzureServicePrincipal *azureServicePrincipalPb `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// Time at which this credential was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount *databricksGcpServiceAccountPb `json:"databricks_gcp_service_account,omitempty"`
	// The full name of the credential.
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the credential.
	Id string `json:"id,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Unique identifier of the parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	Name string `json:"name,omitempty"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Indicates the purpose of the credential.
	Purpose CredentialPurpose `json:"purpose,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool `json:"read_only,omitempty"`
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the credential.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Whether this credential is the current metastore's root storage
	// credential. Only applicable when purpose is **STORAGE**.
	UsedForManagedStorage bool `json:"used_for_managed_storage,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func credentialInfoFromPb(pb *credentialInfoPb) (*CredentialInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CredentialInfo{}
	awsIamRoleField, err := awsIamRoleFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := azureManagedIdentityFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	azureServicePrincipalField, err := azureServicePrincipalFromPb(pb.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalField != nil {
		st.AzureServicePrincipal = azureServicePrincipalField
	}
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	databricksGcpServiceAccountField, err := databricksGcpServiceAccountFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	st.FullName = pb.FullName
	st.Id = pb.Id
	st.IsolationMode = pb.IsolationMode
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.Purpose = pb.Purpose
	st.ReadOnly = pb.ReadOnly
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.UsedForManagedStorage = pb.UsedForManagedStorage

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *credentialInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st credentialInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CredentialPurpose string
type credentialPurposePb string

const CredentialPurposeService CredentialPurpose = `SERVICE`

const CredentialPurposeStorage CredentialPurpose = `STORAGE`

// String representation for [fmt.Print]
func (f *CredentialPurpose) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CredentialPurpose) Set(v string) error {
	switch v {
	case `SERVICE`, `STORAGE`:
		*f = CredentialPurpose(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SERVICE", "STORAGE"`, v)
	}
}

// Type always returns CredentialPurpose to satisfy [pflag.Value] interface
func (f *CredentialPurpose) Type() string {
	return "CredentialPurpose"
}

func credentialPurposeToPb(st *CredentialPurpose) (*credentialPurposePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := credentialPurposePb(*st)
	return &pb, nil
}

func credentialPurposeFromPb(pb *credentialPurposePb) (*CredentialPurpose, error) {
	if pb == nil {
		return nil, nil
	}
	st := CredentialPurpose(*pb)
	return &st, nil
}

// The type of credential.
type CredentialType string
type credentialTypePb string

const CredentialTypeBearerToken CredentialType = `BEARER_TOKEN`

const CredentialTypeUsernamePassword CredentialType = `USERNAME_PASSWORD`

// String representation for [fmt.Print]
func (f *CredentialType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CredentialType) Set(v string) error {
	switch v {
	case `BEARER_TOKEN`, `USERNAME_PASSWORD`:
		*f = CredentialType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BEARER_TOKEN", "USERNAME_PASSWORD"`, v)
	}
}

// Type always returns CredentialType to satisfy [pflag.Value] interface
func (f *CredentialType) Type() string {
	return "CredentialType"
}

func credentialTypeToPb(st *CredentialType) (*credentialTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := credentialTypePb(*st)
	return &pb, nil
}

func credentialTypeFromPb(pb *credentialTypePb) (*CredentialType, error) {
	if pb == nil {
		return nil, nil
	}
	st := CredentialType(*pb)
	return &st, nil
}

type CredentialValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	// Wire name: 'message'
	Message string
	// The results of the tested operation.
	// Wire name: 'result'
	Result ValidateCredentialResult

	ForceSendFields []string `tf:"-"`
}

func credentialValidationResultToPb(st *CredentialValidationResult) (*credentialValidationResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &credentialValidationResultPb{}
	pb.Message = st.Message

	pb.Result = st.Result

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CredentialValidationResult) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &credentialValidationResultPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := credentialValidationResultFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CredentialValidationResult) MarshalJSON() ([]byte, error) {
	pb, err := credentialValidationResultToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type credentialValidationResultPb struct {
	// Error message would exist when the result does not equal to **PASS**.
	Message string `json:"message,omitempty"`
	// The results of the tested operation.
	Result ValidateCredentialResult `json:"result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func credentialValidationResultFromPb(pb *credentialValidationResultPb) (*CredentialValidationResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CredentialValidationResult{}
	st.Message = pb.Message
	st.Result = pb.Result

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *credentialValidationResultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st credentialValidationResultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Currently assigned workspaces
type CurrentWorkspaceBindings struct {
	// A list of workspace IDs.
	// Wire name: 'workspaces'
	Workspaces []int64
}

func currentWorkspaceBindingsToPb(st *CurrentWorkspaceBindings) (*currentWorkspaceBindingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &currentWorkspaceBindingsPb{}
	pb.Workspaces = st.Workspaces

	return pb, nil
}

func (st *CurrentWorkspaceBindings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &currentWorkspaceBindingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := currentWorkspaceBindingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CurrentWorkspaceBindings) MarshalJSON() ([]byte, error) {
	pb, err := currentWorkspaceBindingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type currentWorkspaceBindingsPb struct {
	// A list of workspace IDs.
	Workspaces []int64 `json:"workspaces,omitempty"`
}

func currentWorkspaceBindingsFromPb(pb *currentWorkspaceBindingsPb) (*CurrentWorkspaceBindings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CurrentWorkspaceBindings{}
	st.Workspaces = pb.Workspaces

	return st, nil
}

// Data source format
type DataSourceFormat string
type dataSourceFormatPb string

const DataSourceFormatAvro DataSourceFormat = `AVRO`

const DataSourceFormatBigqueryFormat DataSourceFormat = `BIGQUERY_FORMAT`

const DataSourceFormatCsv DataSourceFormat = `CSV`

const DataSourceFormatDatabricksFormat DataSourceFormat = `DATABRICKS_FORMAT`

const DataSourceFormatDelta DataSourceFormat = `DELTA`

const DataSourceFormatDeltasharing DataSourceFormat = `DELTASHARING`

const DataSourceFormatHiveCustom DataSourceFormat = `HIVE_CUSTOM`

const DataSourceFormatHiveSerde DataSourceFormat = `HIVE_SERDE`

const DataSourceFormatJson DataSourceFormat = `JSON`

const DataSourceFormatMysqlFormat DataSourceFormat = `MYSQL_FORMAT`

const DataSourceFormatNetsuiteFormat DataSourceFormat = `NETSUITE_FORMAT`

const DataSourceFormatOrc DataSourceFormat = `ORC`

const DataSourceFormatParquet DataSourceFormat = `PARQUET`

const DataSourceFormatPostgresqlFormat DataSourceFormat = `POSTGRESQL_FORMAT`

const DataSourceFormatRedshiftFormat DataSourceFormat = `REDSHIFT_FORMAT`

const DataSourceFormatSalesforceFormat DataSourceFormat = `SALESFORCE_FORMAT`

const DataSourceFormatSnowflakeFormat DataSourceFormat = `SNOWFLAKE_FORMAT`

const DataSourceFormatSqldwFormat DataSourceFormat = `SQLDW_FORMAT`

const DataSourceFormatSqlserverFormat DataSourceFormat = `SQLSERVER_FORMAT`

const DataSourceFormatText DataSourceFormat = `TEXT`

const DataSourceFormatUnityCatalog DataSourceFormat = `UNITY_CATALOG`

const DataSourceFormatVectorIndexFormat DataSourceFormat = `VECTOR_INDEX_FORMAT`

const DataSourceFormatWorkdayRaasFormat DataSourceFormat = `WORKDAY_RAAS_FORMAT`

// String representation for [fmt.Print]
func (f *DataSourceFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DataSourceFormat) Set(v string) error {
	switch v {
	case `AVRO`, `BIGQUERY_FORMAT`, `CSV`, `DATABRICKS_FORMAT`, `DELTA`, `DELTASHARING`, `HIVE_CUSTOM`, `HIVE_SERDE`, `JSON`, `MYSQL_FORMAT`, `NETSUITE_FORMAT`, `ORC`, `PARQUET`, `POSTGRESQL_FORMAT`, `REDSHIFT_FORMAT`, `SALESFORCE_FORMAT`, `SNOWFLAKE_FORMAT`, `SQLDW_FORMAT`, `SQLSERVER_FORMAT`, `TEXT`, `UNITY_CATALOG`, `VECTOR_INDEX_FORMAT`, `WORKDAY_RAAS_FORMAT`:
		*f = DataSourceFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVRO", "BIGQUERY_FORMAT", "CSV", "DATABRICKS_FORMAT", "DELTA", "DELTASHARING", "HIVE_CUSTOM", "HIVE_SERDE", "JSON", "MYSQL_FORMAT", "NETSUITE_FORMAT", "ORC", "PARQUET", "POSTGRESQL_FORMAT", "REDSHIFT_FORMAT", "SALESFORCE_FORMAT", "SNOWFLAKE_FORMAT", "SQLDW_FORMAT", "SQLSERVER_FORMAT", "TEXT", "UNITY_CATALOG", "VECTOR_INDEX_FORMAT", "WORKDAY_RAAS_FORMAT"`, v)
	}
}

// Type always returns DataSourceFormat to satisfy [pflag.Value] interface
func (f *DataSourceFormat) Type() string {
	return "DataSourceFormat"
}

func dataSourceFormatToPb(st *DataSourceFormat) (*dataSourceFormatPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := dataSourceFormatPb(*st)
	return &pb, nil
}

func dataSourceFormatFromPb(pb *dataSourceFormatPb) (*DataSourceFormat, error) {
	if pb == nil {
		return nil, nil
	}
	st := DataSourceFormat(*pb)
	return &st, nil
}

// GCP long-lived credential. Databricks-created Google Cloud Storage service
// account.
type DatabricksGcpServiceAccount struct {
	// The Databricks internal ID that represents this managed identity. This
	// field is only used to persist the credential_id once it is fetched from
	// the credentials manager - as we only use the protobuf serializer to store
	// credentials, this ID gets persisted to the database
	// Wire name: 'credential_id'
	CredentialId string
	// The email of the service account.
	// Wire name: 'email'
	Email string
	// The ID that represents the private key for this Service Account
	// Wire name: 'private_key_id'
	PrivateKeyId string

	ForceSendFields []string `tf:"-"`
}

func databricksGcpServiceAccountToPb(st *DatabricksGcpServiceAccount) (*databricksGcpServiceAccountPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databricksGcpServiceAccountPb{}
	pb.CredentialId = st.CredentialId

	pb.Email = st.Email

	pb.PrivateKeyId = st.PrivateKeyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DatabricksGcpServiceAccount) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databricksGcpServiceAccountPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := databricksGcpServiceAccountFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DatabricksGcpServiceAccount) MarshalJSON() ([]byte, error) {
	pb, err := databricksGcpServiceAccountToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type databricksGcpServiceAccountPb struct {
	// The Databricks internal ID that represents this managed identity. This
	// field is only used to persist the credential_id once it is fetched from
	// the credentials manager - as we only use the protobuf serializer to store
	// credentials, this ID gets persisted to the database
	CredentialId string `json:"credential_id,omitempty"`
	// The email of the service account.
	Email string `json:"email,omitempty"`
	// The ID that represents the private key for this Service Account
	PrivateKeyId string `json:"private_key_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func databricksGcpServiceAccountFromPb(pb *databricksGcpServiceAccountPb) (*DatabricksGcpServiceAccount, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabricksGcpServiceAccount{}
	st.CredentialId = pb.CredentialId
	st.Email = pb.Email
	st.PrivateKeyId = pb.PrivateKeyId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *databricksGcpServiceAccountPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st databricksGcpServiceAccountPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DatabricksGcpServiceAccountRequest struct {
}

func databricksGcpServiceAccountRequestToPb(st *DatabricksGcpServiceAccountRequest) (*databricksGcpServiceAccountRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databricksGcpServiceAccountRequestPb{}

	return pb, nil
}

func (st *DatabricksGcpServiceAccountRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databricksGcpServiceAccountRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := databricksGcpServiceAccountRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DatabricksGcpServiceAccountRequest) MarshalJSON() ([]byte, error) {
	pb, err := databricksGcpServiceAccountRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type databricksGcpServiceAccountRequestPb struct {
}

func databricksGcpServiceAccountRequestFromPb(pb *databricksGcpServiceAccountRequestPb) (*DatabricksGcpServiceAccountRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabricksGcpServiceAccountRequest{}

	return st, nil
}

type DatabricksGcpServiceAccountResponse struct {
	// The Databricks internal ID that represents this service account. This is
	// an output-only field.
	// Wire name: 'credential_id'
	CredentialId string
	// The email of the service account. This is an output-only field.
	// Wire name: 'email'
	Email string

	ForceSendFields []string `tf:"-"`
}

func databricksGcpServiceAccountResponseToPb(st *DatabricksGcpServiceAccountResponse) (*databricksGcpServiceAccountResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databricksGcpServiceAccountResponsePb{}
	pb.CredentialId = st.CredentialId

	pb.Email = st.Email

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DatabricksGcpServiceAccountResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databricksGcpServiceAccountResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := databricksGcpServiceAccountResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DatabricksGcpServiceAccountResponse) MarshalJSON() ([]byte, error) {
	pb, err := databricksGcpServiceAccountResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type databricksGcpServiceAccountResponsePb struct {
	// The Databricks internal ID that represents this service account. This is
	// an output-only field.
	CredentialId string `json:"credential_id,omitempty"`
	// The email of the service account. This is an output-only field.
	Email string `json:"email,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func databricksGcpServiceAccountResponseFromPb(pb *databricksGcpServiceAccountResponsePb) (*DatabricksGcpServiceAccountResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabricksGcpServiceAccountResponse{}
	st.CredentialId = pb.CredentialId
	st.Email = pb.Email

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *databricksGcpServiceAccountResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st databricksGcpServiceAccountResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete a metastore assignment
type DeleteAccountMetastoreAssignmentRequest struct {
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func deleteAccountMetastoreAssignmentRequestToPb(st *DeleteAccountMetastoreAssignmentRequest) (*deleteAccountMetastoreAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountMetastoreAssignmentRequestPb{}
	pb.MetastoreId = st.MetastoreId

	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func (st *DeleteAccountMetastoreAssignmentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAccountMetastoreAssignmentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAccountMetastoreAssignmentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAccountMetastoreAssignmentRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteAccountMetastoreAssignmentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteAccountMetastoreAssignmentRequestPb struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func deleteAccountMetastoreAssignmentRequestFromPb(pb *deleteAccountMetastoreAssignmentRequestPb) (*DeleteAccountMetastoreAssignmentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountMetastoreAssignmentRequest{}
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

// Delete a metastore
type DeleteAccountMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	// Wire name: 'force'
	Force bool `tf:"-"`
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func deleteAccountMetastoreRequestToPb(st *DeleteAccountMetastoreRequest) (*deleteAccountMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountMetastoreRequestPb{}
	pb.Force = st.Force

	pb.MetastoreId = st.MetastoreId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeleteAccountMetastoreRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAccountMetastoreRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAccountMetastoreRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAccountMetastoreRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteAccountMetastoreRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteAccountMetastoreRequestPb struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force bool `json:"-" url:"force,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteAccountMetastoreRequestFromPb(pb *deleteAccountMetastoreRequestPb) (*DeleteAccountMetastoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountMetastoreRequest{}
	st.Force = pb.Force
	st.MetastoreId = pb.MetastoreId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteAccountMetastoreRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteAccountMetastoreRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete a storage credential
type DeleteAccountStorageCredentialRequest struct {
	// Force deletion even if the Storage Credential is not empty. Default is
	// false.
	// Wire name: 'force'
	Force bool `tf:"-"`
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Name of the storage credential.
	// Wire name: 'storage_credential_name'
	StorageCredentialName string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func deleteAccountStorageCredentialRequestToPb(st *DeleteAccountStorageCredentialRequest) (*deleteAccountStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountStorageCredentialRequestPb{}
	pb.Force = st.Force

	pb.MetastoreId = st.MetastoreId

	pb.StorageCredentialName = st.StorageCredentialName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeleteAccountStorageCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAccountStorageCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAccountStorageCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAccountStorageCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteAccountStorageCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteAccountStorageCredentialRequestPb struct {
	// Force deletion even if the Storage Credential is not empty. Default is
	// false.
	Force bool `json:"-" url:"force,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Name of the storage credential.
	StorageCredentialName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteAccountStorageCredentialRequestFromPb(pb *deleteAccountStorageCredentialRequestPb) (*DeleteAccountStorageCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountStorageCredentialRequest{}
	st.Force = pb.Force
	st.MetastoreId = pb.MetastoreId
	st.StorageCredentialName = pb.StorageCredentialName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteAccountStorageCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteAccountStorageCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete a Registered Model Alias
type DeleteAliasRequest struct {
	// The name of the alias
	// Wire name: 'alias'
	Alias string `tf:"-"`
	// The three-level (fully qualified) name of the registered model
	// Wire name: 'full_name'
	FullName string `tf:"-"`
}

func deleteAliasRequestToPb(st *DeleteAliasRequest) (*deleteAliasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAliasRequestPb{}
	pb.Alias = st.Alias

	pb.FullName = st.FullName

	return pb, nil
}

func (st *DeleteAliasRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAliasRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAliasRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAliasRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteAliasRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteAliasRequestPb struct {
	// The name of the alias
	Alias string `json:"-" url:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
}

func deleteAliasRequestFromPb(pb *deleteAliasRequestPb) (*DeleteAliasRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAliasRequest{}
	st.Alias = pb.Alias
	st.FullName = pb.FullName

	return st, nil
}

type DeleteAliasResponse struct {
}

func deleteAliasResponseToPb(st *DeleteAliasResponse) (*deleteAliasResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAliasResponsePb{}

	return pb, nil
}

func (st *DeleteAliasResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAliasResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAliasResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAliasResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteAliasResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteAliasResponsePb struct {
}

func deleteAliasResponseFromPb(pb *deleteAliasResponsePb) (*DeleteAliasResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAliasResponse{}

	return st, nil
}

// Delete a catalog
type DeleteCatalogRequest struct {
	// Force deletion even if the catalog is not empty.
	// Wire name: 'force'
	Force bool `tf:"-"`
	// The name of the catalog.
	// Wire name: 'name'
	Name string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func deleteCatalogRequestToPb(st *DeleteCatalogRequest) (*deleteCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCatalogRequestPb{}
	pb.Force = st.Force

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeleteCatalogRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCatalogRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCatalogRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCatalogRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteCatalogRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteCatalogRequestPb struct {
	// Force deletion even if the catalog is not empty.
	Force bool `json:"-" url:"force,omitempty"`
	// The name of the catalog.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteCatalogRequestFromPb(pb *deleteCatalogRequestPb) (*DeleteCatalogRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCatalogRequest{}
	st.Force = pb.Force
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteCatalogRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteCatalogRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete a connection
type DeleteConnectionRequest struct {
	// The name of the connection to be deleted.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func deleteConnectionRequestToPb(st *DeleteConnectionRequest) (*deleteConnectionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteConnectionRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *DeleteConnectionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteConnectionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteConnectionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteConnectionRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteConnectionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteConnectionRequestPb struct {
	// The name of the connection to be deleted.
	Name string `json:"-" url:"-"`
}

func deleteConnectionRequestFromPb(pb *deleteConnectionRequestPb) (*DeleteConnectionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteConnectionRequest{}
	st.Name = pb.Name

	return st, nil
}

// Delete a credential
type DeleteCredentialRequest struct {
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	// Wire name: 'force'
	Force bool `tf:"-"`
	// Name of the credential.
	// Wire name: 'name_arg'
	NameArg string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func deleteCredentialRequestToPb(st *DeleteCredentialRequest) (*deleteCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCredentialRequestPb{}
	pb.Force = st.Force

	pb.NameArg = st.NameArg

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeleteCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteCredentialRequestPb struct {
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	Force bool `json:"-" url:"force,omitempty"`
	// Name of the credential.
	NameArg string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteCredentialRequestFromPb(pb *deleteCredentialRequestPb) (*DeleteCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCredentialRequest{}
	st.Force = pb.Force
	st.NameArg = pb.NameArg

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteCredentialResponse struct {
}

func deleteCredentialResponseToPb(st *DeleteCredentialResponse) (*deleteCredentialResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCredentialResponsePb{}

	return pb, nil
}

func (st *DeleteCredentialResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCredentialResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCredentialResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCredentialResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteCredentialResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteCredentialResponsePb struct {
}

func deleteCredentialResponseFromPb(pb *deleteCredentialResponsePb) (*DeleteCredentialResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCredentialResponse{}

	return st, nil
}

// Delete an external location
type DeleteExternalLocationRequest struct {
	// Force deletion even if there are dependent external tables or mounts.
	// Wire name: 'force'
	Force bool `tf:"-"`
	// Name of the external location.
	// Wire name: 'name'
	Name string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func deleteExternalLocationRequestToPb(st *DeleteExternalLocationRequest) (*deleteExternalLocationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteExternalLocationRequestPb{}
	pb.Force = st.Force

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeleteExternalLocationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteExternalLocationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteExternalLocationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteExternalLocationRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteExternalLocationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteExternalLocationRequestPb struct {
	// Force deletion even if there are dependent external tables or mounts.
	Force bool `json:"-" url:"force,omitempty"`
	// Name of the external location.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteExternalLocationRequestFromPb(pb *deleteExternalLocationRequestPb) (*DeleteExternalLocationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExternalLocationRequest{}
	st.Force = pb.Force
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteExternalLocationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteExternalLocationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete a function
type DeleteFunctionRequest struct {
	// Force deletion even if the function is notempty.
	// Wire name: 'force'
	Force bool `tf:"-"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	// Wire name: 'name'
	Name string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func deleteFunctionRequestToPb(st *DeleteFunctionRequest) (*deleteFunctionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteFunctionRequestPb{}
	pb.Force = st.Force

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeleteFunctionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteFunctionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteFunctionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteFunctionRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteFunctionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteFunctionRequestPb struct {
	// Force deletion even if the function is notempty.
	Force bool `json:"-" url:"force,omitempty"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteFunctionRequestFromPb(pb *deleteFunctionRequestPb) (*DeleteFunctionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteFunctionRequest{}
	st.Force = pb.Force
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteFunctionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteFunctionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete a metastore
type DeleteMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	// Wire name: 'force'
	Force bool `tf:"-"`
	// Unique ID of the metastore.
	// Wire name: 'id'
	Id string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func deleteMetastoreRequestToPb(st *DeleteMetastoreRequest) (*deleteMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteMetastoreRequestPb{}
	pb.Force = st.Force

	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeleteMetastoreRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteMetastoreRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteMetastoreRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteMetastoreRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteMetastoreRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteMetastoreRequestPb struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force bool `json:"-" url:"force,omitempty"`
	// Unique ID of the metastore.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteMetastoreRequestFromPb(pb *deleteMetastoreRequestPb) (*DeleteMetastoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteMetastoreRequest{}
	st.Force = pb.Force
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteMetastoreRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteMetastoreRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete a Model Version
type DeleteModelVersionRequest struct {
	// The three-level (fully qualified) name of the model version
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// The integer version number of the model version
	// Wire name: 'version'
	Version int `tf:"-"`
}

func deleteModelVersionRequestToPb(st *DeleteModelVersionRequest) (*deleteModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelVersionRequestPb{}
	pb.FullName = st.FullName

	pb.Version = st.Version

	return pb, nil
}

func (st *DeleteModelVersionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelVersionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelVersionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelVersionRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelVersionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteModelVersionRequestPb struct {
	// The three-level (fully qualified) name of the model version
	FullName string `json:"-" url:"-"`
	// The integer version number of the model version
	Version int `json:"-" url:"-"`
}

func deleteModelVersionRequestFromPb(pb *deleteModelVersionRequestPb) (*DeleteModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelVersionRequest{}
	st.FullName = pb.FullName
	st.Version = pb.Version

	return st, nil
}

// Delete an Online Table
type DeleteOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func deleteOnlineTableRequestToPb(st *DeleteOnlineTableRequest) (*deleteOnlineTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteOnlineTableRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *DeleteOnlineTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteOnlineTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteOnlineTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteOnlineTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteOnlineTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteOnlineTableRequestPb struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"-" url:"-"`
}

func deleteOnlineTableRequestFromPb(pb *deleteOnlineTableRequestPb) (*DeleteOnlineTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteOnlineTableRequest{}
	st.Name = pb.Name

	return st, nil
}

// Delete a table monitor
type DeleteQualityMonitorRequest struct {
	// Full name of the table.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func deleteQualityMonitorRequestToPb(st *DeleteQualityMonitorRequest) (*deleteQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteQualityMonitorRequestPb{}
	pb.TableName = st.TableName

	return pb, nil
}

func (st *DeleteQualityMonitorRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteQualityMonitorRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteQualityMonitorRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteQualityMonitorRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteQualityMonitorRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteQualityMonitorRequestPb struct {
	// Full name of the table.
	TableName string `json:"-" url:"-"`
}

func deleteQualityMonitorRequestFromPb(pb *deleteQualityMonitorRequestPb) (*DeleteQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteQualityMonitorRequest{}
	st.TableName = pb.TableName

	return st, nil
}

// Delete a Registered Model
type DeleteRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	// Wire name: 'full_name'
	FullName string `tf:"-"`
}

func deleteRegisteredModelRequestToPb(st *DeleteRegisteredModelRequest) (*deleteRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRegisteredModelRequestPb{}
	pb.FullName = st.FullName

	return pb, nil
}

func (st *DeleteRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRegisteredModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRegisteredModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteRegisteredModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteRegisteredModelRequestPb struct {
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
}

func deleteRegisteredModelRequestFromPb(pb *deleteRegisteredModelRequestPb) (*DeleteRegisteredModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRegisteredModelRequest{}
	st.FullName = pb.FullName

	return st, nil
}

type DeleteResponse struct {
}

func deleteResponseToPb(st *DeleteResponse) (*deleteResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteResponsePb{}

	return pb, nil
}

func (st *DeleteResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteResponsePb struct {
}

func deleteResponseFromPb(pb *deleteResponsePb) (*DeleteResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteResponse{}

	return st, nil
}

// Delete a schema
type DeleteSchemaRequest struct {
	// Force deletion even if the schema is not empty.
	// Wire name: 'force'
	Force bool `tf:"-"`
	// Full name of the schema.
	// Wire name: 'full_name'
	FullName string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func deleteSchemaRequestToPb(st *DeleteSchemaRequest) (*deleteSchemaRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteSchemaRequestPb{}
	pb.Force = st.Force

	pb.FullName = st.FullName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeleteSchemaRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteSchemaRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteSchemaRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteSchemaRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteSchemaRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteSchemaRequestPb struct {
	// Force deletion even if the schema is not empty.
	Force bool `json:"-" url:"force,omitempty"`
	// Full name of the schema.
	FullName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteSchemaRequestFromPb(pb *deleteSchemaRequestPb) (*DeleteSchemaRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSchemaRequest{}
	st.Force = pb.Force
	st.FullName = pb.FullName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteSchemaRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteSchemaRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete a credential
type DeleteStorageCredentialRequest struct {
	// Force deletion even if there are dependent external locations or external
	// tables.
	// Wire name: 'force'
	Force bool `tf:"-"`
	// Name of the storage credential.
	// Wire name: 'name'
	Name string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func deleteStorageCredentialRequestToPb(st *DeleteStorageCredentialRequest) (*deleteStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteStorageCredentialRequestPb{}
	pb.Force = st.Force

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeleteStorageCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteStorageCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteStorageCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteStorageCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteStorageCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteStorageCredentialRequestPb struct {
	// Force deletion even if there are dependent external locations or external
	// tables.
	Force bool `json:"-" url:"force,omitempty"`
	// Name of the storage credential.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteStorageCredentialRequestFromPb(pb *deleteStorageCredentialRequestPb) (*DeleteStorageCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteStorageCredentialRequest{}
	st.Force = pb.Force
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteStorageCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteStorageCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete a table constraint
type DeleteTableConstraintRequest struct {
	// If true, try deleting all child constraints of the current constraint. If
	// false, reject this operation if the current constraint has any child
	// constraints.
	// Wire name: 'cascade'
	Cascade bool `tf:"-"`
	// The name of the constraint to delete.
	// Wire name: 'constraint_name'
	ConstraintName string `tf:"-"`
	// Full name of the table referenced by the constraint.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
}

func deleteTableConstraintRequestToPb(st *DeleteTableConstraintRequest) (*deleteTableConstraintRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTableConstraintRequestPb{}
	pb.Cascade = st.Cascade

	pb.ConstraintName = st.ConstraintName

	pb.FullName = st.FullName

	return pb, nil
}

func (st *DeleteTableConstraintRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteTableConstraintRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteTableConstraintRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteTableConstraintRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteTableConstraintRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteTableConstraintRequestPb struct {
	// If true, try deleting all child constraints of the current constraint. If
	// false, reject this operation if the current constraint has any child
	// constraints.
	Cascade bool `json:"-" url:"cascade"`
	// The name of the constraint to delete.
	ConstraintName string `json:"-" url:"constraint_name"`
	// Full name of the table referenced by the constraint.
	FullName string `json:"-" url:"-"`
}

func deleteTableConstraintRequestFromPb(pb *deleteTableConstraintRequestPb) (*DeleteTableConstraintRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTableConstraintRequest{}
	st.Cascade = pb.Cascade
	st.ConstraintName = pb.ConstraintName
	st.FullName = pb.FullName

	return st, nil
}

// Delete a table
type DeleteTableRequest struct {
	// Full name of the table.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
}

func deleteTableRequestToPb(st *DeleteTableRequest) (*deleteTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTableRequestPb{}
	pb.FullName = st.FullName

	return pb, nil
}

func (st *DeleteTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteTableRequestPb struct {
	// Full name of the table.
	FullName string `json:"-" url:"-"`
}

func deleteTableRequestFromPb(pb *deleteTableRequestPb) (*DeleteTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTableRequest{}
	st.FullName = pb.FullName

	return st, nil
}

// Delete a Volume
type DeleteVolumeRequest struct {
	// The three-level (fully qualified) name of the volume
	// Wire name: 'name'
	Name string `tf:"-"`
}

func deleteVolumeRequestToPb(st *DeleteVolumeRequest) (*deleteVolumeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteVolumeRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *DeleteVolumeRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteVolumeRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteVolumeRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteVolumeRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteVolumeRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteVolumeRequestPb struct {
	// The three-level (fully qualified) name of the volume
	Name string `json:"-" url:"-"`
}

func deleteVolumeRequestFromPb(pb *deleteVolumeRequestPb) (*DeleteVolumeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteVolumeRequest{}
	st.Name = pb.Name

	return st, nil
}

// Properties pertaining to the current state of the delta table as given by the
// commit server. This does not contain **delta.*** (input) properties in
// __TableInfo.properties__.
type DeltaRuntimePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	// Wire name: 'delta_runtime_properties'
	DeltaRuntimeProperties map[string]string
}

func deltaRuntimePropertiesKvPairsToPb(st *DeltaRuntimePropertiesKvPairs) (*deltaRuntimePropertiesKvPairsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaRuntimePropertiesKvPairsPb{}
	pb.DeltaRuntimeProperties = st.DeltaRuntimeProperties

	return pb, nil
}

func (st *DeltaRuntimePropertiesKvPairs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deltaRuntimePropertiesKvPairsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deltaRuntimePropertiesKvPairsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeltaRuntimePropertiesKvPairs) MarshalJSON() ([]byte, error) {
	pb, err := deltaRuntimePropertiesKvPairsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deltaRuntimePropertiesKvPairsPb struct {
	// A map of key-value properties attached to the securable.
	DeltaRuntimeProperties map[string]string `json:"delta_runtime_properties"`
}

func deltaRuntimePropertiesKvPairsFromPb(pb *deltaRuntimePropertiesKvPairsPb) (*DeltaRuntimePropertiesKvPairs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaRuntimePropertiesKvPairs{}
	st.DeltaRuntimeProperties = pb.DeltaRuntimeProperties

	return st, nil
}

// A dependency of a SQL object. Either the __table__ field or the __function__
// field must be defined.
type Dependency struct {
	// A function that is dependent on a SQL object.
	// Wire name: 'function'
	Function *FunctionDependency
	// A table that is dependent on a SQL object.
	// Wire name: 'table'
	Table *TableDependency
}

func dependencyToPb(st *Dependency) (*dependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dependencyPb{}
	functionPb, err := functionDependencyToPb(st.Function)
	if err != nil {
		return nil, err
	}
	if functionPb != nil {
		pb.Function = functionPb
	}

	tablePb, err := tableDependencyToPb(st.Table)
	if err != nil {
		return nil, err
	}
	if tablePb != nil {
		pb.Table = tablePb
	}

	return pb, nil
}

func (st *Dependency) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dependencyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dependencyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Dependency) MarshalJSON() ([]byte, error) {
	pb, err := dependencyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type dependencyPb struct {
	// A function that is dependent on a SQL object.
	Function *functionDependencyPb `json:"function,omitempty"`
	// A table that is dependent on a SQL object.
	Table *tableDependencyPb `json:"table,omitempty"`
}

func dependencyFromPb(pb *dependencyPb) (*Dependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Dependency{}
	functionField, err := functionDependencyFromPb(pb.Function)
	if err != nil {
		return nil, err
	}
	if functionField != nil {
		st.Function = functionField
	}
	tableField, err := tableDependencyFromPb(pb.Table)
	if err != nil {
		return nil, err
	}
	if tableField != nil {
		st.Table = tableField
	}

	return st, nil
}

// A list of dependencies.
type DependencyList struct {
	// Array of dependencies.
	// Wire name: 'dependencies'
	Dependencies []Dependency
}

func dependencyListToPb(st *DependencyList) (*dependencyListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dependencyListPb{}

	var dependenciesPb []dependencyPb
	for _, item := range st.Dependencies {
		itemPb, err := dependencyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dependenciesPb = append(dependenciesPb, *itemPb)
		}
	}
	pb.Dependencies = dependenciesPb

	return pb, nil
}

func (st *DependencyList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dependencyListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dependencyListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DependencyList) MarshalJSON() ([]byte, error) {
	pb, err := dependencyListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type dependencyListPb struct {
	// Array of dependencies.
	Dependencies []dependencyPb `json:"dependencies,omitempty"`
}

func dependencyListFromPb(pb *dependencyListPb) (*DependencyList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DependencyList{}

	var dependenciesField []Dependency
	for _, itemPb := range pb.Dependencies {
		item, err := dependencyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dependenciesField = append(dependenciesField, *item)
		}
	}
	st.Dependencies = dependenciesField

	return st, nil
}

// Disable a system schema
type DisableRequest struct {
	// The metastore ID under which the system schema lives.
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Full name of the system schema.
	// Wire name: 'schema_name'
	SchemaName string `tf:"-"`
}

func disableRequestToPb(st *DisableRequest) (*disableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &disableRequestPb{}
	pb.MetastoreId = st.MetastoreId

	pb.SchemaName = st.SchemaName

	return pb, nil
}

func (st *DisableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &disableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := disableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DisableRequest) MarshalJSON() ([]byte, error) {
	pb, err := disableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type disableRequestPb struct {
	// The metastore ID under which the system schema lives.
	MetastoreId string `json:"-" url:"-"`
	// Full name of the system schema.
	SchemaName string `json:"-" url:"-"`
}

func disableRequestFromPb(pb *disableRequestPb) (*DisableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DisableRequest{}
	st.MetastoreId = pb.MetastoreId
	st.SchemaName = pb.SchemaName

	return st, nil
}

type DisableResponse struct {
}

func disableResponseToPb(st *DisableResponse) (*disableResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &disableResponsePb{}

	return pb, nil
}

func (st *DisableResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &disableResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := disableResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DisableResponse) MarshalJSON() ([]byte, error) {
	pb, err := disableResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type disableResponsePb struct {
}

func disableResponseFromPb(pb *disableResponsePb) (*DisableResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DisableResponse{}

	return st, nil
}

type EffectivePermissionsList struct {
	// The privileges conveyed to each principal (either directly or via
	// inheritance)
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []EffectivePrivilegeAssignment
}

func effectivePermissionsListToPb(st *EffectivePermissionsList) (*effectivePermissionsListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &effectivePermissionsListPb{}

	var privilegeAssignmentsPb []effectivePrivilegeAssignmentPb
	for _, item := range st.PrivilegeAssignments {
		itemPb, err := effectivePrivilegeAssignmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			privilegeAssignmentsPb = append(privilegeAssignmentsPb, *itemPb)
		}
	}
	pb.PrivilegeAssignments = privilegeAssignmentsPb

	return pb, nil
}

func (st *EffectivePermissionsList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &effectivePermissionsListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := effectivePermissionsListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EffectivePermissionsList) MarshalJSON() ([]byte, error) {
	pb, err := effectivePermissionsListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type effectivePermissionsListPb struct {
	// The privileges conveyed to each principal (either directly or via
	// inheritance)
	PrivilegeAssignments []effectivePrivilegeAssignmentPb `json:"privilege_assignments,omitempty"`
}

func effectivePermissionsListFromPb(pb *effectivePermissionsListPb) (*EffectivePermissionsList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EffectivePermissionsList{}

	var privilegeAssignmentsField []EffectivePrivilegeAssignment
	for _, itemPb := range pb.PrivilegeAssignments {
		item, err := effectivePrivilegeAssignmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			privilegeAssignmentsField = append(privilegeAssignmentsField, *item)
		}
	}
	st.PrivilegeAssignments = privilegeAssignmentsField

	return st, nil
}

type EffectivePredictiveOptimizationFlag struct {
	// The name of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	// Wire name: 'inherited_from_name'
	InheritedFromName string
	// The type of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	// Wire name: 'inherited_from_type'
	InheritedFromType EffectivePredictiveOptimizationFlagInheritedFromType
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'value'
	Value EnablePredictiveOptimization

	ForceSendFields []string `tf:"-"`
}

func effectivePredictiveOptimizationFlagToPb(st *EffectivePredictiveOptimizationFlag) (*effectivePredictiveOptimizationFlagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &effectivePredictiveOptimizationFlagPb{}
	pb.InheritedFromName = st.InheritedFromName

	pb.InheritedFromType = st.InheritedFromType

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *EffectivePredictiveOptimizationFlag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &effectivePredictiveOptimizationFlagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := effectivePredictiveOptimizationFlagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EffectivePredictiveOptimizationFlag) MarshalJSON() ([]byte, error) {
	pb, err := effectivePredictiveOptimizationFlagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type effectivePredictiveOptimizationFlagPb struct {
	// The name of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromName string `json:"inherited_from_name,omitempty"`
	// The type of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromType EffectivePredictiveOptimizationFlagInheritedFromType `json:"inherited_from_type,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	Value EnablePredictiveOptimization `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func effectivePredictiveOptimizationFlagFromPb(pb *effectivePredictiveOptimizationFlagPb) (*EffectivePredictiveOptimizationFlag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EffectivePredictiveOptimizationFlag{}
	st.InheritedFromName = pb.InheritedFromName
	st.InheritedFromType = pb.InheritedFromType
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *effectivePredictiveOptimizationFlagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st effectivePredictiveOptimizationFlagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The type of the object from which the flag was inherited. If there was no
// inheritance, this field is left blank.
type EffectivePredictiveOptimizationFlagInheritedFromType string
type effectivePredictiveOptimizationFlagInheritedFromTypePb string

const EffectivePredictiveOptimizationFlagInheritedFromTypeCatalog EffectivePredictiveOptimizationFlagInheritedFromType = `CATALOG`

const EffectivePredictiveOptimizationFlagInheritedFromTypeSchema EffectivePredictiveOptimizationFlagInheritedFromType = `SCHEMA`

// String representation for [fmt.Print]
func (f *EffectivePredictiveOptimizationFlagInheritedFromType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EffectivePredictiveOptimizationFlagInheritedFromType) Set(v string) error {
	switch v {
	case `CATALOG`, `SCHEMA`:
		*f = EffectivePredictiveOptimizationFlagInheritedFromType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CATALOG", "SCHEMA"`, v)
	}
}

// Type always returns EffectivePredictiveOptimizationFlagInheritedFromType to satisfy [pflag.Value] interface
func (f *EffectivePredictiveOptimizationFlagInheritedFromType) Type() string {
	return "EffectivePredictiveOptimizationFlagInheritedFromType"
}

func effectivePredictiveOptimizationFlagInheritedFromTypeToPb(st *EffectivePredictiveOptimizationFlagInheritedFromType) (*effectivePredictiveOptimizationFlagInheritedFromTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := effectivePredictiveOptimizationFlagInheritedFromTypePb(*st)
	return &pb, nil
}

func effectivePredictiveOptimizationFlagInheritedFromTypeFromPb(pb *effectivePredictiveOptimizationFlagInheritedFromTypePb) (*EffectivePredictiveOptimizationFlagInheritedFromType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EffectivePredictiveOptimizationFlagInheritedFromType(*pb)
	return &st, nil
}

type EffectivePrivilege struct {
	// The full name of the object that conveys this privilege via inheritance.
	// This field is omitted when privilege is not inherited (it's assigned to
	// the securable itself).
	// Wire name: 'inherited_from_name'
	InheritedFromName string
	// The type of the object that conveys this privilege via inheritance. This
	// field is omitted when privilege is not inherited (it's assigned to the
	// securable itself).
	// Wire name: 'inherited_from_type'
	InheritedFromType SecurableType
	// The privilege assigned to the principal.
	// Wire name: 'privilege'
	Privilege Privilege

	ForceSendFields []string `tf:"-"`
}

func effectivePrivilegeToPb(st *EffectivePrivilege) (*effectivePrivilegePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &effectivePrivilegePb{}
	pb.InheritedFromName = st.InheritedFromName

	pb.InheritedFromType = st.InheritedFromType

	pb.Privilege = st.Privilege

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *EffectivePrivilege) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &effectivePrivilegePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := effectivePrivilegeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EffectivePrivilege) MarshalJSON() ([]byte, error) {
	pb, err := effectivePrivilegeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type effectivePrivilegePb struct {
	// The full name of the object that conveys this privilege via inheritance.
	// This field is omitted when privilege is not inherited (it's assigned to
	// the securable itself).
	InheritedFromName string `json:"inherited_from_name,omitempty"`
	// The type of the object that conveys this privilege via inheritance. This
	// field is omitted when privilege is not inherited (it's assigned to the
	// securable itself).
	InheritedFromType SecurableType `json:"inherited_from_type,omitempty"`
	// The privilege assigned to the principal.
	Privilege Privilege `json:"privilege,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func effectivePrivilegeFromPb(pb *effectivePrivilegePb) (*EffectivePrivilege, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EffectivePrivilege{}
	st.InheritedFromName = pb.InheritedFromName
	st.InheritedFromType = pb.InheritedFromType
	st.Privilege = pb.Privilege

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *effectivePrivilegePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st effectivePrivilegePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EffectivePrivilegeAssignment struct {
	// The principal (user email address or group name).
	// Wire name: 'principal'
	Principal string
	// The privileges conveyed to the principal (either directly or via
	// inheritance).
	// Wire name: 'privileges'
	Privileges []EffectivePrivilege

	ForceSendFields []string `tf:"-"`
}

func effectivePrivilegeAssignmentToPb(st *EffectivePrivilegeAssignment) (*effectivePrivilegeAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &effectivePrivilegeAssignmentPb{}
	pb.Principal = st.Principal

	var privilegesPb []effectivePrivilegePb
	for _, item := range st.Privileges {
		itemPb, err := effectivePrivilegeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			privilegesPb = append(privilegesPb, *itemPb)
		}
	}
	pb.Privileges = privilegesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *EffectivePrivilegeAssignment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &effectivePrivilegeAssignmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := effectivePrivilegeAssignmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EffectivePrivilegeAssignment) MarshalJSON() ([]byte, error) {
	pb, err := effectivePrivilegeAssignmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type effectivePrivilegeAssignmentPb struct {
	// The principal (user email address or group name).
	Principal string `json:"principal,omitempty"`
	// The privileges conveyed to the principal (either directly or via
	// inheritance).
	Privileges []effectivePrivilegePb `json:"privileges,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func effectivePrivilegeAssignmentFromPb(pb *effectivePrivilegeAssignmentPb) (*EffectivePrivilegeAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EffectivePrivilegeAssignment{}
	st.Principal = pb.Principal

	var privilegesField []EffectivePrivilege
	for _, itemPb := range pb.Privileges {
		item, err := effectivePrivilegeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			privilegesField = append(privilegesField, *item)
		}
	}
	st.Privileges = privilegesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *effectivePrivilegeAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st effectivePrivilegeAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Whether predictive optimization should be enabled for this object and objects
// under it.
type EnablePredictiveOptimization string
type enablePredictiveOptimizationPb string

const EnablePredictiveOptimizationDisable EnablePredictiveOptimization = `DISABLE`

const EnablePredictiveOptimizationEnable EnablePredictiveOptimization = `ENABLE`

const EnablePredictiveOptimizationInherit EnablePredictiveOptimization = `INHERIT`

// String representation for [fmt.Print]
func (f *EnablePredictiveOptimization) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EnablePredictiveOptimization) Set(v string) error {
	switch v {
	case `DISABLE`, `ENABLE`, `INHERIT`:
		*f = EnablePredictiveOptimization(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISABLE", "ENABLE", "INHERIT"`, v)
	}
}

// Type always returns EnablePredictiveOptimization to satisfy [pflag.Value] interface
func (f *EnablePredictiveOptimization) Type() string {
	return "EnablePredictiveOptimization"
}

func enablePredictiveOptimizationToPb(st *EnablePredictiveOptimization) (*enablePredictiveOptimizationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := enablePredictiveOptimizationPb(*st)
	return &pb, nil
}

func enablePredictiveOptimizationFromPb(pb *enablePredictiveOptimizationPb) (*EnablePredictiveOptimization, error) {
	if pb == nil {
		return nil, nil
	}
	st := EnablePredictiveOptimization(*pb)
	return &st, nil
}

// Enable a system schema
type EnableRequest struct {
	// The metastore ID under which the system schema lives.
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Full name of the system schema.
	// Wire name: 'schema_name'
	SchemaName string `tf:"-"`
}

func enableRequestToPb(st *EnableRequest) (*enableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enableRequestPb{}
	pb.MetastoreId = st.MetastoreId

	pb.SchemaName = st.SchemaName

	return pb, nil
}

func (st *EnableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnableRequest) MarshalJSON() ([]byte, error) {
	pb, err := enableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type enableRequestPb struct {
	// The metastore ID under which the system schema lives.
	MetastoreId string `json:"-" url:"-"`
	// Full name of the system schema.
	SchemaName string `json:"-" url:"-"`
}

func enableRequestFromPb(pb *enableRequestPb) (*EnableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnableRequest{}
	st.MetastoreId = pb.MetastoreId
	st.SchemaName = pb.SchemaName

	return st, nil
}

type EnableResponse struct {
}

func enableResponseToPb(st *EnableResponse) (*enableResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enableResponsePb{}

	return pb, nil
}

func (st *EnableResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enableResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enableResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnableResponse) MarshalJSON() ([]byte, error) {
	pb, err := enableResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type enableResponsePb struct {
}

func enableResponseFromPb(pb *enableResponsePb) (*EnableResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnableResponse{}

	return st, nil
}

// Encryption options that apply to clients connecting to cloud storage.
type EncryptionDetails struct {
	// Server-Side Encryption properties for clients communicating with AWS s3.
	// Wire name: 'sse_encryption_details'
	SseEncryptionDetails *SseEncryptionDetails
}

func encryptionDetailsToPb(st *EncryptionDetails) (*encryptionDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &encryptionDetailsPb{}
	sseEncryptionDetailsPb, err := sseEncryptionDetailsToPb(st.SseEncryptionDetails)
	if err != nil {
		return nil, err
	}
	if sseEncryptionDetailsPb != nil {
		pb.SseEncryptionDetails = sseEncryptionDetailsPb
	}

	return pb, nil
}

func (st *EncryptionDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &encryptionDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := encryptionDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EncryptionDetails) MarshalJSON() ([]byte, error) {
	pb, err := encryptionDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type encryptionDetailsPb struct {
	// Server-Side Encryption properties for clients communicating with AWS s3.
	SseEncryptionDetails *sseEncryptionDetailsPb `json:"sse_encryption_details,omitempty"`
}

func encryptionDetailsFromPb(pb *encryptionDetailsPb) (*EncryptionDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EncryptionDetails{}
	sseEncryptionDetailsField, err := sseEncryptionDetailsFromPb(pb.SseEncryptionDetails)
	if err != nil {
		return nil, err
	}
	if sseEncryptionDetailsField != nil {
		st.SseEncryptionDetails = sseEncryptionDetailsField
	}

	return st, nil
}

// Get boolean reflecting if table exists
type ExistsRequest struct {
	// Full name of the table.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
}

func existsRequestToPb(st *ExistsRequest) (*existsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &existsRequestPb{}
	pb.FullName = st.FullName

	return pb, nil
}

func (st *ExistsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &existsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := existsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExistsRequest) MarshalJSON() ([]byte, error) {
	pb, err := existsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type existsRequestPb struct {
	// Full name of the table.
	FullName string `json:"-" url:"-"`
}

func existsRequestFromPb(pb *existsRequestPb) (*ExistsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExistsRequest{}
	st.FullName = pb.FullName

	return st, nil
}

type ExternalLocationInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	// Wire name: 'access_point'
	AccessPoint string
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// Time at which this external location was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of external location creator.
	// Wire name: 'created_by'
	CreatedBy string
	// Unique ID of the location's storage credential.
	// Wire name: 'credential_id'
	CredentialId string
	// Name of the storage credential used with this location.
	// Wire name: 'credential_name'
	CredentialName string
	// Encryption options that apply to clients connecting to cloud storage.
	// Wire name: 'encryption_details'
	EncryptionDetails *EncryptionDetails
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	// Wire name: 'fallback'
	Fallback bool

	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode
	// Unique identifier of metastore hosting the external location.
	// Wire name: 'metastore_id'
	MetastoreId string
	// Name of the external location.
	// Wire name: 'name'
	Name string
	// The owner of the external location.
	// Wire name: 'owner'
	Owner string
	// Indicates whether the external location is read-only.
	// Wire name: 'read_only'
	ReadOnly bool
	// Time at which external location this was last modified, in epoch
	// milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// Username of user who last modified the external location.
	// Wire name: 'updated_by'
	UpdatedBy string
	// Path URL of the external location.
	// Wire name: 'url'
	Url string

	ForceSendFields []string `tf:"-"`
}

func externalLocationInfoToPb(st *ExternalLocationInfo) (*externalLocationInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalLocationInfoPb{}
	pb.AccessPoint = st.AccessPoint

	pb.BrowseOnly = st.BrowseOnly

	pb.Comment = st.Comment

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	pb.CredentialId = st.CredentialId

	pb.CredentialName = st.CredentialName

	encryptionDetailsPb, err := encryptionDetailsToPb(st.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsPb != nil {
		pb.EncryptionDetails = encryptionDetailsPb
	}

	pb.Fallback = st.Fallback

	pb.IsolationMode = st.IsolationMode

	pb.MetastoreId = st.MetastoreId

	pb.Name = st.Name

	pb.Owner = st.Owner

	pb.ReadOnly = st.ReadOnly

	pb.UpdatedAt = st.UpdatedAt

	pb.UpdatedBy = st.UpdatedBy

	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ExternalLocationInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalLocationInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalLocationInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalLocationInfo) MarshalJSON() ([]byte, error) {
	pb, err := externalLocationInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type externalLocationInfoPb struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string `json:"access_point,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this external location was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of external location creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique ID of the location's storage credential.
	CredentialId string `json:"credential_id,omitempty"`
	// Name of the storage credential used with this location.
	CredentialName string `json:"credential_name,omitempty"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *encryptionDetailsPb `json:"encryption_details,omitempty"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback bool `json:"fallback,omitempty"`

	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Unique identifier of metastore hosting the external location.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of the external location.
	Name string `json:"name,omitempty"`
	// The owner of the external location.
	Owner string `json:"owner,omitempty"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// Time at which external location this was last modified, in epoch
	// milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the external location.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Path URL of the external location.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func externalLocationInfoFromPb(pb *externalLocationInfoPb) (*ExternalLocationInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLocationInfo{}
	st.AccessPoint = pb.AccessPoint
	st.BrowseOnly = pb.BrowseOnly
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.CredentialId = pb.CredentialId
	st.CredentialName = pb.CredentialName
	encryptionDetailsField, err := encryptionDetailsFromPb(pb.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsField != nil {
		st.EncryptionDetails = encryptionDetailsField
	}
	st.Fallback = pb.Fallback
	st.IsolationMode = pb.IsolationMode
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.ReadOnly = pb.ReadOnly
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *externalLocationInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st externalLocationInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Detailed status of an online table. Shown if the online table is in the
// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
type FailedStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may only be partially synced to the online
	// table. Only populated if the table is still online and available for
	// serving.
	// Wire name: 'last_processed_commit_version'
	LastProcessedCommitVersion int64
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table. Only populated if the table is still online
	// and available for serving.
	// Wire name: 'timestamp'
	Timestamp string

	ForceSendFields []string `tf:"-"`
}

func failedStatusToPb(st *FailedStatus) (*failedStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &failedStatusPb{}
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion

	pb.Timestamp = st.Timestamp

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *FailedStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &failedStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := failedStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FailedStatus) MarshalJSON() ([]byte, error) {
	pb, err := failedStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type failedStatusPb struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may only be partially synced to the online
	// table. Only populated if the table is still online and available for
	// serving.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table. Only populated if the table is still online
	// and available for serving.
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func failedStatusFromPb(pb *failedStatusPb) (*FailedStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FailedStatus{}
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
	st.Timestamp = pb.Timestamp

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *failedStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st failedStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ForeignKeyConstraint struct {
	// Column names for this constraint.
	// Wire name: 'child_columns'
	ChildColumns []string
	// The name of the constraint.
	// Wire name: 'name'
	Name string
	// Column names for this constraint.
	// Wire name: 'parent_columns'
	ParentColumns []string
	// The full name of the parent constraint.
	// Wire name: 'parent_table'
	ParentTable string
}

func foreignKeyConstraintToPb(st *ForeignKeyConstraint) (*foreignKeyConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &foreignKeyConstraintPb{}
	pb.ChildColumns = st.ChildColumns

	pb.Name = st.Name

	pb.ParentColumns = st.ParentColumns

	pb.ParentTable = st.ParentTable

	return pb, nil
}

func (st *ForeignKeyConstraint) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &foreignKeyConstraintPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := foreignKeyConstraintFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ForeignKeyConstraint) MarshalJSON() ([]byte, error) {
	pb, err := foreignKeyConstraintToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type foreignKeyConstraintPb struct {
	// Column names for this constraint.
	ChildColumns []string `json:"child_columns"`
	// The name of the constraint.
	Name string `json:"name"`
	// Column names for this constraint.
	ParentColumns []string `json:"parent_columns"`
	// The full name of the parent constraint.
	ParentTable string `json:"parent_table"`
}

func foreignKeyConstraintFromPb(pb *foreignKeyConstraintPb) (*ForeignKeyConstraint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForeignKeyConstraint{}
	st.ChildColumns = pb.ChildColumns
	st.Name = pb.Name
	st.ParentColumns = pb.ParentColumns
	st.ParentTable = pb.ParentTable

	return st, nil
}

// A function that is dependent on a SQL object.
type FunctionDependency struct {
	// Full name of the dependent function, in the form of
	// __catalog_name__.__schema_name__.__function_name__.
	// Wire name: 'function_full_name'
	FunctionFullName string
}

func functionDependencyToPb(st *FunctionDependency) (*functionDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &functionDependencyPb{}
	pb.FunctionFullName = st.FunctionFullName

	return pb, nil
}

func (st *FunctionDependency) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &functionDependencyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := functionDependencyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FunctionDependency) MarshalJSON() ([]byte, error) {
	pb, err := functionDependencyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type functionDependencyPb struct {
	// Full name of the dependent function, in the form of
	// __catalog_name__.__schema_name__.__function_name__.
	FunctionFullName string `json:"function_full_name"`
}

func functionDependencyFromPb(pb *functionDependencyPb) (*FunctionDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionDependency{}
	st.FunctionFullName = pb.FunctionFullName

	return st, nil
}

type FunctionInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool
	// Name of parent catalog.
	// Wire name: 'catalog_name'
	CatalogName string
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// Time at which this function was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of function creator.
	// Wire name: 'created_by'
	CreatedBy string
	// Scalar function return data type.
	// Wire name: 'data_type'
	DataType ColumnTypeName
	// External function language.
	// Wire name: 'external_language'
	ExternalLanguage string
	// External function name.
	// Wire name: 'external_name'
	ExternalName string
	// Pretty printed function data type.
	// Wire name: 'full_data_type'
	FullDataType string
	// Full name of function, in form of
	// __catalog_name__.__schema_name__.__function__name__
	// Wire name: 'full_name'
	FullName string
	// Id of Function, relative to parent schema.
	// Wire name: 'function_id'
	FunctionId string

	// Wire name: 'input_params'
	InputParams *FunctionParameterInfos
	// Whether the function is deterministic.
	// Wire name: 'is_deterministic'
	IsDeterministic bool
	// Function null call.
	// Wire name: 'is_null_call'
	IsNullCall bool
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string
	// Name of function, relative to parent schema.
	// Wire name: 'name'
	Name string
	// Username of current owner of function.
	// Wire name: 'owner'
	Owner string
	// Function parameter style. **S** is the value for SQL.
	// Wire name: 'parameter_style'
	ParameterStyle FunctionInfoParameterStyle
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	// Wire name: 'properties'
	Properties string
	// Table function return parameters.
	// Wire name: 'return_params'
	ReturnParams *FunctionParameterInfos
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	// Wire name: 'routine_body'
	RoutineBody FunctionInfoRoutineBody
	// Function body.
	// Wire name: 'routine_definition'
	RoutineDefinition string
	// Function dependencies.
	// Wire name: 'routine_dependencies'
	RoutineDependencies *DependencyList
	// Name of parent schema relative to its parent catalog.
	// Wire name: 'schema_name'
	SchemaName string
	// Function security type.
	// Wire name: 'security_type'
	SecurityType FunctionInfoSecurityType
	// Specific name of the function; Reserved for future use.
	// Wire name: 'specific_name'
	SpecificName string
	// Function SQL data access.
	// Wire name: 'sql_data_access'
	SqlDataAccess FunctionInfoSqlDataAccess
	// List of schemes whose objects can be referenced without qualification.
	// Wire name: 'sql_path'
	SqlPath string
	// Time at which this function was created, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// Username of user who last modified function.
	// Wire name: 'updated_by'
	UpdatedBy string

	ForceSendFields []string `tf:"-"`
}

func functionInfoToPb(st *FunctionInfo) (*functionInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &functionInfoPb{}
	pb.BrowseOnly = st.BrowseOnly

	pb.CatalogName = st.CatalogName

	pb.Comment = st.Comment

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	pb.DataType = st.DataType

	pb.ExternalLanguage = st.ExternalLanguage

	pb.ExternalName = st.ExternalName

	pb.FullDataType = st.FullDataType

	pb.FullName = st.FullName

	pb.FunctionId = st.FunctionId

	inputParamsPb, err := functionParameterInfosToPb(st.InputParams)
	if err != nil {
		return nil, err
	}
	if inputParamsPb != nil {
		pb.InputParams = inputParamsPb
	}

	pb.IsDeterministic = st.IsDeterministic

	pb.IsNullCall = st.IsNullCall

	pb.MetastoreId = st.MetastoreId

	pb.Name = st.Name

	pb.Owner = st.Owner

	pb.ParameterStyle = st.ParameterStyle

	pb.Properties = st.Properties

	returnParamsPb, err := functionParameterInfosToPb(st.ReturnParams)
	if err != nil {
		return nil, err
	}
	if returnParamsPb != nil {
		pb.ReturnParams = returnParamsPb
	}

	pb.RoutineBody = st.RoutineBody

	pb.RoutineDefinition = st.RoutineDefinition

	routineDependenciesPb, err := dependencyListToPb(st.RoutineDependencies)
	if err != nil {
		return nil, err
	}
	if routineDependenciesPb != nil {
		pb.RoutineDependencies = routineDependenciesPb
	}

	pb.SchemaName = st.SchemaName

	pb.SecurityType = st.SecurityType

	pb.SpecificName = st.SpecificName

	pb.SqlDataAccess = st.SqlDataAccess

	pb.SqlPath = st.SqlPath

	pb.UpdatedAt = st.UpdatedAt

	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *FunctionInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &functionInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := functionInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FunctionInfo) MarshalJSON() ([]byte, error) {
	pb, err := functionInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type functionInfoPb struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// Name of parent catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this function was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of function creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Scalar function return data type.
	DataType ColumnTypeName `json:"data_type,omitempty"`
	// External function language.
	ExternalLanguage string `json:"external_language,omitempty"`
	// External function name.
	ExternalName string `json:"external_name,omitempty"`
	// Pretty printed function data type.
	FullDataType string `json:"full_data_type,omitempty"`
	// Full name of function, in form of
	// __catalog_name__.__schema_name__.__function__name__
	FullName string `json:"full_name,omitempty"`
	// Id of Function, relative to parent schema.
	FunctionId string `json:"function_id,omitempty"`

	InputParams *functionParameterInfosPb `json:"input_params,omitempty"`
	// Whether the function is deterministic.
	IsDeterministic bool `json:"is_deterministic,omitempty"`
	// Function null call.
	IsNullCall bool `json:"is_null_call,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of function, relative to parent schema.
	Name string `json:"name,omitempty"`
	// Username of current owner of function.
	Owner string `json:"owner,omitempty"`
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle FunctionInfoParameterStyle `json:"parameter_style,omitempty"`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties string `json:"properties,omitempty"`
	// Table function return parameters.
	ReturnParams *functionParameterInfosPb `json:"return_params,omitempty"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody FunctionInfoRoutineBody `json:"routine_body,omitempty"`
	// Function body.
	RoutineDefinition string `json:"routine_definition,omitempty"`
	// Function dependencies.
	RoutineDependencies *dependencyListPb `json:"routine_dependencies,omitempty"`
	// Name of parent schema relative to its parent catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// Function security type.
	SecurityType FunctionInfoSecurityType `json:"security_type,omitempty"`
	// Specific name of the function; Reserved for future use.
	SpecificName string `json:"specific_name,omitempty"`
	// Function SQL data access.
	SqlDataAccess FunctionInfoSqlDataAccess `json:"sql_data_access,omitempty"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string `json:"sql_path,omitempty"`
	// Time at which this function was created, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified function.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func functionInfoFromPb(pb *functionInfoPb) (*FunctionInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionInfo{}
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.DataType = pb.DataType
	st.ExternalLanguage = pb.ExternalLanguage
	st.ExternalName = pb.ExternalName
	st.FullDataType = pb.FullDataType
	st.FullName = pb.FullName
	st.FunctionId = pb.FunctionId
	inputParamsField, err := functionParameterInfosFromPb(pb.InputParams)
	if err != nil {
		return nil, err
	}
	if inputParamsField != nil {
		st.InputParams = inputParamsField
	}
	st.IsDeterministic = pb.IsDeterministic
	st.IsNullCall = pb.IsNullCall
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.ParameterStyle = pb.ParameterStyle
	st.Properties = pb.Properties
	returnParamsField, err := functionParameterInfosFromPb(pb.ReturnParams)
	if err != nil {
		return nil, err
	}
	if returnParamsField != nil {
		st.ReturnParams = returnParamsField
	}
	st.RoutineBody = pb.RoutineBody
	st.RoutineDefinition = pb.RoutineDefinition
	routineDependenciesField, err := dependencyListFromPb(pb.RoutineDependencies)
	if err != nil {
		return nil, err
	}
	if routineDependenciesField != nil {
		st.RoutineDependencies = routineDependenciesField
	}
	st.SchemaName = pb.SchemaName
	st.SecurityType = pb.SecurityType
	st.SpecificName = pb.SpecificName
	st.SqlDataAccess = pb.SqlDataAccess
	st.SqlPath = pb.SqlPath
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *functionInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st functionInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Function parameter style. **S** is the value for SQL.
type FunctionInfoParameterStyle string
type functionInfoParameterStylePb string

const FunctionInfoParameterStyleS FunctionInfoParameterStyle = `S`

// String representation for [fmt.Print]
func (f *FunctionInfoParameterStyle) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionInfoParameterStyle) Set(v string) error {
	switch v {
	case `S`:
		*f = FunctionInfoParameterStyle(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "S"`, v)
	}
}

// Type always returns FunctionInfoParameterStyle to satisfy [pflag.Value] interface
func (f *FunctionInfoParameterStyle) Type() string {
	return "FunctionInfoParameterStyle"
}

func functionInfoParameterStyleToPb(st *FunctionInfoParameterStyle) (*functionInfoParameterStylePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := functionInfoParameterStylePb(*st)
	return &pb, nil
}

func functionInfoParameterStyleFromPb(pb *functionInfoParameterStylePb) (*FunctionInfoParameterStyle, error) {
	if pb == nil {
		return nil, nil
	}
	st := FunctionInfoParameterStyle(*pb)
	return &st, nil
}

// Function language. When **EXTERNAL** is used, the language of the routine
// function should be specified in the __external_language__ field, and the
// __return_params__ of the function cannot be used (as **TABLE** return type is
// not supported), and the __sql_data_access__ field must be **NO_SQL**.
type FunctionInfoRoutineBody string
type functionInfoRoutineBodyPb string

const FunctionInfoRoutineBodyExternal FunctionInfoRoutineBody = `EXTERNAL`

const FunctionInfoRoutineBodySql FunctionInfoRoutineBody = `SQL`

// String representation for [fmt.Print]
func (f *FunctionInfoRoutineBody) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionInfoRoutineBody) Set(v string) error {
	switch v {
	case `EXTERNAL`, `SQL`:
		*f = FunctionInfoRoutineBody(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "SQL"`, v)
	}
}

// Type always returns FunctionInfoRoutineBody to satisfy [pflag.Value] interface
func (f *FunctionInfoRoutineBody) Type() string {
	return "FunctionInfoRoutineBody"
}

func functionInfoRoutineBodyToPb(st *FunctionInfoRoutineBody) (*functionInfoRoutineBodyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := functionInfoRoutineBodyPb(*st)
	return &pb, nil
}

func functionInfoRoutineBodyFromPb(pb *functionInfoRoutineBodyPb) (*FunctionInfoRoutineBody, error) {
	if pb == nil {
		return nil, nil
	}
	st := FunctionInfoRoutineBody(*pb)
	return &st, nil
}

// The security type of the function.
type FunctionInfoSecurityType string
type functionInfoSecurityTypePb string

const FunctionInfoSecurityTypeDefiner FunctionInfoSecurityType = `DEFINER`

// String representation for [fmt.Print]
func (f *FunctionInfoSecurityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionInfoSecurityType) Set(v string) error {
	switch v {
	case `DEFINER`:
		*f = FunctionInfoSecurityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEFINER"`, v)
	}
}

// Type always returns FunctionInfoSecurityType to satisfy [pflag.Value] interface
func (f *FunctionInfoSecurityType) Type() string {
	return "FunctionInfoSecurityType"
}

func functionInfoSecurityTypeToPb(st *FunctionInfoSecurityType) (*functionInfoSecurityTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := functionInfoSecurityTypePb(*st)
	return &pb, nil
}

func functionInfoSecurityTypeFromPb(pb *functionInfoSecurityTypePb) (*FunctionInfoSecurityType, error) {
	if pb == nil {
		return nil, nil
	}
	st := FunctionInfoSecurityType(*pb)
	return &st, nil
}

// Function SQL data access.
type FunctionInfoSqlDataAccess string
type functionInfoSqlDataAccessPb string

const FunctionInfoSqlDataAccessContainsSql FunctionInfoSqlDataAccess = `CONTAINS_SQL`

const FunctionInfoSqlDataAccessNoSql FunctionInfoSqlDataAccess = `NO_SQL`

const FunctionInfoSqlDataAccessReadsSqlData FunctionInfoSqlDataAccess = `READS_SQL_DATA`

// String representation for [fmt.Print]
func (f *FunctionInfoSqlDataAccess) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionInfoSqlDataAccess) Set(v string) error {
	switch v {
	case `CONTAINS_SQL`, `NO_SQL`, `READS_SQL_DATA`:
		*f = FunctionInfoSqlDataAccess(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTAINS_SQL", "NO_SQL", "READS_SQL_DATA"`, v)
	}
}

// Type always returns FunctionInfoSqlDataAccess to satisfy [pflag.Value] interface
func (f *FunctionInfoSqlDataAccess) Type() string {
	return "FunctionInfoSqlDataAccess"
}

func functionInfoSqlDataAccessToPb(st *FunctionInfoSqlDataAccess) (*functionInfoSqlDataAccessPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := functionInfoSqlDataAccessPb(*st)
	return &pb, nil
}

func functionInfoSqlDataAccessFromPb(pb *functionInfoSqlDataAccessPb) (*FunctionInfoSqlDataAccess, error) {
	if pb == nil {
		return nil, nil
	}
	st := FunctionInfoSqlDataAccess(*pb)
	return &st, nil
}

type FunctionParameterInfo struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// Name of parameter.
	// Wire name: 'name'
	Name string
	// Default value of the parameter.
	// Wire name: 'parameter_default'
	ParameterDefault string
	// The mode of the function parameter.
	// Wire name: 'parameter_mode'
	ParameterMode FunctionParameterMode
	// The type of function parameter.
	// Wire name: 'parameter_type'
	ParameterType FunctionParameterType
	// Ordinal position of column (starting at position 0).
	// Wire name: 'position'
	Position int
	// Format of IntervalType.
	// Wire name: 'type_interval_type'
	TypeIntervalType string
	// Full data type spec, JSON-serialized.
	// Wire name: 'type_json'
	TypeJson string

	// Wire name: 'type_name'
	TypeName ColumnTypeName
	// Digits of precision; required on Create for DecimalTypes.
	// Wire name: 'type_precision'
	TypePrecision int
	// Digits to right of decimal; Required on Create for DecimalTypes.
	// Wire name: 'type_scale'
	TypeScale int
	// Full data type spec, SQL/catalogString text.
	// Wire name: 'type_text'
	TypeText string

	ForceSendFields []string `tf:"-"`
}

func functionParameterInfoToPb(st *FunctionParameterInfo) (*functionParameterInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &functionParameterInfoPb{}
	pb.Comment = st.Comment

	pb.Name = st.Name

	pb.ParameterDefault = st.ParameterDefault

	pb.ParameterMode = st.ParameterMode

	pb.ParameterType = st.ParameterType

	pb.Position = st.Position

	pb.TypeIntervalType = st.TypeIntervalType

	pb.TypeJson = st.TypeJson

	pb.TypeName = st.TypeName

	pb.TypePrecision = st.TypePrecision

	pb.TypeScale = st.TypeScale

	pb.TypeText = st.TypeText

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *FunctionParameterInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &functionParameterInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := functionParameterInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FunctionParameterInfo) MarshalJSON() ([]byte, error) {
	pb, err := functionParameterInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type functionParameterInfoPb struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of parameter.
	Name string `json:"name"`
	// Default value of the parameter.
	ParameterDefault string `json:"parameter_default,omitempty"`
	// The mode of the function parameter.
	ParameterMode FunctionParameterMode `json:"parameter_mode,omitempty"`
	// The type of function parameter.
	ParameterType FunctionParameterType `json:"parameter_type,omitempty"`
	// Ordinal position of column (starting at position 0).
	Position int `json:"position"`
	// Format of IntervalType.
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// Full data type spec, JSON-serialized.
	TypeJson string `json:"type_json,omitempty"`

	TypeName ColumnTypeName `json:"type_name"`
	// Digits of precision; required on Create for DecimalTypes.
	TypePrecision int `json:"type_precision,omitempty"`
	// Digits to right of decimal; Required on Create for DecimalTypes.
	TypeScale int `json:"type_scale,omitempty"`
	// Full data type spec, SQL/catalogString text.
	TypeText string `json:"type_text"`

	ForceSendFields []string `json:"-" url:"-"`
}

func functionParameterInfoFromPb(pb *functionParameterInfoPb) (*FunctionParameterInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionParameterInfo{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.ParameterDefault = pb.ParameterDefault
	st.ParameterMode = pb.ParameterMode
	st.ParameterType = pb.ParameterType
	st.Position = pb.Position
	st.TypeIntervalType = pb.TypeIntervalType
	st.TypeJson = pb.TypeJson
	st.TypeName = pb.TypeName
	st.TypePrecision = pb.TypePrecision
	st.TypeScale = pb.TypeScale
	st.TypeText = pb.TypeText

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *functionParameterInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st functionParameterInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FunctionParameterInfos struct {
	// The array of __FunctionParameterInfo__ definitions of the function's
	// parameters.
	// Wire name: 'parameters'
	Parameters []FunctionParameterInfo
}

func functionParameterInfosToPb(st *FunctionParameterInfos) (*functionParameterInfosPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &functionParameterInfosPb{}

	var parametersPb []functionParameterInfoPb
	for _, item := range st.Parameters {
		itemPb, err := functionParameterInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	return pb, nil
}

func (st *FunctionParameterInfos) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &functionParameterInfosPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := functionParameterInfosFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FunctionParameterInfos) MarshalJSON() ([]byte, error) {
	pb, err := functionParameterInfosToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type functionParameterInfosPb struct {
	// The array of __FunctionParameterInfo__ definitions of the function's
	// parameters.
	Parameters []functionParameterInfoPb `json:"parameters,omitempty"`
}

func functionParameterInfosFromPb(pb *functionParameterInfosPb) (*FunctionParameterInfos, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionParameterInfos{}

	var parametersField []FunctionParameterInfo
	for _, itemPb := range pb.Parameters {
		item, err := functionParameterInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField

	return st, nil
}

// The mode of the function parameter.
type FunctionParameterMode string
type functionParameterModePb string

const FunctionParameterModeIn FunctionParameterMode = `IN`

// String representation for [fmt.Print]
func (f *FunctionParameterMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionParameterMode) Set(v string) error {
	switch v {
	case `IN`:
		*f = FunctionParameterMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IN"`, v)
	}
}

// Type always returns FunctionParameterMode to satisfy [pflag.Value] interface
func (f *FunctionParameterMode) Type() string {
	return "FunctionParameterMode"
}

func functionParameterModeToPb(st *FunctionParameterMode) (*functionParameterModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := functionParameterModePb(*st)
	return &pb, nil
}

func functionParameterModeFromPb(pb *functionParameterModePb) (*FunctionParameterMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := FunctionParameterMode(*pb)
	return &st, nil
}

// The type of function parameter.
type FunctionParameterType string
type functionParameterTypePb string

const FunctionParameterTypeColumn FunctionParameterType = `COLUMN`

const FunctionParameterTypeParam FunctionParameterType = `PARAM`

// String representation for [fmt.Print]
func (f *FunctionParameterType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionParameterType) Set(v string) error {
	switch v {
	case `COLUMN`, `PARAM`:
		*f = FunctionParameterType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COLUMN", "PARAM"`, v)
	}
}

// Type always returns FunctionParameterType to satisfy [pflag.Value] interface
func (f *FunctionParameterType) Type() string {
	return "FunctionParameterType"
}

func functionParameterTypeToPb(st *FunctionParameterType) (*functionParameterTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := functionParameterTypePb(*st)
	return &pb, nil
}

func functionParameterTypeFromPb(pb *functionParameterTypePb) (*FunctionParameterType, error) {
	if pb == nil {
		return nil, nil
	}
	st := FunctionParameterType(*pb)
	return &st, nil
}

// GCP temporary credentials for API authentication. Read more at
// https://developers.google.com/identity/protocols/oauth2/service-account
type GcpOauthToken struct {

	// Wire name: 'oauth_token'
	OauthToken string

	ForceSendFields []string `tf:"-"`
}

func gcpOauthTokenToPb(st *GcpOauthToken) (*gcpOauthTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpOauthTokenPb{}
	pb.OauthToken = st.OauthToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GcpOauthToken) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gcpOauthTokenPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gcpOauthTokenFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GcpOauthToken) MarshalJSON() ([]byte, error) {
	pb, err := gcpOauthTokenToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type gcpOauthTokenPb struct {
	OauthToken string `json:"oauth_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func gcpOauthTokenFromPb(pb *gcpOauthTokenPb) (*GcpOauthToken, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpOauthToken{}
	st.OauthToken = pb.OauthToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *gcpOauthTokenPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st gcpOauthTokenPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The Azure cloud options to customize the requested temporary credential
type GenerateTemporaryServiceCredentialAzureOptions struct {
	// The resources to which the temporary Azure credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://learn.microsoft.com/python/api/azure-core/azure.core.credentials.tokencredential?view=azure-python)
	// Wire name: 'resources'
	Resources []string
}

func generateTemporaryServiceCredentialAzureOptionsToPb(st *GenerateTemporaryServiceCredentialAzureOptions) (*generateTemporaryServiceCredentialAzureOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &generateTemporaryServiceCredentialAzureOptionsPb{}
	pb.Resources = st.Resources

	return pb, nil
}

func (st *GenerateTemporaryServiceCredentialAzureOptions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &generateTemporaryServiceCredentialAzureOptionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := generateTemporaryServiceCredentialAzureOptionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenerateTemporaryServiceCredentialAzureOptions) MarshalJSON() ([]byte, error) {
	pb, err := generateTemporaryServiceCredentialAzureOptionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type generateTemporaryServiceCredentialAzureOptionsPb struct {
	// The resources to which the temporary Azure credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://learn.microsoft.com/python/api/azure-core/azure.core.credentials.tokencredential?view=azure-python)
	Resources []string `json:"resources,omitempty"`
}

func generateTemporaryServiceCredentialAzureOptionsFromPb(pb *generateTemporaryServiceCredentialAzureOptionsPb) (*GenerateTemporaryServiceCredentialAzureOptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateTemporaryServiceCredentialAzureOptions{}
	st.Resources = pb.Resources

	return st, nil
}

// The GCP cloud options to customize the requested temporary credential
type GenerateTemporaryServiceCredentialGcpOptions struct {
	// The scopes to which the temporary GCP credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://google-auth.readthedocs.io/en/latest/reference/google.auth.html#google.auth.credentials.Credentials)
	// Wire name: 'scopes'
	Scopes []string
}

func generateTemporaryServiceCredentialGcpOptionsToPb(st *GenerateTemporaryServiceCredentialGcpOptions) (*generateTemporaryServiceCredentialGcpOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &generateTemporaryServiceCredentialGcpOptionsPb{}
	pb.Scopes = st.Scopes

	return pb, nil
}

func (st *GenerateTemporaryServiceCredentialGcpOptions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &generateTemporaryServiceCredentialGcpOptionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := generateTemporaryServiceCredentialGcpOptionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenerateTemporaryServiceCredentialGcpOptions) MarshalJSON() ([]byte, error) {
	pb, err := generateTemporaryServiceCredentialGcpOptionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type generateTemporaryServiceCredentialGcpOptionsPb struct {
	// The scopes to which the temporary GCP credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://google-auth.readthedocs.io/en/latest/reference/google.auth.html#google.auth.credentials.Credentials)
	Scopes []string `json:"scopes,omitempty"`
}

func generateTemporaryServiceCredentialGcpOptionsFromPb(pb *generateTemporaryServiceCredentialGcpOptionsPb) (*GenerateTemporaryServiceCredentialGcpOptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateTemporaryServiceCredentialGcpOptions{}
	st.Scopes = pb.Scopes

	return st, nil
}

type GenerateTemporaryServiceCredentialRequest struct {
	// The Azure cloud options to customize the requested temporary credential
	// Wire name: 'azure_options'
	AzureOptions *GenerateTemporaryServiceCredentialAzureOptions
	// The name of the service credential used to generate a temporary
	// credential
	// Wire name: 'credential_name'
	CredentialName string
	// The GCP cloud options to customize the requested temporary credential
	// Wire name: 'gcp_options'
	GcpOptions *GenerateTemporaryServiceCredentialGcpOptions
}

func generateTemporaryServiceCredentialRequestToPb(st *GenerateTemporaryServiceCredentialRequest) (*generateTemporaryServiceCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &generateTemporaryServiceCredentialRequestPb{}
	azureOptionsPb, err := generateTemporaryServiceCredentialAzureOptionsToPb(st.AzureOptions)
	if err != nil {
		return nil, err
	}
	if azureOptionsPb != nil {
		pb.AzureOptions = azureOptionsPb
	}

	pb.CredentialName = st.CredentialName

	gcpOptionsPb, err := generateTemporaryServiceCredentialGcpOptionsToPb(st.GcpOptions)
	if err != nil {
		return nil, err
	}
	if gcpOptionsPb != nil {
		pb.GcpOptions = gcpOptionsPb
	}

	return pb, nil
}

func (st *GenerateTemporaryServiceCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &generateTemporaryServiceCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := generateTemporaryServiceCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenerateTemporaryServiceCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := generateTemporaryServiceCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type generateTemporaryServiceCredentialRequestPb struct {
	// The Azure cloud options to customize the requested temporary credential
	AzureOptions *generateTemporaryServiceCredentialAzureOptionsPb `json:"azure_options,omitempty"`
	// The name of the service credential used to generate a temporary
	// credential
	CredentialName string `json:"credential_name"`
	// The GCP cloud options to customize the requested temporary credential
	GcpOptions *generateTemporaryServiceCredentialGcpOptionsPb `json:"gcp_options,omitempty"`
}

func generateTemporaryServiceCredentialRequestFromPb(pb *generateTemporaryServiceCredentialRequestPb) (*GenerateTemporaryServiceCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateTemporaryServiceCredentialRequest{}
	azureOptionsField, err := generateTemporaryServiceCredentialAzureOptionsFromPb(pb.AzureOptions)
	if err != nil {
		return nil, err
	}
	if azureOptionsField != nil {
		st.AzureOptions = azureOptionsField
	}
	st.CredentialName = pb.CredentialName
	gcpOptionsField, err := generateTemporaryServiceCredentialGcpOptionsFromPb(pb.GcpOptions)
	if err != nil {
		return nil, err
	}
	if gcpOptionsField != nil {
		st.GcpOptions = gcpOptionsField
	}

	return st, nil
}

type GenerateTemporaryTableCredentialRequest struct {
	// The operation performed against the table data, either READ or
	// READ_WRITE. If READ_WRITE is specified, the credentials returned will
	// have write permissions, otherwise, it will be read only.
	// Wire name: 'operation'
	Operation TableOperation
	// UUID of the table to read or write.
	// Wire name: 'table_id'
	TableId string

	ForceSendFields []string `tf:"-"`
}

func generateTemporaryTableCredentialRequestToPb(st *GenerateTemporaryTableCredentialRequest) (*generateTemporaryTableCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &generateTemporaryTableCredentialRequestPb{}
	pb.Operation = st.Operation

	pb.TableId = st.TableId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GenerateTemporaryTableCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &generateTemporaryTableCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := generateTemporaryTableCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenerateTemporaryTableCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := generateTemporaryTableCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type generateTemporaryTableCredentialRequestPb struct {
	// The operation performed against the table data, either READ or
	// READ_WRITE. If READ_WRITE is specified, the credentials returned will
	// have write permissions, otherwise, it will be read only.
	Operation TableOperation `json:"operation,omitempty"`
	// UUID of the table to read or write.
	TableId string `json:"table_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func generateTemporaryTableCredentialRequestFromPb(pb *generateTemporaryTableCredentialRequestPb) (*GenerateTemporaryTableCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateTemporaryTableCredentialRequest{}
	st.Operation = pb.Operation
	st.TableId = pb.TableId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *generateTemporaryTableCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st generateTemporaryTableCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenerateTemporaryTableCredentialResponse struct {
	// AWS temporary credentials for API authentication. Read more at
	// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
	// Wire name: 'aws_temp_credentials'
	AwsTempCredentials *AwsCredentials
	// Azure Active Directory token, essentially the Oauth token for Azure
	// Service Principal or Managed Identity. Read more at
	// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
	// Wire name: 'azure_aad'
	AzureAad *AzureActiveDirectoryToken
	// Azure temporary credentials for API authentication. Read more at
	// https://docs.microsoft.com/en-us/rest/api/storageservices/create-user-delegation-sas
	// Wire name: 'azure_user_delegation_sas'
	AzureUserDelegationSas *AzureUserDelegationSas
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	// Wire name: 'expiration_time'
	ExpirationTime int64
	// GCP temporary credentials for API authentication. Read more at
	// https://developers.google.com/identity/protocols/oauth2/service-account
	// Wire name: 'gcp_oauth_token'
	GcpOauthToken *GcpOauthToken
	// R2 temporary credentials for API authentication. Read more at
	// https://developers.cloudflare.com/r2/api/s3/tokens/.
	// Wire name: 'r2_temp_credentials'
	R2TempCredentials *R2Credentials
	// The URL of the storage path accessible by the temporary credential.
	// Wire name: 'url'
	Url string

	ForceSendFields []string `tf:"-"`
}

func generateTemporaryTableCredentialResponseToPb(st *GenerateTemporaryTableCredentialResponse) (*generateTemporaryTableCredentialResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &generateTemporaryTableCredentialResponsePb{}
	awsTempCredentialsPb, err := awsCredentialsToPb(st.AwsTempCredentials)
	if err != nil {
		return nil, err
	}
	if awsTempCredentialsPb != nil {
		pb.AwsTempCredentials = awsTempCredentialsPb
	}

	azureAadPb, err := azureActiveDirectoryTokenToPb(st.AzureAad)
	if err != nil {
		return nil, err
	}
	if azureAadPb != nil {
		pb.AzureAad = azureAadPb
	}

	azureUserDelegationSasPb, err := azureUserDelegationSasToPb(st.AzureUserDelegationSas)
	if err != nil {
		return nil, err
	}
	if azureUserDelegationSasPb != nil {
		pb.AzureUserDelegationSas = azureUserDelegationSasPb
	}

	pb.ExpirationTime = st.ExpirationTime

	gcpOauthTokenPb, err := gcpOauthTokenToPb(st.GcpOauthToken)
	if err != nil {
		return nil, err
	}
	if gcpOauthTokenPb != nil {
		pb.GcpOauthToken = gcpOauthTokenPb
	}

	r2TempCredentialsPb, err := r2CredentialsToPb(st.R2TempCredentials)
	if err != nil {
		return nil, err
	}
	if r2TempCredentialsPb != nil {
		pb.R2TempCredentials = r2TempCredentialsPb
	}

	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GenerateTemporaryTableCredentialResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &generateTemporaryTableCredentialResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := generateTemporaryTableCredentialResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenerateTemporaryTableCredentialResponse) MarshalJSON() ([]byte, error) {
	pb, err := generateTemporaryTableCredentialResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type generateTemporaryTableCredentialResponsePb struct {
	// AWS temporary credentials for API authentication. Read more at
	// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
	AwsTempCredentials *awsCredentialsPb `json:"aws_temp_credentials,omitempty"`
	// Azure Active Directory token, essentially the Oauth token for Azure
	// Service Principal or Managed Identity. Read more at
	// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
	AzureAad *azureActiveDirectoryTokenPb `json:"azure_aad,omitempty"`
	// Azure temporary credentials for API authentication. Read more at
	// https://docs.microsoft.com/en-us/rest/api/storageservices/create-user-delegation-sas
	AzureUserDelegationSas *azureUserDelegationSasPb `json:"azure_user_delegation_sas,omitempty"`
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// GCP temporary credentials for API authentication. Read more at
	// https://developers.google.com/identity/protocols/oauth2/service-account
	GcpOauthToken *gcpOauthTokenPb `json:"gcp_oauth_token,omitempty"`
	// R2 temporary credentials for API authentication. Read more at
	// https://developers.cloudflare.com/r2/api/s3/tokens/.
	R2TempCredentials *r2CredentialsPb `json:"r2_temp_credentials,omitempty"`
	// The URL of the storage path accessible by the temporary credential.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func generateTemporaryTableCredentialResponseFromPb(pb *generateTemporaryTableCredentialResponsePb) (*GenerateTemporaryTableCredentialResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateTemporaryTableCredentialResponse{}
	awsTempCredentialsField, err := awsCredentialsFromPb(pb.AwsTempCredentials)
	if err != nil {
		return nil, err
	}
	if awsTempCredentialsField != nil {
		st.AwsTempCredentials = awsTempCredentialsField
	}
	azureAadField, err := azureActiveDirectoryTokenFromPb(pb.AzureAad)
	if err != nil {
		return nil, err
	}
	if azureAadField != nil {
		st.AzureAad = azureAadField
	}
	azureUserDelegationSasField, err := azureUserDelegationSasFromPb(pb.AzureUserDelegationSas)
	if err != nil {
		return nil, err
	}
	if azureUserDelegationSasField != nil {
		st.AzureUserDelegationSas = azureUserDelegationSasField
	}
	st.ExpirationTime = pb.ExpirationTime
	gcpOauthTokenField, err := gcpOauthTokenFromPb(pb.GcpOauthToken)
	if err != nil {
		return nil, err
	}
	if gcpOauthTokenField != nil {
		st.GcpOauthToken = gcpOauthTokenField
	}
	r2TempCredentialsField, err := r2CredentialsFromPb(pb.R2TempCredentials)
	if err != nil {
		return nil, err
	}
	if r2TempCredentialsField != nil {
		st.R2TempCredentials = r2TempCredentialsField
	}
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *generateTemporaryTableCredentialResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st generateTemporaryTableCredentialResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Gets the metastore assignment for a workspace
type GetAccountMetastoreAssignmentRequest struct {
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func getAccountMetastoreAssignmentRequestToPb(st *GetAccountMetastoreAssignmentRequest) (*getAccountMetastoreAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountMetastoreAssignmentRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func (st *GetAccountMetastoreAssignmentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAccountMetastoreAssignmentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAccountMetastoreAssignmentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAccountMetastoreAssignmentRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAccountMetastoreAssignmentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getAccountMetastoreAssignmentRequestPb struct {
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func getAccountMetastoreAssignmentRequestFromPb(pb *getAccountMetastoreAssignmentRequestPb) (*GetAccountMetastoreAssignmentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountMetastoreAssignmentRequest{}
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

// Get a metastore
type GetAccountMetastoreRequest struct {
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
}

func getAccountMetastoreRequestToPb(st *GetAccountMetastoreRequest) (*getAccountMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountMetastoreRequestPb{}
	pb.MetastoreId = st.MetastoreId

	return pb, nil
}

func (st *GetAccountMetastoreRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAccountMetastoreRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAccountMetastoreRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAccountMetastoreRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAccountMetastoreRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getAccountMetastoreRequestPb struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
}

func getAccountMetastoreRequestFromPb(pb *getAccountMetastoreRequestPb) (*GetAccountMetastoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountMetastoreRequest{}
	st.MetastoreId = pb.MetastoreId

	return st, nil
}

// Gets the named storage credential
type GetAccountStorageCredentialRequest struct {
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Name of the storage credential.
	// Wire name: 'storage_credential_name'
	StorageCredentialName string `tf:"-"`
}

func getAccountStorageCredentialRequestToPb(st *GetAccountStorageCredentialRequest) (*getAccountStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountStorageCredentialRequestPb{}
	pb.MetastoreId = st.MetastoreId

	pb.StorageCredentialName = st.StorageCredentialName

	return pb, nil
}

func (st *GetAccountStorageCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAccountStorageCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAccountStorageCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAccountStorageCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAccountStorageCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getAccountStorageCredentialRequestPb struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
	// Name of the storage credential.
	StorageCredentialName string `json:"-" url:"-"`
}

func getAccountStorageCredentialRequestFromPb(pb *getAccountStorageCredentialRequestPb) (*GetAccountStorageCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountStorageCredentialRequest{}
	st.MetastoreId = pb.MetastoreId
	st.StorageCredentialName = pb.StorageCredentialName

	return st, nil
}

// Get an artifact allowlist
type GetArtifactAllowlistRequest struct {
	// The artifact type of the allowlist.
	// Wire name: 'artifact_type'
	ArtifactType ArtifactType `tf:"-"`
}

func getArtifactAllowlistRequestToPb(st *GetArtifactAllowlistRequest) (*getArtifactAllowlistRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getArtifactAllowlistRequestPb{}
	pb.ArtifactType = st.ArtifactType

	return pb, nil
}

func (st *GetArtifactAllowlistRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getArtifactAllowlistRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getArtifactAllowlistRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetArtifactAllowlistRequest) MarshalJSON() ([]byte, error) {
	pb, err := getArtifactAllowlistRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getArtifactAllowlistRequestPb struct {
	// The artifact type of the allowlist.
	ArtifactType ArtifactType `json:"-" url:"-"`
}

func getArtifactAllowlistRequestFromPb(pb *getArtifactAllowlistRequestPb) (*GetArtifactAllowlistRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetArtifactAllowlistRequest{}
	st.ArtifactType = pb.ArtifactType

	return st, nil
}

// Get securable workspace bindings
type GetBindingsRequest struct {
	// Maximum number of workspace bindings to return. - When set to 0, the page
	// length is set to a server configured value (recommended); - When set to a
	// value greater than 0, the page length is the minimum of this value and a
	// server configured value; - When set to a value less than 0, an invalid
	// parameter error is returned; - If not set, all the workspace bindings are
	// returned (not recommended).
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// The name of the securable.
	// Wire name: 'securable_name'
	SecurableName string `tf:"-"`
	// The type of the securable to bind to a workspace.
	// Wire name: 'securable_type'
	SecurableType GetBindingsSecurableType `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func getBindingsRequestToPb(st *GetBindingsRequest) (*getBindingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getBindingsRequestPb{}
	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.SecurableName = st.SecurableName

	pb.SecurableType = st.SecurableType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetBindingsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getBindingsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getBindingsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetBindingsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getBindingsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getBindingsRequestPb struct {
	// Maximum number of workspace bindings to return. - When set to 0, the page
	// length is set to a server configured value (recommended); - When set to a
	// value greater than 0, the page length is the minimum of this value and a
	// server configured value; - When set to a value less than 0, an invalid
	// parameter error is returned; - If not set, all the workspace bindings are
	// returned (not recommended).
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The name of the securable.
	SecurableName string `json:"-" url:"-"`
	// The type of the securable to bind to a workspace.
	SecurableType GetBindingsSecurableType `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getBindingsRequestFromPb(pb *getBindingsRequestPb) (*GetBindingsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetBindingsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.SecurableName = pb.SecurableName
	st.SecurableType = pb.SecurableType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getBindingsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getBindingsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetBindingsSecurableType string
type getBindingsSecurableTypePb string

const GetBindingsSecurableTypeCatalog GetBindingsSecurableType = `catalog`

const GetBindingsSecurableTypeCredential GetBindingsSecurableType = `credential`

const GetBindingsSecurableTypeExternalLocation GetBindingsSecurableType = `external_location`

const GetBindingsSecurableTypeStorageCredential GetBindingsSecurableType = `storage_credential`

// String representation for [fmt.Print]
func (f *GetBindingsSecurableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetBindingsSecurableType) Set(v string) error {
	switch v {
	case `catalog`, `credential`, `external_location`, `storage_credential`:
		*f = GetBindingsSecurableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "catalog", "credential", "external_location", "storage_credential"`, v)
	}
}

// Type always returns GetBindingsSecurableType to satisfy [pflag.Value] interface
func (f *GetBindingsSecurableType) Type() string {
	return "GetBindingsSecurableType"
}

func getBindingsSecurableTypeToPb(st *GetBindingsSecurableType) (*getBindingsSecurableTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := getBindingsSecurableTypePb(*st)
	return &pb, nil
}

func getBindingsSecurableTypeFromPb(pb *getBindingsSecurableTypePb) (*GetBindingsSecurableType, error) {
	if pb == nil {
		return nil, nil
	}
	st := GetBindingsSecurableType(*pb)
	return &st, nil
}

// Get Model Version By Alias
type GetByAliasRequest struct {
	// The name of the alias
	// Wire name: 'alias'
	Alias string `tf:"-"`
	// The three-level (fully qualified) name of the registered model
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// Whether to include aliases associated with the model version in the
	// response
	// Wire name: 'include_aliases'
	IncludeAliases bool `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func getByAliasRequestToPb(st *GetByAliasRequest) (*getByAliasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getByAliasRequestPb{}
	pb.Alias = st.Alias

	pb.FullName = st.FullName

	pb.IncludeAliases = st.IncludeAliases

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetByAliasRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getByAliasRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getByAliasRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetByAliasRequest) MarshalJSON() ([]byte, error) {
	pb, err := getByAliasRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getByAliasRequestPb struct {
	// The name of the alias
	Alias string `json:"-" url:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
	// Whether to include aliases associated with the model version in the
	// response
	IncludeAliases bool `json:"-" url:"include_aliases,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getByAliasRequestFromPb(pb *getByAliasRequestPb) (*GetByAliasRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetByAliasRequest{}
	st.Alias = pb.Alias
	st.FullName = pb.FullName
	st.IncludeAliases = pb.IncludeAliases

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getByAliasRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getByAliasRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get a catalog
type GetCatalogRequest struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// The name of the catalog.
	// Wire name: 'name'
	Name string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func getCatalogRequestToPb(st *GetCatalogRequest) (*getCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCatalogRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetCatalogRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCatalogRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCatalogRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCatalogRequest) MarshalJSON() ([]byte, error) {
	pb, err := getCatalogRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getCatalogRequestPb struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// The name of the catalog.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getCatalogRequestFromPb(pb *getCatalogRequestPb) (*GetCatalogRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCatalogRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getCatalogRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getCatalogRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get a connection
type GetConnectionRequest struct {
	// Name of the connection.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func getConnectionRequestToPb(st *GetConnectionRequest) (*getConnectionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getConnectionRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *GetConnectionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getConnectionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getConnectionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetConnectionRequest) MarshalJSON() ([]byte, error) {
	pb, err := getConnectionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getConnectionRequestPb struct {
	// Name of the connection.
	Name string `json:"-" url:"-"`
}

func getConnectionRequestFromPb(pb *getConnectionRequestPb) (*GetConnectionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetConnectionRequest{}
	st.Name = pb.Name

	return st, nil
}

// Get a credential
type GetCredentialRequest struct {
	// Name of the credential.
	// Wire name: 'name_arg'
	NameArg string `tf:"-"`
}

func getCredentialRequestToPb(st *GetCredentialRequest) (*getCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCredentialRequestPb{}
	pb.NameArg = st.NameArg

	return pb, nil
}

func (st *GetCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := getCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getCredentialRequestPb struct {
	// Name of the credential.
	NameArg string `json:"-" url:"-"`
}

func getCredentialRequestFromPb(pb *getCredentialRequestPb) (*GetCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCredentialRequest{}
	st.NameArg = pb.NameArg

	return st, nil
}

// Get effective permissions
type GetEffectiveRequest struct {
	// Full name of securable.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// If provided, only the effective permissions for the specified principal
	// (user or group) are returned.
	// Wire name: 'principal'
	Principal string `tf:"-"`
	// Type of securable.
	// Wire name: 'securable_type'
	SecurableType SecurableType `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func getEffectiveRequestToPb(st *GetEffectiveRequest) (*getEffectiveRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEffectiveRequestPb{}
	pb.FullName = st.FullName

	pb.Principal = st.Principal

	pb.SecurableType = st.SecurableType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetEffectiveRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getEffectiveRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getEffectiveRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetEffectiveRequest) MarshalJSON() ([]byte, error) {
	pb, err := getEffectiveRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getEffectiveRequestPb struct {
	// Full name of securable.
	FullName string `json:"-" url:"-"`
	// If provided, only the effective permissions for the specified principal
	// (user or group) are returned.
	Principal string `json:"-" url:"principal,omitempty"`
	// Type of securable.
	SecurableType SecurableType `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getEffectiveRequestFromPb(pb *getEffectiveRequestPb) (*GetEffectiveRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEffectiveRequest{}
	st.FullName = pb.FullName
	st.Principal = pb.Principal
	st.SecurableType = pb.SecurableType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getEffectiveRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getEffectiveRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get an external location
type GetExternalLocationRequest struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// Name of the external location.
	// Wire name: 'name'
	Name string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func getExternalLocationRequestToPb(st *GetExternalLocationRequest) (*getExternalLocationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExternalLocationRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetExternalLocationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExternalLocationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExternalLocationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExternalLocationRequest) MarshalJSON() ([]byte, error) {
	pb, err := getExternalLocationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getExternalLocationRequestPb struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Name of the external location.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getExternalLocationRequestFromPb(pb *getExternalLocationRequestPb) (*GetExternalLocationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExternalLocationRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getExternalLocationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getExternalLocationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get a function
type GetFunctionRequest struct {
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	// Wire name: 'name'
	Name string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func getFunctionRequestToPb(st *GetFunctionRequest) (*getFunctionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getFunctionRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetFunctionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getFunctionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getFunctionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetFunctionRequest) MarshalJSON() ([]byte, error) {
	pb, err := getFunctionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getFunctionRequestPb struct {
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getFunctionRequestFromPb(pb *getFunctionRequestPb) (*GetFunctionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetFunctionRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getFunctionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getFunctionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get permissions
type GetGrantRequest struct {
	// Full name of securable.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// If provided, only the permissions for the specified principal (user or
	// group) are returned.
	// Wire name: 'principal'
	Principal string `tf:"-"`
	// Type of securable.
	// Wire name: 'securable_type'
	SecurableType SecurableType `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func getGrantRequestToPb(st *GetGrantRequest) (*getGrantRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getGrantRequestPb{}
	pb.FullName = st.FullName

	pb.Principal = st.Principal

	pb.SecurableType = st.SecurableType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetGrantRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getGrantRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getGrantRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetGrantRequest) MarshalJSON() ([]byte, error) {
	pb, err := getGrantRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getGrantRequestPb struct {
	// Full name of securable.
	FullName string `json:"-" url:"-"`
	// If provided, only the permissions for the specified principal (user or
	// group) are returned.
	Principal string `json:"-" url:"principal,omitempty"`
	// Type of securable.
	SecurableType SecurableType `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getGrantRequestFromPb(pb *getGrantRequestPb) (*GetGrantRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetGrantRequest{}
	st.FullName = pb.FullName
	st.Principal = pb.Principal
	st.SecurableType = pb.SecurableType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getGrantRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getGrantRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get a metastore
type GetMetastoreRequest struct {
	// Unique ID of the metastore.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func getMetastoreRequestToPb(st *GetMetastoreRequest) (*getMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getMetastoreRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func (st *GetMetastoreRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getMetastoreRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getMetastoreRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetMetastoreRequest) MarshalJSON() ([]byte, error) {
	pb, err := getMetastoreRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getMetastoreRequestPb struct {
	// Unique ID of the metastore.
	Id string `json:"-" url:"-"`
}

func getMetastoreRequestFromPb(pb *getMetastoreRequestPb) (*GetMetastoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetMetastoreRequest{}
	st.Id = pb.Id

	return st, nil
}

type GetMetastoreSummaryResponse struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	// Wire name: 'cloud'
	Cloud string
	// Time at which this metastore was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of metastore creator.
	// Wire name: 'created_by'
	CreatedBy string
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	// Wire name: 'default_data_access_config_id'
	DefaultDataAccessConfigId string
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	// Wire name: 'delta_sharing_organization_name'
	DeltaSharingOrganizationName string
	// The lifetime of delta sharing recipient token in seconds.
	// Wire name: 'delta_sharing_recipient_token_lifetime_in_seconds'
	DeltaSharingRecipientTokenLifetimeInSeconds int64
	// The scope of Delta Sharing enabled for the metastore.
	// Wire name: 'delta_sharing_scope'
	DeltaSharingScope GetMetastoreSummaryResponseDeltaSharingScope
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	// Wire name: 'external_access_enabled'
	ExternalAccessEnabled bool
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	// Wire name: 'global_metastore_id'
	GlobalMetastoreId string
	// Unique identifier of metastore.
	// Wire name: 'metastore_id'
	MetastoreId string
	// The user-specified name of the metastore.
	// Wire name: 'name'
	Name string
	// The owner of the metastore.
	// Wire name: 'owner'
	Owner string
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	// Wire name: 'privilege_model_version'
	PrivilegeModelVersion string
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	// Wire name: 'region'
	Region string
	// The storage root URL for metastore
	// Wire name: 'storage_root'
	StorageRoot string
	// UUID of storage credential to access the metastore storage_root.
	// Wire name: 'storage_root_credential_id'
	StorageRootCredentialId string
	// Name of the storage credential to access the metastore storage_root.
	// Wire name: 'storage_root_credential_name'
	StorageRootCredentialName string
	// Time at which the metastore was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// Username of user who last modified the metastore.
	// Wire name: 'updated_by'
	UpdatedBy string

	ForceSendFields []string `tf:"-"`
}

func getMetastoreSummaryResponseToPb(st *GetMetastoreSummaryResponse) (*getMetastoreSummaryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getMetastoreSummaryResponsePb{}
	pb.Cloud = st.Cloud

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	pb.DefaultDataAccessConfigId = st.DefaultDataAccessConfigId

	pb.DeltaSharingOrganizationName = st.DeltaSharingOrganizationName

	pb.DeltaSharingRecipientTokenLifetimeInSeconds = st.DeltaSharingRecipientTokenLifetimeInSeconds

	pb.DeltaSharingScope = st.DeltaSharingScope

	pb.ExternalAccessEnabled = st.ExternalAccessEnabled

	pb.GlobalMetastoreId = st.GlobalMetastoreId

	pb.MetastoreId = st.MetastoreId

	pb.Name = st.Name

	pb.Owner = st.Owner

	pb.PrivilegeModelVersion = st.PrivilegeModelVersion

	pb.Region = st.Region

	pb.StorageRoot = st.StorageRoot

	pb.StorageRootCredentialId = st.StorageRootCredentialId

	pb.StorageRootCredentialName = st.StorageRootCredentialName

	pb.UpdatedAt = st.UpdatedAt

	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetMetastoreSummaryResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getMetastoreSummaryResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getMetastoreSummaryResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetMetastoreSummaryResponse) MarshalJSON() ([]byte, error) {
	pb, err := getMetastoreSummaryResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getMetastoreSummaryResponsePb struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud string `json:"cloud,omitempty"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of metastore creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope GetMetastoreSummaryResponseDeltaSharingScope `json:"delta_sharing_scope,omitempty"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled bool `json:"external_access_enabled,omitempty"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId string `json:"global_metastore_id,omitempty"`
	// Unique identifier of metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The user-specified name of the metastore.
	Name string `json:"name,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	StorageRoot string `json:"storage_root,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName string `json:"storage_root_credential_name,omitempty"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the metastore.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getMetastoreSummaryResponseFromPb(pb *getMetastoreSummaryResponsePb) (*GetMetastoreSummaryResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetMetastoreSummaryResponse{}
	st.Cloud = pb.Cloud
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.DefaultDataAccessConfigId = pb.DefaultDataAccessConfigId
	st.DeltaSharingOrganizationName = pb.DeltaSharingOrganizationName
	st.DeltaSharingRecipientTokenLifetimeInSeconds = pb.DeltaSharingRecipientTokenLifetimeInSeconds
	st.DeltaSharingScope = pb.DeltaSharingScope
	st.ExternalAccessEnabled = pb.ExternalAccessEnabled
	st.GlobalMetastoreId = pb.GlobalMetastoreId
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.PrivilegeModelVersion = pb.PrivilegeModelVersion
	st.Region = pb.Region
	st.StorageRoot = pb.StorageRoot
	st.StorageRootCredentialId = pb.StorageRootCredentialId
	st.StorageRootCredentialName = pb.StorageRootCredentialName
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getMetastoreSummaryResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getMetastoreSummaryResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The scope of Delta Sharing enabled for the metastore.
type GetMetastoreSummaryResponseDeltaSharingScope string
type getMetastoreSummaryResponseDeltaSharingScopePb string

const GetMetastoreSummaryResponseDeltaSharingScopeInternal GetMetastoreSummaryResponseDeltaSharingScope = `INTERNAL`

const GetMetastoreSummaryResponseDeltaSharingScopeInternalAndExternal GetMetastoreSummaryResponseDeltaSharingScope = `INTERNAL_AND_EXTERNAL`

// String representation for [fmt.Print]
func (f *GetMetastoreSummaryResponseDeltaSharingScope) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetMetastoreSummaryResponseDeltaSharingScope) Set(v string) error {
	switch v {
	case `INTERNAL`, `INTERNAL_AND_EXTERNAL`:
		*f = GetMetastoreSummaryResponseDeltaSharingScope(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INTERNAL", "INTERNAL_AND_EXTERNAL"`, v)
	}
}

// Type always returns GetMetastoreSummaryResponseDeltaSharingScope to satisfy [pflag.Value] interface
func (f *GetMetastoreSummaryResponseDeltaSharingScope) Type() string {
	return "GetMetastoreSummaryResponseDeltaSharingScope"
}

func getMetastoreSummaryResponseDeltaSharingScopeToPb(st *GetMetastoreSummaryResponseDeltaSharingScope) (*getMetastoreSummaryResponseDeltaSharingScopePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := getMetastoreSummaryResponseDeltaSharingScopePb(*st)
	return &pb, nil
}

func getMetastoreSummaryResponseDeltaSharingScopeFromPb(pb *getMetastoreSummaryResponseDeltaSharingScopePb) (*GetMetastoreSummaryResponseDeltaSharingScope, error) {
	if pb == nil {
		return nil, nil
	}
	st := GetMetastoreSummaryResponseDeltaSharingScope(*pb)
	return &st, nil
}

// Get a Model Version
type GetModelVersionRequest struct {
	// The three-level (fully qualified) name of the model version
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// Whether to include aliases associated with the model version in the
	// response
	// Wire name: 'include_aliases'
	IncludeAliases bool `tf:"-"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// The integer version number of the model version
	// Wire name: 'version'
	Version int `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func getModelVersionRequestToPb(st *GetModelVersionRequest) (*getModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getModelVersionRequestPb{}
	pb.FullName = st.FullName

	pb.IncludeAliases = st.IncludeAliases

	pb.IncludeBrowse = st.IncludeBrowse

	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetModelVersionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getModelVersionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getModelVersionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetModelVersionRequest) MarshalJSON() ([]byte, error) {
	pb, err := getModelVersionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getModelVersionRequestPb struct {
	// The three-level (fully qualified) name of the model version
	FullName string `json:"-" url:"-"`
	// Whether to include aliases associated with the model version in the
	// response
	IncludeAliases bool `json:"-" url:"include_aliases,omitempty"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// The integer version number of the model version
	Version int `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getModelVersionRequestFromPb(pb *getModelVersionRequestPb) (*GetModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelVersionRequest{}
	st.FullName = pb.FullName
	st.IncludeAliases = pb.IncludeAliases
	st.IncludeBrowse = pb.IncludeBrowse
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getModelVersionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getModelVersionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get an Online Table
type GetOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func getOnlineTableRequestToPb(st *GetOnlineTableRequest) (*getOnlineTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getOnlineTableRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *GetOnlineTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getOnlineTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getOnlineTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetOnlineTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := getOnlineTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getOnlineTableRequestPb struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"-" url:"-"`
}

func getOnlineTableRequestFromPb(pb *getOnlineTableRequestPb) (*GetOnlineTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetOnlineTableRequest{}
	st.Name = pb.Name

	return st, nil
}

// Get a table monitor
type GetQualityMonitorRequest struct {
	// Full name of the table.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func getQualityMonitorRequestToPb(st *GetQualityMonitorRequest) (*getQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getQualityMonitorRequestPb{}
	pb.TableName = st.TableName

	return pb, nil
}

func (st *GetQualityMonitorRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getQualityMonitorRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getQualityMonitorRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetQualityMonitorRequest) MarshalJSON() ([]byte, error) {
	pb, err := getQualityMonitorRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getQualityMonitorRequestPb struct {
	// Full name of the table.
	TableName string `json:"-" url:"-"`
}

func getQualityMonitorRequestFromPb(pb *getQualityMonitorRequestPb) (*GetQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetQualityMonitorRequest{}
	st.TableName = pb.TableName

	return st, nil
}

// Get information for a single resource quota.
type GetQuotaRequest struct {
	// Full name of the parent resource. Provide the metastore ID if the parent
	// is a metastore.
	// Wire name: 'parent_full_name'
	ParentFullName string `tf:"-"`
	// Securable type of the quota parent.
	// Wire name: 'parent_securable_type'
	ParentSecurableType string `tf:"-"`
	// Name of the quota. Follows the pattern of the quota type, with "-quota"
	// added as a suffix.
	// Wire name: 'quota_name'
	QuotaName string `tf:"-"`
}

func getQuotaRequestToPb(st *GetQuotaRequest) (*getQuotaRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getQuotaRequestPb{}
	pb.ParentFullName = st.ParentFullName

	pb.ParentSecurableType = st.ParentSecurableType

	pb.QuotaName = st.QuotaName

	return pb, nil
}

func (st *GetQuotaRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getQuotaRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getQuotaRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetQuotaRequest) MarshalJSON() ([]byte, error) {
	pb, err := getQuotaRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getQuotaRequestPb struct {
	// Full name of the parent resource. Provide the metastore ID if the parent
	// is a metastore.
	ParentFullName string `json:"-" url:"-"`
	// Securable type of the quota parent.
	ParentSecurableType string `json:"-" url:"-"`
	// Name of the quota. Follows the pattern of the quota type, with "-quota"
	// added as a suffix.
	QuotaName string `json:"-" url:"-"`
}

func getQuotaRequestFromPb(pb *getQuotaRequestPb) (*GetQuotaRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetQuotaRequest{}
	st.ParentFullName = pb.ParentFullName
	st.ParentSecurableType = pb.ParentSecurableType
	st.QuotaName = pb.QuotaName

	return st, nil
}

type GetQuotaResponse struct {
	// The returned QuotaInfo.
	// Wire name: 'quota_info'
	QuotaInfo *QuotaInfo
}

func getQuotaResponseToPb(st *GetQuotaResponse) (*getQuotaResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getQuotaResponsePb{}
	quotaInfoPb, err := quotaInfoToPb(st.QuotaInfo)
	if err != nil {
		return nil, err
	}
	if quotaInfoPb != nil {
		pb.QuotaInfo = quotaInfoPb
	}

	return pb, nil
}

func (st *GetQuotaResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getQuotaResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getQuotaResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetQuotaResponse) MarshalJSON() ([]byte, error) {
	pb, err := getQuotaResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getQuotaResponsePb struct {
	// The returned QuotaInfo.
	QuotaInfo *quotaInfoPb `json:"quota_info,omitempty"`
}

func getQuotaResponseFromPb(pb *getQuotaResponsePb) (*GetQuotaResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetQuotaResponse{}
	quotaInfoField, err := quotaInfoFromPb(pb.QuotaInfo)
	if err != nil {
		return nil, err
	}
	if quotaInfoField != nil {
		st.QuotaInfo = quotaInfoField
	}

	return st, nil
}

// Get refresh
type GetRefreshRequest struct {
	// ID of the refresh.
	// Wire name: 'refresh_id'
	RefreshId string `tf:"-"`
	// Full name of the table.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func getRefreshRequestToPb(st *GetRefreshRequest) (*getRefreshRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRefreshRequestPb{}
	pb.RefreshId = st.RefreshId

	pb.TableName = st.TableName

	return pb, nil
}

func (st *GetRefreshRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRefreshRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRefreshRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRefreshRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRefreshRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getRefreshRequestPb struct {
	// ID of the refresh.
	RefreshId string `json:"-" url:"-"`
	// Full name of the table.
	TableName string `json:"-" url:"-"`
}

func getRefreshRequestFromPb(pb *getRefreshRequestPb) (*GetRefreshRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRefreshRequest{}
	st.RefreshId = pb.RefreshId
	st.TableName = pb.TableName

	return st, nil
}

// Get a Registered Model
type GetRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// Whether to include registered model aliases in the response
	// Wire name: 'include_aliases'
	IncludeAliases bool `tf:"-"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func getRegisteredModelRequestToPb(st *GetRegisteredModelRequest) (*getRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRegisteredModelRequestPb{}
	pb.FullName = st.FullName

	pb.IncludeAliases = st.IncludeAliases

	pb.IncludeBrowse = st.IncludeBrowse

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRegisteredModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRegisteredModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRegisteredModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getRegisteredModelRequestPb struct {
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
	// Whether to include registered model aliases in the response
	IncludeAliases bool `json:"-" url:"include_aliases,omitempty"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getRegisteredModelRequestFromPb(pb *getRegisteredModelRequestPb) (*GetRegisteredModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRegisteredModelRequest{}
	st.FullName = pb.FullName
	st.IncludeAliases = pb.IncludeAliases
	st.IncludeBrowse = pb.IncludeBrowse

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getRegisteredModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getRegisteredModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get a schema
type GetSchemaRequest struct {
	// Full name of the schema.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func getSchemaRequestToPb(st *GetSchemaRequest) (*getSchemaRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getSchemaRequestPb{}
	pb.FullName = st.FullName

	pb.IncludeBrowse = st.IncludeBrowse

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetSchemaRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getSchemaRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getSchemaRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetSchemaRequest) MarshalJSON() ([]byte, error) {
	pb, err := getSchemaRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getSchemaRequestPb struct {
	// Full name of the schema.
	FullName string `json:"-" url:"-"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getSchemaRequestFromPb(pb *getSchemaRequestPb) (*GetSchemaRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSchemaRequest{}
	st.FullName = pb.FullName
	st.IncludeBrowse = pb.IncludeBrowse

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getSchemaRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getSchemaRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get a credential
type GetStorageCredentialRequest struct {
	// Name of the storage credential.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func getStorageCredentialRequestToPb(st *GetStorageCredentialRequest) (*getStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStorageCredentialRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *GetStorageCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getStorageCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getStorageCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetStorageCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := getStorageCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getStorageCredentialRequestPb struct {
	// Name of the storage credential.
	Name string `json:"-" url:"-"`
}

func getStorageCredentialRequestFromPb(pb *getStorageCredentialRequestPb) (*GetStorageCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStorageCredentialRequest{}
	st.Name = pb.Name

	return st, nil
}

// Get a table
type GetTableRequest struct {
	// Full name of the table.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// Whether delta metadata should be included in the response.
	// Wire name: 'include_delta_metadata'
	IncludeDeltaMetadata bool `tf:"-"`
	// Whether to include a manifest containing capabilities the table has.
	// Wire name: 'include_manifest_capabilities'
	IncludeManifestCapabilities bool `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func getTableRequestToPb(st *GetTableRequest) (*getTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getTableRequestPb{}
	pb.FullName = st.FullName

	pb.IncludeBrowse = st.IncludeBrowse

	pb.IncludeDeltaMetadata = st.IncludeDeltaMetadata

	pb.IncludeManifestCapabilities = st.IncludeManifestCapabilities

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := getTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getTableRequestPb struct {
	// Full name of the table.
	FullName string `json:"-" url:"-"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata bool `json:"-" url:"include_delta_metadata,omitempty"`
	// Whether to include a manifest containing capabilities the table has.
	IncludeManifestCapabilities bool `json:"-" url:"include_manifest_capabilities,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getTableRequestFromPb(pb *getTableRequestPb) (*GetTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetTableRequest{}
	st.FullName = pb.FullName
	st.IncludeBrowse = pb.IncludeBrowse
	st.IncludeDeltaMetadata = pb.IncludeDeltaMetadata
	st.IncludeManifestCapabilities = pb.IncludeManifestCapabilities

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getTableRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getTableRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get catalog workspace bindings
type GetWorkspaceBindingRequest struct {
	// The name of the catalog.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func getWorkspaceBindingRequestToPb(st *GetWorkspaceBindingRequest) (*getWorkspaceBindingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceBindingRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *GetWorkspaceBindingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getWorkspaceBindingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getWorkspaceBindingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetWorkspaceBindingRequest) MarshalJSON() ([]byte, error) {
	pb, err := getWorkspaceBindingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getWorkspaceBindingRequestPb struct {
	// The name of the catalog.
	Name string `json:"-" url:"-"`
}

func getWorkspaceBindingRequestFromPb(pb *getWorkspaceBindingRequestPb) (*GetWorkspaceBindingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceBindingRequest{}
	st.Name = pb.Name

	return st, nil
}

type IsolationMode string
type isolationModePb string

const IsolationModeIsolationModeIsolated IsolationMode = `ISOLATION_MODE_ISOLATED`

const IsolationModeIsolationModeOpen IsolationMode = `ISOLATION_MODE_OPEN`

// String representation for [fmt.Print]
func (f *IsolationMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *IsolationMode) Set(v string) error {
	switch v {
	case `ISOLATION_MODE_ISOLATED`, `ISOLATION_MODE_OPEN`:
		*f = IsolationMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ISOLATION_MODE_ISOLATED", "ISOLATION_MODE_OPEN"`, v)
	}
}

// Type always returns IsolationMode to satisfy [pflag.Value] interface
func (f *IsolationMode) Type() string {
	return "IsolationMode"
}

func isolationModeToPb(st *IsolationMode) (*isolationModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := isolationModePb(*st)
	return &pb, nil
}

func isolationModeFromPb(pb *isolationModePb) (*IsolationMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := IsolationMode(*pb)
	return &st, nil
}

// Get all workspaces assigned to a metastore
type ListAccountMetastoreAssignmentsRequest struct {
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
}

func listAccountMetastoreAssignmentsRequestToPb(st *ListAccountMetastoreAssignmentsRequest) (*listAccountMetastoreAssignmentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountMetastoreAssignmentsRequestPb{}
	pb.MetastoreId = st.MetastoreId

	return pb, nil
}

func (st *ListAccountMetastoreAssignmentsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAccountMetastoreAssignmentsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAccountMetastoreAssignmentsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAccountMetastoreAssignmentsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listAccountMetastoreAssignmentsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listAccountMetastoreAssignmentsRequestPb struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
}

func listAccountMetastoreAssignmentsRequestFromPb(pb *listAccountMetastoreAssignmentsRequestPb) (*ListAccountMetastoreAssignmentsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountMetastoreAssignmentsRequest{}
	st.MetastoreId = pb.MetastoreId

	return st, nil
}

// The list of workspaces to which the given metastore is assigned.
type ListAccountMetastoreAssignmentsResponse struct {

	// Wire name: 'workspace_ids'
	WorkspaceIds []int64
}

func listAccountMetastoreAssignmentsResponseToPb(st *ListAccountMetastoreAssignmentsResponse) (*listAccountMetastoreAssignmentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountMetastoreAssignmentsResponsePb{}
	pb.WorkspaceIds = st.WorkspaceIds

	return pb, nil
}

func (st *ListAccountMetastoreAssignmentsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAccountMetastoreAssignmentsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAccountMetastoreAssignmentsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAccountMetastoreAssignmentsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listAccountMetastoreAssignmentsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listAccountMetastoreAssignmentsResponsePb struct {
	WorkspaceIds []int64 `json:"workspace_ids,omitempty"`
}

func listAccountMetastoreAssignmentsResponseFromPb(pb *listAccountMetastoreAssignmentsResponsePb) (*ListAccountMetastoreAssignmentsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountMetastoreAssignmentsResponse{}
	st.WorkspaceIds = pb.WorkspaceIds

	return st, nil
}

// Get all storage credentials assigned to a metastore
type ListAccountStorageCredentialsRequest struct {
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
}

func listAccountStorageCredentialsRequestToPb(st *ListAccountStorageCredentialsRequest) (*listAccountStorageCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountStorageCredentialsRequestPb{}
	pb.MetastoreId = st.MetastoreId

	return pb, nil
}

func (st *ListAccountStorageCredentialsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAccountStorageCredentialsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAccountStorageCredentialsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAccountStorageCredentialsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listAccountStorageCredentialsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listAccountStorageCredentialsRequestPb struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" url:"-"`
}

func listAccountStorageCredentialsRequestFromPb(pb *listAccountStorageCredentialsRequestPb) (*ListAccountStorageCredentialsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountStorageCredentialsRequest{}
	st.MetastoreId = pb.MetastoreId

	return st, nil
}

type ListAccountStorageCredentialsResponse struct {
	// An array of metastore storage credentials.
	// Wire name: 'storage_credentials'
	StorageCredentials []StorageCredentialInfo
}

func listAccountStorageCredentialsResponseToPb(st *ListAccountStorageCredentialsResponse) (*listAccountStorageCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountStorageCredentialsResponsePb{}

	var storageCredentialsPb []storageCredentialInfoPb
	for _, item := range st.StorageCredentials {
		itemPb, err := storageCredentialInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			storageCredentialsPb = append(storageCredentialsPb, *itemPb)
		}
	}
	pb.StorageCredentials = storageCredentialsPb

	return pb, nil
}

func (st *ListAccountStorageCredentialsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAccountStorageCredentialsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAccountStorageCredentialsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAccountStorageCredentialsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listAccountStorageCredentialsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listAccountStorageCredentialsResponsePb struct {
	// An array of metastore storage credentials.
	StorageCredentials []storageCredentialInfoPb `json:"storage_credentials,omitempty"`
}

func listAccountStorageCredentialsResponseFromPb(pb *listAccountStorageCredentialsResponsePb) (*ListAccountStorageCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountStorageCredentialsResponse{}

	var storageCredentialsField []StorageCredentialInfo
	for _, itemPb := range pb.StorageCredentials {
		item, err := storageCredentialInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			storageCredentialsField = append(storageCredentialsField, *item)
		}
	}
	st.StorageCredentials = storageCredentialsField

	return st, nil
}

// List catalogs
type ListCatalogsRequest struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// Maximum number of catalogs to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid catalogs are returned (not
	// recommended). - Note: The number of returned catalogs might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further catalogs can be fetched is when the next_page_token is
	// unset from the response.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listCatalogsRequestToPb(st *ListCatalogsRequest) (*listCatalogsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCatalogsRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse

	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListCatalogsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listCatalogsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listCatalogsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListCatalogsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listCatalogsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listCatalogsRequestPb struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Maximum number of catalogs to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid catalogs are returned (not
	// recommended). - Note: The number of returned catalogs might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further catalogs can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCatalogsRequestFromPb(pb *listCatalogsRequestPb) (*ListCatalogsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCatalogsRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCatalogsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCatalogsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCatalogsResponse struct {
	// An array of catalog information objects.
	// Wire name: 'catalogs'
	Catalogs []CatalogInfo
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listCatalogsResponseToPb(st *ListCatalogsResponse) (*listCatalogsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCatalogsResponsePb{}

	var catalogsPb []catalogInfoPb
	for _, item := range st.Catalogs {
		itemPb, err := catalogInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			catalogsPb = append(catalogsPb, *itemPb)
		}
	}
	pb.Catalogs = catalogsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListCatalogsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listCatalogsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listCatalogsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListCatalogsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listCatalogsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listCatalogsResponsePb struct {
	// An array of catalog information objects.
	Catalogs []catalogInfoPb `json:"catalogs,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCatalogsResponseFromPb(pb *listCatalogsResponsePb) (*ListCatalogsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCatalogsResponse{}

	var catalogsField []CatalogInfo
	for _, itemPb := range pb.Catalogs {
		item, err := catalogInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			catalogsField = append(catalogsField, *item)
		}
	}
	st.Catalogs = catalogsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCatalogsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCatalogsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List connections
type ListConnectionsRequest struct {
	// Maximum number of connections to return. - If not set, all connections
	// are returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listConnectionsRequestToPb(st *ListConnectionsRequest) (*listConnectionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listConnectionsRequestPb{}
	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListConnectionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listConnectionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listConnectionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListConnectionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listConnectionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listConnectionsRequestPb struct {
	// Maximum number of connections to return. - If not set, all connections
	// are returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listConnectionsRequestFromPb(pb *listConnectionsRequestPb) (*ListConnectionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListConnectionsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listConnectionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listConnectionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListConnectionsResponse struct {
	// An array of connection information objects.
	// Wire name: 'connections'
	Connections []ConnectionInfo
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listConnectionsResponseToPb(st *ListConnectionsResponse) (*listConnectionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listConnectionsResponsePb{}

	var connectionsPb []connectionInfoPb
	for _, item := range st.Connections {
		itemPb, err := connectionInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			connectionsPb = append(connectionsPb, *itemPb)
		}
	}
	pb.Connections = connectionsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListConnectionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listConnectionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listConnectionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListConnectionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listConnectionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listConnectionsResponsePb struct {
	// An array of connection information objects.
	Connections []connectionInfoPb `json:"connections,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listConnectionsResponseFromPb(pb *listConnectionsResponsePb) (*ListConnectionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListConnectionsResponse{}

	var connectionsField []ConnectionInfo
	for _, itemPb := range pb.Connections {
		item, err := connectionInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			connectionsField = append(connectionsField, *item)
		}
	}
	st.Connections = connectionsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listConnectionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listConnectionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List credentials
type ListCredentialsRequest struct {
	// Maximum number of credentials to return. - If not set, the default max
	// page size is used. - When set to a value greater than 0, the page length
	// is the minimum of this value and a server-configured value. - When set to
	// 0, the page length is set to a server-configured value (recommended). -
	// When set to a value less than 0, an invalid parameter error is returned.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque token to retrieve the next page of results.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// Return only credentials for the specified purpose.
	// Wire name: 'purpose'
	Purpose CredentialPurpose `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listCredentialsRequestToPb(st *ListCredentialsRequest) (*listCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCredentialsRequestPb{}
	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.Purpose = st.Purpose

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListCredentialsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listCredentialsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listCredentialsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListCredentialsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listCredentialsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listCredentialsRequestPb struct {
	// Maximum number of credentials to return. - If not set, the default max
	// page size is used. - When set to a value greater than 0, the page length
	// is the minimum of this value and a server-configured value. - When set to
	// 0, the page length is set to a server-configured value (recommended). -
	// When set to a value less than 0, an invalid parameter error is returned.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque token to retrieve the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Return only credentials for the specified purpose.
	Purpose CredentialPurpose `json:"-" url:"purpose,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCredentialsRequestFromPb(pb *listCredentialsRequestPb) (*ListCredentialsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCredentialsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.Purpose = pb.Purpose

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCredentialsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCredentialsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCredentialsResponse struct {

	// Wire name: 'credentials'
	Credentials []CredentialInfo
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listCredentialsResponseToPb(st *ListCredentialsResponse) (*listCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCredentialsResponsePb{}

	var credentialsPb []credentialInfoPb
	for _, item := range st.Credentials {
		itemPb, err := credentialInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			credentialsPb = append(credentialsPb, *itemPb)
		}
	}
	pb.Credentials = credentialsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListCredentialsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listCredentialsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listCredentialsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListCredentialsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listCredentialsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listCredentialsResponsePb struct {
	Credentials []credentialInfoPb `json:"credentials,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCredentialsResponseFromPb(pb *listCredentialsResponsePb) (*ListCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCredentialsResponse{}

	var credentialsField []CredentialInfo
	for _, itemPb := range pb.Credentials {
		item, err := credentialInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			credentialsField = append(credentialsField, *item)
		}
	}
	st.Credentials = credentialsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCredentialsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCredentialsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List external locations
type ListExternalLocationsRequest struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// Maximum number of external locations to return. If not set, all the
	// external locations are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listExternalLocationsRequestToPb(st *ListExternalLocationsRequest) (*listExternalLocationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExternalLocationsRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse

	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListExternalLocationsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExternalLocationsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExternalLocationsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExternalLocationsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listExternalLocationsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listExternalLocationsRequestPb struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Maximum number of external locations to return. If not set, all the
	// external locations are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExternalLocationsRequestFromPb(pb *listExternalLocationsRequestPb) (*ListExternalLocationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExternalLocationsRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listExternalLocationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listExternalLocationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExternalLocationsResponse struct {
	// An array of external locations.
	// Wire name: 'external_locations'
	ExternalLocations []ExternalLocationInfo
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listExternalLocationsResponseToPb(st *ListExternalLocationsResponse) (*listExternalLocationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExternalLocationsResponsePb{}

	var externalLocationsPb []externalLocationInfoPb
	for _, item := range st.ExternalLocations {
		itemPb, err := externalLocationInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			externalLocationsPb = append(externalLocationsPb, *itemPb)
		}
	}
	pb.ExternalLocations = externalLocationsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListExternalLocationsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExternalLocationsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExternalLocationsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExternalLocationsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listExternalLocationsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listExternalLocationsResponsePb struct {
	// An array of external locations.
	ExternalLocations []externalLocationInfoPb `json:"external_locations,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExternalLocationsResponseFromPb(pb *listExternalLocationsResponsePb) (*ListExternalLocationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExternalLocationsResponse{}

	var externalLocationsField []ExternalLocationInfo
	for _, itemPb := range pb.ExternalLocations {
		item, err := externalLocationInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			externalLocationsField = append(externalLocationsField, *item)
		}
	}
	st.ExternalLocations = externalLocationsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listExternalLocationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listExternalLocationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List functions
type ListFunctionsRequest struct {
	// Name of parent catalog for functions of interest.
	// Wire name: 'catalog_name'
	CatalogName string `tf:"-"`
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// Maximum number of functions to return. If not set, all the functions are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// Parent schema of functions.
	// Wire name: 'schema_name'
	SchemaName string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listFunctionsRequestToPb(st *ListFunctionsRequest) (*listFunctionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFunctionsRequestPb{}
	pb.CatalogName = st.CatalogName

	pb.IncludeBrowse = st.IncludeBrowse

	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListFunctionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listFunctionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listFunctionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListFunctionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listFunctionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listFunctionsRequestPb struct {
	// Name of parent catalog for functions of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Maximum number of functions to return. If not set, all the functions are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent schema of functions.
	SchemaName string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listFunctionsRequestFromPb(pb *listFunctionsRequestPb) (*ListFunctionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFunctionsRequest{}
	st.CatalogName = pb.CatalogName
	st.IncludeBrowse = pb.IncludeBrowse
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.SchemaName = pb.SchemaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listFunctionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listFunctionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListFunctionsResponse struct {
	// An array of function information objects.
	// Wire name: 'functions'
	Functions []FunctionInfo
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listFunctionsResponseToPb(st *ListFunctionsResponse) (*listFunctionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFunctionsResponsePb{}

	var functionsPb []functionInfoPb
	for _, item := range st.Functions {
		itemPb, err := functionInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			functionsPb = append(functionsPb, *itemPb)
		}
	}
	pb.Functions = functionsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListFunctionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listFunctionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listFunctionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListFunctionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listFunctionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listFunctionsResponsePb struct {
	// An array of function information objects.
	Functions []functionInfoPb `json:"functions,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listFunctionsResponseFromPb(pb *listFunctionsResponsePb) (*ListFunctionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFunctionsResponse{}

	var functionsField []FunctionInfo
	for _, itemPb := range pb.Functions {
		item, err := functionInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			functionsField = append(functionsField, *item)
		}
	}
	st.Functions = functionsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listFunctionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listFunctionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListMetastoresResponse struct {
	// An array of metastore information objects.
	// Wire name: 'metastores'
	Metastores []MetastoreInfo
}

func listMetastoresResponseToPb(st *ListMetastoresResponse) (*listMetastoresResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listMetastoresResponsePb{}

	var metastoresPb []metastoreInfoPb
	for _, item := range st.Metastores {
		itemPb, err := metastoreInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			metastoresPb = append(metastoresPb, *itemPb)
		}
	}
	pb.Metastores = metastoresPb

	return pb, nil
}

func (st *ListMetastoresResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listMetastoresResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listMetastoresResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListMetastoresResponse) MarshalJSON() ([]byte, error) {
	pb, err := listMetastoresResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listMetastoresResponsePb struct {
	// An array of metastore information objects.
	Metastores []metastoreInfoPb `json:"metastores,omitempty"`
}

func listMetastoresResponseFromPb(pb *listMetastoresResponsePb) (*ListMetastoresResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListMetastoresResponse{}

	var metastoresField []MetastoreInfo
	for _, itemPb := range pb.Metastores {
		item, err := metastoreInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			metastoresField = append(metastoresField, *item)
		}
	}
	st.Metastores = metastoresField

	return st, nil
}

// List Model Versions
type ListModelVersionsRequest struct {
	// The full three-level name of the registered model under which to list
	// model versions
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// Maximum number of model versions to return. If not set, the page length
	// is set to a server configured value (100, as of 1/3/2024). - when set to
	// a value greater than 0, the page length is the minimum of this value and
	// a server configured value(1000, as of 1/3/2024); - when set to 0, the
	// page length is set to a server configured value (100, as of 1/3/2024)
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listModelVersionsRequestToPb(st *ListModelVersionsRequest) (*listModelVersionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listModelVersionsRequestPb{}
	pb.FullName = st.FullName

	pb.IncludeBrowse = st.IncludeBrowse

	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListModelVersionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listModelVersionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listModelVersionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListModelVersionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listModelVersionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listModelVersionsRequestPb struct {
	// The full three-level name of the registered model under which to list
	// model versions
	FullName string `json:"-" url:"-"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Maximum number of model versions to return. If not set, the page length
	// is set to a server configured value (100, as of 1/3/2024). - when set to
	// a value greater than 0, the page length is the minimum of this value and
	// a server configured value(1000, as of 1/3/2024); - when set to 0, the
	// page length is set to a server configured value (100, as of 1/3/2024)
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listModelVersionsRequestFromPb(pb *listModelVersionsRequestPb) (*ListModelVersionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListModelVersionsRequest{}
	st.FullName = pb.FullName
	st.IncludeBrowse = pb.IncludeBrowse
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listModelVersionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listModelVersionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListModelVersionsResponse struct {

	// Wire name: 'model_versions'
	ModelVersions []ModelVersionInfo
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listModelVersionsResponseToPb(st *ListModelVersionsResponse) (*listModelVersionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listModelVersionsResponsePb{}

	var modelVersionsPb []modelVersionInfoPb
	for _, item := range st.ModelVersions {
		itemPb, err := modelVersionInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			modelVersionsPb = append(modelVersionsPb, *itemPb)
		}
	}
	pb.ModelVersions = modelVersionsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListModelVersionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listModelVersionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listModelVersionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListModelVersionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listModelVersionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listModelVersionsResponsePb struct {
	ModelVersions []modelVersionInfoPb `json:"model_versions,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listModelVersionsResponseFromPb(pb *listModelVersionsResponsePb) (*ListModelVersionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListModelVersionsResponse{}

	var modelVersionsField []ModelVersionInfo
	for _, itemPb := range pb.ModelVersions {
		item, err := modelVersionInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			modelVersionsField = append(modelVersionsField, *item)
		}
	}
	st.ModelVersions = modelVersionsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listModelVersionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listModelVersionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List all resource quotas under a metastore.
type ListQuotasRequest struct {
	// The number of quotas to return.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque token for the next page of results.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listQuotasRequestToPb(st *ListQuotasRequest) (*listQuotasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQuotasRequestPb{}
	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListQuotasRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listQuotasRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listQuotasRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListQuotasRequest) MarshalJSON() ([]byte, error) {
	pb, err := listQuotasRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listQuotasRequestPb struct {
	// The number of quotas to return.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque token for the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listQuotasRequestFromPb(pb *listQuotasRequestPb) (*ListQuotasRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQuotasRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listQuotasRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQuotasRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListQuotasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request.
	// Wire name: 'next_page_token'
	NextPageToken string
	// An array of returned QuotaInfos.
	// Wire name: 'quotas'
	Quotas []QuotaInfo

	ForceSendFields []string `tf:"-"`
}

func listQuotasResponseToPb(st *ListQuotasResponse) (*listQuotasResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQuotasResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var quotasPb []quotaInfoPb
	for _, item := range st.Quotas {
		itemPb, err := quotaInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			quotasPb = append(quotasPb, *itemPb)
		}
	}
	pb.Quotas = quotasPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListQuotasResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listQuotasResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listQuotasResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListQuotasResponse) MarshalJSON() ([]byte, error) {
	pb, err := listQuotasResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listQuotasResponsePb struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request.
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of returned QuotaInfos.
	Quotas []quotaInfoPb `json:"quotas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listQuotasResponseFromPb(pb *listQuotasResponsePb) (*ListQuotasResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQuotasResponse{}
	st.NextPageToken = pb.NextPageToken

	var quotasField []QuotaInfo
	for _, itemPb := range pb.Quotas {
		item, err := quotaInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			quotasField = append(quotasField, *item)
		}
	}
	st.Quotas = quotasField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listQuotasResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQuotasResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List refreshes
type ListRefreshesRequest struct {
	// Full name of the table.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func listRefreshesRequestToPb(st *ListRefreshesRequest) (*listRefreshesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRefreshesRequestPb{}
	pb.TableName = st.TableName

	return pb, nil
}

func (st *ListRefreshesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listRefreshesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listRefreshesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListRefreshesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listRefreshesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listRefreshesRequestPb struct {
	// Full name of the table.
	TableName string `json:"-" url:"-"`
}

func listRefreshesRequestFromPb(pb *listRefreshesRequestPb) (*ListRefreshesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRefreshesRequest{}
	st.TableName = pb.TableName

	return st, nil
}

// List Registered Models
type ListRegisteredModelsRequest struct {
	// The identifier of the catalog under which to list registered models. If
	// specified, schema_name must be specified.
	// Wire name: 'catalog_name'
	CatalogName string `tf:"-"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// Max number of registered models to return.
	//
	// If both catalog and schema are specified: - when max_results is not
	// specified, the page length is set to a server configured value (10000, as
	// of 4/2/2024). - when set to a value greater than 0, the page length is
	// the minimum of this value and a server configured value (10000, as of
	// 4/2/2024); - when set to 0, the page length is set to a server configured
	// value (10000, as of 4/2/2024); - when set to a value less than 0, an
	// invalid parameter error is returned;
	//
	// If neither schema nor catalog is specified: - when max_results is not
	// specified, the page length is set to a server configured value (100, as
	// of 4/2/2024). - when set to a value greater than 0, the page length is
	// the minimum of this value and a server configured value (1000, as of
	// 4/2/2024); - when set to 0, the page length is set to a server configured
	// value (100, as of 4/2/2024); - when set to a value less than 0, an
	// invalid parameter error is returned;
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque token to send for the next page of results (pagination).
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// The identifier of the schema under which to list registered models. If
	// specified, catalog_name must be specified.
	// Wire name: 'schema_name'
	SchemaName string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listRegisteredModelsRequestToPb(st *ListRegisteredModelsRequest) (*listRegisteredModelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRegisteredModelsRequestPb{}
	pb.CatalogName = st.CatalogName

	pb.IncludeBrowse = st.IncludeBrowse

	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListRegisteredModelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listRegisteredModelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listRegisteredModelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListRegisteredModelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listRegisteredModelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listRegisteredModelsRequestPb struct {
	// The identifier of the catalog under which to list registered models. If
	// specified, schema_name must be specified.
	CatalogName string `json:"-" url:"catalog_name,omitempty"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Max number of registered models to return.
	//
	// If both catalog and schema are specified: - when max_results is not
	// specified, the page length is set to a server configured value (10000, as
	// of 4/2/2024). - when set to a value greater than 0, the page length is
	// the minimum of this value and a server configured value (10000, as of
	// 4/2/2024); - when set to 0, the page length is set to a server configured
	// value (10000, as of 4/2/2024); - when set to a value less than 0, an
	// invalid parameter error is returned;
	//
	// If neither schema nor catalog is specified: - when max_results is not
	// specified, the page length is set to a server configured value (100, as
	// of 4/2/2024). - when set to a value greater than 0, the page length is
	// the minimum of this value and a server configured value (1000, as of
	// 4/2/2024); - when set to 0, the page length is set to a server configured
	// value (100, as of 4/2/2024); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque token to send for the next page of results (pagination).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The identifier of the schema under which to list registered models. If
	// specified, catalog_name must be specified.
	SchemaName string `json:"-" url:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listRegisteredModelsRequestFromPb(pb *listRegisteredModelsRequestPb) (*ListRegisteredModelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRegisteredModelsRequest{}
	st.CatalogName = pb.CatalogName
	st.IncludeBrowse = pb.IncludeBrowse
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.SchemaName = pb.SchemaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listRegisteredModelsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listRegisteredModelsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListRegisteredModelsResponse struct {
	// Opaque token for pagination. Omitted if there are no more results.
	// page_token should be set to this value for fetching the next page.
	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'registered_models'
	RegisteredModels []RegisteredModelInfo

	ForceSendFields []string `tf:"-"`
}

func listRegisteredModelsResponseToPb(st *ListRegisteredModelsResponse) (*listRegisteredModelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRegisteredModelsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var registeredModelsPb []registeredModelInfoPb
	for _, item := range st.RegisteredModels {
		itemPb, err := registeredModelInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			registeredModelsPb = append(registeredModelsPb, *itemPb)
		}
	}
	pb.RegisteredModels = registeredModelsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListRegisteredModelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listRegisteredModelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listRegisteredModelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListRegisteredModelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listRegisteredModelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listRegisteredModelsResponsePb struct {
	// Opaque token for pagination. Omitted if there are no more results.
	// page_token should be set to this value for fetching the next page.
	NextPageToken string `json:"next_page_token,omitempty"`

	RegisteredModels []registeredModelInfoPb `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listRegisteredModelsResponseFromPb(pb *listRegisteredModelsResponsePb) (*ListRegisteredModelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRegisteredModelsResponse{}
	st.NextPageToken = pb.NextPageToken

	var registeredModelsField []RegisteredModelInfo
	for _, itemPb := range pb.RegisteredModels {
		item, err := registeredModelInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			registeredModelsField = append(registeredModelsField, *item)
		}
	}
	st.RegisteredModels = registeredModelsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listRegisteredModelsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listRegisteredModelsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List schemas
type ListSchemasRequest struct {
	// Parent catalog for schemas of interest.
	// Wire name: 'catalog_name'
	CatalogName string `tf:"-"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// Maximum number of schemas to return. If not set, all the schemas are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listSchemasRequestToPb(st *ListSchemasRequest) (*listSchemasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSchemasRequestPb{}
	pb.CatalogName = st.CatalogName

	pb.IncludeBrowse = st.IncludeBrowse

	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListSchemasRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listSchemasRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listSchemasRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListSchemasRequest) MarshalJSON() ([]byte, error) {
	pb, err := listSchemasRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listSchemasRequestPb struct {
	// Parent catalog for schemas of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Maximum number of schemas to return. If not set, all the schemas are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSchemasRequestFromPb(pb *listSchemasRequestPb) (*ListSchemasRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSchemasRequest{}
	st.CatalogName = pb.CatalogName
	st.IncludeBrowse = pb.IncludeBrowse
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSchemasRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSchemasRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSchemasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string
	// An array of schema information objects.
	// Wire name: 'schemas'
	Schemas []SchemaInfo

	ForceSendFields []string `tf:"-"`
}

func listSchemasResponseToPb(st *ListSchemasResponse) (*listSchemasResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSchemasResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var schemasPb []schemaInfoPb
	for _, item := range st.Schemas {
		itemPb, err := schemaInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			schemasPb = append(schemasPb, *itemPb)
		}
	}
	pb.Schemas = schemasPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListSchemasResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listSchemasResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listSchemasResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListSchemasResponse) MarshalJSON() ([]byte, error) {
	pb, err := listSchemasResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listSchemasResponsePb struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of schema information objects.
	Schemas []schemaInfoPb `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSchemasResponseFromPb(pb *listSchemasResponsePb) (*ListSchemasResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSchemasResponse{}
	st.NextPageToken = pb.NextPageToken

	var schemasField []SchemaInfo
	for _, itemPb := range pb.Schemas {
		item, err := schemaInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			schemasField = append(schemasField, *item)
		}
	}
	st.Schemas = schemasField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSchemasResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSchemasResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List credentials
type ListStorageCredentialsRequest struct {
	// Maximum number of storage credentials to return. If not set, all the
	// storage credentials are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listStorageCredentialsRequestToPb(st *ListStorageCredentialsRequest) (*listStorageCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listStorageCredentialsRequestPb{}
	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListStorageCredentialsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listStorageCredentialsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listStorageCredentialsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListStorageCredentialsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listStorageCredentialsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listStorageCredentialsRequestPb struct {
	// Maximum number of storage credentials to return. If not set, all the
	// storage credentials are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listStorageCredentialsRequestFromPb(pb *listStorageCredentialsRequestPb) (*ListStorageCredentialsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListStorageCredentialsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listStorageCredentialsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listStorageCredentialsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListStorageCredentialsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'storage_credentials'
	StorageCredentials []StorageCredentialInfo

	ForceSendFields []string `tf:"-"`
}

func listStorageCredentialsResponseToPb(st *ListStorageCredentialsResponse) (*listStorageCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listStorageCredentialsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var storageCredentialsPb []storageCredentialInfoPb
	for _, item := range st.StorageCredentials {
		itemPb, err := storageCredentialInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			storageCredentialsPb = append(storageCredentialsPb, *itemPb)
		}
	}
	pb.StorageCredentials = storageCredentialsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListStorageCredentialsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listStorageCredentialsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listStorageCredentialsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListStorageCredentialsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listStorageCredentialsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listStorageCredentialsResponsePb struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	StorageCredentials []storageCredentialInfoPb `json:"storage_credentials,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listStorageCredentialsResponseFromPb(pb *listStorageCredentialsResponsePb) (*ListStorageCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListStorageCredentialsResponse{}
	st.NextPageToken = pb.NextPageToken

	var storageCredentialsField []StorageCredentialInfo
	for _, itemPb := range pb.StorageCredentials {
		item, err := storageCredentialInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			storageCredentialsField = append(storageCredentialsField, *item)
		}
	}
	st.StorageCredentials = storageCredentialsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listStorageCredentialsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listStorageCredentialsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List table summaries
type ListSummariesRequest struct {
	// Name of parent catalog for tables of interest.
	// Wire name: 'catalog_name'
	CatalogName string `tf:"-"`
	// Whether to include a manifest containing capabilities the table has.
	// Wire name: 'include_manifest_capabilities'
	IncludeManifestCapabilities bool `tf:"-"`
	// Maximum number of summaries for tables to return. If not set, the page
	// length is set to a server configured value (10000, as of 1/5/2024). -
	// when set to a value greater than 0, the page length is the minimum of
	// this value and a server configured value (10000, as of 1/5/2024); - when
	// set to 0, the page length is set to a server configured value (10000, as
	// of 1/5/2024) (recommended); - when set to a value less than 0, an invalid
	// parameter error is returned;
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// A sql LIKE pattern (% and _) for schema names. All schemas will be
	// returned if not set or empty.
	// Wire name: 'schema_name_pattern'
	SchemaNamePattern string `tf:"-"`
	// A sql LIKE pattern (% and _) for table names. All tables will be returned
	// if not set or empty.
	// Wire name: 'table_name_pattern'
	TableNamePattern string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listSummariesRequestToPb(st *ListSummariesRequest) (*listSummariesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSummariesRequestPb{}
	pb.CatalogName = st.CatalogName

	pb.IncludeManifestCapabilities = st.IncludeManifestCapabilities

	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.SchemaNamePattern = st.SchemaNamePattern

	pb.TableNamePattern = st.TableNamePattern

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListSummariesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listSummariesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listSummariesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListSummariesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listSummariesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listSummariesRequestPb struct {
	// Name of parent catalog for tables of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether to include a manifest containing capabilities the table has.
	IncludeManifestCapabilities bool `json:"-" url:"include_manifest_capabilities,omitempty"`
	// Maximum number of summaries for tables to return. If not set, the page
	// length is set to a server configured value (10000, as of 1/5/2024). -
	// when set to a value greater than 0, the page length is the minimum of
	// this value and a server configured value (10000, as of 1/5/2024); - when
	// set to 0, the page length is set to a server configured value (10000, as
	// of 1/5/2024) (recommended); - when set to a value less than 0, an invalid
	// parameter error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// A sql LIKE pattern (% and _) for schema names. All schemas will be
	// returned if not set or empty.
	SchemaNamePattern string `json:"-" url:"schema_name_pattern,omitempty"`
	// A sql LIKE pattern (% and _) for table names. All tables will be returned
	// if not set or empty.
	TableNamePattern string `json:"-" url:"table_name_pattern,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSummariesRequestFromPb(pb *listSummariesRequestPb) (*ListSummariesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSummariesRequest{}
	st.CatalogName = pb.CatalogName
	st.IncludeManifestCapabilities = pb.IncludeManifestCapabilities
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.SchemaNamePattern = pb.SchemaNamePattern
	st.TableNamePattern = pb.TableNamePattern

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSummariesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSummariesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List system schemas
type ListSystemSchemasRequest struct {
	// Maximum number of schemas to return. - When set to 0, the page length is
	// set to a server configured value (recommended); - When set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - When set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all the schemas are returned (not
	// recommended).
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// The ID for the metastore in which the system schema resides.
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listSystemSchemasRequestToPb(st *ListSystemSchemasRequest) (*listSystemSchemasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSystemSchemasRequestPb{}
	pb.MaxResults = st.MaxResults

	pb.MetastoreId = st.MetastoreId

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListSystemSchemasRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listSystemSchemasRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listSystemSchemasRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListSystemSchemasRequest) MarshalJSON() ([]byte, error) {
	pb, err := listSystemSchemasRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listSystemSchemasRequestPb struct {
	// Maximum number of schemas to return. - When set to 0, the page length is
	// set to a server configured value (recommended); - When set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - When set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all the schemas are returned (not
	// recommended).
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// The ID for the metastore in which the system schema resides.
	MetastoreId string `json:"-" url:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSystemSchemasRequestFromPb(pb *listSystemSchemasRequestPb) (*ListSystemSchemasRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSystemSchemasRequest{}
	st.MaxResults = pb.MaxResults
	st.MetastoreId = pb.MetastoreId
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSystemSchemasRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSystemSchemasRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSystemSchemasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string
	// An array of system schema information objects.
	// Wire name: 'schemas'
	Schemas []SystemSchemaInfo

	ForceSendFields []string `tf:"-"`
}

func listSystemSchemasResponseToPb(st *ListSystemSchemasResponse) (*listSystemSchemasResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSystemSchemasResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var schemasPb []systemSchemaInfoPb
	for _, item := range st.Schemas {
		itemPb, err := systemSchemaInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			schemasPb = append(schemasPb, *itemPb)
		}
	}
	pb.Schemas = schemasPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListSystemSchemasResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listSystemSchemasResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listSystemSchemasResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListSystemSchemasResponse) MarshalJSON() ([]byte, error) {
	pb, err := listSystemSchemasResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listSystemSchemasResponsePb struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of system schema information objects.
	Schemas []systemSchemaInfoPb `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSystemSchemasResponseFromPb(pb *listSystemSchemasResponsePb) (*ListSystemSchemasResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSystemSchemasResponse{}
	st.NextPageToken = pb.NextPageToken

	var schemasField []SystemSchemaInfo
	for _, itemPb := range pb.Schemas {
		item, err := systemSchemaInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			schemasField = append(schemasField, *item)
		}
	}
	st.Schemas = schemasField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSystemSchemasResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSystemSchemasResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListTableSummariesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string
	// List of table summaries.
	// Wire name: 'tables'
	Tables []TableSummary

	ForceSendFields []string `tf:"-"`
}

func listTableSummariesResponseToPb(st *ListTableSummariesResponse) (*listTableSummariesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTableSummariesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var tablesPb []tableSummaryPb
	for _, item := range st.Tables {
		itemPb, err := tableSummaryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tablesPb = append(tablesPb, *itemPb)
		}
	}
	pb.Tables = tablesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListTableSummariesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listTableSummariesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listTableSummariesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListTableSummariesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listTableSummariesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listTableSummariesResponsePb struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of table summaries.
	Tables []tableSummaryPb `json:"tables,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listTableSummariesResponseFromPb(pb *listTableSummariesResponsePb) (*ListTableSummariesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTableSummariesResponse{}
	st.NextPageToken = pb.NextPageToken

	var tablesField []TableSummary
	for _, itemPb := range pb.Tables {
		item, err := tableSummaryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tablesField = append(tablesField, *item)
		}
	}
	st.Tables = tablesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listTableSummariesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listTableSummariesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List tables
type ListTablesRequest struct {
	// Name of parent catalog for tables of interest.
	// Wire name: 'catalog_name'
	CatalogName string `tf:"-"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// Whether delta metadata should be included in the response.
	// Wire name: 'include_delta_metadata'
	IncludeDeltaMetadata bool `tf:"-"`
	// Whether to include a manifest containing capabilities the table has.
	// Wire name: 'include_manifest_capabilities'
	IncludeManifestCapabilities bool `tf:"-"`
	// Maximum number of tables to return. If not set, all the tables are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Whether to omit the columns of the table from the response or not.
	// Wire name: 'omit_columns'
	OmitColumns bool `tf:"-"`
	// Whether to omit the properties of the table from the response or not.
	// Wire name: 'omit_properties'
	OmitProperties bool `tf:"-"`
	// Whether to omit the username of the table (e.g. owner, updated_by,
	// created_by) from the response or not.
	// Wire name: 'omit_username'
	OmitUsername bool `tf:"-"`
	// Opaque token to send for the next page of results (pagination).
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// Parent schema of tables.
	// Wire name: 'schema_name'
	SchemaName string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listTablesRequestToPb(st *ListTablesRequest) (*listTablesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTablesRequestPb{}
	pb.CatalogName = st.CatalogName

	pb.IncludeBrowse = st.IncludeBrowse

	pb.IncludeDeltaMetadata = st.IncludeDeltaMetadata

	pb.IncludeManifestCapabilities = st.IncludeManifestCapabilities

	pb.MaxResults = st.MaxResults

	pb.OmitColumns = st.OmitColumns

	pb.OmitProperties = st.OmitProperties

	pb.OmitUsername = st.OmitUsername

	pb.PageToken = st.PageToken

	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListTablesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listTablesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listTablesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListTablesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listTablesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listTablesRequestPb struct {
	// Name of parent catalog for tables of interest.
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata bool `json:"-" url:"include_delta_metadata,omitempty"`
	// Whether to include a manifest containing capabilities the table has.
	IncludeManifestCapabilities bool `json:"-" url:"include_manifest_capabilities,omitempty"`
	// Maximum number of tables to return. If not set, all the tables are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Whether to omit the columns of the table from the response or not.
	OmitColumns bool `json:"-" url:"omit_columns,omitempty"`
	// Whether to omit the properties of the table from the response or not.
	OmitProperties bool `json:"-" url:"omit_properties,omitempty"`
	// Whether to omit the username of the table (e.g. owner, updated_by,
	// created_by) from the response or not.
	OmitUsername bool `json:"-" url:"omit_username,omitempty"`
	// Opaque token to send for the next page of results (pagination).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent schema of tables.
	SchemaName string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listTablesRequestFromPb(pb *listTablesRequestPb) (*ListTablesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTablesRequest{}
	st.CatalogName = pb.CatalogName
	st.IncludeBrowse = pb.IncludeBrowse
	st.IncludeDeltaMetadata = pb.IncludeDeltaMetadata
	st.IncludeManifestCapabilities = pb.IncludeManifestCapabilities
	st.MaxResults = pb.MaxResults
	st.OmitColumns = pb.OmitColumns
	st.OmitProperties = pb.OmitProperties
	st.OmitUsername = pb.OmitUsername
	st.PageToken = pb.PageToken
	st.SchemaName = pb.SchemaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listTablesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listTablesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListTablesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string
	// An array of table information objects.
	// Wire name: 'tables'
	Tables []TableInfo

	ForceSendFields []string `tf:"-"`
}

func listTablesResponseToPb(st *ListTablesResponse) (*listTablesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTablesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var tablesPb []tableInfoPb
	for _, item := range st.Tables {
		itemPb, err := tableInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tablesPb = append(tablesPb, *itemPb)
		}
	}
	pb.Tables = tablesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListTablesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listTablesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listTablesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListTablesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listTablesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listTablesResponsePb struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of table information objects.
	Tables []tableInfoPb `json:"tables,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listTablesResponseFromPb(pb *listTablesResponsePb) (*ListTablesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTablesResponse{}
	st.NextPageToken = pb.NextPageToken

	var tablesField []TableInfo
	for _, itemPb := range pb.Tables {
		item, err := tableInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tablesField = append(tablesField, *item)
		}
	}
	st.Tables = tablesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listTablesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listTablesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List Volumes
type ListVolumesRequest struct {
	// The identifier of the catalog
	// Wire name: 'catalog_name'
	CatalogName string `tf:"-"`
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// Maximum number of volumes to return (page length).
	//
	// If not set, the page length is set to a server configured value (10000,
	// as of 1/29/2024). - when set to a value greater than 0, the page length
	// is the minimum of this value and a server configured value (10000, as of
	// 1/29/2024); - when set to 0, the page length is set to a server
	// configured value (10000, as of 1/29/2024) (recommended); - when set to a
	// value less than 0, an invalid parameter error is returned;
	//
	// Note: this parameter controls only the maximum number of volumes to
	// return. The actual number of volumes returned in a page may be smaller
	// than this value, including 0, even if there are more pages.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque token returned by a previous request. It must be included in the
	// request to retrieve the next page of results (pagination).
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// The identifier of the schema
	// Wire name: 'schema_name'
	SchemaName string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listVolumesRequestToPb(st *ListVolumesRequest) (*listVolumesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listVolumesRequestPb{}
	pb.CatalogName = st.CatalogName

	pb.IncludeBrowse = st.IncludeBrowse

	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListVolumesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listVolumesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listVolumesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListVolumesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listVolumesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listVolumesRequestPb struct {
	// The identifier of the catalog
	CatalogName string `json:"-" url:"catalog_name"`
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// Maximum number of volumes to return (page length).
	//
	// If not set, the page length is set to a server configured value (10000,
	// as of 1/29/2024). - when set to a value greater than 0, the page length
	// is the minimum of this value and a server configured value (10000, as of
	// 1/29/2024); - when set to 0, the page length is set to a server
	// configured value (10000, as of 1/29/2024) (recommended); - when set to a
	// value less than 0, an invalid parameter error is returned;
	//
	// Note: this parameter controls only the maximum number of volumes to
	// return. The actual number of volumes returned in a page may be smaller
	// than this value, including 0, even if there are more pages.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque token returned by a previous request. It must be included in the
	// request to retrieve the next page of results (pagination).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The identifier of the schema
	SchemaName string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listVolumesRequestFromPb(pb *listVolumesRequestPb) (*ListVolumesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListVolumesRequest{}
	st.CatalogName = pb.CatalogName
	st.IncludeBrowse = pb.IncludeBrowse
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.SchemaName = pb.SchemaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listVolumesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listVolumesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListVolumesResponseContent struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request to retrieve the next page of results.
	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'volumes'
	Volumes []VolumeInfo

	ForceSendFields []string `tf:"-"`
}

func listVolumesResponseContentToPb(st *ListVolumesResponseContent) (*listVolumesResponseContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listVolumesResponseContentPb{}
	pb.NextPageToken = st.NextPageToken

	var volumesPb []volumeInfoPb
	for _, item := range st.Volumes {
		itemPb, err := volumeInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			volumesPb = append(volumesPb, *itemPb)
		}
	}
	pb.Volumes = volumesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListVolumesResponseContent) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listVolumesResponseContentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listVolumesResponseContentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListVolumesResponseContent) MarshalJSON() ([]byte, error) {
	pb, err := listVolumesResponseContentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listVolumesResponseContentPb struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request to retrieve the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	Volumes []volumeInfoPb `json:"volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listVolumesResponseContentFromPb(pb *listVolumesResponseContentPb) (*ListVolumesResponseContent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListVolumesResponseContent{}
	st.NextPageToken = pb.NextPageToken

	var volumesField []VolumeInfo
	for _, itemPb := range pb.Volumes {
		item, err := volumeInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			volumesField = append(volumesField, *item)
		}
	}
	st.Volumes = volumesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listVolumesResponseContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listVolumesResponseContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The artifact pattern matching type
type MatchType string
type matchTypePb string

const MatchTypePrefixMatch MatchType = `PREFIX_MATCH`

// String representation for [fmt.Print]
func (f *MatchType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MatchType) Set(v string) error {
	switch v {
	case `PREFIX_MATCH`:
		*f = MatchType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PREFIX_MATCH"`, v)
	}
}

// Type always returns MatchType to satisfy [pflag.Value] interface
func (f *MatchType) Type() string {
	return "MatchType"
}

func matchTypeToPb(st *MatchType) (*matchTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := matchTypePb(*st)
	return &pb, nil
}

func matchTypeFromPb(pb *matchTypePb) (*MatchType, error) {
	if pb == nil {
		return nil, nil
	}
	st := MatchType(*pb)
	return &st, nil
}

type MetastoreAssignment struct {
	// The name of the default catalog in the metastore.
	// Wire name: 'default_catalog_name'
	DefaultCatalogName string
	// The unique ID of the metastore.
	// Wire name: 'metastore_id'
	MetastoreId string
	// The unique ID of the Databricks workspace.
	// Wire name: 'workspace_id'
	WorkspaceId int64

	ForceSendFields []string `tf:"-"`
}

func metastoreAssignmentToPb(st *MetastoreAssignment) (*metastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &metastoreAssignmentPb{}
	pb.DefaultCatalogName = st.DefaultCatalogName

	pb.MetastoreId = st.MetastoreId

	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *MetastoreAssignment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &metastoreAssignmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := metastoreAssignmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MetastoreAssignment) MarshalJSON() ([]byte, error) {
	pb, err := metastoreAssignmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type metastoreAssignmentPb struct {
	// The name of the default catalog in the metastore.
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	// The unique ID of the metastore.
	MetastoreId string `json:"metastore_id"`
	// The unique ID of the Databricks workspace.
	WorkspaceId int64 `json:"workspace_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func metastoreAssignmentFromPb(pb *metastoreAssignmentPb) (*MetastoreAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MetastoreAssignment{}
	st.DefaultCatalogName = pb.DefaultCatalogName
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *metastoreAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st metastoreAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MetastoreInfo struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	// Wire name: 'cloud'
	Cloud string
	// Time at which this metastore was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of metastore creator.
	// Wire name: 'created_by'
	CreatedBy string
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	// Wire name: 'default_data_access_config_id'
	DefaultDataAccessConfigId string
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	// Wire name: 'delta_sharing_organization_name'
	DeltaSharingOrganizationName string
	// The lifetime of delta sharing recipient token in seconds.
	// Wire name: 'delta_sharing_recipient_token_lifetime_in_seconds'
	DeltaSharingRecipientTokenLifetimeInSeconds int64
	// The scope of Delta Sharing enabled for the metastore.
	// Wire name: 'delta_sharing_scope'
	DeltaSharingScope MetastoreInfoDeltaSharingScope
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	// Wire name: 'external_access_enabled'
	ExternalAccessEnabled bool
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	// Wire name: 'global_metastore_id'
	GlobalMetastoreId string
	// Unique identifier of metastore.
	// Wire name: 'metastore_id'
	MetastoreId string
	// The user-specified name of the metastore.
	// Wire name: 'name'
	Name string
	// The owner of the metastore.
	// Wire name: 'owner'
	Owner string
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	// Wire name: 'privilege_model_version'
	PrivilegeModelVersion string
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	// Wire name: 'region'
	Region string
	// The storage root URL for metastore
	// Wire name: 'storage_root'
	StorageRoot string
	// UUID of storage credential to access the metastore storage_root.
	// Wire name: 'storage_root_credential_id'
	StorageRootCredentialId string
	// Name of the storage credential to access the metastore storage_root.
	// Wire name: 'storage_root_credential_name'
	StorageRootCredentialName string
	// Time at which the metastore was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// Username of user who last modified the metastore.
	// Wire name: 'updated_by'
	UpdatedBy string

	ForceSendFields []string `tf:"-"`
}

func metastoreInfoToPb(st *MetastoreInfo) (*metastoreInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &metastoreInfoPb{}
	pb.Cloud = st.Cloud

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	pb.DefaultDataAccessConfigId = st.DefaultDataAccessConfigId

	pb.DeltaSharingOrganizationName = st.DeltaSharingOrganizationName

	pb.DeltaSharingRecipientTokenLifetimeInSeconds = st.DeltaSharingRecipientTokenLifetimeInSeconds

	pb.DeltaSharingScope = st.DeltaSharingScope

	pb.ExternalAccessEnabled = st.ExternalAccessEnabled

	pb.GlobalMetastoreId = st.GlobalMetastoreId

	pb.MetastoreId = st.MetastoreId

	pb.Name = st.Name

	pb.Owner = st.Owner

	pb.PrivilegeModelVersion = st.PrivilegeModelVersion

	pb.Region = st.Region

	pb.StorageRoot = st.StorageRoot

	pb.StorageRootCredentialId = st.StorageRootCredentialId

	pb.StorageRootCredentialName = st.StorageRootCredentialName

	pb.UpdatedAt = st.UpdatedAt

	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *MetastoreInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &metastoreInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := metastoreInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MetastoreInfo) MarshalJSON() ([]byte, error) {
	pb, err := metastoreInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type metastoreInfoPb struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud string `json:"cloud,omitempty"`
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of metastore creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope MetastoreInfoDeltaSharingScope `json:"delta_sharing_scope,omitempty"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled bool `json:"external_access_enabled,omitempty"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId string `json:"global_metastore_id,omitempty"`
	// Unique identifier of metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The user-specified name of the metastore.
	Name string `json:"name,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	StorageRoot string `json:"storage_root,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName string `json:"storage_root_credential_name,omitempty"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the metastore.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func metastoreInfoFromPb(pb *metastoreInfoPb) (*MetastoreInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MetastoreInfo{}
	st.Cloud = pb.Cloud
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.DefaultDataAccessConfigId = pb.DefaultDataAccessConfigId
	st.DeltaSharingOrganizationName = pb.DeltaSharingOrganizationName
	st.DeltaSharingRecipientTokenLifetimeInSeconds = pb.DeltaSharingRecipientTokenLifetimeInSeconds
	st.DeltaSharingScope = pb.DeltaSharingScope
	st.ExternalAccessEnabled = pb.ExternalAccessEnabled
	st.GlobalMetastoreId = pb.GlobalMetastoreId
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.PrivilegeModelVersion = pb.PrivilegeModelVersion
	st.Region = pb.Region
	st.StorageRoot = pb.StorageRoot
	st.StorageRootCredentialId = pb.StorageRootCredentialId
	st.StorageRootCredentialName = pb.StorageRootCredentialName
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *metastoreInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st metastoreInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The scope of Delta Sharing enabled for the metastore.
type MetastoreInfoDeltaSharingScope string
type metastoreInfoDeltaSharingScopePb string

const MetastoreInfoDeltaSharingScopeInternal MetastoreInfoDeltaSharingScope = `INTERNAL`

const MetastoreInfoDeltaSharingScopeInternalAndExternal MetastoreInfoDeltaSharingScope = `INTERNAL_AND_EXTERNAL`

// String representation for [fmt.Print]
func (f *MetastoreInfoDeltaSharingScope) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MetastoreInfoDeltaSharingScope) Set(v string) error {
	switch v {
	case `INTERNAL`, `INTERNAL_AND_EXTERNAL`:
		*f = MetastoreInfoDeltaSharingScope(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INTERNAL", "INTERNAL_AND_EXTERNAL"`, v)
	}
}

// Type always returns MetastoreInfoDeltaSharingScope to satisfy [pflag.Value] interface
func (f *MetastoreInfoDeltaSharingScope) Type() string {
	return "MetastoreInfoDeltaSharingScope"
}

func metastoreInfoDeltaSharingScopeToPb(st *MetastoreInfoDeltaSharingScope) (*metastoreInfoDeltaSharingScopePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := metastoreInfoDeltaSharingScopePb(*st)
	return &pb, nil
}

func metastoreInfoDeltaSharingScopeFromPb(pb *metastoreInfoDeltaSharingScopePb) (*MetastoreInfoDeltaSharingScope, error) {
	if pb == nil {
		return nil, nil
	}
	st := MetastoreInfoDeltaSharingScope(*pb)
	return &st, nil
}

type ModelVersionInfo struct {
	// List of aliases associated with the model version
	// Wire name: 'aliases'
	Aliases []RegisteredModelAlias
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool
	// The name of the catalog containing the model version
	// Wire name: 'catalog_name'
	CatalogName string
	// The comment attached to the model version
	// Wire name: 'comment'
	Comment string

	// Wire name: 'created_at'
	CreatedAt int64
	// The identifier of the user who created the model version
	// Wire name: 'created_by'
	CreatedBy string
	// The unique identifier of the model version
	// Wire name: 'id'
	Id string
	// The unique identifier of the metastore containing the model version
	// Wire name: 'metastore_id'
	MetastoreId string
	// The name of the parent registered model of the model version, relative to
	// parent schema
	// Wire name: 'model_name'
	ModelName string
	// Model version dependencies, for feature-store packaged models
	// Wire name: 'model_version_dependencies'
	ModelVersionDependencies *DependencyList
	// MLflow run ID used when creating the model version, if ``source`` was
	// generated by an experiment run stored in an MLflow tracking server
	// Wire name: 'run_id'
	RunId string
	// ID of the Databricks workspace containing the MLflow run that generated
	// this model version, if applicable
	// Wire name: 'run_workspace_id'
	RunWorkspaceId int
	// The name of the schema containing the model version, relative to parent
	// catalog
	// Wire name: 'schema_name'
	SchemaName string
	// URI indicating the location of the source artifacts (files) for the model
	// version
	// Wire name: 'source'
	Source string
	// Current status of the model version. Newly created model versions start
	// in PENDING_REGISTRATION status, then move to READY status once the model
	// version files are uploaded and the model version is finalized. Only model
	// versions in READY status can be loaded for inference or served.
	// Wire name: 'status'
	Status ModelVersionInfoStatus
	// The storage location on the cloud under which model version data files
	// are stored
	// Wire name: 'storage_location'
	StorageLocation string

	// Wire name: 'updated_at'
	UpdatedAt int64
	// The identifier of the user who updated the model version last time
	// Wire name: 'updated_by'
	UpdatedBy string
	// Integer model version number, used to reference the model version in API
	// requests.
	// Wire name: 'version'
	Version int

	ForceSendFields []string `tf:"-"`
}

func modelVersionInfoToPb(st *ModelVersionInfo) (*modelVersionInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelVersionInfoPb{}

	var aliasesPb []registeredModelAliasPb
	for _, item := range st.Aliases {
		itemPb, err := registeredModelAliasToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			aliasesPb = append(aliasesPb, *itemPb)
		}
	}
	pb.Aliases = aliasesPb

	pb.BrowseOnly = st.BrowseOnly

	pb.CatalogName = st.CatalogName

	pb.Comment = st.Comment

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	pb.Id = st.Id

	pb.MetastoreId = st.MetastoreId

	pb.ModelName = st.ModelName

	modelVersionDependenciesPb, err := dependencyListToPb(st.ModelVersionDependencies)
	if err != nil {
		return nil, err
	}
	if modelVersionDependenciesPb != nil {
		pb.ModelVersionDependencies = modelVersionDependenciesPb
	}

	pb.RunId = st.RunId

	pb.RunWorkspaceId = st.RunWorkspaceId

	pb.SchemaName = st.SchemaName

	pb.Source = st.Source

	pb.Status = st.Status

	pb.StorageLocation = st.StorageLocation

	pb.UpdatedAt = st.UpdatedAt

	pb.UpdatedBy = st.UpdatedBy

	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ModelVersionInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelVersionInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelVersionInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ModelVersionInfo) MarshalJSON() ([]byte, error) {
	pb, err := modelVersionInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type modelVersionInfoPb struct {
	// List of aliases associated with the model version
	Aliases []registeredModelAliasPb `json:"aliases,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The name of the catalog containing the model version
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the model version
	Comment string `json:"comment,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the model version
	CreatedBy string `json:"created_by,omitempty"`
	// The unique identifier of the model version
	Id string `json:"id,omitempty"`
	// The unique identifier of the metastore containing the model version
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the parent registered model of the model version, relative to
	// parent schema
	ModelName string `json:"model_name,omitempty"`
	// Model version dependencies, for feature-store packaged models
	ModelVersionDependencies *dependencyListPb `json:"model_version_dependencies,omitempty"`
	// MLflow run ID used when creating the model version, if ``source`` was
	// generated by an experiment run stored in an MLflow tracking server
	RunId string `json:"run_id,omitempty"`
	// ID of the Databricks workspace containing the MLflow run that generated
	// this model version, if applicable
	RunWorkspaceId int `json:"run_workspace_id,omitempty"`
	// The name of the schema containing the model version, relative to parent
	// catalog
	SchemaName string `json:"schema_name,omitempty"`
	// URI indicating the location of the source artifacts (files) for the model
	// version
	Source string `json:"source,omitempty"`
	// Current status of the model version. Newly created model versions start
	// in PENDING_REGISTRATION status, then move to READY status once the model
	// version files are uploaded and the model version is finalized. Only model
	// versions in READY status can be loaded for inference or served.
	Status ModelVersionInfoStatus `json:"status,omitempty"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string `json:"storage_location,omitempty"`

	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the model version last time
	UpdatedBy string `json:"updated_by,omitempty"`
	// Integer model version number, used to reference the model version in API
	// requests.
	Version int `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func modelVersionInfoFromPb(pb *modelVersionInfoPb) (*ModelVersionInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelVersionInfo{}

	var aliasesField []RegisteredModelAlias
	for _, itemPb := range pb.Aliases {
		item, err := registeredModelAliasFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			aliasesField = append(aliasesField, *item)
		}
	}
	st.Aliases = aliasesField
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.Id = pb.Id
	st.MetastoreId = pb.MetastoreId
	st.ModelName = pb.ModelName
	modelVersionDependenciesField, err := dependencyListFromPb(pb.ModelVersionDependencies)
	if err != nil {
		return nil, err
	}
	if modelVersionDependenciesField != nil {
		st.ModelVersionDependencies = modelVersionDependenciesField
	}
	st.RunId = pb.RunId
	st.RunWorkspaceId = pb.RunWorkspaceId
	st.SchemaName = pb.SchemaName
	st.Source = pb.Source
	st.Status = pb.Status
	st.StorageLocation = pb.StorageLocation
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *modelVersionInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st modelVersionInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Current status of the model version. Newly created model versions start in
// PENDING_REGISTRATION status, then move to READY status once the model version
// files are uploaded and the model version is finalized. Only model versions in
// READY status can be loaded for inference or served.
type ModelVersionInfoStatus string
type modelVersionInfoStatusPb string

const ModelVersionInfoStatusFailedRegistration ModelVersionInfoStatus = `FAILED_REGISTRATION`

const ModelVersionInfoStatusPendingRegistration ModelVersionInfoStatus = `PENDING_REGISTRATION`

const ModelVersionInfoStatusReady ModelVersionInfoStatus = `READY`

// String representation for [fmt.Print]
func (f *ModelVersionInfoStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ModelVersionInfoStatus) Set(v string) error {
	switch v {
	case `FAILED_REGISTRATION`, `PENDING_REGISTRATION`, `READY`:
		*f = ModelVersionInfoStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_REGISTRATION", "PENDING_REGISTRATION", "READY"`, v)
	}
}

// Type always returns ModelVersionInfoStatus to satisfy [pflag.Value] interface
func (f *ModelVersionInfoStatus) Type() string {
	return "ModelVersionInfoStatus"
}

func modelVersionInfoStatusToPb(st *ModelVersionInfoStatus) (*modelVersionInfoStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := modelVersionInfoStatusPb(*st)
	return &pb, nil
}

func modelVersionInfoStatusFromPb(pb *modelVersionInfoStatusPb) (*ModelVersionInfoStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := ModelVersionInfoStatus(*pb)
	return &st, nil
}

type MonitorCronSchedule struct {
	// Read only field that indicates whether a schedule is paused or not.
	// Wire name: 'pause_status'
	PauseStatus MonitorCronSchedulePauseStatus
	// The expression that determines when to run the monitor. See [examples].
	//
	// [examples]: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	// Wire name: 'quartz_cron_expression'
	QuartzCronExpression string
	// The timezone id (e.g., ``"PST"``) in which to evaluate the quartz
	// expression.
	// Wire name: 'timezone_id'
	TimezoneId string
}

func monitorCronScheduleToPb(st *MonitorCronSchedule) (*monitorCronSchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorCronSchedulePb{}
	pb.PauseStatus = st.PauseStatus

	pb.QuartzCronExpression = st.QuartzCronExpression

	pb.TimezoneId = st.TimezoneId

	return pb, nil
}

func (st *MonitorCronSchedule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &monitorCronSchedulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := monitorCronScheduleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MonitorCronSchedule) MarshalJSON() ([]byte, error) {
	pb, err := monitorCronScheduleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type monitorCronSchedulePb struct {
	// Read only field that indicates whether a schedule is paused or not.
	PauseStatus MonitorCronSchedulePauseStatus `json:"pause_status,omitempty"`
	// The expression that determines when to run the monitor. See [examples].
	//
	// [examples]: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression string `json:"quartz_cron_expression"`
	// The timezone id (e.g., ``"PST"``) in which to evaluate the quartz
	// expression.
	TimezoneId string `json:"timezone_id"`
}

func monitorCronScheduleFromPb(pb *monitorCronSchedulePb) (*MonitorCronSchedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorCronSchedule{}
	st.PauseStatus = pb.PauseStatus
	st.QuartzCronExpression = pb.QuartzCronExpression
	st.TimezoneId = pb.TimezoneId

	return st, nil
}

// Read only field that indicates whether a schedule is paused or not.
type MonitorCronSchedulePauseStatus string
type monitorCronSchedulePauseStatusPb string

const MonitorCronSchedulePauseStatusPaused MonitorCronSchedulePauseStatus = `PAUSED`

const MonitorCronSchedulePauseStatusUnpaused MonitorCronSchedulePauseStatus = `UNPAUSED`

// String representation for [fmt.Print]
func (f *MonitorCronSchedulePauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorCronSchedulePauseStatus) Set(v string) error {
	switch v {
	case `PAUSED`, `UNPAUSED`:
		*f = MonitorCronSchedulePauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PAUSED", "UNPAUSED"`, v)
	}
}

// Type always returns MonitorCronSchedulePauseStatus to satisfy [pflag.Value] interface
func (f *MonitorCronSchedulePauseStatus) Type() string {
	return "MonitorCronSchedulePauseStatus"
}

func monitorCronSchedulePauseStatusToPb(st *MonitorCronSchedulePauseStatus) (*monitorCronSchedulePauseStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := monitorCronSchedulePauseStatusPb(*st)
	return &pb, nil
}

func monitorCronSchedulePauseStatusFromPb(pb *monitorCronSchedulePauseStatusPb) (*MonitorCronSchedulePauseStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := MonitorCronSchedulePauseStatus(*pb)
	return &st, nil
}

type MonitorDataClassificationConfig struct {
	// Whether data classification is enabled.
	// Wire name: 'enabled'
	Enabled bool

	ForceSendFields []string `tf:"-"`
}

func monitorDataClassificationConfigToPb(st *MonitorDataClassificationConfig) (*monitorDataClassificationConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorDataClassificationConfigPb{}
	pb.Enabled = st.Enabled

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *MonitorDataClassificationConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &monitorDataClassificationConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := monitorDataClassificationConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MonitorDataClassificationConfig) MarshalJSON() ([]byte, error) {
	pb, err := monitorDataClassificationConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type monitorDataClassificationConfigPb struct {
	// Whether data classification is enabled.
	Enabled bool `json:"enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func monitorDataClassificationConfigFromPb(pb *monitorDataClassificationConfigPb) (*MonitorDataClassificationConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorDataClassificationConfig{}
	st.Enabled = pb.Enabled

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *monitorDataClassificationConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st monitorDataClassificationConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MonitorDestination struct {
	// The list of email addresses to send the notification to. A maximum of 5
	// email addresses is supported.
	// Wire name: 'email_addresses'
	EmailAddresses []string
}

func monitorDestinationToPb(st *MonitorDestination) (*monitorDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorDestinationPb{}
	pb.EmailAddresses = st.EmailAddresses

	return pb, nil
}

func (st *MonitorDestination) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &monitorDestinationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := monitorDestinationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MonitorDestination) MarshalJSON() ([]byte, error) {
	pb, err := monitorDestinationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type monitorDestinationPb struct {
	// The list of email addresses to send the notification to. A maximum of 5
	// email addresses is supported.
	EmailAddresses []string `json:"email_addresses,omitempty"`
}

func monitorDestinationFromPb(pb *monitorDestinationPb) (*MonitorDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorDestination{}
	st.EmailAddresses = pb.EmailAddresses

	return st, nil
}

type MonitorInferenceLog struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	// Wire name: 'granularities'
	Granularities []string
	// Optional column that contains the ground truth for the prediction.
	// Wire name: 'label_col'
	LabelCol string
	// Column that contains the id of the model generating the predictions.
	// Metrics will be computed per model id by default, and also across all
	// model ids.
	// Wire name: 'model_id_col'
	ModelIdCol string
	// Column that contains the output/prediction from the model.
	// Wire name: 'prediction_col'
	PredictionCol string
	// Optional column that contains the prediction probabilities for each class
	// in a classification problem type. The values in this column should be a
	// map, mapping each class label to the prediction probability for a given
	// sample. The map should be of PySpark MapType().
	// Wire name: 'prediction_proba_col'
	PredictionProbaCol string
	// Problem type the model aims to solve. Determines the type of
	// model-quality metrics that will be computed.
	// Wire name: 'problem_type'
	ProblemType MonitorInferenceLogProblemType
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	// Wire name: 'timestamp_col'
	TimestampCol string

	ForceSendFields []string `tf:"-"`
}

func monitorInferenceLogToPb(st *MonitorInferenceLog) (*monitorInferenceLogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorInferenceLogPb{}
	pb.Granularities = st.Granularities

	pb.LabelCol = st.LabelCol

	pb.ModelIdCol = st.ModelIdCol

	pb.PredictionCol = st.PredictionCol

	pb.PredictionProbaCol = st.PredictionProbaCol

	pb.ProblemType = st.ProblemType

	pb.TimestampCol = st.TimestampCol

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *MonitorInferenceLog) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &monitorInferenceLogPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := monitorInferenceLogFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MonitorInferenceLog) MarshalJSON() ([]byte, error) {
	pb, err := monitorInferenceLogToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type monitorInferenceLogPb struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	Granularities []string `json:"granularities"`
	// Optional column that contains the ground truth for the prediction.
	LabelCol string `json:"label_col,omitempty"`
	// Column that contains the id of the model generating the predictions.
	// Metrics will be computed per model id by default, and also across all
	// model ids.
	ModelIdCol string `json:"model_id_col"`
	// Column that contains the output/prediction from the model.
	PredictionCol string `json:"prediction_col"`
	// Optional column that contains the prediction probabilities for each class
	// in a classification problem type. The values in this column should be a
	// map, mapping each class label to the prediction probability for a given
	// sample. The map should be of PySpark MapType().
	PredictionProbaCol string `json:"prediction_proba_col,omitempty"`
	// Problem type the model aims to solve. Determines the type of
	// model-quality metrics that will be computed.
	ProblemType MonitorInferenceLogProblemType `json:"problem_type"`
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	TimestampCol string `json:"timestamp_col"`

	ForceSendFields []string `json:"-" url:"-"`
}

func monitorInferenceLogFromPb(pb *monitorInferenceLogPb) (*MonitorInferenceLog, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorInferenceLog{}
	st.Granularities = pb.Granularities
	st.LabelCol = pb.LabelCol
	st.ModelIdCol = pb.ModelIdCol
	st.PredictionCol = pb.PredictionCol
	st.PredictionProbaCol = pb.PredictionProbaCol
	st.ProblemType = pb.ProblemType
	st.TimestampCol = pb.TimestampCol

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *monitorInferenceLogPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st monitorInferenceLogPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Problem type the model aims to solve. Determines the type of model-quality
// metrics that will be computed.
type MonitorInferenceLogProblemType string
type monitorInferenceLogProblemTypePb string

const MonitorInferenceLogProblemTypeProblemTypeClassification MonitorInferenceLogProblemType = `PROBLEM_TYPE_CLASSIFICATION`

const MonitorInferenceLogProblemTypeProblemTypeRegression MonitorInferenceLogProblemType = `PROBLEM_TYPE_REGRESSION`

// String representation for [fmt.Print]
func (f *MonitorInferenceLogProblemType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorInferenceLogProblemType) Set(v string) error {
	switch v {
	case `PROBLEM_TYPE_CLASSIFICATION`, `PROBLEM_TYPE_REGRESSION`:
		*f = MonitorInferenceLogProblemType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PROBLEM_TYPE_CLASSIFICATION", "PROBLEM_TYPE_REGRESSION"`, v)
	}
}

// Type always returns MonitorInferenceLogProblemType to satisfy [pflag.Value] interface
func (f *MonitorInferenceLogProblemType) Type() string {
	return "MonitorInferenceLogProblemType"
}

func monitorInferenceLogProblemTypeToPb(st *MonitorInferenceLogProblemType) (*monitorInferenceLogProblemTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := monitorInferenceLogProblemTypePb(*st)
	return &pb, nil
}

func monitorInferenceLogProblemTypeFromPb(pb *monitorInferenceLogProblemTypePb) (*MonitorInferenceLogProblemType, error) {
	if pb == nil {
		return nil, nil
	}
	st := MonitorInferenceLogProblemType(*pb)
	return &st, nil
}

type MonitorInfo struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	// Wire name: 'assets_dir'
	AssetsDir string
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	// Wire name: 'baseline_table_name'
	BaselineTableName string
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	// Wire name: 'custom_metrics'
	CustomMetrics []MonitorMetric
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	// Wire name: 'dashboard_id'
	DashboardId string
	// The data classification config for the monitor.
	// Wire name: 'data_classification_config'
	DataClassificationConfig *MonitorDataClassificationConfig
	// The full name of the drift metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	// Wire name: 'drift_metrics_table_name'
	DriftMetricsTableName string
	// Configuration for monitoring inference logs.
	// Wire name: 'inference_log'
	InferenceLog *MonitorInferenceLog
	// The latest failure message of the monitor (if any).
	// Wire name: 'latest_monitor_failure_msg'
	LatestMonitorFailureMsg string
	// The version of the monitor config (e.g. 1,2,3). If negative, the monitor
	// may be corrupted.
	// Wire name: 'monitor_version'
	MonitorVersion string
	// The notification settings for the monitor.
	// Wire name: 'notifications'
	Notifications *MonitorNotifications
	// Schema where output metric tables are created.
	// Wire name: 'output_schema_name'
	OutputSchemaName string
	// The full name of the profile metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	// Wire name: 'profile_metrics_table_name'
	ProfileMetricsTableName string
	// The schedule for automatically updating and refreshing metric tables.
	// Wire name: 'schedule'
	Schedule *MonitorCronSchedule
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	// Wire name: 'slicing_exprs'
	SlicingExprs []string
	// Configuration for monitoring snapshot tables.
	// Wire name: 'snapshot'
	Snapshot *MonitorSnapshot
	// The status of the monitor.
	// Wire name: 'status'
	Status MonitorInfoStatus
	// The full name of the table to monitor. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	// Wire name: 'table_name'
	TableName string
	// Configuration for monitoring time series tables.
	// Wire name: 'time_series'
	TimeSeries *MonitorTimeSeries

	ForceSendFields []string `tf:"-"`
}

func monitorInfoToPb(st *MonitorInfo) (*monitorInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorInfoPb{}
	pb.AssetsDir = st.AssetsDir

	pb.BaselineTableName = st.BaselineTableName

	var customMetricsPb []monitorMetricPb
	for _, item := range st.CustomMetrics {
		itemPb, err := monitorMetricToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customMetricsPb = append(customMetricsPb, *itemPb)
		}
	}
	pb.CustomMetrics = customMetricsPb

	pb.DashboardId = st.DashboardId

	dataClassificationConfigPb, err := monitorDataClassificationConfigToPb(st.DataClassificationConfig)
	if err != nil {
		return nil, err
	}
	if dataClassificationConfigPb != nil {
		pb.DataClassificationConfig = dataClassificationConfigPb
	}

	pb.DriftMetricsTableName = st.DriftMetricsTableName

	inferenceLogPb, err := monitorInferenceLogToPb(st.InferenceLog)
	if err != nil {
		return nil, err
	}
	if inferenceLogPb != nil {
		pb.InferenceLog = inferenceLogPb
	}

	pb.LatestMonitorFailureMsg = st.LatestMonitorFailureMsg

	pb.MonitorVersion = st.MonitorVersion

	notificationsPb, err := monitorNotificationsToPb(st.Notifications)
	if err != nil {
		return nil, err
	}
	if notificationsPb != nil {
		pb.Notifications = notificationsPb
	}

	pb.OutputSchemaName = st.OutputSchemaName

	pb.ProfileMetricsTableName = st.ProfileMetricsTableName

	schedulePb, err := monitorCronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}

	pb.SlicingExprs = st.SlicingExprs

	snapshotPb, err := monitorSnapshotToPb(st.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotPb != nil {
		pb.Snapshot = snapshotPb
	}

	pb.Status = st.Status

	pb.TableName = st.TableName

	timeSeriesPb, err := monitorTimeSeriesToPb(st.TimeSeries)
	if err != nil {
		return nil, err
	}
	if timeSeriesPb != nil {
		pb.TimeSeries = timeSeriesPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *MonitorInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &monitorInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := monitorInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MonitorInfo) MarshalJSON() ([]byte, error) {
	pb, err := monitorInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type monitorInfoPb struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	AssetsDir string `json:"assets_dir,omitempty"`
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName string `json:"baseline_table_name,omitempty"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics []monitorMetricPb `json:"custom_metrics,omitempty"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId string `json:"dashboard_id,omitempty"`
	// The data classification config for the monitor.
	DataClassificationConfig *monitorDataClassificationConfigPb `json:"data_classification_config,omitempty"`
	// The full name of the drift metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	DriftMetricsTableName string `json:"drift_metrics_table_name"`
	// Configuration for monitoring inference logs.
	InferenceLog *monitorInferenceLogPb `json:"inference_log,omitempty"`
	// The latest failure message of the monitor (if any).
	LatestMonitorFailureMsg string `json:"latest_monitor_failure_msg,omitempty"`
	// The version of the monitor config (e.g. 1,2,3). If negative, the monitor
	// may be corrupted.
	MonitorVersion string `json:"monitor_version"`
	// The notification settings for the monitor.
	Notifications *monitorNotificationsPb `json:"notifications,omitempty"`
	// Schema where output metric tables are created.
	OutputSchemaName string `json:"output_schema_name,omitempty"`
	// The full name of the profile metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	ProfileMetricsTableName string `json:"profile_metrics_table_name"`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule *monitorCronSchedulePb `json:"schedule,omitempty"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs []string `json:"slicing_exprs,omitempty"`
	// Configuration for monitoring snapshot tables.
	Snapshot *monitorSnapshotPb `json:"snapshot,omitempty"`
	// The status of the monitor.
	Status MonitorInfoStatus `json:"status"`
	// The full name of the table to monitor. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	TableName string `json:"table_name"`
	// Configuration for monitoring time series tables.
	TimeSeries *monitorTimeSeriesPb `json:"time_series,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func monitorInfoFromPb(pb *monitorInfoPb) (*MonitorInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorInfo{}
	st.AssetsDir = pb.AssetsDir
	st.BaselineTableName = pb.BaselineTableName

	var customMetricsField []MonitorMetric
	for _, itemPb := range pb.CustomMetrics {
		item, err := monitorMetricFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customMetricsField = append(customMetricsField, *item)
		}
	}
	st.CustomMetrics = customMetricsField
	st.DashboardId = pb.DashboardId
	dataClassificationConfigField, err := monitorDataClassificationConfigFromPb(pb.DataClassificationConfig)
	if err != nil {
		return nil, err
	}
	if dataClassificationConfigField != nil {
		st.DataClassificationConfig = dataClassificationConfigField
	}
	st.DriftMetricsTableName = pb.DriftMetricsTableName
	inferenceLogField, err := monitorInferenceLogFromPb(pb.InferenceLog)
	if err != nil {
		return nil, err
	}
	if inferenceLogField != nil {
		st.InferenceLog = inferenceLogField
	}
	st.LatestMonitorFailureMsg = pb.LatestMonitorFailureMsg
	st.MonitorVersion = pb.MonitorVersion
	notificationsField, err := monitorNotificationsFromPb(pb.Notifications)
	if err != nil {
		return nil, err
	}
	if notificationsField != nil {
		st.Notifications = notificationsField
	}
	st.OutputSchemaName = pb.OutputSchemaName
	st.ProfileMetricsTableName = pb.ProfileMetricsTableName
	scheduleField, err := monitorCronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	st.SlicingExprs = pb.SlicingExprs
	snapshotField, err := monitorSnapshotFromPb(pb.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotField != nil {
		st.Snapshot = snapshotField
	}
	st.Status = pb.Status
	st.TableName = pb.TableName
	timeSeriesField, err := monitorTimeSeriesFromPb(pb.TimeSeries)
	if err != nil {
		return nil, err
	}
	if timeSeriesField != nil {
		st.TimeSeries = timeSeriesField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *monitorInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st monitorInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The status of the monitor.
type MonitorInfoStatus string
type monitorInfoStatusPb string

const MonitorInfoStatusMonitorStatusActive MonitorInfoStatus = `MONITOR_STATUS_ACTIVE`

const MonitorInfoStatusMonitorStatusDeletePending MonitorInfoStatus = `MONITOR_STATUS_DELETE_PENDING`

const MonitorInfoStatusMonitorStatusError MonitorInfoStatus = `MONITOR_STATUS_ERROR`

const MonitorInfoStatusMonitorStatusFailed MonitorInfoStatus = `MONITOR_STATUS_FAILED`

const MonitorInfoStatusMonitorStatusPending MonitorInfoStatus = `MONITOR_STATUS_PENDING`

// String representation for [fmt.Print]
func (f *MonitorInfoStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorInfoStatus) Set(v string) error {
	switch v {
	case `MONITOR_STATUS_ACTIVE`, `MONITOR_STATUS_DELETE_PENDING`, `MONITOR_STATUS_ERROR`, `MONITOR_STATUS_FAILED`, `MONITOR_STATUS_PENDING`:
		*f = MonitorInfoStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MONITOR_STATUS_ACTIVE", "MONITOR_STATUS_DELETE_PENDING", "MONITOR_STATUS_ERROR", "MONITOR_STATUS_FAILED", "MONITOR_STATUS_PENDING"`, v)
	}
}

// Type always returns MonitorInfoStatus to satisfy [pflag.Value] interface
func (f *MonitorInfoStatus) Type() string {
	return "MonitorInfoStatus"
}

func monitorInfoStatusToPb(st *MonitorInfoStatus) (*monitorInfoStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := monitorInfoStatusPb(*st)
	return &pb, nil
}

func monitorInfoStatusFromPb(pb *monitorInfoStatusPb) (*MonitorInfoStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := MonitorInfoStatus(*pb)
	return &st, nil
}

type MonitorMetric struct {
	// Jinja template for a SQL expression that specifies how to compute the
	// metric. See [create metric definition].
	//
	// [create metric definition]: https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition
	// Wire name: 'definition'
	Definition string
	// A list of column names in the input table the metric should be computed
	// for. Can use ``":table"`` to indicate that the metric needs information
	// from multiple columns.
	// Wire name: 'input_columns'
	InputColumns []string
	// Name of the metric in the output tables.
	// Wire name: 'name'
	Name string
	// The output type of the custom metric.
	// Wire name: 'output_data_type'
	OutputDataType string
	// Can only be one of ``"CUSTOM_METRIC_TYPE_AGGREGATE"``,
	// ``"CUSTOM_METRIC_TYPE_DERIVED"``, or ``"CUSTOM_METRIC_TYPE_DRIFT"``. The
	// ``"CUSTOM_METRIC_TYPE_AGGREGATE"`` and ``"CUSTOM_METRIC_TYPE_DERIVED"``
	// metrics are computed on a single table, whereas the
	// ``"CUSTOM_METRIC_TYPE_DRIFT"`` compare metrics across baseline and input
	// table, or across the two consecutive time windows. -
	// CUSTOM_METRIC_TYPE_AGGREGATE: only depend on the existing columns in your
	// table - CUSTOM_METRIC_TYPE_DERIVED: depend on previously computed
	// aggregate metrics - CUSTOM_METRIC_TYPE_DRIFT: depend on previously
	// computed aggregate or derived metrics
	// Wire name: 'type'
	Type MonitorMetricType
}

func monitorMetricToPb(st *MonitorMetric) (*monitorMetricPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorMetricPb{}
	pb.Definition = st.Definition

	pb.InputColumns = st.InputColumns

	pb.Name = st.Name

	pb.OutputDataType = st.OutputDataType

	pb.Type = st.Type

	return pb, nil
}

func (st *MonitorMetric) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &monitorMetricPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := monitorMetricFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MonitorMetric) MarshalJSON() ([]byte, error) {
	pb, err := monitorMetricToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type monitorMetricPb struct {
	// Jinja template for a SQL expression that specifies how to compute the
	// metric. See [create metric definition].
	//
	// [create metric definition]: https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition
	Definition string `json:"definition"`
	// A list of column names in the input table the metric should be computed
	// for. Can use ``":table"`` to indicate that the metric needs information
	// from multiple columns.
	InputColumns []string `json:"input_columns"`
	// Name of the metric in the output tables.
	Name string `json:"name"`
	// The output type of the custom metric.
	OutputDataType string `json:"output_data_type"`
	// Can only be one of ``"CUSTOM_METRIC_TYPE_AGGREGATE"``,
	// ``"CUSTOM_METRIC_TYPE_DERIVED"``, or ``"CUSTOM_METRIC_TYPE_DRIFT"``. The
	// ``"CUSTOM_METRIC_TYPE_AGGREGATE"`` and ``"CUSTOM_METRIC_TYPE_DERIVED"``
	// metrics are computed on a single table, whereas the
	// ``"CUSTOM_METRIC_TYPE_DRIFT"`` compare metrics across baseline and input
	// table, or across the two consecutive time windows. -
	// CUSTOM_METRIC_TYPE_AGGREGATE: only depend on the existing columns in your
	// table - CUSTOM_METRIC_TYPE_DERIVED: depend on previously computed
	// aggregate metrics - CUSTOM_METRIC_TYPE_DRIFT: depend on previously
	// computed aggregate or derived metrics
	Type MonitorMetricType `json:"type"`
}

func monitorMetricFromPb(pb *monitorMetricPb) (*MonitorMetric, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorMetric{}
	st.Definition = pb.Definition
	st.InputColumns = pb.InputColumns
	st.Name = pb.Name
	st.OutputDataType = pb.OutputDataType
	st.Type = pb.Type

	return st, nil
}

// Can only be one of “"CUSTOM_METRIC_TYPE_AGGREGATE"“,
// “"CUSTOM_METRIC_TYPE_DERIVED"“, or “"CUSTOM_METRIC_TYPE_DRIFT"“. The
// “"CUSTOM_METRIC_TYPE_AGGREGATE"“ and “"CUSTOM_METRIC_TYPE_DERIVED"“
// metrics are computed on a single table, whereas the
// “"CUSTOM_METRIC_TYPE_DRIFT"“ compare metrics across baseline and input
// table, or across the two consecutive time windows. -
// CUSTOM_METRIC_TYPE_AGGREGATE: only depend on the existing columns in your
// table - CUSTOM_METRIC_TYPE_DERIVED: depend on previously computed aggregate
// metrics - CUSTOM_METRIC_TYPE_DRIFT: depend on previously computed aggregate
// or derived metrics
type MonitorMetricType string
type monitorMetricTypePb string

const MonitorMetricTypeCustomMetricTypeAggregate MonitorMetricType = `CUSTOM_METRIC_TYPE_AGGREGATE`

const MonitorMetricTypeCustomMetricTypeDerived MonitorMetricType = `CUSTOM_METRIC_TYPE_DERIVED`

const MonitorMetricTypeCustomMetricTypeDrift MonitorMetricType = `CUSTOM_METRIC_TYPE_DRIFT`

// String representation for [fmt.Print]
func (f *MonitorMetricType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorMetricType) Set(v string) error {
	switch v {
	case `CUSTOM_METRIC_TYPE_AGGREGATE`, `CUSTOM_METRIC_TYPE_DERIVED`, `CUSTOM_METRIC_TYPE_DRIFT`:
		*f = MonitorMetricType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CUSTOM_METRIC_TYPE_AGGREGATE", "CUSTOM_METRIC_TYPE_DERIVED", "CUSTOM_METRIC_TYPE_DRIFT"`, v)
	}
}

// Type always returns MonitorMetricType to satisfy [pflag.Value] interface
func (f *MonitorMetricType) Type() string {
	return "MonitorMetricType"
}

func monitorMetricTypeToPb(st *MonitorMetricType) (*monitorMetricTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := monitorMetricTypePb(*st)
	return &pb, nil
}

func monitorMetricTypeFromPb(pb *monitorMetricTypePb) (*MonitorMetricType, error) {
	if pb == nil {
		return nil, nil
	}
	st := MonitorMetricType(*pb)
	return &st, nil
}

type MonitorNotifications struct {
	// Who to send notifications to on monitor failure.
	// Wire name: 'on_failure'
	OnFailure *MonitorDestination
	// Who to send notifications to when new data classification tags are
	// detected.
	// Wire name: 'on_new_classification_tag_detected'
	OnNewClassificationTagDetected *MonitorDestination
}

func monitorNotificationsToPb(st *MonitorNotifications) (*monitorNotificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorNotificationsPb{}
	onFailurePb, err := monitorDestinationToPb(st.OnFailure)
	if err != nil {
		return nil, err
	}
	if onFailurePb != nil {
		pb.OnFailure = onFailurePb
	}

	onNewClassificationTagDetectedPb, err := monitorDestinationToPb(st.OnNewClassificationTagDetected)
	if err != nil {
		return nil, err
	}
	if onNewClassificationTagDetectedPb != nil {
		pb.OnNewClassificationTagDetected = onNewClassificationTagDetectedPb
	}

	return pb, nil
}

func (st *MonitorNotifications) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &monitorNotificationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := monitorNotificationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MonitorNotifications) MarshalJSON() ([]byte, error) {
	pb, err := monitorNotificationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type monitorNotificationsPb struct {
	// Who to send notifications to on monitor failure.
	OnFailure *monitorDestinationPb `json:"on_failure,omitempty"`
	// Who to send notifications to when new data classification tags are
	// detected.
	OnNewClassificationTagDetected *monitorDestinationPb `json:"on_new_classification_tag_detected,omitempty"`
}

func monitorNotificationsFromPb(pb *monitorNotificationsPb) (*MonitorNotifications, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorNotifications{}
	onFailureField, err := monitorDestinationFromPb(pb.OnFailure)
	if err != nil {
		return nil, err
	}
	if onFailureField != nil {
		st.OnFailure = onFailureField
	}
	onNewClassificationTagDetectedField, err := monitorDestinationFromPb(pb.OnNewClassificationTagDetected)
	if err != nil {
		return nil, err
	}
	if onNewClassificationTagDetectedField != nil {
		st.OnNewClassificationTagDetected = onNewClassificationTagDetectedField
	}

	return st, nil
}

type MonitorRefreshInfo struct {
	// Time at which refresh operation completed (milliseconds since 1/1/1970
	// UTC).
	// Wire name: 'end_time_ms'
	EndTimeMs int64
	// An optional message to give insight into the current state of the job
	// (e.g. FAILURE messages).
	// Wire name: 'message'
	Message string
	// Unique id of the refresh operation.
	// Wire name: 'refresh_id'
	RefreshId int64
	// Time at which refresh operation was initiated (milliseconds since
	// 1/1/1970 UTC).
	// Wire name: 'start_time_ms'
	StartTimeMs int64
	// The current state of the refresh.
	// Wire name: 'state'
	State MonitorRefreshInfoState
	// The method by which the refresh was triggered.
	// Wire name: 'trigger'
	Trigger MonitorRefreshInfoTrigger

	ForceSendFields []string `tf:"-"`
}

func monitorRefreshInfoToPb(st *MonitorRefreshInfo) (*monitorRefreshInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorRefreshInfoPb{}
	pb.EndTimeMs = st.EndTimeMs

	pb.Message = st.Message

	pb.RefreshId = st.RefreshId

	pb.StartTimeMs = st.StartTimeMs

	pb.State = st.State

	pb.Trigger = st.Trigger

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *MonitorRefreshInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &monitorRefreshInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := monitorRefreshInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MonitorRefreshInfo) MarshalJSON() ([]byte, error) {
	pb, err := monitorRefreshInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type monitorRefreshInfoPb struct {
	// Time at which refresh operation completed (milliseconds since 1/1/1970
	// UTC).
	EndTimeMs int64 `json:"end_time_ms,omitempty"`
	// An optional message to give insight into the current state of the job
	// (e.g. FAILURE messages).
	Message string `json:"message,omitempty"`
	// Unique id of the refresh operation.
	RefreshId int64 `json:"refresh_id"`
	// Time at which refresh operation was initiated (milliseconds since
	// 1/1/1970 UTC).
	StartTimeMs int64 `json:"start_time_ms"`
	// The current state of the refresh.
	State MonitorRefreshInfoState `json:"state"`
	// The method by which the refresh was triggered.
	Trigger MonitorRefreshInfoTrigger `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func monitorRefreshInfoFromPb(pb *monitorRefreshInfoPb) (*MonitorRefreshInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorRefreshInfo{}
	st.EndTimeMs = pb.EndTimeMs
	st.Message = pb.Message
	st.RefreshId = pb.RefreshId
	st.StartTimeMs = pb.StartTimeMs
	st.State = pb.State
	st.Trigger = pb.Trigger

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *monitorRefreshInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st monitorRefreshInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The current state of the refresh.
type MonitorRefreshInfoState string
type monitorRefreshInfoStatePb string

const MonitorRefreshInfoStateCanceled MonitorRefreshInfoState = `CANCELED`

const MonitorRefreshInfoStateFailed MonitorRefreshInfoState = `FAILED`

const MonitorRefreshInfoStatePending MonitorRefreshInfoState = `PENDING`

const MonitorRefreshInfoStateRunning MonitorRefreshInfoState = `RUNNING`

const MonitorRefreshInfoStateSuccess MonitorRefreshInfoState = `SUCCESS`

// String representation for [fmt.Print]
func (f *MonitorRefreshInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorRefreshInfoState) Set(v string) error {
	switch v {
	case `CANCELED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCESS`:
		*f = MonitorRefreshInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "FAILED", "PENDING", "RUNNING", "SUCCESS"`, v)
	}
}

// Type always returns MonitorRefreshInfoState to satisfy [pflag.Value] interface
func (f *MonitorRefreshInfoState) Type() string {
	return "MonitorRefreshInfoState"
}

func monitorRefreshInfoStateToPb(st *MonitorRefreshInfoState) (*monitorRefreshInfoStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := monitorRefreshInfoStatePb(*st)
	return &pb, nil
}

func monitorRefreshInfoStateFromPb(pb *monitorRefreshInfoStatePb) (*MonitorRefreshInfoState, error) {
	if pb == nil {
		return nil, nil
	}
	st := MonitorRefreshInfoState(*pb)
	return &st, nil
}

// The method by which the refresh was triggered.
type MonitorRefreshInfoTrigger string
type monitorRefreshInfoTriggerPb string

const MonitorRefreshInfoTriggerManual MonitorRefreshInfoTrigger = `MANUAL`

const MonitorRefreshInfoTriggerSchedule MonitorRefreshInfoTrigger = `SCHEDULE`

// String representation for [fmt.Print]
func (f *MonitorRefreshInfoTrigger) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorRefreshInfoTrigger) Set(v string) error {
	switch v {
	case `MANUAL`, `SCHEDULE`:
		*f = MonitorRefreshInfoTrigger(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MANUAL", "SCHEDULE"`, v)
	}
}

// Type always returns MonitorRefreshInfoTrigger to satisfy [pflag.Value] interface
func (f *MonitorRefreshInfoTrigger) Type() string {
	return "MonitorRefreshInfoTrigger"
}

func monitorRefreshInfoTriggerToPb(st *MonitorRefreshInfoTrigger) (*monitorRefreshInfoTriggerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := monitorRefreshInfoTriggerPb(*st)
	return &pb, nil
}

func monitorRefreshInfoTriggerFromPb(pb *monitorRefreshInfoTriggerPb) (*MonitorRefreshInfoTrigger, error) {
	if pb == nil {
		return nil, nil
	}
	st := MonitorRefreshInfoTrigger(*pb)
	return &st, nil
}

type MonitorRefreshListResponse struct {
	// List of refreshes.
	// Wire name: 'refreshes'
	Refreshes []MonitorRefreshInfo
}

func monitorRefreshListResponseToPb(st *MonitorRefreshListResponse) (*monitorRefreshListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorRefreshListResponsePb{}

	var refreshesPb []monitorRefreshInfoPb
	for _, item := range st.Refreshes {
		itemPb, err := monitorRefreshInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			refreshesPb = append(refreshesPb, *itemPb)
		}
	}
	pb.Refreshes = refreshesPb

	return pb, nil
}

func (st *MonitorRefreshListResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &monitorRefreshListResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := monitorRefreshListResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MonitorRefreshListResponse) MarshalJSON() ([]byte, error) {
	pb, err := monitorRefreshListResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type monitorRefreshListResponsePb struct {
	// List of refreshes.
	Refreshes []monitorRefreshInfoPb `json:"refreshes,omitempty"`
}

func monitorRefreshListResponseFromPb(pb *monitorRefreshListResponsePb) (*MonitorRefreshListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorRefreshListResponse{}

	var refreshesField []MonitorRefreshInfo
	for _, itemPb := range pb.Refreshes {
		item, err := monitorRefreshInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			refreshesField = append(refreshesField, *item)
		}
	}
	st.Refreshes = refreshesField

	return st, nil
}

type MonitorSnapshot struct {
}

func monitorSnapshotToPb(st *MonitorSnapshot) (*monitorSnapshotPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorSnapshotPb{}

	return pb, nil
}

func (st *MonitorSnapshot) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &monitorSnapshotPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := monitorSnapshotFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MonitorSnapshot) MarshalJSON() ([]byte, error) {
	pb, err := monitorSnapshotToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type monitorSnapshotPb struct {
}

func monitorSnapshotFromPb(pb *monitorSnapshotPb) (*MonitorSnapshot, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorSnapshot{}

	return st, nil
}

type MonitorTimeSeries struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	// Wire name: 'granularities'
	Granularities []string
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	// Wire name: 'timestamp_col'
	TimestampCol string
}

func monitorTimeSeriesToPb(st *MonitorTimeSeries) (*monitorTimeSeriesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorTimeSeriesPb{}
	pb.Granularities = st.Granularities

	pb.TimestampCol = st.TimestampCol

	return pb, nil
}

func (st *MonitorTimeSeries) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &monitorTimeSeriesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := monitorTimeSeriesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MonitorTimeSeries) MarshalJSON() ([]byte, error) {
	pb, err := monitorTimeSeriesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type monitorTimeSeriesPb struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	Granularities []string `json:"granularities"`
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	TimestampCol string `json:"timestamp_col"`
}

func monitorTimeSeriesFromPb(pb *monitorTimeSeriesPb) (*MonitorTimeSeries, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorTimeSeries{}
	st.Granularities = pb.Granularities
	st.TimestampCol = pb.TimestampCol

	return st, nil
}

type NamedTableConstraint struct {
	// The name of the constraint.
	// Wire name: 'name'
	Name string
}

func namedTableConstraintToPb(st *NamedTableConstraint) (*namedTableConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &namedTableConstraintPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *NamedTableConstraint) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &namedTableConstraintPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := namedTableConstraintFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NamedTableConstraint) MarshalJSON() ([]byte, error) {
	pb, err := namedTableConstraintToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type namedTableConstraintPb struct {
	// The name of the constraint.
	Name string `json:"name"`
}

func namedTableConstraintFromPb(pb *namedTableConstraintPb) (*NamedTableConstraint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NamedTableConstraint{}
	st.Name = pb.Name

	return st, nil
}

// Online Table information.
type OnlineTable struct {
	// Full three-part (catalog, schema, table) name of the table.
	// Wire name: 'name'
	Name string
	// Specification of the online table.
	// Wire name: 'spec'
	Spec *OnlineTableSpec
	// Online Table data synchronization status
	// Wire name: 'status'
	Status *OnlineTableStatus
	// Data serving REST API URL for this table
	// Wire name: 'table_serving_url'
	TableServingUrl string
	// The provisioning state of the online table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	// Wire name: 'unity_catalog_provisioning_state'
	UnityCatalogProvisioningState ProvisioningInfoState

	ForceSendFields []string `tf:"-"`
}

func onlineTableToPb(st *OnlineTable) (*onlineTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &onlineTablePb{}
	pb.Name = st.Name

	specPb, err := onlineTableSpecToPb(st.Spec)
	if err != nil {
		return nil, err
	}
	if specPb != nil {
		pb.Spec = specPb
	}

	statusPb, err := onlineTableStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}

	pb.TableServingUrl = st.TableServingUrl

	pb.UnityCatalogProvisioningState = st.UnityCatalogProvisioningState

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *OnlineTable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &onlineTablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := onlineTableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st OnlineTable) MarshalJSON() ([]byte, error) {
	pb, err := onlineTableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type onlineTablePb struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"name,omitempty"`
	// Specification of the online table.
	Spec *onlineTableSpecPb `json:"spec,omitempty"`
	// Online Table data synchronization status
	Status *onlineTableStatusPb `json:"status,omitempty"`
	// Data serving REST API URL for this table
	TableServingUrl string `json:"table_serving_url,omitempty"`
	// The provisioning state of the online table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	UnityCatalogProvisioningState ProvisioningInfoState `json:"unity_catalog_provisioning_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func onlineTableFromPb(pb *onlineTablePb) (*OnlineTable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineTable{}
	st.Name = pb.Name
	specField, err := onlineTableSpecFromPb(pb.Spec)
	if err != nil {
		return nil, err
	}
	if specField != nil {
		st.Spec = specField
	}
	statusField, err := onlineTableStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}
	st.TableServingUrl = pb.TableServingUrl
	st.UnityCatalogProvisioningState = pb.UnityCatalogProvisioningState

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *onlineTablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st onlineTablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Specification of an online table.
type OnlineTableSpec struct {
	// Whether to create a full-copy pipeline -- a pipeline that stops after
	// creates a full copy of the source table upon initialization and does not
	// process any change data feeds (CDFs) afterwards. The pipeline can still
	// be manually triggered afterwards, but it always perform a full copy of
	// the source table and there are no incremental updates. This mode is
	// useful for syncing views or tables without CDFs to online tables. Note
	// that the full-copy pipeline only supports "triggered" scheduling policy.
	// Wire name: 'perform_full_copy'
	PerformFullCopy bool
	// ID of the associated pipeline. Generated by the server - cannot be set by
	// the caller.
	// Wire name: 'pipeline_id'
	PipelineId string
	// Primary Key columns to be used for data insert/update in the destination.
	// Wire name: 'primary_key_columns'
	PrimaryKeyColumns []string
	// Pipeline runs continuously after generating the initial data.
	// Wire name: 'run_continuously'
	RunContinuously *OnlineTableSpecContinuousSchedulingPolicy
	// Pipeline stops after generating the initial data and can be triggered
	// later (manually, through a cron job or through data triggers)
	// Wire name: 'run_triggered'
	RunTriggered *OnlineTableSpecTriggeredSchedulingPolicy
	// Three-part (catalog, schema, table) name of the source Delta table.
	// Wire name: 'source_table_full_name'
	SourceTableFullName string
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	// Wire name: 'timeseries_key'
	TimeseriesKey string

	ForceSendFields []string `tf:"-"`
}

func onlineTableSpecToPb(st *OnlineTableSpec) (*onlineTableSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &onlineTableSpecPb{}
	pb.PerformFullCopy = st.PerformFullCopy

	pb.PipelineId = st.PipelineId

	pb.PrimaryKeyColumns = st.PrimaryKeyColumns

	runContinuouslyPb, err := onlineTableSpecContinuousSchedulingPolicyToPb(st.RunContinuously)
	if err != nil {
		return nil, err
	}
	if runContinuouslyPb != nil {
		pb.RunContinuously = runContinuouslyPb
	}

	runTriggeredPb, err := onlineTableSpecTriggeredSchedulingPolicyToPb(st.RunTriggered)
	if err != nil {
		return nil, err
	}
	if runTriggeredPb != nil {
		pb.RunTriggered = runTriggeredPb
	}

	pb.SourceTableFullName = st.SourceTableFullName

	pb.TimeseriesKey = st.TimeseriesKey

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *OnlineTableSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &onlineTableSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := onlineTableSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st OnlineTableSpec) MarshalJSON() ([]byte, error) {
	pb, err := onlineTableSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type onlineTableSpecPb struct {
	// Whether to create a full-copy pipeline -- a pipeline that stops after
	// creates a full copy of the source table upon initialization and does not
	// process any change data feeds (CDFs) afterwards. The pipeline can still
	// be manually triggered afterwards, but it always perform a full copy of
	// the source table and there are no incremental updates. This mode is
	// useful for syncing views or tables without CDFs to online tables. Note
	// that the full-copy pipeline only supports "triggered" scheduling policy.
	PerformFullCopy bool `json:"perform_full_copy,omitempty"`
	// ID of the associated pipeline. Generated by the server - cannot be set by
	// the caller.
	PipelineId string `json:"pipeline_id,omitempty"`
	// Primary Key columns to be used for data insert/update in the destination.
	PrimaryKeyColumns []string `json:"primary_key_columns,omitempty"`
	// Pipeline runs continuously after generating the initial data.
	RunContinuously *onlineTableSpecContinuousSchedulingPolicyPb `json:"run_continuously,omitempty"`
	// Pipeline stops after generating the initial data and can be triggered
	// later (manually, through a cron job or through data triggers)
	RunTriggered *onlineTableSpecTriggeredSchedulingPolicyPb `json:"run_triggered,omitempty"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	SourceTableFullName string `json:"source_table_full_name,omitempty"`
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	TimeseriesKey string `json:"timeseries_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func onlineTableSpecFromPb(pb *onlineTableSpecPb) (*OnlineTableSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineTableSpec{}
	st.PerformFullCopy = pb.PerformFullCopy
	st.PipelineId = pb.PipelineId
	st.PrimaryKeyColumns = pb.PrimaryKeyColumns
	runContinuouslyField, err := onlineTableSpecContinuousSchedulingPolicyFromPb(pb.RunContinuously)
	if err != nil {
		return nil, err
	}
	if runContinuouslyField != nil {
		st.RunContinuously = runContinuouslyField
	}
	runTriggeredField, err := onlineTableSpecTriggeredSchedulingPolicyFromPb(pb.RunTriggered)
	if err != nil {
		return nil, err
	}
	if runTriggeredField != nil {
		st.RunTriggered = runTriggeredField
	}
	st.SourceTableFullName = pb.SourceTableFullName
	st.TimeseriesKey = pb.TimeseriesKey

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *onlineTableSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st onlineTableSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type OnlineTableSpecContinuousSchedulingPolicy struct {
}

func onlineTableSpecContinuousSchedulingPolicyToPb(st *OnlineTableSpecContinuousSchedulingPolicy) (*onlineTableSpecContinuousSchedulingPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &onlineTableSpecContinuousSchedulingPolicyPb{}

	return pb, nil
}

func (st *OnlineTableSpecContinuousSchedulingPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &onlineTableSpecContinuousSchedulingPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := onlineTableSpecContinuousSchedulingPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st OnlineTableSpecContinuousSchedulingPolicy) MarshalJSON() ([]byte, error) {
	pb, err := onlineTableSpecContinuousSchedulingPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type onlineTableSpecContinuousSchedulingPolicyPb struct {
}

func onlineTableSpecContinuousSchedulingPolicyFromPb(pb *onlineTableSpecContinuousSchedulingPolicyPb) (*OnlineTableSpecContinuousSchedulingPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineTableSpecContinuousSchedulingPolicy{}

	return st, nil
}

type OnlineTableSpecTriggeredSchedulingPolicy struct {
}

func onlineTableSpecTriggeredSchedulingPolicyToPb(st *OnlineTableSpecTriggeredSchedulingPolicy) (*onlineTableSpecTriggeredSchedulingPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &onlineTableSpecTriggeredSchedulingPolicyPb{}

	return pb, nil
}

func (st *OnlineTableSpecTriggeredSchedulingPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &onlineTableSpecTriggeredSchedulingPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := onlineTableSpecTriggeredSchedulingPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st OnlineTableSpecTriggeredSchedulingPolicy) MarshalJSON() ([]byte, error) {
	pb, err := onlineTableSpecTriggeredSchedulingPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type onlineTableSpecTriggeredSchedulingPolicyPb struct {
}

func onlineTableSpecTriggeredSchedulingPolicyFromPb(pb *onlineTableSpecTriggeredSchedulingPolicyPb) (*OnlineTableSpecTriggeredSchedulingPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineTableSpecTriggeredSchedulingPolicy{}

	return st, nil
}

// The state of an online table.
type OnlineTableState string
type onlineTableStatePb string

const OnlineTableStateOffline OnlineTableState = `OFFLINE`

const OnlineTableStateOfflineFailed OnlineTableState = `OFFLINE_FAILED`

const OnlineTableStateOnline OnlineTableState = `ONLINE`

const OnlineTableStateOnlineContinuousUpdate OnlineTableState = `ONLINE_CONTINUOUS_UPDATE`

const OnlineTableStateOnlineNoPendingUpdate OnlineTableState = `ONLINE_NO_PENDING_UPDATE`

const OnlineTableStateOnlinePipelineFailed OnlineTableState = `ONLINE_PIPELINE_FAILED`

const OnlineTableStateOnlineTriggeredUpdate OnlineTableState = `ONLINE_TRIGGERED_UPDATE`

const OnlineTableStateOnlineUpdatingPipelineResources OnlineTableState = `ONLINE_UPDATING_PIPELINE_RESOURCES`

const OnlineTableStateProvisioning OnlineTableState = `PROVISIONING`

const OnlineTableStateProvisioningInitialSnapshot OnlineTableState = `PROVISIONING_INITIAL_SNAPSHOT`

const OnlineTableStateProvisioningPipelineResources OnlineTableState = `PROVISIONING_PIPELINE_RESOURCES`

// String representation for [fmt.Print]
func (f *OnlineTableState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OnlineTableState) Set(v string) error {
	switch v {
	case `OFFLINE`, `OFFLINE_FAILED`, `ONLINE`, `ONLINE_CONTINUOUS_UPDATE`, `ONLINE_NO_PENDING_UPDATE`, `ONLINE_PIPELINE_FAILED`, `ONLINE_TRIGGERED_UPDATE`, `ONLINE_UPDATING_PIPELINE_RESOURCES`, `PROVISIONING`, `PROVISIONING_INITIAL_SNAPSHOT`, `PROVISIONING_PIPELINE_RESOURCES`:
		*f = OnlineTableState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OFFLINE", "OFFLINE_FAILED", "ONLINE", "ONLINE_CONTINUOUS_UPDATE", "ONLINE_NO_PENDING_UPDATE", "ONLINE_PIPELINE_FAILED", "ONLINE_TRIGGERED_UPDATE", "ONLINE_UPDATING_PIPELINE_RESOURCES", "PROVISIONING", "PROVISIONING_INITIAL_SNAPSHOT", "PROVISIONING_PIPELINE_RESOURCES"`, v)
	}
}

// Type always returns OnlineTableState to satisfy [pflag.Value] interface
func (f *OnlineTableState) Type() string {
	return "OnlineTableState"
}

func onlineTableStateToPb(st *OnlineTableState) (*onlineTableStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := onlineTableStatePb(*st)
	return &pb, nil
}

func onlineTableStateFromPb(pb *onlineTableStatePb) (*OnlineTableState, error) {
	if pb == nil {
		return nil, nil
	}
	st := OnlineTableState(*pb)
	return &st, nil
}

// Status of an online table.
type OnlineTableStatus struct {
	// Detailed status of an online table. Shown if the online table is in the
	// ONLINE_CONTINUOUS_UPDATE or the ONLINE_UPDATING_PIPELINE_RESOURCES state.
	// Wire name: 'continuous_update_status'
	ContinuousUpdateStatus *ContinuousUpdateStatus
	// The state of the online table.
	// Wire name: 'detailed_state'
	DetailedState OnlineTableState
	// Detailed status of an online table. Shown if the online table is in the
	// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
	// Wire name: 'failed_status'
	FailedStatus *FailedStatus
	// A text description of the current state of the online table.
	// Wire name: 'message'
	Message string
	// Detailed status of an online table. Shown if the online table is in the
	// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT
	// state.
	// Wire name: 'provisioning_status'
	ProvisioningStatus *ProvisioningStatus
	// Detailed status of an online table. Shown if the online table is in the
	// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
	// Wire name: 'triggered_update_status'
	TriggeredUpdateStatus *TriggeredUpdateStatus

	ForceSendFields []string `tf:"-"`
}

func onlineTableStatusToPb(st *OnlineTableStatus) (*onlineTableStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &onlineTableStatusPb{}
	continuousUpdateStatusPb, err := continuousUpdateStatusToPb(st.ContinuousUpdateStatus)
	if err != nil {
		return nil, err
	}
	if continuousUpdateStatusPb != nil {
		pb.ContinuousUpdateStatus = continuousUpdateStatusPb
	}

	pb.DetailedState = st.DetailedState

	failedStatusPb, err := failedStatusToPb(st.FailedStatus)
	if err != nil {
		return nil, err
	}
	if failedStatusPb != nil {
		pb.FailedStatus = failedStatusPb
	}

	pb.Message = st.Message

	provisioningStatusPb, err := provisioningStatusToPb(st.ProvisioningStatus)
	if err != nil {
		return nil, err
	}
	if provisioningStatusPb != nil {
		pb.ProvisioningStatus = provisioningStatusPb
	}

	triggeredUpdateStatusPb, err := triggeredUpdateStatusToPb(st.TriggeredUpdateStatus)
	if err != nil {
		return nil, err
	}
	if triggeredUpdateStatusPb != nil {
		pb.TriggeredUpdateStatus = triggeredUpdateStatusPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *OnlineTableStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &onlineTableStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := onlineTableStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st OnlineTableStatus) MarshalJSON() ([]byte, error) {
	pb, err := onlineTableStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type onlineTableStatusPb struct {
	// Detailed status of an online table. Shown if the online table is in the
	// ONLINE_CONTINUOUS_UPDATE or the ONLINE_UPDATING_PIPELINE_RESOURCES state.
	ContinuousUpdateStatus *continuousUpdateStatusPb `json:"continuous_update_status,omitempty"`
	// The state of the online table.
	DetailedState OnlineTableState `json:"detailed_state,omitempty"`
	// Detailed status of an online table. Shown if the online table is in the
	// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
	FailedStatus *failedStatusPb `json:"failed_status,omitempty"`
	// A text description of the current state of the online table.
	Message string `json:"message,omitempty"`
	// Detailed status of an online table. Shown if the online table is in the
	// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT
	// state.
	ProvisioningStatus *provisioningStatusPb `json:"provisioning_status,omitempty"`
	// Detailed status of an online table. Shown if the online table is in the
	// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
	TriggeredUpdateStatus *triggeredUpdateStatusPb `json:"triggered_update_status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func onlineTableStatusFromPb(pb *onlineTableStatusPb) (*OnlineTableStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineTableStatus{}
	continuousUpdateStatusField, err := continuousUpdateStatusFromPb(pb.ContinuousUpdateStatus)
	if err != nil {
		return nil, err
	}
	if continuousUpdateStatusField != nil {
		st.ContinuousUpdateStatus = continuousUpdateStatusField
	}
	st.DetailedState = pb.DetailedState
	failedStatusField, err := failedStatusFromPb(pb.FailedStatus)
	if err != nil {
		return nil, err
	}
	if failedStatusField != nil {
		st.FailedStatus = failedStatusField
	}
	st.Message = pb.Message
	provisioningStatusField, err := provisioningStatusFromPb(pb.ProvisioningStatus)
	if err != nil {
		return nil, err
	}
	if provisioningStatusField != nil {
		st.ProvisioningStatus = provisioningStatusField
	}
	triggeredUpdateStatusField, err := triggeredUpdateStatusFromPb(pb.TriggeredUpdateStatus)
	if err != nil {
		return nil, err
	}
	if triggeredUpdateStatusField != nil {
		st.TriggeredUpdateStatus = triggeredUpdateStatusField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *onlineTableStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st onlineTableStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PermissionsChange struct {
	// The set of privileges to add.
	// Wire name: 'add'
	Add []Privilege
	// The principal whose privileges we are changing.
	// Wire name: 'principal'
	Principal string
	// The set of privileges to remove.
	// Wire name: 'remove'
	Remove []Privilege

	ForceSendFields []string `tf:"-"`
}

func permissionsChangeToPb(st *PermissionsChange) (*permissionsChangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permissionsChangePb{}
	pb.Add = st.Add

	pb.Principal = st.Principal

	pb.Remove = st.Remove

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *PermissionsChange) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &permissionsChangePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := permissionsChangeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PermissionsChange) MarshalJSON() ([]byte, error) {
	pb, err := permissionsChangeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type permissionsChangePb struct {
	// The set of privileges to add.
	Add []Privilege `json:"add,omitempty"`
	// The principal whose privileges we are changing.
	Principal string `json:"principal,omitempty"`
	// The set of privileges to remove.
	Remove []Privilege `json:"remove,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func permissionsChangeFromPb(pb *permissionsChangePb) (*PermissionsChange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionsChange{}
	st.Add = pb.Add
	st.Principal = pb.Principal
	st.Remove = pb.Remove

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *permissionsChangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st permissionsChangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PermissionsList struct {
	// The privileges assigned to each principal
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment
}

func permissionsListToPb(st *PermissionsList) (*permissionsListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permissionsListPb{}

	var privilegeAssignmentsPb []privilegeAssignmentPb
	for _, item := range st.PrivilegeAssignments {
		itemPb, err := privilegeAssignmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			privilegeAssignmentsPb = append(privilegeAssignmentsPb, *itemPb)
		}
	}
	pb.PrivilegeAssignments = privilegeAssignmentsPb

	return pb, nil
}

func (st *PermissionsList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &permissionsListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := permissionsListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PermissionsList) MarshalJSON() ([]byte, error) {
	pb, err := permissionsListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type permissionsListPb struct {
	// The privileges assigned to each principal
	PrivilegeAssignments []privilegeAssignmentPb `json:"privilege_assignments,omitempty"`
}

func permissionsListFromPb(pb *permissionsListPb) (*PermissionsList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionsList{}

	var privilegeAssignmentsField []PrivilegeAssignment
	for _, itemPb := range pb.PrivilegeAssignments {
		item, err := privilegeAssignmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			privilegeAssignmentsField = append(privilegeAssignmentsField, *item)
		}
	}
	st.PrivilegeAssignments = privilegeAssignmentsField

	return st, nil
}

// Progress information of the Online Table data synchronization pipeline.
type PipelineProgress struct {
	// The estimated time remaining to complete this update in seconds.
	// Wire name: 'estimated_completion_time_seconds'
	EstimatedCompletionTimeSeconds float64
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	// Wire name: 'latest_version_currently_processing'
	LatestVersionCurrentlyProcessing int64
	// The completion ratio of this update. This is a number between 0 and 1.
	// Wire name: 'sync_progress_completion'
	SyncProgressCompletion float64
	// The number of rows that have been synced in this update.
	// Wire name: 'synced_row_count'
	SyncedRowCount int64
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	// Wire name: 'total_row_count'
	TotalRowCount int64

	ForceSendFields []string `tf:"-"`
}

func pipelineProgressToPb(st *PipelineProgress) (*pipelineProgressPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineProgressPb{}
	pb.EstimatedCompletionTimeSeconds = st.EstimatedCompletionTimeSeconds

	pb.LatestVersionCurrentlyProcessing = st.LatestVersionCurrentlyProcessing

	pb.SyncProgressCompletion = st.SyncProgressCompletion

	pb.SyncedRowCount = st.SyncedRowCount

	pb.TotalRowCount = st.TotalRowCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *PipelineProgress) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelineProgressPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelineProgressFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelineProgress) MarshalJSON() ([]byte, error) {
	pb, err := pipelineProgressToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type pipelineProgressPb struct {
	// The estimated time remaining to complete this update in seconds.
	EstimatedCompletionTimeSeconds float64 `json:"estimated_completion_time_seconds,omitempty"`
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	LatestVersionCurrentlyProcessing int64 `json:"latest_version_currently_processing,omitempty"`
	// The completion ratio of this update. This is a number between 0 and 1.
	SyncProgressCompletion float64 `json:"sync_progress_completion,omitempty"`
	// The number of rows that have been synced in this update.
	SyncedRowCount int64 `json:"synced_row_count,omitempty"`
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	TotalRowCount int64 `json:"total_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineProgressFromPb(pb *pipelineProgressPb) (*PipelineProgress, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineProgress{}
	st.EstimatedCompletionTimeSeconds = pb.EstimatedCompletionTimeSeconds
	st.LatestVersionCurrentlyProcessing = pb.LatestVersionCurrentlyProcessing
	st.SyncProgressCompletion = pb.SyncProgressCompletion
	st.SyncedRowCount = pb.SyncedRowCount
	st.TotalRowCount = pb.TotalRowCount

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineProgressPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineProgressPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PrimaryKeyConstraint struct {
	// Column names for this constraint.
	// Wire name: 'child_columns'
	ChildColumns []string
	// The name of the constraint.
	// Wire name: 'name'
	Name string
}

func primaryKeyConstraintToPb(st *PrimaryKeyConstraint) (*primaryKeyConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &primaryKeyConstraintPb{}
	pb.ChildColumns = st.ChildColumns

	pb.Name = st.Name

	return pb, nil
}

func (st *PrimaryKeyConstraint) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &primaryKeyConstraintPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := primaryKeyConstraintFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PrimaryKeyConstraint) MarshalJSON() ([]byte, error) {
	pb, err := primaryKeyConstraintToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type primaryKeyConstraintPb struct {
	// Column names for this constraint.
	ChildColumns []string `json:"child_columns"`
	// The name of the constraint.
	Name string `json:"name"`
}

func primaryKeyConstraintFromPb(pb *primaryKeyConstraintPb) (*PrimaryKeyConstraint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PrimaryKeyConstraint{}
	st.ChildColumns = pb.ChildColumns
	st.Name = pb.Name

	return st, nil
}

type Privilege string
type privilegePb string

const PrivilegeAccess Privilege = `ACCESS`

const PrivilegeAllPrivileges Privilege = `ALL_PRIVILEGES`

const PrivilegeApplyTag Privilege = `APPLY_TAG`

const PrivilegeBrowse Privilege = `BROWSE`

const PrivilegeCreate Privilege = `CREATE`

const PrivilegeCreateCatalog Privilege = `CREATE_CATALOG`

const PrivilegeCreateCleanRoom Privilege = `CREATE_CLEAN_ROOM`

const PrivilegeCreateConnection Privilege = `CREATE_CONNECTION`

const PrivilegeCreateExternalLocation Privilege = `CREATE_EXTERNAL_LOCATION`

const PrivilegeCreateExternalTable Privilege = `CREATE_EXTERNAL_TABLE`

const PrivilegeCreateExternalVolume Privilege = `CREATE_EXTERNAL_VOLUME`

const PrivilegeCreateForeignCatalog Privilege = `CREATE_FOREIGN_CATALOG`

const PrivilegeCreateForeignSecurable Privilege = `CREATE_FOREIGN_SECURABLE`

const PrivilegeCreateFunction Privilege = `CREATE_FUNCTION`

const PrivilegeCreateManagedStorage Privilege = `CREATE_MANAGED_STORAGE`

const PrivilegeCreateMaterializedView Privilege = `CREATE_MATERIALIZED_VIEW`

const PrivilegeCreateModel Privilege = `CREATE_MODEL`

const PrivilegeCreateProvider Privilege = `CREATE_PROVIDER`

const PrivilegeCreateRecipient Privilege = `CREATE_RECIPIENT`

const PrivilegeCreateSchema Privilege = `CREATE_SCHEMA`

const PrivilegeCreateServiceCredential Privilege = `CREATE_SERVICE_CREDENTIAL`

const PrivilegeCreateShare Privilege = `CREATE_SHARE`

const PrivilegeCreateStorageCredential Privilege = `CREATE_STORAGE_CREDENTIAL`

const PrivilegeCreateTable Privilege = `CREATE_TABLE`

const PrivilegeCreateView Privilege = `CREATE_VIEW`

const PrivilegeCreateVolume Privilege = `CREATE_VOLUME`

const PrivilegeExecute Privilege = `EXECUTE`

const PrivilegeExecuteCleanRoomTask Privilege = `EXECUTE_CLEAN_ROOM_TASK`

const PrivilegeManage Privilege = `MANAGE`

const PrivilegeManageAllowlist Privilege = `MANAGE_ALLOWLIST`

const PrivilegeModify Privilege = `MODIFY`

const PrivilegeModifyCleanRoom Privilege = `MODIFY_CLEAN_ROOM`

const PrivilegeReadFiles Privilege = `READ_FILES`

const PrivilegeReadPrivateFiles Privilege = `READ_PRIVATE_FILES`

const PrivilegeReadVolume Privilege = `READ_VOLUME`

const PrivilegeRefresh Privilege = `REFRESH`

const PrivilegeSelect Privilege = `SELECT`

const PrivilegeSetSharePermission Privilege = `SET_SHARE_PERMISSION`

const PrivilegeUsage Privilege = `USAGE`

const PrivilegeUseCatalog Privilege = `USE_CATALOG`

const PrivilegeUseConnection Privilege = `USE_CONNECTION`

const PrivilegeUseMarketplaceAssets Privilege = `USE_MARKETPLACE_ASSETS`

const PrivilegeUseProvider Privilege = `USE_PROVIDER`

const PrivilegeUseRecipient Privilege = `USE_RECIPIENT`

const PrivilegeUseSchema Privilege = `USE_SCHEMA`

const PrivilegeUseShare Privilege = `USE_SHARE`

const PrivilegeWriteFiles Privilege = `WRITE_FILES`

const PrivilegeWritePrivateFiles Privilege = `WRITE_PRIVATE_FILES`

const PrivilegeWriteVolume Privilege = `WRITE_VOLUME`

// String representation for [fmt.Print]
func (f *Privilege) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Privilege) Set(v string) error {
	switch v {
	case `ACCESS`, `ALL_PRIVILEGES`, `APPLY_TAG`, `BROWSE`, `CREATE`, `CREATE_CATALOG`, `CREATE_CLEAN_ROOM`, `CREATE_CONNECTION`, `CREATE_EXTERNAL_LOCATION`, `CREATE_EXTERNAL_TABLE`, `CREATE_EXTERNAL_VOLUME`, `CREATE_FOREIGN_CATALOG`, `CREATE_FOREIGN_SECURABLE`, `CREATE_FUNCTION`, `CREATE_MANAGED_STORAGE`, `CREATE_MATERIALIZED_VIEW`, `CREATE_MODEL`, `CREATE_PROVIDER`, `CREATE_RECIPIENT`, `CREATE_SCHEMA`, `CREATE_SERVICE_CREDENTIAL`, `CREATE_SHARE`, `CREATE_STORAGE_CREDENTIAL`, `CREATE_TABLE`, `CREATE_VIEW`, `CREATE_VOLUME`, `EXECUTE`, `EXECUTE_CLEAN_ROOM_TASK`, `MANAGE`, `MANAGE_ALLOWLIST`, `MODIFY`, `MODIFY_CLEAN_ROOM`, `READ_FILES`, `READ_PRIVATE_FILES`, `READ_VOLUME`, `REFRESH`, `SELECT`, `SET_SHARE_PERMISSION`, `USAGE`, `USE_CATALOG`, `USE_CONNECTION`, `USE_MARKETPLACE_ASSETS`, `USE_PROVIDER`, `USE_RECIPIENT`, `USE_SCHEMA`, `USE_SHARE`, `WRITE_FILES`, `WRITE_PRIVATE_FILES`, `WRITE_VOLUME`:
		*f = Privilege(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACCESS", "ALL_PRIVILEGES", "APPLY_TAG", "BROWSE", "CREATE", "CREATE_CATALOG", "CREATE_CLEAN_ROOM", "CREATE_CONNECTION", "CREATE_EXTERNAL_LOCATION", "CREATE_EXTERNAL_TABLE", "CREATE_EXTERNAL_VOLUME", "CREATE_FOREIGN_CATALOG", "CREATE_FOREIGN_SECURABLE", "CREATE_FUNCTION", "CREATE_MANAGED_STORAGE", "CREATE_MATERIALIZED_VIEW", "CREATE_MODEL", "CREATE_PROVIDER", "CREATE_RECIPIENT", "CREATE_SCHEMA", "CREATE_SERVICE_CREDENTIAL", "CREATE_SHARE", "CREATE_STORAGE_CREDENTIAL", "CREATE_TABLE", "CREATE_VIEW", "CREATE_VOLUME", "EXECUTE", "EXECUTE_CLEAN_ROOM_TASK", "MANAGE", "MANAGE_ALLOWLIST", "MODIFY", "MODIFY_CLEAN_ROOM", "READ_FILES", "READ_PRIVATE_FILES", "READ_VOLUME", "REFRESH", "SELECT", "SET_SHARE_PERMISSION", "USAGE", "USE_CATALOG", "USE_CONNECTION", "USE_MARKETPLACE_ASSETS", "USE_PROVIDER", "USE_RECIPIENT", "USE_SCHEMA", "USE_SHARE", "WRITE_FILES", "WRITE_PRIVATE_FILES", "WRITE_VOLUME"`, v)
	}
}

// Type always returns Privilege to satisfy [pflag.Value] interface
func (f *Privilege) Type() string {
	return "Privilege"
}

func privilegeToPb(st *Privilege) (*privilegePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := privilegePb(*st)
	return &pb, nil
}

func privilegeFromPb(pb *privilegePb) (*Privilege, error) {
	if pb == nil {
		return nil, nil
	}
	st := Privilege(*pb)
	return &st, nil
}

type PrivilegeAssignment struct {
	// The principal (user email address or group name).
	// Wire name: 'principal'
	Principal string
	// The privileges assigned to the principal.
	// Wire name: 'privileges'
	Privileges []Privilege

	ForceSendFields []string `tf:"-"`
}

func privilegeAssignmentToPb(st *PrivilegeAssignment) (*privilegeAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &privilegeAssignmentPb{}
	pb.Principal = st.Principal

	pb.Privileges = st.Privileges

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *PrivilegeAssignment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &privilegeAssignmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := privilegeAssignmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PrivilegeAssignment) MarshalJSON() ([]byte, error) {
	pb, err := privilegeAssignmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type privilegeAssignmentPb struct {
	// The principal (user email address or group name).
	Principal string `json:"principal,omitempty"`
	// The privileges assigned to the principal.
	Privileges []Privilege `json:"privileges,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func privilegeAssignmentFromPb(pb *privilegeAssignmentPb) (*PrivilegeAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PrivilegeAssignment{}
	st.Principal = pb.Principal
	st.Privileges = pb.Privileges

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *privilegeAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st privilegeAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// An object containing map of key-value properties attached to the connection.

type PropertiesKvPairs map[string]string
type propertiesKvPairsPb PropertiesKvPairs

func propertiesKvPairsToPb(st *PropertiesKvPairs) (*propertiesKvPairsPb, error) {
	if st == nil {
		return nil, nil
	}
	stPb := propertiesKvPairsPb(*st)
	return &stPb, nil
}
func propertiesKvPairsFromPb(stPb *propertiesKvPairsPb) (*PropertiesKvPairs, error) {
	if stPb == nil {
		return nil, nil
	}
	st := PropertiesKvPairs(*stPb)
	return &st, nil
}

// Status of an asynchronously provisioned resource.
type ProvisioningInfo struct {

	// Wire name: 'state'
	State ProvisioningInfoState
}

func provisioningInfoToPb(st *ProvisioningInfo) (*provisioningInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningInfoPb{}
	pb.State = st.State

	return pb, nil
}

func (st *ProvisioningInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &provisioningInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := provisioningInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ProvisioningInfo) MarshalJSON() ([]byte, error) {
	pb, err := provisioningInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type provisioningInfoPb struct {
	State ProvisioningInfoState `json:"state,omitempty"`
}

func provisioningInfoFromPb(pb *provisioningInfoPb) (*ProvisioningInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ProvisioningInfo{}
	st.State = pb.State

	return st, nil
}

type ProvisioningInfoState string
type provisioningInfoStatePb string

const ProvisioningInfoStateActive ProvisioningInfoState = `ACTIVE`

const ProvisioningInfoStateDegraded ProvisioningInfoState = `DEGRADED`

const ProvisioningInfoStateDeleting ProvisioningInfoState = `DELETING`

const ProvisioningInfoStateFailed ProvisioningInfoState = `FAILED`

const ProvisioningInfoStateProvisioning ProvisioningInfoState = `PROVISIONING`

const ProvisioningInfoStateUpdating ProvisioningInfoState = `UPDATING`

// String representation for [fmt.Print]
func (f *ProvisioningInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ProvisioningInfoState) Set(v string) error {
	switch v {
	case `ACTIVE`, `DEGRADED`, `DELETING`, `FAILED`, `PROVISIONING`, `UPDATING`:
		*f = ProvisioningInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DEGRADED", "DELETING", "FAILED", "PROVISIONING", "UPDATING"`, v)
	}
}

// Type always returns ProvisioningInfoState to satisfy [pflag.Value] interface
func (f *ProvisioningInfoState) Type() string {
	return "ProvisioningInfoState"
}

func provisioningInfoStateToPb(st *ProvisioningInfoState) (*provisioningInfoStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := provisioningInfoStatePb(*st)
	return &pb, nil
}

func provisioningInfoStateFromPb(pb *provisioningInfoStatePb) (*ProvisioningInfoState, error) {
	if pb == nil {
		return nil, nil
	}
	st := ProvisioningInfoState(*pb)
	return &st, nil
}

// Detailed status of an online table. Shown if the online table is in the
// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT state.
type ProvisioningStatus struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	// Wire name: 'initial_pipeline_sync_progress'
	InitialPipelineSyncProgress *PipelineProgress
}

func provisioningStatusToPb(st *ProvisioningStatus) (*provisioningStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningStatusPb{}
	initialPipelineSyncProgressPb, err := pipelineProgressToPb(st.InitialPipelineSyncProgress)
	if err != nil {
		return nil, err
	}
	if initialPipelineSyncProgressPb != nil {
		pb.InitialPipelineSyncProgress = initialPipelineSyncProgressPb
	}

	return pb, nil
}

func (st *ProvisioningStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &provisioningStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := provisioningStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ProvisioningStatus) MarshalJSON() ([]byte, error) {
	pb, err := provisioningStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type provisioningStatusPb struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	InitialPipelineSyncProgress *pipelineProgressPb `json:"initial_pipeline_sync_progress,omitempty"`
}

func provisioningStatusFromPb(pb *provisioningStatusPb) (*ProvisioningStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ProvisioningStatus{}
	initialPipelineSyncProgressField, err := pipelineProgressFromPb(pb.InitialPipelineSyncProgress)
	if err != nil {
		return nil, err
	}
	if initialPipelineSyncProgressField != nil {
		st.InitialPipelineSyncProgress = initialPipelineSyncProgressField
	}

	return st, nil
}

type QuotaInfo struct {
	// The timestamp that indicates when the quota count was last updated.
	// Wire name: 'last_refreshed_at'
	LastRefreshedAt int64
	// Name of the parent resource. Returns metastore ID if the parent is a
	// metastore.
	// Wire name: 'parent_full_name'
	ParentFullName string
	// The quota parent securable type.
	// Wire name: 'parent_securable_type'
	ParentSecurableType SecurableType
	// The current usage of the resource quota.
	// Wire name: 'quota_count'
	QuotaCount int
	// The current limit of the resource quota.
	// Wire name: 'quota_limit'
	QuotaLimit int
	// The name of the quota.
	// Wire name: 'quota_name'
	QuotaName string

	ForceSendFields []string `tf:"-"`
}

func quotaInfoToPb(st *QuotaInfo) (*quotaInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &quotaInfoPb{}
	pb.LastRefreshedAt = st.LastRefreshedAt

	pb.ParentFullName = st.ParentFullName

	pb.ParentSecurableType = st.ParentSecurableType

	pb.QuotaCount = st.QuotaCount

	pb.QuotaLimit = st.QuotaLimit

	pb.QuotaName = st.QuotaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *QuotaInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &quotaInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := quotaInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QuotaInfo) MarshalJSON() ([]byte, error) {
	pb, err := quotaInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type quotaInfoPb struct {
	// The timestamp that indicates when the quota count was last updated.
	LastRefreshedAt int64 `json:"last_refreshed_at,omitempty"`
	// Name of the parent resource. Returns metastore ID if the parent is a
	// metastore.
	ParentFullName string `json:"parent_full_name,omitempty"`
	// The quota parent securable type.
	ParentSecurableType SecurableType `json:"parent_securable_type,omitempty"`
	// The current usage of the resource quota.
	QuotaCount int `json:"quota_count,omitempty"`
	// The current limit of the resource quota.
	QuotaLimit int `json:"quota_limit,omitempty"`
	// The name of the quota.
	QuotaName string `json:"quota_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func quotaInfoFromPb(pb *quotaInfoPb) (*QuotaInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QuotaInfo{}
	st.LastRefreshedAt = pb.LastRefreshedAt
	st.ParentFullName = pb.ParentFullName
	st.ParentSecurableType = pb.ParentSecurableType
	st.QuotaCount = pb.QuotaCount
	st.QuotaLimit = pb.QuotaLimit
	st.QuotaName = pb.QuotaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *quotaInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st quotaInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// R2 temporary credentials for API authentication. Read more at
// https://developers.cloudflare.com/r2/api/s3/tokens/.
type R2Credentials struct {
	// The access key ID that identifies the temporary credentials.
	// Wire name: 'access_key_id'
	AccessKeyId string
	// The secret access key associated with the access key.
	// Wire name: 'secret_access_key'
	SecretAccessKey string
	// The generated JWT that users must pass to use the temporary credentials.
	// Wire name: 'session_token'
	SessionToken string

	ForceSendFields []string `tf:"-"`
}

func r2CredentialsToPb(st *R2Credentials) (*r2CredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &r2CredentialsPb{}
	pb.AccessKeyId = st.AccessKeyId

	pb.SecretAccessKey = st.SecretAccessKey

	pb.SessionToken = st.SessionToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *R2Credentials) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &r2CredentialsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := r2CredentialsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st R2Credentials) MarshalJSON() ([]byte, error) {
	pb, err := r2CredentialsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type r2CredentialsPb struct {
	// The access key ID that identifies the temporary credentials.
	AccessKeyId string `json:"access_key_id,omitempty"`
	// The secret access key associated with the access key.
	SecretAccessKey string `json:"secret_access_key,omitempty"`
	// The generated JWT that users must pass to use the temporary credentials.
	SessionToken string `json:"session_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func r2CredentialsFromPb(pb *r2CredentialsPb) (*R2Credentials, error) {
	if pb == nil {
		return nil, nil
	}
	st := &R2Credentials{}
	st.AccessKeyId = pb.AccessKeyId
	st.SecretAccessKey = pb.SecretAccessKey
	st.SessionToken = pb.SessionToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *r2CredentialsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st r2CredentialsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get a Volume
type ReadVolumeRequest struct {
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// The three-level (fully qualified) name of the volume
	// Wire name: 'name'
	Name string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func readVolumeRequestToPb(st *ReadVolumeRequest) (*readVolumeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &readVolumeRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ReadVolumeRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &readVolumeRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := readVolumeRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ReadVolumeRequest) MarshalJSON() ([]byte, error) {
	pb, err := readVolumeRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type readVolumeRequestPb struct {
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" url:"include_browse,omitempty"`
	// The three-level (fully qualified) name of the volume
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func readVolumeRequestFromPb(pb *readVolumeRequestPb) (*ReadVolumeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReadVolumeRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *readVolumeRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st readVolumeRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegenerateDashboardRequest struct {
	// Full name of the table.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
	// Optional argument to specify the warehouse for dashboard regeneration. If
	// not specified, the first running warehouse will be used.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func regenerateDashboardRequestToPb(st *RegenerateDashboardRequest) (*regenerateDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &regenerateDashboardRequestPb{}
	pb.TableName = st.TableName

	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RegenerateDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &regenerateDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := regenerateDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegenerateDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := regenerateDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type regenerateDashboardRequestPb struct {
	// Full name of the table.
	TableName string `json:"-" url:"-"`
	// Optional argument to specify the warehouse for dashboard regeneration. If
	// not specified, the first running warehouse will be used.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func regenerateDashboardRequestFromPb(pb *regenerateDashboardRequestPb) (*RegenerateDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegenerateDashboardRequest{}
	st.TableName = pb.TableName
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *regenerateDashboardRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st regenerateDashboardRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegenerateDashboardResponse struct {
	// Id of the regenerated monitoring dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string
	// The directory where the regenerated dashboard is stored.
	// Wire name: 'parent_folder'
	ParentFolder string

	ForceSendFields []string `tf:"-"`
}

func regenerateDashboardResponseToPb(st *RegenerateDashboardResponse) (*regenerateDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &regenerateDashboardResponsePb{}
	pb.DashboardId = st.DashboardId

	pb.ParentFolder = st.ParentFolder

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RegenerateDashboardResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &regenerateDashboardResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := regenerateDashboardResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegenerateDashboardResponse) MarshalJSON() ([]byte, error) {
	pb, err := regenerateDashboardResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type regenerateDashboardResponsePb struct {
	// Id of the regenerated monitoring dashboard.
	DashboardId string `json:"dashboard_id,omitempty"`
	// The directory where the regenerated dashboard is stored.
	ParentFolder string `json:"parent_folder,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func regenerateDashboardResponseFromPb(pb *regenerateDashboardResponsePb) (*RegenerateDashboardResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegenerateDashboardResponse{}
	st.DashboardId = pb.DashboardId
	st.ParentFolder = pb.ParentFolder

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *regenerateDashboardResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st regenerateDashboardResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Registered model alias.
type RegisteredModelAlias struct {
	// Name of the alias, e.g. 'champion' or 'latest_stable'
	// Wire name: 'alias_name'
	AliasName string
	// Integer version number of the model version to which this alias points.
	// Wire name: 'version_num'
	VersionNum int

	ForceSendFields []string `tf:"-"`
}

func registeredModelAliasToPb(st *RegisteredModelAlias) (*registeredModelAliasPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelAliasPb{}
	pb.AliasName = st.AliasName

	pb.VersionNum = st.VersionNum

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RegisteredModelAlias) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registeredModelAliasPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registeredModelAliasFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegisteredModelAlias) MarshalJSON() ([]byte, error) {
	pb, err := registeredModelAliasToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type registeredModelAliasPb struct {
	// Name of the alias, e.g. 'champion' or 'latest_stable'
	AliasName string `json:"alias_name,omitempty"`
	// Integer version number of the model version to which this alias points.
	VersionNum int `json:"version_num,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registeredModelAliasFromPb(pb *registeredModelAliasPb) (*RegisteredModelAlias, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelAlias{}
	st.AliasName = pb.AliasName
	st.VersionNum = pb.VersionNum

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registeredModelAliasPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registeredModelAliasPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegisteredModelInfo struct {
	// List of aliases associated with the registered model
	// Wire name: 'aliases'
	Aliases []RegisteredModelAlias
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool
	// The name of the catalog where the schema and the registered model reside
	// Wire name: 'catalog_name'
	CatalogName string
	// The comment attached to the registered model
	// Wire name: 'comment'
	Comment string
	// Creation timestamp of the registered model in milliseconds since the Unix
	// epoch
	// Wire name: 'created_at'
	CreatedAt int64
	// The identifier of the user who created the registered model
	// Wire name: 'created_by'
	CreatedBy string
	// The three-level (fully qualified) name of the registered model
	// Wire name: 'full_name'
	FullName string
	// The unique identifier of the metastore
	// Wire name: 'metastore_id'
	MetastoreId string
	// The name of the registered model
	// Wire name: 'name'
	Name string
	// The identifier of the user who owns the registered model
	// Wire name: 'owner'
	Owner string
	// The name of the schema where the registered model resides
	// Wire name: 'schema_name'
	SchemaName string
	// The storage location on the cloud under which model version data files
	// are stored
	// Wire name: 'storage_location'
	StorageLocation string
	// Last-update timestamp of the registered model in milliseconds since the
	// Unix epoch
	// Wire name: 'updated_at'
	UpdatedAt int64
	// The identifier of the user who updated the registered model last time
	// Wire name: 'updated_by'
	UpdatedBy string

	ForceSendFields []string `tf:"-"`
}

func registeredModelInfoToPb(st *RegisteredModelInfo) (*registeredModelInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelInfoPb{}

	var aliasesPb []registeredModelAliasPb
	for _, item := range st.Aliases {
		itemPb, err := registeredModelAliasToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			aliasesPb = append(aliasesPb, *itemPb)
		}
	}
	pb.Aliases = aliasesPb

	pb.BrowseOnly = st.BrowseOnly

	pb.CatalogName = st.CatalogName

	pb.Comment = st.Comment

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	pb.FullName = st.FullName

	pb.MetastoreId = st.MetastoreId

	pb.Name = st.Name

	pb.Owner = st.Owner

	pb.SchemaName = st.SchemaName

	pb.StorageLocation = st.StorageLocation

	pb.UpdatedAt = st.UpdatedAt

	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RegisteredModelInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registeredModelInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registeredModelInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegisteredModelInfo) MarshalJSON() ([]byte, error) {
	pb, err := registeredModelInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type registeredModelInfoPb struct {
	// List of aliases associated with the registered model
	Aliases []registeredModelAliasPb `json:"aliases,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The name of the catalog where the schema and the registered model reside
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the registered model
	Comment string `json:"comment,omitempty"`
	// Creation timestamp of the registered model in milliseconds since the Unix
	// epoch
	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the registered model
	CreatedBy string `json:"created_by,omitempty"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the metastore
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the registered model
	Name string `json:"name,omitempty"`
	// The identifier of the user who owns the registered model
	Owner string `json:"owner,omitempty"`
	// The name of the schema where the registered model resides
	SchemaName string `json:"schema_name,omitempty"`
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string `json:"storage_location,omitempty"`
	// Last-update timestamp of the registered model in milliseconds since the
	// Unix epoch
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the registered model last time
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registeredModelInfoFromPb(pb *registeredModelInfoPb) (*RegisteredModelInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelInfo{}

	var aliasesField []RegisteredModelAlias
	for _, itemPb := range pb.Aliases {
		item, err := registeredModelAliasFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			aliasesField = append(aliasesField, *item)
		}
	}
	st.Aliases = aliasesField
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.FullName = pb.FullName
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.SchemaName = pb.SchemaName
	st.StorageLocation = pb.StorageLocation
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registeredModelInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registeredModelInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Queue a metric refresh for a monitor
type RunRefreshRequest struct {
	// Full name of the table.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func runRefreshRequestToPb(st *RunRefreshRequest) (*runRefreshRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runRefreshRequestPb{}
	pb.TableName = st.TableName

	return pb, nil
}

func (st *RunRefreshRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runRefreshRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runRefreshRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunRefreshRequest) MarshalJSON() ([]byte, error) {
	pb, err := runRefreshRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runRefreshRequestPb struct {
	// Full name of the table.
	TableName string `json:"-" url:"-"`
}

func runRefreshRequestFromPb(pb *runRefreshRequestPb) (*RunRefreshRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunRefreshRequest{}
	st.TableName = pb.TableName

	return st, nil
}

type SchemaInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool
	// Name of parent catalog.
	// Wire name: 'catalog_name'
	CatalogName string
	// The type of the parent catalog.
	// Wire name: 'catalog_type'
	CatalogType string
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// Time at which this schema was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of schema creator.
	// Wire name: 'created_by'
	CreatedBy string

	// Wire name: 'effective_predictive_optimization_flag'
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'enable_predictive_optimization'
	EnablePredictiveOptimization EnablePredictiveOptimization
	// Full name of schema, in form of __catalog_name__.__schema_name__.
	// Wire name: 'full_name'
	FullName string
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string
	// Name of schema, relative to parent catalog.
	// Wire name: 'name'
	Name string
	// Username of current owner of schema.
	// Wire name: 'owner'
	Owner string
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string
	// The unique identifier of the schema.
	// Wire name: 'schema_id'
	SchemaId string
	// Storage location for managed tables within schema.
	// Wire name: 'storage_location'
	StorageLocation string
	// Storage root URL for managed tables within schema.
	// Wire name: 'storage_root'
	StorageRoot string
	// Time at which this schema was created, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// Username of user who last modified schema.
	// Wire name: 'updated_by'
	UpdatedBy string

	ForceSendFields []string `tf:"-"`
}

func schemaInfoToPb(st *SchemaInfo) (*schemaInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &schemaInfoPb{}
	pb.BrowseOnly = st.BrowseOnly

	pb.CatalogName = st.CatalogName

	pb.CatalogType = st.CatalogType

	pb.Comment = st.Comment

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	effectivePredictiveOptimizationFlagPb, err := effectivePredictiveOptimizationFlagToPb(st.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagPb != nil {
		pb.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagPb
	}

	pb.EnablePredictiveOptimization = st.EnablePredictiveOptimization

	pb.FullName = st.FullName

	pb.MetastoreId = st.MetastoreId

	pb.Name = st.Name

	pb.Owner = st.Owner

	pb.Properties = st.Properties

	pb.SchemaId = st.SchemaId

	pb.StorageLocation = st.StorageLocation

	pb.StorageRoot = st.StorageRoot

	pb.UpdatedAt = st.UpdatedAt

	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SchemaInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &schemaInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := schemaInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SchemaInfo) MarshalJSON() ([]byte, error) {
	pb, err := schemaInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type schemaInfoPb struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// Name of parent catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// The type of the parent catalog.
	CatalogType string `json:"catalog_type,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this schema was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of schema creator.
	CreatedBy string `json:"created_by,omitempty"`

	EffectivePredictiveOptimizationFlag *effectivePredictiveOptimizationFlagPb `json:"effective_predictive_optimization_flag,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// Full name of schema, in form of __catalog_name__.__schema_name__.
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of schema, relative to parent catalog.
	Name string `json:"name,omitempty"`
	// Username of current owner of schema.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`
	// The unique identifier of the schema.
	SchemaId string `json:"schema_id,omitempty"`
	// Storage location for managed tables within schema.
	StorageLocation string `json:"storage_location,omitempty"`
	// Storage root URL for managed tables within schema.
	StorageRoot string `json:"storage_root,omitempty"`
	// Time at which this schema was created, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified schema.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func schemaInfoFromPb(pb *schemaInfoPb) (*SchemaInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SchemaInfo{}
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogName = pb.CatalogName
	st.CatalogType = pb.CatalogType
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	effectivePredictiveOptimizationFlagField, err := effectivePredictiveOptimizationFlagFromPb(pb.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagField != nil {
		st.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagField
	}
	st.EnablePredictiveOptimization = pb.EnablePredictiveOptimization
	st.FullName = pb.FullName
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.Properties = pb.Properties
	st.SchemaId = pb.SchemaId
	st.StorageLocation = pb.StorageLocation
	st.StorageRoot = pb.StorageRoot
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *schemaInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st schemaInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// A map of key-value properties attached to the securable.

type SecurableOptionsMap map[string]string
type securableOptionsMapPb SecurableOptionsMap

func securableOptionsMapToPb(st *SecurableOptionsMap) (*securableOptionsMapPb, error) {
	if st == nil {
		return nil, nil
	}
	stPb := securableOptionsMapPb(*st)
	return &stPb, nil
}
func securableOptionsMapFromPb(stPb *securableOptionsMapPb) (*SecurableOptionsMap, error) {
	if stPb == nil {
		return nil, nil
	}
	st := SecurableOptionsMap(*stPb)
	return &st, nil
}

// A map of key-value properties attached to the securable.

type SecurablePropertiesMap map[string]string
type securablePropertiesMapPb SecurablePropertiesMap

func securablePropertiesMapToPb(st *SecurablePropertiesMap) (*securablePropertiesMapPb, error) {
	if st == nil {
		return nil, nil
	}
	stPb := securablePropertiesMapPb(*st)
	return &stPb, nil
}
func securablePropertiesMapFromPb(stPb *securablePropertiesMapPb) (*SecurablePropertiesMap, error) {
	if stPb == nil {
		return nil, nil
	}
	st := SecurablePropertiesMap(*stPb)
	return &st, nil
}

// The type of Unity Catalog securable
type SecurableType string
type securableTypePb string

const SecurableTypeCatalog SecurableType = `CATALOG`

const SecurableTypeCleanRoom SecurableType = `CLEAN_ROOM`

const SecurableTypeConnection SecurableType = `CONNECTION`

const SecurableTypeCredential SecurableType = `CREDENTIAL`

const SecurableTypeExternalLocation SecurableType = `EXTERNAL_LOCATION`

const SecurableTypeFunction SecurableType = `FUNCTION`

const SecurableTypeMetastore SecurableType = `METASTORE`

const SecurableTypePipeline SecurableType = `PIPELINE`

const SecurableTypeProvider SecurableType = `PROVIDER`

const SecurableTypeRecipient SecurableType = `RECIPIENT`

const SecurableTypeSchema SecurableType = `SCHEMA`

const SecurableTypeShare SecurableType = `SHARE`

const SecurableTypeStorageCredential SecurableType = `STORAGE_CREDENTIAL`

const SecurableTypeTable SecurableType = `TABLE`

const SecurableTypeVolume SecurableType = `VOLUME`

// String representation for [fmt.Print]
func (f *SecurableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SecurableType) Set(v string) error {
	switch v {
	case `CATALOG`, `CLEAN_ROOM`, `CONNECTION`, `CREDENTIAL`, `EXTERNAL_LOCATION`, `FUNCTION`, `METASTORE`, `PIPELINE`, `PROVIDER`, `RECIPIENT`, `SCHEMA`, `SHARE`, `STORAGE_CREDENTIAL`, `TABLE`, `VOLUME`:
		*f = SecurableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CATALOG", "CLEAN_ROOM", "CONNECTION", "CREDENTIAL", "EXTERNAL_LOCATION", "FUNCTION", "METASTORE", "PIPELINE", "PROVIDER", "RECIPIENT", "SCHEMA", "SHARE", "STORAGE_CREDENTIAL", "TABLE", "VOLUME"`, v)
	}
}

// Type always returns SecurableType to satisfy [pflag.Value] interface
func (f *SecurableType) Type() string {
	return "SecurableType"
}

func securableTypeToPb(st *SecurableType) (*securableTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := securableTypePb(*st)
	return &pb, nil
}

func securableTypeFromPb(pb *securableTypePb) (*SecurableType, error) {
	if pb == nil {
		return nil, nil
	}
	st := SecurableType(*pb)
	return &st, nil
}

type SetArtifactAllowlist struct {
	// A list of allowed artifact match patterns.
	// Wire name: 'artifact_matchers'
	ArtifactMatchers []ArtifactMatcher
	// The artifact type of the allowlist.
	// Wire name: 'artifact_type'
	ArtifactType ArtifactType `tf:"-"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of the user who set the artifact allowlist.
	// Wire name: 'created_by'
	CreatedBy string
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string

	ForceSendFields []string `tf:"-"`
}

func setArtifactAllowlistToPb(st *SetArtifactAllowlist) (*setArtifactAllowlistPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setArtifactAllowlistPb{}

	var artifactMatchersPb []artifactMatcherPb
	for _, item := range st.ArtifactMatchers {
		itemPb, err := artifactMatcherToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			artifactMatchersPb = append(artifactMatchersPb, *itemPb)
		}
	}
	pb.ArtifactMatchers = artifactMatchersPb

	pb.ArtifactType = st.ArtifactType

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	pb.MetastoreId = st.MetastoreId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SetArtifactAllowlist) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setArtifactAllowlistPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setArtifactAllowlistFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetArtifactAllowlist) MarshalJSON() ([]byte, error) {
	pb, err := setArtifactAllowlistToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type setArtifactAllowlistPb struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers []artifactMatcherPb `json:"artifact_matchers"`
	// The artifact type of the allowlist.
	ArtifactType ArtifactType `json:"-" url:"-"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of the user who set the artifact allowlist.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func setArtifactAllowlistFromPb(pb *setArtifactAllowlistPb) (*SetArtifactAllowlist, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetArtifactAllowlist{}

	var artifactMatchersField []ArtifactMatcher
	for _, itemPb := range pb.ArtifactMatchers {
		item, err := artifactMatcherFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			artifactMatchersField = append(artifactMatchersField, *item)
		}
	}
	st.ArtifactMatchers = artifactMatchersField
	st.ArtifactType = pb.ArtifactType
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.MetastoreId = pb.MetastoreId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *setArtifactAllowlistPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st setArtifactAllowlistPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SetRegisteredModelAliasRequest struct {
	// The name of the alias
	// Wire name: 'alias'
	Alias string
	// Full name of the registered model
	// Wire name: 'full_name'
	FullName string
	// The version number of the model version to which the alias points
	// Wire name: 'version_num'
	VersionNum int
}

func setRegisteredModelAliasRequestToPb(st *SetRegisteredModelAliasRequest) (*setRegisteredModelAliasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setRegisteredModelAliasRequestPb{}
	pb.Alias = st.Alias

	pb.FullName = st.FullName

	pb.VersionNum = st.VersionNum

	return pb, nil
}

func (st *SetRegisteredModelAliasRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setRegisteredModelAliasRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setRegisteredModelAliasRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetRegisteredModelAliasRequest) MarshalJSON() ([]byte, error) {
	pb, err := setRegisteredModelAliasRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type setRegisteredModelAliasRequestPb struct {
	// The name of the alias
	Alias string `json:"alias" url:"-"`
	// Full name of the registered model
	FullName string `json:"full_name" url:"-"`
	// The version number of the model version to which the alias points
	VersionNum int `json:"version_num"`
}

func setRegisteredModelAliasRequestFromPb(pb *setRegisteredModelAliasRequestPb) (*SetRegisteredModelAliasRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetRegisteredModelAliasRequest{}
	st.Alias = pb.Alias
	st.FullName = pb.FullName
	st.VersionNum = pb.VersionNum

	return st, nil
}

// Server-Side Encryption properties for clients communicating with AWS s3.
type SseEncryptionDetails struct {
	// The type of key encryption to use (affects headers from s3 client).
	// Wire name: 'algorithm'
	Algorithm SseEncryptionDetailsAlgorithm
	// When algorithm is **AWS_SSE_KMS** this field specifies the ARN of the SSE
	// key to use.
	// Wire name: 'aws_kms_key_arn'
	AwsKmsKeyArn string

	ForceSendFields []string `tf:"-"`
}

func sseEncryptionDetailsToPb(st *SseEncryptionDetails) (*sseEncryptionDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sseEncryptionDetailsPb{}
	pb.Algorithm = st.Algorithm

	pb.AwsKmsKeyArn = st.AwsKmsKeyArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SseEncryptionDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sseEncryptionDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sseEncryptionDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SseEncryptionDetails) MarshalJSON() ([]byte, error) {
	pb, err := sseEncryptionDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sseEncryptionDetailsPb struct {
	// The type of key encryption to use (affects headers from s3 client).
	Algorithm SseEncryptionDetailsAlgorithm `json:"algorithm,omitempty"`
	// When algorithm is **AWS_SSE_KMS** this field specifies the ARN of the SSE
	// key to use.
	AwsKmsKeyArn string `json:"aws_kms_key_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sseEncryptionDetailsFromPb(pb *sseEncryptionDetailsPb) (*SseEncryptionDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SseEncryptionDetails{}
	st.Algorithm = pb.Algorithm
	st.AwsKmsKeyArn = pb.AwsKmsKeyArn

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sseEncryptionDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sseEncryptionDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The type of key encryption to use (affects headers from s3 client).
type SseEncryptionDetailsAlgorithm string
type sseEncryptionDetailsAlgorithmPb string

const SseEncryptionDetailsAlgorithmAwsSseKms SseEncryptionDetailsAlgorithm = `AWS_SSE_KMS`

const SseEncryptionDetailsAlgorithmAwsSseS3 SseEncryptionDetailsAlgorithm = `AWS_SSE_S3`

// String representation for [fmt.Print]
func (f *SseEncryptionDetailsAlgorithm) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SseEncryptionDetailsAlgorithm) Set(v string) error {
	switch v {
	case `AWS_SSE_KMS`, `AWS_SSE_S3`:
		*f = SseEncryptionDetailsAlgorithm(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AWS_SSE_KMS", "AWS_SSE_S3"`, v)
	}
}

// Type always returns SseEncryptionDetailsAlgorithm to satisfy [pflag.Value] interface
func (f *SseEncryptionDetailsAlgorithm) Type() string {
	return "SseEncryptionDetailsAlgorithm"
}

func sseEncryptionDetailsAlgorithmToPb(st *SseEncryptionDetailsAlgorithm) (*sseEncryptionDetailsAlgorithmPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sseEncryptionDetailsAlgorithmPb(*st)
	return &pb, nil
}

func sseEncryptionDetailsAlgorithmFromPb(pb *sseEncryptionDetailsAlgorithmPb) (*SseEncryptionDetailsAlgorithm, error) {
	if pb == nil {
		return nil, nil
	}
	st := SseEncryptionDetailsAlgorithm(*pb)
	return &st, nil
}

type StorageCredentialInfo struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRoleResponse
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentityResponse
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal
	// The Cloudflare API token configuration.
	// Wire name: 'cloudflare_api_token'
	CloudflareApiToken *CloudflareApiToken
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string
	// Time at which this Credential was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of credential creator.
	// Wire name: 'created_by'
	CreatedBy string
	// The Databricks managed GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountResponse
	// The full name of the credential.
	// Wire name: 'full_name'
	FullName string
	// The unique identifier of the credential.
	// Wire name: 'id'
	Id string

	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string
	// The credential name. The name must be unique within the metastore.
	// Wire name: 'name'
	Name string
	// Username of current owner of credential.
	// Wire name: 'owner'
	Owner string
	// Whether the storage credential is only usable for read operations.
	// Wire name: 'read_only'
	ReadOnly bool
	// Time at which this credential was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// Username of user who last modified the credential.
	// Wire name: 'updated_by'
	UpdatedBy string
	// Whether this credential is the current metastore's root storage
	// credential.
	// Wire name: 'used_for_managed_storage'
	UsedForManagedStorage bool

	ForceSendFields []string `tf:"-"`
}

func storageCredentialInfoToPb(st *StorageCredentialInfo) (*storageCredentialInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &storageCredentialInfoPb{}
	awsIamRolePb, err := awsIamRoleResponseToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}

	azureManagedIdentityPb, err := azureManagedIdentityResponseToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}

	azureServicePrincipalPb, err := azureServicePrincipalToPb(st.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalPb != nil {
		pb.AzureServicePrincipal = azureServicePrincipalPb
	}

	cloudflareApiTokenPb, err := cloudflareApiTokenToPb(st.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenPb != nil {
		pb.CloudflareApiToken = cloudflareApiTokenPb
	}

	pb.Comment = st.Comment

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	databricksGcpServiceAccountPb, err := databricksGcpServiceAccountResponseToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}

	pb.FullName = st.FullName

	pb.Id = st.Id

	pb.IsolationMode = st.IsolationMode

	pb.MetastoreId = st.MetastoreId

	pb.Name = st.Name

	pb.Owner = st.Owner

	pb.ReadOnly = st.ReadOnly

	pb.UpdatedAt = st.UpdatedAt

	pb.UpdatedBy = st.UpdatedBy

	pb.UsedForManagedStorage = st.UsedForManagedStorage

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *StorageCredentialInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &storageCredentialInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := storageCredentialInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StorageCredentialInfo) MarshalJSON() ([]byte, error) {
	pb, err := storageCredentialInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type storageCredentialInfoPb struct {
	// The AWS IAM role configuration.
	AwsIamRole *awsIamRoleResponsePb `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *azureManagedIdentityResponsePb `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *azureServicePrincipalPb `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *cloudflareApiTokenPb `json:"cloudflare_api_token,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// Time at which this Credential was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of credential creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *databricksGcpServiceAccountResponsePb `json:"databricks_gcp_service_account,omitempty"`
	// The full name of the credential.
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the credential.
	Id string `json:"id,omitempty"`

	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The credential name. The name must be unique within the metastore.
	Name string `json:"name,omitempty"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `json:"read_only,omitempty"`
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the credential.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Whether this credential is the current metastore's root storage
	// credential.
	UsedForManagedStorage bool `json:"used_for_managed_storage,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func storageCredentialInfoFromPb(pb *storageCredentialInfoPb) (*StorageCredentialInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StorageCredentialInfo{}
	awsIamRoleField, err := awsIamRoleResponseFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := azureManagedIdentityResponseFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	azureServicePrincipalField, err := azureServicePrincipalFromPb(pb.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalField != nil {
		st.AzureServicePrincipal = azureServicePrincipalField
	}
	cloudflareApiTokenField, err := cloudflareApiTokenFromPb(pb.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenField != nil {
		st.CloudflareApiToken = cloudflareApiTokenField
	}
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	databricksGcpServiceAccountField, err := databricksGcpServiceAccountResponseFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	st.FullName = pb.FullName
	st.Id = pb.Id
	st.IsolationMode = pb.IsolationMode
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.ReadOnly = pb.ReadOnly
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.UsedForManagedStorage = pb.UsedForManagedStorage

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *storageCredentialInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st storageCredentialInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SystemSchemaInfo struct {
	// Name of the system schema.
	// Wire name: 'schema'
	Schema string
	// The current state of enablement for the system schema. An empty string
	// means the system schema is available and ready for opt-in.
	// Wire name: 'state'
	State SystemSchemaInfoState

	ForceSendFields []string `tf:"-"`
}

func systemSchemaInfoToPb(st *SystemSchemaInfo) (*systemSchemaInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &systemSchemaInfoPb{}
	pb.Schema = st.Schema

	pb.State = st.State

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SystemSchemaInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &systemSchemaInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := systemSchemaInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SystemSchemaInfo) MarshalJSON() ([]byte, error) {
	pb, err := systemSchemaInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type systemSchemaInfoPb struct {
	// Name of the system schema.
	Schema string `json:"schema,omitempty"`
	// The current state of enablement for the system schema. An empty string
	// means the system schema is available and ready for opt-in.
	State SystemSchemaInfoState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func systemSchemaInfoFromPb(pb *systemSchemaInfoPb) (*SystemSchemaInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SystemSchemaInfo{}
	st.Schema = pb.Schema
	st.State = pb.State

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *systemSchemaInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st systemSchemaInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The current state of enablement for the system schema. An empty string means
// the system schema is available and ready for opt-in.
type SystemSchemaInfoState string
type systemSchemaInfoStatePb string

const SystemSchemaInfoStateAvailable SystemSchemaInfoState = `AVAILABLE`

const SystemSchemaInfoStateDisableInitialized SystemSchemaInfoState = `DISABLE_INITIALIZED`

const SystemSchemaInfoStateEnableCompleted SystemSchemaInfoState = `ENABLE_COMPLETED`

const SystemSchemaInfoStateEnableInitialized SystemSchemaInfoState = `ENABLE_INITIALIZED`

const SystemSchemaInfoStateUnavailable SystemSchemaInfoState = `UNAVAILABLE`

// String representation for [fmt.Print]
func (f *SystemSchemaInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SystemSchemaInfoState) Set(v string) error {
	switch v {
	case `AVAILABLE`, `DISABLE_INITIALIZED`, `ENABLE_COMPLETED`, `ENABLE_INITIALIZED`, `UNAVAILABLE`:
		*f = SystemSchemaInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVAILABLE", "DISABLE_INITIALIZED", "ENABLE_COMPLETED", "ENABLE_INITIALIZED", "UNAVAILABLE"`, v)
	}
}

// Type always returns SystemSchemaInfoState to satisfy [pflag.Value] interface
func (f *SystemSchemaInfoState) Type() string {
	return "SystemSchemaInfoState"
}

func systemSchemaInfoStateToPb(st *SystemSchemaInfoState) (*systemSchemaInfoStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := systemSchemaInfoStatePb(*st)
	return &pb, nil
}

func systemSchemaInfoStateFromPb(pb *systemSchemaInfoStatePb) (*SystemSchemaInfoState, error) {
	if pb == nil {
		return nil, nil
	}
	st := SystemSchemaInfoState(*pb)
	return &st, nil
}

// A table constraint, as defined by *one* of the following fields being set:
// __primary_key_constraint__, __foreign_key_constraint__,
// __named_table_constraint__.
type TableConstraint struct {

	// Wire name: 'foreign_key_constraint'
	ForeignKeyConstraint *ForeignKeyConstraint

	// Wire name: 'named_table_constraint'
	NamedTableConstraint *NamedTableConstraint

	// Wire name: 'primary_key_constraint'
	PrimaryKeyConstraint *PrimaryKeyConstraint
}

func tableConstraintToPb(st *TableConstraint) (*tableConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableConstraintPb{}
	foreignKeyConstraintPb, err := foreignKeyConstraintToPb(st.ForeignKeyConstraint)
	if err != nil {
		return nil, err
	}
	if foreignKeyConstraintPb != nil {
		pb.ForeignKeyConstraint = foreignKeyConstraintPb
	}

	namedTableConstraintPb, err := namedTableConstraintToPb(st.NamedTableConstraint)
	if err != nil {
		return nil, err
	}
	if namedTableConstraintPb != nil {
		pb.NamedTableConstraint = namedTableConstraintPb
	}

	primaryKeyConstraintPb, err := primaryKeyConstraintToPb(st.PrimaryKeyConstraint)
	if err != nil {
		return nil, err
	}
	if primaryKeyConstraintPb != nil {
		pb.PrimaryKeyConstraint = primaryKeyConstraintPb
	}

	return pb, nil
}

func (st *TableConstraint) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tableConstraintPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tableConstraintFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TableConstraint) MarshalJSON() ([]byte, error) {
	pb, err := tableConstraintToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type tableConstraintPb struct {
	ForeignKeyConstraint *foreignKeyConstraintPb `json:"foreign_key_constraint,omitempty"`

	NamedTableConstraint *namedTableConstraintPb `json:"named_table_constraint,omitempty"`

	PrimaryKeyConstraint *primaryKeyConstraintPb `json:"primary_key_constraint,omitempty"`
}

func tableConstraintFromPb(pb *tableConstraintPb) (*TableConstraint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableConstraint{}
	foreignKeyConstraintField, err := foreignKeyConstraintFromPb(pb.ForeignKeyConstraint)
	if err != nil {
		return nil, err
	}
	if foreignKeyConstraintField != nil {
		st.ForeignKeyConstraint = foreignKeyConstraintField
	}
	namedTableConstraintField, err := namedTableConstraintFromPb(pb.NamedTableConstraint)
	if err != nil {
		return nil, err
	}
	if namedTableConstraintField != nil {
		st.NamedTableConstraint = namedTableConstraintField
	}
	primaryKeyConstraintField, err := primaryKeyConstraintFromPb(pb.PrimaryKeyConstraint)
	if err != nil {
		return nil, err
	}
	if primaryKeyConstraintField != nil {
		st.PrimaryKeyConstraint = primaryKeyConstraintField
	}

	return st, nil
}

// A table that is dependent on a SQL object.
type TableDependency struct {
	// Full name of the dependent table, in the form of
	// __catalog_name__.__schema_name__.__table_name__.
	// Wire name: 'table_full_name'
	TableFullName string
}

func tableDependencyToPb(st *TableDependency) (*tableDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableDependencyPb{}
	pb.TableFullName = st.TableFullName

	return pb, nil
}

func (st *TableDependency) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tableDependencyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tableDependencyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TableDependency) MarshalJSON() ([]byte, error) {
	pb, err := tableDependencyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type tableDependencyPb struct {
	// Full name of the dependent table, in the form of
	// __catalog_name__.__schema_name__.__table_name__.
	TableFullName string `json:"table_full_name"`
}

func tableDependencyFromPb(pb *tableDependencyPb) (*TableDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableDependency{}
	st.TableFullName = pb.TableFullName

	return st, nil
}

type TableExistsResponse struct {
	// Whether the table exists or not.
	// Wire name: 'table_exists'
	TableExists bool

	ForceSendFields []string `tf:"-"`
}

func tableExistsResponseToPb(st *TableExistsResponse) (*tableExistsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableExistsResponsePb{}
	pb.TableExists = st.TableExists

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *TableExistsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tableExistsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tableExistsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TableExistsResponse) MarshalJSON() ([]byte, error) {
	pb, err := tableExistsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type tableExistsResponsePb struct {
	// Whether the table exists or not.
	TableExists bool `json:"table_exists,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tableExistsResponseFromPb(pb *tableExistsResponsePb) (*TableExistsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableExistsResponse{}
	st.TableExists = pb.TableExists

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tableExistsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tableExistsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TableInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	// Wire name: 'access_point'
	AccessPoint string
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool
	// Name of parent catalog.
	// Wire name: 'catalog_name'
	CatalogName string
	// The array of __ColumnInfo__ definitions of the table's columns.
	// Wire name: 'columns'
	Columns []ColumnInfo
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// Time at which this table was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of table creator.
	// Wire name: 'created_by'
	CreatedBy string
	// Unique ID of the Data Access Configuration to use with the table data.
	// Wire name: 'data_access_configuration_id'
	DataAccessConfigurationId string
	// Data source format
	// Wire name: 'data_source_format'
	DataSourceFormat DataSourceFormat
	// Time at which this table was deleted, in epoch milliseconds. Field is
	// omitted if table is not deleted.
	// Wire name: 'deleted_at'
	DeletedAt int64
	// Information pertaining to current state of the delta table.
	// Wire name: 'delta_runtime_properties_kvpairs'
	DeltaRuntimePropertiesKvpairs *DeltaRuntimePropertiesKvPairs

	// Wire name: 'effective_predictive_optimization_flag'
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'enable_predictive_optimization'
	EnablePredictiveOptimization EnablePredictiveOptimization
	// Encryption options that apply to clients connecting to cloud storage.
	// Wire name: 'encryption_details'
	EncryptionDetails *EncryptionDetails
	// Full name of table, in form of
	// __catalog_name__.__schema_name__.__table_name__
	// Wire name: 'full_name'
	FullName string
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string
	// Name of table, relative to parent schema.
	// Wire name: 'name'
	Name string
	// Username of current owner of table.
	// Wire name: 'owner'
	Owner string
	// The pipeline ID of the table. Applicable for tables created by pipelines
	// (Materialized View, Streaming Table, etc.).
	// Wire name: 'pipeline_id'
	PipelineId string
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string

	// Wire name: 'row_filter'
	RowFilter *TableRowFilter
	// Name of parent schema relative to its parent catalog.
	// Wire name: 'schema_name'
	SchemaName string
	// List of schemes whose objects can be referenced without qualification.
	// Wire name: 'sql_path'
	SqlPath string
	// Name of the storage credential, when a storage credential is configured
	// for use with this table.
	// Wire name: 'storage_credential_name'
	StorageCredentialName string
	// Storage root URL for table (for **MANAGED**, **EXTERNAL** tables)
	// Wire name: 'storage_location'
	StorageLocation string
	// List of table constraints. Note: this field is not set in the output of
	// the __listTables__ API.
	// Wire name: 'table_constraints'
	TableConstraints []TableConstraint
	// The unique identifier of the table.
	// Wire name: 'table_id'
	TableId string

	// Wire name: 'table_type'
	TableType TableType
	// Time at which this table was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// Username of user who last modified the table.
	// Wire name: 'updated_by'
	UpdatedBy string
	// View definition SQL (when __table_type__ is **VIEW**,
	// **MATERIALIZED_VIEW**, or **STREAMING_TABLE**)
	// Wire name: 'view_definition'
	ViewDefinition string
	// View dependencies (when table_type == **VIEW** or **MATERIALIZED_VIEW**,
	// **STREAMING_TABLE**) - when DependencyList is None, the dependency is not
	// provided; - when DependencyList is an empty list, the dependency is
	// provided but is empty; - when DependencyList is not an empty list,
	// dependencies are provided and recorded.
	// Wire name: 'view_dependencies'
	ViewDependencies *DependencyList

	ForceSendFields []string `tf:"-"`
}

func tableInfoToPb(st *TableInfo) (*tableInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableInfoPb{}
	pb.AccessPoint = st.AccessPoint

	pb.BrowseOnly = st.BrowseOnly

	pb.CatalogName = st.CatalogName

	var columnsPb []ColumnInfoPb
	for _, item := range st.Columns {
		itemPb, err := ColumnInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			columnsPb = append(columnsPb, *itemPb)
		}
	}
	pb.Columns = columnsPb

	pb.Comment = st.Comment

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	pb.DataAccessConfigurationId = st.DataAccessConfigurationId

	pb.DataSourceFormat = st.DataSourceFormat

	pb.DeletedAt = st.DeletedAt

	deltaRuntimePropertiesKvpairsPb, err := deltaRuntimePropertiesKvPairsToPb(st.DeltaRuntimePropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if deltaRuntimePropertiesKvpairsPb != nil {
		pb.DeltaRuntimePropertiesKvpairs = deltaRuntimePropertiesKvpairsPb
	}

	effectivePredictiveOptimizationFlagPb, err := effectivePredictiveOptimizationFlagToPb(st.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagPb != nil {
		pb.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagPb
	}

	pb.EnablePredictiveOptimization = st.EnablePredictiveOptimization

	encryptionDetailsPb, err := encryptionDetailsToPb(st.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsPb != nil {
		pb.EncryptionDetails = encryptionDetailsPb
	}

	pb.FullName = st.FullName

	pb.MetastoreId = st.MetastoreId

	pb.Name = st.Name

	pb.Owner = st.Owner

	pb.PipelineId = st.PipelineId

	pb.Properties = st.Properties

	rowFilterPb, err := tableRowFilterToPb(st.RowFilter)
	if err != nil {
		return nil, err
	}
	if rowFilterPb != nil {
		pb.RowFilter = rowFilterPb
	}

	pb.SchemaName = st.SchemaName

	pb.SqlPath = st.SqlPath

	pb.StorageCredentialName = st.StorageCredentialName

	pb.StorageLocation = st.StorageLocation

	var tableConstraintsPb []tableConstraintPb
	for _, item := range st.TableConstraints {
		itemPb, err := tableConstraintToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tableConstraintsPb = append(tableConstraintsPb, *itemPb)
		}
	}
	pb.TableConstraints = tableConstraintsPb

	pb.TableId = st.TableId

	pb.TableType = st.TableType

	pb.UpdatedAt = st.UpdatedAt

	pb.UpdatedBy = st.UpdatedBy

	pb.ViewDefinition = st.ViewDefinition

	viewDependenciesPb, err := dependencyListToPb(st.ViewDependencies)
	if err != nil {
		return nil, err
	}
	if viewDependenciesPb != nil {
		pb.ViewDependencies = viewDependenciesPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *TableInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tableInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tableInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TableInfo) MarshalJSON() ([]byte, error) {
	pb, err := tableInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type tableInfoPb struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string `json:"access_point,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// Name of parent catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// The array of __ColumnInfo__ definitions of the table's columns.
	Columns []ColumnInfoPb `json:"columns,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this table was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of table creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Unique ID of the Data Access Configuration to use with the table data.
	DataAccessConfigurationId string `json:"data_access_configuration_id,omitempty"`
	// Data source format
	DataSourceFormat DataSourceFormat `json:"data_source_format,omitempty"`
	// Time at which this table was deleted, in epoch milliseconds. Field is
	// omitted if table is not deleted.
	DeletedAt int64 `json:"deleted_at,omitempty"`
	// Information pertaining to current state of the delta table.
	DeltaRuntimePropertiesKvpairs *deltaRuntimePropertiesKvPairsPb `json:"delta_runtime_properties_kvpairs,omitempty"`

	EffectivePredictiveOptimizationFlag *effectivePredictiveOptimizationFlagPb `json:"effective_predictive_optimization_flag,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *encryptionDetailsPb `json:"encryption_details,omitempty"`
	// Full name of table, in form of
	// __catalog_name__.__schema_name__.__table_name__
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of table, relative to parent schema.
	Name string `json:"name,omitempty"`
	// Username of current owner of table.
	Owner string `json:"owner,omitempty"`
	// The pipeline ID of the table. Applicable for tables created by pipelines
	// (Materialized View, Streaming Table, etc.).
	PipelineId string `json:"pipeline_id,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`

	RowFilter *tableRowFilterPb `json:"row_filter,omitempty"`
	// Name of parent schema relative to its parent catalog.
	SchemaName string `json:"schema_name,omitempty"`
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string `json:"sql_path,omitempty"`
	// Name of the storage credential, when a storage credential is configured
	// for use with this table.
	StorageCredentialName string `json:"storage_credential_name,omitempty"`
	// Storage root URL for table (for **MANAGED**, **EXTERNAL** tables)
	StorageLocation string `json:"storage_location,omitempty"`
	// List of table constraints. Note: this field is not set in the output of
	// the __listTables__ API.
	TableConstraints []tableConstraintPb `json:"table_constraints,omitempty"`
	// The unique identifier of the table.
	TableId string `json:"table_id,omitempty"`

	TableType TableType `json:"table_type,omitempty"`
	// Time at which this table was last modified, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the table.
	UpdatedBy string `json:"updated_by,omitempty"`
	// View definition SQL (when __table_type__ is **VIEW**,
	// **MATERIALIZED_VIEW**, or **STREAMING_TABLE**)
	ViewDefinition string `json:"view_definition,omitempty"`
	// View dependencies (when table_type == **VIEW** or **MATERIALIZED_VIEW**,
	// **STREAMING_TABLE**) - when DependencyList is None, the dependency is not
	// provided; - when DependencyList is an empty list, the dependency is
	// provided but is empty; - when DependencyList is not an empty list,
	// dependencies are provided and recorded.
	ViewDependencies *dependencyListPb `json:"view_dependencies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tableInfoFromPb(pb *tableInfoPb) (*TableInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableInfo{}
	st.AccessPoint = pb.AccessPoint
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogName = pb.CatalogName

	var columnsField []ColumnInfo
	for _, itemPb := range pb.Columns {
		item, err := ColumnInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			columnsField = append(columnsField, *item)
		}
	}
	st.Columns = columnsField
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.DataAccessConfigurationId = pb.DataAccessConfigurationId
	st.DataSourceFormat = pb.DataSourceFormat
	st.DeletedAt = pb.DeletedAt
	deltaRuntimePropertiesKvpairsField, err := deltaRuntimePropertiesKvPairsFromPb(pb.DeltaRuntimePropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if deltaRuntimePropertiesKvpairsField != nil {
		st.DeltaRuntimePropertiesKvpairs = deltaRuntimePropertiesKvpairsField
	}
	effectivePredictiveOptimizationFlagField, err := effectivePredictiveOptimizationFlagFromPb(pb.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagField != nil {
		st.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagField
	}
	st.EnablePredictiveOptimization = pb.EnablePredictiveOptimization
	encryptionDetailsField, err := encryptionDetailsFromPb(pb.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsField != nil {
		st.EncryptionDetails = encryptionDetailsField
	}
	st.FullName = pb.FullName
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.PipelineId = pb.PipelineId
	st.Properties = pb.Properties
	rowFilterField, err := tableRowFilterFromPb(pb.RowFilter)
	if err != nil {
		return nil, err
	}
	if rowFilterField != nil {
		st.RowFilter = rowFilterField
	}
	st.SchemaName = pb.SchemaName
	st.SqlPath = pb.SqlPath
	st.StorageCredentialName = pb.StorageCredentialName
	st.StorageLocation = pb.StorageLocation

	var tableConstraintsField []TableConstraint
	for _, itemPb := range pb.TableConstraints {
		item, err := tableConstraintFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tableConstraintsField = append(tableConstraintsField, *item)
		}
	}
	st.TableConstraints = tableConstraintsField
	st.TableId = pb.TableId
	st.TableType = pb.TableType
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.ViewDefinition = pb.ViewDefinition
	viewDependenciesField, err := dependencyListFromPb(pb.ViewDependencies)
	if err != nil {
		return nil, err
	}
	if viewDependenciesField != nil {
		st.ViewDependencies = viewDependenciesField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tableInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tableInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TableOperation string
type tableOperationPb string

const TableOperationRead TableOperation = `READ`

const TableOperationReadWrite TableOperation = `READ_WRITE`

// String representation for [fmt.Print]
func (f *TableOperation) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TableOperation) Set(v string) error {
	switch v {
	case `READ`, `READ_WRITE`:
		*f = TableOperation(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "READ", "READ_WRITE"`, v)
	}
}

// Type always returns TableOperation to satisfy [pflag.Value] interface
func (f *TableOperation) Type() string {
	return "TableOperation"
}

func tableOperationToPb(st *TableOperation) (*tableOperationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := tableOperationPb(*st)
	return &pb, nil
}

func tableOperationFromPb(pb *tableOperationPb) (*TableOperation, error) {
	if pb == nil {
		return nil, nil
	}
	st := TableOperation(*pb)
	return &st, nil
}

type TableRowFilter struct {
	// The full name of the row filter SQL UDF.
	// Wire name: 'function_name'
	FunctionName string
	// The list of table columns to be passed as input to the row filter
	// function. The column types should match the types of the filter function
	// arguments.
	// Wire name: 'input_column_names'
	InputColumnNames []string
}

func tableRowFilterToPb(st *TableRowFilter) (*tableRowFilterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableRowFilterPb{}
	pb.FunctionName = st.FunctionName

	pb.InputColumnNames = st.InputColumnNames

	return pb, nil
}

func (st *TableRowFilter) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tableRowFilterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tableRowFilterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TableRowFilter) MarshalJSON() ([]byte, error) {
	pb, err := tableRowFilterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type tableRowFilterPb struct {
	// The full name of the row filter SQL UDF.
	FunctionName string `json:"function_name"`
	// The list of table columns to be passed as input to the row filter
	// function. The column types should match the types of the filter function
	// arguments.
	InputColumnNames []string `json:"input_column_names"`
}

func tableRowFilterFromPb(pb *tableRowFilterPb) (*TableRowFilter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableRowFilter{}
	st.FunctionName = pb.FunctionName
	st.InputColumnNames = pb.InputColumnNames

	return st, nil
}

type TableSummary struct {
	// The full name of the table.
	// Wire name: 'full_name'
	FullName string

	// Wire name: 'table_type'
	TableType TableType

	ForceSendFields []string `tf:"-"`
}

func tableSummaryToPb(st *TableSummary) (*tableSummaryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableSummaryPb{}
	pb.FullName = st.FullName

	pb.TableType = st.TableType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *TableSummary) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tableSummaryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tableSummaryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TableSummary) MarshalJSON() ([]byte, error) {
	pb, err := tableSummaryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type tableSummaryPb struct {
	// The full name of the table.
	FullName string `json:"full_name,omitempty"`

	TableType TableType `json:"table_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tableSummaryFromPb(pb *tableSummaryPb) (*TableSummary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableSummary{}
	st.FullName = pb.FullName
	st.TableType = pb.TableType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tableSummaryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tableSummaryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TableType string
type tableTypePb string

const TableTypeExternal TableType = `EXTERNAL`

const TableTypeExternalShallowClone TableType = `EXTERNAL_SHALLOW_CLONE`

const TableTypeForeign TableType = `FOREIGN`

const TableTypeManaged TableType = `MANAGED`

const TableTypeManagedShallowClone TableType = `MANAGED_SHALLOW_CLONE`

const TableTypeMaterializedView TableType = `MATERIALIZED_VIEW`

const TableTypeStreamingTable TableType = `STREAMING_TABLE`

const TableTypeView TableType = `VIEW`

// String representation for [fmt.Print]
func (f *TableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TableType) Set(v string) error {
	switch v {
	case `EXTERNAL`, `EXTERNAL_SHALLOW_CLONE`, `FOREIGN`, `MANAGED`, `MANAGED_SHALLOW_CLONE`, `MATERIALIZED_VIEW`, `STREAMING_TABLE`, `VIEW`:
		*f = TableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "EXTERNAL_SHALLOW_CLONE", "FOREIGN", "MANAGED", "MANAGED_SHALLOW_CLONE", "MATERIALIZED_VIEW", "STREAMING_TABLE", "VIEW"`, v)
	}
}

// Type always returns TableType to satisfy [pflag.Value] interface
func (f *TableType) Type() string {
	return "TableType"
}

func tableTypeToPb(st *TableType) (*tableTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := tableTypePb(*st)
	return &pb, nil
}

func tableTypeFromPb(pb *tableTypePb) (*TableType, error) {
	if pb == nil {
		return nil, nil
	}
	st := TableType(*pb)
	return &st, nil
}

type TagKeyValue struct {
	// name of the tag
	// Wire name: 'key'
	Key string
	// value of the tag associated with the key, could be optional
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func TagKeyValueToPb(st *TagKeyValue) (*TagKeyValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &TagKeyValuePb{}
	pb.Key = st.Key

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *TagKeyValue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &TagKeyValuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TagKeyValueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TagKeyValue) MarshalJSON() ([]byte, error) {
	pb, err := TagKeyValueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TagKeyValuePb struct {
	// name of the tag
	Key string `json:"key,omitempty"`
	// value of the tag associated with the key, could be optional
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func TagKeyValueFromPb(pb *TagKeyValuePb) (*TagKeyValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TagKeyValue{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *TagKeyValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TagKeyValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TemporaryCredentials struct {
	// AWS temporary credentials for API authentication. Read more at
	// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
	// Wire name: 'aws_temp_credentials'
	AwsTempCredentials *AwsCredentials
	// Azure Active Directory token, essentially the Oauth token for Azure
	// Service Principal or Managed Identity. Read more at
	// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
	// Wire name: 'azure_aad'
	AzureAad *AzureActiveDirectoryToken
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	// Wire name: 'expiration_time'
	ExpirationTime int64
	// GCP temporary credentials for API authentication. Read more at
	// https://developers.google.com/identity/protocols/oauth2/service-account
	// Wire name: 'gcp_oauth_token'
	GcpOauthToken *GcpOauthToken

	ForceSendFields []string `tf:"-"`
}

func temporaryCredentialsToPb(st *TemporaryCredentials) (*temporaryCredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &temporaryCredentialsPb{}
	awsTempCredentialsPb, err := awsCredentialsToPb(st.AwsTempCredentials)
	if err != nil {
		return nil, err
	}
	if awsTempCredentialsPb != nil {
		pb.AwsTempCredentials = awsTempCredentialsPb
	}

	azureAadPb, err := azureActiveDirectoryTokenToPb(st.AzureAad)
	if err != nil {
		return nil, err
	}
	if azureAadPb != nil {
		pb.AzureAad = azureAadPb
	}

	pb.ExpirationTime = st.ExpirationTime

	gcpOauthTokenPb, err := gcpOauthTokenToPb(st.GcpOauthToken)
	if err != nil {
		return nil, err
	}
	if gcpOauthTokenPb != nil {
		pb.GcpOauthToken = gcpOauthTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *TemporaryCredentials) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &temporaryCredentialsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := temporaryCredentialsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TemporaryCredentials) MarshalJSON() ([]byte, error) {
	pb, err := temporaryCredentialsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type temporaryCredentialsPb struct {
	// AWS temporary credentials for API authentication. Read more at
	// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
	AwsTempCredentials *awsCredentialsPb `json:"aws_temp_credentials,omitempty"`
	// Azure Active Directory token, essentially the Oauth token for Azure
	// Service Principal or Managed Identity. Read more at
	// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
	AzureAad *azureActiveDirectoryTokenPb `json:"azure_aad,omitempty"`
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// GCP temporary credentials for API authentication. Read more at
	// https://developers.google.com/identity/protocols/oauth2/service-account
	GcpOauthToken *gcpOauthTokenPb `json:"gcp_oauth_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func temporaryCredentialsFromPb(pb *temporaryCredentialsPb) (*TemporaryCredentials, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TemporaryCredentials{}
	awsTempCredentialsField, err := awsCredentialsFromPb(pb.AwsTempCredentials)
	if err != nil {
		return nil, err
	}
	if awsTempCredentialsField != nil {
		st.AwsTempCredentials = awsTempCredentialsField
	}
	azureAadField, err := azureActiveDirectoryTokenFromPb(pb.AzureAad)
	if err != nil {
		return nil, err
	}
	if azureAadField != nil {
		st.AzureAad = azureAadField
	}
	st.ExpirationTime = pb.ExpirationTime
	gcpOauthTokenField, err := gcpOauthTokenFromPb(pb.GcpOauthToken)
	if err != nil {
		return nil, err
	}
	if gcpOauthTokenField != nil {
		st.GcpOauthToken = gcpOauthTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *temporaryCredentialsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st temporaryCredentialsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
type TriggeredUpdateStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	// Wire name: 'last_processed_commit_version'
	LastProcessedCommitVersion int64
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	// Wire name: 'timestamp'
	Timestamp string
	// Progress of the active data synchronization pipeline.
	// Wire name: 'triggered_update_progress'
	TriggeredUpdateProgress *PipelineProgress

	ForceSendFields []string `tf:"-"`
}

func triggeredUpdateStatusToPb(st *TriggeredUpdateStatus) (*triggeredUpdateStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &triggeredUpdateStatusPb{}
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion

	pb.Timestamp = st.Timestamp

	triggeredUpdateProgressPb, err := pipelineProgressToPb(st.TriggeredUpdateProgress)
	if err != nil {
		return nil, err
	}
	if triggeredUpdateProgressPb != nil {
		pb.TriggeredUpdateProgress = triggeredUpdateProgressPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *TriggeredUpdateStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &triggeredUpdateStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := triggeredUpdateStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TriggeredUpdateStatus) MarshalJSON() ([]byte, error) {
	pb, err := triggeredUpdateStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type triggeredUpdateStatusPb struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp string `json:"timestamp,omitempty"`
	// Progress of the active data synchronization pipeline.
	TriggeredUpdateProgress *pipelineProgressPb `json:"triggered_update_progress,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func triggeredUpdateStatusFromPb(pb *triggeredUpdateStatusPb) (*TriggeredUpdateStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TriggeredUpdateStatus{}
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
	st.Timestamp = pb.Timestamp
	triggeredUpdateProgressField, err := pipelineProgressFromPb(pb.TriggeredUpdateProgress)
	if err != nil {
		return nil, err
	}
	if triggeredUpdateProgressField != nil {
		st.TriggeredUpdateProgress = triggeredUpdateProgressField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *triggeredUpdateStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st triggeredUpdateStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete an assignment
type UnassignRequest struct {
	// Query for the ID of the metastore to delete.
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// A workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func unassignRequestToPb(st *UnassignRequest) (*unassignRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &unassignRequestPb{}
	pb.MetastoreId = st.MetastoreId

	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func (st *UnassignRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &unassignRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := unassignRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UnassignRequest) MarshalJSON() ([]byte, error) {
	pb, err := unassignRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type unassignRequestPb struct {
	// Query for the ID of the metastore to delete.
	MetastoreId string `json:"-" url:"metastore_id"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func unassignRequestFromPb(pb *unassignRequestPb) (*UnassignRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UnassignRequest{}
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

type UnassignResponse struct {
}

func unassignResponseToPb(st *UnassignResponse) (*unassignResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &unassignResponsePb{}

	return pb, nil
}

func (st *UnassignResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &unassignResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := unassignResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UnassignResponse) MarshalJSON() ([]byte, error) {
	pb, err := unassignResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type unassignResponsePb struct {
}

func unassignResponseFromPb(pb *unassignResponsePb) (*UnassignResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UnassignResponse{}

	return st, nil
}

type UpdateAssignmentResponse struct {
}

func updateAssignmentResponseToPb(st *UpdateAssignmentResponse) (*updateAssignmentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAssignmentResponsePb{}

	return pb, nil
}

func (st *UpdateAssignmentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateAssignmentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateAssignmentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateAssignmentResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateAssignmentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateAssignmentResponsePb struct {
}

func updateAssignmentResponseFromPb(pb *updateAssignmentResponsePb) (*UpdateAssignmentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAssignmentResponse{}

	return st, nil
}

type UpdateBindingsSecurableType string
type updateBindingsSecurableTypePb string

const UpdateBindingsSecurableTypeCatalog UpdateBindingsSecurableType = `catalog`

const UpdateBindingsSecurableTypeCredential UpdateBindingsSecurableType = `credential`

const UpdateBindingsSecurableTypeExternalLocation UpdateBindingsSecurableType = `external_location`

const UpdateBindingsSecurableTypeStorageCredential UpdateBindingsSecurableType = `storage_credential`

// String representation for [fmt.Print]
func (f *UpdateBindingsSecurableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateBindingsSecurableType) Set(v string) error {
	switch v {
	case `catalog`, `credential`, `external_location`, `storage_credential`:
		*f = UpdateBindingsSecurableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "catalog", "credential", "external_location", "storage_credential"`, v)
	}
}

// Type always returns UpdateBindingsSecurableType to satisfy [pflag.Value] interface
func (f *UpdateBindingsSecurableType) Type() string {
	return "UpdateBindingsSecurableType"
}

func updateBindingsSecurableTypeToPb(st *UpdateBindingsSecurableType) (*updateBindingsSecurableTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := updateBindingsSecurableTypePb(*st)
	return &pb, nil
}

func updateBindingsSecurableTypeFromPb(pb *updateBindingsSecurableTypePb) (*UpdateBindingsSecurableType, error) {
	if pb == nil {
		return nil, nil
	}
	st := UpdateBindingsSecurableType(*pb)
	return &st, nil
}

type UpdateCatalog struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'enable_predictive_optimization'
	EnablePredictiveOptimization EnablePredictiveOptimization
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode CatalogIsolationMode
	// The name of the catalog.
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the catalog.
	// Wire name: 'new_name'
	NewName string
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string
	// Username of current owner of catalog.
	// Wire name: 'owner'
	Owner string
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string

	ForceSendFields []string `tf:"-"`
}

func updateCatalogToPb(st *UpdateCatalog) (*updateCatalogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCatalogPb{}
	pb.Comment = st.Comment

	pb.EnablePredictiveOptimization = st.EnablePredictiveOptimization

	pb.IsolationMode = st.IsolationMode

	pb.Name = st.Name

	pb.NewName = st.NewName

	pb.Options = st.Options

	pb.Owner = st.Owner

	pb.Properties = st.Properties

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateCatalog) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateCatalogPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateCatalogFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateCatalog) MarshalJSON() ([]byte, error) {
	pb, err := updateCatalogToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateCatalogPb struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode CatalogIsolationMode `json:"isolation_mode,omitempty"`
	// The name of the catalog.
	Name string `json:"-" url:"-"`
	// New name for the catalog.
	NewName string `json:"new_name,omitempty"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options,omitempty"`
	// Username of current owner of catalog.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateCatalogFromPb(pb *updateCatalogPb) (*UpdateCatalog, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCatalog{}
	st.Comment = pb.Comment
	st.EnablePredictiveOptimization = pb.EnablePredictiveOptimization
	st.IsolationMode = pb.IsolationMode
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Options = pb.Options
	st.Owner = pb.Owner
	st.Properties = pb.Properties

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateCatalogPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateCatalogPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateConnection struct {
	// Name of the connection.
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the connection.
	// Wire name: 'new_name'
	NewName string
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string
	// Username of current owner of the connection.
	// Wire name: 'owner'
	Owner string

	ForceSendFields []string `tf:"-"`
}

func updateConnectionToPb(st *UpdateConnection) (*updateConnectionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateConnectionPb{}
	pb.Name = st.Name

	pb.NewName = st.NewName

	pb.Options = st.Options

	pb.Owner = st.Owner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateConnection) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateConnectionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateConnectionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateConnection) MarshalJSON() ([]byte, error) {
	pb, err := updateConnectionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateConnectionPb struct {
	// Name of the connection.
	Name string `json:"-" url:"-"`
	// New name for the connection.
	NewName string `json:"new_name,omitempty"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options"`
	// Username of current owner of the connection.
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateConnectionFromPb(pb *updateConnectionPb) (*UpdateConnection, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateConnection{}
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Options = pb.Options
	st.Owner = pb.Owner

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateConnectionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateConnectionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateCredentialRequest struct {
	// The AWS IAM role configuration
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRole
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentity
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	// Wire name: 'force'
	Force bool
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode
	// Name of the credential.
	// Wire name: 'name_arg'
	NameArg string `tf:"-"`
	// New name of credential.
	// Wire name: 'new_name'
	NewName string
	// Username of current owner of credential.
	// Wire name: 'owner'
	Owner string
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	// Wire name: 'read_only'
	ReadOnly bool
	// Supply true to this argument to skip validation of the updated
	// credential.
	// Wire name: 'skip_validation'
	SkipValidation bool

	ForceSendFields []string `tf:"-"`
}

func updateCredentialRequestToPb(st *UpdateCredentialRequest) (*updateCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCredentialRequestPb{}
	awsIamRolePb, err := awsIamRoleToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}

	azureManagedIdentityPb, err := azureManagedIdentityToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}

	azureServicePrincipalPb, err := azureServicePrincipalToPb(st.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalPb != nil {
		pb.AzureServicePrincipal = azureServicePrincipalPb
	}

	pb.Comment = st.Comment

	databricksGcpServiceAccountPb, err := databricksGcpServiceAccountToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}

	pb.Force = st.Force

	pb.IsolationMode = st.IsolationMode

	pb.NameArg = st.NameArg

	pb.NewName = st.NewName

	pb.Owner = st.Owner

	pb.ReadOnly = st.ReadOnly

	pb.SkipValidation = st.SkipValidation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateCredentialRequestPb struct {
	// The AWS IAM role configuration
	AwsIamRole *awsIamRolePb `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *azureManagedIdentityPb `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	AzureServicePrincipal *azureServicePrincipalPb `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount *databricksGcpServiceAccountPb `json:"databricks_gcp_service_account,omitempty"`
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	Force bool `json:"force,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Name of the credential.
	NameArg string `json:"-" url:"-"`
	// New name of credential.
	NewName string `json:"new_name,omitempty"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool `json:"read_only,omitempty"`
	// Supply true to this argument to skip validation of the updated
	// credential.
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateCredentialRequestFromPb(pb *updateCredentialRequestPb) (*UpdateCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCredentialRequest{}
	awsIamRoleField, err := awsIamRoleFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := azureManagedIdentityFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	azureServicePrincipalField, err := azureServicePrincipalFromPb(pb.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalField != nil {
		st.AzureServicePrincipal = azureServicePrincipalField
	}
	st.Comment = pb.Comment
	databricksGcpServiceAccountField, err := databricksGcpServiceAccountFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	st.Force = pb.Force
	st.IsolationMode = pb.IsolationMode
	st.NameArg = pb.NameArg
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateExternalLocation struct {
	// The AWS access point to use when accesing s3 for this external location.
	// Wire name: 'access_point'
	AccessPoint string
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// Name of the storage credential used with this location.
	// Wire name: 'credential_name'
	CredentialName string
	// Encryption options that apply to clients connecting to cloud storage.
	// Wire name: 'encryption_details'
	EncryptionDetails *EncryptionDetails
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	// Wire name: 'fallback'
	Fallback bool
	// Force update even if changing url invalidates dependent external tables
	// or mounts.
	// Wire name: 'force'
	Force bool

	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode
	// Name of the external location.
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the external location.
	// Wire name: 'new_name'
	NewName string
	// The owner of the external location.
	// Wire name: 'owner'
	Owner string
	// Indicates whether the external location is read-only.
	// Wire name: 'read_only'
	ReadOnly bool
	// Skips validation of the storage credential associated with the external
	// location.
	// Wire name: 'skip_validation'
	SkipValidation bool
	// Path URL of the external location.
	// Wire name: 'url'
	Url string

	ForceSendFields []string `tf:"-"`
}

func updateExternalLocationToPb(st *UpdateExternalLocation) (*updateExternalLocationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateExternalLocationPb{}
	pb.AccessPoint = st.AccessPoint

	pb.Comment = st.Comment

	pb.CredentialName = st.CredentialName

	encryptionDetailsPb, err := encryptionDetailsToPb(st.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsPb != nil {
		pb.EncryptionDetails = encryptionDetailsPb
	}

	pb.Fallback = st.Fallback

	pb.Force = st.Force

	pb.IsolationMode = st.IsolationMode

	pb.Name = st.Name

	pb.NewName = st.NewName

	pb.Owner = st.Owner

	pb.ReadOnly = st.ReadOnly

	pb.SkipValidation = st.SkipValidation

	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateExternalLocation) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateExternalLocationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateExternalLocationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateExternalLocation) MarshalJSON() ([]byte, error) {
	pb, err := updateExternalLocationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateExternalLocationPb struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string `json:"access_point,omitempty"`
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of the storage credential used with this location.
	CredentialName string `json:"credential_name,omitempty"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *encryptionDetailsPb `json:"encryption_details,omitempty"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback bool `json:"fallback,omitempty"`
	// Force update even if changing url invalidates dependent external tables
	// or mounts.
	Force bool `json:"force,omitempty"`

	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Name of the external location.
	Name string `json:"-" url:"-"`
	// New name for the external location.
	NewName string `json:"new_name,omitempty"`
	// The owner of the external location.
	Owner string `json:"owner,omitempty"`
	// Indicates whether the external location is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation bool `json:"skip_validation,omitempty"`
	// Path URL of the external location.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateExternalLocationFromPb(pb *updateExternalLocationPb) (*UpdateExternalLocation, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExternalLocation{}
	st.AccessPoint = pb.AccessPoint
	st.Comment = pb.Comment
	st.CredentialName = pb.CredentialName
	encryptionDetailsField, err := encryptionDetailsFromPb(pb.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsField != nil {
		st.EncryptionDetails = encryptionDetailsField
	}
	st.Fallback = pb.Fallback
	st.Force = pb.Force
	st.IsolationMode = pb.IsolationMode
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateExternalLocationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateExternalLocationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateFunction struct {
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	// Wire name: 'name'
	Name string `tf:"-"`
	// Username of current owner of function.
	// Wire name: 'owner'
	Owner string

	ForceSendFields []string `tf:"-"`
}

func updateFunctionToPb(st *UpdateFunction) (*updateFunctionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateFunctionPb{}
	pb.Name = st.Name

	pb.Owner = st.Owner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateFunction) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateFunctionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateFunctionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateFunction) MarshalJSON() ([]byte, error) {
	pb, err := updateFunctionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateFunctionPb struct {
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string `json:"-" url:"-"`
	// Username of current owner of function.
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateFunctionFromPb(pb *updateFunctionPb) (*UpdateFunction, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateFunction{}
	st.Name = pb.Name
	st.Owner = pb.Owner

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateFunctionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateFunctionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateMetastore struct {
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	// Wire name: 'delta_sharing_organization_name'
	DeltaSharingOrganizationName string
	// The lifetime of delta sharing recipient token in seconds.
	// Wire name: 'delta_sharing_recipient_token_lifetime_in_seconds'
	DeltaSharingRecipientTokenLifetimeInSeconds int64
	// The scope of Delta Sharing enabled for the metastore.
	// Wire name: 'delta_sharing_scope'
	DeltaSharingScope UpdateMetastoreDeltaSharingScope
	// Unique ID of the metastore.
	// Wire name: 'id'
	Id string `tf:"-"`
	// New name for the metastore.
	// Wire name: 'new_name'
	NewName string
	// The owner of the metastore.
	// Wire name: 'owner'
	Owner string
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	// Wire name: 'privilege_model_version'
	PrivilegeModelVersion string
	// UUID of storage credential to access the metastore storage_root.
	// Wire name: 'storage_root_credential_id'
	StorageRootCredentialId string

	ForceSendFields []string `tf:"-"`
}

func updateMetastoreToPb(st *UpdateMetastore) (*updateMetastorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateMetastorePb{}
	pb.DeltaSharingOrganizationName = st.DeltaSharingOrganizationName

	pb.DeltaSharingRecipientTokenLifetimeInSeconds = st.DeltaSharingRecipientTokenLifetimeInSeconds

	pb.DeltaSharingScope = st.DeltaSharingScope

	pb.Id = st.Id

	pb.NewName = st.NewName

	pb.Owner = st.Owner

	pb.PrivilegeModelVersion = st.PrivilegeModelVersion

	pb.StorageRootCredentialId = st.StorageRootCredentialId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateMetastore) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateMetastorePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateMetastoreFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateMetastore) MarshalJSON() ([]byte, error) {
	pb, err := updateMetastoreToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateMetastorePb struct {
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope UpdateMetastoreDeltaSharingScope `json:"delta_sharing_scope,omitempty"`
	// Unique ID of the metastore.
	Id string `json:"-" url:"-"`
	// New name for the metastore.
	NewName string `json:"new_name,omitempty"`
	// The owner of the metastore.
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateMetastoreFromPb(pb *updateMetastorePb) (*UpdateMetastore, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateMetastore{}
	st.DeltaSharingOrganizationName = pb.DeltaSharingOrganizationName
	st.DeltaSharingRecipientTokenLifetimeInSeconds = pb.DeltaSharingRecipientTokenLifetimeInSeconds
	st.DeltaSharingScope = pb.DeltaSharingScope
	st.Id = pb.Id
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.PrivilegeModelVersion = pb.PrivilegeModelVersion
	st.StorageRootCredentialId = pb.StorageRootCredentialId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateMetastorePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateMetastorePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateMetastoreAssignment struct {
	// The name of the default catalog in the metastore. This field is
	// depracted. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	// Wire name: 'default_catalog_name'
	DefaultCatalogName string
	// The unique ID of the metastore.
	// Wire name: 'metastore_id'
	MetastoreId string
	// A workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func updateMetastoreAssignmentToPb(st *UpdateMetastoreAssignment) (*updateMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateMetastoreAssignmentPb{}
	pb.DefaultCatalogName = st.DefaultCatalogName

	pb.MetastoreId = st.MetastoreId

	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateMetastoreAssignment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateMetastoreAssignmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateMetastoreAssignmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateMetastoreAssignment) MarshalJSON() ([]byte, error) {
	pb, err := updateMetastoreAssignmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateMetastoreAssignmentPb struct {
	// The name of the default catalog in the metastore. This field is
	// depracted. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	// The unique ID of the metastore.
	MetastoreId string `json:"metastore_id,omitempty"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateMetastoreAssignmentFromPb(pb *updateMetastoreAssignmentPb) (*UpdateMetastoreAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateMetastoreAssignment{}
	st.DefaultCatalogName = pb.DefaultCatalogName
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateMetastoreAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateMetastoreAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The scope of Delta Sharing enabled for the metastore.
type UpdateMetastoreDeltaSharingScope string
type updateMetastoreDeltaSharingScopePb string

const UpdateMetastoreDeltaSharingScopeInternal UpdateMetastoreDeltaSharingScope = `INTERNAL`

const UpdateMetastoreDeltaSharingScopeInternalAndExternal UpdateMetastoreDeltaSharingScope = `INTERNAL_AND_EXTERNAL`

// String representation for [fmt.Print]
func (f *UpdateMetastoreDeltaSharingScope) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateMetastoreDeltaSharingScope) Set(v string) error {
	switch v {
	case `INTERNAL`, `INTERNAL_AND_EXTERNAL`:
		*f = UpdateMetastoreDeltaSharingScope(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INTERNAL", "INTERNAL_AND_EXTERNAL"`, v)
	}
}

// Type always returns UpdateMetastoreDeltaSharingScope to satisfy [pflag.Value] interface
func (f *UpdateMetastoreDeltaSharingScope) Type() string {
	return "UpdateMetastoreDeltaSharingScope"
}

func updateMetastoreDeltaSharingScopeToPb(st *UpdateMetastoreDeltaSharingScope) (*updateMetastoreDeltaSharingScopePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := updateMetastoreDeltaSharingScopePb(*st)
	return &pb, nil
}

func updateMetastoreDeltaSharingScopeFromPb(pb *updateMetastoreDeltaSharingScopePb) (*UpdateMetastoreDeltaSharingScope, error) {
	if pb == nil {
		return nil, nil
	}
	st := UpdateMetastoreDeltaSharingScope(*pb)
	return &st, nil
}

type UpdateModelVersionRequest struct {
	// The comment attached to the model version
	// Wire name: 'comment'
	Comment string
	// The three-level (fully qualified) name of the model version
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// The integer version number of the model version
	// Wire name: 'version'
	Version int `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func updateModelVersionRequestToPb(st *UpdateModelVersionRequest) (*updateModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateModelVersionRequestPb{}
	pb.Comment = st.Comment

	pb.FullName = st.FullName

	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateModelVersionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateModelVersionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateModelVersionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateModelVersionRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateModelVersionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateModelVersionRequestPb struct {
	// The comment attached to the model version
	Comment string `json:"comment,omitempty"`
	// The three-level (fully qualified) name of the model version
	FullName string `json:"-" url:"-"`
	// The integer version number of the model version
	Version int `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateModelVersionRequestFromPb(pb *updateModelVersionRequestPb) (*UpdateModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateModelVersionRequest{}
	st.Comment = pb.Comment
	st.FullName = pb.FullName
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateModelVersionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateModelVersionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateMonitor struct {
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	// Wire name: 'baseline_table_name'
	BaselineTableName string
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	// Wire name: 'custom_metrics'
	CustomMetrics []MonitorMetric
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	// Wire name: 'dashboard_id'
	DashboardId string
	// The data classification config for the monitor.
	// Wire name: 'data_classification_config'
	DataClassificationConfig *MonitorDataClassificationConfig
	// Configuration for monitoring inference logs.
	// Wire name: 'inference_log'
	InferenceLog *MonitorInferenceLog
	// The notification settings for the monitor.
	// Wire name: 'notifications'
	Notifications *MonitorNotifications
	// Schema where output metric tables are created.
	// Wire name: 'output_schema_name'
	OutputSchemaName string
	// The schedule for automatically updating and refreshing metric tables.
	// Wire name: 'schedule'
	Schedule *MonitorCronSchedule
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	// Wire name: 'slicing_exprs'
	SlicingExprs []string
	// Configuration for monitoring snapshot tables.
	// Wire name: 'snapshot'
	Snapshot *MonitorSnapshot
	// Full name of the table.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
	// Configuration for monitoring time series tables.
	// Wire name: 'time_series'
	TimeSeries *MonitorTimeSeries

	ForceSendFields []string `tf:"-"`
}

func updateMonitorToPb(st *UpdateMonitor) (*updateMonitorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateMonitorPb{}
	pb.BaselineTableName = st.BaselineTableName

	var customMetricsPb []monitorMetricPb
	for _, item := range st.CustomMetrics {
		itemPb, err := monitorMetricToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customMetricsPb = append(customMetricsPb, *itemPb)
		}
	}
	pb.CustomMetrics = customMetricsPb

	pb.DashboardId = st.DashboardId

	dataClassificationConfigPb, err := monitorDataClassificationConfigToPb(st.DataClassificationConfig)
	if err != nil {
		return nil, err
	}
	if dataClassificationConfigPb != nil {
		pb.DataClassificationConfig = dataClassificationConfigPb
	}

	inferenceLogPb, err := monitorInferenceLogToPb(st.InferenceLog)
	if err != nil {
		return nil, err
	}
	if inferenceLogPb != nil {
		pb.InferenceLog = inferenceLogPb
	}

	notificationsPb, err := monitorNotificationsToPb(st.Notifications)
	if err != nil {
		return nil, err
	}
	if notificationsPb != nil {
		pb.Notifications = notificationsPb
	}

	pb.OutputSchemaName = st.OutputSchemaName

	schedulePb, err := monitorCronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}

	pb.SlicingExprs = st.SlicingExprs

	snapshotPb, err := monitorSnapshotToPb(st.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotPb != nil {
		pb.Snapshot = snapshotPb
	}

	pb.TableName = st.TableName

	timeSeriesPb, err := monitorTimeSeriesToPb(st.TimeSeries)
	if err != nil {
		return nil, err
	}
	if timeSeriesPb != nil {
		pb.TimeSeries = timeSeriesPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateMonitor) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateMonitorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateMonitorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateMonitor) MarshalJSON() ([]byte, error) {
	pb, err := updateMonitorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateMonitorPb struct {
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName string `json:"baseline_table_name,omitempty"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics []monitorMetricPb `json:"custom_metrics,omitempty"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId string `json:"dashboard_id,omitempty"`
	// The data classification config for the monitor.
	DataClassificationConfig *monitorDataClassificationConfigPb `json:"data_classification_config,omitempty"`
	// Configuration for monitoring inference logs.
	InferenceLog *monitorInferenceLogPb `json:"inference_log,omitempty"`
	// The notification settings for the monitor.
	Notifications *monitorNotificationsPb `json:"notifications,omitempty"`
	// Schema where output metric tables are created.
	OutputSchemaName string `json:"output_schema_name"`
	// The schedule for automatically updating and refreshing metric tables.
	Schedule *monitorCronSchedulePb `json:"schedule,omitempty"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs []string `json:"slicing_exprs,omitempty"`
	// Configuration for monitoring snapshot tables.
	Snapshot *monitorSnapshotPb `json:"snapshot,omitempty"`
	// Full name of the table.
	TableName string `json:"-" url:"-"`
	// Configuration for monitoring time series tables.
	TimeSeries *monitorTimeSeriesPb `json:"time_series,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateMonitorFromPb(pb *updateMonitorPb) (*UpdateMonitor, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateMonitor{}
	st.BaselineTableName = pb.BaselineTableName

	var customMetricsField []MonitorMetric
	for _, itemPb := range pb.CustomMetrics {
		item, err := monitorMetricFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customMetricsField = append(customMetricsField, *item)
		}
	}
	st.CustomMetrics = customMetricsField
	st.DashboardId = pb.DashboardId
	dataClassificationConfigField, err := monitorDataClassificationConfigFromPb(pb.DataClassificationConfig)
	if err != nil {
		return nil, err
	}
	if dataClassificationConfigField != nil {
		st.DataClassificationConfig = dataClassificationConfigField
	}
	inferenceLogField, err := monitorInferenceLogFromPb(pb.InferenceLog)
	if err != nil {
		return nil, err
	}
	if inferenceLogField != nil {
		st.InferenceLog = inferenceLogField
	}
	notificationsField, err := monitorNotificationsFromPb(pb.Notifications)
	if err != nil {
		return nil, err
	}
	if notificationsField != nil {
		st.Notifications = notificationsField
	}
	st.OutputSchemaName = pb.OutputSchemaName
	scheduleField, err := monitorCronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	st.SlicingExprs = pb.SlicingExprs
	snapshotField, err := monitorSnapshotFromPb(pb.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotField != nil {
		st.Snapshot = snapshotField
	}
	st.TableName = pb.TableName
	timeSeriesField, err := monitorTimeSeriesFromPb(pb.TimeSeries)
	if err != nil {
		return nil, err
	}
	if timeSeriesField != nil {
		st.TimeSeries = timeSeriesField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateMonitorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateMonitorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdatePermissions struct {
	// Array of permissions change objects.
	// Wire name: 'changes'
	Changes []PermissionsChange
	// Full name of securable.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// Type of securable.
	// Wire name: 'securable_type'
	SecurableType SecurableType `tf:"-"`
}

func updatePermissionsToPb(st *UpdatePermissions) (*updatePermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updatePermissionsPb{}

	var changesPb []permissionsChangePb
	for _, item := range st.Changes {
		itemPb, err := permissionsChangeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			changesPb = append(changesPb, *itemPb)
		}
	}
	pb.Changes = changesPb

	pb.FullName = st.FullName

	pb.SecurableType = st.SecurableType

	return pb, nil
}

func (st *UpdatePermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updatePermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updatePermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdatePermissions) MarshalJSON() ([]byte, error) {
	pb, err := updatePermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updatePermissionsPb struct {
	// Array of permissions change objects.
	Changes []permissionsChangePb `json:"changes,omitempty"`
	// Full name of securable.
	FullName string `json:"-" url:"-"`
	// Type of securable.
	SecurableType SecurableType `json:"-" url:"-"`
}

func updatePermissionsFromPb(pb *updatePermissionsPb) (*UpdatePermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePermissions{}

	var changesField []PermissionsChange
	for _, itemPb := range pb.Changes {
		item, err := permissionsChangeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			changesField = append(changesField, *item)
		}
	}
	st.Changes = changesField
	st.FullName = pb.FullName
	st.SecurableType = pb.SecurableType

	return st, nil
}

type UpdateRegisteredModelRequest struct {
	// The comment attached to the registered model
	// Wire name: 'comment'
	Comment string
	// The three-level (fully qualified) name of the registered model
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// New name for the registered model.
	// Wire name: 'new_name'
	NewName string
	// The identifier of the user who owns the registered model
	// Wire name: 'owner'
	Owner string

	ForceSendFields []string `tf:"-"`
}

func updateRegisteredModelRequestToPb(st *UpdateRegisteredModelRequest) (*updateRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRegisteredModelRequestPb{}
	pb.Comment = st.Comment

	pb.FullName = st.FullName

	pb.NewName = st.NewName

	pb.Owner = st.Owner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateRegisteredModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateRegisteredModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateRegisteredModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateRegisteredModelRequestPb struct {
	// The comment attached to the registered model
	Comment string `json:"comment,omitempty"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" url:"-"`
	// New name for the registered model.
	NewName string `json:"new_name,omitempty"`
	// The identifier of the user who owns the registered model
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateRegisteredModelRequestFromPb(pb *updateRegisteredModelRequestPb) (*UpdateRegisteredModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRegisteredModelRequest{}
	st.Comment = pb.Comment
	st.FullName = pb.FullName
	st.NewName = pb.NewName
	st.Owner = pb.Owner

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateRegisteredModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateRegisteredModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateResponse struct {
}

func updateResponseToPb(st *UpdateResponse) (*updateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateResponsePb{}

	return pb, nil
}

func (st *UpdateResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateResponsePb struct {
}

func updateResponseFromPb(pb *updateResponsePb) (*UpdateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateResponse{}

	return st, nil
}

type UpdateSchema struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'enable_predictive_optimization'
	EnablePredictiveOptimization EnablePredictiveOptimization
	// Full name of the schema.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// New name for the schema.
	// Wire name: 'new_name'
	NewName string
	// Username of current owner of schema.
	// Wire name: 'owner'
	Owner string
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string

	ForceSendFields []string `tf:"-"`
}

func updateSchemaToPb(st *UpdateSchema) (*updateSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateSchemaPb{}
	pb.Comment = st.Comment

	pb.EnablePredictiveOptimization = st.EnablePredictiveOptimization

	pb.FullName = st.FullName

	pb.NewName = st.NewName

	pb.Owner = st.Owner

	pb.Properties = st.Properties

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateSchema) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateSchemaPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateSchemaFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateSchema) MarshalJSON() ([]byte, error) {
	pb, err := updateSchemaToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateSchemaPb struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// Full name of the schema.
	FullName string `json:"-" url:"-"`
	// New name for the schema.
	NewName string `json:"new_name,omitempty"`
	// Username of current owner of schema.
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	Properties map[string]string `json:"properties,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateSchemaFromPb(pb *updateSchemaPb) (*UpdateSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateSchema{}
	st.Comment = pb.Comment
	st.EnablePredictiveOptimization = pb.EnablePredictiveOptimization
	st.FullName = pb.FullName
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.Properties = pb.Properties

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateSchemaPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateSchemaPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateStorageCredential struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRoleRequest
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentityResponse
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal
	// The Cloudflare API token configuration.
	// Wire name: 'cloudflare_api_token'
	CloudflareApiToken *CloudflareApiToken
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string
	// The Databricks managed GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest
	// Force update even if there are dependent external locations or external
	// tables.
	// Wire name: 'force'
	Force bool

	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode
	// Name of the storage credential.
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the storage credential.
	// Wire name: 'new_name'
	NewName string
	// Username of current owner of credential.
	// Wire name: 'owner'
	Owner string
	// Whether the storage credential is only usable for read operations.
	// Wire name: 'read_only'
	ReadOnly bool
	// Supplying true to this argument skips validation of the updated
	// credential.
	// Wire name: 'skip_validation'
	SkipValidation bool

	ForceSendFields []string `tf:"-"`
}

func updateStorageCredentialToPb(st *UpdateStorageCredential) (*updateStorageCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateStorageCredentialPb{}
	awsIamRolePb, err := awsIamRoleRequestToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}

	azureManagedIdentityPb, err := azureManagedIdentityResponseToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}

	azureServicePrincipalPb, err := azureServicePrincipalToPb(st.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalPb != nil {
		pb.AzureServicePrincipal = azureServicePrincipalPb
	}

	cloudflareApiTokenPb, err := cloudflareApiTokenToPb(st.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenPb != nil {
		pb.CloudflareApiToken = cloudflareApiTokenPb
	}

	pb.Comment = st.Comment

	databricksGcpServiceAccountPb, err := databricksGcpServiceAccountRequestToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}

	pb.Force = st.Force

	pb.IsolationMode = st.IsolationMode

	pb.Name = st.Name

	pb.NewName = st.NewName

	pb.Owner = st.Owner

	pb.ReadOnly = st.ReadOnly

	pb.SkipValidation = st.SkipValidation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateStorageCredential) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateStorageCredentialPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateStorageCredentialFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateStorageCredential) MarshalJSON() ([]byte, error) {
	pb, err := updateStorageCredentialToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateStorageCredentialPb struct {
	// The AWS IAM role configuration.
	AwsIamRole *awsIamRoleRequestPb `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *azureManagedIdentityResponsePb `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *azureServicePrincipalPb `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *cloudflareApiTokenPb `json:"cloudflare_api_token,omitempty"`
	// Comment associated with the credential.
	Comment string `json:"comment,omitempty"`
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *databricksGcpServiceAccountRequestPb `json:"databricks_gcp_service_account,omitempty"`
	// Force update even if there are dependent external locations or external
	// tables.
	Force bool `json:"force,omitempty"`

	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Name of the storage credential.
	Name string `json:"-" url:"-"`
	// New name for the storage credential.
	NewName string `json:"new_name,omitempty"`
	// Username of current owner of credential.
	Owner string `json:"owner,omitempty"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `json:"read_only,omitempty"`
	// Supplying true to this argument skips validation of the updated
	// credential.
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateStorageCredentialFromPb(pb *updateStorageCredentialPb) (*UpdateStorageCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateStorageCredential{}
	awsIamRoleField, err := awsIamRoleRequestFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := azureManagedIdentityResponseFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	azureServicePrincipalField, err := azureServicePrincipalFromPb(pb.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalField != nil {
		st.AzureServicePrincipal = azureServicePrincipalField
	}
	cloudflareApiTokenField, err := cloudflareApiTokenFromPb(pb.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenField != nil {
		st.CloudflareApiToken = cloudflareApiTokenField
	}
	st.Comment = pb.Comment
	databricksGcpServiceAccountField, err := databricksGcpServiceAccountRequestFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	st.Force = pb.Force
	st.IsolationMode = pb.IsolationMode
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateStorageCredentialPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateStorageCredentialPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Update a table owner.
type UpdateTableRequest struct {
	// Full name of the table.
	// Wire name: 'full_name'
	FullName string `tf:"-"`

	// Wire name: 'owner'
	Owner string

	ForceSendFields []string `tf:"-"`
}

func updateTableRequestToPb(st *UpdateTableRequest) (*updateTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateTableRequestPb{}
	pb.FullName = st.FullName

	pb.Owner = st.Owner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateTableRequestPb struct {
	// Full name of the table.
	FullName string `json:"-" url:"-"`

	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateTableRequestFromPb(pb *updateTableRequestPb) (*UpdateTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateTableRequest{}
	st.FullName = pb.FullName
	st.Owner = pb.Owner

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateTableRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateTableRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateVolumeRequestContent struct {
	// The comment attached to the volume
	// Wire name: 'comment'
	Comment string
	// The three-level (fully qualified) name of the volume
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the volume.
	// Wire name: 'new_name'
	NewName string
	// The identifier of the user who owns the volume
	// Wire name: 'owner'
	Owner string

	ForceSendFields []string `tf:"-"`
}

func updateVolumeRequestContentToPb(st *UpdateVolumeRequestContent) (*updateVolumeRequestContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateVolumeRequestContentPb{}
	pb.Comment = st.Comment

	pb.Name = st.Name

	pb.NewName = st.NewName

	pb.Owner = st.Owner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateVolumeRequestContent) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateVolumeRequestContentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateVolumeRequestContentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateVolumeRequestContent) MarshalJSON() ([]byte, error) {
	pb, err := updateVolumeRequestContentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateVolumeRequestContentPb struct {
	// The comment attached to the volume
	Comment string `json:"comment,omitempty"`
	// The three-level (fully qualified) name of the volume
	Name string `json:"-" url:"-"`
	// New name for the volume.
	NewName string `json:"new_name,omitempty"`
	// The identifier of the user who owns the volume
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateVolumeRequestContentFromPb(pb *updateVolumeRequestContentPb) (*UpdateVolumeRequestContent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateVolumeRequestContent{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Owner = pb.Owner

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateVolumeRequestContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateVolumeRequestContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateWorkspaceBindings struct {
	// A list of workspace IDs.
	// Wire name: 'assign_workspaces'
	AssignWorkspaces []int64
	// The name of the catalog.
	// Wire name: 'name'
	Name string `tf:"-"`
	// A list of workspace IDs.
	// Wire name: 'unassign_workspaces'
	UnassignWorkspaces []int64
}

func updateWorkspaceBindingsToPb(st *UpdateWorkspaceBindings) (*updateWorkspaceBindingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateWorkspaceBindingsPb{}
	pb.AssignWorkspaces = st.AssignWorkspaces

	pb.Name = st.Name

	pb.UnassignWorkspaces = st.UnassignWorkspaces

	return pb, nil
}

func (st *UpdateWorkspaceBindings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateWorkspaceBindingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateWorkspaceBindingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateWorkspaceBindings) MarshalJSON() ([]byte, error) {
	pb, err := updateWorkspaceBindingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateWorkspaceBindingsPb struct {
	// A list of workspace IDs.
	AssignWorkspaces []int64 `json:"assign_workspaces,omitempty"`
	// The name of the catalog.
	Name string `json:"-" url:"-"`
	// A list of workspace IDs.
	UnassignWorkspaces []int64 `json:"unassign_workspaces,omitempty"`
}

func updateWorkspaceBindingsFromPb(pb *updateWorkspaceBindingsPb) (*UpdateWorkspaceBindings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWorkspaceBindings{}
	st.AssignWorkspaces = pb.AssignWorkspaces
	st.Name = pb.Name
	st.UnassignWorkspaces = pb.UnassignWorkspaces

	return st, nil
}

type UpdateWorkspaceBindingsParameters struct {
	// List of workspace bindings
	// Wire name: 'add'
	Add []WorkspaceBinding
	// List of workspace bindings
	// Wire name: 'remove'
	Remove []WorkspaceBinding
	// The name of the securable.
	// Wire name: 'securable_name'
	SecurableName string `tf:"-"`
	// The type of the securable to bind to a workspace.
	// Wire name: 'securable_type'
	SecurableType UpdateBindingsSecurableType `tf:"-"`
}

func updateWorkspaceBindingsParametersToPb(st *UpdateWorkspaceBindingsParameters) (*updateWorkspaceBindingsParametersPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateWorkspaceBindingsParametersPb{}

	var addPb []workspaceBindingPb
	for _, item := range st.Add {
		itemPb, err := workspaceBindingToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			addPb = append(addPb, *itemPb)
		}
	}
	pb.Add = addPb

	var removePb []workspaceBindingPb
	for _, item := range st.Remove {
		itemPb, err := workspaceBindingToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			removePb = append(removePb, *itemPb)
		}
	}
	pb.Remove = removePb

	pb.SecurableName = st.SecurableName

	pb.SecurableType = st.SecurableType

	return pb, nil
}

func (st *UpdateWorkspaceBindingsParameters) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateWorkspaceBindingsParametersPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateWorkspaceBindingsParametersFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateWorkspaceBindingsParameters) MarshalJSON() ([]byte, error) {
	pb, err := updateWorkspaceBindingsParametersToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateWorkspaceBindingsParametersPb struct {
	// List of workspace bindings
	Add []workspaceBindingPb `json:"add,omitempty"`
	// List of workspace bindings
	Remove []workspaceBindingPb `json:"remove,omitempty"`
	// The name of the securable.
	SecurableName string `json:"-" url:"-"`
	// The type of the securable to bind to a workspace.
	SecurableType UpdateBindingsSecurableType `json:"-" url:"-"`
}

func updateWorkspaceBindingsParametersFromPb(pb *updateWorkspaceBindingsParametersPb) (*UpdateWorkspaceBindingsParameters, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWorkspaceBindingsParameters{}

	var addField []WorkspaceBinding
	for _, itemPb := range pb.Add {
		item, err := workspaceBindingFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			addField = append(addField, *item)
		}
	}
	st.Add = addField

	var removeField []WorkspaceBinding
	for _, itemPb := range pb.Remove {
		item, err := workspaceBindingFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			removeField = append(removeField, *item)
		}
	}
	st.Remove = removeField
	st.SecurableName = pb.SecurableName
	st.SecurableType = pb.SecurableType

	return st, nil
}

// Next ID: 17
type ValidateCredentialRequest struct {
	// The AWS IAM role configuration
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRole
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentity
	// Required. The name of an existing credential or long-lived cloud
	// credential to validate.
	// Wire name: 'credential_name'
	CredentialName string
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount
	// The name of an existing external location to validate. Only applicable
	// for storage credentials (purpose is **STORAGE**.)
	// Wire name: 'external_location_name'
	ExternalLocationName string
	// The purpose of the credential. This should only be used when the
	// credential is specified.
	// Wire name: 'purpose'
	Purpose CredentialPurpose
	// Whether the credential is only usable for read operations. Only
	// applicable for storage credentials (purpose is **STORAGE**.)
	// Wire name: 'read_only'
	ReadOnly bool
	// The external location url to validate. Only applicable when purpose is
	// **STORAGE**.
	// Wire name: 'url'
	Url string

	ForceSendFields []string `tf:"-"`
}

func validateCredentialRequestToPb(st *ValidateCredentialRequest) (*validateCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &validateCredentialRequestPb{}
	awsIamRolePb, err := awsIamRoleToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}

	azureManagedIdentityPb, err := azureManagedIdentityToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}

	pb.CredentialName = st.CredentialName

	databricksGcpServiceAccountPb, err := databricksGcpServiceAccountToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}

	pb.ExternalLocationName = st.ExternalLocationName

	pb.Purpose = st.Purpose

	pb.ReadOnly = st.ReadOnly

	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ValidateCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &validateCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := validateCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ValidateCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := validateCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type validateCredentialRequestPb struct {
	// The AWS IAM role configuration
	AwsIamRole *awsIamRolePb `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *azureManagedIdentityPb `json:"azure_managed_identity,omitempty"`
	// Required. The name of an existing credential or long-lived cloud
	// credential to validate.
	CredentialName string `json:"credential_name,omitempty"`
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount *databricksGcpServiceAccountPb `json:"databricks_gcp_service_account,omitempty"`
	// The name of an existing external location to validate. Only applicable
	// for storage credentials (purpose is **STORAGE**.)
	ExternalLocationName string `json:"external_location_name,omitempty"`
	// The purpose of the credential. This should only be used when the
	// credential is specified.
	Purpose CredentialPurpose `json:"purpose,omitempty"`
	// Whether the credential is only usable for read operations. Only
	// applicable for storage credentials (purpose is **STORAGE**.)
	ReadOnly bool `json:"read_only,omitempty"`
	// The external location url to validate. Only applicable when purpose is
	// **STORAGE**.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func validateCredentialRequestFromPb(pb *validateCredentialRequestPb) (*ValidateCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ValidateCredentialRequest{}
	awsIamRoleField, err := awsIamRoleFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := azureManagedIdentityFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	st.CredentialName = pb.CredentialName
	databricksGcpServiceAccountField, err := databricksGcpServiceAccountFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	st.ExternalLocationName = pb.ExternalLocationName
	st.Purpose = pb.Purpose
	st.ReadOnly = pb.ReadOnly
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *validateCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st validateCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ValidateCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage. Only
	// applicable for when purpose is **STORAGE**.
	// Wire name: 'isDir'
	IsDir bool
	// The results of the validation check.
	// Wire name: 'results'
	Results []CredentialValidationResult

	ForceSendFields []string `tf:"-"`
}

func validateCredentialResponseToPb(st *ValidateCredentialResponse) (*validateCredentialResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &validateCredentialResponsePb{}
	pb.IsDir = st.IsDir

	var resultsPb []credentialValidationResultPb
	for _, item := range st.Results {
		itemPb, err := credentialValidationResultToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resultsPb = append(resultsPb, *itemPb)
		}
	}
	pb.Results = resultsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ValidateCredentialResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &validateCredentialResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := validateCredentialResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ValidateCredentialResponse) MarshalJSON() ([]byte, error) {
	pb, err := validateCredentialResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type validateCredentialResponsePb struct {
	// Whether the tested location is a directory in cloud storage. Only
	// applicable for when purpose is **STORAGE**.
	IsDir bool `json:"isDir,omitempty"`
	// The results of the validation check.
	Results []credentialValidationResultPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func validateCredentialResponseFromPb(pb *validateCredentialResponsePb) (*ValidateCredentialResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ValidateCredentialResponse{}
	st.IsDir = pb.IsDir

	var resultsField []CredentialValidationResult
	for _, itemPb := range pb.Results {
		item, err := credentialValidationResultFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resultsField = append(resultsField, *item)
		}
	}
	st.Results = resultsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *validateCredentialResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st validateCredentialResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// A enum represents the result of the file operation
type ValidateCredentialResult string
type validateCredentialResultPb string

const ValidateCredentialResultFail ValidateCredentialResult = `FAIL`

const ValidateCredentialResultPass ValidateCredentialResult = `PASS`

const ValidateCredentialResultSkip ValidateCredentialResult = `SKIP`

// String representation for [fmt.Print]
func (f *ValidateCredentialResult) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ValidateCredentialResult) Set(v string) error {
	switch v {
	case `FAIL`, `PASS`, `SKIP`:
		*f = ValidateCredentialResult(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAIL", "PASS", "SKIP"`, v)
	}
}

// Type always returns ValidateCredentialResult to satisfy [pflag.Value] interface
func (f *ValidateCredentialResult) Type() string {
	return "ValidateCredentialResult"
}

func validateCredentialResultToPb(st *ValidateCredentialResult) (*validateCredentialResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := validateCredentialResultPb(*st)
	return &pb, nil
}

func validateCredentialResultFromPb(pb *validateCredentialResultPb) (*ValidateCredentialResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := ValidateCredentialResult(*pb)
	return &st, nil
}

type ValidateStorageCredential struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRoleRequest
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentityRequest
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal
	// The Cloudflare API token configuration.
	// Wire name: 'cloudflare_api_token'
	CloudflareApiToken *CloudflareApiToken
	// The Databricks created GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest
	// The name of an existing external location to validate.
	// Wire name: 'external_location_name'
	ExternalLocationName string
	// Whether the storage credential is only usable for read operations.
	// Wire name: 'read_only'
	ReadOnly bool
	// The name of the storage credential to validate.
	// Wire name: 'storage_credential_name'
	StorageCredentialName string
	// The external location url to validate.
	// Wire name: 'url'
	Url string

	ForceSendFields []string `tf:"-"`
}

func validateStorageCredentialToPb(st *ValidateStorageCredential) (*validateStorageCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &validateStorageCredentialPb{}
	awsIamRolePb, err := awsIamRoleRequestToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}

	azureManagedIdentityPb, err := azureManagedIdentityRequestToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}

	azureServicePrincipalPb, err := azureServicePrincipalToPb(st.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalPb != nil {
		pb.AzureServicePrincipal = azureServicePrincipalPb
	}

	cloudflareApiTokenPb, err := cloudflareApiTokenToPb(st.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenPb != nil {
		pb.CloudflareApiToken = cloudflareApiTokenPb
	}

	databricksGcpServiceAccountPb, err := databricksGcpServiceAccountRequestToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}

	pb.ExternalLocationName = st.ExternalLocationName

	pb.ReadOnly = st.ReadOnly

	pb.StorageCredentialName = st.StorageCredentialName

	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ValidateStorageCredential) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &validateStorageCredentialPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := validateStorageCredentialFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ValidateStorageCredential) MarshalJSON() ([]byte, error) {
	pb, err := validateStorageCredentialToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type validateStorageCredentialPb struct {
	// The AWS IAM role configuration.
	AwsIamRole *awsIamRoleRequestPb `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	AzureManagedIdentity *azureManagedIdentityRequestPb `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	AzureServicePrincipal *azureServicePrincipalPb `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	CloudflareApiToken *cloudflareApiTokenPb `json:"cloudflare_api_token,omitempty"`
	// The Databricks created GCP service account configuration.
	DatabricksGcpServiceAccount *databricksGcpServiceAccountRequestPb `json:"databricks_gcp_service_account,omitempty"`
	// The name of an existing external location to validate.
	ExternalLocationName string `json:"external_location_name,omitempty"`
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool `json:"read_only,omitempty"`
	// The name of the storage credential to validate.
	StorageCredentialName string `json:"storage_credential_name,omitempty"`
	// The external location url to validate.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func validateStorageCredentialFromPb(pb *validateStorageCredentialPb) (*ValidateStorageCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ValidateStorageCredential{}
	awsIamRoleField, err := awsIamRoleRequestFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := azureManagedIdentityRequestFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	azureServicePrincipalField, err := azureServicePrincipalFromPb(pb.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalField != nil {
		st.AzureServicePrincipal = azureServicePrincipalField
	}
	cloudflareApiTokenField, err := cloudflareApiTokenFromPb(pb.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenField != nil {
		st.CloudflareApiToken = cloudflareApiTokenField
	}
	databricksGcpServiceAccountField, err := databricksGcpServiceAccountRequestFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	st.ExternalLocationName = pb.ExternalLocationName
	st.ReadOnly = pb.ReadOnly
	st.StorageCredentialName = pb.StorageCredentialName
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *validateStorageCredentialPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st validateStorageCredentialPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ValidateStorageCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage.
	// Wire name: 'isDir'
	IsDir bool
	// The results of the validation check.
	// Wire name: 'results'
	Results []ValidationResult

	ForceSendFields []string `tf:"-"`
}

func validateStorageCredentialResponseToPb(st *ValidateStorageCredentialResponse) (*validateStorageCredentialResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &validateStorageCredentialResponsePb{}
	pb.IsDir = st.IsDir

	var resultsPb []validationResultPb
	for _, item := range st.Results {
		itemPb, err := validationResultToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resultsPb = append(resultsPb, *itemPb)
		}
	}
	pb.Results = resultsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ValidateStorageCredentialResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &validateStorageCredentialResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := validateStorageCredentialResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ValidateStorageCredentialResponse) MarshalJSON() ([]byte, error) {
	pb, err := validateStorageCredentialResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type validateStorageCredentialResponsePb struct {
	// Whether the tested location is a directory in cloud storage.
	IsDir bool `json:"isDir,omitempty"`
	// The results of the validation check.
	Results []validationResultPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func validateStorageCredentialResponseFromPb(pb *validateStorageCredentialResponsePb) (*ValidateStorageCredentialResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ValidateStorageCredentialResponse{}
	st.IsDir = pb.IsDir

	var resultsField []ValidationResult
	for _, itemPb := range pb.Results {
		item, err := validationResultFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resultsField = append(resultsField, *item)
		}
	}
	st.Results = resultsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *validateStorageCredentialResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st validateStorageCredentialResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	// Wire name: 'message'
	Message string
	// The operation tested.
	// Wire name: 'operation'
	Operation ValidationResultOperation
	// The results of the tested operation.
	// Wire name: 'result'
	Result ValidationResultResult

	ForceSendFields []string `tf:"-"`
}

func validationResultToPb(st *ValidationResult) (*validationResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &validationResultPb{}
	pb.Message = st.Message

	pb.Operation = st.Operation

	pb.Result = st.Result

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ValidationResult) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &validationResultPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := validationResultFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ValidationResult) MarshalJSON() ([]byte, error) {
	pb, err := validationResultToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type validationResultPb struct {
	// Error message would exist when the result does not equal to **PASS**.
	Message string `json:"message,omitempty"`
	// The operation tested.
	Operation ValidationResultOperation `json:"operation,omitempty"`
	// The results of the tested operation.
	Result ValidationResultResult `json:"result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func validationResultFromPb(pb *validationResultPb) (*ValidationResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ValidationResult{}
	st.Message = pb.Message
	st.Operation = pb.Operation
	st.Result = pb.Result

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *validationResultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st validationResultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The operation tested.
type ValidationResultOperation string
type validationResultOperationPb string

const ValidationResultOperationDelete ValidationResultOperation = `DELETE`

const ValidationResultOperationList ValidationResultOperation = `LIST`

const ValidationResultOperationPathExists ValidationResultOperation = `PATH_EXISTS`

const ValidationResultOperationRead ValidationResultOperation = `READ`

const ValidationResultOperationWrite ValidationResultOperation = `WRITE`

// String representation for [fmt.Print]
func (f *ValidationResultOperation) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ValidationResultOperation) Set(v string) error {
	switch v {
	case `DELETE`, `LIST`, `PATH_EXISTS`, `READ`, `WRITE`:
		*f = ValidationResultOperation(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETE", "LIST", "PATH_EXISTS", "READ", "WRITE"`, v)
	}
}

// Type always returns ValidationResultOperation to satisfy [pflag.Value] interface
func (f *ValidationResultOperation) Type() string {
	return "ValidationResultOperation"
}

func validationResultOperationToPb(st *ValidationResultOperation) (*validationResultOperationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := validationResultOperationPb(*st)
	return &pb, nil
}

func validationResultOperationFromPb(pb *validationResultOperationPb) (*ValidationResultOperation, error) {
	if pb == nil {
		return nil, nil
	}
	st := ValidationResultOperation(*pb)
	return &st, nil
}

// The results of the tested operation.
type ValidationResultResult string
type validationResultResultPb string

const ValidationResultResultFail ValidationResultResult = `FAIL`

const ValidationResultResultPass ValidationResultResult = `PASS`

const ValidationResultResultSkip ValidationResultResult = `SKIP`

// String representation for [fmt.Print]
func (f *ValidationResultResult) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ValidationResultResult) Set(v string) error {
	switch v {
	case `FAIL`, `PASS`, `SKIP`:
		*f = ValidationResultResult(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAIL", "PASS", "SKIP"`, v)
	}
}

// Type always returns ValidationResultResult to satisfy [pflag.Value] interface
func (f *ValidationResultResult) Type() string {
	return "ValidationResultResult"
}

func validationResultResultToPb(st *ValidationResultResult) (*validationResultResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := validationResultResultPb(*st)
	return &pb, nil
}

func validationResultResultFromPb(pb *validationResultResultPb) (*ValidationResultResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := ValidationResultResult(*pb)
	return &st, nil
}

type VolumeInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	// Wire name: 'access_point'
	AccessPoint string
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool
	// The name of the catalog where the schema and the volume are
	// Wire name: 'catalog_name'
	CatalogName string
	// The comment attached to the volume
	// Wire name: 'comment'
	Comment string

	// Wire name: 'created_at'
	CreatedAt int64
	// The identifier of the user who created the volume
	// Wire name: 'created_by'
	CreatedBy string
	// Encryption options that apply to clients connecting to cloud storage.
	// Wire name: 'encryption_details'
	EncryptionDetails *EncryptionDetails
	// The three-level (fully qualified) name of the volume
	// Wire name: 'full_name'
	FullName string
	// The unique identifier of the metastore
	// Wire name: 'metastore_id'
	MetastoreId string
	// The name of the volume
	// Wire name: 'name'
	Name string
	// The identifier of the user who owns the volume
	// Wire name: 'owner'
	Owner string
	// The name of the schema where the volume is
	// Wire name: 'schema_name'
	SchemaName string
	// The storage location on the cloud
	// Wire name: 'storage_location'
	StorageLocation string

	// Wire name: 'updated_at'
	UpdatedAt int64
	// The identifier of the user who updated the volume last time
	// Wire name: 'updated_by'
	UpdatedBy string
	// The unique identifier of the volume
	// Wire name: 'volume_id'
	VolumeId string
	// The type of the volume. An external volume is located in the specified
	// external location. A managed volume is located in the default location
	// which is specified by the parent schema, or the parent catalog, or the
	// Metastore. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/aws/en/volumes/managed-vs-external
	// Wire name: 'volume_type'
	VolumeType VolumeType

	ForceSendFields []string `tf:"-"`
}

func volumeInfoToPb(st *VolumeInfo) (*volumeInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &volumeInfoPb{}
	pb.AccessPoint = st.AccessPoint

	pb.BrowseOnly = st.BrowseOnly

	pb.CatalogName = st.CatalogName

	pb.Comment = st.Comment

	pb.CreatedAt = st.CreatedAt

	pb.CreatedBy = st.CreatedBy

	encryptionDetailsPb, err := encryptionDetailsToPb(st.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsPb != nil {
		pb.EncryptionDetails = encryptionDetailsPb
	}

	pb.FullName = st.FullName

	pb.MetastoreId = st.MetastoreId

	pb.Name = st.Name

	pb.Owner = st.Owner

	pb.SchemaName = st.SchemaName

	pb.StorageLocation = st.StorageLocation

	pb.UpdatedAt = st.UpdatedAt

	pb.UpdatedBy = st.UpdatedBy

	pb.VolumeId = st.VolumeId

	pb.VolumeType = st.VolumeType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *VolumeInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &volumeInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := volumeInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st VolumeInfo) MarshalJSON() ([]byte, error) {
	pb, err := volumeInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type volumeInfoPb struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string `json:"access_point,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The name of the catalog where the schema and the volume are
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the volume
	Comment string `json:"comment,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the volume
	CreatedBy string `json:"created_by,omitempty"`
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *encryptionDetailsPb `json:"encryption_details,omitempty"`
	// The three-level (fully qualified) name of the volume
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the metastore
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the volume
	Name string `json:"name,omitempty"`
	// The identifier of the user who owns the volume
	Owner string `json:"owner,omitempty"`
	// The name of the schema where the volume is
	SchemaName string `json:"schema_name,omitempty"`
	// The storage location on the cloud
	StorageLocation string `json:"storage_location,omitempty"`

	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the volume last time
	UpdatedBy string `json:"updated_by,omitempty"`
	// The unique identifier of the volume
	VolumeId string `json:"volume_id,omitempty"`
	// The type of the volume. An external volume is located in the specified
	// external location. A managed volume is located in the default location
	// which is specified by the parent schema, or the parent catalog, or the
	// Metastore. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/aws/en/volumes/managed-vs-external
	VolumeType VolumeType `json:"volume_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func volumeInfoFromPb(pb *volumeInfoPb) (*VolumeInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &VolumeInfo{}
	st.AccessPoint = pb.AccessPoint
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	encryptionDetailsField, err := encryptionDetailsFromPb(pb.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsField != nil {
		st.EncryptionDetails = encryptionDetailsField
	}
	st.FullName = pb.FullName
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.SchemaName = pb.SchemaName
	st.StorageLocation = pb.StorageLocation
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.VolumeId = pb.VolumeId
	st.VolumeType = pb.VolumeType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *volumeInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st volumeInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The type of the volume. An external volume is located in the specified
// external location. A managed volume is located in the default location which
// is specified by the parent schema, or the parent catalog, or the Metastore.
// [Learn more]
//
// [Learn more]: https://docs.databricks.com/aws/en/volumes/managed-vs-external
type VolumeType string
type volumeTypePb string

const VolumeTypeExternal VolumeType = `EXTERNAL`

const VolumeTypeManaged VolumeType = `MANAGED`

// String representation for [fmt.Print]
func (f *VolumeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *VolumeType) Set(v string) error {
	switch v {
	case `EXTERNAL`, `MANAGED`:
		*f = VolumeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "MANAGED"`, v)
	}
}

// Type always returns VolumeType to satisfy [pflag.Value] interface
func (f *VolumeType) Type() string {
	return "VolumeType"
}

func volumeTypeToPb(st *VolumeType) (*volumeTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := volumeTypePb(*st)
	return &pb, nil
}

func volumeTypeFromPb(pb *volumeTypePb) (*VolumeType, error) {
	if pb == nil {
		return nil, nil
	}
	st := VolumeType(*pb)
	return &st, nil
}

type WorkspaceBinding struct {

	// Wire name: 'binding_type'
	BindingType WorkspaceBindingBindingType

	// Wire name: 'workspace_id'
	WorkspaceId int64

	ForceSendFields []string `tf:"-"`
}

func workspaceBindingToPb(st *WorkspaceBinding) (*workspaceBindingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceBindingPb{}
	pb.BindingType = st.BindingType

	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *WorkspaceBinding) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspaceBindingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := workspaceBindingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WorkspaceBinding) MarshalJSON() ([]byte, error) {
	pb, err := workspaceBindingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type workspaceBindingPb struct {
	BindingType WorkspaceBindingBindingType `json:"binding_type,omitempty"`

	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func workspaceBindingFromPb(pb *workspaceBindingPb) (*WorkspaceBinding, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceBinding{}
	st.BindingType = pb.BindingType
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *workspaceBindingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st workspaceBindingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WorkspaceBindingBindingType string
type workspaceBindingBindingTypePb string

const WorkspaceBindingBindingTypeBindingTypeReadOnly WorkspaceBindingBindingType = `BINDING_TYPE_READ_ONLY`

const WorkspaceBindingBindingTypeBindingTypeReadWrite WorkspaceBindingBindingType = `BINDING_TYPE_READ_WRITE`

// String representation for [fmt.Print]
func (f *WorkspaceBindingBindingType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspaceBindingBindingType) Set(v string) error {
	switch v {
	case `BINDING_TYPE_READ_ONLY`, `BINDING_TYPE_READ_WRITE`:
		*f = WorkspaceBindingBindingType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BINDING_TYPE_READ_ONLY", "BINDING_TYPE_READ_WRITE"`, v)
	}
}

// Type always returns WorkspaceBindingBindingType to satisfy [pflag.Value] interface
func (f *WorkspaceBindingBindingType) Type() string {
	return "WorkspaceBindingBindingType"
}

func workspaceBindingBindingTypeToPb(st *WorkspaceBindingBindingType) (*workspaceBindingBindingTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := workspaceBindingBindingTypePb(*st)
	return &pb, nil
}

func workspaceBindingBindingTypeFromPb(pb *workspaceBindingBindingTypePb) (*WorkspaceBindingBindingType, error) {
	if pb == nil {
		return nil, nil
	}
	st := WorkspaceBindingBindingType(*pb)
	return &st, nil
}

// Currently assigned workspace bindings
type WorkspaceBindingsResponse struct {
	// List of workspace bindings
	// Wire name: 'bindings'
	Bindings []WorkspaceBinding
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func workspaceBindingsResponseToPb(st *WorkspaceBindingsResponse) (*workspaceBindingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceBindingsResponsePb{}

	var bindingsPb []workspaceBindingPb
	for _, item := range st.Bindings {
		itemPb, err := workspaceBindingToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			bindingsPb = append(bindingsPb, *itemPb)
		}
	}
	pb.Bindings = bindingsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *WorkspaceBindingsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspaceBindingsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := workspaceBindingsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WorkspaceBindingsResponse) MarshalJSON() ([]byte, error) {
	pb, err := workspaceBindingsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type workspaceBindingsResponsePb struct {
	// List of workspace bindings
	Bindings []workspaceBindingPb `json:"bindings,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func workspaceBindingsResponseFromPb(pb *workspaceBindingsResponsePb) (*WorkspaceBindingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceBindingsResponse{}

	var bindingsField []WorkspaceBinding
	for _, itemPb := range pb.Bindings {
		item, err := workspaceBindingFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			bindingsField = append(bindingsField, *item)
		}
	}
	st.Bindings = bindingsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *workspaceBindingsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st workspaceBindingsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

func durationFromPb(s *string) (*time.Duration, error) {
	if s == nil {
		return nil, nil
	}
	d, err := time.ParseDuration(*s)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func timestampToPb(t *time.Time) (*string, error) {
	if t == nil {
		return nil, nil
	}
	s := t.Format(time.RFC3339)
	return &s, nil
}

func timestampFromPb(s *string) (*time.Time, error) {
	if s == nil {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func fieldMaskToPb(fm *[]string) (*string, error) {
	if fm == nil {
		return nil, nil
	}
	s := strings.Join(*fm, ",")
	return &s, nil
}

func fieldMaskFromPb(s *string) (*[]string, error) {
	if s == nil {
		return nil, nil
	}
	fm := strings.Split(*s, ",")
	return &fm, nil
}
