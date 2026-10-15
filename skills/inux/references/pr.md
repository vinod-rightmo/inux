# Pull requests

Bare `/inux pr` prepares `.inux/pr.md`. Publish only on explicit request.
Use the requested action below; PR inputs are ticket.md and pr.md, not estimate.md.
Require both files and a matching active ticket; never reset the task here.
Preserve accurate developer edits and the draft's labels/order. Read the
[template](../assets/pr.md) only when required labels are missing.

For preparation/publication, show a compact chat checklist: Git validation,
evidence/verification, save/publication. Use Markdown or an available plan tool;
update at phase changes, identify blockers, and keep it out of the PR document.
Stop dependent work on errors. Do not commit, stash, discard work, rename branches,
or rewrite history to make checks pass. Batch independent reads/checks, preserve
each exit status, and reuse unchanged instructions and evidence.

## Git gate — prepare and publish

Complete this before reading full task files, reviewing source, or drafting.

1. Resolve project, active ID from ticket.md's first heading, project `AGENTS.md`
   rules/checks, current branch, GitHub repository, and push remote. Require the
   developer's target from this request or an earlier explicit choice for this PR
   (`/inux pr target BRANCH` or `/inux pr publish target BRANCH`). If absent, ask
   “Which destination branch should this PR target?” before Git checks or edits.
   Never infer it from defaults or an existing PR. Stop on a supplied-ID mismatch.
2. Inspect branch, HEAD, remotes, operation state, and
   `git status --porcelain=v1 --untracked-files=all --ignore-submodules=none` once.
   Require a clean index/tree including untracked files/submodules, attached HEAD,
   and no merge/rebase/cherry-pick/revert. Resolve operation-marker paths through
   Git for worktree compatibility. Check the branch against Naming below.
3. Fetch the target's remote ref and use that same ref as TARGET_REF throughout.
   Require a task branch distinct from the target and a shared ancestor using
   `git merge-base TARGET_REF HEAD`. Failed/missing refs block; never substitute
   another base or require the target to be an ancestor of HEAD.
4. Read PR commits once with
   `git log --reverse --format='%H%n%B%x00' TARGET_REF..HEAD`.
   Require a nonempty range; retain SHAs and full messages, including merge
   commits. Never filter by author/ticket or check inherited target history.
   Check every message against Naming. Report failing SHAs and violated rules.

Record branch, HEAD, target revision, merge base, checked messages, and results.
The agent checks naming directly using Git output; no ix validator or generated
validation script. Required-command failures must propagate without `|| true`
or a trailing command masking the status.

## Naming

Explicit user/project rules override these Inux/QR defaults only where stated;
old nonconforming history is not an override. Use project validators when required.

- Branch: `<type>/<TICKET-ID>-<lowercase-kebab-description>`. Run
  `git check-ref-format --branch NAME`. Require the exact uppercase active ticket,
  type feature/bugfix/hotfix/chore/refactor/perf/test/release, and a nonempty
  description of 1–6 lowercase letter/digit words separated by single hyphens.
  Count only description words, excluding the type and ticket ID.
- Commit: one line, `<TICKET-ID>: <Capitalized present-tense summary>`. Require
  the exact active ticket once, colon and one space, then a meaningful summary
  starting with an uppercase letter. Reject bodies, extra prefixes, and duplicate
  tickets. Add/Fix/Update/Remove are examples, not a closed verb list.

Inspect full messages, preserving spaces/line breaks. Git record separators and
one normal terminal LF are not a body; do not trim other whitespace, reconstruct
messages, or reread each commit. Assess message accuracy with the diff after the gate.

## Prepare

After the gate, read full ticket.md/pr.md, use project rules and checked messages,
and review the full merge-base-to-HEAD diff. Implementation claims require code
evidence; ticket intent and commit messages alone do not prove them. Resolve only
material gaps: name the missing fact, follow exact imports/symbols into direct
dependencies, and read relevant ranges. Stop when claims are supported.

Write one concise line under `## Title`, defaulting to
`<TICKET-ID>: <Capitalized action summary>` unless overridden. Fill Description
using Summary below and [writing guidance](be-human.md), keeping required sections
and accurate developer notes. Complete Verification and recheck before saving.
A successful write completes preparation; no routine readback or byte audit.
Report the path and observed checks. Preparation does not push or publish.

## Summary

Keep the technical summary once inside Description; never create a separate file.
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
blank lines between field bullets. Add `- Important notes:` after Root cause only
for a supported constraint, risk, or follow-up that matters to the reviewer.
Replace all placeholders in a completed draft. Routine passing checks belong in
the completion report; claim success only when observed.

## Verification and recheck

Require completed ticket work, accurate commit messages/summary, no material
unresolved correctness concerns, and passing project-required checks. Reuse checks
only while relevant code/inputs are unchanged. Run missing checks from project
instructions; use one targeted manifest/Makefile lookup if needed. Probe runtime
health only for an actual failure and request required host access on first use.
Independent checks may overlap read-only review/composition; collect every result
before saving or publishing.

Recheck branch, HEAD, TARGET_REF revision, active ticket, operation state, and tree
once immediately before saving/publishing. Restart affected checks only when
relevant state changed. Publication also requires the saved draft to be unchanged.

## Publish

Run the Git gate, then read full ticket.md and the saved draft. Require matching
ticket, a completed single-line Title, a complete Description, and no placeholders. If
incomplete, request preparation/refinement; do not draft during publication.
Compare the saved summary/messages with the full diff, resolve material gaps,
and report stale content instead of rewriting it. Complete Verification.

Use private gh to check authentication and open PRs for this exact repository/head.
Clarify multiple matches or a differing existing target before mutation. After the
final recheck, push without force and verify the remote SHA equals checked HEAD.
Create a draft PR with explicit base/head and self-assignment, or update the
existing open PR. Pass `--repo` explicitly; preserve other assignees when adding
the authenticated user. Use saved Title and a temporary UTF-8 body file containing
only Description, preserving Markdown exactly. Inspect the result, report its URL,
and attach it to this chat when supported. Report partial success; inspect remote
state before retrying uncertain writes. Remove temporary output; keep pr.md unchanged.

## Refine or explain

No target, fetch, or readiness pass is needed. Read ticket.md/pr.md, preserve
unrequested sections, and reuse verified evidence. Check branch/HEAD before
relying on earlier implementation evidence; inspect relevant committed changes
only for a material gap. Apply Summary/writing guidance as needed.

For Title, edit only its concise line using the default format above. For
Description, edit only the requested content; preserve accurate notes and report
unrelated incomplete fields. Recheck the active ticket before saving; other
sections remain byte-for-byte unchanged. No commits, pushes, or publication.
For an explanation, answer from the draft and verified evidence without writing
files. Distinguish previous readiness from current state; do not imply fresh checks.
