// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// all definitions in this file are in alphabetical order

type AppEvents struct {
	EventName string `json:"event_name,omitempty"`

	EventTime string `json:"event_time,omitempty"`

	EventType string `json:"event_type,omitempty"`

	Message string `json:"message,omitempty"`

	ServiceName string `json:"service_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *AppEvents) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppEvents) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppManifest struct {
	// Workspace dependencies.
	Dependencies []any `json:"dependencies,omitempty"`
	// application description
	Description string `json:"description,omitempty"`
	// Ingress rules for app public endpoints
	Ingress any `json:"ingress,omitempty"`
	// Only a-z and dashes (-). Max length of 30.
	Name string `json:"name,omitempty"`
	// Container private registry
	Registry any `json:"registry,omitempty"`
	// list of app services. Restricted to one for now.
	Services any `json:"services,omitempty"`
	// The manifest format version. Must be set to 1.
	Version any `json:"version,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *AppManifest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppManifest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppServiceStatus struct {
	Deployment any `json:"deployment,omitempty"`

	Name string `json:"name,omitempty"`

	Template any `json:"template,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *AppServiceStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppServiceStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Retrieve the logs associated with building the model's environment for a
// given serving endpoint's served model.
type BuildLogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `json:"-" url:"-"`
	// The name of the served model that build logs will be retrieved for. This
	// field is required.
	ServedModelName string `json:"-" url:"-"`
}

type BuildLogsResponse struct {
	// The logs associated with building the served model's environment.
	Logs string `json:"logs"`
}

type CreateServingEndpoint struct {
	// The core config of the serving endpoint.
	Config EndpointCoreConfigInput `json:"config"`
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	Name string `json:"name"`
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	Tags []EndpointTag `json:"tags,omitempty"`
}

type DataframeSplitInput struct {
	Columns []any `json:"columns,omitempty"`

	Data []any `json:"data,omitempty"`

	Index []int `json:"index,omitempty"`
}

// Delete an application
type DeleteAppRequest struct {
	// The name of an application. This field is required.
	Name string `json:"-" url:"-"`
}

type DeleteAppResponse struct {
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteAppResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAppResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a serving endpoint
type DeleteServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name string `json:"-" url:"-"`
}

type DeployAppRequest struct {
	// Manifest that specifies the application requirements
	Manifest AppManifest `json:"manifest"`
	// Information passed at app deployment time to fulfill app dependencies
	Resources any `json:"resources,omitempty"`
}

type DeploymentStatus struct {
	// Container logs.
	ContainerLogs []any `json:"container_logs,omitempty"`
	// description
	DeploymentId string `json:"deployment_id,omitempty"`
	// Supplementary information about pod
	ExtraInfo string `json:"extra_info,omitempty"`
	// State: one of DEPLOYING,SUCCESS, FAILURE, DEPLOYMENT_STATE_UNSPECIFIED
	State DeploymentStatusState `json:"state,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DeploymentStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeploymentStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// State: one of DEPLOYING,SUCCESS, FAILURE, DEPLOYMENT_STATE_UNSPECIFIED
type DeploymentStatusState string

const DeploymentStatusStateDeploying DeploymentStatusState = `DEPLOYING`

const DeploymentStatusStateDeploymentStateUnspecified DeploymentStatusState = `DEPLOYMENT_STATE_UNSPECIFIED`

const DeploymentStatusStateFailure DeploymentStatusState = `FAILURE`

const DeploymentStatusStateSuccess DeploymentStatusState = `SUCCESS`

// String representation for [fmt.Print]
func (f *DeploymentStatusState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeploymentStatusState) Set(v string) error {
	switch v {
	case `DEPLOYING`, `DEPLOYMENT_STATE_UNSPECIFIED`, `FAILURE`, `SUCCESS`:
		*f = DeploymentStatusState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEPLOYING", "DEPLOYMENT_STATE_UNSPECIFIED", "FAILURE", "SUCCESS"`, v)
	}
}

// Type always returns DeploymentStatusState to satisfy [pflag.Value] interface
func (f *DeploymentStatusState) Type() string {
	return "DeploymentStatusState"
}

type EndpointCoreConfigInput struct {
	// The name of the serving endpoint to update. This field is required.
	Name string `json:"-" url:"-"`
	// A list of served models for the endpoint to serve. A serving endpoint can
	// have up to 10 served models.
	ServedModels []ServedModelInput `json:"served_models"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`
}

type EndpointCoreConfigOutput struct {
	// The config version that the serving endpoint is currently serving.
	ConfigVersion int `json:"config_version,omitempty"`
	// The list of served models under the serving endpoint config.
	ServedModels []ServedModelOutput `json:"served_models,omitempty"`
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *EndpointCoreConfigOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointCoreConfigOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointCoreConfigSummary struct {
	// The list of served models under the serving endpoint config.
	ServedModels []ServedModelSpec `json:"served_models,omitempty"`
}

type EndpointPendingConfig struct {
	// The config version that the serving endpoint is currently serving.
	ConfigVersion int `json:"config_version,omitempty"`
	// The list of served models belonging to the last issued update to the
	// serving endpoint.
	ServedModels []ServedModelOutput `json:"served_models,omitempty"`
	// The timestamp when the update to the pending config started.
	StartTime int64 `json:"start_time,omitempty"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *EndpointPendingConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointPendingConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointState struct {
	// The state of an endpoint's config update. This informs the user if the
	// pending_config is in progress, if the update failed, or if there is no
	// update in progress. Note that if the endpoint's config_update state value
	// is IN_PROGRESS, another update can not be made until the update completes
	// or fails."
	ConfigUpdate EndpointStateConfigUpdate `json:"config_update,omitempty"`
	// The state of an endpoint, indicating whether or not the endpoint is
	// queryable. An endpoint is READY if all of the served models in its active
	// configuration are ready. If any of the actively served models are in a
	// non-ready state, the endpoint state will be NOT_READY.
	Ready EndpointStateReady `json:"ready,omitempty"`
}

// The state of an endpoint's config update. This informs the user if the
// pending_config is in progress, if the update failed, or if there is no update
// in progress. Note that if the endpoint's config_update state value is
// IN_PROGRESS, another update can not be made until the update completes or
// fails."
type EndpointStateConfigUpdate string

const EndpointStateConfigUpdateInProgress EndpointStateConfigUpdate = `IN_PROGRESS`

const EndpointStateConfigUpdateNotUpdating EndpointStateConfigUpdate = `NOT_UPDATING`

const EndpointStateConfigUpdateUpdateFailed EndpointStateConfigUpdate = `UPDATE_FAILED`

// String representation for [fmt.Print]
func (f *EndpointStateConfigUpdate) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointStateConfigUpdate) Set(v string) error {
	switch v {
	case `IN_PROGRESS`, `NOT_UPDATING`, `UPDATE_FAILED`:
		*f = EndpointStateConfigUpdate(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IN_PROGRESS", "NOT_UPDATING", "UPDATE_FAILED"`, v)
	}
}

