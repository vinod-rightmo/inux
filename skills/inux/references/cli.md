# Markdown to ADF

Resolve `<adf>` to `<user-home>/.inux/tools/bin/adf` (`adf.exe` on Windows),
or `<storage>/tools/bin/adf` for an explicit Inux storage root. Pass paths as
separate arguments:

```text
<adf> INPUT.md OUTPUT.adf
```

The converter reads UTF-8 Markdown and writes Atlassian Document Format JSON to
OUTPUT.adf. Both paths are required. It preserves the Markdown input, including
hard-link and symbolic-link aliases, and replaces output atomically. Conversion
is silent on success; errors go to stderr with a nonzero exit status. `--help`
shows usage. The binary has no application version or subcommands.

Inux setup downloads the manually built platform binary from the repository's
`bin/` directory using authenticated GitHub CLI access. Go is needed only for
building the converter. If the private executable is missing, request
[setup](../flows/setup/setup.flow); other actions report the missing tool.
