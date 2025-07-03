# NEXT CHANGELOG

## Release v0.74.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added [w.ExternalLineage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLineageAPI) workspace-level service and [w.ExternalMetadata](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalMetadataAPI) workspace-level service.
* Added [w.MaterializedFeatures](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeaturesAPI) workspace-level service.
* Added `DeleteConversation`, `ListConversations` and `TrashSpace` methods for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Added `CreateDatabaseInstanceRole`, `DeleteDatabaseInstanceRole`, `GetDatabaseInstanceRole` and `ListDatabaseInstanceRoles` methods for [w.Database](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseAPI) workspace-level service.
* Added `Connection` and `Credential` fields for [catalog.Dependency](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#Dependency).
* Added `Rely` field for [catalog.ForeignKeyConstraint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ForeignKeyConstraint).
* Added `Rely` field for [catalog.PrimaryKeyConstraint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#PrimaryKeyConstraint).
* Added `SecurableKindManifest` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).
* Added `SecurableKindManifest` field for [catalog.TableSummary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableSummary).
* Added `ChildInstanceRefs`, `EffectiveEnableReadableSecondaries`, `EffectiveNodeCount`, `EffectiveRetentionWindowInDays`, `EnableReadableSecondaries`, `NodeCount`, `ParentInstanceRef`, `ReadOnlyDns` and `RetentionWindowInDays` fields for [database.DatabaseInstance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstance).
* Added `Claims` field for [database.GenerateDatabaseCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#GenerateDatabaseCredentialRequest).
* Added `LastSync` field for [database.SyncedTableStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableStatus).
* Added `Activity` field for [ml.DeleteTransitionRequestResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeleteTransitionRequestResponse).
* Added `MaxResults` field for [ml.ListWebhooksRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ListWebhooksRequest).
* Added `Body` and `StatusCode` fields for [ml.TestRegistryWebhookResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TestRegistryWebhookResponse).
* Added `ModelVersionDatabricks` field for [ml.TransitionStageResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TransitionStageResponse).
* Added `RegisteredModel` field for [ml.UpdateModelResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#UpdateModelResponse).
* Added `ModelVersion` field for [ml.UpdateModelVersionResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#UpdateModelVersionResponse).
* Added `Webhook` field for [ml.UpdateWebhookResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#UpdateWebhookResponse).
* Added `RunAs` field for [pipelines.GetPipelineResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelineResponse).
* Added `Principal` field for [serving.AiGatewayRateLimit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayRateLimit).
* Added `Description` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint).
* Added `ServedEntityName` field for [serving.Route](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#Route).
* Added `AnyStaticCredential` enum value for [catalog.CredentialType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CredentialType).
* Added `DatabricksRowStoreFormat`, `DeltaUniformHudi`, `DeltaUniformIceberg`, `Hive`, `Iceberg`, `MongodbFormat`, `OracleFormat`, `SalesforceDataCloudFormat` and `TeradataFormat` enum values for [catalog.DataSourceFormat](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DataSourceFormat).
* Added `MetricView` enum value for [catalog.TableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableType).
* Added `SecurityAgentsFailedInitialVerification` enum value for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Added `CanCreateRegisteredModel` enum value for [ml.PermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#PermissionLevel).
* Added `Bigquery` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Added `AppendOnly` enum value for [pipelines.TableSpecificConfigScdType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfigScdType).
* Added `ServicePrincipal` and `UserGroup` enum values for [serving.AiGatewayRateLimitKey](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayRateLimitKey).
* [Breaking] Changed `DeleteTransitionRequest`, `UpdateModel`, `UpdateModelVersion` and `UpdateWebhook` methods for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service return type to become non-empty.
* [Breaking] Changed `DeleteWebhook` method for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service with new required argument order.
* [Breaking] Changed `List` method for [w.AlertsLegacy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsLegacyAPI) workspace-level service . New request type is `any`.
* [Breaking] Changed `Update` method for [w.DashboardWidgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardWidgetsAPI) workspace-level service . New request type is [sql.UpdateWidgetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateWidgetRequest).
* [Breaking] Changed `List` method for [w.DataSources](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DataSourcesAPI) workspace-level service . New request type is `any`.
* [Breaking] Changed `Create` method for [w.QueryVisualizationsLegacy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryVisualizationsLegacyAPI) workspace-level service with new required argument order.
* [Breaking] Changed `FromStage` and `ToStage` fields for [ml.Activity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Activity) to type `string`.
* [Breaking] Changed `Stage` field for [ml.ApproveTransitionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ApproveTransitionRequest) to type `string`.
* [Breaking] Changed `Stage` field for [ml.CreateTransitionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CreateTransitionRequest) to type `string`.
* [Breaking] Changed `Stage` field for [ml.DeleteTransitionRequestRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeleteTransitionRequestRequest) to type `string`.
* [Breaking] Changed `Id` field for [ml.DeleteWebhookRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeleteWebhookRequest) to be required.
* Changed `Capacity` field for [ml.OnlineStore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#OnlineStore) to be required.
* [Breaking] Changed `Capacity` field for [ml.OnlineStore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#OnlineStore) to be required.
* [Breaking] Changed `OnlineTableName` field for [ml.PublishSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#PublishSpec) to be required.
* [Breaking] Changed `Stage` field for [ml.RejectTransitionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RejectTransitionRequest) to type `string`.
* [Breaking] Changed `Stage` field for [ml.TransitionModelVersionStageDatabricks](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TransitionModelVersionStageDatabricks) to type `string`.
* [Breaking] Changed `ToStage` field for [ml.TransitionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TransitionRequest) to type `string`.
* Changed `ServedModelName` field for [serving.Route](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#Route) to no longer be required.
* [Breaking] Changed `ServedModelName` field for [serving.Route](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#Route) to no longer be required.
* Changed pagination for [TablesAPI.List](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TablesAPI.List).
* Changed pagination for [TablesAPI.ListSummaries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TablesAPI.ListSummaries).
* [Breaking] Removed `GenerateDownloadFullQueryResult` and `GetDownloadFullQueryResult` methods for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* [Breaking] Removed `IncludeDeltaMetadata` field for [catalog.ListTablesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListTablesRequest).
* [Breaking] Removed `Webhook` field for [ml.TestRegistryWebhookResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TestRegistryWebhookResponse).
* [Breaking] Removed `ModelVersion` field for [ml.TransitionStageResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TransitionStageResponse).
* [Breaking] Removed `UnknownCatalogType` enum value for [catalog.CatalogType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogType).
* [Breaking] Removed `HiveCustom` and `HiveSerde` enum values for [catalog.DataSourceFormat](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DataSourceFormat).
* [Breaking] Removed `UnknownSecurableType` enum value for [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).
* [Breaking] Removed `Archived`, `None`, `Production` and `Staging` enum values for [ml.DeleteTransitionRequestStage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeleteTransitionRequestStage).
* [Breaking] Removed `Archived`, `None`, `Production` and `Staging` enum values for [ml.Stage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Stage).
