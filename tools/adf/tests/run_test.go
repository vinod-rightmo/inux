package tests

import (
	app "adf/internal"
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func put(t *testing.T, path string, content []byte) {
	t.Helper()
	if err := os.WriteFile(path, content, 0600); err != nil {
		t.Fatal(err)
	}
}

func read(t *testing.T, path string) []byte {
	t.Helper()
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return content
}

func TestConvertsMarkdown(t *testing.T) {
	root := t.TempDir()
	input := filepath.Join(root, "source with spaces.md")
	output := filepath.Join(root, "body.adf")
	markdown := []byte("# Estimate\n\nCafé සිංහල 日本語 😀 **bold** [link](https://example.com)\n\n- item\n\n| Area | Score |\n|---|---|\n| UI | 1 |\n")
	put(t, input, markdown)
	put(t, output, []byte("old output"))
	var stdout bytes.Buffer
	if err := app.Run([]string{input, output}, &stdout); err != nil {
		t.Fatal(err)
	}
	converted := read(t, output)
	var document map[string]any
	if err := json.Unmarshal(converted, &document); err != nil {
		t.Fatal(err)
	}
	if document["type"] != "doc" || document["version"] != float64(1) {
		t.Fatal(document)
	}
	for _, node := range []string{"heading", "table", "bulletList", "strong", "link", "සිංහල"} {
		if !strings.Contains(string(converted), node) {
			t.Fatal("missing ADF content", node)
		}
	}
	if !bytes.Equal(read(t, input), markdown) || stdout.Len() != 0 {
		t.Fatal("conversion changed the input or wrote to stdout")
	}
}

func TestEmptyMarkdown(t *testing.T) {
	root := t.TempDir()
	input, output := filepath.Join(root, "empty.md"), filepath.Join(root, "empty.adf")
	put(t, input, nil)
	if err := app.Run([]string{input, output}, io.Discard); err != nil {
		t.Fatal(err)
	}
	var document map[string]any
	if err := json.Unmarshal(read(t, output), &document); err != nil || document["content"] == nil {
		t.Fatal("empty document must have an array for content", err)
	}
}

func TestInputAliasesAreProtected(t *testing.T) {
	root := t.TempDir()
	input := filepath.Join(root, "source.md")
	markdown := []byte("# Keep this Markdown\n")
	put(t, input, markdown)
	for _, kind := range []string{"same path", "hard link", "symbolic link"} {
		t.Run(kind, func(t *testing.T) {
			output := input
			if kind != "same path" {
				output = filepath.Join(root, kind+".adf")
				link := os.Link
				if kind == "symbolic link" {
					link = os.Symlink
				}
				if err := link(input, output); err != nil {
					t.Skip(err)
				}
			}
			if err := app.Run([]string{input, output}, io.Discard); err == nil {
				t.Fatal("accepted input overwrite")
			}
			if !bytes.Equal(read(t, input), markdown) {
				t.Fatal("changed the Markdown input")
			}
		})
	}
}

func TestInvalidInputPreservesOutput(t *testing.T) {
	root := t.TempDir()
	input, output := filepath.Join(root, "source.md"), filepath.Join(root, "body.adf")
	original := []byte("existing output")
	put(t, output, original)
	for _, content := range [][]byte{nil, {0xff}} {
		if content != nil {
			put(t, input, content)
		}
		if err := app.Run([]string{input, output}, io.Discard); err == nil {
			t.Fatal("accepted missing or invalid UTF-8 input")
		}
		if !bytes.Equal(read(t, output), original) {
			t.Fatal("failed conversion changed the output")
		}
	}
}

func TestArguments(t *testing.T) {
	for _, args := range [][]string{nil, {"input.md"}, {"input.md", "output.adf", "extra"}, {"--version"}} {
		if err := app.Run(args, io.Discard); err == nil {
			t.Fatal("accepted invalid arguments", args)
		}
	}
	for _, flag := range []string{"--help", "-h"} {
		var output bytes.Buffer
		if err := app.Run([]string{flag}, &output); err != nil || output.Len() == 0 {
			t.Fatal("help failed", flag, err)
		}
	}
}
