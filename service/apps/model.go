// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type App struct {
	// The active deployment of the app. A deployment is considered active when
	// it has been deployed to the app compute.
	ActiveDeployment *AppDeployment `json:"active_deployment,omitempty"`

	AppStatus *ApplicationStatus `json:"app_status,omitempty"`

	ComputeStatus *ComputeStatus `json:"compute_status,omitempty"`
	// The creation time of the app. Formatted timestamp in ISO 6801.
	CreateTime string `json:"create_time,omitempty"`
	// The email of the user that created the app.
	Creator string `json:"creator,omitempty"`
	// The default workspace file system path of the source code from which app
	// deployment are created. This field tracks the workspace source code path
	// of the last active deployment.
	DefaultSourceCodePath string `json:"default_source_code_path,omitempty"`
	// The description of the app.
	Description string `json:"description,omitempty"`
	// The unique identifier of the app.
	Id string `json:"id,omitempty"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name string `json:"name"`
	// The pending deployment of the app. A deployment is considered pending
	// when it is being prepared for deployment to the app compute.
	PendingDeployment *AppDeployment `json:"pending_deployment,omitempty"`
	// Resources for the app.
	Resources []AppResource `json:"resources,omitempty"`

	ServicePrincipalClientId string `json:"service_principal_client_id,omitempty"`

	ServicePrincipalId int64 `json:"service_principal_id,omitempty"`

	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The update time of the app. Formatted timestamp in ISO 6801.
	UpdateTime string `json:"update_time,omitempty"`
	// The email of the user that last updated the app.
	Updater string `json:"updater,omitempty"`
	// The URL of the app once it is deployed.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *App) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s App) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel AppPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AppAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppAccessControlResponse struct {
	// All permissions.
	AllPermissions []AppPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AppAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppDeployment struct {
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	CreateTime string `json:"create_time,omitempty"`
	// The email of the user creates the deployment.
	Creator string `json:"creator,omitempty"`
	// The deployment artifacts for an app.
	DeploymentArtifacts *AppDeploymentArtifacts `json:"deployment_artifacts,omitempty"`
	// The unique id of the deployment.
	DeploymentId string `json:"deployment_id,omitempty"`
	// The mode of which the deployment will manage the source code.
	Mode AppDeploymentMode `json:"mode,omitempty"`
	// The workspace file system path of the source code used to create the app
	// deployment. This is different from
	// `deployment_artifacts.source_code_path`, which is the path used by the
	// deployed app. The former refers to the original source code location of
	// the app in the workspace during deployment creation, whereas the latter
	// provides a system generated stable snapshotted source code path used by
	// the deployment.
	SourceCodePath string `json:"source_code_path,omitempty"`
	// Status and status message of the deployment
	Status *AppDeploymentStatus `json:"status,omitempty"`
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AppDeployment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppDeployment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppDeploymentArtifacts struct {
	// The snapshotted workspace file system path of the source code loaded by
	// the deployed app.
	SourceCodePath string `json:"source_code_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AppDeploymentArtifacts) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppDeploymentArtifacts) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppDeploymentMode string

const AppDeploymentModeAutoSync AppDeploymentMode = `AUTO_SYNC`

const AppDeploymentModeSnapshot AppDeploymentMode = `SNAPSHOT`

// String representation for [fmt.Print]
func (f *AppDeploymentMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppDeploymentMode) Set(v string) error {
	switch v {
	case `AUTO_SYNC`, `SNAPSHOT`:
		*f = AppDeploymentMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUTO_SYNC", "SNAPSHOT"`, v)
	}
}

// Type always returns AppDeploymentMode to satisfy [pflag.Value] interface
func (f *AppDeploymentMode) Type() string {
	return "AppDeploymentMode"
}

type AppDeploymentState string

const AppDeploymentStateCancelled AppDeploymentState = `CANCELLED`

const AppDeploymentStateFailed AppDeploymentState = `FAILED`

const AppDeploymentStateInProgress AppDeploymentState = `IN_PROGRESS`

