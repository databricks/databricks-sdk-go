# NEXT CHANGELOG

## Release v0.135.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

* Switch workspace addressing header on workspace-scoped API calls from `X-Databricks-Org-Id` to `X-Databricks-Workspace-Id`. The value continues to come from `Config.WorkspaceID` (`DATABRICKS_WORKSPACE_ID`), and now accepts either a classic numeric workspace ID or a CPDR connection ID (server disambiguates). Part of the DECO + Unified Workspace Addressing initiative.

### API Changes
