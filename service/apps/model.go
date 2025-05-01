// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps

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

type App struct {
	// The active deployment of the app. A deployment is considered active when
	// it has been deployed to the app compute.
	ActiveDeployment *AppDeployment

	AppStatus *ApplicationStatus

	BudgetPolicyId string

	ComputeStatus *ComputeStatus
	// The creation time of the app. Formatted timestamp in ISO 6801.
	CreateTime string
	// The email of the user that created the app.
	Creator string
	// The default workspace file system path of the source code from which app
	// deployment are created. This field tracks the workspace source code path
	// of the last active deployment.
	DefaultSourceCodePath string
	// The description of the app.
	Description string

	EffectiveBudgetPolicyId string
	// The effective api scopes granted to the user access token.
	EffectiveUserApiScopes []string
	// The unique identifier of the app.
	Id string
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name string

	Oauth2AppClientId string

	Oauth2AppIntegrationId string
	// The pending deployment of the app. A deployment is considered pending
	// when it is being prepared for deployment to the app compute.
	PendingDeployment *AppDeployment
	// Resources for the app.
	Resources []AppResource

	ServicePrincipalClientId string

	ServicePrincipalId int64

	ServicePrincipalName string
	// The update time of the app. Formatted timestamp in ISO 6801.
	UpdateTime string
	// The email of the user that last updated the app.
	Updater string
	// The URL of the app once it is deployed.
	Url string

	UserApiScopes []string

	ForceSendFields []string
}

