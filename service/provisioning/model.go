// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AwsCredentials struct {

	// Wire name: 'sts_role'
	StsRole *StsRole
}

func awsCredentialsToPb(st *AwsCredentials) (*awsCredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsCredentialsPb{}
	stsRolePb, err := stsRoleToPb(st.StsRole)
	if err != nil {
		return nil, err
	}
	if stsRolePb != nil {
		pb.StsRole = stsRolePb
	}

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
	StsRole *stsRolePb `json:"sts_role,omitempty"`
}

func awsCredentialsFromPb(pb *awsCredentialsPb) (*AwsCredentials, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsCredentials{}
	stsRoleField, err := stsRoleFromPb(pb.StsRole)
	if err != nil {
		return nil, err
	}
	if stsRoleField != nil {
		st.StsRole = stsRoleField
	}

	return st, nil
}

type AwsKeyInfo struct {
	// The AWS KMS key alias.
	// Wire name: 'key_alias'
	KeyAlias string
	// The AWS KMS key's Amazon Resource Name (ARN).
	// Wire name: 'key_arn'
	KeyArn string
	// The AWS KMS key region.
	// Wire name: 'key_region'
	KeyRegion string
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. If you do not want to use this key for encrypting
	// EBS volumes, set to `false`.
	// Wire name: 'reuse_key_for_cluster_volumes'
	ReuseKeyForClusterVolumes bool

	ForceSendFields []string `tf:"-"`
}

