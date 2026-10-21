package app

import (
	"fmt"
	"io"

	"adf/internal/commands"
)

func Run(args []string, out io.Writer) error {
	if len(args) == 1 && (args[0] == "--help" || args[0] == "-h") {
		_, err := fmt.Fprintln(out, "Usage: adf INPUT.md OUTPUT.adf")
		return err
	}
	return commands.ADF(args)
}
