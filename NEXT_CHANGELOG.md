# NEXT CHANGELOG

## Release v0.82.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added `Comment` field for [dashboards.GenieSendMessageFeedbackRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSendMessageFeedbackRequest).
* [Breaking] Added `Rating` field for [dashboards.GenieSendMessageFeedbackRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSendMessageFeedbackRequest).
* Added `EffectiveEnablePgNativeLogin` and `EnablePgNativeLogin` fields for [database.DatabaseInstance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstance).
* Added `TaskRetryMode` field for [jobs.Continuous](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Continuous).
* Added `SourceConfigurations` field for [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
* Added `AppId`, `AppIdSet`, `AuthSecret`, `AuthSecretSet`, `ChannelUrl`, `ChannelUrlSet`, `TenantId` and `TenantIdSet` fields for [settings.MicrosoftTeamsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#MicrosoftTeamsConfig).
* Added `EnsureRerankerCompatible` field for [vectorsearch.GetIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#GetIndexRequest).
* Added `Reranker` field for [vectorsearch.QueryVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexRequest).
* [Breaking] Changed `CreateCleanRoomAssetReview` method for [w.CleanRoomAssets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetsAPI) workspace-level service with new required argument order.
* [Breaking] Changed `SendMessageFeedback` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service with new required argument order.
* Changed `NotebookReview` field for [cleanrooms.CreateCleanRoomAssetReviewRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CreateCleanRoomAssetReviewRequest) to no longer be required.
* [Breaking] Changed `Features` field for [ml.FeatureList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureList) to type [][ml.LinkedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#LinkedFeature).
* [Breaking] Removed `FeedbackRating` and `FeedbackText` fields for [dashboards.GenieSendMessageFeedbackRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSendMessageFeedbackRequest).