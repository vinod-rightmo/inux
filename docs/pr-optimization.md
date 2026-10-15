# Inux PR optimization

## Scope

Improve preparation of `.inux/pr.md` using the observed run, without helper
scripts, new native commands, or automatic publishing. Source lives in
`/Users/vinodliyanage/Developer/Work/inux`. Synchronize the Cursor-specific and
shared skill installations. Review remains manual and on demand.

The developer supplies the destination with `/inux pr target BRANCH`, equivalent
plain language, or an earlier explicit choice for the same PR. Otherwise ask.
Validate commits in the selected remote target's `TARGET_REF..HEAD` range, not
against fixed main history. These preceding corrections remain intact.

## Baseline: 2026-10-14

Chat: `/Users/vinodliyanage/.inux/logs/chats/2056e2b7-d5b8-40df-9655-f8166dec66c8.md`.
SHA-256: `918488f773ffb431ea5c61f7ce2b6610950a9c589700d8756f04048488a5028c`.

Request: `/inux pr target release/v2.9.0 create pr.md`.
The branch and selected-target range resolve correctly to one OQG3-3895 commit.
The agent writes the PR draft without pushing or publishing. The export omits
tool-result bodies and actual token counters; success and readiness claims are
not independent proof of all underlying checks. Exclude chat-save and deduplicate
identical thinking entries.

| Metric | Observed |
|---|---|
| Read calls | 14: six initial skill/task reads, seven source reads, one readback |
| Search calls | 1 Glob and 6 Grep |
| Shell calls | 10 |
| Write calls | 1 |
| Total calls | 32 |
| Generated shell commands | 1,924 characters |
| Commands plus written PR text | 3,188 characters |
| Distinct exported thinking entries | 17 |
| Summed exported thinking duration | 54.080s |
| Task thinking timestamp span | 10:45:13–10:46:49 UTC; 96s |
| Actual total tokens / full response latency | Unavailable |

## Findings and changes

- The root instruction required cli.md for convention checks, although pr.md
  already provided the regex policy. Make PR preparation self-contained with
  the two native check command signatures and success/error behavior; reserve
  cli.md for ADF conversion or explicit CLI usage.
- The agent read assets/pr.md after reading the initialized task draft. Use the
  draft's labels/order; consult the asset only when its structure is missing or
  unclear. Preserve required sections and accurate developer edits.
- The ignored-task-folder Glob and ix version/executable probe were redundant.
  Read known task files directly and invoke the native validator. Missing tools
  fail normally and block dependent work. Setup owns version checks.
- The initial Shell printed porcelain status and full status, then suppressed
  possible earlier failures with a trailing `|| true`. Inspect local facts once,
  reuse them until the final recheck, and preserve required failure statuses.
- The full commit message appeared twice within one Shell call, including a
  truncated byte dump. Read it once without truncation, preserve its whitespace,
  and validate that full value. Keep the terminal-LF rule and semantic checks.
- A small three-file diff led to seven source reads and six searches through the
  QR-name feature. Start from the complete target diff and inspect nearby changed
  functions/direct dependencies only to resolve specific material gaps. Keep
  causal claims as narrow as the evidence supports, rather than expanding the
  summary into an exhaustive explanation of unrelated flows.
- After a successful Write, the agent used Read, od, and Python to confirm two
  trailing spaces. Preserve the existing Markdown hard break in the write payload
  and trust successful writing. Readback/byte checks are for actual write or
  formatting problems, not a default phase. This removes three calls without
  changing the PR template.
- The trace claims readiness passed but contains no observed test execution.
  Do not infer passing checks from a clean tree or absence of failures. Reuse
  current verification evidence and run missing repository-required checks;
  report unperformed checks and actual results accurately. The export alone
  cannot establish whether earlier verification existed.

No QR-specific rules, mandatory call quotas, whitespace-format redesign, or
native Git/Jira wrapper was introduced. The final branch/HEAD/target/ticket/tree
recheck remains, as do publication approval boundaries and remote verification.

## Expected savings

