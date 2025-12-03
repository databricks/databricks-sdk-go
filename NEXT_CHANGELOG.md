# NEXT CHANGELOG

## Release v0.93.0

### New Features and Improvements

### Bug Fixes

* Support for Spark versions with multiple Scala versions ([#1331](https://github.com/databricks/databricks-sdk-go/pull/1331)).

### documentation

### Internal Changes

### API Changes
* Add [w.WorkspaceEntityTagAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/tags#WorkspaceEntityTagAssignmentsAPI) workspace-level service.
* Add `DatasetCatalog` and `DatasetSchema` fields for [dashboards.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateDashboardRequest).
* Add `DatasetCatalog` and `DatasetSchema` fields for [dashboards.UpdateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#UpdateDashboardRequest).
* Add `PurgeData` field for [database.DeleteSyncedDatabaseTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DeleteSyncedDatabaseTableRequest).
* Add `Truncation` field for [pipelines.PipelineEvent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineEvent).
* Add `ForeignTable` and `Volume` enum values for [sharing.SharedDataObjectDataObjectType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharedDataObjectDataObjectType).