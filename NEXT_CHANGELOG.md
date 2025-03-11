# NEXT CHANGELOG

## Release v0.60.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added [w.Forecasting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ForecastingAPI) workspace-level service.
* Added `ExecuteMessageAttachmentQuery` and `GetMessageAttachmentQueryResult` methods for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Added `StatementId` field for [dashboards.GenieQueryAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieQueryAttachment).
* Added `BudgetPolicyId` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint).
* Added `BudgetPolicyId` field for [serving.ServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpoint).
* Added `BudgetPolicyId` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).
* Added `CouldNotGetModelDeploymentsException` enum value for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
