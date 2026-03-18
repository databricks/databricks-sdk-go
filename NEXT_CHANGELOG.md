# NEXT CHANGELOG

## Release v0.124.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes
* Proactively attempt to refresh OAuth tokens expiring within 5 minutes in `PersistentAuth.Token()` to reduce the chance of callers receiving near-expired tokens. If the refresh fails and the token is still valid, the existing token is returned.
* When an expired token has no refresh token, `Token()` now returns `ErrMissingRefreshToken` instead of a less specific oauth2 error.

### API Changes
* Add `PersistentAuth.ForceRefreshToken()` to force-refresh cached U2M OAuth tokens and return an error instead of falling back to an existing token when refresh fails.
* Add `ParentPath` field for [dashboards.GenieSpace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSpace).