// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioningpb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type AwsCredentialsPb struct {
	StsRole *StsRolePb `json:"sts_role,omitempty"`
}

type AwsKeyInfoPb struct {
	KeyAlias                  string `json:"key_alias,omitempty"`
	KeyArn                    string `json:"key_arn"`
	KeyRegion                 string `json:"key_region"`
	ReuseKeyForClusterVolumes bool   `json:"reuse_key_for_cluster_volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AwsKeyInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AwsKeyInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AzureWorkspaceInfoPb struct {
	ResourceGroup  string `json:"resource_group,omitempty"`
	SubscriptionId string `json:"subscription_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AzureWorkspaceInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AzureWorkspaceInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CloudResourceContainerPb struct {
	Gcp *CustomerFacingGcpCloudResourceContainerPb `json:"gcp,omitempty"`
}

type CreateAwsKeyInfoPb struct {
	KeyAlias                  string `json:"key_alias,omitempty"`
	KeyArn                    string `json:"key_arn"`
	ReuseKeyForClusterVolumes bool   `json:"reuse_key_for_cluster_volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateAwsKeyInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateAwsKeyInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCredentialAwsCredentialsPb struct {
	StsRole *CreateCredentialStsRolePb `json:"sts_role,omitempty"`
}

type CreateCredentialRequestPb struct {
	AwsCredentials  CreateCredentialAwsCredentialsPb `json:"aws_credentials"`
	CredentialsName string                           `json:"credentials_name"`
}

type CreateCredentialStsRolePb struct {
	RoleArn string `json:"role_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateCredentialStsRolePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateCredentialStsRolePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCustomerManagedKeyRequestPb struct {
	AwsKeyInfo *CreateAwsKeyInfoPb `json:"aws_key_info,omitempty"`
	GcpKeyInfo *CreateGcpKeyInfoPb `json:"gcp_key_info,omitempty"`
	UseCases   []KeyUseCasePb      `json:"use_cases"`
}

type CreateGcpKeyInfoPb struct {
	KmsKeyId string `json:"kms_key_id"`
}

