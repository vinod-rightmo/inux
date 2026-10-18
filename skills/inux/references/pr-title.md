# Write PR title

Write only the content under `## Title` in the task's `pull-request.md`.

Read the existing title and sibling `ticket.md`, then inspect the current
branch, relevant commits, and actual changed behavior. Preserve the title's
accurate intent and any repository-required ticket ID or prefix. Use repository
PR or commit history only to learn an established title convention.

Produce one concise line that states the main delivered change. Prefer specific
behavior over vague wording such as “updates” or “fixes.” Include a ticket ID,
scope prefix, or conventional-commit style only when the repository consistently
uses it. Do not claim work that is absent from the implementation.

Fill an empty title or refine an existing one. Remove redundancy, unsupported
scope, and implementation detail that belongs in the description.

Preserve `## Description` and every other part of `pull-request.md` exactly.
