# NEXT CHANGELOG

## Release v0.131.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

 * Alias `apierr.ErrResourceAlreadyExists` to `apierr.ErrAlreadyExists` so `errors.Is` matches both the gRPC-canonical `ALREADY_EXISTS` and the Databricks-specific `RESOURCE_ALREADY_EXISTS` error codes, regardless of which one the backend returns.

### Documentation

### Internal Changes

### API Changes
* Add `UndeleteProject` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `ConfidentialComputeType` field for [compute.GcpAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GcpAttributes).
* Add `Purge` field for [postgres.DeleteProjectRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#DeleteProjectRequest).
* Add `ShowDeleted` field for [postgres.ListProjectsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ListProjectsRequest).
* Add `DeleteTime` and `PurgeTime` fields for [postgres.Project](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Project).
* Add `UcConnection` field for [supervisoragents.Tool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#Tool).
* Change `Name` field for [supervisoragents.Connection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#Connection) to no longer be required.
* [Breaking] Change `Name` field for [supervisoragents.Connection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#Connection) to no longer be required.