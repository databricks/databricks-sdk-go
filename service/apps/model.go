// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type App struct {
	// The active deployment of the app. A deployment is considered active when
	// it has been deployed to the app compute.
	ActiveDeployment *AppDeployment `json:"active_deployment,omitempty"`

	AppStatus *ApplicationStatus `json:"app_status,omitempty"`

	BudgetPolicyId string `json:"budget_policy_id,omitempty"`

	ComputeSize ComputeSize `json:"compute_size,omitempty"`

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

	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`

	EffectiveUsagePolicyId string `json:"effective_usage_policy_id,omitempty"`
	// The effective api scopes granted to the user access token.
	EffectiveUserApiScopes []string `json:"effective_user_api_scopes,omitempty"`
	// Git repository configuration for app deployments. When specified,
	// deployments can reference code from this repository by providing only the
	// git reference (branch, tag, or commit).
	GitRepository *GitRepository `json:"git_repository,omitempty"`
	// The unique identifier of the app.
	Id string `json:"id,omitempty"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name string `json:"name"`

	Oauth2AppClientId string `json:"oauth2_app_client_id,omitempty"`

	Oauth2AppIntegrationId string `json:"oauth2_app_integration_id,omitempty"`
	// The pending deployment of the app. A deployment is considered pending
	// when it is being prepared for deployment to the app compute.
	PendingDeployment *AppDeployment `json:"pending_deployment,omitempty"`
	// Resources for the app.
	Resources []AppResource `json:"resources,omitempty"`

	ServicePrincipalClientId string `json:"service_principal_client_id,omitempty"`

	ServicePrincipalId int64 `json:"service_principal_id,omitempty"`

	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Name of the space this app belongs to.
	Space string `json:"space,omitempty"`
	// The update time of the app. Formatted timestamp in ISO 6801.
	UpdateTime string `json:"update_time,omitempty"`
	// The email of the user that last updated the app.
	Updater string `json:"updater,omitempty"`
	// The URL of the app once it is deployed.
	Url string `json:"url,omitempty"`

	UsagePolicyId string `json:"usage_policy_id,omitempty"`

	UserApiScopes []string `json:"user_api_scopes,omitempty"`

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
	// The command with which to run the app. This will override the command
	// specified in the app.yaml file.
	Command []string `json:"command,omitempty"`
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	CreateTime string `json:"create_time,omitempty"`
	// The email of the user creates the deployment.
	Creator string `json:"creator,omitempty"`
	// The deployment artifacts for an app.
	DeploymentArtifacts *AppDeploymentArtifacts `json:"deployment_artifacts,omitempty"`
	// The unique id of the deployment.
	DeploymentId string `json:"deployment_id,omitempty"`
	// The environment variables to set in the app runtime environment. This
	// will override the environment variables specified in the app.yaml file.
	EnvVars []EnvVar `json:"env_vars,omitempty"`
	// Git repository to use as the source for the app deployment.
	GitSource *GitSource `json:"git_source,omitempty"`
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

// Values returns all possible values for AppDeploymentMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppDeploymentMode) Values() []AppDeploymentMode {
	return []AppDeploymentMode{
		AppDeploymentModeAutoSync,
		AppDeploymentModeSnapshot,
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

// Values returns all possible values for AppDeploymentState.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppDeploymentState) Values() []AppDeploymentState {
	return []AppDeploymentState{
		AppDeploymentStateCancelled,
		AppDeploymentStateFailed,
		AppDeploymentStateInProgress,
		AppDeploymentStateSucceeded,
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

// App manifest definition
type AppManifest struct {
	// Description of the app defined by manifest author / publisher
	Description string `json:"description,omitempty"`
	// Name of the app defined by manifest author / publisher
	Name string `json:"name"`

	ResourceSpecs []AppManifestAppResourceSpec `json:"resource_specs,omitempty"`
	// The manifest schema version, for now only 1 is allowed
	Version int `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AppManifest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppManifest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppManifestAppResourceExperimentSpec struct {
	Permission AppManifestAppResourceExperimentSpecExperimentPermission `json:"permission"`
}

type AppManifestAppResourceExperimentSpecExperimentPermission string

const AppManifestAppResourceExperimentSpecExperimentPermissionCanEdit AppManifestAppResourceExperimentSpecExperimentPermission = `CAN_EDIT`

const AppManifestAppResourceExperimentSpecExperimentPermissionCanManage AppManifestAppResourceExperimentSpecExperimentPermission = `CAN_MANAGE`

const AppManifestAppResourceExperimentSpecExperimentPermissionCanRead AppManifestAppResourceExperimentSpecExperimentPermission = `CAN_READ`

// String representation for [fmt.Print]
func (f *AppManifestAppResourceExperimentSpecExperimentPermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppManifestAppResourceExperimentSpecExperimentPermission) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`:
		*f = AppManifestAppResourceExperimentSpecExperimentPermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_READ"`, v)
	}
}

// Values returns all possible values for AppManifestAppResourceExperimentSpecExperimentPermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppManifestAppResourceExperimentSpecExperimentPermission) Values() []AppManifestAppResourceExperimentSpecExperimentPermission {
	return []AppManifestAppResourceExperimentSpecExperimentPermission{
		AppManifestAppResourceExperimentSpecExperimentPermissionCanEdit,
		AppManifestAppResourceExperimentSpecExperimentPermissionCanManage,
		AppManifestAppResourceExperimentSpecExperimentPermissionCanRead,
	}
}

// Type always returns AppManifestAppResourceExperimentSpecExperimentPermission to satisfy [pflag.Value] interface
func (f *AppManifestAppResourceExperimentSpecExperimentPermission) Type() string {
	return "AppManifestAppResourceExperimentSpecExperimentPermission"
}

type AppManifestAppResourceJobSpec struct {
	// Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE",
	// "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW".
	Permission AppManifestAppResourceJobSpecJobPermission `json:"permission"`
}

type AppManifestAppResourceJobSpecJobPermission string

const AppManifestAppResourceJobSpecJobPermissionCanManage AppManifestAppResourceJobSpecJobPermission = `CAN_MANAGE`

const AppManifestAppResourceJobSpecJobPermissionCanManageRun AppManifestAppResourceJobSpecJobPermission = `CAN_MANAGE_RUN`

const AppManifestAppResourceJobSpecJobPermissionCanView AppManifestAppResourceJobSpecJobPermission = `CAN_VIEW`

const AppManifestAppResourceJobSpecJobPermissionIsOwner AppManifestAppResourceJobSpecJobPermission = `IS_OWNER`

// String representation for [fmt.Print]
func (f *AppManifestAppResourceJobSpecJobPermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppManifestAppResourceJobSpecJobPermission) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_MANAGE_RUN`, `CAN_VIEW`, `IS_OWNER`:
		*f = AppManifestAppResourceJobSpecJobPermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_MANAGE_RUN", "CAN_VIEW", "IS_OWNER"`, v)
	}
}

// Values returns all possible values for AppManifestAppResourceJobSpecJobPermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppManifestAppResourceJobSpecJobPermission) Values() []AppManifestAppResourceJobSpecJobPermission {
	return []AppManifestAppResourceJobSpecJobPermission{
		AppManifestAppResourceJobSpecJobPermissionCanManage,
		AppManifestAppResourceJobSpecJobPermissionCanManageRun,
		AppManifestAppResourceJobSpecJobPermissionCanView,
		AppManifestAppResourceJobSpecJobPermissionIsOwner,
	}
}

// Type always returns AppManifestAppResourceJobSpecJobPermission to satisfy [pflag.Value] interface
func (f *AppManifestAppResourceJobSpecJobPermission) Type() string {
	return "AppManifestAppResourceJobSpecJobPermission"
}

type AppManifestAppResourceSecretSpec struct {
	// Permission to grant on the secret scope. For secrets, only one permission
	// is allowed. Permission must be one of: "READ", "WRITE", "MANAGE".
	Permission AppManifestAppResourceSecretSpecSecretPermission `json:"permission"`
}

// Permission to grant on the secret scope. Supported permissions are: "READ",
// "WRITE", "MANAGE".
type AppManifestAppResourceSecretSpecSecretPermission string

const AppManifestAppResourceSecretSpecSecretPermissionManage AppManifestAppResourceSecretSpecSecretPermission = `MANAGE`

const AppManifestAppResourceSecretSpecSecretPermissionRead AppManifestAppResourceSecretSpecSecretPermission = `READ`

const AppManifestAppResourceSecretSpecSecretPermissionWrite AppManifestAppResourceSecretSpecSecretPermission = `WRITE`

