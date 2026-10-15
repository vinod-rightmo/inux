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
   new capability maps to feature. Use repository/user overrides, including
   branch prefixes. Ask only for missing facts needed to choose a name. If a
   supplied ticket disagrees with the active ticket, stop before changing Git.
3. Run `git check-ref-format --branch NAME`. For the default naming rule, use
   absolute private ix with `check-branch PATTERN NAME`, escaping the exact active
   ticket as a Go-regex literal in place of TICKET:

   ```regex
   ^(feature|bugfix|hotfix|chore|refactor|perf|test|release)/TICKET-[a-z0-9]+(-[a-z0-9]+)*$
   ```

   Apply an explicit override's validator instead when provided. If NAME is
   already current, report that and stop. If another local branch has that name,
   report it; do not replace it or silently switch to it.
4. Run `git switch -c NAME` (append the explicit START_REF only when supplied).
   Require success; its output completes verification. Report the branch briefly.

Pass shell arguments safely and separately. No PR/CLI references, tool probes,
tests, fetches, pushes, task initialization, stash, or history changes are needed.
If an explicit start ref is unavailable locally, report it without substituting
another ref. Stop on errors; do not force checkout or retry uncertain creation.
