// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/apps/appspb"
)

type App struct {
	// The active deployment of the app. A deployment is considered active when
	// it has been deployed to the app compute.
	// Wire name: 'active_deployment'
	ActiveDeployment *AppDeployment ``

	// Wire name: 'app_status'
	AppStatus *ApplicationStatus ``

	// Wire name: 'budget_policy_id'
	BudgetPolicyId string ``

	// Wire name: 'compute_status'
	ComputeStatus *ComputeStatus ``
	// The creation time of the app. Formatted timestamp in ISO 6801.
	// Wire name: 'create_time'
	CreateTime *time.Time ``
	// The email of the user that created the app.
	// Wire name: 'creator'
	Creator string ``
	// The default workspace file system path of the source code from which app
	// deployment are created. This field tracks the workspace source code path
	// of the last active deployment.
	// Wire name: 'default_source_code_path'
	DefaultSourceCodePath string ``
	// The description of the app.
	// Wire name: 'description'
	Description string ``

	// Wire name: 'effective_budget_policy_id'
	EffectiveBudgetPolicyId string ``
	// The effective api scopes granted to the user access token.
	// Wire name: 'effective_user_api_scopes'
	EffectiveUserApiScopes []string ``
	// The unique identifier of the app.
	// Wire name: 'id'
	Id string ``
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	// Wire name: 'name'
	Name string ``

	// Wire name: 'oauth2_app_client_id'
	Oauth2AppClientId string ``

	// Wire name: 'oauth2_app_integration_id'
	Oauth2AppIntegrationId string ``
	// The pending deployment of the app. A deployment is considered pending
	// when it is being prepared for deployment to the app compute.
	// Wire name: 'pending_deployment'
	PendingDeployment *AppDeployment ``
	// Resources for the app.
	// Wire name: 'resources'
	Resources []AppResource ``

	// Wire name: 'service_principal_client_id'
	ServicePrincipalClientId string ``

	// Wire name: 'service_principal_id'
	ServicePrincipalId int64 ``

	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// The update time of the app. Formatted timestamp in ISO 6801.
	// Wire name: 'update_time'
	UpdateTime *time.Time ``
	// The email of the user that last updated the app.
	// Wire name: 'updater'
	Updater string ``
	// The URL of the app once it is deployed.
	// Wire name: 'url'
	Url string ``

	// Wire name: 'user_api_scopes'
	UserApiScopes   []string ``
	ForceSendFields []string `tf:"-"`
}

func (s *App) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s App) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	createTimePb, err := timestampToPb(st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}
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
	updateTimePb, err := timestampToPb(st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}
	pb.Updater = st.Updater
	pb.Url = st.Url
	pb.UserApiScopes = st.UserApiScopes

	pb.ForceSendFields = st.ForceSendFields
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
	createTimeField, err := timestampFromPb(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = createTimeField
	}
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
	updateTimeField, err := timestampFromPb(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = updateTimeField
	}
	st.Updater = pb.Updater
	st.Url = pb.Url
	st.UserApiScopes = pb.UserApiScopes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AppAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``

	// Wire name: 'permission_level'
	PermissionLevel AppPermissionLevel ``
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AppAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	pb.ForceSendFields = st.ForceSendFields
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AppAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []AppPermission ``
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string ``
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AppAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	pb.ForceSendFields = st.ForceSendFields
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AppDeployment struct {
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	// Wire name: 'create_time'
	CreateTime *time.Time ``
	// The email of the user creates the deployment.
	// Wire name: 'creator'
	Creator string ``
	// The deployment artifacts for an app.
	// Wire name: 'deployment_artifacts'
	DeploymentArtifacts *AppDeploymentArtifacts ``
	// The unique id of the deployment.
	// Wire name: 'deployment_id'
	DeploymentId string ``
	// The mode of which the deployment will manage the source code.
	// Wire name: 'mode'
	Mode AppDeploymentMode ``
	// The workspace file system path of the source code used to create the app
	// deployment. This is different from
	// `deployment_artifacts.source_code_path`, which is the path used by the
	// deployed app. The former refers to the original source code location of
	// the app in the workspace during deployment creation, whereas the latter
	// provides a system generated stable snapshotted source code path used by
	// the deployment.
	// Wire name: 'source_code_path'
	SourceCodePath string ``
	// Status and status message of the deployment
	// Wire name: 'status'
	Status *AppDeploymentStatus ``
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	// Wire name: 'update_time'
	UpdateTime      *time.Time ``
	ForceSendFields []string   `tf:"-"`
}

