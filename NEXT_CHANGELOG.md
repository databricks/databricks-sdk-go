# NEXT CHANGELOG

## Release v0.164.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `SpecialDestinationSchemaOwner`, `SpecialDestinationTableOwner`, `SpecialDestinationVolumeOwner`, `SpecialDestinationFunctionOwner` and `SpecialDestinationRegisteredModelOwner` enum values for [catalog.SpecialDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SpecialDestination).
* [Breaking] Change `CreateDeployment` method for [w.BundleDeployments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#BundleDeploymentsAPI) workspace-level service with new required argument order.
* [Breaking] Remove `DeploymentId` field for [bundledeployments.CreateDeploymentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#CreateDeploymentRequest).
* Add `PatchTelemetryConfig` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
* Add `FanoutOptions` field for [pipelines.SchemaSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SchemaSpec).