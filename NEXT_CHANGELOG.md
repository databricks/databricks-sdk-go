# NEXT CHANGELOG

## Release v0.77.0

### New Features and Improvements

* Add new error value `ErrHTMLContent` to differentiate parsing errors caused by
  HTML content from other parsing errors.
* Return more detailed error messages when OAuth endpoints cannot be resolved.
* Use a free port in `u2m` authentication flows rather than 8020.

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added [agentbricks](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/agentbricks) package.
* Added [w.CleanRoomAssetRevisions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetRevisionsAPI) workspace-level service and [w.CleanRoomAutoApprovalRules](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAutoApprovalRulesAPI) workspace-level service.
* Added `CreateCleanRoomAssetReview` method for [w.CleanRoomAssets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetsAPI) workspace-level service.
* Added `LatestMonitorFailureMsg` field for [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor).
* Added `LatestMonitorFailureMsg` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor).
* Added `ProvisioningPhase` field for [database.SyncedTablePipelineProgress](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTablePipelineProgress).
* Added `Unspecified` enum value for [catalog.MonitorCronSchedulePauseStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorCronSchedulePauseStatus).
* Added `Unknown` enum value for [catalog.MonitorRefreshInfoState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorRefreshInfoState).
* Added `UnknownTrigger` enum value for [catalog.MonitorRefreshInfoTrigger](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorRefreshInfoTrigger).
* Added `Redshift` and `Sqldw` enum values for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Added `GermanyC5` enum value for [settings.ComplianceStandard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ComplianceStandard).
* [Breaking] Changed `CancelRefresh` method for [w.QualityMonitors](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#QualityMonitorsAPI) workspace-level service to start returning [catalog.CancelRefreshResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CancelRefreshResponse).
* [Breaking] Changed `Create` method for [w.QualityMonitors](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#QualityMonitorsAPI) workspace-level service with new required argument order.
* [Breaking] Changed `Delete` method for [w.QualityMonitors](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#QualityMonitorsAPI) workspace-level service to start returning [catalog.DeleteMonitorResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteMonitorResponse).
* [Breaking] Changed `RefreshId` field for [catalog.CancelRefreshRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CancelRefreshRequest) to type `int64`.
* [Breaking] Changed `RefreshId` field for [catalog.GetRefreshRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRefreshRequest) to type `int64`.
* [Breaking] Changed `MonitorVersion` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to type `int64`.
* Changed `OutputSchemaName` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to be required.
* [Breaking] Changed `AssetType` and `Name` fields for [cleanrooms.CleanRoomAsset](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAsset) to be required.
* Changed `AssetType` and `Name` fields for [cleanrooms.CleanRoomAsset](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAsset) to be required.
* [Breaking] Changed `LocalName` field for [cleanrooms.CleanRoomAssetForeignTableLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetForeignTableLocalDetails) to be required.
* Changed `LocalName` field for [cleanrooms.CleanRoomAssetForeignTableLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetForeignTableLocalDetails) to be required.
* [Breaking] Changed `NotebookContent` field for [cleanrooms.CleanRoomAssetNotebook](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetNotebook) to be required.
* Changed `NotebookContent` field for [cleanrooms.CleanRoomAssetNotebook](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetNotebook) to be required.
* [Breaking] Changed `LocalName` field for [cleanrooms.CleanRoomAssetTableLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetTableLocalDetails) to be required.
* Changed `LocalName` field for [cleanrooms.CleanRoomAssetTableLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetTableLocalDetails) to be required.
* [Breaking] Changed `LocalName` field for [cleanrooms.CleanRoomAssetViewLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetViewLocalDetails) to be required.
* Changed `LocalName` field for [cleanrooms.CleanRoomAssetViewLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetViewLocalDetails) to be required.
* [Breaking] Changed `LocalName` field for [cleanrooms.CleanRoomAssetVolumeLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetVolumeLocalDetails) to be required.
* Changed `LocalName` field for [cleanrooms.CleanRoomAssetVolumeLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetVolumeLocalDetails) to be required.
* [Breaking] Removed [aibuilder](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aibuilder) package.