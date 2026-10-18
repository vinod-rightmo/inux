# Inux

Inux is an agent skill with a small native CLI. The agent uses `ix` for local
task files and Markdown-to-ADF conversion, `gh` for GitHub, and `acli` for Jira.

## Repository layout

```text
skills/
  inux/
    SKILL.md                    Commands and workflow routing
    agents/openai.yaml          Agent display metadata
    bin/ix                      Bundled native CLI (ix.exe on Windows)
    references/                 Setup, task, writing, and publishing workflows
    THIRD_PARTY_NOTICES.txt      Dependency notices
```

The repository distributes the skill and its bundled CLI. It has no editor
extension, VSIX package, or Node build step. The current `ix` binary is for
macOS on Apple Silicon; this checkout does not include its Go source or other
platform builds. The installed skill does not require a compiler.

## Task files

Each project has one task in a flat directory:

```text
<project>/.inux/
  ticket.md
  estimate.md
  pull-request.md
```

Technical summary content is written directly in the PR description. There are
no ticket subdirectories or hidden extension state. `ix init` and `ix reset`
add `.inux/` to the ordinary checkout's local Git exclude without changing
`.gitignore`.

## Install and use

Install Flow and Inux with the Skills CLI:

```sh
npx skills add vinodliyanage/flow --skill flow --global --agent codex cursor
npx skills add vinod-rightmo/inux --skill inux --global --agent codex cursor
```

The complete `skills/inux/` directory, including `bin/ix`, is installed in each
agent's skills directory. Start a new agent conversation after installation or
after updating an installed skill.

Ask the agent for `/inux help` to list all commands, or `/inux setup` to check
the bundled CLI and set up GitHub and Jira tools. Setup stores one shared copy
of `gh` and `acli` under `~/.inux/tools/bin/`, outside every project's task
directory. Credentials remain managed by those CLIs.

Use `/inux new task TICKET-ID` to fetch a ticket and replace the current task
drafts. `/inux estimate`, `/inux pr title`, and `/inux pr description` each
draft empty content or refine existing content. Publishing happens only through
an explicit `/inux publish estimate` or `/inux publish pr` request.

## Verify the bundled CLI

```sh
./skills/inux/bin/ix --help
```

The CLI exposes `init`, `reset`, `remove`, and `adf`. It does not download tools,
authenticate, or publish; the agent invokes `gh` and `acli` directly.
