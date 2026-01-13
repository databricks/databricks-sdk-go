# NEXT CHANGELOG

## Release v0.97.0

### New Features and Improvements

* Increase async cache stale period from 3 to 5 minutes to cover the maximum monthly downtime of a 99.99% uptime SLA.

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `ErrorMessage` field for [settings.CreatePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreatePrivateEndpointRule).
* Add `ErrorMessage` field for [settings.NccPrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccPrivateEndpointRule).
* Add `ErrorMessage` field for [settings.UpdatePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdatePrivateEndpointRule).
* Add `RateLimited` enum value for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Add `Creating` and `CreateFailed` enum values for [settings.NccPrivateEndpointRulePrivateLinkConnectionState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccPrivateEndpointRulePrivateLinkConnectionState).
* Add `RateLimited` enum value for [sql.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TerminationReasonCode).
* [Breaking] Add long-running operation for [PostgresAPI.DeleteBranch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI.DeleteBranch).
* [Breaking] Add long-running operation for [PostgresAPI.DeleteEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI.DeleteEndpoint).
* [Breaking] Add long-running operation for [PostgresAPI.DeleteProject](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI.DeleteProject).
* [Breaking] Change `DeleteBranch`, `DeleteEndpoint` and `DeleteProject` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service to return [postgres.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Operation).
* [Breaking] Remove `PgbouncerSettings` field for [postgres.EndpointSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointSettings).
* [Breaking] Remove `PoolerMode` field for [postgres.EndpointSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointSpec).
* [Breaking] Remove `PoolerMode` field for [postgres.EndpointStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatus).
* [Breaking] Remove `PgbouncerSettings` field for [postgres.ProjectDefaultEndpointSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectDefaultEndpointSettings).
* Add `Owner` field for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* Add `ValidityCheckConfigurations` field for [qualitymonitorv2.QualityMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/qualitymonitorv2#QualityMonitor).