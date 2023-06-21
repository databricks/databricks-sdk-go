# Version changelog

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
