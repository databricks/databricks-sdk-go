# NEXT CHANGELOG

## Release v0.101.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `Force` field for [pipelines.DeletePipelineRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#DeletePipelineRequest).
* Add `PostgresRole` field for [postgres.RoleRoleSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#RoleRoleSpec).
* Add `PostgresRole` field for [postgres.RoleRoleStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#RoleRoleStatus).
* Add `Purge` field for [sql.TrashAlertV2Request](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TrashAlertV2Request).
* [Breaking] Change `CreateRole` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service with new required argument order.
* Change `RoleId` field for [postgres.CreateRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateRoleRequest) to no longer be required.