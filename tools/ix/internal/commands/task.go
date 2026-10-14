package commands

import (
	"io"

	"ix/internal/shared"
)

// Task handles "ix task TEMPLATE [PROJECT]"; PROJECT defaults to the current directory.
// Replaces the entire .inux/ directory and copies TEMPLATE to estimate.md.
// Example: ix task /path/to/assets/estimate.md /path/to/project.
func Task(args []string, out io.Writer) error {
	return shared.RecreateWorkspace(args, out)
}
