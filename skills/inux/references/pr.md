# Pull requests

Load only the requested action and its required references. Bare `/inux pr`
prepares the initialized `.inux/pr.md`; publishing requires explicit intent.

| Request | Action |
|---|---|
| Prepare/create/fill pr.md | [prepare](pr/prepare.md) |
| Publish/create/update the remote PR | [publish](pr/publish.md) |
| Refine Title or Description, or explain the change | [refine](pr/refine.md) |

For preparation/publication, load [readiness](pr/readiness.md) with the action
reference; defer summary/writing guidance until the Git gate passes. Refinements
load only the references required by their selected section.
Batch independent reference reads; reuse unchanged instructions and evidence.
PR task inputs are ticket.md and pr.md; estimate.md is not an input.

Show a compact chat checklist: Git validation, evidence/verification, and
save/publication. Use an already available plan/todo tool or Markdown; do not
discover a tool just for progress. Update at phase changes from observed results,
identify blockers, and keep the checklist out of the PR document. Focused edits
and explanations show only their relevant steps.
