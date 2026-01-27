# NEXT CHANGELOG

## Release v0.102.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `CloneMode` field for [pipelines.ClonePipelineRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ClonePipelineRequest).
* [Breaking] Change `CreateRole` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service with new required argument order.
* Change `RoleId` field for [postgres.CreateRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateRoleRequest) to no longer be required.