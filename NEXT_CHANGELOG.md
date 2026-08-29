# NEXT CHANGELOG

## Release v0.177.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `UpdateDeployment` method for [w.BundleDeployments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#BundleDeploymentsAPI) workspace-level service.
* Add `GenieCancelResponse` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Add `GetExternalGroup`, `GetExternalServicePrincipal` and `GetExternalUser` methods for [a.AccountIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API) account-level service.
* Add `GetExternalGroupProxy`, `GetExternalServicePrincipalProxy` and `GetExternalUserProxy` methods for [w.WorkspaceIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API) workspace-level service.
* Add `CreateSnapshot`, `DeleteSnapshot`, `GetSnapshot`, `GetSnapshotSchedule`, `ListSnapshots` and `UpdateSnapshotSchedule` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `AssumeGroupId` field for [apps.Space](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#Space).
* Add `Operations` field for [bundledeployments.Version](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Version).
* Add `FunctionArgExpression` field for [catalog.FunctionArgument](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionArgument).
* Add `AwsContextId` field for [compute.NodeTypeFlexibility](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#NodeTypeFlexibility).
* Add `Parameters` field for [jobs.AlertTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#AlertTask).
* Add `MaintenanceWindow` field for [jobs.Continuous](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Continuous).
* Add `MaintenanceWindow` field for [jobs.ContinuousTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ContinuousTriggerConfiguration).
* Add `Lateness` field for [ml.DataSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DataSource).
* Add `Delay` and `Offset` fields for [ml.SlidingWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SlidingWindow).
* Add `StartTime` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
* Add `Delay` and `Offset` fields for [ml.TumblingWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TumblingWindow).
* Add `RabbitmqOptions` field for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `SourceSnapshot` field for [postgres.BranchSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchSpec).
* Add `SourceSnapshot` field for [postgres.BranchStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchStatus).
* Add `DeploymentResourceTypeVectorSearchEndpoint`, `DeploymentResourceTypeVectorSearchIndex`, `DeploymentResourceTypeJobRun`, `DeploymentResourceTypePostgresCatalog`, `DeploymentResourceTypePostgresSyncedTable`, `DeploymentResourceTypeGenieSpace`, `DeploymentResourceTypeInstancePool`, `DeploymentResourceTypePostgresDatabase` and `DeploymentResourceTypePostgresRole` enum values for [bundledeployments.DeploymentResourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#DeploymentResourceType).
* Add `OperationStatusPending` enum value for [bundledeployments.OperationStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#OperationStatus).
* Add `Gpu1xH100` enum value for [compute.HardwareAcceleratorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#HardwareAcceleratorType).
* Add `Rabbitmq` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* [Breaking] Change `SourceSchema` field for [pipelines.SchemaSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SchemaSpec) to no longer be required.
* Change `SourceSchema` field for [pipelines.SchemaSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SchemaSpec) to no longer be required.
* [Breaking] Change `SourceTable` field for [pipelines.TableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpec) to no longer be required.
* Change `SourceTable` field for [pipelines.TableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpec) to no longer be required.
* [Breaking] Remove `EntityColumns` and `TimeseriesColumn` fields for [ml.DeltaTableSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeltaTableSource).
* [Breaking] Remove `FilterCondition`, `Inputs` and `TimeWindow` fields for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature).
* [Breaking] Remove `ExtraParameters` and `FunctionType` fields for [ml.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Function).
* [Breaking] Remove `EntityColumnIdentifiers` and `TimeseriesColumnIdentifier` fields for [ml.KafkaSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#KafkaSource).
* [Breaking] Remove `CronSchedule` field for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* [Breaking] Remove `Continuous` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
* Add `Model`, `ModelService`, `McpService` and `ModelProviderService` enum values for [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).