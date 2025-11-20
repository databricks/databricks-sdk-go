# NEXT CHANGELOG

## Release v0.92.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `CreateSpace` and `UpdateSpace` methods for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Add `BatchCreateMaterializedFeatures`, `CreateKafkaConfig`, `DeleteKafkaConfig`, `GetKafkaConfig`, `ListKafkaConfigs` and `UpdateKafkaConfig` methods for [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service.
* Add `DeleteOnlineTable` method for [w.FeatureStore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureStoreAPI) workspace-level service.
* Add `RetrieveUserVisibleMetrics` method for [w.VectorSearchEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchEndpointsAPI) workspace-level service.
* Add `MajorVersion` field for [billing.CreateBillingUsageDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBillingUsageDashboardRequest).
* Add `IncludeSerializedSpace` field for [dashboards.GenieGetSpaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieGetSpaceRequest).
* Add `SerializedSpace` field for [dashboards.GenieSpace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSpace).
* Add `Purpose` field for [dashboards.TextAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TextAttachment).
* Add `BudgetPolicyId` field for [database.NewPipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#NewPipelineSpec).
* Add `Model` field for [jobs.TriggerSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerSettings).
* Add `KafkaSource` field for [ml.DataSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DataSource).
* Add `LineageContext` field for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature).
* Add `ConnectionParameters` field for [pipelines.IngestionGatewayPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionGatewayPipelineDefinition).
* Add `IngestFromUcForeignCatalog` field for [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
* Add `RewindSpec` field for [pipelines.StartUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdate).
* Add `TypeText` field for [vectorsearch.ColumnInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#ColumnInfo).
* Add `AutoscaleV2` enum value for [compute.EventDetailsCause](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventDetailsCause).
* Add `UnsupportedConversationTypeException` enum value for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Add `ForeignCatalog` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Add `Creating` and `CreateFailed` enum values for [settings.CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState).
* Add `Creating` and `CreateFailed` enum values for [settings.NccAzurePrivateEndpointRuleConnectionState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRuleConnectionState).
* Add `RedState` and `YellowState` enum values for [vectorsearch.EndpointStatusState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointStatusState).
* Change `Destinations` field for [catalog.AccessRequestDestinations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccessRequestDestinations) to no longer be required.
* [Breaking] Change `Destinations` field for [catalog.AccessRequestDestinations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccessRequestDestinations) to no longer be required.
* Change `TableNames` field for [jobs.TableUpdateTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TableUpdateTriggerConfiguration) to be required.
* [Breaking] Change `TableNames` field for [jobs.TableUpdateTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TableUpdateTriggerConfiguration) to be required.
* [Breaking] Change `OnlineStoreConfig` field for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature) to type [ml.OnlineStoreConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#OnlineStoreConfig).