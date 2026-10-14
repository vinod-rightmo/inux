package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// RecreateWorkspace validates templates and an optional snapshot before replacing .inux/.
func RecreateWorkspace(args []string, expected string, out io.Writer) error {
	if len(args) < 1 || len(args) > 2 {
		return errors.New("use: ix task TEMPLATES_DIR [PROJECT]")
	}
	estimate, err := os.ReadFile(filepath.Join(args[0], "estimate.md"))
	if err != nil {
		return fmt.Errorf("read estimate template: %w", err)
	}
	pr, err := os.ReadFile(filepath.Join(args[0], "pr.md"))
	if err != nil {
		return fmt.Errorf("read PR template: %w", err)
	}
	project := "."
	if len(args) == 2 {
		project = args[1]
	}
	root, directory, err := workspaceDirectory(project)
	if err != nil {
		return err
	}
	if info, err := os.Lstat(directory); err == nil && !info.IsDir() {
		return errors.New(".inux must be a directory, not a link or file")
	} else if err != nil && !os.IsNotExist(err) {
		return err
	}
	if expected != "" {
		current, err := workspaceSnapshot(directory)
		if err != nil {
			return err
		}
		if current != expected {
			return errors.New(".inux changed since snapshot; task not replaced")
		}
	}
	if err := exclude(root); err != nil {
		return err
	}
	if err := os.RemoveAll(directory); err != nil {
		return err
	}
	if err := os.Mkdir(directory, 0700); err != nil {
		return err
	}
	if err := Save(filepath.Join(directory, "estimate.md"), estimate); err != nil {
		return err
	}
	if err := Save(filepath.Join(directory, "pr.md"), pr); err != nil {
		return err
	}
	if err := Save(filepath.Join(directory, "ticket.md"), nil); err != nil {
		return err
	}
	return json.NewEncoder(out).Encode(map[string]any{
		"ok": true, "directory": directory, "drafts": []string{"estimate.md", "pr.md"},
		"files": []string{"ticket.md", "estimate.md", "pr.md"},
	})
}

func exclude(project string) error {
	gitDir := filepath.Join(project, ".git")
	info, err := os.Stat(gitDir)
	if os.IsNotExist(err) {
		return nil // Plain directories also support Inux.
	}
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return errors.New("MVP requires an ordinary Git checkout; .git files are not supported")
	}
	path := filepath.Join(gitDir, "info", "exclude")
	content, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	for _, line := range strings.Split(string(content), "\n") {
		if strings.TrimSuffix(strings.TrimSpace(line), "/") == ".inux" {
			return nil
		}
	}
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}
	if len(content) > 0 && content[len(content)-1] != '\n' {
		content = append(content, '\n')
	}
	return Save(path, append(content, []byte(".inux/\n")...))
}
