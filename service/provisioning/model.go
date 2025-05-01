// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

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

type AwsCredentials struct {
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
	KeyAlias string
	// The AWS KMS key's Amazon Resource Name (ARN).
	KeyArn string
	// The AWS KMS key region.
	KeyRegion string
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. If you do not want to use this key for encrypting
	// EBS volumes, set to `false`.
	ReuseKeyForClusterVolumes bool

	ForceSendFields []string
}

func awsKeyInfoToPb(st *AwsKeyInfo) (*awsKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsKeyInfoPb{}
	keyAliasPb := &st.KeyAlias
	if keyAliasPb != nil {
		pb.KeyAlias = *keyAliasPb
	}

	keyArnPb := &st.KeyArn
	if keyArnPb != nil {
		pb.KeyArn = *keyArnPb
	}

	keyRegionPb := &st.KeyRegion
	if keyRegionPb != nil {
		pb.KeyRegion = *keyRegionPb
	}

	reuseKeyForClusterVolumesPb := &st.ReuseKeyForClusterVolumes
	if reuseKeyForClusterVolumesPb != nil {
		pb.ReuseKeyForClusterVolumes = *reuseKeyForClusterVolumesPb
	}

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
	keyAliasField := &pb.KeyAlias
	if keyAliasField != nil {
		st.KeyAlias = *keyAliasField
	}
	keyArnField := &pb.KeyArn
	if keyArnField != nil {
		st.KeyArn = *keyArnField
	}
	keyRegionField := &pb.KeyRegion
	if keyRegionField != nil {
		st.KeyRegion = *keyRegionField
	}
	reuseKeyForClusterVolumesField := &pb.ReuseKeyForClusterVolumes
	if reuseKeyForClusterVolumesField != nil {
		st.ReuseKeyForClusterVolumes = *reuseKeyForClusterVolumesField
	}

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
	ResourceGroup string
	// Azure Subscription ID
	SubscriptionId string

	ForceSendFields []string
}

func azureWorkspaceInfoToPb(st *AzureWorkspaceInfo) (*azureWorkspaceInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureWorkspaceInfoPb{}
	resourceGroupPb := &st.ResourceGroup
	if resourceGroupPb != nil {
		pb.ResourceGroup = *resourceGroupPb
	}

	subscriptionIdPb := &st.SubscriptionId
	if subscriptionIdPb != nil {
		pb.SubscriptionId = *subscriptionIdPb
	}

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
	resourceGroupField := &pb.ResourceGroup
	if resourceGroupField != nil {
		st.ResourceGroup = *resourceGroupField
	}
	subscriptionIdField := &pb.SubscriptionId
	if subscriptionIdField != nil {
		st.SubscriptionId = *subscriptionIdField
	}

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
	KeyAlias string
	// The AWS KMS key's Amazon Resource Name (ARN). Note that the key's AWS
	// region is inferred from the ARN.
	KeyArn string
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. To not use this key also for encrypting EBS volumes,
	// set this to `false`.
	ReuseKeyForClusterVolumes bool

	ForceSendFields []string
}

func createAwsKeyInfoToPb(st *CreateAwsKeyInfo) (*createAwsKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAwsKeyInfoPb{}
	keyAliasPb := &st.KeyAlias
	if keyAliasPb != nil {
		pb.KeyAlias = *keyAliasPb
	}

	keyArnPb := &st.KeyArn
	if keyArnPb != nil {
		pb.KeyArn = *keyArnPb
	}

	reuseKeyForClusterVolumesPb := &st.ReuseKeyForClusterVolumes
	if reuseKeyForClusterVolumesPb != nil {
		pb.ReuseKeyForClusterVolumes = *reuseKeyForClusterVolumesPb
	}

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
	keyAliasField := &pb.KeyAlias
	if keyAliasField != nil {
		st.KeyAlias = *keyAliasField
	}
	keyArnField := &pb.KeyArn
	if keyArnField != nil {
		st.KeyArn = *keyArnField
	}
	reuseKeyForClusterVolumesField := &pb.ReuseKeyForClusterVolumes
	if reuseKeyForClusterVolumesField != nil {
		st.ReuseKeyForClusterVolumes = *reuseKeyForClusterVolumesField
	}

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
	AwsCredentials CreateCredentialAwsCredentials
	// The human-readable name of the credential configuration object.
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

	credentialsNamePb := &st.CredentialsName
	if credentialsNamePb != nil {
		pb.CredentialsName = *credentialsNamePb
	}

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
	credentialsNameField := &pb.CredentialsName
	if credentialsNameField != nil {
		st.CredentialsName = *credentialsNameField
	}

	return st, nil
}

type CreateCredentialStsRole struct {
	// The Amazon Resource Name (ARN) of the cross account role.
	RoleArn string

	ForceSendFields []string
}

func createCredentialStsRoleToPb(st *CreateCredentialStsRole) (*createCredentialStsRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCredentialStsRolePb{}
	roleArnPb := &st.RoleArn
	if roleArnPb != nil {
		pb.RoleArn = *roleArnPb
	}

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
	roleArnField := &pb.RoleArn
	if roleArnField != nil {
		st.RoleArn = *roleArnField
	}

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
	AwsKeyInfo *CreateAwsKeyInfo

	GcpKeyInfo *CreateGcpKeyInfo
	// The cases that the key can be used for.
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

	var useCasesPb []KeyUseCase
	for _, item := range st.UseCases {
		itemPb := &item
		if itemPb != nil {
			useCasesPb = append(useCasesPb, *itemPb)
		}
	}
	pb.UseCases = useCasesPb

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

	var useCasesField []KeyUseCase
	for _, item := range pb.UseCases {
		itemField := &item
		if itemField != nil {
			useCasesField = append(useCasesField, *itemField)
		}
	}
	st.UseCases = useCasesField

	return st, nil
}

type CreateGcpKeyInfo struct {
	// The GCP KMS key's resource name
	KmsKeyId string
}

func createGcpKeyInfoToPb(st *CreateGcpKeyInfo) (*createGcpKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createGcpKeyInfoPb{}
	kmsKeyIdPb := &st.KmsKeyId
	if kmsKeyIdPb != nil {
		pb.KmsKeyId = *kmsKeyIdPb
	}

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
	kmsKeyIdField := &pb.KmsKeyId
	if kmsKeyIdField != nil {
		st.KmsKeyId = *kmsKeyIdField
	}

	return st, nil
}

type CreateNetworkRequest struct {
	// The Google Cloud specific information for this network (for example, the
	// VPC ID, subnet ID, and secondary IP ranges).
	GcpNetworkInfo *GcpNetworkInfo
	// The human-readable name of the network configuration.
	NetworkName string
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds []string
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	SubnetIds []string
	// If specified, contains the VPC endpoints used to allow cluster
	// communication from this VPC over [AWS PrivateLink].
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	VpcEndpoints *NetworkVpcEndpoints
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId string

	ForceSendFields []string
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

	networkNamePb := &st.NetworkName
	if networkNamePb != nil {
		pb.NetworkName = *networkNamePb
	}

	var securityGroupIdsPb []string
	for _, item := range st.SecurityGroupIds {
		itemPb := &item
		if itemPb != nil {
			securityGroupIdsPb = append(securityGroupIdsPb, *itemPb)
		}
	}
	pb.SecurityGroupIds = securityGroupIdsPb

	var subnetIdsPb []string
	for _, item := range st.SubnetIds {
		itemPb := &item
		if itemPb != nil {
			subnetIdsPb = append(subnetIdsPb, *itemPb)
		}
	}
	pb.SubnetIds = subnetIdsPb

	vpcEndpointsPb, err := networkVpcEndpointsToPb(st.VpcEndpoints)
	if err != nil {
		return nil, err
	}
	if vpcEndpointsPb != nil {
		pb.VpcEndpoints = vpcEndpointsPb
	}

	vpcIdPb := &st.VpcId
	if vpcIdPb != nil {
		pb.VpcId = *vpcIdPb
	}

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
	networkNameField := &pb.NetworkName
	if networkNameField != nil {
		st.NetworkName = *networkNameField
	}

	var securityGroupIdsField []string
	for _, item := range pb.SecurityGroupIds {
		itemField := &item
		if itemField != nil {
			securityGroupIdsField = append(securityGroupIdsField, *itemField)
		}
	}
	st.SecurityGroupIds = securityGroupIdsField

	var subnetIdsField []string
	for _, item := range pb.SubnetIds {
		itemField := &item
		if itemField != nil {
			subnetIdsField = append(subnetIdsField, *itemField)
		}
	}
	st.SubnetIds = subnetIdsField
	vpcEndpointsField, err := networkVpcEndpointsFromPb(pb.VpcEndpoints)
	if err != nil {
		return nil, err
	}
	if vpcEndpointsField != nil {
		st.VpcEndpoints = vpcEndpointsField
	}
	vpcIdField := &pb.VpcId
	if vpcIdField != nil {
		st.VpcId = *vpcIdField
	}

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
	RootBucketInfo RootBucketInfo
	// The human-readable name of the storage configuration.
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

	storageConfigurationNamePb := &st.StorageConfigurationName
	if storageConfigurationNamePb != nil {
		pb.StorageConfigurationName = *storageConfigurationNamePb
	}

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
	storageConfigurationNameField := &pb.StorageConfigurationName
	if storageConfigurationNameField != nil {
		st.StorageConfigurationName = *storageConfigurationNameField
	}

	return st, nil
}

type CreateVpcEndpointRequest struct {
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId string
	// The Google Cloud specific information for this Private Service Connect
	// endpoint.
	GcpVpcEndpointInfo *GcpVpcEndpointInfo
	// The AWS region in which this VPC endpoint object exists.
	Region string
	// The human-readable name of the storage configuration.
	VpcEndpointName string

	ForceSendFields []string
}