const AppDeploymentStateSucceeded AppDeploymentState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *AppDeploymentState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppDeploymentState) Set(v string) error {
	switch v {
	case `CANCELLED`, `FAILED`, `IN_PROGRESS`, `SUCCEEDED`:
		*f = AppDeploymentState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "FAILED", "IN_PROGRESS", "SUCCEEDED"`, v)
	}
}

// Type always returns AppDeploymentState to satisfy [pflag.Value] interface
func (f *AppDeploymentState) Type() string {
	return "AppDeploymentState"
}

type AppDeploymentStatus struct {
	// Message corresponding with the deployment state.
	Message string `json:"message,omitempty"`
	// State of the deployment.
	State AppDeploymentState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AppDeploymentStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppDeploymentStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel AppPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AppPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type AppPermissionLevel string

const AppPermissionLevelCanManage AppPermissionLevel = `CAN_MANAGE`

const AppPermissionLevelCanUse AppPermissionLevel = `CAN_USE`

// String representation for [fmt.Print]
func (f *AppPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_USE`:
		*f = AppPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_USE"`, v)
	}
}

// Type always returns AppPermissionLevel to satisfy [pflag.Value] interface
func (f *AppPermissionLevel) Type() string {
	return "AppPermissionLevel"
}

type AppPermissions struct {
	AccessControlList []AppAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AppPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppPermissionsDescription struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel AppPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AppPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppPermissionsRequest struct {
	AccessControlList []AppAccessControlRequest `json:"access_control_list,omitempty"`
	// The app for which to get or manage permissions.
	AppName string `json:"-" url:"-"`
}

type AppResource struct {
	// Description of the App Resource.
	Description string `json:"description,omitempty"`

	Job *AppResourceJob `json:"job,omitempty"`
	// Name of the App Resource.
	Name string `json:"name"`

	Secret *AppResourceSecret `json:"secret,omitempty"`

	ServingEndpoint *AppResourceServingEndpoint `json:"serving_endpoint,omitempty"`

	SqlWarehouse *AppResourceSqlWarehouse `json:"sql_warehouse,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AppResource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppResource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppResourceJob struct {
	// Id of the job to grant permission on.
	Id string `json:"id"`
	// Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE",
	// "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW".
	Permission AppResourceJobJobPermission `json:"permission"`
}

type AppResourceJobJobPermission string

const AppResourceJobJobPermissionCanManage AppResourceJobJobPermission = `CAN_MANAGE`

const AppResourceJobJobPermissionCanManageRun AppResourceJobJobPermission = `CAN_MANAGE_RUN`

const AppResourceJobJobPermissionCanView AppResourceJobJobPermission = `CAN_VIEW`

const AppResourceJobJobPermissionIsOwner AppResourceJobJobPermission = `IS_OWNER`

// String representation for [fmt.Print]
func (f *AppResourceJobJobPermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppResourceJobJobPermission) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_MANAGE_RUN`, `CAN_VIEW`, `IS_OWNER`:
		*f = AppResourceJobJobPermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_MANAGE_RUN", "CAN_VIEW", "IS_OWNER"`, v)
	}
}

// Type always returns AppResourceJobJobPermission to satisfy [pflag.Value] interface
func (f *AppResourceJobJobPermission) Type() string {
	return "AppResourceJobJobPermission"
}

type AppResourceSecret struct {
	// Key of the secret to grant permission on.
	Key string `json:"key"`
	// Permission to grant on the secret scope. For secrets, only one permission
	// is allowed. Permission must be one of: "READ", "WRITE", "MANAGE".
	Permission AppResourceSecretSecretPermission `json:"permission"`
	// Scope of the secret to grant permission on.
	Scope string `json:"scope"`
}

// Permission to grant on the secret scope. Supported permissions are: "READ",
// "WRITE", "MANAGE".
type AppResourceSecretSecretPermission string

const AppResourceSecretSecretPermissionManage AppResourceSecretSecretPermission = `MANAGE`

