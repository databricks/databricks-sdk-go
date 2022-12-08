// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package warehouses

import "fmt"

// all definitions in this file are in alphabetical order

type Channel struct {
	DbsqlVersion string `json:"dbsql_version,omitempty"`

	Name ChannelName `json:"name,omitempty"`
}

// Channel information for the SQL warehouse at the time of query execution
type ChannelInfo struct {
	// DBSQL Version the channel is mapped to
	DbsqlVersion string `json:"dbsql_version,omitempty"`
	// Name of the channel
	Name ChannelName `json:"name,omitempty"`
}

// Name of the channel
type ChannelName string

const ChannelNameChannelNameCurrent ChannelName = `CHANNEL_NAME_CURRENT`

const ChannelNameChannelNameCustom ChannelName = `CHANNEL_NAME_CUSTOM`

const ChannelNameChannelNamePreview ChannelName = `CHANNEL_NAME_PREVIEW`

const ChannelNameChannelNamePrevious ChannelName = `CHANNEL_NAME_PREVIOUS`

const ChannelNameChannelNameUnspecified ChannelName = `CHANNEL_NAME_UNSPECIFIED`

// String representation for [fmt.Print]
func (cn *ChannelName) String() string {
	return string(*cn)
}

// Set raw string value and validate it against allowed values
func (cn *ChannelName) Set(v string) error {
	switch v {
	case `CHANNEL_NAME_CURRENT`, `CHANNEL_NAME_CUSTOM`, `CHANNEL_NAME_PREVIEW`, `CHANNEL_NAME_PREVIOUS`, `CHANNEL_NAME_UNSPECIFIED`:
		*cn = ChannelName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CHANNEL_NAME_CURRENT", "CHANNEL_NAME_CUSTOM", "CHANNEL_NAME_PREVIEW", "CHANNEL_NAME_PREVIOUS", "CHANNEL_NAME_UNSPECIFIED"`, v)
	}
}

// Type always returns ChannelName to satisfy [pflag.Value] interface
func (cn *ChannelName) Type() string {
	return "ChannelName"
}

type CreateWarehouseRequest struct {
	// The amount of time in minutes that a SQL Endpoint must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel Details
	Channel *Channel `json:"channel,omitempty"`
	// Size of the clusters allocated for this endpoint. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize string `json:"cluster_size,omitempty"`
	// endpoint creator name
	CreatorName string `json:"creator_name,omitempty"`
	// Configures whether the endpoint should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the endpoint should use Serverless Compute (aka
	// Nephos)
	//
	// Defaults to value in global endpoint settings
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters int `json:"max_num_clusters,omitempty"`
	// Minimum number of available clusters that will be maintained for this SQL
	// Endpoint. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name string `json:"name,omitempty"`
	// Configurations whether the endpoint should use spot instances.
	//
	// Supported values: "COST_OPTIMIZED", "RELIABILITY_OPTIMIZED"
	//
	// Defaults to COST_OPTIMIZED.
	//
	// Please refer to documentation for EndpointSpotInstancePolicy for more
	// details.
	SpotInstancePolicy CreateWarehouseRequestSpotInstancePolicy `json:"spot_instance_policy,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL Endpoints.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags `json:"tags,omitempty"`
	// Warehouse type (Classic/Pro)
	WarehouseType CreateWarehouseRequestWarehouseType `json:"warehouse_type,omitempty"`
}

// Configurations whether the endpoint should use spot instances.
//
// Supported values: "COST_OPTIMIZED", "RELIABILITY_OPTIMIZED"
//
// Defaults to COST_OPTIMIZED.
//
// Please refer to documentation for EndpointSpotInstancePolicy for more
// details.
type CreateWarehouseRequestSpotInstancePolicy string

const CreateWarehouseRequestSpotInstancePolicyCostOptimized CreateWarehouseRequestSpotInstancePolicy = `COST_OPTIMIZED`

const CreateWarehouseRequestSpotInstancePolicyPolicyUnspecified CreateWarehouseRequestSpotInstancePolicy = `POLICY_UNSPECIFIED`

const CreateWarehouseRequestSpotInstancePolicyReliabilityOptimized CreateWarehouseRequestSpotInstancePolicy = `RELIABILITY_OPTIMIZED`

// String representation for [fmt.Print]
func (cwrsip *CreateWarehouseRequestSpotInstancePolicy) String() string {
	return string(*cwrsip)
}

// Set raw string value and validate it against allowed values
func (cwrsip *CreateWarehouseRequestSpotInstancePolicy) Set(v string) error {
	switch v {
	case `COST_OPTIMIZED`, `POLICY_UNSPECIFIED`, `RELIABILITY_OPTIMIZED`:
		*cwrsip = CreateWarehouseRequestSpotInstancePolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COST_OPTIMIZED", "POLICY_UNSPECIFIED", "RELIABILITY_OPTIMIZED"`, v)
	}
}

// Type always returns CreateWarehouseRequestSpotInstancePolicy to satisfy [pflag.Value] interface
func (cwrsip *CreateWarehouseRequestSpotInstancePolicy) Type() string {
	return "CreateWarehouseRequestSpotInstancePolicy"
}

// Warehouse type (Classic/Pro)
type CreateWarehouseRequestWarehouseType string

const CreateWarehouseRequestWarehouseTypeClassic CreateWarehouseRequestWarehouseType = `CLASSIC`

const CreateWarehouseRequestWarehouseTypePro CreateWarehouseRequestWarehouseType = `PRO`

const CreateWarehouseRequestWarehouseTypeTypeUnspecified CreateWarehouseRequestWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (cwrwt *CreateWarehouseRequestWarehouseType) String() string {
	return string(*cwrwt)
}

// Set raw string value and validate it against allowed values
func (cwrwt *CreateWarehouseRequestWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `TYPE_UNSPECIFIED`:
		*cwrwt = CreateWarehouseRequestWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "TYPE_UNSPECIFIED"`, v)
	}
}

// Type always returns CreateWarehouseRequestWarehouseType to satisfy [pflag.Value] interface
func (cwrwt *CreateWarehouseRequestWarehouseType) Type() string {
	return "CreateWarehouseRequestWarehouseType"
}

type CreateWarehouseResponse struct {
	// Id for the SQL warehouse. This value is unique across all SQL warehouses.
	Id string `json:"id,omitempty"`
}

// Delete a warehouse
type DeleteWarehouse struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" url:"-"`
}

type EditWarehouseRequest struct {
	// The amount of time in minutes that a SQL Endpoint must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel Details
	Channel *Channel `json:"channel,omitempty"`
	// Size of the clusters allocated for this endpoint. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize string `json:"cluster_size,omitempty"`
	// endpoint creator name
	CreatorName string `json:"creator_name,omitempty"`
	// Configures whether the endpoint should use Databricks Compute (aka
	// Nephos)
	//
	// Deprecated: Use enable_serverless_compute TODO(SC-79930): Remove the
	// field once clients are updated
	EnableDatabricksCompute bool `json:"enable_databricks_compute,omitempty"`
	// Configures whether the endpoint should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the endpoint should use Serverless Compute (aka
	// Nephos)
	//
	// Defaults to value in global endpoint settings
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Required. Id of the warehouse to configure.
	Id string `json:"-" url:"-"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters int `json:"max_num_clusters,omitempty"`
	// Minimum number of available clusters that will be maintained for this SQL
	// Endpoint. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name string `json:"name,omitempty"`
	// Configurations whether the endpoint should use spot instances.
	//
	// Supported values: "COST_OPTIMIZED", "RELIABILITY_OPTIMIZED"
	//
	// Defaults to COST_OPTIMIZED.
	//
	// Please refer to documentation for EndpointSpotInstancePolicy for more
	// details.
	SpotInstancePolicy EditWarehouseRequestSpotInstancePolicy `json:"spot_instance_policy,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL Endpoints.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags `json:"tags,omitempty"`
	// Warehouse type (Classic/Pro)
	WarehouseType EditWarehouseRequestWarehouseType `json:"warehouse_type,omitempty"`
}

