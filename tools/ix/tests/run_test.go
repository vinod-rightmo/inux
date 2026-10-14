package tests

import (
	"bytes"
	"encoding/json"
	app "ix/internal"
	"os"
	"os/exec"
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

func git(t *testing.T, args ...string) string {
	t.Helper()
	output, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		t.Fatalf("git: %v: %s", err, output)
	}
	return strings.TrimSpace(string(output))
}

func TestLifecycle(t *testing.T) {
	root := t.TempDir()
	template := estimateTemplate(t)
	git(t, "init", "-q", root)
	ignore := filepath.Join(root, ".git", "info", "exclude")
	put(t, ignore, "# Keep existing rule")
	put(t, filepath.Join(root, ".gitignore"), "keep/\n")
	call(t, "task", template, root)
	dir := filepath.Join(root, ".inux")
	for _, name := range []string{"estimate.md", "pr.md", "ticket.md"} {
		put(t, filepath.Join(dir, name), "old task")
	}
	if err := os.Mkdir(filepath.Join(dir, "old-folder"), 0700); err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(dir, "old-folder", "old.md"), "old")
	put(t, filepath.Join(root, "keep.txt"), "keep")
	call(t, "task", template, root)
	files, err := os.ReadDir(dir)
	if err != nil || len(files) != 1 || files[0].Name() != "estimate.md" {
		t.Fatal("task did not replace the complete workspace", files, err)
	}
	if get(t, filepath.Join(dir, "estimate.md")) != get(t, template) {
		t.Fatal("task did not copy the template")
	}
	if get(t, ignore) != "# Keep existing rule\n.inux/\n" || get(t, filepath.Join(root, ".gitignore")) != "keep/\n" {
		t.Fatal("Git exclusion changed incorrectly")
	}
	if get(t, filepath.Join(root, "keep.txt")) != "keep" {
		t.Fatal("task changed project files")
	}
}

func TestFreshTaskAndNoGit(t *testing.T) {
	t.Setenv("PATH", "")
	root := t.TempDir()
	call(t, "task", estimateTemplate(t), root)
	if !strings.Contains(get(t, filepath.Join(root, ".inux", "estimate.md")), "Final Estimate:") {
		t.Fatal("fresh task failed")
	}
	if _, err := os.Stat(filepath.Join(root, ".inux", "pr.md")); !os.IsNotExist(err) {
		t.Fatal("task created a PR draft", err)
	}
}

func TestGitFileUnsupported(t *testing.T) {
	root := t.TempDir()
	put(t, filepath.Join(root, ".git"), "gitdir: elsewhere\n")
	if err := os.Mkdir(filepath.Join(root, ".inux"), 0700); err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(root, ".inux", "ticket.md"), "existing task")
	if err := app.Run([]string{"task", estimateTemplate(t), root}, &bytes.Buffer{}); err == nil {
		t.Fatal("accepted unsupported .git file")
	}
	if get(t, filepath.Join(root, ".inux", "ticket.md")) != "existing task" {
		t.Fatal("unsupported checkout lost the existing task")
	}
}

func TestDirectoryLinksAndFileAliases(t *testing.T) {
	root, outside := t.TempDir(), t.TempDir()
	put(t, filepath.Join(outside, "keep"), "untouched")
	if err := os.Symlink(outside, filepath.Join(root, ".inux")); err != nil {
		t.Skip(err)
	}
	template := estimateTemplate(t)
	if err := app.Run([]string{"task", template, root}, &bytes.Buffer{}); err == nil {
		t.Fatal("accepted directory symlink")
	}
	os.Remove(filepath.Join(root, ".inux"))
	call(t, "task", template, root)
	os.Remove(filepath.Join(root, ".inux", "estimate.md"))
	if err := os.Link(filepath.Join(outside, "keep"), filepath.Join(root, ".inux", "estimate.md")); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, filepath.Join(root, ".inux", "outside")); err != nil {
		t.Fatal(err)
	}
	call(t, "task", template, root)
	if get(t, filepath.Join(outside, "keep")) != "untouched" {
		t.Fatal("task followed a link outside .inux")
	}
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
	for _, args := range [][]string{{"unknown"}, {"adf"}, {"task"}, {"task", "one", "two", "three"}, {"init", "template"}, {"reset", "template"}, {"remove"}, {"adf", "a", "b", "c"}} {
		if err := app.Run(args, &bytes.Buffer{}); err == nil {
			t.Fatal("accepted invalid arguments", args)
		}
	}
}

func estimateTemplate(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "estimate template.md")
	put(t, path, "Final Estimate:      SP\n\nCustom Café template\n")
	return path
}

func TestTemplateInput(t *testing.T) {
	root := t.TempDir()
	template := estimateTemplate(t)
	call(t, "task", template, root)
	estimate := filepath.Join(root, ".inux", "estimate.md")
	put(t, template, "Updated template 日本語\n")
	call(t, "task", template, root)
	if get(t, estimate) != get(t, template) || get(t, template) != "Updated template 日本語\n" {
		t.Fatal("task did not copy the template or changed its source")
	}
	put(t, filepath.Join(root, ".inux", "pr.md"), "reviewed draft")
	put(t, filepath.Join(root, ".inux", "ticket.md"), "existing task")
	missing := filepath.Join(t.TempDir(), "missing.md")
	var out bytes.Buffer
	if err := app.Run([]string{"task", missing, root}, &out); err == nil || out.Len() != 0 {
		t.Fatal("accepted missing template")
	}
	if get(t, estimate) != "Updated template 日本語\n" || get(t, filepath.Join(root, ".inux", "pr.md")) != "reviewed draft" || get(t, filepath.Join(root, ".inux", "ticket.md")) != "existing task" {
		t.Fatal("missing template changed existing task")
	}
	fresh := t.TempDir()
	if err := app.Run([]string{"task", missing, fresh}, &bytes.Buffer{}); err == nil {
		t.Fatal("accepted missing template for a new project")
	}
	if _, err := os.Stat(filepath.Join(fresh, ".inux")); !os.IsNotExist(err) {
		t.Fatal("missing template created a workspace", err)
	}
}
