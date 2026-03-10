# NEXT CHANGELOG

## Release v0.120.0

### Breaking Changes

### New Features and Improvements

* Added discovery login flow for `login.databricks.com` to the U2M OAuth infrastructure. This enables the CLI to authenticate users without requiring a workspace URL upfront.

### Bug Fixes

### Documentation

### Internal Changes
- **Auth cache**: Simplify experimental auth token refresh scheduling by replacing stale-duration/token-state tracking with an `asyncRefreshAllowedAfter` timestamp. Async refreshes now start based on a per-token dynamic lead time and, after a failed async refresh, are deferred by 1 minute before retrying while expiry still forces a blocking refresh.

### API Changes
