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
* Add `UpdateTokenManagement` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service.
* Add `DeploymentId` and `VersionId` fields for [jobs.JobDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobDeployment).
* Add `DeploymentId` and `VersionId` fields for [pipelines.PipelineDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineDeployment).
* Add `AutoscopeEnabled` field for [settings.CreateOboTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateOboTokenRequest).
* Add `AutoscopeEnabled` field for [settings.CreateTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateTokenRequest).
* Add `AutoscopeState`, `BackfillScopes`, `InferredScopes` and `Scopes` fields for [settings.PublicTokenInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#PublicTokenInfo).
* Add `AutoscopeState`, `BackfillScopes`, `InferredScopes` and `Scopes` fields for [settings.TokenInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenInfo).
* Add `ResourceType` field for [bundle.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle#Operation).