// Configurations whether the endpoint should use spot instances.
//
// Supported values: "COST_OPTIMIZED", "RELIABILITY_OPTIMIZED"
//
// Defaults to COST_OPTIMIZED.
//
// Please refer to documentation for EndpointSpotInstancePolicy for more
// details.
type EditWarehouseRequestSpotInstancePolicy string

const EditWarehouseRequestSpotInstancePolicyCostOptimized EditWarehouseRequestSpotInstancePolicy = `COST_OPTIMIZED`

const EditWarehouseRequestSpotInstancePolicyPolicyUnspecified EditWarehouseRequestSpotInstancePolicy = `POLICY_UNSPECIFIED`

const EditWarehouseRequestSpotInstancePolicyReliabilityOptimized EditWarehouseRequestSpotInstancePolicy = `RELIABILITY_OPTIMIZED`

// String representation for [fmt.Print]
func (ewrsip *EditWarehouseRequestSpotInstancePolicy) String() string {
	return string(*ewrsip)
}

// Set raw string value and validate it against allowed values
func (ewrsip *EditWarehouseRequestSpotInstancePolicy) Set(v string) error {
	switch v {
	case `COST_OPTIMIZED`, `POLICY_UNSPECIFIED`, `RELIABILITY_OPTIMIZED`:
		*ewrsip = EditWarehouseRequestSpotInstancePolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COST_OPTIMIZED", "POLICY_UNSPECIFIED", "RELIABILITY_OPTIMIZED"`, v)
	}
}

// Type always returns EditWarehouseRequestSpotInstancePolicy to satisfy [pflag.Value] interface
func (ewrsip *EditWarehouseRequestSpotInstancePolicy) Type() string {
	return "EditWarehouseRequestSpotInstancePolicy"
}

// Warehouse type (Classic/Pro)
type EditWarehouseRequestWarehouseType string

const EditWarehouseRequestWarehouseTypeClassic EditWarehouseRequestWarehouseType = `CLASSIC`

const EditWarehouseRequestWarehouseTypePro EditWarehouseRequestWarehouseType = `PRO`

const EditWarehouseRequestWarehouseTypeTypeUnspecified EditWarehouseRequestWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (ewrwt *EditWarehouseRequestWarehouseType) String() string {
	return string(*ewrwt)
}

// Set raw string value and validate it against allowed values
func (ewrwt *EditWarehouseRequestWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `TYPE_UNSPECIFIED`:
		*ewrwt = EditWarehouseRequestWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "TYPE_UNSPECIFIED"`, v)
	}
}

// Type always returns EditWarehouseRequestWarehouseType to satisfy [pflag.Value] interface
func (ewrwt *EditWarehouseRequestWarehouseType) Type() string {
	return "EditWarehouseRequestWarehouseType"
}

type EndpointConfPair struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

type EndpointHealth struct {
	// Details about errors that are causing current degraded/failed status.
	Details string `json:"details,omitempty"`
	// The reason for failure to bring up clusters for this endpoint. This is
	// available when status is 'FAILED' and sometimes when it is DEGRADED.
	FailureReason *TerminationReason `json:"failure_reason,omitempty"`
	// Deprecated. split into summary and details for security
	Message string `json:"message,omitempty"`
	// Health status of the endpoint.
	Status EndpointHealthStatus `json:"status,omitempty"`
	// A short summary of the health status in case of degraded/failed
	// endpoints.
	Summary string `json:"summary,omitempty"`
}

// Health status of the endpoint.
type EndpointHealthStatus string

const EndpointHealthStatusDegraded EndpointHealthStatus = `DEGRADED`

const EndpointHealthStatusFailed EndpointHealthStatus = `FAILED`

const EndpointHealthStatusHealthy EndpointHealthStatus = `HEALTHY`

const EndpointHealthStatusStatusUnspecified EndpointHealthStatus = `STATUS_UNSPECIFIED`

// String representation for [fmt.Print]
func (ehs *EndpointHealthStatus) String() string {
	return string(*ehs)
}

// Set raw string value and validate it against allowed values
func (ehs *EndpointHealthStatus) Set(v string) error {
	switch v {
	case `DEGRADED`, `FAILED`, `HEALTHY`, `STATUS_UNSPECIFIED`:
		*ehs = EndpointHealthStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEGRADED", "FAILED", "HEALTHY", "STATUS_UNSPECIFIED"`, v)
	}
}

// Type always returns EndpointHealthStatus to satisfy [pflag.Value] interface
func (ehs *EndpointHealthStatus) Type() string {
	return "EndpointHealthStatus"
}

type EndpointInfo struct {
	// The amount of time in minutes that a SQL Endpoint must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel Details
	Channel *Channel `json:"channel,omitempty"`
	// Size of the clusters allocated for this endpoint. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize string `json:"cluster_size,omitempty"`
	// endpoint creator name
	CreatorName string `json:"creator_name,omitempty"`
	// Configures whether the endpoint should use Databricks Compute (aka
	// Nephos)
	//
	// Deprecated: Use enable_serverless_compute TODO(SC-79930): Remove the
	// field once clients are updated
	EnableDatabricksCompute bool `json:"enable_databricks_compute,omitempty"`
	// Configures whether the endpoint should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the endpoint should use Serverless Compute (aka
	// Nephos)
	//
	// Defaults to value in global endpoint settings
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Optional health status. Assume the endpoint is healthy if this field is
	// not set.
	Health *EndpointHealth `json:"health,omitempty"`
	// unique identifier for endpoint
	Id string `json:"id,omitempty"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// the jdbc connection string for this endpoint
	JdbcUrl string `json:"jdbc_url,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters int `json:"max_num_clusters,omitempty"`
	// Minimum number of available clusters that will be maintained for this SQL
	// Endpoint. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name string `json:"name,omitempty"`
	// current number of active sessions for the endpoint
	NumActiveSessions int64 `json:"num_active_sessions,omitempty"`
	// current number of clusters running for the service
	NumClusters int `json:"num_clusters,omitempty"`
	// ODBC parameters for the sql endpoint
	OdbcParams *OdbcParams `json:"odbc_params,omitempty"`
	// Configurations whether the endpoint should use spot instances.
	//
	// Supported values: "COST_OPTIMIZED", "RELIABILITY_OPTIMIZED"
	//
	// Defaults to COST_OPTIMIZED.
	//
	// Please refer to documentation for EndpointSpotInstancePolicy for more
	// details.
	SpotInstancePolicy EndpointInfoSpotInstancePolicy `json:"spot_instance_policy,omitempty"`
	// state of the endpoint
	State EndpointInfoState `json:"state,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL Endpoints.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags `json:"tags,omitempty"`
	// Warehouse type (Classic/Pro)
	WarehouseType EndpointInfoWarehouseType `json:"warehouse_type,omitempty"`
}

// Configurations whether the endpoint should use spot instances.
//
// Supported values: "COST_OPTIMIZED", "RELIABILITY_OPTIMIZED"
//
// Defaults to COST_OPTIMIZED.
//
// Please refer to documentation for EndpointSpotInstancePolicy for more
// details.
type EndpointInfoSpotInstancePolicy string

const EndpointInfoSpotInstancePolicyCostOptimized EndpointInfoSpotInstancePolicy = `COST_OPTIMIZED`

const EndpointInfoSpotInstancePolicyPolicyUnspecified EndpointInfoSpotInstancePolicy = `POLICY_UNSPECIFIED`

const EndpointInfoSpotInstancePolicyReliabilityOptimized EndpointInfoSpotInstancePolicy = `RELIABILITY_OPTIMIZED`

// String representation for [fmt.Print]
func (eisip *EndpointInfoSpotInstancePolicy) String() string {
	return string(*eisip)
}

// Set raw string value and validate it against allowed values
func (eisip *EndpointInfoSpotInstancePolicy) Set(v string) error {
	switch v {
	case `COST_OPTIMIZED`, `POLICY_UNSPECIFIED`, `RELIABILITY_OPTIMIZED`:
		*eisip = EndpointInfoSpotInstancePolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COST_OPTIMIZED", "POLICY_UNSPECIFIED", "RELIABILITY_OPTIMIZED"`, v)
	}
}