type CreateNetworkRequestPb struct {
	GcpNetworkInfo   *GcpNetworkInfoPb      `json:"gcp_network_info,omitempty"`
	NetworkName      string                 `json:"network_name"`
	SecurityGroupIds []string               `json:"security_group_ids,omitempty"`
	SubnetIds        []string               `json:"subnet_ids,omitempty"`
	VpcEndpoints     *NetworkVpcEndpointsPb `json:"vpc_endpoints,omitempty"`
	VpcId            string                 `json:"vpc_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateNetworkRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateNetworkRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateStorageConfigurationRequestPb struct {
	RootBucketInfo           RootBucketInfoPb `json:"root_bucket_info"`
	StorageConfigurationName string           `json:"storage_configuration_name"`
}

type CreateVpcEndpointRequestPb struct {
	AwsVpcEndpointId   string                `json:"aws_vpc_endpoint_id,omitempty"`
	GcpVpcEndpointInfo *GcpVpcEndpointInfoPb `json:"gcp_vpc_endpoint_info,omitempty"`
	Region             string                `json:"region,omitempty"`
	VpcEndpointName    string                `json:"vpc_endpoint_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateVpcEndpointRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateVpcEndpointRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateWorkspaceRequestPb struct {
	AwsRegion                           string                     `json:"aws_region,omitempty"`
	Cloud                               string                     `json:"cloud,omitempty"`
	CloudResourceContainer              *CloudResourceContainerPb  `json:"cloud_resource_container,omitempty"`
	CredentialsId                       string                     `json:"credentials_id,omitempty"`
	CustomTags                          map[string]string          `json:"custom_tags,omitempty"`
	DeploymentName                      string                     `json:"deployment_name,omitempty"`
	GcpManagedNetworkConfig             *GcpManagedNetworkConfigPb `json:"gcp_managed_network_config,omitempty"`
	GkeConfig                           *GkeConfigPb               `json:"gke_config,omitempty"`
	IsNoPublicIpEnabled                 bool                       `json:"is_no_public_ip_enabled,omitempty"`
	Location                            string                     `json:"location,omitempty"`
	ManagedServicesCustomerManagedKeyId string                     `json:"managed_services_customer_managed_key_id,omitempty"`
	NetworkId                           string                     `json:"network_id,omitempty"`
	PricingTier                         PricingTierPb              `json:"pricing_tier,omitempty"`
	PrivateAccessSettingsId             string                     `json:"private_access_settings_id,omitempty"`
	StorageConfigurationId              string                     `json:"storage_configuration_id,omitempty"`
	StorageCustomerManagedKeyId         string                     `json:"storage_customer_managed_key_id,omitempty"`
	WorkspaceName                       string                     `json:"workspace_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateWorkspaceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateWorkspaceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CredentialPb struct {
	AccountId       string            `json:"account_id,omitempty"`
	AwsCredentials  *AwsCredentialsPb `json:"aws_credentials,omitempty"`
	CreationTime    int64             `json:"creation_time,omitempty"`
	CredentialsId   string            `json:"credentials_id,omitempty"`
	CredentialsName string            `json:"credentials_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CredentialPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CredentialPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CustomerFacingGcpCloudResourceContainerPb struct {
	ProjectId string `json:"project_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CustomerFacingGcpCloudResourceContainerPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CustomerFacingGcpCloudResourceContainerPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CustomerManagedKeyPb struct {
	AccountId            string         `json:"account_id,omitempty"`
	AwsKeyInfo           *AwsKeyInfoPb  `json:"aws_key_info,omitempty"`
	CreationTime         int64          `json:"creation_time,omitempty"`
	CustomerManagedKeyId string         `json:"customer_managed_key_id,omitempty"`
	GcpKeyInfo           *GcpKeyInfoPb  `json:"gcp_key_info,omitempty"`
	UseCases             []KeyUseCasePb `json:"use_cases,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CustomerManagedKeyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CustomerManagedKeyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteCredentialRequestPb struct {
	CredentialsId string `json:"-" url:"-"`
}

type DeleteEncryptionKeyRequestPb struct {
	CustomerManagedKeyId string `json:"-" url:"-"`
}

type DeleteNetworkRequestPb struct {
	NetworkId string `json:"-" url:"-"`
}

type DeletePrivateAccesRequestPb struct {
	PrivateAccessSettingsId string `json:"-" url:"-"`
}

type DeleteStorageRequestPb struct {
	StorageConfigurationId string `json:"-" url:"-"`
}

type DeleteVpcEndpointRequestPb struct {
	VpcEndpointId string `json:"-" url:"-"`
}

type DeleteWorkspaceRequestPb struct {
	WorkspaceId int64 `json:"-" url:"-"`
}

type EndpointUseCasePb string

const EndpointUseCaseDataplaneRelayAccess EndpointUseCasePb = `DATAPLANE_RELAY_ACCESS`

const EndpointUseCaseWorkspaceAccess EndpointUseCasePb = `WORKSPACE_ACCESS`

type ErrorTypePb string

const ErrorTypeCredentials ErrorTypePb = `credentials`

const ErrorTypeNetworkAcl ErrorTypePb = `networkAcl`

const ErrorTypeSecurityGroup ErrorTypePb = `securityGroup`

const ErrorTypeSubnet ErrorTypePb = `subnet`

const ErrorTypeVpc ErrorTypePb = `vpc`

type ExternalCustomerInfoPb struct {
	AuthoritativeUserEmail    string `json:"authoritative_user_email,omitempty"`
	AuthoritativeUserFullName string `json:"authoritative_user_full_name,omitempty"`
	CustomerName              string `json:"customer_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalCustomerInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalCustomerInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GcpKeyInfoPb struct {
	KmsKeyId string `json:"kms_key_id"`
}

type GcpManagedNetworkConfigPb struct {
	GkeClusterPodIpRange     string `json:"gke_cluster_pod_ip_range,omitempty"`
	GkeClusterServiceIpRange string `json:"gke_cluster_service_ip_range,omitempty"`
	SubnetCidr               string `json:"subnet_cidr,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GcpManagedNetworkConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GcpManagedNetworkConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GcpNetworkInfoPb struct {
	NetworkProjectId   string `json:"network_project_id"`
	PodIpRangeName     string `json:"pod_ip_range_name"`
	ServiceIpRangeName string `json:"service_ip_range_name"`
	SubnetId           string `json:"subnet_id"`
	SubnetRegion       string `json:"subnet_region"`
	VpcId              string `json:"vpc_id"`
}

