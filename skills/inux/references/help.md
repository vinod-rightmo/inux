# Show Inux help

Return the contents of the response block below as rendered Markdown.
Add examples only when asked. No metadata lookup is needed.

<!-- x-release-please-start-version -->
```markdown
Inux v0.0.1 — Vinod Liyanage

| Command | Purpose |
|---|---|
| `/inux help` | Show commands. |
| `/inux setup` | Set up tools and sign in. |
| `/inux task <id>` | Load a Jira ticket and create task files. |
| `/inux branch create` | Create and switch to a local task branch. |
| `/inux commit changes` | Stage intended changes and create a local commit. |
| `/inux estimate [request]` | Draft, improve, explain, or publish an estimate. |
| `/inux technical summary [request]` | Create, refine, or publish a standalone summary to Jira. |
| `/inux pr [request]` | Prepare, refine, explain, or publish a PR. |

Publishing requires an explicit `/inux estimate publish`,
`/inux technical summary publish`, or `/inux pr publish` request.
```
<!-- x-release-please-end -->
