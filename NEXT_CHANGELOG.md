# NEXT CHANGELOG

## Release v0.187.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `CreateIdentityVisibilityFilter`, `DeleteIdentityVisibilityFilter`, `GetIdentityVisibilityFilter` and `ListIdentityVisibilityFilters` methods for [a.AccountIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API) account-level service.
* Add `EndpointRoute` field for [catalog.ModelProviderServiceConfigModelTargetConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigModelTargetConfig).
* Add `Notifications` field for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* Add `ShufflePartitions` field for [ml.StreamingMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#StreamingMode).
* Add `ParentPath` field for [pipelines.ClonePipelineRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ClonePipelineRequest).
* Add `ParentPath` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Add `ParentPath` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Add `ParentPath` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
* Add `CustomTemplateFormat` field for [sql.AlertV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2).
* Add `StatementTimeout` field for [sql.CreateWarehouseRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateWarehouseRequest).
* Add `StatementTimeout` field for [sql.EditWarehouseRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#EditWarehouseRequest).
* Add `StatementTimeout` field for [sql.EndpointInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#EndpointInfo).
* Add `StatementTimeout` field for [sql.GetWarehouseResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetWarehouseResponse).