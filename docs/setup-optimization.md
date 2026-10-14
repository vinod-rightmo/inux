# Inux setup optimization

## Scope

Source: `/Users/vinodliyanage/Developer/Work/inux/skills/inux/`.
Install updates in both `~/.cursor/skills/inux/` and `~/.agents/skills/inux/`;
inspect the actual loaded path in each evaluation. Review runs only when the user
asks in this chat. Do not create or resume a schedule. Preserve private storage,
checksum verification, user-completed login, cleanup, and publishing boundaries.

## Baseline: 2026-10-14

Saved chat: `/Users/vinodliyanage/.inux/logs/chats/d92f55c1-0e19-4753-88ef-64eff762a57b.md`.

SHA-256: `a7a4afda15d485524a3d8f449866fcce55fff3bf0ba8bb0e30e8088dc753c1f9`.

The setup turn installed missing gh, acli, and ix on macOS arm64 and reported
successful authentication and gh/ix checksum verification. ACLI's checksum was
reported unavailable. Chat-save and its sandbox retry are excluded.

| Metric | Baseline observation |
|---|---|
| Reference reads | 2: setup.md, then 20 lines of SKILL.md for version metadata |
| Shell calls | 8 |
| Other tools | 1 dynamic-tool discovery and 1 web fetch |
| Total tool calls | 12 |
| Generated shell commands | 4,596 characters |
| Exported thinking duration | 41,914 ms across 11 distinct setup entries, after deduplication |
| Timestamp span | 09:00:19–09:02:24 UTC: 125 seconds from first setup thinking entry to final response preparation |
| Actual token counts / complete response latency | Unavailable |

The timestamp span is a proxy, not an end-to-end response measurement. Summed
thinking durations are exporter data, not token counts or total execution time.
Do not count the duplicated thinking entries twice.

## Supported findings

1. Setup repeats the help command's old metadata dependency. A version in the
   setup reference, synchronized by the release updater, removes a separate read.
2. The ix download, checksum/install checks, and cleanup used three shell calls
   even though no user interaction was needed between them. They can form one
   uninterrupted phase. The gh/acli install and validation can also be grouped.
3. Generated commands manually parsed and compared hashes and printed expected/
   actual values. Host checksum-file verification can shorten that work while
   retaining exact-asset selection and failure checks.
4. The first metadata request targeted `/tmp/gh-release.json` outside the unique
   setup directory. Its cleanup was not covered by deleting that directory.
   The transcript does not establish whether a file remained after the failure.
5. One GitHub API request failed with 403 and was retried with different network
   permissions and a User-Agent. The export lacks the error response. Do not
   infer the cause or add a host-specific retry policy from this single example.

## Round 1 changes

- Put the target ix release in setup.md and register its annotation with the
  existing Release Please generic updater. Keep it equal to SKILL.md's version.
- Check availability before invoking absent executables; retain individual exit
  statuses and the existing successful-check fast path.
- Describe uninterrupted installation phases explicitly, including ix validation
  and cleanup, to avoid splitting calls just to list files or inspect hashes.
- Prefer built-in checksum-file verification for the exact selected asset entry.
  Keep all downloaded metadata and artifacts in the unique temporary directory.
- Keep failure details and final version/auth/verification outcomes, while omitting
  successful directory listings and hash dumps. Cleanup covers blocked paths.

No runtime helper, new CLI command, added setup dependency, or live tool
installation is introduced by this optimization. Vendor URLs, authentication
commands, release assets, and private storage rules retain their existing meaning.
The skill remains usable across macOS, Linux, and Windows; Unix checksum commands
are examples where available, not a new Windows requirement.

## Expected savings and measurement limits

The avoided 20-line metadata read contains 792 characters. The revised setup
reference adds 215 characters. At a rough four characters per token, the net
instruction-content saving is about 140 input tokens per run, plus removal of
the Read call's protocol overhead. This is an estimate, not measured model usage;
tokenization, host context, and caching can change it.

For a comparable missing-tools installation, the target is one reference read
and approximately 5–6 shell calls instead of 8 when no login interrupts the work.
Actual network failures and user interaction may require additional calls.
Shorter generated commands/output may save more tokens, but that amount is not
yet measured. Fewer round trips should reduce orchestration time; no numerical
speedup is verified. Download/authentication latency can dominate the 125-second
baseline proxy and may vary independently of the skill.

## Evaluation loop

The user runs setup in a new Cursor chat, saves it, and requests review here.
Do not remove installed tools or credentials to manufacture a comparison.