// String representation for [fmt.Print]
func (f *AppManifestAppResourceSecretSpecSecretPermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppManifestAppResourceSecretSpecSecretPermission) Set(v string) error {
	switch v {
	case `MANAGE`, `READ`, `WRITE`:
		*f = AppManifestAppResourceSecretSpecSecretPermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MANAGE", "READ", "WRITE"`, v)
	}
}

// Values returns all possible values for AppManifestAppResourceSecretSpecSecretPermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppManifestAppResourceSecretSpecSecretPermission) Values() []AppManifestAppResourceSecretSpecSecretPermission {
	return []AppManifestAppResourceSecretSpecSecretPermission{
		AppManifestAppResourceSecretSpecSecretPermissionManage,
		AppManifestAppResourceSecretSpecSecretPermissionRead,
		AppManifestAppResourceSecretSpecSecretPermissionWrite,
	}
}

// Type always returns AppManifestAppResourceSecretSpecSecretPermission to satisfy [pflag.Value] interface
func (f *AppManifestAppResourceSecretSpecSecretPermission) Type() string {
	return "AppManifestAppResourceSecretSpecSecretPermission"
}

type AppManifestAppResourceServingEndpointSpec struct {
	// Permission to grant on the serving endpoint. Supported permissions are:
	// "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW".
	Permission AppManifestAppResourceServingEndpointSpecServingEndpointPermission `json:"permission"`
}

type AppManifestAppResourceServingEndpointSpecServingEndpointPermission string

const AppManifestAppResourceServingEndpointSpecServingEndpointPermissionCanManage AppManifestAppResourceServingEndpointSpecServingEndpointPermission = `CAN_MANAGE`

const AppManifestAppResourceServingEndpointSpecServingEndpointPermissionCanQuery AppManifestAppResourceServingEndpointSpecServingEndpointPermission = `CAN_QUERY`

const AppManifestAppResourceServingEndpointSpecServingEndpointPermissionCanView AppManifestAppResourceServingEndpointSpecServingEndpointPermission = `CAN_VIEW`

// String representation for [fmt.Print]
func (f *AppManifestAppResourceServingEndpointSpecServingEndpointPermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppManifestAppResourceServingEndpointSpecServingEndpointPermission) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_QUERY`, `CAN_VIEW`:
		*f = AppManifestAppResourceServingEndpointSpecServingEndpointPermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW"`, v)
	}
}

// Values returns all possible values for AppManifestAppResourceServingEndpointSpecServingEndpointPermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppManifestAppResourceServingEndpointSpecServingEndpointPermission) Values() []AppManifestAppResourceServingEndpointSpecServingEndpointPermission {
	return []AppManifestAppResourceServingEndpointSpecServingEndpointPermission{
		AppManifestAppResourceServingEndpointSpecServingEndpointPermissionCanManage,
		AppManifestAppResourceServingEndpointSpecServingEndpointPermissionCanQuery,
		AppManifestAppResourceServingEndpointSpecServingEndpointPermissionCanView,
	}
}

// Type always returns AppManifestAppResourceServingEndpointSpecServingEndpointPermission to satisfy [pflag.Value] interface
func (f *AppManifestAppResourceServingEndpointSpecServingEndpointPermission) Type() string {
	return "AppManifestAppResourceServingEndpointSpecServingEndpointPermission"
}

// AppResource related fields are copied from app.proto but excludes resource
// identifiers (e.g. name, id, key, scope, etc.)
type AppManifestAppResourceSpec struct {
	// Description of the App Resource.
	Description string `json:"description,omitempty"`

	ExperimentSpec *AppManifestAppResourceExperimentSpec `json:"experiment_spec,omitempty"`

	JobSpec *AppManifestAppResourceJobSpec `json:"job_spec,omitempty"`
	// Name of the App Resource.
	Name string `json:"name"`

	SecretSpec *AppManifestAppResourceSecretSpec `json:"secret_spec,omitempty"`

	ServingEndpointSpec *AppManifestAppResourceServingEndpointSpec `json:"serving_endpoint_spec,omitempty"`

	SqlWarehouseSpec *AppManifestAppResourceSqlWarehouseSpec `json:"sql_warehouse_spec,omitempty"`

	UcSecurableSpec *AppManifestAppResourceUcSecurableSpec `json:"uc_securable_spec,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AppManifestAppResourceSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppManifestAppResourceSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppManifestAppResourceSqlWarehouseSpec struct {
	// Permission to grant on the SQL warehouse. Supported permissions are:
	// "CAN_MANAGE", "CAN_USE", "IS_OWNER".
	Permission AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermission `json:"permission"`
}

type AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermission string

const AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermissionCanManage AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermission = `CAN_MANAGE`

const AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermissionCanUse AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermission = `CAN_USE`

const AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermissionIsOwner AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermission = `IS_OWNER`

// String representation for [fmt.Print]
func (f *AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermission) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_USE`, `IS_OWNER`:
		*f = AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_USE", "IS_OWNER"`, v)
	}
}

// Values returns all possible values for AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermission) Values() []AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermission {
	return []AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermission{
		AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermissionCanManage,
		AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermissionCanUse,
		AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermissionIsOwner,
	}
}

// Type always returns AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermission to satisfy [pflag.Value] interface
func (f *AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermission) Type() string {
	return "AppManifestAppResourceSqlWarehouseSpecSqlWarehousePermission"
}

type AppManifestAppResourceUcSecurableSpec struct {
	Permission AppManifestAppResourceUcSecurableSpecUcSecurablePermission `json:"permission"`

	SecurableType AppManifestAppResourceUcSecurableSpecUcSecurableType `json:"securable_type"`
}

type AppManifestAppResourceUcSecurableSpecUcSecurablePermission string

const AppManifestAppResourceUcSecurableSpecUcSecurablePermissionExecute AppManifestAppResourceUcSecurableSpecUcSecurablePermission = `EXECUTE`

const AppManifestAppResourceUcSecurableSpecUcSecurablePermissionManage AppManifestAppResourceUcSecurableSpecUcSecurablePermission = `MANAGE`

const AppManifestAppResourceUcSecurableSpecUcSecurablePermissionReadVolume AppManifestAppResourceUcSecurableSpecUcSecurablePermission = `READ_VOLUME`

const AppManifestAppResourceUcSecurableSpecUcSecurablePermissionSelect AppManifestAppResourceUcSecurableSpecUcSecurablePermission = `SELECT`

const AppManifestAppResourceUcSecurableSpecUcSecurablePermissionUseConnection AppManifestAppResourceUcSecurableSpecUcSecurablePermission = `USE_CONNECTION`

const AppManifestAppResourceUcSecurableSpecUcSecurablePermissionWriteVolume AppManifestAppResourceUcSecurableSpecUcSecurablePermission = `WRITE_VOLUME`

// String representation for [fmt.Print]
func (f *AppManifestAppResourceUcSecurableSpecUcSecurablePermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppManifestAppResourceUcSecurableSpecUcSecurablePermission) Set(v string) error {
	switch v {
	case `EXECUTE`, `MANAGE`, `READ_VOLUME`, `SELECT`, `USE_CONNECTION`, `WRITE_VOLUME`:
		*f = AppManifestAppResourceUcSecurableSpecUcSecurablePermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXECUTE", "MANAGE", "READ_VOLUME", "SELECT", "USE_CONNECTION", "WRITE_VOLUME"`, v)
	}
}

// Values returns all possible values for AppManifestAppResourceUcSecurableSpecUcSecurablePermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppManifestAppResourceUcSecurableSpecUcSecurablePermission) Values() []AppManifestAppResourceUcSecurableSpecUcSecurablePermission {
	return []AppManifestAppResourceUcSecurableSpecUcSecurablePermission{
		AppManifestAppResourceUcSecurableSpecUcSecurablePermissionExecute,
		AppManifestAppResourceUcSecurableSpecUcSecurablePermissionManage,
		AppManifestAppResourceUcSecurableSpecUcSecurablePermissionReadVolume,
		AppManifestAppResourceUcSecurableSpecUcSecurablePermissionSelect,
		AppManifestAppResourceUcSecurableSpecUcSecurablePermissionUseConnection,
		AppManifestAppResourceUcSecurableSpecUcSecurablePermissionWriteVolume,
	}
}

// Type always returns AppManifestAppResourceUcSecurableSpecUcSecurablePermission to satisfy [pflag.Value] interface
func (f *AppManifestAppResourceUcSecurableSpecUcSecurablePermission) Type() string {
	return "AppManifestAppResourceUcSecurableSpecUcSecurablePermission"
}

type AppManifestAppResourceUcSecurableSpecUcSecurableType string

const AppManifestAppResourceUcSecurableSpecUcSecurableTypeConnection AppManifestAppResourceUcSecurableSpecUcSecurableType = `CONNECTION`

