# NEXT CHANGELOG

## Release v0.130.0

### Breaking Changes

 * Remove the file-based OAuth token cache from `credentials/u2m/cache`. The removed symbols are `cache.NewFileTokenCache`, `cache.FileTokenCacheOption`, `cache.WithFileLocation`, and the private `tokenCacheFile` struct. The `TokenCache` interface, `ErrNotFound` sentinel, `HostCacheKeyProvider`, and `DiscoveryOAuthArgument` remain exported. `NewPersistentAuth` now defaults to a new in-memory cache (`cache.NewInMemoryTokenCache`) when no `WithTokenCache` option is passed; consumers that relied on the previous file-backed default must supply their own persistent cache. See databricks/cli#5056 for the companion CLI change that moves the file cache into the CLI.

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
