# NEXT CHANGELOG

## Release v0.122.0

### Breaking Changes

### New Features and Improvements

* Add automatic detection of AI coding agents (Claude Code, Cursor, Cline, Codex, Copilot CLI, Gemini CLI, OpenCode, Antigravity, Openclaw) in the user-agent string. The SDK now appends `agent/<name>` to HTTP request headers when running inside a known AI agent environment.

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `AlertOutput` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
* Add `AlertTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Add `AlertTask` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Add `AlertTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Add [environments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/environments) package.
* Add [w.Environments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/environments#EnvironmentsAPI) workspace-level service.
* Add `CanCreateApp` enum value for [iam.PermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionLevel).