func createVpcEndpointRequestToPb(st *CreateVpcEndpointRequest) (*createVpcEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createVpcEndpointRequestPb{}
	awsVpcEndpointIdPb := &st.AwsVpcEndpointId
	if awsVpcEndpointIdPb != nil {
		pb.AwsVpcEndpointId = *awsVpcEndpointIdPb
	}

	gcpVpcEndpointInfoPb, err := gcpVpcEndpointInfoToPb(st.GcpVpcEndpointInfo)
	if err != nil {
		return nil, err
	}
	if gcpVpcEndpointInfoPb != nil {
		pb.GcpVpcEndpointInfo = gcpVpcEndpointInfoPb
	}

	regionPb := &st.Region
	if regionPb != nil {
		pb.Region = *regionPb
	}

	vpcEndpointNamePb := &st.VpcEndpointName
	if vpcEndpointNamePb != nil {
		pb.VpcEndpointName = *vpcEndpointNamePb
	}

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
	awsVpcEndpointIdField := &pb.AwsVpcEndpointId
	if awsVpcEndpointIdField != nil {
		st.AwsVpcEndpointId = *awsVpcEndpointIdField
	}
	gcpVpcEndpointInfoField, err := gcpVpcEndpointInfoFromPb(pb.GcpVpcEndpointInfo)
	if err != nil {
		return nil, err
	}
	if gcpVpcEndpointInfoField != nil {
		st.GcpVpcEndpointInfo = gcpVpcEndpointInfoField
	}
	regionField := &pb.Region
	if regionField != nil {
		st.Region = *regionField
	}
	vpcEndpointNameField := &pb.VpcEndpointName
	if vpcEndpointNameField != nil {
		st.VpcEndpointName = *vpcEndpointNameField
	}

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
	AwsRegion string
	// The cloud provider which the workspace uses. For Google Cloud workspaces,
	// always set this field to `gcp`.
	Cloud string
	// The general workspace configurations that are specific to cloud
	// providers.
	CloudResourceContainer *CloudResourceContainer
	// ID of the workspace's credential configuration object.
	CredentialsId string
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
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
	GcpManagedNetworkConfig *GcpManagedNetworkConfig
	// The configurations for the GKE cluster of a Databricks workspace.
	GkeConfig *GkeConfig
	// Whether no public IP is enabled for the workspace.
	IsNoPublicIpEnabled bool
	// The Google Cloud region of the workspace data plane in your Google
	// account. For example, `us-east4`.
	Location string
	// The ID of the workspace's managed services encryption key configuration
	// object. This is used to help protect and control access to the
	// workspace's notebooks, secrets, Databricks SQL queries, and query
	// history. The provided key configuration object property `use_cases` must
	// contain `MANAGED_SERVICES`.
	ManagedServicesCustomerManagedKeyId string

	NetworkId string
	// The pricing tier of the workspace. For pricing tier information, see [AWS
	// Pricing].
	//
	// [AWS Pricing]: https://databricks.com/product/aws-pricing
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
	PrivateAccessSettingsId string
	// The ID of the workspace's storage configuration object.
	StorageConfigurationId string
	// The ID of the workspace's storage encryption key configuration object.
	// This is used to encrypt the workspace's root S3 bucket (root DBFS and
	// system data) and, optionally, cluster EBS volumes. The provided key
	// configuration object property `use_cases` must contain `STORAGE`.
	StorageCustomerManagedKeyId string
	// The workspace's human-readable name.
	WorkspaceName string

	ForceSendFields []string
}

