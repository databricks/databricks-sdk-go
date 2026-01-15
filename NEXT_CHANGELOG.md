# NEXT CHANGELOG

## Release v0.98.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `FullRefreshWindow` field for [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
* Add `AutoFullRefreshPolicy` field for [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
* Add `Hosts` field for [postgres.EndpointStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatus).
* Add `EndpointTypeReadWrite` and `EndpointTypeReadOnly` enum values for [postgres.EndpointType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointType).
* Add `Deleted` enum value for [vectorsearch.EndpointStatusState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointStatusState).
* [Breaking] Change `CreateBranch`, `CreateEndpoint` and `CreateProject` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service with new required argument order.
* Change `BranchId` field for [postgres.CreateBranchRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateBranchRequest) to no longer be required.
* Change `EndpointId` field for [postgres.CreateEndpointRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateEndpointRequest) to no longer be required.
* Change `ProjectId` field for [postgres.CreateProjectRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateProjectRequest) to no longer be required.
* [Breaking] Remove `Host`, `LastActiveTime`, `StartTime` and `SuspendTime` fields for [postgres.EndpointStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatus).
* [Breaking] Remove `ComputeLastActiveTime` field for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* [Breaking] Remove `ReadWrite` and `ReadOnly` enum values for [postgres.EndpointType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointType).