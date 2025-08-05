# NEXT CHANGELOG

## Release v0.79.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added `StatementIdSignature` field for [dashboards.Result](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Result).
* Added `EffectiveDatabaseInstanceName` and `EffectiveLogicalDatabaseName` fields for [database.SyncedDatabaseTable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedDatabaseTable).
* Added `Table` field for [jobs.TriggerStateProto](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerStateProto).
* Added `EmailNotifications` field for [serving.CreatePtEndpointRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreatePtEndpointRequest).
* Added `EmailNotifications` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint).
* Added `EmailNotifications` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).
* [Breaking] Changed `List` method for [w.ConsumerProviders](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ConsumerProvidersAPI) workspace-level service . New request type is [marketplace.ListConsumerProvidersRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ListConsumerProvidersRequest).
* [Breaking] Changed `Create` method for [a.PrivateAccess](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#PrivateAccessAPI) account-level service . New request type is [provisioning.CreatePrivateAccessSettingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreatePrivateAccessSettingsRequest).
* [Breaking] Changed `Create` method for [a.PrivateAccess](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#PrivateAccessAPI) account-level service with new required argument order.
* [Breaking] Changed `Replace` method for [a.PrivateAccess](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#PrivateAccessAPI) account-level service . New request type is [provisioning.ReplacePrivateAccessSettingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#ReplacePrivateAccessSettingsRequest).
* [Breaking] Removed `IsFeatured` field for [marketplace.ListProvidersRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ListProvidersRequest).