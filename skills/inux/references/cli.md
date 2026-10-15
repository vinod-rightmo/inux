# Inux CLI

Resolve `<cli>` to `<user-home>/.inux/tools/bin/ix` (`ix.exe` on Windows),
or `<storage>/tools/bin/ix` when an explicit Inux storage root is supplied.
Pass arguments separately and use an absolute PROJECT path (the ordinary
repository root in a Git checkout).
The compiled tool needs no Go or Node runtime. --version prints `ix <version>`;
release builds use the same version as the skill, local unversioned builds use dev.

```text
<cli> --version
<cli> task --snapshot PROJECT
<cli> task [--expect SNAPSHOT] TEMPLATES_DIR PROJECT
<cli> adf INPUT.md OUTPUT.json
```

For `task`, pass the skill assets directory; ix reads both estimate.md and
pr.md before replacement, copies both drafts, and creates empty ticket.md.
`--snapshot` returns an opaque state token without writes; `--expect` rejects
changes before replacement. Follow [task.md](task.md) for the complete workflow
and Git exclusion. The agent passes an absolute PROJECT;
direct CLI use may omit it to use the current directory. ix never fetches Jira
or invokes another CLI.

- `adf` reads UTF-8 Markdown and writes ADF JSON. Omit OUTPUT for ADF on stdout.
  It never replaces its Markdown input, including aliases.

The agent performs branch and commit convention checks directly; ix handles
task files and ADF conversion.

File operations and conversion-to-file report JSON on stdout. Errors are JSON
on stderr with exit status 1; never assume success from partial output. No stdin,
selective reset, task history, install/update service, or publishing wrappers.
The skill package contains instructions and templates, without executables. If ix is missing or cannot run,
follow [setup.md](setup.md) when setup is requested; otherwise report that setup
is needed. Do not silently use a PATH binary or another platform's executable.
