# NEXT CHANGELOG

## Release v0.126.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

 * Disable async token refresh for GCP credential providers to avoid wasted refresh attempts caused by double-caching with Google's internal `oauth2.ReuseTokenSource` ([#1549](https://github.com/databricks/databricks-sdk-go/issues/1549)).

### Documentation

### Internal Changes

### API Changes
