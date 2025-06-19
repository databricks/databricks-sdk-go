# NEXT CHANGELOG

## Release v0.73.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added [w.AiBuilder](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aibuilder#AiBuilderAPI) workspace-level service.
* Added [w.FeatureStore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureStoreAPI) workspace-level service.
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.ClusterAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterAttributes).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.ClusterDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterDetails).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.ClusterSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterSpec).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.CreateInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateInstancePool).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.EditCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditCluster).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.GetInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePool).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.InstancePoolAndStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAndStats).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.UpdateClusterResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#UpdateClusterResource).
* Added `ExpirationTime` field for [database.DatabaseCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseCredential).
* Added `EffectiveStopped` field for [database.DatabaseInstance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstance).
* Added `ExistingPipelineId` field for [database.SyncedTableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableSpec).
* Added `PipelineId` field for [database.SyncedTableStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableStatus).
* Added `DbtPlatformOutput` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
* Added `DbtPlatformTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Added `DbtPlatformTask` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Added `DbtPlatformTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Added `Environment` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Added `Environment` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Added `Environment` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
* Added `Description` field for [serving.ServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpoint).
* Added `Description` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).
* Added `R` enum value for [compute.Language](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Language).
* Added `Cancelled`, `Error`, `Queued`, `Running`, `Starting` and `Success` enum values for [jobs.DbtPlatformRunStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DbtPlatformRunStatus).
* Added `Continuous` and `ContinuousRestart` enum values for [jobs.TriggerType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerType).
* [Breaking] Changed `Status` field for [jobs.DbtCloudJobRunStep](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DbtCloudJobRunStep) to type [jobs.DbtPlatformRunStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DbtPlatformRunStatus).
* [Breaking] Removed [w.CustomLlms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aibuilder#CustomLlmsAPI) workspace-level service.
* [Breaking] Removed `TableServingUrl` field for [database.DatabaseTable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseTable).
* [Breaking] Removed `TableServingUrl` field for [database.SyncedDatabaseTable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedDatabaseTable).
* [Breaking] Removed `PipelineId` field for [database.SyncedTableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableSpec).
* [Breaking] Removed `Cancelled`, `Error`, `Queued`, `Running`, `Starting` and `Success` enum values for [jobs.DbtCloudRunStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DbtCloudRunStatus).
