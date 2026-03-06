---
description: "Error handling conventions for the Databricks SDK."
globs: "**/*.go"
alwaysApply: false
---

# Error Handling

- Wrap errors with **caller's context**: `fmt.Errorf("resolving workspace host: %w", err)`.
  Describe what *you* were doing, not what the callee does
- Don't add redundant wrapping — if the underlying error already contains the path/resource,
  add context at the abstraction boundary instead
- Include the specific resource/URL that failed, not just the operation name
- Define sentinel errors as package-level `var` using `errors.New()` — never `fmt.Errorf`.
  They must be comparable with `errors.Is()`
- Use `%q` (not `%s`) for strings in error/log messages — makes empty strings visible as `""`