const AppManifestAppResourceUcSecurableSpecUcSecurableTypeFunction AppManifestAppResourceUcSecurableSpecUcSecurableType = `FUNCTION`

const AppManifestAppResourceUcSecurableSpecUcSecurableTypeTable AppManifestAppResourceUcSecurableSpecUcSecurableType = `TABLE`

const AppManifestAppResourceUcSecurableSpecUcSecurableTypeVolume AppManifestAppResourceUcSecurableSpecUcSecurableType = `VOLUME`

// String representation for [fmt.Print]
func (f *AppManifestAppResourceUcSecurableSpecUcSecurableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppManifestAppResourceUcSecurableSpecUcSecurableType) Set(v string) error {
	switch v {
	case `CONNECTION`, `FUNCTION`, `TABLE`, `VOLUME`:
		*f = AppManifestAppResourceUcSecurableSpecUcSecurableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONNECTION", "FUNCTION", "TABLE", "VOLUME"`, v)
	}
}

// Values returns all possible values for AppManifestAppResourceUcSecurableSpecUcSecurableType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppManifestAppResourceUcSecurableSpecUcSecurableType) Values() []AppManifestAppResourceUcSecurableSpecUcSecurableType {
	return []AppManifestAppResourceUcSecurableSpecUcSecurableType{
		AppManifestAppResourceUcSecurableSpecUcSecurableTypeConnection,
		AppManifestAppResourceUcSecurableSpecUcSecurableTypeFunction,
		AppManifestAppResourceUcSecurableSpecUcSecurableTypeTable,
		AppManifestAppResourceUcSecurableSpecUcSecurableTypeVolume,
	}
}

// Type always returns AppManifestAppResourceUcSecurableSpecUcSecurableType to satisfy [pflag.Value] interface
func (f *AppManifestAppResourceUcSecurableSpecUcSecurableType) Type() string {
	return "AppManifestAppResourceUcSecurableSpecUcSecurableType"
}

type AppPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

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

// Values returns all possible values for AppPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppPermissionLevel) Values() []AppPermissionLevel {
	return []AppPermissionLevel{
		AppPermissionLevelCanManage,
		AppPermissionLevelCanUse,
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
	App *AppResourceApp `json:"app,omitempty"`

	Database *AppResourceDatabase `json:"database,omitempty"`
	// Description of the App Resource.
	Description string `json:"description,omitempty"`

	Experiment *AppResourceExperiment `json:"experiment,omitempty"`

	GenieSpace *AppResourceGenieSpace `json:"genie_space,omitempty"`

	Job *AppResourceJob `json:"job,omitempty"`
	// Name of the App Resource.
	Name string `json:"name"`

	Secret *AppResourceSecret `json:"secret,omitempty"`

	ServingEndpoint *AppResourceServingEndpoint `json:"serving_endpoint,omitempty"`

	SqlWarehouse *AppResourceSqlWarehouse `json:"sql_warehouse,omitempty"`

	UcSecurable *AppResourceUcSecurable `json:"uc_securable,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AppResource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppResource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppResourceApp struct {
}

type AppResourceDatabase struct {
	DatabaseName string `json:"database_name"`

	InstanceName string `json:"instance_name"`

	Permission AppResourceDatabaseDatabasePermission `json:"permission"`
}

type AppResourceDatabaseDatabasePermission string

const AppResourceDatabaseDatabasePermissionCanConnectAndCreate AppResourceDatabaseDatabasePermission = `CAN_CONNECT_AND_CREATE`

// String representation for [fmt.Print]
func (f *AppResourceDatabaseDatabasePermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppResourceDatabaseDatabasePermission) Set(v string) error {
	switch v {
	case `CAN_CONNECT_AND_CREATE`:
		*f = AppResourceDatabaseDatabasePermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_CONNECT_AND_CREATE"`, v)
	}
}

// Values returns all possible values for AppResourceDatabaseDatabasePermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppResourceDatabaseDatabasePermission) Values() []AppResourceDatabaseDatabasePermission {
	return []AppResourceDatabaseDatabasePermission{
		AppResourceDatabaseDatabasePermissionCanConnectAndCreate,
	}
}

// Type always returns AppResourceDatabaseDatabasePermission to satisfy [pflag.Value] interface
func (f *AppResourceDatabaseDatabasePermission) Type() string {
	return "AppResourceDatabaseDatabasePermission"
}

type AppResourceExperiment struct {
	ExperimentId string `json:"experiment_id"`

	Permission AppResourceExperimentExperimentPermission `json:"permission"`
}

type AppResourceExperimentExperimentPermission string

const AppResourceExperimentExperimentPermissionCanEdit AppResourceExperimentExperimentPermission = `CAN_EDIT`

const AppResourceExperimentExperimentPermissionCanManage AppResourceExperimentExperimentPermission = `CAN_MANAGE`

const AppResourceExperimentExperimentPermissionCanRead AppResourceExperimentExperimentPermission = `CAN_READ`

// String representation for [fmt.Print]
func (f *AppResourceExperimentExperimentPermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppResourceExperimentExperimentPermission) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`:
		*f = AppResourceExperimentExperimentPermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_READ"`, v)
	}
}

// Values returns all possible values for AppResourceExperimentExperimentPermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppResourceExperimentExperimentPermission) Values() []AppResourceExperimentExperimentPermission {
	return []AppResourceExperimentExperimentPermission{
		AppResourceExperimentExperimentPermissionCanEdit,
		AppResourceExperimentExperimentPermissionCanManage,
		AppResourceExperimentExperimentPermissionCanRead,
	}
}

// Type always returns AppResourceExperimentExperimentPermission to satisfy [pflag.Value] interface
func (f *AppResourceExperimentExperimentPermission) Type() string {
	return "AppResourceExperimentExperimentPermission"
}

type AppResourceGenieSpace struct {
	Name string `json:"name"`

	Permission AppResourceGenieSpaceGenieSpacePermission `json:"permission"`

	SpaceId string `json:"space_id"`
}

type AppResourceGenieSpaceGenieSpacePermission string

const AppResourceGenieSpaceGenieSpacePermissionCanEdit AppResourceGenieSpaceGenieSpacePermission = `CAN_EDIT`

const AppResourceGenieSpaceGenieSpacePermissionCanManage AppResourceGenieSpaceGenieSpacePermission = `CAN_MANAGE`

const AppResourceGenieSpaceGenieSpacePermissionCanRun AppResourceGenieSpaceGenieSpacePermission = `CAN_RUN`

const AppResourceGenieSpaceGenieSpacePermissionCanView AppResourceGenieSpaceGenieSpacePermission = `CAN_VIEW`

// String representation for [fmt.Print]
func (f *AppResourceGenieSpaceGenieSpacePermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppResourceGenieSpaceGenieSpacePermission) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_RUN`, `CAN_VIEW`:
		*f = AppResourceGenieSpaceGenieSpacePermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_RUN", "CAN_VIEW"`, v)
	}
}

// Values returns all possible values for AppResourceGenieSpaceGenieSpacePermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppResourceGenieSpaceGenieSpacePermission) Values() []AppResourceGenieSpaceGenieSpacePermission {
	return []AppResourceGenieSpaceGenieSpacePermission{
		AppResourceGenieSpaceGenieSpacePermissionCanEdit,
		AppResourceGenieSpaceGenieSpacePermissionCanManage,
		AppResourceGenieSpaceGenieSpacePermissionCanRun,
		AppResourceGenieSpaceGenieSpacePermissionCanView,
	}
}

// Type always returns AppResourceGenieSpaceGenieSpacePermission to satisfy [pflag.Value] interface
func (f *AppResourceGenieSpaceGenieSpacePermission) Type() string {
	return "AppResourceGenieSpaceGenieSpacePermission"
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

// Values returns all possible values for AppResourceJobJobPermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppResourceJobJobPermission) Values() []AppResourceJobJobPermission {
	return []AppResourceJobJobPermission{
		AppResourceJobJobPermissionCanManage,
		AppResourceJobJobPermissionCanManageRun,
		AppResourceJobJobPermissionCanView,
		AppResourceJobJobPermissionIsOwner,
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

// Values returns all possible values for AppResourceSecretSecretPermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppResourceSecretSecretPermission) Values() []AppResourceSecretSecretPermission {
	return []AppResourceSecretSecretPermission{
		AppResourceSecretSecretPermissionManage,
		AppResourceSecretSecretPermissionRead,
		AppResourceSecretSecretPermissionWrite,
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

// Values returns all possible values for AppResourceServingEndpointServingEndpointPermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppResourceServingEndpointServingEndpointPermission) Values() []AppResourceServingEndpointServingEndpointPermission {
	return []AppResourceServingEndpointServingEndpointPermission{
		AppResourceServingEndpointServingEndpointPermissionCanManage,
		AppResourceServingEndpointServingEndpointPermissionCanQuery,
		AppResourceServingEndpointServingEndpointPermissionCanView,
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

// Values returns all possible values for AppResourceSqlWarehouseSqlWarehousePermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppResourceSqlWarehouseSqlWarehousePermission) Values() []AppResourceSqlWarehouseSqlWarehousePermission {
	return []AppResourceSqlWarehouseSqlWarehousePermission{
		AppResourceSqlWarehouseSqlWarehousePermissionCanManage,
		AppResourceSqlWarehouseSqlWarehousePermissionCanUse,
		AppResourceSqlWarehouseSqlWarehousePermissionIsOwner,
	}
}

// Type always returns AppResourceSqlWarehouseSqlWarehousePermission to satisfy [pflag.Value] interface
func (f *AppResourceSqlWarehouseSqlWarehousePermission) Type() string {
	return "AppResourceSqlWarehouseSqlWarehousePermission"
}

type AppResourceUcSecurable struct {
	Permission AppResourceUcSecurableUcSecurablePermission `json:"permission"`

	SecurableFullName string `json:"securable_full_name"`
	// The securable kind from Unity Catalog. See
	// https://docs.databricks.com/api/workspace/tables/get#securable_kind_manifest-securable_kind.
	SecurableKind string `json:"securable_kind,omitempty"`

	SecurableType AppResourceUcSecurableUcSecurableType `json:"securable_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AppResourceUcSecurable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppResourceUcSecurable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppResourceUcSecurableUcSecurablePermission string

const AppResourceUcSecurableUcSecurablePermissionExecute AppResourceUcSecurableUcSecurablePermission = `EXECUTE`

const AppResourceUcSecurableUcSecurablePermissionModify AppResourceUcSecurableUcSecurablePermission = `MODIFY`

const AppResourceUcSecurableUcSecurablePermissionReadVolume AppResourceUcSecurableUcSecurablePermission = `READ_VOLUME`

const AppResourceUcSecurableUcSecurablePermissionSelect AppResourceUcSecurableUcSecurablePermission = `SELECT`

const AppResourceUcSecurableUcSecurablePermissionUseConnection AppResourceUcSecurableUcSecurablePermission = `USE_CONNECTION`

const AppResourceUcSecurableUcSecurablePermissionWriteVolume AppResourceUcSecurableUcSecurablePermission = `WRITE_VOLUME`

// String representation for [fmt.Print]
func (f *AppResourceUcSecurableUcSecurablePermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppResourceUcSecurableUcSecurablePermission) Set(v string) error {
	switch v {
	case `EXECUTE`, `MODIFY`, `READ_VOLUME`, `SELECT`, `USE_CONNECTION`, `WRITE_VOLUME`:
		*f = AppResourceUcSecurableUcSecurablePermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXECUTE", "MODIFY", "READ_VOLUME", "SELECT", "USE_CONNECTION", "WRITE_VOLUME"`, v)
	}
}

// Values returns all possible values for AppResourceUcSecurableUcSecurablePermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppResourceUcSecurableUcSecurablePermission) Values() []AppResourceUcSecurableUcSecurablePermission {
	return []AppResourceUcSecurableUcSecurablePermission{
		AppResourceUcSecurableUcSecurablePermissionExecute,
		AppResourceUcSecurableUcSecurablePermissionModify,
		AppResourceUcSecurableUcSecurablePermissionReadVolume,
		AppResourceUcSecurableUcSecurablePermissionSelect,
		AppResourceUcSecurableUcSecurablePermissionUseConnection,
		AppResourceUcSecurableUcSecurablePermissionWriteVolume,
	}
}

// Type always returns AppResourceUcSecurableUcSecurablePermission to satisfy [pflag.Value] interface
func (f *AppResourceUcSecurableUcSecurablePermission) Type() string {
	return "AppResourceUcSecurableUcSecurablePermission"
}

type AppResourceUcSecurableUcSecurableType string

const AppResourceUcSecurableUcSecurableTypeConnection AppResourceUcSecurableUcSecurableType = `CONNECTION`

const AppResourceUcSecurableUcSecurableTypeFunction AppResourceUcSecurableUcSecurableType = `FUNCTION`

const AppResourceUcSecurableUcSecurableTypeTable AppResourceUcSecurableUcSecurableType = `TABLE`

const AppResourceUcSecurableUcSecurableTypeVolume AppResourceUcSecurableUcSecurableType = `VOLUME`

// String representation for [fmt.Print]
func (f *AppResourceUcSecurableUcSecurableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppResourceUcSecurableUcSecurableType) Set(v string) error {
	switch v {
	case `CONNECTION`, `FUNCTION`, `TABLE`, `VOLUME`:
		*f = AppResourceUcSecurableUcSecurableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONNECTION", "FUNCTION", "TABLE", "VOLUME"`, v)
	}
}

// Values returns all possible values for AppResourceUcSecurableUcSecurableType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppResourceUcSecurableUcSecurableType) Values() []AppResourceUcSecurableUcSecurableType {
	return []AppResourceUcSecurableUcSecurableType{
		AppResourceUcSecurableUcSecurableTypeConnection,
		AppResourceUcSecurableUcSecurableTypeFunction,
		AppResourceUcSecurableUcSecurableTypeTable,
		AppResourceUcSecurableUcSecurableTypeVolume,
	}
}

// Type always returns AppResourceUcSecurableUcSecurableType to satisfy [pflag.Value] interface
func (f *AppResourceUcSecurableUcSecurableType) Type() string {
	return "AppResourceUcSecurableUcSecurableType"
}

type AppUpdate struct {
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`

	ComputeSize ComputeSize `json:"compute_size,omitempty"`

	Description string `json:"description,omitempty"`

	GitRepository *GitRepository `json:"git_repository,omitempty"`

	Resources []AppResource `json:"resources,omitempty"`

	Status *AppUpdateUpdateStatus `json:"status,omitempty"`

	UsagePolicyId string `json:"usage_policy_id,omitempty"`

	UserApiScopes []string `json:"user_api_scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AppUpdate) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppUpdate) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppUpdateUpdateStatus struct {
	Message string `json:"message,omitempty"`

	State AppUpdateUpdateStatusUpdateState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AppUpdateUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppUpdateUpdateStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppUpdateUpdateStatusUpdateState string

const AppUpdateUpdateStatusUpdateStateFailed AppUpdateUpdateStatusUpdateState = `FAILED`

const AppUpdateUpdateStatusUpdateStateInProgress AppUpdateUpdateStatusUpdateState = `IN_PROGRESS`

const AppUpdateUpdateStatusUpdateStateNotUpdated AppUpdateUpdateStatusUpdateState = `NOT_UPDATED`

const AppUpdateUpdateStatusUpdateStateSucceeded AppUpdateUpdateStatusUpdateState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *AppUpdateUpdateStatusUpdateState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppUpdateUpdateStatusUpdateState) Set(v string) error {
	switch v {
	case `FAILED`, `IN_PROGRESS`, `NOT_UPDATED`, `SUCCEEDED`:
		*f = AppUpdateUpdateStatusUpdateState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "IN_PROGRESS", "NOT_UPDATED", "SUCCEEDED"`, v)
	}
}

// Values returns all possible values for AppUpdateUpdateStatusUpdateState.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppUpdateUpdateStatusUpdateState) Values() []AppUpdateUpdateStatusUpdateState {
	return []AppUpdateUpdateStatusUpdateState{
		AppUpdateUpdateStatusUpdateStateFailed,
		AppUpdateUpdateStatusUpdateStateInProgress,
		AppUpdateUpdateStatusUpdateStateNotUpdated,
		AppUpdateUpdateStatusUpdateStateSucceeded,
	}
}

// Type always returns AppUpdateUpdateStatusUpdateState to satisfy [pflag.Value] interface
func (f *AppUpdateUpdateStatusUpdateState) Type() string {
	return "AppUpdateUpdateStatusUpdateState"
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

// Values returns all possible values for ApplicationState.
//
// There is no guarantee on the order of the values in the slice.
func (f *ApplicationState) Values() []ApplicationState {
	return []ApplicationState{
		ApplicationStateCrashed,
		ApplicationStateDeploying,
		ApplicationStateRunning,
		ApplicationStateUnavailable,
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

type AsyncUpdateAppRequest struct {
	App *App `json:"app,omitempty"`

	AppName string `json:"-" url:"-"`
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
	UpdateMask string `json:"update_mask"`
}

type ComputeSize string

const ComputeSizeLarge ComputeSize = `LARGE`

const ComputeSizeMedium ComputeSize = `MEDIUM`

// String representation for [fmt.Print]
func (f *ComputeSize) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ComputeSize) Set(v string) error {
	switch v {
	case `LARGE`, `MEDIUM`:
		*f = ComputeSize(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LARGE", "MEDIUM"`, v)
	}
}

