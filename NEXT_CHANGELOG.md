# NEXT CHANGELOG

## Release v0.69.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added [w.DatabaseInstances](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DatabaseInstancesAPI) workspace-level service.
* Added [w.LlmProxyPartnerPoweredAccount](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#LlmProxyPartnerPoweredAccountAPI) workspace-level service, [w.LlmProxyPartnerPoweredEnforce](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#LlmProxyPartnerPoweredEnforceAPI) workspace-level service and [w.LlmProxyPartnerPoweredWorkspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#LlmProxyPartnerPoweredWorkspaceAPI) workspace-level service.
* Added `CreateProvisionedThroughputEndpoint` and `UpdateProvisionedThroughputEndpointConfig` methods for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
* Added `CatalogName` field for [catalog.EnableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EnableRequest).
* Added `SourceType` field for [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
* Added `Glob` field for [pipelines.PipelineLibrary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineLibrary).
* Added `DescribeQueryInvalidSqlError`, `DescribeQueryTimeout`, `DescribeQueryUnexpectedFailure`, `InvalidSqlMultipleDatasetReferencesException`, `InvalidSqlMultipleStatementsException` and `InvalidSqlUnknownTableException` enum values for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Added `CanCreate` and `CanMonitorOnly` enum values for [iam.PermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionLevel).
* Added `SuccessWithFailures` enum value for [jobs.TerminationCodeCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TerminationCodeCode).
* Added `InfrastructureMaintenance` enum value for [pipelines.StartUpdateCause](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdateCause).
* Added `InfrastructureMaintenance` enum value for [pipelines.UpdateInfoCause](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#UpdateInfoCause).
* [Breaking] Changed `CreateAlert` and `UpdateAlert` methods for [w.AlertsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsV2API) workspace-level service with new required argument order.
* [Breaking] Changed `Set` method for [w.Permissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionsAPI) workspace-level service . New request type is [iam.SetObjectPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#SetObjectPermissions).
* [Breaking] Changed `Update` method for [w.Permissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionsAPI) workspace-level service . New request type is [iam.UpdateObjectPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UpdateObjectPermissions).
* Changed `Schema` and `State` fields for [catalog.SystemSchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SystemSchemaInfo) to be required.
* [Breaking] Changed `State` field for [catalog.SystemSchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SystemSchemaInfo) to type `string`.
* [Breaking] Changed `Alert` field for [sql.CreateAlertV2Request](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateAlertV2Request) to be required.
* [Breaking] Changed `Alert` field for [sql.UpdateAlertV2Request](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateAlertV2Request) to be required.
* [Breaking] Removed `NodeTypeFlexibility` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
* [Breaking] Removed `NodeTypeFlexibility` field for [compute.GetInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePool).
* [Breaking] Removed `NodeTypeFlexibility` field for [compute.InstancePoolAndStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAndStats).
* [Breaking] Removed `Available`, `DisableInitialized`, `EnableCompleted`, `EnableInitialized` and `Unavailable` enum values for [catalog.SystemSchemaInfoState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SystemSchemaInfoState).