func createWorkspaceRequestToPb(st *CreateWorkspaceRequest) (*createWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createWorkspaceRequestPb{}
	awsRegionPb := &st.AwsRegion
	if awsRegionPb != nil {
		pb.AwsRegion = *awsRegionPb
	}

	cloudPb := &st.Cloud
	if cloudPb != nil {
		pb.Cloud = *cloudPb
	}

	cloudResourceContainerPb, err := cloudResourceContainerToPb(st.CloudResourceContainer)
	if err != nil {
		return nil, err
	}
	if cloudResourceContainerPb != nil {
		pb.CloudResourceContainer = cloudResourceContainerPb
	}

	credentialsIdPb := &st.CredentialsId
	if credentialsIdPb != nil {
		pb.CredentialsId = *credentialsIdPb
	}

	customTagsPb := map[string]string{}
	for k, v := range st.CustomTags {
		itemPb := &v
		if itemPb != nil {
			customTagsPb[k] = *itemPb
		}
	}
	pb.CustomTags = customTagsPb

	deploymentNamePb := &st.DeploymentName
	if deploymentNamePb != nil {
		pb.DeploymentName = *deploymentNamePb
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

	isNoPublicIpEnabledPb := &st.IsNoPublicIpEnabled
	if isNoPublicIpEnabledPb != nil {
		pb.IsNoPublicIpEnabled = *isNoPublicIpEnabledPb
	}

	locationPb := &st.Location
	if locationPb != nil {
		pb.Location = *locationPb
	}

	managedServicesCustomerManagedKeyIdPb := &st.ManagedServicesCustomerManagedKeyId
	if managedServicesCustomerManagedKeyIdPb != nil {
		pb.ManagedServicesCustomerManagedKeyId = *managedServicesCustomerManagedKeyIdPb
	}

	networkIdPb := &st.NetworkId
	if networkIdPb != nil {
		pb.NetworkId = *networkIdPb
	}

	pricingTierPb := &st.PricingTier
	if pricingTierPb != nil {
		pb.PricingTier = *pricingTierPb
	}

	privateAccessSettingsIdPb := &st.PrivateAccessSettingsId
	if privateAccessSettingsIdPb != nil {
		pb.PrivateAccessSettingsId = *privateAccessSettingsIdPb
	}

	storageConfigurationIdPb := &st.StorageConfigurationId
	if storageConfigurationIdPb != nil {
		pb.StorageConfigurationId = *storageConfigurationIdPb
	}

	storageCustomerManagedKeyIdPb := &st.StorageCustomerManagedKeyId
	if storageCustomerManagedKeyIdPb != nil {
		pb.StorageCustomerManagedKeyId = *storageCustomerManagedKeyIdPb
	}

	workspaceNamePb := &st.WorkspaceName
	if workspaceNamePb != nil {
		pb.WorkspaceName = *workspaceNamePb
	}

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
	awsRegionField := &pb.AwsRegion
	if awsRegionField != nil {
		st.AwsRegion = *awsRegionField
	}
	cloudField := &pb.Cloud
	if cloudField != nil {
		st.Cloud = *cloudField
	}
	cloudResourceContainerField, err := cloudResourceContainerFromPb(pb.CloudResourceContainer)
	if err != nil {
		return nil, err
	}
	if cloudResourceContainerField != nil {
		st.CloudResourceContainer = cloudResourceContainerField
	}
	credentialsIdField := &pb.CredentialsId
	if credentialsIdField != nil {
		st.CredentialsId = *credentialsIdField
	}

	customTagsField := map[string]string{}
	for k, v := range pb.CustomTags {
		itemField := &v
		if itemField != nil {
			customTagsField[k] = *itemField
		}
	}
	st.CustomTags = customTagsField
	deploymentNameField := &pb.DeploymentName
	if deploymentNameField != nil {
		st.DeploymentName = *deploymentNameField
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
	isNoPublicIpEnabledField := &pb.IsNoPublicIpEnabled
	if isNoPublicIpEnabledField != nil {
		st.IsNoPublicIpEnabled = *isNoPublicIpEnabledField
	}
	locationField := &pb.Location
	if locationField != nil {
		st.Location = *locationField
	}
	managedServicesCustomerManagedKeyIdField := &pb.ManagedServicesCustomerManagedKeyId
	if managedServicesCustomerManagedKeyIdField != nil {
		st.ManagedServicesCustomerManagedKeyId = *managedServicesCustomerManagedKeyIdField
	}
	networkIdField := &pb.NetworkId
	if networkIdField != nil {
		st.NetworkId = *networkIdField
	}
	pricingTierField := &pb.PricingTier
	if pricingTierField != nil {
		st.PricingTier = *pricingTierField
	}
	privateAccessSettingsIdField := &pb.PrivateAccessSettingsId
	if privateAccessSettingsIdField != nil {
		st.PrivateAccessSettingsId = *privateAccessSettingsIdField
	}
	storageConfigurationIdField := &pb.StorageConfigurationId
	if storageConfigurationIdField != nil {
		st.StorageConfigurationId = *storageConfigurationIdField
	}
	storageCustomerManagedKeyIdField := &pb.StorageCustomerManagedKeyId
	if storageCustomerManagedKeyIdField != nil {
		st.StorageCustomerManagedKeyId = *storageCustomerManagedKeyIdField
	}
	workspaceNameField := &pb.WorkspaceName
	if workspaceNameField != nil {
		st.WorkspaceName = *workspaceNameField
	}

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
	AccountId string

	AwsCredentials *AwsCredentials
	// Time in epoch milliseconds when the credential was created.
	CreationTime int64
	// Databricks credential configuration ID.
	CredentialsId string
	// The human-readable name of the credential configuration object.
	CredentialsName string

	ForceSendFields []string
}

func credentialToPb(st *Credential) (*credentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &credentialPb{}
	accountIdPb := &st.AccountId
	if accountIdPb != nil {
		pb.AccountId = *accountIdPb
	}

	awsCredentialsPb, err := awsCredentialsToPb(st.AwsCredentials)
	if err != nil {
		return nil, err
	}
	if awsCredentialsPb != nil {
		pb.AwsCredentials = awsCredentialsPb
	}

	creationTimePb := &st.CreationTime
	if creationTimePb != nil {
		pb.CreationTime = *creationTimePb
	}

	credentialsIdPb := &st.CredentialsId
	if credentialsIdPb != nil {
		pb.CredentialsId = *credentialsIdPb
	}

	credentialsNamePb := &st.CredentialsName
	if credentialsNamePb != nil {
		pb.CredentialsName = *credentialsNamePb
	}

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
	accountIdField := &pb.AccountId
	if accountIdField != nil {
		st.AccountId = *accountIdField
	}
	awsCredentialsField, err := awsCredentialsFromPb(pb.AwsCredentials)
	if err != nil {
		return nil, err
	}
	if awsCredentialsField != nil {
		st.AwsCredentials = awsCredentialsField
	}
	creationTimeField := &pb.CreationTime
	if creationTimeField != nil {
		st.CreationTime = *creationTimeField
	}
	credentialsIdField := &pb.CredentialsId
	if credentialsIdField != nil {
		st.CredentialsId = *credentialsIdField
	}
	credentialsNameField := &pb.CredentialsName
	if credentialsNameField != nil {
		st.CredentialsName = *credentialsNameField
	}

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
	ProjectId string

	ForceSendFields []string
}

func customerFacingGcpCloudResourceContainerToPb(st *CustomerFacingGcpCloudResourceContainer) (*customerFacingGcpCloudResourceContainerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &customerFacingGcpCloudResourceContainerPb{}
	projectIdPb := &st.ProjectId
	if projectIdPb != nil {
		pb.ProjectId = *projectIdPb
	}

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
	projectIdField := &pb.ProjectId
	if projectIdField != nil {
		st.ProjectId = *projectIdField
	}

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
	AccountId string

	AwsKeyInfo *AwsKeyInfo
	// Time in epoch milliseconds when the customer key was created.
	CreationTime int64
	// ID of the encryption key configuration object.
	CustomerManagedKeyId string

	GcpKeyInfo *GcpKeyInfo
	// The cases that the key can be used for.
	UseCases []KeyUseCase

	ForceSendFields []string
}

func customerManagedKeyToPb(st *CustomerManagedKey) (*customerManagedKeyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &customerManagedKeyPb{}
	accountIdPb := &st.AccountId
	if accountIdPb != nil {
		pb.AccountId = *accountIdPb
	}

	awsKeyInfoPb, err := awsKeyInfoToPb(st.AwsKeyInfo)
	if err != nil {
		return nil, err
	}
	if awsKeyInfoPb != nil {
		pb.AwsKeyInfo = awsKeyInfoPb
	}

	creationTimePb := &st.CreationTime
	if creationTimePb != nil {
		pb.CreationTime = *creationTimePb
	}

	customerManagedKeyIdPb := &st.CustomerManagedKeyId
	if customerManagedKeyIdPb != nil {
		pb.CustomerManagedKeyId = *customerManagedKeyIdPb
	}

	gcpKeyInfoPb, err := gcpKeyInfoToPb(st.GcpKeyInfo)
	if err != nil {
		return nil, err
	}
	if gcpKeyInfoPb != nil {
		pb.GcpKeyInfo = gcpKeyInfoPb
	}

	var useCasesPb []KeyUseCase
	for _, item := range st.UseCases {
		itemPb := &item
		if itemPb != nil {
			useCasesPb = append(useCasesPb, *itemPb)
		}
	}
	pb.UseCases = useCasesPb

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
	accountIdField := &pb.AccountId
	if accountIdField != nil {
		st.AccountId = *accountIdField
	}
	awsKeyInfoField, err := awsKeyInfoFromPb(pb.AwsKeyInfo)
	if err != nil {
		return nil, err
	}
	if awsKeyInfoField != nil {
		st.AwsKeyInfo = awsKeyInfoField
	}
	creationTimeField := &pb.CreationTime
	if creationTimeField != nil {
		st.CreationTime = *creationTimeField
	}
	customerManagedKeyIdField := &pb.CustomerManagedKeyId
	if customerManagedKeyIdField != nil {
		st.CustomerManagedKeyId = *customerManagedKeyIdField
	}
	gcpKeyInfoField, err := gcpKeyInfoFromPb(pb.GcpKeyInfo)
	if err != nil {
		return nil, err
	}
	if gcpKeyInfoField != nil {
		st.GcpKeyInfo = gcpKeyInfoField
	}

	var useCasesField []KeyUseCase
	for _, item := range pb.UseCases {
		itemField := &item
		if itemField != nil {
			useCasesField = append(useCasesField, *itemField)
		}
	}
	st.UseCases = useCasesField

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
	CredentialsId string
}

func deleteCredentialRequestToPb(st *DeleteCredentialRequest) (*deleteCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCredentialRequestPb{}
	credentialsIdPb := &st.CredentialsId
	if credentialsIdPb != nil {
		pb.CredentialsId = *credentialsIdPb
	}

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
	credentialsIdField := &pb.CredentialsId
	if credentialsIdField != nil {
		st.CredentialsId = *credentialsIdField
	}

	return st, nil
}

// Delete encryption key configuration
type DeleteEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId string
}

func deleteEncryptionKeyRequestToPb(st *DeleteEncryptionKeyRequest) (*deleteEncryptionKeyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteEncryptionKeyRequestPb{}
	customerManagedKeyIdPb := &st.CustomerManagedKeyId
	if customerManagedKeyIdPb != nil {
		pb.CustomerManagedKeyId = *customerManagedKeyIdPb
	}

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
	customerManagedKeyIdField := &pb.CustomerManagedKeyId
	if customerManagedKeyIdField != nil {
		st.CustomerManagedKeyId = *customerManagedKeyIdField
	}

	return st, nil
}

// Delete a network configuration
type DeleteNetworkRequest struct {
	// Databricks Account API network configuration ID.
	NetworkId string
}

func deleteNetworkRequestToPb(st *DeleteNetworkRequest) (*deleteNetworkRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteNetworkRequestPb{}
	networkIdPb := &st.NetworkId
	if networkIdPb != nil {
		pb.NetworkId = *networkIdPb
	}

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
	networkIdField := &pb.NetworkId
	if networkIdField != nil {
		st.NetworkId = *networkIdField
	}

	return st, nil
}

// Delete a private access settings object
type DeletePrivateAccesRequest struct {
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string
}

func deletePrivateAccesRequestToPb(st *DeletePrivateAccesRequest) (*deletePrivateAccesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePrivateAccesRequestPb{}
	privateAccessSettingsIdPb := &st.PrivateAccessSettingsId
	if privateAccessSettingsIdPb != nil {
		pb.PrivateAccessSettingsId = *privateAccessSettingsIdPb
	}

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
	privateAccessSettingsIdField := &pb.PrivateAccessSettingsId
	if privateAccessSettingsIdField != nil {
		st.PrivateAccessSettingsId = *privateAccessSettingsIdField
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

// Delete storage configuration
type DeleteStorageRequest struct {
	// Databricks Account API storage configuration ID.
	StorageConfigurationId string
}

func deleteStorageRequestToPb(st *DeleteStorageRequest) (*deleteStorageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteStorageRequestPb{}
	storageConfigurationIdPb := &st.StorageConfigurationId
	if storageConfigurationIdPb != nil {
		pb.StorageConfigurationId = *storageConfigurationIdPb
	}

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
	storageConfigurationIdField := &pb.StorageConfigurationId
	if storageConfigurationIdField != nil {
		st.StorageConfigurationId = *storageConfigurationIdField
	}

	return st, nil
}

// Delete VPC endpoint configuration
type DeleteVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId string
}

func deleteVpcEndpointRequestToPb(st *DeleteVpcEndpointRequest) (*deleteVpcEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteVpcEndpointRequestPb{}
	vpcEndpointIdPb := &st.VpcEndpointId
	if vpcEndpointIdPb != nil {
		pb.VpcEndpointId = *vpcEndpointIdPb
	}

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
	vpcEndpointIdField := &pb.VpcEndpointId
	if vpcEndpointIdField != nil {
		st.VpcEndpointId = *vpcEndpointIdField
	}

	return st, nil
}

// Delete a workspace
type DeleteWorkspaceRequest struct {
	// Workspace ID.
	WorkspaceId int64
}

func deleteWorkspaceRequestToPb(st *DeleteWorkspaceRequest) (*deleteWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteWorkspaceRequestPb{}
	workspaceIdPb := &st.WorkspaceId
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

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
	workspaceIdField := &pb.WorkspaceId
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}

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
	AuthoritativeUserEmail string
	// The authoritative user full name.
	AuthoritativeUserFullName string
	// The legal entity name for the external workspace
	CustomerName string

	ForceSendFields []string
}

func externalCustomerInfoToPb(st *ExternalCustomerInfo) (*externalCustomerInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalCustomerInfoPb{}
	authoritativeUserEmailPb := &st.AuthoritativeUserEmail
	if authoritativeUserEmailPb != nil {
		pb.AuthoritativeUserEmail = *authoritativeUserEmailPb
	}

	authoritativeUserFullNamePb := &st.AuthoritativeUserFullName
	if authoritativeUserFullNamePb != nil {
		pb.AuthoritativeUserFullName = *authoritativeUserFullNamePb
	}

	customerNamePb := &st.CustomerName
	if customerNamePb != nil {
		pb.CustomerName = *customerNamePb
	}

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
	authoritativeUserEmailField := &pb.AuthoritativeUserEmail
	if authoritativeUserEmailField != nil {
		st.AuthoritativeUserEmail = *authoritativeUserEmailField
	}
	authoritativeUserFullNameField := &pb.AuthoritativeUserFullName
	if authoritativeUserFullNameField != nil {
		st.AuthoritativeUserFullName = *authoritativeUserFullNameField
	}
	customerNameField := &pb.CustomerName
	if customerNameField != nil {
		st.CustomerName = *customerNameField
	}

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
	KmsKeyId string
}

func gcpKeyInfoToPb(st *GcpKeyInfo) (*gcpKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpKeyInfoPb{}
	kmsKeyIdPb := &st.KmsKeyId
	if kmsKeyIdPb != nil {
		pb.KmsKeyId = *kmsKeyIdPb
	}

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
	kmsKeyIdField := &pb.KmsKeyId
	if kmsKeyIdField != nil {
		st.KmsKeyId = *kmsKeyIdField
	}

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
	GkeClusterPodIpRange string
	// The IP range from which to allocate GKE cluster services. No bigger than
	// `/16` and no smaller than `/27`.
	GkeClusterServiceIpRange string
	// The IP range from which to allocate GKE cluster nodes. No bigger than
	// `/9` and no smaller than `/29`.
	SubnetCidr string

	ForceSendFields []string
}

func gcpManagedNetworkConfigToPb(st *GcpManagedNetworkConfig) (*gcpManagedNetworkConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpManagedNetworkConfigPb{}
	gkeClusterPodIpRangePb := &st.GkeClusterPodIpRange
	if gkeClusterPodIpRangePb != nil {
		pb.GkeClusterPodIpRange = *gkeClusterPodIpRangePb
	}

	gkeClusterServiceIpRangePb := &st.GkeClusterServiceIpRange
	if gkeClusterServiceIpRangePb != nil {
		pb.GkeClusterServiceIpRange = *gkeClusterServiceIpRangePb
	}

	subnetCidrPb := &st.SubnetCidr
	if subnetCidrPb != nil {
		pb.SubnetCidr = *subnetCidrPb
	}

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
	gkeClusterPodIpRangeField := &pb.GkeClusterPodIpRange
	if gkeClusterPodIpRangeField != nil {
		st.GkeClusterPodIpRange = *gkeClusterPodIpRangeField
	}
	gkeClusterServiceIpRangeField := &pb.GkeClusterServiceIpRange
	if gkeClusterServiceIpRangeField != nil {
		st.GkeClusterServiceIpRange = *gkeClusterServiceIpRangeField
	}
	subnetCidrField := &pb.SubnetCidr
	if subnetCidrField != nil {
		st.SubnetCidr = *subnetCidrField
	}

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
	NetworkProjectId string
	// The name of the secondary IP range for pods. A Databricks-managed GKE
	// cluster uses this IP range for its pods. This secondary IP range can be
	// used by only one workspace.
	PodIpRangeName string
	// The name of the secondary IP range for services. A Databricks-managed GKE
	// cluster uses this IP range for its services. This secondary IP range can
	// be used by only one workspace.
	ServiceIpRangeName string
	// The ID of the subnet associated with this network.
	SubnetId string
	// The Google Cloud region of the workspace data plane (for example,
	// `us-east4`).
	SubnetRegion string
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId string
}

func gcpNetworkInfoToPb(st *GcpNetworkInfo) (*gcpNetworkInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpNetworkInfoPb{}
	networkProjectIdPb := &st.NetworkProjectId
	if networkProjectIdPb != nil {
		pb.NetworkProjectId = *networkProjectIdPb
	}

	podIpRangeNamePb := &st.PodIpRangeName
	if podIpRangeNamePb != nil {
		pb.PodIpRangeName = *podIpRangeNamePb
	}

	serviceIpRangeNamePb := &st.ServiceIpRangeName
	if serviceIpRangeNamePb != nil {
		pb.ServiceIpRangeName = *serviceIpRangeNamePb
	}

	subnetIdPb := &st.SubnetId
	if subnetIdPb != nil {
		pb.SubnetId = *subnetIdPb
	}

	subnetRegionPb := &st.SubnetRegion
	if subnetRegionPb != nil {
		pb.SubnetRegion = *subnetRegionPb
	}

	vpcIdPb := &st.VpcId
	if vpcIdPb != nil {
		pb.VpcId = *vpcIdPb
	}

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
	networkProjectIdField := &pb.NetworkProjectId
	if networkProjectIdField != nil {
		st.NetworkProjectId = *networkProjectIdField
	}
	podIpRangeNameField := &pb.PodIpRangeName
	if podIpRangeNameField != nil {
		st.PodIpRangeName = *podIpRangeNameField
	}
	serviceIpRangeNameField := &pb.ServiceIpRangeName
	if serviceIpRangeNameField != nil {
		st.ServiceIpRangeName = *serviceIpRangeNameField
	}
	subnetIdField := &pb.SubnetId
	if subnetIdField != nil {
		st.SubnetId = *subnetIdField
	}
	subnetRegionField := &pb.SubnetRegion
	if subnetRegionField != nil {
		st.SubnetRegion = *subnetRegionField
	}
	vpcIdField := &pb.VpcId
	if vpcIdField != nil {
		st.VpcId = *vpcIdField
	}

	return st, nil
}

// The Google Cloud specific information for this Private Service Connect
// endpoint.
type GcpVpcEndpointInfo struct {
	// Region of the PSC endpoint.
	EndpointRegion string
	// The Google Cloud project ID of the VPC network where the PSC connection
	// resides.
	ProjectId string
	// The unique ID of this PSC connection.
	PscConnectionId string
	// The name of the PSC endpoint in the Google Cloud project.
	PscEndpointName string
	// The service attachment this PSC connection connects to.
	ServiceAttachmentId string

	ForceSendFields []string
}

func gcpVpcEndpointInfoToPb(st *GcpVpcEndpointInfo) (*gcpVpcEndpointInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpVpcEndpointInfoPb{}
	endpointRegionPb := &st.EndpointRegion
	if endpointRegionPb != nil {
		pb.EndpointRegion = *endpointRegionPb
	}

	projectIdPb := &st.ProjectId
	if projectIdPb != nil {
		pb.ProjectId = *projectIdPb
	}

	pscConnectionIdPb := &st.PscConnectionId
	if pscConnectionIdPb != nil {
		pb.PscConnectionId = *pscConnectionIdPb
	}

	pscEndpointNamePb := &st.PscEndpointName
	if pscEndpointNamePb != nil {
		pb.PscEndpointName = *pscEndpointNamePb
	}

	serviceAttachmentIdPb := &st.ServiceAttachmentId
	if serviceAttachmentIdPb != nil {
		pb.ServiceAttachmentId = *serviceAttachmentIdPb
	}

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
	endpointRegionField := &pb.EndpointRegion
	if endpointRegionField != nil {
		st.EndpointRegion = *endpointRegionField
	}
	projectIdField := &pb.ProjectId
	if projectIdField != nil {
		st.ProjectId = *projectIdField
	}
	pscConnectionIdField := &pb.PscConnectionId
	if pscConnectionIdField != nil {
		st.PscConnectionId = *pscConnectionIdField
	}
	pscEndpointNameField := &pb.PscEndpointName
	if pscEndpointNameField != nil {
		st.PscEndpointName = *pscEndpointNameField
	}
	serviceAttachmentIdField := &pb.ServiceAttachmentId
	if serviceAttachmentIdField != nil {
		st.ServiceAttachmentId = *serviceAttachmentIdField
	}

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
	CredentialsId string
}

func getCredentialRequestToPb(st *GetCredentialRequest) (*getCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCredentialRequestPb{}
	credentialsIdPb := &st.CredentialsId
	if credentialsIdPb != nil {
		pb.CredentialsId = *credentialsIdPb
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
	// Databricks Account API credential configuration ID
	CredentialsId string `json:"-" url:"-"`
}

func getCredentialRequestFromPb(pb *getCredentialRequestPb) (*GetCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCredentialRequest{}
	credentialsIdField := &pb.CredentialsId
	if credentialsIdField != nil {
		st.CredentialsId = *credentialsIdField
	}

	return st, nil
}

// Get encryption key configuration
type GetEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId string
}

func getEncryptionKeyRequestToPb(st *GetEncryptionKeyRequest) (*getEncryptionKeyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEncryptionKeyRequestPb{}
	customerManagedKeyIdPb := &st.CustomerManagedKeyId
	if customerManagedKeyIdPb != nil {
		pb.CustomerManagedKeyId = *customerManagedKeyIdPb
	}

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
	customerManagedKeyIdField := &pb.CustomerManagedKeyId
	if customerManagedKeyIdField != nil {
		st.CustomerManagedKeyId = *customerManagedKeyIdField
	}

	return st, nil
}

// Get a network configuration
type GetNetworkRequest struct {
	// Databricks Account API network configuration ID.
	NetworkId string
}

func getNetworkRequestToPb(st *GetNetworkRequest) (*getNetworkRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getNetworkRequestPb{}
	networkIdPb := &st.NetworkId
	if networkIdPb != nil {
		pb.NetworkId = *networkIdPb
	}

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
	networkIdField := &pb.NetworkId
	if networkIdField != nil {
		st.NetworkId = *networkIdField
	}

	return st, nil
}

// Get a private access settings object
type GetPrivateAccesRequest struct {
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string
}

func getPrivateAccesRequestToPb(st *GetPrivateAccesRequest) (*getPrivateAccesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPrivateAccesRequestPb{}
	privateAccessSettingsIdPb := &st.PrivateAccessSettingsId
	if privateAccessSettingsIdPb != nil {
		pb.PrivateAccessSettingsId = *privateAccessSettingsIdPb
	}

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
	privateAccessSettingsIdField := &pb.PrivateAccessSettingsId
	if privateAccessSettingsIdField != nil {
		st.PrivateAccessSettingsId = *privateAccessSettingsIdField
	}

	return st, nil
}

// Get storage configuration
type GetStorageRequest struct {
	// Databricks Account API storage configuration ID.
	StorageConfigurationId string
}

func getStorageRequestToPb(st *GetStorageRequest) (*getStorageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStorageRequestPb{}
	storageConfigurationIdPb := &st.StorageConfigurationId
	if storageConfigurationIdPb != nil {
		pb.StorageConfigurationId = *storageConfigurationIdPb
	}

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
	storageConfigurationIdField := &pb.StorageConfigurationId
	if storageConfigurationIdField != nil {
		st.StorageConfigurationId = *storageConfigurationIdField
	}

	return st, nil
}

// Get a VPC endpoint configuration
type GetVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId string
}

func getVpcEndpointRequestToPb(st *GetVpcEndpointRequest) (*getVpcEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getVpcEndpointRequestPb{}
	vpcEndpointIdPb := &st.VpcEndpointId
	if vpcEndpointIdPb != nil {
		pb.VpcEndpointId = *vpcEndpointIdPb
	}

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
	vpcEndpointIdField := &pb.VpcEndpointId
	if vpcEndpointIdField != nil {
		st.VpcEndpointId = *vpcEndpointIdField
	}

	return st, nil
}

// Get a workspace
type GetWorkspaceRequest struct {
	// Workspace ID.
	WorkspaceId int64
}

func getWorkspaceRequestToPb(st *GetWorkspaceRequest) (*getWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceRequestPb{}
	workspaceIdPb := &st.WorkspaceId
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

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
	workspaceIdField := &pb.WorkspaceId
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}

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
	ConnectivityType GkeConfigConnectivityType
	// The IP range from which to allocate GKE cluster master resources. This
	// field will be ignored if GKE private cluster is not enabled.
	//
	// It must be exactly as big as `/28`.
	MasterIpRange string

	ForceSendFields []string
}

func gkeConfigToPb(st *GkeConfig) (*gkeConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gkeConfigPb{}
	connectivityTypePb := &st.ConnectivityType
	if connectivityTypePb != nil {
		pb.ConnectivityType = *connectivityTypePb
	}

	masterIpRangePb := &st.MasterIpRange
	if masterIpRangePb != nil {
		pb.MasterIpRange = *masterIpRangePb
	}

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
	connectivityTypeField := &pb.ConnectivityType
	if connectivityTypeField != nil {
		st.ConnectivityType = *connectivityTypeField
	}
	masterIpRangeField := &pb.MasterIpRange
	if masterIpRangeField != nil {
		st.MasterIpRange = *masterIpRangeField
	}

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
	AccountId string
	// Time in epoch milliseconds when the network was created.
	CreationTime int64
	// Array of error messages about the network configuration.
	ErrorMessages []NetworkHealth
	// The Google Cloud specific information for this network (for example, the
	// VPC ID, subnet ID, and secondary IP ranges).
	GcpNetworkInfo *GcpNetworkInfo
	// The Databricks network configuration ID.
	NetworkId string
	// The human-readable name of the network configuration.
	NetworkName string

	SecurityGroupIds []string

	SubnetIds []string
	// If specified, contains the VPC endpoints used to allow cluster
	// communication from this VPC over [AWS PrivateLink].
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	VpcEndpoints *NetworkVpcEndpoints
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	VpcId string
	// The status of this network configuration object in terms of its use in a
	// workspace: * `UNATTACHED`: Unattached. * `VALID`: Valid. * `BROKEN`:
	// Broken. * `WARNED`: Warned.
	VpcStatus VpcStatus
	// Array of warning messages about the network configuration.
	WarningMessages []NetworkWarning
	// Workspace ID associated with this network configuration.
	WorkspaceId int64

	ForceSendFields []string
}

