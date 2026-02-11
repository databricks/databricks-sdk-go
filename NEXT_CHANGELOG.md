# NEXT CHANGELOG

## Release v0.107.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `SourceType`, `UpdateTime` and `UpdatedBy` fields for [catalog.EntityTagAssignment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EntityTagAssignment).
* Add `SkipNotify` field for [dashboards.Subscription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Subscription).
* Add `SparseCheckout` field for [jobs.GitSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GitSource).
* Add `DisableAutoOptimization`, `MaxRetries`, `MinRetryIntervalMillis` and `RetryOnTimeout` fields for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Add `DisableAutoOptimization`, `MaxRetries`, `MinRetryIntervalMillis` and `RetryOnTimeout` fields for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Add `EdgegridAkamai` enum value for [catalog.CredentialType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CredentialType).