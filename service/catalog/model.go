// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func identity[T any](obj *T) (*T, error) {
	return obj, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

// Helper to strip trailing zeros in fractional part
func rstripZeros(s string) string {
	for len(s) > 0 && s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}
	if len(s) > 0 && s[len(s)-1] == '.' {
		s = s[:len(s)-1]
	}
	return s
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

type AccountsCreateMetastore struct {
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
	MetastoreAssignment *CreateMetastoreAssignment
	// Unity Catalog metastore ID
	MetastoreId string
	// Workspace ID.
	WorkspaceId int64
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

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	workspaceIdPb := &st.WorkspaceId
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

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
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	workspaceIdField := &pb.WorkspaceId
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}

	return st, nil
}

type AccountsCreateStorageCredential struct {
	CredentialInfo *CreateStorageCredential
	// Unity Catalog metastore ID
	MetastoreId string
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

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

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
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}

	return st, nil
}

type AccountsMetastoreAssignment struct {
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
	MetastoreId string

	MetastoreInfo *UpdateMetastore
}

func accountsUpdateMetastoreToPb(st *AccountsUpdateMetastore) (*accountsUpdateMetastorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsUpdateMetastorePb{}
	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

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
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
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
	MetastoreAssignment *UpdateMetastoreAssignment
	// Unity Catalog metastore ID
	MetastoreId string
	// Workspace ID.
	WorkspaceId int64
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

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	workspaceIdPb := &st.WorkspaceId
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

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
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	workspaceIdField := &pb.WorkspaceId
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}

	return st, nil
}

type AccountsUpdateStorageCredential struct {
	CredentialInfo *UpdateStorageCredential
	// Unity Catalog metastore ID
	MetastoreId string
	// Name of the storage credential.
	StorageCredentialName string
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

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	storageCredentialNamePb := &st.StorageCredentialName
	if storageCredentialNamePb != nil {
		pb.StorageCredentialName = *storageCredentialNamePb
	}

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
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	storageCredentialNameField := &pb.StorageCredentialName
	if storageCredentialNameField != nil {
		st.StorageCredentialName = *storageCredentialNameField
	}

	return st, nil
}

type ArtifactAllowlistInfo struct {
	// A list of allowed artifact match patterns.
	ArtifactMatchers []ArtifactMatcher
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	CreatedAt int64
	// Username of the user who set the artifact allowlist.
	CreatedBy string
	// Unique identifier of parent metastore.
	MetastoreId string

	ForceSendFields []string
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

	createdAtPb := &st.CreatedAt
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

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
	for _, item := range pb.ArtifactMatchers {
		itemField, err := artifactMatcherFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			artifactMatchersField = append(artifactMatchersField, *itemField)
		}
	}
	st.ArtifactMatchers = artifactMatchersField
	createdAtField := &pb.CreatedAt
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}

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
	Artifact string
	// The pattern matching type of the artifact
	MatchType MatchType
}

func artifactMatcherToPb(st *ArtifactMatcher) (*artifactMatcherPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &artifactMatcherPb{}
	artifactPb := &st.Artifact
	if artifactPb != nil {
		pb.Artifact = *artifactPb
	}

	matchTypePb := &st.MatchType
	if matchTypePb != nil {
		pb.MatchType = *matchTypePb
	}

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
	artifactField := &pb.Artifact
	if artifactField != nil {
		st.Artifact = *artifactField
	}
	matchTypeField := &pb.MatchType
	if matchTypeField != nil {
		st.MatchType = *matchTypeField
	}

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
	AccessKeyId string
	// The Amazon Resource Name (ARN) of the S3 access point for temporary
	// credentials related the external location.
	AccessPoint string
	// The secret access key that can be used to sign AWS API requests.
	SecretAccessKey string
	// The token that users must pass to AWS API to use the temporary
	// credentials.
	SessionToken string

	ForceSendFields []string
}

func awsCredentialsToPb(st *AwsCredentials) (*awsCredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsCredentialsPb{}
	accessKeyIdPb := &st.AccessKeyId
	if accessKeyIdPb != nil {
		pb.AccessKeyId = *accessKeyIdPb
	}

	accessPointPb := &st.AccessPoint
	if accessPointPb != nil {
		pb.AccessPoint = *accessPointPb
	}

	secretAccessKeyPb := &st.SecretAccessKey
	if secretAccessKeyPb != nil {
		pb.SecretAccessKey = *secretAccessKeyPb
	}

	sessionTokenPb := &st.SessionToken
	if sessionTokenPb != nil {
		pb.SessionToken = *sessionTokenPb
	}

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
	accessKeyIdField := &pb.AccessKeyId
	if accessKeyIdField != nil {
		st.AccessKeyId = *accessKeyIdField
	}
	accessPointField := &pb.AccessPoint
	if accessPointField != nil {
		st.AccessPoint = *accessPointField
	}
	secretAccessKeyField := &pb.SecretAccessKey
	if secretAccessKeyField != nil {
		st.SecretAccessKey = *secretAccessKeyField
	}
	sessionTokenField := &pb.SessionToken
	if sessionTokenField != nil {
		st.SessionToken = *sessionTokenField
	}

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
	ExternalId string
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	RoleArn string
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn string

	ForceSendFields []string
}

func awsIamRoleToPb(st *AwsIamRole) (*awsIamRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsIamRolePb{}
	externalIdPb := &st.ExternalId
	if externalIdPb != nil {
		pb.ExternalId = *externalIdPb
	}

	roleArnPb := &st.RoleArn
	if roleArnPb != nil {
		pb.RoleArn = *roleArnPb
	}

	unityCatalogIamArnPb := &st.UnityCatalogIamArn
	if unityCatalogIamArnPb != nil {
		pb.UnityCatalogIamArn = *unityCatalogIamArnPb
	}

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
	externalIdField := &pb.ExternalId
	if externalIdField != nil {
		st.ExternalId = *externalIdField
	}
	roleArnField := &pb.RoleArn
	if roleArnField != nil {
		st.RoleArn = *roleArnField
	}
	unityCatalogIamArnField := &pb.UnityCatalogIamArn
	if unityCatalogIamArnField != nil {
		st.UnityCatalogIamArn = *unityCatalogIamArnField
	}

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
	RoleArn string
}

func awsIamRoleRequestToPb(st *AwsIamRoleRequest) (*awsIamRoleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsIamRoleRequestPb{}
	roleArnPb := &st.RoleArn
	if roleArnPb != nil {
		pb.RoleArn = *roleArnPb
	}

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
	roleArnField := &pb.RoleArn
	if roleArnField != nil {
		st.RoleArn = *roleArnField
	}

	return st, nil
}

type AwsIamRoleResponse struct {
	// The external ID used in role assumption to prevent confused deputy
	// problem..
	ExternalId string
	// The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access.
	RoleArn string
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	UnityCatalogIamArn string

	ForceSendFields []string
}

func awsIamRoleResponseToPb(st *AwsIamRoleResponse) (*awsIamRoleResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsIamRoleResponsePb{}
	externalIdPb := &st.ExternalId
	if externalIdPb != nil {
		pb.ExternalId = *externalIdPb
	}

	roleArnPb := &st.RoleArn
	if roleArnPb != nil {
		pb.RoleArn = *roleArnPb
	}

	unityCatalogIamArnPb := &st.UnityCatalogIamArn
	if unityCatalogIamArnPb != nil {
		pb.UnityCatalogIamArn = *unityCatalogIamArnPb
	}

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
	externalIdField := &pb.ExternalId
	if externalIdField != nil {
		st.ExternalId = *externalIdField
	}
	roleArnField := &pb.RoleArn
	if roleArnField != nil {
		st.RoleArn = *roleArnField
	}
	unityCatalogIamArnField := &pb.UnityCatalogIamArn
	if unityCatalogIamArnField != nil {
		st.UnityCatalogIamArn = *unityCatalogIamArnField
	}

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
	AadToken string

	ForceSendFields []string
}

