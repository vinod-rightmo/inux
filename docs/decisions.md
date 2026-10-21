# Implementation decisions

## 2026-10-21: Keep one unversioned Markdown converter

**Decision:** Rename the Go command from ix to adf and accept exactly an input
Markdown path and an output ADF path: `adf input.md output.adf`. Keep conversion,
UTF-8 validation, input protection, and atomic writes. Build platform binaries
manually, commit them in `bin/`, and let setup retrieve the current binary through
GitHub CLI. Remove application versions, Release Please, GoReleaser, and GitHub
workflows. This supersedes earlier ix version and release decisions below.

**Reason:** Conversion is the Go program's only job. Manual builds and pushes
keep its delivery aligned with the skill without a separate release process.

## 2026-10-20: Keep editable drafts outside the repository

**Decision:** Task loading keeps Jira details in conversation context. Create
only the requested estimate, PR, or technical-summary Markdown draft in a unique
private OS temporary directory outside the repository. Developers can edit it
directly or with the agent. Reuse its exact path for the same project, ticket,
and draft kind. Remove the selected file after verified publication or explicit
completion; retain it when publication fails or its outcome is uncertain.

**Reason:** The agent already manages Markdown editing and publication. Removing
the repository task directory also removes ix task initialization, workspace
snapshots, template copying, and Git exclusion logic. ix retains ADF conversion
and version reporting. This supersedes the task-file layout decision below;
existing user task folders are left untouched.

## 2026-10-14: Install Inux tools in private storage

**Problem:** `/inux setup` attempted to install Atlassian CLI with Homebrew,
although Inux tools belong under the configured Inux storage root.

**Cause:** The setup reference stated the private storage location but delegated
ACLI installation to Atlassian's general installation guide. That guide includes
Homebrew and system-directory installation, leaving the installation method
ambiguous. The source and installed skills contained the same wording.

**Decision:** Resolve all tools under `<storage>/tools/bin/` before checking or
installing them. Storage defaults to `~/.inux`, so ACLI belongs at
`~/.inux/tools/bin/acli` (`acli.exe` on Windows). Keep working private copies;
otherwise download the matching official standalone binary or archive. Invoke
`ix`, `gh`, and `acli` through their absolute private paths. A PATH-only tool
does not satisfy setup. Package managers, system installation directories,
sudo, and links to externally managed binaries are excluded from this workflow.

**Reason:** A predictable installation location makes Inux setup consistent
across machines and operating systems, avoids global package-manager changes,
and keeps the tools used by the skill independent of unrelated PATH versions.

**Changed files:** [setup reference](../skills/inux/flows/setup/setup.flow) and
[skill entry point](../skills/inux/SKILL.md). The correction applies to local
Codex and Cursor installations as well as the repository source. It does not
change the CLI commands or require a new tool download when a private copy works.

## 2026-10-14: Support only the current task-file layout

**Problem:** PR guidance allowed evidence from an older task filename, and the
skill and README still mentioned removed commands and earlier implementations.
These instructions made agents consider obsolete workflows during current tasks.

**Decision:** Support the current `.inux/ticket.md`, `.inux/estimate.md`, and
`.inux/pr.md` layout. PR preparation and explanations read the initialized
`pr.md`. Remove the legacy-file fallback and historical command/layout wording.
Existing-file preservation and missing-file checks remain part of current task
handling; they do not introduce a migration or compatibility path.

**Reason:** Inux does not need backward compatibility. Guidance should describe
only supported behavior so agents have fewer branches to consider.

**Review:** A subagent audited the skill package and README. The cleanup changes
instructions only; it does not delete files from users' task directories.

## 2026-10-14: Use normal actions without a separate testing mode

**Problem:** Dedicated test commands and testing sections duplicated the normal
workflows and added fixture, reporting, and routing instructions to every action.

**Decision:** Remove `/inux test` commands, testing-mode sections, and their help
and README examples. Use normal actions with a plain-language “do not publish”
instruction when exercising the skill. Publishing continues to require explicit
intent. Existing action validation, PR readiness checks, CLI unit tests, and
release checks remain part of their respective workflows.

**Reason:** The publishing boundary already controls external writes. A separate
skill testing mode adds instructions and maintenance without being needed for
the user's workflow.

## 2026-10-14: Clean up setup staging files

**Problem:** Setup downloaded and extracted tool packages into temporary storage
but did not explicitly require cleanup, leaving archives, extracted manual pages,
checksums, and checksum-lookup HTML behind.

**Decision:** Use a unique temporary directory for each setup run. Before
returning on success or failure, remove that run's staging directory and files,
and any empty temporary parents it created. Preserve installed tools,
credentials, task data, and pre-existing or unrelated files. Report cleanup
failures with the leftover path and error.

**Reason:** Download and extraction files are installation intermediates, not
part of the working tool installation. Cleanup keeps storage tidy while limiting
deletion to files owned by the current setup run.

## 2026-10-14: Reduce help/setup exploration and repeated checks

**Evidence:** A supplied Cursor conversation contained two help file reads and
18 setup shell calls, with 8,588 characters of generated shell commands. Five
calls investigated ACLI checksum sources. The run also explored multiple ACLI
auth command groups, retried a malformed heredoc, recovered from zsh's special
`path` variable changing PATH, and repeated successful tool checks during cleanup.
The inline skill omitted frontmatter, so the help metadata read was necessary;
it can be limited to the header rather than the first 30 lines.

**Decision:** Reuse provided instructions and metadata, keep help concise, check
existing tools before downloading, and finish immediately when checks pass.
Use the exact Jira auth command, one official ACLI checksum-listing lookup,
existing host checksum tools, and short shell steps with correct failure handling.
Recheck only changed tools/authentication. Keep successful output compact and
avoid full metadata/help dumps. Keep checksum verification for gh/ix, private
installation, cleanup, and explicit publishing intent.

**Reason:** Avoiding unnecessary exploration and retries reduces generated code,
tool responses, and conversation turns. Shorter instructions also reduce the
context loaded for help/setup. This adds no runtime helper or per-run log.

**Validation and limits:** The installed ACLI help and official command reference
confirm `jira auth status` and `jira auth login --web`. Instruction structure,
links, and installation copies are checked locally. The export contains neither
tool results nor token-usage counters, so exact token savings cannot be measured;
a fresh comparable run is needed to measure the behavior change.
