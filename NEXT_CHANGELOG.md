# NEXT CHANGELOG

## Release v0.88.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `CreateMaterializedFeature`, `DeleteMaterializedFeature`, `GetMaterializedFeature`, `ListMaterializedFeatures` and `UpdateMaterializedFeature` methods for [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service.
* Add `FilterCondition` field for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature).
* [Breaking] Change `DisplayName`, `Evaluation`, `QueryText`, `Schedule` and `WarehouseId` fields for [sql.AlertV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2) to be required.
* Change `DisplayName`, `Evaluation`, `QueryText`, `Schedule` and `WarehouseId` fields for [sql.AlertV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2) to be required.
* [Breaking] Change `ComparisonOperator` and `Source` fields for [sql.AlertV2Evaluation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2Evaluation) to be required.
* Change `ComparisonOperator` and `Source` fields for [sql.AlertV2Evaluation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2Evaluation) to be required.
* [Breaking] Change `Name` field for [sql.AlertV2OperandColumn](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2OperandColumn) to be required.
* Change `Name` field for [sql.AlertV2OperandColumn](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2OperandColumn) to be required.
* [Breaking] Change `QuartzCronSchedule` and `TimezoneId` fields for [sql.CronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CronSchedule) to be required.
* Change `QuartzCronSchedule` and `TimezoneId` fields for [sql.CronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CronSchedule) to be required.
* [Breaking] Remove `Results` field for [sql.ListAlertsV2Response](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListAlertsV2Response).