Seven directly avoidable calls are identified: cli.md, the duplicate asset read,
folder Glob, ix probe, and three post-write formatting checks. Removing only
those projects 32 to 25 calls (22% fewer). Focused code investigation may reduce
more; roughly 15–18 calls is plausible for this small diff, but it is not a
required quota and depends on context and necessary verification.

Current instruction/template files corresponding to the baseline total 18,023
characters (root, PR reference, CLI reference, writing guidance, PR asset). The
revised normal preparation reads total 15,384 characters (root, PR reference,
writing guidance), approximately 660 fewer input tokens at four characters per
token. Ticket/draft content is excluded from both totals. This does not measure
total billed tokens, caching, source/tool-result payloads, or reasoning savings.
New execution time and actual call reductions remain unmeasured. The 96s baseline
span is a proxy; do not promise a proportional speedup or convert thinking time
to tokens.

## Validation and installation

The documented private ix commands were run against the exact captured branch
and full stored message; both pass. Read-only Git checks using the captured
immutable revisions confirm the selected-target range contains only the task
commit. All skill relative links resolve. Source metadata is unchanged; YAML
frontmatter is checked with Ruby because the bundled Python quick validator
requires unavailable PyYAML. No dependency was installed for validation.

Synchronize SKILL.md and references/pr.md to both installations after checking
that source and installed copies have not changed. No live PR creation, network
fetch, reset, or modification of the user's QR project is part of this review.

Status: focused instruction changes validated and installed. Await only normal
user-requested review of a future run; no timer or mandatory rerun.

Installed revision fingerprints:

- `SKILL.md` SHA-256: `e1ff9eb9103b61548fd0021642ba175d05a9b47db704a556e77069fc32789a63`
- `references/pr.md` SHA-256: `25e066e64757e513d16e99ef8248da4b8c094ca7d1a4b671d381f2570d46c8de`

## Follow-up: 77dae035 — faster-path changes only partly followed

Reviewed `/Users/vinodliyanage/.inux/logs/chats/77dae035-df6d-4f76-8988-2aacba70a444.md`.
SHA-256: `956d781a0218f51787cf66fec4ab5c43b826d471d3e482eb6ecc1a75b9875f2f`.
The PR reference matches the preceding installed revision. The shared root has
also gained branch/commit commands through other work; preserve those changes.
Both installed copies match current source before this update.

| Metric | Baseline 2056e2b7 | Follow-up 77dae035 |
|---|---|---|
| Read calls | 14 | 14 |
| Search calls | 1 Glob + 6 Grep | 10 Grep |
| Shell calls | 10 | 14 |
| Write calls | 1 | 1 |
| Total calls | 32 | 39; 22% more |
| Shell-command text | 1,924 characters | 2,269 characters |
| Commands plus written PR | 3,188 characters | 3,446 characters |
| Deduplicated thinking | 54.080s | 85.371s; 58% longer |
| Task thinking timestamp span | 96s | 185s; 89s / 93% longer |

The command/write text grew by 258 characters, roughly 65 generated tokens at
four characters per token. Actual total tokens remain absent. No overall token
or speed improvement is demonstrated. Timings exclude chat-save and duplicate
thinking; spans are proxies, not full response latency or causal accounting.

The two runs differ materially in verification. The follow-up invokes type-check,
lint, Prettier, and route-source checks and reports them passing; baseline has
no such calls. The target project's current AGENTS.md, "Local checks before a PR",
prescribes those four checks, plus unit tests when a tested helper changed.
These are legitimate work, not checks to remove merely to improve the benchmark.
The thinking timestamp window surrounding test execution is approximately 55s;
it is not a measured duration for each command. Exported results are unavailable,
so this review does not independently certify that the checks passed.

### What worked

No cli.md read, folder Glob, ix version probe, post-write readback, or trailing-
space audit appears. Target-relative commit validation remains correct and no
publication occurs. Successful writing completes preparation as intended.

### Remaining waste and focused corrections

