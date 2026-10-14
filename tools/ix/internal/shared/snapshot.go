package shared

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func workspaceDirectory(project string) (string, string, error) {
	root, err := filepath.Abs(project)
	if err != nil {
		return "", "", err
	}
	info, err := os.Stat(root)
	if err != nil {
		return "", "", err
	}
	if !info.IsDir() {
		return "", "", errors.New("project must be a directory")
	}
	return root, filepath.Join(root, ".inux"), nil
}

// SnapshotWorkspace returns an opaque, project-bound token without changing files.
func SnapshotWorkspace(project string, out io.Writer) error {
	_, directory, err := workspaceDirectory(project)
	if err != nil {
		return err
	}
	token, err := workspaceSnapshot(directory)
	if err != nil {
		return err
	}
	return json.NewEncoder(out).Encode(map[string]any{"ok": true, "snapshot": token})
}

func workspaceSnapshot(directory string) (string, error) {
	digest := sha256.New()
	fmt.Fprintf(digest, "%q\n", directory)
	info, err := os.Lstat(directory)
	if os.IsNotExist(err) {
		io.WriteString(digest, "absent\n")
		return fmt.Sprintf("%x", digest.Sum(nil)), nil
	}
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		return "", errors.New(".inux must be a directory, not a link or file")
	}
	err = filepath.WalkDir(directory, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(directory, path)
		if err != nil {
			return err
		}
		fmt.Fprintf(digest, "%q %s ", relative, info.Mode())
		switch {
		case info.IsDir():
			fmt.Fprintln(digest)
		case info.Mode()&os.ModeSymlink != 0:
			target, err := os.Readlink(path)
			if err != nil {
				return err
			}
			fmt.Fprintf(digest, "%q\n", target)
		case info.Mode().IsRegular():
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			content := sha256.New()
			_, readErr := io.Copy(content, file)
			if err := errors.Join(readErr, file.Close()); err != nil {
				return err
			}
			fmt.Fprintf(digest, "%x\n", content.Sum(nil))
		default:
			return fmt.Errorf("unsupported entry in .inux: %s", relative)
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", digest.Sum(nil)), nil
}
