package commands

import (
	"encoding/json"
	"errors"
	"os"
	"unicode/utf8"

	"adf/internal/shared"
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
	return shared.Save(args[1], append(converted, '\n'))
}
