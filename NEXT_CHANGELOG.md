# NEXT CHANGELOG

## Release v0.70.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added [a.NetworkPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkPoliciesAPI) account-level service and [a.WorkspaceNetworkConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#WorkspaceNetworkConfigurationAPI) account-level service.
* Added [w.RecipientFederationPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#RecipientFederationPoliciesAPI) workspace-level service.
* Added `CreateLoggedModel`, `DeleteLoggedModel`, `DeleteLoggedModelTag`, `FinalizeLoggedModel`, `GetLoggedModel`, `ListLoggedModelArtifacts`, `LogLoggedModelParams`, `LogOutputs`, `SearchLoggedModels` and `SetLoggedModelTags` methods for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
* Added `UcSecurable` field for [apps.AppResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResource).
* Added `TimeseriesColumns` field for [catalog.PrimaryKeyConstraint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#PrimaryKeyConstraint).
* Added `ReviewState`, `Reviews` and `RunnerCollaboratorAliases` fields for [cleanrooms.CleanRoomAssetNotebook](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetNotebook).
* Added `NotebookEtag` and `NotebookUpdatedAt` fields for [cleanrooms.CleanRoomNotebookTaskRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomNotebookTaskRun).
* Added `RootPath` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Added `RootPath` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Added `RootPath` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
* Added `MaterializationNamespace` field for [sharing.Table](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#Table).
* Added `OmitPermissionsList` field for [sharing.UpdateSharePermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateSharePermissions).
* Added `AutoResolveDisplayName` field for [sql.UpdateAlertRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateAlertRequest).
* Added `AutoResolveDisplayName` field for [sql.UpdateQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateQueryRequest).
* Added `InternalCatalog`, `ManagedOnlineCatalog` and `UnknownCatalogType` enum values for [catalog.CatalogType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogType).
* Added `Catalog`, `CleanRoom`, `Connection`, `Credential`, `ExternalLocation`, `ExternalMetadata`, `Function`, `Metastore`, `Pipeline`, `Provider`, `Recipient`, `Schema`, `Share`, `StagingTable`, `StorageCredential`, `Table`, `UnknownSecurableType` and `Volume` enum values for [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).
* [Breaking] Changed `SecurableType` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo) to type [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).
* Changed `Etag` and `Name` fields for [iam.RuleSetResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#RuleSetResponse) to be required.
