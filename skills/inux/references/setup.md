# Set up Inux tools

Target ix release: `1.0.0`. <!-- x-release-please-version -->

Resolve private `ix`, `gh`, and `acli` under `<storage>/tools/bin/` (`.exe` on
Windows); storage defaults to `<user-home>/.inux`. Honor an explicit override.
Use absolute paths and vendor binaries/archives. No package managers, sudo,
external tool symlinks, or PATH/shell-profile changes. PATH-only tools do not count.

## Check first

Use the target release above; no metadata read is needed. Detect OS/CPU and
batch availability, version, and auth checks in one tool call. Skip commands
for missing executables; retain each exit status so failures do not hide results:

- `ix --version` must equal `ix <target-release>`.
- `gh --version` and `gh auth status --hostname github.com`.
- `acli --version` and `acli jira auth status`.

If all pass and no upgrade is requested, return the status table and stop.
Install only missing/broken tools, mismatched ix, or requested upgrades. Keep
working gh/acli versions. Recheck only tools or authentication changed this run.

## Install when needed

Create one unique temporary directory before fetching metadata or assets. Reuse
it across tool calls; keep all downloads inside it and clean up once at Finish.
Use short host commands; grouping is optional. Do not generate general
installers or create helper script files. Preserve working tools on failure.

Select the exact checksum entry with host text tools; no custom parser is
needed. From its directory, use host verification (`shasum -a 256 -c` on macOS,
`sha256sum -c` on Linux). Gate download, verification, extraction, and copy with
`&&` or explicit error checks; printing exit codes must not replace stopping on
failure. Use native extraction tools. In zsh, avoid the special `path` variable;
keep heredoc terminators on their own line.

1. **gh:** Select this OS/CPU's archive and checksum from the
   [official releases](https://github.com/cli/cli/releases). Verify SHA-256,
   extract, and install privately with Unix executable permission. Verify version
   and GitHub auth. If needed, let the user complete private
   `gh auth login --hostname github.com --web`.
2. **ix:** Use `vinod-rightmo/inux`, tag `v<target-release>`, unless another
   repository is explicitly supplied. Download with private gh:
   `release download TAG --repo OWNER/REPO --pattern ASSET --pattern checksums.txt --dir TEMP_DIRECTORY`.
   Assets: `ix-darwin-amd64`, `ix-darwin-arm64`, `ix-linux-amd64`,
   `ix-linux-arm64`, `ix-windows-amd64.exe`. Verify SHA-256, install as ix/ix.exe,
   set Unix executable permission, and verify the exact target version and
   `--help` once. Do not build from source or substitute a release.
3. **acli:** Use the matching [official standalone binary](https://developer.atlassian.com/cloud/acli/guides/download-supported-packages/):
   `https://acli.atlassian.com/<os>/latest/acli_<os>_<arch>/acli`
   (`darwin`, `linux`, `windows`; `amd64`, `arm64`; append `.exe` on Windows).
   Inspect the listing once with an available page reader; verify an actually
   linked checksum or report unavailable and use the official HTTPS binary.
   Do not build HTML parsers or hypothetical checksum branches, guess URLs,
   or dump headers. Install privately, set Unix executable permission, and
   verify version plus `acli jira auth status`.
   If needed, let the user complete private `acli jira auth login --web`.

Run login as a separate command only after an auth check reports it is needed.
Stop affected steps on unsupported OS/CPU, missing access/assets, or download/
verification failure. Inspect command help only for a syntax/version error.
Credentials stay managed by gh/acli; never request tokens in chat.

## Finish

On success or failure (including metadata or login blockers), remove this run's
temporary directory and any empty parents it created. Preserve installed tools,
credentials, task files, and unrelated files. Report leftover paths/cleanup errors.
Do not repeat successful tool checks during cleanup.

Return one compact tool/version/status table and the private install directory.
Include failures or checksum/cleanup limitations; omit successful hash dumps,
metadata, directory/archive listings, and help output. Setup creates no task
and does not publish.