// Type always returns EndpointStateConfigUpdate to satisfy [pflag.Value] interface
func (f *EndpointStateConfigUpdate) Type() string {
	return "EndpointStateConfigUpdate"
}

// The state of an endpoint, indicating whether or not the endpoint is
// queryable. An endpoint is READY if all of the served models in its active
// configuration are ready. If any of the actively served models are in a
// non-ready state, the endpoint state will be NOT_READY.
type EndpointStateReady string

const EndpointStateReadyNotReady EndpointStateReady = `NOT_READY`

const EndpointStateReadyReady EndpointStateReady = `READY`

// String representation for [fmt.Print]
func (f *EndpointStateReady) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointStateReady) Set(v string) error {
	switch v {
	case `NOT_READY`, `READY`:
		*f = EndpointStateReady(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "NOT_READY", "READY"`, v)
	}
}

// Type always returns EndpointStateReady to satisfy [pflag.Value] interface
func (f *EndpointStateReady) Type() string {
	return "EndpointStateReady"
}

type EndpointTag struct {
	// Key field for a serving endpoint tag.
	Key string `json:"key"`
	// Optional value field for a serving endpoint tag.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *EndpointTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Retrieve the metrics associated with a serving endpoint
type ExportMetricsRequest struct {
	// The name of the serving endpoint to retrieve metrics for. This field is
	// required.
	Name string `json:"-" url:"-"`
}

// Get deployment status for an application
type GetAppDeploymentStatusRequest struct {
	// The deployment id for an application. This field is required.
	DeploymentId string `json:"-" url:"-"`
	// Boolean flag to include application logs
	IncludeAppLog string `json:"-" url:"include_app_log,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetAppDeploymentStatusRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetAppDeploymentStatusRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get definition for an application
type GetAppRequest struct {
	// The name of an application. This field is required.
	Name string `json:"-" url:"-"`
}

type GetAppResponse struct {
	CurrentServices []AppServiceStatus `json:"current_services,omitempty"`

	Name string `json:"name,omitempty"`

	PendingServices []AppServiceStatus `json:"pending_services,omitempty"`

	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetAppResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetAppResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get deployment events for an application
type GetEventsRequest struct {
	// The name of an application. This field is required.
	Name string `json:"-" url:"-"`
}

// Get serving endpoint permission levels
type GetServingEndpointPermissionLevelsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `json:"-" url:"-"`
}

type GetServingEndpointPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ServingEndpointPermissionsDescription `json:"permission_levels,omitempty"`
}

// Get serving endpoint permissions
type GetServingEndpointPermissionsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `json:"-" url:"-"`
}

// Get a single serving endpoint
type GetServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name string `json:"-" url:"-"`
}

