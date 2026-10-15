# Publish a saved PR draft

Require explicit publishing intent and initialized ticket.md/pr.md. Read the
ticket heading and project rules, confirm the developer's target, and rerun the
[Git gate](readiness.md#git-gate) before reviewing implementation or publishing.

Read the full ticket and saved draft. Require a matching ticket, a completed
single-line Title, and a complete Description under [summary](summary.md).
If either section is missing, incomplete, or has template placeholders, request
preparation/refinement; do not draft and publish together.

Compare the saved summary and checked commit messages with the full
merge-base-to-HEAD diff. Resolve material gaps through direct dependencies only.
Report stale content or incomplete work instead of rewriting during publication.
Complete [verification](readiness.md#verification) before any push or PR mutation.

Check private gh authentication and open PRs for this exact repository/head.
Multiple matches require clarification. Require an existing PR to target the
resolved base; resolve a differing target before mutation.

Perform the [final recheck](readiness.md#final-recheck), also confirming the saved
draft is unchanged. Push the task branch without force and verify its remote SHA
equals checked HEAD. Create a draft PR with explicit base/head and self-assignment,
or update the existing open PR for this head. Use explicit `--repo` for gh PR
operations and preserve other assignees when adding the authenticated user.
Use the saved Title and a temporary UTF-8 body file containing only Description,
preserving its Markdown exactly.

Inspect the result, report its URL, and attach the PR to this chat when supported.
If push succeeds but PR creation fails, report both. Inspect remote state before
retrying an uncertain write. Remove temporary output; leave pr.md unchanged.
