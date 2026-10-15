# Commit changes

`/inux commit changes` creates a local commit for the requested changes. The
request authorizes staging and committing; no extra confirmation is needed when
scope and convention are clear. Accept explicit paths or a commit message.

1. Use the explicit/current checkout and applicable repository conventions.
   Batch current branch/HEAD, porcelain status, operation state, and staged and
   unstaged diffs. Stop on detached HEAD, unresolved conflicts, or an active
   merge/rebase/cherry-pick/revert. If there are no changes, report that and stop.
   Read changed untracked files that belong to the requested scope; diff does
   not show their contents. Do not read entire files when the diff is sufficient.
2. Select the requested files, otherwise changes clearly belonging to the active
   task. Ask only if scope is ambiguous, files mix unrelated edits, or unrelated
   staged changes would enter the commit. Preserve unrelated work and the index;
   do not silently unstage it. Exclude task drafts, credentials, and generated
   artifacts unless explicitly intended and allowed by repository policy.
3. Draft one accurate, concise message from the diff, reusing current task context.
   If needed, read only `.inux/ticket.md`'s heading and Summary. Repository/user
   conventions override the Inux default, including Conventional Commits in the
   Inux repository itself. Do not infer an override merely from old history.
   Under the default, require exactly one line:
   `<TICKET-ID>: <Capitalized present-tense summary>`, with the active ticket
   exactly once, followed by a colon and one space. The summary starts with an
   uppercase letter and describes the change in present tense. Check the proposed
   message directly; reject bodies, duplicate ticket tokens, and extra prefixes.
   Stop on a supplied-ticket mismatch. For an override, use its validator if
   required or check its stated rules directly. No ix call or generated validation
   script is needed. Validate only the proposed message, not historical commits.
   No branch naming or PR target validation is required for this action.
4. Reuse observed verification while the relevant changes are unchanged. Run
   missing repository-required checks; an observed failure blocks the commit.
   Let normal commit hooks run; do not disable them. No additional test suite is
   required solely because a commit was requested. Stage explicit selected paths
   with `git add -- PATH...`, then inspect `git diff --cached --stat` and
   `git diff --cached --check`. Read the staged patch only for changes not already
   reviewed or if staging/hooks changed content. Require a nonempty intended diff.
5. Run `git commit -m MESSAGE`, quoting safely, or use a UTF-8 `--file` for a
   repository convention requiring a body. After success, get the resulting SHA,
   full message, and remaining status in one call. Report SHA, message, and any
   remaining changes briefly. On hook failure, report it and preserve the state;
   inspect before any retry and do not blindly retry an uncertain commit.

Keep local evidence within this action; repeat affected checks only if the state
changes. No PR/CLI references, tool/version probes, full history, fetches, pushes,
amend, reset, stash, task initialization, or PR-draft changes are needed.
