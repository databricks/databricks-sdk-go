# NEXT CHANGELOG

## Release v0.150.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Deduplicate identical key/value pairs in the user agent builder ([useragent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/useragent)), so repeatedly injecting the same dimension onto a reused `context.Context` no longer grows the `User-Agent` header without bound. Distinct values for the same key are still preserved.

### Documentation

### Internal Changes

### API Changes