type ListAppEventsResponse struct {
	// App events
	Events []AppEvents `json:"events,omitempty"`
}

type ListAppsResponse struct {
	// Available apps.
	Apps []any `json:"apps,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListAppsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAppsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEndpointsResponse struct {
	// The list of endpoints.
	Endpoints []ServingEndpoint `json:"endpoints,omitempty"`
}

// Retrieve the most recent log lines associated with a given serving endpoint's
// served model
type LogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `json:"-" url:"-"`
	// The name of the served model that logs will be retrieved for. This field
	// is required.
	ServedModelName string `json:"-" url:"-"`
}

type PatchServingEndpointTags struct {
	// List of endpoint tags to add
	AddTags []EndpointTag `json:"add_tags,omitempty"`
	// List of tag keys to delete
	DeleteTags []string `json:"delete_tags,omitempty"`
	// The name of the serving endpoint who's tags to patch. This field is
	// required.
	Name string `json:"-" url:"-"`
}

type QueryEndpointInput struct {
	// Pandas Dataframe input in the records orientation.
	DataframeRecords []any `json:"dataframe_records,omitempty"`
	// Pandas Dataframe input in the split orientation.
	DataframeSplit *DataframeSplitInput `json:"dataframe_split,omitempty"`
	// Tensor-based input in columnar format.
	Inputs any `json:"inputs,omitempty"`
	// Tensor-based input in row format.
	Instances []any `json:"instances,omitempty"`
	// The name of the serving endpoint. This field is required.
	Name string `json:"-" url:"-"`
}

type QueryEndpointResponse struct {
	// The predictions returned by the serving endpoint.
	Predictions []any `json:"predictions"`
}

type Route struct {
	// The name of the served model this route configures traffic for.
	ServedModelName string `json:"served_model_name"`
	// The percentage of endpoint traffic to send to this route. It must be an
	// integer between 0 and 100 inclusive.
	TrafficPercentage int `json:"traffic_percentage"`
}

type ServedModelInput struct {
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this model. Note: this is an
	// experimental feature and subject to change. Example model environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// ARN of the instance profile that the served model will use to access AWS
	// resources.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The name of the model in Databricks Model Registry to be served or if the
	// model resides in Unity Catalog, the full name of model, in the form of
	// __catalog_name__.__schema_name__.__model_name__.
	ModelName string `json:"model_name"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion string `json:"model_version"`
	// The name of a served model. It must be unique across an endpoint. If not
	// specified, this field will default to <model-name>-<model-version>. A
	// served model name can consist of alphanumeric characters, dashes, and
	// underscores.
	Name string `json:"name,omitempty"`
	// Whether the compute resources for the served model should scale down to
	// zero.
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled"`
	// The workload size of the served model. The workload size corresponds to a
	// range of provisioned concurrency that the compute will autoscale between.
	// A single unit of provisioned concurrency can process one request at a
	// time. Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). If scale-to-zero is enabled, the lower bound of
	// the provisioned concurrency for each workload size will be 0.
	WorkloadSize string `json:"workload_size"`
	// The workload type of the served model. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See
	// documentation for all options.
	WorkloadType string `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServedModelInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedModelOutput struct {
	// The creation timestamp of the served model in Unix time.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the served model.
	Creator string `json:"creator,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this model. Note: this is an
	// experimental feature and subject to change. Example model environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// ARN of the instance profile that the served model will use to access AWS
	// resources.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The name of the model in Databricks Model Registry or the full name of
	// the model in Unity Catalog.
	ModelName string `json:"model_name,omitempty"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion string `json:"model_version,omitempty"`
	// The name of the served model.
	Name string `json:"name,omitempty"`
	// Whether the compute resources for the Served Model should scale down to
	// zero.
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`
	// Information corresponding to the state of the Served Model.
	State *ServedModelState `json:"state,omitempty"`
	// The workload size of the served model. The workload size corresponds to a
	// range of provisioned concurrency that the compute will autoscale between.
	// A single unit of provisioned concurrency can process one request at a
	// time. Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). If scale-to-zero is enabled, the lower bound of
	// the provisioned concurrency for each workload size will be 0.
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served model. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See
	// documentation for all options.
	WorkloadType string `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServedModelOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedModelSpec struct {
	// The name of the model in Databricks Model Registry or the full name of
	// the model in Unity Catalog.
	ModelName string `json:"model_name,omitempty"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion string `json:"model_version,omitempty"`
	// The name of the served model.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServedModelSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedModelState struct {
	// The state of the served model deployment. DEPLOYMENT_CREATING indicates
	// that the served model is not ready yet because the deployment is still
	// being created (i.e container image is building, model server is deploying
	// for the first time, etc.). DEPLOYMENT_RECOVERING indicates that the
	// served model was previously in a ready state but no longer is and is
	// attempting to recover. DEPLOYMENT_READY indicates that the served model
	// is ready to receive traffic. DEPLOYMENT_FAILED indicates that there was
	// an error trying to bring up the served model (e.g container image build
	// failed, the model server failed to start due to a model loading error,
	// etc.) DEPLOYMENT_ABORTED indicates that the deployment was terminated
	// likely due to a failure in bringing up another served model under the
	// same endpoint and config version.
	Deployment ServedModelStateDeployment `json:"deployment,omitempty"`
	// More information about the state of the served model, if available.
	DeploymentStateMessage string `json:"deployment_state_message,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServedModelState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelState) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The state of the served model deployment. DEPLOYMENT_CREATING indicates that
// the served model is not ready yet because the deployment is still being
// created (i.e container image is building, model server is deploying for the
// first time, etc.). DEPLOYMENT_RECOVERING indicates that the served model was
// previously in a ready state but no longer is and is attempting to recover.
// DEPLOYMENT_READY indicates that the served model is ready to receive traffic.
// DEPLOYMENT_FAILED indicates that there was an error trying to bring up the
// served model (e.g container image build failed, the model server failed to
// start due to a model loading error, etc.) DEPLOYMENT_ABORTED indicates that
// the deployment was terminated likely due to a failure in bringing up another
// served model under the same endpoint and config version.
type ServedModelStateDeployment string

const ServedModelStateDeploymentDeploymentAborted ServedModelStateDeployment = `DEPLOYMENT_ABORTED`

const ServedModelStateDeploymentDeploymentCreating ServedModelStateDeployment = `DEPLOYMENT_CREATING`

const ServedModelStateDeploymentDeploymentFailed ServedModelStateDeployment = `DEPLOYMENT_FAILED`

const ServedModelStateDeploymentDeploymentReady ServedModelStateDeployment = `DEPLOYMENT_READY`

const ServedModelStateDeploymentDeploymentRecovering ServedModelStateDeployment = `DEPLOYMENT_RECOVERING`

// String representation for [fmt.Print]
func (f *ServedModelStateDeployment) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServedModelStateDeployment) Set(v string) error {
	switch v {
	case `DEPLOYMENT_ABORTED`, `DEPLOYMENT_CREATING`, `DEPLOYMENT_FAILED`, `DEPLOYMENT_READY`, `DEPLOYMENT_RECOVERING`:
		*f = ServedModelStateDeployment(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEPLOYMENT_ABORTED", "DEPLOYMENT_CREATING", "DEPLOYMENT_FAILED", "DEPLOYMENT_READY", "DEPLOYMENT_RECOVERING"`, v)
	}
}

