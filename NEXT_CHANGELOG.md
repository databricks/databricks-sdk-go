# NEXT CHANGELOG

## Release v0.130.0

### Breaking Changes

* Remove the `Experimental_IsUnifiedHost` field (and the `DATABRICKS_EXPERIMENTAL_IS_UNIFIED_HOST` environment variable) from `Config`. Unified host detection is now automatic via the `/.well-known/databricks-config` endpoint.
* Remove the unused `ErrWorkspaceIDInAccountClient` exported variable. It was never returned from any production path, and its message contradicted the unified host workflow where a single profile with both `AccountID` and `WorkspaceID` produces both clients.
* Remove the file-based OAuth token cache from `credentials/u2m/cache`. The removed symbols are `cache.NewFileTokenCache`, `cache.FileTokenCacheOption`, `cache.WithFileLocation`, and the private `tokenCacheFile` struct. The `TokenCache` interface, `ErrNotFound` sentinel, `HostCacheKeyProvider`, and `DiscoveryOAuthArgument` remain exported. `NewPersistentAuth` now defaults to a new in-memory cache (`cache.NewInMemoryTokenCache`) when no `WithTokenCache` option is passed; consumers that relied on the previous file-backed default must supply their own persistent cache. See databricks/cli#5056 for the companion CLI change that moves the file cache into the CLI.

### New Features and Improvements

* Add `u2m.WithDiscoveryHost` option to override the default `https://login.databricks.com` host used by the discovery login flow. Intended for testing and development against non-production environments.
* Add support for unified hosts. A single configuration profile can now be used for both account-level and workspace-level operations when the host supports it and both `AccountID` and `WorkspaceID` are available.

### Bug Fixes

 * Fix CLI token source `--profile` fallback: `--profile` is a global Cobra flag that old CLIs accept silently instead of reporting "unknown flag", making the previous error-based detection dead code. Now uses `databricks version` to detect CLI capabilities at init time ([#1605](https://github.com/databricks/databricks-sdk-go/pull/1605)).

### Documentation

### Internal Changes

 * Pass `--force-refresh` to Databricks CLI `auth token` command to bypass the CLI's internal token cache ([#1628](https://github.com/databricks/databricks-sdk-go/pull/1628)).

### API Changes
* Add [w.TemporaryVolumeCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TemporaryVolumeCredentialsAPI) workspace-level service.
* Add `GetPermissionLevels`, `GetPermissions`, `SetPermissions` and `UpdatePermissions` methods for [w.KnowledgeAssistants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/knowledgeassistants#KnowledgeAssistantsAPI) workspace-level service.
* Add `ThumbnailUrl` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Add `JiraOptions`, `OutlookOptions` and `SmartsheetOptions` fields for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `GoogleAdsConfig` field for [pipelines.SourceConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SourceConfig).
* Add `ReplaceExisting` field for [postgres.CreateBranchRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateBranchRequest).
* Add `ReplaceExisting` field for [postgres.CreateEndpointRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateEndpointRequest).