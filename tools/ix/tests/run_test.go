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
	template := taskTemplates(t)
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
	if err != nil || len(files) != 3 || files[0].Name() != "estimate.md" || files[1].Name() != "pr.md" || files[2].Name() != "ticket.md" {
		t.Fatal("task did not replace the complete workspace", files, err)
	}
	for _, name := range []string{"estimate.md", "pr.md"} {
		if get(t, filepath.Join(dir, name)) != get(t, filepath.Join(template, name)) {
			t.Fatal("task did not copy the template", name)
		}
	}
	if get(t, filepath.Join(dir, "ticket.md")) != "" {
		t.Fatal("task did not reset ticket to an empty file")
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
	call(t, "task", taskTemplates(t), root)
	if !strings.Contains(get(t, filepath.Join(root, ".inux", "estimate.md")), "Final Estimate:") {
		t.Fatal("fresh task failed")
	}
	if !strings.Contains(get(t, filepath.Join(root, ".inux", "pr.md")), "Pull Request") {
		t.Fatal("task did not create the PR draft")
	}
	if get(t, filepath.Join(root, ".inux", "ticket.md")) != "" {
		t.Fatal("task must create an empty ticket for the agent")
	}
}

func TestGitFileUnsupported(t *testing.T) {
	root := t.TempDir()
	put(t, filepath.Join(root, ".git"), "gitdir: elsewhere\n")
	if err := os.Mkdir(filepath.Join(root, ".inux"), 0700); err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(root, ".inux", "ticket.md"), "existing task")
	if err := app.Run([]string{"task", taskTemplates(t), root}, &bytes.Buffer{}); err == nil {
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
	template := taskTemplates(t)
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
	for _, args := range [][]string{{"unknown"}, {"check-branch", ".*", "branch"}, {"check-commit", ".*", "message"}, {"adf"}, {"task"}, {"task", "one", "two", "three"}, {"init", "template"}, {"reset", "template"}, {"remove"}, {"adf", "a", "b", "c"}} {
		if err := app.Run(args, &bytes.Buffer{}); err == nil {
			t.Fatal("accepted invalid arguments", args)
		}
	}
}

func taskTemplates(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "templates with spaces")
	if err := os.Mkdir(path, 0700); err != nil {
		t.Fatal(err)
	}
	put(t, filepath.Join(path, "estimate.md"), "Final Estimate:      SP\n\nCustom Café template\n")
	put(t, filepath.Join(path, "pr.md"), "# Pull Request\r\n\r\nCustom 日本語 template\r\n")
	return path
}

