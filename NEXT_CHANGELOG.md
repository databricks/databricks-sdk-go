# NEXT CHANGELOG

## Release v0.130.0

### Breaking Changes

 * Remove the file-based OAuth token cache from `credentials/u2m/cache`. The removed symbols are `cache.NewFileTokenCache`, `cache.FileTokenCacheOption`, `cache.WithFileLocation`, and the private `tokenCacheFile` struct. The `TokenCache` interface, `ErrNotFound` sentinel, `HostCacheKeyProvider`, and `DiscoveryOAuthArgument` remain exported. `NewPersistentAuth` now defaults to a new in-memory cache (`cache.NewInMemoryTokenCache`) when no `WithTokenCache` option is passed; consumers that relied on the previous file-backed default must supply their own persistent cache. See databricks/cli#5056 for the companion CLI change that moves the file cache into the CLI.

### New Features and Improvements

### Bug Fixes

 * Fix CLI token source `--profile` fallback: `--profile` is a global Cobra flag that old CLIs accept silently instead of reporting "unknown flag", making the previous error-based detection dead code. Now uses `databricks version` to detect CLI capabilities at init time ([#1605](https://github.com/databricks/databricks-sdk-go/pull/1605)).

### Documentation

### Internal Changes

### API Changes
* Add [w.TemporaryVolumeCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TemporaryVolumeCredentialsAPI) workspace-level service.
* Add `GetPermissionLevels`, `GetPermissions`, `SetPermissions` and `UpdatePermissions` methods for [w.KnowledgeAssistants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/knowledgeassistants#KnowledgeAssistantsAPI) workspace-level service.
* Add `ThumbnailUrl` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Add `JiraOptions`, `OutlookOptions` and `SmartsheetOptions` fields for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `GoogleAdsConfig` field for [pipelines.SourceConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SourceConfig).
* Add `ReplaceExisting` field for [postgres.CreateBranchRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateBranchRequest).
* Add `ReplaceExisting` field for [postgres.CreateEndpointRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateEndpointRequest).