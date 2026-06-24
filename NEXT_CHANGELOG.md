# NEXT CHANGELOG

## Release v0.150.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Deduplicate identical key/value pairs in the user agent builder ([useragent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/useragent)), so repeatedly injecting the same dimension onto a reused `context.Context` no longer grows the `User-Agent` header without bound. Distinct values for the same key are still preserved.

### Documentation

### Internal Changes

### API Changes
* Add `CancelPendingClusterEnforcement` method for [w.PolicyComplianceForClusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#PolicyComplianceForClustersAPI) workspace-level service.
* Add `PendingEnforcement` field for [compute.ClusterCompliance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterCompliance).
* Add `EnforceMode` field for [compute.EnforceClusterComplianceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EnforceClusterComplianceRequest).
* Add `EnforceResult` field for [compute.EnforceClusterComplianceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EnforceClusterComplianceResponse).
* Add `PendingEnforcement` field for [compute.GetClusterComplianceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterComplianceResponse).
* Add `DeferredPolicyEnforcementScheduled` and `DeferredPolicyEnforcementFailed` enum values for [compute.EventType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventType).
* Change `ReplicateWorkspaceAssets` field for [disasterrecovery.WorkspaceSet](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery#WorkspaceSet) to no longer be required.
* [Breaking] Change `ReplicateWorkspaceAssets` field for [disasterrecovery.WorkspaceSet](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery#WorkspaceSet) to no longer be required.