- estimate.md is read without being a PR input. State the task input set directly.
- The existing draft structure is explicitly recognized as complete, yet the
  asset template is read. The Technical summary section still linked to the
  asset as its source, conflicting with the preparation fallback rule. Make the
  existing Description fields authoritative and retain just the conditional
  asset lookup for missing labels.
- The full commit is printed twice, once as an od byte dump. Validate the full
  message directly first, normalizing only its terminal LF as already required;
  diagnose an unexpected validator failure afterward.
- Root-cause exploration still expands into speculative edge cases outside the
  immediate summary, including metadata replacement after form submission.
  Explain the immediate behavior/cause supported by the diff, resolve material
  gaps, and stop once the fields are supported. Do not add unsupported wider
  claims that require a full feature audit. Real unresolved concerns still get
  reported rather than ignored.
- Docker status is probed twice due to a permissions retry. There are also two
  package searches and two Makefile reads while rediscovering prescribed checks.
  Use verification commands already provided by repository instructions and
  required host access on their initial invocation. Discover a missing command
  or runtime issue only when needed. For an unmapped script, one targeted lookup
  can remain necessary. Keep the four required checks and accurate reporting.

No QR-specific container/script names were added to the skill, and no helper
script, native wrapper, cached success marker, or test-skipping rule was added.
Do not turn the observed durations into a promised speedup. The initial
15–18-call projection was not realized and should not be treated as established
performance. Approximately seven discovery/template/estimate calls are potential
reductions if documented commands are reused with at most one focused lookup;
additional savings from bounded code investigation are unmeasured.

The root, PR reference, writing guidance and duplicate asset reads in the latest
source total 16,430 characters, excluding ticket/draft/estimate/source/result
payloads. Removing the redundant asset read while updating the PR reference
projects 15,825 characters, about 150 fewer input tokens. This small category
estimate excludes the potentially larger savings from fewer discovery/source
results. Total billed-token savings and a new runtime remain unknown.

### Validation and status

The changes only affect references/pr.md and this evaluation document. Check all
relative links, source/installation equality, and whitespace before finishing.
Preserve current unrelated branch/commit/help changes and all publication gates.
No live fetch, Docker check, PR write, or project modification is performed by
this review. Install the revised reference in Cursor and shared skills. Continue
manual review only when requested, without a timer or mandatory rerun.

Updated PR reference SHA-256: `b96cf7e4c9f1d550709154f5a85cc7a107859edb37208037a50bfdb00dc0f40c`.

## Follow-up: da22335f — fewer calls with required checks retained

Reviewed `/Users/vinodliyanage/.inux/logs/chats/da22335f-328a-4eaa-a6e5-c1f0cb3699f5.md`.
SHA-256: `808d822b4c7361604503b60bd89fed126960c6b449f24fe1de0ef77d6fedcf73`.
Both installed copies match source before review. The PR reference matches the
preceding round's fingerprint; root branch/commit work remains unchanged.

| Metric | Previous 77dae035 | Latest da22335f |
|---|---|---|
| Read calls | 14 | 9 |
| Grep calls | 10 | 4 |
| Shell calls | 14 | 13 |
| Write calls | 1 | 1 |
| Total calls | 39 | 27; 31% fewer |
| Required repository-check invocations | 4 | 4 |
| Shell-command characters | 2,269 | 2,617 |
| Command plus PR write text | 3,446 characters | 3,917 characters |
| Deduplicated thinking duration | 85.371s | 48.912s; 43% shorter |
| Task thinking timestamp span | 185s | 128s; 57s / 31% shorter |

Both runs invoke the same four prescribed checks and report them passing. The
latest has no template/CLI read, ignored-folder discovery, post-write audit, or
Docker permission retry. Code/configuration exploration is reduced. These are
observable improvements, while tool-result bodies and token counters remain
absent. Do not treat the transcript as independent proof of passing checks.