func TestTemplateInput(t *testing.T) {
	root := t.TempDir()
	template := taskTemplates(t)
	call(t, "task", template, root)
	estimate := filepath.Join(root, ".inux", "estimate.md")
	put(t, filepath.Join(template, "estimate.md"), "Updated template 日本語\n")
	call(t, "task", template, root)
	if get(t, estimate) != get(t, filepath.Join(template, "estimate.md")) || get(t, filepath.Join(template, "estimate.md")) != "Updated template 日本語\n" {
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

func TestBothTemplatesRequiredBeforeReplacement(t *testing.T) {
	for _, invalid := range []string{"estimate.md", "pr.md", "pr-directory"} {
		t.Run(invalid, func(t *testing.T) {
			root := t.TempDir()
			git(t, "init", "-q", root)
			templates := taskTemplates(t)
			call(t, "task", templates, root)
			for _, name := range []string{"estimate.md", "pr.md", "ticket.md"} {
				put(t, filepath.Join(root, ".inux", name), "reviewed "+name)
			}
			exclude := filepath.Join(root, ".git", "info", "exclude")
			originalExclude := get(t, exclude)
			missing := invalid
			if invalid == "pr-directory" {
				missing = "pr.md"
			}
			if err := os.Remove(filepath.Join(templates, missing)); err != nil {
				t.Fatal(err)
			}
			if invalid == "pr-directory" {
				if err := os.Mkdir(filepath.Join(templates, missing), 0700); err != nil {
					t.Fatal(err)
				}
			}
			var out bytes.Buffer
			if err := app.Run([]string{"task", templates, root}, &out); err == nil || out.Len() != 0 {
				t.Fatal("accepted invalid template or reported success", invalid)
			}
			for _, name := range []string{"estimate.md", "pr.md", "ticket.md"} {
				if get(t, filepath.Join(root, ".inux", name)) != "reviewed "+name {
					t.Fatal("invalid template changed current task", name)
				}
			}
			if get(t, exclude) != originalExclude {
				t.Fatal("invalid template changed Git exclusion")
			}
			fresh := t.TempDir()
			if err := app.Run([]string{"task", templates, fresh}, &out); err == nil {
				t.Fatal("accepted invalid templates for fresh task")
			}
			if _, err := os.Stat(filepath.Join(fresh, ".inux")); !os.IsNotExist(err) {
				t.Fatal("invalid templates created a task", err)
			}
		})
	}
}

func TestTemplatesLoadedBeforeRemovingTheirDirectory(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, ".inux")
	if err := os.Mkdir(dir, 0700); err != nil {
		t.Fatal(err)
	}
	estimate, pr := "Final Estimate: Café\n", "# Pull Request\r\n日本語\r\n"
	put(t, filepath.Join(dir, "estimate.md"), estimate)
	put(t, filepath.Join(dir, "pr.md"), pr)
	var result struct {
		OK        bool     `json:"ok"`
		Directory string   `json:"directory"`
		Drafts    []string `json:"drafts"`
		Files     []string `json:"files"`
	}
	if err := json.Unmarshal(call(t, "task", dir, root), &result); err != nil {
		t.Fatal(err)
	}
	if !result.OK || result.Directory != dir || len(result.Drafts) != 2 || strings.Join(result.Files, ",") != "ticket.md,estimate.md,pr.md" {
		t.Fatal("missing successful initialization result", result)
	}
	for _, name := range result.Drafts {
		expected, exists := map[string]string{"estimate.md": estimate, "pr.md": pr}[name]
		if !exists || get(t, filepath.Join(dir, name)) != expected {
			t.Fatal("created draft differs from source bytes", name)
		}
	}
}

func snapshot(t *testing.T, project string) string {
	t.Helper()
	var result struct {
		OK       bool   `json:"ok"`
		Snapshot string `json:"snapshot"`
	}
	if err := json.Unmarshal(call(t, "task", "--snapshot", project), &result); err != nil {
		t.Fatal(err)
	}
	if !result.OK || len(result.Snapshot) != 64 {
		t.Fatal("missing snapshot token", result)
	}
	return result.Snapshot
}

func TestGuardedTask(t *testing.T) {
	for _, state := range []string{"absent", "empty", "populated"} {
		t.Run(state, func(t *testing.T) {
			root, templates := t.TempDir(), taskTemplates(t)
			dir := filepath.Join(root, ".inux")
			if state != "absent" {
				if err := os.Mkdir(dir, 0700); err != nil {
					t.Fatal(err)
				}
			}
			if state == "populated" {
				put(t, filepath.Join(dir, "ticket.md"), "old task")
			}
			token := snapshot(t, root)
			if state == "absent" {
				if _, err := os.Stat(dir); !os.IsNotExist(err) {
					t.Fatal("snapshot created a task", err)
				}
			} else if state == "populated" && get(t, filepath.Join(dir, "ticket.md")) != "old task" {
				t.Fatal("snapshot changed current ticket")
			}
			call(t, "task", "--expect", token, templates, root)
			if get(t, filepath.Join(dir, "ticket.md")) != "" {
				t.Fatal("guarded task did not initialize an empty ticket")
			}
		})
	}
}

func TestChangedSnapshotPreservesTaskAndGitRules(t *testing.T) {
	for _, change := range []string{"appeared", "removed", "content", "renamed", "directory"} {
		t.Run(change, func(t *testing.T) {
			root, templates := t.TempDir(), taskTemplates(t)
			git(t, "init", "-q", root)
			exclude := filepath.Join(root, ".git", "info", "exclude")
			put(t, exclude, "# Keep existing rule\n")
			dir := filepath.Join(root, ".inux")
			if change != "appeared" {
				if err := os.Mkdir(dir, 0700); err != nil {
					t.Fatal(err)
				}
				put(t, filepath.Join(dir, "ticket.md"), "old task")
			}
			token := snapshot(t, root)
			switch change {
			case "appeared":
				if err := os.Mkdir(dir, 0700); err != nil {
					t.Fatal(err)
				}
				put(t, filepath.Join(dir, "ticket.md"), "new task")
			case "removed":
				if err := os.RemoveAll(dir); err != nil {
					t.Fatal(err)
				}
			case "content":
				put(t, filepath.Join(dir, "ticket.md"), "new task")
			case "renamed":
				if err := os.Rename(filepath.Join(dir, "ticket.md"), filepath.Join(dir, "custom.md")); err != nil {
					t.Fatal(err)
				}
			case "directory":
				if err := os.Mkdir(filepath.Join(dir, "custom"), 0700); err != nil {
					t.Fatal(err)
				}
			}
			changed := snapshot(t, root)
			var out bytes.Buffer
			if err := app.Run([]string{"task", "--expect", token, templates, root}, &out); err == nil || out.Len() != 0 {
				t.Fatal("changed snapshot accepted or reported success")
			}
			if snapshot(t, root) != changed || get(t, exclude) != "# Keep existing rule\n" {
				t.Fatal("snapshot mismatch modified task or Git rules")
			}
		})
	}
}

func TestSnapshotBoundToProjectAndValidArguments(t *testing.T) {
	root, other, templates := t.TempDir(), t.TempDir(), taskTemplates(t)
	token := snapshot(t, root)
	for _, args := range [][]string{
		{"task", "--expect", token, templates, other},
		{"task", "--expect"},
		{"task", "--expect", "", templates, root},
		{"task", "--expect", strings.Repeat("z", 64), templates, root},
		{"task", "--snapshot", root, "extra"},
		{"task", "--snapshot", filepath.Join(root, "missing")},
	} {
		var out bytes.Buffer
		if err := app.Run(args, &out); err == nil || out.Len() != 0 {
			t.Fatal("accepted invalid snapshot arguments", args)
		}
	}
}

func TestSnapshotLinks(t *testing.T) {
	root, outside := t.TempDir(), t.TempDir()
	dir := filepath.Join(root, ".inux")
	put(t, filepath.Join(outside, "keep"), "untouched")
	put(t, dir, "not a directory")
	if err := app.Run([]string{"task", "--snapshot", root}, &bytes.Buffer{}); err == nil {
		t.Fatal("snapshot accepted a .inux file")
	}
	if err := os.Remove(dir); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, dir); err != nil {
		t.Skip(err)
	}
	if err := app.Run([]string{"task", "--snapshot", root}, &bytes.Buffer{}); err == nil {
		t.Fatal("snapshot accepted a .inux symlink")
	}
	if err := os.Remove(dir); err != nil {
		t.Fatal(err)
	}
	call(t, "task", taskTemplates(t), root)
	if err := os.Symlink(outside, filepath.Join(dir, "external")); err != nil {
		t.Fatal(err)
	}
	token := snapshot(t, root)
	put(t, filepath.Join(outside, "keep"), "changed outside")
	if snapshot(t, root) != token {
		t.Fatal("snapshot followed a nested link")
	}
	call(t, "task", "--expect", token, taskTemplates(t), root)
	if get(t, filepath.Join(outside, "keep")) != "changed outside" {
		t.Fatal("guarded task changed linked files")
	}
}

func TestSnapshotReadFailure(t *testing.T) {
	root := t.TempDir()
	call(t, "task", taskTemplates(t), root)
	path := filepath.Join(root, ".inux", "ticket.md")
	if err := os.Chmod(path, 0000); err != nil {
		t.Skip(err)
	}
	defer os.Chmod(path, 0600)
	if file, err := os.Open(path); err == nil {
		file.Close()
		t.Skip("host can read permission-restricted files")
	}
	var out bytes.Buffer
	if err := app.Run([]string{"task", "--snapshot", root}, &out); err == nil || out.Len() != 0 {
		t.Fatal("snapshot ignored an unreadable file")
	}
}
