# NEXT CHANGELOG

## Release v0.127.0

### Breaking Changes

 * Raise minimum Go version from 1.18 to 1.24 ([#XXXX](https://github.com/databricks/databricks-sdk-go/pull/XXXX)).

### New Features and Improvements

* Added `HostMetadataResolver` interface to allow callers to customize host metadata resolution, e.g. with caching ([#1572](https://github.com/databricks/databricks-sdk-go/pull/1572)).
* Added `NewLimitIterator` to `listing` package for lazy iteration with a cap on output items ([#1555](https://github.com/databricks/databricks-sdk-go/pull/1555)).

### Bug Fixes

### Documentation

### Internal Changes

 * Bump golang.org/x/crypto from 0.21.0 to 0.45.0 in /examples/slog ([#1566](https://github.com/databricks/databricks-sdk-go/pull/1566)).
 * Bump golang.org/x/oauth2 from 0.20.0 to 0.27.0 ([#1563](https://github.com/databricks/databricks-sdk-go/pull/1563)).
 * Bump golang.org/x/crypto from 0.21.0 to 0.45.0 in /examples/zerolog ([#1568](https://github.com/databricks/databricks-sdk-go/pull/1568)).

### API Changes
* Add `CreateCatalog`, `CreateSyncedTable`, `DeleteCatalog`, `DeleteSyncedTable`, `GetCatalog` and `GetSyncedTable` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `EffectiveFileEventQueue` field for [catalog.CreateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateExternalLocation).
* Add `EffectiveFileEventQueue` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo).
* Add `EffectiveFileEventQueue` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
* Add `ColumnSelection` field for [ml.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Function).
* Add `Cascade` field for [pipelines.DeletePipelineRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#DeletePipelineRequest).
* Add `DefaultBranch` field for [postgres.ProjectSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectSpec).
* Add `DefaultBranch` field for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
