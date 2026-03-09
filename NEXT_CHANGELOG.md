# NEXT CHANGELOG

## Release v0.119.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes
- **Auth cache**: Simplify experimental auth token refresh scheduling by replacing stale-duration/token-state tracking with an `asyncRefreshAllowedAfter` timestamp. Async refreshes now start based on a per-token dynamic lead time and, after a failed async refresh, are deferred by 1 minute before retrying while expiry still forces a blocking refresh.

### API Changes
