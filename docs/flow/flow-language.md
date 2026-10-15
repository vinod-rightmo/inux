# Flow — simple commands, plain-English tasks

Group each workflow under a Markdown heading such as `## Flow: Ask and print`,
followed by one fenced `flow` block. The heading names the workflow; the fence
marks its executable instructions. Write one instruction per line.

A file may contain several named workflows. Names must be unique within the
file. Execute only the workflow selected by name; if selection is ambiguous,
ask which one to run. Do not automatically run other blocks or prose.

Use three backticks for ordinary flow blocks. When a flow contains an inline
JavaScript fence, use four backticks for the outer flow and three for JavaScript.
A referenced .js file needs no nested fence.

## Commands

| Command | Meaning | Example |
| --- | --- | --- |
| `ASK` | Ask and wait; optionally name the answer. | `ASK "Your name?" AS name` |
| `SET` | Store or update a named value. | `SET reviewed = reviewed + 1` |
| `THINK` | Assess available information and save a short conclusion. | `THINK AS advice` … `END THINK` |
| `DO` | Perform a plain-English task. | `DO Print name in the terminal.` |
| `RUN` | Execute an exact command, inline code block, or script file. | See command and code forms below. |
| `CHECK` | Verify a condition; stop if false or unverifiable. | `CHECK reviewed == 3` |
| `IF` / `ELSE` | Execute only the selected branch. | `IF accepted == 3` |
| `REPEAT` | Repeat a block a bounded number of times. | `REPEAT 3 TIMES` |
| `ASYNC` | Start a named background task and continue the main flow. | `ASYNC github_setup` … `END ASYNC` |
| `WAIT` | Wait until a named task finishes successfully. | `WAIT github_setup` |
| `STOP` | End the entire run early with a reason. | `STOP "Cancelled by the user."` |
| `END` | Finish the program or close a named block. | `END`, `END IF`, `END REPEAT`, `END THINK`, `END ASYNC` |

Blank lines and lines starting with `#` are ignored. Indentation is for
readability; named END forms determine block boundaries. ELSE belongs to
the nearest unclosed IF, with at most one ELSE. Bare END finishes the program
outside all blocks. It is required and must be the final instruction.
`END with the PR link.` is also allowed and specifies what to report.

## Values and conditions

- ASK takes one quoted question. `AS name` stores the raw answer as a string,
  including an empty answer. Without AS, store it as `answer`. Ask every time
  the instruction is reached, including inside loops; never silently reuse an answer.
- SET creates or updates a value. Evaluate the right-hand side before updating.
  Names begin with a lowercase letter and contain lowercase letters, digits,
  or underscores. Unknown names are errors.
- Expressions support JSON literals, names, fields such as `stdout of result`,
  numeric `+` and `-`, and same-type comparisons `==`, `!=`, `<`, `<=`, `>`, `>=`.
  No implicit conversions or arbitrary expression execution.
- Built-ins: `trim(text)` removes surrounding whitespace; `json(text)` parses
  one complete JSON value, rejecting malformed or trailing content.
- `value is empty` and `value is not empty` test string or array length. Other
  types are errors. IF expressions must produce Booleans. Missing evidence is
  an error, not false and not a reason to select ELSE.
- Use `field of object` to read an existing field, such as `reason of review`.
  Dotted field access remains supported, but prefer the sentence form.
- Use `object is flag` for a Boolean field: `review is valid` reads review's
  `valid` field and succeeds only when it is Boolean true. `review is not valid`
  succeeds only when that field is Boolean false. A missing object, missing
  field, or non-Boolean value is an error; never guess its meaning. The reserved
  `is empty` and `is not empty` forms keep their length-check meaning.
- Descriptive CHECK instructions must resolve to observable Booleans through
  tools or supplied executable validators. Missing rules or validators stop
  execution. Do not invent conventions or judge compliance by feel.
- REPEAT accepts an integer literal or numeric variable from 1 to 100. Capture
  the count on entry. Iterations share named values. Nested blocks are allowed.
  Stop if the whole run exceeds 1,000 executed instructions. Count each reached
  command or block delimiter, including every iteration and END REPEAT; ignore
  code lines, THINK body prose, blank lines, and unchosen branches.

## Background tasks

### Flow: Check tools asynchronously

```flow
ASYNC node_check
    RUN `node --version` AS version
END ASYNC

ASYNC git_check
    RUN `git --version` AS version
END ASYNC

DO Print "Checking tools..." in the terminal.
WAIT node_check
WAIT git_check
END with both task results.
```

ASYNC launches its whole body as one task. END ASYNC closes that body's
definition; it does not wait. The main flow advances once the task is scheduled,
while instructions inside each task remain sequential. WAIT marks a dependency:
do not execute the following main-flow instruction until that task succeeds.
Other tasks can keep running while the main flow waits.

Task names must be unique and distinct from variable names. For now, ASYNC
blocks appear only at the top level; their bodies may contain IF, REPEAT, and
THINK, but not nested ASYNC, WAIT, ASK, or a final END.

Each task receives a copy of the variables available when it starts. Its writes
stay local, so the two `version` outputs above do not conflict. The task handle
tracks `status`; after successful WAIT, `result of task_name` holds the task's
named output values and its execution evidence. Main-flow variables are not
silently overwritten. Tasks cannot read each other's pending outputs.

Tasks must use independent write destinations and temporary files. Validate
dependencies and destinations before launch; do not schedule tasks that need
each other's output or modify the same file. The 1,000-instruction budget is
shared across the main flow and all tasks, not multiplied per task.