// Values returns all possible values for ComputeSize.
//
// There is no guarantee on the order of the values in the slice.
func (f *ComputeSize) Values() []ComputeSize {
	return []ComputeSize{
		ComputeSizeLarge,
		ComputeSizeMedium,
	}
}

// Type always returns ComputeSize to satisfy [pflag.Value] interface
func (f *ComputeSize) Type() string {
	return "ComputeSize"
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

// Values returns all possible values for ComputeState.
//
// There is no guarantee on the order of the values in the slice.
func (f *ComputeState) Values() []ComputeState {
	return []ComputeState{
		ComputeStateActive,
		ComputeStateDeleting,
		ComputeStateError,
		ComputeStateStarting,
		ComputeStateStopped,
		ComputeStateStopping,
		ComputeStateUpdating,
	}
}

// Type always returns ComputeState to satisfy [pflag.Value] interface
func (f *ComputeState) Type() string {
	return "ComputeState"
}

type ComputeStatus struct {
	// The number of compute instances currently serving requests for this
	// application. An instance is considered active if it is reachable and
	// ready to handle requests.
	ActiveInstances int `json:"active_instances,omitempty"`
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

type CreateAppDeploymentRequest struct {
	// The app deployment configuration.
	AppDeployment AppDeployment `json:"app_deployment"`
	// The name of the app.
	AppName string `json:"-" url:"-"`
}

type CreateAppRequest struct {
	App App `json:"app"`
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

type CreateCustomTemplateRequest struct {
	Template CustomTemplate `json:"template"`
}

type CreateSpaceRequest struct {
	Space Space `json:"space"`
}

type CustomTemplate struct {
	Creator string `json:"creator,omitempty"`
	// The description of the template.
	Description string `json:"description,omitempty"`
	// The Git provider of the template.
	GitProvider string `json:"git_provider"`
	// The Git repository URL that the template resides in.
	GitRepo string `json:"git_repo"`
	// The manifest of the template. It defines fields and default values when
	// installing the template.
	Manifest AppManifest `json:"manifest"`
	// The name of the template. It must contain only alphanumeric characters,
	// hyphens, underscores, and whitespaces. It must be unique within the
	// workspace.
	Name string `json:"name"`
	// The path to the template within the Git repository.
	Path string `json:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomTemplate) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomTemplate) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Databricks Error that is returned by all Databricks APIs.
type DatabricksServiceExceptionWithDetailsProto struct {
	Details []json.RawMessage `json:"details,omitempty"`

	ErrorCode ErrorCode `json:"error_code,omitempty"`

	Message string `json:"message,omitempty"`

	StackTrace string `json:"stack_trace,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabricksServiceExceptionWithDetailsProto) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksServiceExceptionWithDetailsProto) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteAppRequest struct {
	// The name of the app.
	Name string `json:"-" url:"-"`
}

type DeleteCustomTemplateRequest struct {
	// The name of the custom template.
	Name string `json:"-" url:"-"`
}

type DeleteSpaceRequest struct {
	// The name of the app space.
	Name string `json:"-" url:"-"`
}

type EnvVar struct {
	// The name of the environment variable.
	Name string `json:"name,omitempty"`
	// The value for the environment variable.
	Value string `json:"value,omitempty"`
	// The name of an external Databricks resource that contains the value, such
	// as a secret or a database table.
	ValueFrom string `json:"value_from,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnvVar) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnvVar) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Error codes returned by Databricks APIs to indicate specific failure
// conditions.
type ErrorCode string

const ErrorCodeAborted ErrorCode = `ABORTED`

const ErrorCodeAlreadyExists ErrorCode = `ALREADY_EXISTS`

const ErrorCodeBadRequest ErrorCode = `BAD_REQUEST`

const ErrorCodeCancelled ErrorCode = `CANCELLED`

const ErrorCodeCatalogAlreadyExists ErrorCode = `CATALOG_ALREADY_EXISTS`

const ErrorCodeCatalogDoesNotExist ErrorCode = `CATALOG_DOES_NOT_EXIST`

const ErrorCodeCatalogNotEmpty ErrorCode = `CATALOG_NOT_EMPTY`

const ErrorCodeCouldNotAcquireLock ErrorCode = `COULD_NOT_ACQUIRE_LOCK`

const ErrorCodeCustomerUnauthorized ErrorCode = `CUSTOMER_UNAUTHORIZED`

const ErrorCodeDacAlreadyExists ErrorCode = `DAC_ALREADY_EXISTS`

const ErrorCodeDacDoesNotExist ErrorCode = `DAC_DOES_NOT_EXIST`

const ErrorCodeDataLoss ErrorCode = `DATA_LOSS`

const ErrorCodeDeadlineExceeded ErrorCode = `DEADLINE_EXCEEDED`

const ErrorCodeDeploymentTimeout ErrorCode = `DEPLOYMENT_TIMEOUT`

const ErrorCodeDirectoryNotEmpty ErrorCode = `DIRECTORY_NOT_EMPTY`

const ErrorCodeDirectoryProtected ErrorCode = `DIRECTORY_PROTECTED`

const ErrorCodeDryRunFailed ErrorCode = `DRY_RUN_FAILED`

const ErrorCodeEndpointNotFound ErrorCode = `ENDPOINT_NOT_FOUND`

const ErrorCodeExternalLocationAlreadyExists ErrorCode = `EXTERNAL_LOCATION_ALREADY_EXISTS`

const ErrorCodeExternalLocationDoesNotExist ErrorCode = `EXTERNAL_LOCATION_DOES_NOT_EXIST`

const ErrorCodeFeatureDisabled ErrorCode = `FEATURE_DISABLED`

const ErrorCodeGitConflict ErrorCode = `GIT_CONFLICT`

const ErrorCodeGitRemoteError ErrorCode = `GIT_REMOTE_ERROR`

const ErrorCodeGitSensitiveTokenDetected ErrorCode = `GIT_SENSITIVE_TOKEN_DETECTED`

const ErrorCodeGitUnknownRef ErrorCode = `GIT_UNKNOWN_REF`

const ErrorCodeGitUrlNotOnAllowList ErrorCode = `GIT_URL_NOT_ON_ALLOW_LIST`

const ErrorCodeInsecurePartnerResponse ErrorCode = `INSECURE_PARTNER_RESPONSE`

const ErrorCodeInternalError ErrorCode = `INTERNAL_ERROR`

const ErrorCodeInvalidParameterValue ErrorCode = `INVALID_PARAMETER_VALUE`

const ErrorCodeInvalidState ErrorCode = `INVALID_STATE`

const ErrorCodeInvalidStateTransition ErrorCode = `INVALID_STATE_TRANSITION`

const ErrorCodeIoError ErrorCode = `IO_ERROR`

const ErrorCodeIpynbFileInRepo ErrorCode = `IPYNB_FILE_IN_REPO`

const ErrorCodeMalformedPartnerResponse ErrorCode = `MALFORMED_PARTNER_RESPONSE`

const ErrorCodeMalformedRequest ErrorCode = `MALFORMED_REQUEST`

const ErrorCodeManagedResourceGroupDoesNotExist ErrorCode = `MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST`

const ErrorCodeMaxBlockSizeExceeded ErrorCode = `MAX_BLOCK_SIZE_EXCEEDED`

const ErrorCodeMaxChildNodeSizeExceeded ErrorCode = `MAX_CHILD_NODE_SIZE_EXCEEDED`

const ErrorCodeMaxListSizeExceeded ErrorCode = `MAX_LIST_SIZE_EXCEEDED`

const ErrorCodeMaxNotebookSizeExceeded ErrorCode = `MAX_NOTEBOOK_SIZE_EXCEEDED`

const ErrorCodeMaxReadSizeExceeded ErrorCode = `MAX_READ_SIZE_EXCEEDED`

const ErrorCodeMetastoreAlreadyExists ErrorCode = `METASTORE_ALREADY_EXISTS`

const ErrorCodeMetastoreDoesNotExist ErrorCode = `METASTORE_DOES_NOT_EXIST`

const ErrorCodeMetastoreNotEmpty ErrorCode = `METASTORE_NOT_EMPTY`

const ErrorCodeNotFound ErrorCode = `NOT_FOUND`

const ErrorCodeNotImplemented ErrorCode = `NOT_IMPLEMENTED`

const ErrorCodePartialDelete ErrorCode = `PARTIAL_DELETE`

const ErrorCodePermissionDenied ErrorCode = `PERMISSION_DENIED`

const ErrorCodePermissionNotPropagated ErrorCode = `PERMISSION_NOT_PROPAGATED`

const ErrorCodePrincipalDoesNotExist ErrorCode = `PRINCIPAL_DOES_NOT_EXIST`

const ErrorCodeProjectsOperationTimeout ErrorCode = `PROJECTS_OPERATION_TIMEOUT`

const ErrorCodeProviderAlreadyExists ErrorCode = `PROVIDER_ALREADY_EXISTS`

const ErrorCodeProviderDoesNotExist ErrorCode = `PROVIDER_DOES_NOT_EXIST`

const ErrorCodeProviderShareNotAccessible ErrorCode = `PROVIDER_SHARE_NOT_ACCESSIBLE`

const ErrorCodeQuotaExceeded ErrorCode = `QUOTA_EXCEEDED`

const ErrorCodeRecipientAlreadyExists ErrorCode = `RECIPIENT_ALREADY_EXISTS`

const ErrorCodeRecipientDoesNotExist ErrorCode = `RECIPIENT_DOES_NOT_EXIST`

const ErrorCodeRequestLimitExceeded ErrorCode = `REQUEST_LIMIT_EXCEEDED`

const ErrorCodeResourceAlreadyExists ErrorCode = `RESOURCE_ALREADY_EXISTS`

const ErrorCodeResourceConflict ErrorCode = `RESOURCE_CONFLICT`

const ErrorCodeResourceDoesNotExist ErrorCode = `RESOURCE_DOES_NOT_EXIST`

const ErrorCodeResourceExhausted ErrorCode = `RESOURCE_EXHAUSTED`

const ErrorCodeResourceLimitExceeded ErrorCode = `RESOURCE_LIMIT_EXCEEDED`

const ErrorCodeSchemaAlreadyExists ErrorCode = `SCHEMA_ALREADY_EXISTS`

const ErrorCodeSchemaDoesNotExist ErrorCode = `SCHEMA_DOES_NOT_EXIST`

const ErrorCodeSchemaNotEmpty ErrorCode = `SCHEMA_NOT_EMPTY`

const ErrorCodeSearchQueryTooLong ErrorCode = `SEARCH_QUERY_TOO_LONG`

const ErrorCodeSearchQueryTooShort ErrorCode = `SEARCH_QUERY_TOO_SHORT`

const ErrorCodeServiceUnderMaintenance ErrorCode = `SERVICE_UNDER_MAINTENANCE`

const ErrorCodeShareAlreadyExists ErrorCode = `SHARE_ALREADY_EXISTS`

const ErrorCodeShareDoesNotExist ErrorCode = `SHARE_DOES_NOT_EXIST`

const ErrorCodeStorageCredentialAlreadyExists ErrorCode = `STORAGE_CREDENTIAL_ALREADY_EXISTS`

const ErrorCodeStorageCredentialDoesNotExist ErrorCode = `STORAGE_CREDENTIAL_DOES_NOT_EXIST`

const ErrorCodeTableAlreadyExists ErrorCode = `TABLE_ALREADY_EXISTS`

const ErrorCodeTableDoesNotExist ErrorCode = `TABLE_DOES_NOT_EXIST`

const ErrorCodeTemporarilyUnavailable ErrorCode = `TEMPORARILY_UNAVAILABLE`

const ErrorCodeUnauthenticated ErrorCode = `UNAUTHENTICATED`

const ErrorCodeUnavailable ErrorCode = `UNAVAILABLE`

const ErrorCodeUnknown ErrorCode = `UNKNOWN`

const ErrorCodeUnparseableHttpError ErrorCode = `UNPARSEABLE_HTTP_ERROR`

const ErrorCodeWorkspaceTemporarilyUnavailable ErrorCode = `WORKSPACE_TEMPORARILY_UNAVAILABLE`

// String representation for [fmt.Print]
func (f *ErrorCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ErrorCode) Set(v string) error {
	switch v {
	case `ABORTED`, `ALREADY_EXISTS`, `BAD_REQUEST`, `CANCELLED`, `CATALOG_ALREADY_EXISTS`, `CATALOG_DOES_NOT_EXIST`, `CATALOG_NOT_EMPTY`, `COULD_NOT_ACQUIRE_LOCK`, `CUSTOMER_UNAUTHORIZED`, `DAC_ALREADY_EXISTS`, `DAC_DOES_NOT_EXIST`, `DATA_LOSS`, `DEADLINE_EXCEEDED`, `DEPLOYMENT_TIMEOUT`, `DIRECTORY_NOT_EMPTY`, `DIRECTORY_PROTECTED`, `DRY_RUN_FAILED`, `ENDPOINT_NOT_FOUND`, `EXTERNAL_LOCATION_ALREADY_EXISTS`, `EXTERNAL_LOCATION_DOES_NOT_EXIST`, `FEATURE_DISABLED`, `GIT_CONFLICT`, `GIT_REMOTE_ERROR`, `GIT_SENSITIVE_TOKEN_DETECTED`, `GIT_UNKNOWN_REF`, `GIT_URL_NOT_ON_ALLOW_LIST`, `INSECURE_PARTNER_RESPONSE`, `INTERNAL_ERROR`, `INVALID_PARAMETER_VALUE`, `INVALID_STATE`, `INVALID_STATE_TRANSITION`, `IO_ERROR`, `IPYNB_FILE_IN_REPO`, `MALFORMED_PARTNER_RESPONSE`, `MALFORMED_REQUEST`, `MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST`, `MAX_BLOCK_SIZE_EXCEEDED`, `MAX_CHILD_NODE_SIZE_EXCEEDED`, `MAX_LIST_SIZE_EXCEEDED`, `MAX_NOTEBOOK_SIZE_EXCEEDED`, `MAX_READ_SIZE_EXCEEDED`, `METASTORE_ALREADY_EXISTS`, `METASTORE_DOES_NOT_EXIST`, `METASTORE_NOT_EMPTY`, `NOT_FOUND`, `NOT_IMPLEMENTED`, `PARTIAL_DELETE`, `PERMISSION_DENIED`, `PERMISSION_NOT_PROPAGATED`, `PRINCIPAL_DOES_NOT_EXIST`, `PROJECTS_OPERATION_TIMEOUT`, `PROVIDER_ALREADY_EXISTS`, `PROVIDER_DOES_NOT_EXIST`, `PROVIDER_SHARE_NOT_ACCESSIBLE`, `QUOTA_EXCEEDED`, `RECIPIENT_ALREADY_EXISTS`, `RECIPIENT_DOES_NOT_EXIST`, `REQUEST_LIMIT_EXCEEDED`, `RESOURCE_ALREADY_EXISTS`, `RESOURCE_CONFLICT`, `RESOURCE_DOES_NOT_EXIST`, `RESOURCE_EXHAUSTED`, `RESOURCE_LIMIT_EXCEEDED`, `SCHEMA_ALREADY_EXISTS`, `SCHEMA_DOES_NOT_EXIST`, `SCHEMA_NOT_EMPTY`, `SEARCH_QUERY_TOO_LONG`, `SEARCH_QUERY_TOO_SHORT`, `SERVICE_UNDER_MAINTENANCE`, `SHARE_ALREADY_EXISTS`, `SHARE_DOES_NOT_EXIST`, `STORAGE_CREDENTIAL_ALREADY_EXISTS`, `STORAGE_CREDENTIAL_DOES_NOT_EXIST`, `TABLE_ALREADY_EXISTS`, `TABLE_DOES_NOT_EXIST`, `TEMPORARILY_UNAVAILABLE`, `UNAUTHENTICATED`, `UNAVAILABLE`, `UNKNOWN`, `UNPARSEABLE_HTTP_ERROR`, `WORKSPACE_TEMPORARILY_UNAVAILABLE`:
		*f = ErrorCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABORTED", "ALREADY_EXISTS", "BAD_REQUEST", "CANCELLED", "CATALOG_ALREADY_EXISTS", "CATALOG_DOES_NOT_EXIST", "CATALOG_NOT_EMPTY", "COULD_NOT_ACQUIRE_LOCK", "CUSTOMER_UNAUTHORIZED", "DAC_ALREADY_EXISTS", "DAC_DOES_NOT_EXIST", "DATA_LOSS", "DEADLINE_EXCEEDED", "DEPLOYMENT_TIMEOUT", "DIRECTORY_NOT_EMPTY", "DIRECTORY_PROTECTED", "DRY_RUN_FAILED", "ENDPOINT_NOT_FOUND", "EXTERNAL_LOCATION_ALREADY_EXISTS", "EXTERNAL_LOCATION_DOES_NOT_EXIST", "FEATURE_DISABLED", "GIT_CONFLICT", "GIT_REMOTE_ERROR", "GIT_SENSITIVE_TOKEN_DETECTED", "GIT_UNKNOWN_REF", "GIT_URL_NOT_ON_ALLOW_LIST", "INSECURE_PARTNER_RESPONSE", "INTERNAL_ERROR", "INVALID_PARAMETER_VALUE", "INVALID_STATE", "INVALID_STATE_TRANSITION", "IO_ERROR", "IPYNB_FILE_IN_REPO", "MALFORMED_PARTNER_RESPONSE", "MALFORMED_REQUEST", "MANAGED_RESOURCE_GROUP_DOES_NOT_EXIST", "MAX_BLOCK_SIZE_EXCEEDED", "MAX_CHILD_NODE_SIZE_EXCEEDED", "MAX_LIST_SIZE_EXCEEDED", "MAX_NOTEBOOK_SIZE_EXCEEDED", "MAX_READ_SIZE_EXCEEDED", "METASTORE_ALREADY_EXISTS", "METASTORE_DOES_NOT_EXIST", "METASTORE_NOT_EMPTY", "NOT_FOUND", "NOT_IMPLEMENTED", "PARTIAL_DELETE", "PERMISSION_DENIED", "PERMISSION_NOT_PROPAGATED", "PRINCIPAL_DOES_NOT_EXIST", "PROJECTS_OPERATION_TIMEOUT", "PROVIDER_ALREADY_EXISTS", "PROVIDER_DOES_NOT_EXIST", "PROVIDER_SHARE_NOT_ACCESSIBLE", "QUOTA_EXCEEDED", "RECIPIENT_ALREADY_EXISTS", "RECIPIENT_DOES_NOT_EXIST", "REQUEST_LIMIT_EXCEEDED", "RESOURCE_ALREADY_EXISTS", "RESOURCE_CONFLICT", "RESOURCE_DOES_NOT_EXIST", "RESOURCE_EXHAUSTED", "RESOURCE_LIMIT_EXCEEDED", "SCHEMA_ALREADY_EXISTS", "SCHEMA_DOES_NOT_EXIST", "SCHEMA_NOT_EMPTY", "SEARCH_QUERY_TOO_LONG", "SEARCH_QUERY_TOO_SHORT", "SERVICE_UNDER_MAINTENANCE", "SHARE_ALREADY_EXISTS", "SHARE_DOES_NOT_EXIST", "STORAGE_CREDENTIAL_ALREADY_EXISTS", "STORAGE_CREDENTIAL_DOES_NOT_EXIST", "TABLE_ALREADY_EXISTS", "TABLE_DOES_NOT_EXIST", "TEMPORARILY_UNAVAILABLE", "UNAUTHENTICATED", "UNAVAILABLE", "UNKNOWN", "UNPARSEABLE_HTTP_ERROR", "WORKSPACE_TEMPORARILY_UNAVAILABLE"`, v)
	}
}

// Values returns all possible values for ErrorCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		ErrorCodeAborted,
		ErrorCodeAlreadyExists,
		ErrorCodeBadRequest,
		ErrorCodeCancelled,
		ErrorCodeCatalogAlreadyExists,
		ErrorCodeCatalogDoesNotExist,
		ErrorCodeCatalogNotEmpty,
		ErrorCodeCouldNotAcquireLock,
		ErrorCodeCustomerUnauthorized,
		ErrorCodeDacAlreadyExists,
		ErrorCodeDacDoesNotExist,
		ErrorCodeDataLoss,
		ErrorCodeDeadlineExceeded,
		ErrorCodeDeploymentTimeout,
		ErrorCodeDirectoryNotEmpty,
		ErrorCodeDirectoryProtected,
		ErrorCodeDryRunFailed,
		ErrorCodeEndpointNotFound,
		ErrorCodeExternalLocationAlreadyExists,
		ErrorCodeExternalLocationDoesNotExist,
		ErrorCodeFeatureDisabled,
		ErrorCodeGitConflict,
		ErrorCodeGitRemoteError,
		ErrorCodeGitSensitiveTokenDetected,
		ErrorCodeGitUnknownRef,
		ErrorCodeGitUrlNotOnAllowList,
		ErrorCodeInsecurePartnerResponse,
		ErrorCodeInternalError,
		ErrorCodeInvalidParameterValue,
		ErrorCodeInvalidState,
		ErrorCodeInvalidStateTransition,
		ErrorCodeIoError,
		ErrorCodeIpynbFileInRepo,
		ErrorCodeMalformedPartnerResponse,
		ErrorCodeMalformedRequest,
		ErrorCodeManagedResourceGroupDoesNotExist,
		ErrorCodeMaxBlockSizeExceeded,
		ErrorCodeMaxChildNodeSizeExceeded,
		ErrorCodeMaxListSizeExceeded,
		ErrorCodeMaxNotebookSizeExceeded,
		ErrorCodeMaxReadSizeExceeded,
		ErrorCodeMetastoreAlreadyExists,
		ErrorCodeMetastoreDoesNotExist,
		ErrorCodeMetastoreNotEmpty,
		ErrorCodeNotFound,
		ErrorCodeNotImplemented,
		ErrorCodePartialDelete,
		ErrorCodePermissionDenied,
		ErrorCodePermissionNotPropagated,
		ErrorCodePrincipalDoesNotExist,
		ErrorCodeProjectsOperationTimeout,
		ErrorCodeProviderAlreadyExists,
		ErrorCodeProviderDoesNotExist,
		ErrorCodeProviderShareNotAccessible,
		ErrorCodeQuotaExceeded,
		ErrorCodeRecipientAlreadyExists,
		ErrorCodeRecipientDoesNotExist,
		ErrorCodeRequestLimitExceeded,
		ErrorCodeResourceAlreadyExists,
		ErrorCodeResourceConflict,
		ErrorCodeResourceDoesNotExist,
		ErrorCodeResourceExhausted,
		ErrorCodeResourceLimitExceeded,
		ErrorCodeSchemaAlreadyExists,
		ErrorCodeSchemaDoesNotExist,
		ErrorCodeSchemaNotEmpty,
		ErrorCodeSearchQueryTooLong,
		ErrorCodeSearchQueryTooShort,
		ErrorCodeServiceUnderMaintenance,
		ErrorCodeShareAlreadyExists,
		ErrorCodeShareDoesNotExist,
		ErrorCodeStorageCredentialAlreadyExists,
		ErrorCodeStorageCredentialDoesNotExist,
		ErrorCodeTableAlreadyExists,
		ErrorCodeTableDoesNotExist,
		ErrorCodeTemporarilyUnavailable,
		ErrorCodeUnauthenticated,
		ErrorCodeUnavailable,
		ErrorCodeUnknown,
		ErrorCodeUnparseableHttpError,
		ErrorCodeWorkspaceTemporarilyUnavailable,
	}
}

