package tests

import (
	"bytes"
	"encoding/json"
	app "ix/internal"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func call(t *testing.T, args ...string) []byte {
	t.Helper()
	var out bytes.Buffer
	if err := app.Run(args, &out); err != nil {
		t.Fatal(err)
	}
	return out.Bytes()
}

func put(t *testing.T, path, value string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(value), 0600); err != nil {
		t.Fatal(err)
	}
}

func get(t *testing.T, path string) string {
	t.Helper()
	value, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(value)
}

func TestADFAndInputProtection(t *testing.T) {
	root := t.TempDir()
	input, output := filepath.Join(root, "source with spaces.md"), filepath.Join(root, "body.json")
	markdown := "# Estimate\n\nCafé සිංහල 日本語 😀 **bold** [link](https://example.com)\n\n- item\n\n| Area | Score |\n|---|---|\n| UI | 1 |\n"
	put(t, input, markdown)
	converted := call(t, "adf", input)
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
	call(t, "adf", input, output)
	if get(t, output) != string(converted) || get(t, input) != markdown {
		t.Fatal("conversion changed content")
	}
	alias := filepath.Join(root, "alias.md")
	if err := os.Link(input, alias); err != nil {
		t.Fatal(err)
	}
	for _, target := range []string{input, alias} {
		if err := app.Run([]string{"adf", input, target}, &bytes.Buffer{}); err == nil {
			t.Fatal("accepted input overwrite")
		}
	}
	put(t, input, string([]byte{0xff}))
	if err := app.Run([]string{"adf", input, output}, &bytes.Buffer{}); err == nil {
		t.Fatal("accepted invalid UTF-8")
	}
	if get(t, output) != string(converted) {
		t.Fatal("failed conversion changed output")
	}
	put(t, input, "")
	if err := json.Unmarshal(call(t, "adf", input), &document); err != nil || document["content"] == nil {
		t.Fatal("empty document must have an array for content", err)
	}
}

func TestInvalidArguments(t *testing.T) {
	for _, args := range [][]string{{"unknown"}, {"check-branch", ".*", "branch"}, {"check-commit", ".*", "message"}, {"adf"}, {"task"}, {"task", "--snapshot"}, {"task", "--expect", "token", "templates"}, {"task", "templates", "project"}, {"init", "template"}, {"reset", "template"}, {"remove"}, {"adf", "a", "b", "c"}} {
		if err := app.Run(args, &bytes.Buffer{}); err == nil {
			t.Fatal("accepted invalid arguments", args)
		}
	}
}

// Removed task commands must never create a workspace or change project files.
func TestRemovedTaskPreservesProject(t *testing.T) {
	root := t.TempDir()
	put(t, filepath.Join(root, "keep.md"), "developer draft")
	for _, args := range [][]string{
		{"task", root, root},
		{"task", "--snapshot", root},
		{"task", "--expect", strings.Repeat("a", 64), root, root},
	} {
		var out bytes.Buffer
		if err := app.Run(args, &out); err == nil || out.Len() != 0 {
			t.Fatal("removed command accepted or reported success", args)
		}
	}
	entries, err := os.ReadDir(root)
	if err != nil || len(entries) != 1 || entries[0].Name() != "keep.md" || get(t, filepath.Join(root, "keep.md")) != "developer draft" {
		t.Fatal("removed task command changed the project", entries, err)
	}
}