// Type always returns ServedModelStateDeployment to satisfy [pflag.Value] interface
func (f *ServedModelStateDeployment) Type() string {
	return "ServedModelStateDeployment"
}

type ServerLogsResponse struct {
	// The most recent log lines of the model server processing invocation
	// requests.
	Logs string `json:"logs"`
}

type ServingEndpoint struct {
	// The config that is currently being served by the endpoint.
	Config *EndpointCoreConfigSummary `json:"config,omitempty"`
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the serving endpoint.
	Creator string `json:"creator,omitempty"`
	// System-generated ID of the endpoint. This is used to refer to the
	// endpoint in the Permissions API
	Id string `json:"id,omitempty"`
	// The timestamp when the endpoint was last updated by a user in Unix time.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The name of the serving endpoint.
	Name string `json:"name,omitempty"`
	// Information corresponding to the state of the serving endpoint.
	State *EndpointState `json:"state,omitempty"`
	// Tags attached to the serving endpoint.
	Tags []EndpointTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServingEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`
	// name of the service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServingEndpointAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointAccessControlResponse struct {
	// All permissions.
	AllPermissions []ServingEndpointPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServingEndpointAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointDetailed struct {
	// The config that is currently being served by the endpoint.
	Config *EndpointCoreConfigOutput `json:"config,omitempty"`
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the serving endpoint.
	Creator string `json:"creator,omitempty"`
	// System-generated ID of the endpoint. This is used to refer to the
	// endpoint in the Permissions API
	Id string `json:"id,omitempty"`
	// The timestamp when the endpoint was last updated by a user in Unix time.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The name of the serving endpoint.
	Name string `json:"name,omitempty"`
	// The config that the endpoint is attempting to update to.
	PendingConfig *EndpointPendingConfig `json:"pending_config,omitempty"`
	// The permission level of the principal making the request.
	PermissionLevel ServingEndpointDetailedPermissionLevel `json:"permission_level,omitempty"`
	// Information corresponding to the state of the serving endpoint.
	State *EndpointState `json:"state,omitempty"`
	// Tags attached to the serving endpoint.
	Tags []EndpointTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServingEndpointDetailed) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointDetailed) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The permission level of the principal making the request.
type ServingEndpointDetailedPermissionLevel string

const ServingEndpointDetailedPermissionLevelCanManage ServingEndpointDetailedPermissionLevel = `CAN_MANAGE`

const ServingEndpointDetailedPermissionLevelCanQuery ServingEndpointDetailedPermissionLevel = `CAN_QUERY`

const ServingEndpointDetailedPermissionLevelCanView ServingEndpointDetailedPermissionLevel = `CAN_VIEW`

// String representation for [fmt.Print]
func (f *ServingEndpointDetailedPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServingEndpointDetailedPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_QUERY`, `CAN_VIEW`:
		*f = ServingEndpointDetailedPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW"`, v)
	}
}