// Type always returns EndpointInfoSpotInstancePolicy to satisfy [pflag.Value] interface
func (eisip *EndpointInfoSpotInstancePolicy) Type() string {
	return "EndpointInfoSpotInstancePolicy"
}

// state of the endpoint
type EndpointInfoState string

const EndpointInfoStateDeleted EndpointInfoState = `DELETED`

const EndpointInfoStateDeleting EndpointInfoState = `DELETING`

const EndpointInfoStateRunning EndpointInfoState = `RUNNING`

const EndpointInfoStateStarting EndpointInfoState = `STARTING`

const EndpointInfoStateStopped EndpointInfoState = `STOPPED`

const EndpointInfoStateStopping EndpointInfoState = `STOPPING`

// String representation for [fmt.Print]
func (eis *EndpointInfoState) String() string {
	return string(*eis)
}

// Set raw string value and validate it against allowed values
func (eis *EndpointInfoState) Set(v string) error {
	switch v {
	case `DELETED`, `DELETING`, `RUNNING`, `STARTING`, `STOPPED`, `STOPPING`:
		*eis = EndpointInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETED", "DELETING", "RUNNING", "STARTING", "STOPPED", "STOPPING"`, v)
	}
}

// Type always returns EndpointInfoState to satisfy [pflag.Value] interface
func (eis *EndpointInfoState) Type() string {
	return "EndpointInfoState"
}

// Warehouse type (Classic/Pro)
type EndpointInfoWarehouseType string

const EndpointInfoWarehouseTypeClassic EndpointInfoWarehouseType = `CLASSIC`

const EndpointInfoWarehouseTypePro EndpointInfoWarehouseType = `PRO`

const EndpointInfoWarehouseTypeTypeUnspecified EndpointInfoWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (eiwt *EndpointInfoWarehouseType) String() string {
	return string(*eiwt)
}

// Set raw string value and validate it against allowed values
func (eiwt *EndpointInfoWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `TYPE_UNSPECIFIED`:
		*eiwt = EndpointInfoWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "TYPE_UNSPECIFIED"`, v)
	}
}

// Type always returns EndpointInfoWarehouseType to satisfy [pflag.Value] interface
func (eiwt *EndpointInfoWarehouseType) Type() string {
	return "EndpointInfoWarehouseType"
}

type EndpointTagPair struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

type EndpointTags struct {
	CustomTags []EndpointTagPair `json:"custom_tags,omitempty"`
}

// Get warehouse info
type GetWarehouse struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" url:"-"`
}

type GetWarehouseResponse struct {
	// The amount of time in minutes that a SQL Endpoint must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel Details
	Channel *Channel `json:"channel,omitempty"`
	// Size of the clusters allocated for this endpoint. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize string `json:"cluster_size,omitempty"`
	// endpoint creator name
	CreatorName string `json:"creator_name,omitempty"`
	// Configures whether the endpoint should use Databricks Compute (aka
	// Nephos)
	//
	// Deprecated: Use enable_serverless_compute TODO(SC-79930): Remove the
	// field once clients are updated
	EnableDatabricksCompute bool `json:"enable_databricks_compute,omitempty"`
	// Configures whether the endpoint should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the endpoint should use Serverless Compute (aka
	// Nephos)
	//
	// Defaults to value in global endpoint settings
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Optional health status. Assume the endpoint is healthy if this field is
	// not set.
	Health *EndpointHealth `json:"health,omitempty"`
	// unique identifier for endpoint
	Id string `json:"id,omitempty"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// the jdbc connection string for this endpoint
	JdbcUrl string `json:"jdbc_url,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters int `json:"max_num_clusters,omitempty"`
	// Minimum number of available clusters that will be maintained for this SQL
	// Endpoint. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name string `json:"name,omitempty"`
	// current number of active sessions for the endpoint
	NumActiveSessions int64 `json:"num_active_sessions,omitempty"`
	// current number of clusters running for the service
	NumClusters int `json:"num_clusters,omitempty"`
	// ODBC parameters for the sql endpoint
	OdbcParams *OdbcParams `json:"odbc_params,omitempty"`
	// Configurations whether the endpoint should use spot instances.
	//
	// Supported values: "COST_OPTIMIZED", "RELIABILITY_OPTIMIZED"
	//
	// Defaults to COST_OPTIMIZED.
	//
	// Please refer to documentation for EndpointSpotInstancePolicy for more
	// details.
	SpotInstancePolicy GetWarehouseResponseSpotInstancePolicy `json:"spot_instance_policy,omitempty"`
	// state of the endpoint
	State GetWarehouseResponseState `json:"state,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL Endpoints.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags `json:"tags,omitempty"`
	// Warehouse type (Classic/Pro)
	WarehouseType GetWarehouseResponseWarehouseType `json:"warehouse_type,omitempty"`
}

// Configurations whether the endpoint should use spot instances.
//
// Supported values: "COST_OPTIMIZED", "RELIABILITY_OPTIMIZED"
//
// Defaults to COST_OPTIMIZED.
//
// Please refer to documentation for EndpointSpotInstancePolicy for more
// details.
type GetWarehouseResponseSpotInstancePolicy string

const GetWarehouseResponseSpotInstancePolicyCostOptimized GetWarehouseResponseSpotInstancePolicy = `COST_OPTIMIZED`

const GetWarehouseResponseSpotInstancePolicyPolicyUnspecified GetWarehouseResponseSpotInstancePolicy = `POLICY_UNSPECIFIED`

const GetWarehouseResponseSpotInstancePolicyReliabilityOptimized GetWarehouseResponseSpotInstancePolicy = `RELIABILITY_OPTIMIZED`

// String representation for [fmt.Print]
func (gwrsip *GetWarehouseResponseSpotInstancePolicy) String() string {
	return string(*gwrsip)
}

// Set raw string value and validate it against allowed values
func (gwrsip *GetWarehouseResponseSpotInstancePolicy) Set(v string) error {
	switch v {
	case `COST_OPTIMIZED`, `POLICY_UNSPECIFIED`, `RELIABILITY_OPTIMIZED`:
		*gwrsip = GetWarehouseResponseSpotInstancePolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COST_OPTIMIZED", "POLICY_UNSPECIFIED", "RELIABILITY_OPTIMIZED"`, v)
	}
}

// Type always returns GetWarehouseResponseSpotInstancePolicy to satisfy [pflag.Value] interface
func (gwrsip *GetWarehouseResponseSpotInstancePolicy) Type() string {
	return "GetWarehouseResponseSpotInstancePolicy"
}

// state of the endpoint
type GetWarehouseResponseState string

const GetWarehouseResponseStateDeleted GetWarehouseResponseState = `DELETED`

const GetWarehouseResponseStateDeleting GetWarehouseResponseState = `DELETING`

const GetWarehouseResponseStateRunning GetWarehouseResponseState = `RUNNING`

const GetWarehouseResponseStateStarting GetWarehouseResponseState = `STARTING`

const GetWarehouseResponseStateStopped GetWarehouseResponseState = `STOPPED`

const GetWarehouseResponseStateStopping GetWarehouseResponseState = `STOPPING`

// String representation for [fmt.Print]
func (gwrs *GetWarehouseResponseState) String() string {
	return string(*gwrs)
}

// Set raw string value and validate it against allowed values
func (gwrs *GetWarehouseResponseState) Set(v string) error {
	switch v {
	case `DELETED`, `DELETING`, `RUNNING`, `STARTING`, `STOPPED`, `STOPPING`:
		*gwrs = GetWarehouseResponseState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETED", "DELETING", "RUNNING", "STARTING", "STOPPED", "STOPPING"`, v)
	}
}

// Type always returns GetWarehouseResponseState to satisfy [pflag.Value] interface
func (gwrs *GetWarehouseResponseState) Type() string {
	return "GetWarehouseResponseState"
}

// Warehouse type (Classic/Pro)
type GetWarehouseResponseWarehouseType string

const GetWarehouseResponseWarehouseTypeClassic GetWarehouseResponseWarehouseType = `CLASSIC`

const GetWarehouseResponseWarehouseTypePro GetWarehouseResponseWarehouseType = `PRO`

