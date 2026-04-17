# NEXT CHANGELOG

## Release v0.128.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

 * Add `X-Databricks-Org-Id` header to `Workspace.Download()` and `Workspace.Upload()` for SPOG host compatibility.
 * `WorkspaceClient.CurrentWorkspaceID()` now returns `Config.WorkspaceID` directly when set, instead of calling `/api/2.0/preview/scim/v2/Me`. This removes an API round-trip on every call where the workspace ID is already known (profile, `?o=` query param, env var, or host metadata) and fixes a failure on SPOG hosts where the unauthenticated probe request was rejected with `Unable to load OAuth Config`.

### Documentation

### Internal Changes

### API Changes