func networkToPb(st *Network) (*networkPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &networkPb{}
	accountIdPb := &st.AccountId
	if accountIdPb != nil {
		pb.AccountId = *accountIdPb
	}

	creationTimePb := &st.CreationTime
	if creationTimePb != nil {
		pb.CreationTime = *creationTimePb
	}

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

	networkIdPb := &st.NetworkId
	if networkIdPb != nil {
		pb.NetworkId = *networkIdPb
	}

	networkNamePb := &st.NetworkName
	if networkNamePb != nil {
		pb.NetworkName = *networkNamePb
	}

	var securityGroupIdsPb []string
	for _, item := range st.SecurityGroupIds {
		itemPb := &item
		if itemPb != nil {
			securityGroupIdsPb = append(securityGroupIdsPb, *itemPb)
		}
	}
	pb.SecurityGroupIds = securityGroupIdsPb

	var subnetIdsPb []string
	for _, item := range st.SubnetIds {
		itemPb := &item
		if itemPb != nil {
			subnetIdsPb = append(subnetIdsPb, *itemPb)
		}
	}
	pb.SubnetIds = subnetIdsPb

	vpcEndpointsPb, err := networkVpcEndpointsToPb(st.VpcEndpoints)
	if err != nil {
		return nil, err
	}
	if vpcEndpointsPb != nil {
		pb.VpcEndpoints = vpcEndpointsPb
	}

	vpcIdPb := &st.VpcId
	if vpcIdPb != nil {
		pb.VpcId = *vpcIdPb
	}

	vpcStatusPb := &st.VpcStatus
	if vpcStatusPb != nil {
		pb.VpcStatus = *vpcStatusPb
	}

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

	workspaceIdPb := &st.WorkspaceId
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

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
	accountIdField := &pb.AccountId
	if accountIdField != nil {
		st.AccountId = *accountIdField
	}
	creationTimeField := &pb.CreationTime
	if creationTimeField != nil {
		st.CreationTime = *creationTimeField
	}

	var errorMessagesField []NetworkHealth
	for _, item := range pb.ErrorMessages {
		itemField, err := networkHealthFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			errorMessagesField = append(errorMessagesField, *itemField)
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
	networkIdField := &pb.NetworkId
	if networkIdField != nil {
		st.NetworkId = *networkIdField
	}
	networkNameField := &pb.NetworkName
	if networkNameField != nil {
		st.NetworkName = *networkNameField
	}

	var securityGroupIdsField []string
	for _, item := range pb.SecurityGroupIds {
		itemField := &item
		if itemField != nil {
			securityGroupIdsField = append(securityGroupIdsField, *itemField)
		}
	}
	st.SecurityGroupIds = securityGroupIdsField

	var subnetIdsField []string
	for _, item := range pb.SubnetIds {
		itemField := &item
		if itemField != nil {
			subnetIdsField = append(subnetIdsField, *itemField)
		}
	}
	st.SubnetIds = subnetIdsField
	vpcEndpointsField, err := networkVpcEndpointsFromPb(pb.VpcEndpoints)
	if err != nil {
		return nil, err
	}
	if vpcEndpointsField != nil {
		st.VpcEndpoints = vpcEndpointsField
	}
	vpcIdField := &pb.VpcId
	if vpcIdField != nil {
		st.VpcId = *vpcIdField
	}
	vpcStatusField := &pb.VpcStatus
	if vpcStatusField != nil {
		st.VpcStatus = *vpcStatusField
	}

	var warningMessagesField []NetworkWarning
	for _, item := range pb.WarningMessages {
		itemField, err := networkWarningFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			warningMessagesField = append(warningMessagesField, *itemField)
		}
	}
	st.WarningMessages = warningMessagesField
	workspaceIdField := &pb.WorkspaceId
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}

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
	ErrorMessage string
	// The AWS resource associated with this error: credentials, VPC, subnet,
	// security group, or network ACL.
	ErrorType ErrorType

	ForceSendFields []string
}

func networkHealthToPb(st *NetworkHealth) (*networkHealthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &networkHealthPb{}
	errorMessagePb := &st.ErrorMessage
	if errorMessagePb != nil {
		pb.ErrorMessage = *errorMessagePb
	}

	errorTypePb := &st.ErrorType
	if errorTypePb != nil {
		pb.ErrorType = *errorTypePb
	}

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
	errorMessageField := &pb.ErrorMessage
	if errorMessageField != nil {
		st.ErrorMessage = *errorMessageField
	}
	errorTypeField := &pb.ErrorType
	if errorTypeField != nil {
		st.ErrorType = *errorTypeField
	}

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
	DataplaneRelay []string
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	RestApi []string
}

func networkVpcEndpointsToPb(st *NetworkVpcEndpoints) (*networkVpcEndpointsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &networkVpcEndpointsPb{}

	var dataplaneRelayPb []string
	for _, item := range st.DataplaneRelay {
		itemPb := &item
		if itemPb != nil {
			dataplaneRelayPb = append(dataplaneRelayPb, *itemPb)
		}
	}
	pb.DataplaneRelay = dataplaneRelayPb

	var restApiPb []string
	for _, item := range st.RestApi {
		itemPb := &item
		if itemPb != nil {
			restApiPb = append(restApiPb, *itemPb)
		}
	}
	pb.RestApi = restApiPb

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

	var dataplaneRelayField []string
	for _, item := range pb.DataplaneRelay {
		itemField := &item
		if itemField != nil {
			dataplaneRelayField = append(dataplaneRelayField, *itemField)
		}
	}
	st.DataplaneRelay = dataplaneRelayField

	var restApiField []string
	for _, item := range pb.RestApi {
		itemField := &item
		if itemField != nil {
			restApiField = append(restApiField, *itemField)
		}
	}
	st.RestApi = restApiField

	return st, nil
}

type NetworkWarning struct {
	// Details of the warning.
	WarningMessage string
	// The AWS resource associated with this warning: a subnet or a security
	// group.
	WarningType WarningType

	ForceSendFields []string
}

