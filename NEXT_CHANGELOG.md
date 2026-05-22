# NEXT CHANGELOG

## Release v0.137.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

* Pass `excludedAttributes=entitlements` on the SCIM `/Me` request made by `WorkspaceClient.CurrentWorkspaceID` ([#1681](https://github.com/databricks/databricks-sdk-go/pull/1681)).

  `CurrentWorkspaceID` only reads the `X-Databricks-Org-Id` response header and discards the body, so it has no use for the `User.Entitlements` field. Skipping that attribute avoids an expensive `getEffectivePermissions` scan on the SCIM backend, which has caused incidents on workspaces with large grant counts.

### API Changes
* Add `Revert` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
* Add `ParentPath` field for [dashboards.GenieUpdateSpaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieUpdateSpaceRequest).