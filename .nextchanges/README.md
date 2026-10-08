# Release note fragments

For a change that needs a release note, add a uniquely named
`.nextchanges/<timestamp>.md` file, for example:

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
