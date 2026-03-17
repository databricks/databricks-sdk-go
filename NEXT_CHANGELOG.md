# NEXT CHANGELOG

## Release v0.123.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes
* Proactively attempt to refresh OAuth tokens expiring within 5 minutes in `PersistentAuth.Token()` to reduce the chance of callers receiving near-expired tokens. If the refresh fails and the token is still valid, the existing token is returned.

### API Changes
