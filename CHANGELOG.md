# Version changelog

## 0.22.0

Breaking API Changes:

 * Changed `List` method for [a.AccountMetastoreAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoreAssignmentsAPI) account-level service to return [catalog.ListAccountMetastoreAssignmentsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListAccountMetastoreAssignmentsResponse).

API Changes:

 * Added [catalog.ListAccountMetastoreAssignmentsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListAccountMetastoreAssignmentsResponse).

Internal SDK Changes:

* Add support for template libraries ([#635](https://github.com/databricks/databricks-sdk-go/pull/635)).

OpenAPI SHA: bcbf6e851e3d82fd910940910dd31c10c059746c, Date: 2023-10-02

## 0.21.0

Breaking API Changes:

 * Changed `ArtifactMatchers` field for [catalog.ArtifactAllowlistInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactAllowlistInfo) to [catalog.ArtifactMatcherList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactMatcherList).
 * Removed `Owner` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
 * Changed `ArtifactMatchers` field for [catalog.SetArtifactAllowlist](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SetArtifactAllowlist) to [catalog.ArtifactMatcherList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactMatcherList).
 * Removed `Options` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog).
 * Changed `CancelAllRuns` method for [w.Jobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsAPI) workspace-level service with new required argument order.
 * Changed `JobId` field for [jobs.CancelAllRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CancelAllRuns) to no longer be required.
 * Changed `JobParameters` field for [jobs.RunNow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunNow) to [jobs.ParamPairs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ParamPairs).
 * Changed `Query` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service . New request type is [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Removed [serving.QueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryRequest).
 * Changed `ExchangeToken` method for [w.CredentialsManager](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CredentialsManagerAPI) workspace-level service with new required argument order.
 * Removed [settings.CredentialPartitionId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CredentialPartitionId).
 * Changed `TokenType` field for [settings.ExchangeToken](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeToken) to [settings.TokenType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenType).
 * Removed `CredentialPartitionId` field for [settings.ExchangeTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeTokenRequest).
 * Changed `List` method for [w.CleanRooms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomsAPI) workspace-level service to require request of [sharing.ListCleanRoomsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListCleanRoomsRequest).

API Changes:

 * Added `AllQueuedRuns` field for [jobs.CancelAllRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CancelAllRuns).
 * Added `Queue` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
 * Added `Queue` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Added `JobParameters` field for [jobs.RepairRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RepairRun).
 * Added `Queue` field for [jobs.RunNow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunNow).
 * Added `JobParameters` field for [jobs.RunParameters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunParameters).
 * Added `QueueReason` field for [jobs.RunState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunState).
 * Added `QueueDuration` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
 * Added `Queue` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
 * Added [jobs.QueueSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#QueueSettings).
 * Added `Notifications` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
 * Added `Notifications` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
 * Added `Notifications` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
 * Added [pipelines.Notifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#Notifications).
 * Added [serving.DataframeSplitInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DataframeSplitInput).
 * Added [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added [w.Settings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SettingsAPI) workspace-level service.
 * Added `PartitionId` field for [settings.ExchangeTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeTokenRequest).
 * Added [settings.DefaultNamespaceSetting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DefaultNamespaceSetting).
 * Added [settings.DeleteDefaultWorkspaceNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultWorkspaceNamespaceRequest).
 * Added [settings.DeleteDefaultWorkspaceNamespaceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultWorkspaceNamespaceResponse).
 * Added [settings.PartitionId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#PartitionId).
 * Added [settings.ReadDefaultWorkspaceNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadDefaultWorkspaceNamespaceRequest).
 * Added [settings.StringMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#StringMessage).
 * Added [settings.UpdateDefaultWorkspaceNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateDefaultWorkspaceNamespaceRequest).
 * Added `NextPageToken` field for [sharing.ListCleanRoomsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListCleanRoomsResponse).
 * Added [sharing.ListCleanRoomsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListCleanRoomsRequest).

SDK Internal Changes:

* Remove use of cancel by job id in integration test ([#632](https://github.com/databricks/databricks-sdk-go/pull/632)).


OpenAPI SHA: 94ddf8ff02df271daebdc5f001075e1ca2ce080d, Date: 2023-09-27

## 0.20.0

* Adjust token expiry window to 40 seconds on Azure ([#617](https://github.com/databricks/databricks-sdk-go/pull/617)).
* Rename variables in Azure code to match across auth types ([#618](https://github.com/databricks/databricks-sdk-go/pull/618)).
* Add details to APIError ([#613](https://github.com/databricks/databricks-sdk-go/pull/613)).
* Fixed code generation of primitive types ([#623](https://github.com/databricks/databricks-sdk-go/pull/623)).
* Updated SDK to changes in OpenAPI specification ([#625](https://github.com/databricks/databricks-sdk-go/pull/625)).

API Changes:

 * Changed `List` method for [a.AccountMetastoreAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoreAssignmentsAPI) account-level service to return [catalog.WorkspaceIdList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceIdList).
 * Added [catalog.WorkspaceId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceId).
 * Added [a.OAuthPublishedApps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#OAuthPublishedAppsAPI) account-level service.
 * Added [oauth2.GetPublishedAppsOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#GetPublishedAppsOutput).
 * Added [oauth2.ListOAuthPublishedAppsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ListOAuthPublishedAppsRequest).
 * Added [oauth2.PublishedAppOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#PublishedAppOutput).
 * Added `Patch` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `Tags` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint).
 * Added `Tags` field for [serving.ServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpoint).
 * Added `Tags` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).
 * Added [serving.EndpointTag](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointTag).
 * Added [serving.PatchServingEndpointTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PatchServingEndpointTags).
 * Added [w.CredentialsManager](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CredentialsManagerAPI) workspace-level service.
 * Added [settings.CredentialPartitionId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CredentialPartitionId).
 * Added [settings.ExchangeToken](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeToken).
 * Added [settings.ExchangeTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeTokenRequest).
 * Added [settings.ExchangeTokenResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeTokenResponse).
 * Added [settings.TokenType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenType).
 * Changed `ExecuteStatement` method for [w.StatementExecution](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#StatementExecutionAPI) workspace-level service with new required argument order.
 * Added `EmptyResultState` field for [sql.AlertOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertOptions).
 * Removed [sql.ChunkInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ChunkInfo).
 * Changed `OnWaitTimeout` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest) to [sql.ExecuteStatementRequestOnWaitTimeout](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequestOnWaitTimeout).
 * Changed `Statement` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest) to be required.
 * Changed `WarehouseId` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest) to be required.
 * Changed `Chunks` field for [sql.ResultManifest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ResultManifest) to [sql.BaseChunkInfoList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#BaseChunkInfoList).
 * Added `Truncated` field for [sql.ResultManifest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ResultManifest).
 * Removed [sql.TimeoutAction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TimeoutAction).
 * Added [sql.AlertOptionsEmptyResultState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertOptionsEmptyResultState).
 * Added [sql.BaseChunkInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#BaseChunkInfo).
 * Added [sql.ExecuteStatementRequestOnWaitTimeout](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequestOnWaitTimeout).

OpenAPI SHA: 51e3d27c0553c891bc418bd0cd07c9240e0476dd, Date: 2023-09-13
Dependency updates:

 * Bump golang.org/x/oauth2 from 0.11.0 to 0.12.0 ([#606](https://github.com/databricks/databricks-sdk-go/pull/606)).
 * Bump google.golang.org/api from 0.138.0 to 0.140.0 ([#612](https://github.com/databricks/databricks-sdk-go/pull/612)).

## 0.19.2

* Fixed case where retry of a request without body fails ([#614](https://github.com/databricks/databricks-sdk-go/pull/614)).


## 0.19.1

* Handled Azure authentication when WorkspaceResourceID is provided ([#597](https://github.com/databricks/databricks-sdk-go/pull/597)).
* Reverted error message changes from client.go ([#604](https://github.com/databricks/databricks-sdk-go/pull/604)).
* Reused tokens after first call to Azure CLI ([#605](https://github.com/databricks/databricks-sdk-go/pull/605)).
* Updated error message assertions ([#607](https://github.com/databricks/databricks-sdk-go/pull/607)).


## 0.19.0

* Added `ErrNotWorkspaceClient` ([#596](https://github.com/databricks/databricks-sdk-go/pull/596)).
* Fix loading of databrickscfg with a password containing a hash ([#595](https://github.com/databricks/databricks-sdk-go/pull/595)).
* Use an operation's request type name if specified ([#598](https://github.com/databricks/databricks-sdk-go/pull/598)).
* Update OpenAPI spec ([#600](https://github.com/databricks/databricks-sdk-go/pull/600)).

API Changes:

* Renamed permissions APIs to no longer include the service name, for example:
  * `GetJobPermissionLevels` -> `GetPermissionLevels`
  * `GetJobPermissions` -> `GetPermissions`
  * `SetJobPermissions` -> `SetPermissions`
  * `UpdateJobPermissions` -> `UpdatePermissions`
* Changed `Create` method for [w.Volumes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumesAPI) workspace-level service with new required argument order.
* Added `SupportsElasticDisk` field for [compute.NodeType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#NodeType).
* Changed `Create` method for [w.Dashboards](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardsAPI) workspace-level service with new required argument order.
* Added [w.DashboardWidgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardWidgetsAPI) workspace-level service.
* Added [w.QueryVisualizations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryVisualizationsAPI) workspace-level service.
* Changed `Name` field for [sql.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateDashboardRequest) to be required.
* Added `DashboardFiltersEnabled` field for [sql.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateDashboardRequest).
* Added `RunAsRole` field for [sql.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateDashboardRequest).
* Added `RunAsRole` field for [sql.Query](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Query).
* Added `RunAsRole` field for [sql.QueryPostContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryPostContent).
* Removed `DashboardId` field for [sql.WidgetOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetOptions).
* Changed `Position` field for [sql.WidgetOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetOptions) to [sql.WidgetPosition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetPosition).
* Removed `Text` field for [sql.WidgetOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetOptions).
* Added `Description` field for [sql.WidgetOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetOptions).
* Added `Title` field for [sql.WidgetOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetOptions).
* Added [sql.CreateQueryVisualizationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateQueryVisualizationRequest).
* Added [sql.CreateWidget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateWidget).
* Added [sql.DeleteDashboardWidgetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DeleteDashboardWidgetRequest).
* Added [sql.DeleteQueryVisualizationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DeleteQueryVisualizationRequest).
* Added [sql.RunAsRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#RunAsRole).
* Added [sql.WidgetPosition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetPosition).

OpenAPI SHA: 09a7fa63d9ae243e5407941f200960ca14d48b07, Date: 2023-09-04

Dependency updates:

* Bump google.golang.org/api from 0.136.0 to 0.138.0 ([#587](https://github.com/databricks/databricks-sdk-go/pull/587)).
## 0.18.0

Breaking Changes:
* Added support Files API (application/octet-stream) in OpenAPI Spec ([#572](https://github.com/databricks/databricks-sdk-go/pull/572)). The signatures of `Upload`, `Download` and `Delete` have changed; these methods now take `UploadRequest`, `DownloadRequest`, and `DeleteRequest` structures, respectively. Shortcut methods are generated for `DownloadByFileName` and `DeleteByFileName` for convenience. The `WriteFile` and `ReadFile` methods are removed.
* Propagated Request Headers to client.go ([#589](https://github.com/databricks/databricks-sdk-go/pull/589)). The `Do` method of `Client` now requires an additional parameter for request headers. Add headers to requests using this parameter, or pass `nil` if no headers are needed.

Breaking API Changes:
 * Removed [w.SecurableTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableTagsAPI) workspace-level service and all associated structures.
 * Removed [w.SubentityTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SubentityTagsAPI) workspace-level service and all associated structures.
 * Renamed `ProvisioningState` field to `ProvisioningInfo` for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Changed [catalog.ProvisioningState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningState) to [catalog.ProvisioningInfoState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningInfoState).
 * Removed `InstancePoolFleetAttributes` field for [compute.CreateInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateInstancePool).
 * Removed `InstancePoolFleetAttributes` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Removed [compute.FleetLaunchTemplateOverride](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#FleetLaunchTemplateOverride).
 * Removed [compute.FleetOnDemandOption](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#FleetOnDemandOption).
 * Removed [compute.FleetOnDemandOptionAllocationStrategy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#FleetOnDemandOptionAllocationStrategy).
 * Removed [compute.FleetSpotOption](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#FleetSpotOption).
 * Removed [compute.FleetSpotOptionAllocationStrategy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#FleetSpotOptionAllocationStrategy).
 * Removed `InstancePoolFleetAttributes` field for [compute.GetInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePool).
 * Removed `InstancePoolFleetAttributes` field for [compute.InstancePoolAndStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAndStats).
 * Removed [compute.InstancePoolFleetAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolFleetAttributes).
 * Changed `GetByName` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service to return [ml.GetExperimentResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentResponse).
 * Changed `GetExperiment` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service to return [ml.GetExperimentResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentResponse).
 * Renamed [ml.GetExperimentByNameResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentByNameResponse) to [ml.GetExperimentResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentResponse).

API Changes:
 * Changed `List` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to return [catalog.StorageCredentialInfoList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialInfoList).
 * Added [w.ModelVersions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionsAPI) workspace-level service.
 * Added [w.RegisteredModels](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelsAPI) workspace-level service.
 * Added `BrowseOnly` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Added `FullName` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Added `ProvisioningInfo` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Added `SecurableKind` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Added `SecurableType` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Added `Options` field for [catalog.CreateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateCatalog).
 * Added `Options` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog).
 * Added [catalog.CreateRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateRegisteredModelRequest).
 * Added [catalog.DeleteAliasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteAliasRequest).
 * Added [catalog.DeleteModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteModelVersionRequest).
 * Added [catalog.DeleteRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteRegisteredModelRequest).
 * Added [catalog.GetByAliasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetByAliasRequest).
 * Added [catalog.GetModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelVersionRequest).
 * Added [catalog.GetRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRegisteredModelRequest).
 * Added [catalog.ListModelVersionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelVersionsRequest).
 * Added [catalog.ListModelVersionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelVersionsResponse).
 * Added [catalog.ListRegisteredModelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRegisteredModelsRequest).
 * Added [catalog.ListRegisteredModelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRegisteredModelsResponse).
 * Added [catalog.ModelVersionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfo).
 * Added [catalog.ModelVersionInfoStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfoStatus).
 * Added [catalog.ProvisioningInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningInfo).
 * Added [catalog.ProvisioningInfoState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningInfoState).
 * Added [catalog.RegisteredModelAlias](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelAlias).
 * Added [catalog.RegisteredModelInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelInfo).
 * Added [catalog.SetRegisteredModelAliasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SetRegisteredModelAliasRequest).
 * Added [catalog.UpdateModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateModelVersionRequest).
 * Added [catalog.UpdateRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateRegisteredModelRequest).
 * Added `Volumes` field for [compute.InitScriptInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptInfo).
 * Added [compute.VolumesStorageInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#VolumesStorageInfo).
 * Added [w.Files](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#FilesAPI) workspace-level service.
 * Added [files.DeleteFileRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DeleteFileRequest).
 * Added [files.DownloadRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DownloadRequest).
 * Added [files.DownloadResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DownloadResponse).
 * Added [files.UploadRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#UploadRequest).
 * Added `CustomTags` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest).
 * Added `CustomTags` field for [provisioning.UpdateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#UpdateWorkspaceRequest).
 * Added `CustomTags` field for [provisioning.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Workspace).
 * Added [provisioning.CustomTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CustomTags).
 * Added `Parameters` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest).
 * Added `RowLimit` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest).
 * Added [sql.StatementParameterListItem](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#StatementParameterListItem).

SDK Internal Changes:
* Added support for io.ReadCloser for request and response bodies ([#590](https://github.com/databricks/databricks-sdk-go/pull/590)).
* Fixed nightly test failures for Go SDK ([#592](https://github.com/databricks/databricks-sdk-go/pull/592)).

OpenAPI Generator Changes:
* Improved error handling for ill-defined parameters ([#582](https://github.com/databricks/databricks-sdk-go/pull/582)).
* Switched to logging consistently rather than printing to stdout ([#588](https://github.com/databricks/databricks-sdk-go/pull/588)).

OpenAPI SHA: 5d0ccbb790d341eae8e85321a685a9e9e2d5bf24, Date: 2023-08-29

## 0.17.0

* Introduced Artifact Allowlist, Securable Tags, and Subentity Tags services.
* Introduced DeleteRuns and RestoreRuns methods in the Experiments API.
* Introduced the GetSecret method in the Secrets API.
* Renamed Auto Maintenance to Predictive Optimization.
* Set necessary headers when authenticating via Azure CLI ([#584](https://github.com/databricks/databricks-sdk-go/pull/584)).

New Services:

 * Added [w.ArtifactAllowlists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactAllowlistsAPI) workspace-level service.
 * Added [w.SecurableTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableTagsAPI) workspace-level service.
 * Added [w.SubentityTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SubentityTagsAPI) workspace-level service.
 * Added [catalog.ArtifactAllowlistInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactAllowlistInfo).
 * Added [catalog.ArtifactMatcher](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactMatcher).
 * Added [catalog.ArtifactType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactType).
 * Added [catalog.GetArtifactAllowlistRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetArtifactAllowlistRequest).
 * Added [catalog.ListSecurableTagsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSecurableTagsRequest).
 * Added [catalog.ListSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSecurableType).
 * Added [catalog.ListSubentityTagsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSubentityTagsRequest).
 * Added [catalog.MatchType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MatchType).
 * Added [catalog.SetArtifactAllowlist](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SetArtifactAllowlist).
 * Added [catalog.TagChanges](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagChanges).
 * Added [catalog.TagKeyValuePair](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagKeyValuePair).
 * Added [catalog.TagSecurable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagSecurable).
 * Added [catalog.TagSecurableAssignment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagSecurableAssignment).
 * Added [catalog.TagSecurableAssignmentsList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagSecurableAssignmentsList).
 * Added [catalog.TagSubentity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagSubentity).
 * Added [catalog.TagSubentityAssignmentsList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagSubentityAssignmentsList).
 * Added [catalog.TagsSubentityAssignment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagsSubentityAssignment).
 * Added [catalog.UpdateSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateSecurableType).
 * Added [catalog.UpdateTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateTags).

New APIs:

 * Added `DeleteRuns` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added `RestoreRuns` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added [ml.DeleteRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeleteRuns).
 * Added [ml.DeleteRunsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeleteRunsResponse).
 * Added [ml.RestoreRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RestoreRuns).
 * Added [ml.RestoreRunsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RestoreRunsResponse).
 * Added `GetSecret` method for [w.Secrets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#SecretsAPI) workspace-level service.
 * Added [workspace.GetSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetSecretRequest).
 * Added [workspace.GetSecretResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetSecretResponse).

Service Renames:

 * Renamed `EffectiveAutoMaintenanceFlag` field to `EffectivePredictiveOptimizationFlag` for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Renamed `EnableAutoMaintenance` field to `EnablePredictiveOptimization` for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Renamed [catalog.EffectiveAutoMaintenanceFlag](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EffectiveAutoMaintenanceFlag) to [catalog.EffectivePredictiveOptimizationFlag](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EffectivePredictiveOptimizationFlag).
 * Renamed [catalog.EffectiveAutoMaintenanceFlagInheritedFromType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EffectiveAutoMaintenanceFlagInheritedFromType) to [catalog.EffectivePredictiveOptimizationFlagInheritedFromType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EffectivePredictiveOptimizationFlagInheritedFromType).
 * Renamed [catalog.EnableAutoMaintenance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EnableAutoMaintenance) to [catalog.EnablePredictiveOptimization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EnablePredictiveOptimization).
 * Renamed `EffectiveAutoMaintenanceFlag` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo) to `EffectivePredictiveOptimizationFlag` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo).
 * Renamed `EnableAutoMaintenance` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo) to `EnablePredictiveOptimization` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo).
 * Renamed `EffectiveAutoMaintenanceFlag` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo) to `EffectivePredictiveOptimizationFlag` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).
 * Renamed `EnableAutoMaintenance` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo) to `EnablePredictiveOptimization` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).

OpenAPI SHA: beff621d7b3e1d59244e2e34fc53a496f310e130, Date: 2023-08-17

## 0.16.0

* Added ability to generate flat names ([#568](https://github.com/databricks/databricks-sdk-go/pull/568)).

API Changes:

 * Changed `Create` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to return [catalog.AccountsStorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsStorageCredentialInfo).
 * Changed `Get` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to return [catalog.AccountsStorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsStorageCredentialInfo).
 * Changed `Update` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to return [catalog.AccountsStorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsStorageCredentialInfo).
 * Changed `Create` method for [w.Connections](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionsAPI) workspace-level service with new required argument order.
 * Changed `Update` method for [w.Connections](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionsAPI) workspace-level service with new required argument order.
 * Removed `OptionsKvpairs` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Removed `PropertiesKvpairs` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Added `Options` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Added `Properties` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Added `ProvisioningState` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Added `SecurableKind` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Added `SecurableType` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Removed `OptionsKvpairs` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
 * Removed `PropertiesKvpairs` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
 * Added `Options` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
 * Added `Properties` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
 * Changed `Algorithm` field for [catalog.SseEncryptionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SseEncryptionDetails) to no longer be required.
 * Removed `OptionsKvpairs` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
 * Added `Options` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
 * Added [catalog.AccountsStorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsStorageCredentialInfo).
 * Added [catalog.ConnectionInfoSecurableKind](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfoSecurableKind).
 * Added [catalog.ProvisioningState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningState).
 * Added `DataSecurityMode` field for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
 * Added `DockerImage` field for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
 * Added `SingleUserName` field for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
 * Removed `Schema` field for [iam.PartialUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PartialUpdate).
 * Added `Schemas` field for [iam.PartialUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PartialUpdate).

OpenAPI SHA: 1e3533f94335f0e6c5d9262bc1fea95b3ddcb0e1, Date: 2023-08-11
Dependency updates:

 * Bump golang.org/x/oauth2 from 0.10.0 to 0.11.0 ([#574](https://github.com/databricks/databricks-sdk-go/pull/574)).
 * Bump google.golang.org/api from 0.134.0 to 0.135.0 ([#575](https://github.com/databricks/databricks-sdk-go/pull/575)).
 * Bump google.golang.org/api from 0.135.0 to 0.136.0 ([#576](https://github.com/databricks/databricks-sdk-go/pull/576)).

## 0.15.0

To simplify documentation and management of object permissions, this release features a major reorganization of how permissions APIs are structured in the SDK. Rather than using a single `Permissions.Get()` API for all services, each service supporting permissions has its own permissions APIs. Follow these steps to migrate to the current SDK:

 * Change `w.Permissions.Get()` and `w.Permissions.GetByRequestObjectIdAndRequestObjectType()` to `w.<Service>.Get<Service>Permissions()`
 * Change `w.Permissions.GetPermissionLevels()` to `w.<Service>.GetServicePermissionLevels()`
 * Change `w.Permissions.Set()` to `w.<Service>.Set<Service>Permissions()`
 * Change `w.Permissions.Update()` to `w.<Service>.Update<Service>Permissions()`

API Changes:

 * Added `GetClusterPolicyPermissionLevels` method for [w.ClusterPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPoliciesAPI) workspace-level service.
 * Added `GetClusterPolicyPermissions` method for [w.ClusterPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPoliciesAPI) workspace-level service.
 * Added `SetClusterPolicyPermissions` method for [w.ClusterPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPoliciesAPI) workspace-level service.
 * Added `UpdateClusterPolicyPermissions` method for [w.ClusterPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPoliciesAPI) workspace-level service.
 * Added `GetClusterPermissionLevels` method for [w.Clusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClustersAPI) workspace-level service.
 * Added `GetClusterPermissions` method for [w.Clusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClustersAPI) workspace-level service.
 * Added `SetClusterPermissions` method for [w.Clusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClustersAPI) workspace-level service.
 * Added `UpdateClusterPermissions` method for [w.Clusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClustersAPI) workspace-level service.
 * Added `GetInstancePoolPermissionLevels` method for [w.InstancePools](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolsAPI) workspace-level service.
 * Added `GetInstancePoolPermissions` method for [w.InstancePools](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolsAPI) workspace-level service.
 * Added `SetInstancePoolPermissions` method for [w.InstancePools](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolsAPI) workspace-level service.
 * Added `UpdateInstancePoolPermissions` method for [w.InstancePools](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolsAPI) workspace-level service.
 * Added [compute.ClusterAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterAccessControlRequest).
 * Added [compute.ClusterAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterAccessControlResponse).
 * Added [compute.ClusterPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPermission).
 * Added [compute.ClusterPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPermissionLevel).
 * Added [compute.ClusterPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPermissions).
 * Added [compute.ClusterPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPermissionsDescription).
 * Added [compute.ClusterPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPermissionsRequest).
 * Added [compute.ClusterPolicyAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyAccessControlRequest).
 * Added [compute.ClusterPolicyAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyAccessControlResponse).
 * Added [compute.ClusterPolicyPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyPermission).
 * Added [compute.ClusterPolicyPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyPermissionLevel).
 * Added [compute.ClusterPolicyPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyPermissions).
 * Added [compute.ClusterPolicyPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyPermissionsDescription).
 * Added [compute.ClusterPolicyPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyPermissionsRequest).
 * Added [compute.GetClusterPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPermissionLevelsRequest).
 * Added [compute.GetClusterPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPermissionLevelsResponse).
 * Added [compute.GetClusterPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPermissionsRequest).
 * Added [compute.GetClusterPolicyPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPolicyPermissionLevelsRequest).
 * Added [compute.GetClusterPolicyPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPolicyPermissionLevelsResponse).
 * Added [compute.GetClusterPolicyPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPolicyPermissionsRequest).
 * Added [compute.GetInstancePoolPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePoolPermissionLevelsRequest).
 * Added [compute.GetInstancePoolPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePoolPermissionLevelsResponse).
 * Added [compute.GetInstancePoolPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePoolPermissionsRequest).
 * Added [compute.InstancePoolAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAccessControlRequest).
 * Added [compute.InstancePoolAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAccessControlResponse).
 * Added [compute.InstancePoolPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolPermission).
 * Added [compute.InstancePoolPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolPermissionLevel).
 * Added [compute.InstancePoolPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolPermissions).
 * Added [compute.InstancePoolPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolPermissionsDescription).
 * Added [compute.InstancePoolPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolPermissionsRequest).
 * Changed `Set` method for [w.Permissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionsAPI) workspace-level service to start returning [iam.ObjectPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ObjectPermissions).
 * Changed `Update` method for [w.Permissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionsAPI) workspace-level service to start returning [iam.ObjectPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ObjectPermissions).
 * Added `GetPasswordPermissionLevels` method for [w.Users](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UsersAPI) workspace-level service.
 * Added `GetPasswordPermissions` method for [w.Users](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UsersAPI) workspace-level service.
 * Added `SetPasswordPermissions` method for [w.Users](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UsersAPI) workspace-level service.
 * Added `UpdatePasswordPermissions` method for [w.Users](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UsersAPI) workspace-level service.
 * Added `DisplayName` field for [iam.AccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#AccessControlResponse).
 * Changed `Roles` field for [iam.GetAssignableRolesForResourceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GetAssignableRolesForResourceResponse) to [iam.RoleList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#RoleList).
 * Added [iam.GetPasswordPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GetPasswordPermissionLevelsResponse).
 * Added [iam.PasswordAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordAccessControlRequest).
 * Added [iam.PasswordAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordAccessControlResponse).
 * Added [iam.PasswordPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordPermission).
 * Added [iam.PasswordPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordPermissionLevel).
 * Added [iam.PasswordPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordPermissions).
 * Added [iam.PasswordPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordPermissionsDescription).
 * Added [iam.PasswordPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordPermissionsRequest).
 * Added [iam.Role](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#Role).
 * Added `GetJobPermissionLevels` method for [w.Jobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsAPI) workspace-level service.
 * Added `GetJobPermissions` method for [w.Jobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsAPI) workspace-level service.
 * Added `SetJobPermissions` method for [w.Jobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsAPI) workspace-level service.
 * Added `UpdateJobPermissions` method for [w.Jobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsAPI) workspace-level service.
 * Added [jobs.GetJobPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetJobPermissionLevelsRequest).
 * Added [jobs.GetJobPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetJobPermissionLevelsResponse).
 * Added [jobs.GetJobPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetJobPermissionsRequest).
 * Added [jobs.JobAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobAccessControlRequest).
 * Added [jobs.JobAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobAccessControlResponse).
 * Added [jobs.JobPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobPermission).
 * Added [jobs.JobPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobPermissionLevel).
 * Added [jobs.JobPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobPermissions).
 * Added [jobs.JobPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobPermissionsDescription).
 * Added [jobs.JobPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobPermissionsRequest).
 * Added `GetExperimentPermissionLevels` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added `GetExperimentPermissions` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added `SetExperimentPermissions` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added `UpdateExperimentPermissions` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added `GetRegisteredModelPermissionLevels` method for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service.
 * Added `GetRegisteredModelPermissions` method for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service.
 * Added `SetRegisteredModelPermissions` method for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service.
 * Added `UpdateRegisteredModelPermissions` method for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service.
 * Added [ml.ExperimentAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentAccessControlRequest).
 * Added [ml.ExperimentAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentAccessControlResponse).
 * Added [ml.ExperimentPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentPermission).
 * Added [ml.ExperimentPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentPermissionLevel).
 * Added [ml.ExperimentPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentPermissions).
 * Added [ml.ExperimentPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentPermissionsDescription).
 * Added [ml.ExperimentPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentPermissionsRequest).
 * Added [ml.GetExperimentPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentPermissionLevelsRequest).
 * Added [ml.GetExperimentPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentPermissionLevelsResponse).
 * Added [ml.GetExperimentPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentPermissionsRequest).
 * Added [ml.GetRegisteredModelPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetRegisteredModelPermissionLevelsRequest).
 * Added [ml.GetRegisteredModelPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetRegisteredModelPermissionLevelsResponse).
 * Added [ml.GetRegisteredModelPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetRegisteredModelPermissionsRequest).
 * Added [ml.RegisteredModelAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelAccessControlRequest).
 * Added [ml.RegisteredModelAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelAccessControlResponse).
 * Added [ml.RegisteredModelPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelPermission).
 * Added [ml.RegisteredModelPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelPermissionLevel).
 * Added [ml.RegisteredModelPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelPermissions).
 * Added [ml.RegisteredModelPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelPermissionsDescription).
 * Added [ml.RegisteredModelPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelPermissionsRequest).
 * Added `Scopes` field for [oauth2.CreateCustomAppIntegration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateCustomAppIntegration).
 * Added `GetPipelinePermissionLevels` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
 * Added `GetPipelinePermissions` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
 * Added `SetPipelinePermissions` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
 * Added `UpdatePipelinePermissions` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
 * Added [pipelines.GetPipelinePermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelinePermissionLevelsRequest).
 * Added [pipelines.GetPipelinePermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelinePermissionLevelsResponse).
 * Added [pipelines.GetPipelinePermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelinePermissionsRequest).
 * Added [pipelines.PipelineAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineAccessControlRequest).
 * Added [pipelines.PipelineAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineAccessControlResponse).
 * Added [pipelines.PipelinePermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinePermission).
 * Added [pipelines.PipelinePermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinePermissionLevel).
 * Added [pipelines.PipelinePermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinePermissions).
 * Added [pipelines.PipelinePermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinePermissionsDescription).
 * Added [pipelines.PipelinePermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinePermissionsRequest).
 * Added `GcpManagedNetworkConfig` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest).
 * Added `GkeConfig` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest).
 * Added `GetServingEndpointPermissionLevels` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `GetServingEndpointPermissions` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `SetServingEndpointPermissions` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `UpdateServingEndpointPermissions` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `InstanceProfileArn` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput).
 * Added `InstanceProfileArn` field for [serving.ServedModelOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelOutput).
 * Added [serving.GetServingEndpointPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetServingEndpointPermissionLevelsRequest).
 * Added [serving.GetServingEndpointPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetServingEndpointPermissionLevelsResponse).
 * Added [serving.GetServingEndpointPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetServingEndpointPermissionsRequest).
 * Added [serving.ServingEndpointAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointAccessControlRequest).
 * Added [serving.ServingEndpointAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointAccessControlResponse).
 * Added [serving.ServingEndpointPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointPermission).
 * Added [serving.ServingEndpointPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointPermissionLevel).
 * Added [serving.ServingEndpointPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointPermissions).
 * Added [serving.ServingEndpointPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointPermissionsDescription).
 * Added [serving.ServingEndpointPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointPermissionsRequest).
 * Added `GetTokenPermissionLevels` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service.
 * Added `GetTokenPermissions` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service.
 * Added `SetTokenPermissions` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service.
 * Added `UpdateTokenPermissions` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service.
 * Added [settings.GetTokenPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetTokenPermissionLevelsResponse).
 * Added [settings.TokenAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenAccessControlRequest).
 * Added [settings.TokenAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenAccessControlResponse).
 * Added [settings.TokenPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenPermission).
 * Added [settings.TokenPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenPermissionLevel).
 * Added [settings.TokenPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenPermissions).
 * Added [settings.TokenPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenPermissionsDescription).
 * Added [settings.TokenPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenPermissionsRequest).
 * Added `GetWarehousePermissionLevels` method for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service.
 * Added `GetWarehousePermissions` method for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service.
 * Added `SetWarehousePermissions` method for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service.
 * Added `UpdateWarehousePermissions` method for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service.
 * Added `CanSubscribeToLiveQuery` field for [sql.QueryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryInfo).
 * Removed `QueuedOverloadTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Removed `QueuedProvisioningTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Removed `TotalFilesCount` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Removed `TotalPartitionsCount` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `MetadataTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `OverloadingQueueStartTimestamp` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `PlanningPhases` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `PlanningTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `ProvisioningQueueStartTimestamp` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `PrunedBytes` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `PrunedFilesCount` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `QueryCompilationStartTimestamp` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `QueryExecutionTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added [sql.GetWarehousePermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetWarehousePermissionLevelsRequest).
 * Added [sql.GetWarehousePermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetWarehousePermissionLevelsResponse).
 * Added [sql.GetWarehousePermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetWarehousePermissionsRequest).
 * Added [sql.WarehouseAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehouseAccessControlRequest).
 * Added [sql.WarehouseAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehouseAccessControlResponse).
 * Added [sql.WarehousePermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermission).
 * Added [sql.WarehousePermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermissionLevel).
 * Added [sql.WarehousePermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermissions).
 * Added [sql.WarehousePermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermissionsDescription).
 * Added [sql.WarehousePermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermissionsRequest).
 * Added `GetRepoPermissionLevels` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * Added `GetRepoPermissions` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * Added `SetRepoPermissions` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * Added `UpdateRepoPermissions` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * Added `GetWorkspaceObjectPermissionLevels` method for [w.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceAPI) workspace-level service.
 * Added `GetWorkspaceObjectPermissions` method for [w.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceAPI) workspace-level service.
 * Added `SetWorkspaceObjectPermissions` method for [w.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceAPI) workspace-level service.
 * Added `UpdateWorkspaceObjectPermissions` method for [w.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceAPI) workspace-level service.
 * Added [workspace.GetRepoPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetRepoPermissionLevelsRequest).
 * Added [workspace.GetRepoPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetRepoPermissionLevelsResponse).
 * Added [workspace.GetRepoPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetRepoPermissionsRequest).
 * Added [workspace.GetWorkspaceObjectPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetWorkspaceObjectPermissionLevelsRequest).
 * Added [workspace.GetWorkspaceObjectPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetWorkspaceObjectPermissionLevelsResponse).
 * Added [workspace.GetWorkspaceObjectPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetWorkspaceObjectPermissionsRequest).
 * Added [workspace.RepoAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoAccessControlRequest).
 * Added [workspace.RepoAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoAccessControlResponse).
 * Added [workspace.RepoPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoPermission).
 * Added [workspace.RepoPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoPermissionLevel).
 * Added [workspace.RepoPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoPermissions).
 * Added [workspace.RepoPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoPermissionsDescription).
 * Added [workspace.RepoPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoPermissionsRequest).
 * Added [workspace.WorkspaceObjectAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectAccessControlRequest).
 * Added [workspace.WorkspaceObjectAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectAccessControlResponse).
 * Added [workspace.WorkspaceObjectPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectPermission).
 * Added [workspace.WorkspaceObjectPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectPermissionLevel).
 * Added [workspace.WorkspaceObjectPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectPermissions).
 * Added [workspace.WorkspaceObjectPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectPermissionsDescription).
 * Added [workspace.WorkspaceObjectPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectPermissionsRequest).

OpenAPI SHA: 386b65ecdc825b9c3ed4aa7ca88e2e5baf9d87df, Date: 2023-08-07
Dependency updates:

 * Bump google.golang.org/api from 0.131.0 to 0.134.0 ([#564](https://github.com/databricks/databricks-sdk-go/pull/564)).

## 0.14.1

* Handle nested query parameters in Client.attempt ([#559](https://github.com/databricks/databricks-sdk-go/pull/559)).
* Support x-databricks-path-style overrides at the operation level ([#562](https://github.com/databricks/databricks-sdk-go/pull/562)).


## 0.14.0

* Fixed names of keyword arguments in examples ([#560](https://github.com/databricks/databricks-sdk-go/pull/560)).

API Changes:

 * Changed `Create` method for [a.AccountMetastoreAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoreAssignmentsAPI) account-level service to no longer return [catalog.CreateMetastoreAssignmentsResponseItemList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMetastoreAssignmentsResponseItemList).
 * Added `ConnectionName` field for [catalog.CreateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateCatalog).
 * Added `AccessPoint` field for [catalog.CreateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateExternalLocation).
 * Added `EncryptionDetails` field for [catalog.CreateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateExternalLocation).
 * Removed [catalog.CreateMetastoreAssignmentsResponseItem](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMetastoreAssignmentsResponseItem).
 * Added `AccessPoint` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo).
 * Added `EncryptionDetails` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo).
 * Added `AccessPoint` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).
 * Added `EncryptionDetails` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).
 * Added `AccessPoint` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
 * Added `EncryptionDetails` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
 * Added `AccessPoint` field for [catalog.VolumeInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumeInfo).
 * Added `EncryptionDetails` field for [catalog.VolumeInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumeInfo).
 * Added [catalog.EncryptionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EncryptionDetails).
 * Added [catalog.SseEncryptionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SseEncryptionDetails).
 * Added [catalog.SseEncryptionDetailsAlgorithm](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SseEncryptionDetailsAlgorithm).
 * Added [a.AccountNetworkPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyAPI) account-level service.
 * Added [settings.AccountNetworkPolicyMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyMessage).
 * Added [settings.DeleteAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAccountNetworkPolicyRequest).
 * Added [settings.DeleteAccountNetworkPolicyResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAccountNetworkPolicyResponse).
 * Added [settings.ReadAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadAccountNetworkPolicyRequest).
 * Added [settings.UpdateAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateAccountNetworkPolicyRequest).

OpenAPI SHA: fbdd0fa3e83fed2c798a58d376529bdb1285b915, Date: 2023-07-26

## 0.13.0

* Add issue templates ([#539](https://github.com/databricks/databricks-sdk-go/pull/539)).
* Added HasRequiredNonBodyField method ([#536](https://github.com/databricks/databricks-sdk-go/pull/536)).
* Make Azure MSI auth account compatible ([#544](https://github.com/databricks/databricks-sdk-go/pull/544)).
* Refactor Handling of Name<->ID Mapping in OpenAPI Generator ([#547](https://github.com/databricks/databricks-sdk-go/pull/547)).
* Regenerate Go SDK from current OpenAPI Specification ([#549](https://github.com/databricks/databricks-sdk-go/pull/549)).
* Parse Camel Case and Pascal Case Enum Values ([#550](https://github.com/databricks/databricks-sdk-go/pull/550)).
* Prepare for auto-releaser infra ([#554](https://github.com/databricks/databricks-sdk-go/pull/554)).
* Added SCIM Patch Acceptance Tests ([#540](https://github.com/databricks/databricks-sdk-go/pull/540)).

API Changes:

 * Removed `Maintenance` method for [w.Metastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoresAPI) workspace-level service.
 * Added `EnableOptimization` method for [w.Metastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoresAPI) workspace-level service.
 * Added `Update` method for [w.Tables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TablesAPI) workspace-level service.
 * Added `Force` field for [catalog.DeleteAccountMetastoreRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteAccountMetastoreRequest).
 * Added `Force` field for [catalog.DeleteAccountStorageCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteAccountStorageCredentialRequest).
 * Removed [catalog.UpdateAutoMaintenance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateAutoMaintenance).
 * Removed [catalog.UpdateAutoMaintenanceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateAutoMaintenanceResponse).
 * Added [catalog.UpdatePredictiveOptimization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePredictiveOptimization).
 * Added [catalog.UpdatePredictiveOptimizationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePredictiveOptimizationResponse).
 * Added [catalog.UpdateTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateTableRequest).
 * Added `Schema` field for [iam.PartialUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PartialUpdate).
 * Added [iam.PatchSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PatchSchema).
 * Added `TriggerInfo` field for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
 * Added `Health` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
 * Added `JobSource` field for [jobs.GitSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GitSource).
 * Added `OnDurationWarningThresholdExceeded` field for [jobs.JobEmailNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobEmailNotifications).
 * Added `Health` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Added `TriggerInfo` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
 * Added `RunJobOutput` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
 * Added `RunJobTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
 * Added `EmailNotifications` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
 * Added `Health` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
 * Added `EmailNotifications` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added `Health` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added `NotificationSettings` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added `Health` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
 * Added `RunJobTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
 * Added `OnDurationWarningThresholdExceeded` field for [jobs.TaskEmailNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TaskEmailNotifications).
 * Added `OnDurationWarningThresholdExceeded` field for [jobs.WebhookNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#WebhookNotifications).
 * Added [jobs.JobSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSource).
 * Added [jobs.JobSourceDirtyState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSourceDirtyState).
 * Added [jobs.JobsHealthMetric](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsHealthMetric).
 * Added [jobs.JobsHealthOperator](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsHealthOperator).
 * Added [jobs.JobsHealthRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsHealthRule).
 * Added [jobs.JobsHealthRules](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsHealthRules).
 * Added [jobs.RunJobOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunJobOutput).
 * Added [jobs.RunJobTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunJobTask).
 * Added [jobs.TriggerInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerInfo).
 * Added [jobs.WebhookNotificationsOnDurationWarningThresholdExceededItem](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#WebhookNotificationsOnDurationWarningThresholdExceededItem).
 * Removed `Whl` field for [pipelines.PipelineLibrary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineLibrary).
 * Changed `DeletePersonalComputeSetting` method for [a.AccountSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountSettingsAPI) account-level service with new required argument order.
 * Changed `ReadPersonalComputeSetting` method for [a.AccountSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountSettingsAPI) account-level service with new required argument order.
 * Changed `Etag` field for [settings.DeletePersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeletePersonalComputeSettingRequest) to be required.
 * Changed `Etag` field for [settings.ReadPersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadPersonalComputeSettingRequest) to be required.
 * Added [w.CleanRooms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomsAPI) workspace-level service.
 * Added [sharing.CentralCleanRoomInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CentralCleanRoomInfo).
 * Added [sharing.CleanRoomAssetInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomAssetInfo).
 * Added [sharing.CleanRoomCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomCatalog).
 * Added [sharing.CleanRoomCatalogUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomCatalogUpdate).
 * Added [sharing.CleanRoomCollaboratorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomCollaboratorInfo).
 * Added [sharing.CleanRoomInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomInfo).
 * Added [sharing.CleanRoomNotebookInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomNotebookInfo).
 * Added [sharing.CleanRoomTableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomTableInfo).
 * Added [sharing.ColumnInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ColumnInfo).
 * Added [sharing.ColumnMask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ColumnMask).
 * Added [sharing.ColumnTypeName](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ColumnTypeName).
 * Added [sharing.CreateCleanRoom](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CreateCleanRoom).
 * Added [sharing.DeleteCleanRoomRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#DeleteCleanRoomRequest).
 * Added [sharing.GetCleanRoomRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#GetCleanRoomRequest).
 * Added [sharing.ListCleanRoomsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListCleanRoomsResponse).
 * Added [sharing.UpdateCleanRoom](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateCleanRoom).
 * Changed `Query` field for [sql.Alert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Alert) to [sql.AlertQuery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertQuery).
 * Changed `Value` field for [sql.AlertOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertOptions) to `any`.
 * Removed `IsDbAdmin` field for [sql.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#User).
 * Removed `ProfileImageUrl` field for [sql.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#User).
 * Added [sql.AlertQuery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertQuery).

OpenAPI SHA: 0a1949ba96f71680dad30e06973eaae85b1307bb, Date: 2023-07-18
Dependency updates:

 * Bump golang.org/x/oauth2 from 0.9.0 to 0.10.0 ([#545](https://github.com/databricks/databricks-sdk-go/pull/545)).
 * Bump golang.org/x/mod from 0.11.0 to 0.12.0 ([#541](https://github.com/databricks/databricks-sdk-go/pull/541)).
 * Bump google.golang.org/api from 0.129.0 to 0.130.0 ([#542](https://github.com/databricks/databricks-sdk-go/pull/542)).
 * Bump google.golang.org/api from 0.130.0 to 0.131.0 ([#551](https://github.com/databricks/databricks-sdk-go/pull/551)).

## 0.12.0

 * Change release type
 * Regenerate SDK according to OpenAPI spec ([#534](https://github.com/databricks/databricks-sdk-go/pull/534)).
 * Bump google.golang.org/api from 0.128.0 to 0.129.0 ([#533](https://github.com/databricks/databricks-sdk-go/pull/533)).

## 0.11.0

 * Regenerated from OpenAPI spec ([#527](https://github.com/databricks/databricks-sdk-go/pull/527)). Includes bugfixes for System Tables.
 * Support accounts-dod as account endpoint ([#530](https://github.com/databricks/databricks-sdk-go/pull/530)).

Breaking API Changes:

 - Renamed compute.ClusterInfo to compute.ClusterDetails
 - Renamed compue.BaseClusterInfo to compute.ClusterSpec
 - Renamed jobs.RunSubmitTaskSettings to jobs.SubmitTask
 - Renamed jobs.JobTaskSettings to jobs.Task
 - Renamed jobs.CreateJobFormat to jobs.Format
 - Renamed jobs.JobsWebhookNotifications to jobs.WebhookNotifications
 - Renamed jobs.CronSchedulePauseStatus to jobs.PauseStatus
 - Renamed jobs.FileArrivalTriggerSettings to jobs.FileArrivalTriggerConfiguration
 - Renamed jobs.GitSourceGitProvider to jobs.GitProvider
 - Consolidated jobs.NotebookTaskSource, jobs.SparkPythonTaskSource to jobs.Source
 - Consolidated jobs.JobWebhookNotificationsOnFailureItem, jobs.JobWebhookNotificationsOnStartItem, jobs.JobWebhookNotificationsOnSuccessItem to jobs.Webhook
 - Renamed jobs.TaskDependenciesItem to job.TaskDependency
 - jobs.SparkSubmitTask.Widgets changed from *SqlDashboardWidgetOutput to []SqlDashboardWidgetOutput
 - (a *SystemSchemasAPI) DisableByMetastoreIdAndSchemaName changed from func(context.Context, string, string) error to func(context.Context, string, DisableSchemaName) error
 - (a *SystemSchemasAPI) Enable changed from func(context.Context) error to func(context.Context, EnableRequest) error
 - catalog.CatalogInfo.OptionsKvpairs changed from *OptionsKvPairs to map[string]string
 - catalog.CreateConnection.OptionsKvpairs changed from *OptionsKvPairs to map[string]string
 - catalog.UpdateConnection.OptionsKvpairs changed from OptionsKvPairs to map[string]string
 - catalog.DisableRequest.SchemaName changed from string to DisableSchemaName
 - catalog.OptionsKvPairs removed
 - catalog.SystemSchemaInfoStateDisableinitialized renamed to catalog.SystemSchemaInfoStateDisabledInitialized
 - catalog.SystemSchemaInfoStateEnabledcompleted renamed to catalog.SystemSchemaInfoStateEnabledCompleted
 - catalog.SystemSchemaInfoStateEnabledinitialized renamed to catalog.SystemSchemaInfoStateEnabledInitialized
 - settings.ReadPersonalComputeSettingsResponse renamed to settings.PersonalComputeSetting
 - workpace.CreateScope.KeyvaultMetadata removed

Other API Changes:

 - catalog.CatalogInfo.ConnectionName added
 - catalog.CatalogInfo.Options added
 - catalog.DisableSchemaName added
 - catalog.EnableRequest added
 - catalog.PrivilegeUseMarketplaceAssets added
 - compute.ClusterEvent added
 - compute.ComputeSpec added
 - compute.ComputeSpecKind added
 - jobs.ConditionTask added
 - jobs.ConditionTaskOp added
 - jobs.CreateJob.Compute added
 - jobs.JobCompute added
 - jobs.JobSettings.Compute added
 - jobs.RunConditionTask added
 - jobs.RunOutput.ConditionTask added
 - jobs.RunTask.ConditionTask added
 - serving.ServedModelInput.EnvironmentVars added
 - serving.ServedModelOutput.EnvironmentVars added
 - settings.DeletePersonalComputeSetting added
 - settings.UpdatePersonalComputeSetting added
 - sharing.PrivilegeUseMarketplaceAssets added
 - workspace.CreateScope.BackendAzureKeyVault added

## 0.10.1

 * Added ClusterID field for Config ([#524](https://github.com/databricks/databricks-sdk-go/pull/524)).
 * Added OnlyPayload property for method ([#522](https://github.com/databricks/databricks-sdk-go/pull/522)).
 * Renamed IsJsonOnly method ([#523](https://github.com/databricks/databricks-sdk-go/pull/523)).

Dependency updates:

* Bump google.golang.org/api from 0.127.0 to 0.128.0 ([#521](https://github.com/databricks/databricks-sdk-go/pull/521)).

## 0.10.0

 * Added log level support to SimpleLogger. Default logger now prints only `INFO` level messages. To replicate more verbose behavior from the previous versions, set the `DEBUG` level in `SimpleLogger` explicitly ([#426](https://github.com/databricks/databricks-sdk-go/pull/426)).
 * Added `Upload` and `Download` methods where applicable ([#423](https://github.com/databricks/databricks-sdk-go/pull/423)).
 * Added auth tests for a databrickscfg file with an empty DEFAULT profile ([#496](https://github.com/databricks/databricks-sdk-go/pull/496)).
 * Added more integration tests and examples ([#503](https://github.com/databricks/databricks-sdk-go/pull/503)).
 * Added preview level functions for services and methods to OpenAPI code ([#497](https://github.com/databricks/databricks-sdk-go/pull/497)).
 * Added IsAllRequiredFieldsPrimitive helper function ([#404](https://github.com/databricks/databricks-sdk-go/pull/404)).
 * Added isCrudCreate method ([#416](https://github.com/databricks/databricks-sdk-go/pull/416)).
 * Added retries for getting refreshable token ([#500](https://github.com/databricks/databricks-sdk-go/pull/500)).
 * Added type update when entity fields are updated ([#411](https://github.com/databricks/databricks-sdk-go/pull/411)).
 * Don't panic but return error from NewAccountClient ([#422](https://github.com/databricks/databricks-sdk-go/pull/422)).
 * Error when opening a DBFS directory for reading ([#415](https://github.com/databricks/databricks-sdk-go/pull/415)).
 * Fixed cleanup logic in SQL warehouse integration test ([#400](https://github.com/databricks/databricks-sdk-go/pull/400)).
 * Fixed error handling in cluster utility function ([#399](https://github.com/databricks/databricks-sdk-go/pull/399)).
 * Fixed example in README.md ([#494](https://github.com/databricks/databricks-sdk-go/pull/494)).
 * Fixed nondeterminism in workspace filesystem integration test ([#401](https://github.com/databricks/databricks-sdk-go/pull/401)).
 * Improve command execution interface ([#410](https://github.com/databricks/databricks-sdk-go/pull/410)).
 * Introduce waiters as top-level methods ([#408](https://github.com/databricks/databricks-sdk-go/pull/408)).
 * Regenerate examples if configured ([#518](https://github.com/databricks/databricks-sdk-go/pull/518)).
 * Respect limit field passed in ListRequests ([#407](https://github.com/databricks/databricks-sdk-go/pull/407)).
 * Updated from OpenAPI spec ([#412](https://github.com/databricks/databricks-sdk-go/pull/412), [#413](https://github.com/databricks/databricks-sdk-go/pull/413), [#421](https://github.com/databricks/databricks-sdk-go/pull/421), [#519](https://github.com/databricks/databricks-sdk-go/pull/519), [#424](https://github.com/databricks/databricks-sdk-go/pull/424)).
 * Updated API test template with page tokens ([#417](https://github.com/databricks/databricks-sdk-go/pull/417)).
 * Use constants instead of hardcoding strings ([#402](https://github.com/databricks/databricks-sdk-go/pull/402)).
 * Use `x-databricks-is-accounts` flag to determine whether a service is an account level service ([#420](https://github.com/databricks/databricks-sdk-go/pull/420)).

API Changes:

github.com/databricks/databricks-sdk-go/service/serving
  - (*ServingEndpointsAPI).Create: changed from func(context.Context, CreateServingEndpoint) (*ServingEndpointDetailed, error) to func(context.Context, CreateServingEndpoint) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error)
  - (*ServingEndpointsAPI).List: removed
  - (*ServingEndpointsAPI).UpdateConfig: changed from func(context.Context, EndpointCoreConfigInput) (*ServingEndpointDetailed, error) to func(context.Context, EndpointCoreConfigInput) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error)
  Compatible changes:
  - (*ServingEndpointsAPI).ListAll: added
  - (*ServingEndpointsAPI).WaitGetServingEndpointNotUpdating: added
  - WaitGetServingEndpointNotUpdating: added

github.com/databricks/databricks-sdk-go
  - WorkspaceClient.CommandExecutor: removed
  - WorkspaceClient.CommandExecution: added
  - WorkspaceClient.Connections: added
  - WorkspaceClient.Files: added
  - WorkspaceClient.SystemSchemas: added

github.com/databricks/databricks-sdk-go/service/sql
  - (*WarehousesAPI).Create: changed from func(context.Context, CreateWarehouseRequest) (*CreateWarehouseResponse, error) to func(context.Context, CreateWarehouseRequest) (*WaitGetWarehouseRunning[CreateWarehouseResponse], error)
  - (*WarehousesAPI).DeleteAndWait: removed
  - (*WarehousesAPI).DeleteByIdAndWait: removed
  - (*WarehousesAPI).Edit: changed from func(context.Context, EditWarehouseRequest) error to func(context.Context, EditWarehouseRequest) (*WaitGetWarehouseRunning[any], error)
  - (*WarehousesAPI).GetAndWait: removed
  - (*WarehousesAPI).GetByIdAndWait: removed
  - (*WarehousesAPI).Start: changed from func(context.Context, StartRequest) error to func(context.Context, StartRequest) (*WaitGetWarehouseRunning[any], error)
  - (*WarehousesAPI).Stop: changed from func(context.Context, StopRequest) error to func(context.Context, StopRequest) (*WaitGetWarehouseStopped[any], error)
  - (*StatementExecutionAPI).ExecuteAndWait: added
  - (*WarehousesAPI).WaitGetWarehouseRunning: added
  - (*WarehousesAPI).WaitGetWarehouseStopped: added
  - FormatCsv: added
  - WaitGetWarehouseRunning: added
  - WaitGetWarehouseStopped: added

github.com/databricks/databricks-sdk-go/service/jobs
  - (*JobsAPI).CancelRun: changed from func(context.Context, CancelRun) error to func(context.Context, CancelRun) (*WaitGetRunJobTerminatedOrSkipped[any], error)
  - (*JobsAPI).GetRunAndWait: removed
  - (*JobsAPI).RepairRun: changed from func(context.Context, RepairRun) (*RepairRunResponse, error) to func(context.Context, RepairRun) (*WaitGetRunJobTerminatedOrSkipped[RepairRunResponse], error)
  - (*JobsAPI).RunNow: changed from func(context.Context, RunNow) (*RunNowResponse, error) to func(context.Context, RunNow) (*WaitGetRunJobTerminatedOrSkipped[RunNowResponse], error)
  - (*JobsAPI).Submit: changed from func(context.Context, SubmitRun) (*SubmitRunResponse, error) to func(context.Context, SubmitRun) (*WaitGetRunJobTerminatedOrSkipped[SubmitRunResponse], error)
  - (*JobsAPI).WaitGetRunJobTerminatedOrSkipped: added
  - CreateJob.RunAs: added
  - JobRunAs: added
  - JobSettings.RunAs: added
  - WaitGetRunJobTerminatedOrSkipped: added

github.com/databricks/databricks-sdk-go/service/compute
  - (*ClustersAPI).Create: changed from func(context.Context, CreateCluster) (*CreateClusterResponse, error) to func(context.Context, CreateCluster) (*WaitGetClusterRunning[CreateClusterResponse], error)
  - (*ClustersAPI).Delete: changed from func(context.Context, DeleteCluster) error to func(context.Context, DeleteCluster) (*WaitGetClusterTerminated[any], error)
  - (*ClustersAPI).Edit: changed from func(context.Context, EditCluster) error to func(context.Context, EditCluster) (*WaitGetClusterRunning[any], error)
  - (*ClustersAPI).GetAndWait: removed
  - (*ClustersAPI).GetByClusterIdAndWait: removed
  - (*ClustersAPI).Resize: changed from func(context.Context, ResizeCluster) error to func(context.Context, ResizeCluster) (*WaitGetClusterRunning[any], error)
  - (*ClustersAPI).Restart: changed from func(context.Context, RestartCluster) error to func(context.Context, RestartCluster) (*WaitGetClusterRunning[any], error)
  - (*ClustersAPI).Start: changed from func(context.Context, StartCluster) error to func(context.Context, StartCluster) (*WaitGetClusterRunning[any], error)
  - (*CommandExecutionAPI).Cancel: changed from func(context.Context, CancelCommand) error to func(context.Context, CancelCommand) (*WaitCommandStatusCommandExecutionCancelled[any], error)
  - (*CommandExecutionAPI).Create: changed from func(context.Context, CreateContext) (*Created, error) to func(context.Context, CreateContext) (*WaitContextStatusCommandExecutionRunning[Created], error)
  - (*CommandExecutionAPI).Execute: changed from func(context.Context, Command) (*Created, error) to func(context.Context, Command) (*WaitCommandStatusCommandExecutionFinishedOrError[Created], error)
  - (*ClustersAPI).WaitGetClusterRunning: added
  - (*ClustersAPI).WaitGetClusterTerminated: added
  - (*CommandExecutionAPI).Start: added
  - (*CommandExecutionAPI).WaitCommandStatusCommandExecutionCancelled: added
  - (*CommandExecutionAPI).WaitCommandStatusCommandExecutionFinishedOrError: added
  - (*CommandExecutionAPI).WaitContextStatusCommandExecutionRunning: added
  - BaseClusterInfo.DataSecurityMode: added
  - BaseClusterInfo.DockerImage: added
  - BaseClusterInfo.SingleUserName: added
  - ClusterAttributes.DataSecurityMode: added
  - ClusterAttributes.DockerImage: added
  - ClusterAttributes.SingleUserName: added
  - ClusterInfo.DockerImage: added
  - CommandExecutorV2: added
  - EditCluster.DataSecurityMode: added
  - EditCluster.DockerImage: added
  - EditCluster.SingleUserName: added
  - WaitCommandStatusCommandExecutionCancelled: added
  - WaitCommandStatusCommandExecutionFinishedOrError: added
  - WaitContextStatusCommandExecutionRunning: added
  - WaitGetClusterRunning: added
  - WaitGetClusterTerminated: added

github.com/databricks/databricks-sdk-go/service/pipelines
  - (*PipelinesAPI).GetAndWait: removed
  - (*PipelinesAPI).GetByPipelineIdAndWait: removed
  - (*PipelinesAPI).Reset: changed from func(context.Context, ResetRequest) error to func(context.Context, ResetRequest) (*WaitGetPipelineRunning[any], error)
  - (*PipelinesAPI).Stop: changed from func(context.Context, StopRequest) error to func(context.Context, StopRequest) (*WaitGetPipelineIdle[any], error)
  - (*PipelinesAPI).WaitGetPipelineIdle: added
  - (*PipelinesAPI).WaitGetPipelineRunning: added
  - WaitGetPipelineIdle: added
  - WaitGetPipelineRunning: added

github.com/databricks/databricks-sdk-go/service/sharing
  - (*ProvidersAPI).ListShares: removed
  - (*ProvidersAPI).ListSharesAll: added

github.com/databricks/databricks-sdk-go/service/workspace
  - ExportFormatAuto: removed
  - ExportRequest.DirectDownload: removed
  - Import.Format: changed from ExportFormat to ImportFormat
  - (*WorkspaceAPI).Download: added
  - (*WorkspaceAPI).ReadFile: added
  - (*WorkspaceAPI).Upload: added
  - (*WorkspaceAPI).WriteFile: added
  - DownloadFormat: added
  - DownloadOption: added
  - ImportFormat: added
  - ImportFormatAuto: added
  - ImportFormatDbc: added
  - ImportFormatHtml: added
  - ImportFormatJupyter: added
  - ImportFormatRMarkdown: added
  - ImportFormatSource: added
  - UploadFormat: added
  - UploadLanguage: added
  - UploadOption: added
  - UploadOverwrite: added

github.com/databricks/databricks-sdk-go/service/catalog
  - (*FunctionsAPI).List: removed
  - (*MetastoresAPI).Assign: changed from func(context.Context, CreateMetastoreAssignment) error to func(context.Context, CreateMetastoreAssignment) error
  - (*MetastoresAPI).Create: changed from func(context.Context, CreateMetastore) (*MetastoreInfo, error) to func(context.Context, CreateMetastore) (*MetastoreInfo, error)
  - (*MetastoresAPI).Update: changed from func(context.Context, UpdateMetastore) (*MetastoreInfo, error) to func(context.Context, UpdateMetastore) (*MetastoreInfo, error)
  - (*MetastoresAPI).UpdateAssignment: changed from func(context.Context, UpdateMetastoreAssignment) error to func(context.Context, UpdateMetastoreAssignment) error
  - (*StorageCredentialsAPI).Create: changed from func(context.Context, CreateStorageCredential) (*StorageCredentialInfo, error) to func(context.Context, CreateStorageCredential) (*StorageCredentialInfo, error)
  - (*StorageCredentialsAPI).Update: changed from func(context.Context, UpdateStorageCredential) (*StorageCredentialInfo, error) to func(context.Context, UpdateStorageCredential) (*StorageCredentialInfo, error)
  - (*TablesAPI).ListSummaries: removed
  - CreateMetastore.Name: removed
  - CreateMetastore.Region: removed
  - CreateMetastore.StorageRoot: removed
  - CreateMetastore: changed from CreateMetastore to CreateMetastore
  - CreateMetastoreAssignment.DefaultCatalogName: removed
  - CreateMetastoreAssignment: changed from CreateMetastoreAssignment to CreateMetastoreAssignment
  - CreateStorageCredential.AwsIamRole: removed
  - CreateStorageCredential.AzureServicePrincipal: removed
  - CreateStorageCredential.Comment: removed
  - CreateStorageCredential.GcpServiceAccountKey: removed
  - CreateStorageCredential.Name: removed
  - CreateStorageCredential.ReadOnly: removed
  - CreateStorageCredential.SkipValidation: removed
  - CreateStorageCredential: changed from CreateStorageCredential to CreateStorageCredential
  - GcpServiceAccountKey: removed
  - ListFunctionsResponse.Schemas: removed
  - MetastoreAssignment.WorkspaceId: changed from string to int64
  - StorageCredentialInfo.GcpServiceAccountKey: removed
  - UpdateMetastore.DeltaSharingOrganizationName: removed
  - UpdateMetastore.DeltaSharingRecipientTokenLifetimeInSeconds: removed
  - UpdateMetastore.DeltaSharingScope: removed
  - UpdateMetastore.Id: removed
  - UpdateMetastore.Name: removed
  - UpdateMetastore.Owner: removed
  - UpdateMetastore.PrivilegeModelVersion: removed
  - UpdateMetastore.StorageRootCredentialId: removed
  - UpdateMetastore: changed from UpdateMetastore to UpdateMetastore
  - UpdateMetastoreAssignment.DefaultCatalogName: removed
  - UpdateMetastoreAssignment: changed from UpdateMetastoreAssignment to UpdateMetastoreAssignment
  - UpdateStorageCredential.AwsIamRole: removed
  - UpdateStorageCredential.AzureServicePrincipal: removed
  - UpdateStorageCredential.Comment: removed
  - UpdateStorageCredential.Force: removed
  - UpdateStorageCredential.GcpServiceAccountKey: removed
  - UpdateStorageCredential.Owner: removed
  - UpdateStorageCredential.ReadOnly: removed
  - UpdateStorageCredential.SkipValidation: removed
  - UpdateStorageCredential: changed from UpdateStorageCredential to UpdateStorageCredential
  - ValidateStorageCredential.GcpServiceAccountKey: removed
  - (*FunctionsAPI).FunctionInfoNameToFullNameMap: added
  - (*FunctionsAPI).ListAll: added
  - (*TablesAPI).ListSummariesAll: added
  - AccountsCreateMetastore: added
  - AccountsCreateMetastoreAssignment: added
  - AccountsCreateStorageCredential: added
  - AccountsUpdateMetastore: added
  - AccountsUpdateMetastoreAssignment: added
  - AccountsUpdateStorageCredential: added
  - AzureManagedIdentity: added
  - ConnectionInfo: added
  - ConnectionType: added
  - ConnectionTypeDatabricks: added
  - ConnectionTypeMysql: added
  - ConnectionTypePostgresql: added
  - ConnectionTypeRedshift: added
  - ConnectionTypeSnowflake: added
  - ConnectionTypeSqldw: added
  - ConnectionTypeSqlserver: added
  - ConnectionsAPI: added
  - ConnectionsService: added
  - CreateConnection: added
  - CreateMetastore.MetastoreInfo: added
  - CreateMetastoreAssignment.MetastoreAssignment: added
  - CreateStorageCredential.CredentialInfo: added
  - CredentialType: added
  - CredentialTypeUsernamePassword: added
  - DatabricksGcpServiceAccountResponse: added
  - DeleteConnectionRequest: added
  - DisableRequest: added
  - GetConnectionRequest: added
  - ListConnectionsResponse: added
  - ListFunctionsResponse.Functions: added
  - ListSystemSchemasRequest: added
  - ListSystemSchemasResponse: added
  - NewConnections: added
  - NewSystemSchemas: added
  - OptionsKvPairs: added
  - PropertiesKvPairs: added
  - StorageCredentialInfo.AzureManagedIdentity: added
  - StorageCredentialInfo.DatabricksGcpServiceAccount: added
  - SystemSchemaInfo: added
  - SystemSchemaInfoState: added
  - SystemSchemaInfoStateDisableinitialized: added
  - SystemSchemaInfoStateEnablecompleted: added
  - SystemSchemaInfoStateEnableinitialized: added
  - SystemSchemaInfoStateUnavailable: added
  - SystemSchemasAPI: added
  - SystemSchemasService: added
  - UpdateConnection: added
  - UpdateMetastore.MetastoreInfo: added
  - UpdateMetastoreAssignment.MetastoreAssignment: added
  - UpdateStorageCredential.CredentialInfo: added
  - ValidateStorageCredential.AzureManagedIdentity: added
  - ValidateStorageCredential.DatabricksGcpServiceAccount: added

github.com/databricks/databricks-sdk-go/service/ml
  - GetModelResponse.RegisteredModel: removed
  - GetModelResponse.RegisteredModelDatabricks: added

github.com/databricks/databricks-sdk-go/service/files
  - FilesAPI: added
  - FilesService: added
  - NewFiles: added

github.com/databricks/databricks-sdk-go/logger
  - SimpleLogger.Level: added

github.com/databricks/databricks-sdk-go/service/provisioning
  - (*WorkspacesAPI).Create: changed from func(context.Context, CreateWorkspaceRequest) (*Workspace, error) to func(context.Context, CreateWorkspaceRequest) (*WaitGetWorkspaceRunning[Workspace], error)
  - (*WorkspacesAPI).Update: changed from func(context.Context, UpdateWorkspaceRequest) error to func(context.Context, UpdateWorkspaceRequest) (*WaitGetWorkspaceRunning[any], error)
  - (*WorkspacesAPI).WaitGetWorkspaceRunning: added
  - WaitGetWorkspaceRunning: added

github.com/databricks/databricks-sdk-go/service/iam
  - AccountAccessControlAPI: added
  - AccountAccessControlProxyAPI: added
  - AccountAccessControlProxyService: added
  - AccountAccessControlService: added
  - GetAssignableRolesForResourceRequest: added
  - GetAssignableRolesForResourceResponse: added
  - GetRuleSetRequest: added
  - GrantRule: added
  - NewAccountAccessControl: added
  - NewAccountAccessControlProxy: added
  - RuleSetResponse: added
  - RuleSetUpdateRequest: added
  - UpdateRuleSetRequest: added

Dependency updates:

 * Bump github.com/stretchr/testify from 1.8.3 to 1.8.4 ([#406](https://github.com/databricks/databricks-sdk-go/pull/406)).
 * Bump golang.org/x/mod from 0.10.0 to 0.11.0 ([#515](https://github.com/databricks/databricks-sdk-go/pull/515)).
 * Bump golang.org/x/oauth2 from 0.8.0 to 0.9.0 ([#498](https://github.com/databricks/databricks-sdk-go/pull/498)).
 * Bump google.golang.org/api from 0.123.0 to 0.127.0 ([#405](https://github.com/databricks/databricks-sdk-go/pull/405), [#425](https://github.com/databricks/databricks-sdk-go/pull/425), [#429](https://github.com/databricks/databricks-sdk-go/pull/429)).

## 0.9.0

 * Added more usage examples for `go doc` and Go Packages ([#389](https://github.com/databricks/databricks-sdk-go/pull/389)).
 * Make u2m authentication work with new CLI ([#394](https://github.com/databricks/databricks-sdk-go/pull/394)).
 * Update from OpenAPI spec (19 may) ([#390](https://github.com/databricks/databricks-sdk-go/pull/390)).

Dependency updates:

 * Bump github.com/stretchr/testify from 1.8.2 to 1.8.3 ([#393](https://github.com/databricks/databricks-sdk-go/pull/393)).
 * Bump google.golang.org/api from 0.122.0 to 0.123.0 ([#392](https://github.com/databricks/databricks-sdk-go/pull/392)).

## 0.8.1

 * Added `in` codegen function ([#387](https://github.com/databricks/databricks-sdk-go/pull/387)).
 * Fixed mlflow acceptance tests ([#378](https://github.com/databricks/databricks-sdk-go/pull/378)).
 * Fixed MLflow integration test and removed workaround for `DELETE` query parameter ([#380](https://github.com/databricks/databricks-sdk-go/pull/380)).
 * Make clusters acceptance tests robust to duplicate cluster names ([#381](https://github.com/databricks/databricks-sdk-go/pull/381)).
 * Remove dead code from apierr/errors.go ([#376](https://github.com/databricks/databricks-sdk-go/pull/376)).
 * Serialize params to request body on delete ([#383](https://github.com/databricks/databricks-sdk-go/pull/383)).

Dependency updates:

 * Bump golang.org/x/oauth2 from 0.7.0 to 0.8.0 ([#385](https://github.com/databricks/databricks-sdk-go/pull/385)).
 * Bump google.golang.org/api from 0.118.0 to 0.122.0 ([#382](https://github.com/databricks/databricks-sdk-go/pull/382), [#386](https://github.com/databricks/databricks-sdk-go/pull/386)).

## 0.8.0

 * Added more code generation utilities ([#369](https://github.com/databricks/databricks-sdk-go/pull/369)).
 * Body logger for non-JSON payload as well ([#365](https://github.com/databricks/databricks-sdk-go/pull/365)).
 * Cleanup ephemeral groups in integration tests ([#368](https://github.com/databricks/databricks-sdk-go/pull/368)).
 * Fixed external entity generation ([#372](https://github.com/databricks/databricks-sdk-go/pull/372)).
 * Skip loading default profile if host is already configured ([#363](https://github.com/databricks/databricks-sdk-go/pull/363)).
 * Update debug messages in config loader to display correct path ([#362](https://github.com/databricks/databricks-sdk-go/pull/362)).

Dependency updates:

 * Bump golang.org/x/oauth2 from 0.6.0 to 0.7.0 ([#364](https://github.com/databricks/databricks-sdk-go/pull/364)).
 * Bump google.golang.org/api from 0.115.0 to 0.116.0 ([#361](https://github.com/databricks/databricks-sdk-go/pull/361)).
 * Bump google.golang.org/api from 0.116.0 to 0.118.0 ([#367](https://github.com/databricks/databricks-sdk-go/pull/367)).
 
API changes:

 * Moved `clusterpolicies` APIs to `compute` package.
 * Moved `clusters` APIs to `compute` package.
 * Moved `commands` APIs to `compute` package.
 * Moved `globalinitscripts` APIs to `compute` package.
 * Moved `instancepools` APIs to `compute` package.
 * Moved `scim` APIs to `iam` package.
 * Moved `permissions` APIs to `iam` package.
 * Moved `ipaccesslists` APIs to `settings` package.
 * Moved `tokenmanagement` APIs to `settings` package.
 * Moved `tokens` APIs to `settings` package.
 * Moved `workspaceconf` APIs to `settings` package.
 * Moved `gitcredentials` APIs to `workspace` package.
 * Moved `repos` APIs to `workspace` package.
 * Moved `secrets` APIs to `workspace` package.
 * Split `unitcatalog` package to `catalog` and `sharing`.
 * Renamed `mlflow` package to `ml`.
 * Renamed `dbfs` package to `files`.
 * Renamed `deployment` package to `provisioning`.
 * Renamed `endpoints` package to `serving`.
 * Renamed `clusters.List` type to `compute.ListClustersRequest`.
 * Renamed `jobs.ListRuns` type to `jobs.ListRunsRequest`.
 * Renamed `jobs.ExportRun` type to `jobs.ExportRunRequest`.
 * Renamed `clusterpolicies.List` type to `compute.ListClusterPoliciesRequest`.
 * Renamed `jobs.List` type to `jobs.ListJobsRequest`.
 * Renamed `permissions.GetPermissionLevels` type to `iam.GetPermissionLevelsRequest`.
 * Renamed `pipelines.ListPipelineEvents` type to `pipelines.ListPipelineEventsRequest`.
 * Renamed `pipelines.ListPipelines` type to `pipelines.ListPipelinesRequest`.
 * Renamed `workspaceconf.GetStatus` type to `settings.GetStatusRequest`.
 * Renamed `repos.List` type to `workspace.ListReposRequest`.
 * Renamed `tokenmanagement.List` type to `settings.ListTokenManagementRequest`.
 * Renamed `workspace.Export` type to `workspace.ExportRequest`.
 * Renamed `workspace.List` type to `workspace.ListWorkspaceRequest`.

## 0.7.0

 * Update from OpenAPI ([#359](https://github.com/databricks/databricks-sdk-go/pull/359)).
 * Experimental credentials provider via local server ([#340](https://github.com/databricks/databricks-sdk-go/pull/340)).
 * Added `isTesting` marker to support `ResourceFixture` in Terraform ([#358](https://github.com/databricks/databricks-sdk-go/pull/358)).

Dependency updates:
 
 * Bump golang.org/x/mod from 0.9.0 to 0.10.0 ([#356](https://github.com/databricks/databricks-sdk-go/pull/356)).
 * Bump google.golang.org/api from 0.114.0 to 0.115.0 ([#357](https://github.com/databricks/databricks-sdk-go/pull/357)).

## 0.6.0

 * Added type to represent a loaded configuration file ([#349](https://github.com/databricks/databricks-sdk-go/pull/349)).
 * Added named `Wait` the level of services ([#348](https://github.com/databricks/databricks-sdk-go/pull/348)).
 * Pass `azure_client_id` to Azure MSI if specified ([#354](https://github.com/databricks/databricks-sdk-go/pull/354)).

Dependency updates:
 
 * Bump google.golang.org/api from 0.112.0 to 0.114.0 ([#344](https://github.com/databricks/databricks-sdk-go/pull/344)).

## 0.5.0

 * Added Delta Live Tables events method to `pipelines.PipelinesAPI` ([#339](https://github.com/databricks/databricks-sdk-go/pull/339)).
 * Added `id` field to the `SparkVersionRequest` for easier integration with Terraform's [databricks_spark_version](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/spark_version) data source ([#335](https://github.com/databricks/databricks-sdk-go/pull/335)).
 * Improve auth fixture capturing ([#336](https://github.com/databricks/databricks-sdk-go/pull/336), [#337](https://github.com/databricks/databricks-sdk-go/pull/337)).
 * Make logger context aware and added examples of usage together with [`zerolog`](https://github.com/databricks/databricks-sdk-go/tree/main/examples/zerolog) and [`slog`](https://github.com/databricks/databricks-sdk-go/tree/main/examples/slog) ([#333](https://github.com/databricks/databricks-sdk-go/pull/333)).
 * Update from OpenAPI spec and added new APIs ([#338](https://github.com/databricks/databricks-sdk-go/pull/338)).
 
Dependency updates:
 
 * Bump google.golang.org/api from 0.111.0 to 0.112.0 ([#334](https://github.com/databricks/databricks-sdk-go/pull/334)).

## 0.4.1

 * Added an option to configure the location of the bricks CLI ([#330](https://github.com/databricks/databricks-sdk-go/pull/330)).
 * Added support for Azure CLI authentication on Azure China and Azure GovCloud ([#331](https://github.com/databricks/databricks-sdk-go/pull/331)).

## 0.4.0

 * Added `id` field to the `NodeTypeRequest` for easier integration with Terraform's [databricks_node_type](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/node_type) data source ([#325](https://github.com/databricks/databricks-sdk-go/pull/325)).
 * Don't load `~/.databrickscfg` if `azure-cli` auth is implicitly configured ([#324](https://github.com/databricks/databricks-sdk-go/pull/324)).
 * Fixed newline in codegen comments ([#326](https://github.com/databricks/databricks-sdk-go/pull/326)).
 * Update from OpenAPI ([#320](https://github.com/databricks/databricks-sdk-go/pull/320), [#328](https://github.com/databricks/databricks-sdk-go/pull/328)).

API changes:

* Renamed `deployment.AwsCredentials` to `deployment.CreateCredentialAwsCredentials`.
* Renamed `deployment.StsRole` to `deployment.CreateCredentialStsRole`.
* Removed schedules support from `sql.AlertsAPI`.

Dependency updates:

 * Bump github.com/stretchr/testify from 1.8.1 to 1.8.2 ([#318](https://github.com/databricks/databricks-sdk-go/pull/318)).
 * Bump golang.org/x/mod from 0.6.0-dev.0.20220419223038-86c51ed26bb4 to 0.8.0 ([#316](https://github.com/databricks/databricks-sdk-go/pull/316)).
 * Bump golang.org/x/mod from 0.8.0 to 0.9.0 ([#323](https://github.com/databricks/databricks-sdk-go/pull/323)).
 * Bump golang.org/x/oauth2 from 0.5.0 to 0.6.0 ([#322](https://github.com/databricks/databricks-sdk-go/pull/322)).
 * Bump golang.org/x/time from 0.0.0-20210723032227-1f47c861a9ac to 0.3.0 ([#317](https://github.com/databricks/databricks-sdk-go/pull/317)).
 * Bump google.golang.org/api from 0.110.0 to 0.111.0 ([#319](https://github.com/databricks/databricks-sdk-go/pull/319)).

## 0.3.3

 * Allow AAD SPN authentication on Databricks Account level ([#311](https://github.com/databricks/databricks-sdk-go/pull/311)).
 * Port auth tests from the JS SDK ([#313](https://github.com/databricks/databricks-sdk-go/pull/313)).
 * Skip loading `~/.databrickscfg` when not required ([#314](https://github.com/databricks/databricks-sdk-go/pull/314)).

Dependency updates:

 * Bump golang.org/x/net from 0.6.0 to 0.7.0 ([#312](https://github.com/databricks/databricks-sdk-go/pull/312)).

## 0.3.2

 * Always use the latest value for user agent key ([#309](https://github.com/databricks/databricks-sdk-go/pull/309)).

## 0.3.1

 * Change APIError to use pointer receiver ([#298](https://github.com/databricks/databricks-sdk-go/pull/298)).
 * Drop duplicate prefix in randomized email ([#299](https://github.com/databricks/databricks-sdk-go/pull/299)).
 * Expand semver pattern to be compliant with https://semver.org ([#302](https://github.com/databricks/databricks-sdk-go/pull/302)).
 * Fix `apierr.APIError` pointer receivers ([#307](https://github.com/databricks/databricks-sdk-go/pull/307)).
 * Skip loading config if auth is already explicitly configured directly ([#306](https://github.com/databricks/databricks-sdk-go/pull/306)).
 * Sync fixes for smallest node selection from Terraform ([#301](https://github.com/databricks/databricks-sdk-go/pull/301)).
 * Updated from OpenAPI spec ([#305](https://github.com/databricks/databricks-sdk-go/pull/305)).

Dependency updates:

 * Bump google.golang.org/api from 0.109.0 to 0.110.0 ([#303](https://github.com/databricks/databricks-sdk-go/pull/303)).

## 0.3.0

 * Added support for GCP deployment APIs in Public Preview.
 * Added new node type selector ([#287](https://github.com/databricks/databricks-sdk-go/pull/287)).
 * Added Azure MSI auth ([#279](https://github.com/databricks/databricks-sdk-go/pull/279)).
 * Added validation of strings in user agent package ([#291](https://github.com/databricks/databricks-sdk-go/pull/291)).
 * Added logging for error responses ([#289](https://github.com/databricks/databricks-sdk-go/pull/289)).
 * Update implementation of smallest node selection to match terraform's ([#273](https://github.com/databricks/databricks-sdk-go/pull/273)).
 * Added experimental OAuth support ([#276](https://github.com/databricks/databricks-sdk-go/pull/276), [#266](https://github.com/databricks/databricks-sdk-go/pull/266), [#292](https://github.com/databricks/databricks-sdk-go/pull/292)).
 * Fixed `google-credentials` to take precedence over `google-id` ([#283](https://github.com/databricks/databricks-sdk-go/pull/283)).
 * Test with Go 1.20 ([#295](https://github.com/databricks/databricks-sdk-go/pull/295)).
 * Regenerate from OpenAPI ([#293](https://github.com/databricks/databricks-sdk-go/pull/293), [#294](https://github.com/databricks/databricks-sdk-go/pull/294), [#296](https://github.com/databricks/databricks-sdk-go/pull/296), [#282](https://github.com/databricks/databricks-sdk-go/pull/282), [#269](https://github.com/databricks/databricks-sdk-go/pull/269), [#272](https://github.com/databricks/databricks-sdk-go/pull/272)).
 * Added `replaceAll` and `lowerFirst` template functions ([#288](https://github.com/databricks/databricks-sdk-go/pull/288), [#277](https://github.com/databricks/databricks-sdk-go/pull/277)).

Dependency updates:

 * Bump google.golang.org/api from 0.105.0 to 0.109.0 ([#271](https://github.com/databricks/databricks-sdk-go/pull/271), [#280](https://github.com/databricks/databricks-sdk-go/pull/280), [#284](https://github.com/databricks/databricks-sdk-go/pull/284), [#290](https://github.com/databricks/databricks-sdk-go/pull/290)).

API changes:

 * Renamed `clusters.CreateCluster` to `clusters.BaseClusterInfo`.
 * Renamed `jobs.Job` to `jobs.BaseJob`.
 * Renamed `jobs.Run` to `jobs.BaseRun`.
 * Fixed `commands.Results` `Schema` type from `[][]any` to `[]map[string]any`.

## 0.2.0

 * Added `DATABRICKS_AUTH_TYPE` environment variable ([#248](https://github.com/databricks/databricks-sdk-go/pull/248)).
 * Added Policy Families API ([#263](https://github.com/databricks/databricks-sdk-go/pull/263)).
 * Added experimental `ErrCannotConfigureAuth` and `ErrNotAccountClient` ([#237](https://github.com/databricks/databricks-sdk-go/pull/237), [#238](https://github.com/databricks/databricks-sdk-go/pull/238)).
 * Added DBFS file handle that supports both reading and writing ([#261](https://github.com/databricks/databricks-sdk-go/pull/261)).
 * Added `io.WriterTo` for DBFS file reader ([#249](https://github.com/databricks/databricks-sdk-go/pull/249)).
 * Added `pflag.Value` interfaces for enums ([#234](https://github.com/databricks/databricks-sdk-go/pull/234)).
 * Added support for adding custom HTTP visitors per request ([#230](https://github.com/databricks/databricks-sdk-go/pull/230)).
 * Added support for raw body as byte slice if requested ([#247](https://github.com/databricks/databricks-sdk-go/pull/247)).
 * Improved callbacks for polling the status of long-running operations ([#258](https://github.com/databricks/databricks-sdk-go/pull/258)).
 * Improved rendering of HTTP links in godoc ([#229](https://github.com/databricks/databricks-sdk-go/pull/229)).
 * Updated field types in the Jobs API from spec ([#259](https://github.com/databricks/databricks-sdk-go/pull/259)).
 * Multiple OpenAPI consistency passes ([#254](https://github.com/databricks/databricks-sdk-go/pull/254), [#241](https://github.com/databricks/databricks-sdk-go/pull/241), [#243](https://github.com/databricks/databricks-sdk-go/pull/243), [#255](https://github.com/databricks/databricks-sdk-go/pull/255), [#236](https://github.com/databricks/databricks-sdk-go/pull/236)).

API changes:
 * Renamed `IsAccountsClient` to `IsAccountClient` ([#231](https://github.com/databricks/databricks-sdk-go/pull/231)).
 * `w.ClusterPolicies.ListAll` now takes `clusterpolicies.List` as an argument.
 * `github.com/databricks/databricks-sdk-go/service/dbsql` package is renamed to `github.com/databricks/databricks-sdk-go/service/sql`.
 * `w.DataSources.ListDataSources` is renamed to `w.DataSources.List`.
 * `w.Queries.CreateQuery` is renamed to `w.Queries.CreateQuery`.
 * `w.Queries.DeleteQueryByQueryId` is renamed to `w.Queries.DeleteByQueryId`.
 * `w.Queries.GetQueryByQueryId` is renamed to `w.Queries.GetByQueryId`.
 * `w.Queries.UpdateQuery` is renamed to `w.Queries.Update`.
 * `w.Alerts.DeleteAlertByAlertId` is renamed to `w.Alerts.DeleteByAlertId`.
 * `w.Alerts.UpdateAlert` is renamed to `w.Alerts.Update`.
 * `w.Alerts.GetAlertByAlertId` is renamed to `w.Alerts.GetByAlertId`.
 * `w.Alerts.ListAlerts` is renamed to `w.Alerts.List`.
 * `w.Dashboards.CreateDashboard` is renamed to `w.Dashboards.Create`.
 * `w.Dashboards.DeleteDashboardByDashboardId` is renamed to `w.Dashboards.GetByDashboardId`.
 * `w.Dashboards.ListDashboardsAll` is renamed to `w.Dashboards.ListAll`.
 * `w.Dashboards.DeleteDashboardByDashboardId` is renamed to `w.Dashboards.DeleteByDashboardId`.
 * `w.Dashboards.RestoreDashboard` is renamed to `w.Dashboards.Restore`.
 * `deployment.CreateCustomerManagedKeyRequest` now takes `deployment.KeyUseCase` enum.
 * `w.GlobalInitScripts.CreateScript` is renamed to `w.GlobalInitScripts.Create`.
 * `w.GlobalInitScripts.DeleteScriptByScriptId` is renamed to `w.GlobalInitScripts.DeleteByScriptId`.
 * `w.GlobalInitScripts.UpdateScript` is renamed to `w.GlobalInitScripts.Update`.
 * `w.GlobalInitScripts.GetScriptByScriptId` is renamed to `w.GlobalInitScripts.GetByScriptId`.
 * `w.GlobalInitScripts.ListScriptsAll` is renamed to `w.GlobalInitScripts.ListAll`.
 * `jobs.ResetJob.NewSettings` is now required field.
 * `w.Pipelines.CreatePipeline` is renamed to `w.Pipelines.Create`.
 * `w.Pipelines.DeletePipelineByPipelineId` is renamed to `w.Pipelines.DeleteByPipelineId`.
 * `w.Pipelines.UpdatePipeline` is renamed to `w.Pipelines.Update`.
 * `w.Pipelines.GetPipelineByPipelineId` is renamed to `w.Pipelines.GetByPipelineId`.
 * `w.StorageCredentials.Update` now also returns an entity.
 * `w.ExternalLocations.Update` now also returns an entity.
 * `w.Metastores.Update` now also returns an entity.
 * `unitycatalog.CreateMetastoreAssignment.WorkspaceId` type changed from `int` to `int64`.
 * `unitycatalog.UnassignRequest.WorkspaceId` type changed from `int` to `int64`.
 * `w.Catalogs.Update` now also returns an entity.
 * `w.Schemas.Update` now also returns an entity.
 * `w.Providers.Update` now also returns an entity.
 * `w.Shares.Update` now also returns an entity.
 * `WarehousesAPI` service moved to `github.com/databricks/databricks-sdk-go/service/sql` package.
 * `w.Warehouses.CreateWarehouseAndWait` renamed to `w.Warehouses.CreateAndWait`.
 * `w.Warehouses.DeleteWarehouseByIdAndWait` renamed to `w.Warehouses.DeleteByIdAndWait`.
 * `w.Warehouses.EditWarehouse` renamed to `w.Warehouses.Edit`.
 * `w.Warehouses.GetWarehouseById` renamed to `w.Warehouses.GetById`.
 * `w.Warehouses.ListWarehousesAll` renamed to `w.Warehouses.ListAll`.
 * Removed `w.Dbfs.Overwrite` in favor of `w.Dbfs.Open("....", dbfs.FileModeOverwrite|dbfs.FileModeWrite)`.
 * Added third required argument to `w.Dbfs.Open`.

Code generation:

 * Added concept of `main` service for the package ([#239](https://github.com/databricks/databricks-sdk-go/pull/239)).
 * Added entity primitives check ([#242](https://github.com/databricks/databricks-sdk-go/pull/242)).
 * Added helpers for CRUD generation ([#246](https://github.com/databricks/databricks-sdk-go/pull/246)).
 * Added more entity-generation utils ([#257](https://github.com/databricks/databricks-sdk-go/pull/257)).
 * Dynamically generate `.gitattributes` ([#244](https://github.com/databricks/databricks-sdk-go/pull/244)).
 * Fixed required order fields ([#245](https://github.com/databricks/databricks-sdk-go/pull/245)).
 * Parse summary from descriptions ([#228](https://github.com/databricks/databricks-sdk-go/pull/228)).
 * Print error on formatter failure ([#235](https://github.com/databricks/databricks-sdk-go/pull/235)).
 * Update usage string in generator ([#260](https://github.com/databricks/databricks-sdk-go/pull/260)).
 * Fixed order of host completion ([#233](https://github.com/databricks/databricks-sdk-go/pull/233)).

Dependency updates:

 * Bump google.golang.org/api from 0.103.0 to 0.105.0 ([#232](https://github.com/databricks/databricks-sdk-go/pull/232), [#252](https://github.com/databricks/databricks-sdk-go/pull/252)).
 
## 0.1.1

* Improved usage instructions and added more documentation for services.

## 0.1.0

 * Initial release of the Databricks SDK for Go.
