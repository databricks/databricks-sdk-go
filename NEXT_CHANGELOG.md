# NEXT CHANGELOG

## Release v0.59.0

### New Features and Improvements

### Bug Fixes

* Fix unlikely issue due to conflicting error details in `APIError`.

### Documentation

### Internal Changes

* Update Jobs ListRuns API to support paginated responses ([#1151](https://github.com/databricks/databricks-sdk-go/pull/1151))
* Add `poll.SimpleError` to mock waiter objects returning errors  ([#1155](https://github.com/databricks/databricks-sdk-go/pull/1155))
* Refactor `APIError` to expose different types of error details ([#1153](https://github.com/databricks/databricks-sdk-go/pull/1153)). 
* Update Jobs ListJobs API to support paginated responses ([#1150](https://github.com/databricks/databricks-sdk-go/pull/1150))
* Introduce automated tagging ([#1148](https://github.com/databricks/databricks-sdk-go/pull/1148)).
* Update Jobs GetJob API to support paginated responses  ([#1133](https://github.com/databricks/databricks-sdk-go/pull/1133)).
* Update Jobs GetRun API to support paginated responses  ([#1132](https://github.com/databricks/databricks-sdk-go/pull/1132)).

### API Changes
* Added `GetSpace` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Added `ListProviderShareAssets` method for [w.Providers](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ProvidersAPI) workspace-level service.
* Added `BudgetPolicyId` and `EffectiveBudgetPolicyId` fields for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Added `Policy` field for [billing.CreateBudgetPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBudgetPolicyRequest).
* Added `DatabricksGcpServiceAccount` field for [catalog.ValidateCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidateCredentialRequest).
* Added `AttachmentId` field for [dashboards.GenieAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAttachment).
* Added `ConversationId` field for [dashboards.GenieConversation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieConversation).
* Added `MessageId` field for [dashboards.GenieMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieMessage).
* Added `Description`, `Id`, `LastUpdatedTimestamp`, `Query`, `QueryResultMetadata` and `Title` fields for [dashboards.GenieQueryAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieQueryAttachment).
* Added `GenAiComputeTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Added `GenAiComputeTask` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Added `GenAiComputeTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Added `RunName` field for [ml.CreateRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CreateRun).
* Added `RunName` field for [ml.RunInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RunInfo).
* Added `RunName` field for [ml.UpdateRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#UpdateRun).
* Added `Lifetime` field for [oauth2.CreateServicePrincipalSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateServicePrincipalSecretRequest).
* Added `ExpireTime` field for [oauth2.CreateServicePrincipalSecretResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateServicePrincipalSecretResponse).
* Added `ExpireTime` field for [oauth2.SecretInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#SecretInfo).
* Added `InstanceProfileArn` field for [serving.AmazonBedrockConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AmazonBedrockConfig).
* Added `Add`, `Principal` and `Remove` fields for [sharing.PermissionsChange](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#PermissionsChange).
* Added `ColumnsToRerank` field for [vectorsearch.QueryVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexRequest).
* Added `Oracle` and `Teradata` enum values for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Added `FunctionArgumentsInvalidTypeException` and `MessageCancelledWhileExecutingException` enum values for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Added `Waiting` enum value for [jobs.RunLifecycleStateV2State](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunLifecycleStateV2State).
* Added `ActiveOnly`, `All` and `DeletedOnly` enum values for [ml.ViewType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ViewType).
* Added `OauthClientCredentials` enum value for [sharing.AuthenticationType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#AuthenticationType).
* Added `Raw` enum value for [workspace.ExportFormat](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ExportFormat).
* [Breaking] Changed `GetByName` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service to return [ml.GetExperimentByNameResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentByNameResponse).
* [Breaking] Changed `LogInputs` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service with new required argument order.
* [Breaking] Changed `SharePermissions` and `UpdatePermissions` methods for [w.Shares](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharesAPI) workspace-level service return type to become non-empty.
* [Breaking] Changed `SharePermissions` method for [w.Shares](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharesAPI) workspace-level service to return [sharing.GetSharePermissionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#GetSharePermissionsResponse).
* [Breaking] Changed `UpdatePermissions` method for [w.Shares](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharesAPI) workspace-level service to return [sharing.UpdateSharePermissionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateSharePermissionsResponse).
* Changed `PolicyId` field for [billing.BudgetPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetPolicy) to no longer be required.
* [Breaking] Changed `PolicyId` field for [billing.BudgetPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetPolicy) to no longer be required.
* [Breaking] Changed `Partitions` field for [cleanrooms.CleanRoomAssetTableLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetTableLocalDetails) to type [cleanrooms.PartitionList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#PartitionList).
* [Breaking] Changed `Query` field for [dashboards.GenieAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAttachment) to type [dashboards.GenieQueryAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieQueryAttachment).
* Changed `Digest`, `Name`, `Source` and `SourceType` fields for [ml.Dataset](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Dataset) to be required.
* [Breaking] Changed `Digest`, `Name`, `Source` and `SourceType` fields for [ml.Dataset](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Dataset) to be required.
* [Breaking] Changed `Dataset` field for [ml.DatasetInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DatasetInput) to be required.
* Changed `Dataset` field for [ml.DatasetInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DatasetInput) to be required.
* Changed `Key` and `Value` fields for [ml.InputTag](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#InputTag) to be required.
* [Breaking] Changed `Key` and `Value` fields for [ml.InputTag](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#InputTag) to be required.
* [Breaking] Changed `ViewType` field for [ml.ListExperimentsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ListExperimentsRequest) to type [ml.ViewType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ViewType).
* [Breaking] Changed `RunId` field for [ml.LogInputs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#LogInputs) to be required.
* [Breaking] Changed `ViewType` field for [ml.SearchExperiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SearchExperiments) to type [ml.ViewType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ViewType).
* [Breaking] Changed `RunViewType` field for [ml.SearchRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SearchRuns) to type [ml.ViewType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ViewType).
* [Breaking] Removed `CustomTags` and `PolicyName` fields for [billing.CreateBudgetPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBudgetPolicyRequest).
* [Breaking] Removed `CachedQuerySchema`, `Description`, `Id`, `InstructionId`, `InstructionTitle`, `LastUpdatedTimestamp`, `Query`, `StatementId` and `Title` fields for [dashboards.QueryAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#QueryAttachment).
* [Breaking] Removed `MaxResults` and `PageToken` fields for [sharing.UpdateSharePermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateSharePermissions).
* [Breaking] Removed `ActiveOnly`, `All` and `DeletedOnly` enum values for [ml.SearchExperimentsViewType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SearchExperimentsViewType).
* [Breaking] Removed `ActiveOnly`, `All` and `DeletedOnly` enum values for [ml.SearchRunsRunViewType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SearchRunsRunViewType).