const GetWarehouseResponseWarehouseTypeTypeUnspecified GetWarehouseResponseWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (gwrwt *GetWarehouseResponseWarehouseType) String() string {
	return string(*gwrwt)
}

// Set raw string value and validate it against allowed values
func (gwrwt *GetWarehouseResponseWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `TYPE_UNSPECIFIED`:
		*gwrwt = GetWarehouseResponseWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "TYPE_UNSPECIFIED"`, v)
	}
}

// Type always returns GetWarehouseResponseWarehouseType to satisfy [pflag.Value] interface
func (gwrwt *GetWarehouseResponseWarehouseType) Type() string {
	return "GetWarehouseResponseWarehouseType"
}

type GetWorkspaceWarehouseConfigResponse struct {
	// Optional: Channel selection details
	Channel *Channel `json:"channel,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam *RepeatedEndpointConfPairs `json:"config_param,omitempty"`
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	DataAccessConfig []EndpointConfPair `json:"data_access_config,omitempty"`
	// Enable Serverless compute for SQL Endpoints
	//
	// Deprecated: Use enable_serverless_compute TODO(SC-79930): Remove the
	// field once clients are updated
	EnableDatabricksCompute bool `json:"enable_databricks_compute,omitempty"`
	// Enable Serverless compute for SQL Endpoints
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	EnabledWarehouseTypes []WarehouseTypePair `json:"enabled_warehouse_types,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	GlobalParam *RepeatedEndpointConfPairs `json:"global_param,omitempty"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount string `json:"google_service_account,omitempty"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Security policy for endpoints
	SecurityPolicy GetWorkspaceWarehouseConfigResponseSecurityPolicy `json:"security_policy,omitempty"`
	// SQL configuration parameters
	SqlConfigurationParameters *RepeatedEndpointConfPairs `json:"sql_configuration_parameters,omitempty"`
}

// Security policy for endpoints
type GetWorkspaceWarehouseConfigResponseSecurityPolicy string

const GetWorkspaceWarehouseConfigResponseSecurityPolicyDataAccessControl GetWorkspaceWarehouseConfigResponseSecurityPolicy = `DATA_ACCESS_CONTROL`

const GetWorkspaceWarehouseConfigResponseSecurityPolicyNone GetWorkspaceWarehouseConfigResponseSecurityPolicy = `NONE`

const GetWorkspaceWarehouseConfigResponseSecurityPolicyPassthrough GetWorkspaceWarehouseConfigResponseSecurityPolicy = `PASSTHROUGH`

// String representation for [fmt.Print]
func (gwwcrsp *GetWorkspaceWarehouseConfigResponseSecurityPolicy) String() string {
	return string(*gwwcrsp)
}

// Set raw string value and validate it against allowed values
func (gwwcrsp *GetWorkspaceWarehouseConfigResponseSecurityPolicy) Set(v string) error {
	switch v {
	case `DATA_ACCESS_CONTROL`, `NONE`, `PASSTHROUGH`:
		*gwwcrsp = GetWorkspaceWarehouseConfigResponseSecurityPolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATA_ACCESS_CONTROL", "NONE", "PASSTHROUGH"`, v)
	}
}

// Type always returns GetWorkspaceWarehouseConfigResponseSecurityPolicy to satisfy [pflag.Value] interface
func (gwwcrsp *GetWorkspaceWarehouseConfigResponseSecurityPolicy) Type() string {
	return "GetWorkspaceWarehouseConfigResponseSecurityPolicy"
}

// List
type ListQueriesRequest struct {
	// A filter to limit query history results. This field is optional.
	FilterBy *QueryFilter `json:"-" url:"filter_by,omitempty"`
	// Whether to include metrics about query.
	IncludeMetrics bool `json:"-" url:"include_metrics,omitempty"`
	// Limit the number of results returned in one page. The default is 100.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// A token that can be used to get the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`
}

type ListQueriesResponse struct {
	// Whether there is another page of results.
	HasNextPage bool `json:"has_next_page,omitempty"`
	// A token that can be used to get the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	Res []QueryInfo `json:"res,omitempty"`
}

// List warehouses
type ListWarehouses struct {
	// Service Principal which will be used to fetch the list of endpoints. If
	// not specified, the user from the session header is used.
	RunAsUserId int `json:"-" url:"run_as_user_id,omitempty"`
}

type ListWarehousesResponse struct {
	// A list of warehouses and their configurations.
	Warehouses []EndpointInfo `json:"warehouses,omitempty"`
}

type OdbcParams struct {
	Hostname string `json:"hostname,omitempty"`

	Path string `json:"path,omitempty"`

	Port int `json:"port,omitempty"`

	Protocol string `json:"protocol,omitempty"`
}

// Whether plans exist for the execution, or the reason why they are missing
type PlansState string

const PlansStateEmpty PlansState = `EMPTY`

const PlansStateExists PlansState = `EXISTS`

const PlansStateIgnoredLargePlansSize PlansState = `IGNORED_LARGE_PLANS_SIZE`

const PlansStateIgnoredSmallDuration PlansState = `IGNORED_SMALL_DURATION`

const PlansStateIgnoredSparkPlanType PlansState = `IGNORED_SPARK_PLAN_TYPE`

const PlansStateUnknown PlansState = `UNKNOWN`

// String representation for [fmt.Print]
func (ps *PlansState) String() string {
	return string(*ps)
}

// Set raw string value and validate it against allowed values
func (ps *PlansState) Set(v string) error {
	switch v {
	case `EMPTY`, `EXISTS`, `IGNORED_LARGE_PLANS_SIZE`, `IGNORED_SMALL_DURATION`, `IGNORED_SPARK_PLAN_TYPE`, `UNKNOWN`:
		*ps = PlansState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EMPTY", "EXISTS", "IGNORED_LARGE_PLANS_SIZE", "IGNORED_SMALL_DURATION", "IGNORED_SPARK_PLAN_TYPE", "UNKNOWN"`, v)
	}
}

// Type always returns PlansState to satisfy [pflag.Value] interface
func (ps *PlansState) Type() string {
	return "PlansState"
}

// A filter to limit query history results. This field is optional.
type QueryFilter struct {
	QueryStartTimeRange *TimeRange `json:"query_start_time_range,omitempty"`

	Statuses []QueryStatus `json:"statuses,omitempty"`
	// A list of user IDs who ran the queries.
	UserIds []int `json:"user_ids,omitempty"`
	// A list of warehouse IDs.
	WarehouseIds []string `json:"warehouse_ids,omitempty"`
}

type QueryInfo struct {
	// Channel information for the SQL warehouse at the time of query execution
	ChannelUsed *ChannelInfo `json:"channel_used,omitempty"`
	// Total execution time of the query from the client’s point of view, in
	// milliseconds.
	Duration int `json:"duration,omitempty"`
	// Alias for `warehouse_id`.
	EndpointId string `json:"endpoint_id,omitempty"`
	// Message describing why the query could not complete.
	ErrorMessage string `json:"error_message,omitempty"`
	// The ID of the user whose credentials were used to run the query.
	ExecutedAsUserId int `json:"executed_as_user_id,omitempty"`
	// The email address or username of the user whose credentials were used to
	// run the query.
	ExecutedAsUserName string `json:"executed_as_user_name,omitempty"`
	// The time execution of the query ended.
	ExecutionEndTimeMs int `json:"execution_end_time_ms,omitempty"`
	// Whether more updates for the query are expected.
	IsFinal bool `json:"is_final,omitempty"`
	// A key that can be used to look up query details.
	LookupKey string `json:"lookup_key,omitempty"`
	// Metrics about query execution.
	Metrics *QueryMetrics `json:"metrics,omitempty"`
	// Whether plans exist for the execution, or the reason why they are missing
	PlansState PlansState `json:"plans_state,omitempty"`
	// The time the query ended.
	QueryEndTimeMs int `json:"query_end_time_ms,omitempty"`
	// The query ID.
	QueryId string `json:"query_id,omitempty"`
	// The time the query started.
	QueryStartTimeMs int `json:"query_start_time_ms,omitempty"`
	// The text of the query.
	QueryText string `json:"query_text,omitempty"`
	// The number of results returned by the query.
	RowsProduced int `json:"rows_produced,omitempty"`
	// URL to the query plan.
	SparkUiUrl string `json:"spark_ui_url,omitempty"`
	// Type of statement for this query
	StatementType QueryStatementType `json:"statement_type,omitempty"`
	// This describes an enum
	Status QueryStatus `json:"status,omitempty"`
	// The ID of the user who ran the query.
	UserId int `json:"user_id,omitempty"`
	// The email address or username of the user who ran the query.
	UserName string `json:"user_name,omitempty"`
	// Warehouse ID.
	WarehouseId string `json:"warehouse_id,omitempty"`
}

