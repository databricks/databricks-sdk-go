# NEXT CHANGELOG

## Release v0.173.0

### Breaking Changes

 * Raise the minimum supported Go version from 1.24 to 1.25

### New Features and Improvements

### Bug Fixes
- Strip custom Databricks credential headers (`X-Databricks-Azure-SP-Management-Token`, `X-Databricks-GCP-SA-Access-Token`) on cross-host redirects, matching how `net/http` handles the `Authorization` header.

### Documentation

### Internal Changes

### API Changes
