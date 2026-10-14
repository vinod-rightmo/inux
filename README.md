# Inux MVP

A small Go CLI plus an agent skill. The skill exposes these commands:

```text
/inux help
/inux setup
/inux task <id>
/inux estimate <request>
/inux pr <request>
/inux test task [id]
/inux test estimate [request]
/inux test pr [request]
```

Requests are plain language: generate or improve an estimate, publish it, prepare
a PR, refine its title/technical summary, or publish the reviewed draft.

The agent fetches the exact Jira ticket first for `/inux task <id>`. Once retrieval
succeeds and both asset templates are readable, it replaces `.inux/`, copies
estimate.md and pr.md templates, and writes ticket.md.
Failed retrieval leaves the previous task intact. This command authorizes replacing
all old task files, including custom files; it does not create a branch or publish.

The CLI also has four commands, each with one handler file:

```text
ix task TEMPLATE [PROJECT]
ix adf INPUT.md [OUTPUT.json]
ix check-branch PATTERN NAME [NAME...]
ix check-commit PATTERN MESSAGE [MESSAGE...]
```

The skill passes its absolute `skills/inux/assets/estimate.md` path as TEMPLATE.
PROJECT defaults to the current directory. For Git, pass the ordinary checkout
root. ix validates the template and checkout before deleting PROJECT/.inux,
recreates it with estimate.md, and locally excludes it using .git/info/exclude.
The project's .gitignore and other files remain unchanged. Linked worktrees and
ancestor discovery are outside this MVP; plain directories also work.

The full `/inux task <id>` workflow validates both templates before replacement.
After ix task succeeds, the agent copies assets/pr.md into `.inux/pr.md` and
saves ticket.md. All three task files exist before estimate/PR writing starts. ix never fetches Jira, invokes another CLI, or publishes.
The old init/reset/remove CLI commands and standalone skill commands are removed.

ADF conversion uses pinned `github.com/ericmason/mdadf v0.1.3` from the earlier
Inux Go CLI. It supports headings, paragraphs, common inline formatting, lists,
code blocks and GFM tables. Images become links; raw HTML remains text. It is not
a lossless Markdown round-trip converter. Review generated output before publishing.

Operations report JSON; errors go to stderr with exit code 1. ADF without an output
filename prints the actual ADF document. Conversion preserves its Markdown input,
including hard-link/symlink aliases. Files are replaced through temporary files.

## Help and testing

Skill version and author are stored in SKILL.md frontmatter under `metadata`.
`/inux help` displays these values above the command list.

`/inux help` uses [help.md](skills/inux/references/help.md) and needs no active
project or ticket. There is one reference per main command: help.md, setup.md,
task.md, [estimate.md](skills/inux/references/estimate.md), and
[pr.md](skills/inux/references/pr.md). Estimate/PR actions, validation, and testing
are sections in those files. Follow only the requested action's section;
readiness runs only for PR preparation/publication. Shared test rules live in
[SKILL.md](skills/inux/SKILL.md), and [be-human.md](skills/inux/references/be-human.md)
remains the writing guide. Add `test` after `/inux` to test an action, such as
`/inux test estimate improve` or `/inux test pr title`. Bare estimate/PR tests
run their non-publishing actions. Refinements reuse verified evidence and inspect more
code only when needed. This reduces unnecessary work; it does not set the host's
reasoning effort or guarantee a runtime. Run `/inux test task`, `/inux test estimate`,
or `/inux test pr`; add a real ticket ID to the task test for live Jira reads.
Tests write only to disposable projects and retain a pass/fail/blocked/skipped
report. They exclude publishing, tool installation, and login. PR tests use
a local Git fixture to exercise readiness without changing your source branch.
Update installed skill copies and start a new conversation to load changes.

## Pull-request workflow

