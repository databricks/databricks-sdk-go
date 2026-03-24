# NEXT CHANGELOG

## Release v0.127.0

### Breaking Changes

 * Raise minimum Go version from 1.18 to 1.24 ([#1558](https://github.com/databricks/databricks-sdk-go/pull/1558)).

### New Features and Improvements

* Added `HostMetadataResolver` hook to allow callers to customize host metadata resolution, e.g. with caching ([#1572](https://github.com/databricks/databricks-sdk-go/pull/1572)).
* Added `NewLimitIterator` to `listing` package for lazy iteration with a cap on output items ([#1555](https://github.com/databricks/databricks-sdk-go/pull/1555)).

### Bug Fixes

 * Fix double-caching of OAuth tokens in Azure client secret credentials ([#1549](https://github.com/databricks/databricks-sdk-go/issues/1549)).
 * Disable async token refresh for GCP credential providers to avoid wasted refresh attempts caused by double-caching with Google's internal `oauth2.ReuseTokenSource` ([#1549](https://github.com/databricks/databricks-sdk-go/issues/1549)).

### Documentation

### Internal Changes

 * Normalize internal token sources on `auth.TokenSource` for proper context propagation ([#1577](https://github.com/databricks/databricks-sdk-go/pull/1577)).
 * Bump golang.org/x/crypto from 0.21.0 to 0.45.0 in /examples/slog ([#1566](https://github.com/databricks/databricks-sdk-go/pull/1566)).
 * Bump golang.org/x/net from 0.23.0 to 0.33.0 in /examples/slog ([#1127](https://github.com/databricks/databricks-sdk-go/pull/1127)).
 * Bump golang.org/x/oauth2 from 0.20.0 to 0.27.0 ([#1563](https://github.com/databricks/databricks-sdk-go/pull/1563)).
 * Bump golang.org/x/crypto from 0.21.0 to 0.45.0 in /examples/zerolog ([#1568](https://github.com/databricks/databricks-sdk-go/pull/1568)).
 * Bump golang.org/x/crypto from 0.31.0 to 0.45.0 ([#1564](https://github.com/databricks/databricks-sdk-go/pull/1564)).
 * Bump google.golang.org/grpc from 1.58.3 to 1.79.3 in /examples/zerolog ([#1565](https://github.com/databricks/databricks-sdk-go/pull/1565)).
 * Bump google.golang.org/grpc from 1.62.0 to 1.79.3 in /examples/slog ([#1569](https://github.com/databricks/databricks-sdk-go/pull/1569)).
 * Bump google.golang.org/grpc from 1.64.1 to 1.79.3 ([#1557](https://github.com/databricks/databricks-sdk-go/pull/1557)).

### API Changes
* Add `CreateCatalog`, `CreateSyncedTable`, `DeleteCatalog`, `DeleteSyncedTable`, `GetCatalog` and `GetSyncedTable` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `EffectiveFileEventQueue` field for [catalog.CreateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateExternalLocation).
* Add `EffectiveFileEventQueue` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo).
* Add `EffectiveFileEventQueue` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
* Add `ColumnSelection` field for [ml.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Function).
* Add `Cascade` field for [pipelines.DeletePipelineRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#DeletePipelineRequest).
* Add `DefaultBranch` field for [postgres.ProjectSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectSpec).
* Add `DefaultBranch` field for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* Add `Ingress` and `IngressDryRun` fields for [settings.AccountNetworkPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicy).