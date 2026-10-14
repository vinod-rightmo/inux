package commands

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"unicode/utf8"

	"github.com/ericmason/mdadf"
	"ix/internal/shared"
)

// ADF handles "ix adf INPUT [OUTPUT]", converting UTF-8 Markdown to ADF JSON.
// Example: ix adf estimate.md body.json; omit body.json to print ADF to stdout.
// OUTPUT must differ from INPUT; the Markdown source is preserved.
func ADF(args []string, out io.Writer) error {
	if len(args) < 1 || len(args) > 2 {
		return errors.New("use: ix adf INPUT [OUTPUT]")
	}
	source, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}
	if !utf8.Valid(source) {
		return errors.New("Markdown must be UTF-8")
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
	converted = append(converted, '\n')
	if len(args) == 1 {
		_, err = out.Write(converted)
		return err
	}
	input, inputErr := os.Stat(args[0])
	output, _ := os.Stat(args[1])
	if inputErr != nil {
		return inputErr
	}
	if output != nil && os.SameFile(input, output) {
		return errors.New("ADF output must differ from the Markdown input")
	}
	if err := shared.Save(args[1], converted); err != nil {
		return err
	}
	return json.NewEncoder(out).Encode(map[string]any{"ok": true, "output": args[1]})
}
