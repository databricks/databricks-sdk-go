# NEXT CHANGELOG

## Release v0.125.0

### Breaking Changes

### New Features and Improvements
* Add `PersistentAuth.ForceRefreshToken()` to force-refresh cached U2M OAuth tokens and return an error instead of falling back to an existing token when refresh fails.
* Add `ErrMissingRefreshToken`, returned when a token refresh is requested but the cached token has no refresh token (i.e. `Token()` or `ForceRefreshToken()`).

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `DisableGovTagCreation` field for [settings.RestrictWorkspaceAdminsMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#RestrictWorkspaceAdminsMessage).
* Add `DisableGovTagCreation` field for [settingsv2.RestrictWorkspaceAdminsMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#RestrictWorkspaceAdminsMessage).
