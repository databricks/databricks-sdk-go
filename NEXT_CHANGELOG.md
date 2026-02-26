# NEXT CHANGELOG

## Release v0.115.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes
* Implement dynamic auth token stale period based on initial token lifetime. Increased up to 20 mins for standard OAuth with proportionally shorter periods for short-lived tokens.
* Move cloud-based credential filtering from individual strategies into `DefaultCredentials`. Azure strategies are skipped on GCP/AWS hosts in auto-detect mode; GCP strategies are skipped on Azure/AWS hosts. When `auth_type` is explicitly set (e.g. `azure-cli`), cloud filtering is bypassed so the named strategy is always attempted regardless of host cloud.

### API Changes
