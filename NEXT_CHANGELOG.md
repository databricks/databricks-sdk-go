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
