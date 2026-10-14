package tests

import (
	"bytes"
	app "ix/internal"
	"testing"
)

func TestVersion(t *testing.T) {
	previous := app.Version
	t.Cleanup(func() { app.Version = previous })
	app.Version = "0.0.1"
	t.Setenv("PATH", "")
	var out bytes.Buffer
	if err := app.Run([]string{"--version"}, &out); err != nil {
		t.Fatal(err)
	}
	if out.String() != "ix 0.0.1\n" {
		t.Fatal("version did not report the build value", out.String())
	}
	out.Reset()
	if err := app.Run([]string{"--version", "extra"}, &out); err == nil || out.Len() != 0 {
		t.Fatal("version accepted an extra argument", err, out.String())
	}
}