Known instruction/template input decreases from 16,430 to 15,825 characters,
roughly 150 input tokens, as projected in the preceding review. Estimate.md is
still read in both runs and is excluded from those figures. Generated command/
write text grows by 471 characters, roughly 118 output tokens; the local-state
and final-recheck shell commands are verbose. Neither text-category estimate
measures total billed tokens or accounts for fewer source/search result payloads
and hidden reasoning. No net total-token reduction is claimed. Thinking duration
is deduplicated and excludes chat-save; the 128s span is not full response
latency or proof that the skill alone caused the improvement.

Remaining redundant behavior includes estimate.md, a preliminary byte dump,
an ix file listing, a Docker health probe, and package/script discovery despite
provided commands. Most are already covered by the instructions. Do not keep
adding near-duplicate prohibitions or QR-specific exceptions to chase individual
model choices; further execution variability does not by itself justify a wider
rewrite or native wrapper.

One concrete Git clarification is warranted. The agent first invokes
`git merge-base --is-ancestor TARGET_REF HEAD`, then sees a normal non-ancestor
result and performs the actual shared-ancestry check separately. A target branch
can advance independently after the task branch forks. Require
`git merge-base TARGET_REF HEAD` directly; shared ancestry does not require the
target to be an ancestor of HEAD. This avoids the unnecessary precheck and
related interpretation without weakening ancestry protection or commit scope.
The observed fetch command also ends with an echo that masks a possible failure;
the existing requirement to propagate required failures remains in force. No
fetch failure is demonstrated by this trace.

Install only this narrow ancestry clarification in both PR references. Preserve
all other pending work. The clarification's time/token effect is unmeasured;
no live fetch/checks, draft writes, or PR publication are performed during this
review. Continue normal on-demand evaluation without a mandatory rerun.

Updated PR reference SHA-256: `409309d9cbfc823ab1f73b24bfa48f2a4f777e75f7f72bf1fd2d3165f94af823`.

## Summary readability preference

The user requested technical summaries that a developer new to the feature can
understand, written as one developer explaining the change to another. Update
be-human.md to assume programming knowledge rather than internal workflow
knowledge, explain unfamiliar terms, connect cause and result, and introduce
code identifiers through their purpose. Keep necessary context rather than
compressing several ideas into one sentence.

Replace the fixed one-or-two-sentence implementation rule in pr.md with a clear
explanation of the user action, resulting behavior, and mechanism. Explain why
the old behavior caused the problem in Root cause, rather than listing internal
calls or state fields. Keep evidence requirements and existing template layout.
Update only the asset's placeholder hints to match; no new runtime reference,
code investigation requirement, or project-specific example is added.

Synchronize all three affected files to Cursor and shared skills. Existing task
PR drafts are not rewritten by this instruction update. No performance savings
are claimed; readability takes precedence over an arbitrary sentence limit.

## Reader feedback: easy opening, useful examples, meaningful notes

The user refined the desired writing style after reviewing a concrete PR draft.
Open the implementation explanation with a short, simple statement of the
result. Introduce conditions, workflow details, and mechanism afterward rather
than packing them into the first sentence. Add a concise concrete example when
it helps the reader picture the behavior. Root cause should explain what went
wrong and why in the same plain language; reuse the example if useful without
repeating the entire fix. Hypothetical examples must be clear and consistent
with verified behavior.

Important notes is optional. Include only constraints, risks, or follow-up that
matter to the reviewer. Routine passing format/lint/type checks belong in the
completion report and do not justify an Important notes bullet.

Apply the preference in be-human.md, the PR technical-summary instructions,
and the PR asset's placeholder hints. Keep the existing template layout and
new-developer audience; no QR-specific example is hardcoded into the skill.
Synchronize all three files to both installed copies. No new references, scripts,
or tool steps are introduced. This is a readability improvement; no measured
token or latency saving is claimed and no existing task draft is rewritten.

- `references/pr.md` SHA-256: `f0f0dc0f8cf7d6733df9c4683bd23d6a2a5922b0a854bbe80c9fdf444d87d776`
- `references/be-human.md` SHA-256: `ff233a4a34df71cb3a79fd90907bb9d18852e87b4c25a8c111ab13ad7de2046b`
- `assets/pr.md` SHA-256: `353f90c3082ac5dc9d8e1fe3140c133d8edeaa82727cc4b757c33b4ced3caa95`

## Analysis: aa2231a9 — simplify decisions rather than add prohibitions

User-supplied export:
`/Users/vinodliyanage/.codex/attachments/8c48da93-e8bc-4cdd-90a7-0639dc137c5a/Pasted text.txt`.
SHA-256: `7f50c9ebf87fa2860d027e8b0225d488b85c4ced33cda4e1123e856afb496f42`.
Task: `/inux pr target release/v2.9.0 create pr.md`.
Both installed root/PR/writing references match current source. The supplied
raw transcript contains role/message records but no usage or timing metadata;
actual token counts are unavailable there as well as in the pasted export.

| Metric | Previous da22335f | Latest aa2231a9 |
|---|---|---|
| Read / Shell / Glob / Grep / Write | 9 / 13 / 0 / 4 / 1 | 12 / 13 / 1 / 6 / 1 |
| Total calls | 27 | 33; 22% more |
| Shell-command characters | 2,617 | 2,444 |
| Command plus PR write characters | 3,917 | 3,883; essentially unchanged |
| Deduplicated thinking duration | 48.912s | 67.403s; 38% longer |
| Task thinking timestamp span | 128s | 217s; 89s / 70% longer |
| Prescribed check invocations | 4 | 4 |

Thinking entries are deduplicated and exclude chat-save. The latest run has 16
unique task thinking entries. Timestamp spans are not full response latency.
The interval from the check-dispatch planning entry (11:35:04) to the next
post-check summary entry (11:37:12) is 128s. It includes waiting and model work;
do not call it isolated test duration or attribute the entire regression to
the skill. All four checks are prescribed by the target project and should stay.

### Observed work and instruction weaknesses

1. The PR reference is 229 lines / 12,814 characters and contains all modes.
   Title/Description/Explain/Publish sections occupy 3,689 characters that this
   preparation request does not use. Root + PR + writing guidance total 17,469
   characters before ticket/draft/code/result content. Direct action-specific
   loading could avoid roughly 920 input tokens at four characters per token,
   provided it does not add a router read or duplicate shared policy. Keep
   safety-critical shared rules available and preserve publication boundaries.

2. Fine-grained commit-byte rules lack a deterministic Git-to-validator handoff.
   The message is read once for display, then read again by command substitution
   for validation. `msg=$(git show ...)` removes all trailing LFs, while the skill
   permits removing at most one. A local synthetic example with an extra trailing
   blank line fails ix when only one LF is removed, but passes after command
   substitution strips every terminal LF. No malformed captured commit is
   demonstrated; this is a confirmed weakness in the generated method. Prefer
   a small Git-aware native validation path that reads exact stored messages,
   normalizes one permitted LF, and reports SHA/results. Preserve every commit
   in TARGET_REF..HEAD and do not filter by ticket/author. No byte-dump fallback
   or persisted helper script is needed. This native change is a proposal only.

3. The run reads estimate.md, globs ix, and then lists the private tool directory.
   Current PR guidance already excludes estimate and root guidance says to invoke
   tools directly. These are execution departures, not missing prohibitions.
   The empty Glob leads to an incorrect missing-tool conclusion despite the
   executable existing. Direct invocation would remove two lookup calls and the
   unused estimate read would remove a third. Repeating the same bans would grow
   prompt cost without a demonstrated compliance benefit.

4. Source exploration expands to eleven reads/searches, versus five in da22335f.
   The seven source Read calls include three separate excerpts of the same store.
   These six additional source calls explain the call-count increase after other
   changes offset. The skill asks for verified cause, complete work, examples,
   and material-gap resolution, but the agent repeatedly expands its explanation
   into extra feature behavior and edge-case review. Bound the summary to the
   smallest verified causal explanation, use exact imported paths and one
   relevant source range per needed fact, and report actual unresolved concerns.
   Examples should illustrate verified behavior without widening investigation.
   Do not replace this with a fixed call quota or skip genuine material issues.

