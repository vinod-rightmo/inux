# Start a task

`/inux task <id>` replaces the entire current `.inux/`, including custom files.
Use the explicit/current project root and absolute private ix/acli paths. The
CLI owns filesystem checks and initializes all three files; the agent fetches
Jira and fills ticket.md. Keep routine local steps quiet.

1. Run `ix task --snapshot ABSOLUTE_PROJECT`. Retain its `snapshot` token unchanged;
   ix checks the directory and captures its state without writing. Do not generate
   shell listings/hashes or separately inspect tools, assets, or task files.
2. Fetch with private acli:
   `jira workitem view TICKET-ID --fields key,issuetype,summary,status,assignee,description,attachment --json`.
   Run this fetch outside the sandbox on the first call, using the shared
   Jira execution rule. It needs Jira network and stored credentials.
   Require success, a matching key, and complete details before replacing files.
   Prepare the layout below, preserving lists/tables and media as links from
   supplied attachment filenames/content URLs. Keep unresolved references as
   given; do not invent mappings or criteria. This reference covers ticket
   formatting; no CLI/PR reference or conversion lookup is needed.
3. Run `ix task --expect SNAPSHOT ABSOLUTE_ASSETS_DIRECTORY ABSOLUTE_PROJECT`,
   using the linked [assets directory](../assets). ix validates both templates
   and rejects a changed snapshot before replacement. It creates empty ticket.md,
   copies estimate.md/pr.md exactly, and preserves existing Git exclusion rules
   and .gitignore. Plain directories work; linked Git worktrees are unsupported.
   Require exit 0 and JSON `ok: true` with all three names in `files`.
4. Replace the returned directory's empty ticket.md with the prepared Markdown.
   A successful CLI result and successful ticket write complete the action; no
   template reads/copies, heading readback, or final filesystem check is needed.
   Report the ticket and three paths briefly. On any error, stop and report the
   actual state; do not reset again or automatically retry a changed snapshot.

## Ticket layout

Fill this layout from the retrieved ticket. Attachments list names and URLs;
use None when there are no attachments. This is the complete ticket format.

```markdown
# TICKET-ID

Type: <type>
Summary: <summary>
Status: <status>
Assignee: <assignee or None>

## Description

<full description or None>

## Criteria

<explicit criteria or None>

## Attachments

- <filename> — <URL>
```
