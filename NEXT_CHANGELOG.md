# NEXT CHANGELOG

## Release v0.136.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `UndeleteBranch` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `Attributes` and `ExcludedAttributes` fields for [iam.MeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#MeRequest).
* Add `IncludeTriggerState` field for [jobs.GetJobRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetJobRequest).
* Add `DeleteTime` and `PurgeTime` fields for [postgres.BranchStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchStatus).
* Add `Purge` field for [postgres.DeleteBranchRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#DeleteBranchRequest).
* Add `ShowDeleted` field for [postgres.ListBranchesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ListBranchesRequest).
* Add `Deleted` enum value for [postgres.BranchStatusState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchStatusState).
* [Breaking] Change `ActionType` and `ResourceId` fields for [bundle.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle#Operation) to be required.
* Change `ActionType` and `ResourceId` fields for [bundle.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle#Operation) to be required.
* [Breaking] Change `CliVersion` field for [bundle.Version](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle#Version) to be required.
* Change `CliVersion` field for [bundle.Version](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle#Version) to be required.
* [Breaking] Change `Tags` field for [marketplace.ListListingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ListListingsRequest) to type [marketplace.ListingTag](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ListingTag).
* [Breaking] Change pagination for [ClustersAPI.Events](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClustersAPI.Events).