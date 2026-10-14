package app

import (
	"errors"
	"fmt"
	"io"

	"ix/internal/commands"
)

const help = `ix --version                              Show the installed CLI version.
ix task --snapshot [PROJECT]              Read a compact .inux state token without writing.
ix task [--expect SNAPSHOT] TEMPLATES_DIR [PROJECT] Replace .inux/ with an empty ticket and both drafts.
ix adf INPUT [OUTPUT]                       Convert Markdown; omit OUTPUT to print ADF JSON.
ix check-branch PATTERN NAME [NAME...]      Check branch names against one project Go regex.
ix check-commit PATTERN MESSAGE [MESSAGE...] Check commit messages against one project Go regex.
PROJECT defaults to the current directory. No downloads, authentication, or publishing.
`

// Version is set from the release version at build time.
var Version = "dev"

// Run dispatches a command and writes its success output.
func Run(args []string, out io.Writer) error {
	if len(args) > 0 && args[0] == "--version" {
		if len(args) != 1 {
			return errors.New("use: ix --version")
		}
		_, err := fmt.Fprintf(out, "ix %s\n", Version)
		return err
	}
	if len(args) == 0 || args[0] == "--help" || args[0] == "help" {
		_, err := fmt.Fprint(out, help)
		return err
	}
	switch args[0] {
	case "task":
		return commands.Task(args[1:], out)
	case "adf":
		return commands.ADF(args[1:], out)
	case "check-branch":
		return commands.CheckBranch(args[1:], out)
	case "check-commit":
		return commands.CheckCommit(args[1:], out)
	default:
		return errors.New("unknown command; see --help")
	}
}
