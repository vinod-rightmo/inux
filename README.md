# Inux

Start Jira tasks, create branches, commit changes, and prepare estimates, technical summaries, and pull requests from your coding agent.

**Codex & Cursor · macOS, Linux & Windows**

## Get started

You need Node.js, Git, and access to this private GitHub repository.

**1. Install the skill**

```sh
npx skills add vinod-rightmo/inux --skill inux --global --agent codex cursor
```

**2. Set up the tools**

Open a new conversation in Codex or Cursor and run:

```text
/inux setup
```

Setup installs or checks `ix`, `gh`, and `acli` in `~/.inux/tools/bin/`,
checks authentication, and cleans up its temporary downloads.

**3. Start a task**

```text
/inux task PROJECT-123
/inux estimate
/inux branch create
```

After implementing your changes, commit and prepare the pull request:

```text
/inux commit changes
/inux pr target release/2026-10
```

## Commands

| Command | What it does |
| --- | --- |
| `/inux help` | Show the available commands. |
| `/inux setup` | Set up tools and check sign-in. |
| `/inux task <id>` | Load a Jira ticket and create the task files. |
| `/inux branch create` | Create and switch to a local task branch. |
| `/inux commit changes` | Stage intended changes and create a local commit. |
| `/inux estimate [request]` | Draft, improve, or explain an estimate. |
| `/inux technical summary [request]` | Create, refine, or publish `.inux/technical-summary.md` to Jira. |
| `/inux pr [request]` | Prepare a PR draft, refine it, or explain it. |

Branch creation starts at current HEAD unless you supply another starting ref.
Both Git actions stay local and follow the target repository's conventions.
They reuse task context and verification without running PR readiness checks.
Commit scope must be clear; unrelated changes are preserved.

Use plain language to refine your drafts:

```text
/inux estimate improve the reasons
/inux pr improve the title
/inux pr improve the description
```

Provide the PR destination branch when preparing or publishing:

```text
/inux pr target release/2026-10
/inux pr publish target release/2026-10
```

If omitted, the agent asks for it. An explicit choice already supplied in the
conversation is reused.
Commit checks use the chosen target branch to exclude inherited history.

## Your task files

Each task starts with three files in your project:

```text
.inux/
├── ticket.md     Jira ticket details
├── estimate.md   Estimate and reasons
└── pr.md         PR title and description
```

Run `/inux technical summary` to create an optional `.inux/technical-summary.md`.
It uses the same summary instructions as the PR description. Each command edits
its own file. You can edit the drafts directly or ask the agent to refine them.

**Starting a task replaces the previous `.inux/` task files after the ticket
is fetched successfully.** The folder is excluded from Git locally.

## Publish when ready

Drafting and refining stay local. Publishing requires an explicit request:

```text
/inux estimate publish
/inux technical summary publish
/inux pr publish
```

Estimate and technical-summary publishing post the saved file as a Jira comment.
PR publishing pushes the task branch and creates or updates a draft GitHub PR
after readiness checks pass.

To try the normal actions without publishing, add **“do not publish”** to your request.

<details>
<summary><strong>Developing and releasing Inux</strong></summary>

### Local checks

The skill instructions live in [skills/inux](skills/inux/SKILL.md).
The Go CLI source lives in [tools/ix](tools/ix/).
`ix task TEMPLATES_DIR [PROJECT]` reads estimate.md and pr.md from the assets
directory before replacing the task, then copies both drafts and creates empty
ticket.md. `ix task --snapshot [PROJECT]` captures task state; pass that token
with `--expect SNAPSHOT` before the assets argument to reject intervening edits.
The agent fetches Jira, fills ticket.md, and checks branch and commit naming
conventions directly. Task filesystem checks and ADF conversion stay inside ix.

```sh
(cd tools/ix && go test ./...)
goreleaser check
goreleaser release --snapshot --clean
```

The last two commands require GoReleaser. Snapshot mode builds locally without
publishing.

### Releases

Use Conventional Commits in this repository:

| Commit | Version change |
| --- | --- |
| `fix: ...` | Patch |
| `feat: ...` | Minor |
| `feat!: ...` | Major |

Merge changes into `main`. Release Please opens a release PR with the version
and changelog. Review and merge that PR; GitHub then tests, builds, and publishes
the downloads with GoReleaser. The skill and `ix` share one version.

Keep the `x-release-please-version` comment in `SKILL.md`; it lets Release Please
update the skill version automatically.

[Release workflow](.github/workflows/release.yml) ·
[Release configuration](release-please-config.json) ·
[Build configuration](.goreleaser.yaml) ·
[CLI reference](skills/inux/references/cli.md) ·
[Implementation decisions](docs/decisions.md)

</details>
