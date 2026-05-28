# NEXT CHANGELOG

## Release v0.139.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `CreateStream`, `DeleteStream`, `GetStream`, `ListStreams` and `UpdateStream` methods for [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service.
* Add `Parameters` field for [jobs.PipelineTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PipelineTask).
* Add `PipelineTask` field for [jobs.ResolvedValues](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ResolvedValues).
* Add `Parameters` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Add `Parameters` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Add `Parameters` field for [pipelines.GetPipelineResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelineResponse).
* [Breaking] Remove `CatalogId` field for [postgres.CatalogCatalogStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CatalogCatalogStatus).
* [Breaking] Remove `SyncedTableId` field for [postgres.SyncedTableSyncedTableStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#SyncedTableSyncedTableStatus).