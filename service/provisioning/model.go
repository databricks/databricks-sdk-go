// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

type AwsCredentials struct {

	// Wire name: 'sts_role'
	StsRole *StsRole `json:"sts_role,omitempty"`
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

type AwsKeyInfo struct {
	// The AWS KMS key alias.
	// Wire name: 'key_alias'
	KeyAlias string `json:"key_alias,omitempty"`
	// The AWS KMS key's Amazon Resource Name (ARN).
	// Wire name: 'key_arn'
	KeyArn string `json:"key_arn"`
	// The AWS KMS key region.
	// Wire name: 'key_region'
	KeyRegion string `json:"key_region"`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. If you do not want to use this key for encrypting
	// EBS volumes, set to `false`.
	// Wire name: 'reuse_key_for_cluster_volumes'
	ReuseKeyForClusterVolumes bool `json:"reuse_key_for_cluster_volumes,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AwsKeyInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := awsKeyInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AwsKeyInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &awsKeyInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := awsKeyInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AwsKeyInfo) MarshalJSON() ([]byte, error) {
	pb, err := awsKeyInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AzureWorkspaceInfo struct {
	// Azure Resource Group name
	// Wire name: 'resource_group'
	ResourceGroup string `json:"resource_group,omitempty"`
	// Azure Subscription ID
	// Wire name: 'subscription_id'
	SubscriptionId string `json:"subscription_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AzureWorkspaceInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := azureWorkspaceInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AzureWorkspaceInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &azureWorkspaceInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := azureWorkspaceInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AzureWorkspaceInfo) MarshalJSON() ([]byte, error) {
	pb, err := azureWorkspaceInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The general workspace configurations that are specific to cloud providers.
type CloudResourceContainer struct {

	// Wire name: 'gcp'
	Gcp *CustomerFacingGcpCloudResourceContainer `json:"gcp,omitempty"`
}

func (st CloudResourceContainer) EncodeValues(key string, v *url.Values) error {
	pb, err := cloudResourceContainerToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CloudResourceContainer) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cloudResourceContainerPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cloudResourceContainerFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CloudResourceContainer) MarshalJSON() ([]byte, error) {
	pb, err := cloudResourceContainerToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateAwsKeyInfo struct {
	// The AWS KMS key alias.
	// Wire name: 'key_alias'
	KeyAlias string `json:"key_alias,omitempty"`
	// The AWS KMS key's Amazon Resource Name (ARN). Note that the key's AWS
	// region is inferred from the ARN.
	// Wire name: 'key_arn'
	KeyArn string `json:"key_arn"`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. To not use this key also for encrypting EBS volumes,
	// set this to `false`.
	// Wire name: 'reuse_key_for_cluster_volumes'
	ReuseKeyForClusterVolumes bool `json:"reuse_key_for_cluster_volumes,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateAwsKeyInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := createAwsKeyInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateAwsKeyInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createAwsKeyInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createAwsKeyInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateAwsKeyInfo) MarshalJSON() ([]byte, error) {
	pb, err := createAwsKeyInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateCredentialAwsCredentials struct {

	// Wire name: 'sts_role'
	StsRole *CreateCredentialStsRole `json:"sts_role,omitempty"`
}

func (st CreateCredentialAwsCredentials) EncodeValues(key string, v *url.Values) error {
	pb, err := createCredentialAwsCredentialsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateCredentialAwsCredentials) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCredentialAwsCredentialsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCredentialAwsCredentialsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCredentialAwsCredentials) MarshalJSON() ([]byte, error) {
	pb, err := createCredentialAwsCredentialsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateCredentialRequest struct {

	// Wire name: 'aws_credentials'
	AwsCredentials CreateCredentialAwsCredentials `json:"aws_credentials"`
	// The human-readable name of the credential configuration object.
	// Wire name: 'credentials_name'
	CredentialsName string `json:"credentials_name"`
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

type CreateCredentialStsRole struct {
	// The Amazon Resource Name (ARN) of the cross account role.
	// Wire name: 'role_arn'
	RoleArn string `json:"role_arn,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateCredentialStsRole) EncodeValues(key string, v *url.Values) error {
	pb, err := createCredentialStsRoleToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateCredentialStsRole) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCredentialStsRolePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCredentialStsRoleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCredentialStsRole) MarshalJSON() ([]byte, error) {
	pb, err := createCredentialStsRoleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateCustomerManagedKeyRequest struct {

	// Wire name: 'aws_key_info'
	AwsKeyInfo *CreateAwsKeyInfo `json:"aws_key_info,omitempty"`

	// Wire name: 'gcp_key_info'
	GcpKeyInfo *CreateGcpKeyInfo `json:"gcp_key_info,omitempty"`
	// The cases that the key can be used for.
	// Wire name: 'use_cases'
	UseCases []KeyUseCase `json:"use_cases"`
}

func (st CreateCustomerManagedKeyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createCustomerManagedKeyRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateCustomerManagedKeyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCustomerManagedKeyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCustomerManagedKeyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCustomerManagedKeyRequest) MarshalJSON() ([]byte, error) {
	pb, err := createCustomerManagedKeyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateGcpKeyInfo struct {
	// The GCP KMS key's resource name
	// Wire name: 'kms_key_id'
	KmsKeyId string `json:"kms_key_id"`
}

func (st CreateGcpKeyInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := createGcpKeyInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateGcpKeyInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createGcpKeyInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createGcpKeyInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateGcpKeyInfo) MarshalJSON() ([]byte, error) {
	pb, err := createGcpKeyInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateNetworkRequest struct {

	// Wire name: 'gcp_network_info'
	GcpNetworkInfo *GcpNetworkInfo `json:"gcp_network_info,omitempty"`
	// The human-readable name of the network configuration.
	// Wire name: 'network_name'
	NetworkName string `json:"network_name"`
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	// Wire name: 'security_group_ids'
	SecurityGroupIds []string `json:"security_group_ids,omitempty"`
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	// Wire name: 'subnet_ids'
	SubnetIds []string `json:"subnet_ids,omitempty"`

	// Wire name: 'vpc_endpoints'
	VpcEndpoints *NetworkVpcEndpoints `json:"vpc_endpoints,omitempty"`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	// Wire name: 'vpc_id'
	VpcId string `json:"vpc_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateNetworkRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createNetworkRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateNetworkRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createNetworkRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createNetworkRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateNetworkRequest) MarshalJSON() ([]byte, error) {
	pb, err := createNetworkRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateStorageConfigurationRequest struct {

	// Wire name: 'root_bucket_info'
	RootBucketInfo RootBucketInfo `json:"root_bucket_info"`
	// The human-readable name of the storage configuration.
	// Wire name: 'storage_configuration_name'
	StorageConfigurationName string `json:"storage_configuration_name"`
}

func (st CreateStorageConfigurationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createStorageConfigurationRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateStorageConfigurationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createStorageConfigurationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createStorageConfigurationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateStorageConfigurationRequest) MarshalJSON() ([]byte, error) {
	pb, err := createStorageConfigurationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateVpcEndpointRequest struct {
	// The ID of the VPC endpoint object in AWS.
	// Wire name: 'aws_vpc_endpoint_id'
	AwsVpcEndpointId string `json:"aws_vpc_endpoint_id,omitempty"`

	// Wire name: 'gcp_vpc_endpoint_info'
	GcpVpcEndpointInfo *GcpVpcEndpointInfo `json:"gcp_vpc_endpoint_info,omitempty"`
	// The AWS region in which this VPC endpoint object exists.
	// Wire name: 'region'
	Region string `json:"region,omitempty"`
	// The human-readable name of the storage configuration.
	// Wire name: 'vpc_endpoint_name'
	VpcEndpointName string `json:"vpc_endpoint_name"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateVpcEndpointRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createVpcEndpointRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateVpcEndpointRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createVpcEndpointRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createVpcEndpointRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateVpcEndpointRequest) MarshalJSON() ([]byte, error) {
	pb, err := createVpcEndpointRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateWorkspaceRequest struct {
	// The AWS region of the workspace's data plane.
	// Wire name: 'aws_region'
	AwsRegion string `json:"aws_region,omitempty"`
	// The cloud provider which the workspace uses. For Google Cloud workspaces,
	// always set this field to `gcp`.
	// Wire name: 'cloud'
	Cloud string `json:"cloud,omitempty"`

	// Wire name: 'cloud_resource_container'
	CloudResourceContainer *CloudResourceContainer `json:"cloud_resource_container,omitempty"`
	// ID of the workspace's credential configuration object.
	// Wire name: 'credentials_id'
	CredentialsId string `json:"credentials_id,omitempty"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	// Wire name: 'custom_tags'
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for the web application and REST APIs is
	// `<workspace-deployment-name>.cloud.databricks.com`. For example, if the
	// deployment name is `abcsales`, your workspace URL will be
	// `https://abcsales.cloud.databricks.com`. Hyphens are allowed. This
	// property supports only the set of characters that are allowed in a
	// subdomain.
	//
	// To set this value, you must have a deployment name prefix. Contact your
	// Databricks account team to add an account deployment name prefix to your
	// account.
	//
	// Workspace deployment names follow the account prefix and a hyphen. For
	// example, if your account's deployment prefix is `acme` and the workspace
	// deployment name is `workspace-1`, the JSON response for the
	// `deployment_name` field becomes `acme-workspace-1`. The workspace URL
	// would be `acme-workspace-1.cloud.databricks.com`.
	//
	// You can also set the `deployment_name` to the reserved keyword `EMPTY` if
	// you want the deployment name to only include the deployment prefix. For
	// example, if your account's deployment prefix is `acme` and the workspace
	// deployment name is `EMPTY`, the `deployment_name` becomes `acme` only and
	// the workspace URL is `acme.cloud.databricks.com`.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	//
	// If a new workspace omits this property, the server generates a unique
	// deployment name for you with the pattern `dbc-xxxxxxxx-xxxx`.
	// Wire name: 'deployment_name'
	DeploymentName string `json:"deployment_name,omitempty"`

	// Wire name: 'gcp_managed_network_config'
	GcpManagedNetworkConfig *GcpManagedNetworkConfig `json:"gcp_managed_network_config,omitempty"`

	// Wire name: 'gke_config'
	GkeConfig *GkeConfig `json:"gke_config,omitempty"`
	// Whether no public IP is enabled for the workspace.
	// Wire name: 'is_no_public_ip_enabled'
	IsNoPublicIpEnabled bool `json:"is_no_public_ip_enabled,omitempty"`
	// The Google Cloud region of the workspace data plane in your Google
	// account. For example, `us-east4`.
	// Wire name: 'location'
	Location string `json:"location,omitempty"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This is used to help protect and control access to the
	// workspace's notebooks, secrets, Databricks SQL queries, and query
	// history. The provided key configuration object property `use_cases` must
	// contain `MANAGED_SERVICES`.
	// Wire name: 'managed_services_customer_managed_key_id'
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`

	// Wire name: 'network_id'
	NetworkId string `json:"network_id,omitempty"`

	// Wire name: 'pricing_tier'
	PricingTier PricingTier `json:"pricing_tier,omitempty"`
	// ID of the workspace's private access settings object. Only used for
	// PrivateLink. This ID must be specified for customers using [AWS
	// PrivateLink] for either front-end (user-to-workspace connection),
	// back-end (data plane to control plane connection), or both connection
	// types.
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink].",
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// The ID of the workspace's storage configuration object.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The ID of the workspace's storage encryption key configuration object.
	// This is used to encrypt the workspace's root S3 bucket (root DBFS and
	// system data) and, optionally, cluster EBS volumes. The provided key
	// configuration object property `use_cases` must contain `STORAGE`.
	// Wire name: 'storage_customer_managed_key_id'
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// The workspace's human-readable name.
	// Wire name: 'workspace_name'
	WorkspaceName string `json:"workspace_name"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateWorkspaceRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createWorkspaceRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := createWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Credential struct {
	// The Databricks account ID that hosts the credential.
	// Wire name: 'account_id'
	AccountId string `json:"account_id,omitempty"`

	// Wire name: 'aws_credentials'
	AwsCredentials *AwsCredentials `json:"aws_credentials,omitempty"`
	// Time in epoch milliseconds when the credential was created.
	// Wire name: 'creation_time'
	CreationTime int64 `json:"creation_time,omitempty"`
	// Databricks credential configuration ID.
	// Wire name: 'credentials_id'
	CredentialsId string `json:"credentials_id,omitempty"`
	// The human-readable name of the credential configuration object.
	// Wire name: 'credentials_name'
	CredentialsName string `json:"credentials_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Credential) EncodeValues(key string, v *url.Values) error {
	pb, err := credentialToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Credential) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &credentialPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := credentialFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Credential) MarshalJSON() ([]byte, error) {
	pb, err := credentialToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The general workspace configurations that are specific to Google Cloud.
type CustomerFacingGcpCloudResourceContainer struct {
	// The Google Cloud project ID, which the workspace uses to instantiate
	// cloud resources for your workspace.
	// Wire name: 'project_id'
	ProjectId string `json:"project_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CustomerFacingGcpCloudResourceContainer) EncodeValues(key string, v *url.Values) error {
	pb, err := customerFacingGcpCloudResourceContainerToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CustomerFacingGcpCloudResourceContainer) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &customerFacingGcpCloudResourceContainerPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := customerFacingGcpCloudResourceContainerFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CustomerFacingGcpCloudResourceContainer) MarshalJSON() ([]byte, error) {
	pb, err := customerFacingGcpCloudResourceContainerToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CustomerManagedKey struct {
	// The Databricks account ID that holds the customer-managed key.
	// Wire name: 'account_id'
	AccountId string `json:"account_id,omitempty"`

	// Wire name: 'aws_key_info'
	AwsKeyInfo *AwsKeyInfo `json:"aws_key_info,omitempty"`
	// Time in epoch milliseconds when the customer key was created.
	// Wire name: 'creation_time'
	CreationTime int64 `json:"creation_time,omitempty"`
	// ID of the encryption key configuration object.
	// Wire name: 'customer_managed_key_id'
	CustomerManagedKeyId string `json:"customer_managed_key_id,omitempty"`

	// Wire name: 'gcp_key_info'
	GcpKeyInfo *GcpKeyInfo `json:"gcp_key_info,omitempty"`
	// The cases that the key can be used for.
	// Wire name: 'use_cases'
	UseCases []KeyUseCase `json:"use_cases,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CustomerManagedKey) EncodeValues(key string, v *url.Values) error {
	pb, err := customerManagedKeyToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CustomerManagedKey) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &customerManagedKeyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := customerManagedKeyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CustomerManagedKey) MarshalJSON() ([]byte, error) {
	pb, err := customerManagedKeyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteCredentialRequest struct {
	// Databricks Account API credential configuration ID
	CredentialsId string `json:"-" tf:"-"`
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

type DeleteEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId string `json:"-" tf:"-"`
}

func (st DeleteEncryptionKeyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteEncryptionKeyRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteEncryptionKeyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteEncryptionKeyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteEncryptionKeyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteEncryptionKeyRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteEncryptionKeyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteNetworkRequest struct {
	// Databricks Account API network configuration ID.
	NetworkId string `json:"-" tf:"-"`
}

func (st DeleteNetworkRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteNetworkRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteNetworkRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteNetworkRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteNetworkRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteNetworkRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteNetworkRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeletePrivateAccesRequest struct {
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string `json:"-" tf:"-"`
}

func (st DeletePrivateAccesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deletePrivateAccesRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeletePrivateAccesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deletePrivateAccesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deletePrivateAccesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeletePrivateAccesRequest) MarshalJSON() ([]byte, error) {
	pb, err := deletePrivateAccesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteStorageRequest struct {
	// Databricks Account API storage configuration ID.
	StorageConfigurationId string `json:"-" tf:"-"`
}

func (st DeleteStorageRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteStorageRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteStorageRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteStorageRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteStorageRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteStorageRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteStorageRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId string `json:"-" tf:"-"`
}

func (st DeleteVpcEndpointRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteVpcEndpointRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteVpcEndpointRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteVpcEndpointRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteVpcEndpointRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteVpcEndpointRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteVpcEndpointRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteWorkspaceRequest struct {
	// Workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`
}

func (st DeleteWorkspaceRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteWorkspaceRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// This enumeration represents the type of Databricks VPC [endpoint service]
// that was used when creating this VPC endpoint.
//
// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
type EndpointUseCase string

const EndpointUseCaseDataplaneRelayAccess EndpointUseCase = `DATAPLANE_RELAY_ACCESS`

const EndpointUseCaseWorkspaceAccess EndpointUseCase = `WORKSPACE_ACCESS`

// String representation for [fmt.Print]
func (f *EndpointUseCase) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointUseCase) Set(v string) error {
	switch v {
	case `DATAPLANE_RELAY_ACCESS`, `WORKSPACE_ACCESS`:
		*f = EndpointUseCase(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATAPLANE_RELAY_ACCESS", "WORKSPACE_ACCESS"`, v)
	}
}

// Values returns all possible values for EndpointUseCase.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointUseCase) Values() []EndpointUseCase {
	return []EndpointUseCase{
		EndpointUseCaseDataplaneRelayAccess,
		EndpointUseCaseWorkspaceAccess,
	}
}

// Type always returns EndpointUseCase to satisfy [pflag.Value] interface
func (f *EndpointUseCase) Type() string {
	return "EndpointUseCase"
}

// The AWS resource associated with this error: credentials, VPC, subnet,
// security group, or network ACL.
type ErrorType string

const ErrorTypeCredentials ErrorType = `credentials`

const ErrorTypeNetworkAcl ErrorType = `networkAcl`

const ErrorTypeSecurityGroup ErrorType = `securityGroup`

const ErrorTypeSubnet ErrorType = `subnet`

const ErrorTypeVpc ErrorType = `vpc`

// String representation for [fmt.Print]
func (f *ErrorType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ErrorType) Set(v string) error {
	switch v {
	case `credentials`, `networkAcl`, `securityGroup`, `subnet`, `vpc`:
		*f = ErrorType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "credentials", "networkAcl", "securityGroup", "subnet", "vpc"`, v)
	}
}

// Values returns all possible values for ErrorType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ErrorType) Values() []ErrorType {
	return []ErrorType{
		ErrorTypeCredentials,
		ErrorTypeNetworkAcl,
		ErrorTypeSecurityGroup,
		ErrorTypeSubnet,
		ErrorTypeVpc,
	}
}

// Type always returns ErrorType to satisfy [pflag.Value] interface
func (f *ErrorType) Type() string {
	return "ErrorType"
}

type ExternalCustomerInfo struct {
	// Email of the authoritative user.
	// Wire name: 'authoritative_user_email'
	AuthoritativeUserEmail string `json:"authoritative_user_email,omitempty"`
	// The authoritative user full name.
	// Wire name: 'authoritative_user_full_name'
	AuthoritativeUserFullName string `json:"authoritative_user_full_name,omitempty"`
	// The legal entity name for the external workspace
	// Wire name: 'customer_name'
	CustomerName string `json:"customer_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalCustomerInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := externalCustomerInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalCustomerInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalCustomerInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalCustomerInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalCustomerInfo) MarshalJSON() ([]byte, error) {
	pb, err := externalCustomerInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GcpKeyInfo struct {
	// The GCP KMS key's resource name
	// Wire name: 'kms_key_id'
	KmsKeyId string `json:"kms_key_id"`
}

func (st GcpKeyInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := gcpKeyInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GcpKeyInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gcpKeyInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gcpKeyInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GcpKeyInfo) MarshalJSON() ([]byte, error) {
	pb, err := gcpKeyInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The network settings for the workspace. The configurations are only for
// Databricks-managed VPCs. It is ignored if you specify a customer-managed VPC
// in the `network_id` field.", All the IP range configurations must be mutually
// exclusive. An attempt to create a workspace fails if Databricks detects an IP
// range overlap.
//
// Specify custom IP ranges in CIDR format. The IP ranges for these fields must
// not overlap, and all IP addresses must be entirely within the following
// ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`, `192.168.0.0/16`, and
// `240.0.0.0/4`.
//
// The sizes of these IP ranges affect the maximum number of nodes for the
// workspace.
//
// **Important**: Confirm the IP ranges used by your Databricks workspace before
// creating the workspace. You cannot change them after your workspace is
// deployed. If the IP address ranges for your Databricks are too small, IP
// exhaustion can occur, causing your Databricks jobs to fail. To determine the
// address range sizes that you need, Databricks provides a calculator as a
// Microsoft Excel spreadsheet. See [calculate subnet sizes for a new
// workspace].
//
// [calculate subnet sizes for a new workspace]: https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html
type GcpManagedNetworkConfig struct {
	// The IP range from which to allocate GKE cluster pods. No bigger than `/9`
	// and no smaller than `/21`.
	// Wire name: 'gke_cluster_pod_ip_range'
	GkeClusterPodIpRange string `json:"gke_cluster_pod_ip_range,omitempty"`
	// The IP range from which to allocate GKE cluster services. No bigger than
	// `/16` and no smaller than `/27`.
	// Wire name: 'gke_cluster_service_ip_range'
	GkeClusterServiceIpRange string `json:"gke_cluster_service_ip_range,omitempty"`
	// The IP range from which to allocate GKE cluster nodes. No bigger than
	// `/9` and no smaller than `/29`.
	// Wire name: 'subnet_cidr'
	SubnetCidr string `json:"subnet_cidr,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GcpManagedNetworkConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := gcpManagedNetworkConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GcpManagedNetworkConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gcpManagedNetworkConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gcpManagedNetworkConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GcpManagedNetworkConfig) MarshalJSON() ([]byte, error) {
	pb, err := gcpManagedNetworkConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The Google Cloud specific information for this network (for example, the VPC
// ID, subnet ID, and secondary IP ranges).
type GcpNetworkInfo struct {
	// The Google Cloud project ID of the VPC network.
	// Wire name: 'network_project_id'
	NetworkProjectId string `json:"network_project_id"`
	// The name of the secondary IP range for pods. A Databricks-managed GKE
	// cluster uses this IP range for its pods. This secondary IP range can be
	// used by only one workspace.
	// Wire name: 'pod_ip_range_name'
	PodIpRangeName string `json:"pod_ip_range_name"`
	// The name of the secondary IP range for services. A Databricks-managed GKE
	// cluster uses this IP range for its services. This secondary IP range can
	// be used by only one workspace.
	// Wire name: 'service_ip_range_name'
	ServiceIpRangeName string `json:"service_ip_range_name"`
	// The ID of the subnet associated with this network.
	// Wire name: 'subnet_id'
	SubnetId string `json:"subnet_id"`
	// The Google Cloud region of the workspace data plane (for example,
	// `us-east4`).
	// Wire name: 'subnet_region'
	SubnetRegion string `json:"subnet_region"`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	// Wire name: 'vpc_id'
	VpcId string `json:"vpc_id"`
}

func (st GcpNetworkInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := gcpNetworkInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GcpNetworkInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gcpNetworkInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gcpNetworkInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GcpNetworkInfo) MarshalJSON() ([]byte, error) {
	pb, err := gcpNetworkInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The Google Cloud specific information for this Private Service Connect
// endpoint.
type GcpVpcEndpointInfo struct {
	// Region of the PSC endpoint.
	// Wire name: 'endpoint_region'
	EndpointRegion string `json:"endpoint_region"`
	// The Google Cloud project ID of the VPC network where the PSC connection
	// resides.
	// Wire name: 'project_id'
	ProjectId string `json:"project_id"`
	// The unique ID of this PSC connection.
	// Wire name: 'psc_connection_id'
	PscConnectionId string `json:"psc_connection_id,omitempty"`
	// The name of the PSC endpoint in the Google Cloud project.
	// Wire name: 'psc_endpoint_name'
	PscEndpointName string `json:"psc_endpoint_name"`
	// The service attachment this PSC connection connects to.
	// Wire name: 'service_attachment_id'
	ServiceAttachmentId string `json:"service_attachment_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GcpVpcEndpointInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := gcpVpcEndpointInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GcpVpcEndpointInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gcpVpcEndpointInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gcpVpcEndpointInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GcpVpcEndpointInfo) MarshalJSON() ([]byte, error) {
	pb, err := gcpVpcEndpointInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetCredentialRequest struct {
	// Databricks Account API credential configuration ID
	CredentialsId string `json:"-" tf:"-"`
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

type GetEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId string `json:"-" tf:"-"`
}

func (st GetEncryptionKeyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getEncryptionKeyRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetEncryptionKeyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getEncryptionKeyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getEncryptionKeyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetEncryptionKeyRequest) MarshalJSON() ([]byte, error) {
	pb, err := getEncryptionKeyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetNetworkRequest struct {
	// Databricks Account API network configuration ID.
	NetworkId string `json:"-" tf:"-"`
}

func (st GetNetworkRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getNetworkRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetNetworkRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getNetworkRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getNetworkRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetNetworkRequest) MarshalJSON() ([]byte, error) {
	pb, err := getNetworkRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetPrivateAccesRequest struct {
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string `json:"-" tf:"-"`
}

func (st GetPrivateAccesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getPrivateAccesRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetPrivateAccesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPrivateAccesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPrivateAccesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPrivateAccesRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPrivateAccesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetStorageRequest struct {
	// Databricks Account API storage configuration ID.
	StorageConfigurationId string `json:"-" tf:"-"`
}

func (st GetStorageRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getStorageRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetStorageRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getStorageRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getStorageRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetStorageRequest) MarshalJSON() ([]byte, error) {
	pb, err := getStorageRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId string `json:"-" tf:"-"`
}

func (st GetVpcEndpointRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getVpcEndpointRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetVpcEndpointRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getVpcEndpointRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getVpcEndpointRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetVpcEndpointRequest) MarshalJSON() ([]byte, error) {
	pb, err := getVpcEndpointRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetWorkspaceRequest struct {
	// Workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`
}

func (st GetWorkspaceRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getWorkspaceRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := getWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The configurations for the GKE cluster of a Databricks workspace.
type GkeConfig struct {
	// Specifies the network connectivity types for the GKE nodes and the GKE
	// master network.
	//
	// Set to `PRIVATE_NODE_PUBLIC_MASTER` for a private GKE cluster for the
	// workspace. The GKE nodes will not have public IPs.
	//
	// Set to `PUBLIC_NODE_PUBLIC_MASTER` for a public GKE cluster. The nodes of
	// a public GKE cluster have public IP addresses.
	// Wire name: 'connectivity_type'
	ConnectivityType GkeConfigConnectivityType `json:"connectivity_type,omitempty"`
	// The IP range from which to allocate GKE cluster master resources. This
	// field will be ignored if GKE private cluster is not enabled.
	//
	// It must be exactly as big as `/28`.
	// Wire name: 'master_ip_range'
	MasterIpRange string `json:"master_ip_range,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GkeConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := gkeConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GkeConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gkeConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gkeConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GkeConfig) MarshalJSON() ([]byte, error) {
	pb, err := gkeConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Specifies the network connectivity types for the GKE nodes and the GKE master
// network.
//
// Set to `PRIVATE_NODE_PUBLIC_MASTER` for a private GKE cluster for the
// workspace. The GKE nodes will not have public IPs.
//
// Set to `PUBLIC_NODE_PUBLIC_MASTER` for a public GKE cluster. The nodes of a
// public GKE cluster have public IP addresses.
type GkeConfigConnectivityType string

const GkeConfigConnectivityTypePrivateNodePublicMaster GkeConfigConnectivityType = `PRIVATE_NODE_PUBLIC_MASTER`

const GkeConfigConnectivityTypePublicNodePublicMaster GkeConfigConnectivityType = `PUBLIC_NODE_PUBLIC_MASTER`

// String representation for [fmt.Print]
func (f *GkeConfigConnectivityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GkeConfigConnectivityType) Set(v string) error {
	switch v {
	case `PRIVATE_NODE_PUBLIC_MASTER`, `PUBLIC_NODE_PUBLIC_MASTER`:
		*f = GkeConfigConnectivityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PRIVATE_NODE_PUBLIC_MASTER", "PUBLIC_NODE_PUBLIC_MASTER"`, v)
	}
}

// Values returns all possible values for GkeConfigConnectivityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *GkeConfigConnectivityType) Values() []GkeConfigConnectivityType {
	return []GkeConfigConnectivityType{
		GkeConfigConnectivityTypePrivateNodePublicMaster,
		GkeConfigConnectivityTypePublicNodePublicMaster,
	}
}

// Type always returns GkeConfigConnectivityType to satisfy [pflag.Value] interface
func (f *GkeConfigConnectivityType) Type() string {
	return "GkeConfigConnectivityType"
}

// Possible values are: * `MANAGED_SERVICES`: Encrypts notebook and secret data
// in the control plane * `STORAGE`: Encrypts the workspace's root S3 bucket
// (root DBFS and system data) and, optionally, cluster EBS volumes.
type KeyUseCase string

// Encrypts notebook and secret data in the control plane
const KeyUseCaseManagedServices KeyUseCase = `MANAGED_SERVICES`

// Encrypts the workspace's root S3 bucket (root DBFS and system data) and,
// optionally, cluster EBS volumes.
const KeyUseCaseStorage KeyUseCase = `STORAGE`

// String representation for [fmt.Print]
func (f *KeyUseCase) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *KeyUseCase) Set(v string) error {
	switch v {
	case `MANAGED_SERVICES`, `STORAGE`:
		*f = KeyUseCase(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MANAGED_SERVICES", "STORAGE"`, v)
	}
}

// Values returns all possible values for KeyUseCase.
//
// There is no guarantee on the order of the values in the slice.
func (f *KeyUseCase) Values() []KeyUseCase {
	return []KeyUseCase{
		KeyUseCaseManagedServices,
		KeyUseCaseStorage,
	}
}

// Type always returns KeyUseCase to satisfy [pflag.Value] interface
func (f *KeyUseCase) Type() string {
	return "KeyUseCase"
}

type Network struct {
	// The Databricks account ID associated with this network configuration.
	// Wire name: 'account_id'
	AccountId string `json:"account_id,omitempty"`
	// Time in epoch milliseconds when the network was created.
	// Wire name: 'creation_time'
	CreationTime int64 `json:"creation_time,omitempty"`
	// Array of error messages about the network configuration.
	// Wire name: 'error_messages'
	ErrorMessages []NetworkHealth `json:"error_messages,omitempty"`

	// Wire name: 'gcp_network_info'
	GcpNetworkInfo *GcpNetworkInfo `json:"gcp_network_info,omitempty"`
	// The Databricks network configuration ID.
	// Wire name: 'network_id'
	NetworkId string `json:"network_id,omitempty"`
	// The human-readable name of the network configuration.
	// Wire name: 'network_name'
	NetworkName string `json:"network_name,omitempty"`

	// Wire name: 'security_group_ids'
	SecurityGroupIds []string `json:"security_group_ids,omitempty"`

	// Wire name: 'subnet_ids'
	SubnetIds []string `json:"subnet_ids,omitempty"`

	// Wire name: 'vpc_endpoints'
	VpcEndpoints *NetworkVpcEndpoints `json:"vpc_endpoints,omitempty"`
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	// Wire name: 'vpc_id'
	VpcId string `json:"vpc_id,omitempty"`

	// Wire name: 'vpc_status'
	VpcStatus VpcStatus `json:"vpc_status,omitempty"`
	// Array of warning messages about the network configuration.
	// Wire name: 'warning_messages'
	WarningMessages []NetworkWarning `json:"warning_messages,omitempty"`
	// Workspace ID associated with this network configuration.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Network) EncodeValues(key string, v *url.Values) error {
	pb, err := networkToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Network) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &networkPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := networkFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Network) MarshalJSON() ([]byte, error) {
	pb, err := networkToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type NetworkHealth struct {
	// Details of the error.
	// Wire name: 'error_message'
	ErrorMessage string `json:"error_message,omitempty"`

	// Wire name: 'error_type'
	ErrorType ErrorType `json:"error_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st NetworkHealth) EncodeValues(key string, v *url.Values) error {
	pb, err := networkHealthToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *NetworkHealth) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &networkHealthPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := networkHealthFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NetworkHealth) MarshalJSON() ([]byte, error) {
	pb, err := networkHealthToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// If specified, contains the VPC endpoints used to allow cluster communication
// from this VPC over [AWS PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
type NetworkVpcEndpoints struct {
	// The VPC endpoint ID used by this network to access the Databricks secure
	// cluster connectivity relay.
	// Wire name: 'dataplane_relay'
	DataplaneRelay []string `json:"dataplane_relay"`
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	// Wire name: 'rest_api'
	RestApi []string `json:"rest_api"`
}

func (st NetworkVpcEndpoints) EncodeValues(key string, v *url.Values) error {
	pb, err := networkVpcEndpointsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *NetworkVpcEndpoints) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &networkVpcEndpointsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := networkVpcEndpointsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NetworkVpcEndpoints) MarshalJSON() ([]byte, error) {
	pb, err := networkVpcEndpointsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type NetworkWarning struct {
	// Details of the warning.
	// Wire name: 'warning_message'
	WarningMessage string `json:"warning_message,omitempty"`

	// Wire name: 'warning_type'
	WarningType WarningType `json:"warning_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st NetworkWarning) EncodeValues(key string, v *url.Values) error {
	pb, err := networkWarningToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *NetworkWarning) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &networkWarningPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := networkWarningFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NetworkWarning) MarshalJSON() ([]byte, error) {
	pb, err := networkWarningToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The pricing tier of the workspace. For pricing tier information, see [AWS
// Pricing].
//
// [AWS Pricing]: https://databricks.com/product/aws-pricing
type PricingTier string

const PricingTierCommunityEdition PricingTier = `COMMUNITY_EDITION`

const PricingTierDedicated PricingTier = `DEDICATED`

const PricingTierEnterprise PricingTier = `ENTERPRISE`

const PricingTierPremium PricingTier = `PREMIUM`

const PricingTierStandard PricingTier = `STANDARD`

const PricingTierUnknown PricingTier = `UNKNOWN`

// String representation for [fmt.Print]
func (f *PricingTier) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PricingTier) Set(v string) error {
	switch v {
	case `COMMUNITY_EDITION`, `DEDICATED`, `ENTERPRISE`, `PREMIUM`, `STANDARD`, `UNKNOWN`:
		*f = PricingTier(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COMMUNITY_EDITION", "DEDICATED", "ENTERPRISE", "PREMIUM", "STANDARD", "UNKNOWN"`, v)
	}
}

// Values returns all possible values for PricingTier.
//
// There is no guarantee on the order of the values in the slice.
func (f *PricingTier) Values() []PricingTier {
	return []PricingTier{
		PricingTierCommunityEdition,
		PricingTierDedicated,
		PricingTierEnterprise,
		PricingTierPremium,
		PricingTierStandard,
		PricingTierUnknown,
	}
}

// Type always returns PricingTier to satisfy [pflag.Value] interface
func (f *PricingTier) Type() string {
	return "PricingTier"
}

// The private access level controls which VPC endpoints can connect to the UI
// or API of any workspace that attaches this private access settings object. *
// `ACCOUNT` level access (the default) allows only VPC endpoints that are
// registered in your Databricks account connect to your workspace. * `ENDPOINT`
// level access allows only specified VPC endpoints connect to your workspace.
// For details, see `allowed_vpc_endpoint_ids`.
type PrivateAccessLevel string

const PrivateAccessLevelAccount PrivateAccessLevel = `ACCOUNT`

const PrivateAccessLevelEndpoint PrivateAccessLevel = `ENDPOINT`

// String representation for [fmt.Print]
func (f *PrivateAccessLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PrivateAccessLevel) Set(v string) error {
	switch v {
	case `ACCOUNT`, `ENDPOINT`:
		*f = PrivateAccessLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACCOUNT", "ENDPOINT"`, v)
	}
}

// Values returns all possible values for PrivateAccessLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *PrivateAccessLevel) Values() []PrivateAccessLevel {
	return []PrivateAccessLevel{
		PrivateAccessLevelAccount,
		PrivateAccessLevelEndpoint,
	}
}

// Type always returns PrivateAccessLevel to satisfy [pflag.Value] interface
func (f *PrivateAccessLevel) Type() string {
	return "PrivateAccessLevel"
}

type PrivateAccessSettings struct {
	// The Databricks account ID that hosts the credential.
	// Wire name: 'account_id'
	AccountId string `json:"account_id,omitempty"`
	// An array of Databricks VPC endpoint IDs.
	// Wire name: 'allowed_vpc_endpoint_ids'
	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`

	// Wire name: 'private_access_level'
	PrivateAccessLevel PrivateAccessLevel `json:"private_access_level,omitempty"`
	// Databricks private access settings ID.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// The human-readable name of the private access settings object.
	// Wire name: 'private_access_settings_name'
	PrivateAccessSettingsName string `json:"private_access_settings_name,omitempty"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	// Wire name: 'public_access_enabled'
	PublicAccessEnabled bool `json:"public_access_enabled,omitempty"`
	// The cloud region for workspaces attached to this private access settings
	// object.
	// Wire name: 'region'
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PrivateAccessSettings) EncodeValues(key string, v *url.Values) error {
	pb, err := privateAccessSettingsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PrivateAccessSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &privateAccessSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := privateAccessSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PrivateAccessSettings) MarshalJSON() ([]byte, error) {
	pb, err := privateAccessSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Root S3 bucket information.
type RootBucketInfo struct {
	// The name of the S3 bucket.
	// Wire name: 'bucket_name'
	BucketName string `json:"bucket_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RootBucketInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := rootBucketInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *RootBucketInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &rootBucketInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := rootBucketInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RootBucketInfo) MarshalJSON() ([]byte, error) {
	pb, err := rootBucketInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StorageConfiguration struct {
	// The Databricks account ID that hosts the credential.
	// Wire name: 'account_id'
	AccountId string `json:"account_id,omitempty"`
	// Time in epoch milliseconds when the storage configuration was created.
	// Wire name: 'creation_time'
	CreationTime int64 `json:"creation_time,omitempty"`

	// Wire name: 'root_bucket_info'
	RootBucketInfo *RootBucketInfo `json:"root_bucket_info,omitempty"`
	// Databricks storage configuration ID.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The human-readable name of the storage configuration.
	// Wire name: 'storage_configuration_name'
	StorageConfigurationName string `json:"storage_configuration_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st StorageConfiguration) EncodeValues(key string, v *url.Values) error {
	pb, err := storageConfigurationToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *StorageConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &storageConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := storageConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StorageConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := storageConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StsRole struct {
	// The external ID that needs to be trusted by the cross-account role. This
	// is always your Databricks account ID.
	// Wire name: 'external_id'
	ExternalId string `json:"external_id,omitempty"`
	// The Amazon Resource Name (ARN) of the cross account role.
	// Wire name: 'role_arn'
	RoleArn string `json:"role_arn,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st StsRole) EncodeValues(key string, v *url.Values) error {
	pb, err := stsRoleToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *StsRole) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &stsRolePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := stsRoleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StsRole) MarshalJSON() ([]byte, error) {
	pb, err := stsRoleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateWorkspaceRequest struct {
	// The AWS region of the workspace's data plane (for example, `us-west-2`).
	// This parameter is available only for updating failed workspaces.
	// Wire name: 'aws_region'
	AwsRegion string `json:"aws_region,omitempty"`
	// ID of the workspace's credential configuration object. This parameter is
	// available for updating both failed and running workspaces.
	// Wire name: 'credentials_id'
	CredentialsId string `json:"credentials_id,omitempty"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	// Wire name: 'custom_tags'
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This parameter is available only for updating failed workspaces.
	// Wire name: 'managed_services_customer_managed_key_id'
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`

	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// The ID of the workspace's network configuration object. Used only if you
	// already use a customer-managed VPC. For failed workspaces only, you can
	// switch from a Databricks-managed VPC to a customer-managed VPC by
	// updating the workspace to add a network configuration ID.
	// Wire name: 'network_id'
	NetworkId string `json:"network_id,omitempty"`
	// The ID of the workspace's private access settings configuration object.
	// This parameter is available only for updating failed workspaces.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// The ID of the workspace's storage configuration object. This parameter is
	// available only for updating failed workspaces.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The ID of the key configuration object for workspace storage. This
	// parameter is available for updating both failed and running workspaces.
	// Wire name: 'storage_customer_managed_key_id'
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateWorkspaceRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateWorkspaceRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpsertPrivateAccessSettingsRequest struct {
	// An array of Databricks VPC endpoint IDs. This is the Databricks ID that
	// is returned when registering the VPC endpoint configuration in your
	// Databricks account. This is not the ID of the VPC endpoint in AWS.
	//
	// Only used when `private_access_level` is set to `ENDPOINT`. This is an
	// allow list of VPC endpoints that in your account that can connect to your
	// workspace over AWS PrivateLink.
	//
	// If hybrid access to your workspace is enabled by setting
	// `public_access_enabled` to `true`, this control only works for
	// PrivateLink connections. To control how your workspace is accessed via
	// public internet, see [IP access lists].
	//
	// [IP access lists]: https://docs.databricks.com/security/network/ip-access-list.html
	// Wire name: 'allowed_vpc_endpoint_ids'
	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`

	// Wire name: 'private_access_level'
	PrivateAccessLevel PrivateAccessLevel `json:"private_access_level,omitempty"`
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string `json:"-" tf:"-"`
	// The human-readable name of the private access settings object.
	// Wire name: 'private_access_settings_name'
	PrivateAccessSettingsName string `json:"private_access_settings_name"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	// Wire name: 'public_access_enabled'
	PublicAccessEnabled bool `json:"public_access_enabled,omitempty"`
	// The cloud region for workspaces associated with this private access
	// settings object.
	// Wire name: 'region'
	Region string `json:"region"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpsertPrivateAccessSettingsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := upsertPrivateAccessSettingsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpsertPrivateAccessSettingsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &upsertPrivateAccessSettingsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := upsertPrivateAccessSettingsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpsertPrivateAccessSettingsRequest) MarshalJSON() ([]byte, error) {
	pb, err := upsertPrivateAccessSettingsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type VpcEndpoint struct {
	// The Databricks account ID that hosts the VPC endpoint configuration.
	// Wire name: 'account_id'
	AccountId string `json:"account_id,omitempty"`
	// The AWS Account in which the VPC endpoint object exists.
	// Wire name: 'aws_account_id'
	AwsAccountId string `json:"aws_account_id,omitempty"`
	// The ID of the Databricks [endpoint service] that this VPC endpoint is
	// connected to. For a list of endpoint service IDs for each supported AWS
	// region, see the [Databricks PrivateLink documentation].
	//
	// [Databricks PrivateLink documentation]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	// Wire name: 'aws_endpoint_service_id'
	AwsEndpointServiceId string `json:"aws_endpoint_service_id,omitempty"`
	// The ID of the VPC endpoint object in AWS.
	// Wire name: 'aws_vpc_endpoint_id'
	AwsVpcEndpointId string `json:"aws_vpc_endpoint_id,omitempty"`

	// Wire name: 'gcp_vpc_endpoint_info'
	GcpVpcEndpointInfo *GcpVpcEndpointInfo `json:"gcp_vpc_endpoint_info,omitempty"`
	// The AWS region in which this VPC endpoint object exists.
	// Wire name: 'region'
	Region string `json:"region,omitempty"`
	// The current state (such as `available` or `rejected`) of the VPC
	// endpoint. Derived from AWS. For the full set of values, see [AWS
	// DescribeVpcEndpoint documentation].
	//
	// [AWS DescribeVpcEndpoint documentation]: https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-vpc-endpoints.html
	// Wire name: 'state'
	State string `json:"state,omitempty"`

	// Wire name: 'use_case'
	UseCase EndpointUseCase `json:"use_case,omitempty"`
	// Databricks VPC endpoint ID. This is the Databricks-specific name of the
	// VPC endpoint. Do not confuse this with the `aws_vpc_endpoint_id`, which
	// is the ID within AWS of the VPC endpoint.
	// Wire name: 'vpc_endpoint_id'
	VpcEndpointId string `json:"vpc_endpoint_id,omitempty"`
	// The human-readable name of the storage configuration.
	// Wire name: 'vpc_endpoint_name'
	VpcEndpointName string `json:"vpc_endpoint_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st VpcEndpoint) EncodeValues(key string, v *url.Values) error {
	pb, err := vpcEndpointToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *VpcEndpoint) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &vpcEndpointPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := vpcEndpointFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st VpcEndpoint) MarshalJSON() ([]byte, error) {
	pb, err := vpcEndpointToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The status of this network configuration object in terms of its use in a
// workspace: * `UNATTACHED`: Unattached. * `VALID`: Valid. * `BROKEN`: Broken.
// * `WARNED`: Warned.
type VpcStatus string

// Broken.
const VpcStatusBroken VpcStatus = `BROKEN`

// Unattached.
const VpcStatusUnattached VpcStatus = `UNATTACHED`

// Valid.
const VpcStatusValid VpcStatus = `VALID`

// Warned.
const VpcStatusWarned VpcStatus = `WARNED`

// String representation for [fmt.Print]
func (f *VpcStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *VpcStatus) Set(v string) error {
	switch v {
	case `BROKEN`, `UNATTACHED`, `VALID`, `WARNED`:
		*f = VpcStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BROKEN", "UNATTACHED", "VALID", "WARNED"`, v)
	}
}

// Values returns all possible values for VpcStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *VpcStatus) Values() []VpcStatus {
	return []VpcStatus{
		VpcStatusBroken,
		VpcStatusUnattached,
		VpcStatusValid,
		VpcStatusWarned,
	}
}

// Type always returns VpcStatus to satisfy [pflag.Value] interface
func (f *VpcStatus) Type() string {
	return "VpcStatus"
}

// The AWS resource associated with this warning: a subnet or a security group.
type WarningType string

const WarningTypeSecurityGroup WarningType = `securityGroup`

const WarningTypeSubnet WarningType = `subnet`

// String representation for [fmt.Print]
func (f *WarningType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WarningType) Set(v string) error {
	switch v {
	case `securityGroup`, `subnet`:
		*f = WarningType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "securityGroup", "subnet"`, v)
	}
}

// Values returns all possible values for WarningType.
//
// There is no guarantee on the order of the values in the slice.
func (f *WarningType) Values() []WarningType {
	return []WarningType{
		WarningTypeSecurityGroup,
		WarningTypeSubnet,
	}
}

// Type always returns WarningType to satisfy [pflag.Value] interface
func (f *WarningType) Type() string {
	return "WarningType"
}

type Workspace struct {
	// Databricks account ID.
	// Wire name: 'account_id'
	AccountId string `json:"account_id,omitempty"`
	// The AWS region of the workspace data plane (for example, `us-west-2`).
	// Wire name: 'aws_region'
	AwsRegion string `json:"aws_region,omitempty"`

	// Wire name: 'azure_workspace_info'
	AzureWorkspaceInfo *AzureWorkspaceInfo `json:"azure_workspace_info,omitempty"`
	// The cloud name. This field always has the value `gcp`.
	// Wire name: 'cloud'
	Cloud string `json:"cloud,omitempty"`

	// Wire name: 'cloud_resource_container'
	CloudResourceContainer *CloudResourceContainer `json:"cloud_resource_container,omitempty"`
	// Time in epoch milliseconds when the workspace was created.
	// Wire name: 'creation_time'
	CreationTime int64 `json:"creation_time,omitempty"`
	// ID of the workspace's credential configuration object.
	// Wire name: 'credentials_id'
	CredentialsId string `json:"credentials_id,omitempty"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	// Wire name: 'custom_tags'
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for web application and REST APIs is
	// `<deployment-name>.cloud.databricks.com`.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	// Wire name: 'deployment_name'
	DeploymentName string `json:"deployment_name,omitempty"`
	// If this workspace is for a external customer, then external_customer_info
	// is populated. If this workspace is not for a external customer, then
	// external_customer_info is empty.
	// Wire name: 'external_customer_info'
	ExternalCustomerInfo *ExternalCustomerInfo `json:"external_customer_info,omitempty"`

	// Wire name: 'gcp_managed_network_config'
	GcpManagedNetworkConfig *GcpManagedNetworkConfig `json:"gcp_managed_network_config,omitempty"`

	// Wire name: 'gke_config'
	GkeConfig *GkeConfig `json:"gke_config,omitempty"`
	// Whether no public IP is enabled for the workspace.
	// Wire name: 'is_no_public_ip_enabled'
	IsNoPublicIpEnabled bool `json:"is_no_public_ip_enabled,omitempty"`
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	// Wire name: 'location'
	Location string `json:"location,omitempty"`
	// ID of the key configuration for encrypting managed services.
	// Wire name: 'managed_services_customer_managed_key_id'
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`
	// The network configuration ID that is attached to the workspace. This
	// field is available only if the network is a customer-managed network.
	// Wire name: 'network_id'
	NetworkId string `json:"network_id,omitempty"`

	// Wire name: 'pricing_tier'
	PricingTier PricingTier `json:"pricing_tier,omitempty"`
	// ID of the workspace's private access settings object. Only used for
	// PrivateLink. You must specify this ID if you are using [AWS PrivateLink]
	// for either front-end (user-to-workspace connection), back-end (data plane
	// to control plane connection), or both connection types.
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink].",
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// ID of the workspace's storage configuration object.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// ID of the key configuration for encrypting workspace storage.
	// Wire name: 'storage_customer_managed_key_id'
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// A unique integer ID for the workspace
	// Wire name: 'workspace_id'
	WorkspaceId int64 `json:"workspace_id,omitempty"`
	// The human-readable name of the workspace.
	// Wire name: 'workspace_name'
	WorkspaceName string `json:"workspace_name,omitempty"`

	// Wire name: 'workspace_status'
	WorkspaceStatus WorkspaceStatus `json:"workspace_status,omitempty"`
	// Message describing the current workspace status.
	// Wire name: 'workspace_status_message'
	WorkspaceStatusMessage string `json:"workspace_status_message,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Workspace) EncodeValues(key string, v *url.Values) error {
	pb, err := workspaceToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Workspace) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := workspaceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Workspace) MarshalJSON() ([]byte, error) {
	pb, err := workspaceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The status of the workspace. For workspace creation, usually it is set to
// `PROVISIONING` initially. Continue to check the status until the status is
// `RUNNING`.
type WorkspaceStatus string

const WorkspaceStatusBanned WorkspaceStatus = `BANNED`

const WorkspaceStatusCancelling WorkspaceStatus = `CANCELLING`

const WorkspaceStatusFailed WorkspaceStatus = `FAILED`

const WorkspaceStatusNotProvisioned WorkspaceStatus = `NOT_PROVISIONED`

const WorkspaceStatusProvisioning WorkspaceStatus = `PROVISIONING`

const WorkspaceStatusRunning WorkspaceStatus = `RUNNING`

// String representation for [fmt.Print]
func (f *WorkspaceStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspaceStatus) Set(v string) error {
	switch v {
	case `BANNED`, `CANCELLING`, `FAILED`, `NOT_PROVISIONED`, `PROVISIONING`, `RUNNING`:
		*f = WorkspaceStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BANNED", "CANCELLING", "FAILED", "NOT_PROVISIONED", "PROVISIONING", "RUNNING"`, v)
	}
}

// Values returns all possible values for WorkspaceStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspaceStatus) Values() []WorkspaceStatus {
	return []WorkspaceStatus{
		WorkspaceStatusBanned,
		WorkspaceStatusCancelling,
		WorkspaceStatusFailed,
		WorkspaceStatusNotProvisioned,
		WorkspaceStatusProvisioning,
		WorkspaceStatusRunning,
	}
}

// Type always returns WorkspaceStatus to satisfy [pflag.Value] interface
func (f *WorkspaceStatus) Type() string {
	return "WorkspaceStatus"
}
