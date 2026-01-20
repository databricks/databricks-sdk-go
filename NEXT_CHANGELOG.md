# NEXT CHANGELOG

## Release v0.99.0

### Breaking Changes

* **[workspace]** Updated Git Credentials List methods to require `ListCredentialsRequest` parameter. The following methods now require a `workspace.ListCredentialsRequest` parameter instead of no parameters:
  - `GitCredentialsAPI.List(ctx, request)`
  - `GitCredentialsAPI.ListAll(ctx, request)`
  - `GitCredentialsAPI.CredentialInfoGitProviderToCredentialIdMap(ctx, request)`

  **Migration Guide:**
  ```go
  // Before
  credentials, err := w.GitCredentials.ListAll(ctx)

  // After
  credentials, err := w.GitCredentials.ListAll(ctx, workspace.ListCredentialsRequest{})
  ```

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `PrincipalId` field for [workspace.CreateCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CreateCredentialsRequest).
* Add `PrincipalId` field for [workspace.DeleteCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#DeleteCredentialsRequest).
* Add `PrincipalId` field for [workspace.GetCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetCredentialsRequest).
* Add `PrincipalId` field for [workspace.ListCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ListCredentialsRequest).
* Add `PrincipalId` field for [workspace.UpdateCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#UpdateCredentialsRequest).
