# NEXT CHANGELOG

## Release v0.115.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes
* Implement dynamic auth token stale period based on initial token lifetime. Increased up to 20 mins for standard OAuth with proportionally shorter periods for short-lived tokens.

### API Changes
* Add `CreateDatabase`, `DeleteDatabase`, `GetDatabase`, `ListDatabases` and `UpdateDatabase` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `Postgres` field for [apps.AppResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResource).
* Add `DataframeSchema`, `FilterCondition` and `TransformationSql` fields for [ml.DeltaTableSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeltaTableSource).
* Add `EnvironmentVersion` field for [pipelines.PipelinesEnvironment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesEnvironment).
* Add `ResetCheckpointSelection` field for [pipelines.StartUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdate).
* [Breaking] Remove `Oauth2AppClientId` and `Oauth2AppIntegrationId` fields for [apps.Space](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#Space).
* Add `EnablePgNativeLogin` field for [postgres.ProjectSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectSpec).
* Add `EnablePgNativeLogin` field for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* [Breaking] Remove `NodeTypeFlexibility` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).