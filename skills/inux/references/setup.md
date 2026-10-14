# Set up Inux tools

Resolve private `ix`, `gh`, and `acli` under `<storage>/tools/bin/` (`.exe` on
Windows). Storage defaults to `<user-home>/.inux`; honor an explicit override.
Use absolute executable paths. PATH-only tools do not satisfy setup. Use vendor
binaries/archives, without package managers, sudo, external tool symlinks, or
changes to PATH/shell profiles.

## Check first

Read the skill version from available metadata; if absent, read only SKILL.md's
frontmatter. Detect OS/CPU and batch independent checks of existing private tools:

- `ix --version` must report `ix <skill-version>`.
- `gh --version` and `gh auth status --hostname github.com`.
- `acli --version` and `acli jira auth status`.

Keep each exit status so one failure does not hide the other results. If all pass
and no upgrade is requested, report success and stop. Install only missing/broken
tools, mismatched ix, or explicitly requested upgrades. Keep other working gh/acli
versions. Recheck only a tool or login changed by this run; reuse other results.

## Install when needed

Create one unique temporary directory only if a download is needed. Verify
checksums before execution, then install; preserve working tools on failure.
Use host checksum/extraction tools instead of generating lengthy scripts when
possible. Batch related steps with clear failure handling. In zsh, avoid `path`
as a variable because it changes PATH. Put heredoc terminators on their own line.

1. **gh:** Select this OS/CPU's archive and checksum from the
   [official releases](https://github.com/cli/cli/releases). Verify its SHA-256,
   extract, and copy gh/gh.exe to the private directory. Set Unix executable
   permission. Verify version and GitHub auth. If login is needed, let the user
   complete `gh auth login --hostname github.com --web` using that private gh.
2. **ix:** Use `vinod-rightmo/inux`, tag `v<skill-version>`, unless another
   repository is explicitly supplied. Download with private gh:
   `release download TAG --repo OWNER/REPO --pattern ASSET --pattern checksums.txt --dir TEMP_DIRECTORY`.
   Assets: `ix-darwin-amd64`, `ix-darwin-arm64`, `ix-linux-amd64`,
   `ix-linux-arm64`, `ix-windows-amd64.exe`. Verify the exact SHA-256 entry,
   install as ix/ix.exe, set Unix executable permission, and verify version
   plus `--help` once. Do not build from source or substitute another release.
3. **acli:** Download the matching [official standalone binary](https://developer.atlassian.com/cloud/acli/guides/download-supported-packages/):
   `https://acli.atlassian.com/<os>/latest/acli_<os>_<arch>/acli`
   (`darwin`, `linux`, `windows`; `amd64`, `arm64`; append `.exe` on Windows).
   Check the official package listing at most once for a linked checksum. Verify
   it if supplied; otherwise report checksum unavailable and use the official
   HTTPS binary. Do not guess checksum URLs, crawl pages, or dump headers.
   Install privately, set Unix executable permission, and verify version and
   `acli jira auth status`. If needed, let the user complete
   `acli jira auth login --web`. Jira status is the relevant auth check for Inux.

Stop affected steps on unsupported OS/CPU, missing access/assets, or download/
verification failure. Keep failure details; keep successful output to versions,
auth status, and required asset names. Avoid full release JSON, archive listings,
and help dumps. Inspect command help only for an actual syntax/version error.
Credentials stay managed by gh/acli; never request tokens in chat.

## Finish

On success or failure, remove this run's temporary directory and its contents;
remove empty temporary parents only if this run created them. Preserve installed
tools, credentials, task files, and pre-existing/unrelated files. Report any
leftover path and cleanup error. Cleanup does not repeat successful tool checks.

Return one compact tool/version/status table and the shared install directory.
Include blockers or checksum/cleanup limitations when present. Setup does not
create a task or publish.