// Type always returns ErrorCode to satisfy [pflag.Value] interface
func (f *ErrorCode) Type() string {
	return "ErrorCode"
}

type GetAppDeploymentRequest struct {
	// The name of the app.
	AppName string `json:"-" url:"-"`
	// The unique id of the deployment.
	DeploymentId string `json:"-" url:"-"`
}

type GetAppPermissionLevelsRequest struct {
	// The app for which to get or manage permissions.
	AppName string `json:"-" url:"-"`
}

type GetAppPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []AppPermissionsDescription `json:"permission_levels,omitempty"`
}

type GetAppPermissionsRequest struct {
	// The app for which to get or manage permissions.
	AppName string `json:"-" url:"-"`
}

type GetAppRequest struct {
	// The name of the app.
	Name string `json:"-" url:"-"`
}

type GetAppUpdateRequest struct {
	// The name of the app.
	AppName string `json:"-" url:"-"`
}

type GetCustomTemplateRequest struct {
	// The name of the custom template.
	Name string `json:"-" url:"-"`
}

type GetOperationRequest struct {
	// The name of the operation resource.
	Name string `json:"-" url:"-"`
}

type GetSpaceRequest struct {
	// The name of the app space.
	Name string `json:"-" url:"-"`
}

// Git repository configuration specifying the location of the repository.
type GitRepository struct {
	// Git provider. Case insensitive. Supported values: gitHub,
	// gitHubEnterprise, bitbucketCloud, bitbucketServer, azureDevOpsServices,
	// gitLab, gitLabEnterpriseEdition, awsCodeCommit.
	Provider string `json:"provider"`
	// URL of the Git repository.
	Url string `json:"url"`
}