func (s *AppDeployment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppDeployment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AppDeploymentToPb(st *AppDeployment) (*appspb.AppDeploymentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppDeploymentPb{}
	createTimePb, err := timestampToPb(st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}
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
	updateTimePb, err := timestampToPb(st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AppDeploymentFromPb(pb *appspb.AppDeploymentPb) (*AppDeployment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppDeployment{}
	createTimeField, err := timestampFromPb(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = createTimeField
	}
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
	updateTimeField, err := timestampFromPb(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = updateTimeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AppDeploymentArtifacts struct {
	// The snapshotted workspace file system path of the source code loaded by
	// the deployed app.
	// Wire name: 'source_code_path'
	SourceCodePath  string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AppDeploymentArtifacts) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppDeploymentArtifacts) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AppDeploymentArtifactsToPb(st *AppDeploymentArtifacts) (*appspb.AppDeploymentArtifactsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.AppDeploymentArtifactsPb{}
	pb.SourceCodePath = st.SourceCodePath

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AppDeploymentArtifactsFromPb(pb *appspb.AppDeploymentArtifactsPb) (*AppDeploymentArtifacts, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppDeploymentArtifacts{}
	st.SourceCodePath = pb.SourceCodePath

	st.ForceSendFields = pb.ForceSendFields
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
	Message string ``
	// State of the deployment.
	// Wire name: 'state'
	State           AppDeploymentState ``
	ForceSendFields []string           `tf:"-"`
}

func (s *AppDeploymentStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppDeploymentStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	pb.ForceSendFields = st.ForceSendFields
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AppPermission struct {

	// Wire name: 'inherited'
	Inherited bool ``

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string ``

	// Wire name: 'permission_level'
	PermissionLevel AppPermissionLevel ``
	ForceSendFields []string           `tf:"-"`
}

func (s *AppPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	pb.ForceSendFields = st.ForceSendFields
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

	st.ForceSendFields = pb.ForceSendFields
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
	AccessControlList []AppAccessControlResponse ``

	// Wire name: 'object_id'
	ObjectId string ``

	// Wire name: 'object_type'
	ObjectType      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AppPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	pb.ForceSendFields = st.ForceSendFields
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AppPermissionsDescription struct {

	// Wire name: 'description'
	Description string ``

	// Wire name: 'permission_level'
	PermissionLevel AppPermissionLevel ``
	ForceSendFields []string           `tf:"-"`
}

func (s *AppPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	pb.ForceSendFields = st.ForceSendFields
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AppPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []AppAccessControlRequest ``
	// The app for which to get or manage permissions.
	// Wire name: 'app_name'
	AppName string `tf:"-"`
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
	Database *AppResourceDatabase ``
	// Description of the App Resource.
	// Wire name: 'description'
	Description string ``

	// Wire name: 'job'
	Job *AppResourceJob ``
	// Name of the App Resource.
	// Wire name: 'name'
	Name string ``

	// Wire name: 'secret'
	Secret *AppResourceSecret ``

	// Wire name: 'serving_endpoint'
	ServingEndpoint *AppResourceServingEndpoint ``

	// Wire name: 'sql_warehouse'
	SqlWarehouse *AppResourceSqlWarehouse ``

	// Wire name: 'uc_securable'
	UcSecurable     *AppResourceUcSecurable ``
	ForceSendFields []string                `tf:"-"`
}

func (s *AppResource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppResource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	pb.ForceSendFields = st.ForceSendFields
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AppResourceDatabase struct {

	// Wire name: 'database_name'
	DatabaseName string ``

	// Wire name: 'instance_name'
	InstanceName string ``

	// Wire name: 'permission'
	Permission AppResourceDatabaseDatabasePermission ``
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
	Id string ``
	// Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE",
	// "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW".
	// Wire name: 'permission'
	Permission AppResourceJobJobPermission ``
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
	Key string ``
	// Permission to grant on the secret scope. For secrets, only one permission
	// is allowed. Permission must be one of: "READ", "WRITE", "MANAGE".
	// Wire name: 'permission'
	Permission AppResourceSecretSecretPermission ``
	// Scope of the secret to grant permission on.
	// Wire name: 'scope'
	Scope string ``
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
	Name string ``
	// Permission to grant on the serving endpoint. Supported permissions are:
	// "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW".
	// Wire name: 'permission'
	Permission AppResourceServingEndpointServingEndpointPermission ``
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
	Id string ``
	// Permission to grant on the SQL warehouse. Supported permissions are:
	// "CAN_MANAGE", "CAN_USE", "IS_OWNER".
	// Wire name: 'permission'
	Permission AppResourceSqlWarehouseSqlWarehousePermission ``
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
	Permission AppResourceUcSecurableUcSecurablePermission ``

	// Wire name: 'securable_full_name'
	SecurableFullName string ``

	// Wire name: 'securable_type'
	SecurableType AppResourceUcSecurableUcSecurableType ``
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
	Message string ``
	// State of the application.
	// Wire name: 'state'
	State           ApplicationState ``
	ForceSendFields []string         `tf:"-"`
}

func (s *ApplicationStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ApplicationStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	pb.ForceSendFields = st.ForceSendFields
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

	st.ForceSendFields = pb.ForceSendFields
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
	Message string ``
	// State of the app compute.
	// Wire name: 'state'
	State           ComputeState ``
	ForceSendFields []string     `tf:"-"`
}

func (s *ComputeStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComputeStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	pb.ForceSendFields = st.ForceSendFields
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateAppDeploymentRequest struct {
	// The app deployment configuration.
	// Wire name: 'app_deployment'
	AppDeployment AppDeployment ``
	// The name of the app.
	// Wire name: 'app_name'
	AppName string `tf:"-"`
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
	App App ``
	// If true, the app will not be started after creation.
	// Wire name: 'no_compute'
	NoCompute       bool     `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *CreateAppRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAppRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	pb.ForceSendFields = st.ForceSendFields
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteAppRequest struct {
	// The name of the app.
	// Wire name: 'name'
	Name string `tf:"-"`
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
	// Wire name: 'app_name'
	AppName string `tf:"-"`
	// The unique id of the deployment.
	// Wire name: 'deployment_id'
	DeploymentId string `tf:"-"`
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
	// Wire name: 'app_name'
	AppName string `tf:"-"`
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
	PermissionLevels []AppPermissionsDescription ``
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
	// Wire name: 'app_name'
	AppName string `tf:"-"`
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
	// Wire name: 'name'
	Name string `tf:"-"`
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
	// Wire name: 'app_name'
	AppName string `tf:"-"`
	// Upper bound for items returned.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListAppDeploymentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAppDeploymentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListAppDeploymentsRequestToPb(st *ListAppDeploymentsRequest) (*appspb.ListAppDeploymentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.ListAppDeploymentsRequestPb{}
	pb.AppName = st.AppName
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListAppDeploymentsResponse struct {
	// Deployment history of the app.
	// Wire name: 'app_deployments'
	AppDeployments []AppDeployment ``
	// Pagination token to request the next page of apps.
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListAppDeploymentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAppDeploymentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	pb.ForceSendFields = st.ForceSendFields
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListAppsRequest struct {
	// Upper bound for items returned.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListAppsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAppsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListAppsRequestToPb(st *ListAppsRequest) (*appspb.ListAppsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appspb.ListAppsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListAppsRequestFromPb(pb *appspb.ListAppsRequestPb) (*ListAppsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAppsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListAppsResponse struct {

	// Wire name: 'apps'
	Apps []App ``
	// Pagination token to request the next page of apps.
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListAppsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAppsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	pb.ForceSendFields = st.ForceSendFields
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type StartAppRequest struct {
	// The name of the app.
	// Wire name: 'name'
	Name string `tf:"-"`
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
	// Wire name: 'name'
	Name string `tf:"-"`
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
	App App ``
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	// Wire name: 'name'
	Name string `tf:"-"`
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
