// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/apps/appspb"
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
	CreateTime string `json:"create_time,omitempty"` //legacy
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
	UpdateTime string `json:"update_time,omitempty"` //legacy
	// The email of the user that last updated the app.
	// Wire name: 'updater'
	Updater string `json:"updater,omitempty"`
	// The URL of the app once it is deployed.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`

	// Wire name: 'user_api_scopes'
	UserApiScopes   []string `json:"user_api_scopes,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st App) MarshalJSON() ([]byte, error) {
	pb, err := AppToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *App) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppToPb(st *App) (*appspb.AppPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppPb{}
	activeDeploymentPb, err := AppDeploymentToPb(st.ActiveDeployment)
	if err != nil {
		return nil, err
	}
	if activeDeploymentPb != nil {
		pb.ActiveDeployment = activeDeploymentPb
	}
	appStatusPb, err := ApplicationStatusToPb(st.AppStatus)
	if err != nil {
		return nil, err
	}
	if appStatusPb != nil {
		pb.AppStatus = appStatusPb
	}
	pb.BudgetPolicyId = st.BudgetPolicyId
	computeStatusPb, err := ComputeStatusToPb(st.ComputeStatus)
	if err != nil {
		return nil, err
	}
	if computeStatusPb != nil {
		pb.ComputeStatus = computeStatusPb
	}
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
	pendingDeploymentPb, err := AppDeploymentToPb(st.PendingDeployment)
	if err != nil {
		return nil, err
	}
	if pendingDeploymentPb != nil {
		pb.PendingDeployment = pendingDeploymentPb
	}

	var resourcesPb []appspb.AppResourcePb
	for _, item := range st.Resources {
		itemPb, err := AppResourceToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resourcesPb = append(resourcesPb, *itemPb)
		}
	}
	pb.Resources = resourcesPb
	pb.ServicePrincipalClientId = st.ServicePrincipalClientId
	pb.ServicePrincipalId = st.ServicePrincipalId
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UpdateTime = st.UpdateTime
	pb.Updater = st.Updater
	pb.Url = st.Url
	pb.UserApiScopes = st.UserApiScopes

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AppFromPb(pb *appspb.AppPb) (*App, error) {
	if pb == nil {
		return nil, nil
	}
	st := &App{}
	activeDeploymentField, err := AppDeploymentFromPb(pb.ActiveDeployment)
	if err != nil {
		return nil, err
	}
	if activeDeploymentField != nil {
		st.ActiveDeployment = activeDeploymentField
	}
	appStatusField, err := ApplicationStatusFromPb(pb.AppStatus)
	if err != nil {
		return nil, err
	}
	if appStatusField != nil {
		st.AppStatus = appStatusField
	}
	st.BudgetPolicyId = pb.BudgetPolicyId
	computeStatusField, err := ComputeStatusFromPb(pb.ComputeStatus)
	if err != nil {
		return nil, err
	}
	if computeStatusField != nil {
		st.ComputeStatus = computeStatusField
	}
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
	pendingDeploymentField, err := AppDeploymentFromPb(pb.PendingDeployment)
	if err != nil {
		return nil, err
	}
	if pendingDeploymentField != nil {
		st.PendingDeployment = pendingDeploymentField
	}

	var resourcesField []AppResource
	for _, itemPb := range pb.Resources {
		item, err := AppResourceFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resourcesField = append(resourcesField, *item)
		}
	}
	st.Resources = resourcesField
	st.ServicePrincipalClientId = pb.ServicePrincipalClientId
	st.ServicePrincipalId = pb.ServicePrincipalId
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UpdateTime = pb.UpdateTime
	st.Updater = pb.Updater
	st.Url = pb.Url
	st.UserApiScopes = pb.UserApiScopes

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type AppAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel AppPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName        string   `json:"user_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AppAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := AppAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppAccessControlRequestToPb(st *AppAccessControlRequest) (*appspb.AppAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := AppPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AppAccessControlRequestFromPb(pb *appspb.AppAccessControlRequestPb) (*AppAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppAccessControlRequest{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := AppPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	UserName        string   `json:"user_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AppAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := AppAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppAccessControlResponseToPb(st *AppAccessControlResponse) (*appspb.AppAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppAccessControlResponsePb{}

	var allPermissionsPb []appspb.AppPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := AppPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AppAccessControlResponseFromPb(pb *appspb.AppAccessControlResponsePb) (*AppAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppAccessControlResponse{}

	var allPermissionsField []AppPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := AppPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type AppDeployment struct {
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"` //legacy
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
	UpdateTime      string   `json:"update_time,omitempty"` //legacy
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AppDeployment) MarshalJSON() ([]byte, error) {
	pb, err := AppDeploymentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppDeployment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppDeploymentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppDeploymentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppDeploymentToPb(st *AppDeployment) (*appspb.AppDeploymentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppDeploymentPb{}
	pb.CreateTime = st.CreateTime
	pb.Creator = st.Creator
	deploymentArtifactsPb, err := AppDeploymentArtifactsToPb(st.DeploymentArtifacts)
	if err != nil {
		return nil, err
	}
	if deploymentArtifactsPb != nil {
		pb.DeploymentArtifacts = deploymentArtifactsPb
	}
	pb.DeploymentId = st.DeploymentId
	modePb, err := AppDeploymentModeToPb(&st.Mode)
	if err != nil {
		return nil, err
	}
	if modePb != nil {
		pb.Mode = *modePb
	}
	pb.SourceCodePath = st.SourceCodePath
	statusPb, err := AppDeploymentStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}
	pb.UpdateTime = st.UpdateTime

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AppDeploymentFromPb(pb *appspb.AppDeploymentPb) (*AppDeployment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppDeployment{}
	st.CreateTime = pb.CreateTime
	st.Creator = pb.Creator
	deploymentArtifactsField, err := AppDeploymentArtifactsFromPb(pb.DeploymentArtifacts)
	if err != nil {
		return nil, err
	}
	if deploymentArtifactsField != nil {
		st.DeploymentArtifacts = deploymentArtifactsField
	}
	st.DeploymentId = pb.DeploymentId
	modeField, err := AppDeploymentModeFromPb(&pb.Mode)
	if err != nil {
		return nil, err
	}
	if modeField != nil {
		st.Mode = *modeField
	}
	st.SourceCodePath = pb.SourceCodePath
	statusField, err := AppDeploymentStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}
	st.UpdateTime = pb.UpdateTime

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type AppDeploymentArtifacts struct {
	// The snapshotted workspace file system path of the source code loaded by
	// the deployed app.
	// Wire name: 'source_code_path'
	SourceCodePath  string   `json:"source_code_path,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AppDeploymentArtifacts) MarshalJSON() ([]byte, error) {
	pb, err := AppDeploymentArtifactsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppDeploymentArtifacts) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppDeploymentArtifactsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppDeploymentArtifactsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppDeploymentArtifactsToPb(st *AppDeploymentArtifacts) (*appspb.AppDeploymentArtifactsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppDeploymentArtifactsPb{}
	pb.SourceCodePath = st.SourceCodePath

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AppDeploymentArtifactsFromPb(pb *appspb.AppDeploymentArtifactsPb) (*AppDeploymentArtifacts, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppDeploymentArtifacts{}
	st.SourceCodePath = pb.SourceCodePath

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func AppDeploymentModeToPb(st *AppDeploymentMode) (*appspb.AppDeploymentModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appspb.AppDeploymentModePb(*st)
	return &pb, nil
}

func AppDeploymentModeFromPb(pb *appspb.AppDeploymentModePb) (*AppDeploymentMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppDeploymentMode(*pb)
	return &st, nil
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

func AppDeploymentStateToPb(st *AppDeploymentState) (*appspb.AppDeploymentStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appspb.AppDeploymentStatePb(*st)
	return &pb, nil
}

func AppDeploymentStateFromPb(pb *appspb.AppDeploymentStatePb) (*AppDeploymentState, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppDeploymentState(*pb)
	return &st, nil
}

type AppDeploymentStatus struct {
	// Message corresponding with the deployment state.
	// Wire name: 'message'
	Message string `json:"message,omitempty"`
	// State of the deployment.
	// Wire name: 'state'
	State           AppDeploymentState `json:"state,omitempty"`
	ForceSendFields []string           `json:"-" tf:"-"`
}

func (st AppDeploymentStatus) MarshalJSON() ([]byte, error) {
	pb, err := AppDeploymentStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppDeploymentStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppDeploymentStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppDeploymentStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppDeploymentStatusToPb(st *AppDeploymentStatus) (*appspb.AppDeploymentStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppDeploymentStatusPb{}
	pb.Message = st.Message
	statePb, err := AppDeploymentStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AppDeploymentStatusFromPb(pb *appspb.AppDeploymentStatusPb) (*AppDeploymentStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppDeploymentStatus{}
	st.Message = pb.Message
	stateField, err := AppDeploymentStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type AppPermission struct {

	// Wire name: 'inherited'
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel AppPermissionLevel `json:"permission_level,omitempty"`
	ForceSendFields []string           `json:"-" tf:"-"`
}

func (st AppPermission) MarshalJSON() ([]byte, error) {
	pb, err := AppPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppPermissionToPb(st *AppPermission) (*appspb.AppPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	permissionLevelPb, err := AppPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AppPermissionFromPb(pb *appspb.AppPermissionPb) (*AppPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	permissionLevelField, err := AppPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func AppPermissionLevelToPb(st *AppPermissionLevel) (*appspb.AppPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appspb.AppPermissionLevelPb(*st)
	return &pb, nil
}

func AppPermissionLevelFromPb(pb *appspb.AppPermissionLevelPb) (*AppPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppPermissionLevel(*pb)
	return &st, nil
}

type AppPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []AppAccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType      string   `json:"object_type,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AppPermissions) MarshalJSON() ([]byte, error) {
	pb, err := AppPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppPermissionsToPb(st *AppPermissions) (*appspb.AppPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppPermissionsPb{}

	var accessControlListPb []appspb.AppAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := AppAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AppPermissionsFromPb(pb *appspb.AppPermissionsPb) (*AppPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppPermissions{}

	var accessControlListField []AppAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := AppAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type AppPermissionsDescription struct {

	// Wire name: 'description'
	Description string `json:"description,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel AppPermissionLevel `json:"permission_level,omitempty"`
	ForceSendFields []string           `json:"-" tf:"-"`
}

func (st AppPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := AppPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppPermissionsDescriptionToPb(st *AppPermissionsDescription) (*appspb.AppPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppPermissionsDescriptionPb{}
	pb.Description = st.Description
	permissionLevelPb, err := AppPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AppPermissionsDescriptionFromPb(pb *appspb.AppPermissionsDescriptionPb) (*AppPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppPermissionsDescription{}
	st.Description = pb.Description
	permissionLevelField, err := AppPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type AppPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []AppAccessControlRequest `json:"access_control_list,omitempty"`
	// The app for which to get or manage permissions.
	AppName string `json:"-" tf:"-"`
}

func (st AppPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := AppPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppPermissionsRequestToPb(st *AppPermissionsRequest) (*appspb.AppPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppPermissionsRequestPb{}

	var accessControlListPb []appspb.AppAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := AppAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.AppName = st.AppName

	return pb, nil
}

func AppPermissionsRequestFromPb(pb *appspb.AppPermissionsRequestPb) (*AppPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppPermissionsRequest{}

	var accessControlListField []AppAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := AppAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.AppName = pb.AppName

	return st, nil
}

type AppResource struct {

	// Wire name: 'database'
	Database *AppResourceDatabase `json:"database,omitempty"`
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
	UcSecurable     *AppResourceUcSecurable `json:"uc_securable,omitempty"`
	ForceSendFields []string                `json:"-" tf:"-"`
}

func (st AppResource) MarshalJSON() ([]byte, error) {
	pb, err := AppResourceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppResource) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppResourcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppResourceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppResourceToPb(st *AppResource) (*appspb.AppResourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppResourcePb{}
	databasePb, err := AppResourceDatabaseToPb(st.Database)
	if err != nil {
		return nil, err
	}
	if databasePb != nil {
		pb.Database = databasePb
	}
	pb.Description = st.Description
	jobPb, err := AppResourceJobToPb(st.Job)
	if err != nil {
		return nil, err
	}
	if jobPb != nil {
		pb.Job = jobPb
	}
	pb.Name = st.Name
	secretPb, err := AppResourceSecretToPb(st.Secret)
	if err != nil {
		return nil, err
	}
	if secretPb != nil {
		pb.Secret = secretPb
	}
	servingEndpointPb, err := AppResourceServingEndpointToPb(st.ServingEndpoint)
	if err != nil {
		return nil, err
	}
	if servingEndpointPb != nil {
		pb.ServingEndpoint = servingEndpointPb
	}
	sqlWarehousePb, err := AppResourceSqlWarehouseToPb(st.SqlWarehouse)
	if err != nil {
		return nil, err
	}
	if sqlWarehousePb != nil {
		pb.SqlWarehouse = sqlWarehousePb
	}
	ucSecurablePb, err := AppResourceUcSecurableToPb(st.UcSecurable)
	if err != nil {
		return nil, err
	}
	if ucSecurablePb != nil {
		pb.UcSecurable = ucSecurablePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AppResourceFromPb(pb *appspb.AppResourcePb) (*AppResource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResource{}
	databaseField, err := AppResourceDatabaseFromPb(pb.Database)
	if err != nil {
		return nil, err
	}
	if databaseField != nil {
		st.Database = databaseField
	}
	st.Description = pb.Description
	jobField, err := AppResourceJobFromPb(pb.Job)
	if err != nil {
		return nil, err
	}
	if jobField != nil {
		st.Job = jobField
	}
	st.Name = pb.Name
	secretField, err := AppResourceSecretFromPb(pb.Secret)
	if err != nil {
		return nil, err
	}
	if secretField != nil {
		st.Secret = secretField
	}
	servingEndpointField, err := AppResourceServingEndpointFromPb(pb.ServingEndpoint)
	if err != nil {
		return nil, err
	}
	if servingEndpointField != nil {
		st.ServingEndpoint = servingEndpointField
	}
	sqlWarehouseField, err := AppResourceSqlWarehouseFromPb(pb.SqlWarehouse)
	if err != nil {
		return nil, err
	}
	if sqlWarehouseField != nil {
		st.SqlWarehouse = sqlWarehouseField
	}
	ucSecurableField, err := AppResourceUcSecurableFromPb(pb.UcSecurable)
	if err != nil {
		return nil, err
	}
	if ucSecurableField != nil {
		st.UcSecurable = ucSecurableField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type AppResourceDatabase struct {

	// Wire name: 'database_name'
	DatabaseName string `json:"database_name"`

	// Wire name: 'instance_name'
	InstanceName string `json:"instance_name"`

	// Wire name: 'permission'
	Permission AppResourceDatabaseDatabasePermission `json:"permission"`
}

func (st AppResourceDatabase) MarshalJSON() ([]byte, error) {
	pb, err := AppResourceDatabaseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppResourceDatabase) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppResourceDatabasePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppResourceDatabaseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppResourceDatabaseToPb(st *AppResourceDatabase) (*appspb.AppResourceDatabasePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppResourceDatabasePb{}
	pb.DatabaseName = st.DatabaseName
	pb.InstanceName = st.InstanceName
	permissionPb, err := AppResourceDatabaseDatabasePermissionToPb(&st.Permission)
	if err != nil {
		return nil, err
	}
	if permissionPb != nil {
		pb.Permission = *permissionPb
	}

	return pb, nil
}

func AppResourceDatabaseFromPb(pb *appspb.AppResourceDatabasePb) (*AppResourceDatabase, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResourceDatabase{}
	st.DatabaseName = pb.DatabaseName
	st.InstanceName = pb.InstanceName
	permissionField, err := AppResourceDatabaseDatabasePermissionFromPb(&pb.Permission)
	if err != nil {
		return nil, err
	}
	if permissionField != nil {
		st.Permission = *permissionField
	}

	return st, nil
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

func AppResourceDatabaseDatabasePermissionToPb(st *AppResourceDatabaseDatabasePermission) (*appspb.AppResourceDatabaseDatabasePermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appspb.AppResourceDatabaseDatabasePermissionPb(*st)
	return &pb, nil
}

func AppResourceDatabaseDatabasePermissionFromPb(pb *appspb.AppResourceDatabaseDatabasePermissionPb) (*AppResourceDatabaseDatabasePermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppResourceDatabaseDatabasePermission(*pb)
	return &st, nil
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

func (st AppResourceJob) MarshalJSON() ([]byte, error) {
	pb, err := AppResourceJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppResourceJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppResourceJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppResourceJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppResourceJobToPb(st *AppResourceJob) (*appspb.AppResourceJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppResourceJobPb{}
	pb.Id = st.Id
	permissionPb, err := AppResourceJobJobPermissionToPb(&st.Permission)
	if err != nil {
		return nil, err
	}
	if permissionPb != nil {
		pb.Permission = *permissionPb
	}

	return pb, nil
}

func AppResourceJobFromPb(pb *appspb.AppResourceJobPb) (*AppResourceJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResourceJob{}
	st.Id = pb.Id
	permissionField, err := AppResourceJobJobPermissionFromPb(&pb.Permission)
	if err != nil {
		return nil, err
	}
	if permissionField != nil {
		st.Permission = *permissionField
	}

	return st, nil
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

func AppResourceJobJobPermissionToPb(st *AppResourceJobJobPermission) (*appspb.AppResourceJobJobPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appspb.AppResourceJobJobPermissionPb(*st)
	return &pb, nil
}

func AppResourceJobJobPermissionFromPb(pb *appspb.AppResourceJobJobPermissionPb) (*AppResourceJobJobPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppResourceJobJobPermission(*pb)
	return &st, nil
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

func (st AppResourceSecret) MarshalJSON() ([]byte, error) {
	pb, err := AppResourceSecretToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppResourceSecret) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppResourceSecretPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppResourceSecretFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppResourceSecretToPb(st *AppResourceSecret) (*appspb.AppResourceSecretPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppResourceSecretPb{}
	pb.Key = st.Key
	permissionPb, err := AppResourceSecretSecretPermissionToPb(&st.Permission)
	if err != nil {
		return nil, err
	}
	if permissionPb != nil {
		pb.Permission = *permissionPb
	}
	pb.Scope = st.Scope

	return pb, nil
}

func AppResourceSecretFromPb(pb *appspb.AppResourceSecretPb) (*AppResourceSecret, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResourceSecret{}
	st.Key = pb.Key
	permissionField, err := AppResourceSecretSecretPermissionFromPb(&pb.Permission)
	if err != nil {
		return nil, err
	}
	if permissionField != nil {
		st.Permission = *permissionField
	}
	st.Scope = pb.Scope

	return st, nil
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

func AppResourceSecretSecretPermissionToPb(st *AppResourceSecretSecretPermission) (*appspb.AppResourceSecretSecretPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appspb.AppResourceSecretSecretPermissionPb(*st)
	return &pb, nil
}

func AppResourceSecretSecretPermissionFromPb(pb *appspb.AppResourceSecretSecretPermissionPb) (*AppResourceSecretSecretPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppResourceSecretSecretPermission(*pb)
	return &st, nil
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

func (st AppResourceServingEndpoint) MarshalJSON() ([]byte, error) {
	pb, err := AppResourceServingEndpointToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppResourceServingEndpoint) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppResourceServingEndpointPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppResourceServingEndpointFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppResourceServingEndpointToPb(st *AppResourceServingEndpoint) (*appspb.AppResourceServingEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppResourceServingEndpointPb{}
	pb.Name = st.Name
	permissionPb, err := AppResourceServingEndpointServingEndpointPermissionToPb(&st.Permission)
	if err != nil {
		return nil, err
	}
	if permissionPb != nil {
		pb.Permission = *permissionPb
	}

	return pb, nil
}

func AppResourceServingEndpointFromPb(pb *appspb.AppResourceServingEndpointPb) (*AppResourceServingEndpoint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResourceServingEndpoint{}
	st.Name = pb.Name
	permissionField, err := AppResourceServingEndpointServingEndpointPermissionFromPb(&pb.Permission)
	if err != nil {
		return nil, err
	}
	if permissionField != nil {
		st.Permission = *permissionField
	}

	return st, nil
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

func AppResourceServingEndpointServingEndpointPermissionToPb(st *AppResourceServingEndpointServingEndpointPermission) (*appspb.AppResourceServingEndpointServingEndpointPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appspb.AppResourceServingEndpointServingEndpointPermissionPb(*st)
	return &pb, nil
}

func AppResourceServingEndpointServingEndpointPermissionFromPb(pb *appspb.AppResourceServingEndpointServingEndpointPermissionPb) (*AppResourceServingEndpointServingEndpointPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppResourceServingEndpointServingEndpointPermission(*pb)
	return &st, nil
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

func (st AppResourceSqlWarehouse) MarshalJSON() ([]byte, error) {
	pb, err := AppResourceSqlWarehouseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppResourceSqlWarehouse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppResourceSqlWarehousePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppResourceSqlWarehouseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppResourceSqlWarehouseToPb(st *AppResourceSqlWarehouse) (*appspb.AppResourceSqlWarehousePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppResourceSqlWarehousePb{}
	pb.Id = st.Id
	permissionPb, err := AppResourceSqlWarehouseSqlWarehousePermissionToPb(&st.Permission)
	if err != nil {
		return nil, err
	}
	if permissionPb != nil {
		pb.Permission = *permissionPb
	}

	return pb, nil
}

func AppResourceSqlWarehouseFromPb(pb *appspb.AppResourceSqlWarehousePb) (*AppResourceSqlWarehouse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResourceSqlWarehouse{}
	st.Id = pb.Id
	permissionField, err := AppResourceSqlWarehouseSqlWarehousePermissionFromPb(&pb.Permission)
	if err != nil {
		return nil, err
	}
	if permissionField != nil {
		st.Permission = *permissionField
	}

	return st, nil
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

func AppResourceSqlWarehouseSqlWarehousePermissionToPb(st *AppResourceSqlWarehouseSqlWarehousePermission) (*appspb.AppResourceSqlWarehouseSqlWarehousePermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appspb.AppResourceSqlWarehouseSqlWarehousePermissionPb(*st)
	return &pb, nil
}

func AppResourceSqlWarehouseSqlWarehousePermissionFromPb(pb *appspb.AppResourceSqlWarehouseSqlWarehousePermissionPb) (*AppResourceSqlWarehouseSqlWarehousePermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppResourceSqlWarehouseSqlWarehousePermission(*pb)
	return &st, nil
}

type AppResourceUcSecurable struct {

	// Wire name: 'permission'
	Permission AppResourceUcSecurableUcSecurablePermission `json:"permission"`

	// Wire name: 'securable_full_name'
	SecurableFullName string `json:"securable_full_name"`

	// Wire name: 'securable_type'
	SecurableType AppResourceUcSecurableUcSecurableType `json:"securable_type"`
}

func (st AppResourceUcSecurable) MarshalJSON() ([]byte, error) {
	pb, err := AppResourceUcSecurableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AppResourceUcSecurable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.AppResourceUcSecurablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AppResourceUcSecurableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AppResourceUcSecurableToPb(st *AppResourceUcSecurable) (*appspb.AppResourceUcSecurablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppResourceUcSecurablePb{}
	permissionPb, err := AppResourceUcSecurableUcSecurablePermissionToPb(&st.Permission)
	if err != nil {
		return nil, err
	}
	if permissionPb != nil {
		pb.Permission = *permissionPb
	}
	pb.SecurableFullName = st.SecurableFullName
	securableTypePb, err := AppResourceUcSecurableUcSecurableTypeToPb(&st.SecurableType)
	if err != nil {
		return nil, err
	}
	if securableTypePb != nil {
		pb.SecurableType = *securableTypePb
	}

	return pb, nil
}

func AppResourceUcSecurableFromPb(pb *appspb.AppResourceUcSecurablePb) (*AppResourceUcSecurable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResourceUcSecurable{}
	permissionField, err := AppResourceUcSecurableUcSecurablePermissionFromPb(&pb.Permission)
	if err != nil {
		return nil, err
	}
	if permissionField != nil {
		st.Permission = *permissionField
	}
	st.SecurableFullName = pb.SecurableFullName
	securableTypeField, err := AppResourceUcSecurableUcSecurableTypeFromPb(&pb.SecurableType)
	if err != nil {
		return nil, err
	}
	if securableTypeField != nil {
		st.SecurableType = *securableTypeField
	}

	return st, nil
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

func AppResourceUcSecurableUcSecurablePermissionToPb(st *AppResourceUcSecurableUcSecurablePermission) (*appspb.AppResourceUcSecurableUcSecurablePermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appspb.AppResourceUcSecurableUcSecurablePermissionPb(*st)
	return &pb, nil
}

func AppResourceUcSecurableUcSecurablePermissionFromPb(pb *appspb.AppResourceUcSecurableUcSecurablePermissionPb) (*AppResourceUcSecurableUcSecurablePermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppResourceUcSecurableUcSecurablePermission(*pb)
	return &st, nil
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

func AppResourceUcSecurableUcSecurableTypeToPb(st *AppResourceUcSecurableUcSecurableType) (*appspb.AppResourceUcSecurableUcSecurableTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appspb.AppResourceUcSecurableUcSecurableTypePb(*st)
	return &pb, nil
}

func AppResourceUcSecurableUcSecurableTypeFromPb(pb *appspb.AppResourceUcSecurableUcSecurableTypePb) (*AppResourceUcSecurableUcSecurableType, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppResourceUcSecurableUcSecurableType(*pb)
	return &st, nil
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

func ApplicationStateToPb(st *ApplicationState) (*appspb.ApplicationStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appspb.ApplicationStatePb(*st)
	return &pb, nil
}

func ApplicationStateFromPb(pb *appspb.ApplicationStatePb) (*ApplicationState, error) {
	if pb == nil {
		return nil, nil
	}
	st := ApplicationState(*pb)
	return &st, nil
}

type ApplicationStatus struct {
	// Application status message
	// Wire name: 'message'
	Message string `json:"message,omitempty"`
	// State of the application.
	// Wire name: 'state'
	State           ApplicationState `json:"state,omitempty"`
	ForceSendFields []string         `json:"-" tf:"-"`
}

func (st ApplicationStatus) MarshalJSON() ([]byte, error) {
	pb, err := ApplicationStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ApplicationStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.ApplicationStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ApplicationStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ApplicationStatusToPb(st *ApplicationStatus) (*appspb.ApplicationStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.ApplicationStatusPb{}
	pb.Message = st.Message
	statePb, err := ApplicationStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ApplicationStatusFromPb(pb *appspb.ApplicationStatusPb) (*ApplicationStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ApplicationStatus{}
	st.Message = pb.Message
	stateField, err := ApplicationStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func ComputeStateToPb(st *ComputeState) (*appspb.ComputeStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appspb.ComputeStatePb(*st)
	return &pb, nil
}

func ComputeStateFromPb(pb *appspb.ComputeStatePb) (*ComputeState, error) {
	if pb == nil {
		return nil, nil
	}
	st := ComputeState(*pb)
	return &st, nil
}

type ComputeStatus struct {
	// Compute status message
	// Wire name: 'message'
	Message string `json:"message,omitempty"`
	// State of the app compute.
	// Wire name: 'state'
	State           ComputeState `json:"state,omitempty"`
	ForceSendFields []string     `json:"-" tf:"-"`
}

func (st ComputeStatus) MarshalJSON() ([]byte, error) {
	pb, err := ComputeStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ComputeStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.ComputeStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ComputeStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ComputeStatusToPb(st *ComputeStatus) (*appspb.ComputeStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.ComputeStatusPb{}
	pb.Message = st.Message
	statePb, err := ComputeStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ComputeStatusFromPb(pb *appspb.ComputeStatusPb) (*ComputeStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComputeStatus{}
	st.Message = pb.Message
	stateField, err := ComputeStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateAppDeploymentRequest struct {
	// The app deployment configuration.
	// Wire name: 'app_deployment'
	AppDeployment AppDeployment `json:"app_deployment"`
	// The name of the app.
	AppName string `json:"-" tf:"-"`
}

func (st CreateAppDeploymentRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateAppDeploymentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateAppDeploymentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.CreateAppDeploymentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateAppDeploymentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateAppDeploymentRequestToPb(st *CreateAppDeploymentRequest) (*appspb.CreateAppDeploymentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.CreateAppDeploymentRequestPb{}
	appDeploymentPb, err := AppDeploymentToPb(&st.AppDeployment)
	if err != nil {
		return nil, err
	}
	if appDeploymentPb != nil {
		pb.AppDeployment = *appDeploymentPb
	}
	pb.AppName = st.AppName

	return pb, nil
}

func CreateAppDeploymentRequestFromPb(pb *appspb.CreateAppDeploymentRequestPb) (*CreateAppDeploymentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAppDeploymentRequest{}
	appDeploymentField, err := AppDeploymentFromPb(&pb.AppDeployment)
	if err != nil {
		return nil, err
	}
	if appDeploymentField != nil {
		st.AppDeployment = *appDeploymentField
	}
	st.AppName = pb.AppName

	return st, nil
}

type CreateAppRequest struct {

	// Wire name: 'app'
	App App `json:"app"`
	// If true, the app will not be started after creation.
	NoCompute       bool     `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateAppRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateAppRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateAppRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.CreateAppRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateAppRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateAppRequestToPb(st *CreateAppRequest) (*appspb.CreateAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.CreateAppRequestPb{}
	appPb, err := AppToPb(&st.App)
	if err != nil {
		return nil, err
	}
	if appPb != nil {
		pb.App = *appPb
	}
	pb.NoCompute = st.NoCompute

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateAppRequestFromPb(pb *appspb.CreateAppRequestPb) (*CreateAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAppRequest{}
	appField, err := AppFromPb(&pb.App)
	if err != nil {
		return nil, err
	}
	if appField != nil {
		st.App = *appField
	}
	st.NoCompute = pb.NoCompute

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DeleteAppRequest struct {
	// The name of the app.
	Name string `json:"-" tf:"-"`
}

func (st DeleteAppRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteAppRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteAppRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.DeleteAppRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteAppRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteAppRequestToPb(st *DeleteAppRequest) (*appspb.DeleteAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.DeleteAppRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func DeleteAppRequestFromPb(pb *appspb.DeleteAppRequestPb) (*DeleteAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAppRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetAppDeploymentRequest struct {
	// The name of the app.
	AppName string `json:"-" tf:"-"`
	// The unique id of the deployment.
	DeploymentId string `json:"-" tf:"-"`
}

func (st GetAppDeploymentRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetAppDeploymentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAppDeploymentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.GetAppDeploymentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAppDeploymentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAppDeploymentRequestToPb(st *GetAppDeploymentRequest) (*appspb.GetAppDeploymentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.GetAppDeploymentRequestPb{}
	pb.AppName = st.AppName
	pb.DeploymentId = st.DeploymentId

	return pb, nil
}

func GetAppDeploymentRequestFromPb(pb *appspb.GetAppDeploymentRequestPb) (*GetAppDeploymentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAppDeploymentRequest{}
	st.AppName = pb.AppName
	st.DeploymentId = pb.DeploymentId

	return st, nil
}

type GetAppPermissionLevelsRequest struct {
	// The app for which to get or manage permissions.
	AppName string `json:"-" tf:"-"`
}

func (st GetAppPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetAppPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAppPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.GetAppPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAppPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAppPermissionLevelsRequestToPb(st *GetAppPermissionLevelsRequest) (*appspb.GetAppPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.GetAppPermissionLevelsRequestPb{}
	pb.AppName = st.AppName

	return pb, nil
}

func GetAppPermissionLevelsRequestFromPb(pb *appspb.GetAppPermissionLevelsRequestPb) (*GetAppPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAppPermissionLevelsRequest{}
	st.AppName = pb.AppName

	return st, nil
}

type GetAppPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []AppPermissionsDescription `json:"permission_levels,omitempty"`
}

func (st GetAppPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetAppPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAppPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.GetAppPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAppPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAppPermissionLevelsResponseToPb(st *GetAppPermissionLevelsResponse) (*appspb.GetAppPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.GetAppPermissionLevelsResponsePb{}

	var permissionLevelsPb []appspb.AppPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := AppPermissionsDescriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionLevelsPb = append(permissionLevelsPb, *itemPb)
		}
	}
	pb.PermissionLevels = permissionLevelsPb

	return pb, nil
}

func GetAppPermissionLevelsResponseFromPb(pb *appspb.GetAppPermissionLevelsResponsePb) (*GetAppPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAppPermissionLevelsResponse{}

	var permissionLevelsField []AppPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := AppPermissionsDescriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionLevelsField = append(permissionLevelsField, *item)
		}
	}
	st.PermissionLevels = permissionLevelsField

	return st, nil
}

type GetAppPermissionsRequest struct {
	// The app for which to get or manage permissions.
	AppName string `json:"-" tf:"-"`
}

func (st GetAppPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetAppPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAppPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.GetAppPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAppPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAppPermissionsRequestToPb(st *GetAppPermissionsRequest) (*appspb.GetAppPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.GetAppPermissionsRequestPb{}
	pb.AppName = st.AppName

	return pb, nil
}

func GetAppPermissionsRequestFromPb(pb *appspb.GetAppPermissionsRequestPb) (*GetAppPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAppPermissionsRequest{}
	st.AppName = pb.AppName

	return st, nil
}

type GetAppRequest struct {
	// The name of the app.
	Name string `json:"-" tf:"-"`
}

func (st GetAppRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetAppRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAppRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.GetAppRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAppRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAppRequestToPb(st *GetAppRequest) (*appspb.GetAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.GetAppRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetAppRequestFromPb(pb *appspb.GetAppRequestPb) (*GetAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAppRequest{}
	st.Name = pb.Name

	return st, nil
}

type ListAppDeploymentsRequest struct {
	// The name of the app.
	AppName string `json:"-" tf:"-"`
	// Upper bound for items returned.
	PageSize int `json:"-" tf:"-"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListAppDeploymentsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListAppDeploymentsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListAppDeploymentsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.ListAppDeploymentsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListAppDeploymentsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListAppDeploymentsRequestToPb(st *ListAppDeploymentsRequest) (*appspb.ListAppDeploymentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.ListAppDeploymentsRequestPb{}
	pb.AppName = st.AppName
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListAppDeploymentsRequestFromPb(pb *appspb.ListAppDeploymentsRequestPb) (*ListAppDeploymentsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAppDeploymentsRequest{}
	st.AppName = pb.AppName
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListAppDeploymentsResponse struct {
	// Deployment history of the app.
	// Wire name: 'app_deployments'
	AppDeployments []AppDeployment `json:"app_deployments,omitempty"`
	// Pagination token to request the next page of apps.
	// Wire name: 'next_page_token'
	NextPageToken   string   `json:"next_page_token,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListAppDeploymentsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListAppDeploymentsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListAppDeploymentsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.ListAppDeploymentsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListAppDeploymentsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListAppDeploymentsResponseToPb(st *ListAppDeploymentsResponse) (*appspb.ListAppDeploymentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.ListAppDeploymentsResponsePb{}

	var appDeploymentsPb []appspb.AppDeploymentPb
	for _, item := range st.AppDeployments {
		itemPb, err := AppDeploymentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			appDeploymentsPb = append(appDeploymentsPb, *itemPb)
		}
	}
	pb.AppDeployments = appDeploymentsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListAppDeploymentsResponseFromPb(pb *appspb.ListAppDeploymentsResponsePb) (*ListAppDeploymentsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAppDeploymentsResponse{}

	var appDeploymentsField []AppDeployment
	for _, itemPb := range pb.AppDeployments {
		item, err := AppDeploymentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			appDeploymentsField = append(appDeploymentsField, *item)
		}
	}
	st.AppDeployments = appDeploymentsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListAppsRequest struct {
	// Upper bound for items returned.
	PageSize int `json:"-" tf:"-"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListAppsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListAppsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListAppsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.ListAppsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListAppsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListAppsRequestToPb(st *ListAppsRequest) (*appspb.ListAppsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.ListAppsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListAppsRequestFromPb(pb *appspb.ListAppsRequestPb) (*ListAppsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAppsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListAppsResponse struct {

	// Wire name: 'apps'
	Apps []App `json:"apps,omitempty"`
	// Pagination token to request the next page of apps.
	// Wire name: 'next_page_token'
	NextPageToken   string   `json:"next_page_token,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListAppsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListAppsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListAppsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.ListAppsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListAppsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListAppsResponseToPb(st *ListAppsResponse) (*appspb.ListAppsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.ListAppsResponsePb{}

	var appsPb []appspb.AppPb
	for _, item := range st.Apps {
		itemPb, err := AppToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			appsPb = append(appsPb, *itemPb)
		}
	}
	pb.Apps = appsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListAppsResponseFromPb(pb *appspb.ListAppsResponsePb) (*ListAppsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAppsResponse{}

	var appsField []App
	for _, itemPb := range pb.Apps {
		item, err := AppFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			appsField = append(appsField, *item)
		}
	}
	st.Apps = appsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type StartAppRequest struct {
	// The name of the app.
	Name string `json:"-" tf:"-"`
}

func (st StartAppRequest) MarshalJSON() ([]byte, error) {
	pb, err := StartAppRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *StartAppRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.StartAppRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := StartAppRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func StartAppRequestToPb(st *StartAppRequest) (*appspb.StartAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.StartAppRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func StartAppRequestFromPb(pb *appspb.StartAppRequestPb) (*StartAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartAppRequest{}
	st.Name = pb.Name

	return st, nil
}

type StopAppRequest struct {
	// The name of the app.
	Name string `json:"-" tf:"-"`
}

func (st StopAppRequest) MarshalJSON() ([]byte, error) {
	pb, err := StopAppRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *StopAppRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.StopAppRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := StopAppRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func StopAppRequestToPb(st *StopAppRequest) (*appspb.StopAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.StopAppRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func StopAppRequestFromPb(pb *appspb.StopAppRequestPb) (*StopAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StopAppRequest{}
	st.Name = pb.Name

	return st, nil
}

type UpdateAppRequest struct {

	// Wire name: 'app'
	App App `json:"app"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name string `json:"-" tf:"-"`
}

func (st UpdateAppRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateAppRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateAppRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &appspb.UpdateAppRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateAppRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateAppRequestToPb(st *UpdateAppRequest) (*appspb.UpdateAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.UpdateAppRequestPb{}
	appPb, err := AppToPb(&st.App)
	if err != nil {
		return nil, err
	}
	if appPb != nil {
		pb.App = *appPb
	}
	pb.Name = st.Name

	return pb, nil
}

func UpdateAppRequestFromPb(pb *appspb.UpdateAppRequestPb) (*UpdateAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAppRequest{}
	appField, err := AppFromPb(&pb.App)
	if err != nil {
		return nil, err
	}
	if appField != nil {
		st.App = *appField
	}
	st.Name = pb.Name

	return st, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%.9fs", d.Seconds())
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
