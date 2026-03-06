---
description: "API design principles for the Databricks SDK."
globs: "**/*.go"
alwaysApply: false
---

# API Design

- Minimize exported API surface: if a type is only used through an interface, keep
  the concrete type unexported
- Pass specific values, not entire config objects — functions should depend on the narrowest
  interface possible
- Prefer new types over modifying existing exported types (additive changes only)
- Unstable APIs go under `experimental/` with `// Experimental: subject to change.` doc comments
- When setting default values (timeouts, retries, cache durations), comment the rationale
- Place helper functions near their callers, not at the top of the file
- Validate required fields in constructors — validate early, fail fast
- Validation logic should live near its point of use, not in generic/shared layers
- Use `switch/case` for type dispatching, not `if/else` chains
- Split functions with multiple concerns (e.g., create temp file + atomic rename)
  so each has one clear responsibility. Use `defer` to enforce cleanup invariants
- Extract shared logic into helpers at 3+ duplications, not before
- Adding a required field to a public type is a breaking change — use nil checks
  and defaults to preserve existing behavior
