# NEXT CHANGELOG

## Release v0.102.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes
* Add integration test support for unified host infrastructure ([#1445](https://github.com/databricks/databricks-sdk-go/pull/1445))

### API Changes
* Add `CloneMode` field for [pipelines.ClonePipelineRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ClonePipelineRequest).
* [Breaking] Change `CreateRole` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service with new required argument order.
* Change `RoleId` field for [postgres.CreateRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateRoleRequest) to no longer be required.