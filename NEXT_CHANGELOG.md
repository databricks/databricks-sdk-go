# NEXT CHANGELOG

## Release v0.105.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fixed `Config` race condition by replacing double-checked locking with `sync.Once` ([#1465](https://github.com/databricks/databricks-sdk-go/pull/1465)). This eliminates race detector warnings and enables proper concurrent usage of the SDK.

### Documentation

### Internal Changes

### API Changes
