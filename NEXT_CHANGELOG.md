# NEXT CHANGELOG

## Release v0.140.0

### Breaking Changes

### New Features and Improvements

* Add `u2m.WithDiscoveryAccountTarget()` option that sets `target=ACCOUNT` on the login.databricks.com authorize URL, so the discovery flow lands the user on the account selector instead of the workspace selector.

### Bug Fixes

### Documentation

### Internal Changes

* Switch workspace addressing header on workspace-scoped API calls from `X-Databricks-Org-Id` to `X-Databricks-Workspace-Id`. The value continues to come from `Config.WorkspaceID` (`DATABRICKS_WORKSPACE_ID`), and now accepts either a classic numeric workspace ID or a CPDR connection ID (server disambiguates). Part of the DECO + Unified Workspace Addressing initiative.

### API Changes
