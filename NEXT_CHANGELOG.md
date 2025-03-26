# NEXT CHANGELOG

## Release v0.61.0

### New Features and Improvements

* Support user-to-machine authentication in the SDK ([#1108](https://github.com/databricks/databricks-sdk-go/pull/1108)).
- Instances of `ApiClient` now share the same connection pool by default ([PR #1190](https://github.com/databricks/databricks-sdk-go/pull/1190)).

### Bug Fixes

### Documentation

### Internal Changes

- Stop recommending users to report an issue when the SDK encounters an unknown
  error ([PR #1189](https://github.com/databricks/databricks-sdk-go/pull/1189)).

### API Changes
* Added `GenerateDownloadFullQueryResult` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Added `EffectiveUserApiScopes`, `Oauth2AppClientId`, `Oauth2AppIntegrationId` and `UserApiScopes` fields for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Added `Abfss`, `Dbfs`, `ErrorMessage`, `ExecutionDurationSeconds`, `File`, `Gcs`, `S3`, `Status`, `Volumes` and `Workspace` fields for [compute.InitScriptInfoAndExecutionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptInfoAndExecutionDetails).
* [Breaking] Added `ForecastGranularity` field for [ml.CreateForecastingExperimentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CreateForecastingExperimentRequest).
* Added `JwksUri` field for [oauth2.OidcFederationPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#OidcFederationPolicy).
* Added `EventLog` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Added `EventLog` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Added `EventLog` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
* Added `FallbackConfig` field for [serving.AiGatewayConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayConfig).
* Added `CustomProviderConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
* Added `FallbackConfig` field for [serving.PutAiGatewayRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PutAiGatewayRequest).
* Added `FallbackConfig` field for [serving.PutAiGatewayResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PutAiGatewayResponse).
* Added `Aliases`, `Comment`, `DataType`, `DependencyList`, `FullDataType`, `Id`, `InputParams`, `Name`, `Properties`, `RoutineDefinition`, `Schema`, `SecurableKind`, `Share`, `ShareId`, `StorageLocation` and `Tags` fields for [sharing.DeltaSharingFunction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#DeltaSharingFunction).
* Added `QuerySource` field for [sql.QueryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryInfo).
* Added `ForeignCatalog` enum value for [catalog.CatalogType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogType).
* Added `Browse` enum value for [catalog.Privilege](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#Privilege).
* Added `AccessTokenFailure`, `AllocationTimeout`, `AllocationTimeoutNodeDaemonNotReady`, `AllocationTimeoutNoHealthyClusters`, `AllocationTimeoutNoMatchedClusters`, `AllocationTimeoutNoReadyClusters`, `AllocationTimeoutNoUnallocatedClusters`, `AllocationTimeoutNoWarmedUpClusters`, `AwsInaccessibleKmsKeyFailure`, `AwsInstanceProfileUpdateFailure`, `AwsInvalidKeyPair`, `AwsInvalidKmsKeyState`, `AwsResourceQuotaExceeded`, `AzurePackedDeploymentPartialFailure`, `BootstrapTimeoutDueToMisconfig`, `BudgetPolicyLimitEnforcementActivated`, `BudgetPolicyResolutionFailure`, `CloudAccountSetupFailure`, `CloudOperationCancelled`, `CloudProviderInstanceNotLaunched`, `CloudProviderLaunchFailureDueToMisconfig`, `CloudProviderResourceStockoutDueToMisconfig`, `ClusterOperationThrottled`, `ClusterOperationTimeout`, `ControlPlaneRequestFailureDueToMisconfig`, `DataAccessConfigChanged`, `DisasterRecoveryReplication`, `DriverEviction`, `DriverLaunchTimeout`, `DriverNodeUnreachable`, `DriverOutOfDisk`, `DriverOutOfMemory`, `DriverPodCreationFailure`, `DriverUnexpectedFailure`, `DynamicSparkConfSizeExceeded`, `EosSparkImage`, `ExecutorPodUnscheduled`, `GcpApiRateQuotaExceeded`, `GcpForbidden`, `GcpIamTimeout`, `GcpInaccessibleKmsKeyFailure`, `GcpInsufficientCapacity`, `GcpIpSpaceExhausted`, `GcpKmsKeyPermissionDenied`, `GcpNotFound`, `GcpResourceQuotaExceeded`, `GcpServiceAccountAccessDenied`, `GcpServiceAccountNotFound`, `GcpSubnetNotReady`, `GcpTrustedImageProjectsViolated`, `GkeBasedClusterTermination`, `InitContainerNotFinished`, `InstancePoolMaxCapacityReached`, `InstancePoolNotFound`, `InstanceUnreachableDueToMisconfig`, `InternalCapacityFailure`, `InvalidAwsParameter`, `InvalidInstancePlacementProtocol`, `InvalidWorkerImageFailure`, `InPenaltyBox`, `LazyAllocationTimeout`, `MaintenanceMode`, `NetvisorSetupTimeout`, `NoMatchedK8s`, `NoMatchedK8sTestingTag`, `PodAssignmentFailure`, `PodSchedulingFailure`, `ResourceUsageBlocked`, `SecretCreationFailure`, `ServerlessLongRunningTerminated`, `SparkImageDownloadThrottled`, `SparkImageNotFound`, `SshBootstrapFailure`, `StorageDownloadFailureDueToMisconfig`, `StorageDownloadFailureSlow`, `StorageDownloadFailureThrottled`, `UnexpectedPodRecreation`, `UserInitiatedVmTermination` and `WorkspaceUpdate` enum values for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Added `GeneratedSqlQueryTooLongException` and `MissingSqlQueryException` enum values for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Added `Balanced` enum value for [jobs.PerformanceTarget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PerformanceTarget).
* Added `ListingResource` enum value for [marketplace.FileParentType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#FileParentType).
* Added `App` enum value for [marketplace.MarketplaceFileType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#MarketplaceFileType).
* Added `Custom` enum value for [serving.ExternalModelProvider](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModelProvider).
* Added `ArclightMultiTenantAzureExchangeToken` and `ArclightMultiTenantAzureExchangeTokenWithUserDelegationKey` enum values for [settings.TokenType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenType).
* [Breaking] Changed `CreateExperiment` method for [w.Forecasting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ForecastingAPI) workspace-level service with new required argument order.
* Changed `InstanceTypeId` field for [compute.NodeInstanceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#NodeInstanceType) to be required.
* Changed `Category` field for [compute.NodeType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#NodeType) to be required.
* [Breaking] Changed `Functions` field for [sharing.ListProviderShareAssetsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListProviderShareAssetsResponse) to type [sharing.DeltaSharingFunctionList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#DeltaSharingFunctionList).
* [Breaking] Removed `ExecutionDetails` and `Script` fields for [compute.InitScriptInfoAndExecutionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptInfoAndExecutionDetails).
* [Breaking] Removed `SupportsElasticDisk` field for [compute.NodeType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#NodeType).
* [Breaking] Removed `DataGranularityQuantity` and `DataGranularityUnit` fields for [ml.CreateForecastingExperimentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CreateForecastingExperimentRequest).
* [Breaking] Removed `Aliases`, `Comment`, `DataType`, `DependencyList`, `FullDataType`, `Id`, `InputParams`, `Name`, `Properties`, `RoutineDefinition`, `Schema`, `SecurableKind`, `Share`, `ShareId`, `StorageLocation` and `Tags` fields for [sharing.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#Function).
