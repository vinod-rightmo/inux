# Inux CLI

Resolve `<cli>` to `<user-home>/.inux/tools/bin/ix` (`ix.exe` on Windows),
or `<storage>/tools/bin/ix` when an explicit Inux storage root is supplied.
Pass arguments separately. The compiled tool needs no Go or Node runtime.
--version prints `ix <version>`; release builds use the same version as the
skill, local unversioned builds use dev.

```text
<cli> --version
<cli> adf INPUT.md OUTPUT.json
```

`adf` reads UTF-8 Markdown and writes ADF JSON. Omit OUTPUT for ADF on stdout.
It never replaces its Markdown input, including aliases. Conversion-to-file
reports JSON on stdout. Errors are JSON on stderr with exit status 1; never
assume success from partial output.

The agent fetches Jira through acli, manages temporary drafts using the selected action's
knowledge contexts, and checks branch and commit conventions directly.
ix does not create task files, manage workspaces, invoke another CLI, install,
authenticate, or publish. The skill package contains instructions and templates,
without executables. If ix is missing or cannot run, follow
[setup.flow](../flows/setup/setup.flow) when setup is requested; otherwise report that
setup is needed. Do not silently use a PATH binary or another platform's executable.