// Type always returns ServingEndpointDetailedPermissionLevel to satisfy [pflag.Value] interface
func (f *ServingEndpointDetailedPermissionLevel) Type() string {
	return "ServingEndpointDetailedPermissionLevel"
}

type ServingEndpointPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServingEndpointPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type ServingEndpointPermissionLevel string

const ServingEndpointPermissionLevelCanManage ServingEndpointPermissionLevel = `CAN_MANAGE`

const ServingEndpointPermissionLevelCanQuery ServingEndpointPermissionLevel = `CAN_QUERY`

const ServingEndpointPermissionLevelCanView ServingEndpointPermissionLevel = `CAN_VIEW`

// String representation for [fmt.Print]
func (f *ServingEndpointPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServingEndpointPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_QUERY`, `CAN_VIEW`:
		*f = ServingEndpointPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW"`, v)
	}
}

// Type always returns ServingEndpointPermissionLevel to satisfy [pflag.Value] interface
func (f *ServingEndpointPermissionLevel) Type() string {
	return "ServingEndpointPermissionLevel"
}

type ServingEndpointPermissions struct {
	AccessControlList []ServingEndpointAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServingEndpointPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointPermissionsDescription struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServingEndpointPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointPermissionsRequest struct {
	AccessControlList []ServingEndpointAccessControlRequest `json:"access_control_list,omitempty"`
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `json:"-" url:"-"`
}

type TrafficConfig struct {
	// The list of routes that define traffic to each served model.
	Routes []Route `json:"routes,omitempty"`
}