A failed task stops scheduling further instructions once its failure is known.
A failed WAIT never permits dependent work. Let in-flight operations settle
safely, perform explicitly required workflow cleanup, and report partial completion; do not replay
or automatically roll back actions. STOP also stops the whole run. Final END
waits for outstanding tasks and succeeds only if all of them succeeded.

This requires a host that supports background execution or a scheduler able
to interleave task instructions while tools run. Track real job/session handles.
If the capability is unavailable, report blocked before launching tasks; do
not silently run them sequentially and claim they were asynchronous.

## Thinking blocks

### Flow: Suggest next step

```flow
ASK "Describe the problem:" AS problem
THINK AS recommendation
    Consider problem and identify the smallest useful next step.
    Give a brief recommendation and mention any material uncertainty.
END THINK
DO Print recommendation in the terminal.
END
```

THINK takes a prose body ending at END THINK. AS is optional; without it,
store the result as `conclusion`. Store only a concise conclusion or
recommendation as a string. Keep internal reasoning private.

Use only information already available in the program. THINK does not run
tools, ask questions, change other variables, or execute instructions inside
its body. Nested blocks inside THINK are unsupported. Use ASK or RUN before
THINK when more information is needed. It may appear inside IF or REPEAT.

A THINK result is a judgment, not validation evidence. It cannot turn a failed
CHECK into a pass or authorize an action. Use CHECK or a script to verify any
factual requirement before acting on the conclusion.

## Running commands and code

### Flow: Check Node.js

```flow
RUN `node --version` AS runtime
END
```

RUN stores an object with `stdout`, `stderr`, and integer `exit_code`.
AS is optional; without it the result is named `result`. Processes use the
host-supplied working directory and a 30-second timeout. If the host cannot
enforce the timeout, stop before running.

Backticks delimit command text. Parse quoted arguments and pass them to the
process API with `shell=false`. Do not expand variables, pipes, redirects, or
shell substitutions. Use an inline code block or script file for dynamic arguments or logic.
Never concatenate user values into executable source.

An attached JavaScript block uses this form. This fragment expects `name` to
be supplied as an initial variable:

### Flow: Greet user

````flow
RUN JavaScript using Node.js AS greeting
```js
const vars = JSON.parse(process.env.FLOW_VARS);
console.log("Hello, " + vars.name);
```
END
````

The code fence belongs to the preceding RUN; together they form one instruction.
Strip only the shared Markdown indentation, preserving relative code indentation.
Run the source unchanged in a fresh Node.js process. Pass a JSON snapshot of
current variables in the `FLOW_VARS` environment variable. Do not substitute
values into source. Capture the result after completion; JavaScript cannot
update program variables directly.

Alternatively, reference a JavaScript file:

### Flow: Validate commit subject

```flow
RUN "check-commit-subject.js" using Node.js AS validation
END
```

Resolve a relative file path against the folder containing the program's
Markdown file, not the process working directory. Absolute paths are also
allowed. If the program was pasted without a source path, require an explicit
program folder before resolving relative paths. Do not guess the folder.

Execute the existing file unchanged with Node.js using its resolved absolute
path. Keep the host-supplied working directory, 30-second timeout,
FLOW_VARS input, and stdout/stderr/exit_code result contract. A file RUN
has no attached code block. Missing or unreadable files stop the run; do not
create, download, rewrite, or substitute a missing script automatically.

Any nonzero process exit stops the run immediately, before a later CHECK or IF.
Arithmetic, comparisons, and JSON parsing must be evaluated in code. DO allows
judgment within its stated task; it cannot change program control flow.

## Flow: Ask and print

```flow
ASK "What would you like to print?" AS message
DO Print message in the terminal.
END
```

## Flow: Publish draft PR

Supply the repository folder as the working folder, GitHub repository, base
branch, final PR title/body, and executable convention checkers. The current
GitHub base commit must be available locally; the head must already be pushed.

```flow
RUN `git symbolic-ref --quiet --short HEAD` AS branch_result
SET branch = trim(branch_result.stdout)
RUN `git rev-parse --verify HEAD` AS head_result
SET head = trim(head_result.stdout)
CHECK Every commit message in base..head passes the supplied commit-message checker, using the current GitHub base commit.
CHECK branch passes the supplied branch-name checker.
CHECK The working tree is clean, including staged changes and untracked files.
CHECK No open PR already exists for this head and base.
CHECK The local branch and HEAD are unchanged, the GitHub head matches head, and the GitHub base is unchanged since validation.
DO Publish a draft PR using gh with the supplied repository, base, title, and body and branch as the head branch.
CHECK The published PR is open, is a draft, targets the supplied base, and contains head as its HEAD commit.
END with the PR link.
```

For publication, use explicit `gh pr create --repo`, `--base`, `--head`,
`--draft`, `--title`, and `--body` arguments. Pass values as literal arguments.
Do not push, fork, edit, commit, stash, fetch, or repair as implicit preparation.
Races between validation and publication remain possible. These checks do not
establish code correctness or replace lasting repository controls.

## Interpreter prompt

Use the shared [Flow interpreter prompt](../../skills/inux/references/flow-interpreter-prompt.md)
with the selected named flow. That file contains the execution instructions;
workflow files reference it instead of repeating interpreter behavior.

## More advanced example and limits

[Review three commit subjects](advanced-commit-review.md) exercises every command,
including a JavaScript file, mutable counters, and a THINK block for advice. Its policy is explicit and
illustrative; it does not change or validate a real Git repository.

Plain-English DO tasks still require interpretation. A strict prompt cannot
guarantee deterministic execution. Hard guarantees require a code runner that
controls actions and evaluates checks. This document defines the language and
prompt; it does not implement such a runner.
