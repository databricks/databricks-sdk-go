# NEXT CHANGELOG

## Release v0.115.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fixed `Config` race condition by replacing double-checked locking with `sync.Once` ([#1465](https://github.com/databricks/databricks-sdk-go/pull/1465)). This eliminates race detector warnings and enables proper concurrent usage of the SDK.

### Documentation

### Internal Changes
* Implement dynamic auth token stale period based on initial token lifetime. Increased up to 20 mins for standard OAuth with proportionally shorter periods for short-lived tokens.

### API Changes
