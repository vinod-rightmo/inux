# Technical summary

Use Summary content below for both `.inux/technical-summary.md` and the
Description in `.inux/pr.md`. For writing, read [writing guidance](be-human.md) once.
PR actions keep their [PR workflow](pr.md). Only `/inux technical summary` uses
the standalone actions below; publication requires explicit intent. Do not
automatically synchronize the files.

## Create or refine

`/inux technical summary [request]` creates or refines `.inux/technical-summary.md`.
Require ticket.md, but allow the summary file to be missing. Read the ticket and
existing summary together, preserving accurate developer notes and unrequested
content. Stop if an existing summary identifies another ticket; never reset the task.

Use the requested implementation scope and current evidence. Inspect branch/HEAD
before reusing earlier code evidence. Review relevant local changes or an explicitly
identified commit range; ask only if scope is unclear. Follow exact imports or
symbols only to resolve material gaps. Wording edits need no code investigation.
Do not run tests solely to write prose; distinguish observed checks from proposed
verification and report incomplete implementation honestly.

For a new file, use `# TICKET-ID` followed by the Description body from the
[PR template](../assets/pr.md), without its PR title or Description heading.
For existing files, preserve the layout and apply Summary content below.
Recheck the active ticket before saving only technical-summary.md. Report the
path briefly. No PR target, fetch, clean-tree gate, commits, or publishing are
required for this local writing action.

## Publish to Jira

`/inux technical summary publish` posts the saved standalone summary as a Jira
comment. Require ticket.md and technical-summary.md with matching ticket IDs,
completed Summary content fields, and no placeholders or unsupported claims.
Report missing/incomplete content; do not create or rewrite the draft while publishing.

Use private acli under the shared Jira execution rule:
`jira workitem comment list --key TICKET-ID --json --paginate`.
Check duplicates and material clarifications. Report identical content without
posting again; never replace unrelated comments. If evidence contradicts the
saved summary, stop and ask whether to revise it, leaving the file unchanged.

Convert with private ix: `adf ABSOLUTE_SUMMARY_PATH TEMPORARY_OUTPUT.json`.
Inspect the ADF and recheck ticket ID and saved content before posting.
Use private acli:
`jira workitem comment create --key TICKET-ID --body-file ABSOLUTE_ADF_PATH --json`.
Consult only exact subcommand help when needed; never rebuild ADF manually.
Inspect remote state before retrying uncertain writes. Remove temporary output
on success or failure. Report the comment/ticket URL; look it up only if missing.
Leave task files unchanged. This action posts only the standalone summary.

## Summary content

Write the summary in the requested destination, preserving its labels and order.
Explain the result, mechanism, old behavior/cause, and relevant impacts or limits.
Use only supported claims; report material uncertainty and incomplete work.

- Branch name: exact current name; keep the label's two trailing spaces.
- Fix / implementation: a short paragraph explaining what happens when the user
  acts, followed by how the app achieves it. Add a brief verified example if useful.
- Root cause: a short paragraph explaining what the app did before, which action
  was missing/incorrect, and the problem the user saw. Use None for non-defect work.
- Impacted services: affected repository names only, such as `qr-web-frontend`;
  no organization prefixes, URLs, or workflow descriptions. Resolve from
  origin/GitHub, falling back to the repository root name.
- Core modules/sections: affected workflows/features, not a file log.
- UI, DB, indexes, environment, infra: brief evidenced changes, or None when
  verified unaffected. Include indirect impacts/other services when code indicates them.

Use familiar words and simple sentences, one action or idea each, in event order.
The opening must make sense on its own. Introduce internal terms only when needed;
avoid repetition and sentences combining several steps. Each explanation stays
as a paragraph on its existing label line, without numbered steps or nested bullets.
Retain field-label bullets, blank lines between sections/before bullets, and no
blank lines between field bullets. Omit `- Important notes:` by default; do not
write an empty field or None. Add it after Root cause only for an evidenced fact
that could change the merge/deployment decision or requires a developer action,
such as a breaking contract, mandatory migration/configuration step, or serious
unresolved correctness risk. Routine implementation details, test results, minor
caveats, optional improvements, and facts already covered elsewhere do not qualify.
If no such fact exists, leave the section out entirely.
Replace all placeholders in a completed draft. Routine passing checks belong in
the completion report; claim success only when observed.
