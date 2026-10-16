# Technical summary

Use Summary content below for both `.inux/technical-summary.md` and the
Description in `.inux/pr.md`. For writing, read [writing guidance](be-human.md) once.
PR actions use the matching program in [Inux routing](../SKILL.md).

For `/inux ts generate`, `/inux ts improve`, or `/inux ts publish`, execute
[ts.flow](ts.flow), which USEs the selected mode file. `ts create` aliases `ts generate`;
`/inux technical summary [request]` remains an alias.
Bare requests generate missing/empty summaries and otherwise improve them.
Publication requires explicit intent. Do not automatically synchronize the files.

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