func networkWarningToPb(st *NetworkWarning) (*networkWarningPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &networkWarningPb{}
	warningMessagePb := &st.WarningMessage
	if warningMessagePb != nil {
		pb.WarningMessage = *warningMessagePb
	}

	warningTypePb := &st.WarningType
	if warningTypePb != nil {
		pb.WarningType = *warningTypePb
	}

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
	warningMessageField := &pb.WarningMessage
	if warningMessageField != nil {
		st.WarningMessage = *warningMessageField
	}
	warningTypeField := &pb.WarningType
	if warningTypeField != nil {
		st.WarningType = *warningTypeField
	}

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
	AccountId string
	// An array of Databricks VPC endpoint IDs.
	AllowedVpcEndpointIds []string
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// * `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see `allowed_vpc_endpoint_ids`.
	PrivateAccessLevel PrivateAccessLevel
	// Databricks private access settings ID.
	PrivateAccessSettingsId string
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName string
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled bool
	// The cloud region for workspaces attached to this private access settings
	// object.
	Region string

	ForceSendFields []string
}

func privateAccessSettingsToPb(st *PrivateAccessSettings) (*privateAccessSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &privateAccessSettingsPb{}
	accountIdPb := &st.AccountId
	if accountIdPb != nil {
		pb.AccountId = *accountIdPb
	}

	var allowedVpcEndpointIdsPb []string
	for _, item := range st.AllowedVpcEndpointIds {
		itemPb := &item
		if itemPb != nil {
			allowedVpcEndpointIdsPb = append(allowedVpcEndpointIdsPb, *itemPb)
		}
	}
	pb.AllowedVpcEndpointIds = allowedVpcEndpointIdsPb

	privateAccessLevelPb := &st.PrivateAccessLevel
	if privateAccessLevelPb != nil {
		pb.PrivateAccessLevel = *privateAccessLevelPb
	}

	privateAccessSettingsIdPb := &st.PrivateAccessSettingsId
	if privateAccessSettingsIdPb != nil {
		pb.PrivateAccessSettingsId = *privateAccessSettingsIdPb
	}

	privateAccessSettingsNamePb := &st.PrivateAccessSettingsName
	if privateAccessSettingsNamePb != nil {
		pb.PrivateAccessSettingsName = *privateAccessSettingsNamePb
	}

	publicAccessEnabledPb := &st.PublicAccessEnabled
	if publicAccessEnabledPb != nil {
		pb.PublicAccessEnabled = *publicAccessEnabledPb
	}

	regionPb := &st.Region
	if regionPb != nil {
		pb.Region = *regionPb
	}

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
	accountIdField := &pb.AccountId
	if accountIdField != nil {
		st.AccountId = *accountIdField
	}

	var allowedVpcEndpointIdsField []string
	for _, item := range pb.AllowedVpcEndpointIds {
		itemField := &item
		if itemField != nil {
			allowedVpcEndpointIdsField = append(allowedVpcEndpointIdsField, *itemField)
		}
	}
	st.AllowedVpcEndpointIds = allowedVpcEndpointIdsField
	privateAccessLevelField := &pb.PrivateAccessLevel
	if privateAccessLevelField != nil {
		st.PrivateAccessLevel = *privateAccessLevelField
	}
	privateAccessSettingsIdField := &pb.PrivateAccessSettingsId
	if privateAccessSettingsIdField != nil {
		st.PrivateAccessSettingsId = *privateAccessSettingsIdField
	}
	privateAccessSettingsNameField := &pb.PrivateAccessSettingsName
	if privateAccessSettingsNameField != nil {
		st.PrivateAccessSettingsName = *privateAccessSettingsNameField
	}
	publicAccessEnabledField := &pb.PublicAccessEnabled
	if publicAccessEnabledField != nil {
		st.PublicAccessEnabled = *publicAccessEnabledField
	}
	regionField := &pb.Region
	if regionField != nil {
		st.Region = *regionField
	}

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
	BucketName string

	ForceSendFields []string
}

func rootBucketInfoToPb(st *RootBucketInfo) (*rootBucketInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &rootBucketInfoPb{}
	bucketNamePb := &st.BucketName
	if bucketNamePb != nil {
		pb.BucketName = *bucketNamePb
	}

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
	bucketNameField := &pb.BucketName
	if bucketNameField != nil {
		st.BucketName = *bucketNameField
	}

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
	AccountId string
	// Time in epoch milliseconds when the storage configuration was created.
	CreationTime int64
	// Root S3 bucket information.
	RootBucketInfo *RootBucketInfo
	// Databricks storage configuration ID.
	StorageConfigurationId string
	// The human-readable name of the storage configuration.
	StorageConfigurationName string

	ForceSendFields []string
}

func storageConfigurationToPb(st *StorageConfiguration) (*storageConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &storageConfigurationPb{}
	accountIdPb := &st.AccountId
	if accountIdPb != nil {
		pb.AccountId = *accountIdPb
	}

	creationTimePb := &st.CreationTime
	if creationTimePb != nil {
		pb.CreationTime = *creationTimePb
	}

	rootBucketInfoPb, err := rootBucketInfoToPb(st.RootBucketInfo)
	if err != nil {
		return nil, err
	}
	if rootBucketInfoPb != nil {
		pb.RootBucketInfo = rootBucketInfoPb
	}

	storageConfigurationIdPb := &st.StorageConfigurationId
	if storageConfigurationIdPb != nil {
		pb.StorageConfigurationId = *storageConfigurationIdPb
	}

	storageConfigurationNamePb := &st.StorageConfigurationName
	if storageConfigurationNamePb != nil {
		pb.StorageConfigurationName = *storageConfigurationNamePb
	}

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
	accountIdField := &pb.AccountId
	if accountIdField != nil {
		st.AccountId = *accountIdField
	}
	creationTimeField := &pb.CreationTime
	if creationTimeField != nil {
		st.CreationTime = *creationTimeField
	}
	rootBucketInfoField, err := rootBucketInfoFromPb(pb.RootBucketInfo)
	if err != nil {
		return nil, err
	}
	if rootBucketInfoField != nil {
		st.RootBucketInfo = rootBucketInfoField
	}
	storageConfigurationIdField := &pb.StorageConfigurationId
	if storageConfigurationIdField != nil {
		st.StorageConfigurationId = *storageConfigurationIdField
	}
	storageConfigurationNameField := &pb.StorageConfigurationName
	if storageConfigurationNameField != nil {
		st.StorageConfigurationName = *storageConfigurationNameField
	}

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
	ExternalId string
	// The Amazon Resource Name (ARN) of the cross account role.
	RoleArn string

	ForceSendFields []string
}

func stsRoleToPb(st *StsRole) (*stsRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stsRolePb{}
	externalIdPb := &st.ExternalId
	if externalIdPb != nil {
		pb.ExternalId = *externalIdPb
	}

	roleArnPb := &st.RoleArn
	if roleArnPb != nil {
		pb.RoleArn = *roleArnPb
	}

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
	externalIdField := &pb.ExternalId
	if externalIdField != nil {
		st.ExternalId = *externalIdField
	}
	roleArnField := &pb.RoleArn
	if roleArnField != nil {
		st.RoleArn = *roleArnField
	}

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
	AwsRegion string
	// ID of the workspace's credential configuration object. This parameter is
	// available for updating both failed and running workspaces.
	CredentialsId string
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags map[string]string
	// The ID of the workspace's managed services encryption key configuration
	// object. This parameter is available only for updating failed workspaces.
	ManagedServicesCustomerManagedKeyId string

	NetworkConnectivityConfigId string
	// The ID of the workspace's network configuration object. Used only if you
	// already use a customer-managed VPC. For failed workspaces only, you can
	// switch from a Databricks-managed VPC to a customer-managed VPC by
	// updating the workspace to add a network configuration ID.
	NetworkId string
	// The ID of the workspace's private access settings configuration object.
	// This parameter is available only for updating failed workspaces.
	PrivateAccessSettingsId string
	// The ID of the workspace's storage configuration object. This parameter is
	// available only for updating failed workspaces.
	StorageConfigurationId string
	// The ID of the key configuration object for workspace storage. This
	// parameter is available for updating both failed and running workspaces.
	StorageCustomerManagedKeyId string
	// Workspace ID.
	WorkspaceId int64

	ForceSendFields []string
}

