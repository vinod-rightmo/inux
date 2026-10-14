package shared

import (
	"errors"
	"os"
	"path/filepath"
)

// Save replaces a file using a temporary file in the same directory.
func Save(path string, content []byte) error {
	file, err := os.CreateTemp(filepath.Dir(path), ".inux-write-*")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	_, writeErr := file.Write(content)
	closeErr := file.Close()
	if writeErr != nil || closeErr != nil {
		return errors.Join(writeErr, closeErr)
	}
	return os.Rename(file.Name(), path)
}
