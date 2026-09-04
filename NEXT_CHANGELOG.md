# NEXT CHANGELOG

## Release v0.178.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add [aifunctions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aifunctions) and [sandbox](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sandbox) packages.
* Add [w.AiFunctions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aifunctions#AiFunctionsAPI) workspace-level service.
* Add [w.Sandbox](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sandbox#SandboxAPI) workspace-level service.
* Add `BackfillFeatures`, `CancelOperation` and `GetOperation` methods for [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service.
* Add `ProvisionedCapacityId` field for [jobs.ComputeSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ComputeSpec).
* Add `Mode` field for [ml.CronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CronSchedule).
* Add `LatestBackfillOperation` field for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* Add `ExcludedColumns` and `RecordTypeFilter` fields for [ml.Stream](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Stream).
* Add `DeploymentResourceTypeSecret` and `DeploymentResourceTypeClusterPolicy` enum values for [bundledeployments.DeploymentResourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#DeploymentResourceType).
* Add `AwsSecretsManager` and `AzureKeyVault` enum values for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `ExternalUseLocation` enum value for [catalog.Privilege](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#Privilege).
* Add `Gpu8xB300` enum value for [compute.HardwareAcceleratorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#HardwareAcceleratorType).
* [Breaking] Remove `Disabled` field for [catalog.InferenceTableConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#InferenceTableConfig).
* [Breaking] Remove `Owner` field for [catalog.McpService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#McpService).
* [Breaking] Remove `Owner` field for [catalog.ModelProviderService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderService).
* [Breaking] Remove `PlanType` field for [catalog.ModelProviderServiceConfigAnthropicProviderRelayedConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigAnthropicProviderRelayedConfig).
* [Breaking] Remove `Owner` field for [catalog.ModelService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelService).
* [Breaking] Remove `FirstTokenTimeout` and `TrafficSplitting` fields for [catalog.ModelServiceConfigRoutingConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelServiceConfigRoutingConfig).
* [Breaking] Remove `RequestTagKey` and `RequestTagValue` fields for [catalog.RateLimit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RateLimit).
* [Breaking] Remove `RateLimitKeyRequestTag` enum value for [catalog.RateLimitRateLimitKey](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RateLimitRateLimitKey).