func updateWorkspaceRequestToPb(st *UpdateWorkspaceRequest) (*updateWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateWorkspaceRequestPb{}
	awsRegionPb := &st.AwsRegion
	if awsRegionPb != nil {
		pb.AwsRegion = *awsRegionPb
	}

	credentialsIdPb := &st.CredentialsId
	if credentialsIdPb != nil {
		pb.CredentialsId = *credentialsIdPb
	}

	customTagsPb := map[string]string{}
	for k, v := range st.CustomTags {
		itemPb := &v
		if itemPb != nil {
			customTagsPb[k] = *itemPb
		}
	}
	pb.CustomTags = customTagsPb

	managedServicesCustomerManagedKeyIdPb := &st.ManagedServicesCustomerManagedKeyId
	if managedServicesCustomerManagedKeyIdPb != nil {
		pb.ManagedServicesCustomerManagedKeyId = *managedServicesCustomerManagedKeyIdPb
	}

	networkConnectivityConfigIdPb := &st.NetworkConnectivityConfigId
	if networkConnectivityConfigIdPb != nil {
		pb.NetworkConnectivityConfigId = *networkConnectivityConfigIdPb
	}

	networkIdPb := &st.NetworkId
	if networkIdPb != nil {
		pb.NetworkId = *networkIdPb
	}

	privateAccessSettingsIdPb := &st.PrivateAccessSettingsId
	if privateAccessSettingsIdPb != nil {
		pb.PrivateAccessSettingsId = *privateAccessSettingsIdPb
	}

	storageConfigurationIdPb := &st.StorageConfigurationId
	if storageConfigurationIdPb != nil {
		pb.StorageConfigurationId = *storageConfigurationIdPb
	}

	storageCustomerManagedKeyIdPb := &st.StorageCustomerManagedKeyId
	if storageCustomerManagedKeyIdPb != nil {
		pb.StorageCustomerManagedKeyId = *storageCustomerManagedKeyIdPb
	}

	workspaceIdPb := &st.WorkspaceId
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

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
	awsRegionField := &pb.AwsRegion
	if awsRegionField != nil {
		st.AwsRegion = *awsRegionField
	}
	credentialsIdField := &pb.CredentialsId
	if credentialsIdField != nil {
		st.CredentialsId = *credentialsIdField
	}

	customTagsField := map[string]string{}
	for k, v := range pb.CustomTags {
		itemField := &v
		if itemField != nil {
			customTagsField[k] = *itemField
		}
	}
	st.CustomTags = customTagsField
	managedServicesCustomerManagedKeyIdField := &pb.ManagedServicesCustomerManagedKeyId
	if managedServicesCustomerManagedKeyIdField != nil {
		st.ManagedServicesCustomerManagedKeyId = *managedServicesCustomerManagedKeyIdField
	}
	networkConnectivityConfigIdField := &pb.NetworkConnectivityConfigId
	if networkConnectivityConfigIdField != nil {
		st.NetworkConnectivityConfigId = *networkConnectivityConfigIdField
	}
	networkIdField := &pb.NetworkId
	if networkIdField != nil {
		st.NetworkId = *networkIdField
	}
	privateAccessSettingsIdField := &pb.PrivateAccessSettingsId
	if privateAccessSettingsIdField != nil {
		st.PrivateAccessSettingsId = *privateAccessSettingsIdField
	}
	storageConfigurationIdField := &pb.StorageConfigurationId
	if storageConfigurationIdField != nil {
		st.StorageConfigurationId = *storageConfigurationIdField
	}
	storageCustomerManagedKeyIdField := &pb.StorageCustomerManagedKeyId
	if storageCustomerManagedKeyIdField != nil {
		st.StorageCustomerManagedKeyId = *storageCustomerManagedKeyIdField
	}
	workspaceIdField := &pb.WorkspaceId
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}

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
	AllowedVpcEndpointIds []string
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// * `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see `allowed_vpc_endpoint_ids`.
	PrivateAccessLevel PrivateAccessLevel
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName string
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled bool
	// The cloud region for workspaces associated with this private access
	// settings object.
	Region string

	ForceSendFields []string
}

func upsertPrivateAccessSettingsRequestToPb(st *UpsertPrivateAccessSettingsRequest) (*upsertPrivateAccessSettingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &upsertPrivateAccessSettingsRequestPb{}

	var allowedVpcEndpointIdsPb []string
	for _, item := range st.AllowedVpcEndpointIds {
		itemPb := &item
		if itemPb != nil {
			allowedVpcEndpointIdsPb = append(allowedVpcEndpointIdsPb, *itemPb)
		}
	}
	pb.AllowedVpcEndpointIds = allowedVpcEndpointIdsPb

	privateAccessLevelPb := &st.PrivateAccessLevel
	if privateAccessLevelPb != nil {
		pb.PrivateAccessLevel = *privateAccessLevelPb
	}

	privateAccessSettingsIdPb := &st.PrivateAccessSettingsId
	if privateAccessSettingsIdPb != nil {
		pb.PrivateAccessSettingsId = *privateAccessSettingsIdPb
	}

	privateAccessSettingsNamePb := &st.PrivateAccessSettingsName
	if privateAccessSettingsNamePb != nil {
		pb.PrivateAccessSettingsName = *privateAccessSettingsNamePb
	}

	publicAccessEnabledPb := &st.PublicAccessEnabled
	if publicAccessEnabledPb != nil {
		pb.PublicAccessEnabled = *publicAccessEnabledPb
	}

	regionPb := &st.Region
	if regionPb != nil {
		pb.Region = *regionPb
	}

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

	var allowedVpcEndpointIdsField []string
	for _, item := range pb.AllowedVpcEndpointIds {
		itemField := &item
		if itemField != nil {
			allowedVpcEndpointIdsField = append(allowedVpcEndpointIdsField, *itemField)
		}
	}
	st.AllowedVpcEndpointIds = allowedVpcEndpointIdsField
	privateAccessLevelField := &pb.PrivateAccessLevel
	if privateAccessLevelField != nil {
		st.PrivateAccessLevel = *privateAccessLevelField
	}
	privateAccessSettingsIdField := &pb.PrivateAccessSettingsId
	if privateAccessSettingsIdField != nil {
		st.PrivateAccessSettingsId = *privateAccessSettingsIdField
	}
	privateAccessSettingsNameField := &pb.PrivateAccessSettingsName
	if privateAccessSettingsNameField != nil {
		st.PrivateAccessSettingsName = *privateAccessSettingsNameField
	}
	publicAccessEnabledField := &pb.PublicAccessEnabled
	if publicAccessEnabledField != nil {
		st.PublicAccessEnabled = *publicAccessEnabledField
	}
	regionField := &pb.Region
	if regionField != nil {
		st.Region = *regionField
	}

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
	AccountId string
	// The AWS Account in which the VPC endpoint object exists.
	AwsAccountId string
	// The ID of the Databricks [endpoint service] that this VPC endpoint is
	// connected to. For a list of endpoint service IDs for each supported AWS
	// region, see the [Databricks PrivateLink documentation].
	//
	// [Databricks PrivateLink documentation]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	AwsEndpointServiceId string
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId string
	// The Google Cloud specific information for this Private Service Connect
	// endpoint.
	GcpVpcEndpointInfo *GcpVpcEndpointInfo
	// The AWS region in which this VPC endpoint object exists.
	Region string
	// The current state (such as `available` or `rejected`) of the VPC
	// endpoint. Derived from AWS. For the full set of values, see [AWS
	// DescribeVpcEndpoint documentation].
	//
	// [AWS DescribeVpcEndpoint documentation]: https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-vpc-endpoints.html
	State string
	// This enumeration represents the type of Databricks VPC [endpoint service]
	// that was used when creating this VPC endpoint.
	//
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	UseCase EndpointUseCase
	// Databricks VPC endpoint ID. This is the Databricks-specific name of the
	// VPC endpoint. Do not confuse this with the `aws_vpc_endpoint_id`, which
	// is the ID within AWS of the VPC endpoint.
	VpcEndpointId string
	// The human-readable name of the storage configuration.
	VpcEndpointName string

	ForceSendFields []string
}

func vpcEndpointToPb(st *VpcEndpoint) (*vpcEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vpcEndpointPb{}
	accountIdPb := &st.AccountId
	if accountIdPb != nil {
		pb.AccountId = *accountIdPb
	}

	awsAccountIdPb := &st.AwsAccountId
	if awsAccountIdPb != nil {
		pb.AwsAccountId = *awsAccountIdPb
	}

	awsEndpointServiceIdPb := &st.AwsEndpointServiceId
	if awsEndpointServiceIdPb != nil {
		pb.AwsEndpointServiceId = *awsEndpointServiceIdPb
	}

	awsVpcEndpointIdPb := &st.AwsVpcEndpointId
	if awsVpcEndpointIdPb != nil {
		pb.AwsVpcEndpointId = *awsVpcEndpointIdPb
	}

	gcpVpcEndpointInfoPb, err := gcpVpcEndpointInfoToPb(st.GcpVpcEndpointInfo)
	if err != nil {
		return nil, err
	}
	if gcpVpcEndpointInfoPb != nil {
		pb.GcpVpcEndpointInfo = gcpVpcEndpointInfoPb
	}

	regionPb := &st.Region
	if regionPb != nil {
		pb.Region = *regionPb
	}

	statePb := &st.State
	if statePb != nil {
		pb.State = *statePb
	}

	useCasePb := &st.UseCase
	if useCasePb != nil {
		pb.UseCase = *useCasePb
	}

	vpcEndpointIdPb := &st.VpcEndpointId
	if vpcEndpointIdPb != nil {
		pb.VpcEndpointId = *vpcEndpointIdPb
	}

	vpcEndpointNamePb := &st.VpcEndpointName
	if vpcEndpointNamePb != nil {
		pb.VpcEndpointName = *vpcEndpointNamePb
	}

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
	accountIdField := &pb.AccountId
	if accountIdField != nil {
		st.AccountId = *accountIdField
	}
	awsAccountIdField := &pb.AwsAccountId
	if awsAccountIdField != nil {
		st.AwsAccountId = *awsAccountIdField
	}
	awsEndpointServiceIdField := &pb.AwsEndpointServiceId
	if awsEndpointServiceIdField != nil {
		st.AwsEndpointServiceId = *awsEndpointServiceIdField
	}
	awsVpcEndpointIdField := &pb.AwsVpcEndpointId
	if awsVpcEndpointIdField != nil {
		st.AwsVpcEndpointId = *awsVpcEndpointIdField
	}
	gcpVpcEndpointInfoField, err := gcpVpcEndpointInfoFromPb(pb.GcpVpcEndpointInfo)
	if err != nil {
		return nil, err
	}
	if gcpVpcEndpointInfoField != nil {
		st.GcpVpcEndpointInfo = gcpVpcEndpointInfoField
	}
	regionField := &pb.Region
	if regionField != nil {
		st.Region = *regionField
	}
	stateField := &pb.State
	if stateField != nil {
		st.State = *stateField
	}
	useCaseField := &pb.UseCase
	if useCaseField != nil {
		st.UseCase = *useCaseField
	}
	vpcEndpointIdField := &pb.VpcEndpointId
	if vpcEndpointIdField != nil {
		st.VpcEndpointId = *vpcEndpointIdField
	}
	vpcEndpointNameField := &pb.VpcEndpointName
	if vpcEndpointNameField != nil {
		st.VpcEndpointName = *vpcEndpointNameField
	}

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
	AccountId string
	// The AWS region of the workspace data plane (for example, `us-west-2`).
	AwsRegion string

	AzureWorkspaceInfo *AzureWorkspaceInfo
	// The cloud name. This field always has the value `gcp`.
	Cloud string
	// The general workspace configurations that are specific to cloud
	// providers.
	CloudResourceContainer *CloudResourceContainer
	// Time in epoch milliseconds when the workspace was created.
	CreationTime int64
	// ID of the workspace's credential configuration object.
	CredentialsId string
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags map[string]string
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for web application and REST APIs is
	// `<deployment-name>.cloud.databricks.com`.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	DeploymentName string
	// If this workspace is for a external customer, then external_customer_info
	// is populated. If this workspace is not for a external customer, then
	// external_customer_info is empty.
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
	GcpManagedNetworkConfig *GcpManagedNetworkConfig
	// The configurations for the GKE cluster of a Databricks workspace.
	GkeConfig *GkeConfig
	// Whether no public IP is enabled for the workspace.
	IsNoPublicIpEnabled bool
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	Location string
	// ID of the key configuration for encrypting managed services.
	ManagedServicesCustomerManagedKeyId string
	// The network configuration ID that is attached to the workspace. This
	// field is available only if the network is a customer-managed network.
	NetworkId string
	// The pricing tier of the workspace. For pricing tier information, see [AWS
	// Pricing].
	//
	// [AWS Pricing]: https://databricks.com/product/aws-pricing
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
	PrivateAccessSettingsId string
	// ID of the workspace's storage configuration object.
	StorageConfigurationId string
	// ID of the key configuration for encrypting workspace storage.
	StorageCustomerManagedKeyId string
	// A unique integer ID for the workspace
	WorkspaceId int64
	// The human-readable name of the workspace.
	WorkspaceName string
	// The status of the workspace. For workspace creation, usually it is set to
	// `PROVISIONING` initially. Continue to check the status until the status
	// is `RUNNING`.
	WorkspaceStatus WorkspaceStatus
	// Message describing the current workspace status.
	WorkspaceStatusMessage string

	ForceSendFields []string
}

