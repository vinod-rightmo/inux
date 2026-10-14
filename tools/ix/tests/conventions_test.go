package tests

import (
	"bytes"
	"encoding/json"
	app "ix/internal"
	"strings"
	"testing"
)

func TestConventionChecks(t *testing.T) {
	cases := []struct {
		name, value, pattern string
		valid                bool
	}{
		{"branch policy", "bugfix/ABC-123-fix", `bugfix/ABC-[0-9]+-[a-z0-9-]+`, true},
		{"commit policy", "fix(ui): correct alignment", `fix\(ui\): [^\r\n]+`, true},
		{"mismatch", "wrong", `right`, false},
		{"prefix", "prefix-abc", `abc`, false},
		{"suffix", "abc-suffix", `abc`, false},
		{"longer alternative", "abc", `a|abc`, true},
		{"alternate prefix", "abc-suffix", `a|abc`, false},
		{"invalid pattern", "abc", `[`, false},
		{"unbalanced groups", "a", `a)|(?:b`, false},
		{"open literal quote", "foo", `\Qfoo`, true},
		{"case flag", "FIX", `(?i)fix`, true},
		{"unicode", "fix: Café 日本語", `fix: .+`, true},
		{"body allowed", "fix: align\n\nDetails", `(?s)fix: .+`, true},
		{"body not allowed", "fix: align\n\nDetails", `fix: [^\r\n]+`, false},
		{"multiline partial", "wrong\nfix: align", `(?m)^fix: .+$`, false},
		{"whitespace preserved", " fix: align", `fix: .+`, false},
		{"empty explicitly allowed", "", ``, true},
		{"empty disallowed", "", `.+`, false},
	}
	for _, command := range []string{"check-branch", "check-commit"} {
		for _, tc := range cases {
			t.Run(command+"/"+tc.name, func(t *testing.T) {
				t.Setenv("PATH", "")
				var out bytes.Buffer
				err := app.Run([]string{command, tc.pattern, tc.value}, &out)
				if (err == nil) != tc.valid {
					t.Fatalf("valid=%v, error=%v", tc.valid, err)
				}
				if !tc.valid {
					if out.Len() != 0 {
						t.Fatal("failed check wrote success output")
					}
					return
				}
				var result map[string]bool
				if err := json.Unmarshal(out.Bytes(), &result); err != nil || !result["ok"] {
					t.Fatal("missing JSON success", out.String(), err)
				}
			})
		}
		for _, args := range [][]string{{}, {"pattern"}} {
			var out bytes.Buffer
			if err := app.Run(append([]string{command}, args...), &out); err == nil || out.Len() != 0 {
				t.Fatal("accepted wrong argument count", command, args)
			}
		}
	}
}

func TestConventionBatches(t *testing.T) {
	cases := []struct {
		name, pattern string
		values        []string
		failed        string
	}{
		{"all valid", `fix: [^\r\n]+`, []string{"fix: align", "fix: tests", "fix: Café"}, ""},
		{"mixed", `fix: [^\r\n]+`, []string{"fix: align", "wrong", "fix: tests", "wrong"}, "[2 4]"},
		{"first and last", `a|abc`, []string{"prefix-abc", "abc", "abc-suffix"}, "[1 3]"},
		{"all invalid", `right`, []string{"wrong", "wrong"}, "[1 2]"},
		{"multiline values", `(?s)fix: .+`, []string{"fix: align\n\nBody", "fix: tests\n\nMore"}, ""},
		{"empty explicitly allowed", `.*`, []string{"", "abc"}, ""},
	}
	for _, command := range []string{"check-branch", "check-commit"} {
		for _, tc := range cases {
			t.Run(command+"/"+tc.name, func(t *testing.T) {
				t.Setenv("PATH", "")
				var out bytes.Buffer
				args := append([]string{command, tc.pattern}, tc.values...)
				err := app.Run(args, &out)
				if tc.failed == "" {
					if err != nil || out.String() != "{\"ok\":true}\n" {
						t.Fatal("batch did not pass", err, out.String())
					}
				} else if err == nil || !strings.Contains(err.Error(), tc.failed) || out.Len() != 0 {
					t.Fatal("missing batch failure positions", err, out.String())
				}
			})
		}
	}
}
