# Technical summary structure and readability

Write technical summary content directly under `## Description` in `pull-request.md` for
a busy reviewer. Inspect enough evidence to
understand the change, then include only what the reader needs to understand
the implementation, cause, impact, and meaningful constraints.

## Fields

Include these labels and fields in order, filling every required field with
supported content or `None` when it does not apply:

```markdown
Branch name:
<exact current Git branch>

Implementation Summary:

- Fix / implementation: <changed behavior and implementation approach>
- Root cause: <defect's causal mechanism, or None for work without a defect>

Code Layer Impact:

- Impacted services: <exact affected repository names>
- Core modules/sections: <affected modules or workflows>

UI Impact:

- New/updated components, layouts, or pages: <actual UI impact or None>

DB Impact:

- New/updated tables: <actual table, schema, or migration changes or None>
- Indexing added/updated: <actual index changes or None>

Infra Impact:

- Env changes: <actual environment changes or None>
- Cloudflare/S3/Redis or other infra changes: <actual infra changes or None>
```

For the current service, use the repository name from its GitHub or `origin`
remote, falling back to the Git root directory name. List other services only
when repository evidence shows they changed. Verify all impact fields against
code, tests, configuration, migrations, and infrastructure; the ticket explains
intent but does not prove completed work. Mention verification only when it was
observed. Replace every placeholder before finishing.

## Structure

- Preserve the section labels and field order above.
- Put the exact branch value on the line immediately below `Branch name:`.
  Use a Markdown hard break after the label to preserve the rendered line break.
- Keep each bullet's answer on the same line as its label, for example:
  `- Root cause: The selected color remained in both palettes.` Natural visual
  wrapping is fine; do not put the answer in a separate paragraph or nested list.
- Leave one blank line between sections and between a section label and its
  bullets. Do not leave blank lines between bullets or their labels and answers.
- Add `- Important notes: ...` after Root cause only for a meaningful, supported
  constraint, exception, risk, uncertainty, or required follow-up. Otherwise omit
  the entire bullet. Bold a critical action or exception when useful.

## Detail level

- Fix / implementation: normally one or two short sentences describing the
  changed behavior and essential implementation approach.
- Root cause: normally one short sentence explaining why the defect happened.
- Impact fields: brief repository/module names or a short description of the
  affected interface, database, or infrastructure.
- Use plain English while keeping useful technical terms. Include code
  identifiers, paths, exact values, or breakpoint details only when they help
  explain the change; avoid code walkthroughs and research history.
- Give each field a purpose: Fix explains what changed, Root cause explains why,
  and UI Impact identifies the affected interface. Do not repeat the fix there.
- Keep notes brief and grounded in evidence; do not add hypothetical concerns
  merely because something was not changed. Preserve material uncertainty.
- Allow more detail when the change requires it. Before finishing, remove
  repetition and details that do not help the reader understand or review it.