func awsKeyInfoToPb(st *AwsKeyInfo) (*awsKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsKeyInfoPb{}
	pb.KeyAlias = st.KeyAlias

	pb.KeyArn = st.KeyArn

	pb.KeyRegion = st.KeyRegion

	pb.ReuseKeyForClusterVolumes = st.ReuseKeyForClusterVolumes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type awsKeyInfoPb struct {
	// The AWS KMS key alias.
	KeyAlias string `json:"key_alias,omitempty"`
	// The AWS KMS key's Amazon Resource Name (ARN).
	KeyArn string `json:"key_arn"`
	// The AWS KMS key region.
	KeyRegion string `json:"key_region"`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. If you do not want to use this key for encrypting
	// EBS volumes, set to `false`.
	ReuseKeyForClusterVolumes bool `json:"reuse_key_for_cluster_volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func awsKeyInfoFromPb(pb *awsKeyInfoPb) (*AwsKeyInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsKeyInfo{}
	st.KeyAlias = pb.KeyAlias
	st.KeyArn = pb.KeyArn
	st.KeyRegion = pb.KeyRegion
	st.ReuseKeyForClusterVolumes = pb.ReuseKeyForClusterVolumes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *awsKeyInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st awsKeyInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AzureWorkspaceInfo struct {
	// Azure Resource Group name
	// Wire name: 'resource_group'
	ResourceGroup string
	// Azure Subscription ID
	// Wire name: 'subscription_id'
	SubscriptionId string

	ForceSendFields []string `tf:"-"`
}

func azureWorkspaceInfoToPb(st *AzureWorkspaceInfo) (*azureWorkspaceInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureWorkspaceInfoPb{}
	pb.ResourceGroup = st.ResourceGroup

	pb.SubscriptionId = st.SubscriptionId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type azureWorkspaceInfoPb struct {
	// Azure Resource Group name
	ResourceGroup string `json:"resource_group,omitempty"`
	// Azure Subscription ID
	SubscriptionId string `json:"subscription_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func azureWorkspaceInfoFromPb(pb *azureWorkspaceInfoPb) (*AzureWorkspaceInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureWorkspaceInfo{}
	st.ResourceGroup = pb.ResourceGroup
	st.SubscriptionId = pb.SubscriptionId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *azureWorkspaceInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st azureWorkspaceInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The general workspace configurations that are specific to cloud providers.
type CloudResourceContainer struct {
	// The general workspace configurations that are specific to Google Cloud.
	// Wire name: 'gcp'
	Gcp *CustomerFacingGcpCloudResourceContainer
}

func cloudResourceContainerToPb(st *CloudResourceContainer) (*cloudResourceContainerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cloudResourceContainerPb{}
	gcpPb, err := customerFacingGcpCloudResourceContainerToPb(st.Gcp)
	if err != nil {
		return nil, err
	}
	if gcpPb != nil {
		pb.Gcp = gcpPb
	}

	return pb, nil
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

type cloudResourceContainerPb struct {
	// The general workspace configurations that are specific to Google Cloud.
	Gcp *customerFacingGcpCloudResourceContainerPb `json:"gcp,omitempty"`
}

func cloudResourceContainerFromPb(pb *cloudResourceContainerPb) (*CloudResourceContainer, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CloudResourceContainer{}
	gcpField, err := customerFacingGcpCloudResourceContainerFromPb(pb.Gcp)
	if err != nil {
		return nil, err
	}
	if gcpField != nil {
		st.Gcp = gcpField
	}

	return st, nil
}

type CreateAwsKeyInfo struct {
	// The AWS KMS key alias.
	// Wire name: 'key_alias'
	KeyAlias string
	// The AWS KMS key's Amazon Resource Name (ARN). Note that the key's AWS
	// region is inferred from the ARN.
	// Wire name: 'key_arn'
	KeyArn string
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. To not use this key also for encrypting EBS volumes,
	// set this to `false`.
	// Wire name: 'reuse_key_for_cluster_volumes'
	ReuseKeyForClusterVolumes bool

	ForceSendFields []string `tf:"-"`
}

func createAwsKeyInfoToPb(st *CreateAwsKeyInfo) (*createAwsKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAwsKeyInfoPb{}
	pb.KeyAlias = st.KeyAlias

	pb.KeyArn = st.KeyArn

	pb.ReuseKeyForClusterVolumes = st.ReuseKeyForClusterVolumes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createAwsKeyInfoPb struct {
	// The AWS KMS key alias.
	KeyAlias string `json:"key_alias,omitempty"`
	// The AWS KMS key's Amazon Resource Name (ARN). Note that the key's AWS
	// region is inferred from the ARN.
	KeyArn string `json:"key_arn"`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. To not use this key also for encrypting EBS volumes,
	// set this to `false`.
	ReuseKeyForClusterVolumes bool `json:"reuse_key_for_cluster_volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createAwsKeyInfoFromPb(pb *createAwsKeyInfoPb) (*CreateAwsKeyInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAwsKeyInfo{}
	st.KeyAlias = pb.KeyAlias
	st.KeyArn = pb.KeyArn
	st.ReuseKeyForClusterVolumes = pb.ReuseKeyForClusterVolumes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createAwsKeyInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createAwsKeyInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCredentialAwsCredentials struct {

	// Wire name: 'sts_role'
	StsRole *CreateCredentialStsRole
}

func createCredentialAwsCredentialsToPb(st *CreateCredentialAwsCredentials) (*createCredentialAwsCredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCredentialAwsCredentialsPb{}
	stsRolePb, err := createCredentialStsRoleToPb(st.StsRole)
	if err != nil {
		return nil, err
	}
	if stsRolePb != nil {
		pb.StsRole = stsRolePb
	}

	return pb, nil
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

type createCredentialAwsCredentialsPb struct {
	StsRole *createCredentialStsRolePb `json:"sts_role,omitempty"`
}

func createCredentialAwsCredentialsFromPb(pb *createCredentialAwsCredentialsPb) (*CreateCredentialAwsCredentials, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialAwsCredentials{}
	stsRoleField, err := createCredentialStsRoleFromPb(pb.StsRole)
	if err != nil {
		return nil, err
	}
	if stsRoleField != nil {
		st.StsRole = stsRoleField
	}

	return st, nil
}

type CreateCredentialRequest struct {

	// Wire name: 'aws_credentials'
	AwsCredentials CreateCredentialAwsCredentials
	// The human-readable name of the credential configuration object.
	// Wire name: 'credentials_name'
	CredentialsName string
}

func createCredentialRequestToPb(st *CreateCredentialRequest) (*createCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCredentialRequestPb{}
	awsCredentialsPb, err := createCredentialAwsCredentialsToPb(&st.AwsCredentials)
	if err != nil {
		return nil, err
	}
	if awsCredentialsPb != nil {
		pb.AwsCredentials = *awsCredentialsPb
	}

	pb.CredentialsName = st.CredentialsName

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
	AwsCredentials createCredentialAwsCredentialsPb `json:"aws_credentials"`
	// The human-readable name of the credential configuration object.
	CredentialsName string `json:"credentials_name"`
}

func createCredentialRequestFromPb(pb *createCredentialRequestPb) (*CreateCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialRequest{}
	awsCredentialsField, err := createCredentialAwsCredentialsFromPb(&pb.AwsCredentials)
	if err != nil {
		return nil, err
	}
	if awsCredentialsField != nil {
		st.AwsCredentials = *awsCredentialsField
	}
	st.CredentialsName = pb.CredentialsName

	return st, nil
}

type CreateCredentialStsRole struct {
	// The Amazon Resource Name (ARN) of the cross account role.
	// Wire name: 'role_arn'
	RoleArn string

	ForceSendFields []string `tf:"-"`
}

func createCredentialStsRoleToPb(st *CreateCredentialStsRole) (*createCredentialStsRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCredentialStsRolePb{}
	pb.RoleArn = st.RoleArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createCredentialStsRolePb struct {
	// The Amazon Resource Name (ARN) of the cross account role.
	RoleArn string `json:"role_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createCredentialStsRoleFromPb(pb *createCredentialStsRolePb) (*CreateCredentialStsRole, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialStsRole{}
	st.RoleArn = pb.RoleArn

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createCredentialStsRolePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createCredentialStsRolePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCustomerManagedKeyRequest struct {

	// Wire name: 'aws_key_info'
	AwsKeyInfo *CreateAwsKeyInfo

	// Wire name: 'gcp_key_info'
	GcpKeyInfo *CreateGcpKeyInfo
	// The cases that the key can be used for.
	// Wire name: 'use_cases'
	UseCases []KeyUseCase
}

func createCustomerManagedKeyRequestToPb(st *CreateCustomerManagedKeyRequest) (*createCustomerManagedKeyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCustomerManagedKeyRequestPb{}
	awsKeyInfoPb, err := createAwsKeyInfoToPb(st.AwsKeyInfo)
	if err != nil {
		return nil, err
	}
	if awsKeyInfoPb != nil {
		pb.AwsKeyInfo = awsKeyInfoPb
	}

	gcpKeyInfoPb, err := createGcpKeyInfoToPb(st.GcpKeyInfo)
	if err != nil {
		return nil, err
	}
	if gcpKeyInfoPb != nil {
		pb.GcpKeyInfo = gcpKeyInfoPb
	}

	pb.UseCases = st.UseCases

	return pb, nil
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

type createCustomerManagedKeyRequestPb struct {
	AwsKeyInfo *createAwsKeyInfoPb `json:"aws_key_info,omitempty"`

	GcpKeyInfo *createGcpKeyInfoPb `json:"gcp_key_info,omitempty"`
	// The cases that the key can be used for.
	UseCases []KeyUseCase `json:"use_cases"`
}

func createCustomerManagedKeyRequestFromPb(pb *createCustomerManagedKeyRequestPb) (*CreateCustomerManagedKeyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCustomerManagedKeyRequest{}
	awsKeyInfoField, err := createAwsKeyInfoFromPb(pb.AwsKeyInfo)
	if err != nil {
		return nil, err
	}
	if awsKeyInfoField != nil {
		st.AwsKeyInfo = awsKeyInfoField
	}
	gcpKeyInfoField, err := createGcpKeyInfoFromPb(pb.GcpKeyInfo)
	if err != nil {
		return nil, err
	}
	if gcpKeyInfoField != nil {
		st.GcpKeyInfo = gcpKeyInfoField
	}
	st.UseCases = pb.UseCases

	return st, nil
}

type CreateGcpKeyInfo struct {
	// The GCP KMS key's resource name
	// Wire name: 'kms_key_id'
	KmsKeyId string
}

func createGcpKeyInfoToPb(st *CreateGcpKeyInfo) (*createGcpKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createGcpKeyInfoPb{}
	pb.KmsKeyId = st.KmsKeyId

	return pb, nil
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

type createGcpKeyInfoPb struct {
	// The GCP KMS key's resource name
	KmsKeyId string `json:"kms_key_id"`
}

func createGcpKeyInfoFromPb(pb *createGcpKeyInfoPb) (*CreateGcpKeyInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateGcpKeyInfo{}
	st.KmsKeyId = pb.KmsKeyId

	return st, nil
}

type CreateNetworkRequest struct {
	// The Google Cloud specific information for this network (for example, the
	// VPC ID, subnet ID, and secondary IP ranges).
	// Wire name: 'gcp_network_info'
	GcpNetworkInfo *GcpNetworkInfo
	// The human-readable name of the network configuration.
	// Wire name: 'network_name'
	NetworkName string
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	// Wire name: 'security_group_ids'
	SecurityGroupIds []string
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	// Wire name: 'subnet_ids'
	SubnetIds []string
	// If specified, contains the VPC endpoints used to allow cluster
	// communication from this VPC over [AWS PrivateLink].
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// Wire name: 'vpc_endpoints'
	VpcEndpoints *NetworkVpcEndpoints
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	// Wire name: 'vpc_id'
	VpcId string

	ForceSendFields []string `tf:"-"`
}

func createNetworkRequestToPb(st *CreateNetworkRequest) (*createNetworkRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createNetworkRequestPb{}
	gcpNetworkInfoPb, err := gcpNetworkInfoToPb(st.GcpNetworkInfo)
	if err != nil {
		return nil, err
	}
	if gcpNetworkInfoPb != nil {
		pb.GcpNetworkInfo = gcpNetworkInfoPb
	}

	pb.NetworkName = st.NetworkName

	pb.SecurityGroupIds = st.SecurityGroupIds

	pb.SubnetIds = st.SubnetIds

	vpcEndpointsPb, err := networkVpcEndpointsToPb(st.VpcEndpoints)
	if err != nil {
		return nil, err
	}
	if vpcEndpointsPb != nil {
		pb.VpcEndpoints = vpcEndpointsPb
	}

	pb.VpcId = st.VpcId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createNetworkRequestPb struct {
	// The Google Cloud specific information for this network (for example, the
	// VPC ID, subnet ID, and secondary IP ranges).
	GcpNetworkInfo *gcpNetworkInfoPb `json:"gcp_network_info,omitempty"`
	// The human-readable name of the network configuration.
	NetworkName string `json:"network_name"`
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds []string `json:"security_group_ids,omitempty"`
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	SubnetIds []string `json:"subnet_ids,omitempty"`
	// If specified, contains the VPC endpoints used to allow cluster
	// communication from this VPC over [AWS PrivateLink].
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	VpcEndpoints *networkVpcEndpointsPb `json:"vpc_endpoints,omitempty"`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId string `json:"vpc_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createNetworkRequestFromPb(pb *createNetworkRequestPb) (*CreateNetworkRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateNetworkRequest{}
	gcpNetworkInfoField, err := gcpNetworkInfoFromPb(pb.GcpNetworkInfo)
	if err != nil {
		return nil, err
	}
	if gcpNetworkInfoField != nil {
		st.GcpNetworkInfo = gcpNetworkInfoField
	}
	st.NetworkName = pb.NetworkName
	st.SecurityGroupIds = pb.SecurityGroupIds
	st.SubnetIds = pb.SubnetIds
	vpcEndpointsField, err := networkVpcEndpointsFromPb(pb.VpcEndpoints)
	if err != nil {
		return nil, err
	}
	if vpcEndpointsField != nil {
		st.VpcEndpoints = vpcEndpointsField
	}
	st.VpcId = pb.VpcId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createNetworkRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createNetworkRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateStorageConfigurationRequest struct {
	// Root S3 bucket information.
	// Wire name: 'root_bucket_info'
	RootBucketInfo RootBucketInfo
	// The human-readable name of the storage configuration.
	// Wire name: 'storage_configuration_name'
	StorageConfigurationName string
}

func createStorageConfigurationRequestToPb(st *CreateStorageConfigurationRequest) (*createStorageConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createStorageConfigurationRequestPb{}
	rootBucketInfoPb, err := rootBucketInfoToPb(&st.RootBucketInfo)
	if err != nil {
		return nil, err
	}
	if rootBucketInfoPb != nil {
		pb.RootBucketInfo = *rootBucketInfoPb
	}

	pb.StorageConfigurationName = st.StorageConfigurationName

	return pb, nil
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

type createStorageConfigurationRequestPb struct {
	// Root S3 bucket information.
	RootBucketInfo rootBucketInfoPb `json:"root_bucket_info"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName string `json:"storage_configuration_name"`
}

func createStorageConfigurationRequestFromPb(pb *createStorageConfigurationRequestPb) (*CreateStorageConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateStorageConfigurationRequest{}
	rootBucketInfoField, err := rootBucketInfoFromPb(&pb.RootBucketInfo)
	if err != nil {
		return nil, err
	}
	if rootBucketInfoField != nil {
		st.RootBucketInfo = *rootBucketInfoField
	}
	st.StorageConfigurationName = pb.StorageConfigurationName

	return st, nil
}

type CreateVpcEndpointRequest struct {
	// The ID of the VPC endpoint object in AWS.
	// Wire name: 'aws_vpc_endpoint_id'
	AwsVpcEndpointId string
	// The Google Cloud specific information for this Private Service Connect
	// endpoint.
	// Wire name: 'gcp_vpc_endpoint_info'
	GcpVpcEndpointInfo *GcpVpcEndpointInfo
	// The AWS region in which this VPC endpoint object exists.
	// Wire name: 'region'
	Region string
	// The human-readable name of the storage configuration.
	// Wire name: 'vpc_endpoint_name'
	VpcEndpointName string

	ForceSendFields []string `tf:"-"`
}

func createVpcEndpointRequestToPb(st *CreateVpcEndpointRequest) (*createVpcEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createVpcEndpointRequestPb{}
	pb.AwsVpcEndpointId = st.AwsVpcEndpointId

	gcpVpcEndpointInfoPb, err := gcpVpcEndpointInfoToPb(st.GcpVpcEndpointInfo)
	if err != nil {
		return nil, err
	}
	if gcpVpcEndpointInfoPb != nil {
		pb.GcpVpcEndpointInfo = gcpVpcEndpointInfoPb
	}

	pb.Region = st.Region

	pb.VpcEndpointName = st.VpcEndpointName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createVpcEndpointRequestPb struct {
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId string `json:"aws_vpc_endpoint_id,omitempty"`
	// The Google Cloud specific information for this Private Service Connect
	// endpoint.
	GcpVpcEndpointInfo *gcpVpcEndpointInfoPb `json:"gcp_vpc_endpoint_info,omitempty"`
	// The AWS region in which this VPC endpoint object exists.
	Region string `json:"region,omitempty"`
	// The human-readable name of the storage configuration.
	VpcEndpointName string `json:"vpc_endpoint_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createVpcEndpointRequestFromPb(pb *createVpcEndpointRequestPb) (*CreateVpcEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateVpcEndpointRequest{}
	st.AwsVpcEndpointId = pb.AwsVpcEndpointId
	gcpVpcEndpointInfoField, err := gcpVpcEndpointInfoFromPb(pb.GcpVpcEndpointInfo)
	if err != nil {
		return nil, err
	}
	if gcpVpcEndpointInfoField != nil {
		st.GcpVpcEndpointInfo = gcpVpcEndpointInfoField
	}
	st.Region = pb.Region
	st.VpcEndpointName = pb.VpcEndpointName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createVpcEndpointRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createVpcEndpointRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateWorkspaceRequest struct {
	// The AWS region of the workspace's data plane.
	// Wire name: 'aws_region'
	AwsRegion string
	// The cloud provider which the workspace uses. For Google Cloud workspaces,
	// always set this field to `gcp`.
	// Wire name: 'cloud'
	Cloud string
	// The general workspace configurations that are specific to cloud
	// providers.
	// Wire name: 'cloud_resource_container'
	CloudResourceContainer *CloudResourceContainer
	// ID of the workspace's credential configuration object.
	// Wire name: 'credentials_id'
	CredentialsId string
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	// Wire name: 'custom_tags'
	CustomTags map[string]string
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
	DeploymentName string
	// The network settings for the workspace. The configurations are only for
	// Databricks-managed VPCs. It is ignored if you specify a customer-managed
	// VPC in the `network_id` field.", All the IP range configurations must be
	// mutually exclusive. An attempt to create a workspace fails if Databricks
	// detects an IP range overlap.
	//
	// Specify custom IP ranges in CIDR format. The IP ranges for these fields
	// must not overlap, and all IP addresses must be entirely within the
	// following ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`,
	// `192.168.0.0/16`, and `240.0.0.0/4`.
	//
	// The sizes of these IP ranges affect the maximum number of nodes for the
	// workspace.
	//
	// **Important**: Confirm the IP ranges used by your Databricks workspace
	// before creating the workspace. You cannot change them after your
	// workspace is deployed. If the IP address ranges for your Databricks are
	// too small, IP exhaustion can occur, causing your Databricks jobs to fail.
	// To determine the address range sizes that you need, Databricks provides a
	// calculator as a Microsoft Excel spreadsheet. See [calculate subnet sizes
	// for a new workspace].
	//
	// [calculate subnet sizes for a new workspace]: https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html
	// Wire name: 'gcp_managed_network_config'
	GcpManagedNetworkConfig *GcpManagedNetworkConfig
	// The configurations for the GKE cluster of a Databricks workspace.
	// Wire name: 'gke_config'
	GkeConfig *GkeConfig
	// Whether no public IP is enabled for the workspace.
	// Wire name: 'is_no_public_ip_enabled'
	IsNoPublicIpEnabled bool
	// The Google Cloud region of the workspace data plane in your Google
	// account. For example, `us-east4`.
	// Wire name: 'location'
	Location string
	// The ID of the workspace's managed services encryption key configuration
	// object. This is used to help protect and control access to the
	// workspace's notebooks, secrets, Databricks SQL queries, and query
	// history. The provided key configuration object property `use_cases` must
	// contain `MANAGED_SERVICES`.
	// Wire name: 'managed_services_customer_managed_key_id'
	ManagedServicesCustomerManagedKeyId string

	// Wire name: 'network_id'
	NetworkId string
	// The pricing tier of the workspace. For pricing tier information, see [AWS
	// Pricing].
	//
	// [AWS Pricing]: https://databricks.com/product/aws-pricing
	// Wire name: 'pricing_tier'
	PricingTier PricingTier
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
	PrivateAccessSettingsId string
	// The ID of the workspace's storage configuration object.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string
	// The ID of the workspace's storage encryption key configuration object.
	// This is used to encrypt the workspace's root S3 bucket (root DBFS and
	// system data) and, optionally, cluster EBS volumes. The provided key
	// configuration object property `use_cases` must contain `STORAGE`.
	// Wire name: 'storage_customer_managed_key_id'
	StorageCustomerManagedKeyId string
	// The workspace's human-readable name.
	// Wire name: 'workspace_name'
	WorkspaceName string

	ForceSendFields []string `tf:"-"`
}

func createWorkspaceRequestToPb(st *CreateWorkspaceRequest) (*createWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createWorkspaceRequestPb{}
	pb.AwsRegion = st.AwsRegion

	pb.Cloud = st.Cloud

	cloudResourceContainerPb, err := cloudResourceContainerToPb(st.CloudResourceContainer)
	if err != nil {
		return nil, err
	}
	if cloudResourceContainerPb != nil {
		pb.CloudResourceContainer = cloudResourceContainerPb
	}

	pb.CredentialsId = st.CredentialsId

	pb.CustomTags = st.CustomTags

	pb.DeploymentName = st.DeploymentName

	gcpManagedNetworkConfigPb, err := gcpManagedNetworkConfigToPb(st.GcpManagedNetworkConfig)
	if err != nil {
		return nil, err
	}
	if gcpManagedNetworkConfigPb != nil {
		pb.GcpManagedNetworkConfig = gcpManagedNetworkConfigPb
	}

	gkeConfigPb, err := gkeConfigToPb(st.GkeConfig)
	if err != nil {
		return nil, err
	}
	if gkeConfigPb != nil {
		pb.GkeConfig = gkeConfigPb
	}

	pb.IsNoPublicIpEnabled = st.IsNoPublicIpEnabled

	pb.Location = st.Location

	pb.ManagedServicesCustomerManagedKeyId = st.ManagedServicesCustomerManagedKeyId

	pb.NetworkId = st.NetworkId

	pb.PricingTier = st.PricingTier

	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	pb.StorageConfigurationId = st.StorageConfigurationId

	pb.StorageCustomerManagedKeyId = st.StorageCustomerManagedKeyId

	pb.WorkspaceName = st.WorkspaceName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createWorkspaceRequestPb struct {
	// The AWS region of the workspace's data plane.
	AwsRegion string `json:"aws_region,omitempty"`
	// The cloud provider which the workspace uses. For Google Cloud workspaces,
	// always set this field to `gcp`.
	Cloud string `json:"cloud,omitempty"`
	// The general workspace configurations that are specific to cloud
	// providers.
	CloudResourceContainer *cloudResourceContainerPb `json:"cloud_resource_container,omitempty"`
	// ID of the workspace's credential configuration object.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
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
	DeploymentName string `json:"deployment_name,omitempty"`
	// The network settings for the workspace. The configurations are only for
	// Databricks-managed VPCs. It is ignored if you specify a customer-managed
	// VPC in the `network_id` field.", All the IP range configurations must be
	// mutually exclusive. An attempt to create a workspace fails if Databricks
	// detects an IP range overlap.
	//
	// Specify custom IP ranges in CIDR format. The IP ranges for these fields
	// must not overlap, and all IP addresses must be entirely within the
	// following ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`,
	// `192.168.0.0/16`, and `240.0.0.0/4`.
	//
	// The sizes of these IP ranges affect the maximum number of nodes for the
	// workspace.
	//
	// **Important**: Confirm the IP ranges used by your Databricks workspace
	// before creating the workspace. You cannot change them after your
	// workspace is deployed. If the IP address ranges for your Databricks are
	// too small, IP exhaustion can occur, causing your Databricks jobs to fail.
	// To determine the address range sizes that you need, Databricks provides a
	// calculator as a Microsoft Excel spreadsheet. See [calculate subnet sizes
	// for a new workspace].
	//
	// [calculate subnet sizes for a new workspace]: https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html
	GcpManagedNetworkConfig *gcpManagedNetworkConfigPb `json:"gcp_managed_network_config,omitempty"`
	// The configurations for the GKE cluster of a Databricks workspace.
	GkeConfig *gkeConfigPb `json:"gke_config,omitempty"`
	// Whether no public IP is enabled for the workspace.
	IsNoPublicIpEnabled bool `json:"is_no_public_ip_enabled,omitempty"`
	// The Google Cloud region of the workspace data plane in your Google
	// account. For example, `us-east4`.
	Location string `json:"location,omitempty"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This is used to help protect and control access to the
	// workspace's notebooks, secrets, Databricks SQL queries, and query
	// history. The provided key configuration object property `use_cases` must
	// contain `MANAGED_SERVICES`.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`

	NetworkId string `json:"network_id,omitempty"`
	// The pricing tier of the workspace. For pricing tier information, see [AWS
	// Pricing].
	//
	// [AWS Pricing]: https://databricks.com/product/aws-pricing
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
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// The ID of the workspace's storage configuration object.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The ID of the workspace's storage encryption key configuration object.
	// This is used to encrypt the workspace's root S3 bucket (root DBFS and
	// system data) and, optionally, cluster EBS volumes. The provided key
	// configuration object property `use_cases` must contain `STORAGE`.
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// The workspace's human-readable name.
	WorkspaceName string `json:"workspace_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createWorkspaceRequestFromPb(pb *createWorkspaceRequestPb) (*CreateWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateWorkspaceRequest{}
	st.AwsRegion = pb.AwsRegion
	st.Cloud = pb.Cloud
	cloudResourceContainerField, err := cloudResourceContainerFromPb(pb.CloudResourceContainer)
	if err != nil {
		return nil, err
	}
	if cloudResourceContainerField != nil {
		st.CloudResourceContainer = cloudResourceContainerField
	}
	st.CredentialsId = pb.CredentialsId
	st.CustomTags = pb.CustomTags
	st.DeploymentName = pb.DeploymentName
	gcpManagedNetworkConfigField, err := gcpManagedNetworkConfigFromPb(pb.GcpManagedNetworkConfig)
	if err != nil {
		return nil, err
	}
	if gcpManagedNetworkConfigField != nil {
		st.GcpManagedNetworkConfig = gcpManagedNetworkConfigField
	}
	gkeConfigField, err := gkeConfigFromPb(pb.GkeConfig)
	if err != nil {
		return nil, err
	}
	if gkeConfigField != nil {
		st.GkeConfig = gkeConfigField
	}
	st.IsNoPublicIpEnabled = pb.IsNoPublicIpEnabled
	st.Location = pb.Location
	st.ManagedServicesCustomerManagedKeyId = pb.ManagedServicesCustomerManagedKeyId
	st.NetworkId = pb.NetworkId
	st.PricingTier = pb.PricingTier
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId
	st.StorageConfigurationId = pb.StorageConfigurationId
	st.StorageCustomerManagedKeyId = pb.StorageCustomerManagedKeyId
	st.WorkspaceName = pb.WorkspaceName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createWorkspaceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createWorkspaceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Credential struct {
	// The Databricks account ID that hosts the credential.
	// Wire name: 'account_id'
	AccountId string

	// Wire name: 'aws_credentials'
	AwsCredentials *AwsCredentials
	// Time in epoch milliseconds when the credential was created.
	// Wire name: 'creation_time'
	CreationTime int64
	// Databricks credential configuration ID.
	// Wire name: 'credentials_id'
	CredentialsId string
	// The human-readable name of the credential configuration object.
	// Wire name: 'credentials_name'
	CredentialsName string

	ForceSendFields []string `tf:"-"`
}

func credentialToPb(st *Credential) (*credentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &credentialPb{}
	pb.AccountId = st.AccountId

	awsCredentialsPb, err := awsCredentialsToPb(st.AwsCredentials)
	if err != nil {
		return nil, err
	}
	if awsCredentialsPb != nil {
		pb.AwsCredentials = awsCredentialsPb
	}

	pb.CreationTime = st.CreationTime

	pb.CredentialsId = st.CredentialsId

	pb.CredentialsName = st.CredentialsName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type credentialPb struct {
	// The Databricks account ID that hosts the credential.
	AccountId string `json:"account_id,omitempty"`

	AwsCredentials *awsCredentialsPb `json:"aws_credentials,omitempty"`
	// Time in epoch milliseconds when the credential was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Databricks credential configuration ID.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The human-readable name of the credential configuration object.
	CredentialsName string `json:"credentials_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func credentialFromPb(pb *credentialPb) (*Credential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Credential{}
	st.AccountId = pb.AccountId
	awsCredentialsField, err := awsCredentialsFromPb(pb.AwsCredentials)
	if err != nil {
		return nil, err
	}
	if awsCredentialsField != nil {
		st.AwsCredentials = awsCredentialsField
	}
	st.CreationTime = pb.CreationTime
	st.CredentialsId = pb.CredentialsId
	st.CredentialsName = pb.CredentialsName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *credentialPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st credentialPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The custom tags key-value pairing that is attached to this workspace. The
// key-value pair is a string of utf-8 characters. The value can be an empty
// string, with maximum length of 255 characters. The key can be of maximum
// length of 127 characters, and cannot be empty.

type CustomTags map[string]string
type customTagsPb CustomTags

func customTagsToPb(st *CustomTags) (*customTagsPb, error) {
	if st == nil {
		return nil, nil
	}
	stPb := customTagsPb(*st)
	return &stPb, nil
}
func customTagsFromPb(stPb *customTagsPb) (*CustomTags, error) {
	if stPb == nil {
		return nil, nil
	}
	st := CustomTags(*stPb)
	return &st, nil
}

// The general workspace configurations that are specific to Google Cloud.
type CustomerFacingGcpCloudResourceContainer struct {
	// The Google Cloud project ID, which the workspace uses to instantiate
	// cloud resources for your workspace.
	// Wire name: 'project_id'
	ProjectId string

	ForceSendFields []string `tf:"-"`
}

func customerFacingGcpCloudResourceContainerToPb(st *CustomerFacingGcpCloudResourceContainer) (*customerFacingGcpCloudResourceContainerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &customerFacingGcpCloudResourceContainerPb{}
	pb.ProjectId = st.ProjectId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type customerFacingGcpCloudResourceContainerPb struct {
	// The Google Cloud project ID, which the workspace uses to instantiate
	// cloud resources for your workspace.
	ProjectId string `json:"project_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func customerFacingGcpCloudResourceContainerFromPb(pb *customerFacingGcpCloudResourceContainerPb) (*CustomerFacingGcpCloudResourceContainer, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomerFacingGcpCloudResourceContainer{}
	st.ProjectId = pb.ProjectId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *customerFacingGcpCloudResourceContainerPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st customerFacingGcpCloudResourceContainerPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CustomerManagedKey struct {
	// The Databricks account ID that holds the customer-managed key.
	// Wire name: 'account_id'
	AccountId string

	// Wire name: 'aws_key_info'
	AwsKeyInfo *AwsKeyInfo
	// Time in epoch milliseconds when the customer key was created.
	// Wire name: 'creation_time'
	CreationTime int64
	// ID of the encryption key configuration object.
	// Wire name: 'customer_managed_key_id'
	CustomerManagedKeyId string

	// Wire name: 'gcp_key_info'
	GcpKeyInfo *GcpKeyInfo
	// The cases that the key can be used for.
	// Wire name: 'use_cases'
	UseCases []KeyUseCase

	ForceSendFields []string `tf:"-"`
}

func customerManagedKeyToPb(st *CustomerManagedKey) (*customerManagedKeyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &customerManagedKeyPb{}
	pb.AccountId = st.AccountId

	awsKeyInfoPb, err := awsKeyInfoToPb(st.AwsKeyInfo)
	if err != nil {
		return nil, err
	}
	if awsKeyInfoPb != nil {
		pb.AwsKeyInfo = awsKeyInfoPb
	}

	pb.CreationTime = st.CreationTime

	pb.CustomerManagedKeyId = st.CustomerManagedKeyId

	gcpKeyInfoPb, err := gcpKeyInfoToPb(st.GcpKeyInfo)
	if err != nil {
		return nil, err
	}
	if gcpKeyInfoPb != nil {
		pb.GcpKeyInfo = gcpKeyInfoPb
	}

	pb.UseCases = st.UseCases

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type customerManagedKeyPb struct {
	// The Databricks account ID that holds the customer-managed key.
	AccountId string `json:"account_id,omitempty"`

	AwsKeyInfo *awsKeyInfoPb `json:"aws_key_info,omitempty"`
	// Time in epoch milliseconds when the customer key was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// ID of the encryption key configuration object.
	CustomerManagedKeyId string `json:"customer_managed_key_id,omitempty"`

	GcpKeyInfo *gcpKeyInfoPb `json:"gcp_key_info,omitempty"`
	// The cases that the key can be used for.
	UseCases []KeyUseCase `json:"use_cases,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func customerManagedKeyFromPb(pb *customerManagedKeyPb) (*CustomerManagedKey, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomerManagedKey{}
	st.AccountId = pb.AccountId
	awsKeyInfoField, err := awsKeyInfoFromPb(pb.AwsKeyInfo)
	if err != nil {
		return nil, err
	}
	if awsKeyInfoField != nil {
		st.AwsKeyInfo = awsKeyInfoField
	}
	st.CreationTime = pb.CreationTime
	st.CustomerManagedKeyId = pb.CustomerManagedKeyId
	gcpKeyInfoField, err := gcpKeyInfoFromPb(pb.GcpKeyInfo)
	if err != nil {
		return nil, err
	}
	if gcpKeyInfoField != nil {
		st.GcpKeyInfo = gcpKeyInfoField
	}
	st.UseCases = pb.UseCases

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *customerManagedKeyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st customerManagedKeyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete credential configuration
type DeleteCredentialRequest struct {
	// Databricks Account API credential configuration ID
	// Wire name: 'credentials_id'
	CredentialsId string `tf:"-"`
}

func deleteCredentialRequestToPb(st *DeleteCredentialRequest) (*deleteCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCredentialRequestPb{}
	pb.CredentialsId = st.CredentialsId

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
	// Databricks Account API credential configuration ID
	CredentialsId string `json:"-" url:"-"`
}

func deleteCredentialRequestFromPb(pb *deleteCredentialRequestPb) (*DeleteCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCredentialRequest{}
	st.CredentialsId = pb.CredentialsId

	return st, nil
}

// Delete encryption key configuration
type DeleteEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	// Wire name: 'customer_managed_key_id'
	CustomerManagedKeyId string `tf:"-"`
}

func deleteEncryptionKeyRequestToPb(st *DeleteEncryptionKeyRequest) (*deleteEncryptionKeyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteEncryptionKeyRequestPb{}
	pb.CustomerManagedKeyId = st.CustomerManagedKeyId

	return pb, nil
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

type deleteEncryptionKeyRequestPb struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId string `json:"-" url:"-"`
}

func deleteEncryptionKeyRequestFromPb(pb *deleteEncryptionKeyRequestPb) (*DeleteEncryptionKeyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteEncryptionKeyRequest{}
	st.CustomerManagedKeyId = pb.CustomerManagedKeyId

	return st, nil
}

// Delete a network configuration
type DeleteNetworkRequest struct {
	// Databricks Account API network configuration ID.
	// Wire name: 'network_id'
	NetworkId string `tf:"-"`
}

func deleteNetworkRequestToPb(st *DeleteNetworkRequest) (*deleteNetworkRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteNetworkRequestPb{}
	pb.NetworkId = st.NetworkId

	return pb, nil
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

type deleteNetworkRequestPb struct {
	// Databricks Account API network configuration ID.
	NetworkId string `json:"-" url:"-"`
}

func deleteNetworkRequestFromPb(pb *deleteNetworkRequestPb) (*DeleteNetworkRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteNetworkRequest{}
	st.NetworkId = pb.NetworkId

	return st, nil
}

// Delete a private access settings object
type DeletePrivateAccesRequest struct {
	// Databricks Account API private access settings ID.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string `tf:"-"`
}

func deletePrivateAccesRequestToPb(st *DeletePrivateAccesRequest) (*deletePrivateAccesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePrivateAccesRequestPb{}
	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	return pb, nil
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

type deletePrivateAccesRequestPb struct {
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string `json:"-" url:"-"`
}

func deletePrivateAccesRequestFromPb(pb *deletePrivateAccesRequestPb) (*DeletePrivateAccesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePrivateAccesRequest{}
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId

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

// Delete storage configuration
type DeleteStorageRequest struct {
	// Databricks Account API storage configuration ID.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string `tf:"-"`
}

func deleteStorageRequestToPb(st *DeleteStorageRequest) (*deleteStorageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteStorageRequestPb{}
	pb.StorageConfigurationId = st.StorageConfigurationId

	return pb, nil
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

type deleteStorageRequestPb struct {
	// Databricks Account API storage configuration ID.
	StorageConfigurationId string `json:"-" url:"-"`
}

func deleteStorageRequestFromPb(pb *deleteStorageRequestPb) (*DeleteStorageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteStorageRequest{}
	st.StorageConfigurationId = pb.StorageConfigurationId

	return st, nil
}

// Delete VPC endpoint configuration
type DeleteVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	// Wire name: 'vpc_endpoint_id'
	VpcEndpointId string `tf:"-"`
}

func deleteVpcEndpointRequestToPb(st *DeleteVpcEndpointRequest) (*deleteVpcEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteVpcEndpointRequestPb{}
	pb.VpcEndpointId = st.VpcEndpointId

	return pb, nil
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

type deleteVpcEndpointRequestPb struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId string `json:"-" url:"-"`
}

func deleteVpcEndpointRequestFromPb(pb *deleteVpcEndpointRequestPb) (*DeleteVpcEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteVpcEndpointRequest{}
	st.VpcEndpointId = pb.VpcEndpointId

	return st, nil
}

// Delete a workspace
type DeleteWorkspaceRequest struct {
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func deleteWorkspaceRequestToPb(st *DeleteWorkspaceRequest) (*deleteWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteWorkspaceRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
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

type deleteWorkspaceRequestPb struct {
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func deleteWorkspaceRequestFromPb(pb *deleteWorkspaceRequestPb) (*DeleteWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteWorkspaceRequest{}
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

// This enumeration represents the type of Databricks VPC [endpoint service]
// that was used when creating this VPC endpoint.
//
// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
type EndpointUseCase string
type endpointUseCasePb string

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

// Type always returns EndpointUseCase to satisfy [pflag.Value] interface
func (f *EndpointUseCase) Type() string {
	return "EndpointUseCase"
}

func endpointUseCaseToPb(st *EndpointUseCase) (*endpointUseCasePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := endpointUseCasePb(*st)
	return &pb, nil
}

func endpointUseCaseFromPb(pb *endpointUseCasePb) (*EndpointUseCase, error) {
	if pb == nil {
		return nil, nil
	}
	st := EndpointUseCase(*pb)
	return &st, nil
}

// The AWS resource associated with this error: credentials, VPC, subnet,
// security group, or network ACL.
type ErrorType string
type errorTypePb string

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

// Type always returns ErrorType to satisfy [pflag.Value] interface
func (f *ErrorType) Type() string {
	return "ErrorType"
}

func errorTypeToPb(st *ErrorType) (*errorTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := errorTypePb(*st)
	return &pb, nil
}

func errorTypeFromPb(pb *errorTypePb) (*ErrorType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ErrorType(*pb)
	return &st, nil
}

type ExternalCustomerInfo struct {
	// Email of the authoritative user.
	// Wire name: 'authoritative_user_email'
	AuthoritativeUserEmail string
	// The authoritative user full name.
	// Wire name: 'authoritative_user_full_name'
	AuthoritativeUserFullName string
	// The legal entity name for the external workspace
	// Wire name: 'customer_name'
	CustomerName string

	ForceSendFields []string `tf:"-"`
}

func externalCustomerInfoToPb(st *ExternalCustomerInfo) (*externalCustomerInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalCustomerInfoPb{}
	pb.AuthoritativeUserEmail = st.AuthoritativeUserEmail

	pb.AuthoritativeUserFullName = st.AuthoritativeUserFullName

	pb.CustomerName = st.CustomerName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type externalCustomerInfoPb struct {
	// Email of the authoritative user.
	AuthoritativeUserEmail string `json:"authoritative_user_email,omitempty"`
	// The authoritative user full name.
	AuthoritativeUserFullName string `json:"authoritative_user_full_name,omitempty"`
	// The legal entity name for the external workspace
	CustomerName string `json:"customer_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func externalCustomerInfoFromPb(pb *externalCustomerInfoPb) (*ExternalCustomerInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalCustomerInfo{}
	st.AuthoritativeUserEmail = pb.AuthoritativeUserEmail
	st.AuthoritativeUserFullName = pb.AuthoritativeUserFullName
	st.CustomerName = pb.CustomerName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *externalCustomerInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st externalCustomerInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GcpKeyInfo struct {
	// The GCP KMS key's resource name
	// Wire name: 'kms_key_id'
	KmsKeyId string
}

func gcpKeyInfoToPb(st *GcpKeyInfo) (*gcpKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpKeyInfoPb{}
	pb.KmsKeyId = st.KmsKeyId

	return pb, nil
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

type gcpKeyInfoPb struct {
	// The GCP KMS key's resource name
	KmsKeyId string `json:"kms_key_id"`
}

func gcpKeyInfoFromPb(pb *gcpKeyInfoPb) (*GcpKeyInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpKeyInfo{}
	st.KmsKeyId = pb.KmsKeyId

	return st, nil
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
	GkeClusterPodIpRange string
	// The IP range from which to allocate GKE cluster services. No bigger than
	// `/16` and no smaller than `/27`.
	// Wire name: 'gke_cluster_service_ip_range'
	GkeClusterServiceIpRange string
	// The IP range from which to allocate GKE cluster nodes. No bigger than
	// `/9` and no smaller than `/29`.
	// Wire name: 'subnet_cidr'
	SubnetCidr string

	ForceSendFields []string `tf:"-"`
}

func gcpManagedNetworkConfigToPb(st *GcpManagedNetworkConfig) (*gcpManagedNetworkConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpManagedNetworkConfigPb{}
	pb.GkeClusterPodIpRange = st.GkeClusterPodIpRange

	pb.GkeClusterServiceIpRange = st.GkeClusterServiceIpRange

	pb.SubnetCidr = st.SubnetCidr

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type gcpManagedNetworkConfigPb struct {
	// The IP range from which to allocate GKE cluster pods. No bigger than `/9`
	// and no smaller than `/21`.
	GkeClusterPodIpRange string `json:"gke_cluster_pod_ip_range,omitempty"`
	// The IP range from which to allocate GKE cluster services. No bigger than
	// `/16` and no smaller than `/27`.
	GkeClusterServiceIpRange string `json:"gke_cluster_service_ip_range,omitempty"`
	// The IP range from which to allocate GKE cluster nodes. No bigger than
	// `/9` and no smaller than `/29`.
	SubnetCidr string `json:"subnet_cidr,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func gcpManagedNetworkConfigFromPb(pb *gcpManagedNetworkConfigPb) (*GcpManagedNetworkConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpManagedNetworkConfig{}
	st.GkeClusterPodIpRange = pb.GkeClusterPodIpRange
	st.GkeClusterServiceIpRange = pb.GkeClusterServiceIpRange
	st.SubnetCidr = pb.SubnetCidr

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *gcpManagedNetworkConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st gcpManagedNetworkConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The Google Cloud specific information for this network (for example, the VPC
// ID, subnet ID, and secondary IP ranges).
type GcpNetworkInfo struct {
	// The Google Cloud project ID of the VPC network.
	// Wire name: 'network_project_id'
	NetworkProjectId string
	// The name of the secondary IP range for pods. A Databricks-managed GKE
	// cluster uses this IP range for its pods. This secondary IP range can be
	// used by only one workspace.
	// Wire name: 'pod_ip_range_name'
	PodIpRangeName string
	// The name of the secondary IP range for services. A Databricks-managed GKE
	// cluster uses this IP range for its services. This secondary IP range can
	// be used by only one workspace.
	// Wire name: 'service_ip_range_name'
	ServiceIpRangeName string
	// The ID of the subnet associated with this network.
	// Wire name: 'subnet_id'
	SubnetId string
	// The Google Cloud region of the workspace data plane (for example,
	// `us-east4`).
	// Wire name: 'subnet_region'
	SubnetRegion string
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	// Wire name: 'vpc_id'
	VpcId string
}

func gcpNetworkInfoToPb(st *GcpNetworkInfo) (*gcpNetworkInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpNetworkInfoPb{}
	pb.NetworkProjectId = st.NetworkProjectId

	pb.PodIpRangeName = st.PodIpRangeName

	pb.ServiceIpRangeName = st.ServiceIpRangeName

	pb.SubnetId = st.SubnetId

	pb.SubnetRegion = st.SubnetRegion

	pb.VpcId = st.VpcId

	return pb, nil
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

type gcpNetworkInfoPb struct {
	// The Google Cloud project ID of the VPC network.
	NetworkProjectId string `json:"network_project_id"`
	// The name of the secondary IP range for pods. A Databricks-managed GKE
	// cluster uses this IP range for its pods. This secondary IP range can be
	// used by only one workspace.
	PodIpRangeName string `json:"pod_ip_range_name"`
	// The name of the secondary IP range for services. A Databricks-managed GKE
	// cluster uses this IP range for its services. This secondary IP range can
	// be used by only one workspace.
	ServiceIpRangeName string `json:"service_ip_range_name"`
	// The ID of the subnet associated with this network.
	SubnetId string `json:"subnet_id"`
	// The Google Cloud region of the workspace data plane (for example,
	// `us-east4`).
	SubnetRegion string `json:"subnet_region"`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId string `json:"vpc_id"`
}

func gcpNetworkInfoFromPb(pb *gcpNetworkInfoPb) (*GcpNetworkInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpNetworkInfo{}
	st.NetworkProjectId = pb.NetworkProjectId
	st.PodIpRangeName = pb.PodIpRangeName
	st.ServiceIpRangeName = pb.ServiceIpRangeName
	st.SubnetId = pb.SubnetId
	st.SubnetRegion = pb.SubnetRegion
	st.VpcId = pb.VpcId

	return st, nil
}

// The Google Cloud specific information for this Private Service Connect
// endpoint.
type GcpVpcEndpointInfo struct {
	// Region of the PSC endpoint.
	// Wire name: 'endpoint_region'
	EndpointRegion string
	// The Google Cloud project ID of the VPC network where the PSC connection
	// resides.
	// Wire name: 'project_id'
	ProjectId string
	// The unique ID of this PSC connection.
	// Wire name: 'psc_connection_id'
	PscConnectionId string
	// The name of the PSC endpoint in the Google Cloud project.
	// Wire name: 'psc_endpoint_name'
	PscEndpointName string
	// The service attachment this PSC connection connects to.
	// Wire name: 'service_attachment_id'
	ServiceAttachmentId string

	ForceSendFields []string `tf:"-"`
}

func gcpVpcEndpointInfoToPb(st *GcpVpcEndpointInfo) (*gcpVpcEndpointInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpVpcEndpointInfoPb{}
	pb.EndpointRegion = st.EndpointRegion

	pb.ProjectId = st.ProjectId

	pb.PscConnectionId = st.PscConnectionId

	pb.PscEndpointName = st.PscEndpointName

	pb.ServiceAttachmentId = st.ServiceAttachmentId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type gcpVpcEndpointInfoPb struct {
	// Region of the PSC endpoint.
	EndpointRegion string `json:"endpoint_region"`
	// The Google Cloud project ID of the VPC network where the PSC connection
	// resides.
	ProjectId string `json:"project_id"`
	// The unique ID of this PSC connection.
	PscConnectionId string `json:"psc_connection_id,omitempty"`
	// The name of the PSC endpoint in the Google Cloud project.
	PscEndpointName string `json:"psc_endpoint_name"`
	// The service attachment this PSC connection connects to.
	ServiceAttachmentId string `json:"service_attachment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func gcpVpcEndpointInfoFromPb(pb *gcpVpcEndpointInfoPb) (*GcpVpcEndpointInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpVpcEndpointInfo{}
	st.EndpointRegion = pb.EndpointRegion
	st.ProjectId = pb.ProjectId
	st.PscConnectionId = pb.PscConnectionId
	st.PscEndpointName = pb.PscEndpointName
	st.ServiceAttachmentId = pb.ServiceAttachmentId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *gcpVpcEndpointInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st gcpVpcEndpointInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get credential configuration
type GetCredentialRequest struct {
	// Databricks Account API credential configuration ID
	// Wire name: 'credentials_id'
	CredentialsId string `tf:"-"`
}

func getCredentialRequestToPb(st *GetCredentialRequest) (*getCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCredentialRequestPb{}
	pb.CredentialsId = st.CredentialsId

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
	// Databricks Account API credential configuration ID
	CredentialsId string `json:"-" url:"-"`
}

func getCredentialRequestFromPb(pb *getCredentialRequestPb) (*GetCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCredentialRequest{}
	st.CredentialsId = pb.CredentialsId

	return st, nil
}

// Get encryption key configuration
type GetEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	// Wire name: 'customer_managed_key_id'
	CustomerManagedKeyId string `tf:"-"`
}

func getEncryptionKeyRequestToPb(st *GetEncryptionKeyRequest) (*getEncryptionKeyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEncryptionKeyRequestPb{}
	pb.CustomerManagedKeyId = st.CustomerManagedKeyId

	return pb, nil
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

type getEncryptionKeyRequestPb struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId string `json:"-" url:"-"`
}

func getEncryptionKeyRequestFromPb(pb *getEncryptionKeyRequestPb) (*GetEncryptionKeyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEncryptionKeyRequest{}
	st.CustomerManagedKeyId = pb.CustomerManagedKeyId

	return st, nil
}

// Get a network configuration
type GetNetworkRequest struct {
	// Databricks Account API network configuration ID.
	// Wire name: 'network_id'
	NetworkId string `tf:"-"`
}

func getNetworkRequestToPb(st *GetNetworkRequest) (*getNetworkRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getNetworkRequestPb{}
	pb.NetworkId = st.NetworkId

	return pb, nil
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

type getNetworkRequestPb struct {
	// Databricks Account API network configuration ID.
	NetworkId string `json:"-" url:"-"`
}

func getNetworkRequestFromPb(pb *getNetworkRequestPb) (*GetNetworkRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetNetworkRequest{}
	st.NetworkId = pb.NetworkId

	return st, nil
}

// Get a private access settings object
type GetPrivateAccesRequest struct {
	// Databricks Account API private access settings ID.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string `tf:"-"`
}

func getPrivateAccesRequestToPb(st *GetPrivateAccesRequest) (*getPrivateAccesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPrivateAccesRequestPb{}
	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	return pb, nil
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

type getPrivateAccesRequestPb struct {
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string `json:"-" url:"-"`
}

func getPrivateAccesRequestFromPb(pb *getPrivateAccesRequestPb) (*GetPrivateAccesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPrivateAccesRequest{}
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId

	return st, nil
}

// Get storage configuration
type GetStorageRequest struct {
	// Databricks Account API storage configuration ID.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string `tf:"-"`
}

func getStorageRequestToPb(st *GetStorageRequest) (*getStorageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStorageRequestPb{}
	pb.StorageConfigurationId = st.StorageConfigurationId

	return pb, nil
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

type getStorageRequestPb struct {
	// Databricks Account API storage configuration ID.
	StorageConfigurationId string `json:"-" url:"-"`
}

func getStorageRequestFromPb(pb *getStorageRequestPb) (*GetStorageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStorageRequest{}
	st.StorageConfigurationId = pb.StorageConfigurationId

	return st, nil
}

// Get a VPC endpoint configuration
type GetVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	// Wire name: 'vpc_endpoint_id'
	VpcEndpointId string `tf:"-"`
}

func getVpcEndpointRequestToPb(st *GetVpcEndpointRequest) (*getVpcEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getVpcEndpointRequestPb{}
	pb.VpcEndpointId = st.VpcEndpointId

	return pb, nil
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

type getVpcEndpointRequestPb struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId string `json:"-" url:"-"`
}

func getVpcEndpointRequestFromPb(pb *getVpcEndpointRequestPb) (*GetVpcEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetVpcEndpointRequest{}
	st.VpcEndpointId = pb.VpcEndpointId

	return st, nil
}

// Get a workspace
type GetWorkspaceRequest struct {
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func getWorkspaceRequestToPb(st *GetWorkspaceRequest) (*getWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
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

type getWorkspaceRequestPb struct {
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func getWorkspaceRequestFromPb(pb *getWorkspaceRequestPb) (*GetWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceRequest{}
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
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
	ConnectivityType GkeConfigConnectivityType
	// The IP range from which to allocate GKE cluster master resources. This
	// field will be ignored if GKE private cluster is not enabled.
	//
	// It must be exactly as big as `/28`.
	// Wire name: 'master_ip_range'
	MasterIpRange string

	ForceSendFields []string `tf:"-"`
}

func gkeConfigToPb(st *GkeConfig) (*gkeConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gkeConfigPb{}
	pb.ConnectivityType = st.ConnectivityType

	pb.MasterIpRange = st.MasterIpRange

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type gkeConfigPb struct {
	// Specifies the network connectivity types for the GKE nodes and the GKE
	// master network.
	//
	// Set to `PRIVATE_NODE_PUBLIC_MASTER` for a private GKE cluster for the
	// workspace. The GKE nodes will not have public IPs.
	//
	// Set to `PUBLIC_NODE_PUBLIC_MASTER` for a public GKE cluster. The nodes of
	// a public GKE cluster have public IP addresses.
	ConnectivityType GkeConfigConnectivityType `json:"connectivity_type,omitempty"`
	// The IP range from which to allocate GKE cluster master resources. This
	// field will be ignored if GKE private cluster is not enabled.
	//
	// It must be exactly as big as `/28`.
	MasterIpRange string `json:"master_ip_range,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func gkeConfigFromPb(pb *gkeConfigPb) (*GkeConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GkeConfig{}
	st.ConnectivityType = pb.ConnectivityType
	st.MasterIpRange = pb.MasterIpRange

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *gkeConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st gkeConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
type gkeConfigConnectivityTypePb string

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

// Type always returns GkeConfigConnectivityType to satisfy [pflag.Value] interface
func (f *GkeConfigConnectivityType) Type() string {
	return "GkeConfigConnectivityType"
}

func gkeConfigConnectivityTypeToPb(st *GkeConfigConnectivityType) (*gkeConfigConnectivityTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := gkeConfigConnectivityTypePb(*st)
	return &pb, nil
}

func gkeConfigConnectivityTypeFromPb(pb *gkeConfigConnectivityTypePb) (*GkeConfigConnectivityType, error) {
	if pb == nil {
		return nil, nil
	}
	st := GkeConfigConnectivityType(*pb)
	return &st, nil
}

// Possible values are: * `MANAGED_SERVICES`: Encrypts notebook and secret data
// in the control plane * `STORAGE`: Encrypts the workspace's root S3 bucket
// (root DBFS and system data) and, optionally, cluster EBS volumes.
type KeyUseCase string
type keyUseCasePb string

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

// Type always returns KeyUseCase to satisfy [pflag.Value] interface
func (f *KeyUseCase) Type() string {
	return "KeyUseCase"
}

func keyUseCaseToPb(st *KeyUseCase) (*keyUseCasePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := keyUseCasePb(*st)
	return &pb, nil
}

func keyUseCaseFromPb(pb *keyUseCasePb) (*KeyUseCase, error) {
	if pb == nil {
		return nil, nil
	}
	st := KeyUseCase(*pb)
	return &st, nil
}

type Network struct {
	// The Databricks account ID associated with this network configuration.
	// Wire name: 'account_id'
	AccountId string
	// Time in epoch milliseconds when the network was created.
	// Wire name: 'creation_time'
	CreationTime int64
	// Array of error messages about the network configuration.
	// Wire name: 'error_messages'
	ErrorMessages []NetworkHealth
	// The Google Cloud specific information for this network (for example, the
	// VPC ID, subnet ID, and secondary IP ranges).
	// Wire name: 'gcp_network_info'
	GcpNetworkInfo *GcpNetworkInfo
	// The Databricks network configuration ID.
	// Wire name: 'network_id'
	NetworkId string
	// The human-readable name of the network configuration.
	// Wire name: 'network_name'
	NetworkName string

	// Wire name: 'security_group_ids'
	SecurityGroupIds []string

	// Wire name: 'subnet_ids'
	SubnetIds []string
	// If specified, contains the VPC endpoints used to allow cluster
	// communication from this VPC over [AWS PrivateLink].
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// Wire name: 'vpc_endpoints'
	VpcEndpoints *NetworkVpcEndpoints
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	// Wire name: 'vpc_id'
	VpcId string
	// The status of this network configuration object in terms of its use in a
	// workspace: * `UNATTACHED`: Unattached. * `VALID`: Valid. * `BROKEN`:
	// Broken. * `WARNED`: Warned.
	// Wire name: 'vpc_status'
	VpcStatus VpcStatus
	// Array of warning messages about the network configuration.
	// Wire name: 'warning_messages'
	WarningMessages []NetworkWarning
	// Workspace ID associated with this network configuration.
	// Wire name: 'workspace_id'
	WorkspaceId int64

	ForceSendFields []string `tf:"-"`
}

func networkToPb(st *Network) (*networkPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &networkPb{}
	pb.AccountId = st.AccountId

	pb.CreationTime = st.CreationTime

	var errorMessagesPb []networkHealthPb
	for _, item := range st.ErrorMessages {
		itemPb, err := networkHealthToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			errorMessagesPb = append(errorMessagesPb, *itemPb)
		}
	}
	pb.ErrorMessages = errorMessagesPb

	gcpNetworkInfoPb, err := gcpNetworkInfoToPb(st.GcpNetworkInfo)
	if err != nil {
		return nil, err
	}
	if gcpNetworkInfoPb != nil {
		pb.GcpNetworkInfo = gcpNetworkInfoPb
	}

	pb.NetworkId = st.NetworkId

	pb.NetworkName = st.NetworkName

	pb.SecurityGroupIds = st.SecurityGroupIds

	pb.SubnetIds = st.SubnetIds

	vpcEndpointsPb, err := networkVpcEndpointsToPb(st.VpcEndpoints)
	if err != nil {
		return nil, err
	}
	if vpcEndpointsPb != nil {
		pb.VpcEndpoints = vpcEndpointsPb
	}

	pb.VpcId = st.VpcId

	pb.VpcStatus = st.VpcStatus

	var warningMessagesPb []networkWarningPb
	for _, item := range st.WarningMessages {
		itemPb, err := networkWarningToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			warningMessagesPb = append(warningMessagesPb, *itemPb)
		}
	}
	pb.WarningMessages = warningMessagesPb

	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type networkPb struct {
	// The Databricks account ID associated with this network configuration.
	AccountId string `json:"account_id,omitempty"`
	// Time in epoch milliseconds when the network was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Array of error messages about the network configuration.
	ErrorMessages []networkHealthPb `json:"error_messages,omitempty"`
	// The Google Cloud specific information for this network (for example, the
	// VPC ID, subnet ID, and secondary IP ranges).
	GcpNetworkInfo *gcpNetworkInfoPb `json:"gcp_network_info,omitempty"`
	// The Databricks network configuration ID.
	NetworkId string `json:"network_id,omitempty"`
	// The human-readable name of the network configuration.
	NetworkName string `json:"network_name,omitempty"`

	SecurityGroupIds []string `json:"security_group_ids,omitempty"`

	SubnetIds []string `json:"subnet_ids,omitempty"`
	// If specified, contains the VPC endpoints used to allow cluster
	// communication from this VPC over [AWS PrivateLink].
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	VpcEndpoints *networkVpcEndpointsPb `json:"vpc_endpoints,omitempty"`
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	VpcId string `json:"vpc_id,omitempty"`
	// The status of this network configuration object in terms of its use in a
	// workspace: * `UNATTACHED`: Unattached. * `VALID`: Valid. * `BROKEN`:
	// Broken. * `WARNED`: Warned.
	VpcStatus VpcStatus `json:"vpc_status,omitempty"`
	// Array of warning messages about the network configuration.
	WarningMessages []networkWarningPb `json:"warning_messages,omitempty"`
	// Workspace ID associated with this network configuration.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func networkFromPb(pb *networkPb) (*Network, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Network{}
	st.AccountId = pb.AccountId
	st.CreationTime = pb.CreationTime

	var errorMessagesField []NetworkHealth
	for _, itemPb := range pb.ErrorMessages {
		item, err := networkHealthFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			errorMessagesField = append(errorMessagesField, *item)
		}
	}
	st.ErrorMessages = errorMessagesField
	gcpNetworkInfoField, err := gcpNetworkInfoFromPb(pb.GcpNetworkInfo)
	if err != nil {
		return nil, err
	}
	if gcpNetworkInfoField != nil {
		st.GcpNetworkInfo = gcpNetworkInfoField
	}
	st.NetworkId = pb.NetworkId
	st.NetworkName = pb.NetworkName
	st.SecurityGroupIds = pb.SecurityGroupIds
	st.SubnetIds = pb.SubnetIds
	vpcEndpointsField, err := networkVpcEndpointsFromPb(pb.VpcEndpoints)
	if err != nil {
		return nil, err
	}
	if vpcEndpointsField != nil {
		st.VpcEndpoints = vpcEndpointsField
	}
	st.VpcId = pb.VpcId
	st.VpcStatus = pb.VpcStatus

	var warningMessagesField []NetworkWarning
	for _, itemPb := range pb.WarningMessages {
		item, err := networkWarningFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			warningMessagesField = append(warningMessagesField, *item)
		}
	}
	st.WarningMessages = warningMessagesField
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *networkPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st networkPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NetworkHealth struct {
	// Details of the error.
	// Wire name: 'error_message'
	ErrorMessage string
	// The AWS resource associated with this error: credentials, VPC, subnet,
	// security group, or network ACL.
	// Wire name: 'error_type'
	ErrorType ErrorType

	ForceSendFields []string `tf:"-"`
}

func networkHealthToPb(st *NetworkHealth) (*networkHealthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &networkHealthPb{}
	pb.ErrorMessage = st.ErrorMessage

	pb.ErrorType = st.ErrorType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type networkHealthPb struct {
	// Details of the error.
	ErrorMessage string `json:"error_message,omitempty"`
	// The AWS resource associated with this error: credentials, VPC, subnet,
	// security group, or network ACL.
	ErrorType ErrorType `json:"error_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func networkHealthFromPb(pb *networkHealthPb) (*NetworkHealth, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NetworkHealth{}
	st.ErrorMessage = pb.ErrorMessage
	st.ErrorType = pb.ErrorType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *networkHealthPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st networkHealthPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// If specified, contains the VPC endpoints used to allow cluster communication
// from this VPC over [AWS PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
type NetworkVpcEndpoints struct {
	// The VPC endpoint ID used by this network to access the Databricks secure
	// cluster connectivity relay.
	// Wire name: 'dataplane_relay'
	DataplaneRelay []string
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	// Wire name: 'rest_api'
	RestApi []string
}

func networkVpcEndpointsToPb(st *NetworkVpcEndpoints) (*networkVpcEndpointsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &networkVpcEndpointsPb{}
	pb.DataplaneRelay = st.DataplaneRelay

	pb.RestApi = st.RestApi

	return pb, nil
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

type networkVpcEndpointsPb struct {
	// The VPC endpoint ID used by this network to access the Databricks secure
	// cluster connectivity relay.
	DataplaneRelay []string `json:"dataplane_relay"`
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	RestApi []string `json:"rest_api"`
}

func networkVpcEndpointsFromPb(pb *networkVpcEndpointsPb) (*NetworkVpcEndpoints, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NetworkVpcEndpoints{}
	st.DataplaneRelay = pb.DataplaneRelay
	st.RestApi = pb.RestApi

	return st, nil
}

type NetworkWarning struct {
	// Details of the warning.
	// Wire name: 'warning_message'
	WarningMessage string
	// The AWS resource associated with this warning: a subnet or a security
	// group.
	// Wire name: 'warning_type'
	WarningType WarningType

	ForceSendFields []string `tf:"-"`
}

func networkWarningToPb(st *NetworkWarning) (*networkWarningPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &networkWarningPb{}
	pb.WarningMessage = st.WarningMessage

	pb.WarningType = st.WarningType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type networkWarningPb struct {
	// Details of the warning.
	WarningMessage string `json:"warning_message,omitempty"`
	// The AWS resource associated with this warning: a subnet or a security
	// group.
	WarningType WarningType `json:"warning_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func networkWarningFromPb(pb *networkWarningPb) (*NetworkWarning, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NetworkWarning{}
	st.WarningMessage = pb.WarningMessage
	st.WarningType = pb.WarningType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *networkWarningPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st networkWarningPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The pricing tier of the workspace. For pricing tier information, see [AWS
// Pricing].
//
// [AWS Pricing]: https://databricks.com/product/aws-pricing
type PricingTier string
type pricingTierPb string

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

// Type always returns PricingTier to satisfy [pflag.Value] interface
func (f *PricingTier) Type() string {
	return "PricingTier"
}

func pricingTierToPb(st *PricingTier) (*pricingTierPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pricingTierPb(*st)
	return &pb, nil
}

func pricingTierFromPb(pb *pricingTierPb) (*PricingTier, error) {
	if pb == nil {
		return nil, nil
	}
	st := PricingTier(*pb)
	return &st, nil
}

// The private access level controls which VPC endpoints can connect to the UI
// or API of any workspace that attaches this private access settings object. *
// `ACCOUNT` level access (the default) allows only VPC endpoints that are
// registered in your Databricks account connect to your workspace. * `ENDPOINT`
// level access allows only specified VPC endpoints connect to your workspace.
// For details, see `allowed_vpc_endpoint_ids`.
type PrivateAccessLevel string
type privateAccessLevelPb string

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

// Type always returns PrivateAccessLevel to satisfy [pflag.Value] interface
func (f *PrivateAccessLevel) Type() string {
	return "PrivateAccessLevel"
}

func privateAccessLevelToPb(st *PrivateAccessLevel) (*privateAccessLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := privateAccessLevelPb(*st)
	return &pb, nil
}

func privateAccessLevelFromPb(pb *privateAccessLevelPb) (*PrivateAccessLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := PrivateAccessLevel(*pb)
	return &st, nil
}

type PrivateAccessSettings struct {
	// The Databricks account ID that hosts the credential.
	// Wire name: 'account_id'
	AccountId string
	// An array of Databricks VPC endpoint IDs.
	// Wire name: 'allowed_vpc_endpoint_ids'
	AllowedVpcEndpointIds []string
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// * `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see `allowed_vpc_endpoint_ids`.
	// Wire name: 'private_access_level'
	PrivateAccessLevel PrivateAccessLevel
	// Databricks private access settings ID.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string
	// The human-readable name of the private access settings object.
	// Wire name: 'private_access_settings_name'
	PrivateAccessSettingsName string
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	// Wire name: 'public_access_enabled'
	PublicAccessEnabled bool
	// The cloud region for workspaces attached to this private access settings
	// object.
	// Wire name: 'region'
	Region string

	ForceSendFields []string `tf:"-"`
}

func privateAccessSettingsToPb(st *PrivateAccessSettings) (*privateAccessSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &privateAccessSettingsPb{}
	pb.AccountId = st.AccountId

	pb.AllowedVpcEndpointIds = st.AllowedVpcEndpointIds

	pb.PrivateAccessLevel = st.PrivateAccessLevel

	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	pb.PrivateAccessSettingsName = st.PrivateAccessSettingsName

	pb.PublicAccessEnabled = st.PublicAccessEnabled

	pb.Region = st.Region

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type privateAccessSettingsPb struct {
	// The Databricks account ID that hosts the credential.
	AccountId string `json:"account_id,omitempty"`
	// An array of Databricks VPC endpoint IDs.
	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// * `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see `allowed_vpc_endpoint_ids`.
	PrivateAccessLevel PrivateAccessLevel `json:"private_access_level,omitempty"`
	// Databricks private access settings ID.
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName string `json:"private_access_settings_name,omitempty"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled bool `json:"public_access_enabled,omitempty"`
	// The cloud region for workspaces attached to this private access settings
	// object.
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func privateAccessSettingsFromPb(pb *privateAccessSettingsPb) (*PrivateAccessSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PrivateAccessSettings{}
	st.AccountId = pb.AccountId
	st.AllowedVpcEndpointIds = pb.AllowedVpcEndpointIds
	st.PrivateAccessLevel = pb.PrivateAccessLevel
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId
	st.PrivateAccessSettingsName = pb.PrivateAccessSettingsName
	st.PublicAccessEnabled = pb.PublicAccessEnabled
	st.Region = pb.Region

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *privateAccessSettingsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st privateAccessSettingsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ReplaceResponse struct {
}

func replaceResponseToPb(st *ReplaceResponse) (*replaceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &replaceResponsePb{}

	return pb, nil
}

func (st *ReplaceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &replaceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := replaceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ReplaceResponse) MarshalJSON() ([]byte, error) {
	pb, err := replaceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type replaceResponsePb struct {
}

func replaceResponseFromPb(pb *replaceResponsePb) (*ReplaceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReplaceResponse{}

	return st, nil
}

// Root S3 bucket information.
type RootBucketInfo struct {
	// The name of the S3 bucket.
	// Wire name: 'bucket_name'
	BucketName string

	ForceSendFields []string `tf:"-"`
}

func rootBucketInfoToPb(st *RootBucketInfo) (*rootBucketInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &rootBucketInfoPb{}
	pb.BucketName = st.BucketName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type rootBucketInfoPb struct {
	// The name of the S3 bucket.
	BucketName string `json:"bucket_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func rootBucketInfoFromPb(pb *rootBucketInfoPb) (*RootBucketInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RootBucketInfo{}
	st.BucketName = pb.BucketName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *rootBucketInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st rootBucketInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StorageConfiguration struct {
	// The Databricks account ID that hosts the credential.
	// Wire name: 'account_id'
	AccountId string
	// Time in epoch milliseconds when the storage configuration was created.
	// Wire name: 'creation_time'
	CreationTime int64
	// Root S3 bucket information.
	// Wire name: 'root_bucket_info'
	RootBucketInfo *RootBucketInfo
	// Databricks storage configuration ID.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string
	// The human-readable name of the storage configuration.
	// Wire name: 'storage_configuration_name'
	StorageConfigurationName string

	ForceSendFields []string `tf:"-"`
}

func storageConfigurationToPb(st *StorageConfiguration) (*storageConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &storageConfigurationPb{}
	pb.AccountId = st.AccountId

	pb.CreationTime = st.CreationTime

	rootBucketInfoPb, err := rootBucketInfoToPb(st.RootBucketInfo)
	if err != nil {
		return nil, err
	}
	if rootBucketInfoPb != nil {
		pb.RootBucketInfo = rootBucketInfoPb
	}

	pb.StorageConfigurationId = st.StorageConfigurationId

	pb.StorageConfigurationName = st.StorageConfigurationName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type storageConfigurationPb struct {
	// The Databricks account ID that hosts the credential.
	AccountId string `json:"account_id,omitempty"`
	// Time in epoch milliseconds when the storage configuration was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Root S3 bucket information.
	RootBucketInfo *rootBucketInfoPb `json:"root_bucket_info,omitempty"`
	// Databricks storage configuration ID.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName string `json:"storage_configuration_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func storageConfigurationFromPb(pb *storageConfigurationPb) (*StorageConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StorageConfiguration{}
	st.AccountId = pb.AccountId
	st.CreationTime = pb.CreationTime
	rootBucketInfoField, err := rootBucketInfoFromPb(pb.RootBucketInfo)
	if err != nil {
		return nil, err
	}
	if rootBucketInfoField != nil {
		st.RootBucketInfo = rootBucketInfoField
	}
	st.StorageConfigurationId = pb.StorageConfigurationId
	st.StorageConfigurationName = pb.StorageConfigurationName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *storageConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st storageConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StsRole struct {
	// The external ID that needs to be trusted by the cross-account role. This
	// is always your Databricks account ID.
	// Wire name: 'external_id'
	ExternalId string
	// The Amazon Resource Name (ARN) of the cross account role.
	// Wire name: 'role_arn'
	RoleArn string

	ForceSendFields []string `tf:"-"`
}

func stsRoleToPb(st *StsRole) (*stsRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stsRolePb{}
	pb.ExternalId = st.ExternalId

	pb.RoleArn = st.RoleArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type stsRolePb struct {
	// The external ID that needs to be trusted by the cross-account role. This
	// is always your Databricks account ID.
	ExternalId string `json:"external_id,omitempty"`
	// The Amazon Resource Name (ARN) of the cross account role.
	RoleArn string `json:"role_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func stsRoleFromPb(pb *stsRolePb) (*StsRole, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StsRole{}
	st.ExternalId = pb.ExternalId
	st.RoleArn = pb.RoleArn

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *stsRolePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st stsRolePb) MarshalJSON() ([]byte, error) {
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

type UpdateWorkspaceRequest struct {
	// The AWS region of the workspace's data plane (for example, `us-west-2`).
	// This parameter is available only for updating failed workspaces.
	// Wire name: 'aws_region'
	AwsRegion string
	// ID of the workspace's credential configuration object. This parameter is
	// available for updating both failed and running workspaces.
	// Wire name: 'credentials_id'
	CredentialsId string
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	// Wire name: 'custom_tags'
	CustomTags map[string]string
	// The ID of the workspace's managed services encryption key configuration
	// object. This parameter is available only for updating failed workspaces.
	// Wire name: 'managed_services_customer_managed_key_id'
	ManagedServicesCustomerManagedKeyId string

	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string
	// The ID of the workspace's network configuration object. Used only if you
	// already use a customer-managed VPC. For failed workspaces only, you can
	// switch from a Databricks-managed VPC to a customer-managed VPC by
	// updating the workspace to add a network configuration ID.
	// Wire name: 'network_id'
	NetworkId string
	// The ID of the workspace's private access settings configuration object.
	// This parameter is available only for updating failed workspaces.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string
	// The ID of the workspace's storage configuration object. This parameter is
	// available only for updating failed workspaces.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string
	// The ID of the key configuration object for workspace storage. This
	// parameter is available for updating both failed and running workspaces.
	// Wire name: 'storage_customer_managed_key_id'
	StorageCustomerManagedKeyId string
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func updateWorkspaceRequestToPb(st *UpdateWorkspaceRequest) (*updateWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateWorkspaceRequestPb{}
	pb.AwsRegion = st.AwsRegion

	pb.CredentialsId = st.CredentialsId

	pb.CustomTags = st.CustomTags

	pb.ManagedServicesCustomerManagedKeyId = st.ManagedServicesCustomerManagedKeyId

	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId

	pb.NetworkId = st.NetworkId

	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	pb.StorageConfigurationId = st.StorageConfigurationId

	pb.StorageCustomerManagedKeyId = st.StorageCustomerManagedKeyId

	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type updateWorkspaceRequestPb struct {
	// The AWS region of the workspace's data plane (for example, `us-west-2`).
	// This parameter is available only for updating failed workspaces.
	AwsRegion string `json:"aws_region,omitempty"`
	// ID of the workspace's credential configuration object. This parameter is
	// available for updating both failed and running workspaces.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This parameter is available only for updating failed workspaces.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`

	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// The ID of the workspace's network configuration object. Used only if you
	// already use a customer-managed VPC. For failed workspaces only, you can
	// switch from a Databricks-managed VPC to a customer-managed VPC by
	// updating the workspace to add a network configuration ID.
	NetworkId string `json:"network_id,omitempty"`
	// The ID of the workspace's private access settings configuration object.
	// This parameter is available only for updating failed workspaces.
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// The ID of the workspace's storage configuration object. This parameter is
	// available only for updating failed workspaces.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The ID of the key configuration object for workspace storage. This
	// parameter is available for updating both failed and running workspaces.
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateWorkspaceRequestFromPb(pb *updateWorkspaceRequestPb) (*UpdateWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWorkspaceRequest{}
	st.AwsRegion = pb.AwsRegion
	st.CredentialsId = pb.CredentialsId
	st.CustomTags = pb.CustomTags
	st.ManagedServicesCustomerManagedKeyId = pb.ManagedServicesCustomerManagedKeyId
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	st.NetworkId = pb.NetworkId
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId
	st.StorageConfigurationId = pb.StorageConfigurationId
	st.StorageCustomerManagedKeyId = pb.StorageCustomerManagedKeyId
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateWorkspaceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateWorkspaceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	AllowedVpcEndpointIds []string
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// * `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see `allowed_vpc_endpoint_ids`.
	// Wire name: 'private_access_level'
	PrivateAccessLevel PrivateAccessLevel
	// Databricks Account API private access settings ID.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string `tf:"-"`
	// The human-readable name of the private access settings object.
	// Wire name: 'private_access_settings_name'
	PrivateAccessSettingsName string
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	// Wire name: 'public_access_enabled'
	PublicAccessEnabled bool
	// The cloud region for workspaces associated with this private access
	// settings object.
	// Wire name: 'region'
	Region string

	ForceSendFields []string `tf:"-"`
}

func upsertPrivateAccessSettingsRequestToPb(st *UpsertPrivateAccessSettingsRequest) (*upsertPrivateAccessSettingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &upsertPrivateAccessSettingsRequestPb{}
	pb.AllowedVpcEndpointIds = st.AllowedVpcEndpointIds

	pb.PrivateAccessLevel = st.PrivateAccessLevel

	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	pb.PrivateAccessSettingsName = st.PrivateAccessSettingsName

	pb.PublicAccessEnabled = st.PublicAccessEnabled

	pb.Region = st.Region

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type upsertPrivateAccessSettingsRequestPb struct {
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
	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// * `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see `allowed_vpc_endpoint_ids`.
	PrivateAccessLevel PrivateAccessLevel `json:"private_access_level,omitempty"`
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string `json:"-" url:"-"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName string `json:"private_access_settings_name"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled bool `json:"public_access_enabled,omitempty"`
	// The cloud region for workspaces associated with this private access
	// settings object.
	Region string `json:"region"`

	ForceSendFields []string `json:"-" url:"-"`
}

func upsertPrivateAccessSettingsRequestFromPb(pb *upsertPrivateAccessSettingsRequestPb) (*UpsertPrivateAccessSettingsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpsertPrivateAccessSettingsRequest{}
	st.AllowedVpcEndpointIds = pb.AllowedVpcEndpointIds
	st.PrivateAccessLevel = pb.PrivateAccessLevel
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId
	st.PrivateAccessSettingsName = pb.PrivateAccessSettingsName
	st.PublicAccessEnabled = pb.PublicAccessEnabled
	st.Region = pb.Region

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *upsertPrivateAccessSettingsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st upsertPrivateAccessSettingsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type VpcEndpoint struct {
	// The Databricks account ID that hosts the VPC endpoint configuration.
	// Wire name: 'account_id'
	AccountId string
	// The AWS Account in which the VPC endpoint object exists.
	// Wire name: 'aws_account_id'
	AwsAccountId string
	// The ID of the Databricks [endpoint service] that this VPC endpoint is
	// connected to. For a list of endpoint service IDs for each supported AWS
	// region, see the [Databricks PrivateLink documentation].
	//
	// [Databricks PrivateLink documentation]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	// Wire name: 'aws_endpoint_service_id'
	AwsEndpointServiceId string
	// The ID of the VPC endpoint object in AWS.
	// Wire name: 'aws_vpc_endpoint_id'
	AwsVpcEndpointId string
	// The Google Cloud specific information for this Private Service Connect
	// endpoint.
	// Wire name: 'gcp_vpc_endpoint_info'
	GcpVpcEndpointInfo *GcpVpcEndpointInfo
	// The AWS region in which this VPC endpoint object exists.
	// Wire name: 'region'
	Region string
	// The current state (such as `available` or `rejected`) of the VPC
	// endpoint. Derived from AWS. For the full set of values, see [AWS
	// DescribeVpcEndpoint documentation].
	//
	// [AWS DescribeVpcEndpoint documentation]: https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-vpc-endpoints.html
	// Wire name: 'state'
	State string
	// This enumeration represents the type of Databricks VPC [endpoint service]
	// that was used when creating this VPC endpoint.
	//
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	// Wire name: 'use_case'
	UseCase EndpointUseCase
	// Databricks VPC endpoint ID. This is the Databricks-specific name of the
	// VPC endpoint. Do not confuse this with the `aws_vpc_endpoint_id`, which
	// is the ID within AWS of the VPC endpoint.
	// Wire name: 'vpc_endpoint_id'
	VpcEndpointId string
	// The human-readable name of the storage configuration.
	// Wire name: 'vpc_endpoint_name'
	VpcEndpointName string

	ForceSendFields []string `tf:"-"`
}

func vpcEndpointToPb(st *VpcEndpoint) (*vpcEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vpcEndpointPb{}
	pb.AccountId = st.AccountId

	pb.AwsAccountId = st.AwsAccountId

	pb.AwsEndpointServiceId = st.AwsEndpointServiceId

	pb.AwsVpcEndpointId = st.AwsVpcEndpointId

	gcpVpcEndpointInfoPb, err := gcpVpcEndpointInfoToPb(st.GcpVpcEndpointInfo)
	if err != nil {
		return nil, err
	}
	if gcpVpcEndpointInfoPb != nil {
		pb.GcpVpcEndpointInfo = gcpVpcEndpointInfoPb
	}

	pb.Region = st.Region

	pb.State = st.State

	pb.UseCase = st.UseCase

	pb.VpcEndpointId = st.VpcEndpointId

	pb.VpcEndpointName = st.VpcEndpointName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type vpcEndpointPb struct {
	// The Databricks account ID that hosts the VPC endpoint configuration.
	AccountId string `json:"account_id,omitempty"`
	// The AWS Account in which the VPC endpoint object exists.
	AwsAccountId string `json:"aws_account_id,omitempty"`
	// The ID of the Databricks [endpoint service] that this VPC endpoint is
	// connected to. For a list of endpoint service IDs for each supported AWS
	// region, see the [Databricks PrivateLink documentation].
	//
	// [Databricks PrivateLink documentation]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	AwsEndpointServiceId string `json:"aws_endpoint_service_id,omitempty"`
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId string `json:"aws_vpc_endpoint_id,omitempty"`
	// The Google Cloud specific information for this Private Service Connect
	// endpoint.
	GcpVpcEndpointInfo *gcpVpcEndpointInfoPb `json:"gcp_vpc_endpoint_info,omitempty"`
	// The AWS region in which this VPC endpoint object exists.
	Region string `json:"region,omitempty"`
	// The current state (such as `available` or `rejected`) of the VPC
	// endpoint. Derived from AWS. For the full set of values, see [AWS
	// DescribeVpcEndpoint documentation].
	//
	// [AWS DescribeVpcEndpoint documentation]: https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-vpc-endpoints.html
	State string `json:"state,omitempty"`
	// This enumeration represents the type of Databricks VPC [endpoint service]
	// that was used when creating this VPC endpoint.
	//
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	UseCase EndpointUseCase `json:"use_case,omitempty"`
	// Databricks VPC endpoint ID. This is the Databricks-specific name of the
	// VPC endpoint. Do not confuse this with the `aws_vpc_endpoint_id`, which
	// is the ID within AWS of the VPC endpoint.
	VpcEndpointId string `json:"vpc_endpoint_id,omitempty"`
	// The human-readable name of the storage configuration.
	VpcEndpointName string `json:"vpc_endpoint_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func vpcEndpointFromPb(pb *vpcEndpointPb) (*VpcEndpoint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &VpcEndpoint{}
	st.AccountId = pb.AccountId
	st.AwsAccountId = pb.AwsAccountId
	st.AwsEndpointServiceId = pb.AwsEndpointServiceId
	st.AwsVpcEndpointId = pb.AwsVpcEndpointId
	gcpVpcEndpointInfoField, err := gcpVpcEndpointInfoFromPb(pb.GcpVpcEndpointInfo)
	if err != nil {
		return nil, err
	}
	if gcpVpcEndpointInfoField != nil {
		st.GcpVpcEndpointInfo = gcpVpcEndpointInfoField
	}
	st.Region = pb.Region
	st.State = pb.State
	st.UseCase = pb.UseCase
	st.VpcEndpointId = pb.VpcEndpointId
	st.VpcEndpointName = pb.VpcEndpointName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *vpcEndpointPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st vpcEndpointPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The status of this network configuration object in terms of its use in a
// workspace: * `UNATTACHED`: Unattached. * `VALID`: Valid. * `BROKEN`: Broken.
// * `WARNED`: Warned.
type VpcStatus string
type vpcStatusPb string

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

// Type always returns VpcStatus to satisfy [pflag.Value] interface
func (f *VpcStatus) Type() string {
	return "VpcStatus"
}

func vpcStatusToPb(st *VpcStatus) (*vpcStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := vpcStatusPb(*st)
	return &pb, nil
}

func vpcStatusFromPb(pb *vpcStatusPb) (*VpcStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := VpcStatus(*pb)
	return &st, nil
}

// The AWS resource associated with this warning: a subnet or a security group.
type WarningType string
type warningTypePb string

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

// Type always returns WarningType to satisfy [pflag.Value] interface
func (f *WarningType) Type() string {
	return "WarningType"
}

func warningTypeToPb(st *WarningType) (*warningTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := warningTypePb(*st)
	return &pb, nil
}

func warningTypeFromPb(pb *warningTypePb) (*WarningType, error) {
	if pb == nil {
		return nil, nil
	}
	st := WarningType(*pb)
	return &st, nil
}

type Workspace struct {
	// Databricks account ID.
	// Wire name: 'account_id'
	AccountId string
	// The AWS region of the workspace data plane (for example, `us-west-2`).
	// Wire name: 'aws_region'
	AwsRegion string

	// Wire name: 'azure_workspace_info'
	AzureWorkspaceInfo *AzureWorkspaceInfo
	// The cloud name. This field always has the value `gcp`.
	// Wire name: 'cloud'
	Cloud string
	// The general workspace configurations that are specific to cloud
	// providers.
	// Wire name: 'cloud_resource_container'
	CloudResourceContainer *CloudResourceContainer
	// Time in epoch milliseconds when the workspace was created.
	// Wire name: 'creation_time'
	CreationTime int64
	// ID of the workspace's credential configuration object.
	// Wire name: 'credentials_id'
	CredentialsId string
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	// Wire name: 'custom_tags'
	CustomTags map[string]string
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for web application and REST APIs is
	// `<deployment-name>.cloud.databricks.com`.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	// Wire name: 'deployment_name'
	DeploymentName string
	// If this workspace is for a external customer, then external_customer_info
	// is populated. If this workspace is not for a external customer, then
	// external_customer_info is empty.
	// Wire name: 'external_customer_info'
	ExternalCustomerInfo *ExternalCustomerInfo
	// The network settings for the workspace. The configurations are only for
	// Databricks-managed VPCs. It is ignored if you specify a customer-managed
	// VPC in the `network_id` field.", All the IP range configurations must be
	// mutually exclusive. An attempt to create a workspace fails if Databricks
	// detects an IP range overlap.
	//
	// Specify custom IP ranges in CIDR format. The IP ranges for these fields
	// must not overlap, and all IP addresses must be entirely within the
	// following ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`,
	// `192.168.0.0/16`, and `240.0.0.0/4`.
	//
	// The sizes of these IP ranges affect the maximum number of nodes for the
	// workspace.
	//
	// **Important**: Confirm the IP ranges used by your Databricks workspace
	// before creating the workspace. You cannot change them after your
	// workspace is deployed. If the IP address ranges for your Databricks are
	// too small, IP exhaustion can occur, causing your Databricks jobs to fail.
	// To determine the address range sizes that you need, Databricks provides a
	// calculator as a Microsoft Excel spreadsheet. See [calculate subnet sizes
	// for a new workspace].
	//
	// [calculate subnet sizes for a new workspace]: https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html
	// Wire name: 'gcp_managed_network_config'
	GcpManagedNetworkConfig *GcpManagedNetworkConfig
	// The configurations for the GKE cluster of a Databricks workspace.
	// Wire name: 'gke_config'
	GkeConfig *GkeConfig
	// Whether no public IP is enabled for the workspace.
	// Wire name: 'is_no_public_ip_enabled'
	IsNoPublicIpEnabled bool
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	// Wire name: 'location'
	Location string
	// ID of the key configuration for encrypting managed services.
	// Wire name: 'managed_services_customer_managed_key_id'
	ManagedServicesCustomerManagedKeyId string
	// The network configuration ID that is attached to the workspace. This
	// field is available only if the network is a customer-managed network.
	// Wire name: 'network_id'
	NetworkId string
	// The pricing tier of the workspace. For pricing tier information, see [AWS
	// Pricing].
	//
	// [AWS Pricing]: https://databricks.com/product/aws-pricing
	// Wire name: 'pricing_tier'
	PricingTier PricingTier
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
	PrivateAccessSettingsId string
	// ID of the workspace's storage configuration object.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string
	// ID of the key configuration for encrypting workspace storage.
	// Wire name: 'storage_customer_managed_key_id'
	StorageCustomerManagedKeyId string
	// A unique integer ID for the workspace
	// Wire name: 'workspace_id'
	WorkspaceId int64
	// The human-readable name of the workspace.
	// Wire name: 'workspace_name'
	WorkspaceName string
	// The status of the workspace. For workspace creation, usually it is set to
	// `PROVISIONING` initially. Continue to check the status until the status
	// is `RUNNING`.
	// Wire name: 'workspace_status'
	WorkspaceStatus WorkspaceStatus
	// Message describing the current workspace status.
	// Wire name: 'workspace_status_message'
	WorkspaceStatusMessage string

	ForceSendFields []string `tf:"-"`
}

func workspaceToPb(st *Workspace) (*workspacePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacePb{}
	pb.AccountId = st.AccountId

	pb.AwsRegion = st.AwsRegion

	azureWorkspaceInfoPb, err := azureWorkspaceInfoToPb(st.AzureWorkspaceInfo)
	if err != nil {
		return nil, err
	}
	if azureWorkspaceInfoPb != nil {
		pb.AzureWorkspaceInfo = azureWorkspaceInfoPb
	}

	pb.Cloud = st.Cloud

	cloudResourceContainerPb, err := cloudResourceContainerToPb(st.CloudResourceContainer)
	if err != nil {
		return nil, err
	}
	if cloudResourceContainerPb != nil {
		pb.CloudResourceContainer = cloudResourceContainerPb
	}

	pb.CreationTime = st.CreationTime

	pb.CredentialsId = st.CredentialsId

	pb.CustomTags = st.CustomTags

	pb.DeploymentName = st.DeploymentName

	externalCustomerInfoPb, err := externalCustomerInfoToPb(st.ExternalCustomerInfo)
	if err != nil {
		return nil, err
	}
	if externalCustomerInfoPb != nil {
		pb.ExternalCustomerInfo = externalCustomerInfoPb
	}

	gcpManagedNetworkConfigPb, err := gcpManagedNetworkConfigToPb(st.GcpManagedNetworkConfig)
	if err != nil {
		return nil, err
	}
	if gcpManagedNetworkConfigPb != nil {
		pb.GcpManagedNetworkConfig = gcpManagedNetworkConfigPb
	}

	gkeConfigPb, err := gkeConfigToPb(st.GkeConfig)
	if err != nil {
		return nil, err
	}
	if gkeConfigPb != nil {
		pb.GkeConfig = gkeConfigPb
	}

	pb.IsNoPublicIpEnabled = st.IsNoPublicIpEnabled

	pb.Location = st.Location

	pb.ManagedServicesCustomerManagedKeyId = st.ManagedServicesCustomerManagedKeyId

	pb.NetworkId = st.NetworkId

	pb.PricingTier = st.PricingTier

	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	pb.StorageConfigurationId = st.StorageConfigurationId

	pb.StorageCustomerManagedKeyId = st.StorageCustomerManagedKeyId

	pb.WorkspaceId = st.WorkspaceId

	pb.WorkspaceName = st.WorkspaceName

	pb.WorkspaceStatus = st.WorkspaceStatus

	pb.WorkspaceStatusMessage = st.WorkspaceStatusMessage

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type workspacePb struct {
	// Databricks account ID.
	AccountId string `json:"account_id,omitempty"`
	// The AWS region of the workspace data plane (for example, `us-west-2`).
	AwsRegion string `json:"aws_region,omitempty"`

	AzureWorkspaceInfo *azureWorkspaceInfoPb `json:"azure_workspace_info,omitempty"`
	// The cloud name. This field always has the value `gcp`.
	Cloud string `json:"cloud,omitempty"`
	// The general workspace configurations that are specific to cloud
	// providers.
	CloudResourceContainer *cloudResourceContainerPb `json:"cloud_resource_container,omitempty"`
	// Time in epoch milliseconds when the workspace was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// ID of the workspace's credential configuration object.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for web application and REST APIs is
	// `<deployment-name>.cloud.databricks.com`.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	DeploymentName string `json:"deployment_name,omitempty"`
	// If this workspace is for a external customer, then external_customer_info
	// is populated. If this workspace is not for a external customer, then
	// external_customer_info is empty.
	ExternalCustomerInfo *externalCustomerInfoPb `json:"external_customer_info,omitempty"`
	// The network settings for the workspace. The configurations are only for
	// Databricks-managed VPCs. It is ignored if you specify a customer-managed
	// VPC in the `network_id` field.", All the IP range configurations must be
	// mutually exclusive. An attempt to create a workspace fails if Databricks
	// detects an IP range overlap.
	//
	// Specify custom IP ranges in CIDR format. The IP ranges for these fields
	// must not overlap, and all IP addresses must be entirely within the
	// following ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`,
	// `192.168.0.0/16`, and `240.0.0.0/4`.
	//
	// The sizes of these IP ranges affect the maximum number of nodes for the
	// workspace.
	//
	// **Important**: Confirm the IP ranges used by your Databricks workspace
	// before creating the workspace. You cannot change them after your
	// workspace is deployed. If the IP address ranges for your Databricks are
	// too small, IP exhaustion can occur, causing your Databricks jobs to fail.
	// To determine the address range sizes that you need, Databricks provides a
	// calculator as a Microsoft Excel spreadsheet. See [calculate subnet sizes
	// for a new workspace].
	//
	// [calculate subnet sizes for a new workspace]: https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html
	GcpManagedNetworkConfig *gcpManagedNetworkConfigPb `json:"gcp_managed_network_config,omitempty"`
	// The configurations for the GKE cluster of a Databricks workspace.
	GkeConfig *gkeConfigPb `json:"gke_config,omitempty"`
	// Whether no public IP is enabled for the workspace.
	IsNoPublicIpEnabled bool `json:"is_no_public_ip_enabled,omitempty"`
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	Location string `json:"location,omitempty"`
	// ID of the key configuration for encrypting managed services.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`
	// The network configuration ID that is attached to the workspace. This
	// field is available only if the network is a customer-managed network.
	NetworkId string `json:"network_id,omitempty"`
	// The pricing tier of the workspace. For pricing tier information, see [AWS
	// Pricing].
	//
	// [AWS Pricing]: https://databricks.com/product/aws-pricing
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
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// ID of the workspace's storage configuration object.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// ID of the key configuration for encrypting workspace storage.
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// A unique integer ID for the workspace
	WorkspaceId int64 `json:"workspace_id,omitempty"`
	// The human-readable name of the workspace.
	WorkspaceName string `json:"workspace_name,omitempty"`
	// The status of the workspace. For workspace creation, usually it is set to
	// `PROVISIONING` initially. Continue to check the status until the status
	// is `RUNNING`.
	WorkspaceStatus WorkspaceStatus `json:"workspace_status,omitempty"`
	// Message describing the current workspace status.
	WorkspaceStatusMessage string `json:"workspace_status_message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func workspaceFromPb(pb *workspacePb) (*Workspace, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Workspace{}
	st.AccountId = pb.AccountId
	st.AwsRegion = pb.AwsRegion
	azureWorkspaceInfoField, err := azureWorkspaceInfoFromPb(pb.AzureWorkspaceInfo)
	if err != nil {
		return nil, err
	}
	if azureWorkspaceInfoField != nil {
		st.AzureWorkspaceInfo = azureWorkspaceInfoField
	}
	st.Cloud = pb.Cloud
	cloudResourceContainerField, err := cloudResourceContainerFromPb(pb.CloudResourceContainer)
	if err != nil {
		return nil, err
	}
	if cloudResourceContainerField != nil {
		st.CloudResourceContainer = cloudResourceContainerField
	}
	st.CreationTime = pb.CreationTime
	st.CredentialsId = pb.CredentialsId
	st.CustomTags = pb.CustomTags
	st.DeploymentName = pb.DeploymentName
	externalCustomerInfoField, err := externalCustomerInfoFromPb(pb.ExternalCustomerInfo)
	if err != nil {
		return nil, err
	}
	if externalCustomerInfoField != nil {
		st.ExternalCustomerInfo = externalCustomerInfoField
	}
	gcpManagedNetworkConfigField, err := gcpManagedNetworkConfigFromPb(pb.GcpManagedNetworkConfig)
	if err != nil {
		return nil, err
	}
	if gcpManagedNetworkConfigField != nil {
		st.GcpManagedNetworkConfig = gcpManagedNetworkConfigField
	}
	gkeConfigField, err := gkeConfigFromPb(pb.GkeConfig)
	if err != nil {
		return nil, err
	}
	if gkeConfigField != nil {
		st.GkeConfig = gkeConfigField
	}
	st.IsNoPublicIpEnabled = pb.IsNoPublicIpEnabled
	st.Location = pb.Location
	st.ManagedServicesCustomerManagedKeyId = pb.ManagedServicesCustomerManagedKeyId
	st.NetworkId = pb.NetworkId
	st.PricingTier = pb.PricingTier
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId
	st.StorageConfigurationId = pb.StorageConfigurationId
	st.StorageCustomerManagedKeyId = pb.StorageCustomerManagedKeyId
	st.WorkspaceId = pb.WorkspaceId
	st.WorkspaceName = pb.WorkspaceName
	st.WorkspaceStatus = pb.WorkspaceStatus
	st.WorkspaceStatusMessage = pb.WorkspaceStatusMessage

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *workspacePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st workspacePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The status of the workspace. For workspace creation, usually it is set to
// `PROVISIONING` initially. Continue to check the status until the status is
// `RUNNING`.
type WorkspaceStatus string
type workspaceStatusPb string

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

// Type always returns WorkspaceStatus to satisfy [pflag.Value] interface
func (f *WorkspaceStatus) Type() string {
	return "WorkspaceStatus"
}

func workspaceStatusToPb(st *WorkspaceStatus) (*workspaceStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := workspaceStatusPb(*st)
	return &pb, nil
}

func workspaceStatusFromPb(pb *workspaceStatusPb) (*WorkspaceStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := WorkspaceStatus(*pb)
	return &st, nil
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
