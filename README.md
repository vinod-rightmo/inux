# Inux

Start Jira tasks, draft estimates, and prepare pull requests from your coding agent.

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
```

After implementing and committing your changes, prepare the pull request:

```text
/inux pr
```

## Commands

| Command | What it does |
| --- | --- |
| `/inux help` | Show the available commands. |
| `/inux setup` | Set up tools and check sign-in. |
| `/inux task <id>` | Load a Jira ticket and create the task files. |
| `/inux estimate [request]` | Draft, improve, or explain an estimate. |
| `/inux pr [request]` | Prepare a PR draft, refine it, or explain it. |

Use plain language to refine your drafts:

```text
/inux estimate improve the reasons
/inux pr improve the title
/inux pr improve the description
```

## Your task files

Each task uses three files in your project:

```text
.inux/
├── ticket.md     Jira ticket details
├── estimate.md   Estimate and reasons
└── pr.md         PR title and description
```

The technical summary lives inside the PR description. You can edit the drafts
directly or ask the agent to refine them.

**Starting a task replaces the previous `.inux/` task files after the ticket
is fetched successfully.** The folder is excluded from Git locally.

## Publish when ready

Drafting and refining stay local. Publishing requires an explicit request:

```text
/inux estimate publish
/inux pr publish
```

Estimate publishing posts a Jira comment. PR publishing pushes the task branch
and creates or updates a draft GitHub PR after readiness checks pass.

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
The agent fetches Jira and fills ticket.md; filesystem checks stay inside ix.

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
