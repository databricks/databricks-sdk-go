# NEXT CHANGELOG

## Release v0.112.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* Fix potential panic in `shouldRetry` function by using comma-ok idiom for type assertion ([#1438](https://github.com/databricks/databricks-sdk-go/pull/1438)).

### Documentation

### Internal Changes
* Remove deprecated `rand.Seed` call in retry backoff logic ([#1438](https://github.com/databricks/databricks-sdk-go/pull/1438)).

### API Changes