`/inux pr prepare` checks the current branch convention, every commit introduced
relative to the convention baseline (remote main by default), and a clean Git
index/working tree, including untracked files and submodules. The PR target is
separate (dev by default). The established team branch/commit rules apply unless
explicit user/repository rules override them; no custom regex input is needed
for the defaults. Missing refs or failed checks block preparation. The agent also reviews the committed
changes and required verification. After checks pass, it fills the already
initialized `.inux/pr.md`, with `## Title` and the technical summary under
`## Description`. Failed checks leave that draft unchanged.
Its complete reusable template is [assets/pr.md](skills/inux/assets/pr.md); writing
instructions live in the [PR technical-summary section](skills/inux/references/pr.md#technical-summary).
The format retains the earlier Inux branch, implementation/root cause, code,
UI, database, and infrastructure fields. The completed summary stays inside
pr.md, so no extra task file is created.

Edit pr.md directly or refine either section through chat. `/inux pr publish`
reruns readiness checks and publishes the saved, reviewed content. Title and
description requests fill/refine only that initialized section without running
preparation. Incomplete or stale drafts need preparation/refinement before publishing;
missing task files indicate incomplete initialization and do not trigger a reset. Preparation does
not push or create a PR. Git inspection and gh publication stay in the skill;
ix performs file management and regex validation only.

## Convention checks

Team policy lives in [PR readiness checks](skills/inux/references/pr.md#readiness-checks):

- Branch: `<type>/<TICKET-ID>-<lowercase-kebab-description>` with a supported type
  and the exact active ticket.
- Commit: `<TICKET-ID>: <Capitalized present-tense implementation summary>`;
  exactly one line, no body, and no duplicate active-ticket token.
- PR title defaults to `<TICKET-ID>: <Capitalized action summary>`.

Explicit user/repository policies override their stated defaults. Nonconforming
history does not redefine the rule. ix remains a policy-neutral Go regex matcher;
the skill supplies the resolved default/override regex. Every entire supplied
value must match; inputs are preserved without trimming. Regex validates
structure, while the agent checks duplicates, meaning, and implementation accuracy.

```sh
ix check-branch '^(feature|bugfix|hotfix|chore|refactor|perf|test|release)/OQG3-3776-[a-z0-9]+(-[a-z0-9]+)*$' 'bugfix/OQG3-3776-home-step-number-centering'
ix check-commit '^OQG3-3776: [A-Z][^\r\n]+$' 'OQG3-3776: Fix QGC home step number centering'
```

The ticket in these examples is resolved from the active task. Read full stored
commit messages and remove at most their single terminal LF before validation;
do not use only subjects or strip bodies/whitespace. Check every introduced
commit, including merges, against the convention baseline; separately inspect
the actual PR target diff. Failed input positions map back to commit SHAs.

Checks return `{"ok":true}` and exit 0 on success; invalid patterns, mismatches,
or missing arguments produce JSON errors on stderr and exit 1. Checks do not
change files or call Git/network. Git ref validity and semantic rules are checked
separately. The CLI accepts alternate policies, including empty input only when
its supplied pattern explicitly allows it; that is not the team's default.

## Build and release

GitHub builds the executables. Commit source/configuration only; dist/ stays
ignored. One version is shared by the skill and ix: the quoted
`metadata.version` in skills/inux/SKILL.md. ix reports it through `--version`.
The Go version in go.mod is a compiler requirement, not the Inux version.

To publish a release:

1. Change metadata.version in SKILL.md, for example to "0.0.2".
2. Commit and push that change to main.
3. Tag that commit and push the matching tag:

```sh
git tag v0.0.2
git push origin v0.0.2
```

[The workflow](.github/workflows/release.yml) tests the CLI on Ubuntu, Windows,
and macOS. A version tag then builds five binaries, packages inux-skill.zip,
generates checksums, and publishes them to GitHub Releases. A tag/version
mismatch fails before publishing. No Changesets, release bot, or manually
uploaded executables are needed. The workflow uses the repository's built-in
GitHub token; no extra release secret is required.

[The release script](scripts/release.py) uses Python's standard library and Go.
To test/build locally, with Go and Python 3 available:

```sh
cd tools/ix
go test ./...
cd ../..
python3 scripts/release.py
```

Outputs are under dist/release/. The script never publishes; only the tag-triggered
workflow does. The release includes Windows x64, Ubuntu/Linux x64 and ARM64,
and macOS Intel and Apple Silicon binaries, the skill ZIP, third-party notices,
and checksums.txt. Users need Git for project operations, but no Go compiler.

Install the skill ZIP, then use `/inux setup` to install matching tools from
`vinod-rightmo/inux`. Setup verifies checksums and ix --version. gh/acli keep
their own vendor versions. Existing projects/tasks are unaffected by setup.

Templates stay under skills/inux/assets/: ix task copies estimate.md; the agent
copies pr.md and writes ticket.md after successful Jira retrieval. PR preparation
later fills the initialized PR template. Both assets are checked before task
replacement. There is no separate technical-summary task file.

## Source structure

```text
tools/ix/
  main.go                   Entry point and JSON error reporting
  go.mod
  go.sum
  internal/
    run.go                  Dispatch and help
    commands/               One handler file per command
      task.go
      adf.go
      check_branch.go
      check_commit.go
    shared/                 Shared operations
      workspace.go
      files.go
      conventions.go
  tests/
    run_test.go
    conventions_test.go
```

The root main package calls the dispatcher in internal/run.go. commands contains
only the four command handlers; shared holds workspace operations, atomic writing,
and regex matching. Existing tests live in tests and exercise the dispatcher’s
Run function. To add a command, add its handler and a dispatcher case, then test
through Run. There is no command registry or CLI framework.
