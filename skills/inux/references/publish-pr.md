# Publish pull request

Only an explicit publishing request authorizes pushing and creating/updating a PR.
Read ticket.md and pull-request.md. Require a matching ticket and nonempty,
single-line Title; Description is optional. Use evidenced repository conventions.

Check private gh authentication, the repository, current branch and working tree.
Fetch origin and inspect the changes intended for the PR. Resolve the base branch
from the request or repository convention; otherwise ask with dev suggested.
Stop on unresolved checks instead of committing, merging or rebasing automatically.

Recheck the branch, HEAD, ticket and saved PR content before mutation. Push without
force. Use private gh to create a draft with explicit base/head and self-assignment,
or update the existing open PR for this exact head branch. Use a temporary UTF-8
body file, preserving the saved Markdown. Multiple matching PRs need clarification.

Inspect the resulting PR and report its URL. If the push succeeds but PR creation
fails, report both outcomes. Inspect remote state before retrying to avoid duplicates.
No Inux CLI publishing command or hidden state is needed.
