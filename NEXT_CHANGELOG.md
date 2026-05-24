# NEXT CHANGELOG

## Release v0.137.0

### Breaking Changes

### New Features and Improvements

* Honor the Vercel `AI_AGENT=<name>` env var as a secondary fallback for
  AI agent detection in the User-Agent header (after the agents.md
  `AGENT=<name>` standard). Unrecognized fallback values now pass through
  the User-Agent sanitized and length-capped at 64 chars instead of being
  coerced to `agent/unknown`, so versioned variants such as
  `claude-code_2-1-141_agent` surface as-is.

### Bug Fixes

### Documentation

### Internal Changes

* Pass `excludedAttributes=entitlements` on the SCIM `/Me` request made by `WorkspaceClient.CurrentWorkspaceID` ([#1681](https://github.com/databricks/databricks-sdk-go/pull/1681)).

  `CurrentWorkspaceID` only reads the `X-Databricks-Org-Id` response header and discards the body, so it has no use for the `User.Entitlements` field. Skipping that attribute avoids an expensive `getEffectivePermissions` scan on the SCIM backend, which has caused incidents on workspaces with large grant counts.

### API Changes
* Add `Revert` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
* Add `ParentPath` field for [dashboards.GenieUpdateSpaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieUpdateSpaceRequest).
* Add `ComputeMaxInstances` and `ComputeMinInstances` fields for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Add `ComputeMaxInstances` and `ComputeMinInstances` fields for [apps.AppUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppUpdate).
* Add `CronScheduleTrigger`, `StreamingMode` and `TableTrigger` fields for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* Add `SyncedTableId` field for [postgres.SyncedTableSyncedTableStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#SyncedTableSyncedTableStatus).