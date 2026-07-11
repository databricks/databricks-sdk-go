# NEXT CHANGELOG

## Release v0.157.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `CreateCdfConfig`, `DeleteCdfConfig`, `GetCdfConfig`, `GetCdfStatus`, `ListCdfConfigs` and `ListCdfStatuses` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `Parent` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
* Add `Parent` field for [catalog.ListConnectionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListConnectionsRequest).
* Add `Minutes` enum value for [jobs.PeriodicTriggerConfigurationTimeUnit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PeriodicTriggerConfigurationTimeUnit).