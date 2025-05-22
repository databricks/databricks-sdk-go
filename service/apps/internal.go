// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

func appToPb(st *App) (*appPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appPb{}
	pb.ActiveDeployment = st.ActiveDeployment

	pb.AppStatus = st.AppStatus

	pb.BudgetPolicyId = st.BudgetPolicyId

	pb.ComputeStatus = st.ComputeStatus

	pb.CreateTime = st.CreateTime

	pb.Creator = st.Creator

	pb.DefaultSourceCodePath = st.DefaultSourceCodePath

	pb.Description = st.Description

	pb.EffectiveBudgetPolicyId = st.EffectiveBudgetPolicyId

	pb.EffectiveUserApiScopes = st.EffectiveUserApiScopes

	pb.Id = st.Id

	pb.Name = st.Name

	pb.Oauth2AppClientId = st.Oauth2AppClientId

	pb.Oauth2AppIntegrationId = st.Oauth2AppIntegrationId

	pb.PendingDeployment = st.PendingDeployment

	pb.Resources = st.Resources

	pb.ServicePrincipalClientId = st.ServicePrincipalClientId

	pb.ServicePrincipalId = st.ServicePrincipalId

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UpdateTime = st.UpdateTime

	pb.Updater = st.Updater

	pb.Url = st.Url

	pb.UserApiScopes = st.UserApiScopes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type appPb struct {
	ActiveDeployment *AppDeployment `json:"active_deployment,omitempty"`

	AppStatus *ApplicationStatus `json:"app_status,omitempty"`

	BudgetPolicyId string `json:"budget_policy_id,omitempty"`

	ComputeStatus *ComputeStatus `json:"compute_status,omitempty"`

	CreateTime string `json:"create_time,omitempty"`

	Creator string `json:"creator,omitempty"`

	DefaultSourceCodePath string `json:"default_source_code_path,omitempty"`

	Description string `json:"description,omitempty"`

	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`

	EffectiveUserApiScopes []string `json:"effective_user_api_scopes,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name"`

	Oauth2AppClientId string `json:"oauth2_app_client_id,omitempty"`

	Oauth2AppIntegrationId string `json:"oauth2_app_integration_id,omitempty"`

	PendingDeployment *AppDeployment `json:"pending_deployment,omitempty"`

	Resources []AppResource `json:"resources,omitempty"`

	ServicePrincipalClientId string `json:"service_principal_client_id,omitempty"`

	ServicePrincipalId int64 `json:"service_principal_id,omitempty"`

	ServicePrincipalName string `json:"service_principal_name,omitempty"`

	UpdateTime string `json:"update_time,omitempty"`

	Updater string `json:"updater,omitempty"`

	Url string `json:"url,omitempty"`

	UserApiScopes []string `json:"user_api_scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appFromPb(pb *appPb) (*App, error) {
	if pb == nil {
		return nil, nil
	}
	st := &App{}
	st.ActiveDeployment = pb.ActiveDeployment
	st.AppStatus = pb.AppStatus
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.ComputeStatus = pb.ComputeStatus
	st.CreateTime = pb.CreateTime
	st.Creator = pb.Creator
	st.DefaultSourceCodePath = pb.DefaultSourceCodePath
	st.Description = pb.Description
	st.EffectiveBudgetPolicyId = pb.EffectiveBudgetPolicyId
	st.EffectiveUserApiScopes = pb.EffectiveUserApiScopes
	st.Id = pb.Id
	st.Name = pb.Name
	st.Oauth2AppClientId = pb.Oauth2AppClientId
	st.Oauth2AppIntegrationId = pb.Oauth2AppIntegrationId
	st.PendingDeployment = pb.PendingDeployment
	st.Resources = pb.Resources
	st.ServicePrincipalClientId = pb.ServicePrincipalClientId
	st.ServicePrincipalId = pb.ServicePrincipalId
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UpdateTime = pb.UpdateTime
	st.Updater = pb.Updater
	st.Url = pb.Url
	st.UserApiScopes = pb.UserApiScopes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func appAccessControlRequestToPb(st *AppAccessControlRequest) (*appAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appAccessControlRequestPb{}
	pb.GroupName = st.GroupName

	pb.PermissionLevel = st.PermissionLevel

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type appAccessControlRequestPb struct {
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel AppPermissionLevel `json:"permission_level,omitempty"`

	ServicePrincipalName string `json:"service_principal_name,omitempty"`

	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appAccessControlRequestFromPb(pb *appAccessControlRequestPb) (*AppAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func appAccessControlResponseToPb(st *AppAccessControlResponse) (*appAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appAccessControlResponsePb{}
	pb.AllPermissions = st.AllPermissions

	pb.DisplayName = st.DisplayName

	pb.GroupName = st.GroupName

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type appAccessControlResponsePb struct {
	AllPermissions []AppPermission `json:"all_permissions,omitempty"`

	DisplayName string `json:"display_name,omitempty"`

	GroupName string `json:"group_name,omitempty"`

	ServicePrincipalName string `json:"service_principal_name,omitempty"`

	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appAccessControlResponseFromPb(pb *appAccessControlResponsePb) (*AppAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppAccessControlResponse{}
	st.AllPermissions = pb.AllPermissions
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func appDeploymentToPb(st *AppDeployment) (*appDeploymentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appDeploymentPb{}
	pb.CreateTime = st.CreateTime

	pb.Creator = st.Creator

	pb.DeploymentArtifacts = st.DeploymentArtifacts

	pb.DeploymentId = st.DeploymentId

	pb.Mode = st.Mode

	pb.SourceCodePath = st.SourceCodePath

	pb.Status = st.Status

	pb.UpdateTime = st.UpdateTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type appDeploymentPb struct {
	CreateTime string `json:"create_time,omitempty"`

	Creator string `json:"creator,omitempty"`

	DeploymentArtifacts *AppDeploymentArtifacts `json:"deployment_artifacts,omitempty"`

	DeploymentId string `json:"deployment_id,omitempty"`

	Mode AppDeploymentMode `json:"mode,omitempty"`

	SourceCodePath string `json:"source_code_path,omitempty"`

	Status *AppDeploymentStatus `json:"status,omitempty"`

	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appDeploymentFromPb(pb *appDeploymentPb) (*AppDeployment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppDeployment{}
	st.CreateTime = pb.CreateTime
	st.Creator = pb.Creator
	st.DeploymentArtifacts = pb.DeploymentArtifacts
	st.DeploymentId = pb.DeploymentId
	st.Mode = pb.Mode
	st.SourceCodePath = pb.SourceCodePath
	st.Status = pb.Status
	st.UpdateTime = pb.UpdateTime

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appDeploymentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appDeploymentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func appDeploymentArtifactsToPb(st *AppDeploymentArtifacts) (*appDeploymentArtifactsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appDeploymentArtifactsPb{}
	pb.SourceCodePath = st.SourceCodePath

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type appDeploymentArtifactsPb struct {
	SourceCodePath string `json:"source_code_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appDeploymentArtifactsFromPb(pb *appDeploymentArtifactsPb) (*AppDeploymentArtifacts, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppDeploymentArtifacts{}
	st.SourceCodePath = pb.SourceCodePath

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appDeploymentArtifactsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appDeploymentArtifactsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func appDeploymentStatusToPb(st *AppDeploymentStatus) (*appDeploymentStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appDeploymentStatusPb{}
	pb.Message = st.Message

	pb.State = st.State

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type appDeploymentStatusPb struct {
	Message string `json:"message,omitempty"`

	State AppDeploymentState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appDeploymentStatusFromPb(pb *appDeploymentStatusPb) (*AppDeploymentStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppDeploymentStatus{}
	st.Message = pb.Message
	st.State = pb.State

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appDeploymentStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appDeploymentStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func appPermissionToPb(st *AppPermission) (*appPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appPermissionPb{}
	pb.Inherited = st.Inherited

	pb.InheritedFromObject = st.InheritedFromObject

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type appPermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel AppPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appPermissionFromPb(pb *appPermissionPb) (*AppPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func appPermissionsToPb(st *AppPermissions) (*appPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appPermissionsPb{}
	pb.AccessControlList = st.AccessControlList

	pb.ObjectId = st.ObjectId

	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type appPermissionsPb struct {
	AccessControlList []AppAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appPermissionsFromPb(pb *appPermissionsPb) (*AppPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppPermissions{}
	st.AccessControlList = pb.AccessControlList
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func appPermissionsDescriptionToPb(st *AppPermissionsDescription) (*appPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appPermissionsDescriptionPb{}
	pb.Description = st.Description

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type appPermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`

	PermissionLevel AppPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appPermissionsDescriptionFromPb(pb *appPermissionsDescriptionPb) (*AppPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppPermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func appPermissionsRequestToPb(st *AppPermissionsRequest) (*appPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appPermissionsRequestPb{}
	pb.AccessControlList = st.AccessControlList

	pb.AppName = st.AppName

	return pb, nil
}

type appPermissionsRequestPb struct {
	AccessControlList []AppAccessControlRequest `json:"access_control_list,omitempty"`

	AppName string `json:"-" url:"-"`
}

func appPermissionsRequestFromPb(pb *appPermissionsRequestPb) (*AppPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppPermissionsRequest{}
	st.AccessControlList = pb.AccessControlList
	st.AppName = pb.AppName

	return st, nil
}

func appResourceToPb(st *AppResource) (*appResourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appResourcePb{}
	pb.Description = st.Description

	pb.Job = st.Job

	pb.Name = st.Name

	pb.Secret = st.Secret

	pb.ServingEndpoint = st.ServingEndpoint

	pb.SqlWarehouse = st.SqlWarehouse

	pb.UcSecurable = st.UcSecurable

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type appResourcePb struct {
	Description string `json:"description,omitempty"`

	Job *AppResourceJob `json:"job,omitempty"`

	Name string `json:"name"`

	Secret *AppResourceSecret `json:"secret,omitempty"`

	ServingEndpoint *AppResourceServingEndpoint `json:"serving_endpoint,omitempty"`

	SqlWarehouse *AppResourceSqlWarehouse `json:"sql_warehouse,omitempty"`

	UcSecurable *AppResourceUcSecurable `json:"uc_securable,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appResourceFromPb(pb *appResourcePb) (*AppResource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResource{}
	st.Description = pb.Description
	st.Job = pb.Job
	st.Name = pb.Name
	st.Secret = pb.Secret
	st.ServingEndpoint = pb.ServingEndpoint
	st.SqlWarehouse = pb.SqlWarehouse
	st.UcSecurable = pb.UcSecurable

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appResourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appResourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func appResourceJobToPb(st *AppResourceJob) (*appResourceJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appResourceJobPb{}
	pb.Id = st.Id

	pb.Permission = st.Permission

	return pb, nil
}

type appResourceJobPb struct {
	Id string `json:"id"`

	Permission AppResourceJobJobPermission `json:"permission"`
}

func appResourceJobFromPb(pb *appResourceJobPb) (*AppResourceJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResourceJob{}
	st.Id = pb.Id
	st.Permission = pb.Permission

	return st, nil
}

func appResourceSecretToPb(st *AppResourceSecret) (*appResourceSecretPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appResourceSecretPb{}
	pb.Key = st.Key

	pb.Permission = st.Permission

	pb.Scope = st.Scope

	return pb, nil
}

type appResourceSecretPb struct {
	Key string `json:"key"`

	Permission AppResourceSecretSecretPermission `json:"permission"`

	Scope string `json:"scope"`
}

func appResourceSecretFromPb(pb *appResourceSecretPb) (*AppResourceSecret, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResourceSecret{}
	st.Key = pb.Key
	st.Permission = pb.Permission
	st.Scope = pb.Scope

	return st, nil
}

func appResourceServingEndpointToPb(st *AppResourceServingEndpoint) (*appResourceServingEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appResourceServingEndpointPb{}
	pb.Name = st.Name

	pb.Permission = st.Permission

	return pb, nil
}

type appResourceServingEndpointPb struct {
	Name string `json:"name"`

	Permission AppResourceServingEndpointServingEndpointPermission `json:"permission"`
}

func appResourceServingEndpointFromPb(pb *appResourceServingEndpointPb) (*AppResourceServingEndpoint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResourceServingEndpoint{}
	st.Name = pb.Name
	st.Permission = pb.Permission

	return st, nil
}

func appResourceSqlWarehouseToPb(st *AppResourceSqlWarehouse) (*appResourceSqlWarehousePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appResourceSqlWarehousePb{}
	pb.Id = st.Id

	pb.Permission = st.Permission

	return pb, nil
}

type appResourceSqlWarehousePb struct {
	Id string `json:"id"`

	Permission AppResourceSqlWarehouseSqlWarehousePermission `json:"permission"`
}

func appResourceSqlWarehouseFromPb(pb *appResourceSqlWarehousePb) (*AppResourceSqlWarehouse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResourceSqlWarehouse{}
	st.Id = pb.Id
	st.Permission = pb.Permission

	return st, nil
}

func appResourceUcSecurableToPb(st *AppResourceUcSecurable) (*appResourceUcSecurablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appResourceUcSecurablePb{}
	pb.Permission = st.Permission

	pb.SecurableFullName = st.SecurableFullName

	pb.SecurableType = st.SecurableType

	return pb, nil
}

type appResourceUcSecurablePb struct {
	Permission AppResourceUcSecurableUcSecurablePermission `json:"permission"`

	SecurableFullName string `json:"securable_full_name"`

	SecurableType AppResourceUcSecurableUcSecurableType `json:"securable_type"`
}

func appResourceUcSecurableFromPb(pb *appResourceUcSecurablePb) (*AppResourceUcSecurable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResourceUcSecurable{}
	st.Permission = pb.Permission
	st.SecurableFullName = pb.SecurableFullName
	st.SecurableType = pb.SecurableType

	return st, nil
}

func applicationStatusToPb(st *ApplicationStatus) (*applicationStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &applicationStatusPb{}
	pb.Message = st.Message

	pb.State = st.State

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type applicationStatusPb struct {
	Message string `json:"message,omitempty"`

	State ApplicationState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func applicationStatusFromPb(pb *applicationStatusPb) (*ApplicationStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ApplicationStatus{}
	st.Message = pb.Message
	st.State = pb.State

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *applicationStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st applicationStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func computeStatusToPb(st *ComputeStatus) (*computeStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computeStatusPb{}
	pb.Message = st.Message

	pb.State = st.State

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type computeStatusPb struct {
	Message string `json:"message,omitempty"`

	State ComputeState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func computeStatusFromPb(pb *computeStatusPb) (*ComputeStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComputeStatus{}
	st.Message = pb.Message
	st.State = pb.State

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *computeStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st computeStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createAppDeploymentRequestToPb(st *CreateAppDeploymentRequest) (*createAppDeploymentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAppDeploymentRequestPb{}
	pb.AppDeployment = st.AppDeployment

	pb.AppName = st.AppName

	return pb, nil
}

type createAppDeploymentRequestPb struct {
	AppDeployment AppDeployment `json:"app_deployment"`

	AppName string `json:"-" url:"-"`
}

func createAppDeploymentRequestFromPb(pb *createAppDeploymentRequestPb) (*CreateAppDeploymentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAppDeploymentRequest{}
	st.AppDeployment = pb.AppDeployment
	st.AppName = pb.AppName

	return st, nil
}

func createAppRequestToPb(st *CreateAppRequest) (*createAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAppRequestPb{}
	pb.App = st.App

	pb.NoCompute = st.NoCompute

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createAppRequestPb struct {
	App App `json:"app"`

	NoCompute bool `json:"-" url:"no_compute,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createAppRequestFromPb(pb *createAppRequestPb) (*CreateAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAppRequest{}
	st.App = pb.App
	st.NoCompute = pb.NoCompute

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createAppRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createAppRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteAppRequestToPb(st *DeleteAppRequest) (*deleteAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAppRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type deleteAppRequestPb struct {
	Name string `json:"-" url:"-"`
}

func deleteAppRequestFromPb(pb *deleteAppRequestPb) (*DeleteAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAppRequest{}
	st.Name = pb.Name

	return st, nil
}

func getAppDeploymentRequestToPb(st *GetAppDeploymentRequest) (*getAppDeploymentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAppDeploymentRequestPb{}
	pb.AppName = st.AppName

	pb.DeploymentId = st.DeploymentId

	return pb, nil
}

type getAppDeploymentRequestPb struct {
	AppName string `json:"-" url:"-"`

	DeploymentId string `json:"-" url:"-"`
}

func getAppDeploymentRequestFromPb(pb *getAppDeploymentRequestPb) (*GetAppDeploymentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAppDeploymentRequest{}
	st.AppName = pb.AppName
	st.DeploymentId = pb.DeploymentId

	return st, nil
}

func getAppPermissionLevelsRequestToPb(st *GetAppPermissionLevelsRequest) (*getAppPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAppPermissionLevelsRequestPb{}
	pb.AppName = st.AppName

	return pb, nil
}

type getAppPermissionLevelsRequestPb struct {
	AppName string `json:"-" url:"-"`
}

func getAppPermissionLevelsRequestFromPb(pb *getAppPermissionLevelsRequestPb) (*GetAppPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAppPermissionLevelsRequest{}
	st.AppName = pb.AppName

	return st, nil
}

func getAppPermissionLevelsResponseToPb(st *GetAppPermissionLevelsResponse) (*getAppPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAppPermissionLevelsResponsePb{}
	pb.PermissionLevels = st.PermissionLevels

	return pb, nil
}

type getAppPermissionLevelsResponsePb struct {
	PermissionLevels []AppPermissionsDescription `json:"permission_levels,omitempty"`
}

func getAppPermissionLevelsResponseFromPb(pb *getAppPermissionLevelsResponsePb) (*GetAppPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAppPermissionLevelsResponse{}
	st.PermissionLevels = pb.PermissionLevels

	return st, nil
}

func getAppPermissionsRequestToPb(st *GetAppPermissionsRequest) (*getAppPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAppPermissionsRequestPb{}
	pb.AppName = st.AppName

	return pb, nil
}

type getAppPermissionsRequestPb struct {
	AppName string `json:"-" url:"-"`
}

func getAppPermissionsRequestFromPb(pb *getAppPermissionsRequestPb) (*GetAppPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAppPermissionsRequest{}
	st.AppName = pb.AppName

	return st, nil
}

func getAppRequestToPb(st *GetAppRequest) (*getAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAppRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type getAppRequestPb struct {
	Name string `json:"-" url:"-"`
}

func getAppRequestFromPb(pb *getAppRequestPb) (*GetAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAppRequest{}
	st.Name = pb.Name

	return st, nil
}

func listAppDeploymentsRequestToPb(st *ListAppDeploymentsRequest) (*listAppDeploymentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAppDeploymentsRequestPb{}
	pb.AppName = st.AppName

	pb.PageSize = st.PageSize

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listAppDeploymentsRequestPb struct {
	AppName string `json:"-" url:"-"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAppDeploymentsRequestFromPb(pb *listAppDeploymentsRequestPb) (*ListAppDeploymentsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAppDeploymentsRequest{}
	st.AppName = pb.AppName
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAppDeploymentsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAppDeploymentsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listAppDeploymentsResponseToPb(st *ListAppDeploymentsResponse) (*listAppDeploymentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAppDeploymentsResponsePb{}
	pb.AppDeployments = st.AppDeployments

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listAppDeploymentsResponsePb struct {
	AppDeployments []AppDeployment `json:"app_deployments,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAppDeploymentsResponseFromPb(pb *listAppDeploymentsResponsePb) (*ListAppDeploymentsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAppDeploymentsResponse{}
	st.AppDeployments = pb.AppDeployments
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAppDeploymentsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAppDeploymentsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listAppsRequestToPb(st *ListAppsRequest) (*listAppsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAppsRequestPb{}
	pb.PageSize = st.PageSize

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listAppsRequestPb struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAppsRequestFromPb(pb *listAppsRequestPb) (*ListAppsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAppsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAppsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAppsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listAppsResponseToPb(st *ListAppsResponse) (*listAppsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAppsResponsePb{}
	pb.Apps = st.Apps

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listAppsResponsePb struct {
	Apps []App `json:"apps,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAppsResponseFromPb(pb *listAppsResponsePb) (*ListAppsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAppsResponse{}
	st.Apps = pb.Apps
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAppsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAppsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func startAppRequestToPb(st *StartAppRequest) (*startAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startAppRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type startAppRequestPb struct {
	Name string `json:"-" url:"-"`
}

func startAppRequestFromPb(pb *startAppRequestPb) (*StartAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartAppRequest{}
	st.Name = pb.Name

	return st, nil
}

func stopAppRequestToPb(st *StopAppRequest) (*stopAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stopAppRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type stopAppRequestPb struct {
	Name string `json:"-" url:"-"`
}

func stopAppRequestFromPb(pb *stopAppRequestPb) (*StopAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StopAppRequest{}
	st.Name = pb.Name

	return st, nil
}

func updateAppRequestToPb(st *UpdateAppRequest) (*updateAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAppRequestPb{}
	pb.App = st.App

	pb.Name = st.Name

	return pb, nil
}

type updateAppRequestPb struct {
	App App `json:"app"`

	Name string `json:"-" url:"-"`
}

func updateAppRequestFromPb(pb *updateAppRequestPb) (*UpdateAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAppRequest{}
	st.App = pb.App
	st.Name = pb.Name

	return st, nil
}