// Metrics about query execution.
type QueryMetrics struct {
	// Time spent loading metadata and optimizing the query, in milliseconds.
	CompilationTimeMs int `json:"compilation_time_ms,omitempty"`
	// Time spent executing the query, in milliseconds.
	ExecutionTimeMs int `json:"execution_time_ms,omitempty"`
	// Total amount of data sent over the network, in bytes.
	NetworkSentBytes int `json:"network_sent_bytes,omitempty"`
	// Total execution time for all individual Photon query engine tasks in the
	// query, in milliseconds.
	PhotonTotalTimeMs int `json:"photon_total_time_ms,omitempty"`
	// Time spent waiting to execute the query because the SQL warehouse is
	// already running the maximum number of concurrent queries, in
	// milliseconds.
	QueuedOverloadTimeMs int `json:"queued_overload_time_ms,omitempty"`
	// Time waiting for compute resources to be provisioned for the SQL
	// warehouse, in milliseconds.
	QueuedProvisioningTimeMs int `json:"queued_provisioning_time_ms,omitempty"`
	// Total size of data read by the query, in bytes.
	ReadBytes int `json:"read_bytes,omitempty"`
	// Size of persistent data read from the cache, in bytes.
	ReadCacheBytes int `json:"read_cache_bytes,omitempty"`
	// Number of files read after pruning.
	ReadFilesCount int `json:"read_files_count,omitempty"`
	// Number of partitions read after pruning.
	ReadPartitionsCount int `json:"read_partitions_count,omitempty"`
	// Size of persistent data read from cloud object storage on your cloud
	// tenant, in bytes.
	ReadRemoteBytes int `json:"read_remote_bytes,omitempty"`
	// Time spent fetching the query results after the execution finished, in
	// milliseconds.
	ResultFetchTimeMs int `json:"result_fetch_time_ms,omitempty"`
	// true if the query result was fetched from cache, false otherwise.
	ResultFromCache bool `json:"result_from_cache,omitempty"`
	// Total number of rows returned by the query.
	RowsProducedCount int `json:"rows_produced_count,omitempty"`
	// Total number of rows read by the query.
	RowsReadCount int `json:"rows_read_count,omitempty"`
	// Size of data temporarily written to disk while executing the query, in
	// bytes.
	SpillToDiskBytes int `json:"spill_to_disk_bytes,omitempty"`
	// Sum of execution time for all of the query’s tasks, in milliseconds.
	TaskTotalTimeMs int `json:"task_total_time_ms,omitempty"`
	// Number of files that would have been read without pruning.
	TotalFilesCount int `json:"total_files_count,omitempty"`
	// Number of partitions that would have been read without pruning.
	TotalPartitionsCount int `json:"total_partitions_count,omitempty"`
	// Total execution time of the query from the client’s point of view, in
	// milliseconds.
	TotalTimeMs int `json:"total_time_ms,omitempty"`
	// Size pf persistent data written to cloud object storage in your cloud
	// tenant, in bytes.
	WriteRemoteBytes int `json:"write_remote_bytes,omitempty"`
}

// Type of statement for this query
type QueryStatementType string

const QueryStatementTypeAlter QueryStatementType = `ALTER`

const QueryStatementTypeAnalyze QueryStatementType = `ANALYZE`

const QueryStatementTypeCopy QueryStatementType = `COPY`

const QueryStatementTypeCreate QueryStatementType = `CREATE`

const QueryStatementTypeDelete QueryStatementType = `DELETE`

const QueryStatementTypeDescribe QueryStatementType = `DESCRIBE`

const QueryStatementTypeDrop QueryStatementType = `DROP`

const QueryStatementTypeExplain QueryStatementType = `EXPLAIN`

const QueryStatementTypeGrant QueryStatementType = `GRANT`

const QueryStatementTypeInsert QueryStatementType = `INSERT`

const QueryStatementTypeMerge QueryStatementType = `MERGE`

const QueryStatementTypeOptimize QueryStatementType = `OPTIMIZE`

const QueryStatementTypeOther QueryStatementType = `OTHER`

const QueryStatementTypeRefresh QueryStatementType = `REFRESH`

const QueryStatementTypeReplace QueryStatementType = `REPLACE`

const QueryStatementTypeRevoke QueryStatementType = `REVOKE`

const QueryStatementTypeSelect QueryStatementType = `SELECT`

const QueryStatementTypeSet QueryStatementType = `SET`

const QueryStatementTypeShow QueryStatementType = `SHOW`

const QueryStatementTypeTruncate QueryStatementType = `TRUNCATE`

const QueryStatementTypeUpdate QueryStatementType = `UPDATE`

const QueryStatementTypeUse QueryStatementType = `USE`

// String representation for [fmt.Print]
func (qst *QueryStatementType) String() string {
	return string(*qst)
}

// Set raw string value and validate it against allowed values
func (qst *QueryStatementType) Set(v string) error {
	switch v {
	case `ALTER`, `ANALYZE`, `COPY`, `CREATE`, `DELETE`, `DESCRIBE`, `DROP`, `EXPLAIN`, `GRANT`, `INSERT`, `MERGE`, `OPTIMIZE`, `OTHER`, `REFRESH`, `REPLACE`, `REVOKE`, `SELECT`, `SET`, `SHOW`, `TRUNCATE`, `UPDATE`, `USE`:
		*qst = QueryStatementType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALTER", "ANALYZE", "COPY", "CREATE", "DELETE", "DESCRIBE", "DROP", "EXPLAIN", "GRANT", "INSERT", "MERGE", "OPTIMIZE", "OTHER", "REFRESH", "REPLACE", "REVOKE", "SELECT", "SET", "SHOW", "TRUNCATE", "UPDATE", "USE"`, v)
	}
}

// Type always returns QueryStatementType to satisfy [pflag.Value] interface
func (qst *QueryStatementType) Type() string {
	return "QueryStatementType"
}

// This describes an enum
type QueryStatus string

// Query has been cancelled by the user.
const QueryStatusCanceled QueryStatus = `CANCELED`

// Query has failed.
const QueryStatusFailed QueryStatus = `FAILED`

// Query has completed.
const QueryStatusFinished QueryStatus = `FINISHED`

// Query has been received and queued.
const QueryStatusQueued QueryStatus = `QUEUED`

// Query has started.
const QueryStatusRunning QueryStatus = `RUNNING`

// String representation for [fmt.Print]
func (qs *QueryStatus) String() string {
	return string(*qs)
}

