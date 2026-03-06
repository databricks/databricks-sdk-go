---
description: "Go code style and naming conventions for the Databricks SDK."
globs: "**/*.go"
alwaysApply: false
---

# Code Style

- Go 1.22+ idioms: range-over-int, `min`/`max` builtins
- Format with `goimports` (handles import grouping) then `gofmt`
- All exported functions and types need doc comments: `// FuncName does X.`
- Interface doc comments should explain the **contract/lifecycle** — when methods are called,
  what callers expect, and when errors are appropriate — not just what they do
- Return early on errors; avoid deeply nested if-else chains
- Use `context.Context` as the first parameter for functions that do I/O

## Naming

- Acronyms are fully capitalized: `ID` not `Id`, `URL` not `Url`, `HTTP` not `Http`
- Avoid abbreviations: `authTokenCommand` not `authTokenCmd`
- Variable names describe content, not origin: `workspaceHost` not `resolvedHost`
- Use specific names when multiple similar types exist: `httpClient` not `client`
