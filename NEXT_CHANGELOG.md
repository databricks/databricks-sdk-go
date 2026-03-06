# NEXT CHANGELOG

## Release v0.119.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes
- **Auth cache**: Add a 1-minute backoff before retrying async token refresh after a failure. When async refresh fails, the cache no longer stays disabled until the next blocking call; it retries after the backoff, improving recovery from transient errors during long stale periods (up to 20 minutes).

### API Changes
