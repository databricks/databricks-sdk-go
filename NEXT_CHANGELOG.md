# NEXT CHANGELOG

## Release v0.162.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes
- Strip custom Databricks credential headers (`X-Databricks-Azure-SP-Management-Token`, `X-Databricks-GCP-SA-Access-Token`) on cross-host redirects, matching how `net/http` handles the `Authorization` header.

### Documentation

### Internal Changes

### API Changes
* Add `EffectiveEntitlements` field for [iamv2.WorkspaceAssignmentDetail](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceAssignmentDetail).
* Add `ServerlessComputeId` field for [jobs.JobCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobCluster).