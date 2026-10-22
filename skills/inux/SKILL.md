---
name: inux
description: Set up Inux tools, start Jira tasks, create branches, commit changes, and draft or refine estimates, technical summaries, and pull requests. Use for /inux commands and equivalent requests; publish only when explicitly requested.
metadata:
  author: "Vinod Liyanage"
---

# Inux

Inux requires Flow 0.4.0. For an Inux execution request, load the installed
`flow` skill and its complete current interpreter, verify that it supports
Flow 0.4.0, then run [flows/inux.flow](flows/inux.flow).

For direct CLI usage, read [cli.md](references/cli.md).

If Flow is unavailable or older than 0.4.0, stop and provide its installation
or update command:

```sh
npx skills add vinodliyanage/flow --skill flow --global --agent codex cursor
```
