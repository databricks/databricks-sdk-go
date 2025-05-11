# NEXT CHANGELOG

## Release v0.69.0

### New Features and Improvements

- Add support to authenticate with Account-wide token federation from the 
  following auth methods: `env-oidc`, `file-oidc`, and `github-oidc`.  

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added [a.LlmProxyPartnerPoweredAccount](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#LlmProxyPartnerPoweredAccountAPI) account-level service, [a.LlmProxyPartnerPoweredEnforce](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#LlmProxyPartnerPoweredEnforceAPI) account-level service and [w.LlmProxyPartnerPoweredWorkspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#LlmProxyPartnerPoweredWorkspaceAPI) workspace-level service.
* Added [w.DatabaseInstances](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DatabaseInstancesAPI) workspace-level service.
* Added `CreateProvisionedThroughputEndpoint` and `UpdateProvisionedThroughputEndpointConfig` methods for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
* Added `CatalogName` field for [catalog.EnableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EnableRequest).
* Added `SourceType` field for [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
* Added `Glob` field for [pipelines.PipelineLibrary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineLibrary).
* Added `ProvisionedModelUnits` field for [serving.ServedEntityInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityInput).
* Added `ProvisionedModelUnits` field for [serving.ServedEntityOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityOutput).
* Added `ProvisionedModelUnits` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput).
* Added `ProvisionedModelUnits` field for [serving.ServedModelOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelOutput).
* Added `DescribeQueryInvalidSqlError`, `DescribeQueryTimeout`, `DescribeQueryUnexpectedFailure`, `InvalidChatCompletionArgumentsJsonException`, `InvalidSqlMultipleDatasetReferencesException`, `InvalidSqlMultipleStatementsException` and `InvalidSqlUnknownTableException` enum values for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Added `CanCreate` and `CanMonitorOnly` enum values for [iam.PermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionLevel).
* Added `SuccessWithFailures` enum value for [jobs.TerminationCodeCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TerminationCodeCode).
* Added `InfrastructureMaintenance` enum value for [pipelines.StartUpdateCause](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdateCause).
* Added `InfrastructureMaintenance` enum value for [pipelines.UpdateInfoCause](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#UpdateInfoCause).
* [Breaking] Changed `CreateAlert` and `UpdateAlert` methods for [w.AlertsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsV2API) workspace-level service with new required argument order.
* [Breaking] Changed `Set` method for [w.Permissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionsAPI) workspace-level service . New request type is [iam.SetObjectPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#SetObjectPermissions).
* [Breaking] Changed `Update` method for [w.Permissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionsAPI) workspace-level service . New request type is [iam.UpdateObjectPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UpdateObjectPermissions).
* [Breaking] Changed `Get` method for [w.WorkspaceBindings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingsAPI) workspace-level service to return [catalog.GetCatalogWorkspaceBindingsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetCatalogWorkspaceBindingsResponse).
* [Breaking] Changed `GetBindings` method for [w.WorkspaceBindings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingsAPI) workspace-level service to return [catalog.GetWorkspaceBindingsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetWorkspaceBindingsResponse).
* [Breaking] Changed `Update` method for [w.WorkspaceBindings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingsAPI) workspace-level service to return [catalog.UpdateCatalogWorkspaceBindingsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalogWorkspaceBindingsResponse).
* [Breaking] Changed `UpdateBindings` method for [w.WorkspaceBindings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingsAPI) workspace-level service to return [catalog.UpdateWorkspaceBindingsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateWorkspaceBindingsResponse).
* [Breaking] Changed `SecurableType` field for [catalog.GetBindingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetBindingsRequest) to type `string`.
* Changed `Schema` and `State` fields for [catalog.SystemSchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SystemSchemaInfo) to be required.
* [Breaking] Changed `State` field for [catalog.SystemSchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SystemSchemaInfo) to type `string`.
* [Breaking] Changed `SecurableType` field for [catalog.UpdateWorkspaceBindingsParameters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateWorkspaceBindingsParameters) to type `string`.
* [Breaking] Changed `WorkspaceId` field for [catalog.WorkspaceBinding](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBinding) to be required.
* [Breaking] Changed `GpuNodePoolId` field for [jobs.ComputeConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ComputeConfig) to no longer be required.
* Changed `GpuNodePoolId` field for [jobs.ComputeConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ComputeConfig) to no longer be required.
* [Breaking] Changed `Alert` field for [sql.CreateAlertV2Request](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateAlertV2Request) to be required.
* [Breaking] Changed `Alert` field for [sql.UpdateAlertV2Request](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateAlertV2Request) to be required.
* [Breaking] Removed `NodeTypeFlexibility` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
* [Breaking] Removed `NodeTypeFlexibility` field for [compute.GetInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePool).
* [Breaking] Removed `NodeTypeFlexibility` field for [compute.InstancePoolAndStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAndStats).
* [Breaking] Removed `Catalog`, `Credential`, `ExternalLocation` and `StorageCredential` enum values for [catalog.GetBindingsSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetBindingsSecurableType).
* [Breaking] Removed `Available`, `DisableInitialized`, `EnableCompleted`, `EnableInitialized` and `Unavailable` enum values for [catalog.SystemSchemaInfoState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SystemSchemaInfoState).
* [Breaking] Removed `Catalog`, `Credential`, `ExternalLocation` and `StorageCredential` enum values for [catalog.UpdateBindingsSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateBindingsSecurableType).
