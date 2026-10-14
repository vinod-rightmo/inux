# Pull requests

Use only the section matching the request. Bare `/inux pr` prepares a draft.
Task initialization already creates pr.md from assets/pr.md. Preparation fills
it; title/description requests fill or refine only their section.
Readiness checks apply to preparation and publication; a wording edit does not
trigger them. Publishing requires explicit intent and the reviewed saved draft.

## Prepare

Read ticket.md, the initialized pr.md, and [Readiness checks](#readiness-checks). Complete all
readiness checks before filling or changing the initialized pr.md. On failure, preserve the
draft and report the blocker; do not fix Git state automatically.

Use the checked evidence, [be-human.md](be-human.md), and
[Technical summary](#technical-summary). Fill the initialized `.inux/pr.md`
from the [PR template](../assets/pr.md): one concise line under `## Title`,
with the technical summary inside `## Description`. Preserve accurate content
rather than resetting the draft to the template. Title defaults to
`<TICKET-ID>: <Capitalized action summary>` unless user/repository rules differ.
Keep required PR-template sections. Preserve accurate existing developer edits.

Recheck branch, HEAD, PR target/baseline revisions, ticket, and tree before
saving. Report the path and readiness outcome. Preparation does not push or publish.

## Title

Read the initialized ticket.md and pr.md. Change only the requested section;
there is no need to run preparation first for a title/description edit.

Use [be-human.md](be-human.md). Write one concise line under `## Title` describing
the delivered change. Default to `<TICKET-ID>: <Capitalized action summary>`
unless explicit user/repository rules differ. Preserve accurate intent. For a
wording edit, use verified draft evidence; inspect the relevant
committed change only when needed to establish the title's accuracy. Check
current branch/HEAD before relying on earlier implementation evidence. Do not rerun
all readiness checks for an existing draft refinement.

Recheck the active ticket and save only Title. Description and every other part
of the file remain byte-for-byte unchanged. No commits, pushes, or publication.

## Description

Read the initialized ticket.md and pr.md. Change only the requested section;
there is no need to run preparation first for a title/description edit.

Use [be-human.md](be-human.md) and
[Technical summary](#technical-summary). Change only the
requested Description content. Preserve accurate developer wording and notes;
correct stale claims, fill relevant gaps, and remove repetition. Use verified
branch/diff evidence already available; inspect the affected changes only when
new evidence is needed. Check current branch/HEAD before reusing implementation
evidence. Do not rerun readiness or fetch for a wording edit. Report unrelated
incomplete fields instead of silently expanding a focused edit.

Recheck the active ticket and save only Description. Title and other sections
remain byte-for-byte unchanged. Keep the summary once inside pr.md; do not
create another summary file or claim unobserved verification.

## Explain

Read ticket.md and the initialized pr.md. Answer the requested question using
[be-human.md](be-human.md), the saved draft, and available verified evidence.
Inspect relevant code or Git state only when needed to answer accurately.
Distinguish a previous readiness result from the current state; an explanation
is not a fresh readiness certification. Do not write files, rerun the full
preparation workflow, push, or publish.

## Publish

Require explicit publishing intent. Read ticket.md and saved pr.md. Require a
matching ticket, completed single-line Title, and complete Description using
[Technical summary](#technical-summary). If missing or
incomplete or containing template placeholders in either section, request
preparation/refinement; do not draft and publish together.

Rerun all [Readiness checks](#readiness-checks) checks. Compare the saved summary with the
checked changes; report stale content instead of rewriting it during publication.
Check private gh authentication and open PRs for this exact repository/head.
Multiple matches need clarification.

Require an existing PR to target the resolved base; resolve a differing target
before mutation. Recheck branch, HEAD, target/baseline revisions, ticket, clean
tree, and saved content. Push the task branch without force and verify the remote
branch SHA equals checked HEAD. Create a draft PR with explicit base/head and
self-assignment, or update
the existing open PR for this head. Use explicit `--repo` for gh PR operations
and preserve other assignees when adding the authenticated user. Use the saved
Title and a temporary UTF-8 body file containing only Description, preserving its Markdown exactly.

Inspect the result, report its URL, and attach the PR to this chat when supported.
If push succeeds but PR creation fails, report both. Inspect remote state before
retrying an uncertain write. Remove temporary output; leave pr.md unchanged.

## Readiness checks

Run these for preparation and publication. Stop on failure and preserve drafts.
Do not auto-commit, rename/rewrite history, stash, or discard work. A clean tree
does not prove completion or passing tests; incomplete work and failed required
checks also block readiness.

### Team conventions

Use these established Inux/QR rules unless the user or repository explicitly
overrides them. Nonconforming history is not an override. Do not ask for regexes
when these defaults apply. Resolve the exact active ticket from ticket.md and
escape it as a regex literal in place of TICKET below.

Branch: `<type>/<TICKET-ID>-<lowercase-kebab-description>`, with the task's exact
uppercase ticket. Types: feature, bugfix, hotfix, chore, refactor, perf, test,
release. Require `git check-ref-format --branch NAME` as well as this Go regex:

```regex
^(feature|bugfix|hotfix|chore|refactor|perf|test|release)/TICKET-[a-z0-9]+(-[a-z0-9]+)*$
```

Commit: exactly one line, `<TICKET-ID>: <Capitalized present-tense summary>`.
The active ticket occurs exactly once; no body or duplicate ticket prefix.
Add/Fix/Update/Remove are examples, not an allowed-verb list. Structure pattern:

```regex
^TICKET: [A-Z][^\r\n]+$
```

Example: `OQG3-3776: Fix QGC home step number centering`. Apply ix check-commit
to the full stored message, not just `%s`. For example, read
`git show -s --format=format:%B SHA` and remove at most one terminal LF from the
stored message; preserve other newlines/whitespace. Regex checks structure;
also check exactly one occurrence of the task's ticket token, present-tense
meaning, and implementation accuracy. Do not claim regex proves those meanings.
An explicit override replaces only its stated rule; report the override and
use its validator or an equivalent check without weakening other requirements.

### Checks

1. Resolve the project, GitHub repository, active ticket, attached current branch,
   push remote, PR target, and convention baseline. User/repository rules take
   precedence; team defaults are target `dev` and baseline remote `main`.
   Require a task branch distinct from both. Fetch the relevant remote
   refs; missing/failed refs block instead of silently substituting another base.
2. Run `git status --porcelain=v1 --untracked-files=all --ignore-submodules=none`.
   Require a clean index/tree, including untracked files/submodules, and no
   merge/rebase/cherry-pick/revert in progress. Do not hide tracked task changes.
3. Check the branch using the resolved default/override regex and Git ref check.
4. Require shared ancestry with the convention baseline. Enumerate every commit
   in `REMOTE/BASELINE..HEAD`, retaining each SHA and full stored message.
   Require a nonempty range and validate every message as above with ix check-commit,
   passing separate arguments and batching if needed. Do not omit merge commits.
   Map failed 1-based positions to SHAs; every batch and semantic check must pass.
5. Inspect commits and the merge-base-to-HEAD diff against the actual PR target,
   separately from the convention baseline. Require a nonempty target diff,
   ticket-related completed work, and required verification. Record branch, HEAD,
   both remote revisions, and checked commits. Recheck these, ticket, and tree
   before saving/publishing; restart affected checks if changed.

Run Git, ix, and gh independently. No hidden state files or hooks are needed.

## Technical summary

Fill the technical-summary fields from the [PR template's Description](../assets/pr.md#description)
directly under `## Description` in `.inux/pr.md`. Keep Title unchanged for a
description-only edit. Preserve the summary's section labels and field order.
Use implementation evidence for delivered behavior; the ticket supplies intent.
Inspect affected code/configuration only as needed to support the fields. Missing
evidence is not proof of no impact: resolve material gaps or report them.

- Branch: exact current branch on the line below `Branch name:`; retain the
  label's two trailing spaces for the Markdown hard break.
- Fix / implementation: one or two sentences about behavior and approach.
- Root cause: verified causal mechanism, or None for work without a defect.
- Services/modules: exact affected repositories and workflows, not a file log.
  Resolve the current repository from origin/GitHub, falling back to root name.
- UI, DB, indexes, environment, infra: brief evidenced changes, or None when
  verified unaffected. Include other services only when evidence shows changes.

Keep answers on their bullet's label line, a blank line between sections and
before bullets, and no blank lines between bullets. Add `- Important notes:`
after Root cause only for a meaningful supported constraint, verification, risk,
or follow-up. Retain required PR-template sections without duplicating the summary.
For a completed summary, replace every placeholder. A focused refinement reports
unrelated gaps without silently rewriting them. Preserve accurate developer notes
and mention tests as passed only when observed. The summary appears once, inside Description.
