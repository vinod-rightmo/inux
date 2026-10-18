# Write estimate

Draft or refine `estimate.md` for the work requested by the sibling `ticket.md`.
Read the complete existing estimate first. Preserve accurate scores, scope
decisions, reasons, and final story-point judgments; fill an empty template from
evidence rather than discarding useful existing content.

## Evidence

1. Read the whole ticket, including acceptance criteria and relevant
   attachments.
2. Inspect nearby implementations, likely affected modules, integrations,
   configuration, tests, and delivery conventions in the repository.
3. Estimate the requested end state, not only code that happens to exist on the
   current branch. Record only unknowns that could materially change delivery.

## Complete the estimate

Keep the six rows, their order, headings, and Markdown structure. Replace vague
reasons with concrete work, dependencies, or risks. Keep each reason in its
correct category and avoid counting the same work repeatedly unless it creates
distinct effort. Put delivery uncertainty in Risk & Unknowns rather than
inflating other rows. Reconcile contradictions between scores, reasons, total,
overall reason, and final estimate. Change an existing final story-point value
only when the breakdown or an evidenced team convention makes it unsupported.

Score every existing row with a non-negative integer and a concise reason tied
to evidence:

- **Implementation Complexity:** domain behavior, control flow, algorithms,
  compatibility, and implementation difficulty.
- **Surface Area:** modules, services, flows, consumers, and repositories that
  must change.
- **UI & Localization Complexity:** interaction, layout, copy, accessibility,
  responsiveness, and translation work.
- **Data / Config / Integration Complexity:** schemas, migrations, APIs,
  third parties, flags, and configuration.
- **Verification & Support Complexity:** automated and manual coverage,
  regression surface, rollout, observability, and support needs.
- **Risk & Unknowns:** unclear requirements, unfamiliar systems, external
  dependencies, compatibility risk, and likely rework.

Use the project’s documented scoring convention when one exists. Otherwise use
`0` for no meaningful work and `1`–`3` for increasing complexity. Explain every
zero; do not add effort to categories merely because they exist.

Set `Total Complexity Score` to the exact sum. Set `Final Estimate` to a numeric
story-point value followed by `SP`, using the team’s convention when evidenced.
Do not invent a score-to-points formula. Write a short `Overall Reason` that
explains the dominant work and risk rather than repeating the table.

## Validate

- All six rows contain an integer score and specific reason.
- The total exactly equals the six scores.
- `Final Estimate` is numeric, ends in `SP`, and is supported by the breakdown.
- No blank fields, `TODO`, unsupported claims, or structural changes remain.
