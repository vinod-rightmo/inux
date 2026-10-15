# PR readiness

Preparation and publication must pass the Git gate before implementation review
and drafting. Verification and the final recheck must also pass before saving or
publishing. Stop dependent work on failure and preserve drafts. Do not auto-commit,
rename/rewrite history, stash, or discard work to make readiness pass.

## Git gate

1. Resolve the project, active ticket from ticket.md's first heading, applicable
   project `AGENTS.md` conventions/check commands, attached current branch,
   GitHub repository, and push remote. Confirm ticket.md/pr.md exist; defer their
   full contents until the gate passes. Require the developer's destination from
   this request or an earlier explicit choice for this PR. Accept
   `/inux pr target BRANCH`, `/inux pr publish target BRANCH`, or clear natural
   language. If missing, ask "Which destination branch should this PR target?"
   before Git checks or draft changes. Do not infer it from defaults, `dev`, or
   an existing PR. Stop on a supplied ticket-ID mismatch.
2. Inspect local state once with
   `git status --porcelain=v1 --untracked-files=all --ignore-submodules=none`,
   branch, HEAD, remotes, and operation state. Require a clean index/tree,
   including untracked files/submodules, and no merge/rebase/cherry-pick/revert.
   Use resolved Git paths for operation markers, including worktree checkouts.
   Validate the branch with [team conventions](#team-conventions). Reuse this
   state until the final recheck; do not add duplicate full/short status output.
3. Fetch the target's remote ref and use that ref as TARGET_REF. Require a task
   branch distinct from the target. Missing/failed refs block rather than using
   another base. Require a shared ancestor with `git merge-base TARGET_REF HEAD`;
   the target need not be an ancestor of HEAD, so no `--is-ancestor` precheck.
4. Enumerate only PR commits with `git rev-list --reverse TARGET_REF..HEAD`.
   Require a nonempty range, retaining each SHA and full stored message. Include
   merge commits; do not filter by ticket/author. Commits already reachable from
   the target are inherited history, outside this task's convention check.
   Validate every message with [team conventions](#team-conventions), batching
   if needed. Map failed 1-based positions to SHAs. Every check must pass.

Record branch, HEAD, target revision, merge base, checked SHAs/messages, and
results for reuse. Structural and present-tense checks happen here; assess
message accuracy against implementation when reviewing the diff after the gate.
Required-command failures must propagate: do not mask them with `|| true` or a
trailing `echo`. Batch independent reads/checks when supported; preserve each
exit status. Invoke Git, private ix, and private gh directly; no persisted helper,
hidden state file, or hook is required.

## Team conventions

Use these Inux/QR defaults unless the user or project explicitly overrides them.
Nonconforming history is not an override. Escape the exact uppercase active
ticket as a regex literal in place of TICKET. Do not ask for regexes when these
defaults apply; an override replaces only its stated rule.

Branch: `<type>/<TICKET-ID>-<lowercase-kebab-description>`. Require
`git check-ref-format --branch NAME` and this Go regex:

```regex
^(feature|bugfix|hotfix|chore|refactor|perf|test|release)/TICKET-[a-z0-9]+(-[a-z0-9]+)*$
```

Commit: one line, `<TICKET-ID>: <Capitalized present-tense summary>`, with exactly
one active-ticket token and no body/duplicate prefix. Add/Fix/Update/Remove are
examples, not a closed verb list. Structure pattern:

```regex
^TICKET: [A-Z][^\r\n]+$
```

Use the absolute private ix; quote/pass each value as a separate argument:

```text
ix check-branch PATTERN BRANCH
ix check-commit PATTERN FULL_MESSAGE [FULL_MESSAGE...]
```

Read each stored message once with `git show -s --format=format:%B SHA`. Capture
that value and pass it directly to ix, removing at most one terminal LF while
preserving other whitespace. Do not manually reproduce the message or use shell
command substitution that strips all trailing LFs. No preliminary byte dump is
needed; diagnose unexpected validator failures afterward.

Checks match the entire value: `{"ok":true}`/exit 0 means success; JSON error and
nonzero exit mean failure. Mismatches report failed 1-based argument positions.
Check exactly one active-ticket token and present-tense meaning as well. Use an
explicit convention override's validator or an equivalent check and report it.

## Verification

Review the full merge-base-to-HEAD diff for completed ticket work and accurate
commit messages. Incomplete work, material unresolved correctness concerns, and
failed required checks block readiness; a clean tree alone proves none of these.

Reuse verification only while the checked implementation and relevant inputs
remain current. Run missing project-required checks using commands already
supplied in project instructions. If a command needs translation or is missing,
perform one targeted lookup in the appropriate manifest/Makefile; reuse its result.
Probe runtime health only to resolve an actual failure. Request required host
access on the first invocation. Independent read-only checks may run together
and overlap evidence review when supported. Collect each exit status and report
only observed successes.

## Final recheck

Record verification, then recheck branch, HEAD, target revision, active ticket,
operation state, and tree once before saving/publishing. Unchanged revisions
preserve the checked range and evidence. Restart affected checks only if relevant
state changed. Publication also confirms the saved draft is unchanged.
