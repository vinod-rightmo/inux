# Show Inux help

Return `Inux v<version> — <author>` from SKILL.md metadata, the table below, and
the publishing sentence. If the host omitted metadata, read only SKILL.md's
frontmatter once; reuse it for later commands. No project inspection, tool
checks, or file writes. Add examples only when asked.

| Command | Purpose |
|---|---|
| `/inux help` | Show commands. |
| `/inux setup` | Set up tools and sign in. |
| `/inux task <id>` | Load a Jira ticket and create task files. |
| `/inux estimate [request]` | Draft, improve, explain, or publish an estimate. |
| `/inux pr [request]` | Prepare, refine, explain, or publish a PR. |

Publishing requires an explicit `/inux estimate publish` or `/inux pr publish` request.
