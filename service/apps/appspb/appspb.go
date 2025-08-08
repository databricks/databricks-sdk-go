// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package appspb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type AppPb struct {
	ActiveDeployment         *AppDeploymentPb     `json:"active_deployment,omitempty"`
	AppStatus                *ApplicationStatusPb `json:"app_status,omitempty"`
	BudgetPolicyId           string               `json:"budget_policy_id,omitempty"`
	ComputeStatus            *ComputeStatusPb     `json:"compute_status,omitempty"`
	CreateTime               string               `json:"create_time,omitempty"`
	Creator                  string               `json:"creator,omitempty"`
	DefaultSourceCodePath    string               `json:"default_source_code_path,omitempty"`
	Description              string               `json:"description,omitempty"`
	EffectiveBudgetPolicyId  string               `json:"effective_budget_policy_id,omitempty"`
	EffectiveUserApiScopes   []string             `json:"effective_user_api_scopes,omitempty"`
	Id                       string               `json:"id,omitempty"`
	Name                     string               `json:"name"`
	Oauth2AppClientId        string               `json:"oauth2_app_client_id,omitempty"`
	Oauth2AppIntegrationId   string               `json:"oauth2_app_integration_id,omitempty"`
	PendingDeployment        *AppDeploymentPb     `json:"pending_deployment,omitempty"`
	Resources                []AppResourcePb      `json:"resources,omitempty"`
	ServicePrincipalClientId string               `json:"service_principal_client_id,omitempty"`
	ServicePrincipalId       int64                `json:"service_principal_id,omitempty"`
	ServicePrincipalName     string               `json:"service_principal_name,omitempty"`
	UpdateTime               string               `json:"update_time,omitempty"`
	Updater                  string               `json:"updater,omitempty"`
	Url                      string               `json:"url,omitempty"`
	UserApiScopes            []string             `json:"user_api_scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AppPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AppPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppAccessControlRequestPb struct {
	GroupName            string               `json:"group_name,omitempty"`
	PermissionLevel      AppPermissionLevelPb `json:"permission_level,omitempty"`
	ServicePrincipalName string               `json:"service_principal_name,omitempty"`
	UserName             string               `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AppAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AppAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppAccessControlResponsePb struct {
	AllPermissions       []AppPermissionPb `json:"all_permissions,omitempty"`
	DisplayName          string            `json:"display_name,omitempty"`
	GroupName            string            `json:"group_name,omitempty"`
	ServicePrincipalName string            `json:"service_principal_name,omitempty"`
	UserName             string            `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AppAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AppAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppDeploymentPb struct {
	CreateTime          string                    `json:"create_time,omitempty"`
	Creator             string                    `json:"creator,omitempty"`
	DeploymentArtifacts *AppDeploymentArtifactsPb `json:"deployment_artifacts,omitempty"`
	DeploymentId        string                    `json:"deployment_id,omitempty"`
	Mode                AppDeploymentModePb       `json:"mode,omitempty"`
	SourceCodePath      string                    `json:"source_code_path,omitempty"`
	Status              *AppDeploymentStatusPb    `json:"status,omitempty"`
	UpdateTime          string                    `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AppDeploymentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AppDeploymentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppDeploymentArtifactsPb struct {
	SourceCodePath string `json:"source_code_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AppDeploymentArtifactsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AppDeploymentArtifactsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppDeploymentModePb string

const AppDeploymentModeAutoSync AppDeploymentModePb = `AUTO_SYNC`

const AppDeploymentModeSnapshot AppDeploymentModePb = `SNAPSHOT`

type AppDeploymentStatePb string

const AppDeploymentStateCancelled AppDeploymentStatePb = `CANCELLED`

const AppDeploymentStateFailed AppDeploymentStatePb = `FAILED`

const AppDeploymentStateInProgress AppDeploymentStatePb = `IN_PROGRESS`

const AppDeploymentStateSucceeded AppDeploymentStatePb = `SUCCEEDED`

type AppDeploymentStatusPb struct {
	Message string               `json:"message,omitempty"`
	State   AppDeploymentStatePb `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AppDeploymentStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AppDeploymentStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppPermissionPb struct {
	Inherited           bool                 `json:"inherited,omitempty"`
	InheritedFromObject []string             `json:"inherited_from_object,omitempty"`
	PermissionLevel     AppPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AppPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AppPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppPermissionLevelPb string

const AppPermissionLevelCanManage AppPermissionLevelPb = `CAN_MANAGE`

const AppPermissionLevelCanUse AppPermissionLevelPb = `CAN_USE`

type AppPermissionsPb struct {
	AccessControlList []AppAccessControlResponsePb `json:"access_control_list,omitempty"`
	ObjectId          string                       `json:"object_id,omitempty"`
	ObjectType        string                       `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AppPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AppPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppPermissionsDescriptionPb struct {
	Description     string               `json:"description,omitempty"`
	PermissionLevel AppPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AppPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AppPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppPermissionsRequestPb struct {
	AccessControlList []AppAccessControlRequestPb `json:"access_control_list,omitempty"`
	AppName           string                      `json:"-" url:"-"`
}

type AppResourcePb struct {
	Database        *AppResourceDatabasePb        `json:"database,omitempty"`
	Description     string                        `json:"description,omitempty"`
	Job             *AppResourceJobPb             `json:"job,omitempty"`
	Name            string                        `json:"name"`
	Secret          *AppResourceSecretPb          `json:"secret,omitempty"`
	ServingEndpoint *AppResourceServingEndpointPb `json:"serving_endpoint,omitempty"`
	SqlWarehouse    *AppResourceSqlWarehousePb    `json:"sql_warehouse,omitempty"`
	UcSecurable     *AppResourceUcSecurablePb     `json:"uc_securable,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AppResourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AppResourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppResourceDatabasePb struct {
	DatabaseName string                                  `json:"database_name"`
	InstanceName string                                  `json:"instance_name"`
	Permission   AppResourceDatabaseDatabasePermissionPb `json:"permission"`
}

type AppResourceDatabaseDatabasePermissionPb string

const AppResourceDatabaseDatabasePermissionCanConnectAndCreate AppResourceDatabaseDatabasePermissionPb = `CAN_CONNECT_AND_CREATE`

type AppResourceJobPb struct {
	Id         string                        `json:"id"`
	Permission AppResourceJobJobPermissionPb `json:"permission"`
}

type AppResourceJobJobPermissionPb string

const AppResourceJobJobPermissionCanManage AppResourceJobJobPermissionPb = `CAN_MANAGE`

const AppResourceJobJobPermissionCanManageRun AppResourceJobJobPermissionPb = `CAN_MANAGE_RUN`

const AppResourceJobJobPermissionCanView AppResourceJobJobPermissionPb = `CAN_VIEW`

const AppResourceJobJobPermissionIsOwner AppResourceJobJobPermissionPb = `IS_OWNER`

type AppResourceSecretPb struct {
	Key        string                              `json:"key"`
	Permission AppResourceSecretSecretPermissionPb `json:"permission"`
	Scope      string                              `json:"scope"`
}

type AppResourceSecretSecretPermissionPb string

const AppResourceSecretSecretPermissionManage AppResourceSecretSecretPermissionPb = `MANAGE`

const AppResourceSecretSecretPermissionRead AppResourceSecretSecretPermissionPb = `READ`

const AppResourceSecretSecretPermissionWrite AppResourceSecretSecretPermissionPb = `WRITE`

type AppResourceServingEndpointPb struct {
	Name       string                                                `json:"name"`
	Permission AppResourceServingEndpointServingEndpointPermissionPb `json:"permission"`
}

type AppResourceServingEndpointServingEndpointPermissionPb string

const AppResourceServingEndpointServingEndpointPermissionCanManage AppResourceServingEndpointServingEndpointPermissionPb = `CAN_MANAGE`

const AppResourceServingEndpointServingEndpointPermissionCanQuery AppResourceServingEndpointServingEndpointPermissionPb = `CAN_QUERY`

const AppResourceServingEndpointServingEndpointPermissionCanView AppResourceServingEndpointServingEndpointPermissionPb = `CAN_VIEW`

type AppResourceSqlWarehousePb struct {
	Id         string                                          `json:"id"`
	Permission AppResourceSqlWarehouseSqlWarehousePermissionPb `json:"permission"`
}

type AppResourceSqlWarehouseSqlWarehousePermissionPb string

const AppResourceSqlWarehouseSqlWarehousePermissionCanManage AppResourceSqlWarehouseSqlWarehousePermissionPb = `CAN_MANAGE`

const AppResourceSqlWarehouseSqlWarehousePermissionCanUse AppResourceSqlWarehouseSqlWarehousePermissionPb = `CAN_USE`

const AppResourceSqlWarehouseSqlWarehousePermissionIsOwner AppResourceSqlWarehouseSqlWarehousePermissionPb = `IS_OWNER`

type AppResourceUcSecurablePb struct {
	Permission        AppResourceUcSecurableUcSecurablePermissionPb `json:"permission"`
	SecurableFullName string                                        `json:"securable_full_name"`
	SecurableType     AppResourceUcSecurableUcSecurableTypePb       `json:"securable_type"`
}

type AppResourceUcSecurableUcSecurablePermissionPb string

const AppResourceUcSecurableUcSecurablePermissionReadVolume AppResourceUcSecurableUcSecurablePermissionPb = `READ_VOLUME`

const AppResourceUcSecurableUcSecurablePermissionWriteVolume AppResourceUcSecurableUcSecurablePermissionPb = `WRITE_VOLUME`

type AppResourceUcSecurableUcSecurableTypePb string

const AppResourceUcSecurableUcSecurableTypeVolume AppResourceUcSecurableUcSecurableTypePb = `VOLUME`

type ApplicationStatePb string

const ApplicationStateCrashed ApplicationStatePb = `CRASHED`

const ApplicationStateDeploying ApplicationStatePb = `DEPLOYING`

const ApplicationStateRunning ApplicationStatePb = `RUNNING`

const ApplicationStateUnavailable ApplicationStatePb = `UNAVAILABLE`

type ApplicationStatusPb struct {
	Message string             `json:"message,omitempty"`
	State   ApplicationStatePb `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ApplicationStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ApplicationStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ComputeStatePb string

const ComputeStateActive ComputeStatePb = `ACTIVE`

const ComputeStateDeleting ComputeStatePb = `DELETING`

const ComputeStateError ComputeStatePb = `ERROR`

const ComputeStateStarting ComputeStatePb = `STARTING`

const ComputeStateStopped ComputeStatePb = `STOPPED`

const ComputeStateStopping ComputeStatePb = `STOPPING`

const ComputeStateUpdating ComputeStatePb = `UPDATING`

type ComputeStatusPb struct {
	Message string         `json:"message,omitempty"`
	State   ComputeStatePb `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ComputeStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ComputeStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateAppDeploymentRequestPb struct {
	AppDeployment AppDeploymentPb `json:"app_deployment"`
	AppName       string          `json:"-" url:"-"`
}

type CreateAppRequestPb struct {
	App       AppPb `json:"app"`
	NoCompute bool  `json:"-" url:"no_compute,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateAppRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateAppRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteAppRequestPb struct {
	Name string `json:"-" url:"-"`
}

type GetAppDeploymentRequestPb struct {
	AppName      string `json:"-" url:"-"`
	DeploymentId string `json:"-" url:"-"`
}

type GetAppPermissionLevelsRequestPb struct {
	AppName string `json:"-" url:"-"`
}

type GetAppPermissionLevelsResponsePb struct {
	PermissionLevels []AppPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

type GetAppPermissionsRequestPb struct {
	AppName string `json:"-" url:"-"`
}

type GetAppRequestPb struct {
	Name string `json:"-" url:"-"`
}

type ListAppDeploymentsRequestPb struct {
	AppName   string `json:"-" url:"-"`
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAppDeploymentsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAppDeploymentsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAppDeploymentsResponsePb struct {
	AppDeployments []AppDeploymentPb `json:"app_deployments,omitempty"`
	NextPageToken  string            `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAppDeploymentsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAppDeploymentsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAppsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAppsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAppsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAppsResponsePb struct {
	Apps          []AppPb `json:"apps,omitempty"`
	NextPageToken string  `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAppsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAppsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StartAppRequestPb struct {
	Name string `json:"-" url:"-"`
}

type StopAppRequestPb struct {
	Name string `json:"-" url:"-"`
}

type UpdateAppRequestPb struct {
	App  AppPb  `json:"app"`
	Name string `json:"-" url:"-"`
}
