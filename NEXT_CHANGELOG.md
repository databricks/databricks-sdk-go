# NEXT CHANGELOG

## Release v0.98.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `CreateDefaultWarehouseOverride`, `DeleteDefaultWarehouseOverride`, `GetDefaultWarehouseOverride`, `ListDefaultWarehouseOverrides` and `UpdateDefaultWarehouseOverride` methods for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service.
* Add `DriverNodeTypeFlexibility` and `WorkerNodeTypeFlexibility` fields for [compute.ClusterAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterAttributes).
* Add `DriverNodeTypeFlexibility` and `WorkerNodeTypeFlexibility` fields for [compute.ClusterDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterDetails).
* Add `DriverNodeTypeFlexibility` and `WorkerNodeTypeFlexibility` fields for [compute.ClusterSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterSpec).
* Add `DriverNodeTypeFlexibility` and `WorkerNodeTypeFlexibility` fields for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
* Add `NodeTypeFlexibility` field for [compute.CreateInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateInstancePool).
* Add `DriverNodeTypeFlexibility` and `WorkerNodeTypeFlexibility` fields for [compute.EditCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditCluster).
* Add `NodeTypeFlexibility` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
* Add `NodeTypeFlexibility` field for [compute.GetInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePool).
* Add `NodeTypeFlexibility` field for [compute.InstancePoolAndStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAndStats).
* Add `DriverNodeTypeFlexibility` and `WorkerNodeTypeFlexibility` fields for [compute.UpdateClusterResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#UpdateClusterResource).
* Add `FullRefreshWindow` field for [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
* Add `AutoFullRefreshPolicy` field for [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
* Add `ExpireTime` and `Ttl` fields for [postgres.BranchSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchSpec).
* Add `ExpireTime` field for [postgres.BranchStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchStatus).
* Add `Hosts` field for [postgres.EndpointStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatus).
* Add `EndpointTypeReadWrite` and `EndpointTypeReadOnly` enum values for [postgres.EndpointType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointType).
* Add `Deleted` enum value for [vectorsearch.EndpointStatusState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointStatusState).
* [Breaking] Remove `Default` field for [postgres.BranchSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchSpec).
* [Breaking] Remove `Host`, `LastActiveTime`, `StartTime` and `SuspendTime` fields for [postgres.EndpointStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatus).
* [Breaking] Remove `Settings` field for [postgres.ProjectSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectSpec).
* [Breaking] Remove `ComputeLastActiveTime` and `Settings` fields for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* [Breaking] Remove `ReadWrite` and `ReadOnly` enum values for [postgres.EndpointType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointType).
* Add `NoExpiry` field for [postgres.BranchSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchSpec).
* Add `Stderr` field for [compute.InitScriptInfoAndExecutionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptInfoAndExecutionDetails).
* Add `GenerateDatabaseCredential` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `GetPublicAccountUserPreference`, `ListAccountUserPreferencesMetadata` and `PatchPublicAccountUserPreference` methods for [a.AccountSettingsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#AccountSettingsV2API) account-level service.