type GcpVpcEndpointInfoPb struct {
	EndpointRegion      string `json:"endpoint_region"`
	ProjectId           string `json:"project_id"`
	PscConnectionId     string `json:"psc_connection_id,omitempty"`
	PscEndpointName     string `json:"psc_endpoint_name"`
	ServiceAttachmentId string `json:"service_attachment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GcpVpcEndpointInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GcpVpcEndpointInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetCredentialRequestPb struct {
	CredentialsId string `json:"-" url:"-"`
}

type GetEncryptionKeyRequestPb struct {
	CustomerManagedKeyId string `json:"-" url:"-"`
}

type GetNetworkRequestPb struct {
	NetworkId string `json:"-" url:"-"`
}

type GetPrivateAccesRequestPb struct {
	PrivateAccessSettingsId string `json:"-" url:"-"`
}

type GetStorageRequestPb struct {
	StorageConfigurationId string `json:"-" url:"-"`
}

type GetVpcEndpointRequestPb struct {
	VpcEndpointId string `json:"-" url:"-"`
}

type GetWorkspaceRequestPb struct {
	WorkspaceId int64 `json:"-" url:"-"`
}

type GkeConfigPb struct {
	ConnectivityType GkeConfigConnectivityTypePb `json:"connectivity_type,omitempty"`
	MasterIpRange    string                      `json:"master_ip_range,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GkeConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GkeConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GkeConfigConnectivityTypePb string

const GkeConfigConnectivityTypePrivateNodePublicMaster GkeConfigConnectivityTypePb = `PRIVATE_NODE_PUBLIC_MASTER`

const GkeConfigConnectivityTypePublicNodePublicMaster GkeConfigConnectivityTypePb = `PUBLIC_NODE_PUBLIC_MASTER`

type KeyUseCasePb string

// Encrypts notebook and secret data in the control plane
const KeyUseCaseManagedServices KeyUseCasePb = `MANAGED_SERVICES`

// Encrypts the workspace's root S3 bucket (root DBFS and system data) and,
// optionally, cluster EBS volumes.
const KeyUseCaseStorage KeyUseCasePb = `STORAGE`

type ListCredentialsRequestPb struct {
}

type ListEncryptionKeysRequestPb struct {
}

type ListNetworksRequestPb struct {
}

type ListPrivateAccessRequestPb struct {
}

type ListStorageRequestPb struct {
}

type ListVpcEndpointsRequestPb struct {
}

type ListWorkspacesRequestPb struct {
}

type NetworkPb struct {
	AccountId        string                 `json:"account_id,omitempty"`
	CreationTime     int64                  `json:"creation_time,omitempty"`
	ErrorMessages    []NetworkHealthPb      `json:"error_messages,omitempty"`
	GcpNetworkInfo   *GcpNetworkInfoPb      `json:"gcp_network_info,omitempty"`
	NetworkId        string                 `json:"network_id,omitempty"`
	NetworkName      string                 `json:"network_name,omitempty"`
	SecurityGroupIds []string               `json:"security_group_ids,omitempty"`
	SubnetIds        []string               `json:"subnet_ids,omitempty"`
	VpcEndpoints     *NetworkVpcEndpointsPb `json:"vpc_endpoints,omitempty"`
	VpcId            string                 `json:"vpc_id,omitempty"`
	VpcStatus        VpcStatusPb            `json:"vpc_status,omitempty"`
	WarningMessages  []NetworkWarningPb     `json:"warning_messages,omitempty"`
	WorkspaceId      int64                  `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NetworkPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NetworkPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NetworkHealthPb struct {
	ErrorMessage string      `json:"error_message,omitempty"`
	ErrorType    ErrorTypePb `json:"error_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NetworkHealthPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NetworkHealthPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NetworkVpcEndpointsPb struct {
	DataplaneRelay []string `json:"dataplane_relay"`
	RestApi        []string `json:"rest_api"`
}

type NetworkWarningPb struct {
	WarningMessage string        `json:"warning_message,omitempty"`
	WarningType    WarningTypePb `json:"warning_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NetworkWarningPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NetworkWarningPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PricingTierPb string

const PricingTierCommunityEdition PricingTierPb = `COMMUNITY_EDITION`

const PricingTierDedicated PricingTierPb = `DEDICATED`

const PricingTierEnterprise PricingTierPb = `ENTERPRISE`

const PricingTierPremium PricingTierPb = `PREMIUM`

const PricingTierStandard PricingTierPb = `STANDARD`

const PricingTierUnknown PricingTierPb = `UNKNOWN`

type PrivateAccessLevelPb string

const PrivateAccessLevelAccount PrivateAccessLevelPb = `ACCOUNT`

const PrivateAccessLevelEndpoint PrivateAccessLevelPb = `ENDPOINT`

type PrivateAccessSettingsPb struct {
	AccountId                 string               `json:"account_id,omitempty"`
	AllowedVpcEndpointIds     []string             `json:"allowed_vpc_endpoint_ids,omitempty"`
	PrivateAccessLevel        PrivateAccessLevelPb `json:"private_access_level,omitempty"`
	PrivateAccessSettingsId   string               `json:"private_access_settings_id,omitempty"`
	PrivateAccessSettingsName string               `json:"private_access_settings_name,omitempty"`
	PublicAccessEnabled       bool                 `json:"public_access_enabled,omitempty"`
	Region                    string               `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PrivateAccessSettingsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PrivateAccessSettingsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RootBucketInfoPb struct {
	BucketName string `json:"bucket_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RootBucketInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RootBucketInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StorageConfigurationPb struct {
	AccountId                string            `json:"account_id,omitempty"`
	CreationTime             int64             `json:"creation_time,omitempty"`
	RootBucketInfo           *RootBucketInfoPb `json:"root_bucket_info,omitempty"`
	StorageConfigurationId   string            `json:"storage_configuration_id,omitempty"`
	StorageConfigurationName string            `json:"storage_configuration_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *StorageConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st StorageConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StsRolePb struct {
	ExternalId string `json:"external_id,omitempty"`
	RoleArn    string `json:"role_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *StsRolePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st StsRolePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateWorkspaceRequestPb struct {
	AwsRegion                           string            `json:"aws_region,omitempty"`
	CredentialsId                       string            `json:"credentials_id,omitempty"`
	CustomTags                          map[string]string `json:"custom_tags,omitempty"`
	ManagedServicesCustomerManagedKeyId string            `json:"managed_services_customer_managed_key_id,omitempty"`
	NetworkConnectivityConfigId         string            `json:"network_connectivity_config_id,omitempty"`
	NetworkId                           string            `json:"network_id,omitempty"`
	PrivateAccessSettingsId             string            `json:"private_access_settings_id,omitempty"`
	StorageConfigurationId              string            `json:"storage_configuration_id,omitempty"`
	StorageCustomerManagedKeyId         string            `json:"storage_customer_managed_key_id,omitempty"`
	WorkspaceId                         int64             `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateWorkspaceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateWorkspaceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpsertPrivateAccessSettingsRequestPb struct {
	AllowedVpcEndpointIds     []string             `json:"allowed_vpc_endpoint_ids,omitempty"`
	PrivateAccessLevel        PrivateAccessLevelPb `json:"private_access_level,omitempty"`
	PrivateAccessSettingsId   string               `json:"-" url:"-"`
	PrivateAccessSettingsName string               `json:"private_access_settings_name"`
	PublicAccessEnabled       bool                 `json:"public_access_enabled,omitempty"`
	Region                    string               `json:"region"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpsertPrivateAccessSettingsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpsertPrivateAccessSettingsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type VpcEndpointPb struct {
	AccountId            string                `json:"account_id,omitempty"`
	AwsAccountId         string                `json:"aws_account_id,omitempty"`
	AwsEndpointServiceId string                `json:"aws_endpoint_service_id,omitempty"`
	AwsVpcEndpointId     string                `json:"aws_vpc_endpoint_id,omitempty"`
	GcpVpcEndpointInfo   *GcpVpcEndpointInfoPb `json:"gcp_vpc_endpoint_info,omitempty"`
	Region               string                `json:"region,omitempty"`
	State                string                `json:"state,omitempty"`
	UseCase              EndpointUseCasePb     `json:"use_case,omitempty"`
	VpcEndpointId        string                `json:"vpc_endpoint_id,omitempty"`
	VpcEndpointName      string                `json:"vpc_endpoint_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *VpcEndpointPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st VpcEndpointPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type VpcStatusPb string

// Broken.
const VpcStatusBroken VpcStatusPb = `BROKEN`

// Unattached.
const VpcStatusUnattached VpcStatusPb = `UNATTACHED`

// Valid.
const VpcStatusValid VpcStatusPb = `VALID`

// Warned.
const VpcStatusWarned VpcStatusPb = `WARNED`

type WarningTypePb string

const WarningTypeSecurityGroup WarningTypePb = `securityGroup`

const WarningTypeSubnet WarningTypePb = `subnet`

type WorkspacePb struct {
	AccountId                           string                     `json:"account_id,omitempty"`
	AwsRegion                           string                     `json:"aws_region,omitempty"`
	AzureWorkspaceInfo                  *AzureWorkspaceInfoPb      `json:"azure_workspace_info,omitempty"`
	Cloud                               string                     `json:"cloud,omitempty"`
	CloudResourceContainer              *CloudResourceContainerPb  `json:"cloud_resource_container,omitempty"`
	CreationTime                        int64                      `json:"creation_time,omitempty"`
	CredentialsId                       string                     `json:"credentials_id,omitempty"`
	CustomTags                          map[string]string          `json:"custom_tags,omitempty"`
	DeploymentName                      string                     `json:"deployment_name,omitempty"`
	ExternalCustomerInfo                *ExternalCustomerInfoPb    `json:"external_customer_info,omitempty"`
	GcpManagedNetworkConfig             *GcpManagedNetworkConfigPb `json:"gcp_managed_network_config,omitempty"`
	GkeConfig                           *GkeConfigPb               `json:"gke_config,omitempty"`
	IsNoPublicIpEnabled                 bool                       `json:"is_no_public_ip_enabled,omitempty"`
	Location                            string                     `json:"location,omitempty"`
	ManagedServicesCustomerManagedKeyId string                     `json:"managed_services_customer_managed_key_id,omitempty"`
	NetworkId                           string                     `json:"network_id,omitempty"`
	PricingTier                         PricingTierPb              `json:"pricing_tier,omitempty"`
	PrivateAccessSettingsId             string                     `json:"private_access_settings_id,omitempty"`
	StorageConfigurationId              string                     `json:"storage_configuration_id,omitempty"`
	StorageCustomerManagedKeyId         string                     `json:"storage_customer_managed_key_id,omitempty"`
	WorkspaceId                         int64                      `json:"workspace_id,omitempty"`
	WorkspaceName                       string                     `json:"workspace_name,omitempty"`
	WorkspaceStatus                     WorkspaceStatusPb          `json:"workspace_status,omitempty"`
	WorkspaceStatusMessage              string                     `json:"workspace_status_message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WorkspacePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WorkspacePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WorkspaceStatusPb string

const WorkspaceStatusBanned WorkspaceStatusPb = `BANNED`

const WorkspaceStatusCancelling WorkspaceStatusPb = `CANCELLING`

const WorkspaceStatusFailed WorkspaceStatusPb = `FAILED`

const WorkspaceStatusNotProvisioned WorkspaceStatusPb = `NOT_PROVISIONED`

const WorkspaceStatusProvisioning WorkspaceStatusPb = `PROVISIONING`

const WorkspaceStatusRunning WorkspaceStatusPb = `RUNNING`
