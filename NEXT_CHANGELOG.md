# NEXT CHANGELOG

## Release v0.101.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* Fix potential panic in `shouldRetry` function by using comma-ok idiom for type assertion ([#1438](https://github.com/databricks/databricks-sdk-go/pull/1438)).

### Documentation

### Internal Changes
* Remove deprecated `rand.Seed` call in retry backoff logic ([#1438](https://github.com/databricks/databricks-sdk-go/pull/1438)).

### API Changes
* Add `Force` field for [pipelines.DeletePipelineRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#DeletePipelineRequest).
* Add `PostgresRole` field for [postgres.RoleRoleSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#RoleRoleSpec).
* Add `PostgresRole` field for [postgres.RoleRoleStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#RoleRoleStatus).
* Add `Purge` field for [sql.TrashAlertV2Request](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TrashAlertV2Request).
