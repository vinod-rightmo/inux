# Inux CLI

Resolve `<cli>` to `<user-home>/.inux/tools/bin/ix` (`ix.exe` on Windows),
or `<storage>/tools/bin/ix` when an explicit Inux storage root is supplied.
Pass arguments separately and use an absolute PROJECT path (the ordinary
repository root in a Git checkout).
The compiled tool needs no Go or Node runtime. --version prints `ix <version>`;
release builds use the same version as the skill, local unversioned builds use dev.

```text
<cli> --version
<cli> task ESTIMATE_TEMPLATE PROJECT
<cli> adf INPUT.md OUTPUT.json
<cli> check-branch PATTERN NAME [NAME...]
<cli> check-commit PATTERN MESSAGE [MESSAGE...]
```

Resolve ESTIMATE_TEMPLATE to [../assets/estimate.md](../assets/estimate.md) relative
to the skill and pass its absolute path. In direct CLI use, PROJECT is optional
and defaults to the current directory. ix never fetches Jira or invokes another CLI.

- `task` replaces the entire PROJECT/.inux directory, including all old task
  files, and creates estimate.md from the supplied template. The full skill
  workflow validates both assets first, calls ix after successful Jira retrieval,
  then copies assets/pr.md and writes ticket.md. ix itself copies only the estimate.
  Missing/unreadable estimate input, directory links, and unsupported checkouts
  fail before deletion.
- `task` adds .inux/ to .git/info/exclude directly, preserving existing entries
  and .gitignore. Pass an ordinary checkout root; linked worktrees and ancestor
  discovery are outside this MVP. Plain directories also work.
- `adf` reads UTF-8 Markdown and writes ADF JSON. Omit OUTPUT for ADF on stdout.
  It never replaces its Markdown input, including aliases.

Convention commands accept a caller-supplied [Go regex](https://pkg.go.dev/regexp/syntax).
Every entire value must match; inputs are not trimmed. ix itself stays policy-neutral.
For PR work, the skill supplies the team rules in [PR readiness checks](pr.md#readiness-checks),
including exact active-ticket matching. Only read that reference for Git policy
work, not for task setup or ADF conversion. The default commit check uses full
single-line messages; subject-only checks apply only to an explicit subject rule.

Resolved examples for active ticket OQG3-3776:

```sh
ix check-branch '^(feature|bugfix|hotfix|chore|refactor|perf|test|release)/OQG3-3776-[a-z0-9]+(-[a-z0-9]+)*$' 'bugfix/OQG3-3776-home-step-number-centering'
ix check-commit '^OQG3-3776: [A-Z][^\r\n]+$' 'OQG3-3776: Fix QGC home step number centering'
```

Quote the regex and each value for the host shell or pass separate arguments.
A matching convention returns `{"ok":true}` and exit status 0. A mismatch, invalid regex,
or missing arguments returns JSON on stderr and exit status 1. All values are
checked; mismatches report every failing value by its 1-based position after
the regex. Patterns that allow empty input also accept an empty supplied value. Checks perform no file
writes or Git/network calls.

File operations and conversion-to-file report JSON on stdout. Errors are JSON
on stderr with exit status 1; never assume success from partial output. No stdin,
selective reset, task history, install/update service, or publishing wrappers.
The skill package contains instructions and templates, without executables. If ix is missing or cannot run,
follow [setup.md](setup.md) when setup is requested; otherwise report that setup
is needed. Do not silently use a PATH binary or another platform's executable.
