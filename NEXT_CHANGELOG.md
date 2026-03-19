# NEXT CHANGELOG

## Release v0.126.0

### Breaking Changes

### New Features and Improvements
* Add `PersistentAuth.ForceRefreshToken()` to force-refresh cached U2M OAuth tokens and return an error instead of falling back to an existing token when refresh fails.
* Add `ErrMissingRefreshToken`, returned when a token refresh is requested but the cached token has no refresh token (i.e. `Token()` or `ForceRefreshToken()`).

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
