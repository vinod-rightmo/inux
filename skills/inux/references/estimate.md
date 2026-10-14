# Estimate

Use only the section matching the request. Read [be-human.md](be-human.md) for
writing. Bare `/inux estimate` fills the initialized template, otherwise improves
existing content. Publication requires explicit intent.

## Generate

Read [be-human.md](be-human.md), ticket.md, and the initialized estimate.md together.
Estimate the ticket's requested end state. Inspect the likely affected code and
nearby verification; widen the search only when a material uncertainty could
change the estimate. Reuse relevant evidence already verified in this chat.
Do not fetch Jira or inspect unrelated repository areas for a complete local ticket.

Fill the existing estimate, preserving supported developer decisions and the
[template](../assets/estimate.md). Apply [Validation](#validation).
Save only estimate.md after checking the active ticket again. Do one final
consistency check and report the file. Do not run project tests merely to estimate
future work; inspect existing coverage when it affects the estimate.

## Improve

Read [be-human.md](be-human.md), ticket.md, and the complete estimate.md. For
an unfilled template, report that generation is needed rather than resetting it.
Refine the requested parts. Preserve accurate reasons, scores, scope decisions,
and final SP. Inspect additional code only when the requested change needs new
evidence; do not repeat the generation investigation for a wording edit.

Use [Validation](#validation) to check changed content and dependent
totals. Report unrelated incomplete fields without filling them or claiming the
whole draft is complete. Change a developer's final SP only
when the breakdown or an evidenced convention makes it unsupported; explain why.
Recheck the active ticket, save only estimate.md, and verify it once. Do not
initialize a task, fetch remote details, or run tests for prose refinement.

## Explain

Read ticket.md and estimate.md. Explain the requested scores, reasons, scope,
or final SP in plain language, using [be-human.md](be-human.md).
Use the saved breakdown and evidence already available. Inspect more code only
if the question cannot be answered from that evidence. Flag a contradiction
instead of silently fixing it. Do not write files, run tests, or publish.

## Publish

Require explicit publishing intent. Read ticket.md and saved estimate.md;
apply [Validation](#validation) without rewriting the draft.
If incomplete, report the missing fields.

Use private acli to inspect existing comments. If identical content exists,
report it without posting a duplicate; never replace an unrelated comment.
Convert the saved Markdown with ix adf to temporary ADF JSON and inspect it.
Use the installed acli comment-create body-file input with the exact ticket key;
check its help for syntax. Do not reconstruct ADF by hand.

Recheck ticket ID and saved content immediately before posting. Inspect remote
comments before retrying an uncertain result. Remove temporary output and report
the comment/ticket URL. Leave task files unchanged.

## Validation

Preserve the template's six rows, their order, fields, and Markdown structure.
Each score is a nonnegative integer with a specific reason. Use the documented
team scale, otherwise 0 for no meaningful work and 1–3 for increasing complexity;
explain zeros. Keep each reason in its category, uncertainty in Risk & Unknowns,
and avoid double-counting distinct effort.

Total Complexity Score equals the six scores. Final Estimate is numeric and
ends in SP, supported by the work and any evidenced team convention. Do not
invent a score-to-points formula. Overall Reason explains the dominant effort
and risk. No blank fields, template placeholders, or unsupported claims remain.
