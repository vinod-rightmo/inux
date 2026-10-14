# Start a task

`/inux task <id>` replaces the current `.inux/` and initializes both draft
files from the skill assets. On success it contains ticket.md, estimate.md,
and pr.md. No branch, commit, PR readiness check, or publication is needed.

1. Resolve the project/requested ID and snapshot the current task files/ID.
   Fetch with private acli:
   `jira workitem view TICKET-ID --fields key,issuetype,summary,status,assignee,description,attachment --json`.
2. Require success, a matching response key, and complete details. Assemble
   Markdown beginning `# TICKET-ID`, then type, summary, status, assignee,
   full description, criteria, and attachment names/URLs. Preserve lists/tables;
   use None for empty optional fields. Failed retrieval leaves the task intact.
3. Resolve and read both [estimate.md](../assets/estimate.md) and
   [pr.md](../assets/pr.md) before replacing anything. If either template is
   missing/unreadable, stop with the old task intact. Keep the PR template bytes
   ready to save. Recheck the project and task snapshot; stop on changed files/ID.
4. Run `ix task ABSOLUTE_ESTIMATE_TEMPLATE ABSOLUTE_PROJECT` using [cli.md](cli.md).
   Only after success, copy the prepared PR template bytes to `.inux/pr.md` and
   save the prepared `.inux/ticket.md`. Do not fill draft placeholders yet.
5. Verify all three files exist and both drafts match their asset templates.
   Report their paths. On any write failure, report the partial task without
   claiming success or automatically resetting it again.
