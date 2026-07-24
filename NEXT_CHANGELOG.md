# NEXT CHANGELOG

## Release v0.164.0

### Breaking Changes

### New Features and Improvements

* Add the Pi coding agent (`PI_CODING_AGENT`) to AI agent detection in the User-Agent header, so CLI usage driven by Pi is attributed to `pi`.

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `SpecialDestinationSchemaOwner`, `SpecialDestinationTableOwner`, `SpecialDestinationVolumeOwner`, `SpecialDestinationFunctionOwner` and `SpecialDestinationRegisteredModelOwner` enum values for [catalog.SpecialDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SpecialDestination).
* [Breaking] Change `CreateDeployment` method for [w.BundleDeployments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#BundleDeploymentsAPI) workspace-level service with new required argument order.
* [Breaking] Remove `DeploymentId` field for [bundledeployments.CreateDeploymentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#CreateDeploymentRequest).