5. The workflow puts verification after source investigation, then the run spends
   another substantial planning block composing/mapping the draft after checks.
   Clarify that only saving the file must wait for readiness: with nonblocking
   host execution, required checks can run while context review and tentative
   drafting continue. Collect results, perform the final state recheck, and save
   only on success. Reuse gathered facts rather than planning the whole summary
   again after results. Exact timing savings require another measured run.

6. The generic summary field taxonomy differs from the initialized draft's labels,
   leading to repeated mapping decisions. Make the initialized draft authoritative
   for layout; keep writing guidance about content without a second competing
   field schema. Optional notes should add a reviewer-relevant fact once rather
   than repeat implementation details. Preserve the user's readable openings,
   concrete examples, and evidence requirements.

### Recommended order and expected savings

Start with action-specific instruction loading and a deterministic native commit
validation path. Then schedule required checks alongside read-only review/draft
composition where the host supports it. Narrow evidence collection and reduce
format-mapping decisions without growing another list of near-duplicate bans.
The known unnecessary estimate/tool lookups account for three calls (33 to 30
if those alone disappear). Further call reductions from source investigation
and the native handoff are plausible but not yet measured. The approximately
920 input-token opportunity covers unused mode text only. Generated command/
write text fell by just 34 characters, roughly nine tokens; total billed-token
savings cannot be inferred from that category or from thinking durations.

No new instruction/native implementation was applied in this analysis pass.
The native normalization reproduction used only synthetic in-memory values and
read-only ix checks; no repository, Jira, Docker, draft, or PR state changed.
The preceding root, branch/commit, target selection, readability, publication,
and verification requirements remain intact. Keep evaluation on demand.

## Visible PR progress checklist

The user explicitly requested a checklist so they can see where the agent is in
an /inux pr run. Add a compact chat-only progress section to the PR reference:
prefer the host's visible plan/todo UI, fall back to Markdown with concise phase
updates, and update from existing results at phase transitions. Completion must
reflect observed results; blocked work is identified. No extra Git/source/test
commands are needed just to refresh progress.

Preparation tracks inputs/target, Git/commit validation, change explanation,
missing prescribed checks, and final state/save. Publication uses corresponding
saved-draft/readiness, remote/auth, push, PR mutation, and verification phases.
Focused edits/explanations use only relevant steps rather than expanding into
readiness. Keep the checklist out of the PR body. Synchronize the PR reference
to Cursor and shared skills. The requested UI adds some instruction/update text;
no token or runtime saving is claimed for the checklist itself.

The earlier aa2231a9 performance recommendations remain analysis proposals.
Only the explicitly requested progress-checklist addition is implemented here.

PR reference SHA-256: `e9b0cf5fe6945864962bd510b9937ed33017a0a248be1cebab2c0921a1526afd`.


## 2026-10-15: split PR actions and validate Git before investigation

User-supplied run: `/Users/vinodliyanage/.inux/logs/chats/5827e71d-e20c-49bf-a876-fc86870c0bba.md`.
SHA-256: `8425cfbc96e926152e2367394fc8621e4e80158ebfd20a23b3a54e1ac5313c85`.
Request: `/inux pr target release/v2.9.0 add content to pr.md file`.
Exclude chat-save; deduplicate identical exported thinking entries.

The preparation run contains 39 calls: 12 Read, 9 Grep, 13 Shell, 1 tool-schema
lookup, 3 TodoWrite updates, and 1 Write. Eighteen unique task thinking records
span 04:15:43–04:18:08 UTC (145s), with 64.803s summed recorded duration. The
export omits tool results and token usage; these figures are not total tokens,
isolated check duration, or independent proof of reported passing checks. All
four prescribed repository checks were submitted in one assistant message, so
the trace does not demonstrate a missed parallelization opportunity between them.
The reviewed commit changes three files, with 23 insertions and one deletion.

The user requested a `references/pr/` folder and the earlier action-based split,
while explicitly requiring Git validation before technical-summary investigation.
Keep references/pr.md as a small router. Add prepare.md, publish.md, refine.md,
readiness.md, and summary.md under references/pr/. Preparation/publication load
their action and readiness together, then defer summary/writing guidance until
the Git gate passes. Focused edits/explanations retain their lighter workflow.
Update cli.md's convention link to the new readiness reference.