// Set raw string value and validate it against allowed values
func (qs *QueryStatus) Set(v string) error {
	switch v {
	case `CANCELED`, `FAILED`, `FINISHED`, `QUEUED`, `RUNNING`:
		*qs = QueryStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "FAILED", "FINISHED", "QUEUED", "RUNNING"`, v)
	}
}

// Type always returns QueryStatus to satisfy [pflag.Value] interface
func (qs *QueryStatus) Type() string {
	return "QueryStatus"
}

type RepeatedEndpointConfPairs struct {
	// Deprecated: Use configuration_pairs
	ConfigPair []EndpointConfPair `json:"config_pair,omitempty"`

	ConfigurationPairs []EndpointConfPair `json:"configuration_pairs,omitempty"`
}

type SetWorkspaceWarehouseConfigRequest struct {
	// Optional: Channel selection details
	Channel *Channel `json:"channel,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam *RepeatedEndpointConfPairs `json:"config_param,omitempty"`
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	DataAccessConfig []EndpointConfPair `json:"data_access_config,omitempty"`
	// Enable Serverless compute for SQL Endpoints
	//
	// Deprecated: Use enable_serverless_compute TODO(SC-79930): Remove the
	// field once clients are updated
	EnableDatabricksCompute bool `json:"enable_databricks_compute,omitempty"`
	// Enable Serverless compute for SQL Endpoints
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	EnabledWarehouseTypes []WarehouseTypePair `json:"enabled_warehouse_types,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	GlobalParam *RepeatedEndpointConfPairs `json:"global_param,omitempty"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount string `json:"google_service_account,omitempty"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Security policy for endpoints
	SecurityPolicy SetWorkspaceWarehouseConfigRequestSecurityPolicy `json:"security_policy,omitempty"`
	// Internal. Used by frontend to save Serverless Compute agreement value.
	ServerlessAgreement bool `json:"serverless_agreement,omitempty"`
	// SQL configuration parameters
	SqlConfigurationParameters *RepeatedEndpointConfPairs `json:"sql_configuration_parameters,omitempty"`
}

// Security policy for endpoints
type SetWorkspaceWarehouseConfigRequestSecurityPolicy string

const SetWorkspaceWarehouseConfigRequestSecurityPolicyDataAccessControl SetWorkspaceWarehouseConfigRequestSecurityPolicy = `DATA_ACCESS_CONTROL`

const SetWorkspaceWarehouseConfigRequestSecurityPolicyNone SetWorkspaceWarehouseConfigRequestSecurityPolicy = `NONE`

const SetWorkspaceWarehouseConfigRequestSecurityPolicyPassthrough SetWorkspaceWarehouseConfigRequestSecurityPolicy = `PASSTHROUGH`

// String representation for [fmt.Print]
func (swwcrsp *SetWorkspaceWarehouseConfigRequestSecurityPolicy) String() string {
	return string(*swwcrsp)
}

// Set raw string value and validate it against allowed values
func (swwcrsp *SetWorkspaceWarehouseConfigRequestSecurityPolicy) Set(v string) error {
	switch v {
	case `DATA_ACCESS_CONTROL`, `NONE`, `PASSTHROUGH`:
		*swwcrsp = SetWorkspaceWarehouseConfigRequestSecurityPolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATA_ACCESS_CONTROL", "NONE", "PASSTHROUGH"`, v)
	}
}

// Type always returns SetWorkspaceWarehouseConfigRequestSecurityPolicy to satisfy [pflag.Value] interface
func (swwcrsp *SetWorkspaceWarehouseConfigRequestSecurityPolicy) Type() string {
	return "SetWorkspaceWarehouseConfigRequestSecurityPolicy"
}

// Start a warehouse
type StartWarehouse struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" url:"-"`
}

// Stop a warehouse
type StopWarehouse struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" url:"-"`
}

type TerminationReason struct {
	// status code indicating why the cluster was terminated
	Code TerminationReasonCode `json:"code,omitempty"`
	// list of parameters that provide additional information about why the
	// cluster was terminated
	Parameters map[string]string `json:"parameters,omitempty"`
	// type of the termination
	Type TerminationReasonType `json:"type,omitempty"`
}

// status code indicating why the cluster was terminated
type TerminationReasonCode string

const TerminationReasonCodeAbuseDetected TerminationReasonCode = `ABUSE_DETECTED`

const TerminationReasonCodeAttachProjectFailure TerminationReasonCode = `ATTACH_PROJECT_FAILURE`

const TerminationReasonCodeAwsAuthorizationFailure TerminationReasonCode = `AWS_AUTHORIZATION_FAILURE`

const TerminationReasonCodeAwsInsufficientFreeAddressesInSubnetFailure TerminationReasonCode = `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`

const TerminationReasonCodeAwsInsufficientInstanceCapacityFailure TerminationReasonCode = `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`

const TerminationReasonCodeAwsMaxSpotInstanceCountExceededFailure TerminationReasonCode = `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`

const TerminationReasonCodeAwsRequestLimitExceeded TerminationReasonCode = `AWS_REQUEST_LIMIT_EXCEEDED`

const TerminationReasonCodeAwsUnsupportedFailure TerminationReasonCode = `AWS_UNSUPPORTED_FAILURE`

const TerminationReasonCodeAzureByokKeyPermissionFailure TerminationReasonCode = `AZURE_BYOK_KEY_PERMISSION_FAILURE`

const TerminationReasonCodeAzureEphemeralDiskFailure TerminationReasonCode = `AZURE_EPHEMERAL_DISK_FAILURE`

const TerminationReasonCodeAzureInvalidDeploymentTemplate TerminationReasonCode = `AZURE_INVALID_DEPLOYMENT_TEMPLATE`

const TerminationReasonCodeAzureOperationNotAllowedException TerminationReasonCode = `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`

const TerminationReasonCodeAzureQuotaExceededException TerminationReasonCode = `AZURE_QUOTA_EXCEEDED_EXCEPTION`

const TerminationReasonCodeAzureResourceManagerThrottling TerminationReasonCode = `AZURE_RESOURCE_MANAGER_THROTTLING`

const TerminationReasonCodeAzureResourceProviderThrottling TerminationReasonCode = `AZURE_RESOURCE_PROVIDER_THROTTLING`

const TerminationReasonCodeAzureUnexpectedDeploymentTemplateFailure TerminationReasonCode = `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`

const TerminationReasonCodeAzureVmExtensionFailure TerminationReasonCode = `AZURE_VM_EXTENSION_FAILURE`

const TerminationReasonCodeAzureVnetConfigurationFailure TerminationReasonCode = `AZURE_VNET_CONFIGURATION_FAILURE`

const TerminationReasonCodeBootstrapTimeout TerminationReasonCode = `BOOTSTRAP_TIMEOUT`

const TerminationReasonCodeBootstrapTimeoutCloudProviderException TerminationReasonCode = `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`

const TerminationReasonCodeCloudProviderDiskSetupFailure TerminationReasonCode = `CLOUD_PROVIDER_DISK_SETUP_FAILURE`

const TerminationReasonCodeCloudProviderLaunchFailure TerminationReasonCode = `CLOUD_PROVIDER_LAUNCH_FAILURE`

const TerminationReasonCodeCloudProviderResourceStockout TerminationReasonCode = `CLOUD_PROVIDER_RESOURCE_STOCKOUT`

const TerminationReasonCodeCloudProviderShutdown TerminationReasonCode = `CLOUD_PROVIDER_SHUTDOWN`

const TerminationReasonCodeCommunicationLost TerminationReasonCode = `COMMUNICATION_LOST`

const TerminationReasonCodeContainerLaunchFailure TerminationReasonCode = `CONTAINER_LAUNCH_FAILURE`

const TerminationReasonCodeControlPlaneRequestFailure TerminationReasonCode = `CONTROL_PLANE_REQUEST_FAILURE`

const TerminationReasonCodeDatabaseConnectionFailure TerminationReasonCode = `DATABASE_CONNECTION_FAILURE`

const TerminationReasonCodeDbfsComponentUnhealthy TerminationReasonCode = `DBFS_COMPONENT_UNHEALTHY`

const TerminationReasonCodeDockerImagePullFailure TerminationReasonCode = `DOCKER_IMAGE_PULL_FAILURE`

const TerminationReasonCodeDriverUnreachable TerminationReasonCode = `DRIVER_UNREACHABLE`

const TerminationReasonCodeDriverUnresponsive TerminationReasonCode = `DRIVER_UNRESPONSIVE`

const TerminationReasonCodeExecutionComponentUnhealthy TerminationReasonCode = `EXECUTION_COMPONENT_UNHEALTHY`

const TerminationReasonCodeGcpQuotaExceeded TerminationReasonCode = `GCP_QUOTA_EXCEEDED`

const TerminationReasonCodeGcpServiceAccountDeleted TerminationReasonCode = `GCP_SERVICE_ACCOUNT_DELETED`

const TerminationReasonCodeGlobalInitScriptFailure TerminationReasonCode = `GLOBAL_INIT_SCRIPT_FAILURE`