const AppResourceSecretSecretPermissionRead AppResourceSecretSecretPermission = `READ`

const AppResourceSecretSecretPermissionWrite AppResourceSecretSecretPermission = `WRITE`

// String representation for [fmt.Print]
func (f *AppResourceSecretSecretPermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppResourceSecretSecretPermission) Set(v string) error {
	switch v {
	case `MANAGE`, `READ`, `WRITE`:
		*f = AppResourceSecretSecretPermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MANAGE", "READ", "WRITE"`, v)
	}
}

// Type always returns AppResourceSecretSecretPermission to satisfy [pflag.Value] interface
func (f *AppResourceSecretSecretPermission) Type() string {
	return "AppResourceSecretSecretPermission"
}

type AppResourceServingEndpoint struct {
	// Name of the serving endpoint to grant permission on.
	Name string `json:"name"`
	// Permission to grant on the serving endpoint. Supported permissions are:
	// "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW".
	Permission AppResourceServingEndpointServingEndpointPermission `json:"permission"`
}

type AppResourceServingEndpointServingEndpointPermission string

const AppResourceServingEndpointServingEndpointPermissionCanManage AppResourceServingEndpointServingEndpointPermission = `CAN_MANAGE`

const AppResourceServingEndpointServingEndpointPermissionCanQuery AppResourceServingEndpointServingEndpointPermission = `CAN_QUERY`

const AppResourceServingEndpointServingEndpointPermissionCanView AppResourceServingEndpointServingEndpointPermission = `CAN_VIEW`

// String representation for [fmt.Print]
func (f *AppResourceServingEndpointServingEndpointPermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppResourceServingEndpointServingEndpointPermission) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_QUERY`, `CAN_VIEW`:
		*f = AppResourceServingEndpointServingEndpointPermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW"`, v)
	}
}

// Type always returns AppResourceServingEndpointServingEndpointPermission to satisfy [pflag.Value] interface
func (f *AppResourceServingEndpointServingEndpointPermission) Type() string {
	return "AppResourceServingEndpointServingEndpointPermission"
}

type AppResourceSqlWarehouse struct {
	// Id of the SQL warehouse to grant permission on.
	Id string `json:"id"`
	// Permission to grant on the SQL warehouse. Supported permissions are:
	// "CAN_MANAGE", "CAN_USE", "IS_OWNER".
	Permission AppResourceSqlWarehouseSqlWarehousePermission `json:"permission"`
}

type AppResourceSqlWarehouseSqlWarehousePermission string

const AppResourceSqlWarehouseSqlWarehousePermissionCanManage AppResourceSqlWarehouseSqlWarehousePermission = `CAN_MANAGE`

const AppResourceSqlWarehouseSqlWarehousePermissionCanUse AppResourceSqlWarehouseSqlWarehousePermission = `CAN_USE`

const AppResourceSqlWarehouseSqlWarehousePermissionIsOwner AppResourceSqlWarehouseSqlWarehousePermission = `IS_OWNER`

// String representation for [fmt.Print]
func (f *AppResourceSqlWarehouseSqlWarehousePermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppResourceSqlWarehouseSqlWarehousePermission) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_USE`, `IS_OWNER`:
		*f = AppResourceSqlWarehouseSqlWarehousePermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_USE", "IS_OWNER"`, v)
	}
}

// Type always returns AppResourceSqlWarehouseSqlWarehousePermission to satisfy [pflag.Value] interface
func (f *AppResourceSqlWarehouseSqlWarehousePermission) Type() string {
	return "AppResourceSqlWarehouseSqlWarehousePermission"
}

type ApplicationState string

const ApplicationStateCrashed ApplicationState = `CRASHED`

const ApplicationStateDeploying ApplicationState = `DEPLOYING`

const ApplicationStateRunning ApplicationState = `RUNNING`

const ApplicationStateUnavailable ApplicationState = `UNAVAILABLE`