func workspaceToPb(st *Workspace) (*workspacePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacePb{}
	accountIdPb := &st.AccountId
	if accountIdPb != nil {
		pb.AccountId = *accountIdPb
	}

	awsRegionPb := &st.AwsRegion
	if awsRegionPb != nil {
		pb.AwsRegion = *awsRegionPb
	}

	azureWorkspaceInfoPb, err := azureWorkspaceInfoToPb(st.AzureWorkspaceInfo)
	if err != nil {
		return nil, err
	}
	if azureWorkspaceInfoPb != nil {
		pb.AzureWorkspaceInfo = azureWorkspaceInfoPb
	}

	cloudPb := &st.Cloud
	if cloudPb != nil {
		pb.Cloud = *cloudPb
	}

	cloudResourceContainerPb, err := cloudResourceContainerToPb(st.CloudResourceContainer)
	if err != nil {
		return nil, err
	}
	if cloudResourceContainerPb != nil {
		pb.CloudResourceContainer = cloudResourceContainerPb
	}

	creationTimePb := &st.CreationTime
	if creationTimePb != nil {
		pb.CreationTime = *creationTimePb
	}

	credentialsIdPb := &st.CredentialsId
	if credentialsIdPb != nil {
		pb.CredentialsId = *credentialsIdPb
	}

	customTagsPb := map[string]string{}
	for k, v := range st.CustomTags {
		itemPb := &v
		if itemPb != nil {
			customTagsPb[k] = *itemPb
		}
	}
	pb.CustomTags = customTagsPb

	deploymentNamePb := &st.DeploymentName
	if deploymentNamePb != nil {
		pb.DeploymentName = *deploymentNamePb
	}

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

	isNoPublicIpEnabledPb := &st.IsNoPublicIpEnabled
	if isNoPublicIpEnabledPb != nil {
		pb.IsNoPublicIpEnabled = *isNoPublicIpEnabledPb
	}

	locationPb := &st.Location
	if locationPb != nil {
		pb.Location = *locationPb
	}

	managedServicesCustomerManagedKeyIdPb := &st.ManagedServicesCustomerManagedKeyId
	if managedServicesCustomerManagedKeyIdPb != nil {
		pb.ManagedServicesCustomerManagedKeyId = *managedServicesCustomerManagedKeyIdPb
	}

	networkIdPb := &st.NetworkId
	if networkIdPb != nil {
		pb.NetworkId = *networkIdPb
	}

	pricingTierPb := &st.PricingTier
	if pricingTierPb != nil {
		pb.PricingTier = *pricingTierPb
	}

	privateAccessSettingsIdPb := &st.PrivateAccessSettingsId
	if privateAccessSettingsIdPb != nil {
		pb.PrivateAccessSettingsId = *privateAccessSettingsIdPb
	}

	storageConfigurationIdPb := &st.StorageConfigurationId
	if storageConfigurationIdPb != nil {
		pb.StorageConfigurationId = *storageConfigurationIdPb
	}

	storageCustomerManagedKeyIdPb := &st.StorageCustomerManagedKeyId
	if storageCustomerManagedKeyIdPb != nil {
		pb.StorageCustomerManagedKeyId = *storageCustomerManagedKeyIdPb
	}

	workspaceIdPb := &st.WorkspaceId
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

	workspaceNamePb := &st.WorkspaceName
	if workspaceNamePb != nil {
		pb.WorkspaceName = *workspaceNamePb
	}

	workspaceStatusPb := &st.WorkspaceStatus
	if workspaceStatusPb != nil {
		pb.WorkspaceStatus = *workspaceStatusPb
	}

	workspaceStatusMessagePb := &st.WorkspaceStatusMessage
	if workspaceStatusMessagePb != nil {
		pb.WorkspaceStatusMessage = *workspaceStatusMessagePb
	}

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
	accountIdField := &pb.AccountId
	if accountIdField != nil {
		st.AccountId = *accountIdField
	}
	awsRegionField := &pb.AwsRegion
	if awsRegionField != nil {
		st.AwsRegion = *awsRegionField
	}
	azureWorkspaceInfoField, err := azureWorkspaceInfoFromPb(pb.AzureWorkspaceInfo)
	if err != nil {
		return nil, err
	}
	if azureWorkspaceInfoField != nil {
		st.AzureWorkspaceInfo = azureWorkspaceInfoField
	}
	cloudField := &pb.Cloud
	if cloudField != nil {
		st.Cloud = *cloudField
	}
	cloudResourceContainerField, err := cloudResourceContainerFromPb(pb.CloudResourceContainer)
	if err != nil {
		return nil, err
	}
	if cloudResourceContainerField != nil {
		st.CloudResourceContainer = cloudResourceContainerField
	}
	creationTimeField := &pb.CreationTime
	if creationTimeField != nil {
		st.CreationTime = *creationTimeField
	}
	credentialsIdField := &pb.CredentialsId
	if credentialsIdField != nil {
		st.CredentialsId = *credentialsIdField
	}

	customTagsField := map[string]string{}
	for k, v := range pb.CustomTags {
		itemField := &v
		if itemField != nil {
			customTagsField[k] = *itemField
		}
	}
	st.CustomTags = customTagsField
	deploymentNameField := &pb.DeploymentName
	if deploymentNameField != nil {
		st.DeploymentName = *deploymentNameField
	}
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
	isNoPublicIpEnabledField := &pb.IsNoPublicIpEnabled
	if isNoPublicIpEnabledField != nil {
		st.IsNoPublicIpEnabled = *isNoPublicIpEnabledField
	}
	locationField := &pb.Location
	if locationField != nil {
		st.Location = *locationField
	}
	managedServicesCustomerManagedKeyIdField := &pb.ManagedServicesCustomerManagedKeyId
	if managedServicesCustomerManagedKeyIdField != nil {
		st.ManagedServicesCustomerManagedKeyId = *managedServicesCustomerManagedKeyIdField
	}
	networkIdField := &pb.NetworkId
	if networkIdField != nil {
		st.NetworkId = *networkIdField
	}
	pricingTierField := &pb.PricingTier
	if pricingTierField != nil {
		st.PricingTier = *pricingTierField
	}
	privateAccessSettingsIdField := &pb.PrivateAccessSettingsId
	if privateAccessSettingsIdField != nil {
		st.PrivateAccessSettingsId = *privateAccessSettingsIdField
	}
	storageConfigurationIdField := &pb.StorageConfigurationId
	if storageConfigurationIdField != nil {
		st.StorageConfigurationId = *storageConfigurationIdField
	}
	storageCustomerManagedKeyIdField := &pb.StorageCustomerManagedKeyId
	if storageCustomerManagedKeyIdField != nil {
		st.StorageCustomerManagedKeyId = *storageCustomerManagedKeyIdField
	}
	workspaceIdField := &pb.WorkspaceId
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}
	workspaceNameField := &pb.WorkspaceName
	if workspaceNameField != nil {
		st.WorkspaceName = *workspaceNameField
	}
	workspaceStatusField := &pb.WorkspaceStatus
	if workspaceStatusField != nil {
		st.WorkspaceStatus = *workspaceStatusField
	}
	workspaceStatusMessageField := &pb.WorkspaceStatusMessage
	if workspaceStatusMessageField != nil {
		st.WorkspaceStatusMessage = *workspaceStatusMessageField
	}

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