const TerminationReasonCodeHiveMetastoreProvisioningFailure TerminationReasonCode = `HIVE_METASTORE_PROVISIONING_FAILURE`

const TerminationReasonCodeImagePullPermissionDenied TerminationReasonCode = `IMAGE_PULL_PERMISSION_DENIED`

const TerminationReasonCodeInactivity TerminationReasonCode = `INACTIVITY`

const TerminationReasonCodeInitScriptFailure TerminationReasonCode = `INIT_SCRIPT_FAILURE`

const TerminationReasonCodeInstancePoolClusterFailure TerminationReasonCode = `INSTANCE_POOL_CLUSTER_FAILURE`

const TerminationReasonCodeInstanceUnreachable TerminationReasonCode = `INSTANCE_UNREACHABLE`

const TerminationReasonCodeInternalError TerminationReasonCode = `INTERNAL_ERROR`

const TerminationReasonCodeInvalidArgument TerminationReasonCode = `INVALID_ARGUMENT`

const TerminationReasonCodeInvalidSparkImage TerminationReasonCode = `INVALID_SPARK_IMAGE`

const TerminationReasonCodeIpExhaustionFailure TerminationReasonCode = `IP_EXHAUSTION_FAILURE`

const TerminationReasonCodeJobFinished TerminationReasonCode = `JOB_FINISHED`

const TerminationReasonCodeKsAutoscalingFailure TerminationReasonCode = `K8S_AUTOSCALING_FAILURE`

const TerminationReasonCodeKsDbrClusterLaunchTimeout TerminationReasonCode = `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`

const TerminationReasonCodeMetastoreComponentUnhealthy TerminationReasonCode = `METASTORE_COMPONENT_UNHEALTHY`

const TerminationReasonCodeNephosResourceManagement TerminationReasonCode = `NEPHOS_RESOURCE_MANAGEMENT`

const TerminationReasonCodeNetworkConfigurationFailure TerminationReasonCode = `NETWORK_CONFIGURATION_FAILURE`

const TerminationReasonCodeNfsMountFailure TerminationReasonCode = `NFS_MOUNT_FAILURE`

const TerminationReasonCodeNpipTunnelSetupFailure TerminationReasonCode = `NPIP_TUNNEL_SETUP_FAILURE`

const TerminationReasonCodeNpipTunnelTokenFailure TerminationReasonCode = `NPIP_TUNNEL_TOKEN_FAILURE`

const TerminationReasonCodeRequestRejected TerminationReasonCode = `REQUEST_REJECTED`

const TerminationReasonCodeRequestThrottled TerminationReasonCode = `REQUEST_THROTTLED`

const TerminationReasonCodeSecretResolutionError TerminationReasonCode = `SECRET_RESOLUTION_ERROR`

const TerminationReasonCodeSecurityDaemonRegistrationException TerminationReasonCode = `SECURITY_DAEMON_REGISTRATION_EXCEPTION`

const TerminationReasonCodeSelfBootstrapFailure TerminationReasonCode = `SELF_BOOTSTRAP_FAILURE`

const TerminationReasonCodeSkippedSlowNodes TerminationReasonCode = `SKIPPED_SLOW_NODES`

const TerminationReasonCodeSlowImageDownload TerminationReasonCode = `SLOW_IMAGE_DOWNLOAD`

const TerminationReasonCodeSparkError TerminationReasonCode = `SPARK_ERROR`

const TerminationReasonCodeSparkImageDownloadFailure TerminationReasonCode = `SPARK_IMAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeSparkStartupFailure TerminationReasonCode = `SPARK_STARTUP_FAILURE`

const TerminationReasonCodeSpotInstanceTermination TerminationReasonCode = `SPOT_INSTANCE_TERMINATION`

const TerminationReasonCodeStorageDownloadFailure TerminationReasonCode = `STORAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeStsClientSetupFailure TerminationReasonCode = `STS_CLIENT_SETUP_FAILURE`

const TerminationReasonCodeSubnetExhaustedFailure TerminationReasonCode = `SUBNET_EXHAUSTED_FAILURE`

const TerminationReasonCodeTemporarilyUnavailable TerminationReasonCode = `TEMPORARILY_UNAVAILABLE`

const TerminationReasonCodeTrialExpired TerminationReasonCode = `TRIAL_EXPIRED`

const TerminationReasonCodeUnexpectedLaunchFailure TerminationReasonCode = `UNEXPECTED_LAUNCH_FAILURE`

const TerminationReasonCodeUnknown TerminationReasonCode = `UNKNOWN`

const TerminationReasonCodeUnsupportedInstanceType TerminationReasonCode = `UNSUPPORTED_INSTANCE_TYPE`

const TerminationReasonCodeUpdateInstanceProfileFailure TerminationReasonCode = `UPDATE_INSTANCE_PROFILE_FAILURE`

const TerminationReasonCodeUserRequest TerminationReasonCode = `USER_REQUEST`

const TerminationReasonCodeWorkerSetupFailure TerminationReasonCode = `WORKER_SETUP_FAILURE`

const TerminationReasonCodeWorkspaceCancelledError TerminationReasonCode = `WORKSPACE_CANCELLED_ERROR`

const TerminationReasonCodeWorkspaceConfigurationError TerminationReasonCode = `WORKSPACE_CONFIGURATION_ERROR`

// String representation for [fmt.Print]
func (trc *TerminationReasonCode) String() string {
	return string(*trc)
}

