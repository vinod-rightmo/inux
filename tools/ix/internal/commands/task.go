package commands

import (
	"encoding/hex"
	"errors"
	"io"

	"ix/internal/shared"
)

// Task snapshots or replaces .inux/; PROJECT defaults to the current directory.
func Task(args []string, out io.Writer) error {
	if len(args) > 0 && args[0] == "--snapshot" {
		if len(args) > 2 {
			return errors.New("use: ix task --snapshot [PROJECT]")
		}
		project := "."
		if len(args) == 2 {
			project = args[1]
		}
		return shared.SnapshotWorkspace(project, out)
	}
	expected := ""
	if len(args) > 0 && args[0] == "--expect" {
		if len(args) < 3 {
			return errors.New("use: ix task --expect SNAPSHOT TEMPLATES_DIR [PROJECT]")
		}
		token, err := hex.DecodeString(args[1])
		if err != nil || len(token) != 32 {
			return errors.New("expected snapshot must be a SHA-256 token from ix task --snapshot")
		}
		expected, args = args[1], args[2:]
	}
	return shared.RecreateWorkspace(args, expected, out)
}
