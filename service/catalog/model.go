// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/catalog/catalogpb"
)

type AccountsCreateMetastore struct {

	// Wire name: 'metastore_info'
	MetastoreInfo *CreateMetastore ``
}

func AccountsCreateMetastoreToPb(st *AccountsCreateMetastore) (*catalogpb.AccountsCreateMetastorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AccountsCreateMetastorePb{}
	metastoreInfoPb, err := CreateMetastoreToPb(st.MetastoreInfo)
	if err != nil {
		return nil, err
	}
	if metastoreInfoPb != nil {
		pb.MetastoreInfo = metastoreInfoPb
	}

	return pb, nil
}

func AccountsCreateMetastoreFromPb(pb *catalogpb.AccountsCreateMetastorePb) (*AccountsCreateMetastore, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsCreateMetastore{}
	metastoreInfoField, err := CreateMetastoreFromPb(pb.MetastoreInfo)
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
	MetastoreAssignment *CreateMetastoreAssignment ``
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func AccountsCreateMetastoreAssignmentToPb(st *AccountsCreateMetastoreAssignment) (*catalogpb.AccountsCreateMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AccountsCreateMetastoreAssignmentPb{}
	metastoreAssignmentPb, err := CreateMetastoreAssignmentToPb(st.MetastoreAssignment)
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

func AccountsCreateMetastoreAssignmentFromPb(pb *catalogpb.AccountsCreateMetastoreAssignmentPb) (*AccountsCreateMetastoreAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsCreateMetastoreAssignment{}
	metastoreAssignmentField, err := CreateMetastoreAssignmentFromPb(pb.MetastoreAssignment)
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
	CredentialInfo *CreateStorageCredential ``
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
}

func AccountsCreateStorageCredentialToPb(st *AccountsCreateStorageCredential) (*catalogpb.AccountsCreateStorageCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AccountsCreateStorageCredentialPb{}
	credentialInfoPb, err := CreateStorageCredentialToPb(st.CredentialInfo)
	if err != nil {
		return nil, err
	}
	if credentialInfoPb != nil {
		pb.CredentialInfo = credentialInfoPb
	}
	pb.MetastoreId = st.MetastoreId

	return pb, nil
}

func AccountsCreateStorageCredentialFromPb(pb *catalogpb.AccountsCreateStorageCredentialPb) (*AccountsCreateStorageCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsCreateStorageCredential{}
	credentialInfoField, err := CreateStorageCredentialFromPb(pb.CredentialInfo)
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
	MetastoreAssignment *MetastoreAssignment ``
}

func AccountsMetastoreAssignmentToPb(st *AccountsMetastoreAssignment) (*catalogpb.AccountsMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AccountsMetastoreAssignmentPb{}
	metastoreAssignmentPb, err := MetastoreAssignmentToPb(st.MetastoreAssignment)
	if err != nil {
		return nil, err
	}
	if metastoreAssignmentPb != nil {
		pb.MetastoreAssignment = metastoreAssignmentPb
	}

	return pb, nil
}

func AccountsMetastoreAssignmentFromPb(pb *catalogpb.AccountsMetastoreAssignmentPb) (*AccountsMetastoreAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsMetastoreAssignment{}
	metastoreAssignmentField, err := MetastoreAssignmentFromPb(pb.MetastoreAssignment)
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
	MetastoreInfo *MetastoreInfo ``
}

func AccountsMetastoreInfoToPb(st *AccountsMetastoreInfo) (*catalogpb.AccountsMetastoreInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AccountsMetastoreInfoPb{}
	metastoreInfoPb, err := MetastoreInfoToPb(st.MetastoreInfo)
	if err != nil {
		return nil, err
	}
	if metastoreInfoPb != nil {
		pb.MetastoreInfo = metastoreInfoPb
	}

	return pb, nil
}

func AccountsMetastoreInfoFromPb(pb *catalogpb.AccountsMetastoreInfoPb) (*AccountsMetastoreInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsMetastoreInfo{}
	metastoreInfoField, err := MetastoreInfoFromPb(pb.MetastoreInfo)
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
	CredentialInfo *StorageCredentialInfo ``
}

func AccountsStorageCredentialInfoToPb(st *AccountsStorageCredentialInfo) (*catalogpb.AccountsStorageCredentialInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AccountsStorageCredentialInfoPb{}
	credentialInfoPb, err := StorageCredentialInfoToPb(st.CredentialInfo)
	if err != nil {
		return nil, err
	}
	if credentialInfoPb != nil {
		pb.CredentialInfo = credentialInfoPb
	}

	return pb, nil
}

func AccountsStorageCredentialInfoFromPb(pb *catalogpb.AccountsStorageCredentialInfoPb) (*AccountsStorageCredentialInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsStorageCredentialInfo{}
	credentialInfoField, err := StorageCredentialInfoFromPb(pb.CredentialInfo)
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
	MetastoreInfo *UpdateMetastore ``
}

func AccountsUpdateMetastoreToPb(st *AccountsUpdateMetastore) (*catalogpb.AccountsUpdateMetastorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AccountsUpdateMetastorePb{}
	pb.MetastoreId = st.MetastoreId
	metastoreInfoPb, err := UpdateMetastoreToPb(st.MetastoreInfo)
	if err != nil {
		return nil, err
	}
	if metastoreInfoPb != nil {
		pb.MetastoreInfo = metastoreInfoPb
	}

	return pb, nil
}

func AccountsUpdateMetastoreFromPb(pb *catalogpb.AccountsUpdateMetastorePb) (*AccountsUpdateMetastore, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsUpdateMetastore{}
	st.MetastoreId = pb.MetastoreId
	metastoreInfoField, err := UpdateMetastoreFromPb(pb.MetastoreInfo)
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
	MetastoreAssignment *UpdateMetastoreAssignment ``
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func AccountsUpdateMetastoreAssignmentToPb(st *AccountsUpdateMetastoreAssignment) (*catalogpb.AccountsUpdateMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AccountsUpdateMetastoreAssignmentPb{}
	metastoreAssignmentPb, err := UpdateMetastoreAssignmentToPb(st.MetastoreAssignment)
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

func AccountsUpdateMetastoreAssignmentFromPb(pb *catalogpb.AccountsUpdateMetastoreAssignmentPb) (*AccountsUpdateMetastoreAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsUpdateMetastoreAssignment{}
	metastoreAssignmentField, err := UpdateMetastoreAssignmentFromPb(pb.MetastoreAssignment)
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
	CredentialInfo *UpdateStorageCredential ``
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Name of the storage credential.
	// Wire name: 'storage_credential_name'
	StorageCredentialName string `tf:"-"`
}

func AccountsUpdateStorageCredentialToPb(st *AccountsUpdateStorageCredential) (*catalogpb.AccountsUpdateStorageCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AccountsUpdateStorageCredentialPb{}
	credentialInfoPb, err := UpdateStorageCredentialToPb(st.CredentialInfo)
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

func AccountsUpdateStorageCredentialFromPb(pb *catalogpb.AccountsUpdateStorageCredentialPb) (*AccountsUpdateStorageCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsUpdateStorageCredential{}
	credentialInfoField, err := UpdateStorageCredentialFromPb(pb.CredentialInfo)
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
	ArtifactMatchers []ArtifactMatcher ``
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of the user who set the artifact allowlist.
	// Wire name: 'created_by'
	CreatedBy string ``
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ArtifactAllowlistInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ArtifactAllowlistInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ArtifactAllowlistInfoToPb(st *ArtifactAllowlistInfo) (*catalogpb.ArtifactAllowlistInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ArtifactAllowlistInfoPb{}

	var artifactMatchersPb []catalogpb.ArtifactMatcherPb
	for _, item := range st.ArtifactMatchers {
		itemPb, err := ArtifactMatcherToPb(&item)
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

func ArtifactAllowlistInfoFromPb(pb *catalogpb.ArtifactAllowlistInfoPb) (*ArtifactAllowlistInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ArtifactAllowlistInfo{}

	var artifactMatchersField []ArtifactMatcher
	for _, itemPb := range pb.ArtifactMatchers {
		item, err := ArtifactMatcherFromPb(&itemPb)
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

type ArtifactMatcher struct {
	// The artifact path or maven coordinate
	// Wire name: 'artifact'
	Artifact string ``
	// The pattern matching type of the artifact
	// Wire name: 'match_type'
	MatchType MatchType ``
}

func ArtifactMatcherToPb(st *ArtifactMatcher) (*catalogpb.ArtifactMatcherPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ArtifactMatcherPb{}
	pb.Artifact = st.Artifact
	matchTypePb, err := MatchTypeToPb(&st.MatchType)
	if err != nil {
		return nil, err
	}
	if matchTypePb != nil {
		pb.MatchType = *matchTypePb
	}

	return pb, nil
}

func ArtifactMatcherFromPb(pb *catalogpb.ArtifactMatcherPb) (*ArtifactMatcher, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ArtifactMatcher{}
	st.Artifact = pb.Artifact
	matchTypeField, err := MatchTypeFromPb(&pb.MatchType)
	if err != nil {
		return nil, err
	}
	if matchTypeField != nil {
		st.MatchType = *matchTypeField
	}

	return st, nil
}

// The artifact type
type ArtifactType string

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

// Values returns all possible values for ArtifactType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ArtifactType) Values() []ArtifactType {
	return []ArtifactType{
		ArtifactTypeInitScript,
		ArtifactTypeLibraryJar,
		ArtifactTypeLibraryMaven,
	}
}

// Type always returns ArtifactType to satisfy [pflag.Value] interface
func (f *ArtifactType) Type() string {
	return "ArtifactType"
}

func ArtifactTypeToPb(st *ArtifactType) (*catalogpb.ArtifactTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.ArtifactTypePb(*st)
	return &pb, nil
}

func ArtifactTypeFromPb(pb *catalogpb.ArtifactTypePb) (*ArtifactType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ArtifactType(*pb)
	return &st, nil
}

// AWS temporary credentials for API authentication. Read more at
// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
type AwsCredentials struct {
	// The access key ID that identifies the temporary credentials.
	// Wire name: 'access_key_id'
	AccessKeyId string ``
	// The Amazon Resource Name (ARN) of the S3 access point for temporary
	// credentials related the external location.
	// Wire name: 'access_point'
	AccessPoint string ``
	// The secret access key that can be used to sign AWS API requests.
	// Wire name: 'secret_access_key'
	SecretAccessKey string ``
	// The token that users must pass to AWS API to use the temporary
	// credentials.
	// Wire name: 'session_token'
	SessionToken    string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AwsCredentials) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsCredentials) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AwsCredentialsToPb(st *AwsCredentials) (*catalogpb.AwsCredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AwsCredentialsPb{}
	pb.AccessKeyId = st.AccessKeyId
	pb.AccessPoint = st.AccessPoint
	pb.SecretAccessKey = st.SecretAccessKey
	pb.SessionToken = st.SessionToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AwsCredentialsFromPb(pb *catalogpb.AwsCredentialsPb) (*AwsCredentials, error) {
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

// The AWS IAM role configuration
type AwsIamRole struct {
	// The external ID used in role assumption to prevent the confused deputy
	// problem.
	// Wire name: 'external_id'
	ExternalId string ``
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	// Wire name: 'role_arn'
	RoleArn string ``
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	// Wire name: 'unity_catalog_iam_arn'
	UnityCatalogIamArn string   ``
	ForceSendFields    []string `tf:"-"`
}

func (s *AwsIamRole) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsIamRole) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AwsIamRoleToPb(st *AwsIamRole) (*catalogpb.AwsIamRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AwsIamRolePb{}
	pb.ExternalId = st.ExternalId
	pb.RoleArn = st.RoleArn
	pb.UnityCatalogIamArn = st.UnityCatalogIamArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AwsIamRoleFromPb(pb *catalogpb.AwsIamRolePb) (*AwsIamRole, error) {
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

// The AWS IAM role configuration
type AwsIamRoleRequest struct {
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	// Wire name: 'role_arn'
	RoleArn string ``
}

func AwsIamRoleRequestToPb(st *AwsIamRoleRequest) (*catalogpb.AwsIamRoleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AwsIamRoleRequestPb{}
	pb.RoleArn = st.RoleArn

	return pb, nil
}

func AwsIamRoleRequestFromPb(pb *catalogpb.AwsIamRoleRequestPb) (*AwsIamRoleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsIamRoleRequest{}
	st.RoleArn = pb.RoleArn

	return st, nil
}

// The AWS IAM role configuration
type AwsIamRoleResponse struct {
	// The external ID used in role assumption to prevent the confused deputy
	// problem.
	// Wire name: 'external_id'
	ExternalId string ``
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	// Wire name: 'role_arn'
	RoleArn string ``
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	// Wire name: 'unity_catalog_iam_arn'
	UnityCatalogIamArn string   ``
	ForceSendFields    []string `tf:"-"`
}

func (s *AwsIamRoleResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsIamRoleResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AwsIamRoleResponseToPb(st *AwsIamRoleResponse) (*catalogpb.AwsIamRoleResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AwsIamRoleResponsePb{}
	pb.ExternalId = st.ExternalId
	pb.RoleArn = st.RoleArn
	pb.UnityCatalogIamArn = st.UnityCatalogIamArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AwsIamRoleResponseFromPb(pb *catalogpb.AwsIamRoleResponsePb) (*AwsIamRoleResponse, error) {
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

type AwsSqsQueue struct {
	// Unique identifier included in the name of file events managed cloud
	// resources.
	// Wire name: 'managed_resource_id'
	ManagedResourceId string ``
	// The AQS queue url in the format
	// https://sqs.{region}.amazonaws.com/{account id}/{queue name} Required for
	// provided_sqs.
	// Wire name: 'queue_url'
	QueueUrl        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AwsSqsQueue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsSqsQueue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AwsSqsQueueToPb(st *AwsSqsQueue) (*catalogpb.AwsSqsQueuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AwsSqsQueuePb{}
	pb.ManagedResourceId = st.ManagedResourceId
	pb.QueueUrl = st.QueueUrl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AwsSqsQueueFromPb(pb *catalogpb.AwsSqsQueuePb) (*AwsSqsQueue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsSqsQueue{}
	st.ManagedResourceId = pb.ManagedResourceId
	st.QueueUrl = pb.QueueUrl

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Azure Active Directory token, essentially the Oauth token for Azure Service
// Principal or Managed Identity. Read more at
// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
type AzureActiveDirectoryToken struct {
	// Opaque token that contains claims that you can use in Azure Active
	// Directory to access cloud services.
	// Wire name: 'aad_token'
	AadToken        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AzureActiveDirectoryToken) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureActiveDirectoryToken) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AzureActiveDirectoryTokenToPb(st *AzureActiveDirectoryToken) (*catalogpb.AzureActiveDirectoryTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AzureActiveDirectoryTokenPb{}
	pb.AadToken = st.AadToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AzureActiveDirectoryTokenFromPb(pb *catalogpb.AzureActiveDirectoryTokenPb) (*AzureActiveDirectoryToken, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureActiveDirectoryToken{}
	st.AadToken = pb.AadToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// The Azure managed identity configuration.
type AzureManagedIdentity struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}`.
	// Wire name: 'access_connector_id'
	AccessConnectorId string ``
	// The Databricks internal ID that represents this managed identity.
	// Wire name: 'credential_id'
	CredentialId string ``
	// The Azure resource ID of the managed identity. Use the format,
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}`
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// using the system-assigned identity.
	// Wire name: 'managed_identity_id'
	ManagedIdentityId string   ``
	ForceSendFields   []string `tf:"-"`
}

func (s *AzureManagedIdentity) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureManagedIdentity) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AzureManagedIdentityToPb(st *AzureManagedIdentity) (*catalogpb.AzureManagedIdentityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AzureManagedIdentityPb{}
	pb.AccessConnectorId = st.AccessConnectorId
	pb.CredentialId = st.CredentialId
	pb.ManagedIdentityId = st.ManagedIdentityId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AzureManagedIdentityFromPb(pb *catalogpb.AzureManagedIdentityPb) (*AzureManagedIdentity, error) {
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

// The Azure managed identity configuration.
type AzureManagedIdentityRequest struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}`.
	// Wire name: 'access_connector_id'
	AccessConnectorId string ``
	// The Azure resource ID of the managed identity. Use the format,
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}`
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// using the system-assigned identity.
	// Wire name: 'managed_identity_id'
	ManagedIdentityId string   ``
	ForceSendFields   []string `tf:"-"`
}

func (s *AzureManagedIdentityRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureManagedIdentityRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AzureManagedIdentityRequestToPb(st *AzureManagedIdentityRequest) (*catalogpb.AzureManagedIdentityRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AzureManagedIdentityRequestPb{}
	pb.AccessConnectorId = st.AccessConnectorId
	pb.ManagedIdentityId = st.ManagedIdentityId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AzureManagedIdentityRequestFromPb(pb *catalogpb.AzureManagedIdentityRequestPb) (*AzureManagedIdentityRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureManagedIdentityRequest{}
	st.AccessConnectorId = pb.AccessConnectorId
	st.ManagedIdentityId = pb.ManagedIdentityId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// The Azure managed identity configuration.
type AzureManagedIdentityResponse struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}`.
	// Wire name: 'access_connector_id'
	AccessConnectorId string ``
	// The Databricks internal ID that represents this managed identity.
	// Wire name: 'credential_id'
	CredentialId string ``
	// The Azure resource ID of the managed identity. Use the format,
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}`
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// using the system-assigned identity.
	// Wire name: 'managed_identity_id'
	ManagedIdentityId string   ``
	ForceSendFields   []string `tf:"-"`
}

func (s *AzureManagedIdentityResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureManagedIdentityResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AzureManagedIdentityResponseToPb(st *AzureManagedIdentityResponse) (*catalogpb.AzureManagedIdentityResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AzureManagedIdentityResponsePb{}
	pb.AccessConnectorId = st.AccessConnectorId
	pb.CredentialId = st.CredentialId
	pb.ManagedIdentityId = st.ManagedIdentityId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AzureManagedIdentityResponseFromPb(pb *catalogpb.AzureManagedIdentityResponsePb) (*AzureManagedIdentityResponse, error) {
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

type AzureQueueStorage struct {
	// Unique identifier included in the name of file events managed cloud
	// resources.
	// Wire name: 'managed_resource_id'
	ManagedResourceId string ``
	// The AQS queue url in the format https://{storage
	// account}.queue.core.windows.net/{queue name} Required for provided_aqs.
	// Wire name: 'queue_url'
	QueueUrl string ``
	// The resource group for the queue, event grid subscription, and external
	// location storage account. Only required for locations with a service
	// principal storage credential
	// Wire name: 'resource_group'
	ResourceGroup string ``
	// Optional subscription id for the queue, event grid subscription, and
	// external location storage account. Required for locations with a service
	// principal storage credential
	// Wire name: 'subscription_id'
	SubscriptionId  string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AzureQueueStorage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureQueueStorage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AzureQueueStorageToPb(st *AzureQueueStorage) (*catalogpb.AzureQueueStoragePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AzureQueueStoragePb{}
	pb.ManagedResourceId = st.ManagedResourceId
	pb.QueueUrl = st.QueueUrl
	pb.ResourceGroup = st.ResourceGroup
	pb.SubscriptionId = st.SubscriptionId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AzureQueueStorageFromPb(pb *catalogpb.AzureQueueStoragePb) (*AzureQueueStorage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureQueueStorage{}
	st.ManagedResourceId = pb.ManagedResourceId
	st.QueueUrl = pb.QueueUrl
	st.ResourceGroup = pb.ResourceGroup
	st.SubscriptionId = pb.SubscriptionId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// The Azure service principal configuration. Only applicable when purpose is
// **STORAGE**.
type AzureServicePrincipal struct {
	// The application ID of the application registration within the referenced
	// AAD tenant.
	// Wire name: 'application_id'
	ApplicationId string ``
	// The client secret generated for the above app ID in AAD.
	// Wire name: 'client_secret'
	ClientSecret string ``
	// The directory ID corresponding to the Azure Active Directory (AAD) tenant
	// of the application.
	// Wire name: 'directory_id'
	DirectoryId string ``
}

func AzureServicePrincipalToPb(st *AzureServicePrincipal) (*catalogpb.AzureServicePrincipalPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AzureServicePrincipalPb{}
	pb.ApplicationId = st.ApplicationId
	pb.ClientSecret = st.ClientSecret
	pb.DirectoryId = st.DirectoryId

	return pb, nil
}

func AzureServicePrincipalFromPb(pb *catalogpb.AzureServicePrincipalPb) (*AzureServicePrincipal, error) {
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
	SasToken        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AzureUserDelegationSas) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureUserDelegationSas) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AzureUserDelegationSasToPb(st *AzureUserDelegationSas) (*catalogpb.AzureUserDelegationSasPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.AzureUserDelegationSasPb{}
	pb.SasToken = st.SasToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AzureUserDelegationSasFromPb(pb *catalogpb.AzureUserDelegationSasPb) (*AzureUserDelegationSas, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureUserDelegationSas{}
	st.SasToken = pb.SasToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CancelRefreshRequest struct {

	// Wire name: 'refresh_id'
	RefreshId int64 `tf:"-"`
	// UC table name in format `catalog.schema.table_name`. table_name is case
	// insensitive and spaces are disallowed.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func CancelRefreshRequestToPb(st *CancelRefreshRequest) (*catalogpb.CancelRefreshRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CancelRefreshRequestPb{}
	pb.RefreshId = st.RefreshId
	pb.TableName = st.TableName

	return pb, nil
}

func CancelRefreshRequestFromPb(pb *catalogpb.CancelRefreshRequestPb) (*CancelRefreshRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelRefreshRequest{}
	st.RefreshId = pb.RefreshId
	st.TableName = pb.TableName

	return st, nil
}

type CatalogInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool ``

	// Wire name: 'catalog_type'
	CatalogType CatalogType ``
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// The name of the connection to an external data source.
	// Wire name: 'connection_name'
	ConnectionName string ``
	// Time at which this catalog was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of catalog creator.
	// Wire name: 'created_by'
	CreatedBy string ``

	// Wire name: 'effective_predictive_optimization_flag'
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag ``
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'enable_predictive_optimization'
	EnablePredictiveOptimization EnablePredictiveOptimization ``
	// The full name of the catalog. Corresponds with the name field.
	// Wire name: 'full_name'
	FullName string ``
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode CatalogIsolationMode ``
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// Name of catalog.
	// Wire name: 'name'
	Name string ``
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string ``
	// Username of current owner of catalog.
	// Wire name: 'owner'
	Owner string ``
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string ``
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	// Wire name: 'provider_name'
	ProviderName string ``

	// Wire name: 'provisioning_info'
	ProvisioningInfo *ProvisioningInfo ``

	// Wire name: 'securable_type'
	SecurableType SecurableType ``
	// The name of the share under the share provider.
	// Wire name: 'share_name'
	ShareName string ``
	// Storage Location URL (full path) for managed tables within catalog.
	// Wire name: 'storage_location'
	StorageLocation string ``
	// Storage root URL for managed tables within catalog.
	// Wire name: 'storage_root'
	StorageRoot string ``
	// Time at which this catalog was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// Username of user who last modified catalog.
	// Wire name: 'updated_by'
	UpdatedBy       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CatalogInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CatalogInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CatalogInfoToPb(st *CatalogInfo) (*catalogpb.CatalogInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CatalogInfoPb{}
	pb.BrowseOnly = st.BrowseOnly
	catalogTypePb, err := CatalogTypeToPb(&st.CatalogType)
	if err != nil {
		return nil, err
	}
	if catalogTypePb != nil {
		pb.CatalogType = *catalogTypePb
	}
	pb.Comment = st.Comment
	pb.ConnectionName = st.ConnectionName
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	effectivePredictiveOptimizationFlagPb, err := EffectivePredictiveOptimizationFlagToPb(st.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagPb != nil {
		pb.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagPb
	}
	enablePredictiveOptimizationPb, err := EnablePredictiveOptimizationToPb(&st.EnablePredictiveOptimization)
	if err != nil {
		return nil, err
	}
	if enablePredictiveOptimizationPb != nil {
		pb.EnablePredictiveOptimization = *enablePredictiveOptimizationPb
	}
	pb.FullName = st.FullName
	isolationModePb, err := CatalogIsolationModeToPb(&st.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Options = st.Options
	pb.Owner = st.Owner
	pb.Properties = st.Properties
	pb.ProviderName = st.ProviderName
	provisioningInfoPb, err := ProvisioningInfoToPb(st.ProvisioningInfo)
	if err != nil {
		return nil, err
	}
	if provisioningInfoPb != nil {
		pb.ProvisioningInfo = provisioningInfoPb
	}
	securableTypePb, err := SecurableTypeToPb(&st.SecurableType)
	if err != nil {
		return nil, err
	}
	if securableTypePb != nil {
		pb.SecurableType = *securableTypePb
	}
	pb.ShareName = st.ShareName
	pb.StorageLocation = st.StorageLocation
	pb.StorageRoot = st.StorageRoot
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CatalogInfoFromPb(pb *catalogpb.CatalogInfoPb) (*CatalogInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CatalogInfo{}
	st.BrowseOnly = pb.BrowseOnly
	catalogTypeField, err := CatalogTypeFromPb(&pb.CatalogType)
	if err != nil {
		return nil, err
	}
	if catalogTypeField != nil {
		st.CatalogType = *catalogTypeField
	}
	st.Comment = pb.Comment
	st.ConnectionName = pb.ConnectionName
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	effectivePredictiveOptimizationFlagField, err := EffectivePredictiveOptimizationFlagFromPb(pb.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagField != nil {
		st.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagField
	}
	enablePredictiveOptimizationField, err := EnablePredictiveOptimizationFromPb(&pb.EnablePredictiveOptimization)
	if err != nil {
		return nil, err
	}
	if enablePredictiveOptimizationField != nil {
		st.EnablePredictiveOptimization = *enablePredictiveOptimizationField
	}
	st.FullName = pb.FullName
	isolationModeField, err := CatalogIsolationModeFromPb(&pb.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Options = pb.Options
	st.Owner = pb.Owner
	st.Properties = pb.Properties
	st.ProviderName = pb.ProviderName
	provisioningInfoField, err := ProvisioningInfoFromPb(pb.ProvisioningInfo)
	if err != nil {
		return nil, err
	}
	if provisioningInfoField != nil {
		st.ProvisioningInfo = provisioningInfoField
	}
	securableTypeField, err := SecurableTypeFromPb(&pb.SecurableType)
	if err != nil {
		return nil, err
	}
	if securableTypeField != nil {
		st.SecurableType = *securableTypeField
	}
	st.ShareName = pb.ShareName
	st.StorageLocation = pb.StorageLocation
	st.StorageRoot = pb.StorageRoot
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CatalogIsolationMode string

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

// Values returns all possible values for CatalogIsolationMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *CatalogIsolationMode) Values() []CatalogIsolationMode {
	return []CatalogIsolationMode{
		CatalogIsolationModeIsolated,
		CatalogIsolationModeOpen,
	}
}

// Type always returns CatalogIsolationMode to satisfy [pflag.Value] interface
func (f *CatalogIsolationMode) Type() string {
	return "CatalogIsolationMode"
}

func CatalogIsolationModeToPb(st *CatalogIsolationMode) (*catalogpb.CatalogIsolationModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.CatalogIsolationModePb(*st)
	return &pb, nil
}

func CatalogIsolationModeFromPb(pb *catalogpb.CatalogIsolationModePb) (*CatalogIsolationMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := CatalogIsolationMode(*pb)
	return &st, nil
}

// The type of the catalog.
type CatalogType string

const CatalogTypeDeltasharingCatalog CatalogType = `DELTASHARING_CATALOG`

const CatalogTypeForeignCatalog CatalogType = `FOREIGN_CATALOG`

const CatalogTypeInternalCatalog CatalogType = `INTERNAL_CATALOG`

const CatalogTypeManagedCatalog CatalogType = `MANAGED_CATALOG`

const CatalogTypeManagedOnlineCatalog CatalogType = `MANAGED_ONLINE_CATALOG`

const CatalogTypeSystemCatalog CatalogType = `SYSTEM_CATALOG`

// String representation for [fmt.Print]
func (f *CatalogType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CatalogType) Set(v string) error {
	switch v {
	case `DELTASHARING_CATALOG`, `FOREIGN_CATALOG`, `INTERNAL_CATALOG`, `MANAGED_CATALOG`, `MANAGED_ONLINE_CATALOG`, `SYSTEM_CATALOG`:
		*f = CatalogType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELTASHARING_CATALOG", "FOREIGN_CATALOG", "INTERNAL_CATALOG", "MANAGED_CATALOG", "MANAGED_ONLINE_CATALOG", "SYSTEM_CATALOG"`, v)
	}
}

// Values returns all possible values for CatalogType.
//
// There is no guarantee on the order of the values in the slice.
func (f *CatalogType) Values() []CatalogType {
	return []CatalogType{
		CatalogTypeDeltasharingCatalog,
		CatalogTypeForeignCatalog,
		CatalogTypeInternalCatalog,
		CatalogTypeManagedCatalog,
		CatalogTypeManagedOnlineCatalog,
		CatalogTypeSystemCatalog,
	}
}

// Type always returns CatalogType to satisfy [pflag.Value] interface
func (f *CatalogType) Type() string {
	return "CatalogType"
}

func CatalogTypeToPb(st *CatalogType) (*catalogpb.CatalogTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.CatalogTypePb(*st)
	return &pb, nil
}

func CatalogTypeFromPb(pb *catalogpb.CatalogTypePb) (*CatalogType, error) {
	if pb == nil {
		return nil, nil
	}
	st := CatalogType(*pb)
	return &st, nil
}

// The Cloudflare API token configuration. Read more at
// https://developers.cloudflare.com/r2/api/s3/tokens/
type CloudflareApiToken struct {
	// The access key ID associated with the API token.
	// Wire name: 'access_key_id'
	AccessKeyId string ``
	// The ID of the account associated with the API token.
	// Wire name: 'account_id'
	AccountId string ``
	// The secret access token generated for the above access key ID.
	// Wire name: 'secret_access_key'
	SecretAccessKey string ``
}

func CloudflareApiTokenToPb(st *CloudflareApiToken) (*catalogpb.CloudflareApiTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CloudflareApiTokenPb{}
	pb.AccessKeyId = st.AccessKeyId
	pb.AccountId = st.AccountId
	pb.SecretAccessKey = st.SecretAccessKey

	return pb, nil
}

func CloudflareApiTokenFromPb(pb *catalogpb.CloudflareApiTokenPb) (*CloudflareApiToken, error) {
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
	Comment string ``

	// Wire name: 'mask'
	Mask *ColumnMask ``
	// Name of Column.
	// Wire name: 'name'
	Name string ``
	// Whether field may be Null (default: true).
	// Wire name: 'nullable'
	Nullable bool ``
	// Partition index for column.
	// Wire name: 'partition_index'
	PartitionIndex int ``
	// Ordinal position of column (starting at position 0).
	// Wire name: 'position'
	Position int ``
	// Format of IntervalType.
	// Wire name: 'type_interval_type'
	TypeIntervalType string ``
	// Full data type specification, JSON-serialized.
	// Wire name: 'type_json'
	TypeJson string ``

	// Wire name: 'type_name'
	TypeName ColumnTypeName ``
	// Digits of precision; required for DecimalTypes.
	// Wire name: 'type_precision'
	TypePrecision int ``
	// Digits to right of decimal; Required for DecimalTypes.
	// Wire name: 'type_scale'
	TypeScale int ``
	// Full data type specification as SQL/catalogString text.
	// Wire name: 'type_text'
	TypeText        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ColumnInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ColumnInfoToPb(st *ColumnInfo) (*catalogpb.ColumnInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ColumnInfoPb{}
	pb.Comment = st.Comment
	maskPb, err := ColumnMaskToPb(st.Mask)
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
	typeNamePb, err := ColumnTypeNameToPb(&st.TypeName)
	if err != nil {
		return nil, err
	}
	if typeNamePb != nil {
		pb.TypeName = *typeNamePb
	}
	pb.TypePrecision = st.TypePrecision
	pb.TypeScale = st.TypeScale
	pb.TypeText = st.TypeText

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ColumnInfoFromPb(pb *catalogpb.ColumnInfoPb) (*ColumnInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ColumnInfo{}
	st.Comment = pb.Comment
	maskField, err := ColumnMaskFromPb(pb.Mask)
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
	typeNameField, err := ColumnTypeNameFromPb(&pb.TypeName)
	if err != nil {
		return nil, err
	}
	if typeNameField != nil {
		st.TypeName = *typeNameField
	}
	st.TypePrecision = pb.TypePrecision
	st.TypeScale = pb.TypeScale
	st.TypeText = pb.TypeText

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ColumnMask struct {
	// The full name of the column mask SQL UDF.
	// Wire name: 'function_name'
	FunctionName string ``
	// The list of additional table columns to be passed as input to the column
	// mask function. The first arg of the mask function should be of the type
	// of the column being masked and the types of the rest of the args should
	// match the types of columns in 'using_column_names'.
	// Wire name: 'using_column_names'
	UsingColumnNames []string ``
	ForceSendFields  []string `tf:"-"`
}

func (s *ColumnMask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnMask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ColumnMaskToPb(st *ColumnMask) (*catalogpb.ColumnMaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ColumnMaskPb{}
	pb.FunctionName = st.FunctionName
	pb.UsingColumnNames = st.UsingColumnNames

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ColumnMaskFromPb(pb *catalogpb.ColumnMaskPb) (*ColumnMask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ColumnMask{}
	st.FunctionName = pb.FunctionName
	st.UsingColumnNames = pb.UsingColumnNames

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ColumnRelationship struct {

	// Wire name: 'source'
	Source string ``

	// Wire name: 'target'
	Target          string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ColumnRelationship) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnRelationship) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ColumnRelationshipToPb(st *ColumnRelationship) (*catalogpb.ColumnRelationshipPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ColumnRelationshipPb{}
	pb.Source = st.Source
	pb.Target = st.Target

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ColumnRelationshipFromPb(pb *catalogpb.ColumnRelationshipPb) (*ColumnRelationship, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ColumnRelationship{}
	st.Source = pb.Source
	st.Target = pb.Target

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ColumnTypeName string

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

// Values returns all possible values for ColumnTypeName.
//
// There is no guarantee on the order of the values in the slice.
func (f *ColumnTypeName) Values() []ColumnTypeName {
	return []ColumnTypeName{
		ColumnTypeNameArray,
		ColumnTypeNameBinary,
		ColumnTypeNameBoolean,
		ColumnTypeNameByte,
		ColumnTypeNameChar,
		ColumnTypeNameDate,
		ColumnTypeNameDecimal,
		ColumnTypeNameDouble,
		ColumnTypeNameFloat,
		ColumnTypeNameGeography,
		ColumnTypeNameGeometry,
		ColumnTypeNameInt,
		ColumnTypeNameInterval,
		ColumnTypeNameLong,
		ColumnTypeNameMap,
		ColumnTypeNameNull,
		ColumnTypeNameShort,
		ColumnTypeNameString,
		ColumnTypeNameStruct,
		ColumnTypeNameTableType,
		ColumnTypeNameTimestamp,
		ColumnTypeNameTimestampNtz,
		ColumnTypeNameUserDefinedType,
		ColumnTypeNameVariant,
	}
}

// Type always returns ColumnTypeName to satisfy [pflag.Value] interface
func (f *ColumnTypeName) Type() string {
	return "ColumnTypeName"
}

func ColumnTypeNameToPb(st *ColumnTypeName) (*catalogpb.ColumnTypeNamePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.ColumnTypeNamePb(*st)
	return &pb, nil
}

func ColumnTypeNameFromPb(pb *catalogpb.ColumnTypeNamePb) (*ColumnTypeName, error) {
	if pb == nil {
		return nil, nil
	}
	st := ColumnTypeName(*pb)
	return &st, nil
}

// A connection that is dependent on a SQL object.
type ConnectionDependency struct {
	// Full name of the dependent connection, in the form of
	// __connection_name__.
	// Wire name: 'connection_name'
	ConnectionName  string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ConnectionDependency) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ConnectionDependency) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ConnectionDependencyToPb(st *ConnectionDependency) (*catalogpb.ConnectionDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ConnectionDependencyPb{}
	pb.ConnectionName = st.ConnectionName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ConnectionDependencyFromPb(pb *catalogpb.ConnectionDependencyPb) (*ConnectionDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ConnectionDependency{}
	st.ConnectionName = pb.ConnectionName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Next ID: 23
type ConnectionInfo struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// Unique identifier of the Connection.
	// Wire name: 'connection_id'
	ConnectionId string ``
	// The type of connection.
	// Wire name: 'connection_type'
	ConnectionType ConnectionType ``
	// Time at which this connection was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of connection creator.
	// Wire name: 'created_by'
	CreatedBy string ``
	// The type of credential.
	// Wire name: 'credential_type'
	CredentialType CredentialType ``
	// [Create,Update:OPT] Connection environment settings as
	// EnvironmentSettings object.
	// Wire name: 'environment_settings'
	EnvironmentSettings *EnvironmentSettings ``
	// Full name of connection.
	// Wire name: 'full_name'
	FullName string ``
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// Name of the connection.
	// Wire name: 'name'
	Name string ``
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string ``
	// Username of current owner of the connection.
	// Wire name: 'owner'
	Owner string ``
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string ``

	// Wire name: 'provisioning_info'
	ProvisioningInfo *ProvisioningInfo ``
	// If the connection is read only.
	// Wire name: 'read_only'
	ReadOnly bool ``

	// Wire name: 'securable_type'
	SecurableType SecurableType ``
	// Time at which this connection was updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// Username of user who last modified connection.
	// Wire name: 'updated_by'
	UpdatedBy string ``
	// URL of the remote data source, extracted from options.
	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ConnectionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ConnectionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ConnectionInfoToPb(st *ConnectionInfo) (*catalogpb.ConnectionInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ConnectionInfoPb{}
	pb.Comment = st.Comment
	pb.ConnectionId = st.ConnectionId
	connectionTypePb, err := ConnectionTypeToPb(&st.ConnectionType)
	if err != nil {
		return nil, err
	}
	if connectionTypePb != nil {
		pb.ConnectionType = *connectionTypePb
	}
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	credentialTypePb, err := CredentialTypeToPb(&st.CredentialType)
	if err != nil {
		return nil, err
	}
	if credentialTypePb != nil {
		pb.CredentialType = *credentialTypePb
	}
	environmentSettingsPb, err := EnvironmentSettingsToPb(st.EnvironmentSettings)
	if err != nil {
		return nil, err
	}
	if environmentSettingsPb != nil {
		pb.EnvironmentSettings = environmentSettingsPb
	}
	pb.FullName = st.FullName
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Options = st.Options
	pb.Owner = st.Owner
	pb.Properties = st.Properties
	provisioningInfoPb, err := ProvisioningInfoToPb(st.ProvisioningInfo)
	if err != nil {
		return nil, err
	}
	if provisioningInfoPb != nil {
		pb.ProvisioningInfo = provisioningInfoPb
	}
	pb.ReadOnly = st.ReadOnly
	securableTypePb, err := SecurableTypeToPb(&st.SecurableType)
	if err != nil {
		return nil, err
	}
	if securableTypePb != nil {
		pb.SecurableType = *securableTypePb
	}
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ConnectionInfoFromPb(pb *catalogpb.ConnectionInfoPb) (*ConnectionInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ConnectionInfo{}
	st.Comment = pb.Comment
	st.ConnectionId = pb.ConnectionId
	connectionTypeField, err := ConnectionTypeFromPb(&pb.ConnectionType)
	if err != nil {
		return nil, err
	}
	if connectionTypeField != nil {
		st.ConnectionType = *connectionTypeField
	}
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	credentialTypeField, err := CredentialTypeFromPb(&pb.CredentialType)
	if err != nil {
		return nil, err
	}
	if credentialTypeField != nil {
		st.CredentialType = *credentialTypeField
	}
	environmentSettingsField, err := EnvironmentSettingsFromPb(pb.EnvironmentSettings)
	if err != nil {
		return nil, err
	}
	if environmentSettingsField != nil {
		st.EnvironmentSettings = environmentSettingsField
	}
	st.FullName = pb.FullName
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Options = pb.Options
	st.Owner = pb.Owner
	st.Properties = pb.Properties
	provisioningInfoField, err := ProvisioningInfoFromPb(pb.ProvisioningInfo)
	if err != nil {
		return nil, err
	}
	if provisioningInfoField != nil {
		st.ProvisioningInfo = provisioningInfoField
	}
	st.ReadOnly = pb.ReadOnly
	securableTypeField, err := SecurableTypeFromPb(&pb.SecurableType)
	if err != nil {
		return nil, err
	}
	if securableTypeField != nil {
		st.SecurableType = *securableTypeField
	}
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Next Id: 37
type ConnectionType string

const ConnectionTypeBigquery ConnectionType = `BIGQUERY`

const ConnectionTypeDatabricks ConnectionType = `DATABRICKS`

const ConnectionTypeGa4RawData ConnectionType = `GA4_RAW_DATA`

const ConnectionTypeGlue ConnectionType = `GLUE`

const ConnectionTypeHiveMetastore ConnectionType = `HIVE_METASTORE`

const ConnectionTypeHttp ConnectionType = `HTTP`

const ConnectionTypeMysql ConnectionType = `MYSQL`

const ConnectionTypeOracle ConnectionType = `ORACLE`

const ConnectionTypePostgresql ConnectionType = `POSTGRESQL`

const ConnectionTypePowerBi ConnectionType = `POWER_BI`

const ConnectionTypeRedshift ConnectionType = `REDSHIFT`

const ConnectionTypeSalesforce ConnectionType = `SALESFORCE`

const ConnectionTypeSalesforceDataCloud ConnectionType = `SALESFORCE_DATA_CLOUD`

const ConnectionTypeServicenow ConnectionType = `SERVICENOW`

const ConnectionTypeSnowflake ConnectionType = `SNOWFLAKE`

const ConnectionTypeSqldw ConnectionType = `SQLDW`

const ConnectionTypeSqlserver ConnectionType = `SQLSERVER`

const ConnectionTypeTeradata ConnectionType = `TERADATA`

const ConnectionTypeUnknownConnectionType ConnectionType = `UNKNOWN_CONNECTION_TYPE`

const ConnectionTypeWorkdayRaas ConnectionType = `WORKDAY_RAAS`

// String representation for [fmt.Print]
func (f *ConnectionType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ConnectionType) Set(v string) error {
	switch v {
	case `BIGQUERY`, `DATABRICKS`, `GA4_RAW_DATA`, `GLUE`, `HIVE_METASTORE`, `HTTP`, `MYSQL`, `ORACLE`, `POSTGRESQL`, `POWER_BI`, `REDSHIFT`, `SALESFORCE`, `SALESFORCE_DATA_CLOUD`, `SERVICENOW`, `SNOWFLAKE`, `SQLDW`, `SQLSERVER`, `TERADATA`, `UNKNOWN_CONNECTION_TYPE`, `WORKDAY_RAAS`:
		*f = ConnectionType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BIGQUERY", "DATABRICKS", "GA4_RAW_DATA", "GLUE", "HIVE_METASTORE", "HTTP", "MYSQL", "ORACLE", "POSTGRESQL", "POWER_BI", "REDSHIFT", "SALESFORCE", "SALESFORCE_DATA_CLOUD", "SERVICENOW", "SNOWFLAKE", "SQLDW", "SQLSERVER", "TERADATA", "UNKNOWN_CONNECTION_TYPE", "WORKDAY_RAAS"`, v)
	}
}

// Values returns all possible values for ConnectionType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ConnectionType) Values() []ConnectionType {
	return []ConnectionType{
		ConnectionTypeBigquery,
		ConnectionTypeDatabricks,
		ConnectionTypeGa4RawData,
		ConnectionTypeGlue,
		ConnectionTypeHiveMetastore,
		ConnectionTypeHttp,
		ConnectionTypeMysql,
		ConnectionTypeOracle,
		ConnectionTypePostgresql,
		ConnectionTypePowerBi,
		ConnectionTypeRedshift,
		ConnectionTypeSalesforce,
		ConnectionTypeSalesforceDataCloud,
		ConnectionTypeServicenow,
		ConnectionTypeSnowflake,
		ConnectionTypeSqldw,
		ConnectionTypeSqlserver,
		ConnectionTypeTeradata,
		ConnectionTypeUnknownConnectionType,
		ConnectionTypeWorkdayRaas,
	}
}

// Type always returns ConnectionType to satisfy [pflag.Value] interface
func (f *ConnectionType) Type() string {
	return "ConnectionType"
}

func ConnectionTypeToPb(st *ConnectionType) (*catalogpb.ConnectionTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.ConnectionTypePb(*st)
	return &pb, nil
}

func ConnectionTypeFromPb(pb *catalogpb.ConnectionTypePb) (*ConnectionType, error) {
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
	InitialPipelineSyncProgress *PipelineProgress ``
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	// Wire name: 'last_processed_commit_version'
	LastProcessedCommitVersion int64 ``
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	// Wire name: 'timestamp'
	Timestamp       *time.Time ``
	ForceSendFields []string   `tf:"-"`
}

func (s *ContinuousUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ContinuousUpdateStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ContinuousUpdateStatusToPb(st *ContinuousUpdateStatus) (*catalogpb.ContinuousUpdateStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ContinuousUpdateStatusPb{}
	initialPipelineSyncProgressPb, err := PipelineProgressToPb(st.InitialPipelineSyncProgress)
	if err != nil {
		return nil, err
	}
	if initialPipelineSyncProgressPb != nil {
		pb.InitialPipelineSyncProgress = initialPipelineSyncProgressPb
	}
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion
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

func ContinuousUpdateStatusFromPb(pb *catalogpb.ContinuousUpdateStatusPb) (*ContinuousUpdateStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ContinuousUpdateStatus{}
	initialPipelineSyncProgressField, err := PipelineProgressFromPb(pb.InitialPipelineSyncProgress)
	if err != nil {
		return nil, err
	}
	if initialPipelineSyncProgressField != nil {
		st.InitialPipelineSyncProgress = initialPipelineSyncProgressField
	}
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
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

type CreateCatalog struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// The name of the connection to an external data source.
	// Wire name: 'connection_name'
	ConnectionName string ``
	// Name of catalog.
	// Wire name: 'name'
	Name string ``
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string ``
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string ``
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	// Wire name: 'provider_name'
	ProviderName string ``
	// The name of the share under the share provider.
	// Wire name: 'share_name'
	ShareName string ``
	// Storage root URL for managed tables within catalog.
	// Wire name: 'storage_root'
	StorageRoot     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateCatalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCatalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateCatalogToPb(st *CreateCatalog) (*catalogpb.CreateCatalogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateCatalogPb{}
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

func CreateCatalogFromPb(pb *catalogpb.CreateCatalogPb) (*CreateCatalog, error) {
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

type CreateConnection struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// The type of connection.
	// Wire name: 'connection_type'
	ConnectionType ConnectionType ``
	// [Create,Update:OPT] Connection environment settings as
	// EnvironmentSettings object.
	// Wire name: 'environment_settings'
	EnvironmentSettings *EnvironmentSettings ``
	// Name of the connection.
	// Wire name: 'name'
	Name string ``
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string ``
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string ``
	// If the connection is read only.
	// Wire name: 'read_only'
	ReadOnly        bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateConnection) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateConnection) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateConnectionToPb(st *CreateConnection) (*catalogpb.CreateConnectionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateConnectionPb{}
	pb.Comment = st.Comment
	connectionTypePb, err := ConnectionTypeToPb(&st.ConnectionType)
	if err != nil {
		return nil, err
	}
	if connectionTypePb != nil {
		pb.ConnectionType = *connectionTypePb
	}
	environmentSettingsPb, err := EnvironmentSettingsToPb(st.EnvironmentSettings)
	if err != nil {
		return nil, err
	}
	if environmentSettingsPb != nil {
		pb.EnvironmentSettings = environmentSettingsPb
	}
	pb.Name = st.Name
	pb.Options = st.Options
	pb.Properties = st.Properties
	pb.ReadOnly = st.ReadOnly

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateConnectionFromPb(pb *catalogpb.CreateConnectionPb) (*CreateConnection, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateConnection{}
	st.Comment = pb.Comment
	connectionTypeField, err := ConnectionTypeFromPb(&pb.ConnectionType)
	if err != nil {
		return nil, err
	}
	if connectionTypeField != nil {
		st.ConnectionType = *connectionTypeField
	}
	environmentSettingsField, err := EnvironmentSettingsFromPb(pb.EnvironmentSettings)
	if err != nil {
		return nil, err
	}
	if environmentSettingsField != nil {
		st.EnvironmentSettings = environmentSettingsField
	}
	st.Name = pb.Name
	st.Options = pb.Options
	st.Properties = pb.Properties
	st.ReadOnly = pb.ReadOnly

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateCredentialRequest struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRole ``
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentity ``
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal ``
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string ``
	// The Databricks managed GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount ``
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	// Wire name: 'name'
	Name string ``
	// Indicates the purpose of the credential.
	// Wire name: 'purpose'
	Purpose CredentialPurpose ``
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	// Wire name: 'read_only'
	ReadOnly bool ``
	// Optional. Supplying true to this argument skips validation of the created
	// set of credentials.
	// Wire name: 'skip_validation'
	SkipValidation  bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateCredentialRequestToPb(st *CreateCredentialRequest) (*catalogpb.CreateCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateCredentialRequestPb{}
	awsIamRolePb, err := AwsIamRoleToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}
	azureManagedIdentityPb, err := AzureManagedIdentityToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}
	azureServicePrincipalPb, err := AzureServicePrincipalToPb(st.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalPb != nil {
		pb.AzureServicePrincipal = azureServicePrincipalPb
	}
	pb.Comment = st.Comment
	databricksGcpServiceAccountPb, err := DatabricksGcpServiceAccountToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}
	pb.Name = st.Name
	purposePb, err := CredentialPurposeToPb(&st.Purpose)
	if err != nil {
		return nil, err
	}
	if purposePb != nil {
		pb.Purpose = *purposePb
	}
	pb.ReadOnly = st.ReadOnly
	pb.SkipValidation = st.SkipValidation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateCredentialRequestFromPb(pb *catalogpb.CreateCredentialRequestPb) (*CreateCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialRequest{}
	awsIamRoleField, err := AwsIamRoleFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := AzureManagedIdentityFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	azureServicePrincipalField, err := AzureServicePrincipalFromPb(pb.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalField != nil {
		st.AzureServicePrincipal = azureServicePrincipalField
	}
	st.Comment = pb.Comment
	databricksGcpServiceAccountField, err := DatabricksGcpServiceAccountFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	st.Name = pb.Name
	purposeField, err := CredentialPurposeFromPb(&pb.Purpose)
	if err != nil {
		return nil, err
	}
	if purposeField != nil {
		st.Purpose = *purposeField
	}
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateExternalLineageRelationshipRequest struct {

	// Wire name: 'external_lineage_relationship'
	ExternalLineageRelationship CreateRequestExternalLineage ``
}

func CreateExternalLineageRelationshipRequestToPb(st *CreateExternalLineageRelationshipRequest) (*catalogpb.CreateExternalLineageRelationshipRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateExternalLineageRelationshipRequestPb{}
	externalLineageRelationshipPb, err := CreateRequestExternalLineageToPb(&st.ExternalLineageRelationship)
	if err != nil {
		return nil, err
	}
	if externalLineageRelationshipPb != nil {
		pb.ExternalLineageRelationship = *externalLineageRelationshipPb
	}

	return pb, nil
}

func CreateExternalLineageRelationshipRequestFromPb(pb *catalogpb.CreateExternalLineageRelationshipRequestPb) (*CreateExternalLineageRelationshipRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExternalLineageRelationshipRequest{}
	externalLineageRelationshipField, err := CreateRequestExternalLineageFromPb(&pb.ExternalLineageRelationship)
	if err != nil {
		return nil, err
	}
	if externalLineageRelationshipField != nil {
		st.ExternalLineageRelationship = *externalLineageRelationshipField
	}

	return st, nil
}

type CreateExternalLocation struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// Name of the storage credential used with this location.
	// Wire name: 'credential_name'
	CredentialName string ``
	// Whether to enable file events on this external location.
	// Wire name: 'enable_file_events'
	EnableFileEvents bool ``

	// Wire name: 'encryption_details'
	EncryptionDetails *EncryptionDetails ``
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	// Wire name: 'fallback'
	Fallback bool ``
	// File event queue settings.
	// Wire name: 'file_event_queue'
	FileEventQueue *FileEventQueue ``
	// Name of the external location.
	// Wire name: 'name'
	Name string ``
	// Indicates whether the external location is read-only.
	// Wire name: 'read_only'
	ReadOnly bool ``
	// Skips validation of the storage credential associated with the external
	// location.
	// Wire name: 'skip_validation'
	SkipValidation bool ``
	// Path URL of the external location.
	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateExternalLocation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateExternalLocation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateExternalLocationToPb(st *CreateExternalLocation) (*catalogpb.CreateExternalLocationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateExternalLocationPb{}
	pb.Comment = st.Comment
	pb.CredentialName = st.CredentialName
	pb.EnableFileEvents = st.EnableFileEvents
	encryptionDetailsPb, err := EncryptionDetailsToPb(st.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsPb != nil {
		pb.EncryptionDetails = encryptionDetailsPb
	}
	pb.Fallback = st.Fallback
	fileEventQueuePb, err := FileEventQueueToPb(st.FileEventQueue)
	if err != nil {
		return nil, err
	}
	if fileEventQueuePb != nil {
		pb.FileEventQueue = fileEventQueuePb
	}
	pb.Name = st.Name
	pb.ReadOnly = st.ReadOnly
	pb.SkipValidation = st.SkipValidation
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateExternalLocationFromPb(pb *catalogpb.CreateExternalLocationPb) (*CreateExternalLocation, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExternalLocation{}
	st.Comment = pb.Comment
	st.CredentialName = pb.CredentialName
	st.EnableFileEvents = pb.EnableFileEvents
	encryptionDetailsField, err := EncryptionDetailsFromPb(pb.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsField != nil {
		st.EncryptionDetails = encryptionDetailsField
	}
	st.Fallback = pb.Fallback
	fileEventQueueField, err := FileEventQueueFromPb(pb.FileEventQueue)
	if err != nil {
		return nil, err
	}
	if fileEventQueueField != nil {
		st.FileEventQueue = fileEventQueueField
	}
	st.Name = pb.Name
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateExternalMetadataRequest struct {

	// Wire name: 'external_metadata'
	ExternalMetadata ExternalMetadata ``
}

func CreateExternalMetadataRequestToPb(st *CreateExternalMetadataRequest) (*catalogpb.CreateExternalMetadataRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateExternalMetadataRequestPb{}
	externalMetadataPb, err := ExternalMetadataToPb(&st.ExternalMetadata)
	if err != nil {
		return nil, err
	}
	if externalMetadataPb != nil {
		pb.ExternalMetadata = *externalMetadataPb
	}

	return pb, nil
}

func CreateExternalMetadataRequestFromPb(pb *catalogpb.CreateExternalMetadataRequestPb) (*CreateExternalMetadataRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExternalMetadataRequest{}
	externalMetadataField, err := ExternalMetadataFromPb(&pb.ExternalMetadata)
	if err != nil {
		return nil, err
	}
	if externalMetadataField != nil {
		st.ExternalMetadata = *externalMetadataField
	}

	return st, nil
}

type CreateFunction struct {
	// Name of parent catalog.
	// Wire name: 'catalog_name'
	CatalogName string ``
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// Scalar function return data type.
	// Wire name: 'data_type'
	DataType ColumnTypeName ``
	// External function language.
	// Wire name: 'external_language'
	ExternalLanguage string ``
	// External function name.
	// Wire name: 'external_name'
	ExternalName string ``
	// Pretty printed function data type.
	// Wire name: 'full_data_type'
	FullDataType string ``

	// Wire name: 'input_params'
	InputParams FunctionParameterInfos ``
	// Whether the function is deterministic.
	// Wire name: 'is_deterministic'
	IsDeterministic bool ``
	// Function null call.
	// Wire name: 'is_null_call'
	IsNullCall bool ``
	// Name of function, relative to parent schema.
	// Wire name: 'name'
	Name string ``
	// Function parameter style. **S** is the value for SQL.
	// Wire name: 'parameter_style'
	ParameterStyle CreateFunctionParameterStyle ``
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	// Wire name: 'properties'
	Properties string ``
	// Table function return parameters.
	// Wire name: 'return_params'
	ReturnParams *FunctionParameterInfos ``
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	// Wire name: 'routine_body'
	RoutineBody CreateFunctionRoutineBody ``
	// Function body.
	// Wire name: 'routine_definition'
	RoutineDefinition string ``
	// Function dependencies.
	// Wire name: 'routine_dependencies'
	RoutineDependencies *DependencyList ``
	// Name of parent schema relative to its parent catalog.
	// Wire name: 'schema_name'
	SchemaName string ``
	// Function security type.
	// Wire name: 'security_type'
	SecurityType CreateFunctionSecurityType ``
	// Specific name of the function; Reserved for future use.
	// Wire name: 'specific_name'
	SpecificName string ``
	// Function SQL data access.
	// Wire name: 'sql_data_access'
	SqlDataAccess CreateFunctionSqlDataAccess ``
	// List of schemes whose objects can be referenced without qualification.
	// Wire name: 'sql_path'
	SqlPath         string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateFunctionToPb(st *CreateFunction) (*catalogpb.CreateFunctionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateFunctionPb{}
	pb.CatalogName = st.CatalogName
	pb.Comment = st.Comment
	dataTypePb, err := ColumnTypeNameToPb(&st.DataType)
	if err != nil {
		return nil, err
	}
	if dataTypePb != nil {
		pb.DataType = *dataTypePb
	}
	pb.ExternalLanguage = st.ExternalLanguage
	pb.ExternalName = st.ExternalName
	pb.FullDataType = st.FullDataType
	inputParamsPb, err := FunctionParameterInfosToPb(&st.InputParams)
	if err != nil {
		return nil, err
	}
	if inputParamsPb != nil {
		pb.InputParams = *inputParamsPb
	}
	pb.IsDeterministic = st.IsDeterministic
	pb.IsNullCall = st.IsNullCall
	pb.Name = st.Name
	parameterStylePb, err := CreateFunctionParameterStyleToPb(&st.ParameterStyle)
	if err != nil {
		return nil, err
	}
	if parameterStylePb != nil {
		pb.ParameterStyle = *parameterStylePb
	}
	pb.Properties = st.Properties
	returnParamsPb, err := FunctionParameterInfosToPb(st.ReturnParams)
	if err != nil {
		return nil, err
	}
	if returnParamsPb != nil {
		pb.ReturnParams = returnParamsPb
	}
	routineBodyPb, err := CreateFunctionRoutineBodyToPb(&st.RoutineBody)
	if err != nil {
		return nil, err
	}
	if routineBodyPb != nil {
		pb.RoutineBody = *routineBodyPb
	}
	pb.RoutineDefinition = st.RoutineDefinition
	routineDependenciesPb, err := DependencyListToPb(st.RoutineDependencies)
	if err != nil {
		return nil, err
	}
	if routineDependenciesPb != nil {
		pb.RoutineDependencies = routineDependenciesPb
	}
	pb.SchemaName = st.SchemaName
	securityTypePb, err := CreateFunctionSecurityTypeToPb(&st.SecurityType)
	if err != nil {
		return nil, err
	}
	if securityTypePb != nil {
		pb.SecurityType = *securityTypePb
	}
	pb.SpecificName = st.SpecificName
	sqlDataAccessPb, err := CreateFunctionSqlDataAccessToPb(&st.SqlDataAccess)
	if err != nil {
		return nil, err
	}
	if sqlDataAccessPb != nil {
		pb.SqlDataAccess = *sqlDataAccessPb
	}
	pb.SqlPath = st.SqlPath

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateFunctionFromPb(pb *catalogpb.CreateFunctionPb) (*CreateFunction, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateFunction{}
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	dataTypeField, err := ColumnTypeNameFromPb(&pb.DataType)
	if err != nil {
		return nil, err
	}
	if dataTypeField != nil {
		st.DataType = *dataTypeField
	}
	st.ExternalLanguage = pb.ExternalLanguage
	st.ExternalName = pb.ExternalName
	st.FullDataType = pb.FullDataType
	inputParamsField, err := FunctionParameterInfosFromPb(&pb.InputParams)
	if err != nil {
		return nil, err
	}
	if inputParamsField != nil {
		st.InputParams = *inputParamsField
	}
	st.IsDeterministic = pb.IsDeterministic
	st.IsNullCall = pb.IsNullCall
	st.Name = pb.Name
	parameterStyleField, err := CreateFunctionParameterStyleFromPb(&pb.ParameterStyle)
	if err != nil {
		return nil, err
	}
	if parameterStyleField != nil {
		st.ParameterStyle = *parameterStyleField
	}
	st.Properties = pb.Properties
	returnParamsField, err := FunctionParameterInfosFromPb(pb.ReturnParams)
	if err != nil {
		return nil, err
	}
	if returnParamsField != nil {
		st.ReturnParams = returnParamsField
	}
	routineBodyField, err := CreateFunctionRoutineBodyFromPb(&pb.RoutineBody)
	if err != nil {
		return nil, err
	}
	if routineBodyField != nil {
		st.RoutineBody = *routineBodyField
	}
	st.RoutineDefinition = pb.RoutineDefinition
	routineDependenciesField, err := DependencyListFromPb(pb.RoutineDependencies)
	if err != nil {
		return nil, err
	}
	if routineDependenciesField != nil {
		st.RoutineDependencies = routineDependenciesField
	}
	st.SchemaName = pb.SchemaName
	securityTypeField, err := CreateFunctionSecurityTypeFromPb(&pb.SecurityType)
	if err != nil {
		return nil, err
	}
	if securityTypeField != nil {
		st.SecurityType = *securityTypeField
	}
	st.SpecificName = pb.SpecificName
	sqlDataAccessField, err := CreateFunctionSqlDataAccessFromPb(&pb.SqlDataAccess)
	if err != nil {
		return nil, err
	}
	if sqlDataAccessField != nil {
		st.SqlDataAccess = *sqlDataAccessField
	}
	st.SqlPath = pb.SqlPath

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Function parameter style. **S** is the value for SQL.
type CreateFunctionParameterStyle string

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

// Values returns all possible values for CreateFunctionParameterStyle.
//
// There is no guarantee on the order of the values in the slice.
func (f *CreateFunctionParameterStyle) Values() []CreateFunctionParameterStyle {
	return []CreateFunctionParameterStyle{
		CreateFunctionParameterStyleS,
	}
}

// Type always returns CreateFunctionParameterStyle to satisfy [pflag.Value] interface
func (f *CreateFunctionParameterStyle) Type() string {
	return "CreateFunctionParameterStyle"
}

func CreateFunctionParameterStyleToPb(st *CreateFunctionParameterStyle) (*catalogpb.CreateFunctionParameterStylePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.CreateFunctionParameterStylePb(*st)
	return &pb, nil
}

func CreateFunctionParameterStyleFromPb(pb *catalogpb.CreateFunctionParameterStylePb) (*CreateFunctionParameterStyle, error) {
	if pb == nil {
		return nil, nil
	}
	st := CreateFunctionParameterStyle(*pb)
	return &st, nil
}

type CreateFunctionRequest struct {
	// Partial __FunctionInfo__ specifying the function to be created.
	// Wire name: 'function_info'
	FunctionInfo CreateFunction ``
}

func CreateFunctionRequestToPb(st *CreateFunctionRequest) (*catalogpb.CreateFunctionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateFunctionRequestPb{}
	functionInfoPb, err := CreateFunctionToPb(&st.FunctionInfo)
	if err != nil {
		return nil, err
	}
	if functionInfoPb != nil {
		pb.FunctionInfo = *functionInfoPb
	}

	return pb, nil
}

func CreateFunctionRequestFromPb(pb *catalogpb.CreateFunctionRequestPb) (*CreateFunctionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateFunctionRequest{}
	functionInfoField, err := CreateFunctionFromPb(&pb.FunctionInfo)
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

// Values returns all possible values for CreateFunctionRoutineBody.
//
// There is no guarantee on the order of the values in the slice.
func (f *CreateFunctionRoutineBody) Values() []CreateFunctionRoutineBody {
	return []CreateFunctionRoutineBody{
		CreateFunctionRoutineBodyExternal,
		CreateFunctionRoutineBodySql,
	}
}

// Type always returns CreateFunctionRoutineBody to satisfy [pflag.Value] interface
func (f *CreateFunctionRoutineBody) Type() string {
	return "CreateFunctionRoutineBody"
}

func CreateFunctionRoutineBodyToPb(st *CreateFunctionRoutineBody) (*catalogpb.CreateFunctionRoutineBodyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.CreateFunctionRoutineBodyPb(*st)
	return &pb, nil
}

func CreateFunctionRoutineBodyFromPb(pb *catalogpb.CreateFunctionRoutineBodyPb) (*CreateFunctionRoutineBody, error) {
	if pb == nil {
		return nil, nil
	}
	st := CreateFunctionRoutineBody(*pb)
	return &st, nil
}

// The security type of the function.
type CreateFunctionSecurityType string

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

// Values returns all possible values for CreateFunctionSecurityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *CreateFunctionSecurityType) Values() []CreateFunctionSecurityType {
	return []CreateFunctionSecurityType{
		CreateFunctionSecurityTypeDefiner,
	}
}

// Type always returns CreateFunctionSecurityType to satisfy [pflag.Value] interface
func (f *CreateFunctionSecurityType) Type() string {
	return "CreateFunctionSecurityType"
}

func CreateFunctionSecurityTypeToPb(st *CreateFunctionSecurityType) (*catalogpb.CreateFunctionSecurityTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.CreateFunctionSecurityTypePb(*st)
	return &pb, nil
}

func CreateFunctionSecurityTypeFromPb(pb *catalogpb.CreateFunctionSecurityTypePb) (*CreateFunctionSecurityType, error) {
	if pb == nil {
		return nil, nil
	}
	st := CreateFunctionSecurityType(*pb)
	return &st, nil
}

// Function SQL data access.
type CreateFunctionSqlDataAccess string

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

// Values returns all possible values for CreateFunctionSqlDataAccess.
//
// There is no guarantee on the order of the values in the slice.
func (f *CreateFunctionSqlDataAccess) Values() []CreateFunctionSqlDataAccess {
	return []CreateFunctionSqlDataAccess{
		CreateFunctionSqlDataAccessContainsSql,
		CreateFunctionSqlDataAccessNoSql,
		CreateFunctionSqlDataAccessReadsSqlData,
	}
}

// Type always returns CreateFunctionSqlDataAccess to satisfy [pflag.Value] interface
func (f *CreateFunctionSqlDataAccess) Type() string {
	return "CreateFunctionSqlDataAccess"
}

func CreateFunctionSqlDataAccessToPb(st *CreateFunctionSqlDataAccess) (*catalogpb.CreateFunctionSqlDataAccessPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.CreateFunctionSqlDataAccessPb(*st)
	return &pb, nil
}

func CreateFunctionSqlDataAccessFromPb(pb *catalogpb.CreateFunctionSqlDataAccessPb) (*CreateFunctionSqlDataAccess, error) {
	if pb == nil {
		return nil, nil
	}
	st := CreateFunctionSqlDataAccess(*pb)
	return &st, nil
}

type CreateMetastore struct {
	// The user-specified name of the metastore.
	// Wire name: 'name'
	Name string ``
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	// Wire name: 'region'
	Region string ``
	// The storage root URL for metastore
	// Wire name: 'storage_root'
	StorageRoot     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateMetastore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateMetastore) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateMetastoreToPb(st *CreateMetastore) (*catalogpb.CreateMetastorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateMetastorePb{}
	pb.Name = st.Name
	pb.Region = st.Region
	pb.StorageRoot = st.StorageRoot

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateMetastoreFromPb(pb *catalogpb.CreateMetastorePb) (*CreateMetastore, error) {
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

type CreateMetastoreAssignment struct {
	// The name of the default catalog in the metastore. This field is
	// deprecated. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	// Wire name: 'default_catalog_name'
	DefaultCatalogName string ``
	// The unique ID of the metastore.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// A workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func CreateMetastoreAssignmentToPb(st *CreateMetastoreAssignment) (*catalogpb.CreateMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateMetastoreAssignmentPb{}
	pb.DefaultCatalogName = st.DefaultCatalogName
	pb.MetastoreId = st.MetastoreId
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func CreateMetastoreAssignmentFromPb(pb *catalogpb.CreateMetastoreAssignmentPb) (*CreateMetastoreAssignment, error) {
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
	// [Create:REQ Update:IGN] Field for specifying the absolute path to a
	// custom directory to store data-monitoring assets. Normally prepopulated
	// to a default user location via UI and Python APIs.
	// Wire name: 'assets_dir'
	AssetsDir string ``
	// [Create:OPT Update:OPT] Baseline table name. Baseline data is used to
	// compute drift from the data in the monitored `table_name`. The baseline
	// table and the monitored table shall have the same schema.
	// Wire name: 'baseline_table_name'
	BaselineTableName string ``
	// [Create:OPT Update:OPT] Custom metrics.
	// Wire name: 'custom_metrics'
	CustomMetrics []MonitorMetric ``
	// [Create:OPT Update:OPT] Data classification related config.
	// Wire name: 'data_classification_config'
	DataClassificationConfig *MonitorDataClassificationConfig ``

	// Wire name: 'inference_log'
	InferenceLog *MonitorInferenceLog ``
	// [Create:ERR Update:IGN] The latest error message for a monitor failure.
	// Wire name: 'latest_monitor_failure_msg'
	LatestMonitorFailureMsg string ``
	// [Create:OPT Update:OPT] Field for specifying notification settings.
	// Wire name: 'notifications'
	Notifications *MonitorNotifications ``
	// [Create:REQ Update:REQ] Schema where output tables are created. Needs to
	// be in 2-level format {catalog}.{schema}
	// Wire name: 'output_schema_name'
	OutputSchemaName string ``
	// [Create:OPT Update:OPT] The monitor schedule.
	// Wire name: 'schedule'
	Schedule *MonitorCronSchedule ``
	// Whether to skip creating a default dashboard summarizing data quality
	// metrics.
	// Wire name: 'skip_builtin_dashboard'
	SkipBuiltinDashboard bool ``
	// [Create:OPT Update:OPT] List of column expressions to slice data with for
	// targeted analysis. The data is grouped by each expression independently,
	// resulting in a separate slice for each predicate and its complements. For
	// example `slicing_exprs=[“col_1”, “col_2 > 10”]` will generate the
	// following slices: two slices for `col_2 > 10` (True and False), and one
	// slice per unique value in `col1`. For high-cardinality columns, only the
	// top 100 unique values by frequency will generate slices.
	// Wire name: 'slicing_exprs'
	SlicingExprs []string ``
	// Configuration for monitoring snapshot tables.
	// Wire name: 'snapshot'
	Snapshot *MonitorSnapshot ``
	// UC table name in format `catalog.schema.table_name`. This field
	// corresponds to the {full_table_name_arg} arg in the endpoint path.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
	// Configuration for monitoring time series tables.
	// Wire name: 'time_series'
	TimeSeries *MonitorTimeSeries ``
	// Optional argument to specify the warehouse for dashboard creation. If not
	// specified, the first running warehouse will be used.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateMonitor) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateMonitor) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateMonitorToPb(st *CreateMonitor) (*catalogpb.CreateMonitorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateMonitorPb{}
	pb.AssetsDir = st.AssetsDir
	pb.BaselineTableName = st.BaselineTableName

	var customMetricsPb []catalogpb.MonitorMetricPb
	for _, item := range st.CustomMetrics {
		itemPb, err := MonitorMetricToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customMetricsPb = append(customMetricsPb, *itemPb)
		}
	}
	pb.CustomMetrics = customMetricsPb
	dataClassificationConfigPb, err := MonitorDataClassificationConfigToPb(st.DataClassificationConfig)
	if err != nil {
		return nil, err
	}
	if dataClassificationConfigPb != nil {
		pb.DataClassificationConfig = dataClassificationConfigPb
	}
	inferenceLogPb, err := MonitorInferenceLogToPb(st.InferenceLog)
	if err != nil {
		return nil, err
	}
	if inferenceLogPb != nil {
		pb.InferenceLog = inferenceLogPb
	}
	pb.LatestMonitorFailureMsg = st.LatestMonitorFailureMsg
	notificationsPb, err := MonitorNotificationsToPb(st.Notifications)
	if err != nil {
		return nil, err
	}
	if notificationsPb != nil {
		pb.Notifications = notificationsPb
	}
	pb.OutputSchemaName = st.OutputSchemaName
	schedulePb, err := MonitorCronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}
	pb.SkipBuiltinDashboard = st.SkipBuiltinDashboard
	pb.SlicingExprs = st.SlicingExprs
	snapshotPb, err := MonitorSnapshotToPb(st.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotPb != nil {
		pb.Snapshot = snapshotPb
	}
	pb.TableName = st.TableName
	timeSeriesPb, err := MonitorTimeSeriesToPb(st.TimeSeries)
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

func CreateMonitorFromPb(pb *catalogpb.CreateMonitorPb) (*CreateMonitor, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateMonitor{}
	st.AssetsDir = pb.AssetsDir
	st.BaselineTableName = pb.BaselineTableName

	var customMetricsField []MonitorMetric
	for _, itemPb := range pb.CustomMetrics {
		item, err := MonitorMetricFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customMetricsField = append(customMetricsField, *item)
		}
	}
	st.CustomMetrics = customMetricsField
	dataClassificationConfigField, err := MonitorDataClassificationConfigFromPb(pb.DataClassificationConfig)
	if err != nil {
		return nil, err
	}
	if dataClassificationConfigField != nil {
		st.DataClassificationConfig = dataClassificationConfigField
	}
	inferenceLogField, err := MonitorInferenceLogFromPb(pb.InferenceLog)
	if err != nil {
		return nil, err
	}
	if inferenceLogField != nil {
		st.InferenceLog = inferenceLogField
	}
	st.LatestMonitorFailureMsg = pb.LatestMonitorFailureMsg
	notificationsField, err := MonitorNotificationsFromPb(pb.Notifications)
	if err != nil {
		return nil, err
	}
	if notificationsField != nil {
		st.Notifications = notificationsField
	}
	st.OutputSchemaName = pb.OutputSchemaName
	scheduleField, err := MonitorCronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	st.SkipBuiltinDashboard = pb.SkipBuiltinDashboard
	st.SlicingExprs = pb.SlicingExprs
	snapshotField, err := MonitorSnapshotFromPb(pb.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotField != nil {
		st.Snapshot = snapshotField
	}
	st.TableName = pb.TableName
	timeSeriesField, err := MonitorTimeSeriesFromPb(pb.TimeSeries)
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

type CreateOnlineTableRequest struct {
	// Specification of the online table to be created.
	// Wire name: 'table'
	Table OnlineTable ``
}

func CreateOnlineTableRequestToPb(st *CreateOnlineTableRequest) (*catalogpb.CreateOnlineTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateOnlineTableRequestPb{}
	tablePb, err := OnlineTableToPb(&st.Table)
	if err != nil {
		return nil, err
	}
	if tablePb != nil {
		pb.Table = *tablePb
	}

	return pb, nil
}

func CreateOnlineTableRequestFromPb(pb *catalogpb.CreateOnlineTableRequestPb) (*CreateOnlineTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateOnlineTableRequest{}
	tableField, err := OnlineTableFromPb(&pb.Table)
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
	CatalogName string ``
	// The comment attached to the registered model
	// Wire name: 'comment'
	Comment string ``
	// The name of the registered model
	// Wire name: 'name'
	Name string ``
	// The name of the schema where the registered model resides
	// Wire name: 'schema_name'
	SchemaName string ``
	// The storage location on the cloud under which model version data files
	// are stored
	// Wire name: 'storage_location'
	StorageLocation string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateRegisteredModelRequestToPb(st *CreateRegisteredModelRequest) (*catalogpb.CreateRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateRegisteredModelRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.SchemaName = st.SchemaName
	pb.StorageLocation = st.StorageLocation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateRegisteredModelRequestFromPb(pb *catalogpb.CreateRegisteredModelRequestPb) (*CreateRegisteredModelRequest, error) {
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

type CreateRequestExternalLineage struct {
	// List of column relationships between source and target objects.
	// Wire name: 'columns'
	Columns []ColumnRelationship ``
	// Unique identifier of the external lineage relationship.
	// Wire name: 'id'
	Id string ``
	// Key-value properties associated with the external lineage relationship.
	// Wire name: 'properties'
	Properties map[string]string ``
	// Source object of the external lineage relationship.
	// Wire name: 'source'
	Source ExternalLineageObject ``
	// Target object of the external lineage relationship.
	// Wire name: 'target'
	Target          ExternalLineageObject ``
	ForceSendFields []string              `tf:"-"`
}

func (s *CreateRequestExternalLineage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRequestExternalLineage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateRequestExternalLineageToPb(st *CreateRequestExternalLineage) (*catalogpb.CreateRequestExternalLineagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateRequestExternalLineagePb{}

	var columnsPb []catalogpb.ColumnRelationshipPb
	for _, item := range st.Columns {
		itemPb, err := ColumnRelationshipToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			columnsPb = append(columnsPb, *itemPb)
		}
	}
	pb.Columns = columnsPb
	pb.Id = st.Id
	pb.Properties = st.Properties
	sourcePb, err := ExternalLineageObjectToPb(&st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = *sourcePb
	}
	targetPb, err := ExternalLineageObjectToPb(&st.Target)
	if err != nil {
		return nil, err
	}
	if targetPb != nil {
		pb.Target = *targetPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateRequestExternalLineageFromPb(pb *catalogpb.CreateRequestExternalLineagePb) (*CreateRequestExternalLineage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRequestExternalLineage{}

	var columnsField []ColumnRelationship
	for _, itemPb := range pb.Columns {
		item, err := ColumnRelationshipFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			columnsField = append(columnsField, *item)
		}
	}
	st.Columns = columnsField
	st.Id = pb.Id
	st.Properties = pb.Properties
	sourceField, err := ExternalLineageObjectFromPb(&pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = *sourceField
	}
	targetField, err := ExternalLineageObjectFromPb(&pb.Target)
	if err != nil {
		return nil, err
	}
	if targetField != nil {
		st.Target = *targetField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateSchema struct {
	// Name of parent catalog.
	// Wire name: 'catalog_name'
	CatalogName string ``
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// Name of schema, relative to parent catalog.
	// Wire name: 'name'
	Name string ``
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string ``
	// Storage root URL for managed tables within schema.
	// Wire name: 'storage_root'
	StorageRoot     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateSchema) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateSchema) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateSchemaToPb(st *CreateSchema) (*catalogpb.CreateSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateSchemaPb{}
	pb.CatalogName = st.CatalogName
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.Properties = st.Properties
	pb.StorageRoot = st.StorageRoot

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateSchemaFromPb(pb *catalogpb.CreateSchemaPb) (*CreateSchema, error) {
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

type CreateStorageCredential struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRoleRequest ``
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentityRequest ``
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal ``
	// The Cloudflare API token configuration.
	// Wire name: 'cloudflare_api_token'
	CloudflareApiToken *CloudflareApiToken ``
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string ``
	// The Databricks managed GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest ``
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	// Wire name: 'name'
	Name string ``
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	// Wire name: 'read_only'
	ReadOnly bool ``
	// Supplying true to this argument skips validation of the created
	// credential.
	// Wire name: 'skip_validation'
	SkipValidation  bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateStorageCredentialToPb(st *CreateStorageCredential) (*catalogpb.CreateStorageCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateStorageCredentialPb{}
	awsIamRolePb, err := AwsIamRoleRequestToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}
	azureManagedIdentityPb, err := AzureManagedIdentityRequestToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}
	azureServicePrincipalPb, err := AzureServicePrincipalToPb(st.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalPb != nil {
		pb.AzureServicePrincipal = azureServicePrincipalPb
	}
	cloudflareApiTokenPb, err := CloudflareApiTokenToPb(st.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenPb != nil {
		pb.CloudflareApiToken = cloudflareApiTokenPb
	}
	pb.Comment = st.Comment
	databricksGcpServiceAccountPb, err := DatabricksGcpServiceAccountRequestToPb(st.DatabricksGcpServiceAccount)
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

func CreateStorageCredentialFromPb(pb *catalogpb.CreateStorageCredentialPb) (*CreateStorageCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateStorageCredential{}
	awsIamRoleField, err := AwsIamRoleRequestFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := AzureManagedIdentityRequestFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	azureServicePrincipalField, err := AzureServicePrincipalFromPb(pb.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalField != nil {
		st.AzureServicePrincipal = azureServicePrincipalField
	}
	cloudflareApiTokenField, err := CloudflareApiTokenFromPb(pb.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenField != nil {
		st.CloudflareApiToken = cloudflareApiTokenField
	}
	st.Comment = pb.Comment
	databricksGcpServiceAccountField, err := DatabricksGcpServiceAccountRequestFromPb(pb.DatabricksGcpServiceAccount)
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

type CreateTableConstraint struct {

	// Wire name: 'constraint'
	Constraint TableConstraint ``
	// The full name of the table referenced by the constraint.
	// Wire name: 'full_name_arg'
	FullNameArg string ``
}

func CreateTableConstraintToPb(st *CreateTableConstraint) (*catalogpb.CreateTableConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateTableConstraintPb{}
	constraintPb, err := TableConstraintToPb(&st.Constraint)
	if err != nil {
		return nil, err
	}
	if constraintPb != nil {
		pb.Constraint = *constraintPb
	}
	pb.FullNameArg = st.FullNameArg

	return pb, nil
}

func CreateTableConstraintFromPb(pb *catalogpb.CreateTableConstraintPb) (*CreateTableConstraint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateTableConstraint{}
	constraintField, err := TableConstraintFromPb(&pb.Constraint)
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
	CatalogName string ``
	// The comment attached to the volume
	// Wire name: 'comment'
	Comment string ``
	// The name of the volume
	// Wire name: 'name'
	Name string ``
	// The name of the schema where the volume is
	// Wire name: 'schema_name'
	SchemaName string ``
	// The storage location on the cloud
	// Wire name: 'storage_location'
	StorageLocation string ``

	// Wire name: 'volume_type'
	VolumeType      VolumeType ``
	ForceSendFields []string   `tf:"-"`
}

func (s *CreateVolumeRequestContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateVolumeRequestContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateVolumeRequestContentToPb(st *CreateVolumeRequestContent) (*catalogpb.CreateVolumeRequestContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CreateVolumeRequestContentPb{}
	pb.CatalogName = st.CatalogName
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.SchemaName = st.SchemaName
	pb.StorageLocation = st.StorageLocation
	volumeTypePb, err := VolumeTypeToPb(&st.VolumeType)
	if err != nil {
		return nil, err
	}
	if volumeTypePb != nil {
		pb.VolumeType = *volumeTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateVolumeRequestContentFromPb(pb *catalogpb.CreateVolumeRequestContentPb) (*CreateVolumeRequestContent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateVolumeRequestContent{}
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.SchemaName = pb.SchemaName
	st.StorageLocation = pb.StorageLocation
	volumeTypeField, err := VolumeTypeFromPb(&pb.VolumeType)
	if err != nil {
		return nil, err
	}
	if volumeTypeField != nil {
		st.VolumeType = *volumeTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// A credential that is dependent on a SQL object.
type CredentialDependency struct {
	// Full name of the dependent credential, in the form of
	// __credential_name__.
	// Wire name: 'credential_name'
	CredentialName  string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CredentialDependency) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CredentialDependency) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CredentialDependencyToPb(st *CredentialDependency) (*catalogpb.CredentialDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CredentialDependencyPb{}
	pb.CredentialName = st.CredentialName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CredentialDependencyFromPb(pb *catalogpb.CredentialDependencyPb) (*CredentialDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CredentialDependency{}
	st.CredentialName = pb.CredentialName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CredentialInfo struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRole ``
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentity ``
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal ``
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string ``
	// Time at which this credential was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of credential creator.
	// Wire name: 'created_by'
	CreatedBy string ``
	// The Databricks managed GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount ``
	// The full name of the credential.
	// Wire name: 'full_name'
	FullName string ``
	// The unique identifier of the credential.
	// Wire name: 'id'
	Id string ``
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode ``
	// Unique identifier of the parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	// Wire name: 'name'
	Name string ``
	// Username of current owner of credential.
	// Wire name: 'owner'
	Owner string ``
	// Indicates the purpose of the credential.
	// Wire name: 'purpose'
	Purpose CredentialPurpose ``
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	// Wire name: 'read_only'
	ReadOnly bool ``
	// Time at which this credential was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// Username of user who last modified the credential.
	// Wire name: 'updated_by'
	UpdatedBy string ``
	// Whether this credential is the current metastore's root storage
	// credential. Only applicable when purpose is **STORAGE**.
	// Wire name: 'used_for_managed_storage'
	UsedForManagedStorage bool     ``
	ForceSendFields       []string `tf:"-"`
}

func (s *CredentialInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CredentialInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CredentialInfoToPb(st *CredentialInfo) (*catalogpb.CredentialInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CredentialInfoPb{}
	awsIamRolePb, err := AwsIamRoleToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}
	azureManagedIdentityPb, err := AzureManagedIdentityToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}
	azureServicePrincipalPb, err := AzureServicePrincipalToPb(st.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalPb != nil {
		pb.AzureServicePrincipal = azureServicePrincipalPb
	}
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	databricksGcpServiceAccountPb, err := DatabricksGcpServiceAccountToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}
	pb.FullName = st.FullName
	pb.Id = st.Id
	isolationModePb, err := IsolationModeToPb(&st.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	purposePb, err := CredentialPurposeToPb(&st.Purpose)
	if err != nil {
		return nil, err
	}
	if purposePb != nil {
		pb.Purpose = *purposePb
	}
	pb.ReadOnly = st.ReadOnly
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy
	pb.UsedForManagedStorage = st.UsedForManagedStorage

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CredentialInfoFromPb(pb *catalogpb.CredentialInfoPb) (*CredentialInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CredentialInfo{}
	awsIamRoleField, err := AwsIamRoleFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := AzureManagedIdentityFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	azureServicePrincipalField, err := AzureServicePrincipalFromPb(pb.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalField != nil {
		st.AzureServicePrincipal = azureServicePrincipalField
	}
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	databricksGcpServiceAccountField, err := DatabricksGcpServiceAccountFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	st.FullName = pb.FullName
	st.Id = pb.Id
	isolationModeField, err := IsolationModeFromPb(&pb.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	purposeField, err := CredentialPurposeFromPb(&pb.Purpose)
	if err != nil {
		return nil, err
	}
	if purposeField != nil {
		st.Purpose = *purposeField
	}
	st.ReadOnly = pb.ReadOnly
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.UsedForManagedStorage = pb.UsedForManagedStorage

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CredentialPurpose string

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

// Values returns all possible values for CredentialPurpose.
//
// There is no guarantee on the order of the values in the slice.
func (f *CredentialPurpose) Values() []CredentialPurpose {
	return []CredentialPurpose{
		CredentialPurposeService,
		CredentialPurposeStorage,
	}
}

// Type always returns CredentialPurpose to satisfy [pflag.Value] interface
func (f *CredentialPurpose) Type() string {
	return "CredentialPurpose"
}

func CredentialPurposeToPb(st *CredentialPurpose) (*catalogpb.CredentialPurposePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.CredentialPurposePb(*st)
	return &pb, nil
}

func CredentialPurposeFromPb(pb *catalogpb.CredentialPurposePb) (*CredentialPurpose, error) {
	if pb == nil {
		return nil, nil
	}
	st := CredentialPurpose(*pb)
	return &st, nil
}

// Next Id: 13
type CredentialType string

const CredentialTypeAnyStaticCredential CredentialType = `ANY_STATIC_CREDENTIAL`

const CredentialTypeBearerToken CredentialType = `BEARER_TOKEN`

const CredentialTypeOauthAccessToken CredentialType = `OAUTH_ACCESS_TOKEN`

const CredentialTypeOauthM2m CredentialType = `OAUTH_M2M`

const CredentialTypeOauthRefreshToken CredentialType = `OAUTH_REFRESH_TOKEN`

const CredentialTypeOauthResourceOwnerPassword CredentialType = `OAUTH_RESOURCE_OWNER_PASSWORD`

const CredentialTypeOauthU2m CredentialType = `OAUTH_U2M`

const CredentialTypeOauthU2mMapping CredentialType = `OAUTH_U2M_MAPPING`

const CredentialTypeOidcToken CredentialType = `OIDC_TOKEN`

const CredentialTypePemPrivateKey CredentialType = `PEM_PRIVATE_KEY`

const CredentialTypeServiceCredential CredentialType = `SERVICE_CREDENTIAL`

const CredentialTypeUnknownCredentialType CredentialType = `UNKNOWN_CREDENTIAL_TYPE`

const CredentialTypeUsernamePassword CredentialType = `USERNAME_PASSWORD`

// String representation for [fmt.Print]
func (f *CredentialType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CredentialType) Set(v string) error {
	switch v {
	case `ANY_STATIC_CREDENTIAL`, `BEARER_TOKEN`, `OAUTH_ACCESS_TOKEN`, `OAUTH_M2M`, `OAUTH_REFRESH_TOKEN`, `OAUTH_RESOURCE_OWNER_PASSWORD`, `OAUTH_U2M`, `OAUTH_U2M_MAPPING`, `OIDC_TOKEN`, `PEM_PRIVATE_KEY`, `SERVICE_CREDENTIAL`, `UNKNOWN_CREDENTIAL_TYPE`, `USERNAME_PASSWORD`:
		*f = CredentialType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ANY_STATIC_CREDENTIAL", "BEARER_TOKEN", "OAUTH_ACCESS_TOKEN", "OAUTH_M2M", "OAUTH_REFRESH_TOKEN", "OAUTH_RESOURCE_OWNER_PASSWORD", "OAUTH_U2M", "OAUTH_U2M_MAPPING", "OIDC_TOKEN", "PEM_PRIVATE_KEY", "SERVICE_CREDENTIAL", "UNKNOWN_CREDENTIAL_TYPE", "USERNAME_PASSWORD"`, v)
	}
}

// Values returns all possible values for CredentialType.
//
// There is no guarantee on the order of the values in the slice.
func (f *CredentialType) Values() []CredentialType {
	return []CredentialType{
		CredentialTypeAnyStaticCredential,
		CredentialTypeBearerToken,
		CredentialTypeOauthAccessToken,
		CredentialTypeOauthM2m,
		CredentialTypeOauthRefreshToken,
		CredentialTypeOauthResourceOwnerPassword,
		CredentialTypeOauthU2m,
		CredentialTypeOauthU2mMapping,
		CredentialTypeOidcToken,
		CredentialTypePemPrivateKey,
		CredentialTypeServiceCredential,
		CredentialTypeUnknownCredentialType,
		CredentialTypeUsernamePassword,
	}
}

// Type always returns CredentialType to satisfy [pflag.Value] interface
func (f *CredentialType) Type() string {
	return "CredentialType"
}

func CredentialTypeToPb(st *CredentialType) (*catalogpb.CredentialTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.CredentialTypePb(*st)
	return &pb, nil
}

func CredentialTypeFromPb(pb *catalogpb.CredentialTypePb) (*CredentialType, error) {
	if pb == nil {
		return nil, nil
	}
	st := CredentialType(*pb)
	return &st, nil
}

type CredentialValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	// Wire name: 'message'
	Message string ``
	// The results of the tested operation.
	// Wire name: 'result'
	Result          ValidateCredentialResult ``
	ForceSendFields []string                 `tf:"-"`
}

func (s *CredentialValidationResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CredentialValidationResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CredentialValidationResultToPb(st *CredentialValidationResult) (*catalogpb.CredentialValidationResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.CredentialValidationResultPb{}
	pb.Message = st.Message
	resultPb, err := ValidateCredentialResultToPb(&st.Result)
	if err != nil {
		return nil, err
	}
	if resultPb != nil {
		pb.Result = *resultPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CredentialValidationResultFromPb(pb *catalogpb.CredentialValidationResultPb) (*CredentialValidationResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CredentialValidationResult{}
	st.Message = pb.Message
	resultField, err := ValidateCredentialResultFromPb(&pb.Result)
	if err != nil {
		return nil, err
	}
	if resultField != nil {
		st.Result = *resultField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Data source format
type DataSourceFormat string

const DataSourceFormatAvro DataSourceFormat = `AVRO`

const DataSourceFormatBigqueryFormat DataSourceFormat = `BIGQUERY_FORMAT`

const DataSourceFormatCsv DataSourceFormat = `CSV`

const DataSourceFormatDatabricksFormat DataSourceFormat = `DATABRICKS_FORMAT`

const DataSourceFormatDatabricksRowStoreFormat DataSourceFormat = `DATABRICKS_ROW_STORE_FORMAT`

const DataSourceFormatDelta DataSourceFormat = `DELTA`

const DataSourceFormatDeltasharing DataSourceFormat = `DELTASHARING`

const DataSourceFormatDeltaUniformHudi DataSourceFormat = `DELTA_UNIFORM_HUDI`

const DataSourceFormatDeltaUniformIceberg DataSourceFormat = `DELTA_UNIFORM_ICEBERG`

const DataSourceFormatHive DataSourceFormat = `HIVE`

const DataSourceFormatIceberg DataSourceFormat = `ICEBERG`

const DataSourceFormatJson DataSourceFormat = `JSON`

const DataSourceFormatMongodbFormat DataSourceFormat = `MONGODB_FORMAT`

const DataSourceFormatMysqlFormat DataSourceFormat = `MYSQL_FORMAT`

const DataSourceFormatNetsuiteFormat DataSourceFormat = `NETSUITE_FORMAT`

const DataSourceFormatOracleFormat DataSourceFormat = `ORACLE_FORMAT`

const DataSourceFormatOrc DataSourceFormat = `ORC`

const DataSourceFormatParquet DataSourceFormat = `PARQUET`

const DataSourceFormatPostgresqlFormat DataSourceFormat = `POSTGRESQL_FORMAT`

const DataSourceFormatRedshiftFormat DataSourceFormat = `REDSHIFT_FORMAT`

const DataSourceFormatSalesforceDataCloudFormat DataSourceFormat = `SALESFORCE_DATA_CLOUD_FORMAT`

const DataSourceFormatSalesforceFormat DataSourceFormat = `SALESFORCE_FORMAT`

const DataSourceFormatSnowflakeFormat DataSourceFormat = `SNOWFLAKE_FORMAT`

const DataSourceFormatSqldwFormat DataSourceFormat = `SQLDW_FORMAT`

const DataSourceFormatSqlserverFormat DataSourceFormat = `SQLSERVER_FORMAT`

const DataSourceFormatTeradataFormat DataSourceFormat = `TERADATA_FORMAT`

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
	case `AVRO`, `BIGQUERY_FORMAT`, `CSV`, `DATABRICKS_FORMAT`, `DATABRICKS_ROW_STORE_FORMAT`, `DELTA`, `DELTASHARING`, `DELTA_UNIFORM_HUDI`, `DELTA_UNIFORM_ICEBERG`, `HIVE`, `ICEBERG`, `JSON`, `MONGODB_FORMAT`, `MYSQL_FORMAT`, `NETSUITE_FORMAT`, `ORACLE_FORMAT`, `ORC`, `PARQUET`, `POSTGRESQL_FORMAT`, `REDSHIFT_FORMAT`, `SALESFORCE_DATA_CLOUD_FORMAT`, `SALESFORCE_FORMAT`, `SNOWFLAKE_FORMAT`, `SQLDW_FORMAT`, `SQLSERVER_FORMAT`, `TERADATA_FORMAT`, `TEXT`, `UNITY_CATALOG`, `VECTOR_INDEX_FORMAT`, `WORKDAY_RAAS_FORMAT`:
		*f = DataSourceFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVRO", "BIGQUERY_FORMAT", "CSV", "DATABRICKS_FORMAT", "DATABRICKS_ROW_STORE_FORMAT", "DELTA", "DELTASHARING", "DELTA_UNIFORM_HUDI", "DELTA_UNIFORM_ICEBERG", "HIVE", "ICEBERG", "JSON", "MONGODB_FORMAT", "MYSQL_FORMAT", "NETSUITE_FORMAT", "ORACLE_FORMAT", "ORC", "PARQUET", "POSTGRESQL_FORMAT", "REDSHIFT_FORMAT", "SALESFORCE_DATA_CLOUD_FORMAT", "SALESFORCE_FORMAT", "SNOWFLAKE_FORMAT", "SQLDW_FORMAT", "SQLSERVER_FORMAT", "TERADATA_FORMAT", "TEXT", "UNITY_CATALOG", "VECTOR_INDEX_FORMAT", "WORKDAY_RAAS_FORMAT"`, v)
	}
}

// Values returns all possible values for DataSourceFormat.
//
// There is no guarantee on the order of the values in the slice.
func (f *DataSourceFormat) Values() []DataSourceFormat {
	return []DataSourceFormat{
		DataSourceFormatAvro,
		DataSourceFormatBigqueryFormat,
		DataSourceFormatCsv,
		DataSourceFormatDatabricksFormat,
		DataSourceFormatDatabricksRowStoreFormat,
		DataSourceFormatDelta,
		DataSourceFormatDeltasharing,
		DataSourceFormatDeltaUniformHudi,
		DataSourceFormatDeltaUniformIceberg,
		DataSourceFormatHive,
		DataSourceFormatIceberg,
		DataSourceFormatJson,
		DataSourceFormatMongodbFormat,
		DataSourceFormatMysqlFormat,
		DataSourceFormatNetsuiteFormat,
		DataSourceFormatOracleFormat,
		DataSourceFormatOrc,
		DataSourceFormatParquet,
		DataSourceFormatPostgresqlFormat,
		DataSourceFormatRedshiftFormat,
		DataSourceFormatSalesforceDataCloudFormat,
		DataSourceFormatSalesforceFormat,
		DataSourceFormatSnowflakeFormat,
		DataSourceFormatSqldwFormat,
		DataSourceFormatSqlserverFormat,
		DataSourceFormatTeradataFormat,
		DataSourceFormatText,
		DataSourceFormatUnityCatalog,
		DataSourceFormatVectorIndexFormat,
		DataSourceFormatWorkdayRaasFormat,
	}
}

// Type always returns DataSourceFormat to satisfy [pflag.Value] interface
func (f *DataSourceFormat) Type() string {
	return "DataSourceFormat"
}

func DataSourceFormatToPb(st *DataSourceFormat) (*catalogpb.DataSourceFormatPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.DataSourceFormatPb(*st)
	return &pb, nil
}

func DataSourceFormatFromPb(pb *catalogpb.DataSourceFormatPb) (*DataSourceFormat, error) {
	if pb == nil {
		return nil, nil
	}
	st := DataSourceFormat(*pb)
	return &st, nil
}

// GCP long-lived credential. Databricks-created Google Cloud Storage service
// account.
type DatabricksGcpServiceAccount struct {
	// The Databricks internal ID that represents this managed identity.
	// Wire name: 'credential_id'
	CredentialId string ``
	// The email of the service account.
	// Wire name: 'email'
	Email string ``
	// The ID that represents the private key for this Service Account
	// Wire name: 'private_key_id'
	PrivateKeyId    string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DatabricksGcpServiceAccount) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksGcpServiceAccount) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DatabricksGcpServiceAccountToPb(st *DatabricksGcpServiceAccount) (*catalogpb.DatabricksGcpServiceAccountPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DatabricksGcpServiceAccountPb{}
	pb.CredentialId = st.CredentialId
	pb.Email = st.Email
	pb.PrivateKeyId = st.PrivateKeyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DatabricksGcpServiceAccountFromPb(pb *catalogpb.DatabricksGcpServiceAccountPb) (*DatabricksGcpServiceAccount, error) {
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

// GCP long-lived credential. Databricks-created Google Cloud Storage service
// account.
type DatabricksGcpServiceAccountRequest struct {
}

func DatabricksGcpServiceAccountRequestToPb(st *DatabricksGcpServiceAccountRequest) (*catalogpb.DatabricksGcpServiceAccountRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DatabricksGcpServiceAccountRequestPb{}

	return pb, nil
}

func DatabricksGcpServiceAccountRequestFromPb(pb *catalogpb.DatabricksGcpServiceAccountRequestPb) (*DatabricksGcpServiceAccountRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabricksGcpServiceAccountRequest{}

	return st, nil
}

// GCP long-lived credential. Databricks-created Google Cloud Storage service
// account.
type DatabricksGcpServiceAccountResponse struct {
	// The Databricks internal ID that represents this managed identity.
	// Wire name: 'credential_id'
	CredentialId string ``
	// The email of the service account.
	// Wire name: 'email'
	Email           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DatabricksGcpServiceAccountResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksGcpServiceAccountResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DatabricksGcpServiceAccountResponseToPb(st *DatabricksGcpServiceAccountResponse) (*catalogpb.DatabricksGcpServiceAccountResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DatabricksGcpServiceAccountResponsePb{}
	pb.CredentialId = st.CredentialId
	pb.Email = st.Email

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DatabricksGcpServiceAccountResponseFromPb(pb *catalogpb.DatabricksGcpServiceAccountResponsePb) (*DatabricksGcpServiceAccountResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabricksGcpServiceAccountResponse{}
	st.CredentialId = pb.CredentialId
	st.Email = pb.Email

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteAccountMetastoreAssignmentRequest struct {
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func DeleteAccountMetastoreAssignmentRequestToPb(st *DeleteAccountMetastoreAssignmentRequest) (*catalogpb.DeleteAccountMetastoreAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteAccountMetastoreAssignmentRequestPb{}
	pb.MetastoreId = st.MetastoreId
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func DeleteAccountMetastoreAssignmentRequestFromPb(pb *catalogpb.DeleteAccountMetastoreAssignmentRequestPb) (*DeleteAccountMetastoreAssignmentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountMetastoreAssignmentRequest{}
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

type DeleteAccountMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	// Wire name: 'force'
	Force bool `tf:"-"`
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId     string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *DeleteAccountMetastoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAccountMetastoreRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteAccountMetastoreRequestToPb(st *DeleteAccountMetastoreRequest) (*catalogpb.DeleteAccountMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteAccountMetastoreRequestPb{}
	pb.Force = st.Force
	pb.MetastoreId = st.MetastoreId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteAccountMetastoreRequestFromPb(pb *catalogpb.DeleteAccountMetastoreRequestPb) (*DeleteAccountMetastoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountMetastoreRequest{}
	st.Force = pb.Force
	st.MetastoreId = pb.MetastoreId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

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
	StorageCredentialName string   `tf:"-"`
	ForceSendFields       []string `tf:"-"`
}

func (s *DeleteAccountStorageCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAccountStorageCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteAccountStorageCredentialRequestToPb(st *DeleteAccountStorageCredentialRequest) (*catalogpb.DeleteAccountStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteAccountStorageCredentialRequestPb{}
	pb.Force = st.Force
	pb.MetastoreId = st.MetastoreId
	pb.StorageCredentialName = st.StorageCredentialName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteAccountStorageCredentialRequestFromPb(pb *catalogpb.DeleteAccountStorageCredentialRequestPb) (*DeleteAccountStorageCredentialRequest, error) {
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

type DeleteAliasRequest struct {
	// The name of the alias
	// Wire name: 'alias'
	Alias string `tf:"-"`
	// The three-level (fully qualified) name of the registered model
	// Wire name: 'full_name'
	FullName string `tf:"-"`
}

func DeleteAliasRequestToPb(st *DeleteAliasRequest) (*catalogpb.DeleteAliasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteAliasRequestPb{}
	pb.Alias = st.Alias
	pb.FullName = st.FullName

	return pb, nil
}

func DeleteAliasRequestFromPb(pb *catalogpb.DeleteAliasRequestPb) (*DeleteAliasRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAliasRequest{}
	st.Alias = pb.Alias
	st.FullName = pb.FullName

	return st, nil
}

type DeleteCatalogRequest struct {
	// Force deletion even if the catalog is not empty.
	// Wire name: 'force'
	Force bool `tf:"-"`
	// The name of the catalog.
	// Wire name: 'name'
	Name            string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *DeleteCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteCatalogRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteCatalogRequestToPb(st *DeleteCatalogRequest) (*catalogpb.DeleteCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteCatalogRequestPb{}
	pb.Force = st.Force
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteCatalogRequestFromPb(pb *catalogpb.DeleteCatalogRequestPb) (*DeleteCatalogRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCatalogRequest{}
	st.Force = pb.Force
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteConnectionRequest struct {
	// The name of the connection to be deleted.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func DeleteConnectionRequestToPb(st *DeleteConnectionRequest) (*catalogpb.DeleteConnectionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteConnectionRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func DeleteConnectionRequestFromPb(pb *catalogpb.DeleteConnectionRequestPb) (*DeleteConnectionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteConnectionRequest{}
	st.Name = pb.Name

	return st, nil
}

type DeleteCredentialRequest struct {
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	// Wire name: 'force'
	Force bool `tf:"-"`
	// Name of the credential.
	// Wire name: 'name_arg'
	NameArg         string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *DeleteCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteCredentialRequestToPb(st *DeleteCredentialRequest) (*catalogpb.DeleteCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteCredentialRequestPb{}
	pb.Force = st.Force
	pb.NameArg = st.NameArg

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteCredentialRequestFromPb(pb *catalogpb.DeleteCredentialRequestPb) (*DeleteCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCredentialRequest{}
	st.Force = pb.Force
	st.NameArg = pb.NameArg

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteExternalLineageRelationshipRequest struct {

	// Wire name: 'external_lineage_relationship'
	ExternalLineageRelationship DeleteRequestExternalLineage `tf:"-"`
}

func DeleteExternalLineageRelationshipRequestToPb(st *DeleteExternalLineageRelationshipRequest) (*catalogpb.DeleteExternalLineageRelationshipRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteExternalLineageRelationshipRequestPb{}
	externalLineageRelationshipPb, err := DeleteRequestExternalLineageToPb(&st.ExternalLineageRelationship)
	if err != nil {
		return nil, err
	}
	if externalLineageRelationshipPb != nil {
		pb.ExternalLineageRelationship = *externalLineageRelationshipPb
	}

	return pb, nil
}

func DeleteExternalLineageRelationshipRequestFromPb(pb *catalogpb.DeleteExternalLineageRelationshipRequestPb) (*DeleteExternalLineageRelationshipRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExternalLineageRelationshipRequest{}
	externalLineageRelationshipField, err := DeleteRequestExternalLineageFromPb(&pb.ExternalLineageRelationship)
	if err != nil {
		return nil, err
	}
	if externalLineageRelationshipField != nil {
		st.ExternalLineageRelationship = *externalLineageRelationshipField
	}

	return st, nil
}

type DeleteExternalLocationRequest struct {
	// Force deletion even if there are dependent external tables or mounts.
	// Wire name: 'force'
	Force bool `tf:"-"`
	// Name of the external location.
	// Wire name: 'name'
	Name            string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *DeleteExternalLocationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteExternalLocationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteExternalLocationRequestToPb(st *DeleteExternalLocationRequest) (*catalogpb.DeleteExternalLocationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteExternalLocationRequestPb{}
	pb.Force = st.Force
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteExternalLocationRequestFromPb(pb *catalogpb.DeleteExternalLocationRequestPb) (*DeleteExternalLocationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExternalLocationRequest{}
	st.Force = pb.Force
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteExternalMetadataRequest struct {

	// Wire name: 'name'
	Name string `tf:"-"`
}

func DeleteExternalMetadataRequestToPb(st *DeleteExternalMetadataRequest) (*catalogpb.DeleteExternalMetadataRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteExternalMetadataRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func DeleteExternalMetadataRequestFromPb(pb *catalogpb.DeleteExternalMetadataRequestPb) (*DeleteExternalMetadataRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExternalMetadataRequest{}
	st.Name = pb.Name

	return st, nil
}

type DeleteFunctionRequest struct {
	// Force deletion even if the function is notempty.
	// Wire name: 'force'
	Force bool `tf:"-"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	// Wire name: 'name'
	Name            string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *DeleteFunctionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteFunctionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteFunctionRequestToPb(st *DeleteFunctionRequest) (*catalogpb.DeleteFunctionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteFunctionRequestPb{}
	pb.Force = st.Force
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteFunctionRequestFromPb(pb *catalogpb.DeleteFunctionRequestPb) (*DeleteFunctionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteFunctionRequest{}
	st.Force = pb.Force
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	// Wire name: 'force'
	Force bool `tf:"-"`
	// Unique ID of the metastore.
	// Wire name: 'id'
	Id              string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *DeleteMetastoreRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteMetastoreRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteMetastoreRequestToPb(st *DeleteMetastoreRequest) (*catalogpb.DeleteMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteMetastoreRequestPb{}
	pb.Force = st.Force
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteMetastoreRequestFromPb(pb *catalogpb.DeleteMetastoreRequestPb) (*DeleteMetastoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteMetastoreRequest{}
	st.Force = pb.Force
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteModelVersionRequest struct {
	// The three-level (fully qualified) name of the model version
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// The integer version number of the model version
	// Wire name: 'version'
	Version int `tf:"-"`
}

func DeleteModelVersionRequestToPb(st *DeleteModelVersionRequest) (*catalogpb.DeleteModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteModelVersionRequestPb{}
	pb.FullName = st.FullName
	pb.Version = st.Version

	return pb, nil
}

func DeleteModelVersionRequestFromPb(pb *catalogpb.DeleteModelVersionRequestPb) (*DeleteModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelVersionRequest{}
	st.FullName = pb.FullName
	st.Version = pb.Version

	return st, nil
}

type DeleteMonitorResponse struct {
}

func DeleteMonitorResponseToPb(st *DeleteMonitorResponse) (*catalogpb.DeleteMonitorResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteMonitorResponsePb{}

	return pb, nil
}

func DeleteMonitorResponseFromPb(pb *catalogpb.DeleteMonitorResponsePb) (*DeleteMonitorResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteMonitorResponse{}

	return st, nil
}

type DeleteOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func DeleteOnlineTableRequestToPb(st *DeleteOnlineTableRequest) (*catalogpb.DeleteOnlineTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteOnlineTableRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func DeleteOnlineTableRequestFromPb(pb *catalogpb.DeleteOnlineTableRequestPb) (*DeleteOnlineTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteOnlineTableRequest{}
	st.Name = pb.Name

	return st, nil
}

type DeleteQualityMonitorRequest struct {
	// UC table name in format `catalog.schema.table_name`. This field
	// corresponds to the {full_table_name_arg} arg in the endpoint path.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func DeleteQualityMonitorRequestToPb(st *DeleteQualityMonitorRequest) (*catalogpb.DeleteQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteQualityMonitorRequestPb{}
	pb.TableName = st.TableName

	return pb, nil
}

func DeleteQualityMonitorRequestFromPb(pb *catalogpb.DeleteQualityMonitorRequestPb) (*DeleteQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteQualityMonitorRequest{}
	st.TableName = pb.TableName

	return st, nil
}

type DeleteRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	// Wire name: 'full_name'
	FullName string `tf:"-"`
}

func DeleteRegisteredModelRequestToPb(st *DeleteRegisteredModelRequest) (*catalogpb.DeleteRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteRegisteredModelRequestPb{}
	pb.FullName = st.FullName

	return pb, nil
}

func DeleteRegisteredModelRequestFromPb(pb *catalogpb.DeleteRegisteredModelRequestPb) (*DeleteRegisteredModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRegisteredModelRequest{}
	st.FullName = pb.FullName

	return st, nil
}

type DeleteRequestExternalLineage struct {
	// Unique identifier of the external lineage relationship.
	// Wire name: 'id'
	Id string ``
	// Source object of the external lineage relationship.
	// Wire name: 'source'
	Source ExternalLineageObject ``
	// Target object of the external lineage relationship.
	// Wire name: 'target'
	Target          ExternalLineageObject ``
	ForceSendFields []string              `tf:"-"`
}

func (s *DeleteRequestExternalLineage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteRequestExternalLineage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteRequestExternalLineageToPb(st *DeleteRequestExternalLineage) (*catalogpb.DeleteRequestExternalLineagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteRequestExternalLineagePb{}
	pb.Id = st.Id
	sourcePb, err := ExternalLineageObjectToPb(&st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = *sourcePb
	}
	targetPb, err := ExternalLineageObjectToPb(&st.Target)
	if err != nil {
		return nil, err
	}
	if targetPb != nil {
		pb.Target = *targetPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteRequestExternalLineageFromPb(pb *catalogpb.DeleteRequestExternalLineagePb) (*DeleteRequestExternalLineage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRequestExternalLineage{}
	st.Id = pb.Id
	sourceField, err := ExternalLineageObjectFromPb(&pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = *sourceField
	}
	targetField, err := ExternalLineageObjectFromPb(&pb.Target)
	if err != nil {
		return nil, err
	}
	if targetField != nil {
		st.Target = *targetField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteSchemaRequest struct {
	// Force deletion even if the schema is not empty.
	// Wire name: 'force'
	Force bool `tf:"-"`
	// Full name of the schema.
	// Wire name: 'full_name'
	FullName        string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *DeleteSchemaRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteSchemaRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteSchemaRequestToPb(st *DeleteSchemaRequest) (*catalogpb.DeleteSchemaRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteSchemaRequestPb{}
	pb.Force = st.Force
	pb.FullName = st.FullName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteSchemaRequestFromPb(pb *catalogpb.DeleteSchemaRequestPb) (*DeleteSchemaRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSchemaRequest{}
	st.Force = pb.Force
	st.FullName = pb.FullName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteStorageCredentialRequest struct {
	// Force an update even if there are dependent external locations or
	// external tables (when purpose is **STORAGE**) or dependent services (when
	// purpose is **SERVICE**).
	// Wire name: 'force'
	Force bool `tf:"-"`
	// Name of the storage credential.
	// Wire name: 'name'
	Name            string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *DeleteStorageCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteStorageCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteStorageCredentialRequestToPb(st *DeleteStorageCredentialRequest) (*catalogpb.DeleteStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteStorageCredentialRequestPb{}
	pb.Force = st.Force
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteStorageCredentialRequestFromPb(pb *catalogpb.DeleteStorageCredentialRequestPb) (*DeleteStorageCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteStorageCredentialRequest{}
	st.Force = pb.Force
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

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

func DeleteTableConstraintRequestToPb(st *DeleteTableConstraintRequest) (*catalogpb.DeleteTableConstraintRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteTableConstraintRequestPb{}
	pb.Cascade = st.Cascade
	pb.ConstraintName = st.ConstraintName
	pb.FullName = st.FullName

	return pb, nil
}

func DeleteTableConstraintRequestFromPb(pb *catalogpb.DeleteTableConstraintRequestPb) (*DeleteTableConstraintRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTableConstraintRequest{}
	st.Cascade = pb.Cascade
	st.ConstraintName = pb.ConstraintName
	st.FullName = pb.FullName

	return st, nil
}

type DeleteTableRequest struct {
	// Full name of the table.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
}

func DeleteTableRequestToPb(st *DeleteTableRequest) (*catalogpb.DeleteTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteTableRequestPb{}
	pb.FullName = st.FullName

	return pb, nil
}

func DeleteTableRequestFromPb(pb *catalogpb.DeleteTableRequestPb) (*DeleteTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTableRequest{}
	st.FullName = pb.FullName

	return st, nil
}

type DeleteVolumeRequest struct {
	// The three-level (fully qualified) name of the volume
	// Wire name: 'name'
	Name string `tf:"-"`
}

func DeleteVolumeRequestToPb(st *DeleteVolumeRequest) (*catalogpb.DeleteVolumeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeleteVolumeRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func DeleteVolumeRequestFromPb(pb *catalogpb.DeleteVolumeRequestPb) (*DeleteVolumeRequest, error) {
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
	DeltaRuntimeProperties map[string]string ``
}

func DeltaRuntimePropertiesKvPairsToPb(st *DeltaRuntimePropertiesKvPairs) (*catalogpb.DeltaRuntimePropertiesKvPairsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DeltaRuntimePropertiesKvPairsPb{}
	pb.DeltaRuntimeProperties = st.DeltaRuntimeProperties

	return pb, nil
}

func DeltaRuntimePropertiesKvPairsFromPb(pb *catalogpb.DeltaRuntimePropertiesKvPairsPb) (*DeltaRuntimePropertiesKvPairs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaRuntimePropertiesKvPairs{}
	st.DeltaRuntimeProperties = pb.DeltaRuntimeProperties

	return st, nil
}

type DeltaSharingScopeEnum string

const DeltaSharingScopeEnumInternal DeltaSharingScopeEnum = `INTERNAL`

const DeltaSharingScopeEnumInternalAndExternal DeltaSharingScopeEnum = `INTERNAL_AND_EXTERNAL`

// String representation for [fmt.Print]
func (f *DeltaSharingScopeEnum) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeltaSharingScopeEnum) Set(v string) error {
	switch v {
	case `INTERNAL`, `INTERNAL_AND_EXTERNAL`:
		*f = DeltaSharingScopeEnum(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INTERNAL", "INTERNAL_AND_EXTERNAL"`, v)
	}
}

// Values returns all possible values for DeltaSharingScopeEnum.
//
// There is no guarantee on the order of the values in the slice.
func (f *DeltaSharingScopeEnum) Values() []DeltaSharingScopeEnum {
	return []DeltaSharingScopeEnum{
		DeltaSharingScopeEnumInternal,
		DeltaSharingScopeEnumInternalAndExternal,
	}
}

// Type always returns DeltaSharingScopeEnum to satisfy [pflag.Value] interface
func (f *DeltaSharingScopeEnum) Type() string {
	return "DeltaSharingScopeEnum"
}

func DeltaSharingScopeEnumToPb(st *DeltaSharingScopeEnum) (*catalogpb.DeltaSharingScopeEnumPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.DeltaSharingScopeEnumPb(*st)
	return &pb, nil
}

func DeltaSharingScopeEnumFromPb(pb *catalogpb.DeltaSharingScopeEnumPb) (*DeltaSharingScopeEnum, error) {
	if pb == nil {
		return nil, nil
	}
	st := DeltaSharingScopeEnum(*pb)
	return &st, nil
}

// A dependency of a SQL object. One of the following fields must be defined:
// __table__, __function__, __connection__, or __credential__.
type Dependency struct {

	// Wire name: 'connection'
	Connection *ConnectionDependency ``

	// Wire name: 'credential'
	Credential *CredentialDependency ``

	// Wire name: 'function'
	Function *FunctionDependency ``

	// Wire name: 'table'
	Table *TableDependency ``
}

func DependencyToPb(st *Dependency) (*catalogpb.DependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DependencyPb{}
	connectionPb, err := ConnectionDependencyToPb(st.Connection)
	if err != nil {
		return nil, err
	}
	if connectionPb != nil {
		pb.Connection = connectionPb
	}
	credentialPb, err := CredentialDependencyToPb(st.Credential)
	if err != nil {
		return nil, err
	}
	if credentialPb != nil {
		pb.Credential = credentialPb
	}
	functionPb, err := FunctionDependencyToPb(st.Function)
	if err != nil {
		return nil, err
	}
	if functionPb != nil {
		pb.Function = functionPb
	}
	tablePb, err := TableDependencyToPb(st.Table)
	if err != nil {
		return nil, err
	}
	if tablePb != nil {
		pb.Table = tablePb
	}

	return pb, nil
}

func DependencyFromPb(pb *catalogpb.DependencyPb) (*Dependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Dependency{}
	connectionField, err := ConnectionDependencyFromPb(pb.Connection)
	if err != nil {
		return nil, err
	}
	if connectionField != nil {
		st.Connection = connectionField
	}
	credentialField, err := CredentialDependencyFromPb(pb.Credential)
	if err != nil {
		return nil, err
	}
	if credentialField != nil {
		st.Credential = credentialField
	}
	functionField, err := FunctionDependencyFromPb(pb.Function)
	if err != nil {
		return nil, err
	}
	if functionField != nil {
		st.Function = functionField
	}
	tableField, err := TableDependencyFromPb(pb.Table)
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
	Dependencies []Dependency ``
}

func DependencyListToPb(st *DependencyList) (*catalogpb.DependencyListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DependencyListPb{}

	var dependenciesPb []catalogpb.DependencyPb
	for _, item := range st.Dependencies {
		itemPb, err := DependencyToPb(&item)
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

func DependencyListFromPb(pb *catalogpb.DependencyListPb) (*DependencyList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DependencyList{}

	var dependenciesField []Dependency
	for _, itemPb := range pb.Dependencies {
		item, err := DependencyFromPb(&itemPb)
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

type DisableRequest struct {
	// The metastore ID under which the system schema lives.
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Full name of the system schema.
	// Wire name: 'schema_name'
	SchemaName string `tf:"-"`
}

func DisableRequestToPb(st *DisableRequest) (*catalogpb.DisableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.DisableRequestPb{}
	pb.MetastoreId = st.MetastoreId
	pb.SchemaName = st.SchemaName

	return pb, nil
}

func DisableRequestFromPb(pb *catalogpb.DisableRequestPb) (*DisableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DisableRequest{}
	st.MetastoreId = pb.MetastoreId
	st.SchemaName = pb.SchemaName

	return st, nil
}

type EffectivePermissionsList struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// The privileges conveyed to each principal (either directly or via
	// inheritance)
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []EffectivePrivilegeAssignment ``
	ForceSendFields      []string                       `tf:"-"`
}

func (s *EffectivePermissionsList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EffectivePermissionsList) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EffectivePermissionsListToPb(st *EffectivePermissionsList) (*catalogpb.EffectivePermissionsListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.EffectivePermissionsListPb{}
	pb.NextPageToken = st.NextPageToken

	var privilegeAssignmentsPb []catalogpb.EffectivePrivilegeAssignmentPb
	for _, item := range st.PrivilegeAssignments {
		itemPb, err := EffectivePrivilegeAssignmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			privilegeAssignmentsPb = append(privilegeAssignmentsPb, *itemPb)
		}
	}
	pb.PrivilegeAssignments = privilegeAssignmentsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EffectivePermissionsListFromPb(pb *catalogpb.EffectivePermissionsListPb) (*EffectivePermissionsList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EffectivePermissionsList{}
	st.NextPageToken = pb.NextPageToken

	var privilegeAssignmentsField []EffectivePrivilegeAssignment
	for _, itemPb := range pb.PrivilegeAssignments {
		item, err := EffectivePrivilegeAssignmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			privilegeAssignmentsField = append(privilegeAssignmentsField, *item)
		}
	}
	st.PrivilegeAssignments = privilegeAssignmentsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type EffectivePredictiveOptimizationFlag struct {
	// The name of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	// Wire name: 'inherited_from_name'
	InheritedFromName string ``
	// The type of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	// Wire name: 'inherited_from_type'
	InheritedFromType EffectivePredictiveOptimizationFlagInheritedFromType ``
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'value'
	Value           EnablePredictiveOptimization ``
	ForceSendFields []string                     `tf:"-"`
}

func (s *EffectivePredictiveOptimizationFlag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EffectivePredictiveOptimizationFlag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EffectivePredictiveOptimizationFlagToPb(st *EffectivePredictiveOptimizationFlag) (*catalogpb.EffectivePredictiveOptimizationFlagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.EffectivePredictiveOptimizationFlagPb{}
	pb.InheritedFromName = st.InheritedFromName
	inheritedFromTypePb, err := EffectivePredictiveOptimizationFlagInheritedFromTypeToPb(&st.InheritedFromType)
	if err != nil {
		return nil, err
	}
	if inheritedFromTypePb != nil {
		pb.InheritedFromType = *inheritedFromTypePb
	}
	valuePb, err := EnablePredictiveOptimizationToPb(&st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = *valuePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EffectivePredictiveOptimizationFlagFromPb(pb *catalogpb.EffectivePredictiveOptimizationFlagPb) (*EffectivePredictiveOptimizationFlag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EffectivePredictiveOptimizationFlag{}
	st.InheritedFromName = pb.InheritedFromName
	inheritedFromTypeField, err := EffectivePredictiveOptimizationFlagInheritedFromTypeFromPb(&pb.InheritedFromType)
	if err != nil {
		return nil, err
	}
	if inheritedFromTypeField != nil {
		st.InheritedFromType = *inheritedFromTypeField
	}
	valueField, err := EnablePredictiveOptimizationFromPb(&pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = *valueField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// The type of the object from which the flag was inherited. If there was no
// inheritance, this field is left blank.
type EffectivePredictiveOptimizationFlagInheritedFromType string

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

// Values returns all possible values for EffectivePredictiveOptimizationFlagInheritedFromType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EffectivePredictiveOptimizationFlagInheritedFromType) Values() []EffectivePredictiveOptimizationFlagInheritedFromType {
	return []EffectivePredictiveOptimizationFlagInheritedFromType{
		EffectivePredictiveOptimizationFlagInheritedFromTypeCatalog,
		EffectivePredictiveOptimizationFlagInheritedFromTypeSchema,
	}
}

// Type always returns EffectivePredictiveOptimizationFlagInheritedFromType to satisfy [pflag.Value] interface
func (f *EffectivePredictiveOptimizationFlagInheritedFromType) Type() string {
	return "EffectivePredictiveOptimizationFlagInheritedFromType"
}

func EffectivePredictiveOptimizationFlagInheritedFromTypeToPb(st *EffectivePredictiveOptimizationFlagInheritedFromType) (*catalogpb.EffectivePredictiveOptimizationFlagInheritedFromTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.EffectivePredictiveOptimizationFlagInheritedFromTypePb(*st)
	return &pb, nil
}

func EffectivePredictiveOptimizationFlagInheritedFromTypeFromPb(pb *catalogpb.EffectivePredictiveOptimizationFlagInheritedFromTypePb) (*EffectivePredictiveOptimizationFlagInheritedFromType, error) {
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
	InheritedFromName string ``
	// The type of the object that conveys this privilege via inheritance. This
	// field is omitted when privilege is not inherited (it's assigned to the
	// securable itself).
	// Wire name: 'inherited_from_type'
	InheritedFromType SecurableType ``
	// The privilege assigned to the principal.
	// Wire name: 'privilege'
	Privilege       Privilege ``
	ForceSendFields []string  `tf:"-"`
}

func (s *EffectivePrivilege) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EffectivePrivilege) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EffectivePrivilegeToPb(st *EffectivePrivilege) (*catalogpb.EffectivePrivilegePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.EffectivePrivilegePb{}
	pb.InheritedFromName = st.InheritedFromName
	inheritedFromTypePb, err := SecurableTypeToPb(&st.InheritedFromType)
	if err != nil {
		return nil, err
	}
	if inheritedFromTypePb != nil {
		pb.InheritedFromType = *inheritedFromTypePb
	}
	privilegePb, err := PrivilegeToPb(&st.Privilege)
	if err != nil {
		return nil, err
	}
	if privilegePb != nil {
		pb.Privilege = *privilegePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EffectivePrivilegeFromPb(pb *catalogpb.EffectivePrivilegePb) (*EffectivePrivilege, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EffectivePrivilege{}
	st.InheritedFromName = pb.InheritedFromName
	inheritedFromTypeField, err := SecurableTypeFromPb(&pb.InheritedFromType)
	if err != nil {
		return nil, err
	}
	if inheritedFromTypeField != nil {
		st.InheritedFromType = *inheritedFromTypeField
	}
	privilegeField, err := PrivilegeFromPb(&pb.Privilege)
	if err != nil {
		return nil, err
	}
	if privilegeField != nil {
		st.Privilege = *privilegeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type EffectivePrivilegeAssignment struct {
	// The principal (user email address or group name).
	// Wire name: 'principal'
	Principal string ``
	// The privileges conveyed to the principal (either directly or via
	// inheritance).
	// Wire name: 'privileges'
	Privileges      []EffectivePrivilege ``
	ForceSendFields []string             `tf:"-"`
}

func (s *EffectivePrivilegeAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EffectivePrivilegeAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EffectivePrivilegeAssignmentToPb(st *EffectivePrivilegeAssignment) (*catalogpb.EffectivePrivilegeAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.EffectivePrivilegeAssignmentPb{}
	pb.Principal = st.Principal

	var privilegesPb []catalogpb.EffectivePrivilegePb
	for _, item := range st.Privileges {
		itemPb, err := EffectivePrivilegeToPb(&item)
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

func EffectivePrivilegeAssignmentFromPb(pb *catalogpb.EffectivePrivilegeAssignmentPb) (*EffectivePrivilegeAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EffectivePrivilegeAssignment{}
	st.Principal = pb.Principal

	var privilegesField []EffectivePrivilege
	for _, itemPb := range pb.Privileges {
		item, err := EffectivePrivilegeFromPb(&itemPb)
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

type EnablePredictiveOptimization string

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

// Values returns all possible values for EnablePredictiveOptimization.
//
// There is no guarantee on the order of the values in the slice.
func (f *EnablePredictiveOptimization) Values() []EnablePredictiveOptimization {
	return []EnablePredictiveOptimization{
		EnablePredictiveOptimizationDisable,
		EnablePredictiveOptimizationEnable,
		EnablePredictiveOptimizationInherit,
	}
}

// Type always returns EnablePredictiveOptimization to satisfy [pflag.Value] interface
func (f *EnablePredictiveOptimization) Type() string {
	return "EnablePredictiveOptimization"
}

func EnablePredictiveOptimizationToPb(st *EnablePredictiveOptimization) (*catalogpb.EnablePredictiveOptimizationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.EnablePredictiveOptimizationPb(*st)
	return &pb, nil
}

func EnablePredictiveOptimizationFromPb(pb *catalogpb.EnablePredictiveOptimizationPb) (*EnablePredictiveOptimization, error) {
	if pb == nil {
		return nil, nil
	}
	st := EnablePredictiveOptimization(*pb)
	return &st, nil
}

type EnableRequest struct {
	// the catalog for which the system schema is to enabled in
	// Wire name: 'catalog_name'
	CatalogName string ``
	// The metastore ID under which the system schema lives.
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Full name of the system schema.
	// Wire name: 'schema_name'
	SchemaName      string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *EnableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EnableRequestToPb(st *EnableRequest) (*catalogpb.EnableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.EnableRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.MetastoreId = st.MetastoreId
	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EnableRequestFromPb(pb *catalogpb.EnableRequestPb) (*EnableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnableRequest{}
	st.CatalogName = pb.CatalogName
	st.MetastoreId = pb.MetastoreId
	st.SchemaName = pb.SchemaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Encryption options that apply to clients connecting to cloud storage.
type EncryptionDetails struct {
	// Server-Side Encryption properties for clients communicating with AWS s3.
	// Wire name: 'sse_encryption_details'
	SseEncryptionDetails *SseEncryptionDetails ``
}

func EncryptionDetailsToPb(st *EncryptionDetails) (*catalogpb.EncryptionDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.EncryptionDetailsPb{}
	sseEncryptionDetailsPb, err := SseEncryptionDetailsToPb(st.SseEncryptionDetails)
	if err != nil {
		return nil, err
	}
	if sseEncryptionDetailsPb != nil {
		pb.SseEncryptionDetails = sseEncryptionDetailsPb
	}

	return pb, nil
}

func EncryptionDetailsFromPb(pb *catalogpb.EncryptionDetailsPb) (*EncryptionDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EncryptionDetails{}
	sseEncryptionDetailsField, err := SseEncryptionDetailsFromPb(pb.SseEncryptionDetails)
	if err != nil {
		return nil, err
	}
	if sseEncryptionDetailsField != nil {
		st.SseEncryptionDetails = sseEncryptionDetailsField
	}

	return st, nil
}

type EnvironmentSettings struct {

	// Wire name: 'environment_version'
	EnvironmentVersion string ``

	// Wire name: 'java_dependencies'
	JavaDependencies []string ``
	ForceSendFields  []string `tf:"-"`
}

func (s *EnvironmentSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnvironmentSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EnvironmentSettingsToPb(st *EnvironmentSettings) (*catalogpb.EnvironmentSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.EnvironmentSettingsPb{}
	pb.EnvironmentVersion = st.EnvironmentVersion
	pb.JavaDependencies = st.JavaDependencies

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EnvironmentSettingsFromPb(pb *catalogpb.EnvironmentSettingsPb) (*EnvironmentSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnvironmentSettings{}
	st.EnvironmentVersion = pb.EnvironmentVersion
	st.JavaDependencies = pb.JavaDependencies

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ExistsRequest struct {
	// Full name of the table.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
}

func ExistsRequestToPb(st *ExistsRequest) (*catalogpb.ExistsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ExistsRequestPb{}
	pb.FullName = st.FullName

	return pb, nil
}

func ExistsRequestFromPb(pb *catalogpb.ExistsRequestPb) (*ExistsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExistsRequest{}
	st.FullName = pb.FullName

	return st, nil
}

type ExternalLineageExternalMetadata struct {

	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExternalLineageExternalMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageExternalMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalLineageExternalMetadataToPb(st *ExternalLineageExternalMetadata) (*catalogpb.ExternalLineageExternalMetadataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ExternalLineageExternalMetadataPb{}
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExternalLineageExternalMetadataFromPb(pb *catalogpb.ExternalLineageExternalMetadataPb) (*ExternalLineageExternalMetadata, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLineageExternalMetadata{}
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Represents the external metadata object in the lineage event.
type ExternalLineageExternalMetadataInfo struct {
	// Type of entity represented by the external metadata object.
	// Wire name: 'entity_type'
	EntityType string ``
	// Timestamp of the lineage event.
	// Wire name: 'event_time'
	EventTime *time.Time ``
	// Name of the external metadata object.
	// Wire name: 'name'
	Name string ``
	// Type of external system.
	// Wire name: 'system_type'
	SystemType      SystemType ``
	ForceSendFields []string   `tf:"-"`
}

func (s *ExternalLineageExternalMetadataInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageExternalMetadataInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalLineageExternalMetadataInfoToPb(st *ExternalLineageExternalMetadataInfo) (*catalogpb.ExternalLineageExternalMetadataInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ExternalLineageExternalMetadataInfoPb{}
	pb.EntityType = st.EntityType
	eventTimePb, err := timestampToPb(st.EventTime)
	if err != nil {
		return nil, err
	}
	if eventTimePb != nil {
		pb.EventTime = *eventTimePb
	}
	pb.Name = st.Name
	systemTypePb, err := SystemTypeToPb(&st.SystemType)
	if err != nil {
		return nil, err
	}
	if systemTypePb != nil {
		pb.SystemType = *systemTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExternalLineageExternalMetadataInfoFromPb(pb *catalogpb.ExternalLineageExternalMetadataInfoPb) (*ExternalLineageExternalMetadataInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLineageExternalMetadataInfo{}
	st.EntityType = pb.EntityType
	eventTimeField, err := timestampFromPb(&pb.EventTime)
	if err != nil {
		return nil, err
	}
	if eventTimeField != nil {
		st.EventTime = eventTimeField
	}
	st.Name = pb.Name
	systemTypeField, err := SystemTypeFromPb(&pb.SystemType)
	if err != nil {
		return nil, err
	}
	if systemTypeField != nil {
		st.SystemType = *systemTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Represents the path information in the lineage event.
type ExternalLineageFileInfo struct {
	// Timestamp of the lineage event.
	// Wire name: 'event_time'
	EventTime *time.Time ``
	// URL of the path.
	// Wire name: 'path'
	Path string ``
	// The full name of the securable on the path.
	// Wire name: 'securable_name'
	SecurableName string ``
	// The securable type of the securable on the path.
	// Wire name: 'securable_type'
	SecurableType string ``
	// The storage location associated with securable on the path.
	// Wire name: 'storage_location'
	StorageLocation string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExternalLineageFileInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageFileInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalLineageFileInfoToPb(st *ExternalLineageFileInfo) (*catalogpb.ExternalLineageFileInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ExternalLineageFileInfoPb{}
	eventTimePb, err := timestampToPb(st.EventTime)
	if err != nil {
		return nil, err
	}
	if eventTimePb != nil {
		pb.EventTime = *eventTimePb
	}
	pb.Path = st.Path
	pb.SecurableName = st.SecurableName
	pb.SecurableType = st.SecurableType
	pb.StorageLocation = st.StorageLocation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExternalLineageFileInfoFromPb(pb *catalogpb.ExternalLineageFileInfoPb) (*ExternalLineageFileInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLineageFileInfo{}
	eventTimeField, err := timestampFromPb(&pb.EventTime)
	if err != nil {
		return nil, err
	}
	if eventTimeField != nil {
		st.EventTime = eventTimeField
	}
	st.Path = pb.Path
	st.SecurableName = pb.SecurableName
	st.SecurableType = pb.SecurableType
	st.StorageLocation = pb.StorageLocation

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Lineage response containing lineage information of a data asset.
type ExternalLineageInfo struct {
	// Information about the edge metadata of the external lineage relationship.
	// Wire name: 'external_lineage_info'
	ExternalLineageInfo *ExternalLineageRelationshipInfo ``
	// Information about external metadata involved in the lineage relationship.
	// Wire name: 'external_metadata_info'
	ExternalMetadataInfo *ExternalLineageExternalMetadataInfo ``
	// Information about the file involved in the lineage relationship.
	// Wire name: 'file_info'
	FileInfo *ExternalLineageFileInfo ``
	// Information about the model version involved in the lineage relationship.
	// Wire name: 'model_info'
	ModelInfo *ExternalLineageModelVersionInfo ``
	// Information about the table involved in the lineage relationship.
	// Wire name: 'table_info'
	TableInfo *ExternalLineageTableInfo ``
}

func ExternalLineageInfoToPb(st *ExternalLineageInfo) (*catalogpb.ExternalLineageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ExternalLineageInfoPb{}
	externalLineageInfoPb, err := ExternalLineageRelationshipInfoToPb(st.ExternalLineageInfo)
	if err != nil {
		return nil, err
	}
	if externalLineageInfoPb != nil {
		pb.ExternalLineageInfo = externalLineageInfoPb
	}
	externalMetadataInfoPb, err := ExternalLineageExternalMetadataInfoToPb(st.ExternalMetadataInfo)
	if err != nil {
		return nil, err
	}
	if externalMetadataInfoPb != nil {
		pb.ExternalMetadataInfo = externalMetadataInfoPb
	}
	fileInfoPb, err := ExternalLineageFileInfoToPb(st.FileInfo)
	if err != nil {
		return nil, err
	}
	if fileInfoPb != nil {
		pb.FileInfo = fileInfoPb
	}
	modelInfoPb, err := ExternalLineageModelVersionInfoToPb(st.ModelInfo)
	if err != nil {
		return nil, err
	}
	if modelInfoPb != nil {
		pb.ModelInfo = modelInfoPb
	}
	tableInfoPb, err := ExternalLineageTableInfoToPb(st.TableInfo)
	if err != nil {
		return nil, err
	}
	if tableInfoPb != nil {
		pb.TableInfo = tableInfoPb
	}

	return pb, nil
}

func ExternalLineageInfoFromPb(pb *catalogpb.ExternalLineageInfoPb) (*ExternalLineageInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLineageInfo{}
	externalLineageInfoField, err := ExternalLineageRelationshipInfoFromPb(pb.ExternalLineageInfo)
	if err != nil {
		return nil, err
	}
	if externalLineageInfoField != nil {
		st.ExternalLineageInfo = externalLineageInfoField
	}
	externalMetadataInfoField, err := ExternalLineageExternalMetadataInfoFromPb(pb.ExternalMetadataInfo)
	if err != nil {
		return nil, err
	}
	if externalMetadataInfoField != nil {
		st.ExternalMetadataInfo = externalMetadataInfoField
	}
	fileInfoField, err := ExternalLineageFileInfoFromPb(pb.FileInfo)
	if err != nil {
		return nil, err
	}
	if fileInfoField != nil {
		st.FileInfo = fileInfoField
	}
	modelInfoField, err := ExternalLineageModelVersionInfoFromPb(pb.ModelInfo)
	if err != nil {
		return nil, err
	}
	if modelInfoField != nil {
		st.ModelInfo = modelInfoField
	}
	tableInfoField, err := ExternalLineageTableInfoFromPb(pb.TableInfo)
	if err != nil {
		return nil, err
	}
	if tableInfoField != nil {
		st.TableInfo = tableInfoField
	}

	return st, nil
}

type ExternalLineageModelVersion struct {

	// Wire name: 'name'
	Name string ``

	// Wire name: 'version'
	Version         string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExternalLineageModelVersion) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageModelVersion) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalLineageModelVersionToPb(st *ExternalLineageModelVersion) (*catalogpb.ExternalLineageModelVersionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ExternalLineageModelVersionPb{}
	pb.Name = st.Name
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExternalLineageModelVersionFromPb(pb *catalogpb.ExternalLineageModelVersionPb) (*ExternalLineageModelVersion, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLineageModelVersion{}
	st.Name = pb.Name
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Represents the model version information in the lineage event.
type ExternalLineageModelVersionInfo struct {
	// Timestamp of the lineage event.
	// Wire name: 'event_time'
	EventTime *time.Time ``
	// Name of the model.
	// Wire name: 'model_name'
	ModelName string ``
	// Version number of the model.
	// Wire name: 'version'
	Version         int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExternalLineageModelVersionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageModelVersionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalLineageModelVersionInfoToPb(st *ExternalLineageModelVersionInfo) (*catalogpb.ExternalLineageModelVersionInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ExternalLineageModelVersionInfoPb{}
	eventTimePb, err := timestampToPb(st.EventTime)
	if err != nil {
		return nil, err
	}
	if eventTimePb != nil {
		pb.EventTime = *eventTimePb
	}
	pb.ModelName = st.ModelName
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExternalLineageModelVersionInfoFromPb(pb *catalogpb.ExternalLineageModelVersionInfoPb) (*ExternalLineageModelVersionInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLineageModelVersionInfo{}
	eventTimeField, err := timestampFromPb(&pb.EventTime)
	if err != nil {
		return nil, err
	}
	if eventTimeField != nil {
		st.EventTime = eventTimeField
	}
	st.ModelName = pb.ModelName
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ExternalLineageObject struct {

	// Wire name: 'external_metadata'
	ExternalMetadata *ExternalLineageExternalMetadata ``

	// Wire name: 'model_version'
	ModelVersion *ExternalLineageModelVersion ``

	// Wire name: 'path'
	Path *ExternalLineagePath ``

	// Wire name: 'table'
	Table *ExternalLineageTable ``
}

func ExternalLineageObjectToPb(st *ExternalLineageObject) (*catalogpb.ExternalLineageObjectPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ExternalLineageObjectPb{}
	externalMetadataPb, err := ExternalLineageExternalMetadataToPb(st.ExternalMetadata)
	if err != nil {
		return nil, err
	}
	if externalMetadataPb != nil {
		pb.ExternalMetadata = externalMetadataPb
	}
	modelVersionPb, err := ExternalLineageModelVersionToPb(st.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionPb != nil {
		pb.ModelVersion = modelVersionPb
	}
	pathPb, err := ExternalLineagePathToPb(st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = pathPb
	}
	tablePb, err := ExternalLineageTableToPb(st.Table)
	if err != nil {
		return nil, err
	}
	if tablePb != nil {
		pb.Table = tablePb
	}

	return pb, nil
}

func ExternalLineageObjectFromPb(pb *catalogpb.ExternalLineageObjectPb) (*ExternalLineageObject, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLineageObject{}
	externalMetadataField, err := ExternalLineageExternalMetadataFromPb(pb.ExternalMetadata)
	if err != nil {
		return nil, err
	}
	if externalMetadataField != nil {
		st.ExternalMetadata = externalMetadataField
	}
	modelVersionField, err := ExternalLineageModelVersionFromPb(pb.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionField != nil {
		st.ModelVersion = modelVersionField
	}
	pathField, err := ExternalLineagePathFromPb(pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = pathField
	}
	tableField, err := ExternalLineageTableFromPb(pb.Table)
	if err != nil {
		return nil, err
	}
	if tableField != nil {
		st.Table = tableField
	}

	return st, nil
}

type ExternalLineagePath struct {

	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExternalLineagePath) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineagePath) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalLineagePathToPb(st *ExternalLineagePath) (*catalogpb.ExternalLineagePathPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ExternalLineagePathPb{}
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExternalLineagePathFromPb(pb *catalogpb.ExternalLineagePathPb) (*ExternalLineagePath, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLineagePath{}
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ExternalLineageRelationship struct {
	// List of column relationships between source and target objects.
	// Wire name: 'columns'
	Columns []ColumnRelationship ``
	// Unique identifier of the external lineage relationship.
	// Wire name: 'id'
	Id string ``
	// Key-value properties associated with the external lineage relationship.
	// Wire name: 'properties'
	Properties map[string]string ``
	// Source object of the external lineage relationship.
	// Wire name: 'source'
	Source ExternalLineageObject ``
	// Target object of the external lineage relationship.
	// Wire name: 'target'
	Target          ExternalLineageObject ``
	ForceSendFields []string              `tf:"-"`
}

func (s *ExternalLineageRelationship) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageRelationship) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalLineageRelationshipToPb(st *ExternalLineageRelationship) (*catalogpb.ExternalLineageRelationshipPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ExternalLineageRelationshipPb{}

	var columnsPb []catalogpb.ColumnRelationshipPb
	for _, item := range st.Columns {
		itemPb, err := ColumnRelationshipToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			columnsPb = append(columnsPb, *itemPb)
		}
	}
	pb.Columns = columnsPb
	pb.Id = st.Id
	pb.Properties = st.Properties
	sourcePb, err := ExternalLineageObjectToPb(&st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = *sourcePb
	}
	targetPb, err := ExternalLineageObjectToPb(&st.Target)
	if err != nil {
		return nil, err
	}
	if targetPb != nil {
		pb.Target = *targetPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExternalLineageRelationshipFromPb(pb *catalogpb.ExternalLineageRelationshipPb) (*ExternalLineageRelationship, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLineageRelationship{}

	var columnsField []ColumnRelationship
	for _, itemPb := range pb.Columns {
		item, err := ColumnRelationshipFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			columnsField = append(columnsField, *item)
		}
	}
	st.Columns = columnsField
	st.Id = pb.Id
	st.Properties = pb.Properties
	sourceField, err := ExternalLineageObjectFromPb(&pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = *sourceField
	}
	targetField, err := ExternalLineageObjectFromPb(&pb.Target)
	if err != nil {
		return nil, err
	}
	if targetField != nil {
		st.Target = *targetField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ExternalLineageRelationshipInfo struct {
	// List of column relationships between source and target objects.
	// Wire name: 'columns'
	Columns []ColumnRelationship ``
	// Unique identifier of the external lineage relationship.
	// Wire name: 'id'
	Id string ``
	// Key-value properties associated with the external lineage relationship.
	// Wire name: 'properties'
	Properties map[string]string ``
	// Source object of the external lineage relationship.
	// Wire name: 'source'
	Source ExternalLineageObject ``
	// Target object of the external lineage relationship.
	// Wire name: 'target'
	Target          ExternalLineageObject ``
	ForceSendFields []string              `tf:"-"`
}

func (s *ExternalLineageRelationshipInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageRelationshipInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalLineageRelationshipInfoToPb(st *ExternalLineageRelationshipInfo) (*catalogpb.ExternalLineageRelationshipInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ExternalLineageRelationshipInfoPb{}

	var columnsPb []catalogpb.ColumnRelationshipPb
	for _, item := range st.Columns {
		itemPb, err := ColumnRelationshipToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			columnsPb = append(columnsPb, *itemPb)
		}
	}
	pb.Columns = columnsPb
	pb.Id = st.Id
	pb.Properties = st.Properties
	sourcePb, err := ExternalLineageObjectToPb(&st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = *sourcePb
	}
	targetPb, err := ExternalLineageObjectToPb(&st.Target)
	if err != nil {
		return nil, err
	}
	if targetPb != nil {
		pb.Target = *targetPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExternalLineageRelationshipInfoFromPb(pb *catalogpb.ExternalLineageRelationshipInfoPb) (*ExternalLineageRelationshipInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLineageRelationshipInfo{}

	var columnsField []ColumnRelationship
	for _, itemPb := range pb.Columns {
		item, err := ColumnRelationshipFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			columnsField = append(columnsField, *item)
		}
	}
	st.Columns = columnsField
	st.Id = pb.Id
	st.Properties = pb.Properties
	sourceField, err := ExternalLineageObjectFromPb(&pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = *sourceField
	}
	targetField, err := ExternalLineageObjectFromPb(&pb.Target)
	if err != nil {
		return nil, err
	}
	if targetField != nil {
		st.Target = *targetField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ExternalLineageTable struct {

	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExternalLineageTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalLineageTableToPb(st *ExternalLineageTable) (*catalogpb.ExternalLineageTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ExternalLineageTablePb{}
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExternalLineageTableFromPb(pb *catalogpb.ExternalLineageTablePb) (*ExternalLineageTable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLineageTable{}
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Represents the table information in the lineage event.
type ExternalLineageTableInfo struct {
	// Name of Catalog.
	// Wire name: 'catalog_name'
	CatalogName string ``
	// Timestamp of the lineage event.
	// Wire name: 'event_time'
	EventTime *time.Time ``
	// Name of Table.
	// Wire name: 'name'
	Name string ``
	// Name of Schema.
	// Wire name: 'schema_name'
	SchemaName      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExternalLineageTableInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLineageTableInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalLineageTableInfoToPb(st *ExternalLineageTableInfo) (*catalogpb.ExternalLineageTableInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ExternalLineageTableInfoPb{}
	pb.CatalogName = st.CatalogName
	eventTimePb, err := timestampToPb(st.EventTime)
	if err != nil {
		return nil, err
	}
	if eventTimePb != nil {
		pb.EventTime = *eventTimePb
	}
	pb.Name = st.Name
	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExternalLineageTableInfoFromPb(pb *catalogpb.ExternalLineageTableInfoPb) (*ExternalLineageTableInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLineageTableInfo{}
	st.CatalogName = pb.CatalogName
	eventTimeField, err := timestampFromPb(&pb.EventTime)
	if err != nil {
		return nil, err
	}
	if eventTimeField != nil {
		st.EventTime = eventTimeField
	}
	st.Name = pb.Name
	st.SchemaName = pb.SchemaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ExternalLocationInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool ``
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// Time at which this external location was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of external location creator.
	// Wire name: 'created_by'
	CreatedBy string ``
	// Unique ID of the location's storage credential.
	// Wire name: 'credential_id'
	CredentialId string ``
	// Name of the storage credential used with this location.
	// Wire name: 'credential_name'
	CredentialName string ``
	// Whether to enable file events on this external location.
	// Wire name: 'enable_file_events'
	EnableFileEvents bool ``

	// Wire name: 'encryption_details'
	EncryptionDetails *EncryptionDetails ``
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	// Wire name: 'fallback'
	Fallback bool ``
	// File event queue settings.
	// Wire name: 'file_event_queue'
	FileEventQueue *FileEventQueue ``

	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode ``
	// Unique identifier of metastore hosting the external location.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// Name of the external location.
	// Wire name: 'name'
	Name string ``
	// The owner of the external location.
	// Wire name: 'owner'
	Owner string ``
	// Indicates whether the external location is read-only.
	// Wire name: 'read_only'
	ReadOnly bool ``
	// Time at which external location this was last modified, in epoch
	// milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// Username of user who last modified the external location.
	// Wire name: 'updated_by'
	UpdatedBy string ``
	// Path URL of the external location.
	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExternalLocationInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLocationInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalLocationInfoToPb(st *ExternalLocationInfo) (*catalogpb.ExternalLocationInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ExternalLocationInfoPb{}
	pb.BrowseOnly = st.BrowseOnly
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.CredentialId = st.CredentialId
	pb.CredentialName = st.CredentialName
	pb.EnableFileEvents = st.EnableFileEvents
	encryptionDetailsPb, err := EncryptionDetailsToPb(st.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsPb != nil {
		pb.EncryptionDetails = encryptionDetailsPb
	}
	pb.Fallback = st.Fallback
	fileEventQueuePb, err := FileEventQueueToPb(st.FileEventQueue)
	if err != nil {
		return nil, err
	}
	if fileEventQueuePb != nil {
		pb.FileEventQueue = fileEventQueuePb
	}
	isolationModePb, err := IsolationModeToPb(&st.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}
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

func ExternalLocationInfoFromPb(pb *catalogpb.ExternalLocationInfoPb) (*ExternalLocationInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLocationInfo{}
	st.BrowseOnly = pb.BrowseOnly
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.CredentialId = pb.CredentialId
	st.CredentialName = pb.CredentialName
	st.EnableFileEvents = pb.EnableFileEvents
	encryptionDetailsField, err := EncryptionDetailsFromPb(pb.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsField != nil {
		st.EncryptionDetails = encryptionDetailsField
	}
	st.Fallback = pb.Fallback
	fileEventQueueField, err := FileEventQueueFromPb(pb.FileEventQueue)
	if err != nil {
		return nil, err
	}
	if fileEventQueueField != nil {
		st.FileEventQueue = fileEventQueueField
	}
	isolationModeField, err := IsolationModeFromPb(&pb.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
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

type ExternalMetadata struct {
	// List of columns associated with the external metadata object.
	// Wire name: 'columns'
	Columns []string ``
	// Time at which this external metadata object was created.
	// Wire name: 'create_time'
	CreateTime *time.Time ``
	// Username of external metadata object creator.
	// Wire name: 'created_by'
	CreatedBy string ``
	// User-provided free-form text description.
	// Wire name: 'description'
	Description string ``
	// Type of entity within the external system.
	// Wire name: 'entity_type'
	EntityType string ``
	// Unique identifier of the external metadata object.
	// Wire name: 'id'
	Id string ``
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// Name of the external metadata object.
	// Wire name: 'name'
	Name string ``
	// Owner of the external metadata object.
	// Wire name: 'owner'
	Owner string ``
	// A map of key-value properties attached to the external metadata object.
	// Wire name: 'properties'
	Properties map[string]string ``
	// Type of external system.
	// Wire name: 'system_type'
	SystemType SystemType ``
	// Time at which this external metadata object was last modified.
	// Wire name: 'update_time'
	UpdateTime *time.Time ``
	// Username of user who last modified external metadata object.
	// Wire name: 'updated_by'
	UpdatedBy string ``
	// URL associated with the external metadata object.
	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExternalMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalMetadataToPb(st *ExternalMetadata) (*catalogpb.ExternalMetadataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ExternalMetadataPb{}
	pb.Columns = st.Columns
	createTimePb, err := timestampToPb(st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}
	pb.CreatedBy = st.CreatedBy
	pb.Description = st.Description
	pb.EntityType = st.EntityType
	pb.Id = st.Id
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	pb.Properties = st.Properties
	systemTypePb, err := SystemTypeToPb(&st.SystemType)
	if err != nil {
		return nil, err
	}
	if systemTypePb != nil {
		pb.SystemType = *systemTypePb
	}
	updateTimePb, err := timestampToPb(st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}
	pb.UpdatedBy = st.UpdatedBy
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExternalMetadataFromPb(pb *catalogpb.ExternalMetadataPb) (*ExternalMetadata, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalMetadata{}
	st.Columns = pb.Columns
	createTimeField, err := timestampFromPb(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = createTimeField
	}
	st.CreatedBy = pb.CreatedBy
	st.Description = pb.Description
	st.EntityType = pb.EntityType
	st.Id = pb.Id
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.Properties = pb.Properties
	systemTypeField, err := SystemTypeFromPb(&pb.SystemType)
	if err != nil {
		return nil, err
	}
	if systemTypeField != nil {
		st.SystemType = *systemTypeField
	}
	updateTimeField, err := timestampFromPb(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = updateTimeField
	}
	st.UpdatedBy = pb.UpdatedBy
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Detailed status of an online table. Shown if the online table is in the
// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
type FailedStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may only be partially synced to the online
	// table. Only populated if the table is still online and available for
	// serving.
	// Wire name: 'last_processed_commit_version'
	LastProcessedCommitVersion int64 ``
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table. Only populated if the table is still online
	// and available for serving.
	// Wire name: 'timestamp'
	Timestamp       *time.Time ``
	ForceSendFields []string   `tf:"-"`
}

func (s *FailedStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FailedStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func FailedStatusToPb(st *FailedStatus) (*catalogpb.FailedStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.FailedStatusPb{}
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion
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

func FailedStatusFromPb(pb *catalogpb.FailedStatusPb) (*FailedStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FailedStatus{}
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
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

type FileEventQueue struct {

	// Wire name: 'managed_aqs'
	ManagedAqs *AzureQueueStorage ``

	// Wire name: 'managed_pubsub'
	ManagedPubsub *GcpPubsub ``

	// Wire name: 'managed_sqs'
	ManagedSqs *AwsSqsQueue ``

	// Wire name: 'provided_aqs'
	ProvidedAqs *AzureQueueStorage ``

	// Wire name: 'provided_pubsub'
	ProvidedPubsub *GcpPubsub ``

	// Wire name: 'provided_sqs'
	ProvidedSqs *AwsSqsQueue ``
}

func FileEventQueueToPb(st *FileEventQueue) (*catalogpb.FileEventQueuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.FileEventQueuePb{}
	managedAqsPb, err := AzureQueueStorageToPb(st.ManagedAqs)
	if err != nil {
		return nil, err
	}
	if managedAqsPb != nil {
		pb.ManagedAqs = managedAqsPb
	}
	managedPubsubPb, err := GcpPubsubToPb(st.ManagedPubsub)
	if err != nil {
		return nil, err
	}
	if managedPubsubPb != nil {
		pb.ManagedPubsub = managedPubsubPb
	}
	managedSqsPb, err := AwsSqsQueueToPb(st.ManagedSqs)
	if err != nil {
		return nil, err
	}
	if managedSqsPb != nil {
		pb.ManagedSqs = managedSqsPb
	}
	providedAqsPb, err := AzureQueueStorageToPb(st.ProvidedAqs)
	if err != nil {
		return nil, err
	}
	if providedAqsPb != nil {
		pb.ProvidedAqs = providedAqsPb
	}
	providedPubsubPb, err := GcpPubsubToPb(st.ProvidedPubsub)
	if err != nil {
		return nil, err
	}
	if providedPubsubPb != nil {
		pb.ProvidedPubsub = providedPubsubPb
	}
	providedSqsPb, err := AwsSqsQueueToPb(st.ProvidedSqs)
	if err != nil {
		return nil, err
	}
	if providedSqsPb != nil {
		pb.ProvidedSqs = providedSqsPb
	}

	return pb, nil
}

func FileEventQueueFromPb(pb *catalogpb.FileEventQueuePb) (*FileEventQueue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileEventQueue{}
	managedAqsField, err := AzureQueueStorageFromPb(pb.ManagedAqs)
	if err != nil {
		return nil, err
	}
	if managedAqsField != nil {
		st.ManagedAqs = managedAqsField
	}
	managedPubsubField, err := GcpPubsubFromPb(pb.ManagedPubsub)
	if err != nil {
		return nil, err
	}
	if managedPubsubField != nil {
		st.ManagedPubsub = managedPubsubField
	}
	managedSqsField, err := AwsSqsQueueFromPb(pb.ManagedSqs)
	if err != nil {
		return nil, err
	}
	if managedSqsField != nil {
		st.ManagedSqs = managedSqsField
	}
	providedAqsField, err := AzureQueueStorageFromPb(pb.ProvidedAqs)
	if err != nil {
		return nil, err
	}
	if providedAqsField != nil {
		st.ProvidedAqs = providedAqsField
	}
	providedPubsubField, err := GcpPubsubFromPb(pb.ProvidedPubsub)
	if err != nil {
		return nil, err
	}
	if providedPubsubField != nil {
		st.ProvidedPubsub = providedPubsubField
	}
	providedSqsField, err := AwsSqsQueueFromPb(pb.ProvidedSqs)
	if err != nil {
		return nil, err
	}
	if providedSqsField != nil {
		st.ProvidedSqs = providedSqsField
	}

	return st, nil
}

type ForeignKeyConstraint struct {
	// Column names for this constraint.
	// Wire name: 'child_columns'
	ChildColumns []string ``
	// The name of the constraint.
	// Wire name: 'name'
	Name string ``
	// Column names for this constraint.
	// Wire name: 'parent_columns'
	ParentColumns []string ``
	// The full name of the parent constraint.
	// Wire name: 'parent_table'
	ParentTable string ``
	// True if the constraint is RELY, false or unset if NORELY.
	// Wire name: 'rely'
	Rely            bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *ForeignKeyConstraint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ForeignKeyConstraint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ForeignKeyConstraintToPb(st *ForeignKeyConstraint) (*catalogpb.ForeignKeyConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ForeignKeyConstraintPb{}
	pb.ChildColumns = st.ChildColumns
	pb.Name = st.Name
	pb.ParentColumns = st.ParentColumns
	pb.ParentTable = st.ParentTable
	pb.Rely = st.Rely

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ForeignKeyConstraintFromPb(pb *catalogpb.ForeignKeyConstraintPb) (*ForeignKeyConstraint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForeignKeyConstraint{}
	st.ChildColumns = pb.ChildColumns
	st.Name = pb.Name
	st.ParentColumns = pb.ParentColumns
	st.ParentTable = pb.ParentTable
	st.Rely = pb.Rely

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// A function that is dependent on a SQL object.
type FunctionDependency struct {
	// Full name of the dependent function, in the form of
	// __catalog_name__.__schema_name__.__function_name__.
	// Wire name: 'function_full_name'
	FunctionFullName string ``
}

func FunctionDependencyToPb(st *FunctionDependency) (*catalogpb.FunctionDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.FunctionDependencyPb{}
	pb.FunctionFullName = st.FunctionFullName

	return pb, nil
}

func FunctionDependencyFromPb(pb *catalogpb.FunctionDependencyPb) (*FunctionDependency, error) {
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
	BrowseOnly bool ``
	// Name of parent catalog.
	// Wire name: 'catalog_name'
	CatalogName string ``
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// Time at which this function was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of function creator.
	// Wire name: 'created_by'
	CreatedBy string ``
	// Scalar function return data type.
	// Wire name: 'data_type'
	DataType ColumnTypeName ``
	// External function language.
	// Wire name: 'external_language'
	ExternalLanguage string ``
	// External function name.
	// Wire name: 'external_name'
	ExternalName string ``
	// Pretty printed function data type.
	// Wire name: 'full_data_type'
	FullDataType string ``
	// Full name of function, in form of
	// __catalog_name__.__schema_name__.__function__name__
	// Wire name: 'full_name'
	FullName string ``
	// Id of Function, relative to parent schema.
	// Wire name: 'function_id'
	FunctionId string ``

	// Wire name: 'input_params'
	InputParams *FunctionParameterInfos ``
	// Whether the function is deterministic.
	// Wire name: 'is_deterministic'
	IsDeterministic bool ``
	// Function null call.
	// Wire name: 'is_null_call'
	IsNullCall bool ``
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// Name of function, relative to parent schema.
	// Wire name: 'name'
	Name string ``
	// Username of current owner of function.
	// Wire name: 'owner'
	Owner string ``
	// Function parameter style. **S** is the value for SQL.
	// Wire name: 'parameter_style'
	ParameterStyle FunctionInfoParameterStyle ``
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	// Wire name: 'properties'
	Properties string ``
	// Table function return parameters.
	// Wire name: 'return_params'
	ReturnParams *FunctionParameterInfos ``
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	// Wire name: 'routine_body'
	RoutineBody FunctionInfoRoutineBody ``
	// Function body.
	// Wire name: 'routine_definition'
	RoutineDefinition string ``
	// Function dependencies.
	// Wire name: 'routine_dependencies'
	RoutineDependencies *DependencyList ``
	// Name of parent schema relative to its parent catalog.
	// Wire name: 'schema_name'
	SchemaName string ``
	// Function security type.
	// Wire name: 'security_type'
	SecurityType FunctionInfoSecurityType ``
	// Specific name of the function; Reserved for future use.
	// Wire name: 'specific_name'
	SpecificName string ``
	// Function SQL data access.
	// Wire name: 'sql_data_access'
	SqlDataAccess FunctionInfoSqlDataAccess ``
	// List of schemes whose objects can be referenced without qualification.
	// Wire name: 'sql_path'
	SqlPath string ``
	// Time at which this function was created, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// Username of user who last modified function.
	// Wire name: 'updated_by'
	UpdatedBy       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *FunctionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FunctionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func FunctionInfoToPb(st *FunctionInfo) (*catalogpb.FunctionInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.FunctionInfoPb{}
	pb.BrowseOnly = st.BrowseOnly
	pb.CatalogName = st.CatalogName
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	dataTypePb, err := ColumnTypeNameToPb(&st.DataType)
	if err != nil {
		return nil, err
	}
	if dataTypePb != nil {
		pb.DataType = *dataTypePb
	}
	pb.ExternalLanguage = st.ExternalLanguage
	pb.ExternalName = st.ExternalName
	pb.FullDataType = st.FullDataType
	pb.FullName = st.FullName
	pb.FunctionId = st.FunctionId
	inputParamsPb, err := FunctionParameterInfosToPb(st.InputParams)
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
	parameterStylePb, err := FunctionInfoParameterStyleToPb(&st.ParameterStyle)
	if err != nil {
		return nil, err
	}
	if parameterStylePb != nil {
		pb.ParameterStyle = *parameterStylePb
	}
	pb.Properties = st.Properties
	returnParamsPb, err := FunctionParameterInfosToPb(st.ReturnParams)
	if err != nil {
		return nil, err
	}
	if returnParamsPb != nil {
		pb.ReturnParams = returnParamsPb
	}
	routineBodyPb, err := FunctionInfoRoutineBodyToPb(&st.RoutineBody)
	if err != nil {
		return nil, err
	}
	if routineBodyPb != nil {
		pb.RoutineBody = *routineBodyPb
	}
	pb.RoutineDefinition = st.RoutineDefinition
	routineDependenciesPb, err := DependencyListToPb(st.RoutineDependencies)
	if err != nil {
		return nil, err
	}
	if routineDependenciesPb != nil {
		pb.RoutineDependencies = routineDependenciesPb
	}
	pb.SchemaName = st.SchemaName
	securityTypePb, err := FunctionInfoSecurityTypeToPb(&st.SecurityType)
	if err != nil {
		return nil, err
	}
	if securityTypePb != nil {
		pb.SecurityType = *securityTypePb
	}
	pb.SpecificName = st.SpecificName
	sqlDataAccessPb, err := FunctionInfoSqlDataAccessToPb(&st.SqlDataAccess)
	if err != nil {
		return nil, err
	}
	if sqlDataAccessPb != nil {
		pb.SqlDataAccess = *sqlDataAccessPb
	}
	pb.SqlPath = st.SqlPath
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func FunctionInfoFromPb(pb *catalogpb.FunctionInfoPb) (*FunctionInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionInfo{}
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	dataTypeField, err := ColumnTypeNameFromPb(&pb.DataType)
	if err != nil {
		return nil, err
	}
	if dataTypeField != nil {
		st.DataType = *dataTypeField
	}
	st.ExternalLanguage = pb.ExternalLanguage
	st.ExternalName = pb.ExternalName
	st.FullDataType = pb.FullDataType
	st.FullName = pb.FullName
	st.FunctionId = pb.FunctionId
	inputParamsField, err := FunctionParameterInfosFromPb(pb.InputParams)
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
	parameterStyleField, err := FunctionInfoParameterStyleFromPb(&pb.ParameterStyle)
	if err != nil {
		return nil, err
	}
	if parameterStyleField != nil {
		st.ParameterStyle = *parameterStyleField
	}
	st.Properties = pb.Properties
	returnParamsField, err := FunctionParameterInfosFromPb(pb.ReturnParams)
	if err != nil {
		return nil, err
	}
	if returnParamsField != nil {
		st.ReturnParams = returnParamsField
	}
	routineBodyField, err := FunctionInfoRoutineBodyFromPb(&pb.RoutineBody)
	if err != nil {
		return nil, err
	}
	if routineBodyField != nil {
		st.RoutineBody = *routineBodyField
	}
	st.RoutineDefinition = pb.RoutineDefinition
	routineDependenciesField, err := DependencyListFromPb(pb.RoutineDependencies)
	if err != nil {
		return nil, err
	}
	if routineDependenciesField != nil {
		st.RoutineDependencies = routineDependenciesField
	}
	st.SchemaName = pb.SchemaName
	securityTypeField, err := FunctionInfoSecurityTypeFromPb(&pb.SecurityType)
	if err != nil {
		return nil, err
	}
	if securityTypeField != nil {
		st.SecurityType = *securityTypeField
	}
	st.SpecificName = pb.SpecificName
	sqlDataAccessField, err := FunctionInfoSqlDataAccessFromPb(&pb.SqlDataAccess)
	if err != nil {
		return nil, err
	}
	if sqlDataAccessField != nil {
		st.SqlDataAccess = *sqlDataAccessField
	}
	st.SqlPath = pb.SqlPath
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Function parameter style. **S** is the value for SQL.
type FunctionInfoParameterStyle string

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

// Values returns all possible values for FunctionInfoParameterStyle.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionInfoParameterStyle) Values() []FunctionInfoParameterStyle {
	return []FunctionInfoParameterStyle{
		FunctionInfoParameterStyleS,
	}
}

// Type always returns FunctionInfoParameterStyle to satisfy [pflag.Value] interface
func (f *FunctionInfoParameterStyle) Type() string {
	return "FunctionInfoParameterStyle"
}

func FunctionInfoParameterStyleToPb(st *FunctionInfoParameterStyle) (*catalogpb.FunctionInfoParameterStylePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.FunctionInfoParameterStylePb(*st)
	return &pb, nil
}

func FunctionInfoParameterStyleFromPb(pb *catalogpb.FunctionInfoParameterStylePb) (*FunctionInfoParameterStyle, error) {
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

// Values returns all possible values for FunctionInfoRoutineBody.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionInfoRoutineBody) Values() []FunctionInfoRoutineBody {
	return []FunctionInfoRoutineBody{
		FunctionInfoRoutineBodyExternal,
		FunctionInfoRoutineBodySql,
	}
}

// Type always returns FunctionInfoRoutineBody to satisfy [pflag.Value] interface
func (f *FunctionInfoRoutineBody) Type() string {
	return "FunctionInfoRoutineBody"
}

func FunctionInfoRoutineBodyToPb(st *FunctionInfoRoutineBody) (*catalogpb.FunctionInfoRoutineBodyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.FunctionInfoRoutineBodyPb(*st)
	return &pb, nil
}

func FunctionInfoRoutineBodyFromPb(pb *catalogpb.FunctionInfoRoutineBodyPb) (*FunctionInfoRoutineBody, error) {
	if pb == nil {
		return nil, nil
	}
	st := FunctionInfoRoutineBody(*pb)
	return &st, nil
}

// The security type of the function.
type FunctionInfoSecurityType string

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

// Values returns all possible values for FunctionInfoSecurityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionInfoSecurityType) Values() []FunctionInfoSecurityType {
	return []FunctionInfoSecurityType{
		FunctionInfoSecurityTypeDefiner,
	}
}

// Type always returns FunctionInfoSecurityType to satisfy [pflag.Value] interface
func (f *FunctionInfoSecurityType) Type() string {
	return "FunctionInfoSecurityType"
}

func FunctionInfoSecurityTypeToPb(st *FunctionInfoSecurityType) (*catalogpb.FunctionInfoSecurityTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.FunctionInfoSecurityTypePb(*st)
	return &pb, nil
}

func FunctionInfoSecurityTypeFromPb(pb *catalogpb.FunctionInfoSecurityTypePb) (*FunctionInfoSecurityType, error) {
	if pb == nil {
		return nil, nil
	}
	st := FunctionInfoSecurityType(*pb)
	return &st, nil
}

// Function SQL data access.
type FunctionInfoSqlDataAccess string

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

// Values returns all possible values for FunctionInfoSqlDataAccess.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionInfoSqlDataAccess) Values() []FunctionInfoSqlDataAccess {
	return []FunctionInfoSqlDataAccess{
		FunctionInfoSqlDataAccessContainsSql,
		FunctionInfoSqlDataAccessNoSql,
		FunctionInfoSqlDataAccessReadsSqlData,
	}
}

// Type always returns FunctionInfoSqlDataAccess to satisfy [pflag.Value] interface
func (f *FunctionInfoSqlDataAccess) Type() string {
	return "FunctionInfoSqlDataAccess"
}

func FunctionInfoSqlDataAccessToPb(st *FunctionInfoSqlDataAccess) (*catalogpb.FunctionInfoSqlDataAccessPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.FunctionInfoSqlDataAccessPb(*st)
	return &pb, nil
}

func FunctionInfoSqlDataAccessFromPb(pb *catalogpb.FunctionInfoSqlDataAccessPb) (*FunctionInfoSqlDataAccess, error) {
	if pb == nil {
		return nil, nil
	}
	st := FunctionInfoSqlDataAccess(*pb)
	return &st, nil
}

type FunctionParameterInfo struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// Name of parameter.
	// Wire name: 'name'
	Name string ``
	// Default value of the parameter.
	// Wire name: 'parameter_default'
	ParameterDefault string ``

	// Wire name: 'parameter_mode'
	ParameterMode FunctionParameterMode ``

	// Wire name: 'parameter_type'
	ParameterType FunctionParameterType ``
	// Ordinal position of column (starting at position 0).
	// Wire name: 'position'
	Position int ``
	// Format of IntervalType.
	// Wire name: 'type_interval_type'
	TypeIntervalType string ``
	// Full data type spec, JSON-serialized.
	// Wire name: 'type_json'
	TypeJson string ``

	// Wire name: 'type_name'
	TypeName ColumnTypeName ``
	// Digits of precision; required on Create for DecimalTypes.
	// Wire name: 'type_precision'
	TypePrecision int ``
	// Digits to right of decimal; Required on Create for DecimalTypes.
	// Wire name: 'type_scale'
	TypeScale int ``
	// Full data type spec, SQL/catalogString text.
	// Wire name: 'type_text'
	TypeText        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *FunctionParameterInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FunctionParameterInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func FunctionParameterInfoToPb(st *FunctionParameterInfo) (*catalogpb.FunctionParameterInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.FunctionParameterInfoPb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.ParameterDefault = st.ParameterDefault
	parameterModePb, err := FunctionParameterModeToPb(&st.ParameterMode)
	if err != nil {
		return nil, err
	}
	if parameterModePb != nil {
		pb.ParameterMode = *parameterModePb
	}
	parameterTypePb, err := FunctionParameterTypeToPb(&st.ParameterType)
	if err != nil {
		return nil, err
	}
	if parameterTypePb != nil {
		pb.ParameterType = *parameterTypePb
	}
	pb.Position = st.Position
	pb.TypeIntervalType = st.TypeIntervalType
	pb.TypeJson = st.TypeJson
	typeNamePb, err := ColumnTypeNameToPb(&st.TypeName)
	if err != nil {
		return nil, err
	}
	if typeNamePb != nil {
		pb.TypeName = *typeNamePb
	}
	pb.TypePrecision = st.TypePrecision
	pb.TypeScale = st.TypeScale
	pb.TypeText = st.TypeText

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func FunctionParameterInfoFromPb(pb *catalogpb.FunctionParameterInfoPb) (*FunctionParameterInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionParameterInfo{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.ParameterDefault = pb.ParameterDefault
	parameterModeField, err := FunctionParameterModeFromPb(&pb.ParameterMode)
	if err != nil {
		return nil, err
	}
	if parameterModeField != nil {
		st.ParameterMode = *parameterModeField
	}
	parameterTypeField, err := FunctionParameterTypeFromPb(&pb.ParameterType)
	if err != nil {
		return nil, err
	}
	if parameterTypeField != nil {
		st.ParameterType = *parameterTypeField
	}
	st.Position = pb.Position
	st.TypeIntervalType = pb.TypeIntervalType
	st.TypeJson = pb.TypeJson
	typeNameField, err := ColumnTypeNameFromPb(&pb.TypeName)
	if err != nil {
		return nil, err
	}
	if typeNameField != nil {
		st.TypeName = *typeNameField
	}
	st.TypePrecision = pb.TypePrecision
	st.TypeScale = pb.TypeScale
	st.TypeText = pb.TypeText

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type FunctionParameterInfos struct {
	// The array of __FunctionParameterInfo__ definitions of the function's
	// parameters.
	// Wire name: 'parameters'
	Parameters []FunctionParameterInfo ``
}

func FunctionParameterInfosToPb(st *FunctionParameterInfos) (*catalogpb.FunctionParameterInfosPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.FunctionParameterInfosPb{}

	var parametersPb []catalogpb.FunctionParameterInfoPb
	for _, item := range st.Parameters {
		itemPb, err := FunctionParameterInfoToPb(&item)
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

func FunctionParameterInfosFromPb(pb *catalogpb.FunctionParameterInfosPb) (*FunctionParameterInfos, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionParameterInfos{}

	var parametersField []FunctionParameterInfo
	for _, itemPb := range pb.Parameters {
		item, err := FunctionParameterInfoFromPb(&itemPb)
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

// Values returns all possible values for FunctionParameterMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionParameterMode) Values() []FunctionParameterMode {
	return []FunctionParameterMode{
		FunctionParameterModeIn,
	}
}

// Type always returns FunctionParameterMode to satisfy [pflag.Value] interface
func (f *FunctionParameterMode) Type() string {
	return "FunctionParameterMode"
}

func FunctionParameterModeToPb(st *FunctionParameterMode) (*catalogpb.FunctionParameterModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.FunctionParameterModePb(*st)
	return &pb, nil
}

func FunctionParameterModeFromPb(pb *catalogpb.FunctionParameterModePb) (*FunctionParameterMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := FunctionParameterMode(*pb)
	return &st, nil
}

// The type of function parameter.
type FunctionParameterType string

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

// Values returns all possible values for FunctionParameterType.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionParameterType) Values() []FunctionParameterType {
	return []FunctionParameterType{
		FunctionParameterTypeColumn,
		FunctionParameterTypeParam,
	}
}

// Type always returns FunctionParameterType to satisfy [pflag.Value] interface
func (f *FunctionParameterType) Type() string {
	return "FunctionParameterType"
}

func FunctionParameterTypeToPb(st *FunctionParameterType) (*catalogpb.FunctionParameterTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.FunctionParameterTypePb(*st)
	return &pb, nil
}

func FunctionParameterTypeFromPb(pb *catalogpb.FunctionParameterTypePb) (*FunctionParameterType, error) {
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
	OauthToken      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GcpOauthToken) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpOauthToken) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GcpOauthTokenToPb(st *GcpOauthToken) (*catalogpb.GcpOauthTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GcpOauthTokenPb{}
	pb.OauthToken = st.OauthToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GcpOauthTokenFromPb(pb *catalogpb.GcpOauthTokenPb) (*GcpOauthToken, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpOauthToken{}
	st.OauthToken = pb.OauthToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GcpPubsub struct {
	// Unique identifier included in the name of file events managed cloud
	// resources.
	// Wire name: 'managed_resource_id'
	ManagedResourceId string ``
	// The Pub/Sub subscription name in the format
	// projects/{project}/subscriptions/{subscription name} Required for
	// provided_pubsub.
	// Wire name: 'subscription_name'
	SubscriptionName string   ``
	ForceSendFields  []string `tf:"-"`
}

func (s *GcpPubsub) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpPubsub) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GcpPubsubToPb(st *GcpPubsub) (*catalogpb.GcpPubsubPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GcpPubsubPb{}
	pb.ManagedResourceId = st.ManagedResourceId
	pb.SubscriptionName = st.SubscriptionName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GcpPubsubFromPb(pb *catalogpb.GcpPubsubPb) (*GcpPubsub, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpPubsub{}
	st.ManagedResourceId = pb.ManagedResourceId
	st.SubscriptionName = pb.SubscriptionName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// The Azure cloud options to customize the requested temporary credential
type GenerateTemporaryServiceCredentialAzureOptions struct {
	// The resources to which the temporary Azure credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://learn.microsoft.com/python/api/azure-core/azure.core.credentials.tokencredential?view=azure-python)
	// Wire name: 'resources'
	Resources []string ``
}

func GenerateTemporaryServiceCredentialAzureOptionsToPb(st *GenerateTemporaryServiceCredentialAzureOptions) (*catalogpb.GenerateTemporaryServiceCredentialAzureOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GenerateTemporaryServiceCredentialAzureOptionsPb{}
	pb.Resources = st.Resources

	return pb, nil
}

func GenerateTemporaryServiceCredentialAzureOptionsFromPb(pb *catalogpb.GenerateTemporaryServiceCredentialAzureOptionsPb) (*GenerateTemporaryServiceCredentialAzureOptions, error) {
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
	Scopes []string ``
}

func GenerateTemporaryServiceCredentialGcpOptionsToPb(st *GenerateTemporaryServiceCredentialGcpOptions) (*catalogpb.GenerateTemporaryServiceCredentialGcpOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GenerateTemporaryServiceCredentialGcpOptionsPb{}
	pb.Scopes = st.Scopes

	return pb, nil
}

func GenerateTemporaryServiceCredentialGcpOptionsFromPb(pb *catalogpb.GenerateTemporaryServiceCredentialGcpOptionsPb) (*GenerateTemporaryServiceCredentialGcpOptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateTemporaryServiceCredentialGcpOptions{}
	st.Scopes = pb.Scopes

	return st, nil
}

type GenerateTemporaryServiceCredentialRequest struct {

	// Wire name: 'azure_options'
	AzureOptions *GenerateTemporaryServiceCredentialAzureOptions ``
	// The name of the service credential used to generate a temporary
	// credential
	// Wire name: 'credential_name'
	CredentialName string ``

	// Wire name: 'gcp_options'
	GcpOptions *GenerateTemporaryServiceCredentialGcpOptions ``
}

func GenerateTemporaryServiceCredentialRequestToPb(st *GenerateTemporaryServiceCredentialRequest) (*catalogpb.GenerateTemporaryServiceCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GenerateTemporaryServiceCredentialRequestPb{}
	azureOptionsPb, err := GenerateTemporaryServiceCredentialAzureOptionsToPb(st.AzureOptions)
	if err != nil {
		return nil, err
	}
	if azureOptionsPb != nil {
		pb.AzureOptions = azureOptionsPb
	}
	pb.CredentialName = st.CredentialName
	gcpOptionsPb, err := GenerateTemporaryServiceCredentialGcpOptionsToPb(st.GcpOptions)
	if err != nil {
		return nil, err
	}
	if gcpOptionsPb != nil {
		pb.GcpOptions = gcpOptionsPb
	}

	return pb, nil
}

func GenerateTemporaryServiceCredentialRequestFromPb(pb *catalogpb.GenerateTemporaryServiceCredentialRequestPb) (*GenerateTemporaryServiceCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateTemporaryServiceCredentialRequest{}
	azureOptionsField, err := GenerateTemporaryServiceCredentialAzureOptionsFromPb(pb.AzureOptions)
	if err != nil {
		return nil, err
	}
	if azureOptionsField != nil {
		st.AzureOptions = azureOptionsField
	}
	st.CredentialName = pb.CredentialName
	gcpOptionsField, err := GenerateTemporaryServiceCredentialGcpOptionsFromPb(pb.GcpOptions)
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
	Operation TableOperation ``
	// UUID of the table to read or write.
	// Wire name: 'table_id'
	TableId         string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GenerateTemporaryTableCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenerateTemporaryTableCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GenerateTemporaryTableCredentialRequestToPb(st *GenerateTemporaryTableCredentialRequest) (*catalogpb.GenerateTemporaryTableCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GenerateTemporaryTableCredentialRequestPb{}
	operationPb, err := TableOperationToPb(&st.Operation)
	if err != nil {
		return nil, err
	}
	if operationPb != nil {
		pb.Operation = *operationPb
	}
	pb.TableId = st.TableId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GenerateTemporaryTableCredentialRequestFromPb(pb *catalogpb.GenerateTemporaryTableCredentialRequestPb) (*GenerateTemporaryTableCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateTemporaryTableCredentialRequest{}
	operationField, err := TableOperationFromPb(&pb.Operation)
	if err != nil {
		return nil, err
	}
	if operationField != nil {
		st.Operation = *operationField
	}
	st.TableId = pb.TableId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GenerateTemporaryTableCredentialResponse struct {

	// Wire name: 'aws_temp_credentials'
	AwsTempCredentials *AwsCredentials ``

	// Wire name: 'azure_aad'
	AzureAad *AzureActiveDirectoryToken ``

	// Wire name: 'azure_user_delegation_sas'
	AzureUserDelegationSas *AzureUserDelegationSas ``
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	// Wire name: 'expiration_time'
	ExpirationTime int64 ``

	// Wire name: 'gcp_oauth_token'
	GcpOauthToken *GcpOauthToken ``

	// Wire name: 'r2_temp_credentials'
	R2TempCredentials *R2Credentials ``
	// The URL of the storage path accessible by the temporary credential.
	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GenerateTemporaryTableCredentialResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenerateTemporaryTableCredentialResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GenerateTemporaryTableCredentialResponseToPb(st *GenerateTemporaryTableCredentialResponse) (*catalogpb.GenerateTemporaryTableCredentialResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GenerateTemporaryTableCredentialResponsePb{}
	awsTempCredentialsPb, err := AwsCredentialsToPb(st.AwsTempCredentials)
	if err != nil {
		return nil, err
	}
	if awsTempCredentialsPb != nil {
		pb.AwsTempCredentials = awsTempCredentialsPb
	}
	azureAadPb, err := AzureActiveDirectoryTokenToPb(st.AzureAad)
	if err != nil {
		return nil, err
	}
	if azureAadPb != nil {
		pb.AzureAad = azureAadPb
	}
	azureUserDelegationSasPb, err := AzureUserDelegationSasToPb(st.AzureUserDelegationSas)
	if err != nil {
		return nil, err
	}
	if azureUserDelegationSasPb != nil {
		pb.AzureUserDelegationSas = azureUserDelegationSasPb
	}
	pb.ExpirationTime = st.ExpirationTime
	gcpOauthTokenPb, err := GcpOauthTokenToPb(st.GcpOauthToken)
	if err != nil {
		return nil, err
	}
	if gcpOauthTokenPb != nil {
		pb.GcpOauthToken = gcpOauthTokenPb
	}
	r2TempCredentialsPb, err := R2CredentialsToPb(st.R2TempCredentials)
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

func GenerateTemporaryTableCredentialResponseFromPb(pb *catalogpb.GenerateTemporaryTableCredentialResponsePb) (*GenerateTemporaryTableCredentialResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateTemporaryTableCredentialResponse{}
	awsTempCredentialsField, err := AwsCredentialsFromPb(pb.AwsTempCredentials)
	if err != nil {
		return nil, err
	}
	if awsTempCredentialsField != nil {
		st.AwsTempCredentials = awsTempCredentialsField
	}
	azureAadField, err := AzureActiveDirectoryTokenFromPb(pb.AzureAad)
	if err != nil {
		return nil, err
	}
	if azureAadField != nil {
		st.AzureAad = azureAadField
	}
	azureUserDelegationSasField, err := AzureUserDelegationSasFromPb(pb.AzureUserDelegationSas)
	if err != nil {
		return nil, err
	}
	if azureUserDelegationSasField != nil {
		st.AzureUserDelegationSas = azureUserDelegationSasField
	}
	st.ExpirationTime = pb.ExpirationTime
	gcpOauthTokenField, err := GcpOauthTokenFromPb(pb.GcpOauthToken)
	if err != nil {
		return nil, err
	}
	if gcpOauthTokenField != nil {
		st.GcpOauthToken = gcpOauthTokenField
	}
	r2TempCredentialsField, err := R2CredentialsFromPb(pb.R2TempCredentials)
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

type GetAccountMetastoreAssignmentRequest struct {
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func GetAccountMetastoreAssignmentRequestToPb(st *GetAccountMetastoreAssignmentRequest) (*catalogpb.GetAccountMetastoreAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetAccountMetastoreAssignmentRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func GetAccountMetastoreAssignmentRequestFromPb(pb *catalogpb.GetAccountMetastoreAssignmentRequestPb) (*GetAccountMetastoreAssignmentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountMetastoreAssignmentRequest{}
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

type GetAccountMetastoreRequest struct {
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
}

func GetAccountMetastoreRequestToPb(st *GetAccountMetastoreRequest) (*catalogpb.GetAccountMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetAccountMetastoreRequestPb{}
	pb.MetastoreId = st.MetastoreId

	return pb, nil
}

func GetAccountMetastoreRequestFromPb(pb *catalogpb.GetAccountMetastoreRequestPb) (*GetAccountMetastoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountMetastoreRequest{}
	st.MetastoreId = pb.MetastoreId

	return st, nil
}

type GetAccountStorageCredentialRequest struct {
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// Name of the storage credential.
	// Wire name: 'storage_credential_name'
	StorageCredentialName string `tf:"-"`
}

func GetAccountStorageCredentialRequestToPb(st *GetAccountStorageCredentialRequest) (*catalogpb.GetAccountStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetAccountStorageCredentialRequestPb{}
	pb.MetastoreId = st.MetastoreId
	pb.StorageCredentialName = st.StorageCredentialName

	return pb, nil
}

func GetAccountStorageCredentialRequestFromPb(pb *catalogpb.GetAccountStorageCredentialRequestPb) (*GetAccountStorageCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountStorageCredentialRequest{}
	st.MetastoreId = pb.MetastoreId
	st.StorageCredentialName = pb.StorageCredentialName

	return st, nil
}

type GetArtifactAllowlistRequest struct {
	// The artifact type of the allowlist.
	// Wire name: 'artifact_type'
	ArtifactType ArtifactType `tf:"-"`
}

func GetArtifactAllowlistRequestToPb(st *GetArtifactAllowlistRequest) (*catalogpb.GetArtifactAllowlistRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetArtifactAllowlistRequestPb{}
	artifactTypePb, err := ArtifactTypeToPb(&st.ArtifactType)
	if err != nil {
		return nil, err
	}
	if artifactTypePb != nil {
		pb.ArtifactType = *artifactTypePb
	}

	return pb, nil
}

func GetArtifactAllowlistRequestFromPb(pb *catalogpb.GetArtifactAllowlistRequestPb) (*GetArtifactAllowlistRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetArtifactAllowlistRequest{}
	artifactTypeField, err := ArtifactTypeFromPb(&pb.ArtifactType)
	if err != nil {
		return nil, err
	}
	if artifactTypeField != nil {
		st.ArtifactType = *artifactTypeField
	}

	return st, nil
}

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
	// The type of the securable to bind to a workspace (catalog,
	// storage_credential, credential, or external_location).
	// Wire name: 'securable_type'
	SecurableType   string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *GetBindingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetBindingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetBindingsRequestToPb(st *GetBindingsRequest) (*catalogpb.GetBindingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetBindingsRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.SecurableName = st.SecurableName
	pb.SecurableType = st.SecurableType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetBindingsRequestFromPb(pb *catalogpb.GetBindingsRequestPb) (*GetBindingsRequest, error) {
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
	IncludeAliases  bool     `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *GetByAliasRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetByAliasRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetByAliasRequestToPb(st *GetByAliasRequest) (*catalogpb.GetByAliasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetByAliasRequestPb{}
	pb.Alias = st.Alias
	pb.FullName = st.FullName
	pb.IncludeAliases = st.IncludeAliases

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetByAliasRequestFromPb(pb *catalogpb.GetByAliasRequestPb) (*GetByAliasRequest, error) {
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

type GetCatalogRequest struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// The name of the catalog.
	// Wire name: 'name'
	Name            string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *GetCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCatalogRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetCatalogRequestToPb(st *GetCatalogRequest) (*catalogpb.GetCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetCatalogRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetCatalogRequestFromPb(pb *catalogpb.GetCatalogRequestPb) (*GetCatalogRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCatalogRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetCatalogWorkspaceBindingsResponse struct {
	// A list of workspace IDs
	// Wire name: 'workspaces'
	Workspaces []int64 ``
}

func GetCatalogWorkspaceBindingsResponseToPb(st *GetCatalogWorkspaceBindingsResponse) (*catalogpb.GetCatalogWorkspaceBindingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetCatalogWorkspaceBindingsResponsePb{}
	pb.Workspaces = st.Workspaces

	return pb, nil
}

func GetCatalogWorkspaceBindingsResponseFromPb(pb *catalogpb.GetCatalogWorkspaceBindingsResponsePb) (*GetCatalogWorkspaceBindingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCatalogWorkspaceBindingsResponse{}
	st.Workspaces = pb.Workspaces

	return st, nil
}

type GetConnectionRequest struct {
	// Name of the connection.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func GetConnectionRequestToPb(st *GetConnectionRequest) (*catalogpb.GetConnectionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetConnectionRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetConnectionRequestFromPb(pb *catalogpb.GetConnectionRequestPb) (*GetConnectionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetConnectionRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetCredentialRequest struct {
	// Name of the credential.
	// Wire name: 'name_arg'
	NameArg string `tf:"-"`
}

func GetCredentialRequestToPb(st *GetCredentialRequest) (*catalogpb.GetCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetCredentialRequestPb{}
	pb.NameArg = st.NameArg

	return pb, nil
}

func GetCredentialRequestFromPb(pb *catalogpb.GetCredentialRequestPb) (*GetCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCredentialRequest{}
	st.NameArg = pb.NameArg

	return st, nil
}

type GetEffectiveRequest struct {
	// Full name of securable.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// Specifies the maximum number of privileges to return (page length). Every
	// EffectivePrivilegeAssignment present in a single page response is
	// guaranteed to contain all the effective privileges granted on (or
	// inherited by) the requested Securable for the respective principal.
	//
	// If not set, all the effective permissions are returned. If set to -
	// lesser than 0: invalid parameter error - 0: page length is set to a
	// server configured value - lesser than 150 but greater than 0: invalid
	// parameter error (this is to ensure that server is able to return at least
	// one complete EffectivePrivilegeAssignment in a single page response) -
	// greater than (or equal to) 150: page length is the minimum of this value
	// and a server configured value
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque token for the next page of results (pagination).
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// If provided, only the effective permissions for the specified principal
	// (user or group) are returned.
	// Wire name: 'principal'
	Principal string `tf:"-"`
	// Type of securable.
	// Wire name: 'securable_type'
	SecurableType   string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *GetEffectiveRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetEffectiveRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetEffectiveRequestToPb(st *GetEffectiveRequest) (*catalogpb.GetEffectiveRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetEffectiveRequestPb{}
	pb.FullName = st.FullName
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.Principal = st.Principal
	pb.SecurableType = st.SecurableType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetEffectiveRequestFromPb(pb *catalogpb.GetEffectiveRequestPb) (*GetEffectiveRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEffectiveRequest{}
	st.FullName = pb.FullName
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.Principal = pb.Principal
	st.SecurableType = pb.SecurableType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetExternalLocationRequest struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// Name of the external location.
	// Wire name: 'name'
	Name            string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *GetExternalLocationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetExternalLocationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetExternalLocationRequestToPb(st *GetExternalLocationRequest) (*catalogpb.GetExternalLocationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetExternalLocationRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetExternalLocationRequestFromPb(pb *catalogpb.GetExternalLocationRequestPb) (*GetExternalLocationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExternalLocationRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetExternalMetadataRequest struct {

	// Wire name: 'name'
	Name string `tf:"-"`
}

func GetExternalMetadataRequestToPb(st *GetExternalMetadataRequest) (*catalogpb.GetExternalMetadataRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetExternalMetadataRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetExternalMetadataRequestFromPb(pb *catalogpb.GetExternalMetadataRequestPb) (*GetExternalMetadataRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExternalMetadataRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetFunctionRequest struct {
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	// Wire name: 'name'
	Name            string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *GetFunctionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetFunctionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetFunctionRequestToPb(st *GetFunctionRequest) (*catalogpb.GetFunctionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetFunctionRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetFunctionRequestFromPb(pb *catalogpb.GetFunctionRequestPb) (*GetFunctionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetFunctionRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetGrantRequest struct {
	// Full name of securable.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// Specifies the maximum number of privileges to return (page length). Every
	// PrivilegeAssignment present in a single page response is guaranteed to
	// contain all the privileges granted on the requested Securable for the
	// respective principal.
	//
	// If not set, all the permissions are returned. If set to - lesser than 0:
	// invalid parameter error - 0: page length is set to a server configured
	// value - lesser than 150 but greater than 0: invalid parameter error (this
	// is to ensure that server is able to return at least one complete
	// PrivilegeAssignment in a single page response) - greater than (or equal
	// to) 150: page length is the minimum of this value and a server configured
	// value
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// If provided, only the permissions for the specified principal (user or
	// group) are returned.
	// Wire name: 'principal'
	Principal string `tf:"-"`
	// Type of securable.
	// Wire name: 'securable_type'
	SecurableType   string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *GetGrantRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetGrantRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetGrantRequestToPb(st *GetGrantRequest) (*catalogpb.GetGrantRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetGrantRequestPb{}
	pb.FullName = st.FullName
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.Principal = st.Principal
	pb.SecurableType = st.SecurableType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetGrantRequestFromPb(pb *catalogpb.GetGrantRequestPb) (*GetGrantRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetGrantRequest{}
	st.FullName = pb.FullName
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.Principal = pb.Principal
	st.SecurableType = pb.SecurableType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetMetastoreRequest struct {
	// Unique ID of the metastore.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func GetMetastoreRequestToPb(st *GetMetastoreRequest) (*catalogpb.GetMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetMetastoreRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func GetMetastoreRequestFromPb(pb *catalogpb.GetMetastoreRequestPb) (*GetMetastoreRequest, error) {
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
	Cloud string ``
	// Time at which this metastore was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of metastore creator.
	// Wire name: 'created_by'
	CreatedBy string ``
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	// Wire name: 'default_data_access_config_id'
	DefaultDataAccessConfigId string ``
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	// Wire name: 'delta_sharing_organization_name'
	DeltaSharingOrganizationName string ``
	// The lifetime of delta sharing recipient token in seconds.
	// Wire name: 'delta_sharing_recipient_token_lifetime_in_seconds'
	DeltaSharingRecipientTokenLifetimeInSeconds int64 ``
	// The scope of Delta Sharing enabled for the metastore.
	// Wire name: 'delta_sharing_scope'
	DeltaSharingScope DeltaSharingScopeEnum ``
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	// Wire name: 'external_access_enabled'
	ExternalAccessEnabled bool ``
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	// Wire name: 'global_metastore_id'
	GlobalMetastoreId string ``
	// Unique identifier of metastore.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// The user-specified name of the metastore.
	// Wire name: 'name'
	Name string ``
	// The owner of the metastore.
	// Wire name: 'owner'
	Owner string ``
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	// Wire name: 'privilege_model_version'
	PrivilegeModelVersion string ``
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	// Wire name: 'region'
	Region string ``
	// The storage root URL for metastore
	// Wire name: 'storage_root'
	StorageRoot string ``
	// UUID of storage credential to access the metastore storage_root.
	// Wire name: 'storage_root_credential_id'
	StorageRootCredentialId string ``
	// Name of the storage credential to access the metastore storage_root.
	// Wire name: 'storage_root_credential_name'
	StorageRootCredentialName string ``
	// Time at which the metastore was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// Username of user who last modified the metastore.
	// Wire name: 'updated_by'
	UpdatedBy       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GetMetastoreSummaryResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetMetastoreSummaryResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetMetastoreSummaryResponseToPb(st *GetMetastoreSummaryResponse) (*catalogpb.GetMetastoreSummaryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetMetastoreSummaryResponsePb{}
	pb.Cloud = st.Cloud
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.DefaultDataAccessConfigId = st.DefaultDataAccessConfigId
	pb.DeltaSharingOrganizationName = st.DeltaSharingOrganizationName
	pb.DeltaSharingRecipientTokenLifetimeInSeconds = st.DeltaSharingRecipientTokenLifetimeInSeconds
	deltaSharingScopePb, err := DeltaSharingScopeEnumToPb(&st.DeltaSharingScope)
	if err != nil {
		return nil, err
	}
	if deltaSharingScopePb != nil {
		pb.DeltaSharingScope = *deltaSharingScopePb
	}
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

func GetMetastoreSummaryResponseFromPb(pb *catalogpb.GetMetastoreSummaryResponsePb) (*GetMetastoreSummaryResponse, error) {
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
	deltaSharingScopeField, err := DeltaSharingScopeEnumFromPb(&pb.DeltaSharingScope)
	if err != nil {
		return nil, err
	}
	if deltaSharingScopeField != nil {
		st.DeltaSharingScope = *deltaSharingScopeField
	}
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
	Version         int      `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *GetModelVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetModelVersionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetModelVersionRequestToPb(st *GetModelVersionRequest) (*catalogpb.GetModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetModelVersionRequestPb{}
	pb.FullName = st.FullName
	pb.IncludeAliases = st.IncludeAliases
	pb.IncludeBrowse = st.IncludeBrowse
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetModelVersionRequestFromPb(pb *catalogpb.GetModelVersionRequestPb) (*GetModelVersionRequest, error) {
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

type GetOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func GetOnlineTableRequestToPb(st *GetOnlineTableRequest) (*catalogpb.GetOnlineTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetOnlineTableRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetOnlineTableRequestFromPb(pb *catalogpb.GetOnlineTableRequestPb) (*GetOnlineTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetOnlineTableRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetPermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// The privileges assigned to each principal
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment ``
	ForceSendFields      []string              `tf:"-"`
}

func (s *GetPermissionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPermissionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetPermissionsResponseToPb(st *GetPermissionsResponse) (*catalogpb.GetPermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetPermissionsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var privilegeAssignmentsPb []catalogpb.PrivilegeAssignmentPb
	for _, item := range st.PrivilegeAssignments {
		itemPb, err := PrivilegeAssignmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			privilegeAssignmentsPb = append(privilegeAssignmentsPb, *itemPb)
		}
	}
	pb.PrivilegeAssignments = privilegeAssignmentsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetPermissionsResponseFromPb(pb *catalogpb.GetPermissionsResponsePb) (*GetPermissionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPermissionsResponse{}
	st.NextPageToken = pb.NextPageToken

	var privilegeAssignmentsField []PrivilegeAssignment
	for _, itemPb := range pb.PrivilegeAssignments {
		item, err := PrivilegeAssignmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			privilegeAssignmentsField = append(privilegeAssignmentsField, *item)
		}
	}
	st.PrivilegeAssignments = privilegeAssignmentsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetQualityMonitorRequest struct {
	// UC table name in format `catalog.schema.table_name`. This field
	// corresponds to the {full_table_name_arg} arg in the endpoint path.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func GetQualityMonitorRequestToPb(st *GetQualityMonitorRequest) (*catalogpb.GetQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetQualityMonitorRequestPb{}
	pb.TableName = st.TableName

	return pb, nil
}

func GetQualityMonitorRequestFromPb(pb *catalogpb.GetQualityMonitorRequestPb) (*GetQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetQualityMonitorRequest{}
	st.TableName = pb.TableName

	return st, nil
}

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

func GetQuotaRequestToPb(st *GetQuotaRequest) (*catalogpb.GetQuotaRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetQuotaRequestPb{}
	pb.ParentFullName = st.ParentFullName
	pb.ParentSecurableType = st.ParentSecurableType
	pb.QuotaName = st.QuotaName

	return pb, nil
}

func GetQuotaRequestFromPb(pb *catalogpb.GetQuotaRequestPb) (*GetQuotaRequest, error) {
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
	QuotaInfo *QuotaInfo ``
}

func GetQuotaResponseToPb(st *GetQuotaResponse) (*catalogpb.GetQuotaResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetQuotaResponsePb{}
	quotaInfoPb, err := QuotaInfoToPb(st.QuotaInfo)
	if err != nil {
		return nil, err
	}
	if quotaInfoPb != nil {
		pb.QuotaInfo = quotaInfoPb
	}

	return pb, nil
}

func GetQuotaResponseFromPb(pb *catalogpb.GetQuotaResponsePb) (*GetQuotaResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetQuotaResponse{}
	quotaInfoField, err := QuotaInfoFromPb(pb.QuotaInfo)
	if err != nil {
		return nil, err
	}
	if quotaInfoField != nil {
		st.QuotaInfo = quotaInfoField
	}

	return st, nil
}

type GetRefreshRequest struct {
	// ID of the refresh.
	// Wire name: 'refresh_id'
	RefreshId int64 `tf:"-"`
	// Full name of the table.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func GetRefreshRequestToPb(st *GetRefreshRequest) (*catalogpb.GetRefreshRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetRefreshRequestPb{}
	pb.RefreshId = st.RefreshId
	pb.TableName = st.TableName

	return pb, nil
}

func GetRefreshRequestFromPb(pb *catalogpb.GetRefreshRequestPb) (*GetRefreshRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRefreshRequest{}
	st.RefreshId = pb.RefreshId
	st.TableName = pb.TableName

	return st, nil
}

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
	IncludeBrowse   bool     `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *GetRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetRegisteredModelRequestToPb(st *GetRegisteredModelRequest) (*catalogpb.GetRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetRegisteredModelRequestPb{}
	pb.FullName = st.FullName
	pb.IncludeAliases = st.IncludeAliases
	pb.IncludeBrowse = st.IncludeBrowse

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetRegisteredModelRequestFromPb(pb *catalogpb.GetRegisteredModelRequestPb) (*GetRegisteredModelRequest, error) {
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

type GetSchemaRequest struct {
	// Full name of the schema.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse   bool     `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *GetSchemaRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetSchemaRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetSchemaRequestToPb(st *GetSchemaRequest) (*catalogpb.GetSchemaRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetSchemaRequestPb{}
	pb.FullName = st.FullName
	pb.IncludeBrowse = st.IncludeBrowse

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetSchemaRequestFromPb(pb *catalogpb.GetSchemaRequestPb) (*GetSchemaRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSchemaRequest{}
	st.FullName = pb.FullName
	st.IncludeBrowse = pb.IncludeBrowse

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetStorageCredentialRequest struct {
	// Name of the storage credential.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func GetStorageCredentialRequestToPb(st *GetStorageCredentialRequest) (*catalogpb.GetStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetStorageCredentialRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetStorageCredentialRequestFromPb(pb *catalogpb.GetStorageCredentialRequestPb) (*GetStorageCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStorageCredentialRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetTableRequest struct {
	// Full name of the table.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for.
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// Whether delta metadata should be included in the response.
	// Wire name: 'include_delta_metadata'
	IncludeDeltaMetadata bool `tf:"-"`
	// Whether to include a manifest containing table capabilities in the
	// response.
	// Wire name: 'include_manifest_capabilities'
	IncludeManifestCapabilities bool     `tf:"-"`
	ForceSendFields             []string `tf:"-"`
}

func (s *GetTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetTableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetTableRequestToPb(st *GetTableRequest) (*catalogpb.GetTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetTableRequestPb{}
	pb.FullName = st.FullName
	pb.IncludeBrowse = st.IncludeBrowse
	pb.IncludeDeltaMetadata = st.IncludeDeltaMetadata
	pb.IncludeManifestCapabilities = st.IncludeManifestCapabilities

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetTableRequestFromPb(pb *catalogpb.GetTableRequestPb) (*GetTableRequest, error) {
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

type GetWorkspaceBindingRequest struct {
	// The name of the catalog.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func GetWorkspaceBindingRequestToPb(st *GetWorkspaceBindingRequest) (*catalogpb.GetWorkspaceBindingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetWorkspaceBindingRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetWorkspaceBindingRequestFromPb(pb *catalogpb.GetWorkspaceBindingRequestPb) (*GetWorkspaceBindingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceBindingRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetWorkspaceBindingsResponse struct {
	// List of workspace bindings
	// Wire name: 'bindings'
	Bindings []WorkspaceBinding ``
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GetWorkspaceBindingsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetWorkspaceBindingsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetWorkspaceBindingsResponseToPb(st *GetWorkspaceBindingsResponse) (*catalogpb.GetWorkspaceBindingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.GetWorkspaceBindingsResponsePb{}

	var bindingsPb []catalogpb.WorkspaceBindingPb
	for _, item := range st.Bindings {
		itemPb, err := WorkspaceBindingToPb(&item)
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

func GetWorkspaceBindingsResponseFromPb(pb *catalogpb.GetWorkspaceBindingsResponsePb) (*GetWorkspaceBindingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceBindingsResponse{}

	var bindingsField []WorkspaceBinding
	for _, itemPb := range pb.Bindings {
		item, err := WorkspaceBindingFromPb(&itemPb)
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

type IsolationMode string

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

// Values returns all possible values for IsolationMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *IsolationMode) Values() []IsolationMode {
	return []IsolationMode{
		IsolationModeIsolationModeIsolated,
		IsolationModeIsolationModeOpen,
	}
}

// Type always returns IsolationMode to satisfy [pflag.Value] interface
func (f *IsolationMode) Type() string {
	return "IsolationMode"
}

func IsolationModeToPb(st *IsolationMode) (*catalogpb.IsolationModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.IsolationModePb(*st)
	return &pb, nil
}

func IsolationModeFromPb(pb *catalogpb.IsolationModePb) (*IsolationMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := IsolationMode(*pb)
	return &st, nil
}

type LineageDirection string

const LineageDirectionDownstream LineageDirection = `DOWNSTREAM`

const LineageDirectionUpstream LineageDirection = `UPSTREAM`

// String representation for [fmt.Print]
func (f *LineageDirection) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LineageDirection) Set(v string) error {
	switch v {
	case `DOWNSTREAM`, `UPSTREAM`:
		*f = LineageDirection(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DOWNSTREAM", "UPSTREAM"`, v)
	}
}

// Values returns all possible values for LineageDirection.
//
// There is no guarantee on the order of the values in the slice.
func (f *LineageDirection) Values() []LineageDirection {
	return []LineageDirection{
		LineageDirectionDownstream,
		LineageDirectionUpstream,
	}
}

// Type always returns LineageDirection to satisfy [pflag.Value] interface
func (f *LineageDirection) Type() string {
	return "LineageDirection"
}

func LineageDirectionToPb(st *LineageDirection) (*catalogpb.LineageDirectionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.LineageDirectionPb(*st)
	return &pb, nil
}

func LineageDirectionFromPb(pb *catalogpb.LineageDirectionPb) (*LineageDirection, error) {
	if pb == nil {
		return nil, nil
	}
	st := LineageDirection(*pb)
	return &st, nil
}

type ListAccountMetastoreAssignmentsRequest struct {
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
}

func ListAccountMetastoreAssignmentsRequestToPb(st *ListAccountMetastoreAssignmentsRequest) (*catalogpb.ListAccountMetastoreAssignmentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListAccountMetastoreAssignmentsRequestPb{}
	pb.MetastoreId = st.MetastoreId

	return pb, nil
}

func ListAccountMetastoreAssignmentsRequestFromPb(pb *catalogpb.ListAccountMetastoreAssignmentsRequestPb) (*ListAccountMetastoreAssignmentsRequest, error) {
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
	WorkspaceIds []int64 ``
}

func ListAccountMetastoreAssignmentsResponseToPb(st *ListAccountMetastoreAssignmentsResponse) (*catalogpb.ListAccountMetastoreAssignmentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListAccountMetastoreAssignmentsResponsePb{}
	pb.WorkspaceIds = st.WorkspaceIds

	return pb, nil
}

func ListAccountMetastoreAssignmentsResponseFromPb(pb *catalogpb.ListAccountMetastoreAssignmentsResponsePb) (*ListAccountMetastoreAssignmentsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountMetastoreAssignmentsResponse{}
	st.WorkspaceIds = pb.WorkspaceIds

	return st, nil
}

type ListAccountStorageCredentialsRequest struct {
	// Unity Catalog metastore ID
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
}

func ListAccountStorageCredentialsRequestToPb(st *ListAccountStorageCredentialsRequest) (*catalogpb.ListAccountStorageCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListAccountStorageCredentialsRequestPb{}
	pb.MetastoreId = st.MetastoreId

	return pb, nil
}

func ListAccountStorageCredentialsRequestFromPb(pb *catalogpb.ListAccountStorageCredentialsRequestPb) (*ListAccountStorageCredentialsRequest, error) {
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
	StorageCredentials []StorageCredentialInfo ``
}

func ListAccountStorageCredentialsResponseToPb(st *ListAccountStorageCredentialsResponse) (*catalogpb.ListAccountStorageCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListAccountStorageCredentialsResponsePb{}

	var storageCredentialsPb []catalogpb.StorageCredentialInfoPb
	for _, item := range st.StorageCredentials {
		itemPb, err := StorageCredentialInfoToPb(&item)
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

func ListAccountStorageCredentialsResponseFromPb(pb *catalogpb.ListAccountStorageCredentialsResponsePb) (*ListAccountStorageCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountStorageCredentialsResponse{}

	var storageCredentialsField []StorageCredentialInfo
	for _, itemPb := range pb.StorageCredentials {
		item, err := StorageCredentialInfoFromPb(&itemPb)
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
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListCatalogsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCatalogsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListCatalogsRequestToPb(st *ListCatalogsRequest) (*catalogpb.ListCatalogsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListCatalogsRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListCatalogsRequestFromPb(pb *catalogpb.ListCatalogsRequestPb) (*ListCatalogsRequest, error) {
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

type ListCatalogsResponse struct {
	// An array of catalog information objects.
	// Wire name: 'catalogs'
	Catalogs []CatalogInfo ``
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListCatalogsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCatalogsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListCatalogsResponseToPb(st *ListCatalogsResponse) (*catalogpb.ListCatalogsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListCatalogsResponsePb{}

	var catalogsPb []catalogpb.CatalogInfoPb
	for _, item := range st.Catalogs {
		itemPb, err := CatalogInfoToPb(&item)
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

func ListCatalogsResponseFromPb(pb *catalogpb.ListCatalogsResponsePb) (*ListCatalogsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCatalogsResponse{}

	var catalogsField []CatalogInfo
	for _, itemPb := range pb.Catalogs {
		item, err := CatalogInfoFromPb(&itemPb)
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
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListConnectionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListConnectionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListConnectionsRequestToPb(st *ListConnectionsRequest) (*catalogpb.ListConnectionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListConnectionsRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListConnectionsRequestFromPb(pb *catalogpb.ListConnectionsRequestPb) (*ListConnectionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListConnectionsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListConnectionsResponse struct {
	// An array of connection information objects.
	// Wire name: 'connections'
	Connections []ConnectionInfo ``
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListConnectionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListConnectionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListConnectionsResponseToPb(st *ListConnectionsResponse) (*catalogpb.ListConnectionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListConnectionsResponsePb{}

	var connectionsPb []catalogpb.ConnectionInfoPb
	for _, item := range st.Connections {
		itemPb, err := ConnectionInfoToPb(&item)
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

func ListConnectionsResponseFromPb(pb *catalogpb.ListConnectionsResponsePb) (*ListConnectionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListConnectionsResponse{}

	var connectionsField []ConnectionInfo
	for _, itemPb := range pb.Connections {
		item, err := ConnectionInfoFromPb(&itemPb)
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
	Purpose         CredentialPurpose `tf:"-"`
	ForceSendFields []string          `tf:"-"`
}

func (s *ListCredentialsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCredentialsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListCredentialsRequestToPb(st *ListCredentialsRequest) (*catalogpb.ListCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListCredentialsRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	purposePb, err := CredentialPurposeToPb(&st.Purpose)
	if err != nil {
		return nil, err
	}
	if purposePb != nil {
		pb.Purpose = *purposePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListCredentialsRequestFromPb(pb *catalogpb.ListCredentialsRequestPb) (*ListCredentialsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCredentialsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	purposeField, err := CredentialPurposeFromPb(&pb.Purpose)
	if err != nil {
		return nil, err
	}
	if purposeField != nil {
		st.Purpose = *purposeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListCredentialsResponse struct {

	// Wire name: 'credentials'
	Credentials []CredentialInfo ``
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListCredentialsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCredentialsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListCredentialsResponseToPb(st *ListCredentialsResponse) (*catalogpb.ListCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListCredentialsResponsePb{}

	var credentialsPb []catalogpb.CredentialInfoPb
	for _, item := range st.Credentials {
		itemPb, err := CredentialInfoToPb(&item)
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

func ListCredentialsResponseFromPb(pb *catalogpb.ListCredentialsResponsePb) (*ListCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCredentialsResponse{}

	var credentialsField []CredentialInfo
	for _, itemPb := range pb.Credentials {
		item, err := CredentialInfoFromPb(&itemPb)
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

type ListExternalLineageRelationshipsRequest struct {
	// The lineage direction to filter on.
	// Wire name: 'lineage_direction'
	LineageDirection LineageDirection `tf:"-"`
	// The object to query external lineage relationships for. Since this field
	// is a query parameter, please flatten the nested fields. For example, if
	// the object is a table, the query parameter should look like:
	// `object_info.table.name=main.sales.customers`
	// Wire name: 'object_info'
	ObjectInfo ExternalLineageObject `tf:"-"`
	// Specifies the maximum number of external lineage relationships to return
	// in a single response. The value must be less than or equal to 1000.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListExternalLineageRelationshipsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalLineageRelationshipsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListExternalLineageRelationshipsRequestToPb(st *ListExternalLineageRelationshipsRequest) (*catalogpb.ListExternalLineageRelationshipsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListExternalLineageRelationshipsRequestPb{}
	lineageDirectionPb, err := LineageDirectionToPb(&st.LineageDirection)
	if err != nil {
		return nil, err
	}
	if lineageDirectionPb != nil {
		pb.LineageDirection = *lineageDirectionPb
	}
	objectInfoPb, err := ExternalLineageObjectToPb(&st.ObjectInfo)
	if err != nil {
		return nil, err
	}
	if objectInfoPb != nil {
		pb.ObjectInfo = *objectInfoPb
	}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListExternalLineageRelationshipsRequestFromPb(pb *catalogpb.ListExternalLineageRelationshipsRequestPb) (*ListExternalLineageRelationshipsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExternalLineageRelationshipsRequest{}
	lineageDirectionField, err := LineageDirectionFromPb(&pb.LineageDirection)
	if err != nil {
		return nil, err
	}
	if lineageDirectionField != nil {
		st.LineageDirection = *lineageDirectionField
	}
	objectInfoField, err := ExternalLineageObjectFromPb(&pb.ObjectInfo)
	if err != nil {
		return nil, err
	}
	if objectInfoField != nil {
		st.ObjectInfo = *objectInfoField
	}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListExternalLineageRelationshipsResponse struct {

	// Wire name: 'external_lineage_relationships'
	ExternalLineageRelationships []ExternalLineageInfo ``

	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListExternalLineageRelationshipsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalLineageRelationshipsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListExternalLineageRelationshipsResponseToPb(st *ListExternalLineageRelationshipsResponse) (*catalogpb.ListExternalLineageRelationshipsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListExternalLineageRelationshipsResponsePb{}

	var externalLineageRelationshipsPb []catalogpb.ExternalLineageInfoPb
	for _, item := range st.ExternalLineageRelationships {
		itemPb, err := ExternalLineageInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			externalLineageRelationshipsPb = append(externalLineageRelationshipsPb, *itemPb)
		}
	}
	pb.ExternalLineageRelationships = externalLineageRelationshipsPb
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListExternalLineageRelationshipsResponseFromPb(pb *catalogpb.ListExternalLineageRelationshipsResponsePb) (*ListExternalLineageRelationshipsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExternalLineageRelationshipsResponse{}

	var externalLineageRelationshipsField []ExternalLineageInfo
	for _, itemPb := range pb.ExternalLineageRelationships {
		item, err := ExternalLineageInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			externalLineageRelationshipsField = append(externalLineageRelationshipsField, *item)
		}
	}
	st.ExternalLineageRelationships = externalLineageRelationshipsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

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
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListExternalLocationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalLocationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListExternalLocationsRequestToPb(st *ListExternalLocationsRequest) (*catalogpb.ListExternalLocationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListExternalLocationsRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListExternalLocationsRequestFromPb(pb *catalogpb.ListExternalLocationsRequestPb) (*ListExternalLocationsRequest, error) {
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

type ListExternalLocationsResponse struct {
	// An array of external locations.
	// Wire name: 'external_locations'
	ExternalLocations []ExternalLocationInfo ``
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListExternalLocationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalLocationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListExternalLocationsResponseToPb(st *ListExternalLocationsResponse) (*catalogpb.ListExternalLocationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListExternalLocationsResponsePb{}

	var externalLocationsPb []catalogpb.ExternalLocationInfoPb
	for _, item := range st.ExternalLocations {
		itemPb, err := ExternalLocationInfoToPb(&item)
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

func ListExternalLocationsResponseFromPb(pb *catalogpb.ListExternalLocationsResponsePb) (*ListExternalLocationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExternalLocationsResponse{}

	var externalLocationsField []ExternalLocationInfo
	for _, itemPb := range pb.ExternalLocations {
		item, err := ExternalLocationInfoFromPb(&itemPb)
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

type ListExternalMetadataRequest struct {
	// Specifies the maximum number of external metadata objects to return in a
	// single response. The value must be less than or equal to 1000.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListExternalMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalMetadataRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListExternalMetadataRequestToPb(st *ListExternalMetadataRequest) (*catalogpb.ListExternalMetadataRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListExternalMetadataRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListExternalMetadataRequestFromPb(pb *catalogpb.ListExternalMetadataRequestPb) (*ListExternalMetadataRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExternalMetadataRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListExternalMetadataResponse struct {

	// Wire name: 'external_metadata'
	ExternalMetadata []ExternalMetadata ``

	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListExternalMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExternalMetadataResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListExternalMetadataResponseToPb(st *ListExternalMetadataResponse) (*catalogpb.ListExternalMetadataResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListExternalMetadataResponsePb{}

	var externalMetadataPb []catalogpb.ExternalMetadataPb
	for _, item := range st.ExternalMetadata {
		itemPb, err := ExternalMetadataToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			externalMetadataPb = append(externalMetadataPb, *itemPb)
		}
	}
	pb.ExternalMetadata = externalMetadataPb
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListExternalMetadataResponseFromPb(pb *catalogpb.ListExternalMetadataResponsePb) (*ListExternalMetadataResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExternalMetadataResponse{}

	var externalMetadataField []ExternalMetadata
	for _, itemPb := range pb.ExternalMetadata {
		item, err := ExternalMetadataFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			externalMetadataField = append(externalMetadataField, *item)
		}
	}
	st.ExternalMetadata = externalMetadataField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

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
	SchemaName      string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListFunctionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFunctionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListFunctionsRequestToPb(st *ListFunctionsRequest) (*catalogpb.ListFunctionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListFunctionsRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.IncludeBrowse = st.IncludeBrowse
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListFunctionsRequestFromPb(pb *catalogpb.ListFunctionsRequestPb) (*ListFunctionsRequest, error) {
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

type ListFunctionsResponse struct {
	// An array of function information objects.
	// Wire name: 'functions'
	Functions []FunctionInfo ``
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListFunctionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFunctionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListFunctionsResponseToPb(st *ListFunctionsResponse) (*catalogpb.ListFunctionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListFunctionsResponsePb{}

	var functionsPb []catalogpb.FunctionInfoPb
	for _, item := range st.Functions {
		itemPb, err := FunctionInfoToPb(&item)
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

func ListFunctionsResponseFromPb(pb *catalogpb.ListFunctionsResponsePb) (*ListFunctionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFunctionsResponse{}

	var functionsField []FunctionInfo
	for _, itemPb := range pb.Functions {
		item, err := FunctionInfoFromPb(&itemPb)
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

type ListMetastoresRequest struct {
	// Maximum number of metastores to return. - when set to a value greater
	// than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned; - If not set, all the metastores are
	// returned (not recommended). - Note: The number of returned metastores
	// might be less than the specified max_results size, even zero. The only
	// definitive indication that no further metastores can be fetched is when
	// the next_page_token is unset from the response.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListMetastoresRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMetastoresRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListMetastoresRequestToPb(st *ListMetastoresRequest) (*catalogpb.ListMetastoresRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListMetastoresRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListMetastoresRequestFromPb(pb *catalogpb.ListMetastoresRequestPb) (*ListMetastoresRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListMetastoresRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListMetastoresResponse struct {
	// An array of metastore information objects.
	// Wire name: 'metastores'
	Metastores []MetastoreInfo ``
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListMetastoresResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMetastoresResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListMetastoresResponseToPb(st *ListMetastoresResponse) (*catalogpb.ListMetastoresResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListMetastoresResponsePb{}

	var metastoresPb []catalogpb.MetastoreInfoPb
	for _, item := range st.Metastores {
		itemPb, err := MetastoreInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			metastoresPb = append(metastoresPb, *itemPb)
		}
	}
	pb.Metastores = metastoresPb
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListMetastoresResponseFromPb(pb *catalogpb.ListMetastoresResponsePb) (*ListMetastoresResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListMetastoresResponse{}

	var metastoresField []MetastoreInfo
	for _, itemPb := range pb.Metastores {
		item, err := MetastoreInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			metastoresField = append(metastoresField, *item)
		}
	}
	st.Metastores = metastoresField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

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
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListModelVersionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelVersionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListModelVersionsRequestToPb(st *ListModelVersionsRequest) (*catalogpb.ListModelVersionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListModelVersionsRequestPb{}
	pb.FullName = st.FullName
	pb.IncludeBrowse = st.IncludeBrowse
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListModelVersionsRequestFromPb(pb *catalogpb.ListModelVersionsRequestPb) (*ListModelVersionsRequest, error) {
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

type ListModelVersionsResponse struct {

	// Wire name: 'model_versions'
	ModelVersions []ModelVersionInfo ``
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListModelVersionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListModelVersionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListModelVersionsResponseToPb(st *ListModelVersionsResponse) (*catalogpb.ListModelVersionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListModelVersionsResponsePb{}

	var modelVersionsPb []catalogpb.ModelVersionInfoPb
	for _, item := range st.ModelVersions {
		itemPb, err := ModelVersionInfoToPb(&item)
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

func ListModelVersionsResponseFromPb(pb *catalogpb.ListModelVersionsResponsePb) (*ListModelVersionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListModelVersionsResponse{}

	var modelVersionsField []ModelVersionInfo
	for _, itemPb := range pb.ModelVersions {
		item, err := ModelVersionInfoFromPb(&itemPb)
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

type ListQuotasRequest struct {
	// The number of quotas to return.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque token for the next page of results.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListQuotasRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQuotasRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListQuotasRequestToPb(st *ListQuotasRequest) (*catalogpb.ListQuotasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListQuotasRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListQuotasRequestFromPb(pb *catalogpb.ListQuotasRequestPb) (*ListQuotasRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQuotasRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListQuotasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// An array of returned QuotaInfos.
	// Wire name: 'quotas'
	Quotas          []QuotaInfo ``
	ForceSendFields []string    `tf:"-"`
}

func (s *ListQuotasResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQuotasResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListQuotasResponseToPb(st *ListQuotasResponse) (*catalogpb.ListQuotasResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListQuotasResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var quotasPb []catalogpb.QuotaInfoPb
	for _, item := range st.Quotas {
		itemPb, err := QuotaInfoToPb(&item)
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

func ListQuotasResponseFromPb(pb *catalogpb.ListQuotasResponsePb) (*ListQuotasResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQuotasResponse{}
	st.NextPageToken = pb.NextPageToken

	var quotasField []QuotaInfo
	for _, itemPb := range pb.Quotas {
		item, err := QuotaInfoFromPb(&itemPb)
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

type ListRefreshesRequest struct {
	// UC table name in format `catalog.schema.table_name`. table_name is case
	// insensitive and spaces are disallowed.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func ListRefreshesRequestToPb(st *ListRefreshesRequest) (*catalogpb.ListRefreshesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListRefreshesRequestPb{}
	pb.TableName = st.TableName

	return pb, nil
}

func ListRefreshesRequestFromPb(pb *catalogpb.ListRefreshesRequestPb) (*ListRefreshesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRefreshesRequest{}
	st.TableName = pb.TableName

	return st, nil
}

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
	SchemaName      string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListRegisteredModelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRegisteredModelsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListRegisteredModelsRequestToPb(st *ListRegisteredModelsRequest) (*catalogpb.ListRegisteredModelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListRegisteredModelsRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.IncludeBrowse = st.IncludeBrowse
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListRegisteredModelsRequestFromPb(pb *catalogpb.ListRegisteredModelsRequestPb) (*ListRegisteredModelsRequest, error) {
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

type ListRegisteredModelsResponse struct {
	// Opaque token for pagination. Omitted if there are no more results.
	// page_token should be set to this value for fetching the next page.
	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'registered_models'
	RegisteredModels []RegisteredModelInfo ``
	ForceSendFields  []string              `tf:"-"`
}

func (s *ListRegisteredModelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRegisteredModelsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListRegisteredModelsResponseToPb(st *ListRegisteredModelsResponse) (*catalogpb.ListRegisteredModelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListRegisteredModelsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var registeredModelsPb []catalogpb.RegisteredModelInfoPb
	for _, item := range st.RegisteredModels {
		itemPb, err := RegisteredModelInfoToPb(&item)
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

func ListRegisteredModelsResponseFromPb(pb *catalogpb.ListRegisteredModelsResponsePb) (*ListRegisteredModelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRegisteredModelsResponse{}
	st.NextPageToken = pb.NextPageToken

	var registeredModelsField []RegisteredModelInfo
	for _, itemPb := range pb.RegisteredModels {
		item, err := RegisteredModelInfoFromPb(&itemPb)
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
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListSchemasRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSchemasRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListSchemasRequestToPb(st *ListSchemasRequest) (*catalogpb.ListSchemasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListSchemasRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.IncludeBrowse = st.IncludeBrowse
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListSchemasRequestFromPb(pb *catalogpb.ListSchemasRequestPb) (*ListSchemasRequest, error) {
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

type ListSchemasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// An array of schema information objects.
	// Wire name: 'schemas'
	Schemas         []SchemaInfo ``
	ForceSendFields []string     `tf:"-"`
}

func (s *ListSchemasResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSchemasResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListSchemasResponseToPb(st *ListSchemasResponse) (*catalogpb.ListSchemasResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListSchemasResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var schemasPb []catalogpb.SchemaInfoPb
	for _, item := range st.Schemas {
		itemPb, err := SchemaInfoToPb(&item)
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

func ListSchemasResponseFromPb(pb *catalogpb.ListSchemasResponsePb) (*ListSchemasResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSchemasResponse{}
	st.NextPageToken = pb.NextPageToken

	var schemasField []SchemaInfo
	for _, itemPb := range pb.Schemas {
		item, err := SchemaInfoFromPb(&itemPb)
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
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListStorageCredentialsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListStorageCredentialsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListStorageCredentialsRequestToPb(st *ListStorageCredentialsRequest) (*catalogpb.ListStorageCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListStorageCredentialsRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListStorageCredentialsRequestFromPb(pb *catalogpb.ListStorageCredentialsRequestPb) (*ListStorageCredentialsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListStorageCredentialsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListStorageCredentialsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'storage_credentials'
	StorageCredentials []StorageCredentialInfo ``
	ForceSendFields    []string                `tf:"-"`
}

func (s *ListStorageCredentialsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListStorageCredentialsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListStorageCredentialsResponseToPb(st *ListStorageCredentialsResponse) (*catalogpb.ListStorageCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListStorageCredentialsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var storageCredentialsPb []catalogpb.StorageCredentialInfoPb
	for _, item := range st.StorageCredentials {
		itemPb, err := StorageCredentialInfoToPb(&item)
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

func ListStorageCredentialsResponseFromPb(pb *catalogpb.ListStorageCredentialsResponsePb) (*ListStorageCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListStorageCredentialsResponse{}
	st.NextPageToken = pb.NextPageToken

	var storageCredentialsField []StorageCredentialInfo
	for _, itemPb := range pb.StorageCredentials {
		item, err := StorageCredentialInfoFromPb(&itemPb)
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

type ListSummariesRequest struct {
	// Name of parent catalog for tables of interest.
	// Wire name: 'catalog_name'
	CatalogName string `tf:"-"`
	// Whether to include a manifest containing table capabilities in the
	// response.
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
	TableNamePattern string   `tf:"-"`
	ForceSendFields  []string `tf:"-"`
}

func (s *ListSummariesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSummariesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListSummariesRequestToPb(st *ListSummariesRequest) (*catalogpb.ListSummariesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListSummariesRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.IncludeManifestCapabilities = st.IncludeManifestCapabilities
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.SchemaNamePattern = st.SchemaNamePattern
	pb.TableNamePattern = st.TableNamePattern

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListSummariesRequestFromPb(pb *catalogpb.ListSummariesRequestPb) (*ListSummariesRequest, error) {
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
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListSystemSchemasRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSystemSchemasRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListSystemSchemasRequestToPb(st *ListSystemSchemasRequest) (*catalogpb.ListSystemSchemasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListSystemSchemasRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.MetastoreId = st.MetastoreId
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListSystemSchemasRequestFromPb(pb *catalogpb.ListSystemSchemasRequestPb) (*ListSystemSchemasRequest, error) {
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

type ListSystemSchemasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// An array of system schema information objects.
	// Wire name: 'schemas'
	Schemas         []SystemSchemaInfo ``
	ForceSendFields []string           `tf:"-"`
}

func (s *ListSystemSchemasResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSystemSchemasResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListSystemSchemasResponseToPb(st *ListSystemSchemasResponse) (*catalogpb.ListSystemSchemasResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListSystemSchemasResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var schemasPb []catalogpb.SystemSchemaInfoPb
	for _, item := range st.Schemas {
		itemPb, err := SystemSchemaInfoToPb(&item)
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

func ListSystemSchemasResponseFromPb(pb *catalogpb.ListSystemSchemasResponsePb) (*ListSystemSchemasResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSystemSchemasResponse{}
	st.NextPageToken = pb.NextPageToken

	var schemasField []SystemSchemaInfo
	for _, itemPb := range pb.Schemas {
		item, err := SystemSchemaInfoFromPb(&itemPb)
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

type ListTableSummariesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// List of table summaries.
	// Wire name: 'tables'
	Tables          []TableSummary ``
	ForceSendFields []string       `tf:"-"`
}

func (s *ListTableSummariesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTableSummariesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListTableSummariesResponseToPb(st *ListTableSummariesResponse) (*catalogpb.ListTableSummariesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListTableSummariesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var tablesPb []catalogpb.TableSummaryPb
	for _, item := range st.Tables {
		itemPb, err := TableSummaryToPb(&item)
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

func ListTableSummariesResponseFromPb(pb *catalogpb.ListTableSummariesResponsePb) (*ListTableSummariesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTableSummariesResponse{}
	st.NextPageToken = pb.NextPageToken

	var tablesField []TableSummary
	for _, itemPb := range pb.Tables {
		item, err := TableSummaryFromPb(&itemPb)
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

type ListTablesRequest struct {
	// Name of parent catalog for tables of interest.
	// Wire name: 'catalog_name'
	CatalogName string `tf:"-"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for.
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// Whether to include a manifest containing table capabilities in the
	// response.
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
	SchemaName      string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListTablesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTablesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListTablesRequestToPb(st *ListTablesRequest) (*catalogpb.ListTablesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListTablesRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.IncludeBrowse = st.IncludeBrowse
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

func ListTablesRequestFromPb(pb *catalogpb.ListTablesRequestPb) (*ListTablesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTablesRequest{}
	st.CatalogName = pb.CatalogName
	st.IncludeBrowse = pb.IncludeBrowse
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

type ListTablesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// An array of table information objects.
	// Wire name: 'tables'
	Tables          []TableInfo ``
	ForceSendFields []string    `tf:"-"`
}

func (s *ListTablesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTablesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListTablesResponseToPb(st *ListTablesResponse) (*catalogpb.ListTablesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListTablesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var tablesPb []catalogpb.TableInfoPb
	for _, item := range st.Tables {
		itemPb, err := TableInfoToPb(&item)
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

func ListTablesResponseFromPb(pb *catalogpb.ListTablesResponsePb) (*ListTablesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTablesResponse{}
	st.NextPageToken = pb.NextPageToken

	var tablesField []TableInfo
	for _, itemPb := range pb.Tables {
		item, err := TableInfoFromPb(&itemPb)
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
	SchemaName      string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListVolumesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVolumesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListVolumesRequestToPb(st *ListVolumesRequest) (*catalogpb.ListVolumesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListVolumesRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.IncludeBrowse = st.IncludeBrowse
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListVolumesRequestFromPb(pb *catalogpb.ListVolumesRequestPb) (*ListVolumesRequest, error) {
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

type ListVolumesResponseContent struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request to retrieve the next page of results.
	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'volumes'
	Volumes         []VolumeInfo ``
	ForceSendFields []string     `tf:"-"`
}

func (s *ListVolumesResponseContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVolumesResponseContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListVolumesResponseContentToPb(st *ListVolumesResponseContent) (*catalogpb.ListVolumesResponseContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ListVolumesResponseContentPb{}
	pb.NextPageToken = st.NextPageToken

	var volumesPb []catalogpb.VolumeInfoPb
	for _, item := range st.Volumes {
		itemPb, err := VolumeInfoToPb(&item)
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

func ListVolumesResponseContentFromPb(pb *catalogpb.ListVolumesResponseContentPb) (*ListVolumesResponseContent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListVolumesResponseContent{}
	st.NextPageToken = pb.NextPageToken

	var volumesField []VolumeInfo
	for _, itemPb := range pb.Volumes {
		item, err := VolumeInfoFromPb(&itemPb)
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

// The artifact pattern matching type
type MatchType string

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

// Values returns all possible values for MatchType.
//
// There is no guarantee on the order of the values in the slice.
func (f *MatchType) Values() []MatchType {
	return []MatchType{
		MatchTypePrefixMatch,
	}
}

// Type always returns MatchType to satisfy [pflag.Value] interface
func (f *MatchType) Type() string {
	return "MatchType"
}

func MatchTypeToPb(st *MatchType) (*catalogpb.MatchTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.MatchTypePb(*st)
	return &pb, nil
}

func MatchTypeFromPb(pb *catalogpb.MatchTypePb) (*MatchType, error) {
	if pb == nil {
		return nil, nil
	}
	st := MatchType(*pb)
	return &st, nil
}

type MetastoreAssignment struct {
	// The name of the default catalog in the metastore.
	// Wire name: 'default_catalog_name'
	DefaultCatalogName string ``
	// The unique ID of the metastore.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// The unique ID of the Databricks workspace.
	// Wire name: 'workspace_id'
	WorkspaceId     int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *MetastoreAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MetastoreAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func MetastoreAssignmentToPb(st *MetastoreAssignment) (*catalogpb.MetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.MetastoreAssignmentPb{}
	pb.DefaultCatalogName = st.DefaultCatalogName
	pb.MetastoreId = st.MetastoreId
	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func MetastoreAssignmentFromPb(pb *catalogpb.MetastoreAssignmentPb) (*MetastoreAssignment, error) {
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

type MetastoreInfo struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	// Wire name: 'cloud'
	Cloud string ``
	// Time at which this metastore was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of metastore creator.
	// Wire name: 'created_by'
	CreatedBy string ``
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	// Wire name: 'default_data_access_config_id'
	DefaultDataAccessConfigId string ``
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	// Wire name: 'delta_sharing_organization_name'
	DeltaSharingOrganizationName string ``
	// The lifetime of delta sharing recipient token in seconds.
	// Wire name: 'delta_sharing_recipient_token_lifetime_in_seconds'
	DeltaSharingRecipientTokenLifetimeInSeconds int64 ``
	// The scope of Delta Sharing enabled for the metastore.
	// Wire name: 'delta_sharing_scope'
	DeltaSharingScope DeltaSharingScopeEnum ``
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	// Wire name: 'external_access_enabled'
	ExternalAccessEnabled bool ``
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	// Wire name: 'global_metastore_id'
	GlobalMetastoreId string ``
	// Unique identifier of metastore.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// The user-specified name of the metastore.
	// Wire name: 'name'
	Name string ``
	// The owner of the metastore.
	// Wire name: 'owner'
	Owner string ``
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	// Wire name: 'privilege_model_version'
	PrivilegeModelVersion string ``
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	// Wire name: 'region'
	Region string ``
	// The storage root URL for metastore
	// Wire name: 'storage_root'
	StorageRoot string ``
	// UUID of storage credential to access the metastore storage_root.
	// Wire name: 'storage_root_credential_id'
	StorageRootCredentialId string ``
	// Name of the storage credential to access the metastore storage_root.
	// Wire name: 'storage_root_credential_name'
	StorageRootCredentialName string ``
	// Time at which the metastore was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// Username of user who last modified the metastore.
	// Wire name: 'updated_by'
	UpdatedBy       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *MetastoreInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MetastoreInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func MetastoreInfoToPb(st *MetastoreInfo) (*catalogpb.MetastoreInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.MetastoreInfoPb{}
	pb.Cloud = st.Cloud
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.DefaultDataAccessConfigId = st.DefaultDataAccessConfigId
	pb.DeltaSharingOrganizationName = st.DeltaSharingOrganizationName
	pb.DeltaSharingRecipientTokenLifetimeInSeconds = st.DeltaSharingRecipientTokenLifetimeInSeconds
	deltaSharingScopePb, err := DeltaSharingScopeEnumToPb(&st.DeltaSharingScope)
	if err != nil {
		return nil, err
	}
	if deltaSharingScopePb != nil {
		pb.DeltaSharingScope = *deltaSharingScopePb
	}
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

func MetastoreInfoFromPb(pb *catalogpb.MetastoreInfoPb) (*MetastoreInfo, error) {
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
	deltaSharingScopeField, err := DeltaSharingScopeEnumFromPb(&pb.DeltaSharingScope)
	if err != nil {
		return nil, err
	}
	if deltaSharingScopeField != nil {
		st.DeltaSharingScope = *deltaSharingScopeField
	}
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

type ModelVersionInfo struct {
	// List of aliases associated with the model version
	// Wire name: 'aliases'
	Aliases []RegisteredModelAlias ``
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool ``
	// The name of the catalog containing the model version
	// Wire name: 'catalog_name'
	CatalogName string ``
	// The comment attached to the model version
	// Wire name: 'comment'
	Comment string ``

	// Wire name: 'created_at'
	CreatedAt int64 ``
	// The identifier of the user who created the model version
	// Wire name: 'created_by'
	CreatedBy string ``
	// The unique identifier of the model version
	// Wire name: 'id'
	Id string ``
	// The unique identifier of the metastore containing the model version
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// The name of the parent registered model of the model version, relative to
	// parent schema
	// Wire name: 'model_name'
	ModelName string ``
	// Model version dependencies, for feature-store packaged models
	// Wire name: 'model_version_dependencies'
	ModelVersionDependencies *DependencyList ``
	// MLflow run ID used when creating the model version, if ``source`` was
	// generated by an experiment run stored in an MLflow tracking server
	// Wire name: 'run_id'
	RunId string ``
	// ID of the Databricks workspace containing the MLflow run that generated
	// this model version, if applicable
	// Wire name: 'run_workspace_id'
	RunWorkspaceId int ``
	// The name of the schema containing the model version, relative to parent
	// catalog
	// Wire name: 'schema_name'
	SchemaName string ``
	// URI indicating the location of the source artifacts (files) for the model
	// version
	// Wire name: 'source'
	Source string ``
	// Current status of the model version. Newly created model versions start
	// in PENDING_REGISTRATION status, then move to READY status once the model
	// version files are uploaded and the model version is finalized. Only model
	// versions in READY status can be loaded for inference or served.
	// Wire name: 'status'
	Status ModelVersionInfoStatus ``
	// The storage location on the cloud under which model version data files
	// are stored
	// Wire name: 'storage_location'
	StorageLocation string ``

	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// The identifier of the user who updated the model version last time
	// Wire name: 'updated_by'
	UpdatedBy string ``
	// Integer model version number, used to reference the model version in API
	// requests.
	// Wire name: 'version'
	Version         int      ``
	ForceSendFields []string `tf:"-"`
}

func (s *ModelVersionInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ModelVersionInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ModelVersionInfoToPb(st *ModelVersionInfo) (*catalogpb.ModelVersionInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ModelVersionInfoPb{}

	var aliasesPb []catalogpb.RegisteredModelAliasPb
	for _, item := range st.Aliases {
		itemPb, err := RegisteredModelAliasToPb(&item)
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
	modelVersionDependenciesPb, err := DependencyListToPb(st.ModelVersionDependencies)
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
	statusPb, err := ModelVersionInfoStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.StorageLocation = st.StorageLocation
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ModelVersionInfoFromPb(pb *catalogpb.ModelVersionInfoPb) (*ModelVersionInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelVersionInfo{}

	var aliasesField []RegisteredModelAlias
	for _, itemPb := range pb.Aliases {
		item, err := RegisteredModelAliasFromPb(&itemPb)
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
	modelVersionDependenciesField, err := DependencyListFromPb(pb.ModelVersionDependencies)
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
	statusField, err := ModelVersionInfoStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.StorageLocation = pb.StorageLocation
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Current status of the model version. Newly created model versions start in
// PENDING_REGISTRATION status, then move to READY status once the model version
// files are uploaded and the model version is finalized. Only model versions in
// READY status can be loaded for inference or served.
type ModelVersionInfoStatus string

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

// Values returns all possible values for ModelVersionInfoStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *ModelVersionInfoStatus) Values() []ModelVersionInfoStatus {
	return []ModelVersionInfoStatus{
		ModelVersionInfoStatusFailedRegistration,
		ModelVersionInfoStatusPendingRegistration,
		ModelVersionInfoStatusReady,
	}
}

// Type always returns ModelVersionInfoStatus to satisfy [pflag.Value] interface
func (f *ModelVersionInfoStatus) Type() string {
	return "ModelVersionInfoStatus"
}

func ModelVersionInfoStatusToPb(st *ModelVersionInfoStatus) (*catalogpb.ModelVersionInfoStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.ModelVersionInfoStatusPb(*st)
	return &pb, nil
}

func ModelVersionInfoStatusFromPb(pb *catalogpb.ModelVersionInfoStatusPb) (*ModelVersionInfoStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := ModelVersionInfoStatus(*pb)
	return &st, nil
}

type MonitorCronSchedule struct {
	// Read only field that indicates whether a schedule is paused or not.
	// Wire name: 'pause_status'
	PauseStatus MonitorCronSchedulePauseStatus ``
	// The expression that determines when to run the monitor. See [examples].
	//
	// [examples]: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	// Wire name: 'quartz_cron_expression'
	QuartzCronExpression string ``
	// The timezone id (e.g., ``PST``) in which to evaluate the quartz
	// expression.
	// Wire name: 'timezone_id'
	TimezoneId string ``
}

func MonitorCronScheduleToPb(st *MonitorCronSchedule) (*catalogpb.MonitorCronSchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.MonitorCronSchedulePb{}
	pauseStatusPb, err := MonitorCronSchedulePauseStatusToPb(&st.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusPb != nil {
		pb.PauseStatus = *pauseStatusPb
	}
	pb.QuartzCronExpression = st.QuartzCronExpression
	pb.TimezoneId = st.TimezoneId

	return pb, nil
}

func MonitorCronScheduleFromPb(pb *catalogpb.MonitorCronSchedulePb) (*MonitorCronSchedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorCronSchedule{}
	pauseStatusField, err := MonitorCronSchedulePauseStatusFromPb(&pb.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusField != nil {
		st.PauseStatus = *pauseStatusField
	}
	st.QuartzCronExpression = pb.QuartzCronExpression
	st.TimezoneId = pb.TimezoneId

	return st, nil
}

// Source link:
// https://src.dev.databricks.com/databricks/universe/-/blob/elastic-spark-common/api/messages/schedule.proto
// Monitoring workflow schedule pause status.
type MonitorCronSchedulePauseStatus string

const MonitorCronSchedulePauseStatusPaused MonitorCronSchedulePauseStatus = `PAUSED`

const MonitorCronSchedulePauseStatusUnpaused MonitorCronSchedulePauseStatus = `UNPAUSED`

const MonitorCronSchedulePauseStatusUnspecified MonitorCronSchedulePauseStatus = `UNSPECIFIED`

// String representation for [fmt.Print]
func (f *MonitorCronSchedulePauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorCronSchedulePauseStatus) Set(v string) error {
	switch v {
	case `PAUSED`, `UNPAUSED`, `UNSPECIFIED`:
		*f = MonitorCronSchedulePauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PAUSED", "UNPAUSED", "UNSPECIFIED"`, v)
	}
}

// Values returns all possible values for MonitorCronSchedulePauseStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *MonitorCronSchedulePauseStatus) Values() []MonitorCronSchedulePauseStatus {
	return []MonitorCronSchedulePauseStatus{
		MonitorCronSchedulePauseStatusPaused,
		MonitorCronSchedulePauseStatusUnpaused,
		MonitorCronSchedulePauseStatusUnspecified,
	}
}

// Type always returns MonitorCronSchedulePauseStatus to satisfy [pflag.Value] interface
func (f *MonitorCronSchedulePauseStatus) Type() string {
	return "MonitorCronSchedulePauseStatus"
}

func MonitorCronSchedulePauseStatusToPb(st *MonitorCronSchedulePauseStatus) (*catalogpb.MonitorCronSchedulePauseStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.MonitorCronSchedulePauseStatusPb(*st)
	return &pb, nil
}

func MonitorCronSchedulePauseStatusFromPb(pb *catalogpb.MonitorCronSchedulePauseStatusPb) (*MonitorCronSchedulePauseStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := MonitorCronSchedulePauseStatus(*pb)
	return &st, nil
}

// Data classification related configuration.
type MonitorDataClassificationConfig struct {
	// Whether to enable data classification.
	// Wire name: 'enabled'
	Enabled         bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *MonitorDataClassificationConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MonitorDataClassificationConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func MonitorDataClassificationConfigToPb(st *MonitorDataClassificationConfig) (*catalogpb.MonitorDataClassificationConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.MonitorDataClassificationConfigPb{}
	pb.Enabled = st.Enabled

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func MonitorDataClassificationConfigFromPb(pb *catalogpb.MonitorDataClassificationConfigPb) (*MonitorDataClassificationConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorDataClassificationConfig{}
	st.Enabled = pb.Enabled

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type MonitorDestination struct {
	// The list of email addresses to send the notification to. A maximum of 5
	// email addresses is supported.
	// Wire name: 'email_addresses'
	EmailAddresses []string ``
}

func MonitorDestinationToPb(st *MonitorDestination) (*catalogpb.MonitorDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.MonitorDestinationPb{}
	pb.EmailAddresses = st.EmailAddresses

	return pb, nil
}

func MonitorDestinationFromPb(pb *catalogpb.MonitorDestinationPb) (*MonitorDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorDestination{}
	st.EmailAddresses = pb.EmailAddresses

	return st, nil
}

type MonitorInferenceLog struct {
	// List of granularities to use when aggregating data into time windows
	// based on their timestamp.
	// Wire name: 'granularities'
	Granularities []string ``
	// Column for the label.
	// Wire name: 'label_col'
	LabelCol string ``
	// Column for the model identifier.
	// Wire name: 'model_id_col'
	ModelIdCol string ``
	// Column for the prediction.
	// Wire name: 'prediction_col'
	PredictionCol string ``
	// Column for prediction probabilities
	// Wire name: 'prediction_proba_col'
	PredictionProbaCol string ``
	// Problem type the model aims to solve.
	// Wire name: 'problem_type'
	ProblemType MonitorInferenceLogProblemType ``
	// Column for the timestamp.
	// Wire name: 'timestamp_col'
	TimestampCol    string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *MonitorInferenceLog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MonitorInferenceLog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func MonitorInferenceLogToPb(st *MonitorInferenceLog) (*catalogpb.MonitorInferenceLogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.MonitorInferenceLogPb{}
	pb.Granularities = st.Granularities
	pb.LabelCol = st.LabelCol
	pb.ModelIdCol = st.ModelIdCol
	pb.PredictionCol = st.PredictionCol
	pb.PredictionProbaCol = st.PredictionProbaCol
	problemTypePb, err := MonitorInferenceLogProblemTypeToPb(&st.ProblemType)
	if err != nil {
		return nil, err
	}
	if problemTypePb != nil {
		pb.ProblemType = *problemTypePb
	}
	pb.TimestampCol = st.TimestampCol

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func MonitorInferenceLogFromPb(pb *catalogpb.MonitorInferenceLogPb) (*MonitorInferenceLog, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorInferenceLog{}
	st.Granularities = pb.Granularities
	st.LabelCol = pb.LabelCol
	st.ModelIdCol = pb.ModelIdCol
	st.PredictionCol = pb.PredictionCol
	st.PredictionProbaCol = pb.PredictionProbaCol
	problemTypeField, err := MonitorInferenceLogProblemTypeFromPb(&pb.ProblemType)
	if err != nil {
		return nil, err
	}
	if problemTypeField != nil {
		st.ProblemType = *problemTypeField
	}
	st.TimestampCol = pb.TimestampCol

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type MonitorInferenceLogProblemType string

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

// Values returns all possible values for MonitorInferenceLogProblemType.
//
// There is no guarantee on the order of the values in the slice.
func (f *MonitorInferenceLogProblemType) Values() []MonitorInferenceLogProblemType {
	return []MonitorInferenceLogProblemType{
		MonitorInferenceLogProblemTypeProblemTypeClassification,
		MonitorInferenceLogProblemTypeProblemTypeRegression,
	}
}

// Type always returns MonitorInferenceLogProblemType to satisfy [pflag.Value] interface
func (f *MonitorInferenceLogProblemType) Type() string {
	return "MonitorInferenceLogProblemType"
}

func MonitorInferenceLogProblemTypeToPb(st *MonitorInferenceLogProblemType) (*catalogpb.MonitorInferenceLogProblemTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.MonitorInferenceLogProblemTypePb(*st)
	return &pb, nil
}

func MonitorInferenceLogProblemTypeFromPb(pb *catalogpb.MonitorInferenceLogProblemTypePb) (*MonitorInferenceLogProblemType, error) {
	if pb == nil {
		return nil, nil
	}
	st := MonitorInferenceLogProblemType(*pb)
	return &st, nil
}

type MonitorInfo struct {
	// [Create:REQ Update:IGN] Field for specifying the absolute path to a
	// custom directory to store data-monitoring assets. Normally prepopulated
	// to a default user location via UI and Python APIs.
	// Wire name: 'assets_dir'
	AssetsDir string ``
	// [Create:OPT Update:OPT] Baseline table name. Baseline data is used to
	// compute drift from the data in the monitored `table_name`. The baseline
	// table and the monitored table shall have the same schema.
	// Wire name: 'baseline_table_name'
	BaselineTableName string ``
	// [Create:OPT Update:OPT] Custom metrics.
	// Wire name: 'custom_metrics'
	CustomMetrics []MonitorMetric ``
	// [Create:ERR Update:OPT] Id of dashboard that visualizes the computed
	// metrics. This can be empty if the monitor is in PENDING state.
	// Wire name: 'dashboard_id'
	DashboardId string ``
	// [Create:OPT Update:OPT] Data classification related config.
	// Wire name: 'data_classification_config'
	DataClassificationConfig *MonitorDataClassificationConfig ``
	// [Create:ERR Update:IGN] Table that stores drift metrics data. Format:
	// `catalog.schema.table_name`.
	// Wire name: 'drift_metrics_table_name'
	DriftMetricsTableName string ``

	// Wire name: 'inference_log'
	InferenceLog *MonitorInferenceLog ``
	// [Create:ERR Update:IGN] The latest error message for a monitor failure.
	// Wire name: 'latest_monitor_failure_msg'
	LatestMonitorFailureMsg string ``
	// [Create:ERR Update:IGN] Represents the current monitor configuration
	// version in use. The version will be represented in a numeric fashion
	// (1,2,3...). The field has flexibility to take on negative values, which
	// can indicate corrupted monitor_version numbers.
	// Wire name: 'monitor_version'
	MonitorVersion int64 ``
	// [Create:OPT Update:OPT] Field for specifying notification settings.
	// Wire name: 'notifications'
	Notifications *MonitorNotifications ``
	// [Create:REQ Update:REQ] Schema where output tables are created. Needs to
	// be in 2-level format {catalog}.{schema}
	// Wire name: 'output_schema_name'
	OutputSchemaName string ``
	// [Create:ERR Update:IGN] Table that stores profile metrics data. Format:
	// `catalog.schema.table_name`.
	// Wire name: 'profile_metrics_table_name'
	ProfileMetricsTableName string ``
	// [Create:OPT Update:OPT] The monitor schedule.
	// Wire name: 'schedule'
	Schedule *MonitorCronSchedule ``
	// [Create:OPT Update:OPT] List of column expressions to slice data with for
	// targeted analysis. The data is grouped by each expression independently,
	// resulting in a separate slice for each predicate and its complements. For
	// example `slicing_exprs=[“col_1”, “col_2 > 10”]` will generate the
	// following slices: two slices for `col_2 > 10` (True and False), and one
	// slice per unique value in `col1`. For high-cardinality columns, only the
	// top 100 unique values by frequency will generate slices.
	// Wire name: 'slicing_exprs'
	SlicingExprs []string ``
	// Configuration for monitoring snapshot tables.
	// Wire name: 'snapshot'
	Snapshot *MonitorSnapshot ``
	// [Create:ERR Update:IGN] The monitor status.
	// Wire name: 'status'
	Status MonitorInfoStatus ``
	// [Create:ERR Update:IGN] UC table to monitor. Format:
	// `catalog.schema.table_name`
	// Wire name: 'table_name'
	TableName string ``
	// Configuration for monitoring time series tables.
	// Wire name: 'time_series'
	TimeSeries      *MonitorTimeSeries ``
	ForceSendFields []string           `tf:"-"`
}

func (s *MonitorInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MonitorInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func MonitorInfoToPb(st *MonitorInfo) (*catalogpb.MonitorInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.MonitorInfoPb{}
	pb.AssetsDir = st.AssetsDir
	pb.BaselineTableName = st.BaselineTableName

	var customMetricsPb []catalogpb.MonitorMetricPb
	for _, item := range st.CustomMetrics {
		itemPb, err := MonitorMetricToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customMetricsPb = append(customMetricsPb, *itemPb)
		}
	}
	pb.CustomMetrics = customMetricsPb
	pb.DashboardId = st.DashboardId
	dataClassificationConfigPb, err := MonitorDataClassificationConfigToPb(st.DataClassificationConfig)
	if err != nil {
		return nil, err
	}
	if dataClassificationConfigPb != nil {
		pb.DataClassificationConfig = dataClassificationConfigPb
	}
	pb.DriftMetricsTableName = st.DriftMetricsTableName
	inferenceLogPb, err := MonitorInferenceLogToPb(st.InferenceLog)
	if err != nil {
		return nil, err
	}
	if inferenceLogPb != nil {
		pb.InferenceLog = inferenceLogPb
	}
	pb.LatestMonitorFailureMsg = st.LatestMonitorFailureMsg
	pb.MonitorVersion = st.MonitorVersion
	notificationsPb, err := MonitorNotificationsToPb(st.Notifications)
	if err != nil {
		return nil, err
	}
	if notificationsPb != nil {
		pb.Notifications = notificationsPb
	}
	pb.OutputSchemaName = st.OutputSchemaName
	pb.ProfileMetricsTableName = st.ProfileMetricsTableName
	schedulePb, err := MonitorCronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}
	pb.SlicingExprs = st.SlicingExprs
	snapshotPb, err := MonitorSnapshotToPb(st.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotPb != nil {
		pb.Snapshot = snapshotPb
	}
	statusPb, err := MonitorInfoStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.TableName = st.TableName
	timeSeriesPb, err := MonitorTimeSeriesToPb(st.TimeSeries)
	if err != nil {
		return nil, err
	}
	if timeSeriesPb != nil {
		pb.TimeSeries = timeSeriesPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func MonitorInfoFromPb(pb *catalogpb.MonitorInfoPb) (*MonitorInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorInfo{}
	st.AssetsDir = pb.AssetsDir
	st.BaselineTableName = pb.BaselineTableName

	var customMetricsField []MonitorMetric
	for _, itemPb := range pb.CustomMetrics {
		item, err := MonitorMetricFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customMetricsField = append(customMetricsField, *item)
		}
	}
	st.CustomMetrics = customMetricsField
	st.DashboardId = pb.DashboardId
	dataClassificationConfigField, err := MonitorDataClassificationConfigFromPb(pb.DataClassificationConfig)
	if err != nil {
		return nil, err
	}
	if dataClassificationConfigField != nil {
		st.DataClassificationConfig = dataClassificationConfigField
	}
	st.DriftMetricsTableName = pb.DriftMetricsTableName
	inferenceLogField, err := MonitorInferenceLogFromPb(pb.InferenceLog)
	if err != nil {
		return nil, err
	}
	if inferenceLogField != nil {
		st.InferenceLog = inferenceLogField
	}
	st.LatestMonitorFailureMsg = pb.LatestMonitorFailureMsg
	st.MonitorVersion = pb.MonitorVersion
	notificationsField, err := MonitorNotificationsFromPb(pb.Notifications)
	if err != nil {
		return nil, err
	}
	if notificationsField != nil {
		st.Notifications = notificationsField
	}
	st.OutputSchemaName = pb.OutputSchemaName
	st.ProfileMetricsTableName = pb.ProfileMetricsTableName
	scheduleField, err := MonitorCronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	st.SlicingExprs = pb.SlicingExprs
	snapshotField, err := MonitorSnapshotFromPb(pb.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotField != nil {
		st.Snapshot = snapshotField
	}
	statusField, err := MonitorInfoStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.TableName = pb.TableName
	timeSeriesField, err := MonitorTimeSeriesFromPb(pb.TimeSeries)
	if err != nil {
		return nil, err
	}
	if timeSeriesField != nil {
		st.TimeSeries = timeSeriesField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type MonitorInfoStatus string

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

// Values returns all possible values for MonitorInfoStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *MonitorInfoStatus) Values() []MonitorInfoStatus {
	return []MonitorInfoStatus{
		MonitorInfoStatusMonitorStatusActive,
		MonitorInfoStatusMonitorStatusDeletePending,
		MonitorInfoStatusMonitorStatusError,
		MonitorInfoStatusMonitorStatusFailed,
		MonitorInfoStatusMonitorStatusPending,
	}
}

// Type always returns MonitorInfoStatus to satisfy [pflag.Value] interface
func (f *MonitorInfoStatus) Type() string {
	return "MonitorInfoStatus"
}

func MonitorInfoStatusToPb(st *MonitorInfoStatus) (*catalogpb.MonitorInfoStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.MonitorInfoStatusPb(*st)
	return &pb, nil
}

func MonitorInfoStatusFromPb(pb *catalogpb.MonitorInfoStatusPb) (*MonitorInfoStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := MonitorInfoStatus(*pb)
	return &st, nil
}

// Custom metric definition.
type MonitorMetric struct {
	// Jinja template for a SQL expression that specifies how to compute the
	// metric. See [create metric definition].
	//
	// [create metric definition]: https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition
	// Wire name: 'definition'
	Definition string ``
	// A list of column names in the input table the metric should be computed
	// for. Can use ``":table"`` to indicate that the metric needs information
	// from multiple columns.
	// Wire name: 'input_columns'
	InputColumns []string ``
	// Name of the metric in the output tables.
	// Wire name: 'name'
	Name string ``
	// The output type of the custom metric.
	// Wire name: 'output_data_type'
	OutputDataType string ``
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
	Type MonitorMetricType ``
}

func MonitorMetricToPb(st *MonitorMetric) (*catalogpb.MonitorMetricPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.MonitorMetricPb{}
	pb.Definition = st.Definition
	pb.InputColumns = st.InputColumns
	pb.Name = st.Name
	pb.OutputDataType = st.OutputDataType
	typePb, err := MonitorMetricTypeToPb(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	return pb, nil
}

func MonitorMetricFromPb(pb *catalogpb.MonitorMetricPb) (*MonitorMetric, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorMetric{}
	st.Definition = pb.Definition
	st.InputColumns = pb.InputColumns
	st.Name = pb.Name
	st.OutputDataType = pb.OutputDataType
	typeField, err := MonitorMetricTypeFromPb(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	return st, nil
}

// Can only be one of “\"CUSTOM_METRIC_TYPE_AGGREGATE\"“,
// “\"CUSTOM_METRIC_TYPE_DERIVED\"“, or “\"CUSTOM_METRIC_TYPE_DRIFT\"“. The
// “\"CUSTOM_METRIC_TYPE_AGGREGATE\"“ and “\"CUSTOM_METRIC_TYPE_DERIVED\"“
// metrics are computed on a single table, whereas the
// “\"CUSTOM_METRIC_TYPE_DRIFT\"“ compare metrics across baseline and input
// table, or across the two consecutive time windows. -
// CUSTOM_METRIC_TYPE_AGGREGATE: only depend on the existing columns in your
// table - CUSTOM_METRIC_TYPE_DERIVED: depend on previously computed aggregate
// metrics - CUSTOM_METRIC_TYPE_DRIFT: depend on previously computed aggregate
// or derived metrics
type MonitorMetricType string

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

// Values returns all possible values for MonitorMetricType.
//
// There is no guarantee on the order of the values in the slice.
func (f *MonitorMetricType) Values() []MonitorMetricType {
	return []MonitorMetricType{
		MonitorMetricTypeCustomMetricTypeAggregate,
		MonitorMetricTypeCustomMetricTypeDerived,
		MonitorMetricTypeCustomMetricTypeDrift,
	}
}

// Type always returns MonitorMetricType to satisfy [pflag.Value] interface
func (f *MonitorMetricType) Type() string {
	return "MonitorMetricType"
}

func MonitorMetricTypeToPb(st *MonitorMetricType) (*catalogpb.MonitorMetricTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.MonitorMetricTypePb(*st)
	return &pb, nil
}

func MonitorMetricTypeFromPb(pb *catalogpb.MonitorMetricTypePb) (*MonitorMetricType, error) {
	if pb == nil {
		return nil, nil
	}
	st := MonitorMetricType(*pb)
	return &st, nil
}

type MonitorNotifications struct {
	// Destinations to send notifications on failure/timeout.
	// Wire name: 'on_failure'
	OnFailure *MonitorDestination ``
	// Destinations to send notifications on new classification tag detected.
	// Wire name: 'on_new_classification_tag_detected'
	OnNewClassificationTagDetected *MonitorDestination ``
}

func MonitorNotificationsToPb(st *MonitorNotifications) (*catalogpb.MonitorNotificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.MonitorNotificationsPb{}
	onFailurePb, err := MonitorDestinationToPb(st.OnFailure)
	if err != nil {
		return nil, err
	}
	if onFailurePb != nil {
		pb.OnFailure = onFailurePb
	}
	onNewClassificationTagDetectedPb, err := MonitorDestinationToPb(st.OnNewClassificationTagDetected)
	if err != nil {
		return nil, err
	}
	if onNewClassificationTagDetectedPb != nil {
		pb.OnNewClassificationTagDetected = onNewClassificationTagDetectedPb
	}

	return pb, nil
}

func MonitorNotificationsFromPb(pb *catalogpb.MonitorNotificationsPb) (*MonitorNotifications, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorNotifications{}
	onFailureField, err := MonitorDestinationFromPb(pb.OnFailure)
	if err != nil {
		return nil, err
	}
	if onFailureField != nil {
		st.OnFailure = onFailureField
	}
	onNewClassificationTagDetectedField, err := MonitorDestinationFromPb(pb.OnNewClassificationTagDetected)
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
	EndTimeMs int64 ``
	// An optional message to give insight into the current state of the job
	// (e.g. FAILURE messages).
	// Wire name: 'message'
	Message string ``
	// Unique id of the refresh operation.
	// Wire name: 'refresh_id'
	RefreshId int64 ``
	// Time at which refresh operation was initiated (milliseconds since
	// 1/1/1970 UTC).
	// Wire name: 'start_time_ms'
	StartTimeMs int64 ``
	// The current state of the refresh.
	// Wire name: 'state'
	State MonitorRefreshInfoState ``
	// The method by which the refresh was triggered.
	// Wire name: 'trigger'
	Trigger         MonitorRefreshInfoTrigger ``
	ForceSendFields []string                  `tf:"-"`
}

func (s *MonitorRefreshInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MonitorRefreshInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func MonitorRefreshInfoToPb(st *MonitorRefreshInfo) (*catalogpb.MonitorRefreshInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.MonitorRefreshInfoPb{}
	pb.EndTimeMs = st.EndTimeMs
	pb.Message = st.Message
	pb.RefreshId = st.RefreshId
	pb.StartTimeMs = st.StartTimeMs
	statePb, err := MonitorRefreshInfoStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}
	triggerPb, err := MonitorRefreshInfoTriggerToPb(&st.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerPb != nil {
		pb.Trigger = *triggerPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func MonitorRefreshInfoFromPb(pb *catalogpb.MonitorRefreshInfoPb) (*MonitorRefreshInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorRefreshInfo{}
	st.EndTimeMs = pb.EndTimeMs
	st.Message = pb.Message
	st.RefreshId = pb.RefreshId
	st.StartTimeMs = pb.StartTimeMs
	stateField, err := MonitorRefreshInfoStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	triggerField, err := MonitorRefreshInfoTriggerFromPb(&pb.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerField != nil {
		st.Trigger = *triggerField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// The current state of the refresh.
type MonitorRefreshInfoState string

const MonitorRefreshInfoStateCanceled MonitorRefreshInfoState = `CANCELED`

const MonitorRefreshInfoStateFailed MonitorRefreshInfoState = `FAILED`

const MonitorRefreshInfoStatePending MonitorRefreshInfoState = `PENDING`

const MonitorRefreshInfoStateRunning MonitorRefreshInfoState = `RUNNING`

const MonitorRefreshInfoStateSuccess MonitorRefreshInfoState = `SUCCESS`

const MonitorRefreshInfoStateUnknown MonitorRefreshInfoState = `UNKNOWN`

// String representation for [fmt.Print]
func (f *MonitorRefreshInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorRefreshInfoState) Set(v string) error {
	switch v {
	case `CANCELED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCESS`, `UNKNOWN`:
		*f = MonitorRefreshInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "FAILED", "PENDING", "RUNNING", "SUCCESS", "UNKNOWN"`, v)
	}
}

// Values returns all possible values for MonitorRefreshInfoState.
//
// There is no guarantee on the order of the values in the slice.
func (f *MonitorRefreshInfoState) Values() []MonitorRefreshInfoState {
	return []MonitorRefreshInfoState{
		MonitorRefreshInfoStateCanceled,
		MonitorRefreshInfoStateFailed,
		MonitorRefreshInfoStatePending,
		MonitorRefreshInfoStateRunning,
		MonitorRefreshInfoStateSuccess,
		MonitorRefreshInfoStateUnknown,
	}
}

// Type always returns MonitorRefreshInfoState to satisfy [pflag.Value] interface
func (f *MonitorRefreshInfoState) Type() string {
	return "MonitorRefreshInfoState"
}

func MonitorRefreshInfoStateToPb(st *MonitorRefreshInfoState) (*catalogpb.MonitorRefreshInfoStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.MonitorRefreshInfoStatePb(*st)
	return &pb, nil
}

func MonitorRefreshInfoStateFromPb(pb *catalogpb.MonitorRefreshInfoStatePb) (*MonitorRefreshInfoState, error) {
	if pb == nil {
		return nil, nil
	}
	st := MonitorRefreshInfoState(*pb)
	return &st, nil
}

type MonitorRefreshInfoTrigger string

const MonitorRefreshInfoTriggerManual MonitorRefreshInfoTrigger = `MANUAL`

const MonitorRefreshInfoTriggerSchedule MonitorRefreshInfoTrigger = `SCHEDULE`

const MonitorRefreshInfoTriggerUnknownTrigger MonitorRefreshInfoTrigger = `UNKNOWN_TRIGGER`

// String representation for [fmt.Print]
func (f *MonitorRefreshInfoTrigger) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MonitorRefreshInfoTrigger) Set(v string) error {
	switch v {
	case `MANUAL`, `SCHEDULE`, `UNKNOWN_TRIGGER`:
		*f = MonitorRefreshInfoTrigger(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MANUAL", "SCHEDULE", "UNKNOWN_TRIGGER"`, v)
	}
}

// Values returns all possible values for MonitorRefreshInfoTrigger.
//
// There is no guarantee on the order of the values in the slice.
func (f *MonitorRefreshInfoTrigger) Values() []MonitorRefreshInfoTrigger {
	return []MonitorRefreshInfoTrigger{
		MonitorRefreshInfoTriggerManual,
		MonitorRefreshInfoTriggerSchedule,
		MonitorRefreshInfoTriggerUnknownTrigger,
	}
}

// Type always returns MonitorRefreshInfoTrigger to satisfy [pflag.Value] interface
func (f *MonitorRefreshInfoTrigger) Type() string {
	return "MonitorRefreshInfoTrigger"
}

func MonitorRefreshInfoTriggerToPb(st *MonitorRefreshInfoTrigger) (*catalogpb.MonitorRefreshInfoTriggerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.MonitorRefreshInfoTriggerPb(*st)
	return &pb, nil
}

func MonitorRefreshInfoTriggerFromPb(pb *catalogpb.MonitorRefreshInfoTriggerPb) (*MonitorRefreshInfoTrigger, error) {
	if pb == nil {
		return nil, nil
	}
	st := MonitorRefreshInfoTrigger(*pb)
	return &st, nil
}

type MonitorRefreshListResponse struct {
	// List of refreshes.
	// Wire name: 'refreshes'
	Refreshes []MonitorRefreshInfo ``
}

func MonitorRefreshListResponseToPb(st *MonitorRefreshListResponse) (*catalogpb.MonitorRefreshListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.MonitorRefreshListResponsePb{}

	var refreshesPb []catalogpb.MonitorRefreshInfoPb
	for _, item := range st.Refreshes {
		itemPb, err := MonitorRefreshInfoToPb(&item)
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

func MonitorRefreshListResponseFromPb(pb *catalogpb.MonitorRefreshListResponsePb) (*MonitorRefreshListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorRefreshListResponse{}

	var refreshesField []MonitorRefreshInfo
	for _, itemPb := range pb.Refreshes {
		item, err := MonitorRefreshInfoFromPb(&itemPb)
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

// Snapshot analysis configuration
type MonitorSnapshot struct {
}

func MonitorSnapshotToPb(st *MonitorSnapshot) (*catalogpb.MonitorSnapshotPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.MonitorSnapshotPb{}

	return pb, nil
}

func MonitorSnapshotFromPb(pb *catalogpb.MonitorSnapshotPb) (*MonitorSnapshot, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorSnapshot{}

	return st, nil
}

// Time series analysis configuration.
type MonitorTimeSeries struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``\"5 minutes\"``, ``\"30 minutes\"``, ``\"1 hour\"``, ``\"1 day\"``,
	// ``\"\u003cn\u003e week(s)\"``, ``\"1 month\"``, ``\"1 year\"``}.
	// Wire name: 'granularities'
	Granularities []string ``
	// Column for the timestamp.
	// Wire name: 'timestamp_col'
	TimestampCol string ``
}

func MonitorTimeSeriesToPb(st *MonitorTimeSeries) (*catalogpb.MonitorTimeSeriesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.MonitorTimeSeriesPb{}
	pb.Granularities = st.Granularities
	pb.TimestampCol = st.TimestampCol

	return pb, nil
}

func MonitorTimeSeriesFromPb(pb *catalogpb.MonitorTimeSeriesPb) (*MonitorTimeSeries, error) {
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
	Name string ``
}

func NamedTableConstraintToPb(st *NamedTableConstraint) (*catalogpb.NamedTableConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.NamedTableConstraintPb{}
	pb.Name = st.Name

	return pb, nil
}

func NamedTableConstraintFromPb(pb *catalogpb.NamedTableConstraintPb) (*NamedTableConstraint, error) {
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
	Name string ``
	// Specification of the online table.
	// Wire name: 'spec'
	Spec *OnlineTableSpec ``
	// Online Table data synchronization status
	// Wire name: 'status'
	Status *OnlineTableStatus ``
	// Data serving REST API URL for this table
	// Wire name: 'table_serving_url'
	TableServingUrl string ``
	// The provisioning state of the online table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	// Wire name: 'unity_catalog_provisioning_state'
	UnityCatalogProvisioningState ProvisioningInfoState ``
	ForceSendFields               []string              `tf:"-"`
}

func (s *OnlineTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OnlineTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func OnlineTableToPb(st *OnlineTable) (*catalogpb.OnlineTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.OnlineTablePb{}
	pb.Name = st.Name
	specPb, err := OnlineTableSpecToPb(st.Spec)
	if err != nil {
		return nil, err
	}
	if specPb != nil {
		pb.Spec = specPb
	}
	statusPb, err := OnlineTableStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}
	pb.TableServingUrl = st.TableServingUrl
	unityCatalogProvisioningStatePb, err := ProvisioningInfoStateToPb(&st.UnityCatalogProvisioningState)
	if err != nil {
		return nil, err
	}
	if unityCatalogProvisioningStatePb != nil {
		pb.UnityCatalogProvisioningState = *unityCatalogProvisioningStatePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func OnlineTableFromPb(pb *catalogpb.OnlineTablePb) (*OnlineTable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineTable{}
	st.Name = pb.Name
	specField, err := OnlineTableSpecFromPb(pb.Spec)
	if err != nil {
		return nil, err
	}
	if specField != nil {
		st.Spec = specField
	}
	statusField, err := OnlineTableStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}
	st.TableServingUrl = pb.TableServingUrl
	unityCatalogProvisioningStateField, err := ProvisioningInfoStateFromPb(&pb.UnityCatalogProvisioningState)
	if err != nil {
		return nil, err
	}
	if unityCatalogProvisioningStateField != nil {
		st.UnityCatalogProvisioningState = *unityCatalogProvisioningStateField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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
	PerformFullCopy bool ``
	// ID of the associated pipeline. Generated by the server - cannot be set by
	// the caller.
	// Wire name: 'pipeline_id'
	PipelineId string ``
	// Primary Key columns to be used for data insert/update in the destination.
	// Wire name: 'primary_key_columns'
	PrimaryKeyColumns []string ``
	// Pipeline runs continuously after generating the initial data.
	// Wire name: 'run_continuously'
	RunContinuously *OnlineTableSpecContinuousSchedulingPolicy ``
	// Pipeline stops after generating the initial data and can be triggered
	// later (manually, through a cron job or through data triggers)
	// Wire name: 'run_triggered'
	RunTriggered *OnlineTableSpecTriggeredSchedulingPolicy ``
	// Three-part (catalog, schema, table) name of the source Delta table.
	// Wire name: 'source_table_full_name'
	SourceTableFullName string ``
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	// Wire name: 'timeseries_key'
	TimeseriesKey   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *OnlineTableSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OnlineTableSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func OnlineTableSpecToPb(st *OnlineTableSpec) (*catalogpb.OnlineTableSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.OnlineTableSpecPb{}
	pb.PerformFullCopy = st.PerformFullCopy
	pb.PipelineId = st.PipelineId
	pb.PrimaryKeyColumns = st.PrimaryKeyColumns
	runContinuouslyPb, err := OnlineTableSpecContinuousSchedulingPolicyToPb(st.RunContinuously)
	if err != nil {
		return nil, err
	}
	if runContinuouslyPb != nil {
		pb.RunContinuously = runContinuouslyPb
	}
	runTriggeredPb, err := OnlineTableSpecTriggeredSchedulingPolicyToPb(st.RunTriggered)
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

func OnlineTableSpecFromPb(pb *catalogpb.OnlineTableSpecPb) (*OnlineTableSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineTableSpec{}
	st.PerformFullCopy = pb.PerformFullCopy
	st.PipelineId = pb.PipelineId
	st.PrimaryKeyColumns = pb.PrimaryKeyColumns
	runContinuouslyField, err := OnlineTableSpecContinuousSchedulingPolicyFromPb(pb.RunContinuously)
	if err != nil {
		return nil, err
	}
	if runContinuouslyField != nil {
		st.RunContinuously = runContinuouslyField
	}
	runTriggeredField, err := OnlineTableSpecTriggeredSchedulingPolicyFromPb(pb.RunTriggered)
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

type OnlineTableSpecContinuousSchedulingPolicy struct {
}

func OnlineTableSpecContinuousSchedulingPolicyToPb(st *OnlineTableSpecContinuousSchedulingPolicy) (*catalogpb.OnlineTableSpecContinuousSchedulingPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.OnlineTableSpecContinuousSchedulingPolicyPb{}

	return pb, nil
}

func OnlineTableSpecContinuousSchedulingPolicyFromPb(pb *catalogpb.OnlineTableSpecContinuousSchedulingPolicyPb) (*OnlineTableSpecContinuousSchedulingPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineTableSpecContinuousSchedulingPolicy{}

	return st, nil
}

type OnlineTableSpecTriggeredSchedulingPolicy struct {
}

func OnlineTableSpecTriggeredSchedulingPolicyToPb(st *OnlineTableSpecTriggeredSchedulingPolicy) (*catalogpb.OnlineTableSpecTriggeredSchedulingPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.OnlineTableSpecTriggeredSchedulingPolicyPb{}

	return pb, nil
}

func OnlineTableSpecTriggeredSchedulingPolicyFromPb(pb *catalogpb.OnlineTableSpecTriggeredSchedulingPolicyPb) (*OnlineTableSpecTriggeredSchedulingPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineTableSpecTriggeredSchedulingPolicy{}

	return st, nil
}

// The state of an online table.
type OnlineTableState string

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

// Values returns all possible values for OnlineTableState.
//
// There is no guarantee on the order of the values in the slice.
func (f *OnlineTableState) Values() []OnlineTableState {
	return []OnlineTableState{
		OnlineTableStateOffline,
		OnlineTableStateOfflineFailed,
		OnlineTableStateOnline,
		OnlineTableStateOnlineContinuousUpdate,
		OnlineTableStateOnlineNoPendingUpdate,
		OnlineTableStateOnlinePipelineFailed,
		OnlineTableStateOnlineTriggeredUpdate,
		OnlineTableStateOnlineUpdatingPipelineResources,
		OnlineTableStateProvisioning,
		OnlineTableStateProvisioningInitialSnapshot,
		OnlineTableStateProvisioningPipelineResources,
	}
}

// Type always returns OnlineTableState to satisfy [pflag.Value] interface
func (f *OnlineTableState) Type() string {
	return "OnlineTableState"
}

func OnlineTableStateToPb(st *OnlineTableState) (*catalogpb.OnlineTableStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.OnlineTableStatePb(*st)
	return &pb, nil
}

func OnlineTableStateFromPb(pb *catalogpb.OnlineTableStatePb) (*OnlineTableState, error) {
	if pb == nil {
		return nil, nil
	}
	st := OnlineTableState(*pb)
	return &st, nil
}

// Status of an online table.
type OnlineTableStatus struct {

	// Wire name: 'continuous_update_status'
	ContinuousUpdateStatus *ContinuousUpdateStatus ``
	// The state of the online table.
	// Wire name: 'detailed_state'
	DetailedState OnlineTableState ``

	// Wire name: 'failed_status'
	FailedStatus *FailedStatus ``
	// A text description of the current state of the online table.
	// Wire name: 'message'
	Message string ``

	// Wire name: 'provisioning_status'
	ProvisioningStatus *ProvisioningStatus ``

	// Wire name: 'triggered_update_status'
	TriggeredUpdateStatus *TriggeredUpdateStatus ``
	ForceSendFields       []string               `tf:"-"`
}

func (s *OnlineTableStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OnlineTableStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func OnlineTableStatusToPb(st *OnlineTableStatus) (*catalogpb.OnlineTableStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.OnlineTableStatusPb{}
	continuousUpdateStatusPb, err := ContinuousUpdateStatusToPb(st.ContinuousUpdateStatus)
	if err != nil {
		return nil, err
	}
	if continuousUpdateStatusPb != nil {
		pb.ContinuousUpdateStatus = continuousUpdateStatusPb
	}
	detailedStatePb, err := OnlineTableStateToPb(&st.DetailedState)
	if err != nil {
		return nil, err
	}
	if detailedStatePb != nil {
		pb.DetailedState = *detailedStatePb
	}
	failedStatusPb, err := FailedStatusToPb(st.FailedStatus)
	if err != nil {
		return nil, err
	}
	if failedStatusPb != nil {
		pb.FailedStatus = failedStatusPb
	}
	pb.Message = st.Message
	provisioningStatusPb, err := ProvisioningStatusToPb(st.ProvisioningStatus)
	if err != nil {
		return nil, err
	}
	if provisioningStatusPb != nil {
		pb.ProvisioningStatus = provisioningStatusPb
	}
	triggeredUpdateStatusPb, err := TriggeredUpdateStatusToPb(st.TriggeredUpdateStatus)
	if err != nil {
		return nil, err
	}
	if triggeredUpdateStatusPb != nil {
		pb.TriggeredUpdateStatus = triggeredUpdateStatusPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func OnlineTableStatusFromPb(pb *catalogpb.OnlineTableStatusPb) (*OnlineTableStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineTableStatus{}
	continuousUpdateStatusField, err := ContinuousUpdateStatusFromPb(pb.ContinuousUpdateStatus)
	if err != nil {
		return nil, err
	}
	if continuousUpdateStatusField != nil {
		st.ContinuousUpdateStatus = continuousUpdateStatusField
	}
	detailedStateField, err := OnlineTableStateFromPb(&pb.DetailedState)
	if err != nil {
		return nil, err
	}
	if detailedStateField != nil {
		st.DetailedState = *detailedStateField
	}
	failedStatusField, err := FailedStatusFromPb(pb.FailedStatus)
	if err != nil {
		return nil, err
	}
	if failedStatusField != nil {
		st.FailedStatus = failedStatusField
	}
	st.Message = pb.Message
	provisioningStatusField, err := ProvisioningStatusFromPb(pb.ProvisioningStatus)
	if err != nil {
		return nil, err
	}
	if provisioningStatusField != nil {
		st.ProvisioningStatus = provisioningStatusField
	}
	triggeredUpdateStatusField, err := TriggeredUpdateStatusFromPb(pb.TriggeredUpdateStatus)
	if err != nil {
		return nil, err
	}
	if triggeredUpdateStatusField != nil {
		st.TriggeredUpdateStatus = triggeredUpdateStatusField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Spec of an allowed option on a securable kind and its attributes. This is
// mostly used by UI to provide user friendly hints and descriptions in order to
// facilitate the securable creation process.
type OptionSpec struct {
	// For drop down / radio button selections, UI will want to know the
	// possible input values, it can also be used by other option types to limit
	// input selections.
	// Wire name: 'allowed_values'
	AllowedValues []string ``
	// The default value of the option, for example, value '443' for 'port'
	// option.
	// Wire name: 'default_value'
	DefaultValue string ``
	// A concise user facing description of what the input value of this option
	// should look like.
	// Wire name: 'description'
	Description string ``
	// The hint is used on the UI to suggest what the input value can possibly
	// be like, for example: example.com for 'host' option. Unlike default
	// value, it will not be applied automatically without user input.
	// Wire name: 'hint'
	Hint string ``
	// Indicates whether an option should be displayed with copy button on the
	// UI.
	// Wire name: 'is_copiable'
	IsCopiable bool ``
	// Indicates whether an option can be provided by users in the create/update
	// path of an entity.
	// Wire name: 'is_creatable'
	IsCreatable bool ``
	// Is the option value not user settable and is thus not shown on the UI.
	// Wire name: 'is_hidden'
	IsHidden bool ``
	// Specifies whether this option is safe to log, i.e. no sensitive
	// information.
	// Wire name: 'is_loggable'
	IsLoggable bool ``
	// Is the option required.
	// Wire name: 'is_required'
	IsRequired bool ``
	// Is the option value considered secret and thus redacted on the UI.
	// Wire name: 'is_secret'
	IsSecret bool ``
	// Is the option updatable by users.
	// Wire name: 'is_updatable'
	IsUpdatable bool ``
	// The unique name of the option.
	// Wire name: 'name'
	Name string ``
	// Specifies when the option value is displayed on the UI within the OAuth
	// flow.
	// Wire name: 'oauth_stage'
	OauthStage OptionSpecOauthStage ``
	// The type of the option.
	// Wire name: 'type'
	Type            OptionSpecOptionType ``
	ForceSendFields []string             `tf:"-"`
}

func (s *OptionSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OptionSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func OptionSpecToPb(st *OptionSpec) (*catalogpb.OptionSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.OptionSpecPb{}
	pb.AllowedValues = st.AllowedValues
	pb.DefaultValue = st.DefaultValue
	pb.Description = st.Description
	pb.Hint = st.Hint
	pb.IsCopiable = st.IsCopiable
	pb.IsCreatable = st.IsCreatable
	pb.IsHidden = st.IsHidden
	pb.IsLoggable = st.IsLoggable
	pb.IsRequired = st.IsRequired
	pb.IsSecret = st.IsSecret
	pb.IsUpdatable = st.IsUpdatable
	pb.Name = st.Name
	oauthStagePb, err := OptionSpecOauthStageToPb(&st.OauthStage)
	if err != nil {
		return nil, err
	}
	if oauthStagePb != nil {
		pb.OauthStage = *oauthStagePb
	}
	typePb, err := OptionSpecOptionTypeToPb(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func OptionSpecFromPb(pb *catalogpb.OptionSpecPb) (*OptionSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OptionSpec{}
	st.AllowedValues = pb.AllowedValues
	st.DefaultValue = pb.DefaultValue
	st.Description = pb.Description
	st.Hint = pb.Hint
	st.IsCopiable = pb.IsCopiable
	st.IsCreatable = pb.IsCreatable
	st.IsHidden = pb.IsHidden
	st.IsLoggable = pb.IsLoggable
	st.IsRequired = pb.IsRequired
	st.IsSecret = pb.IsSecret
	st.IsUpdatable = pb.IsUpdatable
	st.Name = pb.Name
	oauthStageField, err := OptionSpecOauthStageFromPb(&pb.OauthStage)
	if err != nil {
		return nil, err
	}
	if oauthStageField != nil {
		st.OauthStage = *oauthStageField
	}
	typeField, err := OptionSpecOptionTypeFromPb(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// During the OAuth flow, specifies which stage the option should be displayed
// in the UI. OAUTH_STAGE_UNSPECIFIED is the default value for options unrelated
// to the OAuth flow. BEFORE_AUTHORIZATION_CODE corresponds to options necessary
// to initiate the OAuth process. BEFORE_ACCESS_TOKEN corresponds to options
// that are necessary to create a foreign connection, but that should be
// displayed after the authorization code has already been received.
type OptionSpecOauthStage string

const OptionSpecOauthStageBeforeAccessToken OptionSpecOauthStage = `BEFORE_ACCESS_TOKEN`

const OptionSpecOauthStageBeforeAuthorizationCode OptionSpecOauthStage = `BEFORE_AUTHORIZATION_CODE`

// String representation for [fmt.Print]
func (f *OptionSpecOauthStage) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OptionSpecOauthStage) Set(v string) error {
	switch v {
	case `BEFORE_ACCESS_TOKEN`, `BEFORE_AUTHORIZATION_CODE`:
		*f = OptionSpecOauthStage(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BEFORE_ACCESS_TOKEN", "BEFORE_AUTHORIZATION_CODE"`, v)
	}
}

// Values returns all possible values for OptionSpecOauthStage.
//
// There is no guarantee on the order of the values in the slice.
func (f *OptionSpecOauthStage) Values() []OptionSpecOauthStage {
	return []OptionSpecOauthStage{
		OptionSpecOauthStageBeforeAccessToken,
		OptionSpecOauthStageBeforeAuthorizationCode,
	}
}

// Type always returns OptionSpecOauthStage to satisfy [pflag.Value] interface
func (f *OptionSpecOauthStage) Type() string {
	return "OptionSpecOauthStage"
}

func OptionSpecOauthStageToPb(st *OptionSpecOauthStage) (*catalogpb.OptionSpecOauthStagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.OptionSpecOauthStagePb(*st)
	return &pb, nil
}

func OptionSpecOauthStageFromPb(pb *catalogpb.OptionSpecOauthStagePb) (*OptionSpecOauthStage, error) {
	if pb == nil {
		return nil, nil
	}
	st := OptionSpecOauthStage(*pb)
	return &st, nil
}

// Type of the option, we purposely follow JavaScript types so that the UI can
// map the options to JS types. https://www.w3schools.com/js/js_datatypes.asp
// Enum is a special case that it's just string with selections.
type OptionSpecOptionType string

const OptionSpecOptionTypeOptionBigint OptionSpecOptionType = `OPTION_BIGINT`

const OptionSpecOptionTypeOptionBoolean OptionSpecOptionType = `OPTION_BOOLEAN`

const OptionSpecOptionTypeOptionEnum OptionSpecOptionType = `OPTION_ENUM`

const OptionSpecOptionTypeOptionMultilineString OptionSpecOptionType = `OPTION_MULTILINE_STRING`

const OptionSpecOptionTypeOptionNumber OptionSpecOptionType = `OPTION_NUMBER`

const OptionSpecOptionTypeOptionServiceCredential OptionSpecOptionType = `OPTION_SERVICE_CREDENTIAL`

const OptionSpecOptionTypeOptionString OptionSpecOptionType = `OPTION_STRING`

// String representation for [fmt.Print]
func (f *OptionSpecOptionType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OptionSpecOptionType) Set(v string) error {
	switch v {
	case `OPTION_BIGINT`, `OPTION_BOOLEAN`, `OPTION_ENUM`, `OPTION_MULTILINE_STRING`, `OPTION_NUMBER`, `OPTION_SERVICE_CREDENTIAL`, `OPTION_STRING`:
		*f = OptionSpecOptionType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OPTION_BIGINT", "OPTION_BOOLEAN", "OPTION_ENUM", "OPTION_MULTILINE_STRING", "OPTION_NUMBER", "OPTION_SERVICE_CREDENTIAL", "OPTION_STRING"`, v)
	}
}

// Values returns all possible values for OptionSpecOptionType.
//
// There is no guarantee on the order of the values in the slice.
func (f *OptionSpecOptionType) Values() []OptionSpecOptionType {
	return []OptionSpecOptionType{
		OptionSpecOptionTypeOptionBigint,
		OptionSpecOptionTypeOptionBoolean,
		OptionSpecOptionTypeOptionEnum,
		OptionSpecOptionTypeOptionMultilineString,
		OptionSpecOptionTypeOptionNumber,
		OptionSpecOptionTypeOptionServiceCredential,
		OptionSpecOptionTypeOptionString,
	}
}

// Type always returns OptionSpecOptionType to satisfy [pflag.Value] interface
func (f *OptionSpecOptionType) Type() string {
	return "OptionSpecOptionType"
}

func OptionSpecOptionTypeToPb(st *OptionSpecOptionType) (*catalogpb.OptionSpecOptionTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.OptionSpecOptionTypePb(*st)
	return &pb, nil
}

func OptionSpecOptionTypeFromPb(pb *catalogpb.OptionSpecOptionTypePb) (*OptionSpecOptionType, error) {
	if pb == nil {
		return nil, nil
	}
	st := OptionSpecOptionType(*pb)
	return &st, nil
}

type PermissionsChange struct {
	// The set of privileges to add.
	// Wire name: 'add'
	Add []Privilege ``
	// The principal whose privileges we are changing. Only one of principal or
	// principal_id should be specified, never both at the same time.
	// Wire name: 'principal'
	Principal string ``
	// The set of privileges to remove.
	// Wire name: 'remove'
	Remove          []Privilege ``
	ForceSendFields []string    `tf:"-"`
}

func (s *PermissionsChange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PermissionsChange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PermissionsChangeToPb(st *PermissionsChange) (*catalogpb.PermissionsChangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.PermissionsChangePb{}

	var addPb []catalogpb.PrivilegePb
	for _, item := range st.Add {
		itemPb, err := PrivilegeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			addPb = append(addPb, *itemPb)
		}
	}
	pb.Add = addPb
	pb.Principal = st.Principal

	var removePb []catalogpb.PrivilegePb
	for _, item := range st.Remove {
		itemPb, err := PrivilegeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			removePb = append(removePb, *itemPb)
		}
	}
	pb.Remove = removePb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PermissionsChangeFromPb(pb *catalogpb.PermissionsChangePb) (*PermissionsChange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionsChange{}

	var addField []Privilege
	for _, itemPb := range pb.Add {
		item, err := PrivilegeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			addField = append(addField, *item)
		}
	}
	st.Add = addField
	st.Principal = pb.Principal

	var removeField []Privilege
	for _, itemPb := range pb.Remove {
		item, err := PrivilegeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			removeField = append(removeField, *item)
		}
	}
	st.Remove = removeField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Progress information of the Online Table data synchronization pipeline.
type PipelineProgress struct {
	// The estimated time remaining to complete this update in seconds.
	// Wire name: 'estimated_completion_time_seconds'
	EstimatedCompletionTimeSeconds float64 ``
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	// Wire name: 'latest_version_currently_processing'
	LatestVersionCurrentlyProcessing int64 ``
	// The completion ratio of this update. This is a number between 0 and 1.
	// Wire name: 'sync_progress_completion'
	SyncProgressCompletion float64 ``
	// The number of rows that have been synced in this update.
	// Wire name: 'synced_row_count'
	SyncedRowCount int64 ``
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	// Wire name: 'total_row_count'
	TotalRowCount   int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *PipelineProgress) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineProgress) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PipelineProgressToPb(st *PipelineProgress) (*catalogpb.PipelineProgressPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.PipelineProgressPb{}
	pb.EstimatedCompletionTimeSeconds = st.EstimatedCompletionTimeSeconds
	pb.LatestVersionCurrentlyProcessing = st.LatestVersionCurrentlyProcessing
	pb.SyncProgressCompletion = st.SyncProgressCompletion
	pb.SyncedRowCount = st.SyncedRowCount
	pb.TotalRowCount = st.TotalRowCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PipelineProgressFromPb(pb *catalogpb.PipelineProgressPb) (*PipelineProgress, error) {
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

type PrimaryKeyConstraint struct {
	// Column names for this constraint.
	// Wire name: 'child_columns'
	ChildColumns []string ``
	// The name of the constraint.
	// Wire name: 'name'
	Name string ``
	// True if the constraint is RELY, false or unset if NORELY.
	// Wire name: 'rely'
	Rely bool ``
	// Column names that represent a timeseries.
	// Wire name: 'timeseries_columns'
	TimeseriesColumns []string ``
	ForceSendFields   []string `tf:"-"`
}

func (s *PrimaryKeyConstraint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrimaryKeyConstraint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PrimaryKeyConstraintToPb(st *PrimaryKeyConstraint) (*catalogpb.PrimaryKeyConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.PrimaryKeyConstraintPb{}
	pb.ChildColumns = st.ChildColumns
	pb.Name = st.Name
	pb.Rely = st.Rely
	pb.TimeseriesColumns = st.TimeseriesColumns

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PrimaryKeyConstraintFromPb(pb *catalogpb.PrimaryKeyConstraintPb) (*PrimaryKeyConstraint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PrimaryKeyConstraint{}
	st.ChildColumns = pb.ChildColumns
	st.Name = pb.Name
	st.Rely = pb.Rely
	st.TimeseriesColumns = pb.TimeseriesColumns

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type Privilege string

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

// Values returns all possible values for Privilege.
//
// There is no guarantee on the order of the values in the slice.
func (f *Privilege) Values() []Privilege {
	return []Privilege{
		PrivilegeAccess,
		PrivilegeAllPrivileges,
		PrivilegeApplyTag,
		PrivilegeBrowse,
		PrivilegeCreate,
		PrivilegeCreateCatalog,
		PrivilegeCreateCleanRoom,
		PrivilegeCreateConnection,
		PrivilegeCreateExternalLocation,
		PrivilegeCreateExternalTable,
		PrivilegeCreateExternalVolume,
		PrivilegeCreateForeignCatalog,
		PrivilegeCreateForeignSecurable,
		PrivilegeCreateFunction,
		PrivilegeCreateManagedStorage,
		PrivilegeCreateMaterializedView,
		PrivilegeCreateModel,
		PrivilegeCreateProvider,
		PrivilegeCreateRecipient,
		PrivilegeCreateSchema,
		PrivilegeCreateServiceCredential,
		PrivilegeCreateShare,
		PrivilegeCreateStorageCredential,
		PrivilegeCreateTable,
		PrivilegeCreateView,
		PrivilegeCreateVolume,
		PrivilegeExecute,
		PrivilegeExecuteCleanRoomTask,
		PrivilegeManage,
		PrivilegeManageAllowlist,
		PrivilegeModify,
		PrivilegeModifyCleanRoom,
		PrivilegeReadFiles,
		PrivilegeReadPrivateFiles,
		PrivilegeReadVolume,
		PrivilegeRefresh,
		PrivilegeSelect,
		PrivilegeSetSharePermission,
		PrivilegeUsage,
		PrivilegeUseCatalog,
		PrivilegeUseConnection,
		PrivilegeUseMarketplaceAssets,
		PrivilegeUseProvider,
		PrivilegeUseRecipient,
		PrivilegeUseSchema,
		PrivilegeUseShare,
		PrivilegeWriteFiles,
		PrivilegeWritePrivateFiles,
		PrivilegeWriteVolume,
	}
}

// Type always returns Privilege to satisfy [pflag.Value] interface
func (f *Privilege) Type() string {
	return "Privilege"
}

func PrivilegeToPb(st *Privilege) (*catalogpb.PrivilegePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.PrivilegePb(*st)
	return &pb, nil
}

func PrivilegeFromPb(pb *catalogpb.PrivilegePb) (*Privilege, error) {
	if pb == nil {
		return nil, nil
	}
	st := Privilege(*pb)
	return &st, nil
}

type PrivilegeAssignment struct {
	// The principal (user email address or group name). For deleted principals,
	// `principal` is empty while `principal_id` is populated.
	// Wire name: 'principal'
	Principal string ``
	// The privileges assigned to the principal.
	// Wire name: 'privileges'
	Privileges      []Privilege ``
	ForceSendFields []string    `tf:"-"`
}

func (s *PrivilegeAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrivilegeAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PrivilegeAssignmentToPb(st *PrivilegeAssignment) (*catalogpb.PrivilegeAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.PrivilegeAssignmentPb{}
	pb.Principal = st.Principal

	var privilegesPb []catalogpb.PrivilegePb
	for _, item := range st.Privileges {
		itemPb, err := PrivilegeToPb(&item)
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

func PrivilegeAssignmentFromPb(pb *catalogpb.PrivilegeAssignmentPb) (*PrivilegeAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PrivilegeAssignment{}
	st.Principal = pb.Principal

	var privilegesField []Privilege
	for _, itemPb := range pb.Privileges {
		item, err := PrivilegeFromPb(&itemPb)
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

// Status of an asynchronously provisioned resource.
type ProvisioningInfo struct {
	// The provisioning state of the resource.
	// Wire name: 'state'
	State ProvisioningInfoState ``
}

func ProvisioningInfoToPb(st *ProvisioningInfo) (*catalogpb.ProvisioningInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ProvisioningInfoPb{}
	statePb, err := ProvisioningInfoStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	return pb, nil
}

func ProvisioningInfoFromPb(pb *catalogpb.ProvisioningInfoPb) (*ProvisioningInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ProvisioningInfo{}
	stateField, err := ProvisioningInfoStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

	return st, nil
}

type ProvisioningInfoState string

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

// Values returns all possible values for ProvisioningInfoState.
//
// There is no guarantee on the order of the values in the slice.
func (f *ProvisioningInfoState) Values() []ProvisioningInfoState {
	return []ProvisioningInfoState{
		ProvisioningInfoStateActive,
		ProvisioningInfoStateDegraded,
		ProvisioningInfoStateDeleting,
		ProvisioningInfoStateFailed,
		ProvisioningInfoStateProvisioning,
		ProvisioningInfoStateUpdating,
	}
}

// Type always returns ProvisioningInfoState to satisfy [pflag.Value] interface
func (f *ProvisioningInfoState) Type() string {
	return "ProvisioningInfoState"
}

func ProvisioningInfoStateToPb(st *ProvisioningInfoState) (*catalogpb.ProvisioningInfoStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.ProvisioningInfoStatePb(*st)
	return &pb, nil
}

func ProvisioningInfoStateFromPb(pb *catalogpb.ProvisioningInfoStatePb) (*ProvisioningInfoState, error) {
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
	InitialPipelineSyncProgress *PipelineProgress ``
}

func ProvisioningStatusToPb(st *ProvisioningStatus) (*catalogpb.ProvisioningStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ProvisioningStatusPb{}
	initialPipelineSyncProgressPb, err := PipelineProgressToPb(st.InitialPipelineSyncProgress)
	if err != nil {
		return nil, err
	}
	if initialPipelineSyncProgressPb != nil {
		pb.InitialPipelineSyncProgress = initialPipelineSyncProgressPb
	}

	return pb, nil
}

func ProvisioningStatusFromPb(pb *catalogpb.ProvisioningStatusPb) (*ProvisioningStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ProvisioningStatus{}
	initialPipelineSyncProgressField, err := PipelineProgressFromPb(pb.InitialPipelineSyncProgress)
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
	LastRefreshedAt int64 ``
	// Name of the parent resource. Returns metastore ID if the parent is a
	// metastore.
	// Wire name: 'parent_full_name'
	ParentFullName string ``
	// The quota parent securable type.
	// Wire name: 'parent_securable_type'
	ParentSecurableType SecurableType ``
	// The current usage of the resource quota.
	// Wire name: 'quota_count'
	QuotaCount int ``
	// The current limit of the resource quota.
	// Wire name: 'quota_limit'
	QuotaLimit int ``
	// The name of the quota.
	// Wire name: 'quota_name'
	QuotaName       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *QuotaInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QuotaInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func QuotaInfoToPb(st *QuotaInfo) (*catalogpb.QuotaInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.QuotaInfoPb{}
	pb.LastRefreshedAt = st.LastRefreshedAt
	pb.ParentFullName = st.ParentFullName
	parentSecurableTypePb, err := SecurableTypeToPb(&st.ParentSecurableType)
	if err != nil {
		return nil, err
	}
	if parentSecurableTypePb != nil {
		pb.ParentSecurableType = *parentSecurableTypePb
	}
	pb.QuotaCount = st.QuotaCount
	pb.QuotaLimit = st.QuotaLimit
	pb.QuotaName = st.QuotaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func QuotaInfoFromPb(pb *catalogpb.QuotaInfoPb) (*QuotaInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QuotaInfo{}
	st.LastRefreshedAt = pb.LastRefreshedAt
	st.ParentFullName = pb.ParentFullName
	parentSecurableTypeField, err := SecurableTypeFromPb(&pb.ParentSecurableType)
	if err != nil {
		return nil, err
	}
	if parentSecurableTypeField != nil {
		st.ParentSecurableType = *parentSecurableTypeField
	}
	st.QuotaCount = pb.QuotaCount
	st.QuotaLimit = pb.QuotaLimit
	st.QuotaName = pb.QuotaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// R2 temporary credentials for API authentication. Read more at
// https://developers.cloudflare.com/r2/api/s3/tokens/.
type R2Credentials struct {
	// The access key ID that identifies the temporary credentials.
	// Wire name: 'access_key_id'
	AccessKeyId string ``
	// The secret access key associated with the access key.
	// Wire name: 'secret_access_key'
	SecretAccessKey string ``
	// The generated JWT that users must pass to use the temporary credentials.
	// Wire name: 'session_token'
	SessionToken    string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *R2Credentials) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s R2Credentials) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func R2CredentialsToPb(st *R2Credentials) (*catalogpb.R2CredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.R2CredentialsPb{}
	pb.AccessKeyId = st.AccessKeyId
	pb.SecretAccessKey = st.SecretAccessKey
	pb.SessionToken = st.SessionToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func R2CredentialsFromPb(pb *catalogpb.R2CredentialsPb) (*R2Credentials, error) {
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

type ReadVolumeRequest struct {
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	// Wire name: 'include_browse'
	IncludeBrowse bool `tf:"-"`
	// The three-level (fully qualified) name of the volume
	// Wire name: 'name'
	Name            string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ReadVolumeRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ReadVolumeRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ReadVolumeRequestToPb(st *ReadVolumeRequest) (*catalogpb.ReadVolumeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ReadVolumeRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ReadVolumeRequestFromPb(pb *catalogpb.ReadVolumeRequestPb) (*ReadVolumeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReadVolumeRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type RegenerateDashboardRequest struct {
	// UC table name in format `catalog.schema.table_name`. This field
	// corresponds to the {full_table_name_arg} arg in the endpoint path.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
	// Optional argument to specify the warehouse for dashboard regeneration. If
	// not specified, the first running warehouse will be used.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *RegenerateDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegenerateDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RegenerateDashboardRequestToPb(st *RegenerateDashboardRequest) (*catalogpb.RegenerateDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.RegenerateDashboardRequestPb{}
	pb.TableName = st.TableName
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RegenerateDashboardRequestFromPb(pb *catalogpb.RegenerateDashboardRequestPb) (*RegenerateDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegenerateDashboardRequest{}
	st.TableName = pb.TableName
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type RegenerateDashboardResponse struct {

	// Wire name: 'dashboard_id'
	DashboardId string ``
	// Parent folder is equivalent to {assets_dir}/{tableName}
	// Wire name: 'parent_folder'
	ParentFolder    string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *RegenerateDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegenerateDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RegenerateDashboardResponseToPb(st *RegenerateDashboardResponse) (*catalogpb.RegenerateDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.RegenerateDashboardResponsePb{}
	pb.DashboardId = st.DashboardId
	pb.ParentFolder = st.ParentFolder

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RegenerateDashboardResponseFromPb(pb *catalogpb.RegenerateDashboardResponsePb) (*RegenerateDashboardResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegenerateDashboardResponse{}
	st.DashboardId = pb.DashboardId
	st.ParentFolder = pb.ParentFolder

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Registered model alias.
type RegisteredModelAlias struct {
	// Name of the alias, e.g. 'champion' or 'latest_stable'
	// Wire name: 'alias_name'
	AliasName string ``
	// Integer version number of the model version to which this alias points.
	// Wire name: 'version_num'
	VersionNum      int      ``
	ForceSendFields []string `tf:"-"`
}

func (s *RegisteredModelAlias) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelAlias) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RegisteredModelAliasToPb(st *RegisteredModelAlias) (*catalogpb.RegisteredModelAliasPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.RegisteredModelAliasPb{}
	pb.AliasName = st.AliasName
	pb.VersionNum = st.VersionNum

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RegisteredModelAliasFromPb(pb *catalogpb.RegisteredModelAliasPb) (*RegisteredModelAlias, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelAlias{}
	st.AliasName = pb.AliasName
	st.VersionNum = pb.VersionNum

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type RegisteredModelInfo struct {
	// List of aliases associated with the registered model
	// Wire name: 'aliases'
	Aliases []RegisteredModelAlias ``
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool ``
	// The name of the catalog where the schema and the registered model reside
	// Wire name: 'catalog_name'
	CatalogName string ``
	// The comment attached to the registered model
	// Wire name: 'comment'
	Comment string ``
	// Creation timestamp of the registered model in milliseconds since the Unix
	// epoch
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// The identifier of the user who created the registered model
	// Wire name: 'created_by'
	CreatedBy string ``
	// The three-level (fully qualified) name of the registered model
	// Wire name: 'full_name'
	FullName string ``
	// The unique identifier of the metastore
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// The name of the registered model
	// Wire name: 'name'
	Name string ``
	// The identifier of the user who owns the registered model
	// Wire name: 'owner'
	Owner string ``
	// The name of the schema where the registered model resides
	// Wire name: 'schema_name'
	SchemaName string ``
	// The storage location on the cloud under which model version data files
	// are stored
	// Wire name: 'storage_location'
	StorageLocation string ``
	// Last-update timestamp of the registered model in milliseconds since the
	// Unix epoch
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// The identifier of the user who updated the registered model last time
	// Wire name: 'updated_by'
	UpdatedBy       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *RegisteredModelInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RegisteredModelInfoToPb(st *RegisteredModelInfo) (*catalogpb.RegisteredModelInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.RegisteredModelInfoPb{}

	var aliasesPb []catalogpb.RegisteredModelAliasPb
	for _, item := range st.Aliases {
		itemPb, err := RegisteredModelAliasToPb(&item)
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

func RegisteredModelInfoFromPb(pb *catalogpb.RegisteredModelInfoPb) (*RegisteredModelInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelInfo{}

	var aliasesField []RegisteredModelAlias
	for _, itemPb := range pb.Aliases {
		item, err := RegisteredModelAliasFromPb(&itemPb)
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

type RunRefreshRequest struct {
	// UC table name in format `catalog.schema.table_name`. table_name is case
	// insensitive and spaces are disallowed.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func RunRefreshRequestToPb(st *RunRefreshRequest) (*catalogpb.RunRefreshRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.RunRefreshRequestPb{}
	pb.TableName = st.TableName

	return pb, nil
}

func RunRefreshRequestFromPb(pb *catalogpb.RunRefreshRequestPb) (*RunRefreshRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunRefreshRequest{}
	st.TableName = pb.TableName

	return st, nil
}

// Next ID: 40
type SchemaInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool ``
	// Name of parent catalog.
	// Wire name: 'catalog_name'
	CatalogName string ``
	// The type of the parent catalog.
	// Wire name: 'catalog_type'
	CatalogType CatalogType ``
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// Time at which this schema was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of schema creator.
	// Wire name: 'created_by'
	CreatedBy string ``

	// Wire name: 'effective_predictive_optimization_flag'
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag ``
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'enable_predictive_optimization'
	EnablePredictiveOptimization EnablePredictiveOptimization ``
	// Full name of schema, in form of __catalog_name__.__schema_name__.
	// Wire name: 'full_name'
	FullName string ``
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// Name of schema, relative to parent catalog.
	// Wire name: 'name'
	Name string ``
	// Username of current owner of schema.
	// Wire name: 'owner'
	Owner string ``
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string ``
	// The unique identifier of the schema.
	// Wire name: 'schema_id'
	SchemaId string ``
	// Storage location for managed tables within schema.
	// Wire name: 'storage_location'
	StorageLocation string ``
	// Storage root URL for managed tables within schema.
	// Wire name: 'storage_root'
	StorageRoot string ``
	// Time at which this schema was created, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// Username of user who last modified schema.
	// Wire name: 'updated_by'
	UpdatedBy       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *SchemaInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SchemaInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SchemaInfoToPb(st *SchemaInfo) (*catalogpb.SchemaInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.SchemaInfoPb{}
	pb.BrowseOnly = st.BrowseOnly
	pb.CatalogName = st.CatalogName
	catalogTypePb, err := CatalogTypeToPb(&st.CatalogType)
	if err != nil {
		return nil, err
	}
	if catalogTypePb != nil {
		pb.CatalogType = *catalogTypePb
	}
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	effectivePredictiveOptimizationFlagPb, err := EffectivePredictiveOptimizationFlagToPb(st.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagPb != nil {
		pb.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagPb
	}
	enablePredictiveOptimizationPb, err := EnablePredictiveOptimizationToPb(&st.EnablePredictiveOptimization)
	if err != nil {
		return nil, err
	}
	if enablePredictiveOptimizationPb != nil {
		pb.EnablePredictiveOptimization = *enablePredictiveOptimizationPb
	}
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

func SchemaInfoFromPb(pb *catalogpb.SchemaInfoPb) (*SchemaInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SchemaInfo{}
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogName = pb.CatalogName
	catalogTypeField, err := CatalogTypeFromPb(&pb.CatalogType)
	if err != nil {
		return nil, err
	}
	if catalogTypeField != nil {
		st.CatalogType = *catalogTypeField
	}
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	effectivePredictiveOptimizationFlagField, err := EffectivePredictiveOptimizationFlagFromPb(pb.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagField != nil {
		st.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagField
	}
	enablePredictiveOptimizationField, err := EnablePredictiveOptimizationFromPb(&pb.EnablePredictiveOptimization)
	if err != nil {
		return nil, err
	}
	if enablePredictiveOptimizationField != nil {
		st.EnablePredictiveOptimization = *enablePredictiveOptimizationField
	}
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

// Latest kind: CONNECTION_SQLSERVER_OAUTH_M2M = 254; Next id:255
type SecurableKind string

const SecurableKindTableDbStorage SecurableKind = `TABLE_DB_STORAGE`

const SecurableKindTableDelta SecurableKind = `TABLE_DELTA`

const SecurableKindTableDeltasharing SecurableKind = `TABLE_DELTASHARING`

const SecurableKindTableDeltasharingMutable SecurableKind = `TABLE_DELTASHARING_MUTABLE`

const SecurableKindTableDeltaExternal SecurableKind = `TABLE_DELTA_EXTERNAL`

const SecurableKindTableDeltaIcebergDeltasharing SecurableKind = `TABLE_DELTA_ICEBERG_DELTASHARING`

const SecurableKindTableDeltaIcebergManaged SecurableKind = `TABLE_DELTA_ICEBERG_MANAGED`

const SecurableKindTableDeltaUniformHudiExternal SecurableKind = `TABLE_DELTA_UNIFORM_HUDI_EXTERNAL`

const SecurableKindTableDeltaUniformIcebergExternal SecurableKind = `TABLE_DELTA_UNIFORM_ICEBERG_EXTERNAL`

const SecurableKindTableDeltaUniformIcebergForeignHiveMetastoreExternal SecurableKind = `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_HIVE_METASTORE_EXTERNAL`

const SecurableKindTableDeltaUniformIcebergForeignHiveMetastoreManaged SecurableKind = `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_HIVE_METASTORE_MANAGED`

const SecurableKindTableDeltaUniformIcebergForeignSnowflake SecurableKind = `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_SNOWFLAKE`

const SecurableKindTableExternal SecurableKind = `TABLE_EXTERNAL`

const SecurableKindTableFeatureStore SecurableKind = `TABLE_FEATURE_STORE`

const SecurableKindTableFeatureStoreExternal SecurableKind = `TABLE_FEATURE_STORE_EXTERNAL`

const SecurableKindTableForeignBigquery SecurableKind = `TABLE_FOREIGN_BIGQUERY`

const SecurableKindTableForeignDatabricks SecurableKind = `TABLE_FOREIGN_DATABRICKS`

const SecurableKindTableForeignDeltasharing SecurableKind = `TABLE_FOREIGN_DELTASHARING`

const SecurableKindTableForeignHiveMetastore SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE`

const SecurableKindTableForeignHiveMetastoreDbfsExternal SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_DBFS_EXTERNAL`

const SecurableKindTableForeignHiveMetastoreDbfsManaged SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_DBFS_MANAGED`

const SecurableKindTableForeignHiveMetastoreDbfsShallowCloneExternal SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_DBFS_SHALLOW_CLONE_EXTERNAL`

const SecurableKindTableForeignHiveMetastoreDbfsShallowCloneManaged SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_DBFS_SHALLOW_CLONE_MANAGED`

const SecurableKindTableForeignHiveMetastoreDbfsView SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_DBFS_VIEW`

const SecurableKindTableForeignHiveMetastoreExternal SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_EXTERNAL`

const SecurableKindTableForeignHiveMetastoreManaged SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_MANAGED`

const SecurableKindTableForeignHiveMetastoreShallowCloneExternal SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_SHALLOW_CLONE_EXTERNAL`

const SecurableKindTableForeignHiveMetastoreShallowCloneManaged SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_SHALLOW_CLONE_MANAGED`

const SecurableKindTableForeignHiveMetastoreView SecurableKind = `TABLE_FOREIGN_HIVE_METASTORE_VIEW`

const SecurableKindTableForeignMongodb SecurableKind = `TABLE_FOREIGN_MONGODB`

const SecurableKindTableForeignMysql SecurableKind = `TABLE_FOREIGN_MYSQL`

const SecurableKindTableForeignNetsuite SecurableKind = `TABLE_FOREIGN_NETSUITE`

const SecurableKindTableForeignOracle SecurableKind = `TABLE_FOREIGN_ORACLE`

const SecurableKindTableForeignPostgresql SecurableKind = `TABLE_FOREIGN_POSTGRESQL`

const SecurableKindTableForeignRedshift SecurableKind = `TABLE_FOREIGN_REDSHIFT`

const SecurableKindTableForeignSalesforce SecurableKind = `TABLE_FOREIGN_SALESFORCE`

const SecurableKindTableForeignSalesforceDataCloud SecurableKind = `TABLE_FOREIGN_SALESFORCE_DATA_CLOUD`

const SecurableKindTableForeignSalesforceDataCloudFileSharing SecurableKind = `TABLE_FOREIGN_SALESFORCE_DATA_CLOUD_FILE_SHARING`

const SecurableKindTableForeignSalesforceDataCloudFileSharingView SecurableKind = `TABLE_FOREIGN_SALESFORCE_DATA_CLOUD_FILE_SHARING_VIEW`

const SecurableKindTableForeignSnowflake SecurableKind = `TABLE_FOREIGN_SNOWFLAKE`

const SecurableKindTableForeignSqldw SecurableKind = `TABLE_FOREIGN_SQLDW`

const SecurableKindTableForeignSqlserver SecurableKind = `TABLE_FOREIGN_SQLSERVER`

const SecurableKindTableForeignTeradata SecurableKind = `TABLE_FOREIGN_TERADATA`

const SecurableKindTableForeignWorkdayRaas SecurableKind = `TABLE_FOREIGN_WORKDAY_RAAS`

const SecurableKindTableIcebergUniformManaged SecurableKind = `TABLE_ICEBERG_UNIFORM_MANAGED`

const SecurableKindTableInternal SecurableKind = `TABLE_INTERNAL`

const SecurableKindTableManagedPostgresql SecurableKind = `TABLE_MANAGED_POSTGRESQL`

const SecurableKindTableMaterializedView SecurableKind = `TABLE_MATERIALIZED_VIEW`

const SecurableKindTableMaterializedViewDeltasharing SecurableKind = `TABLE_MATERIALIZED_VIEW_DELTASHARING`

const SecurableKindTableMetricView SecurableKind = `TABLE_METRIC_VIEW`

const SecurableKindTableOnlineVectorIndexDirect SecurableKind = `TABLE_ONLINE_VECTOR_INDEX_DIRECT`

const SecurableKindTableOnlineVectorIndexReplica SecurableKind = `TABLE_ONLINE_VECTOR_INDEX_REPLICA`

const SecurableKindTableOnlineView SecurableKind = `TABLE_ONLINE_VIEW`

const SecurableKindTableStandard SecurableKind = `TABLE_STANDARD`

const SecurableKindTableStreamingLiveTable SecurableKind = `TABLE_STREAMING_LIVE_TABLE`

const SecurableKindTableStreamingLiveTableDeltasharing SecurableKind = `TABLE_STREAMING_LIVE_TABLE_DELTASHARING`

const SecurableKindTableSystem SecurableKind = `TABLE_SYSTEM`

const SecurableKindTableSystemDeltasharing SecurableKind = `TABLE_SYSTEM_DELTASHARING`

const SecurableKindTableView SecurableKind = `TABLE_VIEW`

const SecurableKindTableViewDeltasharing SecurableKind = `TABLE_VIEW_DELTASHARING`

// String representation for [fmt.Print]
func (f *SecurableKind) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SecurableKind) Set(v string) error {
	switch v {
	case `TABLE_DB_STORAGE`, `TABLE_DELTA`, `TABLE_DELTASHARING`, `TABLE_DELTASHARING_MUTABLE`, `TABLE_DELTA_EXTERNAL`, `TABLE_DELTA_ICEBERG_DELTASHARING`, `TABLE_DELTA_ICEBERG_MANAGED`, `TABLE_DELTA_UNIFORM_HUDI_EXTERNAL`, `TABLE_DELTA_UNIFORM_ICEBERG_EXTERNAL`, `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_HIVE_METASTORE_EXTERNAL`, `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_HIVE_METASTORE_MANAGED`, `TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_SNOWFLAKE`, `TABLE_EXTERNAL`, `TABLE_FEATURE_STORE`, `TABLE_FEATURE_STORE_EXTERNAL`, `TABLE_FOREIGN_BIGQUERY`, `TABLE_FOREIGN_DATABRICKS`, `TABLE_FOREIGN_DELTASHARING`, `TABLE_FOREIGN_HIVE_METASTORE`, `TABLE_FOREIGN_HIVE_METASTORE_DBFS_EXTERNAL`, `TABLE_FOREIGN_HIVE_METASTORE_DBFS_MANAGED`, `TABLE_FOREIGN_HIVE_METASTORE_DBFS_SHALLOW_CLONE_EXTERNAL`, `TABLE_FOREIGN_HIVE_METASTORE_DBFS_SHALLOW_CLONE_MANAGED`, `TABLE_FOREIGN_HIVE_METASTORE_DBFS_VIEW`, `TABLE_FOREIGN_HIVE_METASTORE_EXTERNAL`, `TABLE_FOREIGN_HIVE_METASTORE_MANAGED`, `TABLE_FOREIGN_HIVE_METASTORE_SHALLOW_CLONE_EXTERNAL`, `TABLE_FOREIGN_HIVE_METASTORE_SHALLOW_CLONE_MANAGED`, `TABLE_FOREIGN_HIVE_METASTORE_VIEW`, `TABLE_FOREIGN_MONGODB`, `TABLE_FOREIGN_MYSQL`, `TABLE_FOREIGN_NETSUITE`, `TABLE_FOREIGN_ORACLE`, `TABLE_FOREIGN_POSTGRESQL`, `TABLE_FOREIGN_REDSHIFT`, `TABLE_FOREIGN_SALESFORCE`, `TABLE_FOREIGN_SALESFORCE_DATA_CLOUD`, `TABLE_FOREIGN_SALESFORCE_DATA_CLOUD_FILE_SHARING`, `TABLE_FOREIGN_SALESFORCE_DATA_CLOUD_FILE_SHARING_VIEW`, `TABLE_FOREIGN_SNOWFLAKE`, `TABLE_FOREIGN_SQLDW`, `TABLE_FOREIGN_SQLSERVER`, `TABLE_FOREIGN_TERADATA`, `TABLE_FOREIGN_WORKDAY_RAAS`, `TABLE_ICEBERG_UNIFORM_MANAGED`, `TABLE_INTERNAL`, `TABLE_MANAGED_POSTGRESQL`, `TABLE_MATERIALIZED_VIEW`, `TABLE_MATERIALIZED_VIEW_DELTASHARING`, `TABLE_METRIC_VIEW`, `TABLE_ONLINE_VECTOR_INDEX_DIRECT`, `TABLE_ONLINE_VECTOR_INDEX_REPLICA`, `TABLE_ONLINE_VIEW`, `TABLE_STANDARD`, `TABLE_STREAMING_LIVE_TABLE`, `TABLE_STREAMING_LIVE_TABLE_DELTASHARING`, `TABLE_SYSTEM`, `TABLE_SYSTEM_DELTASHARING`, `TABLE_VIEW`, `TABLE_VIEW_DELTASHARING`:
		*f = SecurableKind(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TABLE_DB_STORAGE", "TABLE_DELTA", "TABLE_DELTASHARING", "TABLE_DELTASHARING_MUTABLE", "TABLE_DELTA_EXTERNAL", "TABLE_DELTA_ICEBERG_DELTASHARING", "TABLE_DELTA_ICEBERG_MANAGED", "TABLE_DELTA_UNIFORM_HUDI_EXTERNAL", "TABLE_DELTA_UNIFORM_ICEBERG_EXTERNAL", "TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_HIVE_METASTORE_EXTERNAL", "TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_HIVE_METASTORE_MANAGED", "TABLE_DELTA_UNIFORM_ICEBERG_FOREIGN_SNOWFLAKE", "TABLE_EXTERNAL", "TABLE_FEATURE_STORE", "TABLE_FEATURE_STORE_EXTERNAL", "TABLE_FOREIGN_BIGQUERY", "TABLE_FOREIGN_DATABRICKS", "TABLE_FOREIGN_DELTASHARING", "TABLE_FOREIGN_HIVE_METASTORE", "TABLE_FOREIGN_HIVE_METASTORE_DBFS_EXTERNAL", "TABLE_FOREIGN_HIVE_METASTORE_DBFS_MANAGED", "TABLE_FOREIGN_HIVE_METASTORE_DBFS_SHALLOW_CLONE_EXTERNAL", "TABLE_FOREIGN_HIVE_METASTORE_DBFS_SHALLOW_CLONE_MANAGED", "TABLE_FOREIGN_HIVE_METASTORE_DBFS_VIEW", "TABLE_FOREIGN_HIVE_METASTORE_EXTERNAL", "TABLE_FOREIGN_HIVE_METASTORE_MANAGED", "TABLE_FOREIGN_HIVE_METASTORE_SHALLOW_CLONE_EXTERNAL", "TABLE_FOREIGN_HIVE_METASTORE_SHALLOW_CLONE_MANAGED", "TABLE_FOREIGN_HIVE_METASTORE_VIEW", "TABLE_FOREIGN_MONGODB", "TABLE_FOREIGN_MYSQL", "TABLE_FOREIGN_NETSUITE", "TABLE_FOREIGN_ORACLE", "TABLE_FOREIGN_POSTGRESQL", "TABLE_FOREIGN_REDSHIFT", "TABLE_FOREIGN_SALESFORCE", "TABLE_FOREIGN_SALESFORCE_DATA_CLOUD", "TABLE_FOREIGN_SALESFORCE_DATA_CLOUD_FILE_SHARING", "TABLE_FOREIGN_SALESFORCE_DATA_CLOUD_FILE_SHARING_VIEW", "TABLE_FOREIGN_SNOWFLAKE", "TABLE_FOREIGN_SQLDW", "TABLE_FOREIGN_SQLSERVER", "TABLE_FOREIGN_TERADATA", "TABLE_FOREIGN_WORKDAY_RAAS", "TABLE_ICEBERG_UNIFORM_MANAGED", "TABLE_INTERNAL", "TABLE_MANAGED_POSTGRESQL", "TABLE_MATERIALIZED_VIEW", "TABLE_MATERIALIZED_VIEW_DELTASHARING", "TABLE_METRIC_VIEW", "TABLE_ONLINE_VECTOR_INDEX_DIRECT", "TABLE_ONLINE_VECTOR_INDEX_REPLICA", "TABLE_ONLINE_VIEW", "TABLE_STANDARD", "TABLE_STREAMING_LIVE_TABLE", "TABLE_STREAMING_LIVE_TABLE_DELTASHARING", "TABLE_SYSTEM", "TABLE_SYSTEM_DELTASHARING", "TABLE_VIEW", "TABLE_VIEW_DELTASHARING"`, v)
	}
}

// Values returns all possible values for SecurableKind.
//
// There is no guarantee on the order of the values in the slice.
func (f *SecurableKind) Values() []SecurableKind {
	return []SecurableKind{
		SecurableKindTableDbStorage,
		SecurableKindTableDelta,
		SecurableKindTableDeltasharing,
		SecurableKindTableDeltasharingMutable,
		SecurableKindTableDeltaExternal,
		SecurableKindTableDeltaIcebergDeltasharing,
		SecurableKindTableDeltaIcebergManaged,
		SecurableKindTableDeltaUniformHudiExternal,
		SecurableKindTableDeltaUniformIcebergExternal,
		SecurableKindTableDeltaUniformIcebergForeignHiveMetastoreExternal,
		SecurableKindTableDeltaUniformIcebergForeignHiveMetastoreManaged,
		SecurableKindTableDeltaUniformIcebergForeignSnowflake,
		SecurableKindTableExternal,
		SecurableKindTableFeatureStore,
		SecurableKindTableFeatureStoreExternal,
		SecurableKindTableForeignBigquery,
		SecurableKindTableForeignDatabricks,
		SecurableKindTableForeignDeltasharing,
		SecurableKindTableForeignHiveMetastore,
		SecurableKindTableForeignHiveMetastoreDbfsExternal,
		SecurableKindTableForeignHiveMetastoreDbfsManaged,
		SecurableKindTableForeignHiveMetastoreDbfsShallowCloneExternal,
		SecurableKindTableForeignHiveMetastoreDbfsShallowCloneManaged,
		SecurableKindTableForeignHiveMetastoreDbfsView,
		SecurableKindTableForeignHiveMetastoreExternal,
		SecurableKindTableForeignHiveMetastoreManaged,
		SecurableKindTableForeignHiveMetastoreShallowCloneExternal,
		SecurableKindTableForeignHiveMetastoreShallowCloneManaged,
		SecurableKindTableForeignHiveMetastoreView,
		SecurableKindTableForeignMongodb,
		SecurableKindTableForeignMysql,
		SecurableKindTableForeignNetsuite,
		SecurableKindTableForeignOracle,
		SecurableKindTableForeignPostgresql,
		SecurableKindTableForeignRedshift,
		SecurableKindTableForeignSalesforce,
		SecurableKindTableForeignSalesforceDataCloud,
		SecurableKindTableForeignSalesforceDataCloudFileSharing,
		SecurableKindTableForeignSalesforceDataCloudFileSharingView,
		SecurableKindTableForeignSnowflake,
		SecurableKindTableForeignSqldw,
		SecurableKindTableForeignSqlserver,
		SecurableKindTableForeignTeradata,
		SecurableKindTableForeignWorkdayRaas,
		SecurableKindTableIcebergUniformManaged,
		SecurableKindTableInternal,
		SecurableKindTableManagedPostgresql,
		SecurableKindTableMaterializedView,
		SecurableKindTableMaterializedViewDeltasharing,
		SecurableKindTableMetricView,
		SecurableKindTableOnlineVectorIndexDirect,
		SecurableKindTableOnlineVectorIndexReplica,
		SecurableKindTableOnlineView,
		SecurableKindTableStandard,
		SecurableKindTableStreamingLiveTable,
		SecurableKindTableStreamingLiveTableDeltasharing,
		SecurableKindTableSystem,
		SecurableKindTableSystemDeltasharing,
		SecurableKindTableView,
		SecurableKindTableViewDeltasharing,
	}
}

// Type always returns SecurableKind to satisfy [pflag.Value] interface
func (f *SecurableKind) Type() string {
	return "SecurableKind"
}

func SecurableKindToPb(st *SecurableKind) (*catalogpb.SecurableKindPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.SecurableKindPb(*st)
	return &pb, nil
}

func SecurableKindFromPb(pb *catalogpb.SecurableKindPb) (*SecurableKind, error) {
	if pb == nil {
		return nil, nil
	}
	st := SecurableKind(*pb)
	return &st, nil
}

// Manifest of a specific securable kind.
type SecurableKindManifest struct {
	// Privileges that can be assigned to the securable.
	// Wire name: 'assignable_privileges'
	AssignablePrivileges []string ``
	// A list of capabilities in the securable kind.
	// Wire name: 'capabilities'
	Capabilities []string ``
	// Detailed specs of allowed options.
	// Wire name: 'options'
	Options []OptionSpec ``
	// Securable kind to get manifest of.
	// Wire name: 'securable_kind'
	SecurableKind SecurableKind ``
	// Securable Type of the kind.
	// Wire name: 'securable_type'
	SecurableType SecurableType ``
}

func SecurableKindManifestToPb(st *SecurableKindManifest) (*catalogpb.SecurableKindManifestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.SecurableKindManifestPb{}
	pb.AssignablePrivileges = st.AssignablePrivileges
	pb.Capabilities = st.Capabilities

	var optionsPb []catalogpb.OptionSpecPb
	for _, item := range st.Options {
		itemPb, err := OptionSpecToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			optionsPb = append(optionsPb, *itemPb)
		}
	}
	pb.Options = optionsPb
	securableKindPb, err := SecurableKindToPb(&st.SecurableKind)
	if err != nil {
		return nil, err
	}
	if securableKindPb != nil {
		pb.SecurableKind = *securableKindPb
	}
	securableTypePb, err := SecurableTypeToPb(&st.SecurableType)
	if err != nil {
		return nil, err
	}
	if securableTypePb != nil {
		pb.SecurableType = *securableTypePb
	}

	return pb, nil
}

func SecurableKindManifestFromPb(pb *catalogpb.SecurableKindManifestPb) (*SecurableKindManifest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SecurableKindManifest{}
	st.AssignablePrivileges = pb.AssignablePrivileges
	st.Capabilities = pb.Capabilities

	var optionsField []OptionSpec
	for _, itemPb := range pb.Options {
		item, err := OptionSpecFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			optionsField = append(optionsField, *item)
		}
	}
	st.Options = optionsField
	securableKindField, err := SecurableKindFromPb(&pb.SecurableKind)
	if err != nil {
		return nil, err
	}
	if securableKindField != nil {
		st.SecurableKind = *securableKindField
	}
	securableTypeField, err := SecurableTypeFromPb(&pb.SecurableType)
	if err != nil {
		return nil, err
	}
	if securableTypeField != nil {
		st.SecurableType = *securableTypeField
	}

	return st, nil
}

// The type of Unity Catalog securable.
type SecurableType string

const SecurableTypeCatalog SecurableType = `CATALOG`

const SecurableTypeCleanRoom SecurableType = `CLEAN_ROOM`

const SecurableTypeConnection SecurableType = `CONNECTION`

const SecurableTypeCredential SecurableType = `CREDENTIAL`

const SecurableTypeExternalLocation SecurableType = `EXTERNAL_LOCATION`

const SecurableTypeExternalMetadata SecurableType = `EXTERNAL_METADATA`

const SecurableTypeFunction SecurableType = `FUNCTION`

const SecurableTypeMetastore SecurableType = `METASTORE`

const SecurableTypePipeline SecurableType = `PIPELINE`

const SecurableTypeProvider SecurableType = `PROVIDER`

const SecurableTypeRecipient SecurableType = `RECIPIENT`

const SecurableTypeSchema SecurableType = `SCHEMA`

const SecurableTypeShare SecurableType = `SHARE`

const SecurableTypeStagingTable SecurableType = `STAGING_TABLE`

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
	case `CATALOG`, `CLEAN_ROOM`, `CONNECTION`, `CREDENTIAL`, `EXTERNAL_LOCATION`, `EXTERNAL_METADATA`, `FUNCTION`, `METASTORE`, `PIPELINE`, `PROVIDER`, `RECIPIENT`, `SCHEMA`, `SHARE`, `STAGING_TABLE`, `STORAGE_CREDENTIAL`, `TABLE`, `VOLUME`:
		*f = SecurableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CATALOG", "CLEAN_ROOM", "CONNECTION", "CREDENTIAL", "EXTERNAL_LOCATION", "EXTERNAL_METADATA", "FUNCTION", "METASTORE", "PIPELINE", "PROVIDER", "RECIPIENT", "SCHEMA", "SHARE", "STAGING_TABLE", "STORAGE_CREDENTIAL", "TABLE", "VOLUME"`, v)
	}
}

// Values returns all possible values for SecurableType.
//
// There is no guarantee on the order of the values in the slice.
func (f *SecurableType) Values() []SecurableType {
	return []SecurableType{
		SecurableTypeCatalog,
		SecurableTypeCleanRoom,
		SecurableTypeConnection,
		SecurableTypeCredential,
		SecurableTypeExternalLocation,
		SecurableTypeExternalMetadata,
		SecurableTypeFunction,
		SecurableTypeMetastore,
		SecurableTypePipeline,
		SecurableTypeProvider,
		SecurableTypeRecipient,
		SecurableTypeSchema,
		SecurableTypeShare,
		SecurableTypeStagingTable,
		SecurableTypeStorageCredential,
		SecurableTypeTable,
		SecurableTypeVolume,
	}
}

// Type always returns SecurableType to satisfy [pflag.Value] interface
func (f *SecurableType) Type() string {
	return "SecurableType"
}

func SecurableTypeToPb(st *SecurableType) (*catalogpb.SecurableTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.SecurableTypePb(*st)
	return &pb, nil
}

func SecurableTypeFromPb(pb *catalogpb.SecurableTypePb) (*SecurableType, error) {
	if pb == nil {
		return nil, nil
	}
	st := SecurableType(*pb)
	return &st, nil
}

type SetArtifactAllowlist struct {
	// A list of allowed artifact match patterns.
	// Wire name: 'artifact_matchers'
	ArtifactMatchers []ArtifactMatcher ``
	// The artifact type of the allowlist.
	// Wire name: 'artifact_type'
	ArtifactType ArtifactType `tf:"-"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of the user who set the artifact allowlist.
	// Wire name: 'created_by'
	CreatedBy string ``
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *SetArtifactAllowlist) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SetArtifactAllowlist) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SetArtifactAllowlistToPb(st *SetArtifactAllowlist) (*catalogpb.SetArtifactAllowlistPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.SetArtifactAllowlistPb{}

	var artifactMatchersPb []catalogpb.ArtifactMatcherPb
	for _, item := range st.ArtifactMatchers {
		itemPb, err := ArtifactMatcherToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			artifactMatchersPb = append(artifactMatchersPb, *itemPb)
		}
	}
	pb.ArtifactMatchers = artifactMatchersPb
	artifactTypePb, err := ArtifactTypeToPb(&st.ArtifactType)
	if err != nil {
		return nil, err
	}
	if artifactTypePb != nil {
		pb.ArtifactType = *artifactTypePb
	}
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.MetastoreId = st.MetastoreId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SetArtifactAllowlistFromPb(pb *catalogpb.SetArtifactAllowlistPb) (*SetArtifactAllowlist, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetArtifactAllowlist{}

	var artifactMatchersField []ArtifactMatcher
	for _, itemPb := range pb.ArtifactMatchers {
		item, err := ArtifactMatcherFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			artifactMatchersField = append(artifactMatchersField, *item)
		}
	}
	st.ArtifactMatchers = artifactMatchersField
	artifactTypeField, err := ArtifactTypeFromPb(&pb.ArtifactType)
	if err != nil {
		return nil, err
	}
	if artifactTypeField != nil {
		st.ArtifactType = *artifactTypeField
	}
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.MetastoreId = pb.MetastoreId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type SetRegisteredModelAliasRequest struct {
	// The name of the alias
	// Wire name: 'alias'
	Alias string ``
	// Full name of the registered model
	// Wire name: 'full_name'
	FullName string ``
	// The version number of the model version to which the alias points
	// Wire name: 'version_num'
	VersionNum int ``
}

func SetRegisteredModelAliasRequestToPb(st *SetRegisteredModelAliasRequest) (*catalogpb.SetRegisteredModelAliasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.SetRegisteredModelAliasRequestPb{}
	pb.Alias = st.Alias
	pb.FullName = st.FullName
	pb.VersionNum = st.VersionNum

	return pb, nil
}

func SetRegisteredModelAliasRequestFromPb(pb *catalogpb.SetRegisteredModelAliasRequestPb) (*SetRegisteredModelAliasRequest, error) {
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
	// Sets the value of the 'x-amz-server-side-encryption' header in S3
	// request.
	// Wire name: 'algorithm'
	Algorithm SseEncryptionDetailsAlgorithm ``
	// Optional. The ARN of the SSE-KMS key used with the S3 location, when
	// algorithm = "SSE-KMS". Sets the value of the
	// 'x-amz-server-side-encryption-aws-kms-key-id' header.
	// Wire name: 'aws_kms_key_arn'
	AwsKmsKeyArn    string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *SseEncryptionDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SseEncryptionDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SseEncryptionDetailsToPb(st *SseEncryptionDetails) (*catalogpb.SseEncryptionDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.SseEncryptionDetailsPb{}
	algorithmPb, err := SseEncryptionDetailsAlgorithmToPb(&st.Algorithm)
	if err != nil {
		return nil, err
	}
	if algorithmPb != nil {
		pb.Algorithm = *algorithmPb
	}
	pb.AwsKmsKeyArn = st.AwsKmsKeyArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SseEncryptionDetailsFromPb(pb *catalogpb.SseEncryptionDetailsPb) (*SseEncryptionDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SseEncryptionDetails{}
	algorithmField, err := SseEncryptionDetailsAlgorithmFromPb(&pb.Algorithm)
	if err != nil {
		return nil, err
	}
	if algorithmField != nil {
		st.Algorithm = *algorithmField
	}
	st.AwsKmsKeyArn = pb.AwsKmsKeyArn

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type SseEncryptionDetailsAlgorithm string

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

// Values returns all possible values for SseEncryptionDetailsAlgorithm.
//
// There is no guarantee on the order of the values in the slice.
func (f *SseEncryptionDetailsAlgorithm) Values() []SseEncryptionDetailsAlgorithm {
	return []SseEncryptionDetailsAlgorithm{
		SseEncryptionDetailsAlgorithmAwsSseKms,
		SseEncryptionDetailsAlgorithmAwsSseS3,
	}
}

// Type always returns SseEncryptionDetailsAlgorithm to satisfy [pflag.Value] interface
func (f *SseEncryptionDetailsAlgorithm) Type() string {
	return "SseEncryptionDetailsAlgorithm"
}

func SseEncryptionDetailsAlgorithmToPb(st *SseEncryptionDetailsAlgorithm) (*catalogpb.SseEncryptionDetailsAlgorithmPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.SseEncryptionDetailsAlgorithmPb(*st)
	return &pb, nil
}

func SseEncryptionDetailsAlgorithmFromPb(pb *catalogpb.SseEncryptionDetailsAlgorithmPb) (*SseEncryptionDetailsAlgorithm, error) {
	if pb == nil {
		return nil, nil
	}
	st := SseEncryptionDetailsAlgorithm(*pb)
	return &st, nil
}

type StorageCredentialInfo struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRoleResponse ``
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentityResponse ``
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal ``
	// The Cloudflare API token configuration.
	// Wire name: 'cloudflare_api_token'
	CloudflareApiToken *CloudflareApiToken ``
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string ``
	// Time at which this credential was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of credential creator.
	// Wire name: 'created_by'
	CreatedBy string ``
	// The Databricks managed GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountResponse ``
	// The full name of the credential.
	// Wire name: 'full_name'
	FullName string ``
	// The unique identifier of the credential.
	// Wire name: 'id'
	Id string ``
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode ``
	// Unique identifier of the parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	// Wire name: 'name'
	Name string ``
	// Username of current owner of credential.
	// Wire name: 'owner'
	Owner string ``
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	// Wire name: 'read_only'
	ReadOnly bool ``
	// Time at which this credential was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// Username of user who last modified the credential.
	// Wire name: 'updated_by'
	UpdatedBy string ``
	// Whether this credential is the current metastore's root storage
	// credential. Only applicable when purpose is **STORAGE**.
	// Wire name: 'used_for_managed_storage'
	UsedForManagedStorage bool     ``
	ForceSendFields       []string `tf:"-"`
}

func (s *StorageCredentialInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StorageCredentialInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func StorageCredentialInfoToPb(st *StorageCredentialInfo) (*catalogpb.StorageCredentialInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.StorageCredentialInfoPb{}
	awsIamRolePb, err := AwsIamRoleResponseToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}
	azureManagedIdentityPb, err := AzureManagedIdentityResponseToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}
	azureServicePrincipalPb, err := AzureServicePrincipalToPb(st.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalPb != nil {
		pb.AzureServicePrincipal = azureServicePrincipalPb
	}
	cloudflareApiTokenPb, err := CloudflareApiTokenToPb(st.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenPb != nil {
		pb.CloudflareApiToken = cloudflareApiTokenPb
	}
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	databricksGcpServiceAccountPb, err := DatabricksGcpServiceAccountResponseToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}
	pb.FullName = st.FullName
	pb.Id = st.Id
	isolationModePb, err := IsolationModeToPb(&st.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}
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

func StorageCredentialInfoFromPb(pb *catalogpb.StorageCredentialInfoPb) (*StorageCredentialInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StorageCredentialInfo{}
	awsIamRoleField, err := AwsIamRoleResponseFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := AzureManagedIdentityResponseFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	azureServicePrincipalField, err := AzureServicePrincipalFromPb(pb.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalField != nil {
		st.AzureServicePrincipal = azureServicePrincipalField
	}
	cloudflareApiTokenField, err := CloudflareApiTokenFromPb(pb.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenField != nil {
		st.CloudflareApiToken = cloudflareApiTokenField
	}
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	databricksGcpServiceAccountField, err := DatabricksGcpServiceAccountResponseFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	st.FullName = pb.FullName
	st.Id = pb.Id
	isolationModeField, err := IsolationModeFromPb(&pb.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
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

type SystemSchemaInfo struct {
	// Name of the system schema.
	// Wire name: 'schema'
	Schema string ``
	// The current state of enablement for the system schema. An empty string
	// means the system schema is available and ready for opt-in. Possible
	// values: AVAILABLE | ENABLE_INITIALIZED | ENABLE_COMPLETED |
	// DISABLE_INITIALIZED | UNAVAILABLE
	// Wire name: 'state'
	State string ``
}

func SystemSchemaInfoToPb(st *SystemSchemaInfo) (*catalogpb.SystemSchemaInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.SystemSchemaInfoPb{}
	pb.Schema = st.Schema
	pb.State = st.State

	return pb, nil
}

func SystemSchemaInfoFromPb(pb *catalogpb.SystemSchemaInfoPb) (*SystemSchemaInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SystemSchemaInfo{}
	st.Schema = pb.Schema
	st.State = pb.State

	return st, nil
}

type SystemType string

const SystemTypeAmazonRedshift SystemType = `AMAZON_REDSHIFT`

const SystemTypeAzureSynapse SystemType = `AZURE_SYNAPSE`

const SystemTypeConfluent SystemType = `CONFLUENT`

const SystemTypeDatabricks SystemType = `DATABRICKS`

const SystemTypeGoogleBigquery SystemType = `GOOGLE_BIGQUERY`

const SystemTypeKafka SystemType = `KAFKA`

const SystemTypeLooker SystemType = `LOOKER`

const SystemTypeMicrosoftFabric SystemType = `MICROSOFT_FABRIC`

const SystemTypeMicrosoftSqlServer SystemType = `MICROSOFT_SQL_SERVER`

const SystemTypeMongodb SystemType = `MONGODB`

const SystemTypeMysql SystemType = `MYSQL`

const SystemTypeOracle SystemType = `ORACLE`

const SystemTypeOther SystemType = `OTHER`

const SystemTypePostgresql SystemType = `POSTGRESQL`

const SystemTypePowerBi SystemType = `POWER_BI`

const SystemTypeSalesforce SystemType = `SALESFORCE`

const SystemTypeSap SystemType = `SAP`

const SystemTypeServicenow SystemType = `SERVICENOW`

const SystemTypeSnowflake SystemType = `SNOWFLAKE`

const SystemTypeTableau SystemType = `TABLEAU`

const SystemTypeTeradata SystemType = `TERADATA`

const SystemTypeWorkday SystemType = `WORKDAY`

// String representation for [fmt.Print]
func (f *SystemType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SystemType) Set(v string) error {
	switch v {
	case `AMAZON_REDSHIFT`, `AZURE_SYNAPSE`, `CONFLUENT`, `DATABRICKS`, `GOOGLE_BIGQUERY`, `KAFKA`, `LOOKER`, `MICROSOFT_FABRIC`, `MICROSOFT_SQL_SERVER`, `MONGODB`, `MYSQL`, `ORACLE`, `OTHER`, `POSTGRESQL`, `POWER_BI`, `SALESFORCE`, `SAP`, `SERVICENOW`, `SNOWFLAKE`, `TABLEAU`, `TERADATA`, `WORKDAY`:
		*f = SystemType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AMAZON_REDSHIFT", "AZURE_SYNAPSE", "CONFLUENT", "DATABRICKS", "GOOGLE_BIGQUERY", "KAFKA", "LOOKER", "MICROSOFT_FABRIC", "MICROSOFT_SQL_SERVER", "MONGODB", "MYSQL", "ORACLE", "OTHER", "POSTGRESQL", "POWER_BI", "SALESFORCE", "SAP", "SERVICENOW", "SNOWFLAKE", "TABLEAU", "TERADATA", "WORKDAY"`, v)
	}
}

// Values returns all possible values for SystemType.
//
// There is no guarantee on the order of the values in the slice.
func (f *SystemType) Values() []SystemType {
	return []SystemType{
		SystemTypeAmazonRedshift,
		SystemTypeAzureSynapse,
		SystemTypeConfluent,
		SystemTypeDatabricks,
		SystemTypeGoogleBigquery,
		SystemTypeKafka,
		SystemTypeLooker,
		SystemTypeMicrosoftFabric,
		SystemTypeMicrosoftSqlServer,
		SystemTypeMongodb,
		SystemTypeMysql,
		SystemTypeOracle,
		SystemTypeOther,
		SystemTypePostgresql,
		SystemTypePowerBi,
		SystemTypeSalesforce,
		SystemTypeSap,
		SystemTypeServicenow,
		SystemTypeSnowflake,
		SystemTypeTableau,
		SystemTypeTeradata,
		SystemTypeWorkday,
	}
}

// Type always returns SystemType to satisfy [pflag.Value] interface
func (f *SystemType) Type() string {
	return "SystemType"
}

func SystemTypeToPb(st *SystemType) (*catalogpb.SystemTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.SystemTypePb(*st)
	return &pb, nil
}

func SystemTypeFromPb(pb *catalogpb.SystemTypePb) (*SystemType, error) {
	if pb == nil {
		return nil, nil
	}
	st := SystemType(*pb)
	return &st, nil
}

// A table constraint, as defined by *one* of the following fields being set:
// __primary_key_constraint__, __foreign_key_constraint__,
// __named_table_constraint__.
type TableConstraint struct {

	// Wire name: 'foreign_key_constraint'
	ForeignKeyConstraint *ForeignKeyConstraint ``

	// Wire name: 'named_table_constraint'
	NamedTableConstraint *NamedTableConstraint ``

	// Wire name: 'primary_key_constraint'
	PrimaryKeyConstraint *PrimaryKeyConstraint ``
}

func TableConstraintToPb(st *TableConstraint) (*catalogpb.TableConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.TableConstraintPb{}
	foreignKeyConstraintPb, err := ForeignKeyConstraintToPb(st.ForeignKeyConstraint)
	if err != nil {
		return nil, err
	}
	if foreignKeyConstraintPb != nil {
		pb.ForeignKeyConstraint = foreignKeyConstraintPb
	}
	namedTableConstraintPb, err := NamedTableConstraintToPb(st.NamedTableConstraint)
	if err != nil {
		return nil, err
	}
	if namedTableConstraintPb != nil {
		pb.NamedTableConstraint = namedTableConstraintPb
	}
	primaryKeyConstraintPb, err := PrimaryKeyConstraintToPb(st.PrimaryKeyConstraint)
	if err != nil {
		return nil, err
	}
	if primaryKeyConstraintPb != nil {
		pb.PrimaryKeyConstraint = primaryKeyConstraintPb
	}

	return pb, nil
}

func TableConstraintFromPb(pb *catalogpb.TableConstraintPb) (*TableConstraint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableConstraint{}
	foreignKeyConstraintField, err := ForeignKeyConstraintFromPb(pb.ForeignKeyConstraint)
	if err != nil {
		return nil, err
	}
	if foreignKeyConstraintField != nil {
		st.ForeignKeyConstraint = foreignKeyConstraintField
	}
	namedTableConstraintField, err := NamedTableConstraintFromPb(pb.NamedTableConstraint)
	if err != nil {
		return nil, err
	}
	if namedTableConstraintField != nil {
		st.NamedTableConstraint = namedTableConstraintField
	}
	primaryKeyConstraintField, err := PrimaryKeyConstraintFromPb(pb.PrimaryKeyConstraint)
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
	TableFullName string ``
}

func TableDependencyToPb(st *TableDependency) (*catalogpb.TableDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.TableDependencyPb{}
	pb.TableFullName = st.TableFullName

	return pb, nil
}

func TableDependencyFromPb(pb *catalogpb.TableDependencyPb) (*TableDependency, error) {
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
	TableExists     bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *TableExistsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableExistsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TableExistsResponseToPb(st *TableExistsResponse) (*catalogpb.TableExistsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.TableExistsResponsePb{}
	pb.TableExists = st.TableExists

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TableExistsResponseFromPb(pb *catalogpb.TableExistsResponsePb) (*TableExistsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableExistsResponse{}
	st.TableExists = pb.TableExists

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type TableInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	// Wire name: 'access_point'
	AccessPoint string ``
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool ``
	// Name of parent catalog.
	// Wire name: 'catalog_name'
	CatalogName string ``
	// The array of __ColumnInfo__ definitions of the table's columns.
	// Wire name: 'columns'
	Columns []ColumnInfo ``
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// Time at which this table was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of table creator.
	// Wire name: 'created_by'
	CreatedBy string ``
	// Unique ID of the Data Access Configuration to use with the table data.
	// Wire name: 'data_access_configuration_id'
	DataAccessConfigurationId string ``

	// Wire name: 'data_source_format'
	DataSourceFormat DataSourceFormat ``
	// Time at which this table was deleted, in epoch milliseconds. Field is
	// omitted if table is not deleted.
	// Wire name: 'deleted_at'
	DeletedAt int64 ``
	// Information pertaining to current state of the delta table.
	// Wire name: 'delta_runtime_properties_kvpairs'
	DeltaRuntimePropertiesKvpairs *DeltaRuntimePropertiesKvPairs ``

	// Wire name: 'effective_predictive_optimization_flag'
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag ``

	// Wire name: 'enable_predictive_optimization'
	EnablePredictiveOptimization EnablePredictiveOptimization ``

	// Wire name: 'encryption_details'
	EncryptionDetails *EncryptionDetails ``
	// Full name of table, in form of
	// __catalog_name__.__schema_name__.__table_name__
	// Wire name: 'full_name'
	FullName string ``
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// Name of table, relative to parent schema.
	// Wire name: 'name'
	Name string ``
	// Username of current owner of table.
	// Wire name: 'owner'
	Owner string ``
	// The pipeline ID of the table. Applicable for tables created by pipelines
	// (Materialized View, Streaming Table, etc.).
	// Wire name: 'pipeline_id'
	PipelineId string ``
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string ``

	// Wire name: 'row_filter'
	RowFilter *TableRowFilter ``
	// Name of parent schema relative to its parent catalog.
	// Wire name: 'schema_name'
	SchemaName string ``
	// SecurableKindManifest of table, including capabilities the table has.
	// Wire name: 'securable_kind_manifest'
	SecurableKindManifest *SecurableKindManifest ``
	// List of schemes whose objects can be referenced without qualification.
	// Wire name: 'sql_path'
	SqlPath string ``
	// Name of the storage credential, when a storage credential is configured
	// for use with this table.
	// Wire name: 'storage_credential_name'
	StorageCredentialName string ``
	// Storage root URL for table (for **MANAGED**, **EXTERNAL** tables).
	// Wire name: 'storage_location'
	StorageLocation string ``
	// List of table constraints. Note: this field is not set in the output of
	// the __listTables__ API.
	// Wire name: 'table_constraints'
	TableConstraints []TableConstraint ``
	// The unique identifier of the table.
	// Wire name: 'table_id'
	TableId string ``

	// Wire name: 'table_type'
	TableType TableType ``
	// Time at which this table was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// Username of user who last modified the table.
	// Wire name: 'updated_by'
	UpdatedBy string ``
	// View definition SQL (when __table_type__ is **VIEW**,
	// **MATERIALIZED_VIEW**, or **STREAMING_TABLE**)
	// Wire name: 'view_definition'
	ViewDefinition string ``
	// View dependencies (when table_type == **VIEW** or **MATERIALIZED_VIEW**,
	// **STREAMING_TABLE**) - when DependencyList is None, the dependency is not
	// provided; - when DependencyList is an empty list, the dependency is
	// provided but is empty; - when DependencyList is not an empty list,
	// dependencies are provided and recorded.
	// Wire name: 'view_dependencies'
	ViewDependencies *DependencyList ``
	ForceSendFields  []string        `tf:"-"`
}

func (s *TableInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TableInfoToPb(st *TableInfo) (*catalogpb.TableInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.TableInfoPb{}
	pb.AccessPoint = st.AccessPoint
	pb.BrowseOnly = st.BrowseOnly
	pb.CatalogName = st.CatalogName

	var columnsPb []catalogpb.ColumnInfoPb
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
	dataSourceFormatPb, err := DataSourceFormatToPb(&st.DataSourceFormat)
	if err != nil {
		return nil, err
	}
	if dataSourceFormatPb != nil {
		pb.DataSourceFormat = *dataSourceFormatPb
	}
	pb.DeletedAt = st.DeletedAt
	deltaRuntimePropertiesKvpairsPb, err := DeltaRuntimePropertiesKvPairsToPb(st.DeltaRuntimePropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if deltaRuntimePropertiesKvpairsPb != nil {
		pb.DeltaRuntimePropertiesKvpairs = deltaRuntimePropertiesKvpairsPb
	}
	effectivePredictiveOptimizationFlagPb, err := EffectivePredictiveOptimizationFlagToPb(st.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagPb != nil {
		pb.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagPb
	}
	enablePredictiveOptimizationPb, err := EnablePredictiveOptimizationToPb(&st.EnablePredictiveOptimization)
	if err != nil {
		return nil, err
	}
	if enablePredictiveOptimizationPb != nil {
		pb.EnablePredictiveOptimization = *enablePredictiveOptimizationPb
	}
	encryptionDetailsPb, err := EncryptionDetailsToPb(st.EncryptionDetails)
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
	rowFilterPb, err := TableRowFilterToPb(st.RowFilter)
	if err != nil {
		return nil, err
	}
	if rowFilterPb != nil {
		pb.RowFilter = rowFilterPb
	}
	pb.SchemaName = st.SchemaName
	securableKindManifestPb, err := SecurableKindManifestToPb(st.SecurableKindManifest)
	if err != nil {
		return nil, err
	}
	if securableKindManifestPb != nil {
		pb.SecurableKindManifest = securableKindManifestPb
	}
	pb.SqlPath = st.SqlPath
	pb.StorageCredentialName = st.StorageCredentialName
	pb.StorageLocation = st.StorageLocation

	var tableConstraintsPb []catalogpb.TableConstraintPb
	for _, item := range st.TableConstraints {
		itemPb, err := TableConstraintToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tableConstraintsPb = append(tableConstraintsPb, *itemPb)
		}
	}
	pb.TableConstraints = tableConstraintsPb
	pb.TableId = st.TableId
	tableTypePb, err := TableTypeToPb(&st.TableType)
	if err != nil {
		return nil, err
	}
	if tableTypePb != nil {
		pb.TableType = *tableTypePb
	}
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy
	pb.ViewDefinition = st.ViewDefinition
	viewDependenciesPb, err := DependencyListToPb(st.ViewDependencies)
	if err != nil {
		return nil, err
	}
	if viewDependenciesPb != nil {
		pb.ViewDependencies = viewDependenciesPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TableInfoFromPb(pb *catalogpb.TableInfoPb) (*TableInfo, error) {
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
	dataSourceFormatField, err := DataSourceFormatFromPb(&pb.DataSourceFormat)
	if err != nil {
		return nil, err
	}
	if dataSourceFormatField != nil {
		st.DataSourceFormat = *dataSourceFormatField
	}
	st.DeletedAt = pb.DeletedAt
	deltaRuntimePropertiesKvpairsField, err := DeltaRuntimePropertiesKvPairsFromPb(pb.DeltaRuntimePropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if deltaRuntimePropertiesKvpairsField != nil {
		st.DeltaRuntimePropertiesKvpairs = deltaRuntimePropertiesKvpairsField
	}
	effectivePredictiveOptimizationFlagField, err := EffectivePredictiveOptimizationFlagFromPb(pb.EffectivePredictiveOptimizationFlag)
	if err != nil {
		return nil, err
	}
	if effectivePredictiveOptimizationFlagField != nil {
		st.EffectivePredictiveOptimizationFlag = effectivePredictiveOptimizationFlagField
	}
	enablePredictiveOptimizationField, err := EnablePredictiveOptimizationFromPb(&pb.EnablePredictiveOptimization)
	if err != nil {
		return nil, err
	}
	if enablePredictiveOptimizationField != nil {
		st.EnablePredictiveOptimization = *enablePredictiveOptimizationField
	}
	encryptionDetailsField, err := EncryptionDetailsFromPb(pb.EncryptionDetails)
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
	rowFilterField, err := TableRowFilterFromPb(pb.RowFilter)
	if err != nil {
		return nil, err
	}
	if rowFilterField != nil {
		st.RowFilter = rowFilterField
	}
	st.SchemaName = pb.SchemaName
	securableKindManifestField, err := SecurableKindManifestFromPb(pb.SecurableKindManifest)
	if err != nil {
		return nil, err
	}
	if securableKindManifestField != nil {
		st.SecurableKindManifest = securableKindManifestField
	}
	st.SqlPath = pb.SqlPath
	st.StorageCredentialName = pb.StorageCredentialName
	st.StorageLocation = pb.StorageLocation

	var tableConstraintsField []TableConstraint
	for _, itemPb := range pb.TableConstraints {
		item, err := TableConstraintFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tableConstraintsField = append(tableConstraintsField, *item)
		}
	}
	st.TableConstraints = tableConstraintsField
	st.TableId = pb.TableId
	tableTypeField, err := TableTypeFromPb(&pb.TableType)
	if err != nil {
		return nil, err
	}
	if tableTypeField != nil {
		st.TableType = *tableTypeField
	}
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.ViewDefinition = pb.ViewDefinition
	viewDependenciesField, err := DependencyListFromPb(pb.ViewDependencies)
	if err != nil {
		return nil, err
	}
	if viewDependenciesField != nil {
		st.ViewDependencies = viewDependenciesField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type TableOperation string

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

// Values returns all possible values for TableOperation.
//
// There is no guarantee on the order of the values in the slice.
func (f *TableOperation) Values() []TableOperation {
	return []TableOperation{
		TableOperationRead,
		TableOperationReadWrite,
	}
}

// Type always returns TableOperation to satisfy [pflag.Value] interface
func (f *TableOperation) Type() string {
	return "TableOperation"
}

func TableOperationToPb(st *TableOperation) (*catalogpb.TableOperationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.TableOperationPb(*st)
	return &pb, nil
}

func TableOperationFromPb(pb *catalogpb.TableOperationPb) (*TableOperation, error) {
	if pb == nil {
		return nil, nil
	}
	st := TableOperation(*pb)
	return &st, nil
}

type TableRowFilter struct {
	// The full name of the row filter SQL UDF.
	// Wire name: 'function_name'
	FunctionName string ``
	// The list of table columns to be passed as input to the row filter
	// function. The column types should match the types of the filter function
	// arguments.
	// Wire name: 'input_column_names'
	InputColumnNames []string ``
}

func TableRowFilterToPb(st *TableRowFilter) (*catalogpb.TableRowFilterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.TableRowFilterPb{}
	pb.FunctionName = st.FunctionName
	pb.InputColumnNames = st.InputColumnNames

	return pb, nil
}

func TableRowFilterFromPb(pb *catalogpb.TableRowFilterPb) (*TableRowFilter, error) {
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
	FullName string ``
	// SecurableKindManifest of table, including capabilities the table has.
	// Wire name: 'securable_kind_manifest'
	SecurableKindManifest *SecurableKindManifest ``

	// Wire name: 'table_type'
	TableType       TableType ``
	ForceSendFields []string  `tf:"-"`
}

func (s *TableSummary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableSummary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TableSummaryToPb(st *TableSummary) (*catalogpb.TableSummaryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.TableSummaryPb{}
	pb.FullName = st.FullName
	securableKindManifestPb, err := SecurableKindManifestToPb(st.SecurableKindManifest)
	if err != nil {
		return nil, err
	}
	if securableKindManifestPb != nil {
		pb.SecurableKindManifest = securableKindManifestPb
	}
	tableTypePb, err := TableTypeToPb(&st.TableType)
	if err != nil {
		return nil, err
	}
	if tableTypePb != nil {
		pb.TableType = *tableTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TableSummaryFromPb(pb *catalogpb.TableSummaryPb) (*TableSummary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableSummary{}
	st.FullName = pb.FullName
	securableKindManifestField, err := SecurableKindManifestFromPb(pb.SecurableKindManifest)
	if err != nil {
		return nil, err
	}
	if securableKindManifestField != nil {
		st.SecurableKindManifest = securableKindManifestField
	}
	tableTypeField, err := TableTypeFromPb(&pb.TableType)
	if err != nil {
		return nil, err
	}
	if tableTypeField != nil {
		st.TableType = *tableTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type TableType string

const TableTypeExternal TableType = `EXTERNAL`

const TableTypeExternalShallowClone TableType = `EXTERNAL_SHALLOW_CLONE`

const TableTypeForeign TableType = `FOREIGN`

const TableTypeManaged TableType = `MANAGED`

const TableTypeManagedShallowClone TableType = `MANAGED_SHALLOW_CLONE`

const TableTypeMaterializedView TableType = `MATERIALIZED_VIEW`

const TableTypeMetricView TableType = `METRIC_VIEW`

const TableTypeStreamingTable TableType = `STREAMING_TABLE`

const TableTypeView TableType = `VIEW`

// String representation for [fmt.Print]
func (f *TableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TableType) Set(v string) error {
	switch v {
	case `EXTERNAL`, `EXTERNAL_SHALLOW_CLONE`, `FOREIGN`, `MANAGED`, `MANAGED_SHALLOW_CLONE`, `MATERIALIZED_VIEW`, `METRIC_VIEW`, `STREAMING_TABLE`, `VIEW`:
		*f = TableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL", "EXTERNAL_SHALLOW_CLONE", "FOREIGN", "MANAGED", "MANAGED_SHALLOW_CLONE", "MATERIALIZED_VIEW", "METRIC_VIEW", "STREAMING_TABLE", "VIEW"`, v)
	}
}

// Values returns all possible values for TableType.
//
// There is no guarantee on the order of the values in the slice.
func (f *TableType) Values() []TableType {
	return []TableType{
		TableTypeExternal,
		TableTypeExternalShallowClone,
		TableTypeForeign,
		TableTypeManaged,
		TableTypeManagedShallowClone,
		TableTypeMaterializedView,
		TableTypeMetricView,
		TableTypeStreamingTable,
		TableTypeView,
	}
}

// Type always returns TableType to satisfy [pflag.Value] interface
func (f *TableType) Type() string {
	return "TableType"
}

func TableTypeToPb(st *TableType) (*catalogpb.TableTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.TableTypePb(*st)
	return &pb, nil
}

func TableTypeFromPb(pb *catalogpb.TableTypePb) (*TableType, error) {
	if pb == nil {
		return nil, nil
	}
	st := TableType(*pb)
	return &st, nil
}

type TagKeyValue struct {
	// name of the tag
	// Wire name: 'key'
	Key string ``
	// value of the tag associated with the key, could be optional
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *TagKeyValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TagKeyValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TagKeyValueToPb(st *TagKeyValue) (*catalogpb.TagKeyValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.TagKeyValuePb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TagKeyValueFromPb(pb *catalogpb.TagKeyValuePb) (*TagKeyValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TagKeyValue{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type TemporaryCredentials struct {

	// Wire name: 'aws_temp_credentials'
	AwsTempCredentials *AwsCredentials ``

	// Wire name: 'azure_aad'
	AzureAad *AzureActiveDirectoryToken ``
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	// Wire name: 'expiration_time'
	ExpirationTime int64 ``

	// Wire name: 'gcp_oauth_token'
	GcpOauthToken   *GcpOauthToken ``
	ForceSendFields []string       `tf:"-"`
}

func (s *TemporaryCredentials) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TemporaryCredentials) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TemporaryCredentialsToPb(st *TemporaryCredentials) (*catalogpb.TemporaryCredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.TemporaryCredentialsPb{}
	awsTempCredentialsPb, err := AwsCredentialsToPb(st.AwsTempCredentials)
	if err != nil {
		return nil, err
	}
	if awsTempCredentialsPb != nil {
		pb.AwsTempCredentials = awsTempCredentialsPb
	}
	azureAadPb, err := AzureActiveDirectoryTokenToPb(st.AzureAad)
	if err != nil {
		return nil, err
	}
	if azureAadPb != nil {
		pb.AzureAad = azureAadPb
	}
	pb.ExpirationTime = st.ExpirationTime
	gcpOauthTokenPb, err := GcpOauthTokenToPb(st.GcpOauthToken)
	if err != nil {
		return nil, err
	}
	if gcpOauthTokenPb != nil {
		pb.GcpOauthToken = gcpOauthTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TemporaryCredentialsFromPb(pb *catalogpb.TemporaryCredentialsPb) (*TemporaryCredentials, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TemporaryCredentials{}
	awsTempCredentialsField, err := AwsCredentialsFromPb(pb.AwsTempCredentials)
	if err != nil {
		return nil, err
	}
	if awsTempCredentialsField != nil {
		st.AwsTempCredentials = awsTempCredentialsField
	}
	azureAadField, err := AzureActiveDirectoryTokenFromPb(pb.AzureAad)
	if err != nil {
		return nil, err
	}
	if azureAadField != nil {
		st.AzureAad = azureAadField
	}
	st.ExpirationTime = pb.ExpirationTime
	gcpOauthTokenField, err := GcpOauthTokenFromPb(pb.GcpOauthToken)
	if err != nil {
		return nil, err
	}
	if gcpOauthTokenField != nil {
		st.GcpOauthToken = gcpOauthTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
type TriggeredUpdateStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	// Wire name: 'last_processed_commit_version'
	LastProcessedCommitVersion int64 ``
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	// Wire name: 'timestamp'
	Timestamp *time.Time ``
	// Progress of the active data synchronization pipeline.
	// Wire name: 'triggered_update_progress'
	TriggeredUpdateProgress *PipelineProgress ``
	ForceSendFields         []string          `tf:"-"`
}

func (s *TriggeredUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TriggeredUpdateStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TriggeredUpdateStatusToPb(st *TriggeredUpdateStatus) (*catalogpb.TriggeredUpdateStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.TriggeredUpdateStatusPb{}
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion
	timestampPb, err := timestampToPb(st.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampPb != nil {
		pb.Timestamp = *timestampPb
	}
	triggeredUpdateProgressPb, err := PipelineProgressToPb(st.TriggeredUpdateProgress)
	if err != nil {
		return nil, err
	}
	if triggeredUpdateProgressPb != nil {
		pb.TriggeredUpdateProgress = triggeredUpdateProgressPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TriggeredUpdateStatusFromPb(pb *catalogpb.TriggeredUpdateStatusPb) (*TriggeredUpdateStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TriggeredUpdateStatus{}
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
	timestampField, err := timestampFromPb(&pb.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampField != nil {
		st.Timestamp = timestampField
	}
	triggeredUpdateProgressField, err := PipelineProgressFromPb(pb.TriggeredUpdateProgress)
	if err != nil {
		return nil, err
	}
	if triggeredUpdateProgressField != nil {
		st.TriggeredUpdateProgress = triggeredUpdateProgressField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UnassignRequest struct {
	// Query for the ID of the metastore to delete.
	// Wire name: 'metastore_id'
	MetastoreId string `tf:"-"`
	// A workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func UnassignRequestToPb(st *UnassignRequest) (*catalogpb.UnassignRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UnassignRequestPb{}
	pb.MetastoreId = st.MetastoreId
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func UnassignRequestFromPb(pb *catalogpb.UnassignRequestPb) (*UnassignRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UnassignRequest{}
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

type UpdateCatalog struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'enable_predictive_optimization'
	EnablePredictiveOptimization EnablePredictiveOptimization ``
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode CatalogIsolationMode ``
	// The name of the catalog.
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the catalog.
	// Wire name: 'new_name'
	NewName string ``
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string ``
	// Username of current owner of catalog.
	// Wire name: 'owner'
	Owner string ``
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties      map[string]string ``
	ForceSendFields []string          `tf:"-"`
}

func (s *UpdateCatalog) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateCatalog) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateCatalogToPb(st *UpdateCatalog) (*catalogpb.UpdateCatalogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateCatalogPb{}
	pb.Comment = st.Comment
	enablePredictiveOptimizationPb, err := EnablePredictiveOptimizationToPb(&st.EnablePredictiveOptimization)
	if err != nil {
		return nil, err
	}
	if enablePredictiveOptimizationPb != nil {
		pb.EnablePredictiveOptimization = *enablePredictiveOptimizationPb
	}
	isolationModePb, err := CatalogIsolationModeToPb(&st.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Options = st.Options
	pb.Owner = st.Owner
	pb.Properties = st.Properties

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateCatalogFromPb(pb *catalogpb.UpdateCatalogPb) (*UpdateCatalog, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCatalog{}
	st.Comment = pb.Comment
	enablePredictiveOptimizationField, err := EnablePredictiveOptimizationFromPb(&pb.EnablePredictiveOptimization)
	if err != nil {
		return nil, err
	}
	if enablePredictiveOptimizationField != nil {
		st.EnablePredictiveOptimization = *enablePredictiveOptimizationField
	}
	isolationModeField, err := CatalogIsolationModeFromPb(&pb.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Options = pb.Options
	st.Owner = pb.Owner
	st.Properties = pb.Properties

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateCatalogWorkspaceBindingsResponse struct {
	// A list of workspace IDs
	// Wire name: 'workspaces'
	Workspaces []int64 ``
}

func UpdateCatalogWorkspaceBindingsResponseToPb(st *UpdateCatalogWorkspaceBindingsResponse) (*catalogpb.UpdateCatalogWorkspaceBindingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateCatalogWorkspaceBindingsResponsePb{}
	pb.Workspaces = st.Workspaces

	return pb, nil
}

func UpdateCatalogWorkspaceBindingsResponseFromPb(pb *catalogpb.UpdateCatalogWorkspaceBindingsResponsePb) (*UpdateCatalogWorkspaceBindingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCatalogWorkspaceBindingsResponse{}
	st.Workspaces = pb.Workspaces

	return st, nil
}

type UpdateConnection struct {
	// [Create,Update:OPT] Connection environment settings as
	// EnvironmentSettings object.
	// Wire name: 'environment_settings'
	EnvironmentSettings *EnvironmentSettings ``
	// Name of the connection.
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the connection.
	// Wire name: 'new_name'
	NewName string ``
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string ``
	// Username of current owner of the connection.
	// Wire name: 'owner'
	Owner           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateConnection) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateConnection) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateConnectionToPb(st *UpdateConnection) (*catalogpb.UpdateConnectionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateConnectionPb{}
	environmentSettingsPb, err := EnvironmentSettingsToPb(st.EnvironmentSettings)
	if err != nil {
		return nil, err
	}
	if environmentSettingsPb != nil {
		pb.EnvironmentSettings = environmentSettingsPb
	}
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Options = st.Options
	pb.Owner = st.Owner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateConnectionFromPb(pb *catalogpb.UpdateConnectionPb) (*UpdateConnection, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateConnection{}
	environmentSettingsField, err := EnvironmentSettingsFromPb(pb.EnvironmentSettings)
	if err != nil {
		return nil, err
	}
	if environmentSettingsField != nil {
		st.EnvironmentSettings = environmentSettingsField
	}
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Options = pb.Options
	st.Owner = pb.Owner

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateCredentialRequest struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRole ``
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentity ``
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal ``
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string ``
	// The Databricks managed GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount ``
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	// Wire name: 'force'
	Force bool ``
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode ``
	// Name of the credential.
	// Wire name: 'name_arg'
	NameArg string `tf:"-"`
	// New name of credential.
	// Wire name: 'new_name'
	NewName string ``
	// Username of current owner of credential.
	// Wire name: 'owner'
	Owner string ``
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	// Wire name: 'read_only'
	ReadOnly bool ``
	// Supply true to this argument to skip validation of the updated
	// credential.
	// Wire name: 'skip_validation'
	SkipValidation  bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateCredentialRequestToPb(st *UpdateCredentialRequest) (*catalogpb.UpdateCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateCredentialRequestPb{}
	awsIamRolePb, err := AwsIamRoleToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}
	azureManagedIdentityPb, err := AzureManagedIdentityToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}
	azureServicePrincipalPb, err := AzureServicePrincipalToPb(st.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalPb != nil {
		pb.AzureServicePrincipal = azureServicePrincipalPb
	}
	pb.Comment = st.Comment
	databricksGcpServiceAccountPb, err := DatabricksGcpServiceAccountToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}
	pb.Force = st.Force
	isolationModePb, err := IsolationModeToPb(&st.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}
	pb.NameArg = st.NameArg
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	pb.ReadOnly = st.ReadOnly
	pb.SkipValidation = st.SkipValidation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateCredentialRequestFromPb(pb *catalogpb.UpdateCredentialRequestPb) (*UpdateCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCredentialRequest{}
	awsIamRoleField, err := AwsIamRoleFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := AzureManagedIdentityFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	azureServicePrincipalField, err := AzureServicePrincipalFromPb(pb.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalField != nil {
		st.AzureServicePrincipal = azureServicePrincipalField
	}
	st.Comment = pb.Comment
	databricksGcpServiceAccountField, err := DatabricksGcpServiceAccountFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	st.Force = pb.Force
	isolationModeField, err := IsolationModeFromPb(&pb.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
	st.NameArg = pb.NameArg
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateExternalLineageRelationshipRequest struct {

	// Wire name: 'external_lineage_relationship'
	ExternalLineageRelationship UpdateRequestExternalLineage ``
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	// Wire name: 'update_mask'
	UpdateMask []string `tf:"-"`
}

func UpdateExternalLineageRelationshipRequestToPb(st *UpdateExternalLineageRelationshipRequest) (*catalogpb.UpdateExternalLineageRelationshipRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateExternalLineageRelationshipRequestPb{}
	externalLineageRelationshipPb, err := UpdateRequestExternalLineageToPb(&st.ExternalLineageRelationship)
	if err != nil {
		return nil, err
	}
	if externalLineageRelationshipPb != nil {
		pb.ExternalLineageRelationship = *externalLineageRelationshipPb
	}
	updateMaskPb, err := fieldMaskToPb(&st.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskPb != nil {
		pb.UpdateMask = *updateMaskPb
	}

	return pb, nil
}

func UpdateExternalLineageRelationshipRequestFromPb(pb *catalogpb.UpdateExternalLineageRelationshipRequestPb) (*UpdateExternalLineageRelationshipRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExternalLineageRelationshipRequest{}
	externalLineageRelationshipField, err := UpdateRequestExternalLineageFromPb(&pb.ExternalLineageRelationship)
	if err != nil {
		return nil, err
	}
	if externalLineageRelationshipField != nil {
		st.ExternalLineageRelationship = *externalLineageRelationshipField
	}
	updateMaskField, err := fieldMaskFromPb(&pb.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskField != nil {
		st.UpdateMask = *updateMaskField
	}

	return st, nil
}

type UpdateExternalLocation struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// Name of the storage credential used with this location.
	// Wire name: 'credential_name'
	CredentialName string ``
	// Whether to enable file events on this external location.
	// Wire name: 'enable_file_events'
	EnableFileEvents bool ``

	// Wire name: 'encryption_details'
	EncryptionDetails *EncryptionDetails ``
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	// Wire name: 'fallback'
	Fallback bool ``
	// File event queue settings.
	// Wire name: 'file_event_queue'
	FileEventQueue *FileEventQueue ``
	// Force update even if changing url invalidates dependent external tables
	// or mounts.
	// Wire name: 'force'
	Force bool ``

	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode ``
	// Name of the external location.
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the external location.
	// Wire name: 'new_name'
	NewName string ``
	// The owner of the external location.
	// Wire name: 'owner'
	Owner string ``
	// Indicates whether the external location is read-only.
	// Wire name: 'read_only'
	ReadOnly bool ``
	// Skips validation of the storage credential associated with the external
	// location.
	// Wire name: 'skip_validation'
	SkipValidation bool ``
	// Path URL of the external location.
	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateExternalLocation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateExternalLocation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateExternalLocationToPb(st *UpdateExternalLocation) (*catalogpb.UpdateExternalLocationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateExternalLocationPb{}
	pb.Comment = st.Comment
	pb.CredentialName = st.CredentialName
	pb.EnableFileEvents = st.EnableFileEvents
	encryptionDetailsPb, err := EncryptionDetailsToPb(st.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsPb != nil {
		pb.EncryptionDetails = encryptionDetailsPb
	}
	pb.Fallback = st.Fallback
	fileEventQueuePb, err := FileEventQueueToPb(st.FileEventQueue)
	if err != nil {
		return nil, err
	}
	if fileEventQueuePb != nil {
		pb.FileEventQueue = fileEventQueuePb
	}
	pb.Force = st.Force
	isolationModePb, err := IsolationModeToPb(&st.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	pb.ReadOnly = st.ReadOnly
	pb.SkipValidation = st.SkipValidation
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateExternalLocationFromPb(pb *catalogpb.UpdateExternalLocationPb) (*UpdateExternalLocation, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExternalLocation{}
	st.Comment = pb.Comment
	st.CredentialName = pb.CredentialName
	st.EnableFileEvents = pb.EnableFileEvents
	encryptionDetailsField, err := EncryptionDetailsFromPb(pb.EncryptionDetails)
	if err != nil {
		return nil, err
	}
	if encryptionDetailsField != nil {
		st.EncryptionDetails = encryptionDetailsField
	}
	st.Fallback = pb.Fallback
	fileEventQueueField, err := FileEventQueueFromPb(pb.FileEventQueue)
	if err != nil {
		return nil, err
	}
	if fileEventQueueField != nil {
		st.FileEventQueue = fileEventQueueField
	}
	st.Force = pb.Force
	isolationModeField, err := IsolationModeFromPb(&pb.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateExternalMetadataRequest struct {

	// Wire name: 'external_metadata'
	ExternalMetadata ExternalMetadata ``
	// Name of the external metadata object.
	// Wire name: 'name'
	Name string `tf:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	// Wire name: 'update_mask'
	UpdateMask []string `tf:"-"`
}

func UpdateExternalMetadataRequestToPb(st *UpdateExternalMetadataRequest) (*catalogpb.UpdateExternalMetadataRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateExternalMetadataRequestPb{}
	externalMetadataPb, err := ExternalMetadataToPb(&st.ExternalMetadata)
	if err != nil {
		return nil, err
	}
	if externalMetadataPb != nil {
		pb.ExternalMetadata = *externalMetadataPb
	}
	pb.Name = st.Name
	updateMaskPb, err := fieldMaskToPb(&st.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskPb != nil {
		pb.UpdateMask = *updateMaskPb
	}

	return pb, nil
}

func UpdateExternalMetadataRequestFromPb(pb *catalogpb.UpdateExternalMetadataRequestPb) (*UpdateExternalMetadataRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExternalMetadataRequest{}
	externalMetadataField, err := ExternalMetadataFromPb(&pb.ExternalMetadata)
	if err != nil {
		return nil, err
	}
	if externalMetadataField != nil {
		st.ExternalMetadata = *externalMetadataField
	}
	st.Name = pb.Name
	updateMaskField, err := fieldMaskFromPb(&pb.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskField != nil {
		st.UpdateMask = *updateMaskField
	}

	return st, nil
}

type UpdateFunction struct {
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	// Wire name: 'name'
	Name string `tf:"-"`
	// Username of current owner of function.
	// Wire name: 'owner'
	Owner           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateFunctionToPb(st *UpdateFunction) (*catalogpb.UpdateFunctionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateFunctionPb{}
	pb.Name = st.Name
	pb.Owner = st.Owner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateFunctionFromPb(pb *catalogpb.UpdateFunctionPb) (*UpdateFunction, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateFunction{}
	st.Name = pb.Name
	st.Owner = pb.Owner

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateMetastore struct {
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	// Wire name: 'delta_sharing_organization_name'
	DeltaSharingOrganizationName string ``
	// The lifetime of delta sharing recipient token in seconds.
	// Wire name: 'delta_sharing_recipient_token_lifetime_in_seconds'
	DeltaSharingRecipientTokenLifetimeInSeconds int64 ``
	// The scope of Delta Sharing enabled for the metastore.
	// Wire name: 'delta_sharing_scope'
	DeltaSharingScope DeltaSharingScopeEnum ``
	// Unique ID of the metastore.
	// Wire name: 'id'
	Id string `tf:"-"`
	// New name for the metastore.
	// Wire name: 'new_name'
	NewName string ``
	// The owner of the metastore.
	// Wire name: 'owner'
	Owner string ``
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	// Wire name: 'privilege_model_version'
	PrivilegeModelVersion string ``
	// UUID of storage credential to access the metastore storage_root.
	// Wire name: 'storage_root_credential_id'
	StorageRootCredentialId string   ``
	ForceSendFields         []string `tf:"-"`
}

func (s *UpdateMetastore) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateMetastore) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateMetastoreToPb(st *UpdateMetastore) (*catalogpb.UpdateMetastorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateMetastorePb{}
	pb.DeltaSharingOrganizationName = st.DeltaSharingOrganizationName
	pb.DeltaSharingRecipientTokenLifetimeInSeconds = st.DeltaSharingRecipientTokenLifetimeInSeconds
	deltaSharingScopePb, err := DeltaSharingScopeEnumToPb(&st.DeltaSharingScope)
	if err != nil {
		return nil, err
	}
	if deltaSharingScopePb != nil {
		pb.DeltaSharingScope = *deltaSharingScopePb
	}
	pb.Id = st.Id
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	pb.PrivilegeModelVersion = st.PrivilegeModelVersion
	pb.StorageRootCredentialId = st.StorageRootCredentialId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateMetastoreFromPb(pb *catalogpb.UpdateMetastorePb) (*UpdateMetastore, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateMetastore{}
	st.DeltaSharingOrganizationName = pb.DeltaSharingOrganizationName
	st.DeltaSharingRecipientTokenLifetimeInSeconds = pb.DeltaSharingRecipientTokenLifetimeInSeconds
	deltaSharingScopeField, err := DeltaSharingScopeEnumFromPb(&pb.DeltaSharingScope)
	if err != nil {
		return nil, err
	}
	if deltaSharingScopeField != nil {
		st.DeltaSharingScope = *deltaSharingScopeField
	}
	st.Id = pb.Id
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.PrivilegeModelVersion = pb.PrivilegeModelVersion
	st.StorageRootCredentialId = pb.StorageRootCredentialId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateMetastoreAssignment struct {
	// The name of the default catalog in the metastore. This field is
	// deprecated. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	// Wire name: 'default_catalog_name'
	DefaultCatalogName string ``
	// The unique ID of the metastore.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// A workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId     int64    `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateMetastoreAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateMetastoreAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateMetastoreAssignmentToPb(st *UpdateMetastoreAssignment) (*catalogpb.UpdateMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateMetastoreAssignmentPb{}
	pb.DefaultCatalogName = st.DefaultCatalogName
	pb.MetastoreId = st.MetastoreId
	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateMetastoreAssignmentFromPb(pb *catalogpb.UpdateMetastoreAssignmentPb) (*UpdateMetastoreAssignment, error) {
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

type UpdateModelVersionRequest struct {
	// The comment attached to the model version
	// Wire name: 'comment'
	Comment string ``
	// The three-level (fully qualified) name of the model version
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// The integer version number of the model version
	// Wire name: 'version'
	Version         int      `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateModelVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateModelVersionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateModelVersionRequestToPb(st *UpdateModelVersionRequest) (*catalogpb.UpdateModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateModelVersionRequestPb{}
	pb.Comment = st.Comment
	pb.FullName = st.FullName
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateModelVersionRequestFromPb(pb *catalogpb.UpdateModelVersionRequestPb) (*UpdateModelVersionRequest, error) {
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

type UpdateMonitor struct {
	// [Create:OPT Update:OPT] Baseline table name. Baseline data is used to
	// compute drift from the data in the monitored `table_name`. The baseline
	// table and the monitored table shall have the same schema.
	// Wire name: 'baseline_table_name'
	BaselineTableName string ``
	// [Create:OPT Update:OPT] Custom metrics.
	// Wire name: 'custom_metrics'
	CustomMetrics []MonitorMetric ``
	// [Create:ERR Update:OPT] Id of dashboard that visualizes the computed
	// metrics. This can be empty if the monitor is in PENDING state.
	// Wire name: 'dashboard_id'
	DashboardId string ``
	// [Create:OPT Update:OPT] Data classification related config.
	// Wire name: 'data_classification_config'
	DataClassificationConfig *MonitorDataClassificationConfig ``

	// Wire name: 'inference_log'
	InferenceLog *MonitorInferenceLog ``
	// [Create:ERR Update:IGN] The latest error message for a monitor failure.
	// Wire name: 'latest_monitor_failure_msg'
	LatestMonitorFailureMsg string ``
	// [Create:OPT Update:OPT] Field for specifying notification settings.
	// Wire name: 'notifications'
	Notifications *MonitorNotifications ``
	// [Create:REQ Update:REQ] Schema where output tables are created. Needs to
	// be in 2-level format {catalog}.{schema}
	// Wire name: 'output_schema_name'
	OutputSchemaName string ``
	// [Create:OPT Update:OPT] The monitor schedule.
	// Wire name: 'schedule'
	Schedule *MonitorCronSchedule ``
	// [Create:OPT Update:OPT] List of column expressions to slice data with for
	// targeted analysis. The data is grouped by each expression independently,
	// resulting in a separate slice for each predicate and its complements. For
	// example `slicing_exprs=[“col_1”, “col_2 > 10”]` will generate the
	// following slices: two slices for `col_2 > 10` (True and False), and one
	// slice per unique value in `col1`. For high-cardinality columns, only the
	// top 100 unique values by frequency will generate slices.
	// Wire name: 'slicing_exprs'
	SlicingExprs []string ``
	// Configuration for monitoring snapshot tables.
	// Wire name: 'snapshot'
	Snapshot *MonitorSnapshot ``
	// UC table name in format `catalog.schema.table_name`. This field
	// corresponds to the {full_table_name_arg} arg in the endpoint path.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
	// Configuration for monitoring time series tables.
	// Wire name: 'time_series'
	TimeSeries      *MonitorTimeSeries ``
	ForceSendFields []string           `tf:"-"`
}

func (s *UpdateMonitor) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateMonitor) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateMonitorToPb(st *UpdateMonitor) (*catalogpb.UpdateMonitorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateMonitorPb{}
	pb.BaselineTableName = st.BaselineTableName

	var customMetricsPb []catalogpb.MonitorMetricPb
	for _, item := range st.CustomMetrics {
		itemPb, err := MonitorMetricToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customMetricsPb = append(customMetricsPb, *itemPb)
		}
	}
	pb.CustomMetrics = customMetricsPb
	pb.DashboardId = st.DashboardId
	dataClassificationConfigPb, err := MonitorDataClassificationConfigToPb(st.DataClassificationConfig)
	if err != nil {
		return nil, err
	}
	if dataClassificationConfigPb != nil {
		pb.DataClassificationConfig = dataClassificationConfigPb
	}
	inferenceLogPb, err := MonitorInferenceLogToPb(st.InferenceLog)
	if err != nil {
		return nil, err
	}
	if inferenceLogPb != nil {
		pb.InferenceLog = inferenceLogPb
	}
	pb.LatestMonitorFailureMsg = st.LatestMonitorFailureMsg
	notificationsPb, err := MonitorNotificationsToPb(st.Notifications)
	if err != nil {
		return nil, err
	}
	if notificationsPb != nil {
		pb.Notifications = notificationsPb
	}
	pb.OutputSchemaName = st.OutputSchemaName
	schedulePb, err := MonitorCronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}
	pb.SlicingExprs = st.SlicingExprs
	snapshotPb, err := MonitorSnapshotToPb(st.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotPb != nil {
		pb.Snapshot = snapshotPb
	}
	pb.TableName = st.TableName
	timeSeriesPb, err := MonitorTimeSeriesToPb(st.TimeSeries)
	if err != nil {
		return nil, err
	}
	if timeSeriesPb != nil {
		pb.TimeSeries = timeSeriesPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateMonitorFromPb(pb *catalogpb.UpdateMonitorPb) (*UpdateMonitor, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateMonitor{}
	st.BaselineTableName = pb.BaselineTableName

	var customMetricsField []MonitorMetric
	for _, itemPb := range pb.CustomMetrics {
		item, err := MonitorMetricFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customMetricsField = append(customMetricsField, *item)
		}
	}
	st.CustomMetrics = customMetricsField
	st.DashboardId = pb.DashboardId
	dataClassificationConfigField, err := MonitorDataClassificationConfigFromPb(pb.DataClassificationConfig)
	if err != nil {
		return nil, err
	}
	if dataClassificationConfigField != nil {
		st.DataClassificationConfig = dataClassificationConfigField
	}
	inferenceLogField, err := MonitorInferenceLogFromPb(pb.InferenceLog)
	if err != nil {
		return nil, err
	}
	if inferenceLogField != nil {
		st.InferenceLog = inferenceLogField
	}
	st.LatestMonitorFailureMsg = pb.LatestMonitorFailureMsg
	notificationsField, err := MonitorNotificationsFromPb(pb.Notifications)
	if err != nil {
		return nil, err
	}
	if notificationsField != nil {
		st.Notifications = notificationsField
	}
	st.OutputSchemaName = pb.OutputSchemaName
	scheduleField, err := MonitorCronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	st.SlicingExprs = pb.SlicingExprs
	snapshotField, err := MonitorSnapshotFromPb(pb.Snapshot)
	if err != nil {
		return nil, err
	}
	if snapshotField != nil {
		st.Snapshot = snapshotField
	}
	st.TableName = pb.TableName
	timeSeriesField, err := MonitorTimeSeriesFromPb(pb.TimeSeries)
	if err != nil {
		return nil, err
	}
	if timeSeriesField != nil {
		st.TimeSeries = timeSeriesField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdatePermissions struct {
	// Array of permissions change objects.
	// Wire name: 'changes'
	Changes []PermissionsChange ``
	// Full name of securable.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// Type of securable.
	// Wire name: 'securable_type'
	SecurableType string `tf:"-"`
}

func UpdatePermissionsToPb(st *UpdatePermissions) (*catalogpb.UpdatePermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdatePermissionsPb{}

	var changesPb []catalogpb.PermissionsChangePb
	for _, item := range st.Changes {
		itemPb, err := PermissionsChangeToPb(&item)
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

func UpdatePermissionsFromPb(pb *catalogpb.UpdatePermissionsPb) (*UpdatePermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePermissions{}

	var changesField []PermissionsChange
	for _, itemPb := range pb.Changes {
		item, err := PermissionsChangeFromPb(&itemPb)
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

type UpdatePermissionsResponse struct {
	// The privileges assigned to each principal
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment ``
}

func UpdatePermissionsResponseToPb(st *UpdatePermissionsResponse) (*catalogpb.UpdatePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdatePermissionsResponsePb{}

	var privilegeAssignmentsPb []catalogpb.PrivilegeAssignmentPb
	for _, item := range st.PrivilegeAssignments {
		itemPb, err := PrivilegeAssignmentToPb(&item)
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

func UpdatePermissionsResponseFromPb(pb *catalogpb.UpdatePermissionsResponsePb) (*UpdatePermissionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePermissionsResponse{}

	var privilegeAssignmentsField []PrivilegeAssignment
	for _, itemPb := range pb.PrivilegeAssignments {
		item, err := PrivilegeAssignmentFromPb(&itemPb)
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

type UpdateRegisteredModelRequest struct {
	// The comment attached to the registered model
	// Wire name: 'comment'
	Comment string ``
	// The three-level (fully qualified) name of the registered model
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// New name for the registered model.
	// Wire name: 'new_name'
	NewName string ``
	// The identifier of the user who owns the registered model
	// Wire name: 'owner'
	Owner           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateRegisteredModelRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRegisteredModelRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateRegisteredModelRequestToPb(st *UpdateRegisteredModelRequest) (*catalogpb.UpdateRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateRegisteredModelRequestPb{}
	pb.Comment = st.Comment
	pb.FullName = st.FullName
	pb.NewName = st.NewName
	pb.Owner = st.Owner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateRegisteredModelRequestFromPb(pb *catalogpb.UpdateRegisteredModelRequestPb) (*UpdateRegisteredModelRequest, error) {
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

type UpdateRequestExternalLineage struct {
	// List of column relationships between source and target objects.
	// Wire name: 'columns'
	Columns []ColumnRelationship ``
	// Unique identifier of the external lineage relationship.
	// Wire name: 'id'
	Id string ``
	// Key-value properties associated with the external lineage relationship.
	// Wire name: 'properties'
	Properties map[string]string ``
	// Source object of the external lineage relationship.
	// Wire name: 'source'
	Source ExternalLineageObject ``
	// Target object of the external lineage relationship.
	// Wire name: 'target'
	Target          ExternalLineageObject ``
	ForceSendFields []string              `tf:"-"`
}

func (s *UpdateRequestExternalLineage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRequestExternalLineage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateRequestExternalLineageToPb(st *UpdateRequestExternalLineage) (*catalogpb.UpdateRequestExternalLineagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateRequestExternalLineagePb{}

	var columnsPb []catalogpb.ColumnRelationshipPb
	for _, item := range st.Columns {
		itemPb, err := ColumnRelationshipToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			columnsPb = append(columnsPb, *itemPb)
		}
	}
	pb.Columns = columnsPb
	pb.Id = st.Id
	pb.Properties = st.Properties
	sourcePb, err := ExternalLineageObjectToPb(&st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = *sourcePb
	}
	targetPb, err := ExternalLineageObjectToPb(&st.Target)
	if err != nil {
		return nil, err
	}
	if targetPb != nil {
		pb.Target = *targetPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateRequestExternalLineageFromPb(pb *catalogpb.UpdateRequestExternalLineagePb) (*UpdateRequestExternalLineage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRequestExternalLineage{}

	var columnsField []ColumnRelationship
	for _, itemPb := range pb.Columns {
		item, err := ColumnRelationshipFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			columnsField = append(columnsField, *item)
		}
	}
	st.Columns = columnsField
	st.Id = pb.Id
	st.Properties = pb.Properties
	sourceField, err := ExternalLineageObjectFromPb(&pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = *sourceField
	}
	targetField, err := ExternalLineageObjectFromPb(&pb.Target)
	if err != nil {
		return nil, err
	}
	if targetField != nil {
		st.Target = *targetField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateSchema struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'enable_predictive_optimization'
	EnablePredictiveOptimization EnablePredictiveOptimization ``
	// Full name of the schema.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// New name for the schema.
	// Wire name: 'new_name'
	NewName string ``
	// Username of current owner of schema.
	// Wire name: 'owner'
	Owner string ``
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties      map[string]string ``
	ForceSendFields []string          `tf:"-"`
}

func (s *UpdateSchema) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateSchema) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateSchemaToPb(st *UpdateSchema) (*catalogpb.UpdateSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateSchemaPb{}
	pb.Comment = st.Comment
	enablePredictiveOptimizationPb, err := EnablePredictiveOptimizationToPb(&st.EnablePredictiveOptimization)
	if err != nil {
		return nil, err
	}
	if enablePredictiveOptimizationPb != nil {
		pb.EnablePredictiveOptimization = *enablePredictiveOptimizationPb
	}
	pb.FullName = st.FullName
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	pb.Properties = st.Properties

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateSchemaFromPb(pb *catalogpb.UpdateSchemaPb) (*UpdateSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateSchema{}
	st.Comment = pb.Comment
	enablePredictiveOptimizationField, err := EnablePredictiveOptimizationFromPb(&pb.EnablePredictiveOptimization)
	if err != nil {
		return nil, err
	}
	if enablePredictiveOptimizationField != nil {
		st.EnablePredictiveOptimization = *enablePredictiveOptimizationField
	}
	st.FullName = pb.FullName
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.Properties = pb.Properties

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateStorageCredential struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRoleRequest ``
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentityResponse ``
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal ``
	// The Cloudflare API token configuration.
	// Wire name: 'cloudflare_api_token'
	CloudflareApiToken *CloudflareApiToken ``
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string ``
	// The Databricks managed GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest ``
	// Force update even if there are dependent external locations or external
	// tables.
	// Wire name: 'force'
	Force bool ``
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode ``
	// Name of the storage credential.
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the storage credential.
	// Wire name: 'new_name'
	NewName string ``
	// Username of current owner of credential.
	// Wire name: 'owner'
	Owner string ``
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	// Wire name: 'read_only'
	ReadOnly bool ``
	// Supplying true to this argument skips validation of the updated
	// credential.
	// Wire name: 'skip_validation'
	SkipValidation  bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateStorageCredentialToPb(st *UpdateStorageCredential) (*catalogpb.UpdateStorageCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateStorageCredentialPb{}
	awsIamRolePb, err := AwsIamRoleRequestToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}
	azureManagedIdentityPb, err := AzureManagedIdentityResponseToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}
	azureServicePrincipalPb, err := AzureServicePrincipalToPb(st.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalPb != nil {
		pb.AzureServicePrincipal = azureServicePrincipalPb
	}
	cloudflareApiTokenPb, err := CloudflareApiTokenToPb(st.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenPb != nil {
		pb.CloudflareApiToken = cloudflareApiTokenPb
	}
	pb.Comment = st.Comment
	databricksGcpServiceAccountPb, err := DatabricksGcpServiceAccountRequestToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}
	pb.Force = st.Force
	isolationModePb, err := IsolationModeToPb(&st.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModePb != nil {
		pb.IsolationMode = *isolationModePb
	}
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	pb.ReadOnly = st.ReadOnly
	pb.SkipValidation = st.SkipValidation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateStorageCredentialFromPb(pb *catalogpb.UpdateStorageCredentialPb) (*UpdateStorageCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateStorageCredential{}
	awsIamRoleField, err := AwsIamRoleRequestFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := AzureManagedIdentityResponseFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	azureServicePrincipalField, err := AzureServicePrincipalFromPb(pb.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalField != nil {
		st.AzureServicePrincipal = azureServicePrincipalField
	}
	cloudflareApiTokenField, err := CloudflareApiTokenFromPb(pb.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenField != nil {
		st.CloudflareApiToken = cloudflareApiTokenField
	}
	st.Comment = pb.Comment
	databricksGcpServiceAccountField, err := DatabricksGcpServiceAccountRequestFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	st.Force = pb.Force
	isolationModeField, err := IsolationModeFromPb(&pb.IsolationMode)
	if err != nil {
		return nil, err
	}
	if isolationModeField != nil {
		st.IsolationMode = *isolationModeField
	}
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateTableRequest struct {
	// Full name of the table.
	// Wire name: 'full_name'
	FullName string `tf:"-"`
	// Username of current owner of table.
	// Wire name: 'owner'
	Owner           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateTableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateTableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateTableRequestToPb(st *UpdateTableRequest) (*catalogpb.UpdateTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateTableRequestPb{}
	pb.FullName = st.FullName
	pb.Owner = st.Owner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateTableRequestFromPb(pb *catalogpb.UpdateTableRequestPb) (*UpdateTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateTableRequest{}
	st.FullName = pb.FullName
	st.Owner = pb.Owner

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateVolumeRequestContent struct {
	// The comment attached to the volume
	// Wire name: 'comment'
	Comment string ``
	// The three-level (fully qualified) name of the volume
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the volume.
	// Wire name: 'new_name'
	NewName string ``
	// The identifier of the user who owns the volume
	// Wire name: 'owner'
	Owner           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateVolumeRequestContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateVolumeRequestContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateVolumeRequestContentToPb(st *UpdateVolumeRequestContent) (*catalogpb.UpdateVolumeRequestContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateVolumeRequestContentPb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Owner = st.Owner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateVolumeRequestContentFromPb(pb *catalogpb.UpdateVolumeRequestContentPb) (*UpdateVolumeRequestContent, error) {
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

type UpdateWorkspaceBindings struct {
	// A list of workspace IDs.
	// Wire name: 'assign_workspaces'
	AssignWorkspaces []int64 ``
	// The name of the catalog.
	// Wire name: 'name'
	Name string `tf:"-"`
	// A list of workspace IDs.
	// Wire name: 'unassign_workspaces'
	UnassignWorkspaces []int64 ``
}

func UpdateWorkspaceBindingsToPb(st *UpdateWorkspaceBindings) (*catalogpb.UpdateWorkspaceBindingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateWorkspaceBindingsPb{}
	pb.AssignWorkspaces = st.AssignWorkspaces
	pb.Name = st.Name
	pb.UnassignWorkspaces = st.UnassignWorkspaces

	return pb, nil
}

func UpdateWorkspaceBindingsFromPb(pb *catalogpb.UpdateWorkspaceBindingsPb) (*UpdateWorkspaceBindings, error) {
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
	// List of workspace bindings.
	// Wire name: 'add'
	Add []WorkspaceBinding ``
	// List of workspace bindings.
	// Wire name: 'remove'
	Remove []WorkspaceBinding ``
	// The name of the securable.
	// Wire name: 'securable_name'
	SecurableName string `tf:"-"`
	// The type of the securable to bind to a workspace (catalog,
	// storage_credential, credential, or external_location).
	// Wire name: 'securable_type'
	SecurableType string `tf:"-"`
}

func UpdateWorkspaceBindingsParametersToPb(st *UpdateWorkspaceBindingsParameters) (*catalogpb.UpdateWorkspaceBindingsParametersPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateWorkspaceBindingsParametersPb{}

	var addPb []catalogpb.WorkspaceBindingPb
	for _, item := range st.Add {
		itemPb, err := WorkspaceBindingToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			addPb = append(addPb, *itemPb)
		}
	}
	pb.Add = addPb

	var removePb []catalogpb.WorkspaceBindingPb
	for _, item := range st.Remove {
		itemPb, err := WorkspaceBindingToPb(&item)
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

func UpdateWorkspaceBindingsParametersFromPb(pb *catalogpb.UpdateWorkspaceBindingsParametersPb) (*UpdateWorkspaceBindingsParameters, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWorkspaceBindingsParameters{}

	var addField []WorkspaceBinding
	for _, itemPb := range pb.Add {
		item, err := WorkspaceBindingFromPb(&itemPb)
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
		item, err := WorkspaceBindingFromPb(&itemPb)
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

// A list of workspace IDs that are bound to the securable
type UpdateWorkspaceBindingsResponse struct {
	// List of workspace bindings.
	// Wire name: 'bindings'
	Bindings []WorkspaceBinding ``
}

func UpdateWorkspaceBindingsResponseToPb(st *UpdateWorkspaceBindingsResponse) (*catalogpb.UpdateWorkspaceBindingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.UpdateWorkspaceBindingsResponsePb{}

	var bindingsPb []catalogpb.WorkspaceBindingPb
	for _, item := range st.Bindings {
		itemPb, err := WorkspaceBindingToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			bindingsPb = append(bindingsPb, *itemPb)
		}
	}
	pb.Bindings = bindingsPb

	return pb, nil
}

func UpdateWorkspaceBindingsResponseFromPb(pb *catalogpb.UpdateWorkspaceBindingsResponsePb) (*UpdateWorkspaceBindingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWorkspaceBindingsResponse{}

	var bindingsField []WorkspaceBinding
	for _, itemPb := range pb.Bindings {
		item, err := WorkspaceBindingFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			bindingsField = append(bindingsField, *item)
		}
	}
	st.Bindings = bindingsField

	return st, nil
}

// Next ID: 17
type ValidateCredentialRequest struct {

	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRole ``

	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentity ``
	// Required. The name of an existing credential or long-lived cloud
	// credential to validate.
	// Wire name: 'credential_name'
	CredentialName string ``

	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount ``
	// The name of an existing external location to validate. Only applicable
	// for storage credentials (purpose is **STORAGE**.)
	// Wire name: 'external_location_name'
	ExternalLocationName string ``
	// The purpose of the credential. This should only be used when the
	// credential is specified.
	// Wire name: 'purpose'
	Purpose CredentialPurpose ``
	// Whether the credential is only usable for read operations. Only
	// applicable for storage credentials (purpose is **STORAGE**.)
	// Wire name: 'read_only'
	ReadOnly bool ``
	// The external location url to validate. Only applicable when purpose is
	// **STORAGE**.
	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ValidateCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidateCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ValidateCredentialRequestToPb(st *ValidateCredentialRequest) (*catalogpb.ValidateCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ValidateCredentialRequestPb{}
	awsIamRolePb, err := AwsIamRoleToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}
	azureManagedIdentityPb, err := AzureManagedIdentityToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}
	pb.CredentialName = st.CredentialName
	databricksGcpServiceAccountPb, err := DatabricksGcpServiceAccountToPb(st.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountPb != nil {
		pb.DatabricksGcpServiceAccount = databricksGcpServiceAccountPb
	}
	pb.ExternalLocationName = st.ExternalLocationName
	purposePb, err := CredentialPurposeToPb(&st.Purpose)
	if err != nil {
		return nil, err
	}
	if purposePb != nil {
		pb.Purpose = *purposePb
	}
	pb.ReadOnly = st.ReadOnly
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ValidateCredentialRequestFromPb(pb *catalogpb.ValidateCredentialRequestPb) (*ValidateCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ValidateCredentialRequest{}
	awsIamRoleField, err := AwsIamRoleFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := AzureManagedIdentityFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	st.CredentialName = pb.CredentialName
	databricksGcpServiceAccountField, err := DatabricksGcpServiceAccountFromPb(pb.DatabricksGcpServiceAccount)
	if err != nil {
		return nil, err
	}
	if databricksGcpServiceAccountField != nil {
		st.DatabricksGcpServiceAccount = databricksGcpServiceAccountField
	}
	st.ExternalLocationName = pb.ExternalLocationName
	purposeField, err := CredentialPurposeFromPb(&pb.Purpose)
	if err != nil {
		return nil, err
	}
	if purposeField != nil {
		st.Purpose = *purposeField
	}
	st.ReadOnly = pb.ReadOnly
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ValidateCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage. Only
	// applicable for when purpose is **STORAGE**.
	// Wire name: 'isDir'
	IsDir bool ``
	// The results of the validation check.
	// Wire name: 'results'
	Results         []CredentialValidationResult ``
	ForceSendFields []string                     `tf:"-"`
}

func (s *ValidateCredentialResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidateCredentialResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ValidateCredentialResponseToPb(st *ValidateCredentialResponse) (*catalogpb.ValidateCredentialResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ValidateCredentialResponsePb{}
	pb.IsDir = st.IsDir

	var resultsPb []catalogpb.CredentialValidationResultPb
	for _, item := range st.Results {
		itemPb, err := CredentialValidationResultToPb(&item)
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

func ValidateCredentialResponseFromPb(pb *catalogpb.ValidateCredentialResponsePb) (*ValidateCredentialResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ValidateCredentialResponse{}
	st.IsDir = pb.IsDir

	var resultsField []CredentialValidationResult
	for _, itemPb := range pb.Results {
		item, err := CredentialValidationResultFromPb(&itemPb)
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

// A enum represents the result of the file operation
type ValidateCredentialResult string

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

// Values returns all possible values for ValidateCredentialResult.
//
// There is no guarantee on the order of the values in the slice.
func (f *ValidateCredentialResult) Values() []ValidateCredentialResult {
	return []ValidateCredentialResult{
		ValidateCredentialResultFail,
		ValidateCredentialResultPass,
		ValidateCredentialResultSkip,
	}
}

// Type always returns ValidateCredentialResult to satisfy [pflag.Value] interface
func (f *ValidateCredentialResult) Type() string {
	return "ValidateCredentialResult"
}

func ValidateCredentialResultToPb(st *ValidateCredentialResult) (*catalogpb.ValidateCredentialResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.ValidateCredentialResultPb(*st)
	return &pb, nil
}

func ValidateCredentialResultFromPb(pb *catalogpb.ValidateCredentialResultPb) (*ValidateCredentialResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := ValidateCredentialResult(*pb)
	return &st, nil
}

type ValidateStorageCredential struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRoleRequest ``
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentityRequest ``
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal ``
	// The Cloudflare API token configuration.
	// Wire name: 'cloudflare_api_token'
	CloudflareApiToken *CloudflareApiToken ``
	// The Databricks created GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest ``
	// The name of an existing external location to validate.
	// Wire name: 'external_location_name'
	ExternalLocationName string ``
	// Whether the storage credential is only usable for read operations.
	// Wire name: 'read_only'
	ReadOnly bool ``
	// Required. The name of an existing credential or long-lived cloud
	// credential to validate.
	// Wire name: 'storage_credential_name'
	StorageCredentialName string ``
	// The external location url to validate.
	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ValidateStorageCredential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidateStorageCredential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ValidateStorageCredentialToPb(st *ValidateStorageCredential) (*catalogpb.ValidateStorageCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ValidateStorageCredentialPb{}
	awsIamRolePb, err := AwsIamRoleRequestToPb(st.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRolePb != nil {
		pb.AwsIamRole = awsIamRolePb
	}
	azureManagedIdentityPb, err := AzureManagedIdentityRequestToPb(st.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityPb != nil {
		pb.AzureManagedIdentity = azureManagedIdentityPb
	}
	azureServicePrincipalPb, err := AzureServicePrincipalToPb(st.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalPb != nil {
		pb.AzureServicePrincipal = azureServicePrincipalPb
	}
	cloudflareApiTokenPb, err := CloudflareApiTokenToPb(st.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenPb != nil {
		pb.CloudflareApiToken = cloudflareApiTokenPb
	}
	databricksGcpServiceAccountPb, err := DatabricksGcpServiceAccountRequestToPb(st.DatabricksGcpServiceAccount)
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

func ValidateStorageCredentialFromPb(pb *catalogpb.ValidateStorageCredentialPb) (*ValidateStorageCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ValidateStorageCredential{}
	awsIamRoleField, err := AwsIamRoleRequestFromPb(pb.AwsIamRole)
	if err != nil {
		return nil, err
	}
	if awsIamRoleField != nil {
		st.AwsIamRole = awsIamRoleField
	}
	azureManagedIdentityField, err := AzureManagedIdentityRequestFromPb(pb.AzureManagedIdentity)
	if err != nil {
		return nil, err
	}
	if azureManagedIdentityField != nil {
		st.AzureManagedIdentity = azureManagedIdentityField
	}
	azureServicePrincipalField, err := AzureServicePrincipalFromPb(pb.AzureServicePrincipal)
	if err != nil {
		return nil, err
	}
	if azureServicePrincipalField != nil {
		st.AzureServicePrincipal = azureServicePrincipalField
	}
	cloudflareApiTokenField, err := CloudflareApiTokenFromPb(pb.CloudflareApiToken)
	if err != nil {
		return nil, err
	}
	if cloudflareApiTokenField != nil {
		st.CloudflareApiToken = cloudflareApiTokenField
	}
	databricksGcpServiceAccountField, err := DatabricksGcpServiceAccountRequestFromPb(pb.DatabricksGcpServiceAccount)
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

type ValidateStorageCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage.
	// Wire name: 'isDir'
	IsDir bool ``
	// The results of the validation check.
	// Wire name: 'results'
	Results         []ValidationResult ``
	ForceSendFields []string           `tf:"-"`
}

func (s *ValidateStorageCredentialResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidateStorageCredentialResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ValidateStorageCredentialResponseToPb(st *ValidateStorageCredentialResponse) (*catalogpb.ValidateStorageCredentialResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ValidateStorageCredentialResponsePb{}
	pb.IsDir = st.IsDir

	var resultsPb []catalogpb.ValidationResultPb
	for _, item := range st.Results {
		itemPb, err := ValidationResultToPb(&item)
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

func ValidateStorageCredentialResponseFromPb(pb *catalogpb.ValidateStorageCredentialResponsePb) (*ValidateStorageCredentialResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ValidateStorageCredentialResponse{}
	st.IsDir = pb.IsDir

	var resultsField []ValidationResult
	for _, itemPb := range pb.Results {
		item, err := ValidationResultFromPb(&itemPb)
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

type ValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	// Wire name: 'message'
	Message string ``
	// The operation tested.
	// Wire name: 'operation'
	Operation ValidationResultOperation ``
	// The results of the tested operation.
	// Wire name: 'result'
	Result          ValidationResultResult ``
	ForceSendFields []string               `tf:"-"`
}

func (s *ValidationResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ValidationResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ValidationResultToPb(st *ValidationResult) (*catalogpb.ValidationResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.ValidationResultPb{}
	pb.Message = st.Message
	operationPb, err := ValidationResultOperationToPb(&st.Operation)
	if err != nil {
		return nil, err
	}
	if operationPb != nil {
		pb.Operation = *operationPb
	}
	resultPb, err := ValidationResultResultToPb(&st.Result)
	if err != nil {
		return nil, err
	}
	if resultPb != nil {
		pb.Result = *resultPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ValidationResultFromPb(pb *catalogpb.ValidationResultPb) (*ValidationResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ValidationResult{}
	st.Message = pb.Message
	operationField, err := ValidationResultOperationFromPb(&pb.Operation)
	if err != nil {
		return nil, err
	}
	if operationField != nil {
		st.Operation = *operationField
	}
	resultField, err := ValidationResultResultFromPb(&pb.Result)
	if err != nil {
		return nil, err
	}
	if resultField != nil {
		st.Result = *resultField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// A enum represents the file operation performed on the external location with
// the storage credential
type ValidationResultOperation string

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

// Values returns all possible values for ValidationResultOperation.
//
// There is no guarantee on the order of the values in the slice.
func (f *ValidationResultOperation) Values() []ValidationResultOperation {
	return []ValidationResultOperation{
		ValidationResultOperationDelete,
		ValidationResultOperationList,
		ValidationResultOperationPathExists,
		ValidationResultOperationRead,
		ValidationResultOperationWrite,
	}
}

// Type always returns ValidationResultOperation to satisfy [pflag.Value] interface
func (f *ValidationResultOperation) Type() string {
	return "ValidationResultOperation"
}

func ValidationResultOperationToPb(st *ValidationResultOperation) (*catalogpb.ValidationResultOperationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.ValidationResultOperationPb(*st)
	return &pb, nil
}

func ValidationResultOperationFromPb(pb *catalogpb.ValidationResultOperationPb) (*ValidationResultOperation, error) {
	if pb == nil {
		return nil, nil
	}
	st := ValidationResultOperation(*pb)
	return &st, nil
}

// A enum represents the result of the file operation
type ValidationResultResult string

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

// Values returns all possible values for ValidationResultResult.
//
// There is no guarantee on the order of the values in the slice.
func (f *ValidationResultResult) Values() []ValidationResultResult {
	return []ValidationResultResult{
		ValidationResultResultFail,
		ValidationResultResultPass,
		ValidationResultResultSkip,
	}
}

// Type always returns ValidationResultResult to satisfy [pflag.Value] interface
func (f *ValidationResultResult) Type() string {
	return "ValidationResultResult"
}

func ValidationResultResultToPb(st *ValidationResultResult) (*catalogpb.ValidationResultResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.ValidationResultResultPb(*st)
	return &pb, nil
}

func ValidationResultResultFromPb(pb *catalogpb.ValidationResultResultPb) (*ValidationResultResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := ValidationResultResult(*pb)
	return &st, nil
}

type VolumeInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	// Wire name: 'access_point'
	AccessPoint string ``
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool ``
	// The name of the catalog where the schema and the volume are
	// Wire name: 'catalog_name'
	CatalogName string ``
	// The comment attached to the volume
	// Wire name: 'comment'
	Comment string ``

	// Wire name: 'created_at'
	CreatedAt int64 ``
	// The identifier of the user who created the volume
	// Wire name: 'created_by'
	CreatedBy string ``

	// Wire name: 'encryption_details'
	EncryptionDetails *EncryptionDetails ``
	// The three-level (fully qualified) name of the volume
	// Wire name: 'full_name'
	FullName string ``
	// The unique identifier of the metastore
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// The name of the volume
	// Wire name: 'name'
	Name string ``
	// The identifier of the user who owns the volume
	// Wire name: 'owner'
	Owner string ``
	// The name of the schema where the volume is
	// Wire name: 'schema_name'
	SchemaName string ``
	// The storage location on the cloud
	// Wire name: 'storage_location'
	StorageLocation string ``

	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// The identifier of the user who updated the volume last time
	// Wire name: 'updated_by'
	UpdatedBy string ``
	// The unique identifier of the volume
	// Wire name: 'volume_id'
	VolumeId string ``

	// Wire name: 'volume_type'
	VolumeType      VolumeType ``
	ForceSendFields []string   `tf:"-"`
}

func (s *VolumeInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VolumeInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func VolumeInfoToPb(st *VolumeInfo) (*catalogpb.VolumeInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.VolumeInfoPb{}
	pb.AccessPoint = st.AccessPoint
	pb.BrowseOnly = st.BrowseOnly
	pb.CatalogName = st.CatalogName
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	encryptionDetailsPb, err := EncryptionDetailsToPb(st.EncryptionDetails)
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
	volumeTypePb, err := VolumeTypeToPb(&st.VolumeType)
	if err != nil {
		return nil, err
	}
	if volumeTypePb != nil {
		pb.VolumeType = *volumeTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func VolumeInfoFromPb(pb *catalogpb.VolumeInfoPb) (*VolumeInfo, error) {
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
	encryptionDetailsField, err := EncryptionDetailsFromPb(pb.EncryptionDetails)
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
	volumeTypeField, err := VolumeTypeFromPb(&pb.VolumeType)
	if err != nil {
		return nil, err
	}
	if volumeTypeField != nil {
		st.VolumeType = *volumeTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// The type of the volume. An external volume is located in the specified
// external location. A managed volume is located in the default location which
// is specified by the parent schema, or the parent catalog, or the Metastore.
// [Learn more]
//
// [Learn more]: https://docs.databricks.com/aws/en/volumes/managed-vs-external
type VolumeType string

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

// Values returns all possible values for VolumeType.
//
// There is no guarantee on the order of the values in the slice.
func (f *VolumeType) Values() []VolumeType {
	return []VolumeType{
		VolumeTypeExternal,
		VolumeTypeManaged,
	}
}

// Type always returns VolumeType to satisfy [pflag.Value] interface
func (f *VolumeType) Type() string {
	return "VolumeType"
}

func VolumeTypeToPb(st *VolumeType) (*catalogpb.VolumeTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.VolumeTypePb(*st)
	return &pb, nil
}

func VolumeTypeFromPb(pb *catalogpb.VolumeTypePb) (*VolumeType, error) {
	if pb == nil {
		return nil, nil
	}
	st := VolumeType(*pb)
	return &st, nil
}

type WorkspaceBinding struct {
	// One of READ_WRITE/READ_ONLY. Default is READ_WRITE.
	// Wire name: 'binding_type'
	BindingType WorkspaceBindingBindingType ``
	// Required
	// Wire name: 'workspace_id'
	WorkspaceId int64 ``
}

func WorkspaceBindingToPb(st *WorkspaceBinding) (*catalogpb.WorkspaceBindingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogpb.WorkspaceBindingPb{}
	bindingTypePb, err := WorkspaceBindingBindingTypeToPb(&st.BindingType)
	if err != nil {
		return nil, err
	}
	if bindingTypePb != nil {
		pb.BindingType = *bindingTypePb
	}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func WorkspaceBindingFromPb(pb *catalogpb.WorkspaceBindingPb) (*WorkspaceBinding, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceBinding{}
	bindingTypeField, err := WorkspaceBindingBindingTypeFromPb(&pb.BindingType)
	if err != nil {
		return nil, err
	}
	if bindingTypeField != nil {
		st.BindingType = *bindingTypeField
	}
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

// Using `BINDING_TYPE_` prefix here to avoid conflict with `TableOperation`
// enum in `credentials_common.proto`.
type WorkspaceBindingBindingType string

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

// Values returns all possible values for WorkspaceBindingBindingType.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspaceBindingBindingType) Values() []WorkspaceBindingBindingType {
	return []WorkspaceBindingBindingType{
		WorkspaceBindingBindingTypeBindingTypeReadOnly,
		WorkspaceBindingBindingTypeBindingTypeReadWrite,
	}
}

// Type always returns WorkspaceBindingBindingType to satisfy [pflag.Value] interface
func (f *WorkspaceBindingBindingType) Type() string {
	return "WorkspaceBindingBindingType"
}

func WorkspaceBindingBindingTypeToPb(st *WorkspaceBindingBindingType) (*catalogpb.WorkspaceBindingBindingTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := catalogpb.WorkspaceBindingBindingTypePb(*st)
	return &pb, nil
}

func WorkspaceBindingBindingTypeFromPb(pb *catalogpb.WorkspaceBindingBindingTypePb) (*WorkspaceBindingBindingType, error) {
	if pb == nil {
		return nil, nil
	}
	st := WorkspaceBindingBindingType(*pb)
	return &st, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%.9fs", d.Seconds())
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
