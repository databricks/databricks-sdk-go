# NEXT CHANGELOG

## Release v0.128.0

### Breaking Changes

### New Features and Improvements

 * Add `config.DefaultHostMetadataResolverFactory` — a package-level variable consulted when `Config.HostMetadataResolver` is unset. Lets programs install a shared resolver (e.g. a caching one) once from an `init()` block (typically in a blank-imported package) instead of wiring per-Config. Experimental.

### Bug Fixes

 * Add `X-Databricks-Org-Id` header to `Workspace.Download()` and `Workspace.Upload()` for SPOG host compatibility.
 * `WorkspaceClient.CurrentWorkspaceID()` now returns `Config.WorkspaceID` directly when set, instead of calling `/api/2.0/preview/scim/v2/Me`. This removes an API round-trip on every call where the workspace ID is already known (profile, `?o=` query param, env var, or host metadata) and fixes a failure on SPOG hosts where the unauthenticated probe request was rejected with `Unable to load OAuth Config`.
 * Add `X-Databricks-Org-Id` header to `SharesAPI.internalList()` (`service/sharing/ext_api.go`) for SPOG host compatibility. Same class of bug as #1634 - a hand-written extension method was calling `a.client.Do()` without the header that the generated code paths set per-call.

### Documentation

### Internal Changes

### API Changes
