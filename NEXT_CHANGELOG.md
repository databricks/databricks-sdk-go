# NEXT CHANGELOG

## Release v0.120.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `UpdateRole` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `Entities` and `TimeseriesColumn` fields for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature).
* Add `AggregationFunction` field for [ml.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Function).
* Add `FilterCondition` field for [ml.KafkaSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#KafkaSource).
* Add `Attributes` field for [postgres.RoleRoleStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#RoleRoleStatus).
* [Breaking] Change `EntityColumns` and `TimeseriesColumn` fields for [ml.DeltaTableSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeltaTableSource) to no longer be required.
* Change `EntityColumns` and `TimeseriesColumn` fields for [ml.DeltaTableSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeltaTableSource) to no longer be required.
* [Breaking] Change `Inputs` field for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature) to no longer be required.
* Change `Inputs` field for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature) to no longer be required.
* [Breaking] Change `FunctionType` field for [ml.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Function) to no longer be required.
* Change `FunctionType` field for [ml.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Function) to no longer be required.
* Change `EntityColumnIdentifiers` and `TimeseriesColumnIdentifier` fields for [ml.KafkaSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#KafkaSource) to no longer be required.
* [Breaking] Change `EntityColumnIdentifiers` and `TimeseriesColumnIdentifier` fields for [ml.KafkaSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#KafkaSource) to no longer be required.