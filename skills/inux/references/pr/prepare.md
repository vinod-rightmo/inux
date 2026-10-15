# Prepare a PR draft

## Validate first

Require initialized `.inux/ticket.md` and `.inux/pr.md` without resetting them.
Read the ticket heading for its ID and applicable project `AGENTS.md` rules;
confirm the developer's target and complete the [Git gate](readiness.md#git-gate).
Stop on failure before reading the full task files, reviewing source, or drafting.

## Gather evidence and compose

After the gate passes, read the full ticket and initialized PR draft, then use
[summary](summary.md) and [writing guidance](../be-human.md). The primary evidence
is the ticket, project rules, checked commit messages, and full merge-base-to-HEAD
diff. Reuse messages already captured during validation.

Follow the draft's existing labels/order. Read the [template](../../assets/pr.md)
only if required labels are missing. Write one concise line under `## Title` and
the technical summary inside `## Description`. Title defaults to
`<TICKET-ID>: <Capitalized action summary>` unless user/project rules differ.
Preserve required sections and accurate developer edits.

Resolve only material gaps in the summary using direct dependencies. Name the
missing fact first, follow exact imported paths/symbols, and read the relevant
source range. Stop when the summary's claims are supported; report unresolved
material questions. Examples illustrate established behavior, not extra scope.

Complete [required verification](readiness.md#verification). Once check scope is
known, independent checks and read-only evidence review/tentative composition may
overlap when the host supports it. Collect every result before saving; reuse the
composed explanation instead of planning it again after checks finish.

## Save

Perform the [final recheck](readiness.md#final-recheck), then save pr.md only after
readiness succeeds. A successful write completes preparation; no readback or
byte/whitespace audit is needed unless the write fails or a formatting defect is
observed. Report the path and checks actually completed. Preparation does not
commit, push, or publish.
