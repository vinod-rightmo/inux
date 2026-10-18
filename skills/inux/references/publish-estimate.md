# Publish estimate

Only an explicit publishing request authorizes posting a Jira comment. Read the
saved estimate and ticket.md. Require a matching ticket, completed scores/reasons,
an accurate sum, numeric Final Estimate in SP, and a complete Overall Reason.
Report missing fields instead of inventing content or rewriting the saved draft.

Use private acli to inspect existing comments. If identical content already exists,
report it instead of creating a duplicate. Never replace an unrelated comment.

Run `<ix> adf ABSOLUTE_ESTIMATE_PATH TEMPORARY_OUTPUT.json`. Check exit status
and inspect the ADF. Use that file with the installed acli's comment-create
body-file input and explicit ticket key; do not reconstruct ADF by hand. CLI help
and [Atlassian's reference](https://developer.atlassian.com/cloud/acli/reference/commands/jira-workitem-comment-create/)
provide the command syntax.

Recheck ticket ID and saved estimate before posting. If either changed, stop.
Inspect remote comments before retrying an uncertain result. Clean up temporary
output and report the comment or ticket URL. Keep the task files unchanged.
