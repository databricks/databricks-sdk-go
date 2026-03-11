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
- No `Get` prefix on methods — not idiomatic in Go. Use `HostType()` not `GetHostType()`
- Avoid abbreviations: `authTokenCommand` not `authTokenCmd`
- Variable names describe content, not origin: `workspaceHost` not `resolvedHost`
- Use specific names when multiple similar types exist: `httpClient` not `client`

## Go Idioms

- **Exhaustive switches**: prefer `switch` over `if-else` chains for type/enum dispatching.
  Always include a `default` case that returns an error for unknown values
- **`sync.Once`** over manual double-checked locking with mutexes — it is safer and
  respects Go's memory model
- **Scope variables tightly**: use `:=` inside `if` blocks when the variable is not needed
  outside: `if err := json.Unmarshal(raw, &f); err != nil {`
- **Use stdlib functions** over manual equivalents: `math.Ceil`, `json.Marshal`,
  `slices.Sort`, `slices.Compact`, `sort.Strings`
- **No forced line length**: Go does not mandate 80-char lines. Prefer keeping function
  signatures on one line, or use one argument per line for very long signatures
- **Config structs over long parameter lists**: when a function takes many parameters,
  group them in a config struct
  ([Google Go style](https://google.github.io/styleguide/go/best-practices#option-structure))
- **Don't alias imports unnecessarily**: only alias when there is a genuine name collision
- **Use `sync.Mutex` directly**, not through an interface — the zero value is valid and
  needs no initialization

## Code Organization

- **Keep related code close together**: place constructors near the types they return,
  move helper functions near their callers
- **Declare variables close to where they're used**: don't declare at the top of a function
  if the variable is only needed 30 lines later
- **Inline simple helpers** when they obscure intent — a one-line function called once
  adds indirection without value. Remove unnecessary intermediate variables that suggest
  the value might be mutated when it won't be
- **Move branching out of callbacks**: if a visitor/callback is only needed conditionally,
  check the condition first, then register the callback — don't put the `if` inside

## Comments and Documentation

- **Full sentences**: comments start with a capital letter and end with a period.
  Ref: [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments#comment-sentences)
- **Explain why, not what**: the code is self-explanatory for *what* it does; comments
  should explain *why* a decision was made or why a non-obvious approach was chosen
- **Use Go doc links**: reference types with bracket syntax: `// implements the [TokenCache] interface`
- **Document public types thoroughly**: explain the purpose, which fields are optional,
  and what the zero value means
- **Don't duplicate godoc**: if a method's doc comment says the same thing as the type's
  doc comment, remove the redundancy
