# NEXT CHANGELOG

## Release v0.160.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `ListCleanRoomTaskRunsHandler` method for [w.CleanRoomTaskRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomTaskRunsAPI) workspace-level service.
* Add `InitialParentPath` field for [bundledeployments.Deployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Deployment).
* Add `JarAnalysis` field for [cleanrooms.CleanRoomAsset](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAsset).
* Add `JarAnalysisReview` field for [cleanrooms.CreateCleanRoomAssetReviewRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CreateCleanRoomAssetReviewRequest).
* Add `JarAnalysisReviewState` and `JarAnalysisReviews` fields for [cleanrooms.CreateCleanRoomAssetReviewResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CreateCleanRoomAssetReviewResponse).
* Add `CapacityReservationGroup` field for [compute.InstancePoolAzureAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAzureAttributes).
* Add `Lifetime` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
* Add `Jdbc` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `JarAnalysis` enum value for [cleanrooms.CleanRoomAssetAssetType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetAssetType).