package commands

import (
	"io"

	"ix/internal/shared"
)

// CheckBranch handles "ix check-branch PATTERN NAME [NAME...]" with a project-specific Go regex.
// Example: ix check-branch 'bugfix/ABC-[0-9]+-[a-z0-9-]+' bugfix/ABC-123-fix bugfix/ABC-124-test
func CheckBranch(args []string, out io.Writer) error {
	return shared.CheckConvention(args, out, "ix check-branch PATTERN NAME [NAME...]")
}
