# .nextchanges

Pending changelog entries for the next release. Each change is a
`.nextchanges/<timestamp>.md` file that `genkit prepare-release` collects, folds
into `CHANGELOG.md` and the `.codegen/releases.jsonl` ledger, and then deletes.

An entry looks like:

```markdown
---
version_bump: patch
section: Bug Fixes
---
* Describe the change.
```

`version_bump` is `minor` or `patch`. `section` must match a section in
[.codegen/changelog.json](../.codegen/changelog.json). Changes that do not need
a release note do not require a fragment. API change entries are generated
from the spec diff during release preparation.