// Complete git source specification including repository location and
// reference.
type GitSource struct {
	// Git branch to checkout.
	Branch string `json:"branch,omitempty"`
	// Git commit SHA to checkout.
	Commit string `json:"commit,omitempty"`
	// Git repository configuration. Populated from the app's git_repository
	// configuration.
	GitRepository *GitRepository `json:"git_repository,omitempty"`
	// The resolved commit SHA that was actually used for the deployment. This
	// is populated by the system after resolving the reference (branch, tag, or
	// commit). If commit is specified directly, this will match commit. If a
	// branch or tag is specified, this contains the commit SHA that the branch
	// or tag pointed to at deployment time.
	ResolvedCommit string `json:"resolved_commit,omitempty"`
	// Relative path to the app source code within the Git repository. If not
	// specified, the root of the repository is used.
	SourceCodePath string `json:"source_code_path,omitempty"`
	// Git tag to checkout.
	Tag string `json:"tag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GitSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GitSource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

type ListAppsRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Filter apps by app space name. When specified, only apps belonging to
	// this space are returned.
	Space string `json:"-" url:"space,omitempty"`

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

type ListCustomTemplatesRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of custom templates. Requests
	// first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCustomTemplatesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCustomTemplatesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListCustomTemplatesResponse struct {
	// Pagination token to request the next page of custom templates.
	NextPageToken string `json:"next_page_token,omitempty"`

	Templates []CustomTemplate `json:"templates,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCustomTemplatesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCustomTemplatesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSpacesRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of app spaces. Requests first
	// page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSpacesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSpacesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSpacesResponse struct {
	// Pagination token to request the next page of app spaces.
	NextPageToken string `json:"next_page_token,omitempty"`

	Spaces []Space `json:"spaces,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSpacesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSpacesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// If the value is `false`, it means the operation is still in progress. If
	// `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `json:"done,omitempty"`
	// The error result of the operation in case of failure or cancellation.
	Error *DatabricksServiceExceptionWithDetailsProto `json:"error,omitempty"`
	// Service-specific metadata associated with the operation. It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.
	Metadata json.RawMessage `json:"metadata,omitempty"`
	// The server-assigned name, which is only unique within the same service
	// that originally returns it. If you use the default HTTP mapping, the
	// `name` should be a resource name ending with `operations/{unique_id}`.
	Name string `json:"name,omitempty"`
	// The normal, successful response of the operation.
	Response json.RawMessage `json:"response,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Operation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Operation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Space struct {
	// The creation time of the app space. Formatted timestamp in ISO 6801.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The email of the user that created the app space.
	Creator string `json:"creator,omitempty"`
	// The description of the app space.
	Description string `json:"description,omitempty"`
	// The effective usage policy ID used by apps in the space.
	EffectiveUsagePolicyId string `json:"effective_usage_policy_id,omitempty"`
	// The effective api scopes granted to the user access token.
	EffectiveUserApiScopes []string `json:"effective_user_api_scopes,omitempty"`
	// The unique identifier of the app space.
	Id string `json:"id,omitempty"`
	// The name of the app space. The name must contain only lowercase
	// alphanumeric characters and hyphens. It must be unique within the
	// workspace.
	Name string `json:"name"`
	// Resources for the app space. Resources configured at the space level are
	// available to all apps in the space.
	Resources []AppResource `json:"resources,omitempty"`
	// The service principal client ID for the app space.
	ServicePrincipalClientId string `json:"service_principal_client_id,omitempty"`
	// The service principal ID for the app space.
	ServicePrincipalId int64 `json:"service_principal_id,omitempty"`
	// The service principal name for the app space.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The status of the app space.
	Status *SpaceStatus `json:"status,omitempty"`
	// The update time of the app space. Formatted timestamp in ISO 6801.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// The email of the user that last updated the app space.
	Updater string `json:"updater,omitempty"`
	// The usage policy ID for managing cost at the space level.
	UsagePolicyId string `json:"usage_policy_id,omitempty"`
	// OAuth scopes for apps in the space.
	UserApiScopes []string `json:"user_api_scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Space) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Space) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SpaceStatus struct {
	// Message providing context about the current state.
	Message string `json:"message,omitempty"`
	// The state of the app space.
	State SpaceStatusSpaceState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SpaceStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SpaceStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SpaceStatusSpaceState string

const SpaceStatusSpaceStateSpaceActive SpaceStatusSpaceState = `SPACE_ACTIVE`

const SpaceStatusSpaceStateSpaceCreating SpaceStatusSpaceState = `SPACE_CREATING`

const SpaceStatusSpaceStateSpaceDeleted SpaceStatusSpaceState = `SPACE_DELETED`

const SpaceStatusSpaceStateSpaceDeleting SpaceStatusSpaceState = `SPACE_DELETING`

const SpaceStatusSpaceStateSpaceError SpaceStatusSpaceState = `SPACE_ERROR`

const SpaceStatusSpaceStateSpaceUpdating SpaceStatusSpaceState = `SPACE_UPDATING`

// String representation for [fmt.Print]
func (f *SpaceStatusSpaceState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SpaceStatusSpaceState) Set(v string) error {
	switch v {
	case `SPACE_ACTIVE`, `SPACE_CREATING`, `SPACE_DELETED`, `SPACE_DELETING`, `SPACE_ERROR`, `SPACE_UPDATING`:
		*f = SpaceStatusSpaceState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SPACE_ACTIVE", "SPACE_CREATING", "SPACE_DELETED", "SPACE_DELETING", "SPACE_ERROR", "SPACE_UPDATING"`, v)
	}
}

// Values returns all possible values for SpaceStatusSpaceState.
//
// There is no guarantee on the order of the values in the slice.
func (f *SpaceStatusSpaceState) Values() []SpaceStatusSpaceState {
	return []SpaceStatusSpaceState{
		SpaceStatusSpaceStateSpaceActive,
		SpaceStatusSpaceStateSpaceCreating,
		SpaceStatusSpaceStateSpaceDeleted,
		SpaceStatusSpaceStateSpaceDeleting,
		SpaceStatusSpaceStateSpaceError,
		SpaceStatusSpaceStateSpaceUpdating,
	}
}

// Type always returns SpaceStatusSpaceState to satisfy [pflag.Value] interface
func (f *SpaceStatusSpaceState) Type() string {
	return "SpaceStatusSpaceState"
}

// Tracks app space update information.
type SpaceUpdate struct {
	Description string `json:"description,omitempty"`

	Resources []AppResource `json:"resources,omitempty"`

	Status *SpaceUpdateStatus `json:"status,omitempty"`

	UsagePolicyId string `json:"usage_policy_id,omitempty"`

	UserApiScopes []string `json:"user_api_scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SpaceUpdate) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SpaceUpdate) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SpaceUpdateState string

const SpaceUpdateStateFailed SpaceUpdateState = `FAILED`

const SpaceUpdateStateInProgress SpaceUpdateState = `IN_PROGRESS`

const SpaceUpdateStateNotUpdated SpaceUpdateState = `NOT_UPDATED`

const SpaceUpdateStateSucceeded SpaceUpdateState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *SpaceUpdateState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SpaceUpdateState) Set(v string) error {
	switch v {
	case `FAILED`, `IN_PROGRESS`, `NOT_UPDATED`, `SUCCEEDED`:
		*f = SpaceUpdateState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "IN_PROGRESS", "NOT_UPDATED", "SUCCEEDED"`, v)
	}
}

