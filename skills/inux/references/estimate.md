# Estimate

Use only the requested section. Bare `/inux estimate` fills an empty template or
improves an existing draft. Read [be-human.md](be-human.md) once for writing.
Publication requires explicit intent.

## Work efficiently

Batch independent reads/searches. Reuse current evidence; read only missing or
changed context. Invoke known tool paths directly. Keep routine checks quiet.

## Generate

Read ticket.md and estimate.md together; use its initialized structure without
rereading the template asset. Estimate the requested end state, including completed
work. Search the affected flow once, then read relevant code and coverage together.
Stop when scope, effort, and material risks are understood. Further searches
must answer a question that could change the estimate. Check project scoring
guidance once; otherwise use the fallback below. Do not run tests.

Confirm expected behavior before choosing a fix. If ambiguity changes scope,
check relevant Jira comments/attachments using the shared Jira execution rule;
otherwise use local evidence. Record unresolved assumptions in Risk & Unknowns;
keep proposed fixes conditional. Ask only when a useful estimate is blocked.

Fill the draft, preserving supported developer decisions. Validate, recheck the
active ticket, save only estimate.md, and verify the saved file once.

## Improve

Read ticket.md and the complete estimate together. Refine requested parts;
wording edits preserve supported scores, scope, and SP without code investigation.
Inspect new evidence only when needed. Change final SP only when unsupported by
the breakdown or evidenced conventions; explain why.
Validate changed fields and dependent totals; report unrelated incomplete fields.
For empty templates, use Generate. Recheck the ticket, save only estimate.md,
and verify once. Do not reset tasks, fetch Jira for wording edits, or run tests.

## Explain

Use the saved breakdown and current evidence; inspect more only when needed.
Flag contradictions. Do not edit or publish.

## Publish

Read ticket.md and estimate.md together; validate without rewriting.
Report missing fields. Use private acli:
`jira workitem comment list --key TICKET-ID --json --paginate`.
Check duplicates and clarifications. If evidence contradicts the draft, stop and
ask whether to revise; leave it unchanged. Report identical comments without
posting; never replace unrelated comments.

Convert with private ix: `adf ABSOLUTE_ESTIMATE_PATH TEMPORARY_OUTPUT.json`.
Inspect ADF and recheck ticket ID and saved content together before posting.
Use private acli:
`jira workitem comment create --key TICKET-ID --body-file ABSOLUTE_ADF_PATH --json`.
Consult only exact subcommand help when needed. Never rebuild ADF manually.
Inspect remote state before retrying uncertain writes. Remove temporary output.
Report the returned comment/ticket URL; look it up only if missing. Preserve task files.

## Writing and validation

Preserve six rows, order, fields, and Markdown structure. Scores are nonnegative
integers. Use the documented team scale, otherwise 0 for no meaningful work and
1–3 for increasing complexity. Explain zeros. Keep reasons in their categories,
uncertainty in Risk & Unknowns, and avoid double-counting.

Each reason is one plain sentence, usually 10–20 words. Omit code walkthroughs.
Overall Reason covers the main effort and risk in one or two sentences, usually
under 40 words. Chat replies need ticket, SP, file or URL, and material assumptions;
expand when asked.

Total Complexity Score equals the six scores. Final Estimate is numeric, ends
in SP, and is supported by the work; never invent a score-to-points formula.
No blank fields, placeholders, or unsupported claims remain.