// Set raw string value and validate it against allowed values
func (trc *TerminationReasonCode) Set(v string) error {
	switch v {
	case `ABUSE_DETECTED`, `ATTACH_PROJECT_FAILURE`, `AWS_AUTHORIZATION_FAILURE`, `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`, `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`, `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`, `AWS_REQUEST_LIMIT_EXCEEDED`, `AWS_UNSUPPORTED_FAILURE`, `AZURE_BYOK_KEY_PERMISSION_FAILURE`, `AZURE_EPHEMERAL_DISK_FAILURE`, `AZURE_INVALID_DEPLOYMENT_TEMPLATE`, `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`, `AZURE_QUOTA_EXCEEDED_EXCEPTION`, `AZURE_RESOURCE_MANAGER_THROTTLING`, `AZURE_RESOURCE_PROVIDER_THROTTLING`, `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`, `AZURE_VM_EXTENSION_FAILURE`, `AZURE_VNET_CONFIGURATION_FAILURE`, `BOOTSTRAP_TIMEOUT`, `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`, `CLOUD_PROVIDER_DISK_SETUP_FAILURE`, `CLOUD_PROVIDER_LAUNCH_FAILURE`, `CLOUD_PROVIDER_RESOURCE_STOCKOUT`, `CLOUD_PROVIDER_SHUTDOWN`, `COMMUNICATION_LOST`, `CONTAINER_LAUNCH_FAILURE`, `CONTROL_PLANE_REQUEST_FAILURE`, `DATABASE_CONNECTION_FAILURE`, `DBFS_COMPONENT_UNHEALTHY`, `DOCKER_IMAGE_PULL_FAILURE`, `DRIVER_UNREACHABLE`, `DRIVER_UNRESPONSIVE`, `EXECUTION_COMPONENT_UNHEALTHY`, `GCP_QUOTA_EXCEEDED`, `GCP_SERVICE_ACCOUNT_DELETED`, `GLOBAL_INIT_SCRIPT_FAILURE`, `HIVE_METASTORE_PROVISIONING_FAILURE`, `IMAGE_PULL_PERMISSION_DENIED`, `INACTIVITY`, `INIT_SCRIPT_FAILURE`, `INSTANCE_POOL_CLUSTER_FAILURE`, `INSTANCE_UNREACHABLE`, `INTERNAL_ERROR`, `INVALID_ARGUMENT`, `INVALID_SPARK_IMAGE`, `IP_EXHAUSTION_FAILURE`, `JOB_FINISHED`, `K8S_AUTOSCALING_FAILURE`, `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`, `METASTORE_COMPONENT_UNHEALTHY`, `NEPHOS_RESOURCE_MANAGEMENT`, `NETWORK_CONFIGURATION_FAILURE`, `NFS_MOUNT_FAILURE`, `NPIP_TUNNEL_SETUP_FAILURE`, `NPIP_TUNNEL_TOKEN_FAILURE`, `REQUEST_REJECTED`, `REQUEST_THROTTLED`, `SECRET_RESOLUTION_ERROR`, `SECURITY_DAEMON_REGISTRATION_EXCEPTION`, `SELF_BOOTSTRAP_FAILURE`, `SKIPPED_SLOW_NODES`, `SLOW_IMAGE_DOWNLOAD`, `SPARK_ERROR`, `SPARK_IMAGE_DOWNLOAD_FAILURE`, `SPARK_STARTUP_FAILURE`, `SPOT_INSTANCE_TERMINATION`, `STORAGE_DOWNLOAD_FAILURE`, `STS_CLIENT_SETUP_FAILURE`, `SUBNET_EXHAUSTED_FAILURE`, `TEMPORARILY_UNAVAILABLE`, `TRIAL_EXPIRED`, `UNEXPECTED_LAUNCH_FAILURE`, `UNKNOWN`, `UNSUPPORTED_INSTANCE_TYPE`, `UPDATE_INSTANCE_PROFILE_FAILURE`, `USER_REQUEST`, `WORKER_SETUP_FAILURE`, `WORKSPACE_CANCELLED_ERROR`, `WORKSPACE_CONFIGURATION_ERROR`:
		*trc = TerminationReasonCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABUSE_DETECTED", "ATTACH_PROJECT_FAILURE", "AWS_AUTHORIZATION_FAILURE", "AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE", "AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE", "AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE", "AWS_REQUEST_LIMIT_EXCEEDED", "AWS_UNSUPPORTED_FAILURE", "AZURE_BYOK_KEY_PERMISSION_FAILURE", "AZURE_EPHEMERAL_DISK_FAILURE", "AZURE_INVALID_DEPLOYMENT_TEMPLATE", "AZURE_OPERATION_NOT_ALLOWED_EXCEPTION", "AZURE_QUOTA_EXCEEDED_EXCEPTION", "AZURE_RESOURCE_MANAGER_THROTTLING", "AZURE_RESOURCE_PROVIDER_THROTTLING", "AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE", "AZURE_VM_EXTENSION_FAILURE", "AZURE_VNET_CONFIGURATION_FAILURE", "BOOTSTRAP_TIMEOUT", "BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION", "CLOUD_PROVIDER_DISK_SETUP_FAILURE", "CLOUD_PROVIDER_LAUNCH_FAILURE", "CLOUD_PROVIDER_RESOURCE_STOCKOUT", "CLOUD_PROVIDER_SHUTDOWN", "COMMUNICATION_LOST", "CONTAINER_LAUNCH_FAILURE", "CONTROL_PLANE_REQUEST_FAILURE", "DATABASE_CONNECTION_FAILURE", "DBFS_COMPONENT_UNHEALTHY", "DOCKER_IMAGE_PULL_FAILURE", "DRIVER_UNREACHABLE", "DRIVER_UNRESPONSIVE", "EXECUTION_COMPONENT_UNHEALTHY", "GCP_QUOTA_EXCEEDED", "GCP_SERVICE_ACCOUNT_DELETED", "GLOBAL_INIT_SCRIPT_FAILURE", "HIVE_METASTORE_PROVISIONING_FAILURE", "IMAGE_PULL_PERMISSION_DENIED", "INACTIVITY", "INIT_SCRIPT_FAILURE", "INSTANCE_POOL_CLUSTER_FAILURE", "INSTANCE_UNREACHABLE", "INTERNAL_ERROR", "INVALID_ARGUMENT", "INVALID_SPARK_IMAGE", "IP_EXHAUSTION_FAILURE", "JOB_FINISHED", "K8S_AUTOSCALING_FAILURE", "K8S_DBR_CLUSTER_LAUNCH_TIMEOUT", "METASTORE_COMPONENT_UNHEALTHY", "NEPHOS_RESOURCE_MANAGEMENT", "NETWORK_CONFIGURATION_FAILURE", "NFS_MOUNT_FAILURE", "NPIP_TUNNEL_SETUP_FAILURE", "NPIP_TUNNEL_TOKEN_FAILURE", "REQUEST_REJECTED", "REQUEST_THROTTLED", "SECRET_RESOLUTION_ERROR", "SECURITY_DAEMON_REGISTRATION_EXCEPTION", "SELF_BOOTSTRAP_FAILURE", "SKIPPED_SLOW_NODES", "SLOW_IMAGE_DOWNLOAD", "SPARK_ERROR", "SPARK_IMAGE_DOWNLOAD_FAILURE", "SPARK_STARTUP_FAILURE", "SPOT_INSTANCE_TERMINATION", "STORAGE_DOWNLOAD_FAILURE", "STS_CLIENT_SETUP_FAILURE", "SUBNET_EXHAUSTED_FAILURE", "TEMPORARILY_UNAVAILABLE", "TRIAL_EXPIRED", "UNEXPECTED_LAUNCH_FAILURE", "UNKNOWN", "UNSUPPORTED_INSTANCE_TYPE", "UPDATE_INSTANCE_PROFILE_FAILURE", "USER_REQUEST", "WORKER_SETUP_FAILURE", "WORKSPACE_CANCELLED_ERROR", "WORKSPACE_CONFIGURATION_ERROR"`, v)
	}
}

// Type always returns TerminationReasonCode to satisfy [pflag.Value] interface
func (trc *TerminationReasonCode) Type() string {
	return "TerminationReasonCode"
}

// type of the termination
type TerminationReasonType string

const TerminationReasonTypeClientError TerminationReasonType = `CLIENT_ERROR`

const TerminationReasonTypeCloudFailure TerminationReasonType = `CLOUD_FAILURE`

const TerminationReasonTypeServiceFault TerminationReasonType = `SERVICE_FAULT`

const TerminationReasonTypeSuccess TerminationReasonType = `SUCCESS`

// String representation for [fmt.Print]
func (trt *TerminationReasonType) String() string {
	return string(*trt)
}

// Set raw string value and validate it against allowed values
func (trt *TerminationReasonType) Set(v string) error {
	switch v {
	case `CLIENT_ERROR`, `CLOUD_FAILURE`, `SERVICE_FAULT`, `SUCCESS`:
		*trt = TerminationReasonType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLIENT_ERROR", "CLOUD_FAILURE", "SERVICE_FAULT", "SUCCESS"`, v)
	}
}

// Type always returns TerminationReasonType to satisfy [pflag.Value] interface
func (trt *TerminationReasonType) Type() string {
	return "TerminationReasonType"
}

type TimeRange struct {
	// Limit results to queries that started before this time.
	EndTimeMs int `json:"end_time_ms,omitempty"`
	// Limit results to queries that started after this time.
	StartTimeMs int `json:"start_time_ms,omitempty"`
}

type WarehouseTypePair struct {
	// If set to false the specific warehouse type will not be be allowed as a
	// value for warehouse_type in CreateWarehouse and EditWarehouse
	Enabled bool `json:"enabled,omitempty"`

	WarehouseType WarehouseTypePairWarehouseType `json:"warehouse_type,omitempty"`
}

type WarehouseTypePairWarehouseType string

const WarehouseTypePairWarehouseTypeClassic WarehouseTypePairWarehouseType = `CLASSIC`

const WarehouseTypePairWarehouseTypePro WarehouseTypePairWarehouseType = `PRO`

const WarehouseTypePairWarehouseTypeTypeUnspecified WarehouseTypePairWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (wtpwt *WarehouseTypePairWarehouseType) String() string {
	return string(*wtpwt)
}

// Set raw string value and validate it against allowed values
func (wtpwt *WarehouseTypePairWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `TYPE_UNSPECIFIED`:
		*wtpwt = WarehouseTypePairWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "TYPE_UNSPECIFIED"`, v)
	}
}

// Type always returns WarehouseTypePairWarehouseType to satisfy [pflag.Value] interface
func (wtpwt *WarehouseTypePairWarehouseType) Type() string {
	return "WarehouseTypePairWarehouseType"
}