// String representation for [fmt.Print]
func (f *ApplicationState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ApplicationState) Set(v string) error {
	switch v {
	case `CRASHED`, `DEPLOYING`, `RUNNING`, `UNAVAILABLE`:
		*f = ApplicationState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CRASHED", "DEPLOYING", "RUNNING", "UNAVAILABLE"`, v)
	}
}

// Type always returns ApplicationState to satisfy [pflag.Value] interface
func (f *ApplicationState) Type() string {
	return "ApplicationState"
}

type ApplicationStatus struct {
	// Application status message
	Message string `json:"message,omitempty"`
	// State of the application.
	State ApplicationState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ApplicationStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ApplicationStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ComputeState string

const ComputeStateActive ComputeState = `ACTIVE`

const ComputeStateDeleting ComputeState = `DELETING`

const ComputeStateError ComputeState = `ERROR`

const ComputeStateStarting ComputeState = `STARTING`

const ComputeStateStopped ComputeState = `STOPPED`

const ComputeStateStopping ComputeState = `STOPPING`

const ComputeStateUpdating ComputeState = `UPDATING`

// String representation for [fmt.Print]
func (f *ComputeState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ComputeState) Set(v string) error {
	switch v {
	case `ACTIVE`, `DELETING`, `ERROR`, `STARTING`, `STOPPED`, `STOPPING`, `UPDATING`:
		*f = ComputeState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DELETING", "ERROR", "STARTING", "STOPPED", "STOPPING", "UPDATING"`, v)
	}
}

// Type always returns ComputeState to satisfy [pflag.Value] interface
func (f *ComputeState) Type() string {
	return "ComputeState"
}

type ComputeStatus struct {
	// Compute status message
	Message string `json:"message,omitempty"`
	// State of the app compute.
	State ComputeState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ComputeStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComputeStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Create an app deployment
type CreateAppDeploymentRequest struct {
	AppDeployment *AppDeployment `json:"app_deployment,omitempty"`
	// The name of the app.
	AppName string `json:"-" url:"-"`
}

// Create an app
type CreateAppRequest struct {
	App *App `json:"app,omitempty"`
	// If true, the app will not be started after creation.
	NoCompute bool `json:"-" url:"no_compute,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateAppRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAppRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete an app
type DeleteAppRequest struct {
	// The name of the app.
	Name string `json:"-" url:"-"`
}

// Get an app deployment
type GetAppDeploymentRequest struct {
	// The name of the app.
	AppName string `json:"-" url:"-"`
	// The unique id of the deployment.
	DeploymentId string `json:"-" url:"-"`
}

// Get app permission levels
type GetAppPermissionLevelsRequest struct {
	// The app for which to get or manage permissions.
	AppName string `json:"-" url:"-"`
}

type GetAppPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []AppPermissionsDescription `json:"permission_levels,omitempty"`
}

// Get app permissions
type GetAppPermissionsRequest struct {
	// The app for which to get or manage permissions.
	AppName string `json:"-" url:"-"`
}

// Get an app
type GetAppRequest struct {
	// The name of the app.
	Name string `json:"-" url:"-"`
}

// List app deployments
type ListAppDeploymentsRequest struct {
	// The name of the app.
	AppName string `json:"-" url:"-"`
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAppDeploymentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAppDeploymentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAppDeploymentsResponse struct {
	// Deployment history of the app.
	AppDeployments []AppDeployment `json:"app_deployments,omitempty"`
	// Pagination token to request the next page of apps.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAppDeploymentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAppDeploymentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List apps
type ListAppsRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAppsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAppsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAppsResponse struct {
	Apps []App `json:"apps,omitempty"`
	// Pagination token to request the next page of apps.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAppsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAppsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StartAppRequest struct {
	// The name of the app.
	Name string `json:"-" url:"-"`
}

type StopAppRequest struct {
	// The name of the app.
	Name string `json:"-" url:"-"`
}

// Update an app
type UpdateAppRequest struct {
	App *App `json:"app,omitempty"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name string `json:"-" url:"-"`
}
