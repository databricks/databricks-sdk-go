# NEXT CHANGELOG

## Release v0.166.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add [w.AiGateway](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AiGatewayAPI) workspace-level service.
* Add `UpdateOperation` method for [w.BundleDeployments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#BundleDeploymentsAPI) workspace-level service.
* Add `LastSuccessfulVersionId` and `UpdatedBy` fields for [bundledeployments.Deployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Deployment).
* Add `DashboardMetadata`, `SequenceId` and `UpdateTime` fields for [bundledeployments.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Operation).
* Add `DashboardMetadata` field for [bundledeployments.Resource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Resource).
* Add `PreviousVersionId` field for [bundledeployments.Version](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Version).