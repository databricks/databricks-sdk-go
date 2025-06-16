// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps

import (
	"encoding/json"
	"fmt"
)

type App struct {
	// The active deployment of the app. A deployment is considered active when
	// it has been deployed to the app compute.
	// Wire name: 'active_deployment'
	ActiveDeployment *AppDeployment `json:"active_deployment,omitempty"`

	// Wire name: 'app_status'
	AppStatus *ApplicationStatus `json:"app_status,omitempty"`

	// Wire name: 'budget_policy_id'
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`

	// Wire name: 'compute_status'
	ComputeStatus *ComputeStatus `json:"compute_status,omitempty"`
	// The creation time of the app. Formatted timestamp in ISO 6801.
	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"`
	// The email of the user that created the app.
	// Wire name: 'creator'
	Creator string `json:"creator,omitempty"`
	// The default workspace file system path of the source code from which app
	// deployment are created. This field tracks the workspace source code path
	// of the last active deployment.
	// Wire name: 'default_source_code_path'
	DefaultSourceCodePath string `json:"default_source_code_path,omitempty"`
	// The description of the app.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`

	// Wire name: 'effective_budget_policy_id'
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`
	// The effective api scopes granted to the user access token.
	// Wire name: 'effective_user_api_scopes'
	EffectiveUserApiScopes []string `json:"effective_user_api_scopes,omitempty"`
	// The unique identifier of the app.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	// Wire name: 'name'
	Name string `json:"name"`

	// Wire name: 'oauth2_app_client_id'
	Oauth2AppClientId string `json:"oauth2_app_client_id,omitempty"`

	// Wire name: 'oauth2_app_integration_id'
	Oauth2AppIntegrationId string `json:"oauth2_app_integration_id,omitempty"`
	// The pending deployment of the app. A deployment is considered pending
	// when it is being prepared for deployment to the app compute.
	// Wire name: 'pending_deployment'
	PendingDeployment *AppDeployment `json:"pending_deployment,omitempty"`
	// Resources for the app.
	// Wire name: 'resources'
	Resources []AppResource `json:"resources,omitempty"`

	// Wire name: 'service_principal_client_id'
	ServicePrincipalClientId string `json:"service_principal_client_id,omitempty"`

	// Wire name: 'service_principal_id'
	ServicePrincipalId int64 `json:"service_principal_id,omitempty"`

	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The update time of the app. Formatted timestamp in ISO 6801.
	// Wire name: 'update_time'
	UpdateTime string `json:"update_time,omitempty"`
	// The email of the user that last updated the app.
	// Wire name: 'updater'
	Updater string `json:"updater,omitempty"`
	// The URL of the app once it is deployed.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`

	// Wire name: 'user_api_scopes'
	UserApiScopes []string `json:"user_api_scopes,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *App) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st App) MarshalJSON() ([]byte, error) {
	pb, err := appToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AppAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel AppPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AppAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AppAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := appAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AppAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []AppPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AppAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AppAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := appAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AppDeployment struct {
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"`
	// The email of the user creates the deployment.
	// Wire name: 'creator'
	Creator string `json:"creator,omitempty"`
	// The deployment artifacts for an app.
	// Wire name: 'deployment_artifacts'
	DeploymentArtifacts *AppDeploymentArtifacts `json:"deployment_artifacts,omitempty"`
	// The unique id of the deployment.
	// Wire name: 'deployment_id'
	DeploymentId string `json:"deployment_id,omitempty"`
	// The mode of which the deployment will manage the source code.
	// Wire name: 'mode'
	Mode AppDeploymentMode `json:"mode,omitempty"`
	// The workspace file system path of the source code used to create the app
	// deployment. This is different from
	// `deployment_artifacts.source_code_path`, which is the path used by the
	// deployed app. The former refers to the original source code location of
	// the app in the workspace during deployment creation, whereas the latter
	// provides a system generated stable snapshotted source code path used by
	// the deployment.
	// Wire name: 'source_code_path'
	SourceCodePath string `json:"source_code_path,omitempty"`
	// Status and status message of the deployment
	// Wire name: 'status'
	Status *AppDeploymentStatus `json:"status,omitempty"`
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	// Wire name: 'update_time'
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AppDeployment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appDeploymentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appDeploymentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AppDeployment) MarshalJSON() ([]byte, error) {
	pb, err := appDeploymentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AppDeploymentArtifacts struct {
	// The snapshotted workspace file system path of the source code loaded by
	// the deployed app.
	// Wire name: 'source_code_path'
	SourceCodePath string `json:"source_code_path,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AppDeploymentArtifacts) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appDeploymentArtifactsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appDeploymentArtifactsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AppDeploymentArtifacts) MarshalJSON() ([]byte, error) {
	pb, err := appDeploymentArtifactsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'message'
	Message string `json:"message,omitempty"`
	// State of the deployment.
	// Wire name: 'state'
	State AppDeploymentState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AppDeploymentStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appDeploymentStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appDeploymentStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AppDeploymentStatus) MarshalJSON() ([]byte, error) {
	pb, err := appDeploymentStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AppPermission struct {

	// Wire name: 'inherited'
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel AppPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AppPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AppPermission) MarshalJSON() ([]byte, error) {
	pb, err := appPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'access_control_list'
	AccessControlList []AppAccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AppPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AppPermissions) MarshalJSON() ([]byte, error) {
	pb, err := appPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AppPermissionsDescription struct {

	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel AppPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AppPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AppPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := appPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AppPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []AppAccessControlRequest `json:"access_control_list,omitempty"`
	// The app for which to get or manage permissions.
	AppName string `json:"-" tf:"-"`
}

func (st *AppPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AppPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := appPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AppResource struct {
	// Description of the App Resource.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`

	// Wire name: 'job'
	Job *AppResourceJob `json:"job,omitempty"`
	// Name of the App Resource.
	// Wire name: 'name'
	Name string `json:"name"`

	// Wire name: 'secret'
	Secret *AppResourceSecret `json:"secret,omitempty"`

	// Wire name: 'serving_endpoint'
	ServingEndpoint *AppResourceServingEndpoint `json:"serving_endpoint,omitempty"`

	// Wire name: 'sql_warehouse'
	SqlWarehouse *AppResourceSqlWarehouse `json:"sql_warehouse,omitempty"`

	// Wire name: 'uc_securable'
	UcSecurable *AppResourceUcSecurable `json:"uc_securable,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AppResource) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appResourcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appResourceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AppResource) MarshalJSON() ([]byte, error) {
	pb, err := appResourceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AppResourceJob struct {
	// Id of the job to grant permission on.
	// Wire name: 'id'
	Id string `json:"id"`
	// Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE",
	// "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW".
	// Wire name: 'permission'
	Permission AppResourceJobJobPermission `json:"permission"`
}

func (st *AppResourceJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appResourceJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appResourceJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AppResourceJob) MarshalJSON() ([]byte, error) {
	pb, err := appResourceJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'key'
	Key string `json:"key"`
	// Permission to grant on the secret scope. For secrets, only one permission
	// is allowed. Permission must be one of: "READ", "WRITE", "MANAGE".
	// Wire name: 'permission'
	Permission AppResourceSecretSecretPermission `json:"permission"`
	// Scope of the secret to grant permission on.
	// Wire name: 'scope'
	Scope string `json:"scope"`
}

func (st *AppResourceSecret) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appResourceSecretPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appResourceSecretFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AppResourceSecret) MarshalJSON() ([]byte, error) {
	pb, err := appResourceSecretToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'name'
	Name string `json:"name"`
	// Permission to grant on the serving endpoint. Supported permissions are:
	// "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW".
	// Wire name: 'permission'
	Permission AppResourceServingEndpointServingEndpointPermission `json:"permission"`
}

func (st *AppResourceServingEndpoint) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appResourceServingEndpointPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appResourceServingEndpointFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AppResourceServingEndpoint) MarshalJSON() ([]byte, error) {
	pb, err := appResourceServingEndpointToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'id'
	Id string `json:"id"`
	// Permission to grant on the SQL warehouse. Supported permissions are:
	// "CAN_MANAGE", "CAN_USE", "IS_OWNER".
	// Wire name: 'permission'
	Permission AppResourceSqlWarehouseSqlWarehousePermission `json:"permission"`
}

func (st *AppResourceSqlWarehouse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appResourceSqlWarehousePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appResourceSqlWarehouseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AppResourceSqlWarehouse) MarshalJSON() ([]byte, error) {
	pb, err := appResourceSqlWarehouseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'permission'
	Permission AppResourceUcSecurableUcSecurablePermission `json:"permission"`

	// Wire name: 'securable_full_name'
	SecurableFullName string `json:"securable_full_name"`

	// Wire name: 'securable_type'
	SecurableType AppResourceUcSecurableUcSecurableType `json:"securable_type"`
}

func (st *AppResourceUcSecurable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appResourceUcSecurablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := appResourceUcSecurableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AppResourceUcSecurable) MarshalJSON() ([]byte, error) {
	pb, err := appResourceUcSecurableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AppResourceUcSecurableUcSecurablePermission string

const AppResourceUcSecurableUcSecurablePermissionReadVolume AppResourceUcSecurableUcSecurablePermission = `READ_VOLUME`

const AppResourceUcSecurableUcSecurablePermissionWriteVolume AppResourceUcSecurableUcSecurablePermission = `WRITE_VOLUME`

// String representation for [fmt.Print]
func (f *AppResourceUcSecurableUcSecurablePermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppResourceUcSecurableUcSecurablePermission) Set(v string) error {
	switch v {
	case `READ_VOLUME`, `WRITE_VOLUME`:
		*f = AppResourceUcSecurableUcSecurablePermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "READ_VOLUME", "WRITE_VOLUME"`, v)
	}
}

// Values returns all possible values for AppResourceUcSecurableUcSecurablePermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppResourceUcSecurableUcSecurablePermission) Values() []AppResourceUcSecurableUcSecurablePermission {
	return []AppResourceUcSecurableUcSecurablePermission{
		AppResourceUcSecurableUcSecurablePermissionReadVolume,
		AppResourceUcSecurableUcSecurablePermissionWriteVolume,
	}
}

// Type always returns AppResourceUcSecurableUcSecurablePermission to satisfy [pflag.Value] interface
func (f *AppResourceUcSecurableUcSecurablePermission) Type() string {
	return "AppResourceUcSecurableUcSecurablePermission"
}

type AppResourceUcSecurableUcSecurableType string

const AppResourceUcSecurableUcSecurableTypeVolume AppResourceUcSecurableUcSecurableType = `VOLUME`

// String representation for [fmt.Print]
func (f *AppResourceUcSecurableUcSecurableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppResourceUcSecurableUcSecurableType) Set(v string) error {
	switch v {
	case `VOLUME`:
		*f = AppResourceUcSecurableUcSecurableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "VOLUME"`, v)
	}
}

// Values returns all possible values for AppResourceUcSecurableUcSecurableType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AppResourceUcSecurableUcSecurableType) Values() []AppResourceUcSecurableUcSecurableType {
	return []AppResourceUcSecurableUcSecurableType{
		AppResourceUcSecurableUcSecurableTypeVolume,
	}
}

// Type always returns AppResourceUcSecurableUcSecurableType to satisfy [pflag.Value] interface
func (f *AppResourceUcSecurableUcSecurableType) Type() string {
	return "AppResourceUcSecurableUcSecurableType"
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
	// Wire name: 'message'
	Message string `json:"message,omitempty"`
	// State of the application.
	// Wire name: 'state'
	State ApplicationState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ApplicationStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &applicationStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := applicationStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ApplicationStatus) MarshalJSON() ([]byte, error) {
	pb, err := applicationStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Compute status message
	// Wire name: 'message'
	Message string `json:"message,omitempty"`
	// State of the app compute.
	// Wire name: 'state'
	State ComputeState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ComputeStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computeStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := computeStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ComputeStatus) MarshalJSON() ([]byte, error) {
	pb, err := computeStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Create an app deployment
type CreateAppDeploymentRequest struct {

	// Wire name: 'app_deployment'
	AppDeployment AppDeployment `json:"app_deployment"`
	// The name of the app.
	AppName string `json:"-" tf:"-"`
}

func (st *CreateAppDeploymentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createAppDeploymentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createAppDeploymentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateAppDeploymentRequest) MarshalJSON() ([]byte, error) {
	pb, err := createAppDeploymentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Create an app
type CreateAppRequest struct {

	// Wire name: 'app'
	App App `json:"app"`
	// If true, the app will not be started after creation.
	NoCompute bool `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateAppRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createAppRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createAppRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateAppRequest) MarshalJSON() ([]byte, error) {
	pb, err := createAppRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete an app
type DeleteAppRequest struct {
	// The name of the app.
	Name string `json:"-" tf:"-"`
}

func (st *DeleteAppRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAppRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAppRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAppRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteAppRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get an app deployment
type GetAppDeploymentRequest struct {
	// The name of the app.
	AppName string `json:"-" tf:"-"`
	// The unique id of the deployment.
	DeploymentId string `json:"-" tf:"-"`
}

func (st *GetAppDeploymentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAppDeploymentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAppDeploymentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAppDeploymentRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAppDeploymentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get app permission levels
type GetAppPermissionLevelsRequest struct {
	// The app for which to get or manage permissions.
	AppName string `json:"-" tf:"-"`
}

func (st *GetAppPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAppPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAppPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAppPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAppPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetAppPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []AppPermissionsDescription `json:"permission_levels,omitempty"`
}

func (st *GetAppPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAppPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAppPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAppPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getAppPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get app permissions
type GetAppPermissionsRequest struct {
	// The app for which to get or manage permissions.
	AppName string `json:"-" tf:"-"`
}

func (st *GetAppPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAppPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAppPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAppPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAppPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get an app
type GetAppRequest struct {
	// The name of the app.
	Name string `json:"-" tf:"-"`
}

func (st *GetAppRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAppRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAppRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAppRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAppRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List app deployments
type ListAppDeploymentsRequest struct {
	// The name of the app.
	AppName string `json:"-" tf:"-"`
	// Upper bound for items returned.
	PageSize int `json:"-" tf:"-"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListAppDeploymentsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAppDeploymentsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAppDeploymentsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAppDeploymentsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listAppDeploymentsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListAppDeploymentsResponse struct {
	// Deployment history of the app.
	// Wire name: 'app_deployments'
	AppDeployments []AppDeployment `json:"app_deployments,omitempty"`
	// Pagination token to request the next page of apps.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListAppDeploymentsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAppDeploymentsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAppDeploymentsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAppDeploymentsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listAppDeploymentsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List apps
type ListAppsRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" tf:"-"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListAppsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAppsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAppsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAppsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listAppsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListAppsResponse struct {

	// Wire name: 'apps'
	Apps []App `json:"apps,omitempty"`
	// Pagination token to request the next page of apps.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListAppsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAppsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAppsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAppsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listAppsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StartAppRequest struct {
	// The name of the app.
	Name string `json:"-" tf:"-"`
}

func (st *StartAppRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &startAppRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := startAppRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StartAppRequest) MarshalJSON() ([]byte, error) {
	pb, err := startAppRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StopAppRequest struct {
	// The name of the app.
	Name string `json:"-" tf:"-"`
}

func (st *StopAppRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &stopAppRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := stopAppRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StopAppRequest) MarshalJSON() ([]byte, error) {
	pb, err := stopAppRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Update an app
type UpdateAppRequest struct {

	// Wire name: 'app'
	App App `json:"app"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name string `json:"-" tf:"-"`
}

func (st *UpdateAppRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateAppRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateAppRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateAppRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateAppRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}
