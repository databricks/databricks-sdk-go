# NEXT CHANGELOG

## Release v0.67.0

### New Features and Improvements

### Bug Fixes
* Fixed the deserialization of responses in VectorSearchAPI's `QueryIndex()` method ([#1214](https://github.com/databricks/databricks-sdk-py/pull/1214)).

### Documentation

### Internal Changes

### API Changes
* Added `FutureFeatureDataPath` field for [ml.CreateForecastingExperimentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CreateForecastingExperimentRequest).
* Added `ExcludeColumns` and `IncludeColumns` fields for [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
* Added `NetworkCheckControlPlaneFailure`, `NetworkCheckDnsServerFailure`, `NetworkCheckMetadataEndpointFailure`, `NetworkCheckMultipleComponentsFailure`, `NetworkCheckNicFailure`, `NetworkCheckStorageFailure` and `SecretPermissionDenied` enum values for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* [Breaking] Changed [vectorsearch.ListValue](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#ListValue) to.
* [Breaking] Changed `PipelineId` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline) to be required.
* Changed `ConnectionName`, `GatewayStorageCatalog` and `GatewayStorageSchema` fields for [pipelines.IngestionGatewayPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionGatewayPipelineDefinition) to be required.
* [Breaking] Changed `ConnectionName`, `GatewayStorageCatalog` and `GatewayStorageSchema` fields for [pipelines.IngestionGatewayPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionGatewayPipelineDefinition) to be required.
* [Breaking] Changed `Kind` field for [pipelines.PipelineDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineDeployment) to be required.
* Changed `Kind` field for [pipelines.PipelineDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineDeployment) to be required.
* [Breaking] Changed `DestinationCatalog`, `DestinationSchema` and `SourceUrl` fields for [pipelines.ReportSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ReportSpec) to be required.
* Changed `DestinationCatalog`, `DestinationSchema` and `SourceUrl` fields for [pipelines.ReportSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ReportSpec) to be required.
* [Breaking] Changed `DestinationCatalog`, `DestinationSchema` and `SourceSchema` fields for [pipelines.SchemaSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SchemaSpec) to be required.
* Changed `DestinationCatalog`, `DestinationSchema` and `SourceSchema` fields for [pipelines.SchemaSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SchemaSpec) to be required.
* Changed `DestinationCatalog`, `DestinationSchema` and `SourceTable` fields for [pipelines.TableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpec) to be required.
* [Breaking] Changed `DestinationCatalog`, `DestinationSchema` and `SourceTable` fields for [pipelines.TableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpec) to be required.
* [Breaking] Changed pagination for [AlertsV2API.ListAlerts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsV2API.ListAlerts).
* [Breaking] Changed waiter for [GenieAPI.CreateMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI.CreateMessage).
