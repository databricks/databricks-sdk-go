# NEXT CHANGELOG

## Release v0.91.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `BatchCreateMaterializedFeatures` method for [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service.
* Add `RetrieveUserVisibleMetrics` method for [w.VectorSearchEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchEndpointsAPI) workspace-level service.
* Add `Purpose` field for [dashboards.TextAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TextAttachment).
* Add `LineageContext` field for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature).
* Add `ConnectionParameters` field for [pipelines.IngestionGatewayPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionGatewayPipelineDefinition).
* Add `IngestFromUcForeignCatalog` field for [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
* Add `AutoscaleV2` enum value for [compute.EventDetailsCause](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventDetailsCause).
* Add `UnsupportedConversationTypeException` enum value for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Add `Creating` and `CreateFailed` enum values for [settings.CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState).
* Add `Creating` and `CreateFailed` enum values for [settings.NccAzurePrivateEndpointRuleConnectionState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRuleConnectionState).
* Add `RedState` and `YellowState` enum values for [vectorsearch.EndpointStatusState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointStatusState).
* [Breaking] Change `Destinations` field for [catalog.AccessRequestDestinations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccessRequestDestinations) to no longer be required.
* Change `Destinations` field for [catalog.AccessRequestDestinations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccessRequestDestinations) to no longer be required.
* [Breaking] Change `TableNames` field for [jobs.TableUpdateTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TableUpdateTriggerConfiguration) to be required.
* Change `TableNames` field for [jobs.TableUpdateTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TableUpdateTriggerConfiguration) to be required.
* [Breaking] Change `OnlineStoreConfig` field for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature) to type [ml.OnlineStoreConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#OnlineStoreConfig).