func appToPb(st *App) (*appPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appPb{}
	activeDeploymentPb, err := appDeploymentToPb(st.ActiveDeployment)
	if err != nil {
		return nil, err
	}
	if activeDeploymentPb != nil {
		pb.ActiveDeployment = activeDeploymentPb
	}

	appStatusPb, err := applicationStatusToPb(st.AppStatus)
	if err != nil {
		return nil, err
	}
	if appStatusPb != nil {
		pb.AppStatus = appStatusPb
	}

	budgetPolicyIdPb, err := identity(&st.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdPb != nil {
		pb.BudgetPolicyId = *budgetPolicyIdPb
	}

	computeStatusPb, err := computeStatusToPb(st.ComputeStatus)
	if err != nil {
		return nil, err
	}
	if computeStatusPb != nil {
		pb.ComputeStatus = computeStatusPb
	}

	createTimePb, err := identity(&st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	creatorPb, err := identity(&st.Creator)
	if err != nil {
		return nil, err
	}
	if creatorPb != nil {
		pb.Creator = *creatorPb
	}

	defaultSourceCodePathPb, err := identity(&st.DefaultSourceCodePath)
	if err != nil {
		return nil, err
	}
	if defaultSourceCodePathPb != nil {
		pb.DefaultSourceCodePath = *defaultSourceCodePathPb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	effectiveBudgetPolicyIdPb, err := identity(&st.EffectiveBudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if effectiveBudgetPolicyIdPb != nil {
		pb.EffectiveBudgetPolicyId = *effectiveBudgetPolicyIdPb
	}

	var effectiveUserApiScopesPb []string
	for _, item := range st.EffectiveUserApiScopes {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			effectiveUserApiScopesPb = append(effectiveUserApiScopesPb, *itemPb)
		}
	}
	pb.EffectiveUserApiScopes = effectiveUserApiScopesPb

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	oauth2AppClientIdPb, err := identity(&st.Oauth2AppClientId)
	if err != nil {
		return nil, err
	}
	if oauth2AppClientIdPb != nil {
		pb.Oauth2AppClientId = *oauth2AppClientIdPb
	}

	oauth2AppIntegrationIdPb, err := identity(&st.Oauth2AppIntegrationId)
	if err != nil {
		return nil, err
	}
	if oauth2AppIntegrationIdPb != nil {
		pb.Oauth2AppIntegrationId = *oauth2AppIntegrationIdPb
	}

	pendingDeploymentPb, err := appDeploymentToPb(st.PendingDeployment)
	if err != nil {
		return nil, err
	}
	if pendingDeploymentPb != nil {
		pb.PendingDeployment = pendingDeploymentPb
	}

	var resourcesPb []appResourcePb
	for _, item := range st.Resources {
		itemPb, err := appResourceToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resourcesPb = append(resourcesPb, *itemPb)
		}
	}
	pb.Resources = resourcesPb

	servicePrincipalClientIdPb, err := identity(&st.ServicePrincipalClientId)
	if err != nil {
		return nil, err
	}
	if servicePrincipalClientIdPb != nil {
		pb.ServicePrincipalClientId = *servicePrincipalClientIdPb
	}

	servicePrincipalIdPb, err := identity(&st.ServicePrincipalId)
	if err != nil {
		return nil, err
	}
	if servicePrincipalIdPb != nil {
		pb.ServicePrincipalId = *servicePrincipalIdPb
	}

	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	updateTimePb, err := identity(&st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

	updaterPb, err := identity(&st.Updater)
	if err != nil {
		return nil, err
	}
	if updaterPb != nil {
		pb.Updater = *updaterPb
	}

	urlPb, err := identity(&st.Url)
	if err != nil {
		return nil, err
	}
	if urlPb != nil {
		pb.Url = *urlPb
	}

	var userApiScopesPb []string
	for _, item := range st.UserApiScopes {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			userApiScopesPb = append(userApiScopesPb, *itemPb)
		}
	}
	pb.UserApiScopes = userApiScopesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type appPb struct {
	// The active deployment of the app. A deployment is considered active when
	// it has been deployed to the app compute.
	ActiveDeployment *appDeploymentPb `json:"active_deployment,omitempty"`

	AppStatus *applicationStatusPb `json:"app_status,omitempty"`

	BudgetPolicyId string `json:"budget_policy_id,omitempty"`

	ComputeStatus *computeStatusPb `json:"compute_status,omitempty"`
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
	// The effective api scopes granted to the user access token.
	EffectiveUserApiScopes []string `json:"effective_user_api_scopes,omitempty"`
	// The unique identifier of the app.
	Id string `json:"id,omitempty"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name string `json:"name"`

	Oauth2AppClientId string `json:"oauth2_app_client_id,omitempty"`

	Oauth2AppIntegrationId string `json:"oauth2_app_integration_id,omitempty"`
	// The pending deployment of the app. A deployment is considered pending
	// when it is being prepared for deployment to the app compute.
	PendingDeployment *appDeploymentPb `json:"pending_deployment,omitempty"`
	// Resources for the app.
	Resources []appResourcePb `json:"resources,omitempty"`

	ServicePrincipalClientId string `json:"service_principal_client_id,omitempty"`

	ServicePrincipalId int64 `json:"service_principal_id,omitempty"`

	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The update time of the app. Formatted timestamp in ISO 6801.
	UpdateTime string `json:"update_time,omitempty"`
	// The email of the user that last updated the app.
	Updater string `json:"updater,omitempty"`
	// The URL of the app once it is deployed.
	Url string `json:"url,omitempty"`

	UserApiScopes []string `json:"user_api_scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appFromPb(pb *appPb) (*App, error) {
	if pb == nil {
		return nil, nil
	}
	st := &App{}
	activeDeploymentField, err := appDeploymentFromPb(pb.ActiveDeployment)
	if err != nil {
		return nil, err
	}
	if activeDeploymentField != nil {
		st.ActiveDeployment = activeDeploymentField
	}
	appStatusField, err := applicationStatusFromPb(pb.AppStatus)
	if err != nil {
		return nil, err
	}
	if appStatusField != nil {
		st.AppStatus = appStatusField
	}
	budgetPolicyIdField, err := identity(&pb.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdField != nil {
		st.BudgetPolicyId = *budgetPolicyIdField
	}
	computeStatusField, err := computeStatusFromPb(pb.ComputeStatus)
	if err != nil {
		return nil, err
	}
	if computeStatusField != nil {
		st.ComputeStatus = computeStatusField
	}
	createTimeField, err := identity(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = *createTimeField
	}
	creatorField, err := identity(&pb.Creator)
	if err != nil {
		return nil, err
	}
	if creatorField != nil {
		st.Creator = *creatorField
	}
	defaultSourceCodePathField, err := identity(&pb.DefaultSourceCodePath)
	if err != nil {
		return nil, err
	}
	if defaultSourceCodePathField != nil {
		st.DefaultSourceCodePath = *defaultSourceCodePathField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	effectiveBudgetPolicyIdField, err := identity(&pb.EffectiveBudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if effectiveBudgetPolicyIdField != nil {
		st.EffectiveBudgetPolicyId = *effectiveBudgetPolicyIdField
	}

	var effectiveUserApiScopesField []string
	for _, itemPb := range pb.EffectiveUserApiScopes {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			effectiveUserApiScopesField = append(effectiveUserApiScopesField, *item)
		}
	}
	st.EffectiveUserApiScopes = effectiveUserApiScopesField
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	oauth2AppClientIdField, err := identity(&pb.Oauth2AppClientId)
	if err != nil {
		return nil, err
	}
	if oauth2AppClientIdField != nil {
		st.Oauth2AppClientId = *oauth2AppClientIdField
	}
	oauth2AppIntegrationIdField, err := identity(&pb.Oauth2AppIntegrationId)
	if err != nil {
		return nil, err
	}
	if oauth2AppIntegrationIdField != nil {
		st.Oauth2AppIntegrationId = *oauth2AppIntegrationIdField
	}
	pendingDeploymentField, err := appDeploymentFromPb(pb.PendingDeployment)
	if err != nil {
		return nil, err
	}
	if pendingDeploymentField != nil {
		st.PendingDeployment = pendingDeploymentField
	}

	var resourcesField []AppResource
	for _, itemPb := range pb.Resources {
		item, err := appResourceFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resourcesField = append(resourcesField, *item)
		}
	}
	st.Resources = resourcesField
	servicePrincipalClientIdField, err := identity(&pb.ServicePrincipalClientId)
	if err != nil {
		return nil, err
	}
	if servicePrincipalClientIdField != nil {
		st.ServicePrincipalClientId = *servicePrincipalClientIdField
	}
	servicePrincipalIdField, err := identity(&pb.ServicePrincipalId)
	if err != nil {
		return nil, err
	}
	if servicePrincipalIdField != nil {
		st.ServicePrincipalId = *servicePrincipalIdField
	}
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	updateTimeField, err := identity(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = *updateTimeField
	}
	updaterField, err := identity(&pb.Updater)
	if err != nil {
		return nil, err
	}
	if updaterField != nil {
		st.Updater = *updaterField
	}
	urlField, err := identity(&pb.Url)
	if err != nil {
		return nil, err
	}
	if urlField != nil {
		st.Url = *urlField
	}

	var userApiScopesField []string
	for _, itemPb := range pb.UserApiScopes {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			userApiScopesField = append(userApiScopesField, *item)
		}
	}
	st.UserApiScopes = userApiScopesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppAccessControlRequest struct {
	// name of the group
	GroupName string
	// Permission level
	PermissionLevel AppPermissionLevel
	// application ID of a service principal
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
}

func appAccessControlRequestToPb(st *AppAccessControlRequest) (*appAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appAccessControlRequestPb{}
	groupNamePb, err := identity(&st.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNamePb != nil {
		pb.GroupName = *groupNamePb
	}

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type appAccessControlRequestPb struct {
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

func appAccessControlRequestFromPb(pb *appAccessControlRequestPb) (*AppAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppAccessControlRequest{}
	groupNameField, err := identity(&pb.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNameField != nil {
		st.GroupName = *groupNameField
	}
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppAccessControlResponse struct {
	// All permissions.
	AllPermissions []AppPermission
	// Display name of the user or service principal.
	DisplayName string
	// name of the group
	GroupName string
	// Name of the service principal.
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
}

func appAccessControlResponseToPb(st *AppAccessControlResponse) (*appAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appAccessControlResponsePb{}

	var allPermissionsPb []appPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := appPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	groupNamePb, err := identity(&st.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNamePb != nil {
		pb.GroupName = *groupNamePb
	}

	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type appAccessControlResponsePb struct {
	// All permissions.
	AllPermissions []appPermissionPb `json:"all_permissions,omitempty"`
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

func appAccessControlResponseFromPb(pb *appAccessControlResponsePb) (*AppAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppAccessControlResponse{}

	var allPermissionsField []AppPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := appPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	groupNameField, err := identity(&pb.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNameField != nil {
		st.GroupName = *groupNameField
	}
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppDeployment struct {
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	CreateTime string
	// The email of the user creates the deployment.
	Creator string
	// The deployment artifacts for an app.
	DeploymentArtifacts *AppDeploymentArtifacts
	// The unique id of the deployment.
	DeploymentId string
	// The mode of which the deployment will manage the source code.
	Mode AppDeploymentMode
	// The workspace file system path of the source code used to create the app
	// deployment. This is different from
	// `deployment_artifacts.source_code_path`, which is the path used by the
	// deployed app. The former refers to the original source code location of
	// the app in the workspace during deployment creation, whereas the latter
	// provides a system generated stable snapshotted source code path used by
	// the deployment.
	SourceCodePath string
	// Status and status message of the deployment
	Status *AppDeploymentStatus
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	UpdateTime string

	ForceSendFields []string
}

func appDeploymentToPb(st *AppDeployment) (*appDeploymentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appDeploymentPb{}
	createTimePb, err := identity(&st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	creatorPb, err := identity(&st.Creator)
	if err != nil {
		return nil, err
	}
	if creatorPb != nil {
		pb.Creator = *creatorPb
	}

	deploymentArtifactsPb, err := appDeploymentArtifactsToPb(st.DeploymentArtifacts)
	if err != nil {
		return nil, err
	}
	if deploymentArtifactsPb != nil {
		pb.DeploymentArtifacts = deploymentArtifactsPb
	}

	deploymentIdPb, err := identity(&st.DeploymentId)
	if err != nil {
		return nil, err
	}
	if deploymentIdPb != nil {
		pb.DeploymentId = *deploymentIdPb
	}

	modePb, err := identity(&st.Mode)
	if err != nil {
		return nil, err
	}
	if modePb != nil {
		pb.Mode = *modePb
	}

	sourceCodePathPb, err := identity(&st.SourceCodePath)
	if err != nil {
		return nil, err
	}
	if sourceCodePathPb != nil {
		pb.SourceCodePath = *sourceCodePathPb
	}

	statusPb, err := appDeploymentStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}

	updateTimePb, err := identity(&st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type appDeploymentPb struct {
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	CreateTime string `json:"create_time,omitempty"`
	// The email of the user creates the deployment.
	Creator string `json:"creator,omitempty"`
	// The deployment artifacts for an app.
	DeploymentArtifacts *appDeploymentArtifactsPb `json:"deployment_artifacts,omitempty"`
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
	Status *appDeploymentStatusPb `json:"status,omitempty"`
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appDeploymentFromPb(pb *appDeploymentPb) (*AppDeployment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppDeployment{}
	createTimeField, err := identity(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = *createTimeField
	}
	creatorField, err := identity(&pb.Creator)
	if err != nil {
		return nil, err
	}
	if creatorField != nil {
		st.Creator = *creatorField
	}
	deploymentArtifactsField, err := appDeploymentArtifactsFromPb(pb.DeploymentArtifacts)
	if err != nil {
		return nil, err
	}
	if deploymentArtifactsField != nil {
		st.DeploymentArtifacts = deploymentArtifactsField
	}
	deploymentIdField, err := identity(&pb.DeploymentId)
	if err != nil {
		return nil, err
	}
	if deploymentIdField != nil {
		st.DeploymentId = *deploymentIdField
	}
	modeField, err := identity(&pb.Mode)
	if err != nil {
		return nil, err
	}
	if modeField != nil {
		st.Mode = *modeField
	}
	sourceCodePathField, err := identity(&pb.SourceCodePath)
	if err != nil {
		return nil, err
	}
	if sourceCodePathField != nil {
		st.SourceCodePath = *sourceCodePathField
	}
	statusField, err := appDeploymentStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}
	updateTimeField, err := identity(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = *updateTimeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appDeploymentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appDeploymentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppDeploymentArtifacts struct {
	// The snapshotted workspace file system path of the source code loaded by
	// the deployed app.
	SourceCodePath string

	ForceSendFields []string
}

func appDeploymentArtifactsToPb(st *AppDeploymentArtifacts) (*appDeploymentArtifactsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appDeploymentArtifactsPb{}
	sourceCodePathPb, err := identity(&st.SourceCodePath)
	if err != nil {
		return nil, err
	}
	if sourceCodePathPb != nil {
		pb.SourceCodePath = *sourceCodePathPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type appDeploymentArtifactsPb struct {
	// The snapshotted workspace file system path of the source code loaded by
	// the deployed app.
	SourceCodePath string `json:"source_code_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appDeploymentArtifactsFromPb(pb *appDeploymentArtifactsPb) (*AppDeploymentArtifacts, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppDeploymentArtifacts{}
	sourceCodePathField, err := identity(&pb.SourceCodePath)
	if err != nil {
		return nil, err
	}
	if sourceCodePathField != nil {
		st.SourceCodePath = *sourceCodePathField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appDeploymentArtifactsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appDeploymentArtifactsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppDeploymentMode string
type appDeploymentModePb string

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

func appDeploymentModeToPb(st *AppDeploymentMode) (*appDeploymentModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appDeploymentModePb(*st)
	return &pb, nil
}

func appDeploymentModeFromPb(pb *appDeploymentModePb) (*AppDeploymentMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppDeploymentMode(*pb)
	return &st, nil
}

type AppDeploymentState string
type appDeploymentStatePb string

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

func appDeploymentStateToPb(st *AppDeploymentState) (*appDeploymentStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appDeploymentStatePb(*st)
	return &pb, nil
}

func appDeploymentStateFromPb(pb *appDeploymentStatePb) (*AppDeploymentState, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppDeploymentState(*pb)
	return &st, nil
}

type AppDeploymentStatus struct {
	// Message corresponding with the deployment state.
	Message string
	// State of the deployment.
	State AppDeploymentState

	ForceSendFields []string
}

func appDeploymentStatusToPb(st *AppDeploymentStatus) (*appDeploymentStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appDeploymentStatusPb{}
	messagePb, err := identity(&st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = *messagePb
	}

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type appDeploymentStatusPb struct {
	// Message corresponding with the deployment state.
	Message string `json:"message,omitempty"`
	// State of the deployment.
	State AppDeploymentState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appDeploymentStatusFromPb(pb *appDeploymentStatusPb) (*AppDeploymentStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppDeploymentStatus{}
	messageField, err := identity(&pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = *messageField
	}
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appDeploymentStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appDeploymentStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppPermission struct {
	Inherited bool

	InheritedFromObject []string
	// Permission level
	PermissionLevel AppPermissionLevel

	ForceSendFields []string
}

func appPermissionToPb(st *AppPermission) (*appPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appPermissionPb{}
	inheritedPb, err := identity(&st.Inherited)
	if err != nil {
		return nil, err
	}
	if inheritedPb != nil {
		pb.Inherited = *inheritedPb
	}

	var inheritedFromObjectPb []string
	for _, item := range st.InheritedFromObject {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			inheritedFromObjectPb = append(inheritedFromObjectPb, *itemPb)
		}
	}
	pb.InheritedFromObject = inheritedFromObjectPb

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type appPermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel AppPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appPermissionFromPb(pb *appPermissionPb) (*AppPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppPermission{}
	inheritedField, err := identity(&pb.Inherited)
	if err != nil {
		return nil, err
	}
	if inheritedField != nil {
		st.Inherited = *inheritedField
	}

	var inheritedFromObjectField []string
	for _, itemPb := range pb.InheritedFromObject {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			inheritedFromObjectField = append(inheritedFromObjectField, *item)
		}
	}
	st.InheritedFromObject = inheritedFromObjectField
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Permission level
type AppPermissionLevel string
type appPermissionLevelPb string

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

func appPermissionLevelToPb(st *AppPermissionLevel) (*appPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appPermissionLevelPb(*st)
	return &pb, nil
}

func appPermissionLevelFromPb(pb *appPermissionLevelPb) (*AppPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppPermissionLevel(*pb)
	return &st, nil
}

type AppPermissions struct {
	AccessControlList []AppAccessControlResponse

	ObjectId string

	ObjectType string

	ForceSendFields []string
}

func appPermissionsToPb(st *AppPermissions) (*appPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appPermissionsPb{}

	var accessControlListPb []appAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := appAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	objectIdPb, err := identity(&st.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdPb != nil {
		pb.ObjectId = *objectIdPb
	}

	objectTypePb, err := identity(&st.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypePb != nil {
		pb.ObjectType = *objectTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type appPermissionsPb struct {
	AccessControlList []appAccessControlResponsePb `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appPermissionsFromPb(pb *appPermissionsPb) (*AppPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppPermissions{}

	var accessControlListField []AppAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := appAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	objectIdField, err := identity(&pb.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdField != nil {
		st.ObjectId = *objectIdField
	}
	objectTypeField, err := identity(&pb.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypeField != nil {
		st.ObjectType = *objectTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppPermissionsDescription struct {
	Description string
	// Permission level
	PermissionLevel AppPermissionLevel

	ForceSendFields []string
}

func appPermissionsDescriptionToPb(st *AppPermissionsDescription) (*appPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appPermissionsDescriptionPb{}
	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type appPermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel AppPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appPermissionsDescriptionFromPb(pb *appPermissionsDescriptionPb) (*AppPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppPermissionsDescription{}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppPermissionsRequest struct {
	AccessControlList []AppAccessControlRequest
	// The app for which to get or manage permissions.
	AppName string
}

func appPermissionsRequestToPb(st *AppPermissionsRequest) (*appPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appPermissionsRequestPb{}

	var accessControlListPb []appAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := appAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	appNamePb, err := identity(&st.AppName)
	if err != nil {
		return nil, err
	}
	if appNamePb != nil {
		pb.AppName = *appNamePb
	}

	return pb, nil
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

type appPermissionsRequestPb struct {
	AccessControlList []appAccessControlRequestPb `json:"access_control_list,omitempty"`
	// The app for which to get or manage permissions.
	AppName string `json:"-" url:"-"`
}

func appPermissionsRequestFromPb(pb *appPermissionsRequestPb) (*AppPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppPermissionsRequest{}

	var accessControlListField []AppAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := appAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	appNameField, err := identity(&pb.AppName)
	if err != nil {
		return nil, err
	}
	if appNameField != nil {
		st.AppName = *appNameField
	}

	return st, nil
}

type AppResource struct {
	// Description of the App Resource.
	Description string

	Job *AppResourceJob
	// Name of the App Resource.
	Name string

	Secret *AppResourceSecret

	ServingEndpoint *AppResourceServingEndpoint

	SqlWarehouse *AppResourceSqlWarehouse

	ForceSendFields []string
}

func appResourceToPb(st *AppResource) (*appResourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appResourcePb{}
	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	jobPb, err := appResourceJobToPb(st.Job)
	if err != nil {
		return nil, err
	}
	if jobPb != nil {
		pb.Job = jobPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	secretPb, err := appResourceSecretToPb(st.Secret)
	if err != nil {
		return nil, err
	}
	if secretPb != nil {
		pb.Secret = secretPb
	}

	servingEndpointPb, err := appResourceServingEndpointToPb(st.ServingEndpoint)
	if err != nil {
		return nil, err
	}
	if servingEndpointPb != nil {
		pb.ServingEndpoint = servingEndpointPb
	}

	sqlWarehousePb, err := appResourceSqlWarehouseToPb(st.SqlWarehouse)
	if err != nil {
		return nil, err
	}
	if sqlWarehousePb != nil {
		pb.SqlWarehouse = sqlWarehousePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type appResourcePb struct {
	// Description of the App Resource.
	Description string `json:"description,omitempty"`

	Job *appResourceJobPb `json:"job,omitempty"`
	// Name of the App Resource.
	Name string `json:"name"`

	Secret *appResourceSecretPb `json:"secret,omitempty"`

	ServingEndpoint *appResourceServingEndpointPb `json:"serving_endpoint,omitempty"`

	SqlWarehouse *appResourceSqlWarehousePb `json:"sql_warehouse,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func appResourceFromPb(pb *appResourcePb) (*AppResource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResource{}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	jobField, err := appResourceJobFromPb(pb.Job)
	if err != nil {
		return nil, err
	}
	if jobField != nil {
		st.Job = jobField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	secretField, err := appResourceSecretFromPb(pb.Secret)
	if err != nil {
		return nil, err
	}
	if secretField != nil {
		st.Secret = secretField
	}
	servingEndpointField, err := appResourceServingEndpointFromPb(pb.ServingEndpoint)
	if err != nil {
		return nil, err
	}
	if servingEndpointField != nil {
		st.ServingEndpoint = servingEndpointField
	}
	sqlWarehouseField, err := appResourceSqlWarehouseFromPb(pb.SqlWarehouse)
	if err != nil {
		return nil, err
	}
	if sqlWarehouseField != nil {
		st.SqlWarehouse = sqlWarehouseField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *appResourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st appResourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AppResourceJob struct {
	// Id of the job to grant permission on.
	Id string
	// Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE",
	// "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW".
	Permission AppResourceJobJobPermission
}

func appResourceJobToPb(st *AppResourceJob) (*appResourceJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appResourceJobPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	permissionPb, err := identity(&st.Permission)
	if err != nil {
		return nil, err
	}
	if permissionPb != nil {
		pb.Permission = *permissionPb
	}

	return pb, nil
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

type appResourceJobPb struct {
	// Id of the job to grant permission on.
	Id string `json:"id"`
	// Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE",
	// "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW".
	Permission AppResourceJobJobPermission `json:"permission"`
}

func appResourceJobFromPb(pb *appResourceJobPb) (*AppResourceJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResourceJob{}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	permissionField, err := identity(&pb.Permission)
	if err != nil {
		return nil, err
	}
	if permissionField != nil {
		st.Permission = *permissionField
	}

	return st, nil
}

type AppResourceJobJobPermission string
type appResourceJobJobPermissionPb string

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

func appResourceJobJobPermissionToPb(st *AppResourceJobJobPermission) (*appResourceJobJobPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appResourceJobJobPermissionPb(*st)
	return &pb, nil
}

func appResourceJobJobPermissionFromPb(pb *appResourceJobJobPermissionPb) (*AppResourceJobJobPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppResourceJobJobPermission(*pb)
	return &st, nil
}

type AppResourceSecret struct {
	// Key of the secret to grant permission on.
	Key string
	// Permission to grant on the secret scope. For secrets, only one permission
	// is allowed. Permission must be one of: "READ", "WRITE", "MANAGE".
	Permission AppResourceSecretSecretPermission
	// Scope of the secret to grant permission on.
	Scope string
}

func appResourceSecretToPb(st *AppResourceSecret) (*appResourceSecretPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appResourceSecretPb{}
	keyPb, err := identity(&st.Key)
	if err != nil {
		return nil, err
	}
	if keyPb != nil {
		pb.Key = *keyPb
	}

	permissionPb, err := identity(&st.Permission)
	if err != nil {
		return nil, err
	}
	if permissionPb != nil {
		pb.Permission = *permissionPb
	}

	scopePb, err := identity(&st.Scope)
	if err != nil {
		return nil, err
	}
	if scopePb != nil {
		pb.Scope = *scopePb
	}

	return pb, nil
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

type appResourceSecretPb struct {
	// Key of the secret to grant permission on.
	Key string `json:"key"`
	// Permission to grant on the secret scope. For secrets, only one permission
	// is allowed. Permission must be one of: "READ", "WRITE", "MANAGE".
	Permission AppResourceSecretSecretPermission `json:"permission"`
	// Scope of the secret to grant permission on.
	Scope string `json:"scope"`
}

func appResourceSecretFromPb(pb *appResourceSecretPb) (*AppResourceSecret, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResourceSecret{}
	keyField, err := identity(&pb.Key)
	if err != nil {
		return nil, err
	}
	if keyField != nil {
		st.Key = *keyField
	}
	permissionField, err := identity(&pb.Permission)
	if err != nil {
		return nil, err
	}
	if permissionField != nil {
		st.Permission = *permissionField
	}
	scopeField, err := identity(&pb.Scope)
	if err != nil {
		return nil, err
	}
	if scopeField != nil {
		st.Scope = *scopeField
	}

	return st, nil
}

// Permission to grant on the secret scope. Supported permissions are: "READ",
// "WRITE", "MANAGE".
type AppResourceSecretSecretPermission string
type appResourceSecretSecretPermissionPb string

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

func appResourceSecretSecretPermissionToPb(st *AppResourceSecretSecretPermission) (*appResourceSecretSecretPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appResourceSecretSecretPermissionPb(*st)
	return &pb, nil
}

func appResourceSecretSecretPermissionFromPb(pb *appResourceSecretSecretPermissionPb) (*AppResourceSecretSecretPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppResourceSecretSecretPermission(*pb)
	return &st, nil
}

type AppResourceServingEndpoint struct {
	// Name of the serving endpoint to grant permission on.
	Name string
	// Permission to grant on the serving endpoint. Supported permissions are:
	// "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW".
	Permission AppResourceServingEndpointServingEndpointPermission
}

func appResourceServingEndpointToPb(st *AppResourceServingEndpoint) (*appResourceServingEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appResourceServingEndpointPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	permissionPb, err := identity(&st.Permission)
	if err != nil {
		return nil, err
	}
	if permissionPb != nil {
		pb.Permission = *permissionPb
	}

	return pb, nil
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

type appResourceServingEndpointPb struct {
	// Name of the serving endpoint to grant permission on.
	Name string `json:"name"`
	// Permission to grant on the serving endpoint. Supported permissions are:
	// "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW".
	Permission AppResourceServingEndpointServingEndpointPermission `json:"permission"`
}

func appResourceServingEndpointFromPb(pb *appResourceServingEndpointPb) (*AppResourceServingEndpoint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResourceServingEndpoint{}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	permissionField, err := identity(&pb.Permission)
	if err != nil {
		return nil, err
	}
	if permissionField != nil {
		st.Permission = *permissionField
	}

	return st, nil
}

type AppResourceServingEndpointServingEndpointPermission string
type appResourceServingEndpointServingEndpointPermissionPb string

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

func appResourceServingEndpointServingEndpointPermissionToPb(st *AppResourceServingEndpointServingEndpointPermission) (*appResourceServingEndpointServingEndpointPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appResourceServingEndpointServingEndpointPermissionPb(*st)
	return &pb, nil
}

func appResourceServingEndpointServingEndpointPermissionFromPb(pb *appResourceServingEndpointServingEndpointPermissionPb) (*AppResourceServingEndpointServingEndpointPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppResourceServingEndpointServingEndpointPermission(*pb)
	return &st, nil
}

type AppResourceSqlWarehouse struct {
	// Id of the SQL warehouse to grant permission on.
	Id string
	// Permission to grant on the SQL warehouse. Supported permissions are:
	// "CAN_MANAGE", "CAN_USE", "IS_OWNER".
	Permission AppResourceSqlWarehouseSqlWarehousePermission
}

func appResourceSqlWarehouseToPb(st *AppResourceSqlWarehouse) (*appResourceSqlWarehousePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &appResourceSqlWarehousePb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	permissionPb, err := identity(&st.Permission)
	if err != nil {
		return nil, err
	}
	if permissionPb != nil {
		pb.Permission = *permissionPb
	}

	return pb, nil
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

type appResourceSqlWarehousePb struct {
	// Id of the SQL warehouse to grant permission on.
	Id string `json:"id"`
	// Permission to grant on the SQL warehouse. Supported permissions are:
	// "CAN_MANAGE", "CAN_USE", "IS_OWNER".
	Permission AppResourceSqlWarehouseSqlWarehousePermission `json:"permission"`
}

func appResourceSqlWarehouseFromPb(pb *appResourceSqlWarehousePb) (*AppResourceSqlWarehouse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AppResourceSqlWarehouse{}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	permissionField, err := identity(&pb.Permission)
	if err != nil {
		return nil, err
	}
	if permissionField != nil {
		st.Permission = *permissionField
	}

	return st, nil
}

type AppResourceSqlWarehouseSqlWarehousePermission string
type appResourceSqlWarehouseSqlWarehousePermissionPb string

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

func appResourceSqlWarehouseSqlWarehousePermissionToPb(st *AppResourceSqlWarehouseSqlWarehousePermission) (*appResourceSqlWarehouseSqlWarehousePermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := appResourceSqlWarehouseSqlWarehousePermissionPb(*st)
	return &pb, nil
}

func appResourceSqlWarehouseSqlWarehousePermissionFromPb(pb *appResourceSqlWarehouseSqlWarehousePermissionPb) (*AppResourceSqlWarehouseSqlWarehousePermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := AppResourceSqlWarehouseSqlWarehousePermission(*pb)
	return &st, nil
}

type ApplicationState string
type applicationStatePb string

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

func applicationStateToPb(st *ApplicationState) (*applicationStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := applicationStatePb(*st)
	return &pb, nil
}

func applicationStateFromPb(pb *applicationStatePb) (*ApplicationState, error) {
	if pb == nil {
		return nil, nil
	}
	st := ApplicationState(*pb)
	return &st, nil
}

type ApplicationStatus struct {
	// Application status message
	Message string
	// State of the application.
	State ApplicationState

	ForceSendFields []string
}

func applicationStatusToPb(st *ApplicationStatus) (*applicationStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &applicationStatusPb{}
	messagePb, err := identity(&st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = *messagePb
	}

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type applicationStatusPb struct {
	// Application status message
	Message string `json:"message,omitempty"`
	// State of the application.
	State ApplicationState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func applicationStatusFromPb(pb *applicationStatusPb) (*ApplicationStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ApplicationStatus{}
	messageField, err := identity(&pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = *messageField
	}
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *applicationStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st applicationStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ComputeState string
type computeStatePb string

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

func computeStateToPb(st *ComputeState) (*computeStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computeStatePb(*st)
	return &pb, nil
}

func computeStateFromPb(pb *computeStatePb) (*ComputeState, error) {
	if pb == nil {
		return nil, nil
	}
	st := ComputeState(*pb)
	return &st, nil
}

type ComputeStatus struct {
	// Compute status message
	Message string
	// State of the app compute.
	State ComputeState

	ForceSendFields []string
}

func computeStatusToPb(st *ComputeStatus) (*computeStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computeStatusPb{}
	messagePb, err := identity(&st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = *messagePb
	}

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type computeStatusPb struct {
	// Compute status message
	Message string `json:"message,omitempty"`
	// State of the app compute.
	State ComputeState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func computeStatusFromPb(pb *computeStatusPb) (*ComputeStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComputeStatus{}
	messageField, err := identity(&pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = *messageField
	}
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *computeStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st computeStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Create an app deployment
type CreateAppDeploymentRequest struct {
	AppDeployment AppDeployment
	// The name of the app.
	AppName string
}

func createAppDeploymentRequestToPb(st *CreateAppDeploymentRequest) (*createAppDeploymentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAppDeploymentRequestPb{}
	appDeploymentPb, err := appDeploymentToPb(&st.AppDeployment)
	if err != nil {
		return nil, err
	}
	if appDeploymentPb != nil {
		pb.AppDeployment = *appDeploymentPb
	}

	appNamePb, err := identity(&st.AppName)
	if err != nil {
		return nil, err
	}
	if appNamePb != nil {
		pb.AppName = *appNamePb
	}

	return pb, nil
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

type createAppDeploymentRequestPb struct {
	AppDeployment appDeploymentPb `json:"app_deployment"`
	// The name of the app.
	AppName string `json:"-" url:"-"`
}

func createAppDeploymentRequestFromPb(pb *createAppDeploymentRequestPb) (*CreateAppDeploymentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAppDeploymentRequest{}
	appDeploymentField, err := appDeploymentFromPb(&pb.AppDeployment)
	if err != nil {
		return nil, err
	}
	if appDeploymentField != nil {
		st.AppDeployment = *appDeploymentField
	}
	appNameField, err := identity(&pb.AppName)
	if err != nil {
		return nil, err
	}
	if appNameField != nil {
		st.AppName = *appNameField
	}

	return st, nil
}

// Create an app
type CreateAppRequest struct {
	App App
	// If true, the app will not be started after creation.
	NoCompute bool

	ForceSendFields []string
}

func createAppRequestToPb(st *CreateAppRequest) (*createAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAppRequestPb{}
	appPb, err := appToPb(&st.App)
	if err != nil {
		return nil, err
	}
	if appPb != nil {
		pb.App = *appPb
	}

	noComputePb, err := identity(&st.NoCompute)
	if err != nil {
		return nil, err
	}
	if noComputePb != nil {
		pb.NoCompute = *noComputePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createAppRequestPb struct {
	App appPb `json:"app"`
	// If true, the app will not be started after creation.
	NoCompute bool `json:"-" url:"no_compute,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createAppRequestFromPb(pb *createAppRequestPb) (*CreateAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAppRequest{}
	appField, err := appFromPb(&pb.App)
	if err != nil {
		return nil, err
	}
	if appField != nil {
		st.App = *appField
	}
	noComputeField, err := identity(&pb.NoCompute)
	if err != nil {
		return nil, err
	}
	if noComputeField != nil {
		st.NoCompute = *noComputeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createAppRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createAppRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete an app
type DeleteAppRequest struct {
	// The name of the app.
	Name string
}

func deleteAppRequestToPb(st *DeleteAppRequest) (*deleteAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAppRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	return pb, nil
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

type deleteAppRequestPb struct {
	// The name of the app.
	Name string `json:"-" url:"-"`
}

func deleteAppRequestFromPb(pb *deleteAppRequestPb) (*DeleteAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAppRequest{}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

// Get an app deployment
type GetAppDeploymentRequest struct {
	// The name of the app.
	AppName string
	// The unique id of the deployment.
	DeploymentId string
}

func getAppDeploymentRequestToPb(st *GetAppDeploymentRequest) (*getAppDeploymentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAppDeploymentRequestPb{}
	appNamePb, err := identity(&st.AppName)
	if err != nil {
		return nil, err
	}
	if appNamePb != nil {
		pb.AppName = *appNamePb
	}

	deploymentIdPb, err := identity(&st.DeploymentId)
	if err != nil {
		return nil, err
	}
	if deploymentIdPb != nil {
		pb.DeploymentId = *deploymentIdPb
	}

	return pb, nil
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

type getAppDeploymentRequestPb struct {
	// The name of the app.
	AppName string `json:"-" url:"-"`
	// The unique id of the deployment.
	DeploymentId string `json:"-" url:"-"`
}

func getAppDeploymentRequestFromPb(pb *getAppDeploymentRequestPb) (*GetAppDeploymentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAppDeploymentRequest{}
	appNameField, err := identity(&pb.AppName)
	if err != nil {
		return nil, err
	}
	if appNameField != nil {
		st.AppName = *appNameField
	}
	deploymentIdField, err := identity(&pb.DeploymentId)
	if err != nil {
		return nil, err
	}
	if deploymentIdField != nil {
		st.DeploymentId = *deploymentIdField
	}

	return st, nil
}

// Get app permission levels
type GetAppPermissionLevelsRequest struct {
	// The app for which to get or manage permissions.
	AppName string
}

func getAppPermissionLevelsRequestToPb(st *GetAppPermissionLevelsRequest) (*getAppPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAppPermissionLevelsRequestPb{}
	appNamePb, err := identity(&st.AppName)
	if err != nil {
		return nil, err
	}
	if appNamePb != nil {
		pb.AppName = *appNamePb
	}

	return pb, nil
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

type getAppPermissionLevelsRequestPb struct {
	// The app for which to get or manage permissions.
	AppName string `json:"-" url:"-"`
}

func getAppPermissionLevelsRequestFromPb(pb *getAppPermissionLevelsRequestPb) (*GetAppPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAppPermissionLevelsRequest{}
	appNameField, err := identity(&pb.AppName)
	if err != nil {
		return nil, err
	}
	if appNameField != nil {
		st.AppName = *appNameField
	}

	return st, nil
}

type GetAppPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []AppPermissionsDescription
}

func getAppPermissionLevelsResponseToPb(st *GetAppPermissionLevelsResponse) (*getAppPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAppPermissionLevelsResponsePb{}

	var permissionLevelsPb []appPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := appPermissionsDescriptionToPb(&item)
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

type getAppPermissionLevelsResponsePb struct {
	// Specific permission levels
	PermissionLevels []appPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

func getAppPermissionLevelsResponseFromPb(pb *getAppPermissionLevelsResponsePb) (*GetAppPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAppPermissionLevelsResponse{}

	var permissionLevelsField []AppPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := appPermissionsDescriptionFromPb(&itemPb)
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

// Get app permissions
type GetAppPermissionsRequest struct {
	// The app for which to get or manage permissions.
	AppName string
}

func getAppPermissionsRequestToPb(st *GetAppPermissionsRequest) (*getAppPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAppPermissionsRequestPb{}
	appNamePb, err := identity(&st.AppName)
	if err != nil {
		return nil, err
	}
	if appNamePb != nil {
		pb.AppName = *appNamePb
	}

	return pb, nil
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

type getAppPermissionsRequestPb struct {
	// The app for which to get or manage permissions.
	AppName string `json:"-" url:"-"`
}

func getAppPermissionsRequestFromPb(pb *getAppPermissionsRequestPb) (*GetAppPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAppPermissionsRequest{}
	appNameField, err := identity(&pb.AppName)
	if err != nil {
		return nil, err
	}
	if appNameField != nil {
		st.AppName = *appNameField
	}

	return st, nil
}

// Get an app
type GetAppRequest struct {
	// The name of the app.
	Name string
}

func getAppRequestToPb(st *GetAppRequest) (*getAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAppRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	return pb, nil
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

type getAppRequestPb struct {
	// The name of the app.
	Name string `json:"-" url:"-"`
}

func getAppRequestFromPb(pb *getAppRequestPb) (*GetAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAppRequest{}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

// List app deployments
type ListAppDeploymentsRequest struct {
	// The name of the app.
	AppName string
	// Upper bound for items returned.
	PageSize int
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken string

	ForceSendFields []string
}

func listAppDeploymentsRequestToPb(st *ListAppDeploymentsRequest) (*listAppDeploymentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAppDeploymentsRequestPb{}
	appNamePb, err := identity(&st.AppName)
	if err != nil {
		return nil, err
	}
	if appNamePb != nil {
		pb.AppName = *appNamePb
	}

	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listAppDeploymentsRequestPb struct {
	// The name of the app.
	AppName string `json:"-" url:"-"`
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAppDeploymentsRequestFromPb(pb *listAppDeploymentsRequestPb) (*ListAppDeploymentsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAppDeploymentsRequest{}
	appNameField, err := identity(&pb.AppName)
	if err != nil {
		return nil, err
	}
	if appNameField != nil {
		st.AppName = *appNameField
	}
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAppDeploymentsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAppDeploymentsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAppDeploymentsResponse struct {
	// Deployment history of the app.
	AppDeployments []AppDeployment
	// Pagination token to request the next page of apps.
	NextPageToken string

	ForceSendFields []string
}

func listAppDeploymentsResponseToPb(st *ListAppDeploymentsResponse) (*listAppDeploymentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAppDeploymentsResponsePb{}

	var appDeploymentsPb []appDeploymentPb
	for _, item := range st.AppDeployments {
		itemPb, err := appDeploymentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			appDeploymentsPb = append(appDeploymentsPb, *itemPb)
		}
	}
	pb.AppDeployments = appDeploymentsPb

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listAppDeploymentsResponsePb struct {
	// Deployment history of the app.
	AppDeployments []appDeploymentPb `json:"app_deployments,omitempty"`
	// Pagination token to request the next page of apps.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAppDeploymentsResponseFromPb(pb *listAppDeploymentsResponsePb) (*ListAppDeploymentsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAppDeploymentsResponse{}

	var appDeploymentsField []AppDeployment
	for _, itemPb := range pb.AppDeployments {
		item, err := appDeploymentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			appDeploymentsField = append(appDeploymentsField, *item)
		}
	}
	st.AppDeployments = appDeploymentsField
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAppDeploymentsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAppDeploymentsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List apps
type ListAppsRequest struct {
	// Upper bound for items returned.
	PageSize int
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken string

	ForceSendFields []string
}

func listAppsRequestToPb(st *ListAppsRequest) (*listAppsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAppsRequestPb{}
	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listAppsRequestPb struct {
	// Upper bound for items returned.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAppsRequestFromPb(pb *listAppsRequestPb) (*ListAppsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAppsRequest{}
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAppsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAppsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAppsResponse struct {
	Apps []App
	// Pagination token to request the next page of apps.
	NextPageToken string

	ForceSendFields []string
}

func listAppsResponseToPb(st *ListAppsResponse) (*listAppsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAppsResponsePb{}

	var appsPb []appPb
	for _, item := range st.Apps {
		itemPb, err := appToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			appsPb = append(appsPb, *itemPb)
		}
	}
	pb.Apps = appsPb

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listAppsResponsePb struct {
	Apps []appPb `json:"apps,omitempty"`
	// Pagination token to request the next page of apps.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAppsResponseFromPb(pb *listAppsResponsePb) (*ListAppsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAppsResponse{}

	var appsField []App
	for _, itemPb := range pb.Apps {
		item, err := appFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			appsField = append(appsField, *item)
		}
	}
	st.Apps = appsField
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAppsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAppsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StartAppRequest struct {
	// The name of the app.
	Name string
}

func startAppRequestToPb(st *StartAppRequest) (*startAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startAppRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	return pb, nil
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

type startAppRequestPb struct {
	// The name of the app.
	Name string `json:"-" url:"-"`
}

func startAppRequestFromPb(pb *startAppRequestPb) (*StartAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartAppRequest{}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

type StopAppRequest struct {
	// The name of the app.
	Name string
}

func stopAppRequestToPb(st *StopAppRequest) (*stopAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stopAppRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	return pb, nil
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

type stopAppRequestPb struct {
	// The name of the app.
	Name string `json:"-" url:"-"`
}

func stopAppRequestFromPb(pb *stopAppRequestPb) (*StopAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StopAppRequest{}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

// Update an app
type UpdateAppRequest struct {
	App App
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name string
}

func updateAppRequestToPb(st *UpdateAppRequest) (*updateAppRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAppRequestPb{}
	appPb, err := appToPb(&st.App)
	if err != nil {
		return nil, err
	}
	if appPb != nil {
		pb.App = *appPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	return pb, nil
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

type updateAppRequestPb struct {
	App appPb `json:"app"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name string `json:"-" url:"-"`
}

func updateAppRequestFromPb(pb *updateAppRequestPb) (*UpdateAppRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAppRequest{}
	appField, err := appFromPb(&pb.App)
	if err != nil {
		return nil, err
	}
	if appField != nil {
		st.App = *appField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}
