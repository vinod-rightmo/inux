# Show Inux help

Show `Inux v<version> — <author>` using metadata.version and metadata.author
from the already loaded SKILL.md frontmatter, followed by this table. Explain
that requests use plain language and stop. Keep the values only in SKILL.md;
do not duplicate them here.
Do not inspect a project, require a ticket, check tools, or write files.

| Command | Purpose |
|---|---|
| `/inux help` | Show commands. |
| `/inux setup` | Install/check shared tools and sign in. |
| `/inux task <id>` | Fetch a ticket and initialize ticket.md, estimate.md, and pr.md. |
| `/inux estimate [request]` | Generate, improve, explain, or explicitly publish the estimate. |
| `/inux pr [request]` | Prepare a PR draft, refine title/description, explain, or explicitly publish. |
| `/inux test task [id]` | Test task handling and check tools in a temporary project. |
| `/inux test estimate [request]` | Test estimate writing, explanations, and ADF conversion. |
| `/inux test pr [request]` | Test readiness, preparation, and PR refinements. |

Add `test` after `/inux` to test any action, such as `/inux test setup` or
`/inux test pr title`. Bare estimate/PR tests run their non-publishing actions.

Examples: `/inux estimate improve the reasons`, `/inux pr prepare`,
`/inux pr improve the title`, `/inux pr improve the description`.
Bare `/inux estimate` fills its initialized template or improves existing content;
bare `/inux pr` prepares a draft. Publishing requires an explicit
`/inux estimate publish` or `/inux pr publish` request. Test mode never publishes.

## Testing mode

Show the help table and check that the action requires no active project,
ticket, or tool. Report PASS/FAIL in chat without inspecting or writing files.
