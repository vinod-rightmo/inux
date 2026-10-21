package commands

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"unicode/utf8"

	"github.com/ericmason/mdadf"
)

// ADF converts a UTF-8 Markdown input file to an ADF output file.
func ADF(args []string) error {
	if len(args) != 2 {
		return errors.New("usage: adf INPUT.md OUTPUT.adf")
	}

	source, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}
	if !utf8.Valid(source) {
		return errors.New("Markdown input must be UTF-8")
	}
	input, err := os.Stat(args[0])
	if err != nil {
		return err
	}
	output, err := os.Stat(args[1])
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if output != nil && os.SameFile(input, output) {
		return errors.New("ADF output must differ from the Markdown input")
	}

	document, err := mdadf.ConvertToDoc(string(source))
	if err != nil {
		return err
	}
	if document.Content == nil {
		document.Content = []mdadf.Node{}
	}
	converted, err := json.Marshal(document)
	if err != nil {
		return err
	}
	file, err := os.CreateTemp(filepath.Dir(args[1]), ".adf-*")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	_, writeErr := file.Write(append(converted, '\n'))
	closeErr := file.Close()
	if writeErr != nil || closeErr != nil {
		return errors.Join(writeErr, closeErr)
	}
	return os.Rename(file.Name(), args[1])
}
