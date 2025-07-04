// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

type AccountsCreateMetastore struct {

	// Wire name: 'metastore_info'
	MetastoreInfo *CreateMetastore `json:"metastore_info,omitempty"`
}

func (st AccountsCreateMetastore) EncodeValues(key string, v *url.Values) error {
	pb, err := accountsCreateMetastoreToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type AccountsCreateMetastoreAssignment struct {

	// Wire name: 'metastore_assignment'
	MetastoreAssignment *CreateMetastoreAssignment `json:"metastore_assignment,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" tf:"-"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`
}

func (st AccountsCreateMetastoreAssignment) EncodeValues(key string, v *url.Values) error {
	pb, err := accountsCreateMetastoreAssignmentToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type AccountsCreateStorageCredential struct {

	// Wire name: 'credential_info'
	CredentialInfo *CreateStorageCredential `json:"credential_info,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" tf:"-"`
}

func (st AccountsCreateStorageCredential) EncodeValues(key string, v *url.Values) error {
	pb, err := accountsCreateStorageCredentialToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type AccountsMetastoreAssignment struct {

	// Wire name: 'metastore_assignment'
	MetastoreAssignment *MetastoreAssignment `json:"metastore_assignment,omitempty"`
}

func (st AccountsMetastoreAssignment) EncodeValues(key string, v *url.Values) error {
	pb, err := accountsMetastoreAssignmentToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type AccountsMetastoreInfo struct {

	// Wire name: 'metastore_info'
	MetastoreInfo *MetastoreInfo `json:"metastore_info,omitempty"`
}

func (st AccountsMetastoreInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := accountsMetastoreInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type AccountsStorageCredentialInfo struct {

	// Wire name: 'credential_info'
	CredentialInfo *StorageCredentialInfo `json:"credential_info,omitempty"`
}

func (st AccountsStorageCredentialInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := accountsStorageCredentialInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type AccountsUpdateMetastore struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" tf:"-"`

	// Wire name: 'metastore_info'
	MetastoreInfo *UpdateMetastore `json:"metastore_info,omitempty"`
}

func (st AccountsUpdateMetastore) EncodeValues(key string, v *url.Values) error {
	pb, err := accountsUpdateMetastoreToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type AccountsUpdateMetastoreAssignment struct {

	// Wire name: 'metastore_assignment'
	MetastoreAssignment *UpdateMetastoreAssignment `json:"metastore_assignment,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" tf:"-"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`
}

func (st AccountsUpdateMetastoreAssignment) EncodeValues(key string, v *url.Values) error {
	pb, err := accountsUpdateMetastoreAssignmentToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type AccountsUpdateStorageCredential struct {

	// Wire name: 'credential_info'
	CredentialInfo *UpdateStorageCredential `json:"credential_info,omitempty"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" tf:"-"`
	// Name of the storage credential.
	StorageCredentialName string `json:"-" tf:"-"`
}

func (st AccountsUpdateStorageCredential) EncodeValues(key string, v *url.Values) error {
	pb, err := accountsUpdateStorageCredentialToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ArtifactAllowlistInfo struct {
	// A list of allowed artifact match patterns.
	// Wire name: 'artifact_matchers'
	ArtifactMatchers []ArtifactMatcher `json:"artifact_matchers,omitempty"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of the user who set the artifact allowlist.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ArtifactAllowlistInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := artifactAllowlistInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ArtifactMatcher struct {
	// The artifact path or maven coordinate
	// Wire name: 'artifact'
	Artifact string `json:"artifact"`
	// The pattern matching type of the artifact
	// Wire name: 'match_type'
	MatchType MatchType `json:"match_type"`
}

func (st ArtifactMatcher) EncodeValues(key string, v *url.Values) error {
	pb, err := artifactMatcherToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// AWS temporary credentials for API authentication. Read more at
// https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html.
type AwsCredentials struct {
	// The access key ID that identifies the temporary credentials.
	// Wire name: 'access_key_id'
	AccessKeyId string `json:"access_key_id,omitempty"`
	// The Amazon Resource Name (ARN) of the S3 access point for temporary
	// credentials related the external location.
	// Wire name: 'access_point'
	AccessPoint string `json:"access_point,omitempty"`
	// The secret access key that can be used to sign AWS API requests.
	// Wire name: 'secret_access_key'
	SecretAccessKey string `json:"secret_access_key,omitempty"`
	// The token that users must pass to AWS API to use the temporary
	// credentials.
	// Wire name: 'session_token'
	SessionToken string `json:"session_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AwsCredentials) EncodeValues(key string, v *url.Values) error {
	pb, err := awsCredentialsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// The AWS IAM role configuration
type AwsIamRole struct {
	// The external ID used in role assumption to prevent the confused deputy
	// problem.
	// Wire name: 'external_id'
	ExternalId string `json:"external_id,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	// Wire name: 'role_arn'
	RoleArn string `json:"role_arn,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	// Wire name: 'unity_catalog_iam_arn'
	UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AwsIamRole) EncodeValues(key string, v *url.Values) error {
	pb, err := awsIamRoleToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// The AWS IAM role configuration
type AwsIamRoleRequest struct {
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	// Wire name: 'role_arn'
	RoleArn string `json:"role_arn"`
}

func (st AwsIamRoleRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := awsIamRoleRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// The AWS IAM role configuration
type AwsIamRoleResponse struct {
	// The external ID used in role assumption to prevent the confused deputy
	// problem.
	// Wire name: 'external_id'
	ExternalId string `json:"external_id,omitempty"`
	// The Amazon Resource Name (ARN) of the AWS IAM role used to vend temporary
	// credentials.
	// Wire name: 'role_arn'
	RoleArn string `json:"role_arn"`
	// The Amazon Resource Name (ARN) of the AWS IAM user managed by Databricks.
	// This is the identity that is going to assume the AWS IAM role.
	// Wire name: 'unity_catalog_iam_arn'
	UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AwsIamRoleResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := awsIamRoleResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type AwsSqsQueue struct {
	// Unique identifier included in the name of file events managed cloud
	// resources.
	// Wire name: 'managed_resource_id'
	ManagedResourceId string `json:"managed_resource_id,omitempty"`
	// The AQS queue url in the format
	// https://sqs.{region}.amazonaws.com/{account id}/{queue name} Required for
	// provided_sqs.
	// Wire name: 'queue_url'
	QueueUrl string `json:"queue_url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AwsSqsQueue) EncodeValues(key string, v *url.Values) error {
	pb, err := awsSqsQueueToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AwsSqsQueue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &awsSqsQueuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := awsSqsQueueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AwsSqsQueue) MarshalJSON() ([]byte, error) {
	pb, err := awsSqsQueueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Azure Active Directory token, essentially the Oauth token for Azure Service
// Principal or Managed Identity. Read more at
// https://learn.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/aad/service-prin-aad-token
type AzureActiveDirectoryToken struct {
	// Opaque token that contains claims that you can use in Azure Active
	// Directory to access cloud services.
	// Wire name: 'aad_token'
	AadToken string `json:"aad_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AzureActiveDirectoryToken) EncodeValues(key string, v *url.Values) error {
	pb, err := azureActiveDirectoryTokenToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// The Azure managed identity configuration.
type AzureManagedIdentity struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}`.
	// Wire name: 'access_connector_id'
	AccessConnectorId string `json:"access_connector_id"`
	// The Databricks internal ID that represents this managed identity.
	// Wire name: 'credential_id'
	CredentialId string `json:"credential_id,omitempty"`
	// The Azure resource ID of the managed identity. Use the format,
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}`
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// using the system-assigned identity.
	// Wire name: 'managed_identity_id'
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AzureManagedIdentity) EncodeValues(key string, v *url.Values) error {
	pb, err := azureManagedIdentityToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// The Azure managed identity configuration.
type AzureManagedIdentityRequest struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}`.
	// Wire name: 'access_connector_id'
	AccessConnectorId string `json:"access_connector_id"`
	// The Azure resource ID of the managed identity. Use the format,
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}`
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// using the system-assigned identity.
	// Wire name: 'managed_identity_id'
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AzureManagedIdentityRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := azureManagedIdentityRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// The Azure managed identity configuration.
type AzureManagedIdentityResponse struct {
	// The Azure resource ID of the Azure Databricks Access Connector. Use the
	// format
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.Databricks/accessConnectors/{connector-name}`.
	// Wire name: 'access_connector_id'
	AccessConnectorId string `json:"access_connector_id"`
	// The Databricks internal ID that represents this managed identity.
	// Wire name: 'credential_id'
	CredentialId string `json:"credential_id,omitempty"`
	// The Azure resource ID of the managed identity. Use the format,
	// `/subscriptions/{guid}/resourceGroups/{rg-name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identity-name}`
	// This is only available for user-assgined identities. For system-assigned
	// identities, the access_connector_id is used to identify the identity. If
	// this field is not provided, then we assume the AzureManagedIdentity is
	// using the system-assigned identity.
	// Wire name: 'managed_identity_id'
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AzureManagedIdentityResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := azureManagedIdentityResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type AzureQueueStorage struct {
	// Unique identifier included in the name of file events managed cloud
	// resources.
	// Wire name: 'managed_resource_id'
	ManagedResourceId string `json:"managed_resource_id,omitempty"`
	// The AQS queue url in the format https://{storage
	// account}.queue.core.windows.net/{queue name} Required for provided_aqs.
	// Wire name: 'queue_url'
	QueueUrl string `json:"queue_url,omitempty"`
	// The resource group for the queue, event grid subscription, and external
	// location storage account. Only required for locations with a service
	// principal storage credential
	// Wire name: 'resource_group'
	ResourceGroup string `json:"resource_group,omitempty"`
	// Optional subscription id for the queue, event grid subscription, and
	// external location storage account. Required for locations with a service
	// principal storage credential
	// Wire name: 'subscription_id'
	SubscriptionId string `json:"subscription_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AzureQueueStorage) EncodeValues(key string, v *url.Values) error {
	pb, err := azureQueueStorageToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AzureQueueStorage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &azureQueueStoragePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := azureQueueStorageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AzureQueueStorage) MarshalJSON() ([]byte, error) {
	pb, err := azureQueueStorageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The Azure service principal configuration. Only applicable when purpose is
// **STORAGE**.
type AzureServicePrincipal struct {
	// The application ID of the application registration within the referenced
	// AAD tenant.
	// Wire name: 'application_id'
	ApplicationId string `json:"application_id"`
	// The client secret generated for the above app ID in AAD.
	// Wire name: 'client_secret'
	ClientSecret string `json:"client_secret"`
	// The directory ID corresponding to the Azure Active Directory (AAD) tenant
	// of the application.
	// Wire name: 'directory_id'
	DirectoryId string `json:"directory_id"`
}

func (st AzureServicePrincipal) EncodeValues(key string, v *url.Values) error {
	pb, err := azureServicePrincipalToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Azure temporary credentials for API authentication. Read more at
// https://docs.microsoft.com/en-us/rest/api/storageservices/create-user-delegation-sas
type AzureUserDelegationSas struct {
	// The signed URI (SAS Token) used to access blob services for a given path
	// Wire name: 'sas_token'
	SasToken string `json:"sas_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AzureUserDelegationSas) EncodeValues(key string, v *url.Values) error {
	pb, err := azureUserDelegationSasToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CancelRefreshRequest struct {
	// ID of the refresh.
	RefreshId string `json:"-" tf:"-"`
	// Full name of the table.
	TableName string `json:"-" tf:"-"`
}

func (st CancelRefreshRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := cancelRefreshRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CatalogInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool `json:"browse_only,omitempty"`

	// Wire name: 'catalog_type'
	CatalogType CatalogType `json:"catalog_type,omitempty"`
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The name of the connection to an external data source.
	// Wire name: 'connection_name'
	ConnectionName string `json:"connection_name,omitempty"`
	// Time at which this catalog was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of catalog creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`

	// Wire name: 'effective_predictive_optimization_flag'
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag `json:"effective_predictive_optimization_flag,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'enable_predictive_optimization'
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// The full name of the catalog. Corresponds with the name field.
	// Wire name: 'full_name'
	FullName string `json:"full_name,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode CatalogIsolationMode `json:"isolation_mode,omitempty"`
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of catalog.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string `json:"options,omitempty"`
	// Username of current owner of catalog.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string `json:"properties,omitempty"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	// Wire name: 'provider_name'
	ProviderName string `json:"provider_name,omitempty"`

	// Wire name: 'provisioning_info'
	ProvisioningInfo *ProvisioningInfo `json:"provisioning_info,omitempty"`

	// Wire name: 'securable_type'
	SecurableType SecurableType `json:"securable_type,omitempty"`
	// The name of the share under the share provider.
	// Wire name: 'share_name'
	ShareName string `json:"share_name,omitempty"`
	// Storage Location URL (full path) for managed tables within catalog.
	// Wire name: 'storage_location'
	StorageLocation string `json:"storage_location,omitempty"`
	// Storage root URL for managed tables within catalog.
	// Wire name: 'storage_root'
	StorageRoot string `json:"storage_root,omitempty"`
	// Time at which this catalog was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified catalog.
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CatalogInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := catalogInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// The Cloudflare API token configuration. Read more at
// https://developers.cloudflare.com/r2/api/s3/tokens/
type CloudflareApiToken struct {
	// The access key ID associated with the API token.
	// Wire name: 'access_key_id'
	AccessKeyId string `json:"access_key_id"`
	// The ID of the account associated with the API token.
	// Wire name: 'account_id'
	AccountId string `json:"account_id"`
	// The secret access token generated for the above access key ID.
	// Wire name: 'secret_access_key'
	SecretAccessKey string `json:"secret_access_key"`
}

func (st CloudflareApiToken) EncodeValues(key string, v *url.Values) error {
	pb, err := cloudflareApiTokenToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ColumnInfo struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`

	// Wire name: 'mask'
	Mask *ColumnMask `json:"mask,omitempty"`
	// Name of Column.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Whether field may be Null (default: true).
	// Wire name: 'nullable'
	Nullable bool `json:"nullable,omitempty"`
	// Partition index for column.
	// Wire name: 'partition_index'
	PartitionIndex int `json:"partition_index,omitempty"`
	// Ordinal position of column (starting at position 0).
	// Wire name: 'position'
	Position int `json:"position,omitempty"`
	// Format of IntervalType.
	// Wire name: 'type_interval_type'
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// Full data type specification, JSON-serialized.
	// Wire name: 'type_json'
	TypeJson string `json:"type_json,omitempty"`

	// Wire name: 'type_name'
	TypeName ColumnTypeName `json:"type_name,omitempty"`
	// Digits of precision; required for DecimalTypes.
	// Wire name: 'type_precision'
	TypePrecision int `json:"type_precision,omitempty"`
	// Digits to right of decimal; Required for DecimalTypes.
	// Wire name: 'type_scale'
	TypeScale int `json:"type_scale,omitempty"`
	// Full data type specification as SQL/catalogString text.
	// Wire name: 'type_text'
	TypeText string `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ColumnInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := columnInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ColumnInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &columnInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := columnInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ColumnInfo) MarshalJSON() ([]byte, error) {
	pb, err := columnInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ColumnMask struct {
	// The full name of the column mask SQL UDF.
	// Wire name: 'function_name'
	FunctionName string `json:"function_name,omitempty"`
	// The list of additional table columns to be passed as input to the column
	// mask function. The first arg of the mask function should be of the type
	// of the column being masked and the types of the rest of the args should
	// match the types of columns in 'using_column_names'.
	// Wire name: 'using_column_names'
	UsingColumnNames []string `json:"using_column_names,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ColumnMask) EncodeValues(key string, v *url.Values) error {
	pb, err := columnMaskToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ColumnRelationship struct {

	// Wire name: 'source'
	Source string `json:"source,omitempty"`

	// Wire name: 'target'
	Target string `json:"target,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ColumnRelationship) EncodeValues(key string, v *url.Values) error {
	pb, err := columnRelationshipToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ColumnRelationship) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &columnRelationshipPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := columnRelationshipFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ColumnRelationship) MarshalJSON() ([]byte, error) {
	pb, err := columnRelationshipToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// A connection that is dependent on a SQL object.
type ConnectionDependency struct {
	// Full name of the dependent connection, in the form of
	// __connection_name__.
	// Wire name: 'connection_name'
	ConnectionName string `json:"connection_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ConnectionDependency) EncodeValues(key string, v *url.Values) error {
	pb, err := connectionDependencyToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ConnectionDependency) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &connectionDependencyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := connectionDependencyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ConnectionDependency) MarshalJSON() ([]byte, error) {
	pb, err := connectionDependencyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ConnectionInfo struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Unique identifier of the Connection.
	// Wire name: 'connection_id'
	ConnectionId string `json:"connection_id,omitempty"`
	// The type of connection.
	// Wire name: 'connection_type'
	ConnectionType ConnectionType `json:"connection_type,omitempty"`
	// Time at which this connection was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of connection creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// The type of credential.
	// Wire name: 'credential_type'
	CredentialType CredentialType `json:"credential_type,omitempty"`
	// Full name of connection.
	// Wire name: 'full_name'
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of the connection.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string `json:"options,omitempty"`
	// Username of current owner of the connection.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string `json:"properties,omitempty"`

	// Wire name: 'provisioning_info'
	ProvisioningInfo *ProvisioningInfo `json:"provisioning_info,omitempty"`
	// If the connection is read only.
	// Wire name: 'read_only'
	ReadOnly bool `json:"read_only,omitempty"`

	// Wire name: 'securable_type'
	SecurableType SecurableType `json:"securable_type,omitempty"`
	// Time at which this connection was updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified connection.
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`
	// URL of the remote data source, extracted from options.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ConnectionInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := connectionInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Next Id: 36
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

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_CONTINUOUS_UPDATE or the ONLINE_UPDATING_PIPELINE_RESOURCES state.
type ContinuousUpdateStatus struct {
	// Progress of the initial data synchronization.
	// Wire name: 'initial_pipeline_sync_progress'
	InitialPipelineSyncProgress *PipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	// Wire name: 'last_processed_commit_version'
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	// Wire name: 'timestamp'
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ContinuousUpdateStatus) EncodeValues(key string, v *url.Values) error {
	pb, err := continuousUpdateStatusToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateCatalog struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The name of the connection to an external data source.
	// Wire name: 'connection_name'
	ConnectionName string `json:"connection_name,omitempty"`
	// Name of catalog.
	// Wire name: 'name'
	Name string `json:"name"`
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string `json:"options,omitempty"`
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string `json:"properties,omitempty"`
	// The name of delta sharing provider.
	//
	// A Delta Sharing catalog is a catalog that is based on a Delta share on a
	// remote sharing server.
	// Wire name: 'provider_name'
	ProviderName string `json:"provider_name,omitempty"`
	// The name of the share under the share provider.
	// Wire name: 'share_name'
	ShareName string `json:"share_name,omitempty"`
	// Storage root URL for managed tables within catalog.
	// Wire name: 'storage_root'
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateCatalog) EncodeValues(key string, v *url.Values) error {
	pb, err := createCatalogToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateConnection struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The type of connection.
	// Wire name: 'connection_type'
	ConnectionType ConnectionType `json:"connection_type"`
	// Name of the connection.
	// Wire name: 'name'
	Name string `json:"name"`
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string `json:"options"`
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string `json:"properties,omitempty"`
	// If the connection is read only.
	// Wire name: 'read_only'
	ReadOnly bool `json:"read_only,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateConnection) EncodeValues(key string, v *url.Values) error {
	pb, err := createConnectionToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateCredentialRequest struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentity `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The Databricks managed GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount `json:"databricks_gcp_service_account,omitempty"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	// Wire name: 'name'
	Name string `json:"name"`
	// Indicates the purpose of the credential.
	// Wire name: 'purpose'
	Purpose CredentialPurpose `json:"purpose,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	// Wire name: 'read_only'
	ReadOnly bool `json:"read_only,omitempty"`
	// Optional. Supplying true to this argument skips validation of the created
	// set of credentials.
	// Wire name: 'skip_validation'
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateCredentialRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createCredentialRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateExternalLineageRelationshipRequest struct {

	// Wire name: 'external_lineage_relationship'
	ExternalLineageRelationship CreateRequestExternalLineage `json:"external_lineage_relationship"`
}

func (st CreateExternalLineageRelationshipRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createExternalLineageRelationshipRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateExternalLineageRelationshipRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createExternalLineageRelationshipRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createExternalLineageRelationshipRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateExternalLineageRelationshipRequest) MarshalJSON() ([]byte, error) {
	pb, err := createExternalLineageRelationshipRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateExternalLocation struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Name of the storage credential used with this location.
	// Wire name: 'credential_name'
	CredentialName string `json:"credential_name"`
	// Whether to enable file events on this external location.
	// Wire name: 'enable_file_events'
	EnableFileEvents bool `json:"enable_file_events,omitempty"`

	// Wire name: 'encryption_details'
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	// Wire name: 'fallback'
	Fallback bool `json:"fallback,omitempty"`
	// File event queue settings.
	// Wire name: 'file_event_queue'
	FileEventQueue *FileEventQueue `json:"file_event_queue,omitempty"`
	// Name of the external location.
	// Wire name: 'name'
	Name string `json:"name"`
	// Indicates whether the external location is read-only.
	// Wire name: 'read_only'
	ReadOnly bool `json:"read_only,omitempty"`
	// Skips validation of the storage credential associated with the external
	// location.
	// Wire name: 'skip_validation'
	SkipValidation bool `json:"skip_validation,omitempty"`
	// Path URL of the external location.
	// Wire name: 'url'
	Url string `json:"url"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateExternalLocation) EncodeValues(key string, v *url.Values) error {
	pb, err := createExternalLocationToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateExternalMetadataRequest struct {

	// Wire name: 'external_metadata'
	ExternalMetadata ExternalMetadata `json:"external_metadata"`
}

func (st CreateExternalMetadataRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createExternalMetadataRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateExternalMetadataRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createExternalMetadataRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createExternalMetadataRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateExternalMetadataRequest) MarshalJSON() ([]byte, error) {
	pb, err := createExternalMetadataRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateFunction struct {
	// Name of parent catalog.
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name"`
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Scalar function return data type.
	// Wire name: 'data_type'
	DataType ColumnTypeName `json:"data_type"`
	// External function language.
	// Wire name: 'external_language'
	ExternalLanguage string `json:"external_language,omitempty"`
	// External function name.
	// Wire name: 'external_name'
	ExternalName string `json:"external_name,omitempty"`
	// Pretty printed function data type.
	// Wire name: 'full_data_type'
	FullDataType string `json:"full_data_type"`

	// Wire name: 'input_params'
	InputParams FunctionParameterInfos `json:"input_params"`
	// Whether the function is deterministic.
	// Wire name: 'is_deterministic'
	IsDeterministic bool `json:"is_deterministic"`
	// Function null call.
	// Wire name: 'is_null_call'
	IsNullCall bool `json:"is_null_call"`
	// Name of function, relative to parent schema.
	// Wire name: 'name'
	Name string `json:"name"`
	// Function parameter style. **S** is the value for SQL.
	// Wire name: 'parameter_style'
	ParameterStyle CreateFunctionParameterStyle `json:"parameter_style"`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	// Wire name: 'properties'
	Properties string `json:"properties,omitempty"`
	// Table function return parameters.
	// Wire name: 'return_params'
	ReturnParams *FunctionParameterInfos `json:"return_params,omitempty"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	// Wire name: 'routine_body'
	RoutineBody CreateFunctionRoutineBody `json:"routine_body"`
	// Function body.
	// Wire name: 'routine_definition'
	RoutineDefinition string `json:"routine_definition"`
	// Function dependencies.
	// Wire name: 'routine_dependencies'
	RoutineDependencies *DependencyList `json:"routine_dependencies,omitempty"`
	// Name of parent schema relative to its parent catalog.
	// Wire name: 'schema_name'
	SchemaName string `json:"schema_name"`
	// Function security type.
	// Wire name: 'security_type'
	SecurityType CreateFunctionSecurityType `json:"security_type"`
	// Specific name of the function; Reserved for future use.
	// Wire name: 'specific_name'
	SpecificName string `json:"specific_name"`
	// Function SQL data access.
	// Wire name: 'sql_data_access'
	SqlDataAccess CreateFunctionSqlDataAccess `json:"sql_data_access"`
	// List of schemes whose objects can be referenced without qualification.
	// Wire name: 'sql_path'
	SqlPath string `json:"sql_path,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateFunction) EncodeValues(key string, v *url.Values) error {
	pb, err := createFunctionToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateFunctionRequest struct {
	// Partial __FunctionInfo__ specifying the function to be created.
	// Wire name: 'function_info'
	FunctionInfo CreateFunction `json:"function_info"`
}

func (st CreateFunctionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createFunctionRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateMetastore struct {
	// The user-specified name of the metastore.
	// Wire name: 'name'
	Name string `json:"name"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	// Wire name: 'region'
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	// Wire name: 'storage_root'
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateMetastore) EncodeValues(key string, v *url.Values) error {
	pb, err := createMetastoreToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateMetastoreAssignment struct {
	// The name of the default catalog in the metastore. This field is
	// deprecated. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	// Wire name: 'default_catalog_name'
	DefaultCatalogName string `json:"default_catalog_name"`
	// The unique ID of the metastore.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`
}

func (st CreateMetastoreAssignment) EncodeValues(key string, v *url.Values) error {
	pb, err := createMetastoreAssignmentToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateMonitor struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	// Wire name: 'assets_dir'
	AssetsDir string `json:"assets_dir"`
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	// Wire name: 'baseline_table_name'
	BaselineTableName string `json:"baseline_table_name,omitempty"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	// Wire name: 'custom_metrics'
	CustomMetrics []MonitorMetric `json:"custom_metrics,omitempty"`
	// The data classification config for the monitor.
	// Wire name: 'data_classification_config'
	DataClassificationConfig *MonitorDataClassificationConfig `json:"data_classification_config,omitempty"`
	// Configuration for monitoring inference logs.
	// Wire name: 'inference_log'
	InferenceLog *MonitorInferenceLog `json:"inference_log,omitempty"`
	// The notification settings for the monitor.
	// Wire name: 'notifications'
	Notifications *MonitorNotifications `json:"notifications,omitempty"`
	// Schema where output metric tables are created.
	// Wire name: 'output_schema_name'
	OutputSchemaName string `json:"output_schema_name"`
	// The schedule for automatically updating and refreshing metric tables.
	// Wire name: 'schedule'
	Schedule *MonitorCronSchedule `json:"schedule,omitempty"`
	// Whether to skip creating a default dashboard summarizing data quality
	// metrics.
	// Wire name: 'skip_builtin_dashboard'
	SkipBuiltinDashboard bool `json:"skip_builtin_dashboard,omitempty"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	// Wire name: 'slicing_exprs'
	SlicingExprs []string `json:"slicing_exprs,omitempty"`
	// Configuration for monitoring snapshot tables.
	// Wire name: 'snapshot'
	Snapshot *MonitorSnapshot `json:"snapshot,omitempty"`
	// Full name of the table.
	TableName string `json:"-" tf:"-"`
	// Configuration for monitoring time series tables.
	// Wire name: 'time_series'
	TimeSeries *MonitorTimeSeries `json:"time_series,omitempty"`
	// Optional argument to specify the warehouse for dashboard creation. If not
	// specified, the first running warehouse will be used.
	// Wire name: 'warehouse_id'
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateMonitor) EncodeValues(key string, v *url.Values) error {
	pb, err := createMonitorToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateOnlineTableRequest struct {
	// Specification of the online table to be created.
	// Wire name: 'table'
	Table OnlineTable `json:"table"`
}

func (st CreateOnlineTableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createOnlineTableRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateRegisteredModelRequest struct {
	// The name of the catalog where the schema and the registered model reside
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name"`
	// The comment attached to the registered model
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The name of the registered model
	// Wire name: 'name'
	Name string `json:"name"`
	// The name of the schema where the registered model resides
	// Wire name: 'schema_name'
	SchemaName string `json:"schema_name"`
	// The storage location on the cloud under which model version data files
	// are stored
	// Wire name: 'storage_location'
	StorageLocation string `json:"storage_location,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateRegisteredModelRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createRegisteredModelRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateRequestExternalLineage struct {
	// List of column relationships between source and target objects.
	// Wire name: 'columns'
	Columns []ColumnRelationship `json:"columns,omitempty"`
	// Unique identifier of the external lineage relationship.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Key-value properties associated with the external lineage relationship.
	// Wire name: 'properties'
	Properties map[string]string `json:"properties,omitempty"`
	// Source object of the external lineage relationship.
	// Wire name: 'source'
	Source ExternalLineageObject `json:"source"`
	// Target object of the external lineage relationship.
	// Wire name: 'target'
	Target ExternalLineageObject `json:"target"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateRequestExternalLineage) EncodeValues(key string, v *url.Values) error {
	pb, err := createRequestExternalLineageToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateRequestExternalLineage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createRequestExternalLineagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createRequestExternalLineageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateRequestExternalLineage) MarshalJSON() ([]byte, error) {
	pb, err := createRequestExternalLineageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateSchema struct {
	// Name of parent catalog.
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name"`
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Name of schema, relative to parent catalog.
	// Wire name: 'name'
	Name string `json:"name"`
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string `json:"properties,omitempty"`
	// Storage root URL for managed tables within schema.
	// Wire name: 'storage_root'
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateSchema) EncodeValues(key string, v *url.Values) error {
	pb, err := createSchemaToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateStorageCredential struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRoleRequest `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentityRequest `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	// Wire name: 'cloudflare_api_token'
	CloudflareApiToken *CloudflareApiToken `json:"cloudflare_api_token,omitempty"`
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The Databricks managed GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `json:"databricks_gcp_service_account,omitempty"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	// Wire name: 'name'
	Name string `json:"name"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	// Wire name: 'read_only'
	ReadOnly bool `json:"read_only,omitempty"`
	// Supplying true to this argument skips validation of the created
	// credential.
	// Wire name: 'skip_validation'
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateStorageCredential) EncodeValues(key string, v *url.Values) error {
	pb, err := createStorageCredentialToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateTableConstraint struct {

	// Wire name: 'constraint'
	Constraint TableConstraint `json:"constraint"`
	// The full name of the table referenced by the constraint.
	// Wire name: 'full_name_arg'
	FullNameArg string `json:"full_name_arg"`
}

func (st CreateTableConstraint) EncodeValues(key string, v *url.Values) error {
	pb, err := createTableConstraintToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateVolumeRequestContent struct {
	// The name of the catalog where the schema and the volume are
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name"`
	// The comment attached to the volume
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The name of the volume
	// Wire name: 'name'
	Name string `json:"name"`
	// The name of the schema where the volume is
	// Wire name: 'schema_name'
	SchemaName string `json:"schema_name"`
	// The storage location on the cloud
	// Wire name: 'storage_location'
	StorageLocation string `json:"storage_location,omitempty"`

	// Wire name: 'volume_type'
	VolumeType VolumeType `json:"volume_type"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateVolumeRequestContent) EncodeValues(key string, v *url.Values) error {
	pb, err := createVolumeRequestContentToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// A credential that is dependent on a SQL object.
type CredentialDependency struct {
	// Full name of the dependent credential, in the form of
	// __credential_name__.
	// Wire name: 'credential_name'
	CredentialName string `json:"credential_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CredentialDependency) EncodeValues(key string, v *url.Values) error {
	pb, err := credentialDependencyToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CredentialDependency) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &credentialDependencyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := credentialDependencyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CredentialDependency) MarshalJSON() ([]byte, error) {
	pb, err := credentialDependencyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CredentialInfo struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentity `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Time at which this credential was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of credential creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// The Databricks managed GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount `json:"databricks_gcp_service_account,omitempty"`
	// The full name of the credential.
	// Wire name: 'full_name'
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the credential.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Unique identifier of the parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Username of current owner of credential.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Indicates the purpose of the credential.
	// Wire name: 'purpose'
	Purpose CredentialPurpose `json:"purpose,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	// Wire name: 'read_only'
	ReadOnly bool `json:"read_only,omitempty"`
	// Time at which this credential was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the credential.
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`
	// Whether this credential is the current metastore's root storage
	// credential. Only applicable when purpose is **STORAGE**.
	// Wire name: 'used_for_managed_storage'
	UsedForManagedStorage bool `json:"used_for_managed_storage,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CredentialInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := credentialInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CredentialValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	// Wire name: 'message'
	Message string `json:"message,omitempty"`
	// The results of the tested operation.
	// Wire name: 'result'
	Result ValidateCredentialResult `json:"result,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CredentialValidationResult) EncodeValues(key string, v *url.Values) error {
	pb, err := credentialValidationResultToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// GCP long-lived credential. Databricks-created Google Cloud Storage service
// account.
type DatabricksGcpServiceAccount struct {
	// The Databricks internal ID that represents this managed identity.
	// Wire name: 'credential_id'
	CredentialId string `json:"credential_id,omitempty"`
	// The email of the service account.
	// Wire name: 'email'
	Email string `json:"email,omitempty"`
	// The ID that represents the private key for this Service Account
	// Wire name: 'private_key_id'
	PrivateKeyId string `json:"private_key_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabricksGcpServiceAccount) EncodeValues(key string, v *url.Values) error {
	pb, err := databricksGcpServiceAccountToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// GCP long-lived credential. Databricks-created Google Cloud Storage service
// account.
type DatabricksGcpServiceAccountRequest struct {
}

func (st DatabricksGcpServiceAccountRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := databricksGcpServiceAccountRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// GCP long-lived credential. Databricks-created Google Cloud Storage service
// account.
type DatabricksGcpServiceAccountResponse struct {
	// The Databricks internal ID that represents this managed identity.
	// Wire name: 'credential_id'
	CredentialId string `json:"credential_id,omitempty"`
	// The email of the service account.
	// Wire name: 'email'
	Email string `json:"email,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabricksGcpServiceAccountResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := databricksGcpServiceAccountResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteAccountMetastoreAssignmentRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" tf:"-"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`
}

func (st DeleteAccountMetastoreAssignmentRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteAccountMetastoreAssignmentRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteAccountMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force bool `json:"-" tf:"-"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteAccountMetastoreRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteAccountMetastoreRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteAccountStorageCredentialRequest struct {
	// Force deletion even if the Storage Credential is not empty. Default is
	// false.
	Force bool `json:"-" tf:"-"`
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" tf:"-"`
	// Name of the storage credential.
	StorageCredentialName string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteAccountStorageCredentialRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteAccountStorageCredentialRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteAliasRequest struct {
	// The name of the alias
	Alias string `json:"-" tf:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" tf:"-"`
}

func (st DeleteAliasRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteAliasRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteCatalogRequest struct {
	// Force deletion even if the catalog is not empty.
	Force bool `json:"-" tf:"-"`
	// The name of the catalog.
	Name string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteCatalogRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteCatalogRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteConnectionRequest struct {
	// The name of the connection to be deleted.
	Name string `json:"-" tf:"-"`
}

func (st DeleteConnectionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteConnectionRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteCredentialRequest struct {
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	Force bool `json:"-" tf:"-"`
	// Name of the credential.
	NameArg string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteCredentialRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteCredentialRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteExternalLineageRelationshipRequest struct {
	ExternalLineageRelationship DeleteRequestExternalLineage `json:"-" tf:"-"`
}

func (st DeleteExternalLineageRelationshipRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteExternalLineageRelationshipRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteExternalLineageRelationshipRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteExternalLineageRelationshipRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteExternalLineageRelationshipRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteExternalLineageRelationshipRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteExternalLineageRelationshipRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteExternalLineageRelationshipResponse struct {
}

func (st DeleteExternalLineageRelationshipResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteExternalLineageRelationshipResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteExternalLineageRelationshipResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteExternalLineageRelationshipResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteExternalLineageRelationshipResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteExternalLineageRelationshipResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteExternalLineageRelationshipResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteExternalLocationRequest struct {
	// Force deletion even if there are dependent external tables or mounts.
	Force bool `json:"-" tf:"-"`
	// Name of the external location.
	Name string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteExternalLocationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteExternalLocationRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteExternalMetadataRequest struct {
	Name string `json:"-" tf:"-"`
}

func (st DeleteExternalMetadataRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteExternalMetadataRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteExternalMetadataRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteExternalMetadataRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteExternalMetadataRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteExternalMetadataRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteExternalMetadataRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteExternalMetadataResponse struct {
}

func (st DeleteExternalMetadataResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteExternalMetadataResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteExternalMetadataResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteExternalMetadataResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteExternalMetadataResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteExternalMetadataResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteExternalMetadataResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteFunctionRequest struct {
	// Force deletion even if the function is notempty.
	Force bool `json:"-" tf:"-"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteFunctionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteFunctionRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteMetastoreRequest struct {
	// Force deletion even if the metastore is not empty. Default is false.
	Force bool `json:"-" tf:"-"`
	// Unique ID of the metastore.
	Id string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteMetastoreRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteMetastoreRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteModelVersionRequest struct {
	// The three-level (fully qualified) name of the model version
	FullName string `json:"-" tf:"-"`
	// The integer version number of the model version
	Version int `json:"-" tf:"-"`
}

func (st DeleteModelVersionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteModelVersionRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"-" tf:"-"`
}

func (st DeleteOnlineTableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteOnlineTableRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteQualityMonitorRequest struct {
	// Full name of the table.
	TableName string `json:"-" tf:"-"`
}

func (st DeleteQualityMonitorRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteQualityMonitorRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" tf:"-"`
}

func (st DeleteRegisteredModelRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteRegisteredModelRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteRequestExternalLineage struct {
	// Unique identifier of the external lineage relationship.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Source object of the external lineage relationship.
	// Wire name: 'source'
	Source ExternalLineageObject `json:"source"`
	// Target object of the external lineage relationship.
	// Wire name: 'target'
	Target ExternalLineageObject `json:"target"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteRequestExternalLineage) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteRequestExternalLineageToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteRequestExternalLineage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRequestExternalLineagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRequestExternalLineageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRequestExternalLineage) MarshalJSON() ([]byte, error) {
	pb, err := deleteRequestExternalLineageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteSchemaRequest struct {
	// Force deletion even if the schema is not empty.
	Force bool `json:"-" tf:"-"`
	// Full name of the schema.
	FullName string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteSchemaRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteSchemaRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteStorageCredentialRequest struct {
	// Force an update even if there are dependent external locations or
	// external tables (when purpose is **STORAGE**) or dependent services (when
	// purpose is **SERVICE**).
	Force bool `json:"-" tf:"-"`
	// Name of the storage credential.
	Name string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteStorageCredentialRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteStorageCredentialRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteTableConstraintRequest struct {
	// If true, try deleting all child constraints of the current constraint. If
	// false, reject this operation if the current constraint has any child
	// constraints.
	Cascade bool `json:"-" tf:"-"`
	// The name of the constraint to delete.
	ConstraintName string `json:"-" tf:"-"`
	// Full name of the table referenced by the constraint.
	FullName string `json:"-" tf:"-"`
}

func (st DeleteTableConstraintRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteTableConstraintRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteTableRequest struct {
	// Full name of the table.
	FullName string `json:"-" tf:"-"`
}

func (st DeleteTableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteTableRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteVolumeRequest struct {
	// The three-level (fully qualified) name of the volume
	Name string `json:"-" tf:"-"`
}

func (st DeleteVolumeRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteVolumeRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Properties pertaining to the current state of the delta table as given by the
// commit server. This does not contain **delta.*** (input) properties in
// __TableInfo.properties__.
type DeltaRuntimePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	// Wire name: 'delta_runtime_properties'
	DeltaRuntimeProperties map[string]string `json:"delta_runtime_properties"`
}

func (st DeltaRuntimePropertiesKvPairs) EncodeValues(key string, v *url.Values) error {
	pb, err := deltaRuntimePropertiesKvPairsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// A dependency of a SQL object. One of the following fields must be defined:
// __table__, __function__, __connection__, or __credential__.
type Dependency struct {

	// Wire name: 'connection'
	Connection *ConnectionDependency `json:"connection,omitempty"`

	// Wire name: 'credential'
	Credential *CredentialDependency `json:"credential,omitempty"`

	// Wire name: 'function'
	Function *FunctionDependency `json:"function,omitempty"`

	// Wire name: 'table'
	Table *TableDependency `json:"table,omitempty"`
}

func (st Dependency) EncodeValues(key string, v *url.Values) error {
	pb, err := dependencyToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// A list of dependencies.
type DependencyList struct {
	// Array of dependencies.
	// Wire name: 'dependencies'
	Dependencies []Dependency `json:"dependencies,omitempty"`
}

func (st DependencyList) EncodeValues(key string, v *url.Values) error {
	pb, err := dependencyListToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DisableRequest struct {
	// The metastore ID under which the system schema lives.
	MetastoreId string `json:"-" tf:"-"`
	// Full name of the system schema.
	SchemaName string `json:"-" tf:"-"`
}

func (st DisableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := disableRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type EffectivePermissionsList struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// The privileges conveyed to each principal (either directly or via
	// inheritance)
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []EffectivePrivilegeAssignment `json:"privilege_assignments,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EffectivePermissionsList) EncodeValues(key string, v *url.Values) error {
	pb, err := effectivePermissionsListToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type EffectivePredictiveOptimizationFlag struct {
	// The name of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	// Wire name: 'inherited_from_name'
	InheritedFromName string `json:"inherited_from_name,omitempty"`
	// The type of the object from which the flag was inherited. If there was no
	// inheritance, this field is left blank.
	// Wire name: 'inherited_from_type'
	InheritedFromType EffectivePredictiveOptimizationFlagInheritedFromType `json:"inherited_from_type,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'value'
	Value EnablePredictiveOptimization `json:"value"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EffectivePredictiveOptimizationFlag) EncodeValues(key string, v *url.Values) error {
	pb, err := effectivePredictiveOptimizationFlagToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type EffectivePrivilege struct {
	// The full name of the object that conveys this privilege via inheritance.
	// This field is omitted when privilege is not inherited (it's assigned to
	// the securable itself).
	// Wire name: 'inherited_from_name'
	InheritedFromName string `json:"inherited_from_name,omitempty"`
	// The type of the object that conveys this privilege via inheritance. This
	// field is omitted when privilege is not inherited (it's assigned to the
	// securable itself).
	// Wire name: 'inherited_from_type'
	InheritedFromType SecurableType `json:"inherited_from_type,omitempty"`
	// The privilege assigned to the principal.
	// Wire name: 'privilege'
	Privilege Privilege `json:"privilege,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EffectivePrivilege) EncodeValues(key string, v *url.Values) error {
	pb, err := effectivePrivilegeToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type EffectivePrivilegeAssignment struct {
	// The principal (user email address or group name).
	// Wire name: 'principal'
	Principal string `json:"principal,omitempty"`
	// The privileges conveyed to the principal (either directly or via
	// inheritance).
	// Wire name: 'privileges'
	Privileges []EffectivePrivilege `json:"privileges,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EffectivePrivilegeAssignment) EncodeValues(key string, v *url.Values) error {
	pb, err := effectivePrivilegeAssignmentToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type EnableRequest struct {
	// the catalog for which the system schema is to enabled in
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name,omitempty"`
	// The metastore ID under which the system schema lives.
	MetastoreId string `json:"-" tf:"-"`
	// Full name of the system schema.
	SchemaName string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EnableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := enableRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Encryption options that apply to clients connecting to cloud storage.
type EncryptionDetails struct {
	// Server-Side Encryption properties for clients communicating with AWS s3.
	// Wire name: 'sse_encryption_details'
	SseEncryptionDetails *SseEncryptionDetails `json:"sse_encryption_details,omitempty"`
}

func (st EncryptionDetails) EncodeValues(key string, v *url.Values) error {
	pb, err := encryptionDetailsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ExistsRequest struct {
	// Full name of the table.
	FullName string `json:"-" tf:"-"`
}

func (st ExistsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := existsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ExternalLineageExternalMetadata struct {

	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalLineageExternalMetadata) EncodeValues(key string, v *url.Values) error {
	pb, err := externalLineageExternalMetadataToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalLineageExternalMetadata) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalLineageExternalMetadataPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalLineageExternalMetadataFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalLineageExternalMetadata) MarshalJSON() ([]byte, error) {
	pb, err := externalLineageExternalMetadataToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Represents the external metadata object in the lineage event.
type ExternalLineageExternalMetadataInfo struct {
	// Type of entity represented by the external metadata object.
	// Wire name: 'entity_type'
	EntityType string `json:"entity_type,omitempty"`
	// Timestamp of the lineage event.
	// Wire name: 'event_time'
	EventTime string `json:"event_time,omitempty"`
	// Name of the external metadata object.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Type of external system.
	// Wire name: 'system_type'
	SystemType SystemType `json:"system_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalLineageExternalMetadataInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := externalLineageExternalMetadataInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalLineageExternalMetadataInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalLineageExternalMetadataInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalLineageExternalMetadataInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalLineageExternalMetadataInfo) MarshalJSON() ([]byte, error) {
	pb, err := externalLineageExternalMetadataInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Represents the path information in the lineage event.
type ExternalLineageFileInfo struct {
	// Timestamp of the lineage event.
	// Wire name: 'event_time'
	EventTime string `json:"event_time,omitempty"`
	// URL of the path.
	// Wire name: 'path'
	Path string `json:"path,omitempty"`
	// The full name of the securable on the path.
	// Wire name: 'securable_name'
	SecurableName string `json:"securable_name,omitempty"`
	// The securable type of the securable on the path.
	// Wire name: 'securable_type'
	SecurableType string `json:"securable_type,omitempty"`
	// The storage location associated with securable on the path.
	// Wire name: 'storage_location'
	StorageLocation string `json:"storage_location,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalLineageFileInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := externalLineageFileInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalLineageFileInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalLineageFileInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalLineageFileInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalLineageFileInfo) MarshalJSON() ([]byte, error) {
	pb, err := externalLineageFileInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Lineage response containing lineage information of a data asset.
type ExternalLineageInfo struct {
	// Information about the edge metadata of the external lineage relationship.
	// Wire name: 'external_lineage_info'
	ExternalLineageInfo *ExternalLineageRelationshipInfo `json:"external_lineage_info,omitempty"`
	// Information about external metadata involved in the lineage relationship.
	// Wire name: 'external_metadata_info'
	ExternalMetadataInfo *ExternalLineageExternalMetadataInfo `json:"external_metadata_info,omitempty"`
	// Information about the file involved in the lineage relationship.
	// Wire name: 'file_info'
	FileInfo *ExternalLineageFileInfo `json:"file_info,omitempty"`
	// Information about the model version involved in the lineage relationship.
	// Wire name: 'model_info'
	ModelInfo *ExternalLineageModelVersionInfo `json:"model_info,omitempty"`
	// Information about the table involved in the lineage relationship.
	// Wire name: 'table_info'
	TableInfo *ExternalLineageTableInfo `json:"table_info,omitempty"`
}

func (st ExternalLineageInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := externalLineageInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalLineageInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalLineageInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalLineageInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalLineageInfo) MarshalJSON() ([]byte, error) {
	pb, err := externalLineageInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExternalLineageModelVersion struct {

	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	// Wire name: 'version'
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalLineageModelVersion) EncodeValues(key string, v *url.Values) error {
	pb, err := externalLineageModelVersionToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalLineageModelVersion) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalLineageModelVersionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalLineageModelVersionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalLineageModelVersion) MarshalJSON() ([]byte, error) {
	pb, err := externalLineageModelVersionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Represents the model version information in the lineage event.
type ExternalLineageModelVersionInfo struct {
	// Timestamp of the lineage event.
	// Wire name: 'event_time'
	EventTime string `json:"event_time,omitempty"`
	// Name of the model.
	// Wire name: 'model_name'
	ModelName string `json:"model_name,omitempty"`
	// Version number of the model.
	// Wire name: 'version'
	Version int64 `json:"version,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalLineageModelVersionInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := externalLineageModelVersionInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalLineageModelVersionInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalLineageModelVersionInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalLineageModelVersionInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalLineageModelVersionInfo) MarshalJSON() ([]byte, error) {
	pb, err := externalLineageModelVersionInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExternalLineageObject struct {

	// Wire name: 'external_metadata'
	ExternalMetadata *ExternalLineageExternalMetadata `json:"external_metadata,omitempty"`

	// Wire name: 'model_version'
	ModelVersion *ExternalLineageModelVersion `json:"model_version,omitempty"`

	// Wire name: 'path'
	Path *ExternalLineagePath `json:"path,omitempty"`

	// Wire name: 'table'
	Table *ExternalLineageTable `json:"table,omitempty"`
}

func (st ExternalLineageObject) EncodeValues(key string, v *url.Values) error {
	pb, err := externalLineageObjectToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalLineageObject) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalLineageObjectPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalLineageObjectFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalLineageObject) MarshalJSON() ([]byte, error) {
	pb, err := externalLineageObjectToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExternalLineagePath struct {

	// Wire name: 'url'
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalLineagePath) EncodeValues(key string, v *url.Values) error {
	pb, err := externalLineagePathToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalLineagePath) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalLineagePathPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalLineagePathFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalLineagePath) MarshalJSON() ([]byte, error) {
	pb, err := externalLineagePathToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExternalLineageRelationship struct {
	// List of column relationships between source and target objects.
	// Wire name: 'columns'
	Columns []ColumnRelationship `json:"columns,omitempty"`
	// Unique identifier of the external lineage relationship.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Key-value properties associated with the external lineage relationship.
	// Wire name: 'properties'
	Properties map[string]string `json:"properties,omitempty"`
	// Source object of the external lineage relationship.
	// Wire name: 'source'
	Source ExternalLineageObject `json:"source"`
	// Target object of the external lineage relationship.
	// Wire name: 'target'
	Target ExternalLineageObject `json:"target"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalLineageRelationship) EncodeValues(key string, v *url.Values) error {
	pb, err := externalLineageRelationshipToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalLineageRelationship) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalLineageRelationshipPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalLineageRelationshipFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalLineageRelationship) MarshalJSON() ([]byte, error) {
	pb, err := externalLineageRelationshipToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExternalLineageRelationshipInfo struct {
	// List of column relationships between source and target objects.
	// Wire name: 'columns'
	Columns []ColumnRelationship `json:"columns,omitempty"`
	// Unique identifier of the external lineage relationship.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Key-value properties associated with the external lineage relationship.
	// Wire name: 'properties'
	Properties map[string]string `json:"properties,omitempty"`
	// Source object of the external lineage relationship.
	// Wire name: 'source'
	Source ExternalLineageObject `json:"source"`
	// Target object of the external lineage relationship.
	// Wire name: 'target'
	Target ExternalLineageObject `json:"target"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalLineageRelationshipInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := externalLineageRelationshipInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalLineageRelationshipInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalLineageRelationshipInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalLineageRelationshipInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalLineageRelationshipInfo) MarshalJSON() ([]byte, error) {
	pb, err := externalLineageRelationshipInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExternalLineageTable struct {

	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalLineageTable) EncodeValues(key string, v *url.Values) error {
	pb, err := externalLineageTableToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalLineageTable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalLineageTablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalLineageTableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalLineageTable) MarshalJSON() ([]byte, error) {
	pb, err := externalLineageTableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Represents the table information in the lineage event.
type ExternalLineageTableInfo struct {
	// Name of Catalog.
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name,omitempty"`
	// Timestamp of the lineage event.
	// Wire name: 'event_time'
	EventTime string `json:"event_time,omitempty"`
	// Name of Table.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Name of Schema.
	// Wire name: 'schema_name'
	SchemaName string `json:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalLineageTableInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := externalLineageTableInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalLineageTableInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalLineageTableInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalLineageTableInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalLineageTableInfo) MarshalJSON() ([]byte, error) {
	pb, err := externalLineageTableInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExternalLocationInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool `json:"browse_only,omitempty"`
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Time at which this external location was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of external location creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// Unique ID of the location's storage credential.
	// Wire name: 'credential_id'
	CredentialId string `json:"credential_id,omitempty"`
	// Name of the storage credential used with this location.
	// Wire name: 'credential_name'
	CredentialName string `json:"credential_name,omitempty"`
	// Whether to enable file events on this external location.
	// Wire name: 'enable_file_events'
	EnableFileEvents bool `json:"enable_file_events,omitempty"`

	// Wire name: 'encryption_details'
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	// Wire name: 'fallback'
	Fallback bool `json:"fallback,omitempty"`
	// File event queue settings.
	// Wire name: 'file_event_queue'
	FileEventQueue *FileEventQueue `json:"file_event_queue,omitempty"`

	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Unique identifier of metastore hosting the external location.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of the external location.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The owner of the external location.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Indicates whether the external location is read-only.
	// Wire name: 'read_only'
	ReadOnly bool `json:"read_only,omitempty"`
	// Time at which external location this was last modified, in epoch
	// milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the external location.
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`
	// Path URL of the external location.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalLocationInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := externalLocationInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ExternalMetadata struct {
	// List of columns associated with the external metadata object.
	// Wire name: 'columns'
	Columns []string `json:"columns,omitempty"`
	// Time at which this external metadata object was created.
	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"`
	// Username of external metadata object creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// User-provided free-form text description.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Type of entity within the external system.
	// Wire name: 'entity_type'
	EntityType string `json:"entity_type"`
	// Unique identifier of the external metadata object.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of the external metadata object.
	// Wire name: 'name'
	Name string `json:"name"`
	// Owner of the external metadata object.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the external metadata object.
	// Wire name: 'properties'
	Properties map[string]string `json:"properties,omitempty"`
	// Type of external system.
	// Wire name: 'system_type'
	SystemType SystemType `json:"system_type"`
	// Time at which this external metadata object was last modified.
	// Wire name: 'update_time'
	UpdateTime string `json:"update_time,omitempty"`
	// Username of user who last modified external metadata object.
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`
	// URL associated with the external metadata object.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalMetadata) EncodeValues(key string, v *url.Values) error {
	pb, err := externalMetadataToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalMetadata) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalMetadataPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalMetadataFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalMetadata) MarshalJSON() ([]byte, error) {
	pb, err := externalMetadataToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Detailed status of an online table. Shown if the online table is in the
// OFFLINE_FAILED or the ONLINE_PIPELINE_FAILED state.
type FailedStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may only be partially synced to the online
	// table. Only populated if the table is still online and available for
	// serving.
	// Wire name: 'last_processed_commit_version'
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table. Only populated if the table is still online
	// and available for serving.
	// Wire name: 'timestamp'
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st FailedStatus) EncodeValues(key string, v *url.Values) error {
	pb, err := failedStatusToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type FileEventQueue struct {

	// Wire name: 'managed_aqs'
	ManagedAqs *AzureQueueStorage `json:"managed_aqs,omitempty"`

	// Wire name: 'managed_pubsub'
	ManagedPubsub *GcpPubsub `json:"managed_pubsub,omitempty"`

	// Wire name: 'managed_sqs'
	ManagedSqs *AwsSqsQueue `json:"managed_sqs,omitempty"`

	// Wire name: 'provided_aqs'
	ProvidedAqs *AzureQueueStorage `json:"provided_aqs,omitempty"`

	// Wire name: 'provided_pubsub'
	ProvidedPubsub *GcpPubsub `json:"provided_pubsub,omitempty"`

	// Wire name: 'provided_sqs'
	ProvidedSqs *AwsSqsQueue `json:"provided_sqs,omitempty"`
}

func (st FileEventQueue) EncodeValues(key string, v *url.Values) error {
	pb, err := fileEventQueueToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *FileEventQueue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &fileEventQueuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := fileEventQueueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FileEventQueue) MarshalJSON() ([]byte, error) {
	pb, err := fileEventQueueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ForeignKeyConstraint struct {
	// Column names for this constraint.
	// Wire name: 'child_columns'
	ChildColumns []string `json:"child_columns"`
	// The name of the constraint.
	// Wire name: 'name'
	Name string `json:"name"`
	// Column names for this constraint.
	// Wire name: 'parent_columns'
	ParentColumns []string `json:"parent_columns"`
	// The full name of the parent constraint.
	// Wire name: 'parent_table'
	ParentTable string `json:"parent_table"`
	// True if the constraint is RELY, false or unset if NORELY.
	// Wire name: 'rely'
	Rely bool `json:"rely,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ForeignKeyConstraint) EncodeValues(key string, v *url.Values) error {
	pb, err := foreignKeyConstraintToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// A function that is dependent on a SQL object.
type FunctionDependency struct {
	// Full name of the dependent function, in the form of
	// __catalog_name__.__schema_name__.__function_name__.
	// Wire name: 'function_full_name'
	FunctionFullName string `json:"function_full_name"`
}

func (st FunctionDependency) EncodeValues(key string, v *url.Values) error {
	pb, err := functionDependencyToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type FunctionInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool `json:"browse_only,omitempty"`
	// Name of parent catalog.
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name,omitempty"`
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Time at which this function was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of function creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// Scalar function return data type.
	// Wire name: 'data_type'
	DataType ColumnTypeName `json:"data_type,omitempty"`
	// External function language.
	// Wire name: 'external_language'
	ExternalLanguage string `json:"external_language,omitempty"`
	// External function name.
	// Wire name: 'external_name'
	ExternalName string `json:"external_name,omitempty"`
	// Pretty printed function data type.
	// Wire name: 'full_data_type'
	FullDataType string `json:"full_data_type,omitempty"`
	// Full name of function, in form of
	// __catalog_name__.__schema_name__.__function__name__
	// Wire name: 'full_name'
	FullName string `json:"full_name,omitempty"`
	// Id of Function, relative to parent schema.
	// Wire name: 'function_id'
	FunctionId string `json:"function_id,omitempty"`

	// Wire name: 'input_params'
	InputParams *FunctionParameterInfos `json:"input_params,omitempty"`
	// Whether the function is deterministic.
	// Wire name: 'is_deterministic'
	IsDeterministic bool `json:"is_deterministic,omitempty"`
	// Function null call.
	// Wire name: 'is_null_call'
	IsNullCall bool `json:"is_null_call,omitempty"`
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of function, relative to parent schema.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Username of current owner of function.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Function parameter style. **S** is the value for SQL.
	// Wire name: 'parameter_style'
	ParameterStyle FunctionInfoParameterStyle `json:"parameter_style,omitempty"`
	// JSON-serialized key-value pair map, encoded (escaped) as a string.
	// Wire name: 'properties'
	Properties string `json:"properties,omitempty"`
	// Table function return parameters.
	// Wire name: 'return_params'
	ReturnParams *FunctionParameterInfos `json:"return_params,omitempty"`
	// Function language. When **EXTERNAL** is used, the language of the routine
	// function should be specified in the __external_language__ field, and the
	// __return_params__ of the function cannot be used (as **TABLE** return
	// type is not supported), and the __sql_data_access__ field must be
	// **NO_SQL**.
	// Wire name: 'routine_body'
	RoutineBody FunctionInfoRoutineBody `json:"routine_body,omitempty"`
	// Function body.
	// Wire name: 'routine_definition'
	RoutineDefinition string `json:"routine_definition,omitempty"`
	// Function dependencies.
	// Wire name: 'routine_dependencies'
	RoutineDependencies *DependencyList `json:"routine_dependencies,omitempty"`
	// Name of parent schema relative to its parent catalog.
	// Wire name: 'schema_name'
	SchemaName string `json:"schema_name,omitempty"`
	// Function security type.
	// Wire name: 'security_type'
	SecurityType FunctionInfoSecurityType `json:"security_type,omitempty"`
	// Specific name of the function; Reserved for future use.
	// Wire name: 'specific_name'
	SpecificName string `json:"specific_name,omitempty"`
	// Function SQL data access.
	// Wire name: 'sql_data_access'
	SqlDataAccess FunctionInfoSqlDataAccess `json:"sql_data_access,omitempty"`
	// List of schemes whose objects can be referenced without qualification.
	// Wire name: 'sql_path'
	SqlPath string `json:"sql_path,omitempty"`
	// Time at which this function was created, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified function.
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st FunctionInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := functionInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type FunctionParameterInfo struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Name of parameter.
	// Wire name: 'name'
	Name string `json:"name"`
	// Default value of the parameter.
	// Wire name: 'parameter_default'
	ParameterDefault string `json:"parameter_default,omitempty"`

	// Wire name: 'parameter_mode'
	ParameterMode FunctionParameterMode `json:"parameter_mode,omitempty"`

	// Wire name: 'parameter_type'
	ParameterType FunctionParameterType `json:"parameter_type,omitempty"`
	// Ordinal position of column (starting at position 0).
	// Wire name: 'position'
	Position int `json:"position"`
	// Format of IntervalType.
	// Wire name: 'type_interval_type'
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// Full data type spec, JSON-serialized.
	// Wire name: 'type_json'
	TypeJson string `json:"type_json,omitempty"`

	// Wire name: 'type_name'
	TypeName ColumnTypeName `json:"type_name"`
	// Digits of precision; required on Create for DecimalTypes.
	// Wire name: 'type_precision'
	TypePrecision int `json:"type_precision,omitempty"`
	// Digits to right of decimal; Required on Create for DecimalTypes.
	// Wire name: 'type_scale'
	TypeScale int `json:"type_scale,omitempty"`
	// Full data type spec, SQL/catalogString text.
	// Wire name: 'type_text'
	TypeText string `json:"type_text"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st FunctionParameterInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := functionParameterInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type FunctionParameterInfos struct {
	// The array of __FunctionParameterInfo__ definitions of the function's
	// parameters.
	// Wire name: 'parameters'
	Parameters []FunctionParameterInfo `json:"parameters,omitempty"`
}

func (st FunctionParameterInfos) EncodeValues(key string, v *url.Values) error {
	pb, err := functionParameterInfosToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// GCP temporary credentials for API authentication. Read more at
// https://developers.google.com/identity/protocols/oauth2/service-account
type GcpOauthToken struct {

	// Wire name: 'oauth_token'
	OauthToken string `json:"oauth_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GcpOauthToken) EncodeValues(key string, v *url.Values) error {
	pb, err := gcpOauthTokenToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GcpPubsub struct {
	// Unique identifier included in the name of file events managed cloud
	// resources.
	// Wire name: 'managed_resource_id'
	ManagedResourceId string `json:"managed_resource_id,omitempty"`
	// The Pub/Sub subscription name in the format
	// projects/{project}/subscriptions/{subscription name} Required for
	// provided_pubsub.
	// Wire name: 'subscription_name'
	SubscriptionName string `json:"subscription_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GcpPubsub) EncodeValues(key string, v *url.Values) error {
	pb, err := gcpPubsubToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GcpPubsub) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gcpPubsubPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gcpPubsubFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GcpPubsub) MarshalJSON() ([]byte, error) {
	pb, err := gcpPubsubToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The Azure cloud options to customize the requested temporary credential
type GenerateTemporaryServiceCredentialAzureOptions struct {
	// The resources to which the temporary Azure credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://learn.microsoft.com/python/api/azure-core/azure.core.credentials.tokencredential?view=azure-python)
	// Wire name: 'resources'
	Resources []string `json:"resources,omitempty"`
}

func (st GenerateTemporaryServiceCredentialAzureOptions) EncodeValues(key string, v *url.Values) error {
	pb, err := generateTemporaryServiceCredentialAzureOptionsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// The GCP cloud options to customize the requested temporary credential
type GenerateTemporaryServiceCredentialGcpOptions struct {
	// The scopes to which the temporary GCP credential should apply. These
	// resources are the scopes that are passed to the token provider (see
	// https://google-auth.readthedocs.io/en/latest/reference/google.auth.html#google.auth.credentials.Credentials)
	// Wire name: 'scopes'
	Scopes []string `json:"scopes,omitempty"`
}

func (st GenerateTemporaryServiceCredentialGcpOptions) EncodeValues(key string, v *url.Values) error {
	pb, err := generateTemporaryServiceCredentialGcpOptionsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GenerateTemporaryServiceCredentialRequest struct {

	// Wire name: 'azure_options'
	AzureOptions *GenerateTemporaryServiceCredentialAzureOptions `json:"azure_options,omitempty"`
	// The name of the service credential used to generate a temporary
	// credential
	// Wire name: 'credential_name'
	CredentialName string `json:"credential_name"`

	// Wire name: 'gcp_options'
	GcpOptions *GenerateTemporaryServiceCredentialGcpOptions `json:"gcp_options,omitempty"`
}

func (st GenerateTemporaryServiceCredentialRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := generateTemporaryServiceCredentialRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GenerateTemporaryTableCredentialRequest struct {
	// The operation performed against the table data, either READ or
	// READ_WRITE. If READ_WRITE is specified, the credentials returned will
	// have write permissions, otherwise, it will be read only.
	// Wire name: 'operation'
	Operation TableOperation `json:"operation,omitempty"`
	// UUID of the table to read or write.
	// Wire name: 'table_id'
	TableId string `json:"table_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GenerateTemporaryTableCredentialRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := generateTemporaryTableCredentialRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GenerateTemporaryTableCredentialResponse struct {

	// Wire name: 'aws_temp_credentials'
	AwsTempCredentials *AwsCredentials `json:"aws_temp_credentials,omitempty"`

	// Wire name: 'azure_aad'
	AzureAad *AzureActiveDirectoryToken `json:"azure_aad,omitempty"`

	// Wire name: 'azure_user_delegation_sas'
	AzureUserDelegationSas *AzureUserDelegationSas `json:"azure_user_delegation_sas,omitempty"`
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	// Wire name: 'expiration_time'
	ExpirationTime int64 `json:"expiration_time,omitempty"`

	// Wire name: 'gcp_oauth_token'
	GcpOauthToken *GcpOauthToken `json:"gcp_oauth_token,omitempty"`

	// Wire name: 'r2_temp_credentials'
	R2TempCredentials *R2Credentials `json:"r2_temp_credentials,omitempty"`
	// The URL of the storage path accessible by the temporary credential.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GenerateTemporaryTableCredentialResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := generateTemporaryTableCredentialResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetAccountMetastoreAssignmentRequest struct {
	// Workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`
}

func (st GetAccountMetastoreAssignmentRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getAccountMetastoreAssignmentRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetAccountMetastoreRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" tf:"-"`
}

func (st GetAccountMetastoreRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getAccountMetastoreRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetAccountStorageCredentialRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" tf:"-"`
	// Name of the storage credential.
	StorageCredentialName string `json:"-" tf:"-"`
}

func (st GetAccountStorageCredentialRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getAccountStorageCredentialRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetArtifactAllowlistRequest struct {
	// The artifact type of the allowlist.
	ArtifactType ArtifactType `json:"-" tf:"-"`
}

func (st GetArtifactAllowlistRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getArtifactAllowlistRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetBindingsRequest struct {
	// Maximum number of workspace bindings to return. - When set to 0, the page
	// length is set to a server configured value (recommended); - When set to a
	// value greater than 0, the page length is the minimum of this value and a
	// server configured value; - When set to a value less than 0, an invalid
	// parameter error is returned; - If not set, all the workspace bindings are
	// returned (not recommended).
	MaxResults int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`
	// The name of the securable.
	SecurableName string `json:"-" tf:"-"`
	// The type of the securable to bind to a workspace (catalog,
	// storage_credential, credential, or external_location).
	SecurableType string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetBindingsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getBindingsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetByAliasRequest struct {
	// The name of the alias
	Alias string `json:"-" tf:"-"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" tf:"-"`
	// Whether to include aliases associated with the model version in the
	// response
	IncludeAliases bool `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetByAliasRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getByAliasRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetCatalogRequest struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" tf:"-"`
	// The name of the catalog.
	Name string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetCatalogRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getCatalogRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetCatalogWorkspaceBindingsResponse struct {
	// A list of workspace IDs
	// Wire name: 'workspaces'
	Workspaces []int64 `json:"workspaces,omitempty"`
}

func (st GetCatalogWorkspaceBindingsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getCatalogWorkspaceBindingsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetCatalogWorkspaceBindingsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCatalogWorkspaceBindingsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCatalogWorkspaceBindingsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCatalogWorkspaceBindingsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getCatalogWorkspaceBindingsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetConnectionRequest struct {
	// Name of the connection.
	Name string `json:"-" tf:"-"`
}

func (st GetConnectionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getConnectionRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetCredentialRequest struct {
	// Name of the credential.
	NameArg string `json:"-" tf:"-"`
}

func (st GetCredentialRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getCredentialRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetEffectiveRequest struct {
	// Full name of securable.
	FullName string `json:"-" tf:"-"`
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
	MaxResults int `json:"-" tf:"-"`
	// Opaque token for the next page of results (pagination).
	PageToken string `json:"-" tf:"-"`
	// If provided, only the effective permissions for the specified principal
	// (user or group) are returned.
	Principal string `json:"-" tf:"-"`
	// Type of securable.
	SecurableType string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetEffectiveRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getEffectiveRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetExternalLocationRequest struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `json:"-" tf:"-"`
	// Name of the external location.
	Name string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetExternalLocationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getExternalLocationRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetExternalMetadataRequest struct {
	Name string `json:"-" tf:"-"`
}

func (st GetExternalMetadataRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getExternalMetadataRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetExternalMetadataRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExternalMetadataRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExternalMetadataRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExternalMetadataRequest) MarshalJSON() ([]byte, error) {
	pb, err := getExternalMetadataRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetFunctionRequest struct {
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" tf:"-"`
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetFunctionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getFunctionRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetGrantRequest struct {
	// Full name of securable.
	FullName string `json:"-" tf:"-"`
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
	MaxResults int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`
	// If provided, only the permissions for the specified principal (user or
	// group) are returned.
	Principal string `json:"-" tf:"-"`
	// Type of securable.
	SecurableType string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetGrantRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getGrantRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetMetastoreRequest struct {
	// Unique ID of the metastore.
	Id string `json:"-" tf:"-"`
}

func (st GetMetastoreRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getMetastoreRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetMetastoreSummaryResponse struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	// Wire name: 'cloud'
	Cloud string `json:"cloud,omitempty"`
	// Time at which this metastore was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of metastore creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	// Wire name: 'default_data_access_config_id'
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	// Wire name: 'delta_sharing_organization_name'
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	// Wire name: 'delta_sharing_recipient_token_lifetime_in_seconds'
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	// Wire name: 'delta_sharing_scope'
	DeltaSharingScope DeltaSharingScopeEnum `json:"delta_sharing_scope,omitempty"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	// Wire name: 'external_access_enabled'
	ExternalAccessEnabled bool `json:"external_access_enabled,omitempty"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	// Wire name: 'global_metastore_id'
	GlobalMetastoreId string `json:"global_metastore_id,omitempty"`
	// Unique identifier of metastore.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// The user-specified name of the metastore.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The owner of the metastore.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	// Wire name: 'privilege_model_version'
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	// Wire name: 'region'
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	// Wire name: 'storage_root'
	StorageRoot string `json:"storage_root,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	// Wire name: 'storage_root_credential_id'
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// Name of the storage credential to access the metastore storage_root.
	// Wire name: 'storage_root_credential_name'
	StorageRootCredentialName string `json:"storage_root_credential_name,omitempty"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the metastore.
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetMetastoreSummaryResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getMetastoreSummaryResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetModelVersionRequest struct {
	// The three-level (fully qualified) name of the model version
	FullName string `json:"-" tf:"-"`
	// Whether to include aliases associated with the model version in the
	// response
	IncludeAliases bool `json:"-" tf:"-"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse bool `json:"-" tf:"-"`
	// The integer version number of the model version
	Version int `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetModelVersionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getModelVersionRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetOnlineTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name string `json:"-" tf:"-"`
}

func (st GetOnlineTableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getOnlineTableRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetPermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// The privileges assigned to each principal
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetPermissionsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getPermissionsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetPermissionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPermissionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPermissionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPermissionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getPermissionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetQualityMonitorRequest struct {
	// Full name of the table.
	TableName string `json:"-" tf:"-"`
}

func (st GetQualityMonitorRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getQualityMonitorRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetQuotaRequest struct {
	// Full name of the parent resource. Provide the metastore ID if the parent
	// is a metastore.
	ParentFullName string `json:"-" tf:"-"`
	// Securable type of the quota parent.
	ParentSecurableType string `json:"-" tf:"-"`
	// Name of the quota. Follows the pattern of the quota type, with "-quota"
	// added as a suffix.
	QuotaName string `json:"-" tf:"-"`
}

func (st GetQuotaRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getQuotaRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetQuotaResponse struct {
	// The returned QuotaInfo.
	// Wire name: 'quota_info'
	QuotaInfo *QuotaInfo `json:"quota_info,omitempty"`
}

func (st GetQuotaResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getQuotaResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetRefreshRequest struct {
	// ID of the refresh.
	RefreshId string `json:"-" tf:"-"`
	// Full name of the table.
	TableName string `json:"-" tf:"-"`
}

func (st GetRefreshRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getRefreshRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetRegisteredModelRequest struct {
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" tf:"-"`
	// Whether to include registered model aliases in the response
	IncludeAliases bool `json:"-" tf:"-"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetRegisteredModelRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getRegisteredModelRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetSchemaRequest struct {
	// Full name of the schema.
	FullName string `json:"-" tf:"-"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetSchemaRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getSchemaRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetStorageCredentialRequest struct {
	// Name of the storage credential.
	Name string `json:"-" tf:"-"`
}

func (st GetStorageCredentialRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getStorageCredentialRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetTableRequest struct {
	// Full name of the table.
	FullName string `json:"-" tf:"-"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for.
	IncludeBrowse bool `json:"-" tf:"-"`
	// Whether delta metadata should be included in the response.
	IncludeDeltaMetadata bool `json:"-" tf:"-"`
	// Whether to include a manifest containing table capabilities in the
	// response.
	IncludeManifestCapabilities bool `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetTableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getTableRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetWorkspaceBindingRequest struct {
	// The name of the catalog.
	Name string `json:"-" tf:"-"`
}

func (st GetWorkspaceBindingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getWorkspaceBindingRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetWorkspaceBindingsResponse struct {
	// List of workspace bindings
	// Wire name: 'bindings'
	Bindings []WorkspaceBinding `json:"bindings,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetWorkspaceBindingsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getWorkspaceBindingsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetWorkspaceBindingsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getWorkspaceBindingsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getWorkspaceBindingsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetWorkspaceBindingsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getWorkspaceBindingsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type ListAccountMetastoreAssignmentsRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" tf:"-"`
}

func (st ListAccountMetastoreAssignmentsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listAccountMetastoreAssignmentsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// The list of workspaces to which the given metastore is assigned.
type ListAccountMetastoreAssignmentsResponse struct {

	// Wire name: 'workspace_ids'
	WorkspaceIds []int64 `json:"workspace_ids,omitempty"`
}

func (st ListAccountMetastoreAssignmentsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listAccountMetastoreAssignmentsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListAccountStorageCredentialsRequest struct {
	// Unity Catalog metastore ID
	MetastoreId string `json:"-" tf:"-"`
}

func (st ListAccountStorageCredentialsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listAccountStorageCredentialsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListAccountStorageCredentialsResponse struct {
	// An array of metastore storage credentials.
	// Wire name: 'storage_credentials'
	StorageCredentials []StorageCredentialInfo `json:"storage_credentials,omitempty"`
}

func (st ListAccountStorageCredentialsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listAccountStorageCredentialsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListCatalogsRequest struct {
	// Whether to include catalogs in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" tf:"-"`
	// Maximum number of catalogs to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid catalogs are returned (not
	// recommended). - Note: The number of returned catalogs might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further catalogs can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCatalogsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listCatalogsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListCatalogsResponse struct {
	// An array of catalog information objects.
	// Wire name: 'catalogs'
	Catalogs []CatalogInfo `json:"catalogs,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCatalogsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listCatalogsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListConnectionsRequest struct {
	// Maximum number of connections to return. - If not set, all connections
	// are returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListConnectionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listConnectionsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListConnectionsResponse struct {
	// An array of connection information objects.
	// Wire name: 'connections'
	Connections []ConnectionInfo `json:"connections,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListConnectionsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listConnectionsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListCredentialsRequest struct {
	// Maximum number of credentials to return. - If not set, the default max
	// page size is used. - When set to a value greater than 0, the page length
	// is the minimum of this value and a server-configured value. - When set to
	// 0, the page length is set to a server-configured value (recommended). -
	// When set to a value less than 0, an invalid parameter error is returned.
	MaxResults int `json:"-" tf:"-"`
	// Opaque token to retrieve the next page of results.
	PageToken string `json:"-" tf:"-"`
	// Return only credentials for the specified purpose.
	Purpose CredentialPurpose `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCredentialsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listCredentialsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListCredentialsResponse struct {

	// Wire name: 'credentials'
	Credentials []CredentialInfo `json:"credentials,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCredentialsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listCredentialsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListExternalLineageRelationshipsRequest struct {
	// The lineage direction to filter on.
	LineageDirection LineageDirection `json:"-" tf:"-"`
	// The object to query external lineage relationship on.
	ObjectInfo ExternalLineageObject `json:"-" tf:"-"`

	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListExternalLineageRelationshipsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listExternalLineageRelationshipsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListExternalLineageRelationshipsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExternalLineageRelationshipsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExternalLineageRelationshipsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExternalLineageRelationshipsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listExternalLineageRelationshipsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListExternalLineageRelationshipsResponse struct {

	// Wire name: 'external_lineage_relationships'
	ExternalLineageRelationships []ExternalLineageInfo `json:"external_lineage_relationships,omitempty"`

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListExternalLineageRelationshipsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listExternalLineageRelationshipsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListExternalLineageRelationshipsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExternalLineageRelationshipsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExternalLineageRelationshipsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExternalLineageRelationshipsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listExternalLineageRelationshipsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListExternalLocationsRequest struct {
	// Whether to include external locations in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `json:"-" tf:"-"`
	// Maximum number of external locations to return. If not set, all the
	// external locations are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListExternalLocationsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listExternalLocationsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListExternalLocationsResponse struct {
	// An array of external locations.
	// Wire name: 'external_locations'
	ExternalLocations []ExternalLocationInfo `json:"external_locations,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListExternalLocationsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listExternalLocationsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListExternalMetadataRequest struct {
	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListExternalMetadataRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listExternalMetadataRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListExternalMetadataRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExternalMetadataRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExternalMetadataRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExternalMetadataRequest) MarshalJSON() ([]byte, error) {
	pb, err := listExternalMetadataRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListExternalMetadataResponse struct {

	// Wire name: 'external_metadata'
	ExternalMetadata []ExternalMetadata `json:"external_metadata,omitempty"`

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListExternalMetadataResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listExternalMetadataResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListExternalMetadataResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExternalMetadataResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExternalMetadataResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExternalMetadataResponse) MarshalJSON() ([]byte, error) {
	pb, err := listExternalMetadataResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListFunctionsRequest struct {
	// Name of parent catalog for functions of interest.
	CatalogName string `json:"-" tf:"-"`
	// Whether to include functions in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" tf:"-"`
	// Maximum number of functions to return. If not set, all the functions are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`
	// Parent schema of functions.
	SchemaName string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListFunctionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listFunctionsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListFunctionsResponse struct {
	// An array of function information objects.
	// Wire name: 'functions'
	Functions []FunctionInfo `json:"functions,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListFunctionsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listFunctionsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	MaxResults int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListMetastoresRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listMetastoresRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListMetastoresRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listMetastoresRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listMetastoresRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListMetastoresRequest) MarshalJSON() ([]byte, error) {
	pb, err := listMetastoresRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListMetastoresResponse struct {
	// An array of metastore information objects.
	// Wire name: 'metastores'
	Metastores []MetastoreInfo `json:"metastores,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListMetastoresResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listMetastoresResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListModelVersionsRequest struct {
	// The full three-level name of the registered model under which to list
	// model versions
	FullName string `json:"-" tf:"-"`
	// Whether to include model versions in the response for which the principal
	// can only access selective metadata for
	IncludeBrowse bool `json:"-" tf:"-"`
	// Maximum number of model versions to return. If not set, the page length
	// is set to a server configured value (100, as of 1/3/2024). - when set to
	// a value greater than 0, the page length is the minimum of this value and
	// a server configured value(1000, as of 1/3/2024); - when set to 0, the
	// page length is set to a server configured value (100, as of 1/3/2024)
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListModelVersionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listModelVersionsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListModelVersionsResponse struct {

	// Wire name: 'model_versions'
	ModelVersions []ModelVersionInfo `json:"model_versions,omitempty"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListModelVersionsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listModelVersionsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListQuotasRequest struct {
	// The number of quotas to return.
	MaxResults int `json:"-" tf:"-"`
	// Opaque token for the next page of results.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListQuotasRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listQuotasRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListQuotasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of returned QuotaInfos.
	// Wire name: 'quotas'
	Quotas []QuotaInfo `json:"quotas,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListQuotasResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listQuotasResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListRefreshesRequest struct {
	// Full name of the table.
	TableName string `json:"-" tf:"-"`
}

func (st ListRefreshesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listRefreshesRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListRegisteredModelsRequest struct {
	// The identifier of the catalog under which to list registered models. If
	// specified, schema_name must be specified.
	CatalogName string `json:"-" tf:"-"`
	// Whether to include registered models in the response for which the
	// principal can only access selective metadata for
	IncludeBrowse bool `json:"-" tf:"-"`
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
	MaxResults int `json:"-" tf:"-"`
	// Opaque token to send for the next page of results (pagination).
	PageToken string `json:"-" tf:"-"`
	// The identifier of the schema under which to list registered models. If
	// specified, catalog_name must be specified.
	SchemaName string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListRegisteredModelsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listRegisteredModelsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListRegisteredModelsResponse struct {
	// Opaque token for pagination. Omitted if there are no more results.
	// page_token should be set to this value for fetching the next page.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'registered_models'
	RegisteredModels []RegisteredModelInfo `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListRegisteredModelsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listRegisteredModelsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListSchemasRequest struct {
	// Parent catalog for schemas of interest.
	CatalogName string `json:"-" tf:"-"`
	// Whether to include schemas in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" tf:"-"`
	// Maximum number of schemas to return. If not set, all the schemas are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListSchemasRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listSchemasRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListSchemasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of schema information objects.
	// Wire name: 'schemas'
	Schemas []SchemaInfo `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListSchemasResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listSchemasResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListStorageCredentialsRequest struct {
	// Maximum number of storage credentials to return. If not set, all the
	// storage credentials are returned (not recommended). - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to 0, the page length is set to a server
	// configured value (recommended); - when set to a value less than 0, an
	// invalid parameter error is returned;
	MaxResults int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListStorageCredentialsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listStorageCredentialsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListStorageCredentialsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'storage_credentials'
	StorageCredentials []StorageCredentialInfo `json:"storage_credentials,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListStorageCredentialsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listStorageCredentialsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListSummariesRequest struct {
	// Name of parent catalog for tables of interest.
	CatalogName string `json:"-" tf:"-"`
	// Whether to include a manifest containing table capabilities in the
	// response.
	IncludeManifestCapabilities bool `json:"-" tf:"-"`
	// Maximum number of summaries for tables to return. If not set, the page
	// length is set to a server configured value (10000, as of 1/5/2024). -
	// when set to a value greater than 0, the page length is the minimum of
	// this value and a server configured value (10000, as of 1/5/2024); - when
	// set to 0, the page length is set to a server configured value (10000, as
	// of 1/5/2024) (recommended); - when set to a value less than 0, an invalid
	// parameter error is returned;
	MaxResults int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`
	// A sql LIKE pattern (% and _) for schema names. All schemas will be
	// returned if not set or empty.
	SchemaNamePattern string `json:"-" tf:"-"`
	// A sql LIKE pattern (% and _) for table names. All tables will be returned
	// if not set or empty.
	TableNamePattern string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListSummariesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listSummariesRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListSystemSchemasRequest struct {
	// Maximum number of schemas to return. - When set to 0, the page length is
	// set to a server configured value (recommended); - When set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - When set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all the schemas are returned (not
	// recommended).
	MaxResults int `json:"-" tf:"-"`
	// The ID for the metastore in which the system schema resides.
	MetastoreId string `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListSystemSchemasRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listSystemSchemasRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListSystemSchemasResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of system schema information objects.
	// Wire name: 'schemas'
	Schemas []SystemSchemaInfo `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListSystemSchemasResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listSystemSchemasResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListTableSummariesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of table summaries.
	// Wire name: 'tables'
	Tables []TableSummary `json:"tables,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListTableSummariesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listTableSummariesResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListTablesRequest struct {
	// Name of parent catalog for tables of interest.
	CatalogName string `json:"-" tf:"-"`
	// Whether to include tables in the response for which the principal can
	// only access selective metadata for.
	IncludeBrowse bool `json:"-" tf:"-"`
	// Whether to include a manifest containing table capabilities in the
	// response.
	IncludeManifestCapabilities bool `json:"-" tf:"-"`
	// Maximum number of tables to return. If not set, all the tables are
	// returned (not recommended). - when set to a value greater than 0, the
	// page length is the minimum of this value and a server configured value; -
	// when set to 0, the page length is set to a server configured value
	// (recommended); - when set to a value less than 0, an invalid parameter
	// error is returned;
	MaxResults int `json:"-" tf:"-"`
	// Whether to omit the columns of the table from the response or not.
	OmitColumns bool `json:"-" tf:"-"`
	// Whether to omit the properties of the table from the response or not.
	OmitProperties bool `json:"-" tf:"-"`
	// Whether to omit the username of the table (e.g. owner, updated_by,
	// created_by) from the response or not.
	OmitUsername bool `json:"-" tf:"-"`
	// Opaque token to send for the next page of results (pagination).
	PageToken string `json:"-" tf:"-"`
	// Parent schema of tables.
	SchemaName string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListTablesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listTablesRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListTablesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of table information objects.
	// Wire name: 'tables'
	Tables []TableInfo `json:"tables,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListTablesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listTablesResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListVolumesRequest struct {
	// The identifier of the catalog
	CatalogName string `json:"-" tf:"-"`
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" tf:"-"`
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
	MaxResults int `json:"-" tf:"-"`
	// Opaque token returned by a previous request. It must be included in the
	// request to retrieve the next page of results (pagination).
	PageToken string `json:"-" tf:"-"`
	// The identifier of the schema
	SchemaName string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListVolumesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listVolumesRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListVolumesResponseContent struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request to retrieve the next page of results.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'volumes'
	Volumes []VolumeInfo `json:"volumes,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListVolumesResponseContent) EncodeValues(key string, v *url.Values) error {
	pb, err := listVolumesResponseContentToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type MetastoreAssignment struct {
	// The name of the default catalog in the metastore.
	// Wire name: 'default_catalog_name'
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	// The unique ID of the metastore.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id"`
	// The unique ID of the Databricks workspace.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `json:"workspace_id"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st MetastoreAssignment) EncodeValues(key string, v *url.Values) error {
	pb, err := metastoreAssignmentToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type MetastoreInfo struct {
	// Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
	// Wire name: 'cloud'
	Cloud string `json:"cloud,omitempty"`
	// Time at which this metastore was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of metastore creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of the metastore's (Default) Data Access Configuration.
	// Wire name: 'default_data_access_config_id'
	DefaultDataAccessConfigId string `json:"default_data_access_config_id,omitempty"`
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	// Wire name: 'delta_sharing_organization_name'
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	// Wire name: 'delta_sharing_recipient_token_lifetime_in_seconds'
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	// Wire name: 'delta_sharing_scope'
	DeltaSharingScope DeltaSharingScopeEnum `json:"delta_sharing_scope,omitempty"`
	// Whether to allow non-DBR clients to directly access entities under the
	// metastore.
	// Wire name: 'external_access_enabled'
	ExternalAccessEnabled bool `json:"external_access_enabled,omitempty"`
	// Globally unique metastore ID across clouds and regions, of the form
	// `cloud:region:metastore_id`.
	// Wire name: 'global_metastore_id'
	GlobalMetastoreId string `json:"global_metastore_id,omitempty"`
	// Unique identifier of metastore.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// The user-specified name of the metastore.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The owner of the metastore.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	// Wire name: 'privilege_model_version'
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
	// Wire name: 'region'
	Region string `json:"region,omitempty"`
	// The storage root URL for metastore
	// Wire name: 'storage_root'
	StorageRoot string `json:"storage_root,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	// Wire name: 'storage_root_credential_id'
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`
	// Name of the storage credential to access the metastore storage_root.
	// Wire name: 'storage_root_credential_name'
	StorageRootCredentialName string `json:"storage_root_credential_name,omitempty"`
	// Time at which the metastore was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the metastore.
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st MetastoreInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := metastoreInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ModelVersionInfo struct {
	// List of aliases associated with the model version
	// Wire name: 'aliases'
	Aliases []RegisteredModelAlias `json:"aliases,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The name of the catalog containing the model version
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the model version
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`

	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the model version
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// The unique identifier of the model version
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// The unique identifier of the metastore containing the model version
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the parent registered model of the model version, relative to
	// parent schema
	// Wire name: 'model_name'
	ModelName string `json:"model_name,omitempty"`
	// Model version dependencies, for feature-store packaged models
	// Wire name: 'model_version_dependencies'
	ModelVersionDependencies *DependencyList `json:"model_version_dependencies,omitempty"`
	// MLflow run ID used when creating the model version, if ``source`` was
	// generated by an experiment run stored in an MLflow tracking server
	// Wire name: 'run_id'
	RunId string `json:"run_id,omitempty"`
	// ID of the Databricks workspace containing the MLflow run that generated
	// this model version, if applicable
	// Wire name: 'run_workspace_id'
	RunWorkspaceId int `json:"run_workspace_id,omitempty"`
	// The name of the schema containing the model version, relative to parent
	// catalog
	// Wire name: 'schema_name'
	SchemaName string `json:"schema_name,omitempty"`
	// URI indicating the location of the source artifacts (files) for the model
	// version
	// Wire name: 'source'
	Source string `json:"source,omitempty"`
	// Current status of the model version. Newly created model versions start
	// in PENDING_REGISTRATION status, then move to READY status once the model
	// version files are uploaded and the model version is finalized. Only model
	// versions in READY status can be loaded for inference or served.
	// Wire name: 'status'
	Status ModelVersionInfoStatus `json:"status,omitempty"`
	// The storage location on the cloud under which model version data files
	// are stored
	// Wire name: 'storage_location'
	StorageLocation string `json:"storage_location,omitempty"`

	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the model version last time
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`
	// Integer model version number, used to reference the model version in API
	// requests.
	// Wire name: 'version'
	Version int `json:"version,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ModelVersionInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := modelVersionInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type MonitorCronSchedule struct {
	// Read only field that indicates whether a schedule is paused or not.
	// Wire name: 'pause_status'
	PauseStatus MonitorCronSchedulePauseStatus `json:"pause_status,omitempty"`
	// The expression that determines when to run the monitor. See [examples].
	//
	// [examples]: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	// Wire name: 'quartz_cron_expression'
	QuartzCronExpression string `json:"quartz_cron_expression"`
	// The timezone id (e.g., ``"PST"``) in which to evaluate the quartz
	// expression.
	// Wire name: 'timezone_id'
	TimezoneId string `json:"timezone_id"`
}

func (st MonitorCronSchedule) EncodeValues(key string, v *url.Values) error {
	pb, err := monitorCronScheduleToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Read only field that indicates whether a schedule is paused or not.
type MonitorCronSchedulePauseStatus string

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

// Values returns all possible values for MonitorCronSchedulePauseStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *MonitorCronSchedulePauseStatus) Values() []MonitorCronSchedulePauseStatus {
	return []MonitorCronSchedulePauseStatus{
		MonitorCronSchedulePauseStatusPaused,
		MonitorCronSchedulePauseStatusUnpaused,
	}
}

// Type always returns MonitorCronSchedulePauseStatus to satisfy [pflag.Value] interface
func (f *MonitorCronSchedulePauseStatus) Type() string {
	return "MonitorCronSchedulePauseStatus"
}

type MonitorDataClassificationConfig struct {
	// Whether data classification is enabled.
	// Wire name: 'enabled'
	Enabled bool `json:"enabled,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st MonitorDataClassificationConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := monitorDataClassificationConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type MonitorDestination struct {
	// The list of email addresses to send the notification to. A maximum of 5
	// email addresses is supported.
	// Wire name: 'email_addresses'
	EmailAddresses []string `json:"email_addresses,omitempty"`
}

func (st MonitorDestination) EncodeValues(key string, v *url.Values) error {
	pb, err := monitorDestinationToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type MonitorInferenceLog struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	// Wire name: 'granularities'
	Granularities []string `json:"granularities"`
	// Optional column that contains the ground truth for the prediction.
	// Wire name: 'label_col'
	LabelCol string `json:"label_col,omitempty"`
	// Column that contains the id of the model generating the predictions.
	// Metrics will be computed per model id by default, and also across all
	// model ids.
	// Wire name: 'model_id_col'
	ModelIdCol string `json:"model_id_col"`
	// Column that contains the output/prediction from the model.
	// Wire name: 'prediction_col'
	PredictionCol string `json:"prediction_col"`
	// Optional column that contains the prediction probabilities for each class
	// in a classification problem type. The values in this column should be a
	// map, mapping each class label to the prediction probability for a given
	// sample. The map should be of PySpark MapType().
	// Wire name: 'prediction_proba_col'
	PredictionProbaCol string `json:"prediction_proba_col,omitempty"`
	// Problem type the model aims to solve. Determines the type of
	// model-quality metrics that will be computed.
	// Wire name: 'problem_type'
	ProblemType MonitorInferenceLogProblemType `json:"problem_type"`
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	// Wire name: 'timestamp_col'
	TimestampCol string `json:"timestamp_col"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st MonitorInferenceLog) EncodeValues(key string, v *url.Values) error {
	pb, err := monitorInferenceLogToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Problem type the model aims to solve. Determines the type of model-quality
// metrics that will be computed.
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

type MonitorInfo struct {
	// The directory to store monitoring assets (e.g. dashboard, metric tables).
	// Wire name: 'assets_dir'
	AssetsDir string `json:"assets_dir,omitempty"`
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	// Wire name: 'baseline_table_name'
	BaselineTableName string `json:"baseline_table_name,omitempty"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	// Wire name: 'custom_metrics'
	CustomMetrics []MonitorMetric `json:"custom_metrics,omitempty"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	// Wire name: 'dashboard_id'
	DashboardId string `json:"dashboard_id,omitempty"`
	// The data classification config for the monitor.
	// Wire name: 'data_classification_config'
	DataClassificationConfig *MonitorDataClassificationConfig `json:"data_classification_config,omitempty"`
	// The full name of the drift metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	// Wire name: 'drift_metrics_table_name'
	DriftMetricsTableName string `json:"drift_metrics_table_name"`
	// Configuration for monitoring inference logs.
	// Wire name: 'inference_log'
	InferenceLog *MonitorInferenceLog `json:"inference_log,omitempty"`
	// The latest failure message of the monitor (if any).
	// Wire name: 'latest_monitor_failure_msg'
	LatestMonitorFailureMsg string `json:"latest_monitor_failure_msg,omitempty"`
	// The version of the monitor config (e.g. 1,2,3). If negative, the monitor
	// may be corrupted.
	// Wire name: 'monitor_version'
	MonitorVersion string `json:"monitor_version"`
	// The notification settings for the monitor.
	// Wire name: 'notifications'
	Notifications *MonitorNotifications `json:"notifications,omitempty"`
	// Schema where output metric tables are created.
	// Wire name: 'output_schema_name'
	OutputSchemaName string `json:"output_schema_name,omitempty"`
	// The full name of the profile metrics table. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	// Wire name: 'profile_metrics_table_name'
	ProfileMetricsTableName string `json:"profile_metrics_table_name"`
	// The schedule for automatically updating and refreshing metric tables.
	// Wire name: 'schedule'
	Schedule *MonitorCronSchedule `json:"schedule,omitempty"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	// Wire name: 'slicing_exprs'
	SlicingExprs []string `json:"slicing_exprs,omitempty"`
	// Configuration for monitoring snapshot tables.
	// Wire name: 'snapshot'
	Snapshot *MonitorSnapshot `json:"snapshot,omitempty"`

	// Wire name: 'status'
	Status MonitorInfoStatus `json:"status"`
	// The full name of the table to monitor. Format:
	// __catalog_name__.__schema_name__.__table_name__.
	// Wire name: 'table_name'
	TableName string `json:"table_name"`
	// Configuration for monitoring time series tables.
	// Wire name: 'time_series'
	TimeSeries *MonitorTimeSeries `json:"time_series,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st MonitorInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := monitorInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// The status of the monitor.
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

type MonitorMetric struct {
	// Jinja template for a SQL expression that specifies how to compute the
	// metric. See [create metric definition].
	//
	// [create metric definition]: https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition
	// Wire name: 'definition'
	Definition string `json:"definition"`
	// A list of column names in the input table the metric should be computed
	// for. Can use ``":table"`` to indicate that the metric needs information
	// from multiple columns.
	// Wire name: 'input_columns'
	InputColumns []string `json:"input_columns"`
	// Name of the metric in the output tables.
	// Wire name: 'name'
	Name string `json:"name"`
	// The output type of the custom metric.
	// Wire name: 'output_data_type'
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
	// Wire name: 'type'
	Type MonitorMetricType `json:"type"`
}

func (st MonitorMetric) EncodeValues(key string, v *url.Values) error {
	pb, err := monitorMetricToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type MonitorNotifications struct {
	// Who to send notifications to on monitor failure.
	// Wire name: 'on_failure'
	OnFailure *MonitorDestination `json:"on_failure,omitempty"`
	// Who to send notifications to when new data classification tags are
	// detected.
	// Wire name: 'on_new_classification_tag_detected'
	OnNewClassificationTagDetected *MonitorDestination `json:"on_new_classification_tag_detected,omitempty"`
}

func (st MonitorNotifications) EncodeValues(key string, v *url.Values) error {
	pb, err := monitorNotificationsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type MonitorRefreshInfo struct {
	// Time at which refresh operation completed (milliseconds since 1/1/1970
	// UTC).
	// Wire name: 'end_time_ms'
	EndTimeMs int64 `json:"end_time_ms,omitempty"`
	// An optional message to give insight into the current state of the job
	// (e.g. FAILURE messages).
	// Wire name: 'message'
	Message string `json:"message,omitempty"`
	// Unique id of the refresh operation.
	// Wire name: 'refresh_id'
	RefreshId int64 `json:"refresh_id"`
	// Time at which refresh operation was initiated (milliseconds since
	// 1/1/1970 UTC).
	// Wire name: 'start_time_ms'
	StartTimeMs int64 `json:"start_time_ms"`
	// The current state of the refresh.
	// Wire name: 'state'
	State MonitorRefreshInfoState `json:"state"`
	// The method by which the refresh was triggered.
	// Wire name: 'trigger'
	Trigger MonitorRefreshInfoTrigger `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st MonitorRefreshInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := monitorRefreshInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// The current state of the refresh.
type MonitorRefreshInfoState string

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
	}
}

// Type always returns MonitorRefreshInfoState to satisfy [pflag.Value] interface
func (f *MonitorRefreshInfoState) Type() string {
	return "MonitorRefreshInfoState"
}

// The method by which the refresh was triggered.
type MonitorRefreshInfoTrigger string

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

// Values returns all possible values for MonitorRefreshInfoTrigger.
//
// There is no guarantee on the order of the values in the slice.
func (f *MonitorRefreshInfoTrigger) Values() []MonitorRefreshInfoTrigger {
	return []MonitorRefreshInfoTrigger{
		MonitorRefreshInfoTriggerManual,
		MonitorRefreshInfoTriggerSchedule,
	}
}

// Type always returns MonitorRefreshInfoTrigger to satisfy [pflag.Value] interface
func (f *MonitorRefreshInfoTrigger) Type() string {
	return "MonitorRefreshInfoTrigger"
}

type MonitorRefreshListResponse struct {
	// List of refreshes.
	// Wire name: 'refreshes'
	Refreshes []MonitorRefreshInfo `json:"refreshes,omitempty"`
}

func (st MonitorRefreshListResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := monitorRefreshListResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type MonitorSnapshot struct {
}

func (st MonitorSnapshot) EncodeValues(key string, v *url.Values) error {
	pb, err := monitorSnapshotToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type MonitorTimeSeries struct {
	// Granularities for aggregating data into time windows based on their
	// timestamp. Currently the following static granularities are supported:
	// {``"5 minutes"``, ``"30 minutes"``, ``"1 hour"``, ``"1 day"``, ``"<n>
	// week(s)"``, ``"1 month"``, ``"1 year"``}.
	// Wire name: 'granularities'
	Granularities []string `json:"granularities"`
	// Column that contains the timestamps of requests. The column must be one
	// of the following: - A ``TimestampType`` column - A column whose values
	// can be converted to timestamps through the pyspark ``to_timestamp``
	// [function].
	//
	// [function]: https://spark.apache.org/docs/latest/api/python/reference/pyspark.sql/api/pyspark.sql.functions.to_timestamp.html
	// Wire name: 'timestamp_col'
	TimestampCol string `json:"timestamp_col"`
}

func (st MonitorTimeSeries) EncodeValues(key string, v *url.Values) error {
	pb, err := monitorTimeSeriesToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type NamedTableConstraint struct {
	// The name of the constraint.
	// Wire name: 'name'
	Name string `json:"name"`
}

func (st NamedTableConstraint) EncodeValues(key string, v *url.Values) error {
	pb, err := namedTableConstraintToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Online Table information.
type OnlineTable struct {
	// Full three-part (catalog, schema, table) name of the table.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Specification of the online table.
	// Wire name: 'spec'
	Spec *OnlineTableSpec `json:"spec,omitempty"`
	// Online Table data synchronization status
	// Wire name: 'status'
	Status *OnlineTableStatus `json:"status,omitempty"`
	// Data serving REST API URL for this table
	// Wire name: 'table_serving_url'
	TableServingUrl string `json:"table_serving_url,omitempty"`
	// The provisioning state of the online table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	// Wire name: 'unity_catalog_provisioning_state'
	UnityCatalogProvisioningState ProvisioningInfoState `json:"unity_catalog_provisioning_state,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st OnlineTable) EncodeValues(key string, v *url.Values) error {
	pb, err := onlineTableToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	PerformFullCopy bool `json:"perform_full_copy,omitempty"`
	// ID of the associated pipeline. Generated by the server - cannot be set by
	// the caller.
	// Wire name: 'pipeline_id'
	PipelineId string `json:"pipeline_id,omitempty"`
	// Primary Key columns to be used for data insert/update in the destination.
	// Wire name: 'primary_key_columns'
	PrimaryKeyColumns []string `json:"primary_key_columns,omitempty"`
	// Pipeline runs continuously after generating the initial data.
	// Wire name: 'run_continuously'
	RunContinuously *OnlineTableSpecContinuousSchedulingPolicy `json:"run_continuously,omitempty"`
	// Pipeline stops after generating the initial data and can be triggered
	// later (manually, through a cron job or through data triggers)
	// Wire name: 'run_triggered'
	RunTriggered *OnlineTableSpecTriggeredSchedulingPolicy `json:"run_triggered,omitempty"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	// Wire name: 'source_table_full_name'
	SourceTableFullName string `json:"source_table_full_name,omitempty"`
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	// Wire name: 'timeseries_key'
	TimeseriesKey string `json:"timeseries_key,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st OnlineTableSpec) EncodeValues(key string, v *url.Values) error {
	pb, err := onlineTableSpecToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type OnlineTableSpecContinuousSchedulingPolicy struct {
}

func (st OnlineTableSpecContinuousSchedulingPolicy) EncodeValues(key string, v *url.Values) error {
	pb, err := onlineTableSpecContinuousSchedulingPolicyToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type OnlineTableSpecTriggeredSchedulingPolicy struct {
}

func (st OnlineTableSpecTriggeredSchedulingPolicy) EncodeValues(key string, v *url.Values) error {
	pb, err := onlineTableSpecTriggeredSchedulingPolicyToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Status of an online table.
type OnlineTableStatus struct {

	// Wire name: 'continuous_update_status'
	ContinuousUpdateStatus *ContinuousUpdateStatus `json:"continuous_update_status,omitempty"`
	// The state of the online table.
	// Wire name: 'detailed_state'
	DetailedState OnlineTableState `json:"detailed_state,omitempty"`

	// Wire name: 'failed_status'
	FailedStatus *FailedStatus `json:"failed_status,omitempty"`
	// A text description of the current state of the online table.
	// Wire name: 'message'
	Message string `json:"message,omitempty"`

	// Wire name: 'provisioning_status'
	ProvisioningStatus *ProvisioningStatus `json:"provisioning_status,omitempty"`

	// Wire name: 'triggered_update_status'
	TriggeredUpdateStatus *TriggeredUpdateStatus `json:"triggered_update_status,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st OnlineTableStatus) EncodeValues(key string, v *url.Values) error {
	pb, err := onlineTableStatusToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Spec of an allowed option on a securable kind and its attributes. This is
// mostly used by UI to provide user friendly hints and descriptions in order to
// facilitate the securable creation process.
type OptionSpec struct {
	// For drop down / radio button selections, UI will want to know the
	// possible input values, it can also be used by other option types to limit
	// input selections.
	// Wire name: 'allowed_values'
	AllowedValues []string `json:"allowed_values,omitempty"`
	// The default value of the option, for example, value '443' for 'port'
	// option.
	// Wire name: 'default_value'
	DefaultValue string `json:"default_value,omitempty"`
	// A concise user facing description of what the input value of this option
	// should look like.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// The hint is used on the UI to suggest what the input value can possibly
	// be like, for example: example.com for 'host' option. Unlike default
	// value, it will not be applied automatically without user input.
	// Wire name: 'hint'
	Hint string `json:"hint,omitempty"`
	// Indicates whether an option should be displayed with copy button on the
	// UI.
	// Wire name: 'is_copiable'
	IsCopiable bool `json:"is_copiable,omitempty"`
	// Indicates whether an option can be provided by users in the create/update
	// path of an entity.
	// Wire name: 'is_creatable'
	IsCreatable bool `json:"is_creatable,omitempty"`
	// Is the option value not user settable and is thus not shown on the UI.
	// Wire name: 'is_hidden'
	IsHidden bool `json:"is_hidden,omitempty"`
	// Specifies whether this option is safe to log, i.e. no sensitive
	// information.
	// Wire name: 'is_loggable'
	IsLoggable bool `json:"is_loggable,omitempty"`
	// Is the option required.
	// Wire name: 'is_required'
	IsRequired bool `json:"is_required,omitempty"`
	// Is the option value considered secret and thus redacted on the UI.
	// Wire name: 'is_secret'
	IsSecret bool `json:"is_secret,omitempty"`
	// Is the option updatable by users.
	// Wire name: 'is_updatable'
	IsUpdatable bool `json:"is_updatable,omitempty"`
	// The unique name of the option.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Specifies when the option value is displayed on the UI within the OAuth
	// flow.
	// Wire name: 'oauth_stage'
	OauthStage OptionSpecOauthStage `json:"oauth_stage,omitempty"`
	// The type of the option.
	// Wire name: 'type'
	Type OptionSpecOptionType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st OptionSpec) EncodeValues(key string, v *url.Values) error {
	pb, err := optionSpecToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *OptionSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &optionSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := optionSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st OptionSpec) MarshalJSON() ([]byte, error) {
	pb, err := optionSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type PermissionsChange struct {
	// The set of privileges to add.
	// Wire name: 'add'
	Add []Privilege `json:"add,omitempty"`
	// The principal whose privileges we are changing.
	// Wire name: 'principal'
	Principal string `json:"principal,omitempty"`
	// The set of privileges to remove.
	// Wire name: 'remove'
	Remove []Privilege `json:"remove,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PermissionsChange) EncodeValues(key string, v *url.Values) error {
	pb, err := permissionsChangeToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Progress information of the Online Table data synchronization pipeline.
type PipelineProgress struct {
	// The estimated time remaining to complete this update in seconds.
	// Wire name: 'estimated_completion_time_seconds'
	EstimatedCompletionTimeSeconds float64 `json:"estimated_completion_time_seconds,omitempty"`
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	// Wire name: 'latest_version_currently_processing'
	LatestVersionCurrentlyProcessing int64 `json:"latest_version_currently_processing,omitempty"`
	// The completion ratio of this update. This is a number between 0 and 1.
	// Wire name: 'sync_progress_completion'
	SyncProgressCompletion float64 `json:"sync_progress_completion,omitempty"`
	// The number of rows that have been synced in this update.
	// Wire name: 'synced_row_count'
	SyncedRowCount int64 `json:"synced_row_count,omitempty"`
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	// Wire name: 'total_row_count'
	TotalRowCount int64 `json:"total_row_count,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PipelineProgress) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelineProgressToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type PrimaryKeyConstraint struct {
	// Column names for this constraint.
	// Wire name: 'child_columns'
	ChildColumns []string `json:"child_columns"`
	// The name of the constraint.
	// Wire name: 'name'
	Name string `json:"name"`
	// True if the constraint is RELY, false or unset if NORELY.
	// Wire name: 'rely'
	Rely bool `json:"rely,omitempty"`
	// Column names that represent a timeseries.
	// Wire name: 'timeseries_columns'
	TimeseriesColumns []string `json:"timeseries_columns,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PrimaryKeyConstraint) EncodeValues(key string, v *url.Values) error {
	pb, err := primaryKeyConstraintToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type PrivilegeAssignment struct {
	// The principal (user email address or group name). For deleted principals,
	// `principal` is empty while `principal_id` is populated.
	// Wire name: 'principal'
	Principal string `json:"principal,omitempty"`
	// The privileges assigned to the principal.
	// Wire name: 'privileges'
	Privileges []Privilege `json:"privileges,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PrivilegeAssignment) EncodeValues(key string, v *url.Values) error {
	pb, err := privilegeAssignmentToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Status of an asynchronously provisioned resource.
type ProvisioningInfo struct {
	// The provisioning state of the resource.
	// Wire name: 'state'
	State ProvisioningInfoState `json:"state,omitempty"`
}

func (st ProvisioningInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := provisioningInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Detailed status of an online table. Shown if the online table is in the
// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT state.
type ProvisioningStatus struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	// Wire name: 'initial_pipeline_sync_progress'
	InitialPipelineSyncProgress *PipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
}

func (st ProvisioningStatus) EncodeValues(key string, v *url.Values) error {
	pb, err := provisioningStatusToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type QuotaInfo struct {
	// The timestamp that indicates when the quota count was last updated.
	// Wire name: 'last_refreshed_at'
	LastRefreshedAt int64 `json:"last_refreshed_at,omitempty"`
	// Name of the parent resource. Returns metastore ID if the parent is a
	// metastore.
	// Wire name: 'parent_full_name'
	ParentFullName string `json:"parent_full_name,omitempty"`
	// The quota parent securable type.
	// Wire name: 'parent_securable_type'
	ParentSecurableType SecurableType `json:"parent_securable_type,omitempty"`
	// The current usage of the resource quota.
	// Wire name: 'quota_count'
	QuotaCount int `json:"quota_count,omitempty"`
	// The current limit of the resource quota.
	// Wire name: 'quota_limit'
	QuotaLimit int `json:"quota_limit,omitempty"`
	// The name of the quota.
	// Wire name: 'quota_name'
	QuotaName string `json:"quota_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st QuotaInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := quotaInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// R2 temporary credentials for API authentication. Read more at
// https://developers.cloudflare.com/r2/api/s3/tokens/.
type R2Credentials struct {
	// The access key ID that identifies the temporary credentials.
	// Wire name: 'access_key_id'
	AccessKeyId string `json:"access_key_id,omitempty"`
	// The secret access key associated with the access key.
	// Wire name: 'secret_access_key'
	SecretAccessKey string `json:"secret_access_key,omitempty"`
	// The generated JWT that users must pass to use the temporary credentials.
	// Wire name: 'session_token'
	SessionToken string `json:"session_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st R2Credentials) EncodeValues(key string, v *url.Values) error {
	pb, err := r2CredentialsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ReadVolumeRequest struct {
	// Whether to include volumes in the response for which the principal can
	// only access selective metadata for
	IncludeBrowse bool `json:"-" tf:"-"`
	// The three-level (fully qualified) name of the volume
	Name string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ReadVolumeRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := readVolumeRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type RegenerateDashboardRequest struct {
	// Full name of the table.
	TableName string `json:"-" tf:"-"`
	// Optional argument to specify the warehouse for dashboard regeneration. If
	// not specified, the first running warehouse will be used.
	// Wire name: 'warehouse_id'
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RegenerateDashboardRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := regenerateDashboardRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type RegenerateDashboardResponse struct {
	// Id of the regenerated monitoring dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `json:"dashboard_id,omitempty"`
	// The directory where the regenerated dashboard is stored.
	// Wire name: 'parent_folder'
	ParentFolder string `json:"parent_folder,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RegenerateDashboardResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := regenerateDashboardResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Registered model alias.
type RegisteredModelAlias struct {
	// Name of the alias, e.g. 'champion' or 'latest_stable'
	// Wire name: 'alias_name'
	AliasName string `json:"alias_name,omitempty"`
	// Integer version number of the model version to which this alias points.
	// Wire name: 'version_num'
	VersionNum int `json:"version_num,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RegisteredModelAlias) EncodeValues(key string, v *url.Values) error {
	pb, err := registeredModelAliasToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type RegisteredModelInfo struct {
	// List of aliases associated with the registered model
	// Wire name: 'aliases'
	Aliases []RegisteredModelAlias `json:"aliases,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The name of the catalog where the schema and the registered model reside
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the registered model
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Creation timestamp of the registered model in milliseconds since the Unix
	// epoch
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the registered model
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// The three-level (fully qualified) name of the registered model
	// Wire name: 'full_name'
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the metastore
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the registered model
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The identifier of the user who owns the registered model
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// The name of the schema where the registered model resides
	// Wire name: 'schema_name'
	SchemaName string `json:"schema_name,omitempty"`
	// The storage location on the cloud under which model version data files
	// are stored
	// Wire name: 'storage_location'
	StorageLocation string `json:"storage_location,omitempty"`
	// Last-update timestamp of the registered model in milliseconds since the
	// Unix epoch
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the registered model last time
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RegisteredModelInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := registeredModelInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type RunRefreshRequest struct {
	// Full name of the table.
	TableName string `json:"-" tf:"-"`
}

func (st RunRefreshRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := runRefreshRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Next ID: 40
type SchemaInfo struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool `json:"browse_only,omitempty"`
	// Name of parent catalog.
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name,omitempty"`
	// The type of the parent catalog.
	// Wire name: 'catalog_type'
	CatalogType CatalogType `json:"catalog_type,omitempty"`
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Time at which this schema was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of schema creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`

	// Wire name: 'effective_predictive_optimization_flag'
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag `json:"effective_predictive_optimization_flag,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'enable_predictive_optimization'
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// Full name of schema, in form of __catalog_name__.__schema_name__.
	// Wire name: 'full_name'
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of schema, relative to parent catalog.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Username of current owner of schema.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string `json:"properties,omitempty"`
	// The unique identifier of the schema.
	// Wire name: 'schema_id'
	SchemaId string `json:"schema_id,omitempty"`
	// Storage location for managed tables within schema.
	// Wire name: 'storage_location'
	StorageLocation string `json:"storage_location,omitempty"`
	// Storage root URL for managed tables within schema.
	// Wire name: 'storage_root'
	StorageRoot string `json:"storage_root,omitempty"`
	// Time at which this schema was created, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified schema.
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SchemaInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := schemaInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Latest kind: TABLE_DELTA_ICEBERG_DELTASHARING = 252; Next id:253
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

// Manifest of a specific securable kind.
type SecurableKindManifest struct {
	// Privileges that can be assigned to the securable.
	// Wire name: 'assignable_privileges'
	AssignablePrivileges []string `json:"assignable_privileges,omitempty"`
	// A list of capabilities in the securable kind.
	// Wire name: 'capabilities'
	Capabilities []string `json:"capabilities,omitempty"`
	// Detailed specs of allowed options.
	// Wire name: 'options'
	Options []OptionSpec `json:"options,omitempty"`
	// Securable kind to get manifest of.
	// Wire name: 'securable_kind'
	SecurableKind SecurableKind `json:"securable_kind,omitempty"`
	// Securable Type of the kind.
	// Wire name: 'securable_type'
	SecurableType SecurableType `json:"securable_type,omitempty"`
}

func (st SecurableKindManifest) EncodeValues(key string, v *url.Values) error {
	pb, err := securableKindManifestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SecurableKindManifest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &securableKindManifestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := securableKindManifestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SecurableKindManifest) MarshalJSON() ([]byte, error) {
	pb, err := securableKindManifestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type SetArtifactAllowlist struct {
	// A list of allowed artifact match patterns.
	// Wire name: 'artifact_matchers'
	ArtifactMatchers []ArtifactMatcher `json:"artifact_matchers"`
	// The artifact type of the allowlist.
	ArtifactType ArtifactType `json:"-" tf:"-"`
	// Time at which this artifact allowlist was set, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of the user who set the artifact allowlist.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SetArtifactAllowlist) EncodeValues(key string, v *url.Values) error {
	pb, err := setArtifactAllowlistToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type SetRegisteredModelAliasRequest struct {
	// The name of the alias
	// Wire name: 'alias'
	Alias string `json:"alias"`
	// Full name of the registered model
	// Wire name: 'full_name'
	FullName string `json:"full_name"`
	// The version number of the model version to which the alias points
	// Wire name: 'version_num'
	VersionNum int `json:"version_num"`
}

func (st SetRegisteredModelAliasRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := setRegisteredModelAliasRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Server-Side Encryption properties for clients communicating with AWS s3.
type SseEncryptionDetails struct {
	// Sets the value of the 'x-amz-server-side-encryption' header in S3
	// request.
	// Wire name: 'algorithm'
	Algorithm SseEncryptionDetailsAlgorithm `json:"algorithm,omitempty"`
	// Optional. The ARN of the SSE-KMS key used with the S3 location, when
	// algorithm = "SSE-KMS". Sets the value of the
	// 'x-amz-server-side-encryption-aws-kms-key-id' header.
	// Wire name: 'aws_kms_key_arn'
	AwsKmsKeyArn string `json:"aws_kms_key_arn,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SseEncryptionDetails) EncodeValues(key string, v *url.Values) error {
	pb, err := sseEncryptionDetailsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type StorageCredentialInfo struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRoleResponse `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentityResponse `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	// Wire name: 'cloudflare_api_token'
	CloudflareApiToken *CloudflareApiToken `json:"cloudflare_api_token,omitempty"`
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Time at which this credential was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of credential creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// The Databricks managed GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountResponse `json:"databricks_gcp_service_account,omitempty"`
	// The full name of the credential.
	// Wire name: 'full_name'
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the credential.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Unique identifier of the parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// The credential name. The name must be unique among storage and service
	// credentials within the metastore.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Username of current owner of credential.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	// Wire name: 'read_only'
	ReadOnly bool `json:"read_only,omitempty"`
	// Time at which this credential was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the credential.
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`
	// Whether this credential is the current metastore's root storage
	// credential. Only applicable when purpose is **STORAGE**.
	// Wire name: 'used_for_managed_storage'
	UsedForManagedStorage bool `json:"used_for_managed_storage,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st StorageCredentialInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := storageCredentialInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type SystemSchemaInfo struct {
	// Name of the system schema.
	// Wire name: 'schema'
	Schema string `json:"schema"`
	// The current state of enablement for the system schema. An empty string
	// means the system schema is available and ready for opt-in. Possible
	// values: AVAILABLE | ENABLE_INITIALIZED | ENABLE_COMPLETED |
	// DISABLE_INITIALIZED | UNAVAILABLE
	// Wire name: 'state'
	State string `json:"state"`
}

func (st SystemSchemaInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := systemSchemaInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type SystemType string

const SystemTypeAmazonRedshift SystemType = `AMAZON_REDSHIFT`

const SystemTypeAzureSynapse SystemType = `AZURE_SYNAPSE`

const SystemTypeConfluent SystemType = `CONFLUENT`

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
	case `AMAZON_REDSHIFT`, `AZURE_SYNAPSE`, `CONFLUENT`, `GOOGLE_BIGQUERY`, `KAFKA`, `LOOKER`, `MICROSOFT_FABRIC`, `MICROSOFT_SQL_SERVER`, `MONGODB`, `MYSQL`, `ORACLE`, `OTHER`, `POSTGRESQL`, `POWER_BI`, `SALESFORCE`, `SAP`, `SERVICENOW`, `SNOWFLAKE`, `TABLEAU`, `TERADATA`, `WORKDAY`:
		*f = SystemType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AMAZON_REDSHIFT", "AZURE_SYNAPSE", "CONFLUENT", "GOOGLE_BIGQUERY", "KAFKA", "LOOKER", "MICROSOFT_FABRIC", "MICROSOFT_SQL_SERVER", "MONGODB", "MYSQL", "ORACLE", "OTHER", "POSTGRESQL", "POWER_BI", "SALESFORCE", "SAP", "SERVICENOW", "SNOWFLAKE", "TABLEAU", "TERADATA", "WORKDAY"`, v)
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

// A table constraint, as defined by *one* of the following fields being set:
// __primary_key_constraint__, __foreign_key_constraint__,
// __named_table_constraint__.
type TableConstraint struct {

	// Wire name: 'foreign_key_constraint'
	ForeignKeyConstraint *ForeignKeyConstraint `json:"foreign_key_constraint,omitempty"`

	// Wire name: 'named_table_constraint'
	NamedTableConstraint *NamedTableConstraint `json:"named_table_constraint,omitempty"`

	// Wire name: 'primary_key_constraint'
	PrimaryKeyConstraint *PrimaryKeyConstraint `json:"primary_key_constraint,omitempty"`
}

func (st TableConstraint) EncodeValues(key string, v *url.Values) error {
	pb, err := tableConstraintToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// A table that is dependent on a SQL object.
type TableDependency struct {
	// Full name of the dependent table, in the form of
	// __catalog_name__.__schema_name__.__table_name__.
	// Wire name: 'table_full_name'
	TableFullName string `json:"table_full_name"`
}

func (st TableDependency) EncodeValues(key string, v *url.Values) error {
	pb, err := tableDependencyToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type TableExistsResponse struct {
	// Whether the table exists or not.
	// Wire name: 'table_exists'
	TableExists bool `json:"table_exists,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TableExistsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := tableExistsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type TableInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	// Wire name: 'access_point'
	AccessPoint string `json:"access_point,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool `json:"browse_only,omitempty"`
	// Name of parent catalog.
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name,omitempty"`
	// The array of __ColumnInfo__ definitions of the table's columns.
	// Wire name: 'columns'
	Columns []ColumnInfo `json:"columns,omitempty"`
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Time at which this table was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of table creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// Unique ID of the Data Access Configuration to use with the table data.
	// Wire name: 'data_access_configuration_id'
	DataAccessConfigurationId string `json:"data_access_configuration_id,omitempty"`

	// Wire name: 'data_source_format'
	DataSourceFormat DataSourceFormat `json:"data_source_format,omitempty"`
	// Time at which this table was deleted, in epoch milliseconds. Field is
	// omitted if table is not deleted.
	// Wire name: 'deleted_at'
	DeletedAt int64 `json:"deleted_at,omitempty"`
	// Information pertaining to current state of the delta table.
	// Wire name: 'delta_runtime_properties_kvpairs'
	DeltaRuntimePropertiesKvpairs *DeltaRuntimePropertiesKvPairs `json:"delta_runtime_properties_kvpairs,omitempty"`

	// Wire name: 'effective_predictive_optimization_flag'
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag `json:"effective_predictive_optimization_flag,omitempty"`

	// Wire name: 'enable_predictive_optimization'
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`

	// Wire name: 'encryption_details'
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	// Full name of table, in form of
	// __catalog_name__.__schema_name__.__table_name__
	// Wire name: 'full_name'
	FullName string `json:"full_name,omitempty"`
	// Unique identifier of parent metastore.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of table, relative to parent schema.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Username of current owner of table.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// The pipeline ID of the table. Applicable for tables created by pipelines
	// (Materialized View, Streaming Table, etc.).
	// Wire name: 'pipeline_id'
	PipelineId string `json:"pipeline_id,omitempty"`
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string `json:"properties,omitempty"`

	// Wire name: 'row_filter'
	RowFilter *TableRowFilter `json:"row_filter,omitempty"`
	// Name of parent schema relative to its parent catalog.
	// Wire name: 'schema_name'
	SchemaName string `json:"schema_name,omitempty"`
	// SecurableKindManifest of table, including capabilities the table has.
	// Wire name: 'securable_kind_manifest'
	SecurableKindManifest *SecurableKindManifest `json:"securable_kind_manifest,omitempty"`
	// List of schemes whose objects can be referenced without qualification.
	// Wire name: 'sql_path'
	SqlPath string `json:"sql_path,omitempty"`
	// Name of the storage credential, when a storage credential is configured
	// for use with this table.
	// Wire name: 'storage_credential_name'
	StorageCredentialName string `json:"storage_credential_name,omitempty"`
	// Storage root URL for table (for **MANAGED**, **EXTERNAL** tables).
	// Wire name: 'storage_location'
	StorageLocation string `json:"storage_location,omitempty"`
	// List of table constraints. Note: this field is not set in the output of
	// the __listTables__ API.
	// Wire name: 'table_constraints'
	TableConstraints []TableConstraint `json:"table_constraints,omitempty"`
	// The unique identifier of the table.
	// Wire name: 'table_id'
	TableId string `json:"table_id,omitempty"`

	// Wire name: 'table_type'
	TableType TableType `json:"table_type,omitempty"`
	// Time at which this table was last modified, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified the table.
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`
	// View definition SQL (when __table_type__ is **VIEW**,
	// **MATERIALIZED_VIEW**, or **STREAMING_TABLE**)
	// Wire name: 'view_definition'
	ViewDefinition string `json:"view_definition,omitempty"`
	// View dependencies (when table_type == **VIEW** or **MATERIALIZED_VIEW**,
	// **STREAMING_TABLE**) - when DependencyList is None, the dependency is not
	// provided; - when DependencyList is an empty list, the dependency is
	// provided but is empty; - when DependencyList is not an empty list,
	// dependencies are provided and recorded.
	// Wire name: 'view_dependencies'
	ViewDependencies *DependencyList `json:"view_dependencies,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TableInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := tableInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type TableRowFilter struct {
	// The full name of the row filter SQL UDF.
	// Wire name: 'function_name'
	FunctionName string `json:"function_name"`
	// The list of table columns to be passed as input to the row filter
	// function. The column types should match the types of the filter function
	// arguments.
	// Wire name: 'input_column_names'
	InputColumnNames []string `json:"input_column_names"`
}

func (st TableRowFilter) EncodeValues(key string, v *url.Values) error {
	pb, err := tableRowFilterToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type TableSummary struct {
	// The full name of the table.
	// Wire name: 'full_name'
	FullName string `json:"full_name,omitempty"`
	// SecurableKindManifest of table, including capabilities the table has.
	// Wire name: 'securable_kind_manifest'
	SecurableKindManifest *SecurableKindManifest `json:"securable_kind_manifest,omitempty"`

	// Wire name: 'table_type'
	TableType TableType `json:"table_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TableSummary) EncodeValues(key string, v *url.Values) error {
	pb, err := tableSummaryToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type TagKeyValue struct {
	// name of the tag
	// Wire name: 'key'
	Key string `json:"key,omitempty"`
	// value of the tag associated with the key, could be optional
	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TagKeyValue) EncodeValues(key string, v *url.Values) error {
	pb, err := tagKeyValueToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *TagKeyValue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tagKeyValuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tagKeyValueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TagKeyValue) MarshalJSON() ([]byte, error) {
	pb, err := tagKeyValueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TemporaryCredentials struct {

	// Wire name: 'aws_temp_credentials'
	AwsTempCredentials *AwsCredentials `json:"aws_temp_credentials,omitempty"`

	// Wire name: 'azure_aad'
	AzureAad *AzureActiveDirectoryToken `json:"azure_aad,omitempty"`
	// Server time when the credential will expire, in epoch milliseconds. The
	// API client is advised to cache the credential given this expiration time.
	// Wire name: 'expiration_time'
	ExpirationTime int64 `json:"expiration_time,omitempty"`

	// Wire name: 'gcp_oauth_token'
	GcpOauthToken *GcpOauthToken `json:"gcp_oauth_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TemporaryCredentials) EncodeValues(key string, v *url.Values) error {
	pb, err := temporaryCredentialsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Detailed status of an online table. Shown if the online table is in the
// ONLINE_TRIGGERED_UPDATE or the ONLINE_NO_PENDING_UPDATE state.
type TriggeredUpdateStatus struct {
	// The last source table Delta version that was synced to the online table.
	// Note that this Delta version may not be completely synced to the online
	// table yet.
	// Wire name: 'last_processed_commit_version'
	LastProcessedCommitVersion int64 `json:"last_processed_commit_version,omitempty"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the online table.
	// Wire name: 'timestamp'
	Timestamp string `json:"timestamp,omitempty"`
	// Progress of the active data synchronization pipeline.
	// Wire name: 'triggered_update_progress'
	TriggeredUpdateProgress *PipelineProgress `json:"triggered_update_progress,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TriggeredUpdateStatus) EncodeValues(key string, v *url.Values) error {
	pb, err := triggeredUpdateStatusToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UnassignRequest struct {
	// Query for the ID of the metastore to delete.
	MetastoreId string `json:"-" tf:"-"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`
}

func (st UnassignRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := unassignRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateCatalog struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'enable_predictive_optimization'
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode CatalogIsolationMode `json:"isolation_mode,omitempty"`
	// The name of the catalog.
	Name string `json:"-" tf:"-"`
	// New name for the catalog.
	// Wire name: 'new_name'
	NewName string `json:"new_name,omitempty"`
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string `json:"options,omitempty"`
	// Username of current owner of catalog.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string `json:"properties,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateCatalog) EncodeValues(key string, v *url.Values) error {
	pb, err := updateCatalogToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateCatalogWorkspaceBindingsResponse struct {
	// A list of workspace IDs
	// Wire name: 'workspaces'
	Workspaces []int64 `json:"workspaces,omitempty"`
}

func (st UpdateCatalogWorkspaceBindingsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updateCatalogWorkspaceBindingsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateCatalogWorkspaceBindingsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateCatalogWorkspaceBindingsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateCatalogWorkspaceBindingsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateCatalogWorkspaceBindingsResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateCatalogWorkspaceBindingsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateConnection struct {
	// Name of the connection.
	Name string `json:"-" tf:"-"`
	// New name for the connection.
	// Wire name: 'new_name'
	NewName string `json:"new_name,omitempty"`
	// A map of key-value properties attached to the securable.
	// Wire name: 'options'
	Options map[string]string `json:"options"`
	// Username of current owner of the connection.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateConnection) EncodeValues(key string, v *url.Values) error {
	pb, err := updateConnectionToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateCredentialRequest struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentity `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The Databricks managed GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount `json:"databricks_gcp_service_account,omitempty"`
	// Force an update even if there are dependent services (when purpose is
	// **SERVICE**) or dependent external locations and external tables (when
	// purpose is **STORAGE**).
	// Wire name: 'force'
	Force bool `json:"force,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Name of the credential.
	NameArg string `json:"-" tf:"-"`
	// New name of credential.
	// Wire name: 'new_name'
	NewName string `json:"new_name,omitempty"`
	// Username of current owner of credential.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	// Wire name: 'read_only'
	ReadOnly bool `json:"read_only,omitempty"`
	// Supply true to this argument to skip validation of the updated
	// credential.
	// Wire name: 'skip_validation'
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateCredentialRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateCredentialRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateExternalLineageRelationshipRequest struct {

	// Wire name: 'external_lineage_relationship'
	ExternalLineageRelationship UpdateRequestExternalLineage `json:"external_lineage_relationship"`
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
	UpdateMask string `json:"-" tf:"-"`
}

func (st UpdateExternalLineageRelationshipRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateExternalLineageRelationshipRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateExternalLineageRelationshipRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateExternalLineageRelationshipRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateExternalLineageRelationshipRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateExternalLineageRelationshipRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateExternalLineageRelationshipRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateExternalLocation struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Name of the storage credential used with this location.
	// Wire name: 'credential_name'
	CredentialName string `json:"credential_name,omitempty"`
	// Whether to enable file events on this external location.
	// Wire name: 'enable_file_events'
	EnableFileEvents bool `json:"enable_file_events,omitempty"`

	// Wire name: 'encryption_details'
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	// Indicates whether fallback mode is enabled for this external location.
	// When fallback mode is enabled, the access to the location falls back to
	// cluster credentials if UC credentials are not sufficient.
	// Wire name: 'fallback'
	Fallback bool `json:"fallback,omitempty"`
	// File event queue settings.
	// Wire name: 'file_event_queue'
	FileEventQueue *FileEventQueue `json:"file_event_queue,omitempty"`
	// Force update even if changing url invalidates dependent external tables
	// or mounts.
	// Wire name: 'force'
	Force bool `json:"force,omitempty"`

	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Name of the external location.
	Name string `json:"-" tf:"-"`
	// New name for the external location.
	// Wire name: 'new_name'
	NewName string `json:"new_name,omitempty"`
	// The owner of the external location.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Indicates whether the external location is read-only.
	// Wire name: 'read_only'
	ReadOnly bool `json:"read_only,omitempty"`
	// Skips validation of the storage credential associated with the external
	// location.
	// Wire name: 'skip_validation'
	SkipValidation bool `json:"skip_validation,omitempty"`
	// Path URL of the external location.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateExternalLocation) EncodeValues(key string, v *url.Values) error {
	pb, err := updateExternalLocationToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateExternalMetadataRequest struct {

	// Wire name: 'external_metadata'
	ExternalMetadata ExternalMetadata `json:"external_metadata"`
	// Name of the external metadata object.
	Name string `json:"-" tf:"-"`
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
	UpdateMask string `json:"-" tf:"-"`
}

func (st UpdateExternalMetadataRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateExternalMetadataRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateExternalMetadataRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateExternalMetadataRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateExternalMetadataRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateExternalMetadataRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateExternalMetadataRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateFunction struct {
	// The fully-qualified name of the function (of the form
	// __catalog_name__.__schema_name__.__function__name__).
	Name string `json:"-" tf:"-"`
	// Username of current owner of function.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateFunction) EncodeValues(key string, v *url.Values) error {
	pb, err := updateFunctionToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateMetastore struct {
	// The organization name of a Delta Sharing entity, to be used in
	// Databricks-to-Databricks Delta Sharing as the official name.
	// Wire name: 'delta_sharing_organization_name'
	DeltaSharingOrganizationName string `json:"delta_sharing_organization_name,omitempty"`
	// The lifetime of delta sharing recipient token in seconds.
	// Wire name: 'delta_sharing_recipient_token_lifetime_in_seconds'
	DeltaSharingRecipientTokenLifetimeInSeconds int64 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	// The scope of Delta Sharing enabled for the metastore.
	// Wire name: 'delta_sharing_scope'
	DeltaSharingScope DeltaSharingScopeEnum `json:"delta_sharing_scope,omitempty"`
	// Unique ID of the metastore.
	Id string `json:"-" tf:"-"`
	// New name for the metastore.
	// Wire name: 'new_name'
	NewName string `json:"new_name,omitempty"`
	// The owner of the metastore.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Privilege model version of the metastore, of the form `major.minor`
	// (e.g., `1.0`).
	// Wire name: 'privilege_model_version'
	PrivilegeModelVersion string `json:"privilege_model_version,omitempty"`
	// UUID of storage credential to access the metastore storage_root.
	// Wire name: 'storage_root_credential_id'
	StorageRootCredentialId string `json:"storage_root_credential_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateMetastore) EncodeValues(key string, v *url.Values) error {
	pb, err := updateMetastoreToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateMetastoreAssignment struct {
	// The name of the default catalog in the metastore. This field is
	// deprecated. Please use "Default Namespace API" to configure the default
	// catalog for a Databricks workspace.
	// Wire name: 'default_catalog_name'
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	// The unique ID of the metastore.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// A workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateMetastoreAssignment) EncodeValues(key string, v *url.Values) error {
	pb, err := updateMetastoreAssignmentToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateModelVersionRequest struct {
	// The comment attached to the model version
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The three-level (fully qualified) name of the model version
	FullName string `json:"-" tf:"-"`
	// The integer version number of the model version
	Version int `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateModelVersionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateModelVersionRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateMonitor struct {
	// Name of the baseline table from which drift metrics are computed from.
	// Columns in the monitored table should also be present in the baseline
	// table.
	// Wire name: 'baseline_table_name'
	BaselineTableName string `json:"baseline_table_name,omitempty"`
	// Custom metrics to compute on the monitored table. These can be aggregate
	// metrics, derived metrics (from already computed aggregate metrics), or
	// drift metrics (comparing metrics across time windows).
	// Wire name: 'custom_metrics'
	CustomMetrics []MonitorMetric `json:"custom_metrics,omitempty"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	// Wire name: 'dashboard_id'
	DashboardId string `json:"dashboard_id,omitempty"`
	// The data classification config for the monitor.
	// Wire name: 'data_classification_config'
	DataClassificationConfig *MonitorDataClassificationConfig `json:"data_classification_config,omitempty"`
	// Configuration for monitoring inference logs.
	// Wire name: 'inference_log'
	InferenceLog *MonitorInferenceLog `json:"inference_log,omitempty"`
	// The notification settings for the monitor.
	// Wire name: 'notifications'
	Notifications *MonitorNotifications `json:"notifications,omitempty"`
	// Schema where output metric tables are created.
	// Wire name: 'output_schema_name'
	OutputSchemaName string `json:"output_schema_name"`
	// The schedule for automatically updating and refreshing metric tables.
	// Wire name: 'schedule'
	Schedule *MonitorCronSchedule `json:"schedule,omitempty"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For high-cardinality
	// columns, only the top 100 unique values by frequency will generate
	// slices.
	// Wire name: 'slicing_exprs'
	SlicingExprs []string `json:"slicing_exprs,omitempty"`
	// Configuration for monitoring snapshot tables.
	// Wire name: 'snapshot'
	Snapshot *MonitorSnapshot `json:"snapshot,omitempty"`
	// Full name of the table.
	TableName string `json:"-" tf:"-"`
	// Configuration for monitoring time series tables.
	// Wire name: 'time_series'
	TimeSeries *MonitorTimeSeries `json:"time_series,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateMonitor) EncodeValues(key string, v *url.Values) error {
	pb, err := updateMonitorToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdatePermissions struct {
	// Array of permissions change objects.
	// Wire name: 'changes'
	Changes []PermissionsChange `json:"changes,omitempty"`
	// Full name of securable.
	FullName string `json:"-" tf:"-"`
	// Type of securable.
	SecurableType string `json:"-" tf:"-"`
}

func (st UpdatePermissions) EncodeValues(key string, v *url.Values) error {
	pb, err := updatePermissionsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdatePermissionsResponse struct {
	// The privileges assigned to each principal
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

func (st UpdatePermissionsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updatePermissionsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdatePermissionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updatePermissionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updatePermissionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdatePermissionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := updatePermissionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateRegisteredModelRequest struct {
	// The comment attached to the registered model
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The three-level (fully qualified) name of the registered model
	FullName string `json:"-" tf:"-"`
	// New name for the registered model.
	// Wire name: 'new_name'
	NewName string `json:"new_name,omitempty"`
	// The identifier of the user who owns the registered model
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateRegisteredModelRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateRegisteredModelRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateRequestExternalLineage struct {
	// List of column relationships between source and target objects.
	// Wire name: 'columns'
	Columns []ColumnRelationship `json:"columns,omitempty"`
	// Unique identifier of the external lineage relationship.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Key-value properties associated with the external lineage relationship.
	// Wire name: 'properties'
	Properties map[string]string `json:"properties,omitempty"`
	// Source object of the external lineage relationship.
	// Wire name: 'source'
	Source ExternalLineageObject `json:"source"`
	// Target object of the external lineage relationship.
	// Wire name: 'target'
	Target ExternalLineageObject `json:"target"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateRequestExternalLineage) EncodeValues(key string, v *url.Values) error {
	pb, err := updateRequestExternalLineageToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateRequestExternalLineage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateRequestExternalLineagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateRequestExternalLineageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateRequestExternalLineage) MarshalJSON() ([]byte, error) {
	pb, err := updateRequestExternalLineageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateSchema struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Whether predictive optimization should be enabled for this object and
	// objects under it.
	// Wire name: 'enable_predictive_optimization'
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	// Full name of the schema.
	FullName string `json:"-" tf:"-"`
	// New name for the schema.
	// Wire name: 'new_name'
	NewName string `json:"new_name,omitempty"`
	// Username of current owner of schema.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string `json:"properties,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateSchema) EncodeValues(key string, v *url.Values) error {
	pb, err := updateSchemaToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateStorageCredential struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRoleRequest `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentityResponse `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	// Wire name: 'cloudflare_api_token'
	CloudflareApiToken *CloudflareApiToken `json:"cloudflare_api_token,omitempty"`
	// Comment associated with the credential.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The Databricks managed GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `json:"databricks_gcp_service_account,omitempty"`
	// Force update even if there are dependent external locations or external
	// tables.
	// Wire name: 'force'
	Force bool `json:"force,omitempty"`
	// Whether the current securable is accessible from all workspaces or a
	// specific set of workspaces.
	// Wire name: 'isolation_mode'
	IsolationMode IsolationMode `json:"isolation_mode,omitempty"`
	// Name of the storage credential.
	Name string `json:"-" tf:"-"`
	// New name for the storage credential.
	// Wire name: 'new_name'
	NewName string `json:"new_name,omitempty"`
	// Username of current owner of credential.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Whether the credential is usable only for read operations. Only
	// applicable when purpose is **STORAGE**.
	// Wire name: 'read_only'
	ReadOnly bool `json:"read_only,omitempty"`
	// Supplying true to this argument skips validation of the updated
	// credential.
	// Wire name: 'skip_validation'
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateStorageCredential) EncodeValues(key string, v *url.Values) error {
	pb, err := updateStorageCredentialToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateTableRequest struct {
	// Full name of the table.
	FullName string `json:"-" tf:"-"`
	// Username of current owner of table.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateTableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateTableRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateVolumeRequestContent struct {
	// The comment attached to the volume
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The three-level (fully qualified) name of the volume
	Name string `json:"-" tf:"-"`
	// New name for the volume.
	// Wire name: 'new_name'
	NewName string `json:"new_name,omitempty"`
	// The identifier of the user who owns the volume
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateVolumeRequestContent) EncodeValues(key string, v *url.Values) error {
	pb, err := updateVolumeRequestContentToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateWorkspaceBindings struct {
	// A list of workspace IDs.
	// Wire name: 'assign_workspaces'
	AssignWorkspaces []int64 `json:"assign_workspaces,omitempty"`
	// The name of the catalog.
	Name string `json:"-" tf:"-"`
	// A list of workspace IDs.
	// Wire name: 'unassign_workspaces'
	UnassignWorkspaces []int64 `json:"unassign_workspaces,omitempty"`
}

func (st UpdateWorkspaceBindings) EncodeValues(key string, v *url.Values) error {
	pb, err := updateWorkspaceBindingsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateWorkspaceBindingsParameters struct {
	// List of workspace bindings.
	// Wire name: 'add'
	Add []WorkspaceBinding `json:"add,omitempty"`
	// List of workspace bindings.
	// Wire name: 'remove'
	Remove []WorkspaceBinding `json:"remove,omitempty"`
	// The name of the securable.
	SecurableName string `json:"-" tf:"-"`
	// The type of the securable to bind to a workspace (catalog,
	// storage_credential, credential, or external_location).
	SecurableType string `json:"-" tf:"-"`
}

func (st UpdateWorkspaceBindingsParameters) EncodeValues(key string, v *url.Values) error {
	pb, err := updateWorkspaceBindingsParametersToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// A list of workspace IDs that are bound to the securable
type UpdateWorkspaceBindingsResponse struct {
	// List of workspace bindings.
	// Wire name: 'bindings'
	Bindings []WorkspaceBinding `json:"bindings,omitempty"`
}

func (st UpdateWorkspaceBindingsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updateWorkspaceBindingsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateWorkspaceBindingsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateWorkspaceBindingsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateWorkspaceBindingsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateWorkspaceBindingsResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateWorkspaceBindingsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Next ID: 17
type ValidateCredentialRequest struct {

	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRole `json:"aws_iam_role,omitempty"`

	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentity `json:"azure_managed_identity,omitempty"`
	// Required. The name of an existing credential or long-lived cloud
	// credential to validate.
	// Wire name: 'credential_name'
	CredentialName string `json:"credential_name,omitempty"`

	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount `json:"databricks_gcp_service_account,omitempty"`
	// The name of an existing external location to validate. Only applicable
	// for storage credentials (purpose is **STORAGE**.)
	// Wire name: 'external_location_name'
	ExternalLocationName string `json:"external_location_name,omitempty"`
	// The purpose of the credential. This should only be used when the
	// credential is specified.
	// Wire name: 'purpose'
	Purpose CredentialPurpose `json:"purpose,omitempty"`
	// Whether the credential is only usable for read operations. Only
	// applicable for storage credentials (purpose is **STORAGE**.)
	// Wire name: 'read_only'
	ReadOnly bool `json:"read_only,omitempty"`
	// The external location url to validate. Only applicable when purpose is
	// **STORAGE**.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ValidateCredentialRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := validateCredentialRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ValidateCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage. Only
	// applicable for when purpose is **STORAGE**.
	// Wire name: 'isDir'
	IsDir bool `json:"isDir,omitempty"`
	// The results of the validation check.
	// Wire name: 'results'
	Results []CredentialValidationResult `json:"results,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ValidateCredentialResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := validateCredentialResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ValidateStorageCredential struct {
	// The AWS IAM role configuration.
	// Wire name: 'aws_iam_role'
	AwsIamRole *AwsIamRoleRequest `json:"aws_iam_role,omitempty"`
	// The Azure managed identity configuration.
	// Wire name: 'azure_managed_identity'
	AzureManagedIdentity *AzureManagedIdentityRequest `json:"azure_managed_identity,omitempty"`
	// The Azure service principal configuration.
	// Wire name: 'azure_service_principal'
	AzureServicePrincipal *AzureServicePrincipal `json:"azure_service_principal,omitempty"`
	// The Cloudflare API token configuration.
	// Wire name: 'cloudflare_api_token'
	CloudflareApiToken *CloudflareApiToken `json:"cloudflare_api_token,omitempty"`
	// The Databricks created GCP service account configuration.
	// Wire name: 'databricks_gcp_service_account'
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `json:"databricks_gcp_service_account,omitempty"`
	// The name of an existing external location to validate.
	// Wire name: 'external_location_name'
	ExternalLocationName string `json:"external_location_name,omitempty"`
	// Whether the storage credential is only usable for read operations.
	// Wire name: 'read_only'
	ReadOnly bool `json:"read_only,omitempty"`
	// Required. The name of an existing credential or long-lived cloud
	// credential to validate.
	// Wire name: 'storage_credential_name'
	StorageCredentialName string `json:"storage_credential_name,omitempty"`
	// The external location url to validate.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ValidateStorageCredential) EncodeValues(key string, v *url.Values) error {
	pb, err := validateStorageCredentialToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ValidateStorageCredentialResponse struct {
	// Whether the tested location is a directory in cloud storage.
	// Wire name: 'isDir'
	IsDir bool `json:"isDir,omitempty"`
	// The results of the validation check.
	// Wire name: 'results'
	Results []ValidationResult `json:"results,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ValidateStorageCredentialResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := validateStorageCredentialResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ValidationResult struct {
	// Error message would exist when the result does not equal to **PASS**.
	// Wire name: 'message'
	Message string `json:"message,omitempty"`
	// The operation tested.
	// Wire name: 'operation'
	Operation ValidationResultOperation `json:"operation,omitempty"`
	// The results of the tested operation.
	// Wire name: 'result'
	Result ValidationResultResult `json:"result,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ValidationResult) EncodeValues(key string, v *url.Values) error {
	pb, err := validationResultToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type VolumeInfo struct {
	// The AWS access point to use when accesing s3 for this external location.
	// Wire name: 'access_point'
	AccessPoint string `json:"access_point,omitempty"`
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the BROWSE privilege when include_browse is
	// enabled in the request.
	// Wire name: 'browse_only'
	BrowseOnly bool `json:"browse_only,omitempty"`
	// The name of the catalog where the schema and the volume are
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name,omitempty"`
	// The comment attached to the volume
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`

	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// The identifier of the user who created the volume
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`

	// Wire name: 'encryption_details'
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	// The three-level (fully qualified) name of the volume
	// Wire name: 'full_name'
	FullName string `json:"full_name,omitempty"`
	// The unique identifier of the metastore
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the volume
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The identifier of the user who owns the volume
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// The name of the schema where the volume is
	// Wire name: 'schema_name'
	SchemaName string `json:"schema_name,omitempty"`
	// The storage location on the cloud
	// Wire name: 'storage_location'
	StorageLocation string `json:"storage_location,omitempty"`

	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The identifier of the user who updated the volume last time
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`
	// The unique identifier of the volume
	// Wire name: 'volume_id'
	VolumeId string `json:"volume_id,omitempty"`

	// Wire name: 'volume_type'
	VolumeType VolumeType `json:"volume_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st VolumeInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := volumeInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type WorkspaceBinding struct {
	// One of READ_WRITE/READ_ONLY. Default is READ_WRITE.
	// Wire name: 'binding_type'
	BindingType WorkspaceBindingBindingType `json:"binding_type,omitempty"`
	// Required
	// Wire name: 'workspace_id'
	WorkspaceId int64 `json:"workspace_id"`
}

func (st WorkspaceBinding) EncodeValues(key string, v *url.Values) error {
	pb, err := workspaceBindingToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
