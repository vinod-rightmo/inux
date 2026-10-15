# Create a branch

`/inux branch create` creates and switches to a local branch. Accept an explicit
name or starting ref in the request. Otherwise use current HEAD; a PR destination
does not change this starting point. This command authorizes creation and checkout.

1. Use the explicit/current checkout and applicable repository conventions.
   Batch local branch/HEAD, porcelain status, and Git operation-state inspection.
   Stop for unresolved conflicts or an active merge/rebase/cherry-pick/revert.
   Uncommitted changes may carry onto the new branch; preserve them.
2. Use the supplied name, or derive one from the known active task. If needed,
   read only `.inux/ticket.md`'s heading, Type, and Summary. Default to
   `<type>/<TICKET-ID>-<lowercase-kebab-description>`; types are feature, bugfix,
   hotfix, chore, refactor, perf, test, release. A Bug maps to bugfix; a Story or
   new capability maps to feature. Keep the description to 1–6 hyphen-separated
   words, excluding the type and ticket ID. Describe the change concisely rather
   than repeating the issue sentence. Use repository/user overrides, including
   branch prefixes. Ask only for missing facts needed to choose a name. If a
   supplied ticket disagrees with the active ticket, stop before changing Git.
3. Run `git check-ref-format --branch NAME`, then check the name directly
   against the rule in step 2: an allowed type, the exact uppercase active ticket,
   and a description of 1–6 lowercase letter/digit words separated by single
   hyphens. Count only description words, excluding the type and ticket ID.
   Follow explicit repository/user overrides and use a project-supplied
   validator when required. Stop on a mismatch; no ix call or generated validation
   script is needed. If NAME is already current, report that and stop.
   If another local branch has that name,
   report it; do not replace it or silently switch to it.
4. Run `git switch -c NAME` (append the explicit START_REF only when supplied).
   Require success; its output completes verification. Report the branch briefly.

Pass shell arguments safely and separately. No PR/CLI references, tool probes,
tests, fetches, pushes, task initialization, stash, or history changes are needed.
If an explicit start ref is unavailable locally, report it without substituting
another ref. Stop on errors; do not force checkout or retry uncertain creation.