The Git gate resolves the ticket heading, applicable project AGENTS.md rules,
required task-file presence, and developer-supplied destination before checking
local state, branch conventions, fetched target/shared ancestry, and all PR-range
commit messages. Failure preserves the draft and stops full ticket/draft reads,
source exploration, and composition. Message implementation accuracy remains a
diff-review decision after structural validation.

Primary summary evidence is ticket.md, project AGENTS.md, checked commit messages,
and the full merge-base-to-HEAD diff. Reuse this evidence. Follow exact imported
paths/symbols only for a named material missing fact, and stop when the summary's
claims are supported. Keep the initialized draft authoritative for labels/order
and preserve the established readable opening, useful examples, and optional
reviewer-relevant notes.

Consolidate the demonstrated waste into these action-specific instructions:
exclude estimate.md; use one initial porcelain status; read stored commit messages
once, preserve their exact value except at most one terminal LF, and pass them
directly to ix without retyping or lossy command substitution; reserve byte dumps
for actual failures; preserve required-command exit statuses instead of hiding
them behind echo; reuse provided verification commands and perform only a targeted
lookup when translation is needed. Independent checks may overlap read-only
review/composition once their scope is known; every required result and final
state recheck still gates saving/publishing.

Preserve the user's visible progress checklist using an already available plan
UI or Markdown, without a separate tool-schema lookup. No fixed call quota,
QR-specific runtime rule, new helper/native command, or automatic publication is
introduced. Existing drafts and unrelated source changes remain untouched.

Update the existing skill archive configuration with explicit top-level Markdown
and references/pr/*.md entries so nested references retain references/pr/ paths.
Use the documented src/dst/strip_parent behavior:
[GoReleaser archives](https://goreleaser.com/customization/package/archives/).

Preparation reference text (router, action, shared readiness/summary, and writing
reference, excluding unchanged SKILL.md and all evidence/tool outputs) changes
from 15,400 to 13,062 characters (15.2% shorter).
The split increases reference-file reads from two to five; batching and deferred
loading limit round trips and failed-gate work, but no runtime or billed-token
improvement is claimed without a new run.

Validation: all packaged Markdown links/anchors and whitespace pass; frontmatter
and archive YAML parse with Ruby Psych. The standard quick_validate.py cannot run
because PyYAML is absent in system and bundled Python; no dependency is installed.
A local ZIP-layout smoke check using the configured archive globs contains all
18 skill files at their expected paths with matching bytes. This is not an actual
GoReleaser build or release. Routing/publication/edit boundaries were reviewed
statically; no live Jira, Git fetch, Docker check, task draft write, or publication
was exercised. Synchronize only changed skill files to both installed copies,
with baseline-drift guards and backups. Further evaluation remains on demand.

## Reader feedback: simple chronological paragraphs

The user found both the generated summary and an initial rewrite difficult to
understand on first reading. Numbered steps made the sequence clearer, but the
user explicitly wants paragraph explanations. They approved a short paragraph
that first states the user-visible result, optionally illustrates it, then says
how the app makes it happen. The approved Root cause describes the old behavior,
the missing action, and the visible failure in simple sentences.

Replace the corresponding field guidance in references/pr/summary.md and align
assets/pr.md's hints. Require one action or idea per sentence, events in order,
an opening that stands on its own, familiar words, and no repeated facts. Keep
each explanation as a paragraph on the existing label line, without numbered
steps or nested bullets. Preserve the template's field labels, verified evidence,
useful examples, and reviewer-relevant limitations. Do not impose a sentence
quota or add a QR-specific example to the reusable skill.

This is a focused readability change. Shared be-human.md already supports the
style and stays unchanged. Validate the changed Markdown and links, synchronize
both skill files to Codex/shared and Cursor installations, and preserve existing
task drafts. The next normal run can show whether the wording improves; no token
or latency improvement is claimed.
