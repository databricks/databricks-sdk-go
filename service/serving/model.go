// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import "fmt"

// all definitions in this file are in alphabetical order

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
}

// Delete a serving endpoint
type DeleteServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name string `json:"-" url:"-"`
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

// Retrieve the metrics corresponding to a serving endpoint for the current time
// in Prometheus or OpenMetrics exposition format
type ExportMetricsRequest struct {
	// The name of the serving endpoint to retrieve metrics for. This field is
	// required.
	Name string `json:"-" url:"-"`
}

// Get a single serving endpoint
type GetServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name string `json:"-" url:"-"`
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

type QueryEndpointResponse struct {
	// The predictions returned by the serving endpoint.
	Predictions []any `json:"predictions"`
}

// Query a serving endpoint with provided model input.
type QueryRequest struct {
	// The name of the serving endpoint. This field is required.
	Name string `json:"-" url:"-"`
}

type Route struct {
	// The name of the served model this route configures traffic for.
	ServedModelName string `json:"served_model_name"`
	// The percentage of endpoint traffic to send to this route. It must be an
	// integer between 0 and 100 inclusive.
	TrafficPercentage int `json:"traffic_percentage"`
}

type ServedModelInput struct {
	// The name of the model in Databricks Model Registry to be served.
	ModelName string `json:"model_name"`
	// The version of the model in Databricks Model Registry to be served.
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
}

type ServedModelOutput struct {
	// The creation timestamp of the served model in Unix time.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the served model.
	Creator string `json:"creator,omitempty"`
	// The name of the model in Databricks Model Registry.
	ModelName string `json:"model_name,omitempty"`
	// The version of the model in Databricks Model Registry.
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
}

type ServedModelSpec struct {
	// The name of the model in Databricks Model Registry.
	ModelName string `json:"model_name,omitempty"`
	// The version of the model in Databricks Model Registry.
	ModelVersion string `json:"model_version,omitempty"`
	// The name of the served model.
	Name string `json:"name,omitempty"`
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

type TrafficConfig struct {
	// The list of routes that define traffic to each served model.
	Routes []Route `json:"routes,omitempty"`
}
