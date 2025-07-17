# NEXT CHANGELOG

## Release v0.76.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added `EnvironmentSettings` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
* Added `EnvironmentSettings` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
* Added `EnvironmentSettings` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
* Added `QueryBasedConnectorConfig` field for [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
* Added `MetaMarketing` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* [Breaking] Removed `Create` method for [w.Dashboards](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardsAPI) workspace-level service.
