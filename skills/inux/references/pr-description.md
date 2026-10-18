# Write PR description

Write only the content under `## Description` in the task's `pull-request.md`.

Read the whole existing description and sibling `ticket.md`, then inspect the
current branch, its base or merge-base when available, relevant commits, and
uncommitted changes. Use an existing repository PR template or strong local
convention when present; otherwise write compact Markdown covering:

- what behavior changed;
- why the change is needed; and
- verification that was actually observed.

Write the complete technical summary directly in the description using
[technical-summary-writing.md](technical-summary-writing.md) for its fields,
structure, and detail level. Derive it from the actual implementation; there is
no separate technical summary file to read or create. Avoid repeating the same
information in a separate overview.

Fill an empty description or refine an existing draft. Preserve accurate
developer-authored context and intentional emphasis. Add missing high-value
impact, correct stale or unsupported claims, and merge overlapping content so
the technical summary appears once. Keep useful repository template headings.

Describe behavior and important technical decisions, not a file-by-file log.
Do not copy the full Jira ticket, invent test results, or claim intended work as
implemented. Mention material limitations or follow-up work only when evidenced.

Preserve `## Title` and every other part of `pull-request.md` exactly.
