# Inux CLI

Resolve `<cli>` to the binary beside this skill: `../bin/ix` relative to
this reference, or `ix.exe` on Windows. Pass arguments separately and use
an absolute PROJECT path (the repository root in a Git checkout). The compiled tool needs no Go or Node runtime.

```text
<cli> init PROJECT
<cli> reset PROJECT
<cli> remove PROJECT
<cli> adf INPUT.md OUTPUT.json
```

- `init` creates .inux/ and missing estimate/PR templates, preserving drafts.
- `reset` creates .inux/ if needed and replaces both template files. It preserves
  ticket.md and other files. Use once for a new task after Jira retrieval succeeds.
- Both add .inux/ to .git/info/exclude directly in an ordinary Git checkout.
  Existing entries and .gitignore are preserved. This MVP does not support
  linked worktrees or discover a repository above PROJECT.
- `remove` deletes only PROJECT/.inux and everything within it. It leaves the
  Git exclusion in place. Only use when deletion is requested.
- `adf` reads UTF-8 Markdown and writes ADF JSON. Omit OUTPUT to receive the actual
  ADF document on stdout. It never replaces its Markdown input, including aliases.

File operations and conversion-to-file report JSON on stdout. Errors are JSON
on stderr with exit status 1; never assume success from partial output. No stdin,
selective reset, task history, install/update service, or publishing wrappers.
The native MVP package contains one platform's binary. If it cannot run, report
that limitation; do not invent a release URL or silently use an unrelated CLI.
