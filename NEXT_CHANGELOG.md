# NEXT CHANGELOG

## Release v0.158.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes
- Strip custom Databricks credential headers (`X-Databricks-Azure-SP-Management-Token`, `X-Databricks-GCP-SA-Access-Token`) on cross-host redirects, matching how `net/http` handles the `Authorization` header.

### Documentation

### Internal Changes

### API Changes
