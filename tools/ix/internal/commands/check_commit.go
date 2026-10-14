package commands

import (
	"io"

	"ix/internal/shared"
)

// CheckCommit handles "ix check-commit PATTERN MESSAGE [MESSAGE...]" with a project-specific Go regex.
// Example: ix check-commit '^OQG3-3776: [A-Z][^\r\n]+$' 'OQG3-3776: Fix home step centering'
// Pass each subject or complete message as one argument; every entire value is checked.
func CheckCommit(args []string, out io.Writer) error {
	return shared.CheckConvention(args, out, "ix check-commit PATTERN MESSAGE [MESSAGE...]")
}