An already-installed run tests the fast path: one reference read, one batch of
availability/version/auth checks, and a compact final status with no downloads
or writes. Its total runtime cannot be compared directly with the fresh-install
baseline as proof of optimization. A naturally needed install/upgrade can test
the changed installation phases and cleanup without destructive preparation.

For each new trace, record its hash, loaded installation, installed revision,
tool counts, generated-command size, verification/auth outcomes, and cleanup.
Distinguish measured changes from estimates in the user report. Preserve the
skill when the behavior is correct; change it only for a demonstrated meaningful
issue. Stop when the relevant changes are verified or further variation belongs
to network/host behavior. Extra reruns are optional, not a fixed quota.

Validation: target release agrees with skill metadata; the existing generic
release updater registers setup.md exactly once. Native checksum-file mode
accepts matching fixture bytes and rejects altered bytes. Source and both
installed references match after installation. No tools were downloaded or
credentials changed during this update.

Installed setup.md SHA-256: `2c2e90163c9370ea5c32f3f45706cc88ffb98a07726d3d660bdafff7ab49b762`.

Status: round 1 evaluated in the next section. Reference reads and shell calls
decreased, but generated command size and exported thinking duration increased.
The follow-up justifies a small instruction correction, not a larger installer.

## Follow-up: 5bc85258-5912-489c-b66e-77ad56ae1a0a

Saved chat: `/Users/vinodliyanage/.inux/logs/chats/5bc85258-5912-489c-b66e-77ad56ae1a0a.md`.

SHA-256: `8d55a2e2a5e0d1dea15a33f43e688ad33a9be6a4b1462361e18a08d46c7dce51`.

This was another missing-tools installation, not an already-installed fast-path
run. The trace loaded the updated Cursor setup reference and installed all three
tools. It reports successful authentication and cleanup. The user measured
1 minute 58 seconds. Both installed references and source matched round 1.

| Metric | Baseline | Follow-up | Observation |
|---|---|---|---|
| Reference reads | 2 | 1 | 50% fewer |
| Shell calls | 8 | 4 | 50% fewer |
| Total tool calls | 12 | 5 | 58% fewer |
| Generated command characters | 4,596 | 9,211 | 100% more |
| Deduplicated exported thinking | 41.914s | 78.609s | About 88% more |
| Setup thinking timestamp span | 125s | 115s | 8% shorter proxy; not a response-latency measurement |
| User-reported total duration | Unavailable | 118s | No equivalent measured baseline |

Fewer calls did not establish lower total token usage. At four characters per
token, the extra 4,615 generated command characters correspond to roughly
1,150 additional output tokens, offsetting the earlier small input-read saving.
This is only a heuristic for command text; reasoning-token and total usage
counters remain unavailable. Timing includes variable downloads/authentication,
so the 8% proxy change cannot be attributed solely to the instructions.

The generated commands expanded into repeated cleanup functions, custom Python
checksum parsing, and a large HTML/checksum-format parser for ACLI even though
no checksum was linked. The instruction to combine entire phases in one call
appears to have encouraged this boilerplate. Its download/checksum failures were
also often printed as exit codes before continuing to copy; `set -u` at the start
was not fail-fast protection. Successful installation does not validate those
failure paths.

## Round 2: modest wording correction; no helper scripts

The user prefers instruction-only improvements because setup is infrequent.
A reusable installer was considered, then abandoned before any script file was
written. No helper, Node dependency, CLI command, or packaging change is added.

Relax the one-call-per-phase instruction. Prefer short host commands, reuse one
temporary directory, and inspect the actual linked checksum before writing
verification logic. Discourage custom checksum/HTML parsers and speculative
format/login branches. Login is a separate command only after an auth check
requires it. Gate dependent steps with `&&` or explicit checks so verification
failure prevents installation. Keep private paths, exact ix version, checksum
requirements, cleanup, and the one-reference-read flow.

The reference grows only 36 characters (about nine input tokens heuristically).
The intended benefit is less generated boilerplate and clearer failure handling,
not another reduction in instruction size or a guaranteed shell-call count.
Further token/time savings are unmeasured; do not promise a numerical speedup.

Validation: the setup target still matches SKILL.md and remains registered with
Release Please. A local checksum fixture proves the suggested `&&` sequence
preserves an existing destination on mismatch and copies only on a matching
checksum. Both installed references match source after installation. No real
tool download or credential change was performed during this edit.

Current setup.md SHA-256: `da452f3a28ec2a5e8d61e8a5228d4ca828397de64a48bd082474f92384f14185`.

Status: round 2 installed on 2026-10-14 in both skill directories. Leave further
review on demand; a rerun is optional. Judge a future comparable trace by correct
verification/cleanup and reduced generated work, rather than fewer calls alone.
