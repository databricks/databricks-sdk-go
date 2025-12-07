# NEXT CHANGELOG

## Release v0.94.0

### New Features and Improvements

### Bug Fixes

### documentation

### Internal Changes

### API Changes
* Add `Clone` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
* Add `CronSchedule` field for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* Add `GcpServiceAccount` field for [provisioning.CreateGcpKeyInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateGcpKeyInfo).
* Add `GcpServiceAccount` field for [provisioning.GcpKeyInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#GcpKeyInfo).
* Change `TimeWindow` field for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature) to no longer be required.
* [Breaking] Change `TimeWindow` field for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature) to no longer be required.