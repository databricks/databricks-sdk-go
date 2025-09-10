# NEXT CHANGELOG

## Release v0.84.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added `JavaDependencies` field for [compute.Environment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Environment).
* Added `EffectiveCapacity` field for [database.DatabaseInstance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstance).
* [Breaking] Changed `CreationTime` field for [agentbricks.CustomLlm](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/agentbricks#CustomLlm) to type `string`.
* [Breaking] Changed `UpdateMask` field for [agentbricks.UpdateCustomLlmRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/agentbricks#UpdateCustomLlmRequest) to type `string`.
* [Breaking] Changed `CreateTime` and `UpdateTime` fields for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App) to type `string`.
* [Breaking] Changed `CreateTime` and `UpdateTime` fields for [apps.AppDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppDeployment) to type `string`.
* [Breaking] Changed `Timestamp` field for [catalog.ContinuousUpdateStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ContinuousUpdateStatus) to type `string`.
* [Breaking] Changed `EventTime` field for [catalog.ExternalLineageExternalMetadataInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLineageExternalMetadataInfo) to type `string`.
* [Breaking] Changed `EventTime` field for [catalog.ExternalLineageFileInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLineageFileInfo) to type `string`.
* [Breaking] Changed `EventTime` field for [catalog.ExternalLineageModelVersionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLineageModelVersionInfo) to type `string`.
* [Breaking] Changed `EventTime` field for [catalog.ExternalLineageTableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLineageTableInfo) to type `string`.
* [Breaking] Changed `CreateTime` and `UpdateTime` fields for [catalog.ExternalMetadata](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalMetadata) to type `string`.
* [Breaking] Changed `Timestamp` field for [catalog.FailedStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FailedStatus) to type `string`.
* [Breaking] Changed `Timestamp` field for [catalog.TriggeredUpdateStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TriggeredUpdateStatus) to type `string`.
* [Breaking] Changed `UpdateMask` field for [catalog.UpdateAccessRequestDestinationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateAccessRequestDestinationsRequest) to type `string`.
* [Breaking] Changed `UpdateMask` field for [catalog.UpdateEntityTagAssignmentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateEntityTagAssignmentRequest) to type `string`.
* [Breaking] Changed `UpdateMask` field for [catalog.UpdateExternalLineageRelationshipRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLineageRelationshipRequest) to type `string`.
* [Breaking] Changed `UpdateMask` field for [catalog.UpdateExternalMetadataRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalMetadataRequest) to type `string`.
* [Breaking] Changed `UpdateMask` field for [catalog.UpdatePolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePolicyRequest) to type `string`.
* [Breaking] Changed `UpdateMask` field for [compute.UpdateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#UpdateCluster) to type `string`.
* [Breaking] Changed `CreateTime` and `UpdateTime` fields for [dashboards.Dashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Dashboard) to type `string`.
* [Breaking] Changed `RevisionCreateTime` field for [dashboards.PublishedDashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#PublishedDashboard) to type `string`.
* [Breaking] Changed `CreateTime` and `UpdateTime` fields for [dashboards.Schedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Schedule) to type `string`.
* [Breaking] Changed `CreateTime` and `UpdateTime` fields for [dashboards.Subscription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Subscription) to type `string`.
* [Breaking] Changed `ExpirationTime` field for [database.DatabaseCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseCredential) to type `string`.
* [Breaking] Changed `CreationTime` field for [database.DatabaseInstance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstance) to type `string`.
* [Breaking] Changed `BranchTime` field for [database.DatabaseInstanceRef](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstanceRef) to type `string`.
* [Breaking] Changed `DeltaCommitTimestamp` field for [database.DeltaTableSyncInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DeltaTableSyncInfo) to type `string`.
* [Breaking] Changed `Timestamp` field for [database.SyncedTableContinuousUpdateStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableContinuousUpdateStatus) to type `string`.
* [Breaking] Changed `Timestamp` field for [database.SyncedTableFailedStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableFailedStatus) to type `string`.
* [Breaking] Changed `SyncEndTimestamp` and `SyncStartTimestamp` fields for [database.SyncedTablePosition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTablePosition) to type `string`.
* [Breaking] Changed `Timestamp` field for [database.SyncedTableTriggeredUpdateStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableTriggeredUpdateStatus) to type `string`.
* [Breaking] Changed `UpdateMask` field for [database.UpdateDatabaseCatalogRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#UpdateDatabaseCatalogRequest) to type `string`.
* [Breaking] Changed `UpdateMask` field for [database.UpdateDatabaseInstanceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#UpdateDatabaseInstanceRequest) to type `string`.
* [Breaking] Changed `UpdateMask` field for [database.UpdateSyncedDatabaseTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#UpdateSyncedDatabaseTableRequest) to type `string`.
* [Breaking] Changed `CreationTime` field for [ml.OnlineStore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#OnlineStore) to type `string`.
* [Breaking] Changed `UpdateMask` field for [ml.UpdateFeatureTagRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#UpdateFeatureTagRequest) to type `string`.
* [Breaking] Changed `UpdateMask` field for [ml.UpdateOnlineStoreRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#UpdateOnlineStoreRequest) to type `string`.
* [Breaking] Changed `Lifetime` field for [oauth2.CreateServicePrincipalSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateServicePrincipalSecretRequest) to type `string`.
* [Breaking] Changed `ExpireTime` field for [oauth2.CreateServicePrincipalSecretResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateServicePrincipalSecretResponse) to type `string`.
* [Breaking] Changed `CreateTime` and `UpdateTime` fields for [oauth2.FederationPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#FederationPolicy) to type `string`.
* [Breaking] Changed `ExpireTime` field for [oauth2.SecretInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#SecretInfo) to type `string`.
* [Breaking] Changed `UpdateMask` field for [oauth2.UpdateAccountFederationPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#UpdateAccountFederationPolicyRequest) to type `string`.
* [Breaking] Changed `UpdateMask` field for [oauth2.UpdateServicePrincipalFederationPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#UpdateServicePrincipalFederationPolicyRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateAccountIpAccessEnableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateAccountIpAccessEnableRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateAutomaticClusterUpdateSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateAutomaticClusterUpdateSettingRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateComplianceSecurityProfileSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateComplianceSecurityProfileSettingRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateCspEnablementAccountSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateCspEnablementAccountSettingRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateDashboardEmailSubscriptionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateDashboardEmailSubscriptionsRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateDefaultNamespaceSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateDefaultNamespaceSettingRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateDefaultWarehouseIdRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateDefaultWarehouseIdRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateDisableLegacyAccessRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateDisableLegacyAccessRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateDisableLegacyDbfsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateDisableLegacyDbfsRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateDisableLegacyFeaturesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateDisableLegacyFeaturesRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateEnableExportNotebookRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateEnableExportNotebookRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateEnableNotebookTableClipboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateEnableNotebookTableClipboardRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateEnableResultsDownloadingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateEnableResultsDownloadingRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateEnhancedSecurityMonitoringSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateEnhancedSecurityMonitoringSettingRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateEsmEnablementAccountSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateEsmEnablementAccountSettingRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateLlmProxyPartnerPoweredAccountRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateLlmProxyPartnerPoweredAccountRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateLlmProxyPartnerPoweredEnforceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateLlmProxyPartnerPoweredEnforceRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateLlmProxyPartnerPoweredWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateLlmProxyPartnerPoweredWorkspaceRequest) to type `string`.
* [Breaking] Changed `UpdateMask` field for [settings.UpdateNccPrivateEndpointRuleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateNccPrivateEndpointRuleRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdatePersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdatePersonalComputeSettingRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateRestrictWorkspaceAdminsSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateRestrictWorkspaceAdminsSettingRequest) to type `string`.
* [Breaking] Changed `FieldMask` field for [settings.UpdateSqlResultsDownloadRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateSqlResultsDownloadRequest) to type `string`.
* [Breaking] Changed `CreateTime` and `UpdateTime` fields for [sharing.FederationPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#FederationPolicy) to type `string`.
* [Breaking] Changed `UpdateMask` field for [sharing.UpdateFederationPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateFederationPolicyRequest) to type `string`.
* [Breaking] Changed `CreateTime`, `TriggerTime` and `UpdateTime` fields for [sql.Alert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Alert) to type `string`.
* [Breaking] Changed `CreateTime` and `UpdateTime` fields for [sql.AlertV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2) to type `string`.
* [Breaking] Changed `LastEvaluatedAt` field for [sql.AlertV2Evaluation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2Evaluation) to type `string`.
* [Breaking] Changed `CreateTime`, `TriggerTime` and `UpdateTime` fields for [sql.ListAlertsResponseAlert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListAlertsResponseAlert) to type `string`.
* [Breaking] Changed `CreateTime` and `UpdateTime` fields for [sql.ListQueryObjectsResponseQuery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListQueryObjectsResponseQuery) to type `string`.
* [Breaking] Changed `CreateTime` and `UpdateTime` fields for [sql.Query](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Query) to type `string`.
* [Breaking] Changed `UpdateMask` field for [sql.UpdateAlertRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateAlertRequest) to type `string`.
* [Breaking] Changed `UpdateMask` field for [sql.UpdateAlertV2Request](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateAlertV2Request) to type `string`.
* [Breaking] Changed `UpdateMask` field for [sql.UpdateQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateQueryRequest) to type `string`.
* [Breaking] Changed `UpdateMask` field for [sql.UpdateVisualizationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateVisualizationRequest) to type `string`.
* [Breaking] Changed `CreateTime` and `UpdateTime` fields for [sql.Visualization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Visualization) to type `string`.
* [Breaking] Changed `UpdateMask` field for [tags.UpdateTagPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/tags#UpdateTagPolicyRequest) to type `string`.
* [Breaking] Removed `DefaultDataSecurityMode` and `EffectiveDefaultDataSecurityMode` fields for [settingsv2.Setting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#Setting).