# Technical summary

Use ticket.md for intent, project `AGENTS.md` for conventions/required checks,
checked commit messages for stated changes, and the full merge-base-to-HEAD diff
for delivered behavior. Implementation evidence governs implementation claims;
the ticket and messages alone do not prove them. Reuse current evidence rather
than rereading it. Read direct dependencies only for a material missing fact.

Fill the existing fields under `## Description` in `.inux/pr.md`. Its labels/order
are authoritative; keep the summary once, preserve required sections and accurate
developer notes, and keep Title unchanged for a description-only edit. Use
[writing guidance](../be-human.md) for readable prose.

Support the result, mechanism, old behavior/cause, and reviewer-relevant impacts
or limitations. Stop when those claims are established; omit unnecessary
background. Keep causal claims within the evidence and report material uncertainty.
Use the full diff for changed services/configuration and inspect indirect impacts
when the implementation indicates them.

- Branch: exact current name below `Branch name:`; retain the label's two trailing
  spaces for the Markdown hard break.
- Fix / implementation: write a short paragraph. First explain what now happens
  when the user performs the action. Add a brief concrete example when useful,
  then explain how the app makes that happen. Examples must match verified behavior.
- Root cause: write a short paragraph explaining the old sequence: what the app
  did, which action was missing or incorrect, and the problem the user saw as a
  result. Explain the failure rather than general feature background. Use None
  for work without a defect.
- Impacted services: affected repository names only, without owner/organization
  prefixes, URLs, or workflow descriptions. Resolve the name from origin/GitHub,
  falling back to the repository root name.
- Core modules/sections: affected workflows or feature areas, not a file log.
- UI, DB, indexes, environment, infra: brief evidenced changes, or None when
  verified unaffected. Include other services only when evidence shows changes.

Use simple sentences with one action or idea each, and explain events in the
order they happen. Make the opening understandable on its own. Use familiar
words for the user action, what the app does next, and the result; introduce
internal terms only when needed. Do not repeat the same fact or compress several
steps into one sentence. Each explanation is a paragraph on its existing label
line, without numbered steps or nested bullets; retain the template's field labels.

Keep answers on their bullet's label line, a blank line between sections and
before bullets, and no blank lines between bullets. Add `- Important notes:`
after Root cause only for a supported constraint, risk, or follow-up that matters
to the reviewer. Routine passing format/lint/type checks belong in the completion
report. Replace every placeholder for a completed draft; a focused refinement
reports unrelated gaps without rewriting them. Claim checks passed only when
observed. Do not create a separate summary file.
