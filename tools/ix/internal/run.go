package app

import (
	"errors"
	"fmt"
	"io"

	"ix/internal/commands"
)

const help = `ix --version                              Show the installed CLI version.
ix adf INPUT [OUTPUT]                       Convert Markdown; omit OUTPUT to print ADF JSON.
No task files, downloads, authentication, or publishing.
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
	case "adf":
		return commands.ADF(args[1:], out)
	default:
		return errors.New("unknown command; see --help")
	}
}