func azureActiveDirectoryTokenToPb(st *AzureActiveDirectoryToken) (*azureActiveDirectoryTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureActiveDirectoryTokenPb{}
	aadTokenPb := &st.AadToken
	if aadTokenPb != nil {
		pb.AadToken = *aadTokenPb
	}

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
	aadTokenField := &pb.AadToken
	if aadTokenField != nil {
		st.AadToken = *aadTokenField
	}

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
	AccessConnectorId string
	// The Databricks internal ID that represents this managed identity. This
	// field is only used to persist the credential_id once it is fetched from
	// the credentials manager - as we only use the protobuf serializer to store
	// credentials, this ID gets persisted to the database. .
	CredentialId string
	// The Azure resource ID of the managed identity. Use the format,
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}`
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// using the system-assigned identity.
	ManagedIdentityId string

	ForceSendFields []string
}

func azureManagedIdentityToPb(st *AzureManagedIdentity) (*azureManagedIdentityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureManagedIdentityPb{}
	accessConnectorIdPb := &st.AccessConnectorId
	if accessConnectorIdPb != nil {
		pb.AccessConnectorId = *accessConnectorIdPb
	}

	credentialIdPb := &st.CredentialId
	if credentialIdPb != nil {
		pb.CredentialId = *credentialIdPb
	}

	managedIdentityIdPb := &st.ManagedIdentityId
	if managedIdentityIdPb != nil {
		pb.ManagedIdentityId = *managedIdentityIdPb
	}

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
	accessConnectorIdField := &pb.AccessConnectorId
	if accessConnectorIdField != nil {
		st.AccessConnectorId = *accessConnectorIdField
	}
	credentialIdField := &pb.CredentialId
	if credentialIdField != nil {
		st.CredentialId = *credentialIdField
	}
	managedIdentityIdField := &pb.ManagedIdentityId
	if managedIdentityIdField != nil {
		st.ManagedIdentityId = *managedIdentityIdField
	}

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
	AccessConnectorId string
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	ManagedIdentityId string

	ForceSendFields []string
}

func azureManagedIdentityRequestToPb(st *AzureManagedIdentityRequest) (*azureManagedIdentityRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureManagedIdentityRequestPb{}
	accessConnectorIdPb := &st.AccessConnectorId
	if accessConnectorIdPb != nil {
		pb.AccessConnectorId = *accessConnectorIdPb
	}

	managedIdentityIdPb := &st.ManagedIdentityId
	if managedIdentityIdPb != nil {
		pb.ManagedIdentityId = *managedIdentityIdPb
	}

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
	accessConnectorIdField := &pb.AccessConnectorId
	if accessConnectorIdField != nil {
		st.AccessConnectorId = *accessConnectorIdField
	}
	managedIdentityIdField := &pb.ManagedIdentityId
	if managedIdentityIdField != nil {
		st.ManagedIdentityId = *managedIdentityIdField
	}

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
	AccessConnectorId string
	// The Databricks internal ID that represents this managed identity.
	CredentialId string
	// The Azure resource ID of the managed identity. Use the format
	// /subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}.
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// for a system-assigned identity.
	ManagedIdentityId string

	ForceSendFields []string
}

func azureManagedIdentityResponseToPb(st *AzureManagedIdentityResponse) (*azureManagedIdentityResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureManagedIdentityResponsePb{}
	accessConnectorIdPb := &st.AccessConnectorId
	if accessConnectorIdPb != nil {
		pb.AccessConnectorId = *accessConnectorIdPb
	}

	credentialIdPb := &st.CredentialId
	if credentialIdPb != nil {
		pb.CredentialId = *credentialIdPb
	}

	managedIdentityIdPb := &st.ManagedIdentityId
	if managedIdentityIdPb != nil {
		pb.ManagedIdentityId = *managedIdentityIdPb
	}

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
	accessConnectorIdField := &pb.AccessConnectorId
	if accessConnectorIdField != nil {
		st.AccessConnectorId = *accessConnectorIdField
	}
	credentialIdField := &pb.CredentialId
	if credentialIdField != nil {
		st.CredentialId = *credentialIdField
	}
	managedIdentityIdField := &pb.ManagedIdentityId
	if managedIdentityIdField != nil {
		st.ManagedIdentityId = *managedIdentityIdField
	}

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
	ApplicationId string
	// The client secret generated for the above app ID in AAD.
	ClientSecret string
	// The directory ID corresponding to the Azure Active Directory (AAD) tenant
	// of the application.
	DirectoryId string
}

func azureServicePrincipalToPb(st *AzureServicePrincipal) (*azureServicePrincipalPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureServicePrincipalPb{}
	applicationIdPb := &st.ApplicationId
	if applicationIdPb != nil {
		pb.ApplicationId = *applicationIdPb
	}

	clientSecretPb := &st.ClientSecret
	if clientSecretPb != nil {
		pb.ClientSecret = *clientSecretPb
	}

	directoryIdPb := &st.DirectoryId
	if directoryIdPb != nil {
		pb.DirectoryId = *directoryIdPb
	}

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
	applicationIdField := &pb.ApplicationId
	if applicationIdField != nil {
		st.ApplicationId = *applicationIdField
	}
	clientSecretField := &pb.ClientSecret
	if clientSecretField != nil {
		st.ClientSecret = *clientSecretField
	}
	directoryIdField := &pb.DirectoryId
	if directoryIdField != nil {
		st.DirectoryId = *directoryIdField
	}

	return st, nil
}

// Azure temporary credentials for API authentication. Read more at
// https://docs.microsoft.com/en-us/rest/api/storageservices/create-user-delegation-sas
type AzureUserDelegationSas struct {
	// The signed URI (SAS Token) used to access blob services for a given path
	SasToken string

	ForceSendFields []string
}

func azureUserDelegationSasToPb(st *AzureUserDelegationSas) (*azureUserDelegationSasPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureUserDelegationSasPb{}
	sasTokenPb := &st.SasToken
	if sasTokenPb != nil {
		pb.SasToken = *sasTokenPb
	}

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
	sasTokenField := &pb.SasToken
	if sasTokenField != nil {
		st.SasToken = *sasTokenField
	}

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
	RefreshId string
	// Full name of the table.
	TableName string
}

func cancelRefreshRequestToPb(st *CancelRefreshRequest) (*cancelRefreshRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelRefreshRequestPb{}
	refreshIdPb := &st.RefreshId
	if refreshIdPb != nil {
		pb.RefreshId = *refreshIdPb
	}

	tableNamePb := &st.TableName
	if tableNamePb != nil {
		pb.TableName = *tableNamePb
	}

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
	refreshIdField := &pb.RefreshId
	if refreshIdField != nil {
		st.RefreshId = *refreshIdField
	}
	tableNameField := &pb.TableName
	if tableNameField != nil {
		st.TableName = *tableNameField
	}

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
	BrowseOnly bool
	// The type of the catalog.
	CatalogType CatalogType
	// User-provided free-form text description.
	Comment string
	// The name of the connection to an external data source.
	ConnectionName string
	// Time at which this catalog was created, in epoch milliseconds.
	CreatedAt int64
	// Username of catalog creator.
	CreatedBy string

	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization
	// The full name of the catalog. Corresponds with the name field.
	FullName string
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode CatalogIsolationMode
	// Unique identifier of parent metastore.
	MetastoreId string
	// Name of catalog.
	Name string
	// A map of key-value properties attached to the securable.
	Options map[string]string
	// Username of current owner of catalog.
	Owner string
	// A map of key-value properties attached to the securable.
	Properties map[string]string
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName string
	// Status of an asynchronously provisioned resource.
	ProvisioningInfo *ProvisioningInfo

	SecurableType string
	// The name of the share under the share provider.
	ShareName string
	// Storage Location URL (full path) for managed tables within catalog.
	StorageLocation string
	// Storage root URL for managed tables within catalog.
	StorageRoot string
	// Time at which this catalog was last modified, in epoch milliseconds.
	UpdatedAt int64
	// Username of user who last modified catalog.
	UpdatedBy string

	ForceSendFields []string
}

func catalogInfoToPb(st *CatalogInfo) (*catalogInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogInfoPb{}
	browseOnlyPb := &st.BrowseOnly
	if browseOnlyPb != nil {
		pb.BrowseOnly = *browseOnlyPb
	}

	catalogTypePb := &st.CatalogType
	if catalogTypePb != nil {
		pb.CatalogType = *catalogTypePb
	}

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	connectionNamePb := &st.ConnectionName
	if connectionNamePb != nil {
		pb.ConnectionName = *connectionNamePb
	}

	createdAtPb := &st.CreatedAt
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	effectivePredictiveOptimizationFlagPb, err := effectivePredictiveOptimizationFlagToPb(st.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagPb != nil {
		pb.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagPb
	}

	enablePredictiveOptimizationPb := &st.EnablePredictiveOptimization
	if enablePredictiveOptimizationPb != nil {
		pb.EnablePredictiveOptimization = *enablePredictiveOptimizationPb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	isolationModePb := &st.IsolationMode
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	optionsPb := map[string]string{}
	for k, v := range st.Options {
		itemPb := &v
		if itemPb != nil {
			optionsPb[k] = *itemPb
		}
	}
	pb.Options = optionsPb

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	propertiesPb := map[string]string{}
	for k, v := range st.Properties {
		itemPb := &v
		if itemPb != nil {
			propertiesPb[k] = *itemPb
		}
	}
	pb.Properties = propertiesPb

	providerNamePb := &st.ProviderName
	if providerNamePb != nil {
		pb.ProviderName = *providerNamePb
	}

	provisioningInfoPb, err := provisioningInfoToPb(st.ProvisioningInfo)
	if err != nil {
		return nil, err
	}
	if provisioningInfoPb != nil {
		pb.ProvisioningInfo = provisioningInfoPb
	}

	securableTypePb := &st.SecurableType
	if securableTypePb != nil {
		pb.SecurableType = *securableTypePb
	}

	shareNamePb := &st.ShareName
	if shareNamePb != nil {
		pb.ShareName = *shareNamePb
	}

	storageLocationPb := &st.StorageLocation
	if storageLocationPb != nil {
		pb.StorageLocation = *storageLocationPb
	}

	storageRootPb := &st.StorageRoot
	if storageRootPb != nil {
		pb.StorageRoot = *storageRootPb
	}

	updatedAtPb := &st.UpdatedAt
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb := &st.UpdatedBy
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

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
	browseOnlyField := &pb.BrowseOnly
	if browseOnlyField != nil {
		st.BrowseOnly = *browseOnlyField
	}
	catalogTypeField := &pb.CatalogType
	if catalogTypeField != nil {
		st.CatalogType = *catalogTypeField
	}
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	connectionNameField := &pb.ConnectionName
	if connectionNameField != nil {
		st.ConnectionName = *connectionNameField
	}
	createdAtField := &pb.CreatedAt
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	effectivePredictiveOptimizationFlagField, err := effectivePredictiveOptimizationFlagFromPb(pb.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagField != nil {
		st.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagField
	}
	enablePredictiveOptimizationField := &pb.EnablePredictiveOptimization
	if enablePredictiveOptimizationField != nil {
		st.EnablePredictiveOptimization = *enablePredictiveOptimizationField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	isolationModeField := &pb.IsolationMode
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	optionsField := map[string]string{}
	for k, v := range pb.Options {
		itemField := &v
		if itemField != nil {
			optionsField[k] = *itemField
		}
	}
	st.Options = optionsField
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}

	propertiesField := map[string]string{}
	for k, v := range pb.Properties {
		itemField := &v
		if itemField != nil {
			propertiesField[k] = *itemField
		}
	}
	st.Properties = propertiesField
	providerNameField := &pb.ProviderName
	if providerNameField != nil {
		st.ProviderName = *providerNameField
	}
	provisioningInfoField, err := provisioningInfoFromPb(pb.ProvisioningInfo)
	if err != nil {
		return nil, err
	}
	if provisioningInfoField != nil {
		st.ProvisioningInfo = provisioningInfoField
	}
	securableTypeField := &pb.SecurableType
	if securableTypeField != nil {
		st.SecurableType = *securableTypeField
	}
	shareNameField := &pb.ShareName
	if shareNameField != nil {
		st.ShareName = *shareNameField
	}
	storageLocationField := &pb.StorageLocation
	if storageLocationField != nil {
		st.StorageLocation = *storageLocationField
	}
	storageRootField := &pb.StorageRoot
	if storageRootField != nil {
		st.StorageRoot = *storageRootField
	}
	updatedAtField := &pb.UpdatedAt
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField := &pb.UpdatedBy
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}

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
	AccessKeyId string
	// The account id associated with the API token.
	AccountId string
	// The secret access token generated for the access key id
	SecretAccessKey string
}

func cloudflareApiTokenToPb(st *CloudflareApiToken) (*cloudflareApiTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cloudflareApiTokenPb{}
	accessKeyIdPb := &st.AccessKeyId
	if accessKeyIdPb != nil {
		pb.AccessKeyId = *accessKeyIdPb
	}

	accountIdPb := &st.AccountId
	if accountIdPb != nil {
		pb.AccountId = *accountIdPb
	}

	secretAccessKeyPb := &st.SecretAccessKey
	if secretAccessKeyPb != nil {
		pb.SecretAccessKey = *secretAccessKeyPb
	}

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
	accessKeyIdField := &pb.AccessKeyId
	if accessKeyIdField != nil {
		st.AccessKeyId = *accessKeyIdField
	}
	accountIdField := &pb.AccountId
	if accountIdField != nil {
		st.AccountId = *accountIdField
	}
	secretAccessKeyField := &pb.SecretAccessKey
	if secretAccessKeyField != nil {
		st.SecretAccessKey = *secretAccessKeyField
	}

	return st, nil
}

type ColumnInfo struct {
	// User-provided free-form text description.
	Comment string

	Mask *ColumnMask
	// Name of Column.
	Name string
	// Whether field may be Null (default: true).
	Nullable bool
	// Partition index for column.
	PartitionIndex int
	// Ordinal position of column (starting at position 0).
	Position int
	// Format of IntervalType.
	TypeIntervalType string
	// Full data type specification, JSON-serialized.
	TypeJson string

	TypeName ColumnTypeName
	// Digits of precision; required for DecimalTypes.
	TypePrecision int
	// Digits to right of decimal; Required for DecimalTypes.
	TypeScale int
	// Full data type specification as SQL/catalogString text.
	TypeText string

	ForceSendFields []string
}

func ColumnInfoToPb(st *ColumnInfo) (*ColumnInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ColumnInfoPb{}
	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	maskPb, err := columnMaskToPb(st.Mask)
	if err != nil {
		return nil, err
	}
	if maskPb != nil {
		pb.Mask = maskPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	nullablePb := &st.Nullable
	if nullablePb != nil {
		pb.Nullable = *nullablePb
	}

	partitionIndexPb := &st.PartitionIndex
	if partitionIndexPb != nil {
		pb.PartitionIndex = *partitionIndexPb
	}

	positionPb := &st.Position
	if positionPb != nil {
		pb.Position = *positionPb
	}

	typeIntervalTypePb := &st.TypeIntervalType
	if typeIntervalTypePb != nil {
		pb.TypeIntervalType = *typeIntervalTypePb
	}

	typeJsonPb := &st.TypeJson
	if typeJsonPb != nil {
		pb.TypeJson = *typeJsonPb
	}

	typeNamePb := &st.TypeName
	if typeNamePb != nil {
		pb.TypeName = *typeNamePb
	}

	typePrecisionPb := &st.TypePrecision
	if typePrecisionPb != nil {
		pb.TypePrecision = *typePrecisionPb
	}

	typeScalePb := &st.TypeScale
	if typeScalePb != nil {
		pb.TypeScale = *typeScalePb
	}

	typeTextPb := &st.TypeText
	if typeTextPb != nil {
		pb.TypeText = *typeTextPb
	}

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	maskField, err := columnMaskFromPb(pb.Mask)
	if err != nil {
		return nil, err
	}
	if maskField != nil {
		st.Mask = maskField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	nullableField := &pb.Nullable
	if nullableField != nil {
		st.Nullable = *nullableField
	}
	partitionIndexField := &pb.PartitionIndex
	if partitionIndexField != nil {
		st.PartitionIndex = *partitionIndexField
	}
	positionField := &pb.Position
	if positionField != nil {
		st.Position = *positionField
	}
	typeIntervalTypeField := &pb.TypeIntervalType
	if typeIntervalTypeField != nil {
		st.TypeIntervalType = *typeIntervalTypeField
	}
	typeJsonField := &pb.TypeJson
	if typeJsonField != nil {
		st.TypeJson = *typeJsonField
	}
	typeNameField := &pb.TypeName
	if typeNameField != nil {
		st.TypeName = *typeNameField
	}
	typePrecisionField := &pb.TypePrecision
	if typePrecisionField != nil {
		st.TypePrecision = *typePrecisionField
	}
	typeScaleField := &pb.TypeScale
	if typeScaleField != nil {
		st.TypeScale = *typeScaleField
	}
	typeTextField := &pb.TypeText
	if typeTextField != nil {
		st.TypeText = *typeTextField
	}

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
	FunctionName string
	// The list of additional table columns to be passed as input to the column
	// mask function. The first arg of the mask function should be of the type
	// of the column being masked and the types of the rest of the args should
	// match the types of columns in 'using_column_names'.
	UsingColumnNames []string

	ForceSendFields []string
}

func columnMaskToPb(st *ColumnMask) (*columnMaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &columnMaskPb{}
	functionNamePb := &st.FunctionName
	if functionNamePb != nil {
		pb.FunctionName = *functionNamePb
	}

	var usingColumnNamesPb []string
	for _, item := range st.UsingColumnNames {
		itemPb := &item
		if itemPb != nil {
			usingColumnNamesPb = append(usingColumnNamesPb, *itemPb)
		}
	}
	pb.UsingColumnNames = usingColumnNamesPb

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
	functionNameField := &pb.FunctionName
	if functionNameField != nil {
		st.FunctionName = *functionNameField
	}

	var usingColumnNamesField []string
	for _, item := range pb.UsingColumnNames {
		itemField := &item
		if itemField != nil {
			usingColumnNamesField = append(usingColumnNamesField, *itemField)
		}
	}
	st.UsingColumnNames = usingColumnNamesField

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
	Comment string
	// Unique identifier of the Connection.
	ConnectionId string
	// The type of connection.
	ConnectionType ConnectionType
	// Time at which this connection was created, in epoch milliseconds.
	CreatedAt int64
	// Username of connection creator.
	CreatedBy string
	// The type of credential.
	CredentialType CredentialType
	// Full name of connection.
	FullName string
	// Unique identifier of parent metastore.
	MetastoreId string
	// Name of the connection.
	Name string
	// A map of key-value properties attached to the securable.
	Options map[string]string
	// Username of current owner of the connection.
	Owner string
	// An object containing map of key-value properties attached to the
	// connection.
	Properties map[string]string
	// Status of an asynchronously provisioned resource.
	ProvisioningInfo *ProvisioningInfo
	// If the connection is read only.
	ReadOnly bool

	SecurableType string
	// Time at which this connection was updated, in epoch milliseconds.
	UpdatedAt int64
	// Username of user who last modified connection.
	UpdatedBy string
	// URL of the remote data source, extracted from options.
	Url string

	ForceSendFields []string
}

func connectionInfoToPb(st *ConnectionInfo) (*connectionInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &connectionInfoPb{}
	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	connectionIdPb := &st.ConnectionId
	if connectionIdPb != nil {
		pb.ConnectionId = *connectionIdPb
	}

	connectionTypePb := &st.ConnectionType
	if connectionTypePb != nil {
		pb.ConnectionType = *connectionTypePb
	}

	createdAtPb := &st.CreatedAt
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	credentialTypePb := &st.CredentialType
	if credentialTypePb != nil {
		pb.CredentialType = *credentialTypePb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	optionsPb := map[string]string{}
	for k, v := range st.Options {
		itemPb := &v
		if itemPb != nil {
			optionsPb[k] = *itemPb
		}
	}
	pb.Options = optionsPb

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	propertiesPb := map[string]string{}
	for k, v := range st.Properties {
		itemPb := &v
		if itemPb != nil {
			propertiesPb[k] = *itemPb
		}
	}
	pb.Properties = propertiesPb

	provisioningInfoPb, err := provisioningInfoToPb(st.ProvisioningInfo)
	if err != nil {
		return nil, err
	}
	if provisioningInfoPb != nil {
		pb.ProvisioningInfo = provisioningInfoPb
	}

	readOnlyPb := &st.ReadOnly
	if readOnlyPb != nil {
		pb.ReadOnly = *readOnlyPb
	}

	securableTypePb := &st.SecurableType
	if securableTypePb != nil {
		pb.SecurableType = *securableTypePb
	}

	updatedAtPb := &st.UpdatedAt
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb := &st.UpdatedBy
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

	urlPb := &st.Url
	if urlPb != nil {
		pb.Url = *urlPb
	}

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	connectionIdField := &pb.ConnectionId
	if connectionIdField != nil {
		st.ConnectionId = *connectionIdField
	}
	connectionTypeField := &pb.ConnectionType
	if connectionTypeField != nil {
		st.ConnectionType = *connectionTypeField
	}
	createdAtField := &pb.CreatedAt
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	credentialTypeField := &pb.CredentialType
	if credentialTypeField != nil {
		st.CredentialType = *credentialTypeField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	optionsField := map[string]string{}
	for k, v := range pb.Options {
		itemField := &v
		if itemField != nil {
			optionsField[k] = *itemField
		}
	}
	st.Options = optionsField
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}

	propertiesField := map[string]string{}
	for k, v := range pb.Properties {
		itemField := &v
		if itemField != nil {
			propertiesField[k] = *itemField
		}
	}
	st.Properties = propertiesField
	provisioningInfoField, err := provisioningInfoFromPb(pb.ProvisioningInfo)
	if err != nil {
		return nil, err
	}
	if provisioningInfoField != nil {
		st.ProvisioningInfo = provisioningInfoField
	}
	readOnlyField := &pb.ReadOnly
	if readOnlyField != nil {
		st.ReadOnly = *readOnlyField
	}
	securableTypeField := &pb.SecurableType
	if securableTypeField != nil {
		st.SecurableType = *securableTypeField
	}
	updatedAtField := &pb.UpdatedAt
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField := &pb.UpdatedBy
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}
	urlField := &pb.Url
	if urlField != nil {
		st.Url = *urlField
	}

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
	InitialPipelineSyncProgress *PipelineProgress
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	LastProcessedCommitVersion int64
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp string

	ForceSendFields []string
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

	lastProcessedCommitVersionPb := &st.LastProcessedCommitVersion
	if lastProcessedCommitVersionPb != nil {
		pb.LastProcessedCommitVersion = *lastProcessedCommitVersionPb
	}

	timestampPb := &st.Timestamp
	if timestampPb != nil {
		pb.Timestamp = *timestampPb
	}

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
	lastProcessedCommitVersionField := &pb.LastProcessedCommitVersion
	if lastProcessedCommitVersionField != nil {
		st.LastProcessedCommitVersion = *lastProcessedCommitVersionField
	}
	timestampField := &pb.Timestamp
	if timestampField != nil {
		st.Timestamp = *timestampField
	}

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
	Comment string
	// The name of the connection to an external data source.
	ConnectionName string
	// Name of catalog.
	Name string
	// A map of key-value properties attached to the securable.
	Options map[string]string
	// A map of key-value properties attached to the securable.
	Properties map[string]string
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	ProviderName string
	// The name of the share under the share provider.
	ShareName string
	// Storage root URL for managed tables within catalog.
	StorageRoot string

	ForceSendFields []string
}

func createCatalogToPb(st *CreateCatalog) (*createCatalogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCatalogPb{}
	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	connectionNamePb := &st.ConnectionName
	if connectionNamePb != nil {
		pb.ConnectionName = *connectionNamePb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	optionsPb := map[string]string{}
	for k, v := range st.Options {
		itemPb := &v
		if itemPb != nil {
			optionsPb[k] = *itemPb
		}
	}
	pb.Options = optionsPb

	propertiesPb := map[string]string{}
	for k, v := range st.Properties {
		itemPb := &v
		if itemPb != nil {
			propertiesPb[k] = *itemPb
		}
	}
	pb.Properties = propertiesPb

	providerNamePb := &st.ProviderName
	if providerNamePb != nil {
		pb.ProviderName = *providerNamePb
	}

	shareNamePb := &st.ShareName
	if shareNamePb != nil {
		pb.ShareName = *shareNamePb
	}

	storageRootPb := &st.StorageRoot
	if storageRootPb != nil {
		pb.StorageRoot = *storageRootPb
	}

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	connectionNameField := &pb.ConnectionName
	if connectionNameField != nil {
		st.ConnectionName = *connectionNameField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	optionsField := map[string]string{}
	for k, v := range pb.Options {
		itemField := &v
		if itemField != nil {
			optionsField[k] = *itemField
		}
	}
	st.Options = optionsField

	propertiesField := map[string]string{}
	for k, v := range pb.Properties {
		itemField := &v
		if itemField != nil {
			propertiesField[k] = *itemField
		}
	}
	st.Properties = propertiesField
	providerNameField := &pb.ProviderName
	if providerNameField != nil {
		st.ProviderName = *providerNameField
	}
	shareNameField := &pb.ShareName
	if shareNameField != nil {
		st.ShareName = *shareNameField
	}
	storageRootField := &pb.StorageRoot
	if storageRootField != nil {
		st.StorageRoot = *storageRootField
	}

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
	Comment string
	// The type of connection.
	ConnectionType ConnectionType
	// Name of the connection.
	Name string
	// A map of key-value properties attached to the securable.
	Options map[string]string
	// An object containing map of key-value properties attached to the
	// connection.
	Properties map[string]string
	// If the connection is read only.
	ReadOnly bool

	ForceSendFields []string
}

func createConnectionToPb(st *CreateConnection) (*createConnectionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createConnectionPb{}
	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	connectionTypePb := &st.ConnectionType
	if connectionTypePb != nil {
		pb.ConnectionType = *connectionTypePb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	optionsPb := map[string]string{}
	for k, v := range st.Options {
		itemPb := &v
		if itemPb != nil {
			optionsPb[k] = *itemPb
		}
	}
	pb.Options = optionsPb

	propertiesPb := map[string]string{}
	for k, v := range st.Properties {
		itemPb := &v
		if itemPb != nil {
			propertiesPb[k] = *itemPb
		}
	}
	pb.Properties = propertiesPb

	readOnlyPb := &st.ReadOnly
	if readOnlyPb != nil {
		pb.ReadOnly = *readOnlyPb
	}

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	connectionTypeField := &pb.ConnectionType
	if connectionTypeField != nil {
		st.ConnectionType = *connectionTypeField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	optionsField := map[string]string{}
	for k, v := range pb.Options {
		itemField := &v
		if itemField != nil {
			optionsField[k] = *itemField
		}
	}
	st.Options = optionsField

	propertiesField := map[string]string{}
	for k, v := range pb.Properties {
		itemField := &v
		if itemField != nil {
			propertiesField[k] = *itemField
		}
	}
	st.Properties = propertiesField
	readOnlyField := &pb.ReadOnly
	if readOnlyField != nil {
		st.ReadOnly = *readOnlyField
	}

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
	AwsIamRole *AwsIamRole
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentity
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	AzureServicePrincipal *AzureServicePrincipal
	// Comment associated with the credential.
	Comment string
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	Name string
	// Indicates the purpose of the credential.
	Purpose CredentialPurpose
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool
	// Optional. Supplying true to this argument skips validation of the created
	// set of credentials.
	SkipValidation bool

	ForceSendFields []string
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

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	databricksGcpServiceAccountPb, err := databricksGcpServiceAccountToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	purposePb := &st.Purpose
	if purposePb != nil {
		pb.Purpose = *purposePb
	}

	readOnlyPb := &st.ReadOnly
	if readOnlyPb != nil {
		pb.ReadOnly = *readOnlyPb
	}

	skipValidationPb := &st.SkipValidation
	if skipValidationPb != nil {
		pb.SkipValidation = *skipValidationPb
	}

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	databricksGcpServiceAccountField, err := databricksGcpServiceAccountFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	purposeField := &pb.Purpose
	if purposeField != nil {
		st.Purpose = *purposeField
	}
	readOnlyField := &pb.ReadOnly
	if readOnlyField != nil {
		st.ReadOnly = *readOnlyField
	}
	skipValidationField := &pb.SkipValidation
	if skipValidationField != nil {
		st.SkipValidation = *skipValidationField
	}

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
	AccessPoint string
	// User-provided free-form text description.
	Comment string
	// Name of the storage credential used with this location.
	CredentialName string
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback bool
	// Name of the external location.
	Name string
	// Indicates whether the external location is read-only.
	ReadOnly bool
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation bool
	// Path URL of the external location.
	Url string

	ForceSendFields []string
}

func createExternalLocationToPb(st *CreateExternalLocation) (*createExternalLocationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createExternalLocationPb{}
	accessPointPb := &st.AccessPoint
	if accessPointPb != nil {
		pb.AccessPoint = *accessPointPb
	}

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	credentialNamePb := &st.CredentialName
	if credentialNamePb != nil {
		pb.CredentialName = *credentialNamePb
	}

	encryptionDetailsPb, err := encryptionDetailsToPb(st.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsPb != nil {
		pb.EncryptionDetails = encryptionDetailsPb
	}

	fallbackPb := &st.Fallback
	if fallbackPb != nil {
		pb.Fallback = *fallbackPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	readOnlyPb := &st.ReadOnly
	if readOnlyPb != nil {
		pb.ReadOnly = *readOnlyPb
	}

	skipValidationPb := &st.SkipValidation
	if skipValidationPb != nil {
		pb.SkipValidation = *skipValidationPb
	}

	urlPb := &st.Url
	if urlPb != nil {
		pb.Url = *urlPb
	}

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
	accessPointField := &pb.AccessPoint
	if accessPointField != nil {
		st.AccessPoint = *accessPointField
	}
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	credentialNameField := &pb.CredentialName
	if credentialNameField != nil {
		st.CredentialName = *credentialNameField
	}
	encryptionDetailsField, err := encryptionDetailsFromPb(pb.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsField != nil {
		st.EncryptionDetails = encryptionDetailsField
	}
	fallbackField := &pb.Fallback
	if fallbackField != nil {
		st.Fallback = *fallbackField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	readOnlyField := &pb.ReadOnly
	if readOnlyField != nil {
		st.ReadOnly = *readOnlyField
	}
	skipValidationField := &pb.SkipValidation
	if skipValidationField != nil {
		st.SkipValidation = *skipValidationField
	}
	urlField := &pb.Url
	if urlField != nil {
		st.Url = *urlField
	}

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
	CatalogName string
	// User-provided free-form text description.
	Comment string
	// Scalar function return data type.
	DataType ColumnTypeName
	// External function language.
	ExternalLanguage string
	// External function name.
	ExternalName string
	// Pretty printed function data type.
	FullDataType string

	InputParams FunctionParameterInfos
	// Whether the function is deterministic.
	IsDeterministic bool
	// Function null call.
	IsNullCall bool
	// Name of function, relative to parent schema.
	Name string
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle CreateFunctionParameterStyle
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties string
	// Table function return parameters.
	ReturnParams *FunctionParameterInfos
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody CreateFunctionRoutineBody
	// Function body.
	RoutineDefinition string
	// Function dependencies.
	RoutineDependencies *DependencyList
	// Name of parent schema relative to its parent catalog.
	SchemaName string
	// Function security type.
	SecurityType CreateFunctionSecurityType
	// Specific name of the function; Reserved for future use.
	SpecificName string
	// Function SQL data access.
	SqlDataAccess CreateFunctionSqlDataAccess
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string

	ForceSendFields []string
}

func createFunctionToPb(st *CreateFunction) (*createFunctionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createFunctionPb{}
	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	dataTypePb := &st.DataType
	if dataTypePb != nil {
		pb.DataType = *dataTypePb
	}

	externalLanguagePb := &st.ExternalLanguage
	if externalLanguagePb != nil {
		pb.ExternalLanguage = *externalLanguagePb
	}

	externalNamePb := &st.ExternalName
	if externalNamePb != nil {
		pb.ExternalName = *externalNamePb
	}

	fullDataTypePb := &st.FullDataType
	if fullDataTypePb != nil {
		pb.FullDataType = *fullDataTypePb
	}

	inputParamsPb, err := functionParameterInfosToPb(&st.InputParams)
	if err != nil {
		return nil, err
	}
	if inputParamsPb != nil {
		pb.InputParams = *inputParamsPb
	}

	isDeterministicPb := &st.IsDeterministic
	if isDeterministicPb != nil {
		pb.IsDeterministic = *isDeterministicPb
	}

	isNullCallPb := &st.IsNullCall
	if isNullCallPb != nil {
		pb.IsNullCall = *isNullCallPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	parameterStylePb := &st.ParameterStyle
	if parameterStylePb != nil {
		pb.ParameterStyle = *parameterStylePb
	}

	propertiesPb := &st.Properties
	if propertiesPb != nil {
		pb.Properties = *propertiesPb
	}

	returnParamsPb, err := functionParameterInfosToPb(st.ReturnParams)
	if err != nil {
		return nil, err
	}
	if returnParamsPb != nil {
		pb.ReturnParams = returnParamsPb
	}

	routineBodyPb := &st.RoutineBody
	if routineBodyPb != nil {
		pb.RoutineBody = *routineBodyPb
	}

	routineDefinitionPb := &st.RoutineDefinition
	if routineDefinitionPb != nil {
		pb.RoutineDefinition = *routineDefinitionPb
	}

	routineDependenciesPb, err := dependencyListToPb(st.RoutineDependencies)
	if err != nil {
		return nil, err
	}
	if routineDependenciesPb != nil {
		pb.RoutineDependencies = routineDependenciesPb
	}

	schemaNamePb := &st.SchemaName
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

	securityTypePb := &st.SecurityType
	if securityTypePb != nil {
		pb.SecurityType = *securityTypePb
	}

	specificNamePb := &st.SpecificName
	if specificNamePb != nil {
		pb.SpecificName = *specificNamePb
	}

	sqlDataAccessPb := &st.SqlDataAccess
	if sqlDataAccessPb != nil {
		pb.SqlDataAccess = *sqlDataAccessPb
	}

	sqlPathPb := &st.SqlPath
	if sqlPathPb != nil {
		pb.SqlPath = *sqlPathPb
	}

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
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	dataTypeField := &pb.DataType
	if dataTypeField != nil {
		st.DataType = *dataTypeField
	}
	externalLanguageField := &pb.ExternalLanguage
	if externalLanguageField != nil {
		st.ExternalLanguage = *externalLanguageField
	}
	externalNameField := &pb.ExternalName
	if externalNameField != nil {
		st.ExternalName = *externalNameField
	}
	fullDataTypeField := &pb.FullDataType
	if fullDataTypeField != nil {
		st.FullDataType = *fullDataTypeField
	}
	inputParamsField, err := functionParameterInfosFromPb(&pb.InputParams)
	if err != nil {
		return nil, err
	}
	if inputParamsField != nil {
		st.InputParams = *inputParamsField
	}
	isDeterministicField := &pb.IsDeterministic
	if isDeterministicField != nil {
		st.IsDeterministic = *isDeterministicField
	}
	isNullCallField := &pb.IsNullCall
	if isNullCallField != nil {
		st.IsNullCall = *isNullCallField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	parameterStyleField := &pb.ParameterStyle
	if parameterStyleField != nil {
		st.ParameterStyle = *parameterStyleField
	}
	propertiesField := &pb.Properties
	if propertiesField != nil {
		st.Properties = *propertiesField
	}
	returnParamsField, err := functionParameterInfosFromPb(pb.ReturnParams)
	if err != nil {
		return nil, err
	}
	if returnParamsField != nil {
		st.ReturnParams = returnParamsField
	}
	routineBodyField := &pb.RoutineBody
	if routineBodyField != nil {
		st.RoutineBody = *routineBodyField
	}
	routineDefinitionField := &pb.RoutineDefinition
	if routineDefinitionField != nil {
		st.RoutineDefinition = *routineDefinitionField
	}
	routineDependenciesField, err := dependencyListFromPb(pb.RoutineDependencies)
	if err != nil {
		return nil, err
	}
	if routineDependenciesField != nil {
		st.RoutineDependencies = routineDependenciesField
	}
	schemaNameField := &pb.SchemaName
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}
	securityTypeField := &pb.SecurityType
	if securityTypeField != nil {
		st.SecurityType = *securityTypeField
	}
	specificNameField := &pb.SpecificName
	if specificNameField != nil {
		st.SpecificName = *specificNameField
	}
	sqlDataAccessField := &pb.SqlDataAccess
	if sqlDataAccessField != nil {
		st.SqlDataAccess = *sqlDataAccessField
	}
	sqlPathField := &pb.SqlPath
	if sqlPathField != nil {
		st.SqlPath = *sqlPathField
	}

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
	Name string
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	// The field can be omitted in the __workspace-level__ __API__ but not in
	// the __account-level__ __API__. If this field is omitted, the region of
	// the workspace receiving the request will be used.
	Region string
	// The storage root URL for metastore
	StorageRoot string

	ForceSendFields []string
}

func createMetastoreToPb(st *CreateMetastore) (*createMetastorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createMetastorePb{}
	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	regionPb := &st.Region
	if regionPb != nil {
		pb.Region = *regionPb
	}

	storageRootPb := &st.StorageRoot
	if storageRootPb != nil {
		pb.StorageRoot = *storageRootPb
	}

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
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	regionField := &pb.Region
	if regionField != nil {
		st.Region = *regionField
	}
	storageRootField := &pb.StorageRoot
	if storageRootField != nil {
		st.StorageRoot = *storageRootField
	}

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
	DefaultCatalogName string
	// The unique ID of the metastore.
	MetastoreId string
	// A workspace ID.
	WorkspaceId int64
}

func createMetastoreAssignmentToPb(st *CreateMetastoreAssignment) (*createMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createMetastoreAssignmentPb{}
	defaultCatalogNamePb := &st.DefaultCatalogName
	if defaultCatalogNamePb != nil {
		pb.DefaultCatalogName = *defaultCatalogNamePb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	workspaceIdPb := &st.WorkspaceId
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

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
	defaultCatalogNameField := &pb.DefaultCatalogName
	if defaultCatalogNameField != nil {
		st.DefaultCatalogName = *defaultCatalogNameField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	workspaceIdField := &pb.WorkspaceId
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}

	return st, nil
}

type CreateMonitor struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	AssetsDir string
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName string
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics []MonitorMetric
	// The data classification config for the monitor.
	DataClassificationConfig *MonitorDataClassificationConfig
	// Configuration for monitoring inference logs.
	InferenceLog *MonitorInferenceLog
	// The notification settings for the monitor.
	Notifications *MonitorNotifications
	// Schema where output metric tables are created.
	OutputSchemaName string
	// The schedule for automatically updating and refreshing metric tables.
	Schedule *MonitorCronSchedule
	// Whether to skip creating a default dashboard summarizing data quality
	// metrics.
	SkipBuiltinDashboard bool
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs []string
	// Configuration for monitoring snapshot tables.
	Snapshot *MonitorSnapshot
	// Full name of the table.
	TableName string
	// Configuration for monitoring time series tables.
	TimeSeries *MonitorTimeSeries
	// Optional argument to specify the warehouse for dashboard creation. If not
	// specified, the first running warehouse will be used.
	WarehouseId string

	ForceSendFields []string
}

func createMonitorToPb(st *CreateMonitor) (*createMonitorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createMonitorPb{}
	assetsDirPb := &st.AssetsDir
	if assetsDirPb != nil {
		pb.AssetsDir = *assetsDirPb
	}

	baselineTableNamePb := &st.BaselineTableName
	if baselineTableNamePb != nil {
		pb.BaselineTableName = *baselineTableNamePb
	}

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

	outputSchemaNamePb := &st.OutputSchemaName
	if outputSchemaNamePb != nil {
		pb.OutputSchemaName = *outputSchemaNamePb
	}

	schedulePb, err := monitorCronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}

	skipBuiltinDashboardPb := &st.SkipBuiltinDashboard
	if skipBuiltinDashboardPb != nil {
		pb.SkipBuiltinDashboard = *skipBuiltinDashboardPb
	}

	var slicingExprsPb []string
	for _, item := range st.SlicingExprs {
		itemPb := &item
		if itemPb != nil {
			slicingExprsPb = append(slicingExprsPb, *itemPb)
		}
	}
	pb.SlicingExprs = slicingExprsPb

	snapshotPb, err := monitorSnapshotToPb(st.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotPb != nil {
		pb.Snapshot = snapshotPb
	}

	tableNamePb := &st.TableName
	if tableNamePb != nil {
		pb.TableName = *tableNamePb
	}

	timeSeriesPb, err := monitorTimeSeriesToPb(st.TimeSeries)
	if err != nil {
		return nil, err
	}
	if timeSeriesPb != nil {
		pb.TimeSeries = timeSeriesPb
	}

	warehouseIdPb := &st.WarehouseId
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

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
	assetsDirField := &pb.AssetsDir
	if assetsDirField != nil {
		st.AssetsDir = *assetsDirField
	}
	baselineTableNameField := &pb.BaselineTableName
	if baselineTableNameField != nil {
		st.BaselineTableName = *baselineTableNameField
	}

	var customMetricsField []MonitorMetric
	for _, item := range pb.CustomMetrics {
		itemField, err := monitorMetricFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			customMetricsField = append(customMetricsField, *itemField)
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
	outputSchemaNameField := &pb.OutputSchemaName
	if outputSchemaNameField != nil {
		st.OutputSchemaName = *outputSchemaNameField
	}
	scheduleField, err := monitorCronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	skipBuiltinDashboardField := &pb.SkipBuiltinDashboard
	if skipBuiltinDashboardField != nil {
		st.SkipBuiltinDashboard = *skipBuiltinDashboardField
	}

	var slicingExprsField []string
	for _, item := range pb.SlicingExprs {
		itemField := &item
		if itemField != nil {
			slicingExprsField = append(slicingExprsField, *itemField)
		}
	}
	st.SlicingExprs = slicingExprsField
	snapshotField, err := monitorSnapshotFromPb(pb.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotField != nil {
		st.Snapshot = snapshotField
	}
	tableNameField := &pb.TableName
	if tableNameField != nil {
		st.TableName = *tableNameField
	}
	timeSeriesField, err := monitorTimeSeriesFromPb(pb.TimeSeries)
	if err != nil {
		return nil, err
	}
	if timeSeriesField != nil {
		st.TimeSeries = timeSeriesField
	}
	warehouseIdField := &pb.WarehouseId
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

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
	CatalogName string
	// The comment attached to the registered model
	Comment string
	// The name of the registered model
	Name string
	// The name of the schema where the registered model resides
	SchemaName string
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string

	ForceSendFields []string
}

func createRegisteredModelRequestToPb(st *CreateRegisteredModelRequest) (*createRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createRegisteredModelRequestPb{}
	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	schemaNamePb := &st.SchemaName
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

	storageLocationPb := &st.StorageLocation
	if storageLocationPb != nil {
		pb.StorageLocation = *storageLocationPb
	}

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
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	schemaNameField := &pb.SchemaName
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}
	storageLocationField := &pb.StorageLocation
	if storageLocationField != nil {
		st.StorageLocation = *storageLocationField
	}

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
	CatalogName string
	// User-provided free-form text description.
	Comment string
	// Name of schema, relative to parent catalog.
	Name string
	// A map of key-value properties attached to the securable.
	Properties map[string]string
	// Storage root URL for managed tables within schema.
	StorageRoot string

	ForceSendFields []string
}

func createSchemaToPb(st *CreateSchema) (*createSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createSchemaPb{}
	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	propertiesPb := map[string]string{}
	for k, v := range st.Properties {
		itemPb := &v
		if itemPb != nil {
			propertiesPb[k] = *itemPb
		}
	}
	pb.Properties = propertiesPb

	storageRootPb := &st.StorageRoot
	if storageRootPb != nil {
		pb.StorageRoot = *storageRootPb
	}

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
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	propertiesField := map[string]string{}
	for k, v := range pb.Properties {
		itemField := &v
		if itemField != nil {
			propertiesField[k] = *itemField
		}
	}
	st.Properties = propertiesField
	storageRootField := &pb.StorageRoot
	if storageRootField != nil {
		st.StorageRoot = *storageRootField
	}

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
	AwsIamRole *AwsIamRoleRequest
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityRequest
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken
	// Comment associated with the credential.
	Comment string
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest
	// The credential name. The name must be unique within the metastore.
	Name string
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool
	// Supplying true to this argument skips validation of the created
	// credential.
	SkipValidation bool

	ForceSendFields []string
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

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	databricksGcpServiceAccountPb, err := databricksGcpServiceAccountRequestToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	readOnlyPb := &st.ReadOnly
	if readOnlyPb != nil {
		pb.ReadOnly = *readOnlyPb
	}

	skipValidationPb := &st.SkipValidation
	if skipValidationPb != nil {
		pb.SkipValidation = *skipValidationPb
	}

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	databricksGcpServiceAccountField, err := databricksGcpServiceAccountRequestFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	readOnlyField := &pb.ReadOnly
	if readOnlyField != nil {
		st.ReadOnly = *readOnlyField
	}
	skipValidationField := &pb.SkipValidation
	if skipValidationField != nil {
		st.SkipValidation = *skipValidationField
	}

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
	Constraint TableConstraint
	// The full name of the table referenced by the constraint.
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

	fullNameArgPb := &st.FullNameArg
	if fullNameArgPb != nil {
		pb.FullNameArg = *fullNameArgPb
	}

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
	fullNameArgField := &pb.FullNameArg
	if fullNameArgField != nil {
		st.FullNameArg = *fullNameArgField
	}

	return st, nil
}

type CreateVolumeRequestContent struct {
	// The name of the catalog where the schema and the volume are
	CatalogName string
	// The comment attached to the volume
	Comment string
	// The name of the volume
	Name string
	// The name of the schema where the volume is
	SchemaName string
	// The storage location on the cloud
	StorageLocation string
	// The type of the volume. An external volume is located in the specified
	// external location. A managed volume is located in the default location
	// which is specified by the parent schema, or the parent catalog, or the
	// Metastore. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/aws/en/volumes/managed-vs-external
	VolumeType VolumeType

	ForceSendFields []string
}

func createVolumeRequestContentToPb(st *CreateVolumeRequestContent) (*createVolumeRequestContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createVolumeRequestContentPb{}
	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	schemaNamePb := &st.SchemaName
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

	storageLocationPb := &st.StorageLocation
	if storageLocationPb != nil {
		pb.StorageLocation = *storageLocationPb
	}

	volumeTypePb := &st.VolumeType
	if volumeTypePb != nil {
		pb.VolumeType = *volumeTypePb
	}

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
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	schemaNameField := &pb.SchemaName
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}
	storageLocationField := &pb.StorageLocation
	if storageLocationField != nil {
		st.StorageLocation = *storageLocationField
	}
	volumeTypeField := &pb.VolumeType
	if volumeTypeField != nil {
		st.VolumeType = *volumeTypeField
	}

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
	AwsIamRole *AwsIamRole
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentity
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	AzureServicePrincipal *AzureServicePrincipal
	// Comment associated with the credential.
	Comment string
	// Time at which this credential was created, in epoch milliseconds.
	CreatedAt int64
	// Username of credential creator.
	CreatedBy string
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount
	// The full name of the credential.
	FullName string
	// The unique identifier of the credential.
	Id string
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode IsolationMode
	// Unique identifier of the parent metastore.
	MetastoreId string
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	Name string
	// Username of current owner of credential.
	Owner string
	// Indicates the purpose of the credential.
	Purpose CredentialPurpose
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt int64
	// Username of user who last modified the credential.
	UpdatedBy string
	// Whether this credential is the current metastore's root storage
	// credential. Only applicable when purpose is **STORAGE**.
	UsedForManagedStorage bool

	ForceSendFields []string
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

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	createdAtPb := &st.CreatedAt
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	databricksGcpServiceAccountPb, err := databricksGcpServiceAccountToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	idPb := &st.Id
	if idPb != nil {
		pb.Id = *idPb
	}

	isolationModePb := &st.IsolationMode
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	purposePb := &st.Purpose
	if purposePb != nil {
		pb.Purpose = *purposePb
	}

	readOnlyPb := &st.ReadOnly
	if readOnlyPb != nil {
		pb.ReadOnly = *readOnlyPb
	}

	updatedAtPb := &st.UpdatedAt
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb := &st.UpdatedBy
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

	usedForManagedStoragePb := &st.UsedForManagedStorage
	if usedForManagedStoragePb != nil {
		pb.UsedForManagedStorage = *usedForManagedStoragePb
	}

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	createdAtField := &pb.CreatedAt
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	databricksGcpServiceAccountField, err := databricksGcpServiceAccountFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	idField := &pb.Id
	if idField != nil {
		st.Id = *idField
	}
	isolationModeField := &pb.IsolationMode
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}
	purposeField := &pb.Purpose
	if purposeField != nil {
		st.Purpose = *purposeField
	}
	readOnlyField := &pb.ReadOnly
	if readOnlyField != nil {
		st.ReadOnly = *readOnlyField
	}
	updatedAtField := &pb.UpdatedAt
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField := &pb.UpdatedBy
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}
	usedForManagedStorageField := &pb.UsedForManagedStorage
	if usedForManagedStorageField != nil {
		st.UsedForManagedStorage = *usedForManagedStorageField
	}

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
	Message string
	// The results of the tested operation.
	Result ValidateCredentialResult

	ForceSendFields []string
}

func credentialValidationResultToPb(st *CredentialValidationResult) (*credentialValidationResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &credentialValidationResultPb{}
	messagePb := &st.Message
	if messagePb != nil {
		pb.Message = *messagePb
	}

	resultPb := &st.Result
	if resultPb != nil {
		pb.Result = *resultPb
	}

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
	messageField := &pb.Message
	if messageField != nil {
		st.Message = *messageField
	}
	resultField := &pb.Result
	if resultField != nil {
		st.Result = *resultField
	}

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
	Workspaces []int64
}

func currentWorkspaceBindingsToPb(st *CurrentWorkspaceBindings) (*currentWorkspaceBindingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &currentWorkspaceBindingsPb{}

	var workspacesPb []int64
	for _, item := range st.Workspaces {
		itemPb := &item
		if itemPb != nil {
			workspacesPb = append(workspacesPb, *itemPb)
		}
	}
	pb.Workspaces = workspacesPb

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

	var workspacesField []int64
	for _, item := range pb.Workspaces {
		itemField := &item
		if itemField != nil {
			workspacesField = append(workspacesField, *itemField)
		}
	}
	st.Workspaces = workspacesField

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
	CredentialId string
	// The email of the service account.
	Email string
	// The ID that represents the private key for this Service Account
	PrivateKeyId string

	ForceSendFields []string
}

func databricksGcpServiceAccountToPb(st *DatabricksGcpServiceAccount) (*databricksGcpServiceAccountPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databricksGcpServiceAccountPb{}
	credentialIdPb := &st.CredentialId
	if credentialIdPb != nil {
		pb.CredentialId = *credentialIdPb
	}

	emailPb := &st.Email
	if emailPb != nil {
		pb.Email = *emailPb
	}

	privateKeyIdPb := &st.PrivateKeyId
	if privateKeyIdPb != nil {
		pb.PrivateKeyId = *privateKeyIdPb
	}

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
	credentialIdField := &pb.CredentialId
	if credentialIdField != nil {
		st.CredentialId = *credentialIdField
	}
	emailField := &pb.Email
	if emailField != nil {
		st.Email = *emailField
	}
	privateKeyIdField := &pb.PrivateKeyId
	if privateKeyIdField != nil {
		st.PrivateKeyId = *privateKeyIdField
	}

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
	CredentialId string
	// The email of the service account. This is an output-only field.
	Email string

	ForceSendFields []string
}

func databricksGcpServiceAccountResponseToPb(st *DatabricksGcpServiceAccountResponse) (*databricksGcpServiceAccountResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databricksGcpServiceAccountResponsePb{}
	credentialIdPb := &st.CredentialId
	if credentialIdPb != nil {
		pb.CredentialId = *credentialIdPb
	}

	emailPb := &st.Email
	if emailPb != nil {
		pb.Email = *emailPb
	}

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
	credentialIdField := &pb.CredentialId
	if credentialIdField != nil {
		st.CredentialId = *credentialIdField
	}
	emailField := &pb.Email
	if emailField != nil {
		st.Email = *emailField
	}

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
	MetastoreId string
	// Workspace ID.
	WorkspaceId int64
}

func deleteAccountMetastoreAssignmentRequestToPb(st *DeleteAccountMetastoreAssignmentRequest) (*deleteAccountMetastoreAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountMetastoreAssignmentRequestPb{}
	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	workspaceIdPb := &st.WorkspaceId
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

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
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	workspaceIdField := &pb.WorkspaceId
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}

	return st, nil
}

// Delete a metastore
type DeleteAccountMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force bool
	// Unity Catalog metastore ID
	MetastoreId string

	ForceSendFields []string
}

func deleteAccountMetastoreRequestToPb(st *DeleteAccountMetastoreRequest) (*deleteAccountMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountMetastoreRequestPb{}
	forcePb := &st.Force
	if forcePb != nil {
		pb.Force = *forcePb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

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
	forceField := &pb.Force
	if forceField != nil {
		st.Force = *forceField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}

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
	Force bool
	// Unity Catalog metastore ID
	MetastoreId string
	// Name of the storage credential.
	StorageCredentialName string

	ForceSendFields []string
}

func deleteAccountStorageCredentialRequestToPb(st *DeleteAccountStorageCredentialRequest) (*deleteAccountStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountStorageCredentialRequestPb{}
	forcePb := &st.Force
	if forcePb != nil {
		pb.Force = *forcePb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	storageCredentialNamePb := &st.StorageCredentialName
	if storageCredentialNamePb != nil {
		pb.StorageCredentialName = *storageCredentialNamePb
	}

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
	forceField := &pb.Force
	if forceField != nil {
		st.Force = *forceField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	storageCredentialNameField := &pb.StorageCredentialName
	if storageCredentialNameField != nil {
		st.StorageCredentialName = *storageCredentialNameField
	}

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
	Alias string
	// The three-level (fully qualified) name of the registered model
	FullName string
}

func deleteAliasRequestToPb(st *DeleteAliasRequest) (*deleteAliasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAliasRequestPb{}
	aliasPb := &st.Alias
	if aliasPb != nil {
		pb.Alias = *aliasPb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

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
	aliasField := &pb.Alias
	if aliasField != nil {
		st.Alias = *aliasField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}

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
	Force bool
	// The name of the catalog.
	Name string

	ForceSendFields []string
}

func deleteCatalogRequestToPb(st *DeleteCatalogRequest) (*deleteCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCatalogRequestPb{}
	forcePb := &st.Force
	if forcePb != nil {
		pb.Force = *forcePb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	forceField := &pb.Force
	if forceField != nil {
		st.Force = *forceField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

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
	Name string
}

func deleteConnectionRequestToPb(st *DeleteConnectionRequest) (*deleteConnectionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteConnectionRequestPb{}
	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

// Delete a credential
type DeleteCredentialRequest struct {
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	Force bool
	// Name of the credential.
	NameArg string

	ForceSendFields []string
}

func deleteCredentialRequestToPb(st *DeleteCredentialRequest) (*deleteCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCredentialRequestPb{}
	forcePb := &st.Force
	if forcePb != nil {
		pb.Force = *forcePb
	}

	nameArgPb := &st.NameArg
	if nameArgPb != nil {
		pb.NameArg = *nameArgPb
	}

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
	forceField := &pb.Force
	if forceField != nil {
		st.Force = *forceField
	}
	nameArgField := &pb.NameArg
	if nameArgField != nil {
		st.NameArg = *nameArgField
	}

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
	Force bool
	// Name of the external location.
	Name string

	ForceSendFields []string
}

func deleteExternalLocationRequestToPb(st *DeleteExternalLocationRequest) (*deleteExternalLocationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteExternalLocationRequestPb{}
	forcePb := &st.Force
	if forcePb != nil {
		pb.Force = *forcePb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	forceField := &pb.Force
	if forceField != nil {
		st.Force = *forceField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

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
	Force bool
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string

	ForceSendFields []string
}

func deleteFunctionRequestToPb(st *DeleteFunctionRequest) (*deleteFunctionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteFunctionRequestPb{}
	forcePb := &st.Force
	if forcePb != nil {
		pb.Force = *forcePb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	forceField := &pb.Force
	if forceField != nil {
		st.Force = *forceField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

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
	Force bool
	// Unique ID of the metastore.
	Id string

	ForceSendFields []string
}

func deleteMetastoreRequestToPb(st *DeleteMetastoreRequest) (*deleteMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteMetastoreRequestPb{}
	forcePb := &st.Force
	if forcePb != nil {
		pb.Force = *forcePb
	}

	idPb := &st.Id
	if idPb != nil {
		pb.Id = *idPb
	}

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
	forceField := &pb.Force
	if forceField != nil {
		st.Force = *forceField
	}
	idField := &pb.Id
	if idField != nil {
		st.Id = *idField
	}

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
	FullName string
	// The integer version number of the model version
	Version int
}

func deleteModelVersionRequestToPb(st *DeleteModelVersionRequest) (*deleteModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelVersionRequestPb{}
	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	versionPb := &st.Version
	if versionPb != nil {
		pb.Version = *versionPb
	}

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
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	versionField := &pb.Version
	if versionField != nil {
		st.Version = *versionField
	}

	return st, nil
}

// Delete an Online Table
type DeleteOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string
}

func deleteOnlineTableRequestToPb(st *DeleteOnlineTableRequest) (*deleteOnlineTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteOnlineTableRequestPb{}
	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

// Delete a table monitor
type DeleteQualityMonitorRequest struct {
	// Full name of the table.
	TableName string
}

func deleteQualityMonitorRequestToPb(st *DeleteQualityMonitorRequest) (*deleteQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteQualityMonitorRequestPb{}
	tableNamePb := &st.TableName
	if tableNamePb != nil {
		pb.TableName = *tableNamePb
	}

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
	tableNameField := &pb.TableName
	if tableNameField != nil {
		st.TableName = *tableNameField
	}

	return st, nil
}

// Delete a Registered Model
type DeleteRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	FullName string
}

func deleteRegisteredModelRequestToPb(st *DeleteRegisteredModelRequest) (*deleteRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRegisteredModelRequestPb{}
	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

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
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}

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
	Force bool
	// Full name of the schema.
	FullName string

	ForceSendFields []string
}

func deleteSchemaRequestToPb(st *DeleteSchemaRequest) (*deleteSchemaRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteSchemaRequestPb{}
	forcePb := &st.Force
	if forcePb != nil {
		pb.Force = *forcePb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

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
	forceField := &pb.Force
	if forceField != nil {
		st.Force = *forceField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}

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
	Force bool
	// Name of the storage credential.
	Name string

	ForceSendFields []string
}

func deleteStorageCredentialRequestToPb(st *DeleteStorageCredentialRequest) (*deleteStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteStorageCredentialRequestPb{}
	forcePb := &st.Force
	if forcePb != nil {
		pb.Force = *forcePb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	forceField := &pb.Force
	if forceField != nil {
		st.Force = *forceField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

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
	Cascade bool
	// The name of the constraint to delete.
	ConstraintName string
	// Full name of the table referenced by the constraint.
	FullName string
}

func deleteTableConstraintRequestToPb(st *DeleteTableConstraintRequest) (*deleteTableConstraintRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTableConstraintRequestPb{}
	cascadePb := &st.Cascade
	if cascadePb != nil {
		pb.Cascade = *cascadePb
	}

	constraintNamePb := &st.ConstraintName
	if constraintNamePb != nil {
		pb.ConstraintName = *constraintNamePb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

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
	cascadeField := &pb.Cascade
	if cascadeField != nil {
		st.Cascade = *cascadeField
	}
	constraintNameField := &pb.ConstraintName
	if constraintNameField != nil {
		st.ConstraintName = *constraintNameField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}

	return st, nil
}

// Delete a table
type DeleteTableRequest struct {
	// Full name of the table.
	FullName string
}

func deleteTableRequestToPb(st *DeleteTableRequest) (*deleteTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTableRequestPb{}
	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

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
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}

	return st, nil
}

// Delete a Volume
type DeleteVolumeRequest struct {
	// The three-level (fully qualified) name of the volume
	Name string
}

func deleteVolumeRequestToPb(st *DeleteVolumeRequest) (*deleteVolumeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteVolumeRequestPb{}
	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

// Properties pertaining to the current state of the delta table as given by the
// commit server. This does not contain **delta.*** (input) properties in
// __TableInfo.properties__.
type DeltaRuntimePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	DeltaRuntimeProperties map[string]string
}

func deltaRuntimePropertiesKvPairsToPb(st *DeltaRuntimePropertiesKvPairs) (*deltaRuntimePropertiesKvPairsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaRuntimePropertiesKvPairsPb{}

	deltaRuntimePropertiesPb := map[string]string{}
	for k, v := range st.DeltaRuntimeProperties {
		itemPb := &v
		if itemPb != nil {
			deltaRuntimePropertiesPb[k] = *itemPb
		}
	}
	pb.DeltaRuntimeProperties = deltaRuntimePropertiesPb

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

	deltaRuntimePropertiesField := map[string]string{}
	for k, v := range pb.DeltaRuntimeProperties {
		itemField := &v
		if itemField != nil {
			deltaRuntimePropertiesField[k] = *itemField
		}
	}
	st.DeltaRuntimeProperties = deltaRuntimePropertiesField

	return st, nil
}

// A dependency of a SQL object. Either the __table__ field or the __function__
// field must be defined.
type Dependency struct {
	// A function that is dependent on a SQL object.
	Function *FunctionDependency
	// A table that is dependent on a SQL object.
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
	for _, item := range pb.Dependencies {
		itemField, err := dependencyFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			dependenciesField = append(dependenciesField, *itemField)
		}
	}
	st.Dependencies = dependenciesField

	return st, nil
}

// Disable a system schema
type DisableRequest struct {
	// The metastore ID under which the system schema lives.
	MetastoreId string
	// Full name of the system schema.
	SchemaName string
}

func disableRequestToPb(st *DisableRequest) (*disableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &disableRequestPb{}
	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	schemaNamePb := &st.SchemaName
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

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
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	schemaNameField := &pb.SchemaName
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}

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
	for _, item := range pb.PrivilegeAssignments {
		itemField, err := effectivePrivilegeAssignmentFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			privilegeAssignmentsField = append(privilegeAssignmentsField, *itemField)
		}
	}
	st.PrivilegeAssignments = privilegeAssignmentsField

	return st, nil
}

type EffectivePredictiveOptimizationFlag struct {
	// The name of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromName string
	// The type of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	InheritedFromType EffectivePredictiveOptimizationFlagInheritedFromType
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	Value EnablePredictiveOptimization

	ForceSendFields []string
}

func effectivePredictiveOptimizationFlagToPb(st *EffectivePredictiveOptimizationFlag) (*effectivePredictiveOptimizationFlagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &effectivePredictiveOptimizationFlagPb{}
	inheritedFromNamePb := &st.InheritedFromName
	if inheritedFromNamePb != nil {
		pb.InheritedFromName = *inheritedFromNamePb
	}

	inheritedFromTypePb := &st.InheritedFromType
	if inheritedFromTypePb != nil {
		pb.InheritedFromType = *inheritedFromTypePb
	}

	valuePb := &st.Value
	if valuePb != nil {
		pb.Value = *valuePb
	}

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
	inheritedFromNameField := &pb.InheritedFromName
	if inheritedFromNameField != nil {
		st.InheritedFromName = *inheritedFromNameField
	}
	inheritedFromTypeField := &pb.InheritedFromType
	if inheritedFromTypeField != nil {
		st.InheritedFromType = *inheritedFromTypeField
	}
	valueField := &pb.Value
	if valueField != nil {
		st.Value = *valueField
	}

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
	InheritedFromName string
	// The type of the object that conveys this privilege via inheritance. This
	// field is omitted when privilege is not inherited (it's assigned to the
	// securable itself).
	InheritedFromType SecurableType
	// The privilege assigned to the principal.
	Privilege Privilege

	ForceSendFields []string
}

func effectivePrivilegeToPb(st *EffectivePrivilege) (*effectivePrivilegePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &effectivePrivilegePb{}
	inheritedFromNamePb := &st.InheritedFromName
	if inheritedFromNamePb != nil {
		pb.InheritedFromName = *inheritedFromNamePb
	}

	inheritedFromTypePb := &st.InheritedFromType
	if inheritedFromTypePb != nil {
		pb.InheritedFromType = *inheritedFromTypePb
	}

	privilegePb := &st.Privilege
	if privilegePb != nil {
		pb.Privilege = *privilegePb
	}

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
	inheritedFromNameField := &pb.InheritedFromName
	if inheritedFromNameField != nil {
		st.InheritedFromName = *inheritedFromNameField
	}
	inheritedFromTypeField := &pb.InheritedFromType
	if inheritedFromTypeField != nil {
		st.InheritedFromType = *inheritedFromTypeField
	}
	privilegeField := &pb.Privilege
	if privilegeField != nil {
		st.Privilege = *privilegeField
	}

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
	Principal string
	// The privileges conveyed to the principal (either directly or via
	// inheritance).
	Privileges []EffectivePrivilege

	ForceSendFields []string
}

func effectivePrivilegeAssignmentToPb(st *EffectivePrivilegeAssignment) (*effectivePrivilegeAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &effectivePrivilegeAssignmentPb{}
	principalPb := &st.Principal
	if principalPb != nil {
		pb.Principal = *principalPb
	}

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
	principalField := &pb.Principal
	if principalField != nil {
		st.Principal = *principalField
	}

	var privilegesField []EffectivePrivilege
	for _, item := range pb.Privileges {
		itemField, err := effectivePrivilegeFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			privilegesField = append(privilegesField, *itemField)
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
	MetastoreId string
	// Full name of the system schema.
	SchemaName string
}

func enableRequestToPb(st *EnableRequest) (*enableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enableRequestPb{}
	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	schemaNamePb := &st.SchemaName
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

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
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	schemaNameField := &pb.SchemaName
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}

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
	FullName string
}

func existsRequestToPb(st *ExistsRequest) (*existsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &existsRequestPb{}
	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

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
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}

	return st, nil
}

type ExternalLocationInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	AccessPoint string
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool
	// User-provided free-form text description.
	Comment string
	// Time at which this external location was created, in epoch milliseconds.
	CreatedAt int64
	// Username of external location creator.
	CreatedBy string
	// Unique ID of the location's storage credential.
	CredentialId string
	// Name of the storage credential used with this location.
	CredentialName string
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback bool

	IsolationMode IsolationMode
	// Unique identifier of metastore hosting the external location.
	MetastoreId string
	// Name of the external location.
	Name string
	// The owner of the external location.
	Owner string
	// Indicates whether the external location is read-only.
	ReadOnly bool
	// Time at which external location this was last modified, in epoch
	// milliseconds.
	UpdatedAt int64
	// Username of user who last modified the external location.
	UpdatedBy string
	// Path URL of the external location.
	Url string

	ForceSendFields []string
}

func externalLocationInfoToPb(st *ExternalLocationInfo) (*externalLocationInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalLocationInfoPb{}
	accessPointPb := &st.AccessPoint
	if accessPointPb != nil {
		pb.AccessPoint = *accessPointPb
	}

	browseOnlyPb := &st.BrowseOnly
	if browseOnlyPb != nil {
		pb.BrowseOnly = *browseOnlyPb
	}

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	createdAtPb := &st.CreatedAt
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	credentialIdPb := &st.CredentialId
	if credentialIdPb != nil {
		pb.CredentialId = *credentialIdPb
	}

	credentialNamePb := &st.CredentialName
	if credentialNamePb != nil {
		pb.CredentialName = *credentialNamePb
	}

	encryptionDetailsPb, err := encryptionDetailsToPb(st.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsPb != nil {
		pb.EncryptionDetails = encryptionDetailsPb
	}

	fallbackPb := &st.Fallback
	if fallbackPb != nil {
		pb.Fallback = *fallbackPb
	}

	isolationModePb := &st.IsolationMode
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	readOnlyPb := &st.ReadOnly
	if readOnlyPb != nil {
		pb.ReadOnly = *readOnlyPb
	}

	updatedAtPb := &st.UpdatedAt
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb := &st.UpdatedBy
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

	urlPb := &st.Url
	if urlPb != nil {
		pb.Url = *urlPb
	}

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
	accessPointField := &pb.AccessPoint
	if accessPointField != nil {
		st.AccessPoint = *accessPointField
	}
	browseOnlyField := &pb.BrowseOnly
	if browseOnlyField != nil {
		st.BrowseOnly = *browseOnlyField
	}
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	createdAtField := &pb.CreatedAt
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	credentialIdField := &pb.CredentialId
	if credentialIdField != nil {
		st.CredentialId = *credentialIdField
	}
	credentialNameField := &pb.CredentialName
	if credentialNameField != nil {
		st.CredentialName = *credentialNameField
	}
	encryptionDetailsField, err := encryptionDetailsFromPb(pb.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsField != nil {
		st.EncryptionDetails = encryptionDetailsField
	}
	fallbackField := &pb.Fallback
	if fallbackField != nil {
		st.Fallback = *fallbackField
	}
	isolationModeField := &pb.IsolationMode
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}
	readOnlyField := &pb.ReadOnly
	if readOnlyField != nil {
		st.ReadOnly = *readOnlyField
	}
	updatedAtField := &pb.UpdatedAt
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField := &pb.UpdatedBy
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}
	urlField := &pb.Url
	if urlField != nil {
		st.Url = *urlField
	}

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
	LastProcessedCommitVersion int64
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table. Only populated if the table is still online
	// and available for serving.
	Timestamp *time.Time

	ForceSendFields []string
}

func failedStatusToPb(st *FailedStatus) (*failedStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &failedStatusPb{}
	lastProcessedCommitVersionPb := &st.LastProcessedCommitVersion
	if lastProcessedCommitVersionPb != nil {
		pb.LastProcessedCommitVersion = *lastProcessedCommitVersionPb
	}

	timestampPb, err := timestampToPb(st.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampPb != nil {
		pb.Timestamp = *timestampPb
	}

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
	lastProcessedCommitVersionField := &pb.LastProcessedCommitVersion
	if lastProcessedCommitVersionField != nil {
		st.LastProcessedCommitVersion = *lastProcessedCommitVersionField
	}
	timestampField, err := timestampFromPb(&pb.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampField != nil {
		st.Timestamp = timestampField
	}

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
	ChildColumns []string
	// The name of the constraint.
	Name string
	// Column names for this constraint.
	ParentColumns []string
	// The full name of the parent constraint.
	ParentTable string
}

func foreignKeyConstraintToPb(st *ForeignKeyConstraint) (*foreignKeyConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &foreignKeyConstraintPb{}

	var childColumnsPb []string
	for _, item := range st.ChildColumns {
		itemPb := &item
		if itemPb != nil {
			childColumnsPb = append(childColumnsPb, *itemPb)
		}
	}
	pb.ChildColumns = childColumnsPb

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	var parentColumnsPb []string
	for _, item := range st.ParentColumns {
		itemPb := &item
		if itemPb != nil {
			parentColumnsPb = append(parentColumnsPb, *itemPb)
		}
	}
	pb.ParentColumns = parentColumnsPb

	parentTablePb := &st.ParentTable
	if parentTablePb != nil {
		pb.ParentTable = *parentTablePb
	}

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

	var childColumnsField []string
	for _, item := range pb.ChildColumns {
		itemField := &item
		if itemField != nil {
			childColumnsField = append(childColumnsField, *itemField)
		}
	}
	st.ChildColumns = childColumnsField
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	var parentColumnsField []string
	for _, item := range pb.ParentColumns {
		itemField := &item
		if itemField != nil {
			parentColumnsField = append(parentColumnsField, *itemField)
		}
	}
	st.ParentColumns = parentColumnsField
	parentTableField := &pb.ParentTable
	if parentTableField != nil {
		st.ParentTable = *parentTableField
	}

	return st, nil
}

// A function that is dependent on a SQL object.
type FunctionDependency struct {
	// Full name of the dependent function, in the form of
	// __catalog_name__.__schema_name__.__function_name__.
	FunctionFullName string
}

func functionDependencyToPb(st *FunctionDependency) (*functionDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &functionDependencyPb{}
	functionFullNamePb := &st.FunctionFullName
	if functionFullNamePb != nil {
		pb.FunctionFullName = *functionFullNamePb
	}

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
	functionFullNameField := &pb.FunctionFullName
	if functionFullNameField != nil {
		st.FunctionFullName = *functionFullNameField
	}

	return st, nil
}

type FunctionInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool
	// Name of parent catalog.
	CatalogName string
	// User-provided free-form text description.
	Comment string
	// Time at which this function was created, in epoch milliseconds.
	CreatedAt int64
	// Username of function creator.
	CreatedBy string
	// Scalar function return data type.
	DataType ColumnTypeName
	// External function language.
	ExternalLanguage string
	// External function name.
	ExternalName string
	// Pretty printed function data type.
	FullDataType string
	// Full name of function, in form of
	// __catalog_name__.__schema_name__.__function__name__
	FullName string
	// Id of Function, relative to parent schema.
	FunctionId string

	InputParams *FunctionParameterInfos
	// Whether the function is deterministic.
	IsDeterministic bool
	// Function null call.
	IsNullCall bool
	// Unique identifier of parent metastore.
	MetastoreId string
	// Name of function, relative to parent schema.
	Name string
	// Username of current owner of function.
	Owner string
	// Function parameter style. **S** is the value for SQL.
	ParameterStyle FunctionInfoParameterStyle
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	Properties string
	// Table function return parameters.
	ReturnParams *FunctionParameterInfos
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	RoutineBody FunctionInfoRoutineBody
	// Function body.
	RoutineDefinition string
	// Function dependencies.
	RoutineDependencies *DependencyList
	// Name of parent schema relative to its parent catalog.
	SchemaName string
	// Function security type.
	SecurityType FunctionInfoSecurityType
	// Specific name of the function; Reserved for future use.
	SpecificName string
	// Function SQL data access.
	SqlDataAccess FunctionInfoSqlDataAccess
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string
	// Time at which this function was created, in epoch milliseconds.
	UpdatedAt int64
	// Username of user who last modified function.
	UpdatedBy string

	ForceSendFields []string
}

func functionInfoToPb(st *FunctionInfo) (*functionInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &functionInfoPb{}
	browseOnlyPb := &st.BrowseOnly
	if browseOnlyPb != nil {
		pb.BrowseOnly = *browseOnlyPb
	}

	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	createdAtPb := &st.CreatedAt
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	dataTypePb := &st.DataType
	if dataTypePb != nil {
		pb.DataType = *dataTypePb
	}

	externalLanguagePb := &st.ExternalLanguage
	if externalLanguagePb != nil {
		pb.ExternalLanguage = *externalLanguagePb
	}

	externalNamePb := &st.ExternalName
	if externalNamePb != nil {
		pb.ExternalName = *externalNamePb
	}

	fullDataTypePb := &st.FullDataType
	if fullDataTypePb != nil {
		pb.FullDataType = *fullDataTypePb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	functionIdPb := &st.FunctionId
	if functionIdPb != nil {
		pb.FunctionId = *functionIdPb
	}

	inputParamsPb, err := functionParameterInfosToPb(st.InputParams)
	if err != nil {
		return nil, err
	}
	if inputParamsPb != nil {
		pb.InputParams = inputParamsPb
	}

	isDeterministicPb := &st.IsDeterministic
	if isDeterministicPb != nil {
		pb.IsDeterministic = *isDeterministicPb
	}

	isNullCallPb := &st.IsNullCall
	if isNullCallPb != nil {
		pb.IsNullCall = *isNullCallPb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	parameterStylePb := &st.ParameterStyle
	if parameterStylePb != nil {
		pb.ParameterStyle = *parameterStylePb
	}

	propertiesPb := &st.Properties
	if propertiesPb != nil {
		pb.Properties = *propertiesPb
	}

	returnParamsPb, err := functionParameterInfosToPb(st.ReturnParams)
	if err != nil {
		return nil, err
	}
	if returnParamsPb != nil {
		pb.ReturnParams = returnParamsPb
	}

	routineBodyPb := &st.RoutineBody
	if routineBodyPb != nil {
		pb.RoutineBody = *routineBodyPb
	}

	routineDefinitionPb := &st.RoutineDefinition
	if routineDefinitionPb != nil {
		pb.RoutineDefinition = *routineDefinitionPb
	}

	routineDependenciesPb, err := dependencyListToPb(st.RoutineDependencies)
	if err != nil {
		return nil, err
	}
	if routineDependenciesPb != nil {
		pb.RoutineDependencies = routineDependenciesPb
	}

	schemaNamePb := &st.SchemaName
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

	securityTypePb := &st.SecurityType
	if securityTypePb != nil {
		pb.SecurityType = *securityTypePb
	}

	specificNamePb := &st.SpecificName
	if specificNamePb != nil {
		pb.SpecificName = *specificNamePb
	}

	sqlDataAccessPb := &st.SqlDataAccess
	if sqlDataAccessPb != nil {
		pb.SqlDataAccess = *sqlDataAccessPb
	}

	sqlPathPb := &st.SqlPath
	if sqlPathPb != nil {
		pb.SqlPath = *sqlPathPb
	}

	updatedAtPb := &st.UpdatedAt
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb := &st.UpdatedBy
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

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
	browseOnlyField := &pb.BrowseOnly
	if browseOnlyField != nil {
		st.BrowseOnly = *browseOnlyField
	}
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	createdAtField := &pb.CreatedAt
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	dataTypeField := &pb.DataType
	if dataTypeField != nil {
		st.DataType = *dataTypeField
	}
	externalLanguageField := &pb.ExternalLanguage
	if externalLanguageField != nil {
		st.ExternalLanguage = *externalLanguageField
	}
	externalNameField := &pb.ExternalName
	if externalNameField != nil {
		st.ExternalName = *externalNameField
	}
	fullDataTypeField := &pb.FullDataType
	if fullDataTypeField != nil {
		st.FullDataType = *fullDataTypeField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	functionIdField := &pb.FunctionId
	if functionIdField != nil {
		st.FunctionId = *functionIdField
	}
	inputParamsField, err := functionParameterInfosFromPb(pb.InputParams)
	if err != nil {
		return nil, err
	}
	if inputParamsField != nil {
		st.InputParams = inputParamsField
	}
	isDeterministicField := &pb.IsDeterministic
	if isDeterministicField != nil {
		st.IsDeterministic = *isDeterministicField
	}
	isNullCallField := &pb.IsNullCall
	if isNullCallField != nil {
		st.IsNullCall = *isNullCallField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}
	parameterStyleField := &pb.ParameterStyle
	if parameterStyleField != nil {
		st.ParameterStyle = *parameterStyleField
	}
	propertiesField := &pb.Properties
	if propertiesField != nil {
		st.Properties = *propertiesField
	}
	returnParamsField, err := functionParameterInfosFromPb(pb.ReturnParams)
	if err != nil {
		return nil, err
	}
	if returnParamsField != nil {
		st.ReturnParams = returnParamsField
	}
	routineBodyField := &pb.RoutineBody
	if routineBodyField != nil {
		st.RoutineBody = *routineBodyField
	}
	routineDefinitionField := &pb.RoutineDefinition
	if routineDefinitionField != nil {
		st.RoutineDefinition = *routineDefinitionField
	}
	routineDependenciesField, err := dependencyListFromPb(pb.RoutineDependencies)
	if err != nil {
		return nil, err
	}
	if routineDependenciesField != nil {
		st.RoutineDependencies = routineDependenciesField
	}
	schemaNameField := &pb.SchemaName
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}
	securityTypeField := &pb.SecurityType
	if securityTypeField != nil {
		st.SecurityType = *securityTypeField
	}
	specificNameField := &pb.SpecificName
	if specificNameField != nil {
		st.SpecificName = *specificNameField
	}
	sqlDataAccessField := &pb.SqlDataAccess
	if sqlDataAccessField != nil {
		st.SqlDataAccess = *sqlDataAccessField
	}
	sqlPathField := &pb.SqlPath
	if sqlPathField != nil {
		st.SqlPath = *sqlPathField
	}
	updatedAtField := &pb.UpdatedAt
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField := &pb.UpdatedBy
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}

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
	Comment string
	// Name of parameter.
	Name string
	// Default value of the parameter.
	ParameterDefault string
	// The mode of the function parameter.
	ParameterMode FunctionParameterMode
	// The type of function parameter.
	ParameterType FunctionParameterType
	// Ordinal position of column (starting at position 0).
	Position int
	// Format of IntervalType.
	TypeIntervalType string
	// Full data type spec, JSON-serialized.
	TypeJson string

	TypeName ColumnTypeName
	// Digits of precision; required on Create for DecimalTypes.
	TypePrecision int
	// Digits to right of decimal; Required on Create for DecimalTypes.
	TypeScale int
	// Full data type spec, SQL/catalogString text.
	TypeText string

	ForceSendFields []string
}

func functionParameterInfoToPb(st *FunctionParameterInfo) (*functionParameterInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &functionParameterInfoPb{}
	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	parameterDefaultPb := &st.ParameterDefault
	if parameterDefaultPb != nil {
		pb.ParameterDefault = *parameterDefaultPb
	}

	parameterModePb := &st.ParameterMode
	if parameterModePb != nil {
		pb.ParameterMode = *parameterModePb
	}

	parameterTypePb := &st.ParameterType
	if parameterTypePb != nil {
		pb.ParameterType = *parameterTypePb
	}

	positionPb := &st.Position
	if positionPb != nil {
		pb.Position = *positionPb
	}

	typeIntervalTypePb := &st.TypeIntervalType
	if typeIntervalTypePb != nil {
		pb.TypeIntervalType = *typeIntervalTypePb
	}

	typeJsonPb := &st.TypeJson
	if typeJsonPb != nil {
		pb.TypeJson = *typeJsonPb
	}

	typeNamePb := &st.TypeName
	if typeNamePb != nil {
		pb.TypeName = *typeNamePb
	}

	typePrecisionPb := &st.TypePrecision
	if typePrecisionPb != nil {
		pb.TypePrecision = *typePrecisionPb
	}

	typeScalePb := &st.TypeScale
	if typeScalePb != nil {
		pb.TypeScale = *typeScalePb
	}

	typeTextPb := &st.TypeText
	if typeTextPb != nil {
		pb.TypeText = *typeTextPb
	}

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	parameterDefaultField := &pb.ParameterDefault
	if parameterDefaultField != nil {
		st.ParameterDefault = *parameterDefaultField
	}
	parameterModeField := &pb.ParameterMode
	if parameterModeField != nil {
		st.ParameterMode = *parameterModeField
	}
	parameterTypeField := &pb.ParameterType
	if parameterTypeField != nil {
		st.ParameterType = *parameterTypeField
	}
	positionField := &pb.Position
	if positionField != nil {
		st.Position = *positionField
	}
	typeIntervalTypeField := &pb.TypeIntervalType
	if typeIntervalTypeField != nil {
		st.TypeIntervalType = *typeIntervalTypeField
	}
	typeJsonField := &pb.TypeJson
	if typeJsonField != nil {
		st.TypeJson = *typeJsonField
	}
	typeNameField := &pb.TypeName
	if typeNameField != nil {
		st.TypeName = *typeNameField
	}
	typePrecisionField := &pb.TypePrecision
	if typePrecisionField != nil {
		st.TypePrecision = *typePrecisionField
	}
	typeScaleField := &pb.TypeScale
	if typeScaleField != nil {
		st.TypeScale = *typeScaleField
	}
	typeTextField := &pb.TypeText
	if typeTextField != nil {
		st.TypeText = *typeTextField
	}

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
	for _, item := range pb.Parameters {
		itemField, err := functionParameterInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			parametersField = append(parametersField, *itemField)
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
	OauthToken string

	ForceSendFields []string
}

func gcpOauthTokenToPb(st *GcpOauthToken) (*gcpOauthTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpOauthTokenPb{}
	oauthTokenPb := &st.OauthToken
	if oauthTokenPb != nil {
		pb.OauthToken = *oauthTokenPb
	}

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
	oauthTokenField := &pb.OauthToken
	if oauthTokenField != nil {
		st.OauthToken = *oauthTokenField
	}

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
	Resources []string
}

func generateTemporaryServiceCredentialAzureOptionsToPb(st *GenerateTemporaryServiceCredentialAzureOptions) (*generateTemporaryServiceCredentialAzureOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &generateTemporaryServiceCredentialAzureOptionsPb{}

	var resourcesPb []string
	for _, item := range st.Resources {
		itemPb := &item
		if itemPb != nil {
			resourcesPb = append(resourcesPb, *itemPb)
		}
	}
	pb.Resources = resourcesPb

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

	var resourcesField []string
	for _, item := range pb.Resources {
		itemField := &item
		if itemField != nil {
			resourcesField = append(resourcesField, *itemField)
		}
	}
	st.Resources = resourcesField

	return st, nil
}

// The GCP cloud options to customize the requested temporary credential
type GenerateTemporaryServiceCredentialGcpOptions struct {
	// The scopes to which the temporary GCP credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://google-auth.readthedocs.io/en/latest/reference/google.auth.html#google.auth.credentials.Credentials)
	Scopes []string
}

func generateTemporaryServiceCredentialGcpOptionsToPb(st *GenerateTemporaryServiceCredentialGcpOptions) (*generateTemporaryServiceCredentialGcpOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &generateTemporaryServiceCredentialGcpOptionsPb{}

	var scopesPb []string
	for _, item := range st.Scopes {
		itemPb := &item
		if itemPb != nil {
			scopesPb = append(scopesPb, *itemPb)
		}
	}
	pb.Scopes = scopesPb

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

	var scopesField []string
	for _, item := range pb.Scopes {
		itemField := &item
		if itemField != nil {
			scopesField = append(scopesField, *itemField)
		}
	}
	st.Scopes = scopesField

	return st, nil
}

type GenerateTemporaryServiceCredentialRequest struct {
	// The Azure cloud options to customize the requested temporary credential
	AzureOptions *GenerateTemporaryServiceCredentialAzureOptions
	// The name of the service credential used to generate a temporary
	// credential
	CredentialName string
	// The GCP cloud options to customize the requested temporary credential
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

	credentialNamePb := &st.CredentialName
	if credentialNamePb != nil {
		pb.CredentialName = *credentialNamePb
	}

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
	credentialNameField := &pb.CredentialName
	if credentialNameField != nil {
		st.CredentialName = *credentialNameField
	}
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
	Operation TableOperation
	// UUID of the table to read or write.
	TableId string

	ForceSendFields []string
}

func generateTemporaryTableCredentialRequestToPb(st *GenerateTemporaryTableCredentialRequest) (*generateTemporaryTableCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &generateTemporaryTableCredentialRequestPb{}
	operationPb := &st.Operation
	if operationPb != nil {
		pb.Operation = *operationPb
	}

	tableIdPb := &st.TableId
	if tableIdPb != nil {
		pb.TableId = *tableIdPb
	}

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
	operationField := &pb.Operation
	if operationField != nil {
		st.Operation = *operationField
	}
	tableIdField := &pb.TableId
	if tableIdField != nil {
		st.TableId = *tableIdField
	}

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
	AwsTempCredentials *AwsCredentials
	// Azure Active Directory token, essentially the Oauth token for Azure
	// Service Principal or Managed Identity. Read more at
	// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
	AzureAad *AzureActiveDirectoryToken
	// Azure temporary credentials for API authentication. Read more at
	// https://docs.microsoft.com/en-us/rest/api/storageservices/create-user-delegation-sas
	AzureUserDelegationSas *AzureUserDelegationSas
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	ExpirationTime int64
	// GCP temporary credentials for API authentication. Read more at
	// https://developers.google.com/identity/protocols/oauth2/service-account
	GcpOauthToken *GcpOauthToken
	// R2 temporary credentials for API authentication. Read more at
	// https://developers.cloudflare.com/r2/api/s3/tokens/.
	R2TempCredentials *R2Credentials
	// The URL of the storage path accessible by the temporary credential.
	Url string

	ForceSendFields []string
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

	expirationTimePb := &st.ExpirationTime
	if expirationTimePb != nil {
		pb.ExpirationTime = *expirationTimePb
	}

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

	urlPb := &st.Url
	if urlPb != nil {
		pb.Url = *urlPb
	}

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
	expirationTimeField := &pb.ExpirationTime
	if expirationTimeField != nil {
		st.ExpirationTime = *expirationTimeField
	}
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
	urlField := &pb.Url
	if urlField != nil {
		st.Url = *urlField
	}

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
	WorkspaceId int64
}

func getAccountMetastoreAssignmentRequestToPb(st *GetAccountMetastoreAssignmentRequest) (*getAccountMetastoreAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountMetastoreAssignmentRequestPb{}
	workspaceIdPb := &st.WorkspaceId
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

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
	workspaceIdField := &pb.WorkspaceId
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}

	return st, nil
}

// Get a metastore
type GetAccountMetastoreRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string
}

func getAccountMetastoreRequestToPb(st *GetAccountMetastoreRequest) (*getAccountMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountMetastoreRequestPb{}
	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

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
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}

	return st, nil
}

// Gets the named storage credential
type GetAccountStorageCredentialRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string
	// Name of the storage credential.
	StorageCredentialName string
}

func getAccountStorageCredentialRequestToPb(st *GetAccountStorageCredentialRequest) (*getAccountStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountStorageCredentialRequestPb{}
	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	storageCredentialNamePb := &st.StorageCredentialName
	if storageCredentialNamePb != nil {
		pb.StorageCredentialName = *storageCredentialNamePb
	}

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
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	storageCredentialNameField := &pb.StorageCredentialName
	if storageCredentialNameField != nil {
		st.StorageCredentialName = *storageCredentialNameField
	}

	return st, nil
}

// Get an artifact allowlist
type GetArtifactAllowlistRequest struct {
	// The artifact type of the allowlist.
	ArtifactType ArtifactType
}

func getArtifactAllowlistRequestToPb(st *GetArtifactAllowlistRequest) (*getArtifactAllowlistRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getArtifactAllowlistRequestPb{}
	artifactTypePb := &st.ArtifactType
	if artifactTypePb != nil {
		pb.ArtifactType = *artifactTypePb
	}

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
	artifactTypeField := &pb.ArtifactType
	if artifactTypeField != nil {
		st.ArtifactType = *artifactTypeField
	}

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
	MaxResults int
	// Opaque pagination token to go to next page based on previous query.
	PageToken string
	// The name of the securable.
	SecurableName string
	// The type of the securable to bind to a workspace.
	SecurableType GetBindingsSecurableType

	ForceSendFields []string
}

func getBindingsRequestToPb(st *GetBindingsRequest) (*getBindingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getBindingsRequestPb{}
	maxResultsPb := &st.MaxResults
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	securableNamePb := &st.SecurableName
	if securableNamePb != nil {
		pb.SecurableName = *securableNamePb
	}

	securableTypePb := &st.SecurableType
	if securableTypePb != nil {
		pb.SecurableType = *securableTypePb
	}

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
	maxResultsField := &pb.MaxResults
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	securableNameField := &pb.SecurableName
	if securableNameField != nil {
		st.SecurableName = *securableNameField
	}
	securableTypeField := &pb.SecurableType
	if securableTypeField != nil {
		st.SecurableType = *securableTypeField
	}

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
	Alias string
	// The three-level (fully qualified) name of the registered model
	FullName string
	// Whether to include aliases associated with the model version in the
	// response
	IncludeAliases bool

	ForceSendFields []string
}

func getByAliasRequestToPb(st *GetByAliasRequest) (*getByAliasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getByAliasRequestPb{}
	aliasPb := &st.Alias
	if aliasPb != nil {
		pb.Alias = *aliasPb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	includeAliasesPb := &st.IncludeAliases
	if includeAliasesPb != nil {
		pb.IncludeAliases = *includeAliasesPb
	}

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
	aliasField := &pb.Alias
	if aliasField != nil {
		st.Alias = *aliasField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	includeAliasesField := &pb.IncludeAliases
	if includeAliasesField != nil {
		st.IncludeAliases = *includeAliasesField
	}

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
	IncludeBrowse bool
	// The name of the catalog.
	Name string

	ForceSendFields []string
}

func getCatalogRequestToPb(st *GetCatalogRequest) (*getCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCatalogRequestPb{}
	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

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
	Name string
}

func getConnectionRequestToPb(st *GetConnectionRequest) (*getConnectionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getConnectionRequestPb{}
	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

// Get a credential
type GetCredentialRequest struct {
	// Name of the credential.
	NameArg string
}

func getCredentialRequestToPb(st *GetCredentialRequest) (*getCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCredentialRequestPb{}
	nameArgPb := &st.NameArg
	if nameArgPb != nil {
		pb.NameArg = *nameArgPb
	}

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
	nameArgField := &pb.NameArg
	if nameArgField != nil {
		st.NameArg = *nameArgField
	}

	return st, nil
}

// Get effective permissions
type GetEffectiveRequest struct {
	// Full name of securable.
	FullName string
	// If provided, only the effective permissions for the specified principal
	// (user or group) are returned.
	Principal string
	// Type of securable.
	SecurableType SecurableType

	ForceSendFields []string
}

func getEffectiveRequestToPb(st *GetEffectiveRequest) (*getEffectiveRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEffectiveRequestPb{}
	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	principalPb := &st.Principal
	if principalPb != nil {
		pb.Principal = *principalPb
	}

	securableTypePb := &st.SecurableType
	if securableTypePb != nil {
		pb.SecurableType = *securableTypePb
	}

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
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	principalField := &pb.Principal
	if principalField != nil {
		st.Principal = *principalField
	}
	securableTypeField := &pb.SecurableType
	if securableTypeField != nil {
		st.SecurableType = *securableTypeField
	}

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
	IncludeBrowse bool
	// Name of the external location.
	Name string

	ForceSendFields []string
}

func getExternalLocationRequestToPb(st *GetExternalLocationRequest) (*getExternalLocationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExternalLocationRequestPb{}
	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

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
	IncludeBrowse bool
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string

	ForceSendFields []string
}

func getFunctionRequestToPb(st *GetFunctionRequest) (*getFunctionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getFunctionRequestPb{}
	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

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
	FullName string
	// If provided, only the permissions for the specified principal (user or
	// group) are returned.
	Principal string
	// Type of securable.
	SecurableType SecurableType

	ForceSendFields []string
}

func getGrantRequestToPb(st *GetGrantRequest) (*getGrantRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getGrantRequestPb{}
	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	principalPb := &st.Principal
	if principalPb != nil {
		pb.Principal = *principalPb
	}

	securableTypePb := &st.SecurableType
	if securableTypePb != nil {
		pb.SecurableType = *securableTypePb
	}

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
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	principalField := &pb.Principal
	if principalField != nil {
		st.Principal = *principalField
	}
	securableTypeField := &pb.SecurableType
	if securableTypeField != nil {
		st.SecurableType = *securableTypeField
	}

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
	Id string
}

func getMetastoreRequestToPb(st *GetMetastoreRequest) (*getMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getMetastoreRequestPb{}
	idPb := &st.Id
	if idPb != nil {
		pb.Id = *idPb
	}

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
	idField := &pb.Id
	if idField != nil {
		st.Id = *idField
	}

	return st, nil
}

type GetMetastoreSummaryResponse struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	Cloud string
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt int64
	// Username of metastore creator.
	CreatedBy string
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId string
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope GetMetastoreSummaryResponseDeltaSharingScope
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled bool
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId string
	// Unique identifier of metastore.
	MetastoreId string
	// The user-specified name of the metastore.
	Name string
	// The owner of the metastore.
	Owner string
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region string
	// The storage root URL for metastore
	StorageRoot string
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName string
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt int64
	// Username of user who last modified the metastore.
	UpdatedBy string

	ForceSendFields []string
}

func getMetastoreSummaryResponseToPb(st *GetMetastoreSummaryResponse) (*getMetastoreSummaryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getMetastoreSummaryResponsePb{}
	cloudPb := &st.Cloud
	if cloudPb != nil {
		pb.Cloud = *cloudPb
	}

	createdAtPb := &st.CreatedAt
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	defaultDataAccessConfigIdPb := &st.DefaultDataAccessConfigId
	if defaultDataAccessConfigIdPb != nil {
		pb.DefaultDataAccessConfigId = *defaultDataAccessConfigIdPb
	}

	deltaSharingOrganizationNamePb := &st.DeltaSharingOrganizationName
	if deltaSharingOrganizationNamePb != nil {
		pb.DeltaSharingOrganizationName = *deltaSharingOrganizationNamePb
	}

	deltaSharingRecipientTokenLifetimeInSecondsPb := &st.DeltaSharingRecipientTokenLifetimeInSeconds
	if deltaSharingRecipientTokenLifetimeInSecondsPb != nil {
		pb.DeltaSharingRecipientTokenLifetimeInSeconds = *deltaSharingRecipientTokenLifetimeInSecondsPb
	}

	deltaSharingScopePb := &st.DeltaSharingScope
	if deltaSharingScopePb != nil {
		pb.DeltaSharingScope = *deltaSharingScopePb
	}

	externalAccessEnabledPb := &st.ExternalAccessEnabled
	if externalAccessEnabledPb != nil {
		pb.ExternalAccessEnabled = *externalAccessEnabledPb
	}

	globalMetastoreIdPb := &st.GlobalMetastoreId
	if globalMetastoreIdPb != nil {
		pb.GlobalMetastoreId = *globalMetastoreIdPb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	privilegeModelVersionPb := &st.PrivilegeModelVersion
	if privilegeModelVersionPb != nil {
		pb.PrivilegeModelVersion = *privilegeModelVersionPb
	}

	regionPb := &st.Region
	if regionPb != nil {
		pb.Region = *regionPb
	}

	storageRootPb := &st.StorageRoot
	if storageRootPb != nil {
		pb.StorageRoot = *storageRootPb
	}

	storageRootCredentialIdPb := &st.StorageRootCredentialId
	if storageRootCredentialIdPb != nil {
		pb.StorageRootCredentialId = *storageRootCredentialIdPb
	}

	storageRootCredentialNamePb := &st.StorageRootCredentialName
	if storageRootCredentialNamePb != nil {
		pb.StorageRootCredentialName = *storageRootCredentialNamePb
	}

	updatedAtPb := &st.UpdatedAt
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb := &st.UpdatedBy
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

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
	cloudField := &pb.Cloud
	if cloudField != nil {
		st.Cloud = *cloudField
	}
	createdAtField := &pb.CreatedAt
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	defaultDataAccessConfigIdField := &pb.DefaultDataAccessConfigId
	if defaultDataAccessConfigIdField != nil {
		st.DefaultDataAccessConfigId = *defaultDataAccessConfigIdField
	}
	deltaSharingOrganizationNameField := &pb.DeltaSharingOrganizationName
	if deltaSharingOrganizationNameField != nil {
		st.DeltaSharingOrganizationName = *deltaSharingOrganizationNameField
	}
	deltaSharingRecipientTokenLifetimeInSecondsField := &pb.DeltaSharingRecipientTokenLifetimeInSeconds
	if deltaSharingRecipientTokenLifetimeInSecondsField != nil {
		st.DeltaSharingRecipientTokenLifetimeInSeconds = *deltaSharingRecipientTokenLifetimeInSecondsField
	}
	deltaSharingScopeField := &pb.DeltaSharingScope
	if deltaSharingScopeField != nil {
		st.DeltaSharingScope = *deltaSharingScopeField
	}
	externalAccessEnabledField := &pb.ExternalAccessEnabled
	if externalAccessEnabledField != nil {
		st.ExternalAccessEnabled = *externalAccessEnabledField
	}
	globalMetastoreIdField := &pb.GlobalMetastoreId
	if globalMetastoreIdField != nil {
		st.GlobalMetastoreId = *globalMetastoreIdField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}
	privilegeModelVersionField := &pb.PrivilegeModelVersion
	if privilegeModelVersionField != nil {
		st.PrivilegeModelVersion = *privilegeModelVersionField
	}
	regionField := &pb.Region
	if regionField != nil {
		st.Region = *regionField
	}
	storageRootField := &pb.StorageRoot
	if storageRootField != nil {
		st.StorageRoot = *storageRootField
	}
	storageRootCredentialIdField := &pb.StorageRootCredentialId
	if storageRootCredentialIdField != nil {
		st.StorageRootCredentialId = *storageRootCredentialIdField
	}
	storageRootCredentialNameField := &pb.StorageRootCredentialName
	if storageRootCredentialNameField != nil {
		st.StorageRootCredentialName = *storageRootCredentialNameField
	}
	updatedAtField := &pb.UpdatedAt
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField := &pb.UpdatedBy
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}

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
	FullName string
	// Whether to include aliases associated with the model version in the
	// response
	IncludeAliases bool
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse bool
	// The integer version number of the model version
	Version int

	ForceSendFields []string
}

func getModelVersionRequestToPb(st *GetModelVersionRequest) (*getModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getModelVersionRequestPb{}
	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	includeAliasesPb := &st.IncludeAliases
	if includeAliasesPb != nil {
		pb.IncludeAliases = *includeAliasesPb
	}

	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

	versionPb := &st.Version
	if versionPb != nil {
		pb.Version = *versionPb
	}

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
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	includeAliasesField := &pb.IncludeAliases
	if includeAliasesField != nil {
		st.IncludeAliases = *includeAliasesField
	}
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}
	versionField := &pb.Version
	if versionField != nil {
		st.Version = *versionField
	}

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
	Name string
}

func getOnlineTableRequestToPb(st *GetOnlineTableRequest) (*getOnlineTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getOnlineTableRequestPb{}
	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

// Get a table monitor
type GetQualityMonitorRequest struct {
	// Full name of the table.
	TableName string
}

func getQualityMonitorRequestToPb(st *GetQualityMonitorRequest) (*getQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getQualityMonitorRequestPb{}
	tableNamePb := &st.TableName
	if tableNamePb != nil {
		pb.TableName = *tableNamePb
	}

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
	tableNameField := &pb.TableName
	if tableNameField != nil {
		st.TableName = *tableNameField
	}

	return st, nil
}

// Get information for a single resource quota.
type GetQuotaRequest struct {
	// Full name of the parent resource. Provide the metastore ID if the parent
	// is a metastore.
	ParentFullName string
	// Securable type of the quota parent.
	ParentSecurableType string
	// Name of the quota. Follows the pattern of the quota type, with "-quota"
	// added as a suffix.
	QuotaName string
}

func getQuotaRequestToPb(st *GetQuotaRequest) (*getQuotaRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getQuotaRequestPb{}
	parentFullNamePb := &st.ParentFullName
	if parentFullNamePb != nil {
		pb.ParentFullName = *parentFullNamePb
	}

	parentSecurableTypePb := &st.ParentSecurableType
	if parentSecurableTypePb != nil {
		pb.ParentSecurableType = *parentSecurableTypePb
	}

	quotaNamePb := &st.QuotaName
	if quotaNamePb != nil {
		pb.QuotaName = *quotaNamePb
	}

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
	parentFullNameField := &pb.ParentFullName
	if parentFullNameField != nil {
		st.ParentFullName = *parentFullNameField
	}
	parentSecurableTypeField := &pb.ParentSecurableType
	if parentSecurableTypeField != nil {
		st.ParentSecurableType = *parentSecurableTypeField
	}
	quotaNameField := &pb.QuotaName
	if quotaNameField != nil {
		st.QuotaName = *quotaNameField
	}

	return st, nil
}

type GetQuotaResponse struct {
	// The returned QuotaInfo.
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
	RefreshId string
	// Full name of the table.
	TableName string
}

func getRefreshRequestToPb(st *GetRefreshRequest) (*getRefreshRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRefreshRequestPb{}
	refreshIdPb := &st.RefreshId
	if refreshIdPb != nil {
		pb.RefreshId = *refreshIdPb
	}

	tableNamePb := &st.TableName
	if tableNamePb != nil {
		pb.TableName = *tableNamePb
	}

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
	refreshIdField := &pb.RefreshId
	if refreshIdField != nil {
		st.RefreshId = *refreshIdField
	}
	tableNameField := &pb.TableName
	if tableNameField != nil {
		st.TableName = *tableNameField
	}

	return st, nil
}

// Get a Registered Model
type GetRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	FullName string
	// Whether to include registered model aliases in the response
	IncludeAliases bool
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool

	ForceSendFields []string
}

func getRegisteredModelRequestToPb(st *GetRegisteredModelRequest) (*getRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRegisteredModelRequestPb{}
	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	includeAliasesPb := &st.IncludeAliases
	if includeAliasesPb != nil {
		pb.IncludeAliases = *includeAliasesPb
	}

	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

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
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	includeAliasesField := &pb.IncludeAliases
	if includeAliasesField != nil {
		st.IncludeAliases = *includeAliasesField
	}
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}

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
	FullName string
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool

	ForceSendFields []string
}

func getSchemaRequestToPb(st *GetSchemaRequest) (*getSchemaRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getSchemaRequestPb{}
	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

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
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}

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
	Name string
}

func getStorageCredentialRequestToPb(st *GetStorageCredentialRequest) (*getStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStorageCredentialRequestPb{}
	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

// Get a table
type GetTableRequest struct {
	// Full name of the table.
	FullName string
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata bool
	// Whether to include a manifest containing capabilities the table has.
	IncludeManifestCapabilities bool

	ForceSendFields []string
}

func getTableRequestToPb(st *GetTableRequest) (*getTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getTableRequestPb{}
	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

	includeDeltaMetadataPb := &st.IncludeDeltaMetadata
	if includeDeltaMetadataPb != nil {
		pb.IncludeDeltaMetadata = *includeDeltaMetadataPb
	}

	includeManifestCapabilitiesPb := &st.IncludeManifestCapabilities
	if includeManifestCapabilitiesPb != nil {
		pb.IncludeManifestCapabilities = *includeManifestCapabilitiesPb
	}

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
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}
	includeDeltaMetadataField := &pb.IncludeDeltaMetadata
	if includeDeltaMetadataField != nil {
		st.IncludeDeltaMetadata = *includeDeltaMetadataField
	}
	includeManifestCapabilitiesField := &pb.IncludeManifestCapabilities
	if includeManifestCapabilitiesField != nil {
		st.IncludeManifestCapabilities = *includeManifestCapabilitiesField
	}

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
	Name string
}

func getWorkspaceBindingRequestToPb(st *GetWorkspaceBindingRequest) (*getWorkspaceBindingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceBindingRequestPb{}
	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

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
	MetastoreId string
}

func listAccountMetastoreAssignmentsRequestToPb(st *ListAccountMetastoreAssignmentsRequest) (*listAccountMetastoreAssignmentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountMetastoreAssignmentsRequestPb{}
	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

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
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}

	return st, nil
}

// The list of workspaces to which the given metastore is assigned.
type ListAccountMetastoreAssignmentsResponse struct {
	WorkspaceIds []int64
}

func listAccountMetastoreAssignmentsResponseToPb(st *ListAccountMetastoreAssignmentsResponse) (*listAccountMetastoreAssignmentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountMetastoreAssignmentsResponsePb{}

	var workspaceIdsPb []int64
	for _, item := range st.WorkspaceIds {
		itemPb := &item
		if itemPb != nil {
			workspaceIdsPb = append(workspaceIdsPb, *itemPb)
		}
	}
	pb.WorkspaceIds = workspaceIdsPb

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

	var workspaceIdsField []int64
	for _, item := range pb.WorkspaceIds {
		itemField := &item
		if itemField != nil {
			workspaceIdsField = append(workspaceIdsField, *itemField)
		}
	}
	st.WorkspaceIds = workspaceIdsField

	return st, nil
}

// Get all storage credentials assigned to a metastore
type ListAccountStorageCredentialsRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string
}

func listAccountStorageCredentialsRequestToPb(st *ListAccountStorageCredentialsRequest) (*listAccountStorageCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountStorageCredentialsRequestPb{}
	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

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
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}

	return st, nil
}

type ListAccountStorageCredentialsResponse struct {
	// An array of metastore storage credentials.
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
	for _, item := range pb.StorageCredentials {
		itemField, err := storageCredentialInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			storageCredentialsField = append(storageCredentialsField, *itemField)
		}
	}
	st.StorageCredentials = storageCredentialsField

	return st, nil
}

// List catalogs
type ListCatalogsRequest struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool
	// Maximum number of catalogs to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid catalogs are returned (not
	// recommended). - Note: The number of returned catalogs might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further catalogs can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults int
	// Opaque pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func listCatalogsRequestToPb(st *ListCatalogsRequest) (*listCatalogsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCatalogsRequestPb{}
	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

	maxResultsPb := &st.MaxResults
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}
	maxResultsField := &pb.MaxResults
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	Catalogs []CatalogInfo
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string

	ForceSendFields []string
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

	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	for _, item := range pb.Catalogs {
		itemField, err := catalogInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			catalogsField = append(catalogsField, *itemField)
		}
	}
	st.Catalogs = catalogsField
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	MaxResults int
	// Opaque pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func listConnectionsRequestToPb(st *ListConnectionsRequest) (*listConnectionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listConnectionsRequestPb{}
	maxResultsPb := &st.MaxResults
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	maxResultsField := &pb.MaxResults
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	Connections []ConnectionInfo
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string

	ForceSendFields []string
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

	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	for _, item := range pb.Connections {
		itemField, err := connectionInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			connectionsField = append(connectionsField, *itemField)
		}
	}
	st.Connections = connectionsField
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	MaxResults int
	// Opaque token to retrieve the next page of results.
	PageToken string
	// Return only credentials for the specified purpose.
	Purpose CredentialPurpose

	ForceSendFields []string
}

func listCredentialsRequestToPb(st *ListCredentialsRequest) (*listCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCredentialsRequestPb{}
	maxResultsPb := &st.MaxResults
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	purposePb := &st.Purpose
	if purposePb != nil {
		pb.Purpose = *purposePb
	}

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
	maxResultsField := &pb.MaxResults
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	purposeField := &pb.Purpose
	if purposeField != nil {
		st.Purpose = *purposeField
	}

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
	Credentials []CredentialInfo
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string

	ForceSendFields []string
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

	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	for _, item := range pb.Credentials {
		itemField, err := credentialInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			credentialsField = append(credentialsField, *itemField)
		}
	}
	st.Credentials = credentialsField
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	IncludeBrowse bool
	// Maximum number of external locations to return. If not set, all the
	// external locations are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults int
	// Opaque pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func listExternalLocationsRequestToPb(st *ListExternalLocationsRequest) (*listExternalLocationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExternalLocationsRequestPb{}
	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

	maxResultsPb := &st.MaxResults
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}
	maxResultsField := &pb.MaxResults
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	ExternalLocations []ExternalLocationInfo
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string

	ForceSendFields []string
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

	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	for _, item := range pb.ExternalLocations {
		itemField, err := externalLocationInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			externalLocationsField = append(externalLocationsField, *itemField)
		}
	}
	st.ExternalLocations = externalLocationsField
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	CatalogName string
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool
	// Maximum number of functions to return. If not set, all the functions are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int
	// Opaque pagination token to go to next page based on previous query.
	PageToken string
	// Parent schema of functions.
	SchemaName string

	ForceSendFields []string
}

func listFunctionsRequestToPb(st *ListFunctionsRequest) (*listFunctionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFunctionsRequestPb{}
	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

	maxResultsPb := &st.MaxResults
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	schemaNamePb := &st.SchemaName
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

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
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}
	maxResultsField := &pb.MaxResults
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	schemaNameField := &pb.SchemaName
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}

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
	Functions []FunctionInfo
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string

	ForceSendFields []string
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

	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	for _, item := range pb.Functions {
		itemField, err := functionInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			functionsField = append(functionsField, *itemField)
		}
	}
	st.Functions = functionsField
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	for _, item := range pb.Metastores {
		itemField, err := metastoreInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			metastoresField = append(metastoresField, *itemField)
		}
	}
	st.Metastores = metastoresField

	return st, nil
}

// List Model Versions
type ListModelVersionsRequest struct {
	// The full three-level name of the registered model under which to list
	// model versions
	FullName string
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse bool
	// Maximum number of model versions to return. If not set, the page length
	// is set to a server configured value (100, as of 1/3/2024). - when set to
	// a value greater than 0, the page length is the minimum of this value and
	// a server configured value(1000, as of 1/3/2024); - when set to 0, the
	// page length is set to a server configured value (100, as of 1/3/2024)
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int
	// Opaque pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func listModelVersionsRequestToPb(st *ListModelVersionsRequest) (*listModelVersionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listModelVersionsRequestPb{}
	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

	maxResultsPb := &st.MaxResults
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}
	maxResultsField := &pb.MaxResults
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	ModelVersions []ModelVersionInfo
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string

	ForceSendFields []string
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

	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	for _, item := range pb.ModelVersions {
		itemField, err := modelVersionInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			modelVersionsField = append(modelVersionsField, *itemField)
		}
	}
	st.ModelVersions = modelVersionsField
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	MaxResults int
	// Opaque token for the next page of results.
	PageToken string

	ForceSendFields []string
}

func listQuotasRequestToPb(st *ListQuotasRequest) (*listQuotasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQuotasRequestPb{}
	maxResultsPb := &st.MaxResults
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	maxResultsField := &pb.MaxResults
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	NextPageToken string
	// An array of returned QuotaInfos.
	Quotas []QuotaInfo

	ForceSendFields []string
}

func listQuotasResponseToPb(st *ListQuotasResponse) (*listQuotasResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQuotasResponsePb{}
	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	var quotasField []QuotaInfo
	for _, item := range pb.Quotas {
		itemField, err := quotaInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			quotasField = append(quotasField, *itemField)
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
	TableName string
}

func listRefreshesRequestToPb(st *ListRefreshesRequest) (*listRefreshesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRefreshesRequestPb{}
	tableNamePb := &st.TableName
	if tableNamePb != nil {
		pb.TableName = *tableNamePb
	}

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
	tableNameField := &pb.TableName
	if tableNameField != nil {
		st.TableName = *tableNameField
	}

	return st, nil
}

// List Registered Models
type ListRegisteredModelsRequest struct {
	// The identifier of the catalog under which to list registered models. If
	// specified, schema_name must be specified.
	CatalogName string
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool
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
	MaxResults int
	// Opaque token to send for the next page of results (pagination).
	PageToken string
	// The identifier of the schema under which to list registered models. If
	// specified, catalog_name must be specified.
	SchemaName string

	ForceSendFields []string
}

func listRegisteredModelsRequestToPb(st *ListRegisteredModelsRequest) (*listRegisteredModelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRegisteredModelsRequestPb{}
	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

	maxResultsPb := &st.MaxResults
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	schemaNamePb := &st.SchemaName
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

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
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}
	maxResultsField := &pb.MaxResults
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	schemaNameField := &pb.SchemaName
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}

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
	NextPageToken string

	RegisteredModels []RegisteredModelInfo

	ForceSendFields []string
}

func listRegisteredModelsResponseToPb(st *ListRegisteredModelsResponse) (*listRegisteredModelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRegisteredModelsResponsePb{}
	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	var registeredModelsField []RegisteredModelInfo
	for _, item := range pb.RegisteredModels {
		itemField, err := registeredModelInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			registeredModelsField = append(registeredModelsField, *itemField)
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
	CatalogName string
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool
	// Maximum number of schemas to return. If not set, all the schemas are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int
	// Opaque pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func listSchemasRequestToPb(st *ListSchemasRequest) (*listSchemasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSchemasRequestPb{}
	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

	maxResultsPb := &st.MaxResults
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}
	maxResultsField := &pb.MaxResults
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	NextPageToken string
	// An array of schema information objects.
	Schemas []SchemaInfo

	ForceSendFields []string
}

func listSchemasResponseToPb(st *ListSchemasResponse) (*listSchemasResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSchemasResponsePb{}
	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	var schemasField []SchemaInfo
	for _, item := range pb.Schemas {
		itemField, err := schemaInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			schemasField = append(schemasField, *itemField)
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
	MaxResults int
	// Opaque pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func listStorageCredentialsRequestToPb(st *ListStorageCredentialsRequest) (*listStorageCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listStorageCredentialsRequestPb{}
	maxResultsPb := &st.MaxResults
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	maxResultsField := &pb.MaxResults
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	NextPageToken string

	StorageCredentials []StorageCredentialInfo

	ForceSendFields []string
}

func listStorageCredentialsResponseToPb(st *ListStorageCredentialsResponse) (*listStorageCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listStorageCredentialsResponsePb{}
	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	var storageCredentialsField []StorageCredentialInfo
	for _, item := range pb.StorageCredentials {
		itemField, err := storageCredentialInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			storageCredentialsField = append(storageCredentialsField, *itemField)
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
	CatalogName string
	// Whether to include a manifest containing capabilities the table has.
	IncludeManifestCapabilities bool
	// Maximum number of summaries for tables to return. If not set, the page
	// length is set to a server configured value (10000, as of 1/5/2024). -
	// when set to a value greater than 0, the page length is the minimum of
	// this value and a server configured value (10000, as of 1/5/2024); - when
	// set to 0, the page length is set to a server configured value (10000, as
	// of 1/5/2024) (recommended); - when set to a value less than 0, an invalid
	// parameter error is returned;
	MaxResults int
	// Opaque pagination token to go to next page based on previous query.
	PageToken string
	// A sql LIKE pattern (% and _) for schema names. All schemas will be
	// returned if not set or empty.
	SchemaNamePattern string
	// A sql LIKE pattern (% and _) for table names. All tables will be returned
	// if not set or empty.
	TableNamePattern string

	ForceSendFields []string
}

func listSummariesRequestToPb(st *ListSummariesRequest) (*listSummariesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSummariesRequestPb{}
	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	includeManifestCapabilitiesPb := &st.IncludeManifestCapabilities
	if includeManifestCapabilitiesPb != nil {
		pb.IncludeManifestCapabilities = *includeManifestCapabilitiesPb
	}

	maxResultsPb := &st.MaxResults
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	schemaNamePatternPb := &st.SchemaNamePattern
	if schemaNamePatternPb != nil {
		pb.SchemaNamePattern = *schemaNamePatternPb
	}

	tableNamePatternPb := &st.TableNamePattern
	if tableNamePatternPb != nil {
		pb.TableNamePattern = *tableNamePatternPb
	}

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
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	includeManifestCapabilitiesField := &pb.IncludeManifestCapabilities
	if includeManifestCapabilitiesField != nil {
		st.IncludeManifestCapabilities = *includeManifestCapabilitiesField
	}
	maxResultsField := &pb.MaxResults
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	schemaNamePatternField := &pb.SchemaNamePattern
	if schemaNamePatternField != nil {
		st.SchemaNamePattern = *schemaNamePatternField
	}
	tableNamePatternField := &pb.TableNamePattern
	if tableNamePatternField != nil {
		st.TableNamePattern = *tableNamePatternField
	}

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
	MaxResults int
	// The ID for the metastore in which the system schema resides.
	MetastoreId string
	// Opaque pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func listSystemSchemasRequestToPb(st *ListSystemSchemasRequest) (*listSystemSchemasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSystemSchemasRequestPb{}
	maxResultsPb := &st.MaxResults
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	maxResultsField := &pb.MaxResults
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	NextPageToken string
	// An array of system schema information objects.
	Schemas []SystemSchemaInfo

	ForceSendFields []string
}

func listSystemSchemasResponseToPb(st *ListSystemSchemasResponse) (*listSystemSchemasResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSystemSchemasResponsePb{}
	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	var schemasField []SystemSchemaInfo
	for _, item := range pb.Schemas {
		itemField, err := systemSchemaInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			schemasField = append(schemasField, *itemField)
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
	NextPageToken string
	// List of table summaries.
	Tables []TableSummary

	ForceSendFields []string
}

func listTableSummariesResponseToPb(st *ListTableSummariesResponse) (*listTableSummariesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTableSummariesResponsePb{}
	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	var tablesField []TableSummary
	for _, item := range pb.Tables {
		itemField, err := tableSummaryFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			tablesField = append(tablesField, *itemField)
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
	CatalogName string
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata bool
	// Whether to include a manifest containing capabilities the table has.
	IncludeManifestCapabilities bool
	// Maximum number of tables to return. If not set, all the tables are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int
	// Whether to omit the columns of the table from the response or not.
	OmitColumns bool
	// Whether to omit the properties of the table from the response or not.
	OmitProperties bool
	// Whether to omit the username of the table (e.g. owner, updated_by,
	// created_by) from the response or not.
	OmitUsername bool
	// Opaque token to send for the next page of results (pagination).
	PageToken string
	// Parent schema of tables.
	SchemaName string

	ForceSendFields []string
}

func listTablesRequestToPb(st *ListTablesRequest) (*listTablesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTablesRequestPb{}
	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

	includeDeltaMetadataPb := &st.IncludeDeltaMetadata
	if includeDeltaMetadataPb != nil {
		pb.IncludeDeltaMetadata = *includeDeltaMetadataPb
	}

	includeManifestCapabilitiesPb := &st.IncludeManifestCapabilities
	if includeManifestCapabilitiesPb != nil {
		pb.IncludeManifestCapabilities = *includeManifestCapabilitiesPb
	}

	maxResultsPb := &st.MaxResults
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	omitColumnsPb := &st.OmitColumns
	if omitColumnsPb != nil {
		pb.OmitColumns = *omitColumnsPb
	}

	omitPropertiesPb := &st.OmitProperties
	if omitPropertiesPb != nil {
		pb.OmitProperties = *omitPropertiesPb
	}

	omitUsernamePb := &st.OmitUsername
	if omitUsernamePb != nil {
		pb.OmitUsername = *omitUsernamePb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	schemaNamePb := &st.SchemaName
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

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
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}
	includeDeltaMetadataField := &pb.IncludeDeltaMetadata
	if includeDeltaMetadataField != nil {
		st.IncludeDeltaMetadata = *includeDeltaMetadataField
	}
	includeManifestCapabilitiesField := &pb.IncludeManifestCapabilities
	if includeManifestCapabilitiesField != nil {
		st.IncludeManifestCapabilities = *includeManifestCapabilitiesField
	}
	maxResultsField := &pb.MaxResults
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	omitColumnsField := &pb.OmitColumns
	if omitColumnsField != nil {
		st.OmitColumns = *omitColumnsField
	}
	omitPropertiesField := &pb.OmitProperties
	if omitPropertiesField != nil {
		st.OmitProperties = *omitPropertiesField
	}
	omitUsernameField := &pb.OmitUsername
	if omitUsernameField != nil {
		st.OmitUsername = *omitUsernameField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	schemaNameField := &pb.SchemaName
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}

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
	NextPageToken string
	// An array of table information objects.
	Tables []TableInfo

	ForceSendFields []string
}

func listTablesResponseToPb(st *ListTablesResponse) (*listTablesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTablesResponsePb{}
	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	var tablesField []TableInfo
	for _, item := range pb.Tables {
		itemField, err := tableInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			tablesField = append(tablesField, *itemField)
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
	CatalogName string
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool
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
	MaxResults int
	// Opaque token returned by a previous request. It must be included in the
	// request to retrieve the next page of results (pagination).
	PageToken string
	// The identifier of the schema
	SchemaName string

	ForceSendFields []string
}

func listVolumesRequestToPb(st *ListVolumesRequest) (*listVolumesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listVolumesRequestPb{}
	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

	maxResultsPb := &st.MaxResults
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	schemaNamePb := &st.SchemaName
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

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
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}
	maxResultsField := &pb.MaxResults
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	schemaNameField := &pb.SchemaName
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}

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
	NextPageToken string

	Volumes []VolumeInfo

	ForceSendFields []string
}

func listVolumesResponseContentToPb(st *ListVolumesResponseContent) (*listVolumesResponseContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listVolumesResponseContentPb{}
	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	var volumesField []VolumeInfo
	for _, item := range pb.Volumes {
		itemField, err := volumeInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			volumesField = append(volumesField, *itemField)
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
	DefaultCatalogName string
	// The unique ID of the metastore.
	MetastoreId string
	// The unique ID of the Databricks workspace.
	WorkspaceId int64

	ForceSendFields []string
}

func metastoreAssignmentToPb(st *MetastoreAssignment) (*metastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &metastoreAssignmentPb{}
	defaultCatalogNamePb := &st.DefaultCatalogName
	if defaultCatalogNamePb != nil {
		pb.DefaultCatalogName = *defaultCatalogNamePb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	workspaceIdPb := &st.WorkspaceId
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

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
	defaultCatalogNameField := &pb.DefaultCatalogName
	if defaultCatalogNameField != nil {
		st.DefaultCatalogName = *defaultCatalogNameField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	workspaceIdField := &pb.WorkspaceId
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}

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
	Cloud string
	// Time at which this metastore was created, in epoch milliseconds.
	CreatedAt int64
	// Username of metastore creator.
	CreatedBy string
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	DefaultDataAccessConfigId string
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	DeltaSharingOrganizationName string
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope MetastoreInfoDeltaSharingScope
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	ExternalAccessEnabled bool
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	GlobalMetastoreId string
	// Unique identifier of metastore.
	MetastoreId string
	// The user-specified name of the metastore.
	Name string
	// The owner of the metastore.
	Owner string
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	Region string
	// The storage root URL for metastore
	StorageRoot string
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string
	// Name of the storage credential to access the metastore storage_root.
	StorageRootCredentialName string
	// Time at which the metastore was last modified, in epoch milliseconds.
	UpdatedAt int64
	// Username of user who last modified the metastore.
	UpdatedBy string

	ForceSendFields []string
}

func metastoreInfoToPb(st *MetastoreInfo) (*metastoreInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &metastoreInfoPb{}
	cloudPb := &st.Cloud
	if cloudPb != nil {
		pb.Cloud = *cloudPb
	}

	createdAtPb := &st.CreatedAt
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	defaultDataAccessConfigIdPb := &st.DefaultDataAccessConfigId
	if defaultDataAccessConfigIdPb != nil {
		pb.DefaultDataAccessConfigId = *defaultDataAccessConfigIdPb
	}

	deltaSharingOrganizationNamePb := &st.DeltaSharingOrganizationName
	if deltaSharingOrganizationNamePb != nil {
		pb.DeltaSharingOrganizationName = *deltaSharingOrganizationNamePb
	}

	deltaSharingRecipientTokenLifetimeInSecondsPb := &st.DeltaSharingRecipientTokenLifetimeInSeconds
	if deltaSharingRecipientTokenLifetimeInSecondsPb != nil {
		pb.DeltaSharingRecipientTokenLifetimeInSeconds = *deltaSharingRecipientTokenLifetimeInSecondsPb
	}

	deltaSharingScopePb := &st.DeltaSharingScope
	if deltaSharingScopePb != nil {
		pb.DeltaSharingScope = *deltaSharingScopePb
	}

	externalAccessEnabledPb := &st.ExternalAccessEnabled
	if externalAccessEnabledPb != nil {
		pb.ExternalAccessEnabled = *externalAccessEnabledPb
	}

	globalMetastoreIdPb := &st.GlobalMetastoreId
	if globalMetastoreIdPb != nil {
		pb.GlobalMetastoreId = *globalMetastoreIdPb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	privilegeModelVersionPb := &st.PrivilegeModelVersion
	if privilegeModelVersionPb != nil {
		pb.PrivilegeModelVersion = *privilegeModelVersionPb
	}

	regionPb := &st.Region
	if regionPb != nil {
		pb.Region = *regionPb
	}

	storageRootPb := &st.StorageRoot
	if storageRootPb != nil {
		pb.StorageRoot = *storageRootPb
	}

	storageRootCredentialIdPb := &st.StorageRootCredentialId
	if storageRootCredentialIdPb != nil {
		pb.StorageRootCredentialId = *storageRootCredentialIdPb
	}

	storageRootCredentialNamePb := &st.StorageRootCredentialName
	if storageRootCredentialNamePb != nil {
		pb.StorageRootCredentialName = *storageRootCredentialNamePb
	}

	updatedAtPb := &st.UpdatedAt
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb := &st.UpdatedBy
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

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
	cloudField := &pb.Cloud
	if cloudField != nil {
		st.Cloud = *cloudField
	}
	createdAtField := &pb.CreatedAt
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	defaultDataAccessConfigIdField := &pb.DefaultDataAccessConfigId
	if defaultDataAccessConfigIdField != nil {
		st.DefaultDataAccessConfigId = *defaultDataAccessConfigIdField
	}
	deltaSharingOrganizationNameField := &pb.DeltaSharingOrganizationName
	if deltaSharingOrganizationNameField != nil {
		st.DeltaSharingOrganizationName = *deltaSharingOrganizationNameField
	}
	deltaSharingRecipientTokenLifetimeInSecondsField := &pb.DeltaSharingRecipientTokenLifetimeInSeconds
	if deltaSharingRecipientTokenLifetimeInSecondsField != nil {
		st.DeltaSharingRecipientTokenLifetimeInSeconds = *deltaSharingRecipientTokenLifetimeInSecondsField
	}
	deltaSharingScopeField := &pb.DeltaSharingScope
	if deltaSharingScopeField != nil {
		st.DeltaSharingScope = *deltaSharingScopeField
	}
	externalAccessEnabledField := &pb.ExternalAccessEnabled
	if externalAccessEnabledField != nil {
		st.ExternalAccessEnabled = *externalAccessEnabledField
	}
	globalMetastoreIdField := &pb.GlobalMetastoreId
	if globalMetastoreIdField != nil {
		st.GlobalMetastoreId = *globalMetastoreIdField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}
	privilegeModelVersionField := &pb.PrivilegeModelVersion
	if privilegeModelVersionField != nil {
		st.PrivilegeModelVersion = *privilegeModelVersionField
	}
	regionField := &pb.Region
	if regionField != nil {
		st.Region = *regionField
	}
	storageRootField := &pb.StorageRoot
	if storageRootField != nil {
		st.StorageRoot = *storageRootField
	}
	storageRootCredentialIdField := &pb.StorageRootCredentialId
	if storageRootCredentialIdField != nil {
		st.StorageRootCredentialId = *storageRootCredentialIdField
	}
	storageRootCredentialNameField := &pb.StorageRootCredentialName
	if storageRootCredentialNameField != nil {
		st.StorageRootCredentialName = *storageRootCredentialNameField
	}
	updatedAtField := &pb.UpdatedAt
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField := &pb.UpdatedBy
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}

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
	Aliases []RegisteredModelAlias
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool
	// The name of the catalog containing the model version
	CatalogName string
	// The comment attached to the model version
	Comment string

	CreatedAt int64
	// The identifier of the user who created the model version
	CreatedBy string
	// The unique identifier of the model version
	Id string
	// The unique identifier of the metastore containing the model version
	MetastoreId string
	// The name of the parent registered model of the model version, relative to
	// parent schema
	ModelName string
	// Model version dependencies, for feature-store packaged models
	ModelVersionDependencies *DependencyList
	// MLflow run ID used when creating the model version, if ``source`` was
	// generated by an experiment run stored in an MLflow tracking server
	RunId string
	// ID of the Databricks workspace containing the MLflow run that generated
	// this model version, if applicable
	RunWorkspaceId int
	// The name of the schema containing the model version, relative to parent
	// catalog
	SchemaName string
	// URI indicating the location of the source artifacts (files) for the model
	// version
	Source string
	// Current status of the model version. Newly created model versions start
	// in PENDING_REGISTRATION status, then move to READY status once the model
	// version files are uploaded and the model version is finalized. Only model
	// versions in READY status can be loaded for inference or served.
	Status ModelVersionInfoStatus
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string

	UpdatedAt int64
	// The identifier of the user who updated the model version last time
	UpdatedBy string
	// Integer model version number, used to reference the model version in API
	// requests.
	Version int

	ForceSendFields []string
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

	browseOnlyPb := &st.BrowseOnly
	if browseOnlyPb != nil {
		pb.BrowseOnly = *browseOnlyPb
	}

	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	createdAtPb := &st.CreatedAt
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	idPb := &st.Id
	if idPb != nil {
		pb.Id = *idPb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	modelNamePb := &st.ModelName
	if modelNamePb != nil {
		pb.ModelName = *modelNamePb
	}

	modelVersionDependenciesPb, err := dependencyListToPb(st.ModelVersionDependencies)
	if err != nil {
		return nil, err
	}
	if modelVersionDependenciesPb != nil {
		pb.ModelVersionDependencies = modelVersionDependenciesPb
	}

	runIdPb := &st.RunId
	if runIdPb != nil {
		pb.RunId = *runIdPb
	}

	runWorkspaceIdPb := &st.RunWorkspaceId
	if runWorkspaceIdPb != nil {
		pb.RunWorkspaceId = *runWorkspaceIdPb
	}

	schemaNamePb := &st.SchemaName
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

	sourcePb := &st.Source
	if sourcePb != nil {
		pb.Source = *sourcePb
	}

	statusPb := &st.Status
	if statusPb != nil {
		pb.Status = *statusPb
	}

	storageLocationPb := &st.StorageLocation
	if storageLocationPb != nil {
		pb.StorageLocation = *storageLocationPb
	}

	updatedAtPb := &st.UpdatedAt
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb := &st.UpdatedBy
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

	versionPb := &st.Version
	if versionPb != nil {
		pb.Version = *versionPb
	}

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
	for _, item := range pb.Aliases {
		itemField, err := registeredModelAliasFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			aliasesField = append(aliasesField, *itemField)
		}
	}
	st.Aliases = aliasesField
	browseOnlyField := &pb.BrowseOnly
	if browseOnlyField != nil {
		st.BrowseOnly = *browseOnlyField
	}
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	createdAtField := &pb.CreatedAt
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	idField := &pb.Id
	if idField != nil {
		st.Id = *idField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	modelNameField := &pb.ModelName
	if modelNameField != nil {
		st.ModelName = *modelNameField
	}
	modelVersionDependenciesField, err := dependencyListFromPb(pb.ModelVersionDependencies)
	if err != nil {
		return nil, err
	}
	if modelVersionDependenciesField != nil {
		st.ModelVersionDependencies = modelVersionDependenciesField
	}
	runIdField := &pb.RunId
	if runIdField != nil {
		st.RunId = *runIdField
	}
	runWorkspaceIdField := &pb.RunWorkspaceId
	if runWorkspaceIdField != nil {
		st.RunWorkspaceId = *runWorkspaceIdField
	}
	schemaNameField := &pb.SchemaName
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}
	sourceField := &pb.Source
	if sourceField != nil {
		st.Source = *sourceField
	}
	statusField := &pb.Status
	if statusField != nil {
		st.Status = *statusField
	}
	storageLocationField := &pb.StorageLocation
	if storageLocationField != nil {
		st.StorageLocation = *storageLocationField
	}
	updatedAtField := &pb.UpdatedAt
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField := &pb.UpdatedBy
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}
	versionField := &pb.Version
	if versionField != nil {
		st.Version = *versionField
	}

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
	PauseStatus MonitorCronSchedulePauseStatus
	// The expression that determines when to run the monitor. See [examples].
	//
	// [examples]: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression string
	// The timezone id (e.g., ``"PST"``) in which to evaluate the quartz
	// expression.
	TimezoneId string
}

func monitorCronScheduleToPb(st *MonitorCronSchedule) (*monitorCronSchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorCronSchedulePb{}
	pauseStatusPb := &st.PauseStatus
	if pauseStatusPb != nil {
		pb.PauseStatus = *pauseStatusPb
	}

	quartzCronExpressionPb := &st.QuartzCronExpression
	if quartzCronExpressionPb != nil {
		pb.QuartzCronExpression = *quartzCronExpressionPb
	}

	timezoneIdPb := &st.TimezoneId
	if timezoneIdPb != nil {
		pb.TimezoneId = *timezoneIdPb
	}

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
	pauseStatusField := &pb.PauseStatus
	if pauseStatusField != nil {
		st.PauseStatus = *pauseStatusField
	}
	quartzCronExpressionField := &pb.QuartzCronExpression
	if quartzCronExpressionField != nil {
		st.QuartzCronExpression = *quartzCronExpressionField
	}
	timezoneIdField := &pb.TimezoneId
	if timezoneIdField != nil {
		st.TimezoneId = *timezoneIdField
	}

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
	Enabled bool

	ForceSendFields []string
}

func monitorDataClassificationConfigToPb(st *MonitorDataClassificationConfig) (*monitorDataClassificationConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorDataClassificationConfigPb{}
	enabledPb := &st.Enabled
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

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
	enabledField := &pb.Enabled
	if enabledField != nil {
		st.Enabled = *enabledField
	}

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
	EmailAddresses []string
}

func monitorDestinationToPb(st *MonitorDestination) (*monitorDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorDestinationPb{}

	var emailAddressesPb []string
	for _, item := range st.EmailAddresses {
		itemPb := &item
		if itemPb != nil {
			emailAddressesPb = append(emailAddressesPb, *itemPb)
		}
	}
	pb.EmailAddresses = emailAddressesPb

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

	var emailAddressesField []string
	for _, item := range pb.EmailAddresses {
		itemField := &item
		if itemField != nil {
			emailAddressesField = append(emailAddressesField, *itemField)
		}
	}
	st.EmailAddresses = emailAddressesField

	return st, nil
}

type MonitorInferenceLog struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	Granularities []string
	// Optional column that contains the ground truth for the prediction.
	LabelCol string
	// Column that contains the id of the model generating the predictions.
	// Metrics will be computed per model id by default, and also across all
	// model ids.
	ModelIdCol string
	// Column that contains the output/prediction from the model.
	PredictionCol string
	// Optional column that contains the prediction probabilities for each class
	// in a classification problem type. The values in this column should be a
	// map, mapping each class label to the prediction probability for a given
	// sample. The map should be of PySpark MapType().
	PredictionProbaCol string
	// Problem type the model aims to solve. Determines the type of
	// model-quality metrics that will be computed.
	ProblemType MonitorInferenceLogProblemType
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	TimestampCol string

	ForceSendFields []string
}

func monitorInferenceLogToPb(st *MonitorInferenceLog) (*monitorInferenceLogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorInferenceLogPb{}

	var granularitiesPb []string
	for _, item := range st.Granularities {
		itemPb := &item
		if itemPb != nil {
			granularitiesPb = append(granularitiesPb, *itemPb)
		}
	}
	pb.Granularities = granularitiesPb

	labelColPb := &st.LabelCol
	if labelColPb != nil {
		pb.LabelCol = *labelColPb
	}

	modelIdColPb := &st.ModelIdCol
	if modelIdColPb != nil {
		pb.ModelIdCol = *modelIdColPb
	}

	predictionColPb := &st.PredictionCol
	if predictionColPb != nil {
		pb.PredictionCol = *predictionColPb
	}

	predictionProbaColPb := &st.PredictionProbaCol
	if predictionProbaColPb != nil {
		pb.PredictionProbaCol = *predictionProbaColPb
	}

	problemTypePb := &st.ProblemType
	if problemTypePb != nil {
		pb.ProblemType = *problemTypePb
	}

	timestampColPb := &st.TimestampCol
	if timestampColPb != nil {
		pb.TimestampCol = *timestampColPb
	}

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

	var granularitiesField []string
	for _, item := range pb.Granularities {
		itemField := &item
		if itemField != nil {
			granularitiesField = append(granularitiesField, *itemField)
		}
	}
	st.Granularities = granularitiesField
	labelColField := &pb.LabelCol
	if labelColField != nil {
		st.LabelCol = *labelColField
	}
	modelIdColField := &pb.ModelIdCol
	if modelIdColField != nil {
		st.ModelIdCol = *modelIdColField
	}
	predictionColField := &pb.PredictionCol
	if predictionColField != nil {
		st.PredictionCol = *predictionColField
	}
	predictionProbaColField := &pb.PredictionProbaCol
	if predictionProbaColField != nil {
		st.PredictionProbaCol = *predictionProbaColField
	}
	problemTypeField := &pb.ProblemType
	if problemTypeField != nil {
		st.ProblemType = *problemTypeField
	}
	timestampColField := &pb.TimestampCol
	if timestampColField != nil {
		st.TimestampCol = *timestampColField
	}

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
	AssetsDir string
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	BaselineTableName string
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics []MonitorMetric
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId string
	// The data classification config for the monitor.
	DataClassificationConfig *MonitorDataClassificationConfig
	// The full name of the drift metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	DriftMetricsTableName string
	// Configuration for monitoring inference logs.
	InferenceLog *MonitorInferenceLog
	// The latest failure message of the monitor (if any).
	LatestMonitorFailureMsg string
	// The version of the monitor config (e.g. 1,2,3). If negative, the monitor
	// may be corrupted.
	MonitorVersion string
	// The notification settings for the monitor.
	Notifications *MonitorNotifications
	// Schema where output metric tables are created.
	OutputSchemaName string
	// The full name of the profile metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	ProfileMetricsTableName string
	// The schedule for automatically updating and refreshing metric tables.
	Schedule *MonitorCronSchedule
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs []string
	// Configuration for monitoring snapshot tables.
	Snapshot *MonitorSnapshot
	// The status of the monitor.
	Status MonitorInfoStatus
	// The full name of the table to monitor. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	TableName string
	// Configuration for monitoring time series tables.
	TimeSeries *MonitorTimeSeries

	ForceSendFields []string
}

func monitorInfoToPb(st *MonitorInfo) (*monitorInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorInfoPb{}
	assetsDirPb := &st.AssetsDir
	if assetsDirPb != nil {
		pb.AssetsDir = *assetsDirPb
	}

	baselineTableNamePb := &st.BaselineTableName
	if baselineTableNamePb != nil {
		pb.BaselineTableName = *baselineTableNamePb
	}

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

	dashboardIdPb := &st.DashboardId
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	dataClassificationConfigPb, err := monitorDataClassificationConfigToPb(st.DataClassificationConfig)
	if err != nil {
		return nil, err
	}
	if dataClassificationConfigPb != nil {
		pb.DataClassificationConfig = dataClassificationConfigPb
	}

	driftMetricsTableNamePb := &st.DriftMetricsTableName
	if driftMetricsTableNamePb != nil {
		pb.DriftMetricsTableName = *driftMetricsTableNamePb
	}

	inferenceLogPb, err := monitorInferenceLogToPb(st.InferenceLog)
	if err != nil {
		return nil, err
	}
	if inferenceLogPb != nil {
		pb.InferenceLog = inferenceLogPb
	}

	latestMonitorFailureMsgPb := &st.LatestMonitorFailureMsg
	if latestMonitorFailureMsgPb != nil {
		pb.LatestMonitorFailureMsg = *latestMonitorFailureMsgPb
	}

	monitorVersionPb := &st.MonitorVersion
	if monitorVersionPb != nil {
		pb.MonitorVersion = *monitorVersionPb
	}

	notificationsPb, err := monitorNotificationsToPb(st.Notifications)
	if err != nil {
		return nil, err
	}
	if notificationsPb != nil {
		pb.Notifications = notificationsPb
	}

	outputSchemaNamePb := &st.OutputSchemaName
	if outputSchemaNamePb != nil {
		pb.OutputSchemaName = *outputSchemaNamePb
	}

	profileMetricsTableNamePb := &st.ProfileMetricsTableName
	if profileMetricsTableNamePb != nil {
		pb.ProfileMetricsTableName = *profileMetricsTableNamePb
	}

	schedulePb, err := monitorCronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}

	var slicingExprsPb []string
	for _, item := range st.SlicingExprs {
		itemPb := &item
		if itemPb != nil {
			slicingExprsPb = append(slicingExprsPb, *itemPb)
		}
	}
	pb.SlicingExprs = slicingExprsPb

	snapshotPb, err := monitorSnapshotToPb(st.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotPb != nil {
		pb.Snapshot = snapshotPb
	}

	statusPb := &st.Status
	if statusPb != nil {
		pb.Status = *statusPb
	}

	tableNamePb := &st.TableName
	if tableNamePb != nil {
		pb.TableName = *tableNamePb
	}

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
	assetsDirField := &pb.AssetsDir
	if assetsDirField != nil {
		st.AssetsDir = *assetsDirField
	}
	baselineTableNameField := &pb.BaselineTableName
	if baselineTableNameField != nil {
		st.BaselineTableName = *baselineTableNameField
	}

	var customMetricsField []MonitorMetric
	for _, item := range pb.CustomMetrics {
		itemField, err := monitorMetricFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			customMetricsField = append(customMetricsField, *itemField)
		}
	}
	st.CustomMetrics = customMetricsField
	dashboardIdField := &pb.DashboardId
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	dataClassificationConfigField, err := monitorDataClassificationConfigFromPb(pb.DataClassificationConfig)
	if err != nil {
		return nil, err
	}
	if dataClassificationConfigField != nil {
		st.DataClassificationConfig = dataClassificationConfigField
	}
	driftMetricsTableNameField := &pb.DriftMetricsTableName
	if driftMetricsTableNameField != nil {
		st.DriftMetricsTableName = *driftMetricsTableNameField
	}
	inferenceLogField, err := monitorInferenceLogFromPb(pb.InferenceLog)
	if err != nil {
		return nil, err
	}
	if inferenceLogField != nil {
		st.InferenceLog = inferenceLogField
	}
	latestMonitorFailureMsgField := &pb.LatestMonitorFailureMsg
	if latestMonitorFailureMsgField != nil {
		st.LatestMonitorFailureMsg = *latestMonitorFailureMsgField
	}
	monitorVersionField := &pb.MonitorVersion
	if monitorVersionField != nil {
		st.MonitorVersion = *monitorVersionField
	}
	notificationsField, err := monitorNotificationsFromPb(pb.Notifications)
	if err != nil {
		return nil, err
	}
	if notificationsField != nil {
		st.Notifications = notificationsField
	}
	outputSchemaNameField := &pb.OutputSchemaName
	if outputSchemaNameField != nil {
		st.OutputSchemaName = *outputSchemaNameField
	}
	profileMetricsTableNameField := &pb.ProfileMetricsTableName
	if profileMetricsTableNameField != nil {
		st.ProfileMetricsTableName = *profileMetricsTableNameField
	}
	scheduleField, err := monitorCronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}

	var slicingExprsField []string
	for _, item := range pb.SlicingExprs {
		itemField := &item
		if itemField != nil {
			slicingExprsField = append(slicingExprsField, *itemField)
		}
	}
	st.SlicingExprs = slicingExprsField
	snapshotField, err := monitorSnapshotFromPb(pb.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotField != nil {
		st.Snapshot = snapshotField
	}
	statusField := &pb.Status
	if statusField != nil {
		st.Status = *statusField
	}
	tableNameField := &pb.TableName
	if tableNameField != nil {
		st.TableName = *tableNameField
	}
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
	Definition string
	// A list of column names in the input table the metric should be computed
	// for. Can use ``":table"`` to indicate that the metric needs information
	// from multiple columns.
	InputColumns []string
	// Name of the metric in the output tables.
	Name string
	// The output type of the custom metric.
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
	Type MonitorMetricType
}

func monitorMetricToPb(st *MonitorMetric) (*monitorMetricPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorMetricPb{}
	definitionPb := &st.Definition
	if definitionPb != nil {
		pb.Definition = *definitionPb
	}

	var inputColumnsPb []string
	for _, item := range st.InputColumns {
		itemPb := &item
		if itemPb != nil {
			inputColumnsPb = append(inputColumnsPb, *itemPb)
		}
	}
	pb.InputColumns = inputColumnsPb

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	outputDataTypePb := &st.OutputDataType
	if outputDataTypePb != nil {
		pb.OutputDataType = *outputDataTypePb
	}

	typePb := &st.Type
	if typePb != nil {
		pb.Type = *typePb
	}

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
	definitionField := &pb.Definition
	if definitionField != nil {
		st.Definition = *definitionField
	}

	var inputColumnsField []string
	for _, item := range pb.InputColumns {
		itemField := &item
		if itemField != nil {
			inputColumnsField = append(inputColumnsField, *itemField)
		}
	}
	st.InputColumns = inputColumnsField
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	outputDataTypeField := &pb.OutputDataType
	if outputDataTypeField != nil {
		st.OutputDataType = *outputDataTypeField
	}
	typeField := &pb.Type
	if typeField != nil {
		st.Type = *typeField
	}

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
	OnFailure *MonitorDestination
	// Who to send notifications to when new data classification tags are
	// detected.
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
	EndTimeMs int64
	// An optional message to give insight into the current state of the job
	// (e.g. FAILURE messages).
	Message string
	// Unique id of the refresh operation.
	RefreshId int64
	// Time at which refresh operation was initiated (milliseconds since
	// 1/1/1970 UTC).
	StartTimeMs int64
	// The current state of the refresh.
	State MonitorRefreshInfoState
	// The method by which the refresh was triggered.
	Trigger MonitorRefreshInfoTrigger

	ForceSendFields []string
}

func monitorRefreshInfoToPb(st *MonitorRefreshInfo) (*monitorRefreshInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorRefreshInfoPb{}
	endTimeMsPb := &st.EndTimeMs
	if endTimeMsPb != nil {
		pb.EndTimeMs = *endTimeMsPb
	}

	messagePb := &st.Message
	if messagePb != nil {
		pb.Message = *messagePb
	}

	refreshIdPb := &st.RefreshId
	if refreshIdPb != nil {
		pb.RefreshId = *refreshIdPb
	}

	startTimeMsPb := &st.StartTimeMs
	if startTimeMsPb != nil {
		pb.StartTimeMs = *startTimeMsPb
	}

	statePb := &st.State
	if statePb != nil {
		pb.State = *statePb
	}

	triggerPb := &st.Trigger
	if triggerPb != nil {
		pb.Trigger = *triggerPb
	}

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
	endTimeMsField := &pb.EndTimeMs
	if endTimeMsField != nil {
		st.EndTimeMs = *endTimeMsField
	}
	messageField := &pb.Message
	if messageField != nil {
		st.Message = *messageField
	}
	refreshIdField := &pb.RefreshId
	if refreshIdField != nil {
		st.RefreshId = *refreshIdField
	}
	startTimeMsField := &pb.StartTimeMs
	if startTimeMsField != nil {
		st.StartTimeMs = *startTimeMsField
	}
	stateField := &pb.State
	if stateField != nil {
		st.State = *stateField
	}
	triggerField := &pb.Trigger
	if triggerField != nil {
		st.Trigger = *triggerField
	}

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
	for _, item := range pb.Refreshes {
		itemField, err := monitorRefreshInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			refreshesField = append(refreshesField, *itemField)
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
	Granularities []string
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	TimestampCol string
}

func monitorTimeSeriesToPb(st *MonitorTimeSeries) (*monitorTimeSeriesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorTimeSeriesPb{}

	var granularitiesPb []string
	for _, item := range st.Granularities {
		itemPb := &item
		if itemPb != nil {
			granularitiesPb = append(granularitiesPb, *itemPb)
		}
	}
	pb.Granularities = granularitiesPb

	timestampColPb := &st.TimestampCol
	if timestampColPb != nil {
		pb.TimestampCol = *timestampColPb
	}

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

	var granularitiesField []string
	for _, item := range pb.Granularities {
		itemField := &item
		if itemField != nil {
			granularitiesField = append(granularitiesField, *itemField)
		}
	}
	st.Granularities = granularitiesField
	timestampColField := &pb.TimestampCol
	if timestampColField != nil {
		st.TimestampCol = *timestampColField
	}

	return st, nil
}

type NamedTableConstraint struct {
	// The name of the constraint.
	Name string
}

func namedTableConstraintToPb(st *NamedTableConstraint) (*namedTableConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &namedTableConstraintPb{}
	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

// Online Table information.
type OnlineTable struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string
	// Specification of the online table.
	Spec *OnlineTableSpec
	// Online Table data synchronization status
	Status *OnlineTableStatus
	// Data serving REST API URL for this table
	TableServingUrl string
	// The provisioning state of the online table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	UnityCatalogProvisioningState ProvisioningInfoState

	ForceSendFields []string
}

func onlineTableToPb(st *OnlineTable) (*onlineTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &onlineTablePb{}
	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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

	tableServingUrlPb := &st.TableServingUrl
	if tableServingUrlPb != nil {
		pb.TableServingUrl = *tableServingUrlPb
	}

	unityCatalogProvisioningStatePb := &st.UnityCatalogProvisioningState
	if unityCatalogProvisioningStatePb != nil {
		pb.UnityCatalogProvisioningState = *unityCatalogProvisioningStatePb
	}

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
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
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
	tableServingUrlField := &pb.TableServingUrl
	if tableServingUrlField != nil {
		st.TableServingUrl = *tableServingUrlField
	}
	unityCatalogProvisioningStateField := &pb.UnityCatalogProvisioningState
	if unityCatalogProvisioningStateField != nil {
		st.UnityCatalogProvisioningState = *unityCatalogProvisioningStateField
	}

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
	PerformFullCopy bool
	// ID of the associated pipeline. Generated by the server - cannot be set by
	// the caller.
	PipelineId string
	// Primary Key columns to be used for data insert/update in the destination.
	PrimaryKeyColumns []string
	// Pipeline runs continuously after generating the initial data.
	RunContinuously *OnlineTableSpecContinuousSchedulingPolicy
	// Pipeline stops after generating the initial data and can be triggered
	// later (manually, through a cron job or through data triggers)
	RunTriggered *OnlineTableSpecTriggeredSchedulingPolicy
	// Three-part (catalog, schema, table) name of the source Delta table.
	SourceTableFullName string
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	TimeseriesKey string

	ForceSendFields []string
}

func onlineTableSpecToPb(st *OnlineTableSpec) (*onlineTableSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &onlineTableSpecPb{}
	performFullCopyPb := &st.PerformFullCopy
	if performFullCopyPb != nil {
		pb.PerformFullCopy = *performFullCopyPb
	}

	pipelineIdPb := &st.PipelineId
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

	var primaryKeyColumnsPb []string
	for _, item := range st.PrimaryKeyColumns {
		itemPb := &item
		if itemPb != nil {
			primaryKeyColumnsPb = append(primaryKeyColumnsPb, *itemPb)
		}
	}
	pb.PrimaryKeyColumns = primaryKeyColumnsPb

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

	sourceTableFullNamePb := &st.SourceTableFullName
	if sourceTableFullNamePb != nil {
		pb.SourceTableFullName = *sourceTableFullNamePb
	}

	timeseriesKeyPb := &st.TimeseriesKey
	if timeseriesKeyPb != nil {
		pb.TimeseriesKey = *timeseriesKeyPb
	}

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
	performFullCopyField := &pb.PerformFullCopy
	if performFullCopyField != nil {
		st.PerformFullCopy = *performFullCopyField
	}
	pipelineIdField := &pb.PipelineId
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}

	var primaryKeyColumnsField []string
	for _, item := range pb.PrimaryKeyColumns {
		itemField := &item
		if itemField != nil {
			primaryKeyColumnsField = append(primaryKeyColumnsField, *itemField)
		}
	}
	st.PrimaryKeyColumns = primaryKeyColumnsField
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
	sourceTableFullNameField := &pb.SourceTableFullName
	if sourceTableFullNameField != nil {
		st.SourceTableFullName = *sourceTableFullNameField
	}
	timeseriesKeyField := &pb.TimeseriesKey
	if timeseriesKeyField != nil {
		st.TimeseriesKey = *timeseriesKeyField
	}

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
	ContinuousUpdateStatus *ContinuousUpdateStatus
	// The state of the online table.
	DetailedState OnlineTableState
	// Detailed status of an online table. Shown if the online table is in the
	// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
	FailedStatus *FailedStatus
	// A text description of the current state of the online table.
	Message string
	// Detailed status of an online table. Shown if the online table is in the
	// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT
	// state.
	ProvisioningStatus *ProvisioningStatus
	// Detailed status of an online table. Shown if the online table is in the
	// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
	TriggeredUpdateStatus *TriggeredUpdateStatus

	ForceSendFields []string
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

	detailedStatePb := &st.DetailedState
	if detailedStatePb != nil {
		pb.DetailedState = *detailedStatePb
	}

	failedStatusPb, err := failedStatusToPb(st.FailedStatus)
	if err != nil {
		return nil, err
	}
	if failedStatusPb != nil {
		pb.FailedStatus = failedStatusPb
	}

	messagePb := &st.Message
	if messagePb != nil {
		pb.Message = *messagePb
	}

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
	detailedStateField := &pb.DetailedState
	if detailedStateField != nil {
		st.DetailedState = *detailedStateField
	}
	failedStatusField, err := failedStatusFromPb(pb.FailedStatus)
	if err != nil {
		return nil, err
	}
	if failedStatusField != nil {
		st.FailedStatus = failedStatusField
	}
	messageField := &pb.Message
	if messageField != nil {
		st.Message = *messageField
	}
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
	Add []Privilege
	// The principal whose privileges we are changing.
	Principal string
	// The set of privileges to remove.
	Remove []Privilege

	ForceSendFields []string
}

func permissionsChangeToPb(st *PermissionsChange) (*permissionsChangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permissionsChangePb{}

	var addPb []Privilege
	for _, item := range st.Add {
		itemPb := &item
		if itemPb != nil {
			addPb = append(addPb, *itemPb)
		}
	}
	pb.Add = addPb

	principalPb := &st.Principal
	if principalPb != nil {
		pb.Principal = *principalPb
	}

	var removePb []Privilege
	for _, item := range st.Remove {
		itemPb := &item
		if itemPb != nil {
			removePb = append(removePb, *itemPb)
		}
	}
	pb.Remove = removePb

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

	var addField []Privilege
	for _, item := range pb.Add {
		itemField := &item
		if itemField != nil {
			addField = append(addField, *itemField)
		}
	}
	st.Add = addField
	principalField := &pb.Principal
	if principalField != nil {
		st.Principal = *principalField
	}

	var removeField []Privilege
	for _, item := range pb.Remove {
		itemField := &item
		if itemField != nil {
			removeField = append(removeField, *itemField)
		}
	}
	st.Remove = removeField

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
	for _, item := range pb.PrivilegeAssignments {
		itemField, err := privilegeAssignmentFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			privilegeAssignmentsField = append(privilegeAssignmentsField, *itemField)
		}
	}
	st.PrivilegeAssignments = privilegeAssignmentsField

	return st, nil
}

// Progress information of the Online Table data synchronization pipeline.
type PipelineProgress struct {
	// The estimated time remaining to complete this update in seconds.
	EstimatedCompletionTimeSeconds float64
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	LatestVersionCurrentlyProcessing int64
	// The completion ratio of this update. This is a number between 0 and 1.
	SyncProgressCompletion float64
	// The number of rows that have been synced in this update.
	SyncedRowCount int64
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	TotalRowCount int64

	ForceSendFields []string
}

func pipelineProgressToPb(st *PipelineProgress) (*pipelineProgressPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineProgressPb{}
	estimatedCompletionTimeSecondsPb := &st.EstimatedCompletionTimeSeconds
	if estimatedCompletionTimeSecondsPb != nil {
		pb.EstimatedCompletionTimeSeconds = *estimatedCompletionTimeSecondsPb
	}

	latestVersionCurrentlyProcessingPb := &st.LatestVersionCurrentlyProcessing
	if latestVersionCurrentlyProcessingPb != nil {
		pb.LatestVersionCurrentlyProcessing = *latestVersionCurrentlyProcessingPb
	}

	syncProgressCompletionPb := &st.SyncProgressCompletion
	if syncProgressCompletionPb != nil {
		pb.SyncProgressCompletion = *syncProgressCompletionPb
	}

	syncedRowCountPb := &st.SyncedRowCount
	if syncedRowCountPb != nil {
		pb.SyncedRowCount = *syncedRowCountPb
	}

	totalRowCountPb := &st.TotalRowCount
	if totalRowCountPb != nil {
		pb.TotalRowCount = *totalRowCountPb
	}

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
	estimatedCompletionTimeSecondsField := &pb.EstimatedCompletionTimeSeconds
	if estimatedCompletionTimeSecondsField != nil {
		st.EstimatedCompletionTimeSeconds = *estimatedCompletionTimeSecondsField
	}
	latestVersionCurrentlyProcessingField := &pb.LatestVersionCurrentlyProcessing
	if latestVersionCurrentlyProcessingField != nil {
		st.LatestVersionCurrentlyProcessing = *latestVersionCurrentlyProcessingField
	}
	syncProgressCompletionField := &pb.SyncProgressCompletion
	if syncProgressCompletionField != nil {
		st.SyncProgressCompletion = *syncProgressCompletionField
	}
	syncedRowCountField := &pb.SyncedRowCount
	if syncedRowCountField != nil {
		st.SyncedRowCount = *syncedRowCountField
	}
	totalRowCountField := &pb.TotalRowCount
	if totalRowCountField != nil {
		st.TotalRowCount = *totalRowCountField
	}

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
	ChildColumns []string
	// The name of the constraint.
	Name string
}

func primaryKeyConstraintToPb(st *PrimaryKeyConstraint) (*primaryKeyConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &primaryKeyConstraintPb{}

	var childColumnsPb []string
	for _, item := range st.ChildColumns {
		itemPb := &item
		if itemPb != nil {
			childColumnsPb = append(childColumnsPb, *itemPb)
		}
	}
	pb.ChildColumns = childColumnsPb

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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

	var childColumnsField []string
	for _, item := range pb.ChildColumns {
		itemField := &item
		if itemField != nil {
			childColumnsField = append(childColumnsField, *itemField)
		}
	}
	st.ChildColumns = childColumnsField
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

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
	Principal string
	// The privileges assigned to the principal.
	Privileges []Privilege

	ForceSendFields []string
}

func privilegeAssignmentToPb(st *PrivilegeAssignment) (*privilegeAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &privilegeAssignmentPb{}
	principalPb := &st.Principal
	if principalPb != nil {
		pb.Principal = *principalPb
	}

	var privilegesPb []Privilege
	for _, item := range st.Privileges {
		itemPb := &item
		if itemPb != nil {
			privilegesPb = append(privilegesPb, *itemPb)
		}
	}
	pb.Privileges = privilegesPb

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
	principalField := &pb.Principal
	if principalField != nil {
		st.Principal = *principalField
	}

	var privilegesField []Privilege
	for _, item := range pb.Privileges {
		itemField := &item
		if itemField != nil {
			privilegesField = append(privilegesField, *itemField)
		}
	}
	st.Privileges = privilegesField

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
	State ProvisioningInfoState
}

func provisioningInfoToPb(st *ProvisioningInfo) (*provisioningInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningInfoPb{}
	statePb := &st.State
	if statePb != nil {
		pb.State = *statePb
	}

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
	stateField := &pb.State
	if stateField != nil {
		st.State = *stateField
	}

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
	LastRefreshedAt int64
	// Name of the parent resource. Returns metastore ID if the parent is a
	// metastore.
	ParentFullName string
	// The quota parent securable type.
	ParentSecurableType SecurableType
	// The current usage of the resource quota.
	QuotaCount int
	// The current limit of the resource quota.
	QuotaLimit int
	// The name of the quota.
	QuotaName string

	ForceSendFields []string
}

func quotaInfoToPb(st *QuotaInfo) (*quotaInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &quotaInfoPb{}
	lastRefreshedAtPb := &st.LastRefreshedAt
	if lastRefreshedAtPb != nil {
		pb.LastRefreshedAt = *lastRefreshedAtPb
	}

	parentFullNamePb := &st.ParentFullName
	if parentFullNamePb != nil {
		pb.ParentFullName = *parentFullNamePb
	}

	parentSecurableTypePb := &st.ParentSecurableType
	if parentSecurableTypePb != nil {
		pb.ParentSecurableType = *parentSecurableTypePb
	}

	quotaCountPb := &st.QuotaCount
	if quotaCountPb != nil {
		pb.QuotaCount = *quotaCountPb
	}

	quotaLimitPb := &st.QuotaLimit
	if quotaLimitPb != nil {
		pb.QuotaLimit = *quotaLimitPb
	}

	quotaNamePb := &st.QuotaName
	if quotaNamePb != nil {
		pb.QuotaName = *quotaNamePb
	}

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
	lastRefreshedAtField := &pb.LastRefreshedAt
	if lastRefreshedAtField != nil {
		st.LastRefreshedAt = *lastRefreshedAtField
	}
	parentFullNameField := &pb.ParentFullName
	if parentFullNameField != nil {
		st.ParentFullName = *parentFullNameField
	}
	parentSecurableTypeField := &pb.ParentSecurableType
	if parentSecurableTypeField != nil {
		st.ParentSecurableType = *parentSecurableTypeField
	}
	quotaCountField := &pb.QuotaCount
	if quotaCountField != nil {
		st.QuotaCount = *quotaCountField
	}
	quotaLimitField := &pb.QuotaLimit
	if quotaLimitField != nil {
		st.QuotaLimit = *quotaLimitField
	}
	quotaNameField := &pb.QuotaName
	if quotaNameField != nil {
		st.QuotaName = *quotaNameField
	}

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
	AccessKeyId string
	// The secret access key associated with the access key.
	SecretAccessKey string
	// The generated JWT that users must pass to use the temporary credentials.
	SessionToken string

	ForceSendFields []string
}

func r2CredentialsToPb(st *R2Credentials) (*r2CredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &r2CredentialsPb{}
	accessKeyIdPb := &st.AccessKeyId
	if accessKeyIdPb != nil {
		pb.AccessKeyId = *accessKeyIdPb
	}

	secretAccessKeyPb := &st.SecretAccessKey
	if secretAccessKeyPb != nil {
		pb.SecretAccessKey = *secretAccessKeyPb
	}

	sessionTokenPb := &st.SessionToken
	if sessionTokenPb != nil {
		pb.SessionToken = *sessionTokenPb
	}

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
	accessKeyIdField := &pb.AccessKeyId
	if accessKeyIdField != nil {
		st.AccessKeyId = *accessKeyIdField
	}
	secretAccessKeyField := &pb.SecretAccessKey
	if secretAccessKeyField != nil {
		st.SecretAccessKey = *secretAccessKeyField
	}
	sessionTokenField := &pb.SessionToken
	if sessionTokenField != nil {
		st.SessionToken = *sessionTokenField
	}

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
	IncludeBrowse bool
	// The three-level (fully qualified) name of the volume
	Name string

	ForceSendFields []string
}

func readVolumeRequestToPb(st *ReadVolumeRequest) (*readVolumeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &readVolumeRequestPb{}
	includeBrowsePb := &st.IncludeBrowse
	if includeBrowsePb != nil {
		pb.IncludeBrowse = *includeBrowsePb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

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
	includeBrowseField := &pb.IncludeBrowse
	if includeBrowseField != nil {
		st.IncludeBrowse = *includeBrowseField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

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
	TableName string
	// Optional argument to specify the warehouse for dashboard regeneration. If
	// not specified, the first running warehouse will be used.
	WarehouseId string

	ForceSendFields []string
}

func regenerateDashboardRequestToPb(st *RegenerateDashboardRequest) (*regenerateDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &regenerateDashboardRequestPb{}
	tableNamePb := &st.TableName
	if tableNamePb != nil {
		pb.TableName = *tableNamePb
	}

	warehouseIdPb := &st.WarehouseId
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

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
	tableNameField := &pb.TableName
	if tableNameField != nil {
		st.TableName = *tableNameField
	}
	warehouseIdField := &pb.WarehouseId
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

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
	DashboardId string
	// The directory where the regenerated dashboard is stored.
	ParentFolder string

	ForceSendFields []string
}

func regenerateDashboardResponseToPb(st *RegenerateDashboardResponse) (*regenerateDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &regenerateDashboardResponsePb{}
	dashboardIdPb := &st.DashboardId
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	parentFolderPb := &st.ParentFolder
	if parentFolderPb != nil {
		pb.ParentFolder = *parentFolderPb
	}

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
	dashboardIdField := &pb.DashboardId
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	parentFolderField := &pb.ParentFolder
	if parentFolderField != nil {
		st.ParentFolder = *parentFolderField
	}

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
	AliasName string
	// Integer version number of the model version to which this alias points.
	VersionNum int

	ForceSendFields []string
}

func registeredModelAliasToPb(st *RegisteredModelAlias) (*registeredModelAliasPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelAliasPb{}
	aliasNamePb := &st.AliasName
	if aliasNamePb != nil {
		pb.AliasName = *aliasNamePb
	}

	versionNumPb := &st.VersionNum
	if versionNumPb != nil {
		pb.VersionNum = *versionNumPb
	}

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
	aliasNameField := &pb.AliasName
	if aliasNameField != nil {
		st.AliasName = *aliasNameField
	}
	versionNumField := &pb.VersionNum
	if versionNumField != nil {
		st.VersionNum = *versionNumField
	}

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
	Aliases []RegisteredModelAlias
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool
	// The name of the catalog where the schema and the registered model reside
	CatalogName string
	// The comment attached to the registered model
	Comment string
	// Creation timestamp of the registered model in milliseconds since the Unix
	// epoch
	CreatedAt int64
	// The identifier of the user who created the registered model
	CreatedBy string
	// The three-level (fully qualified) name of the registered model
	FullName string
	// The unique identifier of the metastore
	MetastoreId string
	// The name of the registered model
	Name string
	// The identifier of the user who owns the registered model
	Owner string
	// The name of the schema where the registered model resides
	SchemaName string
	// The storage location on the cloud under which model version data files
	// are stored
	StorageLocation string
	// Last-update timestamp of the registered model in milliseconds since the
	// Unix epoch
	UpdatedAt int64
	// The identifier of the user who updated the registered model last time
	UpdatedBy string

	ForceSendFields []string
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

	browseOnlyPb := &st.BrowseOnly
	if browseOnlyPb != nil {
		pb.BrowseOnly = *browseOnlyPb
	}

	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	createdAtPb := &st.CreatedAt
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	schemaNamePb := &st.SchemaName
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

	storageLocationPb := &st.StorageLocation
	if storageLocationPb != nil {
		pb.StorageLocation = *storageLocationPb
	}

	updatedAtPb := &st.UpdatedAt
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb := &st.UpdatedBy
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

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
	for _, item := range pb.Aliases {
		itemField, err := registeredModelAliasFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			aliasesField = append(aliasesField, *itemField)
		}
	}
	st.Aliases = aliasesField
	browseOnlyField := &pb.BrowseOnly
	if browseOnlyField != nil {
		st.BrowseOnly = *browseOnlyField
	}
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	createdAtField := &pb.CreatedAt
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}
	schemaNameField := &pb.SchemaName
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}
	storageLocationField := &pb.StorageLocation
	if storageLocationField != nil {
		st.StorageLocation = *storageLocationField
	}
	updatedAtField := &pb.UpdatedAt
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField := &pb.UpdatedBy
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}

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
	TableName string
}

func runRefreshRequestToPb(st *RunRefreshRequest) (*runRefreshRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runRefreshRequestPb{}
	tableNamePb := &st.TableName
	if tableNamePb != nil {
		pb.TableName = *tableNamePb
	}

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
	tableNameField := &pb.TableName
	if tableNameField != nil {
		st.TableName = *tableNameField
	}

	return st, nil
}

type SchemaInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool
	// Name of parent catalog.
	CatalogName string
	// The type of the parent catalog.
	CatalogType string
	// User-provided free-form text description.
	Comment string
	// Time at which this schema was created, in epoch milliseconds.
	CreatedAt int64
	// Username of schema creator.
	CreatedBy string

	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization
	// Full name of schema, in form of __catalog_name__.__schema_name__.
	FullName string
	// Unique identifier of parent metastore.
	MetastoreId string
	// Name of schema, relative to parent catalog.
	Name string
	// Username of current owner of schema.
	Owner string
	// A map of key-value properties attached to the securable.
	Properties map[string]string
	// The unique identifier of the schema.
	SchemaId string
	// Storage location for managed tables within schema.
	StorageLocation string
	// Storage root URL for managed tables within schema.
	StorageRoot string
	// Time at which this schema was created, in epoch milliseconds.
	UpdatedAt int64
	// Username of user who last modified schema.
	UpdatedBy string

	ForceSendFields []string
}

func schemaInfoToPb(st *SchemaInfo) (*schemaInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &schemaInfoPb{}
	browseOnlyPb := &st.BrowseOnly
	if browseOnlyPb != nil {
		pb.BrowseOnly = *browseOnlyPb
	}

	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	catalogTypePb := &st.CatalogType
	if catalogTypePb != nil {
		pb.CatalogType = *catalogTypePb
	}

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	createdAtPb := &st.CreatedAt
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	effectivePredictiveOptimizationFlagPb, err := effectivePredictiveOptimizationFlagToPb(st.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagPb != nil {
		pb.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagPb
	}

	enablePredictiveOptimizationPb := &st.EnablePredictiveOptimization
	if enablePredictiveOptimizationPb != nil {
		pb.EnablePredictiveOptimization = *enablePredictiveOptimizationPb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	propertiesPb := map[string]string{}
	for k, v := range st.Properties {
		itemPb := &v
		if itemPb != nil {
			propertiesPb[k] = *itemPb
		}
	}
	pb.Properties = propertiesPb

	schemaIdPb := &st.SchemaId
	if schemaIdPb != nil {
		pb.SchemaId = *schemaIdPb
	}

	storageLocationPb := &st.StorageLocation
	if storageLocationPb != nil {
		pb.StorageLocation = *storageLocationPb
	}

	storageRootPb := &st.StorageRoot
	if storageRootPb != nil {
		pb.StorageRoot = *storageRootPb
	}

	updatedAtPb := &st.UpdatedAt
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb := &st.UpdatedBy
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

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
	browseOnlyField := &pb.BrowseOnly
	if browseOnlyField != nil {
		st.BrowseOnly = *browseOnlyField
	}
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	catalogTypeField := &pb.CatalogType
	if catalogTypeField != nil {
		st.CatalogType = *catalogTypeField
	}
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	createdAtField := &pb.CreatedAt
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	effectivePredictiveOptimizationFlagField, err := effectivePredictiveOptimizationFlagFromPb(pb.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagField != nil {
		st.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagField
	}
	enablePredictiveOptimizationField := &pb.EnablePredictiveOptimization
	if enablePredictiveOptimizationField != nil {
		st.EnablePredictiveOptimization = *enablePredictiveOptimizationField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}

	propertiesField := map[string]string{}
	for k, v := range pb.Properties {
		itemField := &v
		if itemField != nil {
			propertiesField[k] = *itemField
		}
	}
	st.Properties = propertiesField
	schemaIdField := &pb.SchemaId
	if schemaIdField != nil {
		st.SchemaId = *schemaIdField
	}
	storageLocationField := &pb.StorageLocation
	if storageLocationField != nil {
		st.StorageLocation = *storageLocationField
	}
	storageRootField := &pb.StorageRoot
	if storageRootField != nil {
		st.StorageRoot = *storageRootField
	}
	updatedAtField := &pb.UpdatedAt
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField := &pb.UpdatedBy
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}

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
	ArtifactMatchers []ArtifactMatcher
	// The artifact type of the allowlist.
	ArtifactType ArtifactType
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	CreatedAt int64
	// Username of the user who set the artifact allowlist.
	CreatedBy string
	// Unique identifier of parent metastore.
	MetastoreId string

	ForceSendFields []string
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

	artifactTypePb := &st.ArtifactType
	if artifactTypePb != nil {
		pb.ArtifactType = *artifactTypePb
	}

	createdAtPb := &st.CreatedAt
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

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
	for _, item := range pb.ArtifactMatchers {
		itemField, err := artifactMatcherFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			artifactMatchersField = append(artifactMatchersField, *itemField)
		}
	}
	st.ArtifactMatchers = artifactMatchersField
	artifactTypeField := &pb.ArtifactType
	if artifactTypeField != nil {
		st.ArtifactType = *artifactTypeField
	}
	createdAtField := &pb.CreatedAt
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}

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
	Alias string
	// Full name of the registered model
	FullName string
	// The version number of the model version to which the alias points
	VersionNum int
}

func setRegisteredModelAliasRequestToPb(st *SetRegisteredModelAliasRequest) (*setRegisteredModelAliasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setRegisteredModelAliasRequestPb{}
	aliasPb := &st.Alias
	if aliasPb != nil {
		pb.Alias = *aliasPb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	versionNumPb := &st.VersionNum
	if versionNumPb != nil {
		pb.VersionNum = *versionNumPb
	}

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
	aliasField := &pb.Alias
	if aliasField != nil {
		st.Alias = *aliasField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	versionNumField := &pb.VersionNum
	if versionNumField != nil {
		st.VersionNum = *versionNumField
	}

	return st, nil
}

// Server-Side Encryption properties for clients communicating with AWS s3.
type SseEncryptionDetails struct {
	// The type of key encryption to use (affects headers from s3 client).
	Algorithm SseEncryptionDetailsAlgorithm
	// When algorithm is **AWS_SSE_KMS** this field specifies the ARN of the SSE
	// key to use.
	AwsKmsKeyArn string

	ForceSendFields []string
}

func sseEncryptionDetailsToPb(st *SseEncryptionDetails) (*sseEncryptionDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sseEncryptionDetailsPb{}
	algorithmPb := &st.Algorithm
	if algorithmPb != nil {
		pb.Algorithm = *algorithmPb
	}

	awsKmsKeyArnPb := &st.AwsKmsKeyArn
	if awsKmsKeyArnPb != nil {
		pb.AwsKmsKeyArn = *awsKmsKeyArnPb
	}

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
	algorithmField := &pb.Algorithm
	if algorithmField != nil {
		st.Algorithm = *algorithmField
	}
	awsKmsKeyArnField := &pb.AwsKmsKeyArn
	if awsKmsKeyArnField != nil {
		st.AwsKmsKeyArn = *awsKmsKeyArnField
	}

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
	AwsIamRole *AwsIamRoleResponse
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityResponse
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken
	// Comment associated with the credential.
	Comment string
	// Time at which this Credential was created, in epoch milliseconds.
	CreatedAt int64
	// Username of credential creator.
	CreatedBy string
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountResponse
	// The full name of the credential.
	FullName string
	// The unique identifier of the credential.
	Id string

	IsolationMode IsolationMode
	// Unique identifier of parent metastore.
	MetastoreId string
	// The credential name. The name must be unique within the metastore.
	Name string
	// Username of current owner of credential.
	Owner string
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool
	// Time at which this credential was last modified, in epoch milliseconds.
	UpdatedAt int64
	// Username of user who last modified the credential.
	UpdatedBy string
	// Whether this credential is the current metastore's root storage
	// credential.
	UsedForManagedStorage bool

	ForceSendFields []string
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

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	createdAtPb := &st.CreatedAt
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	databricksGcpServiceAccountPb, err := databricksGcpServiceAccountResponseToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	idPb := &st.Id
	if idPb != nil {
		pb.Id = *idPb
	}

	isolationModePb := &st.IsolationMode
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	readOnlyPb := &st.ReadOnly
	if readOnlyPb != nil {
		pb.ReadOnly = *readOnlyPb
	}

	updatedAtPb := &st.UpdatedAt
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb := &st.UpdatedBy
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

	usedForManagedStoragePb := &st.UsedForManagedStorage
	if usedForManagedStoragePb != nil {
		pb.UsedForManagedStorage = *usedForManagedStoragePb
	}

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	createdAtField := &pb.CreatedAt
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	databricksGcpServiceAccountField, err := databricksGcpServiceAccountResponseFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	idField := &pb.Id
	if idField != nil {
		st.Id = *idField
	}
	isolationModeField := &pb.IsolationMode
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}
	readOnlyField := &pb.ReadOnly
	if readOnlyField != nil {
		st.ReadOnly = *readOnlyField
	}
	updatedAtField := &pb.UpdatedAt
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField := &pb.UpdatedBy
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}
	usedForManagedStorageField := &pb.UsedForManagedStorage
	if usedForManagedStorageField != nil {
		st.UsedForManagedStorage = *usedForManagedStorageField
	}

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
	Schema string
	// The current state of enablement for the system schema. An empty string
	// means the system schema is available and ready for opt-in.
	State SystemSchemaInfoState

	ForceSendFields []string
}

func systemSchemaInfoToPb(st *SystemSchemaInfo) (*systemSchemaInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &systemSchemaInfoPb{}
	schemaPb := &st.Schema
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

	statePb := &st.State
	if statePb != nil {
		pb.State = *statePb
	}

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
	schemaField := &pb.Schema
	if schemaField != nil {
		st.Schema = *schemaField
	}
	stateField := &pb.State
	if stateField != nil {
		st.State = *stateField
	}

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
	ForeignKeyConstraint *ForeignKeyConstraint

	NamedTableConstraint *NamedTableConstraint

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
	TableFullName string
}

func tableDependencyToPb(st *TableDependency) (*tableDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableDependencyPb{}
	tableFullNamePb := &st.TableFullName
	if tableFullNamePb != nil {
		pb.TableFullName = *tableFullNamePb
	}

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
	tableFullNameField := &pb.TableFullName
	if tableFullNameField != nil {
		st.TableFullName = *tableFullNameField
	}

	return st, nil
}

type TableExistsResponse struct {
	// Whether the table exists or not.
	TableExists bool

	ForceSendFields []string
}

func tableExistsResponseToPb(st *TableExistsResponse) (*tableExistsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableExistsResponsePb{}
	tableExistsPb := &st.TableExists
	if tableExistsPb != nil {
		pb.TableExists = *tableExistsPb
	}

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
	tableExistsField := &pb.TableExists
	if tableExistsField != nil {
		st.TableExists = *tableExistsField
	}

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
	AccessPoint string
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool
	// Name of parent catalog.
	CatalogName string
	// The array of __ColumnInfo__ definitions of the table's columns.
	Columns []ColumnInfo
	// User-provided free-form text description.
	Comment string
	// Time at which this table was created, in epoch milliseconds.
	CreatedAt int64
	// Username of table creator.
	CreatedBy string
	// Unique ID of the Data Access Configuration to use with the table data.
	DataAccessConfigurationId string
	// Data source format
	DataSourceFormat DataSourceFormat
	// Time at which this table was deleted, in epoch milliseconds. Field is
	// omitted if table is not deleted.
	DeletedAt int64
	// Information pertaining to current state of the delta table.
	DeltaRuntimePropertiesKvpairs *DeltaRuntimePropertiesKvPairs

	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails
	// Full name of table, in form of
	// __catalog_name__.__schema_name__.__table_name__
	FullName string
	// Unique identifier of parent metastore.
	MetastoreId string
	// Name of table, relative to parent schema.
	Name string
	// Username of current owner of table.
	Owner string
	// The pipeline ID of the table. Applicable for tables created by pipelines
	// (Materialized View, Streaming Table, etc.).
	PipelineId string
	// A map of key-value properties attached to the securable.
	Properties map[string]string

	RowFilter *TableRowFilter
	// Name of parent schema relative to its parent catalog.
	SchemaName string
	// List of schemes whose objects can be referenced without qualification.
	SqlPath string
	// Name of the storage credential, when a storage credential is configured
	// for use with this table.
	StorageCredentialName string
	// Storage root URL for table (for **MANAGED**, **EXTERNAL** tables)
	StorageLocation string
	// List of table constraints. Note: this field is not set in the output of
	// the __listTables__ API.
	TableConstraints []TableConstraint
	// The unique identifier of the table.
	TableId string

	TableType TableType
	// Time at which this table was last modified, in epoch milliseconds.
	UpdatedAt int64
	// Username of user who last modified the table.
	UpdatedBy string
	// View definition SQL (when __table_type__ is **VIEW**,
	// **MATERIALIZED_VIEW**, or **STREAMING_TABLE**)
	ViewDefinition string
	// View dependencies (when table_type == **VIEW** or **MATERIALIZED_VIEW**,
	// **STREAMING_TABLE**) - when DependencyList is None, the dependency is not
	// provided; - when DependencyList is an empty list, the dependency is
	// provided but is empty; - when DependencyList is not an empty list,
	// dependencies are provided and recorded.
	ViewDependencies *DependencyList

	ForceSendFields []string
}

func tableInfoToPb(st *TableInfo) (*tableInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableInfoPb{}
	accessPointPb := &st.AccessPoint
	if accessPointPb != nil {
		pb.AccessPoint = *accessPointPb
	}

	browseOnlyPb := &st.BrowseOnly
	if browseOnlyPb != nil {
		pb.BrowseOnly = *browseOnlyPb
	}

	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

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

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	createdAtPb := &st.CreatedAt
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	dataAccessConfigurationIdPb := &st.DataAccessConfigurationId
	if dataAccessConfigurationIdPb != nil {
		pb.DataAccessConfigurationId = *dataAccessConfigurationIdPb
	}

	dataSourceFormatPb := &st.DataSourceFormat
	if dataSourceFormatPb != nil {
		pb.DataSourceFormat = *dataSourceFormatPb
	}

	deletedAtPb := &st.DeletedAt
	if deletedAtPb != nil {
		pb.DeletedAt = *deletedAtPb
	}

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

	enablePredictiveOptimizationPb := &st.EnablePredictiveOptimization
	if enablePredictiveOptimizationPb != nil {
		pb.EnablePredictiveOptimization = *enablePredictiveOptimizationPb
	}

	encryptionDetailsPb, err := encryptionDetailsToPb(st.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsPb != nil {
		pb.EncryptionDetails = encryptionDetailsPb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	pipelineIdPb := &st.PipelineId
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

	propertiesPb := map[string]string{}
	for k, v := range st.Properties {
		itemPb := &v
		if itemPb != nil {
			propertiesPb[k] = *itemPb
		}
	}
	pb.Properties = propertiesPb

	rowFilterPb, err := tableRowFilterToPb(st.RowFilter)
	if err != nil {
		return nil, err
	}
	if rowFilterPb != nil {
		pb.RowFilter = rowFilterPb
	}

	schemaNamePb := &st.SchemaName
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

	sqlPathPb := &st.SqlPath
	if sqlPathPb != nil {
		pb.SqlPath = *sqlPathPb
	}

	storageCredentialNamePb := &st.StorageCredentialName
	if storageCredentialNamePb != nil {
		pb.StorageCredentialName = *storageCredentialNamePb
	}

	storageLocationPb := &st.StorageLocation
	if storageLocationPb != nil {
		pb.StorageLocation = *storageLocationPb
	}

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

	tableIdPb := &st.TableId
	if tableIdPb != nil {
		pb.TableId = *tableIdPb
	}

	tableTypePb := &st.TableType
	if tableTypePb != nil {
		pb.TableType = *tableTypePb
	}

	updatedAtPb := &st.UpdatedAt
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb := &st.UpdatedBy
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

	viewDefinitionPb := &st.ViewDefinition
	if viewDefinitionPb != nil {
		pb.ViewDefinition = *viewDefinitionPb
	}

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
	accessPointField := &pb.AccessPoint
	if accessPointField != nil {
		st.AccessPoint = *accessPointField
	}
	browseOnlyField := &pb.BrowseOnly
	if browseOnlyField != nil {
		st.BrowseOnly = *browseOnlyField
	}
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}

	var columnsField []ColumnInfo
	for _, item := range pb.Columns {
		itemField, err := ColumnInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			columnsField = append(columnsField, *itemField)
		}
	}
	st.Columns = columnsField
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	createdAtField := &pb.CreatedAt
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	dataAccessConfigurationIdField := &pb.DataAccessConfigurationId
	if dataAccessConfigurationIdField != nil {
		st.DataAccessConfigurationId = *dataAccessConfigurationIdField
	}
	dataSourceFormatField := &pb.DataSourceFormat
	if dataSourceFormatField != nil {
		st.DataSourceFormat = *dataSourceFormatField
	}
	deletedAtField := &pb.DeletedAt
	if deletedAtField != nil {
		st.DeletedAt = *deletedAtField
	}
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
	enablePredictiveOptimizationField := &pb.EnablePredictiveOptimization
	if enablePredictiveOptimizationField != nil {
		st.EnablePredictiveOptimization = *enablePredictiveOptimizationField
	}
	encryptionDetailsField, err := encryptionDetailsFromPb(pb.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsField != nil {
		st.EncryptionDetails = encryptionDetailsField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}
	pipelineIdField := &pb.PipelineId
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}

	propertiesField := map[string]string{}
	for k, v := range pb.Properties {
		itemField := &v
		if itemField != nil {
			propertiesField[k] = *itemField
		}
	}
	st.Properties = propertiesField
	rowFilterField, err := tableRowFilterFromPb(pb.RowFilter)
	if err != nil {
		return nil, err
	}
	if rowFilterField != nil {
		st.RowFilter = rowFilterField
	}
	schemaNameField := &pb.SchemaName
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}
	sqlPathField := &pb.SqlPath
	if sqlPathField != nil {
		st.SqlPath = *sqlPathField
	}
	storageCredentialNameField := &pb.StorageCredentialName
	if storageCredentialNameField != nil {
		st.StorageCredentialName = *storageCredentialNameField
	}
	storageLocationField := &pb.StorageLocation
	if storageLocationField != nil {
		st.StorageLocation = *storageLocationField
	}

	var tableConstraintsField []TableConstraint
	for _, item := range pb.TableConstraints {
		itemField, err := tableConstraintFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			tableConstraintsField = append(tableConstraintsField, *itemField)
		}
	}
	st.TableConstraints = tableConstraintsField
	tableIdField := &pb.TableId
	if tableIdField != nil {
		st.TableId = *tableIdField
	}
	tableTypeField := &pb.TableType
	if tableTypeField != nil {
		st.TableType = *tableTypeField
	}
	updatedAtField := &pb.UpdatedAt
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField := &pb.UpdatedBy
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}
	viewDefinitionField := &pb.ViewDefinition
	if viewDefinitionField != nil {
		st.ViewDefinition = *viewDefinitionField
	}
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
	FunctionName string
	// The list of table columns to be passed as input to the row filter
	// function. The column types should match the types of the filter function
	// arguments.
	InputColumnNames []string
}

func tableRowFilterToPb(st *TableRowFilter) (*tableRowFilterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableRowFilterPb{}
	functionNamePb := &st.FunctionName
	if functionNamePb != nil {
		pb.FunctionName = *functionNamePb
	}

	var inputColumnNamesPb []string
	for _, item := range st.InputColumnNames {
		itemPb := &item
		if itemPb != nil {
			inputColumnNamesPb = append(inputColumnNamesPb, *itemPb)
		}
	}
	pb.InputColumnNames = inputColumnNamesPb

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
	functionNameField := &pb.FunctionName
	if functionNameField != nil {
		st.FunctionName = *functionNameField
	}

	var inputColumnNamesField []string
	for _, item := range pb.InputColumnNames {
		itemField := &item
		if itemField != nil {
			inputColumnNamesField = append(inputColumnNamesField, *itemField)
		}
	}
	st.InputColumnNames = inputColumnNamesField

	return st, nil
}

type TableSummary struct {
	// The full name of the table.
	FullName string

	TableType TableType

	ForceSendFields []string
}

func tableSummaryToPb(st *TableSummary) (*tableSummaryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableSummaryPb{}
	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	tableTypePb := &st.TableType
	if tableTypePb != nil {
		pb.TableType = *tableTypePb
	}

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
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	tableTypeField := &pb.TableType
	if tableTypeField != nil {
		st.TableType = *tableTypeField
	}

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
	Key string
	// value of the tag associated with the key, could be optional
	Value string

	ForceSendFields []string
}

func TagKeyValueToPb(st *TagKeyValue) (*TagKeyValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &TagKeyValuePb{}
	keyPb := &st.Key
	if keyPb != nil {
		pb.Key = *keyPb
	}

	valuePb := &st.Value
	if valuePb != nil {
		pb.Value = *valuePb
	}

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
	keyField := &pb.Key
	if keyField != nil {
		st.Key = *keyField
	}
	valueField := &pb.Value
	if valueField != nil {
		st.Value = *valueField
	}

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
	AwsTempCredentials *AwsCredentials
	// Azure Active Directory token, essentially the Oauth token for Azure
	// Service Principal or Managed Identity. Read more at
	// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
	AzureAad *AzureActiveDirectoryToken
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	ExpirationTime int64
	// GCP temporary credentials for API authentication. Read more at
	// https://developers.google.com/identity/protocols/oauth2/service-account
	GcpOauthToken *GcpOauthToken

	ForceSendFields []string
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

	expirationTimePb := &st.ExpirationTime
	if expirationTimePb != nil {
		pb.ExpirationTime = *expirationTimePb
	}

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
	expirationTimeField := &pb.ExpirationTime
	if expirationTimeField != nil {
		st.ExpirationTime = *expirationTimeField
	}
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
	LastProcessedCommitVersion int64
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	Timestamp string
	// Progress of the active data synchronization pipeline.
	TriggeredUpdateProgress *PipelineProgress

	ForceSendFields []string
}

func triggeredUpdateStatusToPb(st *TriggeredUpdateStatus) (*triggeredUpdateStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &triggeredUpdateStatusPb{}
	lastProcessedCommitVersionPb := &st.LastProcessedCommitVersion
	if lastProcessedCommitVersionPb != nil {
		pb.LastProcessedCommitVersion = *lastProcessedCommitVersionPb
	}

	timestampPb := &st.Timestamp
	if timestampPb != nil {
		pb.Timestamp = *timestampPb
	}

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
	lastProcessedCommitVersionField := &pb.LastProcessedCommitVersion
	if lastProcessedCommitVersionField != nil {
		st.LastProcessedCommitVersion = *lastProcessedCommitVersionField
	}
	timestampField := &pb.Timestamp
	if timestampField != nil {
		st.Timestamp = *timestampField
	}
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
	MetastoreId string
	// A workspace ID.
	WorkspaceId int64
}

func unassignRequestToPb(st *UnassignRequest) (*unassignRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &unassignRequestPb{}
	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	workspaceIdPb := &st.WorkspaceId
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

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
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	workspaceIdField := &pb.WorkspaceId
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}

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
	Comment string
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode CatalogIsolationMode
	// The name of the catalog.
	Name string
	// New name for the catalog.
	NewName string
	// A map of key-value properties attached to the securable.
	Options map[string]string
	// Username of current owner of catalog.
	Owner string
	// A map of key-value properties attached to the securable.
	Properties map[string]string

	ForceSendFields []string
}

func updateCatalogToPb(st *UpdateCatalog) (*updateCatalogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCatalogPb{}
	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	enablePredictiveOptimizationPb := &st.EnablePredictiveOptimization
	if enablePredictiveOptimizationPb != nil {
		pb.EnablePredictiveOptimization = *enablePredictiveOptimizationPb
	}

	isolationModePb := &st.IsolationMode
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	newNamePb := &st.NewName
	if newNamePb != nil {
		pb.NewName = *newNamePb
	}

	optionsPb := map[string]string{}
	for k, v := range st.Options {
		itemPb := &v
		if itemPb != nil {
			optionsPb[k] = *itemPb
		}
	}
	pb.Options = optionsPb

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	propertiesPb := map[string]string{}
	for k, v := range st.Properties {
		itemPb := &v
		if itemPb != nil {
			propertiesPb[k] = *itemPb
		}
	}
	pb.Properties = propertiesPb

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	enablePredictiveOptimizationField := &pb.EnablePredictiveOptimization
	if enablePredictiveOptimizationField != nil {
		st.EnablePredictiveOptimization = *enablePredictiveOptimizationField
	}
	isolationModeField := &pb.IsolationMode
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	newNameField := &pb.NewName
	if newNameField != nil {
		st.NewName = *newNameField
	}

	optionsField := map[string]string{}
	for k, v := range pb.Options {
		itemField := &v
		if itemField != nil {
			optionsField[k] = *itemField
		}
	}
	st.Options = optionsField
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}

	propertiesField := map[string]string{}
	for k, v := range pb.Properties {
		itemField := &v
		if itemField != nil {
			propertiesField[k] = *itemField
		}
	}
	st.Properties = propertiesField

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
	Name string
	// New name for the connection.
	NewName string
	// A map of key-value properties attached to the securable.
	Options map[string]string
	// Username of current owner of the connection.
	Owner string

	ForceSendFields []string
}

func updateConnectionToPb(st *UpdateConnection) (*updateConnectionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateConnectionPb{}
	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	newNamePb := &st.NewName
	if newNamePb != nil {
		pb.NewName = *newNamePb
	}

	optionsPb := map[string]string{}
	for k, v := range st.Options {
		itemPb := &v
		if itemPb != nil {
			optionsPb[k] = *itemPb
		}
	}
	pb.Options = optionsPb

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

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
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	newNameField := &pb.NewName
	if newNameField != nil {
		st.NewName = *newNameField
	}

	optionsField := map[string]string{}
	for k, v := range pb.Options {
		itemField := &v
		if itemField != nil {
			optionsField[k] = *itemField
		}
	}
	st.Options = optionsField
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}

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
	AwsIamRole *AwsIamRole
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentity
	// The Azure service principal configuration. Only applicable when purpose
	// is **STORAGE**.
	AzureServicePrincipal *AzureServicePrincipal
	// Comment associated with the credential.
	Comment string
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	Force bool
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	IsolationMode IsolationMode
	// Name of the credential.
	NameArg string
	// New name of credential.
	NewName string
	// Username of current owner of credential.
	Owner string
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	ReadOnly bool
	// Supply true to this argument to skip validation of the updated
	// credential.
	SkipValidation bool

	ForceSendFields []string
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

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	databricksGcpServiceAccountPb, err := databricksGcpServiceAccountToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}

	forcePb := &st.Force
	if forcePb != nil {
		pb.Force = *forcePb
	}

	isolationModePb := &st.IsolationMode
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}

	nameArgPb := &st.NameArg
	if nameArgPb != nil {
		pb.NameArg = *nameArgPb
	}

	newNamePb := &st.NewName
	if newNamePb != nil {
		pb.NewName = *newNamePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	readOnlyPb := &st.ReadOnly
	if readOnlyPb != nil {
		pb.ReadOnly = *readOnlyPb
	}

	skipValidationPb := &st.SkipValidation
	if skipValidationPb != nil {
		pb.SkipValidation = *skipValidationPb
	}

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	databricksGcpServiceAccountField, err := databricksGcpServiceAccountFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	forceField := &pb.Force
	if forceField != nil {
		st.Force = *forceField
	}
	isolationModeField := &pb.IsolationMode
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
	nameArgField := &pb.NameArg
	if nameArgField != nil {
		st.NameArg = *nameArgField
	}
	newNameField := &pb.NewName
	if newNameField != nil {
		st.NewName = *newNameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}
	readOnlyField := &pb.ReadOnly
	if readOnlyField != nil {
		st.ReadOnly = *readOnlyField
	}
	skipValidationField := &pb.SkipValidation
	if skipValidationField != nil {
		st.SkipValidation = *skipValidationField
	}

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
	AccessPoint string
	// User-provided free-form text description.
	Comment string
	// Name of the storage credential used with this location.
	CredentialName string
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	Fallback bool
	// Force update even if changing url invalidates dependent external tables
	// or mounts.
	Force bool

	IsolationMode IsolationMode
	// Name of the external location.
	Name string
	// New name for the external location.
	NewName string
	// The owner of the external location.
	Owner string
	// Indicates whether the external location is read-only.
	ReadOnly bool
	// Skips validation of the storage credential associated with the external
	// location.
	SkipValidation bool
	// Path URL of the external location.
	Url string

	ForceSendFields []string
}

func updateExternalLocationToPb(st *UpdateExternalLocation) (*updateExternalLocationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateExternalLocationPb{}
	accessPointPb := &st.AccessPoint
	if accessPointPb != nil {
		pb.AccessPoint = *accessPointPb
	}

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	credentialNamePb := &st.CredentialName
	if credentialNamePb != nil {
		pb.CredentialName = *credentialNamePb
	}

	encryptionDetailsPb, err := encryptionDetailsToPb(st.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsPb != nil {
		pb.EncryptionDetails = encryptionDetailsPb
	}

	fallbackPb := &st.Fallback
	if fallbackPb != nil {
		pb.Fallback = *fallbackPb
	}

	forcePb := &st.Force
	if forcePb != nil {
		pb.Force = *forcePb
	}

	isolationModePb := &st.IsolationMode
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	newNamePb := &st.NewName
	if newNamePb != nil {
		pb.NewName = *newNamePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	readOnlyPb := &st.ReadOnly
	if readOnlyPb != nil {
		pb.ReadOnly = *readOnlyPb
	}

	skipValidationPb := &st.SkipValidation
	if skipValidationPb != nil {
		pb.SkipValidation = *skipValidationPb
	}

	urlPb := &st.Url
	if urlPb != nil {
		pb.Url = *urlPb
	}

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
	accessPointField := &pb.AccessPoint
	if accessPointField != nil {
		st.AccessPoint = *accessPointField
	}
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	credentialNameField := &pb.CredentialName
	if credentialNameField != nil {
		st.CredentialName = *credentialNameField
	}
	encryptionDetailsField, err := encryptionDetailsFromPb(pb.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsField != nil {
		st.EncryptionDetails = encryptionDetailsField
	}
	fallbackField := &pb.Fallback
	if fallbackField != nil {
		st.Fallback = *fallbackField
	}
	forceField := &pb.Force
	if forceField != nil {
		st.Force = *forceField
	}
	isolationModeField := &pb.IsolationMode
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	newNameField := &pb.NewName
	if newNameField != nil {
		st.NewName = *newNameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}
	readOnlyField := &pb.ReadOnly
	if readOnlyField != nil {
		st.ReadOnly = *readOnlyField
	}
	skipValidationField := &pb.SkipValidation
	if skipValidationField != nil {
		st.SkipValidation = *skipValidationField
	}
	urlField := &pb.Url
	if urlField != nil {
		st.Url = *urlField
	}

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
	Name string
	// Username of current owner of function.
	Owner string

	ForceSendFields []string
}

func updateFunctionToPb(st *UpdateFunction) (*updateFunctionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateFunctionPb{}
	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

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
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}

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
	DeltaSharingOrganizationName string
	// The lifetime of delta sharing recipient token in seconds.
	DeltaSharingRecipientTokenLifetimeInSeconds int64
	// The scope of Delta Sharing enabled for the metastore.
	DeltaSharingScope UpdateMetastoreDeltaSharingScope
	// Unique ID of the metastore.
	Id string
	// New name for the metastore.
	NewName string
	// The owner of the metastore.
	Owner string
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	PrivilegeModelVersion string
	// UUID of storage credential to access the metastore storage_root.
	StorageRootCredentialId string

	ForceSendFields []string
}

func updateMetastoreToPb(st *UpdateMetastore) (*updateMetastorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateMetastorePb{}
	deltaSharingOrganizationNamePb := &st.DeltaSharingOrganizationName
	if deltaSharingOrganizationNamePb != nil {
		pb.DeltaSharingOrganizationName = *deltaSharingOrganizationNamePb
	}

	deltaSharingRecipientTokenLifetimeInSecondsPb := &st.DeltaSharingRecipientTokenLifetimeInSeconds
	if deltaSharingRecipientTokenLifetimeInSecondsPb != nil {
		pb.DeltaSharingRecipientTokenLifetimeInSeconds = *deltaSharingRecipientTokenLifetimeInSecondsPb
	}

	deltaSharingScopePb := &st.DeltaSharingScope
	if deltaSharingScopePb != nil {
		pb.DeltaSharingScope = *deltaSharingScopePb
	}

	idPb := &st.Id
	if idPb != nil {
		pb.Id = *idPb
	}

	newNamePb := &st.NewName
	if newNamePb != nil {
		pb.NewName = *newNamePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	privilegeModelVersionPb := &st.PrivilegeModelVersion
	if privilegeModelVersionPb != nil {
		pb.PrivilegeModelVersion = *privilegeModelVersionPb
	}

	storageRootCredentialIdPb := &st.StorageRootCredentialId
	if storageRootCredentialIdPb != nil {
		pb.StorageRootCredentialId = *storageRootCredentialIdPb
	}

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
	deltaSharingOrganizationNameField := &pb.DeltaSharingOrganizationName
	if deltaSharingOrganizationNameField != nil {
		st.DeltaSharingOrganizationName = *deltaSharingOrganizationNameField
	}
	deltaSharingRecipientTokenLifetimeInSecondsField := &pb.DeltaSharingRecipientTokenLifetimeInSeconds
	if deltaSharingRecipientTokenLifetimeInSecondsField != nil {
		st.DeltaSharingRecipientTokenLifetimeInSeconds = *deltaSharingRecipientTokenLifetimeInSecondsField
	}
	deltaSharingScopeField := &pb.DeltaSharingScope
	if deltaSharingScopeField != nil {
		st.DeltaSharingScope = *deltaSharingScopeField
	}
	idField := &pb.Id
	if idField != nil {
		st.Id = *idField
	}
	newNameField := &pb.NewName
	if newNameField != nil {
		st.NewName = *newNameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}
	privilegeModelVersionField := &pb.PrivilegeModelVersion
	if privilegeModelVersionField != nil {
		st.PrivilegeModelVersion = *privilegeModelVersionField
	}
	storageRootCredentialIdField := &pb.StorageRootCredentialId
	if storageRootCredentialIdField != nil {
		st.StorageRootCredentialId = *storageRootCredentialIdField
	}

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
	DefaultCatalogName string
	// The unique ID of the metastore.
	MetastoreId string
	// A workspace ID.
	WorkspaceId int64

	ForceSendFields []string
}

func updateMetastoreAssignmentToPb(st *UpdateMetastoreAssignment) (*updateMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateMetastoreAssignmentPb{}
	defaultCatalogNamePb := &st.DefaultCatalogName
	if defaultCatalogNamePb != nil {
		pb.DefaultCatalogName = *defaultCatalogNamePb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	workspaceIdPb := &st.WorkspaceId
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

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
	defaultCatalogNameField := &pb.DefaultCatalogName
	if defaultCatalogNameField != nil {
		st.DefaultCatalogName = *defaultCatalogNameField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	workspaceIdField := &pb.WorkspaceId
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}

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
	Comment string
	// The three-level (fully qualified) name of the model version
	FullName string
	// The integer version number of the model version
	Version int

	ForceSendFields []string
}

func updateModelVersionRequestToPb(st *UpdateModelVersionRequest) (*updateModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateModelVersionRequestPb{}
	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	versionPb := &st.Version
	if versionPb != nil {
		pb.Version = *versionPb
	}

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	versionField := &pb.Version
	if versionField != nil {
		st.Version = *versionField
	}

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
	BaselineTableName string
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	CustomMetrics []MonitorMetric
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId string
	// The data classification config for the monitor.
	DataClassificationConfig *MonitorDataClassificationConfig
	// Configuration for monitoring inference logs.
	InferenceLog *MonitorInferenceLog
	// The notification settings for the monitor.
	Notifications *MonitorNotifications
	// Schema where output metric tables are created.
	OutputSchemaName string
	// The schedule for automatically updating and refreshing metric tables.
	Schedule *MonitorCronSchedule
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	SlicingExprs []string
	// Configuration for monitoring snapshot tables.
	Snapshot *MonitorSnapshot
	// Full name of the table.
	TableName string
	// Configuration for monitoring time series tables.
	TimeSeries *MonitorTimeSeries

	ForceSendFields []string
}

func updateMonitorToPb(st *UpdateMonitor) (*updateMonitorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateMonitorPb{}
	baselineTableNamePb := &st.BaselineTableName
	if baselineTableNamePb != nil {
		pb.BaselineTableName = *baselineTableNamePb
	}

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

	dashboardIdPb := &st.DashboardId
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

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

	outputSchemaNamePb := &st.OutputSchemaName
	if outputSchemaNamePb != nil {
		pb.OutputSchemaName = *outputSchemaNamePb
	}

	schedulePb, err := monitorCronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}

	var slicingExprsPb []string
	for _, item := range st.SlicingExprs {
		itemPb := &item
		if itemPb != nil {
			slicingExprsPb = append(slicingExprsPb, *itemPb)
		}
	}
	pb.SlicingExprs = slicingExprsPb

	snapshotPb, err := monitorSnapshotToPb(st.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotPb != nil {
		pb.Snapshot = snapshotPb
	}

	tableNamePb := &st.TableName
	if tableNamePb != nil {
		pb.TableName = *tableNamePb
	}

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
	baselineTableNameField := &pb.BaselineTableName
	if baselineTableNameField != nil {
		st.BaselineTableName = *baselineTableNameField
	}

	var customMetricsField []MonitorMetric
	for _, item := range pb.CustomMetrics {
		itemField, err := monitorMetricFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			customMetricsField = append(customMetricsField, *itemField)
		}
	}
	st.CustomMetrics = customMetricsField
	dashboardIdField := &pb.DashboardId
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
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
	outputSchemaNameField := &pb.OutputSchemaName
	if outputSchemaNameField != nil {
		st.OutputSchemaName = *outputSchemaNameField
	}
	scheduleField, err := monitorCronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}

	var slicingExprsField []string
	for _, item := range pb.SlicingExprs {
		itemField := &item
		if itemField != nil {
			slicingExprsField = append(slicingExprsField, *itemField)
		}
	}
	st.SlicingExprs = slicingExprsField
	snapshotField, err := monitorSnapshotFromPb(pb.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotField != nil {
		st.Snapshot = snapshotField
	}
	tableNameField := &pb.TableName
	if tableNameField != nil {
		st.TableName = *tableNameField
	}
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
	Changes []PermissionsChange
	// Full name of securable.
	FullName string
	// Type of securable.
	SecurableType SecurableType
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

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	securableTypePb := &st.SecurableType
	if securableTypePb != nil {
		pb.SecurableType = *securableTypePb
	}

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
	for _, item := range pb.Changes {
		itemField, err := permissionsChangeFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			changesField = append(changesField, *itemField)
		}
	}
	st.Changes = changesField
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	securableTypeField := &pb.SecurableType
	if securableTypeField != nil {
		st.SecurableType = *securableTypeField
	}

	return st, nil
}

type UpdateRegisteredModelRequest struct {
	// The comment attached to the registered model
	Comment string
	// The three-level (fully qualified) name of the registered model
	FullName string
	// New name for the registered model.
	NewName string
	// The identifier of the user who owns the registered model
	Owner string

	ForceSendFields []string
}

func updateRegisteredModelRequestToPb(st *UpdateRegisteredModelRequest) (*updateRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRegisteredModelRequestPb{}
	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	newNamePb := &st.NewName
	if newNamePb != nil {
		pb.NewName = *newNamePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	newNameField := &pb.NewName
	if newNameField != nil {
		st.NewName = *newNameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}

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
	Comment string
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	EnablePredictiveOptimization EnablePredictiveOptimization
	// Full name of the schema.
	FullName string
	// New name for the schema.
	NewName string
	// Username of current owner of schema.
	Owner string
	// A map of key-value properties attached to the securable.
	Properties map[string]string

	ForceSendFields []string
}

func updateSchemaToPb(st *UpdateSchema) (*updateSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateSchemaPb{}
	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	enablePredictiveOptimizationPb := &st.EnablePredictiveOptimization
	if enablePredictiveOptimizationPb != nil {
		pb.EnablePredictiveOptimization = *enablePredictiveOptimizationPb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	newNamePb := &st.NewName
	if newNamePb != nil {
		pb.NewName = *newNamePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	propertiesPb := map[string]string{}
	for k, v := range st.Properties {
		itemPb := &v
		if itemPb != nil {
			propertiesPb[k] = *itemPb
		}
	}
	pb.Properties = propertiesPb

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	enablePredictiveOptimizationField := &pb.EnablePredictiveOptimization
	if enablePredictiveOptimizationField != nil {
		st.EnablePredictiveOptimization = *enablePredictiveOptimizationField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	newNameField := &pb.NewName
	if newNameField != nil {
		st.NewName = *newNameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}

	propertiesField := map[string]string{}
	for k, v := range pb.Properties {
		itemField := &v
		if itemField != nil {
			propertiesField[k] = *itemField
		}
	}
	st.Properties = propertiesField

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
	AwsIamRole *AwsIamRoleRequest
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityResponse
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken
	// Comment associated with the credential.
	Comment string
	// The Databricks managed GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest
	// Force update even if there are dependent external locations or external
	// tables.
	Force bool

	IsolationMode IsolationMode
	// Name of the storage credential.
	Name string
	// New name for the storage credential.
	NewName string
	// Username of current owner of credential.
	Owner string
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool
	// Supplying true to this argument skips validation of the updated
	// credential.
	SkipValidation bool

	ForceSendFields []string
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

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	databricksGcpServiceAccountPb, err := databricksGcpServiceAccountRequestToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}

	forcePb := &st.Force
	if forcePb != nil {
		pb.Force = *forcePb
	}

	isolationModePb := &st.IsolationMode
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	newNamePb := &st.NewName
	if newNamePb != nil {
		pb.NewName = *newNamePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	readOnlyPb := &st.ReadOnly
	if readOnlyPb != nil {
		pb.ReadOnly = *readOnlyPb
	}

	skipValidationPb := &st.SkipValidation
	if skipValidationPb != nil {
		pb.SkipValidation = *skipValidationPb
	}

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	databricksGcpServiceAccountField, err := databricksGcpServiceAccountRequestFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	forceField := &pb.Force
	if forceField != nil {
		st.Force = *forceField
	}
	isolationModeField := &pb.IsolationMode
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	newNameField := &pb.NewName
	if newNameField != nil {
		st.NewName = *newNameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}
	readOnlyField := &pb.ReadOnly
	if readOnlyField != nil {
		st.ReadOnly = *readOnlyField
	}
	skipValidationField := &pb.SkipValidation
	if skipValidationField != nil {
		st.SkipValidation = *skipValidationField
	}

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
	FullName string

	Owner string

	ForceSendFields []string
}

func updateTableRequestToPb(st *UpdateTableRequest) (*updateTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateTableRequestPb{}
	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

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
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}

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
	Comment string
	// The three-level (fully qualified) name of the volume
	Name string
	// New name for the volume.
	NewName string
	// The identifier of the user who owns the volume
	Owner string

	ForceSendFields []string
}

func updateVolumeRequestContentToPb(st *UpdateVolumeRequestContent) (*updateVolumeRequestContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateVolumeRequestContentPb{}
	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	newNamePb := &st.NewName
	if newNamePb != nil {
		pb.NewName = *newNamePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

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
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	newNameField := &pb.NewName
	if newNameField != nil {
		st.NewName = *newNameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}

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
	AssignWorkspaces []int64
	// The name of the catalog.
	Name string
	// A list of workspace IDs.
	UnassignWorkspaces []int64
}

func updateWorkspaceBindingsToPb(st *UpdateWorkspaceBindings) (*updateWorkspaceBindingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateWorkspaceBindingsPb{}

	var assignWorkspacesPb []int64
	for _, item := range st.AssignWorkspaces {
		itemPb := &item
		if itemPb != nil {
			assignWorkspacesPb = append(assignWorkspacesPb, *itemPb)
		}
	}
	pb.AssignWorkspaces = assignWorkspacesPb

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	var unassignWorkspacesPb []int64
	for _, item := range st.UnassignWorkspaces {
		itemPb := &item
		if itemPb != nil {
			unassignWorkspacesPb = append(unassignWorkspacesPb, *itemPb)
		}
	}
	pb.UnassignWorkspaces = unassignWorkspacesPb

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

	var assignWorkspacesField []int64
	for _, item := range pb.AssignWorkspaces {
		itemField := &item
		if itemField != nil {
			assignWorkspacesField = append(assignWorkspacesField, *itemField)
		}
	}
	st.AssignWorkspaces = assignWorkspacesField
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	var unassignWorkspacesField []int64
	for _, item := range pb.UnassignWorkspaces {
		itemField := &item
		if itemField != nil {
			unassignWorkspacesField = append(unassignWorkspacesField, *itemField)
		}
	}
	st.UnassignWorkspaces = unassignWorkspacesField

	return st, nil
}

type UpdateWorkspaceBindingsParameters struct {
	// List of workspace bindings
	Add []WorkspaceBinding
	// List of workspace bindings
	Remove []WorkspaceBinding
	// The name of the securable.
	SecurableName string
	// The type of the securable to bind to a workspace.
	SecurableType UpdateBindingsSecurableType
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

	securableNamePb := &st.SecurableName
	if securableNamePb != nil {
		pb.SecurableName = *securableNamePb
	}

	securableTypePb := &st.SecurableType
	if securableTypePb != nil {
		pb.SecurableType = *securableTypePb
	}

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
	for _, item := range pb.Add {
		itemField, err := workspaceBindingFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			addField = append(addField, *itemField)
		}
	}
	st.Add = addField

	var removeField []WorkspaceBinding
	for _, item := range pb.Remove {
		itemField, err := workspaceBindingFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			removeField = append(removeField, *itemField)
		}
	}
	st.Remove = removeField
	securableNameField := &pb.SecurableName
	if securableNameField != nil {
		st.SecurableName = *securableNameField
	}
	securableTypeField := &pb.SecurableType
	if securableTypeField != nil {
		st.SecurableType = *securableTypeField
	}

	return st, nil
}

// Next ID: 17
type ValidateCredentialRequest struct {
	// The AWS IAM role configuration
	AwsIamRole *AwsIamRole
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentity
	// Required. The name of an existing credential or long-lived cloud
	// credential to validate.
	CredentialName string
	// GCP long-lived credential. Databricks-created Google Cloud Storage
	// service account.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount
	// The name of an existing external location to validate. Only applicable
	// for storage credentials (purpose is **STORAGE**.)
	ExternalLocationName string
	// The purpose of the credential. This should only be used when the
	// credential is specified.
	Purpose CredentialPurpose
	// Whether the credential is only usable for read operations. Only
	// applicable for storage credentials (purpose is **STORAGE**.)
	ReadOnly bool
	// The external location url to validate. Only applicable when purpose is
	// **STORAGE**.
	Url string

	ForceSendFields []string
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

	credentialNamePb := &st.CredentialName
	if credentialNamePb != nil {
		pb.CredentialName = *credentialNamePb
	}

	databricksGcpServiceAccountPb, err := databricksGcpServiceAccountToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}

	externalLocationNamePb := &st.ExternalLocationName
	if externalLocationNamePb != nil {
		pb.ExternalLocationName = *externalLocationNamePb
	}

	purposePb := &st.Purpose
	if purposePb != nil {
		pb.Purpose = *purposePb
	}

	readOnlyPb := &st.ReadOnly
	if readOnlyPb != nil {
		pb.ReadOnly = *readOnlyPb
	}

	urlPb := &st.Url
	if urlPb != nil {
		pb.Url = *urlPb
	}

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
	credentialNameField := &pb.CredentialName
	if credentialNameField != nil {
		st.CredentialName = *credentialNameField
	}
	databricksGcpServiceAccountField, err := databricksGcpServiceAccountFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	externalLocationNameField := &pb.ExternalLocationName
	if externalLocationNameField != nil {
		st.ExternalLocationName = *externalLocationNameField
	}
	purposeField := &pb.Purpose
	if purposeField != nil {
		st.Purpose = *purposeField
	}
	readOnlyField := &pb.ReadOnly
	if readOnlyField != nil {
		st.ReadOnly = *readOnlyField
	}
	urlField := &pb.Url
	if urlField != nil {
		st.Url = *urlField
	}

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
	IsDir bool
	// The results of the validation check.
	Results []CredentialValidationResult

	ForceSendFields []string
}

func validateCredentialResponseToPb(st *ValidateCredentialResponse) (*validateCredentialResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &validateCredentialResponsePb{}
	isDirPb := &st.IsDir
	if isDirPb != nil {
		pb.IsDir = *isDirPb
	}

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
	isDirField := &pb.IsDir
	if isDirField != nil {
		st.IsDir = *isDirField
	}

	var resultsField []CredentialValidationResult
	for _, item := range pb.Results {
		itemField, err := credentialValidationResultFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			resultsField = append(resultsField, *itemField)
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
	AwsIamRole *AwsIamRoleRequest
	// The Azure managed identity configuration.
	AzureManagedIdentity *AzureManagedIdentityRequest
	// The Azure service principal configuration.
	AzureServicePrincipal *AzureServicePrincipal
	// The Cloudflare API token configuration.
	CloudflareApiToken *CloudflareApiToken
	// The Databricks created GCP service account configuration.
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest
	// The name of an existing external location to validate.
	ExternalLocationName string
	// Whether the storage credential is only usable for read operations.
	ReadOnly bool
	// The name of the storage credential to validate.
	StorageCredentialName string
	// The external location url to validate.
	Url string

	ForceSendFields []string
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

	externalLocationNamePb := &st.ExternalLocationName
	if externalLocationNamePb != nil {
		pb.ExternalLocationName = *externalLocationNamePb
	}

	readOnlyPb := &st.ReadOnly
	if readOnlyPb != nil {
		pb.ReadOnly = *readOnlyPb
	}

	storageCredentialNamePb := &st.StorageCredentialName
	if storageCredentialNamePb != nil {
		pb.StorageCredentialName = *storageCredentialNamePb
	}

	urlPb := &st.Url
	if urlPb != nil {
		pb.Url = *urlPb
	}

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
	externalLocationNameField := &pb.ExternalLocationName
	if externalLocationNameField != nil {
		st.ExternalLocationName = *externalLocationNameField
	}
	readOnlyField := &pb.ReadOnly
	if readOnlyField != nil {
		st.ReadOnly = *readOnlyField
	}
	storageCredentialNameField := &pb.StorageCredentialName
	if storageCredentialNameField != nil {
		st.StorageCredentialName = *storageCredentialNameField
	}
	urlField := &pb.Url
	if urlField != nil {
		st.Url = *urlField
	}

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
	IsDir bool
	// The results of the validation check.
	Results []ValidationResult

	ForceSendFields []string
}

func validateStorageCredentialResponseToPb(st *ValidateStorageCredentialResponse) (*validateStorageCredentialResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &validateStorageCredentialResponsePb{}
	isDirPb := &st.IsDir
	if isDirPb != nil {
		pb.IsDir = *isDirPb
	}

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
	isDirField := &pb.IsDir
	if isDirField != nil {
		st.IsDir = *isDirField
	}

	var resultsField []ValidationResult
	for _, item := range pb.Results {
		itemField, err := validationResultFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			resultsField = append(resultsField, *itemField)
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
	Message string
	// The operation tested.
	Operation ValidationResultOperation
	// The results of the tested operation.
	Result ValidationResultResult

	ForceSendFields []string
}

func validationResultToPb(st *ValidationResult) (*validationResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &validationResultPb{}
	messagePb := &st.Message
	if messagePb != nil {
		pb.Message = *messagePb
	}

	operationPb := &st.Operation
	if operationPb != nil {
		pb.Operation = *operationPb
	}

	resultPb := &st.Result
	if resultPb != nil {
		pb.Result = *resultPb
	}

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
	messageField := &pb.Message
	if messageField != nil {
		st.Message = *messageField
	}
	operationField := &pb.Operation
	if operationField != nil {
		st.Operation = *operationField
	}
	resultField := &pb.Result
	if resultField != nil {
		st.Result = *resultField
	}

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
	AccessPoint string
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	BrowseOnly bool
	// The name of the catalog where the schema and the volume are
	CatalogName string
	// The comment attached to the volume
	Comment string

	CreatedAt int64
	// The identifier of the user who created the volume
	CreatedBy string
	// Encryption options that apply to clients connecting to cloud storage.
	EncryptionDetails *EncryptionDetails
	// The three-level (fully qualified) name of the volume
	FullName string
	// The unique identifier of the metastore
	MetastoreId string
	// The name of the volume
	Name string
	// The identifier of the user who owns the volume
	Owner string
	// The name of the schema where the volume is
	SchemaName string
	// The storage location on the cloud
	StorageLocation string

	UpdatedAt int64
	// The identifier of the user who updated the volume last time
	UpdatedBy string
	// The unique identifier of the volume
	VolumeId string
	// The type of the volume. An external volume is located in the specified
	// external location. A managed volume is located in the default location
	// which is specified by the parent schema, or the parent catalog, or the
	// Metastore. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/aws/en/volumes/managed-vs-external
	VolumeType VolumeType

	ForceSendFields []string
}

func volumeInfoToPb(st *VolumeInfo) (*volumeInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &volumeInfoPb{}
	accessPointPb := &st.AccessPoint
	if accessPointPb != nil {
		pb.AccessPoint = *accessPointPb
	}

	browseOnlyPb := &st.BrowseOnly
	if browseOnlyPb != nil {
		pb.BrowseOnly = *browseOnlyPb
	}

	catalogNamePb := &st.CatalogName
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	commentPb := &st.Comment
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	createdAtPb := &st.CreatedAt
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	encryptionDetailsPb, err := encryptionDetailsToPb(st.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsPb != nil {
		pb.EncryptionDetails = encryptionDetailsPb
	}

	fullNamePb := &st.FullName
	if fullNamePb != nil {
		pb.FullName = *fullNamePb
	}

	metastoreIdPb := &st.MetastoreId
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	ownerPb := &st.Owner
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	schemaNamePb := &st.SchemaName
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

	storageLocationPb := &st.StorageLocation
	if storageLocationPb != nil {
		pb.StorageLocation = *storageLocationPb
	}

	updatedAtPb := &st.UpdatedAt
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb := &st.UpdatedBy
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

	volumeIdPb := &st.VolumeId
	if volumeIdPb != nil {
		pb.VolumeId = *volumeIdPb
	}

	volumeTypePb := &st.VolumeType
	if volumeTypePb != nil {
		pb.VolumeType = *volumeTypePb
	}

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
	accessPointField := &pb.AccessPoint
	if accessPointField != nil {
		st.AccessPoint = *accessPointField
	}
	browseOnlyField := &pb.BrowseOnly
	if browseOnlyField != nil {
		st.BrowseOnly = *browseOnlyField
	}
	catalogNameField := &pb.CatalogName
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	commentField := &pb.Comment
	if commentField != nil {
		st.Comment = *commentField
	}
	createdAtField := &pb.CreatedAt
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	encryptionDetailsField, err := encryptionDetailsFromPb(pb.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsField != nil {
		st.EncryptionDetails = encryptionDetailsField
	}
	fullNameField := &pb.FullName
	if fullNameField != nil {
		st.FullName = *fullNameField
	}
	metastoreIdField := &pb.MetastoreId
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	ownerField := &pb.Owner
	if ownerField != nil {
		st.Owner = *ownerField
	}
	schemaNameField := &pb.SchemaName
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}
	storageLocationField := &pb.StorageLocation
	if storageLocationField != nil {
		st.StorageLocation = *storageLocationField
	}
	updatedAtField := &pb.UpdatedAt
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField := &pb.UpdatedBy
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}
	volumeIdField := &pb.VolumeId
	if volumeIdField != nil {
		st.VolumeId = *volumeIdField
	}
	volumeTypeField := &pb.VolumeType
	if volumeTypeField != nil {
		st.VolumeType = *volumeTypeField
	}

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
	BindingType WorkspaceBindingBindingType

	WorkspaceId int64

	ForceSendFields []string
}

func workspaceBindingToPb(st *WorkspaceBinding) (*workspaceBindingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceBindingPb{}
	bindingTypePb := &st.BindingType
	if bindingTypePb != nil {
		pb.BindingType = *bindingTypePb
	}

	workspaceIdPb := &st.WorkspaceId
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

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
	bindingTypeField := &pb.BindingType
	if bindingTypeField != nil {
		st.BindingType = *bindingTypeField
	}
	workspaceIdField := &pb.WorkspaceId
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}

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
	Bindings []WorkspaceBinding
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string

	ForceSendFields []string
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

	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	for _, item := range pb.Bindings {
		itemField, err := workspaceBindingFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			bindingsField = append(bindingsField, *itemField)
		}
	}
	st.Bindings = bindingsField
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *workspaceBindingsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st workspaceBindingsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