// Values returns all possible values for SpaceUpdateState.
//
// There is no guarantee on the order of the values in the slice.
func (f *SpaceUpdateState) Values() []SpaceUpdateState {
	return []SpaceUpdateState{
		SpaceUpdateStateFailed,
		SpaceUpdateStateInProgress,
		SpaceUpdateStateNotUpdated,
		SpaceUpdateStateSucceeded,
	}
}

// Type always returns SpaceUpdateState to satisfy [pflag.Value] interface
func (f *SpaceUpdateState) Type() string {
	return "SpaceUpdateState"
}

// Status of an app space update operation
type SpaceUpdateStatus struct {
	Message string `json:"message,omitempty"`

	State SpaceUpdateState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SpaceUpdateStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SpaceUpdateStatus) MarshalJSON() ([]byte, error) {
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

type UpdateAppRequest struct {
	App App `json:"app"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name string `json:"-" url:"-"`
}

type UpdateCustomTemplateRequest struct {
	// The name of the template. It must contain only alphanumeric characters,
	// hyphens, underscores, and whitespaces. It must be unique within the
	// workspace.
	Name string `json:"-" url:"-"`

	Template CustomTemplate `json:"template"`
}

type UpdateSpaceRequest struct {
	// The name of the app space. The name must contain only lowercase
	// alphanumeric characters and hyphens. It must be unique within the
	// workspace.
	Name string `json:"-" url:"-"`

	Space Space `json:"space"`
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
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}
