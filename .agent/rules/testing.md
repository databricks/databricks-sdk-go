---
description: "Testing standards and anti-patterns for the Databricks SDK."
globs: "**/*_test.go"
alwaysApply: false
---

# Testing

`testify` is **discouraged in new packages** and **tolerated in existing packages**.
The SDK aims to minimize third-party dependencies.

- **New packages**: use Go's standard `testing` package with `google/go-cmp` for complex comparisons
- **Existing packages that already use testify**: keep using it for consistency within that package
- Do NOT introduce testify into a package that doesn't already use it

Table-driven tests are preferred for API client methods:

```go
for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
        // ...
        if diff := cmp.Diff(tc.expected, actual); diff != "" {
            t.Errorf("mismatch (-want +got):\n%s", diff)
        }
    })
}
```

- Unit tests: `*_test.go` alongside source files
- Integration tests: use `//go:build cloud` build tag
- HTTP fixtures: `httpclient/fixtures/` for mocking HTTP responses
- Generated mocks: `experimental/mocks/` (do not hand-write mocks for generated services)
- Test helpers: `qa/` package for integration test utilities
- Call `t.Helper()` at the start of test helper functions (correct line numbers in failures)
- For configurable options, always test three cases: (1) explicit values, (2) nil/unset default, (3) empty value
- When a feature spans multiple implementations, test the same scenarios consistently in each
- Inject `time.Now` as `func() time.Time` for time-dependent logic — makes tests deterministic
- Always check errors — never use `_` for error returns in tests
- Assert on exact expected values, not partial matches (e.g., full header value, not just prefix)
- Extract repeated test values to shared `const`/`var` to convey intent and avoid silent drift
- Add explicit timeouts to channel/goroutine-based tests to prevent hanging
- Use `defer close(ch)` for channel cleanup to handle error paths safely

## Anti-Patterns (never do these)

- Deleting or skipping flaky tests instead of fixing the root cause
- Over-mocking — prefer HTTP fixtures over deep mock chains
- Tests depending on execution order or shared mutable state
- Non-descriptive test names (`TestFoo`, `TestBasic`) — use `TestFunctionName_Scenario_ExpectedBehavior`
- `time.Sleep` in tests — inject a clock or use polling with `assert.Eventually`
- Testing generated code in `service/` — test your hand-written code that uses it
- Exporting methods solely for test access — use `cmp.Diff` to compare structs,
  or write internal (`_test.go` in same package) tests instead
- Unnecessary concurrency in tests — if goroutines/channels aren't clearly needed, don't use them
