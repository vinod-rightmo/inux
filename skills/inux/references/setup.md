# Set up Inux tools

Use `<storage>/tools/bin/`, with `.exe` on Windows. Storage defaults to
`<user-home>/.inux` and may be explicitly overridden. Keep working gh/acli unless broken or an upgrade
is requested. ix must match the loaded skill version; setup installs the matching
release when missing, unversioned, or incompatible. Detect host OS/CPU and use only matching assets. Download into a temporary
directory, verify before execution, then install. Preserve working tools on failure.

Verify gh/acli with --version. Read metadata.version from SKILL.md and require
`ix --version` to report `ix <version>`; also check ix --help.

1. Install missing gh from [official releases](https://github.com/cli/cli/releases),
   verifying its checksum. Run absolute-path `gh --version` and
   `gh auth status --hostname github.com`. If needed, let the user complete
   `gh auth login --hostname github.com --web` with private-repository/SSO access.
2. For ix, use `vinod-rightmo/inux` and tag `v<metadata.version>` by default;
   honor an explicitly supplied repository. The same release version covers
   skill and CLI. Use authenticated gh:
   `release view TAG --repo OWNER/REPO --json tagName,assets`, then
   `release download TAG --repo OWNER/REPO --pattern ASSET --pattern CHECKSUM_FILE --dir TEMP_DIRECTORY`.
   Select `ix-darwin-amd64`, `ix-darwin-arm64`, `ix-linux-amd64`,
   `ix-linux-arm64`, or `ix-windows-amd64.exe`, plus `checksums.txt`.
   These are raw binaries; verify the exact SHA-256 entry before executing.
   Stop on missing access/assets or unsupported OS/CPU. Install as ix/ix.exe,
   set Unix executable permission, and verify --version/--help. No compiler build.
   Never silently substitute another release or overwrite a working tool on failure.
3. Install missing acli using [official instructions](https://developer.atlassian.com/cloud/acli/guides/install-acli/)
   for this host, checking official checksums when supplied. Verify `acli --version`
   and its supported auth-status command; let the user complete
   `acli jira auth login --web` when needed.

Credentials stay managed by gh/acli; never request tokens in chat. Do not change
system PATH or shell profiles. Report verified paths/status for each tool and any
partial failure. Setup does not create a task or publish.

## Testing mode

Check private ix --version/--help, gh/acli --version, and read-only auth status.
Compare ix with the skill version; mismatch is BLOCKED without installing.
Report missing tools/login as BLOCKED. Do not install, log in, or print credentials.
When called from the task suite, include results in its temporary report.
