# NEXT CHANGELOG

## Release v0.72.0

### New Features and Improvements

- Enums now have a `Values` function that returns the list of all possible
  enum values ([PR #1228](https://github.com/databricks/databricks-sdk-go/pull/1228))

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added [aibuilder](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aibuilder) package, [database](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database) package and [qualitymonitorv2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/qualitymonitorv2) package.
* Added [w.CustomLlms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aibuilder#CustomLlmsAPI) workspace-level service.
* Added [w.DashboardEmailSubscriptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DashboardEmailSubscriptionsAPI) workspace-level service and [w.SqlResultsDownload](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SqlResultsDownloadAPI) workspace-level service.
* Added [w.Database](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseAPI) workspace-level service.
* Added [w.QualityMonitorV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/qualitymonitorv2#QualityMonitorV2API) workspace-level service.
* Added `UpdatePrivateEndpointRule` method for [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service.
* Added `ListSpaces` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Added `PageToken` field for [billing.ListLogDeliveryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#ListLogDeliveryRequest).
* Added `NextPageToken` field for [billing.WrappedLogDeliveryConfigurations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#WrappedLogDeliveryConfigurations).
* Added `NextPageToken` field for [catalog.EffectivePermissionsList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EffectivePermissionsList).
* Added `MaxResults` and `PageToken` fields for [catalog.GetEffectiveRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetEffectiveRequest).
* Added `MaxResults` and `PageToken` fields for [catalog.GetGrantRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetGrantRequest).
* Added `NextPageToken` field for [catalog.ListMetastoresResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListMetastoresResponse).
* Added `CleanRoomName` field for [cleanrooms.CleanRoomAsset](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAsset).
* [Breaking] Added `Name` field for [cleanrooms.DeleteCleanRoomAssetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#DeleteCleanRoomAssetRequest).
* [Breaking] Added `Name` field for [cleanrooms.GetCleanRoomAssetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#GetCleanRoomAssetRequest).
* Added `TriggerState` field for [jobs.BaseJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseJob).
* Added `TriggerState` field for [jobs.Job](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Job).
* Added `DbtCloudOutput` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
* Added `DbtCloudTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Added `DbtCloudTask` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Added `DbtCloudTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Added `Tags` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Added `Tags` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Added `Tags` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
* Added `MaxProvisionedConcurrency` and `MinProvisionedConcurrency` fields for [serving.ServedEntityInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityInput).
* Added `MaxProvisionedConcurrency` and `MinProvisionedConcurrency` fields for [serving.ServedEntityOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityOutput).
* Added `MaxProvisionedConcurrency` and `MinProvisionedConcurrency` fields for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput).
* Added `MaxProvisionedConcurrency` and `MinProvisionedConcurrency` fields for [serving.ServedModelOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelOutput).
* Added `EndpointService` and `ResourceNames` fields for [settings.CreatePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreatePrivateEndpointRule).
* Added `AwsPrivateEndpointRules` field for [settings.NccEgressTargetRules](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccEgressTargetRules).
* Added `TaskTimeOverTimeRange` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
* Added `DeltasharingCatalog`, `ForeignCatalog`, `InternalCatalog`, `ManagedCatalog`, `ManagedOnlineCatalog`, `SystemCatalog` and `UnknownCatalogType` enum values for [catalog.CatalogType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogType).
* Added `Ga4RawData`, `PowerBi`, `Salesforce`, `SalesforceDataCloud`, `Servicenow`, `UnknownConnectionType` and `WorkdayRaas` enum values for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Added `OauthAccessToken`, `OauthM2m`, `OauthRefreshToken`, `OauthResourceOwnerPassword`, `OauthU2m`, `OauthU2mMapping`, `OidcToken`, `PemPrivateKey`, `ServiceCredential` and `UnknownCredentialType` enum values for [catalog.CredentialType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CredentialType).
* Added `Internal` and `InternalAndExternal` enum values for [catalog.DeltaSharingScopeEnum](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeltaSharingScopeEnum).
* Added `Catalog`, `CleanRoom`, `Connection`, `Credential`, `ExternalLocation`, `ExternalMetadata`, `Function`, `Metastore`, `Pipeline`, `Provider`, `Recipient`, `Schema`, `Share`, `StagingTable`, `StorageCredential`, `Table`, `UnknownSecurableType` and `Volume` enum values for [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).
* Added `ClusterMigrated` enum value for [compute.EventType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventType).
* Added `DriverUnhealthy` enum value for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Added `Teradata` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Added `OidcFederation` enum value for [sharing.AuthenticationType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#AuthenticationType).
* [Breaking] Changed `Create` method for [a.LogDelivery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#LogDeliveryAPI) account-level service with new required argument order.
* [Breaking] Changed `Get` method for [a.LogDelivery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#LogDeliveryAPI) account-level service to return [billing.GetLogDeliveryConfigurationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#GetLogDeliveryConfigurationResponse).
* [Breaking] Changed `CreatePrivateEndpointRule`, `DeletePrivateEndpointRule` and `GetPrivateEndpointRule` methods for [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service to return [settings.NccPrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccPrivateEndpointRule).
* [Breaking] Changed `ListPrivateEndpointRules` method for [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service to return [settings.ListPrivateEndpointRulesResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListPrivateEndpointRulesResponse).
* [Breaking] Changed `Delete` and `Get` methods for [w.CleanRoomAssets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetsAPI) workspace-level service with new required argument order.
* [Breaking] Changed `Delete` and `Get` methods for [w.CleanRoomAssets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetsAPI) workspace-level service . Method path has changed.
* [Breaking] Changed `Get` method for [w.Grants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GrantsAPI) workspace-level service to return [catalog.GetPermissionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetPermissionsResponse).
* [Breaking] Changed `Update` method for [w.Grants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GrantsAPI) workspace-level service to return [catalog.UpdatePermissionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePermissionsResponse).
* [Breaking] Changed `List` method for [w.Metastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoresAPI) workspace-level service to require request of [catalog.ListMetastoresRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListMetastoresRequest).
* Changed `AccountId`, `CredentialsId`, `LogType`, `OutputFormat` and `StorageConfigurationId` fields for [billing.LogDeliveryConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#LogDeliveryConfiguration) to be required.
* Changed `Message` and `Status` fields for [billing.LogDeliveryStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#LogDeliveryStatus) to be required.
* [Breaking] Changed `LogDeliveryConfiguration` field for [billing.WrappedCreateLogDeliveryConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#WrappedCreateLogDeliveryConfiguration) to be required.
* [Breaking] Changed `SecurableType` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo) to type [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).
* [Breaking] Changed `SecurableType` field for [catalog.GetEffectiveRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetEffectiveRequest) to type `string`.
* [Breaking] Changed `SecurableType` field for [catalog.GetGrantRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetGrantRequest) to type `string`.
* [Breaking] Changed `DeltaSharingScope` field for [catalog.GetMetastoreSummaryResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetMetastoreSummaryResponse) to type [catalog.DeltaSharingScopeEnum](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeltaSharingScopeEnum).
* [Breaking] Changed `DeltaSharingScope` field for [catalog.MetastoreInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoreInfo) to type [catalog.DeltaSharingScopeEnum](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeltaSharingScopeEnum).
* [Breaking] Changed `CatalogType` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo) to type [catalog.CatalogType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogType).
* [Breaking] Changed `DeltaSharingScope` field for [catalog.UpdateMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMetastore) to type [catalog.DeltaSharingScopeEnum](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeltaSharingScopeEnum).
* [Breaking] Changed `SecurableType` field for [catalog.UpdatePermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePermissions) to type `string`.
* Changed `ResourceId` field for [settings.CreatePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreatePrivateEndpointRule) to no longer be required.
* [Breaking] Changed pagination for [NetworkConnectivityAPI.ListPrivateEndpointRules](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI.ListPrivateEndpointRules).
* [Breaking] Removed [w.DatabaseInstances](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DatabaseInstancesAPI) workspace-level service.
* [Breaking] Removed [w.QueryExecution](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#QueryExecutionAPI) workspace-level service.
* [Breaking] Removed `UpdateNccAzurePrivateEndpointRulePublic` method for [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service.
* [Breaking] Removed `GetCredentialsForTraceDataDownload`, `GetCredentialsForTraceDataUpload` and `ListLoggedModelArtifacts` methods for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
* [Breaking] Removed `GetPublishedDashboardEmbedded` method for [w.LakeviewEmbedded](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewEmbeddedAPI) workspace-level service.
* [Breaking] Removed `AssetFullName` field for [cleanrooms.DeleteCleanRoomAssetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#DeleteCleanRoomAssetRequest).
* [Breaking] Removed `AssetFullName` field for [cleanrooms.GetCleanRoomAssetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#GetCleanRoomAssetRequest).
* [Breaking] Removed `Internal` and `InternalAndExternal` enum values for [catalog.GetMetastoreSummaryResponseDeltaSharingScope](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetMetastoreSummaryResponseDeltaSharingScope).
* [Breaking] Removed `Internal` and `InternalAndExternal` enum values for [catalog.MetastoreInfoDeltaSharingScope](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoreInfoDeltaSharingScope).
* [Breaking] Removed `Catalog`, `CleanRoom`, `Connection`, `Credential`, `ExternalLocation`, `ExternalMetadata`, `Function`, `Metastore`, `Pipeline`, `Provider`, `Recipient`, `Schema`, `Share`, `StagingTable`, `StorageCredential`, `Table`, `UnknownSecurableType` and `Volume` enum values for [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).
* [Breaking] Removed `Internal` and `InternalAndExternal` enum values for [catalog.UpdateMetastoreDeltaSharingScope](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMetastoreDeltaSharingScope).
